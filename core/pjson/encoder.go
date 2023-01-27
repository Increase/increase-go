package pjson

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/tidwall/sjson"
)

type encoder struct {
	encoderCache sync.Map // map[reflect.Type]encoderFunc
}

type stringifier func(value reflect.Value) ([]byte, error)

type structFieldEncoder struct {
	parsedStructTag
	stringifier
	structFieldName string
}

const pjsonStructTag = "pjson"

func (e *encoder) newPrimitiveTypeEncoder(t reflect.Type) stringifier {
	switch t.Kind() {
	// Note that we could use `gsjon` to encode these types but it would complicate our
	// code more and this current code shouldn't cause any issues
	case reflect.String:
		return func(v reflect.Value) ([]byte, error) {
			return []byte(fmt.Sprintf("%q", v.String())), nil
		}
	case reflect.Bool:
		return func(v reflect.Value) ([]byte, error) {
			if v.Bool() {
				return []byte("true"), nil
			}
			return []byte("false"), nil
		}
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		return func(v reflect.Value) ([]byte, error) {
			return []byte(strconv.FormatInt(v.Int(), 10)), nil
		}
	case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return func(v reflect.Value) ([]byte, error) {
			return []byte(strconv.FormatUint(v.Uint(), 10)), nil
		}
	case reflect.Float32:
		return func(v reflect.Value) ([]byte, error) {
			return []byte(strconv.FormatFloat(v.Float(), 'f', -1, 32)), nil
		}
	case reflect.Float64:
		return func(v reflect.Value) ([]byte, error) {
			return []byte(strconv.FormatFloat(v.Float(), 'f', -1, 64)), nil
		}
	default:
		return func(v reflect.Value) ([]byte, error) {
			return nil, fmt.Errorf("unknown type received at primitive encoder: %s", t.String())
		}
	}
}

func (e *encoder) valueEncoder(v reflect.Value) stringifier {
	if !v.IsValid() {
		return func(value reflect.Value) ([]byte, error) { return nil, nil }
	}
	return e.typeEncoder(v.Type())
}

type parsedStructTag struct {
	name                 string
	required             bool
	storeExtraProperties bool
}

func parseStructTag(raw string) (tag parsedStructTag) {
	parts := strings.Split(raw, ",")
	if len(parts) > 0 {
		tag.name = parts[0]
		parts = parts[1:]
	}
	for len(parts) > 0 {
		switch parts[0] {
		case "required":
			tag.required = true
		case "extras":
			tag.storeExtraProperties = true
		}
		parts = parts[1:]
	}
	return
}

type sortStructFields []structFieldEncoder

func (s sortStructFields) Len() int {
	return len(s)
}
func (s sortStructFields) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s sortStructFields) Less(i, j int) bool {
	return s[i].name > s[j].name
}

func (e *encoder) newArrayTypeEncoder(t reflect.Type) stringifier {
	itemEncoder := e.typeEncoder(t.Elem())

	return func(value reflect.Value) ([]byte, error) {
		json := []byte("[]")
		for i := 0; i < value.Len(); i++ {
			var value, err = itemEncoder(value.Index(i))
			if err != nil {
				return nil, err
			} else if value == nil {
				// Assume that empty items should be inserted as `null` so that the output array
				// will be the same length as the input array
				value = []byte("null")
			}

			json, err = sjson.SetRawBytes(json, "-1", value)
			if err != nil {
				return nil, err
			}
		}

		return json, nil
	}
}

func (e *encoder) newStructTypeEncoder(t reflect.Type) stringifier {
	fieldEncoders := []structFieldEncoder{}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Anonymous {
			t := field.Type
			if t.Kind() == reflect.Pointer {
				t = t.Elem()
			}
			if !field.IsExported() && t.Kind() != reflect.Struct {
				// Ignore embedded fields of unexported non-struct types.
				continue
			}
			// Do not ignore embedded fields of unexported struct types
			// since they may have exported fields.
		}

		tag, ok := field.Tag.Lookup(pjsonStructTag)
		if !ok {
			continue
		}

		fieldEncoder := structFieldEncoder{parseStructTag(tag), e.typeEncoder(field.Type), field.Name}

		// We only want to support unexported fields if they're tagged with `extras` because
		// that field shouldn't be part of the public API.
		if !field.IsExported() && !fieldEncoder.storeExtraProperties {
			continue
		}

		fieldEncoders = append(fieldEncoders, fieldEncoder)
	}

	// Ensure deterministic output
	sort.Sort(sortStructFields(fieldEncoders))

	return func(value reflect.Value) ([]byte, error) {
		json := []byte("{}")
		for i, field := range fieldEncoders {
			// Special cased handling for struct fields that are used to store unknown properties
			if field.storeExtraProperties {
				fieldValue := value.Field(i)
				if fieldValue.Kind() == reflect.Pointer {
					fieldValue = fieldValue.Elem()
				}

				if fieldValue.Type().Kind() != reflect.Map {
					panic(fmt.Sprintf("Expected a map type for field %s because the `extra` tag is present but received %s", field.structFieldName, fieldValue.Type().Kind()))
				}

				var err error
				json, err = e.encodeMapEntries(json, fieldValue)
				if err != nil {
					return nil, err
				}

				continue
			}

			var key = field.name
			var value, err = field.stringifier(value.Field(i))
			if err != nil {
				return nil, err
			} else if value == nil {
				if field.required {
					value = []byte("null")
				} else {
					continue
				}
			}

			json, err = sjson.SetRawBytes(json, key, value)
			if err != nil {
				return nil, err
			}
		}

		return json, nil
	}
}

// Given a []byte of json (may either be an empty object or an object that already contains entries)
// encode all of the entries in the map to the json byte array.
func (e *encoder) encodeMapEntries(json []byte, v reflect.Value) ([]byte, error) {
	pairs := []mapPair{}
	keyEncoder := e.typeEncoder(v.Type().Key())
	iter := v.MapRange()
	for iter.Next() {
		var encodedKey []byte
		if iter.Key().Type().Kind() == reflect.String {
			encodedKey = []byte(iter.Key().String())
		} else {
			var err error
			encodedKey, err = keyEncoder(iter.Key())
			if err != nil {
				return nil, err
			}
		}

		pairs = append(pairs, mapPair{key: encodedKey, value: iter.Value()})
	}

	// Ensure deterministic output
	sort.Sort(sortMapPairs(pairs))

	elementEncoder := e.typeEncoder(v.Type().Elem())

	for _, p := range pairs {
		encodedValue, err := elementEncoder(p.value)
		if err != nil {
			return nil, err
		}

		json, err = sjson.SetRawBytes(json, string(p.key), encodedValue)
		if err != nil {
			return nil, err
		}
	}

	return json, nil
}

type mapPair struct {
	key   []byte
	value reflect.Value
}

type sortMapPairs []mapPair

func (a sortMapPairs) Len() int           { return len(a) }
func (a sortMapPairs) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortMapPairs) Less(i, j int) bool { return bytes.Compare(a[i].key, a[j].key) == 1 }

func (e *encoder) newMapEncoder(t reflect.Type) stringifier {
	return func(value reflect.Value) ([]byte, error) {
		json := []byte("{}")
		var err error
		json, err = e.encodeMapEntries(json, value)
		if err != nil {
			return nil, err
		}
		return json, nil
	}
}

func (e *encoder) newTypeEncoder(t reflect.Type) stringifier {
	switch t.Kind() {
	case reflect.Pointer:
		inner := t.Elem()

		innerEncoder := e.typeEncoder(inner)
		return func(v reflect.Value) ([]byte, error) {
			if !v.IsValid() || v.IsNil() {
				return nil, nil
			}
			return innerEncoder(v.Elem())
		}
	case reflect.Struct:
		return e.newStructTypeEncoder(t)
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		return e.newArrayTypeEncoder(t)
	case reflect.Map:
		return e.newMapEncoder(t)
	case reflect.Interface:
		return func(value reflect.Value) ([]byte, error) {
			value = value.Elem()
			return e.valueEncoder(value)(value)
		}
	default:
		return e.newPrimitiveTypeEncoder(t)
	}
}

func (e *encoder) typeEncoder(t reflect.Type) stringifier {
	if fi, ok := e.encoderCache.Load(t); ok {
		return fi.(stringifier)
	}

	// To deal with recursive types, populate the map with an
	// indirect func before we build it. This type waits on the
	// real func (f) to be ready and then calls it. This indirect
	// func is only used for recursive types.
	var (
		wg sync.WaitGroup
		f  stringifier
	)
	wg.Add(1)
	fi, loaded := e.encoderCache.LoadOrStore(t, stringifier(func(v reflect.Value) ([]byte, error) {
		wg.Wait()
		return f(v)
	}))
	if loaded {
		return fi.(stringifier)
	}

	// Compute the real encoder and replace the indirect func with it.
	f = e.newTypeEncoder(t)
	wg.Done()
	e.encoderCache.Store(t, f)
	return f
}

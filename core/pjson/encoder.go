package pjson

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"sync"

	"github.com/tidwall/sjson"
)

var encoders sync.Map // map[reflect.Type]encoderFunc

func Marshal(value interface{}) ([]byte, error) {
	e := &encoder{}
	val := reflect.ValueOf(value)
	if !val.IsValid() {
		return nil, nil
	}
	typ := val.Type()
	enc := e.typeEncoder(typ)
	return enc(val)
}

type encoder struct{}

type encoderFunc func(value reflect.Value) ([]byte, error)

type encoderField struct {
	tag parsedStructTag
	fn  encoderFunc
	idx []int
}

func (e *encoder) typeEncoder(t reflect.Type) encoderFunc {
	if fi, ok := encoders.Load(t); ok {
		return fi.(encoderFunc)
	}

	// To deal with recursive types, populate the map with an
	// indirect func before we build it. This type waits on the
	// real func (f) to be ready and then calls it. This indirect
	// func is only used for recursive types.
	var (
		wg sync.WaitGroup
		f  encoderFunc
	)
	wg.Add(1)
	fi, loaded := encoders.LoadOrStore(t, encoderFunc(func(v reflect.Value) ([]byte, error) {
		wg.Wait()
		return f(v)
	}))
	if loaded {
		return fi.(encoderFunc)
	}

	// Compute the real encoder and replace the indirect func with it.
	f = e.newTypeEncoder(t)
	wg.Done()
	encoders.Store(t, f)
	return f
}

func (e *encoder) newTypeEncoder(t reflect.Type) encoderFunc {
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
			if !value.IsValid() {
				return nil, nil
			}
			return e.typeEncoder(value.Type())(value)
		}
	default:
		return e.newPrimitiveTypeEncoder(t)
	}
}

func (e *encoder) newPrimitiveTypeEncoder(t reflect.Type) encoderFunc {
	switch t.Kind() {
	// Note that we could use `gjson` to encode these types but it would complicate our
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

func (e *encoder) newArrayTypeEncoder(t reflect.Type) encoderFunc {
	itemEncoder := e.typeEncoder(t.Elem())

	return func(value reflect.Value) ([]byte, error) {
		json := []byte("[]")
		for i := 0; i < value.Len(); i++ {
			var value, err = itemEncoder(value.Index(i))
			if err != nil {
				return nil, err
			}
			if value == nil {
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

func (e *encoder) newStructTypeEncoder(t reflect.Type) encoderFunc {
	encoderFields := []encoderField{}

	// This helper allows us to recursively collect field encoders into a flat
	// array. The parameter `index` keeps track of the access patterns necessary
	// to get to some field.
	var collectFieldEncoders func(r reflect.Type, index []int)
	collectFieldEncoders = func(r reflect.Type, index []int) {
		for i := 0; i < r.NumField(); i++ {
			idx := append(index, i)
			field := t.FieldByIndex(idx)
			// If this is an embedded struct, traverse one level deeper to extract
			// the fields and get their encoders as well.
			if field.Anonymous {
				collectFieldEncoders(field.Type, idx)
				continue
			}
			tag, ok := field.Tag.Lookup(pjsonStructTag)
			if !ok {
				continue
			}
			ptag := parseStructTag(tag)

			// We only want to support unexported fields if they're tagged with `extras` because
			// that field shouldn't be part of the public API.
			if !field.IsExported() && !ptag.storeExtraProperties {
				continue
			}
			encoderFields = append(encoderFields, encoderField{ptag, e.typeEncoder(field.Type), idx})
		}
	}
	collectFieldEncoders(t, []int{})

	// Ensure deterministic output by sorting by lexicographic order
	sort.Slice(encoderFields, func(i, j int) bool {
		if encoderFields[i].tag.storeExtraProperties {
			return false
		}
		if encoderFields[j].tag.storeExtraProperties {
			return true
		}
		return encoderFields[i].tag.name < encoderFields[j].tag.name
	})

	return func(value reflect.Value) (json []byte, err error) {
		json = []byte("{}")

		for _, ef := range encoderFields {
			field := value.FieldByIndex(ef.idx)
			enc, err := ef.fn(field)
			if err != nil {
				return nil, err
			}

			// Special cased handling for struct fields that are used to store unknown properties
			if ef.tag.storeExtraProperties {
				json, err = e.encodeMapEntries(json, field)
				if err != nil {
					return nil, err
				}
				continue
			}

			if enc == nil && ef.tag.required {
				enc = []byte("null")
			}
			if enc == nil {
				continue
			}

			json, err = sjson.SetRawBytes(json, ef.tag.name, enc)
			if err != nil {
				return nil, err
			}
		}

		return
	}
}

type mapPair struct {
	key   []byte
	value reflect.Value
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
	sort.Slice(pairs, func(i, j int) bool {
		return bytes.Compare(pairs[i].key, pairs[j].key) > 0
	})

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

func (e *encoder) newMapEncoder(t reflect.Type) encoderFunc {
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

package query

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

type encoder struct {
	settings     QuerySettings
	encoderCache sync.Map // map[reflect.Type]encoderFunc
}

type Pair struct {
	key   string
	value string
}
type stringifier func(key string, value reflect.Value) []Pair

type parsedStructTag struct {
	name      string
	omitempty bool
	inline    bool
}

func parseStructTag(raw string) (tag parsedStructTag) {
	parts := strings.Split(raw, ",")
	if len(parts) > 0 {
		tag.name = parts[0]
		parts = parts[1:]
	}
	for len(parts) > 0 {
		switch parts[0] {
		case "omitempty":
			tag.omitempty = true
		case "inline":
			tag.inline = true
		}
		parts = parts[1:]
	}
	return
}

type structField struct {
	parsedStructTag
	stringifier
}

func (e *encoder) newStructTypeEncoder(t reflect.Type) stringifier {
	fieldEncoders := map[int]structField{}

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
		} else if !field.IsExported() {
			// Ignore unexported fields
			continue
		}
		tag, ok := field.Tag.Lookup(queryStructTag)
		if !ok {
			continue
		}
		fieldEncoders[i] = structField{parseStructTag(tag), e.typeEncoder(field.Type)}
	}

	return func(key string, value reflect.Value) (pairs []Pair) {
		for i, field := range fieldEncoders {
			if field.omitempty {
				if !value.IsValid() || value.IsZero() {
					continue
				}
			}
			var subkey string = field.name
			if field.inline {
				subkey = key
			} else {
				subkey = e.renderKeyPath(key, subkey)
			}
			for _, pair := range field.stringifier(subkey, value.Field(i)) {
				if field.omitempty && len(pair.value) == 0 {
					continue
				}
				pairs = append(pairs, pair)
			}
		}
		return
	}
}

func (e *encoder) newMapEncoder(t reflect.Type) stringifier {
	keyEncoder := e.typeEncoder(t.Key())
	elementEncoder := e.typeEncoder(t.Elem())
	return func(key string, value reflect.Value) (pairs []Pair) {
		iter := value.MapRange()
		for iter.Next() {
			encodedKey := keyEncoder("", iter.Key())
			if len(encodedKey) != 1 {
				panic("Unexpected number of parts for encoded map key. Are you using a non-primitive for this map?")
			}
			subkey := encodedKey[0].value
			keyPath := e.renderKeyPath(key, subkey)
			pairs = append(pairs, elementEncoder(keyPath, iter.Value())...)
		}
		return
	}
}

func (e *encoder) renderKeyPath(key string, subkey string) string {
	if len(key) == 0 {
		return subkey
	}
	if e.settings.NestedFormat == NestedQueryFormatDots {
		return fmt.Sprintf("%s.%s", key, subkey)
	}
	return fmt.Sprintf("%s[%s]", key, subkey)
}

func (e *encoder) newArrayTypeEncoder(t reflect.Type) stringifier {
	switch e.settings.ArrayFormat {
	case ArrayQueryFormatComma:
		innerEncoder := e.typeEncoder(t.Elem())
		return func(key string, v reflect.Value) []Pair {
			elements := []string{}
			for i := 0; i < v.Len(); i++ {
				for _, pair := range innerEncoder("", v.Index(i)) {
					elements = append(elements, pair.value)
				}
			}
			if len(elements) == 0 {
				return []Pair{}
			}
			return []Pair{{key, strings.Join(elements, ",")}}
		}
	case ArrayQueryFormatRepeat:
		innerEncoder := e.typeEncoder(t.Elem())
		return func(key string, value reflect.Value) (pairs []Pair) {
			for i := 0; i < value.Len(); i++ {
				pairs = append(pairs, innerEncoder(key, value.Index(i))...)
			}
			return pairs
		}
	case ArrayQueryFormatIndices:
		panic("The array indices format is not supported yet")
	case ArrayQueryFormatBrackets:
		innerEncoder := e.typeEncoder(t.Elem())
		return func(key string, value reflect.Value) []Pair {
			pairs := []Pair{}
			for i := 0; i < value.Len(); i++ {
				pairs = append(pairs, innerEncoder(key+"[]", value)...)
			}
			return pairs
		}
	default:
		panic(fmt.Sprintf("Unknown ArrayFormat value: %s", e.settings.ArrayFormat))
	}
}

func (e *encoder) newPrimitiveTypeEncoder(t reflect.Type) stringifier {
	switch t.Kind() {
	case reflect.Pointer:
		inner := t.Elem()

		innerEncoder := e.newPrimitiveTypeEncoder(inner)
		return func(key string, v reflect.Value) []Pair {
			if !v.IsValid() || v.IsNil() {
				return nil
			}
			return innerEncoder(key, v.Elem())
		}
	case reflect.String:
		return func(key string, v reflect.Value) []Pair {
			return []Pair{{key, v.String()}}
		}
	case reflect.Bool:
		return func(key string, v reflect.Value) []Pair {
			if v.Bool() {
				return []Pair{{key, "true"}}
			}
			return []Pair{{key, "false"}}
		}
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		return func(key string, v reflect.Value) []Pair {
			return []Pair{{key, strconv.FormatInt(v.Int(), 10)}}
		}
	case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return func(key string, v reflect.Value) []Pair {
			return []Pair{{key, strconv.FormatUint(v.Uint(), 10)}}
		}
	case reflect.Float32, reflect.Float64:
		return func(key string, v reflect.Value) []Pair {
			return []Pair{{key, strconv.FormatFloat(v.Float(), 'f', -1, 64)}}
		}
	case reflect.Complex64, reflect.Complex128:
		bitSize := 64
		if t.Kind() == reflect.Complex128 {
			bitSize = 128
		}
		return func(key string, v reflect.Value) []Pair {
			return []Pair{{key, strconv.FormatComplex(v.Complex(), 'f', -1, bitSize)}}
		}
	default:
		return func(key string, v reflect.Value) []Pair {
			return nil
		}
	}
}

func (e *encoder) newTypeEncoder(t reflect.Type) stringifier {
	switch t.Kind() {
	case reflect.Pointer:
		encoder := e.typeEncoder(t.Elem())
		return func(key string, value reflect.Value) (pairs []Pair) {
			if !value.IsValid() || value.IsNil() {
				return
			}
			pairs = encoder(key, value.Elem())
			return
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
		return func(key string, value reflect.Value) []Pair {
			value = value.Elem()
			return e.valueEncoder(value)(key, value)
		}
	default:
		return e.newPrimitiveTypeEncoder(t)
	}
}

func (e *encoder) valueEncoder(v reflect.Value) stringifier {
	if !v.IsValid() {
		return func(key string, value reflect.Value) []Pair { return nil }
	}
	return e.typeEncoder(v.Type())
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
	fi, loaded := e.encoderCache.LoadOrStore(t, stringifier(func(key string, v reflect.Value) []Pair {
		wg.Wait()
		return f(key, v)
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

package pjson

import (
	"fmt"
	"reflect"
	"sync"
	"unsafe"

	"github.com/tidwall/gjson"
)

type decoder struct {
	decoderCache sync.Map // map[reflect.Type]decoderFunc
}

type decoderFunc func(node gjson.Result, value reflect.Value) error

type structFieldDecoder struct {
	parsedStructTag
	decoderFunc
	structFieldName string
	fieldType       reflect.Type
}

func (d *decoder) newTypeDecoder(t reflect.Type) decoderFunc {
	switch t.Kind() {
	case reflect.Pointer:
		inner := t.Elem()
		innerDecoder := d.typeDecoder(inner)

		return func(n gjson.Result, v reflect.Value) error {
			if !v.IsValid() {
				return fmt.Errorf("unexpected invalid reflection value %+#v", v)
			}

			newValue := reflect.New(inner).Elem()
			err := innerDecoder(n, newValue)
			if err != nil {
				return err
			}

			v.Set(newValue.Addr())
			return nil
		}
	case reflect.Struct:
		return d.newStructTypeDecoder(t)
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		return d.newArrayTypeDecoder(t)
	case reflect.Map:
		return d.newMapDecoder(t)
	case reflect.Interface:
		return func(node gjson.Result, value reflect.Value) error {
			value.Set(reflect.ValueOf(node.Value()))
			return nil
		}
	default:
		return d.newPrimitiveTypeDecoder(t)
	}
}

func (d *decoder) newMapDecoder(t reflect.Type) decoderFunc {
	keyType := t.Key()
	itemType := t.Elem()
	itemDecoder := d.typeDecoder(itemType)

	return func(node gjson.Result, value reflect.Value) (err error) {
		mapValue := reflect.MakeMapWithSize(t, len(node.Map()))

		node.ForEach(func(key, value gjson.Result) bool {
			// It's fine for us to just use `ValueOf` here because the key types will
			// always be primitive types so we don't need to decode it using the standard pattern
			keyValue := reflect.ValueOf(key.Value())
			if keyValue.Type() != keyType {
				err = fmt.Errorf("expected key type %v but got %v", keyType, keyValue.Type())
				return false
			}

			itemValue := reflect.New(itemType).Elem()
			err = itemDecoder(value, itemValue)
			if err != nil {
				return false
			}

			mapValue.SetMapIndex(keyValue, itemValue)
			return true
		})

		if err != nil {
			return err
		}

		value.Set(mapValue)
		return nil
	}
}

func (d *decoder) newArrayTypeDecoder(t reflect.Type) decoderFunc {
	itemDecoder := d.typeDecoder(t.Elem())

	return func(node gjson.Result, value reflect.Value) (err error) {
		arrayNode := node.Array()

		arrayValue := reflect.MakeSlice(reflect.SliceOf(t.Elem()), len(arrayNode), len(arrayNode))
		for i, itemNode := range arrayNode {
			err = itemDecoder(itemNode, arrayValue.Index(i))
			if err != nil {
				return err
			}
		}

		value.Set(arrayValue)
		return nil
	}
}

func (d *decoder) newStructTypeDecoder(t reflect.Type) decoderFunc {
	// map of json field name to struct field decoders
	fieldDecoders := map[string]structFieldDecoder{}

	var extraFieldsDecoder structFieldDecoder
	var extraFieldsItemsDecoderFunc decoderFunc
	var extraFieldsType reflect.Type
	var extraFieldsItemsType reflect.Type

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

		decoder := structFieldDecoder{parseStructTag(tag), d.typeDecoder(field.Type), field.Name, field.Type}

		if decoder.storeExtraProperties {
			if extraFieldsType != nil {
				panic("Multiple struct fields encountered with the `extra` tag set. Only a single field can be tagged with `extra` at once.")
			}

			if field.Type.Kind() == reflect.Pointer {
				extraFieldsType = field.Type.Elem()
			} else {
				extraFieldsType = field.Type
			}

			extraFieldsDecoder = decoder
			extraFieldsItemsType = extraFieldsType.Elem()
			extraFieldsItemsDecoderFunc = d.typeDecoder(extraFieldsItemsType)
			continue
		}

		// We only want to support unexported fields if they're tagged with `extras` because
		// that field shouldn't be part of the public API.
		if !field.IsExported() {
			continue
		}

		fieldDecoders[decoder.name] = decoder
	}

	return func(node gjson.Result, value reflect.Value) (err error) {
		mapNode := node.Map()
		extraFields := map[string]gjson.Result{}

		for fieldName, itemNode := range mapNode {
			decoder, ok := fieldDecoders[fieldName]
			if !ok {
				extraFields[fieldName] = itemNode
				continue
			}

			if decoder.storeExtraProperties {
				panic("Unexpected field decoder tagged with `extra`")
			}

			err = decoder.decoderFunc(itemNode, value.FieldByName(decoder.structFieldName))
			if err != nil {
				return err
			}
		}

		if len(extraFields) == 0 {
			return nil
		}

		mapValue := reflect.MakeMapWithSize(extraFieldsType, len(extraFields))

		// TODO: don't re-implement the map decoder
		for k, itemNode := range extraFields {
			keyVal := reflect.ValueOf(k)

			itemValue := reflect.New(extraFieldsItemsType).Elem()
			err := extraFieldsItemsDecoderFunc(itemNode, itemValue)
			if err != nil {
				return err
			}

			mapValue.SetMapIndex(keyVal, itemValue)
		}

		// We need to setup a new value like this so that we can safely call `Addr()`
		newValue := reflect.New(extraFieldsType).Elem()
		newValue.Set(mapValue)

		fieldValue := makeSetable(value.FieldByName(extraFieldsDecoder.structFieldName))
		if fieldValue.Kind() == reflect.Pointer {
			fieldValue.Set(newValue.Addr())
		} else {
			fieldValue.Set(newValue)
		}

		return nil
	}
}

// Given a reflect.Value that points to an unexported struct field, returns a new reflect.Value
// that will not panic when v.Set() is called.
//
// https://stackoverflow.com/questions/17981651/is-there-any-way-to-access-private-fields-of-a-struct-from-another-package#comment74658463_17982725
func makeSetable(v reflect.Value) reflect.Value {
	return reflect.NewAt(v.Type(), unsafe.Pointer(v.UnsafeAddr())).Elem()
}

func (d *decoder) newPrimitiveTypeDecoder(t reflect.Type) decoderFunc {
	switch t.Kind() {
	case reflect.String:
		return func(n gjson.Result, v reflect.Value) error {
			v.SetString(n.String())
			return nil
		}
	case reflect.Bool:
		return func(n gjson.Result, v reflect.Value) error {
			v.SetBool(n.Bool())
			return nil
		}
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		return func(n gjson.Result, v reflect.Value) error {
			v.SetInt(n.Int())
			return nil
		}
	case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return func(n gjson.Result, v reflect.Value) error {
			v.SetUint(n.Uint())
			return nil
		}
	case reflect.Float32, reflect.Float64:
		return func(n gjson.Result, v reflect.Value) error {
			v.SetFloat(n.Float())
			return nil
		}
	default:
		return func(node gjson.Result, v reflect.Value) error {
			return fmt.Errorf("unknown type received at primitive decoder: %s", t.String())
		}
	}
}

func (d *decoder) valueDecoder(v reflect.Value) decoderFunc {
	if !v.IsValid() {
		return func(node gjson.Result, value reflect.Value) error { return nil }
	}
	return d.typeDecoder(v.Type())
}

func (d *decoder) typeDecoder(t reflect.Type) decoderFunc {
	if fi, ok := d.decoderCache.Load(t); ok {
		return fi.(decoderFunc)
	}

	// To deal with recursive types, populate the map with an
	// indirect func before we build it. This type waits on the
	// real func (f) to be ready and then calls it. This indirect
	// func is only used for recursive types.
	var (
		wg sync.WaitGroup
		f  decoderFunc
	)
	wg.Add(1)
	fi, loaded := d.decoderCache.LoadOrStore(t, decoderFunc(func(node gjson.Result, v reflect.Value) error {
		wg.Wait()
		return f(node, v)
	}))
	if loaded {
		return fi.(decoderFunc)
	}

	// Compute the real decoder and replace the indirect func with it.
	f = d.newTypeDecoder(t)
	wg.Done()
	d.decoderCache.Store(t, f)
	return f
}

package core

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

func urlEncode(f reflect.Value) string {
	var value string
	switch f.Kind() {
	case reflect.String:
		value = fmt.Sprintf("%v", f.String())
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		value = fmt.Sprintf("%v", f.Int())
	case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		value = fmt.Sprintf("%v", f.Uint())
	case reflect.Bool:
		value = fmt.Sprintf("%v", f.Bool())
	case reflect.Float32, reflect.Float64:
		value = fmt.Sprintf("%v", f.Float())
	case reflect.Pointer:
		value = urlEncode(f.Elem())
	default:
		panic(fmt.Sprintf("%v is not supported yet %v", f.String(), f.Kind()))
	}
	return value
}

func MarshalIntoURLQuery(url *url.URL, value interface{}) {
	vo := reflect.ValueOf(value)
	if vo.IsZero() {
		return
	}
	rv := reflect.Indirect(vo)
	rt := rv.Type()
	query := url.Query()
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)

		if tag, ok := field.Tag.Lookup("json"); ok {
			if components := strings.Split(tag, ","); len(components) == 0 || components[0] == "_" {
				continue
			} else {
				omitempty := false
				for _, component := range components[1:] {
					omitempty = component == "omitempty"
					if omitempty {
						break
					}
				}
				f := rv.FieldByName(field.Name)
				if omitempty && f.IsZero() || f.IsNil() {
					continue
				}
				query.Add(components[0], urlEncode(f))
			}
		}
	}
	url.RawQuery = query.Encode()
}

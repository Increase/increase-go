package core

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

func LoadEnvironmentVariables(settings any) error {
	rv := reflect.ValueOf(settings).Elem()
	rt := rv.Type()
	for i := 0; i < rt.NumField(); i++ {
		fieldv := rv.Field(i)
		if !fieldv.IsZero() {
			continue
		}
		fieldt := rt.Field(i)
		if env, ok := fieldt.Tag.Lookup("env"); ok {
			if value, ok := os.LookupEnv(env); ok {
				if fieldt.Type.Kind() == reflect.String {
					fieldv.Set(reflect.ValueOf(value))
				} else {
					var parsed interface{}
					err := json.Unmarshal([]byte(value), &parsed)
					if err != nil {
						return fmt.Errorf("error while parsing env variable %s: %w", env, err)
					}
					pv := reflect.ValueOf(parsed)
					if pv.Kind() == fieldv.Kind() {
						fieldv.Set(pv.Convert(fieldv.Type()))
						continue
					} else {
						return fmt.Errorf("Unexpectedly found type %v (wanted %s) when parsing env %s", pv.Kind().String(), fieldv.Kind().String(), env)
					}
				}
			}
		}
		if defaultValue, ok := fieldt.Tag.Lookup("default"); ok {
			var parsed interface{}
			if fieldv.Kind() == reflect.String {
				parsed = defaultValue
			} else {
				err := json.Unmarshal([]byte(defaultValue), &parsed)
				if err != nil {
					return fmt.Errorf("error while parsing default value for field %s: %w", fieldt.Name, err)
				}
			}
			pv := reflect.ValueOf(parsed)
			if pv.Kind() == fieldv.Kind() {
				fieldv.Set(pv.Convert(fieldv.Type()))
			} else {
				return fmt.Errorf("Unexpectedly found type %v (wanted %s) when parsing default value for field %s", pv.Kind().String(), fieldv.Kind().String(), rv.Type().Field(i).Name)
			}
		}
	}
	return nil
}

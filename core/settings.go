package core

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

type environmentField struct {
	v reflect.Value
	t reflect.StructField
}

func (field environmentField) tag(name string) (string, bool) {
	return field.t.Tag.Lookup(name)
}

func (field environmentField) mustTag(name string) string {
	tag, ok := field.tag(name)
	if !ok {
		panic(fmt.Errorf("expected to find tag %s on type %v but got nothing", name, field.t))
	}
	return tag
}

func (field environmentField) checkEnv() (value string, exists bool) {
	env, exists := field.t.Tag.Lookup("env")
	if !exists {
		return
	}
	return os.LookupEnv(env)
}

func (field environmentField) store(value string) error {
	var parsed interface{}
	var pv reflect.Value
	if field.v.Kind() != reflect.String {
		err := json.Unmarshal([]byte(value), &parsed)
		if err != nil {
			// unmarshal failure
			return fmt.Errorf("error while parsing env variable %s: %w", field.mustTag("env"), err)
		}
		pv = reflect.ValueOf(parsed)
		if pv.Kind() != field.v.Kind() {
			// type mismatch
			return fmt.Errorf("unexpectedly found type %v (wanted %s) when parsing %s", pv.Kind().String(), field.v.Kind().String(), value)
		}
	} else {
		parsed = value
		pv = reflect.ValueOf(value)
	}
	field.v.Set(pv.Convert(field.v.Type()))
	return nil
}

func (field environmentField) attemptLoad() error {
	rawValue, ok := field.checkEnv()
	if !ok {
		// fallback to defaults
		rawValue, ok = field.tag("default")
	}
	if !ok {
		// TODO: fail?
		return nil
	}
	err := field.store(rawValue)
	if err != nil {
		return fmt.Errorf("error storing: %w", err)
	}
	return nil
}

type environmentLoader struct {
	v reflect.Value
	t reflect.Type
}

func (loader environmentLoader) field(index int) environmentField {
	return environmentField{loader.v.Field(index), loader.t.Field(index)}
}

// LoadEnvironmentVariables enumerates all field on the provided struct-value,
// attempting to set empty values to their environment values (when defined by the `env` struct tag),
// or from the default value (defined by the `default` tag).
//
// Values are parsed as JSON, except when the field is a string, in which case
// the value is assigned as-is.
//
// Examples:
//
//	type Settings struct {
//	  Name string `env:"NAME"`
//		Age int `env:"AGE"`
//	}
//
// NAME=Eric AGE=20 go run .
//
//	Settings.Name == "Eric"
//	Settings.Age  == 20
//
// NAME='"Eric"' AGE=20 go run .
//
//	Settings.Name == "\"Eric\""
//	Settings.Age  == 20
func LoadEnvironmentVariables(settings any) error {
	rv := reflect.ValueOf(settings).Elem()
	rt := rv.Type()
	loader := environmentLoader{rv, rt}
	for i := 0; i < rt.NumField(); i++ {
		field := loader.field(i)
		if !field.v.IsZero() {
			continue
		}
		if err := field.attemptLoad(); err != nil {
			return fmt.Errorf("error loading %s from environment/defaults: %w", field.t.Name, err)
		}
	}
	return nil
}

package pjson

import (
	"reflect"
	"sync"

	"github.com/tidwall/gjson"
)

var encoders sync.Map // map[reflect.Type]*encoder

func Marshal(value interface{}) ([]byte, error) {
	e := &encoder{encoders}
	reflectValue := reflect.ValueOf(value)
	return e.valueEncoder(reflectValue)(reflectValue)
}

var decoders sync.Map // map[reflect.Type]*decoder

func Unmarshal(raw []byte, to any) error {
	d := &decoder{decoders}
	value := reflect.ValueOf(to).Elem()
	result := gjson.ParseBytes(raw)
	return d.valueDecoder(value)(result, value)
}

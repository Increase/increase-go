package pjson

import (
	"math"
	"testing"

	"increase/core/pointers"
)

func serialize(v interface{}) (string, error) {
	raw, err := Marshal(v)
	if err != nil {
		return "", err
	}
	foo := string(raw[:])
	return foo, nil
}

func assertEncode(t *testing.T, v interface{}, expectation string) {
	if serialized, err := serialize(v); err != nil {
		t.Fatalf("serialization of %v failed with error %v", v, err)
	} else if serialized != expectation {
		t.Fatalf("expected %+#v to serialize to %s but got %s", v, expectation, serialized)
	}
}

// TODO: test extreme values

func TestPrimitiveTypes(t *testing.T) {
	assertEncode(t, true, "true")
	assertEncode(t, false, "false")
	assertEncode(t, nil, "")
	assertEncode(t, "a", "\"a\"")
	assertEncode(t, 1, "1")
	assertEncode(t, uint(1), "1")
	assertEncode(t, float32(1.34), "1.34")
	assertEncode(t, float64(1.34), "1.34")
	assertEncode(t, math.NaN(), "NaN")

	assertEncode(t, pointers.P(1.23456), "1.23456")
	assertEncode(t, pointers.P(true), "true")
	assertEncode(t, pointers.P(false), "false")
	assertEncode(t, pointers.P("a"), "\"a\"")
	assertEncode(t, pointers.P(1), "1")
	assertEncode(t, pointers.P(uint(1)), "1")
	assertEncode(t, pointers.P(float32(1.34)), "1.34")
	assertEncode(t, pointers.P(float64(1.34)), "1.34")
}

type BasicStruct struct {
	A interface{} `pjson:"a"`
}

func TestBasicStruct(t *testing.T) {
	assertEncode(t, BasicStruct{A: true}, "{\"a\":true}")
	assertEncode(t, BasicStruct{}, "{}")
	assertEncode(t, BasicStruct{A: pointers.P("foo")}, "{\"a\":\"foo\"}")
	assertEncode(t, BasicStruct{A: &[]int{1, 2}}, "{\"a\":[1,2]}")
}

type RequiredProperties struct {
	A interface{} `pjson:"a"`
	B interface{} `pjson:"B,required"`
}

func TestStructRequiredProperties(t *testing.T) {
	assertEncode(t, RequiredProperties{A: true, B: 1}, "{\"B\":1,\"a\":true}")
	assertEncode(t, RequiredProperties{A: true}, "{\"B\":null,\"a\":true}")
	assertEncode(t, RequiredProperties{}, "{\"B\":null}")
}

func TestArray(t *testing.T) {
	assertEncode(t, []int{1, 2}, "[1,2]")
	assertEncode(t, []bool{false, true, false}, "[false,true,false]")
	assertEncode(t, []BasicStruct{{A: false}}, "[{\"a\":false}]")
	assertEncode(t, []BasicStruct{{A: BasicStruct{A: "foo"}}}, "[{\"a\":{\"a\":\"foo\"}}]")
	assertEncode(t, &[]BasicStruct{{A: false}}, "[{\"a\":false}]")
}

func TestMap(t *testing.T) {
	assertEncode(t, map[string]interface{}{"hello": "world", "goodbye": true}, "{\"hello\":\"world\",\"goodbye\":true}")
	assertEncode(t, map[bool]interface{}{true: "value"}, "{\"true\":\"value\"}")
}

type StructWithExtras struct {
	A               interface{} `pjson:"a"`
	ExtraProperties JSONExtras  `pjson:"-,extras"`
}

func TestExtras(t *testing.T) {
	obj1 := StructWithExtras{A: true, ExtraProperties: JSONExtras{"foo": "bar"}}
	assertEncode(t, obj1, "{\"a\":true,\"foo\":\"bar\"}")

	obj2 := StructWithExtras{A: true, ExtraProperties: JSONExtras{"foo": true, "bar": "a"}}
	assertEncode(t, obj2, "{\"a\":true,\"foo\":true,\"bar\":\"a\"}")

	obj3 := StructWithExtras{A: true, ExtraProperties: JSONExtras{"foo": true, "bar": JSONExtras{"a": 1}}}
	assertEncode(t, obj3, "{\"a\":true,\"foo\":true,\"bar\":{\"a\":1}}")
}

type StructWithPrivateJsonFields struct {
	A          interface{} `pjson:"a"`
	jsonFields JSONExtras  `pjson:"-,extras"`
}

func TestPrivateJsonFields(t *testing.T) {
	obj := StructWithPrivateJsonFields{A: true, jsonFields: JSONExtras{"foo": "bar"}}
	assertEncode(t, obj, "{\"a\":true,\"foo\":\"bar\"}")
}

type StructWithPointers struct {
	A          *string    `pjson:"a"`
	B          *bool      `pjson:"b"`
	jsonFields JSONExtras `pjson:"-,extras"`
}

func TestStructWithPointers(t *testing.T) {
	assertEncode(t, StructWithPointers{A: pointers.P("Robert"), jsonFields: JSONExtras{"foo": "bar"}}, "{\"a\":\"Robert\",\"foo\":\"bar\"}")
	assertEncode(t, StructWithPointers{A: pointers.P("Robert"), B: pointers.P(false), jsonFields: JSONExtras{"foo": "bar"}}, "{\"a\":\"Robert\",\"b\":false,\"foo\":\"bar\"}")
}

type StructWithEmbedded struct {
	StructWithPrivateExtrasField
	A          *int       `pjson:"number2"`
	jsonFields JSONExtras `pjson:"-,extras"`
}

func TestStructWithEmbedded(t *testing.T) {
	assertEncode(t, StructWithEmbedded{
		StructWithPrivateExtrasField: StructWithPrivateExtrasField{
			A:          pointers.P(32),
			jsonFields: JSONExtras{"x1": "hi"},
		},
		A:          pointers.P(5),
		jsonFields: JSONExtras{"x2": "yo"},
	}, `{"number":32,"number2":5,"x1":"hi","x2":"yo"}`)
}

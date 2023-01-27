package pjson

import (
	"increase/core/pointers"
	"testing"
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
		t.Fatalf("expected %v to serialize to %s but got %s", v, expectation, serialized)
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
	assertEncode(t, RequiredProperties{A: true, B: 1}, "{\"a\":true,\"B\":1}")
	assertEncode(t, RequiredProperties{A: true}, "{\"a\":true,\"B\":null}")
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
	A               interface{}            `pjson:"a"`
	ExtraProperties map[string]interface{} `pjson:"-,extras"`
}

func TestExtras(t *testing.T) {
	assertEncode(t, StructWithExtras{A: true, ExtraProperties: map[string]interface{}{"foo": "bar"}}, "{\"a\":true,\"foo\":\"bar\"}")
	assertEncode(t, StructWithExtras{A: true, ExtraProperties: map[string]interface{}{"foo": true, "bar": "a"}}, "{\"a\":true,\"foo\":true,\"bar\":\"a\"}")
	assertEncode(t, StructWithExtras{A: true, ExtraProperties: map[string]interface{}{"foo": true, "bar": BasicStruct{A: 1}}}, "{\"a\":true,\"foo\":true,\"bar\":{\"a\":1}}")
}

type StructWithPrivateJsonFields struct {
	A          interface{}            `pjson:"a"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

func TestPrivateJsonFields(t *testing.T) {
	assertEncode(t, StructWithPrivateJsonFields{A: true, jsonFields: map[string]interface{}{"foo": "bar"}}, "{\"a\":true,\"foo\":\"bar\"}")
}

type StructWithPointers struct {
	A          *string                 `pjson:"a"`
	jsonFields *map[string]interface{} `pjson:"-,extras"`
}

func TestStructWithPointers(t *testing.T) {
	assertEncode(t, StructWithPointers{A: pointers.P("Robert"), jsonFields: &map[string]interface{}{"foo": "bar"}}, "{\"a\":\"Robert\",\"foo\":\"bar\"}")
}

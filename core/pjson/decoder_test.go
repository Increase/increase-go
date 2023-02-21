package pjson

import (
	"increase/core/pointers"
	"reflect"
	"testing"
)

func assertDecode[P any](t *testing.T, raw []byte, expectation P) {
	var result P
	if err := Unmarshal(raw, &result); err != nil {
		t.Fatalf("deserialization of %v failed with error %v", result, err)
	}

	if !reflect.DeepEqual(result, expectation) {
		t.Fatalf("expected to deserialize to %#v but got %#v", expectation, result)
	}
}

func assertError[P any](t *testing.T, raw []byte, expectErr string) {
	var result P
	err := Unmarshal(raw, &result)
	if err == nil {
		t.Fatalf("expected deserialize to fail with %q but it did not fail and deserialized to %v", expectErr, result)
	}

	if err.Error() != expectErr {
		t.Fatalf("expected deserialize to fail with %q but it failed with %q", expectErr, err)
	}
}

func TestPrimitive(t *testing.T) {
	assertDecode(t, []byte("true"), true)
	assertDecode(t, []byte("false"), false)

	assertDecode(t, []byte("1"), 1)
	assertDecode(t, []byte("12324"), 12324)
	assertDecode(t, []byte("\"65\""), 65)

	assertDecode(t, []byte("1"), uint(1))
	assertDecode(t, []byte("12324"), uint(12324))
	assertDecode(t, []byte("\"65\""), uint(65))

	assertDecode(t, []byte("1.54"), float32(1.54))
	assertDecode(t, []byte("1.89"), float64(1.89))
}

type PrimitveStruct struct {
	A bool    `pjson:"a"`
	B int     `pjson:"b"`
	C uint    `pjson:"c"`
	D float64 `pjson:"d"`
	E float32 `pjson:"e"`
}

func TestPrimitiveStruct(t *testing.T) {
	assertDecode(t, []byte("{\"a\":true}"), PrimitveStruct{A: true})
	assertDecode(t, []byte("{\"b\":true}"), PrimitveStruct{B: 1})
	assertDecode(t, []byte("{\"b\":165}"), PrimitveStruct{B: 165})
	assertDecode(t, []byte("{\"b\":\"185\"}"), PrimitveStruct{B: 185})
	assertDecode(t, []byte("{\"c\":0}"), PrimitveStruct{C: uint(0)})
	assertDecode(t, []byte("{\"c\":true}"), PrimitveStruct{C: uint(1)})
	assertDecode(t, []byte("{\"d\":55.79}"), PrimitveStruct{D: 55.79})
	assertDecode(t, []byte("{\"e\":746287.263572}"), PrimitveStruct{E: 746287.263572})
	assertDecode(
		t,
		[]byte("{\"a\":false,\"b\":237628372683,\"c\":654,\"d\":9999.43,\"e\":43.76}"),
		PrimitveStruct{
			A: false,
			B: 237628372683,
			C: uint(654),
			D: 9999.43,
			E: 43.76,
		},
	)
}

func TestPointers(t *testing.T) {
	assertDecode(t, []byte("true"), pointers.P(true))
	assertDecode(t, []byte("false"), pointers.P(false))
	assertDecode(t, []byte("165"), pointers.P(165))
	assertDecode(t, []byte("165.45"), pointers.P(165.45))
	assertDecode(t, []byte("\"foo\""), pointers.P("foo"))
}

type PrimitvePointerStruct struct {
	A *bool    `pjson:"a"`
	B *int     `pjson:"b"`
	C *uint    `pjson:"c"`
	D *float64 `pjson:"d"`
	E *float32 `pjson:"e"`
	F *[]int   `pjson:"f"`
}

func TestPrimitivePointerStruct(t *testing.T) {
	assertDecode(t, []byte("{\"a\":true}"), PrimitvePointerStruct{A: pointers.P(true)})
	assertDecode(t, []byte("{\"b\":true}"), PrimitvePointerStruct{B: pointers.P(1)})
	assertDecode(t, []byte("{\"b\":165}"), PrimitvePointerStruct{B: pointers.P(165)})
	assertDecode(t, []byte("{\"b\":\"185\"}"), PrimitvePointerStruct{B: pointers.P(185)})
	assertDecode(t, []byte("{\"c\":0}"), PrimitvePointerStruct{C: pointers.P(uint(0))})
	assertDecode(t, []byte("{\"c\":true}"), PrimitvePointerStruct{C: pointers.P(uint(1))})
	assertDecode(t, []byte("{\"d\":55.79}"), PrimitvePointerStruct{D: pointers.P(55.79)})
	assertDecode(t, []byte("{\"e\":746287.263572}"), PrimitvePointerStruct{E: pointers.P(float32(746287.263572))})
	assertDecode(
		t,
		[]byte("{\"a\":false,\"b\":237628372683,\"c\":654,\"d\":9999.43,\"e\":43.76,\"f\":[1,2,3,4,5]}"),
		PrimitvePointerStruct{
			A: pointers.P(false),
			B: pointers.P(237628372683),
			C: pointers.P(uint(654)),
			D: pointers.P(9999.43),
			E: pointers.P(float32(43.76)),
			F: &[]int{1, 2, 3, 4, 5},
		},
	)
}

type StructWithExtraProperties struct {
	A      bool       `pjson:"a"`
	Extras JSONExtras `pjson:"-,extras"`
}

func TestStructExtraProperties(t *testing.T) {
	obj1 := StructWithExtraProperties{
		A:      true,
		Extras: JSONExtras{"foo": true},
	}
	assertDecode(t, []byte("{\"a\": true, \"foo\": true}"), obj1)
	obj2 := StructWithExtraProperties{
		A: true,
		Extras: JSONExtras{
			"bar": "value",
			"foo": true,
		},
	}
	assertDecode(t, []byte("{\"a\": true, \"foo\": true, \"bar\": \"value\"}"), obj2)
}

func TestDecodeMap(t *testing.T) {
	assertError[map[int]string](t, []byte("{\"foo\": \"bar\", \"baz\": \"hello\"}"), "pjson: expected key type int but got string")
	assertDecode(t, []byte("{\"foo\": \"bar\"}"), map[string]string{"foo": "bar"})
}

func TestDecodeArray(t *testing.T) {
	assertDecode(t, []byte("[\"foo\", \"bar\"]"), []string{"foo", "bar"})
	assertDecode(t, []byte("[\"1\", 2]"), []int{1, 2})
}

type RecursiveStruct struct {
	Name  string           `pjson:"name"`
	Child *RecursiveStruct `pjson:"child"`
}

func TestRecursiveStruct(t *testing.T) {
	assertDecode(t, []byte("{\"name\": \"Robert\", \"child\": {\"name\":\"Alex\"}}"), RecursiveStruct{Name: "Robert", Child: &RecursiveStruct{Name: "Alex"}})
}

type StructWithPrivateExtrasField struct {
	A          *int       `pjson:"number"`
	jsonFields JSONExtras `pjson:"-,extras"`
}

func TestPrivateExtrasField(t *testing.T) {
	expect := StructWithPrivateExtrasField{A: pointers.P(100), jsonFields: JSONExtras{"foo": true}}
	assertDecode(t, []byte("{\"number\": 100, \"foo\": true}"), expect)

	extras := expect.jsonFields
	if extras["foo"] != true {
		t.Errorf("Expected extras[\"foo\"] to be true but got %v", extras["foo"])
	}
}

func TestStringCoercion(t *testing.T) {
	assertDecode(t, []byte("65"), "65")
	assertDecode(t, []byte("\"65\""), "65")
}

func TestStructWithEmbeddedDecode(t *testing.T) {
	assertDecode(t, []byte(`{"number":32,"x1":"hi","number2":5,"x2":"yo"}`), StructWithEmbedded{
		StructWithPrivateExtrasField: StructWithPrivateExtrasField{A: pointers.P(32)},

		A: pointers.P(5),
		jsonFields: JSONExtras{
			"x1": "hi",
			"x2": "yo",
		},
	})
}

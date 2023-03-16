package pjson

import (
	"reflect"
	"strings"
)

const pjsonStructTag = "pjson"
const formatStructTag = "format"

type parsedStructTag struct {
	name                 string
	required             bool
	storeExtraProperties bool
}

func parseStructTag(field reflect.StructField) (tag parsedStructTag, ok bool) {
	raw, ok := field.Tag.Lookup(pjsonStructTag)
	if !ok {
		return
	}
	parts := strings.Split(raw, ",")
	if len(parts) > 0 {
		tag.name = parts[0]
		parts = parts[1:]
	}
	for len(parts) > 0 {
		switch parts[0] {
		case "required":
			tag.required = true
		case "extras":
			tag.storeExtraProperties = true
		}
		parts = parts[1:]
	}
	return
}

func parseFormatStructTag(field reflect.StructField) (format string, ok bool) {
	format, ok = field.Tag.Lookup(formatStructTag)
	return
}

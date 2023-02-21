package pjson

import (
	"strings"
)

const pjsonStructTag = "pjson"

type parsedStructTag struct {
	name                 string
	required             bool
	storeExtraProperties bool
}

func parseStructTag(raw string) (tag parsedStructTag) {
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

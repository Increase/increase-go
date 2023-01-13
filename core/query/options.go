package query

type NestedQueryFormat string

const (
	NestedQueryFormatDots     NestedQueryFormat = "dots"
	NestedQueryFormatBrackets NestedQueryFormat = "brackets"
	NestedQueryFormatDefault  NestedQueryFormat = NestedQueryFormatBrackets
)

type ArrayQueryFormat string

const (
	ArrayQueryFormatComma    ArrayQueryFormat = "comma"
	ArrayQueryFormatRepeat   ArrayQueryFormat = "repeat"
	ArrayQueryFormatIndices  ArrayQueryFormat = "indices"
	ArrayQueryFormatBrackets ArrayQueryFormat = "brackets"
	ArrayQueryFormatDefault  ArrayQueryFormat = ArrayQueryFormatComma
)

type QuerySettings struct {
	NestedFormat NestedQueryFormat
	ArrayFormat  ArrayQueryFormat
}

func (settings *QuerySettings) assignDefaults() {
	if len(settings.NestedFormat) == 0 {
		settings.NestedFormat = NestedQueryFormatDefault
	}
	if len(settings.ArrayFormat) == 0 {
		settings.ArrayFormat = ArrayQueryFormatDefault
	}
}

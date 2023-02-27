package query

import (
	"net/url"
	"reflect"
	"sync"
)

const queryStructTag = "query"
const pathParamStructTag = "pathparam"

var encoders sync.Map // map[QuerySettings]*encoder

func MarshalWithSettings(value interface{}, settings QuerySettings) (v url.Values) {
	if value == nil {
		return
	}
	settings.assignDefaults()
	e, _ := encoders.LoadOrStore(settings, &encoder{settings, sync.Map{}})

	v = make(url.Values)

	rv := reflect.ValueOf(value)
	for _, pair := range e.(*encoder).valueEncoder(rv)("", rv) {
		v.Add(pair.key, pair.value)
	}

	return
}

func Marshal(value interface{}) url.Values {
	return MarshalWithSettings(value, QuerySettings{})
}

type Queryer interface {
	URLQuery() url.Values
}

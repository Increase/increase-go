package requests

import (
	"net/url"

	"github.com/increase/increase-go/core/field"
	"github.com/increase/increase-go/core/query"
)

type RoutingNumberListParams struct {
	// Return the page of entries after this one.
	Cursor field.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit field.Field[int64] `query:"limit"`
	// Filter financial institutions by routing number.
	RoutingNumber field.Field[string] `query:"routing_number,required"`
}

// URLQuery serializes RoutingNumberListParams into a url.Values of the query
// parameters associated with this value
func (r RoutingNumberListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

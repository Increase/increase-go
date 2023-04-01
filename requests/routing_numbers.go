package requests

import (
	"fmt"
	"net/url"

	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/fields"
)

type RoutingNumberListParams struct {
	// Return the page of entries after this one.
	Cursor fields.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit fields.Field[int64] `query:"limit"`
	// Filter financial institutions by routing number.
	RoutingNumber fields.Field[string] `query:"routing_number,required"`
}

// URLQuery serializes RoutingNumberListParams into a url.Values of the query
// parameters associated with this value
func (r *RoutingNumberListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r RoutingNumberListParams) String() (result string) {
	return fmt.Sprintf("&RoutingNumberListParams{Cursor:%s Limit:%s RoutingNumber:%s}", r.Cursor, r.Limit, r.RoutingNumber)
}

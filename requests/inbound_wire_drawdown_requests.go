package requests

import (
	"fmt"
	"net/url"

	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/fields"
)

type InboundWireDrawdownRequestListParams struct {
	// Return the page of entries after this one.
	Cursor fields.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit fields.Field[int64] `query:"limit"`
}

// URLQuery serializes InboundWireDrawdownRequestListParams into a url.Values of
// the query parameters associated with this value
func (r *InboundWireDrawdownRequestListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r InboundWireDrawdownRequestListParams) String() (result string) {
	return fmt.Sprintf("&InboundWireDrawdownRequestListParams{Cursor:%s Limit:%s}", r.Cursor, r.Limit)
}

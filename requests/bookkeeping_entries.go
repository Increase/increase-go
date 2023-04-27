package requests

import (
	"net/url"

	"github.com/increase/increase-go/internal/field"
	"github.com/increase/increase-go/internal/query"
)

type BookkeepingEntryListParams struct {
	// Return the page of entries after this one.
	Cursor field.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit field.Field[int64] `query:"limit"`
}

// URLQuery serializes BookkeepingEntryListParams into a url.Values of the query
// parameters associated with this value
func (r BookkeepingEntryListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

package requests

import (
	"fmt"
	"net/url"

	"github.com/increase/increase-go/core/fields"
	"github.com/increase/increase-go/core/query"
)

type OauthConnectionListParams struct {
	// Return the page of entries after this one.
	Cursor fields.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit fields.Field[int64] `query:"limit"`
}

// URLQuery serializes OauthConnectionListParams into a url.Values of the query
// parameters associated with this value
func (r *OauthConnectionListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r OauthConnectionListParams) String() (result string) {
	return fmt.Sprintf("&OauthConnectionListParams{Cursor:%s Limit:%s}", r.Cursor, r.Limit)
}

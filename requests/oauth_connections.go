package requests

import (
	"fmt"
	"net/url"
	"time"

	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/fields"
)

type OauthConnection struct {
	// The OAuth Connection's identifier.
	ID fields.Field[string] `json:"id,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp when the OAuth
	// Connection was created.
	CreatedAt fields.Field[time.Time] `json:"created_at,required" format:"date-time"`
	// The identifier of the Group that has authorized your OAuth application.
	GroupID fields.Field[string] `json:"group_id,required"`
	// Whether the connection is active.
	Status fields.Field[OauthConnectionStatus] `json:"status,required"`
	// A constant representing the object's type. For this resource it will always be
	// `oauth_connection`.
	Type fields.Field[OauthConnectionType] `json:"type,required"`
}

// MarshalJSON serializes OauthConnection into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *OauthConnection) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r OauthConnection) String() (result string) {
	return fmt.Sprintf("&OauthConnection{ID:%s CreatedAt:%s GroupID:%s Status:%s Type:%s}", r.ID, r.CreatedAt, r.GroupID, r.Status, r.Type)
}

type OauthConnectionStatus string

const (
	OauthConnectionStatusActive   OauthConnectionStatus = "active"
	OauthConnectionStatusInactive OauthConnectionStatus = "inactive"
)

type OauthConnectionType string

const (
	OauthConnectionTypeOauthConnection OauthConnectionType = "oauth_connection"
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

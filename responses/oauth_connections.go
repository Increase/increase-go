package responses

import (
	"time"

	pjson "github.com/increase/increase-go/core/json"
)

type OauthConnection struct {
	// The OAuth Connection's identifier.
	ID string `json:"id,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp when the OAuth
	// Connection was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The identifier of the Group that has authorized your OAuth application.
	GroupID string `json:"group_id,required"`
	// Whether the connection is active.
	Status OauthConnectionStatus `json:"status,required"`
	// A constant representing the object's type. For this resource it will always be
	// `oauth_connection`.
	Type OauthConnectionType `json:"type,required"`
	JSON OauthConnectionJSON
}

type OauthConnectionJSON struct {
	ID        pjson.Metadata
	CreatedAt pjson.Metadata
	GroupID   pjson.Metadata
	Status    pjson.Metadata
	Type      pjson.Metadata
	Raw       []byte
	Extras    map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into OauthConnection using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *OauthConnection) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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

type OauthConnectionListResponse struct {
	// The contents of the list.
	Data []OauthConnection `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       OauthConnectionListResponseJSON
}

type OauthConnectionListResponseJSON struct {
	Data       pjson.Metadata
	NextCursor pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into OauthConnectionListResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *OauthConnectionListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

package responses

import (
	"time"

	apijson "github.com/increase/increase-go/internal/json"
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
	ID        apijson.Metadata
	CreatedAt apijson.Metadata
	GroupID   apijson.Metadata
	Status    apijson.Metadata
	Type      apijson.Metadata
	raw       string
	Extras    map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into OauthConnection using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *OauthConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	Data       apijson.Metadata
	NextCursor apijson.Metadata
	raw        string
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into OauthConnectionListResponse
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *OauthConnectionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

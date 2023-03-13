package types

import (
	"fmt"
	"net/url"

	"github.com/increase/increase-go/core"
	"github.com/increase/increase-go/core/pjson"
	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/pagination"
)

type OauthConnection struct {
	// The OAuth Connection's identifier.
	ID *string `pjson:"id"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp when the OAuth
	// Connection was created.
	CreatedAt *string `pjson:"created_at"`
	// The identifier of the Group that has authorized your OAuth application.
	GroupID *string `pjson:"group_id"`
	// Whether the connection is active.
	Status *OauthConnectionStatus `pjson:"status"`
	// A constant representing the object's type. For this resource it will always be
	// `oauth_connection`.
	Type       *OauthConnectionType   `pjson:"type"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into OauthConnection using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *OauthConnection) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes OauthConnection into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *OauthConnection) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The OAuth Connection's identifier.
func (r *OauthConnection) GetID() (ID string) {
	if r != nil && r.ID != nil {
		ID = *r.ID
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp when the OAuth
// Connection was created.
func (r *OauthConnection) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The identifier of the Group that has authorized your OAuth application.
func (r *OauthConnection) GetGroupID() (GroupID string) {
	if r != nil && r.GroupID != nil {
		GroupID = *r.GroupID
	}
	return
}

// Whether the connection is active.
func (r *OauthConnection) GetStatus() (Status OauthConnectionStatus) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `oauth_connection`.
func (r *OauthConnection) GetType() (Type OauthConnectionType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

func (r OauthConnection) String() (result string) {
	return fmt.Sprintf("&OauthConnection{ID:%s CreatedAt:%s GroupID:%s Status:%s Type:%s}", core.FmtP(r.ID), core.FmtP(r.CreatedAt), core.FmtP(r.GroupID), core.FmtP(r.Status), core.FmtP(r.Type))
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
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit      *int64                 `query:"limit"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into OauthConnectionListParams
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *OauthConnectionListParams) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes OauthConnectionListParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *OauthConnectionListParams) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes OauthConnectionListParams into a url.Values of the query
// parameters associated with this value
func (r *OauthConnectionListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// Return the page of entries after this one.
func (r *OauthConnectionListParams) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *OauthConnectionListParams) GetLimit() (Limit int64) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

func (r OauthConnectionListParams) String() (result string) {
	return fmt.Sprintf("&OauthConnectionListParams{Cursor:%s Limit:%s}", core.FmtP(r.Cursor), core.FmtP(r.Limit))
}

type OauthConnectionList struct {
	// The contents of the list.
	Data *[]OauthConnection `pjson:"data"`
	// A pointer to a place in the list.
	NextCursor *string                `pjson:"next_cursor"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into OauthConnectionList using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *OauthConnectionList) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes OauthConnectionList into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *OauthConnectionList) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes OauthConnectionList into a url.Values of the query
// parameters associated with this value
func (r *OauthConnectionList) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// The contents of the list.
func (r *OauthConnectionList) GetData() (Data []OauthConnection) {
	if r != nil && r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r *OauthConnectionList) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

func (r OauthConnectionList) String() (result string) {
	return fmt.Sprintf("&OauthConnectionList{Data:%s NextCursor:%s}", core.Fmt(r.Data), core.FmtP(r.NextCursor))
}

type OauthConnectionsPage struct {
	*pagination.Page[OauthConnection]
}

func (r *OauthConnectionsPage) OauthConnection() *OauthConnection {
	return r.Current()
}

func (r *OauthConnectionsPage) NextPage() (*OauthConnectionsPage, error) {
	if page, err := r.Page.NextPage(); err != nil {
		return nil, err
	} else {
		return &OauthConnectionsPage{page}, nil
	}
}

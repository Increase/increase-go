package oauth_connections

type OauthConnection struct {
	// The OAuth Connection's identifier.
	Id *string `json:"id,omitempty"`

	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp when the OAuth
	// Connection was created.
	CreatedAt *string `json:"created_at,omitempty"`

	// The identifier of the Group that has authorized your OAuth application.
	GroupId *string `json:"group_id,omitempty"`

	// Whether the connection is active.
	Status *string `json:"status,omitempty"`

	// A constant representing the object's type. For this resource it will always be
	// `oauth_connection`.
	Type *string `json:"type,omitempty"`
}

// The OAuth Connection's identifier.
func (r *OauthConnection) GetId() (Id string) {
	if r != nil && r.Id != nil {
		Id = *r.Id
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
func (r *OauthConnection) GetGroupId() (GroupId string) {
	if r != nil && r.GroupId != nil {
		GroupId = *r.GroupId
	}
	return
}

// Whether the connection is active.
func (r *OauthConnection) GetStatus() (Status string) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `oauth_connection`.
func (r *OauthConnection) GetType() (Type string) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

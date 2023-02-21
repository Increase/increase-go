package types

import (
	"fmt"
	"increase/core"
	"increase/core/pjson"
	"increase/pagination"
)

type Document struct {
	// The Document identifier.
	ID *string `pjson:"id"`
	// The type of document.
	Category *DocumentCategory `pjson:"category"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the
	// Document was created.
	CreatedAt *string `pjson:"created_at"`
	// The identifier of the Entity the document was generated for.
	EntityID *string `pjson:"entity_id"`
	// The identifier of the File containing the Document's contents.
	FileID *string `pjson:"file_id"`
	// A constant representing the object's type. For this resource it will always be
	// `document`.
	Type       *DocumentType          `pjson:"type"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into Document using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Document) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes Document into an array of bytes using the gjson library.
// Members of the `jsonFields` field are serialized into the top-level, and will
// overwrite known members of the same name.
func (r *Document) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The Document identifier.
func (r *Document) GetID() (ID string) {
	if r != nil && r.ID != nil {
		ID = *r.ID
	}
	return
}

// The type of document.
func (r *Document) GetCategory() (Category DocumentCategory) {
	if r != nil && r.Category != nil {
		Category = *r.Category
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the
// Document was created.
func (r *Document) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The identifier of the Entity the document was generated for.
func (r *Document) GetEntityID() (EntityID string) {
	if r != nil && r.EntityID != nil {
		EntityID = *r.EntityID
	}
	return
}

// The identifier of the File containing the Document's contents.
func (r *Document) GetFileID() (FileID string) {
	if r != nil && r.FileID != nil {
		FileID = *r.FileID
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `document`.
func (r *Document) GetType() (Type DocumentType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

func (r Document) String() (result string) {
	return fmt.Sprintf("&Document{ID:%s Category:%s CreatedAt:%s EntityID:%s FileID:%s Type:%s}", core.FmtP(r.ID), core.FmtP(r.Category), core.FmtP(r.CreatedAt), core.FmtP(r.EntityID), core.FmtP(r.FileID), core.FmtP(r.Type))
}

type DocumentCategory string

const (
	DocumentCategoryForm_1099Int DocumentCategory = "form_1099_int"
)

type DocumentType string

const (
	DocumentTypeDocument DocumentType = "document"
)

type DocumentListParams struct {
	// Return the page of entries after this one.
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int64 `query:"limit"`
	// Filter Documents to ones belonging to the specified Entity.
	EntityID   *string                       `query:"entity_id"`
	Category   *DocumentsListParamsCategory  `query:"category"`
	CreatedAt  *DocumentsListParamsCreatedAt `query:"created_at"`
	jsonFields map[string]interface{}        `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into DocumentListParams using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *DocumentListParams) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes DocumentListParams into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *DocumentListParams) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// Return the page of entries after this one.
func (r *DocumentListParams) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *DocumentListParams) GetLimit() (Limit int64) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// Filter Documents to ones belonging to the specified Entity.
func (r *DocumentListParams) GetEntityID() (EntityID string) {
	if r != nil && r.EntityID != nil {
		EntityID = *r.EntityID
	}
	return
}

func (r *DocumentListParams) GetCategory() (Category DocumentsListParamsCategory) {
	if r != nil && r.Category != nil {
		Category = *r.Category
	}
	return
}

func (r *DocumentListParams) GetCreatedAt() (CreatedAt DocumentsListParamsCreatedAt) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

func (r DocumentListParams) String() (result string) {
	return fmt.Sprintf("&DocumentListParams{Cursor:%s Limit:%s EntityID:%s Category:%s CreatedAt:%s}", core.FmtP(r.Cursor), core.FmtP(r.Limit), core.FmtP(r.EntityID), r.Category, r.CreatedAt)
}

type DocumentsListParamsCategory struct {
	// Return results whose value is in the provided list. For GET requests, this
	// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
	In         *[]DocumentsListParamsCategoryIn `pjson:"in"`
	jsonFields map[string]interface{}           `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into DocumentsListParamsCategory
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *DocumentsListParamsCategory) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes DocumentsListParamsCategory into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *DocumentsListParamsCategory) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// Return results whose value is in the provided list. For GET requests, this
// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
func (r *DocumentsListParamsCategory) GetIn() (In []DocumentsListParamsCategoryIn) {
	if r != nil && r.In != nil {
		In = *r.In
	}
	return
}

func (r DocumentsListParamsCategory) String() (result string) {
	return fmt.Sprintf("&DocumentsListParamsCategory{In:%s}", core.Fmt(r.In))
}

type DocumentsListParamsCategoryIn string

const (
	DocumentsListParamsCategoryInForm_1099Int DocumentsListParamsCategoryIn = "form_1099_int"
)

type DocumentsListParamsCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After *string `pjson:"after"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before *string `pjson:"before"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter *string `pjson:"on_or_after"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore *string                `pjson:"on_or_before"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into DocumentsListParamsCreatedAt
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *DocumentsListParamsCreatedAt) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes DocumentsListParamsCreatedAt into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *DocumentsListParamsCreatedAt) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r *DocumentsListParamsCreatedAt) GetAfter() (After string) {
	if r != nil && r.After != nil {
		After = *r.After
	}
	return
}

// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r *DocumentsListParamsCreatedAt) GetBefore() (Before string) {
	if r != nil && r.Before != nil {
		Before = *r.Before
	}
	return
}

// Return results on or after this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r *DocumentsListParamsCreatedAt) GetOnOrAfter() (OnOrAfter string) {
	if r != nil && r.OnOrAfter != nil {
		OnOrAfter = *r.OnOrAfter
	}
	return
}

// Return results on or before this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r *DocumentsListParamsCreatedAt) GetOnOrBefore() (OnOrBefore string) {
	if r != nil && r.OnOrBefore != nil {
		OnOrBefore = *r.OnOrBefore
	}
	return
}

func (r DocumentsListParamsCreatedAt) String() (result string) {
	return fmt.Sprintf("&DocumentsListParamsCreatedAt{After:%s Before:%s OnOrAfter:%s OnOrBefore:%s}", core.FmtP(r.After), core.FmtP(r.Before), core.FmtP(r.OnOrAfter), core.FmtP(r.OnOrBefore))
}

type DocumentList struct {
	// The contents of the list.
	Data *[]Document `pjson:"data"`
	// A pointer to a place in the list.
	NextCursor *string                `pjson:"next_cursor"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into DocumentList using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *DocumentList) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes DocumentList into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *DocumentList) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The contents of the list.
func (r *DocumentList) GetData() (Data []Document) {
	if r != nil && r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r *DocumentList) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

func (r DocumentList) String() (result string) {
	return fmt.Sprintf("&DocumentList{Data:%s NextCursor:%s}", core.Fmt(r.Data), core.FmtP(r.NextCursor))
}

type DocumentsPage struct {
	*pagination.Page[Document]
}

func (r *DocumentsPage) Document() *Document {
	return r.Current()
}

func (r *DocumentsPage) NextPage() (*DocumentsPage, error) {
	if page, err := r.Page.NextPage(); err != nil {
		return nil, err
	} else {
		return &DocumentsPage{page}, nil
	}
}

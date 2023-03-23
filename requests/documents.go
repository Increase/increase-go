package requests

import (
	"fmt"
	"net/url"
	"time"

	"github.com/increase/increase-go/core"
	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/fields"
)

type Document struct {
	// The Document identifier.
	ID fields.Field[string] `json:"id,required"`
	// The type of document.
	Category fields.Field[DocumentCategory] `json:"category,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the
	// Document was created.
	CreatedAt fields.Field[time.Time] `json:"created_at,required" format:"date-time"`
	// The identifier of the Entity the document was generated for.
	EntityID fields.Field[string] `json:"entity_id,required,nullable"`
	// The identifier of the File containing the Document's contents.
	FileID fields.Field[string] `json:"file_id,required"`
	// A constant representing the object's type. For this resource it will always be
	// `document`.
	Type fields.Field[DocumentType] `json:"type,required"`
}

// MarshalJSON serializes Document into an array of bytes using the gjson library.
// Members of the `jsonFields` field are serialized into the top-level, and will
// overwrite known members of the same name.
func (r *Document) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r Document) String() (result string) {
	return fmt.Sprintf("&Document{ID:%s Category:%s CreatedAt:%s EntityID:%s FileID:%s Type:%s}", r.ID, r.Category, r.CreatedAt, r.EntityID, r.FileID, r.Type)
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
	Cursor fields.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit fields.Field[int64] `query:"limit"`
	// Filter Documents to ones belonging to the specified Entity.
	EntityID  fields.Field[string]                      `query:"entity_id"`
	Category  fields.Field[DocumentListParamsCategory]  `query:"category"`
	CreatedAt fields.Field[DocumentListParamsCreatedAt] `query:"created_at"`
}

// URLQuery serializes DocumentListParams into a url.Values of the query parameters
// associated with this value
func (r *DocumentListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r DocumentListParams) String() (result string) {
	return fmt.Sprintf("&DocumentListParams{Cursor:%s Limit:%s EntityID:%s Category:%s CreatedAt:%s}", r.Cursor, r.Limit, r.EntityID, r.Category, r.CreatedAt)
}

type DocumentListParamsCategory struct {
	// Return results whose value is in the provided list. For GET requests, this
	// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
	In fields.Field[[]DocumentListParamsCategoryIn] `query:"in"`
}

// URLQuery serializes DocumentListParamsCategory into a url.Values of the query
// parameters associated with this value
func (r *DocumentListParamsCategory) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r DocumentListParamsCategory) String() (result string) {
	return fmt.Sprintf("&DocumentListParamsCategory{In:%s}", core.Fmt(r.In))
}

type DocumentListParamsCategoryIn string

const (
	DocumentListParamsCategoryInForm_1099Int DocumentListParamsCategoryIn = "form_1099_int"
)

type DocumentListParamsCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After fields.Field[time.Time] `query:"after" format:"date-time"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before fields.Field[time.Time] `query:"before" format:"date-time"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter fields.Field[time.Time] `query:"on_or_after" format:"date-time"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore fields.Field[time.Time] `query:"on_or_before" format:"date-time"`
}

// URLQuery serializes DocumentListParamsCreatedAt into a url.Values of the query
// parameters associated with this value
func (r *DocumentListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r DocumentListParamsCreatedAt) String() (result string) {
	return fmt.Sprintf("&DocumentListParamsCreatedAt{After:%s Before:%s OnOrAfter:%s OnOrBefore:%s}", r.After, r.Before, r.OnOrAfter, r.OnOrBefore)
}

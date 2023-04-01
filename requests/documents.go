package requests

import (
	"fmt"
	"net/url"
	"time"

	"github.com/increase/increase-go/core"
	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/fields"
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

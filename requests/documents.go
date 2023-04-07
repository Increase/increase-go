package requests

import (
	"net/url"
	"time"

	"github.com/increase/increase-go/core/field"
	"github.com/increase/increase-go/core/query"
)

type DocumentListParams struct {
	// Return the page of entries after this one.
	Cursor field.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit field.Field[int64] `query:"limit"`
	// Filter Documents to ones belonging to the specified Entity.
	EntityID  field.Field[string]                      `query:"entity_id"`
	Category  field.Field[DocumentListParamsCategory]  `query:"category"`
	CreatedAt field.Field[DocumentListParamsCreatedAt] `query:"created_at"`
}

// URLQuery serializes DocumentListParams into a url.Values of the query parameters
// associated with this value
func (r *DocumentListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type DocumentListParamsCategory struct {
	// Return results whose value is in the provided list. For GET requests, this
	// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
	In field.Field[[]DocumentListParamsCategoryIn] `query:"in"`
}

// URLQuery serializes DocumentListParamsCategory into a url.Values of the query
// parameters associated with this value
func (r *DocumentListParamsCategory) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type DocumentListParamsCategoryIn string

const (
	DocumentListParamsCategoryInForm_1099Int DocumentListParamsCategoryIn = "form_1099_int"
)

type DocumentListParamsCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After field.Field[time.Time] `query:"after" format:"date-time"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before field.Field[time.Time] `query:"before" format:"date-time"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter field.Field[time.Time] `query:"on_or_after" format:"date-time"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore field.Field[time.Time] `query:"on_or_before" format:"date-time"`
}

// URLQuery serializes DocumentListParamsCreatedAt into a url.Values of the query
// parameters associated with this value
func (r *DocumentListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

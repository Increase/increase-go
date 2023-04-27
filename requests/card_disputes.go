package requests

import (
	"net/url"
	"time"

	"github.com/increase/increase-go/internal/field"
	apijson "github.com/increase/increase-go/internal/json"
	"github.com/increase/increase-go/internal/query"
)

type CardDisputeNewParams struct {
	// The Transaction you wish to dispute. This Transaction must have a `source_type`
	// of `card_settlement`.
	DisputedTransactionID field.Field[string] `json:"disputed_transaction_id,required"`
	// Why you are disputing this Transaction.
	Explanation field.Field[string] `json:"explanation,required"`
}

// MarshalJSON serializes CardDisputeNewParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r CardDisputeNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CardDisputeListParams struct {
	// Return the page of entries after this one.
	Cursor field.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit     field.Field[int64]                          `query:"limit"`
	CreatedAt field.Field[CardDisputeListParamsCreatedAt] `query:"created_at"`
	Status    field.Field[CardDisputeListParamsStatus]    `query:"status"`
}

// URLQuery serializes CardDisputeListParams into a url.Values of the query
// parameters associated with this value
func (r CardDisputeListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type CardDisputeListParamsCreatedAt struct {
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

// URLQuery serializes CardDisputeListParamsCreatedAt into a url.Values of the
// query parameters associated with this value
func (r CardDisputeListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type CardDisputeListParamsStatus struct {
	// Filter Card Disputes for those with the specified status or statuses. For GET
	// requests, this should be encoded as a comma-delimited string, such as
	// `?in=one,two,three`.
	In field.Field[[]CardDisputeListParamsStatusIn] `query:"in"`
}

// URLQuery serializes CardDisputeListParamsStatus into a url.Values of the query
// parameters associated with this value
func (r CardDisputeListParamsStatus) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type CardDisputeListParamsStatusIn string

const (
	CardDisputeListParamsStatusInPendingReviewing CardDisputeListParamsStatusIn = "pending_reviewing"
	CardDisputeListParamsStatusInAccepted         CardDisputeListParamsStatusIn = "accepted"
	CardDisputeListParamsStatusInRejected         CardDisputeListParamsStatusIn = "rejected"
)

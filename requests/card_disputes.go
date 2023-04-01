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

type CreateACardDisputeParameters struct {
	// The Transaction you wish to dispute. This Transaction must have a `source_type`
	// of `card_settlement`.
	DisputedTransactionID fields.Field[string] `json:"disputed_transaction_id,required"`
	// Why you are disputing this Transaction.
	Explanation fields.Field[string] `json:"explanation,required"`
}

// MarshalJSON serializes CreateACardDisputeParameters into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CreateACardDisputeParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CreateACardDisputeParameters) String() (result string) {
	return fmt.Sprintf("&CreateACardDisputeParameters{DisputedTransactionID:%s Explanation:%s}", r.DisputedTransactionID, r.Explanation)
}

type CardDisputeListParams struct {
	// Return the page of entries after this one.
	Cursor fields.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit     fields.Field[int64]                          `query:"limit"`
	CreatedAt fields.Field[CardDisputeListParamsCreatedAt] `query:"created_at"`
	Status    fields.Field[CardDisputeListParamsStatus]    `query:"status"`
}

// URLQuery serializes CardDisputeListParams into a url.Values of the query
// parameters associated with this value
func (r *CardDisputeListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r CardDisputeListParams) String() (result string) {
	return fmt.Sprintf("&CardDisputeListParams{Cursor:%s Limit:%s CreatedAt:%s Status:%s}", r.Cursor, r.Limit, r.CreatedAt, r.Status)
}

type CardDisputeListParamsCreatedAt struct {
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

// URLQuery serializes CardDisputeListParamsCreatedAt into a url.Values of the
// query parameters associated with this value
func (r *CardDisputeListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r CardDisputeListParamsCreatedAt) String() (result string) {
	return fmt.Sprintf("&CardDisputeListParamsCreatedAt{After:%s Before:%s OnOrAfter:%s OnOrBefore:%s}", r.After, r.Before, r.OnOrAfter, r.OnOrBefore)
}

type CardDisputeListParamsStatus struct {
	// Return results whose value is in the provided list. For GET requests, this
	// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
	In fields.Field[[]CardDisputeListParamsStatusIn] `query:"in"`
}

// URLQuery serializes CardDisputeListParamsStatus into a url.Values of the query
// parameters associated with this value
func (r *CardDisputeListParamsStatus) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r CardDisputeListParamsStatus) String() (result string) {
	return fmt.Sprintf("&CardDisputeListParamsStatus{In:%s}", core.Fmt(r.In))
}

type CardDisputeListParamsStatusIn string

const (
	CardDisputeListParamsStatusInPendingReviewing CardDisputeListParamsStatusIn = "pending_reviewing"
	CardDisputeListParamsStatusInAccepted         CardDisputeListParamsStatusIn = "accepted"
	CardDisputeListParamsStatusInRejected         CardDisputeListParamsStatusIn = "rejected"
)

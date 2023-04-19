package requests

import (
	"net/url"
	"time"

	"github.com/increase/increase-go/core/field"
	"github.com/increase/increase-go/core/query"
)

type PendingTransactionListParams struct {
	// Return the page of entries after this one.
	Cursor field.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit field.Field[int64] `query:"limit"`
	// Filter pending transactions to those belonging to the specified Account.
	AccountID field.Field[string] `query:"account_id"`
	// Filter pending transactions to those belonging to the specified Route.
	RouteID field.Field[string] `query:"route_id"`
	// Filter pending transactions to those caused by the specified source.
	SourceID  field.Field[string]                                `query:"source_id"`
	Status    field.Field[PendingTransactionListParamsStatus]    `query:"status"`
	CreatedAt field.Field[PendingTransactionListParamsCreatedAt] `query:"created_at"`
}

// URLQuery serializes PendingTransactionListParams into a url.Values of the query
// parameters associated with this value
func (r PendingTransactionListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type PendingTransactionListParamsStatus struct {
	// Filter Pending Transactions for those with the specified status. By default only
	// Pending Transactions in with status `pending` will be returned. For GET
	// requests, this should be encoded as a comma-delimited string, such as
	// `?in=one,two,three`.
	In field.Field[[]PendingTransactionListParamsStatusIn] `query:"in"`
}

// URLQuery serializes PendingTransactionListParamsStatus into a url.Values of the
// query parameters associated with this value
func (r PendingTransactionListParamsStatus) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type PendingTransactionListParamsStatusIn string

const (
	PendingTransactionListParamsStatusInPending  PendingTransactionListParamsStatusIn = "pending"
	PendingTransactionListParamsStatusInComplete PendingTransactionListParamsStatusIn = "complete"
)

type PendingTransactionListParamsCreatedAt struct {
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

// URLQuery serializes PendingTransactionListParamsCreatedAt into a url.Values of
// the query parameters associated with this value
func (r PendingTransactionListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

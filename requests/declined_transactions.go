package requests

import (
	"net/url"
	"time"

	"github.com/increase/increase-go/core/field"
	"github.com/increase/increase-go/core/query"
)

type DeclinedTransactionListParams struct {
	// Return the page of entries after this one.
	Cursor field.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit field.Field[int64] `query:"limit"`
	// Filter Declined Transactions to ones belonging to the specified Account.
	AccountID field.Field[string] `query:"account_id"`
	// Filter Declined Transactions to those belonging to the specified route.
	RouteID   field.Field[string]                                 `query:"route_id"`
	CreatedAt field.Field[DeclinedTransactionListParamsCreatedAt] `query:"created_at"`
}

// URLQuery serializes DeclinedTransactionListParams into a url.Values of the query
// parameters associated with this value
func (r *DeclinedTransactionListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type DeclinedTransactionListParamsCreatedAt struct {
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

// URLQuery serializes DeclinedTransactionListParamsCreatedAt into a url.Values of
// the query parameters associated with this value
func (r *DeclinedTransactionListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

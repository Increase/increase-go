package requests

import (
	"fmt"
	"net/url"
	"time"

	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/fields"
)

type DeclinedTransactionListParams struct {
	// Return the page of entries after this one.
	Cursor fields.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit fields.Field[int64] `query:"limit"`
	// Filter Declined Transactions to ones belonging to the specified Account.
	AccountID fields.Field[string] `query:"account_id"`
	// Filter Declined Transactions to those belonging to the specified route.
	RouteID   fields.Field[string]                                 `query:"route_id"`
	CreatedAt fields.Field[DeclinedTransactionListParamsCreatedAt] `query:"created_at"`
}

// URLQuery serializes DeclinedTransactionListParams into a url.Values of the query
// parameters associated with this value
func (r *DeclinedTransactionListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r DeclinedTransactionListParams) String() (result string) {
	return fmt.Sprintf("&DeclinedTransactionListParams{Cursor:%s Limit:%s AccountID:%s RouteID:%s CreatedAt:%s}", r.Cursor, r.Limit, r.AccountID, r.RouteID, r.CreatedAt)
}

type DeclinedTransactionListParamsCreatedAt struct {
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

// URLQuery serializes DeclinedTransactionListParamsCreatedAt into a url.Values of
// the query parameters associated with this value
func (r *DeclinedTransactionListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r DeclinedTransactionListParamsCreatedAt) String() (result string) {
	return fmt.Sprintf("&DeclinedTransactionListParamsCreatedAt{After:%s Before:%s OnOrAfter:%s OnOrBefore:%s}", r.After, r.Before, r.OnOrAfter, r.OnOrBefore)
}

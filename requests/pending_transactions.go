package requests

import (
	"net/url"

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
	SourceID field.Field[string]                             `query:"source_id"`
	Status   field.Field[PendingTransactionListParamsStatus] `query:"status"`
}

// URLQuery serializes PendingTransactionListParams into a url.Values of the query
// parameters associated with this value
func (r *PendingTransactionListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type PendingTransactionListParamsStatus struct {
	// Return results whose value is in the provided list. For GET requests, this
	// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
	In field.Field[[]PendingTransactionListParamsStatusIn] `query:"in"`
}

// URLQuery serializes PendingTransactionListParamsStatus into a url.Values of the
// query parameters associated with this value
func (r *PendingTransactionListParamsStatus) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type PendingTransactionListParamsStatusIn string

const (
	PendingTransactionListParamsStatusInPending  PendingTransactionListParamsStatusIn = "pending"
	PendingTransactionListParamsStatusInComplete PendingTransactionListParamsStatusIn = "complete"
)

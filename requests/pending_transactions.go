package requests

import (
	"fmt"
	"net/url"

	"github.com/increase/increase-go/core"
	"github.com/increase/increase-go/core/fields"
	"github.com/increase/increase-go/core/query"
)

type PendingTransactionListParams struct {
	// Return the page of entries after this one.
	Cursor fields.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit fields.Field[int64] `query:"limit"`
	// Filter pending transactions to those belonging to the specified Account.
	AccountID fields.Field[string] `query:"account_id"`
	// Filter pending transactions to those belonging to the specified Route.
	RouteID fields.Field[string] `query:"route_id"`
	// Filter pending transactions to those caused by the specified source.
	SourceID fields.Field[string]                             `query:"source_id"`
	Status   fields.Field[PendingTransactionListParamsStatus] `query:"status"`
}

// URLQuery serializes PendingTransactionListParams into a url.Values of the query
// parameters associated with this value
func (r *PendingTransactionListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r PendingTransactionListParams) String() (result string) {
	return fmt.Sprintf("&PendingTransactionListParams{Cursor:%s Limit:%s AccountID:%s RouteID:%s SourceID:%s Status:%s}", r.Cursor, r.Limit, r.AccountID, r.RouteID, r.SourceID, r.Status)
}

type PendingTransactionListParamsStatus struct {
	// Return results whose value is in the provided list. For GET requests, this
	// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
	In fields.Field[[]PendingTransactionListParamsStatusIn] `query:"in"`
}

// URLQuery serializes PendingTransactionListParamsStatus into a url.Values of the
// query parameters associated with this value
func (r *PendingTransactionListParamsStatus) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r PendingTransactionListParamsStatus) String() (result string) {
	return fmt.Sprintf("&PendingTransactionListParamsStatus{In:%s}", core.Fmt(r.In))
}

type PendingTransactionListParamsStatusIn string

const (
	PendingTransactionListParamsStatusInPending  PendingTransactionListParamsStatusIn = "pending"
	PendingTransactionListParamsStatusInComplete PendingTransactionListParamsStatusIn = "complete"
)

package requests

import (
	"net/url"
	"time"

	"github.com/increase/increase-go/internal/field"
	"github.com/increase/increase-go/internal/query"
)

type AccountStatementListParams struct {
	// Return the page of entries after this one.
	Cursor field.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit field.Field[int64] `query:"limit"`
	// Filter Account Statements to those belonging to the specified Account.
	AccountID            field.Field[string]                                         `query:"account_id"`
	StatementPeriodStart field.Field[AccountStatementListParamsStatementPeriodStart] `query:"statement_period_start"`
}

// URLQuery serializes AccountStatementListParams into a url.Values of the query
// parameters associated with this value
func (r AccountStatementListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type AccountStatementListParamsStatementPeriodStart struct {
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

// URLQuery serializes AccountStatementListParamsStatementPeriodStart into a
// url.Values of the query parameters associated with this value
func (r AccountStatementListParamsStatementPeriodStart) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

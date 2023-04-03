package requests

import (
	"fmt"
	"net/url"
	"time"

	"github.com/increase/increase-go/core/fields"
	"github.com/increase/increase-go/core/query"
)

type AccountStatementListParams struct {
	// Return the page of entries after this one.
	Cursor fields.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit fields.Field[int64] `query:"limit"`
	// Filter Account Statements to those belonging to the specified Account.
	AccountID            fields.Field[string]                                         `query:"account_id"`
	StatementPeriodStart fields.Field[AccountStatementListParamsStatementPeriodStart] `query:"statement_period_start"`
}

// URLQuery serializes AccountStatementListParams into a url.Values of the query
// parameters associated with this value
func (r *AccountStatementListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r AccountStatementListParams) String() (result string) {
	return fmt.Sprintf("&AccountStatementListParams{Cursor:%s Limit:%s AccountID:%s StatementPeriodStart:%s}", r.Cursor, r.Limit, r.AccountID, r.StatementPeriodStart)
}

type AccountStatementListParamsStatementPeriodStart struct {
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

// URLQuery serializes AccountStatementListParamsStatementPeriodStart into a
// url.Values of the query parameters associated with this value
func (r *AccountStatementListParamsStatementPeriodStart) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r AccountStatementListParamsStatementPeriodStart) String() (result string) {
	return fmt.Sprintf("&AccountStatementListParamsStatementPeriodStart{After:%s Before:%s OnOrAfter:%s OnOrBefore:%s}", r.After, r.Before, r.OnOrAfter, r.OnOrBefore)
}

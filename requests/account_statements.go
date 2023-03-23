package requests

import (
	"fmt"
	"net/url"
	"time"

	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/fields"
)

type AccountStatement struct {
	// The Account Statement identifier.
	ID fields.Field[string] `json:"id,required"`
	// The identifier for the Account this Account Statement belongs to.
	AccountID fields.Field[string] `json:"account_id,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Account
	// Statement was created.
	CreatedAt fields.Field[time.Time] `json:"created_at,required" format:"date-time"`
	// The identifier of the File containing a PDF of the statement.
	FileID fields.Field[string] `json:"file_id,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time representing the
	// start of the period the Account Statement covers.
	StatementPeriodStart fields.Field[time.Time] `json:"statement_period_start,required" format:"date-time"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time representing the end
	// of the period the Account Statement covers.
	StatementPeriodEnd fields.Field[time.Time] `json:"statement_period_end,required" format:"date-time"`
	// The Account's balance at the start of its statement period.
	StartingBalance fields.Field[int64] `json:"starting_balance,required"`
	// The Account's balance at the start of its statement period.
	EndingBalance fields.Field[int64] `json:"ending_balance,required"`
	// A constant representing the object's type. For this resource it will always be
	// `account_statement`.
	Type fields.Field[AccountStatementType] `json:"type,required"`
}

// MarshalJSON serializes AccountStatement into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *AccountStatement) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r AccountStatement) String() (result string) {
	return fmt.Sprintf("&AccountStatement{ID:%s AccountID:%s CreatedAt:%s FileID:%s StatementPeriodStart:%s StatementPeriodEnd:%s StartingBalance:%s EndingBalance:%s Type:%s}", r.ID, r.AccountID, r.CreatedAt, r.FileID, r.StatementPeriodStart, r.StatementPeriodEnd, r.StartingBalance, r.EndingBalance, r.Type)
}

type AccountStatementType string

const (
	AccountStatementTypeAccountStatement AccountStatementType = "account_statement"
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

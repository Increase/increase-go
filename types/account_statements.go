package types

import (
	"fmt"
	"net/url"
	"time"

	"github.com/increase/increase-go/core"
	"github.com/increase/increase-go/core/pjson"
	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/pagination"
)

type AccountStatement struct {
	// The Account Statement identifier.
	ID *string `pjson:"id"`
	// The identifier for the Account this Account Statement belongs to.
	AccountID *string `pjson:"account_id"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Account
	// Statement was created.
	CreatedAt *time.Time `pjson:"created_at" format:"2006-01-02T15:04:05Z07:00"`
	// The identifier of the File containing a PDF of the statement.
	FileID *string `pjson:"file_id"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time representing the
	// start of the period the Account Statement covers.
	StatementPeriodStart *time.Time `pjson:"statement_period_start" format:"2006-01-02T15:04:05Z07:00"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time representing the end
	// of the period the Account Statement covers.
	StatementPeriodEnd *time.Time `pjson:"statement_period_end" format:"2006-01-02T15:04:05Z07:00"`
	// The Account's balance at the start of its statement period.
	StartingBalance *int64 `pjson:"starting_balance"`
	// The Account's balance at the start of its statement period.
	EndingBalance *int64 `pjson:"ending_balance"`
	// A constant representing the object's type. For this resource it will always be
	// `account_statement`.
	Type       *AccountStatementType  `pjson:"type"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into AccountStatement using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *AccountStatement) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes AccountStatement into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *AccountStatement) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The Account Statement identifier.
func (r AccountStatement) GetID() (ID string) {
	if r.ID != nil {
		ID = *r.ID
	}
	return
}

// The identifier for the Account this Account Statement belongs to.
func (r AccountStatement) GetAccountID() (AccountID string) {
	if r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Account
// Statement was created.
func (r AccountStatement) GetCreatedAt() (CreatedAt time.Time) {
	if r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The identifier of the File containing a PDF of the statement.
func (r AccountStatement) GetFileID() (FileID string) {
	if r.FileID != nil {
		FileID = *r.FileID
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time representing the
// start of the period the Account Statement covers.
func (r AccountStatement) GetStatementPeriodStart() (StatementPeriodStart time.Time) {
	if r.StatementPeriodStart != nil {
		StatementPeriodStart = *r.StatementPeriodStart
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time representing the end
// of the period the Account Statement covers.
func (r AccountStatement) GetStatementPeriodEnd() (StatementPeriodEnd time.Time) {
	if r.StatementPeriodEnd != nil {
		StatementPeriodEnd = *r.StatementPeriodEnd
	}
	return
}

// The Account's balance at the start of its statement period.
func (r AccountStatement) GetStartingBalance() (StartingBalance int64) {
	if r.StartingBalance != nil {
		StartingBalance = *r.StartingBalance
	}
	return
}

// The Account's balance at the start of its statement period.
func (r AccountStatement) GetEndingBalance() (EndingBalance int64) {
	if r.EndingBalance != nil {
		EndingBalance = *r.EndingBalance
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `account_statement`.
func (r AccountStatement) GetType() (Type AccountStatementType) {
	if r.Type != nil {
		Type = *r.Type
	}
	return
}

func (r AccountStatement) String() (result string) {
	return fmt.Sprintf("&AccountStatement{ID:%s AccountID:%s CreatedAt:%s FileID:%s StatementPeriodStart:%s StatementPeriodEnd:%s StartingBalance:%s EndingBalance:%s Type:%s}", core.FmtP(r.ID), core.FmtP(r.AccountID), core.FmtP(r.CreatedAt), core.FmtP(r.FileID), core.FmtP(r.StatementPeriodStart), core.FmtP(r.StatementPeriodEnd), core.FmtP(r.StartingBalance), core.FmtP(r.EndingBalance), core.FmtP(r.Type))
}

type AccountStatementType string

const (
	AccountStatementTypeAccountStatement AccountStatementType = "account_statement"
)

type AccountStatementListParams struct {
	// Return the page of entries after this one.
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int64 `query:"limit"`
	// Filter Account Statements to those belonging to the specified Account.
	AccountID            *string                                         `query:"account_id"`
	StatementPeriodStart *AccountStatementListParamsStatementPeriodStart `query:"statement_period_start"`
	jsonFields           map[string]interface{}                          `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into AccountStatementListParams
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *AccountStatementListParams) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes AccountStatementListParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *AccountStatementListParams) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes AccountStatementListParams into a url.Values of the query
// parameters associated with this value
func (r *AccountStatementListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// Return the page of entries after this one.
func (r AccountStatementListParams) GetCursor() (Cursor string) {
	if r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r AccountStatementListParams) GetLimit() (Limit int64) {
	if r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// Filter Account Statements to those belonging to the specified Account.
func (r AccountStatementListParams) GetAccountID() (AccountID string) {
	if r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

func (r AccountStatementListParams) GetStatementPeriodStart() (StatementPeriodStart AccountStatementListParamsStatementPeriodStart) {
	if r.StatementPeriodStart != nil {
		StatementPeriodStart = *r.StatementPeriodStart
	}
	return
}

func (r AccountStatementListParams) String() (result string) {
	return fmt.Sprintf("&AccountStatementListParams{Cursor:%s Limit:%s AccountID:%s StatementPeriodStart:%s}", core.FmtP(r.Cursor), core.FmtP(r.Limit), core.FmtP(r.AccountID), r.StatementPeriodStart)
}

type AccountStatementListParamsStatementPeriodStart struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After *time.Time `pjson:"after" format:"2006-01-02T15:04:05Z07:00"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before *time.Time `pjson:"before" format:"2006-01-02T15:04:05Z07:00"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter *time.Time `pjson:"on_or_after" format:"2006-01-02T15:04:05Z07:00"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore *time.Time             `pjson:"on_or_before" format:"2006-01-02T15:04:05Z07:00"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// AccountStatementListParamsStatementPeriodStart using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *AccountStatementListParamsStatementPeriodStart) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes AccountStatementListParamsStatementPeriodStart into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *AccountStatementListParamsStatementPeriodStart) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes AccountStatementListParamsStatementPeriodStart into a
// url.Values of the query parameters associated with this value
func (r *AccountStatementListParamsStatementPeriodStart) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r AccountStatementListParamsStatementPeriodStart) GetAfter() (After time.Time) {
	if r.After != nil {
		After = *r.After
	}
	return
}

// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r AccountStatementListParamsStatementPeriodStart) GetBefore() (Before time.Time) {
	if r.Before != nil {
		Before = *r.Before
	}
	return
}

// Return results on or after this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r AccountStatementListParamsStatementPeriodStart) GetOnOrAfter() (OnOrAfter time.Time) {
	if r.OnOrAfter != nil {
		OnOrAfter = *r.OnOrAfter
	}
	return
}

// Return results on or before this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r AccountStatementListParamsStatementPeriodStart) GetOnOrBefore() (OnOrBefore time.Time) {
	if r.OnOrBefore != nil {
		OnOrBefore = *r.OnOrBefore
	}
	return
}

func (r AccountStatementListParamsStatementPeriodStart) String() (result string) {
	return fmt.Sprintf("&AccountStatementListParamsStatementPeriodStart{After:%s Before:%s OnOrAfter:%s OnOrBefore:%s}", core.FmtP(r.After), core.FmtP(r.Before), core.FmtP(r.OnOrAfter), core.FmtP(r.OnOrBefore))
}

type AccountStatementList struct {
	// The contents of the list.
	Data *[]AccountStatement `pjson:"data"`
	// A pointer to a place in the list.
	NextCursor *string                `pjson:"next_cursor"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into AccountStatementList using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *AccountStatementList) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes AccountStatementList into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *AccountStatementList) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes AccountStatementList into a url.Values of the query
// parameters associated with this value
func (r *AccountStatementList) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// The contents of the list.
func (r AccountStatementList) GetData() (Data []AccountStatement) {
	if r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r AccountStatementList) GetNextCursor() (NextCursor string) {
	if r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

func (r AccountStatementList) String() (result string) {
	return fmt.Sprintf("&AccountStatementList{Data:%s NextCursor:%s}", core.Fmt(r.Data), core.FmtP(r.NextCursor))
}

type AccountStatementsPage struct {
	*pagination.Page[AccountStatement]
}

func (r *AccountStatementsPage) AccountStatement() *AccountStatement {
	return r.Current()
}

func (r *AccountStatementsPage) NextPage() (*AccountStatementsPage, error) {
	if page, err := r.Page.NextPage(); err != nil {
		return nil, err
	} else {
		return &AccountStatementsPage{page}, nil
	}
}

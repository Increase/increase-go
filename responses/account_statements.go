package responses

import (
	"time"

	apijson "github.com/increase/increase-go/internal/json"
)

type AccountStatement struct {
	// The Account Statement identifier.
	ID string `json:"id,required"`
	// The identifier for the Account this Account Statement belongs to.
	AccountID string `json:"account_id,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Account
	// Statement was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The identifier of the File containing a PDF of the statement.
	FileID string `json:"file_id,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time representing the
	// start of the period the Account Statement covers.
	StatementPeriodStart time.Time `json:"statement_period_start,required" format:"date-time"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time representing the end
	// of the period the Account Statement covers.
	StatementPeriodEnd time.Time `json:"statement_period_end,required" format:"date-time"`
	// The Account's balance at the start of its statement period.
	StartingBalance int64 `json:"starting_balance,required"`
	// The Account's balance at the start of its statement period.
	EndingBalance int64 `json:"ending_balance,required"`
	// A constant representing the object's type. For this resource it will always be
	// `account_statement`.
	Type AccountStatementType `json:"type,required"`
	JSON AccountStatementJSON
}

type AccountStatementJSON struct {
	ID                   apijson.Metadata
	AccountID            apijson.Metadata
	CreatedAt            apijson.Metadata
	FileID               apijson.Metadata
	StatementPeriodStart apijson.Metadata
	StatementPeriodEnd   apijson.Metadata
	StartingBalance      apijson.Metadata
	EndingBalance        apijson.Metadata
	Type                 apijson.Metadata
	raw                  string
	Extras               map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AccountStatement using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *AccountStatement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStatementType string

const (
	AccountStatementTypeAccountStatement AccountStatementType = "account_statement"
)

type AccountStatementListResponse struct {
	// The contents of the list.
	Data []AccountStatement `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       AccountStatementListResponseJSON
}

type AccountStatementListResponseJSON struct {
	Data       apijson.Metadata
	NextCursor apijson.Metadata
	raw        string
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AccountStatementListResponse
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *AccountStatementListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

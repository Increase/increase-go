package responses

import (
	"time"

	apijson "github.com/increase/increase-go/internal/json"
)

type AccountNumber struct {
	// The identifier for the account this Account Number belongs to.
	AccountID string `json:"account_id,required"`
	// The account number.
	AccountNumber string `json:"account_number,required"`
	// The Account Number identifier.
	ID string `json:"id,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Account
	// Number was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The name you choose for the Account Number.
	Name string `json:"name,required"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber string `json:"routing_number,required"`
	// This indicates if payments can be made to the Account Number.
	Status AccountNumberStatus `json:"status,required"`
	// A constant representing the object's type. For this resource it will always be
	// `account_number`.
	Type AccountNumberType `json:"type,required"`
	JSON AccountNumberJSON
}

type AccountNumberJSON struct {
	AccountID     apijson.Metadata
	AccountNumber apijson.Metadata
	ID            apijson.Metadata
	CreatedAt     apijson.Metadata
	Name          apijson.Metadata
	RoutingNumber apijson.Metadata
	Status        apijson.Metadata
	Type          apijson.Metadata
	raw           string
	Extras        map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AccountNumber using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *AccountNumber) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountNumberStatus string

const (
	AccountNumberStatusActive   AccountNumberStatus = "active"
	AccountNumberStatusDisabled AccountNumberStatus = "disabled"
	AccountNumberStatusCanceled AccountNumberStatus = "canceled"
)

type AccountNumberType string

const (
	AccountNumberTypeAccountNumber AccountNumberType = "account_number"
)

type AccountNumberListResponse struct {
	// The contents of the list.
	Data []AccountNumber `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       AccountNumberListResponseJSON
}

type AccountNumberListResponseJSON struct {
	Data       apijson.Metadata
	NextCursor apijson.Metadata
	raw        string
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AccountNumberListResponse
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *AccountNumberListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

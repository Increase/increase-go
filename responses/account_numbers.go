package responses

import (
	"time"

	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/pagination"
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
	AccountID     pjson.Metadata
	AccountNumber pjson.Metadata
	ID            pjson.Metadata
	CreatedAt     pjson.Metadata
	Name          pjson.Metadata
	RoutingNumber pjson.Metadata
	Status        pjson.Metadata
	Type          pjson.Metadata
	Raw           []byte
	Extras        map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AccountNumber using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *AccountNumber) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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

type AccountNumberList struct {
	// The contents of the list.
	Data []AccountNumber `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       AccountNumberListJSON
}

type AccountNumberListJSON struct {
	Data       pjson.Metadata
	NextCursor pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AccountNumberList using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *AccountNumberList) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type AccountNumbersPage struct {
	*pagination.Page[AccountNumber]
}

func (r *AccountNumbersPage) AccountNumber() *AccountNumber {
	return r.Current()
}

func (r *AccountNumbersPage) NextPage() (*AccountNumbersPage, error) {
	if page, err := r.Page.NextPage(); err != nil {
		return nil, err
	} else {
		return &AccountNumbersPage{page}, nil
	}
}

package responses

import (
	"time"

	pjson "github.com/increase/increase-go/core/json"
)

type Account struct {
	// The Account's balances in the minor unit of its currency. For dollars, for
	// example, these values will represent cents.
	Balances AccountBalances `json:"balances,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Account
	// was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Account
	// currency.
	Currency AccountCurrency `json:"currency,required"`
	// The identifier for the Entity the Account belongs to.
	EntityID string `json:"entity_id,required,nullable"`
	// The identifier of an Entity that, while not owning the Account, is associated
	// with its activity.
	InformationalEntityID string `json:"informational_entity_id,required,nullable"`
	// The Account identifier.
	ID string `json:"id,required"`
	// The interest accrued but not yet paid, expressed as a string containing a
	// floating-point value.
	InterestAccrued string `json:"interest_accrued,required"`
	// The latest [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which
	// interest was accrued.
	InterestAccruedAt time.Time `json:"interest_accrued_at,required,nullable" format:"date"`
	// The name you choose for the Account.
	Name string `json:"name,required"`
	// The status of the Account.
	Status AccountStatus `json:"status,required"`
	// A constant representing the object's type. For this resource it will always be
	// `account`.
	Type AccountType `json:"type,required"`
	JSON AccountJSON
}

type AccountJSON struct {
	Balances              pjson.Metadata
	CreatedAt             pjson.Metadata
	Currency              pjson.Metadata
	EntityID              pjson.Metadata
	InformationalEntityID pjson.Metadata
	ID                    pjson.Metadata
	InterestAccrued       pjson.Metadata
	InterestAccruedAt     pjson.Metadata
	Name                  pjson.Metadata
	Status                pjson.Metadata
	Type                  pjson.Metadata
	Raw                   []byte
	Extras                map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Account using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Account) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type AccountBalances struct {
	// The Account's current balance, representing the sum of all posted Transactions
	// on the Account.
	CurrentBalance int64 `json:"current_balance,required"`
	// The Account's available balance, representing the current balance less any open
	// Pending Transactions on the Account.
	AvailableBalance int64 `json:"available_balance,required"`
	JSON             AccountBalancesJSON
}

type AccountBalancesJSON struct {
	CurrentBalance   pjson.Metadata
	AvailableBalance pjson.Metadata
	Raw              []byte
	Extras           map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AccountBalances using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *AccountBalances) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type AccountCurrency string

const (
	AccountCurrencyCad AccountCurrency = "CAD"
	AccountCurrencyChf AccountCurrency = "CHF"
	AccountCurrencyEur AccountCurrency = "EUR"
	AccountCurrencyGbp AccountCurrency = "GBP"
	AccountCurrencyJpy AccountCurrency = "JPY"
	AccountCurrencyUsd AccountCurrency = "USD"
)

type AccountStatus string

const (
	AccountStatusOpen   AccountStatus = "open"
	AccountStatusClosed AccountStatus = "closed"
)

type AccountType string

const (
	AccountTypeAccount AccountType = "account"
)

type AccountListResponse struct {
	// The contents of the list.
	Data []Account `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       AccountListResponseJSON
}

type AccountListResponseJSON struct {
	Data       pjson.Metadata
	NextCursor pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AccountListResponse using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *AccountListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

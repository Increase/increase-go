package responses

import (
	"time"

	apijson "github.com/increase/increase-go/internal/json"
)

type Account struct {
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
	CreatedAt             apijson.Metadata
	Currency              apijson.Metadata
	EntityID              apijson.Metadata
	InformationalEntityID apijson.Metadata
	ID                    apijson.Metadata
	InterestAccrued       apijson.Metadata
	InterestAccruedAt     apijson.Metadata
	Name                  apijson.Metadata
	Status                apijson.Metadata
	Type                  apijson.Metadata
	raw                   string
	Extras                map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Account using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Account) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	Data       apijson.Metadata
	NextCursor apijson.Metadata
	raw        string
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AccountListResponse using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *AccountListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

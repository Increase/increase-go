package requests

import (
	"fmt"
	"net/url"
	"time"

	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/fields"
)

type Account struct {
	// The Account's balances in the minor unit of its currency. For dollars, for
	// example, these values will represent cents.
	Balances fields.Field[AccountBalances] `json:"balances,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Account
	// was created.
	CreatedAt fields.Field[time.Time] `json:"created_at,required" format:"date-time"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Account
	// currency.
	Currency fields.Field[AccountCurrency] `json:"currency,required"`
	// The identifier for the Entity the Account belongs to.
	EntityID fields.Field[string] `json:"entity_id,required,nullable"`
	// The identifier of an Entity that, while not owning the Account, is associated
	// with its activity.
	InformationalEntityID fields.Field[string] `json:"informational_entity_id,required,nullable"`
	// The Account identifier.
	ID fields.Field[string] `json:"id,required"`
	// The interest accrued but not yet paid, expressed as a string containing a
	// floating-point value.
	InterestAccrued fields.Field[string] `json:"interest_accrued,required"`
	// The latest [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which
	// interest was accrued.
	InterestAccruedAt fields.Field[time.Time] `json:"interest_accrued_at,required,nullable" format:"date"`
	// The name you choose for the Account.
	Name fields.Field[string] `json:"name,required"`
	// The status of the Account.
	Status fields.Field[AccountStatus] `json:"status,required"`
	// A constant representing the object's type. For this resource it will always be
	// `account`.
	Type fields.Field[AccountType] `json:"type,required"`
}

// MarshalJSON serializes Account into an array of bytes using the gjson library.
// Members of the `jsonFields` field are serialized into the top-level, and will
// overwrite known members of the same name.
func (r *Account) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r Account) String() (result string) {
	return fmt.Sprintf("&Account{Balances:%s CreatedAt:%s Currency:%s EntityID:%s InformationalEntityID:%s ID:%s InterestAccrued:%s InterestAccruedAt:%s Name:%s Status:%s Type:%s}", r.Balances, r.CreatedAt, r.Currency, r.EntityID, r.InformationalEntityID, r.ID, r.InterestAccrued, r.InterestAccruedAt, r.Name, r.Status, r.Type)
}

type AccountBalances struct {
	// The Account's current balance, representing the sum of all posted Transactions
	// on the Account.
	CurrentBalance fields.Field[int64] `json:"current_balance,required"`
	// The Account's available balance, representing the current balance less any open
	// Pending Transactions on the Account.
	AvailableBalance fields.Field[int64] `json:"available_balance,required"`
}

// MarshalJSON serializes AccountBalances into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *AccountBalances) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r AccountBalances) String() (result string) {
	return fmt.Sprintf("&AccountBalances{CurrentBalance:%s AvailableBalance:%s}", r.CurrentBalance, r.AvailableBalance)
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

type CreateAnAccountParameters struct {
	// The identifier for the Entity that will own the Account.
	EntityID fields.Field[string] `json:"entity_id"`
	// The identifier for the Program that this Account falls under.
	ProgramID fields.Field[string] `json:"program_id"`
	// The identifier of an Entity that, while not owning the Account, is associated
	// with its activity. Its relationship to your group must be `informational`.
	InformationalEntityID fields.Field[string] `json:"informational_entity_id"`
	// The name you choose for the Account.
	Name fields.Field[string] `json:"name,required"`
}

// MarshalJSON serializes CreateAnAccountParameters into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CreateAnAccountParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CreateAnAccountParameters) String() (result string) {
	return fmt.Sprintf("&CreateAnAccountParameters{EntityID:%s ProgramID:%s InformationalEntityID:%s Name:%s}", r.EntityID, r.ProgramID, r.InformationalEntityID, r.Name)
}

type UpdateAnAccountParameters struct {
	// The new name of the Account.
	Name fields.Field[string] `json:"name"`
}

// MarshalJSON serializes UpdateAnAccountParameters into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *UpdateAnAccountParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r UpdateAnAccountParameters) String() (result string) {
	return fmt.Sprintf("&UpdateAnAccountParameters{Name:%s}", r.Name)
}

type AccountListParams struct {
	// Return the page of entries after this one.
	Cursor fields.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit fields.Field[int64] `query:"limit"`
	// Filter Accounts for those belonging to the specified Entity.
	EntityID fields.Field[string] `query:"entity_id"`
	// Filter Accounts for those with the specified status.
	Status fields.Field[AccountListParamsStatus] `query:"status"`
}

// URLQuery serializes AccountListParams into a url.Values of the query parameters
// associated with this value
func (r *AccountListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r AccountListParams) String() (result string) {
	return fmt.Sprintf("&AccountListParams{Cursor:%s Limit:%s EntityID:%s Status:%s}", r.Cursor, r.Limit, r.EntityID, r.Status)
}

type AccountListParamsStatus string

const (
	AccountListParamsStatusOpen   AccountListParamsStatus = "open"
	AccountListParamsStatusClosed AccountListParamsStatus = "closed"
)

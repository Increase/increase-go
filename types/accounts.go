package types

import (
	"fmt"
	"increase/core"
	"increase/core/pjson"
	"increase/core/query"
	"increase/pagination"
	"net/url"
)

type Account struct {
	// The Account's balances in the minor unit of its currency. For dollars, for
	// example, these values will represent cents.
	Balances *AccountBalances `pjson:"balances"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Account
	// was created.
	CreatedAt *string `pjson:"created_at"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Account
	// currency.
	Currency *AccountCurrency `pjson:"currency"`
	// The identifier for the Entity the Account belongs to.
	EntityID *string `pjson:"entity_id"`
	// The identifier of an Entity that, while not owning the Account, is associated
	// with its activity.
	InformationalEntityID *string `pjson:"informational_entity_id"`
	// The Account identifier.
	ID *string `pjson:"id"`
	// The interest accrued but not yet paid, expressed as a string containing a
	// floating-point value.
	InterestAccrued *string `pjson:"interest_accrued"`
	// The latest [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which
	// interest was accrued.
	InterestAccruedAt *string `pjson:"interest_accrued_at"`
	// The name you choose for the Account.
	Name *string `pjson:"name"`
	// The status of the Account.
	Status *AccountStatus `pjson:"status"`
	// A constant representing the object's type. For this resource it will always be
	// `account`.
	Type       *AccountType           `pjson:"type"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into Account using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Account) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes Account into an array of bytes using the gjson library.
// Members of the `jsonFields` field are serialized into the top-level, and will
// overwrite known members of the same name.
func (r *Account) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The Account's balances in the minor unit of its currency. For dollars, for
// example, these values will represent cents.
func (r *Account) GetBalances() (Balances AccountBalances) {
	if r != nil && r.Balances != nil {
		Balances = *r.Balances
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Account
// was created.
func (r *Account) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Account
// currency.
func (r *Account) GetCurrency() (Currency AccountCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The identifier for the Entity the Account belongs to.
func (r *Account) GetEntityID() (EntityID string) {
	if r != nil && r.EntityID != nil {
		EntityID = *r.EntityID
	}
	return
}

// The identifier of an Entity that, while not owning the Account, is associated
// with its activity.
func (r *Account) GetInformationalEntityID() (InformationalEntityID string) {
	if r != nil && r.InformationalEntityID != nil {
		InformationalEntityID = *r.InformationalEntityID
	}
	return
}

// The Account identifier.
func (r *Account) GetID() (ID string) {
	if r != nil && r.ID != nil {
		ID = *r.ID
	}
	return
}

// The interest accrued but not yet paid, expressed as a string containing a
// floating-point value.
func (r *Account) GetInterestAccrued() (InterestAccrued string) {
	if r != nil && r.InterestAccrued != nil {
		InterestAccrued = *r.InterestAccrued
	}
	return
}

// The latest [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which
// interest was accrued.
func (r *Account) GetInterestAccruedAt() (InterestAccruedAt string) {
	if r != nil && r.InterestAccruedAt != nil {
		InterestAccruedAt = *r.InterestAccruedAt
	}
	return
}

// The name you choose for the Account.
func (r *Account) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}

// The status of the Account.
func (r *Account) GetStatus() (Status AccountStatus) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `account`.
func (r *Account) GetType() (Type AccountType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

func (r Account) String() (result string) {
	return fmt.Sprintf("&Account{Balances:%s CreatedAt:%s Currency:%s EntityID:%s InformationalEntityID:%s ID:%s InterestAccrued:%s InterestAccruedAt:%s Name:%s Status:%s Type:%s}", r.Balances, core.FmtP(r.CreatedAt), core.FmtP(r.Currency), core.FmtP(r.EntityID), core.FmtP(r.InformationalEntityID), core.FmtP(r.ID), core.FmtP(r.InterestAccrued), core.FmtP(r.InterestAccruedAt), core.FmtP(r.Name), core.FmtP(r.Status), core.FmtP(r.Type))
}

type AccountBalances struct {
	// The Account's current balance, representing the sum of all posted Transactions
	// on the Account.
	CurrentBalance *int64 `pjson:"current_balance"`
	// The Account's available balance, representing the current balance less any open
	// Pending Transactions on the Account.
	AvailableBalance *int64                 `pjson:"available_balance"`
	jsonFields       map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into AccountBalances using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *AccountBalances) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes AccountBalances into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *AccountBalances) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The Account's current balance, representing the sum of all posted Transactions
// on the Account.
func (r *AccountBalances) GetCurrentBalance() (CurrentBalance int64) {
	if r != nil && r.CurrentBalance != nil {
		CurrentBalance = *r.CurrentBalance
	}
	return
}

// The Account's available balance, representing the current balance less any open
// Pending Transactions on the Account.
func (r *AccountBalances) GetAvailableBalance() (AvailableBalance int64) {
	if r != nil && r.AvailableBalance != nil {
		AvailableBalance = *r.AvailableBalance
	}
	return
}

func (r AccountBalances) String() (result string) {
	return fmt.Sprintf("&AccountBalances{CurrentBalance:%s AvailableBalance:%s}", core.FmtP(r.CurrentBalance), core.FmtP(r.AvailableBalance))
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
	EntityID *string `pjson:"entity_id"`
	// The identifier of an Entity that, while not owning the Account, is associated
	// with its activity. Its relationship to your group must be `informational`.
	InformationalEntityID *string `pjson:"informational_entity_id"`
	// The name you choose for the Account.
	Name       *string                `pjson:"name"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into CreateAnAccountParameters
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *CreateAnAccountParameters) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CreateAnAccountParameters into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CreateAnAccountParameters) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The identifier for the Entity that will own the Account.
func (r *CreateAnAccountParameters) GetEntityID() (EntityID string) {
	if r != nil && r.EntityID != nil {
		EntityID = *r.EntityID
	}
	return
}

// The identifier of an Entity that, while not owning the Account, is associated
// with its activity. Its relationship to your group must be `informational`.
func (r *CreateAnAccountParameters) GetInformationalEntityID() (InformationalEntityID string) {
	if r != nil && r.InformationalEntityID != nil {
		InformationalEntityID = *r.InformationalEntityID
	}
	return
}

// The name you choose for the Account.
func (r *CreateAnAccountParameters) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}

func (r CreateAnAccountParameters) String() (result string) {
	return fmt.Sprintf("&CreateAnAccountParameters{EntityID:%s InformationalEntityID:%s Name:%s}", core.FmtP(r.EntityID), core.FmtP(r.InformationalEntityID), core.FmtP(r.Name))
}

type UpdateAnAccountParameters struct {
	// The new name of the Account.
	Name       *string                `pjson:"name"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into UpdateAnAccountParameters
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *UpdateAnAccountParameters) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes UpdateAnAccountParameters into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *UpdateAnAccountParameters) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The new name of the Account.
func (r *UpdateAnAccountParameters) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}

func (r UpdateAnAccountParameters) String() (result string) {
	return fmt.Sprintf("&UpdateAnAccountParameters{Name:%s}", core.FmtP(r.Name))
}

type AccountListParams struct {
	// Return the page of entries after this one.
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int64 `query:"limit"`
	// Filter Accounts for those belonging to the specified Entity.
	EntityID *string `query:"entity_id"`
	// Filter Accounts for those with the specified status.
	Status     *AccountsListParamsStatus `query:"status"`
	jsonFields map[string]interface{}    `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into AccountListParams using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *AccountListParams) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes AccountListParams into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *AccountListParams) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes AccountListParams into a url.Values of the query parameters
// associated with this value
func (r *AccountListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// Return the page of entries after this one.
func (r *AccountListParams) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *AccountListParams) GetLimit() (Limit int64) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// Filter Accounts for those belonging to the specified Entity.
func (r *AccountListParams) GetEntityID() (EntityID string) {
	if r != nil && r.EntityID != nil {
		EntityID = *r.EntityID
	}
	return
}

// Filter Accounts for those with the specified status.
func (r *AccountListParams) GetStatus() (Status AccountsListParamsStatus) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

func (r AccountListParams) String() (result string) {
	return fmt.Sprintf("&AccountListParams{Cursor:%s Limit:%s EntityID:%s Status:%s}", core.FmtP(r.Cursor), core.FmtP(r.Limit), core.FmtP(r.EntityID), core.FmtP(r.Status))
}

type AccountsListParamsStatus string

const (
	AccountsListParamsStatusOpen   AccountsListParamsStatus = "open"
	AccountsListParamsStatusClosed AccountsListParamsStatus = "closed"
)

type AccountList struct {
	// The contents of the list.
	Data *[]Account `pjson:"data"`
	// A pointer to a place in the list.
	NextCursor *string                `pjson:"next_cursor"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into AccountList using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *AccountList) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes AccountList into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *AccountList) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes AccountList into a url.Values of the query parameters
// associated with this value
func (r *AccountList) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// The contents of the list.
func (r *AccountList) GetData() (Data []Account) {
	if r != nil && r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r *AccountList) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

func (r AccountList) String() (result string) {
	return fmt.Sprintf("&AccountList{Data:%s NextCursor:%s}", core.Fmt(r.Data), core.FmtP(r.NextCursor))
}

type AccountsPage struct {
	*pagination.Page[Account]
}

func (r *AccountsPage) Account() *Account {
	return r.Current()
}

func (r *AccountsPage) NextPage() (*AccountsPage, error) {
	if page, err := r.Page.NextPage(); err != nil {
		return nil, err
	} else {
		return &AccountsPage{page}, nil
	}
}

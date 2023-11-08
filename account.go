// File generated from our OpenAPI spec by Stainless.

package increase

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/apiquery"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/internal/shared"
	"github.com/increase/increase-go/option"
)

// AccountService contains methods and other services that help with interacting
// with the increase API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewAccountService] method instead.
type AccountService struct {
	Options []option.RequestOption
}

// NewAccountService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccountService(opts ...option.RequestOption) (r *AccountService) {
	r = &AccountService{}
	r.Options = opts
	return
}

// Create an Account
func (r *AccountService) New(ctx context.Context, body AccountNewParams, opts ...option.RequestOption) (res *Account, err error) {
	opts = append(r.Options[:], opts...)
	path := "accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve an Account
func (r *AccountService) Get(ctx context.Context, accountID string, opts ...option.RequestOption) (res *Account, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an Account
func (r *AccountService) Update(ctx context.Context, accountID string, body AccountUpdateParams, opts ...option.RequestOption) (res *Account, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List Accounts
func (r *AccountService) List(ctx context.Context, query AccountListParams, opts ...option.RequestOption) (res *shared.Page[Account], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "accounts"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List Accounts
func (r *AccountService) ListAutoPaging(ctx context.Context, query AccountListParams, opts ...option.RequestOption) *shared.PageAutoPager[Account] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Retrieve an Account Balance
func (r *AccountService) Balance(ctx context.Context, accountID string, query AccountBalanceParams, opts ...option.RequestOption) (res *BalanceLookup, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/balance", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Close an Account
func (r *AccountService) Close(ctx context.Context, accountID string, opts ...option.RequestOption) (res *Account, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/close", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Accounts are your bank accounts with Increase. They store money, receive
// transfers, and send payments. They earn interest and have depository insurance.
type Account struct {
	// The Account identifier.
	ID string `json:"id,required"`
	// The bank the Account is with.
	Bank AccountBank `json:"bank,required"`
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
	// The interest accrued but not yet paid, expressed as a string containing a
	// floating-point value.
	InterestAccrued string `json:"interest_accrued,required"`
	// The latest [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which
	// interest was accrued.
	InterestAccruedAt time.Time `json:"interest_accrued_at,required,nullable" format:"date"`
	// The Interest Rate currently being earned on the account, as a string containing
	// a decimal number. For example, a 1% interest rate would be represented as
	// "0.01".
	InterestRate string `json:"interest_rate,required"`
	// The name you choose for the Account.
	Name string `json:"name,required"`
	// The status of the Account.
	Status AccountStatus `json:"status,required"`
	// A constant representing the object's type. For this resource it will always be
	// `account`.
	Type AccountType `json:"type,required"`
	JSON accountJSON
}

// accountJSON contains the JSON metadata for the struct [Account]
type accountJSON struct {
	ID                    apijson.Field
	Bank                  apijson.Field
	CreatedAt             apijson.Field
	Currency              apijson.Field
	EntityID              apijson.Field
	InformationalEntityID apijson.Field
	InterestAccrued       apijson.Field
	InterestAccruedAt     apijson.Field
	InterestRate          apijson.Field
	Name                  apijson.Field
	Status                apijson.Field
	Type                  apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *Account) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The bank the Account is with.
type AccountBank string

const (
	// Blue Ridge Bank, N.A.
	AccountBankBlueRidgeBank AccountBank = "blue_ridge_bank"
	// First Internet Bank of Indiana
	AccountBankFirstInternetBank AccountBank = "first_internet_bank"
)

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Account
// currency.
type AccountCurrency string

const (
	// Canadian Dollar (CAD)
	AccountCurrencyCad AccountCurrency = "CAD"
	// Swiss Franc (CHF)
	AccountCurrencyChf AccountCurrency = "CHF"
	// Euro (EUR)
	AccountCurrencyEur AccountCurrency = "EUR"
	// British Pound (GBP)
	AccountCurrencyGbp AccountCurrency = "GBP"
	// Japanese Yen (JPY)
	AccountCurrencyJpy AccountCurrency = "JPY"
	// US Dollar (USD)
	AccountCurrencyUsd AccountCurrency = "USD"
)

// The status of the Account.
type AccountStatus string

const (
	// Open Accounts that are ready to use.
	AccountStatusOpen AccountStatus = "open"
	// Closed Accounts on which no new activity can occur.
	AccountStatusClosed AccountStatus = "closed"
)

// A constant representing the object's type. For this resource it will always be
// `account`.
type AccountType string

const (
	AccountTypeAccount AccountType = "account"
)

// Represents a request to lookup the balance of an Account at a given point in
// time.
type BalanceLookup struct {
	// The identifier for the account for which the balance was queried.
	AccountID string `json:"account_id,required"`
	// The Account's available balance, representing the current balance less any open
	// Pending Transactions on the Account.
	AvailableBalance int64 `json:"available_balance,required"`
	// The Account's current balance, representing the sum of all posted Transactions
	// on the Account.
	CurrentBalance int64 `json:"current_balance,required"`
	// A constant representing the object's type. For this resource it will always be
	// `balance_lookup`.
	Type BalanceLookupType `json:"type,required"`
	JSON balanceLookupJSON
}

// balanceLookupJSON contains the JSON metadata for the struct [BalanceLookup]
type balanceLookupJSON struct {
	AccountID        apijson.Field
	AvailableBalance apijson.Field
	CurrentBalance   apijson.Field
	Type             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *BalanceLookup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A constant representing the object's type. For this resource it will always be
// `balance_lookup`.
type BalanceLookupType string

const (
	BalanceLookupTypeBalanceLookup BalanceLookupType = "balance_lookup"
)

type AccountNewParams struct {
	// The name you choose for the Account.
	Name param.Field[string] `json:"name,required"`
	// The identifier for the Entity that will own the Account.
	EntityID param.Field[string] `json:"entity_id"`
	// The identifier of an Entity that, while not owning the Account, is associated
	// with its activity. Its relationship to your group must be `informational`.
	InformationalEntityID param.Field[string] `json:"informational_entity_id"`
	// The identifier for the Program that this Account falls under. Required if you
	// operate more than one Program.
	ProgramID param.Field[string] `json:"program_id"`
}

func (r AccountNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountUpdateParams struct {
	// The new name of the Account.
	Name param.Field[string] `json:"name"`
}

func (r AccountUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountListParams struct {
	CreatedAt param.Field[AccountListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Filter Accounts for those belonging to the specified Entity.
	EntityID param.Field[string] `query:"entity_id"`
	// Filter Accounts for those belonging to the specified Entity as informational.
	InformationalEntityID param.Field[string] `query:"informational_entity_id"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
	// Filter Accounts for those with the specified status.
	Status param.Field[AccountListParamsStatus] `query:"status"`
}

// URLQuery serializes [AccountListParams]'s query parameters as `url.Values`.
func (r AccountListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AccountListParamsCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After param.Field[time.Time] `query:"after" format:"date-time"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before param.Field[time.Time] `query:"before" format:"date-time"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter param.Field[time.Time] `query:"on_or_after" format:"date-time"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore param.Field[time.Time] `query:"on_or_before" format:"date-time"`
}

// URLQuery serializes [AccountListParamsCreatedAt]'s query parameters as
// `url.Values`.
func (r AccountListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Filter Accounts for those with the specified status.
type AccountListParamsStatus string

const (
	// Open Accounts that are ready to use.
	AccountListParamsStatusOpen AccountListParamsStatus = "open"
	// Closed Accounts on which no new activity can occur.
	AccountListParamsStatusClosed AccountListParamsStatus = "closed"
)

type AccountBalanceParams struct {
	// The moment to query the balance at. If not set, returns the current balances.
	AtTime param.Field[time.Time] `query:"at_time" format:"date-time"`
}

// URLQuery serializes [AccountBalanceParams]'s query parameters as `url.Values`.
func (r AccountBalanceParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

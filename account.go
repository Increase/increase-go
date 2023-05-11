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
func (r *AccountService) Get(ctx context.Context, account_id string, opts ...option.RequestOption) (res *Account, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s", account_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an Account
func (r *AccountService) Update(ctx context.Context, account_id string, body AccountUpdateParams, opts ...option.RequestOption) (res *Account, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s", account_id)
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

// Close an Account
func (r *AccountService) Close(ctx context.Context, account_id string, opts ...option.RequestOption) (res *Account, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/close", account_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Accounts are your bank accounts with Increase. They store money, receive
// transfers, and send payments. They earn interest and have depository insurance.
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
	JSON accountJSON
}

// accountJSON contains the JSON metadata for the struct [Account]
type accountJSON struct {
	CreatedAt             apijson.Field
	Currency              apijson.Field
	EntityID              apijson.Field
	InformationalEntityID apijson.Field
	ID                    apijson.Field
	InterestAccrued       apijson.Field
	InterestAccruedAt     apijson.Field
	Name                  apijson.Field
	Status                apijson.Field
	Type                  apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

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

type AccountNewParams struct {
	// The identifier for the Entity that will own the Account.
	EntityID param.Field[string] `json:"entity_id"`
	// The identifier for the Program that this Account falls under.
	ProgramID param.Field[string] `json:"program_id"`
	// The identifier of an Entity that, while not owning the Account, is associated
	// with its activity. Its relationship to your group must be `informational`.
	InformationalEntityID param.Field[string] `json:"informational_entity_id"`
	// The name you choose for the Account.
	Name param.Field[string] `json:"name,required"`
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
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
	// Filter Accounts for those belonging to the specified Entity.
	EntityID param.Field[string] `query:"entity_id"`
	// Filter Accounts for those with the specified status.
	Status    param.Field[AccountListParamsStatus]    `query:"status"`
	CreatedAt param.Field[AccountListParamsCreatedAt] `query:"created_at"`
}

// URLQuery serializes [AccountListParams]'s query parameters as `url.Values`.
func (r AccountListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AccountListParamsStatus string

const (
	AccountListParamsStatusOpen   AccountListParamsStatus = "open"
	AccountListParamsStatusClosed AccountListParamsStatus = "closed"
)

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

// A list of Account objects
type AccountListResponse struct {
	// The contents of the list.
	Data []Account `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       accountListResponseJSON
}

// accountListResponseJSON contains the JSON metadata for the struct
// [AccountListResponse]
type accountListResponseJSON struct {
	Data        apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

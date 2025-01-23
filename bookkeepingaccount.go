// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Increase/increase-go/internal/apijson"
	"github.com/Increase/increase-go/internal/apiquery"
	"github.com/Increase/increase-go/internal/param"
	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
	"github.com/Increase/increase-go/packages/pagination"
)

// BookkeepingAccountService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBookkeepingAccountService] method instead.
type BookkeepingAccountService struct {
	Options []option.RequestOption
}

// NewBookkeepingAccountService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBookkeepingAccountService(opts ...option.RequestOption) (r *BookkeepingAccountService) {
	r = &BookkeepingAccountService{}
	r.Options = opts
	return
}

// Create a Bookkeeping Account
func (r *BookkeepingAccountService) New(ctx context.Context, body BookkeepingAccountNewParams, opts ...option.RequestOption) (res *BookkeepingAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := "bookkeeping_accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update a Bookkeeping Account
func (r *BookkeepingAccountService) Update(ctx context.Context, bookkeepingAccountID string, body BookkeepingAccountUpdateParams, opts ...option.RequestOption) (res *BookkeepingAccount, err error) {
	opts = append(r.Options[:], opts...)
	if bookkeepingAccountID == "" {
		err = errors.New("missing required bookkeeping_account_id parameter")
		return
	}
	path := fmt.Sprintf("bookkeeping_accounts/%s", bookkeepingAccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List Bookkeeping Accounts
func (r *BookkeepingAccountService) List(ctx context.Context, query BookkeepingAccountListParams, opts ...option.RequestOption) (res *pagination.Page[BookkeepingAccount], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "bookkeeping_accounts"
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

// List Bookkeeping Accounts
func (r *BookkeepingAccountService) ListAutoPaging(ctx context.Context, query BookkeepingAccountListParams, opts ...option.RequestOption) *pagination.PageAutoPager[BookkeepingAccount] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Retrieve a Bookkeeping Account Balance
func (r *BookkeepingAccountService) Balance(ctx context.Context, bookkeepingAccountID string, query BookkeepingAccountBalanceParams, opts ...option.RequestOption) (res *BookkeepingBalanceLookup, err error) {
	opts = append(r.Options[:], opts...)
	if bookkeepingAccountID == "" {
		err = errors.New("missing required bookkeeping_account_id parameter")
		return
	}
	path := fmt.Sprintf("bookkeeping_accounts/%s/balance", bookkeepingAccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Accounts are T-accounts. They can store accounting entries. Your compliance
// setup might require annotating money movements using this API. Learn more in our
// [guide to Bookkeeping](https://increase.com/documentation/bookkeeping#bookkeeping).
type BookkeepingAccount struct {
	// The account identifier.
	ID string `json:"id,required"`
	// The API Account associated with this bookkeeping account.
	AccountID string `json:"account_id,required,nullable"`
	// The compliance category of the account.
	ComplianceCategory BookkeepingAccountComplianceCategory `json:"compliance_category,required,nullable"`
	// The Entity associated with this bookkeeping account.
	EntityID string `json:"entity_id,required,nullable"`
	// The idempotency key you chose for this object. This value is unique across
	// Increase and is used to ensure that a request is only processed once. Learn more
	// about [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey string `json:"idempotency_key,required,nullable"`
	// The name you choose for the account.
	Name string `json:"name,required"`
	// A constant representing the object's type. For this resource it will always be
	// `bookkeeping_account`.
	Type BookkeepingAccountType `json:"type,required"`
	JSON bookkeepingAccountJSON `json:"-"`
}

// bookkeepingAccountJSON contains the JSON metadata for the struct
// [BookkeepingAccount]
type bookkeepingAccountJSON struct {
	ID                 apijson.Field
	AccountID          apijson.Field
	ComplianceCategory apijson.Field
	EntityID           apijson.Field
	IdempotencyKey     apijson.Field
	Name               apijson.Field
	Type               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *BookkeepingAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bookkeepingAccountJSON) RawJSON() string {
	return r.raw
}

// The compliance category of the account.
type BookkeepingAccountComplianceCategory string

const (
	BookkeepingAccountComplianceCategoryCommingledCash  BookkeepingAccountComplianceCategory = "commingled_cash"
	BookkeepingAccountComplianceCategoryCustomerBalance BookkeepingAccountComplianceCategory = "customer_balance"
)

func (r BookkeepingAccountComplianceCategory) IsKnown() bool {
	switch r {
	case BookkeepingAccountComplianceCategoryCommingledCash, BookkeepingAccountComplianceCategoryCustomerBalance:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `bookkeeping_account`.
type BookkeepingAccountType string

const (
	BookkeepingAccountTypeBookkeepingAccount BookkeepingAccountType = "bookkeeping_account"
)

func (r BookkeepingAccountType) IsKnown() bool {
	switch r {
	case BookkeepingAccountTypeBookkeepingAccount:
		return true
	}
	return false
}

// Represents a request to lookup the balance of an Bookkeeping Account at a given
// point in time.
type BookkeepingBalanceLookup struct {
	// The Bookkeeping Account's current balance, representing the sum of all
	// Bookkeeping Entries on the Bookkeeping Account.
	Balance int64 `json:"balance,required"`
	// The identifier for the account for which the balance was queried.
	BookkeepingAccountID string `json:"bookkeeping_account_id,required"`
	// A constant representing the object's type. For this resource it will always be
	// `bookkeeping_balance_lookup`.
	Type BookkeepingBalanceLookupType `json:"type,required"`
	JSON bookkeepingBalanceLookupJSON `json:"-"`
}

// bookkeepingBalanceLookupJSON contains the JSON metadata for the struct
// [BookkeepingBalanceLookup]
type bookkeepingBalanceLookupJSON struct {
	Balance              apijson.Field
	BookkeepingAccountID apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *BookkeepingBalanceLookup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bookkeepingBalanceLookupJSON) RawJSON() string {
	return r.raw
}

// A constant representing the object's type. For this resource it will always be
// `bookkeeping_balance_lookup`.
type BookkeepingBalanceLookupType string

const (
	BookkeepingBalanceLookupTypeBookkeepingBalanceLookup BookkeepingBalanceLookupType = "bookkeeping_balance_lookup"
)

func (r BookkeepingBalanceLookupType) IsKnown() bool {
	switch r {
	case BookkeepingBalanceLookupTypeBookkeepingBalanceLookup:
		return true
	}
	return false
}

type BookkeepingAccountNewParams struct {
	// The name you choose for the account.
	Name param.Field[string] `json:"name,required"`
	// The entity, if `compliance_category` is `commingled_cash`.
	AccountID param.Field[string] `json:"account_id"`
	// The account compliance category.
	ComplianceCategory param.Field[BookkeepingAccountNewParamsComplianceCategory] `json:"compliance_category"`
	// The entity, if `compliance_category` is `customer_balance`.
	EntityID param.Field[string] `json:"entity_id"`
}

func (r BookkeepingAccountNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The account compliance category.
type BookkeepingAccountNewParamsComplianceCategory string

const (
	BookkeepingAccountNewParamsComplianceCategoryCommingledCash  BookkeepingAccountNewParamsComplianceCategory = "commingled_cash"
	BookkeepingAccountNewParamsComplianceCategoryCustomerBalance BookkeepingAccountNewParamsComplianceCategory = "customer_balance"
)

func (r BookkeepingAccountNewParamsComplianceCategory) IsKnown() bool {
	switch r {
	case BookkeepingAccountNewParamsComplianceCategoryCommingledCash, BookkeepingAccountNewParamsComplianceCategoryCustomerBalance:
		return true
	}
	return false
}

type BookkeepingAccountUpdateParams struct {
	// The name you choose for the account.
	Name param.Field[string] `json:"name,required"`
}

func (r BookkeepingAccountUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BookkeepingAccountListParams struct {
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Filter records to the one with the specified `idempotency_key` you chose for
	// that object. This value is unique across Increase and is used to ensure that a
	// request is only processed once. Learn more about
	// [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey param.Field[string] `query:"idempotency_key"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [BookkeepingAccountListParams]'s query parameters as
// `url.Values`.
func (r BookkeepingAccountListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type BookkeepingAccountBalanceParams struct {
	// The moment to query the balance at. If not set, returns the current balances.
	AtTime param.Field[time.Time] `query:"at_time" format:"date-time"`
}

// URLQuery serializes [BookkeepingAccountBalanceParams]'s query parameters as
// `url.Values`.
func (r BookkeepingAccountBalanceParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

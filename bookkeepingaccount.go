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

// BookkeepingAccountService contains methods and other services that help with
// interacting with the increase API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewBookkeepingAccountService] method
// instead.
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
	path := fmt.Sprintf("bookkeeping_accounts/%s", bookkeepingAccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List Bookkeeping Accounts
func (r *BookkeepingAccountService) List(ctx context.Context, query BookkeepingAccountListParams, opts ...option.RequestOption) (res *shared.Page[BookkeepingAccount], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
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
func (r *BookkeepingAccountService) ListAutoPaging(ctx context.Context, query BookkeepingAccountListParams, opts ...option.RequestOption) *shared.PageAutoPager[BookkeepingAccount] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Retrieve a Bookkeeping Account Balance
func (r *BookkeepingAccountService) Balance(ctx context.Context, bookkeepingAccountID string, query BookkeepingAccountBalanceParams, opts ...option.RequestOption) (res *BookkeepingBalanceLookup, err error) {
	opts = append(r.Options[:], opts...)
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
	// The name you choose for the account.
	Name string `json:"name,required"`
	// A constant representing the object's type. For this resource it will always be
	// `bookkeeping_account`.
	Type BookkeepingAccountType `json:"type,required"`
	JSON bookkeepingAccountJSON
}

// bookkeepingAccountJSON contains the JSON metadata for the struct
// [BookkeepingAccount]
type bookkeepingAccountJSON struct {
	ID                 apijson.Field
	AccountID          apijson.Field
	ComplianceCategory apijson.Field
	EntityID           apijson.Field
	Name               apijson.Field
	Type               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *BookkeepingAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The compliance category of the account.
type BookkeepingAccountComplianceCategory string

const (
	// A cash in an commingled Increase Account.
	BookkeepingAccountComplianceCategoryCommingledCash BookkeepingAccountComplianceCategory = "commingled_cash"
	// A customer balance.
	BookkeepingAccountComplianceCategoryCustomerBalance BookkeepingAccountComplianceCategory = "customer_balance"
)

// A constant representing the object's type. For this resource it will always be
// `bookkeeping_account`.
type BookkeepingAccountType string

const (
	BookkeepingAccountTypeBookkeepingAccount BookkeepingAccountType = "bookkeeping_account"
)

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
	JSON bookkeepingBalanceLookupJSON
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

// A constant representing the object's type. For this resource it will always be
// `bookkeeping_balance_lookup`.
type BookkeepingBalanceLookupType string

const (
	BookkeepingBalanceLookupTypeBookkeepingBalanceLookup BookkeepingBalanceLookupType = "bookkeeping_balance_lookup"
)

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
	// A cash in an commingled Increase Account.
	BookkeepingAccountNewParamsComplianceCategoryCommingledCash BookkeepingAccountNewParamsComplianceCategory = "commingled_cash"
	// A customer balance.
	BookkeepingAccountNewParamsComplianceCategoryCustomerBalance BookkeepingAccountNewParamsComplianceCategory = "customer_balance"
)

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

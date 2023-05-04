package increase

import (
	"context"
	"net/http"
	"net/url"

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

// Accounts are T-accounts. They can store accounting entries.
type BookkeepingAccount struct {
	// The account identifier.
	ID string `json:"id,required"`
	// The compliance category of the account.
	ComplianceCategory BookkeepingAccountComplianceCategory `json:"compliance_category,required,nullable"`
	// The API Account associated with this bookkeeping account.
	AccountID string `json:"account_id,required,nullable"`
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
	ComplianceCategory apijson.Field
	AccountID          apijson.Field
	EntityID           apijson.Field
	Name               apijson.Field
	Type               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *BookkeepingAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BookkeepingAccountComplianceCategory string

const (
	BookkeepingAccountComplianceCategoryCommingledCash  BookkeepingAccountComplianceCategory = "commingled_cash"
	BookkeepingAccountComplianceCategoryCustomerBalance BookkeepingAccountComplianceCategory = "customer_balance"
)

type BookkeepingAccountType string

const (
	BookkeepingAccountTypeBookkeepingAccount BookkeepingAccountType = "bookkeeping_account"
)

type BookkeepingAccountNewParams struct {
	// The account compliance category.
	ComplianceCategory param.Field[BookkeepingAccountNewParamsComplianceCategory] `json:"compliance_category"`
	// The entity, if `compliance_category` is `customer_balance`.
	EntityID param.Field[string] `json:"entity_id"`
	// The entity, if `compliance_category` is `commingled_cash`.
	AccountID param.Field[string] `json:"account_id"`
	// The name you choose for the account.
	Name param.Field[string] `json:"name,required"`
}

func (r BookkeepingAccountNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BookkeepingAccountNewParamsComplianceCategory string

const (
	BookkeepingAccountNewParamsComplianceCategoryCommingledCash  BookkeepingAccountNewParamsComplianceCategory = "commingled_cash"
	BookkeepingAccountNewParamsComplianceCategoryCustomerBalance BookkeepingAccountNewParamsComplianceCategory = "customer_balance"
)

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
	return apiquery.Marshal(r)
}

// A list of Bookkeeping Account objects
type BookkeepingAccountListResponse struct {
	// The contents of the list.
	Data []BookkeepingAccount `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       bookkeepingAccountListResponseJSON
}

// bookkeepingAccountListResponseJSON contains the JSON metadata for the struct
// [BookkeepingAccountListResponse]
type bookkeepingAccountListResponseJSON struct {
	Data        apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BookkeepingAccountListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

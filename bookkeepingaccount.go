package increase

import (
	"context"
	"net/http"
	"net/url"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/apiquery"
	"github.com/increase/increase-go/internal/field"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/internal/shared"
	"github.com/increase/increase-go/option"
)

type BookkeepingAccountService struct {
	Options []option.RequestOption
}

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
	JSON BookkeepingAccountJSON
}

type BookkeepingAccountJSON struct {
	ID                 apijson.Metadata
	ComplianceCategory apijson.Metadata
	AccountID          apijson.Metadata
	EntityID           apijson.Metadata
	Name               apijson.Metadata
	Type               apijson.Metadata
	raw                string
	Extras             map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into BookkeepingAccount using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
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
	ComplianceCategory field.Field[BookkeepingAccountNewParamsComplianceCategory] `json:"compliance_category"`
	// The entity, if `compliance_category` is `customer_balance`.
	EntityID field.Field[string] `json:"entity_id"`
	// The entity, if `compliance_category` is `commingled_cash`.
	AccountID field.Field[string] `json:"account_id"`
	// The name you choose for the account.
	Name field.Field[string] `json:"name,required"`
}

// MarshalJSON serializes BookkeepingAccountNewParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
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
	Cursor field.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit field.Field[int64] `query:"limit"`
}

// URLQuery serializes BookkeepingAccountListParams into a url.Values of the query
// parameters associated with this value
func (r BookkeepingAccountListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

type BookkeepingAccountListResponse struct {
	// The contents of the list.
	Data []BookkeepingAccount `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       BookkeepingAccountListResponseJSON
}

type BookkeepingAccountListResponseJSON struct {
	Data       apijson.Metadata
	NextCursor apijson.Metadata
	raw        string
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// BookkeepingAccountListResponse using the internal json library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *BookkeepingAccountListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

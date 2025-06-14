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

// ExternalAccountService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExternalAccountService] method instead.
type ExternalAccountService struct {
	Options []option.RequestOption
}

// NewExternalAccountService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewExternalAccountService(opts ...option.RequestOption) (r *ExternalAccountService) {
	r = &ExternalAccountService{}
	r.Options = opts
	return
}

// Create an External Account
func (r *ExternalAccountService) New(ctx context.Context, body ExternalAccountNewParams, opts ...option.RequestOption) (res *ExternalAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := "external_accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve an External Account
func (r *ExternalAccountService) Get(ctx context.Context, externalAccountID string, opts ...option.RequestOption) (res *ExternalAccount, err error) {
	opts = append(r.Options[:], opts...)
	if externalAccountID == "" {
		err = errors.New("missing required external_account_id parameter")
		return
	}
	path := fmt.Sprintf("external_accounts/%s", externalAccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an External Account
func (r *ExternalAccountService) Update(ctx context.Context, externalAccountID string, body ExternalAccountUpdateParams, opts ...option.RequestOption) (res *ExternalAccount, err error) {
	opts = append(r.Options[:], opts...)
	if externalAccountID == "" {
		err = errors.New("missing required external_account_id parameter")
		return
	}
	path := fmt.Sprintf("external_accounts/%s", externalAccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List External Accounts
func (r *ExternalAccountService) List(ctx context.Context, query ExternalAccountListParams, opts ...option.RequestOption) (res *pagination.Page[ExternalAccount], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "external_accounts"
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

// List External Accounts
func (r *ExternalAccountService) ListAutoPaging(ctx context.Context, query ExternalAccountListParams, opts ...option.RequestOption) *pagination.PageAutoPager[ExternalAccount] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// External Accounts represent accounts at financial institutions other than
// Increase. You can use this API to store their details for reuse.
type ExternalAccount struct {
	// The External Account's identifier.
	ID string `json:"id,required"`
	// The type of entity that owns the External Account.
	AccountHolder ExternalAccountAccountHolder `json:"account_holder,required"`
	// The destination account number.
	AccountNumber string `json:"account_number,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the External Account was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The External Account's description for display purposes.
	Description string `json:"description,required"`
	// The type of the account to which the transfer will be sent.
	Funding ExternalAccountFunding `json:"funding,required"`
	// The idempotency key you chose for this object. This value is unique across
	// Increase and is used to ensure that a request is only processed once. Learn more
	// about [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey string `json:"idempotency_key,required,nullable"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber string `json:"routing_number,required"`
	// The External Account's status.
	Status ExternalAccountStatus `json:"status,required"`
	// A constant representing the object's type. For this resource it will always be
	// `external_account`.
	Type ExternalAccountType `json:"type,required"`
	JSON externalAccountJSON `json:"-"`
}

// externalAccountJSON contains the JSON metadata for the struct [ExternalAccount]
type externalAccountJSON struct {
	ID             apijson.Field
	AccountHolder  apijson.Field
	AccountNumber  apijson.Field
	CreatedAt      apijson.Field
	Description    apijson.Field
	Funding        apijson.Field
	IdempotencyKey apijson.Field
	RoutingNumber  apijson.Field
	Status         apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ExternalAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r externalAccountJSON) RawJSON() string {
	return r.raw
}

// The type of entity that owns the External Account.
type ExternalAccountAccountHolder string

const (
	ExternalAccountAccountHolderBusiness   ExternalAccountAccountHolder = "business"
	ExternalAccountAccountHolderIndividual ExternalAccountAccountHolder = "individual"
	ExternalAccountAccountHolderUnknown    ExternalAccountAccountHolder = "unknown"
)

func (r ExternalAccountAccountHolder) IsKnown() bool {
	switch r {
	case ExternalAccountAccountHolderBusiness, ExternalAccountAccountHolderIndividual, ExternalAccountAccountHolderUnknown:
		return true
	}
	return false
}

// The type of the account to which the transfer will be sent.
type ExternalAccountFunding string

const (
	ExternalAccountFundingChecking ExternalAccountFunding = "checking"
	ExternalAccountFundingSavings  ExternalAccountFunding = "savings"
	ExternalAccountFundingOther    ExternalAccountFunding = "other"
)

func (r ExternalAccountFunding) IsKnown() bool {
	switch r {
	case ExternalAccountFundingChecking, ExternalAccountFundingSavings, ExternalAccountFundingOther:
		return true
	}
	return false
}

// The External Account's status.
type ExternalAccountStatus string

const (
	ExternalAccountStatusActive   ExternalAccountStatus = "active"
	ExternalAccountStatusArchived ExternalAccountStatus = "archived"
)

func (r ExternalAccountStatus) IsKnown() bool {
	switch r {
	case ExternalAccountStatusActive, ExternalAccountStatusArchived:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `external_account`.
type ExternalAccountType string

const (
	ExternalAccountTypeExternalAccount ExternalAccountType = "external_account"
)

func (r ExternalAccountType) IsKnown() bool {
	switch r {
	case ExternalAccountTypeExternalAccount:
		return true
	}
	return false
}

type ExternalAccountNewParams struct {
	// The account number for the destination account.
	AccountNumber param.Field[string] `json:"account_number,required"`
	// The name you choose for the Account.
	Description param.Field[string] `json:"description,required"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
	// destination account.
	RoutingNumber param.Field[string] `json:"routing_number,required"`
	// The type of entity that owns the External Account.
	AccountHolder param.Field[ExternalAccountNewParamsAccountHolder] `json:"account_holder"`
	// The type of the destination account. Defaults to `checking`.
	Funding param.Field[ExternalAccountNewParamsFunding] `json:"funding"`
}

func (r ExternalAccountNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of entity that owns the External Account.
type ExternalAccountNewParamsAccountHolder string

const (
	ExternalAccountNewParamsAccountHolderBusiness   ExternalAccountNewParamsAccountHolder = "business"
	ExternalAccountNewParamsAccountHolderIndividual ExternalAccountNewParamsAccountHolder = "individual"
	ExternalAccountNewParamsAccountHolderUnknown    ExternalAccountNewParamsAccountHolder = "unknown"
)

func (r ExternalAccountNewParamsAccountHolder) IsKnown() bool {
	switch r {
	case ExternalAccountNewParamsAccountHolderBusiness, ExternalAccountNewParamsAccountHolderIndividual, ExternalAccountNewParamsAccountHolderUnknown:
		return true
	}
	return false
}

// The type of the destination account. Defaults to `checking`.
type ExternalAccountNewParamsFunding string

const (
	ExternalAccountNewParamsFundingChecking ExternalAccountNewParamsFunding = "checking"
	ExternalAccountNewParamsFundingSavings  ExternalAccountNewParamsFunding = "savings"
	ExternalAccountNewParamsFundingOther    ExternalAccountNewParamsFunding = "other"
)

func (r ExternalAccountNewParamsFunding) IsKnown() bool {
	switch r {
	case ExternalAccountNewParamsFundingChecking, ExternalAccountNewParamsFundingSavings, ExternalAccountNewParamsFundingOther:
		return true
	}
	return false
}

type ExternalAccountUpdateParams struct {
	// The type of entity that owns the External Account.
	AccountHolder param.Field[ExternalAccountUpdateParamsAccountHolder] `json:"account_holder"`
	// The description you choose to give the external account.
	Description param.Field[string] `json:"description"`
	// The funding type of the External Account.
	Funding param.Field[ExternalAccountUpdateParamsFunding] `json:"funding"`
	// The status of the External Account.
	Status param.Field[ExternalAccountUpdateParamsStatus] `json:"status"`
}

func (r ExternalAccountUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of entity that owns the External Account.
type ExternalAccountUpdateParamsAccountHolder string

const (
	ExternalAccountUpdateParamsAccountHolderBusiness   ExternalAccountUpdateParamsAccountHolder = "business"
	ExternalAccountUpdateParamsAccountHolderIndividual ExternalAccountUpdateParamsAccountHolder = "individual"
)

func (r ExternalAccountUpdateParamsAccountHolder) IsKnown() bool {
	switch r {
	case ExternalAccountUpdateParamsAccountHolderBusiness, ExternalAccountUpdateParamsAccountHolderIndividual:
		return true
	}
	return false
}

// The funding type of the External Account.
type ExternalAccountUpdateParamsFunding string

const (
	ExternalAccountUpdateParamsFundingChecking ExternalAccountUpdateParamsFunding = "checking"
	ExternalAccountUpdateParamsFundingSavings  ExternalAccountUpdateParamsFunding = "savings"
	ExternalAccountUpdateParamsFundingOther    ExternalAccountUpdateParamsFunding = "other"
)

func (r ExternalAccountUpdateParamsFunding) IsKnown() bool {
	switch r {
	case ExternalAccountUpdateParamsFundingChecking, ExternalAccountUpdateParamsFundingSavings, ExternalAccountUpdateParamsFundingOther:
		return true
	}
	return false
}

// The status of the External Account.
type ExternalAccountUpdateParamsStatus string

const (
	ExternalAccountUpdateParamsStatusActive   ExternalAccountUpdateParamsStatus = "active"
	ExternalAccountUpdateParamsStatusArchived ExternalAccountUpdateParamsStatus = "archived"
)

func (r ExternalAccountUpdateParamsStatus) IsKnown() bool {
	switch r {
	case ExternalAccountUpdateParamsStatusActive, ExternalAccountUpdateParamsStatusArchived:
		return true
	}
	return false
}

type ExternalAccountListParams struct {
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
	// Filter External Accounts to those with the specified Routing Number.
	RoutingNumber param.Field[string]                          `query:"routing_number"`
	Status        param.Field[ExternalAccountListParamsStatus] `query:"status"`
}

// URLQuery serializes [ExternalAccountListParams]'s query parameters as
// `url.Values`.
func (r ExternalAccountListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ExternalAccountListParamsStatus struct {
	// Filter External Accounts for those with the specified status or statuses. For
	// GET requests, this should be encoded as a comma-delimited string, such as
	// `?in=one,two,three`.
	In param.Field[[]ExternalAccountListParamsStatusIn] `query:"in"`
}

// URLQuery serializes [ExternalAccountListParamsStatus]'s query parameters as
// `url.Values`.
func (r ExternalAccountListParamsStatus) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ExternalAccountListParamsStatusIn string

const (
	ExternalAccountListParamsStatusInActive   ExternalAccountListParamsStatusIn = "active"
	ExternalAccountListParamsStatusInArchived ExternalAccountListParamsStatusIn = "archived"
)

func (r ExternalAccountListParamsStatusIn) IsKnown() bool {
	switch r {
	case ExternalAccountListParamsStatusInActive, ExternalAccountListParamsStatusInArchived:
		return true
	}
	return false
}

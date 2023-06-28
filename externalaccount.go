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

// ExternalAccountService contains methods and other services that help with
// interacting with the increase API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewExternalAccountService] method
// instead.
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
	path := fmt.Sprintf("external_accounts/%s", externalAccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an External Account
func (r *ExternalAccountService) Update(ctx context.Context, externalAccountID string, body ExternalAccountUpdateParams, opts ...option.RequestOption) (res *ExternalAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("external_accounts/%s", externalAccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List External Accounts
func (r *ExternalAccountService) List(ctx context.Context, query ExternalAccountListParams, opts ...option.RequestOption) (res *shared.Page[ExternalAccount], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
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
func (r *ExternalAccountService) ListAutoPaging(ctx context.Context, query ExternalAccountListParams, opts ...option.RequestOption) *shared.PageAutoPager[ExternalAccount] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// External Accounts represent accounts at financial institutions other than
// Increase. You can use this API to store their details for reuse.
type ExternalAccount struct {
	// The External Account's identifier.
	ID string `json:"id,required"`
	// The destination account number.
	AccountNumber string `json:"account_number,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the External Account was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The External Account's description for display purposes.
	Description string `json:"description,required"`
	// The type of the account to which the transfer will be sent.
	Funding ExternalAccountFunding `json:"funding,required"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber string `json:"routing_number,required"`
	// The External Account's status.
	Status ExternalAccountStatus `json:"status,required"`
	// A constant representing the object's type. For this resource it will always be
	// `external_account`.
	Type ExternalAccountType `json:"type,required"`
	// If you have verified ownership of the External Account.
	VerificationStatus ExternalAccountVerificationStatus `json:"verification_status,required"`
	JSON               externalAccountJSON
}

// externalAccountJSON contains the JSON metadata for the struct [ExternalAccount]
type externalAccountJSON struct {
	ID                 apijson.Field
	AccountNumber      apijson.Field
	CreatedAt          apijson.Field
	Description        apijson.Field
	Funding            apijson.Field
	RoutingNumber      apijson.Field
	Status             apijson.Field
	Type               apijson.Field
	VerificationStatus apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ExternalAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the account to which the transfer will be sent.
type ExternalAccountFunding string

const (
	ExternalAccountFundingChecking ExternalAccountFunding = "checking"
	ExternalAccountFundingSavings  ExternalAccountFunding = "savings"
	ExternalAccountFundingOther    ExternalAccountFunding = "other"
)

// The External Account's status.
type ExternalAccountStatus string

const (
	ExternalAccountStatusActive   ExternalAccountStatus = "active"
	ExternalAccountStatusArchived ExternalAccountStatus = "archived"
)

// A constant representing the object's type. For this resource it will always be
// `external_account`.
type ExternalAccountType string

const (
	ExternalAccountTypeExternalAccount ExternalAccountType = "external_account"
)

// If you have verified ownership of the External Account.
type ExternalAccountVerificationStatus string

const (
	ExternalAccountVerificationStatusUnverified ExternalAccountVerificationStatus = "unverified"
	ExternalAccountVerificationStatusPending    ExternalAccountVerificationStatus = "pending"
	ExternalAccountVerificationStatusVerified   ExternalAccountVerificationStatus = "verified"
)

type ExternalAccountNewParams struct {
	// The account number for the destination account.
	AccountNumber param.Field[string] `json:"account_number,required"`
	// The name you choose for the Account.
	Description param.Field[string] `json:"description,required"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
	// destination account.
	RoutingNumber param.Field[string] `json:"routing_number,required"`
	// The type of the destination account. Defaults to `checking`.
	Funding param.Field[ExternalAccountNewParamsFunding] `json:"funding"`
}

func (r ExternalAccountNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of the destination account. Defaults to `checking`.
type ExternalAccountNewParamsFunding string

const (
	ExternalAccountNewParamsFundingChecking ExternalAccountNewParamsFunding = "checking"
	ExternalAccountNewParamsFundingSavings  ExternalAccountNewParamsFunding = "savings"
	ExternalAccountNewParamsFundingOther    ExternalAccountNewParamsFunding = "other"
)

type ExternalAccountUpdateParams struct {
	// The description you choose to give the external account.
	Description param.Field[string] `json:"description"`
	// The status of the External Account.
	Status param.Field[ExternalAccountUpdateParamsStatus] `json:"status"`
}

func (r ExternalAccountUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The status of the External Account.
type ExternalAccountUpdateParamsStatus string

const (
	ExternalAccountUpdateParamsStatusActive   ExternalAccountUpdateParamsStatus = "active"
	ExternalAccountUpdateParamsStatusArchived ExternalAccountUpdateParamsStatus = "archived"
)

type ExternalAccountListParams struct {
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit  param.Field[int64]                           `query:"limit"`
	Status param.Field[ExternalAccountListParamsStatus] `query:"status"`
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

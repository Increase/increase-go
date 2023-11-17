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

// AccountNumberService contains methods and other services that help with
// interacting with the increase API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountNumberService] method
// instead.
type AccountNumberService struct {
	Options []option.RequestOption
}

// NewAccountNumberService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountNumberService(opts ...option.RequestOption) (r *AccountNumberService) {
	r = &AccountNumberService{}
	r.Options = opts
	return
}

// Create an Account Number
func (r *AccountNumberService) New(ctx context.Context, body AccountNumberNewParams, opts ...option.RequestOption) (res *AccountNumber, err error) {
	opts = append(r.Options[:], opts...)
	path := "account_numbers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve an Account Number
func (r *AccountNumberService) Get(ctx context.Context, accountNumberID string, opts ...option.RequestOption) (res *AccountNumber, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("account_numbers/%s", accountNumberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an Account Number
func (r *AccountNumberService) Update(ctx context.Context, accountNumberID string, body AccountNumberUpdateParams, opts ...option.RequestOption) (res *AccountNumber, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("account_numbers/%s", accountNumberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List Account Numbers
func (r *AccountNumberService) List(ctx context.Context, query AccountNumberListParams, opts ...option.RequestOption) (res *shared.Page[AccountNumber], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "account_numbers"
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

// List Account Numbers
func (r *AccountNumberService) ListAutoPaging(ctx context.Context, query AccountNumberListParams, opts ...option.RequestOption) *shared.PageAutoPager[AccountNumber] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Each account can have multiple account and routing numbers. We recommend that
// you use a set per vendor. This is similar to how you use different passwords for
// different websites. Account numbers can also be used to seamlessly reconcile
// inbound payments. Generating a unique account number per vendor ensures you
// always know the originator of an incoming payment.
type AccountNumber struct {
	// The Account Number identifier.
	ID string `json:"id,required"`
	// The identifier for the account this Account Number belongs to.
	AccountID string `json:"account_id,required"`
	// The account number.
	AccountNumber string `json:"account_number,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Account
	// Number was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Properties related to how this Account Number handles inbound ACH transfers.
	InboundACH AccountNumberInboundACH `json:"inbound_ach,required"`
	// Properties related to how this Account Number should handle inbound check
	// withdrawls.
	InboundChecks AccountNumberInboundChecks `json:"inbound_checks,required"`
	// The name you choose for the Account Number.
	Name string `json:"name,required"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber string `json:"routing_number,required"`
	// This indicates if payments can be made to the Account Number.
	Status AccountNumberStatus `json:"status,required"`
	// A constant representing the object's type. For this resource it will always be
	// `account_number`.
	Type AccountNumberType `json:"type,required"`
	JSON accountNumberJSON `json:"-"`
}

// accountNumberJSON contains the JSON metadata for the struct [AccountNumber]
type accountNumberJSON struct {
	ID            apijson.Field
	AccountID     apijson.Field
	AccountNumber apijson.Field
	CreatedAt     apijson.Field
	InboundACH    apijson.Field
	InboundChecks apijson.Field
	Name          apijson.Field
	RoutingNumber apijson.Field
	Status        apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountNumber) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Properties related to how this Account Number handles inbound ACH transfers.
type AccountNumberInboundACH struct {
	// Whether ACH debits are allowed against this Account Number. Note that they will
	// still be declined if this is `allowed` if the Account Number is not active.
	DebitStatus AccountNumberInboundACHDebitStatus `json:"debit_status,required"`
	JSON        accountNumberInboundACHJSON        `json:"-"`
}

// accountNumberInboundACHJSON contains the JSON metadata for the struct
// [AccountNumberInboundACH]
type accountNumberInboundACHJSON struct {
	DebitStatus apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountNumberInboundACH) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether ACH debits are allowed against this Account Number. Note that they will
// still be declined if this is `allowed` if the Account Number is not active.
type AccountNumberInboundACHDebitStatus string

const (
	// ACH Debits are allowed.
	AccountNumberInboundACHDebitStatusAllowed AccountNumberInboundACHDebitStatus = "allowed"
	// ACH Debits are blocked.
	AccountNumberInboundACHDebitStatusBlocked AccountNumberInboundACHDebitStatus = "blocked"
)

// Properties related to how this Account Number should handle inbound check
// withdrawls.
type AccountNumberInboundChecks struct {
	// How Increase should process checks with this account number printed on them.
	Status AccountNumberInboundChecksStatus `json:"status,required"`
	JSON   accountNumberInboundChecksJSON   `json:"-"`
}

// accountNumberInboundChecksJSON contains the JSON metadata for the struct
// [AccountNumberInboundChecks]
type accountNumberInboundChecksJSON struct {
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountNumberInboundChecks) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// How Increase should process checks with this account number printed on them.
type AccountNumberInboundChecksStatus string

const (
	// Checks with this Account Number will be processed even if they are not
	// associated with a Check Transfer.
	AccountNumberInboundChecksStatusAllowed AccountNumberInboundChecksStatus = "allowed"
	// Checks with this Account Number will be processed only if they can be matched to
	// an existing Check Transfer.
	AccountNumberInboundChecksStatusCheckTransfersOnly AccountNumberInboundChecksStatus = "check_transfers_only"
)

// This indicates if payments can be made to the Account Number.
type AccountNumberStatus string

const (
	// The account number is active.
	AccountNumberStatusActive AccountNumberStatus = "active"
	// The account number is temporarily disabled.
	AccountNumberStatusDisabled AccountNumberStatus = "disabled"
	// The account number is permanently disabled.
	AccountNumberStatusCanceled AccountNumberStatus = "canceled"
)

// A constant representing the object's type. For this resource it will always be
// `account_number`.
type AccountNumberType string

const (
	AccountNumberTypeAccountNumber AccountNumberType = "account_number"
)

type AccountNumberNewParams struct {
	// The Account the Account Number should belong to.
	AccountID param.Field[string] `json:"account_id,required"`
	// The name you choose for the Account Number.
	Name param.Field[string] `json:"name,required"`
	// Options related to how this Account Number should handle inbound ACH transfers.
	InboundACH param.Field[AccountNumberNewParamsInboundACH] `json:"inbound_ach"`
	// Options related to how this Account Number should handle inbound check
	// withdrawls.
	InboundChecks param.Field[AccountNumberNewParamsInboundChecks] `json:"inbound_checks"`
}

func (r AccountNumberNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Options related to how this Account Number should handle inbound ACH transfers.
type AccountNumberNewParamsInboundACH struct {
	// Whether ACH debits are allowed against this Account Number. Note that ACH debits
	// will be declined if this is `allowed` but the Account Number is not active.
	DebitStatus param.Field[AccountNumberNewParamsInboundACHDebitStatus] `json:"debit_status,required"`
}

func (r AccountNumberNewParamsInboundACH) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether ACH debits are allowed against this Account Number. Note that ACH debits
// will be declined if this is `allowed` but the Account Number is not active.
type AccountNumberNewParamsInboundACHDebitStatus string

const (
	// ACH Debits are allowed.
	AccountNumberNewParamsInboundACHDebitStatusAllowed AccountNumberNewParamsInboundACHDebitStatus = "allowed"
	// ACH Debits are blocked.
	AccountNumberNewParamsInboundACHDebitStatusBlocked AccountNumberNewParamsInboundACHDebitStatus = "blocked"
)

// Options related to how this Account Number should handle inbound check
// withdrawls.
type AccountNumberNewParamsInboundChecks struct {
	// How Increase should process checks with this account number printed on them.
	Status param.Field[AccountNumberNewParamsInboundChecksStatus] `json:"status,required"`
}

func (r AccountNumberNewParamsInboundChecks) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// How Increase should process checks with this account number printed on them.
type AccountNumberNewParamsInboundChecksStatus string

const (
	// Checks with this Account Number will be processed even if they are not
	// associated with a Check Transfer.
	AccountNumberNewParamsInboundChecksStatusAllowed AccountNumberNewParamsInboundChecksStatus = "allowed"
	// Checks with this Account Number will be processed only if they can be matched to
	// an existing Check Transfer.
	AccountNumberNewParamsInboundChecksStatusCheckTransfersOnly AccountNumberNewParamsInboundChecksStatus = "check_transfers_only"
)

type AccountNumberUpdateParams struct {
	// Options related to how this Account Number handles inbound ACH transfers.
	InboundACH param.Field[AccountNumberUpdateParamsInboundACH] `json:"inbound_ach"`
	// The name you choose for the Account Number.
	Name param.Field[string] `json:"name"`
	// This indicates if transfers can be made to the Account Number.
	Status param.Field[AccountNumberUpdateParamsStatus] `json:"status"`
}

func (r AccountNumberUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Options related to how this Account Number handles inbound ACH transfers.
type AccountNumberUpdateParamsInboundACH struct {
	// Whether ACH debits are allowed against this Account Number. Note that ACH debits
	// will be declined if this is `allowed` but the Account Number is not active.
	DebitStatus param.Field[AccountNumberUpdateParamsInboundACHDebitStatus] `json:"debit_status"`
}

func (r AccountNumberUpdateParamsInboundACH) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether ACH debits are allowed against this Account Number. Note that ACH debits
// will be declined if this is `allowed` but the Account Number is not active.
type AccountNumberUpdateParamsInboundACHDebitStatus string

const (
	// ACH Debits are allowed.
	AccountNumberUpdateParamsInboundACHDebitStatusAllowed AccountNumberUpdateParamsInboundACHDebitStatus = "allowed"
	// ACH Debits are blocked.
	AccountNumberUpdateParamsInboundACHDebitStatusBlocked AccountNumberUpdateParamsInboundACHDebitStatus = "blocked"
)

// This indicates if transfers can be made to the Account Number.
type AccountNumberUpdateParamsStatus string

const (
	// The account number is active.
	AccountNumberUpdateParamsStatusActive AccountNumberUpdateParamsStatus = "active"
	// The account number is temporarily disabled.
	AccountNumberUpdateParamsStatusDisabled AccountNumberUpdateParamsStatus = "disabled"
	// The account number is permanently disabled.
	AccountNumberUpdateParamsStatusCanceled AccountNumberUpdateParamsStatus = "canceled"
)

type AccountNumberListParams struct {
	// Filter Account Numbers to those belonging to the specified Account.
	AccountID param.Field[string]                           `query:"account_id"`
	CreatedAt param.Field[AccountNumberListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
	// The status to retrieve Account Numbers for.
	Status param.Field[AccountNumberListParamsStatus] `query:"status"`
}

// URLQuery serializes [AccountNumberListParams]'s query parameters as
// `url.Values`.
func (r AccountNumberListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AccountNumberListParamsCreatedAt struct {
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

// URLQuery serializes [AccountNumberListParamsCreatedAt]'s query parameters as
// `url.Values`.
func (r AccountNumberListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The status to retrieve Account Numbers for.
type AccountNumberListParamsStatus string

const (
	// The account number is active.
	AccountNumberListParamsStatusActive AccountNumberListParamsStatus = "active"
	// The account number is temporarily disabled.
	AccountNumberListParamsStatusDisabled AccountNumberListParamsStatus = "disabled"
	// The account number is permanently disabled.
	AccountNumberListParamsStatusCanceled AccountNumberListParamsStatus = "canceled"
)

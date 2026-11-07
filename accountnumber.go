// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Increase/increase-go/internal/apijson"
	"github.com/Increase/increase-go/internal/apiquery"
	"github.com/Increase/increase-go/internal/param"
	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
	"github.com/Increase/increase-go/packages/pagination"
)

// AccountNumberService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountNumberService] method instead.
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
	opts = slices.Concat(r.Options, opts)
	path := "account_numbers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve an Account Number
func (r *AccountNumberService) Get(ctx context.Context, accountNumberID string, opts ...option.RequestOption) (res *AccountNumber, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountNumberID == "" {
		err = errors.New("missing required account_number_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("account_numbers/%s", accountNumberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update an Account Number
func (r *AccountNumberService) Update(ctx context.Context, accountNumberID string, body AccountNumberUpdateParams, opts ...option.RequestOption) (res *AccountNumber, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountNumberID == "" {
		err = errors.New("missing required account_number_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("account_numbers/%s", accountNumberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List Account Numbers
func (r *AccountNumberService) List(ctx context.Context, query AccountNumberListParams, opts ...option.RequestOption) (res *pagination.Page[AccountNumber], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
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
func (r *AccountNumberService) ListAutoPaging(ctx context.Context, query AccountNumberListParams, opts ...option.RequestOption) *pagination.PageAutoPager[AccountNumber] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Each account can have multiple account and routing numbers. We recommend that
// you use a set per vendor. This is similar to how you use different passwords for
// different websites. Account numbers can also be used to seamlessly reconcile
// inbound payments. Generating a unique account number per vendor ensures you
// always know the originator of an incoming payment.
type AccountNumber struct {
	// The Account Number identifier.
	ID string `json:"id" api:"required"`
	// The identifier for the account this Account Number belongs to.
	AccountID string `json:"account_id" api:"required"`
	// The account number.
	AccountNumber string `json:"account_number" api:"required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Account
	// Number was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The idempotency key you chose for this object. This value is unique across
	// Increase and is used to ensure that a request is only processed once. Learn more
	// about [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey string `json:"idempotency_key" api:"required,nullable"`
	// Whether inbound ACH debits are allowed against this Account Number. Note that
	// they will still be declined if this is `allowed` but the Account Number is not
	// active.
	InboundACHDebits AccountNumberInboundACHDebits `json:"inbound_ach_debits" api:"required"`
	// The name you choose for the Account Number.
	Name string `json:"name" api:"required"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber string `json:"routing_number" api:"required"`
	// This indicates if payments can be made to the Account Number.
	Status AccountNumberStatus `json:"status" api:"required"`
	// A constant representing the object's type. For this resource it will always be
	// `account_number`.
	Type AccountNumberType `json:"type" api:"required"`
	// Whether unregistered inbound checks for this Account Number can be deposited.
	UnregisteredInboundChecks AccountNumberUnregisteredInboundChecks `json:"unregistered_inbound_checks" api:"required"`
	ExtraFields               map[string]interface{}                 `json:"-" api:"extrafields"`
	JSON                      accountNumberJSON                      `json:"-"`
}

// accountNumberJSON contains the JSON metadata for the struct [AccountNumber]
type accountNumberJSON struct {
	ID                        apijson.Field
	AccountID                 apijson.Field
	AccountNumber             apijson.Field
	CreatedAt                 apijson.Field
	IdempotencyKey            apijson.Field
	InboundACHDebits          apijson.Field
	Name                      apijson.Field
	RoutingNumber             apijson.Field
	Status                    apijson.Field
	Type                      apijson.Field
	UnregisteredInboundChecks apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *AccountNumber) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountNumberJSON) RawJSON() string {
	return r.raw
}

// Whether inbound ACH debits are allowed against this Account Number. Note that
// they will still be declined if this is `allowed` but the Account Number is not
// active.
type AccountNumberInboundACHDebits string

const (
	AccountNumberInboundACHDebitsAllowed AccountNumberInboundACHDebits = "allowed"
	AccountNumberInboundACHDebitsBlocked AccountNumberInboundACHDebits = "blocked"
)

func (r AccountNumberInboundACHDebits) IsKnown() bool {
	switch r {
	case AccountNumberInboundACHDebitsAllowed, AccountNumberInboundACHDebitsBlocked:
		return true
	}
	return false
}

// This indicates if payments can be made to the Account Number.
type AccountNumberStatus string

const (
	AccountNumberStatusActive   AccountNumberStatus = "active"
	AccountNumberStatusDisabled AccountNumberStatus = "disabled"
	AccountNumberStatusCanceled AccountNumberStatus = "canceled"
)

func (r AccountNumberStatus) IsKnown() bool {
	switch r {
	case AccountNumberStatusActive, AccountNumberStatusDisabled, AccountNumberStatusCanceled:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `account_number`.
type AccountNumberType string

const (
	AccountNumberTypeAccountNumber AccountNumberType = "account_number"
)

func (r AccountNumberType) IsKnown() bool {
	switch r {
	case AccountNumberTypeAccountNumber:
		return true
	}
	return false
}

// Whether unregistered inbound checks for this Account Number can be deposited.
type AccountNumberUnregisteredInboundChecks string

const (
	AccountNumberUnregisteredInboundChecksAllowed AccountNumberUnregisteredInboundChecks = "allowed"
	AccountNumberUnregisteredInboundChecksBlocked AccountNumberUnregisteredInboundChecks = "blocked"
)

func (r AccountNumberUnregisteredInboundChecks) IsKnown() bool {
	switch r {
	case AccountNumberUnregisteredInboundChecksAllowed, AccountNumberUnregisteredInboundChecksBlocked:
		return true
	}
	return false
}

type AccountNumberNewParams struct {
	// The Account the Account Number should belong to.
	AccountID param.Field[string] `json:"account_id" api:"required"`
	// The name you choose for the Account Number.
	Name param.Field[string] `json:"name" api:"required"`
	// Whether inbound ACH debits are allowed against this Account Number. Note that
	// ACH debits will be declined if this is `allowed` but the Account Number is not
	// active. If you do not specify this field, the default is `allowed`.
	InboundACHDebits param.Field[AccountNumberNewParamsInboundACHDebits] `json:"inbound_ach_debits"`
	// Whether unregistered inbound checks for this Account Number can be deposited. If
	// you do not specify this field, the default is `blocked`.
	UnregisteredInboundChecks param.Field[AccountNumberNewParamsUnregisteredInboundChecks] `json:"unregistered_inbound_checks"`
}

func (r AccountNumberNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether inbound ACH debits are allowed against this Account Number. Note that
// ACH debits will be declined if this is `allowed` but the Account Number is not
// active. If you do not specify this field, the default is `allowed`.
type AccountNumberNewParamsInboundACHDebits string

const (
	AccountNumberNewParamsInboundACHDebitsAllowed AccountNumberNewParamsInboundACHDebits = "allowed"
	AccountNumberNewParamsInboundACHDebitsBlocked AccountNumberNewParamsInboundACHDebits = "blocked"
)

func (r AccountNumberNewParamsInboundACHDebits) IsKnown() bool {
	switch r {
	case AccountNumberNewParamsInboundACHDebitsAllowed, AccountNumberNewParamsInboundACHDebitsBlocked:
		return true
	}
	return false
}

// Whether unregistered inbound checks for this Account Number can be deposited. If
// you do not specify this field, the default is `blocked`.
type AccountNumberNewParamsUnregisteredInboundChecks string

const (
	AccountNumberNewParamsUnregisteredInboundChecksAllowed AccountNumberNewParamsUnregisteredInboundChecks = "allowed"
	AccountNumberNewParamsUnregisteredInboundChecksBlocked AccountNumberNewParamsUnregisteredInboundChecks = "blocked"
)

func (r AccountNumberNewParamsUnregisteredInboundChecks) IsKnown() bool {
	switch r {
	case AccountNumberNewParamsUnregisteredInboundChecksAllowed, AccountNumberNewParamsUnregisteredInboundChecksBlocked:
		return true
	}
	return false
}

type AccountNumberUpdateParams struct {
	// Whether inbound ACH debits are allowed against this Account Number. Note that
	// ACH debits will be declined if this is `allowed` but the Account Number is not
	// active.
	InboundACHDebits param.Field[AccountNumberUpdateParamsInboundACHDebits] `json:"inbound_ach_debits"`
	// The name you choose for the Account Number.
	Name param.Field[string] `json:"name"`
	// This indicates if transfers can be made to the Account Number.
	Status param.Field[AccountNumberUpdateParamsStatus] `json:"status"`
	// Whether unregistered inbound checks for this Account Number can be deposited.
	UnregisteredInboundChecks param.Field[AccountNumberUpdateParamsUnregisteredInboundChecks] `json:"unregistered_inbound_checks"`
}

func (r AccountNumberUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether inbound ACH debits are allowed against this Account Number. Note that
// ACH debits will be declined if this is `allowed` but the Account Number is not
// active.
type AccountNumberUpdateParamsInboundACHDebits string

const (
	AccountNumberUpdateParamsInboundACHDebitsAllowed AccountNumberUpdateParamsInboundACHDebits = "allowed"
	AccountNumberUpdateParamsInboundACHDebitsBlocked AccountNumberUpdateParamsInboundACHDebits = "blocked"
)

func (r AccountNumberUpdateParamsInboundACHDebits) IsKnown() bool {
	switch r {
	case AccountNumberUpdateParamsInboundACHDebitsAllowed, AccountNumberUpdateParamsInboundACHDebitsBlocked:
		return true
	}
	return false
}

// This indicates if transfers can be made to the Account Number.
type AccountNumberUpdateParamsStatus string

const (
	AccountNumberUpdateParamsStatusActive   AccountNumberUpdateParamsStatus = "active"
	AccountNumberUpdateParamsStatusDisabled AccountNumberUpdateParamsStatus = "disabled"
	AccountNumberUpdateParamsStatusCanceled AccountNumberUpdateParamsStatus = "canceled"
)

func (r AccountNumberUpdateParamsStatus) IsKnown() bool {
	switch r {
	case AccountNumberUpdateParamsStatusActive, AccountNumberUpdateParamsStatusDisabled, AccountNumberUpdateParamsStatusCanceled:
		return true
	}
	return false
}

// Whether unregistered inbound checks for this Account Number can be deposited.
type AccountNumberUpdateParamsUnregisteredInboundChecks string

const (
	AccountNumberUpdateParamsUnregisteredInboundChecksAllowed AccountNumberUpdateParamsUnregisteredInboundChecks = "allowed"
	AccountNumberUpdateParamsUnregisteredInboundChecksBlocked AccountNumberUpdateParamsUnregisteredInboundChecks = "blocked"
)

func (r AccountNumberUpdateParamsUnregisteredInboundChecks) IsKnown() bool {
	switch r {
	case AccountNumberUpdateParamsUnregisteredInboundChecksAllowed, AccountNumberUpdateParamsUnregisteredInboundChecksBlocked:
		return true
	}
	return false
}

type AccountNumberListParams struct {
	// Filter Account Numbers to those belonging to the specified Account.
	AccountID param.Field[string]                           `query:"account_id"`
	CreatedAt param.Field[AccountNumberListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Filter records to the one with the specified `idempotency_key` you chose for
	// that object. This value is unique across Increase and is used to ensure that a
	// request is only processed once. Learn more about
	// [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey   param.Field[string]                                  `query:"idempotency_key"`
	InboundACHDebits param.Field[AccountNumberListParamsInboundACHDebits] `query:"inbound_ach_debits"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	//
	// Defaults to `100`.
	Limit  param.Field[int64]                         `query:"limit"`
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

type AccountNumberListParamsInboundACHDebits struct {
	// The inbound ACH debits status to retrieve Account Numbers for. For GET requests,
	// this should be encoded as a comma-delimited string, such as `?in=one,two,three`.
	In param.Field[[]AccountNumberListParamsInboundACHDebitsIn] `query:"in"`
}

// URLQuery serializes [AccountNumberListParamsInboundACHDebits]'s query parameters
// as `url.Values`.
func (r AccountNumberListParamsInboundACHDebits) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AccountNumberListParamsInboundACHDebitsIn string

const (
	AccountNumberListParamsInboundACHDebitsInAllowed AccountNumberListParamsInboundACHDebitsIn = "allowed"
	AccountNumberListParamsInboundACHDebitsInBlocked AccountNumberListParamsInboundACHDebitsIn = "blocked"
)

func (r AccountNumberListParamsInboundACHDebitsIn) IsKnown() bool {
	switch r {
	case AccountNumberListParamsInboundACHDebitsInAllowed, AccountNumberListParamsInboundACHDebitsInBlocked:
		return true
	}
	return false
}

type AccountNumberListParamsStatus struct {
	// The status to retrieve Account Numbers for. For GET requests, this should be
	// encoded as a comma-delimited string, such as `?in=one,two,three`.
	In param.Field[[]AccountNumberListParamsStatusIn] `query:"in"`
}

// URLQuery serializes [AccountNumberListParamsStatus]'s query parameters as
// `url.Values`.
func (r AccountNumberListParamsStatus) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AccountNumberListParamsStatusIn string

const (
	AccountNumberListParamsStatusInActive   AccountNumberListParamsStatusIn = "active"
	AccountNumberListParamsStatusInDisabled AccountNumberListParamsStatusIn = "disabled"
	AccountNumberListParamsStatusInCanceled AccountNumberListParamsStatusIn = "canceled"
)

func (r AccountNumberListParamsStatusIn) IsKnown() bool {
	switch r {
	case AccountNumberListParamsStatusInActive, AccountNumberListParamsStatusInDisabled, AccountNumberListParamsStatusInCanceled:
		return true
	}
	return false
}

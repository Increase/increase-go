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
	return
}

// Retrieve an Account Number
func (r *AccountNumberService) Get(ctx context.Context, accountNumberID string, opts ...option.RequestOption) (res *AccountNumber, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountNumberID == "" {
		err = errors.New("missing required account_number_id parameter")
		return
	}
	path := fmt.Sprintf("account_numbers/%s", accountNumberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an Account Number
func (r *AccountNumberService) Update(ctx context.Context, accountNumberID string, body AccountNumberUpdateParams, opts ...option.RequestOption) (res *AccountNumber, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountNumberID == "" {
		err = errors.New("missing required account_number_id parameter")
		return
	}
	path := fmt.Sprintf("account_numbers/%s", accountNumberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
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
	// Properties related to how this Account Number handles inbound ACH transfers.
	InboundACH AccountNumberInboundACH `json:"inbound_ach" api:"required"`
	// Properties related to how this Account Number should handle inbound check
	// withdrawals.
	InboundChecks AccountNumberInboundChecks `json:"inbound_checks" api:"required"`
	// The name you choose for the Account Number.
	Name string `json:"name" api:"required"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber string `json:"routing_number" api:"required"`
	// This indicates if payments can be made to the Account Number.
	Status AccountNumberStatus `json:"status" api:"required"`
	// A constant representing the object's type. For this resource it will always be
	// `account_number`.
	Type        AccountNumberType      `json:"type" api:"required"`
	ExtraFields map[string]interface{} `json:"-" api:"extrafields"`
	JSON        accountNumberJSON      `json:"-"`
}

// accountNumberJSON contains the JSON metadata for the struct [AccountNumber]
type accountNumberJSON struct {
	ID             apijson.Field
	AccountID      apijson.Field
	AccountNumber  apijson.Field
	CreatedAt      apijson.Field
	IdempotencyKey apijson.Field
	InboundACH     apijson.Field
	InboundChecks  apijson.Field
	Name           apijson.Field
	RoutingNumber  apijson.Field
	Status         apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountNumber) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountNumberJSON) RawJSON() string {
	return r.raw
}

// Properties related to how this Account Number handles inbound ACH transfers.
type AccountNumberInboundACH struct {
	// Whether ACH debits are allowed against this Account Number. Note that they will
	// still be declined if this is `allowed` if the Account Number is not active.
	DebitStatus AccountNumberInboundACHDebitStatus `json:"debit_status" api:"required"`
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

func (r accountNumberInboundACHJSON) RawJSON() string {
	return r.raw
}

// Whether ACH debits are allowed against this Account Number. Note that they will
// still be declined if this is `allowed` if the Account Number is not active.
type AccountNumberInboundACHDebitStatus string

const (
	AccountNumberInboundACHDebitStatusAllowed AccountNumberInboundACHDebitStatus = "allowed"
	AccountNumberInboundACHDebitStatusBlocked AccountNumberInboundACHDebitStatus = "blocked"
)

func (r AccountNumberInboundACHDebitStatus) IsKnown() bool {
	switch r {
	case AccountNumberInboundACHDebitStatusAllowed, AccountNumberInboundACHDebitStatusBlocked:
		return true
	}
	return false
}

// Properties related to how this Account Number should handle inbound check
// withdrawals.
type AccountNumberInboundChecks struct {
	// How Increase should process checks with this account number printed on them.
	Status AccountNumberInboundChecksStatus `json:"status" api:"required"`
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

func (r accountNumberInboundChecksJSON) RawJSON() string {
	return r.raw
}

// How Increase should process checks with this account number printed on them.
type AccountNumberInboundChecksStatus string

const (
	AccountNumberInboundChecksStatusAllowed            AccountNumberInboundChecksStatus = "allowed"
	AccountNumberInboundChecksStatusCheckTransfersOnly AccountNumberInboundChecksStatus = "check_transfers_only"
)

func (r AccountNumberInboundChecksStatus) IsKnown() bool {
	switch r {
	case AccountNumberInboundChecksStatusAllowed, AccountNumberInboundChecksStatusCheckTransfersOnly:
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

type AccountNumberNewParams struct {
	// The Account the Account Number should belong to.
	AccountID param.Field[string] `json:"account_id" api:"required"`
	// The name you choose for the Account Number.
	Name param.Field[string] `json:"name" api:"required"`
	// Options related to how this Account Number should handle inbound ACH transfers.
	InboundACH param.Field[AccountNumberNewParamsInboundACH] `json:"inbound_ach"`
	// Options related to how this Account Number should handle inbound check
	// withdrawals.
	InboundChecks param.Field[AccountNumberNewParamsInboundChecks] `json:"inbound_checks"`
}

func (r AccountNumberNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Options related to how this Account Number should handle inbound ACH transfers.
type AccountNumberNewParamsInboundACH struct {
	// Whether ACH debits are allowed against this Account Number. Note that ACH debits
	// will be declined if this is `allowed` but the Account Number is not active. If
	// you do not specify this field, the default is `allowed`.
	DebitStatus param.Field[AccountNumberNewParamsInboundACHDebitStatus] `json:"debit_status" api:"required"`
}

func (r AccountNumberNewParamsInboundACH) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether ACH debits are allowed against this Account Number. Note that ACH debits
// will be declined if this is `allowed` but the Account Number is not active. If
// you do not specify this field, the default is `allowed`.
type AccountNumberNewParamsInboundACHDebitStatus string

const (
	AccountNumberNewParamsInboundACHDebitStatusAllowed AccountNumberNewParamsInboundACHDebitStatus = "allowed"
	AccountNumberNewParamsInboundACHDebitStatusBlocked AccountNumberNewParamsInboundACHDebitStatus = "blocked"
)

func (r AccountNumberNewParamsInboundACHDebitStatus) IsKnown() bool {
	switch r {
	case AccountNumberNewParamsInboundACHDebitStatusAllowed, AccountNumberNewParamsInboundACHDebitStatusBlocked:
		return true
	}
	return false
}

// Options related to how this Account Number should handle inbound check
// withdrawals.
type AccountNumberNewParamsInboundChecks struct {
	// How Increase should process checks with this account number printed on them. If
	// you do not specify this field, the default is `check_transfers_only`.
	Status param.Field[AccountNumberNewParamsInboundChecksStatus] `json:"status" api:"required"`
}

func (r AccountNumberNewParamsInboundChecks) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// How Increase should process checks with this account number printed on them. If
// you do not specify this field, the default is `check_transfers_only`.
type AccountNumberNewParamsInboundChecksStatus string

const (
	AccountNumberNewParamsInboundChecksStatusAllowed            AccountNumberNewParamsInboundChecksStatus = "allowed"
	AccountNumberNewParamsInboundChecksStatusCheckTransfersOnly AccountNumberNewParamsInboundChecksStatus = "check_transfers_only"
)

func (r AccountNumberNewParamsInboundChecksStatus) IsKnown() bool {
	switch r {
	case AccountNumberNewParamsInboundChecksStatusAllowed, AccountNumberNewParamsInboundChecksStatusCheckTransfersOnly:
		return true
	}
	return false
}

type AccountNumberUpdateParams struct {
	// Options related to how this Account Number handles inbound ACH transfers.
	InboundACH param.Field[AccountNumberUpdateParamsInboundACH] `json:"inbound_ach"`
	// Options related to how this Account Number should handle inbound check
	// withdrawals.
	InboundChecks param.Field[AccountNumberUpdateParamsInboundChecks] `json:"inbound_checks"`
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
	AccountNumberUpdateParamsInboundACHDebitStatusAllowed AccountNumberUpdateParamsInboundACHDebitStatus = "allowed"
	AccountNumberUpdateParamsInboundACHDebitStatusBlocked AccountNumberUpdateParamsInboundACHDebitStatus = "blocked"
)

func (r AccountNumberUpdateParamsInboundACHDebitStatus) IsKnown() bool {
	switch r {
	case AccountNumberUpdateParamsInboundACHDebitStatusAllowed, AccountNumberUpdateParamsInboundACHDebitStatusBlocked:
		return true
	}
	return false
}

// Options related to how this Account Number should handle inbound check
// withdrawals.
type AccountNumberUpdateParamsInboundChecks struct {
	// How Increase should process checks with this account number printed on them.
	Status param.Field[AccountNumberUpdateParamsInboundChecksStatus] `json:"status" api:"required"`
}

func (r AccountNumberUpdateParamsInboundChecks) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// How Increase should process checks with this account number printed on them.
type AccountNumberUpdateParamsInboundChecksStatus string

const (
	AccountNumberUpdateParamsInboundChecksStatusAllowed            AccountNumberUpdateParamsInboundChecksStatus = "allowed"
	AccountNumberUpdateParamsInboundChecksStatusCheckTransfersOnly AccountNumberUpdateParamsInboundChecksStatus = "check_transfers_only"
)

func (r AccountNumberUpdateParamsInboundChecksStatus) IsKnown() bool {
	switch r {
	case AccountNumberUpdateParamsInboundChecksStatusAllowed, AccountNumberUpdateParamsInboundChecksStatusCheckTransfersOnly:
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

type AccountNumberListParams struct {
	// Filter Account Numbers to those belonging to the specified Account.
	AccountID      param.Field[string]                                `query:"account_id"`
	ACHDebitStatus param.Field[AccountNumberListParamsACHDebitStatus] `query:"ach_debit_status"`
	CreatedAt      param.Field[AccountNumberListParamsCreatedAt]      `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Filter records to the one with the specified `idempotency_key` you chose for
	// that object. This value is unique across Increase and is used to ensure that a
	// request is only processed once. Learn more about
	// [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey param.Field[string] `query:"idempotency_key"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
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

type AccountNumberListParamsACHDebitStatus struct {
	// The ACH Debit status to retrieve Account Numbers for. For GET requests, this
	// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
	In param.Field[[]AccountNumberListParamsACHDebitStatusIn] `query:"in"`
}

// URLQuery serializes [AccountNumberListParamsACHDebitStatus]'s query parameters
// as `url.Values`.
func (r AccountNumberListParamsACHDebitStatus) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AccountNumberListParamsACHDebitStatusIn string

const (
	AccountNumberListParamsACHDebitStatusInAllowed AccountNumberListParamsACHDebitStatusIn = "allowed"
	AccountNumberListParamsACHDebitStatusInBlocked AccountNumberListParamsACHDebitStatusIn = "blocked"
)

func (r AccountNumberListParamsACHDebitStatusIn) IsKnown() bool {
	switch r {
	case AccountNumberListParamsACHDebitStatusInAllowed, AccountNumberListParamsACHDebitStatusInBlocked:
		return true
	}
	return false
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

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

// SwiftTransferService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSwiftTransferService] method instead.
type SwiftTransferService struct {
	Options []option.RequestOption
}

// NewSwiftTransferService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSwiftTransferService(opts ...option.RequestOption) (r *SwiftTransferService) {
	r = &SwiftTransferService{}
	r.Options = opts
	return
}

// Create a Swift Transfer
func (r *SwiftTransferService) New(ctx context.Context, body SwiftTransferNewParams, opts ...option.RequestOption) (res *SwiftTransfer, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "swift_transfers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a Swift Transfer
func (r *SwiftTransferService) Get(ctx context.Context, swiftTransferID string, opts ...option.RequestOption) (res *SwiftTransfer, err error) {
	opts = slices.Concat(r.Options, opts)
	if swiftTransferID == "" {
		err = errors.New("missing required swift_transfer_id parameter")
		return
	}
	path := fmt.Sprintf("swift_transfers/%s", swiftTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Swift Transfers
func (r *SwiftTransferService) List(ctx context.Context, query SwiftTransferListParams, opts ...option.RequestOption) (res *pagination.Page[SwiftTransfer], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "swift_transfers"
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

// List Swift Transfers
func (r *SwiftTransferService) ListAutoPaging(ctx context.Context, query SwiftTransferListParams, opts ...option.RequestOption) *pagination.PageAutoPager[SwiftTransfer] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Approve a Swift Transfer
func (r *SwiftTransferService) Approve(ctx context.Context, swiftTransferID string, opts ...option.RequestOption) (res *SwiftTransfer, err error) {
	opts = slices.Concat(r.Options, opts)
	if swiftTransferID == "" {
		err = errors.New("missing required swift_transfer_id parameter")
		return
	}
	path := fmt.Sprintf("swift_transfers/%s/approve", swiftTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Cancel a pending Swift Transfer
func (r *SwiftTransferService) Cancel(ctx context.Context, swiftTransferID string, opts ...option.RequestOption) (res *SwiftTransfer, err error) {
	opts = slices.Concat(r.Options, opts)
	if swiftTransferID == "" {
		err = errors.New("missing required swift_transfer_id parameter")
		return
	}
	path := fmt.Sprintf("swift_transfers/%s/cancel", swiftTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Swift Transfers send funds internationally.
type SwiftTransfer struct {
	// The Swift transfer's identifier.
	ID string `json:"id,required"`
	// The Account to which the transfer belongs.
	AccountID string `json:"account_id,required"`
	// The creditor's account number.
	AccountNumber string `json:"account_number,required"`
	// The transfer amount in USD cents.
	Amount int64 `json:"amount,required"`
	// The bank identification code (BIC) of the creditor.
	BankIdentificationCode string `json:"bank_identification_code,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// What object created the transfer, either via the API or the dashboard.
	CreatedBy SwiftTransferCreatedBy `json:"created_by,required"`
	// The creditor's address.
	CreditorAddress SwiftTransferCreditorAddress `json:"creditor_address,required"`
	// The creditor's name.
	CreditorName string `json:"creditor_name,required"`
	// The debtor's address.
	DebtorAddress SwiftTransferDebtorAddress `json:"debtor_address,required"`
	// The debtor's name.
	DebtorName string `json:"debtor_name,required"`
	// The idempotency key you chose for this object. This value is unique across
	// Increase and is used to ensure that a request is only processed once. Learn more
	// about [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey string `json:"idempotency_key,required,nullable"`
	// The amount that was instructed to be transferred in minor units of the
	// `instructed_currency`.
	InstructedAmount int64 `json:"instructed_amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code of the
	// instructed amount.
	InstructedCurrency SwiftTransferInstructedCurrency `json:"instructed_currency,required"`
	// The ID for the pending transaction representing the transfer.
	PendingTransactionID string `json:"pending_transaction_id,required,nullable"`
	// The creditor's bank account routing or transit number. Required in certain
	// countries.
	RoutingNumber string `json:"routing_number,required,nullable"`
	// The Account Number included in the transfer as the debtor's account number.
	SourceAccountNumberID string `json:"source_account_number_id,required"`
	// The lifecycle status of the transfer.
	Status SwiftTransferStatus `json:"status,required"`
	// The ID for the transaction funding the transfer. This will be populated after
	// the transfer is initiated.
	TransactionID string `json:"transaction_id,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `swift_transfer`.
	Type SwiftTransferType `json:"type,required"`
	// The Unique End-to-end Transaction Reference
	// ([UETR](https://www.swift.com/payments/what-unique-end-end-transaction-reference-uetr))
	// for the transfer.
	UniqueEndToEndTransactionReference string `json:"unique_end_to_end_transaction_reference,required"`
	// The unstructured remittance information that was included with the transfer.
	UnstructuredRemittanceInformation string            `json:"unstructured_remittance_information,required"`
	JSON                              swiftTransferJSON `json:"-"`
}

// swiftTransferJSON contains the JSON metadata for the struct [SwiftTransfer]
type swiftTransferJSON struct {
	ID                                 apijson.Field
	AccountID                          apijson.Field
	AccountNumber                      apijson.Field
	Amount                             apijson.Field
	BankIdentificationCode             apijson.Field
	CreatedAt                          apijson.Field
	CreatedBy                          apijson.Field
	CreditorAddress                    apijson.Field
	CreditorName                       apijson.Field
	DebtorAddress                      apijson.Field
	DebtorName                         apijson.Field
	IdempotencyKey                     apijson.Field
	InstructedAmount                   apijson.Field
	InstructedCurrency                 apijson.Field
	PendingTransactionID               apijson.Field
	RoutingNumber                      apijson.Field
	SourceAccountNumberID              apijson.Field
	Status                             apijson.Field
	TransactionID                      apijson.Field
	Type                               apijson.Field
	UniqueEndToEndTransactionReference apijson.Field
	UnstructuredRemittanceInformation  apijson.Field
	raw                                string
	ExtraFields                        map[string]apijson.Field
}

func (r *SwiftTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r swiftTransferJSON) RawJSON() string {
	return r.raw
}

// What object created the transfer, either via the API or the dashboard.
type SwiftTransferCreatedBy struct {
	// If present, details about the API key that created the transfer.
	APIKey SwiftTransferCreatedByAPIKey `json:"api_key,required,nullable"`
	// The type of object that created this transfer.
	Category SwiftTransferCreatedByCategory `json:"category,required"`
	// If present, details about the OAuth Application that created the transfer.
	OAuthApplication SwiftTransferCreatedByOAuthApplication `json:"oauth_application,required,nullable"`
	// If present, details about the User that created the transfer.
	User SwiftTransferCreatedByUser `json:"user,required,nullable"`
	JSON swiftTransferCreatedByJSON `json:"-"`
}

// swiftTransferCreatedByJSON contains the JSON metadata for the struct
// [SwiftTransferCreatedBy]
type swiftTransferCreatedByJSON struct {
	APIKey           apijson.Field
	Category         apijson.Field
	OAuthApplication apijson.Field
	User             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SwiftTransferCreatedBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r swiftTransferCreatedByJSON) RawJSON() string {
	return r.raw
}

// If present, details about the API key that created the transfer.
type SwiftTransferCreatedByAPIKey struct {
	// The description set for the API key when it was created.
	Description string                           `json:"description,required,nullable"`
	JSON        swiftTransferCreatedByAPIKeyJSON `json:"-"`
}

// swiftTransferCreatedByAPIKeyJSON contains the JSON metadata for the struct
// [SwiftTransferCreatedByAPIKey]
type swiftTransferCreatedByAPIKeyJSON struct {
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SwiftTransferCreatedByAPIKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r swiftTransferCreatedByAPIKeyJSON) RawJSON() string {
	return r.raw
}

// The type of object that created this transfer.
type SwiftTransferCreatedByCategory string

const (
	SwiftTransferCreatedByCategoryAPIKey           SwiftTransferCreatedByCategory = "api_key"
	SwiftTransferCreatedByCategoryOAuthApplication SwiftTransferCreatedByCategory = "oauth_application"
	SwiftTransferCreatedByCategoryUser             SwiftTransferCreatedByCategory = "user"
)

func (r SwiftTransferCreatedByCategory) IsKnown() bool {
	switch r {
	case SwiftTransferCreatedByCategoryAPIKey, SwiftTransferCreatedByCategoryOAuthApplication, SwiftTransferCreatedByCategoryUser:
		return true
	}
	return false
}

// If present, details about the OAuth Application that created the transfer.
type SwiftTransferCreatedByOAuthApplication struct {
	// The name of the OAuth Application.
	Name string                                     `json:"name,required"`
	JSON swiftTransferCreatedByOAuthApplicationJSON `json:"-"`
}

// swiftTransferCreatedByOAuthApplicationJSON contains the JSON metadata for the
// struct [SwiftTransferCreatedByOAuthApplication]
type swiftTransferCreatedByOAuthApplicationJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SwiftTransferCreatedByOAuthApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r swiftTransferCreatedByOAuthApplicationJSON) RawJSON() string {
	return r.raw
}

// If present, details about the User that created the transfer.
type SwiftTransferCreatedByUser struct {
	// The email address of the User.
	Email string                         `json:"email,required"`
	JSON  swiftTransferCreatedByUserJSON `json:"-"`
}

// swiftTransferCreatedByUserJSON contains the JSON metadata for the struct
// [SwiftTransferCreatedByUser]
type swiftTransferCreatedByUserJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SwiftTransferCreatedByUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r swiftTransferCreatedByUserJSON) RawJSON() string {
	return r.raw
}

// The creditor's address.
type SwiftTransferCreditorAddress struct {
	// The city, district, town, or village of the address.
	City string `json:"city,required,nullable"`
	// The two-letter
	// [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code for
	// the country of the address.
	Country string `json:"country,required"`
	// The first line of the address.
	Line1 string `json:"line1,required"`
	// The second line of the address.
	Line2 string `json:"line2,required,nullable"`
	// The ZIP or postal code of the address.
	PostalCode string `json:"postal_code,required,nullable"`
	// The state, province, or region of the address. Required in certain countries.
	State string                           `json:"state,required,nullable"`
	JSON  swiftTransferCreditorAddressJSON `json:"-"`
}

// swiftTransferCreditorAddressJSON contains the JSON metadata for the struct
// [SwiftTransferCreditorAddress]
type swiftTransferCreditorAddressJSON struct {
	City        apijson.Field
	Country     apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SwiftTransferCreditorAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r swiftTransferCreditorAddressJSON) RawJSON() string {
	return r.raw
}

// The debtor's address.
type SwiftTransferDebtorAddress struct {
	// The city, district, town, or village of the address.
	City string `json:"city,required,nullable"`
	// The two-letter
	// [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code for
	// the country of the address.
	Country string `json:"country,required"`
	// The first line of the address.
	Line1 string `json:"line1,required"`
	// The second line of the address.
	Line2 string `json:"line2,required,nullable"`
	// The ZIP or postal code of the address.
	PostalCode string `json:"postal_code,required,nullable"`
	// The state, province, or region of the address. Required in certain countries.
	State string                         `json:"state,required,nullable"`
	JSON  swiftTransferDebtorAddressJSON `json:"-"`
}

// swiftTransferDebtorAddressJSON contains the JSON metadata for the struct
// [SwiftTransferDebtorAddress]
type swiftTransferDebtorAddressJSON struct {
	City        apijson.Field
	Country     apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SwiftTransferDebtorAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r swiftTransferDebtorAddressJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code of the
// instructed amount.
type SwiftTransferInstructedCurrency string

const (
	SwiftTransferInstructedCurrencyUsd SwiftTransferInstructedCurrency = "USD"
)

func (r SwiftTransferInstructedCurrency) IsKnown() bool {
	switch r {
	case SwiftTransferInstructedCurrencyUsd:
		return true
	}
	return false
}

// The lifecycle status of the transfer.
type SwiftTransferStatus string

const (
	SwiftTransferStatusPendingApproval   SwiftTransferStatus = "pending_approval"
	SwiftTransferStatusCanceled          SwiftTransferStatus = "canceled"
	SwiftTransferStatusPendingReviewing  SwiftTransferStatus = "pending_reviewing"
	SwiftTransferStatusRequiresAttention SwiftTransferStatus = "requires_attention"
	SwiftTransferStatusPendingInitiating SwiftTransferStatus = "pending_initiating"
	SwiftTransferStatusInitiated         SwiftTransferStatus = "initiated"
	SwiftTransferStatusRejected          SwiftTransferStatus = "rejected"
	SwiftTransferStatusReturned          SwiftTransferStatus = "returned"
)

func (r SwiftTransferStatus) IsKnown() bool {
	switch r {
	case SwiftTransferStatusPendingApproval, SwiftTransferStatusCanceled, SwiftTransferStatusPendingReviewing, SwiftTransferStatusRequiresAttention, SwiftTransferStatusPendingInitiating, SwiftTransferStatusInitiated, SwiftTransferStatusRejected, SwiftTransferStatusReturned:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `swift_transfer`.
type SwiftTransferType string

const (
	SwiftTransferTypeSwiftTransfer SwiftTransferType = "swift_transfer"
)

func (r SwiftTransferType) IsKnown() bool {
	switch r {
	case SwiftTransferTypeSwiftTransfer:
		return true
	}
	return false
}

type SwiftTransferNewParams struct {
	// The identifier for the account that will send the transfer.
	AccountID param.Field[string] `json:"account_id,required"`
	// The creditor's account number.
	AccountNumber param.Field[string] `json:"account_number,required"`
	// The bank identification code (BIC) of the creditor. If it ends with the
	// three-character branch code, this must be 11 characters long. Otherwise this
	// must be 8 characters and the branch code will be assumed to be `XXX`.
	BankIdentificationCode param.Field[string] `json:"bank_identification_code,required"`
	// The creditor's address.
	CreditorAddress param.Field[SwiftTransferNewParamsCreditorAddress] `json:"creditor_address,required"`
	// The creditor's name.
	CreditorName param.Field[string] `json:"creditor_name,required"`
	// The debtor's address.
	DebtorAddress param.Field[SwiftTransferNewParamsDebtorAddress] `json:"debtor_address,required"`
	// The debtor's name.
	DebtorName param.Field[string] `json:"debtor_name,required"`
	// The amount, in minor units of `instructed_currency`, to send to the creditor.
	InstructedAmount param.Field[int64] `json:"instructed_amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code of the
	// instructed amount.
	InstructedCurrency param.Field[SwiftTransferNewParamsInstructedCurrency] `json:"instructed_currency,required"`
	// The Account Number to include in the transfer as the debtor's account number.
	SourceAccountNumberID param.Field[string] `json:"source_account_number_id,required"`
	// Unstructured remittance information to include in the transfer.
	UnstructuredRemittanceInformation param.Field[string] `json:"unstructured_remittance_information,required"`
	// Whether the transfer requires explicit approval via the dashboard or API.
	RequireApproval param.Field[bool] `json:"require_approval"`
	// The creditor's bank account routing or transit number. Required in certain
	// countries.
	RoutingNumber param.Field[string] `json:"routing_number"`
}

func (r SwiftTransferNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The creditor's address.
type SwiftTransferNewParamsCreditorAddress struct {
	// The city, district, town, or village of the address.
	City param.Field[string] `json:"city,required"`
	// The two-letter
	// [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code for
	// the country of the address.
	Country param.Field[string] `json:"country,required"`
	// The first line of the address. This is usually the street number and street.
	Line1 param.Field[string] `json:"line1,required"`
	// The second line of the address. This might be the floor or room number.
	Line2 param.Field[string] `json:"line2"`
	// The ZIP or postal code of the address. Required in certain countries.
	PostalCode param.Field[string] `json:"postal_code"`
	// The state, province, or region of the address. Required in certain countries.
	State param.Field[string] `json:"state"`
}

func (r SwiftTransferNewParamsCreditorAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The debtor's address.
type SwiftTransferNewParamsDebtorAddress struct {
	// The city, district, town, or village of the address.
	City param.Field[string] `json:"city,required"`
	// The two-letter
	// [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code for
	// the country of the address.
	Country param.Field[string] `json:"country,required"`
	// The first line of the address. This is usually the street number and street.
	Line1 param.Field[string] `json:"line1,required"`
	// The second line of the address. This might be the floor or room number.
	Line2 param.Field[string] `json:"line2"`
	// The ZIP or postal code of the address. Required in certain countries.
	PostalCode param.Field[string] `json:"postal_code"`
	// The state, province, or region of the address. Required in certain countries.
	State param.Field[string] `json:"state"`
}

func (r SwiftTransferNewParamsDebtorAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code of the
// instructed amount.
type SwiftTransferNewParamsInstructedCurrency string

const (
	SwiftTransferNewParamsInstructedCurrencyUsd SwiftTransferNewParamsInstructedCurrency = "USD"
)

func (r SwiftTransferNewParamsInstructedCurrency) IsKnown() bool {
	switch r {
	case SwiftTransferNewParamsInstructedCurrencyUsd:
		return true
	}
	return false
}

type SwiftTransferListParams struct {
	// Filter Swift Transfers to those that originated from the specified Account.
	AccountID param.Field[string]                           `query:"account_id"`
	CreatedAt param.Field[SwiftTransferListParamsCreatedAt] `query:"created_at"`
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
	Status param.Field[SwiftTransferListParamsStatus] `query:"status"`
}

// URLQuery serializes [SwiftTransferListParams]'s query parameters as
// `url.Values`.
func (r SwiftTransferListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type SwiftTransferListParamsCreatedAt struct {
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

// URLQuery serializes [SwiftTransferListParamsCreatedAt]'s query parameters as
// `url.Values`.
func (r SwiftTransferListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type SwiftTransferListParamsStatus struct {
	// Return results whose value is in the provided list. For GET requests, this
	// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
	In param.Field[[]SwiftTransferListParamsStatusIn] `query:"in"`
}

// URLQuery serializes [SwiftTransferListParamsStatus]'s query parameters as
// `url.Values`.
func (r SwiftTransferListParamsStatus) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type SwiftTransferListParamsStatusIn string

const (
	SwiftTransferListParamsStatusInPendingApproval   SwiftTransferListParamsStatusIn = "pending_approval"
	SwiftTransferListParamsStatusInCanceled          SwiftTransferListParamsStatusIn = "canceled"
	SwiftTransferListParamsStatusInPendingReviewing  SwiftTransferListParamsStatusIn = "pending_reviewing"
	SwiftTransferListParamsStatusInRequiresAttention SwiftTransferListParamsStatusIn = "requires_attention"
	SwiftTransferListParamsStatusInPendingInitiating SwiftTransferListParamsStatusIn = "pending_initiating"
	SwiftTransferListParamsStatusInInitiated         SwiftTransferListParamsStatusIn = "initiated"
	SwiftTransferListParamsStatusInRejected          SwiftTransferListParamsStatusIn = "rejected"
	SwiftTransferListParamsStatusInReturned          SwiftTransferListParamsStatusIn = "returned"
)

func (r SwiftTransferListParamsStatusIn) IsKnown() bool {
	switch r {
	case SwiftTransferListParamsStatusInPendingApproval, SwiftTransferListParamsStatusInCanceled, SwiftTransferListParamsStatusInPendingReviewing, SwiftTransferListParamsStatusInRequiresAttention, SwiftTransferListParamsStatusInPendingInitiating, SwiftTransferListParamsStatusInInitiated, SwiftTransferListParamsStatusInRejected, SwiftTransferListParamsStatusInReturned:
		return true
	}
	return false
}

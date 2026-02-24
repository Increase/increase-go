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

// FednowTransferService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFednowTransferService] method instead.
type FednowTransferService struct {
	Options []option.RequestOption
}

// NewFednowTransferService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFednowTransferService(opts ...option.RequestOption) (r *FednowTransferService) {
	r = &FednowTransferService{}
	r.Options = opts
	return
}

// Create a FedNow Transfer
func (r *FednowTransferService) New(ctx context.Context, body FednowTransferNewParams, opts ...option.RequestOption) (res *FednowTransfer, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "fednow_transfers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a FedNow Transfer
func (r *FednowTransferService) Get(ctx context.Context, fednowTransferID string, opts ...option.RequestOption) (res *FednowTransfer, err error) {
	opts = slices.Concat(r.Options, opts)
	if fednowTransferID == "" {
		err = errors.New("missing required fednow_transfer_id parameter")
		return
	}
	path := fmt.Sprintf("fednow_transfers/%s", fednowTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List FedNow Transfers
func (r *FednowTransferService) List(ctx context.Context, query FednowTransferListParams, opts ...option.RequestOption) (res *pagination.Page[FednowTransfer], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "fednow_transfers"
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

// List FedNow Transfers
func (r *FednowTransferService) ListAutoPaging(ctx context.Context, query FednowTransferListParams, opts ...option.RequestOption) *pagination.PageAutoPager[FednowTransfer] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Approve a FedNow Transfer
func (r *FednowTransferService) Approve(ctx context.Context, fednowTransferID string, opts ...option.RequestOption) (res *FednowTransfer, err error) {
	opts = slices.Concat(r.Options, opts)
	if fednowTransferID == "" {
		err = errors.New("missing required fednow_transfer_id parameter")
		return
	}
	path := fmt.Sprintf("fednow_transfers/%s/approve", fednowTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Cancel a pending FedNow Transfer
func (r *FednowTransferService) Cancel(ctx context.Context, fednowTransferID string, opts ...option.RequestOption) (res *FednowTransfer, err error) {
	opts = slices.Concat(r.Options, opts)
	if fednowTransferID == "" {
		err = errors.New("missing required fednow_transfer_id parameter")
		return
	}
	path := fmt.Sprintf("fednow_transfers/%s/cancel", fednowTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// FedNow transfers move funds, within seconds, between your Increase account and
// any other account supporting FedNow.
type FednowTransfer struct {
	// The FedNow Transfer's identifier.
	ID string `json:"id" api:"required"`
	// The Account from which the transfer was sent.
	AccountID string `json:"account_id" api:"required"`
	// The destination account number.
	AccountNumber string `json:"account_number" api:"required"`
	// If the transfer is acknowledged by the recipient bank, this will contain
	// supplemental details.
	Acknowledgement FednowTransferAcknowledgement `json:"acknowledgement" api:"required,nullable"`
	// The transfer amount in USD cents.
	Amount int64 `json:"amount" api:"required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// What object created the transfer, either via the API or the dashboard.
	CreatedBy FednowTransferCreatedBy `json:"created_by" api:"required,nullable"`
	// The name of the transfer's recipient. This is set by the sender when creating
	// the transfer.
	CreditorName string `json:"creditor_name" api:"required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transfer's
	// currency. For FedNow transfers this is always equal to `USD`.
	Currency FednowTransferCurrency `json:"currency" api:"required"`
	// The name of the transfer's sender. If not provided, defaults to the name of the
	// account's entity.
	DebtorName string `json:"debtor_name" api:"required"`
	// The identifier of the External Account the transfer was made to, if any.
	ExternalAccountID string `json:"external_account_id" api:"required,nullable"`
	// The idempotency key you chose for this object. This value is unique across
	// Increase and is used to ensure that a request is only processed once. Learn more
	// about [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey string `json:"idempotency_key" api:"required,nullable"`
	// The ID for the pending transaction representing the transfer.
	PendingTransactionID string `json:"pending_transaction_id" api:"required,nullable"`
	// If the transfer is rejected by FedNow or the destination financial institution,
	// this will contain supplemental details.
	Rejection FednowTransferRejection `json:"rejection" api:"required,nullable"`
	// The destination American Bankers' Association (ABA) Routing Transit Number
	// (RTN).
	RoutingNumber string `json:"routing_number" api:"required"`
	// The Account Number the recipient will see as having sent the transfer.
	SourceAccountNumberID string `json:"source_account_number_id" api:"required"`
	// The lifecycle status of the transfer.
	Status FednowTransferStatus `json:"status" api:"required"`
	// After the transfer is submitted to FedNow, this will contain supplemental
	// details.
	Submission FednowTransferSubmission `json:"submission" api:"required,nullable"`
	// The Transaction funding the transfer once it is complete.
	TransactionID string `json:"transaction_id" api:"required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `fednow_transfer`.
	Type FednowTransferType `json:"type" api:"required"`
	// The Unique End-to-end Transaction Reference
	// ([UETR](https://www.swift.com/payments/what-unique-end-end-transaction-reference-uetr))
	// of the transfer.
	UniqueEndToEndTransactionReference string `json:"unique_end_to_end_transaction_reference" api:"required"`
	// Unstructured information that will show on the recipient's bank statement.
	UnstructuredRemittanceInformation string             `json:"unstructured_remittance_information" api:"required"`
	JSON                              fednowTransferJSON `json:"-"`
}

// fednowTransferJSON contains the JSON metadata for the struct [FednowTransfer]
type fednowTransferJSON struct {
	ID                                 apijson.Field
	AccountID                          apijson.Field
	AccountNumber                      apijson.Field
	Acknowledgement                    apijson.Field
	Amount                             apijson.Field
	CreatedAt                          apijson.Field
	CreatedBy                          apijson.Field
	CreditorName                       apijson.Field
	Currency                           apijson.Field
	DebtorName                         apijson.Field
	ExternalAccountID                  apijson.Field
	IdempotencyKey                     apijson.Field
	PendingTransactionID               apijson.Field
	Rejection                          apijson.Field
	RoutingNumber                      apijson.Field
	SourceAccountNumberID              apijson.Field
	Status                             apijson.Field
	Submission                         apijson.Field
	TransactionID                      apijson.Field
	Type                               apijson.Field
	UniqueEndToEndTransactionReference apijson.Field
	UnstructuredRemittanceInformation  apijson.Field
	raw                                string
	ExtraFields                        map[string]apijson.Field
}

func (r *FednowTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fednowTransferJSON) RawJSON() string {
	return r.raw
}

// If the transfer is acknowledged by the recipient bank, this will contain
// supplemental details.
type FednowTransferAcknowledgement struct {
	// When the transfer was acknowledged.
	AcknowledgedAt time.Time                         `json:"acknowledged_at" api:"required" format:"date-time"`
	JSON           fednowTransferAcknowledgementJSON `json:"-"`
}

// fednowTransferAcknowledgementJSON contains the JSON metadata for the struct
// [FednowTransferAcknowledgement]
type fednowTransferAcknowledgementJSON struct {
	AcknowledgedAt apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *FednowTransferAcknowledgement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fednowTransferAcknowledgementJSON) RawJSON() string {
	return r.raw
}

// What object created the transfer, either via the API or the dashboard.
type FednowTransferCreatedBy struct {
	// The type of object that created this transfer.
	Category FednowTransferCreatedByCategory `json:"category" api:"required"`
	// If present, details about the API key that created the transfer.
	APIKey FednowTransferCreatedByAPIKey `json:"api_key" api:"nullable"`
	// If present, details about the OAuth Application that created the transfer.
	OAuthApplication FednowTransferCreatedByOAuthApplication `json:"oauth_application" api:"nullable"`
	// If present, details about the User that created the transfer.
	User FednowTransferCreatedByUser `json:"user" api:"nullable"`
	JSON fednowTransferCreatedByJSON `json:"-"`
}

// fednowTransferCreatedByJSON contains the JSON metadata for the struct
// [FednowTransferCreatedBy]
type fednowTransferCreatedByJSON struct {
	Category         apijson.Field
	APIKey           apijson.Field
	OAuthApplication apijson.Field
	User             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *FednowTransferCreatedBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fednowTransferCreatedByJSON) RawJSON() string {
	return r.raw
}

// The type of object that created this transfer.
type FednowTransferCreatedByCategory string

const (
	FednowTransferCreatedByCategoryAPIKey           FednowTransferCreatedByCategory = "api_key"
	FednowTransferCreatedByCategoryOAuthApplication FednowTransferCreatedByCategory = "oauth_application"
	FednowTransferCreatedByCategoryUser             FednowTransferCreatedByCategory = "user"
)

func (r FednowTransferCreatedByCategory) IsKnown() bool {
	switch r {
	case FednowTransferCreatedByCategoryAPIKey, FednowTransferCreatedByCategoryOAuthApplication, FednowTransferCreatedByCategoryUser:
		return true
	}
	return false
}

// If present, details about the API key that created the transfer.
type FednowTransferCreatedByAPIKey struct {
	// The description set for the API key when it was created.
	Description string                            `json:"description" api:"required,nullable"`
	JSON        fednowTransferCreatedByAPIKeyJSON `json:"-"`
}

// fednowTransferCreatedByAPIKeyJSON contains the JSON metadata for the struct
// [FednowTransferCreatedByAPIKey]
type fednowTransferCreatedByAPIKeyJSON struct {
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FednowTransferCreatedByAPIKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fednowTransferCreatedByAPIKeyJSON) RawJSON() string {
	return r.raw
}

// If present, details about the OAuth Application that created the transfer.
type FednowTransferCreatedByOAuthApplication struct {
	// The name of the OAuth Application.
	Name string                                      `json:"name" api:"required"`
	JSON fednowTransferCreatedByOAuthApplicationJSON `json:"-"`
}

// fednowTransferCreatedByOAuthApplicationJSON contains the JSON metadata for the
// struct [FednowTransferCreatedByOAuthApplication]
type fednowTransferCreatedByOAuthApplicationJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FednowTransferCreatedByOAuthApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fednowTransferCreatedByOAuthApplicationJSON) RawJSON() string {
	return r.raw
}

// If present, details about the User that created the transfer.
type FednowTransferCreatedByUser struct {
	// The email address of the User.
	Email string                          `json:"email" api:"required"`
	JSON  fednowTransferCreatedByUserJSON `json:"-"`
}

// fednowTransferCreatedByUserJSON contains the JSON metadata for the struct
// [FednowTransferCreatedByUser]
type fednowTransferCreatedByUserJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FednowTransferCreatedByUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fednowTransferCreatedByUserJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transfer's
// currency. For FedNow transfers this is always equal to `USD`.
type FednowTransferCurrency string

const (
	FednowTransferCurrencyUsd FednowTransferCurrency = "USD"
)

func (r FednowTransferCurrency) IsKnown() bool {
	switch r {
	case FednowTransferCurrencyUsd:
		return true
	}
	return false
}

// If the transfer is rejected by FedNow or the destination financial institution,
// this will contain supplemental details.
type FednowTransferRejection struct {
	// Additional information about the rejection provided by the recipient bank.
	RejectReasonAdditionalInformation string `json:"reject_reason_additional_information" api:"required,nullable"`
	// The reason the transfer was rejected as provided by the recipient bank or the
	// FedNow network.
	RejectReasonCode FednowTransferRejectionRejectReasonCode `json:"reject_reason_code" api:"required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was rejected.
	RejectedAt time.Time                   `json:"rejected_at" api:"required,nullable" format:"date-time"`
	JSON       fednowTransferRejectionJSON `json:"-"`
}

// fednowTransferRejectionJSON contains the JSON metadata for the struct
// [FednowTransferRejection]
type fednowTransferRejectionJSON struct {
	RejectReasonAdditionalInformation apijson.Field
	RejectReasonCode                  apijson.Field
	RejectedAt                        apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *FednowTransferRejection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fednowTransferRejectionJSON) RawJSON() string {
	return r.raw
}

// The reason the transfer was rejected as provided by the recipient bank or the
// FedNow network.
type FednowTransferRejectionRejectReasonCode string

const (
	FednowTransferRejectionRejectReasonCodeAccountClosed                                 FednowTransferRejectionRejectReasonCode = "account_closed"
	FednowTransferRejectionRejectReasonCodeAccountBlocked                                FednowTransferRejectionRejectReasonCode = "account_blocked"
	FednowTransferRejectionRejectReasonCodeInvalidCreditorAccountType                    FednowTransferRejectionRejectReasonCode = "invalid_creditor_account_type"
	FednowTransferRejectionRejectReasonCodeInvalidCreditorAccountNumber                  FednowTransferRejectionRejectReasonCode = "invalid_creditor_account_number"
	FednowTransferRejectionRejectReasonCodeInvalidCreditorFinancialInstitutionIdentifier FednowTransferRejectionRejectReasonCode = "invalid_creditor_financial_institution_identifier"
	FednowTransferRejectionRejectReasonCodeEndCustomerDeceased                           FednowTransferRejectionRejectReasonCode = "end_customer_deceased"
	FednowTransferRejectionRejectReasonCodeNarrative                                     FednowTransferRejectionRejectReasonCode = "narrative"
	FednowTransferRejectionRejectReasonCodeTransactionForbidden                          FednowTransferRejectionRejectReasonCode = "transaction_forbidden"
	FednowTransferRejectionRejectReasonCodeTransactionTypeNotSupported                   FednowTransferRejectionRejectReasonCode = "transaction_type_not_supported"
	FednowTransferRejectionRejectReasonCodeAmountExceedsBankLimits                       FednowTransferRejectionRejectReasonCode = "amount_exceeds_bank_limits"
	FednowTransferRejectionRejectReasonCodeInvalidCreditorAddress                        FednowTransferRejectionRejectReasonCode = "invalid_creditor_address"
	FednowTransferRejectionRejectReasonCodeInvalidDebtorAddress                          FednowTransferRejectionRejectReasonCode = "invalid_debtor_address"
	FednowTransferRejectionRejectReasonCodeTimeout                                       FednowTransferRejectionRejectReasonCode = "timeout"
	FednowTransferRejectionRejectReasonCodeProcessingError                               FednowTransferRejectionRejectReasonCode = "processing_error"
	FednowTransferRejectionRejectReasonCodeOther                                         FednowTransferRejectionRejectReasonCode = "other"
)

func (r FednowTransferRejectionRejectReasonCode) IsKnown() bool {
	switch r {
	case FednowTransferRejectionRejectReasonCodeAccountClosed, FednowTransferRejectionRejectReasonCodeAccountBlocked, FednowTransferRejectionRejectReasonCodeInvalidCreditorAccountType, FednowTransferRejectionRejectReasonCodeInvalidCreditorAccountNumber, FednowTransferRejectionRejectReasonCodeInvalidCreditorFinancialInstitutionIdentifier, FednowTransferRejectionRejectReasonCodeEndCustomerDeceased, FednowTransferRejectionRejectReasonCodeNarrative, FednowTransferRejectionRejectReasonCodeTransactionForbidden, FednowTransferRejectionRejectReasonCodeTransactionTypeNotSupported, FednowTransferRejectionRejectReasonCodeAmountExceedsBankLimits, FednowTransferRejectionRejectReasonCodeInvalidCreditorAddress, FednowTransferRejectionRejectReasonCodeInvalidDebtorAddress, FednowTransferRejectionRejectReasonCodeTimeout, FednowTransferRejectionRejectReasonCodeProcessingError, FednowTransferRejectionRejectReasonCodeOther:
		return true
	}
	return false
}

// The lifecycle status of the transfer.
type FednowTransferStatus string

const (
	FednowTransferStatusPendingReviewing  FednowTransferStatus = "pending_reviewing"
	FednowTransferStatusCanceled          FednowTransferStatus = "canceled"
	FednowTransferStatusReviewingRejected FednowTransferStatus = "reviewing_rejected"
	FednowTransferStatusRequiresAttention FednowTransferStatus = "requires_attention"
	FednowTransferStatusPendingApproval   FednowTransferStatus = "pending_approval"
	FednowTransferStatusPendingSubmitting FednowTransferStatus = "pending_submitting"
	FednowTransferStatusPendingResponse   FednowTransferStatus = "pending_response"
	FednowTransferStatusComplete          FednowTransferStatus = "complete"
	FednowTransferStatusRejected          FednowTransferStatus = "rejected"
)

func (r FednowTransferStatus) IsKnown() bool {
	switch r {
	case FednowTransferStatusPendingReviewing, FednowTransferStatusCanceled, FednowTransferStatusReviewingRejected, FednowTransferStatusRequiresAttention, FednowTransferStatusPendingApproval, FednowTransferStatusPendingSubmitting, FednowTransferStatusPendingResponse, FednowTransferStatusComplete, FednowTransferStatusRejected:
		return true
	}
	return false
}

// After the transfer is submitted to FedNow, this will contain supplemental
// details.
type FednowTransferSubmission struct {
	// The FedNow network identification of the message submitted.
	MessageIdentification string `json:"message_identification" api:"required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was submitted to FedNow.
	SubmittedAt time.Time                    `json:"submitted_at" api:"required,nullable" format:"date-time"`
	JSON        fednowTransferSubmissionJSON `json:"-"`
}

// fednowTransferSubmissionJSON contains the JSON metadata for the struct
// [FednowTransferSubmission]
type fednowTransferSubmissionJSON struct {
	MessageIdentification apijson.Field
	SubmittedAt           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *FednowTransferSubmission) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fednowTransferSubmissionJSON) RawJSON() string {
	return r.raw
}

// A constant representing the object's type. For this resource it will always be
// `fednow_transfer`.
type FednowTransferType string

const (
	FednowTransferTypeFednowTransfer FednowTransferType = "fednow_transfer"
)

func (r FednowTransferType) IsKnown() bool {
	switch r {
	case FednowTransferTypeFednowTransfer:
		return true
	}
	return false
}

type FednowTransferNewParams struct {
	// The identifier for the account that will send the transfer.
	AccountID param.Field[string] `json:"account_id" api:"required"`
	// The amount, in minor units, to send to the creditor.
	Amount param.Field[int64] `json:"amount" api:"required"`
	// The creditor's name.
	CreditorName param.Field[string] `json:"creditor_name" api:"required"`
	// The debtor's name.
	DebtorName param.Field[string] `json:"debtor_name" api:"required"`
	// The Account Number to include in the transfer as the debtor's account number.
	SourceAccountNumberID param.Field[string] `json:"source_account_number_id" api:"required"`
	// Unstructured remittance information to include in the transfer.
	UnstructuredRemittanceInformation param.Field[string] `json:"unstructured_remittance_information" api:"required"`
	// The creditor's account number.
	AccountNumber param.Field[string] `json:"account_number"`
	// The creditor's address.
	CreditorAddress param.Field[FednowTransferNewParamsCreditorAddress] `json:"creditor_address"`
	// The debtor's address.
	DebtorAddress param.Field[FednowTransferNewParamsDebtorAddress] `json:"debtor_address"`
	// The ID of an External Account to initiate a transfer to. If this parameter is
	// provided, `account_number` and `routing_number` must be absent.
	ExternalAccountID param.Field[string] `json:"external_account_id"`
	// Whether the transfer requires explicit approval via the dashboard or API.
	RequireApproval param.Field[bool] `json:"require_approval"`
	// The creditor's bank account routing number.
	RoutingNumber param.Field[string] `json:"routing_number"`
}

func (r FednowTransferNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The creditor's address.
type FednowTransferNewParamsCreditorAddress struct {
	// The city, district, town, or village of the address.
	City param.Field[string] `json:"city" api:"required"`
	// The postal code component of the address.
	PostalCode param.Field[string] `json:"postal_code" api:"required"`
	// The US state component of the address.
	State param.Field[string] `json:"state" api:"required"`
	// The first line of the address. This is usually the street number and street.
	Line1 param.Field[string] `json:"line1"`
}

func (r FednowTransferNewParamsCreditorAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The debtor's address.
type FednowTransferNewParamsDebtorAddress struct {
	// The city, district, town, or village of the address.
	City param.Field[string] `json:"city" api:"required"`
	// The postal code component of the address.
	PostalCode param.Field[string] `json:"postal_code" api:"required"`
	// The US state component of the address.
	State param.Field[string] `json:"state" api:"required"`
	// The first line of the address. This is usually the street number and street.
	Line1 param.Field[string] `json:"line1"`
}

func (r FednowTransferNewParamsDebtorAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FednowTransferListParams struct {
	// Filter FedNow Transfers to those that originated from the specified Account.
	AccountID param.Field[string]                            `query:"account_id"`
	CreatedAt param.Field[FednowTransferListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Filter FedNow Transfers to those made to the specified External Account.
	ExternalAccountID param.Field[string] `query:"external_account_id"`
	// Filter records to the one with the specified `idempotency_key` you chose for
	// that object. This value is unique across Increase and is used to ensure that a
	// request is only processed once. Learn more about
	// [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey param.Field[string] `query:"idempotency_key"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit  param.Field[int64]                          `query:"limit"`
	Status param.Field[FednowTransferListParamsStatus] `query:"status"`
}

// URLQuery serializes [FednowTransferListParams]'s query parameters as
// `url.Values`.
func (r FednowTransferListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type FednowTransferListParamsCreatedAt struct {
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

// URLQuery serializes [FednowTransferListParamsCreatedAt]'s query parameters as
// `url.Values`.
func (r FednowTransferListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type FednowTransferListParamsStatus struct {
	// Return results whose value is in the provided list. For GET requests, this
	// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
	In param.Field[[]FednowTransferListParamsStatusIn] `query:"in"`
}

// URLQuery serializes [FednowTransferListParamsStatus]'s query parameters as
// `url.Values`.
func (r FednowTransferListParamsStatus) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type FednowTransferListParamsStatusIn string

const (
	FednowTransferListParamsStatusInPendingReviewing  FednowTransferListParamsStatusIn = "pending_reviewing"
	FednowTransferListParamsStatusInCanceled          FednowTransferListParamsStatusIn = "canceled"
	FednowTransferListParamsStatusInReviewingRejected FednowTransferListParamsStatusIn = "reviewing_rejected"
	FednowTransferListParamsStatusInRequiresAttention FednowTransferListParamsStatusIn = "requires_attention"
	FednowTransferListParamsStatusInPendingApproval   FednowTransferListParamsStatusIn = "pending_approval"
	FednowTransferListParamsStatusInPendingSubmitting FednowTransferListParamsStatusIn = "pending_submitting"
	FednowTransferListParamsStatusInPendingResponse   FednowTransferListParamsStatusIn = "pending_response"
	FednowTransferListParamsStatusInComplete          FednowTransferListParamsStatusIn = "complete"
	FednowTransferListParamsStatusInRejected          FednowTransferListParamsStatusIn = "rejected"
)

func (r FednowTransferListParamsStatusIn) IsKnown() bool {
	switch r {
	case FednowTransferListParamsStatusInPendingReviewing, FednowTransferListParamsStatusInCanceled, FednowTransferListParamsStatusInReviewingRejected, FednowTransferListParamsStatusInRequiresAttention, FednowTransferListParamsStatusInPendingApproval, FednowTransferListParamsStatusInPendingSubmitting, FednowTransferListParamsStatusInPendingResponse, FednowTransferListParamsStatusInComplete, FednowTransferListParamsStatusInRejected:
		return true
	}
	return false
}

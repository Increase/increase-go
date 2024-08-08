// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/apiquery"
	"github.com/increase/increase-go/internal/pagination"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/option"
)

// RealTimePaymentsTransferService contains methods and other services that help
// with interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRealTimePaymentsTransferService] method instead.
type RealTimePaymentsTransferService struct {
	Options []option.RequestOption
}

// NewRealTimePaymentsTransferService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRealTimePaymentsTransferService(opts ...option.RequestOption) (r *RealTimePaymentsTransferService) {
	r = &RealTimePaymentsTransferService{}
	r.Options = opts
	return
}

// Create a Real-Time Payments Transfer
func (r *RealTimePaymentsTransferService) New(ctx context.Context, body RealTimePaymentsTransferNewParams, opts ...option.RequestOption) (res *RealTimePaymentsTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := "real_time_payments_transfers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a Real-Time Payments Transfer
func (r *RealTimePaymentsTransferService) Get(ctx context.Context, realTimePaymentsTransferID string, opts ...option.RequestOption) (res *RealTimePaymentsTransfer, err error) {
	opts = append(r.Options[:], opts...)
	if realTimePaymentsTransferID == "" {
		err = errors.New("missing required real_time_payments_transfer_id parameter")
		return
	}
	path := fmt.Sprintf("real_time_payments_transfers/%s", realTimePaymentsTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Real-Time Payments Transfers
func (r *RealTimePaymentsTransferService) List(ctx context.Context, query RealTimePaymentsTransferListParams, opts ...option.RequestOption) (res *pagination.Page[RealTimePaymentsTransfer], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "real_time_payments_transfers"
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

// List Real-Time Payments Transfers
func (r *RealTimePaymentsTransferService) ListAutoPaging(ctx context.Context, query RealTimePaymentsTransferListParams, opts ...option.RequestOption) *pagination.PageAutoPager[RealTimePaymentsTransfer] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Real-Time Payments transfers move funds, within seconds, between your Increase
// account and any other account on the Real-Time Payments network.
type RealTimePaymentsTransfer struct {
	// The Real-Time Payments Transfer's identifier.
	ID string `json:"id,required"`
	// The Account from which the transfer was sent.
	AccountID string `json:"account_id,required"`
	// The transfer amount in USD cents.
	Amount int64 `json:"amount,required"`
	// If your account requires approvals for transfers and the transfer was approved,
	// this will contain details of the approval.
	Approval RealTimePaymentsTransferApproval `json:"approval,required,nullable"`
	// If your account requires approvals for transfers and the transfer was not
	// approved, this will contain details of the cancellation.
	Cancellation RealTimePaymentsTransferCancellation `json:"cancellation,required,nullable"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// What object created the transfer, either via the API or the dashboard.
	CreatedBy RealTimePaymentsTransferCreatedBy `json:"created_by,required,nullable"`
	// The name of the transfer's recipient. This is set by the sender when creating
	// the transfer.
	CreditorName string `json:"creditor_name,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transfer's
	// currency. For real-time payments transfers this is always equal to `USD`.
	Currency RealTimePaymentsTransferCurrency `json:"currency,required"`
	// The name of the transfer's sender. If not provided, defaults to the name of the
	// account's entity.
	DebtorName string `json:"debtor_name,required,nullable"`
	// The destination account number.
	DestinationAccountNumber string `json:"destination_account_number,required"`
	// The destination American Bankers' Association (ABA) Routing Transit Number
	// (RTN).
	DestinationRoutingNumber string `json:"destination_routing_number,required"`
	// The identifier of the External Account the transfer was made to, if any.
	ExternalAccountID string `json:"external_account_id,required,nullable"`
	// The idempotency key you chose for this object. This value is unique across
	// Increase and is used to ensure that a request is only processed once. Learn more
	// about [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey string `json:"idempotency_key,required,nullable"`
	// The ID for the pending transaction representing the transfer. A pending
	// transaction is created when the transfer
	// [requires approval](https://increase.com/documentation/transfer-approvals#transfer-approvals)
	// by someone else in your organization.
	PendingTransactionID string `json:"pending_transaction_id,required,nullable"`
	// If the transfer is rejected by Real-Time Payments or the destination financial
	// institution, this will contain supplemental details.
	Rejection RealTimePaymentsTransferRejection `json:"rejection,required,nullable"`
	// Unstructured information that will show on the recipient's bank statement.
	RemittanceInformation string `json:"remittance_information,required"`
	// The Account Number the recipient will see as having sent the transfer.
	SourceAccountNumberID string `json:"source_account_number_id,required"`
	// The lifecycle status of the transfer.
	Status RealTimePaymentsTransferStatus `json:"status,required"`
	// After the transfer is submitted to Real-Time Payments, this will contain
	// supplemental details.
	Submission RealTimePaymentsTransferSubmission `json:"submission,required,nullable"`
	// The Transaction funding the transfer once it is complete.
	TransactionID string `json:"transaction_id,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `real_time_payments_transfer`.
	Type RealTimePaymentsTransferType `json:"type,required"`
	// The name of the ultimate recipient of the transfer. Set this if the creditor is
	// an intermediary receiving the payment for someone else.
	UltimateCreditorName string `json:"ultimate_creditor_name,required,nullable"`
	// The name of the ultimate sender of the transfer. Set this if the funds are being
	// sent on behalf of someone who is not the account holder at Increase.
	UltimateDebtorName string                       `json:"ultimate_debtor_name,required,nullable"`
	JSON               realTimePaymentsTransferJSON `json:"-"`
}

// realTimePaymentsTransferJSON contains the JSON metadata for the struct
// [RealTimePaymentsTransfer]
type realTimePaymentsTransferJSON struct {
	ID                       apijson.Field
	AccountID                apijson.Field
	Amount                   apijson.Field
	Approval                 apijson.Field
	Cancellation             apijson.Field
	CreatedAt                apijson.Field
	CreatedBy                apijson.Field
	CreditorName             apijson.Field
	Currency                 apijson.Field
	DebtorName               apijson.Field
	DestinationAccountNumber apijson.Field
	DestinationRoutingNumber apijson.Field
	ExternalAccountID        apijson.Field
	IdempotencyKey           apijson.Field
	PendingTransactionID     apijson.Field
	Rejection                apijson.Field
	RemittanceInformation    apijson.Field
	SourceAccountNumberID    apijson.Field
	Status                   apijson.Field
	Submission               apijson.Field
	TransactionID            apijson.Field
	Type                     apijson.Field
	UltimateCreditorName     apijson.Field
	UltimateDebtorName       apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *RealTimePaymentsTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimePaymentsTransferJSON) RawJSON() string {
	return r.raw
}

// If your account requires approvals for transfers and the transfer was approved,
// this will contain details of the approval.
type RealTimePaymentsTransferApproval struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was approved.
	ApprovedAt time.Time `json:"approved_at,required" format:"date-time"`
	// If the Transfer was approved by a user in the dashboard, the email address of
	// that user.
	ApprovedBy string                               `json:"approved_by,required,nullable"`
	JSON       realTimePaymentsTransferApprovalJSON `json:"-"`
}

// realTimePaymentsTransferApprovalJSON contains the JSON metadata for the struct
// [RealTimePaymentsTransferApproval]
type realTimePaymentsTransferApprovalJSON struct {
	ApprovedAt  apijson.Field
	ApprovedBy  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RealTimePaymentsTransferApproval) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimePaymentsTransferApprovalJSON) RawJSON() string {
	return r.raw
}

// If your account requires approvals for transfers and the transfer was not
// approved, this will contain details of the cancellation.
type RealTimePaymentsTransferCancellation struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Transfer was canceled.
	CanceledAt time.Time `json:"canceled_at,required" format:"date-time"`
	// If the Transfer was canceled by a user in the dashboard, the email address of
	// that user.
	CanceledBy string                                   `json:"canceled_by,required,nullable"`
	JSON       realTimePaymentsTransferCancellationJSON `json:"-"`
}

// realTimePaymentsTransferCancellationJSON contains the JSON metadata for the
// struct [RealTimePaymentsTransferCancellation]
type realTimePaymentsTransferCancellationJSON struct {
	CanceledAt  apijson.Field
	CanceledBy  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RealTimePaymentsTransferCancellation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimePaymentsTransferCancellationJSON) RawJSON() string {
	return r.raw
}

// What object created the transfer, either via the API or the dashboard.
type RealTimePaymentsTransferCreatedBy struct {
	// If present, details about the API key that created the transfer.
	APIKey RealTimePaymentsTransferCreatedByAPIKey `json:"api_key,required,nullable"`
	// The type of object that created this transfer.
	Category RealTimePaymentsTransferCreatedByCategory `json:"category,required"`
	// If present, details about the OAuth Application that created the transfer.
	OAuthApplication RealTimePaymentsTransferCreatedByOAuthApplication `json:"oauth_application,required,nullable"`
	// If present, details about the User that created the transfer.
	User RealTimePaymentsTransferCreatedByUser `json:"user,required,nullable"`
	JSON realTimePaymentsTransferCreatedByJSON `json:"-"`
}

// realTimePaymentsTransferCreatedByJSON contains the JSON metadata for the struct
// [RealTimePaymentsTransferCreatedBy]
type realTimePaymentsTransferCreatedByJSON struct {
	APIKey           apijson.Field
	Category         apijson.Field
	OAuthApplication apijson.Field
	User             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RealTimePaymentsTransferCreatedBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimePaymentsTransferCreatedByJSON) RawJSON() string {
	return r.raw
}

// If present, details about the API key that created the transfer.
type RealTimePaymentsTransferCreatedByAPIKey struct {
	// The description set for the API key when it was created.
	Description string                                      `json:"description,required,nullable"`
	JSON        realTimePaymentsTransferCreatedByAPIKeyJSON `json:"-"`
}

// realTimePaymentsTransferCreatedByAPIKeyJSON contains the JSON metadata for the
// struct [RealTimePaymentsTransferCreatedByAPIKey]
type realTimePaymentsTransferCreatedByAPIKeyJSON struct {
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RealTimePaymentsTransferCreatedByAPIKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimePaymentsTransferCreatedByAPIKeyJSON) RawJSON() string {
	return r.raw
}

// The type of object that created this transfer.
type RealTimePaymentsTransferCreatedByCategory string

const (
	// An API key. Details will be under the `api_key` object.
	RealTimePaymentsTransferCreatedByCategoryAPIKey RealTimePaymentsTransferCreatedByCategory = "api_key"
	// An OAuth application you connected to Increase. Details will be under the
	// `oauth_application` object.
	RealTimePaymentsTransferCreatedByCategoryOAuthApplication RealTimePaymentsTransferCreatedByCategory = "oauth_application"
	// A User in the Increase dashboard. Details will be under the `user` object.
	RealTimePaymentsTransferCreatedByCategoryUser RealTimePaymentsTransferCreatedByCategory = "user"
)

func (r RealTimePaymentsTransferCreatedByCategory) IsKnown() bool {
	switch r {
	case RealTimePaymentsTransferCreatedByCategoryAPIKey, RealTimePaymentsTransferCreatedByCategoryOAuthApplication, RealTimePaymentsTransferCreatedByCategoryUser:
		return true
	}
	return false
}

// If present, details about the OAuth Application that created the transfer.
type RealTimePaymentsTransferCreatedByOAuthApplication struct {
	// The name of the OAuth Application.
	Name string                                                `json:"name,required"`
	JSON realTimePaymentsTransferCreatedByOAuthApplicationJSON `json:"-"`
}

// realTimePaymentsTransferCreatedByOAuthApplicationJSON contains the JSON metadata
// for the struct [RealTimePaymentsTransferCreatedByOAuthApplication]
type realTimePaymentsTransferCreatedByOAuthApplicationJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RealTimePaymentsTransferCreatedByOAuthApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimePaymentsTransferCreatedByOAuthApplicationJSON) RawJSON() string {
	return r.raw
}

// If present, details about the User that created the transfer.
type RealTimePaymentsTransferCreatedByUser struct {
	// The email address of the User.
	Email string                                    `json:"email,required"`
	JSON  realTimePaymentsTransferCreatedByUserJSON `json:"-"`
}

// realTimePaymentsTransferCreatedByUserJSON contains the JSON metadata for the
// struct [RealTimePaymentsTransferCreatedByUser]
type realTimePaymentsTransferCreatedByUserJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RealTimePaymentsTransferCreatedByUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimePaymentsTransferCreatedByUserJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transfer's
// currency. For real-time payments transfers this is always equal to `USD`.
type RealTimePaymentsTransferCurrency string

const (
	// Canadian Dollar (CAD)
	RealTimePaymentsTransferCurrencyCad RealTimePaymentsTransferCurrency = "CAD"
	// Swiss Franc (CHF)
	RealTimePaymentsTransferCurrencyChf RealTimePaymentsTransferCurrency = "CHF"
	// Euro (EUR)
	RealTimePaymentsTransferCurrencyEur RealTimePaymentsTransferCurrency = "EUR"
	// British Pound (GBP)
	RealTimePaymentsTransferCurrencyGbp RealTimePaymentsTransferCurrency = "GBP"
	// Japanese Yen (JPY)
	RealTimePaymentsTransferCurrencyJpy RealTimePaymentsTransferCurrency = "JPY"
	// US Dollar (USD)
	RealTimePaymentsTransferCurrencyUsd RealTimePaymentsTransferCurrency = "USD"
)

func (r RealTimePaymentsTransferCurrency) IsKnown() bool {
	switch r {
	case RealTimePaymentsTransferCurrencyCad, RealTimePaymentsTransferCurrencyChf, RealTimePaymentsTransferCurrencyEur, RealTimePaymentsTransferCurrencyGbp, RealTimePaymentsTransferCurrencyJpy, RealTimePaymentsTransferCurrencyUsd:
		return true
	}
	return false
}

// If the transfer is rejected by Real-Time Payments or the destination financial
// institution, this will contain supplemental details.
type RealTimePaymentsTransferRejection struct {
	// Additional information about the rejection provided by the recipient bank when
	// the `reject_reason_code` is `NARRATIVE`.
	RejectReasonAdditionalInformation string `json:"reject_reason_additional_information,required,nullable"`
	// The reason the transfer was rejected as provided by the recipient bank or the
	// Real-Time Payments network.
	RejectReasonCode RealTimePaymentsTransferRejectionRejectReasonCode `json:"reject_reason_code,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was rejected.
	RejectedAt time.Time                             `json:"rejected_at,required,nullable" format:"date-time"`
	JSON       realTimePaymentsTransferRejectionJSON `json:"-"`
}

// realTimePaymentsTransferRejectionJSON contains the JSON metadata for the struct
// [RealTimePaymentsTransferRejection]
type realTimePaymentsTransferRejectionJSON struct {
	RejectReasonAdditionalInformation apijson.Field
	RejectReasonCode                  apijson.Field
	RejectedAt                        apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *RealTimePaymentsTransferRejection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimePaymentsTransferRejectionJSON) RawJSON() string {
	return r.raw
}

// The reason the transfer was rejected as provided by the recipient bank or the
// Real-Time Payments network.
type RealTimePaymentsTransferRejectionRejectReasonCode string

const (
	// The destination account is closed. Corresponds to the Real-Time Payments reason
	// code `AC04`.
	RealTimePaymentsTransferRejectionRejectReasonCodeAccountClosed RealTimePaymentsTransferRejectionRejectReasonCode = "account_closed"
	// The destination account is currently blocked from receiving transactions.
	// Corresponds to the Real-Time Payments reason code `AC06`.
	RealTimePaymentsTransferRejectionRejectReasonCodeAccountBlocked RealTimePaymentsTransferRejectionRejectReasonCode = "account_blocked"
	// The destination account is ineligible to receive Real-Time Payments transfers.
	// Corresponds to the Real-Time Payments reason code `AC14`.
	RealTimePaymentsTransferRejectionRejectReasonCodeInvalidCreditorAccountType RealTimePaymentsTransferRejectionRejectReasonCode = "invalid_creditor_account_type"
	// The destination account does not exist. Corresponds to the Real-Time Payments
	// reason code `AC03`.
	RealTimePaymentsTransferRejectionRejectReasonCodeInvalidCreditorAccountNumber RealTimePaymentsTransferRejectionRejectReasonCode = "invalid_creditor_account_number"
	// The destination routing number is invalid. Corresponds to the Real-Time Payments
	// reason code `RC04`.
	RealTimePaymentsTransferRejectionRejectReasonCodeInvalidCreditorFinancialInstitutionIdentifier RealTimePaymentsTransferRejectionRejectReasonCode = "invalid_creditor_financial_institution_identifier"
	// The destination account holder is deceased. Corresponds to the Real-Time
	// Payments reason code `MD07`.
	RealTimePaymentsTransferRejectionRejectReasonCodeEndCustomerDeceased RealTimePaymentsTransferRejectionRejectReasonCode = "end_customer_deceased"
	// The reason is provided as narrative information in the additional information
	// field.
	RealTimePaymentsTransferRejectionRejectReasonCodeNarrative RealTimePaymentsTransferRejectionRejectReasonCode = "narrative"
	// Real-Time Payments transfers are not allowed to the destination account.
	// Corresponds to the Real-Time Payments reason code `AG01`.
	RealTimePaymentsTransferRejectionRejectReasonCodeTransactionForbidden RealTimePaymentsTransferRejectionRejectReasonCode = "transaction_forbidden"
	// Real-Time Payments transfers are not enabled for the destination account.
	// Corresponds to the Real-Time Payments reason code `AG03`.
	RealTimePaymentsTransferRejectionRejectReasonCodeTransactionTypeNotSupported RealTimePaymentsTransferRejectionRejectReasonCode = "transaction_type_not_supported"
	// The amount of the transfer is different than expected by the recipient.
	// Corresponds to the Real-Time Payments reason code `AM09`.
	RealTimePaymentsTransferRejectionRejectReasonCodeUnexpectedAmount RealTimePaymentsTransferRejectionRejectReasonCode = "unexpected_amount"
	// The amount is higher than the recipient is authorized to send or receive.
	// Corresponds to the Real-Time Payments reason code `AM14`.
	RealTimePaymentsTransferRejectionRejectReasonCodeAmountExceedsBankLimits RealTimePaymentsTransferRejectionRejectReasonCode = "amount_exceeds_bank_limits"
	// The creditor's address is required, but missing or invalid. Corresponds to the
	// Real-Time Payments reason code `BE04`.
	RealTimePaymentsTransferRejectionRejectReasonCodeInvalidCreditorAddress RealTimePaymentsTransferRejectionRejectReasonCode = "invalid_creditor_address"
	// The specified creditor is unknown. Corresponds to the Real-Time Payments reason
	// code `BE06`.
	RealTimePaymentsTransferRejectionRejectReasonCodeUnknownEndCustomer RealTimePaymentsTransferRejectionRejectReasonCode = "unknown_end_customer"
	// The debtor's address is required, but missing or invalid. Corresponds to the
	// Real-Time Payments reason code `BE07`.
	RealTimePaymentsTransferRejectionRejectReasonCodeInvalidDebtorAddress RealTimePaymentsTransferRejectionRejectReasonCode = "invalid_debtor_address"
	// There was a timeout processing the transfer. Corresponds to the Real-Time
	// Payments reason code `DS24`.
	RealTimePaymentsTransferRejectionRejectReasonCodeTimeout RealTimePaymentsTransferRejectionRejectReasonCode = "timeout"
	// Real-Time Payments transfers are not enabled for the destination account.
	// Corresponds to the Real-Time Payments reason code `NOAT`.
	RealTimePaymentsTransferRejectionRejectReasonCodeUnsupportedMessageForRecipient RealTimePaymentsTransferRejectionRejectReasonCode = "unsupported_message_for_recipient"
	// The destination financial institution is currently not connected to Real-Time
	// Payments. Corresponds to the Real-Time Payments reason code `9912`.
	RealTimePaymentsTransferRejectionRejectReasonCodeRecipientConnectionNotAvailable RealTimePaymentsTransferRejectionRejectReasonCode = "recipient_connection_not_available"
	// Real-Time Payments is currently unavailable. Corresponds to the Real-Time
	// Payments reason code `9948`.
	RealTimePaymentsTransferRejectionRejectReasonCodeRealTimePaymentsSuspended RealTimePaymentsTransferRejectionRejectReasonCode = "real_time_payments_suspended"
	// The destination financial institution is currently signed off of Real-Time
	// Payments. Corresponds to the Real-Time Payments reason code `9910`.
	RealTimePaymentsTransferRejectionRejectReasonCodeInstructedAgentSignedOff RealTimePaymentsTransferRejectionRejectReasonCode = "instructed_agent_signed_off"
	// The transfer was rejected due to an internal Increase issue. We have been
	// notified.
	RealTimePaymentsTransferRejectionRejectReasonCodeProcessingError RealTimePaymentsTransferRejectionRejectReasonCode = "processing_error"
	// Some other error or issue has occurred.
	RealTimePaymentsTransferRejectionRejectReasonCodeOther RealTimePaymentsTransferRejectionRejectReasonCode = "other"
)

func (r RealTimePaymentsTransferRejectionRejectReasonCode) IsKnown() bool {
	switch r {
	case RealTimePaymentsTransferRejectionRejectReasonCodeAccountClosed, RealTimePaymentsTransferRejectionRejectReasonCodeAccountBlocked, RealTimePaymentsTransferRejectionRejectReasonCodeInvalidCreditorAccountType, RealTimePaymentsTransferRejectionRejectReasonCodeInvalidCreditorAccountNumber, RealTimePaymentsTransferRejectionRejectReasonCodeInvalidCreditorFinancialInstitutionIdentifier, RealTimePaymentsTransferRejectionRejectReasonCodeEndCustomerDeceased, RealTimePaymentsTransferRejectionRejectReasonCodeNarrative, RealTimePaymentsTransferRejectionRejectReasonCodeTransactionForbidden, RealTimePaymentsTransferRejectionRejectReasonCodeTransactionTypeNotSupported, RealTimePaymentsTransferRejectionRejectReasonCodeUnexpectedAmount, RealTimePaymentsTransferRejectionRejectReasonCodeAmountExceedsBankLimits, RealTimePaymentsTransferRejectionRejectReasonCodeInvalidCreditorAddress, RealTimePaymentsTransferRejectionRejectReasonCodeUnknownEndCustomer, RealTimePaymentsTransferRejectionRejectReasonCodeInvalidDebtorAddress, RealTimePaymentsTransferRejectionRejectReasonCodeTimeout, RealTimePaymentsTransferRejectionRejectReasonCodeUnsupportedMessageForRecipient, RealTimePaymentsTransferRejectionRejectReasonCodeRecipientConnectionNotAvailable, RealTimePaymentsTransferRejectionRejectReasonCodeRealTimePaymentsSuspended, RealTimePaymentsTransferRejectionRejectReasonCodeInstructedAgentSignedOff, RealTimePaymentsTransferRejectionRejectReasonCodeProcessingError, RealTimePaymentsTransferRejectionRejectReasonCodeOther:
		return true
	}
	return false
}

// The lifecycle status of the transfer.
type RealTimePaymentsTransferStatus string

const (
	// The transfer is pending approval.
	RealTimePaymentsTransferStatusPendingApproval RealTimePaymentsTransferStatus = "pending_approval"
	// The transfer has been canceled.
	RealTimePaymentsTransferStatusCanceled RealTimePaymentsTransferStatus = "canceled"
	// The transfer is pending review by Increase.
	RealTimePaymentsTransferStatusPendingReviewing RealTimePaymentsTransferStatus = "pending_reviewing"
	// The transfer is queued to be submitted to Real-Time Payments.
	RealTimePaymentsTransferStatusPendingSubmission RealTimePaymentsTransferStatus = "pending_submission"
	// The transfer has been submitted and is pending a response from Real-Time
	// Payments.
	RealTimePaymentsTransferStatusSubmitted RealTimePaymentsTransferStatus = "submitted"
	// The transfer has been sent successfully and is complete.
	RealTimePaymentsTransferStatusComplete RealTimePaymentsTransferStatus = "complete"
	// The transfer was rejected by the network or the recipient's bank.
	RealTimePaymentsTransferStatusRejected RealTimePaymentsTransferStatus = "rejected"
	// The transfer requires attention from an Increase operator.
	RealTimePaymentsTransferStatusRequiresAttention RealTimePaymentsTransferStatus = "requires_attention"
)

func (r RealTimePaymentsTransferStatus) IsKnown() bool {
	switch r {
	case RealTimePaymentsTransferStatusPendingApproval, RealTimePaymentsTransferStatusCanceled, RealTimePaymentsTransferStatusPendingReviewing, RealTimePaymentsTransferStatusPendingSubmission, RealTimePaymentsTransferStatusSubmitted, RealTimePaymentsTransferStatusComplete, RealTimePaymentsTransferStatusRejected, RealTimePaymentsTransferStatusRequiresAttention:
		return true
	}
	return false
}

// After the transfer is submitted to Real-Time Payments, this will contain
// supplemental details.
type RealTimePaymentsTransferSubmission struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was submitted to The Clearing House.
	SubmittedAt time.Time `json:"submitted_at,required,nullable" format:"date-time"`
	// The Real-Time Payments network identification of the transfer.
	TransactionIdentification string                                 `json:"transaction_identification,required"`
	JSON                      realTimePaymentsTransferSubmissionJSON `json:"-"`
}

// realTimePaymentsTransferSubmissionJSON contains the JSON metadata for the struct
// [RealTimePaymentsTransferSubmission]
type realTimePaymentsTransferSubmissionJSON struct {
	SubmittedAt               apijson.Field
	TransactionIdentification apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *RealTimePaymentsTransferSubmission) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimePaymentsTransferSubmissionJSON) RawJSON() string {
	return r.raw
}

// A constant representing the object's type. For this resource it will always be
// `real_time_payments_transfer`.
type RealTimePaymentsTransferType string

const (
	RealTimePaymentsTransferTypeRealTimePaymentsTransfer RealTimePaymentsTransferType = "real_time_payments_transfer"
)

func (r RealTimePaymentsTransferType) IsKnown() bool {
	switch r {
	case RealTimePaymentsTransferTypeRealTimePaymentsTransfer:
		return true
	}
	return false
}

type RealTimePaymentsTransferNewParams struct {
	// The transfer amount in USD cents. For Real-Time Payments transfers, must be
	// positive.
	Amount param.Field[int64] `json:"amount,required"`
	// The name of the transfer's recipient.
	CreditorName param.Field[string] `json:"creditor_name,required"`
	// Unstructured information that will show on the recipient's bank statement.
	RemittanceInformation param.Field[string] `json:"remittance_information,required"`
	// The identifier of the Account Number from which to send the transfer.
	SourceAccountNumberID param.Field[string] `json:"source_account_number_id,required"`
	// The name of the transfer's sender. If not provided, defaults to the name of the
	// account's entity.
	DebtorName param.Field[string] `json:"debtor_name"`
	// The destination account number.
	DestinationAccountNumber param.Field[string] `json:"destination_account_number"`
	// The destination American Bankers' Association (ABA) Routing Transit Number
	// (RTN).
	DestinationRoutingNumber param.Field[string] `json:"destination_routing_number"`
	// The ID of an External Account to initiate a transfer to. If this parameter is
	// provided, `destination_account_number` and `destination_routing_number` must be
	// absent.
	ExternalAccountID param.Field[string] `json:"external_account_id"`
	// Whether the transfer requires explicit approval via the dashboard or API.
	RequireApproval param.Field[bool] `json:"require_approval"`
	// The name of the ultimate recipient of the transfer. Set this if the creditor is
	// an intermediary receiving the payment for someone else.
	UltimateCreditorName param.Field[string] `json:"ultimate_creditor_name"`
	// The name of the ultimate sender of the transfer. Set this if the funds are being
	// sent on behalf of someone who is not the account holder at Increase.
	UltimateDebtorName param.Field[string] `json:"ultimate_debtor_name"`
}

func (r RealTimePaymentsTransferNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RealTimePaymentsTransferListParams struct {
	// Filter Real-Time Payments Transfers to those belonging to the specified Account.
	AccountID param.Field[string]                                      `query:"account_id"`
	CreatedAt param.Field[RealTimePaymentsTransferListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Filter Real-Time Payments Transfers to those made to the specified External
	// Account.
	ExternalAccountID param.Field[string] `query:"external_account_id"`
	// Filter records to the one with the specified `idempotency_key` you chose for
	// that object. This value is unique across Increase and is used to ensure that a
	// request is only processed once. Learn more about
	// [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey param.Field[string] `query:"idempotency_key"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [RealTimePaymentsTransferListParams]'s query parameters as
// `url.Values`.
func (r RealTimePaymentsTransferListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type RealTimePaymentsTransferListParamsCreatedAt struct {
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

// URLQuery serializes [RealTimePaymentsTransferListParamsCreatedAt]'s query
// parameters as `url.Values`.
func (r RealTimePaymentsTransferListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

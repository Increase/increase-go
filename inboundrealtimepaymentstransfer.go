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

// InboundRealTimePaymentsTransferService contains methods and other services that
// help with interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInboundRealTimePaymentsTransferService] method instead.
type InboundRealTimePaymentsTransferService struct {
	Options []option.RequestOption
}

// NewInboundRealTimePaymentsTransferService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewInboundRealTimePaymentsTransferService(opts ...option.RequestOption) (r *InboundRealTimePaymentsTransferService) {
	r = &InboundRealTimePaymentsTransferService{}
	r.Options = opts
	return
}

// Retrieve an Inbound Real-Time Payments Transfer
func (r *InboundRealTimePaymentsTransferService) Get(ctx context.Context, inboundRealTimePaymentsTransferID string, opts ...option.RequestOption) (res *InboundRealTimePaymentsTransfer, err error) {
	opts = slices.Concat(r.Options, opts)
	if inboundRealTimePaymentsTransferID == "" {
		err = errors.New("missing required inbound_real_time_payments_transfer_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("inbound_real_time_payments_transfers/%s", inboundRealTimePaymentsTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List Inbound Real-Time Payments Transfers
func (r *InboundRealTimePaymentsTransferService) List(ctx context.Context, query InboundRealTimePaymentsTransferListParams, opts ...option.RequestOption) (res *pagination.Page[InboundRealTimePaymentsTransfer], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "inbound_real_time_payments_transfers"
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

// List Inbound Real-Time Payments Transfers
func (r *InboundRealTimePaymentsTransferService) ListAutoPaging(ctx context.Context, query InboundRealTimePaymentsTransferListParams, opts ...option.RequestOption) *pagination.PageAutoPager[InboundRealTimePaymentsTransfer] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Return an Inbound Real-Time Payments Transfer
func (r *InboundRealTimePaymentsTransferService) Return(ctx context.Context, inboundRealTimePaymentsTransferID string, body InboundRealTimePaymentsTransferReturnParams, opts ...option.RequestOption) (res *RealTimePaymentsTransfer, err error) {
	opts = slices.Concat(r.Options, opts)
	if inboundRealTimePaymentsTransferID == "" {
		err = errors.New("missing required inbound_real_time_payments_transfer_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("inbound_real_time_payments_transfers/%s/return", inboundRealTimePaymentsTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// An Inbound Real-Time Payments Transfer is a Real-Time Payments transfer
// initiated outside of Increase to your account.
type InboundRealTimePaymentsTransfer struct {
	// The inbound Real-Time Payments transfer's identifier.
	ID string `json:"id" api:"required"`
	// The Account to which the transfer was sent.
	AccountID string `json:"account_id" api:"required"`
	// The identifier of the Account Number to which this transfer was sent.
	AccountNumberID string `json:"account_number_id" api:"required"`
	// The amount in USD cents.
	Amount int64 `json:"amount" api:"required"`
	// If your transfer is confirmed, this will contain details of the confirmation.
	Confirmation InboundRealTimePaymentsTransferConfirmation `json:"confirmation" api:"required,nullable"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The name the sender of the transfer specified as the recipient of the transfer.
	CreditorName string `json:"creditor_name" api:"required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the transfer's
	// currency. This will always be "USD" for a Real-Time Payments transfer.
	Currency InboundRealTimePaymentsTransferCurrency `json:"currency" api:"required"`
	// The account number of the account that sent the transfer.
	DebtorAccountNumber string `json:"debtor_account_number" api:"required"`
	// The name provided by the sender of the transfer.
	DebtorName string `json:"debtor_name" api:"required"`
	// The routing number of the account that sent the transfer.
	DebtorRoutingNumber string `json:"debtor_routing_number" api:"required"`
	// If your transfer is declined, this will contain details of the decline.
	Decline InboundRealTimePaymentsTransferDecline `json:"decline" api:"required,nullable"`
	// A free-form reference string set by the sender, to help identify the transfer.
	EndToEndIdentification string `json:"end_to_end_identification" api:"required,nullable"`
	// The identifier of the Real-Time Payments Transfer this transfer returns, if any.
	ReturnOfRealTimePaymentsTransferID string `json:"return_of_real_time_payments_transfer_id" api:"required,nullable"`
	// The lifecycle status of the transfer.
	Status InboundRealTimePaymentsTransferStatus `json:"status" api:"required"`
	// The Real-Time Payments network identification of the transfer.
	TransactionIdentification string `json:"transaction_identification" api:"required"`
	// A constant representing the object's type. For this resource it will always be
	// `inbound_real_time_payments_transfer`.
	Type InboundRealTimePaymentsTransferType `json:"type" api:"required"`
	// Additional information included with the transfer.
	UnstructuredRemittanceInformation string                              `json:"unstructured_remittance_information" api:"required,nullable"`
	ExtraFields                       map[string]interface{}              `json:"-" api:"extrafields"`
	JSON                              inboundRealTimePaymentsTransferJSON `json:"-"`
}

// inboundRealTimePaymentsTransferJSON contains the JSON metadata for the struct
// [InboundRealTimePaymentsTransfer]
type inboundRealTimePaymentsTransferJSON struct {
	ID                                 apijson.Field
	AccountID                          apijson.Field
	AccountNumberID                    apijson.Field
	Amount                             apijson.Field
	Confirmation                       apijson.Field
	CreatedAt                          apijson.Field
	CreditorName                       apijson.Field
	Currency                           apijson.Field
	DebtorAccountNumber                apijson.Field
	DebtorName                         apijson.Field
	DebtorRoutingNumber                apijson.Field
	Decline                            apijson.Field
	EndToEndIdentification             apijson.Field
	ReturnOfRealTimePaymentsTransferID apijson.Field
	Status                             apijson.Field
	TransactionIdentification          apijson.Field
	Type                               apijson.Field
	UnstructuredRemittanceInformation  apijson.Field
	raw                                string
	ExtraFields                        map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboundRealTimePaymentsTransferJSON) RawJSON() string {
	return r.raw
}

// If your transfer is confirmed, this will contain details of the confirmation.
type InboundRealTimePaymentsTransferConfirmation struct {
	// The time at which the transfer was confirmed.
	ConfirmedAt time.Time `json:"confirmed_at" api:"required" format:"date-time"`
	// The id of the transaction for the confirmed transfer.
	TransactionID string                                          `json:"transaction_id" api:"required"`
	JSON          inboundRealTimePaymentsTransferConfirmationJSON `json:"-"`
}

// inboundRealTimePaymentsTransferConfirmationJSON contains the JSON metadata for
// the struct [InboundRealTimePaymentsTransferConfirmation]
type inboundRealTimePaymentsTransferConfirmationJSON struct {
	ConfirmedAt   apijson.Field
	TransactionID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferConfirmation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboundRealTimePaymentsTransferConfirmationJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the transfer's
// currency. This will always be "USD" for a Real-Time Payments transfer.
type InboundRealTimePaymentsTransferCurrency string

const (
	InboundRealTimePaymentsTransferCurrencyUsd InboundRealTimePaymentsTransferCurrency = "USD"
)

func (r InboundRealTimePaymentsTransferCurrency) IsKnown() bool {
	switch r {
	case InboundRealTimePaymentsTransferCurrencyUsd:
		return true
	}
	return false
}

// If your transfer is declined, this will contain details of the decline.
type InboundRealTimePaymentsTransferDecline struct {
	// The time at which the transfer was declined.
	DeclinedAt time.Time `json:"declined_at" api:"required" format:"date-time"`
	// The id of the transaction for the declined transfer.
	DeclinedTransactionID string `json:"declined_transaction_id" api:"required"`
	// The reason for the transfer decline.
	Reason InboundRealTimePaymentsTransferDeclineReason `json:"reason" api:"required"`
	JSON   inboundRealTimePaymentsTransferDeclineJSON   `json:"-"`
}

// inboundRealTimePaymentsTransferDeclineJSON contains the JSON metadata for the
// struct [InboundRealTimePaymentsTransferDecline]
type inboundRealTimePaymentsTransferDeclineJSON struct {
	DeclinedAt            apijson.Field
	DeclinedTransactionID apijson.Field
	Reason                apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *InboundRealTimePaymentsTransferDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboundRealTimePaymentsTransferDeclineJSON) RawJSON() string {
	return r.raw
}

// The reason for the transfer decline.
type InboundRealTimePaymentsTransferDeclineReason string

const (
	InboundRealTimePaymentsTransferDeclineReasonAccountNumberCanceled      InboundRealTimePaymentsTransferDeclineReason = "account_number_canceled"
	InboundRealTimePaymentsTransferDeclineReasonAccountNumberDisabled      InboundRealTimePaymentsTransferDeclineReason = "account_number_disabled"
	InboundRealTimePaymentsTransferDeclineReasonAccountRestricted          InboundRealTimePaymentsTransferDeclineReason = "account_restricted"
	InboundRealTimePaymentsTransferDeclineReasonGroupLocked                InboundRealTimePaymentsTransferDeclineReason = "group_locked"
	InboundRealTimePaymentsTransferDeclineReasonEntityNotActive            InboundRealTimePaymentsTransferDeclineReason = "entity_not_active"
	InboundRealTimePaymentsTransferDeclineReasonRealTimePaymentsNotEnabled InboundRealTimePaymentsTransferDeclineReason = "real_time_payments_not_enabled"
	InboundRealTimePaymentsTransferDeclineReasonTransactionNotAllowed      InboundRealTimePaymentsTransferDeclineReason = "transaction_not_allowed"
)

func (r InboundRealTimePaymentsTransferDeclineReason) IsKnown() bool {
	switch r {
	case InboundRealTimePaymentsTransferDeclineReasonAccountNumberCanceled, InboundRealTimePaymentsTransferDeclineReasonAccountNumberDisabled, InboundRealTimePaymentsTransferDeclineReasonAccountRestricted, InboundRealTimePaymentsTransferDeclineReasonGroupLocked, InboundRealTimePaymentsTransferDeclineReasonEntityNotActive, InboundRealTimePaymentsTransferDeclineReasonRealTimePaymentsNotEnabled, InboundRealTimePaymentsTransferDeclineReasonTransactionNotAllowed:
		return true
	}
	return false
}

// The lifecycle status of the transfer.
type InboundRealTimePaymentsTransferStatus string

const (
	InboundRealTimePaymentsTransferStatusPendingConfirming InboundRealTimePaymentsTransferStatus = "pending_confirming"
	InboundRealTimePaymentsTransferStatusTimedOut          InboundRealTimePaymentsTransferStatus = "timed_out"
	InboundRealTimePaymentsTransferStatusConfirmed         InboundRealTimePaymentsTransferStatus = "confirmed"
	InboundRealTimePaymentsTransferStatusDeclined          InboundRealTimePaymentsTransferStatus = "declined"
)

func (r InboundRealTimePaymentsTransferStatus) IsKnown() bool {
	switch r {
	case InboundRealTimePaymentsTransferStatusPendingConfirming, InboundRealTimePaymentsTransferStatusTimedOut, InboundRealTimePaymentsTransferStatusConfirmed, InboundRealTimePaymentsTransferStatusDeclined:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `inbound_real_time_payments_transfer`.
type InboundRealTimePaymentsTransferType string

const (
	InboundRealTimePaymentsTransferTypeInboundRealTimePaymentsTransfer InboundRealTimePaymentsTransferType = "inbound_real_time_payments_transfer"
)

func (r InboundRealTimePaymentsTransferType) IsKnown() bool {
	switch r {
	case InboundRealTimePaymentsTransferTypeInboundRealTimePaymentsTransfer:
		return true
	}
	return false
}

// Real-Time Payments transfers move funds, within seconds, between your Increase
// account and any other account on the Real-Time Payments network.
type RealTimePaymentsTransfer struct {
	// The Real-Time Payments Transfer's identifier.
	ID string `json:"id" api:"required"`
	// The Account from which the transfer was sent.
	AccountID string `json:"account_id" api:"required"`
	// The destination account number.
	AccountNumber string `json:"account_number" api:"required"`
	// If the transfer is acknowledged by the recipient bank, this will contain
	// supplemental details.
	Acknowledgement RealTimePaymentsTransferAcknowledgement `json:"acknowledgement" api:"required,nullable"`
	// The transfer amount in USD cents.
	Amount int64 `json:"amount" api:"required"`
	// If your account requires approvals for transfers and the transfer was approved,
	// this will contain details of the approval.
	Approval RealTimePaymentsTransferApproval `json:"approval" api:"required,nullable"`
	// If your account requires approvals for transfers and the transfer was not
	// approved, this will contain details of the cancellation.
	Cancellation RealTimePaymentsTransferCancellation `json:"cancellation" api:"required,nullable"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// What object created the transfer, either via the API or the dashboard.
	CreatedBy RealTimePaymentsTransferCreatedBy `json:"created_by" api:"required,nullable"`
	// The name of the transfer's recipient. This is set by the sender when creating
	// the transfer.
	CreditorName string `json:"creditor_name" api:"required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transfer's
	// currency. For real-time payments transfers this is always equal to `USD`.
	Currency RealTimePaymentsTransferCurrency `json:"currency" api:"required"`
	// The name of the transfer's sender. If not provided, defaults to the name of the
	// account's entity.
	DebtorName string `json:"debtor_name" api:"required,nullable"`
	// The identifier of the External Account the transfer was made to, if any.
	ExternalAccountID string `json:"external_account_id" api:"required,nullable"`
	// The idempotency key you chose for this object. This value is unique across
	// Increase and is used to ensure that a request is only processed once. Learn more
	// about [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey string `json:"idempotency_key" api:"required,nullable"`
	// The identifier of the Inbound Real-Time Payments Request for Payment this
	// transfer was sent in response to, if any.
	InboundRealTimePaymentsRequestForPaymentID string `json:"inbound_real_time_payments_request_for_payment_id" api:"required,nullable"`
	// The ID for the pending transaction representing the transfer. A pending
	// transaction is created when the transfer
	// [requires approval](https://increase.com/documentation/transfer-approvals#transfer-approvals)
	// by someone else in your organization.
	PendingTransactionID string `json:"pending_transaction_id" api:"required,nullable"`
	// If the transfer is rejected by Real-Time Payments or the destination financial
	// institution, this will contain supplemental details.
	Rejection RealTimePaymentsTransferRejection `json:"rejection" api:"required,nullable"`
	// The identifier of the Inbound Real-Time Payments Transfer this transfer returns,
	// if any.
	ReturnOfInboundRealTimePaymentsTransferID string `json:"return_of_inbound_real_time_payments_transfer_id" api:"required,nullable"`
	// The destination American Bankers' Association (ABA) Routing Transit Number
	// (RTN).
	RoutingNumber string `json:"routing_number" api:"required"`
	// The Account Number the recipient will see as having sent the transfer.
	SourceAccountNumberID string `json:"source_account_number_id" api:"required"`
	// The lifecycle status of the transfer.
	Status RealTimePaymentsTransferStatus `json:"status" api:"required"`
	// After the transfer is submitted to Real-Time Payments, this will contain
	// supplemental details.
	Submission RealTimePaymentsTransferSubmission `json:"submission" api:"required,nullable"`
	// The Transaction funding the transfer once it is complete.
	TransactionID string `json:"transaction_id" api:"required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `real_time_payments_transfer`.
	Type RealTimePaymentsTransferType `json:"type" api:"required"`
	// The name of the ultimate recipient of the transfer. Set this if the creditor is
	// an intermediary receiving the payment for someone else.
	UltimateCreditorName string `json:"ultimate_creditor_name" api:"required,nullable"`
	// The name of the ultimate sender of the transfer. Set this if the funds are being
	// sent on behalf of someone who is not the account holder at Increase.
	UltimateDebtorName string `json:"ultimate_debtor_name" api:"required,nullable"`
	// Unstructured information that will show on the recipient's bank statement.
	UnstructuredRemittanceInformation string                       `json:"unstructured_remittance_information" api:"required"`
	ExtraFields                       map[string]interface{}       `json:"-" api:"extrafields"`
	JSON                              realTimePaymentsTransferJSON `json:"-"`
}

// realTimePaymentsTransferJSON contains the JSON metadata for the struct
// [RealTimePaymentsTransfer]
type realTimePaymentsTransferJSON struct {
	ID                                         apijson.Field
	AccountID                                  apijson.Field
	AccountNumber                              apijson.Field
	Acknowledgement                            apijson.Field
	Amount                                     apijson.Field
	Approval                                   apijson.Field
	Cancellation                               apijson.Field
	CreatedAt                                  apijson.Field
	CreatedBy                                  apijson.Field
	CreditorName                               apijson.Field
	Currency                                   apijson.Field
	DebtorName                                 apijson.Field
	ExternalAccountID                          apijson.Field
	IdempotencyKey                             apijson.Field
	InboundRealTimePaymentsRequestForPaymentID apijson.Field
	PendingTransactionID                       apijson.Field
	Rejection                                  apijson.Field
	ReturnOfInboundRealTimePaymentsTransferID  apijson.Field
	RoutingNumber                              apijson.Field
	SourceAccountNumberID                      apijson.Field
	Status                                     apijson.Field
	Submission                                 apijson.Field
	TransactionID                              apijson.Field
	Type                                       apijson.Field
	UltimateCreditorName                       apijson.Field
	UltimateDebtorName                         apijson.Field
	UnstructuredRemittanceInformation          apijson.Field
	raw                                        string
	ExtraFields                                map[string]apijson.Field
}

func (r *RealTimePaymentsTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimePaymentsTransferJSON) RawJSON() string {
	return r.raw
}

// If the transfer is acknowledged by the recipient bank, this will contain
// supplemental details.
type RealTimePaymentsTransferAcknowledgement struct {
	// When the transfer was acknowledged.
	AcknowledgedAt time.Time                                   `json:"acknowledged_at" api:"required" format:"date-time"`
	JSON           realTimePaymentsTransferAcknowledgementJSON `json:"-"`
}

// realTimePaymentsTransferAcknowledgementJSON contains the JSON metadata for the
// struct [RealTimePaymentsTransferAcknowledgement]
type realTimePaymentsTransferAcknowledgementJSON struct {
	AcknowledgedAt apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RealTimePaymentsTransferAcknowledgement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimePaymentsTransferAcknowledgementJSON) RawJSON() string {
	return r.raw
}

// If your account requires approvals for transfers and the transfer was approved,
// this will contain details of the approval.
type RealTimePaymentsTransferApproval struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was approved.
	ApprovedAt time.Time `json:"approved_at" api:"required" format:"date-time"`
	// If the Transfer was approved by a user in the dashboard, the email address of
	// that user.
	ApprovedBy string                               `json:"approved_by" api:"required,nullable"`
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
	CanceledAt time.Time `json:"canceled_at" api:"required" format:"date-time"`
	// If the Transfer was canceled by a user in the dashboard, the email address of
	// that user.
	CanceledBy string                                   `json:"canceled_by" api:"required,nullable"`
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
	// The type of object that created this transfer.
	Category RealTimePaymentsTransferCreatedByCategory `json:"category" api:"required"`
	// If present, details about the API key that created the transfer.
	APIKey RealTimePaymentsTransferCreatedByAPIKey `json:"api_key" api:"nullable"`
	// If present, details about the OAuth Application that created the transfer.
	OAuthApplication RealTimePaymentsTransferCreatedByOAuthApplication `json:"oauth_application" api:"nullable"`
	// If present, details about the User that created the transfer.
	User RealTimePaymentsTransferCreatedByUser `json:"user" api:"nullable"`
	JSON realTimePaymentsTransferCreatedByJSON `json:"-"`
}

// realTimePaymentsTransferCreatedByJSON contains the JSON metadata for the struct
// [RealTimePaymentsTransferCreatedBy]
type realTimePaymentsTransferCreatedByJSON struct {
	Category         apijson.Field
	APIKey           apijson.Field
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

// The type of object that created this transfer.
type RealTimePaymentsTransferCreatedByCategory string

const (
	RealTimePaymentsTransferCreatedByCategoryAPIKey           RealTimePaymentsTransferCreatedByCategory = "api_key"
	RealTimePaymentsTransferCreatedByCategoryOAuthApplication RealTimePaymentsTransferCreatedByCategory = "oauth_application"
	RealTimePaymentsTransferCreatedByCategoryUser             RealTimePaymentsTransferCreatedByCategory = "user"
)

func (r RealTimePaymentsTransferCreatedByCategory) IsKnown() bool {
	switch r {
	case RealTimePaymentsTransferCreatedByCategoryAPIKey, RealTimePaymentsTransferCreatedByCategoryOAuthApplication, RealTimePaymentsTransferCreatedByCategoryUser:
		return true
	}
	return false
}

// If present, details about the API key that created the transfer.
type RealTimePaymentsTransferCreatedByAPIKey struct {
	// The description set for the API key when it was created.
	Description string                                      `json:"description" api:"required,nullable"`
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

// If present, details about the OAuth Application that created the transfer.
type RealTimePaymentsTransferCreatedByOAuthApplication struct {
	// The name of the OAuth Application.
	Name string                                                `json:"name" api:"required"`
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
	Email string                                    `json:"email" api:"required"`
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
	RealTimePaymentsTransferCurrencyUsd RealTimePaymentsTransferCurrency = "USD"
)

func (r RealTimePaymentsTransferCurrency) IsKnown() bool {
	switch r {
	case RealTimePaymentsTransferCurrencyUsd:
		return true
	}
	return false
}

// If the transfer is rejected by Real-Time Payments or the destination financial
// institution, this will contain supplemental details.
type RealTimePaymentsTransferRejection struct {
	// Additional information about the rejection provided by the recipient bank when
	// the `reject_reason_code` is `NARRATIVE`.
	RejectReasonAdditionalInformation string `json:"reject_reason_additional_information" api:"required,nullable"`
	// The reason the transfer was rejected as provided by the recipient bank or the
	// Real-Time Payments network.
	RejectReasonCode RealTimePaymentsTransferRejectionRejectReasonCode `json:"reject_reason_code" api:"required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was rejected.
	RejectedAt time.Time                             `json:"rejected_at" api:"required,nullable" format:"date-time"`
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
	RealTimePaymentsTransferRejectionRejectReasonCodeAccountClosed                                 RealTimePaymentsTransferRejectionRejectReasonCode = "account_closed"
	RealTimePaymentsTransferRejectionRejectReasonCodeAccountBlocked                                RealTimePaymentsTransferRejectionRejectReasonCode = "account_blocked"
	RealTimePaymentsTransferRejectionRejectReasonCodeInvalidCreditorAccountType                    RealTimePaymentsTransferRejectionRejectReasonCode = "invalid_creditor_account_type"
	RealTimePaymentsTransferRejectionRejectReasonCodeInvalidCreditorAccountNumber                  RealTimePaymentsTransferRejectionRejectReasonCode = "invalid_creditor_account_number"
	RealTimePaymentsTransferRejectionRejectReasonCodeInvalidCreditorFinancialInstitutionIdentifier RealTimePaymentsTransferRejectionRejectReasonCode = "invalid_creditor_financial_institution_identifier"
	RealTimePaymentsTransferRejectionRejectReasonCodeEndCustomerDeceased                           RealTimePaymentsTransferRejectionRejectReasonCode = "end_customer_deceased"
	RealTimePaymentsTransferRejectionRejectReasonCodeNarrative                                     RealTimePaymentsTransferRejectionRejectReasonCode = "narrative"
	RealTimePaymentsTransferRejectionRejectReasonCodeTransactionForbidden                          RealTimePaymentsTransferRejectionRejectReasonCode = "transaction_forbidden"
	RealTimePaymentsTransferRejectionRejectReasonCodeTransactionTypeNotSupported                   RealTimePaymentsTransferRejectionRejectReasonCode = "transaction_type_not_supported"
	RealTimePaymentsTransferRejectionRejectReasonCodeUnexpectedAmount                              RealTimePaymentsTransferRejectionRejectReasonCode = "unexpected_amount"
	RealTimePaymentsTransferRejectionRejectReasonCodeAmountExceedsBankLimits                       RealTimePaymentsTransferRejectionRejectReasonCode = "amount_exceeds_bank_limits"
	RealTimePaymentsTransferRejectionRejectReasonCodeInvalidCreditorAddress                        RealTimePaymentsTransferRejectionRejectReasonCode = "invalid_creditor_address"
	RealTimePaymentsTransferRejectionRejectReasonCodeUnknownEndCustomer                            RealTimePaymentsTransferRejectionRejectReasonCode = "unknown_end_customer"
	RealTimePaymentsTransferRejectionRejectReasonCodeInvalidDebtorAddress                          RealTimePaymentsTransferRejectionRejectReasonCode = "invalid_debtor_address"
	RealTimePaymentsTransferRejectionRejectReasonCodeTimeout                                       RealTimePaymentsTransferRejectionRejectReasonCode = "timeout"
	RealTimePaymentsTransferRejectionRejectReasonCodeUnsupportedMessageForRecipient                RealTimePaymentsTransferRejectionRejectReasonCode = "unsupported_message_for_recipient"
	RealTimePaymentsTransferRejectionRejectReasonCodeRecipientConnectionNotAvailable               RealTimePaymentsTransferRejectionRejectReasonCode = "recipient_connection_not_available"
	RealTimePaymentsTransferRejectionRejectReasonCodeRealTimePaymentsSuspended                     RealTimePaymentsTransferRejectionRejectReasonCode = "real_time_payments_suspended"
	RealTimePaymentsTransferRejectionRejectReasonCodeInstructedAgentSignedOff                      RealTimePaymentsTransferRejectionRejectReasonCode = "instructed_agent_signed_off"
	RealTimePaymentsTransferRejectionRejectReasonCodeProcessingError                               RealTimePaymentsTransferRejectionRejectReasonCode = "processing_error"
	RealTimePaymentsTransferRejectionRejectReasonCodeOther                                         RealTimePaymentsTransferRejectionRejectReasonCode = "other"
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
	RealTimePaymentsTransferStatusPendingApproval   RealTimePaymentsTransferStatus = "pending_approval"
	RealTimePaymentsTransferStatusCanceled          RealTimePaymentsTransferStatus = "canceled"
	RealTimePaymentsTransferStatusPendingSubmission RealTimePaymentsTransferStatus = "pending_submission"
	RealTimePaymentsTransferStatusPendingReviewing  RealTimePaymentsTransferStatus = "pending_reviewing"
	RealTimePaymentsTransferStatusRequiresAttention RealTimePaymentsTransferStatus = "requires_attention"
	RealTimePaymentsTransferStatusRejected          RealTimePaymentsTransferStatus = "rejected"
	RealTimePaymentsTransferStatusSubmitted         RealTimePaymentsTransferStatus = "submitted"
	RealTimePaymentsTransferStatusComplete          RealTimePaymentsTransferStatus = "complete"
)

func (r RealTimePaymentsTransferStatus) IsKnown() bool {
	switch r {
	case RealTimePaymentsTransferStatusPendingApproval, RealTimePaymentsTransferStatusCanceled, RealTimePaymentsTransferStatusPendingSubmission, RealTimePaymentsTransferStatusPendingReviewing, RealTimePaymentsTransferStatusRequiresAttention, RealTimePaymentsTransferStatusRejected, RealTimePaymentsTransferStatusSubmitted, RealTimePaymentsTransferStatusComplete:
		return true
	}
	return false
}

// After the transfer is submitted to Real-Time Payments, this will contain
// supplemental details.
type RealTimePaymentsTransferSubmission struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was submitted to The Clearing House.
	SubmittedAt time.Time `json:"submitted_at" api:"required,nullable" format:"date-time"`
	// The Real-Time Payments network identification of the transfer.
	TransactionIdentification string                                 `json:"transaction_identification" api:"required"`
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

type InboundRealTimePaymentsTransferListParams struct {
	// Filter Inbound Real-Time Payments Transfers to those belonging to the specified
	// Account.
	AccountID param.Field[string] `query:"account_id"`
	// Filter Inbound Real-Time Payments Transfers to ones belonging to the specified
	// Account Number.
	AccountNumberID param.Field[string]                                             `query:"account_number_id"`
	CreatedAt       param.Field[InboundRealTimePaymentsTransferListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	//
	// Defaults to `100`.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [InboundRealTimePaymentsTransferListParams]'s query
// parameters as `url.Values`.
func (r InboundRealTimePaymentsTransferListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type InboundRealTimePaymentsTransferListParamsCreatedAt struct {
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

// URLQuery serializes [InboundRealTimePaymentsTransferListParamsCreatedAt]'s query
// parameters as `url.Values`.
func (r InboundRealTimePaymentsTransferListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type InboundRealTimePaymentsTransferReturnParams struct {
	// The amount to return, in USD cents. Defaults to the full amount of the Inbound
	// Real-Time Payments Transfer, and can't be higher than it.
	Amount param.Field[int64] `json:"amount"`
}

func (r InboundRealTimePaymentsTransferReturnParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

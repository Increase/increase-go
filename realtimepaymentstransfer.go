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

// RealTimePaymentsTransferService contains methods and other services that help
// with interacting with the increase API. Note, unlike clients, this service does
// not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRealTimePaymentsTransferService] method instead.
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

// Create a Real Time Payments Transfer
func (r *RealTimePaymentsTransferService) New(ctx context.Context, body RealTimePaymentsTransferNewParams, opts ...option.RequestOption) (res *RealTimePaymentsTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := "real_time_payments_transfers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a Real Time Payments Transfer
func (r *RealTimePaymentsTransferService) Get(ctx context.Context, realTimePaymentsTransferID string, opts ...option.RequestOption) (res *RealTimePaymentsTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("real_time_payments_transfers/%s", realTimePaymentsTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Real Time Payments Transfers
func (r *RealTimePaymentsTransferService) List(ctx context.Context, query RealTimePaymentsTransferListParams, opts ...option.RequestOption) (res *shared.Page[RealTimePaymentsTransfer], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
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

// List Real Time Payments Transfers
func (r *RealTimePaymentsTransferService) ListAutoPaging(ctx context.Context, query RealTimePaymentsTransferListParams, opts ...option.RequestOption) *shared.PageAutoPager[RealTimePaymentsTransfer] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Real Time Payments transfers move funds, within seconds, between your Increase
// account and any other account on the Real Time Payments network.
type RealTimePaymentsTransfer struct {
	// The Real Time Payments Transfer's identifier.
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
	// The name of the transfer's recipient as provided by the sender.
	CreditorName string `json:"creditor_name,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transfer's
	// currency. For real time payments transfers this is always equal to `USD`.
	Currency RealTimePaymentsTransferCurrency `json:"currency,required"`
	// The destination account number.
	DestinationAccountNumber string `json:"destination_account_number,required"`
	// The destination American Bankers' Association (ABA) Routing Transit Number
	// (RTN).
	DestinationRoutingNumber string `json:"destination_routing_number,required"`
	// The identifier of the External Account the transfer was made to, if any.
	ExternalAccountID string `json:"external_account_id,required,nullable"`
	// If the transfer is rejected by Real Time Payments or the destination financial
	// institution, this will contain supplemental details.
	Rejection RealTimePaymentsTransferRejection `json:"rejection,required,nullable"`
	// Unstructured information that will show on the recipient's bank statement.
	RemittanceInformation string `json:"remittance_information,required"`
	// The Account Number the recipient will see as having sent the transfer.
	SourceAccountNumberID string `json:"source_account_number_id,required"`
	// The lifecycle status of the transfer.
	Status RealTimePaymentsTransferStatus `json:"status,required"`
	// After the transfer is submitted to Real Time Payments, this will contain
	// supplemental details.
	Submission RealTimePaymentsTransferSubmission `json:"submission,required,nullable"`
	// The Transaction funding the transfer once it is complete.
	TransactionID string `json:"transaction_id,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `real_time_payments_transfer`.
	Type RealTimePaymentsTransferType `json:"type,required"`
	JSON realTimePaymentsTransferJSON
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
	CreditorName             apijson.Field
	Currency                 apijson.Field
	DestinationAccountNumber apijson.Field
	DestinationRoutingNumber apijson.Field
	ExternalAccountID        apijson.Field
	Rejection                apijson.Field
	RemittanceInformation    apijson.Field
	SourceAccountNumberID    apijson.Field
	Status                   apijson.Field
	Submission               apijson.Field
	TransactionID            apijson.Field
	Type                     apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *RealTimePaymentsTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// If your account requires approvals for transfers and the transfer was approved,
// this will contain details of the approval.
type RealTimePaymentsTransferApproval struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was approved.
	ApprovedAt time.Time `json:"approved_at,required" format:"date-time"`
	// If the Transfer was approved by a user in the dashboard, the email address of
	// that user.
	ApprovedBy string `json:"approved_by,required,nullable"`
	JSON       realTimePaymentsTransferApprovalJSON
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

// If your account requires approvals for transfers and the transfer was not
// approved, this will contain details of the cancellation.
type RealTimePaymentsTransferCancellation struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Transfer was canceled.
	CanceledAt time.Time `json:"canceled_at,required" format:"date-time"`
	// If the Transfer was canceled by a user in the dashboard, the email address of
	// that user.
	CanceledBy string `json:"canceled_by,required,nullable"`
	JSON       realTimePaymentsTransferCancellationJSON
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

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transfer's
// currency. For real time payments transfers this is always equal to `USD`.
type RealTimePaymentsTransferCurrency string

const (
	RealTimePaymentsTransferCurrencyCad RealTimePaymentsTransferCurrency = "CAD"
	RealTimePaymentsTransferCurrencyChf RealTimePaymentsTransferCurrency = "CHF"
	RealTimePaymentsTransferCurrencyEur RealTimePaymentsTransferCurrency = "EUR"
	RealTimePaymentsTransferCurrencyGbp RealTimePaymentsTransferCurrency = "GBP"
	RealTimePaymentsTransferCurrencyJpy RealTimePaymentsTransferCurrency = "JPY"
	RealTimePaymentsTransferCurrencyUsd RealTimePaymentsTransferCurrency = "USD"
)

// If the transfer is rejected by Real Time Payments or the destination financial
// institution, this will contain supplemental details.
type RealTimePaymentsTransferRejection struct {
	// Additional information about the rejection provided by the recipient bank when
	// the `reject_reason_code` is `NARRATIVE`.
	RejectReasonAdditionalInformation string `json:"reject_reason_additional_information,required,nullable"`
	// The reason the transfer was rejected as provided by the recipient bank or the
	// Real Time Payments network.
	RejectReasonCode RealTimePaymentsTransferRejectionRejectReasonCode `json:"reject_reason_code,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was rejected.
	RejectedAt time.Time `json:"rejected_at,required,nullable" format:"date-time"`
	JSON       realTimePaymentsTransferRejectionJSON
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

// The reason the transfer was rejected as provided by the recipient bank or the
// Real Time Payments network.
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

// The lifecycle status of the transfer.
type RealTimePaymentsTransferStatus string

const (
	RealTimePaymentsTransferStatusPendingApproval   RealTimePaymentsTransferStatus = "pending_approval"
	RealTimePaymentsTransferStatusCanceled          RealTimePaymentsTransferStatus = "canceled"
	RealTimePaymentsTransferStatusPendingSubmission RealTimePaymentsTransferStatus = "pending_submission"
	RealTimePaymentsTransferStatusSubmitted         RealTimePaymentsTransferStatus = "submitted"
	RealTimePaymentsTransferStatusComplete          RealTimePaymentsTransferStatus = "complete"
	RealTimePaymentsTransferStatusRejected          RealTimePaymentsTransferStatus = "rejected"
	RealTimePaymentsTransferStatusRequiresAttention RealTimePaymentsTransferStatus = "requires_attention"
)

// After the transfer is submitted to Real Time Payments, this will contain
// supplemental details.
type RealTimePaymentsTransferSubmission struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was submitted to The Clearing House.
	SubmittedAt time.Time `json:"submitted_at,required,nullable" format:"date-time"`
	// The Real Time Payments network identification of the transfer.
	TransactionIdentification string `json:"transaction_identification,required"`
	JSON                      realTimePaymentsTransferSubmissionJSON
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

// A constant representing the object's type. For this resource it will always be
// `real_time_payments_transfer`.
type RealTimePaymentsTransferType string

const (
	RealTimePaymentsTransferTypeRealTimePaymentsTransfer RealTimePaymentsTransferType = "real_time_payments_transfer"
)

type RealTimePaymentsTransferNewParams struct {
	// The transfer amount in USD cents. For Real Time Payments transfers, must be
	// positive.
	Amount param.Field[int64] `json:"amount,required"`
	// The name of the transfer's recipient.
	CreditorName param.Field[string] `json:"creditor_name,required"`
	// Unstructured information that will show on the recipient's bank statement.
	RemittanceInformation param.Field[string] `json:"remittance_information,required"`
	// The identifier of the Account Number from which to send the transfer.
	SourceAccountNumberID param.Field[string] `json:"source_account_number_id,required"`
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
}

func (r RealTimePaymentsTransferNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RealTimePaymentsTransferListParams struct {
	// Filter Real Time Payments Transfers to those belonging to the specified Account.
	AccountID param.Field[string]                                      `query:"account_id"`
	CreatedAt param.Field[RealTimePaymentsTransferListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Filter Real Time Payments Transfers to those made to the specified External
	// Account.
	ExternalAccountID param.Field[string] `query:"external_account_id"`
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

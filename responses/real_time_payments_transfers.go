package responses

import (
	"time"

	pjson "github.com/increase/increase-go/core/json"
)

type RealTimePaymentsTransfer struct {
	// A constant representing the object's type. For this resource it will always be
	// `real_time_payments_transfer`.
	Type RealTimePaymentsTransferType `json:"type,required"`
	// The Real Time Payments Transfer's identifier.
	ID string `json:"id,required"`
	// If your account requires approvals for transfers and the transfer was approved,
	// this will contain details of the approval.
	Approval RealTimePaymentsTransferApproval `json:"approval,required,nullable"`
	// If your account requires approvals for transfers and the transfer was not
	// approved, this will contain details of the cancellation.
	Cancellation RealTimePaymentsTransferCancellation `json:"cancellation,required,nullable"`
	// The lifecycle status of the transfer.
	Status RealTimePaymentsTransferStatus `json:"status,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The Account from which the transfer was sent.
	AccountID string `json:"account_id,required"`
	// The identifier of the External Account the transfer was made to, if any.
	ExternalAccountID string `json:"external_account_id,required,nullable"`
	// The Account Number the recipient will see as having sent the transfer.
	SourceAccountNumberID string `json:"source_account_number_id,required"`
	// The name of the transfer's recipient as provided by the sender.
	CreditorName string `json:"creditor_name,required"`
	// Unstructured information that will show on the recipient's bank statement.
	RemittanceInformation string `json:"remittance_information,required"`
	// The transfer amount in USD cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transfer's
	// currency. For real time payments transfers this is always equal to `USD`.
	Currency RealTimePaymentsTransferCurrency `json:"currency,required"`
	// The destination account number.
	DestinationAccountNumber string `json:"destination_account_number,required"`
	// The destination American Bankers' Association (ABA) Routing Transit Number
	// (RTN).
	DestinationRoutingNumber string `json:"destination_routing_number,required"`
	// The Transaction funding the transfer once it is complete.
	TransactionID string `json:"transaction_id,required,nullable"`
	// After the transfer is submitted to Real Time Payments, this will contain
	// supplemental details.
	Submission RealTimePaymentsTransferSubmission `json:"submission,required,nullable"`
	// If the transfer is rejected by Real Time Payments or the destination financial
	// institution, this will contain supplemental details.
	Rejection RealTimePaymentsTransferRejection `json:"rejection,required,nullable"`
	JSON      RealTimePaymentsTransferJSON
}

type RealTimePaymentsTransferJSON struct {
	Type                     pjson.Metadata
	ID                       pjson.Metadata
	Approval                 pjson.Metadata
	Cancellation             pjson.Metadata
	Status                   pjson.Metadata
	CreatedAt                pjson.Metadata
	AccountID                pjson.Metadata
	ExternalAccountID        pjson.Metadata
	SourceAccountNumberID    pjson.Metadata
	CreditorName             pjson.Metadata
	RemittanceInformation    pjson.Metadata
	Amount                   pjson.Metadata
	Currency                 pjson.Metadata
	DestinationAccountNumber pjson.Metadata
	DestinationRoutingNumber pjson.Metadata
	TransactionID            pjson.Metadata
	Submission               pjson.Metadata
	Rejection                pjson.Metadata
	Raw                      []byte
	Extras                   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into RealTimePaymentsTransfer
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *RealTimePaymentsTransfer) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type RealTimePaymentsTransferType string

const (
	RealTimePaymentsTransferTypeRealTimePaymentsTransfer RealTimePaymentsTransferType = "real_time_payments_transfer"
)

type RealTimePaymentsTransferApproval struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was approved.
	ApprovedAt time.Time `json:"approved_at,required" format:"date-time"`
	// If the Transfer was approved by a user in the dashboard, the email address of
	// that user.
	ApprovedBy string `json:"approved_by,required,nullable"`
	JSON       RealTimePaymentsTransferApprovalJSON
}

type RealTimePaymentsTransferApprovalJSON struct {
	ApprovedAt pjson.Metadata
	ApprovedBy pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// RealTimePaymentsTransferApproval using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *RealTimePaymentsTransferApproval) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type RealTimePaymentsTransferCancellation struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Transfer was canceled.
	CanceledAt time.Time `json:"canceled_at,required" format:"date-time"`
	// If the Transfer was canceled by a user in the dashboard, the email address of
	// that user.
	CanceledBy string `json:"canceled_by,required,nullable"`
	JSON       RealTimePaymentsTransferCancellationJSON
}

type RealTimePaymentsTransferCancellationJSON struct {
	CanceledAt pjson.Metadata
	CanceledBy pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// RealTimePaymentsTransferCancellation using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *RealTimePaymentsTransferCancellation) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

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

type RealTimePaymentsTransferCurrency string

const (
	RealTimePaymentsTransferCurrencyCad RealTimePaymentsTransferCurrency = "CAD"
	RealTimePaymentsTransferCurrencyChf RealTimePaymentsTransferCurrency = "CHF"
	RealTimePaymentsTransferCurrencyEur RealTimePaymentsTransferCurrency = "EUR"
	RealTimePaymentsTransferCurrencyGbp RealTimePaymentsTransferCurrency = "GBP"
	RealTimePaymentsTransferCurrencyJpy RealTimePaymentsTransferCurrency = "JPY"
	RealTimePaymentsTransferCurrencyUsd RealTimePaymentsTransferCurrency = "USD"
)

type RealTimePaymentsTransferSubmission struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was submitted to The Clearing House.
	SubmittedAt time.Time `json:"submitted_at,required,nullable" format:"date-time"`
	// The Real Time Payments network identification of the transfer.
	TransactionIdentification string `json:"transaction_identification,required"`
	JSON                      RealTimePaymentsTransferSubmissionJSON
}

type RealTimePaymentsTransferSubmissionJSON struct {
	SubmittedAt               pjson.Metadata
	TransactionIdentification pjson.Metadata
	Raw                       []byte
	Extras                    map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// RealTimePaymentsTransferSubmission using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *RealTimePaymentsTransferSubmission) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type RealTimePaymentsTransferRejection struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was rejected.
	RejectedAt time.Time `json:"rejected_at,required,nullable" format:"date-time"`
	// The reason the transfer was rejected as provided by the recipient bank or the
	// Real Time Payments network.
	RejectReasonCode RealTimePaymentsTransferRejectionRejectReasonCode `json:"reject_reason_code,required"`
	// Additional information about the rejection provided by the recipient bank when
	// the `reject_reason_code` is `NARRATIVE`.
	RejectReasonAdditionalInformation string `json:"reject_reason_additional_information,required,nullable"`
	JSON                              RealTimePaymentsTransferRejectionJSON
}

type RealTimePaymentsTransferRejectionJSON struct {
	RejectedAt                        pjson.Metadata
	RejectReasonCode                  pjson.Metadata
	RejectReasonAdditionalInformation pjson.Metadata
	Raw                               []byte
	Extras                            map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// RealTimePaymentsTransferRejection using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *RealTimePaymentsTransferRejection) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

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

type RealTimePaymentsTransferListResponse struct {
	// The contents of the list.
	Data []RealTimePaymentsTransfer `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       RealTimePaymentsTransferListResponseJSON
}

type RealTimePaymentsTransferListResponseJSON struct {
	Data       pjson.Metadata
	NextCursor pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// RealTimePaymentsTransferListResponse using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *RealTimePaymentsTransferListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

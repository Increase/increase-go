package responses

import (
	"time"

	pjson "github.com/increase/increase-go/core/json"
)

type WireTransfer struct {
	// The wire transfer's identifier.
	ID string `json:"id,required"`
	// The message that will show on the recipient's bank statement.
	MessageToRecipient string `json:"message_to_recipient,required,nullable"`
	// The transfer amount in USD cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transfer's
	// currency. For wire transfers this is always equal to `usd`.
	Currency WireTransferCurrency `json:"currency,required"`
	// The destination account number.
	AccountNumber string `json:"account_number,required"`
	// The beneficiary's name.
	BeneficiaryName string `json:"beneficiary_name,required,nullable"`
	// The beneficiary's address line 1.
	BeneficiaryAddressLine1 string `json:"beneficiary_address_line1,required,nullable"`
	// The beneficiary's address line 2.
	BeneficiaryAddressLine2 string `json:"beneficiary_address_line2,required,nullable"`
	// The beneficiary's address line 3.
	BeneficiaryAddressLine3 string `json:"beneficiary_address_line3,required,nullable"`
	// The Account to which the transfer belongs.
	AccountID string `json:"account_id,required"`
	// The identifier of the External Account the transfer was made to, if any.
	ExternalAccountID string `json:"external_account_id,required,nullable"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber string `json:"routing_number,required"`
	// If your account requires approvals for transfers and the transfer was approved,
	// this will contain details of the approval.
	Approval WireTransferApproval `json:"approval,required,nullable"`
	// If your account requires approvals for transfers and the transfer was not
	// approved, this will contain details of the cancellation.
	Cancellation WireTransferCancellation `json:"cancellation,required,nullable"`
	// If your transfer is reversed, this will contain details of the reversal.
	Reversal WireTransferReversal `json:"reversal,required,nullable"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The transfer's network.
	Network WireTransferNetwork `json:"network,required"`
	// The lifecycle status of the transfer.
	Status WireTransferStatus `json:"status,required"`
	// After the transfer is submitted to Fedwire, this will contain supplemental
	// details.
	Submission WireTransferSubmission `json:"submission,required,nullable"`
	// The ID for the transaction funding the transfer.
	TransactionID string `json:"transaction_id,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `wire_transfer`.
	Type WireTransferType `json:"type,required"`
	JSON WireTransferJSON
}

type WireTransferJSON struct {
	ID                      pjson.Metadata
	MessageToRecipient      pjson.Metadata
	Amount                  pjson.Metadata
	Currency                pjson.Metadata
	AccountNumber           pjson.Metadata
	BeneficiaryName         pjson.Metadata
	BeneficiaryAddressLine1 pjson.Metadata
	BeneficiaryAddressLine2 pjson.Metadata
	BeneficiaryAddressLine3 pjson.Metadata
	AccountID               pjson.Metadata
	ExternalAccountID       pjson.Metadata
	RoutingNumber           pjson.Metadata
	Approval                pjson.Metadata
	Cancellation            pjson.Metadata
	Reversal                pjson.Metadata
	CreatedAt               pjson.Metadata
	Network                 pjson.Metadata
	Status                  pjson.Metadata
	Submission              pjson.Metadata
	TransactionID           pjson.Metadata
	Type                    pjson.Metadata
	Raw                     []byte
	Extras                  map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into WireTransfer using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *WireTransfer) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type WireTransferCurrency string

const (
	WireTransferCurrencyCad WireTransferCurrency = "CAD"
	WireTransferCurrencyChf WireTransferCurrency = "CHF"
	WireTransferCurrencyEur WireTransferCurrency = "EUR"
	WireTransferCurrencyGbp WireTransferCurrency = "GBP"
	WireTransferCurrencyJpy WireTransferCurrency = "JPY"
	WireTransferCurrencyUsd WireTransferCurrency = "USD"
)

type WireTransferApproval struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was approved.
	ApprovedAt time.Time `json:"approved_at,required" format:"date-time"`
	// If the Transfer was approved by a user in the dashboard, the email address of
	// that user.
	ApprovedBy string `json:"approved_by,required,nullable"`
	JSON       WireTransferApprovalJSON
}

type WireTransferApprovalJSON struct {
	ApprovedAt pjson.Metadata
	ApprovedBy pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into WireTransferApproval using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *WireTransferApproval) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type WireTransferCancellation struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Transfer was canceled.
	CanceledAt time.Time `json:"canceled_at,required" format:"date-time"`
	// If the Transfer was canceled by a user in the dashboard, the email address of
	// that user.
	CanceledBy string `json:"canceled_by,required,nullable"`
	JSON       WireTransferCancellationJSON
}

type WireTransferCancellationJSON struct {
	CanceledAt pjson.Metadata
	CanceledBy pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into WireTransferCancellation
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *WireTransferCancellation) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type WireTransferReversal struct {
	// The amount that was reversed.
	Amount int64 `json:"amount,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the reversal was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The description on the reversal message from Fedwire.
	Description string `json:"description,required"`
	// The Fedwire cycle date for the wire reversal.
	InputCycleDate time.Time `json:"input_cycle_date,required" format:"date"`
	// The Fedwire sequence number.
	InputSequenceNumber string `json:"input_sequence_number,required"`
	// The Fedwire input source identifier.
	InputSource string `json:"input_source,required"`
	// The Fedwire transaction identifier.
	InputMessageAccountabilityData string `json:"input_message_accountability_data,required"`
	// The Fedwire transaction identifier for the wire transfer that was reversed.
	PreviousMessageInputMessageAccountabilityData string `json:"previous_message_input_message_accountability_data,required"`
	// The Fedwire cycle date for the wire transfer that was reversed.
	PreviousMessageInputCycleDate time.Time `json:"previous_message_input_cycle_date,required" format:"date"`
	// The Fedwire sequence number for the wire transfer that was reversed.
	PreviousMessageInputSequenceNumber string `json:"previous_message_input_sequence_number,required"`
	// The Fedwire input source identifier for the wire transfer that was reversed.
	PreviousMessageInputSource string `json:"previous_message_input_source,required"`
	// Information included in the wire reversal for the receiving financial
	// institution.
	ReceiverFinancialInstitutionInformation string `json:"receiver_financial_institution_information,required,nullable"`
	// Additional financial institution information included in the wire reversal.
	FinancialInstitutionToFinancialInstitutionInformation string `json:"financial_institution_to_financial_institution_information,required,nullable"`
	// The ID for the Transaction associated with the transfer reversal.
	TransactionID string `json:"transaction_id,required,nullable"`
	JSON          WireTransferReversalJSON
}

type WireTransferReversalJSON struct {
	Amount                                                pjson.Metadata
	CreatedAt                                             pjson.Metadata
	Description                                           pjson.Metadata
	InputCycleDate                                        pjson.Metadata
	InputSequenceNumber                                   pjson.Metadata
	InputSource                                           pjson.Metadata
	InputMessageAccountabilityData                        pjson.Metadata
	PreviousMessageInputMessageAccountabilityData         pjson.Metadata
	PreviousMessageInputCycleDate                         pjson.Metadata
	PreviousMessageInputSequenceNumber                    pjson.Metadata
	PreviousMessageInputSource                            pjson.Metadata
	ReceiverFinancialInstitutionInformation               pjson.Metadata
	FinancialInstitutionToFinancialInstitutionInformation pjson.Metadata
	TransactionID                                         pjson.Metadata
	Raw                                                   []byte
	Extras                                                map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into WireTransferReversal using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *WireTransferReversal) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type WireTransferNetwork string

const (
	WireTransferNetworkWire WireTransferNetwork = "wire"
)

type WireTransferStatus string

const (
	WireTransferStatusCanceled          WireTransferStatus = "canceled"
	WireTransferStatusRequiresAttention WireTransferStatus = "requires_attention"
	WireTransferStatusPendingApproval   WireTransferStatus = "pending_approval"
	WireTransferStatusRejected          WireTransferStatus = "rejected"
	WireTransferStatusReversed          WireTransferStatus = "reversed"
	WireTransferStatusComplete          WireTransferStatus = "complete"
	WireTransferStatusPendingCreating   WireTransferStatus = "pending_creating"
)

type WireTransferSubmission struct {
	// The accountability data for the submission.
	InputMessageAccountabilityData string `json:"input_message_accountability_data,required"`
	// When this wire transfer was submitted to Fedwire.
	SubmittedAt time.Time `json:"submitted_at,required" format:"date-time"`
	JSON        WireTransferSubmissionJSON
}

type WireTransferSubmissionJSON struct {
	InputMessageAccountabilityData pjson.Metadata
	SubmittedAt                    pjson.Metadata
	Raw                            []byte
	Extras                         map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into WireTransferSubmission using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *WireTransferSubmission) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type WireTransferType string

const (
	WireTransferTypeWireTransfer WireTransferType = "wire_transfer"
)

type WireTransferListResponse struct {
	// The contents of the list.
	Data []WireTransfer `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       WireTransferListResponseJSON
}

type WireTransferListResponseJSON struct {
	Data       pjson.Metadata
	NextCursor pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into WireTransferListResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *WireTransferListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

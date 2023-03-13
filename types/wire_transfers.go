package types

import (
	"fmt"
	"net/url"

	"github.com/increase/increase-go/core"
	"github.com/increase/increase-go/core/pjson"
	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/pagination"
)

type WireTransfer struct {
	// The wire transfer's identifier.
	ID *string `pjson:"id"`
	// The message that will show on the recipient's bank statement.
	MessageToRecipient *string `pjson:"message_to_recipient"`
	// The transfer amount in USD cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transfer's
	// currency. For wire transfers this is always equal to `usd`.
	Currency *WireTransferCurrency `pjson:"currency"`
	// The destination account number.
	AccountNumber *string `pjson:"account_number"`
	// The beneficiary's name.
	BeneficiaryName *string `pjson:"beneficiary_name"`
	// The beneficiary's address line 1.
	BeneficiaryAddressLine1 *string `pjson:"beneficiary_address_line1"`
	// The beneficiary's address line 2.
	BeneficiaryAddressLine2 *string `pjson:"beneficiary_address_line2"`
	// The beneficiary's address line 3.
	BeneficiaryAddressLine3 *string `pjson:"beneficiary_address_line3"`
	// The Account to which the transfer belongs.
	AccountID *string `pjson:"account_id"`
	// The identifier of the External Account the transfer was made to, if any.
	ExternalAccountID *string `pjson:"external_account_id"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber *string `pjson:"routing_number"`
	// If your account requires approvals for transfers and the transfer was approved,
	// this will contain details of the approval.
	Approval *WireTransferApproval `pjson:"approval"`
	// If your account requires approvals for transfers and the transfer was not
	// approved, this will contain details of the cancellation.
	Cancellation *WireTransferCancellation `pjson:"cancellation"`
	// If your transfer is reversed, this will contain details of the reversal.
	Reversal *WireTransferReversal `pjson:"reversal"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt *string `pjson:"created_at"`
	// The transfer's network.
	Network *WireTransferNetwork `pjson:"network"`
	// The lifecycle status of the transfer.
	Status *WireTransferStatus `pjson:"status"`
	// After the transfer is submitted to Fedwire, this will contain supplemental
	// details.
	Submission *WireTransferSubmission `pjson:"submission"`
	// If the transfer was created from a template, this will be the template's ID.
	TemplateID *string `pjson:"template_id"`
	// The ID for the transaction funding the transfer.
	TransactionID *string `pjson:"transaction_id"`
	// A constant representing the object's type. For this resource it will always be
	// `wire_transfer`.
	Type       *WireTransferType      `pjson:"type"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into WireTransfer using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *WireTransfer) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes WireTransfer into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *WireTransfer) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The wire transfer's identifier.
func (r WireTransfer) GetID() (ID string) {
	if r.ID != nil {
		ID = *r.ID
	}
	return
}

// The message that will show on the recipient's bank statement.
func (r WireTransfer) GetMessageToRecipient() (MessageToRecipient string) {
	if r.MessageToRecipient != nil {
		MessageToRecipient = *r.MessageToRecipient
	}
	return
}

// The transfer amount in USD cents.
func (r WireTransfer) GetAmount() (Amount int64) {
	if r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transfer's
// currency. For wire transfers this is always equal to `usd`.
func (r WireTransfer) GetCurrency() (Currency WireTransferCurrency) {
	if r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The destination account number.
func (r WireTransfer) GetAccountNumber() (AccountNumber string) {
	if r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

// The beneficiary's name.
func (r WireTransfer) GetBeneficiaryName() (BeneficiaryName string) {
	if r.BeneficiaryName != nil {
		BeneficiaryName = *r.BeneficiaryName
	}
	return
}

// The beneficiary's address line 1.
func (r WireTransfer) GetBeneficiaryAddressLine1() (BeneficiaryAddressLine1 string) {
	if r.BeneficiaryAddressLine1 != nil {
		BeneficiaryAddressLine1 = *r.BeneficiaryAddressLine1
	}
	return
}

// The beneficiary's address line 2.
func (r WireTransfer) GetBeneficiaryAddressLine2() (BeneficiaryAddressLine2 string) {
	if r.BeneficiaryAddressLine2 != nil {
		BeneficiaryAddressLine2 = *r.BeneficiaryAddressLine2
	}
	return
}

// The beneficiary's address line 3.
func (r WireTransfer) GetBeneficiaryAddressLine3() (BeneficiaryAddressLine3 string) {
	if r.BeneficiaryAddressLine3 != nil {
		BeneficiaryAddressLine3 = *r.BeneficiaryAddressLine3
	}
	return
}

// The Account to which the transfer belongs.
func (r WireTransfer) GetAccountID() (AccountID string) {
	if r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

// The identifier of the External Account the transfer was made to, if any.
func (r WireTransfer) GetExternalAccountID() (ExternalAccountID string) {
	if r.ExternalAccountID != nil {
		ExternalAccountID = *r.ExternalAccountID
	}
	return
}

// The American Bankers' Association (ABA) Routing Transit Number (RTN).
func (r WireTransfer) GetRoutingNumber() (RoutingNumber string) {
	if r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

// If your account requires approvals for transfers and the transfer was approved,
// this will contain details of the approval.
func (r WireTransfer) GetApproval() (Approval WireTransferApproval) {
	if r.Approval != nil {
		Approval = *r.Approval
	}
	return
}

// If your account requires approvals for transfers and the transfer was not
// approved, this will contain details of the cancellation.
func (r WireTransfer) GetCancellation() (Cancellation WireTransferCancellation) {
	if r.Cancellation != nil {
		Cancellation = *r.Cancellation
	}
	return
}

// If your transfer is reversed, this will contain details of the reversal.
func (r WireTransfer) GetReversal() (Reversal WireTransferReversal) {
	if r.Reversal != nil {
		Reversal = *r.Reversal
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the transfer was created.
func (r WireTransfer) GetCreatedAt() (CreatedAt string) {
	if r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The transfer's network.
func (r WireTransfer) GetNetwork() (Network WireTransferNetwork) {
	if r.Network != nil {
		Network = *r.Network
	}
	return
}

// The lifecycle status of the transfer.
func (r WireTransfer) GetStatus() (Status WireTransferStatus) {
	if r.Status != nil {
		Status = *r.Status
	}
	return
}

// After the transfer is submitted to Fedwire, this will contain supplemental
// details.
func (r WireTransfer) GetSubmission() (Submission WireTransferSubmission) {
	if r.Submission != nil {
		Submission = *r.Submission
	}
	return
}

// If the transfer was created from a template, this will be the template's ID.
func (r WireTransfer) GetTemplateID() (TemplateID string) {
	if r.TemplateID != nil {
		TemplateID = *r.TemplateID
	}
	return
}

// The ID for the transaction funding the transfer.
func (r WireTransfer) GetTransactionID() (TransactionID string) {
	if r.TransactionID != nil {
		TransactionID = *r.TransactionID
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `wire_transfer`.
func (r WireTransfer) GetType() (Type WireTransferType) {
	if r.Type != nil {
		Type = *r.Type
	}
	return
}

func (r WireTransfer) String() (result string) {
	return fmt.Sprintf("&WireTransfer{ID:%s MessageToRecipient:%s Amount:%s Currency:%s AccountNumber:%s BeneficiaryName:%s BeneficiaryAddressLine1:%s BeneficiaryAddressLine2:%s BeneficiaryAddressLine3:%s AccountID:%s ExternalAccountID:%s RoutingNumber:%s Approval:%s Cancellation:%s Reversal:%s CreatedAt:%s Network:%s Status:%s Submission:%s TemplateID:%s TransactionID:%s Type:%s}", core.FmtP(r.ID), core.FmtP(r.MessageToRecipient), core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.AccountNumber), core.FmtP(r.BeneficiaryName), core.FmtP(r.BeneficiaryAddressLine1), core.FmtP(r.BeneficiaryAddressLine2), core.FmtP(r.BeneficiaryAddressLine3), core.FmtP(r.AccountID), core.FmtP(r.ExternalAccountID), core.FmtP(r.RoutingNumber), r.Approval, r.Cancellation, r.Reversal, core.FmtP(r.CreatedAt), core.FmtP(r.Network), core.FmtP(r.Status), r.Submission, core.FmtP(r.TemplateID), core.FmtP(r.TransactionID), core.FmtP(r.Type))
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
	ApprovedAt *string                `pjson:"approved_at"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into WireTransferApproval using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *WireTransferApproval) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes WireTransferApproval into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *WireTransferApproval) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the transfer was approved.
func (r WireTransferApproval) GetApprovedAt() (ApprovedAt string) {
	if r.ApprovedAt != nil {
		ApprovedAt = *r.ApprovedAt
	}
	return
}

func (r WireTransferApproval) String() (result string) {
	return fmt.Sprintf("&WireTransferApproval{ApprovedAt:%s}", core.FmtP(r.ApprovedAt))
}

type WireTransferCancellation struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Transfer was canceled.
	CanceledAt *string                `pjson:"canceled_at"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into WireTransferCancellation
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *WireTransferCancellation) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes WireTransferCancellation into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *WireTransferCancellation) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the Transfer was canceled.
func (r WireTransferCancellation) GetCanceledAt() (CanceledAt string) {
	if r.CanceledAt != nil {
		CanceledAt = *r.CanceledAt
	}
	return
}

func (r WireTransferCancellation) String() (result string) {
	return fmt.Sprintf("&WireTransferCancellation{CanceledAt:%s}", core.FmtP(r.CanceledAt))
}

type WireTransferReversal struct {
	// The amount that was reversed.
	Amount *int64 `pjson:"amount"`
	// The description on the reversal message from Fedwire.
	Description *string `pjson:"description"`
	// The Fedwire cycle date for the wire reversal.
	InputCycleDate *string `pjson:"input_cycle_date"`
	// The Fedwire sequence number.
	InputSequenceNumber *string `pjson:"input_sequence_number"`
	// The Fedwire input source identifier.
	InputSource *string `pjson:"input_source"`
	// The Fedwire transaction identifier.
	InputMessageAccountabilityData *string `pjson:"input_message_accountability_data"`
	// The Fedwire transaction identifier for the wire transfer that was reversed.
	PreviousMessageInputMessageAccountabilityData *string `pjson:"previous_message_input_message_accountability_data"`
	// The Fedwire cycle date for the wire transfer that was reversed.
	PreviousMessageInputCycleDate *string `pjson:"previous_message_input_cycle_date"`
	// The Fedwire sequence number for the wire transfer that was reversed.
	PreviousMessageInputSequenceNumber *string `pjson:"previous_message_input_sequence_number"`
	// The Fedwire input source identifier for the wire transfer that was reversed.
	PreviousMessageInputSource *string `pjson:"previous_message_input_source"`
	// Information included in the wire reversal for the receiving financial
	// institution.
	ReceiverFinancialInstitutionInformation *string `pjson:"receiver_financial_institution_information"`
	// Additional financial institution information included in the wire reversal.
	FinancialInstitutionToFinancialInstitutionInformation *string                `pjson:"financial_institution_to_financial_institution_information"`
	jsonFields                                            map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into WireTransferReversal using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *WireTransferReversal) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes WireTransferReversal into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *WireTransferReversal) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount that was reversed.
func (r WireTransferReversal) GetAmount() (Amount int64) {
	if r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The description on the reversal message from Fedwire.
func (r WireTransferReversal) GetDescription() (Description string) {
	if r.Description != nil {
		Description = *r.Description
	}
	return
}

// The Fedwire cycle date for the wire reversal.
func (r WireTransferReversal) GetInputCycleDate() (InputCycleDate string) {
	if r.InputCycleDate != nil {
		InputCycleDate = *r.InputCycleDate
	}
	return
}

// The Fedwire sequence number.
func (r WireTransferReversal) GetInputSequenceNumber() (InputSequenceNumber string) {
	if r.InputSequenceNumber != nil {
		InputSequenceNumber = *r.InputSequenceNumber
	}
	return
}

// The Fedwire input source identifier.
func (r WireTransferReversal) GetInputSource() (InputSource string) {
	if r.InputSource != nil {
		InputSource = *r.InputSource
	}
	return
}

// The Fedwire transaction identifier.
func (r WireTransferReversal) GetInputMessageAccountabilityData() (InputMessageAccountabilityData string) {
	if r.InputMessageAccountabilityData != nil {
		InputMessageAccountabilityData = *r.InputMessageAccountabilityData
	}
	return
}

// The Fedwire transaction identifier for the wire transfer that was reversed.
func (r WireTransferReversal) GetPreviousMessageInputMessageAccountabilityData() (PreviousMessageInputMessageAccountabilityData string) {
	if r.PreviousMessageInputMessageAccountabilityData != nil {
		PreviousMessageInputMessageAccountabilityData = *r.PreviousMessageInputMessageAccountabilityData
	}
	return
}

// The Fedwire cycle date for the wire transfer that was reversed.
func (r WireTransferReversal) GetPreviousMessageInputCycleDate() (PreviousMessageInputCycleDate string) {
	if r.PreviousMessageInputCycleDate != nil {
		PreviousMessageInputCycleDate = *r.PreviousMessageInputCycleDate
	}
	return
}

// The Fedwire sequence number for the wire transfer that was reversed.
func (r WireTransferReversal) GetPreviousMessageInputSequenceNumber() (PreviousMessageInputSequenceNumber string) {
	if r.PreviousMessageInputSequenceNumber != nil {
		PreviousMessageInputSequenceNumber = *r.PreviousMessageInputSequenceNumber
	}
	return
}

// The Fedwire input source identifier for the wire transfer that was reversed.
func (r WireTransferReversal) GetPreviousMessageInputSource() (PreviousMessageInputSource string) {
	if r.PreviousMessageInputSource != nil {
		PreviousMessageInputSource = *r.PreviousMessageInputSource
	}
	return
}

// Information included in the wire reversal for the receiving financial
// institution.
func (r WireTransferReversal) GetReceiverFinancialInstitutionInformation() (ReceiverFinancialInstitutionInformation string) {
	if r.ReceiverFinancialInstitutionInformation != nil {
		ReceiverFinancialInstitutionInformation = *r.ReceiverFinancialInstitutionInformation
	}
	return
}

// Additional financial institution information included in the wire reversal.
func (r WireTransferReversal) GetFinancialInstitutionToFinancialInstitutionInformation() (FinancialInstitutionToFinancialInstitutionInformation string) {
	if r.FinancialInstitutionToFinancialInstitutionInformation != nil {
		FinancialInstitutionToFinancialInstitutionInformation = *r.FinancialInstitutionToFinancialInstitutionInformation
	}
	return
}

func (r WireTransferReversal) String() (result string) {
	return fmt.Sprintf("&WireTransferReversal{Amount:%s Description:%s InputCycleDate:%s InputSequenceNumber:%s InputSource:%s InputMessageAccountabilityData:%s PreviousMessageInputMessageAccountabilityData:%s PreviousMessageInputCycleDate:%s PreviousMessageInputSequenceNumber:%s PreviousMessageInputSource:%s ReceiverFinancialInstitutionInformation:%s FinancialInstitutionToFinancialInstitutionInformation:%s}", core.FmtP(r.Amount), core.FmtP(r.Description), core.FmtP(r.InputCycleDate), core.FmtP(r.InputSequenceNumber), core.FmtP(r.InputSource), core.FmtP(r.InputMessageAccountabilityData), core.FmtP(r.PreviousMessageInputMessageAccountabilityData), core.FmtP(r.PreviousMessageInputCycleDate), core.FmtP(r.PreviousMessageInputSequenceNumber), core.FmtP(r.PreviousMessageInputSource), core.FmtP(r.ReceiverFinancialInstitutionInformation), core.FmtP(r.FinancialInstitutionToFinancialInstitutionInformation))
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
	InputMessageAccountabilityData *string `pjson:"input_message_accountability_data"`
	// When this wire transfer was submitted to Fedwire.
	SubmittedAt *string                `pjson:"submitted_at"`
	jsonFields  map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into WireTransferSubmission using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *WireTransferSubmission) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes WireTransferSubmission into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *WireTransferSubmission) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The accountability data for the submission.
func (r WireTransferSubmission) GetInputMessageAccountabilityData() (InputMessageAccountabilityData string) {
	if r.InputMessageAccountabilityData != nil {
		InputMessageAccountabilityData = *r.InputMessageAccountabilityData
	}
	return
}

// When this wire transfer was submitted to Fedwire.
func (r WireTransferSubmission) GetSubmittedAt() (SubmittedAt string) {
	if r.SubmittedAt != nil {
		SubmittedAt = *r.SubmittedAt
	}
	return
}

func (r WireTransferSubmission) String() (result string) {
	return fmt.Sprintf("&WireTransferSubmission{InputMessageAccountabilityData:%s SubmittedAt:%s}", core.FmtP(r.InputMessageAccountabilityData), core.FmtP(r.SubmittedAt))
}

type WireTransferType string

const (
	WireTransferTypeWireTransfer WireTransferType = "wire_transfer"
)

type CreateAWireTransferParameters struct {
	// The identifier for the account that will send the transfer.
	AccountID *string `pjson:"account_id"`
	// The account number for the destination account.
	AccountNumber *string `pjson:"account_number"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
	// destination account.
	RoutingNumber *string `pjson:"routing_number"`
	// The ID of an External Account to initiate a transfer to. If this parameter is
	// provided, `account_number` and `routing_number` must be absent.
	ExternalAccountID *string `pjson:"external_account_id"`
	// The transfer amount in cents.
	Amount *int64 `pjson:"amount"`
	// The message that will show on the recipient's bank statement.
	MessageToRecipient *string `pjson:"message_to_recipient"`
	// The beneficiary's name.
	BeneficiaryName *string `pjson:"beneficiary_name"`
	// The beneficiary's address line 1.
	BeneficiaryAddressLine1 *string `pjson:"beneficiary_address_line1"`
	// The beneficiary's address line 2.
	BeneficiaryAddressLine2 *string `pjson:"beneficiary_address_line2"`
	// The beneficiary's address line 3.
	BeneficiaryAddressLine3 *string `pjson:"beneficiary_address_line3"`
	// Whether the transfer requires explicit approval via the dashboard or API.
	RequireApproval *bool                  `pjson:"require_approval"`
	jsonFields      map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into CreateAWireTransferParameters
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *CreateAWireTransferParameters) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CreateAWireTransferParameters into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *CreateAWireTransferParameters) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The identifier for the account that will send the transfer.
func (r CreateAWireTransferParameters) GetAccountID() (AccountID string) {
	if r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

// The account number for the destination account.
func (r CreateAWireTransferParameters) GetAccountNumber() (AccountNumber string) {
	if r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
// destination account.
func (r CreateAWireTransferParameters) GetRoutingNumber() (RoutingNumber string) {
	if r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

// The ID of an External Account to initiate a transfer to. If this parameter is
// provided, `account_number` and `routing_number` must be absent.
func (r CreateAWireTransferParameters) GetExternalAccountID() (ExternalAccountID string) {
	if r.ExternalAccountID != nil {
		ExternalAccountID = *r.ExternalAccountID
	}
	return
}

// The transfer amount in cents.
func (r CreateAWireTransferParameters) GetAmount() (Amount int64) {
	if r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The message that will show on the recipient's bank statement.
func (r CreateAWireTransferParameters) GetMessageToRecipient() (MessageToRecipient string) {
	if r.MessageToRecipient != nil {
		MessageToRecipient = *r.MessageToRecipient
	}
	return
}

// The beneficiary's name.
func (r CreateAWireTransferParameters) GetBeneficiaryName() (BeneficiaryName string) {
	if r.BeneficiaryName != nil {
		BeneficiaryName = *r.BeneficiaryName
	}
	return
}

// The beneficiary's address line 1.
func (r CreateAWireTransferParameters) GetBeneficiaryAddressLine1() (BeneficiaryAddressLine1 string) {
	if r.BeneficiaryAddressLine1 != nil {
		BeneficiaryAddressLine1 = *r.BeneficiaryAddressLine1
	}
	return
}

// The beneficiary's address line 2.
func (r CreateAWireTransferParameters) GetBeneficiaryAddressLine2() (BeneficiaryAddressLine2 string) {
	if r.BeneficiaryAddressLine2 != nil {
		BeneficiaryAddressLine2 = *r.BeneficiaryAddressLine2
	}
	return
}

// The beneficiary's address line 3.
func (r CreateAWireTransferParameters) GetBeneficiaryAddressLine3() (BeneficiaryAddressLine3 string) {
	if r.BeneficiaryAddressLine3 != nil {
		BeneficiaryAddressLine3 = *r.BeneficiaryAddressLine3
	}
	return
}

// Whether the transfer requires explicit approval via the dashboard or API.
func (r CreateAWireTransferParameters) GetRequireApproval() (RequireApproval bool) {
	if r.RequireApproval != nil {
		RequireApproval = *r.RequireApproval
	}
	return
}

func (r CreateAWireTransferParameters) String() (result string) {
	return fmt.Sprintf("&CreateAWireTransferParameters{AccountID:%s AccountNumber:%s RoutingNumber:%s ExternalAccountID:%s Amount:%s MessageToRecipient:%s BeneficiaryName:%s BeneficiaryAddressLine1:%s BeneficiaryAddressLine2:%s BeneficiaryAddressLine3:%s RequireApproval:%s}", core.FmtP(r.AccountID), core.FmtP(r.AccountNumber), core.FmtP(r.RoutingNumber), core.FmtP(r.ExternalAccountID), core.FmtP(r.Amount), core.FmtP(r.MessageToRecipient), core.FmtP(r.BeneficiaryName), core.FmtP(r.BeneficiaryAddressLine1), core.FmtP(r.BeneficiaryAddressLine2), core.FmtP(r.BeneficiaryAddressLine3), core.FmtP(r.RequireApproval))
}

type WireTransferListParams struct {
	// Return the page of entries after this one.
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int64 `query:"limit"`
	// Filter Wire Transfers to those belonging to the specified Account.
	AccountID *string `query:"account_id"`
	// Filter Wire Transfers to those made to the specified External Account.
	ExternalAccountID *string                           `query:"external_account_id"`
	CreatedAt         *WireTransfersListParamsCreatedAt `query:"created_at"`
	jsonFields        map[string]interface{}            `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into WireTransferListParams using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *WireTransferListParams) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes WireTransferListParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *WireTransferListParams) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes WireTransferListParams into a url.Values of the query
// parameters associated with this value
func (r *WireTransferListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// Return the page of entries after this one.
func (r WireTransferListParams) GetCursor() (Cursor string) {
	if r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r WireTransferListParams) GetLimit() (Limit int64) {
	if r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// Filter Wire Transfers to those belonging to the specified Account.
func (r WireTransferListParams) GetAccountID() (AccountID string) {
	if r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

// Filter Wire Transfers to those made to the specified External Account.
func (r WireTransferListParams) GetExternalAccountID() (ExternalAccountID string) {
	if r.ExternalAccountID != nil {
		ExternalAccountID = *r.ExternalAccountID
	}
	return
}

func (r WireTransferListParams) GetCreatedAt() (CreatedAt WireTransfersListParamsCreatedAt) {
	if r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

func (r WireTransferListParams) String() (result string) {
	return fmt.Sprintf("&WireTransferListParams{Cursor:%s Limit:%s AccountID:%s ExternalAccountID:%s CreatedAt:%s}", core.FmtP(r.Cursor), core.FmtP(r.Limit), core.FmtP(r.AccountID), core.FmtP(r.ExternalAccountID), r.CreatedAt)
}

type WireTransfersListParamsCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After *string `pjson:"after"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before *string `pjson:"before"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter *string `pjson:"on_or_after"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore *string                `pjson:"on_or_before"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// WireTransfersListParamsCreatedAt using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *WireTransfersListParamsCreatedAt) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes WireTransfersListParamsCreatedAt into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *WireTransfersListParamsCreatedAt) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes WireTransfersListParamsCreatedAt into a url.Values of the
// query parameters associated with this value
func (r *WireTransfersListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r WireTransfersListParamsCreatedAt) GetAfter() (After string) {
	if r.After != nil {
		After = *r.After
	}
	return
}

// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r WireTransfersListParamsCreatedAt) GetBefore() (Before string) {
	if r.Before != nil {
		Before = *r.Before
	}
	return
}

// Return results on or after this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r WireTransfersListParamsCreatedAt) GetOnOrAfter() (OnOrAfter string) {
	if r.OnOrAfter != nil {
		OnOrAfter = *r.OnOrAfter
	}
	return
}

// Return results on or before this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r WireTransfersListParamsCreatedAt) GetOnOrBefore() (OnOrBefore string) {
	if r.OnOrBefore != nil {
		OnOrBefore = *r.OnOrBefore
	}
	return
}

func (r WireTransfersListParamsCreatedAt) String() (result string) {
	return fmt.Sprintf("&WireTransfersListParamsCreatedAt{After:%s Before:%s OnOrAfter:%s OnOrBefore:%s}", core.FmtP(r.After), core.FmtP(r.Before), core.FmtP(r.OnOrAfter), core.FmtP(r.OnOrBefore))
}

type WireTransferList struct {
	// The contents of the list.
	Data *[]WireTransfer `pjson:"data"`
	// A pointer to a place in the list.
	NextCursor *string                `pjson:"next_cursor"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into WireTransferList using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *WireTransferList) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes WireTransferList into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *WireTransferList) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes WireTransferList into a url.Values of the query parameters
// associated with this value
func (r *WireTransferList) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// The contents of the list.
func (r WireTransferList) GetData() (Data []WireTransfer) {
	if r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r WireTransferList) GetNextCursor() (NextCursor string) {
	if r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

func (r WireTransferList) String() (result string) {
	return fmt.Sprintf("&WireTransferList{Data:%s NextCursor:%s}", core.Fmt(r.Data), core.FmtP(r.NextCursor))
}

type WireTransfersPage struct {
	*pagination.Page[WireTransfer]
}

func (r *WireTransfersPage) WireTransfer() *WireTransfer {
	return r.Current()
}

func (r *WireTransfersPage) NextPage() (*WireTransfersPage, error) {
	if page, err := r.Page.NextPage(); err != nil {
		return nil, err
	} else {
		return &WireTransfersPage{page}, nil
	}
}

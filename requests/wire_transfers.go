package requests

import (
	"fmt"
	"net/url"
	"time"

	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/fields"
)

type WireTransfer struct {
	// The wire transfer's identifier.
	ID fields.Field[string] `json:"id,required"`
	// The message that will show on the recipient's bank statement.
	MessageToRecipient fields.Field[string] `json:"message_to_recipient,required,nullable"`
	// The transfer amount in USD cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transfer's
	// currency. For wire transfers this is always equal to `usd`.
	Currency fields.Field[WireTransferCurrency] `json:"currency,required"`
	// The destination account number.
	AccountNumber fields.Field[string] `json:"account_number,required"`
	// The beneficiary's name.
	BeneficiaryName fields.Field[string] `json:"beneficiary_name,required,nullable"`
	// The beneficiary's address line 1.
	BeneficiaryAddressLine1 fields.Field[string] `json:"beneficiary_address_line1,required,nullable"`
	// The beneficiary's address line 2.
	BeneficiaryAddressLine2 fields.Field[string] `json:"beneficiary_address_line2,required,nullable"`
	// The beneficiary's address line 3.
	BeneficiaryAddressLine3 fields.Field[string] `json:"beneficiary_address_line3,required,nullable"`
	// The Account to which the transfer belongs.
	AccountID fields.Field[string] `json:"account_id,required"`
	// The identifier of the External Account the transfer was made to, if any.
	ExternalAccountID fields.Field[string] `json:"external_account_id,required,nullable"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber fields.Field[string] `json:"routing_number,required"`
	// If your account requires approvals for transfers and the transfer was approved,
	// this will contain details of the approval.
	Approval fields.Field[WireTransferApproval] `json:"approval,required,nullable"`
	// If your account requires approvals for transfers and the transfer was not
	// approved, this will contain details of the cancellation.
	Cancellation fields.Field[WireTransferCancellation] `json:"cancellation,required,nullable"`
	// If your transfer is reversed, this will contain details of the reversal.
	Reversal fields.Field[WireTransferReversal] `json:"reversal,required,nullable"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt fields.Field[time.Time] `json:"created_at,required" format:"date-time"`
	// The transfer's network.
	Network fields.Field[WireTransferNetwork] `json:"network,required"`
	// The lifecycle status of the transfer.
	Status fields.Field[WireTransferStatus] `json:"status,required"`
	// After the transfer is submitted to Fedwire, this will contain supplemental
	// details.
	Submission fields.Field[WireTransferSubmission] `json:"submission,required,nullable"`
	// If the transfer was created from a template, this will be the template's ID.
	TemplateID fields.Field[string] `json:"template_id,required,nullable"`
	// The ID for the transaction funding the transfer.
	TransactionID fields.Field[string] `json:"transaction_id,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `wire_transfer`.
	Type fields.Field[WireTransferType] `json:"type,required"`
}

// MarshalJSON serializes WireTransfer into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *WireTransfer) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WireTransfer) String() (result string) {
	return fmt.Sprintf("&WireTransfer{ID:%s MessageToRecipient:%s Amount:%s Currency:%s AccountNumber:%s BeneficiaryName:%s BeneficiaryAddressLine1:%s BeneficiaryAddressLine2:%s BeneficiaryAddressLine3:%s AccountID:%s ExternalAccountID:%s RoutingNumber:%s Approval:%s Cancellation:%s Reversal:%s CreatedAt:%s Network:%s Status:%s Submission:%s TemplateID:%s TransactionID:%s Type:%s}", r.ID, r.MessageToRecipient, r.Amount, r.Currency, r.AccountNumber, r.BeneficiaryName, r.BeneficiaryAddressLine1, r.BeneficiaryAddressLine2, r.BeneficiaryAddressLine3, r.AccountID, r.ExternalAccountID, r.RoutingNumber, r.Approval, r.Cancellation, r.Reversal, r.CreatedAt, r.Network, r.Status, r.Submission, r.TemplateID, r.TransactionID, r.Type)
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
	ApprovedAt fields.Field[time.Time] `json:"approved_at,required" format:"date-time"`
}

// MarshalJSON serializes WireTransferApproval into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *WireTransferApproval) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WireTransferApproval) String() (result string) {
	return fmt.Sprintf("&WireTransferApproval{ApprovedAt:%s}", r.ApprovedAt)
}

type WireTransferCancellation struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Transfer was canceled.
	CanceledAt fields.Field[time.Time] `json:"canceled_at,required" format:"date-time"`
}

// MarshalJSON serializes WireTransferCancellation into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *WireTransferCancellation) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WireTransferCancellation) String() (result string) {
	return fmt.Sprintf("&WireTransferCancellation{CanceledAt:%s}", r.CanceledAt)
}

type WireTransferReversal struct {
	// The amount that was reversed.
	Amount fields.Field[int64] `json:"amount,required"`
	// The description on the reversal message from Fedwire.
	Description fields.Field[string] `json:"description,required"`
	// The Fedwire cycle date for the wire reversal.
	InputCycleDate fields.Field[time.Time] `json:"input_cycle_date,required" format:"date"`
	// The Fedwire sequence number.
	InputSequenceNumber fields.Field[string] `json:"input_sequence_number,required"`
	// The Fedwire input source identifier.
	InputSource fields.Field[string] `json:"input_source,required"`
	// The Fedwire transaction identifier.
	InputMessageAccountabilityData fields.Field[string] `json:"input_message_accountability_data,required"`
	// The Fedwire transaction identifier for the wire transfer that was reversed.
	PreviousMessageInputMessageAccountabilityData fields.Field[string] `json:"previous_message_input_message_accountability_data,required"`
	// The Fedwire cycle date for the wire transfer that was reversed.
	PreviousMessageInputCycleDate fields.Field[time.Time] `json:"previous_message_input_cycle_date,required" format:"date"`
	// The Fedwire sequence number for the wire transfer that was reversed.
	PreviousMessageInputSequenceNumber fields.Field[string] `json:"previous_message_input_sequence_number,required"`
	// The Fedwire input source identifier for the wire transfer that was reversed.
	PreviousMessageInputSource fields.Field[string] `json:"previous_message_input_source,required"`
	// Information included in the wire reversal for the receiving financial
	// institution.
	ReceiverFinancialInstitutionInformation fields.Field[string] `json:"receiver_financial_institution_information,required,nullable"`
	// Additional financial institution information included in the wire reversal.
	FinancialInstitutionToFinancialInstitutionInformation fields.Field[string] `json:"financial_institution_to_financial_institution_information,required,nullable"`
}

// MarshalJSON serializes WireTransferReversal into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *WireTransferReversal) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WireTransferReversal) String() (result string) {
	return fmt.Sprintf("&WireTransferReversal{Amount:%s Description:%s InputCycleDate:%s InputSequenceNumber:%s InputSource:%s InputMessageAccountabilityData:%s PreviousMessageInputMessageAccountabilityData:%s PreviousMessageInputCycleDate:%s PreviousMessageInputSequenceNumber:%s PreviousMessageInputSource:%s ReceiverFinancialInstitutionInformation:%s FinancialInstitutionToFinancialInstitutionInformation:%s}", r.Amount, r.Description, r.InputCycleDate, r.InputSequenceNumber, r.InputSource, r.InputMessageAccountabilityData, r.PreviousMessageInputMessageAccountabilityData, r.PreviousMessageInputCycleDate, r.PreviousMessageInputSequenceNumber, r.PreviousMessageInputSource, r.ReceiverFinancialInstitutionInformation, r.FinancialInstitutionToFinancialInstitutionInformation)
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
	InputMessageAccountabilityData fields.Field[string] `json:"input_message_accountability_data,required"`
	// When this wire transfer was submitted to Fedwire.
	SubmittedAt fields.Field[time.Time] `json:"submitted_at,required" format:"date-time"`
}

// MarshalJSON serializes WireTransferSubmission into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *WireTransferSubmission) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WireTransferSubmission) String() (result string) {
	return fmt.Sprintf("&WireTransferSubmission{InputMessageAccountabilityData:%s SubmittedAt:%s}", r.InputMessageAccountabilityData, r.SubmittedAt)
}

type WireTransferType string

const (
	WireTransferTypeWireTransfer WireTransferType = "wire_transfer"
)

type CreateAWireTransferParameters struct {
	// The identifier for the account that will send the transfer.
	AccountID fields.Field[string] `json:"account_id,required"`
	// The account number for the destination account.
	AccountNumber fields.Field[string] `json:"account_number"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
	// destination account.
	RoutingNumber fields.Field[string] `json:"routing_number"`
	// The ID of an External Account to initiate a transfer to. If this parameter is
	// provided, `account_number` and `routing_number` must be absent.
	ExternalAccountID fields.Field[string] `json:"external_account_id"`
	// The transfer amount in cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The message that will show on the recipient's bank statement.
	MessageToRecipient fields.Field[string] `json:"message_to_recipient,required"`
	// The beneficiary's name.
	BeneficiaryName fields.Field[string] `json:"beneficiary_name,required"`
	// The beneficiary's address line 1.
	BeneficiaryAddressLine1 fields.Field[string] `json:"beneficiary_address_line1"`
	// The beneficiary's address line 2.
	BeneficiaryAddressLine2 fields.Field[string] `json:"beneficiary_address_line2"`
	// The beneficiary's address line 3.
	BeneficiaryAddressLine3 fields.Field[string] `json:"beneficiary_address_line3"`
	// Whether the transfer requires explicit approval via the dashboard or API.
	RequireApproval fields.Field[bool] `json:"require_approval"`
}

// MarshalJSON serializes CreateAWireTransferParameters into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *CreateAWireTransferParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CreateAWireTransferParameters) String() (result string) {
	return fmt.Sprintf("&CreateAWireTransferParameters{AccountID:%s AccountNumber:%s RoutingNumber:%s ExternalAccountID:%s Amount:%s MessageToRecipient:%s BeneficiaryName:%s BeneficiaryAddressLine1:%s BeneficiaryAddressLine2:%s BeneficiaryAddressLine3:%s RequireApproval:%s}", r.AccountID, r.AccountNumber, r.RoutingNumber, r.ExternalAccountID, r.Amount, r.MessageToRecipient, r.BeneficiaryName, r.BeneficiaryAddressLine1, r.BeneficiaryAddressLine2, r.BeneficiaryAddressLine3, r.RequireApproval)
}

type WireTransferListParams struct {
	// Return the page of entries after this one.
	Cursor fields.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit fields.Field[int64] `query:"limit"`
	// Filter Wire Transfers to those belonging to the specified Account.
	AccountID fields.Field[string] `query:"account_id"`
	// Filter Wire Transfers to those made to the specified External Account.
	ExternalAccountID fields.Field[string]                          `query:"external_account_id"`
	CreatedAt         fields.Field[WireTransferListParamsCreatedAt] `query:"created_at"`
}

// URLQuery serializes WireTransferListParams into a url.Values of the query
// parameters associated with this value
func (r *WireTransferListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r WireTransferListParams) String() (result string) {
	return fmt.Sprintf("&WireTransferListParams{Cursor:%s Limit:%s AccountID:%s ExternalAccountID:%s CreatedAt:%s}", r.Cursor, r.Limit, r.AccountID, r.ExternalAccountID, r.CreatedAt)
}

type WireTransferListParamsCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After fields.Field[time.Time] `query:"after" format:"date-time"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before fields.Field[time.Time] `query:"before" format:"date-time"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter fields.Field[time.Time] `query:"on_or_after" format:"date-time"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore fields.Field[time.Time] `query:"on_or_before" format:"date-time"`
}

// URLQuery serializes WireTransferListParamsCreatedAt into a url.Values of the
// query parameters associated with this value
func (r *WireTransferListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r WireTransferListParamsCreatedAt) String() (result string) {
	return fmt.Sprintf("&WireTransferListParamsCreatedAt{After:%s Before:%s OnOrAfter:%s OnOrBefore:%s}", r.After, r.Before, r.OnOrAfter, r.OnOrBefore)
}

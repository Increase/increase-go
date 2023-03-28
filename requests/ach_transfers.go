package requests

import (
	"fmt"
	"net/url"
	"time"

	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/fields"
)

type ACHTransfer struct {
	// The Account to which the transfer belongs.
	AccountID fields.Field[string] `json:"account_id,required"`
	// The destination account number.
	AccountNumber fields.Field[string] `json:"account_number,required"`
	// Additional information that will be sent to the recipient.
	Addendum fields.Field[string] `json:"addendum,required,nullable"`
	// The transfer amount in USD cents. A positive amount indicates a credit transfer
	// pushing funds to the receiving account. A negative amount indicates a debit
	// transfer pulling funds from the receiving account.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transfer's
	// currency. For ACH transfers this is always equal to `usd`.
	Currency fields.Field[ACHTransferCurrency] `json:"currency,required"`
	// If your account requires approvals for transfers and the transfer was approved,
	// this will contain details of the approval.
	Approval fields.Field[ACHTransferApproval] `json:"approval,required,nullable"`
	// If your account requires approvals for transfers and the transfer was not
	// approved, this will contain details of the cancellation.
	Cancellation fields.Field[ACHTransferCancellation] `json:"cancellation,required,nullable"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt fields.Field[time.Time] `json:"created_at,required" format:"date-time"`
	// The identifier of the External Account the transfer was made to, if any.
	ExternalAccountID fields.Field[string] `json:"external_account_id,required,nullable"`
	// The ACH transfer's identifier.
	ID fields.Field[string] `json:"id,required"`
	// The transfer's network.
	Network fields.Field[ACHTransferNetwork] `json:"network,required"`
	// If the receiving bank accepts the transfer but notifies that future transfers
	// should use different details, this will contain those details.
	NotificationOfChange fields.Field[ACHTransferNotificationOfChange] `json:"notification_of_change,required,nullable"`
	// If your transfer is returned, this will contain details of the return.
	Return fields.Field[ACHTransferReturn] `json:"return,required,nullable"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber fields.Field[string] `json:"routing_number,required"`
	// The descriptor that will show on the recipient's bank statement.
	StatementDescriptor fields.Field[string] `json:"statement_descriptor,required"`
	// The lifecycle status of the transfer.
	Status fields.Field[ACHTransferStatus] `json:"status,required"`
	// After the transfer is submitted to FedACH, this will contain supplemental
	// details.
	Submission fields.Field[ACHTransferSubmission] `json:"submission,required,nullable"`
	// If the transfer was created from a template, this will be the template's ID.
	TemplateID fields.Field[string] `json:"template_id,required,nullable"`
	// The ID for the transaction funding the transfer.
	TransactionID fields.Field[string] `json:"transaction_id,required,nullable"`
	// The description of the date of the transfer.
	CompanyDescriptiveDate fields.Field[string] `json:"company_descriptive_date,required,nullable"`
	// The data you chose to associate with the transfer.
	CompanyDiscretionaryData fields.Field[string] `json:"company_discretionary_data,required,nullable"`
	// The description of the transfer you set to be shown to the recipient.
	CompanyEntryDescription fields.Field[string] `json:"company_entry_description,required,nullable"`
	// The name by which the recipient knows you.
	CompanyName fields.Field[string] `json:"company_name,required,nullable"`
	// The type of the account to which the transfer will be sent.
	Funding fields.Field[ACHTransferFunding] `json:"funding,required"`
	// Your identifer for the transfer recipient.
	IndividualID fields.Field[string] `json:"individual_id,required,nullable"`
	// The name of the transfer recipient. This value is information and not verified
	// by the recipient's bank.
	IndividualName fields.Field[string] `json:"individual_name,required,nullable"`
	// The transfer effective date in
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	EffectiveDate fields.Field[time.Time] `json:"effective_date,required,nullable" format:"date"`
	// The Standard Entry Class (SEC) code to use for the transfer.
	StandardEntryClassCode fields.Field[ACHTransferStandardEntryClassCode] `json:"standard_entry_class_code,required"`
	// A constant representing the object's type. For this resource it will always be
	// `ach_transfer`.
	Type fields.Field[ACHTransferType] `json:"type,required"`
}

// MarshalJSON serializes ACHTransfer into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *ACHTransfer) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransfer) String() (result string) {
	return fmt.Sprintf("&ACHTransfer{AccountID:%s AccountNumber:%s Addendum:%s Amount:%s Currency:%s Approval:%s Cancellation:%s CreatedAt:%s ExternalAccountID:%s ID:%s Network:%s NotificationOfChange:%s Return:%s RoutingNumber:%s StatementDescriptor:%s Status:%s Submission:%s TemplateID:%s TransactionID:%s CompanyDescriptiveDate:%s CompanyDiscretionaryData:%s CompanyEntryDescription:%s CompanyName:%s Funding:%s IndividualID:%s IndividualName:%s EffectiveDate:%s StandardEntryClassCode:%s Type:%s}", r.AccountID, r.AccountNumber, r.Addendum, r.Amount, r.Currency, r.Approval, r.Cancellation, r.CreatedAt, r.ExternalAccountID, r.ID, r.Network, r.NotificationOfChange, r.Return, r.RoutingNumber, r.StatementDescriptor, r.Status, r.Submission, r.TemplateID, r.TransactionID, r.CompanyDescriptiveDate, r.CompanyDiscretionaryData, r.CompanyEntryDescription, r.CompanyName, r.Funding, r.IndividualID, r.IndividualName, r.EffectiveDate, r.StandardEntryClassCode, r.Type)
}

type ACHTransferCurrency string

const (
	ACHTransferCurrencyCad ACHTransferCurrency = "CAD"
	ACHTransferCurrencyChf ACHTransferCurrency = "CHF"
	ACHTransferCurrencyEur ACHTransferCurrency = "EUR"
	ACHTransferCurrencyGbp ACHTransferCurrency = "GBP"
	ACHTransferCurrencyJpy ACHTransferCurrency = "JPY"
	ACHTransferCurrencyUsd ACHTransferCurrency = "USD"
)

type ACHTransferApproval struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was approved.
	ApprovedAt fields.Field[time.Time] `json:"approved_at,required" format:"date-time"`
}

// MarshalJSON serializes ACHTransferApproval into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *ACHTransferApproval) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferApproval) String() (result string) {
	return fmt.Sprintf("&ACHTransferApproval{ApprovedAt:%s}", r.ApprovedAt)
}

type ACHTransferCancellation struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Transfer was canceled.
	CanceledAt fields.Field[time.Time] `json:"canceled_at,required" format:"date-time"`
}

// MarshalJSON serializes ACHTransferCancellation into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *ACHTransferCancellation) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferCancellation) String() (result string) {
	return fmt.Sprintf("&ACHTransferCancellation{CanceledAt:%s}", r.CanceledAt)
}

type ACHTransferNetwork string

const (
	ACHTransferNetworkACH ACHTransferNetwork = "ach"
)

type ACHTransferNotificationOfChange struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the notification occurred.
	CreatedAt fields.Field[time.Time] `json:"created_at,required" format:"date-time"`
	// The type of change that occurred.
	ChangeCode fields.Field[string] `json:"change_code,required"`
	// The corrected data.
	CorrectedData fields.Field[string] `json:"corrected_data,required"`
}

// MarshalJSON serializes ACHTransferNotificationOfChange into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *ACHTransferNotificationOfChange) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferNotificationOfChange) String() (result string) {
	return fmt.Sprintf("&ACHTransferNotificationOfChange{CreatedAt:%s ChangeCode:%s CorrectedData:%s}", r.CreatedAt, r.ChangeCode, r.CorrectedData)
}

type ACHTransferReturn struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt fields.Field[time.Time] `json:"created_at,required" format:"date-time"`
	// Why the ACH Transfer was returned.
	ReturnReasonCode fields.Field[ACHTransferReturnReturnReasonCode] `json:"return_reason_code,required"`
	// The identifier of the ACH Transfer associated with this return.
	TransferID fields.Field[string] `json:"transfer_id,required"`
	// The identifier of the Tranasaction associated with this return.
	TransactionID fields.Field[string] `json:"transaction_id,required"`
}

// MarshalJSON serializes ACHTransferReturn into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *ACHTransferReturn) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferReturn) String() (result string) {
	return fmt.Sprintf("&ACHTransferReturn{CreatedAt:%s ReturnReasonCode:%s TransferID:%s TransactionID:%s}", r.CreatedAt, r.ReturnReasonCode, r.TransferID, r.TransactionID)
}

type ACHTransferReturnReturnReasonCode string

const (
	ACHTransferReturnReturnReasonCodeInsufficientFund                                          ACHTransferReturnReturnReasonCode = "insufficient_fund"
	ACHTransferReturnReturnReasonCodeNoAccount                                                 ACHTransferReturnReturnReasonCode = "no_account"
	ACHTransferReturnReturnReasonCodeAccountClosed                                             ACHTransferReturnReturnReasonCode = "account_closed"
	ACHTransferReturnReturnReasonCodeInvalidAccountNumberStructure                             ACHTransferReturnReturnReasonCode = "invalid_account_number_structure"
	ACHTransferReturnReturnReasonCodeAccountFrozenEntryReturnedPerOfacInstruction              ACHTransferReturnReturnReasonCode = "account_frozen_entry_returned_per_ofac_instruction"
	ACHTransferReturnReturnReasonCodeCreditEntryRefusedByReceiver                              ACHTransferReturnReturnReasonCode = "credit_entry_refused_by_receiver"
	ACHTransferReturnReturnReasonCodeUnauthorizedDebitToConsumerAccountUsingCorporateSecCode   ACHTransferReturnReturnReasonCode = "unauthorized_debit_to_consumer_account_using_corporate_sec_code"
	ACHTransferReturnReturnReasonCodeCorporateCustomerAdvisedNotAuthorized                     ACHTransferReturnReturnReasonCode = "corporate_customer_advised_not_authorized"
	ACHTransferReturnReturnReasonCodePaymentStopped                                            ACHTransferReturnReturnReasonCode = "payment_stopped"
	ACHTransferReturnReturnReasonCodeNonTransactionAccount                                     ACHTransferReturnReturnReasonCode = "non_transaction_account"
	ACHTransferReturnReturnReasonCodeUncollectedFunds                                          ACHTransferReturnReturnReasonCode = "uncollected_funds"
	ACHTransferReturnReturnReasonCodeRoutingNumberCheckDigitError                              ACHTransferReturnReturnReasonCode = "routing_number_check_digit_error"
	ACHTransferReturnReturnReasonCodeCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete ACHTransferReturnReturnReasonCode = "customer_advised_unauthorized_improper_ineligible_or_incomplete"
	ACHTransferReturnReturnReasonCodeAmountFieldError                                          ACHTransferReturnReturnReasonCode = "amount_field_error"
	ACHTransferReturnReturnReasonCodeAuthorizationRevokedByCustomer                            ACHTransferReturnReturnReasonCode = "authorization_revoked_by_customer"
	ACHTransferReturnReturnReasonCodeInvalidACHRoutingNumber                                   ACHTransferReturnReturnReasonCode = "invalid_ach_routing_number"
	ACHTransferReturnReturnReasonCodeFileRecordEditCriteria                                    ACHTransferReturnReturnReasonCode = "file_record_edit_criteria"
	ACHTransferReturnReturnReasonCodeEnrInvalidIndividualName                                  ACHTransferReturnReturnReasonCode = "enr_invalid_individual_name"
	ACHTransferReturnReturnReasonCodeReturnedPerOdfiRequest                                    ACHTransferReturnReturnReasonCode = "returned_per_odfi_request"
	ACHTransferReturnReturnReasonCodeAddendaError                                              ACHTransferReturnReturnReasonCode = "addenda_error"
	ACHTransferReturnReturnReasonCodeLimitedParticipationDfi                                   ACHTransferReturnReturnReasonCode = "limited_participation_dfi"
	ACHTransferReturnReturnReasonCodeIncorrectlyCodedOutboundInternationalPayment              ACHTransferReturnReturnReasonCode = "incorrectly_coded_outbound_international_payment"
	ACHTransferReturnReturnReasonCodeOther                                                     ACHTransferReturnReturnReasonCode = "other"
)

type ACHTransferStatus string

const (
	ACHTransferStatusPendingApproval   ACHTransferStatus = "pending_approval"
	ACHTransferStatusCanceled          ACHTransferStatus = "canceled"
	ACHTransferStatusPendingSubmission ACHTransferStatus = "pending_submission"
	ACHTransferStatusSubmitted         ACHTransferStatus = "submitted"
	ACHTransferStatusReturned          ACHTransferStatus = "returned"
	ACHTransferStatusRequiresAttention ACHTransferStatus = "requires_attention"
	ACHTransferStatusRejected          ACHTransferStatus = "rejected"
)

type ACHTransferSubmission struct {
	// The trace number for the submission.
	TraceNumber fields.Field[string] `json:"trace_number,required"`
	// When the ACH transfer was sent to FedACH.
	SubmittedAt fields.Field[time.Time] `json:"submitted_at,required" format:"date-time"`
}

// MarshalJSON serializes ACHTransferSubmission into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *ACHTransferSubmission) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHTransferSubmission) String() (result string) {
	return fmt.Sprintf("&ACHTransferSubmission{TraceNumber:%s SubmittedAt:%s}", r.TraceNumber, r.SubmittedAt)
}

type ACHTransferFunding string

const (
	ACHTransferFundingChecking ACHTransferFunding = "checking"
	ACHTransferFundingSavings  ACHTransferFunding = "savings"
)

type ACHTransferStandardEntryClassCode string

const (
	ACHTransferStandardEntryClassCodeCorporateCreditOrDebit        ACHTransferStandardEntryClassCode = "corporate_credit_or_debit"
	ACHTransferStandardEntryClassCodePrearrangedPaymentsAndDeposit ACHTransferStandardEntryClassCode = "prearranged_payments_and_deposit"
	ACHTransferStandardEntryClassCodeInternetInitiated             ACHTransferStandardEntryClassCode = "internet_initiated"
)

type ACHTransferType string

const (
	ACHTransferTypeACHTransfer ACHTransferType = "ach_transfer"
)

type CreateAnACHTransferParameters struct {
	// The Increase identifier for the account that will send the transfer.
	AccountID fields.Field[string] `json:"account_id,required"`
	// The account number for the destination account.
	AccountNumber fields.Field[string] `json:"account_number"`
	// Additional information that will be sent to the recipient. This is included in
	// the transfer data sent to the receiving bank.
	Addendum fields.Field[string] `json:"addendum"`
	// The transfer amount in cents. A positive amount originates a credit transfer
	// pushing funds to the receiving account. A negative amount originates a debit
	// transfer pulling funds from the receiving account.
	Amount fields.Field[int64] `json:"amount,required"`
	// The description of the date of the transfer, usually in the format `YYYYMMDD`.
	// This is included in the transfer data sent to the receiving bank.
	CompanyDescriptiveDate fields.Field[string] `json:"company_descriptive_date"`
	// The data you choose to associate with the transfer. This is included in the
	// transfer data sent to the receiving bank.
	CompanyDiscretionaryData fields.Field[string] `json:"company_discretionary_data"`
	// A description of the transfer. This is included in the transfer data sent to the
	// receiving bank.
	CompanyEntryDescription fields.Field[string] `json:"company_entry_description"`
	// The name by which the recipient knows you. This is included in the transfer data
	// sent to the receiving bank.
	CompanyName fields.Field[string] `json:"company_name"`
	// The transfer effective date in
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	EffectiveDate fields.Field[time.Time] `json:"effective_date" format:"date"`
	// The ID of an External Account to initiate a transfer to. If this parameter is
	// provided, `account_number`, `routing_number`, and `funding` must be absent.
	ExternalAccountID fields.Field[string] `json:"external_account_id"`
	// The type of the account to which the transfer will be sent.
	Funding fields.Field[CreateAnACHTransferParametersFunding] `json:"funding"`
	// Your identifer for the transfer recipient.
	IndividualID fields.Field[string] `json:"individual_id"`
	// The name of the transfer recipient. This value is informational and not verified
	// by the recipient's bank.
	IndividualName fields.Field[string] `json:"individual_name"`
	// Whether the transfer requires explicit approval via the dashboard or API.
	RequireApproval fields.Field[bool] `json:"require_approval"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
	// destination account.
	RoutingNumber fields.Field[string] `json:"routing_number"`
	// The Standard Entry Class (SEC) code to use for the transfer.
	StandardEntryClassCode fields.Field[CreateAnACHTransferParametersStandardEntryClassCode] `json:"standard_entry_class_code"`
	// A description you choose to give the transfer. This will be saved with the
	// transfer details, displayed in the dashboard, and returned by the API. If
	// `individual_name` and `company_name` are not explicitly set by this API, the
	// `statement_descriptor` will be sent in those fields to the receiving bank to
	// help the customer recognize the transfer. You are highly encouraged to pass
	// `individual_name` and `company_name` instead of relying on this fallback.
	StatementDescriptor fields.Field[string] `json:"statement_descriptor,required"`
}

// MarshalJSON serializes CreateAnACHTransferParameters into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *CreateAnACHTransferParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CreateAnACHTransferParameters) String() (result string) {
	return fmt.Sprintf("&CreateAnACHTransferParameters{AccountID:%s AccountNumber:%s Addendum:%s Amount:%s CompanyDescriptiveDate:%s CompanyDiscretionaryData:%s CompanyEntryDescription:%s CompanyName:%s EffectiveDate:%s ExternalAccountID:%s Funding:%s IndividualID:%s IndividualName:%s RequireApproval:%s RoutingNumber:%s StandardEntryClassCode:%s StatementDescriptor:%s}", r.AccountID, r.AccountNumber, r.Addendum, r.Amount, r.CompanyDescriptiveDate, r.CompanyDiscretionaryData, r.CompanyEntryDescription, r.CompanyName, r.EffectiveDate, r.ExternalAccountID, r.Funding, r.IndividualID, r.IndividualName, r.RequireApproval, r.RoutingNumber, r.StandardEntryClassCode, r.StatementDescriptor)
}

type CreateAnACHTransferParametersFunding string

const (
	CreateAnACHTransferParametersFundingChecking CreateAnACHTransferParametersFunding = "checking"
	CreateAnACHTransferParametersFundingSavings  CreateAnACHTransferParametersFunding = "savings"
)

type CreateAnACHTransferParametersStandardEntryClassCode string

const (
	CreateAnACHTransferParametersStandardEntryClassCodeCorporateCreditOrDebit        CreateAnACHTransferParametersStandardEntryClassCode = "corporate_credit_or_debit"
	CreateAnACHTransferParametersStandardEntryClassCodePrearrangedPaymentsAndDeposit CreateAnACHTransferParametersStandardEntryClassCode = "prearranged_payments_and_deposit"
	CreateAnACHTransferParametersStandardEntryClassCodeInternetInitiated             CreateAnACHTransferParametersStandardEntryClassCode = "internet_initiated"
)

type ACHTransferListParams struct {
	// Return the page of entries after this one.
	Cursor fields.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit fields.Field[int64] `query:"limit"`
	// Filter ACH Transfers to those that originated from the specified Account.
	AccountID fields.Field[string] `query:"account_id"`
	// Filter ACH Transfers to those made to the specified External Account.
	ExternalAccountID fields.Field[string]                         `query:"external_account_id"`
	CreatedAt         fields.Field[ACHTransferListParamsCreatedAt] `query:"created_at"`
}

// URLQuery serializes ACHTransferListParams into a url.Values of the query
// parameters associated with this value
func (r *ACHTransferListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r ACHTransferListParams) String() (result string) {
	return fmt.Sprintf("&ACHTransferListParams{Cursor:%s Limit:%s AccountID:%s ExternalAccountID:%s CreatedAt:%s}", r.Cursor, r.Limit, r.AccountID, r.ExternalAccountID, r.CreatedAt)
}

type ACHTransferListParamsCreatedAt struct {
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

// URLQuery serializes ACHTransferListParamsCreatedAt into a url.Values of the
// query parameters associated with this value
func (r *ACHTransferListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r ACHTransferListParamsCreatedAt) String() (result string) {
	return fmt.Sprintf("&ACHTransferListParamsCreatedAt{After:%s Before:%s OnOrAfter:%s OnOrBefore:%s}", r.After, r.Before, r.OnOrAfter, r.OnOrBefore)
}

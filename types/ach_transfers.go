package types

import (
	"fmt"
	"increase/core"
	"increase/core/pjson"
	"increase/core/query"
	"increase/pagination"
	"net/url"
)

type ACHTransfer struct {
	// The Account to which the transfer belongs.
	AccountID *string `pjson:"account_id"`
	// The destination account number.
	AccountNumber *string `pjson:"account_number"`
	// Additional information that will be sent to the recipient.
	Addendum *string `pjson:"addendum"`
	// The transfer amount in USD cents. A positive amount indicates a credit transfer
	// pushing funds to the receiving account. A negative amount indicates a debit
	// transfer pulling funds from the receiving account.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transfer's
	// currency. For ACH transfers this is always equal to `usd`.
	Currency *ACHTransferCurrency `pjson:"currency"`
	// If your account requires approvals for transfers and the transfer was approved,
	// this will contain details of the approval.
	Approval *ACHTransferApproval `pjson:"approval"`
	// If your account requires approvals for transfers and the transfer was not
	// approved, this will contain details of the cancellation.
	Cancellation *ACHTransferCancellation `pjson:"cancellation"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt *string `pjson:"created_at"`
	// The identifier of the External Account the transfer was made to, if any.
	ExternalAccountID *string `pjson:"external_account_id"`
	// The ACH transfer's identifier.
	ID *string `pjson:"id"`
	// The transfer's network.
	Network *ACHTransferNetwork `pjson:"network"`
	// If the receiving bank accepts the transfer but notifies that future transfers
	// should use different details, this will contain those details.
	NotificationOfChange *ACHTransferNotificationOfChange `pjson:"notification_of_change"`
	// If your transfer is returned, this will contain details of the return.
	Return *ACHTransferReturn `pjson:"return"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber *string `pjson:"routing_number"`
	// The descriptor that will show on the recipient's bank statement.
	StatementDescriptor *string `pjson:"statement_descriptor"`
	// The lifecycle status of the transfer.
	Status *ACHTransferStatus `pjson:"status"`
	// After the transfer is submitted to FedACH, this will contain supplemental
	// details.
	Submission *ACHTransferSubmission `pjson:"submission"`
	// If the transfer was created from a template, this will be the template's ID.
	TemplateID *string `pjson:"template_id"`
	// The ID for the transaction funding the transfer.
	TransactionID *string `pjson:"transaction_id"`
	// The description of the date of the transfer.
	CompanyDescriptiveDate *string `pjson:"company_descriptive_date"`
	// The data you chose to associate with the transfer.
	CompanyDiscretionaryData *string `pjson:"company_discretionary_data"`
	// The description of the transfer you set to be shown to the recipient.
	CompanyEntryDescription *string `pjson:"company_entry_description"`
	// The name by which the recipient knows you.
	CompanyName *string `pjson:"company_name"`
	// The type of the account to which the transfer will be sent.
	Funding *ACHTransferFunding `pjson:"funding"`
	// Your identifer for the transfer recipient.
	IndividualID *string `pjson:"individual_id"`
	// The name of the transfer recipient. This value is information and not verified
	// by the recipient's bank.
	IndividualName *string `pjson:"individual_name"`
	// The Standard Entry Class (SEC) code to use for the transfer.
	StandardEntryClassCode *ACHTransferStandardEntryClassCode `pjson:"standard_entry_class_code"`
	// A constant representing the object's type. For this resource it will always be
	// `ach_transfer`.
	Type       *ACHTransferType       `pjson:"type"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into ACHTransfer using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ACHTransfer) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes ACHTransfer into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *ACHTransfer) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The Account to which the transfer belongs.
func (r *ACHTransfer) GetAccountID() (AccountID string) {
	if r != nil && r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

// The destination account number.
func (r *ACHTransfer) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

// Additional information that will be sent to the recipient.
func (r *ACHTransfer) GetAddendum() (Addendum string) {
	if r != nil && r.Addendum != nil {
		Addendum = *r.Addendum
	}
	return
}

// The transfer amount in USD cents. A positive amount indicates a credit transfer
// pushing funds to the receiving account. A negative amount indicates a debit
// transfer pulling funds from the receiving account.
func (r *ACHTransfer) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transfer's
// currency. For ACH transfers this is always equal to `usd`.
func (r *ACHTransfer) GetCurrency() (Currency ACHTransferCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// If your account requires approvals for transfers and the transfer was approved,
// this will contain details of the approval.
func (r *ACHTransfer) GetApproval() (Approval ACHTransferApproval) {
	if r != nil && r.Approval != nil {
		Approval = *r.Approval
	}
	return
}

// If your account requires approvals for transfers and the transfer was not
// approved, this will contain details of the cancellation.
func (r *ACHTransfer) GetCancellation() (Cancellation ACHTransferCancellation) {
	if r != nil && r.Cancellation != nil {
		Cancellation = *r.Cancellation
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the transfer was created.
func (r *ACHTransfer) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The identifier of the External Account the transfer was made to, if any.
func (r *ACHTransfer) GetExternalAccountID() (ExternalAccountID string) {
	if r != nil && r.ExternalAccountID != nil {
		ExternalAccountID = *r.ExternalAccountID
	}
	return
}

// The ACH transfer's identifier.
func (r *ACHTransfer) GetID() (ID string) {
	if r != nil && r.ID != nil {
		ID = *r.ID
	}
	return
}

// The transfer's network.
func (r *ACHTransfer) GetNetwork() (Network ACHTransferNetwork) {
	if r != nil && r.Network != nil {
		Network = *r.Network
	}
	return
}

// If the receiving bank accepts the transfer but notifies that future transfers
// should use different details, this will contain those details.
func (r *ACHTransfer) GetNotificationOfChange() (NotificationOfChange ACHTransferNotificationOfChange) {
	if r != nil && r.NotificationOfChange != nil {
		NotificationOfChange = *r.NotificationOfChange
	}
	return
}

// If your transfer is returned, this will contain details of the return.
func (r *ACHTransfer) GetReturn() (Return ACHTransferReturn) {
	if r != nil && r.Return != nil {
		Return = *r.Return
	}
	return
}

// The American Bankers' Association (ABA) Routing Transit Number (RTN).
func (r *ACHTransfer) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

// The descriptor that will show on the recipient's bank statement.
func (r *ACHTransfer) GetStatementDescriptor() (StatementDescriptor string) {
	if r != nil && r.StatementDescriptor != nil {
		StatementDescriptor = *r.StatementDescriptor
	}
	return
}

// The lifecycle status of the transfer.
func (r *ACHTransfer) GetStatus() (Status ACHTransferStatus) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

// After the transfer is submitted to FedACH, this will contain supplemental
// details.
func (r *ACHTransfer) GetSubmission() (Submission ACHTransferSubmission) {
	if r != nil && r.Submission != nil {
		Submission = *r.Submission
	}
	return
}

// If the transfer was created from a template, this will be the template's ID.
func (r *ACHTransfer) GetTemplateID() (TemplateID string) {
	if r != nil && r.TemplateID != nil {
		TemplateID = *r.TemplateID
	}
	return
}

// The ID for the transaction funding the transfer.
func (r *ACHTransfer) GetTransactionID() (TransactionID string) {
	if r != nil && r.TransactionID != nil {
		TransactionID = *r.TransactionID
	}
	return
}

// The description of the date of the transfer.
func (r *ACHTransfer) GetCompanyDescriptiveDate() (CompanyDescriptiveDate string) {
	if r != nil && r.CompanyDescriptiveDate != nil {
		CompanyDescriptiveDate = *r.CompanyDescriptiveDate
	}
	return
}

// The data you chose to associate with the transfer.
func (r *ACHTransfer) GetCompanyDiscretionaryData() (CompanyDiscretionaryData string) {
	if r != nil && r.CompanyDiscretionaryData != nil {
		CompanyDiscretionaryData = *r.CompanyDiscretionaryData
	}
	return
}

// The description of the transfer you set to be shown to the recipient.
func (r *ACHTransfer) GetCompanyEntryDescription() (CompanyEntryDescription string) {
	if r != nil && r.CompanyEntryDescription != nil {
		CompanyEntryDescription = *r.CompanyEntryDescription
	}
	return
}

// The name by which the recipient knows you.
func (r *ACHTransfer) GetCompanyName() (CompanyName string) {
	if r != nil && r.CompanyName != nil {
		CompanyName = *r.CompanyName
	}
	return
}

// The type of the account to which the transfer will be sent.
func (r *ACHTransfer) GetFunding() (Funding ACHTransferFunding) {
	if r != nil && r.Funding != nil {
		Funding = *r.Funding
	}
	return
}

// Your identifer for the transfer recipient.
func (r *ACHTransfer) GetIndividualID() (IndividualID string) {
	if r != nil && r.IndividualID != nil {
		IndividualID = *r.IndividualID
	}
	return
}

// The name of the transfer recipient. This value is information and not verified
// by the recipient's bank.
func (r *ACHTransfer) GetIndividualName() (IndividualName string) {
	if r != nil && r.IndividualName != nil {
		IndividualName = *r.IndividualName
	}
	return
}

// The Standard Entry Class (SEC) code to use for the transfer.
func (r *ACHTransfer) GetStandardEntryClassCode() (StandardEntryClassCode ACHTransferStandardEntryClassCode) {
	if r != nil && r.StandardEntryClassCode != nil {
		StandardEntryClassCode = *r.StandardEntryClassCode
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `ach_transfer`.
func (r *ACHTransfer) GetType() (Type ACHTransferType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

func (r ACHTransfer) String() (result string) {
	return fmt.Sprintf("&ACHTransfer{AccountID:%s AccountNumber:%s Addendum:%s Amount:%s Currency:%s Approval:%s Cancellation:%s CreatedAt:%s ExternalAccountID:%s ID:%s Network:%s NotificationOfChange:%s Return:%s RoutingNumber:%s StatementDescriptor:%s Status:%s Submission:%s TemplateID:%s TransactionID:%s CompanyDescriptiveDate:%s CompanyDiscretionaryData:%s CompanyEntryDescription:%s CompanyName:%s Funding:%s IndividualID:%s IndividualName:%s StandardEntryClassCode:%s Type:%s}", core.FmtP(r.AccountID), core.FmtP(r.AccountNumber), core.FmtP(r.Addendum), core.FmtP(r.Amount), core.FmtP(r.Currency), r.Approval, r.Cancellation, core.FmtP(r.CreatedAt), core.FmtP(r.ExternalAccountID), core.FmtP(r.ID), core.FmtP(r.Network), r.NotificationOfChange, r.Return, core.FmtP(r.RoutingNumber), core.FmtP(r.StatementDescriptor), core.FmtP(r.Status), r.Submission, core.FmtP(r.TemplateID), core.FmtP(r.TransactionID), core.FmtP(r.CompanyDescriptiveDate), core.FmtP(r.CompanyDiscretionaryData), core.FmtP(r.CompanyEntryDescription), core.FmtP(r.CompanyName), core.FmtP(r.Funding), core.FmtP(r.IndividualID), core.FmtP(r.IndividualName), core.FmtP(r.StandardEntryClassCode), core.FmtP(r.Type))
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
	ApprovedAt *string                `pjson:"approved_at"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into ACHTransferApproval using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ACHTransferApproval) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes ACHTransferApproval into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *ACHTransferApproval) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the transfer was approved.
func (r *ACHTransferApproval) GetApprovedAt() (ApprovedAt string) {
	if r != nil && r.ApprovedAt != nil {
		ApprovedAt = *r.ApprovedAt
	}
	return
}

func (r ACHTransferApproval) String() (result string) {
	return fmt.Sprintf("&ACHTransferApproval{ApprovedAt:%s}", core.FmtP(r.ApprovedAt))
}

type ACHTransferCancellation struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Transfer was canceled.
	CanceledAt *string                `pjson:"canceled_at"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into ACHTransferCancellation using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ACHTransferCancellation) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes ACHTransferCancellation into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *ACHTransferCancellation) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the Transfer was canceled.
func (r *ACHTransferCancellation) GetCanceledAt() (CanceledAt string) {
	if r != nil && r.CanceledAt != nil {
		CanceledAt = *r.CanceledAt
	}
	return
}

func (r ACHTransferCancellation) String() (result string) {
	return fmt.Sprintf("&ACHTransferCancellation{CanceledAt:%s}", core.FmtP(r.CanceledAt))
}

type ACHTransferNetwork string

const (
	ACHTransferNetworkACH ACHTransferNetwork = "ach"
)

type ACHTransferNotificationOfChange struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the notification occurred.
	CreatedAt *string `pjson:"created_at"`
	// The type of change that occurred.
	ChangeCode *string `pjson:"change_code"`
	// The corrected data.
	CorrectedData *string                `pjson:"corrected_data"`
	jsonFields    map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferNotificationOfChange using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *ACHTransferNotificationOfChange) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes ACHTransferNotificationOfChange into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *ACHTransferNotificationOfChange) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the notification occurred.
func (r *ACHTransferNotificationOfChange) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The type of change that occurred.
func (r *ACHTransferNotificationOfChange) GetChangeCode() (ChangeCode string) {
	if r != nil && r.ChangeCode != nil {
		ChangeCode = *r.ChangeCode
	}
	return
}

// The corrected data.
func (r *ACHTransferNotificationOfChange) GetCorrectedData() (CorrectedData string) {
	if r != nil && r.CorrectedData != nil {
		CorrectedData = *r.CorrectedData
	}
	return
}

func (r ACHTransferNotificationOfChange) String() (result string) {
	return fmt.Sprintf("&ACHTransferNotificationOfChange{CreatedAt:%s ChangeCode:%s CorrectedData:%s}", core.FmtP(r.CreatedAt), core.FmtP(r.ChangeCode), core.FmtP(r.CorrectedData))
}

type ACHTransferReturn struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt *string `pjson:"created_at"`
	// Why the ACH Transfer was returned.
	ReturnReasonCode *ACHTransferReturnReturnReasonCode `pjson:"return_reason_code"`
	// The identifier of the ACH Transfer associated with this return.
	TransferID *string `pjson:"transfer_id"`
	// The identifier of the Tranasaction associated with this return.
	TransactionID *string                `pjson:"transaction_id"`
	jsonFields    map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into ACHTransferReturn using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ACHTransferReturn) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes ACHTransferReturn into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *ACHTransferReturn) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the transfer was created.
func (r *ACHTransferReturn) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// Why the ACH Transfer was returned.
func (r *ACHTransferReturn) GetReturnReasonCode() (ReturnReasonCode ACHTransferReturnReturnReasonCode) {
	if r != nil && r.ReturnReasonCode != nil {
		ReturnReasonCode = *r.ReturnReasonCode
	}
	return
}

// The identifier of the ACH Transfer associated with this return.
func (r *ACHTransferReturn) GetTransferID() (TransferID string) {
	if r != nil && r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

// The identifier of the Tranasaction associated with this return.
func (r *ACHTransferReturn) GetTransactionID() (TransactionID string) {
	if r != nil && r.TransactionID != nil {
		TransactionID = *r.TransactionID
	}
	return
}

func (r ACHTransferReturn) String() (result string) {
	return fmt.Sprintf("&ACHTransferReturn{CreatedAt:%s ReturnReasonCode:%s TransferID:%s TransactionID:%s}", core.FmtP(r.CreatedAt), core.FmtP(r.ReturnReasonCode), core.FmtP(r.TransferID), core.FmtP(r.TransactionID))
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
	TraceNumber *string `pjson:"trace_number"`
	// When the ACH transfer was sent to FedACH.
	SubmittedAt *string                `pjson:"submitted_at"`
	jsonFields  map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into ACHTransferSubmission using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ACHTransferSubmission) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes ACHTransferSubmission into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *ACHTransferSubmission) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The trace number for the submission.
func (r *ACHTransferSubmission) GetTraceNumber() (TraceNumber string) {
	if r != nil && r.TraceNumber != nil {
		TraceNumber = *r.TraceNumber
	}
	return
}

// When the ACH transfer was sent to FedACH.
func (r *ACHTransferSubmission) GetSubmittedAt() (SubmittedAt string) {
	if r != nil && r.SubmittedAt != nil {
		SubmittedAt = *r.SubmittedAt
	}
	return
}

func (r ACHTransferSubmission) String() (result string) {
	return fmt.Sprintf("&ACHTransferSubmission{TraceNumber:%s SubmittedAt:%s}", core.FmtP(r.TraceNumber), core.FmtP(r.SubmittedAt))
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
	AccountID *string `pjson:"account_id"`
	// The account number for the destination account.
	AccountNumber *string `pjson:"account_number"`
	// Additional information that will be sent to the recipient. This is included in
	// the transfer data sent to the receiving bank.
	Addendum *string `pjson:"addendum"`
	// The transfer amount in cents. A positive amount originates a credit transfer
	// pushing funds to the receiving account. A negative amount originates a debit
	// transfer pulling funds from the receiving account.
	Amount *int64 `pjson:"amount"`
	// The description of the date of the transfer, usually in the format `YYYYMMDD`.
	// This is included in the transfer data sent to the receiving bank.
	CompanyDescriptiveDate *string `pjson:"company_descriptive_date"`
	// The data you choose to associate with the transfer. This is included in the
	// transfer data sent to the receiving bank.
	CompanyDiscretionaryData *string `pjson:"company_discretionary_data"`
	// A description of the transfer. This is included in the transfer data sent to the
	// receiving bank.
	CompanyEntryDescription *string `pjson:"company_entry_description"`
	// The name by which the recipient knows you. This is included in the transfer data
	// sent to the receiving bank.
	CompanyName *string `pjson:"company_name"`
	// The transfer effective date in
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	EffectiveDate *string `pjson:"effective_date"`
	// The ID of an External Account to initiate a transfer to. If this parameter is
	// provided, `account_number`, `routing_number`, and `funding` must be absent.
	ExternalAccountID *string `pjson:"external_account_id"`
	// The type of the account to which the transfer will be sent.
	Funding *CreateAnACHTransferParametersFunding `pjson:"funding"`
	// Your identifer for the transfer recipient.
	IndividualID *string `pjson:"individual_id"`
	// The name of the transfer recipient. This value is informational and not verified
	// by the recipient's bank.
	IndividualName *string `pjson:"individual_name"`
	// Whether the transfer requires explicit approval via the dashboard or API.
	RequireApproval *bool `pjson:"require_approval"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
	// destination account.
	RoutingNumber *string `pjson:"routing_number"`
	// The Standard Entry Class (SEC) code to use for the transfer.
	StandardEntryClassCode *CreateAnACHTransferParametersStandardEntryClassCode `pjson:"standard_entry_class_code"`
	// A description you choose to give the transfer. This will be saved with the
	// transfer details, displayed in the dashboard, and returned by the API. If
	// `individual_name` and `company_name` are not explicitly set by this API, the
	// `statement_descriptor` will be sent in those fields to the receiving bank to
	// help the customer recognize the transfer. You are highly encouraged to pass
	// `individual_name` and `company_name` instead of relying on this fallback.
	StatementDescriptor *string                `pjson:"statement_descriptor"`
	jsonFields          map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into CreateAnACHTransferParameters
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *CreateAnACHTransferParameters) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CreateAnACHTransferParameters into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *CreateAnACHTransferParameters) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The Increase identifier for the account that will send the transfer.
func (r *CreateAnACHTransferParameters) GetAccountID() (AccountID string) {
	if r != nil && r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

// The account number for the destination account.
func (r *CreateAnACHTransferParameters) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

// Additional information that will be sent to the recipient. This is included in
// the transfer data sent to the receiving bank.
func (r *CreateAnACHTransferParameters) GetAddendum() (Addendum string) {
	if r != nil && r.Addendum != nil {
		Addendum = *r.Addendum
	}
	return
}

// The transfer amount in cents. A positive amount originates a credit transfer
// pushing funds to the receiving account. A negative amount originates a debit
// transfer pulling funds from the receiving account.
func (r *CreateAnACHTransferParameters) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The description of the date of the transfer, usually in the format `YYYYMMDD`.
// This is included in the transfer data sent to the receiving bank.
func (r *CreateAnACHTransferParameters) GetCompanyDescriptiveDate() (CompanyDescriptiveDate string) {
	if r != nil && r.CompanyDescriptiveDate != nil {
		CompanyDescriptiveDate = *r.CompanyDescriptiveDate
	}
	return
}

// The data you choose to associate with the transfer. This is included in the
// transfer data sent to the receiving bank.
func (r *CreateAnACHTransferParameters) GetCompanyDiscretionaryData() (CompanyDiscretionaryData string) {
	if r != nil && r.CompanyDiscretionaryData != nil {
		CompanyDiscretionaryData = *r.CompanyDiscretionaryData
	}
	return
}

// A description of the transfer. This is included in the transfer data sent to the
// receiving bank.
func (r *CreateAnACHTransferParameters) GetCompanyEntryDescription() (CompanyEntryDescription string) {
	if r != nil && r.CompanyEntryDescription != nil {
		CompanyEntryDescription = *r.CompanyEntryDescription
	}
	return
}

// The name by which the recipient knows you. This is included in the transfer data
// sent to the receiving bank.
func (r *CreateAnACHTransferParameters) GetCompanyName() (CompanyName string) {
	if r != nil && r.CompanyName != nil {
		CompanyName = *r.CompanyName
	}
	return
}

// The transfer effective date in
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
func (r *CreateAnACHTransferParameters) GetEffectiveDate() (EffectiveDate string) {
	if r != nil && r.EffectiveDate != nil {
		EffectiveDate = *r.EffectiveDate
	}
	return
}

// The ID of an External Account to initiate a transfer to. If this parameter is
// provided, `account_number`, `routing_number`, and `funding` must be absent.
func (r *CreateAnACHTransferParameters) GetExternalAccountID() (ExternalAccountID string) {
	if r != nil && r.ExternalAccountID != nil {
		ExternalAccountID = *r.ExternalAccountID
	}
	return
}

// The type of the account to which the transfer will be sent.
func (r *CreateAnACHTransferParameters) GetFunding() (Funding CreateAnACHTransferParametersFunding) {
	if r != nil && r.Funding != nil {
		Funding = *r.Funding
	}
	return
}

// Your identifer for the transfer recipient.
func (r *CreateAnACHTransferParameters) GetIndividualID() (IndividualID string) {
	if r != nil && r.IndividualID != nil {
		IndividualID = *r.IndividualID
	}
	return
}

// The name of the transfer recipient. This value is informational and not verified
// by the recipient's bank.
func (r *CreateAnACHTransferParameters) GetIndividualName() (IndividualName string) {
	if r != nil && r.IndividualName != nil {
		IndividualName = *r.IndividualName
	}
	return
}

// Whether the transfer requires explicit approval via the dashboard or API.
func (r *CreateAnACHTransferParameters) GetRequireApproval() (RequireApproval bool) {
	if r != nil && r.RequireApproval != nil {
		RequireApproval = *r.RequireApproval
	}
	return
}

// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
// destination account.
func (r *CreateAnACHTransferParameters) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

// The Standard Entry Class (SEC) code to use for the transfer.
func (r *CreateAnACHTransferParameters) GetStandardEntryClassCode() (StandardEntryClassCode CreateAnACHTransferParametersStandardEntryClassCode) {
	if r != nil && r.StandardEntryClassCode != nil {
		StandardEntryClassCode = *r.StandardEntryClassCode
	}
	return
}

// A description you choose to give the transfer. This will be saved with the
// transfer details, displayed in the dashboard, and returned by the API. If
// `individual_name` and `company_name` are not explicitly set by this API, the
// `statement_descriptor` will be sent in those fields to the receiving bank to
// help the customer recognize the transfer. You are highly encouraged to pass
// `individual_name` and `company_name` instead of relying on this fallback.
func (r *CreateAnACHTransferParameters) GetStatementDescriptor() (StatementDescriptor string) {
	if r != nil && r.StatementDescriptor != nil {
		StatementDescriptor = *r.StatementDescriptor
	}
	return
}

func (r CreateAnACHTransferParameters) String() (result string) {
	return fmt.Sprintf("&CreateAnACHTransferParameters{AccountID:%s AccountNumber:%s Addendum:%s Amount:%s CompanyDescriptiveDate:%s CompanyDiscretionaryData:%s CompanyEntryDescription:%s CompanyName:%s EffectiveDate:%s ExternalAccountID:%s Funding:%s IndividualID:%s IndividualName:%s RequireApproval:%s RoutingNumber:%s StandardEntryClassCode:%s StatementDescriptor:%s}", core.FmtP(r.AccountID), core.FmtP(r.AccountNumber), core.FmtP(r.Addendum), core.FmtP(r.Amount), core.FmtP(r.CompanyDescriptiveDate), core.FmtP(r.CompanyDiscretionaryData), core.FmtP(r.CompanyEntryDescription), core.FmtP(r.CompanyName), core.FmtP(r.EffectiveDate), core.FmtP(r.ExternalAccountID), core.FmtP(r.Funding), core.FmtP(r.IndividualID), core.FmtP(r.IndividualName), core.FmtP(r.RequireApproval), core.FmtP(r.RoutingNumber), core.FmtP(r.StandardEntryClassCode), core.FmtP(r.StatementDescriptor))
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
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int64 `query:"limit"`
	// Filter ACH Transfers to those that originated from the specified Account.
	AccountID *string `query:"account_id"`
	// Filter ACH Transfers to those made to the specified External Account.
	ExternalAccountID *string                          `query:"external_account_id"`
	CreatedAt         *ACHTransfersListParamsCreatedAt `query:"created_at"`
	jsonFields        map[string]interface{}           `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into ACHTransferListParams using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ACHTransferListParams) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes ACHTransferListParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *ACHTransferListParams) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes ACHTransferListParams into a url.Values of the query
// parameters associated with this value
func (r *ACHTransferListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// Return the page of entries after this one.
func (r *ACHTransferListParams) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *ACHTransferListParams) GetLimit() (Limit int64) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// Filter ACH Transfers to those that originated from the specified Account.
func (r *ACHTransferListParams) GetAccountID() (AccountID string) {
	if r != nil && r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

// Filter ACH Transfers to those made to the specified External Account.
func (r *ACHTransferListParams) GetExternalAccountID() (ExternalAccountID string) {
	if r != nil && r.ExternalAccountID != nil {
		ExternalAccountID = *r.ExternalAccountID
	}
	return
}

func (r *ACHTransferListParams) GetCreatedAt() (CreatedAt ACHTransfersListParamsCreatedAt) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

func (r ACHTransferListParams) String() (result string) {
	return fmt.Sprintf("&ACHTransferListParams{Cursor:%s Limit:%s AccountID:%s ExternalAccountID:%s CreatedAt:%s}", core.FmtP(r.Cursor), core.FmtP(r.Limit), core.FmtP(r.AccountID), core.FmtP(r.ExternalAccountID), r.CreatedAt)
}

type ACHTransfersListParamsCreatedAt struct {
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
// ACHTransfersListParamsCreatedAt using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *ACHTransfersListParamsCreatedAt) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes ACHTransfersListParamsCreatedAt into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *ACHTransfersListParamsCreatedAt) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes ACHTransfersListParamsCreatedAt into a url.Values of the
// query parameters associated with this value
func (r *ACHTransfersListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r *ACHTransfersListParamsCreatedAt) GetAfter() (After string) {
	if r != nil && r.After != nil {
		After = *r.After
	}
	return
}

// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r *ACHTransfersListParamsCreatedAt) GetBefore() (Before string) {
	if r != nil && r.Before != nil {
		Before = *r.Before
	}
	return
}

// Return results on or after this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r *ACHTransfersListParamsCreatedAt) GetOnOrAfter() (OnOrAfter string) {
	if r != nil && r.OnOrAfter != nil {
		OnOrAfter = *r.OnOrAfter
	}
	return
}

// Return results on or before this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r *ACHTransfersListParamsCreatedAt) GetOnOrBefore() (OnOrBefore string) {
	if r != nil && r.OnOrBefore != nil {
		OnOrBefore = *r.OnOrBefore
	}
	return
}

func (r ACHTransfersListParamsCreatedAt) String() (result string) {
	return fmt.Sprintf("&ACHTransfersListParamsCreatedAt{After:%s Before:%s OnOrAfter:%s OnOrBefore:%s}", core.FmtP(r.After), core.FmtP(r.Before), core.FmtP(r.OnOrAfter), core.FmtP(r.OnOrBefore))
}

type ACHTransferList struct {
	// The contents of the list.
	Data *[]ACHTransfer `pjson:"data"`
	// A pointer to a place in the list.
	NextCursor *string                `pjson:"next_cursor"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into ACHTransferList using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ACHTransferList) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes ACHTransferList into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *ACHTransferList) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes ACHTransferList into a url.Values of the query parameters
// associated with this value
func (r *ACHTransferList) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// The contents of the list.
func (r *ACHTransferList) GetData() (Data []ACHTransfer) {
	if r != nil && r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r *ACHTransferList) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

func (r ACHTransferList) String() (result string) {
	return fmt.Sprintf("&ACHTransferList{Data:%s NextCursor:%s}", core.Fmt(r.Data), core.FmtP(r.NextCursor))
}

type ACHTransfersPage struct {
	*pagination.Page[ACHTransfer]
}

func (r *ACHTransfersPage) ACHTransfer() *ACHTransfer {
	return r.Current()
}

func (r *ACHTransfersPage) NextPage() (*ACHTransfersPage, error) {
	if page, err := r.Page.NextPage(); err != nil {
		return nil, err
	} else {
		return &ACHTransfersPage{page}, nil
	}
}

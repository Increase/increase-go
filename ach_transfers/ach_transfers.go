package ach_transfers

import "context"
import "increase/core"
import "fmt"
import "increase/pagination"

type ACHTransferService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewACHTransferService(requester core.Requester) (r *ACHTransferService) {
	r = &ACHTransferService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

type PreloadedACHTransferService struct {
	ACHTransfers *ACHTransferService
}

func (r *PreloadedACHTransferService) Init(service *ACHTransferService) {
	r.ACHTransfers = service
}

func NewPreloadedACHTransferService(service *ACHTransferService) (r *PreloadedACHTransferService) {
	r = &PreloadedACHTransferService{}
	r.Init(service)
	return
}

//
type ACHTransfer struct {
	// The Account to which the transfer belongs.
	AccountID *string `json:"account_id"`
	// The destination account number.
	AccountNumber *string `json:"account_number"`
	// Additional information that will be sent to the recipient.
	Addendum *string `json:"addendum"`
	// The transfer amount in USD cents. A positive amount indicates a credit transfer
	// pushing funds to the receiving account. A negative amount indicates a debit
	// transfer pulling funds from the receiving account.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transfer's
	// currency. For ACH transfers this is always equal to `usd`.
	Currency *ACHTransferCurrency `json:"currency"`
	// If your account requires approvals for transfers and the transfer was approved,
	// this will contain details of the approval.
	Approval *ACHTransferApproval `json:"approval"`
	// If your account requires approvals for transfers and the transfer was not
	// approved, this will contain details of the cancellation.
	Cancellation *ACHTransferCancellation `json:"cancellation"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt *string `json:"created_at"`
	// The identifier of the External Account the transfer was made to, if any.
	ExternalAccountID *string `json:"external_account_id"`
	// The ACH transfer's identifier.
	ID *string `json:"id"`
	// The transfer's network.
	Network *ACHTransferNetwork `json:"network"`
	// If the receiving bank accepts the transfer but notifies that future transfers
	// should use different details, this will contain those details.
	NotificationOfChange *ACHTransferNotificationOfChange `json:"notification_of_change"`
	// If your transfer is returned, this will contain details of the return.
	Return *ACHTransferReturn `json:"return"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber *string `json:"routing_number"`
	// The descriptor that will show on the recipient's bank statement.
	StatementDescriptor *string `json:"statement_descriptor"`
	// The lifecycle status of the transfer.
	Status *ACHTransferStatus `json:"status"`
	// After the transfer is submitted to FedACH, this will contain supplemental
	// details.
	Submission *ACHTransferSubmission `json:"submission"`
	// If the transfer was created from a template, this will be the template's ID.
	TemplateID *string `json:"template_id"`
	// The ID for the transaction funding the transfer.
	TransactionID *string `json:"transaction_id"`
	// A constant representing the object's type. For this resource it will always be
	// `ach_transfer`.
	Type *ACHTransferType `json:"type"`
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
func (r *ACHTransfer) GetAmount() (Amount int) {
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

// A constant representing the object's type. For this resource it will always be
// `ach_transfer`.
func (r *ACHTransfer) GetType() (Type ACHTransferType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
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

//
type ACHTransferApproval struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was approved.
	ApprovedAt *string `json:"approved_at"`
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the transfer was approved.
func (r *ACHTransferApproval) GetApprovedAt() (ApprovedAt string) {
	if r != nil && r.ApprovedAt != nil {
		ApprovedAt = *r.ApprovedAt
	}
	return
}

//
type ACHTransferCancellation struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Transfer was canceled.
	CanceledAt *string `json:"canceled_at"`
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the Transfer was canceled.
func (r *ACHTransferCancellation) GetCanceledAt() (CanceledAt string) {
	if r != nil && r.CanceledAt != nil {
		CanceledAt = *r.CanceledAt
	}
	return
}

type ACHTransferNetwork string

const (
	ACHTransferNetworkACH ACHTransferNetwork = "ach"
)

//
type ACHTransferNotificationOfChange struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the notification occurred.
	CreatedAt *string `json:"created_at"`
	// The type of change that occurred.
	ChangeCode *string `json:"change_code"`
	// The corrected data.
	CorrectedData *string `json:"corrected_data"`
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

//
type ACHTransferReturn struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt *string `json:"created_at"`
	// Why the ACH Transfer was returned.
	ReturnReasonCode *ACHTransferReturnReturnReasonCode `json:"return_reason_code"`
	// The identifier of the ACH Transfer associated with this return.
	TransferID *string `json:"transfer_id"`
	// The identifier of the Tranasaction associated with this return.
	TransactionID *string `json:"transaction_id"`
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

//
type ACHTransferSubmission struct {
	// The trace number for the submission.
	TraceNumber *string `json:"trace_number"`
}

// The trace number for the submission.
func (r *ACHTransferSubmission) GetTraceNumber() (TraceNumber string) {
	if r != nil && r.TraceNumber != nil {
		TraceNumber = *r.TraceNumber
	}
	return
}

type ACHTransferType string

const (
	ACHTransferTypeACHTransfer ACHTransferType = "ach_transfer"
)

type CreateAnACHTransferParameters struct {
	// The identifier for the account that will send the transfer.
	AccountID *string `json:"account_id"`
	// The account number for the destination account.
	AccountNumber *string `json:"account_number,omitempty"`
	// Additional information that will be sent to the recipient.
	Addendum *string `json:"addendum,omitempty"`
	// The transfer amount in cents. A positive amount originates a credit transfer
	// pushing funds to the receiving account. A negative amount originates a debit
	// transfer pulling funds from the receiving account.
	Amount *int `json:"amount"`
	// The description of the date of the transfer.
	CompanyDescriptiveDate *string `json:"company_descriptive_date,omitempty"`
	// The data you choose to associate with the transfer.
	CompanyDiscretionaryData *string `json:"company_discretionary_data,omitempty"`
	// The description of the transfer you wish to be shown to the recipient.
	CompanyEntryDescription *string `json:"company_entry_description,omitempty"`
	// The name by which the recipient knows you.
	CompanyName *string `json:"company_name,omitempty"`
	// The transfer effective date in
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	EffectiveDate *string `json:"effective_date,omitempty"`
	// The ID of an External Account to initiate a transfer to. If this parameter is
	// provided, `account_number`, `routing_number`, and `funding` must be absent.
	ExternalAccountID *string `json:"external_account_id,omitempty"`
	// The type of the account to which the transfer will be sent.
	Funding *CreateAnACHTransferParametersFunding `json:"funding,omitempty"`
	// Your identifer for the transfer recipient.
	IndividualID *string `json:"individual_id,omitempty"`
	// The name of the transfer recipient. This value is information and not verified
	// by the recipient's bank.
	IndividualName *string `json:"individual_name,omitempty"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
	// destination account.
	RoutingNumber *string `json:"routing_number,omitempty"`
	// The Standard Entry Class (SEC) code to use for the transfer.
	StandardEntryClassCode *CreateAnACHTransferParametersStandardEntryClassCode `json:"standard_entry_class_code,omitempty"`
	// The description you choose to give the transfer. This will be shown to the
	// recipient.
	StatementDescriptor *string `json:"statement_descriptor"`
}

// The identifier for the account that will send the transfer.
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

// Additional information that will be sent to the recipient.
func (r *CreateAnACHTransferParameters) GetAddendum() (Addendum string) {
	if r != nil && r.Addendum != nil {
		Addendum = *r.Addendum
	}
	return
}

// The transfer amount in cents. A positive amount originates a credit transfer
// pushing funds to the receiving account. A negative amount originates a debit
// transfer pulling funds from the receiving account.
func (r *CreateAnACHTransferParameters) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The description of the date of the transfer.
func (r *CreateAnACHTransferParameters) GetCompanyDescriptiveDate() (CompanyDescriptiveDate string) {
	if r != nil && r.CompanyDescriptiveDate != nil {
		CompanyDescriptiveDate = *r.CompanyDescriptiveDate
	}
	return
}

// The data you choose to associate with the transfer.
func (r *CreateAnACHTransferParameters) GetCompanyDiscretionaryData() (CompanyDiscretionaryData string) {
	if r != nil && r.CompanyDiscretionaryData != nil {
		CompanyDiscretionaryData = *r.CompanyDiscretionaryData
	}
	return
}

// The description of the transfer you wish to be shown to the recipient.
func (r *CreateAnACHTransferParameters) GetCompanyEntryDescription() (CompanyEntryDescription string) {
	if r != nil && r.CompanyEntryDescription != nil {
		CompanyEntryDescription = *r.CompanyEntryDescription
	}
	return
}

// The name by which the recipient knows you.
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

// The name of the transfer recipient. This value is information and not verified
// by the recipient's bank.
func (r *CreateAnACHTransferParameters) GetIndividualName() (IndividualName string) {
	if r != nil && r.IndividualName != nil {
		IndividualName = *r.IndividualName
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

// The description you choose to give the transfer. This will be shown to the
// recipient.
func (r *CreateAnACHTransferParameters) GetStatementDescriptor() (StatementDescriptor string) {
	if r != nil && r.StatementDescriptor != nil {
		StatementDescriptor = *r.StatementDescriptor
	}
	return
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

type ListACHTransfersQuery struct {
	// Return the page of entries after this one.
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int `query:"limit"`
	// Filter ACH Transfers to those that originated from the specified Account.
	AccountID *string `query:"account_id"`
	// Filter ACH Transfers to those made to the specified External Account.
	ExternalAccountID *string                         `query:"external_account_id"`
	CreatedAt         *ListACHTransfersQueryCreatedAt `query:"created_at"`
}

// Return the page of entries after this one.
func (r *ListACHTransfersQuery) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *ListACHTransfersQuery) GetLimit() (Limit int) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// Filter ACH Transfers to those that originated from the specified Account.
func (r *ListACHTransfersQuery) GetAccountID() (AccountID string) {
	if r != nil && r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

// Filter ACH Transfers to those made to the specified External Account.
func (r *ListACHTransfersQuery) GetExternalAccountID() (ExternalAccountID string) {
	if r != nil && r.ExternalAccountID != nil {
		ExternalAccountID = *r.ExternalAccountID
	}
	return
}

func (r *ListACHTransfersQuery) GetCreatedAt() (CreatedAt ListACHTransfersQueryCreatedAt) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

type ListACHTransfersQueryCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After *string `json:"after,omitempty"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before *string `json:"before,omitempty"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter *string `json:"on_or_after,omitempty"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore *string `json:"on_or_before,omitempty"`
}

// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r *ListACHTransfersQueryCreatedAt) GetAfter() (After string) {
	if r != nil && r.After != nil {
		After = *r.After
	}
	return
}

// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r *ListACHTransfersQueryCreatedAt) GetBefore() (Before string) {
	if r != nil && r.Before != nil {
		Before = *r.Before
	}
	return
}

// Return results on or after this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r *ListACHTransfersQueryCreatedAt) GetOnOrAfter() (OnOrAfter string) {
	if r != nil && r.OnOrAfter != nil {
		OnOrAfter = *r.OnOrAfter
	}
	return
}

// Return results on or before this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r *ListACHTransfersQueryCreatedAt) GetOnOrBefore() (OnOrBefore string) {
	if r != nil && r.OnOrBefore != nil {
		OnOrBefore = *r.OnOrBefore
	}
	return
}

//
type ACHTransferList struct {
	// The contents of the list.
	Data *[]ACHTransfer `json:"data"`
	// A pointer to a place in the list.
	NextCursor *string `json:"next_cursor"`
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

func (r *ACHTransferService) Create(ctx context.Context, body *CreateAnACHTransferParameters, opts ...*core.RequestOpts) (res *ACHTransfer, err error) {
	err = r.post(
		ctx,
		"/ach_transfers",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *PreloadedACHTransferService) Create(ctx context.Context, body *CreateAnACHTransferParameters, opts ...*core.RequestOpts) (res *ACHTransfer, err error) {
	err = r.ACHTransfers.post(
		ctx,
		"/ach_transfers",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *ACHTransferService) Retrieve(ctx context.Context, ach_transfer_id string, opts ...*core.RequestOpts) (res *ACHTransfer, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/ach_transfers/%s", ach_transfer_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *PreloadedACHTransferService) Retrieve(ctx context.Context, ach_transfer_id string, opts ...*core.RequestOpts) (res *ACHTransfer, err error) {
	err = r.ACHTransfers.get(
		ctx,
		fmt.Sprintf("/ach_transfers/%s", ach_transfer_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

type ACHTransfersPage struct {
	*pagination.Page[ACHTransfer]
}

func (r *ACHTransfersPage) ACHTransfer() *ACHTransfer {
	return r.Current()
}

func (r *ACHTransfersPage) GetNextPage() (*ACHTransfersPage, error) {
	if page, err := r.Page.GetNextPage(); err != nil {
		return nil, err
	} else {
		return &ACHTransfersPage{page}, nil
	}
}

func (r *ACHTransferService) List(ctx context.Context, query *ListACHTransfersQuery, opts ...*core.RequestOpts) (res *ACHTransfersPage, err error) {
	page := &ACHTransfersPage{
		Page: &pagination.Page[ACHTransfer]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/ach_transfers",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

func (r *PreloadedACHTransferService) List(ctx context.Context, query *ListACHTransfersQuery, opts ...*core.RequestOpts) (res *ACHTransfersPage, err error) {
	page := &ACHTransfersPage{
		Page: &pagination.Page[ACHTransfer]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/ach_transfers",
			},
			Requester: r.ACHTransfers.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

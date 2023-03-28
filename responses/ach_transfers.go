package responses

import (
	"time"

	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/pagination"
)

type ACHTransfer struct {
	// The Account to which the transfer belongs.
	AccountID string `json:"account_id,required"`
	// The destination account number.
	AccountNumber string `json:"account_number,required"`
	// Additional information that will be sent to the recipient.
	Addendum string `json:"addendum,required,nullable"`
	// The transfer amount in USD cents. A positive amount indicates a credit transfer
	// pushing funds to the receiving account. A negative amount indicates a debit
	// transfer pulling funds from the receiving account.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transfer's
	// currency. For ACH transfers this is always equal to `usd`.
	Currency ACHTransferCurrency `json:"currency,required"`
	// If your account requires approvals for transfers and the transfer was approved,
	// this will contain details of the approval.
	Approval ACHTransferApproval `json:"approval,required,nullable"`
	// If your account requires approvals for transfers and the transfer was not
	// approved, this will contain details of the cancellation.
	Cancellation ACHTransferCancellation `json:"cancellation,required,nullable"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The identifier of the External Account the transfer was made to, if any.
	ExternalAccountID string `json:"external_account_id,required,nullable"`
	// The ACH transfer's identifier.
	ID string `json:"id,required"`
	// The transfer's network.
	Network ACHTransferNetwork `json:"network,required"`
	// If the receiving bank accepts the transfer but notifies that future transfers
	// should use different details, this will contain those details.
	NotificationOfChange ACHTransferNotificationOfChange `json:"notification_of_change,required,nullable"`
	// If your transfer is returned, this will contain details of the return.
	Return ACHTransferReturn `json:"return,required,nullable"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber string `json:"routing_number,required"`
	// The descriptor that will show on the recipient's bank statement.
	StatementDescriptor string `json:"statement_descriptor,required"`
	// The lifecycle status of the transfer.
	Status ACHTransferStatus `json:"status,required"`
	// After the transfer is submitted to FedACH, this will contain supplemental
	// details.
	Submission ACHTransferSubmission `json:"submission,required,nullable"`
	// If the transfer was created from a template, this will be the template's ID.
	TemplateID string `json:"template_id,required,nullable"`
	// The ID for the transaction funding the transfer.
	TransactionID string `json:"transaction_id,required,nullable"`
	// The description of the date of the transfer.
	CompanyDescriptiveDate string `json:"company_descriptive_date,required,nullable"`
	// The data you chose to associate with the transfer.
	CompanyDiscretionaryData string `json:"company_discretionary_data,required,nullable"`
	// The description of the transfer you set to be shown to the recipient.
	CompanyEntryDescription string `json:"company_entry_description,required,nullable"`
	// The name by which the recipient knows you.
	CompanyName string `json:"company_name,required,nullable"`
	// The type of the account to which the transfer will be sent.
	Funding ACHTransferFunding `json:"funding,required"`
	// Your identifer for the transfer recipient.
	IndividualID string `json:"individual_id,required,nullable"`
	// The name of the transfer recipient. This value is information and not verified
	// by the recipient's bank.
	IndividualName string `json:"individual_name,required,nullable"`
	// The transfer effective date in
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	EffectiveDate time.Time `json:"effective_date,required,nullable" format:"date"`
	// The Standard Entry Class (SEC) code to use for the transfer.
	StandardEntryClassCode ACHTransferStandardEntryClassCode `json:"standard_entry_class_code,required"`
	// A constant representing the object's type. For this resource it will always be
	// `ach_transfer`.
	Type ACHTransferType `json:"type,required"`
	JSON ACHTransferJSON
}

type ACHTransferJSON struct {
	AccountID                pjson.Metadata
	AccountNumber            pjson.Metadata
	Addendum                 pjson.Metadata
	Amount                   pjson.Metadata
	Currency                 pjson.Metadata
	Approval                 pjson.Metadata
	Cancellation             pjson.Metadata
	CreatedAt                pjson.Metadata
	ExternalAccountID        pjson.Metadata
	ID                       pjson.Metadata
	Network                  pjson.Metadata
	NotificationOfChange     pjson.Metadata
	Return                   pjson.Metadata
	RoutingNumber            pjson.Metadata
	StatementDescriptor      pjson.Metadata
	Status                   pjson.Metadata
	Submission               pjson.Metadata
	TemplateID               pjson.Metadata
	TransactionID            pjson.Metadata
	CompanyDescriptiveDate   pjson.Metadata
	CompanyDiscretionaryData pjson.Metadata
	CompanyEntryDescription  pjson.Metadata
	CompanyName              pjson.Metadata
	Funding                  pjson.Metadata
	IndividualID             pjson.Metadata
	IndividualName           pjson.Metadata
	EffectiveDate            pjson.Metadata
	StandardEntryClassCode   pjson.Metadata
	Type                     pjson.Metadata
	Raw                      []byte
	Extras                   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into ACHTransfer using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ACHTransfer) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	ApprovedAt time.Time `json:"approved_at,required" format:"date-time"`
	JSON       ACHTransferApprovalJSON
}

type ACHTransferApprovalJSON struct {
	ApprovedAt pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into ACHTransferApproval using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ACHTransferApproval) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type ACHTransferCancellation struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Transfer was canceled.
	CanceledAt time.Time `json:"canceled_at,required" format:"date-time"`
	JSON       ACHTransferCancellationJSON
}

type ACHTransferCancellationJSON struct {
	CanceledAt pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into ACHTransferCancellation using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ACHTransferCancellation) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type ACHTransferNetwork string

const (
	ACHTransferNetworkACH ACHTransferNetwork = "ach"
)

type ACHTransferNotificationOfChange struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the notification occurred.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The type of change that occurred.
	ChangeCode string `json:"change_code,required"`
	// The corrected data.
	CorrectedData string `json:"corrected_data,required"`
	JSON          ACHTransferNotificationOfChangeJSON
}

type ACHTransferNotificationOfChangeJSON struct {
	CreatedAt     pjson.Metadata
	ChangeCode    pjson.Metadata
	CorrectedData pjson.Metadata
	Raw           []byte
	Extras        map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// ACHTransferNotificationOfChange using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *ACHTransferNotificationOfChange) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type ACHTransferReturn struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Why the ACH Transfer was returned.
	ReturnReasonCode ACHTransferReturnReturnReasonCode `json:"return_reason_code,required"`
	// The identifier of the ACH Transfer associated with this return.
	TransferID string `json:"transfer_id,required"`
	// The identifier of the Tranasaction associated with this return.
	TransactionID string `json:"transaction_id,required"`
	JSON          ACHTransferReturnJSON
}

type ACHTransferReturnJSON struct {
	CreatedAt        pjson.Metadata
	ReturnReasonCode pjson.Metadata
	TransferID       pjson.Metadata
	TransactionID    pjson.Metadata
	Raw              []byte
	Extras           map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into ACHTransferReturn using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ACHTransferReturn) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	TraceNumber string `json:"trace_number,required"`
	// When the ACH transfer was sent to FedACH.
	SubmittedAt time.Time `json:"submitted_at,required" format:"date-time"`
	JSON        ACHTransferSubmissionJSON
}

type ACHTransferSubmissionJSON struct {
	TraceNumber pjson.Metadata
	SubmittedAt pjson.Metadata
	Raw         []byte
	Extras      map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into ACHTransferSubmission using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ACHTransferSubmission) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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

type ACHTransferList struct {
	// The contents of the list.
	Data []ACHTransfer `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       ACHTransferListJSON
}

type ACHTransferListJSON struct {
	Data       pjson.Metadata
	NextCursor pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into ACHTransferList using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ACHTransferList) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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

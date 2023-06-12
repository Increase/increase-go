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

// ACHTransferService contains methods and other services that help with
// interacting with the increase API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewACHTransferService] method
// instead.
type ACHTransferService struct {
	Options []option.RequestOption
}

// NewACHTransferService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewACHTransferService(opts ...option.RequestOption) (r *ACHTransferService) {
	r = &ACHTransferService{}
	r.Options = opts
	return
}

// Create an ACH Transfer
func (r *ACHTransferService) New(ctx context.Context, body ACHTransferNewParams, opts ...option.RequestOption) (res *ACHTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := "ach_transfers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve an ACH Transfer
func (r *ACHTransferService) Get(ctx context.Context, achTransferID string, opts ...option.RequestOption) (res *ACHTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("ach_transfers/%s", achTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List ACH Transfers
func (r *ACHTransferService) List(ctx context.Context, query ACHTransferListParams, opts ...option.RequestOption) (res *shared.Page[ACHTransfer], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "ach_transfers"
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

// List ACH Transfers
func (r *ACHTransferService) ListAutoPaging(ctx context.Context, query ACHTransferListParams, opts ...option.RequestOption) *shared.PageAutoPager[ACHTransfer] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Approves an ACH Transfer in a pending_approval state.
func (r *ACHTransferService) Approve(ctx context.Context, achTransferID string, opts ...option.RequestOption) (res *ACHTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("ach_transfers/%s/approve", achTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Cancels an ACH Transfer in a pending_approval state.
func (r *ACHTransferService) Cancel(ctx context.Context, achTransferID string, opts ...option.RequestOption) (res *ACHTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("ach_transfers/%s/cancel", achTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// ACH transfers move funds between your Increase account and any other account
// accessible by the Automated Clearing House (ACH).
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
	NotificationsOfChange []ACHTransferNotificationsOfChange `json:"notifications_of_change,required"`
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
	JSON achTransferJSON
}

// achTransferJSON contains the JSON metadata for the struct [ACHTransfer]
type achTransferJSON struct {
	AccountID                apijson.Field
	AccountNumber            apijson.Field
	Addendum                 apijson.Field
	Amount                   apijson.Field
	Currency                 apijson.Field
	Approval                 apijson.Field
	Cancellation             apijson.Field
	CreatedAt                apijson.Field
	ExternalAccountID        apijson.Field
	ID                       apijson.Field
	Network                  apijson.Field
	NotificationsOfChange    apijson.Field
	Return                   apijson.Field
	RoutingNumber            apijson.Field
	StatementDescriptor      apijson.Field
	Status                   apijson.Field
	Submission               apijson.Field
	TransactionID            apijson.Field
	CompanyDescriptiveDate   apijson.Field
	CompanyDiscretionaryData apijson.Field
	CompanyEntryDescription  apijson.Field
	CompanyName              apijson.Field
	Funding                  apijson.Field
	IndividualID             apijson.Field
	IndividualName           apijson.Field
	EffectiveDate            apijson.Field
	StandardEntryClassCode   apijson.Field
	Type                     apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *ACHTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transfer's
// currency. For ACH transfers this is always equal to `usd`.
type ACHTransferCurrency string

const (
	ACHTransferCurrencyCad ACHTransferCurrency = "CAD"
	ACHTransferCurrencyChf ACHTransferCurrency = "CHF"
	ACHTransferCurrencyEur ACHTransferCurrency = "EUR"
	ACHTransferCurrencyGbp ACHTransferCurrency = "GBP"
	ACHTransferCurrencyJpy ACHTransferCurrency = "JPY"
	ACHTransferCurrencyUsd ACHTransferCurrency = "USD"
)

// If your account requires approvals for transfers and the transfer was approved,
// this will contain details of the approval.
type ACHTransferApproval struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was approved.
	ApprovedAt time.Time `json:"approved_at,required" format:"date-time"`
	// If the Transfer was approved by a user in the dashboard, the email address of
	// that user.
	ApprovedBy string `json:"approved_by,required,nullable"`
	JSON       achTransferApprovalJSON
}

// achTransferApprovalJSON contains the JSON metadata for the struct
// [ACHTransferApproval]
type achTransferApprovalJSON struct {
	ApprovedAt  apijson.Field
	ApprovedBy  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACHTransferApproval) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// If your account requires approvals for transfers and the transfer was not
// approved, this will contain details of the cancellation.
type ACHTransferCancellation struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Transfer was canceled.
	CanceledAt time.Time `json:"canceled_at,required" format:"date-time"`
	// If the Transfer was canceled by a user in the dashboard, the email address of
	// that user.
	CanceledBy string `json:"canceled_by,required,nullable"`
	JSON       achTransferCancellationJSON
}

// achTransferCancellationJSON contains the JSON metadata for the struct
// [ACHTransferCancellation]
type achTransferCancellationJSON struct {
	CanceledAt  apijson.Field
	CanceledBy  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACHTransferCancellation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The transfer's network.
type ACHTransferNetwork string

const (
	ACHTransferNetworkACH ACHTransferNetwork = "ach"
)

type ACHTransferNotificationsOfChange struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the notification occurred.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The type of change that occurred.
	ChangeCode string `json:"change_code,required"`
	// The corrected data.
	CorrectedData string `json:"corrected_data,required"`
	JSON          achTransferNotificationsOfChangeJSON
}

// achTransferNotificationsOfChangeJSON contains the JSON metadata for the struct
// [ACHTransferNotificationsOfChange]
type achTransferNotificationsOfChangeJSON struct {
	CreatedAt     apijson.Field
	ChangeCode    apijson.Field
	CorrectedData apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ACHTransferNotificationsOfChange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// If your transfer is returned, this will contain details of the return.
type ACHTransferReturn struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Why the ACH Transfer was returned.
	ReturnReasonCode ACHTransferReturnReturnReasonCode `json:"return_reason_code,required"`
	// The three character ACH return code, in the range R01 to R85.
	RawReturnReasonCode string `json:"raw_return_reason_code,required"`
	// The identifier of the ACH Transfer associated with this return.
	TransferID string `json:"transfer_id,required"`
	// The identifier of the Tranasaction associated with this return.
	TransactionID string `json:"transaction_id,required"`
	JSON          achTransferReturnJSON
}

// achTransferReturnJSON contains the JSON metadata for the struct
// [ACHTransferReturn]
type achTransferReturnJSON struct {
	CreatedAt           apijson.Field
	ReturnReasonCode    apijson.Field
	RawReturnReasonCode apijson.Field
	TransferID          apijson.Field
	TransactionID       apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ACHTransferReturn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Why the ACH Transfer was returned.
type ACHTransferReturnReturnReasonCode string

const (
	ACHTransferReturnReturnReasonCodeInsufficientFund                                            ACHTransferReturnReturnReasonCode = "insufficient_fund"
	ACHTransferReturnReturnReasonCodeNoAccount                                                   ACHTransferReturnReturnReasonCode = "no_account"
	ACHTransferReturnReturnReasonCodeAccountClosed                                               ACHTransferReturnReturnReasonCode = "account_closed"
	ACHTransferReturnReturnReasonCodeInvalidAccountNumberStructure                               ACHTransferReturnReturnReasonCode = "invalid_account_number_structure"
	ACHTransferReturnReturnReasonCodeAccountFrozenEntryReturnedPerOfacInstruction                ACHTransferReturnReturnReasonCode = "account_frozen_entry_returned_per_ofac_instruction"
	ACHTransferReturnReturnReasonCodeCreditEntryRefusedByReceiver                                ACHTransferReturnReturnReasonCode = "credit_entry_refused_by_receiver"
	ACHTransferReturnReturnReasonCodeUnauthorizedDebitToConsumerAccountUsingCorporateSecCode     ACHTransferReturnReturnReasonCode = "unauthorized_debit_to_consumer_account_using_corporate_sec_code"
	ACHTransferReturnReturnReasonCodeCorporateCustomerAdvisedNotAuthorized                       ACHTransferReturnReturnReasonCode = "corporate_customer_advised_not_authorized"
	ACHTransferReturnReturnReasonCodePaymentStopped                                              ACHTransferReturnReturnReasonCode = "payment_stopped"
	ACHTransferReturnReturnReasonCodeNonTransactionAccount                                       ACHTransferReturnReturnReasonCode = "non_transaction_account"
	ACHTransferReturnReturnReasonCodeUncollectedFunds                                            ACHTransferReturnReturnReasonCode = "uncollected_funds"
	ACHTransferReturnReturnReasonCodeRoutingNumberCheckDigitError                                ACHTransferReturnReturnReasonCode = "routing_number_check_digit_error"
	ACHTransferReturnReturnReasonCodeCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete   ACHTransferReturnReturnReasonCode = "customer_advised_unauthorized_improper_ineligible_or_incomplete"
	ACHTransferReturnReturnReasonCodeAmountFieldError                                            ACHTransferReturnReturnReasonCode = "amount_field_error"
	ACHTransferReturnReturnReasonCodeAuthorizationRevokedByCustomer                              ACHTransferReturnReturnReasonCode = "authorization_revoked_by_customer"
	ACHTransferReturnReturnReasonCodeInvalidACHRoutingNumber                                     ACHTransferReturnReturnReasonCode = "invalid_ach_routing_number"
	ACHTransferReturnReturnReasonCodeFileRecordEditCriteria                                      ACHTransferReturnReturnReasonCode = "file_record_edit_criteria"
	ACHTransferReturnReturnReasonCodeEnrInvalidIndividualName                                    ACHTransferReturnReturnReasonCode = "enr_invalid_individual_name"
	ACHTransferReturnReturnReasonCodeReturnedPerOdfiRequest                                      ACHTransferReturnReturnReasonCode = "returned_per_odfi_request"
	ACHTransferReturnReturnReasonCodeLimitedParticipationDfi                                     ACHTransferReturnReturnReasonCode = "limited_participation_dfi"
	ACHTransferReturnReturnReasonCodeIncorrectlyCodedOutboundInternationalPayment                ACHTransferReturnReturnReasonCode = "incorrectly_coded_outbound_international_payment"
	ACHTransferReturnReturnReasonCodeOther                                                       ACHTransferReturnReturnReasonCode = "other"
	ACHTransferReturnReturnReasonCodeAccountSoldToAnotherDfi                                     ACHTransferReturnReturnReasonCode = "account_sold_to_another_dfi"
	ACHTransferReturnReturnReasonCodeAddendaError                                                ACHTransferReturnReturnReasonCode = "addenda_error"
	ACHTransferReturnReturnReasonCodeBeneficiaryOrAccountHolderDeceased                          ACHTransferReturnReturnReasonCode = "beneficiary_or_account_holder_deceased"
	ACHTransferReturnReturnReasonCodeCustomerAdvisedNotWithinAuthorizationTerms                  ACHTransferReturnReturnReasonCode = "customer_advised_not_within_authorization_terms"
	ACHTransferReturnReturnReasonCodeCorrectedReturn                                             ACHTransferReturnReturnReasonCode = "corrected_return"
	ACHTransferReturnReturnReasonCodeDuplicateEntry                                              ACHTransferReturnReturnReasonCode = "duplicate_entry"
	ACHTransferReturnReturnReasonCodeDuplicateReturn                                             ACHTransferReturnReturnReasonCode = "duplicate_return"
	ACHTransferReturnReturnReasonCodeEnrDuplicateEnrollment                                      ACHTransferReturnReturnReasonCode = "enr_duplicate_enrollment"
	ACHTransferReturnReturnReasonCodeEnrInvalidDfiAccountNumber                                  ACHTransferReturnReturnReasonCode = "enr_invalid_dfi_account_number"
	ACHTransferReturnReturnReasonCodeEnrInvalidIndividualIDNumber                                ACHTransferReturnReturnReasonCode = "enr_invalid_individual_id_number"
	ACHTransferReturnReturnReasonCodeEnrInvalidRepresentativePayeeIndicator                      ACHTransferReturnReturnReasonCode = "enr_invalid_representative_payee_indicator"
	ACHTransferReturnReturnReasonCodeEnrInvalidTransactionCode                                   ACHTransferReturnReturnReasonCode = "enr_invalid_transaction_code"
	ACHTransferReturnReturnReasonCodeEnrReturnOfEnrEntry                                         ACHTransferReturnReturnReasonCode = "enr_return_of_enr_entry"
	ACHTransferReturnReturnReasonCodeEnrRoutingNumberCheckDigitError                             ACHTransferReturnReturnReasonCode = "enr_routing_number_check_digit_error"
	ACHTransferReturnReturnReasonCodeEntryNotProcessedByGateway                                  ACHTransferReturnReturnReasonCode = "entry_not_processed_by_gateway"
	ACHTransferReturnReturnReasonCodeFieldError                                                  ACHTransferReturnReturnReasonCode = "field_error"
	ACHTransferReturnReturnReasonCodeForeignReceivingDfiUnableToSettle                           ACHTransferReturnReturnReasonCode = "foreign_receiving_dfi_unable_to_settle"
	ACHTransferReturnReturnReasonCodeIatEntryCodingError                                         ACHTransferReturnReturnReasonCode = "iat_entry_coding_error"
	ACHTransferReturnReturnReasonCodeImproperEffectiveEntryDate                                  ACHTransferReturnReturnReasonCode = "improper_effective_entry_date"
	ACHTransferReturnReturnReasonCodeImproperSourceDocumentSourceDocumentPresented               ACHTransferReturnReturnReasonCode = "improper_source_document_source_document_presented"
	ACHTransferReturnReturnReasonCodeInvalidCompanyID                                            ACHTransferReturnReturnReasonCode = "invalid_company_id"
	ACHTransferReturnReturnReasonCodeInvalidForeignReceivingDfiIdentification                    ACHTransferReturnReturnReasonCode = "invalid_foreign_receiving_dfi_identification"
	ACHTransferReturnReturnReasonCodeInvalidIndividualIDNumber                                   ACHTransferReturnReturnReasonCode = "invalid_individual_id_number"
	ACHTransferReturnReturnReasonCodeItemAndRckEntryPresentedForPayment                          ACHTransferReturnReturnReasonCode = "item_and_rck_entry_presented_for_payment"
	ACHTransferReturnReturnReasonCodeItemRelatedToRckEntryIsIneligible                           ACHTransferReturnReturnReasonCode = "item_related_to_rck_entry_is_ineligible"
	ACHTransferReturnReturnReasonCodeMandatoryFieldError                                         ACHTransferReturnReturnReasonCode = "mandatory_field_error"
	ACHTransferReturnReturnReasonCodeMisroutedDishonoredReturn                                   ACHTransferReturnReturnReasonCode = "misrouted_dishonored_return"
	ACHTransferReturnReturnReasonCodeMisroutedReturn                                             ACHTransferReturnReturnReasonCode = "misrouted_return"
	ACHTransferReturnReturnReasonCodeNoErrorsFound                                               ACHTransferReturnReturnReasonCode = "no_errors_found"
	ACHTransferReturnReturnReasonCodeNonAcceptanceOfR62DishonoredReturn                          ACHTransferReturnReturnReasonCode = "non_acceptance_of_r62_dishonored_return"
	ACHTransferReturnReturnReasonCodeNonParticipantInIatProgram                                  ACHTransferReturnReturnReasonCode = "non_participant_in_iat_program"
	ACHTransferReturnReturnReasonCodePermissibleReturnEntry                                      ACHTransferReturnReturnReasonCode = "permissible_return_entry"
	ACHTransferReturnReturnReasonCodePermissibleReturnEntryNotAccepted                           ACHTransferReturnReturnReasonCode = "permissible_return_entry_not_accepted"
	ACHTransferReturnReturnReasonCodeRdfiNonSettlement                                           ACHTransferReturnReturnReasonCode = "rdfi_non_settlement"
	ACHTransferReturnReturnReasonCodeRdfiParticipantInCheckTruncationProgram                     ACHTransferReturnReturnReasonCode = "rdfi_participant_in_check_truncation_program"
	ACHTransferReturnReturnReasonCodeRepresentativePayeeDeceasedOrUnableToContinueInThatCapacity ACHTransferReturnReturnReasonCode = "representative_payee_deceased_or_unable_to_continue_in_that_capacity"
	ACHTransferReturnReturnReasonCodeReturnNotADuplicate                                         ACHTransferReturnReturnReasonCode = "return_not_a_duplicate"
	ACHTransferReturnReturnReasonCodeReturnOfErroneousOrReversingDebit                           ACHTransferReturnReturnReasonCode = "return_of_erroneous_or_reversing_debit"
	ACHTransferReturnReturnReasonCodeReturnOfImproperCreditEntry                                 ACHTransferReturnReturnReasonCode = "return_of_improper_credit_entry"
	ACHTransferReturnReturnReasonCodeReturnOfImproperDebitEntry                                  ACHTransferReturnReturnReasonCode = "return_of_improper_debit_entry"
	ACHTransferReturnReturnReasonCodeReturnOfXckEntry                                            ACHTransferReturnReturnReasonCode = "return_of_xck_entry"
	ACHTransferReturnReturnReasonCodeSourceDocumentPresentedForPayment                           ACHTransferReturnReturnReasonCode = "source_document_presented_for_payment"
	ACHTransferReturnReturnReasonCodeStateLawAffectingRckAcceptance                              ACHTransferReturnReturnReasonCode = "state_law_affecting_rck_acceptance"
	ACHTransferReturnReturnReasonCodeStopPaymentOnItemRelatedToRckEntry                          ACHTransferReturnReturnReasonCode = "stop_payment_on_item_related_to_rck_entry"
	ACHTransferReturnReturnReasonCodeStopPaymentOnSourceDocument                                 ACHTransferReturnReturnReasonCode = "stop_payment_on_source_document"
	ACHTransferReturnReturnReasonCodeTimelyOriginalReturn                                        ACHTransferReturnReturnReasonCode = "timely_original_return"
	ACHTransferReturnReturnReasonCodeTraceNumberError                                            ACHTransferReturnReturnReasonCode = "trace_number_error"
	ACHTransferReturnReturnReasonCodeUntimelyDishonoredReturn                                    ACHTransferReturnReturnReasonCode = "untimely_dishonored_return"
	ACHTransferReturnReturnReasonCodeUntimelyReturn                                              ACHTransferReturnReturnReasonCode = "untimely_return"
)

// The lifecycle status of the transfer.
type ACHTransferStatus string

const (
	ACHTransferStatusPendingApproval   ACHTransferStatus = "pending_approval"
	ACHTransferStatusCanceled          ACHTransferStatus = "canceled"
	ACHTransferStatusPendingReviewing  ACHTransferStatus = "pending_reviewing"
	ACHTransferStatusPendingSubmission ACHTransferStatus = "pending_submission"
	ACHTransferStatusSubmitted         ACHTransferStatus = "submitted"
	ACHTransferStatusReturned          ACHTransferStatus = "returned"
	ACHTransferStatusRequiresAttention ACHTransferStatus = "requires_attention"
	ACHTransferStatusRejected          ACHTransferStatus = "rejected"
)

// After the transfer is submitted to FedACH, this will contain supplemental
// details.
type ACHTransferSubmission struct {
	// The trace number for the submission.
	TraceNumber string `json:"trace_number,required"`
	// When the ACH transfer was sent to FedACH.
	SubmittedAt time.Time `json:"submitted_at,required" format:"date-time"`
	JSON        achTransferSubmissionJSON
}

// achTransferSubmissionJSON contains the JSON metadata for the struct
// [ACHTransferSubmission]
type achTransferSubmissionJSON struct {
	TraceNumber apijson.Field
	SubmittedAt apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACHTransferSubmission) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the account to which the transfer will be sent.
type ACHTransferFunding string

const (
	ACHTransferFundingChecking ACHTransferFunding = "checking"
	ACHTransferFundingSavings  ACHTransferFunding = "savings"
)

// The Standard Entry Class (SEC) code to use for the transfer.
type ACHTransferStandardEntryClassCode string

const (
	ACHTransferStandardEntryClassCodeCorporateCreditOrDebit        ACHTransferStandardEntryClassCode = "corporate_credit_or_debit"
	ACHTransferStandardEntryClassCodePrearrangedPaymentsAndDeposit ACHTransferStandardEntryClassCode = "prearranged_payments_and_deposit"
	ACHTransferStandardEntryClassCodeInternetInitiated             ACHTransferStandardEntryClassCode = "internet_initiated"
)

// A constant representing the object's type. For this resource it will always be
// `ach_transfer`.
type ACHTransferType string

const (
	ACHTransferTypeACHTransfer ACHTransferType = "ach_transfer"
)

type ACHTransferNewParams struct {
	// The Increase identifier for the account that will send the transfer.
	AccountID param.Field[string] `json:"account_id,required"`
	// The transfer amount in cents. A positive amount originates a credit transfer
	// pushing funds to the receiving account. A negative amount originates a debit
	// transfer pulling funds from the receiving account.
	Amount param.Field[int64] `json:"amount,required"`
	// A description you choose to give the transfer. This will be saved with the
	// transfer details, displayed in the dashboard, and returned by the API. If
	// `individual_name` and `company_name` are not explicitly set by this API, the
	// `statement_descriptor` will be sent in those fields to the receiving bank to
	// help the customer recognize the transfer. You are highly encouraged to pass
	// `individual_name` and `company_name` instead of relying on this fallback.
	StatementDescriptor param.Field[string] `json:"statement_descriptor,required"`
	// The account number for the destination account.
	AccountNumber param.Field[string] `json:"account_number"`
	// Additional information that will be sent to the recipient. This is included in
	// the transfer data sent to the receiving bank.
	Addendum param.Field[string] `json:"addendum"`
	// The description of the date of the transfer, usually in the format `YYMMDD`.
	// This is included in the transfer data sent to the receiving bank.
	CompanyDescriptiveDate param.Field[string] `json:"company_descriptive_date"`
	// The data you choose to associate with the transfer. This is included in the
	// transfer data sent to the receiving bank.
	CompanyDiscretionaryData param.Field[string] `json:"company_discretionary_data"`
	// A description of the transfer. This is included in the transfer data sent to the
	// receiving bank.
	CompanyEntryDescription param.Field[string] `json:"company_entry_description"`
	// The name by which the recipient knows you. This is included in the transfer data
	// sent to the receiving bank.
	CompanyName param.Field[string] `json:"company_name"`
	// The transfer effective date in
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	EffectiveDate param.Field[time.Time] `json:"effective_date" format:"date"`
	// The ID of an External Account to initiate a transfer to. If this parameter is
	// provided, `account_number`, `routing_number`, and `funding` must be absent.
	ExternalAccountID param.Field[string] `json:"external_account_id"`
	// The type of the account to which the transfer will be sent.
	Funding param.Field[ACHTransferNewParamsFunding] `json:"funding"`
	// Your identifer for the transfer recipient.
	IndividualID param.Field[string] `json:"individual_id"`
	// The name of the transfer recipient. This value is informational and not verified
	// by the recipient's bank.
	IndividualName param.Field[string] `json:"individual_name"`
	// Whether the transfer requires explicit approval via the dashboard or API.
	RequireApproval param.Field[bool] `json:"require_approval"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
	// destination account.
	RoutingNumber param.Field[string] `json:"routing_number"`
	// The Standard Entry Class (SEC) code to use for the transfer.
	StandardEntryClassCode param.Field[ACHTransferNewParamsStandardEntryClassCode] `json:"standard_entry_class_code"`
}

func (r ACHTransferNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of the account to which the transfer will be sent.
type ACHTransferNewParamsFunding string

const (
	ACHTransferNewParamsFundingChecking ACHTransferNewParamsFunding = "checking"
	ACHTransferNewParamsFundingSavings  ACHTransferNewParamsFunding = "savings"
)

// The Standard Entry Class (SEC) code to use for the transfer.
type ACHTransferNewParamsStandardEntryClassCode string

const (
	ACHTransferNewParamsStandardEntryClassCodeCorporateCreditOrDebit        ACHTransferNewParamsStandardEntryClassCode = "corporate_credit_or_debit"
	ACHTransferNewParamsStandardEntryClassCodePrearrangedPaymentsAndDeposit ACHTransferNewParamsStandardEntryClassCode = "prearranged_payments_and_deposit"
	ACHTransferNewParamsStandardEntryClassCodeInternetInitiated             ACHTransferNewParamsStandardEntryClassCode = "internet_initiated"
)

type ACHTransferListParams struct {
	// Filter ACH Transfers to those that originated from the specified Account.
	AccountID param.Field[string]                         `query:"account_id"`
	CreatedAt param.Field[ACHTransferListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Filter ACH Transfers to those made to the specified External Account.
	ExternalAccountID param.Field[string] `query:"external_account_id"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [ACHTransferListParams]'s query parameters as `url.Values`.
func (r ACHTransferListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ACHTransferListParamsCreatedAt struct {
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

// URLQuery serializes [ACHTransferListParamsCreatedAt]'s query parameters as
// `url.Values`.
func (r ACHTransferListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// A list of ACH Transfer objects
type ACHTransferListResponse struct {
	// The contents of the list.
	Data []ACHTransfer `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       achTransferListResponseJSON
}

// achTransferListResponseJSON contains the JSON metadata for the struct
// [ACHTransferListResponse]
type achTransferListResponseJSON struct {
	Data        apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACHTransferListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

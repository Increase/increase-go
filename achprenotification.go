// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Increase/increase-go/internal/apijson"
	"github.com/Increase/increase-go/internal/apiquery"
	"github.com/Increase/increase-go/internal/param"
	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
)

// ACHPrenotificationService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewACHPrenotificationService] method instead.
type ACHPrenotificationService struct {
	Options []option.RequestOption
}

// NewACHPrenotificationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewACHPrenotificationService(opts ...option.RequestOption) (r *ACHPrenotificationService) {
	r = &ACHPrenotificationService{}
	r.Options = opts
	return
}

// Create an ACH Prenotification
func (r *ACHPrenotificationService) New(ctx context.Context, body ACHPrenotificationNewParams, opts ...option.RequestOption) (res *ACHPrenotification, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ach_prenotifications"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve an ACH Prenotification
func (r *ACHPrenotificationService) Get(ctx context.Context, achPrenotificationID string, opts ...option.RequestOption) (res *ACHPrenotification, err error) {
	opts = slices.Concat(r.Options, opts)
	if achPrenotificationID == "" {
		err = errors.New("missing required ach_prenotification_id parameter")
		return
	}
	path := fmt.Sprintf("ach_prenotifications/%s", achPrenotificationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List ACH Prenotifications
func (r *ACHPrenotificationService) List(ctx context.Context, query ACHPrenotificationListParams, opts ...option.RequestOption) (res *ACHPrenotificationListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ach_prenotifications"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// ACH Prenotifications are one way you can verify account and routing numbers by
// Automated Clearing House (ACH).
type ACHPrenotification struct {
	// The ACH Prenotification's identifier.
	ID string `json:"id,required"`
	// The account that sent the ACH Prenotification.
	AccountID string `json:"account_id,required,nullable"`
	// The destination account number.
	AccountNumber string `json:"account_number,required"`
	// Additional information for the recipient.
	Addendum string `json:"addendum,required,nullable"`
	// The description of the date of the notification.
	CompanyDescriptiveDate string `json:"company_descriptive_date,required,nullable"`
	// Optional data associated with the notification.
	CompanyDiscretionaryData string `json:"company_discretionary_data,required,nullable"`
	// The description of the notification.
	CompanyEntryDescription string `json:"company_entry_description,required,nullable"`
	// The name by which you know the company.
	CompanyName string `json:"company_name,required,nullable"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the prenotification was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// If the notification is for a future credit or debit.
	CreditDebitIndicator ACHPrenotificationCreditDebitIndicator `json:"credit_debit_indicator,required,nullable"`
	// The effective date in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	EffectiveDate time.Time `json:"effective_date,required,nullable" format:"date-time"`
	// The idempotency key you chose for this object. This value is unique across
	// Increase and is used to ensure that a request is only processed once. Learn more
	// about [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey string `json:"idempotency_key,required,nullable"`
	// Your identifier for the recipient.
	IndividualID string `json:"individual_id,required,nullable"`
	// The name of the recipient. This value is informational and not verified by the
	// recipient's bank.
	IndividualName string `json:"individual_name,required,nullable"`
	// If the receiving bank notifies that future transfers should use different
	// details, this will contain those details.
	NotificationsOfChange []ACHPrenotificationNotificationsOfChange `json:"notifications_of_change,required"`
	// If your prenotification is returned, this will contain details of the return.
	PrenotificationReturn ACHPrenotificationPrenotificationReturn `json:"prenotification_return,required,nullable"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber string `json:"routing_number,required"`
	// The Standard Entry Class (SEC) code to use for the ACH Prenotification.
	StandardEntryClassCode ACHPrenotificationStandardEntryClassCode `json:"standard_entry_class_code,required,nullable"`
	// The lifecycle status of the ACH Prenotification.
	Status ACHPrenotificationStatus `json:"status,required"`
	// A constant representing the object's type. For this resource it will always be
	// `ach_prenotification`.
	Type ACHPrenotificationType `json:"type,required"`
	JSON achPrenotificationJSON `json:"-"`
}

// achPrenotificationJSON contains the JSON metadata for the struct
// [ACHPrenotification]
type achPrenotificationJSON struct {
	ID                       apijson.Field
	AccountID                apijson.Field
	AccountNumber            apijson.Field
	Addendum                 apijson.Field
	CompanyDescriptiveDate   apijson.Field
	CompanyDiscretionaryData apijson.Field
	CompanyEntryDescription  apijson.Field
	CompanyName              apijson.Field
	CreatedAt                apijson.Field
	CreditDebitIndicator     apijson.Field
	EffectiveDate            apijson.Field
	IdempotencyKey           apijson.Field
	IndividualID             apijson.Field
	IndividualName           apijson.Field
	NotificationsOfChange    apijson.Field
	PrenotificationReturn    apijson.Field
	RoutingNumber            apijson.Field
	StandardEntryClassCode   apijson.Field
	Status                   apijson.Field
	Type                     apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *ACHPrenotification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r achPrenotificationJSON) RawJSON() string {
	return r.raw
}

// If the notification is for a future credit or debit.
type ACHPrenotificationCreditDebitIndicator string

const (
	ACHPrenotificationCreditDebitIndicatorCredit ACHPrenotificationCreditDebitIndicator = "credit"
	ACHPrenotificationCreditDebitIndicatorDebit  ACHPrenotificationCreditDebitIndicator = "debit"
)

func (r ACHPrenotificationCreditDebitIndicator) IsKnown() bool {
	switch r {
	case ACHPrenotificationCreditDebitIndicatorCredit, ACHPrenotificationCreditDebitIndicatorDebit:
		return true
	}
	return false
}

type ACHPrenotificationNotificationsOfChange struct {
	// The required type of change that is being signaled by the receiving financial
	// institution.
	ChangeCode ACHPrenotificationNotificationsOfChangeChangeCode `json:"change_code,required"`
	// The corrected data that should be used in future ACHs to this account. This may
	// contain the suggested new account number or routing number. When the
	// `change_code` is `incorrect_transaction_code`, this field contains an integer.
	// Numbers starting with a 2 encourage changing the `funding` parameter to
	// checking; numbers starting with a 3 encourage changing to savings.
	CorrectedData string `json:"corrected_data,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the notification occurred.
	CreatedAt time.Time                                   `json:"created_at,required" format:"date-time"`
	JSON      achPrenotificationNotificationsOfChangeJSON `json:"-"`
}

// achPrenotificationNotificationsOfChangeJSON contains the JSON metadata for the
// struct [ACHPrenotificationNotificationsOfChange]
type achPrenotificationNotificationsOfChangeJSON struct {
	ChangeCode    apijson.Field
	CorrectedData apijson.Field
	CreatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ACHPrenotificationNotificationsOfChange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r achPrenotificationNotificationsOfChangeJSON) RawJSON() string {
	return r.raw
}

// The required type of change that is being signaled by the receiving financial
// institution.
type ACHPrenotificationNotificationsOfChangeChangeCode string

const (
	ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectAccountNumber                                                      ACHPrenotificationNotificationsOfChangeChangeCode = "incorrect_account_number"
	ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectRoutingNumber                                                      ACHPrenotificationNotificationsOfChangeChangeCode = "incorrect_routing_number"
	ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectRoutingNumberAndAccountNumber                                      ACHPrenotificationNotificationsOfChangeChangeCode = "incorrect_routing_number_and_account_number"
	ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectTransactionCode                                                    ACHPrenotificationNotificationsOfChangeChangeCode = "incorrect_transaction_code"
	ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectAccountNumberAndTransactionCode                                    ACHPrenotificationNotificationsOfChangeChangeCode = "incorrect_account_number_and_transaction_code"
	ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectRoutingNumberAccountNumberAndTransactionCode                       ACHPrenotificationNotificationsOfChangeChangeCode = "incorrect_routing_number_account_number_and_transaction_code"
	ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectReceivingDepositoryFinancialInstitutionIdentification              ACHPrenotificationNotificationsOfChangeChangeCode = "incorrect_receiving_depository_financial_institution_identification"
	ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectIndividualIdentificationNumber                                     ACHPrenotificationNotificationsOfChangeChangeCode = "incorrect_individual_identification_number"
	ACHPrenotificationNotificationsOfChangeChangeCodeAddendaFormatError                                                          ACHPrenotificationNotificationsOfChangeChangeCode = "addenda_format_error"
	ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectStandardEntryClassCodeForOutboundInternationalPayment              ACHPrenotificationNotificationsOfChangeChangeCode = "incorrect_standard_entry_class_code_for_outbound_international_payment"
	ACHPrenotificationNotificationsOfChangeChangeCodeMisroutedNotificationOfChange                                               ACHPrenotificationNotificationsOfChangeChangeCode = "misrouted_notification_of_change"
	ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectTraceNumber                                                        ACHPrenotificationNotificationsOfChangeChangeCode = "incorrect_trace_number"
	ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectCompanyIdentificationNumber                                        ACHPrenotificationNotificationsOfChangeChangeCode = "incorrect_company_identification_number"
	ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectIdentificationNumber                                               ACHPrenotificationNotificationsOfChangeChangeCode = "incorrect_identification_number"
	ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectlyFormattedCorrectedData                                           ACHPrenotificationNotificationsOfChangeChangeCode = "incorrectly_formatted_corrected_data"
	ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectDiscretionaryData                                                  ACHPrenotificationNotificationsOfChangeChangeCode = "incorrect_discretionary_data"
	ACHPrenotificationNotificationsOfChangeChangeCodeRoutingNumberNotFromOriginalEntryDetailRecord                               ACHPrenotificationNotificationsOfChangeChangeCode = "routing_number_not_from_original_entry_detail_record"
	ACHPrenotificationNotificationsOfChangeChangeCodeDepositoryFinancialInstitutionAccountNumberNotFromOriginalEntryDetailRecord ACHPrenotificationNotificationsOfChangeChangeCode = "depository_financial_institution_account_number_not_from_original_entry_detail_record"
	ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectTransactionCodeByOriginatingDepositoryFinancialInstitution         ACHPrenotificationNotificationsOfChangeChangeCode = "incorrect_transaction_code_by_originating_depository_financial_institution"
)

func (r ACHPrenotificationNotificationsOfChangeChangeCode) IsKnown() bool {
	switch r {
	case ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectAccountNumber, ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectRoutingNumber, ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectRoutingNumberAndAccountNumber, ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectTransactionCode, ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectAccountNumberAndTransactionCode, ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectRoutingNumberAccountNumberAndTransactionCode, ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectReceivingDepositoryFinancialInstitutionIdentification, ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectIndividualIdentificationNumber, ACHPrenotificationNotificationsOfChangeChangeCodeAddendaFormatError, ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectStandardEntryClassCodeForOutboundInternationalPayment, ACHPrenotificationNotificationsOfChangeChangeCodeMisroutedNotificationOfChange, ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectTraceNumber, ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectCompanyIdentificationNumber, ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectIdentificationNumber, ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectlyFormattedCorrectedData, ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectDiscretionaryData, ACHPrenotificationNotificationsOfChangeChangeCodeRoutingNumberNotFromOriginalEntryDetailRecord, ACHPrenotificationNotificationsOfChangeChangeCodeDepositoryFinancialInstitutionAccountNumberNotFromOriginalEntryDetailRecord, ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectTransactionCodeByOriginatingDepositoryFinancialInstitution:
		return true
	}
	return false
}

// If your prenotification is returned, this will contain details of the return.
type ACHPrenotificationPrenotificationReturn struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Prenotification was returned.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Why the Prenotification was returned.
	ReturnReasonCode ACHPrenotificationPrenotificationReturnReturnReasonCode `json:"return_reason_code,required"`
	JSON             achPrenotificationPrenotificationReturnJSON             `json:"-"`
}

// achPrenotificationPrenotificationReturnJSON contains the JSON metadata for the
// struct [ACHPrenotificationPrenotificationReturn]
type achPrenotificationPrenotificationReturnJSON struct {
	CreatedAt        apijson.Field
	ReturnReasonCode apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ACHPrenotificationPrenotificationReturn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r achPrenotificationPrenotificationReturnJSON) RawJSON() string {
	return r.raw
}

// Why the Prenotification was returned.
type ACHPrenotificationPrenotificationReturnReturnReasonCode string

const (
	ACHPrenotificationPrenotificationReturnReturnReasonCodeInsufficientFund                                            ACHPrenotificationPrenotificationReturnReturnReasonCode = "insufficient_fund"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeNoAccount                                                   ACHPrenotificationPrenotificationReturnReturnReasonCode = "no_account"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeAccountClosed                                               ACHPrenotificationPrenotificationReturnReturnReasonCode = "account_closed"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeInvalidAccountNumberStructure                               ACHPrenotificationPrenotificationReturnReturnReasonCode = "invalid_account_number_structure"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeAccountFrozenEntryReturnedPerOfacInstruction                ACHPrenotificationPrenotificationReturnReturnReasonCode = "account_frozen_entry_returned_per_ofac_instruction"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeCreditEntryRefusedByReceiver                                ACHPrenotificationPrenotificationReturnReturnReasonCode = "credit_entry_refused_by_receiver"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeUnauthorizedDebitToConsumerAccountUsingCorporateSecCode     ACHPrenotificationPrenotificationReturnReturnReasonCode = "unauthorized_debit_to_consumer_account_using_corporate_sec_code"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeCorporateCustomerAdvisedNotAuthorized                       ACHPrenotificationPrenotificationReturnReturnReasonCode = "corporate_customer_advised_not_authorized"
	ACHPrenotificationPrenotificationReturnReturnReasonCodePaymentStopped                                              ACHPrenotificationPrenotificationReturnReturnReasonCode = "payment_stopped"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeNonTransactionAccount                                       ACHPrenotificationPrenotificationReturnReturnReasonCode = "non_transaction_account"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeUncollectedFunds                                            ACHPrenotificationPrenotificationReturnReturnReasonCode = "uncollected_funds"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeRoutingNumberCheckDigitError                                ACHPrenotificationPrenotificationReturnReturnReasonCode = "routing_number_check_digit_error"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete   ACHPrenotificationPrenotificationReturnReturnReasonCode = "customer_advised_unauthorized_improper_ineligible_or_incomplete"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeAmountFieldError                                            ACHPrenotificationPrenotificationReturnReturnReasonCode = "amount_field_error"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeAuthorizationRevokedByCustomer                              ACHPrenotificationPrenotificationReturnReturnReasonCode = "authorization_revoked_by_customer"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeInvalidACHRoutingNumber                                     ACHPrenotificationPrenotificationReturnReturnReasonCode = "invalid_ach_routing_number"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeFileRecordEditCriteria                                      ACHPrenotificationPrenotificationReturnReturnReasonCode = "file_record_edit_criteria"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeEnrInvalidIndividualName                                    ACHPrenotificationPrenotificationReturnReturnReasonCode = "enr_invalid_individual_name"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeReturnedPerOdfiRequest                                      ACHPrenotificationPrenotificationReturnReturnReasonCode = "returned_per_odfi_request"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeLimitedParticipationDfi                                     ACHPrenotificationPrenotificationReturnReturnReasonCode = "limited_participation_dfi"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeIncorrectlyCodedOutboundInternationalPayment                ACHPrenotificationPrenotificationReturnReturnReasonCode = "incorrectly_coded_outbound_international_payment"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeAccountSoldToAnotherDfi                                     ACHPrenotificationPrenotificationReturnReturnReasonCode = "account_sold_to_another_dfi"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeAddendaError                                                ACHPrenotificationPrenotificationReturnReturnReasonCode = "addenda_error"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeBeneficiaryOrAccountHolderDeceased                          ACHPrenotificationPrenotificationReturnReturnReasonCode = "beneficiary_or_account_holder_deceased"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeCustomerAdvisedNotWithinAuthorizationTerms                  ACHPrenotificationPrenotificationReturnReturnReasonCode = "customer_advised_not_within_authorization_terms"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeCorrectedReturn                                             ACHPrenotificationPrenotificationReturnReturnReasonCode = "corrected_return"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeDuplicateEntry                                              ACHPrenotificationPrenotificationReturnReturnReasonCode = "duplicate_entry"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeDuplicateReturn                                             ACHPrenotificationPrenotificationReturnReturnReasonCode = "duplicate_return"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeEnrDuplicateEnrollment                                      ACHPrenotificationPrenotificationReturnReturnReasonCode = "enr_duplicate_enrollment"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeEnrInvalidDfiAccountNumber                                  ACHPrenotificationPrenotificationReturnReturnReasonCode = "enr_invalid_dfi_account_number"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeEnrInvalidIndividualIDNumber                                ACHPrenotificationPrenotificationReturnReturnReasonCode = "enr_invalid_individual_id_number"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeEnrInvalidRepresentativePayeeIndicator                      ACHPrenotificationPrenotificationReturnReturnReasonCode = "enr_invalid_representative_payee_indicator"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeEnrInvalidTransactionCode                                   ACHPrenotificationPrenotificationReturnReturnReasonCode = "enr_invalid_transaction_code"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeEnrReturnOfEnrEntry                                         ACHPrenotificationPrenotificationReturnReturnReasonCode = "enr_return_of_enr_entry"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeEnrRoutingNumberCheckDigitError                             ACHPrenotificationPrenotificationReturnReturnReasonCode = "enr_routing_number_check_digit_error"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeEntryNotProcessedByGateway                                  ACHPrenotificationPrenotificationReturnReturnReasonCode = "entry_not_processed_by_gateway"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeFieldError                                                  ACHPrenotificationPrenotificationReturnReturnReasonCode = "field_error"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeForeignReceivingDfiUnableToSettle                           ACHPrenotificationPrenotificationReturnReturnReasonCode = "foreign_receiving_dfi_unable_to_settle"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeIatEntryCodingError                                         ACHPrenotificationPrenotificationReturnReturnReasonCode = "iat_entry_coding_error"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeImproperEffectiveEntryDate                                  ACHPrenotificationPrenotificationReturnReturnReasonCode = "improper_effective_entry_date"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeImproperSourceDocumentSourceDocumentPresented               ACHPrenotificationPrenotificationReturnReturnReasonCode = "improper_source_document_source_document_presented"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeInvalidCompanyID                                            ACHPrenotificationPrenotificationReturnReturnReasonCode = "invalid_company_id"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeInvalidForeignReceivingDfiIdentification                    ACHPrenotificationPrenotificationReturnReturnReasonCode = "invalid_foreign_receiving_dfi_identification"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeInvalidIndividualIDNumber                                   ACHPrenotificationPrenotificationReturnReturnReasonCode = "invalid_individual_id_number"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeItemAndRckEntryPresentedForPayment                          ACHPrenotificationPrenotificationReturnReturnReasonCode = "item_and_rck_entry_presented_for_payment"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeItemRelatedToRckEntryIsIneligible                           ACHPrenotificationPrenotificationReturnReturnReasonCode = "item_related_to_rck_entry_is_ineligible"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeMandatoryFieldError                                         ACHPrenotificationPrenotificationReturnReturnReasonCode = "mandatory_field_error"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeMisroutedDishonoredReturn                                   ACHPrenotificationPrenotificationReturnReturnReasonCode = "misrouted_dishonored_return"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeMisroutedReturn                                             ACHPrenotificationPrenotificationReturnReturnReasonCode = "misrouted_return"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeNoErrorsFound                                               ACHPrenotificationPrenotificationReturnReturnReasonCode = "no_errors_found"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeNonAcceptanceOfR62DishonoredReturn                          ACHPrenotificationPrenotificationReturnReturnReasonCode = "non_acceptance_of_r62_dishonored_return"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeNonParticipantInIatProgram                                  ACHPrenotificationPrenotificationReturnReturnReasonCode = "non_participant_in_iat_program"
	ACHPrenotificationPrenotificationReturnReturnReasonCodePermissibleReturnEntry                                      ACHPrenotificationPrenotificationReturnReturnReasonCode = "permissible_return_entry"
	ACHPrenotificationPrenotificationReturnReturnReasonCodePermissibleReturnEntryNotAccepted                           ACHPrenotificationPrenotificationReturnReturnReasonCode = "permissible_return_entry_not_accepted"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeRdfiNonSettlement                                           ACHPrenotificationPrenotificationReturnReturnReasonCode = "rdfi_non_settlement"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeRdfiParticipantInCheckTruncationProgram                     ACHPrenotificationPrenotificationReturnReturnReasonCode = "rdfi_participant_in_check_truncation_program"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeRepresentativePayeeDeceasedOrUnableToContinueInThatCapacity ACHPrenotificationPrenotificationReturnReturnReasonCode = "representative_payee_deceased_or_unable_to_continue_in_that_capacity"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeReturnNotADuplicate                                         ACHPrenotificationPrenotificationReturnReturnReasonCode = "return_not_a_duplicate"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeReturnOfErroneousOrReversingDebit                           ACHPrenotificationPrenotificationReturnReturnReasonCode = "return_of_erroneous_or_reversing_debit"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeReturnOfImproperCreditEntry                                 ACHPrenotificationPrenotificationReturnReturnReasonCode = "return_of_improper_credit_entry"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeReturnOfImproperDebitEntry                                  ACHPrenotificationPrenotificationReturnReturnReasonCode = "return_of_improper_debit_entry"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeReturnOfXckEntry                                            ACHPrenotificationPrenotificationReturnReturnReasonCode = "return_of_xck_entry"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeSourceDocumentPresentedForPayment                           ACHPrenotificationPrenotificationReturnReturnReasonCode = "source_document_presented_for_payment"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeStateLawAffectingRckAcceptance                              ACHPrenotificationPrenotificationReturnReturnReasonCode = "state_law_affecting_rck_acceptance"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeStopPaymentOnItemRelatedToRckEntry                          ACHPrenotificationPrenotificationReturnReturnReasonCode = "stop_payment_on_item_related_to_rck_entry"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeStopPaymentOnSourceDocument                                 ACHPrenotificationPrenotificationReturnReturnReasonCode = "stop_payment_on_source_document"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeTimelyOriginalReturn                                        ACHPrenotificationPrenotificationReturnReturnReasonCode = "timely_original_return"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeTraceNumberError                                            ACHPrenotificationPrenotificationReturnReturnReasonCode = "trace_number_error"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeUntimelyDishonoredReturn                                    ACHPrenotificationPrenotificationReturnReturnReasonCode = "untimely_dishonored_return"
	ACHPrenotificationPrenotificationReturnReturnReasonCodeUntimelyReturn                                              ACHPrenotificationPrenotificationReturnReturnReasonCode = "untimely_return"
)

func (r ACHPrenotificationPrenotificationReturnReturnReasonCode) IsKnown() bool {
	switch r {
	case ACHPrenotificationPrenotificationReturnReturnReasonCodeInsufficientFund, ACHPrenotificationPrenotificationReturnReturnReasonCodeNoAccount, ACHPrenotificationPrenotificationReturnReturnReasonCodeAccountClosed, ACHPrenotificationPrenotificationReturnReturnReasonCodeInvalidAccountNumberStructure, ACHPrenotificationPrenotificationReturnReturnReasonCodeAccountFrozenEntryReturnedPerOfacInstruction, ACHPrenotificationPrenotificationReturnReturnReasonCodeCreditEntryRefusedByReceiver, ACHPrenotificationPrenotificationReturnReturnReasonCodeUnauthorizedDebitToConsumerAccountUsingCorporateSecCode, ACHPrenotificationPrenotificationReturnReturnReasonCodeCorporateCustomerAdvisedNotAuthorized, ACHPrenotificationPrenotificationReturnReturnReasonCodePaymentStopped, ACHPrenotificationPrenotificationReturnReturnReasonCodeNonTransactionAccount, ACHPrenotificationPrenotificationReturnReturnReasonCodeUncollectedFunds, ACHPrenotificationPrenotificationReturnReturnReasonCodeRoutingNumberCheckDigitError, ACHPrenotificationPrenotificationReturnReturnReasonCodeCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete, ACHPrenotificationPrenotificationReturnReturnReasonCodeAmountFieldError, ACHPrenotificationPrenotificationReturnReturnReasonCodeAuthorizationRevokedByCustomer, ACHPrenotificationPrenotificationReturnReturnReasonCodeInvalidACHRoutingNumber, ACHPrenotificationPrenotificationReturnReturnReasonCodeFileRecordEditCriteria, ACHPrenotificationPrenotificationReturnReturnReasonCodeEnrInvalidIndividualName, ACHPrenotificationPrenotificationReturnReturnReasonCodeReturnedPerOdfiRequest, ACHPrenotificationPrenotificationReturnReturnReasonCodeLimitedParticipationDfi, ACHPrenotificationPrenotificationReturnReturnReasonCodeIncorrectlyCodedOutboundInternationalPayment, ACHPrenotificationPrenotificationReturnReturnReasonCodeAccountSoldToAnotherDfi, ACHPrenotificationPrenotificationReturnReturnReasonCodeAddendaError, ACHPrenotificationPrenotificationReturnReturnReasonCodeBeneficiaryOrAccountHolderDeceased, ACHPrenotificationPrenotificationReturnReturnReasonCodeCustomerAdvisedNotWithinAuthorizationTerms, ACHPrenotificationPrenotificationReturnReturnReasonCodeCorrectedReturn, ACHPrenotificationPrenotificationReturnReturnReasonCodeDuplicateEntry, ACHPrenotificationPrenotificationReturnReturnReasonCodeDuplicateReturn, ACHPrenotificationPrenotificationReturnReturnReasonCodeEnrDuplicateEnrollment, ACHPrenotificationPrenotificationReturnReturnReasonCodeEnrInvalidDfiAccountNumber, ACHPrenotificationPrenotificationReturnReturnReasonCodeEnrInvalidIndividualIDNumber, ACHPrenotificationPrenotificationReturnReturnReasonCodeEnrInvalidRepresentativePayeeIndicator, ACHPrenotificationPrenotificationReturnReturnReasonCodeEnrInvalidTransactionCode, ACHPrenotificationPrenotificationReturnReturnReasonCodeEnrReturnOfEnrEntry, ACHPrenotificationPrenotificationReturnReturnReasonCodeEnrRoutingNumberCheckDigitError, ACHPrenotificationPrenotificationReturnReturnReasonCodeEntryNotProcessedByGateway, ACHPrenotificationPrenotificationReturnReturnReasonCodeFieldError, ACHPrenotificationPrenotificationReturnReturnReasonCodeForeignReceivingDfiUnableToSettle, ACHPrenotificationPrenotificationReturnReturnReasonCodeIatEntryCodingError, ACHPrenotificationPrenotificationReturnReturnReasonCodeImproperEffectiveEntryDate, ACHPrenotificationPrenotificationReturnReturnReasonCodeImproperSourceDocumentSourceDocumentPresented, ACHPrenotificationPrenotificationReturnReturnReasonCodeInvalidCompanyID, ACHPrenotificationPrenotificationReturnReturnReasonCodeInvalidForeignReceivingDfiIdentification, ACHPrenotificationPrenotificationReturnReturnReasonCodeInvalidIndividualIDNumber, ACHPrenotificationPrenotificationReturnReturnReasonCodeItemAndRckEntryPresentedForPayment, ACHPrenotificationPrenotificationReturnReturnReasonCodeItemRelatedToRckEntryIsIneligible, ACHPrenotificationPrenotificationReturnReturnReasonCodeMandatoryFieldError, ACHPrenotificationPrenotificationReturnReturnReasonCodeMisroutedDishonoredReturn, ACHPrenotificationPrenotificationReturnReturnReasonCodeMisroutedReturn, ACHPrenotificationPrenotificationReturnReturnReasonCodeNoErrorsFound, ACHPrenotificationPrenotificationReturnReturnReasonCodeNonAcceptanceOfR62DishonoredReturn, ACHPrenotificationPrenotificationReturnReturnReasonCodeNonParticipantInIatProgram, ACHPrenotificationPrenotificationReturnReturnReasonCodePermissibleReturnEntry, ACHPrenotificationPrenotificationReturnReturnReasonCodePermissibleReturnEntryNotAccepted, ACHPrenotificationPrenotificationReturnReturnReasonCodeRdfiNonSettlement, ACHPrenotificationPrenotificationReturnReturnReasonCodeRdfiParticipantInCheckTruncationProgram, ACHPrenotificationPrenotificationReturnReturnReasonCodeRepresentativePayeeDeceasedOrUnableToContinueInThatCapacity, ACHPrenotificationPrenotificationReturnReturnReasonCodeReturnNotADuplicate, ACHPrenotificationPrenotificationReturnReturnReasonCodeReturnOfErroneousOrReversingDebit, ACHPrenotificationPrenotificationReturnReturnReasonCodeReturnOfImproperCreditEntry, ACHPrenotificationPrenotificationReturnReturnReasonCodeReturnOfImproperDebitEntry, ACHPrenotificationPrenotificationReturnReturnReasonCodeReturnOfXckEntry, ACHPrenotificationPrenotificationReturnReturnReasonCodeSourceDocumentPresentedForPayment, ACHPrenotificationPrenotificationReturnReturnReasonCodeStateLawAffectingRckAcceptance, ACHPrenotificationPrenotificationReturnReturnReasonCodeStopPaymentOnItemRelatedToRckEntry, ACHPrenotificationPrenotificationReturnReturnReasonCodeStopPaymentOnSourceDocument, ACHPrenotificationPrenotificationReturnReturnReasonCodeTimelyOriginalReturn, ACHPrenotificationPrenotificationReturnReturnReasonCodeTraceNumberError, ACHPrenotificationPrenotificationReturnReturnReasonCodeUntimelyDishonoredReturn, ACHPrenotificationPrenotificationReturnReturnReasonCodeUntimelyReturn:
		return true
	}
	return false
}

// The Standard Entry Class (SEC) code to use for the ACH Prenotification.
type ACHPrenotificationStandardEntryClassCode string

const (
	ACHPrenotificationStandardEntryClassCodeCorporateCreditOrDebit        ACHPrenotificationStandardEntryClassCode = "corporate_credit_or_debit"
	ACHPrenotificationStandardEntryClassCodeCorporateTradeExchange        ACHPrenotificationStandardEntryClassCode = "corporate_trade_exchange"
	ACHPrenotificationStandardEntryClassCodePrearrangedPaymentsAndDeposit ACHPrenotificationStandardEntryClassCode = "prearranged_payments_and_deposit"
	ACHPrenotificationStandardEntryClassCodeInternetInitiated             ACHPrenotificationStandardEntryClassCode = "internet_initiated"
)

func (r ACHPrenotificationStandardEntryClassCode) IsKnown() bool {
	switch r {
	case ACHPrenotificationStandardEntryClassCodeCorporateCreditOrDebit, ACHPrenotificationStandardEntryClassCodeCorporateTradeExchange, ACHPrenotificationStandardEntryClassCodePrearrangedPaymentsAndDeposit, ACHPrenotificationStandardEntryClassCodeInternetInitiated:
		return true
	}
	return false
}

// The lifecycle status of the ACH Prenotification.
type ACHPrenotificationStatus string

const (
	ACHPrenotificationStatusPendingSubmitting ACHPrenotificationStatus = "pending_submitting"
	ACHPrenotificationStatusRequiresAttention ACHPrenotificationStatus = "requires_attention"
	ACHPrenotificationStatusReturned          ACHPrenotificationStatus = "returned"
	ACHPrenotificationStatusSubmitted         ACHPrenotificationStatus = "submitted"
)

func (r ACHPrenotificationStatus) IsKnown() bool {
	switch r {
	case ACHPrenotificationStatusPendingSubmitting, ACHPrenotificationStatusRequiresAttention, ACHPrenotificationStatusReturned, ACHPrenotificationStatusSubmitted:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `ach_prenotification`.
type ACHPrenotificationType string

const (
	ACHPrenotificationTypeACHPrenotification ACHPrenotificationType = "ach_prenotification"
)

func (r ACHPrenotificationType) IsKnown() bool {
	switch r {
	case ACHPrenotificationTypeACHPrenotification:
		return true
	}
	return false
}

// A list of ACH Prenotification objects.
type ACHPrenotificationListResponse struct {
	// The contents of the list.
	Data []ACHPrenotification `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor  string                             `json:"next_cursor,required,nullable"`
	ExtraFields map[string]interface{}             `json:"-,extras"`
	JSON        achPrenotificationListResponseJSON `json:"-"`
}

// achPrenotificationListResponseJSON contains the JSON metadata for the struct
// [ACHPrenotificationListResponse]
type achPrenotificationListResponseJSON struct {
	Data        apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACHPrenotificationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r achPrenotificationListResponseJSON) RawJSON() string {
	return r.raw
}

type ACHPrenotificationNewParams struct {
	// The Increase identifier for the account that will send the ACH Prenotification.
	AccountID param.Field[string] `json:"account_id,required"`
	// The account number for the destination account.
	AccountNumber param.Field[string] `json:"account_number,required"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
	// destination account.
	RoutingNumber param.Field[string] `json:"routing_number,required"`
	// Additional information that will be sent to the recipient.
	Addendum param.Field[string] `json:"addendum"`
	// The description of the date of the ACH Prenotification.
	CompanyDescriptiveDate param.Field[string] `json:"company_descriptive_date"`
	// The data you choose to associate with the ACH Prenotification.
	CompanyDiscretionaryData param.Field[string] `json:"company_discretionary_data"`
	// The description you wish to be shown to the recipient.
	CompanyEntryDescription param.Field[string] `json:"company_entry_description"`
	// The name by which the recipient knows you.
	CompanyName param.Field[string] `json:"company_name"`
	// Whether the Prenotification is for a future debit or credit.
	CreditDebitIndicator param.Field[ACHPrenotificationNewParamsCreditDebitIndicator] `json:"credit_debit_indicator"`
	// The ACH Prenotification effective date in
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	EffectiveDate param.Field[time.Time] `json:"effective_date" format:"date"`
	// Your identifier for the recipient.
	IndividualID param.Field[string] `json:"individual_id"`
	// The name of therecipient. This value is informational and not verified by the
	// recipient's bank.
	IndividualName param.Field[string] `json:"individual_name"`
	// The Standard Entry Class (SEC) code to use for the ACH Prenotification.
	StandardEntryClassCode param.Field[ACHPrenotificationNewParamsStandardEntryClassCode] `json:"standard_entry_class_code"`
}

func (r ACHPrenotificationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether the Prenotification is for a future debit or credit.
type ACHPrenotificationNewParamsCreditDebitIndicator string

const (
	ACHPrenotificationNewParamsCreditDebitIndicatorCredit ACHPrenotificationNewParamsCreditDebitIndicator = "credit"
	ACHPrenotificationNewParamsCreditDebitIndicatorDebit  ACHPrenotificationNewParamsCreditDebitIndicator = "debit"
)

func (r ACHPrenotificationNewParamsCreditDebitIndicator) IsKnown() bool {
	switch r {
	case ACHPrenotificationNewParamsCreditDebitIndicatorCredit, ACHPrenotificationNewParamsCreditDebitIndicatorDebit:
		return true
	}
	return false
}

// The Standard Entry Class (SEC) code to use for the ACH Prenotification.
type ACHPrenotificationNewParamsStandardEntryClassCode string

const (
	ACHPrenotificationNewParamsStandardEntryClassCodeCorporateCreditOrDebit        ACHPrenotificationNewParamsStandardEntryClassCode = "corporate_credit_or_debit"
	ACHPrenotificationNewParamsStandardEntryClassCodeCorporateTradeExchange        ACHPrenotificationNewParamsStandardEntryClassCode = "corporate_trade_exchange"
	ACHPrenotificationNewParamsStandardEntryClassCodePrearrangedPaymentsAndDeposit ACHPrenotificationNewParamsStandardEntryClassCode = "prearranged_payments_and_deposit"
	ACHPrenotificationNewParamsStandardEntryClassCodeInternetInitiated             ACHPrenotificationNewParamsStandardEntryClassCode = "internet_initiated"
)

func (r ACHPrenotificationNewParamsStandardEntryClassCode) IsKnown() bool {
	switch r {
	case ACHPrenotificationNewParamsStandardEntryClassCodeCorporateCreditOrDebit, ACHPrenotificationNewParamsStandardEntryClassCodeCorporateTradeExchange, ACHPrenotificationNewParamsStandardEntryClassCodePrearrangedPaymentsAndDeposit, ACHPrenotificationNewParamsStandardEntryClassCodeInternetInitiated:
		return true
	}
	return false
}

type ACHPrenotificationListParams struct {
	CreatedAt param.Field[ACHPrenotificationListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Filter records to the one with the specified `idempotency_key` you chose for
	// that object. This value is unique across Increase and is used to ensure that a
	// request is only processed once. Learn more about
	// [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey param.Field[string] `query:"idempotency_key"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [ACHPrenotificationListParams]'s query parameters as
// `url.Values`.
func (r ACHPrenotificationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ACHPrenotificationListParamsCreatedAt struct {
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

// URLQuery serializes [ACHPrenotificationListParamsCreatedAt]'s query parameters
// as `url.Values`.
func (r ACHPrenotificationListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Increase/increase-go/internal/apijson"
	"github.com/Increase/increase-go/internal/apiquery"
	"github.com/Increase/increase-go/internal/pagination"
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
	opts = append(r.Options[:], opts...)
	path := "ach_prenotifications"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve an ACH Prenotification
func (r *ACHPrenotificationService) Get(ctx context.Context, achPrenotificationID string, opts ...option.RequestOption) (res *ACHPrenotification, err error) {
	opts = append(r.Options[:], opts...)
	if achPrenotificationID == "" {
		err = errors.New("missing required ach_prenotification_id parameter")
		return
	}
	path := fmt.Sprintf("ach_prenotifications/%s", achPrenotificationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List ACH Prenotifications
func (r *ACHPrenotificationService) List(ctx context.Context, query ACHPrenotificationListParams, opts ...option.RequestOption) (res *pagination.Page[ACHPrenotification], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "ach_prenotifications"
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

// List ACH Prenotifications
func (r *ACHPrenotificationService) ListAutoPaging(ctx context.Context, query ACHPrenotificationListParams, opts ...option.RequestOption) *pagination.PageAutoPager[ACHPrenotification] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// ACH Prenotifications are one way you can verify account and routing numbers by
// Automated Clearing House (ACH).
type ACHPrenotification struct {
	// The ACH Prenotification's identifier.
	ID string `json:"id,required"`
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
	// If the receiving bank notifies that future transfers should use different
	// details, this will contain those details.
	NotificationsOfChange []ACHPrenotificationNotificationsOfChange `json:"notifications_of_change,required"`
	// If your prenotification is returned, this will contain details of the return.
	PrenotificationReturn ACHPrenotificationPrenotificationReturn `json:"prenotification_return,required,nullable"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber string `json:"routing_number,required"`
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
	NotificationsOfChange    apijson.Field
	PrenotificationReturn    apijson.Field
	RoutingNumber            apijson.Field
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
	// The Prenotification is for an anticipated credit.
	ACHPrenotificationCreditDebitIndicatorCredit ACHPrenotificationCreditDebitIndicator = "credit"
	// The Prenotification is for an anticipated debit.
	ACHPrenotificationCreditDebitIndicatorDebit ACHPrenotificationCreditDebitIndicator = "debit"
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
	// The account number was incorrect.
	ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectAccountNumber ACHPrenotificationNotificationsOfChangeChangeCode = "incorrect_account_number"
	// The routing number was incorrect.
	ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectRoutingNumber ACHPrenotificationNotificationsOfChangeChangeCode = "incorrect_routing_number"
	// Both the routing number and the account number were incorrect.
	ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectRoutingNumberAndAccountNumber ACHPrenotificationNotificationsOfChangeChangeCode = "incorrect_routing_number_and_account_number"
	// The transaction code was incorrect. Try changing the `funding` parameter from
	// checking to savings or vice-versa.
	ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectTransactionCode ACHPrenotificationNotificationsOfChangeChangeCode = "incorrect_transaction_code"
	// The account number and the transaction code were incorrect.
	ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectAccountNumberAndTransactionCode ACHPrenotificationNotificationsOfChangeChangeCode = "incorrect_account_number_and_transaction_code"
	// The routing number, account number, and transaction code were incorrect.
	ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectRoutingNumberAccountNumberAndTransactionCode ACHPrenotificationNotificationsOfChangeChangeCode = "incorrect_routing_number_account_number_and_transaction_code"
	// The receiving depository financial institution identification was incorrect.
	ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectReceivingDepositoryFinancialInstitutionIdentification ACHPrenotificationNotificationsOfChangeChangeCode = "incorrect_receiving_depository_financial_institution_identification"
	// The individual identification number was incorrect.
	ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectIndividualIdentificationNumber ACHPrenotificationNotificationsOfChangeChangeCode = "incorrect_individual_identification_number"
	// The addenda had an incorrect format.
	ACHPrenotificationNotificationsOfChangeChangeCodeAddendaFormatError ACHPrenotificationNotificationsOfChangeChangeCode = "addenda_format_error"
	// The standard entry class code was incorrect for an outbound international
	// payment.
	ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectStandardEntryClassCodeForOutboundInternationalPayment ACHPrenotificationNotificationsOfChangeChangeCode = "incorrect_standard_entry_class_code_for_outbound_international_payment"
	// The notification of change was misrouted.
	ACHPrenotificationNotificationsOfChangeChangeCodeMisroutedNotificationOfChange ACHPrenotificationNotificationsOfChangeChangeCode = "misrouted_notification_of_change"
	// The trace number was incorrect.
	ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectTraceNumber ACHPrenotificationNotificationsOfChangeChangeCode = "incorrect_trace_number"
	// The company identification number was incorrect.
	ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectCompanyIdentificationNumber ACHPrenotificationNotificationsOfChangeChangeCode = "incorrect_company_identification_number"
	// The individual identification number or identification number was incorrect.
	ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectIdentificationNumber ACHPrenotificationNotificationsOfChangeChangeCode = "incorrect_identification_number"
	// The corrected data was incorrectly formatted.
	ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectlyFormattedCorrectedData ACHPrenotificationNotificationsOfChangeChangeCode = "incorrectly_formatted_corrected_data"
	// The discretionary data was incorrect.
	ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectDiscretionaryData ACHPrenotificationNotificationsOfChangeChangeCode = "incorrect_discretionary_data"
	// The routing number was not from the original entry detail record.
	ACHPrenotificationNotificationsOfChangeChangeCodeRoutingNumberNotFromOriginalEntryDetailRecord ACHPrenotificationNotificationsOfChangeChangeCode = "routing_number_not_from_original_entry_detail_record"
	// The depository financial institution account number was not from the original
	// entry detail record.
	ACHPrenotificationNotificationsOfChangeChangeCodeDepositoryFinancialInstitutionAccountNumberNotFromOriginalEntryDetailRecord ACHPrenotificationNotificationsOfChangeChangeCode = "depository_financial_institution_account_number_not_from_original_entry_detail_record"
	// The transaction code was incorrect, initiated by the originating depository
	// financial institution.
	ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectTransactionCodeByOriginatingDepositoryFinancialInstitution ACHPrenotificationNotificationsOfChangeChangeCode = "incorrect_transaction_code_by_originating_depository_financial_institution"
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
	// Code R01. Insufficient funds in the receiving account. Sometimes abbreviated to
	// NSF.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeInsufficientFund ACHPrenotificationPrenotificationReturnReturnReasonCode = "insufficient_fund"
	// Code R03. The account does not exist or the receiving bank was unable to locate
	// it.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeNoAccount ACHPrenotificationPrenotificationReturnReturnReasonCode = "no_account"
	// Code R02. The account is closed at the receiving bank.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeAccountClosed ACHPrenotificationPrenotificationReturnReturnReasonCode = "account_closed"
	// Code R04. The account number is invalid at the receiving bank.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeInvalidAccountNumberStructure ACHPrenotificationPrenotificationReturnReturnReasonCode = "invalid_account_number_structure"
	// Code R16. The account at the receiving bank was frozen per the Office of Foreign
	// Assets Control.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeAccountFrozenEntryReturnedPerOfacInstruction ACHPrenotificationPrenotificationReturnReturnReasonCode = "account_frozen_entry_returned_per_ofac_instruction"
	// Code R23. The receiving bank account refused a credit transfer.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeCreditEntryRefusedByReceiver ACHPrenotificationPrenotificationReturnReturnReasonCode = "credit_entry_refused_by_receiver"
	// Code R05. The receiving bank rejected because of an incorrect Standard Entry
	// Class code.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeUnauthorizedDebitToConsumerAccountUsingCorporateSecCode ACHPrenotificationPrenotificationReturnReturnReasonCode = "unauthorized_debit_to_consumer_account_using_corporate_sec_code"
	// Code R29. The corporate customer at the receiving bank reversed the transfer.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeCorporateCustomerAdvisedNotAuthorized ACHPrenotificationPrenotificationReturnReturnReasonCode = "corporate_customer_advised_not_authorized"
	// Code R08. The receiving bank stopped payment on this transfer.
	ACHPrenotificationPrenotificationReturnReturnReasonCodePaymentStopped ACHPrenotificationPrenotificationReturnReturnReasonCode = "payment_stopped"
	// Code R20. The receiving bank account does not perform transfers.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeNonTransactionAccount ACHPrenotificationPrenotificationReturnReturnReasonCode = "non_transaction_account"
	// Code R09. The receiving bank account does not have enough available balance for
	// the transfer.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeUncollectedFunds ACHPrenotificationPrenotificationReturnReturnReasonCode = "uncollected_funds"
	// Code R28. The routing number is incorrect.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeRoutingNumberCheckDigitError ACHPrenotificationPrenotificationReturnReturnReasonCode = "routing_number_check_digit_error"
	// Code R10. The customer at the receiving bank reversed the transfer.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete ACHPrenotificationPrenotificationReturnReturnReasonCode = "customer_advised_unauthorized_improper_ineligible_or_incomplete"
	// Code R19. The amount field is incorrect or too large.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeAmountFieldError ACHPrenotificationPrenotificationReturnReturnReasonCode = "amount_field_error"
	// Code R07. The customer at the receiving institution informed their bank that
	// they have revoked authorization for a previously authorized transfer.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeAuthorizationRevokedByCustomer ACHPrenotificationPrenotificationReturnReturnReasonCode = "authorization_revoked_by_customer"
	// Code R13. The routing number is invalid.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeInvalidACHRoutingNumber ACHPrenotificationPrenotificationReturnReturnReasonCode = "invalid_ach_routing_number"
	// Code R17. The receiving bank is unable to process a field in the transfer.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeFileRecordEditCriteria ACHPrenotificationPrenotificationReturnReturnReasonCode = "file_record_edit_criteria"
	// Code R45. The individual name field was invalid.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeEnrInvalidIndividualName ACHPrenotificationPrenotificationReturnReturnReasonCode = "enr_invalid_individual_name"
	// Code R06. The originating financial institution asked for this transfer to be
	// returned. The receiving bank is complying with the request.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeReturnedPerOdfiRequest ACHPrenotificationPrenotificationReturnReturnReasonCode = "returned_per_odfi_request"
	// Code R34. The receiving bank's regulatory supervisor has limited their
	// participation in the ACH network.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeLimitedParticipationDfi ACHPrenotificationPrenotificationReturnReturnReasonCode = "limited_participation_dfi"
	// Code R85. The outbound international ACH transfer was incorrect.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeIncorrectlyCodedOutboundInternationalPayment ACHPrenotificationPrenotificationReturnReturnReasonCode = "incorrectly_coded_outbound_international_payment"
	// Code R12. A rare return reason. The account was sold to another bank.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeAccountSoldToAnotherDfi ACHPrenotificationPrenotificationReturnReturnReasonCode = "account_sold_to_another_dfi"
	// Code R25. The addenda record is incorrect or missing.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeAddendaError ACHPrenotificationPrenotificationReturnReturnReasonCode = "addenda_error"
	// Code R15. A rare return reason. The account holder is deceased.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeBeneficiaryOrAccountHolderDeceased ACHPrenotificationPrenotificationReturnReturnReasonCode = "beneficiary_or_account_holder_deceased"
	// Code R11. A rare return reason. The customer authorized some payment to the
	// sender, but this payment was not in error.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeCustomerAdvisedNotWithinAuthorizationTerms ACHPrenotificationPrenotificationReturnReturnReasonCode = "customer_advised_not_within_authorization_terms"
	// Code R74. A rare return reason. Sent in response to a return that was returned
	// with code `field_error`. The latest return should include the corrected
	// field(s).
	ACHPrenotificationPrenotificationReturnReturnReasonCodeCorrectedReturn ACHPrenotificationPrenotificationReturnReturnReasonCode = "corrected_return"
	// Code R24. A rare return reason. The receiving bank received an exact duplicate
	// entry with the same trace number and amount.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeDuplicateEntry ACHPrenotificationPrenotificationReturnReturnReasonCode = "duplicate_entry"
	// Code R67. A rare return reason. The return this message refers to was a
	// duplicate.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeDuplicateReturn ACHPrenotificationPrenotificationReturnReturnReasonCode = "duplicate_return"
	// Code R47. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeEnrDuplicateEnrollment ACHPrenotificationPrenotificationReturnReturnReasonCode = "enr_duplicate_enrollment"
	// Code R43. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeEnrInvalidDfiAccountNumber ACHPrenotificationPrenotificationReturnReturnReasonCode = "enr_invalid_dfi_account_number"
	// Code R44. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeEnrInvalidIndividualIDNumber ACHPrenotificationPrenotificationReturnReturnReasonCode = "enr_invalid_individual_id_number"
	// Code R46. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeEnrInvalidRepresentativePayeeIndicator ACHPrenotificationPrenotificationReturnReturnReasonCode = "enr_invalid_representative_payee_indicator"
	// Code R41. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeEnrInvalidTransactionCode ACHPrenotificationPrenotificationReturnReturnReasonCode = "enr_invalid_transaction_code"
	// Code R40. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeEnrReturnOfEnrEntry ACHPrenotificationPrenotificationReturnReturnReasonCode = "enr_return_of_enr_entry"
	// Code R42. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeEnrRoutingNumberCheckDigitError ACHPrenotificationPrenotificationReturnReturnReasonCode = "enr_routing_number_check_digit_error"
	// Code R84. A rare return reason. The International ACH Transfer cannot be
	// processed by the gateway.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeEntryNotProcessedByGateway ACHPrenotificationPrenotificationReturnReturnReasonCode = "entry_not_processed_by_gateway"
	// Code R69. A rare return reason. One or more of the fields in the ACH were
	// malformed.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeFieldError ACHPrenotificationPrenotificationReturnReturnReasonCode = "field_error"
	// Code R83. A rare return reason. The Foreign receiving bank was unable to settle
	// this ACH transfer.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeForeignReceivingDfiUnableToSettle ACHPrenotificationPrenotificationReturnReturnReasonCode = "foreign_receiving_dfi_unable_to_settle"
	// Code R80. A rare return reason. The International ACH Transfer is malformed.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeIatEntryCodingError ACHPrenotificationPrenotificationReturnReturnReasonCode = "iat_entry_coding_error"
	// Code R18. A rare return reason. The ACH has an improper effective entry date
	// field.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeImproperEffectiveEntryDate ACHPrenotificationPrenotificationReturnReturnReasonCode = "improper_effective_entry_date"
	// Code R39. A rare return reason. The source document related to this ACH, usually
	// an ACH check conversion, was presented to the bank.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeImproperSourceDocumentSourceDocumentPresented ACHPrenotificationPrenotificationReturnReturnReasonCode = "improper_source_document_source_document_presented"
	// Code R21. A rare return reason. The Company ID field of the ACH was invalid.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeInvalidCompanyID ACHPrenotificationPrenotificationReturnReturnReasonCode = "invalid_company_id"
	// Code R82. A rare return reason. The foreign receiving bank identifier for an
	// International ACH Transfer was invalid.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeInvalidForeignReceivingDfiIdentification ACHPrenotificationPrenotificationReturnReturnReasonCode = "invalid_foreign_receiving_dfi_identification"
	// Code R22. A rare return reason. The Individual ID number field of the ACH was
	// invalid.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeInvalidIndividualIDNumber ACHPrenotificationPrenotificationReturnReturnReasonCode = "invalid_individual_id_number"
	// Code R53. A rare return reason. Both the Represented Check ("RCK") entry and the
	// original check were presented to the bank.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeItemAndRckEntryPresentedForPayment ACHPrenotificationPrenotificationReturnReturnReasonCode = "item_and_rck_entry_presented_for_payment"
	// Code R51. A rare return reason. The Represented Check ("RCK") entry is
	// ineligible.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeItemRelatedToRckEntryIsIneligible ACHPrenotificationPrenotificationReturnReturnReasonCode = "item_related_to_rck_entry_is_ineligible"
	// Code R26. A rare return reason. The ACH is missing a required field.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeMandatoryFieldError ACHPrenotificationPrenotificationReturnReturnReasonCode = "mandatory_field_error"
	// Code R71. A rare return reason. The receiving bank does not recognize the
	// routing number in a dishonored return entry.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeMisroutedDishonoredReturn ACHPrenotificationPrenotificationReturnReturnReasonCode = "misrouted_dishonored_return"
	// Code R61. A rare return reason. The receiving bank does not recognize the
	// routing number in a return entry.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeMisroutedReturn ACHPrenotificationPrenotificationReturnReturnReasonCode = "misrouted_return"
	// Code R76. A rare return reason. Sent in response to a return, the bank does not
	// find the errors alleged by the returning bank.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeNoErrorsFound ACHPrenotificationPrenotificationReturnReturnReasonCode = "no_errors_found"
	// Code R77. A rare return reason. The receiving bank does not accept the return of
	// the erroneous debit. The funds are not available at the receiving bank.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeNonAcceptanceOfR62DishonoredReturn ACHPrenotificationPrenotificationReturnReturnReasonCode = "non_acceptance_of_r62_dishonored_return"
	// Code R81. A rare return reason. The receiving bank does not accept International
	// ACH Transfers.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeNonParticipantInIatProgram ACHPrenotificationPrenotificationReturnReturnReasonCode = "non_participant_in_iat_program"
	// Code R31. A rare return reason. A return that has been agreed to be accepted by
	// the receiving bank, despite falling outside of the usual return timeframe.
	ACHPrenotificationPrenotificationReturnReturnReasonCodePermissibleReturnEntry ACHPrenotificationPrenotificationReturnReturnReasonCode = "permissible_return_entry"
	// Code R70. A rare return reason. The receiving bank had not approved this return.
	ACHPrenotificationPrenotificationReturnReturnReasonCodePermissibleReturnEntryNotAccepted ACHPrenotificationPrenotificationReturnReturnReasonCode = "permissible_return_entry_not_accepted"
	// Code R32. A rare return reason. The receiving bank could not settle this
	// transaction.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeRdfiNonSettlement ACHPrenotificationPrenotificationReturnReturnReasonCode = "rdfi_non_settlement"
	// Code R30. A rare return reason. The receiving bank does not accept Check
	// Truncation ACH transfers.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeRdfiParticipantInCheckTruncationProgram ACHPrenotificationPrenotificationReturnReturnReasonCode = "rdfi_participant_in_check_truncation_program"
	// Code R14. A rare return reason. The payee is deceased.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeRepresentativePayeeDeceasedOrUnableToContinueInThatCapacity ACHPrenotificationPrenotificationReturnReturnReasonCode = "representative_payee_deceased_or_unable_to_continue_in_that_capacity"
	// Code R75. A rare return reason. The originating bank disputes that an earlier
	// `duplicate_entry` return was actually a duplicate.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeReturnNotADuplicate ACHPrenotificationPrenotificationReturnReturnReasonCode = "return_not_a_duplicate"
	// Code R62. A rare return reason. The originating financial institution made a
	// mistake and this return corrects it.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeReturnOfErroneousOrReversingDebit ACHPrenotificationPrenotificationReturnReturnReasonCode = "return_of_erroneous_or_reversing_debit"
	// Code R36. A rare return reason. Return of a malformed credit entry.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeReturnOfImproperCreditEntry ACHPrenotificationPrenotificationReturnReturnReasonCode = "return_of_improper_credit_entry"
	// Code R35. A rare return reason. Return of a malformed debit entry.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeReturnOfImproperDebitEntry ACHPrenotificationPrenotificationReturnReturnReasonCode = "return_of_improper_debit_entry"
	// Code R33. A rare return reason. Return of a Destroyed Check ("XKC") entry.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeReturnOfXckEntry ACHPrenotificationPrenotificationReturnReturnReasonCode = "return_of_xck_entry"
	// Code R37. A rare return reason. The source document related to this ACH, usually
	// an ACH check conversion, was presented to the bank.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeSourceDocumentPresentedForPayment ACHPrenotificationPrenotificationReturnReturnReasonCode = "source_document_presented_for_payment"
	// Code R50. A rare return reason. State law prevents the bank from accepting the
	// Represented Check ("RCK") entry.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeStateLawAffectingRckAcceptance ACHPrenotificationPrenotificationReturnReturnReasonCode = "state_law_affecting_rck_acceptance"
	// Code R52. A rare return reason. A stop payment was issued on a Represented Check
	// ("RCK") entry.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeStopPaymentOnItemRelatedToRckEntry ACHPrenotificationPrenotificationReturnReturnReasonCode = "stop_payment_on_item_related_to_rck_entry"
	// Code R38. A rare return reason. The source attached to the ACH, usually an ACH
	// check conversion, includes a stop payment.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeStopPaymentOnSourceDocument ACHPrenotificationPrenotificationReturnReturnReasonCode = "stop_payment_on_source_document"
	// Code R73. A rare return reason. The bank receiving an `untimely_return` believes
	// it was on time.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeTimelyOriginalReturn ACHPrenotificationPrenotificationReturnReturnReasonCode = "timely_original_return"
	// Code R27. A rare return reason. An ACH return's trace number does not match an
	// originated ACH.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeTraceNumberError ACHPrenotificationPrenotificationReturnReturnReasonCode = "trace_number_error"
	// Code R72. A rare return reason. The dishonored return was sent too late.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeUntimelyDishonoredReturn ACHPrenotificationPrenotificationReturnReturnReasonCode = "untimely_dishonored_return"
	// Code R68. A rare return reason. The return was sent too late.
	ACHPrenotificationPrenotificationReturnReturnReasonCodeUntimelyReturn ACHPrenotificationPrenotificationReturnReturnReasonCode = "untimely_return"
)

func (r ACHPrenotificationPrenotificationReturnReturnReasonCode) IsKnown() bool {
	switch r {
	case ACHPrenotificationPrenotificationReturnReturnReasonCodeInsufficientFund, ACHPrenotificationPrenotificationReturnReturnReasonCodeNoAccount, ACHPrenotificationPrenotificationReturnReturnReasonCodeAccountClosed, ACHPrenotificationPrenotificationReturnReturnReasonCodeInvalidAccountNumberStructure, ACHPrenotificationPrenotificationReturnReturnReasonCodeAccountFrozenEntryReturnedPerOfacInstruction, ACHPrenotificationPrenotificationReturnReturnReasonCodeCreditEntryRefusedByReceiver, ACHPrenotificationPrenotificationReturnReturnReasonCodeUnauthorizedDebitToConsumerAccountUsingCorporateSecCode, ACHPrenotificationPrenotificationReturnReturnReasonCodeCorporateCustomerAdvisedNotAuthorized, ACHPrenotificationPrenotificationReturnReturnReasonCodePaymentStopped, ACHPrenotificationPrenotificationReturnReturnReasonCodeNonTransactionAccount, ACHPrenotificationPrenotificationReturnReturnReasonCodeUncollectedFunds, ACHPrenotificationPrenotificationReturnReturnReasonCodeRoutingNumberCheckDigitError, ACHPrenotificationPrenotificationReturnReturnReasonCodeCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete, ACHPrenotificationPrenotificationReturnReturnReasonCodeAmountFieldError, ACHPrenotificationPrenotificationReturnReturnReasonCodeAuthorizationRevokedByCustomer, ACHPrenotificationPrenotificationReturnReturnReasonCodeInvalidACHRoutingNumber, ACHPrenotificationPrenotificationReturnReturnReasonCodeFileRecordEditCriteria, ACHPrenotificationPrenotificationReturnReturnReasonCodeEnrInvalidIndividualName, ACHPrenotificationPrenotificationReturnReturnReasonCodeReturnedPerOdfiRequest, ACHPrenotificationPrenotificationReturnReturnReasonCodeLimitedParticipationDfi, ACHPrenotificationPrenotificationReturnReturnReasonCodeIncorrectlyCodedOutboundInternationalPayment, ACHPrenotificationPrenotificationReturnReturnReasonCodeAccountSoldToAnotherDfi, ACHPrenotificationPrenotificationReturnReturnReasonCodeAddendaError, ACHPrenotificationPrenotificationReturnReturnReasonCodeBeneficiaryOrAccountHolderDeceased, ACHPrenotificationPrenotificationReturnReturnReasonCodeCustomerAdvisedNotWithinAuthorizationTerms, ACHPrenotificationPrenotificationReturnReturnReasonCodeCorrectedReturn, ACHPrenotificationPrenotificationReturnReturnReasonCodeDuplicateEntry, ACHPrenotificationPrenotificationReturnReturnReasonCodeDuplicateReturn, ACHPrenotificationPrenotificationReturnReturnReasonCodeEnrDuplicateEnrollment, ACHPrenotificationPrenotificationReturnReturnReasonCodeEnrInvalidDfiAccountNumber, ACHPrenotificationPrenotificationReturnReturnReasonCodeEnrInvalidIndividualIDNumber, ACHPrenotificationPrenotificationReturnReturnReasonCodeEnrInvalidRepresentativePayeeIndicator, ACHPrenotificationPrenotificationReturnReturnReasonCodeEnrInvalidTransactionCode, ACHPrenotificationPrenotificationReturnReturnReasonCodeEnrReturnOfEnrEntry, ACHPrenotificationPrenotificationReturnReturnReasonCodeEnrRoutingNumberCheckDigitError, ACHPrenotificationPrenotificationReturnReturnReasonCodeEntryNotProcessedByGateway, ACHPrenotificationPrenotificationReturnReturnReasonCodeFieldError, ACHPrenotificationPrenotificationReturnReturnReasonCodeForeignReceivingDfiUnableToSettle, ACHPrenotificationPrenotificationReturnReturnReasonCodeIatEntryCodingError, ACHPrenotificationPrenotificationReturnReturnReasonCodeImproperEffectiveEntryDate, ACHPrenotificationPrenotificationReturnReturnReasonCodeImproperSourceDocumentSourceDocumentPresented, ACHPrenotificationPrenotificationReturnReturnReasonCodeInvalidCompanyID, ACHPrenotificationPrenotificationReturnReturnReasonCodeInvalidForeignReceivingDfiIdentification, ACHPrenotificationPrenotificationReturnReturnReasonCodeInvalidIndividualIDNumber, ACHPrenotificationPrenotificationReturnReturnReasonCodeItemAndRckEntryPresentedForPayment, ACHPrenotificationPrenotificationReturnReturnReasonCodeItemRelatedToRckEntryIsIneligible, ACHPrenotificationPrenotificationReturnReturnReasonCodeMandatoryFieldError, ACHPrenotificationPrenotificationReturnReturnReasonCodeMisroutedDishonoredReturn, ACHPrenotificationPrenotificationReturnReturnReasonCodeMisroutedReturn, ACHPrenotificationPrenotificationReturnReturnReasonCodeNoErrorsFound, ACHPrenotificationPrenotificationReturnReturnReasonCodeNonAcceptanceOfR62DishonoredReturn, ACHPrenotificationPrenotificationReturnReturnReasonCodeNonParticipantInIatProgram, ACHPrenotificationPrenotificationReturnReturnReasonCodePermissibleReturnEntry, ACHPrenotificationPrenotificationReturnReturnReasonCodePermissibleReturnEntryNotAccepted, ACHPrenotificationPrenotificationReturnReturnReasonCodeRdfiNonSettlement, ACHPrenotificationPrenotificationReturnReturnReasonCodeRdfiParticipantInCheckTruncationProgram, ACHPrenotificationPrenotificationReturnReturnReasonCodeRepresentativePayeeDeceasedOrUnableToContinueInThatCapacity, ACHPrenotificationPrenotificationReturnReturnReasonCodeReturnNotADuplicate, ACHPrenotificationPrenotificationReturnReturnReasonCodeReturnOfErroneousOrReversingDebit, ACHPrenotificationPrenotificationReturnReturnReasonCodeReturnOfImproperCreditEntry, ACHPrenotificationPrenotificationReturnReturnReasonCodeReturnOfImproperDebitEntry, ACHPrenotificationPrenotificationReturnReturnReasonCodeReturnOfXckEntry, ACHPrenotificationPrenotificationReturnReturnReasonCodeSourceDocumentPresentedForPayment, ACHPrenotificationPrenotificationReturnReturnReasonCodeStateLawAffectingRckAcceptance, ACHPrenotificationPrenotificationReturnReturnReasonCodeStopPaymentOnItemRelatedToRckEntry, ACHPrenotificationPrenotificationReturnReturnReasonCodeStopPaymentOnSourceDocument, ACHPrenotificationPrenotificationReturnReturnReasonCodeTimelyOriginalReturn, ACHPrenotificationPrenotificationReturnReturnReasonCodeTraceNumberError, ACHPrenotificationPrenotificationReturnReturnReasonCodeUntimelyDishonoredReturn, ACHPrenotificationPrenotificationReturnReturnReasonCodeUntimelyReturn:
		return true
	}
	return false
}

// The lifecycle status of the ACH Prenotification.
type ACHPrenotificationStatus string

const (
	// The Prenotification is pending submission.
	ACHPrenotificationStatusPendingSubmitting ACHPrenotificationStatus = "pending_submitting"
	// The Prenotification requires attention.
	ACHPrenotificationStatusRequiresAttention ACHPrenotificationStatus = "requires_attention"
	// The Prenotification has been returned.
	ACHPrenotificationStatusReturned ACHPrenotificationStatus = "returned"
	// The Prenotification is complete.
	ACHPrenotificationStatusSubmitted ACHPrenotificationStatus = "submitted"
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

type ACHPrenotificationNewParams struct {
	// The Increase identifier for the account that will send the transfer.
	AccountID param.Field[string] `json:"account_id,required"`
	// The account number for the destination account.
	AccountNumber param.Field[string] `json:"account_number,required"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
	// destination account.
	RoutingNumber param.Field[string] `json:"routing_number,required"`
	// Additional information that will be sent to the recipient.
	Addendum param.Field[string] `json:"addendum"`
	// The description of the date of the transfer.
	CompanyDescriptiveDate param.Field[string] `json:"company_descriptive_date"`
	// The data you choose to associate with the transfer.
	CompanyDiscretionaryData param.Field[string] `json:"company_discretionary_data"`
	// The description of the transfer you wish to be shown to the recipient.
	CompanyEntryDescription param.Field[string] `json:"company_entry_description"`
	// The name by which the recipient knows you.
	CompanyName param.Field[string] `json:"company_name"`
	// Whether the Prenotification is for a future debit or credit.
	CreditDebitIndicator param.Field[ACHPrenotificationNewParamsCreditDebitIndicator] `json:"credit_debit_indicator"`
	// The transfer effective date in
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	EffectiveDate param.Field[time.Time] `json:"effective_date" format:"date"`
	// Your identifier for the transfer recipient.
	IndividualID param.Field[string] `json:"individual_id"`
	// The name of the transfer recipient. This value is information and not verified
	// by the recipient's bank.
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
	// The Prenotification is for an anticipated credit.
	ACHPrenotificationNewParamsCreditDebitIndicatorCredit ACHPrenotificationNewParamsCreditDebitIndicator = "credit"
	// The Prenotification is for an anticipated debit.
	ACHPrenotificationNewParamsCreditDebitIndicatorDebit ACHPrenotificationNewParamsCreditDebitIndicator = "debit"
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
	// Corporate Credit and Debit (CCD).
	ACHPrenotificationNewParamsStandardEntryClassCodeCorporateCreditOrDebit ACHPrenotificationNewParamsStandardEntryClassCode = "corporate_credit_or_debit"
	// Corporate Trade Exchange (CTX).
	ACHPrenotificationNewParamsStandardEntryClassCodeCorporateTradeExchange ACHPrenotificationNewParamsStandardEntryClassCode = "corporate_trade_exchange"
	// Prearranged Payments and Deposits (PPD).
	ACHPrenotificationNewParamsStandardEntryClassCodePrearrangedPaymentsAndDeposit ACHPrenotificationNewParamsStandardEntryClassCode = "prearranged_payments_and_deposit"
	// Internet Initiated (WEB).
	ACHPrenotificationNewParamsStandardEntryClassCodeInternetInitiated ACHPrenotificationNewParamsStandardEntryClassCode = "internet_initiated"
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

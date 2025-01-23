// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/Increase/increase-go/internal/apijson"
	"github.com/Increase/increase-go/internal/param"
	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
)

// SimulationACHTransferService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimulationACHTransferService] method instead.
type SimulationACHTransferService struct {
	Options []option.RequestOption
}

// NewSimulationACHTransferService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSimulationACHTransferService(opts ...option.RequestOption) (r *SimulationACHTransferService) {
	r = &SimulationACHTransferService{}
	r.Options = opts
	return
}

// Simulates the acknowledgement of an [ACH Transfer](#ach-transfers) by the
// Federal Reserve. This transfer must first have a `status` of `submitted` . In
// production, the Federal Reserve generally acknowledges submitted ACH files
// within 30 minutes. Since sandbox ACH Transfers are not submitted to the Federal
// Reserve, this endpoint allows you to skip that delay and add the acknowledgment
// subresource to the ACH Transfer.
func (r *SimulationACHTransferService) Acknowledge(ctx context.Context, achTransferID string, opts ...option.RequestOption) (res *ACHTransfer, err error) {
	opts = append(r.Options[:], opts...)
	if achTransferID == "" {
		err = errors.New("missing required ach_transfer_id parameter")
		return
	}
	path := fmt.Sprintf("simulations/ach_transfers/%s/acknowledge", achTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Simulates receiving a Notification of Change for an
// [ACH Transfer](#ach-transfers).
func (r *SimulationACHTransferService) NewNotificationOfChange(ctx context.Context, achTransferID string, body SimulationACHTransferNewNotificationOfChangeParams, opts ...option.RequestOption) (res *ACHTransfer, err error) {
	opts = append(r.Options[:], opts...)
	if achTransferID == "" {
		err = errors.New("missing required ach_transfer_id parameter")
		return
	}
	path := fmt.Sprintf("simulations/ach_transfers/%s/create_notification_of_change", achTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Simulates the return of an [ACH Transfer](#ach-transfers) by the Federal Reserve
// due to an error condition. This will also create a Transaction to account for
// the returned funds. This transfer must first have a `status` of `submitted`.
func (r *SimulationACHTransferService) Return(ctx context.Context, achTransferID string, body SimulationACHTransferReturnParams, opts ...option.RequestOption) (res *ACHTransfer, err error) {
	opts = append(r.Options[:], opts...)
	if achTransferID == "" {
		err = errors.New("missing required ach_transfer_id parameter")
		return
	}
	path := fmt.Sprintf("simulations/ach_transfers/%s/return", achTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Simulates the settlement of an [ACH Transfer](#ach-transfers) by the Federal
// Reserve. This transfer must first have a `status` of `submitted`. Without this
// simulation the transfer will eventually settle on its own following the same
// Federal Reserve timeline as in production.
func (r *SimulationACHTransferService) Settle(ctx context.Context, achTransferID string, opts ...option.RequestOption) (res *ACHTransfer, err error) {
	opts = append(r.Options[:], opts...)
	if achTransferID == "" {
		err = errors.New("missing required ach_transfer_id parameter")
		return
	}
	path := fmt.Sprintf("simulations/ach_transfers/%s/settle", achTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Simulates the submission of an [ACH Transfer](#ach-transfers) to the Federal
// Reserve. This transfer must first have a `status` of `pending_approval` or
// `pending_submission`. In production, Increase submits ACH Transfers to the
// Federal Reserve three times per day on weekdays. Since sandbox ACH Transfers are
// not submitted to the Federal Reserve, this endpoint allows you to skip that
// delay and transition the ACH Transfer to a status of `submitted`.
func (r *SimulationACHTransferService) Submit(ctx context.Context, achTransferID string, opts ...option.RequestOption) (res *ACHTransfer, err error) {
	opts = append(r.Options[:], opts...)
	if achTransferID == "" {
		err = errors.New("missing required ach_transfer_id parameter")
		return
	}
	path := fmt.Sprintf("simulations/ach_transfers/%s/submit", achTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type SimulationACHTransferNewNotificationOfChangeParams struct {
	// The reason for the notification of change.
	ChangeCode param.Field[SimulationACHTransferNewNotificationOfChangeParamsChangeCode] `json:"change_code,required"`
	// The corrected data for the notification of change (e.g., a new routing number).
	CorrectedData param.Field[string] `json:"corrected_data,required"`
}

func (r SimulationACHTransferNewNotificationOfChangeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The reason for the notification of change.
type SimulationACHTransferNewNotificationOfChangeParamsChangeCode string

const (
	SimulationACHTransferNewNotificationOfChangeParamsChangeCodeIncorrectAccountNumber                                                      SimulationACHTransferNewNotificationOfChangeParamsChangeCode = "incorrect_account_number"
	SimulationACHTransferNewNotificationOfChangeParamsChangeCodeIncorrectRoutingNumber                                                      SimulationACHTransferNewNotificationOfChangeParamsChangeCode = "incorrect_routing_number"
	SimulationACHTransferNewNotificationOfChangeParamsChangeCodeIncorrectRoutingNumberAndAccountNumber                                      SimulationACHTransferNewNotificationOfChangeParamsChangeCode = "incorrect_routing_number_and_account_number"
	SimulationACHTransferNewNotificationOfChangeParamsChangeCodeIncorrectTransactionCode                                                    SimulationACHTransferNewNotificationOfChangeParamsChangeCode = "incorrect_transaction_code"
	SimulationACHTransferNewNotificationOfChangeParamsChangeCodeIncorrectAccountNumberAndTransactionCode                                    SimulationACHTransferNewNotificationOfChangeParamsChangeCode = "incorrect_account_number_and_transaction_code"
	SimulationACHTransferNewNotificationOfChangeParamsChangeCodeIncorrectRoutingNumberAccountNumberAndTransactionCode                       SimulationACHTransferNewNotificationOfChangeParamsChangeCode = "incorrect_routing_number_account_number_and_transaction_code"
	SimulationACHTransferNewNotificationOfChangeParamsChangeCodeIncorrectReceivingDepositoryFinancialInstitutionIdentification              SimulationACHTransferNewNotificationOfChangeParamsChangeCode = "incorrect_receiving_depository_financial_institution_identification"
	SimulationACHTransferNewNotificationOfChangeParamsChangeCodeIncorrectIndividualIdentificationNumber                                     SimulationACHTransferNewNotificationOfChangeParamsChangeCode = "incorrect_individual_identification_number"
	SimulationACHTransferNewNotificationOfChangeParamsChangeCodeAddendaFormatError                                                          SimulationACHTransferNewNotificationOfChangeParamsChangeCode = "addenda_format_error"
	SimulationACHTransferNewNotificationOfChangeParamsChangeCodeIncorrectStandardEntryClassCodeForOutboundInternationalPayment              SimulationACHTransferNewNotificationOfChangeParamsChangeCode = "incorrect_standard_entry_class_code_for_outbound_international_payment"
	SimulationACHTransferNewNotificationOfChangeParamsChangeCodeMisroutedNotificationOfChange                                               SimulationACHTransferNewNotificationOfChangeParamsChangeCode = "misrouted_notification_of_change"
	SimulationACHTransferNewNotificationOfChangeParamsChangeCodeIncorrectTraceNumber                                                        SimulationACHTransferNewNotificationOfChangeParamsChangeCode = "incorrect_trace_number"
	SimulationACHTransferNewNotificationOfChangeParamsChangeCodeIncorrectCompanyIdentificationNumber                                        SimulationACHTransferNewNotificationOfChangeParamsChangeCode = "incorrect_company_identification_number"
	SimulationACHTransferNewNotificationOfChangeParamsChangeCodeIncorrectIdentificationNumber                                               SimulationACHTransferNewNotificationOfChangeParamsChangeCode = "incorrect_identification_number"
	SimulationACHTransferNewNotificationOfChangeParamsChangeCodeIncorrectlyFormattedCorrectedData                                           SimulationACHTransferNewNotificationOfChangeParamsChangeCode = "incorrectly_formatted_corrected_data"
	SimulationACHTransferNewNotificationOfChangeParamsChangeCodeIncorrectDiscretionaryData                                                  SimulationACHTransferNewNotificationOfChangeParamsChangeCode = "incorrect_discretionary_data"
	SimulationACHTransferNewNotificationOfChangeParamsChangeCodeRoutingNumberNotFromOriginalEntryDetailRecord                               SimulationACHTransferNewNotificationOfChangeParamsChangeCode = "routing_number_not_from_original_entry_detail_record"
	SimulationACHTransferNewNotificationOfChangeParamsChangeCodeDepositoryFinancialInstitutionAccountNumberNotFromOriginalEntryDetailRecord SimulationACHTransferNewNotificationOfChangeParamsChangeCode = "depository_financial_institution_account_number_not_from_original_entry_detail_record"
	SimulationACHTransferNewNotificationOfChangeParamsChangeCodeIncorrectTransactionCodeByOriginatingDepositoryFinancialInstitution         SimulationACHTransferNewNotificationOfChangeParamsChangeCode = "incorrect_transaction_code_by_originating_depository_financial_institution"
)

func (r SimulationACHTransferNewNotificationOfChangeParamsChangeCode) IsKnown() bool {
	switch r {
	case SimulationACHTransferNewNotificationOfChangeParamsChangeCodeIncorrectAccountNumber, SimulationACHTransferNewNotificationOfChangeParamsChangeCodeIncorrectRoutingNumber, SimulationACHTransferNewNotificationOfChangeParamsChangeCodeIncorrectRoutingNumberAndAccountNumber, SimulationACHTransferNewNotificationOfChangeParamsChangeCodeIncorrectTransactionCode, SimulationACHTransferNewNotificationOfChangeParamsChangeCodeIncorrectAccountNumberAndTransactionCode, SimulationACHTransferNewNotificationOfChangeParamsChangeCodeIncorrectRoutingNumberAccountNumberAndTransactionCode, SimulationACHTransferNewNotificationOfChangeParamsChangeCodeIncorrectReceivingDepositoryFinancialInstitutionIdentification, SimulationACHTransferNewNotificationOfChangeParamsChangeCodeIncorrectIndividualIdentificationNumber, SimulationACHTransferNewNotificationOfChangeParamsChangeCodeAddendaFormatError, SimulationACHTransferNewNotificationOfChangeParamsChangeCodeIncorrectStandardEntryClassCodeForOutboundInternationalPayment, SimulationACHTransferNewNotificationOfChangeParamsChangeCodeMisroutedNotificationOfChange, SimulationACHTransferNewNotificationOfChangeParamsChangeCodeIncorrectTraceNumber, SimulationACHTransferNewNotificationOfChangeParamsChangeCodeIncorrectCompanyIdentificationNumber, SimulationACHTransferNewNotificationOfChangeParamsChangeCodeIncorrectIdentificationNumber, SimulationACHTransferNewNotificationOfChangeParamsChangeCodeIncorrectlyFormattedCorrectedData, SimulationACHTransferNewNotificationOfChangeParamsChangeCodeIncorrectDiscretionaryData, SimulationACHTransferNewNotificationOfChangeParamsChangeCodeRoutingNumberNotFromOriginalEntryDetailRecord, SimulationACHTransferNewNotificationOfChangeParamsChangeCodeDepositoryFinancialInstitutionAccountNumberNotFromOriginalEntryDetailRecord, SimulationACHTransferNewNotificationOfChangeParamsChangeCodeIncorrectTransactionCodeByOriginatingDepositoryFinancialInstitution:
		return true
	}
	return false
}

type SimulationACHTransferReturnParams struct {
	// The reason why the Federal Reserve or destination bank returned this transfer.
	// Defaults to `no_account`.
	Reason param.Field[SimulationACHTransferReturnParamsReason] `json:"reason"`
}

func (r SimulationACHTransferReturnParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The reason why the Federal Reserve or destination bank returned this transfer.
// Defaults to `no_account`.
type SimulationACHTransferReturnParamsReason string

const (
	SimulationACHTransferReturnParamsReasonInsufficientFund                                            SimulationACHTransferReturnParamsReason = "insufficient_fund"
	SimulationACHTransferReturnParamsReasonNoAccount                                                   SimulationACHTransferReturnParamsReason = "no_account"
	SimulationACHTransferReturnParamsReasonAccountClosed                                               SimulationACHTransferReturnParamsReason = "account_closed"
	SimulationACHTransferReturnParamsReasonInvalidAccountNumberStructure                               SimulationACHTransferReturnParamsReason = "invalid_account_number_structure"
	SimulationACHTransferReturnParamsReasonAccountFrozenEntryReturnedPerOfacInstruction                SimulationACHTransferReturnParamsReason = "account_frozen_entry_returned_per_ofac_instruction"
	SimulationACHTransferReturnParamsReasonCreditEntryRefusedByReceiver                                SimulationACHTransferReturnParamsReason = "credit_entry_refused_by_receiver"
	SimulationACHTransferReturnParamsReasonUnauthorizedDebitToConsumerAccountUsingCorporateSecCode     SimulationACHTransferReturnParamsReason = "unauthorized_debit_to_consumer_account_using_corporate_sec_code"
	SimulationACHTransferReturnParamsReasonCorporateCustomerAdvisedNotAuthorized                       SimulationACHTransferReturnParamsReason = "corporate_customer_advised_not_authorized"
	SimulationACHTransferReturnParamsReasonPaymentStopped                                              SimulationACHTransferReturnParamsReason = "payment_stopped"
	SimulationACHTransferReturnParamsReasonNonTransactionAccount                                       SimulationACHTransferReturnParamsReason = "non_transaction_account"
	SimulationACHTransferReturnParamsReasonUncollectedFunds                                            SimulationACHTransferReturnParamsReason = "uncollected_funds"
	SimulationACHTransferReturnParamsReasonRoutingNumberCheckDigitError                                SimulationACHTransferReturnParamsReason = "routing_number_check_digit_error"
	SimulationACHTransferReturnParamsReasonCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete   SimulationACHTransferReturnParamsReason = "customer_advised_unauthorized_improper_ineligible_or_incomplete"
	SimulationACHTransferReturnParamsReasonAmountFieldError                                            SimulationACHTransferReturnParamsReason = "amount_field_error"
	SimulationACHTransferReturnParamsReasonAuthorizationRevokedByCustomer                              SimulationACHTransferReturnParamsReason = "authorization_revoked_by_customer"
	SimulationACHTransferReturnParamsReasonInvalidACHRoutingNumber                                     SimulationACHTransferReturnParamsReason = "invalid_ach_routing_number"
	SimulationACHTransferReturnParamsReasonFileRecordEditCriteria                                      SimulationACHTransferReturnParamsReason = "file_record_edit_criteria"
	SimulationACHTransferReturnParamsReasonEnrInvalidIndividualName                                    SimulationACHTransferReturnParamsReason = "enr_invalid_individual_name"
	SimulationACHTransferReturnParamsReasonReturnedPerOdfiRequest                                      SimulationACHTransferReturnParamsReason = "returned_per_odfi_request"
	SimulationACHTransferReturnParamsReasonLimitedParticipationDfi                                     SimulationACHTransferReturnParamsReason = "limited_participation_dfi"
	SimulationACHTransferReturnParamsReasonIncorrectlyCodedOutboundInternationalPayment                SimulationACHTransferReturnParamsReason = "incorrectly_coded_outbound_international_payment"
	SimulationACHTransferReturnParamsReasonAccountSoldToAnotherDfi                                     SimulationACHTransferReturnParamsReason = "account_sold_to_another_dfi"
	SimulationACHTransferReturnParamsReasonAddendaError                                                SimulationACHTransferReturnParamsReason = "addenda_error"
	SimulationACHTransferReturnParamsReasonBeneficiaryOrAccountHolderDeceased                          SimulationACHTransferReturnParamsReason = "beneficiary_or_account_holder_deceased"
	SimulationACHTransferReturnParamsReasonCustomerAdvisedNotWithinAuthorizationTerms                  SimulationACHTransferReturnParamsReason = "customer_advised_not_within_authorization_terms"
	SimulationACHTransferReturnParamsReasonCorrectedReturn                                             SimulationACHTransferReturnParamsReason = "corrected_return"
	SimulationACHTransferReturnParamsReasonDuplicateEntry                                              SimulationACHTransferReturnParamsReason = "duplicate_entry"
	SimulationACHTransferReturnParamsReasonDuplicateReturn                                             SimulationACHTransferReturnParamsReason = "duplicate_return"
	SimulationACHTransferReturnParamsReasonEnrDuplicateEnrollment                                      SimulationACHTransferReturnParamsReason = "enr_duplicate_enrollment"
	SimulationACHTransferReturnParamsReasonEnrInvalidDfiAccountNumber                                  SimulationACHTransferReturnParamsReason = "enr_invalid_dfi_account_number"
	SimulationACHTransferReturnParamsReasonEnrInvalidIndividualIDNumber                                SimulationACHTransferReturnParamsReason = "enr_invalid_individual_id_number"
	SimulationACHTransferReturnParamsReasonEnrInvalidRepresentativePayeeIndicator                      SimulationACHTransferReturnParamsReason = "enr_invalid_representative_payee_indicator"
	SimulationACHTransferReturnParamsReasonEnrInvalidTransactionCode                                   SimulationACHTransferReturnParamsReason = "enr_invalid_transaction_code"
	SimulationACHTransferReturnParamsReasonEnrReturnOfEnrEntry                                         SimulationACHTransferReturnParamsReason = "enr_return_of_enr_entry"
	SimulationACHTransferReturnParamsReasonEnrRoutingNumberCheckDigitError                             SimulationACHTransferReturnParamsReason = "enr_routing_number_check_digit_error"
	SimulationACHTransferReturnParamsReasonEntryNotProcessedByGateway                                  SimulationACHTransferReturnParamsReason = "entry_not_processed_by_gateway"
	SimulationACHTransferReturnParamsReasonFieldError                                                  SimulationACHTransferReturnParamsReason = "field_error"
	SimulationACHTransferReturnParamsReasonForeignReceivingDfiUnableToSettle                           SimulationACHTransferReturnParamsReason = "foreign_receiving_dfi_unable_to_settle"
	SimulationACHTransferReturnParamsReasonIatEntryCodingError                                         SimulationACHTransferReturnParamsReason = "iat_entry_coding_error"
	SimulationACHTransferReturnParamsReasonImproperEffectiveEntryDate                                  SimulationACHTransferReturnParamsReason = "improper_effective_entry_date"
	SimulationACHTransferReturnParamsReasonImproperSourceDocumentSourceDocumentPresented               SimulationACHTransferReturnParamsReason = "improper_source_document_source_document_presented"
	SimulationACHTransferReturnParamsReasonInvalidCompanyID                                            SimulationACHTransferReturnParamsReason = "invalid_company_id"
	SimulationACHTransferReturnParamsReasonInvalidForeignReceivingDfiIdentification                    SimulationACHTransferReturnParamsReason = "invalid_foreign_receiving_dfi_identification"
	SimulationACHTransferReturnParamsReasonInvalidIndividualIDNumber                                   SimulationACHTransferReturnParamsReason = "invalid_individual_id_number"
	SimulationACHTransferReturnParamsReasonItemAndRckEntryPresentedForPayment                          SimulationACHTransferReturnParamsReason = "item_and_rck_entry_presented_for_payment"
	SimulationACHTransferReturnParamsReasonItemRelatedToRckEntryIsIneligible                           SimulationACHTransferReturnParamsReason = "item_related_to_rck_entry_is_ineligible"
	SimulationACHTransferReturnParamsReasonMandatoryFieldError                                         SimulationACHTransferReturnParamsReason = "mandatory_field_error"
	SimulationACHTransferReturnParamsReasonMisroutedDishonoredReturn                                   SimulationACHTransferReturnParamsReason = "misrouted_dishonored_return"
	SimulationACHTransferReturnParamsReasonMisroutedReturn                                             SimulationACHTransferReturnParamsReason = "misrouted_return"
	SimulationACHTransferReturnParamsReasonNoErrorsFound                                               SimulationACHTransferReturnParamsReason = "no_errors_found"
	SimulationACHTransferReturnParamsReasonNonAcceptanceOfR62DishonoredReturn                          SimulationACHTransferReturnParamsReason = "non_acceptance_of_r62_dishonored_return"
	SimulationACHTransferReturnParamsReasonNonParticipantInIatProgram                                  SimulationACHTransferReturnParamsReason = "non_participant_in_iat_program"
	SimulationACHTransferReturnParamsReasonPermissibleReturnEntry                                      SimulationACHTransferReturnParamsReason = "permissible_return_entry"
	SimulationACHTransferReturnParamsReasonPermissibleReturnEntryNotAccepted                           SimulationACHTransferReturnParamsReason = "permissible_return_entry_not_accepted"
	SimulationACHTransferReturnParamsReasonRdfiNonSettlement                                           SimulationACHTransferReturnParamsReason = "rdfi_non_settlement"
	SimulationACHTransferReturnParamsReasonRdfiParticipantInCheckTruncationProgram                     SimulationACHTransferReturnParamsReason = "rdfi_participant_in_check_truncation_program"
	SimulationACHTransferReturnParamsReasonRepresentativePayeeDeceasedOrUnableToContinueInThatCapacity SimulationACHTransferReturnParamsReason = "representative_payee_deceased_or_unable_to_continue_in_that_capacity"
	SimulationACHTransferReturnParamsReasonReturnNotADuplicate                                         SimulationACHTransferReturnParamsReason = "return_not_a_duplicate"
	SimulationACHTransferReturnParamsReasonReturnOfErroneousOrReversingDebit                           SimulationACHTransferReturnParamsReason = "return_of_erroneous_or_reversing_debit"
	SimulationACHTransferReturnParamsReasonReturnOfImproperCreditEntry                                 SimulationACHTransferReturnParamsReason = "return_of_improper_credit_entry"
	SimulationACHTransferReturnParamsReasonReturnOfImproperDebitEntry                                  SimulationACHTransferReturnParamsReason = "return_of_improper_debit_entry"
	SimulationACHTransferReturnParamsReasonReturnOfXckEntry                                            SimulationACHTransferReturnParamsReason = "return_of_xck_entry"
	SimulationACHTransferReturnParamsReasonSourceDocumentPresentedForPayment                           SimulationACHTransferReturnParamsReason = "source_document_presented_for_payment"
	SimulationACHTransferReturnParamsReasonStateLawAffectingRckAcceptance                              SimulationACHTransferReturnParamsReason = "state_law_affecting_rck_acceptance"
	SimulationACHTransferReturnParamsReasonStopPaymentOnItemRelatedToRckEntry                          SimulationACHTransferReturnParamsReason = "stop_payment_on_item_related_to_rck_entry"
	SimulationACHTransferReturnParamsReasonStopPaymentOnSourceDocument                                 SimulationACHTransferReturnParamsReason = "stop_payment_on_source_document"
	SimulationACHTransferReturnParamsReasonTimelyOriginalReturn                                        SimulationACHTransferReturnParamsReason = "timely_original_return"
	SimulationACHTransferReturnParamsReasonTraceNumberError                                            SimulationACHTransferReturnParamsReason = "trace_number_error"
	SimulationACHTransferReturnParamsReasonUntimelyDishonoredReturn                                    SimulationACHTransferReturnParamsReason = "untimely_dishonored_return"
	SimulationACHTransferReturnParamsReasonUntimelyReturn                                              SimulationACHTransferReturnParamsReason = "untimely_return"
)

func (r SimulationACHTransferReturnParamsReason) IsKnown() bool {
	switch r {
	case SimulationACHTransferReturnParamsReasonInsufficientFund, SimulationACHTransferReturnParamsReasonNoAccount, SimulationACHTransferReturnParamsReasonAccountClosed, SimulationACHTransferReturnParamsReasonInvalidAccountNumberStructure, SimulationACHTransferReturnParamsReasonAccountFrozenEntryReturnedPerOfacInstruction, SimulationACHTransferReturnParamsReasonCreditEntryRefusedByReceiver, SimulationACHTransferReturnParamsReasonUnauthorizedDebitToConsumerAccountUsingCorporateSecCode, SimulationACHTransferReturnParamsReasonCorporateCustomerAdvisedNotAuthorized, SimulationACHTransferReturnParamsReasonPaymentStopped, SimulationACHTransferReturnParamsReasonNonTransactionAccount, SimulationACHTransferReturnParamsReasonUncollectedFunds, SimulationACHTransferReturnParamsReasonRoutingNumberCheckDigitError, SimulationACHTransferReturnParamsReasonCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete, SimulationACHTransferReturnParamsReasonAmountFieldError, SimulationACHTransferReturnParamsReasonAuthorizationRevokedByCustomer, SimulationACHTransferReturnParamsReasonInvalidACHRoutingNumber, SimulationACHTransferReturnParamsReasonFileRecordEditCriteria, SimulationACHTransferReturnParamsReasonEnrInvalidIndividualName, SimulationACHTransferReturnParamsReasonReturnedPerOdfiRequest, SimulationACHTransferReturnParamsReasonLimitedParticipationDfi, SimulationACHTransferReturnParamsReasonIncorrectlyCodedOutboundInternationalPayment, SimulationACHTransferReturnParamsReasonAccountSoldToAnotherDfi, SimulationACHTransferReturnParamsReasonAddendaError, SimulationACHTransferReturnParamsReasonBeneficiaryOrAccountHolderDeceased, SimulationACHTransferReturnParamsReasonCustomerAdvisedNotWithinAuthorizationTerms, SimulationACHTransferReturnParamsReasonCorrectedReturn, SimulationACHTransferReturnParamsReasonDuplicateEntry, SimulationACHTransferReturnParamsReasonDuplicateReturn, SimulationACHTransferReturnParamsReasonEnrDuplicateEnrollment, SimulationACHTransferReturnParamsReasonEnrInvalidDfiAccountNumber, SimulationACHTransferReturnParamsReasonEnrInvalidIndividualIDNumber, SimulationACHTransferReturnParamsReasonEnrInvalidRepresentativePayeeIndicator, SimulationACHTransferReturnParamsReasonEnrInvalidTransactionCode, SimulationACHTransferReturnParamsReasonEnrReturnOfEnrEntry, SimulationACHTransferReturnParamsReasonEnrRoutingNumberCheckDigitError, SimulationACHTransferReturnParamsReasonEntryNotProcessedByGateway, SimulationACHTransferReturnParamsReasonFieldError, SimulationACHTransferReturnParamsReasonForeignReceivingDfiUnableToSettle, SimulationACHTransferReturnParamsReasonIatEntryCodingError, SimulationACHTransferReturnParamsReasonImproperEffectiveEntryDate, SimulationACHTransferReturnParamsReasonImproperSourceDocumentSourceDocumentPresented, SimulationACHTransferReturnParamsReasonInvalidCompanyID, SimulationACHTransferReturnParamsReasonInvalidForeignReceivingDfiIdentification, SimulationACHTransferReturnParamsReasonInvalidIndividualIDNumber, SimulationACHTransferReturnParamsReasonItemAndRckEntryPresentedForPayment, SimulationACHTransferReturnParamsReasonItemRelatedToRckEntryIsIneligible, SimulationACHTransferReturnParamsReasonMandatoryFieldError, SimulationACHTransferReturnParamsReasonMisroutedDishonoredReturn, SimulationACHTransferReturnParamsReasonMisroutedReturn, SimulationACHTransferReturnParamsReasonNoErrorsFound, SimulationACHTransferReturnParamsReasonNonAcceptanceOfR62DishonoredReturn, SimulationACHTransferReturnParamsReasonNonParticipantInIatProgram, SimulationACHTransferReturnParamsReasonPermissibleReturnEntry, SimulationACHTransferReturnParamsReasonPermissibleReturnEntryNotAccepted, SimulationACHTransferReturnParamsReasonRdfiNonSettlement, SimulationACHTransferReturnParamsReasonRdfiParticipantInCheckTruncationProgram, SimulationACHTransferReturnParamsReasonRepresentativePayeeDeceasedOrUnableToContinueInThatCapacity, SimulationACHTransferReturnParamsReasonReturnNotADuplicate, SimulationACHTransferReturnParamsReasonReturnOfErroneousOrReversingDebit, SimulationACHTransferReturnParamsReasonReturnOfImproperCreditEntry, SimulationACHTransferReturnParamsReasonReturnOfImproperDebitEntry, SimulationACHTransferReturnParamsReasonReturnOfXckEntry, SimulationACHTransferReturnParamsReasonSourceDocumentPresentedForPayment, SimulationACHTransferReturnParamsReasonStateLawAffectingRckAcceptance, SimulationACHTransferReturnParamsReasonStopPaymentOnItemRelatedToRckEntry, SimulationACHTransferReturnParamsReasonStopPaymentOnSourceDocument, SimulationACHTransferReturnParamsReasonTimelyOriginalReturn, SimulationACHTransferReturnParamsReasonTraceNumberError, SimulationACHTransferReturnParamsReasonUntimelyDishonoredReturn, SimulationACHTransferReturnParamsReasonUntimelyReturn:
		return true
	}
	return false
}

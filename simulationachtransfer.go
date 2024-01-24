// File generated from our OpenAPI spec by Stainless.

package increase

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/option"
)

// SimulationACHTransferService contains methods and other services that help with
// interacting with the increase API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSimulationACHTransferService]
// method instead.
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

// Simulates an inbound ACH transfer to your account. This imitates initiating a
// transfer to an Increase account from a different financial institution. The
// transfer may be either a credit or a debit depending on if the `amount` is
// positive or negative. The result of calling this API will contain the created
// transfer. You can pass a `resolve_at` parameter to allow for a window to
// [action on the Inbound ACH Transfer](https://increase.com/documentation/inbound-ach-transfers#inbound-ach-transfers).
// Alternatively, if you don't pass the `resolve_at` parameter the result will
// contain either a [Transaction](#transactions) or a
// [Declined Transaction](#declined-transactions) depending on whether or not the
// transfer is allowed.
func (r *SimulationACHTransferService) NewInbound(ctx context.Context, body SimulationACHTransferNewInboundParams, opts ...option.RequestOption) (res *InboundACHTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/inbound_ach_transfers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Simulates the return of an [ACH Transfer](#ach-transfers) by the Federal Reserve
// due to an error condition. This will also create a Transaction to account for
// the returned funds. This transfer must first have a `status` of `submitted`.
func (r *SimulationACHTransferService) Return(ctx context.Context, achTransferID string, body SimulationACHTransferReturnParams, opts ...option.RequestOption) (res *ACHTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("simulations/ach_transfers/%s/return", achTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
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
	path := fmt.Sprintf("simulations/ach_transfers/%s/submit", achTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type SimulationACHTransferNewInboundParams struct {
	// The identifier of the Account Number the inbound ACH Transfer is for.
	AccountNumberID param.Field[string] `json:"account_number_id,required"`
	// The transfer amount in cents. A positive amount originates a credit transfer
	// pushing funds to the receiving account. A negative amount originates a debit
	// transfer pulling funds from the receiving account.
	Amount param.Field[int64] `json:"amount,required"`
	// The description of the date of the transfer.
	CompanyDescriptiveDate param.Field[string] `json:"company_descriptive_date"`
	// Data associated with the transfer set by the sender.
	CompanyDiscretionaryData param.Field[string] `json:"company_discretionary_data"`
	// The description of the transfer set by the sender.
	CompanyEntryDescription param.Field[string] `json:"company_entry_description"`
	// The sender's company id.
	CompanyID param.Field[string] `json:"company_id"`
	// The name of the sender.
	CompanyName param.Field[string] `json:"company_name"`
	// The time at which the transfer should be resolved. If not provided will resolve
	// immediately.
	ResolveAt param.Field[time.Time] `json:"resolve_at" format:"date-time"`
}

func (r SimulationACHTransferNewInboundParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	// Code R01. Insufficient funds in the receiving account. Sometimes abbreviated to
	// NSF.
	SimulationACHTransferReturnParamsReasonInsufficientFund SimulationACHTransferReturnParamsReason = "insufficient_fund"
	// Code R03. The account does not exist or the receiving bank was unable to locate
	// it.
	SimulationACHTransferReturnParamsReasonNoAccount SimulationACHTransferReturnParamsReason = "no_account"
	// Code R02. The account is closed at the receiving bank.
	SimulationACHTransferReturnParamsReasonAccountClosed SimulationACHTransferReturnParamsReason = "account_closed"
	// Code R04. The account number is invalid at the receiving bank.
	SimulationACHTransferReturnParamsReasonInvalidAccountNumberStructure SimulationACHTransferReturnParamsReason = "invalid_account_number_structure"
	// Code R16. The account at the receiving bank was frozen per the Office of Foreign
	// Assets Control.
	SimulationACHTransferReturnParamsReasonAccountFrozenEntryReturnedPerOfacInstruction SimulationACHTransferReturnParamsReason = "account_frozen_entry_returned_per_ofac_instruction"
	// Code R23. The receiving bank account refused a credit transfer.
	SimulationACHTransferReturnParamsReasonCreditEntryRefusedByReceiver SimulationACHTransferReturnParamsReason = "credit_entry_refused_by_receiver"
	// Code R05. The receiving bank rejected because of an incorrect Standard Entry
	// Class code.
	SimulationACHTransferReturnParamsReasonUnauthorizedDebitToConsumerAccountUsingCorporateSecCode SimulationACHTransferReturnParamsReason = "unauthorized_debit_to_consumer_account_using_corporate_sec_code"
	// Code R29. The corporate customer at the receiving bank reversed the transfer.
	SimulationACHTransferReturnParamsReasonCorporateCustomerAdvisedNotAuthorized SimulationACHTransferReturnParamsReason = "corporate_customer_advised_not_authorized"
	// Code R08. The receiving bank stopped payment on this transfer.
	SimulationACHTransferReturnParamsReasonPaymentStopped SimulationACHTransferReturnParamsReason = "payment_stopped"
	// Code R20. The receiving bank account does not perform transfers.
	SimulationACHTransferReturnParamsReasonNonTransactionAccount SimulationACHTransferReturnParamsReason = "non_transaction_account"
	// Code R09. The receiving bank account does not have enough available balance for
	// the transfer.
	SimulationACHTransferReturnParamsReasonUncollectedFunds SimulationACHTransferReturnParamsReason = "uncollected_funds"
	// Code R28. The routing number is incorrect.
	SimulationACHTransferReturnParamsReasonRoutingNumberCheckDigitError SimulationACHTransferReturnParamsReason = "routing_number_check_digit_error"
	// Code R10. The customer at the receiving bank reversed the transfer.
	SimulationACHTransferReturnParamsReasonCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete SimulationACHTransferReturnParamsReason = "customer_advised_unauthorized_improper_ineligible_or_incomplete"
	// Code R19. The amount field is incorrect or too large.
	SimulationACHTransferReturnParamsReasonAmountFieldError SimulationACHTransferReturnParamsReason = "amount_field_error"
	// Code R07. The customer at the receiving institution informed their bank that
	// they have revoked authorization for a previously authorized transfer.
	SimulationACHTransferReturnParamsReasonAuthorizationRevokedByCustomer SimulationACHTransferReturnParamsReason = "authorization_revoked_by_customer"
	// Code R13. The routing number is invalid.
	SimulationACHTransferReturnParamsReasonInvalidACHRoutingNumber SimulationACHTransferReturnParamsReason = "invalid_ach_routing_number"
	// Code R17. The receiving bank is unable to process a field in the transfer.
	SimulationACHTransferReturnParamsReasonFileRecordEditCriteria SimulationACHTransferReturnParamsReason = "file_record_edit_criteria"
	// Code R45. The individual name field was invalid.
	SimulationACHTransferReturnParamsReasonEnrInvalidIndividualName SimulationACHTransferReturnParamsReason = "enr_invalid_individual_name"
	// Code R06. The originating financial institution asked for this transfer to be
	// returned. The receiving bank is complying with the request.
	SimulationACHTransferReturnParamsReasonReturnedPerOdfiRequest SimulationACHTransferReturnParamsReason = "returned_per_odfi_request"
	// Code R34. The receiving bank's regulatory supervisor has limited their
	// participation in the ACH network.
	SimulationACHTransferReturnParamsReasonLimitedParticipationDfi SimulationACHTransferReturnParamsReason = "limited_participation_dfi"
	// Code R85. The outbound international ACH transfer was incorrect.
	SimulationACHTransferReturnParamsReasonIncorrectlyCodedOutboundInternationalPayment SimulationACHTransferReturnParamsReason = "incorrectly_coded_outbound_international_payment"
	// Code R12. A rare return reason. The account was sold to another bank.
	SimulationACHTransferReturnParamsReasonAccountSoldToAnotherDfi SimulationACHTransferReturnParamsReason = "account_sold_to_another_dfi"
	// Code R25. The addenda record is incorrect or missing.
	SimulationACHTransferReturnParamsReasonAddendaError SimulationACHTransferReturnParamsReason = "addenda_error"
	// Code R15. A rare return reason. The account holder is deceased.
	SimulationACHTransferReturnParamsReasonBeneficiaryOrAccountHolderDeceased SimulationACHTransferReturnParamsReason = "beneficiary_or_account_holder_deceased"
	// Code R11. A rare return reason. The customer authorized some payment to the
	// sender, but this payment was not in error.
	SimulationACHTransferReturnParamsReasonCustomerAdvisedNotWithinAuthorizationTerms SimulationACHTransferReturnParamsReason = "customer_advised_not_within_authorization_terms"
	// Code R74. A rare return reason. Sent in response to a return that was returned
	// with code `field_error`. The latest return should include the corrected
	// field(s).
	SimulationACHTransferReturnParamsReasonCorrectedReturn SimulationACHTransferReturnParamsReason = "corrected_return"
	// Code R24. A rare return reason. The receiving bank received an exact duplicate
	// entry with the same trace number and amount.
	SimulationACHTransferReturnParamsReasonDuplicateEntry SimulationACHTransferReturnParamsReason = "duplicate_entry"
	// Code R67. A rare return reason. The return this message refers to was a
	// duplicate.
	SimulationACHTransferReturnParamsReasonDuplicateReturn SimulationACHTransferReturnParamsReason = "duplicate_return"
	// Code R47. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	SimulationACHTransferReturnParamsReasonEnrDuplicateEnrollment SimulationACHTransferReturnParamsReason = "enr_duplicate_enrollment"
	// Code R43. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	SimulationACHTransferReturnParamsReasonEnrInvalidDfiAccountNumber SimulationACHTransferReturnParamsReason = "enr_invalid_dfi_account_number"
	// Code R44. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	SimulationACHTransferReturnParamsReasonEnrInvalidIndividualIDNumber SimulationACHTransferReturnParamsReason = "enr_invalid_individual_id_number"
	// Code R46. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	SimulationACHTransferReturnParamsReasonEnrInvalidRepresentativePayeeIndicator SimulationACHTransferReturnParamsReason = "enr_invalid_representative_payee_indicator"
	// Code R41. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	SimulationACHTransferReturnParamsReasonEnrInvalidTransactionCode SimulationACHTransferReturnParamsReason = "enr_invalid_transaction_code"
	// Code R40. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	SimulationACHTransferReturnParamsReasonEnrReturnOfEnrEntry SimulationACHTransferReturnParamsReason = "enr_return_of_enr_entry"
	// Code R42. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	SimulationACHTransferReturnParamsReasonEnrRoutingNumberCheckDigitError SimulationACHTransferReturnParamsReason = "enr_routing_number_check_digit_error"
	// Code R84. A rare return reason. The International ACH Transfer cannot be
	// processed by the gateway.
	SimulationACHTransferReturnParamsReasonEntryNotProcessedByGateway SimulationACHTransferReturnParamsReason = "entry_not_processed_by_gateway"
	// Code R69. A rare return reason. One or more of the fields in the ACH were
	// malformed.
	SimulationACHTransferReturnParamsReasonFieldError SimulationACHTransferReturnParamsReason = "field_error"
	// Code R83. A rare return reason. The Foreign receiving bank was unable to settle
	// this ACH transfer.
	SimulationACHTransferReturnParamsReasonForeignReceivingDfiUnableToSettle SimulationACHTransferReturnParamsReason = "foreign_receiving_dfi_unable_to_settle"
	// Code R80. A rare return reason. The International ACH Transfer is malformed.
	SimulationACHTransferReturnParamsReasonIatEntryCodingError SimulationACHTransferReturnParamsReason = "iat_entry_coding_error"
	// Code R18. A rare return reason. The ACH has an improper effective entry date
	// field.
	SimulationACHTransferReturnParamsReasonImproperEffectiveEntryDate SimulationACHTransferReturnParamsReason = "improper_effective_entry_date"
	// Code R39. A rare return reason. The source document related to this ACH, usually
	// an ACH check conversion, was presented to the bank.
	SimulationACHTransferReturnParamsReasonImproperSourceDocumentSourceDocumentPresented SimulationACHTransferReturnParamsReason = "improper_source_document_source_document_presented"
	// Code R21. A rare return reason. The Company ID field of the ACH was invalid.
	SimulationACHTransferReturnParamsReasonInvalidCompanyID SimulationACHTransferReturnParamsReason = "invalid_company_id"
	// Code R82. A rare return reason. The foreign receiving bank identifier for an
	// International ACH Transfer was invalid.
	SimulationACHTransferReturnParamsReasonInvalidForeignReceivingDfiIdentification SimulationACHTransferReturnParamsReason = "invalid_foreign_receiving_dfi_identification"
	// Code R22. A rare return reason. The Individual ID number field of the ACH was
	// invalid.
	SimulationACHTransferReturnParamsReasonInvalidIndividualIDNumber SimulationACHTransferReturnParamsReason = "invalid_individual_id_number"
	// Code R53. A rare return reason. Both the Represented Check ("RCK") entry and the
	// original check were presented to the bank.
	SimulationACHTransferReturnParamsReasonItemAndRckEntryPresentedForPayment SimulationACHTransferReturnParamsReason = "item_and_rck_entry_presented_for_payment"
	// Code R51. A rare return reason. The Represented Check ("RCK") entry is
	// ineligible.
	SimulationACHTransferReturnParamsReasonItemRelatedToRckEntryIsIneligible SimulationACHTransferReturnParamsReason = "item_related_to_rck_entry_is_ineligible"
	// Code R26. A rare return reason. The ACH is missing a required field.
	SimulationACHTransferReturnParamsReasonMandatoryFieldError SimulationACHTransferReturnParamsReason = "mandatory_field_error"
	// Code R71. A rare return reason. The receiving bank does not recognize the
	// routing number in a dishonored return entry.
	SimulationACHTransferReturnParamsReasonMisroutedDishonoredReturn SimulationACHTransferReturnParamsReason = "misrouted_dishonored_return"
	// Code R61. A rare return reason. The receiving bank does not recognize the
	// routing number in a return entry.
	SimulationACHTransferReturnParamsReasonMisroutedReturn SimulationACHTransferReturnParamsReason = "misrouted_return"
	// Code R76. A rare return reason. Sent in response to a return, the bank does not
	// find the errors alleged by the returning bank.
	SimulationACHTransferReturnParamsReasonNoErrorsFound SimulationACHTransferReturnParamsReason = "no_errors_found"
	// Code R77. A rare return reason. The receiving bank does not accept the return of
	// the erroneous debit. The funds are not available at the receiving bank.
	SimulationACHTransferReturnParamsReasonNonAcceptanceOfR62DishonoredReturn SimulationACHTransferReturnParamsReason = "non_acceptance_of_r62_dishonored_return"
	// Code R81. A rare return reason. The receiving bank does not accept International
	// ACH Transfers.
	SimulationACHTransferReturnParamsReasonNonParticipantInIatProgram SimulationACHTransferReturnParamsReason = "non_participant_in_iat_program"
	// Code R31. A rare return reason. A return that has been agreed to be accepted by
	// the receiving bank, despite falling outside of the usual return timeframe.
	SimulationACHTransferReturnParamsReasonPermissibleReturnEntry SimulationACHTransferReturnParamsReason = "permissible_return_entry"
	// Code R70. A rare return reason. The receiving bank had not approved this return.
	SimulationACHTransferReturnParamsReasonPermissibleReturnEntryNotAccepted SimulationACHTransferReturnParamsReason = "permissible_return_entry_not_accepted"
	// Code R32. A rare return reason. The receiving bank could not settle this
	// transaction.
	SimulationACHTransferReturnParamsReasonRdfiNonSettlement SimulationACHTransferReturnParamsReason = "rdfi_non_settlement"
	// Code R30. A rare return reason. The receiving bank does not accept Check
	// Truncation ACH transfers.
	SimulationACHTransferReturnParamsReasonRdfiParticipantInCheckTruncationProgram SimulationACHTransferReturnParamsReason = "rdfi_participant_in_check_truncation_program"
	// Code R14. A rare return reason. The payee is deceased.
	SimulationACHTransferReturnParamsReasonRepresentativePayeeDeceasedOrUnableToContinueInThatCapacity SimulationACHTransferReturnParamsReason = "representative_payee_deceased_or_unable_to_continue_in_that_capacity"
	// Code R75. A rare return reason. The originating bank disputes that an earlier
	// `duplicate_entry` return was actually a duplicate.
	SimulationACHTransferReturnParamsReasonReturnNotADuplicate SimulationACHTransferReturnParamsReason = "return_not_a_duplicate"
	// Code R62. A rare return reason. The originating financial institution made a
	// mistake and this return corrects it.
	SimulationACHTransferReturnParamsReasonReturnOfErroneousOrReversingDebit SimulationACHTransferReturnParamsReason = "return_of_erroneous_or_reversing_debit"
	// Code R36. A rare return reason. Return of a malformed credit entry.
	SimulationACHTransferReturnParamsReasonReturnOfImproperCreditEntry SimulationACHTransferReturnParamsReason = "return_of_improper_credit_entry"
	// Code R35. A rare return reason. Return of a malformed debit entry.
	SimulationACHTransferReturnParamsReasonReturnOfImproperDebitEntry SimulationACHTransferReturnParamsReason = "return_of_improper_debit_entry"
	// Code R33. A rare return reason. Return of a Destroyed Check ("XKC") entry.
	SimulationACHTransferReturnParamsReasonReturnOfXckEntry SimulationACHTransferReturnParamsReason = "return_of_xck_entry"
	// Code R37. A rare return reason. The source document related to this ACH, usually
	// an ACH check conversion, was presented to the bank.
	SimulationACHTransferReturnParamsReasonSourceDocumentPresentedForPayment SimulationACHTransferReturnParamsReason = "source_document_presented_for_payment"
	// Code R50. A rare return reason. State law prevents the bank from accepting the
	// Represented Check ("RCK") entry.
	SimulationACHTransferReturnParamsReasonStateLawAffectingRckAcceptance SimulationACHTransferReturnParamsReason = "state_law_affecting_rck_acceptance"
	// Code R52. A rare return reason. A stop payment was issued on a Represented Check
	// ("RCK") entry.
	SimulationACHTransferReturnParamsReasonStopPaymentOnItemRelatedToRckEntry SimulationACHTransferReturnParamsReason = "stop_payment_on_item_related_to_rck_entry"
	// Code R38. A rare return reason. The source attached to the ACH, usually an ACH
	// check conversion, includes a stop payment.
	SimulationACHTransferReturnParamsReasonStopPaymentOnSourceDocument SimulationACHTransferReturnParamsReason = "stop_payment_on_source_document"
	// Code R73. A rare return reason. The bank receiving an `untimely_return` believes
	// it was on time.
	SimulationACHTransferReturnParamsReasonTimelyOriginalReturn SimulationACHTransferReturnParamsReason = "timely_original_return"
	// Code R27. A rare return reason. An ACH return's trace number does not match an
	// originated ACH.
	SimulationACHTransferReturnParamsReasonTraceNumberError SimulationACHTransferReturnParamsReason = "trace_number_error"
	// Code R72. A rare return reason. The dishonored return was sent too late.
	SimulationACHTransferReturnParamsReasonUntimelyDishonoredReturn SimulationACHTransferReturnParamsReason = "untimely_dishonored_return"
	// Code R68. A rare return reason. The return was sent too late.
	SimulationACHTransferReturnParamsReasonUntimelyReturn SimulationACHTransferReturnParamsReason = "untimely_return"
)

package requests

import (
	"github.com/increase/increase-go/internal/field"
	apijson "github.com/increase/increase-go/internal/json"
)

type SimulationACHTransferNewInboundParams struct {
	// The identifier of the Account Number the inbound ACH Transfer is for.
	AccountNumberID field.Field[string] `json:"account_number_id,required"`
	// The transfer amount in cents. A positive amount originates a credit transfer
	// pushing funds to the receiving account. A negative amount originates a debit
	// transfer pulling funds from the receiving account.
	Amount field.Field[int64] `json:"amount,required"`
	// The description of the date of the transfer.
	CompanyDescriptiveDate field.Field[string] `json:"company_descriptive_date"`
	// Data associated with the transfer set by the sender.
	CompanyDiscretionaryData field.Field[string] `json:"company_discretionary_data"`
	// The description of the transfer set by the sender.
	CompanyEntryDescription field.Field[string] `json:"company_entry_description"`
	// The name of the sender.
	CompanyName field.Field[string] `json:"company_name"`
	// The sender's company id.
	CompanyID field.Field[string] `json:"company_id"`
}

// MarshalJSON serializes SimulationACHTransferNewInboundParams into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r SimulationACHTransferNewInboundParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SimulationACHTransferReturnParams struct {
	// The reason why the Federal Reserve or destination bank returned this transfer.
	// Defaults to `no_account`.
	Reason field.Field[SimulationACHTransferReturnParamsReason] `json:"reason"`
}

// MarshalJSON serializes SimulationACHTransferReturnParams into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r SimulationACHTransferReturnParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SimulationACHTransferReturnParamsReason string

const (
	SimulationACHTransferReturnParamsReasonInsufficientFund                                          SimulationACHTransferReturnParamsReason = "insufficient_fund"
	SimulationACHTransferReturnParamsReasonNoAccount                                                 SimulationACHTransferReturnParamsReason = "no_account"
	SimulationACHTransferReturnParamsReasonAccountClosed                                             SimulationACHTransferReturnParamsReason = "account_closed"
	SimulationACHTransferReturnParamsReasonInvalidAccountNumberStructure                             SimulationACHTransferReturnParamsReason = "invalid_account_number_structure"
	SimulationACHTransferReturnParamsReasonAccountFrozenEntryReturnedPerOfacInstruction              SimulationACHTransferReturnParamsReason = "account_frozen_entry_returned_per_ofac_instruction"
	SimulationACHTransferReturnParamsReasonCreditEntryRefusedByReceiver                              SimulationACHTransferReturnParamsReason = "credit_entry_refused_by_receiver"
	SimulationACHTransferReturnParamsReasonUnauthorizedDebitToConsumerAccountUsingCorporateSecCode   SimulationACHTransferReturnParamsReason = "unauthorized_debit_to_consumer_account_using_corporate_sec_code"
	SimulationACHTransferReturnParamsReasonCorporateCustomerAdvisedNotAuthorized                     SimulationACHTransferReturnParamsReason = "corporate_customer_advised_not_authorized"
	SimulationACHTransferReturnParamsReasonPaymentStopped                                            SimulationACHTransferReturnParamsReason = "payment_stopped"
	SimulationACHTransferReturnParamsReasonNonTransactionAccount                                     SimulationACHTransferReturnParamsReason = "non_transaction_account"
	SimulationACHTransferReturnParamsReasonUncollectedFunds                                          SimulationACHTransferReturnParamsReason = "uncollected_funds"
	SimulationACHTransferReturnParamsReasonRoutingNumberCheckDigitError                              SimulationACHTransferReturnParamsReason = "routing_number_check_digit_error"
	SimulationACHTransferReturnParamsReasonCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete SimulationACHTransferReturnParamsReason = "customer_advised_unauthorized_improper_ineligible_or_incomplete"
	SimulationACHTransferReturnParamsReasonAmountFieldError                                          SimulationACHTransferReturnParamsReason = "amount_field_error"
	SimulationACHTransferReturnParamsReasonAuthorizationRevokedByCustomer                            SimulationACHTransferReturnParamsReason = "authorization_revoked_by_customer"
	SimulationACHTransferReturnParamsReasonInvalidACHRoutingNumber                                   SimulationACHTransferReturnParamsReason = "invalid_ach_routing_number"
	SimulationACHTransferReturnParamsReasonFileRecordEditCriteria                                    SimulationACHTransferReturnParamsReason = "file_record_edit_criteria"
	SimulationACHTransferReturnParamsReasonEnrInvalidIndividualName                                  SimulationACHTransferReturnParamsReason = "enr_invalid_individual_name"
	SimulationACHTransferReturnParamsReasonReturnedPerOdfiRequest                                    SimulationACHTransferReturnParamsReason = "returned_per_odfi_request"
	SimulationACHTransferReturnParamsReasonAddendaError                                              SimulationACHTransferReturnParamsReason = "addenda_error"
	SimulationACHTransferReturnParamsReasonLimitedParticipationDfi                                   SimulationACHTransferReturnParamsReason = "limited_participation_dfi"
	SimulationACHTransferReturnParamsReasonIncorrectlyCodedOutboundInternationalPayment              SimulationACHTransferReturnParamsReason = "incorrectly_coded_outbound_international_payment"
	SimulationACHTransferReturnParamsReasonOther                                                     SimulationACHTransferReturnParamsReason = "other"
)

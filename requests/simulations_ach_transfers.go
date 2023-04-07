package requests

import (
	"github.com/increase/increase-go/core/field"
	pjson "github.com/increase/increase-go/core/json"
)

type ACHTransferNewInboundParams struct {
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

// MarshalJSON serializes ACHTransferNewInboundParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *ACHTransferNewInboundParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type ACHTransferReturnParams struct {
	// The reason why the Federal Reserve or destination bank returned this transfer.
	// Defaults to `no_account`.
	Reason field.Field[ACHTransferReturnParamsReason] `json:"reason"`
}

// MarshalJSON serializes ACHTransferReturnParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *ACHTransferReturnParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type ACHTransferReturnParamsReason string

const (
	ACHTransferReturnParamsReasonInsufficientFund                                          ACHTransferReturnParamsReason = "insufficient_fund"
	ACHTransferReturnParamsReasonNoAccount                                                 ACHTransferReturnParamsReason = "no_account"
	ACHTransferReturnParamsReasonAccountClosed                                             ACHTransferReturnParamsReason = "account_closed"
	ACHTransferReturnParamsReasonInvalidAccountNumberStructure                             ACHTransferReturnParamsReason = "invalid_account_number_structure"
	ACHTransferReturnParamsReasonAccountFrozenEntryReturnedPerOfacInstruction              ACHTransferReturnParamsReason = "account_frozen_entry_returned_per_ofac_instruction"
	ACHTransferReturnParamsReasonCreditEntryRefusedByReceiver                              ACHTransferReturnParamsReason = "credit_entry_refused_by_receiver"
	ACHTransferReturnParamsReasonUnauthorizedDebitToConsumerAccountUsingCorporateSecCode   ACHTransferReturnParamsReason = "unauthorized_debit_to_consumer_account_using_corporate_sec_code"
	ACHTransferReturnParamsReasonCorporateCustomerAdvisedNotAuthorized                     ACHTransferReturnParamsReason = "corporate_customer_advised_not_authorized"
	ACHTransferReturnParamsReasonPaymentStopped                                            ACHTransferReturnParamsReason = "payment_stopped"
	ACHTransferReturnParamsReasonNonTransactionAccount                                     ACHTransferReturnParamsReason = "non_transaction_account"
	ACHTransferReturnParamsReasonUncollectedFunds                                          ACHTransferReturnParamsReason = "uncollected_funds"
	ACHTransferReturnParamsReasonRoutingNumberCheckDigitError                              ACHTransferReturnParamsReason = "routing_number_check_digit_error"
	ACHTransferReturnParamsReasonCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete ACHTransferReturnParamsReason = "customer_advised_unauthorized_improper_ineligible_or_incomplete"
	ACHTransferReturnParamsReasonAmountFieldError                                          ACHTransferReturnParamsReason = "amount_field_error"
	ACHTransferReturnParamsReasonAuthorizationRevokedByCustomer                            ACHTransferReturnParamsReason = "authorization_revoked_by_customer"
	ACHTransferReturnParamsReasonInvalidACHRoutingNumber                                   ACHTransferReturnParamsReason = "invalid_ach_routing_number"
	ACHTransferReturnParamsReasonFileRecordEditCriteria                                    ACHTransferReturnParamsReason = "file_record_edit_criteria"
	ACHTransferReturnParamsReasonEnrInvalidIndividualName                                  ACHTransferReturnParamsReason = "enr_invalid_individual_name"
	ACHTransferReturnParamsReasonReturnedPerOdfiRequest                                    ACHTransferReturnParamsReason = "returned_per_odfi_request"
	ACHTransferReturnParamsReasonAddendaError                                              ACHTransferReturnParamsReason = "addenda_error"
	ACHTransferReturnParamsReasonLimitedParticipationDfi                                   ACHTransferReturnParamsReason = "limited_participation_dfi"
	ACHTransferReturnParamsReasonIncorrectlyCodedOutboundInternationalPayment              ACHTransferReturnParamsReason = "incorrectly_coded_outbound_international_payment"
	ACHTransferReturnParamsReasonOther                                                     ACHTransferReturnParamsReason = "other"
)

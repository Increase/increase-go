package requests

import (
	"fmt"

	"github.com/increase/increase-go/core/fields"
	pjson "github.com/increase/increase-go/core/json"
)

type SimulateAnACHTransferToYourAccountParameters struct {
	// The identifier of the Account Number the inbound ACH Transfer is for.
	AccountNumberID fields.Field[string] `json:"account_number_id,required"`
	// The transfer amount in cents. A positive amount originates a credit transfer
	// pushing funds to the receiving account. A negative amount originates a debit
	// transfer pulling funds from the receiving account.
	Amount fields.Field[int64] `json:"amount,required"`
	// The description of the date of the transfer.
	CompanyDescriptiveDate fields.Field[string] `json:"company_descriptive_date"`
	// Data associated with the transfer set by the sender.
	CompanyDiscretionaryData fields.Field[string] `json:"company_discretionary_data"`
	// The description of the transfer set by the sender.
	CompanyEntryDescription fields.Field[string] `json:"company_entry_description"`
	// The name of the sender.
	CompanyName fields.Field[string] `json:"company_name"`
	// The sender's company id.
	CompanyID fields.Field[string] `json:"company_id"`
}

// MarshalJSON serializes SimulateAnACHTransferToYourAccountParameters into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *SimulateAnACHTransferToYourAccountParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r SimulateAnACHTransferToYourAccountParameters) String() (result string) {
	return fmt.Sprintf("&SimulateAnACHTransferToYourAccountParameters{AccountNumberID:%s Amount:%s CompanyDescriptiveDate:%s CompanyDiscretionaryData:%s CompanyEntryDescription:%s CompanyName:%s CompanyID:%s}", r.AccountNumberID, r.Amount, r.CompanyDescriptiveDate, r.CompanyDiscretionaryData, r.CompanyEntryDescription, r.CompanyName, r.CompanyID)
}

type ReturnASandboxACHTransferParameters struct {
	// The reason why the Federal Reserve or destination bank returned this transfer.
	// Defaults to `no_account`.
	Reason fields.Field[ReturnASandboxACHTransferParametersReason] `json:"reason"`
}

// MarshalJSON serializes ReturnASandboxACHTransferParameters into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *ReturnASandboxACHTransferParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ReturnASandboxACHTransferParameters) String() (result string) {
	return fmt.Sprintf("&ReturnASandboxACHTransferParameters{Reason:%s}", r.Reason)
}

type ReturnASandboxACHTransferParametersReason string

const (
	ReturnASandboxACHTransferParametersReasonInsufficientFund                                          ReturnASandboxACHTransferParametersReason = "insufficient_fund"
	ReturnASandboxACHTransferParametersReasonNoAccount                                                 ReturnASandboxACHTransferParametersReason = "no_account"
	ReturnASandboxACHTransferParametersReasonAccountClosed                                             ReturnASandboxACHTransferParametersReason = "account_closed"
	ReturnASandboxACHTransferParametersReasonInvalidAccountNumberStructure                             ReturnASandboxACHTransferParametersReason = "invalid_account_number_structure"
	ReturnASandboxACHTransferParametersReasonAccountFrozenEntryReturnedPerOfacInstruction              ReturnASandboxACHTransferParametersReason = "account_frozen_entry_returned_per_ofac_instruction"
	ReturnASandboxACHTransferParametersReasonCreditEntryRefusedByReceiver                              ReturnASandboxACHTransferParametersReason = "credit_entry_refused_by_receiver"
	ReturnASandboxACHTransferParametersReasonUnauthorizedDebitToConsumerAccountUsingCorporateSecCode   ReturnASandboxACHTransferParametersReason = "unauthorized_debit_to_consumer_account_using_corporate_sec_code"
	ReturnASandboxACHTransferParametersReasonCorporateCustomerAdvisedNotAuthorized                     ReturnASandboxACHTransferParametersReason = "corporate_customer_advised_not_authorized"
	ReturnASandboxACHTransferParametersReasonPaymentStopped                                            ReturnASandboxACHTransferParametersReason = "payment_stopped"
	ReturnASandboxACHTransferParametersReasonNonTransactionAccount                                     ReturnASandboxACHTransferParametersReason = "non_transaction_account"
	ReturnASandboxACHTransferParametersReasonUncollectedFunds                                          ReturnASandboxACHTransferParametersReason = "uncollected_funds"
	ReturnASandboxACHTransferParametersReasonRoutingNumberCheckDigitError                              ReturnASandboxACHTransferParametersReason = "routing_number_check_digit_error"
	ReturnASandboxACHTransferParametersReasonCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete ReturnASandboxACHTransferParametersReason = "customer_advised_unauthorized_improper_ineligible_or_incomplete"
	ReturnASandboxACHTransferParametersReasonAmountFieldError                                          ReturnASandboxACHTransferParametersReason = "amount_field_error"
	ReturnASandboxACHTransferParametersReasonAuthorizationRevokedByCustomer                            ReturnASandboxACHTransferParametersReason = "authorization_revoked_by_customer"
	ReturnASandboxACHTransferParametersReasonInvalidACHRoutingNumber                                   ReturnASandboxACHTransferParametersReason = "invalid_ach_routing_number"
	ReturnASandboxACHTransferParametersReasonFileRecordEditCriteria                                    ReturnASandboxACHTransferParametersReason = "file_record_edit_criteria"
	ReturnASandboxACHTransferParametersReasonEnrInvalidIndividualName                                  ReturnASandboxACHTransferParametersReason = "enr_invalid_individual_name"
	ReturnASandboxACHTransferParametersReasonReturnedPerOdfiRequest                                    ReturnASandboxACHTransferParametersReason = "returned_per_odfi_request"
	ReturnASandboxACHTransferParametersReasonAddendaError                                              ReturnASandboxACHTransferParametersReason = "addenda_error"
	ReturnASandboxACHTransferParametersReasonLimitedParticipationDfi                                   ReturnASandboxACHTransferParametersReason = "limited_participation_dfi"
	ReturnASandboxACHTransferParametersReasonIncorrectlyCodedOutboundInternationalPayment              ReturnASandboxACHTransferParametersReason = "incorrectly_coded_outbound_international_payment"
	ReturnASandboxACHTransferParametersReasonOther                                                     ReturnASandboxACHTransferParametersReason = "other"
)

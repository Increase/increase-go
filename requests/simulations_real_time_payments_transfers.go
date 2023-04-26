package requests

import (
	"github.com/increase/increase-go/core/field"
	apijson "github.com/increase/increase-go/core/json"
)

type SimulationRealTimePaymentsTransferCompleteParams struct {
	// If set, the simulation will reject the transfer.
	Rejection field.Field[SimulationRealTimePaymentsTransferCompleteParamsRejection] `json:"rejection"`
}

// MarshalJSON serializes SimulationRealTimePaymentsTransferCompleteParams into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r SimulationRealTimePaymentsTransferCompleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SimulationRealTimePaymentsTransferCompleteParamsRejection struct {
	// The reason code that the simulated rejection will have.
	RejectReasonCode field.Field[SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode] `json:"reject_reason_code,required"`
}

type SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode string

const (
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeAccountClosed                                 SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "account_closed"
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeAccountBlocked                                SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "account_blocked"
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeInvalidCreditorAccountType                    SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "invalid_creditor_account_type"
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeInvalidCreditorAccountNumber                  SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "invalid_creditor_account_number"
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeInvalidCreditorFinancialInstitutionIdentifier SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "invalid_creditor_financial_institution_identifier"
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeEndCustomerDeceased                           SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "end_customer_deceased"
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeNarrative                                     SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "narrative"
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeTransactionForbidden                          SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "transaction_forbidden"
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeTransactionTypeNotSupported                   SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "transaction_type_not_supported"
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeUnexpectedAmount                              SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "unexpected_amount"
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeAmountExceedsBankLimits                       SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "amount_exceeds_bank_limits"
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeInvalidCreditorAddress                        SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "invalid_creditor_address"
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeUnknownEndCustomer                            SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "unknown_end_customer"
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeInvalidDebtorAddress                          SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "invalid_debtor_address"
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeTimeout                                       SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "timeout"
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeUnsupportedMessageForRecipient                SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "unsupported_message_for_recipient"
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeRecipientConnectionNotAvailable               SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "recipient_connection_not_available"
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeRealTimePaymentsSuspended                     SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "real_time_payments_suspended"
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeInstructedAgentSignedOff                      SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "instructed_agent_signed_off"
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeProcessingError                               SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "processing_error"
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeOther                                         SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "other"
)

type SimulationRealTimePaymentsTransferNewInboundParams struct {
	// The identifier of the Account Number the inbound Real Time Payments Transfer is
	// for.
	AccountNumberID field.Field[string] `json:"account_number_id,required"`
	// The transfer amount in USD cents. Must be positive.
	Amount field.Field[int64] `json:"amount,required"`
	// The identifier of a pending Request for Payment that this transfer will fulfill.
	RequestForPaymentID field.Field[string] `json:"request_for_payment_id"`
	// The name provided by the sender of the transfer.
	DebtorName field.Field[string] `json:"debtor_name"`
	// The account number of the account that sent the transfer.
	DebtorAccountNumber field.Field[string] `json:"debtor_account_number"`
	// The routing number of the account that sent the transfer.
	DebtorRoutingNumber field.Field[string] `json:"debtor_routing_number"`
	// Additional information included with the transfer.
	RemittanceInformation field.Field[string] `json:"remittance_information"`
}

// MarshalJSON serializes SimulationRealTimePaymentsTransferNewInboundParams into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r SimulationRealTimePaymentsTransferNewInboundParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

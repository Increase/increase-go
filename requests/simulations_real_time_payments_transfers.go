package requests

import (
	"github.com/increase/increase-go/core/field"
	pjson "github.com/increase/increase-go/core/json"
)

type RealTimePaymentsTransferCompleteParams struct {
	// If set, the simulation will reject the transfer.
	Rejection field.Field[RealTimePaymentsTransferCompleteParamsRejection] `json:"rejection"`
}

// MarshalJSON serializes RealTimePaymentsTransferCompleteParams into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r RealTimePaymentsTransferCompleteParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type RealTimePaymentsTransferCompleteParamsRejection struct {
	// The reason code that the simulated rejection will have.
	RejectReasonCode field.Field[RealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode] `json:"reject_reason_code,required"`
}

type RealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode string

const (
	RealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeAccountClosed                                 RealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "account_closed"
	RealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeAccountBlocked                                RealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "account_blocked"
	RealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeInvalidCreditorAccountType                    RealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "invalid_creditor_account_type"
	RealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeInvalidCreditorAccountNumber                  RealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "invalid_creditor_account_number"
	RealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeInvalidCreditorFinancialInstitutionIdentifier RealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "invalid_creditor_financial_institution_identifier"
	RealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeEndCustomerDeceased                           RealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "end_customer_deceased"
	RealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeNarrative                                     RealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "narrative"
	RealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeTransactionForbidden                          RealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "transaction_forbidden"
	RealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeTransactionTypeNotSupported                   RealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "transaction_type_not_supported"
	RealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeUnexpectedAmount                              RealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "unexpected_amount"
	RealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeAmountExceedsBankLimits                       RealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "amount_exceeds_bank_limits"
	RealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeInvalidCreditorAddress                        RealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "invalid_creditor_address"
	RealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeUnknownEndCustomer                            RealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "unknown_end_customer"
	RealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeInvalidDebtorAddress                          RealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "invalid_debtor_address"
	RealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeTimeout                                       RealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "timeout"
	RealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeUnsupportedMessageForRecipient                RealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "unsupported_message_for_recipient"
	RealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeRecipientConnectionNotAvailable               RealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "recipient_connection_not_available"
	RealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeRealTimePaymentsSuspended                     RealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "real_time_payments_suspended"
	RealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeInstructedAgentSignedOff                      RealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "instructed_agent_signed_off"
	RealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeProcessingError                               RealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "processing_error"
	RealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeOther                                         RealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "other"
)

type RealTimePaymentsTransferNewInboundParams struct {
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

// MarshalJSON serializes RealTimePaymentsTransferNewInboundParams into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r RealTimePaymentsTransferNewInboundParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

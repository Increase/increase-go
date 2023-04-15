package requests

import (
	"github.com/increase/increase-go/core/field"
	pjson "github.com/increase/increase-go/core/json"
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

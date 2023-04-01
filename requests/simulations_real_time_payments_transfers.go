package requests

import (
	"fmt"

	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/fields"
)

type SimulateARealTimePaymentsTransferToYourAccountParameters struct {
	// The identifier of the Account Number the inbound Real Time Payments Transfer is
	// for.
	AccountNumberID fields.Field[string] `json:"account_number_id,required"`
	// The transfer amount in USD cents. Must be positive.
	Amount fields.Field[int64] `json:"amount,required"`
	// The identifier of a pending Request for Payment that this transfer will fulfill.
	RequestForPaymentID fields.Field[string] `json:"request_for_payment_id"`
	// The name provided by the sender of the transfer.
	DebtorName fields.Field[string] `json:"debtor_name"`
	// The account number of the account that sent the transfer.
	DebtorAccountNumber fields.Field[string] `json:"debtor_account_number"`
	// The routing number of the account that sent the transfer.
	DebtorRoutingNumber fields.Field[string] `json:"debtor_routing_number"`
	// Additional information included with the transfer.
	RemittanceInformation fields.Field[string] `json:"remittance_information"`
}

// MarshalJSON serializes SimulateARealTimePaymentsTransferToYourAccountParameters
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *SimulateARealTimePaymentsTransferToYourAccountParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r SimulateARealTimePaymentsTransferToYourAccountParameters) String() (result string) {
	return fmt.Sprintf("&SimulateARealTimePaymentsTransferToYourAccountParameters{AccountNumberID:%s Amount:%s RequestForPaymentID:%s DebtorName:%s DebtorAccountNumber:%s DebtorRoutingNumber:%s RemittanceInformation:%s}", r.AccountNumberID, r.Amount, r.RequestForPaymentID, r.DebtorName, r.DebtorAccountNumber, r.DebtorRoutingNumber, r.RemittanceInformation)
}

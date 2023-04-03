package requests

import (
	"fmt"

	"github.com/increase/increase-go/core/fields"
	pjson "github.com/increase/increase-go/core/json"
)

type SimulateARefundOnACardParameters struct {
	// The identifier for the Transaction to refund. The Transaction's source must have
	// a category of card_settlement.
	TransactionID fields.Field[string] `json:"transaction_id,required"`
}

// MarshalJSON serializes SimulateARefundOnACardParameters into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *SimulateARefundOnACardParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r SimulateARefundOnACardParameters) String() (result string) {
	return fmt.Sprintf("&SimulateARefundOnACardParameters{TransactionID:%s}", r.TransactionID)
}

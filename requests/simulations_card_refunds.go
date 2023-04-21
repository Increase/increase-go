package requests

import (
	"github.com/increase/increase-go/core/field"
	pjson "github.com/increase/increase-go/core/json"
)

type SimulationCardRefundNewParams struct {
	// The identifier for the Transaction to refund. The Transaction's source must have
	// a category of card_settlement.
	TransactionID field.Field[string] `json:"transaction_id,required"`
}

// MarshalJSON serializes SimulationCardRefundNewParams into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r SimulationCardRefundNewParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

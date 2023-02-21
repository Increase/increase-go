package types

import (
	"fmt"
	"increase/core"
	"increase/core/pjson"
)

type SimulateARefundOnACardParameters struct {
	// The identifier for the Transaction to refund. The Transaction's source must have
	// a category of card_settlement.
	TransactionID *string                `pjson:"transaction_id"`
	jsonFields    map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// SimulateARefundOnACardParameters using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *SimulateARefundOnACardParameters) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes SimulateARefundOnACardParameters into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *SimulateARefundOnACardParameters) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The identifier for the Transaction to refund. The Transaction's source must have
// a category of card_settlement.
func (r *SimulateARefundOnACardParameters) GetTransactionID() (TransactionID string) {
	if r != nil && r.TransactionID != nil {
		TransactionID = *r.TransactionID
	}
	return
}

func (r SimulateARefundOnACardParameters) String() (result string) {
	return fmt.Sprintf("&SimulateARefundOnACardParameters{TransactionID:%s}", core.FmtP(r.TransactionID))
}

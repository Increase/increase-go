package requests

import (
	"github.com/increase/increase-go/internal/field"
	apijson "github.com/increase/increase-go/internal/json"
)

type SimulationInterestPaymentNewParams struct {
	// The identifier of the Account Number the Interest Payment is for.
	AccountID field.Field[string] `json:"account_id,required"`
	// The interest amount in cents. Must be positive.
	Amount field.Field[int64] `json:"amount,required"`
}

// MarshalJSON serializes SimulationInterestPaymentNewParams into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r SimulationInterestPaymentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

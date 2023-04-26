package requests

import (
	"github.com/increase/increase-go/core/field"
	apijson "github.com/increase/increase-go/core/json"
)

type SimulationAccountStatementNewParams struct {
	// The identifier of the Account the statement is for.
	AccountID field.Field[string] `json:"account_id,required"`
}

// MarshalJSON serializes SimulationAccountStatementNewParams into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r SimulationAccountStatementNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

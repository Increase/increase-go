package requests

import (
	"fmt"

	"github.com/increase/increase-go/core/fields"
	pjson "github.com/increase/increase-go/core/json"
)

type SimulateAnAccountStatementBeingCreatedParameters struct {
	// The identifier of the Account the statement is for.
	AccountID fields.Field[string] `json:"account_id,required"`
}

// MarshalJSON serializes SimulateAnAccountStatementBeingCreatedParameters into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *SimulateAnAccountStatementBeingCreatedParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r SimulateAnAccountStatementBeingCreatedParameters) String() (result string) {
	return fmt.Sprintf("&SimulateAnAccountStatementBeingCreatedParameters{AccountID:%s}", r.AccountID)
}

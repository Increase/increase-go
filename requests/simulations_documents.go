package requests

import (
	"fmt"

	"github.com/increase/increase-go/core/fields"
	pjson "github.com/increase/increase-go/core/json"
)

type SimulateATaxDocumentBeingCreatedParameters struct {
	// The identifier of the Account the tax document is for.
	AccountID fields.Field[string] `json:"account_id,required"`
}

// MarshalJSON serializes SimulateATaxDocumentBeingCreatedParameters into an array
// of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *SimulateATaxDocumentBeingCreatedParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r SimulateATaxDocumentBeingCreatedParameters) String() (result string) {
	return fmt.Sprintf("&SimulateATaxDocumentBeingCreatedParameters{AccountID:%s}", r.AccountID)
}

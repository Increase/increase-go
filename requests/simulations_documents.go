package requests

import (
	"github.com/increase/increase-go/core/field"
	pjson "github.com/increase/increase-go/core/json"
)

type DocumentNewParams struct {
	// The identifier of the Account the tax document is for.
	AccountID field.Field[string] `json:"account_id,required"`
}

// MarshalJSON serializes DocumentNewParams into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *DocumentNewParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

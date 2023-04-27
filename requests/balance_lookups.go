package requests

import (
	"github.com/increase/increase-go/internal/field"
	apijson "github.com/increase/increase-go/internal/json"
)

type BalanceLookupLookupParams struct {
	// The Account to query the balance for.
	AccountID field.Field[string] `json:"account_id,required"`
}

// MarshalJSON serializes BalanceLookupLookupParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r BalanceLookupLookupParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

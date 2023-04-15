package requests

import (
	"github.com/increase/increase-go/core/field"
	pjson "github.com/increase/increase-go/core/json"
)

type DigitalWalletTokenRequestNewParams struct {
	// The identifier of the Card to be authorized.
	CardID field.Field[string] `json:"card_id,required"`
}

// MarshalJSON serializes DigitalWalletTokenRequestNewParams into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r DigitalWalletTokenRequestNewParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

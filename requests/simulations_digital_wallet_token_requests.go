package requests

import (
	"github.com/increase/increase-go/internal/field"
	apijson "github.com/increase/increase-go/internal/json"
)

type SimulationDigitalWalletTokenRequestNewParams struct {
	// The identifier of the Card to be authorized.
	CardID field.Field[string] `json:"card_id,required"`
}

// MarshalJSON serializes SimulationDigitalWalletTokenRequestNewParams into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r SimulationDigitalWalletTokenRequestNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

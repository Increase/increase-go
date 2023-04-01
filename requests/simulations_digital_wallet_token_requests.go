package requests

import (
	"fmt"

	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/fields"
)

type SimulateDigitalWalletProvisioningForACardParameters struct {
	// The identifier of the Card to be authorized.
	CardID fields.Field[string] `json:"card_id,required"`
}

// MarshalJSON serializes SimulateDigitalWalletProvisioningForACardParameters into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *SimulateDigitalWalletProvisioningForACardParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r SimulateDigitalWalletProvisioningForACardParameters) String() (result string) {
	return fmt.Sprintf("&SimulateDigitalWalletProvisioningForACardParameters{CardID:%s}", r.CardID)
}

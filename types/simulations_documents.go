package types

import (
	"fmt"
	"increase/core"
	"increase/core/pjson"
)

type SimulateATaxDocumentBeingCreatedParameters struct {
	// The identifier of the Account the tax document is for.
	AccountID  *string                `pjson:"account_id"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// SimulateATaxDocumentBeingCreatedParameters using the internal pjson library.
// Unrecognized fields are stored in the `Extras` property.
func (r *SimulateATaxDocumentBeingCreatedParameters) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes SimulateATaxDocumentBeingCreatedParameters into an array
// of bytes using the gjson library. Members of the `Extras` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *SimulateATaxDocumentBeingCreatedParameters) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The identifier of the Account the tax document is for.
func (r *SimulateATaxDocumentBeingCreatedParameters) GetAccountID() (AccountID string) {
	if r != nil && r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

func (r SimulateATaxDocumentBeingCreatedParameters) String() (result string) {
	return fmt.Sprintf("&SimulateATaxDocumentBeingCreatedParameters{AccountID:%s}", core.FmtP(r.AccountID))
}

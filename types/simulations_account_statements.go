package types

import (
	"fmt"
	"increase/core"
	"increase/core/pjson"
)

type SimulateAnAccountStatementBeingCreatedParameters struct {
	// The identifier of the Account the statement is for.
	AccountID  *string                `pjson:"account_id"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// SimulateAnAccountStatementBeingCreatedParameters using the internal pjson
// library. Unrecognized fields are stored in the `Extras` property.
func (r *SimulateAnAccountStatementBeingCreatedParameters) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes SimulateAnAccountStatementBeingCreatedParameters into an
// array of bytes using the gjson library. Members of the `Extras` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *SimulateAnAccountStatementBeingCreatedParameters) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The identifier of the Account the statement is for.
func (r *SimulateAnAccountStatementBeingCreatedParameters) GetAccountID() (AccountID string) {
	if r != nil && r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

func (r SimulateAnAccountStatementBeingCreatedParameters) String() (result string) {
	return fmt.Sprintf("&SimulateAnAccountStatementBeingCreatedParameters{AccountID:%s}", core.FmtP(r.AccountID))
}

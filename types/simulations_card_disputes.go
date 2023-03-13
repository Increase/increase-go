package types

import (
	"fmt"

	"github.com/increase/increase-go/core"
	"github.com/increase/increase-go/core/pjson"
)

type SimulatesAdvancingTheStateOfACardDisputeParameters struct {
	// The status to move the dispute to.
	Status *SimulatesAdvancingTheStateOfACardDisputeParametersStatus `pjson:"status"`
	// Why the dispute was rejected. Not required for accepting disputes.
	Explanation *string                `pjson:"explanation"`
	jsonFields  map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// SimulatesAdvancingTheStateOfACardDisputeParameters using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *SimulatesAdvancingTheStateOfACardDisputeParameters) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes SimulatesAdvancingTheStateOfACardDisputeParameters into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *SimulatesAdvancingTheStateOfACardDisputeParameters) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The status to move the dispute to.
func (r SimulatesAdvancingTheStateOfACardDisputeParameters) GetStatus() (Status SimulatesAdvancingTheStateOfACardDisputeParametersStatus) {
	if r.Status != nil {
		Status = *r.Status
	}
	return
}

// Why the dispute was rejected. Not required for accepting disputes.
func (r SimulatesAdvancingTheStateOfACardDisputeParameters) GetExplanation() (Explanation string) {
	if r.Explanation != nil {
		Explanation = *r.Explanation
	}
	return
}

func (r SimulatesAdvancingTheStateOfACardDisputeParameters) String() (result string) {
	return fmt.Sprintf("&SimulatesAdvancingTheStateOfACardDisputeParameters{Status:%s Explanation:%s}", core.FmtP(r.Status), core.FmtP(r.Explanation))
}

type SimulatesAdvancingTheStateOfACardDisputeParametersStatus string

const (
	SimulatesAdvancingTheStateOfACardDisputeParametersStatusAccepted SimulatesAdvancingTheStateOfACardDisputeParametersStatus = "accepted"
	SimulatesAdvancingTheStateOfACardDisputeParametersStatusRejected SimulatesAdvancingTheStateOfACardDisputeParametersStatus = "rejected"
)

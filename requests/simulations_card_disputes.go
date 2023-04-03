package requests

import (
	"fmt"

	"github.com/increase/increase-go/core/fields"
	pjson "github.com/increase/increase-go/core/json"
)

type SimulatesAdvancingTheStateOfACardDisputeParameters struct {
	// The status to move the dispute to.
	Status fields.Field[SimulatesAdvancingTheStateOfACardDisputeParametersStatus] `json:"status,required"`
	// Why the dispute was rejected. Not required for accepting disputes.
	Explanation fields.Field[string] `json:"explanation"`
}

// MarshalJSON serializes SimulatesAdvancingTheStateOfACardDisputeParameters into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *SimulatesAdvancingTheStateOfACardDisputeParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r SimulatesAdvancingTheStateOfACardDisputeParameters) String() (result string) {
	return fmt.Sprintf("&SimulatesAdvancingTheStateOfACardDisputeParameters{Status:%s Explanation:%s}", r.Status, r.Explanation)
}

type SimulatesAdvancingTheStateOfACardDisputeParametersStatus string

const (
	SimulatesAdvancingTheStateOfACardDisputeParametersStatusAccepted SimulatesAdvancingTheStateOfACardDisputeParametersStatus = "accepted"
	SimulatesAdvancingTheStateOfACardDisputeParametersStatusRejected SimulatesAdvancingTheStateOfACardDisputeParametersStatus = "rejected"
)

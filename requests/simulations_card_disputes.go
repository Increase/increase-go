package requests

import (
	"github.com/increase/increase-go/core/field"
	pjson "github.com/increase/increase-go/core/json"
)

type CardDisputeActionParams struct {
	// The status to move the dispute to.
	Status field.Field[CardDisputeActionParamsStatus] `json:"status,required"`
	// Why the dispute was rejected. Not required for accepting disputes.
	Explanation field.Field[string] `json:"explanation"`
}

// MarshalJSON serializes CardDisputeActionParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CardDisputeActionParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type CardDisputeActionParamsStatus string

const (
	CardDisputeActionParamsStatusAccepted CardDisputeActionParamsStatus = "accepted"
	CardDisputeActionParamsStatusRejected CardDisputeActionParamsStatus = "rejected"
)

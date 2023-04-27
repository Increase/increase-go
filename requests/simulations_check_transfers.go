package requests

import (
	"github.com/increase/increase-go/internal/field"
	apijson "github.com/increase/increase-go/internal/json"
)

type SimulationCheckTransferReturnParams struct {
	// The reason why the Check Transfer was returned to Increase.
	Reason field.Field[SimulationCheckTransferReturnParamsReason] `json:"reason,required"`
}

// MarshalJSON serializes SimulationCheckTransferReturnParams into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r SimulationCheckTransferReturnParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SimulationCheckTransferReturnParamsReason string

const (
	SimulationCheckTransferReturnParamsReasonMailDeliveryFailure SimulationCheckTransferReturnParamsReason = "mail_delivery_failure"
	SimulationCheckTransferReturnParamsReasonRefusedByRecipient  SimulationCheckTransferReturnParamsReason = "refused_by_recipient"
)

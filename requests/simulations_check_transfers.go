package requests

import (
	"github.com/increase/increase-go/core/field"
	pjson "github.com/increase/increase-go/core/json"
)

type CheckTransferReturnParams struct {
	// The reason why the Check Transfer was returned to Increase.
	Reason field.Field[CheckTransferReturnParamsReason] `json:"reason,required"`
}

// MarshalJSON serializes CheckTransferReturnParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r CheckTransferReturnParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type CheckTransferReturnParamsReason string

const (
	CheckTransferReturnParamsReasonMailDeliveryFailure CheckTransferReturnParamsReason = "mail_delivery_failure"
	CheckTransferReturnParamsReasonRefusedByRecipient  CheckTransferReturnParamsReason = "refused_by_recipient"
)

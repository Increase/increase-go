package requests

import (
	"fmt"

	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/fields"
)

type ReturnASandboxCheckTransferParameters struct {
	// The reason why the Check Transfer was returned to Increase.
	Reason fields.Field[ReturnASandboxCheckTransferParametersReason] `json:"reason,required"`
}

// MarshalJSON serializes ReturnASandboxCheckTransferParameters into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *ReturnASandboxCheckTransferParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ReturnASandboxCheckTransferParameters) String() (result string) {
	return fmt.Sprintf("&ReturnASandboxCheckTransferParameters{Reason:%s}", r.Reason)
}

type ReturnASandboxCheckTransferParametersReason string

const (
	ReturnASandboxCheckTransferParametersReasonMailDeliveryFailure ReturnASandboxCheckTransferParametersReason = "mail_delivery_failure"
	ReturnASandboxCheckTransferParametersReasonRefusedByRecipient  ReturnASandboxCheckTransferParametersReason = "refused_by_recipient"
)

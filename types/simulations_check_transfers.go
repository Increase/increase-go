package types

import (
	"fmt"

	"github.com/increase/increase-go/core"
	"github.com/increase/increase-go/core/pjson"
)

type ReturnASandboxCheckTransferParameters struct {
	// The reason why the Check Transfer was returned to Increase.
	Reason     *ReturnASandboxCheckTransferParametersReason `pjson:"reason"`
	jsonFields map[string]interface{}                       `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// ReturnASandboxCheckTransferParameters using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *ReturnASandboxCheckTransferParameters) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes ReturnASandboxCheckTransferParameters into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *ReturnASandboxCheckTransferParameters) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The reason why the Check Transfer was returned to Increase.
func (r ReturnASandboxCheckTransferParameters) GetReason() (Reason ReturnASandboxCheckTransferParametersReason) {
	if r.Reason != nil {
		Reason = *r.Reason
	}
	return
}

func (r ReturnASandboxCheckTransferParameters) String() (result string) {
	return fmt.Sprintf("&ReturnASandboxCheckTransferParameters{Reason:%s}", core.FmtP(r.Reason))
}

type ReturnASandboxCheckTransferParametersReason string

const (
	ReturnASandboxCheckTransferParametersReasonMailDeliveryFailure ReturnASandboxCheckTransferParametersReason = "mail_delivery_failure"
	ReturnASandboxCheckTransferParametersReasonRefusedByRecipient  ReturnASandboxCheckTransferParametersReason = "refused_by_recipient"
)

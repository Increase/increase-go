package requests

import (
	"fmt"
	"time"

	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/fields"
)

type Group struct {
	// If the Group is activated or not.
	ActivationStatus fields.Field[GroupActivationStatus] `json:"activation_status,required"`
	// If the Group is allowed to create ACH debits.
	ACHDebitStatus fields.Field[GroupACHDebitStatus] `json:"ach_debit_status,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Group
	// was created.
	CreatedAt fields.Field[time.Time] `json:"created_at,required" format:"date-time"`
	// The Group identifier.
	ID fields.Field[string] `json:"id,required"`
	// A constant representing the object's type. For this resource it will always be
	// `group`.
	Type fields.Field[GroupType] `json:"type,required"`
}

// MarshalJSON serializes Group into an array of bytes using the gjson library.
// Members of the `jsonFields` field are serialized into the top-level, and will
// overwrite known members of the same name.
func (r *Group) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r Group) String() (result string) {
	return fmt.Sprintf("&Group{ActivationStatus:%s ACHDebitStatus:%s CreatedAt:%s ID:%s Type:%s}", r.ActivationStatus, r.ACHDebitStatus, r.CreatedAt, r.ID, r.Type)
}

type GroupActivationStatus string

const (
	GroupActivationStatusUnactivated GroupActivationStatus = "unactivated"
	GroupActivationStatusActivated   GroupActivationStatus = "activated"
)

type GroupACHDebitStatus string

const (
	GroupACHDebitStatusDisabled GroupACHDebitStatus = "disabled"
	GroupACHDebitStatusEnabled  GroupACHDebitStatus = "enabled"
)

type GroupType string

const (
	GroupTypeGroup GroupType = "group"
)

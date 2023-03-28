package responses

import (
	"time"

	pjson "github.com/increase/increase-go/core/json"
)

type Group struct {
	// If the Group is activated or not.
	ActivationStatus GroupActivationStatus `json:"activation_status,required"`
	// If the Group is allowed to create ACH debits.
	ACHDebitStatus GroupACHDebitStatus `json:"ach_debit_status,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Group
	// was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The Group identifier.
	ID string `json:"id,required"`
	// A constant representing the object's type. For this resource it will always be
	// `group`.
	Type GroupType `json:"type,required"`
	JSON GroupJSON
}

type GroupJSON struct {
	ActivationStatus pjson.Metadata
	ACHDebitStatus   pjson.Metadata
	CreatedAt        pjson.Metadata
	ID               pjson.Metadata
	Type             pjson.Metadata
	Raw              []byte
	Extras           map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Group using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Group) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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

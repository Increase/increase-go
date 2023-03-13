package types

import (
	"fmt"

	"github.com/increase/increase-go/core"
	"github.com/increase/increase-go/core/pjson"
)

type Group struct {
	// If the Group is activated or not.
	ActivationStatus *GroupActivationStatus `pjson:"activation_status"`
	// If the Group is allowed to create ACH debits.
	ACHDebitStatus *GroupACHDebitStatus `pjson:"ach_debit_status"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Group
	// was created.
	CreatedAt *string `pjson:"created_at"`
	// The Group identifier.
	ID *string `pjson:"id"`
	// A constant representing the object's type. For this resource it will always be
	// `group`.
	Type       *GroupType             `pjson:"type"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into Group using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Group) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes Group into an array of bytes using the gjson library.
// Members of the `jsonFields` field are serialized into the top-level, and will
// overwrite known members of the same name.
func (r *Group) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// If the Group is activated or not.
func (r *Group) GetActivationStatus() (ActivationStatus GroupActivationStatus) {
	if r != nil && r.ActivationStatus != nil {
		ActivationStatus = *r.ActivationStatus
	}
	return
}

// If the Group is allowed to create ACH debits.
func (r *Group) GetACHDebitStatus() (ACHDebitStatus GroupACHDebitStatus) {
	if r != nil && r.ACHDebitStatus != nil {
		ACHDebitStatus = *r.ACHDebitStatus
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Group
// was created.
func (r *Group) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The Group identifier.
func (r *Group) GetID() (ID string) {
	if r != nil && r.ID != nil {
		ID = *r.ID
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `group`.
func (r *Group) GetType() (Type GroupType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

func (r Group) String() (result string) {
	return fmt.Sprintf("&Group{ActivationStatus:%s ACHDebitStatus:%s CreatedAt:%s ID:%s Type:%s}", core.FmtP(r.ActivationStatus), core.FmtP(r.ACHDebitStatus), core.FmtP(r.CreatedAt), core.FmtP(r.ID), core.FmtP(r.Type))
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

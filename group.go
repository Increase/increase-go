// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/Increase/increase-go/internal/apijson"
	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
)

// GroupService contains methods and other services that help with interacting with
// the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGroupService] method instead.
type GroupService struct {
	Options []option.RequestOption
}

// NewGroupService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewGroupService(opts ...option.RequestOption) (r *GroupService) {
	r = &GroupService{}
	r.Options = opts
	return
}

// Returns details for the currently authenticated Group.
func (r *GroupService) Get(ctx context.Context, opts ...option.RequestOption) (res *Group, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "groups/current"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Groups represent organizations using Increase. You can retrieve information
// about your own organization via the API. More commonly, OAuth platforms can
// retrieve information about the organizations that have granted them access.
// Learn more about OAuth [here](https://increase.com/documentation/oauth).
type Group struct {
	// The Group identifier.
	ID string `json:"id" api:"required"`
	// If the Group is allowed to create ACH debits.
	ACHDebitStatus GroupACHDebitStatus `json:"ach_debit_status" api:"required"`
	// If the Group is activated or not.
	ActivationStatus GroupActivationStatus `json:"activation_status" api:"required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Group
	// was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// A constant representing the object's type. For this resource it will always be
	// `group`.
	Type GroupType `json:"type" api:"required"`
	JSON groupJSON `json:"-"`
}

// groupJSON contains the JSON metadata for the struct [Group]
type groupJSON struct {
	ID               apijson.Field
	ACHDebitStatus   apijson.Field
	ActivationStatus apijson.Field
	CreatedAt        apijson.Field
	Type             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *Group) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r groupJSON) RawJSON() string {
	return r.raw
}

// If the Group is allowed to create ACH debits.
type GroupACHDebitStatus string

const (
	GroupACHDebitStatusDisabled GroupACHDebitStatus = "disabled"
	GroupACHDebitStatusEnabled  GroupACHDebitStatus = "enabled"
)

func (r GroupACHDebitStatus) IsKnown() bool {
	switch r {
	case GroupACHDebitStatusDisabled, GroupACHDebitStatusEnabled:
		return true
	}
	return false
}

// If the Group is activated or not.
type GroupActivationStatus string

const (
	GroupActivationStatusUnactivated GroupActivationStatus = "unactivated"
	GroupActivationStatusActivated   GroupActivationStatus = "activated"
)

func (r GroupActivationStatus) IsKnown() bool {
	switch r {
	case GroupActivationStatusUnactivated, GroupActivationStatusActivated:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `group`.
type GroupType string

const (
	GroupTypeGroup GroupType = "group"
)

func (r GroupType) IsKnown() bool {
	switch r {
	case GroupTypeGroup:
		return true
	}
	return false
}

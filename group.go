package increase

import (
	"context"
	"net/http"
	"time"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/option"
)

// GroupService contains methods and other services that help with interacting with
// the increase API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewGroupService] method instead.
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
func (r *GroupService) GetDetails(ctx context.Context, opts ...option.RequestOption) (res *Group, err error) {
	opts = append(r.Options[:], opts...)
	path := "groups/current"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Groups represent organizations using Increase. You can retrieve information
// about your own organization via the API, or (more commonly) OAuth platforms can
// retrieve information about the organizations that have granted them access.
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
	JSON groupJSON
}

// groupJSON contains the JSON metadata for the struct [Group]
type groupJSON struct {
	ActivationStatus apijson.Field
	ACHDebitStatus   apijson.Field
	CreatedAt        apijson.Field
	ID               apijson.Field
	Type             apijson.Field
	raw              string
	Extras           map[string]apijson.Field
}

func (r *Group) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

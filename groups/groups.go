package groups

import "context"
import "increase/core"

type GroupService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewGroupService(requester core.Requester) (r *GroupService) {
	r = &GroupService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

//
type Group struct {
	// If the Group is activated or not.
	ActivationStatus GroupActivationStatus `json:"activation_status"`
	// If the Group is allowed to create ACH debits.
	ACHDebitStatus GroupACHDebitStatus `json:"ach_debit_status"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Group
	// was created.
	CreatedAt string `json:"created_at"`
	// The Group identifier.
	ID string `json:"id"`
	// A constant representing the object's type. For this resource it will always be
	// `group`.
	Type GroupType `json:"type"`
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

// Returns details for the currently authenticated Group.
func (r *GroupService) RetrieveDetails(ctx context.Context, opts ...*core.RequestOpts) (res *Group, err error) {
	err = r.get(
		ctx,
		"/groups/current",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)

	return
}

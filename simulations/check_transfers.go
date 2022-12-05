package simulations

import "increase/core"
import "increase/check_transfers"
import "fmt"

type CheckTransfers struct {
	Requester core.Requester
	get       func(string, *core.CoreRequest, interface{}) error
	post      func(string, *core.CoreRequest, interface{}) error
	patch     func(string, *core.CoreRequest, interface{}) error
	put       func(string, *core.CoreRequest, interface{}) error
	delete    func(string, *core.CoreRequest, interface{}) error
}

func NewCheckTransfers(requster core.Requester) (r *CheckTransfers) {
	r = &CheckTransfers{}
	r.Requester = requster
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

// Simulates the mailing of a Check Transfer. This transfer must first have a
// `status` of `pending_approval` or `pending_submission`.
func (r *CheckTransfers) Mail(check_transfer_id string, opts ...*core.RequestOpts) (res check_transfers.CheckTransfer, err error) {
	err = r.post(
		fmt.Sprintf("/simulations/check_transfers/%s/mail", check_transfer_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

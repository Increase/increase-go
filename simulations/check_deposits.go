package simulations

import "increase/core"
import "fmt"

type CheckDeposits struct {
	Requester core.Requester
	get       func(string, *core.CoreRequest, interface{}) error
	post      func(string, *core.CoreRequest, interface{}) error
	patch     func(string, *core.CoreRequest, interface{}) error
	put       func(string, *core.CoreRequest, interface{}) error
	delete    func(string, *core.CoreRequest, interface{}) error
}

func NewCheckDeposits(requster core.Requester) (r *CheckDeposits) {
	r = &CheckDeposits{}
	r.Requester = requster
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

// Simulates the rejection of a Check Deposit by Increase due to factors like poor
// image quality. This Check Deposit must first have a `status` of `pending`.
func (r *CheckDeposits) Reject(check_deposit_id string, opts ...*core.RequestOpts) (res CheckDeposit, err error) {
	err = r.post(
		fmt.Sprintf("/simulations/check_deposits/%s/reject", check_deposit_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

// Simulates the submission of a Check Deposit to the Federal Reserve. This Check
// Deposit must first have a `status` of `pending`.
func (r *CheckDeposits) Submit(check_deposit_id string, opts ...*core.RequestOpts) (res CheckDeposit, err error) {
	err = r.post(
		fmt.Sprintf("/simulations/check_deposits/%s/submit", check_deposit_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

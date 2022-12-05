package check_deposits

import "increase/core"
import "increase/simulations"
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

func (r *CheckDeposits) Create(body CheckDepositsCreateParams, opts ...*core.RequestOpts) (res simulations.CheckDeposit, err error) {
	err = r.post(
		"/check_deposits",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *CheckDeposits) Retrieve(check_deposit_id string, opts ...*core.RequestOpts) (res simulations.CheckDeposit, err error) {
	err = r.get(
		fmt.Sprintf("/check_deposits/%s", check_deposit_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *CheckDeposits) List(query *CheckDepositsListQuery, opts ...*core.RequestOpts) (res CheckDepositsListResponse, err error) {
	err = r.get(
		"/check_deposits",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Query:  query,
		},
		&res,
	)
	return
}

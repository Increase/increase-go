package check_transfers

import "increase/core"
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

func (r *CheckTransfers) Create(body CheckTransfersCreateParams, opts ...*core.RequestOpts) (res CheckTransfer, err error) {
	err = r.post(
		"/check_transfers",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *CheckTransfers) Retrieve(check_transfer_id string, opts ...*core.RequestOpts) (res CheckTransfer, err error) {
	err = r.get(
		fmt.Sprintf("/check_transfers/%s", check_transfer_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *CheckTransfers) List(query *CheckTransfersListQuery, opts ...*core.RequestOpts) (res CheckTransfersListResponse, err error) {
	err = r.get(
		"/check_transfers",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Query:  query,
		},
		&res,
	)
	return
}

func (r *CheckTransfers) StopPayment(check_transfer_id string, opts ...*core.RequestOpts) (res CheckTransfer, err error) {
	err = r.post(
		fmt.Sprintf("/check_transfers/%s/stop_payment", check_transfer_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

package ach_transfers

import "increase/core"
import "fmt"

type ACHTransfers struct {
	Requester core.Requester
	get       func(string, *core.CoreRequest, interface{}) error
	post      func(string, *core.CoreRequest, interface{}) error
	patch     func(string, *core.CoreRequest, interface{}) error
	put       func(string, *core.CoreRequest, interface{}) error
	delete    func(string, *core.CoreRequest, interface{}) error
}

func NewACHTransfers(requster core.Requester) (r *ACHTransfers) {
	r = &ACHTransfers{}
	r.Requester = requster
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

func (r *ACHTransfers) Create(body ACHTransfersCreateParams, opts ...*core.RequestOpts) (res ACHTransfer, err error) {
	err = r.post(
		"/ach_transfers",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *ACHTransfers) Retrieve(ach_transfer_id string, opts ...*core.RequestOpts) (res ACHTransfer, err error) {
	err = r.get(
		fmt.Sprintf("/ach_transfers/%s", ach_transfer_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *ACHTransfers) List(query *ACHTransfersListQuery, opts ...*core.RequestOpts) (res ACHTransfersListResponse, err error) {
	err = r.get(
		"/ach_transfers",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Query:  query,
		},
		&res,
	)
	return
}

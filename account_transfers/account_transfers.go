package account_transfers

import "increase/core"
import "fmt"

type AccountTransfers struct {
	Requester core.Requester
	get       func(string, *core.CoreRequest, interface{}) error
	post      func(string, *core.CoreRequest, interface{}) error
	patch     func(string, *core.CoreRequest, interface{}) error
	put       func(string, *core.CoreRequest, interface{}) error
	delete    func(string, *core.CoreRequest, interface{}) error
}

func NewAccountTransfers(requster core.Requester) (r *AccountTransfers) {
	r = &AccountTransfers{}
	r.Requester = requster
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

func (r *AccountTransfers) Create(body AccountTransfersCreateParams, opts ...*core.RequestOpts) (res AccountTransfer, err error) {
	err = r.post(
		"/account_transfers",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *AccountTransfers) Retrieve(account_transfer_id string, opts ...*core.RequestOpts) (res AccountTransfer, err error) {
	err = r.get(
		fmt.Sprintf("/account_transfers/%s", account_transfer_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *AccountTransfers) List(query *AccountTransfersListQuery, opts ...*core.RequestOpts) (res AccountTransfersListResponse, err error) {
	err = r.get(
		"/account_transfers",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Query:  query,
		},
		&res,
	)
	return
}

package simulations

import "increase/core"
import "increase/account_transfers"
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

// Simulates the completion of an Account Transfer. This transfer must first have a
// `status` of `pending_approval`.
func (r *AccountTransfers) Complete(account_transfer_id string, opts ...*core.RequestOpts) (res account_transfers.AccountTransfer, err error) {
	err = r.post(
		fmt.Sprintf("/simulations/account_transfers/%s/complete", account_transfer_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

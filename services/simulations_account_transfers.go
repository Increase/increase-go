package services

import "context"
import "increase/core"
import "increase/types"
import "fmt"

type SimulationsAccountTransferService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewSimulationsAccountTransferService(requester core.Requester) (r *SimulationsAccountTransferService) {
	r = &SimulationsAccountTransferService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

// Simulates the completion of an Account Transfer. This transfer must first have a
// `status` of `pending_approval`.
func (r *SimulationsAccountTransferService) Complete(ctx context.Context, account_transfer_id string, opts ...*core.RequestOpts) (res *types.AccountTransfer, err error) {
	err = r.post(
		ctx,
		fmt.Sprintf("/simulations/account_transfers/%s/complete", account_transfer_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)

	return
}

package simulations

import "context"
import "increase/core"
import "increase/account_transfers"
import "fmt"

type AccountTransferService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewAccountTransferService(requester core.Requester) (r *AccountTransferService) {
	r = &AccountTransferService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

type PreloadedAccountTransferService struct {
	AccountTransfers *AccountTransferService
}

func (r *PreloadedAccountTransferService) Init(service *AccountTransferService) {
	r.AccountTransfers = service
}

func NewPreloadedAccountTransferService(service *AccountTransferService) (r *PreloadedAccountTransferService) {
	r = &PreloadedAccountTransferService{}
	r.Init(service)
	return
}

// Simulates the completion of an Account Transfer. This transfer must first have a
// `status` of `pending_approval`.
func (r *AccountTransferService) Complete(ctx context.Context, account_transfer_id string, opts ...*core.RequestOpts) (res *account_transfers.AccountTransfer, err error) {
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

// Simulates the completion of an Account Transfer. This transfer must first have a
// `status` of `pending_approval`.
func (r *PreloadedAccountTransferService) Complete(ctx context.Context, account_transfer_id string, opts ...*core.RequestOpts) (res *account_transfers.AccountTransfer, err error) {
	err = r.AccountTransfers.post(
		ctx,
		fmt.Sprintf("/simulations/account_transfers/%s/complete", account_transfer_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

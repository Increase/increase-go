package services

import "context"
import "increase/core"
import "increase/types"
import "fmt"

type SimulationsACHTransferService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewSimulationsACHTransferService(requester core.Requester) (r *SimulationsACHTransferService) {
	r = &SimulationsACHTransferService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

// Simulates an inbound ACH transfer to your account. The transfer may be either a
// credit or a debit depending on if the `amount` is positive or negative. This
// will result in either a Transaction or a Declined Transaction depending on if
// the transfer is allowed.
func (r *SimulationsACHTransferService) CreateInbound(ctx context.Context, body *types.SimulateAnACHTransferToYourAccountParameters, opts ...*core.RequestOpts) (res *types.ACHTransferSimulation, err error) {
	err = r.post(
		ctx,
		"/simulations/inbound_ach_transfers",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)

	return
}

// Simulates the return of an ACH Transfer by the Federal Reserve due to error
// conditions. This will also create a Transaction to account for the returned
// funds. This transfer must first have a `status` of `submitted`.
func (r *SimulationsACHTransferService) Return(ctx context.Context, ach_transfer_id string, opts ...*core.RequestOpts) (res *types.ACHTransfer, err error) {
	err = r.post(
		ctx,
		fmt.Sprintf("/simulations/ach_transfers/%s/return", ach_transfer_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)

	return
}

// Simulates the submission of an ACH Transfer to the Federal Reserve. This
// transfer must first have a `status` of `pending_approval` or
// `pending_submission`.
func (r *SimulationsACHTransferService) Submit(ctx context.Context, ach_transfer_id string, opts ...*core.RequestOpts) (res *types.ACHTransfer, err error) {
	err = r.post(
		ctx,
		fmt.Sprintf("/simulations/ach_transfers/%s/submit", ach_transfer_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)

	return
}

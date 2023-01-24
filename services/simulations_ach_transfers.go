package services

import (
	"context"
	"fmt"
	"increase/core"
	"increase/types"
)

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

// Simulates an inbound ACH transfer to your account. This imitates initiating a
// transaction to an Increase account from a different financial institution. The
// transfer may be either a credit or a debit depending on if the `amount` is
// positive or negative. The result of calling this API will be either a
// [Transaction](#transactions) or a [Declined Transaction](#declined-transactions)
// depending on whether or not the transfer is allowed.
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

// Simulates the return of an [ACH Transfer](#ach-transfers) by the Federal Reserve
// due to an error condition. This will also create a Transaction to account for
// the returned funds. This transfer must first have a `status` of `submitted`.
func (r *SimulationsACHTransferService) Return(ctx context.Context, ach_transfer_id string, body *types.ReturnASandboxACHTransferParameters, opts ...*core.RequestOpts) (res *types.ACHTransfer, err error) {
	err = r.post(
		ctx,
		fmt.Sprintf("/simulations/ach_transfers/%s/return", ach_transfer_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)

	return
}

// Simulates the submission of an [ACH Transfer](#ach-transfers) to the Federal
// Reserve. This transfer must first have a `status` of `pending_approval` or
// `pending_submission`. In production, Increase submits ACH Transfers to the
// Federal Reserve three times per day on weekdays. Since sandbox ACH Transfers are
// not submitted to the Federal Reserve, this endpoint allows you to skip that
// delay and transition the ACH Transfer to a status of `submitted`.
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

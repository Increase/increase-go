package services

import (
	"context"
	"fmt"
	"increase/options"
	"increase/types"
)

type SimulationsACHTransferService struct {
	Options []options.RequestOption
}

func NewSimulationsACHTransferService(opts ...options.RequestOption) (r *SimulationsACHTransferService) {
	r = &SimulationsACHTransferService{}
	r.Options = opts
	return
}

// Simulates an inbound ACH transfer to your account. This imitates initiating a
// transaction to an Increase account from a different financial institution. The
// transfer may be either a credit or a debit depending on if the `amount` is
// positive or negative. The result of calling this API will be either a
// [Transaction](#transactions) or a [Declined Transaction](#declined-transactions)
// depending on whether or not the transfer is allowed.
func (r *SimulationsACHTransferService) NewInbound(ctx context.Context, body *types.SimulateAnACHTransferToYourAccountParameters, opts ...options.RequestOption) (res *types.ACHTransferSimulation, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/inbound_ach_transfers"
	err = options.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Simulates the return of an [ACH Transfer](#ach-transfers) by the Federal Reserve
// due to an error condition. This will also create a Transaction to account for
// the returned funds. This transfer must first have a `status` of `submitted`.
func (r *SimulationsACHTransferService) Return(ctx context.Context, ach_transfer_id string, body *types.ReturnASandboxACHTransferParameters, opts ...options.RequestOption) (res *types.ACHTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("simulations/ach_transfers/%s/return", ach_transfer_id)
	err = options.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Simulates the submission of an [ACH Transfer](#ach-transfers) to the Federal
// Reserve. This transfer must first have a `status` of `pending_approval` or
// `pending_submission`. In production, Increase submits ACH Transfers to the
// Federal Reserve three times per day on weekdays. Since sandbox ACH Transfers are
// not submitted to the Federal Reserve, this endpoint allows you to skip that
// delay and transition the ACH Transfer to a status of `submitted`.
func (r *SimulationsACHTransferService) Submit(ctx context.Context, ach_transfer_id string, opts ...options.RequestOption) (res *types.ACHTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("simulations/ach_transfers/%s/submit", ach_transfer_id)
	err = options.ExecuteNewRequest(ctx, "POST", path, nil, &res, opts...)
	return
}

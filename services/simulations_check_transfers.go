package services

import (
	"context"
	"fmt"
	"increase/core"
	"increase/types"
)

type SimulationsCheckTransferService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewSimulationsCheckTransferService(requester core.Requester) (r *SimulationsCheckTransferService) {
	r = &SimulationsCheckTransferService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

// Simulates a [Check Transfer](#check-transfers) being deposited at a bank. This
// transfer must first have a `status` of `mailed`.
func (r *SimulationsCheckTransferService) Deposit(ctx context.Context, check_transfer_id string, opts ...*core.RequestOpts) (res *types.CheckTransfer, err error) {
	path := fmt.Sprintf("/simulations/check_transfers/%s/deposit", check_transfer_id)
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
	}
	err = r.post(ctx, path, req, &res)

	return
}

// Simulates the mailing of a [Check Transfer](#check-transfers), which happens
// once per weekday in production but can be sped up in sandbox. This transfer must
// first have a `status` of `pending_approval` or `pending_submission`.
func (r *SimulationsCheckTransferService) Mail(ctx context.Context, check_transfer_id string, opts ...*core.RequestOpts) (res *types.CheckTransfer, err error) {
	path := fmt.Sprintf("/simulations/check_transfers/%s/mail", check_transfer_id)
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
	}
	err = r.post(ctx, path, req, &res)

	return
}

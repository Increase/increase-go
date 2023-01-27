package services

import (
	"context"
	"fmt"
	"increase/core"
	"increase/types"
)

type SimulationsCheckDepositService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewSimulationsCheckDepositService(requester core.Requester) (r *SimulationsCheckDepositService) {
	r = &SimulationsCheckDepositService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

// Simulates the rejection of a [Check Deposit](#check-deposits) by Increase due to
// factors like poor image quality. This Check Deposit must first have a `status`
// of `pending`.
func (r *SimulationsCheckDepositService) Reject(ctx context.Context, check_deposit_id string, opts ...*core.RequestOpts) (res *types.CheckDeposit, err error) {
	path := fmt.Sprintf("/simulations/check_deposits/%s/reject", check_deposit_id)
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
	}
	err = r.post(ctx, path, req, &res)

	return
}

// Simulates the submission of a [Check Deposit](#check-deposits) to the Federal
// Reserve. This Check Deposit must first have a `status` of `pending`.
func (r *SimulationsCheckDepositService) Submit(ctx context.Context, check_deposit_id string, opts ...*core.RequestOpts) (res *types.CheckDeposit, err error) {
	path := fmt.Sprintf("/simulations/check_deposits/%s/submit", check_deposit_id)
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
	}
	err = r.post(ctx, path, req, &res)

	return
}

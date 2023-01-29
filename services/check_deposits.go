package services

import (
	"context"
	"fmt"
	"increase/core"
	"increase/pagination"
	"increase/types"
)

type CheckDepositService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewCheckDepositService(requester core.Requester) (r *CheckDepositService) {
	r = &CheckDepositService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

// Create a Check Deposit
func (r *CheckDepositService) New(ctx context.Context, body *types.CreateACheckDepositParameters, opts ...*core.RequestOpts) (res *types.CheckDeposit, err error) {
	path := "/check_deposits"
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
		Body:   body,
	}
	err = r.post(ctx, path, req, &res)

	return
}

// Retrieve a Check Deposit
func (r *CheckDepositService) Get(ctx context.Context, check_deposit_id string, opts ...*core.RequestOpts) (res *types.CheckDeposit, err error) {
	path := fmt.Sprintf("/check_deposits/%s", check_deposit_id)
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
	}
	err = r.get(ctx, path, req, &res)

	return
}

// List Check Deposits
func (r *CheckDepositService) List(ctx context.Context, query *types.CheckDepositListParams, opts ...*core.RequestOpts) (res *types.CheckDepositsPage, err error) {
	page := &types.CheckDepositsPage{
		Page: &pagination.Page[types.CheckDeposit]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/check_deposits",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

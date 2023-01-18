package services

import "context"
import "increase/core"
import "increase/types"
import "fmt"
import "increase/pagination"

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

func (r *CheckDepositService) Create(ctx context.Context, body *types.CreateACheckDepositParameters, opts ...*core.RequestOpts) (res *types.CheckDeposit, err error) {
	err = r.post(
		ctx,
		"/check_deposits",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)

	return
}

func (r *CheckDepositService) Retrieve(ctx context.Context, check_deposit_id string, opts ...*core.RequestOpts) (res *types.CheckDeposit, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/check_deposits/%s", check_deposit_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)

	return
}

func (r *CheckDepositService) List(ctx context.Context, query *types.ListCheckDepositsQuery, opts ...*core.RequestOpts) (res *types.CheckDepositsPage, err error) {
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

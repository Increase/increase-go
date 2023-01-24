package services

import (
	"context"
	"fmt"
	"increase/core"
	"increase/pagination"
	"increase/types"
)

type LimitService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewLimitService(requester core.Requester) (r *LimitService) {
	r = &LimitService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

func (r *LimitService) Create(ctx context.Context, body *types.CreateALimitParameters, opts ...*core.RequestOpts) (res *types.Limit, err error) {
	err = r.post(
		ctx,
		"/limits",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)

	return
}

func (r *LimitService) Retrieve(ctx context.Context, limit_id string, opts ...*core.RequestOpts) (res *types.Limit, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/limits/%s", limit_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)

	return
}

func (r *LimitService) Update(ctx context.Context, limit_id string, body *types.UpdateALimitParameters, opts ...*core.RequestOpts) (res *types.Limit, err error) {
	err = r.patch(
		ctx,
		fmt.Sprintf("/limits/%s", limit_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)

	return
}

func (r *LimitService) List(ctx context.Context, query *types.ListLimitsQuery, opts ...*core.RequestOpts) (res *types.LimitsPage, err error) {
	page := &types.LimitsPage{
		Page: &pagination.Page[types.Limit]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/limits",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

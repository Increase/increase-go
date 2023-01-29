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

// Create a Limit
func (r *LimitService) New(ctx context.Context, body *types.CreateALimitParameters, opts ...*core.RequestOpts) (res *types.Limit, err error) {
	path := "/limits"
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
		Body:   body,
	}
	err = r.post(ctx, path, req, &res)

	return
}

// Retrieve a Limit
func (r *LimitService) Get(ctx context.Context, limit_id string, opts ...*core.RequestOpts) (res *types.Limit, err error) {
	path := fmt.Sprintf("/limits/%s", limit_id)
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
	}
	err = r.get(ctx, path, req, &res)

	return
}

// Update a Limit
func (r *LimitService) Update(ctx context.Context, limit_id string, body *types.UpdateALimitParameters, opts ...*core.RequestOpts) (res *types.Limit, err error) {
	path := fmt.Sprintf("/limits/%s", limit_id)
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
		Body:   body,
	}
	err = r.patch(ctx, path, req, &res)

	return
}

// List Limits
func (r *LimitService) List(ctx context.Context, query *types.LimitListParams, opts ...*core.RequestOpts) (res *types.LimitsPage, err error) {
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

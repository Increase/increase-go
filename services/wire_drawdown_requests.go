package services

import (
	"context"
	"fmt"
	"increase/core"
	"increase/pagination"
	"increase/types"
)

type WireDrawdownRequestService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewWireDrawdownRequestService(requester core.Requester) (r *WireDrawdownRequestService) {
	r = &WireDrawdownRequestService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

// Create a Wire Drawdown Request
func (r *WireDrawdownRequestService) Create(ctx context.Context, body *types.CreateAWireDrawdownRequestParameters, opts ...*core.RequestOpts) (res *types.WireDrawdownRequest, err error) {
	err = r.post(
		ctx,
		"/wire_drawdown_requests",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)

	return
}

// Retrieve a Wire Drawdown Request
func (r *WireDrawdownRequestService) Retrieve(ctx context.Context, wire_drawdown_request_id string, opts ...*core.RequestOpts) (res *types.WireDrawdownRequest, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/wire_drawdown_requests/%s", wire_drawdown_request_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)

	return
}

// List Wire Drawdown Requests
func (r *WireDrawdownRequestService) List(ctx context.Context, query *types.ListWireDrawdownRequestsQuery, opts ...*core.RequestOpts) (res *types.WireDrawdownRequestsPage, err error) {
	page := &types.WireDrawdownRequestsPage{
		Page: &pagination.Page[types.WireDrawdownRequest]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/wire_drawdown_requests",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

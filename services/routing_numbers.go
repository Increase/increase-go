package services

import (
	"context"
	"increase/core"
	"increase/pagination"
	"increase/types"
)

type RoutingNumberService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewRoutingNumberService(requester core.Requester) (r *RoutingNumberService) {
	r = &RoutingNumberService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

// You can use this API to confirm if a routing number is valid, such as when a
// user is providing you with bank account details. Since routing numbers uniquely
// identify a bank, this will always return 0 or 1 entry. In Sandbox, the only
// valid routing number for this method is 110000000.
func (r *RoutingNumberService) List(ctx context.Context, query *types.ListRoutingNumbersQuery, opts ...*core.RequestOpts) (res *types.RoutingNumbersPage, err error) {
	page := &types.RoutingNumbersPage{
		Page: &pagination.Page[types.RoutingNumber]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/routing_numbers",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

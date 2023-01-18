package services

import "context"
import "increase/core"
import "increase/types"
import "fmt"
import "increase/pagination"

type CardDisputeService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewCardDisputeService(requester core.Requester) (r *CardDisputeService) {
	r = &CardDisputeService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

func (r *CardDisputeService) Create(ctx context.Context, body *types.CreateACardDisputeParameters, opts ...*core.RequestOpts) (res *types.CardDispute, err error) {
	err = r.post(
		ctx,
		"/card_disputes",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)

	return
}

func (r *CardDisputeService) Retrieve(ctx context.Context, card_dispute_id string, opts ...*core.RequestOpts) (res *types.CardDispute, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/card_disputes/%s", card_dispute_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)

	return
}

func (r *CardDisputeService) List(ctx context.Context, query *types.ListCardDisputesQuery, opts ...*core.RequestOpts) (res *types.CardDisputesPage, err error) {
	page := &types.CardDisputesPage{
		Page: &pagination.Page[types.CardDispute]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/card_disputes",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

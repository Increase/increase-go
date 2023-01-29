package services

import (
	"context"
	"fmt"
	"increase/core"
	"increase/pagination"
	"increase/types"
)

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

// Create a Card Dispute
func (r *CardDisputeService) New(ctx context.Context, body *types.CreateACardDisputeParameters, opts ...*core.RequestOpts) (res *types.CardDispute, err error) {
	path := "/card_disputes"
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
		Body:   body,
	}
	err = r.post(ctx, path, req, &res)

	return
}

// Retrieve a Card Dispute
func (r *CardDisputeService) Get(ctx context.Context, card_dispute_id string, opts ...*core.RequestOpts) (res *types.CardDispute, err error) {
	path := fmt.Sprintf("/card_disputes/%s", card_dispute_id)
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
	}
	err = r.get(ctx, path, req, &res)

	return
}

// List Card Disputes
func (r *CardDisputeService) List(ctx context.Context, query *types.CardDisputeListParams, opts ...*core.RequestOpts) (res *types.CardDisputesPage, err error) {
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

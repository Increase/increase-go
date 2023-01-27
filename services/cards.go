package services

import (
	"context"
	"fmt"
	"increase/core"
	"increase/pagination"
	"increase/types"
)

type CardService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewCardService(requester core.Requester) (r *CardService) {
	r = &CardService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

// Create a Card
func (r *CardService) New(ctx context.Context, body *types.CreateACardParameters, opts ...*core.RequestOpts) (res *types.Card, err error) {
	path := "/cards"
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
		Body:   body,
	}
	err = r.post(ctx, path, req, &res)

	return
}

// Retrieve a Card
func (r *CardService) Get(ctx context.Context, card_id string, opts ...*core.RequestOpts) (res *types.Card, err error) {
	path := fmt.Sprintf("/cards/%s", card_id)
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
	}
	err = r.get(ctx, path, req, &res)

	return
}

// Update a Card
func (r *CardService) Update(ctx context.Context, card_id string, body *types.UpdateACardParameters, opts ...*core.RequestOpts) (res *types.Card, err error) {
	path := fmt.Sprintf("/cards/%s", card_id)
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
		Body:   body,
	}
	err = r.patch(ctx, path, req, &res)

	return
}

// List Cards
func (r *CardService) List(ctx context.Context, query *types.ListCardsQuery, opts ...*core.RequestOpts) (res *types.CardsPage, err error) {
	page := &types.CardsPage{
		Page: &pagination.Page[types.Card]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/cards",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

// Retrieve sensitive details for a Card
func (r *CardService) GetSensitiveDetails(ctx context.Context, card_id string, opts ...*core.RequestOpts) (res *types.CardDetails, err error) {
	path := fmt.Sprintf("/cards/%s/details", card_id)
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
	}
	err = r.get(ctx, path, req, &res)

	return
}

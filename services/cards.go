package services

import "context"
import "increase/core"
import "increase/types"
import "fmt"
import "increase/pagination"

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

func (r *CardService) Create(ctx context.Context, body *types.CreateACardParameters, opts ...*core.RequestOpts) (res *types.Card, err error) {
	err = r.post(
		ctx,
		"/cards",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)

	return
}

func (r *CardService) Retrieve(ctx context.Context, card_id string, opts ...*core.RequestOpts) (res *types.Card, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/cards/%s", card_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)

	return
}

func (r *CardService) Update(ctx context.Context, card_id string, body *types.UpdateACardParameters, opts ...*core.RequestOpts) (res *types.Card, err error) {
	err = r.patch(
		ctx,
		fmt.Sprintf("/cards/%s", card_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)

	return
}

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

func (r *CardService) RetrieveSensitiveDetails(ctx context.Context, card_id string, opts ...*core.RequestOpts) (res *types.CardDetails, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/cards/%s/details", card_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)

	return
}

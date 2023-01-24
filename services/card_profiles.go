package services

import (
	"context"
	"fmt"
	"increase/core"
	"increase/pagination"
	"increase/types"
)

type CardProfileService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewCardProfileService(requester core.Requester) (r *CardProfileService) {
	r = &CardProfileService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

func (r *CardProfileService) Create(ctx context.Context, body *types.CreateACardProfileParameters, opts ...*core.RequestOpts) (res *types.CardProfile, err error) {
	err = r.post(
		ctx,
		"/card_profiles",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)

	return
}

func (r *CardProfileService) Retrieve(ctx context.Context, card_profile_id string, opts ...*core.RequestOpts) (res *types.CardProfile, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/card_profiles/%s", card_profile_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)

	return
}

func (r *CardProfileService) List(ctx context.Context, query *types.ListCardProfilesQuery, opts ...*core.RequestOpts) (res *types.CardProfilesPage, err error) {
	page := &types.CardProfilesPage{
		Page: &pagination.Page[types.CardProfile]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/card_profiles",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

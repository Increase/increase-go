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

// Create a Card Profile
func (r *CardProfileService) New(ctx context.Context, body *types.CreateACardProfileParameters, opts ...*core.RequestOpts) (res *types.CardProfile, err error) {
	path := "/card_profiles"
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
		Body:   body,
	}
	err = r.post(ctx, path, req, &res)

	return
}

// Retrieve a Card Profile
func (r *CardProfileService) Get(ctx context.Context, card_profile_id string, opts ...*core.RequestOpts) (res *types.CardProfile, err error) {
	path := fmt.Sprintf("/card_profiles/%s", card_profile_id)
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
	}
	err = r.get(ctx, path, req, &res)

	return
}

// List Card Profiles
func (r *CardProfileService) List(ctx context.Context, query *types.CardProfileListParams, opts ...*core.RequestOpts) (res *types.CardProfilesPage, err error) {
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

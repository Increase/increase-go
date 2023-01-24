package services

import (
	"context"
	"fmt"
	"increase/core"
	"increase/pagination"
	"increase/types"
)

type ExternalAccountService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewExternalAccountService(requester core.Requester) (r *ExternalAccountService) {
	r = &ExternalAccountService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

func (r *ExternalAccountService) Create(ctx context.Context, body *types.CreateAnExternalAccountParameters, opts ...*core.RequestOpts) (res *types.ExternalAccount, err error) {
	err = r.post(
		ctx,
		"/external_accounts",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)

	return
}

func (r *ExternalAccountService) Retrieve(ctx context.Context, external_account_id string, opts ...*core.RequestOpts) (res *types.ExternalAccount, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/external_accounts/%s", external_account_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)

	return
}

func (r *ExternalAccountService) Update(ctx context.Context, external_account_id string, body *types.UpdateAnExternalAccountParameters, opts ...*core.RequestOpts) (res *types.ExternalAccount, err error) {
	err = r.patch(
		ctx,
		fmt.Sprintf("/external_accounts/%s", external_account_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)

	return
}

func (r *ExternalAccountService) List(ctx context.Context, query *types.ListExternalAccountsQuery, opts ...*core.RequestOpts) (res *types.ExternalAccountsPage, err error) {
	page := &types.ExternalAccountsPage{
		Page: &pagination.Page[types.ExternalAccount]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/external_accounts",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

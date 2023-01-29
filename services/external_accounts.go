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

// Create an External Account
func (r *ExternalAccountService) New(ctx context.Context, body *types.CreateAnExternalAccountParameters, opts ...*core.RequestOpts) (res *types.ExternalAccount, err error) {
	path := "/external_accounts"
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
		Body:   body,
	}
	err = r.post(ctx, path, req, &res)

	return
}

// Retrieve an External Account
func (r *ExternalAccountService) Get(ctx context.Context, external_account_id string, opts ...*core.RequestOpts) (res *types.ExternalAccount, err error) {
	path := fmt.Sprintf("/external_accounts/%s", external_account_id)
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
	}
	err = r.get(ctx, path, req, &res)

	return
}

// Update an External Account
func (r *ExternalAccountService) Update(ctx context.Context, external_account_id string, body *types.UpdateAnExternalAccountParameters, opts ...*core.RequestOpts) (res *types.ExternalAccount, err error) {
	path := fmt.Sprintf("/external_accounts/%s", external_account_id)
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
		Body:   body,
	}
	err = r.patch(ctx, path, req, &res)

	return
}

// List External Accounts
func (r *ExternalAccountService) List(ctx context.Context, query *types.ExternalAccountListParams, opts ...*core.RequestOpts) (res *types.ExternalAccountsPage, err error) {
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

package services

import (
	"context"
	"fmt"
	"increase/core"
	"increase/pagination"
	"increase/types"
)

type AccountService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewAccountService(requester core.Requester) (r *AccountService) {
	r = &AccountService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

// Create an Account
func (r *AccountService) New(ctx context.Context, body *types.CreateAnAccountParameters, opts ...*core.RequestOpts) (res *types.Account, err error) {
	path := "/accounts"
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
		Body:   body,
	}
	err = r.post(ctx, path, req, &res)

	return
}

// Retrieve an Account
func (r *AccountService) Get(ctx context.Context, account_id string, opts ...*core.RequestOpts) (res *types.Account, err error) {
	path := fmt.Sprintf("/accounts/%s", account_id)
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
	}
	err = r.get(ctx, path, req, &res)

	return
}

// Update an Account
func (r *AccountService) Update(ctx context.Context, account_id string, body *types.UpdateAnAccountParameters, opts ...*core.RequestOpts) (res *types.Account, err error) {
	path := fmt.Sprintf("/accounts/%s", account_id)
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
		Body:   body,
	}
	err = r.patch(ctx, path, req, &res)

	return
}

// List Accounts
func (r *AccountService) List(ctx context.Context, query *types.AccountListParams, opts ...*core.RequestOpts) (res *types.AccountsPage, err error) {
	page := &types.AccountsPage{
		Page: &pagination.Page[types.Account]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/accounts",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

// Close an Account
func (r *AccountService) Close(ctx context.Context, account_id string, opts ...*core.RequestOpts) (res *types.Account, err error) {
	path := fmt.Sprintf("/accounts/%s/close", account_id)
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
	}
	err = r.post(ctx, path, req, &res)

	return
}

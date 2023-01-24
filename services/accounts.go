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

func (r *AccountService) Create(ctx context.Context, body *types.CreateAnAccountParameters, opts ...*core.RequestOpts) (res *types.Account, err error) {
	err = r.post(
		ctx,
		"/accounts",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)

	return
}

func (r *AccountService) Retrieve(ctx context.Context, account_id string, opts ...*core.RequestOpts) (res *types.Account, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/accounts/%s", account_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)

	return
}

func (r *AccountService) Update(ctx context.Context, account_id string, body *types.UpdateAnAccountParameters, opts ...*core.RequestOpts) (res *types.Account, err error) {
	err = r.patch(
		ctx,
		fmt.Sprintf("/accounts/%s", account_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)

	return
}

func (r *AccountService) List(ctx context.Context, query *types.ListAccountsQuery, opts ...*core.RequestOpts) (res *types.AccountsPage, err error) {
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

func (r *AccountService) Close(ctx context.Context, account_id string, opts ...*core.RequestOpts) (res *types.Account, err error) {
	err = r.post(
		ctx,
		fmt.Sprintf("/accounts/%s/close", account_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)

	return
}

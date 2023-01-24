package services

import (
	"context"
	"fmt"
	"increase/core"
	"increase/pagination"
	"increase/types"
)

type AccountNumberService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewAccountNumberService(requester core.Requester) (r *AccountNumberService) {
	r = &AccountNumberService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

func (r *AccountNumberService) Create(ctx context.Context, body *types.CreateAnAccountNumberParameters, opts ...*core.RequestOpts) (res *types.AccountNumber, err error) {
	err = r.post(
		ctx,
		"/account_numbers",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)

	return
}

func (r *AccountNumberService) Retrieve(ctx context.Context, account_number_id string, opts ...*core.RequestOpts) (res *types.AccountNumber, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/account_numbers/%s", account_number_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)

	return
}

func (r *AccountNumberService) Update(ctx context.Context, account_number_id string, body *types.UpdateAnAccountNumberParameters, opts ...*core.RequestOpts) (res *types.AccountNumber, err error) {
	err = r.patch(
		ctx,
		fmt.Sprintf("/account_numbers/%s", account_number_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)

	return
}

func (r *AccountNumberService) List(ctx context.Context, query *types.ListAccountNumbersQuery, opts ...*core.RequestOpts) (res *types.AccountNumbersPage, err error) {
	page := &types.AccountNumbersPage{
		Page: &pagination.Page[types.AccountNumber]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/account_numbers",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

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

// Create an Account Number
func (r *AccountNumberService) New(ctx context.Context, body *types.CreateAnAccountNumberParameters, opts ...*core.RequestOpts) (res *types.AccountNumber, err error) {
	path := "/account_numbers"
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
		Body:   body,
	}
	err = r.post(ctx, path, req, &res)

	return
}

// Retrieve an Account Number
func (r *AccountNumberService) Get(ctx context.Context, account_number_id string, opts ...*core.RequestOpts) (res *types.AccountNumber, err error) {
	path := fmt.Sprintf("/account_numbers/%s", account_number_id)
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
	}
	err = r.get(ctx, path, req, &res)

	return
}

// Update an Account Number
func (r *AccountNumberService) Update(ctx context.Context, account_number_id string, body *types.UpdateAnAccountNumberParameters, opts ...*core.RequestOpts) (res *types.AccountNumber, err error) {
	path := fmt.Sprintf("/account_numbers/%s", account_number_id)
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
		Body:   body,
	}
	err = r.patch(ctx, path, req, &res)

	return
}

// List Account Numbers
func (r *AccountNumberService) List(ctx context.Context, query *types.AccountNumberListParams, opts ...*core.RequestOpts) (res *types.AccountNumbersPage, err error) {
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

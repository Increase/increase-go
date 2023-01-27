package services

import (
	"context"
	"fmt"
	"increase/core"
	"increase/pagination"
	"increase/types"
)

type AccountTransferService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewAccountTransferService(requester core.Requester) (r *AccountTransferService) {
	r = &AccountTransferService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

// Create an Account Transfer
func (r *AccountTransferService) New(ctx context.Context, body *types.CreateAnAccountTransferParameters, opts ...*core.RequestOpts) (res *types.AccountTransfer, err error) {
	path := "/account_transfers"
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
		Body:   body,
	}
	err = r.post(ctx, path, req, &res)

	return
}

// Retrieve an Account Transfer
func (r *AccountTransferService) Get(ctx context.Context, account_transfer_id string, opts ...*core.RequestOpts) (res *types.AccountTransfer, err error) {
	path := fmt.Sprintf("/account_transfers/%s", account_transfer_id)
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
	}
	err = r.get(ctx, path, req, &res)

	return
}

// List Account Transfers
func (r *AccountTransferService) List(ctx context.Context, query *types.ListAccountTransfersQuery, opts ...*core.RequestOpts) (res *types.AccountTransfersPage, err error) {
	page := &types.AccountTransfersPage{
		Page: &pagination.Page[types.AccountTransfer]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/account_transfers",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

// Approve an Account Transfer
func (r *AccountTransferService) Approve(ctx context.Context, account_transfer_id string, opts ...*core.RequestOpts) (res *types.AccountTransfer, err error) {
	path := fmt.Sprintf("/account_transfers/%s/approve", account_transfer_id)
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
	}
	err = r.post(ctx, path, req, &res)

	return
}

// Cancel an Account Transfer
func (r *AccountTransferService) Cancel(ctx context.Context, account_transfer_id string, opts ...*core.RequestOpts) (res *types.AccountTransfer, err error) {
	path := fmt.Sprintf("/account_transfers/%s/cancel", account_transfer_id)
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
	}
	err = r.post(ctx, path, req, &res)

	return
}

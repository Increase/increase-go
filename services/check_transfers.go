package services

import (
	"context"
	"fmt"
	"increase/core"
	"increase/pagination"
	"increase/types"
)

type CheckTransferService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewCheckTransferService(requester core.Requester) (r *CheckTransferService) {
	r = &CheckTransferService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

// Create a Check Transfer
func (r *CheckTransferService) New(ctx context.Context, body *types.CreateACheckTransferParameters, opts ...*core.RequestOpts) (res *types.CheckTransfer, err error) {
	path := "/check_transfers"
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
		Body:   body,
	}
	err = r.post(ctx, path, req, &res)

	return
}

// Retrieve a Check Transfer
func (r *CheckTransferService) Get(ctx context.Context, check_transfer_id string, opts ...*core.RequestOpts) (res *types.CheckTransfer, err error) {
	path := fmt.Sprintf("/check_transfers/%s", check_transfer_id)
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
	}
	err = r.get(ctx, path, req, &res)

	return
}

// List Check Transfers
func (r *CheckTransferService) List(ctx context.Context, query *types.CheckTransferListParams, opts ...*core.RequestOpts) (res *types.CheckTransfersPage, err error) {
	page := &types.CheckTransfersPage{
		Page: &pagination.Page[types.CheckTransfer]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/check_transfers",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

// Approve a Check Transfer
func (r *CheckTransferService) Approve(ctx context.Context, check_transfer_id string, opts ...*core.RequestOpts) (res *types.CheckTransfer, err error) {
	path := fmt.Sprintf("/check_transfers/%s/approve", check_transfer_id)
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
	}
	err = r.post(ctx, path, req, &res)

	return
}

// Cancel a pending Check Transfer
func (r *CheckTransferService) Cancel(ctx context.Context, check_transfer_id string, opts ...*core.RequestOpts) (res *types.CheckTransfer, err error) {
	path := fmt.Sprintf("/check_transfers/%s/cancel", check_transfer_id)
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
	}
	err = r.post(ctx, path, req, &res)

	return
}

// Request a stop payment on a Check Transfer
func (r *CheckTransferService) StopPayment(ctx context.Context, check_transfer_id string, opts ...*core.RequestOpts) (res *types.CheckTransfer, err error) {
	path := fmt.Sprintf("/check_transfers/%s/stop_payment", check_transfer_id)
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
	}
	err = r.post(ctx, path, req, &res)

	return
}

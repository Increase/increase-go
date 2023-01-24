package services

import (
	"context"
	"fmt"
	"increase/core"
	"increase/pagination"
	"increase/types"
)

type ACHTransferService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewACHTransferService(requester core.Requester) (r *ACHTransferService) {
	r = &ACHTransferService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

func (r *ACHTransferService) Create(ctx context.Context, body *types.CreateAnACHTransferParameters, opts ...*core.RequestOpts) (res *types.ACHTransfer, err error) {
	err = r.post(
		ctx,
		"/ach_transfers",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)

	return
}

func (r *ACHTransferService) Retrieve(ctx context.Context, ach_transfer_id string, opts ...*core.RequestOpts) (res *types.ACHTransfer, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/ach_transfers/%s", ach_transfer_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)

	return
}

func (r *ACHTransferService) List(ctx context.Context, query *types.ListACHTransfersQuery, opts ...*core.RequestOpts) (res *types.ACHTransfersPage, err error) {
	page := &types.ACHTransfersPage{
		Page: &pagination.Page[types.ACHTransfer]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/ach_transfers",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

func (r *ACHTransferService) Approve(ctx context.Context, ach_transfer_id string, opts ...*core.RequestOpts) (res *types.ACHTransfer, err error) {
	err = r.post(
		ctx,
		fmt.Sprintf("/ach_transfers/%s/approve", ach_transfer_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)

	return
}

func (r *ACHTransferService) Cancel(ctx context.Context, ach_transfer_id string, opts ...*core.RequestOpts) (res *types.ACHTransfer, err error) {
	err = r.post(
		ctx,
		fmt.Sprintf("/ach_transfers/%s/cancel", ach_transfer_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)

	return
}

package services

import (
	"context"
	"fmt"
	"increase/core"
	"increase/pagination"
	"increase/types"
)

type WireTransferService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewWireTransferService(requester core.Requester) (r *WireTransferService) {
	r = &WireTransferService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

// Create a Wire Transfer
func (r *WireTransferService) New(ctx context.Context, body *types.CreateAWireTransferParameters, opts ...*core.RequestOpts) (res *types.WireTransfer, err error) {
	path := "/wire_transfers"
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
		Body:   body,
	}
	err = r.post(ctx, path, req, &res)

	return
}

// Retrieve a Wire Transfer
func (r *WireTransferService) Get(ctx context.Context, wire_transfer_id string, opts ...*core.RequestOpts) (res *types.WireTransfer, err error) {
	path := fmt.Sprintf("/wire_transfers/%s", wire_transfer_id)
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
	}
	err = r.get(ctx, path, req, &res)

	return
}

// List Wire Transfers
func (r *WireTransferService) List(ctx context.Context, query *types.ListWireTransfersQuery, opts ...*core.RequestOpts) (res *types.WireTransfersPage, err error) {
	page := &types.WireTransfersPage{
		Page: &pagination.Page[types.WireTransfer]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/wire_transfers",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

// Approve a Wire Transfer
func (r *WireTransferService) Approve(ctx context.Context, wire_transfer_id string, opts ...*core.RequestOpts) (res *types.WireTransfer, err error) {
	path := fmt.Sprintf("/wire_transfers/%s/approve", wire_transfer_id)
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
	}
	err = r.post(ctx, path, req, &res)

	return
}

// Cancel a pending Wire Transfer
func (r *WireTransferService) Cancel(ctx context.Context, wire_transfer_id string, opts ...*core.RequestOpts) (res *types.WireTransfer, err error) {
	path := fmt.Sprintf("/wire_transfers/%s/cancel", wire_transfer_id)
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
	}
	err = r.post(ctx, path, req, &res)

	return
}

// Simulates the reversal of a [Wire Transfer](#wire-transfers) by the Federal
// Reserve due to error conditions. This will also create a
// [Transaction](#transaction) to account for the returned funds. This Wire
// Transfer must first have a `status` of `complete`.'
func (r *WireTransferService) Reverse(ctx context.Context, wire_transfer_id string, opts ...*core.RequestOpts) (res *types.WireTransfer, err error) {
	path := fmt.Sprintf("/simulations/wire_transfers/%s/reverse", wire_transfer_id)
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
	}
	err = r.post(ctx, path, req, &res)

	return
}

// Simulates the submission of a [Wire Transfer](#wire-transfers) to the Federal
// Reserve. This transfer must first have a `status` of `pending_approval` or
// `pending_creating`.
func (r *WireTransferService) Submit(ctx context.Context, wire_transfer_id string, opts ...*core.RequestOpts) (res *types.WireTransfer, err error) {
	path := fmt.Sprintf("/simulations/wire_transfers/%s/submit", wire_transfer_id)
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
	}
	err = r.post(ctx, path, req, &res)

	return
}

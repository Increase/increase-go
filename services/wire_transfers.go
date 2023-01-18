package services

import "context"
import "increase/core"
import "increase/types"
import "fmt"
import "increase/pagination"

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

func (r *WireTransferService) Create(ctx context.Context, body *types.CreateAWireTransferParameters, opts ...*core.RequestOpts) (res *types.WireTransfer, err error) {
	err = r.post(
		ctx,
		"/wire_transfers",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)

	return
}

func (r *WireTransferService) Retrieve(ctx context.Context, wire_transfer_id string, opts ...*core.RequestOpts) (res *types.WireTransfer, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/wire_transfers/%s", wire_transfer_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)

	return
}

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

// Simulates the reversal of an Wire Transfer by the Federal Reserve due to error
// conditions. This will also create a Transaction to account for the returned
// funds. This transfer must first have a `status` of `complete`.
func (r *WireTransferService) Reverse(ctx context.Context, wire_transfer_id string, opts ...*core.RequestOpts) (res *types.WireTransfer, err error) {
	err = r.post(
		ctx,
		fmt.Sprintf("/simulations/wire_transfers/%s/reverse", wire_transfer_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)

	return
}

// Simulates the submission of a Wire Transfer to the Federal Reserve. This
// transfer must first have a `status` of `pending_approval` or `pending_creating`.
func (r *WireTransferService) Submit(ctx context.Context, wire_transfer_id string, opts ...*core.RequestOpts) (res *types.WireTransfer, err error) {
	err = r.post(
		ctx,
		fmt.Sprintf("/simulations/wire_transfers/%s/submit", wire_transfer_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)

	return
}

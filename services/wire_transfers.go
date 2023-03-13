package services

import (
	"context"
	"fmt"

	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/pagination"
	"github.com/increase/increase-go/types"
)

type WireTransferService struct {
	Options []options.RequestOption
}

func NewWireTransferService(opts ...options.RequestOption) (r *WireTransferService) {
	r = &WireTransferService{}
	r.Options = opts
	return
}

// Create a Wire Transfer
func (r *WireTransferService) New(ctx context.Context, body *types.CreateAWireTransferParameters, opts ...options.RequestOption) (res *types.WireTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := "wire_transfers"
	err = options.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Retrieve a Wire Transfer
func (r *WireTransferService) Get(ctx context.Context, wire_transfer_id string, opts ...options.RequestOption) (res *types.WireTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("wire_transfers/%s", wire_transfer_id)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// List Wire Transfers
func (r *WireTransferService) List(ctx context.Context, query *types.WireTransferListParams, opts ...options.RequestOption) (res *types.WireTransfersPage, err error) {
	opts = append(r.Options, opts...)
	path := "wire_transfers"
	cfg, err := options.NewRequestConfig(ctx, "GET", path, query, nil, opts...)
	if err != nil {
		return
	}
	res = &types.WireTransfersPage{
		Page: &pagination.Page[types.WireTransfer]{
			Config:  *cfg,
			Options: opts,
		},
	}
	return res, res.Fire()
}

// Approve a Wire Transfer
func (r *WireTransferService) Approve(ctx context.Context, wire_transfer_id string, opts ...options.RequestOption) (res *types.WireTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("wire_transfers/%s/approve", wire_transfer_id)
	err = options.ExecuteNewRequest(ctx, "POST", path, nil, &res, opts...)
	return
}

// Cancel a pending Wire Transfer
func (r *WireTransferService) Cancel(ctx context.Context, wire_transfer_id string, opts ...options.RequestOption) (res *types.WireTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("wire_transfers/%s/cancel", wire_transfer_id)
	err = options.ExecuteNewRequest(ctx, "POST", path, nil, &res, opts...)
	return
}

// Simulates the reversal of a [Wire Transfer](#wire-transfers) by the Federal
// Reserve due to error conditions. This will also create a
// [Transaction](#transaction) to account for the returned funds. This Wire
// Transfer must first have a `status` of `complete`.'
func (r *WireTransferService) Reverse(ctx context.Context, wire_transfer_id string, opts ...options.RequestOption) (res *types.WireTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("simulations/wire_transfers/%s/reverse", wire_transfer_id)
	err = options.ExecuteNewRequest(ctx, "POST", path, nil, &res, opts...)
	return
}

// Simulates the submission of a [Wire Transfer](#wire-transfers) to the Federal
// Reserve. This transfer must first have a `status` of `pending_approval` or
// `pending_creating`.
func (r *WireTransferService) Submit(ctx context.Context, wire_transfer_id string, opts ...options.RequestOption) (res *types.WireTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("simulations/wire_transfers/%s/submit", wire_transfer_id)
	err = options.ExecuteNewRequest(ctx, "POST", path, nil, &res, opts...)
	return
}

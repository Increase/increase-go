package services

import (
	"context"
	"fmt"

	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/pagination"
	"github.com/increase/increase-go/types"
)

type CheckTransferService struct {
	Options []options.RequestOption
}

func NewCheckTransferService(opts ...options.RequestOption) (r *CheckTransferService) {
	r = &CheckTransferService{}
	r.Options = opts
	return
}

// Create a Check Transfer
func (r *CheckTransferService) New(ctx context.Context, body *types.CreateACheckTransferParameters, opts ...options.RequestOption) (res *types.CheckTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := "check_transfers"
	err = options.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Retrieve a Check Transfer
func (r *CheckTransferService) Get(ctx context.Context, check_transfer_id string, opts ...options.RequestOption) (res *types.CheckTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("check_transfers/%s", check_transfer_id)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// List Check Transfers
func (r *CheckTransferService) List(ctx context.Context, query *types.CheckTransferListParams, opts ...options.RequestOption) (res *types.CheckTransfersPage, err error) {
	opts = append(r.Options, opts...)
	path := "check_transfers"
	cfg, err := options.NewRequestConfig(ctx, "GET", path, query, nil, opts...)
	if err != nil {
		return
	}
	res = &types.CheckTransfersPage{
		Page: &pagination.Page[types.CheckTransfer]{
			Config:  *cfg,
			Options: opts,
		},
	}
	return res, res.Fire()
}

// Approve a Check Transfer
func (r *CheckTransferService) Approve(ctx context.Context, check_transfer_id string, opts ...options.RequestOption) (res *types.CheckTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("check_transfers/%s/approve", check_transfer_id)
	err = options.ExecuteNewRequest(ctx, "POST", path, nil, &res, opts...)
	return
}

// Cancel a pending Check Transfer
func (r *CheckTransferService) Cancel(ctx context.Context, check_transfer_id string, opts ...options.RequestOption) (res *types.CheckTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("check_transfers/%s/cancel", check_transfer_id)
	err = options.ExecuteNewRequest(ctx, "POST", path, nil, &res, opts...)
	return
}

// Request a stop payment on a Check Transfer
func (r *CheckTransferService) StopPayment(ctx context.Context, check_transfer_id string, opts ...options.RequestOption) (res *types.CheckTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("check_transfers/%s/stop_payment", check_transfer_id)
	err = options.ExecuteNewRequest(ctx, "POST", path, nil, &res, opts...)
	return
}

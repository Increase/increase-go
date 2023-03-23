package services

import (
	"context"
	"fmt"

	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/pagination"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type ACHTransferService struct {
	Options []options.RequestOption
}

func NewACHTransferService(opts ...options.RequestOption) (r *ACHTransferService) {
	r = &ACHTransferService{}
	r.Options = opts
	return
}

// Create an ACH Transfer
func (r *ACHTransferService) New(ctx context.Context, body *requests.CreateAnACHTransferParameters, opts ...options.RequestOption) (res *responses.ACHTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := "ach_transfers"
	err = options.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Retrieve an ACH Transfer
func (r *ACHTransferService) Get(ctx context.Context, ach_transfer_id string, opts ...options.RequestOption) (res *responses.ACHTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("ach_transfers/%s", ach_transfer_id)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// List ACH Transfers
func (r *ACHTransferService) List(ctx context.Context, query *requests.ACHTransferListParams, opts ...options.RequestOption) (res *responses.ACHTransfersPage, err error) {
	opts = append(r.Options, opts...)
	path := "ach_transfers"
	cfg, err := options.NewRequestConfig(ctx, "GET", path, query, nil, opts...)
	if err != nil {
		return
	}
	res = &responses.ACHTransfersPage{
		Page: &pagination.Page[responses.ACHTransfer]{
			Config:  *cfg,
			Options: opts,
		},
	}
	return res, res.Fire()
}

// Approves an ACH Transfer in a pending_approval state.
func (r *ACHTransferService) Approve(ctx context.Context, ach_transfer_id string, opts ...options.RequestOption) (res *responses.ACHTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("ach_transfers/%s/approve", ach_transfer_id)
	err = options.ExecuteNewRequest(ctx, "POST", path, nil, &res, opts...)
	return
}

// Cancels an ACH Transfer in a pending_approval state.
func (r *ACHTransferService) Cancel(ctx context.Context, ach_transfer_id string, opts ...options.RequestOption) (res *responses.ACHTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("ach_transfers/%s/cancel", ach_transfer_id)
	err = options.ExecuteNewRequest(ctx, "POST", path, nil, &res, opts...)
	return
}

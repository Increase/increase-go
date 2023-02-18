package services

import (
	"bytes"
	"context"
	"fmt"
	"increase/options"
	"increase/pagination"
	"increase/types"
	"io"
	"net/url"
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
func (r *ACHTransferService) New(ctx context.Context, body *types.CreateAnACHTransferParameters, opts ...options.RequestOption) (res *types.ACHTransfer, err error) {
	opts = append(r.Options[:], opts...)
	b, err := body.MarshalJSON()
	content := io.NopCloser(bytes.NewBuffer(b))
	u, err := url.Parse(fmt.Sprintf("ach_transfers"))
	if err != nil {
		return
	}
	cfg, err := options.NewRequestConfig(ctx, "POST", u, content, opts...)
	if err != nil {
		return
	}
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return
	}

	return
}

// Retrieve an ACH Transfer
func (r *ACHTransferService) Get(ctx context.Context, ach_transfer_id string, opts ...options.RequestOption) (res *types.ACHTransfer, err error) {
	opts = append(r.Options[:], opts...)
	u, err := url.Parse(fmt.Sprintf("ach_transfers/%s", ach_transfer_id))
	if err != nil {
		return
	}
	cfg, err := options.NewRequestConfig(ctx, "GET", u, nil, opts...)
	if err != nil {
		return
	}
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return
	}

	return
}

// List ACH Transfers
func (r *ACHTransferService) List(ctx context.Context, query *types.ACHTransferListParams, opts ...options.RequestOption) (res *types.ACHTransfersPage, err error) {
	u, err := url.Parse(fmt.Sprintf("ach_transfers"))
	if err != nil {
		return
	}
	opts = append(r.Options, opts...)
	cfg, err := options.NewRequestConfig(ctx, "GET", u, nil, opts...)
	if err != nil {
		return
	}
	res = &types.ACHTransfersPage{
		Page: &pagination.Page[types.ACHTransfer]{
			Config:  *cfg,
			Options: opts,
		},
	}
	err = res.Fire()
	return
}

// Approve an ACH Transfer
func (r *ACHTransferService) Approve(ctx context.Context, ach_transfer_id string, opts ...options.RequestOption) (res *types.ACHTransfer, err error) {
	opts = append(r.Options[:], opts...)
	u, err := url.Parse(fmt.Sprintf("ach_transfers/%s/approve", ach_transfer_id))
	if err != nil {
		return
	}
	cfg, err := options.NewRequestConfig(ctx, "POST", u, nil, opts...)
	if err != nil {
		return
	}
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return
	}

	return
}

// Cancel a pending ACH Transfer
func (r *ACHTransferService) Cancel(ctx context.Context, ach_transfer_id string, opts ...options.RequestOption) (res *types.ACHTransfer, err error) {
	opts = append(r.Options[:], opts...)
	u, err := url.Parse(fmt.Sprintf("ach_transfers/%s/cancel", ach_transfer_id))
	if err != nil {
		return
	}
	cfg, err := options.NewRequestConfig(ctx, "POST", u, nil, opts...)
	if err != nil {
		return
	}
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return
	}

	return
}

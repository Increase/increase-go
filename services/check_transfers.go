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
	b, err := body.MarshalJSON()
	content := io.NopCloser(bytes.NewBuffer(b))
	u, err := url.Parse(fmt.Sprintf("check_transfers"))
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

// Retrieve a Check Transfer
func (r *CheckTransferService) Get(ctx context.Context, check_transfer_id string, opts ...options.RequestOption) (res *types.CheckTransfer, err error) {
	opts = append(r.Options[:], opts...)
	u, err := url.Parse(fmt.Sprintf("check_transfers/%s", check_transfer_id))
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

// List Check Transfers
func (r *CheckTransferService) List(ctx context.Context, query *types.CheckTransferListParams, opts ...options.RequestOption) (res *types.CheckTransfersPage, err error) {
	u, err := url.Parse(fmt.Sprintf("check_transfers"))
	if err != nil {
		return
	}
	opts = append(r.Options, opts...)
	cfg, err := options.NewRequestConfig(ctx, "GET", u, nil, opts...)
	if err != nil {
		return
	}
	res = &types.CheckTransfersPage{
		Page: &pagination.Page[types.CheckTransfer]{
			Config:  *cfg,
			Options: opts,
		},
	}
	err = res.Fire()
	return
}

// Approve a Check Transfer
func (r *CheckTransferService) Approve(ctx context.Context, check_transfer_id string, opts ...options.RequestOption) (res *types.CheckTransfer, err error) {
	opts = append(r.Options[:], opts...)
	u, err := url.Parse(fmt.Sprintf("check_transfers/%s/approve", check_transfer_id))
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

// Cancel a pending Check Transfer
func (r *CheckTransferService) Cancel(ctx context.Context, check_transfer_id string, opts ...options.RequestOption) (res *types.CheckTransfer, err error) {
	opts = append(r.Options[:], opts...)
	u, err := url.Parse(fmt.Sprintf("check_transfers/%s/cancel", check_transfer_id))
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

// Request a stop payment on a Check Transfer
func (r *CheckTransferService) StopPayment(ctx context.Context, check_transfer_id string, opts ...options.RequestOption) (res *types.CheckTransfer, err error) {
	opts = append(r.Options[:], opts...)
	u, err := url.Parse(fmt.Sprintf("check_transfers/%s/stop_payment", check_transfer_id))
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

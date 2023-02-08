package services

import (
	"context"
	"fmt"
	"increase/options"
	"increase/pagination"
	"increase/types"
	"net/url"
)

type AccountTransferService struct {
	Options []options.RequestOption
}

func NewAccountTransferService(opts ...options.RequestOption) (r *AccountTransferService) {
	r = &AccountTransferService{}
	r.Options = opts
	return
}

// Create an Account Transfer
func (r *AccountTransferService) New(ctx context.Context, body *types.CreateAnAccountTransferParameters, opts ...options.RequestOption) (res *types.AccountTransfer, err error) {
	opts = append(r.Options, opts...)
	u, err := url.Parse(fmt.Sprintf("account_transfers"))
	if err != nil {
		return
	}
	cfg := options.NewRequestConfig(ctx, "POST", u, opts...)
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return
	}

	return
}

// Retrieve an Account Transfer
func (r *AccountTransferService) Get(ctx context.Context, account_transfer_id string, opts ...options.RequestOption) (res *types.AccountTransfer, err error) {
	opts = append(r.Options, opts...)
	u, err := url.Parse(fmt.Sprintf("account_transfers/%s", account_transfer_id))
	if err != nil {
		return
	}
	cfg := options.NewRequestConfig(ctx, "GET", u, opts...)
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return
	}

	return
}

// List Account Transfers
func (r *AccountTransferService) List(ctx context.Context, query *types.AccountTransferListParams, opts ...options.RequestOption) (res *types.AccountTransfersPage, err error) {
	u, err := url.Parse(fmt.Sprintf("account_transfers"))
	if err != nil {
		return
	}
	opts = append(r.Options, opts...)
	cfg := options.NewRequestConfig(ctx, "GET", u, opts...)
	res = &types.AccountTransfersPage{
		Page: &pagination.Page[types.AccountTransfer]{
			Config:  *cfg,
			Options: opts,
		},
	}
	err = res.Fire()
	return
}

// Approve an Account Transfer
func (r *AccountTransferService) Approve(ctx context.Context, account_transfer_id string, opts ...options.RequestOption) (res *types.AccountTransfer, err error) {
	opts = append(r.Options, opts...)
	u, err := url.Parse(fmt.Sprintf("account_transfers/%s/approve", account_transfer_id))
	if err != nil {
		return
	}
	cfg := options.NewRequestConfig(ctx, "POST", u, opts...)
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return
	}

	return
}

// Cancel an Account Transfer
func (r *AccountTransferService) Cancel(ctx context.Context, account_transfer_id string, opts ...options.RequestOption) (res *types.AccountTransfer, err error) {
	opts = append(r.Options, opts...)
	u, err := url.Parse(fmt.Sprintf("account_transfers/%s/cancel", account_transfer_id))
	if err != nil {
		return
	}
	cfg := options.NewRequestConfig(ctx, "POST", u, opts...)
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return
	}

	return
}

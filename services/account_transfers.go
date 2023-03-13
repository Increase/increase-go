package services

import (
	"context"
	"fmt"

	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/pagination"
	"github.com/increase/increase-go/types"
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
	opts = append(r.Options[:], opts...)
	path := "account_transfers"
	err = options.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Retrieve an Account Transfer
func (r *AccountTransferService) Get(ctx context.Context, account_transfer_id string, opts ...options.RequestOption) (res *types.AccountTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("account_transfers/%s", account_transfer_id)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// List Account Transfers
func (r *AccountTransferService) List(ctx context.Context, query *types.AccountTransferListParams, opts ...options.RequestOption) (res *types.AccountTransfersPage, err error) {
	opts = append(r.Options, opts...)
	path := "account_transfers"
	cfg, err := options.NewRequestConfig(ctx, "GET", path, query, nil, opts...)
	if err != nil {
		return
	}
	res = &types.AccountTransfersPage{
		Page: &pagination.Page[types.AccountTransfer]{
			Config:  *cfg,
			Options: opts,
		},
	}
	return res, res.Fire()
}

// Approve an Account Transfer
func (r *AccountTransferService) Approve(ctx context.Context, account_transfer_id string, opts ...options.RequestOption) (res *types.AccountTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("account_transfers/%s/approve", account_transfer_id)
	err = options.ExecuteNewRequest(ctx, "POST", path, nil, &res, opts...)
	return
}

// Cancel an Account Transfer
func (r *AccountTransferService) Cancel(ctx context.Context, account_transfer_id string, opts ...options.RequestOption) (res *types.AccountTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("account_transfers/%s/cancel", account_transfer_id)
	err = options.ExecuteNewRequest(ctx, "POST", path, nil, &res, opts...)
	return
}

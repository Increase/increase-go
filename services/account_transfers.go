package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type AccountTransferService struct {
	Options []option.RequestOption
}

func NewAccountTransferService(opts ...option.RequestOption) (r *AccountTransferService) {
	r = &AccountTransferService{}
	r.Options = opts
	return
}

// Create an Account Transfer
func (r *AccountTransferService) New(ctx context.Context, body requests.AccountTransferNewParams, opts ...option.RequestOption) (res *responses.AccountTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := "account_transfers"
	err = option.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve an Account Transfer
func (r *AccountTransferService) Get(ctx context.Context, account_transfer_id string, opts ...option.RequestOption) (res *responses.AccountTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("account_transfers/%s", account_transfer_id)
	err = option.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Account Transfers
func (r *AccountTransferService) List(ctx context.Context, query requests.AccountTransferListParams, opts ...option.RequestOption) (res *responses.Page[responses.AccountTransfer], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "account_transfers"
	cfg, err := option.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List Account Transfers
func (r *AccountTransferService) ListAutoPaging(ctx context.Context, query requests.AccountTransferListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.AccountTransfer] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Approve an Account Transfer
func (r *AccountTransferService) Approve(ctx context.Context, account_transfer_id string, opts ...option.RequestOption) (res *responses.AccountTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("account_transfers/%s/approve", account_transfer_id)
	err = option.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Cancel an Account Transfer
func (r *AccountTransferService) Cancel(ctx context.Context, account_transfer_id string, opts ...option.RequestOption) (res *responses.AccountTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("account_transfers/%s/cancel", account_transfer_id)
	err = option.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

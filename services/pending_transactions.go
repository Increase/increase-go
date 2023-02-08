package services

import (
	"context"
	"fmt"
	"increase/options"
	"increase/pagination"
	"increase/types"
	"net/url"
)

type PendingTransactionService struct {
	Options []options.RequestOption
}

func NewPendingTransactionService(opts ...options.RequestOption) (r *PendingTransactionService) {
	r = &PendingTransactionService{}
	r.Options = opts
	return
}

// Retrieve a Pending Transaction
func (r *PendingTransactionService) Get(ctx context.Context, pending_transaction_id string, opts ...options.RequestOption) (res *types.PendingTransaction, err error) {
	opts = append(r.Options, opts...)
	u, err := url.Parse(fmt.Sprintf("pending_transactions/%s", pending_transaction_id))
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

// List Pending Transactions
func (r *PendingTransactionService) List(ctx context.Context, query *types.PendingTransactionListParams, opts ...options.RequestOption) (res *types.PendingTransactionsPage, err error) {
	u, err := url.Parse(fmt.Sprintf("pending_transactions"))
	if err != nil {
		return
	}
	opts = append(r.Options, opts...)
	cfg := options.NewRequestConfig(ctx, "GET", u, opts...)
	res = &types.PendingTransactionsPage{
		Page: &pagination.Page[types.PendingTransaction]{
			Config:  *cfg,
			Options: opts,
		},
	}
	err = res.Fire()
	return
}

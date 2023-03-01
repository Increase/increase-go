package services

import (
	"context"
	"fmt"
	"increase/options"
	"increase/pagination"
	"increase/types"
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
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("pending_transactions/%s", pending_transaction_id)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// List Pending Transactions
func (r *PendingTransactionService) List(ctx context.Context, query *types.PendingTransactionListParams, opts ...options.RequestOption) (res *types.PendingTransactionsPage, err error) {
	opts = append(r.Options, opts...)
	path := "pending_transactions"
	cfg, err := options.NewRequestConfig(ctx, "GET", path, query, nil, opts...)
	if err != nil {
		return
	}
	res = &types.PendingTransactionsPage{
		Page: &pagination.Page[types.PendingTransaction]{
			Config:  *cfg,
			Options: opts,
		},
	}
	return res, res.Fire()
}

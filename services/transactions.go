package services

import (
	"context"
	"fmt"
	"increase/options"
	"increase/pagination"
	"increase/types"
	"net/url"
)

type TransactionService struct {
	Options []options.RequestOption
}

func NewTransactionService(opts ...options.RequestOption) (r *TransactionService) {
	r = &TransactionService{}
	r.Options = opts
	return
}

// Retrieve a Transaction
func (r *TransactionService) Get(ctx context.Context, transaction_id string, opts ...options.RequestOption) (res *types.Transaction, err error) {
	opts = append(r.Options[:], opts...)
	u, err := url.Parse(fmt.Sprintf("transactions/%s", transaction_id))
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

// List Transactions
func (r *TransactionService) List(ctx context.Context, query *types.TransactionListParams, opts ...options.RequestOption) (res *types.TransactionsPage, err error) {
	u, err := url.Parse(fmt.Sprintf("transactions"))
	if err != nil {
		return
	}
	opts = append(r.Options, opts...)
	cfg, err := options.NewRequestConfig(ctx, "GET", u, nil, opts...)
	if err != nil {
		return
	}
	res = &types.TransactionsPage{
		Page: &pagination.Page[types.Transaction]{
			Config:  *cfg,
			Options: opts,
		},
	}
	err = res.Fire()
	return
}

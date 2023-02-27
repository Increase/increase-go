package services

import (
	"context"
	"fmt"
	"increase/options"
	"increase/pagination"
	"increase/types"
	"net/url"
)

type DeclinedTransactionService struct {
	Options []options.RequestOption
}

func NewDeclinedTransactionService(opts ...options.RequestOption) (r *DeclinedTransactionService) {
	r = &DeclinedTransactionService{}
	r.Options = opts
	return
}

// Retrieve a Declined Transaction
func (r *DeclinedTransactionService) Get(ctx context.Context, declined_transaction_id string, opts ...options.RequestOption) (res *types.DeclinedTransaction, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]options.RequestOption{options.WithHeader("Content-Type", "")}, opts...)
	u, err := url.Parse(fmt.Sprintf("declined_transactions/%s", declined_transaction_id))
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

// List Declined Transactions
func (r *DeclinedTransactionService) List(ctx context.Context, query *types.DeclinedTransactionListParams, opts ...options.RequestOption) (res *types.DeclinedTransactionsPage, err error) {
	u, err := url.Parse(fmt.Sprintf("declined_transactions"))
	if err != nil {
		return
	}
	opts = append(r.Options, opts...)
	cfg, err := options.NewRequestConfig(ctx, "GET", u, query, opts...)
	if err != nil {
		return
	}
	res = &types.DeclinedTransactionsPage{
		Page: &pagination.Page[types.DeclinedTransaction]{
			Config:  *cfg,
			Options: opts,
		},
	}
	err = res.Fire()
	return
}

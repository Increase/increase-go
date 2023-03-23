package services

import (
	"context"
	"fmt"

	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/pagination"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
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
func (r *TransactionService) Get(ctx context.Context, transaction_id string, opts ...options.RequestOption) (res *responses.Transaction, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("transactions/%s", transaction_id)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// List Transactions
func (r *TransactionService) List(ctx context.Context, query *requests.TransactionListParams, opts ...options.RequestOption) (res *responses.TransactionsPage, err error) {
	opts = append(r.Options, opts...)
	path := "transactions"
	cfg, err := options.NewRequestConfig(ctx, "GET", path, query, nil, opts...)
	if err != nil {
		return
	}
	res = &responses.TransactionsPage{
		Page: &pagination.Page[responses.Transaction]{
			Config:  *cfg,
			Options: opts,
		},
	}
	return res, res.Fire()
}

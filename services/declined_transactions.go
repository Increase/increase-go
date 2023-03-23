package services

import (
	"context"
	"fmt"

	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/pagination"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
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
func (r *DeclinedTransactionService) Get(ctx context.Context, declined_transaction_id string, opts ...options.RequestOption) (res *responses.DeclinedTransaction, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("declined_transactions/%s", declined_transaction_id)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// List Declined Transactions
func (r *DeclinedTransactionService) List(ctx context.Context, query *requests.DeclinedTransactionListParams, opts ...options.RequestOption) (res *responses.DeclinedTransactionsPage, err error) {
	opts = append(r.Options, opts...)
	path := "declined_transactions"
	cfg, err := options.NewRequestConfig(ctx, "GET", path, query, nil, opts...)
	if err != nil {
		return
	}
	res = &responses.DeclinedTransactionsPage{
		Page: &pagination.Page[responses.DeclinedTransaction]{
			Config:  *cfg,
			Options: opts,
		},
	}
	return res, res.Fire()
}

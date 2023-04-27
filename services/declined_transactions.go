package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type DeclinedTransactionService struct {
	Options []option.RequestOption
}

func NewDeclinedTransactionService(opts ...option.RequestOption) (r *DeclinedTransactionService) {
	r = &DeclinedTransactionService{}
	r.Options = opts
	return
}

// Retrieve a Declined Transaction
func (r *DeclinedTransactionService) Get(ctx context.Context, declined_transaction_id string, opts ...option.RequestOption) (res *responses.DeclinedTransaction, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("declined_transactions/%s", declined_transaction_id)
	err = option.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Declined Transactions
func (r *DeclinedTransactionService) List(ctx context.Context, query requests.DeclinedTransactionListParams, opts ...option.RequestOption) (res *responses.Page[responses.DeclinedTransaction], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "declined_transactions"
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

// List Declined Transactions
func (r *DeclinedTransactionService) ListAutoPaging(ctx context.Context, query requests.DeclinedTransactionListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.DeclinedTransaction] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}

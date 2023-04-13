package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type PendingTransactionService struct {
	Options []option.RequestOption
}

func NewPendingTransactionService(opts ...option.RequestOption) (r *PendingTransactionService) {
	r = &PendingTransactionService{}
	r.Options = opts
	return
}

// Retrieve a Pending Transaction
func (r *PendingTransactionService) Get(ctx context.Context, pending_transaction_id string, opts ...option.RequestOption) (res *responses.PendingTransaction, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("pending_transactions/%s", pending_transaction_id)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// List Pending Transactions
func (r *PendingTransactionService) List(ctx context.Context, query *requests.PendingTransactionListParams, opts ...option.RequestOption) (res *responses.Page[responses.PendingTransaction], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "pending_transactions"
	cfg, err := option.NewRequestConfig(ctx, "GET", path, query, &res, opts...)
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

// List Pending Transactions
func (r *PendingTransactionService) ListAutoPager(ctx context.Context, query *requests.PendingTransactionListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.PendingTransaction] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}

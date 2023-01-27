package services

import (
	"context"
	"fmt"
	"increase/core"
	"increase/pagination"
	"increase/types"
)

type TransactionService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewTransactionService(requester core.Requester) (r *TransactionService) {
	r = &TransactionService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

// Retrieve a Transaction
func (r *TransactionService) Get(ctx context.Context, transaction_id string, opts ...*core.RequestOpts) (res *types.Transaction, err error) {
	path := fmt.Sprintf("/transactions/%s", transaction_id)
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
	}
	err = r.get(ctx, path, req, &res)

	return
}

// List Transactions
func (r *TransactionService) List(ctx context.Context, query *types.ListTransactionsQuery, opts ...*core.RequestOpts) (res *types.TransactionsPage, err error) {
	page := &types.TransactionsPage{
		Page: &pagination.Page[types.Transaction]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/transactions",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

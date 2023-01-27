package services

import (
	"context"
	"fmt"
	"increase/core"
	"increase/pagination"
	"increase/types"
)

type DeclinedTransactionService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewDeclinedTransactionService(requester core.Requester) (r *DeclinedTransactionService) {
	r = &DeclinedTransactionService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

// Retrieve a Declined Transaction
func (r *DeclinedTransactionService) Get(ctx context.Context, declined_transaction_id string, opts ...*core.RequestOpts) (res *types.DeclinedTransaction, err error) {
	path := fmt.Sprintf("/declined_transactions/%s", declined_transaction_id)
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
	}
	err = r.get(ctx, path, req, &res)

	return
}

// List Declined Transactions
func (r *DeclinedTransactionService) List(ctx context.Context, query *types.ListDeclinedTransactionsQuery, opts ...*core.RequestOpts) (res *types.DeclinedTransactionsPage, err error) {
	page := &types.DeclinedTransactionsPage{
		Page: &pagination.Page[types.DeclinedTransaction]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/declined_transactions",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

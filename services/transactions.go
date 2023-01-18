package services

import "context"
import "increase/core"
import "increase/types"
import "fmt"
import "increase/pagination"

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

func (r *TransactionService) Retrieve(ctx context.Context, transaction_id string, opts ...*core.RequestOpts) (res *types.Transaction, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/transactions/%s", transaction_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)

	return
}

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

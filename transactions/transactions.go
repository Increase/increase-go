package transactions

import "increase/core"
import "fmt"

type Transactions struct {
	Requester core.Requester
	get       func(string, *core.CoreRequest, interface{}) error
	post      func(string, *core.CoreRequest, interface{}) error
	patch     func(string, *core.CoreRequest, interface{}) error
	put       func(string, *core.CoreRequest, interface{}) error
	delete    func(string, *core.CoreRequest, interface{}) error
}

func NewTransactions(requster core.Requester) (r *Transactions) {
	r = &Transactions{}
	r.Requester = requster
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

func (r *Transactions) Retrieve(transaction_id string, opts ...*core.RequestOpts) (res Transaction, err error) {
	err = r.get(
		fmt.Sprintf("/transactions/%s", transaction_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *Transactions) List(query *TransactionsListQuery, opts ...*core.RequestOpts) (res TransactionsListResponse, err error) {
	err = r.get(
		"/transactions",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Query:  query,
		},
		&res,
	)
	return
}

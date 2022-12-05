package pending_transactions

import "increase/core"
import "fmt"

type PendingTransactions struct {
	Requester core.Requester
	get       func(string, *core.CoreRequest, interface{}) error
	post      func(string, *core.CoreRequest, interface{}) error
	patch     func(string, *core.CoreRequest, interface{}) error
	put       func(string, *core.CoreRequest, interface{}) error
	delete    func(string, *core.CoreRequest, interface{}) error
}

func NewPendingTransactions(requster core.Requester) (r *PendingTransactions) {
	r = &PendingTransactions{}
	r.Requester = requster
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

func (r *PendingTransactions) Retrieve(pending_transaction_id string, opts ...*core.RequestOpts) (res PendingTransaction, err error) {
	err = r.get(
		fmt.Sprintf("/pending_transactions/%s", pending_transaction_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *PendingTransactions) List(query *PendingTransactionsListQuery, opts ...*core.RequestOpts) (res PendingTransactionsListResponse, err error) {
	err = r.get(
		"/pending_transactions",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Query:  query,
		},
		&res,
	)
	return
}

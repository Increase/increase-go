package declined_transactions

import "increase/core"
import "fmt"

type DeclinedTransactions struct {
	Requester core.Requester
	get       func(string, *core.CoreRequest, interface{}) error
	post      func(string, *core.CoreRequest, interface{}) error
	patch     func(string, *core.CoreRequest, interface{}) error
	put       func(string, *core.CoreRequest, interface{}) error
	delete    func(string, *core.CoreRequest, interface{}) error
}

func NewDeclinedTransactions(requster core.Requester) (r *DeclinedTransactions) {
	r = &DeclinedTransactions{}
	r.Requester = requster
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

func (r *DeclinedTransactions) Retrieve(declined_transaction_id string, opts ...*core.RequestOpts) (res DeclinedTransaction, err error) {
	err = r.get(
		fmt.Sprintf("/declined_transactions/%s", declined_transaction_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *DeclinedTransactions) List(query *DeclinedTransactionsListQuery, opts ...*core.RequestOpts) (res DeclinedTransactionsListResponse, err error) {
	err = r.get(
		"/declined_transactions",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Query:  query,
		},
		&res,
	)
	return
}

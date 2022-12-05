package account_statements

import "increase/core"
import "fmt"

type AccountStatements struct {
	Requester core.Requester
	get       func(string, *core.CoreRequest, interface{}) error
	post      func(string, *core.CoreRequest, interface{}) error
	patch     func(string, *core.CoreRequest, interface{}) error
	put       func(string, *core.CoreRequest, interface{}) error
	delete    func(string, *core.CoreRequest, interface{}) error
}

func NewAccountStatements(requster core.Requester) (r *AccountStatements) {
	r = &AccountStatements{}
	r.Requester = requster
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

func (r *AccountStatements) Retrieve(account_statement_id string, opts ...*core.RequestOpts) (res AccountStatement, err error) {
	err = r.get(
		fmt.Sprintf("/account_statements/%s", account_statement_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *AccountStatements) List(query *AccountStatementsListQuery, opts ...*core.RequestOpts) (res AccountStatementsListResponse, err error) {
	err = r.get(
		"/account_statements",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Query:  query,
		},
		&res,
	)
	return
}

package simulations

import "increase/core"
import "increase/account_statements"

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

// Simulates an Account Statement being created for an account.
func (r *AccountStatements) Create(body AccountStatementsCreateParams, opts ...*core.RequestOpts) (res account_statements.AccountStatement, err error) {
	err = r.post(
		"/simulations/account_statements",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

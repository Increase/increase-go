package accounts

import "increase/core"
import "fmt"

type Accounts struct {
	Requester core.Requester
	get       func(string, *core.CoreRequest, interface{}) error
	post      func(string, *core.CoreRequest, interface{}) error
	patch     func(string, *core.CoreRequest, interface{}) error
	put       func(string, *core.CoreRequest, interface{}) error
	delete    func(string, *core.CoreRequest, interface{}) error
}

func NewAccounts(requster core.Requester) (r *Accounts) {
	r = &Accounts{}
	r.Requester = requster
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

func (r *Accounts) Create(body AccountsCreateParams, opts ...*core.RequestOpts) (res Account, err error) {
	err = r.post(
		"/accounts",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *Accounts) Retrieve(account_id string, opts ...*core.RequestOpts) (res Account, err error) {
	err = r.get(
		fmt.Sprintf("/accounts/%s", account_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *Accounts) Update(account_id string, body AccountsUpdateParams, opts ...*core.RequestOpts) (res Account, err error) {
	err = r.patch(
		fmt.Sprintf("/accounts/%s", account_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *Accounts) List(query *AccountsListQuery, opts ...*core.RequestOpts) (res AccountsListResponse, err error) {
	err = r.get(
		"/accounts",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Query:  query,
		},
		&res,
	)
	return
}

func (r *Accounts) Close(account_id string, opts ...*core.RequestOpts) (res Account, err error) {
	err = r.post(
		fmt.Sprintf("/accounts/%s/close", account_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

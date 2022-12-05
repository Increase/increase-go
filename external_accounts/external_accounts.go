package external_accounts

import "increase/core"
import "fmt"

type ExternalAccounts struct {
	Requester core.Requester
	get       func(string, *core.CoreRequest, interface{}) error
	post      func(string, *core.CoreRequest, interface{}) error
	patch     func(string, *core.CoreRequest, interface{}) error
	put       func(string, *core.CoreRequest, interface{}) error
	delete    func(string, *core.CoreRequest, interface{}) error
}

func NewExternalAccounts(requster core.Requester) (r *ExternalAccounts) {
	r = &ExternalAccounts{}
	r.Requester = requster
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

func (r *ExternalAccounts) Create(body ExternalAccountsCreateParams, opts ...*core.RequestOpts) (res ExternalAccount, err error) {
	err = r.post(
		"/external_accounts",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *ExternalAccounts) Retrieve(external_account_id string, opts ...*core.RequestOpts) (res ExternalAccount, err error) {
	err = r.get(
		fmt.Sprintf("/external_accounts/%s", external_account_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *ExternalAccounts) Update(external_account_id string, body ExternalAccountsUpdateParams, opts ...*core.RequestOpts) (res ExternalAccount, err error) {
	err = r.patch(
		fmt.Sprintf("/external_accounts/%s", external_account_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *ExternalAccounts) List(query *ExternalAccountsListQuery, opts ...*core.RequestOpts) (res ExternalAccountsListResponse, err error) {
	err = r.get(
		"/external_accounts",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Query:  query,
		},
		&res,
	)
	return
}

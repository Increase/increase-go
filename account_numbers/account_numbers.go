package account_numbers

import "increase/core"
import "fmt"

type AccountNumbers struct {
	Requester core.Requester
	get       func(string, *core.CoreRequest, interface{}) error
	post      func(string, *core.CoreRequest, interface{}) error
	patch     func(string, *core.CoreRequest, interface{}) error
	put       func(string, *core.CoreRequest, interface{}) error
	delete    func(string, *core.CoreRequest, interface{}) error
}

func NewAccountNumbers(requster core.Requester) (r *AccountNumbers) {
	r = &AccountNumbers{}
	r.Requester = requster
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

func (r *AccountNumbers) Create(body AccountNumbersCreateParams, opts ...*core.RequestOpts) (res AccountNumber, err error) {
	err = r.post(
		"/account_numbers",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *AccountNumbers) Retrieve(account_number_id string, opts ...*core.RequestOpts) (res AccountNumber, err error) {
	err = r.get(
		fmt.Sprintf("/account_numbers/%s", account_number_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *AccountNumbers) Update(account_number_id string, body AccountNumbersUpdateParams, opts ...*core.RequestOpts) (res AccountNumber, err error) {
	err = r.patch(
		fmt.Sprintf("/account_numbers/%s", account_number_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *AccountNumbers) List(query *AccountNumbersListQuery, opts ...*core.RequestOpts) (res AccountNumbersListResponse, err error) {
	err = r.get(
		"/account_numbers",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Query:  query,
		},
		&res,
	)
	return
}

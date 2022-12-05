package card_disputes

import "increase/core"
import "fmt"

type CardDisputes struct {
	Requester core.Requester
	get       func(string, *core.CoreRequest, interface{}) error
	post      func(string, *core.CoreRequest, interface{}) error
	patch     func(string, *core.CoreRequest, interface{}) error
	put       func(string, *core.CoreRequest, interface{}) error
	delete    func(string, *core.CoreRequest, interface{}) error
}

func NewCardDisputes(requster core.Requester) (r *CardDisputes) {
	r = &CardDisputes{}
	r.Requester = requster
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

func (r *CardDisputes) Create(body CardDisputesCreateParams, opts ...*core.RequestOpts) (res CardDispute, err error) {
	err = r.post(
		"/card_disputes",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *CardDisputes) Retrieve(card_dispute_id string, opts ...*core.RequestOpts) (res CardDispute, err error) {
	err = r.get(
		fmt.Sprintf("/card_disputes/%s", card_dispute_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *CardDisputes) List(query *CardDisputesListQuery, opts ...*core.RequestOpts) (res CardDisputesListResponse, err error) {
	err = r.get(
		"/card_disputes",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Query:  query,
		},
		&res,
	)
	return
}

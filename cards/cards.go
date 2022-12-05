package cards

import "increase/core"
import "fmt"

type Cards struct {
	Requester core.Requester
	get       func(string, *core.CoreRequest, interface{}) error
	post      func(string, *core.CoreRequest, interface{}) error
	patch     func(string, *core.CoreRequest, interface{}) error
	put       func(string, *core.CoreRequest, interface{}) error
	delete    func(string, *core.CoreRequest, interface{}) error
}

func NewCards(requster core.Requester) (r *Cards) {
	r = &Cards{}
	r.Requester = requster
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

func (r *Cards) Create(body CardsCreateParams, opts ...*core.RequestOpts) (res Card, err error) {
	err = r.post(
		"/cards",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *Cards) Retrieve(card_id string, opts ...*core.RequestOpts) (res Card, err error) {
	err = r.get(
		fmt.Sprintf("/cards/%s", card_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *Cards) Update(card_id string, body CardsUpdateParams, opts ...*core.RequestOpts) (res Card, err error) {
	err = r.patch(
		fmt.Sprintf("/cards/%s", card_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *Cards) List(query *CardsListQuery, opts ...*core.RequestOpts) (res CardsListResponse, err error) {
	err = r.get(
		"/cards",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Query:  query,
		},
		&res,
	)
	return
}

func (r *Cards) RetrieveSensitiveDetails(card_id string, opts ...*core.RequestOpts) (res CardDetails, err error) {
	err = r.get(
		fmt.Sprintf("/cards/%s/details", card_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

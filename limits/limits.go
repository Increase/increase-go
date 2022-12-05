package limits

import "increase/core"
import "fmt"

type Limits struct {
	Requester core.Requester
	get       func(string, *core.CoreRequest, interface{}) error
	post      func(string, *core.CoreRequest, interface{}) error
	patch     func(string, *core.CoreRequest, interface{}) error
	put       func(string, *core.CoreRequest, interface{}) error
	delete    func(string, *core.CoreRequest, interface{}) error
}

func NewLimits(requster core.Requester) (r *Limits) {
	r = &Limits{}
	r.Requester = requster
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

func (r *Limits) Create(body LimitsCreateParams, opts ...*core.RequestOpts) (res Limit, err error) {
	err = r.post(
		"/limits",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *Limits) Retrieve(limit_id string, opts ...*core.RequestOpts) (res Limit, err error) {
	err = r.get(
		fmt.Sprintf("/limits/%s", limit_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *Limits) Update(limit_id string, body LimitsUpdateParams, opts ...*core.RequestOpts) (res Limit, err error) {
	err = r.patch(
		fmt.Sprintf("/limits/%s", limit_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *Limits) List(query *LimitsListQuery, opts ...*core.RequestOpts) (res LimitsListResponse, err error) {
	err = r.get(
		"/limits",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Query:  query,
		},
		&res,
	)
	return
}

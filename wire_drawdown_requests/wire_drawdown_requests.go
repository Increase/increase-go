package wire_drawdown_requests

import "increase/core"
import "fmt"

type WireDrawdownRequests struct {
	Requester core.Requester
	get       func(string, *core.CoreRequest, interface{}) error
	post      func(string, *core.CoreRequest, interface{}) error
	patch     func(string, *core.CoreRequest, interface{}) error
	put       func(string, *core.CoreRequest, interface{}) error
	delete    func(string, *core.CoreRequest, interface{}) error
}

func NewWireDrawdownRequests(requster core.Requester) (r *WireDrawdownRequests) {
	r = &WireDrawdownRequests{}
	r.Requester = requster
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

func (r *WireDrawdownRequests) Create(body WireDrawdownRequestsCreateParams, opts ...*core.RequestOpts) (res WireDrawdownRequest, err error) {
	err = r.post(
		"/wire_drawdown_requests",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *WireDrawdownRequests) Retrieve(wire_drawdown_request_id string, opts ...*core.RequestOpts) (res WireDrawdownRequest, err error) {
	err = r.get(
		fmt.Sprintf("/wire_drawdown_requests/%s", wire_drawdown_request_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *WireDrawdownRequests) List(query *WireDrawdownRequestsListQuery, opts ...*core.RequestOpts) (res WireDrawdownRequestsListResponse, err error) {
	err = r.get(
		"/wire_drawdown_requests",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Query:  query,
		},
		&res,
	)
	return
}

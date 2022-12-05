package events

import "increase/core"
import "fmt"

type Events struct {
	Requester core.Requester
	get       func(string, *core.CoreRequest, interface{}) error
	post      func(string, *core.CoreRequest, interface{}) error
	patch     func(string, *core.CoreRequest, interface{}) error
	put       func(string, *core.CoreRequest, interface{}) error
	delete    func(string, *core.CoreRequest, interface{}) error
}

func NewEvents(requster core.Requester) (r *Events) {
	r = &Events{}
	r.Requester = requster
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

func (r *Events) Retrieve(event_id string, opts ...*core.RequestOpts) (res Event, err error) {
	err = r.get(
		fmt.Sprintf("/events/%s", event_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *Events) List(query *EventsListQuery, opts ...*core.RequestOpts) (res EventsListResponse, err error) {
	err = r.get(
		"/events",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Query:  query,
		},
		&res,
	)
	return
}

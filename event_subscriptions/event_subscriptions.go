package event_subscriptions

import "increase/core"
import "fmt"

type EventSubscriptions struct {
	Requester core.Requester
	get       func(string, *core.CoreRequest, interface{}) error
	post      func(string, *core.CoreRequest, interface{}) error
	patch     func(string, *core.CoreRequest, interface{}) error
	put       func(string, *core.CoreRequest, interface{}) error
	delete    func(string, *core.CoreRequest, interface{}) error
}

func NewEventSubscriptions(requster core.Requester) (r *EventSubscriptions) {
	r = &EventSubscriptions{}
	r.Requester = requster
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

func (r *EventSubscriptions) Create(body EventSubscriptionsCreateParams, opts ...*core.RequestOpts) (res EventSubscription, err error) {
	err = r.post(
		"/event_subscriptions",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *EventSubscriptions) Retrieve(event_subscription_id string, opts ...*core.RequestOpts) (res EventSubscription, err error) {
	err = r.get(
		fmt.Sprintf("/event_subscriptions/%s", event_subscription_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *EventSubscriptions) Update(event_subscription_id string, body EventSubscriptionsUpdateParams, opts ...*core.RequestOpts) (res EventSubscription, err error) {
	err = r.patch(
		fmt.Sprintf("/event_subscriptions/%s", event_subscription_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *EventSubscriptions) List(query *EventSubscriptionsListQuery, opts ...*core.RequestOpts) (res EventSubscriptionsListResponse, err error) {
	err = r.get(
		"/event_subscriptions",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Query:  query,
		},
		&res,
	)
	return
}

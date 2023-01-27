package services

import (
	"context"
	"fmt"
	"increase/core"
	"increase/pagination"
	"increase/types"
)

type EventSubscriptionService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewEventSubscriptionService(requester core.Requester) (r *EventSubscriptionService) {
	r = &EventSubscriptionService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

// Create an Event Subscription
func (r *EventSubscriptionService) New(ctx context.Context, body *types.CreateAnEventSubscriptionParameters, opts ...*core.RequestOpts) (res *types.EventSubscription, err error) {
	path := "/event_subscriptions"
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
		Body:   body,
	}
	err = r.post(ctx, path, req, &res)

	return
}

// Retrieve an Event Subscription
func (r *EventSubscriptionService) Get(ctx context.Context, event_subscription_id string, opts ...*core.RequestOpts) (res *types.EventSubscription, err error) {
	path := fmt.Sprintf("/event_subscriptions/%s", event_subscription_id)
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
	}
	err = r.get(ctx, path, req, &res)

	return
}

// Update an Event Subscription
func (r *EventSubscriptionService) Update(ctx context.Context, event_subscription_id string, body *types.UpdateAnEventSubscriptionParameters, opts ...*core.RequestOpts) (res *types.EventSubscription, err error) {
	path := fmt.Sprintf("/event_subscriptions/%s", event_subscription_id)
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
		Body:   body,
	}
	err = r.patch(ctx, path, req, &res)

	return
}

// List Event Subscriptions
func (r *EventSubscriptionService) List(ctx context.Context, query *types.ListEventSubscriptionsQuery, opts ...*core.RequestOpts) (res *types.EventSubscriptionsPage, err error) {
	page := &types.EventSubscriptionsPage{
		Page: &pagination.Page[types.EventSubscription]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/event_subscriptions",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

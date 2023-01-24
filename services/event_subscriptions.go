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

func (r *EventSubscriptionService) Create(ctx context.Context, body *types.CreateAnEventSubscriptionParameters, opts ...*core.RequestOpts) (res *types.EventSubscription, err error) {
	err = r.post(
		ctx,
		"/event_subscriptions",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)

	return
}

func (r *EventSubscriptionService) Retrieve(ctx context.Context, event_subscription_id string, opts ...*core.RequestOpts) (res *types.EventSubscription, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/event_subscriptions/%s", event_subscription_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)

	return
}

func (r *EventSubscriptionService) Update(ctx context.Context, event_subscription_id string, body *types.UpdateAnEventSubscriptionParameters, opts ...*core.RequestOpts) (res *types.EventSubscription, err error) {
	err = r.patch(
		ctx,
		fmt.Sprintf("/event_subscriptions/%s", event_subscription_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)

	return
}

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

package services

import (
	"context"
	"fmt"
	"increase/options"
	"increase/pagination"
	"increase/types"
)

type EventSubscriptionService struct {
	Options []options.RequestOption
}

func NewEventSubscriptionService(opts ...options.RequestOption) (r *EventSubscriptionService) {
	r = &EventSubscriptionService{}
	r.Options = opts
	return
}

// Create an Event Subscription
func (r *EventSubscriptionService) New(ctx context.Context, body *types.CreateAnEventSubscriptionParameters, opts ...options.RequestOption) (res *types.EventSubscription, err error) {
	opts = append(r.Options[:], opts...)
	path := "event_subscriptions"
	err = options.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Retrieve an Event Subscription
func (r *EventSubscriptionService) Get(ctx context.Context, event_subscription_id string, opts ...options.RequestOption) (res *types.EventSubscription, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("event_subscriptions/%s", event_subscription_id)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// Update an Event Subscription
func (r *EventSubscriptionService) Update(ctx context.Context, event_subscription_id string, body *types.UpdateAnEventSubscriptionParameters, opts ...options.RequestOption) (res *types.EventSubscription, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("event_subscriptions/%s", event_subscription_id)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// List Event Subscriptions
func (r *EventSubscriptionService) List(ctx context.Context, query *types.EventSubscriptionListParams, opts ...options.RequestOption) (res *types.EventSubscriptionsPage, err error) {
	opts = append(r.Options, opts...)
	path := "event_subscriptions"
	cfg, err := options.NewRequestConfig(ctx, "GET", path, query, nil, opts...)
	if err != nil {
		return
	}
	res = &types.EventSubscriptionsPage{
		Page: &pagination.Page[types.EventSubscription]{
			Config:  *cfg,
			Options: opts,
		},
	}
	return res, res.Fire()
}

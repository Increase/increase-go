package services

import (
	"context"
	"fmt"
	"increase/options"
	"increase/pagination"
	"increase/types"
	"net/url"
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
	u, err := url.Parse(fmt.Sprintf("event_subscriptions"))
	if err != nil {
		return
	}
	cfg, err := options.NewRequestConfig(ctx, "POST", u, body, opts...)
	if err != nil {
		return
	}
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return
	}

	return
}

// Retrieve an Event Subscription
func (r *EventSubscriptionService) Get(ctx context.Context, event_subscription_id string, opts ...options.RequestOption) (res *types.EventSubscription, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]options.RequestOption{options.WithHeader("Content-Type", "")}, opts...)
	u, err := url.Parse(fmt.Sprintf("event_subscriptions/%s", event_subscription_id))
	if err != nil {
		return
	}
	cfg, err := options.NewRequestConfig(ctx, "GET", u, nil, opts...)
	if err != nil {
		return
	}
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return
	}

	return
}

// Update an Event Subscription
func (r *EventSubscriptionService) Update(ctx context.Context, event_subscription_id string, body *types.UpdateAnEventSubscriptionParameters, opts ...options.RequestOption) (res *types.EventSubscription, err error) {
	opts = append(r.Options[:], opts...)
	u, err := url.Parse(fmt.Sprintf("event_subscriptions/%s", event_subscription_id))
	if err != nil {
		return
	}
	cfg, err := options.NewRequestConfig(ctx, "PATCH", u, body, opts...)
	if err != nil {
		return
	}
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return
	}

	return
}

// List Event Subscriptions
func (r *EventSubscriptionService) List(ctx context.Context, query *types.EventSubscriptionListParams, opts ...options.RequestOption) (res *types.EventSubscriptionsPage, err error) {
	u, err := url.Parse(fmt.Sprintf("event_subscriptions"))
	if err != nil {
		return
	}
	opts = append(r.Options, opts...)
	cfg, err := options.NewRequestConfig(ctx, "GET", u, query, opts...)
	if err != nil {
		return
	}
	res = &types.EventSubscriptionsPage{
		Page: &pagination.Page[types.EventSubscription]{
			Config:  *cfg,
			Options: opts,
		},
	}
	err = res.Fire()
	return
}

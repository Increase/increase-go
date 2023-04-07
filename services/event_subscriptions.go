package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type EventSubscriptionService struct {
	Options []option.RequestOption
}

func NewEventSubscriptionService(opts ...option.RequestOption) (r *EventSubscriptionService) {
	r = &EventSubscriptionService{}
	r.Options = opts
	return
}

// Create an Event Subscription
func (r *EventSubscriptionService) New(ctx context.Context, body *requests.EventSubscriptionNewParams, opts ...option.RequestOption) (res *responses.EventSubscription, err error) {
	opts = append(r.Options[:], opts...)
	path := "event_subscriptions"
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Retrieve an Event Subscription
func (r *EventSubscriptionService) Get(ctx context.Context, event_subscription_id string, opts ...option.RequestOption) (res *responses.EventSubscription, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("event_subscriptions/%s", event_subscription_id)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// Update an Event Subscription
func (r *EventSubscriptionService) Update(ctx context.Context, event_subscription_id string, body *requests.EventSubscriptionUpdateParams, opts ...option.RequestOption) (res *responses.EventSubscription, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("event_subscriptions/%s", event_subscription_id)
	err = option.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// List Event Subscriptions
func (r *EventSubscriptionService) List(ctx context.Context, query *requests.EventSubscriptionListParams, opts ...option.RequestOption) (res *responses.Page[responses.EventSubscription], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "event_subscriptions"
	cfg, err := option.NewRequestConfig(ctx, "GET", path, query, &res, opts...)
	if err != nil {
		return
	}
	err = cfg.Execute()
	if err != nil {
		return
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List Event Subscriptions
func (r *EventSubscriptionService) ListAutoPager(ctx context.Context, query *requests.EventSubscriptionListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.EventSubscription] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}

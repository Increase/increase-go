package services

import (
	"context"
	"fmt"
	"increase/options"
	"increase/pagination"
	"increase/types"
	"net/url"
)

type EventService struct {
	Options []options.RequestOption
}

func NewEventService(opts ...options.RequestOption) (r *EventService) {
	r = &EventService{}
	r.Options = opts
	return
}

// Retrieve an Event
func (r *EventService) Get(ctx context.Context, event_id string, opts ...options.RequestOption) (res *types.Event, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]options.RequestOption{options.WithHeader("Content-Type", "")}, opts...)
	u, err := url.Parse(fmt.Sprintf("events/%s", event_id))
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

// List Events
func (r *EventService) List(ctx context.Context, query *types.EventListParams, opts ...options.RequestOption) (res *types.EventsPage, err error) {
	u, err := url.Parse(fmt.Sprintf("events"))
	if err != nil {
		return
	}
	opts = append(r.Options, opts...)
	cfg, err := options.NewRequestConfig(ctx, "GET", u, query, opts...)
	if err != nil {
		return
	}
	res = &types.EventsPage{
		Page: &pagination.Page[types.Event]{
			Config:  *cfg,
			Options: opts,
		},
	}
	err = res.Fire()
	return
}

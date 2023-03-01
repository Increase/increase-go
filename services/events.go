package services

import (
	"context"
	"fmt"
	"increase/options"
	"increase/pagination"
	"increase/types"
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
	path := fmt.Sprintf("events/%s", event_id)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// List Events
func (r *EventService) List(ctx context.Context, query *types.EventListParams, opts ...options.RequestOption) (res *types.EventsPage, err error) {
	opts = append(r.Options, opts...)
	path := "events"
	cfg, err := options.NewRequestConfig(ctx, "GET", path, query, nil, opts...)
	if err != nil {
		return
	}
	res = &types.EventsPage{
		Page: &pagination.Page[types.Event]{
			Config:  *cfg,
			Options: opts,
		},
	}
	return res, res.Fire()
}

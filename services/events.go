package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type EventService struct {
	Options []option.RequestOption
}

func NewEventService(opts ...option.RequestOption) (r *EventService) {
	r = &EventService{}
	r.Options = opts
	return
}

// Retrieve an Event
func (r *EventService) Get(ctx context.Context, event_id string, opts ...option.RequestOption) (res *responses.Event, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("events/%s", event_id)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// List Events
func (r *EventService) List(ctx context.Context, query *requests.EventListParams, opts ...option.RequestOption) (res *responses.Page[responses.Event], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "events"
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

// List Events
func (r *EventService) ListAutoPager(ctx context.Context, query *requests.EventListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.Event] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}

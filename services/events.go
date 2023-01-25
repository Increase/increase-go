package services

import (
	"context"
	"fmt"
	"increase/core"
	"increase/pagination"
	"increase/types"
)

type EventService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewEventService(requester core.Requester) (r *EventService) {
	r = &EventService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

// Retrieve an Event
func (r *EventService) Retrieve(ctx context.Context, event_id string, opts ...*core.RequestOpts) (res *types.Event, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/events/%s", event_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)

	return
}

// List Events
func (r *EventService) List(ctx context.Context, query *types.ListEventsQuery, opts ...*core.RequestOpts) (res *types.EventsPage, err error) {
	page := &types.EventsPage{
		Page: &pagination.Page[types.Event]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/events",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

package services

import (
	"context"
	"fmt"
	"increase/core"
	"increase/pagination"
	"increase/types"
)

type ACHPrenotificationService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewACHPrenotificationService(requester core.Requester) (r *ACHPrenotificationService) {
	r = &ACHPrenotificationService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

// Create an ACH Prenotification
func (r *ACHPrenotificationService) New(ctx context.Context, body *types.CreateAnACHPrenotificationParameters, opts ...*core.RequestOpts) (res *types.ACHPrenotification, err error) {
	path := "/ach_prenotifications"
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
		Body:   body,
	}
	err = r.post(ctx, path, req, &res)

	return
}

// Retrieve an ACH Prenotification
func (r *ACHPrenotificationService) Get(ctx context.Context, ach_prenotification_id string, opts ...*core.RequestOpts) (res *types.ACHPrenotification, err error) {
	path := fmt.Sprintf("/ach_prenotifications/%s", ach_prenotification_id)
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
	}
	err = r.get(ctx, path, req, &res)

	return
}

// List ACH Prenotifications
func (r *ACHPrenotificationService) List(ctx context.Context, query *types.ACHPrenotificationListParams, opts ...*core.RequestOpts) (res *types.ACHPrenotificationsPage, err error) {
	page := &types.ACHPrenotificationsPage{
		Page: &pagination.Page[types.ACHPrenotification]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/ach_prenotifications",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

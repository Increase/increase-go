package services

import "context"
import "increase/core"
import "increase/types"
import "fmt"
import "increase/pagination"

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

func (r *ACHPrenotificationService) Create(ctx context.Context, body *types.CreateAnACHPrenotificationParameters, opts ...*core.RequestOpts) (res *types.ACHPrenotification, err error) {
	err = r.post(
		ctx,
		"/ach_prenotifications",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)

	return
}

func (r *ACHPrenotificationService) Retrieve(ctx context.Context, ach_prenotification_id string, opts ...*core.RequestOpts) (res *types.ACHPrenotification, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/ach_prenotifications/%s", ach_prenotification_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)

	return
}

func (r *ACHPrenotificationService) List(ctx context.Context, query *types.ListACHPrenotificationsQuery, opts ...*core.RequestOpts) (res *types.ACHPrenotificationsPage, err error) {
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

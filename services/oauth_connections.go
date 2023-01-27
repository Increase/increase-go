package services

import (
	"context"
	"fmt"
	"increase/core"
	"increase/pagination"
	"increase/types"
)

type OauthConnectionService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewOauthConnectionService(requester core.Requester) (r *OauthConnectionService) {
	r = &OauthConnectionService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

// Retrieve an OAuth Connection
func (r *OauthConnectionService) Get(ctx context.Context, oauth_connection_id string, opts ...*core.RequestOpts) (res *types.OauthConnection, err error) {
	path := fmt.Sprintf("/oauth_connections/%s", oauth_connection_id)
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
	}
	err = r.get(ctx, path, req, &res)

	return
}

// List OAuth Connections
func (r *OauthConnectionService) List(ctx context.Context, query *types.ListOauthConnectionsQuery, opts ...*core.RequestOpts) (res *types.OauthConnectionsPage, err error) {
	page := &types.OauthConnectionsPage{
		Page: &pagination.Page[types.OauthConnection]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/oauth_connections",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

package services

import (
	"context"
	"fmt"
	"increase/options"
	"increase/pagination"
	"increase/types"
	"net/url"
)

type OauthConnectionService struct {
	Options []options.RequestOption
}

func NewOauthConnectionService(opts ...options.RequestOption) (r *OauthConnectionService) {
	r = &OauthConnectionService{}
	r.Options = opts
	return
}

// Retrieve an OAuth Connection
func (r *OauthConnectionService) Get(ctx context.Context, oauth_connection_id string, opts ...options.RequestOption) (res *types.OauthConnection, err error) {
	opts = append(r.Options, opts...)
	u, err := url.Parse(fmt.Sprintf("oauth_connections/%s", oauth_connection_id))
	if err != nil {
		return
	}
	cfg := options.NewRequestConfig(ctx, "GET", u, opts...)
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return
	}

	return
}

// List OAuth Connections
func (r *OauthConnectionService) List(ctx context.Context, query *types.OauthConnectionListParams, opts ...options.RequestOption) (res *types.OauthConnectionsPage, err error) {
	u, err := url.Parse(fmt.Sprintf("oauth_connections"))
	if err != nil {
		return
	}
	opts = append(r.Options, opts...)
	cfg := options.NewRequestConfig(ctx, "GET", u, opts...)
	res = &types.OauthConnectionsPage{
		Page: &pagination.Page[types.OauthConnection]{
			Config:  *cfg,
			Options: opts,
		},
	}
	err = res.Fire()
	return
}

package services

import (
	"context"
	"fmt"

	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/pagination"
	"github.com/increase/increase-go/types"
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
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("oauth_connections/%s", oauth_connection_id)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// List OAuth Connections
func (r *OauthConnectionService) List(ctx context.Context, query *types.OauthConnectionListParams, opts ...options.RequestOption) (res *types.OauthConnectionsPage, err error) {
	opts = append(r.Options, opts...)
	path := "oauth_connections"
	cfg, err := options.NewRequestConfig(ctx, "GET", path, query, nil, opts...)
	if err != nil {
		return
	}
	res = &types.OauthConnectionsPage{
		Page: &pagination.Page[types.OauthConnection]{
			Config:  *cfg,
			Options: opts,
		},
	}
	return res, res.Fire()
}

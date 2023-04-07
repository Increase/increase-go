package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type OauthConnectionService struct {
	Options []option.RequestOption
}

func NewOauthConnectionService(opts ...option.RequestOption) (r *OauthConnectionService) {
	r = &OauthConnectionService{}
	r.Options = opts
	return
}

// Retrieve an OAuth Connection
func (r *OauthConnectionService) Get(ctx context.Context, oauth_connection_id string, opts ...option.RequestOption) (res *responses.OauthConnection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("oauth_connections/%s", oauth_connection_id)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// List OAuth Connections
func (r *OauthConnectionService) List(ctx context.Context, query *requests.OauthConnectionListParams, opts ...option.RequestOption) (res *responses.Page[responses.OauthConnection], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "oauth_connections"
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

// List OAuth Connections
func (r *OauthConnectionService) ListAutoPager(ctx context.Context, query *requests.OauthConnectionListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.OauthConnection] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}

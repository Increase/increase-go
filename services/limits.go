package services

import (
	"context"
	"fmt"
	"increase/options"
	"increase/pagination"
	"increase/types"
	"net/url"
)

type LimitService struct {
	Options []options.RequestOption
}

func NewLimitService(opts ...options.RequestOption) (r *LimitService) {
	r = &LimitService{}
	r.Options = opts
	return
}

// Create a Limit
func (r *LimitService) New(ctx context.Context, body *types.CreateALimitParameters, opts ...options.RequestOption) (res *types.Limit, err error) {
	opts = append(r.Options, opts...)
	u, err := url.Parse(fmt.Sprintf("limits"))
	if err != nil {
		return
	}
	cfg := options.NewRequestConfig(ctx, "POST", u, opts...)
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return
	}

	return
}

// Retrieve a Limit
func (r *LimitService) Get(ctx context.Context, limit_id string, opts ...options.RequestOption) (res *types.Limit, err error) {
	opts = append(r.Options, opts...)
	u, err := url.Parse(fmt.Sprintf("limits/%s", limit_id))
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

// Update a Limit
func (r *LimitService) Update(ctx context.Context, limit_id string, body *types.UpdateALimitParameters, opts ...options.RequestOption) (res *types.Limit, err error) {
	opts = append(r.Options, opts...)
	u, err := url.Parse(fmt.Sprintf("limits/%s", limit_id))
	if err != nil {
		return
	}
	cfg := options.NewRequestConfig(ctx, "PATCH", u, opts...)
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return
	}

	return
}

// List Limits
func (r *LimitService) List(ctx context.Context, query *types.LimitListParams, opts ...options.RequestOption) (res *types.LimitsPage, err error) {
	u, err := url.Parse(fmt.Sprintf("limits"))
	if err != nil {
		return
	}
	opts = append(r.Options, opts...)
	cfg := options.NewRequestConfig(ctx, "GET", u, opts...)
	res = &types.LimitsPage{
		Page: &pagination.Page[types.Limit]{
			Config:  *cfg,
			Options: opts,
		},
	}
	err = res.Fire()
	return
}

package services

import (
	"context"
	"fmt"
	"increase/options"
	"increase/pagination"
	"increase/types"
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
	opts = append(r.Options[:], opts...)
	path := "limits"
	err = options.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Retrieve a Limit
func (r *LimitService) Get(ctx context.Context, limit_id string, opts ...options.RequestOption) (res *types.Limit, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("limits/%s", limit_id)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// Update a Limit
func (r *LimitService) Update(ctx context.Context, limit_id string, body *types.UpdateALimitParameters, opts ...options.RequestOption) (res *types.Limit, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("limits/%s", limit_id)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// List Limits
func (r *LimitService) List(ctx context.Context, query *types.LimitListParams, opts ...options.RequestOption) (res *types.LimitsPage, err error) {
	opts = append(r.Options, opts...)
	path := "limits"
	cfg, err := options.NewRequestConfig(ctx, "GET", path, query, nil, opts...)
	if err != nil {
		return
	}
	res = &types.LimitsPage{
		Page: &pagination.Page[types.Limit]{
			Config:  *cfg,
			Options: opts,
		},
	}
	return res, res.Fire()
}

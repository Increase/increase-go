package services

import (
	"context"
	"fmt"
	"increase/options"
	"increase/pagination"
	"increase/types"
	"net/url"
)

type WireDrawdownRequestService struct {
	Options []options.RequestOption
}

func NewWireDrawdownRequestService(opts ...options.RequestOption) (r *WireDrawdownRequestService) {
	r = &WireDrawdownRequestService{}
	r.Options = opts
	return
}

// Create a Wire Drawdown Request
func (r *WireDrawdownRequestService) New(ctx context.Context, body *types.CreateAWireDrawdownRequestParameters, opts ...options.RequestOption) (res *types.WireDrawdownRequest, err error) {
	opts = append(r.Options, opts...)
	u, err := url.Parse(fmt.Sprintf("wire_drawdown_requests"))
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

// Retrieve a Wire Drawdown Request
func (r *WireDrawdownRequestService) Get(ctx context.Context, wire_drawdown_request_id string, opts ...options.RequestOption) (res *types.WireDrawdownRequest, err error) {
	opts = append(r.Options, opts...)
	u, err := url.Parse(fmt.Sprintf("wire_drawdown_requests/%s", wire_drawdown_request_id))
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

// List Wire Drawdown Requests
func (r *WireDrawdownRequestService) List(ctx context.Context, query *types.WireDrawdownRequestListParams, opts ...options.RequestOption) (res *types.WireDrawdownRequestsPage, err error) {
	u, err := url.Parse(fmt.Sprintf("wire_drawdown_requests"))
	if err != nil {
		return
	}
	opts = append(r.Options, opts...)
	cfg := options.NewRequestConfig(ctx, "GET", u, opts...)
	res = &types.WireDrawdownRequestsPage{
		Page: &pagination.Page[types.WireDrawdownRequest]{
			Config:  *cfg,
			Options: opts,
		},
	}
	err = res.Fire()
	return
}

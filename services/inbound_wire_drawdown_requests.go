package services

import (
	"context"
	"fmt"
	"increase/options"
	"increase/pagination"
	"increase/types"
	"net/url"
)

type InboundWireDrawdownRequestService struct {
	Options []options.RequestOption
}

func NewInboundWireDrawdownRequestService(opts ...options.RequestOption) (r *InboundWireDrawdownRequestService) {
	r = &InboundWireDrawdownRequestService{}
	r.Options = opts
	return
}

// Retrieve an Inbound Wire Drawdown Request
func (r *InboundWireDrawdownRequestService) Get(ctx context.Context, inbound_wire_drawdown_request_id string, opts ...options.RequestOption) (res *types.InboundWireDrawdownRequest, err error) {
	opts = append(r.Options[:], opts...)
	u, err := url.Parse(fmt.Sprintf("inbound_wire_drawdown_requests/%s", inbound_wire_drawdown_request_id))
	if err != nil {
		return
	}
	cfg, err := options.NewRequestConfig(ctx, "GET", u, nil, opts...)
	if err != nil {
		return
	}
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return
	}

	return
}

// List Inbound Wire Drawdown Requests
func (r *InboundWireDrawdownRequestService) List(ctx context.Context, query *types.InboundWireDrawdownRequestListParams, opts ...options.RequestOption) (res *types.InboundWireDrawdownRequestsPage, err error) {
	u, err := url.Parse(fmt.Sprintf("inbound_wire_drawdown_requests"))
	if err != nil {
		return
	}
	opts = append(r.Options, opts...)
	cfg, err := options.NewRequestConfig(ctx, "GET", u, nil, opts...)
	if err != nil {
		return
	}
	res = &types.InboundWireDrawdownRequestsPage{
		Page: &pagination.Page[types.InboundWireDrawdownRequest]{
			Config:  *cfg,
			Options: opts,
		},
	}
	err = res.Fire()
	return
}

package services

import (
	"context"
	"fmt"
	"increase/options"
	"increase/pagination"
	"increase/types"
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
	path := fmt.Sprintf("inbound_wire_drawdown_requests/%s", inbound_wire_drawdown_request_id)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// List Inbound Wire Drawdown Requests
func (r *InboundWireDrawdownRequestService) List(ctx context.Context, query *types.InboundWireDrawdownRequestListParams, opts ...options.RequestOption) (res *types.InboundWireDrawdownRequestsPage, err error) {
	opts = append(r.Options, opts...)
	path := "inbound_wire_drawdown_requests"
	cfg, err := options.NewRequestConfig(ctx, "GET", path, query, nil, opts...)
	if err != nil {
		return
	}
	res = &types.InboundWireDrawdownRequestsPage{
		Page: &pagination.Page[types.InboundWireDrawdownRequest]{
			Config:  *cfg,
			Options: opts,
		},
	}
	return res, res.Fire()
}

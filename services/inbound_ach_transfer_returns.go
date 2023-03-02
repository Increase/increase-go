package services

import (
	"context"
	"fmt"
	"increase/options"
	"increase/pagination"
	"increase/types"
)

type InboundACHTransferReturnService struct {
	Options []options.RequestOption
}

func NewInboundACHTransferReturnService(opts ...options.RequestOption) (r *InboundACHTransferReturnService) {
	r = &InboundACHTransferReturnService{}
	r.Options = opts
	return
}

// Create an ACH Return
func (r *InboundACHTransferReturnService) New(ctx context.Context, body *types.CreateAnACHReturnParameters, opts ...options.RequestOption) (res *types.InboundACHTransferReturn, err error) {
	opts = append(r.Options[:], opts...)
	path := "inbound_ach_transfer_returns"
	err = options.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Retrieve an Inbound ACH Transfer Return
func (r *InboundACHTransferReturnService) Get(ctx context.Context, inbound_ach_transfer_return_id string, opts ...options.RequestOption) (res *types.InboundACHTransferReturn, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("inbound_ach_transfer_returns/%s", inbound_ach_transfer_return_id)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// List Inbound ACH Transfer Returns
func (r *InboundACHTransferReturnService) List(ctx context.Context, query *types.InboundACHTransferReturnListParams, opts ...options.RequestOption) (res *types.InboundACHTransferReturnsPage, err error) {
	opts = append(r.Options, opts...)
	path := "inbound_ach_transfer_returns"
	cfg, err := options.NewRequestConfig(ctx, "GET", path, query, nil, opts...)
	if err != nil {
		return
	}
	res = &types.InboundACHTransferReturnsPage{
		Page: &pagination.Page[types.InboundACHTransferReturn]{
			Config:  *cfg,
			Options: opts,
		},
	}
	return res, res.Fire()
}

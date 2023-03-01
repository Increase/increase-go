package services

import (
	"context"
	"increase/options"
	"increase/types"
)

type SimulationsInboundWireDrawdownRequestService struct {
	Options []options.RequestOption
}

func NewSimulationsInboundWireDrawdownRequestService(opts ...options.RequestOption) (r *SimulationsInboundWireDrawdownRequestService) {
	r = &SimulationsInboundWireDrawdownRequestService{}
	r.Options = opts
	return
}

// Simulates the receival of an
// [Inbound Wire Drawdown Request](#inbound-wire-drawdown-requests).
func (r *SimulationsInboundWireDrawdownRequestService) New(ctx context.Context, body *types.SimulateAnInboundWireDrawdownRequestBeingCreatedParameters, opts ...options.RequestOption) (res *types.InboundWireDrawdownRequest, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/inbound_wire_drawdown_requests"
	err = options.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

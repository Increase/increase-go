package services

import (
	"context"
	"fmt"
	"increase/options"
	"increase/types"
	"net/url"
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
	u, err := url.Parse(fmt.Sprintf("simulations/inbound_wire_drawdown_requests"))
	if err != nil {
		return
	}
	cfg, err := options.NewRequestConfig(ctx, "POST", u, body, opts...)
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

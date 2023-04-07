package services

import (
	"context"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type SimulationsInboundWireDrawdownRequestService struct {
	Options []option.RequestOption
}

func NewSimulationsInboundWireDrawdownRequestService(opts ...option.RequestOption) (r *SimulationsInboundWireDrawdownRequestService) {
	r = &SimulationsInboundWireDrawdownRequestService{}
	r.Options = opts
	return
}

// Simulates the receival of an
// [Inbound Wire Drawdown Request](#inbound-wire-drawdown-requests).
func (r *SimulationsInboundWireDrawdownRequestService) New(ctx context.Context, body *requests.InboundWireDrawdownRequestNewParams, opts ...option.RequestOption) (res *responses.InboundWireDrawdownRequest, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/inbound_wire_drawdown_requests"
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

package services

import (
	"context"
	"net/http"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type SimulationInboundWireDrawdownRequestService struct {
	Options []option.RequestOption
}

func NewSimulationInboundWireDrawdownRequestService(opts ...option.RequestOption) (r *SimulationInboundWireDrawdownRequestService) {
	r = &SimulationInboundWireDrawdownRequestService{}
	r.Options = opts
	return
}

// Simulates the receival of an
// [Inbound Wire Drawdown Request](#inbound-wire-drawdown-requests).
func (r *SimulationInboundWireDrawdownRequestService) New(ctx context.Context, body *requests.SimulationInboundWireDrawdownRequestNewParams, opts ...option.RequestOption) (res *responses.InboundWireDrawdownRequest, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/inbound_wire_drawdown_requests"
	err = option.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

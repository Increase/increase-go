package services

import (
	"context"
	"net/http"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type SimulationWireTransferService struct {
	Options []option.RequestOption
}

func NewSimulationWireTransferService(opts ...option.RequestOption) (r *SimulationWireTransferService) {
	r = &SimulationWireTransferService{}
	r.Options = opts
	return
}

// Simulates an inbound Wire Transfer to your account.
func (r *SimulationWireTransferService) NewInbound(ctx context.Context, body requests.SimulationWireTransferNewInboundParams, opts ...option.RequestOption) (res *responses.WireTransferSimulation, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/inbound_wire_transfers"
	err = option.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

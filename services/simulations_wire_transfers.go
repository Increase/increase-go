package services

import (
	"context"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type SimulationsWireTransferService struct {
	Options []option.RequestOption
}

func NewSimulationsWireTransferService(opts ...option.RequestOption) (r *SimulationsWireTransferService) {
	r = &SimulationsWireTransferService{}
	r.Options = opts
	return
}

// Simulates an inbound Wire Transfer to your account.
func (r *SimulationsWireTransferService) NewInbound(ctx context.Context, body *requests.WireTransferNewInboundParams, opts ...option.RequestOption) (res *responses.WireTransferSimulation, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/inbound_wire_transfers"
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

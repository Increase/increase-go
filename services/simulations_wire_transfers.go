package services

import (
	"context"
	"increase/options"
	"increase/types"
)

type SimulationsWireTransferService struct {
	Options []options.RequestOption
}

func NewSimulationsWireTransferService(opts ...options.RequestOption) (r *SimulationsWireTransferService) {
	r = &SimulationsWireTransferService{}
	r.Options = opts
	return
}

// Simulates an inbound Wire Transfer to your account.
func (r *SimulationsWireTransferService) NewInbound(ctx context.Context, body *types.SimulateAWireTransferToYourAccountParameters, opts ...options.RequestOption) (res *types.WireTransferSimulation, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/inbound_wire_transfers"
	err = options.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

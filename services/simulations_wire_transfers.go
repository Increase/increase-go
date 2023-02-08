package services

import (
	"context"
	"fmt"
	"increase/options"
	"increase/types"
	"net/url"
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
	opts = append(r.Options, opts...)
	u, err := url.Parse(fmt.Sprintf("simulations/inbound_wire_transfers"))
	if err != nil {
		return
	}
	cfg := options.NewRequestConfig(ctx, "POST", u, opts...)
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return
	}

	return
}

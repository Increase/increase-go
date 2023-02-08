package services

import (
	"context"
	"fmt"
	"increase/options"
	"increase/types"
	"net/url"
)

type SimulationsRealTimePaymentsTransferService struct {
	Options []options.RequestOption
}

func NewSimulationsRealTimePaymentsTransferService(opts ...options.RequestOption) (r *SimulationsRealTimePaymentsTransferService) {
	r = &SimulationsRealTimePaymentsTransferService{}
	r.Options = opts
	return
}

// Simulates an inbound Real Time Payments transfer to your account. Real Time
// Payments are a beta feature.
func (r *SimulationsRealTimePaymentsTransferService) NewInbound(ctx context.Context, body *types.SimulateARealTimePaymentsTransferToYourAccountParameters, opts ...options.RequestOption) (res *types.InboundRealTimePaymentsTransferSimulationResult, err error) {
	opts = append(r.Options, opts...)
	u, err := url.Parse(fmt.Sprintf("simulations/inbound_real_time_payments_transfers"))
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

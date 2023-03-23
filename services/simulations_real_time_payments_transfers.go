package services

import (
	"context"

	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
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
func (r *SimulationsRealTimePaymentsTransferService) NewInbound(ctx context.Context, body *requests.SimulateARealTimePaymentsTransferToYourAccountParameters, opts ...options.RequestOption) (res *responses.InboundRealTimePaymentsTransferSimulationResult, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/inbound_real_time_payments_transfers"
	err = options.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

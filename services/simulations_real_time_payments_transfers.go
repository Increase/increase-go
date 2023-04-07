package services

import (
	"context"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type SimulationsRealTimePaymentsTransferService struct {
	Options []option.RequestOption
}

func NewSimulationsRealTimePaymentsTransferService(opts ...option.RequestOption) (r *SimulationsRealTimePaymentsTransferService) {
	r = &SimulationsRealTimePaymentsTransferService{}
	r.Options = opts
	return
}

// Simulates an inbound Real Time Payments transfer to your account. Real Time
// Payments are a beta feature.
func (r *SimulationsRealTimePaymentsTransferService) NewInbound(ctx context.Context, body *requests.RealTimePaymentsTransferNewInboundParams, opts ...option.RequestOption) (res *responses.InboundRealTimePaymentsTransferSimulationResult, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/inbound_real_time_payments_transfers"
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

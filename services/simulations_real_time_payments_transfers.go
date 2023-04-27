package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type SimulationRealTimePaymentsTransferService struct {
	Options []option.RequestOption
}

func NewSimulationRealTimePaymentsTransferService(opts ...option.RequestOption) (r *SimulationRealTimePaymentsTransferService) {
	r = &SimulationRealTimePaymentsTransferService{}
	r.Options = opts
	return
}

// Simulates submission of a Real Time Payments transfer and handling the response
// from the destination financial institution. This transfer must first have a
// `status` of `pending_submission`.
func (r *SimulationRealTimePaymentsTransferService) Complete(ctx context.Context, real_time_payments_transfer_id string, body requests.SimulationRealTimePaymentsTransferCompleteParams, opts ...option.RequestOption) (res *responses.RealTimePaymentsTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("simulations/real_time_payments_transfers/%s/complete", real_time_payments_transfer_id)
	err = option.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Simulates an inbound Real Time Payments transfer to your account. Real Time
// Payments are a beta feature.
func (r *SimulationRealTimePaymentsTransferService) NewInbound(ctx context.Context, body requests.SimulationRealTimePaymentsTransferNewInboundParams, opts ...option.RequestOption) (res *responses.InboundRealTimePaymentsTransferSimulationResult, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/inbound_real_time_payments_transfers"
	err = option.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

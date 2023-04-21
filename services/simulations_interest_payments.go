package services

import (
	"context"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type SimulationInterestPaymentService struct {
	Options []option.RequestOption
}

func NewSimulationInterestPaymentService(opts ...option.RequestOption) (r *SimulationInterestPaymentService) {
	r = &SimulationInterestPaymentService{}
	r.Options = opts
	return
}

// Simulates an interest payment to your account. In production, this happens
// automatically on the first of each month.
func (r *SimulationInterestPaymentService) New(ctx context.Context, body *requests.SimulationInterestPaymentNewParams, opts ...option.RequestOption) (res *responses.InterestPaymentSimulationResult, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/interest_payment"
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

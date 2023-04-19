package services

import (
	"context"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type SimulationsInterestPaymentService struct {
	Options []option.RequestOption
}

func NewSimulationsInterestPaymentService(opts ...option.RequestOption) (r *SimulationsInterestPaymentService) {
	r = &SimulationsInterestPaymentService{}
	r.Options = opts
	return
}

// Simulates an interest payment to your account. In production, this happens
// automatically on the first of each month.
func (r *SimulationsInterestPaymentService) New(ctx context.Context, body *requests.InterestPaymentNewParams, opts ...option.RequestOption) (res *responses.InterestPaymentSimulationResult, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/interest_payment"
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

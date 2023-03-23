package services

import (
	"context"

	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type SimulationsCardRefundService struct {
	Options []options.RequestOption
}

func NewSimulationsCardRefundService(opts ...options.RequestOption) (r *SimulationsCardRefundService) {
	r = &SimulationsCardRefundService{}
	r.Options = opts
	return
}

// Simulates refunding a card transaction. The full value of the original sandbox
// transaction is refunded.
func (r *SimulationsCardRefundService) New(ctx context.Context, body *requests.SimulateARefundOnACardParameters, opts ...options.RequestOption) (res *responses.Transaction, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/card_refunds"
	err = options.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

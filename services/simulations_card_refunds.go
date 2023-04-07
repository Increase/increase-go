package services

import (
	"context"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type SimulationsCardRefundService struct {
	Options []option.RequestOption
}

func NewSimulationsCardRefundService(opts ...option.RequestOption) (r *SimulationsCardRefundService) {
	r = &SimulationsCardRefundService{}
	r.Options = opts
	return
}

// Simulates refunding a card transaction. The full value of the original sandbox
// transaction is refunded.
func (r *SimulationsCardRefundService) New(ctx context.Context, body *requests.CardRefundNewParams, opts ...option.RequestOption) (res *responses.Transaction, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/card_refunds"
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

package services

import (
	"context"
	"net/http"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type SimulationCardRefundService struct {
	Options []option.RequestOption
}

func NewSimulationCardRefundService(opts ...option.RequestOption) (r *SimulationCardRefundService) {
	r = &SimulationCardRefundService{}
	r.Options = opts
	return
}

// Simulates refunding a card transaction. The full value of the original sandbox
// transaction is refunded.
func (r *SimulationCardRefundService) New(ctx context.Context, body requests.SimulationCardRefundNewParams, opts ...option.RequestOption) (res *responses.Transaction, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/card_refunds"
	err = option.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

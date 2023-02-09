package services

import (
	"bytes"
	"context"
	"fmt"
	"increase/options"
	"increase/types"
	"io"
	"net/url"
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
func (r *SimulationsCardRefundService) New(ctx context.Context, body *types.SimulateARefundOnACardParameters, opts ...options.RequestOption) (res *types.Transaction, err error) {
	opts = append(r.Options[:], opts...)
	b, err := body.MarshalJSON()
	content := io.NopCloser(bytes.NewBuffer(b))
	u, err := url.Parse(fmt.Sprintf("simulations/card_refunds"))
	if err != nil {
		return
	}
	cfg := options.NewRequestConfig(ctx, "POST", u, content, opts...)
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return
	}

	return
}

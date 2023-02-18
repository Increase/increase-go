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

type SimulationsCardDisputeService struct {
	Options []options.RequestOption
}

func NewSimulationsCardDisputeService(opts ...options.RequestOption) (r *SimulationsCardDisputeService) {
	r = &SimulationsCardDisputeService{}
	r.Options = opts
	return
}

// After a [Card Dispute](#card-disputes) is created in production, the dispute
// will be reviewed. Since no review happens in sandbox, this endpoint simulates
// moving a Card Dispute into a rejected or accepted state. A Card Dispute can only
// be actioned one time and must have a status of `pending_reviewing`.
func (r *SimulationsCardDisputeService) Action(ctx context.Context, card_dispute_id string, body *types.SimulatesAdvancingTheStateOfACardDisputeParameters, opts ...options.RequestOption) (res *types.CardDispute, err error) {
	opts = append(r.Options[:], opts...)
	b, err := body.MarshalJSON()
	content := io.NopCloser(bytes.NewBuffer(b))
	u, err := url.Parse(fmt.Sprintf("simulations/card_disputes/%s/action", card_dispute_id))
	if err != nil {
		return
	}
	cfg, err := options.NewRequestConfig(ctx, "POST", u, content, opts...)
	if err != nil {
		return
	}
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return
	}

	return
}

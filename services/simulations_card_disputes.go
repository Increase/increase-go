package services

import (
	"context"
	"fmt"
	"increase/options"
	"increase/types"
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
	path := fmt.Sprintf("simulations/card_disputes/%s/action", card_dispute_id)
	err = options.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

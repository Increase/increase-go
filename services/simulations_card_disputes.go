package services

import (
	"context"
	"fmt"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type SimulationCardDisputeService struct {
	Options []option.RequestOption
}

func NewSimulationCardDisputeService(opts ...option.RequestOption) (r *SimulationCardDisputeService) {
	r = &SimulationCardDisputeService{}
	r.Options = opts
	return
}

// After a [Card Dispute](#card-disputes) is created in production, the dispute
// will be reviewed. Since no review happens in sandbox, this endpoint simulates
// moving a Card Dispute into a rejected or accepted state. A Card Dispute can only
// be actioned one time and must have a status of `pending_reviewing`.
func (r *SimulationCardDisputeService) Action(ctx context.Context, card_dispute_id string, body *requests.SimulationCardDisputeActionParams, opts ...option.RequestOption) (res *responses.CardDispute, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("simulations/card_disputes/%s/action", card_dispute_id)
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

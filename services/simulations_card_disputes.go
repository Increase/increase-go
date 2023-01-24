package services

import (
	"context"
	"fmt"
	"increase/core"
	"increase/types"
)

type SimulationsCardDisputeService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewSimulationsCardDisputeService(requester core.Requester) (r *SimulationsCardDisputeService) {
	r = &SimulationsCardDisputeService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

// After a [Card Dispute](#card-disputes) is created in production, the dispute
// will be reviewed. Since no review happens in sandbox, this endpoint simulates
// moving a Card Dispute into a rejected or accepted state. A Card Dispute can only
// be actioned one time and must have a status of `pending_reviewing`.
func (r *SimulationsCardDisputeService) Action(ctx context.Context, card_dispute_id string, body *types.SimulatesAdvancingTheStateOfACardDisputeParameters, opts ...*core.RequestOpts) (res *types.CardDispute, err error) {
	err = r.post(
		ctx,
		fmt.Sprintf("/simulations/card_disputes/%s/action", card_dispute_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)

	return
}

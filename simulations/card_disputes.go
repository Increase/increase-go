package simulations

import "context"
import "increase/core"
import "increase/card_disputes"
import "fmt"

type CardDisputeService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewCardDisputeService(requester core.Requester) (r *CardDisputeService) {
	r = &CardDisputeService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

// Simulates moving a card dispute into a rejected or accepted state. A dispute can
// only be actioned once and must have a status of `pending_reviewing`.
func (r *CardDisputeService) Action(ctx context.Context, card_dispute_id string, body *SimulatesAdvancingTheStateOfACardDisputeParameters, opts ...*core.RequestOpts) (res *card_disputes.CardDispute, err error) {
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

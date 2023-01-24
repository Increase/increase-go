package services

import (
	"context"
	"increase/core"
	"increase/types"
)

type SimulationsCardService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewSimulationsCardService(requester core.Requester) (r *SimulationsCardService) {
	r = &SimulationsCardService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

// Simulates activity on a Card. You can pass either a Card id or a Digital Wallet
// Token id to simulate the two different ways purchases can be made.
func (r *SimulationsCardService) Authorize(ctx context.Context, body *types.SimulateAnAuthorizationOnACardParameters, opts ...*core.RequestOpts) (res *types.CardAuthorizationSimulation, err error) {
	err = r.post(
		ctx,
		"/simulations/card_authorizations",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)

	return
}

// Simulates the settlement of an authorization by a card acquirer.
func (r *SimulationsCardService) Settlement(ctx context.Context, body *types.SimulateSettlingACardAuthorizationParameters, opts ...*core.RequestOpts) (res *types.Transaction, err error) {
	err = r.post(
		ctx,
		"/simulations/card_settlements",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)

	return
}

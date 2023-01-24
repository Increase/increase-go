package services

import (
	"context"
	"increase/core"
	"increase/types"
)

type SimulationsDigitalWalletTokenRequestService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewSimulationsDigitalWalletTokenRequestService(requester core.Requester) (r *SimulationsDigitalWalletTokenRequestService) {
	r = &SimulationsDigitalWalletTokenRequestService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

// Simulates a user attempting add a [Card](#cards) to a digital wallet such as
// Apple Pay.
func (r *SimulationsDigitalWalletTokenRequestService) Create(ctx context.Context, body *types.SimulateDigitalWalletProvisioningForACardParameters, opts ...*core.RequestOpts) (res *types.DigitalWalletTokenRequestCreateResponse, err error) {
	err = r.post(
		ctx,
		"/simulations/digital_wallet_token_requests",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)

	return
}

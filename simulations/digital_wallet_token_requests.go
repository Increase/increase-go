package simulations

import "context"
import "increase/core"
import "increase/shared"

type DigitalWalletTokenRequestService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewDigitalWalletTokenRequestService(requester core.Requester) (r *DigitalWalletTokenRequestService) {
	r = &DigitalWalletTokenRequestService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

type PreloadedDigitalWalletTokenRequestService struct {
	DigitalWalletTokenRequests *DigitalWalletTokenRequestService
}

func (r *PreloadedDigitalWalletTokenRequestService) Init(service *DigitalWalletTokenRequestService) {
	r.DigitalWalletTokenRequests = service
}

func NewPreloadedDigitalWalletTokenRequestService(service *DigitalWalletTokenRequestService) (r *PreloadedDigitalWalletTokenRequestService) {
	r = &PreloadedDigitalWalletTokenRequestService{}
	r.Init(service)
	return
}

// Simulates a user attempting add a Card to a digital wallet such as Apple Pay.
func (r *DigitalWalletTokenRequestService) Create(ctx context.Context, body *SimulateDigitalWalletActivityOnACardParameters, opts ...*core.RequestOpts) (res *shared.InboundDigitalWalletTokenRequestSimulationResult, err error) {
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

// Simulates a user attempting add a Card to a digital wallet such as Apple Pay.
func (r *PreloadedDigitalWalletTokenRequestService) Create(ctx context.Context, body *SimulateDigitalWalletActivityOnACardParameters, opts ...*core.RequestOpts) (res *shared.InboundDigitalWalletTokenRequestSimulationResult, err error) {
	err = r.DigitalWalletTokenRequests.post(
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

package simulations

import "context"
import "increase/core"

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

// Simulates a user attempting add a Card to a digital wallet such as Apple Pay.
func (r *DigitalWalletTokenRequestService) Create(ctx context.Context, body *SimulateDigitalWalletActivityOnACardParameters, opts ...*core.RequestOpts) (res *DigitalWalletTokenRequestCreateResponse, err error) {
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

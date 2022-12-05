package simulations

import "increase/core"
import "increase/shared"

type DigitalWalletTokenRequests struct {
	Requester core.Requester
	get       func(string, *core.CoreRequest, interface{}) error
	post      func(string, *core.CoreRequest, interface{}) error
	patch     func(string, *core.CoreRequest, interface{}) error
	put       func(string, *core.CoreRequest, interface{}) error
	delete    func(string, *core.CoreRequest, interface{}) error
}

func NewDigitalWalletTokenRequests(requster core.Requester) (r *DigitalWalletTokenRequests) {
	r = &DigitalWalletTokenRequests{}
	r.Requester = requster
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

// Simulates a user attempting add a Card to a digital wallet such as Apple Pay.
func (r *DigitalWalletTokenRequests) Create(body DigitalWalletTokenRequestsCreateParams, opts ...*core.RequestOpts) (res shared.InboundDigitalWalletTokenRequestSimulationResult, err error) {
	err = r.post(
		"/simulations/digital_wallet_token_requests",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

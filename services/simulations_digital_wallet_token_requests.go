package services

import (
	"context"
	"increase/options"
	"increase/types"
)

type SimulationsDigitalWalletTokenRequestService struct {
	Options []options.RequestOption
}

func NewSimulationsDigitalWalletTokenRequestService(opts ...options.RequestOption) (r *SimulationsDigitalWalletTokenRequestService) {
	r = &SimulationsDigitalWalletTokenRequestService{}
	r.Options = opts
	return
}

// Simulates a user attempting add a [Card](#cards) to a digital wallet such as
// Apple Pay.
func (r *SimulationsDigitalWalletTokenRequestService) New(ctx context.Context, body *types.SimulateDigitalWalletProvisioningForACardParameters, opts ...options.RequestOption) (res *types.DigitalWalletTokenRequestCreateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/digital_wallet_token_requests"
	err = options.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

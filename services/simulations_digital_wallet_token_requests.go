package services

import (
	"context"
	"fmt"
	"increase/options"
	"increase/types"
	"net/url"
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
	opts = append(r.Options, opts...)
	u, err := url.Parse(fmt.Sprintf("simulations/digital_wallet_token_requests"))
	if err != nil {
		return
	}
	cfg := options.NewRequestConfig(ctx, "POST", u, opts...)
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return
	}

	return
}

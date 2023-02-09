package services

import (
	"bytes"
	"context"
	"fmt"
	"increase/options"
	"increase/types"
	"io"
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
	opts = append(r.Options[:], opts...)
	b, err := body.MarshalJSON()
	content := io.NopCloser(bytes.NewBuffer(b))
	u, err := url.Parse(fmt.Sprintf("simulations/digital_wallet_token_requests"))
	if err != nil {
		return
	}
	cfg := options.NewRequestConfig(ctx, "POST", u, content, opts...)
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return
	}

	return
}

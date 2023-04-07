package services

import (
	"context"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type SimulationsDigitalWalletTokenRequestService struct {
	Options []option.RequestOption
}

func NewSimulationsDigitalWalletTokenRequestService(opts ...option.RequestOption) (r *SimulationsDigitalWalletTokenRequestService) {
	r = &SimulationsDigitalWalletTokenRequestService{}
	r.Options = opts
	return
}

// Simulates a user attempting add a [Card](#cards) to a digital wallet such as
// Apple Pay.
func (r *SimulationsDigitalWalletTokenRequestService) New(ctx context.Context, body *requests.DigitalWalletTokenRequestNewParams, opts ...option.RequestOption) (res *responses.DigitalWalletTokenRequestCreateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/digital_wallet_token_requests"
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

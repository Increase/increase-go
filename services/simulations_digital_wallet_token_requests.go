package services

import (
	"context"
	"net/http"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type SimulationDigitalWalletTokenRequestService struct {
	Options []option.RequestOption
}

func NewSimulationDigitalWalletTokenRequestService(opts ...option.RequestOption) (r *SimulationDigitalWalletTokenRequestService) {
	r = &SimulationDigitalWalletTokenRequestService{}
	r.Options = opts
	return
}

// Simulates a user attempting add a [Card](#cards) to a digital wallet such as
// Apple Pay.
func (r *SimulationDigitalWalletTokenRequestService) New(ctx context.Context, body requests.SimulationDigitalWalletTokenRequestNewParams, opts ...option.RequestOption) (res *responses.DigitalWalletTokenRequestCreateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/digital_wallet_token_requests"
	err = option.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

package services

import (
	"context"
	"net/http"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type SimulationCardService struct {
	Options []option.RequestOption
}

func NewSimulationCardService(opts ...option.RequestOption) (r *SimulationCardService) {
	r = &SimulationCardService{}
	r.Options = opts
	return
}

// Simulates a purchase authorization on a [Card](#cards). Depending on the balance
// available to the card and the `amount` submitted, the authorization activity
// will result in a [Pending Transaction](#pending-transactions) of type
// `card_authorization` or a [Declined Transaction](#declined-transactions) of type
// `card_decline`. You can pass either a Card id or a
// [Digital Wallet Token](#digital-wallet-tokens) id to simulate the two different
// ways purchases can be made.
func (r *SimulationCardService) Authorize(ctx context.Context, body requests.SimulationCardAuthorizeParams, opts ...option.RequestOption) (res *responses.CardAuthorizationSimulation, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/card_authorizations"
	err = option.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Simulates the settlement of an authorization by a card acquirer. After a card
// authorization is created, the merchant will eventually send a settlement. This
// simulates that event, which may occur many days after the purchase in
// production. The amount settled can be different from the amount originally
// authorized, for example, when adding a tip to a restaurant bill.
func (r *SimulationCardService) Settlement(ctx context.Context, body requests.SimulationCardSettlementParams, opts ...option.RequestOption) (res *responses.Transaction, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/card_settlements"
	err = option.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

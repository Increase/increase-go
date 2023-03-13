package services

import (
	"context"

	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/types"
)

type SimulationsCardService struct {
	Options []options.RequestOption
}

func NewSimulationsCardService(opts ...options.RequestOption) (r *SimulationsCardService) {
	r = &SimulationsCardService{}
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
func (r *SimulationsCardService) Authorize(ctx context.Context, body *types.SimulateAnAuthorizationOnACardParameters, opts ...options.RequestOption) (res *types.CardAuthorizationSimulation, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/card_authorizations"
	err = options.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Simulates the settlement of an authorization by a card acquirer. After a card
// authorization is created, the merchant will eventually send a settlement. This
// simulates that event, which may occur many days after the purchase in
// production. The amount settled can be different from the amount originally
// authorized, for example, when adding a tip to a restaurant bill.
func (r *SimulationsCardService) Settlement(ctx context.Context, body *types.SimulateSettlingACardAuthorizationParameters, opts ...options.RequestOption) (res *types.Transaction, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/card_settlements"
	err = options.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

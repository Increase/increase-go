// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"net/http"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/option"
)

// SimulationCardAuthorizationService contains methods and other services that help
// with interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimulationCardAuthorizationService] method instead.
type SimulationCardAuthorizationService struct {
	Options []option.RequestOption
}

// NewSimulationCardAuthorizationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSimulationCardAuthorizationService(opts ...option.RequestOption) (r *SimulationCardAuthorizationService) {
	r = &SimulationCardAuthorizationService{}
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
func (r *SimulationCardAuthorizationService) New(ctx context.Context, body SimulationCardAuthorizationNewParams, opts ...option.RequestOption) (res *SimulationCardAuthorizationNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/card_authorizations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// The results of a Card Authorization simulation.
type SimulationCardAuthorizationNewResponse struct {
	// If the authorization attempt fails, this will contain the resulting
	// [Declined Transaction](#declined-transactions) object. The Declined
	// Transaction's `source` will be of `category: card_decline`.
	DeclinedTransaction DeclinedTransaction `json:"declined_transaction,required,nullable"`
	// If the authorization attempt succeeds, this will contain the resulting Pending
	// Transaction object. The Pending Transaction's `source` will be of
	// `category: card_authorization`.
	PendingTransaction PendingTransaction `json:"pending_transaction,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `inbound_card_authorization_simulation_result`.
	Type SimulationCardAuthorizationNewResponseType `json:"type,required"`
	JSON simulationCardAuthorizationNewResponseJSON `json:"-"`
}

// simulationCardAuthorizationNewResponseJSON contains the JSON metadata for the
// struct [SimulationCardAuthorizationNewResponse]
type simulationCardAuthorizationNewResponseJSON struct {
	DeclinedTransaction apijson.Field
	PendingTransaction  apijson.Field
	Type                apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *SimulationCardAuthorizationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r simulationCardAuthorizationNewResponseJSON) RawJSON() string {
	return r.raw
}

// A constant representing the object's type. For this resource it will always be
// `inbound_card_authorization_simulation_result`.
type SimulationCardAuthorizationNewResponseType string

const (
	SimulationCardAuthorizationNewResponseTypeInboundCardAuthorizationSimulationResult SimulationCardAuthorizationNewResponseType = "inbound_card_authorization_simulation_result"
)

func (r SimulationCardAuthorizationNewResponseType) IsKnown() bool {
	switch r {
	case SimulationCardAuthorizationNewResponseTypeInboundCardAuthorizationSimulationResult:
		return true
	}
	return false
}

type SimulationCardAuthorizationNewParams struct {
	// The authorization amount in cents.
	Amount param.Field[int64] `json:"amount,required"`
	// The identifier of the Card to be authorized.
	CardID param.Field[string] `json:"card_id"`
	// The identifier of the Digital Wallet Token to be authorized.
	DigitalWalletTokenID param.Field[string] `json:"digital_wallet_token_id"`
	// The identifier of the Event Subscription to use. If provided, will override the
	// default real time event subscription. Because you can only create one real time
	// decision event subscription, you can use this field to route events to any
	// specified event subscription for testing purposes.
	EventSubscriptionID param.Field[string] `json:"event_subscription_id"`
	// The merchant identifier (commonly abbreviated as MID) of the merchant the card
	// is transacting with.
	MerchantAcceptorID param.Field[string] `json:"merchant_acceptor_id"`
	// The Merchant Category Code (commonly abbreviated as MCC) of the merchant the
	// card is transacting with.
	MerchantCategoryCode param.Field[string] `json:"merchant_category_code"`
	// The city the merchant resides in.
	MerchantCity param.Field[string] `json:"merchant_city"`
	// The country the merchant resides in.
	MerchantCountry param.Field[string] `json:"merchant_country"`
	// The merchant descriptor of the merchant the card is transacting with.
	MerchantDescriptor param.Field[string] `json:"merchant_descriptor"`
	// The identifier of the Physical Card to be authorized.
	PhysicalCardID param.Field[string] `json:"physical_card_id"`
}

func (r SimulationCardAuthorizationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

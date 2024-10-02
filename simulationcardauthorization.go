// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"net/http"

	"github.com/Increase/increase-go/internal/apijson"
	"github.com/Increase/increase-go/internal/param"
	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
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
	// Forces a card decline with a specific reason. No real time decision will be
	// sent.
	DeclineReason param.Field[SimulationCardAuthorizationNewParamsDeclineReason] `json:"decline_reason"`
	// The identifier of the Digital Wallet Token to be authorized.
	DigitalWalletTokenID param.Field[string] `json:"digital_wallet_token_id"`
	// The direction describes the direction the funds will move, either from the
	// cardholder to the merchant or from the merchant to the cardholder.
	Direction param.Field[SimulationCardAuthorizationNewParamsDirection] `json:"direction"`
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

// Forces a card decline with a specific reason. No real time decision will be
// sent.
type SimulationCardAuthorizationNewParamsDeclineReason string

const (
	// The Card was not active.
	SimulationCardAuthorizationNewParamsDeclineReasonCardNotActive SimulationCardAuthorizationNewParamsDeclineReason = "card_not_active"
	// The Physical Card was not active.
	SimulationCardAuthorizationNewParamsDeclineReasonPhysicalCardNotActive SimulationCardAuthorizationNewParamsDeclineReason = "physical_card_not_active"
	// The account's entity was not active.
	SimulationCardAuthorizationNewParamsDeclineReasonEntityNotActive SimulationCardAuthorizationNewParamsDeclineReason = "entity_not_active"
	// The account was inactive.
	SimulationCardAuthorizationNewParamsDeclineReasonGroupLocked SimulationCardAuthorizationNewParamsDeclineReason = "group_locked"
	// The Card's Account did not have a sufficient available balance.
	SimulationCardAuthorizationNewParamsDeclineReasonInsufficientFunds SimulationCardAuthorizationNewParamsDeclineReason = "insufficient_funds"
	// The given CVV2 did not match the card's value.
	SimulationCardAuthorizationNewParamsDeclineReasonCvv2Mismatch SimulationCardAuthorizationNewParamsDeclineReason = "cvv2_mismatch"
	// The given expiration date did not match the card's value. Only applies when a
	// CVV2 is present.
	SimulationCardAuthorizationNewParamsDeclineReasonCardExpirationMismatch SimulationCardAuthorizationNewParamsDeclineReason = "card_expiration_mismatch"
	// The attempted card transaction is not allowed per Increase's terms.
	SimulationCardAuthorizationNewParamsDeclineReasonTransactionNotAllowed SimulationCardAuthorizationNewParamsDeclineReason = "transaction_not_allowed"
	// The transaction was blocked by a Limit.
	SimulationCardAuthorizationNewParamsDeclineReasonBreachesLimit SimulationCardAuthorizationNewParamsDeclineReason = "breaches_limit"
	// Your application declined the transaction via webhook.
	SimulationCardAuthorizationNewParamsDeclineReasonWebhookDeclined SimulationCardAuthorizationNewParamsDeclineReason = "webhook_declined"
	// Your application webhook did not respond without the required timeout.
	SimulationCardAuthorizationNewParamsDeclineReasonWebhookTimedOut SimulationCardAuthorizationNewParamsDeclineReason = "webhook_timed_out"
	// Declined by stand-in processing.
	SimulationCardAuthorizationNewParamsDeclineReasonDeclinedByStandInProcessing SimulationCardAuthorizationNewParamsDeclineReason = "declined_by_stand_in_processing"
	// The card read had an invalid CVV, dCVV, or authorization request cryptogram.
	SimulationCardAuthorizationNewParamsDeclineReasonInvalidPhysicalCard SimulationCardAuthorizationNewParamsDeclineReason = "invalid_physical_card"
	// The original card authorization for this incremental authorization does not
	// exist.
	SimulationCardAuthorizationNewParamsDeclineReasonMissingOriginalAuthorization SimulationCardAuthorizationNewParamsDeclineReason = "missing_original_authorization"
	// The transaction was suspected to be fraudulent. Please reach out to
	// support@increase.com for more information.
	SimulationCardAuthorizationNewParamsDeclineReasonSuspectedFraud SimulationCardAuthorizationNewParamsDeclineReason = "suspected_fraud"
)

func (r SimulationCardAuthorizationNewParamsDeclineReason) IsKnown() bool {
	switch r {
	case SimulationCardAuthorizationNewParamsDeclineReasonCardNotActive, SimulationCardAuthorizationNewParamsDeclineReasonPhysicalCardNotActive, SimulationCardAuthorizationNewParamsDeclineReasonEntityNotActive, SimulationCardAuthorizationNewParamsDeclineReasonGroupLocked, SimulationCardAuthorizationNewParamsDeclineReasonInsufficientFunds, SimulationCardAuthorizationNewParamsDeclineReasonCvv2Mismatch, SimulationCardAuthorizationNewParamsDeclineReasonCardExpirationMismatch, SimulationCardAuthorizationNewParamsDeclineReasonTransactionNotAllowed, SimulationCardAuthorizationNewParamsDeclineReasonBreachesLimit, SimulationCardAuthorizationNewParamsDeclineReasonWebhookDeclined, SimulationCardAuthorizationNewParamsDeclineReasonWebhookTimedOut, SimulationCardAuthorizationNewParamsDeclineReasonDeclinedByStandInProcessing, SimulationCardAuthorizationNewParamsDeclineReasonInvalidPhysicalCard, SimulationCardAuthorizationNewParamsDeclineReasonMissingOriginalAuthorization, SimulationCardAuthorizationNewParamsDeclineReasonSuspectedFraud:
		return true
	}
	return false
}

// The direction describes the direction the funds will move, either from the
// cardholder to the merchant or from the merchant to the cardholder.
type SimulationCardAuthorizationNewParamsDirection string

const (
	// A regular card authorization where funds are debited from the cardholder.
	SimulationCardAuthorizationNewParamsDirectionSettlement SimulationCardAuthorizationNewParamsDirection = "settlement"
	// A refund card authorization, sometimes referred to as a credit voucher
	// authorization, where funds are credited to the cardholder.
	SimulationCardAuthorizationNewParamsDirectionRefund SimulationCardAuthorizationNewParamsDirection = "refund"
)

func (r SimulationCardAuthorizationNewParamsDirection) IsKnown() bool {
	switch r {
	case SimulationCardAuthorizationNewParamsDirectionSettlement, SimulationCardAuthorizationNewParamsDirectionRefund:
		return true
	}
	return false
}

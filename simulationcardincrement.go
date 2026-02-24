// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"net/http"
	"slices"

	"github.com/Increase/increase-go/internal/apijson"
	"github.com/Increase/increase-go/internal/param"
	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
)

// SimulationCardIncrementService contains methods and other services that help
// with interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimulationCardIncrementService] method instead.
type SimulationCardIncrementService struct {
	Options []option.RequestOption
}

// NewSimulationCardIncrementService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSimulationCardIncrementService(opts ...option.RequestOption) (r *SimulationCardIncrementService) {
	r = &SimulationCardIncrementService{}
	r.Options = opts
	return
}

// Simulates the increment of an authorization by a card acquirer. An authorization
// can be incremented multiple times.
func (r *SimulationCardIncrementService) New(ctx context.Context, body SimulationCardIncrementNewParams, opts ...option.RequestOption) (res *CardPayment, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "simulations/card_increments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SimulationCardIncrementNewParams struct {
	// The amount of the increment in minor units in the card authorization's currency.
	Amount param.Field[int64] `json:"amount" api:"required"`
	// The identifier of the Card Payment to create an increment on.
	CardPaymentID param.Field[string] `json:"card_payment_id" api:"required"`
	// The identifier of the Event Subscription to use. If provided, will override the
	// default real time event subscription. Because you can only create one real time
	// decision event subscription, you can use this field to route events to any
	// specified event subscription for testing purposes.
	EventSubscriptionID param.Field[string] `json:"event_subscription_id"`
}

func (r SimulationCardIncrementNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

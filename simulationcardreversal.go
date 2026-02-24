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

// SimulationCardReversalService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimulationCardReversalService] method instead.
type SimulationCardReversalService struct {
	Options []option.RequestOption
}

// NewSimulationCardReversalService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSimulationCardReversalService(opts ...option.RequestOption) (r *SimulationCardReversalService) {
	r = &SimulationCardReversalService{}
	r.Options = opts
	return
}

// Simulates the reversal of an authorization by a card acquirer. An authorization
// can be partially reversed multiple times, up until the total authorized amount.
// Marks the pending transaction as complete if the authorization is fully
// reversed.
func (r *SimulationCardReversalService) New(ctx context.Context, body SimulationCardReversalNewParams, opts ...option.RequestOption) (res *CardPayment, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "simulations/card_reversals"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SimulationCardReversalNewParams struct {
	// The identifier of the Card Payment to create a reversal on.
	CardPaymentID param.Field[string] `json:"card_payment_id" api:"required"`
	// The amount of the reversal in minor units in the card authorization's currency.
	// This defaults to the authorization amount.
	Amount param.Field[int64] `json:"amount"`
}

func (r SimulationCardReversalNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

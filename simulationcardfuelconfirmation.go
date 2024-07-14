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

// SimulationCardFuelConfirmationService contains methods and other services that
// help with interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimulationCardFuelConfirmationService] method instead.
type SimulationCardFuelConfirmationService struct {
	Options []option.RequestOption
}

// NewSimulationCardFuelConfirmationService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSimulationCardFuelConfirmationService(opts ...option.RequestOption) (r *SimulationCardFuelConfirmationService) {
	r = &SimulationCardFuelConfirmationService{}
	r.Options = opts
	return
}

// Simulates the fuel confirmation of an authorization by a card acquirer. This
// happens asynchronously right after a fuel pump transaction is completed. A fuel
// confirmation can only happen once per authorization.
func (r *SimulationCardFuelConfirmationService) New(ctx context.Context, body SimulationCardFuelConfirmationNewParams, opts ...option.RequestOption) (res *CardPayment, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/card_fuel_confirmations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SimulationCardFuelConfirmationNewParams struct {
	// The amount of the fuel_confirmation in minor units in the card authorization's
	// currency.
	Amount param.Field[int64] `json:"amount,required"`
	// The identifier of the Card Payment to create a fuel_confirmation on.
	CardPaymentID param.Field[string] `json:"card_payment_id,required"`
}

func (r SimulationCardFuelConfirmationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

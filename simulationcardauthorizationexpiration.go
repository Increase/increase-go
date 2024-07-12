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

// SimulationCardAuthorizationExpirationService contains methods and other services
// that help with interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimulationCardAuthorizationExpirationService] method instead.
type SimulationCardAuthorizationExpirationService struct {
	Options []option.RequestOption
}

// NewSimulationCardAuthorizationExpirationService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewSimulationCardAuthorizationExpirationService(opts ...option.RequestOption) (r *SimulationCardAuthorizationExpirationService) {
	r = &SimulationCardAuthorizationExpirationService{}
	r.Options = opts
	return
}

// Simulates expiring a card authorization immediately.
func (r *SimulationCardAuthorizationExpirationService) New(ctx context.Context, body SimulationCardAuthorizationExpirationNewParams, opts ...option.RequestOption) (res *CardPayment, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/card_authorization_expirations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SimulationCardAuthorizationExpirationNewParams struct {
	// The identifier of the Card Payment to expire.
	CardPaymentID param.Field[string] `json:"card_payment_id,required"`
}

func (r SimulationCardAuthorizationExpirationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/Increase/increase-go/internal/apijson"
	"github.com/Increase/increase-go/internal/param"
	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
)

// SimulationCardTokenService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimulationCardTokenService] method instead.
type SimulationCardTokenService struct {
	Options []option.RequestOption
}

// NewSimulationCardTokenService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSimulationCardTokenService(opts ...option.RequestOption) (r *SimulationCardTokenService) {
	r = &SimulationCardTokenService{}
	r.Options = opts
	return
}

// Simulates tokenizing a card in the sandbox environment.
func (r *SimulationCardTokenService) New(ctx context.Context, body SimulationCardTokenNewParams, opts ...option.RequestOption) (res *CardToken, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "simulations/card_tokens"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SimulationCardTokenNewParams struct {
	// The capabilities of the outbound card token.
	Capabilities param.Field[[]SimulationCardTokenNewParamsCapability] `json:"capabilities"`
	// The expiration date of the card.
	Expiration param.Field[time.Time] `json:"expiration" format:"date"`
	// The last 4 digits of the card number.
	Last4 param.Field[string] `json:"last4"`
	// The prefix of the card number, usually the first 8 digits.
	Prefix param.Field[string] `json:"prefix"`
	// The total length of the card number, including prefix and last4.
	PrimaryAccountNumberLength param.Field[int64] `json:"primary_account_number_length"`
}

func (r SimulationCardTokenNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SimulationCardTokenNewParamsCapability struct {
	// The cross-border push transfers capability.
	CrossBorderPushTransfers param.Field[SimulationCardTokenNewParamsCapabilitiesCrossBorderPushTransfers] `json:"cross_border_push_transfers" api:"required"`
	// The domestic push transfers capability.
	DomesticPushTransfers param.Field[SimulationCardTokenNewParamsCapabilitiesDomesticPushTransfers] `json:"domestic_push_transfers" api:"required"`
	// The route of the capability.
	Route param.Field[SimulationCardTokenNewParamsCapabilitiesRoute] `json:"route" api:"required"`
}

func (r SimulationCardTokenNewParamsCapability) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cross-border push transfers capability.
type SimulationCardTokenNewParamsCapabilitiesCrossBorderPushTransfers string

const (
	SimulationCardTokenNewParamsCapabilitiesCrossBorderPushTransfersSupported    SimulationCardTokenNewParamsCapabilitiesCrossBorderPushTransfers = "supported"
	SimulationCardTokenNewParamsCapabilitiesCrossBorderPushTransfersNotSupported SimulationCardTokenNewParamsCapabilitiesCrossBorderPushTransfers = "not_supported"
)

func (r SimulationCardTokenNewParamsCapabilitiesCrossBorderPushTransfers) IsKnown() bool {
	switch r {
	case SimulationCardTokenNewParamsCapabilitiesCrossBorderPushTransfersSupported, SimulationCardTokenNewParamsCapabilitiesCrossBorderPushTransfersNotSupported:
		return true
	}
	return false
}

// The domestic push transfers capability.
type SimulationCardTokenNewParamsCapabilitiesDomesticPushTransfers string

const (
	SimulationCardTokenNewParamsCapabilitiesDomesticPushTransfersSupported    SimulationCardTokenNewParamsCapabilitiesDomesticPushTransfers = "supported"
	SimulationCardTokenNewParamsCapabilitiesDomesticPushTransfersNotSupported SimulationCardTokenNewParamsCapabilitiesDomesticPushTransfers = "not_supported"
)

func (r SimulationCardTokenNewParamsCapabilitiesDomesticPushTransfers) IsKnown() bool {
	switch r {
	case SimulationCardTokenNewParamsCapabilitiesDomesticPushTransfersSupported, SimulationCardTokenNewParamsCapabilitiesDomesticPushTransfersNotSupported:
		return true
	}
	return false
}

// The route of the capability.
type SimulationCardTokenNewParamsCapabilitiesRoute string

const (
	SimulationCardTokenNewParamsCapabilitiesRouteVisa       SimulationCardTokenNewParamsCapabilitiesRoute = "visa"
	SimulationCardTokenNewParamsCapabilitiesRouteMastercard SimulationCardTokenNewParamsCapabilitiesRoute = "mastercard"
)

func (r SimulationCardTokenNewParamsCapabilitiesRoute) IsKnown() bool {
	switch r {
	case SimulationCardTokenNewParamsCapabilitiesRouteVisa, SimulationCardTokenNewParamsCapabilitiesRouteMastercard:
		return true
	}
	return false
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/option"
)

// SimulationInboundFundsHoldService contains methods and other services that help
// with interacting with the increase API. Note, unlike clients, this service does
// not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewSimulationInboundFundsHoldService] method instead.
type SimulationInboundFundsHoldService struct {
	Options []option.RequestOption
}

// NewSimulationInboundFundsHoldService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSimulationInboundFundsHoldService(opts ...option.RequestOption) (r *SimulationInboundFundsHoldService) {
	r = &SimulationInboundFundsHoldService{}
	r.Options = opts
	return
}

// This endpoint simulates immediately releasing an inbound funds hold, which might
// be created as a result of e.g., an ACH debit.
func (r *SimulationInboundFundsHoldService) Release(ctx context.Context, inboundFundsHoldID string, opts ...option.RequestOption) (res *SimulationInboundFundsHoldReleaseResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("simulations/inbound_funds_holds/%s/release", inboundFundsHoldID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// We hold funds for certain transaction types to account for return windows where
// funds might still be clawed back by the sending institution.
type SimulationInboundFundsHoldReleaseResponse struct {
	// The Inbound Funds Hold identifier.
	ID string `json:"id,required"`
	// The held amount in the minor unit of the account's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// When the hold will be released automatically. Certain conditions may cause it to
	// be released before this time.
	AutomaticallyReleasesAt time.Time `json:"automatically_releases_at,required" format:"date-time"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the hold
	// was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the hold's
	// currency.
	Currency SimulationInboundFundsHoldReleaseResponseCurrency `json:"currency,required"`
	// The ID of the Transaction for which funds were held.
	HeldTransactionID string `json:"held_transaction_id,required,nullable"`
	// The ID of the Pending Transaction representing the held funds.
	PendingTransactionID string `json:"pending_transaction_id,required,nullable"`
	// When the hold was released (if it has been released).
	ReleasedAt time.Time `json:"released_at,required,nullable" format:"date-time"`
	// The status of the hold.
	Status SimulationInboundFundsHoldReleaseResponseStatus `json:"status,required"`
	// A constant representing the object's type. For this resource it will always be
	// `inbound_funds_hold`.
	Type SimulationInboundFundsHoldReleaseResponseType `json:"type,required"`
	JSON simulationInboundFundsHoldReleaseResponseJSON `json:"-"`
}

// simulationInboundFundsHoldReleaseResponseJSON contains the JSON metadata for the
// struct [SimulationInboundFundsHoldReleaseResponse]
type simulationInboundFundsHoldReleaseResponseJSON struct {
	ID                      apijson.Field
	Amount                  apijson.Field
	AutomaticallyReleasesAt apijson.Field
	CreatedAt               apijson.Field
	Currency                apijson.Field
	HeldTransactionID       apijson.Field
	PendingTransactionID    apijson.Field
	ReleasedAt              apijson.Field
	Status                  apijson.Field
	Type                    apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *SimulationInboundFundsHoldReleaseResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r simulationInboundFundsHoldReleaseResponseJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the hold's
// currency.
type SimulationInboundFundsHoldReleaseResponseCurrency string

const (
	// Canadian Dollar (CAD)
	SimulationInboundFundsHoldReleaseResponseCurrencyCad SimulationInboundFundsHoldReleaseResponseCurrency = "CAD"
	// Swiss Franc (CHF)
	SimulationInboundFundsHoldReleaseResponseCurrencyChf SimulationInboundFundsHoldReleaseResponseCurrency = "CHF"
	// Euro (EUR)
	SimulationInboundFundsHoldReleaseResponseCurrencyEur SimulationInboundFundsHoldReleaseResponseCurrency = "EUR"
	// British Pound (GBP)
	SimulationInboundFundsHoldReleaseResponseCurrencyGbp SimulationInboundFundsHoldReleaseResponseCurrency = "GBP"
	// Japanese Yen (JPY)
	SimulationInboundFundsHoldReleaseResponseCurrencyJpy SimulationInboundFundsHoldReleaseResponseCurrency = "JPY"
	// US Dollar (USD)
	SimulationInboundFundsHoldReleaseResponseCurrencyUsd SimulationInboundFundsHoldReleaseResponseCurrency = "USD"
)

// The status of the hold.
type SimulationInboundFundsHoldReleaseResponseStatus string

const (
	// Funds are still being held.
	SimulationInboundFundsHoldReleaseResponseStatusHeld SimulationInboundFundsHoldReleaseResponseStatus = "held"
	// Funds have been released.
	SimulationInboundFundsHoldReleaseResponseStatusComplete SimulationInboundFundsHoldReleaseResponseStatus = "complete"
)

// A constant representing the object's type. For this resource it will always be
// `inbound_funds_hold`.
type SimulationInboundFundsHoldReleaseResponseType string

const (
	SimulationInboundFundsHoldReleaseResponseTypeInboundFundsHold SimulationInboundFundsHoldReleaseResponseType = "inbound_funds_hold"
)

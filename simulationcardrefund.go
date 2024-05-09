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

// SimulationCardRefundService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimulationCardRefundService] method instead.
type SimulationCardRefundService struct {
	Options []option.RequestOption
}

// NewSimulationCardRefundService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSimulationCardRefundService(opts ...option.RequestOption) (r *SimulationCardRefundService) {
	r = &SimulationCardRefundService{}
	r.Options = opts
	return
}

// Simulates refunding a card transaction. The full value of the original sandbox
// transaction is refunded.
func (r *SimulationCardRefundService) New(ctx context.Context, body SimulationCardRefundNewParams, opts ...option.RequestOption) (res *Transaction, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/card_refunds"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SimulationCardRefundNewParams struct {
	// The identifier for the Transaction to refund. The Transaction's source must have
	// a category of card_settlement.
	TransactionID param.Field[string] `json:"transaction_id,required"`
}

func (r SimulationCardRefundNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

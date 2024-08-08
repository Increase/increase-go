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

// SimulationCardSettlementService contains methods and other services that help
// with interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimulationCardSettlementService] method instead.
type SimulationCardSettlementService struct {
	Options []option.RequestOption
}

// NewSimulationCardSettlementService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSimulationCardSettlementService(opts ...option.RequestOption) (r *SimulationCardSettlementService) {
	r = &SimulationCardSettlementService{}
	r.Options = opts
	return
}

// Simulates the settlement of an authorization by a card acquirer. After a card
// authorization is created, the merchant will eventually send a settlement. This
// simulates that event, which may occur many days after the purchase in
// production. The amount settled can be different from the amount originally
// authorized, for example, when adding a tip to a restaurant bill.
func (r *SimulationCardSettlementService) New(ctx context.Context, body SimulationCardSettlementNewParams, opts ...option.RequestOption) (res *Transaction, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/card_settlements"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SimulationCardSettlementNewParams struct {
	// The identifier of the Card to create a settlement on.
	CardID param.Field[string] `json:"card_id,required"`
	// The identifier of the Pending Transaction for the Card Authorization you wish to
	// settle.
	PendingTransactionID param.Field[string] `json:"pending_transaction_id,required"`
	// The amount to be settled. This defaults to the amount of the Pending Transaction
	// being settled.
	Amount param.Field[int64] `json:"amount"`
}

func (r SimulationCardSettlementNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

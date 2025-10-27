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
	opts = slices.Concat(r.Options, opts)
	path := "simulations/card_refunds"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SimulationCardRefundNewParams struct {
	// The refund amount in cents. Pulled off the `pending_transaction` or the
	// `transaction` if not provided.
	Amount param.Field[int64] `json:"amount"`
	// The identifier of the Pending Transaction for the refund authorization. If this
	// is provided, `transaction` must not be provided as a refund with a refund
	// authorized can not be linked to a regular transaction.
	PendingTransactionID param.Field[string] `json:"pending_transaction_id"`
	// The identifier for the Transaction to refund. The Transaction's source must have
	// a category of card_settlement.
	TransactionID param.Field[string] `json:"transaction_id"`
}

func (r SimulationCardRefundNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

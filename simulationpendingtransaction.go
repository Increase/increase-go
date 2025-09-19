// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
)

// SimulationPendingTransactionService contains methods and other services that
// help with interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimulationPendingTransactionService] method instead.
type SimulationPendingTransactionService struct {
	Options []option.RequestOption
}

// NewSimulationPendingTransactionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSimulationPendingTransactionService(opts ...option.RequestOption) (r *SimulationPendingTransactionService) {
	r = &SimulationPendingTransactionService{}
	r.Options = opts
	return
}

// This endpoint simulates immediately releasing an Inbound Funds Hold, which might
// be created as a result of, for example, an ACH debit.
func (r *SimulationPendingTransactionService) ReleaseInboundFundsHold(ctx context.Context, pendingTransactionID string, opts ...option.RequestOption) (res *PendingTransaction, err error) {
	opts = slices.Concat(r.Options, opts)
	if pendingTransactionID == "" {
		err = errors.New("missing required pending_transaction_id parameter")
		return
	}
	path := fmt.Sprintf("simulations/pending_transactions/%s/release_inbound_funds_hold", pendingTransactionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

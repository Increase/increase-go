// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
)

// SimulationWireTransferService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimulationWireTransferService] method instead.
type SimulationWireTransferService struct {
	Options []option.RequestOption
}

// NewSimulationWireTransferService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSimulationWireTransferService(opts ...option.RequestOption) (r *SimulationWireTransferService) {
	r = &SimulationWireTransferService{}
	r.Options = opts
	return
}

// Simulates the reversal of a [Wire Transfer](#wire-transfers) by the Federal
// Reserve due to error conditions. This will also create a
// [Transaction](#transaction) to account for the returned funds. This Wire
// Transfer must first have a `status` of `complete`.
func (r *SimulationWireTransferService) Reverse(ctx context.Context, wireTransferID string, opts ...option.RequestOption) (res *WireTransfer, err error) {
	opts = append(r.Options[:], opts...)
	if wireTransferID == "" {
		err = errors.New("missing required wire_transfer_id parameter")
		return
	}
	path := fmt.Sprintf("simulations/wire_transfers/%s/reverse", wireTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Simulates the submission of a [Wire Transfer](#wire-transfers) to the Federal
// Reserve. This transfer must first have a `status` of `pending_approval` or
// `pending_creating`.
func (r *SimulationWireTransferService) Submit(ctx context.Context, wireTransferID string, opts ...option.RequestOption) (res *WireTransfer, err error) {
	opts = append(r.Options[:], opts...)
	if wireTransferID == "" {
		err = errors.New("missing required wire_transfer_id parameter")
		return
	}
	path := fmt.Sprintf("simulations/wire_transfers/%s/submit", wireTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

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

// SimulationWireDrawdownRequestService contains methods and other services that
// help with interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimulationWireDrawdownRequestService] method instead.
type SimulationWireDrawdownRequestService struct {
	Options []option.RequestOption
}

// NewSimulationWireDrawdownRequestService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSimulationWireDrawdownRequestService(opts ...option.RequestOption) (r *SimulationWireDrawdownRequestService) {
	r = &SimulationWireDrawdownRequestService{}
	r.Options = opts
	return
}

// Simulates a Wire Drawdown Request being refused by the debtor.
func (r *SimulationWireDrawdownRequestService) Refuse(ctx context.Context, wireDrawdownRequestID string, opts ...option.RequestOption) (res *WireDrawdownRequest, err error) {
	opts = append(r.Options[:], opts...)
	if wireDrawdownRequestID == "" {
		err = errors.New("missing required wire_drawdown_request_id parameter")
		return
	}
	path := fmt.Sprintf("simulations/wire_drawdown_requests/%s/refuse", wireDrawdownRequestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Simulates a Wire Drawdown Request being submitted to Fedwire.
func (r *SimulationWireDrawdownRequestService) Submit(ctx context.Context, wireDrawdownRequestID string, opts ...option.RequestOption) (res *WireDrawdownRequest, err error) {
	opts = append(r.Options[:], opts...)
	if wireDrawdownRequestID == "" {
		err = errors.New("missing required wire_drawdown_request_id parameter")
		return
	}
	path := fmt.Sprintf("simulations/wire_drawdown_requests/%s/submit", wireDrawdownRequestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

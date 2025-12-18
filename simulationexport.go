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

// SimulationExportService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimulationExportService] method instead.
type SimulationExportService struct {
	Options []option.RequestOption
}

// NewSimulationExportService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSimulationExportService(opts ...option.RequestOption) (r *SimulationExportService) {
	r = &SimulationExportService{}
	r.Options = opts
	return
}

// Simulates a tax form export being generated.
func (r *SimulationExportService) New(ctx context.Context, body SimulationExportNewParams, opts ...option.RequestOption) (res *Export, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "simulations/exports"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SimulationExportNewParams struct {
	// The identifier of the Account the tax document is for.
	AccountID param.Field[string] `json:"account_id,required"`
}

func (r SimulationExportNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

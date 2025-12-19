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

// Many exports are created by you via POST /exports or in the Dashboard. Some
// exports are created automatically by Increase. For example, tax documents are
// published once a year. In sandbox, you can trigger the arrival of an export that
// would normally only be created automatically via this simulation.
func (r *SimulationExportService) New(ctx context.Context, body SimulationExportNewParams, opts ...option.RequestOption) (res *Export, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "simulations/exports"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SimulationExportNewParams struct {
	// The type of Export to create.
	Category param.Field[SimulationExportNewParamsCategory] `json:"category,required"`
	// Options for the created export. Required if `category` is equal to
	// `form_1099_int`.
	Form1099Int param.Field[SimulationExportNewParamsForm1099Int] `json:"form_1099_int"`
}

func (r SimulationExportNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of Export to create.
type SimulationExportNewParamsCategory string

const (
	SimulationExportNewParamsCategoryForm1099Int SimulationExportNewParamsCategory = "form_1099_int"
)

func (r SimulationExportNewParamsCategory) IsKnown() bool {
	switch r {
	case SimulationExportNewParamsCategoryForm1099Int:
		return true
	}
	return false
}

// Options for the created export. Required if `category` is equal to
// `form_1099_int`.
type SimulationExportNewParamsForm1099Int struct {
	// The identifier of the Account the tax document is for.
	AccountID param.Field[string] `json:"account_id,required"`
}

func (r SimulationExportNewParamsForm1099Int) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

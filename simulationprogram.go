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

// SimulationProgramService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimulationProgramService] method instead.
type SimulationProgramService struct {
	Options []option.RequestOption
}

// NewSimulationProgramService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSimulationProgramService(opts ...option.RequestOption) (r *SimulationProgramService) {
	r = &SimulationProgramService{}
	r.Options = opts
	return
}

// Simulates a [Program](#programs) being created in your group. By default, your
// group has one program called Commercial Banking. Note that when your group
// operates more than one program, `program_id` is a required field when creating
// accounts.
func (r *SimulationProgramService) New(ctx context.Context, body SimulationProgramNewParams, opts ...option.RequestOption) (res *Program, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "simulations/programs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SimulationProgramNewParams struct {
	// The name of the program being added.
	Name param.Field[string] `json:"name,required"`
	// The bank for the program's accounts, defaults to First Internet Bank.
	Bank param.Field[SimulationProgramNewParamsBank] `json:"bank"`
	// The identifier of the Account the Program should be added to is for.
	ReserveAccountID param.Field[string] `json:"reserve_account_id"`
}

func (r SimulationProgramNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The bank for the program's accounts, defaults to First Internet Bank.
type SimulationProgramNewParamsBank string

const (
	SimulationProgramNewParamsBankBlueRidgeBank         SimulationProgramNewParamsBank = "blue_ridge_bank"
	SimulationProgramNewParamsBankCoreBank              SimulationProgramNewParamsBank = "core_bank"
	SimulationProgramNewParamsBankFirstInternetBank     SimulationProgramNewParamsBank = "first_internet_bank"
	SimulationProgramNewParamsBankGlobalInnovationsBank SimulationProgramNewParamsBank = "global_innovations_bank"
	SimulationProgramNewParamsBankGrasshopperBank       SimulationProgramNewParamsBank = "grasshopper_bank"
)

func (r SimulationProgramNewParamsBank) IsKnown() bool {
	switch r {
	case SimulationProgramNewParamsBankBlueRidgeBank, SimulationProgramNewParamsBankCoreBank, SimulationProgramNewParamsBankFirstInternetBank, SimulationProgramNewParamsBankGlobalInnovationsBank, SimulationProgramNewParamsBankGrasshopperBank:
		return true
	}
	return false
}

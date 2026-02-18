// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/Increase/increase-go/internal/apijson"
	"github.com/Increase/increase-go/internal/param"
	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
)

// SimulationCheckDepositService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimulationCheckDepositService] method instead.
type SimulationCheckDepositService struct {
	Options []option.RequestOption
}

// NewSimulationCheckDepositService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSimulationCheckDepositService(opts ...option.RequestOption) (r *SimulationCheckDepositService) {
	r = &SimulationCheckDepositService{}
	r.Options = opts
	return
}

// Simulates the rejection of a [Check Deposit](#check-deposits) by Increase due to
// factors like poor image quality. This Check Deposit must first have a `status`
// of `pending`.
func (r *SimulationCheckDepositService) Reject(ctx context.Context, checkDepositID string, opts ...option.RequestOption) (res *CheckDeposit, err error) {
	opts = slices.Concat(r.Options, opts)
	if checkDepositID == "" {
		err = errors.New("missing required check_deposit_id parameter")
		return
	}
	path := fmt.Sprintf("simulations/check_deposits/%s/reject", checkDepositID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Simulates the return of a [Check Deposit](#check-deposits). This Check Deposit
// must first have a `status` of `submitted`.
func (r *SimulationCheckDepositService) Return(ctx context.Context, checkDepositID string, opts ...option.RequestOption) (res *CheckDeposit, err error) {
	opts = slices.Concat(r.Options, opts)
	if checkDepositID == "" {
		err = errors.New("missing required check_deposit_id parameter")
		return
	}
	path := fmt.Sprintf("simulations/check_deposits/%s/return", checkDepositID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Simulates the submission of a [Check Deposit](#check-deposits) to the Federal
// Reserve. This Check Deposit must first have a `status` of `pending`.
func (r *SimulationCheckDepositService) Submit(ctx context.Context, checkDepositID string, body SimulationCheckDepositSubmitParams, opts ...option.RequestOption) (res *CheckDeposit, err error) {
	opts = slices.Concat(r.Options, opts)
	if checkDepositID == "" {
		err = errors.New("missing required check_deposit_id parameter")
		return
	}
	path := fmt.Sprintf("simulations/check_deposits/%s/submit", checkDepositID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SimulationCheckDepositSubmitParams struct {
	// If set, the simulation will use these values for the check's scanned MICR data.
	Scan param.Field[SimulationCheckDepositSubmitParamsScan] `json:"scan"`
}

func (r SimulationCheckDepositSubmitParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If set, the simulation will use these values for the check's scanned MICR data.
type SimulationCheckDepositSubmitParamsScan struct {
	// The account number to be returned in the check deposit's scan data.
	AccountNumber param.Field[string] `json:"account_number,required"`
	// The routing number to be returned in the check deposit's scan data.
	RoutingNumber param.Field[string] `json:"routing_number,required"`
	// The auxiliary on-us data to be returned in the check deposit's scan data.
	AuxiliaryOnUs param.Field[string] `json:"auxiliary_on_us"`
}

func (r SimulationCheckDepositSubmitParamsScan) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

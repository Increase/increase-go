// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"fmt"
	"net/http"

	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/option"
)

// SimulationCheckDepositService contains methods and other services that help with
// interacting with the increase API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSimulationCheckDepositService]
// method instead.
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
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("simulations/check_deposits/%s/reject", checkDepositID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Simulates the return of a [Check Deposit](#check-deposits). This Check Deposit
// must first have a `status` of `submitted`.
func (r *SimulationCheckDepositService) Return(ctx context.Context, checkDepositID string, opts ...option.RequestOption) (res *CheckDeposit, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("simulations/check_deposits/%s/return", checkDepositID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Simulates the submission of a [Check Deposit](#check-deposits) to the Federal
// Reserve. This Check Deposit must first have a `status` of `pending`.
func (r *SimulationCheckDepositService) Submit(ctx context.Context, checkDepositID string, opts ...option.RequestOption) (res *CheckDeposit, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("simulations/check_deposits/%s/submit", checkDepositID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

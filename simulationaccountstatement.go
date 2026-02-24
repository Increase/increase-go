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

// SimulationAccountStatementService contains methods and other services that help
// with interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimulationAccountStatementService] method instead.
type SimulationAccountStatementService struct {
	Options []option.RequestOption
}

// NewSimulationAccountStatementService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSimulationAccountStatementService(opts ...option.RequestOption) (r *SimulationAccountStatementService) {
	r = &SimulationAccountStatementService{}
	r.Options = opts
	return
}

// Simulates an [Account Statement](#account-statements) being created for an
// account. In production, Account Statements are generated once per month.
func (r *SimulationAccountStatementService) New(ctx context.Context, body SimulationAccountStatementNewParams, opts ...option.RequestOption) (res *AccountStatement, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "simulations/account_statements"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SimulationAccountStatementNewParams struct {
	// The identifier of the Account the statement is for.
	AccountID param.Field[string] `json:"account_id" api:"required"`
}

func (r SimulationAccountStatementNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

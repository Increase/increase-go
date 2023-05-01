package increase

import (
	"context"
	"net/http"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/field"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/option"
)

type SimulationAccountStatementService struct {
	Options []option.RequestOption
}

func NewSimulationAccountStatementService(opts ...option.RequestOption) (r *SimulationAccountStatementService) {
	r = &SimulationAccountStatementService{}
	r.Options = opts
	return
}

// Simulates an [Account Statement](#account-statements) being created for an
// account. In production, Account Statements are generated once per month.
func (r *SimulationAccountStatementService) New(ctx context.Context, body SimulationAccountStatementNewParams, opts ...option.RequestOption) (res *AccountStatement, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/account_statements"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SimulationAccountStatementNewParams struct {
	// The identifier of the Account the statement is for.
	AccountID field.Field[string] `json:"account_id,required"`
}

// MarshalJSON serializes SimulationAccountStatementNewParams into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r SimulationAccountStatementNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

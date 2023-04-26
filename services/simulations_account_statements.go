package services

import (
	"context"
	"net/http"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
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
func (r *SimulationAccountStatementService) New(ctx context.Context, body *requests.SimulationAccountStatementNewParams, opts ...option.RequestOption) (res *responses.AccountStatement, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/account_statements"
	err = option.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

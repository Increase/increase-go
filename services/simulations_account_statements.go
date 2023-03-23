package services

import (
	"context"

	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type SimulationsAccountStatementService struct {
	Options []options.RequestOption
}

func NewSimulationsAccountStatementService(opts ...options.RequestOption) (r *SimulationsAccountStatementService) {
	r = &SimulationsAccountStatementService{}
	r.Options = opts
	return
}

// Simulates an [Account Statement](#account-statements) being created for an
// account. In production, Account Statements are generated once per month.
func (r *SimulationsAccountStatementService) New(ctx context.Context, body *requests.SimulateAnAccountStatementBeingCreatedParameters, opts ...options.RequestOption) (res *responses.AccountStatement, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/account_statements"
	err = options.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

package services

import (
	"context"
	"increase/options"
	"increase/types"
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
func (r *SimulationsAccountStatementService) New(ctx context.Context, body *types.SimulateAnAccountStatementBeingCreatedParameters, opts ...options.RequestOption) (res *types.AccountStatement, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/account_statements"
	err = options.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

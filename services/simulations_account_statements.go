package services

import (
	"context"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type SimulationsAccountStatementService struct {
	Options []option.RequestOption
}

func NewSimulationsAccountStatementService(opts ...option.RequestOption) (r *SimulationsAccountStatementService) {
	r = &SimulationsAccountStatementService{}
	r.Options = opts
	return
}

// Simulates an [Account Statement](#account-statements) being created for an
// account. In production, Account Statements are generated once per month.
func (r *SimulationsAccountStatementService) New(ctx context.Context, body *requests.AccountStatementNewParams, opts ...option.RequestOption) (res *responses.AccountStatement, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/account_statements"
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

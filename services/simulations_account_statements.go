package services

import (
	"bytes"
	"context"
	"fmt"
	"increase/options"
	"increase/types"
	"io"
	"net/url"
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
	b, err := body.MarshalJSON()
	content := io.NopCloser(bytes.NewBuffer(b))
	u, err := url.Parse(fmt.Sprintf("simulations/account_statements"))
	if err != nil {
		return
	}
	cfg := options.NewRequestConfig(ctx, "POST", u, content, opts...)
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return
	}

	return
}

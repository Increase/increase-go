package services

import (
	"context"
	"fmt"
	"increase/options"
	"increase/pagination"
	"increase/types"
	"net/url"
)

type AccountStatementService struct {
	Options []options.RequestOption
}

func NewAccountStatementService(opts ...options.RequestOption) (r *AccountStatementService) {
	r = &AccountStatementService{}
	r.Options = opts
	return
}

// Retrieve an Account Statement
func (r *AccountStatementService) Get(ctx context.Context, account_statement_id string, opts ...options.RequestOption) (res *types.AccountStatement, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]options.RequestOption{options.WithHeader("Content-Type", "")}, opts...)
	u, err := url.Parse(fmt.Sprintf("account_statements/%s", account_statement_id))
	if err != nil {
		return
	}
	cfg, err := options.NewRequestConfig(ctx, "GET", u, nil, opts...)
	if err != nil {
		return
	}
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return
	}

	return
}

// List Account Statements
func (r *AccountStatementService) List(ctx context.Context, query *types.AccountStatementListParams, opts ...options.RequestOption) (res *types.AccountStatementsPage, err error) {
	u, err := url.Parse(fmt.Sprintf("account_statements"))
	if err != nil {
		return
	}
	opts = append(r.Options, opts...)
	cfg, err := options.NewRequestConfig(ctx, "GET", u, query, opts...)
	if err != nil {
		return
	}
	res = &types.AccountStatementsPage{
		Page: &pagination.Page[types.AccountStatement]{
			Config:  *cfg,
			Options: opts,
		},
	}
	err = res.Fire()
	return
}

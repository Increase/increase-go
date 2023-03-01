package services

import (
	"context"
	"fmt"
	"increase/options"
	"increase/pagination"
	"increase/types"
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
	path := fmt.Sprintf("account_statements/%s", account_statement_id)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// List Account Statements
func (r *AccountStatementService) List(ctx context.Context, query *types.AccountStatementListParams, opts ...options.RequestOption) (res *types.AccountStatementsPage, err error) {
	opts = append(r.Options, opts...)
	path := "account_statements"
	cfg, err := options.NewRequestConfig(ctx, "GET", path, query, nil, opts...)
	if err != nil {
		return
	}
	res = &types.AccountStatementsPage{
		Page: &pagination.Page[types.AccountStatement]{
			Config:  *cfg,
			Options: opts,
		},
	}
	return res, res.Fire()
}

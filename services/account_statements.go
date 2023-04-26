package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type AccountStatementService struct {
	Options []option.RequestOption
}

func NewAccountStatementService(opts ...option.RequestOption) (r *AccountStatementService) {
	r = &AccountStatementService{}
	r.Options = opts
	return
}

// Retrieve an Account Statement
func (r *AccountStatementService) Get(ctx context.Context, account_statement_id string, opts ...option.RequestOption) (res *responses.AccountStatement, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("account_statements/%s", account_statement_id)
	err = option.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Account Statements
func (r *AccountStatementService) List(ctx context.Context, query *requests.AccountStatementListParams, opts ...option.RequestOption) (res *responses.Page[responses.AccountStatement], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "account_statements"
	cfg, err := option.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List Account Statements
func (r *AccountStatementService) ListAutoPager(ctx context.Context, query *requests.AccountStatementListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.AccountStatement] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}

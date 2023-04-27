package services

import (
	"context"
	"net/http"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type BookkeepingAccountService struct {
	Options []option.RequestOption
}

func NewBookkeepingAccountService(opts ...option.RequestOption) (r *BookkeepingAccountService) {
	r = &BookkeepingAccountService{}
	r.Options = opts
	return
}

// Create a Bookkeeping Account
func (r *BookkeepingAccountService) New(ctx context.Context, body requests.BookkeepingAccountNewParams, opts ...option.RequestOption) (res *responses.BookkeepingAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := "bookkeeping_accounts"
	err = option.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List Bookkeeping Accounts
func (r *BookkeepingAccountService) List(ctx context.Context, query requests.BookkeepingAccountListParams, opts ...option.RequestOption) (res *responses.Page[responses.BookkeepingAccount], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "bookkeeping_accounts"
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

// List Bookkeeping Accounts
func (r *BookkeepingAccountService) ListAutoPaging(ctx context.Context, query requests.BookkeepingAccountListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.BookkeepingAccount] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}

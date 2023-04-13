package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type AccountService struct {
	Options []option.RequestOption
}

func NewAccountService(opts ...option.RequestOption) (r *AccountService) {
	r = &AccountService{}
	r.Options = opts
	return
}

// Create an Account
func (r *AccountService) New(ctx context.Context, body *requests.AccountNewParams, opts ...option.RequestOption) (res *responses.Account, err error) {
	opts = append(r.Options[:], opts...)
	path := "accounts"
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Retrieve an Account
func (r *AccountService) Get(ctx context.Context, account_id string, opts ...option.RequestOption) (res *responses.Account, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s", account_id)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// Update an Account
func (r *AccountService) Update(ctx context.Context, account_id string, body *requests.AccountUpdateParams, opts ...option.RequestOption) (res *responses.Account, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s", account_id)
	err = option.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// List Accounts
func (r *AccountService) List(ctx context.Context, query *requests.AccountListParams, opts ...option.RequestOption) (res *responses.Page[responses.Account], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "accounts"
	cfg, err := option.NewRequestConfig(ctx, "GET", path, query, &res, opts...)
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

// List Accounts
func (r *AccountService) ListAutoPager(ctx context.Context, query *requests.AccountListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.Account] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Close an Account
func (r *AccountService) Close(ctx context.Context, account_id string, opts ...option.RequestOption) (res *responses.Account, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/close", account_id)
	err = option.ExecuteNewRequest(ctx, "POST", path, nil, &res, opts...)
	return
}

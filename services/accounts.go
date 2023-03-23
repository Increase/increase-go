package services

import (
	"context"
	"fmt"

	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/pagination"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type AccountService struct {
	Options []options.RequestOption
}

func NewAccountService(opts ...options.RequestOption) (r *AccountService) {
	r = &AccountService{}
	r.Options = opts
	return
}

// Create an Account
func (r *AccountService) New(ctx context.Context, body *requests.CreateAnAccountParameters, opts ...options.RequestOption) (res *responses.Account, err error) {
	opts = append(r.Options[:], opts...)
	path := "accounts"
	err = options.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Retrieve an Account
func (r *AccountService) Get(ctx context.Context, account_id string, opts ...options.RequestOption) (res *responses.Account, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s", account_id)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// Update an Account
func (r *AccountService) Update(ctx context.Context, account_id string, body *requests.UpdateAnAccountParameters, opts ...options.RequestOption) (res *responses.Account, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s", account_id)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// List Accounts
func (r *AccountService) List(ctx context.Context, query *requests.AccountListParams, opts ...options.RequestOption) (res *responses.AccountsPage, err error) {
	opts = append(r.Options, opts...)
	path := "accounts"
	cfg, err := options.NewRequestConfig(ctx, "GET", path, query, nil, opts...)
	if err != nil {
		return
	}
	res = &responses.AccountsPage{
		Page: &pagination.Page[responses.Account]{
			Config:  *cfg,
			Options: opts,
		},
	}
	return res, res.Fire()
}

// Close an Account
func (r *AccountService) Close(ctx context.Context, account_id string, opts ...options.RequestOption) (res *responses.Account, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/close", account_id)
	err = options.ExecuteNewRequest(ctx, "POST", path, nil, &res, opts...)
	return
}

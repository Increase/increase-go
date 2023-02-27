package services

import (
	"context"
	"fmt"
	"increase/options"
	"increase/pagination"
	"increase/types"
	"net/url"
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
func (r *AccountService) New(ctx context.Context, body *types.CreateAnAccountParameters, opts ...options.RequestOption) (res *types.Account, err error) {
	opts = append(r.Options[:], opts...)
	u, err := url.Parse(fmt.Sprintf("accounts"))
	if err != nil {
		return
	}
	cfg, err := options.NewRequestConfig(ctx, "POST", u, body, opts...)
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

// Retrieve an Account
func (r *AccountService) Get(ctx context.Context, account_id string, opts ...options.RequestOption) (res *types.Account, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]options.RequestOption{options.WithHeader("Content-Type", "")}, opts...)
	u, err := url.Parse(fmt.Sprintf("accounts/%s", account_id))
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

// Update an Account
func (r *AccountService) Update(ctx context.Context, account_id string, body *types.UpdateAnAccountParameters, opts ...options.RequestOption) (res *types.Account, err error) {
	opts = append(r.Options[:], opts...)
	u, err := url.Parse(fmt.Sprintf("accounts/%s", account_id))
	if err != nil {
		return
	}
	cfg, err := options.NewRequestConfig(ctx, "PATCH", u, body, opts...)
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

// List Accounts
func (r *AccountService) List(ctx context.Context, query *types.AccountListParams, opts ...options.RequestOption) (res *types.AccountsPage, err error) {
	u, err := url.Parse(fmt.Sprintf("accounts"))
	if err != nil {
		return
	}
	opts = append(r.Options, opts...)
	cfg, err := options.NewRequestConfig(ctx, "GET", u, query, opts...)
	if err != nil {
		return
	}
	res = &types.AccountsPage{
		Page: &pagination.Page[types.Account]{
			Config:  *cfg,
			Options: opts,
		},
	}
	err = res.Fire()
	return
}

// Close an Account
func (r *AccountService) Close(ctx context.Context, account_id string, opts ...options.RequestOption) (res *types.Account, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]options.RequestOption{options.WithHeader("Content-Type", "")}, opts...)
	u, err := url.Parse(fmt.Sprintf("accounts/%s/close", account_id))
	if err != nil {
		return
	}
	cfg, err := options.NewRequestConfig(ctx, "POST", u, nil, opts...)
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

package services

import (
	"context"
	"fmt"
	"increase/options"
	"increase/pagination"
	"increase/types"
	"net/url"
)

type ExternalAccountService struct {
	Options []options.RequestOption
}

func NewExternalAccountService(opts ...options.RequestOption) (r *ExternalAccountService) {
	r = &ExternalAccountService{}
	r.Options = opts
	return
}

// Create an External Account
func (r *ExternalAccountService) New(ctx context.Context, body *types.CreateAnExternalAccountParameters, opts ...options.RequestOption) (res *types.ExternalAccount, err error) {
	opts = append(r.Options[:], opts...)
	u, err := url.Parse(fmt.Sprintf("external_accounts"))
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

// Retrieve an External Account
func (r *ExternalAccountService) Get(ctx context.Context, external_account_id string, opts ...options.RequestOption) (res *types.ExternalAccount, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]options.RequestOption{options.WithHeader("Content-Type", "")}, opts...)
	u, err := url.Parse(fmt.Sprintf("external_accounts/%s", external_account_id))
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

// Update an External Account
func (r *ExternalAccountService) Update(ctx context.Context, external_account_id string, body *types.UpdateAnExternalAccountParameters, opts ...options.RequestOption) (res *types.ExternalAccount, err error) {
	opts = append(r.Options[:], opts...)
	u, err := url.Parse(fmt.Sprintf("external_accounts/%s", external_account_id))
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

// List External Accounts
func (r *ExternalAccountService) List(ctx context.Context, query *types.ExternalAccountListParams, opts ...options.RequestOption) (res *types.ExternalAccountsPage, err error) {
	u, err := url.Parse(fmt.Sprintf("external_accounts"))
	if err != nil {
		return
	}
	opts = append(r.Options, opts...)
	cfg, err := options.NewRequestConfig(ctx, "GET", u, query, opts...)
	if err != nil {
		return
	}
	res = &types.ExternalAccountsPage{
		Page: &pagination.Page[types.ExternalAccount]{
			Config:  *cfg,
			Options: opts,
		},
	}
	err = res.Fire()
	return
}

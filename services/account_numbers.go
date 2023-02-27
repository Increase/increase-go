package services

import (
	"context"
	"fmt"
	"increase/options"
	"increase/pagination"
	"increase/types"
	"net/url"
)

type AccountNumberService struct {
	Options []options.RequestOption
}

func NewAccountNumberService(opts ...options.RequestOption) (r *AccountNumberService) {
	r = &AccountNumberService{}
	r.Options = opts
	return
}

// Create an Account Number
func (r *AccountNumberService) New(ctx context.Context, body *types.CreateAnAccountNumberParameters, opts ...options.RequestOption) (res *types.AccountNumber, err error) {
	opts = append(r.Options[:], opts...)
	u, err := url.Parse(fmt.Sprintf("account_numbers"))
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

// Retrieve an Account Number
func (r *AccountNumberService) Get(ctx context.Context, account_number_id string, opts ...options.RequestOption) (res *types.AccountNumber, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]options.RequestOption{options.WithHeader("Content-Type", "")}, opts...)
	u, err := url.Parse(fmt.Sprintf("account_numbers/%s", account_number_id))
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

// Update an Account Number
func (r *AccountNumberService) Update(ctx context.Context, account_number_id string, body *types.UpdateAnAccountNumberParameters, opts ...options.RequestOption) (res *types.AccountNumber, err error) {
	opts = append(r.Options[:], opts...)
	u, err := url.Parse(fmt.Sprintf("account_numbers/%s", account_number_id))
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

// List Account Numbers
func (r *AccountNumberService) List(ctx context.Context, query *types.AccountNumberListParams, opts ...options.RequestOption) (res *types.AccountNumbersPage, err error) {
	u, err := url.Parse(fmt.Sprintf("account_numbers"))
	if err != nil {
		return
	}
	opts = append(r.Options, opts...)
	cfg, err := options.NewRequestConfig(ctx, "GET", u, query, opts...)
	if err != nil {
		return
	}
	res = &types.AccountNumbersPage{
		Page: &pagination.Page[types.AccountNumber]{
			Config:  *cfg,
			Options: opts,
		},
	}
	err = res.Fire()
	return
}

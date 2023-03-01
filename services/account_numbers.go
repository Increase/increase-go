package services

import (
	"context"
	"fmt"
	"increase/options"
	"increase/pagination"
	"increase/types"
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
	path := "account_numbers"
	err = options.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Retrieve an Account Number
func (r *AccountNumberService) Get(ctx context.Context, account_number_id string, opts ...options.RequestOption) (res *types.AccountNumber, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("account_numbers/%s", account_number_id)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// Update an Account Number
func (r *AccountNumberService) Update(ctx context.Context, account_number_id string, body *types.UpdateAnAccountNumberParameters, opts ...options.RequestOption) (res *types.AccountNumber, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("account_numbers/%s", account_number_id)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// List Account Numbers
func (r *AccountNumberService) List(ctx context.Context, query *types.AccountNumberListParams, opts ...options.RequestOption) (res *types.AccountNumbersPage, err error) {
	opts = append(r.Options, opts...)
	path := "account_numbers"
	cfg, err := options.NewRequestConfig(ctx, "GET", path, query, nil, opts...)
	if err != nil {
		return
	}
	res = &types.AccountNumbersPage{
		Page: &pagination.Page[types.AccountNumber]{
			Config:  *cfg,
			Options: opts,
		},
	}
	return res, res.Fire()
}

package services

import (
	"context"
	"fmt"

	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/pagination"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
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
func (r *AccountNumberService) New(ctx context.Context, body *requests.CreateAnAccountNumberParameters, opts ...options.RequestOption) (res *responses.AccountNumber, err error) {
	opts = append(r.Options[:], opts...)
	path := "account_numbers"
	err = options.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Retrieve an Account Number
func (r *AccountNumberService) Get(ctx context.Context, account_number_id string, opts ...options.RequestOption) (res *responses.AccountNumber, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("account_numbers/%s", account_number_id)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// Update an Account Number
func (r *AccountNumberService) Update(ctx context.Context, account_number_id string, body *requests.UpdateAnAccountNumberParameters, opts ...options.RequestOption) (res *responses.AccountNumber, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("account_numbers/%s", account_number_id)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// List Account Numbers
func (r *AccountNumberService) List(ctx context.Context, query *requests.AccountNumberListParams, opts ...options.RequestOption) (res *responses.AccountNumbersPage, err error) {
	opts = append(r.Options, opts...)
	path := "account_numbers"
	cfg, err := options.NewRequestConfig(ctx, "GET", path, query, nil, opts...)
	if err != nil {
		return
	}
	res = &responses.AccountNumbersPage{
		Page: &pagination.Page[responses.AccountNumber]{
			Config:  *cfg,
			Options: opts,
		},
	}
	return res, res.Fire()
}

package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type AccountNumberService struct {
	Options []option.RequestOption
}

func NewAccountNumberService(opts ...option.RequestOption) (r *AccountNumberService) {
	r = &AccountNumberService{}
	r.Options = opts
	return
}

// Create an Account Number
func (r *AccountNumberService) New(ctx context.Context, body *requests.AccountNumberNewParams, opts ...option.RequestOption) (res *responses.AccountNumber, err error) {
	opts = append(r.Options[:], opts...)
	path := "account_numbers"
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Retrieve an Account Number
func (r *AccountNumberService) Get(ctx context.Context, account_number_id string, opts ...option.RequestOption) (res *responses.AccountNumber, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("account_numbers/%s", account_number_id)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// Update an Account Number
func (r *AccountNumberService) Update(ctx context.Context, account_number_id string, body *requests.AccountNumberUpdateParams, opts ...option.RequestOption) (res *responses.AccountNumber, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("account_numbers/%s", account_number_id)
	err = option.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// List Account Numbers
func (r *AccountNumberService) List(ctx context.Context, query *requests.AccountNumberListParams, opts ...option.RequestOption) (res *responses.Page[responses.AccountNumber], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "account_numbers"
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

// List Account Numbers
func (r *AccountNumberService) ListAutoPager(ctx context.Context, query *requests.AccountNumberListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.AccountNumber] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}

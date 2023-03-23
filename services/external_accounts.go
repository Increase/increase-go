package services

import (
	"context"
	"fmt"

	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/pagination"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
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
func (r *ExternalAccountService) New(ctx context.Context, body *requests.CreateAnExternalAccountParameters, opts ...options.RequestOption) (res *responses.ExternalAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := "external_accounts"
	err = options.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Retrieve an External Account
func (r *ExternalAccountService) Get(ctx context.Context, external_account_id string, opts ...options.RequestOption) (res *responses.ExternalAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("external_accounts/%s", external_account_id)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// Update an External Account
func (r *ExternalAccountService) Update(ctx context.Context, external_account_id string, body *requests.UpdateAnExternalAccountParameters, opts ...options.RequestOption) (res *responses.ExternalAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("external_accounts/%s", external_account_id)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// List External Accounts
func (r *ExternalAccountService) List(ctx context.Context, query *requests.ExternalAccountListParams, opts ...options.RequestOption) (res *responses.ExternalAccountsPage, err error) {
	opts = append(r.Options, opts...)
	path := "external_accounts"
	cfg, err := options.NewRequestConfig(ctx, "GET", path, query, nil, opts...)
	if err != nil {
		return
	}
	res = &responses.ExternalAccountsPage{
		Page: &pagination.Page[responses.ExternalAccount]{
			Config:  *cfg,
			Options: opts,
		},
	}
	return res, res.Fire()
}

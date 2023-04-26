package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type ExternalAccountService struct {
	Options []option.RequestOption
}

func NewExternalAccountService(opts ...option.RequestOption) (r *ExternalAccountService) {
	r = &ExternalAccountService{}
	r.Options = opts
	return
}

// Create an External Account
func (r *ExternalAccountService) New(ctx context.Context, body *requests.ExternalAccountNewParams, opts ...option.RequestOption) (res *responses.ExternalAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := "external_accounts"
	err = option.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve an External Account
func (r *ExternalAccountService) Get(ctx context.Context, external_account_id string, opts ...option.RequestOption) (res *responses.ExternalAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("external_accounts/%s", external_account_id)
	err = option.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an External Account
func (r *ExternalAccountService) Update(ctx context.Context, external_account_id string, body *requests.ExternalAccountUpdateParams, opts ...option.RequestOption) (res *responses.ExternalAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("external_accounts/%s", external_account_id)
	err = option.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List External Accounts
func (r *ExternalAccountService) List(ctx context.Context, query *requests.ExternalAccountListParams, opts ...option.RequestOption) (res *responses.Page[responses.ExternalAccount], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "external_accounts"
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

// List External Accounts
func (r *ExternalAccountService) ListAutoPager(ctx context.Context, query *requests.ExternalAccountListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.ExternalAccount] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}

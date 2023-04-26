package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type CheckDepositService struct {
	Options []option.RequestOption
}

func NewCheckDepositService(opts ...option.RequestOption) (r *CheckDepositService) {
	r = &CheckDepositService{}
	r.Options = opts
	return
}

// Create a Check Deposit
func (r *CheckDepositService) New(ctx context.Context, body *requests.CheckDepositNewParams, opts ...option.RequestOption) (res *responses.CheckDeposit, err error) {
	opts = append(r.Options[:], opts...)
	path := "check_deposits"
	err = option.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a Check Deposit
func (r *CheckDepositService) Get(ctx context.Context, check_deposit_id string, opts ...option.RequestOption) (res *responses.CheckDeposit, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("check_deposits/%s", check_deposit_id)
	err = option.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Check Deposits
func (r *CheckDepositService) List(ctx context.Context, query *requests.CheckDepositListParams, opts ...option.RequestOption) (res *responses.Page[responses.CheckDeposit], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "check_deposits"
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

// List Check Deposits
func (r *CheckDepositService) ListAutoPager(ctx context.Context, query *requests.CheckDepositListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.CheckDeposit] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}

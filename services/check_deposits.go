package services

import (
	"context"
	"fmt"
	"increase/options"
	"increase/pagination"
	"increase/types"
	"net/url"
)

type CheckDepositService struct {
	Options []options.RequestOption
}

func NewCheckDepositService(opts ...options.RequestOption) (r *CheckDepositService) {
	r = &CheckDepositService{}
	r.Options = opts
	return
}

// Create a Check Deposit
func (r *CheckDepositService) New(ctx context.Context, body *types.CreateACheckDepositParameters, opts ...options.RequestOption) (res *types.CheckDeposit, err error) {
	opts = append(r.Options, opts...)
	u, err := url.Parse(fmt.Sprintf("check_deposits"))
	if err != nil {
		return
	}
	cfg := options.NewRequestConfig(ctx, "POST", u, opts...)
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return
	}

	return
}

// Retrieve a Check Deposit
func (r *CheckDepositService) Get(ctx context.Context, check_deposit_id string, opts ...options.RequestOption) (res *types.CheckDeposit, err error) {
	opts = append(r.Options, opts...)
	u, err := url.Parse(fmt.Sprintf("check_deposits/%s", check_deposit_id))
	if err != nil {
		return
	}
	cfg := options.NewRequestConfig(ctx, "GET", u, opts...)
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return
	}

	return
}

// List Check Deposits
func (r *CheckDepositService) List(ctx context.Context, query *types.CheckDepositListParams, opts ...options.RequestOption) (res *types.CheckDepositsPage, err error) {
	u, err := url.Parse(fmt.Sprintf("check_deposits"))
	if err != nil {
		return
	}
	opts = append(r.Options, opts...)
	cfg := options.NewRequestConfig(ctx, "GET", u, opts...)
	res = &types.CheckDepositsPage{
		Page: &pagination.Page[types.CheckDeposit]{
			Config:  *cfg,
			Options: opts,
		},
	}
	err = res.Fire()
	return
}

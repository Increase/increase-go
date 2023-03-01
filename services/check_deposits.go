package services

import (
	"context"
	"fmt"
	"increase/options"
	"increase/pagination"
	"increase/types"
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
	opts = append(r.Options[:], opts...)
	path := "check_deposits"
	err = options.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Retrieve a Check Deposit
func (r *CheckDepositService) Get(ctx context.Context, check_deposit_id string, opts ...options.RequestOption) (res *types.CheckDeposit, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("check_deposits/%s", check_deposit_id)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// List Check Deposits
func (r *CheckDepositService) List(ctx context.Context, query *types.CheckDepositListParams, opts ...options.RequestOption) (res *types.CheckDepositsPage, err error) {
	opts = append(r.Options, opts...)
	path := "check_deposits"
	cfg, err := options.NewRequestConfig(ctx, "GET", path, query, nil, opts...)
	if err != nil {
		return
	}
	res = &types.CheckDepositsPage{
		Page: &pagination.Page[types.CheckDeposit]{
			Config:  *cfg,
			Options: opts,
		},
	}
	return res, res.Fire()
}

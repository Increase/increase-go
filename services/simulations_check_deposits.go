package services

import (
	"context"
	"fmt"
	"increase/options"
	"increase/types"
	"net/url"
)

type SimulationsCheckDepositService struct {
	Options []options.RequestOption
}

func NewSimulationsCheckDepositService(opts ...options.RequestOption) (r *SimulationsCheckDepositService) {
	r = &SimulationsCheckDepositService{}
	r.Options = opts
	return
}

// Simulates the rejection of a [Check Deposit](#check-deposits) by Increase due to
// factors like poor image quality. This Check Deposit must first have a `status`
// of `pending`.
func (r *SimulationsCheckDepositService) Reject(ctx context.Context, check_deposit_id string, opts ...options.RequestOption) (res *types.CheckDeposit, err error) {
	opts = append(r.Options[:], opts...)
	u, err := url.Parse(fmt.Sprintf("simulations/check_deposits/%s/reject", check_deposit_id))
	if err != nil {
		return
	}
	cfg, err := options.NewRequestConfig(ctx, "POST", u, nil, opts...)
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

// Simulates the submission of a [Check Deposit](#check-deposits) to the Federal
// Reserve. This Check Deposit must first have a `status` of `pending`.
func (r *SimulationsCheckDepositService) Submit(ctx context.Context, check_deposit_id string, opts ...options.RequestOption) (res *types.CheckDeposit, err error) {
	opts = append(r.Options[:], opts...)
	u, err := url.Parse(fmt.Sprintf("simulations/check_deposits/%s/submit", check_deposit_id))
	if err != nil {
		return
	}
	cfg, err := options.NewRequestConfig(ctx, "POST", u, nil, opts...)
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

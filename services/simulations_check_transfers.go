package services

import (
	"context"
	"fmt"
	"increase/options"
	"increase/types"
	"net/url"
)

type SimulationsCheckTransferService struct {
	Options []options.RequestOption
}

func NewSimulationsCheckTransferService(opts ...options.RequestOption) (r *SimulationsCheckTransferService) {
	r = &SimulationsCheckTransferService{}
	r.Options = opts
	return
}

// Simulates a [Check Transfer](#check-transfers) being deposited at a bank. This
// transfer must first have a `status` of `mailed`.
func (r *SimulationsCheckTransferService) Deposit(ctx context.Context, check_transfer_id string, opts ...options.RequestOption) (res *types.CheckTransfer, err error) {
	opts = append(r.Options, opts...)
	u, err := url.Parse(fmt.Sprintf("simulations/check_transfers/%s/deposit", check_transfer_id))
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

// Simulates the mailing of a [Check Transfer](#check-transfers), which happens
// once per weekday in production but can be sped up in sandbox. This transfer must
// first have a `status` of `pending_approval` or `pending_submission`.
func (r *SimulationsCheckTransferService) Mail(ctx context.Context, check_transfer_id string, opts ...options.RequestOption) (res *types.CheckTransfer, err error) {
	opts = append(r.Options, opts...)
	u, err := url.Parse(fmt.Sprintf("simulations/check_transfers/%s/mail", check_transfer_id))
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

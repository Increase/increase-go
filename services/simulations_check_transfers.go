package services

import (
	"context"
	"fmt"

	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/types"
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
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("simulations/check_transfers/%s/deposit", check_transfer_id)
	err = options.ExecuteNewRequest(ctx, "POST", path, nil, &res, opts...)
	return
}

// Simulates the mailing of a [Check Transfer](#check-transfers), which happens
// once per weekday in production but can be sped up in sandbox. This transfer must
// first have a `status` of `pending_approval` or `pending_submission`.
func (r *SimulationsCheckTransferService) Mail(ctx context.Context, check_transfer_id string, opts ...options.RequestOption) (res *types.CheckTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("simulations/check_transfers/%s/mail", check_transfer_id)
	err = options.ExecuteNewRequest(ctx, "POST", path, nil, &res, opts...)
	return
}

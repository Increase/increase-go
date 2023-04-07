package services

import (
	"context"
	"fmt"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type SimulationsCheckTransferService struct {
	Options []option.RequestOption
}

func NewSimulationsCheckTransferService(opts ...option.RequestOption) (r *SimulationsCheckTransferService) {
	r = &SimulationsCheckTransferService{}
	r.Options = opts
	return
}

// Simulates a [Check Transfer](#check-transfers) being deposited at a bank. This
// transfer must first have a `status` of `mailed`.
func (r *SimulationsCheckTransferService) Deposit(ctx context.Context, check_transfer_id string, opts ...option.RequestOption) (res *responses.CheckTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("simulations/check_transfers/%s/deposit", check_transfer_id)
	err = option.ExecuteNewRequest(ctx, "POST", path, nil, &res, opts...)
	return
}

// Simulates the mailing of a [Check Transfer](#check-transfers), which happens
// once per weekday in production but can be sped up in sandbox. This transfer must
// first have a `status` of `pending_approval` or `pending_submission`.
func (r *SimulationsCheckTransferService) Mail(ctx context.Context, check_transfer_id string, opts ...option.RequestOption) (res *responses.CheckTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("simulations/check_transfers/%s/mail", check_transfer_id)
	err = option.ExecuteNewRequest(ctx, "POST", path, nil, &res, opts...)
	return
}

// Simulates a [Check Transfer](#check-transfers) being returned via USPS to
// Increase. This transfer must first have a `status` of `mailed`.
func (r *SimulationsCheckTransferService) Return(ctx context.Context, check_transfer_id string, body *requests.CheckTransferReturnParams, opts ...option.RequestOption) (res *responses.CheckTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("simulations/check_transfers/%s/return", check_transfer_id)
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

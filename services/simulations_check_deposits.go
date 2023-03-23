package services

import (
	"context"
	"fmt"

	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/responses"
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
func (r *SimulationsCheckDepositService) Reject(ctx context.Context, check_deposit_id string, opts ...options.RequestOption) (res *responses.CheckDeposit, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("simulations/check_deposits/%s/reject", check_deposit_id)
	err = options.ExecuteNewRequest(ctx, "POST", path, nil, &res, opts...)
	return
}

// Simulates the return of a [Check Deposit](#check-deposits). This Check Deposit
// must first have a `status` of `submitted`.
func (r *SimulationsCheckDepositService) Return(ctx context.Context, check_deposit_id string, opts ...options.RequestOption) (res *responses.CheckDeposit, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("simulations/check_deposits/%s/return", check_deposit_id)
	err = options.ExecuteNewRequest(ctx, "POST", path, nil, &res, opts...)
	return
}

// Simulates the submission of a [Check Deposit](#check-deposits) to the Federal
// Reserve. This Check Deposit must first have a `status` of `pending`.
func (r *SimulationsCheckDepositService) Submit(ctx context.Context, check_deposit_id string, opts ...options.RequestOption) (res *responses.CheckDeposit, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("simulations/check_deposits/%s/submit", check_deposit_id)
	err = options.ExecuteNewRequest(ctx, "POST", path, nil, &res, opts...)
	return
}

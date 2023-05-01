package increase

import (
	"context"
	"fmt"
	"net/http"

	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/option"
)

type SimulationCheckDepositService struct {
	Options []option.RequestOption
}

func NewSimulationCheckDepositService(opts ...option.RequestOption) (r *SimulationCheckDepositService) {
	r = &SimulationCheckDepositService{}
	r.Options = opts
	return
}

// Simulates the rejection of a [Check Deposit](#check-deposits) by Increase due to
// factors like poor image quality. This Check Deposit must first have a `status`
// of `pending`.
func (r *SimulationCheckDepositService) Reject(ctx context.Context, check_deposit_id string, opts ...option.RequestOption) (res *CheckDeposit, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("simulations/check_deposits/%s/reject", check_deposit_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Simulates the return of a [Check Deposit](#check-deposits). This Check Deposit
// must first have a `status` of `submitted`.
func (r *SimulationCheckDepositService) Return(ctx context.Context, check_deposit_id string, opts ...option.RequestOption) (res *CheckDeposit, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("simulations/check_deposits/%s/return", check_deposit_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Simulates the submission of a [Check Deposit](#check-deposits) to the Federal
// Reserve. This Check Deposit must first have a `status` of `pending`.
func (r *SimulationCheckDepositService) Submit(ctx context.Context, check_deposit_id string, opts ...option.RequestOption) (res *CheckDeposit, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("simulations/check_deposits/%s/submit", check_deposit_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

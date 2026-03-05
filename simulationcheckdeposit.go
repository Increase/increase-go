// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/Increase/increase-go/internal/apijson"
	"github.com/Increase/increase-go/internal/param"
	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
)

// SimulationCheckDepositService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimulationCheckDepositService] method instead.
type SimulationCheckDepositService struct {
	Options []option.RequestOption
}

// NewSimulationCheckDepositService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSimulationCheckDepositService(opts ...option.RequestOption) (r *SimulationCheckDepositService) {
	r = &SimulationCheckDepositService{}
	r.Options = opts
	return
}

// Simulates the creation of a
// [Check Deposit Adjustment](#check-deposit-adjustments) on a
// [Check Deposit](#check-deposits). This Check Deposit must first have a `status`
// of `submitted`.
func (r *SimulationCheckDepositService) Adjustment(ctx context.Context, checkDepositID string, body SimulationCheckDepositAdjustmentParams, opts ...option.RequestOption) (res *CheckDeposit, err error) {
	opts = slices.Concat(r.Options, opts)
	if checkDepositID == "" {
		err = errors.New("missing required check_deposit_id parameter")
		return
	}
	path := fmt.Sprintf("simulations/check_deposits/%s/adjustment", checkDepositID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Simulates the rejection of a [Check Deposit](#check-deposits) by Increase due to
// factors like poor image quality. This Check Deposit must first have a `status`
// of `pending`.
func (r *SimulationCheckDepositService) Reject(ctx context.Context, checkDepositID string, opts ...option.RequestOption) (res *CheckDeposit, err error) {
	opts = slices.Concat(r.Options, opts)
	if checkDepositID == "" {
		err = errors.New("missing required check_deposit_id parameter")
		return
	}
	path := fmt.Sprintf("simulations/check_deposits/%s/reject", checkDepositID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Simulates the return of a [Check Deposit](#check-deposits). This Check Deposit
// must first have a `status` of `submitted`.
func (r *SimulationCheckDepositService) Return(ctx context.Context, checkDepositID string, opts ...option.RequestOption) (res *CheckDeposit, err error) {
	opts = slices.Concat(r.Options, opts)
	if checkDepositID == "" {
		err = errors.New("missing required check_deposit_id parameter")
		return
	}
	path := fmt.Sprintf("simulations/check_deposits/%s/return", checkDepositID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Simulates the submission of a [Check Deposit](#check-deposits) to the Federal
// Reserve. This Check Deposit must first have a `status` of `pending`.
func (r *SimulationCheckDepositService) Submit(ctx context.Context, checkDepositID string, body SimulationCheckDepositSubmitParams, opts ...option.RequestOption) (res *CheckDeposit, err error) {
	opts = slices.Concat(r.Options, opts)
	if checkDepositID == "" {
		err = errors.New("missing required check_deposit_id parameter")
		return
	}
	path := fmt.Sprintf("simulations/check_deposits/%s/submit", checkDepositID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SimulationCheckDepositAdjustmentParams struct {
	// The adjustment amount in the minor unit of the Check Deposit's currency (e.g.,
	// cents). A negative amount means that the funds are being clawed back by the
	// other bank and is a debit to your account. Defaults to the negative of the Check
	// Deposit amount.
	Amount param.Field[int64] `json:"amount"`
	// The reason for the adjustment. Defaults to `non_conforming_item`, which is often
	// used for a low quality image that the recipient wasn't able to handle.
	Reason param.Field[SimulationCheckDepositAdjustmentParamsReason] `json:"reason"`
}

func (r SimulationCheckDepositAdjustmentParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The reason for the adjustment. Defaults to `non_conforming_item`, which is often
// used for a low quality image that the recipient wasn't able to handle.
type SimulationCheckDepositAdjustmentParamsReason string

const (
	SimulationCheckDepositAdjustmentParamsReasonLateReturn        SimulationCheckDepositAdjustmentParamsReason = "late_return"
	SimulationCheckDepositAdjustmentParamsReasonWrongPayeeCredit  SimulationCheckDepositAdjustmentParamsReason = "wrong_payee_credit"
	SimulationCheckDepositAdjustmentParamsReasonAdjustedAmount    SimulationCheckDepositAdjustmentParamsReason = "adjusted_amount"
	SimulationCheckDepositAdjustmentParamsReasonNonConformingItem SimulationCheckDepositAdjustmentParamsReason = "non_conforming_item"
	SimulationCheckDepositAdjustmentParamsReasonPaid              SimulationCheckDepositAdjustmentParamsReason = "paid"
)

func (r SimulationCheckDepositAdjustmentParamsReason) IsKnown() bool {
	switch r {
	case SimulationCheckDepositAdjustmentParamsReasonLateReturn, SimulationCheckDepositAdjustmentParamsReasonWrongPayeeCredit, SimulationCheckDepositAdjustmentParamsReasonAdjustedAmount, SimulationCheckDepositAdjustmentParamsReasonNonConformingItem, SimulationCheckDepositAdjustmentParamsReasonPaid:
		return true
	}
	return false
}

type SimulationCheckDepositSubmitParams struct {
	// If set, the simulation will use these values for the check's scanned MICR data.
	Scan param.Field[SimulationCheckDepositSubmitParamsScan] `json:"scan"`
}

func (r SimulationCheckDepositSubmitParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If set, the simulation will use these values for the check's scanned MICR data.
type SimulationCheckDepositSubmitParamsScan struct {
	// The account number to be returned in the check deposit's scan data.
	AccountNumber param.Field[string] `json:"account_number" api:"required"`
	// The routing number to be returned in the check deposit's scan data.
	RoutingNumber param.Field[string] `json:"routing_number" api:"required"`
	// The auxiliary on-us data to be returned in the check deposit's scan data.
	AuxiliaryOnUs param.Field[string] `json:"auxiliary_on_us"`
}

func (r SimulationCheckDepositSubmitParamsScan) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

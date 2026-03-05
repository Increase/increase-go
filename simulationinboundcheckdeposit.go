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

// SimulationInboundCheckDepositService contains methods and other services that
// help with interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimulationInboundCheckDepositService] method instead.
type SimulationInboundCheckDepositService struct {
	Options []option.RequestOption
}

// NewSimulationInboundCheckDepositService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSimulationInboundCheckDepositService(opts ...option.RequestOption) (r *SimulationInboundCheckDepositService) {
	r = &SimulationInboundCheckDepositService{}
	r.Options = opts
	return
}

// Simulates an Inbound Check Deposit against your account. This imitates someone
// depositing a check at their bank that was issued from your account. It may or
// may not be associated with a Check Transfer. Increase will evaluate the Inbound
// Check Deposit as we would in production and either create a Transaction or a
// Declined Transaction as a result. You can inspect the resulting Inbound Check
// Deposit object to see the result.
func (r *SimulationInboundCheckDepositService) New(ctx context.Context, body SimulationInboundCheckDepositNewParams, opts ...option.RequestOption) (res *InboundCheckDeposit, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "simulations/inbound_check_deposits"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Simulates an adjustment on an Inbound Check Deposit. The Inbound Check Deposit
// must have a `status` of `accepted`.
func (r *SimulationInboundCheckDepositService) Adjustment(ctx context.Context, inboundCheckDepositID string, body SimulationInboundCheckDepositAdjustmentParams, opts ...option.RequestOption) (res *InboundCheckDeposit, err error) {
	opts = slices.Concat(r.Options, opts)
	if inboundCheckDepositID == "" {
		err = errors.New("missing required inbound_check_deposit_id parameter")
		return
	}
	path := fmt.Sprintf("simulations/inbound_check_deposits/%s/adjustment", inboundCheckDepositID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SimulationInboundCheckDepositNewParams struct {
	// The identifier of the Account Number the Inbound Check Deposit will be against.
	AccountNumberID param.Field[string] `json:"account_number_id" api:"required"`
	// The check amount in cents.
	Amount param.Field[int64] `json:"amount" api:"required"`
	// The check number on the check to be deposited.
	CheckNumber param.Field[string] `json:"check_number" api:"required"`
	// Simulate the outcome of
	// [payee name checking](https://increase.com/documentation/positive-pay#payee-name-mismatches).
	// Defaults to `not_evaluated`.
	PayeeNameAnalysis param.Field[SimulationInboundCheckDepositNewParamsPayeeNameAnalysis] `json:"payee_name_analysis"`
}

func (r SimulationInboundCheckDepositNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Simulate the outcome of
// [payee name checking](https://increase.com/documentation/positive-pay#payee-name-mismatches).
// Defaults to `not_evaluated`.
type SimulationInboundCheckDepositNewParamsPayeeNameAnalysis string

const (
	SimulationInboundCheckDepositNewParamsPayeeNameAnalysisNameMatches  SimulationInboundCheckDepositNewParamsPayeeNameAnalysis = "name_matches"
	SimulationInboundCheckDepositNewParamsPayeeNameAnalysisDoesNotMatch SimulationInboundCheckDepositNewParamsPayeeNameAnalysis = "does_not_match"
	SimulationInboundCheckDepositNewParamsPayeeNameAnalysisNotEvaluated SimulationInboundCheckDepositNewParamsPayeeNameAnalysis = "not_evaluated"
)

func (r SimulationInboundCheckDepositNewParamsPayeeNameAnalysis) IsKnown() bool {
	switch r {
	case SimulationInboundCheckDepositNewParamsPayeeNameAnalysisNameMatches, SimulationInboundCheckDepositNewParamsPayeeNameAnalysisDoesNotMatch, SimulationInboundCheckDepositNewParamsPayeeNameAnalysisNotEvaluated:
		return true
	}
	return false
}

type SimulationInboundCheckDepositAdjustmentParams struct {
	// The adjustment amount in cents. Defaults to the amount of the Inbound Check
	// Deposit.
	Amount param.Field[int64] `json:"amount"`
	// The reason for the adjustment. Defaults to `wrong_payee_credit`.
	Reason param.Field[SimulationInboundCheckDepositAdjustmentParamsReason] `json:"reason"`
}

func (r SimulationInboundCheckDepositAdjustmentParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The reason for the adjustment. Defaults to `wrong_payee_credit`.
type SimulationInboundCheckDepositAdjustmentParamsReason string

const (
	SimulationInboundCheckDepositAdjustmentParamsReasonLateReturn        SimulationInboundCheckDepositAdjustmentParamsReason = "late_return"
	SimulationInboundCheckDepositAdjustmentParamsReasonWrongPayeeCredit  SimulationInboundCheckDepositAdjustmentParamsReason = "wrong_payee_credit"
	SimulationInboundCheckDepositAdjustmentParamsReasonAdjustedAmount    SimulationInboundCheckDepositAdjustmentParamsReason = "adjusted_amount"
	SimulationInboundCheckDepositAdjustmentParamsReasonNonConformingItem SimulationInboundCheckDepositAdjustmentParamsReason = "non_conforming_item"
	SimulationInboundCheckDepositAdjustmentParamsReasonPaid              SimulationInboundCheckDepositAdjustmentParamsReason = "paid"
)

func (r SimulationInboundCheckDepositAdjustmentParamsReason) IsKnown() bool {
	switch r {
	case SimulationInboundCheckDepositAdjustmentParamsReasonLateReturn, SimulationInboundCheckDepositAdjustmentParamsReasonWrongPayeeCredit, SimulationInboundCheckDepositAdjustmentParamsReasonAdjustedAmount, SimulationInboundCheckDepositAdjustmentParamsReasonNonConformingItem, SimulationInboundCheckDepositAdjustmentParamsReasonPaid:
		return true
	}
	return false
}

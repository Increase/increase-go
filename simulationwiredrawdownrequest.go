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

// SimulationWireDrawdownRequestService contains methods and other services that
// help with interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimulationWireDrawdownRequestService] method instead.
type SimulationWireDrawdownRequestService struct {
	Options []option.RequestOption
}

// NewSimulationWireDrawdownRequestService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSimulationWireDrawdownRequestService(opts ...option.RequestOption) (r *SimulationWireDrawdownRequestService) {
	r = &SimulationWireDrawdownRequestService{}
	r.Options = opts
	return
}

// Simulates a Wire Drawdown Request being refused by the debtor.
func (r *SimulationWireDrawdownRequestService) Refuse(ctx context.Context, wireDrawdownRequestID string, body SimulationWireDrawdownRequestRefuseParams, opts ...option.RequestOption) (res *WireDrawdownRequest, err error) {
	opts = slices.Concat(r.Options, opts)
	if wireDrawdownRequestID == "" {
		err = errors.New("missing required wire_drawdown_request_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("simulations/wire_drawdown_requests/%s/refuse", wireDrawdownRequestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Simulates a Wire Drawdown Request being submitted to Fedwire.
func (r *SimulationWireDrawdownRequestService) Submit(ctx context.Context, wireDrawdownRequestID string, opts ...option.RequestOption) (res *WireDrawdownRequest, err error) {
	opts = slices.Concat(r.Options, opts)
	if wireDrawdownRequestID == "" {
		err = errors.New("missing required wire_drawdown_request_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("simulations/wire_drawdown_requests/%s/submit", wireDrawdownRequestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type SimulationWireDrawdownRequestRefuseParams struct {
	// Additional information about the refusal that the simulated refusal will have.
	RefusalReasonAdditionalInformation param.Field[string] `json:"refusal_reason_additional_information"`
	// The reason code that the simulated refusal will have.
	//
	// Defaults to `other`.
	RefusalReasonCode param.Field[SimulationWireDrawdownRequestRefuseParamsRefusalReasonCode] `json:"refusal_reason_code"`
}

func (r SimulationWireDrawdownRequestRefuseParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The reason code that the simulated refusal will have.
type SimulationWireDrawdownRequestRefuseParamsRefusalReasonCode string

const (
	SimulationWireDrawdownRequestRefuseParamsRefusalReasonCodeInvalidDebtorAccountNumber  SimulationWireDrawdownRequestRefuseParamsRefusalReasonCode = "invalid_debtor_account_number"
	SimulationWireDrawdownRequestRefuseParamsRefusalReasonCodeTransactionForbidden        SimulationWireDrawdownRequestRefuseParamsRefusalReasonCode = "transaction_forbidden"
	SimulationWireDrawdownRequestRefuseParamsRefusalReasonCodeUnsuccessfulDirectDebit     SimulationWireDrawdownRequestRefuseParamsRefusalReasonCode = "unsuccessful_direct_debit"
	SimulationWireDrawdownRequestRefuseParamsRefusalReasonCodeInconsistentWithEndCustomer SimulationWireDrawdownRequestRefuseParamsRefusalReasonCode = "inconsistent_with_end_customer"
	SimulationWireDrawdownRequestRefuseParamsRefusalReasonCodeNarrative                   SimulationWireDrawdownRequestRefuseParamsRefusalReasonCode = "narrative"
	SimulationWireDrawdownRequestRefuseParamsRefusalReasonCodeOther                       SimulationWireDrawdownRequestRefuseParamsRefusalReasonCode = "other"
)

func (r SimulationWireDrawdownRequestRefuseParamsRefusalReasonCode) IsKnown() bool {
	switch r {
	case SimulationWireDrawdownRequestRefuseParamsRefusalReasonCodeInvalidDebtorAccountNumber, SimulationWireDrawdownRequestRefuseParamsRefusalReasonCodeTransactionForbidden, SimulationWireDrawdownRequestRefuseParamsRefusalReasonCodeUnsuccessfulDirectDebit, SimulationWireDrawdownRequestRefuseParamsRefusalReasonCodeInconsistentWithEndCustomer, SimulationWireDrawdownRequestRefuseParamsRefusalReasonCodeNarrative, SimulationWireDrawdownRequestRefuseParamsRefusalReasonCodeOther:
		return true
	}
	return false
}

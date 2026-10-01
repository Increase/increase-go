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

// SimulationFednowTransferService contains methods and other services that help
// with interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimulationFednowTransferService] method instead.
type SimulationFednowTransferService struct {
	Options []option.RequestOption
}

// NewSimulationFednowTransferService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSimulationFednowTransferService(opts ...option.RequestOption) (r *SimulationFednowTransferService) {
	r = &SimulationFednowTransferService{}
	r.Options = opts
	return
}

// Simulates submission of a [FedNow Transfer](#fednow-transfers) and handling the
// response from the destination financial institution. This transfer must first
// have a `status` of `pending_submitting`.
func (r *SimulationFednowTransferService) Complete(ctx context.Context, fednowTransferID string, body SimulationFednowTransferCompleteParams, opts ...option.RequestOption) (res *FednowTransfer, err error) {
	opts = slices.Concat(r.Options, opts)
	if fednowTransferID == "" {
		err = errors.New("missing required fednow_transfer_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("simulations/fednow_transfers/%s/complete", fednowTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type SimulationFednowTransferCompleteParams struct {
	// If set, the simulation will reject the transfer.
	Rejection param.Field[SimulationFednowTransferCompleteParamsRejection] `json:"rejection"`
}

func (r SimulationFednowTransferCompleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If set, the simulation will reject the transfer.
type SimulationFednowTransferCompleteParamsRejection struct {
	// The reason code that the simulated rejection will have.
	RejectReasonCode param.Field[SimulationFednowTransferCompleteParamsRejectionRejectReasonCode] `json:"reject_reason_code" api:"required"`
}

func (r SimulationFednowTransferCompleteParamsRejection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The reason code that the simulated rejection will have.
type SimulationFednowTransferCompleteParamsRejectionRejectReasonCode string

const (
	SimulationFednowTransferCompleteParamsRejectionRejectReasonCodeAccountClosed                                 SimulationFednowTransferCompleteParamsRejectionRejectReasonCode = "account_closed"
	SimulationFednowTransferCompleteParamsRejectionRejectReasonCodeAccountBlocked                                SimulationFednowTransferCompleteParamsRejectionRejectReasonCode = "account_blocked"
	SimulationFednowTransferCompleteParamsRejectionRejectReasonCodeInvalidCreditorAccountType                    SimulationFednowTransferCompleteParamsRejectionRejectReasonCode = "invalid_creditor_account_type"
	SimulationFednowTransferCompleteParamsRejectionRejectReasonCodeInvalidCreditorAccountNumber                  SimulationFednowTransferCompleteParamsRejectionRejectReasonCode = "invalid_creditor_account_number"
	SimulationFednowTransferCompleteParamsRejectionRejectReasonCodeInvalidCreditorFinancialInstitutionIdentifier SimulationFednowTransferCompleteParamsRejectionRejectReasonCode = "invalid_creditor_financial_institution_identifier"
	SimulationFednowTransferCompleteParamsRejectionRejectReasonCodeEndCustomerDeceased                           SimulationFednowTransferCompleteParamsRejectionRejectReasonCode = "end_customer_deceased"
	SimulationFednowTransferCompleteParamsRejectionRejectReasonCodeNarrative                                     SimulationFednowTransferCompleteParamsRejectionRejectReasonCode = "narrative"
	SimulationFednowTransferCompleteParamsRejectionRejectReasonCodeTransactionForbidden                          SimulationFednowTransferCompleteParamsRejectionRejectReasonCode = "transaction_forbidden"
	SimulationFednowTransferCompleteParamsRejectionRejectReasonCodeTransactionTypeNotSupported                   SimulationFednowTransferCompleteParamsRejectionRejectReasonCode = "transaction_type_not_supported"
	SimulationFednowTransferCompleteParamsRejectionRejectReasonCodeAmountExceedsBankLimits                       SimulationFednowTransferCompleteParamsRejectionRejectReasonCode = "amount_exceeds_bank_limits"
	SimulationFednowTransferCompleteParamsRejectionRejectReasonCodeInvalidCreditorAddress                        SimulationFednowTransferCompleteParamsRejectionRejectReasonCode = "invalid_creditor_address"
	SimulationFednowTransferCompleteParamsRejectionRejectReasonCodeInvalidDebtorAddress                          SimulationFednowTransferCompleteParamsRejectionRejectReasonCode = "invalid_debtor_address"
	SimulationFednowTransferCompleteParamsRejectionRejectReasonCodeTimeout                                       SimulationFednowTransferCompleteParamsRejectionRejectReasonCode = "timeout"
	SimulationFednowTransferCompleteParamsRejectionRejectReasonCodeProcessingError                               SimulationFednowTransferCompleteParamsRejectionRejectReasonCode = "processing_error"
	SimulationFednowTransferCompleteParamsRejectionRejectReasonCodeOther                                         SimulationFednowTransferCompleteParamsRejectionRejectReasonCode = "other"
)

func (r SimulationFednowTransferCompleteParamsRejectionRejectReasonCode) IsKnown() bool {
	switch r {
	case SimulationFednowTransferCompleteParamsRejectionRejectReasonCodeAccountClosed, SimulationFednowTransferCompleteParamsRejectionRejectReasonCodeAccountBlocked, SimulationFednowTransferCompleteParamsRejectionRejectReasonCodeInvalidCreditorAccountType, SimulationFednowTransferCompleteParamsRejectionRejectReasonCodeInvalidCreditorAccountNumber, SimulationFednowTransferCompleteParamsRejectionRejectReasonCodeInvalidCreditorFinancialInstitutionIdentifier, SimulationFednowTransferCompleteParamsRejectionRejectReasonCodeEndCustomerDeceased, SimulationFednowTransferCompleteParamsRejectionRejectReasonCodeNarrative, SimulationFednowTransferCompleteParamsRejectionRejectReasonCodeTransactionForbidden, SimulationFednowTransferCompleteParamsRejectionRejectReasonCodeTransactionTypeNotSupported, SimulationFednowTransferCompleteParamsRejectionRejectReasonCodeAmountExceedsBankLimits, SimulationFednowTransferCompleteParamsRejectionRejectReasonCodeInvalidCreditorAddress, SimulationFednowTransferCompleteParamsRejectionRejectReasonCodeInvalidDebtorAddress, SimulationFednowTransferCompleteParamsRejectionRejectReasonCodeTimeout, SimulationFednowTransferCompleteParamsRejectionRejectReasonCodeProcessingError, SimulationFednowTransferCompleteParamsRejectionRejectReasonCodeOther:
		return true
	}
	return false
}

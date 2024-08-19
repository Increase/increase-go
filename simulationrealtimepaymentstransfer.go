// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/option"
)

// SimulationRealTimePaymentsTransferService contains methods and other services
// that help with interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimulationRealTimePaymentsTransferService] method instead.
type SimulationRealTimePaymentsTransferService struct {
	Options []option.RequestOption
}

// NewSimulationRealTimePaymentsTransferService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewSimulationRealTimePaymentsTransferService(opts ...option.RequestOption) (r *SimulationRealTimePaymentsTransferService) {
	r = &SimulationRealTimePaymentsTransferService{}
	r.Options = opts
	return
}

// Simulates submission of a Real-Time Payments transfer and handling the response
// from the destination financial institution. This transfer must first have a
// `status` of `pending_submission`.
func (r *SimulationRealTimePaymentsTransferService) Complete(ctx context.Context, realTimePaymentsTransferID string, body SimulationRealTimePaymentsTransferCompleteParams, opts ...option.RequestOption) (res *RealTimePaymentsTransfer, err error) {
	opts = append(r.Options[:], opts...)
	if realTimePaymentsTransferID == "" {
		err = errors.New("missing required real_time_payments_transfer_id parameter")
		return
	}
	path := fmt.Sprintf("simulations/real_time_payments_transfers/%s/complete", realTimePaymentsTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SimulationRealTimePaymentsTransferCompleteParams struct {
	// If set, the simulation will reject the transfer.
	Rejection param.Field[SimulationRealTimePaymentsTransferCompleteParamsRejection] `json:"rejection"`
}

func (r SimulationRealTimePaymentsTransferCompleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If set, the simulation will reject the transfer.
type SimulationRealTimePaymentsTransferCompleteParamsRejection struct {
	// The reason code that the simulated rejection will have.
	RejectReasonCode param.Field[SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode] `json:"reject_reason_code,required"`
}

func (r SimulationRealTimePaymentsTransferCompleteParamsRejection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The reason code that the simulated rejection will have.
type SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode string

const (
	// The destination account is closed. Corresponds to the Real-Time Payments reason
	// code `AC04`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeAccountClosed SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "account_closed"
	// The destination account is currently blocked from receiving transactions.
	// Corresponds to the Real-Time Payments reason code `AC06`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeAccountBlocked SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "account_blocked"
	// The destination account is ineligible to receive Real-Time Payments transfers.
	// Corresponds to the Real-Time Payments reason code `AC14`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeInvalidCreditorAccountType SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "invalid_creditor_account_type"
	// The destination account does not exist. Corresponds to the Real-Time Payments
	// reason code `AC03`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeInvalidCreditorAccountNumber SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "invalid_creditor_account_number"
	// The destination routing number is invalid. Corresponds to the Real-Time Payments
	// reason code `RC04`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeInvalidCreditorFinancialInstitutionIdentifier SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "invalid_creditor_financial_institution_identifier"
	// The destination account holder is deceased. Corresponds to the Real-Time
	// Payments reason code `MD07`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeEndCustomerDeceased SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "end_customer_deceased"
	// The reason is provided as narrative information in the additional information
	// field.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeNarrative SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "narrative"
	// Real-Time Payments transfers are not allowed to the destination account.
	// Corresponds to the Real-Time Payments reason code `AG01`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeTransactionForbidden SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "transaction_forbidden"
	// Real-Time Payments transfers are not enabled for the destination account.
	// Corresponds to the Real-Time Payments reason code `AG03`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeTransactionTypeNotSupported SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "transaction_type_not_supported"
	// The amount of the transfer is different than expected by the recipient.
	// Corresponds to the Real-Time Payments reason code `AM09`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeUnexpectedAmount SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "unexpected_amount"
	// The amount is higher than the recipient is authorized to send or receive.
	// Corresponds to the Real-Time Payments reason code `AM14`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeAmountExceedsBankLimits SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "amount_exceeds_bank_limits"
	// The creditor's address is required, but missing or invalid. Corresponds to the
	// Real-Time Payments reason code `BE04`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeInvalidCreditorAddress SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "invalid_creditor_address"
	// The specified creditor is unknown. Corresponds to the Real-Time Payments reason
	// code `BE06`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeUnknownEndCustomer SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "unknown_end_customer"
	// The debtor's address is required, but missing or invalid. Corresponds to the
	// Real-Time Payments reason code `BE07`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeInvalidDebtorAddress SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "invalid_debtor_address"
	// There was a timeout processing the transfer. Corresponds to the Real-Time
	// Payments reason code `DS24`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeTimeout SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "timeout"
	// Real-Time Payments transfers are not enabled for the destination account.
	// Corresponds to the Real-Time Payments reason code `NOAT`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeUnsupportedMessageForRecipient SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "unsupported_message_for_recipient"
	// The destination financial institution is currently not connected to Real-Time
	// Payments. Corresponds to the Real-Time Payments reason code `9912`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeRecipientConnectionNotAvailable SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "recipient_connection_not_available"
	// Real-Time Payments is currently unavailable. Corresponds to the Real-Time
	// Payments reason code `9948`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeRealTimePaymentsSuspended SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "real_time_payments_suspended"
	// The destination financial institution is currently signed off of Real-Time
	// Payments. Corresponds to the Real-Time Payments reason code `9910`.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeInstructedAgentSignedOff SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "instructed_agent_signed_off"
	// The transfer was rejected due to an internal Increase issue. We have been
	// notified.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeProcessingError SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "processing_error"
	// Some other error or issue has occurred.
	SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeOther SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode = "other"
)

func (r SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCode) IsKnown() bool {
	switch r {
	case SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeAccountClosed, SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeAccountBlocked, SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeInvalidCreditorAccountType, SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeInvalidCreditorAccountNumber, SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeInvalidCreditorFinancialInstitutionIdentifier, SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeEndCustomerDeceased, SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeNarrative, SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeTransactionForbidden, SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeTransactionTypeNotSupported, SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeUnexpectedAmount, SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeAmountExceedsBankLimits, SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeInvalidCreditorAddress, SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeUnknownEndCustomer, SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeInvalidDebtorAddress, SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeTimeout, SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeUnsupportedMessageForRecipient, SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeRecipientConnectionNotAvailable, SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeRealTimePaymentsSuspended, SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeInstructedAgentSignedOff, SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeProcessingError, SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeOther:
		return true
	}
	return false
}

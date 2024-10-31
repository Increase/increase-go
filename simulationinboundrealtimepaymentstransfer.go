// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"net/http"

	"github.com/Increase/increase-go/internal/apijson"
	"github.com/Increase/increase-go/internal/param"
	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
)

// SimulationInboundRealTimePaymentsTransferService contains methods and other
// services that help with interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimulationInboundRealTimePaymentsTransferService] method instead.
type SimulationInboundRealTimePaymentsTransferService struct {
	Options []option.RequestOption
}

// NewSimulationInboundRealTimePaymentsTransferService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewSimulationInboundRealTimePaymentsTransferService(opts ...option.RequestOption) (r *SimulationInboundRealTimePaymentsTransferService) {
	r = &SimulationInboundRealTimePaymentsTransferService{}
	r.Options = opts
	return
}

// Simulates an
// [Inbound Real-Time Payments Transfer](#inbound-real-time-payments-transfers) to
// your account. Real-Time Payments are a beta feature.
func (r *SimulationInboundRealTimePaymentsTransferService) New(ctx context.Context, body SimulationInboundRealTimePaymentsTransferNewParams, opts ...option.RequestOption) (res *InboundRealTimePaymentsTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/inbound_real_time_payments_transfers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SimulationInboundRealTimePaymentsTransferNewParams struct {
	// The identifier of the Account Number the inbound Real-Time Payments Transfer is
	// for.
	AccountNumberID param.Field[string] `json:"account_number_id,required"`
	// The transfer amount in USD cents. Must be positive.
	Amount param.Field[int64] `json:"amount,required"`
	// The account number of the account that sent the transfer.
	DebtorAccountNumber param.Field[string] `json:"debtor_account_number"`
	// The name provided by the sender of the transfer.
	DebtorName param.Field[string] `json:"debtor_name"`
	// The routing number of the account that sent the transfer.
	DebtorRoutingNumber param.Field[string] `json:"debtor_routing_number"`
	// Additional information included with the transfer.
	RemittanceInformation param.Field[string] `json:"remittance_information"`
	// The identifier of a pending Request for Payment that this transfer will fulfill.
	RequestForPaymentID param.Field[string] `json:"request_for_payment_id"`
}

func (r SimulationInboundRealTimePaymentsTransferNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

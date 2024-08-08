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

// Simulates an inbound Real-Time Payments transfer to your account. Real-Time
// Payments are a beta feature.
func (r *SimulationInboundRealTimePaymentsTransferService) New(ctx context.Context, body SimulationInboundRealTimePaymentsTransferNewParams, opts ...option.RequestOption) (res *SimulationInboundRealTimePaymentsTransferNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/inbound_real_time_payments_transfers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// The results of an inbound Real-Time Payments Transfer simulation.
type SimulationInboundRealTimePaymentsTransferNewResponse struct {
	// If the Real-Time Payments Transfer attempt fails, this will contain the
	// resulting [Declined Transaction](#declined-transactions) object. The Declined
	// Transaction's `source` will be of
	// `category: inbound_real_time_payments_transfer_decline`.
	DeclinedTransaction DeclinedTransaction `json:"declined_transaction,required,nullable"`
	// If the Real-Time Payments Transfer attempt succeeds, this will contain the
	// resulting [Transaction](#transactions) object. The Transaction's `source` will
	// be of `category: inbound_real_time_payments_transfer_confirmation`.
	Transaction Transaction `json:"transaction,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `inbound_real_time_payments_transfer_simulation_result`.
	Type SimulationInboundRealTimePaymentsTransferNewResponseType `json:"type,required"`
	JSON simulationInboundRealTimePaymentsTransferNewResponseJSON `json:"-"`
}

// simulationInboundRealTimePaymentsTransferNewResponseJSON contains the JSON
// metadata for the struct [SimulationInboundRealTimePaymentsTransferNewResponse]
type simulationInboundRealTimePaymentsTransferNewResponseJSON struct {
	DeclinedTransaction apijson.Field
	Transaction         apijson.Field
	Type                apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *SimulationInboundRealTimePaymentsTransferNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r simulationInboundRealTimePaymentsTransferNewResponseJSON) RawJSON() string {
	return r.raw
}

// A constant representing the object's type. For this resource it will always be
// `inbound_real_time_payments_transfer_simulation_result`.
type SimulationInboundRealTimePaymentsTransferNewResponseType string

const (
	SimulationInboundRealTimePaymentsTransferNewResponseTypeInboundRealTimePaymentsTransferSimulationResult SimulationInboundRealTimePaymentsTransferNewResponseType = "inbound_real_time_payments_transfer_simulation_result"
)

func (r SimulationInboundRealTimePaymentsTransferNewResponseType) IsKnown() bool {
	switch r {
	case SimulationInboundRealTimePaymentsTransferNewResponseTypeInboundRealTimePaymentsTransferSimulationResult:
		return true
	}
	return false
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

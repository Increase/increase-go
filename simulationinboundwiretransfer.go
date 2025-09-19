// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"net/http"
	"slices"

	"github.com/Increase/increase-go/internal/apijson"
	"github.com/Increase/increase-go/internal/param"
	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
)

// SimulationInboundWireTransferService contains methods and other services that
// help with interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimulationInboundWireTransferService] method instead.
type SimulationInboundWireTransferService struct {
	Options []option.RequestOption
}

// NewSimulationInboundWireTransferService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSimulationInboundWireTransferService(opts ...option.RequestOption) (r *SimulationInboundWireTransferService) {
	r = &SimulationInboundWireTransferService{}
	r.Options = opts
	return
}

// Simulates an [Inbound Wire Transfer](#inbound-wire-transfers) to your account.
func (r *SimulationInboundWireTransferService) New(ctx context.Context, body SimulationInboundWireTransferNewParams, opts ...option.RequestOption) (res *InboundWireTransfer, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "simulations/inbound_wire_transfers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SimulationInboundWireTransferNewParams struct {
	// The identifier of the Account Number the inbound Wire Transfer is for.
	AccountNumberID param.Field[string] `json:"account_number_id,required"`
	// The transfer amount in cents. Must be positive.
	Amount param.Field[int64] `json:"amount,required"`
	// The sending bank will set creditor_address_line1 in production. You can simulate
	// any value here.
	CreditorAddressLine1 param.Field[string] `json:"creditor_address_line1"`
	// The sending bank will set creditor_address_line2 in production. You can simulate
	// any value here.
	CreditorAddressLine2 param.Field[string] `json:"creditor_address_line2"`
	// The sending bank will set creditor_address_line3 in production. You can simulate
	// any value here.
	CreditorAddressLine3 param.Field[string] `json:"creditor_address_line3"`
	// The sending bank will set creditor_name in production. You can simulate any
	// value here.
	CreditorName param.Field[string] `json:"creditor_name"`
	// The sending bank will set debtor_address_line1 in production. You can simulate
	// any value here.
	DebtorAddressLine1 param.Field[string] `json:"debtor_address_line1"`
	// The sending bank will set debtor_address_line2 in production. You can simulate
	// any value here.
	DebtorAddressLine2 param.Field[string] `json:"debtor_address_line2"`
	// The sending bank will set debtor_address_line3 in production. You can simulate
	// any value here.
	DebtorAddressLine3 param.Field[string] `json:"debtor_address_line3"`
	// The sending bank will set debtor_name in production. You can simulate any value
	// here.
	DebtorName param.Field[string] `json:"debtor_name"`
	// The sending bank will set end_to_end_identification in production. You can
	// simulate any value here.
	EndToEndIdentification param.Field[string] `json:"end_to_end_identification"`
	// The sending bank will set instructing_agent_routing_number in production. You
	// can simulate any value here.
	InstructingAgentRoutingNumber param.Field[string] `json:"instructing_agent_routing_number"`
	// The sending bank will set instruction_identification in production. You can
	// simulate any value here.
	InstructionIdentification param.Field[string] `json:"instruction_identification"`
	// The sending bank will set unique_end_to_end_transaction_reference in production.
	// You can simulate any value here.
	UniqueEndToEndTransactionReference param.Field[string] `json:"unique_end_to_end_transaction_reference"`
	// The sending bank will set unstructured_remittance_information in production. You
	// can simulate any value here.
	UnstructuredRemittanceInformation param.Field[string] `json:"unstructured_remittance_information"`
	// The identifier of a Wire Drawdown Request the inbound Wire Transfer is
	// fulfilling.
	WireDrawdownRequestID param.Field[string] `json:"wire_drawdown_request_id"`
}

func (r SimulationInboundWireTransferNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

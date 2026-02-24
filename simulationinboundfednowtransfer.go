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

// SimulationInboundFednowTransferService contains methods and other services that
// help with interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimulationInboundFednowTransferService] method instead.
type SimulationInboundFednowTransferService struct {
	Options []option.RequestOption
}

// NewSimulationInboundFednowTransferService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSimulationInboundFednowTransferService(opts ...option.RequestOption) (r *SimulationInboundFednowTransferService) {
	r = &SimulationInboundFednowTransferService{}
	r.Options = opts
	return
}

// Simulates an [Inbound FedNow Transfer](#inbound-fednow-transfers) to your
// account.
func (r *SimulationInboundFednowTransferService) New(ctx context.Context, body SimulationInboundFednowTransferNewParams, opts ...option.RequestOption) (res *InboundFednowTransfer, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "simulations/inbound_fednow_transfers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SimulationInboundFednowTransferNewParams struct {
	// The identifier of the Account Number the inbound FedNow Transfer is for.
	AccountNumberID param.Field[string] `json:"account_number_id" api:"required"`
	// The transfer amount in USD cents. Must be positive.
	Amount param.Field[int64] `json:"amount" api:"required"`
	// The account number of the account that sent the transfer.
	DebtorAccountNumber param.Field[string] `json:"debtor_account_number"`
	// The name provided by the sender of the transfer.
	DebtorName param.Field[string] `json:"debtor_name"`
	// The routing number of the account that sent the transfer.
	DebtorRoutingNumber param.Field[string] `json:"debtor_routing_number"`
	// Additional information included with the transfer.
	UnstructuredRemittanceInformation param.Field[string] `json:"unstructured_remittance_information"`
}

func (r SimulationInboundFednowTransferNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

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

// SimulationInboundWireDrawdownRequestService contains methods and other services
// that help with interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimulationInboundWireDrawdownRequestService] method instead.
type SimulationInboundWireDrawdownRequestService struct {
	Options []option.RequestOption
}

// NewSimulationInboundWireDrawdownRequestService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewSimulationInboundWireDrawdownRequestService(opts ...option.RequestOption) (r *SimulationInboundWireDrawdownRequestService) {
	r = &SimulationInboundWireDrawdownRequestService{}
	r.Options = opts
	return
}

// Simulates receiving an
// [Inbound Wire Drawdown Request](#inbound-wire-drawdown-requests).
func (r *SimulationInboundWireDrawdownRequestService) New(ctx context.Context, body SimulationInboundWireDrawdownRequestNewParams, opts ...option.RequestOption) (res *InboundWireDrawdownRequest, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "simulations/inbound_wire_drawdown_requests"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SimulationInboundWireDrawdownRequestNewParams struct {
	// The amount being requested in cents.
	Amount param.Field[int64] `json:"amount" api:"required"`
	// The creditor's account number.
	CreditorAccountNumber param.Field[string] `json:"creditor_account_number" api:"required"`
	// The creditor's routing number.
	CreditorRoutingNumber param.Field[string] `json:"creditor_routing_number" api:"required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the amount being
	// requested. Will always be "USD".
	Currency param.Field[string] `json:"currency" api:"required"`
	// The Account Number to which the recipient of this request is being requested to
	// send funds from.
	RecipientAccountNumberID param.Field[string] `json:"recipient_account_number_id" api:"required"`
	// A free-form address field set by the sender representing the first line of the
	// creditor's address.
	CreditorAddressLine1 param.Field[string] `json:"creditor_address_line1"`
	// A free-form address field set by the sender representing the second line of the
	// creditor's address.
	CreditorAddressLine2 param.Field[string] `json:"creditor_address_line2"`
	// A free-form address field set by the sender representing the third line of the
	// creditor's address.
	CreditorAddressLine3 param.Field[string] `json:"creditor_address_line3"`
	// A free-form name field set by the sender representing the creditor's name.
	CreditorName param.Field[string] `json:"creditor_name"`
	// The debtor's account number.
	DebtorAccountNumber param.Field[string] `json:"debtor_account_number"`
	// A free-form address field set by the sender representing the first line of the
	// debtor's address.
	DebtorAddressLine1 param.Field[string] `json:"debtor_address_line1"`
	// A free-form address field set by the sender representing the second line of the
	// debtor's address.
	DebtorAddressLine2 param.Field[string] `json:"debtor_address_line2"`
	// A free-form address field set by the sender.
	DebtorAddressLine3 param.Field[string] `json:"debtor_address_line3"`
	// A free-form name field set by the sender representing the debtor's name.
	DebtorName param.Field[string] `json:"debtor_name"`
	// The debtor's routing number.
	DebtorRoutingNumber param.Field[string] `json:"debtor_routing_number"`
	// A free-form reference string set by the sender, to help identify the transfer.
	EndToEndIdentification param.Field[string] `json:"end_to_end_identification"`
	// The sending bank's identifier for the wire transfer.
	InstructionIdentification param.Field[string] `json:"instruction_identification"`
	// The Unique End-to-end Transaction Reference
	// ([UETR](https://www.swift.com/payments/what-unique-end-end-transaction-reference-uetr))
	// of the transfer.
	UniqueEndToEndTransactionReference param.Field[string] `json:"unique_end_to_end_transaction_reference"`
	// A free-form message set by the sender.
	UnstructuredRemittanceInformation param.Field[string] `json:"unstructured_remittance_information"`
}

func (r SimulationInboundWireDrawdownRequestNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

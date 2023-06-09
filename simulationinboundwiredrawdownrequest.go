// File generated from our OpenAPI spec by Stainless.

package increase

import (
	"context"
	"net/http"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/option"
)

// SimulationInboundWireDrawdownRequestService contains methods and other services
// that help with interacting with the increase API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewSimulationInboundWireDrawdownRequestService] method instead.
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

// Simulates the receival of an
// [Inbound Wire Drawdown Request](#inbound-wire-drawdown-requests).
func (r *SimulationInboundWireDrawdownRequestService) New(ctx context.Context, body SimulationInboundWireDrawdownRequestNewParams, opts ...option.RequestOption) (res *InboundWireDrawdownRequest, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/inbound_wire_drawdown_requests"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SimulationInboundWireDrawdownRequestNewParams struct {
	// The amount being requested in cents.
	Amount param.Field[int64] `json:"amount,required"`
	// The drawdown request's beneficiary's account number.
	BeneficiaryAccountNumber param.Field[string] `json:"beneficiary_account_number,required"`
	// The drawdown request's beneficiary's routing number.
	BeneficiaryRoutingNumber param.Field[string] `json:"beneficiary_routing_number,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the amount being
	// requested. Will always be "USD".
	Currency param.Field[string] `json:"currency,required"`
	// A message from the drawdown request's originator.
	MessageToRecipient param.Field[string] `json:"message_to_recipient,required"`
	// The drawdown request's originator's account number.
	OriginatorAccountNumber param.Field[string] `json:"originator_account_number,required"`
	// The drawdown request's originator's routing number.
	OriginatorRoutingNumber param.Field[string] `json:"originator_routing_number,required"`
	// The Account Number to which the recipient of this request is being requested to
	// send funds from.
	RecipientAccountNumberID param.Field[string] `json:"recipient_account_number_id,required"`
	// Line 1 of the drawdown request's beneficiary's address.
	BeneficiaryAddressLine1 param.Field[string] `json:"beneficiary_address_line1"`
	// Line 2 of the drawdown request's beneficiary's address.
	BeneficiaryAddressLine2 param.Field[string] `json:"beneficiary_address_line2"`
	// Line 3 of the drawdown request's beneficiary's address.
	BeneficiaryAddressLine3 param.Field[string] `json:"beneficiary_address_line3"`
	// The drawdown request's beneficiary's name.
	BeneficiaryName param.Field[string] `json:"beneficiary_name"`
	// Line 1 of the drawdown request's originator's address.
	OriginatorAddressLine1 param.Field[string] `json:"originator_address_line1"`
	// Line 2 of the drawdown request's originator's address.
	OriginatorAddressLine2 param.Field[string] `json:"originator_address_line2"`
	// Line 3 of the drawdown request's originator's address.
	OriginatorAddressLine3 param.Field[string] `json:"originator_address_line3"`
	// The drawdown request's originator's name.
	OriginatorName param.Field[string] `json:"originator_name"`
	// Line 1 of the information conveyed from the originator of the message to the
	// beneficiary.
	OriginatorToBeneficiaryInformationLine1 param.Field[string] `json:"originator_to_beneficiary_information_line1"`
	// Line 2 of the information conveyed from the originator of the message to the
	// beneficiary.
	OriginatorToBeneficiaryInformationLine2 param.Field[string] `json:"originator_to_beneficiary_information_line2"`
	// Line 3 of the information conveyed from the originator of the message to the
	// beneficiary.
	OriginatorToBeneficiaryInformationLine3 param.Field[string] `json:"originator_to_beneficiary_information_line3"`
	// Line 4 of the information conveyed from the originator of the message to the
	// beneficiary.
	OriginatorToBeneficiaryInformationLine4 param.Field[string] `json:"originator_to_beneficiary_information_line4"`
}

func (r SimulationInboundWireDrawdownRequestNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

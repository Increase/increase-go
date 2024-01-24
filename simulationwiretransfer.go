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

// SimulationWireTransferService contains methods and other services that help with
// interacting with the increase API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSimulationWireTransferService]
// method instead.
type SimulationWireTransferService struct {
	Options []option.RequestOption
}

// NewSimulationWireTransferService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSimulationWireTransferService(opts ...option.RequestOption) (r *SimulationWireTransferService) {
	r = &SimulationWireTransferService{}
	r.Options = opts
	return
}

// Simulates an inbound Wire Transfer to your account.
func (r *SimulationWireTransferService) NewInbound(ctx context.Context, body SimulationWireTransferNewInboundParams, opts ...option.RequestOption) (res *InboundWireTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/inbound_wire_transfers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SimulationWireTransferNewInboundParams struct {
	// The identifier of the Account Number the inbound Wire Transfer is for.
	AccountNumberID param.Field[string] `json:"account_number_id,required"`
	// The transfer amount in cents. Must be positive.
	Amount param.Field[int64] `json:"amount,required"`
	// The sending bank will set beneficiary_address_line1 in production. You can
	// simulate any value here.
	BeneficiaryAddressLine1 param.Field[string] `json:"beneficiary_address_line1"`
	// The sending bank will set beneficiary_address_line2 in production. You can
	// simulate any value here.
	BeneficiaryAddressLine2 param.Field[string] `json:"beneficiary_address_line2"`
	// The sending bank will set beneficiary_address_line3 in production. You can
	// simulate any value here.
	BeneficiaryAddressLine3 param.Field[string] `json:"beneficiary_address_line3"`
	// The sending bank will set beneficiary_name in production. You can simulate any
	// value here.
	BeneficiaryName param.Field[string] `json:"beneficiary_name"`
	// The sending bank will set beneficiary_reference in production. You can simulate
	// any value here.
	BeneficiaryReference param.Field[string] `json:"beneficiary_reference"`
	// The sending bank will set originator_address_line1 in production. You can
	// simulate any value here.
	OriginatorAddressLine1 param.Field[string] `json:"originator_address_line1"`
	// The sending bank will set originator_address_line2 in production. You can
	// simulate any value here.
	OriginatorAddressLine2 param.Field[string] `json:"originator_address_line2"`
	// The sending bank will set originator_address_line3 in production. You can
	// simulate any value here.
	OriginatorAddressLine3 param.Field[string] `json:"originator_address_line3"`
	// The sending bank will set originator_name in production. You can simulate any
	// value here.
	OriginatorName param.Field[string] `json:"originator_name"`
	// The sending bank will set originator_routing_number in production. You can
	// simulate any value here.
	OriginatorRoutingNumber param.Field[string] `json:"originator_routing_number"`
	// The sending bank will set originator_to_beneficiary_information_line1 in
	// production. You can simulate any value here.
	OriginatorToBeneficiaryInformationLine1 param.Field[string] `json:"originator_to_beneficiary_information_line1"`
	// The sending bank will set originator_to_beneficiary_information_line2 in
	// production. You can simulate any value here.
	OriginatorToBeneficiaryInformationLine2 param.Field[string] `json:"originator_to_beneficiary_information_line2"`
	// The sending bank will set originator_to_beneficiary_information_line3 in
	// production. You can simulate any value here.
	OriginatorToBeneficiaryInformationLine3 param.Field[string] `json:"originator_to_beneficiary_information_line3"`
	// The sending bank will set originator_to_beneficiary_information_line4 in
	// production. You can simulate any value here.
	OriginatorToBeneficiaryInformationLine4 param.Field[string] `json:"originator_to_beneficiary_information_line4"`
}

func (r SimulationWireTransferNewInboundParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

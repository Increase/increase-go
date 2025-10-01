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

// SimulationCardDisputeService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimulationCardDisputeService] method instead.
type SimulationCardDisputeService struct {
	Options []option.RequestOption
}

// NewSimulationCardDisputeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSimulationCardDisputeService(opts ...option.RequestOption) (r *SimulationCardDisputeService) {
	r = &SimulationCardDisputeService{}
	r.Options = opts
	return
}

// After a [Card Dispute](#card-disputes) is created in production, the dispute
// will initially be in a `pending_user_submission_reviewing` state. Since no
// review or further action happens in sandbox, this endpoint simulates moving a
// Card Dispute through its various states.
func (r *SimulationCardDisputeService) Action(ctx context.Context, cardDisputeID string, body SimulationCardDisputeActionParams, opts ...option.RequestOption) (res *CardDispute, err error) {
	opts = slices.Concat(r.Options, opts)
	if cardDisputeID == "" {
		err = errors.New("missing required card_dispute_id parameter")
		return
	}
	path := fmt.Sprintf("simulations/card_disputes/%s/action", cardDisputeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SimulationCardDisputeActionParams struct {
	// The network of the Card Dispute. Details specific to the network are required
	// under the sub-object with the same identifier as the network.
	Network param.Field[SimulationCardDisputeActionParamsNetwork] `json:"network,required"`
	// The Visa-specific parameters for the taking action on the dispute. Required if
	// and only if `network` is `visa`.
	Visa param.Field[SimulationCardDisputeActionParamsVisa] `json:"visa"`
}

func (r SimulationCardDisputeActionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The network of the Card Dispute. Details specific to the network are required
// under the sub-object with the same identifier as the network.
type SimulationCardDisputeActionParamsNetwork string

const (
	SimulationCardDisputeActionParamsNetworkVisa SimulationCardDisputeActionParamsNetwork = "visa"
)

func (r SimulationCardDisputeActionParamsNetwork) IsKnown() bool {
	switch r {
	case SimulationCardDisputeActionParamsNetworkVisa:
		return true
	}
	return false
}

// The Visa-specific parameters for the taking action on the dispute. Required if
// and only if `network` is `visa`.
type SimulationCardDisputeActionParamsVisa struct {
	// The action to take. Details specific to the action are required under the
	// sub-object with the same identifier as the action.
	Action param.Field[SimulationCardDisputeActionParamsVisaAction] `json:"action,required"`
	// The parameters for accepting the chargeback. Required if and only if `action` is
	// `accept_chargeback`.
	AcceptChargeback param.Field[interface{}] `json:"accept_chargeback"`
	// The parameters for accepting the user submission. Required if and only if
	// `action` is `accept_user_submission`.
	AcceptUserSubmission param.Field[interface{}] `json:"accept_user_submission"`
	// The parameters for declining the prearbitration. Required if and only if
	// `action` is `decline_user_prearbitration`.
	DeclineUserPrearbitration param.Field[interface{}] `json:"decline_user_prearbitration"`
	// The parameters for receiving the prearbitration. Required if and only if
	// `action` is `receive_merchant_prearbitration`.
	ReceiveMerchantPrearbitration param.Field[interface{}] `json:"receive_merchant_prearbitration"`
	// The parameters for re-presenting the dispute. Required if and only if `action`
	// is `represent`.
	Represent param.Field[interface{}] `json:"represent"`
	// The parameters for requesting further information from the user. Required if and
	// only if `action` is `request_further_information`.
	RequestFurtherInformation param.Field[SimulationCardDisputeActionParamsVisaRequestFurtherInformation] `json:"request_further_information"`
	// The parameters for timing out the chargeback. Required if and only if `action`
	// is `time_out_chargeback`.
	TimeOutChargeback param.Field[interface{}] `json:"time_out_chargeback"`
	// The parameters for timing out the merchant prearbitration. Required if and only
	// if `action` is `time_out_merchant_prearbitration`.
	TimeOutMerchantPrearbitration param.Field[interface{}] `json:"time_out_merchant_prearbitration"`
	// The parameters for timing out the re-presentment. Required if and only if
	// `action` is `time_out_representment`.
	TimeOutRepresentment param.Field[interface{}] `json:"time_out_representment"`
	// The parameters for timing out the user prearbitration. Required if and only if
	// `action` is `time_out_user_prearbitration`.
	TimeOutUserPrearbitration param.Field[interface{}] `json:"time_out_user_prearbitration"`
}

func (r SimulationCardDisputeActionParamsVisa) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to take. Details specific to the action are required under the
// sub-object with the same identifier as the action.
type SimulationCardDisputeActionParamsVisaAction string

const (
	SimulationCardDisputeActionParamsVisaActionAcceptChargeback              SimulationCardDisputeActionParamsVisaAction = "accept_chargeback"
	SimulationCardDisputeActionParamsVisaActionAcceptUserSubmission          SimulationCardDisputeActionParamsVisaAction = "accept_user_submission"
	SimulationCardDisputeActionParamsVisaActionDeclineUserPrearbitration     SimulationCardDisputeActionParamsVisaAction = "decline_user_prearbitration"
	SimulationCardDisputeActionParamsVisaActionReceiveMerchantPrearbitration SimulationCardDisputeActionParamsVisaAction = "receive_merchant_prearbitration"
	SimulationCardDisputeActionParamsVisaActionRepresent                     SimulationCardDisputeActionParamsVisaAction = "represent"
	SimulationCardDisputeActionParamsVisaActionRequestFurtherInformation     SimulationCardDisputeActionParamsVisaAction = "request_further_information"
	SimulationCardDisputeActionParamsVisaActionTimeOutChargeback             SimulationCardDisputeActionParamsVisaAction = "time_out_chargeback"
	SimulationCardDisputeActionParamsVisaActionTimeOutMerchantPrearbitration SimulationCardDisputeActionParamsVisaAction = "time_out_merchant_prearbitration"
	SimulationCardDisputeActionParamsVisaActionTimeOutRepresentment          SimulationCardDisputeActionParamsVisaAction = "time_out_representment"
	SimulationCardDisputeActionParamsVisaActionTimeOutUserPrearbitration     SimulationCardDisputeActionParamsVisaAction = "time_out_user_prearbitration"
)

func (r SimulationCardDisputeActionParamsVisaAction) IsKnown() bool {
	switch r {
	case SimulationCardDisputeActionParamsVisaActionAcceptChargeback, SimulationCardDisputeActionParamsVisaActionAcceptUserSubmission, SimulationCardDisputeActionParamsVisaActionDeclineUserPrearbitration, SimulationCardDisputeActionParamsVisaActionReceiveMerchantPrearbitration, SimulationCardDisputeActionParamsVisaActionRepresent, SimulationCardDisputeActionParamsVisaActionRequestFurtherInformation, SimulationCardDisputeActionParamsVisaActionTimeOutChargeback, SimulationCardDisputeActionParamsVisaActionTimeOutMerchantPrearbitration, SimulationCardDisputeActionParamsVisaActionTimeOutRepresentment, SimulationCardDisputeActionParamsVisaActionTimeOutUserPrearbitration:
		return true
	}
	return false
}

// The parameters for requesting further information from the user. Required if and
// only if `action` is `request_further_information`.
type SimulationCardDisputeActionParamsVisaRequestFurtherInformation struct {
	// The reason for requesting further information from the user.
	Reason param.Field[string] `json:"reason,required"`
}

func (r SimulationCardDisputeActionParamsVisaRequestFurtherInformation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

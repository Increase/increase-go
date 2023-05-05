package increase

import (
	"context"
	"fmt"
	"net/http"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/option"
)

// SimulationCheckTransferService contains methods and other services that help
// with interacting with the increase API. Note, unlike clients, this service does
// not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewSimulationCheckTransferService] method instead.
type SimulationCheckTransferService struct {
	Options []option.RequestOption
}

// NewSimulationCheckTransferService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSimulationCheckTransferService(opts ...option.RequestOption) (r *SimulationCheckTransferService) {
	r = &SimulationCheckTransferService{}
	r.Options = opts
	return
}

// Simulates a [Check Transfer](#check-transfers) being deposited at a bank. This
// transfer must first have a `status` of `mailed`.
func (r *SimulationCheckTransferService) Deposit(ctx context.Context, check_transfer_id string, opts ...option.RequestOption) (res *CheckTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("simulations/check_transfers/%s/deposit", check_transfer_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Simulates the mailing of a [Check Transfer](#check-transfers), which happens
// once per weekday in production but can be sped up in sandbox. This transfer must
// first have a `status` of `pending_approval` or `pending_submission`.
func (r *SimulationCheckTransferService) Mail(ctx context.Context, check_transfer_id string, opts ...option.RequestOption) (res *CheckTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("simulations/check_transfers/%s/mail", check_transfer_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Simulates a [Check Transfer](#check-transfers) being returned via USPS to
// Increase. This transfer must first have a `status` of `mailed`.
func (r *SimulationCheckTransferService) Return(ctx context.Context, check_transfer_id string, body SimulationCheckTransferReturnParams, opts ...option.RequestOption) (res *CheckTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("simulations/check_transfers/%s/return", check_transfer_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SimulationCheckTransferReturnParams struct {
	// The reason why the Check Transfer was returned to Increase.
	Reason param.Field[SimulationCheckTransferReturnParamsReason] `json:"reason,required"`
}

func (r SimulationCheckTransferReturnParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SimulationCheckTransferReturnParamsReason string

const (
	SimulationCheckTransferReturnParamsReasonMailDeliveryFailure   SimulationCheckTransferReturnParamsReason = "mail_delivery_failure"
	SimulationCheckTransferReturnParamsReasonRefusedByRecipient    SimulationCheckTransferReturnParamsReason = "refused_by_recipient"
	SimulationCheckTransferReturnParamsReasonReturnedNotAuthorized SimulationCheckTransferReturnParamsReason = "returned_not_authorized"
)

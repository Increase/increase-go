// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"net/http"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/option"
)

// SimulationInboundCheckDepositService contains methods and other services that
// help with interacting with the increase API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewSimulationInboundCheckDepositService] method instead.
type SimulationInboundCheckDepositService struct {
	Options []option.RequestOption
}

// NewSimulationInboundCheckDepositService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSimulationInboundCheckDepositService(opts ...option.RequestOption) (r *SimulationInboundCheckDepositService) {
	r = &SimulationInboundCheckDepositService{}
	r.Options = opts
	return
}

// Simulates an Inbound Check Deposit against your account. This imitates someone
// depositing a check at their bank that was issued from your account. It may or
// may not be associated with a Check Transfer. Increase will evaluate the Check
// Deposit as we would in production and either create a Transaction or a Declined
// Transaction as a result. You can inspect the resulting Inbound Check Deposit
// object to see the result.
func (r *SimulationInboundCheckDepositService) New(ctx context.Context, body SimulationInboundCheckDepositNewParams, opts ...option.RequestOption) (res *InboundCheckDeposit, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/inbound_check_deposits"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SimulationInboundCheckDepositNewParams struct {
	// The identifier of the Account Number the Inbound Check Deposit will be against.
	AccountNumberID param.Field[string] `json:"account_number_id,required"`
	// The check amount in cents.
	Amount param.Field[int64] `json:"amount,required"`
	// The check number on the check to be deposited.
	CheckNumber param.Field[string] `json:"check_number,required"`
}

func (r SimulationInboundCheckDepositNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

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

// SimulationFeePaymentService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimulationFeePaymentService] method instead.
type SimulationFeePaymentService struct {
	Options []option.RequestOption
}

// NewSimulationFeePaymentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSimulationFeePaymentService(opts ...option.RequestOption) (r *SimulationFeePaymentService) {
	r = &SimulationFeePaymentService{}
	r.Options = opts
	return
}

// A fee payment is how Increase charges you for technology fees incurred each
// month. In production, these payments will be charged to your program's billing
// account.
func (r *SimulationFeePaymentService) New(ctx context.Context, body SimulationFeePaymentNewParams, opts ...option.RequestOption) (res *Transaction, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/fee_payments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SimulationFeePaymentNewParams struct {
	// The identifier of the Account to use as the billing account for the fee payment.
	AccountID param.Field[string] `json:"account_id,required"`
	// The fee amount in cents. Must be positive.
	Amount param.Field[int64] `json:"amount,required"`
}

func (r SimulationFeePaymentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

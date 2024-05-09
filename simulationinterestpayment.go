// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"net/http"
	"time"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/option"
)

// SimulationInterestPaymentService contains methods and other services that help
// with interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimulationInterestPaymentService] method instead.
type SimulationInterestPaymentService struct {
	Options []option.RequestOption
}

// NewSimulationInterestPaymentService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSimulationInterestPaymentService(opts ...option.RequestOption) (r *SimulationInterestPaymentService) {
	r = &SimulationInterestPaymentService{}
	r.Options = opts
	return
}

// Simulates an interest payment to your account. In production, this happens
// automatically on the first of each month.
func (r *SimulationInterestPaymentService) New(ctx context.Context, body SimulationInterestPaymentNewParams, opts ...option.RequestOption) (res *Transaction, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/interest_payment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SimulationInterestPaymentNewParams struct {
	// The identifier of the Account Number the Interest Payment is for.
	AccountID param.Field[string] `json:"account_id,required"`
	// The interest amount in cents. Must be positive.
	Amount param.Field[int64] `json:"amount,required"`
	// The end of the interest period. If not provided, defaults to the current time.
	PeriodEnd param.Field[time.Time] `json:"period_end" format:"date-time"`
	// The start of the interest period. If not provided, defaults to the current time.
	PeriodStart param.Field[time.Time] `json:"period_start" format:"date-time"`
}

func (r SimulationInterestPaymentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

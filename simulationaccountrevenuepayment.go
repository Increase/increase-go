// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/Increase/increase-go/internal/apijson"
	"github.com/Increase/increase-go/internal/param"
	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
)

// SimulationAccountRevenuePaymentService contains methods and other services that
// help with interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimulationAccountRevenuePaymentService] method instead.
type SimulationAccountRevenuePaymentService struct {
	Options []option.RequestOption
}

// NewSimulationAccountRevenuePaymentService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSimulationAccountRevenuePaymentService(opts ...option.RequestOption) (r *SimulationAccountRevenuePaymentService) {
	r = &SimulationAccountRevenuePaymentService{}
	r.Options = opts
	return
}

// Simulates an account revenue payment to your account. In production, this
// happens automatically on the first of each month.
func (r *SimulationAccountRevenuePaymentService) New(ctx context.Context, body SimulationAccountRevenuePaymentNewParams, opts ...option.RequestOption) (res *Transaction, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "simulations/account_revenue_payments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type SimulationAccountRevenuePaymentNewParams struct {
	// The identifier of the Account the Account Revenue Payment should be paid to.
	AccountID param.Field[string] `json:"account_id" api:"required"`
	// The account revenue amount in cents. Must be positive.
	Amount param.Field[int64] `json:"amount" api:"required"`
	// The identifier of the Account the account revenue accrued on. Defaults to
	// `account_id`.
	AccruedOnAccountID param.Field[string] `json:"accrued_on_account_id"`
	// The end of the account revenue period. If not provided, defaults to the current
	// time.
	PeriodEnd param.Field[time.Time] `json:"period_end" format:"date-time"`
	// The start of the account revenue period. If not provided, defaults to the
	// current time.
	PeriodStart param.Field[time.Time] `json:"period_start" format:"date-time"`
}

func (r SimulationAccountRevenuePaymentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

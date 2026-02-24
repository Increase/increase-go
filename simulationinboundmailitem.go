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

// SimulationInboundMailItemService contains methods and other services that help
// with interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimulationInboundMailItemService] method instead.
type SimulationInboundMailItemService struct {
	Options []option.RequestOption
}

// NewSimulationInboundMailItemService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSimulationInboundMailItemService(opts ...option.RequestOption) (r *SimulationInboundMailItemService) {
	r = &SimulationInboundMailItemService{}
	r.Options = opts
	return
}

// Simulates an inbound mail item to your account, as if someone had mailed a
// physical check to one of your account's Lockboxes.
func (r *SimulationInboundMailItemService) New(ctx context.Context, body SimulationInboundMailItemNewParams, opts ...option.RequestOption) (res *InboundMailItem, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "simulations/inbound_mail_items"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SimulationInboundMailItemNewParams struct {
	// The amount of the check to be simulated, in cents.
	Amount param.Field[int64] `json:"amount" api:"required"`
	// The identifier of the Lockbox to simulate inbound mail to.
	LockboxID param.Field[string] `json:"lockbox_id" api:"required"`
	// The file containing the PDF contents. If not present, a default check image file
	// will be used.
	ContentsFileID param.Field[string] `json:"contents_file_id"`
}

func (r SimulationInboundMailItemNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

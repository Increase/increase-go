package increase

import (
	"context"
	"fmt"
	"net/http"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/field"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/option"
)

// SimulationCardDisputeService contains methods and other services that help with
// interacting with the increase API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSimulationCardDisputeService]
// method instead.
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
// will be reviewed. Since no review happens in sandbox, this endpoint simulates
// moving a Card Dispute into a rejected or accepted state. A Card Dispute can only
// be actioned one time and must have a status of `pending_reviewing`.
func (r *SimulationCardDisputeService) Action(ctx context.Context, card_dispute_id string, body SimulationCardDisputeActionParams, opts ...option.RequestOption) (res *CardDispute, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("simulations/card_disputes/%s/action", card_dispute_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SimulationCardDisputeActionParams struct {
	// The status to move the dispute to.
	Status field.Field[SimulationCardDisputeActionParamsStatus] `json:"status,required"`
	// Why the dispute was rejected. Not required for accepting disputes.
	Explanation field.Field[string] `json:"explanation"`
}

func (r SimulationCardDisputeActionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SimulationCardDisputeActionParamsStatus string

const (
	SimulationCardDisputeActionParamsStatusAccepted SimulationCardDisputeActionParamsStatus = "accepted"
	SimulationCardDisputeActionParamsStatusRejected SimulationCardDisputeActionParamsStatus = "rejected"
)

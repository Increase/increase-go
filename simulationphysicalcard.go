// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

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

// SimulationPhysicalCardService contains methods and other services that help with
// interacting with the increase API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSimulationPhysicalCardService]
// method instead.
type SimulationPhysicalCardService struct {
	Options []option.RequestOption
}

// NewSimulationPhysicalCardService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSimulationPhysicalCardService(opts ...option.RequestOption) (r *SimulationPhysicalCardService) {
	r = &SimulationPhysicalCardService{}
	r.Options = opts
	return
}

// This endpoint allows you to simulate advancing the shipment status of a Physical
// Card, to simulate e.g., that a physical card was attempted shipped but then
// failed delivery.
func (r *SimulationPhysicalCardService) ShipmentAdvance(ctx context.Context, physicalCardID string, body SimulationPhysicalCardShipmentAdvanceParams, opts ...option.RequestOption) (res *PhysicalCard, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("simulations/physical_cards/%s/shipment_advance", physicalCardID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SimulationPhysicalCardShipmentAdvanceParams struct {
	// The shipment status to move the Physical Card to.
	ShipmentStatus param.Field[SimulationPhysicalCardShipmentAdvanceParamsShipmentStatus] `json:"shipment_status,required"`
}

func (r SimulationPhysicalCardShipmentAdvanceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The shipment status to move the Physical Card to.
type SimulationPhysicalCardShipmentAdvanceParamsShipmentStatus string

const (
	// The physical card has not yet been shipped.
	SimulationPhysicalCardShipmentAdvanceParamsShipmentStatusPending SimulationPhysicalCardShipmentAdvanceParamsShipmentStatus = "pending"
	// The physical card shipment was canceled prior to submission.
	SimulationPhysicalCardShipmentAdvanceParamsShipmentStatusCanceled SimulationPhysicalCardShipmentAdvanceParamsShipmentStatus = "canceled"
	// The physical card shipment has been submitted to the card fulfillment provider.
	SimulationPhysicalCardShipmentAdvanceParamsShipmentStatusSubmitted SimulationPhysicalCardShipmentAdvanceParamsShipmentStatus = "submitted"
	// The physical card shipment has been acknowledged by the card fulfillment
	// provider and will be processed in their next batch.
	SimulationPhysicalCardShipmentAdvanceParamsShipmentStatusAcknowledged SimulationPhysicalCardShipmentAdvanceParamsShipmentStatus = "acknowledged"
	// The physical card shipment was rejected by the card printer due to an error.
	SimulationPhysicalCardShipmentAdvanceParamsShipmentStatusRejected SimulationPhysicalCardShipmentAdvanceParamsShipmentStatus = "rejected"
	// The physical card has been shipped.
	SimulationPhysicalCardShipmentAdvanceParamsShipmentStatusShipped SimulationPhysicalCardShipmentAdvanceParamsShipmentStatus = "shipped"
	// The physical card shipment was returned to the sender and destroyed by the
	// production facility.
	SimulationPhysicalCardShipmentAdvanceParamsShipmentStatusReturned SimulationPhysicalCardShipmentAdvanceParamsShipmentStatus = "returned"
)

func (r SimulationPhysicalCardShipmentAdvanceParamsShipmentStatus) IsKnown() bool {
	switch r {
	case SimulationPhysicalCardShipmentAdvanceParamsShipmentStatusPending, SimulationPhysicalCardShipmentAdvanceParamsShipmentStatusCanceled, SimulationPhysicalCardShipmentAdvanceParamsShipmentStatusSubmitted, SimulationPhysicalCardShipmentAdvanceParamsShipmentStatusAcknowledged, SimulationPhysicalCardShipmentAdvanceParamsShipmentStatusRejected, SimulationPhysicalCardShipmentAdvanceParamsShipmentStatusShipped, SimulationPhysicalCardShipmentAdvanceParamsShipmentStatusReturned:
		return true
	}
	return false
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/Increase/increase-go/internal/apijson"
	"github.com/Increase/increase-go/internal/param"
	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
)

// SimulationPhysicalCardService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimulationPhysicalCardService] method instead.
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

// This endpoint allows you to simulate receiving a tracking update for a Physical
// Card, to simulate the progress of a shipment.
func (r *SimulationPhysicalCardService) New(ctx context.Context, physicalCardID string, body SimulationPhysicalCardNewParams, opts ...option.RequestOption) (res *PhysicalCard, err error) {
	opts = slices.Concat(r.Options, opts)
	if physicalCardID == "" {
		err = errors.New("missing required physical_card_id parameter")
		return
	}
	path := fmt.Sprintf("simulations/physical_cards/%s/tracking_updates", physicalCardID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint allows you to simulate advancing the shipment status of a Physical
// Card, to simulate e.g., that a physical card was attempted shipped but then
// failed delivery.
func (r *SimulationPhysicalCardService) AdvanceShipment(ctx context.Context, physicalCardID string, body SimulationPhysicalCardAdvanceShipmentParams, opts ...option.RequestOption) (res *PhysicalCard, err error) {
	opts = slices.Concat(r.Options, opts)
	if physicalCardID == "" {
		err = errors.New("missing required physical_card_id parameter")
		return
	}
	path := fmt.Sprintf("simulations/physical_cards/%s/advance_shipment", physicalCardID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SimulationPhysicalCardNewParams struct {
	// The type of tracking event.
	Category param.Field[SimulationPhysicalCardNewParamsCategory] `json:"category" api:"required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time when the
	// carrier expects the card to be delivered.
	CarrierEstimatedDeliveryAt param.Field[time.Time] `json:"carrier_estimated_delivery_at" format:"date-time"`
	// The city where the event took place.
	City param.Field[string] `json:"city"`
	// The postal code where the event took place.
	PostalCode param.Field[string] `json:"postal_code"`
	// The state where the event took place.
	State param.Field[string] `json:"state"`
}

func (r SimulationPhysicalCardNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of tracking event.
type SimulationPhysicalCardNewParamsCategory string

const (
	SimulationPhysicalCardNewParamsCategoryInTransit            SimulationPhysicalCardNewParamsCategory = "in_transit"
	SimulationPhysicalCardNewParamsCategoryProcessedForDelivery SimulationPhysicalCardNewParamsCategory = "processed_for_delivery"
	SimulationPhysicalCardNewParamsCategoryDelivered            SimulationPhysicalCardNewParamsCategory = "delivered"
	SimulationPhysicalCardNewParamsCategoryDeliveryIssue        SimulationPhysicalCardNewParamsCategory = "delivery_issue"
	SimulationPhysicalCardNewParamsCategoryReturnedToSender     SimulationPhysicalCardNewParamsCategory = "returned_to_sender"
)

func (r SimulationPhysicalCardNewParamsCategory) IsKnown() bool {
	switch r {
	case SimulationPhysicalCardNewParamsCategoryInTransit, SimulationPhysicalCardNewParamsCategoryProcessedForDelivery, SimulationPhysicalCardNewParamsCategoryDelivered, SimulationPhysicalCardNewParamsCategoryDeliveryIssue, SimulationPhysicalCardNewParamsCategoryReturnedToSender:
		return true
	}
	return false
}

type SimulationPhysicalCardAdvanceShipmentParams struct {
	// The shipment status to move the Physical Card to.
	ShipmentStatus param.Field[SimulationPhysicalCardAdvanceShipmentParamsShipmentStatus] `json:"shipment_status" api:"required"`
}

func (r SimulationPhysicalCardAdvanceShipmentParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The shipment status to move the Physical Card to.
type SimulationPhysicalCardAdvanceShipmentParamsShipmentStatus string

const (
	SimulationPhysicalCardAdvanceShipmentParamsShipmentStatusPending           SimulationPhysicalCardAdvanceShipmentParamsShipmentStatus = "pending"
	SimulationPhysicalCardAdvanceShipmentParamsShipmentStatusCanceled          SimulationPhysicalCardAdvanceShipmentParamsShipmentStatus = "canceled"
	SimulationPhysicalCardAdvanceShipmentParamsShipmentStatusSubmitted         SimulationPhysicalCardAdvanceShipmentParamsShipmentStatus = "submitted"
	SimulationPhysicalCardAdvanceShipmentParamsShipmentStatusAcknowledged      SimulationPhysicalCardAdvanceShipmentParamsShipmentStatus = "acknowledged"
	SimulationPhysicalCardAdvanceShipmentParamsShipmentStatusRejected          SimulationPhysicalCardAdvanceShipmentParamsShipmentStatus = "rejected"
	SimulationPhysicalCardAdvanceShipmentParamsShipmentStatusShipped           SimulationPhysicalCardAdvanceShipmentParamsShipmentStatus = "shipped"
	SimulationPhysicalCardAdvanceShipmentParamsShipmentStatusReturned          SimulationPhysicalCardAdvanceShipmentParamsShipmentStatus = "returned"
	SimulationPhysicalCardAdvanceShipmentParamsShipmentStatusRequiresAttention SimulationPhysicalCardAdvanceShipmentParamsShipmentStatus = "requires_attention"
)

func (r SimulationPhysicalCardAdvanceShipmentParamsShipmentStatus) IsKnown() bool {
	switch r {
	case SimulationPhysicalCardAdvanceShipmentParamsShipmentStatusPending, SimulationPhysicalCardAdvanceShipmentParamsShipmentStatusCanceled, SimulationPhysicalCardAdvanceShipmentParamsShipmentStatusSubmitted, SimulationPhysicalCardAdvanceShipmentParamsShipmentStatusAcknowledged, SimulationPhysicalCardAdvanceShipmentParamsShipmentStatusRejected, SimulationPhysicalCardAdvanceShipmentParamsShipmentStatusShipped, SimulationPhysicalCardAdvanceShipmentParamsShipmentStatusReturned, SimulationPhysicalCardAdvanceShipmentParamsShipmentStatusRequiresAttention:
		return true
	}
	return false
}

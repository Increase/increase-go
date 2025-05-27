// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/Increase/increase-go"
	"github.com/Increase/increase-go/internal/testutil"
	"github.com/Increase/increase-go/option"
)

func TestSimulationPhysicalCardAdvanceShipment(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := increase.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Simulations.PhysicalCards.AdvanceShipment(
		context.TODO(),
		"physical_card_ode8duyq5v2ynhjoharl",
		increase.SimulationPhysicalCardAdvanceShipmentParams{
			ShipmentStatus: increase.F(increase.SimulationPhysicalCardAdvanceShipmentParamsShipmentStatusShipped),
		},
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSimulationPhysicalCardTrackingUpdatesWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := increase.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Simulations.PhysicalCards.TrackingUpdates(
		context.TODO(),
		"physical_card_ode8duyq5v2ynhjoharl",
		increase.SimulationPhysicalCardTrackingUpdatesParams{
			Category:                   increase.F(increase.SimulationPhysicalCardTrackingUpdatesParamsCategoryDelivered),
			CarrierEstimatedDeliveryAt: increase.F(time.Now()),
			City:                       increase.F("New York"),
			PostalCode:                 increase.F("10045"),
			State:                      increase.F("NY"),
		},
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

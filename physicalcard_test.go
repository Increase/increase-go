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

func TestPhysicalCardNewWithOptionalParams(t *testing.T) {
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
	_, err := client.PhysicalCards.New(context.TODO(), increase.PhysicalCardNewParams{
		CardID: increase.F("card_oubs0hwk5rn6knuecxg2"),
		Cardholder: increase.F(increase.PhysicalCardNewParamsCardholder{
			FirstName: increase.F("Ian"),
			LastName:  increase.F("Crease"),
		}),
		Shipment: increase.F(increase.PhysicalCardNewParamsShipment{
			Address: increase.F(increase.PhysicalCardNewParamsShipmentAddress{
				City:        increase.F("New York"),
				Line1:       increase.F("33 Liberty Street"),
				Line2:       increase.F("Unit 2"),
				Line3:       increase.F("x"),
				Name:        increase.F("Ian Crease"),
				PhoneNumber: increase.F("x"),
				PostalCode:  increase.F("10045"),
				State:       increase.F("NY"),
			}),
			Method: increase.F(increase.PhysicalCardNewParamsShipmentMethodUsps),
		}),
		PhysicalCardProfileID: increase.F("physical_card_profile_id"),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPhysicalCardGet(t *testing.T) {
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
	_, err := client.PhysicalCards.Get(context.TODO(), "physical_card_ode8duyq5v2ynhjoharl")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPhysicalCardUpdate(t *testing.T) {
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
	_, err := client.PhysicalCards.Update(
		context.TODO(),
		"physical_card_ode8duyq5v2ynhjoharl",
		increase.PhysicalCardUpdateParams{
			Status: increase.F(increase.PhysicalCardUpdateParamsStatusDisabled),
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

func TestPhysicalCardListWithOptionalParams(t *testing.T) {
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
	_, err := client.PhysicalCards.List(context.TODO(), increase.PhysicalCardListParams{
		CardID: increase.F("card_id"),
		CreatedAt: increase.F(increase.PhysicalCardListParamsCreatedAt{
			After:      increase.F(time.Now()),
			Before:     increase.F(time.Now()),
			OnOrAfter:  increase.F(time.Now()),
			OnOrBefore: increase.F(time.Now()),
		}),
		Cursor:         increase.F("cursor"),
		IdempotencyKey: increase.F("x"),
		Limit:          increase.F(int64(1)),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

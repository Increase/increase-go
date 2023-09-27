// File generated from our OpenAPI spec by Stainless.

package increase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/internal/testutil"
	"github.com/increase/increase-go/option"
)

func TestPhysicalCardNew(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := increase.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.PhysicalCards.New(context.TODO(), increase.PhysicalCardNewParams{
		CardID:        increase.F("string"),
		CardProfileID: increase.F("string"),
		Cardholder: increase.F(increase.PhysicalCardNewParamsCardholder{
			FirstName: increase.F("x"),
			LastName:  increase.F("x"),
		}),
		Shipment: increase.F(increase.PhysicalCardNewParamsShipment{
			Method: increase.F(increase.PhysicalCardNewParamsShipmentMethodUsps),
			Address: increase.F(increase.PhysicalCardNewParamsShipmentAddress{
				Name:        increase.F("x"),
				Line1:       increase.F("x"),
				Line2:       increase.F("x"),
				Line3:       increase.F("x"),
				PhoneNumber: increase.F("x"),
				City:        increase.F("x"),
				State:       increase.F("x"),
				PostalCode:  increase.F("x"),
			}),
		}),
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
	if !testutil.CheckTestServer(t) {
		return
	}
	client := increase.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
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
	if !testutil.CheckTestServer(t) {
		return
	}
	client := increase.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.PhysicalCards.Update(
		context.TODO(),
		"physical_card_ode8duyq5v2ynhjoharl",
		increase.PhysicalCardUpdateParams{
			Status: increase.F(increase.PhysicalCardUpdateParamsStatusActive),
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
	if !testutil.CheckTestServer(t) {
		return
	}
	client := increase.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.PhysicalCards.List(context.TODO(), increase.PhysicalCardListParams{
		CardID: increase.F("string"),
		CreatedAt: increase.F(increase.PhysicalCardListParamsCreatedAt{
			After:      increase.F(time.Now()),
			Before:     increase.F(time.Now()),
			OnOrAfter:  increase.F(time.Now()),
			OnOrBefore: increase.F(time.Now()),
		}),
		Cursor: increase.F("string"),
		Limit:  increase.F(int64(0)),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

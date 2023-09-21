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

func TestCardNewWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := increase.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.Cards.New(context.TODO(), increase.CardNewParams{
		AccountID: increase.F("string"),
		BillingAddress: increase.F(increase.CardNewParamsBillingAddress{
			Line1:      increase.F("x"),
			Line2:      increase.F("x"),
			City:       increase.F("x"),
			State:      increase.F("x"),
			PostalCode: increase.F("x"),
		}),
		Description: increase.F("x"),
		DigitalWallet: increase.F(increase.CardNewParamsDigitalWallet{
			Email:         increase.F("x"),
			Phone:         increase.F("x"),
			CardProfileID: increase.F("string"),
		}),
		EntityID: increase.F("string"),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardGet(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := increase.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.Cards.Get(context.TODO(), "card_oubs0hwk5rn6knuecxg2")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardUpdateWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := increase.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.Cards.Update(
		context.TODO(),
		"card_oubs0hwk5rn6knuecxg2",
		increase.CardUpdateParams{
			BillingAddress: increase.F(increase.CardUpdateParamsBillingAddress{
				Line1:      increase.F("x"),
				Line2:      increase.F("x"),
				City:       increase.F("x"),
				State:      increase.F("x"),
				PostalCode: increase.F("x"),
			}),
			Description: increase.F("x"),
			DigitalWallet: increase.F(increase.CardUpdateParamsDigitalWallet{
				Email:         increase.F("x"),
				Phone:         increase.F("x"),
				CardProfileID: increase.F("string"),
			}),
			Status: increase.F(increase.CardUpdateParamsStatusActive),
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

func TestCardListWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := increase.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.Cards.List(context.TODO(), increase.CardListParams{
		AccountID: increase.F("string"),
		CreatedAt: increase.F(increase.CardListParamsCreatedAt{
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

func TestCardGetSensitiveDetails(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := increase.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.Cards.GetSensitiveDetails(context.TODO(), "card_oubs0hwk5rn6knuecxg2")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

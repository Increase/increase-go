package increase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/option"
)

func TestCardNewWithOptionalParams(t *testing.T) {
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Cards.New(context.TODO(), increase.CardNewParams{AccountID: increase.F("account_in71c4amph0vgo2qllky"), Description: increase.F("Card for Ian Crease"), BillingAddress: increase.F(increase.CardNewParamsBillingAddress{Line1: increase.F("x"), Line2: increase.F("x"), City: increase.F("x"), State: increase.F("x"), PostalCode: increase.F("x")}), DigitalWallet: increase.F(increase.CardNewParamsDigitalWallet{Email: increase.F("x"), Phone: increase.F("x"), CardProfileID: increase.F("string")})})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardGet(t *testing.T) {
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Cards.Get(
		context.TODO(),
		"card_oubs0hwk5rn6knuecxg2",
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardUpdateWithOptionalParams(t *testing.T) {
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Cards.Update(
		context.TODO(),
		"card_oubs0hwk5rn6knuecxg2",
		increase.CardUpdateParams{Description: increase.F("New description"), Status: increase.F(increase.CardUpdateParamsStatusActive), BillingAddress: increase.F(increase.CardUpdateParamsBillingAddress{Line1: increase.F("x"), Line2: increase.F("x"), City: increase.F("x"), State: increase.F("x"), PostalCode: increase.F("x")}), DigitalWallet: increase.F(increase.CardUpdateParamsDigitalWallet{Email: increase.F("x"), Phone: increase.F("x"), CardProfileID: increase.F("string")})},
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardListWithOptionalParams(t *testing.T) {
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Cards.List(context.TODO(), increase.CardListParams{Cursor: increase.F("string"), Limit: increase.F(int64(0)), AccountID: increase.F("string"), CreatedAt: increase.F(increase.CardListParamsCreatedAt{After: increase.F(time.Now()), Before: increase.F(time.Now()), OnOrAfter: increase.F(time.Now()), OnOrBefore: increase.F(time.Now())})})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardGetSensitiveDetails(t *testing.T) {
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Cards.GetSensitiveDetails(
		context.TODO(),
		"card_oubs0hwk5rn6knuecxg2",
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

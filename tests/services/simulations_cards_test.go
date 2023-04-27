package services

import (
	"context"
	"errors"
	"testing"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
)

func TestSimulationCardAuthorizeWithOptionalParams(t *testing.T) {
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Simulations.Cards.Authorize(context.TODO(), requests.SimulationCardAuthorizeParams{Amount: increase.F(int64(1000)), CardID: increase.F("card_oubs0hwk5rn6knuecxg2"), DigitalWalletTokenID: increase.F("string"), EventSubscriptionID: increase.F("event_subscription_001dzz0r20rcdxgb013zqb8m04g")})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSimulationCardSettlementWithOptionalParams(t *testing.T) {
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Simulations.Cards.Settlement(context.TODO(), requests.SimulationCardSettlementParams{CardID: increase.F("card_oubs0hwk5rn6knuecxg2"), PendingTransactionID: increase.F("pending_transaction_k1sfetcau2qbvjbzgju4"), Amount: increase.F(int64(1))})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

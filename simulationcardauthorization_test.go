// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/Increase/increase-go"
	"github.com/Increase/increase-go/internal/testutil"
	"github.com/Increase/increase-go/option"
)

func TestSimulationCardAuthorizationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Simulations.CardAuthorizations.New(context.TODO(), increase.SimulationCardAuthorizationNewParams{
		Amount:                     increase.F(int64(1000)),
		AuthenticatedCardPaymentID: increase.F("authenticated_card_payment_id"),
		CardID:                     increase.F("card_oubs0hwk5rn6knuecxg2"),
		DeclineReason:              increase.F(increase.SimulationCardAuthorizationNewParamsDeclineReasonAccountClosed),
		DigitalWalletTokenID:       increase.F("digital_wallet_token_id"),
		EventSubscriptionID:        increase.F("event_subscription_001dzz0r20rcdxgb013zqb8m04g"),
		MerchantAcceptorID:         increase.F("5665270011000168"),
		MerchantCategoryCode:       increase.F("5734"),
		MerchantCity:               increase.F("New York"),
		MerchantCountry:            increase.F("US"),
		MerchantDescriptor:         increase.F("AMAZON.COM"),
		MerchantState:              increase.F("NY"),
		NetworkDetails: increase.F(increase.SimulationCardAuthorizationNewParamsNetworkDetails{
			Visa: increase.F(increase.SimulationCardAuthorizationNewParamsNetworkDetailsVisa{
				StandInProcessingReason: increase.F(increase.SimulationCardAuthorizationNewParamsNetworkDetailsVisaStandInProcessingReasonIssuerError),
			}),
		}),
		NetworkRiskScore: increase.F(int64(0)),
		PhysicalCardID:   increase.F("physical_card_id"),
		ProcessingCategory: increase.F(increase.SimulationCardAuthorizationNewParamsProcessingCategory{
			Category: increase.F(increase.SimulationCardAuthorizationNewParamsProcessingCategoryCategoryAccountFunding),
		}),
		TerminalID: increase.F("x"),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

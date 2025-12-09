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

func TestSimulationCardBalanceInquiryNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Simulations.CardBalanceInquiries.New(context.TODO(), increase.SimulationCardBalanceInquiryNewParams{
		Balance:              increase.F(int64(1000000)),
		CardID:               increase.F("card_oubs0hwk5rn6knuecxg2"),
		DeclineReason:        increase.F(increase.SimulationCardBalanceInquiryNewParamsDeclineReasonAccountClosed),
		DigitalWalletTokenID: increase.F("digital_wallet_token_id"),
		EventSubscriptionID:  increase.F("event_subscription_001dzz0r20rcdxgb013zqb8m04g"),
		MerchantAcceptorID:   increase.F("5665270011000168"),
		MerchantCategoryCode: increase.F("5734"),
		MerchantCity:         increase.F("New York"),
		MerchantCountry:      increase.F("US"),
		MerchantDescriptor:   increase.F("CITIBANK"),
		MerchantState:        increase.F("NY"),
		NetworkDetails: increase.F(increase.SimulationCardBalanceInquiryNewParamsNetworkDetails{
			Visa: increase.F(increase.SimulationCardBalanceInquiryNewParamsNetworkDetailsVisa{
				StandInProcessingReason: increase.F(increase.SimulationCardBalanceInquiryNewParamsNetworkDetailsVisaStandInProcessingReasonIssuerError),
			}),
		}),
		NetworkRiskScore: increase.F(int64(0)),
		PhysicalCardID:   increase.F("physical_card_id"),
		TerminalID:       increase.F("x"),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

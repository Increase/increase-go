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

func TestSimulationCardTokenNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Simulations.CardTokens.New(context.TODO(), increase.SimulationCardTokenNewParams{
		Capabilities: increase.F([]increase.SimulationCardTokenNewParamsCapability{{
			CrossBorderPushTransfers: increase.F(increase.SimulationCardTokenNewParamsCapabilitiesCrossBorderPushTransfersSupported),
			DomesticPushTransfers:    increase.F(increase.SimulationCardTokenNewParamsCapabilitiesDomesticPushTransfersSupported),
			Route:                    increase.F(increase.SimulationCardTokenNewParamsCapabilitiesRouteVisa),
		}}),
		Expiration:                 increase.F(time.Now()),
		Last4:                      increase.F("1234"),
		Prefix:                     increase.F("41234567"),
		PrimaryAccountNumberLength: increase.F(int64(16)),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

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

func TestSimulationCardDisputeActionWithOptionalParams(t *testing.T) {
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
	_, err := client.Simulations.CardDisputes.Action(
		context.TODO(),
		"card_dispute_h9sc95nbl1cgltpp7men",
		increase.SimulationCardDisputeActionParams{
			Network: increase.F(increase.SimulationCardDisputeActionParamsNetworkVisa),
			Visa: increase.F(increase.SimulationCardDisputeActionParamsVisa{
				Action:                        increase.F(increase.SimulationCardDisputeActionParamsVisaActionAcceptUserSubmission),
				AcceptChargeback:              increase.F(increase.SimulationCardDisputeActionParamsVisaAcceptChargeback{}),
				AcceptUserSubmission:          increase.F(increase.SimulationCardDisputeActionParamsVisaAcceptUserSubmission{}),
				DeclineUserPrearbitration:     increase.F(increase.SimulationCardDisputeActionParamsVisaDeclineUserPrearbitration{}),
				ReceiveMerchantPrearbitration: increase.F(increase.SimulationCardDisputeActionParamsVisaReceiveMerchantPrearbitration{}),
				Represent:                     increase.F(increase.SimulationCardDisputeActionParamsVisaRepresent{}),
				RequestFurtherInformation: increase.F(increase.SimulationCardDisputeActionParamsVisaRequestFurtherInformation{
					Reason: increase.F("x"),
				}),
				TimeOutChargeback:             increase.F(increase.SimulationCardDisputeActionParamsVisaTimeOutChargeback{}),
				TimeOutMerchantPrearbitration: increase.F(increase.SimulationCardDisputeActionParamsVisaTimeOutMerchantPrearbitration{}),
				TimeOutRepresentment:          increase.F(increase.SimulationCardDisputeActionParamsVisaTimeOutRepresentment{}),
				TimeOutUserPrearbitration:     increase.F(increase.SimulationCardDisputeActionParamsVisaTimeOutUserPrearbitration{}),
			}),
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

package increase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/option"
)

func TestSimulationCardDisputeActionWithOptionalParams(t *testing.T) {
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Simulations.CardDisputes.Action(
		context.TODO(),
		"card_dispute_h9sc95nbl1cgltpp7men",
		increase.SimulationCardDisputeActionParams{
			Status:      increase.F(increase.SimulationCardDisputeActionParamsStatusAccepted),
			Explanation: increase.F("x"),
		},
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

package increase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/option"
)

func TestSimulationInterestPaymentNew(t *testing.T) {
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Simulations.InterestPayments.New(context.TODO(), increase.SimulationInterestPaymentNewParams{
		AccountID: increase.F("string"),
		Amount:    increase.F(int64(1)),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

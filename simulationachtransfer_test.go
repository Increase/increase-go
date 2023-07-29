// File generated from our OpenAPI spec by Stainless.

package increase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/internal/testutil"
	"github.com/increase/increase-go/option"
)

func TestSimulationACHTransferNewInboundWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := increase.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.Simulations.ACHTransfers.NewInbound(context.TODO(), increase.SimulationACHTransferNewInboundParams{
		AccountNumberID:          increase.F("string"),
		Amount:                   increase.F(int64(0)),
		CompanyDescriptiveDate:   increase.F("x"),
		CompanyDiscretionaryData: increase.F("x"),
		CompanyEntryDescription:  increase.F("x"),
		CompanyID:                increase.F("x"),
		CompanyName:              increase.F("x"),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSimulationACHTransferReturnWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	t.Skip("Prism incorrectly returns an invalid JSON error")
	client := increase.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.Simulations.ACHTransfers.Return(
		context.TODO(),
		"ach_transfer_uoxatyh3lt5evrsdvo7q",
		increase.SimulationACHTransferReturnParams{
			Reason: increase.F(increase.SimulationACHTransferReturnParamsReasonInsufficientFund),
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

func TestSimulationACHTransferSubmit(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	t.Skip("Prism incorrectly returns an invalid JSON error")
	client := increase.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.Simulations.ACHTransfers.Submit(context.TODO(), "ach_transfer_uoxatyh3lt5evrsdvo7q")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

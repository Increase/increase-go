// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/internal/testutil"
	"github.com/increase/increase-go/option"
)

func TestSimulationACHTransferNewInboundWithOptionalParams(t *testing.T) {
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
	_, err := client.Simulations.ACHTransfers.NewInbound(context.TODO(), increase.SimulationACHTransferNewInboundParams{
		AccountNumberID:          increase.F("account_number_v18nkfqm6afpsrvy82b2"),
		Amount:                   increase.F(int64(1000)),
		CompanyDescriptiveDate:   increase.F("x"),
		CompanyDiscretionaryData: increase.F("x"),
		CompanyEntryDescription:  increase.F("x"),
		CompanyID:                increase.F("x"),
		CompanyName:              increase.F("x"),
		ReceiverIDNumber:         increase.F("x"),
		ReceiverName:             increase.F("x"),
		ResolveAt:                increase.F(time.Now()),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSimulationACHTransferNotificationOfChange(t *testing.T) {
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
	_, err := client.Simulations.ACHTransfers.NotificationOfChange(
		context.TODO(),
		"ach_transfer_uoxatyh3lt5evrsdvo7q",
		increase.SimulationACHTransferNotificationOfChangeParams{
			ChangeCode:    increase.F(increase.SimulationACHTransferNotificationOfChangeParamsChangeCodeIncorrectRoutingNumber),
			CorrectedData: increase.F("123456789"),
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

func TestSimulationACHTransferReturnWithOptionalParams(t *testing.T) {
	t.Skip("Prism incorrectly returns an invalid JSON error")
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
	t.Skip("Prism incorrectly returns an invalid JSON error")
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
	_, err := client.Simulations.ACHTransfers.Submit(context.TODO(), "ach_transfer_uoxatyh3lt5evrsdvo7q")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

package services

import (
	"context"
	"errors"
	"net/http/httputil"
	"testing"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/core"
	"github.com/increase/increase-go/fields"
	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/requests"
)

func TestACHTransfersNewInboundWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Simulations.ACHTransfers.NewInbound(context.TODO(), &requests.SimulateAnACHTransferToYourAccountParameters{AccountNumberID: fields.F("account_number_v18nkfqm6afpsrvy82b2"), Amount: fields.F(int64(1000)), CompanyDescriptiveDate: fields.F("x"), CompanyDiscretionaryData: fields.F("x"), CompanyEntryDescription: fields.F("x"), CompanyName: fields.F("x"), CompanyID: fields.F("x")})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestACHTransfersReturnWithOptionalParams(t *testing.T) {
	t.Skip("Prism incorrectly returns an invalid JSON error")
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Simulations.ACHTransfers.Return(
		context.TODO(),
		"ach_transfer_uoxatyh3lt5evrsdvo7q",
		&requests.ReturnASandboxACHTransferParameters{Reason: fields.F(requests.ReturnASandboxACHTransferParametersReasonInsufficientFund)},
	)
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestACHTransfersSubmit(t *testing.T) {
	t.Skip("Prism incorrectly returns an invalid JSON error")
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Simulations.ACHTransfers.Submit(
		context.TODO(),
		"ach_transfer_uoxatyh3lt5evrsdvo7q",
	)
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

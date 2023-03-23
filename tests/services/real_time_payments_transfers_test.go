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

func TestRealTimePaymentsTransfersNewInboundWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Simulations.RealTimePaymentsTransfers.NewInbound(context.TODO(), &requests.SimulateARealTimePaymentsTransferToYourAccountParameters{AccountNumberID: fields.F("account_number_v18nkfqm6afpsrvy82b2"), Amount: fields.F(int64(1000)), RequestForPaymentID: fields.F("real_time_payments_request_for_payment_28kcliz1oevcnqyn9qp7"), DebtorName: fields.F("x"), DebtorAccountNumber: fields.F("x"), DebtorRoutingNumber: fields.F("xxxxxxxxx"), RemittanceInformation: fields.F("x")})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

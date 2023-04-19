package services

import (
	"context"
	"errors"
	"net/http/httputil"
	"testing"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/core"
	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
)

func TestRealTimePaymentsTransfersCompleteWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Simulations.RealTimePaymentsTransfers.Complete(
		context.TODO(),
		"real_time_payments_transfer_iyuhl5kdn7ssmup83mvq",
		&requests.RealTimePaymentsTransferCompleteParams{Rejection: increase.F(requests.RealTimePaymentsTransferCompleteParamsRejection{RejectReasonCode: increase.F(requests.RealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeAccountClosed)})},
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

func TestRealTimePaymentsTransfersNewInboundWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Simulations.RealTimePaymentsTransfers.NewInbound(context.TODO(), &requests.RealTimePaymentsTransferNewInboundParams{AccountNumberID: increase.F("account_number_v18nkfqm6afpsrvy82b2"), Amount: increase.F(int64(1000)), RequestForPaymentID: increase.F("real_time_payments_request_for_payment_28kcliz1oevcnqyn9qp7"), DebtorName: increase.F("x"), DebtorAccountNumber: increase.F("x"), DebtorRoutingNumber: increase.F("xxxxxxxxx"), RemittanceInformation: increase.F("x")})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

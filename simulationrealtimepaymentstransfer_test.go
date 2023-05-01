package increase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/option"
)

func TestSimulationRealTimePaymentsTransferCompleteWithOptionalParams(t *testing.T) {
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Simulations.RealTimePaymentsTransfers.Complete(
		context.TODO(),
		"real_time_payments_transfer_iyuhl5kdn7ssmup83mvq",
		increase.SimulationRealTimePaymentsTransferCompleteParams{Rejection: increase.F(increase.SimulationRealTimePaymentsTransferCompleteParamsRejection{RejectReasonCode: increase.F(increase.SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeAccountClosed)})},
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSimulationRealTimePaymentsTransferNewInboundWithOptionalParams(t *testing.T) {
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Simulations.RealTimePaymentsTransfers.NewInbound(context.TODO(), increase.SimulationRealTimePaymentsTransferNewInboundParams{AccountNumberID: increase.F("account_number_v18nkfqm6afpsrvy82b2"), Amount: increase.F(int64(1000)), RequestForPaymentID: increase.F("real_time_payments_request_for_payment_28kcliz1oevcnqyn9qp7"), DebtorName: increase.F("x"), DebtorAccountNumber: increase.F("x"), DebtorRoutingNumber: increase.F("xxxxxxxxx"), RemittanceInformation: increase.F("x")})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

package services

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
)

func TestRealTimePaymentsTransferNewWithOptionalParams(t *testing.T) {
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.RealTimePaymentsTransfers.New(context.TODO(), requests.RealTimePaymentsTransferNewParams{SourceAccountNumberID: increase.F("account_number_v18nkfqm6afpsrvy82b2"), DestinationAccountNumber: increase.F("987654321"), DestinationRoutingNumber: increase.F("101050001"), ExternalAccountID: increase.F("string"), Amount: increase.F(int64(100)), CreditorName: increase.F("Ian Crease"), RemittanceInformation: increase.F("Invoice 29582"), RequireApproval: increase.F(true)})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRealTimePaymentsTransferGet(t *testing.T) {
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.RealTimePaymentsTransfers.Get(
		context.TODO(),
		"real_time_payments_transfer_iyuhl5kdn7ssmup83mvq",
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRealTimePaymentsTransferListWithOptionalParams(t *testing.T) {
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.RealTimePaymentsTransfers.List(context.TODO(), requests.RealTimePaymentsTransferListParams{Cursor: increase.F("string"), Limit: increase.F(int64(0)), AccountID: increase.F("string"), ExternalAccountID: increase.F("string"), CreatedAt: increase.F(requests.RealTimePaymentsTransferListParamsCreatedAt{After: increase.F(time.Now()), Before: increase.F(time.Now()), OnOrAfter: increase.F(time.Now()), OnOrBefore: increase.F(time.Now())})})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

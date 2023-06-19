// File generated from our OpenAPI spec by Stainless.

package increase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/internal/testutil"
	"github.com/increase/increase-go/option"
)

func TestRealTimePaymentsTransferNewWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := increase.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := client.RealTimePaymentsTransfers.New(context.TODO(), increase.RealTimePaymentsTransferNewParams{
		Amount:                   increase.F(int64(1)),
		CreditorName:             increase.F("x"),
		RemittanceInformation:    increase.F("x"),
		SourceAccountNumberID:    increase.F("string"),
		DestinationAccountNumber: increase.F("x"),
		DestinationRoutingNumber: increase.F("xxxxxxxxx"),
		ExternalAccountID:        increase.F("string"),
		RequireApproval:          increase.F(true),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRealTimePaymentsTransferGet(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := increase.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := client.RealTimePaymentsTransfers.Get(context.TODO(), "real_time_payments_transfer_iyuhl5kdn7ssmup83mvq")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRealTimePaymentsTransferListWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := increase.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := client.RealTimePaymentsTransfers.List(context.TODO(), increase.RealTimePaymentsTransferListParams{
		AccountID: increase.F("string"),
		CreatedAt: increase.F(increase.RealTimePaymentsTransferListParamsCreatedAt{
			After:      increase.F(time.Now()),
			Before:     increase.F(time.Now()),
			OnOrAfter:  increase.F(time.Now()),
			OnOrBefore: increase.F(time.Now()),
		}),
		Cursor:            increase.F("string"),
		ExternalAccountID: increase.F("string"),
		Limit:             increase.F(int64(0)),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

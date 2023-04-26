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

func TestACHTransferNewWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.ACHTransfers.New(context.TODO(), &requests.ACHTransferNewParams{AccountID: increase.F("account_in71c4amph0vgo2qllky"), AccountNumber: increase.F("987654321"), Addendum: increase.F("x"), Amount: increase.F(int64(100)), CompanyDescriptiveDate: increase.F("x"), CompanyDiscretionaryData: increase.F("x"), CompanyEntryDescription: increase.F("x"), CompanyName: increase.F("x"), EffectiveDate: increase.F(time.Now()), ExternalAccountID: increase.F("string"), Funding: increase.F(requests.ACHTransferNewParamsFundingChecking), IndividualID: increase.F("x"), IndividualName: increase.F("x"), RequireApproval: increase.F(true), RoutingNumber: increase.F("101050001"), StandardEntryClassCode: increase.F(requests.ACHTransferNewParamsStandardEntryClassCodeCorporateCreditOrDebit), StatementDescriptor: increase.F("New ACH transfer")})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestACHTransferGet(t *testing.T) {
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.ACHTransfers.Get(
		context.TODO(),
		"ach_transfer_uoxatyh3lt5evrsdvo7q",
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestACHTransferListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.ACHTransfers.List(context.TODO(), &requests.ACHTransferListParams{Cursor: increase.F("string"), Limit: increase.F(int64(0)), AccountID: increase.F("string"), ExternalAccountID: increase.F("string"), CreatedAt: increase.F(requests.ACHTransferListParamsCreatedAt{After: increase.F(time.Now()), Before: increase.F(time.Now()), OnOrAfter: increase.F(time.Now()), OnOrBefore: increase.F(time.Now())})})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestACHTransferApprove(t *testing.T) {
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.ACHTransfers.Approve(
		context.TODO(),
		"ach_transfer_uoxatyh3lt5evrsdvo7q",
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestACHTransferCancel(t *testing.T) {
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.ACHTransfers.Cancel(
		context.TODO(),
		"ach_transfer_uoxatyh3lt5evrsdvo7q",
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

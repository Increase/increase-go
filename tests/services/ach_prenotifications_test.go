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

func TestACHPrenotificationNewWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.ACHPrenotifications.New(context.TODO(), &requests.ACHPrenotificationNewParams{AccountNumber: increase.F("987654321"), Addendum: increase.F("x"), CompanyDescriptiveDate: increase.F("x"), CompanyDiscretionaryData: increase.F("x"), CompanyEntryDescription: increase.F("x"), CompanyName: increase.F("x"), CreditDebitIndicator: increase.F(requests.ACHPrenotificationNewParamsCreditDebitIndicatorCredit), EffectiveDate: increase.F(time.Now()), IndividualID: increase.F("x"), IndividualName: increase.F("x"), RoutingNumber: increase.F("101050001"), StandardEntryClassCode: increase.F(requests.ACHPrenotificationNewParamsStandardEntryClassCodeCorporateCreditOrDebit)})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestACHPrenotificationGet(t *testing.T) {
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.ACHPrenotifications.Get(
		context.TODO(),
		"ach_prenotification_ubjf9qqsxl3obbcn1u34",
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestACHPrenotificationListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.ACHPrenotifications.List(context.TODO(), &requests.ACHPrenotificationListParams{Cursor: increase.F("string"), Limit: increase.F(int64(0)), CreatedAt: increase.F(requests.ACHPrenotificationListParamsCreatedAt{After: increase.F(time.Now()), Before: increase.F(time.Now()), OnOrAfter: increase.F(time.Now()), OnOrBefore: increase.F(time.Now())})})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

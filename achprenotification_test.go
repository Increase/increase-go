package increase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/option"
)

func TestACHPrenotificationNewWithOptionalParams(t *testing.T) {
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.ACHPrenotifications.New(context.TODO(), increase.ACHPrenotificationNewParams{
		AccountNumber:            increase.F("x"),
		RoutingNumber:            increase.F("xxxxxxxxx"),
		Addendum:                 increase.F("x"),
		CompanyDescriptiveDate:   increase.F("x"),
		CompanyDiscretionaryData: increase.F("x"),
		CompanyEntryDescription:  increase.F("x"),
		CompanyName:              increase.F("x"),
		CreditDebitIndicator:     increase.F(increase.ACHPrenotificationNewParamsCreditDebitIndicatorCredit),
		EffectiveDate:            increase.F(time.Now()),
		IndividualID:             increase.F("x"),
		IndividualName:           increase.F("x"),
		StandardEntryClassCode:   increase.F(increase.ACHPrenotificationNewParamsStandardEntryClassCodeCorporateCreditOrDebit),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestACHPrenotificationGet(t *testing.T) {
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
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
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.ACHPrenotifications.List(context.TODO(), increase.ACHPrenotificationListParams{
		CreatedAt: increase.F(increase.ACHPrenotificationListParamsCreatedAt{After: increase.F(time.Now()), Before: increase.F(time.Now()), OnOrAfter: increase.F(time.Now()), OnOrBefore: increase.F(time.Now())}),
		Cursor:    increase.F("string"),
		Limit:     increase.F(int64(0)),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

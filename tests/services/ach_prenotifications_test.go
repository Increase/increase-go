package services

import (
	"context"
	"errors"
	"net/http/httputil"
	"testing"
	"time"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/core"
	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
)

func TestACHPrenotificationsNewWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.ACHPrenotifications.New(context.TODO(), &requests.ACHPrenotificationNewParams{AccountNumber: increase.F("987654321"), Addendum: increase.F("x"), CompanyDescriptiveDate: increase.F("x"), CompanyDiscretionaryData: increase.F("x"), CompanyEntryDescription: increase.F("x"), CompanyName: increase.F("x"), CreditDebitIndicator: increase.F(requests.ACHPrenotificationNewParamsCreditDebitIndicatorCredit), EffectiveDate: increase.F(time.Now()), IndividualID: increase.F("x"), IndividualName: increase.F("x"), RoutingNumber: increase.F("101050001"), StandardEntryClassCode: increase.F(requests.ACHPrenotificationNewParamsStandardEntryClassCodeCorporateCreditOrDebit)})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestACHPrenotificationsGet(t *testing.T) {
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.ACHPrenotifications.Get(
		context.TODO(),
		"ach_prenotification_ubjf9qqsxl3obbcn1u34",
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

func TestACHPrenotificationsListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.ACHPrenotifications.List(context.TODO(), &requests.ACHPrenotificationListParams{Cursor: increase.F("string"), Limit: increase.F(int64(0)), CreatedAt: increase.F(requests.ACHPrenotificationListParamsCreatedAt{After: increase.F(time.Now()), Before: increase.F(time.Now()), OnOrAfter: increase.F(time.Now()), OnOrBefore: increase.F(time.Now())})})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

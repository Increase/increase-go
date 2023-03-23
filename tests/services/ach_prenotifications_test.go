package services

import (
	"context"
	"errors"
	"net/http/httputil"
	"testing"
	"time"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/core"
	"github.com/increase/increase-go/fields"
	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/requests"
)

func TestACHPrenotificationsNewWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.ACHPrenotifications.New(context.TODO(), &requests.CreateAnACHPrenotificationParameters{AccountNumber: fields.F("987654321"), Addendum: fields.F("x"), CompanyDescriptiveDate: fields.F("x"), CompanyDiscretionaryData: fields.F("x"), CompanyEntryDescription: fields.F("x"), CompanyName: fields.F("x"), CreditDebitIndicator: fields.F(requests.CreateAnACHPrenotificationParametersCreditDebitIndicatorCredit), EffectiveDate: fields.F(time.Now()), IndividualID: fields.F("x"), IndividualName: fields.F("x"), RoutingNumber: fields.F("101050001"), StandardEntryClassCode: fields.F(requests.CreateAnACHPrenotificationParametersStandardEntryClassCodeCorporateCreditOrDebit)})
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
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
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
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.ACHPrenotifications.List(context.TODO(), &requests.ACHPrenotificationListParams{Cursor: fields.F("string"), Limit: fields.F(int64(0)), CreatedAt: fields.F(requests.ACHPrenotificationListParamsCreatedAt{After: fields.F(time.Now()), Before: fields.F(time.Now()), OnOrAfter: fields.F(time.Now()), OnOrBefore: fields.F(time.Now())})})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

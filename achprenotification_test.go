// File generated from our OpenAPI spec by Stainless.

package increase_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/internal/testutil"
	"github.com/increase/increase-go/option"
)

func TestACHPrenotificationNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := increase.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.ACHPrenotifications.New(context.TODO(), increase.ACHPrenotificationNewParams{
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
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestACHPrenotificationGet(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := increase.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.ACHPrenotifications.Get(context.TODO(), "ach_prenotification_ubjf9qqsxl3obbcn1u34")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestACHPrenotificationListWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := increase.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.ACHPrenotifications.List(context.TODO(), increase.ACHPrenotificationListParams{
		CreatedAt: increase.F(increase.ACHPrenotificationListParamsCreatedAt{
			After:      increase.F(time.Now()),
			Before:     increase.F(time.Now()),
			OnOrAfter:  increase.F(time.Now()),
			OnOrBefore: increase.F(time.Now()),
		}),
		Cursor: increase.F("string"),
		Limit:  increase.F(int64(0)),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

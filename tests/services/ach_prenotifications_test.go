package services

import (
	"context"
	"testing"
	"time"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/types"
)

func TestACHPrenotificationsNewWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.ACHPrenotifications.New(context.TODO(), &types.CreateAnACHPrenotificationParameters{AccountNumber: increase.P("987654321"), Addendum: increase.P("x"), CompanyDescriptiveDate: increase.P("x"), CompanyDiscretionaryData: increase.P("x"), CompanyEntryDescription: increase.P("x"), CompanyName: increase.P("x"), CreditDebitIndicator: increase.P(types.CreateAnACHPrenotificationParametersCreditDebitIndicatorCredit), EffectiveDate: increase.P(time.Now()), IndividualID: increase.P("x"), IndividualName: increase.P("x"), RoutingNumber: increase.P("101050001"), StandardEntryClassCode: increase.P(types.CreateAnACHPrenotificationParametersStandardEntryClassCodeCorporateCreditOrDebit)})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestACHPrenotificationsGet(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.ACHPrenotifications.Get(
		context.TODO(),
		"ach_prenotification_ubjf9qqsxl3obbcn1u34",
	)
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestACHPrenotificationsListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.ACHPrenotifications.List(context.TODO(), &types.ACHPrenotificationListParams{Cursor: increase.P("string"), Limit: increase.P(int64(0)), CreatedAt: increase.P(types.ACHPrenotificationListParamsCreatedAt{After: increase.P(time.Now()), Before: increase.P(time.Now()), OnOrAfter: increase.P(time.Now()), OnOrBefore: increase.P(time.Now())})})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

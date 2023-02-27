package services

import (
	"context"
	"increase"
	"increase/options"
	"increase/types"
	"testing"
)

func TestACHPrenotificationsNewWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.ACHPrenotifications.New(context.TODO(), &types.CreateAnACHPrenotificationParameters{AccountNumber: increase.P("987654321"), Addendum: increase.P("x"), CompanyDescriptiveDate: increase.P("x"), CompanyDiscretionaryData: increase.P("x"), CompanyEntryDescription: increase.P("x"), CompanyName: increase.P("x"), CreditDebitIndicator: increase.P(types.CreateAnACHPrenotificationParametersCreditDebitIndicatorCredit), EffectiveDate: increase.P("2019-12-27"), IndividualID: increase.P("x"), IndividualName: increase.P("x"), RoutingNumber: increase.P("101050001"), StandardEntryClassCode: increase.P(types.CreateAnACHPrenotificationParametersStandardEntryClassCodeCorporateCreditOrDebit)})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestACHPrenotificationsGet(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.ACHPrenotifications.Get(context.TODO(), "ach_prenotification_ubjf9qqsxl3obbcn1u34")
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestACHPrenotificationsListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.ACHPrenotifications.List(context.TODO(), &types.ACHPrenotificationListParams{Cursor: increase.P("string"), Limit: increase.P(int64(0)), CreatedAt: &types.ACHPrenotificationsListParamsCreatedAt{After: increase.P("2019-12-27T18:11:19.117Z"), Before: increase.P("2019-12-27T18:11:19.117Z"), OnOrAfter: increase.P("2019-12-27T18:11:19.117Z"), OnOrBefore: increase.P("2019-12-27T18:11:19.117Z")}})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

package increase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/option"
)

func TestExternalAccountNewWithOptionalParams(t *testing.T) {
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.ExternalAccounts.New(context.TODO(), increase.ExternalAccountNewParams{
		AccountNumber: increase.F("x"),
		Description:   increase.F("x"),
		RoutingNumber: increase.F("xxxxxxxxx"),
		Funding:       increase.F(increase.ExternalAccountNewParamsFundingChecking),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExternalAccountGet(t *testing.T) {
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.ExternalAccounts.Get(
		context.TODO(),
		"external_account_ukk55lr923a3ac0pp7iv",
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExternalAccountUpdateWithOptionalParams(t *testing.T) {
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.ExternalAccounts.Update(
		context.TODO(),
		"external_account_ukk55lr923a3ac0pp7iv",
		increase.ExternalAccountUpdateParams{
			Description: increase.F("x"),
			Status:      increase.F(increase.ExternalAccountUpdateParamsStatusActive),
		},
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExternalAccountListWithOptionalParams(t *testing.T) {
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.ExternalAccounts.List(context.TODO(), increase.ExternalAccountListParams{
		Cursor: increase.F("string"),
		Limit:  increase.F(int64(0)),
		Status: increase.F(increase.ExternalAccountListParamsStatus{In: increase.F([]increase.ExternalAccountListParamsStatusIn{increase.ExternalAccountListParamsStatusInActive, increase.ExternalAccountListParamsStatusInActive, increase.ExternalAccountListParamsStatusInActive})}),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

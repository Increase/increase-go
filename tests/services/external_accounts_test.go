package services

import (
	"context"
	"errors"
	"net/http/httputil"
	"testing"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/core"
	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
)

func TestExternalAccountNewWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.ExternalAccounts.New(context.TODO(), &requests.ExternalAccountNewParams{RoutingNumber: increase.F("101050001"), AccountNumber: increase.F("987654321"), Funding: increase.F(requests.ExternalAccountNewParamsFundingChecking), Description: increase.F("Landlord")})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExternalAccountGet(t *testing.T) {
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.ExternalAccounts.Get(
		context.TODO(),
		"external_account_ukk55lr923a3ac0pp7iv",
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

func TestExternalAccountUpdateWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.ExternalAccounts.Update(
		context.TODO(),
		"external_account_ukk55lr923a3ac0pp7iv",
		&requests.ExternalAccountUpdateParams{Description: increase.F("New description"), Status: increase.F(requests.ExternalAccountUpdateParamsStatusActive)},
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

func TestExternalAccountListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.ExternalAccounts.List(context.TODO(), &requests.ExternalAccountListParams{Cursor: increase.F("string"), Limit: increase.F(int64(0)), Status: increase.F(requests.ExternalAccountListParamsStatus{In: increase.F([]requests.ExternalAccountListParamsStatusIn{requests.ExternalAccountListParamsStatusInActive, requests.ExternalAccountListParamsStatusInActive, requests.ExternalAccountListParamsStatusInActive})})})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

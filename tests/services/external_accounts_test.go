package services

import (
	"context"
	"errors"
	"net/http/httputil"
	"testing"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/core"
	"github.com/increase/increase-go/fields"
	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/requests"
)

func TestExternalAccountsNewWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.ExternalAccounts.New(context.TODO(), &requests.CreateAnExternalAccountParameters{RoutingNumber: fields.F("101050001"), AccountNumber: fields.F("987654321"), Funding: fields.F(requests.CreateAnExternalAccountParametersFundingChecking), Description: fields.F("Landlord")})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExternalAccountsGet(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
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

func TestExternalAccountsUpdateWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.ExternalAccounts.Update(
		context.TODO(),
		"external_account_ukk55lr923a3ac0pp7iv",
		&requests.UpdateAnExternalAccountParameters{Description: fields.F("New description"), Status: fields.F(requests.UpdateAnExternalAccountParametersStatusActive)},
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

func TestExternalAccountsListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.ExternalAccounts.List(context.TODO(), &requests.ExternalAccountListParams{Cursor: fields.F("string"), Limit: fields.F(int64(0)), Status: fields.F(requests.ExternalAccountListParamsStatus{In: fields.F([]requests.ExternalAccountListParamsStatusIn{requests.ExternalAccountListParamsStatusInActive, requests.ExternalAccountListParamsStatusInActive, requests.ExternalAccountListParamsStatusInActive})})})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

package services

import (
	"context"
	"increase"
	"increase/options"
	"increase/types"
	"testing"
)

func TestExternalAccountsNewWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.ExternalAccounts.New(context.TODO(), &types.CreateAnExternalAccountParameters{RoutingNumber: increase.P("101050001"), AccountNumber: increase.P("987654321"), Funding: increase.P(types.CreateAnExternalAccountParametersFundingChecking), Description: increase.P("Landlord")})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestExternalAccountsGet(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.ExternalAccounts.Get(context.TODO(), "external_account_ukk55lr923a3ac0pp7iv")
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestExternalAccountsUpdateWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.ExternalAccounts.Update(
		context.TODO(),
		"external_account_ukk55lr923a3ac0pp7iv",
		&types.UpdateAnExternalAccountParameters{Description: increase.P("New description"), Status: increase.P(types.UpdateAnExternalAccountParametersStatusActive)},
	)
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestExternalAccountsListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.ExternalAccounts.List(context.TODO(), &types.ExternalAccountListParams{Cursor: increase.P("string"), Limit: increase.P(int64(0)), Status: &types.ExternalAccountsListParamsStatus{In: &[]types.ExternalAccountsListParamsStatusIn{types.ExternalAccountsListParamsStatusInActive, types.ExternalAccountsListParamsStatusInActive, types.ExternalAccountsListParamsStatusInActive}}})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

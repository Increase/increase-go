// File generated from our OpenAPI spec by Stainless.

package increase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/internal/testutil"
	"github.com/increase/increase-go/option"
)

func TestAccountNewWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := increase.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := client.Accounts.New(context.TODO(), increase.AccountNewParams{
		Name:                  increase.F("x"),
		EntityID:              increase.F("string"),
		InformationalEntityID: increase.F("string"),
		ProgramID:             increase.F("string"),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountGet(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := increase.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := client.Accounts.Get(context.TODO(), "account_in71c4amph0vgo2qllky")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountUpdateWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := increase.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := client.Accounts.Update(
		context.TODO(),
		"account_in71c4amph0vgo2qllky",
		increase.AccountUpdateParams{
			Name: increase.F("x"),
		},
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountListWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := increase.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := client.Accounts.List(context.TODO(), increase.AccountListParams{
		CreatedAt: increase.F(increase.AccountListParamsCreatedAt{
			After:      increase.F(time.Now()),
			Before:     increase.F(time.Now()),
			OnOrAfter:  increase.F(time.Now()),
			OnOrBefore: increase.F(time.Now()),
		}),
		Cursor:                increase.F("string"),
		EntityID:              increase.F("string"),
		InformationalEntityID: increase.F("string"),
		Limit:                 increase.F(int64(0)),
		Status:                increase.F(increase.AccountListParamsStatusOpen),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountClose(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	t.Skip("Prism tests are broken")
	client := increase.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := client.Accounts.Close(context.TODO(), "account_in71c4amph0vgo2qllky")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

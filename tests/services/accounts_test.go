package services

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
)

func TestAccountNewWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Accounts.New(context.TODO(), &requests.AccountNewParams{EntityID: increase.F("string"), ProgramID: increase.F("string"), InformationalEntityID: increase.F("string"), Name: increase.F("New Account!")})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountGet(t *testing.T) {
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Accounts.Get(
		context.TODO(),
		"account_in71c4amph0vgo2qllky",
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountUpdateWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Accounts.Update(
		context.TODO(),
		"account_in71c4amph0vgo2qllky",
		&requests.AccountUpdateParams{Name: increase.F("My renamed account")},
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Accounts.List(context.TODO(), &requests.AccountListParams{Cursor: increase.F("string"), Limit: increase.F(int64(0)), EntityID: increase.F("string"), Status: increase.F(requests.AccountListParamsStatusOpen), CreatedAt: increase.F(requests.AccountListParamsCreatedAt{After: increase.F(time.Now()), Before: increase.F(time.Now()), OnOrAfter: increase.F(time.Now()), OnOrBefore: increase.F(time.Now())})})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountClose(t *testing.T) {
	t.Skip("Prism tests are broken")
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Accounts.Close(
		context.TODO(),
		"account_in71c4amph0vgo2qllky",
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

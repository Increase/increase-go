package services

import (
	"context"
	"errors"
	"testing"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
)

func TestBookkeepingAccountNewWithOptionalParams(t *testing.T) {
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.BookkeepingAccounts.New(context.TODO(), requests.BookkeepingAccountNewParams{ComplianceCategory: increase.F(requests.BookkeepingAccountNewParamsComplianceCategoryCommingledCash), EntityID: increase.F("string"), AccountID: increase.F("string"), Name: increase.F("New Account!")})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBookkeepingAccountListWithOptionalParams(t *testing.T) {
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.BookkeepingAccounts.List(context.TODO(), requests.BookkeepingAccountListParams{Cursor: increase.F("string"), Limit: increase.F(int64(0))})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

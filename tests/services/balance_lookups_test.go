package services

import (
	"context"
	"errors"
	"testing"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
)

func TestBalanceLookupLookup(t *testing.T) {
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.BalanceLookups.Lookup(context.TODO(), &requests.BalanceLookupLookupParams{AccountID: increase.F("account_in71c4amph0vgo2qllky")})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

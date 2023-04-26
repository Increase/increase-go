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

func TestDocumentGet(t *testing.T) {
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Documents.Get(
		context.TODO(),
		"document_qjtqc6s4c14ve2q89izm",
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDocumentListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Documents.List(context.TODO(), &requests.DocumentListParams{Cursor: increase.F("string"), Limit: increase.F(int64(0)), EntityID: increase.F("string"), Category: increase.F(requests.DocumentListParamsCategory{In: increase.F([]requests.DocumentListParamsCategoryIn{requests.DocumentListParamsCategoryInAccountOpeningDisclosures, requests.DocumentListParamsCategoryInAccountOpeningDisclosures, requests.DocumentListParamsCategoryInAccountOpeningDisclosures})}), CreatedAt: increase.F(requests.DocumentListParamsCreatedAt{After: increase.F(time.Now()), Before: increase.F(time.Now()), OnOrAfter: increase.F(time.Now()), OnOrBefore: increase.F(time.Now())})})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

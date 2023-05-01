package increase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/option"
)

func TestDocumentGet(t *testing.T) {
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
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
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Documents.List(context.TODO(), increase.DocumentListParams{Cursor: increase.F("string"), Limit: increase.F(int64(0)), EntityID: increase.F("string"), Category: increase.F(increase.DocumentListParamsCategory{In: increase.F([]increase.DocumentListParamsCategoryIn{increase.DocumentListParamsCategoryInAccountOpeningDisclosures, increase.DocumentListParamsCategoryInAccountOpeningDisclosures, increase.DocumentListParamsCategoryInAccountOpeningDisclosures})}), CreatedAt: increase.F(increase.DocumentListParamsCreatedAt{After: increase.F(time.Now()), Before: increase.F(time.Now()), OnOrAfter: increase.F(time.Now()), OnOrBefore: increase.F(time.Now())})})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

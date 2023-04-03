package services

import (
	"context"
	"errors"
	"net/http/httputil"
	"testing"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/core"
	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/requests"
)

func TestSupplementalDocumentsNew(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Entities.SupplementalDocuments.New(
		context.TODO(),
		"entity_n8y8tnk2p9339ti393yi",
		&requests.CreateASupplementalDocumentForAnEntityParameters{FileID: increase.F("file_makxrc67oh9l6sg7w9yc")},
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

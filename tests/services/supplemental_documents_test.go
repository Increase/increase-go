package services

import (
	"context"
	"increase"
	"increase/options"
	"increase/types"
	"testing"
)

func TestSupplementalDocumentsNew(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("something1234"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Entities.SupplementalDocuments.New(
		context.TODO(),
		"entity_n8y8tnk2p9339ti393yi",
		&types.CreateASupplementalDocumentForAnEntityParameters{FileID: increase.P("file_makxrc67oh9l6sg7w9yc")},
	)
	if err == nil {
		t.Fatal("err should not be nil", err)
	}
}

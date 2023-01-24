package services

import (
	"context"
	"increase"
	"testing"

	client "increase"
	"increase/types"
)

func TestSupplementalDocumentsCreate(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.SupplementalDocuments.Entities.Create(
		context.TODO(),
		"entity_n8y8tnk2p9339ti393yi",
		&types.CreateASupplementalDocumentForAnEntityParameters{FileID: increase.P("file_makxrc67oh9l6sg7w9yc")},
	)
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}

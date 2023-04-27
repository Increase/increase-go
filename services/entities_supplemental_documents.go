package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type EntitySupplementalDocumentService struct {
	Options []option.RequestOption
}

func NewEntitySupplementalDocumentService(opts ...option.RequestOption) (r *EntitySupplementalDocumentService) {
	r = &EntitySupplementalDocumentService{}
	r.Options = opts
	return
}

// Create a supplemental document for an Entity
func (r *EntitySupplementalDocumentService) New(ctx context.Context, entity_id string, body requests.EntitySupplementalDocumentNewParams, opts ...option.RequestOption) (res *responses.Entity, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("entities/%s/supplemental_documents", entity_id)
	err = option.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

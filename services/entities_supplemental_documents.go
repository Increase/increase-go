package services

import (
	"context"
	"fmt"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type EntitiesSupplementalDocumentService struct {
	Options []option.RequestOption
}

func NewEntitiesSupplementalDocumentService(opts ...option.RequestOption) (r *EntitiesSupplementalDocumentService) {
	r = &EntitiesSupplementalDocumentService{}
	r.Options = opts
	return
}

// Create a supplemental document for an Entity
func (r *EntitiesSupplementalDocumentService) New(ctx context.Context, entity_id string, body *requests.SupplementalDocumentNewParams, opts ...option.RequestOption) (res *responses.Entity, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("entities/%s/supplemental_documents", entity_id)
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

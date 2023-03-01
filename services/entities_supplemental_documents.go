package services

import (
	"context"
	"fmt"
	"increase/options"
	"increase/types"
)

type EntitiesSupplementalDocumentService struct {
	Options []options.RequestOption
}

func NewEntitiesSupplementalDocumentService(opts ...options.RequestOption) (r *EntitiesSupplementalDocumentService) {
	r = &EntitiesSupplementalDocumentService{}
	r.Options = opts
	return
}

// Create a supplemental document for an Entity
func (r *EntitiesSupplementalDocumentService) New(ctx context.Context, entity_id string, body *types.CreateASupplementalDocumentForAnEntityParameters, opts ...options.RequestOption) (res *types.Entity, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("entities/%s/supplemental_documents", entity_id)
	err = options.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

package services

import (
	"context"
	"fmt"
	"increase/options"
	"increase/types"
	"net/url"
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
	u, err := url.Parse(fmt.Sprintf("entities/%s/supplemental_documents", entity_id))
	if err != nil {
		return
	}
	cfg, err := options.NewRequestConfig(ctx, "POST", u, body, opts...)
	if err != nil {
		return
	}
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return
	}

	return
}

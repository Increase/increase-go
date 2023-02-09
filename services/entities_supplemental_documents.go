package services

import (
	"bytes"
	"context"
	"fmt"
	"increase/options"
	"increase/types"
	"io"
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
	b, err := body.MarshalJSON()
	content := io.NopCloser(bytes.NewBuffer(b))
	u, err := url.Parse(fmt.Sprintf("entities/%s/supplemental_documents", entity_id))
	if err != nil {
		return
	}
	cfg := options.NewRequestConfig(ctx, "POST", u, content, opts...)
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return
	}

	return
}

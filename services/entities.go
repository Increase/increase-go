package services

import (
	"context"
	"fmt"
	"increase/options"
	"increase/pagination"
	"increase/types"
	"net/url"
)

type EntityService struct {
	Options               []options.RequestOption
	SupplementalDocuments *EntitiesSupplementalDocumentService
}

func NewEntityService(opts ...options.RequestOption) (r *EntityService) {
	r = &EntityService{}
	r.Options = opts
	r.SupplementalDocuments = NewEntitiesSupplementalDocumentService(opts...)
	return
}

// Create an Entity
func (r *EntityService) New(ctx context.Context, body *types.CreateAnEntityParameters, opts ...options.RequestOption) (res *types.Entity, err error) {
	opts = append(r.Options, opts...)
	u, err := url.Parse(fmt.Sprintf("entities"))
	if err != nil {
		return
	}
	cfg := options.NewRequestConfig(ctx, "POST", u, opts...)
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return
	}

	return
}

// Retrieve an Entity
func (r *EntityService) Get(ctx context.Context, entity_id string, opts ...options.RequestOption) (res *types.Entity, err error) {
	opts = append(r.Options, opts...)
	u, err := url.Parse(fmt.Sprintf("entities/%s", entity_id))
	if err != nil {
		return
	}
	cfg := options.NewRequestConfig(ctx, "GET", u, opts...)
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return
	}

	return
}

// List Entities
func (r *EntityService) List(ctx context.Context, query *types.EntityListParams, opts ...options.RequestOption) (res *types.EntitiesPage, err error) {
	u, err := url.Parse(fmt.Sprintf("entities"))
	if err != nil {
		return
	}
	opts = append(r.Options, opts...)
	cfg := options.NewRequestConfig(ctx, "GET", u, opts...)
	res = &types.EntitiesPage{
		Page: &pagination.Page[types.Entity]{
			Config:  *cfg,
			Options: opts,
		},
	}
	err = res.Fire()
	return
}

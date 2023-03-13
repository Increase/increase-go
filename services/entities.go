package services

import (
	"context"
	"fmt"

	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/pagination"
	"github.com/increase/increase-go/types"
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
	opts = append(r.Options[:], opts...)
	path := "entities"
	err = options.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Retrieve an Entity
func (r *EntityService) Get(ctx context.Context, entity_id string, opts ...options.RequestOption) (res *types.Entity, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("entities/%s", entity_id)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// List Entities
func (r *EntityService) List(ctx context.Context, query *types.EntityListParams, opts ...options.RequestOption) (res *types.EntitiesPage, err error) {
	opts = append(r.Options, opts...)
	path := "entities"
	cfg, err := options.NewRequestConfig(ctx, "GET", path, query, nil, opts...)
	if err != nil {
		return
	}
	res = &types.EntitiesPage{
		Page: &pagination.Page[types.Entity]{
			Config:  *cfg,
			Options: opts,
		},
	}
	return res, res.Fire()
}

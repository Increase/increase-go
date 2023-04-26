package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type EntityService struct {
	Options               []option.RequestOption
	SupplementalDocuments *EntitySupplementalDocumentService
}

func NewEntityService(opts ...option.RequestOption) (r *EntityService) {
	r = &EntityService{}
	r.Options = opts
	r.SupplementalDocuments = NewEntitySupplementalDocumentService(opts...)
	return
}

// Create an Entity
func (r *EntityService) New(ctx context.Context, body *requests.EntityNewParams, opts ...option.RequestOption) (res *responses.Entity, err error) {
	opts = append(r.Options[:], opts...)
	path := "entities"
	err = option.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve an Entity
func (r *EntityService) Get(ctx context.Context, entity_id string, opts ...option.RequestOption) (res *responses.Entity, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("entities/%s", entity_id)
	err = option.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Entities
func (r *EntityService) List(ctx context.Context, query *requests.EntityListParams, opts ...option.RequestOption) (res *responses.Page[responses.Entity], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "entities"
	cfg, err := option.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List Entities
func (r *EntityService) ListAutoPager(ctx context.Context, query *requests.EntityListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.Entity] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}

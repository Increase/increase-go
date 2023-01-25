package services

import (
	"context"
	"fmt"
	"increase/core"
	"increase/pagination"
	"increase/types"
)

type EntityService struct {
	Requester             core.Requester
	get                   func(context.Context, string, *core.CoreRequest, interface{}) error
	post                  func(context.Context, string, *core.CoreRequest, interface{}) error
	patch                 func(context.Context, string, *core.CoreRequest, interface{}) error
	put                   func(context.Context, string, *core.CoreRequest, interface{}) error
	delete                func(context.Context, string, *core.CoreRequest, interface{}) error
	SupplementalDocuments *EntitiesSupplementalDocumentService
}

func NewEntityService(requester core.Requester) (r *EntityService) {
	r = &EntityService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	r.SupplementalDocuments = NewEntitiesSupplementalDocumentService(requester)
	return
}

// Create an Entity
func (r *EntityService) Create(ctx context.Context, body *types.CreateAnEntityParameters, opts ...*core.RequestOpts) (res *types.Entity, err error) {
	err = r.post(
		ctx,
		"/entities",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)

	return
}

// Retrieve an Entity
func (r *EntityService) Retrieve(ctx context.Context, entity_id string, opts ...*core.RequestOpts) (res *types.Entity, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/entities/%s", entity_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)

	return
}

// List Entities
func (r *EntityService) List(ctx context.Context, query *types.ListEntitiesQuery, opts ...*core.RequestOpts) (res *types.EntitiesPage, err error) {
	page := &types.EntitiesPage{
		Page: &pagination.Page[types.Entity]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/entities",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

package services

import (
	"context"
	"fmt"
	"increase/options"
	"increase/pagination"
	"increase/types"
)

type DocumentService struct {
	Options []options.RequestOption
}

func NewDocumentService(opts ...options.RequestOption) (r *DocumentService) {
	r = &DocumentService{}
	r.Options = opts
	return
}

// Retrieve a Document
func (r *DocumentService) Get(ctx context.Context, document_id string, opts ...options.RequestOption) (res *types.Document, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("documents/%s", document_id)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// List Documents
func (r *DocumentService) List(ctx context.Context, query *types.DocumentListParams, opts ...options.RequestOption) (res *types.DocumentsPage, err error) {
	opts = append(r.Options, opts...)
	path := "documents"
	cfg, err := options.NewRequestConfig(ctx, "GET", path, query, nil, opts...)
	if err != nil {
		return
	}
	res = &types.DocumentsPage{
		Page: &pagination.Page[types.Document]{
			Config:  *cfg,
			Options: opts,
		},
	}
	return res, res.Fire()
}

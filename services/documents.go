package services

import (
	"context"
	"fmt"

	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/pagination"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
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
func (r *DocumentService) Get(ctx context.Context, document_id string, opts ...options.RequestOption) (res *responses.Document, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("documents/%s", document_id)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// List Documents
func (r *DocumentService) List(ctx context.Context, query *requests.DocumentListParams, opts ...options.RequestOption) (res *responses.DocumentsPage, err error) {
	opts = append(r.Options, opts...)
	path := "documents"
	cfg, err := options.NewRequestConfig(ctx, "GET", path, query, nil, opts...)
	if err != nil {
		return
	}
	res = &responses.DocumentsPage{
		Page: &pagination.Page[responses.Document]{
			Config:  *cfg,
			Options: opts,
		},
	}
	return res, res.Fire()
}

package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type DocumentService struct {
	Options []option.RequestOption
}

func NewDocumentService(opts ...option.RequestOption) (r *DocumentService) {
	r = &DocumentService{}
	r.Options = opts
	return
}

// Retrieve a Document
func (r *DocumentService) Get(ctx context.Context, document_id string, opts ...option.RequestOption) (res *responses.Document, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("documents/%s", document_id)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// List Documents
func (r *DocumentService) List(ctx context.Context, query *requests.DocumentListParams, opts ...option.RequestOption) (res *responses.Page[responses.Document], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "documents"
	cfg, err := option.NewRequestConfig(ctx, "GET", path, query, &res, opts...)
	if err != nil {
		return
	}
	err = cfg.Execute()
	if err != nil {
		return
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List Documents
func (r *DocumentService) ListAutoPager(ctx context.Context, query *requests.DocumentListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.Document] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}

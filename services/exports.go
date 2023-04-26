package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type ExportService struct {
	Options []option.RequestOption
}

func NewExportService(opts ...option.RequestOption) (r *ExportService) {
	r = &ExportService{}
	r.Options = opts
	return
}

// Create an Export
func (r *ExportService) New(ctx context.Context, body *requests.ExportNewParams, opts ...option.RequestOption) (res *responses.Export, err error) {
	opts = append(r.Options[:], opts...)
	path := "exports"
	err = option.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve an Export
func (r *ExportService) Get(ctx context.Context, export_id string, opts ...option.RequestOption) (res *responses.Export, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("exports/%s", export_id)
	err = option.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Exports
func (r *ExportService) List(ctx context.Context, query *requests.ExportListParams, opts ...option.RequestOption) (res *responses.Page[responses.Export], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "exports"
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

// List Exports
func (r *ExportService) ListAutoPager(ctx context.Context, query *requests.ExportListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.Export] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}

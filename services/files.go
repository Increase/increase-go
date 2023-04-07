package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type FileService struct {
	Options []option.RequestOption
}

func NewFileService(opts ...option.RequestOption) (r *FileService) {
	r = &FileService{}
	r.Options = opts
	return
}

// To upload a file to Increase, you'll need to send a request of Content-Type
// `multipart/form-data`. The request should contain the file you would like to
// upload, as well as the parameters for creating a file.
func (r *FileService) New(ctx context.Context, body *requests.FileNewParams, opts ...option.RequestOption) (res *responses.File, err error) {
	opts = append(r.Options[:], opts...)
	path := "files"
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Retrieve a File
func (r *FileService) Get(ctx context.Context, file_id string, opts ...option.RequestOption) (res *responses.File, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("files/%s", file_id)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// List Files
func (r *FileService) List(ctx context.Context, query *requests.FileListParams, opts ...option.RequestOption) (res *responses.Page[responses.File], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "files"
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

// List Files
func (r *FileService) ListAutoPager(ctx context.Context, query *requests.FileListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.File] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}

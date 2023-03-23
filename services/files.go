package services

import (
	"context"
	"fmt"

	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/pagination"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type FileService struct {
	Options []options.RequestOption
}

func NewFileService(opts ...options.RequestOption) (r *FileService) {
	r = &FileService{}
	r.Options = opts
	return
}

// To upload a file to Increase, you'll need to send a request of Content-Type
// `multipart/form-data`. The request should contain the file you would like to
// upload, as well as the parameters for creating a file.
func (r *FileService) New(ctx context.Context, body *requests.CreateAFileParameters, opts ...options.RequestOption) (res *responses.File, err error) {
	opts = append(r.Options[:], opts...)
	path := "files"
	err = options.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Retrieve a File
func (r *FileService) Get(ctx context.Context, file_id string, opts ...options.RequestOption) (res *responses.File, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("files/%s", file_id)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// List Files
func (r *FileService) List(ctx context.Context, query *requests.FileListParams, opts ...options.RequestOption) (res *responses.FilesPage, err error) {
	opts = append(r.Options, opts...)
	path := "files"
	cfg, err := options.NewRequestConfig(ctx, "GET", path, query, nil, opts...)
	if err != nil {
		return
	}
	res = &responses.FilesPage{
		Page: &pagination.Page[responses.File]{
			Config:  *cfg,
			Options: opts,
		},
	}
	return res, res.Fire()
}

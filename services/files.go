package services

import (
	"context"
	"fmt"
	"increase/core"
	"increase/pagination"
	"increase/types"
)

type FileService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewFileService(requester core.Requester) (r *FileService) {
	r = &FileService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

// To upload a file to Increase, you'll need to send a request of Content-Type
// `multipart/form-data`. The request should contain the file you would like to
// upload, as well as the parameters for creating a file.
func (r *FileService) New(ctx context.Context, body *types.CreateAFileParameters, opts ...*core.RequestOpts) (res *types.File, err error) {
	path := "/files"
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
		Body:   body,
	}
	err = r.post(ctx, path, req, &res)

	return
}

// Retrieve a File
func (r *FileService) Get(ctx context.Context, file_id string, opts ...*core.RequestOpts) (res *types.File, err error) {
	path := fmt.Sprintf("/files/%s", file_id)
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
	}
	err = r.get(ctx, path, req, &res)

	return
}

// List Files
func (r *FileService) List(ctx context.Context, query *types.FileListParams, opts ...*core.RequestOpts) (res *types.FilesPage, err error) {
	page := &types.FilesPage{
		Page: &pagination.Page[types.File]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/files",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

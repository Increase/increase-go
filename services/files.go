package services

import "context"
import "increase/core"
import "increase/types"
import "fmt"
import "increase/pagination"

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
func (r *FileService) Create(ctx context.Context, body *types.CreateAFileParameters, opts ...*core.RequestOpts) (res *types.File, err error) {
	err = r.post(
		ctx,
		"/files",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)

	return
}

func (r *FileService) Retrieve(ctx context.Context, file_id string, opts ...*core.RequestOpts) (res *types.File, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/files/%s", file_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)

	return
}

func (r *FileService) List(ctx context.Context, query *types.ListFilesQuery, opts ...*core.RequestOpts) (res *types.FilesPage, err error) {
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

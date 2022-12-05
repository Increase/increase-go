package files

import "increase/core"
import "fmt"

type Files struct {
	Requester core.Requester
	get       func(string, *core.CoreRequest, interface{}) error
	post      func(string, *core.CoreRequest, interface{}) error
	patch     func(string, *core.CoreRequest, interface{}) error
	put       func(string, *core.CoreRequest, interface{}) error
	delete    func(string, *core.CoreRequest, interface{}) error
}

func NewFiles(requster core.Requester) (r *Files) {
	r = &Files{}
	r.Requester = requster
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
func (r *Files) Create(body FilesCreateParams, opts ...*core.RequestOpts) (res File, err error) {
	err = r.post(
		"/files",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *Files) Retrieve(file_id string, opts ...*core.RequestOpts) (res File, err error) {
	err = r.get(
		fmt.Sprintf("/files/%s", file_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *Files) List(query *FilesListQuery, opts ...*core.RequestOpts) (res FilesListResponse, err error) {
	err = r.get(
		"/files",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Query:  query,
		},
		&res,
	)
	return
}

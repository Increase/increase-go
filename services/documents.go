package services

import (
	"context"
	"fmt"
	"increase/core"
	"increase/pagination"
	"increase/types"
)

type DocumentService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewDocumentService(requester core.Requester) (r *DocumentService) {
	r = &DocumentService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

func (r *DocumentService) Retrieve(ctx context.Context, document_id string, opts ...*core.RequestOpts) (res *types.Document, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/documents/%s", document_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)

	return
}

func (r *DocumentService) List(ctx context.Context, query *types.ListDocumentsQuery, opts ...*core.RequestOpts) (res *types.DocumentsPage, err error) {
	page := &types.DocumentsPage{
		Page: &pagination.Page[types.Document]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/documents",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

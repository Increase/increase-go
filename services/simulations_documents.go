package services

import (
	"context"
	"increase/core"
	"increase/types"
)

type SimulationsDocumentService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewSimulationsDocumentService(requester core.Requester) (r *SimulationsDocumentService) {
	r = &SimulationsDocumentService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

// Simulates an tax document being created for an account.
func (r *SimulationsDocumentService) New(ctx context.Context, body *types.SimulateATaxDocumentBeingCreatedParameters, opts ...*core.RequestOpts) (res *types.Document, err error) {
	path := "/simulations/documents"
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
		Body:   body,
	}
	err = r.post(ctx, path, req, &res)

	return
}

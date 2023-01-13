package entities

import "context"
import "increase/core"
import "increase/shared"
import "fmt"

type SupplementalDocumentService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewSupplementalDocumentService(requester core.Requester) (r *SupplementalDocumentService) {
	r = &SupplementalDocumentService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

type PreloadedSupplementalDocumentService struct {
	SupplementalDocuments *SupplementalDocumentService
}

func (r *PreloadedSupplementalDocumentService) Init(service *SupplementalDocumentService) {
	r.SupplementalDocuments = service
}

func NewPreloadedSupplementalDocumentService(service *SupplementalDocumentService) (r *PreloadedSupplementalDocumentService) {
	r = &PreloadedSupplementalDocumentService{}
	r.Init(service)
	return
}

func (r *SupplementalDocumentService) Create(ctx context.Context, entity_id string, body *CreateASupplementalDocumentForAnEntityParameters, opts ...*core.RequestOpts) (res *shared.Entity, err error) {
	err = r.post(
		ctx,
		fmt.Sprintf("/entities/%s/supplemental_documents", entity_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *PreloadedSupplementalDocumentService) Create(ctx context.Context, entity_id string, body *CreateASupplementalDocumentForAnEntityParameters, opts ...*core.RequestOpts) (res *shared.Entity, err error) {
	err = r.SupplementalDocuments.post(
		ctx,
		fmt.Sprintf("/entities/%s/supplemental_documents", entity_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

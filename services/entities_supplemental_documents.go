package services

import "context"
import "increase/core"
import "increase/types"
import "fmt"

type EntitiesSupplementalDocumentService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewEntitiesSupplementalDocumentService(requester core.Requester) (r *EntitiesSupplementalDocumentService) {
	r = &EntitiesSupplementalDocumentService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

func (r *EntitiesSupplementalDocumentService) Create(ctx context.Context, entity_id string, body *types.CreateASupplementalDocumentForAnEntityParameters, opts ...*core.RequestOpts) (res *types.Entity, err error) {
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

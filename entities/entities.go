package entities

import "increase/core"
import "fmt"

type Entities struct {
	Requester core.Requester
	get       func(string, *core.CoreRequest, interface{}) error
	post      func(string, *core.CoreRequest, interface{}) error
	patch     func(string, *core.CoreRequest, interface{}) error
	put       func(string, *core.CoreRequest, interface{}) error
	delete    func(string, *core.CoreRequest, interface{}) error
}

func NewEntities(requster core.Requester) (r *Entities) {
	r = &Entities{}
	r.Requester = requster
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

func (r *Entities) Create(body EntitiesCreateParams, opts ...*core.RequestOpts) (res Entity, err error) {
	err = r.post(
		"/entities",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *Entities) Retrieve(entity_id string, opts ...*core.RequestOpts) (res Entity, err error) {
	err = r.get(
		fmt.Sprintf("/entities/%s", entity_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *Entities) List(query *EntitiesListQuery, opts ...*core.RequestOpts) (res EntitiesListResponse, err error) {
	err = r.get(
		"/entities",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Query:  query,
		},
		&res,
	)
	return
}

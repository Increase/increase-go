package card_profiles

import "increase/core"
import "fmt"

type CardProfiles struct {
	Requester core.Requester
	get       func(string, *core.CoreRequest, interface{}) error
	post      func(string, *core.CoreRequest, interface{}) error
	patch     func(string, *core.CoreRequest, interface{}) error
	put       func(string, *core.CoreRequest, interface{}) error
	delete    func(string, *core.CoreRequest, interface{}) error
}

func NewCardProfiles(requster core.Requester) (r *CardProfiles) {
	r = &CardProfiles{}
	r.Requester = requster
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

func (r *CardProfiles) Create(body CardProfilesCreateParams, opts ...*core.RequestOpts) (res CardProfile, err error) {
	err = r.post(
		"/card_profiles",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *CardProfiles) Retrieve(card_profile_id string, opts ...*core.RequestOpts) (res CardProfile, err error) {
	err = r.get(
		fmt.Sprintf("/card_profiles/%s", card_profile_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *CardProfiles) List(query *CardProfilesListQuery, opts ...*core.RequestOpts) (res CardProfilesListResponse, err error) {
	err = r.get(
		"/card_profiles",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Query:  query,
		},
		&res,
	)
	return
}

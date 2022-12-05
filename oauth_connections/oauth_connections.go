package oauth_connections

import "increase/core"
import "fmt"

type OauthConnections struct {
	Requester core.Requester
	get       func(string, *core.CoreRequest, interface{}) error
	post      func(string, *core.CoreRequest, interface{}) error
	patch     func(string, *core.CoreRequest, interface{}) error
	put       func(string, *core.CoreRequest, interface{}) error
	delete    func(string, *core.CoreRequest, interface{}) error
}

func NewOauthConnections(requster core.Requester) (r *OauthConnections) {
	r = &OauthConnections{}
	r.Requester = requster
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

func (r *OauthConnections) Retrieve(oauth_connection_id string, opts ...*core.RequestOpts) (res OauthConnection, err error) {
	err = r.get(
		fmt.Sprintf("/oauth_connections/%s", oauth_connection_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *OauthConnections) List(query *OauthConnectionsListQuery, opts ...*core.RequestOpts) (res OauthConnectionsListResponse, err error) {
	err = r.get(
		"/oauth_connections",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Query:  query,
		},
		&res,
	)
	return
}

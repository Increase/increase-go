package routing_numbers

import "increase/core"

type RoutingNumbers struct {
	Requester core.Requester
	get       func(string, *core.CoreRequest, interface{}) error
	post      func(string, *core.CoreRequest, interface{}) error
	patch     func(string, *core.CoreRequest, interface{}) error
	put       func(string, *core.CoreRequest, interface{}) error
	delete    func(string, *core.CoreRequest, interface{}) error
}

func NewRoutingNumbers(requster core.Requester) (r *RoutingNumbers) {
	r = &RoutingNumbers{}
	r.Requester = requster
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

// You can use this API to confirm if a routing number is valid, such as when a
// user is providing you with bank account details. Since routing numbers uniquely
// identify a bank, this will always return 0 or 1 entry. In Sandbox, the only
// valid routing number for this method is 110000000.
func (r *RoutingNumbers) List(query *RoutingNumbersListQuery, opts ...*core.RequestOpts) (res RoutingNumbersListResponse, err error) {
	err = r.get(
		"/routing_numbers",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Query:  query,
		},
		&res,
	)
	return
}

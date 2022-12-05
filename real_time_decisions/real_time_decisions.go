package real_time_decisions

import "increase/core"
import "fmt"

type RealTimeDecisions struct {
	Requester core.Requester
	get       func(string, *core.CoreRequest, interface{}) error
	post      func(string, *core.CoreRequest, interface{}) error
	patch     func(string, *core.CoreRequest, interface{}) error
	put       func(string, *core.CoreRequest, interface{}) error
	delete    func(string, *core.CoreRequest, interface{}) error
}

func NewRealTimeDecisions(requster core.Requester) (r *RealTimeDecisions) {
	r = &RealTimeDecisions{}
	r.Requester = requster
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

func (r *RealTimeDecisions) Retrieve(real_time_decision_id string, opts ...*core.RequestOpts) (res RealTimeDecision, err error) {
	err = r.get(
		fmt.Sprintf("/real_time_decisions/%s", real_time_decision_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *RealTimeDecisions) Action(real_time_decision_id string, body RealTimeDecisionsActionParams, opts ...*core.RequestOpts) (res RealTimeDecision, err error) {
	err = r.post(
		fmt.Sprintf("/real_time_decisions/%s/action", real_time_decision_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

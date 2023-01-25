package services

import (
	"context"
	"fmt"
	"increase/core"
	"increase/types"
)

type RealTimeDecisionService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewRealTimeDecisionService(requester core.Requester) (r *RealTimeDecisionService) {
	r = &RealTimeDecisionService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

// Retrieve a Real-Time Decision
func (r *RealTimeDecisionService) Retrieve(ctx context.Context, real_time_decision_id string, opts ...*core.RequestOpts) (res *types.RealTimeDecision, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/real_time_decisions/%s", real_time_decision_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)

	return
}

// Action a Real-Time Decision
func (r *RealTimeDecisionService) Action(ctx context.Context, real_time_decision_id string, body *types.ActionARealTimeDecisionParameters, opts ...*core.RequestOpts) (res *types.RealTimeDecision, err error) {
	err = r.post(
		ctx,
		fmt.Sprintf("/real_time_decisions/%s/action", real_time_decision_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)

	return
}

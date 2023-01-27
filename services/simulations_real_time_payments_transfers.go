package services

import (
	"context"
	"increase/core"
	"increase/types"
)

type SimulationsRealTimePaymentsTransferService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewSimulationsRealTimePaymentsTransferService(requester core.Requester) (r *SimulationsRealTimePaymentsTransferService) {
	r = &SimulationsRealTimePaymentsTransferService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

// Simulates an inbound Real Time Payments transfer to your account. Real Time
// Payments are a beta feature.
func (r *SimulationsRealTimePaymentsTransferService) NewInbound(ctx context.Context, body *types.SimulateARealTimePaymentsTransferToYourAccountParameters, opts ...*core.RequestOpts) (res *types.InboundRealTimePaymentsTransferSimulationResult, err error) {
	path := "/simulations/inbound_real_time_payments_transfers"
	req := &core.CoreRequest{
		Params: core.MergeRequestOpts(opts...),
		Body:   body,
	}
	err = r.post(ctx, path, req, &res)

	return
}

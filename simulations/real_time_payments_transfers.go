package simulations

import "context"
import "increase/core"

type RealTimePaymentsTransferService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewRealTimePaymentsTransferService(requester core.Requester) (r *RealTimePaymentsTransferService) {
	r = &RealTimePaymentsTransferService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

// Simulates an inbound Real Time Payments transfer to your account.
func (r *RealTimePaymentsTransferService) CreateInbound(ctx context.Context, body *SimulateARealTimePaymentsTransferToYourAccountParameters, opts ...*core.RequestOpts) (res *InboundRealTimePaymentsTransferSimulationResult, err error) {
	err = r.post(
		ctx,
		"/simulations/inbound_real_time_payments_transfers",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)

	return
}

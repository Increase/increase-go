package simulations

import "increase/core"

type RealTimePaymentsTransfers struct {
	Requester core.Requester
	get       func(string, *core.CoreRequest, interface{}) error
	post      func(string, *core.CoreRequest, interface{}) error
	patch     func(string, *core.CoreRequest, interface{}) error
	put       func(string, *core.CoreRequest, interface{}) error
	delete    func(string, *core.CoreRequest, interface{}) error
}

func NewRealTimePaymentsTransfers(requster core.Requester) (r *RealTimePaymentsTransfers) {
	r = &RealTimePaymentsTransfers{}
	r.Requester = requster
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

// Simulates an inbound Real Time Payments transfer to your account.
func (r *RealTimePaymentsTransfers) CreateInbound(body RealTimePaymentsTransfersCreateInboundParams, opts ...*core.RequestOpts) (res InboundRealTimePaymentsTransferSimulationResult, err error) {
	err = r.post(
		"/simulations/inbound_real_time_payments_transfers",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

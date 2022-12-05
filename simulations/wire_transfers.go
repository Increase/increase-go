package simulations

import "increase/core"

type WireTransfers struct {
	Requester core.Requester
	get       func(string, *core.CoreRequest, interface{}) error
	post      func(string, *core.CoreRequest, interface{}) error
	patch     func(string, *core.CoreRequest, interface{}) error
	put       func(string, *core.CoreRequest, interface{}) error
	delete    func(string, *core.CoreRequest, interface{}) error
}

func NewWireTransfers(requster core.Requester) (r *WireTransfers) {
	r = &WireTransfers{}
	r.Requester = requster
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

// Simulates an inbound Wire transfer to your account.
func (r *WireTransfers) CreateInbound(body WireTransfersCreateInboundParams, opts ...*core.RequestOpts) (res WireTransferSimulation, err error) {
	err = r.post(
		"/simulations/inbound_wire_transfers",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

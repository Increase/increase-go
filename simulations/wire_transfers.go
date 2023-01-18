package simulations

import "context"
import "increase/core"

type WireTransferService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewWireTransferService(requester core.Requester) (r *WireTransferService) {
	r = &WireTransferService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

// Simulates an inbound Wire transfer to your account.
func (r *WireTransferService) CreateInbound(ctx context.Context, body *SimulateAWireTransferToYourAccountParameters, opts ...*core.RequestOpts) (res *WireTransferSimulation, err error) {
	err = r.post(
		ctx,
		"/simulations/inbound_wire_transfers",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)

	return
}

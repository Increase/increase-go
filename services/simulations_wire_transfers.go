package services

import (
	"context"
	"increase/core"
	"increase/types"
)

type SimulationsWireTransferService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewSimulationsWireTransferService(requester core.Requester) (r *SimulationsWireTransferService) {
	r = &SimulationsWireTransferService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

// Simulates an inbound Wire Transfer to your account.
func (r *SimulationsWireTransferService) CreateInbound(ctx context.Context, body *types.SimulateAWireTransferToYourAccountParameters, opts ...*core.RequestOpts) (res *types.WireTransferSimulation, err error) {
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

package services

import (
	"context"
	"increase/core"
	"increase/types"
)

type SimulationsCardRefundService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewSimulationsCardRefundService(requester core.Requester) (r *SimulationsCardRefundService) {
	r = &SimulationsCardRefundService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

// Simulates refunding a card transaction. The full value of the original sandbox
// transaction is refunded.
func (r *SimulationsCardRefundService) Create(ctx context.Context, body *types.SimulateARefundOnACardParameters, opts ...*core.RequestOpts) (res *types.Transaction, err error) {
	err = r.post(
		ctx,
		"/simulations/card_refunds",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)

	return
}

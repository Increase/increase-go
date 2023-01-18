package services

import "context"
import "increase/core"
import "increase/types"

type SimulationsAccountStatementService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewSimulationsAccountStatementService(requester core.Requester) (r *SimulationsAccountStatementService) {
	r = &SimulationsAccountStatementService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

// Simulates an Account Statement being created for an account.
func (r *SimulationsAccountStatementService) Create(ctx context.Context, body *types.SimulateAnAccountStatementBeingCreatedParameters, opts ...*core.RequestOpts) (res *types.AccountStatement, err error) {
	err = r.post(
		ctx,
		"/simulations/account_statements",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)

	return
}

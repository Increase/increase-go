package simulations

import "context"
import "increase/core"
import "increase/account_statements"

type AccountStatementService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewAccountStatementService(requester core.Requester) (r *AccountStatementService) {
	r = &AccountStatementService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

// Simulates an Account Statement being created for an account.
func (r *AccountStatementService) Create(ctx context.Context, body *SimulateAnAccountStatementBeingCreatedParameters, opts ...*core.RequestOpts) (res *account_statements.AccountStatement, err error) {
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

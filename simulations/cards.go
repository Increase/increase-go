package simulations

import "increase/core"
import "increase/transactions"

type Cards struct {
	Requester core.Requester
	get       func(string, *core.CoreRequest, interface{}) error
	post      func(string, *core.CoreRequest, interface{}) error
	patch     func(string, *core.CoreRequest, interface{}) error
	put       func(string, *core.CoreRequest, interface{}) error
	delete    func(string, *core.CoreRequest, interface{}) error
}

func NewCards(requster core.Requester) (r *Cards) {
	r = &Cards{}
	r.Requester = requster
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

// Simulates activity on a Card. You can pass either a Card id or a Digital Wallet
// Token id to simulate the two different ways purchases can be made.
func (r *Cards) Authorize(body CardsAuthorizeParams, opts ...*core.RequestOpts) (res CardAuthorizationSimulation, err error) {
	err = r.post(
		"/simulations/card_authorizations",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

// Simulates the settlement of an authorization by a card acquirer.
func (r *Cards) Settlement(body CardsSettlementParams, opts ...*core.RequestOpts) (res transactions.Transaction, err error) {
	err = r.post(
		"/simulations/card_settlements",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

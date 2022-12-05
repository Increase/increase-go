package simulations

import "increase/core"

type Simulations struct {
	Requester                  core.Requester
	get                        func(string, *core.CoreRequest, interface{}) error
	post                       func(string, *core.CoreRequest, interface{}) error
	patch                      func(string, *core.CoreRequest, interface{}) error
	put                        func(string, *core.CoreRequest, interface{}) error
	delete                     func(string, *core.CoreRequest, interface{}) error
	AccountTransfers           *AccountTransfers
	AccountStatements          *AccountStatements
	ACHTransfers               *ACHTransfers
	CheckTransfers             *CheckTransfers
	DigitalWalletTokenRequests *DigitalWalletTokenRequests
	CheckDeposits              *CheckDeposits
	WireTransfers              *WireTransfers
	Cards                      *Cards
	RealTimePaymentsTransfers  *RealTimePaymentsTransfers
}

func NewSimulations(requster core.Requester) (r *Simulations) {
	r = &Simulations{}
	r.Requester = requster
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	r.AccountTransfers = NewAccountTransfers(requster)
	r.AccountStatements = NewAccountStatements(requster)
	r.ACHTransfers = NewACHTransfers(requster)
	r.CheckTransfers = NewCheckTransfers(requster)
	r.DigitalWalletTokenRequests = NewDigitalWalletTokenRequests(requster)
	r.CheckDeposits = NewCheckDeposits(requster)
	r.WireTransfers = NewWireTransfers(requster)
	r.Cards = NewCards(requster)
	r.RealTimePaymentsTransfers = NewRealTimePaymentsTransfers(requster)
	return
}

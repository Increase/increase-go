package services

import (
	"context"
	"increase/core"
)

type SimulationService struct {
	Requester                  core.Requester
	get                        func(context.Context, string, *core.CoreRequest, interface{}) error
	post                       func(context.Context, string, *core.CoreRequest, interface{}) error
	patch                      func(context.Context, string, *core.CoreRequest, interface{}) error
	put                        func(context.Context, string, *core.CoreRequest, interface{}) error
	delete                     func(context.Context, string, *core.CoreRequest, interface{}) error
	AccountTransfers           *SimulationsAccountTransferService
	AccountStatements          *SimulationsAccountStatementService
	ACHTransfers               *SimulationsACHTransferService
	CardDisputes               *SimulationsCardDisputeService
	CardRefunds                *SimulationsCardRefundService
	CheckTransfers             *SimulationsCheckTransferService
	Documents                  *SimulationsDocumentService
	DigitalWalletTokenRequests *SimulationsDigitalWalletTokenRequestService
	CheckDeposits              *SimulationsCheckDepositService
	WireTransfers              *SimulationsWireTransferService
	Cards                      *SimulationsCardService
	RealTimePaymentsTransfers  *SimulationsRealTimePaymentsTransferService
}

func NewSimulationService(requester core.Requester) (r *SimulationService) {
	r = &SimulationService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	r.AccountTransfers = NewSimulationsAccountTransferService(requester)
	r.AccountStatements = NewSimulationsAccountStatementService(requester)
	r.ACHTransfers = NewSimulationsACHTransferService(requester)
	r.CardDisputes = NewSimulationsCardDisputeService(requester)
	r.CardRefunds = NewSimulationsCardRefundService(requester)
	r.CheckTransfers = NewSimulationsCheckTransferService(requester)
	r.Documents = NewSimulationsDocumentService(requester)
	r.DigitalWalletTokenRequests = NewSimulationsDigitalWalletTokenRequestService(requester)
	r.CheckDeposits = NewSimulationsCheckDepositService(requester)
	r.WireTransfers = NewSimulationsWireTransferService(requester)
	r.Cards = NewSimulationsCardService(requester)
	r.RealTimePaymentsTransfers = NewSimulationsRealTimePaymentsTransferService(requester)
	return
}

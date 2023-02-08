package services

import "increase/options"

type SimulationService struct {
	Options                    []options.RequestOption
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

func NewSimulationService(opts ...options.RequestOption) (r *SimulationService) {
	r = &SimulationService{}
	r.Options = opts
	r.AccountTransfers = NewSimulationsAccountTransferService(opts...)
	r.AccountStatements = NewSimulationsAccountStatementService(opts...)
	r.ACHTransfers = NewSimulationsACHTransferService(opts...)
	r.CardDisputes = NewSimulationsCardDisputeService(opts...)
	r.CardRefunds = NewSimulationsCardRefundService(opts...)
	r.CheckTransfers = NewSimulationsCheckTransferService(opts...)
	r.Documents = NewSimulationsDocumentService(opts...)
	r.DigitalWalletTokenRequests = NewSimulationsDigitalWalletTokenRequestService(opts...)
	r.CheckDeposits = NewSimulationsCheckDepositService(opts...)
	r.WireTransfers = NewSimulationsWireTransferService(opts...)
	r.Cards = NewSimulationsCardService(opts...)
	r.RealTimePaymentsTransfers = NewSimulationsRealTimePaymentsTransferService(opts...)
	return
}

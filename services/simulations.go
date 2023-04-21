package services

import (
	"github.com/increase/increase-go/option"
)

type SimulationService struct {
	Options                     []option.RequestOption
	AccountTransfers            *SimulationAccountTransferService
	AccountStatements           *SimulationAccountStatementService
	ACHTransfers                *SimulationACHTransferService
	CardDisputes                *SimulationCardDisputeService
	CardRefunds                 *SimulationCardRefundService
	CheckTransfers              *SimulationCheckTransferService
	Documents                   *SimulationDocumentService
	DigitalWalletTokenRequests  *SimulationDigitalWalletTokenRequestService
	CheckDeposits               *SimulationCheckDepositService
	InboundWireDrawdownRequests *SimulationInboundWireDrawdownRequestService
	InterestPayments            *SimulationInterestPaymentService
	WireTransfers               *SimulationWireTransferService
	Cards                       *SimulationCardService
	RealTimePaymentsTransfers   *SimulationRealTimePaymentsTransferService
}

func NewSimulationService(opts ...option.RequestOption) (r *SimulationService) {
	r = &SimulationService{}
	r.Options = opts
	r.AccountTransfers = NewSimulationAccountTransferService(opts...)
	r.AccountStatements = NewSimulationAccountStatementService(opts...)
	r.ACHTransfers = NewSimulationACHTransferService(opts...)
	r.CardDisputes = NewSimulationCardDisputeService(opts...)
	r.CardRefunds = NewSimulationCardRefundService(opts...)
	r.CheckTransfers = NewSimulationCheckTransferService(opts...)
	r.Documents = NewSimulationDocumentService(opts...)
	r.DigitalWalletTokenRequests = NewSimulationDigitalWalletTokenRequestService(opts...)
	r.CheckDeposits = NewSimulationCheckDepositService(opts...)
	r.InboundWireDrawdownRequests = NewSimulationInboundWireDrawdownRequestService(opts...)
	r.InterestPayments = NewSimulationInterestPaymentService(opts...)
	r.WireTransfers = NewSimulationWireTransferService(opts...)
	r.Cards = NewSimulationCardService(opts...)
	r.RealTimePaymentsTransfers = NewSimulationRealTimePaymentsTransferService(opts...)
	return
}

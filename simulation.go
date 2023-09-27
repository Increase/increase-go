// File generated from our OpenAPI spec by Stainless.

package increase

import (
	"github.com/increase/increase-go/option"
)

// SimulationService contains methods and other services that help with interacting
// with the increase API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewSimulationService] method instead.
type SimulationService struct {
	Options                     []option.RequestOption
	AccountTransfers            *SimulationAccountTransferService
	AccountStatements           *SimulationAccountStatementService
	ACHTransfers                *SimulationACHTransferService
	CardDisputes                *SimulationCardDisputeService
	CardProfiles                *SimulationCardProfileService
	CardRefunds                 *SimulationCardRefundService
	CheckTransfers              *SimulationCheckTransferService
	Documents                   *SimulationDocumentService
	DigitalWalletTokenRequests  *SimulationDigitalWalletTokenRequestService
	CheckDeposits               *SimulationCheckDepositService
	Programs                    *SimulationProgramService
	InboundWireDrawdownRequests *SimulationInboundWireDrawdownRequestService
	InboundFundsHolds           *SimulationInboundFundsHoldService
	InterestPayments            *SimulationInterestPaymentService
	WireTransfers               *SimulationWireTransferService
	Cards                       *SimulationCardService
	RealTimePaymentsTransfers   *SimulationRealTimePaymentsTransferService
	PhysicalCards               *SimulationPhysicalCardService
}

// NewSimulationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSimulationService(opts ...option.RequestOption) (r *SimulationService) {
	r = &SimulationService{}
	r.Options = opts
	r.AccountTransfers = NewSimulationAccountTransferService(opts...)
	r.AccountStatements = NewSimulationAccountStatementService(opts...)
	r.ACHTransfers = NewSimulationACHTransferService(opts...)
	r.CardDisputes = NewSimulationCardDisputeService(opts...)
	r.CardProfiles = NewSimulationCardProfileService(opts...)
	r.CardRefunds = NewSimulationCardRefundService(opts...)
	r.CheckTransfers = NewSimulationCheckTransferService(opts...)
	r.Documents = NewSimulationDocumentService(opts...)
	r.DigitalWalletTokenRequests = NewSimulationDigitalWalletTokenRequestService(opts...)
	r.CheckDeposits = NewSimulationCheckDepositService(opts...)
	r.Programs = NewSimulationProgramService(opts...)
	r.InboundWireDrawdownRequests = NewSimulationInboundWireDrawdownRequestService(opts...)
	r.InboundFundsHolds = NewSimulationInboundFundsHoldService(opts...)
	r.InterestPayments = NewSimulationInterestPaymentService(opts...)
	r.WireTransfers = NewSimulationWireTransferService(opts...)
	r.Cards = NewSimulationCardService(opts...)
	r.RealTimePaymentsTransfers = NewSimulationRealTimePaymentsTransferService(opts...)
	r.PhysicalCards = NewSimulationPhysicalCardService(opts...)
	return
}

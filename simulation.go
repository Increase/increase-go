// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"github.com/increase/increase-go/option"
)

// SimulationService contains methods and other services that help with interacting
// with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimulationService] method instead.
type SimulationService struct {
	Options                          []option.RequestOption
	AccountTransfers                 *SimulationAccountTransferService
	InboundACHTransfers              *SimulationInboundACHTransferService
	ACHTransfers                     *SimulationACHTransferService
	CheckTransfers                   *SimulationCheckTransferService
	InboundCheckDeposits             *SimulationInboundCheckDepositService
	CheckDeposits                    *SimulationCheckDepositService
	InboundWireTransfers             *SimulationInboundWireTransferService
	WireTransfers                    *SimulationWireTransferService
	InboundWireDrawdownRequests      *SimulationInboundWireDrawdownRequestService
	InboundRealTimePaymentsTransfers *SimulationInboundRealTimePaymentsTransferService
	InboundFundsHolds                *SimulationInboundFundsHoldService
	RealTimePaymentsTransfers        *SimulationRealTimePaymentsTransferService
	InboundInternationalACHTransfers *SimulationInboundInternationalACHTransferService
	CardAuthorizations               *SimulationCardAuthorizationService
	CardSettlements                  *SimulationCardSettlementService
	CardReversals                    *SimulationCardReversalService
	CardIncrements                   *SimulationCardIncrementService
	CardAuthorizationExpirations     *SimulationCardAuthorizationExpirationService
	CardFuelConfirmations            *SimulationCardFuelConfirmationService
	CardRefunds                      *SimulationCardRefundService
	CardDisputes                     *SimulationCardDisputeService
	DigitalWalletTokenRequests       *SimulationDigitalWalletTokenRequestService
	PhysicalCards                    *SimulationPhysicalCardService
	InterestPayments                 *SimulationInterestPaymentService
	AccountStatements                *SimulationAccountStatementService
	Documents                        *SimulationDocumentService
	Programs                         *SimulationProgramService
}

// NewSimulationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSimulationService(opts ...option.RequestOption) (r *SimulationService) {
	r = &SimulationService{}
	r.Options = opts
	r.AccountTransfers = NewSimulationAccountTransferService(opts...)
	r.InboundACHTransfers = NewSimulationInboundACHTransferService(opts...)
	r.ACHTransfers = NewSimulationACHTransferService(opts...)
	r.CheckTransfers = NewSimulationCheckTransferService(opts...)
	r.InboundCheckDeposits = NewSimulationInboundCheckDepositService(opts...)
	r.CheckDeposits = NewSimulationCheckDepositService(opts...)
	r.InboundWireTransfers = NewSimulationInboundWireTransferService(opts...)
	r.WireTransfers = NewSimulationWireTransferService(opts...)
	r.InboundWireDrawdownRequests = NewSimulationInboundWireDrawdownRequestService(opts...)
	r.InboundRealTimePaymentsTransfers = NewSimulationInboundRealTimePaymentsTransferService(opts...)
	r.InboundFundsHolds = NewSimulationInboundFundsHoldService(opts...)
	r.RealTimePaymentsTransfers = NewSimulationRealTimePaymentsTransferService(opts...)
	r.InboundInternationalACHTransfers = NewSimulationInboundInternationalACHTransferService(opts...)
	r.CardAuthorizations = NewSimulationCardAuthorizationService(opts...)
	r.CardSettlements = NewSimulationCardSettlementService(opts...)
	r.CardReversals = NewSimulationCardReversalService(opts...)
	r.CardIncrements = NewSimulationCardIncrementService(opts...)
	r.CardAuthorizationExpirations = NewSimulationCardAuthorizationExpirationService(opts...)
	r.CardFuelConfirmations = NewSimulationCardFuelConfirmationService(opts...)
	r.CardRefunds = NewSimulationCardRefundService(opts...)
	r.CardDisputes = NewSimulationCardDisputeService(opts...)
	r.DigitalWalletTokenRequests = NewSimulationDigitalWalletTokenRequestService(opts...)
	r.PhysicalCards = NewSimulationPhysicalCardService(opts...)
	r.InterestPayments = NewSimulationInterestPaymentService(opts...)
	r.AccountStatements = NewSimulationAccountStatementService(opts...)
	r.Documents = NewSimulationDocumentService(opts...)
	r.Programs = NewSimulationProgramService(opts...)
	return
}

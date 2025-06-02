// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"github.com/Increase/increase-go/option"
)

// SimulationService contains methods and other services that help with interacting
// with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimulationService] method instead.
type SimulationService struct {
	Options                          []option.RequestOption
	InterestPayments                 *SimulationInterestPaymentService
	FeePayments                      *SimulationFeePaymentService
	CardAuthorizations               *SimulationCardAuthorizationService
	CardAuthorizationExpirations     *SimulationCardAuthorizationExpirationService
	CardSettlements                  *SimulationCardSettlementService
	CardReversals                    *SimulationCardReversalService
	CardIncrements                   *SimulationCardIncrementService
	CardFuelConfirmations            *SimulationCardFuelConfirmationService
	CardRefunds                      *SimulationCardRefundService
	CardDisputes                     *SimulationCardDisputeService
	PhysicalCards                    *SimulationPhysicalCardService
	DigitalWalletTokenRequests       *SimulationDigitalWalletTokenRequestService
	InboundFundsHolds                *SimulationInboundFundsHoldService
	AccountTransfers                 *SimulationAccountTransferService
	ACHTransfers                     *SimulationACHTransferService
	InboundACHTransfers              *SimulationInboundACHTransferService
	WireTransfers                    *SimulationWireTransferService
	InboundWireTransfers             *SimulationInboundWireTransferService
	InboundWireDrawdownRequests      *SimulationInboundWireDrawdownRequestService
	CheckTransfers                   *SimulationCheckTransferService
	InboundCheckDeposits             *SimulationInboundCheckDepositService
	RealTimePaymentsTransfers        *SimulationRealTimePaymentsTransferService
	InboundRealTimePaymentsTransfers *SimulationInboundRealTimePaymentsTransferService
	CheckDeposits                    *SimulationCheckDepositService
	InboundMailItems                 *SimulationInboundMailItemService
	Programs                         *SimulationProgramService
	AccountStatements                *SimulationAccountStatementService
	Documents                        *SimulationDocumentService
}

// NewSimulationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSimulationService(opts ...option.RequestOption) (r *SimulationService) {
	r = &SimulationService{}
	r.Options = opts
	r.InterestPayments = NewSimulationInterestPaymentService(opts...)
	r.FeePayments = NewSimulationFeePaymentService(opts...)
	r.CardAuthorizations = NewSimulationCardAuthorizationService(opts...)
	r.CardAuthorizationExpirations = NewSimulationCardAuthorizationExpirationService(opts...)
	r.CardSettlements = NewSimulationCardSettlementService(opts...)
	r.CardReversals = NewSimulationCardReversalService(opts...)
	r.CardIncrements = NewSimulationCardIncrementService(opts...)
	r.CardFuelConfirmations = NewSimulationCardFuelConfirmationService(opts...)
	r.CardRefunds = NewSimulationCardRefundService(opts...)
	r.CardDisputes = NewSimulationCardDisputeService(opts...)
	r.PhysicalCards = NewSimulationPhysicalCardService(opts...)
	r.DigitalWalletTokenRequests = NewSimulationDigitalWalletTokenRequestService(opts...)
	r.InboundFundsHolds = NewSimulationInboundFundsHoldService(opts...)
	r.AccountTransfers = NewSimulationAccountTransferService(opts...)
	r.ACHTransfers = NewSimulationACHTransferService(opts...)
	r.InboundACHTransfers = NewSimulationInboundACHTransferService(opts...)
	r.WireTransfers = NewSimulationWireTransferService(opts...)
	r.InboundWireTransfers = NewSimulationInboundWireTransferService(opts...)
	r.InboundWireDrawdownRequests = NewSimulationInboundWireDrawdownRequestService(opts...)
	r.CheckTransfers = NewSimulationCheckTransferService(opts...)
	r.InboundCheckDeposits = NewSimulationInboundCheckDepositService(opts...)
	r.RealTimePaymentsTransfers = NewSimulationRealTimePaymentsTransferService(opts...)
	r.InboundRealTimePaymentsTransfers = NewSimulationInboundRealTimePaymentsTransferService(opts...)
	r.CheckDeposits = NewSimulationCheckDepositService(opts...)
	r.InboundMailItems = NewSimulationInboundMailItemService(opts...)
	r.Programs = NewSimulationProgramService(opts...)
	r.AccountStatements = NewSimulationAccountStatementService(opts...)
	r.Documents = NewSimulationDocumentService(opts...)
	return
}

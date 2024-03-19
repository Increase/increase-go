// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"net/http"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
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

// Simulates expiring a card authorization immediately.
func (r *SimulationService) CardAuthorizationExpirations(ctx context.Context, body SimulationCardAuthorizationExpirationsParams, opts ...option.RequestOption) (res *CardPayment, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/card_authorization_expirations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Simulates the fuel confirmation of an authorization by a card acquirer. This
// happens asynchronously right after a fuel pump transaction is completed. A fuel
// confirmation can only happen once per authorization.
func (r *SimulationService) CardFuelConfirmations(ctx context.Context, body SimulationCardFuelConfirmationsParams, opts ...option.RequestOption) (res *CardPayment, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/card_fuel_confirmations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Simulates the increment of an authorization by a card acquirer. An authorization
// can be incremented multiple times.
func (r *SimulationService) CardIncrements(ctx context.Context, body SimulationCardIncrementsParams, opts ...option.RequestOption) (res *CardPayment, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/card_increments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Simulates the reversal of an authorization by a card acquirer. An authorization
// can be partially reversed multiple times, up until the total authorized amount.
// Marks the pending transaction as complete if the authorization is fully
// reversed.
func (r *SimulationService) CardReversals(ctx context.Context, body SimulationCardReversalsParams, opts ...option.RequestOption) (res *CardPayment, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/card_reversals"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SimulationCardAuthorizationExpirationsParams struct {
	// The identifier of the Card Payment to expire.
	CardPaymentID param.Field[string] `json:"card_payment_id,required"`
}

func (r SimulationCardAuthorizationExpirationsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SimulationCardFuelConfirmationsParams struct {
	// The amount of the fuel_confirmation in minor units in the card authorization's
	// currency.
	Amount param.Field[int64] `json:"amount,required"`
	// The identifier of the Card Payment to create a fuel_confirmation on.
	CardPaymentID param.Field[string] `json:"card_payment_id,required"`
}

func (r SimulationCardFuelConfirmationsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SimulationCardIncrementsParams struct {
	// The amount of the increment in minor units in the card authorization's currency.
	Amount param.Field[int64] `json:"amount,required"`
	// The identifier of the Card Payment to create a increment on.
	CardPaymentID param.Field[string] `json:"card_payment_id,required"`
	// The identifier of the Event Subscription to use. If provided, will override the
	// default real time event subscription. Because you can only create one real time
	// decision event subscription, you can use this field to route events to any
	// specified event subscription for testing purposes.
	EventSubscriptionID param.Field[string] `json:"event_subscription_id"`
}

func (r SimulationCardIncrementsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SimulationCardReversalsParams struct {
	// The identifier of the Card Payment to create a reversal on.
	CardPaymentID param.Field[string] `json:"card_payment_id,required"`
	// The amount of the reversal in minor units in the card authorization's currency.
	// This defaults to the authorization amount.
	Amount param.Field[int64] `json:"amount"`
}

func (r SimulationCardReversalsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

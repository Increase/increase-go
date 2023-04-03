package increase

import (
	"os"

	"github.com/increase/increase-go/core/fields"
	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/services"
)

func F[T any](value T) fields.Field[T]          { return fields.Field[T]{Value: value, Present: true} }
func NullField[T any]() fields.Field[T]         { return fields.Field[T]{Null: true, Present: true} }
func RawField[T any](value any) fields.Field[T] { return fields.Field[T]{Raw: value, Present: true} }

func Float[T float32 | float64](value T) fields.Field[float64] {
	return fields.Field[float64]{Value: float64(value), Present: true}
}
func Int[T int | int8 | int16 | int32 | int64](value T) fields.Field[int64] {
	return fields.Field[int64]{Value: int64(value), Present: true}
}
func UInt[T uint | uint8 | uint16 | uint32 | uint64](value T) fields.Field[uint64] {
	return fields.Field[uint64]{Value: uint64(value), Present: true}
}
func Str(str string) fields.Field[string] { return F(str) }

type Increase struct {
	Options                     []options.RequestOption
	Accounts                    *services.AccountService
	AccountNumbers              *services.AccountNumberService
	RealTimeDecisions           *services.RealTimeDecisionService
	Cards                       *services.CardService
	CardDisputes                *services.CardDisputeService
	CardProfiles                *services.CardProfileService
	ExternalAccounts            *services.ExternalAccountService
	DigitalWalletTokens         *services.DigitalWalletTokenService
	Transactions                *services.TransactionService
	PendingTransactions         *services.PendingTransactionService
	DeclinedTransactions        *services.DeclinedTransactionService
	Limits                      *services.LimitService
	AccountTransfers            *services.AccountTransferService
	ACHTransfers                *services.ACHTransferService
	InboundACHTransferReturns   *services.InboundACHTransferReturnService
	ACHPrenotifications         *services.ACHPrenotificationService
	Documents                   *services.DocumentService
	WireTransfers               *services.WireTransferService
	CheckTransfers              *services.CheckTransferService
	Entities                    *services.EntityService
	InboundWireDrawdownRequests *services.InboundWireDrawdownRequestService
	WireDrawdownRequests        *services.WireDrawdownRequestService
	Events                      *services.EventService
	EventSubscriptions          *services.EventSubscriptionService
	Files                       *services.FileService
	Groups                      *services.GroupService
	OauthConnections            *services.OauthConnectionService
	CheckDeposits               *services.CheckDepositService
	RoutingNumbers              *services.RoutingNumberService
	AccountStatements           *services.AccountStatementService
	Simulations                 *services.SimulationService
}

// NewIncrease generates a new client with the default options read from the
// environment ("INCREASE_API_KEY"). The options passed in as arguments are applied
// after these default arguments, and all options will be passed down to the
// services and requests that this client makes.
func NewIncrease(opts ...options.RequestOption) (r *Increase) {
	defaults := []options.RequestOption{options.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("INCREASE_API_KEY"); ok {
		defaults = append(defaults, options.WithAPIKey(o))
	}
	opts = append(defaults, opts...)

	r = &Increase{Options: opts}

	r.Accounts = services.NewAccountService(opts...)
	r.AccountNumbers = services.NewAccountNumberService(opts...)
	r.RealTimeDecisions = services.NewRealTimeDecisionService(opts...)
	r.Cards = services.NewCardService(opts...)
	r.CardDisputes = services.NewCardDisputeService(opts...)
	r.CardProfiles = services.NewCardProfileService(opts...)
	r.ExternalAccounts = services.NewExternalAccountService(opts...)
	r.DigitalWalletTokens = services.NewDigitalWalletTokenService(opts...)
	r.Transactions = services.NewTransactionService(opts...)
	r.PendingTransactions = services.NewPendingTransactionService(opts...)
	r.DeclinedTransactions = services.NewDeclinedTransactionService(opts...)
	r.Limits = services.NewLimitService(opts...)
	r.AccountTransfers = services.NewAccountTransferService(opts...)
	r.ACHTransfers = services.NewACHTransferService(opts...)
	r.InboundACHTransferReturns = services.NewInboundACHTransferReturnService(opts...)
	r.ACHPrenotifications = services.NewACHPrenotificationService(opts...)
	r.Documents = services.NewDocumentService(opts...)
	r.WireTransfers = services.NewWireTransferService(opts...)
	r.CheckTransfers = services.NewCheckTransferService(opts...)
	r.Entities = services.NewEntityService(opts...)
	r.InboundWireDrawdownRequests = services.NewInboundWireDrawdownRequestService(opts...)
	r.WireDrawdownRequests = services.NewWireDrawdownRequestService(opts...)
	r.Events = services.NewEventService(opts...)
	r.EventSubscriptions = services.NewEventSubscriptionService(opts...)
	r.Files = services.NewFileService(opts...)
	r.Groups = services.NewGroupService(opts...)
	r.OauthConnections = services.NewOauthConnectionService(opts...)
	r.CheckDeposits = services.NewCheckDepositService(opts...)
	r.RoutingNumbers = services.NewRoutingNumberService(opts...)
	r.AccountStatements = services.NewAccountStatementService(opts...)
	r.Simulations = services.NewSimulationService(opts...)

	return
}

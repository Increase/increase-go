package increase

import (
	"os"

	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/services"
)

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

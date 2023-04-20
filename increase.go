package increase

import (
	"os"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/services"
)

type Increase struct {
	Options                     []option.RequestOption
	Accounts                    *services.AccountService
	AccountNumbers              *services.AccountNumberService
	BookkeepingAccounts         *services.BookkeepingAccountService
	BookkeepingEntrySets        *services.BookkeepingEntrySetService
	BookkeepingEntries          *services.BookkeepingEntryService
	RealTimeDecisions           *services.RealTimeDecisionService
	RealTimePaymentsTransfers   *services.RealTimePaymentsTransferService
	BalanceLookups              *services.BalanceLookupService
	Cards                       *services.CardService
	CardDisputes                *services.CardDisputeService
	CardProfiles                *services.CardProfileService
	ExternalAccounts            *services.ExternalAccountService
	Exports                     *services.ExportService
	DigitalWalletTokens         *services.DigitalWalletTokenService
	Transactions                *services.TransactionService
	PendingTransactions         *services.PendingTransactionService
	Programs                    *services.ProgramService
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

// NewIncrease generates a new client with the default option read from the
// environment ("INCREASE_API_KEY"). The option passed in as arguments are applied
// after these default arguments, and all option will be passed down to the
// services and requests that this client makes.
func NewIncrease(opts ...option.RequestOption) (r *Increase) {
	defaults := []option.RequestOption{option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("INCREASE_API_KEY"); ok {
		defaults = append(defaults, option.WithAPIKey(o))
	}
	opts = append(defaults, opts...)

	r = &Increase{Options: opts}

	r.Accounts = services.NewAccountService(opts...)
	r.AccountNumbers = services.NewAccountNumberService(opts...)
	r.BookkeepingAccounts = services.NewBookkeepingAccountService(opts...)
	r.BookkeepingEntrySets = services.NewBookkeepingEntrySetService(opts...)
	r.BookkeepingEntries = services.NewBookkeepingEntryService(opts...)
	r.RealTimeDecisions = services.NewRealTimeDecisionService(opts...)
	r.RealTimePaymentsTransfers = services.NewRealTimePaymentsTransferService(opts...)
	r.BalanceLookups = services.NewBalanceLookupService(opts...)
	r.Cards = services.NewCardService(opts...)
	r.CardDisputes = services.NewCardDisputeService(opts...)
	r.CardProfiles = services.NewCardProfileService(opts...)
	r.ExternalAccounts = services.NewExternalAccountService(opts...)
	r.Exports = services.NewExportService(opts...)
	r.DigitalWalletTokens = services.NewDigitalWalletTokenService(opts...)
	r.Transactions = services.NewTransactionService(opts...)
	r.PendingTransactions = services.NewPendingTransactionService(opts...)
	r.Programs = services.NewProgramService(opts...)
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

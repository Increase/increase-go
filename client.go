package increase

import (
	"os"

	"github.com/increase/increase-go/option"
)

type Client struct {
	Options                     []option.RequestOption
	Accounts                    *AccountService
	AccountNumbers              *AccountNumberService
	BookkeepingAccounts         *BookkeepingAccountService
	BookkeepingEntrySets        *BookkeepingEntrySetService
	BookkeepingEntries          *BookkeepingEntryService
	RealTimeDecisions           *RealTimeDecisionService
	RealTimePaymentsTransfers   *RealTimePaymentsTransferService
	BalanceLookups              *BalanceLookupService
	Cards                       *CardService
	CardDisputes                *CardDisputeService
	CardProfiles                *CardProfileService
	ExternalAccounts            *ExternalAccountService
	Exports                     *ExportService
	DigitalWalletTokens         *DigitalWalletTokenService
	Transactions                *TransactionService
	PendingTransactions         *PendingTransactionService
	Programs                    *ProgramService
	DeclinedTransactions        *DeclinedTransactionService
	Limits                      *LimitService
	AccountTransfers            *AccountTransferService
	ACHTransfers                *ACHTransferService
	InboundACHTransferReturns   *InboundACHTransferReturnService
	ACHPrenotifications         *ACHPrenotificationService
	Documents                   *DocumentService
	WireTransfers               *WireTransferService
	CheckTransfers              *CheckTransferService
	Entities                    *EntityService
	InboundWireDrawdownRequests *InboundWireDrawdownRequestService
	WireDrawdownRequests        *WireDrawdownRequestService
	Events                      *EventService
	EventSubscriptions          *EventSubscriptionService
	Files                       *FileService
	Groups                      *GroupService
	OauthConnections            *OauthConnectionService
	CheckDeposits               *CheckDepositService
	RoutingNumbers              *RoutingNumberService
	AccountStatements           *AccountStatementService
	Simulations                 *SimulationService
}

// NewClient generates a new client with the default option read from the
// environment ("INCREASE_API_KEY"). The option passed in as arguments are applied
// after these default arguments, and all option will be passed down to the
// services and requests that this client makes.
func NewClient(opts ...option.RequestOption) (r *Client) {
	defaults := []option.RequestOption{option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("INCREASE_API_KEY"); ok {
		defaults = append(defaults, option.WithAPIKey(o))
	}
	opts = append(defaults, opts...)

	r = &Client{Options: opts}

	r.Accounts = NewAccountService(opts...)
	r.AccountNumbers = NewAccountNumberService(opts...)
	r.BookkeepingAccounts = NewBookkeepingAccountService(opts...)
	r.BookkeepingEntrySets = NewBookkeepingEntrySetService(opts...)
	r.BookkeepingEntries = NewBookkeepingEntryService(opts...)
	r.RealTimeDecisions = NewRealTimeDecisionService(opts...)
	r.RealTimePaymentsTransfers = NewRealTimePaymentsTransferService(opts...)
	r.BalanceLookups = NewBalanceLookupService(opts...)
	r.Cards = NewCardService(opts...)
	r.CardDisputes = NewCardDisputeService(opts...)
	r.CardProfiles = NewCardProfileService(opts...)
	r.ExternalAccounts = NewExternalAccountService(opts...)
	r.Exports = NewExportService(opts...)
	r.DigitalWalletTokens = NewDigitalWalletTokenService(opts...)
	r.Transactions = NewTransactionService(opts...)
	r.PendingTransactions = NewPendingTransactionService(opts...)
	r.Programs = NewProgramService(opts...)
	r.DeclinedTransactions = NewDeclinedTransactionService(opts...)
	r.Limits = NewLimitService(opts...)
	r.AccountTransfers = NewAccountTransferService(opts...)
	r.ACHTransfers = NewACHTransferService(opts...)
	r.InboundACHTransferReturns = NewInboundACHTransferReturnService(opts...)
	r.ACHPrenotifications = NewACHPrenotificationService(opts...)
	r.Documents = NewDocumentService(opts...)
	r.WireTransfers = NewWireTransferService(opts...)
	r.CheckTransfers = NewCheckTransferService(opts...)
	r.Entities = NewEntityService(opts...)
	r.InboundWireDrawdownRequests = NewInboundWireDrawdownRequestService(opts...)
	r.WireDrawdownRequests = NewWireDrawdownRequestService(opts...)
	r.Events = NewEventService(opts...)
	r.EventSubscriptions = NewEventSubscriptionService(opts...)
	r.Files = NewFileService(opts...)
	r.Groups = NewGroupService(opts...)
	r.OauthConnections = NewOauthConnectionService(opts...)
	r.CheckDeposits = NewCheckDepositService(opts...)
	r.RoutingNumbers = NewRoutingNumberService(opts...)
	r.AccountStatements = NewAccountStatementService(opts...)
	r.Simulations = NewSimulationService(opts...)

	return
}

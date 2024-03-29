// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"os"

	"github.com/increase/increase-go/option"
)

// Client creates a struct with services and top level methods that help with
// interacting with the increase API. You should not instantiate this client
// directly, and instead use the [NewClient] method instead.
type Client struct {
	Options                                []option.RequestOption
	Accounts                               *AccountService
	AccountNumbers                         *AccountNumberService
	BookkeepingAccounts                    *BookkeepingAccountService
	BookkeepingEntrySets                   *BookkeepingEntrySetService
	BookkeepingEntries                     *BookkeepingEntryService
	RealTimeDecisions                      *RealTimeDecisionService
	RealTimePaymentsTransfers              *RealTimePaymentsTransferService
	Cards                                  *CardService
	CardDisputes                           *CardDisputeService
	CardPurchaseSupplements                *CardPurchaseSupplementService
	ExternalAccounts                       *ExternalAccountService
	Exports                                *ExportService
	DigitalWalletTokens                    *DigitalWalletTokenService
	Transactions                           *TransactionService
	PendingTransactions                    *PendingTransactionService
	Programs                               *ProgramService
	DeclinedTransactions                   *DeclinedTransactionService
	AccountTransfers                       *AccountTransferService
	ACHTransfers                           *ACHTransferService
	ACHPrenotifications                    *ACHPrenotificationService
	Documents                              *DocumentService
	WireTransfers                          *WireTransferService
	CheckTransfers                         *CheckTransferService
	Entities                               *EntityService
	InboundACHTransfers                    *InboundACHTransferService
	InboundWireDrawdownRequests            *InboundWireDrawdownRequestService
	WireDrawdownRequests                   *WireDrawdownRequestService
	Events                                 *EventService
	EventSubscriptions                     *EventSubscriptionService
	Files                                  *FileService
	Groups                                 *GroupService
	OAuthConnections                       *OAuthConnectionService
	CheckDeposits                          *CheckDepositService
	RoutingNumbers                         *RoutingNumberService
	AccountStatements                      *AccountStatementService
	Simulations                            *SimulationService
	PhysicalCards                          *PhysicalCardService
	CardPayments                           *CardPaymentService
	ProofOfAuthorizationRequests           *ProofOfAuthorizationRequestService
	ProofOfAuthorizationRequestSubmissions *ProofOfAuthorizationRequestSubmissionService
	Intrafi                                *IntrafiService
	RealTimePaymentsRequestForPayments     *RealTimePaymentsRequestForPaymentService
	Webhooks                               *WebhookService
	OAuthTokens                            *OAuthTokenService
	InboundWireTransfers                   *InboundWireTransferService
	DigitalCardProfiles                    *DigitalCardProfileService
	PhysicalCardProfiles                   *PhysicalCardProfileService
}

// NewClient generates a new client with the default option read from the
// environment (INCREASE_API_KEY, INCREASE_WEBHOOK_SECRET). The option passed in as
// arguments are applied after these default arguments, and all option will be
// passed down to the services and requests that this client makes.
func NewClient(opts ...option.RequestOption) (r *Client) {
	defaults := []option.RequestOption{option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("INCREASE_API_KEY"); ok {
		defaults = append(defaults, option.WithAPIKey(o))
	}
	if o, ok := os.LookupEnv("INCREASE_WEBHOOK_SECRET"); ok {
		defaults = append(defaults, option.WithWebhookSecret(o))
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
	r.Cards = NewCardService(opts...)
	r.CardDisputes = NewCardDisputeService(opts...)
	r.CardPurchaseSupplements = NewCardPurchaseSupplementService(opts...)
	r.ExternalAccounts = NewExternalAccountService(opts...)
	r.Exports = NewExportService(opts...)
	r.DigitalWalletTokens = NewDigitalWalletTokenService(opts...)
	r.Transactions = NewTransactionService(opts...)
	r.PendingTransactions = NewPendingTransactionService(opts...)
	r.Programs = NewProgramService(opts...)
	r.DeclinedTransactions = NewDeclinedTransactionService(opts...)
	r.AccountTransfers = NewAccountTransferService(opts...)
	r.ACHTransfers = NewACHTransferService(opts...)
	r.ACHPrenotifications = NewACHPrenotificationService(opts...)
	r.Documents = NewDocumentService(opts...)
	r.WireTransfers = NewWireTransferService(opts...)
	r.CheckTransfers = NewCheckTransferService(opts...)
	r.Entities = NewEntityService(opts...)
	r.InboundACHTransfers = NewInboundACHTransferService(opts...)
	r.InboundWireDrawdownRequests = NewInboundWireDrawdownRequestService(opts...)
	r.WireDrawdownRequests = NewWireDrawdownRequestService(opts...)
	r.Events = NewEventService(opts...)
	r.EventSubscriptions = NewEventSubscriptionService(opts...)
	r.Files = NewFileService(opts...)
	r.Groups = NewGroupService(opts...)
	r.OAuthConnections = NewOAuthConnectionService(opts...)
	r.CheckDeposits = NewCheckDepositService(opts...)
	r.RoutingNumbers = NewRoutingNumberService(opts...)
	r.AccountStatements = NewAccountStatementService(opts...)
	r.Simulations = NewSimulationService(opts...)
	r.PhysicalCards = NewPhysicalCardService(opts...)
	r.CardPayments = NewCardPaymentService(opts...)
	r.ProofOfAuthorizationRequests = NewProofOfAuthorizationRequestService(opts...)
	r.ProofOfAuthorizationRequestSubmissions = NewProofOfAuthorizationRequestSubmissionService(opts...)
	r.Intrafi = NewIntrafiService(opts...)
	r.RealTimePaymentsRequestForPayments = NewRealTimePaymentsRequestForPaymentService(opts...)
	r.Webhooks = NewWebhookService(opts...)
	r.OAuthTokens = NewOAuthTokenService(opts...)
	r.InboundWireTransfers = NewInboundWireTransferService(opts...)
	r.DigitalCardProfiles = NewDigitalCardProfileService(opts...)
	r.PhysicalCardProfiles = NewPhysicalCardProfileService(opts...)

	return
}

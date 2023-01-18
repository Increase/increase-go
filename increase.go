package increase

import "context"
import "increase/core"
import "increase/accounts"
import "increase/account_numbers"
import "increase/real_time_decisions"
import "increase/cards"
import "increase/card_disputes"
import "increase/card_profiles"
import "increase/external_accounts"
import "increase/digital_wallet_tokens"
import "increase/transactions"
import "increase/pending_transactions"
import "increase/declined_transactions"
import "increase/limits"
import "increase/account_transfers"
import "increase/ach_transfers"
import "increase/ach_prenotifications"
import "increase/wire_transfers"
import "increase/check_transfers"
import "increase/entities"
import "increase/wire_drawdown_requests"
import "increase/events"
import "increase/event_subscriptions"
import "increase/files"
import "increase/groups"
import "increase/oauth_connections"
import "increase/check_deposits"
import "increase/routing_numbers"
import "increase/account_statements"
import "increase/simulations"
import "fmt"

type RequestOpts = core.RequestOpts

func P[T ~bool | ~float32 | ~float64 | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~string | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~complex64 | ~complex128](v T) *T {
	return &v
}

type Increase struct {
	*ClientOptions
	get                  func(context.Context, string, *core.CoreRequest, interface{}) error
	post                 func(context.Context, string, *core.CoreRequest, interface{}) error
	patch                func(context.Context, string, *core.CoreRequest, interface{}) error
	put                  func(context.Context, string, *core.CoreRequest, interface{}) error
	delete               func(context.Context, string, *core.CoreRequest, interface{}) error
	Accounts             *accounts.AccountService
	AccountNumbers       *account_numbers.AccountNumberService
	RealTimeDecisions    *real_time_decisions.RealTimeDecisionService
	Cards                *cards.CardService
	CardDisputes         *card_disputes.CardDisputeService
	CardProfiles         *card_profiles.CardProfileService
	ExternalAccounts     *external_accounts.ExternalAccountService
	DigitalWalletTokens  *digital_wallet_tokens.DigitalWalletTokenService
	Transactions         *transactions.TransactionService
	PendingTransactions  *pending_transactions.PendingTransactionService
	DeclinedTransactions *declined_transactions.DeclinedTransactionService
	Limits               *limits.LimitService
	AccountTransfers     *account_transfers.AccountTransferService
	ACHTransfers         *ach_transfers.ACHTransferService
	ACHPrenotifications  *ach_prenotifications.ACHPrenotificationService
	WireTransfers        *wire_transfers.WireTransferService
	CheckTransfers       *check_transfers.CheckTransferService
	Entities             *entities.EntityService
	WireDrawdownRequests *wire_drawdown_requests.WireDrawdownRequestService
	Events               *events.EventService
	EventSubscriptions   *event_subscriptions.EventSubscriptionService
	Files                *files.FileService
	Groups               *groups.GroupService
	OauthConnections     *oauth_connections.OauthConnectionService
	CheckDeposits        *check_deposits.CheckDepositService
	RoutingNumbers       *routing_numbers.RoutingNumberService
	AccountStatements    *account_statements.AccountStatementService
	Simulations          *simulations.SimulationService
}

func NewIncreaseWithOptions(p ClientOptions) (r *Increase) {
	r = &Increase{}
	p.LoadDefaults()
	r.ClientOptions = &p

	if p.Requester == nil {
		p.Requester = &core.CoreClient{}
	}
	r.Requester = p.Requester

	if len(r.Requester.BaseURL) == 0 {
		if len(r.BaseURL) != 0 {
			r.Requester.BaseURL = r.BaseURL
		} else {
			r.Requester.BaseURL = "https://api.increase.com"
		}
	}

	r.Requester.AuthHeaders = func() map[string]string {
		return map[string]string{"Authorization": fmt.Sprintf("Bearer %s", r.APIKey)}
	}()

	r.Accounts = accounts.NewAccountService(r.Requester)

	r.AccountNumbers = account_numbers.NewAccountNumberService(r.Requester)

	r.RealTimeDecisions = real_time_decisions.NewRealTimeDecisionService(r.Requester)

	r.Cards = cards.NewCardService(r.Requester)

	r.CardDisputes = card_disputes.NewCardDisputeService(r.Requester)

	r.CardProfiles = card_profiles.NewCardProfileService(r.Requester)

	r.ExternalAccounts = external_accounts.NewExternalAccountService(r.Requester)

	r.DigitalWalletTokens = digital_wallet_tokens.NewDigitalWalletTokenService(r.Requester)

	r.Transactions = transactions.NewTransactionService(r.Requester)

	r.PendingTransactions = pending_transactions.NewPendingTransactionService(r.Requester)

	r.DeclinedTransactions = declined_transactions.NewDeclinedTransactionService(r.Requester)

	r.Limits = limits.NewLimitService(r.Requester)

	r.AccountTransfers = account_transfers.NewAccountTransferService(r.Requester)

	r.ACHTransfers = ach_transfers.NewACHTransferService(r.Requester)

	r.ACHPrenotifications = ach_prenotifications.NewACHPrenotificationService(r.Requester)

	r.WireTransfers = wire_transfers.NewWireTransferService(r.Requester)

	r.CheckTransfers = check_transfers.NewCheckTransferService(r.Requester)

	r.Entities = entities.NewEntityService(r.Requester)

	r.WireDrawdownRequests = wire_drawdown_requests.NewWireDrawdownRequestService(r.Requester)

	r.Events = events.NewEventService(r.Requester)

	r.EventSubscriptions = event_subscriptions.NewEventSubscriptionService(r.Requester)

	r.Files = files.NewFileService(r.Requester)

	r.Groups = groups.NewGroupService(r.Requester)

	r.OauthConnections = oauth_connections.NewOauthConnectionService(r.Requester)

	r.CheckDeposits = check_deposits.NewCheckDepositService(r.Requester)

	r.RoutingNumbers = routing_numbers.NewRoutingNumberService(r.Requester)

	r.AccountStatements = account_statements.NewAccountStatementService(r.Requester)

	r.Simulations = simulations.NewSimulationService(r.Requester)

	return
}

func NewIncrease() *Increase {
	return NewIncreaseWithOptions(ClientOptions{})
}

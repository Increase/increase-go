package increase

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

type Increase struct {
	Requester            *core.CoreClient
	get                  func(string, *core.CoreRequest, interface{}) error
	post                 func(string, *core.CoreRequest, interface{}) error
	patch                func(string, *core.CoreRequest, interface{}) error
	put                  func(string, *core.CoreRequest, interface{}) error
	delete               func(string, *core.CoreRequest, interface{}) error
	APIKey               string
	Accounts             *accounts.Accounts
	AccountNumbers       *account_numbers.AccountNumbers
	RealTimeDecisions    *real_time_decisions.RealTimeDecisions
	Cards                *cards.Cards
	CardDisputes         *card_disputes.CardDisputes
	CardProfiles         *card_profiles.CardProfiles
	ExternalAccounts     *external_accounts.ExternalAccounts
	DigitalWalletTokens  *digital_wallet_tokens.DigitalWalletTokens
	Transactions         *transactions.Transactions
	PendingTransactions  *pending_transactions.PendingTransactions
	DeclinedTransactions *declined_transactions.DeclinedTransactions
	Limits               *limits.Limits
	AccountTransfers     *account_transfers.AccountTransfers
	ACHTransfers         *ach_transfers.ACHTransfers
	ACHPrenotifications  *ach_prenotifications.ACHPrenotifications
	WireTransfers        *wire_transfers.WireTransfers
	CheckTransfers       *check_transfers.CheckTransfers
	Entities             *entities.Entities
	WireDrawdownRequests *wire_drawdown_requests.WireDrawdownRequests
	Events               *events.Events
	EventSubscriptions   *event_subscriptions.EventSubscriptions
	Files                *files.Files
	Groups               *groups.Groups
	OauthConnections     *oauth_connections.OauthConnections
	CheckDeposits        *check_deposits.CheckDeposits
	RoutingNumbers       *routing_numbers.RoutingNumbers
	AccountStatements    *account_statements.AccountStatements
	Simulations          *simulations.Simulations
}

func NewIncrease(requster *core.CoreClient) (r *Increase) {
	r.Requester = requster
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	r = NewIncreaseWithParams(IncreaseParams{})
	return
}

func NewIncreaseWithParams(p IncreaseParams) (r *Increase) {
	r = &Increase{}

	if p.Client == nil {
		p.Client = &core.CoreClient{}
	}
	r.Requester = p.Client
	if len(r.Requester.BaseURL) == 0 {
		if len(p.BaseURL) != 0 {
			r.Requester.BaseURL = p.BaseURL
		} else {
			r.Requester.BaseURL = "https://api.increase.com"
		}
	}
	r.APIKey = p.APIKey
	r.Requester.AuthHeaders = func() map[string]string {
		return map[string]string{"Authorization": fmt.Sprintf("Bearer %s", r.APIKey)}
	}
	r.Accounts = accounts.NewAccounts(r.Requester)
	r.AccountNumbers = account_numbers.NewAccountNumbers(r.Requester)
	r.RealTimeDecisions = real_time_decisions.NewRealTimeDecisions(r.Requester)
	r.Cards = cards.NewCards(r.Requester)
	r.CardDisputes = card_disputes.NewCardDisputes(r.Requester)
	r.CardProfiles = card_profiles.NewCardProfiles(r.Requester)
	r.ExternalAccounts = external_accounts.NewExternalAccounts(r.Requester)
	r.DigitalWalletTokens = digital_wallet_tokens.NewDigitalWalletTokens(r.Requester)
	r.Transactions = transactions.NewTransactions(r.Requester)
	r.PendingTransactions = pending_transactions.NewPendingTransactions(r.Requester)
	r.DeclinedTransactions = declined_transactions.NewDeclinedTransactions(r.Requester)
	r.Limits = limits.NewLimits(r.Requester)
	r.AccountTransfers = account_transfers.NewAccountTransfers(r.Requester)
	r.ACHTransfers = ach_transfers.NewACHTransfers(r.Requester)
	r.ACHPrenotifications = ach_prenotifications.NewACHPrenotifications(r.Requester)
	r.WireTransfers = wire_transfers.NewWireTransfers(r.Requester)
	r.CheckTransfers = check_transfers.NewCheckTransfers(r.Requester)
	r.Entities = entities.NewEntities(r.Requester)
	r.WireDrawdownRequests = wire_drawdown_requests.NewWireDrawdownRequests(r.Requester)
	r.Events = events.NewEvents(r.Requester)
	r.EventSubscriptions = event_subscriptions.NewEventSubscriptions(r.Requester)
	r.Files = files.NewFiles(r.Requester)
	r.Groups = groups.NewGroups(r.Requester)
	r.OauthConnections = oauth_connections.NewOauthConnections(r.Requester)
	r.CheckDeposits = check_deposits.NewCheckDeposits(r.Requester)
	r.RoutingNumbers = routing_numbers.NewRoutingNumbers(r.Requester)
	r.AccountStatements = account_statements.NewAccountStatements(r.Requester)
	r.Simulations = simulations.NewSimulations(r.Requester)
	return
}

type IncreaseParams struct {
	Client  *core.CoreClient
	BaseURL string
	APIKey  string
}

type RequestOpts = core.RequestOpts

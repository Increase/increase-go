package increase

import (
	"context"
	"fmt"
	"increase/core"
	"increase/services"
)

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
	Accounts             *services.AccountService
	AccountNumbers       *services.AccountNumberService
	RealTimeDecisions    *services.RealTimeDecisionService
	Cards                *services.CardService
	CardDisputes         *services.CardDisputeService
	CardProfiles         *services.CardProfileService
	ExternalAccounts     *services.ExternalAccountService
	DigitalWalletTokens  *services.DigitalWalletTokenService
	Transactions         *services.TransactionService
	PendingTransactions  *services.PendingTransactionService
	DeclinedTransactions *services.DeclinedTransactionService
	Limits               *services.LimitService
	AccountTransfers     *services.AccountTransferService
	ACHTransfers         *services.ACHTransferService
	ACHPrenotifications  *services.ACHPrenotificationService
	WireTransfers        *services.WireTransferService
	CheckTransfers       *services.CheckTransferService
	Entities             *services.EntityService
	WireDrawdownRequests *services.WireDrawdownRequestService
	Events               *services.EventService
	EventSubscriptions   *services.EventSubscriptionService
	Files                *services.FileService
	Groups               *services.GroupService
	OauthConnections     *services.OauthConnectionService
	CheckDeposits        *services.CheckDepositService
	RoutingNumbers       *services.RoutingNumberService
	AccountStatements    *services.AccountStatementService
	Simulations          *services.SimulationService
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

	r.Accounts = services.NewAccountService(r.Requester)

	r.AccountNumbers = services.NewAccountNumberService(r.Requester)

	r.RealTimeDecisions = services.NewRealTimeDecisionService(r.Requester)

	r.Cards = services.NewCardService(r.Requester)

	r.CardDisputes = services.NewCardDisputeService(r.Requester)

	r.CardProfiles = services.NewCardProfileService(r.Requester)

	r.ExternalAccounts = services.NewExternalAccountService(r.Requester)

	r.DigitalWalletTokens = services.NewDigitalWalletTokenService(r.Requester)

	r.Transactions = services.NewTransactionService(r.Requester)

	r.PendingTransactions = services.NewPendingTransactionService(r.Requester)

	r.DeclinedTransactions = services.NewDeclinedTransactionService(r.Requester)

	r.Limits = services.NewLimitService(r.Requester)

	r.AccountTransfers = services.NewAccountTransferService(r.Requester)

	r.ACHTransfers = services.NewACHTransferService(r.Requester)

	r.ACHPrenotifications = services.NewACHPrenotificationService(r.Requester)

	r.WireTransfers = services.NewWireTransferService(r.Requester)

	r.CheckTransfers = services.NewCheckTransferService(r.Requester)

	r.Entities = services.NewEntityService(r.Requester)

	r.WireDrawdownRequests = services.NewWireDrawdownRequestService(r.Requester)

	r.Events = services.NewEventService(r.Requester)

	r.EventSubscriptions = services.NewEventSubscriptionService(r.Requester)

	r.Files = services.NewFileService(r.Requester)

	r.Groups = services.NewGroupService(r.Requester)

	r.OauthConnections = services.NewOauthConnectionService(r.Requester)

	r.CheckDeposits = services.NewCheckDepositService(r.Requester)

	r.RoutingNumbers = services.NewRoutingNumberService(r.Requester)

	r.AccountStatements = services.NewAccountStatementService(r.Requester)

	r.Simulations = services.NewSimulationService(r.Requester)

	return
}

func NewIncrease() *Increase {
	return NewIncreaseWithOptions(ClientOptions{})
}

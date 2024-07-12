// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"net/http"
	"os"

	"github.com/increase/increase-go/internal/requestconfig"
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
	InboundCheckDeposits                   *InboundCheckDepositService
	InboundMailItems                       *InboundMailItemService
	Lockboxes                              *LockboxService
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
	r.InboundCheckDeposits = NewInboundCheckDepositService(opts...)
	r.InboundMailItems = NewInboundMailItemService(opts...)
	r.Lockboxes = NewLockboxService(opts...)

	return
}

// Execute makes a request with the given context, method, URL, request params,
// response, and request options. This is useful for hitting undocumented endpoints
// while retaining the base URL, auth, retries, and other options from the client.
//
// If a byte slice or an [io.Reader] is supplied to params, it will be used as-is
// for the request body.
//
// The params is by default serialized into the body using [encoding/json]. If your
// type implements a MarshalJSON function, it will be used instead to serialize the
// request. If a URLQuery method is implemented, the returned [url.Values] will be
// used as query strings to the url.
//
// If your params struct uses [param.Field], you must provide either [MarshalJSON],
// [URLQuery], and/or [MarshalForm] functions. It is undefined behavior to use a
// struct uses [param.Field] without specifying how it is serialized.
//
// Any "…Params" object defined in this library can be used as the request
// argument. Note that 'path' arguments will not be forwarded into the url.
//
// The response body will be deserialized into the res variable, depending on its
// type:
//
//   - A pointer to a [*http.Response] is populated by the raw response.
//   - A pointer to a byte array will be populated with the contents of the request
//     body.
//   - A pointer to any other type uses this library's default JSON decoding, which
//     respects UnmarshalJSON if it is defined on the type.
//   - A nil value will not read the response body.
//
// For even greater flexibility, see [option.WithResponseInto] and
// [option.WithResponseBodyInto].
func (r *Client) Execute(ctx context.Context, method string, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	opts = append(r.Options, opts...)
	return requestconfig.ExecuteNewRequest(ctx, method, path, params, res, opts...)
}

// Get makes a GET request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Get(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodGet, path, params, res, opts...)
}

// Post makes a POST request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Post(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPost, path, params, res, opts...)
}

// Put makes a PUT request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Put(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPut, path, params, res, opts...)
}

// Patch makes a PATCH request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Patch(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPatch, path, params, res, opts...)
}

// Delete makes a DELETE request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Delete(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodDelete, path, params, res, opts...)
}

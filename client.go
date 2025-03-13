// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"net/http"
	"os"

	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
)

// Client creates a struct with services and top level methods that help with
// interacting with the increase API. You should not instantiate this client
// directly, and instead use the [NewClient] method instead.
type Client struct {
	Options                                []option.RequestOption
	Accounts                               *AccountService
	AccountNumbers                         *AccountNumberService
	Cards                                  *CardService
	CardPayments                           *CardPaymentService
	CardPurchaseSupplements                *CardPurchaseSupplementService
	CardDisputes                           *CardDisputeService
	PhysicalCards                          *PhysicalCardService
	DigitalCardProfiles                    *DigitalCardProfileService
	PhysicalCardProfiles                   *PhysicalCardProfileService
	DigitalWalletTokens                    *DigitalWalletTokenService
	Transactions                           *TransactionService
	PendingTransactions                    *PendingTransactionService
	DeclinedTransactions                   *DeclinedTransactionService
	AccountTransfers                       *AccountTransferService
	ACHTransfers                           *ACHTransferService
	ACHPrenotifications                    *ACHPrenotificationService
	InboundACHTransfers                    *InboundACHTransferService
	WireTransfers                          *WireTransferService
	InboundWireTransfers                   *InboundWireTransferService
	WireDrawdownRequests                   *WireDrawdownRequestService
	InboundWireDrawdownRequests            *InboundWireDrawdownRequestService
	CheckTransfers                         *CheckTransferService
	InboundCheckDeposits                   *InboundCheckDepositService
	RealTimePaymentsTransfers              *RealTimePaymentsTransferService
	InboundRealTimePaymentsTransfers       *InboundRealTimePaymentsTransferService
	CheckDeposits                          *CheckDepositService
	Lockboxes                              *LockboxService
	InboundMailItems                       *InboundMailItemService
	RoutingNumbers                         *RoutingNumberService
	ExternalAccounts                       *ExternalAccountService
	Entities                               *EntityService
	SupplementalDocuments                  *SupplementalDocumentService
	Programs                               *ProgramService
	ProofOfAuthorizationRequests           *ProofOfAuthorizationRequestService
	ProofOfAuthorizationRequestSubmissions *ProofOfAuthorizationRequestSubmissionService
	AccountStatements                      *AccountStatementService
	Files                                  *FileService
	Documents                              *DocumentService
	Exports                                *ExportService
	Events                                 *EventService
	EventSubscriptions                     *EventSubscriptionService
	RealTimeDecisions                      *RealTimeDecisionService
	BookkeepingAccounts                    *BookkeepingAccountService
	BookkeepingEntrySets                   *BookkeepingEntrySetService
	BookkeepingEntries                     *BookkeepingEntryService
	Groups                                 *GroupService
	OAuthApplications                      *OAuthApplicationService
	OAuthConnections                       *OAuthConnectionService
	Webhooks                               *WebhookService
	OAuthTokens                            *OAuthTokenService
	IntrafiAccountEnrollments              *IntrafiAccountEnrollmentService
	IntrafiBalances                        *IntrafiBalanceService
	IntrafiExclusions                      *IntrafiExclusionService
	Simulations                            *SimulationService
}

// DefaultClientOptions read from the environment (INCREASE_API_KEY,
// INCREASE_WEBHOOK_SECRET). This should be used to initialize new clients.
func DefaultClientOptions() []option.RequestOption {
	defaults := []option.RequestOption{option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("INCREASE_API_KEY"); ok {
		defaults = append(defaults, option.WithAPIKey(o))
	}
	if o, ok := os.LookupEnv("INCREASE_WEBHOOK_SECRET"); ok {
		defaults = append(defaults, option.WithWebhookSecret(o))
	}
	return defaults
}

// NewClient generates a new client with the default option read from the
// environment (INCREASE_API_KEY, INCREASE_WEBHOOK_SECRET). The option passed in as
// arguments are applied after these default arguments, and all option will be
// passed down to the services and requests that this client makes.
func NewClient(opts ...option.RequestOption) (r *Client) {
	opts = append(DefaultClientOptions(), opts...)

	r = &Client{Options: opts}

	r.Accounts = NewAccountService(opts...)
	r.AccountNumbers = NewAccountNumberService(opts...)
	r.Cards = NewCardService(opts...)
	r.CardPayments = NewCardPaymentService(opts...)
	r.CardPurchaseSupplements = NewCardPurchaseSupplementService(opts...)
	r.CardDisputes = NewCardDisputeService(opts...)
	r.PhysicalCards = NewPhysicalCardService(opts...)
	r.DigitalCardProfiles = NewDigitalCardProfileService(opts...)
	r.PhysicalCardProfiles = NewPhysicalCardProfileService(opts...)
	r.DigitalWalletTokens = NewDigitalWalletTokenService(opts...)
	r.Transactions = NewTransactionService(opts...)
	r.PendingTransactions = NewPendingTransactionService(opts...)
	r.DeclinedTransactions = NewDeclinedTransactionService(opts...)
	r.AccountTransfers = NewAccountTransferService(opts...)
	r.ACHTransfers = NewACHTransferService(opts...)
	r.ACHPrenotifications = NewACHPrenotificationService(opts...)
	r.InboundACHTransfers = NewInboundACHTransferService(opts...)
	r.WireTransfers = NewWireTransferService(opts...)
	r.InboundWireTransfers = NewInboundWireTransferService(opts...)
	r.WireDrawdownRequests = NewWireDrawdownRequestService(opts...)
	r.InboundWireDrawdownRequests = NewInboundWireDrawdownRequestService(opts...)
	r.CheckTransfers = NewCheckTransferService(opts...)
	r.InboundCheckDeposits = NewInboundCheckDepositService(opts...)
	r.RealTimePaymentsTransfers = NewRealTimePaymentsTransferService(opts...)
	r.InboundRealTimePaymentsTransfers = NewInboundRealTimePaymentsTransferService(opts...)
	r.CheckDeposits = NewCheckDepositService(opts...)
	r.Lockboxes = NewLockboxService(opts...)
	r.InboundMailItems = NewInboundMailItemService(opts...)
	r.RoutingNumbers = NewRoutingNumberService(opts...)
	r.ExternalAccounts = NewExternalAccountService(opts...)
	r.Entities = NewEntityService(opts...)
	r.SupplementalDocuments = NewSupplementalDocumentService(opts...)
	r.Programs = NewProgramService(opts...)
	r.ProofOfAuthorizationRequests = NewProofOfAuthorizationRequestService(opts...)
	r.ProofOfAuthorizationRequestSubmissions = NewProofOfAuthorizationRequestSubmissionService(opts...)
	r.AccountStatements = NewAccountStatementService(opts...)
	r.Files = NewFileService(opts...)
	r.Documents = NewDocumentService(opts...)
	r.Exports = NewExportService(opts...)
	r.Events = NewEventService(opts...)
	r.EventSubscriptions = NewEventSubscriptionService(opts...)
	r.RealTimeDecisions = NewRealTimeDecisionService(opts...)
	r.BookkeepingAccounts = NewBookkeepingAccountService(opts...)
	r.BookkeepingEntrySets = NewBookkeepingEntrySetService(opts...)
	r.BookkeepingEntries = NewBookkeepingEntryService(opts...)
	r.Groups = NewGroupService(opts...)
	r.OAuthApplications = NewOAuthApplicationService(opts...)
	r.OAuthConnections = NewOAuthConnectionService(opts...)
	r.Webhooks = NewWebhookService(opts...)
	r.OAuthTokens = NewOAuthTokenService(opts...)
	r.IntrafiAccountEnrollments = NewIntrafiAccountEnrollmentService(opts...)
	r.IntrafiBalances = NewIntrafiBalanceService(opts...)
	r.IntrafiExclusions = NewIntrafiExclusionService(opts...)
	r.Simulations = NewSimulationService(opts...)

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

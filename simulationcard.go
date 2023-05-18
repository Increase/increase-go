package increase

import (
	"context"
	"net/http"
	"time"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/internal/shared"
	"github.com/increase/increase-go/option"
)

// SimulationCardService contains methods and other services that help with
// interacting with the increase API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSimulationCardService] method
// instead.
type SimulationCardService struct {
	Options []option.RequestOption
}

// NewSimulationCardService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSimulationCardService(opts ...option.RequestOption) (r *SimulationCardService) {
	r = &SimulationCardService{}
	r.Options = opts
	return
}

// Simulates a purchase authorization on a [Card](#cards). Depending on the balance
// available to the card and the `amount` submitted, the authorization activity
// will result in a [Pending Transaction](#pending-transactions) of type
// `card_authorization` or a [Declined Transaction](#declined-transactions) of type
// `card_decline`. You can pass either a Card id or a
// [Digital Wallet Token](#digital-wallet-tokens) id to simulate the two different
// ways purchases can be made.
func (r *SimulationCardService) Authorize(ctx context.Context, body SimulationCardAuthorizeParams, opts ...option.RequestOption) (res *CardAuthorizationSimulation, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/card_authorizations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Simulates the settlement of an authorization by a card acquirer. After a card
// authorization is created, the merchant will eventually send a settlement. This
// simulates that event, which may occur many days after the purchase in
// production. The amount settled can be different from the amount originally
// authorized, for example, when adding a tip to a restaurant bill.
func (r *SimulationCardService) Settlement(ctx context.Context, body SimulationCardSettlementParams, opts ...option.RequestOption) (res *Transaction, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/card_settlements"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// The results of a Card Authorization simulation.
type CardAuthorizationSimulation struct {
	// If the authorization attempt succeeds, this will contain the resulting Pending
	// Transaction object. The Pending Transaction's `source` will be of
	// `category: card_authorization`.
	PendingTransaction CardAuthorizationSimulationPendingTransaction `json:"pending_transaction,required,nullable"`
	// If the authorization attempt fails, this will contain the resulting
	// [Declined Transaction](#declined-transactions) object. The Declined
	// Transaction's `source` will be of `category: card_decline`.
	DeclinedTransaction CardAuthorizationSimulationDeclinedTransaction `json:"declined_transaction,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `inbound_card_authorization_simulation_result`.
	Type CardAuthorizationSimulationType `json:"type,required"`
	JSON cardAuthorizationSimulationJSON
}

// cardAuthorizationSimulationJSON contains the JSON metadata for the struct
// [CardAuthorizationSimulation]
type cardAuthorizationSimulationJSON struct {
	PendingTransaction  apijson.Field
	DeclinedTransaction apijson.Field
	Type                apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CardAuthorizationSimulation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// If the authorization attempt succeeds, this will contain the resulting Pending
// Transaction object. The Pending Transaction's `source` will be of
// `category: card_authorization`.
type CardAuthorizationSimulationPendingTransaction struct {
	// The identifier for the account this Pending Transaction belongs to.
	AccountID string `json:"account_id,required"`
	// The Pending Transaction amount in the minor unit of its currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Pending
	// Transaction's currency. This will match the currency on the Pending
	// Transcation's Account.
	Currency CardAuthorizationSimulationPendingTransactionCurrency `json:"currency,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the Pending
	// Transaction was completed.
	CompletedAt time.Time `json:"completed_at,required,nullable" format:"date-time"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the Pending
	// Transaction occured.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// For a Pending Transaction related to a transfer, this is the description you
	// provide. For a Pending Transaction related to a payment, this is the description
	// the vendor provides.
	Description string `json:"description,required"`
	// The Pending Transaction identifier.
	ID string `json:"id,required"`
	// The identifier for the route this Pending Transaction came through. Routes are
	// things like cards and ACH details.
	RouteID string `json:"route_id,required,nullable"`
	// The type of the route this Pending Transaction came through.
	RouteType CardAuthorizationSimulationPendingTransactionRouteType `json:"route_type,required,nullable"`
	// This is an object giving more details on the network-level event that caused the
	// Pending Transaction. For example, for a card transaction this lists the
	// merchant's industry and location.
	Source CardAuthorizationSimulationPendingTransactionSource `json:"source,required"`
	// Whether the Pending Transaction has been confirmed and has an associated
	// Transaction.
	Status CardAuthorizationSimulationPendingTransactionStatus `json:"status,required"`
	// A constant representing the object's type. For this resource it will always be
	// `pending_transaction`.
	Type CardAuthorizationSimulationPendingTransactionType `json:"type,required"`
	JSON cardAuthorizationSimulationPendingTransactionJSON
}

// cardAuthorizationSimulationPendingTransactionJSON contains the JSON metadata for
// the struct [CardAuthorizationSimulationPendingTransaction]
type cardAuthorizationSimulationPendingTransactionJSON struct {
	AccountID   apijson.Field
	Amount      apijson.Field
	Currency    apijson.Field
	CompletedAt apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	ID          apijson.Field
	RouteID     apijson.Field
	RouteType   apijson.Field
	Source      apijson.Field
	Status      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationSimulationPendingTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CardAuthorizationSimulationPendingTransactionCurrency string

const (
	CardAuthorizationSimulationPendingTransactionCurrencyCad CardAuthorizationSimulationPendingTransactionCurrency = "CAD"
	CardAuthorizationSimulationPendingTransactionCurrencyChf CardAuthorizationSimulationPendingTransactionCurrency = "CHF"
	CardAuthorizationSimulationPendingTransactionCurrencyEur CardAuthorizationSimulationPendingTransactionCurrency = "EUR"
	CardAuthorizationSimulationPendingTransactionCurrencyGbp CardAuthorizationSimulationPendingTransactionCurrency = "GBP"
	CardAuthorizationSimulationPendingTransactionCurrencyJpy CardAuthorizationSimulationPendingTransactionCurrency = "JPY"
	CardAuthorizationSimulationPendingTransactionCurrencyUsd CardAuthorizationSimulationPendingTransactionCurrency = "USD"
)

type CardAuthorizationSimulationPendingTransactionRouteType string

const (
	CardAuthorizationSimulationPendingTransactionRouteTypeAccountNumber CardAuthorizationSimulationPendingTransactionRouteType = "account_number"
	CardAuthorizationSimulationPendingTransactionRouteTypeCard          CardAuthorizationSimulationPendingTransactionRouteType = "card"
)

// This is an object giving more details on the network-level event that caused the
// Pending Transaction. For example, for a card transaction this lists the
// merchant's industry and location.
type CardAuthorizationSimulationPendingTransactionSource struct {
	// The type of transaction that took place. We may add additional possible values
	// for this enum over time; your application should be able to handle such
	// additions gracefully.
	Category CardAuthorizationSimulationPendingTransactionSourceCategory `json:"category,required"`
	// A Account Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `account_transfer_instruction`.
	AccountTransferInstruction CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstruction `json:"account_transfer_instruction,required,nullable"`
	// A ACH Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `ach_transfer_instruction`.
	ACHTransferInstruction CardAuthorizationSimulationPendingTransactionSourceACHTransferInstruction `json:"ach_transfer_instruction,required,nullable"`
	// A Card Authorization object. This field will be present in the JSON response if
	// and only if `category` is equal to `card_authorization`.
	CardAuthorization CardAuthorizationSimulationPendingTransactionSourceCardAuthorization `json:"card_authorization,required,nullable"`
	// A Check Deposit Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_deposit_instruction`.
	CheckDepositInstruction CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstruction `json:"check_deposit_instruction,required,nullable"`
	// A Check Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_instruction`.
	CheckTransferInstruction CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstruction `json:"check_transfer_instruction,required,nullable"`
	// A Inbound Funds Hold object. This field will be present in the JSON response if
	// and only if `category` is equal to `inbound_funds_hold`.
	InboundFundsHold CardAuthorizationSimulationPendingTransactionSourceInboundFundsHold `json:"inbound_funds_hold,required,nullable"`
	// A Deprecated Card Authorization object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_route_authorization`.
	CardRouteAuthorization CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorization `json:"card_route_authorization,required,nullable"`
	// A Real Time Payments Transfer Instruction object. This field will be present in
	// the JSON response if and only if `category` is equal to
	// `real_time_payments_transfer_instruction`.
	RealTimePaymentsTransferInstruction CardAuthorizationSimulationPendingTransactionSourceRealTimePaymentsTransferInstruction `json:"real_time_payments_transfer_instruction,required,nullable"`
	// A Wire Drawdown Payment Instruction object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `wire_drawdown_payment_instruction`.
	WireDrawdownPaymentInstruction CardAuthorizationSimulationPendingTransactionSourceWireDrawdownPaymentInstruction `json:"wire_drawdown_payment_instruction,required,nullable"`
	// A Wire Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_instruction`.
	WireTransferInstruction CardAuthorizationSimulationPendingTransactionSourceWireTransferInstruction `json:"wire_transfer_instruction,required,nullable"`
	JSON                    cardAuthorizationSimulationPendingTransactionSourceJSON
}

// cardAuthorizationSimulationPendingTransactionSourceJSON contains the JSON
// metadata for the struct [CardAuthorizationSimulationPendingTransactionSource]
type cardAuthorizationSimulationPendingTransactionSourceJSON struct {
	Category                            apijson.Field
	AccountTransferInstruction          apijson.Field
	ACHTransferInstruction              apijson.Field
	CardAuthorization                   apijson.Field
	CheckDepositInstruction             apijson.Field
	CheckTransferInstruction            apijson.Field
	InboundFundsHold                    apijson.Field
	CardRouteAuthorization              apijson.Field
	RealTimePaymentsTransferInstruction apijson.Field
	WireDrawdownPaymentInstruction      apijson.Field
	WireTransferInstruction             apijson.Field
	raw                                 string
	ExtraFields                         map[string]apijson.Field
}

func (r *CardAuthorizationSimulationPendingTransactionSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CardAuthorizationSimulationPendingTransactionSourceCategory string

const (
	CardAuthorizationSimulationPendingTransactionSourceCategoryAccountTransferInstruction          CardAuthorizationSimulationPendingTransactionSourceCategory = "account_transfer_instruction"
	CardAuthorizationSimulationPendingTransactionSourceCategoryACHTransferInstruction              CardAuthorizationSimulationPendingTransactionSourceCategory = "ach_transfer_instruction"
	CardAuthorizationSimulationPendingTransactionSourceCategoryCardAuthorization                   CardAuthorizationSimulationPendingTransactionSourceCategory = "card_authorization"
	CardAuthorizationSimulationPendingTransactionSourceCategoryCheckDepositInstruction             CardAuthorizationSimulationPendingTransactionSourceCategory = "check_deposit_instruction"
	CardAuthorizationSimulationPendingTransactionSourceCategoryCheckTransferInstruction            CardAuthorizationSimulationPendingTransactionSourceCategory = "check_transfer_instruction"
	CardAuthorizationSimulationPendingTransactionSourceCategoryInboundFundsHold                    CardAuthorizationSimulationPendingTransactionSourceCategory = "inbound_funds_hold"
	CardAuthorizationSimulationPendingTransactionSourceCategoryCardRouteAuthorization              CardAuthorizationSimulationPendingTransactionSourceCategory = "card_route_authorization"
	CardAuthorizationSimulationPendingTransactionSourceCategoryRealTimePaymentsTransferInstruction CardAuthorizationSimulationPendingTransactionSourceCategory = "real_time_payments_transfer_instruction"
	CardAuthorizationSimulationPendingTransactionSourceCategoryWireDrawdownPaymentInstruction      CardAuthorizationSimulationPendingTransactionSourceCategory = "wire_drawdown_payment_instruction"
	CardAuthorizationSimulationPendingTransactionSourceCategoryWireTransferInstruction             CardAuthorizationSimulationPendingTransactionSourceCategory = "wire_transfer_instruction"
	CardAuthorizationSimulationPendingTransactionSourceCategoryOther                               CardAuthorizationSimulationPendingTransactionSourceCategory = "other"
)

// A Account Transfer Instruction object. This field will be present in the JSON
// response if and only if `category` is equal to `account_transfer_instruction`.
type CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrency `json:"currency,required"`
	// The identifier of the Account Transfer that led to this Pending Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       cardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionJSON
}

// cardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionJSON
// contains the JSON metadata for the struct
// [CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstruction]
type cardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	TransferID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstruction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrency string

const (
	CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrencyCad CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrency = "CAD"
	CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrencyChf CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrency = "CHF"
	CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrencyEur CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrency = "EUR"
	CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrencyGbp CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrency = "GBP"
	CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrencyJpy CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrency = "JPY"
	CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrencyUsd CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrency = "USD"
)

// A ACH Transfer Instruction object. This field will be present in the JSON
// response if and only if `category` is equal to `ach_transfer_instruction`.
type CardAuthorizationSimulationPendingTransactionSourceACHTransferInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The identifier of the ACH Transfer that led to this Pending Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       cardAuthorizationSimulationPendingTransactionSourceACHTransferInstructionJSON
}

// cardAuthorizationSimulationPendingTransactionSourceACHTransferInstructionJSON
// contains the JSON metadata for the struct
// [CardAuthorizationSimulationPendingTransactionSourceACHTransferInstruction]
type cardAuthorizationSimulationPendingTransactionSourceACHTransferInstructionJSON struct {
	Amount      apijson.Field
	TransferID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationSimulationPendingTransactionSourceACHTransferInstruction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Card Authorization object. This field will be present in the JSON response if
// and only if `category` is equal to `card_authorization`.
type CardAuthorizationSimulationPendingTransactionSourceCardAuthorization struct {
	// The Card Authorization identifier.
	ID string `json:"id,required"`
	// The merchant identifier (commonly abbreviated as MID) of the merchant the card
	// is transacting with.
	MerchantAcceptorID string `json:"merchant_acceptor_id,required"`
	// The merchant descriptor of the merchant the card is transacting with.
	MerchantDescriptor string `json:"merchant_descriptor,required"`
	// The Merchant Category Code (commonly abbreviated as MCC) of the merchant the
	// card is transacting with.
	MerchantCategoryCode string `json:"merchant_category_code,required,nullable"`
	// The city the merchant resides in.
	MerchantCity string `json:"merchant_city,required,nullable"`
	// The country the merchant resides in.
	MerchantCountry string `json:"merchant_country,required,nullable"`
	// The payment network used to process this card authorization
	Network CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetwork `json:"network,required"`
	// Fields specific to the `network`
	NetworkDetails CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetails `json:"network_details,required"`
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrency `json:"currency,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) when this authorization
	// will expire and the pending transaction will be released.
	ExpiresAt time.Time `json:"expires_at,required" format:"date-time"`
	// The identifier of the Real-Time Decision sent to approve or decline this
	// transaction.
	RealTimeDecisionID string `json:"real_time_decision_id,required,nullable"`
	// If the authorization was made via a Digital Wallet Token (such as an Apple Pay
	// purchase), the identifier of the token that was used.
	DigitalWalletTokenID string `json:"digital_wallet_token_id,required,nullable"`
	// The identifier of the Pending Transaction associated with this Transaction.
	PendingTransactionID string `json:"pending_transaction_id,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `card_authorization`.
	Type CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationType `json:"type,required"`
	JSON cardAuthorizationSimulationPendingTransactionSourceCardAuthorizationJSON
}

// cardAuthorizationSimulationPendingTransactionSourceCardAuthorizationJSON
// contains the JSON metadata for the struct
// [CardAuthorizationSimulationPendingTransactionSourceCardAuthorization]
type cardAuthorizationSimulationPendingTransactionSourceCardAuthorizationJSON struct {
	ID                   apijson.Field
	MerchantAcceptorID   apijson.Field
	MerchantDescriptor   apijson.Field
	MerchantCategoryCode apijson.Field
	MerchantCity         apijson.Field
	MerchantCountry      apijson.Field
	Network              apijson.Field
	NetworkDetails       apijson.Field
	Amount               apijson.Field
	Currency             apijson.Field
	ExpiresAt            apijson.Field
	RealTimeDecisionID   apijson.Field
	DigitalWalletTokenID apijson.Field
	PendingTransactionID apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CardAuthorizationSimulationPendingTransactionSourceCardAuthorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetwork string

const (
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkVisa CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetwork = "visa"
)

// Fields specific to the `network`
type CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetails struct {
	// Fields specific to the `visa` network
	Visa CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisa `json:"visa,required"`
	JSON cardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsJSON
}

// cardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsJSON
// contains the JSON metadata for the struct
// [CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetails]
type cardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsJSON struct {
	Visa        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Fields specific to the `visa` network
type CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisa struct {
	// For electronic commerce transactions, this identifies the level of security used
	// in obtaining the customer's payment credential. For mail or telephone order
	// transactions, identifies the type of mail or telephone order.
	ElectronicCommerceIndicator CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator `json:"electronic_commerce_indicator,required,nullable"`
	// The method used to enter the cardholder's primary account number and card
	// expiration date
	PointOfServiceEntryMode shared.PointOfServiceEntryMode `json:"point_of_service_entry_mode,required,nullable"`
	JSON                    cardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaJSON
}

// cardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaJSON
// contains the JSON metadata for the struct
// [CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisa]
type cardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaJSON struct {
	ElectronicCommerceIndicator apijson.Field
	PointOfServiceEntryMode     apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisa) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator string

const (
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorMailPhoneOrder                                          CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "mail_phone_order"
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorRecurring                                               CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "recurring"
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorInstallment                                             CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "installment"
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorUnknownMailPhoneOrder                                   CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "unknown_mail_phone_order"
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorSecureElectronicCommerce                                CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "secure_electronic_commerce"
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransactionAt3DSCapableMerchant CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction_at_3ds_capable_merchant"
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransaction                     CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction"
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonSecureTransaction                                    CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "non_secure_transaction"
)

type CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrency string

const (
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrencyCad CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrency = "CAD"
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrencyChf CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrency = "CHF"
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrencyEur CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrency = "EUR"
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrencyGbp CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrency = "GBP"
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrencyJpy CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrency = "JPY"
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrencyUsd CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrency = "USD"
)

type CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationType string

const (
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationTypeCardAuthorization CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationType = "card_authorization"
)

// A Check Deposit Instruction object. This field will be present in the JSON
// response if and only if `category` is equal to `check_deposit_instruction`.
type CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrency `json:"currency,required"`
	// The identifier of the File containing the image of the front of the check that
	// was deposited.
	FrontImageFileID string `json:"front_image_file_id,required"`
	// The identifier of the File containing the image of the back of the check that
	// was deposited.
	BackImageFileID string `json:"back_image_file_id,required,nullable"`
	// The identifier of the Check Deposit.
	CheckDepositID string `json:"check_deposit_id,required,nullable"`
	JSON           cardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionJSON
}

// cardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionJSON
// contains the JSON metadata for the struct
// [CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstruction]
type cardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionJSON struct {
	Amount           apijson.Field
	Currency         apijson.Field
	FrontImageFileID apijson.Field
	BackImageFileID  apijson.Field
	CheckDepositID   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstruction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrency string

const (
	CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrencyCad CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrency = "CAD"
	CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrencyChf CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrency = "CHF"
	CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrencyEur CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrency = "EUR"
	CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrencyGbp CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrency = "GBP"
	CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrencyJpy CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrency = "JPY"
	CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrencyUsd CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrency = "USD"
)

// A Check Transfer Instruction object. This field will be present in the JSON
// response if and only if `category` is equal to `check_transfer_instruction`.
type CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
	// currency.
	Currency CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrency `json:"currency,required"`
	// The identifier of the Check Transfer that led to this Pending Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       cardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionJSON
}

// cardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionJSON
// contains the JSON metadata for the struct
// [CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstruction]
type cardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	TransferID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstruction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrency string

const (
	CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrencyCad CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrency = "CAD"
	CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrencyChf CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrency = "CHF"
	CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrencyEur CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrency = "EUR"
	CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrencyGbp CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrency = "GBP"
	CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrencyJpy CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrency = "JPY"
	CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrencyUsd CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrency = "USD"
)

// A Inbound Funds Hold object. This field will be present in the JSON response if
// and only if `category` is equal to `inbound_funds_hold`.
type CardAuthorizationSimulationPendingTransactionSourceInboundFundsHold struct {
	// The held amount in the minor unit of the account's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the hold
	// was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the hold's
	// currency.
	Currency CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrency `json:"currency,required"`
	// When the hold will be released automatically. Certain conditions may cause it to
	// be released before this time.
	AutomaticallyReleasesAt time.Time `json:"automatically_releases_at,required" format:"date-time"`
	// When the hold was released (if it has been released).
	ReleasedAt time.Time `json:"released_at,required,nullable" format:"date-time"`
	// The status of the hold.
	Status CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldStatus `json:"status,required"`
	// The ID of the Transaction for which funds were held.
	HeldTransactionID string `json:"held_transaction_id,required,nullable"`
	// The ID of the Pending Transaction representing the held funds.
	PendingTransactionID string `json:"pending_transaction_id,required,nullable"`
	JSON                 cardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldJSON
}

// cardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldJSON contains
// the JSON metadata for the struct
// [CardAuthorizationSimulationPendingTransactionSourceInboundFundsHold]
type cardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldJSON struct {
	Amount                  apijson.Field
	CreatedAt               apijson.Field
	Currency                apijson.Field
	AutomaticallyReleasesAt apijson.Field
	ReleasedAt              apijson.Field
	Status                  apijson.Field
	HeldTransactionID       apijson.Field
	PendingTransactionID    apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *CardAuthorizationSimulationPendingTransactionSourceInboundFundsHold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrency string

const (
	CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrencyCad CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrency = "CAD"
	CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrencyChf CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrency = "CHF"
	CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrencyEur CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrency = "EUR"
	CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrencyGbp CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrency = "GBP"
	CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrencyJpy CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrency = "JPY"
	CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrencyUsd CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrency = "USD"
)

type CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldStatus string

const (
	CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldStatusHeld     CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldStatus = "held"
	CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldStatusComplete CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldStatus = "complete"
)

// A Deprecated Card Authorization object. This field will be present in the JSON
// response if and only if `category` is equal to `card_route_authorization`.
type CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorization struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency             CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationCurrency `json:"currency,required"`
	MerchantAcceptorID   string                                                                            `json:"merchant_acceptor_id,required"`
	MerchantCity         string                                                                            `json:"merchant_city,required,nullable"`
	MerchantCountry      string                                                                            `json:"merchant_country,required"`
	MerchantDescriptor   string                                                                            `json:"merchant_descriptor,required"`
	MerchantCategoryCode string                                                                            `json:"merchant_category_code,required"`
	MerchantState        string                                                                            `json:"merchant_state,required,nullable"`
	JSON                 cardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationJSON
}

// cardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationJSON
// contains the JSON metadata for the struct
// [CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorization]
type cardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationJSON struct {
	Amount               apijson.Field
	Currency             apijson.Field
	MerchantAcceptorID   apijson.Field
	MerchantCity         apijson.Field
	MerchantCountry      apijson.Field
	MerchantDescriptor   apijson.Field
	MerchantCategoryCode apijson.Field
	MerchantState        apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationCurrency string

const (
	CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationCurrencyCad CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationCurrency = "CAD"
	CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationCurrencyChf CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationCurrency = "CHF"
	CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationCurrencyEur CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationCurrency = "EUR"
	CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationCurrencyGbp CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationCurrency = "GBP"
	CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationCurrencyJpy CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationCurrency = "JPY"
	CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationCurrencyUsd CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationCurrency = "USD"
)

// A Real Time Payments Transfer Instruction object. This field will be present in
// the JSON response if and only if `category` is equal to
// `real_time_payments_transfer_instruction`.
type CardAuthorizationSimulationPendingTransactionSourceRealTimePaymentsTransferInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The identifier of the Real Time Payments Transfer that led to this Pending
	// Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       cardAuthorizationSimulationPendingTransactionSourceRealTimePaymentsTransferInstructionJSON
}

// cardAuthorizationSimulationPendingTransactionSourceRealTimePaymentsTransferInstructionJSON
// contains the JSON metadata for the struct
// [CardAuthorizationSimulationPendingTransactionSourceRealTimePaymentsTransferInstruction]
type cardAuthorizationSimulationPendingTransactionSourceRealTimePaymentsTransferInstructionJSON struct {
	Amount      apijson.Field
	TransferID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationSimulationPendingTransactionSourceRealTimePaymentsTransferInstruction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Wire Drawdown Payment Instruction object. This field will be present in the
// JSON response if and only if `category` is equal to
// `wire_drawdown_payment_instruction`.
type CardAuthorizationSimulationPendingTransactionSourceWireDrawdownPaymentInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount             int64  `json:"amount,required"`
	AccountNumber      string `json:"account_number,required"`
	RoutingNumber      string `json:"routing_number,required"`
	MessageToRecipient string `json:"message_to_recipient,required"`
	JSON               cardAuthorizationSimulationPendingTransactionSourceWireDrawdownPaymentInstructionJSON
}

// cardAuthorizationSimulationPendingTransactionSourceWireDrawdownPaymentInstructionJSON
// contains the JSON metadata for the struct
// [CardAuthorizationSimulationPendingTransactionSourceWireDrawdownPaymentInstruction]
type cardAuthorizationSimulationPendingTransactionSourceWireDrawdownPaymentInstructionJSON struct {
	Amount             apijson.Field
	AccountNumber      apijson.Field
	RoutingNumber      apijson.Field
	MessageToRecipient apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CardAuthorizationSimulationPendingTransactionSourceWireDrawdownPaymentInstruction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Wire Transfer Instruction object. This field will be present in the JSON
// response if and only if `category` is equal to `wire_transfer_instruction`.
type CardAuthorizationSimulationPendingTransactionSourceWireTransferInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount             int64  `json:"amount,required"`
	AccountNumber      string `json:"account_number,required"`
	RoutingNumber      string `json:"routing_number,required"`
	MessageToRecipient string `json:"message_to_recipient,required"`
	TransferID         string `json:"transfer_id,required"`
	JSON               cardAuthorizationSimulationPendingTransactionSourceWireTransferInstructionJSON
}

// cardAuthorizationSimulationPendingTransactionSourceWireTransferInstructionJSON
// contains the JSON metadata for the struct
// [CardAuthorizationSimulationPendingTransactionSourceWireTransferInstruction]
type cardAuthorizationSimulationPendingTransactionSourceWireTransferInstructionJSON struct {
	Amount             apijson.Field
	AccountNumber      apijson.Field
	RoutingNumber      apijson.Field
	MessageToRecipient apijson.Field
	TransferID         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CardAuthorizationSimulationPendingTransactionSourceWireTransferInstruction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CardAuthorizationSimulationPendingTransactionStatus string

const (
	CardAuthorizationSimulationPendingTransactionStatusPending  CardAuthorizationSimulationPendingTransactionStatus = "pending"
	CardAuthorizationSimulationPendingTransactionStatusComplete CardAuthorizationSimulationPendingTransactionStatus = "complete"
)

type CardAuthorizationSimulationPendingTransactionType string

const (
	CardAuthorizationSimulationPendingTransactionTypePendingTransaction CardAuthorizationSimulationPendingTransactionType = "pending_transaction"
)

// If the authorization attempt fails, this will contain the resulting
// [Declined Transaction](#declined-transactions) object. The Declined
// Transaction's `source` will be of `category: card_decline`.
type CardAuthorizationSimulationDeclinedTransaction struct {
	// The identifier for the Account the Declined Transaction belongs to.
	AccountID string `json:"account_id,required"`
	// The Declined Transaction amount in the minor unit of its currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Declined
	// Transaction's currency. This will match the currency on the Declined
	// Transcation's Account.
	Currency CardAuthorizationSimulationDeclinedTransactionCurrency `json:"currency,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the
	// Transaction occured.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// This is the description the vendor provides.
	Description string `json:"description,required"`
	// The Declined Transaction identifier.
	ID string `json:"id,required"`
	// The identifier for the route this Declined Transaction came through. Routes are
	// things like cards and ACH details.
	RouteID string `json:"route_id,required,nullable"`
	// The type of the route this Declined Transaction came through.
	RouteType CardAuthorizationSimulationDeclinedTransactionRouteType `json:"route_type,required,nullable"`
	// This is an object giving more details on the network-level event that caused the
	// Declined Transaction. For example, for a card transaction this lists the
	// merchant's industry and location. Note that for backwards compatibility reasons,
	// additional undocumented keys may appear in this object. These should be treated
	// as deprecated and will be removed in the future.
	Source CardAuthorizationSimulationDeclinedTransactionSource `json:"source,required"`
	// A constant representing the object's type. For this resource it will always be
	// `declined_transaction`.
	Type CardAuthorizationSimulationDeclinedTransactionType `json:"type,required"`
	JSON cardAuthorizationSimulationDeclinedTransactionJSON
}

// cardAuthorizationSimulationDeclinedTransactionJSON contains the JSON metadata
// for the struct [CardAuthorizationSimulationDeclinedTransaction]
type cardAuthorizationSimulationDeclinedTransactionJSON struct {
	AccountID   apijson.Field
	Amount      apijson.Field
	Currency    apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	ID          apijson.Field
	RouteID     apijson.Field
	RouteType   apijson.Field
	Source      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationSimulationDeclinedTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CardAuthorizationSimulationDeclinedTransactionCurrency string

const (
	CardAuthorizationSimulationDeclinedTransactionCurrencyCad CardAuthorizationSimulationDeclinedTransactionCurrency = "CAD"
	CardAuthorizationSimulationDeclinedTransactionCurrencyChf CardAuthorizationSimulationDeclinedTransactionCurrency = "CHF"
	CardAuthorizationSimulationDeclinedTransactionCurrencyEur CardAuthorizationSimulationDeclinedTransactionCurrency = "EUR"
	CardAuthorizationSimulationDeclinedTransactionCurrencyGbp CardAuthorizationSimulationDeclinedTransactionCurrency = "GBP"
	CardAuthorizationSimulationDeclinedTransactionCurrencyJpy CardAuthorizationSimulationDeclinedTransactionCurrency = "JPY"
	CardAuthorizationSimulationDeclinedTransactionCurrencyUsd CardAuthorizationSimulationDeclinedTransactionCurrency = "USD"
)

type CardAuthorizationSimulationDeclinedTransactionRouteType string

const (
	CardAuthorizationSimulationDeclinedTransactionRouteTypeAccountNumber CardAuthorizationSimulationDeclinedTransactionRouteType = "account_number"
	CardAuthorizationSimulationDeclinedTransactionRouteTypeCard          CardAuthorizationSimulationDeclinedTransactionRouteType = "card"
)

// This is an object giving more details on the network-level event that caused the
// Declined Transaction. For example, for a card transaction this lists the
// merchant's industry and location. Note that for backwards compatibility reasons,
// additional undocumented keys may appear in this object. These should be treated
// as deprecated and will be removed in the future.
type CardAuthorizationSimulationDeclinedTransactionSource struct {
	// The type of decline that took place. We may add additional possible values for
	// this enum over time; your application should be able to handle such additions
	// gracefully.
	Category CardAuthorizationSimulationDeclinedTransactionSourceCategory `json:"category,required"`
	// A ACH Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `ach_decline`.
	ACHDecline CardAuthorizationSimulationDeclinedTransactionSourceACHDecline `json:"ach_decline,required,nullable"`
	// A Card Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_decline`.
	CardDecline CardAuthorizationSimulationDeclinedTransactionSourceCardDecline `json:"card_decline,required,nullable"`
	// A Check Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `check_decline`.
	CheckDecline CardAuthorizationSimulationDeclinedTransactionSourceCheckDecline `json:"check_decline,required,nullable"`
	// A Inbound Real Time Payments Transfer Decline object. This field will be present
	// in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_decline`.
	InboundRealTimePaymentsTransferDecline CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline `json:"inbound_real_time_payments_transfer_decline,required,nullable"`
	// A International ACH Decline object. This field will be present in the JSON
	// response if and only if `category` is equal to `international_ach_decline`.
	InternationalACHDecline CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline `json:"international_ach_decline,required,nullable"`
	// A Deprecated Card Decline object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_route_decline`.
	CardRouteDecline CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDecline `json:"card_route_decline,required,nullable"`
	// A Wire Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `wire_decline`.
	WireDecline CardAuthorizationSimulationDeclinedTransactionSourceWireDecline `json:"wire_decline,required,nullable"`
	JSON        cardAuthorizationSimulationDeclinedTransactionSourceJSON
}

// cardAuthorizationSimulationDeclinedTransactionSourceJSON contains the JSON
// metadata for the struct [CardAuthorizationSimulationDeclinedTransactionSource]
type cardAuthorizationSimulationDeclinedTransactionSourceJSON struct {
	Category                               apijson.Field
	ACHDecline                             apijson.Field
	CardDecline                            apijson.Field
	CheckDecline                           apijson.Field
	InboundRealTimePaymentsTransferDecline apijson.Field
	InternationalACHDecline                apijson.Field
	CardRouteDecline                       apijson.Field
	WireDecline                            apijson.Field
	raw                                    string
	ExtraFields                            map[string]apijson.Field
}

func (r *CardAuthorizationSimulationDeclinedTransactionSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CardAuthorizationSimulationDeclinedTransactionSourceCategory string

const (
	CardAuthorizationSimulationDeclinedTransactionSourceCategoryACHDecline                             CardAuthorizationSimulationDeclinedTransactionSourceCategory = "ach_decline"
	CardAuthorizationSimulationDeclinedTransactionSourceCategoryCardDecline                            CardAuthorizationSimulationDeclinedTransactionSourceCategory = "card_decline"
	CardAuthorizationSimulationDeclinedTransactionSourceCategoryCheckDecline                           CardAuthorizationSimulationDeclinedTransactionSourceCategory = "check_decline"
	CardAuthorizationSimulationDeclinedTransactionSourceCategoryInboundRealTimePaymentsTransferDecline CardAuthorizationSimulationDeclinedTransactionSourceCategory = "inbound_real_time_payments_transfer_decline"
	CardAuthorizationSimulationDeclinedTransactionSourceCategoryInternationalACHDecline                CardAuthorizationSimulationDeclinedTransactionSourceCategory = "international_ach_decline"
	CardAuthorizationSimulationDeclinedTransactionSourceCategoryCardRouteDecline                       CardAuthorizationSimulationDeclinedTransactionSourceCategory = "card_route_decline"
	CardAuthorizationSimulationDeclinedTransactionSourceCategoryWireDecline                            CardAuthorizationSimulationDeclinedTransactionSourceCategory = "wire_decline"
	CardAuthorizationSimulationDeclinedTransactionSourceCategoryOther                                  CardAuthorizationSimulationDeclinedTransactionSourceCategory = "other"
)

// A ACH Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `ach_decline`.
type CardAuthorizationSimulationDeclinedTransactionSourceACHDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount                             int64  `json:"amount,required"`
	OriginatorCompanyName              string `json:"originator_company_name,required"`
	OriginatorCompanyDescriptiveDate   string `json:"originator_company_descriptive_date,required,nullable"`
	OriginatorCompanyDiscretionaryData string `json:"originator_company_discretionary_data,required,nullable"`
	OriginatorCompanyID                string `json:"originator_company_id,required"`
	// Why the ACH transfer was declined.
	Reason           CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason `json:"reason,required"`
	ReceiverIDNumber string                                                               `json:"receiver_id_number,required,nullable"`
	ReceiverName     string                                                               `json:"receiver_name,required,nullable"`
	TraceNumber      string                                                               `json:"trace_number,required"`
	JSON             cardAuthorizationSimulationDeclinedTransactionSourceACHDeclineJSON
}

// cardAuthorizationSimulationDeclinedTransactionSourceACHDeclineJSON contains the
// JSON metadata for the struct
// [CardAuthorizationSimulationDeclinedTransactionSourceACHDecline]
type cardAuthorizationSimulationDeclinedTransactionSourceACHDeclineJSON struct {
	Amount                             apijson.Field
	OriginatorCompanyName              apijson.Field
	OriginatorCompanyDescriptiveDate   apijson.Field
	OriginatorCompanyDiscretionaryData apijson.Field
	OriginatorCompanyID                apijson.Field
	Reason                             apijson.Field
	ReceiverIDNumber                   apijson.Field
	ReceiverName                       apijson.Field
	TraceNumber                        apijson.Field
	raw                                string
	ExtraFields                        map[string]apijson.Field
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceACHDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason string

const (
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonACHRouteCanceled             CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "ach_route_canceled"
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonACHRouteDisabled             CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "ach_route_disabled"
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonBreachesLimit                CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "breaches_limit"
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonCreditEntryRefusedByReceiver CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "credit_entry_refused_by_receiver"
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonDuplicateReturn              CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "duplicate_return"
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonEntityNotActive              CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "entity_not_active"
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonGroupLocked                  CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "group_locked"
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonInsufficientFunds            CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "insufficient_funds"
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonMisroutedReturn              CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "misrouted_return"
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonNoACHRoute                   CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "no_ach_route"
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonOriginatorRequest            CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "originator_request"
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonTransactionNotAllowed        CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "transaction_not_allowed"
)

// A Card Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `card_decline`.
type CardAuthorizationSimulationDeclinedTransactionSourceCardDecline struct {
	// The merchant identifier (commonly abbreviated as MID) of the merchant the card
	// is transacting with.
	MerchantAcceptorID string `json:"merchant_acceptor_id,required"`
	// The merchant descriptor of the merchant the card is transacting with.
	MerchantDescriptor string `json:"merchant_descriptor,required"`
	// The Merchant Category Code (commonly abbreviated as MCC) of the merchant the
	// card is transacting with.
	MerchantCategoryCode string `json:"merchant_category_code,required,nullable"`
	// The city the merchant resides in.
	MerchantCity string `json:"merchant_city,required,nullable"`
	// The country the merchant resides in.
	MerchantCountry string `json:"merchant_country,required,nullable"`
	// The payment network used to process this card authorization
	Network CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetwork `json:"network,required"`
	// Fields specific to the `network`
	NetworkDetails CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetails `json:"network_details,required"`
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrency `json:"currency,required"`
	// Why the transaction was declined.
	Reason CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason `json:"reason,required"`
	// The state the merchant resides in.
	MerchantState string `json:"merchant_state,required,nullable"`
	// The identifier of the Real-Time Decision sent to approve or decline this
	// transaction.
	RealTimeDecisionID string `json:"real_time_decision_id,required,nullable"`
	// If the authorization was attempted using a Digital Wallet Token (such as an
	// Apple Pay purchase), the identifier of the token that was used.
	DigitalWalletTokenID string `json:"digital_wallet_token_id,required,nullable"`
	JSON                 cardAuthorizationSimulationDeclinedTransactionSourceCardDeclineJSON
}

// cardAuthorizationSimulationDeclinedTransactionSourceCardDeclineJSON contains the
// JSON metadata for the struct
// [CardAuthorizationSimulationDeclinedTransactionSourceCardDecline]
type cardAuthorizationSimulationDeclinedTransactionSourceCardDeclineJSON struct {
	MerchantAcceptorID   apijson.Field
	MerchantDescriptor   apijson.Field
	MerchantCategoryCode apijson.Field
	MerchantCity         apijson.Field
	MerchantCountry      apijson.Field
	Network              apijson.Field
	NetworkDetails       apijson.Field
	Amount               apijson.Field
	Currency             apijson.Field
	Reason               apijson.Field
	MerchantState        apijson.Field
	RealTimeDecisionID   apijson.Field
	DigitalWalletTokenID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetwork string

const (
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkVisa CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetwork = "visa"
)

// Fields specific to the `network`
type CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetails struct {
	// Fields specific to the `visa` network
	Visa CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa `json:"visa,required"`
	JSON cardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsJSON
}

// cardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsJSON
// contains the JSON metadata for the struct
// [CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetails]
type cardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsJSON struct {
	Visa        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Fields specific to the `visa` network
type CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa struct {
	// For electronic commerce transactions, this identifies the level of security used
	// in obtaining the customer's payment credential. For mail or telephone order
	// transactions, identifies the type of mail or telephone order.
	ElectronicCommerceIndicator CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator `json:"electronic_commerce_indicator,required,nullable"`
	// The method used to enter the cardholder's primary account number and card
	// expiration date
	PointOfServiceEntryMode shared.PointOfServiceEntryMode `json:"point_of_service_entry_mode,required,nullable"`
	JSON                    cardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaJSON
}

// cardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaJSON
// contains the JSON metadata for the struct
// [CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa]
type cardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaJSON struct {
	ElectronicCommerceIndicator apijson.Field
	PointOfServiceEntryMode     apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator string

const (
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorMailPhoneOrder                                          CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "mail_phone_order"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorRecurring                                               CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "recurring"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorInstallment                                             CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "installment"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorUnknownMailPhoneOrder                                   CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "unknown_mail_phone_order"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorSecureElectronicCommerce                                CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "secure_electronic_commerce"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransactionAt3DSCapableMerchant CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction_at_3ds_capable_merchant"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransaction                     CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonSecureTransaction                                    CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "non_secure_transaction"
)

type CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrency string

const (
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrencyCad CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrency = "CAD"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrencyChf CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrency = "CHF"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrencyEur CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrency = "EUR"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrencyGbp CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrency = "GBP"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrencyJpy CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrency = "JPY"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrencyUsd CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrency = "USD"
)

type CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason string

const (
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonCardNotActive                CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "card_not_active"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonEntityNotActive              CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "entity_not_active"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonGroupLocked                  CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "group_locked"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonInsufficientFunds            CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "insufficient_funds"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonCvv2Mismatch                 CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "cvv2_mismatch"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonTransactionNotAllowed        CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "transaction_not_allowed"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonBreachesInternalLimit        CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "breaches_internal_limit"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonBreachesLimit                CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "breaches_limit"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonWebhookDeclined              CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "webhook_declined"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonWebhookTimedOut              CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "webhook_timed_out"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonDeclinedByStandInProcessing  CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "declined_by_stand_in_processing"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonInvalidPhysicalCard          CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "invalid_physical_card"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonMissingOriginalAuthorization CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "missing_original_authorization"
)

// A Check Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `check_decline`.
type CardAuthorizationSimulationDeclinedTransactionSourceCheckDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount        int64  `json:"amount,required"`
	AuxiliaryOnUs string `json:"auxiliary_on_us,required,nullable"`
	// Why the check was declined.
	Reason CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason `json:"reason,required"`
	JSON   cardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineJSON
}

// cardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineJSON contains
// the JSON metadata for the struct
// [CardAuthorizationSimulationDeclinedTransactionSourceCheckDecline]
type cardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineJSON struct {
	Amount        apijson.Field
	AuxiliaryOnUs apijson.Field
	Reason        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceCheckDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason string

const (
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonACHRouteCanceled      CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "ach_route_canceled"
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonACHRouteDisabled      CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "ach_route_disabled"
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonBreachesLimit         CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "breaches_limit"
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonEntityNotActive       CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "entity_not_active"
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonGroupLocked           CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "group_locked"
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonInsufficientFunds     CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "insufficient_funds"
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonUnableToLocateAccount CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "unable_to_locate_account"
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonUnableToProcess       CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "unable_to_process"
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonReferToImage          CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "refer_to_image"
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonStopPaymentRequested  CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "stop_payment_requested"
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonReturned              CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "returned"
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonDuplicatePresentment  CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "duplicate_presentment"
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonNotAuthorized         CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "not_authorized"
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonAlteredOrFictitious   CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "altered_or_fictitious"
)

// A Inbound Real Time Payments Transfer Decline object. This field will be present
// in the JSON response if and only if `category` is equal to
// `inbound_real_time_payments_transfer_decline`.
type CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the declined
	// transfer's currency. This will always be "USD" for a Real Time Payments
	// transfer.
	Currency CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency `json:"currency,required"`
	// Why the transfer was declined.
	Reason CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason `json:"reason,required"`
	// The name the sender of the transfer specified as the recipient of the transfer.
	CreditorName string `json:"creditor_name,required"`
	// The name provided by the sender of the transfer.
	DebtorName string `json:"debtor_name,required"`
	// The account number of the account that sent the transfer.
	DebtorAccountNumber string `json:"debtor_account_number,required"`
	// The routing number of the account that sent the transfer.
	DebtorRoutingNumber string `json:"debtor_routing_number,required"`
	// The Real Time Payments network identification of the declined transfer.
	TransactionIdentification string `json:"transaction_identification,required"`
	// Additional information included with the transfer.
	RemittanceInformation string `json:"remittance_information,required,nullable"`
	JSON                  cardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineJSON
}

// cardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineJSON
// contains the JSON metadata for the struct
// [CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline]
type cardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineJSON struct {
	Amount                    apijson.Field
	Currency                  apijson.Field
	Reason                    apijson.Field
	CreditorName              apijson.Field
	DebtorName                apijson.Field
	DebtorAccountNumber       apijson.Field
	DebtorRoutingNumber       apijson.Field
	TransactionIdentification apijson.Field
	RemittanceInformation     apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency string

const (
	CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyCad CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "CAD"
	CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyChf CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "CHF"
	CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyEur CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "EUR"
	CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyGbp CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "GBP"
	CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyJpy CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "JPY"
	CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyUsd CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "USD"
)

type CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason string

const (
	CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonAccountNumberCanceled      CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "account_number_canceled"
	CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonAccountNumberDisabled      CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "account_number_disabled"
	CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonAccountRestricted          CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "account_restricted"
	CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonGroupLocked                CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "group_locked"
	CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonEntityNotActive            CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "entity_not_active"
	CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonRealTimePaymentsNotEnabled CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "real_time_payments_not_enabled"
)

// A International ACH Decline object. This field will be present in the JSON
// response if and only if `category` is equal to `international_ach_decline`.
type CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount                                                 int64  `json:"amount,required"`
	ForeignExchangeIndicator                               string `json:"foreign_exchange_indicator,required"`
	ForeignExchangeReferenceIndicator                      string `json:"foreign_exchange_reference_indicator,required"`
	ForeignExchangeReference                               string `json:"foreign_exchange_reference,required,nullable"`
	DestinationCountryCode                                 string `json:"destination_country_code,required"`
	DestinationCurrencyCode                                string `json:"destination_currency_code,required"`
	ForeignPaymentAmount                                   int64  `json:"foreign_payment_amount,required"`
	ForeignTraceNumber                                     string `json:"foreign_trace_number,required,nullable"`
	InternationalTransactionTypeCode                       string `json:"international_transaction_type_code,required"`
	OriginatingCurrencyCode                                string `json:"originating_currency_code,required"`
	OriginatingDepositoryFinancialInstitutionName          string `json:"originating_depository_financial_institution_name,required"`
	OriginatingDepositoryFinancialInstitutionIDQualifier   string `json:"originating_depository_financial_institution_id_qualifier,required"`
	OriginatingDepositoryFinancialInstitutionID            string `json:"originating_depository_financial_institution_id,required"`
	OriginatingDepositoryFinancialInstitutionBranchCountry string `json:"originating_depository_financial_institution_branch_country,required"`
	OriginatorCity                                         string `json:"originator_city,required"`
	OriginatorCompanyEntryDescription                      string `json:"originator_company_entry_description,required"`
	OriginatorCountry                                      string `json:"originator_country,required"`
	OriginatorIdentification                               string `json:"originator_identification,required"`
	OriginatorName                                         string `json:"originator_name,required"`
	OriginatorPostalCode                                   string `json:"originator_postal_code,required,nullable"`
	OriginatorStreetAddress                                string `json:"originator_street_address,required"`
	OriginatorStateOrProvince                              string `json:"originator_state_or_province,required,nullable"`
	PaymentRelatedInformation                              string `json:"payment_related_information,required,nullable"`
	PaymentRelatedInformation2                             string `json:"payment_related_information2,required,nullable"`
	ReceiverIdentificationNumber                           string `json:"receiver_identification_number,required,nullable"`
	ReceiverStreetAddress                                  string `json:"receiver_street_address,required"`
	ReceiverCity                                           string `json:"receiver_city,required"`
	ReceiverStateOrProvince                                string `json:"receiver_state_or_province,required,nullable"`
	ReceiverCountry                                        string `json:"receiver_country,required"`
	ReceiverPostalCode                                     string `json:"receiver_postal_code,required,nullable"`
	ReceivingCompanyOrIndividualName                       string `json:"receiving_company_or_individual_name,required"`
	ReceivingDepositoryFinancialInstitutionName            string `json:"receiving_depository_financial_institution_name,required"`
	ReceivingDepositoryFinancialInstitutionIDQualifier     string `json:"receiving_depository_financial_institution_id_qualifier,required"`
	ReceivingDepositoryFinancialInstitutionID              string `json:"receiving_depository_financial_institution_id,required"`
	ReceivingDepositoryFinancialInstitutionCountry         string `json:"receiving_depository_financial_institution_country,required"`
	TraceNumber                                            string `json:"trace_number,required"`
	JSON                                                   cardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineJSON
}

// cardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineJSON
// contains the JSON metadata for the struct
// [CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline]
type cardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineJSON struct {
	Amount                                                 apijson.Field
	ForeignExchangeIndicator                               apijson.Field
	ForeignExchangeReferenceIndicator                      apijson.Field
	ForeignExchangeReference                               apijson.Field
	DestinationCountryCode                                 apijson.Field
	DestinationCurrencyCode                                apijson.Field
	ForeignPaymentAmount                                   apijson.Field
	ForeignTraceNumber                                     apijson.Field
	InternationalTransactionTypeCode                       apijson.Field
	OriginatingCurrencyCode                                apijson.Field
	OriginatingDepositoryFinancialInstitutionName          apijson.Field
	OriginatingDepositoryFinancialInstitutionIDQualifier   apijson.Field
	OriginatingDepositoryFinancialInstitutionID            apijson.Field
	OriginatingDepositoryFinancialInstitutionBranchCountry apijson.Field
	OriginatorCity                                         apijson.Field
	OriginatorCompanyEntryDescription                      apijson.Field
	OriginatorCountry                                      apijson.Field
	OriginatorIdentification                               apijson.Field
	OriginatorName                                         apijson.Field
	OriginatorPostalCode                                   apijson.Field
	OriginatorStreetAddress                                apijson.Field
	OriginatorStateOrProvince                              apijson.Field
	PaymentRelatedInformation                              apijson.Field
	PaymentRelatedInformation2                             apijson.Field
	ReceiverIdentificationNumber                           apijson.Field
	ReceiverStreetAddress                                  apijson.Field
	ReceiverCity                                           apijson.Field
	ReceiverStateOrProvince                                apijson.Field
	ReceiverCountry                                        apijson.Field
	ReceiverPostalCode                                     apijson.Field
	ReceivingCompanyOrIndividualName                       apijson.Field
	ReceivingDepositoryFinancialInstitutionName            apijson.Field
	ReceivingDepositoryFinancialInstitutionIDQualifier     apijson.Field
	ReceivingDepositoryFinancialInstitutionID              apijson.Field
	ReceivingDepositoryFinancialInstitutionCountry         apijson.Field
	TraceNumber                                            apijson.Field
	raw                                                    string
	ExtraFields                                            map[string]apijson.Field
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Deprecated Card Decline object. This field will be present in the JSON
// response if and only if `category` is equal to `card_route_decline`.
type CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency             CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDeclineCurrency `json:"currency,required"`
	MerchantAcceptorID   string                                                                       `json:"merchant_acceptor_id,required"`
	MerchantCity         string                                                                       `json:"merchant_city,required,nullable"`
	MerchantCountry      string                                                                       `json:"merchant_country,required"`
	MerchantDescriptor   string                                                                       `json:"merchant_descriptor,required"`
	MerchantState        string                                                                       `json:"merchant_state,required,nullable"`
	MerchantCategoryCode string                                                                       `json:"merchant_category_code,required,nullable"`
	JSON                 cardAuthorizationSimulationDeclinedTransactionSourceCardRouteDeclineJSON
}

// cardAuthorizationSimulationDeclinedTransactionSourceCardRouteDeclineJSON
// contains the JSON metadata for the struct
// [CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDecline]
type cardAuthorizationSimulationDeclinedTransactionSourceCardRouteDeclineJSON struct {
	Amount               apijson.Field
	Currency             apijson.Field
	MerchantAcceptorID   apijson.Field
	MerchantCity         apijson.Field
	MerchantCountry      apijson.Field
	MerchantDescriptor   apijson.Field
	MerchantState        apijson.Field
	MerchantCategoryCode apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDeclineCurrency string

const (
	CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDeclineCurrencyCad CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDeclineCurrency = "CAD"
	CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDeclineCurrencyChf CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDeclineCurrency = "CHF"
	CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDeclineCurrencyEur CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDeclineCurrency = "EUR"
	CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDeclineCurrencyGbp CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDeclineCurrency = "GBP"
	CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDeclineCurrencyJpy CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDeclineCurrency = "JPY"
	CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDeclineCurrencyUsd CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDeclineCurrency = "USD"
)

// A Wire Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `wire_decline`.
type CardAuthorizationSimulationDeclinedTransactionSourceWireDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// Why the wire transfer was declined.
	Reason                                  CardAuthorizationSimulationDeclinedTransactionSourceWireDeclineReason `json:"reason,required"`
	Description                             string                                                                `json:"description,required"`
	BeneficiaryAddressLine1                 string                                                                `json:"beneficiary_address_line1,required,nullable"`
	BeneficiaryAddressLine2                 string                                                                `json:"beneficiary_address_line2,required,nullable"`
	BeneficiaryAddressLine3                 string                                                                `json:"beneficiary_address_line3,required,nullable"`
	BeneficiaryName                         string                                                                `json:"beneficiary_name,required,nullable"`
	BeneficiaryReference                    string                                                                `json:"beneficiary_reference,required,nullable"`
	InputMessageAccountabilityData          string                                                                `json:"input_message_accountability_data,required,nullable"`
	OriginatorAddressLine1                  string                                                                `json:"originator_address_line1,required,nullable"`
	OriginatorAddressLine2                  string                                                                `json:"originator_address_line2,required,nullable"`
	OriginatorAddressLine3                  string                                                                `json:"originator_address_line3,required,nullable"`
	OriginatorName                          string                                                                `json:"originator_name,required,nullable"`
	OriginatorToBeneficiaryInformationLine1 string                                                                `json:"originator_to_beneficiary_information_line1,required,nullable"`
	OriginatorToBeneficiaryInformationLine2 string                                                                `json:"originator_to_beneficiary_information_line2,required,nullable"`
	OriginatorToBeneficiaryInformationLine3 string                                                                `json:"originator_to_beneficiary_information_line3,required,nullable"`
	OriginatorToBeneficiaryInformationLine4 string                                                                `json:"originator_to_beneficiary_information_line4,required,nullable"`
	JSON                                    cardAuthorizationSimulationDeclinedTransactionSourceWireDeclineJSON
}

// cardAuthorizationSimulationDeclinedTransactionSourceWireDeclineJSON contains the
// JSON metadata for the struct
// [CardAuthorizationSimulationDeclinedTransactionSourceWireDecline]
type cardAuthorizationSimulationDeclinedTransactionSourceWireDeclineJSON struct {
	Amount                                  apijson.Field
	Reason                                  apijson.Field
	Description                             apijson.Field
	BeneficiaryAddressLine1                 apijson.Field
	BeneficiaryAddressLine2                 apijson.Field
	BeneficiaryAddressLine3                 apijson.Field
	BeneficiaryName                         apijson.Field
	BeneficiaryReference                    apijson.Field
	InputMessageAccountabilityData          apijson.Field
	OriginatorAddressLine1                  apijson.Field
	OriginatorAddressLine2                  apijson.Field
	OriginatorAddressLine3                  apijson.Field
	OriginatorName                          apijson.Field
	OriginatorToBeneficiaryInformationLine1 apijson.Field
	OriginatorToBeneficiaryInformationLine2 apijson.Field
	OriginatorToBeneficiaryInformationLine3 apijson.Field
	OriginatorToBeneficiaryInformationLine4 apijson.Field
	raw                                     string
	ExtraFields                             map[string]apijson.Field
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceWireDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CardAuthorizationSimulationDeclinedTransactionSourceWireDeclineReason string

const (
	CardAuthorizationSimulationDeclinedTransactionSourceWireDeclineReasonAccountNumberCanceled CardAuthorizationSimulationDeclinedTransactionSourceWireDeclineReason = "account_number_canceled"
	CardAuthorizationSimulationDeclinedTransactionSourceWireDeclineReasonAccountNumberDisabled CardAuthorizationSimulationDeclinedTransactionSourceWireDeclineReason = "account_number_disabled"
	CardAuthorizationSimulationDeclinedTransactionSourceWireDeclineReasonEntityNotActive       CardAuthorizationSimulationDeclinedTransactionSourceWireDeclineReason = "entity_not_active"
	CardAuthorizationSimulationDeclinedTransactionSourceWireDeclineReasonGroupLocked           CardAuthorizationSimulationDeclinedTransactionSourceWireDeclineReason = "group_locked"
	CardAuthorizationSimulationDeclinedTransactionSourceWireDeclineReasonNoAccountNumber       CardAuthorizationSimulationDeclinedTransactionSourceWireDeclineReason = "no_account_number"
	CardAuthorizationSimulationDeclinedTransactionSourceWireDeclineReasonTransactionNotAllowed CardAuthorizationSimulationDeclinedTransactionSourceWireDeclineReason = "transaction_not_allowed"
)

type CardAuthorizationSimulationDeclinedTransactionType string

const (
	CardAuthorizationSimulationDeclinedTransactionTypeDeclinedTransaction CardAuthorizationSimulationDeclinedTransactionType = "declined_transaction"
)

type CardAuthorizationSimulationType string

const (
	CardAuthorizationSimulationTypeInboundCardAuthorizationSimulationResult CardAuthorizationSimulationType = "inbound_card_authorization_simulation_result"
)

type SimulationCardAuthorizeParams struct {
	// The authorization amount in cents.
	Amount param.Field[int64] `json:"amount,required"`
	// The identifier of the Card to be authorized.
	CardID param.Field[string] `json:"card_id"`
	// The identifier of the Digital Wallet Token to be authorized.
	DigitalWalletTokenID param.Field[string] `json:"digital_wallet_token_id"`
	// The identifier of the Event Subscription to use. If provided, will override the
	// default real time event subscription. Because you can only create one real time
	// decision event subscription, you can use this field to route events to any
	// specified event subscription for testing purposes.
	EventSubscriptionID param.Field[string] `json:"event_subscription_id"`
}

func (r SimulationCardAuthorizeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SimulationCardSettlementParams struct {
	// The identifier of the Card to create a settlement on.
	CardID param.Field[string] `json:"card_id,required"`
	// The identifier of the Pending Transaction for the Card Authorization you wish to
	// settle.
	PendingTransactionID param.Field[string] `json:"pending_transaction_id,required"`
	// The amount to be settled. This defaults to the amount of the Pending Transaction
	// being settled.
	Amount param.Field[int64] `json:"amount"`
}

func (r SimulationCardSettlementParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

package responses

import (
	"time"

	apijson "github.com/increase/increase-go/core/json"
)

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
	JSON CardAuthorizationSimulationJSON
}

type CardAuthorizationSimulationJSON struct {
	PendingTransaction  apijson.Metadata
	DeclinedTransaction apijson.Metadata
	Type                apijson.Metadata
	Raw                 []byte
	Extras              map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CardAuthorizationSimulation
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *CardAuthorizationSimulation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

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
	JSON CardAuthorizationSimulationPendingTransactionJSON
}

type CardAuthorizationSimulationPendingTransactionJSON struct {
	AccountID   apijson.Metadata
	Amount      apijson.Metadata
	Currency    apijson.Metadata
	CompletedAt apijson.Metadata
	CreatedAt   apijson.Metadata
	Description apijson.Metadata
	ID          apijson.Metadata
	RouteID     apijson.Metadata
	RouteType   apijson.Metadata
	Source      apijson.Metadata
	Status      apijson.Metadata
	Type        apijson.Metadata
	Raw         []byte
	Extras      map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// CardAuthorizationSimulationPendingTransaction using the internal json library.
// Unrecognized fields are stored in the `jsonFields` property.
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
	JSON                    CardAuthorizationSimulationPendingTransactionSourceJSON
}

type CardAuthorizationSimulationPendingTransactionSourceJSON struct {
	Category                            apijson.Metadata
	AccountTransferInstruction          apijson.Metadata
	ACHTransferInstruction              apijson.Metadata
	CardAuthorization                   apijson.Metadata
	CheckDepositInstruction             apijson.Metadata
	CheckTransferInstruction            apijson.Metadata
	InboundFundsHold                    apijson.Metadata
	CardRouteAuthorization              apijson.Metadata
	RealTimePaymentsTransferInstruction apijson.Metadata
	WireDrawdownPaymentInstruction      apijson.Metadata
	WireTransferInstruction             apijson.Metadata
	Raw                                 []byte
	Extras                              map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// CardAuthorizationSimulationPendingTransactionSource using the internal json
// library. Unrecognized fields are stored in the `jsonFields` property.
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

type CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrency `json:"currency,required"`
	// The identifier of the Account Transfer that led to this Pending Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionJSON
}

type CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionJSON struct {
	Amount     apijson.Metadata
	Currency   apijson.Metadata
	TransferID apijson.Metadata
	Raw        []byte
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstruction
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
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

type CardAuthorizationSimulationPendingTransactionSourceACHTransferInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The identifier of the ACH Transfer that led to this Pending Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       CardAuthorizationSimulationPendingTransactionSourceACHTransferInstructionJSON
}

type CardAuthorizationSimulationPendingTransactionSourceACHTransferInstructionJSON struct {
	Amount     apijson.Metadata
	TransferID apijson.Metadata
	Raw        []byte
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// CardAuthorizationSimulationPendingTransactionSourceACHTransferInstruction using
// the internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CardAuthorizationSimulationPendingTransactionSourceACHTransferInstruction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

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
	// The identifier of the Real-Time Decision sent to approve or decline this
	// transaction.
	RealTimeDecisionID string `json:"real_time_decision_id,required,nullable"`
	// If the authorization was made via a Digital Wallet Token (such as an Apple Pay
	// purchase), the identifier of the token that was used.
	DigitalWalletTokenID string `json:"digital_wallet_token_id,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `card_authorization`.
	Type CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationType `json:"type,required"`
	JSON CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationJSON
}

type CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationJSON struct {
	ID                   apijson.Metadata
	MerchantAcceptorID   apijson.Metadata
	MerchantDescriptor   apijson.Metadata
	MerchantCategoryCode apijson.Metadata
	MerchantCity         apijson.Metadata
	MerchantCountry      apijson.Metadata
	Network              apijson.Metadata
	NetworkDetails       apijson.Metadata
	Amount               apijson.Metadata
	Currency             apijson.Metadata
	RealTimeDecisionID   apijson.Metadata
	DigitalWalletTokenID apijson.Metadata
	Type                 apijson.Metadata
	Raw                  []byte
	Extras               map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// CardAuthorizationSimulationPendingTransactionSourceCardAuthorization using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CardAuthorizationSimulationPendingTransactionSourceCardAuthorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetwork string

const (
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkVisa CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetwork = "visa"
)

type CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetails struct {
	// Fields specific to the `visa` network
	Visa CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisa `json:"visa,required"`
	JSON CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsJSON
}

type CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsJSON struct {
	Visa   apijson.Metadata
	Raw    []byte
	Extras map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetails
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisa struct {
	// For electronic commerce transactions, this identifies the level of security used
	// in obtaining the customer's payment credential. For mail or telephone order
	// transactions, identifies the type of mail or telephone order.
	ElectronicCommerceIndicator CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator `json:"electronic_commerce_indicator,required,nullable"`
	// The method used to enter the cardholder's primary account number and card
	// expiration date
	PointOfServiceEntryMode PointOfServiceEntryMode `json:"point_of_service_entry_mode,required,nullable"`
	JSON                    CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaJSON
}

type CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaJSON struct {
	ElectronicCommerceIndicator apijson.Metadata
	PointOfServiceEntryMode     apijson.Metadata
	Raw                         []byte
	Extras                      map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisa
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisa) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator string

const (
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorMailPhoneOrder                                           CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "mail_phone_order"
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorRecurring                                                CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "recurring"
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorInstallment                                              CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "installment"
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorUnknownMailPhoneOrder                                    CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "unknown_mail_phone_order"
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorSecureElectronicCommerce                                 CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "secure_electronic_commerce"
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransactionAt_3dsCapableMerchant CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction_at_3ds_capable_merchant"
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransaction                      CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction"
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonSecureTransaction                                     CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "non_secure_transaction"
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
	JSON           CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionJSON
}

type CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionJSON struct {
	Amount           apijson.Metadata
	Currency         apijson.Metadata
	FrontImageFileID apijson.Metadata
	BackImageFileID  apijson.Metadata
	CheckDepositID   apijson.Metadata
	Raw              []byte
	Extras           map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstruction using
// the internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
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

type CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
	// currency.
	Currency CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrency `json:"currency,required"`
	// The identifier of the Check Transfer that led to this Pending Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionJSON
}

type CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionJSON struct {
	Amount     apijson.Metadata
	Currency   apijson.Metadata
	TransferID apijson.Metadata
	Raw        []byte
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstruction
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
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
	JSON                 CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldJSON
}

type CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldJSON struct {
	Amount                  apijson.Metadata
	CreatedAt               apijson.Metadata
	Currency                apijson.Metadata
	AutomaticallyReleasesAt apijson.Metadata
	ReleasedAt              apijson.Metadata
	Status                  apijson.Metadata
	HeldTransactionID       apijson.Metadata
	PendingTransactionID    apijson.Metadata
	Raw                     []byte
	Extras                  map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// CardAuthorizationSimulationPendingTransactionSourceInboundFundsHold using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
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
	JSON                 CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationJSON
}

type CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationJSON struct {
	Amount               apijson.Metadata
	Currency             apijson.Metadata
	MerchantAcceptorID   apijson.Metadata
	MerchantCity         apijson.Metadata
	MerchantCountry      apijson.Metadata
	MerchantDescriptor   apijson.Metadata
	MerchantCategoryCode apijson.Metadata
	MerchantState        apijson.Metadata
	Raw                  []byte
	Extras               map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorization using
// the internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
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

type CardAuthorizationSimulationPendingTransactionSourceRealTimePaymentsTransferInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The identifier of the Real Time Payments Transfer that led to this Pending
	// Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       CardAuthorizationSimulationPendingTransactionSourceRealTimePaymentsTransferInstructionJSON
}

type CardAuthorizationSimulationPendingTransactionSourceRealTimePaymentsTransferInstructionJSON struct {
	Amount     apijson.Metadata
	TransferID apijson.Metadata
	Raw        []byte
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// CardAuthorizationSimulationPendingTransactionSourceRealTimePaymentsTransferInstruction
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *CardAuthorizationSimulationPendingTransactionSourceRealTimePaymentsTransferInstruction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CardAuthorizationSimulationPendingTransactionSourceWireDrawdownPaymentInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount             int64  `json:"amount,required"`
	AccountNumber      string `json:"account_number,required"`
	RoutingNumber      string `json:"routing_number,required"`
	MessageToRecipient string `json:"message_to_recipient,required"`
	JSON               CardAuthorizationSimulationPendingTransactionSourceWireDrawdownPaymentInstructionJSON
}

type CardAuthorizationSimulationPendingTransactionSourceWireDrawdownPaymentInstructionJSON struct {
	Amount             apijson.Metadata
	AccountNumber      apijson.Metadata
	RoutingNumber      apijson.Metadata
	MessageToRecipient apijson.Metadata
	Raw                []byte
	Extras             map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// CardAuthorizationSimulationPendingTransactionSourceWireDrawdownPaymentInstruction
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *CardAuthorizationSimulationPendingTransactionSourceWireDrawdownPaymentInstruction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CardAuthorizationSimulationPendingTransactionSourceWireTransferInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount             int64  `json:"amount,required"`
	AccountNumber      string `json:"account_number,required"`
	RoutingNumber      string `json:"routing_number,required"`
	MessageToRecipient string `json:"message_to_recipient,required"`
	TransferID         string `json:"transfer_id,required"`
	JSON               CardAuthorizationSimulationPendingTransactionSourceWireTransferInstructionJSON
}

type CardAuthorizationSimulationPendingTransactionSourceWireTransferInstructionJSON struct {
	Amount             apijson.Metadata
	AccountNumber      apijson.Metadata
	RoutingNumber      apijson.Metadata
	MessageToRecipient apijson.Metadata
	TransferID         apijson.Metadata
	Raw                []byte
	Extras             map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// CardAuthorizationSimulationPendingTransactionSourceWireTransferInstruction using
// the internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
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
	JSON CardAuthorizationSimulationDeclinedTransactionJSON
}

type CardAuthorizationSimulationDeclinedTransactionJSON struct {
	AccountID   apijson.Metadata
	Amount      apijson.Metadata
	Currency    apijson.Metadata
	CreatedAt   apijson.Metadata
	Description apijson.Metadata
	ID          apijson.Metadata
	RouteID     apijson.Metadata
	RouteType   apijson.Metadata
	Source      apijson.Metadata
	Type        apijson.Metadata
	Raw         []byte
	Extras      map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// CardAuthorizationSimulationDeclinedTransaction using the internal json library.
// Unrecognized fields are stored in the `jsonFields` property.
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
	JSON             CardAuthorizationSimulationDeclinedTransactionSourceJSON
}

type CardAuthorizationSimulationDeclinedTransactionSourceJSON struct {
	Category                               apijson.Metadata
	ACHDecline                             apijson.Metadata
	CardDecline                            apijson.Metadata
	CheckDecline                           apijson.Metadata
	InboundRealTimePaymentsTransferDecline apijson.Metadata
	InternationalACHDecline                apijson.Metadata
	CardRouteDecline                       apijson.Metadata
	Raw                                    []byte
	Extras                                 map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// CardAuthorizationSimulationDeclinedTransactionSource using the internal json
// library. Unrecognized fields are stored in the `jsonFields` property.
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
	CardAuthorizationSimulationDeclinedTransactionSourceCategoryOther                                  CardAuthorizationSimulationDeclinedTransactionSourceCategory = "other"
)

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
	JSON             CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineJSON
}

type CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineJSON struct {
	Amount                             apijson.Metadata
	OriginatorCompanyName              apijson.Metadata
	OriginatorCompanyDescriptiveDate   apijson.Metadata
	OriginatorCompanyDiscretionaryData apijson.Metadata
	OriginatorCompanyID                apijson.Metadata
	Reason                             apijson.Metadata
	ReceiverIDNumber                   apijson.Metadata
	ReceiverName                       apijson.Metadata
	TraceNumber                        apijson.Metadata
	Raw                                []byte
	Extras                             map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// CardAuthorizationSimulationDeclinedTransactionSourceACHDecline using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
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
	JSON                 CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineJSON
}

type CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineJSON struct {
	MerchantAcceptorID   apijson.Metadata
	MerchantDescriptor   apijson.Metadata
	MerchantCategoryCode apijson.Metadata
	MerchantCity         apijson.Metadata
	MerchantCountry      apijson.Metadata
	Network              apijson.Metadata
	NetworkDetails       apijson.Metadata
	Amount               apijson.Metadata
	Currency             apijson.Metadata
	Reason               apijson.Metadata
	MerchantState        apijson.Metadata
	RealTimeDecisionID   apijson.Metadata
	DigitalWalletTokenID apijson.Metadata
	Raw                  []byte
	Extras               map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// CardAuthorizationSimulationDeclinedTransactionSourceCardDecline using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetwork string

const (
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkVisa CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetwork = "visa"
)

type CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetails struct {
	// Fields specific to the `visa` network
	Visa CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa `json:"visa,required"`
	JSON CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsJSON
}

type CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsJSON struct {
	Visa   apijson.Metadata
	Raw    []byte
	Extras map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetails
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa struct {
	// For electronic commerce transactions, this identifies the level of security used
	// in obtaining the customer's payment credential. For mail or telephone order
	// transactions, identifies the type of mail or telephone order.
	ElectronicCommerceIndicator CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator `json:"electronic_commerce_indicator,required,nullable"`
	// The method used to enter the cardholder's primary account number and card
	// expiration date
	PointOfServiceEntryMode PointOfServiceEntryMode `json:"point_of_service_entry_mode,required,nullable"`
	JSON                    CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaJSON
}

type CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaJSON struct {
	ElectronicCommerceIndicator apijson.Metadata
	PointOfServiceEntryMode     apijson.Metadata
	Raw                         []byte
	Extras                      map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator string

const (
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorMailPhoneOrder                                           CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "mail_phone_order"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorRecurring                                                CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "recurring"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorInstallment                                              CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "installment"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorUnknownMailPhoneOrder                                    CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "unknown_mail_phone_order"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorSecureElectronicCommerce                                 CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "secure_electronic_commerce"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransactionAt_3dsCapableMerchant CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction_at_3ds_capable_merchant"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransaction                      CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonSecureTransaction                                     CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "non_secure_transaction"
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

type CardAuthorizationSimulationDeclinedTransactionSourceCheckDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount        int64  `json:"amount,required"`
	AuxiliaryOnUs string `json:"auxiliary_on_us,required,nullable"`
	// Why the check was declined.
	Reason CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason `json:"reason,required"`
	JSON   CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineJSON
}

type CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineJSON struct {
	Amount        apijson.Metadata
	AuxiliaryOnUs apijson.Metadata
	Reason        apijson.Metadata
	Raw           []byte
	Extras        map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// CardAuthorizationSimulationDeclinedTransactionSourceCheckDecline using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
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
)

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
	JSON                  CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineJSON
}

type CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineJSON struct {
	Amount                    apijson.Metadata
	Currency                  apijson.Metadata
	Reason                    apijson.Metadata
	CreditorName              apijson.Metadata
	DebtorName                apijson.Metadata
	DebtorAccountNumber       apijson.Metadata
	DebtorRoutingNumber       apijson.Metadata
	TransactionIdentification apijson.Metadata
	RemittanceInformation     apijson.Metadata
	Raw                       []byte
	Extras                    map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
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
	CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonGroupLocked                CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "group_locked"
	CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonEntityNotActive            CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "entity_not_active"
	CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonRealTimePaymentsNotEnabled CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "real_time_payments_not_enabled"
)

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
	JSON                                                   CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineJSON
}

type CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDeclineJSON struct {
	Amount                                                 apijson.Metadata
	ForeignExchangeIndicator                               apijson.Metadata
	ForeignExchangeReferenceIndicator                      apijson.Metadata
	ForeignExchangeReference                               apijson.Metadata
	DestinationCountryCode                                 apijson.Metadata
	DestinationCurrencyCode                                apijson.Metadata
	ForeignPaymentAmount                                   apijson.Metadata
	ForeignTraceNumber                                     apijson.Metadata
	InternationalTransactionTypeCode                       apijson.Metadata
	OriginatingCurrencyCode                                apijson.Metadata
	OriginatingDepositoryFinancialInstitutionName          apijson.Metadata
	OriginatingDepositoryFinancialInstitutionIDQualifier   apijson.Metadata
	OriginatingDepositoryFinancialInstitutionID            apijson.Metadata
	OriginatingDepositoryFinancialInstitutionBranchCountry apijson.Metadata
	OriginatorCity                                         apijson.Metadata
	OriginatorCompanyEntryDescription                      apijson.Metadata
	OriginatorCountry                                      apijson.Metadata
	OriginatorIdentification                               apijson.Metadata
	OriginatorName                                         apijson.Metadata
	OriginatorPostalCode                                   apijson.Metadata
	OriginatorStreetAddress                                apijson.Metadata
	OriginatorStateOrProvince                              apijson.Metadata
	PaymentRelatedInformation                              apijson.Metadata
	PaymentRelatedInformation2                             apijson.Metadata
	ReceiverIdentificationNumber                           apijson.Metadata
	ReceiverStreetAddress                                  apijson.Metadata
	ReceiverCity                                           apijson.Metadata
	ReceiverStateOrProvince                                apijson.Metadata
	ReceiverCountry                                        apijson.Metadata
	ReceiverPostalCode                                     apijson.Metadata
	ReceivingCompanyOrIndividualName                       apijson.Metadata
	ReceivingDepositoryFinancialInstitutionName            apijson.Metadata
	ReceivingDepositoryFinancialInstitutionIDQualifier     apijson.Metadata
	ReceivingDepositoryFinancialInstitutionID              apijson.Metadata
	ReceivingDepositoryFinancialInstitutionCountry         apijson.Metadata
	TraceNumber                                            apijson.Metadata
	Raw                                                    []byte
	Extras                                                 map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

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
	JSON                 CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDeclineJSON
}

type CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDeclineJSON struct {
	Amount               apijson.Metadata
	Currency             apijson.Metadata
	MerchantAcceptorID   apijson.Metadata
	MerchantCity         apijson.Metadata
	MerchantCountry      apijson.Metadata
	MerchantDescriptor   apijson.Metadata
	MerchantState        apijson.Metadata
	MerchantCategoryCode apijson.Metadata
	Raw                  []byte
	Extras               map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDecline using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
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

type CardAuthorizationSimulationDeclinedTransactionType string

const (
	CardAuthorizationSimulationDeclinedTransactionTypeDeclinedTransaction CardAuthorizationSimulationDeclinedTransactionType = "declined_transaction"
)

type CardAuthorizationSimulationType string

const (
	CardAuthorizationSimulationTypeInboundCardAuthorizationSimulationResult CardAuthorizationSimulationType = "inbound_card_authorization_simulation_result"
)

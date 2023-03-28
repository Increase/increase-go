package requests

import (
	"fmt"
	"time"

	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/fields"
)

type CardAuthorizationSimulation struct {
	// If the authorization attempt succeeds, this will contain the resulting Pending
	// Transaction object. The Pending Transaction's `source` will be of
	// `category: card_authorization`.
	PendingTransaction fields.Field[CardAuthorizationSimulationPendingTransaction] `json:"pending_transaction,required,nullable"`
	// If the authorization attempt fails, this will contain the resulting
	// [Declined Transaction](#declined-transactions) object. The Declined
	// Transaction's `source` will be of `category: card_decline`.
	DeclinedTransaction fields.Field[CardAuthorizationSimulationDeclinedTransaction] `json:"declined_transaction,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `inbound_card_authorization_simulation_result`.
	Type fields.Field[CardAuthorizationSimulationType] `json:"type,required"`
}

// MarshalJSON serializes CardAuthorizationSimulation into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CardAuthorizationSimulation) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CardAuthorizationSimulation) String() (result string) {
	return fmt.Sprintf("&CardAuthorizationSimulation{PendingTransaction:%s DeclinedTransaction:%s Type:%s}", r.PendingTransaction, r.DeclinedTransaction, r.Type)
}

type CardAuthorizationSimulationPendingTransaction struct {
	// The identifier for the account this Pending Transaction belongs to.
	AccountID fields.Field[string] `json:"account_id,required"`
	// The Pending Transaction amount in the minor unit of its currency. For dollars,
	// for example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Pending
	// Transaction's currency. This will match the currency on the Pending
	// Transcation's Account.
	Currency fields.Field[CardAuthorizationSimulationPendingTransactionCurrency] `json:"currency,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the Pending
	// Transaction occured.
	CreatedAt fields.Field[time.Time] `json:"created_at,required" format:"date-time"`
	// For a Pending Transaction related to a transfer, this is the description you
	// provide. For a Pending Transaction related to a payment, this is the description
	// the vendor provides.
	Description fields.Field[string] `json:"description,required"`
	// The Pending Transaction identifier.
	ID fields.Field[string] `json:"id,required"`
	// The identifier for the route this Pending Transaction came through. Routes are
	// things like cards and ACH details.
	RouteID fields.Field[string] `json:"route_id,required,nullable"`
	// The type of the route this Pending Transaction came through.
	RouteType fields.Field[CardAuthorizationSimulationPendingTransactionRouteType] `json:"route_type,required,nullable"`
	// This is an object giving more details on the network-level event that caused the
	// Pending Transaction. For example, for a card transaction this lists the
	// merchant's industry and location.
	Source fields.Field[CardAuthorizationSimulationPendingTransactionSource] `json:"source,required"`
	// Whether the Pending Transaction has been confirmed and has an associated
	// Transaction.
	Status fields.Field[CardAuthorizationSimulationPendingTransactionStatus] `json:"status,required"`
	// A constant representing the object's type. For this resource it will always be
	// `pending_transaction`.
	Type fields.Field[CardAuthorizationSimulationPendingTransactionType] `json:"type,required"`
}

// MarshalJSON serializes CardAuthorizationSimulationPendingTransaction into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CardAuthorizationSimulationPendingTransaction) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CardAuthorizationSimulationPendingTransaction) String() (result string) {
	return fmt.Sprintf("&CardAuthorizationSimulationPendingTransaction{AccountID:%s Amount:%s Currency:%s CreatedAt:%s Description:%s ID:%s RouteID:%s RouteType:%s Source:%s Status:%s Type:%s}", r.AccountID, r.Amount, r.Currency, r.CreatedAt, r.Description, r.ID, r.RouteID, r.RouteType, r.Source, r.Status, r.Type)
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
	Category fields.Field[CardAuthorizationSimulationPendingTransactionSourceCategory] `json:"category,required"`
	// A Account Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `account_transfer_instruction`.
	AccountTransferInstruction fields.Field[CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstruction] `json:"account_transfer_instruction,required,nullable"`
	// A ACH Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `ach_transfer_instruction`.
	ACHTransferInstruction fields.Field[CardAuthorizationSimulationPendingTransactionSourceACHTransferInstruction] `json:"ach_transfer_instruction,required,nullable"`
	// A Card Authorization object. This field will be present in the JSON response if
	// and only if `category` is equal to `card_authorization`.
	CardAuthorization fields.Field[CardAuthorizationSimulationPendingTransactionSourceCardAuthorization] `json:"card_authorization,required,nullable"`
	// A Check Deposit Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_deposit_instruction`.
	CheckDepositInstruction fields.Field[CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstruction] `json:"check_deposit_instruction,required,nullable"`
	// A Check Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_instruction`.
	CheckTransferInstruction fields.Field[CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstruction] `json:"check_transfer_instruction,required,nullable"`
	// A Inbound Funds Hold object. This field will be present in the JSON response if
	// and only if `category` is equal to `inbound_funds_hold`.
	InboundFundsHold fields.Field[CardAuthorizationSimulationPendingTransactionSourceInboundFundsHold] `json:"inbound_funds_hold,required,nullable"`
	// A Deprecated Card Authorization object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_route_authorization`.
	CardRouteAuthorization fields.Field[CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorization] `json:"card_route_authorization,required,nullable"`
	// A Wire Drawdown Payment Instruction object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `wire_drawdown_payment_instruction`.
	WireDrawdownPaymentInstruction fields.Field[CardAuthorizationSimulationPendingTransactionSourceWireDrawdownPaymentInstruction] `json:"wire_drawdown_payment_instruction,required,nullable"`
	// A Wire Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_instruction`.
	WireTransferInstruction fields.Field[CardAuthorizationSimulationPendingTransactionSourceWireTransferInstruction] `json:"wire_transfer_instruction,required,nullable"`
}

// MarshalJSON serializes CardAuthorizationSimulationPendingTransactionSource into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CardAuthorizationSimulationPendingTransactionSource) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CardAuthorizationSimulationPendingTransactionSource) String() (result string) {
	return fmt.Sprintf("&CardAuthorizationSimulationPendingTransactionSource{Category:%s AccountTransferInstruction:%s ACHTransferInstruction:%s CardAuthorization:%s CheckDepositInstruction:%s CheckTransferInstruction:%s InboundFundsHold:%s CardRouteAuthorization:%s WireDrawdownPaymentInstruction:%s WireTransferInstruction:%s}", r.Category, r.AccountTransferInstruction, r.ACHTransferInstruction, r.CardAuthorization, r.CheckDepositInstruction, r.CheckTransferInstruction, r.InboundFundsHold, r.CardRouteAuthorization, r.WireDrawdownPaymentInstruction, r.WireTransferInstruction)
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
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency fields.Field[CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrency] `json:"currency,required"`
	// The identifier of the Account Transfer that led to this Pending Transaction.
	TransferID fields.Field[string] `json:"transfer_id,required"`
}

// MarshalJSON serializes
// CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstruction
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstruction) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstruction) String() (result string) {
	return fmt.Sprintf("&CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstruction{Amount:%s Currency:%s TransferID:%s}", r.Amount, r.Currency, r.TransferID)
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
	Amount fields.Field[int64] `json:"amount,required"`
	// The identifier of the ACH Transfer that led to this Pending Transaction.
	TransferID fields.Field[string] `json:"transfer_id,required"`
}

// MarshalJSON serializes
// CardAuthorizationSimulationPendingTransactionSourceACHTransferInstruction into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CardAuthorizationSimulationPendingTransactionSourceACHTransferInstruction) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CardAuthorizationSimulationPendingTransactionSourceACHTransferInstruction) String() (result string) {
	return fmt.Sprintf("&CardAuthorizationSimulationPendingTransactionSourceACHTransferInstruction{Amount:%s TransferID:%s}", r.Amount, r.TransferID)
}

type CardAuthorizationSimulationPendingTransactionSourceCardAuthorization struct {
	// The merchant identifier (commonly abbreviated as MID) of the merchant the card
	// is transacting with.
	MerchantAcceptorID fields.Field[string] `json:"merchant_acceptor_id,required"`
	// The merchant descriptor of the merchant the card is transacting with.
	MerchantDescriptor fields.Field[string] `json:"merchant_descriptor,required"`
	// The Merchant Category Code (commonly abbreviated as MCC) of the merchant the
	// card is transacting with.
	MerchantCategoryCode fields.Field[string] `json:"merchant_category_code,required,nullable"`
	// The city the merchant resides in.
	MerchantCity fields.Field[string] `json:"merchant_city,required,nullable"`
	// The country the merchant resides in.
	MerchantCountry fields.Field[string] `json:"merchant_country,required,nullable"`
	// The payment network used to process this card authorization
	Network fields.Field[CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetwork] `json:"network,required"`
	// Fields specific to the `network`
	NetworkDetails fields.Field[CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetails] `json:"network_details,required"`
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency fields.Field[CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrency] `json:"currency,required"`
	// The identifier of the Real-Time Decision sent to approve or decline this
	// transaction.
	RealTimeDecisionID fields.Field[string] `json:"real_time_decision_id,required,nullable"`
	// If the authorization was made via a Digital Wallet Token (such as an Apple Pay
	// purchase), the identifier of the token that was used.
	DigitalWalletTokenID fields.Field[string] `json:"digital_wallet_token_id,required,nullable"`
}

// MarshalJSON serializes
// CardAuthorizationSimulationPendingTransactionSourceCardAuthorization into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CardAuthorizationSimulationPendingTransactionSourceCardAuthorization) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CardAuthorizationSimulationPendingTransactionSourceCardAuthorization) String() (result string) {
	return fmt.Sprintf("&CardAuthorizationSimulationPendingTransactionSourceCardAuthorization{MerchantAcceptorID:%s MerchantDescriptor:%s MerchantCategoryCode:%s MerchantCity:%s MerchantCountry:%s Network:%s NetworkDetails:%s Amount:%s Currency:%s RealTimeDecisionID:%s DigitalWalletTokenID:%s}", r.MerchantAcceptorID, r.MerchantDescriptor, r.MerchantCategoryCode, r.MerchantCity, r.MerchantCountry, r.Network, r.NetworkDetails, r.Amount, r.Currency, r.RealTimeDecisionID, r.DigitalWalletTokenID)
}

type CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetwork string

const (
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkVisa CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetwork = "visa"
)

type CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetails struct {
	// Fields specific to the `visa` network
	Visa fields.Field[CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisa] `json:"visa,required"`
}

// MarshalJSON serializes
// CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetails
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetails) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetails) String() (result string) {
	return fmt.Sprintf("&CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetails{Visa:%s}", r.Visa)
}

type CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisa struct {
	// For electronic commerce transactions, this identifies the level of security used
	// in obtaining the customer's payment credential. For mail or telephone order
	// transactions, identifies the type of mail or telephone order.
	ElectronicCommerceIndicator fields.Field[CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator] `json:"electronic_commerce_indicator,required,nullable"`
	// The method used to enter the cardholder's primary account number and card
	// expiration date
	PointOfServiceEntryMode fields.Field[PointOfServiceEntryMode] `json:"point_of_service_entry_mode,required,nullable"`
}

// MarshalJSON serializes
// CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisa
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisa) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisa) String() (result string) {
	return fmt.Sprintf("&CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisa{ElectronicCommerceIndicator:%s PointOfServiceEntryMode:%s}", r.ElectronicCommerceIndicator, r.PointOfServiceEntryMode)
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

type CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency fields.Field[CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrency] `json:"currency,required"`
	// The identifier of the File containing the image of the front of the check that
	// was deposited.
	FrontImageFileID fields.Field[string] `json:"front_image_file_id,required"`
	// The identifier of the File containing the image of the back of the check that
	// was deposited.
	BackImageFileID fields.Field[string] `json:"back_image_file_id,required,nullable"`
	// The identifier of the Check Deposit.
	CheckDepositID fields.Field[string] `json:"check_deposit_id,required,nullable"`
}

// MarshalJSON serializes
// CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstruction into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstruction) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstruction) String() (result string) {
	return fmt.Sprintf("&CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstruction{Amount:%s Currency:%s FrontImageFileID:%s BackImageFileID:%s CheckDepositID:%s}", r.Amount, r.Currency, r.FrontImageFileID, r.BackImageFileID, r.CheckDepositID)
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
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
	// currency.
	Currency fields.Field[CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrency] `json:"currency,required"`
	// The identifier of the Check Transfer that led to this Pending Transaction.
	TransferID fields.Field[string] `json:"transfer_id,required"`
}

// MarshalJSON serializes
// CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstruction into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstruction) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstruction) String() (result string) {
	return fmt.Sprintf("&CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstruction{Amount:%s Currency:%s TransferID:%s}", r.Amount, r.Currency, r.TransferID)
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
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the hold's
	// currency.
	Currency fields.Field[CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrency] `json:"currency,required"`
	// When the hold will be released automatically. Certain conditions may cause it to
	// be released before this time.
	AutomaticallyReleasesAt fields.Field[time.Time] `json:"automatically_releases_at,required" format:"date-time"`
	// When the hold was released (if it has been released).
	ReleasedAt fields.Field[time.Time] `json:"released_at,required,nullable" format:"date-time"`
	// The status of the hold.
	Status fields.Field[CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldStatus] `json:"status,required"`
	// The ID of the Transaction for which funds were held.
	HeldTransactionID fields.Field[string] `json:"held_transaction_id,required,nullable"`
}

// MarshalJSON serializes
// CardAuthorizationSimulationPendingTransactionSourceInboundFundsHold into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CardAuthorizationSimulationPendingTransactionSourceInboundFundsHold) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CardAuthorizationSimulationPendingTransactionSourceInboundFundsHold) String() (result string) {
	return fmt.Sprintf("&CardAuthorizationSimulationPendingTransactionSourceInboundFundsHold{Amount:%s Currency:%s AutomaticallyReleasesAt:%s ReleasedAt:%s Status:%s HeldTransactionID:%s}", r.Amount, r.Currency, r.AutomaticallyReleasesAt, r.ReleasedAt, r.Status, r.HeldTransactionID)
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
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency             fields.Field[CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationCurrency] `json:"currency,required"`
	MerchantAcceptorID   fields.Field[string]                                                                            `json:"merchant_acceptor_id,required"`
	MerchantCity         fields.Field[string]                                                                            `json:"merchant_city,required,nullable"`
	MerchantCountry      fields.Field[string]                                                                            `json:"merchant_country,required"`
	MerchantDescriptor   fields.Field[string]                                                                            `json:"merchant_descriptor,required"`
	MerchantCategoryCode fields.Field[string]                                                                            `json:"merchant_category_code,required"`
	MerchantState        fields.Field[string]                                                                            `json:"merchant_state,required,nullable"`
}

// MarshalJSON serializes
// CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorization into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorization) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorization) String() (result string) {
	return fmt.Sprintf("&CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorization{Amount:%s Currency:%s MerchantAcceptorID:%s MerchantCity:%s MerchantCountry:%s MerchantDescriptor:%s MerchantCategoryCode:%s MerchantState:%s}", r.Amount, r.Currency, r.MerchantAcceptorID, r.MerchantCity, r.MerchantCountry, r.MerchantDescriptor, r.MerchantCategoryCode, r.MerchantState)
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

type CardAuthorizationSimulationPendingTransactionSourceWireDrawdownPaymentInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount             fields.Field[int64]  `json:"amount,required"`
	AccountNumber      fields.Field[string] `json:"account_number,required"`
	RoutingNumber      fields.Field[string] `json:"routing_number,required"`
	MessageToRecipient fields.Field[string] `json:"message_to_recipient,required"`
}

// MarshalJSON serializes
// CardAuthorizationSimulationPendingTransactionSourceWireDrawdownPaymentInstruction
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *CardAuthorizationSimulationPendingTransactionSourceWireDrawdownPaymentInstruction) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CardAuthorizationSimulationPendingTransactionSourceWireDrawdownPaymentInstruction) String() (result string) {
	return fmt.Sprintf("&CardAuthorizationSimulationPendingTransactionSourceWireDrawdownPaymentInstruction{Amount:%s AccountNumber:%s RoutingNumber:%s MessageToRecipient:%s}", r.Amount, r.AccountNumber, r.RoutingNumber, r.MessageToRecipient)
}

type CardAuthorizationSimulationPendingTransactionSourceWireTransferInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount             fields.Field[int64]  `json:"amount,required"`
	AccountNumber      fields.Field[string] `json:"account_number,required"`
	RoutingNumber      fields.Field[string] `json:"routing_number,required"`
	MessageToRecipient fields.Field[string] `json:"message_to_recipient,required"`
	TransferID         fields.Field[string] `json:"transfer_id,required"`
}

// MarshalJSON serializes
// CardAuthorizationSimulationPendingTransactionSourceWireTransferInstruction into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CardAuthorizationSimulationPendingTransactionSourceWireTransferInstruction) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CardAuthorizationSimulationPendingTransactionSourceWireTransferInstruction) String() (result string) {
	return fmt.Sprintf("&CardAuthorizationSimulationPendingTransactionSourceWireTransferInstruction{Amount:%s AccountNumber:%s RoutingNumber:%s MessageToRecipient:%s TransferID:%s}", r.Amount, r.AccountNumber, r.RoutingNumber, r.MessageToRecipient, r.TransferID)
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
	AccountID fields.Field[string] `json:"account_id,required"`
	// The Declined Transaction amount in the minor unit of its currency. For dollars,
	// for example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Declined
	// Transaction's currency. This will match the currency on the Declined
	// Transcation's Account.
	Currency fields.Field[CardAuthorizationSimulationDeclinedTransactionCurrency] `json:"currency,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the
	// Transaction occured.
	CreatedAt fields.Field[time.Time] `json:"created_at,required" format:"date-time"`
	// This is the description the vendor provides.
	Description fields.Field[string] `json:"description,required"`
	// The Declined Transaction identifier.
	ID fields.Field[string] `json:"id,required"`
	// The identifier for the route this Declined Transaction came through. Routes are
	// things like cards and ACH details.
	RouteID fields.Field[string] `json:"route_id,required,nullable"`
	// The type of the route this Declined Transaction came through.
	RouteType fields.Field[CardAuthorizationSimulationDeclinedTransactionRouteType] `json:"route_type,required,nullable"`
	// This is an object giving more details on the network-level event that caused the
	// Declined Transaction. For example, for a card transaction this lists the
	// merchant's industry and location. Note that for backwards compatibility reasons,
	// additional undocumented keys may appear in this object. These should be treated
	// as deprecated and will be removed in the future.
	Source fields.Field[CardAuthorizationSimulationDeclinedTransactionSource] `json:"source,required"`
	// A constant representing the object's type. For this resource it will always be
	// `declined_transaction`.
	Type fields.Field[CardAuthorizationSimulationDeclinedTransactionType] `json:"type,required"`
}

// MarshalJSON serializes CardAuthorizationSimulationDeclinedTransaction into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CardAuthorizationSimulationDeclinedTransaction) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CardAuthorizationSimulationDeclinedTransaction) String() (result string) {
	return fmt.Sprintf("&CardAuthorizationSimulationDeclinedTransaction{AccountID:%s Amount:%s Currency:%s CreatedAt:%s Description:%s ID:%s RouteID:%s RouteType:%s Source:%s Type:%s}", r.AccountID, r.Amount, r.Currency, r.CreatedAt, r.Description, r.ID, r.RouteID, r.RouteType, r.Source, r.Type)
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
	Category fields.Field[CardAuthorizationSimulationDeclinedTransactionSourceCategory] `json:"category,required"`
	// A ACH Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `ach_decline`.
	ACHDecline fields.Field[CardAuthorizationSimulationDeclinedTransactionSourceACHDecline] `json:"ach_decline,required,nullable"`
	// A Card Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_decline`.
	CardDecline fields.Field[CardAuthorizationSimulationDeclinedTransactionSourceCardDecline] `json:"card_decline,required,nullable"`
	// A Check Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `check_decline`.
	CheckDecline fields.Field[CardAuthorizationSimulationDeclinedTransactionSourceCheckDecline] `json:"check_decline,required,nullable"`
	// A Inbound Real Time Payments Transfer Decline object. This field will be present
	// in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_decline`.
	InboundRealTimePaymentsTransferDecline fields.Field[CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline] `json:"inbound_real_time_payments_transfer_decline,required,nullable"`
	// A International ACH Decline object. This field will be present in the JSON
	// response if and only if `category` is equal to `international_ach_decline`.
	InternationalACHDecline fields.Field[CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline] `json:"international_ach_decline,required,nullable"`
	// A Deprecated Card Decline object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_route_decline`.
	CardRouteDecline fields.Field[CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDecline] `json:"card_route_decline,required,nullable"`
}

// MarshalJSON serializes CardAuthorizationSimulationDeclinedTransactionSource into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CardAuthorizationSimulationDeclinedTransactionSource) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CardAuthorizationSimulationDeclinedTransactionSource) String() (result string) {
	return fmt.Sprintf("&CardAuthorizationSimulationDeclinedTransactionSource{Category:%s ACHDecline:%s CardDecline:%s CheckDecline:%s InboundRealTimePaymentsTransferDecline:%s InternationalACHDecline:%s CardRouteDecline:%s}", r.Category, r.ACHDecline, r.CardDecline, r.CheckDecline, r.InboundRealTimePaymentsTransferDecline, r.InternationalACHDecline, r.CardRouteDecline)
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
	Amount                             fields.Field[int64]  `json:"amount,required"`
	OriginatorCompanyName              fields.Field[string] `json:"originator_company_name,required"`
	OriginatorCompanyDescriptiveDate   fields.Field[string] `json:"originator_company_descriptive_date,required,nullable"`
	OriginatorCompanyDiscretionaryData fields.Field[string] `json:"originator_company_discretionary_data,required,nullable"`
	OriginatorCompanyID                fields.Field[string] `json:"originator_company_id,required"`
	// Why the ACH transfer was declined.
	Reason           fields.Field[CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason] `json:"reason,required"`
	ReceiverIDNumber fields.Field[string]                                                               `json:"receiver_id_number,required,nullable"`
	ReceiverName     fields.Field[string]                                                               `json:"receiver_name,required,nullable"`
	TraceNumber      fields.Field[string]                                                               `json:"trace_number,required"`
}

// MarshalJSON serializes
// CardAuthorizationSimulationDeclinedTransactionSourceACHDecline into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceACHDecline) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CardAuthorizationSimulationDeclinedTransactionSourceACHDecline) String() (result string) {
	return fmt.Sprintf("&CardAuthorizationSimulationDeclinedTransactionSourceACHDecline{Amount:%s OriginatorCompanyName:%s OriginatorCompanyDescriptiveDate:%s OriginatorCompanyDiscretionaryData:%s OriginatorCompanyID:%s Reason:%s ReceiverIDNumber:%s ReceiverName:%s TraceNumber:%s}", r.Amount, r.OriginatorCompanyName, r.OriginatorCompanyDescriptiveDate, r.OriginatorCompanyDiscretionaryData, r.OriginatorCompanyID, r.Reason, r.ReceiverIDNumber, r.ReceiverName, r.TraceNumber)
}

type CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason string

const (
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonACHRouteCanceled             CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "ach_route_canceled"
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonACHRouteDisabled             CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "ach_route_disabled"
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonBreachesLimit                CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "breaches_limit"
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonCreditEntryRefusedByReceiver CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "credit_entry_refused_by_receiver"
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonDuplicateReturn              CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "duplicate_return"
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonEntityNotActive              CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "entity_not_active"
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonTransactionNotAllowed        CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "transaction_not_allowed"
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonGroupLocked                  CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "group_locked"
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonInsufficientFunds            CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "insufficient_funds"
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonNoACHRoute                   CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "no_ach_route"
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonOriginatorRequest            CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "originator_request"
)

type CardAuthorizationSimulationDeclinedTransactionSourceCardDecline struct {
	// The merchant identifier (commonly abbreviated as MID) of the merchant the card
	// is transacting with.
	MerchantAcceptorID fields.Field[string] `json:"merchant_acceptor_id,required"`
	// The merchant descriptor of the merchant the card is transacting with.
	MerchantDescriptor fields.Field[string] `json:"merchant_descriptor,required"`
	// The Merchant Category Code (commonly abbreviated as MCC) of the merchant the
	// card is transacting with.
	MerchantCategoryCode fields.Field[string] `json:"merchant_category_code,required,nullable"`
	// The city the merchant resides in.
	MerchantCity fields.Field[string] `json:"merchant_city,required,nullable"`
	// The country the merchant resides in.
	MerchantCountry fields.Field[string] `json:"merchant_country,required,nullable"`
	// The payment network used to process this card authorization
	Network fields.Field[CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetwork] `json:"network,required"`
	// Fields specific to the `network`
	NetworkDetails fields.Field[CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetails] `json:"network_details,required"`
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency fields.Field[CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrency] `json:"currency,required"`
	// Why the transaction was declined.
	Reason fields.Field[CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason] `json:"reason,required"`
	// The state the merchant resides in.
	MerchantState fields.Field[string] `json:"merchant_state,required,nullable"`
	// The identifier of the Real-Time Decision sent to approve or decline this
	// transaction.
	RealTimeDecisionID fields.Field[string] `json:"real_time_decision_id,required,nullable"`
	// If the authorization was attempted using a Digital Wallet Token (such as an
	// Apple Pay purchase), the identifier of the token that was used.
	DigitalWalletTokenID fields.Field[string] `json:"digital_wallet_token_id,required,nullable"`
}

// MarshalJSON serializes
// CardAuthorizationSimulationDeclinedTransactionSourceCardDecline into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDecline) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CardAuthorizationSimulationDeclinedTransactionSourceCardDecline) String() (result string) {
	return fmt.Sprintf("&CardAuthorizationSimulationDeclinedTransactionSourceCardDecline{MerchantAcceptorID:%s MerchantDescriptor:%s MerchantCategoryCode:%s MerchantCity:%s MerchantCountry:%s Network:%s NetworkDetails:%s Amount:%s Currency:%s Reason:%s MerchantState:%s RealTimeDecisionID:%s DigitalWalletTokenID:%s}", r.MerchantAcceptorID, r.MerchantDescriptor, r.MerchantCategoryCode, r.MerchantCity, r.MerchantCountry, r.Network, r.NetworkDetails, r.Amount, r.Currency, r.Reason, r.MerchantState, r.RealTimeDecisionID, r.DigitalWalletTokenID)
}

type CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetwork string

const (
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkVisa CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetwork = "visa"
)

type CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetails struct {
	// Fields specific to the `visa` network
	Visa fields.Field[CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa] `json:"visa,required"`
}

// MarshalJSON serializes
// CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetails
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetails) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetails) String() (result string) {
	return fmt.Sprintf("&CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetails{Visa:%s}", r.Visa)
}

type CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa struct {
	// For electronic commerce transactions, this identifies the level of security used
	// in obtaining the customer's payment credential. For mail or telephone order
	// transactions, identifies the type of mail or telephone order.
	ElectronicCommerceIndicator fields.Field[CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator] `json:"electronic_commerce_indicator,required,nullable"`
	// The method used to enter the cardholder's primary account number and card
	// expiration date
	PointOfServiceEntryMode fields.Field[PointOfServiceEntryMode] `json:"point_of_service_entry_mode,required,nullable"`
}

// MarshalJSON serializes
// CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa) String() (result string) {
	return fmt.Sprintf("&CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa{ElectronicCommerceIndicator:%s PointOfServiceEntryMode:%s}", r.ElectronicCommerceIndicator, r.PointOfServiceEntryMode)
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
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonCardNotActive               CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "card_not_active"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonEntityNotActive             CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "entity_not_active"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonGroupLocked                 CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "group_locked"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonInsufficientFunds           CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "insufficient_funds"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonCvv2Mismatch                CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "cvv2_mismatch"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonTransactionNotAllowed       CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "transaction_not_allowed"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonBreachesLimit               CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "breaches_limit"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonWebhookDeclined             CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "webhook_declined"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonWebhookTimedOut             CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "webhook_timed_out"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonDeclinedByStandInProcessing CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "declined_by_stand_in_processing"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonInvalidPhysicalCard         CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "invalid_physical_card"
)

type CardAuthorizationSimulationDeclinedTransactionSourceCheckDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount        fields.Field[int64]  `json:"amount,required"`
	AuxiliaryOnUs fields.Field[string] `json:"auxiliary_on_us,required,nullable"`
	// Why the check was declined.
	Reason fields.Field[CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason] `json:"reason,required"`
}

// MarshalJSON serializes
// CardAuthorizationSimulationDeclinedTransactionSourceCheckDecline into an array
// of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceCheckDecline) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CardAuthorizationSimulationDeclinedTransactionSourceCheckDecline) String() (result string) {
	return fmt.Sprintf("&CardAuthorizationSimulationDeclinedTransactionSourceCheckDecline{Amount:%s AuxiliaryOnUs:%s Reason:%s}", r.Amount, r.AuxiliaryOnUs, r.Reason)
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
)

type CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the declined
	// transfer's currency. This will always be "USD" for a Real Time Payments
	// transfer.
	Currency fields.Field[CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency] `json:"currency,required"`
	// Why the transfer was declined.
	Reason fields.Field[CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason] `json:"reason,required"`
	// The name the sender of the transfer specified as the recipient of the transfer.
	CreditorName fields.Field[string] `json:"creditor_name,required"`
	// The name provided by the sender of the transfer.
	DebtorName fields.Field[string] `json:"debtor_name,required"`
	// The account number of the account that sent the transfer.
	DebtorAccountNumber fields.Field[string] `json:"debtor_account_number,required"`
	// The routing number of the account that sent the transfer.
	DebtorRoutingNumber fields.Field[string] `json:"debtor_routing_number,required"`
	// The Real Time Payments network identification of the declined transfer.
	TransactionIdentification fields.Field[string] `json:"transaction_identification,required"`
	// Additional information included with the transfer.
	RemittanceInformation fields.Field[string] `json:"remittance_information,required,nullable"`
}

// MarshalJSON serializes
// CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) String() (result string) {
	return fmt.Sprintf("&CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline{Amount:%s Currency:%s Reason:%s CreditorName:%s DebtorName:%s DebtorAccountNumber:%s DebtorRoutingNumber:%s TransactionIdentification:%s RemittanceInformation:%s}", r.Amount, r.Currency, r.Reason, r.CreditorName, r.DebtorName, r.DebtorAccountNumber, r.DebtorRoutingNumber, r.TransactionIdentification, r.RemittanceInformation)
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
	Amount                                                 fields.Field[int64]  `json:"amount,required"`
	ForeignExchangeIndicator                               fields.Field[string] `json:"foreign_exchange_indicator,required"`
	ForeignExchangeReferenceIndicator                      fields.Field[string] `json:"foreign_exchange_reference_indicator,required"`
	ForeignExchangeReference                               fields.Field[string] `json:"foreign_exchange_reference,required,nullable"`
	DestinationCountryCode                                 fields.Field[string] `json:"destination_country_code,required"`
	DestinationCurrencyCode                                fields.Field[string] `json:"destination_currency_code,required"`
	ForeignPaymentAmount                                   fields.Field[int64]  `json:"foreign_payment_amount,required"`
	ForeignTraceNumber                                     fields.Field[string] `json:"foreign_trace_number,required,nullable"`
	InternationalTransactionTypeCode                       fields.Field[string] `json:"international_transaction_type_code,required"`
	OriginatingCurrencyCode                                fields.Field[string] `json:"originating_currency_code,required"`
	OriginatingDepositoryFinancialInstitutionName          fields.Field[string] `json:"originating_depository_financial_institution_name,required"`
	OriginatingDepositoryFinancialInstitutionIDQualifier   fields.Field[string] `json:"originating_depository_financial_institution_id_qualifier,required"`
	OriginatingDepositoryFinancialInstitutionID            fields.Field[string] `json:"originating_depository_financial_institution_id,required"`
	OriginatingDepositoryFinancialInstitutionBranchCountry fields.Field[string] `json:"originating_depository_financial_institution_branch_country,required"`
	OriginatorCity                                         fields.Field[string] `json:"originator_city,required"`
	OriginatorCompanyEntryDescription                      fields.Field[string] `json:"originator_company_entry_description,required"`
	OriginatorCountry                                      fields.Field[string] `json:"originator_country,required"`
	OriginatorIdentification                               fields.Field[string] `json:"originator_identification,required"`
	OriginatorName                                         fields.Field[string] `json:"originator_name,required"`
	OriginatorPostalCode                                   fields.Field[string] `json:"originator_postal_code,required,nullable"`
	OriginatorStreetAddress                                fields.Field[string] `json:"originator_street_address,required"`
	OriginatorStateOrProvince                              fields.Field[string] `json:"originator_state_or_province,required,nullable"`
	PaymentRelatedInformation                              fields.Field[string] `json:"payment_related_information,required,nullable"`
	PaymentRelatedInformation2                             fields.Field[string] `json:"payment_related_information2,required,nullable"`
	ReceiverIdentificationNumber                           fields.Field[string] `json:"receiver_identification_number,required,nullable"`
	ReceiverStreetAddress                                  fields.Field[string] `json:"receiver_street_address,required"`
	ReceiverCity                                           fields.Field[string] `json:"receiver_city,required"`
	ReceiverStateOrProvince                                fields.Field[string] `json:"receiver_state_or_province,required,nullable"`
	ReceiverCountry                                        fields.Field[string] `json:"receiver_country,required"`
	ReceiverPostalCode                                     fields.Field[string] `json:"receiver_postal_code,required,nullable"`
	ReceivingCompanyOrIndividualName                       fields.Field[string] `json:"receiving_company_or_individual_name,required"`
	ReceivingDepositoryFinancialInstitutionName            fields.Field[string] `json:"receiving_depository_financial_institution_name,required"`
	ReceivingDepositoryFinancialInstitutionIDQualifier     fields.Field[string] `json:"receiving_depository_financial_institution_id_qualifier,required"`
	ReceivingDepositoryFinancialInstitutionID              fields.Field[string] `json:"receiving_depository_financial_institution_id,required"`
	ReceivingDepositoryFinancialInstitutionCountry         fields.Field[string] `json:"receiving_depository_financial_institution_country,required"`
	TraceNumber                                            fields.Field[string] `json:"trace_number,required"`
}

// MarshalJSON serializes
// CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) String() (result string) {
	return fmt.Sprintf("&CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline{Amount:%s ForeignExchangeIndicator:%s ForeignExchangeReferenceIndicator:%s ForeignExchangeReference:%s DestinationCountryCode:%s DestinationCurrencyCode:%s ForeignPaymentAmount:%s ForeignTraceNumber:%s InternationalTransactionTypeCode:%s OriginatingCurrencyCode:%s OriginatingDepositoryFinancialInstitutionName:%s OriginatingDepositoryFinancialInstitutionIDQualifier:%s OriginatingDepositoryFinancialInstitutionID:%s OriginatingDepositoryFinancialInstitutionBranchCountry:%s OriginatorCity:%s OriginatorCompanyEntryDescription:%s OriginatorCountry:%s OriginatorIdentification:%s OriginatorName:%s OriginatorPostalCode:%s OriginatorStreetAddress:%s OriginatorStateOrProvince:%s PaymentRelatedInformation:%s PaymentRelatedInformation2:%s ReceiverIdentificationNumber:%s ReceiverStreetAddress:%s ReceiverCity:%s ReceiverStateOrProvince:%s ReceiverCountry:%s ReceiverPostalCode:%s ReceivingCompanyOrIndividualName:%s ReceivingDepositoryFinancialInstitutionName:%s ReceivingDepositoryFinancialInstitutionIDQualifier:%s ReceivingDepositoryFinancialInstitutionID:%s ReceivingDepositoryFinancialInstitutionCountry:%s TraceNumber:%s}", r.Amount, r.ForeignExchangeIndicator, r.ForeignExchangeReferenceIndicator, r.ForeignExchangeReference, r.DestinationCountryCode, r.DestinationCurrencyCode, r.ForeignPaymentAmount, r.ForeignTraceNumber, r.InternationalTransactionTypeCode, r.OriginatingCurrencyCode, r.OriginatingDepositoryFinancialInstitutionName, r.OriginatingDepositoryFinancialInstitutionIDQualifier, r.OriginatingDepositoryFinancialInstitutionID, r.OriginatingDepositoryFinancialInstitutionBranchCountry, r.OriginatorCity, r.OriginatorCompanyEntryDescription, r.OriginatorCountry, r.OriginatorIdentification, r.OriginatorName, r.OriginatorPostalCode, r.OriginatorStreetAddress, r.OriginatorStateOrProvince, r.PaymentRelatedInformation, r.PaymentRelatedInformation2, r.ReceiverIdentificationNumber, r.ReceiverStreetAddress, r.ReceiverCity, r.ReceiverStateOrProvince, r.ReceiverCountry, r.ReceiverPostalCode, r.ReceivingCompanyOrIndividualName, r.ReceivingDepositoryFinancialInstitutionName, r.ReceivingDepositoryFinancialInstitutionIDQualifier, r.ReceivingDepositoryFinancialInstitutionID, r.ReceivingDepositoryFinancialInstitutionCountry, r.TraceNumber)
}

type CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency             fields.Field[CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDeclineCurrency] `json:"currency,required"`
	MerchantAcceptorID   fields.Field[string]                                                                       `json:"merchant_acceptor_id,required"`
	MerchantCity         fields.Field[string]                                                                       `json:"merchant_city,required,nullable"`
	MerchantCountry      fields.Field[string]                                                                       `json:"merchant_country,required"`
	MerchantDescriptor   fields.Field[string]                                                                       `json:"merchant_descriptor,required"`
	MerchantState        fields.Field[string]                                                                       `json:"merchant_state,required,nullable"`
	MerchantCategoryCode fields.Field[string]                                                                       `json:"merchant_category_code,required,nullable"`
}

// MarshalJSON serializes
// CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDecline into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDecline) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDecline) String() (result string) {
	return fmt.Sprintf("&CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDecline{Amount:%s Currency:%s MerchantAcceptorID:%s MerchantCity:%s MerchantCountry:%s MerchantDescriptor:%s MerchantState:%s MerchantCategoryCode:%s}", r.Amount, r.Currency, r.MerchantAcceptorID, r.MerchantCity, r.MerchantCountry, r.MerchantDescriptor, r.MerchantState, r.MerchantCategoryCode)
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

type SimulateAnAuthorizationOnACardParameters struct {
	// The authorization amount in cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The identifier of the Card to be authorized.
	CardID fields.Field[string] `json:"card_id"`
	// The identifier of the Digital Wallet Token to be authorized.
	DigitalWalletTokenID fields.Field[string] `json:"digital_wallet_token_id"`
}

// MarshalJSON serializes SimulateAnAuthorizationOnACardParameters into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *SimulateAnAuthorizationOnACardParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r SimulateAnAuthorizationOnACardParameters) String() (result string) {
	return fmt.Sprintf("&SimulateAnAuthorizationOnACardParameters{Amount:%s CardID:%s DigitalWalletTokenID:%s}", r.Amount, r.CardID, r.DigitalWalletTokenID)
}

type SimulateSettlingACardAuthorizationParameters struct {
	// The identifier of the Card to create a settlement on.
	CardID fields.Field[string] `json:"card_id,required"`
	// The identifier of the Pending Transaction for the Card Authorization you wish to
	// settle.
	PendingTransactionID fields.Field[string] `json:"pending_transaction_id,required"`
	// The amount to be settled. This defaults to the amount of the Pending Transaction
	// being settled.
	Amount fields.Field[int64] `json:"amount"`
}

// MarshalJSON serializes SimulateSettlingACardAuthorizationParameters into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *SimulateSettlingACardAuthorizationParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r SimulateSettlingACardAuthorizationParameters) String() (result string) {
	return fmt.Sprintf("&SimulateSettlingACardAuthorizationParameters{CardID:%s PendingTransactionID:%s Amount:%s}", r.CardID, r.PendingTransactionID, r.Amount)
}

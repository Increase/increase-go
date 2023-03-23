package requests

import (
	"fmt"
	"net/url"
	"time"

	"github.com/increase/increase-go/core"
	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/fields"
)

type PendingTransaction struct {
	// The identifier for the account this Pending Transaction belongs to.
	AccountID fields.Field[string] `json:"account_id,required"`
	// The Pending Transaction amount in the minor unit of its currency. For dollars,
	// for example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Pending
	// Transaction's currency. This will match the currency on the Pending
	// Transcation's Account.
	Currency fields.Field[PendingTransactionCurrency] `json:"currency,required"`
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
	RouteType fields.Field[string] `json:"route_type,required,nullable"`
	// This is an object giving more details on the network-level event that caused the
	// Pending Transaction. For example, for a card transaction this lists the
	// merchant's industry and location.
	Source fields.Field[PendingTransactionSource] `json:"source,required"`
	// Whether the Pending Transaction has been confirmed and has an associated
	// Transaction.
	Status fields.Field[PendingTransactionStatus] `json:"status,required"`
	// A constant representing the object's type. For this resource it will always be
	// `pending_transaction`.
	Type fields.Field[PendingTransactionType] `json:"type,required"`
}

// MarshalJSON serializes PendingTransaction into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *PendingTransaction) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r PendingTransaction) String() (result string) {
	return fmt.Sprintf("&PendingTransaction{AccountID:%s Amount:%s Currency:%s CreatedAt:%s Description:%s ID:%s RouteID:%s RouteType:%s Source:%s Status:%s Type:%s}", r.AccountID, r.Amount, r.Currency, r.CreatedAt, r.Description, r.ID, r.RouteID, r.RouteType, r.Source, r.Status, r.Type)
}

type PendingTransactionCurrency string

const (
	PendingTransactionCurrencyCad PendingTransactionCurrency = "CAD"
	PendingTransactionCurrencyChf PendingTransactionCurrency = "CHF"
	PendingTransactionCurrencyEur PendingTransactionCurrency = "EUR"
	PendingTransactionCurrencyGbp PendingTransactionCurrency = "GBP"
	PendingTransactionCurrencyJpy PendingTransactionCurrency = "JPY"
	PendingTransactionCurrencyUsd PendingTransactionCurrency = "USD"
)

type PendingTransactionSource struct {
	// The type of transaction that took place. We may add additional possible values
	// for this enum over time; your application should be able to handle such
	// additions gracefully.
	Category fields.Field[PendingTransactionSourceCategory] `json:"category,required"`
	// A Account Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `account_transfer_instruction`.
	AccountTransferInstruction fields.Field[PendingTransactionSourceAccountTransferInstruction] `json:"account_transfer_instruction,required,nullable"`
	// A ACH Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `ach_transfer_instruction`.
	ACHTransferInstruction fields.Field[PendingTransactionSourceACHTransferInstruction] `json:"ach_transfer_instruction,required,nullable"`
	// A Card Authorization object. This field will be present in the JSON response if
	// and only if `category` is equal to `card_authorization`.
	CardAuthorization fields.Field[PendingTransactionSourceCardAuthorization] `json:"card_authorization,required,nullable"`
	// A Check Deposit Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_deposit_instruction`.
	CheckDepositInstruction fields.Field[PendingTransactionSourceCheckDepositInstruction] `json:"check_deposit_instruction,required,nullable"`
	// A Check Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_instruction`.
	CheckTransferInstruction fields.Field[PendingTransactionSourceCheckTransferInstruction] `json:"check_transfer_instruction,required,nullable"`
	// A Inbound Funds Hold object. This field will be present in the JSON response if
	// and only if `category` is equal to `inbound_funds_hold`.
	InboundFundsHold fields.Field[PendingTransactionSourceInboundFundsHold] `json:"inbound_funds_hold,required,nullable"`
	// A Deprecated Card Authorization object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_route_authorization`.
	CardRouteAuthorization fields.Field[PendingTransactionSourceCardRouteAuthorization] `json:"card_route_authorization,required,nullable"`
	// A Wire Drawdown Payment Instruction object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `wire_drawdown_payment_instruction`.
	WireDrawdownPaymentInstruction fields.Field[PendingTransactionSourceWireDrawdownPaymentInstruction] `json:"wire_drawdown_payment_instruction,required,nullable"`
	// A Wire Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_instruction`.
	WireTransferInstruction fields.Field[PendingTransactionSourceWireTransferInstruction] `json:"wire_transfer_instruction,required,nullable"`
}

// MarshalJSON serializes PendingTransactionSource into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *PendingTransactionSource) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r PendingTransactionSource) String() (result string) {
	return fmt.Sprintf("&PendingTransactionSource{Category:%s AccountTransferInstruction:%s ACHTransferInstruction:%s CardAuthorization:%s CheckDepositInstruction:%s CheckTransferInstruction:%s InboundFundsHold:%s CardRouteAuthorization:%s WireDrawdownPaymentInstruction:%s WireTransferInstruction:%s}", r.Category, r.AccountTransferInstruction, r.ACHTransferInstruction, r.CardAuthorization, r.CheckDepositInstruction, r.CheckTransferInstruction, r.InboundFundsHold, r.CardRouteAuthorization, r.WireDrawdownPaymentInstruction, r.WireTransferInstruction)
}

type PendingTransactionSourceCategory string

const (
	PendingTransactionSourceCategoryAccountTransferInstruction          PendingTransactionSourceCategory = "account_transfer_instruction"
	PendingTransactionSourceCategoryACHTransferInstruction              PendingTransactionSourceCategory = "ach_transfer_instruction"
	PendingTransactionSourceCategoryCardAuthorization                   PendingTransactionSourceCategory = "card_authorization"
	PendingTransactionSourceCategoryCheckDepositInstruction             PendingTransactionSourceCategory = "check_deposit_instruction"
	PendingTransactionSourceCategoryCheckTransferInstruction            PendingTransactionSourceCategory = "check_transfer_instruction"
	PendingTransactionSourceCategoryInboundFundsHold                    PendingTransactionSourceCategory = "inbound_funds_hold"
	PendingTransactionSourceCategoryCardRouteAuthorization              PendingTransactionSourceCategory = "card_route_authorization"
	PendingTransactionSourceCategoryRealTimePaymentsTransferInstruction PendingTransactionSourceCategory = "real_time_payments_transfer_instruction"
	PendingTransactionSourceCategoryWireDrawdownPaymentInstruction      PendingTransactionSourceCategory = "wire_drawdown_payment_instruction"
	PendingTransactionSourceCategoryWireTransferInstruction             PendingTransactionSourceCategory = "wire_transfer_instruction"
	PendingTransactionSourceCategoryOther                               PendingTransactionSourceCategory = "other"
)

type PendingTransactionSourceAccountTransferInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency fields.Field[PendingTransactionSourceAccountTransferInstructionCurrency] `json:"currency,required"`
	// The identifier of the Account Transfer that led to this Pending Transaction.
	TransferID fields.Field[string] `json:"transfer_id,required"`
}

// MarshalJSON serializes PendingTransactionSourceAccountTransferInstruction into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *PendingTransactionSourceAccountTransferInstruction) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r PendingTransactionSourceAccountTransferInstruction) String() (result string) {
	return fmt.Sprintf("&PendingTransactionSourceAccountTransferInstruction{Amount:%s Currency:%s TransferID:%s}", r.Amount, r.Currency, r.TransferID)
}

type PendingTransactionSourceAccountTransferInstructionCurrency string

const (
	PendingTransactionSourceAccountTransferInstructionCurrencyCad PendingTransactionSourceAccountTransferInstructionCurrency = "CAD"
	PendingTransactionSourceAccountTransferInstructionCurrencyChf PendingTransactionSourceAccountTransferInstructionCurrency = "CHF"
	PendingTransactionSourceAccountTransferInstructionCurrencyEur PendingTransactionSourceAccountTransferInstructionCurrency = "EUR"
	PendingTransactionSourceAccountTransferInstructionCurrencyGbp PendingTransactionSourceAccountTransferInstructionCurrency = "GBP"
	PendingTransactionSourceAccountTransferInstructionCurrencyJpy PendingTransactionSourceAccountTransferInstructionCurrency = "JPY"
	PendingTransactionSourceAccountTransferInstructionCurrencyUsd PendingTransactionSourceAccountTransferInstructionCurrency = "USD"
)

type PendingTransactionSourceACHTransferInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The identifier of the ACH Transfer that led to this Pending Transaction.
	TransferID fields.Field[string] `json:"transfer_id,required"`
}

// MarshalJSON serializes PendingTransactionSourceACHTransferInstruction into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *PendingTransactionSourceACHTransferInstruction) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r PendingTransactionSourceACHTransferInstruction) String() (result string) {
	return fmt.Sprintf("&PendingTransactionSourceACHTransferInstruction{Amount:%s TransferID:%s}", r.Amount, r.TransferID)
}

type PendingTransactionSourceCardAuthorization struct {
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
	Network fields.Field[PendingTransactionSourceCardAuthorizationNetwork] `json:"network,required"`
	// Fields specific to the `network`
	NetworkDetails fields.Field[PendingTransactionSourceCardAuthorizationNetworkDetails] `json:"network_details,required"`
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency fields.Field[PendingTransactionSourceCardAuthorizationCurrency] `json:"currency,required"`
	// The identifier of the Real-Time Decision sent to approve or decline this
	// transaction.
	RealTimeDecisionID fields.Field[string] `json:"real_time_decision_id,required,nullable"`
	// If the authorization was made via a Digital Wallet Token (such as an Apple Pay
	// purchase), the identifier of the token that was used.
	DigitalWalletTokenID fields.Field[string] `json:"digital_wallet_token_id,required,nullable"`
}

// MarshalJSON serializes PendingTransactionSourceCardAuthorization into an array
// of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *PendingTransactionSourceCardAuthorization) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r PendingTransactionSourceCardAuthorization) String() (result string) {
	return fmt.Sprintf("&PendingTransactionSourceCardAuthorization{MerchantAcceptorID:%s MerchantDescriptor:%s MerchantCategoryCode:%s MerchantCity:%s MerchantCountry:%s Network:%s NetworkDetails:%s Amount:%s Currency:%s RealTimeDecisionID:%s DigitalWalletTokenID:%s}", r.MerchantAcceptorID, r.MerchantDescriptor, r.MerchantCategoryCode, r.MerchantCity, r.MerchantCountry, r.Network, r.NetworkDetails, r.Amount, r.Currency, r.RealTimeDecisionID, r.DigitalWalletTokenID)
}

type PendingTransactionSourceCardAuthorizationNetwork string

const (
	PendingTransactionSourceCardAuthorizationNetworkVisa PendingTransactionSourceCardAuthorizationNetwork = "visa"
)

type PendingTransactionSourceCardAuthorizationNetworkDetails struct {
	// Fields specific to the `visa` network
	Visa fields.Field[PendingTransactionSourceCardAuthorizationNetworkDetailsVisa] `json:"visa,required"`
}

// MarshalJSON serializes PendingTransactionSourceCardAuthorizationNetworkDetails
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *PendingTransactionSourceCardAuthorizationNetworkDetails) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r PendingTransactionSourceCardAuthorizationNetworkDetails) String() (result string) {
	return fmt.Sprintf("&PendingTransactionSourceCardAuthorizationNetworkDetails{Visa:%s}", r.Visa)
}

type PendingTransactionSourceCardAuthorizationNetworkDetailsVisa struct {
	// For electronic commerce transactions, this identifies the level of security used
	// in obtaining the customer's payment credential. For mail or telephone order
	// transactions, identifies the type of mail or telephone order.
	ElectronicCommerceIndicator fields.Field[PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator] `json:"electronic_commerce_indicator,required,nullable"`
	// The method used to enter the cardholder's primary account number and card
	// expiration date
	PointOfServiceEntryMode fields.Field[PointOfServiceEntryMode] `json:"point_of_service_entry_mode,required,nullable"`
}

// MarshalJSON serializes
// PendingTransactionSourceCardAuthorizationNetworkDetailsVisa into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *PendingTransactionSourceCardAuthorizationNetworkDetailsVisa) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r PendingTransactionSourceCardAuthorizationNetworkDetailsVisa) String() (result string) {
	return fmt.Sprintf("&PendingTransactionSourceCardAuthorizationNetworkDetailsVisa{ElectronicCommerceIndicator:%s PointOfServiceEntryMode:%s}", r.ElectronicCommerceIndicator, r.PointOfServiceEntryMode)
}

type PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator string

const (
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorMailPhoneOrder                                           PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "mail_phone_order"
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorRecurring                                                PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "recurring"
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorInstallment                                              PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "installment"
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorUnknownMailPhoneOrder                                    PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "unknown_mail_phone_order"
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorSecureElectronicCommerce                                 PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "secure_electronic_commerce"
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransactionAt_3dsCapableMerchant PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction_at_3ds_capable_merchant"
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransaction                      PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction"
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonSecureTransaction                                     PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "non_secure_transaction"
)

type PendingTransactionSourceCardAuthorizationCurrency string

const (
	PendingTransactionSourceCardAuthorizationCurrencyCad PendingTransactionSourceCardAuthorizationCurrency = "CAD"
	PendingTransactionSourceCardAuthorizationCurrencyChf PendingTransactionSourceCardAuthorizationCurrency = "CHF"
	PendingTransactionSourceCardAuthorizationCurrencyEur PendingTransactionSourceCardAuthorizationCurrency = "EUR"
	PendingTransactionSourceCardAuthorizationCurrencyGbp PendingTransactionSourceCardAuthorizationCurrency = "GBP"
	PendingTransactionSourceCardAuthorizationCurrencyJpy PendingTransactionSourceCardAuthorizationCurrency = "JPY"
	PendingTransactionSourceCardAuthorizationCurrencyUsd PendingTransactionSourceCardAuthorizationCurrency = "USD"
)

type PendingTransactionSourceCheckDepositInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency fields.Field[PendingTransactionSourceCheckDepositInstructionCurrency] `json:"currency,required"`
	// The identifier of the File containing the image of the front of the check that
	// was deposited.
	FrontImageFileID fields.Field[string] `json:"front_image_file_id,required"`
	// The identifier of the File containing the image of the back of the check that
	// was deposited.
	BackImageFileID fields.Field[string] `json:"back_image_file_id,required,nullable"`
	// The identifier of the Check Deposit.
	CheckDepositID fields.Field[string] `json:"check_deposit_id,required,nullable"`
}

// MarshalJSON serializes PendingTransactionSourceCheckDepositInstruction into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *PendingTransactionSourceCheckDepositInstruction) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r PendingTransactionSourceCheckDepositInstruction) String() (result string) {
	return fmt.Sprintf("&PendingTransactionSourceCheckDepositInstruction{Amount:%s Currency:%s FrontImageFileID:%s BackImageFileID:%s CheckDepositID:%s}", r.Amount, r.Currency, r.FrontImageFileID, r.BackImageFileID, r.CheckDepositID)
}

type PendingTransactionSourceCheckDepositInstructionCurrency string

const (
	PendingTransactionSourceCheckDepositInstructionCurrencyCad PendingTransactionSourceCheckDepositInstructionCurrency = "CAD"
	PendingTransactionSourceCheckDepositInstructionCurrencyChf PendingTransactionSourceCheckDepositInstructionCurrency = "CHF"
	PendingTransactionSourceCheckDepositInstructionCurrencyEur PendingTransactionSourceCheckDepositInstructionCurrency = "EUR"
	PendingTransactionSourceCheckDepositInstructionCurrencyGbp PendingTransactionSourceCheckDepositInstructionCurrency = "GBP"
	PendingTransactionSourceCheckDepositInstructionCurrencyJpy PendingTransactionSourceCheckDepositInstructionCurrency = "JPY"
	PendingTransactionSourceCheckDepositInstructionCurrencyUsd PendingTransactionSourceCheckDepositInstructionCurrency = "USD"
)

type PendingTransactionSourceCheckTransferInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
	// currency.
	Currency fields.Field[PendingTransactionSourceCheckTransferInstructionCurrency] `json:"currency,required"`
	// The identifier of the Check Transfer that led to this Pending Transaction.
	TransferID fields.Field[string] `json:"transfer_id,required"`
}

// MarshalJSON serializes PendingTransactionSourceCheckTransferInstruction into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *PendingTransactionSourceCheckTransferInstruction) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r PendingTransactionSourceCheckTransferInstruction) String() (result string) {
	return fmt.Sprintf("&PendingTransactionSourceCheckTransferInstruction{Amount:%s Currency:%s TransferID:%s}", r.Amount, r.Currency, r.TransferID)
}

type PendingTransactionSourceCheckTransferInstructionCurrency string

const (
	PendingTransactionSourceCheckTransferInstructionCurrencyCad PendingTransactionSourceCheckTransferInstructionCurrency = "CAD"
	PendingTransactionSourceCheckTransferInstructionCurrencyChf PendingTransactionSourceCheckTransferInstructionCurrency = "CHF"
	PendingTransactionSourceCheckTransferInstructionCurrencyEur PendingTransactionSourceCheckTransferInstructionCurrency = "EUR"
	PendingTransactionSourceCheckTransferInstructionCurrencyGbp PendingTransactionSourceCheckTransferInstructionCurrency = "GBP"
	PendingTransactionSourceCheckTransferInstructionCurrencyJpy PendingTransactionSourceCheckTransferInstructionCurrency = "JPY"
	PendingTransactionSourceCheckTransferInstructionCurrencyUsd PendingTransactionSourceCheckTransferInstructionCurrency = "USD"
)

type PendingTransactionSourceInboundFundsHold struct {
	// The held amount in the minor unit of the account's currency. For dollars, for
	// example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the hold's
	// currency.
	Currency fields.Field[PendingTransactionSourceInboundFundsHoldCurrency] `json:"currency,required"`
	// When the hold will be released automatically. Certain conditions may cause it to
	// be released before this time.
	AutomaticallyReleasesAt fields.Field[time.Time] `json:"automatically_releases_at,required" format:"date-time"`
	// When the hold was released (if it has been released).
	ReleasedAt fields.Field[time.Time] `json:"released_at,required,nullable" format:"date-time"`
	// The status of the hold.
	Status fields.Field[PendingTransactionSourceInboundFundsHoldStatus] `json:"status,required"`
	// The ID of the Transaction for which funds were held.
	HeldTransactionID fields.Field[string] `json:"held_transaction_id,required,nullable"`
}

// MarshalJSON serializes PendingTransactionSourceInboundFundsHold into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *PendingTransactionSourceInboundFundsHold) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r PendingTransactionSourceInboundFundsHold) String() (result string) {
	return fmt.Sprintf("&PendingTransactionSourceInboundFundsHold{Amount:%s Currency:%s AutomaticallyReleasesAt:%s ReleasedAt:%s Status:%s HeldTransactionID:%s}", r.Amount, r.Currency, r.AutomaticallyReleasesAt, r.ReleasedAt, r.Status, r.HeldTransactionID)
}

type PendingTransactionSourceInboundFundsHoldCurrency string

const (
	PendingTransactionSourceInboundFundsHoldCurrencyCad PendingTransactionSourceInboundFundsHoldCurrency = "CAD"
	PendingTransactionSourceInboundFundsHoldCurrencyChf PendingTransactionSourceInboundFundsHoldCurrency = "CHF"
	PendingTransactionSourceInboundFundsHoldCurrencyEur PendingTransactionSourceInboundFundsHoldCurrency = "EUR"
	PendingTransactionSourceInboundFundsHoldCurrencyGbp PendingTransactionSourceInboundFundsHoldCurrency = "GBP"
	PendingTransactionSourceInboundFundsHoldCurrencyJpy PendingTransactionSourceInboundFundsHoldCurrency = "JPY"
	PendingTransactionSourceInboundFundsHoldCurrencyUsd PendingTransactionSourceInboundFundsHoldCurrency = "USD"
)

type PendingTransactionSourceInboundFundsHoldStatus string

const (
	PendingTransactionSourceInboundFundsHoldStatusHeld     PendingTransactionSourceInboundFundsHoldStatus = "held"
	PendingTransactionSourceInboundFundsHoldStatusComplete PendingTransactionSourceInboundFundsHoldStatus = "complete"
)

type PendingTransactionSourceCardRouteAuthorization struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency             fields.Field[PendingTransactionSourceCardRouteAuthorizationCurrency] `json:"currency,required"`
	MerchantAcceptorID   fields.Field[string]                                                 `json:"merchant_acceptor_id,required"`
	MerchantCity         fields.Field[string]                                                 `json:"merchant_city,required,nullable"`
	MerchantCountry      fields.Field[string]                                                 `json:"merchant_country,required"`
	MerchantDescriptor   fields.Field[string]                                                 `json:"merchant_descriptor,required"`
	MerchantCategoryCode fields.Field[string]                                                 `json:"merchant_category_code,required"`
	MerchantState        fields.Field[string]                                                 `json:"merchant_state,required,nullable"`
}

// MarshalJSON serializes PendingTransactionSourceCardRouteAuthorization into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *PendingTransactionSourceCardRouteAuthorization) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r PendingTransactionSourceCardRouteAuthorization) String() (result string) {
	return fmt.Sprintf("&PendingTransactionSourceCardRouteAuthorization{Amount:%s Currency:%s MerchantAcceptorID:%s MerchantCity:%s MerchantCountry:%s MerchantDescriptor:%s MerchantCategoryCode:%s MerchantState:%s}", r.Amount, r.Currency, r.MerchantAcceptorID, r.MerchantCity, r.MerchantCountry, r.MerchantDescriptor, r.MerchantCategoryCode, r.MerchantState)
}

type PendingTransactionSourceCardRouteAuthorizationCurrency string

const (
	PendingTransactionSourceCardRouteAuthorizationCurrencyCad PendingTransactionSourceCardRouteAuthorizationCurrency = "CAD"
	PendingTransactionSourceCardRouteAuthorizationCurrencyChf PendingTransactionSourceCardRouteAuthorizationCurrency = "CHF"
	PendingTransactionSourceCardRouteAuthorizationCurrencyEur PendingTransactionSourceCardRouteAuthorizationCurrency = "EUR"
	PendingTransactionSourceCardRouteAuthorizationCurrencyGbp PendingTransactionSourceCardRouteAuthorizationCurrency = "GBP"
	PendingTransactionSourceCardRouteAuthorizationCurrencyJpy PendingTransactionSourceCardRouteAuthorizationCurrency = "JPY"
	PendingTransactionSourceCardRouteAuthorizationCurrencyUsd PendingTransactionSourceCardRouteAuthorizationCurrency = "USD"
)

type PendingTransactionSourceWireDrawdownPaymentInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount             fields.Field[int64]  `json:"amount,required"`
	AccountNumber      fields.Field[string] `json:"account_number,required"`
	RoutingNumber      fields.Field[string] `json:"routing_number,required"`
	MessageToRecipient fields.Field[string] `json:"message_to_recipient,required"`
}

// MarshalJSON serializes PendingTransactionSourceWireDrawdownPaymentInstruction
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *PendingTransactionSourceWireDrawdownPaymentInstruction) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r PendingTransactionSourceWireDrawdownPaymentInstruction) String() (result string) {
	return fmt.Sprintf("&PendingTransactionSourceWireDrawdownPaymentInstruction{Amount:%s AccountNumber:%s RoutingNumber:%s MessageToRecipient:%s}", r.Amount, r.AccountNumber, r.RoutingNumber, r.MessageToRecipient)
}

type PendingTransactionSourceWireTransferInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount             fields.Field[int64]  `json:"amount,required"`
	AccountNumber      fields.Field[string] `json:"account_number,required"`
	RoutingNumber      fields.Field[string] `json:"routing_number,required"`
	MessageToRecipient fields.Field[string] `json:"message_to_recipient,required"`
	TransferID         fields.Field[string] `json:"transfer_id,required"`
}

// MarshalJSON serializes PendingTransactionSourceWireTransferInstruction into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *PendingTransactionSourceWireTransferInstruction) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r PendingTransactionSourceWireTransferInstruction) String() (result string) {
	return fmt.Sprintf("&PendingTransactionSourceWireTransferInstruction{Amount:%s AccountNumber:%s RoutingNumber:%s MessageToRecipient:%s TransferID:%s}", r.Amount, r.AccountNumber, r.RoutingNumber, r.MessageToRecipient, r.TransferID)
}

type PendingTransactionStatus string

const (
	PendingTransactionStatusPending  PendingTransactionStatus = "pending"
	PendingTransactionStatusComplete PendingTransactionStatus = "complete"
)

type PendingTransactionType string

const (
	PendingTransactionTypePendingTransaction PendingTransactionType = "pending_transaction"
)

type PendingTransactionListParams struct {
	// Return the page of entries after this one.
	Cursor fields.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit fields.Field[int64] `query:"limit"`
	// Filter pending transactions to those belonging to the specified Account.
	AccountID fields.Field[string] `query:"account_id"`
	// Filter pending transactions to those belonging to the specified Route.
	RouteID fields.Field[string] `query:"route_id"`
	// Filter pending transactions to those caused by the specified source.
	SourceID fields.Field[string]                             `query:"source_id"`
	Status   fields.Field[PendingTransactionListParamsStatus] `query:"status"`
}

// URLQuery serializes PendingTransactionListParams into a url.Values of the query
// parameters associated with this value
func (r *PendingTransactionListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r PendingTransactionListParams) String() (result string) {
	return fmt.Sprintf("&PendingTransactionListParams{Cursor:%s Limit:%s AccountID:%s RouteID:%s SourceID:%s Status:%s}", r.Cursor, r.Limit, r.AccountID, r.RouteID, r.SourceID, r.Status)
}

type PendingTransactionListParamsStatus struct {
	// Return results whose value is in the provided list. For GET requests, this
	// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
	In fields.Field[[]PendingTransactionListParamsStatusIn] `query:"in"`
}

// URLQuery serializes PendingTransactionListParamsStatus into a url.Values of the
// query parameters associated with this value
func (r *PendingTransactionListParamsStatus) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r PendingTransactionListParamsStatus) String() (result string) {
	return fmt.Sprintf("&PendingTransactionListParamsStatus{In:%s}", core.Fmt(r.In))
}

type PendingTransactionListParamsStatusIn string

const (
	PendingTransactionListParamsStatusInPending  PendingTransactionListParamsStatusIn = "pending"
	PendingTransactionListParamsStatusInComplete PendingTransactionListParamsStatusIn = "complete"
)

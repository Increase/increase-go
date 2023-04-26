package responses

import (
	"time"

	apijson "github.com/increase/increase-go/core/json"
)

type PendingTransaction struct {
	// The identifier for the account this Pending Transaction belongs to.
	AccountID string `json:"account_id,required"`
	// The Pending Transaction amount in the minor unit of its currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Pending
	// Transaction's currency. This will match the currency on the Pending
	// Transcation's Account.
	Currency PendingTransactionCurrency `json:"currency,required"`
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
	RouteType PendingTransactionRouteType `json:"route_type,required,nullable"`
	// This is an object giving more details on the network-level event that caused the
	// Pending Transaction. For example, for a card transaction this lists the
	// merchant's industry and location.
	Source PendingTransactionSource `json:"source,required"`
	// Whether the Pending Transaction has been confirmed and has an associated
	// Transaction.
	Status PendingTransactionStatus `json:"status,required"`
	// A constant representing the object's type. For this resource it will always be
	// `pending_transaction`.
	Type PendingTransactionType `json:"type,required"`
	JSON PendingTransactionJSON
}

type PendingTransactionJSON struct {
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

// UnmarshalJSON deserializes the provided bytes into PendingTransaction using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *PendingTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

type PendingTransactionRouteType string

const (
	PendingTransactionRouteTypeAccountNumber PendingTransactionRouteType = "account_number"
	PendingTransactionRouteTypeCard          PendingTransactionRouteType = "card"
)

type PendingTransactionSource struct {
	// The type of transaction that took place. We may add additional possible values
	// for this enum over time; your application should be able to handle such
	// additions gracefully.
	Category PendingTransactionSourceCategory `json:"category,required"`
	// A Account Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `account_transfer_instruction`.
	AccountTransferInstruction PendingTransactionSourceAccountTransferInstruction `json:"account_transfer_instruction,required,nullable"`
	// A ACH Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `ach_transfer_instruction`.
	ACHTransferInstruction PendingTransactionSourceACHTransferInstruction `json:"ach_transfer_instruction,required,nullable"`
	// A Card Authorization object. This field will be present in the JSON response if
	// and only if `category` is equal to `card_authorization`.
	CardAuthorization PendingTransactionSourceCardAuthorization `json:"card_authorization,required,nullable"`
	// A Check Deposit Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_deposit_instruction`.
	CheckDepositInstruction PendingTransactionSourceCheckDepositInstruction `json:"check_deposit_instruction,required,nullable"`
	// A Check Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_instruction`.
	CheckTransferInstruction PendingTransactionSourceCheckTransferInstruction `json:"check_transfer_instruction,required,nullable"`
	// A Inbound Funds Hold object. This field will be present in the JSON response if
	// and only if `category` is equal to `inbound_funds_hold`.
	InboundFundsHold PendingTransactionSourceInboundFundsHold `json:"inbound_funds_hold,required,nullable"`
	// A Deprecated Card Authorization object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_route_authorization`.
	CardRouteAuthorization PendingTransactionSourceCardRouteAuthorization `json:"card_route_authorization,required,nullable"`
	// A Real Time Payments Transfer Instruction object. This field will be present in
	// the JSON response if and only if `category` is equal to
	// `real_time_payments_transfer_instruction`.
	RealTimePaymentsTransferInstruction PendingTransactionSourceRealTimePaymentsTransferInstruction `json:"real_time_payments_transfer_instruction,required,nullable"`
	// A Wire Drawdown Payment Instruction object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `wire_drawdown_payment_instruction`.
	WireDrawdownPaymentInstruction PendingTransactionSourceWireDrawdownPaymentInstruction `json:"wire_drawdown_payment_instruction,required,nullable"`
	// A Wire Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_instruction`.
	WireTransferInstruction PendingTransactionSourceWireTransferInstruction `json:"wire_transfer_instruction,required,nullable"`
	JSON                    PendingTransactionSourceJSON
}

type PendingTransactionSourceJSON struct {
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

// UnmarshalJSON deserializes the provided bytes into PendingTransactionSource
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *PendingTransactionSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency PendingTransactionSourceAccountTransferInstructionCurrency `json:"currency,required"`
	// The identifier of the Account Transfer that led to this Pending Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       PendingTransactionSourceAccountTransferInstructionJSON
}

type PendingTransactionSourceAccountTransferInstructionJSON struct {
	Amount     apijson.Metadata
	Currency   apijson.Metadata
	TransferID apijson.Metadata
	Raw        []byte
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// PendingTransactionSourceAccountTransferInstruction using the internal json
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *PendingTransactionSourceAccountTransferInstruction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	Amount int64 `json:"amount,required"`
	// The identifier of the ACH Transfer that led to this Pending Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       PendingTransactionSourceACHTransferInstructionJSON
}

type PendingTransactionSourceACHTransferInstructionJSON struct {
	Amount     apijson.Metadata
	TransferID apijson.Metadata
	Raw        []byte
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// PendingTransactionSourceACHTransferInstruction using the internal json library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *PendingTransactionSourceACHTransferInstruction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PendingTransactionSourceCardAuthorization struct {
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
	Network PendingTransactionSourceCardAuthorizationNetwork `json:"network,required"`
	// Fields specific to the `network`
	NetworkDetails PendingTransactionSourceCardAuthorizationNetworkDetails `json:"network_details,required"`
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency PendingTransactionSourceCardAuthorizationCurrency `json:"currency,required"`
	// The identifier of the Real-Time Decision sent to approve or decline this
	// transaction.
	RealTimeDecisionID string `json:"real_time_decision_id,required,nullable"`
	// If the authorization was made via a Digital Wallet Token (such as an Apple Pay
	// purchase), the identifier of the token that was used.
	DigitalWalletTokenID string `json:"digital_wallet_token_id,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `card_authorization`.
	Type PendingTransactionSourceCardAuthorizationType `json:"type,required"`
	JSON PendingTransactionSourceCardAuthorizationJSON
}

type PendingTransactionSourceCardAuthorizationJSON struct {
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
// PendingTransactionSourceCardAuthorization using the internal json library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *PendingTransactionSourceCardAuthorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PendingTransactionSourceCardAuthorizationNetwork string

const (
	PendingTransactionSourceCardAuthorizationNetworkVisa PendingTransactionSourceCardAuthorizationNetwork = "visa"
)

type PendingTransactionSourceCardAuthorizationNetworkDetails struct {
	// Fields specific to the `visa` network
	Visa PendingTransactionSourceCardAuthorizationNetworkDetailsVisa `json:"visa,required"`
	JSON PendingTransactionSourceCardAuthorizationNetworkDetailsJSON
}

type PendingTransactionSourceCardAuthorizationNetworkDetailsJSON struct {
	Visa   apijson.Metadata
	Raw    []byte
	Extras map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// PendingTransactionSourceCardAuthorizationNetworkDetails using the internal json
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *PendingTransactionSourceCardAuthorizationNetworkDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PendingTransactionSourceCardAuthorizationNetworkDetailsVisa struct {
	// For electronic commerce transactions, this identifies the level of security used
	// in obtaining the customer's payment credential. For mail or telephone order
	// transactions, identifies the type of mail or telephone order.
	ElectronicCommerceIndicator PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator `json:"electronic_commerce_indicator,required,nullable"`
	// The method used to enter the cardholder's primary account number and card
	// expiration date
	PointOfServiceEntryMode PointOfServiceEntryMode `json:"point_of_service_entry_mode,required,nullable"`
	JSON                    PendingTransactionSourceCardAuthorizationNetworkDetailsVisaJSON
}

type PendingTransactionSourceCardAuthorizationNetworkDetailsVisaJSON struct {
	ElectronicCommerceIndicator apijson.Metadata
	PointOfServiceEntryMode     apijson.Metadata
	Raw                         []byte
	Extras                      map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// PendingTransactionSourceCardAuthorizationNetworkDetailsVisa using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *PendingTransactionSourceCardAuthorizationNetworkDetailsVisa) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

type PendingTransactionSourceCardAuthorizationType string

const (
	PendingTransactionSourceCardAuthorizationTypeCardAuthorization PendingTransactionSourceCardAuthorizationType = "card_authorization"
)

type PendingTransactionSourceCheckDepositInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency PendingTransactionSourceCheckDepositInstructionCurrency `json:"currency,required"`
	// The identifier of the File containing the image of the front of the check that
	// was deposited.
	FrontImageFileID string `json:"front_image_file_id,required"`
	// The identifier of the File containing the image of the back of the check that
	// was deposited.
	BackImageFileID string `json:"back_image_file_id,required,nullable"`
	// The identifier of the Check Deposit.
	CheckDepositID string `json:"check_deposit_id,required,nullable"`
	JSON           PendingTransactionSourceCheckDepositInstructionJSON
}

type PendingTransactionSourceCheckDepositInstructionJSON struct {
	Amount           apijson.Metadata
	Currency         apijson.Metadata
	FrontImageFileID apijson.Metadata
	BackImageFileID  apijson.Metadata
	CheckDepositID   apijson.Metadata
	Raw              []byte
	Extras           map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// PendingTransactionSourceCheckDepositInstruction using the internal json library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *PendingTransactionSourceCheckDepositInstruction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
	// currency.
	Currency PendingTransactionSourceCheckTransferInstructionCurrency `json:"currency,required"`
	// The identifier of the Check Transfer that led to this Pending Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       PendingTransactionSourceCheckTransferInstructionJSON
}

type PendingTransactionSourceCheckTransferInstructionJSON struct {
	Amount     apijson.Metadata
	Currency   apijson.Metadata
	TransferID apijson.Metadata
	Raw        []byte
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// PendingTransactionSourceCheckTransferInstruction using the internal json
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *PendingTransactionSourceCheckTransferInstruction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	Amount int64 `json:"amount,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the hold
	// was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the hold's
	// currency.
	Currency PendingTransactionSourceInboundFundsHoldCurrency `json:"currency,required"`
	// When the hold will be released automatically. Certain conditions may cause it to
	// be released before this time.
	AutomaticallyReleasesAt time.Time `json:"automatically_releases_at,required" format:"date-time"`
	// When the hold was released (if it has been released).
	ReleasedAt time.Time `json:"released_at,required,nullable" format:"date-time"`
	// The status of the hold.
	Status PendingTransactionSourceInboundFundsHoldStatus `json:"status,required"`
	// The ID of the Transaction for which funds were held.
	HeldTransactionID string `json:"held_transaction_id,required,nullable"`
	// The ID of the Pending Transaction representing the held funds.
	PendingTransactionID string `json:"pending_transaction_id,required,nullable"`
	JSON                 PendingTransactionSourceInboundFundsHoldJSON
}

type PendingTransactionSourceInboundFundsHoldJSON struct {
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
// PendingTransactionSourceInboundFundsHold using the internal json library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *PendingTransactionSourceInboundFundsHold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency             PendingTransactionSourceCardRouteAuthorizationCurrency `json:"currency,required"`
	MerchantAcceptorID   string                                                 `json:"merchant_acceptor_id,required"`
	MerchantCity         string                                                 `json:"merchant_city,required,nullable"`
	MerchantCountry      string                                                 `json:"merchant_country,required"`
	MerchantDescriptor   string                                                 `json:"merchant_descriptor,required"`
	MerchantCategoryCode string                                                 `json:"merchant_category_code,required"`
	MerchantState        string                                                 `json:"merchant_state,required,nullable"`
	JSON                 PendingTransactionSourceCardRouteAuthorizationJSON
}

type PendingTransactionSourceCardRouteAuthorizationJSON struct {
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
// PendingTransactionSourceCardRouteAuthorization using the internal json library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *PendingTransactionSourceCardRouteAuthorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

type PendingTransactionSourceRealTimePaymentsTransferInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The identifier of the Real Time Payments Transfer that led to this Pending
	// Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       PendingTransactionSourceRealTimePaymentsTransferInstructionJSON
}

type PendingTransactionSourceRealTimePaymentsTransferInstructionJSON struct {
	Amount     apijson.Metadata
	TransferID apijson.Metadata
	Raw        []byte
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// PendingTransactionSourceRealTimePaymentsTransferInstruction using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *PendingTransactionSourceRealTimePaymentsTransferInstruction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PendingTransactionSourceWireDrawdownPaymentInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount             int64  `json:"amount,required"`
	AccountNumber      string `json:"account_number,required"`
	RoutingNumber      string `json:"routing_number,required"`
	MessageToRecipient string `json:"message_to_recipient,required"`
	JSON               PendingTransactionSourceWireDrawdownPaymentInstructionJSON
}

type PendingTransactionSourceWireDrawdownPaymentInstructionJSON struct {
	Amount             apijson.Metadata
	AccountNumber      apijson.Metadata
	RoutingNumber      apijson.Metadata
	MessageToRecipient apijson.Metadata
	Raw                []byte
	Extras             map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// PendingTransactionSourceWireDrawdownPaymentInstruction using the internal json
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *PendingTransactionSourceWireDrawdownPaymentInstruction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PendingTransactionSourceWireTransferInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount             int64  `json:"amount,required"`
	AccountNumber      string `json:"account_number,required"`
	RoutingNumber      string `json:"routing_number,required"`
	MessageToRecipient string `json:"message_to_recipient,required"`
	TransferID         string `json:"transfer_id,required"`
	JSON               PendingTransactionSourceWireTransferInstructionJSON
}

type PendingTransactionSourceWireTransferInstructionJSON struct {
	Amount             apijson.Metadata
	AccountNumber      apijson.Metadata
	RoutingNumber      apijson.Metadata
	MessageToRecipient apijson.Metadata
	TransferID         apijson.Metadata
	Raw                []byte
	Extras             map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// PendingTransactionSourceWireTransferInstruction using the internal json library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *PendingTransactionSourceWireTransferInstruction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

type PendingTransactionListResponse struct {
	// The contents of the list.
	Data []PendingTransaction `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       PendingTransactionListResponseJSON
}

type PendingTransactionListResponseJSON struct {
	Data       apijson.Metadata
	NextCursor apijson.Metadata
	Raw        []byte
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// PendingTransactionListResponse using the internal json library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *PendingTransactionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

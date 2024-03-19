// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/apiquery"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/internal/shared"
	"github.com/increase/increase-go/option"
)

// PendingTransactionService contains methods and other services that help with
// interacting with the increase API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewPendingTransactionService] method
// instead.
type PendingTransactionService struct {
	Options []option.RequestOption
}

// NewPendingTransactionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPendingTransactionService(opts ...option.RequestOption) (r *PendingTransactionService) {
	r = &PendingTransactionService{}
	r.Options = opts
	return
}

// Retrieve a Pending Transaction
func (r *PendingTransactionService) Get(ctx context.Context, pendingTransactionID string, opts ...option.RequestOption) (res *PendingTransaction, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("pending_transactions/%s", pendingTransactionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Pending Transactions
func (r *PendingTransactionService) List(ctx context.Context, query PendingTransactionListParams, opts ...option.RequestOption) (res *shared.Page[PendingTransaction], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "pending_transactions"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List Pending Transactions
func (r *PendingTransactionService) ListAutoPaging(ctx context.Context, query PendingTransactionListParams, opts ...option.RequestOption) *shared.PageAutoPager[PendingTransaction] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Pending Transactions are potential future additions and removals of money from
// your bank account.
type PendingTransaction struct {
	// The Pending Transaction identifier.
	ID string `json:"id,required"`
	// The identifier for the account this Pending Transaction belongs to.
	AccountID string `json:"account_id,required"`
	// The Pending Transaction amount in the minor unit of its currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the Pending
	// Transaction was completed.
	CompletedAt time.Time `json:"completed_at,required,nullable" format:"date-time"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the Pending
	// Transaction occurred.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Pending
	// Transaction's currency. This will match the currency on the Pending
	// Transaction's Account.
	Currency PendingTransactionCurrency `json:"currency,required"`
	// For a Pending Transaction related to a transfer, this is the description you
	// provide. For a Pending Transaction related to a payment, this is the description
	// the vendor provides.
	Description string `json:"description,required"`
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
	JSON pendingTransactionJSON `json:"-"`
}

// pendingTransactionJSON contains the JSON metadata for the struct
// [PendingTransaction]
type pendingTransactionJSON struct {
	ID          apijson.Field
	AccountID   apijson.Field
	Amount      apijson.Field
	CompletedAt apijson.Field
	CreatedAt   apijson.Field
	Currency    apijson.Field
	Description apijson.Field
	RouteID     apijson.Field
	RouteType   apijson.Field
	Source      apijson.Field
	Status      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PendingTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pendingTransactionJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Pending
// Transaction's currency. This will match the currency on the Pending
// Transaction's Account.
type PendingTransactionCurrency string

const (
	// Canadian Dollar (CAD)
	PendingTransactionCurrencyCad PendingTransactionCurrency = "CAD"
	// Swiss Franc (CHF)
	PendingTransactionCurrencyChf PendingTransactionCurrency = "CHF"
	// Euro (EUR)
	PendingTransactionCurrencyEur PendingTransactionCurrency = "EUR"
	// British Pound (GBP)
	PendingTransactionCurrencyGbp PendingTransactionCurrency = "GBP"
	// Japanese Yen (JPY)
	PendingTransactionCurrencyJpy PendingTransactionCurrency = "JPY"
	// US Dollar (USD)
	PendingTransactionCurrencyUsd PendingTransactionCurrency = "USD"
)

// The type of the route this Pending Transaction came through.
type PendingTransactionRouteType string

const (
	// An Account Number.
	PendingTransactionRouteTypeAccountNumber PendingTransactionRouteType = "account_number"
	// A Card.
	PendingTransactionRouteTypeCard PendingTransactionRouteType = "card"
)

// This is an object giving more details on the network-level event that caused the
// Pending Transaction. For example, for a card transaction this lists the
// merchant's industry and location.
type PendingTransactionSource struct {
	// An Account Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `account_transfer_instruction`.
	AccountTransferInstruction PendingTransactionSourceAccountTransferInstruction `json:"account_transfer_instruction,required,nullable"`
	// An ACH Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `ach_transfer_instruction`.
	ACHTransferInstruction PendingTransactionSourceACHTransferInstruction `json:"ach_transfer_instruction,required,nullable"`
	// A Card Authorization object. This field will be present in the JSON response if
	// and only if `category` is equal to `card_authorization`.
	CardAuthorization PendingTransactionSourceCardAuthorization `json:"card_authorization,required,nullable"`
	// The type of the resource. We may add additional possible values for this enum
	// over time; your application should be able to handle such additions gracefully.
	Category PendingTransactionSourceCategory `json:"category,required"`
	// A Check Deposit Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_deposit_instruction`.
	CheckDepositInstruction PendingTransactionSourceCheckDepositInstruction `json:"check_deposit_instruction,required,nullable"`
	// A Check Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_instruction`.
	CheckTransferInstruction PendingTransactionSourceCheckTransferInstruction `json:"check_transfer_instruction,required,nullable"`
	// An Inbound Funds Hold object. This field will be present in the JSON response if
	// and only if `category` is equal to `inbound_funds_hold`.
	InboundFundsHold PendingTransactionSourceInboundFundsHold `json:"inbound_funds_hold,required,nullable"`
	// A Real-Time Payments Transfer Instruction object. This field will be present in
	// the JSON response if and only if `category` is equal to
	// `real_time_payments_transfer_instruction`.
	RealTimePaymentsTransferInstruction PendingTransactionSourceRealTimePaymentsTransferInstruction `json:"real_time_payments_transfer_instruction,required,nullable"`
	// A Wire Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_instruction`.
	WireTransferInstruction PendingTransactionSourceWireTransferInstruction `json:"wire_transfer_instruction,required,nullable"`
	JSON                    pendingTransactionSourceJSON                    `json:"-"`
}

// pendingTransactionSourceJSON contains the JSON metadata for the struct
// [PendingTransactionSource]
type pendingTransactionSourceJSON struct {
	AccountTransferInstruction          apijson.Field
	ACHTransferInstruction              apijson.Field
	CardAuthorization                   apijson.Field
	Category                            apijson.Field
	CheckDepositInstruction             apijson.Field
	CheckTransferInstruction            apijson.Field
	InboundFundsHold                    apijson.Field
	RealTimePaymentsTransferInstruction apijson.Field
	WireTransferInstruction             apijson.Field
	raw                                 string
	ExtraFields                         map[string]apijson.Field
}

func (r *PendingTransactionSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pendingTransactionSourceJSON) RawJSON() string {
	return r.raw
}

// An Account Transfer Instruction object. This field will be present in the JSON
// response if and only if `category` is equal to `account_transfer_instruction`.
type PendingTransactionSourceAccountTransferInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency PendingTransactionSourceAccountTransferInstructionCurrency `json:"currency,required"`
	// The identifier of the Account Transfer that led to this Pending Transaction.
	TransferID string                                                 `json:"transfer_id,required"`
	JSON       pendingTransactionSourceAccountTransferInstructionJSON `json:"-"`
}

// pendingTransactionSourceAccountTransferInstructionJSON contains the JSON
// metadata for the struct [PendingTransactionSourceAccountTransferInstruction]
type pendingTransactionSourceAccountTransferInstructionJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	TransferID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PendingTransactionSourceAccountTransferInstruction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pendingTransactionSourceAccountTransferInstructionJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
// account currency.
type PendingTransactionSourceAccountTransferInstructionCurrency string

const (
	// Canadian Dollar (CAD)
	PendingTransactionSourceAccountTransferInstructionCurrencyCad PendingTransactionSourceAccountTransferInstructionCurrency = "CAD"
	// Swiss Franc (CHF)
	PendingTransactionSourceAccountTransferInstructionCurrencyChf PendingTransactionSourceAccountTransferInstructionCurrency = "CHF"
	// Euro (EUR)
	PendingTransactionSourceAccountTransferInstructionCurrencyEur PendingTransactionSourceAccountTransferInstructionCurrency = "EUR"
	// British Pound (GBP)
	PendingTransactionSourceAccountTransferInstructionCurrencyGbp PendingTransactionSourceAccountTransferInstructionCurrency = "GBP"
	// Japanese Yen (JPY)
	PendingTransactionSourceAccountTransferInstructionCurrencyJpy PendingTransactionSourceAccountTransferInstructionCurrency = "JPY"
	// US Dollar (USD)
	PendingTransactionSourceAccountTransferInstructionCurrencyUsd PendingTransactionSourceAccountTransferInstructionCurrency = "USD"
)

// An ACH Transfer Instruction object. This field will be present in the JSON
// response if and only if `category` is equal to `ach_transfer_instruction`.
type PendingTransactionSourceACHTransferInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The identifier of the ACH Transfer that led to this Pending Transaction.
	TransferID string                                             `json:"transfer_id,required"`
	JSON       pendingTransactionSourceACHTransferInstructionJSON `json:"-"`
}

// pendingTransactionSourceACHTransferInstructionJSON contains the JSON metadata
// for the struct [PendingTransactionSourceACHTransferInstruction]
type pendingTransactionSourceACHTransferInstructionJSON struct {
	Amount      apijson.Field
	TransferID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PendingTransactionSourceACHTransferInstruction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pendingTransactionSourceACHTransferInstructionJSON) RawJSON() string {
	return r.raw
}

// A Card Authorization object. This field will be present in the JSON response if
// and only if `category` is equal to `card_authorization`.
type PendingTransactionSourceCardAuthorization struct {
	// The Card Authorization identifier.
	ID string `json:"id,required"`
	// Whether this authorization was approved by Increase, the card network through
	// stand-in processing, or the user through a real-time decision.
	Actioner PendingTransactionSourceCardAuthorizationActioner `json:"actioner,required"`
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The ID of the Card Payment this transaction belongs to.
	CardPaymentID string `json:"card_payment_id,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency PendingTransactionSourceCardAuthorizationCurrency `json:"currency,required"`
	// If the authorization was made via a Digital Wallet Token (such as an Apple Pay
	// purchase), the identifier of the token that was used.
	DigitalWalletTokenID string `json:"digital_wallet_token_id,required,nullable"`
	// The direction descibes the direction the funds will move, either from the
	// cardholder to the merchant or from the merchant to the cardholder.
	Direction PendingTransactionSourceCardAuthorizationDirection `json:"direction,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) when this authorization
	// will expire and the pending transaction will be released.
	ExpiresAt time.Time `json:"expires_at,required" format:"date-time"`
	// The merchant identifier (commonly abbreviated as MID) of the merchant the card
	// is transacting with.
	MerchantAcceptorID string `json:"merchant_acceptor_id,required"`
	// The Merchant Category Code (commonly abbreviated as MCC) of the merchant the
	// card is transacting with.
	MerchantCategoryCode string `json:"merchant_category_code,required,nullable"`
	// The city the merchant resides in.
	MerchantCity string `json:"merchant_city,required,nullable"`
	// The country the merchant resides in.
	MerchantCountry string `json:"merchant_country,required,nullable"`
	// The merchant descriptor of the merchant the card is transacting with.
	MerchantDescriptor string `json:"merchant_descriptor,required"`
	// Fields specific to the `network`.
	NetworkDetails PendingTransactionSourceCardAuthorizationNetworkDetails `json:"network_details,required"`
	// Network-specific identifiers for a specific request or transaction.
	NetworkIdentifiers PendingTransactionSourceCardAuthorizationNetworkIdentifiers `json:"network_identifiers,required"`
	// The identifier of the Pending Transaction associated with this Transaction.
	PendingTransactionID string `json:"pending_transaction_id,required,nullable"`
	// If the authorization was made in-person with a physical card, the Physical Card
	// that was used.
	PhysicalCardID string `json:"physical_card_id,required,nullable"`
	// The processing category describes the intent behind the authorization, such as
	// whether it was used for bill payments or an automatic fuel dispenser.
	ProcessingCategory PendingTransactionSourceCardAuthorizationProcessingCategory `json:"processing_category,required"`
	// The identifier of the Real-Time Decision sent to approve or decline this
	// transaction.
	RealTimeDecisionID string `json:"real_time_decision_id,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `card_authorization`.
	Type PendingTransactionSourceCardAuthorizationType `json:"type,required"`
	// Fields related to verification of cardholder-provided values.
	Verification PendingTransactionSourceCardAuthorizationVerification `json:"verification,required"`
	JSON         pendingTransactionSourceCardAuthorizationJSON         `json:"-"`
}

// pendingTransactionSourceCardAuthorizationJSON contains the JSON metadata for the
// struct [PendingTransactionSourceCardAuthorization]
type pendingTransactionSourceCardAuthorizationJSON struct {
	ID                   apijson.Field
	Actioner             apijson.Field
	Amount               apijson.Field
	CardPaymentID        apijson.Field
	Currency             apijson.Field
	DigitalWalletTokenID apijson.Field
	Direction            apijson.Field
	ExpiresAt            apijson.Field
	MerchantAcceptorID   apijson.Field
	MerchantCategoryCode apijson.Field
	MerchantCity         apijson.Field
	MerchantCountry      apijson.Field
	MerchantDescriptor   apijson.Field
	NetworkDetails       apijson.Field
	NetworkIdentifiers   apijson.Field
	PendingTransactionID apijson.Field
	PhysicalCardID       apijson.Field
	ProcessingCategory   apijson.Field
	RealTimeDecisionID   apijson.Field
	Type                 apijson.Field
	Verification         apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *PendingTransactionSourceCardAuthorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pendingTransactionSourceCardAuthorizationJSON) RawJSON() string {
	return r.raw
}

// Whether this authorization was approved by Increase, the card network through
// stand-in processing, or the user through a real-time decision.
type PendingTransactionSourceCardAuthorizationActioner string

const (
	// This object was actioned by the user through a real-time decision.
	PendingTransactionSourceCardAuthorizationActionerUser PendingTransactionSourceCardAuthorizationActioner = "user"
	// This object was actioned by Increase without user intervention.
	PendingTransactionSourceCardAuthorizationActionerIncrease PendingTransactionSourceCardAuthorizationActioner = "increase"
	// This object was actioned by the network, through stand-in processing.
	PendingTransactionSourceCardAuthorizationActionerNetwork PendingTransactionSourceCardAuthorizationActioner = "network"
)

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
type PendingTransactionSourceCardAuthorizationCurrency string

const (
	// Canadian Dollar (CAD)
	PendingTransactionSourceCardAuthorizationCurrencyCad PendingTransactionSourceCardAuthorizationCurrency = "CAD"
	// Swiss Franc (CHF)
	PendingTransactionSourceCardAuthorizationCurrencyChf PendingTransactionSourceCardAuthorizationCurrency = "CHF"
	// Euro (EUR)
	PendingTransactionSourceCardAuthorizationCurrencyEur PendingTransactionSourceCardAuthorizationCurrency = "EUR"
	// British Pound (GBP)
	PendingTransactionSourceCardAuthorizationCurrencyGbp PendingTransactionSourceCardAuthorizationCurrency = "GBP"
	// Japanese Yen (JPY)
	PendingTransactionSourceCardAuthorizationCurrencyJpy PendingTransactionSourceCardAuthorizationCurrency = "JPY"
	// US Dollar (USD)
	PendingTransactionSourceCardAuthorizationCurrencyUsd PendingTransactionSourceCardAuthorizationCurrency = "USD"
)

// The direction descibes the direction the funds will move, either from the
// cardholder to the merchant or from the merchant to the cardholder.
type PendingTransactionSourceCardAuthorizationDirection string

const (
	// A regular card authorization where funds are debited from the cardholder.
	PendingTransactionSourceCardAuthorizationDirectionSettlement PendingTransactionSourceCardAuthorizationDirection = "settlement"
	// A refund card authorization, sometimes referred to as a credit voucher
	// authorization, where funds are credited to the cardholder.
	PendingTransactionSourceCardAuthorizationDirectionRefund PendingTransactionSourceCardAuthorizationDirection = "refund"
)

// Fields specific to the `network`.
type PendingTransactionSourceCardAuthorizationNetworkDetails struct {
	// The payment network used to process this card authorization.
	Category PendingTransactionSourceCardAuthorizationNetworkDetailsCategory `json:"category,required"`
	// Fields specific to the `visa` network.
	Visa PendingTransactionSourceCardAuthorizationNetworkDetailsVisa `json:"visa,required,nullable"`
	JSON pendingTransactionSourceCardAuthorizationNetworkDetailsJSON `json:"-"`
}

// pendingTransactionSourceCardAuthorizationNetworkDetailsJSON contains the JSON
// metadata for the struct
// [PendingTransactionSourceCardAuthorizationNetworkDetails]
type pendingTransactionSourceCardAuthorizationNetworkDetailsJSON struct {
	Category    apijson.Field
	Visa        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PendingTransactionSourceCardAuthorizationNetworkDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pendingTransactionSourceCardAuthorizationNetworkDetailsJSON) RawJSON() string {
	return r.raw
}

// The payment network used to process this card authorization.
type PendingTransactionSourceCardAuthorizationNetworkDetailsCategory string

const (
	// Visa
	PendingTransactionSourceCardAuthorizationNetworkDetailsCategoryVisa PendingTransactionSourceCardAuthorizationNetworkDetailsCategory = "visa"
)

// Fields specific to the `visa` network.
type PendingTransactionSourceCardAuthorizationNetworkDetailsVisa struct {
	// For electronic commerce transactions, this identifies the level of security used
	// in obtaining the customer's payment credential. For mail or telephone order
	// transactions, identifies the type of mail or telephone order.
	ElectronicCommerceIndicator PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator `json:"electronic_commerce_indicator,required,nullable"`
	// The method used to enter the cardholder's primary account number and card
	// expiration date.
	PointOfServiceEntryMode PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode `json:"point_of_service_entry_mode,required,nullable"`
	JSON                    pendingTransactionSourceCardAuthorizationNetworkDetailsVisaJSON                    `json:"-"`
}

// pendingTransactionSourceCardAuthorizationNetworkDetailsVisaJSON contains the
// JSON metadata for the struct
// [PendingTransactionSourceCardAuthorizationNetworkDetailsVisa]
type pendingTransactionSourceCardAuthorizationNetworkDetailsVisaJSON struct {
	ElectronicCommerceIndicator apijson.Field
	PointOfServiceEntryMode     apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *PendingTransactionSourceCardAuthorizationNetworkDetailsVisa) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pendingTransactionSourceCardAuthorizationNetworkDetailsVisaJSON) RawJSON() string {
	return r.raw
}

// For electronic commerce transactions, this identifies the level of security used
// in obtaining the customer's payment credential. For mail or telephone order
// transactions, identifies the type of mail or telephone order.
type PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator string

const (
	// Single transaction of a mail/phone order: Use to indicate that the transaction
	// is a mail/phone order purchase, not a recurring transaction or installment
	// payment. For domestic transactions in the US region, this value may also
	// indicate one bill payment transaction in the card-present or card-absent
	// environments.
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorMailPhoneOrder PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "mail_phone_order"
	// Recurring transaction: Payment indicator used to indicate a recurring
	// transaction that originates from an acquirer in the US region.
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorRecurring PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "recurring"
	// Installment payment: Payment indicator used to indicate one purchase of goods or
	// services that is billed to the account in multiple charges over a period of time
	// agreed upon by the cardholder and merchant from transactions that originate from
	// an acquirer in the US region.
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorInstallment PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "installment"
	// Unknown classification: other mail order: Use to indicate that the type of
	// mail/telephone order is unknown.
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorUnknownMailPhoneOrder PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "unknown_mail_phone_order"
	// Secure electronic commerce transaction: Use to indicate that the electronic
	// commerce transaction has been authenticated using e.g., 3-D Secure
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorSecureElectronicCommerce PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "secure_electronic_commerce"
	// Non-authenticated security transaction at a 3-D Secure-capable merchant, and
	// merchant attempted to authenticate the cardholder using 3-D Secure: Use to
	// identify an electronic commerce transaction where the merchant attempted to
	// authenticate the cardholder using 3-D Secure, but was unable to complete the
	// authentication because the issuer or cardholder does not participate in the 3-D
	// Secure program.
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransactionAt3DSCapableMerchant PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction_at_3ds_capable_merchant"
	// Non-authenticated security transaction: Use to identify an electronic commerce
	// transaction that uses data encryption for security however , cardholder
	// authentication is not performed using 3-D Secure.
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransaction PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction"
	// Non-secure transaction: Use to identify an electronic commerce transaction that
	// has no data protection.
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonSecureTransaction PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "non_secure_transaction"
)

// The method used to enter the cardholder's primary account number and card
// expiration date.
type PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode string

const (
	// Unknown
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeUnknown PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "unknown"
	// Manual key entry
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeManual PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "manual"
	// Magnetic stripe read, without card verification value
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeMagneticStripeNoCvv PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "magnetic_stripe_no_cvv"
	// Optical code
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeOpticalCode PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "optical_code"
	// Contact chip card
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCard PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "integrated_circuit_card"
	// Contactless read of chip card
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeContactless PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "contactless"
	// Transaction initiated using a credential that has previously been stored on file
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeCredentialOnFile PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "credential_on_file"
	// Magnetic stripe read
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeMagneticStripe PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "magnetic_stripe"
	// Contactless read of magnetic stripe data
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeContactlessMagneticStripe PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "contactless_magnetic_stripe"
	// Contact chip card, without card verification value
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCardNoCvv PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "integrated_circuit_card_no_cvv"
)

// Network-specific identifiers for a specific request or transaction.
type PendingTransactionSourceCardAuthorizationNetworkIdentifiers struct {
	// A life-cycle identifier used across e.g., an authorization and a reversal.
	// Expected to be unique per acquirer within a window of time. For some card
	// networks the retrieval reference number includes the trace counter.
	RetrievalReferenceNumber string `json:"retrieval_reference_number,required,nullable"`
	// A counter used to verify an individual authorization. Expected to be unique per
	// acquirer within a window of time.
	TraceNumber string `json:"trace_number,required,nullable"`
	// A globally unique transaction identifier provided by the card network, used
	// across multiple life-cycle requests.
	TransactionID string                                                          `json:"transaction_id,required,nullable"`
	JSON          pendingTransactionSourceCardAuthorizationNetworkIdentifiersJSON `json:"-"`
}

// pendingTransactionSourceCardAuthorizationNetworkIdentifiersJSON contains the
// JSON metadata for the struct
// [PendingTransactionSourceCardAuthorizationNetworkIdentifiers]
type pendingTransactionSourceCardAuthorizationNetworkIdentifiersJSON struct {
	RetrievalReferenceNumber apijson.Field
	TraceNumber              apijson.Field
	TransactionID            apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *PendingTransactionSourceCardAuthorizationNetworkIdentifiers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pendingTransactionSourceCardAuthorizationNetworkIdentifiersJSON) RawJSON() string {
	return r.raw
}

// The processing category describes the intent behind the authorization, such as
// whether it was used for bill payments or an automatic fuel dispenser.
type PendingTransactionSourceCardAuthorizationProcessingCategory string

const (
	// Account funding transactions are transactions used to e.g., fund an account or
	// transfer funds between accounts.
	PendingTransactionSourceCardAuthorizationProcessingCategoryAccountFunding PendingTransactionSourceCardAuthorizationProcessingCategory = "account_funding"
	// Automatic fuel dispenser authorizations occur when a card is used at a gas pump,
	// prior to the actual transaction amount being known. They are followed by an
	// advice message that updates the amount of the pending transaction.
	PendingTransactionSourceCardAuthorizationProcessingCategoryAutomaticFuelDispenser PendingTransactionSourceCardAuthorizationProcessingCategory = "automatic_fuel_dispenser"
	// A transaction used to pay a bill.
	PendingTransactionSourceCardAuthorizationProcessingCategoryBillPayment PendingTransactionSourceCardAuthorizationProcessingCategory = "bill_payment"
	// A regular purchase.
	PendingTransactionSourceCardAuthorizationProcessingCategoryPurchase PendingTransactionSourceCardAuthorizationProcessingCategory = "purchase"
	// Quasi-cash transactions represent purchases of items which may be convertible to
	// cash.
	PendingTransactionSourceCardAuthorizationProcessingCategoryQuasiCash PendingTransactionSourceCardAuthorizationProcessingCategory = "quasi_cash"
	// A refund card authorization, sometimes referred to as a credit voucher
	// authorization, where funds are credited to the cardholder.
	PendingTransactionSourceCardAuthorizationProcessingCategoryRefund PendingTransactionSourceCardAuthorizationProcessingCategory = "refund"
)

// A constant representing the object's type. For this resource it will always be
// `card_authorization`.
type PendingTransactionSourceCardAuthorizationType string

const (
	PendingTransactionSourceCardAuthorizationTypeCardAuthorization PendingTransactionSourceCardAuthorizationType = "card_authorization"
)

// Fields related to verification of cardholder-provided values.
type PendingTransactionSourceCardAuthorizationVerification struct {
	// Fields related to verification of the Card Verification Code, a 3-digit code on
	// the back of the card.
	CardVerificationCode PendingTransactionSourceCardAuthorizationVerificationCardVerificationCode `json:"card_verification_code,required"`
	// Cardholder address provided in the authorization request and the address on file
	// we verified it against.
	CardholderAddress PendingTransactionSourceCardAuthorizationVerificationCardholderAddress `json:"cardholder_address,required"`
	JSON              pendingTransactionSourceCardAuthorizationVerificationJSON              `json:"-"`
}

// pendingTransactionSourceCardAuthorizationVerificationJSON contains the JSON
// metadata for the struct [PendingTransactionSourceCardAuthorizationVerification]
type pendingTransactionSourceCardAuthorizationVerificationJSON struct {
	CardVerificationCode apijson.Field
	CardholderAddress    apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *PendingTransactionSourceCardAuthorizationVerification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pendingTransactionSourceCardAuthorizationVerificationJSON) RawJSON() string {
	return r.raw
}

// Fields related to verification of the Card Verification Code, a 3-digit code on
// the back of the card.
type PendingTransactionSourceCardAuthorizationVerificationCardVerificationCode struct {
	// The result of verifying the Card Verification Code.
	Result PendingTransactionSourceCardAuthorizationVerificationCardVerificationCodeResult `json:"result,required"`
	JSON   pendingTransactionSourceCardAuthorizationVerificationCardVerificationCodeJSON   `json:"-"`
}

// pendingTransactionSourceCardAuthorizationVerificationCardVerificationCodeJSON
// contains the JSON metadata for the struct
// [PendingTransactionSourceCardAuthorizationVerificationCardVerificationCode]
type pendingTransactionSourceCardAuthorizationVerificationCardVerificationCodeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PendingTransactionSourceCardAuthorizationVerificationCardVerificationCode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pendingTransactionSourceCardAuthorizationVerificationCardVerificationCodeJSON) RawJSON() string {
	return r.raw
}

// The result of verifying the Card Verification Code.
type PendingTransactionSourceCardAuthorizationVerificationCardVerificationCodeResult string

const (
	// No card verification code was provided in the authorization request.
	PendingTransactionSourceCardAuthorizationVerificationCardVerificationCodeResultNotChecked PendingTransactionSourceCardAuthorizationVerificationCardVerificationCodeResult = "not_checked"
	// The card verification code matched the one on file.
	PendingTransactionSourceCardAuthorizationVerificationCardVerificationCodeResultMatch PendingTransactionSourceCardAuthorizationVerificationCardVerificationCodeResult = "match"
	// The card verification code did not match the one on file.
	PendingTransactionSourceCardAuthorizationVerificationCardVerificationCodeResultNoMatch PendingTransactionSourceCardAuthorizationVerificationCardVerificationCodeResult = "no_match"
)

// Cardholder address provided in the authorization request and the address on file
// we verified it against.
type PendingTransactionSourceCardAuthorizationVerificationCardholderAddress struct {
	// Line 1 of the address on file for the cardholder.
	ActualLine1 string `json:"actual_line1,required,nullable"`
	// The postal code of the address on file for the cardholder.
	ActualPostalCode string `json:"actual_postal_code,required,nullable"`
	// The cardholder address line 1 provided for verification in the authorization
	// request.
	ProvidedLine1 string `json:"provided_line1,required,nullable"`
	// The postal code provided for verification in the authorization request.
	ProvidedPostalCode string `json:"provided_postal_code,required,nullable"`
	// The address verification result returned to the card network.
	Result PendingTransactionSourceCardAuthorizationVerificationCardholderAddressResult `json:"result,required"`
	JSON   pendingTransactionSourceCardAuthorizationVerificationCardholderAddressJSON   `json:"-"`
}

// pendingTransactionSourceCardAuthorizationVerificationCardholderAddressJSON
// contains the JSON metadata for the struct
// [PendingTransactionSourceCardAuthorizationVerificationCardholderAddress]
type pendingTransactionSourceCardAuthorizationVerificationCardholderAddressJSON struct {
	ActualLine1        apijson.Field
	ActualPostalCode   apijson.Field
	ProvidedLine1      apijson.Field
	ProvidedPostalCode apijson.Field
	Result             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PendingTransactionSourceCardAuthorizationVerificationCardholderAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pendingTransactionSourceCardAuthorizationVerificationCardholderAddressJSON) RawJSON() string {
	return r.raw
}

// The address verification result returned to the card network.
type PendingTransactionSourceCardAuthorizationVerificationCardholderAddressResult string

const (
	// No adress was provided in the authorization request.
	PendingTransactionSourceCardAuthorizationVerificationCardholderAddressResultNotChecked PendingTransactionSourceCardAuthorizationVerificationCardholderAddressResult = "not_checked"
	// Postal code matches, but the street address was not verified.
	PendingTransactionSourceCardAuthorizationVerificationCardholderAddressResultPostalCodeMatchAddressNotChecked PendingTransactionSourceCardAuthorizationVerificationCardholderAddressResult = "postal_code_match_address_not_checked"
	// Postal code matches, but the street address does not match.
	PendingTransactionSourceCardAuthorizationVerificationCardholderAddressResultPostalCodeMatchAddressNoMatch PendingTransactionSourceCardAuthorizationVerificationCardholderAddressResult = "postal_code_match_address_no_match"
	// Postal code does not match, but the street address matches.
	PendingTransactionSourceCardAuthorizationVerificationCardholderAddressResultPostalCodeNoMatchAddressMatch PendingTransactionSourceCardAuthorizationVerificationCardholderAddressResult = "postal_code_no_match_address_match"
	// Postal code and street address match.
	PendingTransactionSourceCardAuthorizationVerificationCardholderAddressResultMatch PendingTransactionSourceCardAuthorizationVerificationCardholderAddressResult = "match"
	// Postal code and street address do not match.
	PendingTransactionSourceCardAuthorizationVerificationCardholderAddressResultNoMatch PendingTransactionSourceCardAuthorizationVerificationCardholderAddressResult = "no_match"
)

// The type of the resource. We may add additional possible values for this enum
// over time; your application should be able to handle such additions gracefully.
type PendingTransactionSourceCategory string

const (
	// Account Transfer Instruction: details will be under the
	// `account_transfer_instruction` object.
	PendingTransactionSourceCategoryAccountTransferInstruction PendingTransactionSourceCategory = "account_transfer_instruction"
	// ACH Transfer Instruction: details will be under the `ach_transfer_instruction`
	// object.
	PendingTransactionSourceCategoryACHTransferInstruction PendingTransactionSourceCategory = "ach_transfer_instruction"
	// Card Authorization: details will be under the `card_authorization` object.
	PendingTransactionSourceCategoryCardAuthorization PendingTransactionSourceCategory = "card_authorization"
	// Check Deposit Instruction: details will be under the `check_deposit_instruction`
	// object.
	PendingTransactionSourceCategoryCheckDepositInstruction PendingTransactionSourceCategory = "check_deposit_instruction"
	// Check Transfer Instruction: details will be under the
	// `check_transfer_instruction` object.
	PendingTransactionSourceCategoryCheckTransferInstruction PendingTransactionSourceCategory = "check_transfer_instruction"
	// Inbound Funds Hold: details will be under the `inbound_funds_hold` object.
	PendingTransactionSourceCategoryInboundFundsHold PendingTransactionSourceCategory = "inbound_funds_hold"
	// Real-Time Payments Transfer Instruction: details will be under the
	// `real_time_payments_transfer_instruction` object.
	PendingTransactionSourceCategoryRealTimePaymentsTransferInstruction PendingTransactionSourceCategory = "real_time_payments_transfer_instruction"
	// Wire Transfer Instruction: details will be under the `wire_transfer_instruction`
	// object.
	PendingTransactionSourceCategoryWireTransferInstruction PendingTransactionSourceCategory = "wire_transfer_instruction"
	// The Pending Transaction was made for an undocumented or deprecated reason.
	PendingTransactionSourceCategoryOther PendingTransactionSourceCategory = "other"
)

// A Check Deposit Instruction object. This field will be present in the JSON
// response if and only if `category` is equal to `check_deposit_instruction`.
type PendingTransactionSourceCheckDepositInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The identifier of the File containing the image of the back of the check that
	// was deposited.
	BackImageFileID string `json:"back_image_file_id,required,nullable"`
	// The identifier of the Check Deposit.
	CheckDepositID string `json:"check_deposit_id,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency PendingTransactionSourceCheckDepositInstructionCurrency `json:"currency,required"`
	// The identifier of the File containing the image of the front of the check that
	// was deposited.
	FrontImageFileID string                                              `json:"front_image_file_id,required"`
	JSON             pendingTransactionSourceCheckDepositInstructionJSON `json:"-"`
}

// pendingTransactionSourceCheckDepositInstructionJSON contains the JSON metadata
// for the struct [PendingTransactionSourceCheckDepositInstruction]
type pendingTransactionSourceCheckDepositInstructionJSON struct {
	Amount           apijson.Field
	BackImageFileID  apijson.Field
	CheckDepositID   apijson.Field
	Currency         apijson.Field
	FrontImageFileID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *PendingTransactionSourceCheckDepositInstruction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pendingTransactionSourceCheckDepositInstructionJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
type PendingTransactionSourceCheckDepositInstructionCurrency string

const (
	// Canadian Dollar (CAD)
	PendingTransactionSourceCheckDepositInstructionCurrencyCad PendingTransactionSourceCheckDepositInstructionCurrency = "CAD"
	// Swiss Franc (CHF)
	PendingTransactionSourceCheckDepositInstructionCurrencyChf PendingTransactionSourceCheckDepositInstructionCurrency = "CHF"
	// Euro (EUR)
	PendingTransactionSourceCheckDepositInstructionCurrencyEur PendingTransactionSourceCheckDepositInstructionCurrency = "EUR"
	// British Pound (GBP)
	PendingTransactionSourceCheckDepositInstructionCurrencyGbp PendingTransactionSourceCheckDepositInstructionCurrency = "GBP"
	// Japanese Yen (JPY)
	PendingTransactionSourceCheckDepositInstructionCurrencyJpy PendingTransactionSourceCheckDepositInstructionCurrency = "JPY"
	// US Dollar (USD)
	PendingTransactionSourceCheckDepositInstructionCurrencyUsd PendingTransactionSourceCheckDepositInstructionCurrency = "USD"
)

// A Check Transfer Instruction object. This field will be present in the JSON
// response if and only if `category` is equal to `check_transfer_instruction`.
type PendingTransactionSourceCheckTransferInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
	// currency.
	Currency PendingTransactionSourceCheckTransferInstructionCurrency `json:"currency,required"`
	// The identifier of the Check Transfer that led to this Pending Transaction.
	TransferID string                                               `json:"transfer_id,required"`
	JSON       pendingTransactionSourceCheckTransferInstructionJSON `json:"-"`
}

// pendingTransactionSourceCheckTransferInstructionJSON contains the JSON metadata
// for the struct [PendingTransactionSourceCheckTransferInstruction]
type pendingTransactionSourceCheckTransferInstructionJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	TransferID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PendingTransactionSourceCheckTransferInstruction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pendingTransactionSourceCheckTransferInstructionJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
// currency.
type PendingTransactionSourceCheckTransferInstructionCurrency string

const (
	// Canadian Dollar (CAD)
	PendingTransactionSourceCheckTransferInstructionCurrencyCad PendingTransactionSourceCheckTransferInstructionCurrency = "CAD"
	// Swiss Franc (CHF)
	PendingTransactionSourceCheckTransferInstructionCurrencyChf PendingTransactionSourceCheckTransferInstructionCurrency = "CHF"
	// Euro (EUR)
	PendingTransactionSourceCheckTransferInstructionCurrencyEur PendingTransactionSourceCheckTransferInstructionCurrency = "EUR"
	// British Pound (GBP)
	PendingTransactionSourceCheckTransferInstructionCurrencyGbp PendingTransactionSourceCheckTransferInstructionCurrency = "GBP"
	// Japanese Yen (JPY)
	PendingTransactionSourceCheckTransferInstructionCurrencyJpy PendingTransactionSourceCheckTransferInstructionCurrency = "JPY"
	// US Dollar (USD)
	PendingTransactionSourceCheckTransferInstructionCurrencyUsd PendingTransactionSourceCheckTransferInstructionCurrency = "USD"
)

// An Inbound Funds Hold object. This field will be present in the JSON response if
// and only if `category` is equal to `inbound_funds_hold`.
type PendingTransactionSourceInboundFundsHold struct {
	// The Inbound Funds Hold identifier.
	ID string `json:"id,required"`
	// The held amount in the minor unit of the account's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// When the hold will be released automatically. Certain conditions may cause it to
	// be released before this time.
	AutomaticallyReleasesAt time.Time `json:"automatically_releases_at,required" format:"date-time"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the hold
	// was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the hold's
	// currency.
	Currency PendingTransactionSourceInboundFundsHoldCurrency `json:"currency,required"`
	// The ID of the Transaction for which funds were held.
	HeldTransactionID string `json:"held_transaction_id,required,nullable"`
	// The ID of the Pending Transaction representing the held funds.
	PendingTransactionID string `json:"pending_transaction_id,required,nullable"`
	// When the hold was released (if it has been released).
	ReleasedAt time.Time `json:"released_at,required,nullable" format:"date-time"`
	// The status of the hold.
	Status PendingTransactionSourceInboundFundsHoldStatus `json:"status,required"`
	// A constant representing the object's type. For this resource it will always be
	// `inbound_funds_hold`.
	Type PendingTransactionSourceInboundFundsHoldType `json:"type,required"`
	JSON pendingTransactionSourceInboundFundsHoldJSON `json:"-"`
}

// pendingTransactionSourceInboundFundsHoldJSON contains the JSON metadata for the
// struct [PendingTransactionSourceInboundFundsHold]
type pendingTransactionSourceInboundFundsHoldJSON struct {
	ID                      apijson.Field
	Amount                  apijson.Field
	AutomaticallyReleasesAt apijson.Field
	CreatedAt               apijson.Field
	Currency                apijson.Field
	HeldTransactionID       apijson.Field
	PendingTransactionID    apijson.Field
	ReleasedAt              apijson.Field
	Status                  apijson.Field
	Type                    apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *PendingTransactionSourceInboundFundsHold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pendingTransactionSourceInboundFundsHoldJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the hold's
// currency.
type PendingTransactionSourceInboundFundsHoldCurrency string

const (
	// Canadian Dollar (CAD)
	PendingTransactionSourceInboundFundsHoldCurrencyCad PendingTransactionSourceInboundFundsHoldCurrency = "CAD"
	// Swiss Franc (CHF)
	PendingTransactionSourceInboundFundsHoldCurrencyChf PendingTransactionSourceInboundFundsHoldCurrency = "CHF"
	// Euro (EUR)
	PendingTransactionSourceInboundFundsHoldCurrencyEur PendingTransactionSourceInboundFundsHoldCurrency = "EUR"
	// British Pound (GBP)
	PendingTransactionSourceInboundFundsHoldCurrencyGbp PendingTransactionSourceInboundFundsHoldCurrency = "GBP"
	// Japanese Yen (JPY)
	PendingTransactionSourceInboundFundsHoldCurrencyJpy PendingTransactionSourceInboundFundsHoldCurrency = "JPY"
	// US Dollar (USD)
	PendingTransactionSourceInboundFundsHoldCurrencyUsd PendingTransactionSourceInboundFundsHoldCurrency = "USD"
)

// The status of the hold.
type PendingTransactionSourceInboundFundsHoldStatus string

const (
	// Funds are still being held.
	PendingTransactionSourceInboundFundsHoldStatusHeld PendingTransactionSourceInboundFundsHoldStatus = "held"
	// Funds have been released.
	PendingTransactionSourceInboundFundsHoldStatusComplete PendingTransactionSourceInboundFundsHoldStatus = "complete"
)

// A constant representing the object's type. For this resource it will always be
// `inbound_funds_hold`.
type PendingTransactionSourceInboundFundsHoldType string

const (
	PendingTransactionSourceInboundFundsHoldTypeInboundFundsHold PendingTransactionSourceInboundFundsHoldType = "inbound_funds_hold"
)

// A Real-Time Payments Transfer Instruction object. This field will be present in
// the JSON response if and only if `category` is equal to
// `real_time_payments_transfer_instruction`.
type PendingTransactionSourceRealTimePaymentsTransferInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The identifier of the Real-Time Payments Transfer that led to this Pending
	// Transaction.
	TransferID string                                                          `json:"transfer_id,required"`
	JSON       pendingTransactionSourceRealTimePaymentsTransferInstructionJSON `json:"-"`
}

// pendingTransactionSourceRealTimePaymentsTransferInstructionJSON contains the
// JSON metadata for the struct
// [PendingTransactionSourceRealTimePaymentsTransferInstruction]
type pendingTransactionSourceRealTimePaymentsTransferInstructionJSON struct {
	Amount      apijson.Field
	TransferID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PendingTransactionSourceRealTimePaymentsTransferInstruction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pendingTransactionSourceRealTimePaymentsTransferInstructionJSON) RawJSON() string {
	return r.raw
}

// A Wire Transfer Instruction object. This field will be present in the JSON
// response if and only if `category` is equal to `wire_transfer_instruction`.
type PendingTransactionSourceWireTransferInstruction struct {
	// The account number for the destination account.
	AccountNumber string `json:"account_number,required"`
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The message that will show on the recipient's bank statement.
	MessageToRecipient string `json:"message_to_recipient,required"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
	// destination account.
	RoutingNumber string `json:"routing_number,required"`
	// The identifier of the Wire Transfer that led to this Pending Transaction.
	TransferID string                                              `json:"transfer_id,required"`
	JSON       pendingTransactionSourceWireTransferInstructionJSON `json:"-"`
}

// pendingTransactionSourceWireTransferInstructionJSON contains the JSON metadata
// for the struct [PendingTransactionSourceWireTransferInstruction]
type pendingTransactionSourceWireTransferInstructionJSON struct {
	AccountNumber      apijson.Field
	Amount             apijson.Field
	MessageToRecipient apijson.Field
	RoutingNumber      apijson.Field
	TransferID         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PendingTransactionSourceWireTransferInstruction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pendingTransactionSourceWireTransferInstructionJSON) RawJSON() string {
	return r.raw
}

// Whether the Pending Transaction has been confirmed and has an associated
// Transaction.
type PendingTransactionStatus string

const (
	// The Pending Transaction is still awaiting confirmation.
	PendingTransactionStatusPending PendingTransactionStatus = "pending"
	// The Pending Transaction is confirmed. An associated Transaction exists for this
	// object. The Pending Transaction will no longer count against your balance and
	// can generally be hidden from UIs, etc.
	PendingTransactionStatusComplete PendingTransactionStatus = "complete"
)

// A constant representing the object's type. For this resource it will always be
// `pending_transaction`.
type PendingTransactionType string

const (
	PendingTransactionTypePendingTransaction PendingTransactionType = "pending_transaction"
)

type PendingTransactionListParams struct {
	// Filter pending transactions to those belonging to the specified Account.
	AccountID param.Field[string]                                `query:"account_id"`
	Category  param.Field[PendingTransactionListParamsCategory]  `query:"category"`
	CreatedAt param.Field[PendingTransactionListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
	// Filter pending transactions to those belonging to the specified Route.
	RouteID param.Field[string] `query:"route_id"`
	// Filter pending transactions to those caused by the specified source.
	SourceID param.Field[string]                             `query:"source_id"`
	Status   param.Field[PendingTransactionListParamsStatus] `query:"status"`
}

// URLQuery serializes [PendingTransactionListParams]'s query parameters as
// `url.Values`.
func (r PendingTransactionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type PendingTransactionListParamsCategory struct {
	// Return results whose value is in the provided list. For GET requests, this
	// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
	In param.Field[[]PendingTransactionListParamsCategoryIn] `query:"in"`
}

// URLQuery serializes [PendingTransactionListParamsCategory]'s query parameters as
// `url.Values`.
func (r PendingTransactionListParamsCategory) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type PendingTransactionListParamsCategoryIn string

const (
	// Account Transfer Instruction: details will be under the
	// `account_transfer_instruction` object.
	PendingTransactionListParamsCategoryInAccountTransferInstruction PendingTransactionListParamsCategoryIn = "account_transfer_instruction"
	// ACH Transfer Instruction: details will be under the `ach_transfer_instruction`
	// object.
	PendingTransactionListParamsCategoryInACHTransferInstruction PendingTransactionListParamsCategoryIn = "ach_transfer_instruction"
	// Card Authorization: details will be under the `card_authorization` object.
	PendingTransactionListParamsCategoryInCardAuthorization PendingTransactionListParamsCategoryIn = "card_authorization"
	// Check Deposit Instruction: details will be under the `check_deposit_instruction`
	// object.
	PendingTransactionListParamsCategoryInCheckDepositInstruction PendingTransactionListParamsCategoryIn = "check_deposit_instruction"
	// Check Transfer Instruction: details will be under the
	// `check_transfer_instruction` object.
	PendingTransactionListParamsCategoryInCheckTransferInstruction PendingTransactionListParamsCategoryIn = "check_transfer_instruction"
	// Inbound Funds Hold: details will be under the `inbound_funds_hold` object.
	PendingTransactionListParamsCategoryInInboundFundsHold PendingTransactionListParamsCategoryIn = "inbound_funds_hold"
	// Real-Time Payments Transfer Instruction: details will be under the
	// `real_time_payments_transfer_instruction` object.
	PendingTransactionListParamsCategoryInRealTimePaymentsTransferInstruction PendingTransactionListParamsCategoryIn = "real_time_payments_transfer_instruction"
	// Wire Transfer Instruction: details will be under the `wire_transfer_instruction`
	// object.
	PendingTransactionListParamsCategoryInWireTransferInstruction PendingTransactionListParamsCategoryIn = "wire_transfer_instruction"
	// The Pending Transaction was made for an undocumented or deprecated reason.
	PendingTransactionListParamsCategoryInOther PendingTransactionListParamsCategoryIn = "other"
)

type PendingTransactionListParamsCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After param.Field[time.Time] `query:"after" format:"date-time"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before param.Field[time.Time] `query:"before" format:"date-time"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter param.Field[time.Time] `query:"on_or_after" format:"date-time"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore param.Field[time.Time] `query:"on_or_before" format:"date-time"`
}

// URLQuery serializes [PendingTransactionListParamsCreatedAt]'s query parameters
// as `url.Values`.
func (r PendingTransactionListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type PendingTransactionListParamsStatus struct {
	// Filter Pending Transactions for those with the specified status. By default only
	// Pending Transactions in with status `pending` will be returned. For GET
	// requests, this should be encoded as a comma-delimited string, such as
	// `?in=one,two,three`.
	In param.Field[[]PendingTransactionListParamsStatusIn] `query:"in"`
}

// URLQuery serializes [PendingTransactionListParamsStatus]'s query parameters as
// `url.Values`.
func (r PendingTransactionListParamsStatus) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type PendingTransactionListParamsStatusIn string

const (
	// The Pending Transaction is still awaiting confirmation.
	PendingTransactionListParamsStatusInPending PendingTransactionListParamsStatusIn = "pending"
	// The Pending Transaction is confirmed. An associated Transaction exists for this
	// object. The Pending Transaction will no longer count against your balance and
	// can generally be hidden from UIs, etc.
	PendingTransactionListParamsStatusInComplete PendingTransactionListParamsStatusIn = "complete"
)

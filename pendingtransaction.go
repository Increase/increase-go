// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Increase/increase-go/internal/apijson"
	"github.com/Increase/increase-go/internal/apiquery"
	"github.com/Increase/increase-go/internal/param"
	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
	"github.com/Increase/increase-go/packages/pagination"
)

// PendingTransactionService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPendingTransactionService] method instead.
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
	if pendingTransactionID == "" {
		err = errors.New("missing required pending_transaction_id parameter")
		return
	}
	path := fmt.Sprintf("pending_transactions/%s", pendingTransactionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Pending Transactions
func (r *PendingTransactionService) List(ctx context.Context, query PendingTransactionListParams, opts ...option.RequestOption) (res *pagination.Page[PendingTransaction], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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
func (r *PendingTransactionService) ListAutoPaging(ctx context.Context, query PendingTransactionListParams, opts ...option.RequestOption) *pagination.PageAutoPager[PendingTransaction] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
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
	PendingTransactionCurrencyCad PendingTransactionCurrency = "CAD"
	PendingTransactionCurrencyChf PendingTransactionCurrency = "CHF"
	PendingTransactionCurrencyEur PendingTransactionCurrency = "EUR"
	PendingTransactionCurrencyGbp PendingTransactionCurrency = "GBP"
	PendingTransactionCurrencyJpy PendingTransactionCurrency = "JPY"
	PendingTransactionCurrencyUsd PendingTransactionCurrency = "USD"
)

func (r PendingTransactionCurrency) IsKnown() bool {
	switch r {
	case PendingTransactionCurrencyCad, PendingTransactionCurrencyChf, PendingTransactionCurrencyEur, PendingTransactionCurrencyGbp, PendingTransactionCurrencyJpy, PendingTransactionCurrencyUsd:
		return true
	}
	return false
}

// The type of the route this Pending Transaction came through.
type PendingTransactionRouteType string

const (
	PendingTransactionRouteTypeAccountNumber PendingTransactionRouteType = "account_number"
	PendingTransactionRouteTypeCard          PendingTransactionRouteType = "card"
	PendingTransactionRouteTypeLockbox       PendingTransactionRouteType = "lockbox"
)

func (r PendingTransactionRouteType) IsKnown() bool {
	switch r {
	case PendingTransactionRouteTypeAccountNumber, PendingTransactionRouteTypeCard, PendingTransactionRouteTypeLockbox:
		return true
	}
	return false
}

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
	// An Inbound Wire Transfer Reversal object. This field will be present in the JSON
	// response if and only if `category` is equal to `inbound_wire_transfer_reversal`.
	InboundWireTransferReversal PendingTransactionSourceInboundWireTransferReversal `json:"inbound_wire_transfer_reversal,required,nullable"`
	// If the category of this Transaction source is equal to `other`, this field will
	// contain an empty object, otherwise it will contain null.
	Other interface{} `json:"other,required,nullable"`
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
	InboundWireTransferReversal         apijson.Field
	Other                               apijson.Field
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
	PendingTransactionSourceAccountTransferInstructionCurrencyCad PendingTransactionSourceAccountTransferInstructionCurrency = "CAD"
	PendingTransactionSourceAccountTransferInstructionCurrencyChf PendingTransactionSourceAccountTransferInstructionCurrency = "CHF"
	PendingTransactionSourceAccountTransferInstructionCurrencyEur PendingTransactionSourceAccountTransferInstructionCurrency = "EUR"
	PendingTransactionSourceAccountTransferInstructionCurrencyGbp PendingTransactionSourceAccountTransferInstructionCurrency = "GBP"
	PendingTransactionSourceAccountTransferInstructionCurrencyJpy PendingTransactionSourceAccountTransferInstructionCurrency = "JPY"
	PendingTransactionSourceAccountTransferInstructionCurrencyUsd PendingTransactionSourceAccountTransferInstructionCurrency = "USD"
)

func (r PendingTransactionSourceAccountTransferInstructionCurrency) IsKnown() bool {
	switch r {
	case PendingTransactionSourceAccountTransferInstructionCurrencyCad, PendingTransactionSourceAccountTransferInstructionCurrencyChf, PendingTransactionSourceAccountTransferInstructionCurrencyEur, PendingTransactionSourceAccountTransferInstructionCurrencyGbp, PendingTransactionSourceAccountTransferInstructionCurrencyJpy, PendingTransactionSourceAccountTransferInstructionCurrencyUsd:
		return true
	}
	return false
}

// An ACH Transfer Instruction object. This field will be present in the JSON
// response if and only if `category` is equal to `ach_transfer_instruction`.
type PendingTransactionSourceACHTransferInstruction struct {
	// The pending amount in USD cents.
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
	CardPaymentID string `json:"card_payment_id,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency PendingTransactionSourceCardAuthorizationCurrency `json:"currency,required"`
	// If the authorization was made via a Digital Wallet Token (such as an Apple Pay
	// purchase), the identifier of the token that was used.
	DigitalWalletTokenID string `json:"digital_wallet_token_id,required,nullable"`
	// The direction describes the direction the funds will move, either from the
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
	MerchantCategoryCode string `json:"merchant_category_code,required"`
	// The city the merchant resides in.
	MerchantCity string `json:"merchant_city,required,nullable"`
	// The country the merchant resides in.
	MerchantCountry string `json:"merchant_country,required"`
	// The merchant descriptor of the merchant the card is transacting with.
	MerchantDescriptor string `json:"merchant_descriptor,required"`
	// The merchant's postal code. For US merchants this is either a 5-digit or 9-digit
	// ZIP code, where the first 5 and last 4 are separated by a dash.
	MerchantPostalCode string `json:"merchant_postal_code,required,nullable"`
	// The state the merchant resides in.
	MerchantState string `json:"merchant_state,required,nullable"`
	// Fields specific to the `network`.
	NetworkDetails PendingTransactionSourceCardAuthorizationNetworkDetails `json:"network_details,required"`
	// Network-specific identifiers for a specific request or transaction.
	NetworkIdentifiers PendingTransactionSourceCardAuthorizationNetworkIdentifiers `json:"network_identifiers,required"`
	// The risk score generated by the card network. For Visa this is the Visa Advanced
	// Authorization risk score, from 0 to 99, where 99 is the riskiest.
	NetworkRiskScore int64 `json:"network_risk_score,required,nullable"`
	// The identifier of the Pending Transaction associated with this Transaction.
	PendingTransactionID string `json:"pending_transaction_id,required,nullable"`
	// If the authorization was made in-person with a physical card, the Physical Card
	// that was used.
	PhysicalCardID string `json:"physical_card_id,required,nullable"`
	// The pending amount in the minor unit of the transaction's presentment currency.
	PresentmentAmount int64 `json:"presentment_amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's presentment currency.
	PresentmentCurrency string `json:"presentment_currency,required"`
	// The processing category describes the intent behind the authorization, such as
	// whether it was used for bill payments or an automatic fuel dispenser.
	ProcessingCategory PendingTransactionSourceCardAuthorizationProcessingCategory `json:"processing_category,required"`
	// The identifier of the Real-Time Decision sent to approve or decline this
	// transaction.
	RealTimeDecisionID string `json:"real_time_decision_id,required,nullable"`
	// The terminal identifier (commonly abbreviated as TID) of the terminal the card
	// is transacting with.
	TerminalID string `json:"terminal_id,required,nullable"`
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
	MerchantPostalCode   apijson.Field
	MerchantState        apijson.Field
	NetworkDetails       apijson.Field
	NetworkIdentifiers   apijson.Field
	NetworkRiskScore     apijson.Field
	PendingTransactionID apijson.Field
	PhysicalCardID       apijson.Field
	PresentmentAmount    apijson.Field
	PresentmentCurrency  apijson.Field
	ProcessingCategory   apijson.Field
	RealTimeDecisionID   apijson.Field
	TerminalID           apijson.Field
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
	PendingTransactionSourceCardAuthorizationActionerUser     PendingTransactionSourceCardAuthorizationActioner = "user"
	PendingTransactionSourceCardAuthorizationActionerIncrease PendingTransactionSourceCardAuthorizationActioner = "increase"
	PendingTransactionSourceCardAuthorizationActionerNetwork  PendingTransactionSourceCardAuthorizationActioner = "network"
)

func (r PendingTransactionSourceCardAuthorizationActioner) IsKnown() bool {
	switch r {
	case PendingTransactionSourceCardAuthorizationActionerUser, PendingTransactionSourceCardAuthorizationActionerIncrease, PendingTransactionSourceCardAuthorizationActionerNetwork:
		return true
	}
	return false
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
type PendingTransactionSourceCardAuthorizationCurrency string

const (
	PendingTransactionSourceCardAuthorizationCurrencyCad PendingTransactionSourceCardAuthorizationCurrency = "CAD"
	PendingTransactionSourceCardAuthorizationCurrencyChf PendingTransactionSourceCardAuthorizationCurrency = "CHF"
	PendingTransactionSourceCardAuthorizationCurrencyEur PendingTransactionSourceCardAuthorizationCurrency = "EUR"
	PendingTransactionSourceCardAuthorizationCurrencyGbp PendingTransactionSourceCardAuthorizationCurrency = "GBP"
	PendingTransactionSourceCardAuthorizationCurrencyJpy PendingTransactionSourceCardAuthorizationCurrency = "JPY"
	PendingTransactionSourceCardAuthorizationCurrencyUsd PendingTransactionSourceCardAuthorizationCurrency = "USD"
)

func (r PendingTransactionSourceCardAuthorizationCurrency) IsKnown() bool {
	switch r {
	case PendingTransactionSourceCardAuthorizationCurrencyCad, PendingTransactionSourceCardAuthorizationCurrencyChf, PendingTransactionSourceCardAuthorizationCurrencyEur, PendingTransactionSourceCardAuthorizationCurrencyGbp, PendingTransactionSourceCardAuthorizationCurrencyJpy, PendingTransactionSourceCardAuthorizationCurrencyUsd:
		return true
	}
	return false
}

// The direction describes the direction the funds will move, either from the
// cardholder to the merchant or from the merchant to the cardholder.
type PendingTransactionSourceCardAuthorizationDirection string

const (
	PendingTransactionSourceCardAuthorizationDirectionSettlement PendingTransactionSourceCardAuthorizationDirection = "settlement"
	PendingTransactionSourceCardAuthorizationDirectionRefund     PendingTransactionSourceCardAuthorizationDirection = "refund"
)

func (r PendingTransactionSourceCardAuthorizationDirection) IsKnown() bool {
	switch r {
	case PendingTransactionSourceCardAuthorizationDirectionSettlement, PendingTransactionSourceCardAuthorizationDirectionRefund:
		return true
	}
	return false
}

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
	PendingTransactionSourceCardAuthorizationNetworkDetailsCategoryVisa PendingTransactionSourceCardAuthorizationNetworkDetailsCategory = "visa"
)

func (r PendingTransactionSourceCardAuthorizationNetworkDetailsCategory) IsKnown() bool {
	switch r {
	case PendingTransactionSourceCardAuthorizationNetworkDetailsCategoryVisa:
		return true
	}
	return false
}

// Fields specific to the `visa` network.
type PendingTransactionSourceCardAuthorizationNetworkDetailsVisa struct {
	// For electronic commerce transactions, this identifies the level of security used
	// in obtaining the customer's payment credential. For mail or telephone order
	// transactions, identifies the type of mail or telephone order.
	ElectronicCommerceIndicator PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator `json:"electronic_commerce_indicator,required,nullable"`
	// The method used to enter the cardholder's primary account number and card
	// expiration date.
	PointOfServiceEntryMode PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode `json:"point_of_service_entry_mode,required,nullable"`
	// Only present when `actioner: network`. Describes why a card authorization was
	// approved or declined by Visa through stand-in processing.
	StandInProcessingReason PendingTransactionSourceCardAuthorizationNetworkDetailsVisaStandInProcessingReason `json:"stand_in_processing_reason,required,nullable"`
	JSON                    pendingTransactionSourceCardAuthorizationNetworkDetailsVisaJSON                    `json:"-"`
}

// pendingTransactionSourceCardAuthorizationNetworkDetailsVisaJSON contains the
// JSON metadata for the struct
// [PendingTransactionSourceCardAuthorizationNetworkDetailsVisa]
type pendingTransactionSourceCardAuthorizationNetworkDetailsVisaJSON struct {
	ElectronicCommerceIndicator apijson.Field
	PointOfServiceEntryMode     apijson.Field
	StandInProcessingReason     apijson.Field
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
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorMailPhoneOrder                                          PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "mail_phone_order"
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorRecurring                                               PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "recurring"
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorInstallment                                             PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "installment"
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorUnknownMailPhoneOrder                                   PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "unknown_mail_phone_order"
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorSecureElectronicCommerce                                PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "secure_electronic_commerce"
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransactionAt3DSCapableMerchant PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction_at_3ds_capable_merchant"
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransaction                     PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction"
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonSecureTransaction                                    PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "non_secure_transaction"
)

func (r PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator) IsKnown() bool {
	switch r {
	case PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorMailPhoneOrder, PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorRecurring, PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorInstallment, PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorUnknownMailPhoneOrder, PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorSecureElectronicCommerce, PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransactionAt3DSCapableMerchant, PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransaction, PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonSecureTransaction:
		return true
	}
	return false
}

// The method used to enter the cardholder's primary account number and card
// expiration date.
type PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode string

const (
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeUnknown                    PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "unknown"
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeManual                     PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "manual"
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeMagneticStripeNoCvv        PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "magnetic_stripe_no_cvv"
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeOpticalCode                PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "optical_code"
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCard      PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "integrated_circuit_card"
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeContactless                PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "contactless"
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeCredentialOnFile           PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "credential_on_file"
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeMagneticStripe             PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "magnetic_stripe"
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeContactlessMagneticStripe  PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "contactless_magnetic_stripe"
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCardNoCvv PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "integrated_circuit_card_no_cvv"
)

func (r PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode) IsKnown() bool {
	switch r {
	case PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeUnknown, PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeManual, PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeMagneticStripeNoCvv, PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeOpticalCode, PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCard, PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeContactless, PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeCredentialOnFile, PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeMagneticStripe, PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeContactlessMagneticStripe, PendingTransactionSourceCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCardNoCvv:
		return true
	}
	return false
}

// Only present when `actioner: network`. Describes why a card authorization was
// approved or declined by Visa through stand-in processing.
type PendingTransactionSourceCardAuthorizationNetworkDetailsVisaStandInProcessingReason string

const (
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaStandInProcessingReasonIssuerError                                              PendingTransactionSourceCardAuthorizationNetworkDetailsVisaStandInProcessingReason = "issuer_error"
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaStandInProcessingReasonInvalidPhysicalCard                                      PendingTransactionSourceCardAuthorizationNetworkDetailsVisaStandInProcessingReason = "invalid_physical_card"
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaStandInProcessingReasonInvalidCardholderAuthenticationVerificationValue         PendingTransactionSourceCardAuthorizationNetworkDetailsVisaStandInProcessingReason = "invalid_cardholder_authentication_verification_value"
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaStandInProcessingReasonInternalVisaError                                        PendingTransactionSourceCardAuthorizationNetworkDetailsVisaStandInProcessingReason = "internal_visa_error"
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaStandInProcessingReasonMerchantTransactionAdvisoryServiceAuthenticationRequired PendingTransactionSourceCardAuthorizationNetworkDetailsVisaStandInProcessingReason = "merchant_transaction_advisory_service_authentication_required"
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaStandInProcessingReasonPaymentFraudDisruptionAcquirerBlock                      PendingTransactionSourceCardAuthorizationNetworkDetailsVisaStandInProcessingReason = "payment_fraud_disruption_acquirer_block"
	PendingTransactionSourceCardAuthorizationNetworkDetailsVisaStandInProcessingReasonOther                                                    PendingTransactionSourceCardAuthorizationNetworkDetailsVisaStandInProcessingReason = "other"
)

func (r PendingTransactionSourceCardAuthorizationNetworkDetailsVisaStandInProcessingReason) IsKnown() bool {
	switch r {
	case PendingTransactionSourceCardAuthorizationNetworkDetailsVisaStandInProcessingReasonIssuerError, PendingTransactionSourceCardAuthorizationNetworkDetailsVisaStandInProcessingReasonInvalidPhysicalCard, PendingTransactionSourceCardAuthorizationNetworkDetailsVisaStandInProcessingReasonInvalidCardholderAuthenticationVerificationValue, PendingTransactionSourceCardAuthorizationNetworkDetailsVisaStandInProcessingReasonInternalVisaError, PendingTransactionSourceCardAuthorizationNetworkDetailsVisaStandInProcessingReasonMerchantTransactionAdvisoryServiceAuthenticationRequired, PendingTransactionSourceCardAuthorizationNetworkDetailsVisaStandInProcessingReasonPaymentFraudDisruptionAcquirerBlock, PendingTransactionSourceCardAuthorizationNetworkDetailsVisaStandInProcessingReasonOther:
		return true
	}
	return false
}

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
	PendingTransactionSourceCardAuthorizationProcessingCategoryAccountFunding         PendingTransactionSourceCardAuthorizationProcessingCategory = "account_funding"
	PendingTransactionSourceCardAuthorizationProcessingCategoryAutomaticFuelDispenser PendingTransactionSourceCardAuthorizationProcessingCategory = "automatic_fuel_dispenser"
	PendingTransactionSourceCardAuthorizationProcessingCategoryBillPayment            PendingTransactionSourceCardAuthorizationProcessingCategory = "bill_payment"
	PendingTransactionSourceCardAuthorizationProcessingCategoryPurchase               PendingTransactionSourceCardAuthorizationProcessingCategory = "purchase"
	PendingTransactionSourceCardAuthorizationProcessingCategoryQuasiCash              PendingTransactionSourceCardAuthorizationProcessingCategory = "quasi_cash"
	PendingTransactionSourceCardAuthorizationProcessingCategoryRefund                 PendingTransactionSourceCardAuthorizationProcessingCategory = "refund"
)

func (r PendingTransactionSourceCardAuthorizationProcessingCategory) IsKnown() bool {
	switch r {
	case PendingTransactionSourceCardAuthorizationProcessingCategoryAccountFunding, PendingTransactionSourceCardAuthorizationProcessingCategoryAutomaticFuelDispenser, PendingTransactionSourceCardAuthorizationProcessingCategoryBillPayment, PendingTransactionSourceCardAuthorizationProcessingCategoryPurchase, PendingTransactionSourceCardAuthorizationProcessingCategoryQuasiCash, PendingTransactionSourceCardAuthorizationProcessingCategoryRefund:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `card_authorization`.
type PendingTransactionSourceCardAuthorizationType string

const (
	PendingTransactionSourceCardAuthorizationTypeCardAuthorization PendingTransactionSourceCardAuthorizationType = "card_authorization"
)

func (r PendingTransactionSourceCardAuthorizationType) IsKnown() bool {
	switch r {
	case PendingTransactionSourceCardAuthorizationTypeCardAuthorization:
		return true
	}
	return false
}

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
	PendingTransactionSourceCardAuthorizationVerificationCardVerificationCodeResultNotChecked PendingTransactionSourceCardAuthorizationVerificationCardVerificationCodeResult = "not_checked"
	PendingTransactionSourceCardAuthorizationVerificationCardVerificationCodeResultMatch      PendingTransactionSourceCardAuthorizationVerificationCardVerificationCodeResult = "match"
	PendingTransactionSourceCardAuthorizationVerificationCardVerificationCodeResultNoMatch    PendingTransactionSourceCardAuthorizationVerificationCardVerificationCodeResult = "no_match"
)

func (r PendingTransactionSourceCardAuthorizationVerificationCardVerificationCodeResult) IsKnown() bool {
	switch r {
	case PendingTransactionSourceCardAuthorizationVerificationCardVerificationCodeResultNotChecked, PendingTransactionSourceCardAuthorizationVerificationCardVerificationCodeResultMatch, PendingTransactionSourceCardAuthorizationVerificationCardVerificationCodeResultNoMatch:
		return true
	}
	return false
}

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
	PendingTransactionSourceCardAuthorizationVerificationCardholderAddressResultNotChecked                       PendingTransactionSourceCardAuthorizationVerificationCardholderAddressResult = "not_checked"
	PendingTransactionSourceCardAuthorizationVerificationCardholderAddressResultPostalCodeMatchAddressNotChecked PendingTransactionSourceCardAuthorizationVerificationCardholderAddressResult = "postal_code_match_address_not_checked"
	PendingTransactionSourceCardAuthorizationVerificationCardholderAddressResultPostalCodeMatchAddressNoMatch    PendingTransactionSourceCardAuthorizationVerificationCardholderAddressResult = "postal_code_match_address_no_match"
	PendingTransactionSourceCardAuthorizationVerificationCardholderAddressResultPostalCodeNoMatchAddressMatch    PendingTransactionSourceCardAuthorizationVerificationCardholderAddressResult = "postal_code_no_match_address_match"
	PendingTransactionSourceCardAuthorizationVerificationCardholderAddressResultMatch                            PendingTransactionSourceCardAuthorizationVerificationCardholderAddressResult = "match"
	PendingTransactionSourceCardAuthorizationVerificationCardholderAddressResultNoMatch                          PendingTransactionSourceCardAuthorizationVerificationCardholderAddressResult = "no_match"
)

func (r PendingTransactionSourceCardAuthorizationVerificationCardholderAddressResult) IsKnown() bool {
	switch r {
	case PendingTransactionSourceCardAuthorizationVerificationCardholderAddressResultNotChecked, PendingTransactionSourceCardAuthorizationVerificationCardholderAddressResultPostalCodeMatchAddressNotChecked, PendingTransactionSourceCardAuthorizationVerificationCardholderAddressResultPostalCodeMatchAddressNoMatch, PendingTransactionSourceCardAuthorizationVerificationCardholderAddressResultPostalCodeNoMatchAddressMatch, PendingTransactionSourceCardAuthorizationVerificationCardholderAddressResultMatch, PendingTransactionSourceCardAuthorizationVerificationCardholderAddressResultNoMatch:
		return true
	}
	return false
}

// The type of the resource. We may add additional possible values for this enum
// over time; your application should be able to handle such additions gracefully.
type PendingTransactionSourceCategory string

const (
	PendingTransactionSourceCategoryAccountTransferInstruction          PendingTransactionSourceCategory = "account_transfer_instruction"
	PendingTransactionSourceCategoryACHTransferInstruction              PendingTransactionSourceCategory = "ach_transfer_instruction"
	PendingTransactionSourceCategoryCardAuthorization                   PendingTransactionSourceCategory = "card_authorization"
	PendingTransactionSourceCategoryCheckDepositInstruction             PendingTransactionSourceCategory = "check_deposit_instruction"
	PendingTransactionSourceCategoryCheckTransferInstruction            PendingTransactionSourceCategory = "check_transfer_instruction"
	PendingTransactionSourceCategoryInboundFundsHold                    PendingTransactionSourceCategory = "inbound_funds_hold"
	PendingTransactionSourceCategoryRealTimePaymentsTransferInstruction PendingTransactionSourceCategory = "real_time_payments_transfer_instruction"
	PendingTransactionSourceCategoryWireTransferInstruction             PendingTransactionSourceCategory = "wire_transfer_instruction"
	PendingTransactionSourceCategoryInboundWireTransferReversal         PendingTransactionSourceCategory = "inbound_wire_transfer_reversal"
	PendingTransactionSourceCategoryOther                               PendingTransactionSourceCategory = "other"
)

func (r PendingTransactionSourceCategory) IsKnown() bool {
	switch r {
	case PendingTransactionSourceCategoryAccountTransferInstruction, PendingTransactionSourceCategoryACHTransferInstruction, PendingTransactionSourceCategoryCardAuthorization, PendingTransactionSourceCategoryCheckDepositInstruction, PendingTransactionSourceCategoryCheckTransferInstruction, PendingTransactionSourceCategoryInboundFundsHold, PendingTransactionSourceCategoryRealTimePaymentsTransferInstruction, PendingTransactionSourceCategoryWireTransferInstruction, PendingTransactionSourceCategoryInboundWireTransferReversal, PendingTransactionSourceCategoryOther:
		return true
	}
	return false
}

// A Check Deposit Instruction object. This field will be present in the JSON
// response if and only if `category` is equal to `check_deposit_instruction`.
type PendingTransactionSourceCheckDepositInstruction struct {
	// The pending amount in USD cents.
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
	PendingTransactionSourceCheckDepositInstructionCurrencyCad PendingTransactionSourceCheckDepositInstructionCurrency = "CAD"
	PendingTransactionSourceCheckDepositInstructionCurrencyChf PendingTransactionSourceCheckDepositInstructionCurrency = "CHF"
	PendingTransactionSourceCheckDepositInstructionCurrencyEur PendingTransactionSourceCheckDepositInstructionCurrency = "EUR"
	PendingTransactionSourceCheckDepositInstructionCurrencyGbp PendingTransactionSourceCheckDepositInstructionCurrency = "GBP"
	PendingTransactionSourceCheckDepositInstructionCurrencyJpy PendingTransactionSourceCheckDepositInstructionCurrency = "JPY"
	PendingTransactionSourceCheckDepositInstructionCurrencyUsd PendingTransactionSourceCheckDepositInstructionCurrency = "USD"
)

func (r PendingTransactionSourceCheckDepositInstructionCurrency) IsKnown() bool {
	switch r {
	case PendingTransactionSourceCheckDepositInstructionCurrencyCad, PendingTransactionSourceCheckDepositInstructionCurrencyChf, PendingTransactionSourceCheckDepositInstructionCurrencyEur, PendingTransactionSourceCheckDepositInstructionCurrencyGbp, PendingTransactionSourceCheckDepositInstructionCurrencyJpy, PendingTransactionSourceCheckDepositInstructionCurrencyUsd:
		return true
	}
	return false
}

// A Check Transfer Instruction object. This field will be present in the JSON
// response if and only if `category` is equal to `check_transfer_instruction`.
type PendingTransactionSourceCheckTransferInstruction struct {
	// The transfer amount in USD cents.
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
	PendingTransactionSourceCheckTransferInstructionCurrencyCad PendingTransactionSourceCheckTransferInstructionCurrency = "CAD"
	PendingTransactionSourceCheckTransferInstructionCurrencyChf PendingTransactionSourceCheckTransferInstructionCurrency = "CHF"
	PendingTransactionSourceCheckTransferInstructionCurrencyEur PendingTransactionSourceCheckTransferInstructionCurrency = "EUR"
	PendingTransactionSourceCheckTransferInstructionCurrencyGbp PendingTransactionSourceCheckTransferInstructionCurrency = "GBP"
	PendingTransactionSourceCheckTransferInstructionCurrencyJpy PendingTransactionSourceCheckTransferInstructionCurrency = "JPY"
	PendingTransactionSourceCheckTransferInstructionCurrencyUsd PendingTransactionSourceCheckTransferInstructionCurrency = "USD"
)

func (r PendingTransactionSourceCheckTransferInstructionCurrency) IsKnown() bool {
	switch r {
	case PendingTransactionSourceCheckTransferInstructionCurrencyCad, PendingTransactionSourceCheckTransferInstructionCurrencyChf, PendingTransactionSourceCheckTransferInstructionCurrencyEur, PendingTransactionSourceCheckTransferInstructionCurrencyGbp, PendingTransactionSourceCheckTransferInstructionCurrencyJpy, PendingTransactionSourceCheckTransferInstructionCurrencyUsd:
		return true
	}
	return false
}

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
	PendingTransactionSourceInboundFundsHoldCurrencyCad PendingTransactionSourceInboundFundsHoldCurrency = "CAD"
	PendingTransactionSourceInboundFundsHoldCurrencyChf PendingTransactionSourceInboundFundsHoldCurrency = "CHF"
	PendingTransactionSourceInboundFundsHoldCurrencyEur PendingTransactionSourceInboundFundsHoldCurrency = "EUR"
	PendingTransactionSourceInboundFundsHoldCurrencyGbp PendingTransactionSourceInboundFundsHoldCurrency = "GBP"
	PendingTransactionSourceInboundFundsHoldCurrencyJpy PendingTransactionSourceInboundFundsHoldCurrency = "JPY"
	PendingTransactionSourceInboundFundsHoldCurrencyUsd PendingTransactionSourceInboundFundsHoldCurrency = "USD"
)

func (r PendingTransactionSourceInboundFundsHoldCurrency) IsKnown() bool {
	switch r {
	case PendingTransactionSourceInboundFundsHoldCurrencyCad, PendingTransactionSourceInboundFundsHoldCurrencyChf, PendingTransactionSourceInboundFundsHoldCurrencyEur, PendingTransactionSourceInboundFundsHoldCurrencyGbp, PendingTransactionSourceInboundFundsHoldCurrencyJpy, PendingTransactionSourceInboundFundsHoldCurrencyUsd:
		return true
	}
	return false
}

// The status of the hold.
type PendingTransactionSourceInboundFundsHoldStatus string

const (
	PendingTransactionSourceInboundFundsHoldStatusHeld     PendingTransactionSourceInboundFundsHoldStatus = "held"
	PendingTransactionSourceInboundFundsHoldStatusComplete PendingTransactionSourceInboundFundsHoldStatus = "complete"
)

func (r PendingTransactionSourceInboundFundsHoldStatus) IsKnown() bool {
	switch r {
	case PendingTransactionSourceInboundFundsHoldStatusHeld, PendingTransactionSourceInboundFundsHoldStatusComplete:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `inbound_funds_hold`.
type PendingTransactionSourceInboundFundsHoldType string

const (
	PendingTransactionSourceInboundFundsHoldTypeInboundFundsHold PendingTransactionSourceInboundFundsHoldType = "inbound_funds_hold"
)

func (r PendingTransactionSourceInboundFundsHoldType) IsKnown() bool {
	switch r {
	case PendingTransactionSourceInboundFundsHoldTypeInboundFundsHold:
		return true
	}
	return false
}

// An Inbound Wire Transfer Reversal object. This field will be present in the JSON
// response if and only if `category` is equal to `inbound_wire_transfer_reversal`.
type PendingTransactionSourceInboundWireTransferReversal struct {
	// The ID of the Inbound Wire Transfer that is being reversed.
	InboundWireTransferID string                                                  `json:"inbound_wire_transfer_id,required"`
	JSON                  pendingTransactionSourceInboundWireTransferReversalJSON `json:"-"`
}

// pendingTransactionSourceInboundWireTransferReversalJSON contains the JSON
// metadata for the struct [PendingTransactionSourceInboundWireTransferReversal]
type pendingTransactionSourceInboundWireTransferReversalJSON struct {
	InboundWireTransferID apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *PendingTransactionSourceInboundWireTransferReversal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pendingTransactionSourceInboundWireTransferReversalJSON) RawJSON() string {
	return r.raw
}

// A Real-Time Payments Transfer Instruction object. This field will be present in
// the JSON response if and only if `category` is equal to
// `real_time_payments_transfer_instruction`.
type PendingTransactionSourceRealTimePaymentsTransferInstruction struct {
	// The transfer amount in USD cents.
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
	// The transfer amount in USD cents.
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
	PendingTransactionStatusPending  PendingTransactionStatus = "pending"
	PendingTransactionStatusComplete PendingTransactionStatus = "complete"
)

func (r PendingTransactionStatus) IsKnown() bool {
	switch r {
	case PendingTransactionStatusPending, PendingTransactionStatusComplete:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `pending_transaction`.
type PendingTransactionType string

const (
	PendingTransactionTypePendingTransaction PendingTransactionType = "pending_transaction"
)

func (r PendingTransactionType) IsKnown() bool {
	switch r {
	case PendingTransactionTypePendingTransaction:
		return true
	}
	return false
}

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
	RouteID param.Field[string]                             `query:"route_id"`
	Status  param.Field[PendingTransactionListParamsStatus] `query:"status"`
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
	PendingTransactionListParamsCategoryInAccountTransferInstruction          PendingTransactionListParamsCategoryIn = "account_transfer_instruction"
	PendingTransactionListParamsCategoryInACHTransferInstruction              PendingTransactionListParamsCategoryIn = "ach_transfer_instruction"
	PendingTransactionListParamsCategoryInCardAuthorization                   PendingTransactionListParamsCategoryIn = "card_authorization"
	PendingTransactionListParamsCategoryInCheckDepositInstruction             PendingTransactionListParamsCategoryIn = "check_deposit_instruction"
	PendingTransactionListParamsCategoryInCheckTransferInstruction            PendingTransactionListParamsCategoryIn = "check_transfer_instruction"
	PendingTransactionListParamsCategoryInInboundFundsHold                    PendingTransactionListParamsCategoryIn = "inbound_funds_hold"
	PendingTransactionListParamsCategoryInRealTimePaymentsTransferInstruction PendingTransactionListParamsCategoryIn = "real_time_payments_transfer_instruction"
	PendingTransactionListParamsCategoryInWireTransferInstruction             PendingTransactionListParamsCategoryIn = "wire_transfer_instruction"
	PendingTransactionListParamsCategoryInInboundWireTransferReversal         PendingTransactionListParamsCategoryIn = "inbound_wire_transfer_reversal"
	PendingTransactionListParamsCategoryInOther                               PendingTransactionListParamsCategoryIn = "other"
)

func (r PendingTransactionListParamsCategoryIn) IsKnown() bool {
	switch r {
	case PendingTransactionListParamsCategoryInAccountTransferInstruction, PendingTransactionListParamsCategoryInACHTransferInstruction, PendingTransactionListParamsCategoryInCardAuthorization, PendingTransactionListParamsCategoryInCheckDepositInstruction, PendingTransactionListParamsCategoryInCheckTransferInstruction, PendingTransactionListParamsCategoryInInboundFundsHold, PendingTransactionListParamsCategoryInRealTimePaymentsTransferInstruction, PendingTransactionListParamsCategoryInWireTransferInstruction, PendingTransactionListParamsCategoryInInboundWireTransferReversal, PendingTransactionListParamsCategoryInOther:
		return true
	}
	return false
}

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
	PendingTransactionListParamsStatusInPending  PendingTransactionListParamsStatusIn = "pending"
	PendingTransactionListParamsStatusInComplete PendingTransactionListParamsStatusIn = "complete"
)

func (r PendingTransactionListParamsStatusIn) IsKnown() bool {
	switch r {
	case PendingTransactionListParamsStatusInPending, PendingTransactionListParamsStatusInComplete:
		return true
	}
	return false
}

// File generated from our OpenAPI spec by Stainless.

package increase

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/option"
)

// SimulationACHTransferService contains methods and other services that help with
// interacting with the increase API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSimulationACHTransferService]
// method instead.
type SimulationACHTransferService struct {
	Options []option.RequestOption
}

// NewSimulationACHTransferService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSimulationACHTransferService(opts ...option.RequestOption) (r *SimulationACHTransferService) {
	r = &SimulationACHTransferService{}
	r.Options = opts
	return
}

// Simulates an inbound ACH transfer to your account. This imitates initiating a
// transaction to an Increase account from a different financial institution. The
// transfer may be either a credit or a debit depending on if the `amount` is
// positive or negative. The result of calling this API will be either a
// [Transaction](#transactions) or a [Declined Transaction](#declined-transactions)
// depending on whether or not the transfer is allowed.
func (r *SimulationACHTransferService) NewInbound(ctx context.Context, body SimulationACHTransferNewInboundParams, opts ...option.RequestOption) (res *ACHTransferSimulation, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/inbound_ach_transfers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Simulates the return of an [ACH Transfer](#ach-transfers) by the Federal Reserve
// due to an error condition. This will also create a Transaction to account for
// the returned funds. This transfer must first have a `status` of `submitted`.
func (r *SimulationACHTransferService) Return(ctx context.Context, achTransferID string, body SimulationACHTransferReturnParams, opts ...option.RequestOption) (res *ACHTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("simulations/ach_transfers/%s/return", achTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Simulates the submission of an [ACH Transfer](#ach-transfers) to the Federal
// Reserve. This transfer must first have a `status` of `pending_approval` or
// `pending_submission`. In production, Increase submits ACH Transfers to the
// Federal Reserve three times per day on weekdays. Since sandbox ACH Transfers are
// not submitted to the Federal Reserve, this endpoint allows you to skip that
// delay and transition the ACH Transfer to a status of `submitted`.
func (r *SimulationACHTransferService) Submit(ctx context.Context, achTransferID string, opts ...option.RequestOption) (res *ACHTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("simulations/ach_transfers/%s/submit", achTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// The results of an inbound ACH Transfer simulation.
type ACHTransferSimulation struct {
	// If the ACH Transfer attempt fails, this will contain the resulting
	// [Declined Transaction](#declined-transactions) object. The Declined
	// Transaction's `source` will be of `category: inbound_ach_transfer`.
	DeclinedTransaction ACHTransferSimulationDeclinedTransaction `json:"declined_transaction,required,nullable"`
	// If the ACH Transfer attempt succeeds, this will contain the resulting
	// [Transaction](#transactions) object. The Transaction's `source` will be of
	// `category: inbound_ach_transfer`.
	Transaction ACHTransferSimulationTransaction `json:"transaction,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `inbound_ach_transfer_simulation_result`.
	Type ACHTransferSimulationType `json:"type,required"`
	JSON achTransferSimulationJSON
}

// achTransferSimulationJSON contains the JSON metadata for the struct
// [ACHTransferSimulation]
type achTransferSimulationJSON struct {
	DeclinedTransaction apijson.Field
	Transaction         apijson.Field
	Type                apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ACHTransferSimulation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// If the ACH Transfer attempt fails, this will contain the resulting
// [Declined Transaction](#declined-transactions) object. The Declined
// Transaction's `source` will be of `category: inbound_ach_transfer`.
type ACHTransferSimulationDeclinedTransaction struct {
	// The Declined Transaction identifier.
	ID string `json:"id,required"`
	// The identifier for the Account the Declined Transaction belongs to.
	AccountID string `json:"account_id,required"`
	// The Declined Transaction amount in the minor unit of its currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the
	// Transaction occurred.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Declined
	// Transaction's currency. This will match the currency on the Declined
	// Transaction's Account.
	Currency ACHTransferSimulationDeclinedTransactionCurrency `json:"currency,required"`
	// This is the description the vendor provides.
	Description string `json:"description,required"`
	// The identifier for the route this Declined Transaction came through. Routes are
	// things like cards and ACH details.
	RouteID string `json:"route_id,required,nullable"`
	// The type of the route this Declined Transaction came through.
	RouteType ACHTransferSimulationDeclinedTransactionRouteType `json:"route_type,required,nullable"`
	// This is an object giving more details on the network-level event that caused the
	// Declined Transaction. For example, for a card transaction this lists the
	// merchant's industry and location. Note that for backwards compatibility reasons,
	// additional undocumented keys may appear in this object. These should be treated
	// as deprecated and will be removed in the future.
	Source ACHTransferSimulationDeclinedTransactionSource `json:"source,required"`
	// A constant representing the object's type. For this resource it will always be
	// `declined_transaction`.
	Type ACHTransferSimulationDeclinedTransactionType `json:"type,required"`
	JSON achTransferSimulationDeclinedTransactionJSON
}

// achTransferSimulationDeclinedTransactionJSON contains the JSON metadata for the
// struct [ACHTransferSimulationDeclinedTransaction]
type achTransferSimulationDeclinedTransactionJSON struct {
	ID          apijson.Field
	AccountID   apijson.Field
	Amount      apijson.Field
	CreatedAt   apijson.Field
	Currency    apijson.Field
	Description apijson.Field
	RouteID     apijson.Field
	RouteType   apijson.Field
	Source      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACHTransferSimulationDeclinedTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Declined
// Transaction's currency. This will match the currency on the Declined
// Transaction's Account.
type ACHTransferSimulationDeclinedTransactionCurrency string

const (
	// Canadian Dollar (CAD)
	ACHTransferSimulationDeclinedTransactionCurrencyCad ACHTransferSimulationDeclinedTransactionCurrency = "CAD"
	// Swiss Franc (CHF)
	ACHTransferSimulationDeclinedTransactionCurrencyChf ACHTransferSimulationDeclinedTransactionCurrency = "CHF"
	// Euro (EUR)
	ACHTransferSimulationDeclinedTransactionCurrencyEur ACHTransferSimulationDeclinedTransactionCurrency = "EUR"
	// British Pound (GBP)
	ACHTransferSimulationDeclinedTransactionCurrencyGbp ACHTransferSimulationDeclinedTransactionCurrency = "GBP"
	// Japanese Yen (JPY)
	ACHTransferSimulationDeclinedTransactionCurrencyJpy ACHTransferSimulationDeclinedTransactionCurrency = "JPY"
	// US Dollar (USD)
	ACHTransferSimulationDeclinedTransactionCurrencyUsd ACHTransferSimulationDeclinedTransactionCurrency = "USD"
)

// The type of the route this Declined Transaction came through.
type ACHTransferSimulationDeclinedTransactionRouteType string

const (
	// An Account Number.
	ACHTransferSimulationDeclinedTransactionRouteTypeAccountNumber ACHTransferSimulationDeclinedTransactionRouteType = "account_number"
	// A Card.
	ACHTransferSimulationDeclinedTransactionRouteTypeCard ACHTransferSimulationDeclinedTransactionRouteType = "card"
)

// This is an object giving more details on the network-level event that caused the
// Declined Transaction. For example, for a card transaction this lists the
// merchant's industry and location. Note that for backwards compatibility reasons,
// additional undocumented keys may appear in this object. These should be treated
// as deprecated and will be removed in the future.
type ACHTransferSimulationDeclinedTransactionSource struct {
	// An ACH Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `ach_decline`.
	ACHDecline ACHTransferSimulationDeclinedTransactionSourceACHDecline `json:"ach_decline,required,nullable"`
	// A Card Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_decline`.
	CardDecline ACHTransferSimulationDeclinedTransactionSourceCardDecline `json:"card_decline,required,nullable"`
	// The type of the resource. We may add additional possible values for this enum
	// over time; your application should be able to handle such additions gracefully.
	Category ACHTransferSimulationDeclinedTransactionSourceCategory `json:"category,required"`
	// A Check Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `check_decline`.
	CheckDecline ACHTransferSimulationDeclinedTransactionSourceCheckDecline `json:"check_decline,required,nullable"`
	// An Inbound Real-Time Payments Transfer Decline object. This field will be
	// present in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_decline`.
	InboundRealTimePaymentsTransferDecline ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline `json:"inbound_real_time_payments_transfer_decline,required,nullable"`
	// An International ACH Decline object. This field will be present in the JSON
	// response if and only if `category` is equal to `international_ach_decline`.
	InternationalACHDecline ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline `json:"international_ach_decline,required,nullable"`
	// A Wire Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `wire_decline`.
	WireDecline ACHTransferSimulationDeclinedTransactionSourceWireDecline `json:"wire_decline,required,nullable"`
	JSON        achTransferSimulationDeclinedTransactionSourceJSON
}

// achTransferSimulationDeclinedTransactionSourceJSON contains the JSON metadata
// for the struct [ACHTransferSimulationDeclinedTransactionSource]
type achTransferSimulationDeclinedTransactionSourceJSON struct {
	ACHDecline                             apijson.Field
	CardDecline                            apijson.Field
	Category                               apijson.Field
	CheckDecline                           apijson.Field
	InboundRealTimePaymentsTransferDecline apijson.Field
	InternationalACHDecline                apijson.Field
	WireDecline                            apijson.Field
	raw                                    string
	ExtraFields                            map[string]apijson.Field
}

func (r *ACHTransferSimulationDeclinedTransactionSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// An ACH Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `ach_decline`.
type ACHTransferSimulationDeclinedTransactionSourceACHDecline struct {
	// The ACH Decline's identifier.
	ID string `json:"id,required"`
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The descriptive date of the transfer.
	OriginatorCompanyDescriptiveDate string `json:"originator_company_descriptive_date,required,nullable"`
	// The additional information included with the transfer.
	OriginatorCompanyDiscretionaryData string `json:"originator_company_discretionary_data,required,nullable"`
	// The identifier of the company that initiated the transfer.
	OriginatorCompanyID string `json:"originator_company_id,required"`
	// The name of the company that initiated the transfer.
	OriginatorCompanyName string `json:"originator_company_name,required"`
	// Why the ACH transfer was declined.
	Reason ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason `json:"reason,required"`
	// The id of the receiver of the transfer.
	ReceiverIDNumber string `json:"receiver_id_number,required,nullable"`
	// The name of the receiver of the transfer.
	ReceiverName string `json:"receiver_name,required,nullable"`
	// The trace number of the transfer.
	TraceNumber string `json:"trace_number,required"`
	// A constant representing the object's type. For this resource it will always be
	// `ach_decline`.
	Type ACHTransferSimulationDeclinedTransactionSourceACHDeclineType `json:"type,required"`
	JSON achTransferSimulationDeclinedTransactionSourceACHDeclineJSON
}

// achTransferSimulationDeclinedTransactionSourceACHDeclineJSON contains the JSON
// metadata for the struct
// [ACHTransferSimulationDeclinedTransactionSourceACHDecline]
type achTransferSimulationDeclinedTransactionSourceACHDeclineJSON struct {
	ID                                 apijson.Field
	Amount                             apijson.Field
	OriginatorCompanyDescriptiveDate   apijson.Field
	OriginatorCompanyDiscretionaryData apijson.Field
	OriginatorCompanyID                apijson.Field
	OriginatorCompanyName              apijson.Field
	Reason                             apijson.Field
	ReceiverIDNumber                   apijson.Field
	ReceiverName                       apijson.Field
	TraceNumber                        apijson.Field
	Type                               apijson.Field
	raw                                string
	ExtraFields                        map[string]apijson.Field
}

func (r *ACHTransferSimulationDeclinedTransactionSourceACHDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Why the ACH transfer was declined.
type ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason string

const (
	// The account number is canceled.
	ACHTransferSimulationDeclinedTransactionSourceACHDeclineReasonACHRouteCanceled ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason = "ach_route_canceled"
	// The account number is disabled.
	ACHTransferSimulationDeclinedTransactionSourceACHDeclineReasonACHRouteDisabled ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason = "ach_route_disabled"
	// The transaction would cause a limit to be exceeded.
	ACHTransferSimulationDeclinedTransactionSourceACHDeclineReasonBreachesLimit ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason = "breaches_limit"
	// A credit was refused.
	ACHTransferSimulationDeclinedTransactionSourceACHDeclineReasonCreditEntryRefusedByReceiver ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason = "credit_entry_refused_by_receiver"
	// Other.
	ACHTransferSimulationDeclinedTransactionSourceACHDeclineReasonDuplicateReturn ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason = "duplicate_return"
	// The account's entity is not active.
	ACHTransferSimulationDeclinedTransactionSourceACHDeclineReasonEntityNotActive ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason = "entity_not_active"
	// Your account is inactive.
	ACHTransferSimulationDeclinedTransactionSourceACHDeclineReasonGroupLocked ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason = "group_locked"
	// Your account contains insufficient funds.
	ACHTransferSimulationDeclinedTransactionSourceACHDeclineReasonInsufficientFunds ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason = "insufficient_funds"
	// Other.
	ACHTransferSimulationDeclinedTransactionSourceACHDeclineReasonMisroutedReturn ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason = "misrouted_return"
	// Other.
	ACHTransferSimulationDeclinedTransactionSourceACHDeclineReasonReturnOfErroneousOrReversingDebit ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason = "return_of_erroneous_or_reversing_debit"
	// The account number that was debited does not exist.
	ACHTransferSimulationDeclinedTransactionSourceACHDeclineReasonNoACHRoute ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason = "no_ach_route"
	// Other.
	ACHTransferSimulationDeclinedTransactionSourceACHDeclineReasonOriginatorRequest ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason = "originator_request"
	// The transaction is not allowed per Increase's terms.
	ACHTransferSimulationDeclinedTransactionSourceACHDeclineReasonTransactionNotAllowed ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason = "transaction_not_allowed"
	// The user initiated the decline.
	ACHTransferSimulationDeclinedTransactionSourceACHDeclineReasonUserInitiated ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason = "user_initiated"
)

// A constant representing the object's type. For this resource it will always be
// `ach_decline`.
type ACHTransferSimulationDeclinedTransactionSourceACHDeclineType string

const (
	ACHTransferSimulationDeclinedTransactionSourceACHDeclineTypeACHDecline ACHTransferSimulationDeclinedTransactionSourceACHDeclineType = "ach_decline"
)

// A Card Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `card_decline`.
type ACHTransferSimulationDeclinedTransactionSourceCardDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency ACHTransferSimulationDeclinedTransactionSourceCardDeclineCurrency `json:"currency,required"`
	// If the authorization was attempted using a Digital Wallet Token (such as an
	// Apple Pay purchase), the identifier of the token that was used.
	DigitalWalletTokenID string `json:"digital_wallet_token_id,required,nullable"`
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
	// The state the merchant resides in.
	MerchantState string `json:"merchant_state,required,nullable"`
	// Fields specific to the `network`
	NetworkDetails ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetails `json:"network_details,required"`
	// If the authorization was made in-person with a physical card, the Physical Card
	// that was used.
	PhysicalCardID string `json:"physical_card_id,required,nullable"`
	// The identifier of the Real-Time Decision sent to approve or decline this
	// transaction.
	RealTimeDecisionID string `json:"real_time_decision_id,required,nullable"`
	// Why the transaction was declined.
	Reason ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason `json:"reason,required"`
	JSON   achTransferSimulationDeclinedTransactionSourceCardDeclineJSON
}

// achTransferSimulationDeclinedTransactionSourceCardDeclineJSON contains the JSON
// metadata for the struct
// [ACHTransferSimulationDeclinedTransactionSourceCardDecline]
type achTransferSimulationDeclinedTransactionSourceCardDeclineJSON struct {
	Amount               apijson.Field
	Currency             apijson.Field
	DigitalWalletTokenID apijson.Field
	MerchantAcceptorID   apijson.Field
	MerchantCategoryCode apijson.Field
	MerchantCity         apijson.Field
	MerchantCountry      apijson.Field
	MerchantDescriptor   apijson.Field
	MerchantState        apijson.Field
	NetworkDetails       apijson.Field
	PhysicalCardID       apijson.Field
	RealTimeDecisionID   apijson.Field
	Reason               apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ACHTransferSimulationDeclinedTransactionSourceCardDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
// account currency.
type ACHTransferSimulationDeclinedTransactionSourceCardDeclineCurrency string

const (
	// Canadian Dollar (CAD)
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineCurrencyCad ACHTransferSimulationDeclinedTransactionSourceCardDeclineCurrency = "CAD"
	// Swiss Franc (CHF)
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineCurrencyChf ACHTransferSimulationDeclinedTransactionSourceCardDeclineCurrency = "CHF"
	// Euro (EUR)
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineCurrencyEur ACHTransferSimulationDeclinedTransactionSourceCardDeclineCurrency = "EUR"
	// British Pound (GBP)
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineCurrencyGbp ACHTransferSimulationDeclinedTransactionSourceCardDeclineCurrency = "GBP"
	// Japanese Yen (JPY)
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineCurrencyJpy ACHTransferSimulationDeclinedTransactionSourceCardDeclineCurrency = "JPY"
	// US Dollar (USD)
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineCurrencyUsd ACHTransferSimulationDeclinedTransactionSourceCardDeclineCurrency = "USD"
)

// Fields specific to the `network`
type ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetails struct {
	// The payment network used to process this card authorization
	Category ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsCategory `json:"category,required"`
	// Fields specific to the `visa` network
	Visa ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa `json:"visa,required,nullable"`
	JSON achTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsJSON
}

// achTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsJSON
// contains the JSON metadata for the struct
// [ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetails]
type achTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsJSON struct {
	Category    apijson.Field
	Visa        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The payment network used to process this card authorization
type ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsCategory string

const (
	// Visa
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsCategoryVisa ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsCategory = "visa"
)

// Fields specific to the `visa` network
type ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa struct {
	// For electronic commerce transactions, this identifies the level of security used
	// in obtaining the customer's payment credential. For mail or telephone order
	// transactions, identifies the type of mail or telephone order.
	ElectronicCommerceIndicator ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator `json:"electronic_commerce_indicator,required,nullable"`
	// The method used to enter the cardholder's primary account number and card
	// expiration date
	PointOfServiceEntryMode ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode `json:"point_of_service_entry_mode,required,nullable"`
	JSON                    achTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaJSON
}

// achTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaJSON
// contains the JSON metadata for the struct
// [ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa]
type achTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaJSON struct {
	ElectronicCommerceIndicator apijson.Field
	PointOfServiceEntryMode     apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// For electronic commerce transactions, this identifies the level of security used
// in obtaining the customer's payment credential. For mail or telephone order
// transactions, identifies the type of mail or telephone order.
type ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator string

const (
	// Single transaction of a mail/phone order: Use to indicate that the transaction
	// is a mail/phone order purchase, not a recurring transaction or installment
	// payment. For domestic transactions in the US region, this value may also
	// indicate one bill payment transaction in the card-present or card-absent
	// environments.
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorMailPhoneOrder ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "mail_phone_order"
	// Recurring transaction: Payment indicator used to indicate a recurring
	// transaction that originates from an acquirer in the US region.
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorRecurring ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "recurring"
	// Installment payment: Payment indicator used to indicate one purchase of goods or
	// services that is billed to the account in multiple charges over a period of time
	// agreed upon by the cardholder and merchant from transactions that originate from
	// an acquirer in the US region.
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorInstallment ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "installment"
	// Unknown classification: other mail order: Use to indicate that the type of
	// mail/telephone order is unknown.
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorUnknownMailPhoneOrder ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "unknown_mail_phone_order"
	// Secure electronic commerce transaction: Use to indicate that the electronic
	// commerce transaction has been authenticated using e.g., 3-D Secure
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorSecureElectronicCommerce ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "secure_electronic_commerce"
	// Non-authenticated security transaction at a 3-D Secure-capable merchant, and
	// merchant attempted to authenticate the cardholder using 3-D Secure: Use to
	// identify an electronic commerce transaction where the merchant attempted to
	// authenticate the cardholder using 3-D Secure, but was unable to complete the
	// authentication because the issuer or cardholder does not participate in the 3-D
	// Secure program.
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransactionAt3DSCapableMerchant ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction_at_3ds_capable_merchant"
	// Non-authenticated security transaction: Use to identify an electronic commerce
	// transaction that uses data encryption for security however , cardholder
	// authentication is not performed using 3-D Secure.
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransaction ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction"
	// Non-secure transaction: Use to identify an electronic commerce transaction that
	// has no data protection.
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonSecureTransaction ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "non_secure_transaction"
)

// The method used to enter the cardholder's primary account number and card
// expiration date
type ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode string

const (
	// Unknown
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeUnknown ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "unknown"
	// Manual key entry
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeManual ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "manual"
	// Magnetic stripe read, without card verification value
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeMagneticStripeNoCvv ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "magnetic_stripe_no_cvv"
	// Optical code
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeOpticalCode ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "optical_code"
	// Contact chip card
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCard ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "integrated_circuit_card"
	// Contactless read of chip card
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeContactless ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "contactless"
	// Transaction initiated using a credential that has previously been stored on file
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeCredentialOnFile ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "credential_on_file"
	// Magnetic stripe read
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeMagneticStripe ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "magnetic_stripe"
	// Contactless read of magnetic stripe data
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeContactlessMagneticStripe ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "contactless_magnetic_stripe"
	// Contact chip card, without card verification value
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCardNoCvv ACHTransferSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaPointOfServiceEntryMode = "integrated_circuit_card_no_cvv"
)

// Why the transaction was declined.
type ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason string

const (
	// The Card was not active.
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineReasonCardNotActive ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason = "card_not_active"
	// The Physical Card was not active.
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineReasonPhysicalCardNotActive ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason = "physical_card_not_active"
	// The account's entity was not active.
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineReasonEntityNotActive ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason = "entity_not_active"
	// The account was inactive.
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineReasonGroupLocked ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason = "group_locked"
	// The Card's Account did not have a sufficient available balance.
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineReasonInsufficientFunds ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason = "insufficient_funds"
	// The given CVV2 did not match the card's value.
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineReasonCvv2Mismatch ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason = "cvv2_mismatch"
	// The attempted card transaction is not allowed per Increase's terms.
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineReasonTransactionNotAllowed ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason = "transaction_not_allowed"
	// The transaction was blocked by a Limit.
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineReasonBreachesLimit ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason = "breaches_limit"
	// Your application declined the transaction via webhook.
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineReasonWebhookDeclined ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason = "webhook_declined"
	// Your application webhook did not respond without the required timeout.
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineReasonWebhookTimedOut ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason = "webhook_timed_out"
	// Declined by stand-in processing.
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineReasonDeclinedByStandInProcessing ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason = "declined_by_stand_in_processing"
	// The card read had an invalid CVV, dCVV, or authorization request cryptogram.
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineReasonInvalidPhysicalCard ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason = "invalid_physical_card"
	// The original card authorization for this incremental authorization does not
	// exist.
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineReasonMissingOriginalAuthorization ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason = "missing_original_authorization"
	// The transaction was suspected to be fraudulent. Please reach out to
	// support@increase.com for more information.
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineReasonSuspectedFraud ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason = "suspected_fraud"
)

// The type of the resource. We may add additional possible values for this enum
// over time; your application should be able to handle such additions gracefully.
type ACHTransferSimulationDeclinedTransactionSourceCategory string

const (
	// ACH Decline: details will be under the `ach_decline` object.
	ACHTransferSimulationDeclinedTransactionSourceCategoryACHDecline ACHTransferSimulationDeclinedTransactionSourceCategory = "ach_decline"
	// Card Decline: details will be under the `card_decline` object.
	ACHTransferSimulationDeclinedTransactionSourceCategoryCardDecline ACHTransferSimulationDeclinedTransactionSourceCategory = "card_decline"
	// Check Decline: details will be under the `check_decline` object.
	ACHTransferSimulationDeclinedTransactionSourceCategoryCheckDecline ACHTransferSimulationDeclinedTransactionSourceCategory = "check_decline"
	// Inbound Real-Time Payments Transfer Decline: details will be under the
	// `inbound_real_time_payments_transfer_decline` object.
	ACHTransferSimulationDeclinedTransactionSourceCategoryInboundRealTimePaymentsTransferDecline ACHTransferSimulationDeclinedTransactionSourceCategory = "inbound_real_time_payments_transfer_decline"
	// International ACH Decline: details will be under the `international_ach_decline`
	// object.
	ACHTransferSimulationDeclinedTransactionSourceCategoryInternationalACHDecline ACHTransferSimulationDeclinedTransactionSourceCategory = "international_ach_decline"
	// Wire Decline: details will be under the `wire_decline` object.
	ACHTransferSimulationDeclinedTransactionSourceCategoryWireDecline ACHTransferSimulationDeclinedTransactionSourceCategory = "wire_decline"
	// The Declined Transaction was made for an undocumented or deprecated reason.
	ACHTransferSimulationDeclinedTransactionSourceCategoryOther ACHTransferSimulationDeclinedTransactionSourceCategory = "other"
)

// A Check Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `check_decline`.
type ACHTransferSimulationDeclinedTransactionSourceCheckDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount        int64  `json:"amount,required"`
	AuxiliaryOnUs string `json:"auxiliary_on_us,required,nullable"`
	// Why the check was declined.
	Reason ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReason `json:"reason,required"`
	JSON   achTransferSimulationDeclinedTransactionSourceCheckDeclineJSON
}

// achTransferSimulationDeclinedTransactionSourceCheckDeclineJSON contains the JSON
// metadata for the struct
// [ACHTransferSimulationDeclinedTransactionSourceCheckDecline]
type achTransferSimulationDeclinedTransactionSourceCheckDeclineJSON struct {
	Amount        apijson.Field
	AuxiliaryOnUs apijson.Field
	Reason        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ACHTransferSimulationDeclinedTransactionSourceCheckDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Why the check was declined.
type ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReason string

const (
	// The account number is canceled.
	ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReasonACHRouteCanceled ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReason = "ach_route_canceled"
	// The account number is disabled.
	ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReasonACHRouteDisabled ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReason = "ach_route_disabled"
	// The transaction would cause a limit to be exceeded.
	ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReasonBreachesLimit ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReason = "breaches_limit"
	// The account's entity is not active.
	ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReasonEntityNotActive ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReason = "entity_not_active"
	// Your account is inactive.
	ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReasonGroupLocked ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReason = "group_locked"
	// Your account contains insufficient funds.
	ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReasonInsufficientFunds ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReason = "insufficient_funds"
	// Unable to locate account.
	ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReasonUnableToLocateAccount ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReason = "unable_to_locate_account"
	// Routing number on the check is not ours.
	ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReasonNotOurItem ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReason = "not_our_item"
	// Unable to process.
	ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReasonUnableToProcess ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReason = "unable_to_process"
	// Refer to image.
	ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReasonReferToImage ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReason = "refer_to_image"
	// Stop payment requested for this check.
	ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReasonStopPaymentRequested ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReason = "stop_payment_requested"
	// Check was returned to sender.
	ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReasonReturned ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReason = "returned"
	// The check was a duplicate deposit.
	ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReasonDuplicatePresentment ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReason = "duplicate_presentment"
	// The transaction is not allowed.
	ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReasonNotAuthorized ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReason = "not_authorized"
	// The check was altered or fictitious.
	ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReasonAlteredOrFictitious ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReason = "altered_or_fictitious"
)

// An Inbound Real-Time Payments Transfer Decline object. This field will be
// present in the JSON response if and only if `category` is equal to
// `inbound_real_time_payments_transfer_decline`.
type ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The name the sender of the transfer specified as the recipient of the transfer.
	CreditorName string `json:"creditor_name,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the declined
	// transfer's currency. This will always be "USD" for a Real-Time Payments
	// transfer.
	Currency ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency `json:"currency,required"`
	// The account number of the account that sent the transfer.
	DebtorAccountNumber string `json:"debtor_account_number,required"`
	// The name provided by the sender of the transfer.
	DebtorName string `json:"debtor_name,required"`
	// The routing number of the account that sent the transfer.
	DebtorRoutingNumber string `json:"debtor_routing_number,required"`
	// Why the transfer was declined.
	Reason ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason `json:"reason,required"`
	// Additional information included with the transfer.
	RemittanceInformation string `json:"remittance_information,required,nullable"`
	// The Real-Time Payments network identification of the declined transfer.
	TransactionIdentification string `json:"transaction_identification,required"`
	JSON                      achTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineJSON
}

// achTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineJSON
// contains the JSON metadata for the struct
// [ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline]
type achTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineJSON struct {
	Amount                    apijson.Field
	CreditorName              apijson.Field
	Currency                  apijson.Field
	DebtorAccountNumber       apijson.Field
	DebtorName                apijson.Field
	DebtorRoutingNumber       apijson.Field
	Reason                    apijson.Field
	RemittanceInformation     apijson.Field
	TransactionIdentification apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the declined
// transfer's currency. This will always be "USD" for a Real-Time Payments
// transfer.
type ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency string

const (
	// Canadian Dollar (CAD)
	ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyCad ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "CAD"
	// Swiss Franc (CHF)
	ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyChf ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "CHF"
	// Euro (EUR)
	ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyEur ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "EUR"
	// British Pound (GBP)
	ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyGbp ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "GBP"
	// Japanese Yen (JPY)
	ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyJpy ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "JPY"
	// US Dollar (USD)
	ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyUsd ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "USD"
)

// Why the transfer was declined.
type ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason string

const (
	// The account number is canceled.
	ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonAccountNumberCanceled ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "account_number_canceled"
	// The account number is disabled.
	ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonAccountNumberDisabled ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "account_number_disabled"
	// Your account is restricted.
	ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonAccountRestricted ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "account_restricted"
	// Your account is inactive.
	ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonGroupLocked ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "group_locked"
	// The account's entity is not active.
	ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonEntityNotActive ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "entity_not_active"
	// Your account is not enabled to receive Real-Time Payments transfers.
	ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonRealTimePaymentsNotEnabled ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "real_time_payments_not_enabled"
)

// An International ACH Decline object. This field will be present in the JSON
// response if and only if `category` is equal to `international_ach_decline`.
type ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount                                                 int64  `json:"amount,required"`
	DestinationCountryCode                                 string `json:"destination_country_code,required"`
	DestinationCurrencyCode                                string `json:"destination_currency_code,required"`
	ForeignExchangeIndicator                               string `json:"foreign_exchange_indicator,required"`
	ForeignExchangeReference                               string `json:"foreign_exchange_reference,required,nullable"`
	ForeignExchangeReferenceIndicator                      string `json:"foreign_exchange_reference_indicator,required"`
	ForeignPaymentAmount                                   int64  `json:"foreign_payment_amount,required"`
	ForeignTraceNumber                                     string `json:"foreign_trace_number,required,nullable"`
	InternationalTransactionTypeCode                       string `json:"international_transaction_type_code,required"`
	OriginatingCurrencyCode                                string `json:"originating_currency_code,required"`
	OriginatingDepositoryFinancialInstitutionBranchCountry string `json:"originating_depository_financial_institution_branch_country,required"`
	OriginatingDepositoryFinancialInstitutionID            string `json:"originating_depository_financial_institution_id,required"`
	OriginatingDepositoryFinancialInstitutionIDQualifier   string `json:"originating_depository_financial_institution_id_qualifier,required"`
	OriginatingDepositoryFinancialInstitutionName          string `json:"originating_depository_financial_institution_name,required"`
	OriginatorCity                                         string `json:"originator_city,required"`
	OriginatorCompanyEntryDescription                      string `json:"originator_company_entry_description,required"`
	OriginatorCountry                                      string `json:"originator_country,required"`
	OriginatorIdentification                               string `json:"originator_identification,required"`
	OriginatorName                                         string `json:"originator_name,required"`
	OriginatorPostalCode                                   string `json:"originator_postal_code,required,nullable"`
	OriginatorStateOrProvince                              string `json:"originator_state_or_province,required,nullable"`
	OriginatorStreetAddress                                string `json:"originator_street_address,required"`
	PaymentRelatedInformation                              string `json:"payment_related_information,required,nullable"`
	PaymentRelatedInformation2                             string `json:"payment_related_information2,required,nullable"`
	ReceiverCity                                           string `json:"receiver_city,required"`
	ReceiverCountry                                        string `json:"receiver_country,required"`
	ReceiverIdentificationNumber                           string `json:"receiver_identification_number,required,nullable"`
	ReceiverPostalCode                                     string `json:"receiver_postal_code,required,nullable"`
	ReceiverStateOrProvince                                string `json:"receiver_state_or_province,required,nullable"`
	ReceiverStreetAddress                                  string `json:"receiver_street_address,required"`
	ReceivingCompanyOrIndividualName                       string `json:"receiving_company_or_individual_name,required"`
	ReceivingDepositoryFinancialInstitutionCountry         string `json:"receiving_depository_financial_institution_country,required"`
	ReceivingDepositoryFinancialInstitutionID              string `json:"receiving_depository_financial_institution_id,required"`
	ReceivingDepositoryFinancialInstitutionIDQualifier     string `json:"receiving_depository_financial_institution_id_qualifier,required"`
	ReceivingDepositoryFinancialInstitutionName            string `json:"receiving_depository_financial_institution_name,required"`
	TraceNumber                                            string `json:"trace_number,required"`
	JSON                                                   achTransferSimulationDeclinedTransactionSourceInternationalACHDeclineJSON
}

// achTransferSimulationDeclinedTransactionSourceInternationalACHDeclineJSON
// contains the JSON metadata for the struct
// [ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline]
type achTransferSimulationDeclinedTransactionSourceInternationalACHDeclineJSON struct {
	Amount                                                 apijson.Field
	DestinationCountryCode                                 apijson.Field
	DestinationCurrencyCode                                apijson.Field
	ForeignExchangeIndicator                               apijson.Field
	ForeignExchangeReference                               apijson.Field
	ForeignExchangeReferenceIndicator                      apijson.Field
	ForeignPaymentAmount                                   apijson.Field
	ForeignTraceNumber                                     apijson.Field
	InternationalTransactionTypeCode                       apijson.Field
	OriginatingCurrencyCode                                apijson.Field
	OriginatingDepositoryFinancialInstitutionBranchCountry apijson.Field
	OriginatingDepositoryFinancialInstitutionID            apijson.Field
	OriginatingDepositoryFinancialInstitutionIDQualifier   apijson.Field
	OriginatingDepositoryFinancialInstitutionName          apijson.Field
	OriginatorCity                                         apijson.Field
	OriginatorCompanyEntryDescription                      apijson.Field
	OriginatorCountry                                      apijson.Field
	OriginatorIdentification                               apijson.Field
	OriginatorName                                         apijson.Field
	OriginatorPostalCode                                   apijson.Field
	OriginatorStateOrProvince                              apijson.Field
	OriginatorStreetAddress                                apijson.Field
	PaymentRelatedInformation                              apijson.Field
	PaymentRelatedInformation2                             apijson.Field
	ReceiverCity                                           apijson.Field
	ReceiverCountry                                        apijson.Field
	ReceiverIdentificationNumber                           apijson.Field
	ReceiverPostalCode                                     apijson.Field
	ReceiverStateOrProvince                                apijson.Field
	ReceiverStreetAddress                                  apijson.Field
	ReceivingCompanyOrIndividualName                       apijson.Field
	ReceivingDepositoryFinancialInstitutionCountry         apijson.Field
	ReceivingDepositoryFinancialInstitutionID              apijson.Field
	ReceivingDepositoryFinancialInstitutionIDQualifier     apijson.Field
	ReceivingDepositoryFinancialInstitutionName            apijson.Field
	TraceNumber                                            apijson.Field
	raw                                                    string
	ExtraFields                                            map[string]apijson.Field
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Wire Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `wire_decline`.
type ACHTransferSimulationDeclinedTransactionSourceWireDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount                                  int64  `json:"amount,required"`
	BeneficiaryAddressLine1                 string `json:"beneficiary_address_line1,required,nullable"`
	BeneficiaryAddressLine2                 string `json:"beneficiary_address_line2,required,nullable"`
	BeneficiaryAddressLine3                 string `json:"beneficiary_address_line3,required,nullable"`
	BeneficiaryName                         string `json:"beneficiary_name,required,nullable"`
	BeneficiaryReference                    string `json:"beneficiary_reference,required,nullable"`
	Description                             string `json:"description,required"`
	InputMessageAccountabilityData          string `json:"input_message_accountability_data,required,nullable"`
	OriginatorAddressLine1                  string `json:"originator_address_line1,required,nullable"`
	OriginatorAddressLine2                  string `json:"originator_address_line2,required,nullable"`
	OriginatorAddressLine3                  string `json:"originator_address_line3,required,nullable"`
	OriginatorName                          string `json:"originator_name,required,nullable"`
	OriginatorToBeneficiaryInformationLine1 string `json:"originator_to_beneficiary_information_line1,required,nullable"`
	OriginatorToBeneficiaryInformationLine2 string `json:"originator_to_beneficiary_information_line2,required,nullable"`
	OriginatorToBeneficiaryInformationLine3 string `json:"originator_to_beneficiary_information_line3,required,nullable"`
	OriginatorToBeneficiaryInformationLine4 string `json:"originator_to_beneficiary_information_line4,required,nullable"`
	// Why the wire transfer was declined.
	Reason ACHTransferSimulationDeclinedTransactionSourceWireDeclineReason `json:"reason,required"`
	JSON   achTransferSimulationDeclinedTransactionSourceWireDeclineJSON
}

// achTransferSimulationDeclinedTransactionSourceWireDeclineJSON contains the JSON
// metadata for the struct
// [ACHTransferSimulationDeclinedTransactionSourceWireDecline]
type achTransferSimulationDeclinedTransactionSourceWireDeclineJSON struct {
	Amount                                  apijson.Field
	BeneficiaryAddressLine1                 apijson.Field
	BeneficiaryAddressLine2                 apijson.Field
	BeneficiaryAddressLine3                 apijson.Field
	BeneficiaryName                         apijson.Field
	BeneficiaryReference                    apijson.Field
	Description                             apijson.Field
	InputMessageAccountabilityData          apijson.Field
	OriginatorAddressLine1                  apijson.Field
	OriginatorAddressLine2                  apijson.Field
	OriginatorAddressLine3                  apijson.Field
	OriginatorName                          apijson.Field
	OriginatorToBeneficiaryInformationLine1 apijson.Field
	OriginatorToBeneficiaryInformationLine2 apijson.Field
	OriginatorToBeneficiaryInformationLine3 apijson.Field
	OriginatorToBeneficiaryInformationLine4 apijson.Field
	Reason                                  apijson.Field
	raw                                     string
	ExtraFields                             map[string]apijson.Field
}

func (r *ACHTransferSimulationDeclinedTransactionSourceWireDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Why the wire transfer was declined.
type ACHTransferSimulationDeclinedTransactionSourceWireDeclineReason string

const (
	// The account number is canceled.
	ACHTransferSimulationDeclinedTransactionSourceWireDeclineReasonAccountNumberCanceled ACHTransferSimulationDeclinedTransactionSourceWireDeclineReason = "account_number_canceled"
	// The account number is disabled.
	ACHTransferSimulationDeclinedTransactionSourceWireDeclineReasonAccountNumberDisabled ACHTransferSimulationDeclinedTransactionSourceWireDeclineReason = "account_number_disabled"
	// The account's entity is not active.
	ACHTransferSimulationDeclinedTransactionSourceWireDeclineReasonEntityNotActive ACHTransferSimulationDeclinedTransactionSourceWireDeclineReason = "entity_not_active"
	// Your account is inactive.
	ACHTransferSimulationDeclinedTransactionSourceWireDeclineReasonGroupLocked ACHTransferSimulationDeclinedTransactionSourceWireDeclineReason = "group_locked"
	// The beneficiary account number does not exist.
	ACHTransferSimulationDeclinedTransactionSourceWireDeclineReasonNoAccountNumber ACHTransferSimulationDeclinedTransactionSourceWireDeclineReason = "no_account_number"
	// The transaction is not allowed per Increase's terms.
	ACHTransferSimulationDeclinedTransactionSourceWireDeclineReasonTransactionNotAllowed ACHTransferSimulationDeclinedTransactionSourceWireDeclineReason = "transaction_not_allowed"
)

// A constant representing the object's type. For this resource it will always be
// `declined_transaction`.
type ACHTransferSimulationDeclinedTransactionType string

const (
	ACHTransferSimulationDeclinedTransactionTypeDeclinedTransaction ACHTransferSimulationDeclinedTransactionType = "declined_transaction"
)

// If the ACH Transfer attempt succeeds, this will contain the resulting
// [Transaction](#transactions) object. The Transaction's `source` will be of
// `category: inbound_ach_transfer`.
type ACHTransferSimulationTransaction struct {
	// The Transaction identifier.
	ID string `json:"id,required"`
	// The identifier for the Account the Transaction belongs to.
	AccountID string `json:"account_id,required"`
	// The Transaction amount in the minor unit of its currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the
	// Transaction occurred.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// Transaction's currency. This will match the currency on the Transaction's
	// Account.
	Currency ACHTransferSimulationTransactionCurrency `json:"currency,required"`
	// An informational message describing this transaction. Use the fields in `source`
	// to get more detailed information. This field appears as the line-item on the
	// statement.
	Description string `json:"description,required"`
	// The identifier for the route this Transaction came through. Routes are things
	// like cards and ACH details.
	RouteID string `json:"route_id,required,nullable"`
	// The type of the route this Transaction came through.
	RouteType ACHTransferSimulationTransactionRouteType `json:"route_type,required,nullable"`
	// This is an object giving more details on the network-level event that caused the
	// Transaction. Note that for backwards compatibility reasons, additional
	// undocumented keys may appear in this object. These should be treated as
	// deprecated and will be removed in the future.
	Source ACHTransferSimulationTransactionSource `json:"source,required"`
	// A constant representing the object's type. For this resource it will always be
	// `transaction`.
	Type ACHTransferSimulationTransactionType `json:"type,required"`
	JSON achTransferSimulationTransactionJSON
}

// achTransferSimulationTransactionJSON contains the JSON metadata for the struct
// [ACHTransferSimulationTransaction]
type achTransferSimulationTransactionJSON struct {
	ID          apijson.Field
	AccountID   apijson.Field
	Amount      apijson.Field
	CreatedAt   apijson.Field
	Currency    apijson.Field
	Description apijson.Field
	RouteID     apijson.Field
	RouteType   apijson.Field
	Source      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACHTransferSimulationTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// Transaction's currency. This will match the currency on the Transaction's
// Account.
type ACHTransferSimulationTransactionCurrency string

const (
	// Canadian Dollar (CAD)
	ACHTransferSimulationTransactionCurrencyCad ACHTransferSimulationTransactionCurrency = "CAD"
	// Swiss Franc (CHF)
	ACHTransferSimulationTransactionCurrencyChf ACHTransferSimulationTransactionCurrency = "CHF"
	// Euro (EUR)
	ACHTransferSimulationTransactionCurrencyEur ACHTransferSimulationTransactionCurrency = "EUR"
	// British Pound (GBP)
	ACHTransferSimulationTransactionCurrencyGbp ACHTransferSimulationTransactionCurrency = "GBP"
	// Japanese Yen (JPY)
	ACHTransferSimulationTransactionCurrencyJpy ACHTransferSimulationTransactionCurrency = "JPY"
	// US Dollar (USD)
	ACHTransferSimulationTransactionCurrencyUsd ACHTransferSimulationTransactionCurrency = "USD"
)

// The type of the route this Transaction came through.
type ACHTransferSimulationTransactionRouteType string

const (
	// An Account Number.
	ACHTransferSimulationTransactionRouteTypeAccountNumber ACHTransferSimulationTransactionRouteType = "account_number"
	// A Card.
	ACHTransferSimulationTransactionRouteTypeCard ACHTransferSimulationTransactionRouteType = "card"
)

// This is an object giving more details on the network-level event that caused the
// Transaction. Note that for backwards compatibility reasons, additional
// undocumented keys may appear in this object. These should be treated as
// deprecated and will be removed in the future.
type ACHTransferSimulationTransactionSource struct {
	// An Account Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `account_transfer_intention`.
	AccountTransferIntention ACHTransferSimulationTransactionSourceAccountTransferIntention `json:"account_transfer_intention,required,nullable"`
	// An ACH Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `ach_transfer_intention`.
	ACHTransferIntention ACHTransferSimulationTransactionSourceACHTransferIntention `json:"ach_transfer_intention,required,nullable"`
	// An ACH Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `ach_transfer_rejection`.
	ACHTransferRejection ACHTransferSimulationTransactionSourceACHTransferRejection `json:"ach_transfer_rejection,required,nullable"`
	// An ACH Transfer Return object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_transfer_return`.
	ACHTransferReturn ACHTransferSimulationTransactionSourceACHTransferReturn `json:"ach_transfer_return,required,nullable"`
	// A Card Dispute Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_dispute_acceptance`.
	CardDisputeAcceptance ACHTransferSimulationTransactionSourceCardDisputeAcceptance `json:"card_dispute_acceptance,required,nullable"`
	// A Card Refund object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_refund`.
	CardRefund ACHTransferSimulationTransactionSourceCardRefund `json:"card_refund,required,nullable"`
	// A Card Revenue Payment object. This field will be present in the JSON response
	// if and only if `category` is equal to `card_revenue_payment`.
	CardRevenuePayment ACHTransferSimulationTransactionSourceCardRevenuePayment `json:"card_revenue_payment,required,nullable"`
	// A Card Settlement object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_settlement`.
	CardSettlement ACHTransferSimulationTransactionSourceCardSettlement `json:"card_settlement,required,nullable"`
	// The type of the resource. We may add additional possible values for this enum
	// over time; your application should be able to handle such additions gracefully.
	Category ACHTransferSimulationTransactionSourceCategory `json:"category,required"`
	// A Check Deposit Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_deposit_acceptance`.
	CheckDepositAcceptance ACHTransferSimulationTransactionSourceCheckDepositAcceptance `json:"check_deposit_acceptance,required,nullable"`
	// A Check Deposit Return object. This field will be present in the JSON response
	// if and only if `category` is equal to `check_deposit_return`.
	CheckDepositReturn ACHTransferSimulationTransactionSourceCheckDepositReturn `json:"check_deposit_return,required,nullable"`
	// A Check Transfer Deposit object. This field will be present in the JSON response
	// if and only if `category` is equal to `check_transfer_deposit`.
	CheckTransferDeposit ACHTransferSimulationTransactionSourceCheckTransferDeposit `json:"check_transfer_deposit,required,nullable"`
	// A Check Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_intention`.
	CheckTransferIntention ACHTransferSimulationTransactionSourceCheckTransferIntention `json:"check_transfer_intention,required,nullable"`
	// A Check Transfer Stop Payment Request object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `check_transfer_stop_payment_request`.
	CheckTransferStopPaymentRequest ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequest `json:"check_transfer_stop_payment_request,required,nullable"`
	// A Fee Payment object. This field will be present in the JSON response if and
	// only if `category` is equal to `fee_payment`.
	FeePayment ACHTransferSimulationTransactionSourceFeePayment `json:"fee_payment,required,nullable"`
	// An Inbound ACH Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `inbound_ach_transfer`.
	InboundACHTransfer ACHTransferSimulationTransactionSourceInboundACHTransfer `json:"inbound_ach_transfer,required,nullable"`
	// An Inbound Check object. This field will be present in the JSON response if and
	// only if `category` is equal to `inbound_check`.
	InboundCheck ACHTransferSimulationTransactionSourceInboundCheck `json:"inbound_check,required,nullable"`
	// An Inbound International ACH Transfer object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `inbound_international_ach_transfer`.
	InboundInternationalACHTransfer ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer `json:"inbound_international_ach_transfer,required,nullable"`
	// An Inbound Real-Time Payments Transfer Confirmation object. This field will be
	// present in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_confirmation`.
	InboundRealTimePaymentsTransferConfirmation ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation `json:"inbound_real_time_payments_transfer_confirmation,required,nullable"`
	// An Inbound Wire Drawdown Payment object. This field will be present in the JSON
	// response if and only if `category` is equal to `inbound_wire_drawdown_payment`.
	InboundWireDrawdownPayment ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment `json:"inbound_wire_drawdown_payment,required,nullable"`
	// An Inbound Wire Drawdown Payment Reversal object. This field will be present in
	// the JSON response if and only if `category` is equal to
	// `inbound_wire_drawdown_payment_reversal`.
	InboundWireDrawdownPaymentReversal ACHTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal `json:"inbound_wire_drawdown_payment_reversal,required,nullable"`
	// An Inbound Wire Reversal object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_wire_reversal`.
	InboundWireReversal ACHTransferSimulationTransactionSourceInboundWireReversal `json:"inbound_wire_reversal,required,nullable"`
	// An Inbound Wire Transfer object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_wire_transfer`.
	InboundWireTransfer ACHTransferSimulationTransactionSourceInboundWireTransfer `json:"inbound_wire_transfer,required,nullable"`
	// An Interest Payment object. This field will be present in the JSON response if
	// and only if `category` is equal to `interest_payment`.
	InterestPayment ACHTransferSimulationTransactionSourceInterestPayment `json:"interest_payment,required,nullable"`
	// An Internal Source object. This field will be present in the JSON response if
	// and only if `category` is equal to `internal_source`.
	InternalSource ACHTransferSimulationTransactionSourceInternalSource `json:"internal_source,required,nullable"`
	// A Real-Time Payments Transfer Acknowledgement object. This field will be present
	// in the JSON response if and only if `category` is equal to
	// `real_time_payments_transfer_acknowledgement`.
	RealTimePaymentsTransferAcknowledgement ACHTransferSimulationTransactionSourceRealTimePaymentsTransferAcknowledgement `json:"real_time_payments_transfer_acknowledgement,required,nullable"`
	// A Sample Funds object. This field will be present in the JSON response if and
	// only if `category` is equal to `sample_funds`.
	SampleFunds ACHTransferSimulationTransactionSourceSampleFunds `json:"sample_funds,required,nullable"`
	// A Wire Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_intention`.
	WireTransferIntention ACHTransferSimulationTransactionSourceWireTransferIntention `json:"wire_transfer_intention,required,nullable"`
	// A Wire Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_rejection`.
	WireTransferRejection ACHTransferSimulationTransactionSourceWireTransferRejection `json:"wire_transfer_rejection,required,nullable"`
	JSON                  achTransferSimulationTransactionSourceJSON
}

// achTransferSimulationTransactionSourceJSON contains the JSON metadata for the
// struct [ACHTransferSimulationTransactionSource]
type achTransferSimulationTransactionSourceJSON struct {
	AccountTransferIntention                    apijson.Field
	ACHTransferIntention                        apijson.Field
	ACHTransferRejection                        apijson.Field
	ACHTransferReturn                           apijson.Field
	CardDisputeAcceptance                       apijson.Field
	CardRefund                                  apijson.Field
	CardRevenuePayment                          apijson.Field
	CardSettlement                              apijson.Field
	Category                                    apijson.Field
	CheckDepositAcceptance                      apijson.Field
	CheckDepositReturn                          apijson.Field
	CheckTransferDeposit                        apijson.Field
	CheckTransferIntention                      apijson.Field
	CheckTransferStopPaymentRequest             apijson.Field
	FeePayment                                  apijson.Field
	InboundACHTransfer                          apijson.Field
	InboundCheck                                apijson.Field
	InboundInternationalACHTransfer             apijson.Field
	InboundRealTimePaymentsTransferConfirmation apijson.Field
	InboundWireDrawdownPayment                  apijson.Field
	InboundWireDrawdownPaymentReversal          apijson.Field
	InboundWireReversal                         apijson.Field
	InboundWireTransfer                         apijson.Field
	InterestPayment                             apijson.Field
	InternalSource                              apijson.Field
	RealTimePaymentsTransferAcknowledgement     apijson.Field
	SampleFunds                                 apijson.Field
	WireTransferIntention                       apijson.Field
	WireTransferRejection                       apijson.Field
	raw                                         string
	ExtraFields                                 map[string]apijson.Field
}

func (r *ACHTransferSimulationTransactionSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// An Account Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `account_transfer_intention`.
type ACHTransferSimulationTransactionSourceAccountTransferIntention struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency ACHTransferSimulationTransactionSourceAccountTransferIntentionCurrency `json:"currency,required"`
	// The description you chose to give the transfer.
	Description string `json:"description,required"`
	// The identifier of the Account to where the Account Transfer was sent.
	DestinationAccountID string `json:"destination_account_id,required"`
	// The identifier of the Account from where the Account Transfer was sent.
	SourceAccountID string `json:"source_account_id,required"`
	// The identifier of the Account Transfer that led to this Pending Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       achTransferSimulationTransactionSourceAccountTransferIntentionJSON
}

// achTransferSimulationTransactionSourceAccountTransferIntentionJSON contains the
// JSON metadata for the struct
// [ACHTransferSimulationTransactionSourceAccountTransferIntention]
type achTransferSimulationTransactionSourceAccountTransferIntentionJSON struct {
	Amount               apijson.Field
	Currency             apijson.Field
	Description          apijson.Field
	DestinationAccountID apijson.Field
	SourceAccountID      apijson.Field
	TransferID           apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ACHTransferSimulationTransactionSourceAccountTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
// account currency.
type ACHTransferSimulationTransactionSourceAccountTransferIntentionCurrency string

const (
	// Canadian Dollar (CAD)
	ACHTransferSimulationTransactionSourceAccountTransferIntentionCurrencyCad ACHTransferSimulationTransactionSourceAccountTransferIntentionCurrency = "CAD"
	// Swiss Franc (CHF)
	ACHTransferSimulationTransactionSourceAccountTransferIntentionCurrencyChf ACHTransferSimulationTransactionSourceAccountTransferIntentionCurrency = "CHF"
	// Euro (EUR)
	ACHTransferSimulationTransactionSourceAccountTransferIntentionCurrencyEur ACHTransferSimulationTransactionSourceAccountTransferIntentionCurrency = "EUR"
	// British Pound (GBP)
	ACHTransferSimulationTransactionSourceAccountTransferIntentionCurrencyGbp ACHTransferSimulationTransactionSourceAccountTransferIntentionCurrency = "GBP"
	// Japanese Yen (JPY)
	ACHTransferSimulationTransactionSourceAccountTransferIntentionCurrencyJpy ACHTransferSimulationTransactionSourceAccountTransferIntentionCurrency = "JPY"
	// US Dollar (USD)
	ACHTransferSimulationTransactionSourceAccountTransferIntentionCurrencyUsd ACHTransferSimulationTransactionSourceAccountTransferIntentionCurrency = "USD"
)

// An ACH Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `ach_transfer_intention`.
type ACHTransferSimulationTransactionSourceACHTransferIntention struct {
	AccountNumber string `json:"account_number,required"`
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount              int64  `json:"amount,required"`
	RoutingNumber       string `json:"routing_number,required"`
	StatementDescriptor string `json:"statement_descriptor,required"`
	// The identifier of the ACH Transfer that led to this Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       achTransferSimulationTransactionSourceACHTransferIntentionJSON
}

// achTransferSimulationTransactionSourceACHTransferIntentionJSON contains the JSON
// metadata for the struct
// [ACHTransferSimulationTransactionSourceACHTransferIntention]
type achTransferSimulationTransactionSourceACHTransferIntentionJSON struct {
	AccountNumber       apijson.Field
	Amount              apijson.Field
	RoutingNumber       apijson.Field
	StatementDescriptor apijson.Field
	TransferID          apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ACHTransferSimulationTransactionSourceACHTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// An ACH Transfer Rejection object. This field will be present in the JSON
// response if and only if `category` is equal to `ach_transfer_rejection`.
type ACHTransferSimulationTransactionSourceACHTransferRejection struct {
	// The identifier of the ACH Transfer that led to this Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       achTransferSimulationTransactionSourceACHTransferRejectionJSON
}

// achTransferSimulationTransactionSourceACHTransferRejectionJSON contains the JSON
// metadata for the struct
// [ACHTransferSimulationTransactionSourceACHTransferRejection]
type achTransferSimulationTransactionSourceACHTransferRejectionJSON struct {
	TransferID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACHTransferSimulationTransactionSourceACHTransferRejection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// An ACH Transfer Return object. This field will be present in the JSON response
// if and only if `category` is equal to `ach_transfer_return`.
type ACHTransferSimulationTransactionSourceACHTransferReturn struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The three character ACH return code, in the range R01 to R85.
	RawReturnReasonCode string `json:"raw_return_reason_code,required"`
	// Why the ACH Transfer was returned.
	ReturnReasonCode ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode `json:"return_reason_code,required"`
	// The identifier of the Transaction associated with this return.
	TransactionID string `json:"transaction_id,required"`
	// The identifier of the ACH Transfer associated with this return.
	TransferID string `json:"transfer_id,required"`
	JSON       achTransferSimulationTransactionSourceACHTransferReturnJSON
}

// achTransferSimulationTransactionSourceACHTransferReturnJSON contains the JSON
// metadata for the struct
// [ACHTransferSimulationTransactionSourceACHTransferReturn]
type achTransferSimulationTransactionSourceACHTransferReturnJSON struct {
	CreatedAt           apijson.Field
	RawReturnReasonCode apijson.Field
	ReturnReasonCode    apijson.Field
	TransactionID       apijson.Field
	TransferID          apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ACHTransferSimulationTransactionSourceACHTransferReturn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Why the ACH Transfer was returned.
type ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode string

const (
	// Code R01. Insufficient funds in the source account.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeInsufficientFund ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "insufficient_fund"
	// Code R03. The account does not exist or the receiving bank was unable to locate
	// it.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeNoAccount ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "no_account"
	// Code R02. The account is closed.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeAccountClosed ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "account_closed"
	// Code R04. The account number is invalid at the receiving bank.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeInvalidAccountNumberStructure ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "invalid_account_number_structure"
	// Code R16. The account was frozen per the Office of Foreign Assets Control.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeAccountFrozenEntryReturnedPerOfacInstruction ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "account_frozen_entry_returned_per_ofac_instruction"
	// Code R23. The receiving bank account refused a credit transfer.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeCreditEntryRefusedByReceiver ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "credit_entry_refused_by_receiver"
	// Code R05. The receiving bank rejected because of an incorrect Standard Entry
	// Class code.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeUnauthorizedDebitToConsumerAccountUsingCorporateSecCode ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "unauthorized_debit_to_consumer_account_using_corporate_sec_code"
	// Code R29. The corporate customer reversed the transfer.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeCorporateCustomerAdvisedNotAuthorized ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "corporate_customer_advised_not_authorized"
	// Code R08. The receiving bank stopped payment on this transfer.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodePaymentStopped ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "payment_stopped"
	// Code R20. The receiving bank account does not perform transfers.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeNonTransactionAccount ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "non_transaction_account"
	// Code R09. The receiving bank account does not have enough available balance for
	// the transfer.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeUncollectedFunds ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "uncollected_funds"
	// Code R28. The routing number is incorrect.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeRoutingNumberCheckDigitError ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "routing_number_check_digit_error"
	// Code R10. The customer reversed the transfer.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "customer_advised_unauthorized_improper_ineligible_or_incomplete"
	// Code R19. The amount field is incorrect or too large.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeAmountFieldError ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "amount_field_error"
	// Code R07. The customer who initiated the transfer revoked authorization.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeAuthorizationRevokedByCustomer ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "authorization_revoked_by_customer"
	// Code R13. The routing number is invalid.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeInvalidACHRoutingNumber ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "invalid_ach_routing_number"
	// Code R17. The receiving bank is unable to process a field in the transfer.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeFileRecordEditCriteria ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "file_record_edit_criteria"
	// Code R45. The individual name field was invalid.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeEnrInvalidIndividualName ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "enr_invalid_individual_name"
	// Code R06. The originating financial institution reversed the transfer.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeReturnedPerOdfiRequest ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "returned_per_odfi_request"
	// Code R34. The receiving bank's regulatory supervisor has limited their
	// participation.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeLimitedParticipationDfi ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "limited_participation_dfi"
	// Code R85. The outbound international ACH transfer was incorrect.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeIncorrectlyCodedOutboundInternationalPayment ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "incorrectly_coded_outbound_international_payment"
	// Code R12. A rare return reason. The account was sold to another bank.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeAccountSoldToAnotherDfi ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "account_sold_to_another_dfi"
	// Code R25. The addenda record is incorrect or missing.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeAddendaError ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "addenda_error"
	// Code R15. A rare return reason. The account holder is deceased.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeBeneficiaryOrAccountHolderDeceased ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "beneficiary_or_account_holder_deceased"
	// Code R11. A rare return reason. The customer authorized some payment to the
	// sender, but this payment was not in error.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeCustomerAdvisedNotWithinAuthorizationTerms ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "customer_advised_not_within_authorization_terms"
	// Code R74. A rare return reason. Sent in response to a return that was returned
	// with code `field_error`. The latest return should include the corrected
	// field(s).
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeCorrectedReturn ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "corrected_return"
	// Code R24. A rare return reason. The receiving bank received an exact duplicate
	// entry with the same trace number and amount.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeDuplicateEntry ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "duplicate_entry"
	// Code R67. A rare return reason. The return this message refers to was a
	// duplicate.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeDuplicateReturn ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "duplicate_return"
	// Code R47. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeEnrDuplicateEnrollment ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "enr_duplicate_enrollment"
	// Code R43. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeEnrInvalidDfiAccountNumber ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "enr_invalid_dfi_account_number"
	// Code R44. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeEnrInvalidIndividualIDNumber ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "enr_invalid_individual_id_number"
	// Code R46. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeEnrInvalidRepresentativePayeeIndicator ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "enr_invalid_representative_payee_indicator"
	// Code R41. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeEnrInvalidTransactionCode ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "enr_invalid_transaction_code"
	// Code R40. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeEnrReturnOfEnrEntry ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "enr_return_of_enr_entry"
	// Code R42. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeEnrRoutingNumberCheckDigitError ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "enr_routing_number_check_digit_error"
	// Code R84. A rare return reason. The International ACH Transfer cannot be
	// processed by the gateway.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeEntryNotProcessedByGateway ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "entry_not_processed_by_gateway"
	// Code R69. A rare return reason. One or more of the fields in the ACH were
	// malformed.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeFieldError ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "field_error"
	// Code R83. A rare return reason. The Foreign receiving bank was unable to settle
	// this ACH transfer.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeForeignReceivingDfiUnableToSettle ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "foreign_receiving_dfi_unable_to_settle"
	// Code R80. A rare return reason. The International ACH Transfer is malformed.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeIatEntryCodingError ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "iat_entry_coding_error"
	// Code R18. A rare return reason. The ACH has an improper effective entry date
	// field.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeImproperEffectiveEntryDate ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "improper_effective_entry_date"
	// Code R39. A rare return reason. The source document related to this ACH, usually
	// an ACH check conversion, was presented to the bank.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeImproperSourceDocumentSourceDocumentPresented ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "improper_source_document_source_document_presented"
	// Code R21. A rare return reason. The Company ID field of the ACH was invalid.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeInvalidCompanyID ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "invalid_company_id"
	// Code R82. A rare return reason. The foreign receiving bank identifier for an
	// International ACH Transfer was invalid.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeInvalidForeignReceivingDfiIdentification ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "invalid_foreign_receiving_dfi_identification"
	// Code R22. A rare return reason. The Individual ID number field of the ACH was
	// invalid.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeInvalidIndividualIDNumber ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "invalid_individual_id_number"
	// Code R53. A rare return reason. Both the Represented Check ("RCK") entry and the
	// original check were presented to the bank.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeItemAndRckEntryPresentedForPayment ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "item_and_rck_entry_presented_for_payment"
	// Code R51. A rare return reason. The Represented Check ("RCK") entry is
	// ineligible.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeItemRelatedToRckEntryIsIneligible ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "item_related_to_rck_entry_is_ineligible"
	// Code R26. A rare return reason. The ACH is missing a required field.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeMandatoryFieldError ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "mandatory_field_error"
	// Code R71. A rare return reason. The receiving bank does not recognize the
	// routing number in a dishonored return entry.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeMisroutedDishonoredReturn ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "misrouted_dishonored_return"
	// Code R61. A rare return reason. The receiving bank does not recognize the
	// routing number in a return entry.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeMisroutedReturn ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "misrouted_return"
	// Code R76. A rare return reason. Sent in response to a return, the bank does not
	// find the errors alleged by the returning bank.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeNoErrorsFound ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "no_errors_found"
	// Code R77. A rare return reason. The receiving bank does not accept the return of
	// the erroneous debit. The funds are not available at the receiving bank.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeNonAcceptanceOfR62DishonoredReturn ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "non_acceptance_of_r62_dishonored_return"
	// Code R81. A rare return reason. The receiving bank does not accept International
	// ACH Transfers.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeNonParticipantInIatProgram ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "non_participant_in_iat_program"
	// Code R31. A rare return reason. A return that has been agreed to be accepted by
	// the receiving bank, despite falling outside of the usual return timeframe.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodePermissibleReturnEntry ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "permissible_return_entry"
	// Code R70. A rare return reason. The receiving bank had not approved this return.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodePermissibleReturnEntryNotAccepted ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "permissible_return_entry_not_accepted"
	// Code R32. A rare return reason. The receiving bank could not settle this
	// transaction.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeRdfiNonSettlement ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "rdfi_non_settlement"
	// Code R30. A rare return reason. The receiving bank does not accept Check
	// Truncation ACH transfers.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeRdfiParticipantInCheckTruncationProgram ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "rdfi_participant_in_check_truncation_program"
	// Code R14. A rare return reason. The payee is deceased.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeRepresentativePayeeDeceasedOrUnableToContinueInThatCapacity ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "representative_payee_deceased_or_unable_to_continue_in_that_capacity"
	// Code R75. A rare return reason. The originating bank disputes that an earlier
	// `duplicate_entry` return was actually a duplicate.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeReturnNotADuplicate ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "return_not_a_duplicate"
	// Code R62. A rare return reason. The originating bank made a mistake earlier and
	// this return corrects it.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeReturnOfErroneousOrReversingDebit ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "return_of_erroneous_or_reversing_debit"
	// Code R36. A rare return reason. Return of a malformed credit entry.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeReturnOfImproperCreditEntry ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "return_of_improper_credit_entry"
	// Code R35. A rare return reason. Return of a malformed debit entry.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeReturnOfImproperDebitEntry ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "return_of_improper_debit_entry"
	// Code R33. A rare return reason. Return of a Destroyed Check ("XKC") entry.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeReturnOfXckEntry ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "return_of_xck_entry"
	// Code R37. A rare return reason. The source document related to this ACH, usually
	// an ACH check conversion, was presented to the bank.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeSourceDocumentPresentedForPayment ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "source_document_presented_for_payment"
	// Code R50. A rare return reason. State law prevents the bank from accepting the
	// Represented Check ("RCK") entry.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeStateLawAffectingRckAcceptance ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "state_law_affecting_rck_acceptance"
	// Code R52. A rare return reason. A stop payment was issued on a Represented Check
	// ("RCK") entry.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeStopPaymentOnItemRelatedToRckEntry ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "stop_payment_on_item_related_to_rck_entry"
	// Code R38. A rare return reason. The source attached to the ACH, usually an ACH
	// check conversion, includes a stop payment.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeStopPaymentOnSourceDocument ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "stop_payment_on_source_document"
	// Code R73. A rare return reason. The bank receiving an `untimely_return` believes
	// it was on time.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeTimelyOriginalReturn ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "timely_original_return"
	// Code R27. A rare return reason. An ACH Return's trace number does not match an
	// originated ACH.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeTraceNumberError ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "trace_number_error"
	// Code R72. A rare return reason. The dishonored return was sent too late.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeUntimelyDishonoredReturn ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "untimely_dishonored_return"
	// Code R68. A rare return reason. The return was sent too late.
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeUntimelyReturn ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "untimely_return"
)

// A Card Dispute Acceptance object. This field will be present in the JSON
// response if and only if `category` is equal to `card_dispute_acceptance`.
type ACHTransferSimulationTransactionSourceCardDisputeAcceptance struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card Dispute was accepted.
	AcceptedAt time.Time `json:"accepted_at,required" format:"date-time"`
	// The identifier of the Card Dispute that was accepted.
	CardDisputeID string `json:"card_dispute_id,required"`
	// The identifier of the Transaction that was created to return the disputed funds
	// to your account.
	TransactionID string `json:"transaction_id,required"`
	JSON          achTransferSimulationTransactionSourceCardDisputeAcceptanceJSON
}

// achTransferSimulationTransactionSourceCardDisputeAcceptanceJSON contains the
// JSON metadata for the struct
// [ACHTransferSimulationTransactionSourceCardDisputeAcceptance]
type achTransferSimulationTransactionSourceCardDisputeAcceptanceJSON struct {
	AcceptedAt    apijson.Field
	CardDisputeID apijson.Field
	TransactionID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ACHTransferSimulationTransactionSourceCardDisputeAcceptance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Card Refund object. This field will be present in the JSON response if and
// only if `category` is equal to `card_refund`.
type ACHTransferSimulationTransactionSourceCardRefund struct {
	// The Card Refund identifier.
	ID string `json:"id,required"`
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency ACHTransferSimulationTransactionSourceCardRefundCurrency `json:"currency,required"`
	// The merchant identifier (commonly abbreviated as MID) of the merchant the card
	// is transacting with.
	MerchantAcceptorID string `json:"merchant_acceptor_id,required,nullable"`
	// The 4-digit MCC describing the merchant's business.
	MerchantCategoryCode string `json:"merchant_category_code,required"`
	// The city the merchant resides in.
	MerchantCity string `json:"merchant_city,required,nullable"`
	// The country the merchant resides in.
	MerchantCountry string `json:"merchant_country,required"`
	// The name of the merchant.
	MerchantName string `json:"merchant_name,required,nullable"`
	// The state the merchant resides in.
	MerchantState string `json:"merchant_state,required,nullable"`
	// Additional details about the card purchase, such as tax and industry-specific
	// fields.
	PurchaseDetails ACHTransferSimulationTransactionSourceCardRefundPurchaseDetails `json:"purchase_details,required,nullable"`
	// The identifier of the Transaction associated with this Transaction.
	TransactionID string `json:"transaction_id,required"`
	// A constant representing the object's type. For this resource it will always be
	// `card_refund`.
	Type ACHTransferSimulationTransactionSourceCardRefundType `json:"type,required"`
	JSON achTransferSimulationTransactionSourceCardRefundJSON
}

// achTransferSimulationTransactionSourceCardRefundJSON contains the JSON metadata
// for the struct [ACHTransferSimulationTransactionSourceCardRefund]
type achTransferSimulationTransactionSourceCardRefundJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	Currency             apijson.Field
	MerchantAcceptorID   apijson.Field
	MerchantCategoryCode apijson.Field
	MerchantCity         apijson.Field
	MerchantCountry      apijson.Field
	MerchantName         apijson.Field
	MerchantState        apijson.Field
	PurchaseDetails      apijson.Field
	TransactionID        apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ACHTransferSimulationTransactionSourceCardRefund) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
type ACHTransferSimulationTransactionSourceCardRefundCurrency string

const (
	// Canadian Dollar (CAD)
	ACHTransferSimulationTransactionSourceCardRefundCurrencyCad ACHTransferSimulationTransactionSourceCardRefundCurrency = "CAD"
	// Swiss Franc (CHF)
	ACHTransferSimulationTransactionSourceCardRefundCurrencyChf ACHTransferSimulationTransactionSourceCardRefundCurrency = "CHF"
	// Euro (EUR)
	ACHTransferSimulationTransactionSourceCardRefundCurrencyEur ACHTransferSimulationTransactionSourceCardRefundCurrency = "EUR"
	// British Pound (GBP)
	ACHTransferSimulationTransactionSourceCardRefundCurrencyGbp ACHTransferSimulationTransactionSourceCardRefundCurrency = "GBP"
	// Japanese Yen (JPY)
	ACHTransferSimulationTransactionSourceCardRefundCurrencyJpy ACHTransferSimulationTransactionSourceCardRefundCurrency = "JPY"
	// US Dollar (USD)
	ACHTransferSimulationTransactionSourceCardRefundCurrencyUsd ACHTransferSimulationTransactionSourceCardRefundCurrency = "USD"
)

// Additional details about the card purchase, such as tax and industry-specific
// fields.
type ACHTransferSimulationTransactionSourceCardRefundPurchaseDetails struct {
	// Fields specific to car rentals.
	CarRental ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRental `json:"car_rental,required,nullable"`
	// An identifier from the merchant for the customer or consumer.
	CustomerReferenceIdentifier string `json:"customer_reference_identifier,required,nullable"`
	// The state or provincial tax amount in minor units.
	LocalTaxAmount int64 `json:"local_tax_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the local tax
	// assessed.
	LocalTaxCurrency string `json:"local_tax_currency,required,nullable"`
	// Fields specific to lodging.
	Lodging ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodging `json:"lodging,required,nullable"`
	// The national tax amount in minor units.
	NationalTaxAmount int64 `json:"national_tax_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the local tax
	// assessed.
	NationalTaxCurrency string `json:"national_tax_currency,required,nullable"`
	// An identifier from the merchant for the purchase to the issuer and cardholder.
	PurchaseIdentifier string `json:"purchase_identifier,required,nullable"`
	// The format of the purchase identifier.
	PurchaseIdentifierFormat ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormat `json:"purchase_identifier_format,required,nullable"`
	// Fields specific to travel.
	Travel ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravel `json:"travel,required,nullable"`
	JSON   achTransferSimulationTransactionSourceCardRefundPurchaseDetailsJSON
}

// achTransferSimulationTransactionSourceCardRefundPurchaseDetailsJSON contains the
// JSON metadata for the struct
// [ACHTransferSimulationTransactionSourceCardRefundPurchaseDetails]
type achTransferSimulationTransactionSourceCardRefundPurchaseDetailsJSON struct {
	CarRental                   apijson.Field
	CustomerReferenceIdentifier apijson.Field
	LocalTaxAmount              apijson.Field
	LocalTaxCurrency            apijson.Field
	Lodging                     apijson.Field
	NationalTaxAmount           apijson.Field
	NationalTaxCurrency         apijson.Field
	PurchaseIdentifier          apijson.Field
	PurchaseIdentifierFormat    apijson.Field
	Travel                      apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *ACHTransferSimulationTransactionSourceCardRefundPurchaseDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Fields specific to car rentals.
type ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRental struct {
	// Code indicating the vehicle's class.
	CarClassCode string `json:"car_class_code,required,nullable"`
	// Date the customer picked up the car or, in the case of a no-show or pre-pay
	// transaction, the scheduled pick up date.
	CheckoutDate time.Time `json:"checkout_date,required,nullable" format:"date"`
	// Daily rate being charged for the vehicle.
	DailyRentalRateAmount int64 `json:"daily_rental_rate_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the daily rental
	// rate.
	DailyRentalRateCurrency string `json:"daily_rental_rate_currency,required,nullable"`
	// Number of days the vehicle was rented.
	DaysRented int64 `json:"days_rented,required,nullable"`
	// Additional charges (gas, late fee, etc.) being billed.
	ExtraCharges ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRentalExtraCharges `json:"extra_charges,required,nullable"`
	// Fuel charges for the vehicle.
	FuelChargesAmount int64 `json:"fuel_charges_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the fuel charges
	// assessed.
	FuelChargesCurrency string `json:"fuel_charges_currency,required,nullable"`
	// Any insurance being charged for the vehicle.
	InsuranceChargesAmount int64 `json:"insurance_charges_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the insurance
	// charges assessed.
	InsuranceChargesCurrency string `json:"insurance_charges_currency,required,nullable"`
	// An indicator that the cardholder is being billed for a reserved vehicle that was
	// not actually rented (that is, a "no-show" charge).
	NoShowIndicator ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRentalNoShowIndicator `json:"no_show_indicator,required,nullable"`
	// Charges for returning the vehicle at a different location than where it was
	// picked up.
	OneWayDropOffChargesAmount int64 `json:"one_way_drop_off_charges_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the one-way
	// drop-off charges assessed.
	OneWayDropOffChargesCurrency string `json:"one_way_drop_off_charges_currency,required,nullable"`
	// Name of the person renting the vehicle.
	RenterName string `json:"renter_name,required,nullable"`
	// Weekly rate being charged for the vehicle.
	WeeklyRentalRateAmount int64 `json:"weekly_rental_rate_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the weekly
	// rental rate.
	WeeklyRentalRateCurrency string `json:"weekly_rental_rate_currency,required,nullable"`
	JSON                     achTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRentalJSON
}

// achTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRentalJSON
// contains the JSON metadata for the struct
// [ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRental]
type achTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRentalJSON struct {
	CarClassCode                 apijson.Field
	CheckoutDate                 apijson.Field
	DailyRentalRateAmount        apijson.Field
	DailyRentalRateCurrency      apijson.Field
	DaysRented                   apijson.Field
	ExtraCharges                 apijson.Field
	FuelChargesAmount            apijson.Field
	FuelChargesCurrency          apijson.Field
	InsuranceChargesAmount       apijson.Field
	InsuranceChargesCurrency     apijson.Field
	NoShowIndicator              apijson.Field
	OneWayDropOffChargesAmount   apijson.Field
	OneWayDropOffChargesCurrency apijson.Field
	RenterName                   apijson.Field
	WeeklyRentalRateAmount       apijson.Field
	WeeklyRentalRateCurrency     apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRental) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Additional charges (gas, late fee, etc.) being billed.
type ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRentalExtraCharges string

const (
	// No extra charge
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRentalExtraChargesNoExtraCharge ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRentalExtraCharges = "no_extra_charge"
	// Gas
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRentalExtraChargesGas ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRentalExtraCharges = "gas"
	// Extra mileage
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRentalExtraChargesExtraMileage ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRentalExtraCharges = "extra_mileage"
	// Late return
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRentalExtraChargesLateReturn ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRentalExtraCharges = "late_return"
	// One way service fee
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRentalExtraChargesOneWayServiceFee ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRentalExtraCharges = "one_way_service_fee"
	// Parking violation
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRentalExtraChargesParkingViolation ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRentalExtraCharges = "parking_violation"
)

// An indicator that the cardholder is being billed for a reserved vehicle that was
// not actually rented (that is, a "no-show" charge).
type ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRentalNoShowIndicator string

const (
	// Not applicable
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRentalNoShowIndicatorNotApplicable ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRentalNoShowIndicator = "not_applicable"
	// No show for specialized vehicle
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRentalNoShowIndicatorNoShowForSpecializedVehicle ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsCarRentalNoShowIndicator = "no_show_for_specialized_vehicle"
)

// Fields specific to lodging.
type ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodging struct {
	// Date the customer checked in.
	CheckInDate time.Time `json:"check_in_date,required,nullable" format:"date"`
	// Daily rate being charged for the room.
	DailyRoomRateAmount int64 `json:"daily_room_rate_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the daily room
	// rate.
	DailyRoomRateCurrency string `json:"daily_room_rate_currency,required,nullable"`
	// Additional charges (phone, late check-out, etc.) being billed.
	ExtraCharges ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingExtraCharges `json:"extra_charges,required,nullable"`
	// Folio cash advances for the room.
	FolioCashAdvancesAmount int64 `json:"folio_cash_advances_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the folio cash
	// advances.
	FolioCashAdvancesCurrency string `json:"folio_cash_advances_currency,required,nullable"`
	// Food and beverage charges for the room.
	FoodBeverageChargesAmount int64 `json:"food_beverage_charges_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the food and
	// beverage charges.
	FoodBeverageChargesCurrency string `json:"food_beverage_charges_currency,required,nullable"`
	// Indicator that the cardholder is being billed for a reserved room that was not
	// actually used.
	NoShowIndicator ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingNoShowIndicator `json:"no_show_indicator,required,nullable"`
	// Prepaid expenses being charged for the room.
	PrepaidExpensesAmount int64 `json:"prepaid_expenses_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the prepaid
	// expenses.
	PrepaidExpensesCurrency string `json:"prepaid_expenses_currency,required,nullable"`
	// Number of nights the room was rented.
	RoomNights int64 `json:"room_nights,required,nullable"`
	// Total room tax being charged.
	TotalRoomTaxAmount int64 `json:"total_room_tax_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the total room
	// tax.
	TotalRoomTaxCurrency string `json:"total_room_tax_currency,required,nullable"`
	// Total tax being charged for the room.
	TotalTaxAmount int64 `json:"total_tax_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the total tax
	// assessed.
	TotalTaxCurrency string `json:"total_tax_currency,required,nullable"`
	JSON             achTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingJSON
}

// achTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingJSON
// contains the JSON metadata for the struct
// [ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodging]
type achTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingJSON struct {
	CheckInDate                 apijson.Field
	DailyRoomRateAmount         apijson.Field
	DailyRoomRateCurrency       apijson.Field
	ExtraCharges                apijson.Field
	FolioCashAdvancesAmount     apijson.Field
	FolioCashAdvancesCurrency   apijson.Field
	FoodBeverageChargesAmount   apijson.Field
	FoodBeverageChargesCurrency apijson.Field
	NoShowIndicator             apijson.Field
	PrepaidExpensesAmount       apijson.Field
	PrepaidExpensesCurrency     apijson.Field
	RoomNights                  apijson.Field
	TotalRoomTaxAmount          apijson.Field
	TotalRoomTaxCurrency        apijson.Field
	TotalTaxAmount              apijson.Field
	TotalTaxCurrency            apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Additional charges (phone, late check-out, etc.) being billed.
type ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingExtraCharges string

const (
	// No extra charge
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingExtraChargesNoExtraCharge ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingExtraCharges = "no_extra_charge"
	// Restaurant
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingExtraChargesRestaurant ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingExtraCharges = "restaurant"
	// Gift shop
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingExtraChargesGiftShop ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingExtraCharges = "gift_shop"
	// Mini bar
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingExtraChargesMiniBar ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingExtraCharges = "mini_bar"
	// Telephone
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingExtraChargesTelephone ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingExtraCharges = "telephone"
	// Other
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingExtraChargesOther ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingExtraCharges = "other"
	// Laundry
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingExtraChargesLaundry ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingExtraCharges = "laundry"
)

// Indicator that the cardholder is being billed for a reserved room that was not
// actually used.
type ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingNoShowIndicator string

const (
	// Not applicable
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingNoShowIndicatorNotApplicable ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingNoShowIndicator = "not_applicable"
	// No show
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingNoShowIndicatorNoShow ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsLodgingNoShowIndicator = "no_show"
)

// The format of the purchase identifier.
type ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormat string

const (
	// Free text
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormatFreeText ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormat = "free_text"
	// Order number
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormatOrderNumber ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormat = "order_number"
	// Rental agreement number
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormatRentalAgreementNumber ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormat = "rental_agreement_number"
	// Hotel folio number
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormatHotelFolioNumber ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormat = "hotel_folio_number"
	// Invoice number
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormatInvoiceNumber ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsPurchaseIdentifierFormat = "invoice_number"
)

// Fields specific to travel.
type ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravel struct {
	// Ancillary purchases in addition to the airfare.
	Ancillary ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillary `json:"ancillary,required,nullable"`
	// Indicates the computerized reservation system used to book the ticket.
	ComputerizedReservationSystem string `json:"computerized_reservation_system,required,nullable"`
	// Indicates the reason for a credit to the cardholder.
	CreditReasonIndicator ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicator `json:"credit_reason_indicator,required,nullable"`
	// Date of departure.
	DepartureDate time.Time `json:"departure_date,required,nullable" format:"date"`
	// Code for the originating city or airport.
	OriginationCityAirportCode string `json:"origination_city_airport_code,required,nullable"`
	// Name of the passenger.
	PassengerName string `json:"passenger_name,required,nullable"`
	// Indicates whether this ticket is non-refundable.
	RestrictedTicketIndicator ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelRestrictedTicketIndicator `json:"restricted_ticket_indicator,required,nullable"`
	// Indicates why a ticket was changed.
	TicketChangeIndicator ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelTicketChangeIndicator `json:"ticket_change_indicator,required,nullable"`
	// Ticket number.
	TicketNumber string `json:"ticket_number,required,nullable"`
	// Code for the travel agency if the ticket was issued by a travel agency.
	TravelAgencyCode string `json:"travel_agency_code,required,nullable"`
	// Name of the travel agency if the ticket was issued by a travel agency.
	TravelAgencyName string `json:"travel_agency_name,required,nullable"`
	// Fields specific to each leg of the journey.
	TripLegs []ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelTripLeg `json:"trip_legs,required,nullable"`
	JSON     achTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelJSON
}

// achTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelJSON
// contains the JSON metadata for the struct
// [ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravel]
type achTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelJSON struct {
	Ancillary                     apijson.Field
	ComputerizedReservationSystem apijson.Field
	CreditReasonIndicator         apijson.Field
	DepartureDate                 apijson.Field
	OriginationCityAirportCode    apijson.Field
	PassengerName                 apijson.Field
	RestrictedTicketIndicator     apijson.Field
	TicketChangeIndicator         apijson.Field
	TicketNumber                  apijson.Field
	TravelAgencyCode              apijson.Field
	TravelAgencyName              apijson.Field
	TripLegs                      apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Ancillary purchases in addition to the airfare.
type ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillary struct {
	// If this purchase has a connection or relationship to another purchase, such as a
	// baggage fee for a passenger transport ticket, this field should contain the
	// ticket document number for the other purchase.
	ConnectedTicketDocumentNumber string `json:"connected_ticket_document_number,required,nullable"`
	// Indicates the reason for a credit to the cardholder.
	CreditReasonIndicator ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicator `json:"credit_reason_indicator,required,nullable"`
	// Name of the passenger or description of the ancillary purchase.
	PassengerNameOrDescription string `json:"passenger_name_or_description,required,nullable"`
	// Additional travel charges, such as baggage fees.
	Services []ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryService `json:"services,required"`
	// Ticket document number.
	TicketDocumentNumber string `json:"ticket_document_number,required,nullable"`
	JSON                 achTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryJSON
}

// achTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryJSON
// contains the JSON metadata for the struct
// [ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillary]
type achTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryJSON struct {
	ConnectedTicketDocumentNumber apijson.Field
	CreditReasonIndicator         apijson.Field
	PassengerNameOrDescription    apijson.Field
	Services                      apijson.Field
	TicketDocumentNumber          apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the reason for a credit to the cardholder.
type ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicator string

const (
	// No credit
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicatorNoCredit ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicator = "no_credit"
	// Passenger transport ancillary purchase cancellation
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicatorPassengerTransportAncillaryPurchaseCancellation ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicator = "passenger_transport_ancillary_purchase_cancellation"
	// Airline ticket and passenger transport ancillary purchase cancellation
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicatorAirlineTicketAndPassengerTransportAncillaryPurchaseCancellation ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicator = "airline_ticket_and_passenger_transport_ancillary_purchase_cancellation"
	// Other
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicatorOther ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryCreditReasonIndicator = "other"
)

type ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryService struct {
	// Category of the ancillary service.
	Category ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory `json:"category,required,nullable"`
	// Sub-category of the ancillary service, free-form.
	SubCategory string `json:"sub_category,required,nullable"`
	JSON        achTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServiceJSON
}

// achTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServiceJSON
// contains the JSON metadata for the struct
// [ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryService]
type achTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServiceJSON struct {
	Category    apijson.Field
	SubCategory apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Category of the ancillary service.
type ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory string

const (
	// None
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryNone ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "none"
	// Bundled service
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryBundledService ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "bundled_service"
	// Baggage fee
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryBaggageFee ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "baggage_fee"
	// Change fee
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryChangeFee ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "change_fee"
	// Cargo
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryCargo ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "cargo"
	// Carbon offset
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryCarbonOffset ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "carbon_offset"
	// Frequent flyer
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryFrequentFlyer ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "frequent_flyer"
	// Gift card
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryGiftCard ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "gift_card"
	// Ground transport
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryGroundTransport ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "ground_transport"
	// In-flight entertainment
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryInFlightEntertainment ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "in_flight_entertainment"
	// Lounge
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryLounge ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "lounge"
	// Medical
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryMedical ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "medical"
	// Meal beverage
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryMealBeverage ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "meal_beverage"
	// Other
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryOther ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "other"
	// Passenger assist fee
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryPassengerAssistFee ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "passenger_assist_fee"
	// Pets
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryPets ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "pets"
	// Seat fees
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategorySeatFees ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "seat_fees"
	// Standby
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryStandby ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "standby"
	// Service fee
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryServiceFee ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "service_fee"
	// Store
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryStore ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "store"
	// Travel service
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryTravelService ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "travel_service"
	// Unaccompanied travel
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryUnaccompaniedTravel ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "unaccompanied_travel"
	// Upgrades
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryUpgrades ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "upgrades"
	// Wi-fi
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategoryWifi ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelAncillaryServicesCategory = "wifi"
)

// Indicates the reason for a credit to the cardholder.
type ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicator string

const (
	// No credit
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicatorNoCredit ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicator = "no_credit"
	// Passenger transport ancillary purchase cancellation
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicatorPassengerTransportAncillaryPurchaseCancellation ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicator = "passenger_transport_ancillary_purchase_cancellation"
	// Airline ticket and passenger transport ancillary purchase cancellation
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicatorAirlineTicketAndPassengerTransportAncillaryPurchaseCancellation ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicator = "airline_ticket_and_passenger_transport_ancillary_purchase_cancellation"
	// Airline ticket cancellation
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicatorAirlineTicketCancellation ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicator = "airline_ticket_cancellation"
	// Other
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicatorOther ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicator = "other"
	// Partial refund of airline ticket
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicatorPartialRefundOfAirlineTicket ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelCreditReasonIndicator = "partial_refund_of_airline_ticket"
)

// Indicates whether this ticket is non-refundable.
type ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelRestrictedTicketIndicator string

const (
	// No restrictions
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelRestrictedTicketIndicatorNoRestrictions ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelRestrictedTicketIndicator = "no_restrictions"
	// Restricted non-refundable ticket
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelRestrictedTicketIndicatorRestrictedNonRefundableTicket ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelRestrictedTicketIndicator = "restricted_non_refundable_ticket"
)

// Indicates why a ticket was changed.
type ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelTicketChangeIndicator string

const (
	// None
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelTicketChangeIndicatorNone ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelTicketChangeIndicator = "none"
	// Change to existing ticket
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelTicketChangeIndicatorChangeToExistingTicket ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelTicketChangeIndicator = "change_to_existing_ticket"
	// New ticket
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelTicketChangeIndicatorNewTicket ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelTicketChangeIndicator = "new_ticket"
)

type ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelTripLeg struct {
	// Carrier code (e.g., United Airlines, Jet Blue, etc.).
	CarrierCode string `json:"carrier_code,required,nullable"`
	// Code for the destination city or airport.
	DestinationCityAirportCode string `json:"destination_city_airport_code,required,nullable"`
	// Fare basis code.
	FareBasisCode string `json:"fare_basis_code,required,nullable"`
	// Flight number.
	FlightNumber string `json:"flight_number,required,nullable"`
	// Service class (e.g., first class, business class, etc.).
	ServiceClass string `json:"service_class,required,nullable"`
	// Indicates whether a stopover is allowed on this ticket.
	StopOverCode ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelTripLegsStopOverCode `json:"stop_over_code,required,nullable"`
	JSON         achTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelTripLegJSON
}

// achTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelTripLegJSON
// contains the JSON metadata for the struct
// [ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelTripLeg]
type achTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelTripLegJSON struct {
	CarrierCode                apijson.Field
	DestinationCityAirportCode apijson.Field
	FareBasisCode              apijson.Field
	FlightNumber               apijson.Field
	ServiceClass               apijson.Field
	StopOverCode               apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelTripLeg) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates whether a stopover is allowed on this ticket.
type ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelTripLegsStopOverCode string

const (
	// None
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelTripLegsStopOverCodeNone ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelTripLegsStopOverCode = "none"
	// Stop over allowed
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelTripLegsStopOverCodeStopOverAllowed ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelTripLegsStopOverCode = "stop_over_allowed"
	// Stop over not allowed
	ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelTripLegsStopOverCodeStopOverNotAllowed ACHTransferSimulationTransactionSourceCardRefundPurchaseDetailsTravelTripLegsStopOverCode = "stop_over_not_allowed"
)

// A constant representing the object's type. For this resource it will always be
// `card_refund`.
type ACHTransferSimulationTransactionSourceCardRefundType string

const (
	ACHTransferSimulationTransactionSourceCardRefundTypeCardRefund ACHTransferSimulationTransactionSourceCardRefundType = "card_refund"
)

// A Card Revenue Payment object. This field will be present in the JSON response
// if and only if `category` is equal to `card_revenue_payment`.
type ACHTransferSimulationTransactionSourceCardRevenuePayment struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency ACHTransferSimulationTransactionSourceCardRevenuePaymentCurrency `json:"currency,required"`
	// The end of the period for which this transaction paid interest.
	PeriodEnd time.Time `json:"period_end,required" format:"date-time"`
	// The start of the period for which this transaction paid interest.
	PeriodStart time.Time `json:"period_start,required" format:"date-time"`
	// The account the card belonged to.
	TransactedOnAccountID string `json:"transacted_on_account_id,required,nullable"`
	JSON                  achTransferSimulationTransactionSourceCardRevenuePaymentJSON
}

// achTransferSimulationTransactionSourceCardRevenuePaymentJSON contains the JSON
// metadata for the struct
// [ACHTransferSimulationTransactionSourceCardRevenuePayment]
type achTransferSimulationTransactionSourceCardRevenuePaymentJSON struct {
	Amount                apijson.Field
	Currency              apijson.Field
	PeriodEnd             apijson.Field
	PeriodStart           apijson.Field
	TransactedOnAccountID apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ACHTransferSimulationTransactionSourceCardRevenuePayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
// currency.
type ACHTransferSimulationTransactionSourceCardRevenuePaymentCurrency string

const (
	// Canadian Dollar (CAD)
	ACHTransferSimulationTransactionSourceCardRevenuePaymentCurrencyCad ACHTransferSimulationTransactionSourceCardRevenuePaymentCurrency = "CAD"
	// Swiss Franc (CHF)
	ACHTransferSimulationTransactionSourceCardRevenuePaymentCurrencyChf ACHTransferSimulationTransactionSourceCardRevenuePaymentCurrency = "CHF"
	// Euro (EUR)
	ACHTransferSimulationTransactionSourceCardRevenuePaymentCurrencyEur ACHTransferSimulationTransactionSourceCardRevenuePaymentCurrency = "EUR"
	// British Pound (GBP)
	ACHTransferSimulationTransactionSourceCardRevenuePaymentCurrencyGbp ACHTransferSimulationTransactionSourceCardRevenuePaymentCurrency = "GBP"
	// Japanese Yen (JPY)
	ACHTransferSimulationTransactionSourceCardRevenuePaymentCurrencyJpy ACHTransferSimulationTransactionSourceCardRevenuePaymentCurrency = "JPY"
	// US Dollar (USD)
	ACHTransferSimulationTransactionSourceCardRevenuePaymentCurrencyUsd ACHTransferSimulationTransactionSourceCardRevenuePaymentCurrency = "USD"
)

// A Card Settlement object. This field will be present in the JSON response if and
// only if `category` is equal to `card_settlement`.
type ACHTransferSimulationTransactionSourceCardSettlement struct {
	// The Card Settlement identifier.
	ID string `json:"id,required"`
	// The amount in the minor unit of the transaction's settlement currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The Card Authorization that was created prior to this Card Settlement, if on
	// exists.
	CardAuthorization string `json:"card_authorization,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's settlement currency.
	Currency ACHTransferSimulationTransactionSourceCardSettlementCurrency `json:"currency,required"`
	// The merchant identifier (commonly abbreviated as MID) of the merchant the card
	// is transacting with.
	MerchantAcceptorID string `json:"merchant_acceptor_id,required,nullable"`
	// The 4-digit MCC describing the merchant's business.
	MerchantCategoryCode string `json:"merchant_category_code,required"`
	// The city the merchant resides in.
	MerchantCity string `json:"merchant_city,required,nullable"`
	// The country the merchant resides in.
	MerchantCountry string `json:"merchant_country,required"`
	// The name of the merchant.
	MerchantName string `json:"merchant_name,required,nullable"`
	// The state the merchant resides in.
	MerchantState string `json:"merchant_state,required,nullable"`
	// The identifier of the Pending Transaction associated with this Transaction.
	PendingTransactionID string `json:"pending_transaction_id,required,nullable"`
	// The amount in the minor unit of the transaction's presentment currency.
	PresentmentAmount int64 `json:"presentment_amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's presentment currency.
	PresentmentCurrency string `json:"presentment_currency,required"`
	// Additional details about the card purchase, such as tax and industry-specific
	// fields.
	PurchaseDetails ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetails `json:"purchase_details,required,nullable"`
	// The identifier of the Transaction associated with this Transaction.
	TransactionID string `json:"transaction_id,required"`
	// A constant representing the object's type. For this resource it will always be
	// `card_settlement`.
	Type ACHTransferSimulationTransactionSourceCardSettlementType `json:"type,required"`
	JSON achTransferSimulationTransactionSourceCardSettlementJSON
}

// achTransferSimulationTransactionSourceCardSettlementJSON contains the JSON
// metadata for the struct [ACHTransferSimulationTransactionSourceCardSettlement]
type achTransferSimulationTransactionSourceCardSettlementJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	CardAuthorization    apijson.Field
	Currency             apijson.Field
	MerchantAcceptorID   apijson.Field
	MerchantCategoryCode apijson.Field
	MerchantCity         apijson.Field
	MerchantCountry      apijson.Field
	MerchantName         apijson.Field
	MerchantState        apijson.Field
	PendingTransactionID apijson.Field
	PresentmentAmount    apijson.Field
	PresentmentCurrency  apijson.Field
	PurchaseDetails      apijson.Field
	TransactionID        apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ACHTransferSimulationTransactionSourceCardSettlement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's settlement currency.
type ACHTransferSimulationTransactionSourceCardSettlementCurrency string

const (
	// Canadian Dollar (CAD)
	ACHTransferSimulationTransactionSourceCardSettlementCurrencyCad ACHTransferSimulationTransactionSourceCardSettlementCurrency = "CAD"
	// Swiss Franc (CHF)
	ACHTransferSimulationTransactionSourceCardSettlementCurrencyChf ACHTransferSimulationTransactionSourceCardSettlementCurrency = "CHF"
	// Euro (EUR)
	ACHTransferSimulationTransactionSourceCardSettlementCurrencyEur ACHTransferSimulationTransactionSourceCardSettlementCurrency = "EUR"
	// British Pound (GBP)
	ACHTransferSimulationTransactionSourceCardSettlementCurrencyGbp ACHTransferSimulationTransactionSourceCardSettlementCurrency = "GBP"
	// Japanese Yen (JPY)
	ACHTransferSimulationTransactionSourceCardSettlementCurrencyJpy ACHTransferSimulationTransactionSourceCardSettlementCurrency = "JPY"
	// US Dollar (USD)
	ACHTransferSimulationTransactionSourceCardSettlementCurrencyUsd ACHTransferSimulationTransactionSourceCardSettlementCurrency = "USD"
)

// Additional details about the card purchase, such as tax and industry-specific
// fields.
type ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetails struct {
	// Fields specific to car rentals.
	CarRental ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRental `json:"car_rental,required,nullable"`
	// An identifier from the merchant for the customer or consumer.
	CustomerReferenceIdentifier string `json:"customer_reference_identifier,required,nullable"`
	// The state or provincial tax amount in minor units.
	LocalTaxAmount int64 `json:"local_tax_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the local tax
	// assessed.
	LocalTaxCurrency string `json:"local_tax_currency,required,nullable"`
	// Fields specific to lodging.
	Lodging ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodging `json:"lodging,required,nullable"`
	// The national tax amount in minor units.
	NationalTaxAmount int64 `json:"national_tax_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the local tax
	// assessed.
	NationalTaxCurrency string `json:"national_tax_currency,required,nullable"`
	// An identifier from the merchant for the purchase to the issuer and cardholder.
	PurchaseIdentifier string `json:"purchase_identifier,required,nullable"`
	// The format of the purchase identifier.
	PurchaseIdentifierFormat ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormat `json:"purchase_identifier_format,required,nullable"`
	// Fields specific to travel.
	Travel ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravel `json:"travel,required,nullable"`
	JSON   achTransferSimulationTransactionSourceCardSettlementPurchaseDetailsJSON
}

// achTransferSimulationTransactionSourceCardSettlementPurchaseDetailsJSON contains
// the JSON metadata for the struct
// [ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetails]
type achTransferSimulationTransactionSourceCardSettlementPurchaseDetailsJSON struct {
	CarRental                   apijson.Field
	CustomerReferenceIdentifier apijson.Field
	LocalTaxAmount              apijson.Field
	LocalTaxCurrency            apijson.Field
	Lodging                     apijson.Field
	NationalTaxAmount           apijson.Field
	NationalTaxCurrency         apijson.Field
	PurchaseIdentifier          apijson.Field
	PurchaseIdentifierFormat    apijson.Field
	Travel                      apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Fields specific to car rentals.
type ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRental struct {
	// Code indicating the vehicle's class.
	CarClassCode string `json:"car_class_code,required,nullable"`
	// Date the customer picked up the car or, in the case of a no-show or pre-pay
	// transaction, the scheduled pick up date.
	CheckoutDate time.Time `json:"checkout_date,required,nullable" format:"date"`
	// Daily rate being charged for the vehicle.
	DailyRentalRateAmount int64 `json:"daily_rental_rate_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the daily rental
	// rate.
	DailyRentalRateCurrency string `json:"daily_rental_rate_currency,required,nullable"`
	// Number of days the vehicle was rented.
	DaysRented int64 `json:"days_rented,required,nullable"`
	// Additional charges (gas, late fee, etc.) being billed.
	ExtraCharges ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRentalExtraCharges `json:"extra_charges,required,nullable"`
	// Fuel charges for the vehicle.
	FuelChargesAmount int64 `json:"fuel_charges_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the fuel charges
	// assessed.
	FuelChargesCurrency string `json:"fuel_charges_currency,required,nullable"`
	// Any insurance being charged for the vehicle.
	InsuranceChargesAmount int64 `json:"insurance_charges_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the insurance
	// charges assessed.
	InsuranceChargesCurrency string `json:"insurance_charges_currency,required,nullable"`
	// An indicator that the cardholder is being billed for a reserved vehicle that was
	// not actually rented (that is, a "no-show" charge).
	NoShowIndicator ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRentalNoShowIndicator `json:"no_show_indicator,required,nullable"`
	// Charges for returning the vehicle at a different location than where it was
	// picked up.
	OneWayDropOffChargesAmount int64 `json:"one_way_drop_off_charges_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the one-way
	// drop-off charges assessed.
	OneWayDropOffChargesCurrency string `json:"one_way_drop_off_charges_currency,required,nullable"`
	// Name of the person renting the vehicle.
	RenterName string `json:"renter_name,required,nullable"`
	// Weekly rate being charged for the vehicle.
	WeeklyRentalRateAmount int64 `json:"weekly_rental_rate_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the weekly
	// rental rate.
	WeeklyRentalRateCurrency string `json:"weekly_rental_rate_currency,required,nullable"`
	JSON                     achTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRentalJSON
}

// achTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRentalJSON
// contains the JSON metadata for the struct
// [ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRental]
type achTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRentalJSON struct {
	CarClassCode                 apijson.Field
	CheckoutDate                 apijson.Field
	DailyRentalRateAmount        apijson.Field
	DailyRentalRateCurrency      apijson.Field
	DaysRented                   apijson.Field
	ExtraCharges                 apijson.Field
	FuelChargesAmount            apijson.Field
	FuelChargesCurrency          apijson.Field
	InsuranceChargesAmount       apijson.Field
	InsuranceChargesCurrency     apijson.Field
	NoShowIndicator              apijson.Field
	OneWayDropOffChargesAmount   apijson.Field
	OneWayDropOffChargesCurrency apijson.Field
	RenterName                   apijson.Field
	WeeklyRentalRateAmount       apijson.Field
	WeeklyRentalRateCurrency     apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRental) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Additional charges (gas, late fee, etc.) being billed.
type ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRentalExtraCharges string

const (
	// No extra charge
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRentalExtraChargesNoExtraCharge ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRentalExtraCharges = "no_extra_charge"
	// Gas
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRentalExtraChargesGas ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRentalExtraCharges = "gas"
	// Extra mileage
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRentalExtraChargesExtraMileage ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRentalExtraCharges = "extra_mileage"
	// Late return
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRentalExtraChargesLateReturn ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRentalExtraCharges = "late_return"
	// One way service fee
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRentalExtraChargesOneWayServiceFee ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRentalExtraCharges = "one_way_service_fee"
	// Parking violation
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRentalExtraChargesParkingViolation ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRentalExtraCharges = "parking_violation"
)

// An indicator that the cardholder is being billed for a reserved vehicle that was
// not actually rented (that is, a "no-show" charge).
type ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRentalNoShowIndicator string

const (
	// Not applicable
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRentalNoShowIndicatorNotApplicable ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRentalNoShowIndicator = "not_applicable"
	// No show for specialized vehicle
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRentalNoShowIndicatorNoShowForSpecializedVehicle ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsCarRentalNoShowIndicator = "no_show_for_specialized_vehicle"
)

// Fields specific to lodging.
type ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodging struct {
	// Date the customer checked in.
	CheckInDate time.Time `json:"check_in_date,required,nullable" format:"date"`
	// Daily rate being charged for the room.
	DailyRoomRateAmount int64 `json:"daily_room_rate_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the daily room
	// rate.
	DailyRoomRateCurrency string `json:"daily_room_rate_currency,required,nullable"`
	// Additional charges (phone, late check-out, etc.) being billed.
	ExtraCharges ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingExtraCharges `json:"extra_charges,required,nullable"`
	// Folio cash advances for the room.
	FolioCashAdvancesAmount int64 `json:"folio_cash_advances_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the folio cash
	// advances.
	FolioCashAdvancesCurrency string `json:"folio_cash_advances_currency,required,nullable"`
	// Food and beverage charges for the room.
	FoodBeverageChargesAmount int64 `json:"food_beverage_charges_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the food and
	// beverage charges.
	FoodBeverageChargesCurrency string `json:"food_beverage_charges_currency,required,nullable"`
	// Indicator that the cardholder is being billed for a reserved room that was not
	// actually used.
	NoShowIndicator ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingNoShowIndicator `json:"no_show_indicator,required,nullable"`
	// Prepaid expenses being charged for the room.
	PrepaidExpensesAmount int64 `json:"prepaid_expenses_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the prepaid
	// expenses.
	PrepaidExpensesCurrency string `json:"prepaid_expenses_currency,required,nullable"`
	// Number of nights the room was rented.
	RoomNights int64 `json:"room_nights,required,nullable"`
	// Total room tax being charged.
	TotalRoomTaxAmount int64 `json:"total_room_tax_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the total room
	// tax.
	TotalRoomTaxCurrency string `json:"total_room_tax_currency,required,nullable"`
	// Total tax being charged for the room.
	TotalTaxAmount int64 `json:"total_tax_amount,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the total tax
	// assessed.
	TotalTaxCurrency string `json:"total_tax_currency,required,nullable"`
	JSON             achTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingJSON
}

// achTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingJSON
// contains the JSON metadata for the struct
// [ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodging]
type achTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingJSON struct {
	CheckInDate                 apijson.Field
	DailyRoomRateAmount         apijson.Field
	DailyRoomRateCurrency       apijson.Field
	ExtraCharges                apijson.Field
	FolioCashAdvancesAmount     apijson.Field
	FolioCashAdvancesCurrency   apijson.Field
	FoodBeverageChargesAmount   apijson.Field
	FoodBeverageChargesCurrency apijson.Field
	NoShowIndicator             apijson.Field
	PrepaidExpensesAmount       apijson.Field
	PrepaidExpensesCurrency     apijson.Field
	RoomNights                  apijson.Field
	TotalRoomTaxAmount          apijson.Field
	TotalRoomTaxCurrency        apijson.Field
	TotalTaxAmount              apijson.Field
	TotalTaxCurrency            apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Additional charges (phone, late check-out, etc.) being billed.
type ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingExtraCharges string

const (
	// No extra charge
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingExtraChargesNoExtraCharge ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingExtraCharges = "no_extra_charge"
	// Restaurant
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingExtraChargesRestaurant ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingExtraCharges = "restaurant"
	// Gift shop
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingExtraChargesGiftShop ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingExtraCharges = "gift_shop"
	// Mini bar
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingExtraChargesMiniBar ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingExtraCharges = "mini_bar"
	// Telephone
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingExtraChargesTelephone ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingExtraCharges = "telephone"
	// Other
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingExtraChargesOther ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingExtraCharges = "other"
	// Laundry
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingExtraChargesLaundry ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingExtraCharges = "laundry"
)

// Indicator that the cardholder is being billed for a reserved room that was not
// actually used.
type ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingNoShowIndicator string

const (
	// Not applicable
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingNoShowIndicatorNotApplicable ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingNoShowIndicator = "not_applicable"
	// No show
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingNoShowIndicatorNoShow ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsLodgingNoShowIndicator = "no_show"
)

// The format of the purchase identifier.
type ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormat string

const (
	// Free text
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormatFreeText ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormat = "free_text"
	// Order number
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormatOrderNumber ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormat = "order_number"
	// Rental agreement number
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormatRentalAgreementNumber ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormat = "rental_agreement_number"
	// Hotel folio number
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormatHotelFolioNumber ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormat = "hotel_folio_number"
	// Invoice number
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormatInvoiceNumber ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsPurchaseIdentifierFormat = "invoice_number"
)

// Fields specific to travel.
type ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravel struct {
	// Ancillary purchases in addition to the airfare.
	Ancillary ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillary `json:"ancillary,required,nullable"`
	// Indicates the computerized reservation system used to book the ticket.
	ComputerizedReservationSystem string `json:"computerized_reservation_system,required,nullable"`
	// Indicates the reason for a credit to the cardholder.
	CreditReasonIndicator ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicator `json:"credit_reason_indicator,required,nullable"`
	// Date of departure.
	DepartureDate time.Time `json:"departure_date,required,nullable" format:"date"`
	// Code for the originating city or airport.
	OriginationCityAirportCode string `json:"origination_city_airport_code,required,nullable"`
	// Name of the passenger.
	PassengerName string `json:"passenger_name,required,nullable"`
	// Indicates whether this ticket is non-refundable.
	RestrictedTicketIndicator ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelRestrictedTicketIndicator `json:"restricted_ticket_indicator,required,nullable"`
	// Indicates why a ticket was changed.
	TicketChangeIndicator ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelTicketChangeIndicator `json:"ticket_change_indicator,required,nullable"`
	// Ticket number.
	TicketNumber string `json:"ticket_number,required,nullable"`
	// Code for the travel agency if the ticket was issued by a travel agency.
	TravelAgencyCode string `json:"travel_agency_code,required,nullable"`
	// Name of the travel agency if the ticket was issued by a travel agency.
	TravelAgencyName string `json:"travel_agency_name,required,nullable"`
	// Fields specific to each leg of the journey.
	TripLegs []ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelTripLeg `json:"trip_legs,required,nullable"`
	JSON     achTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelJSON
}

// achTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelJSON
// contains the JSON metadata for the struct
// [ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravel]
type achTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelJSON struct {
	Ancillary                     apijson.Field
	ComputerizedReservationSystem apijson.Field
	CreditReasonIndicator         apijson.Field
	DepartureDate                 apijson.Field
	OriginationCityAirportCode    apijson.Field
	PassengerName                 apijson.Field
	RestrictedTicketIndicator     apijson.Field
	TicketChangeIndicator         apijson.Field
	TicketNumber                  apijson.Field
	TravelAgencyCode              apijson.Field
	TravelAgencyName              apijson.Field
	TripLegs                      apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Ancillary purchases in addition to the airfare.
type ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillary struct {
	// If this purchase has a connection or relationship to another purchase, such as a
	// baggage fee for a passenger transport ticket, this field should contain the
	// ticket document number for the other purchase.
	ConnectedTicketDocumentNumber string `json:"connected_ticket_document_number,required,nullable"`
	// Indicates the reason for a credit to the cardholder.
	CreditReasonIndicator ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicator `json:"credit_reason_indicator,required,nullable"`
	// Name of the passenger or description of the ancillary purchase.
	PassengerNameOrDescription string `json:"passenger_name_or_description,required,nullable"`
	// Additional travel charges, such as baggage fees.
	Services []ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryService `json:"services,required"`
	// Ticket document number.
	TicketDocumentNumber string `json:"ticket_document_number,required,nullable"`
	JSON                 achTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryJSON
}

// achTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryJSON
// contains the JSON metadata for the struct
// [ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillary]
type achTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryJSON struct {
	ConnectedTicketDocumentNumber apijson.Field
	CreditReasonIndicator         apijson.Field
	PassengerNameOrDescription    apijson.Field
	Services                      apijson.Field
	TicketDocumentNumber          apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the reason for a credit to the cardholder.
type ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicator string

const (
	// No credit
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicatorNoCredit ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicator = "no_credit"
	// Passenger transport ancillary purchase cancellation
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicatorPassengerTransportAncillaryPurchaseCancellation ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicator = "passenger_transport_ancillary_purchase_cancellation"
	// Airline ticket and passenger transport ancillary purchase cancellation
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicatorAirlineTicketAndPassengerTransportAncillaryPurchaseCancellation ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicator = "airline_ticket_and_passenger_transport_ancillary_purchase_cancellation"
	// Other
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicatorOther ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryCreditReasonIndicator = "other"
)

type ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryService struct {
	// Category of the ancillary service.
	Category ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory `json:"category,required,nullable"`
	// Sub-category of the ancillary service, free-form.
	SubCategory string `json:"sub_category,required,nullable"`
	JSON        achTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServiceJSON
}

// achTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServiceJSON
// contains the JSON metadata for the struct
// [ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryService]
type achTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServiceJSON struct {
	Category    apijson.Field
	SubCategory apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Category of the ancillary service.
type ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory string

const (
	// None
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryNone ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "none"
	// Bundled service
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryBundledService ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "bundled_service"
	// Baggage fee
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryBaggageFee ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "baggage_fee"
	// Change fee
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryChangeFee ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "change_fee"
	// Cargo
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryCargo ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "cargo"
	// Carbon offset
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryCarbonOffset ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "carbon_offset"
	// Frequent flyer
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryFrequentFlyer ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "frequent_flyer"
	// Gift card
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryGiftCard ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "gift_card"
	// Ground transport
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryGroundTransport ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "ground_transport"
	// In-flight entertainment
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryInFlightEntertainment ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "in_flight_entertainment"
	// Lounge
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryLounge ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "lounge"
	// Medical
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryMedical ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "medical"
	// Meal beverage
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryMealBeverage ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "meal_beverage"
	// Other
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryOther ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "other"
	// Passenger assist fee
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryPassengerAssistFee ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "passenger_assist_fee"
	// Pets
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryPets ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "pets"
	// Seat fees
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategorySeatFees ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "seat_fees"
	// Standby
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryStandby ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "standby"
	// Service fee
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryServiceFee ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "service_fee"
	// Store
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryStore ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "store"
	// Travel service
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryTravelService ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "travel_service"
	// Unaccompanied travel
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryUnaccompaniedTravel ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "unaccompanied_travel"
	// Upgrades
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryUpgrades ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "upgrades"
	// Wi-fi
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategoryWifi ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelAncillaryServicesCategory = "wifi"
)

// Indicates the reason for a credit to the cardholder.
type ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicator string

const (
	// No credit
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicatorNoCredit ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicator = "no_credit"
	// Passenger transport ancillary purchase cancellation
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicatorPassengerTransportAncillaryPurchaseCancellation ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicator = "passenger_transport_ancillary_purchase_cancellation"
	// Airline ticket and passenger transport ancillary purchase cancellation
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicatorAirlineTicketAndPassengerTransportAncillaryPurchaseCancellation ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicator = "airline_ticket_and_passenger_transport_ancillary_purchase_cancellation"
	// Airline ticket cancellation
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicatorAirlineTicketCancellation ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicator = "airline_ticket_cancellation"
	// Other
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicatorOther ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicator = "other"
	// Partial refund of airline ticket
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicatorPartialRefundOfAirlineTicket ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelCreditReasonIndicator = "partial_refund_of_airline_ticket"
)

// Indicates whether this ticket is non-refundable.
type ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelRestrictedTicketIndicator string

const (
	// No restrictions
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelRestrictedTicketIndicatorNoRestrictions ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelRestrictedTicketIndicator = "no_restrictions"
	// Restricted non-refundable ticket
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelRestrictedTicketIndicatorRestrictedNonRefundableTicket ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelRestrictedTicketIndicator = "restricted_non_refundable_ticket"
)

// Indicates why a ticket was changed.
type ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelTicketChangeIndicator string

const (
	// None
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelTicketChangeIndicatorNone ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelTicketChangeIndicator = "none"
	// Change to existing ticket
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelTicketChangeIndicatorChangeToExistingTicket ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelTicketChangeIndicator = "change_to_existing_ticket"
	// New ticket
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelTicketChangeIndicatorNewTicket ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelTicketChangeIndicator = "new_ticket"
)

type ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelTripLeg struct {
	// Carrier code (e.g., United Airlines, Jet Blue, etc.).
	CarrierCode string `json:"carrier_code,required,nullable"`
	// Code for the destination city or airport.
	DestinationCityAirportCode string `json:"destination_city_airport_code,required,nullable"`
	// Fare basis code.
	FareBasisCode string `json:"fare_basis_code,required,nullable"`
	// Flight number.
	FlightNumber string `json:"flight_number,required,nullable"`
	// Service class (e.g., first class, business class, etc.).
	ServiceClass string `json:"service_class,required,nullable"`
	// Indicates whether a stopover is allowed on this ticket.
	StopOverCode ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelTripLegsStopOverCode `json:"stop_over_code,required,nullable"`
	JSON         achTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelTripLegJSON
}

// achTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelTripLegJSON
// contains the JSON metadata for the struct
// [ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelTripLeg]
type achTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelTripLegJSON struct {
	CarrierCode                apijson.Field
	DestinationCityAirportCode apijson.Field
	FareBasisCode              apijson.Field
	FlightNumber               apijson.Field
	ServiceClass               apijson.Field
	StopOverCode               apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelTripLeg) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates whether a stopover is allowed on this ticket.
type ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelTripLegsStopOverCode string

const (
	// None
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelTripLegsStopOverCodeNone ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelTripLegsStopOverCode = "none"
	// Stop over allowed
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelTripLegsStopOverCodeStopOverAllowed ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelTripLegsStopOverCode = "stop_over_allowed"
	// Stop over not allowed
	ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelTripLegsStopOverCodeStopOverNotAllowed ACHTransferSimulationTransactionSourceCardSettlementPurchaseDetailsTravelTripLegsStopOverCode = "stop_over_not_allowed"
)

// A constant representing the object's type. For this resource it will always be
// `card_settlement`.
type ACHTransferSimulationTransactionSourceCardSettlementType string

const (
	ACHTransferSimulationTransactionSourceCardSettlementTypeCardSettlement ACHTransferSimulationTransactionSourceCardSettlementType = "card_settlement"
)

// The type of the resource. We may add additional possible values for this enum
// over time; your application should be able to handle such additions gracefully.
type ACHTransferSimulationTransactionSourceCategory string

const (
	// Account Transfer Intention: details will be under the
	// `account_transfer_intention` object.
	ACHTransferSimulationTransactionSourceCategoryAccountTransferIntention ACHTransferSimulationTransactionSourceCategory = "account_transfer_intention"
	// ACH Transfer Intention: details will be under the `ach_transfer_intention`
	// object.
	ACHTransferSimulationTransactionSourceCategoryACHTransferIntention ACHTransferSimulationTransactionSourceCategory = "ach_transfer_intention"
	// ACH Transfer Rejection: details will be under the `ach_transfer_rejection`
	// object.
	ACHTransferSimulationTransactionSourceCategoryACHTransferRejection ACHTransferSimulationTransactionSourceCategory = "ach_transfer_rejection"
	// ACH Transfer Return: details will be under the `ach_transfer_return` object.
	ACHTransferSimulationTransactionSourceCategoryACHTransferReturn ACHTransferSimulationTransactionSourceCategory = "ach_transfer_return"
	// Card Dispute Acceptance: details will be under the `card_dispute_acceptance`
	// object.
	ACHTransferSimulationTransactionSourceCategoryCardDisputeAcceptance ACHTransferSimulationTransactionSourceCategory = "card_dispute_acceptance"
	// Card Refund: details will be under the `card_refund` object.
	ACHTransferSimulationTransactionSourceCategoryCardRefund ACHTransferSimulationTransactionSourceCategory = "card_refund"
	// Card Revenue Payment: details will be under the `card_revenue_payment` object.
	ACHTransferSimulationTransactionSourceCategoryCardRevenuePayment ACHTransferSimulationTransactionSourceCategory = "card_revenue_payment"
	// Card Settlement: details will be under the `card_settlement` object.
	ACHTransferSimulationTransactionSourceCategoryCardSettlement ACHTransferSimulationTransactionSourceCategory = "card_settlement"
	// Check Deposit Acceptance: details will be under the `check_deposit_acceptance`
	// object.
	ACHTransferSimulationTransactionSourceCategoryCheckDepositAcceptance ACHTransferSimulationTransactionSourceCategory = "check_deposit_acceptance"
	// Check Deposit Return: details will be under the `check_deposit_return` object.
	ACHTransferSimulationTransactionSourceCategoryCheckDepositReturn ACHTransferSimulationTransactionSourceCategory = "check_deposit_return"
	// Check Transfer Deposit: details will be under the `check_transfer_deposit`
	// object.
	ACHTransferSimulationTransactionSourceCategoryCheckTransferDeposit ACHTransferSimulationTransactionSourceCategory = "check_transfer_deposit"
	// Check Transfer Intention: details will be under the `check_transfer_intention`
	// object.
	ACHTransferSimulationTransactionSourceCategoryCheckTransferIntention ACHTransferSimulationTransactionSourceCategory = "check_transfer_intention"
	// Check Transfer Stop Payment Request: details will be under the
	// `check_transfer_stop_payment_request` object.
	ACHTransferSimulationTransactionSourceCategoryCheckTransferStopPaymentRequest ACHTransferSimulationTransactionSourceCategory = "check_transfer_stop_payment_request"
	// Fee Payment: details will be under the `fee_payment` object.
	ACHTransferSimulationTransactionSourceCategoryFeePayment ACHTransferSimulationTransactionSourceCategory = "fee_payment"
	// Inbound ACH Transfer Intention: details will be under the `inbound_ach_transfer`
	// object.
	ACHTransferSimulationTransactionSourceCategoryInboundACHTransfer ACHTransferSimulationTransactionSourceCategory = "inbound_ach_transfer"
	// Inbound ACH Transfer Return Intention: details will be under the
	// `inbound_ach_transfer_return_intention` object.
	ACHTransferSimulationTransactionSourceCategoryInboundACHTransferReturnIntention ACHTransferSimulationTransactionSourceCategory = "inbound_ach_transfer_return_intention"
	// Inbound Check: details will be under the `inbound_check` object.
	ACHTransferSimulationTransactionSourceCategoryInboundCheck ACHTransferSimulationTransactionSourceCategory = "inbound_check"
	// Inbound International ACH Transfer: details will be under the
	// `inbound_international_ach_transfer` object.
	ACHTransferSimulationTransactionSourceCategoryInboundInternationalACHTransfer ACHTransferSimulationTransactionSourceCategory = "inbound_international_ach_transfer"
	// Inbound Real-Time Payments Transfer Confirmation: details will be under the
	// `inbound_real_time_payments_transfer_confirmation` object.
	ACHTransferSimulationTransactionSourceCategoryInboundRealTimePaymentsTransferConfirmation ACHTransferSimulationTransactionSourceCategory = "inbound_real_time_payments_transfer_confirmation"
	// Inbound Wire Drawdown Payment: details will be under the
	// `inbound_wire_drawdown_payment` object.
	ACHTransferSimulationTransactionSourceCategoryInboundWireDrawdownPayment ACHTransferSimulationTransactionSourceCategory = "inbound_wire_drawdown_payment"
	// Inbound Wire Drawdown Payment Reversal: details will be under the
	// `inbound_wire_drawdown_payment_reversal` object.
	ACHTransferSimulationTransactionSourceCategoryInboundWireDrawdownPaymentReversal ACHTransferSimulationTransactionSourceCategory = "inbound_wire_drawdown_payment_reversal"
	// Inbound Wire Reversal: details will be under the `inbound_wire_reversal` object.
	ACHTransferSimulationTransactionSourceCategoryInboundWireReversal ACHTransferSimulationTransactionSourceCategory = "inbound_wire_reversal"
	// Inbound Wire Transfer: details will be under the `inbound_wire_transfer` object.
	ACHTransferSimulationTransactionSourceCategoryInboundWireTransfer ACHTransferSimulationTransactionSourceCategory = "inbound_wire_transfer"
	// Interest Payment: details will be under the `interest_payment` object.
	ACHTransferSimulationTransactionSourceCategoryInterestPayment ACHTransferSimulationTransactionSourceCategory = "interest_payment"
	// Internal Source: details will be under the `internal_source` object.
	ACHTransferSimulationTransactionSourceCategoryInternalSource ACHTransferSimulationTransactionSourceCategory = "internal_source"
	// Real-Time Payments Transfer Acknowledgement: details will be under the
	// `real_time_payments_transfer_acknowledgement` object.
	ACHTransferSimulationTransactionSourceCategoryRealTimePaymentsTransferAcknowledgement ACHTransferSimulationTransactionSourceCategory = "real_time_payments_transfer_acknowledgement"
	// Sample Funds: details will be under the `sample_funds` object.
	ACHTransferSimulationTransactionSourceCategorySampleFunds ACHTransferSimulationTransactionSourceCategory = "sample_funds"
	// Wire Transfer Intention: details will be under the `wire_transfer_intention`
	// object.
	ACHTransferSimulationTransactionSourceCategoryWireTransferIntention ACHTransferSimulationTransactionSourceCategory = "wire_transfer_intention"
	// Wire Transfer Rejection: details will be under the `wire_transfer_rejection`
	// object.
	ACHTransferSimulationTransactionSourceCategoryWireTransferRejection ACHTransferSimulationTransactionSourceCategory = "wire_transfer_rejection"
	// The Transaction was made for an undocumented or deprecated reason.
	ACHTransferSimulationTransactionSourceCategoryOther ACHTransferSimulationTransactionSourceCategory = "other"
)

// A Check Deposit Acceptance object. This field will be present in the JSON
// response if and only if `category` is equal to `check_deposit_acceptance`.
type ACHTransferSimulationTransactionSourceCheckDepositAcceptance struct {
	// The account number printed on the check.
	AccountNumber string `json:"account_number,required"`
	// The amount to be deposited in the minor unit of the transaction's currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// An additional line of metadata printed on the check. This typically includes the
	// check number for business checks.
	AuxiliaryOnUs string `json:"auxiliary_on_us,required,nullable"`
	// The ID of the Check Deposit that was accepted.
	CheckDepositID string `json:"check_deposit_id,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency ACHTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency `json:"currency,required"`
	// The routing number printed on the check.
	RoutingNumber string `json:"routing_number,required"`
	// The check serial number, if present, for consumer checks. For business checks,
	// the serial number is usually in the `auxiliary_on_us` field.
	SerialNumber string `json:"serial_number,required,nullable"`
	JSON         achTransferSimulationTransactionSourceCheckDepositAcceptanceJSON
}

// achTransferSimulationTransactionSourceCheckDepositAcceptanceJSON contains the
// JSON metadata for the struct
// [ACHTransferSimulationTransactionSourceCheckDepositAcceptance]
type achTransferSimulationTransactionSourceCheckDepositAcceptanceJSON struct {
	AccountNumber  apijson.Field
	Amount         apijson.Field
	AuxiliaryOnUs  apijson.Field
	CheckDepositID apijson.Field
	Currency       apijson.Field
	RoutingNumber  apijson.Field
	SerialNumber   apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ACHTransferSimulationTransactionSourceCheckDepositAcceptance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
type ACHTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency string

const (
	// Canadian Dollar (CAD)
	ACHTransferSimulationTransactionSourceCheckDepositAcceptanceCurrencyCad ACHTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency = "CAD"
	// Swiss Franc (CHF)
	ACHTransferSimulationTransactionSourceCheckDepositAcceptanceCurrencyChf ACHTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency = "CHF"
	// Euro (EUR)
	ACHTransferSimulationTransactionSourceCheckDepositAcceptanceCurrencyEur ACHTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency = "EUR"
	// British Pound (GBP)
	ACHTransferSimulationTransactionSourceCheckDepositAcceptanceCurrencyGbp ACHTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency = "GBP"
	// Japanese Yen (JPY)
	ACHTransferSimulationTransactionSourceCheckDepositAcceptanceCurrencyJpy ACHTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency = "JPY"
	// US Dollar (USD)
	ACHTransferSimulationTransactionSourceCheckDepositAcceptanceCurrencyUsd ACHTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency = "USD"
)

// A Check Deposit Return object. This field will be present in the JSON response
// if and only if `category` is equal to `check_deposit_return`.
type ACHTransferSimulationTransactionSourceCheckDepositReturn struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The identifier of the Check Deposit that was returned.
	CheckDepositID string `json:"check_deposit_id,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency     ACHTransferSimulationTransactionSourceCheckDepositReturnCurrency     `json:"currency,required"`
	ReturnReason ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReason `json:"return_reason,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check deposit was returned.
	ReturnedAt time.Time `json:"returned_at,required" format:"date-time"`
	// The identifier of the transaction that reversed the original check deposit
	// transaction.
	TransactionID string `json:"transaction_id,required"`
	JSON          achTransferSimulationTransactionSourceCheckDepositReturnJSON
}

// achTransferSimulationTransactionSourceCheckDepositReturnJSON contains the JSON
// metadata for the struct
// [ACHTransferSimulationTransactionSourceCheckDepositReturn]
type achTransferSimulationTransactionSourceCheckDepositReturnJSON struct {
	Amount         apijson.Field
	CheckDepositID apijson.Field
	Currency       apijson.Field
	ReturnReason   apijson.Field
	ReturnedAt     apijson.Field
	TransactionID  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ACHTransferSimulationTransactionSourceCheckDepositReturn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
type ACHTransferSimulationTransactionSourceCheckDepositReturnCurrency string

const (
	// Canadian Dollar (CAD)
	ACHTransferSimulationTransactionSourceCheckDepositReturnCurrencyCad ACHTransferSimulationTransactionSourceCheckDepositReturnCurrency = "CAD"
	// Swiss Franc (CHF)
	ACHTransferSimulationTransactionSourceCheckDepositReturnCurrencyChf ACHTransferSimulationTransactionSourceCheckDepositReturnCurrency = "CHF"
	// Euro (EUR)
	ACHTransferSimulationTransactionSourceCheckDepositReturnCurrencyEur ACHTransferSimulationTransactionSourceCheckDepositReturnCurrency = "EUR"
	// British Pound (GBP)
	ACHTransferSimulationTransactionSourceCheckDepositReturnCurrencyGbp ACHTransferSimulationTransactionSourceCheckDepositReturnCurrency = "GBP"
	// Japanese Yen (JPY)
	ACHTransferSimulationTransactionSourceCheckDepositReturnCurrencyJpy ACHTransferSimulationTransactionSourceCheckDepositReturnCurrency = "JPY"
	// US Dollar (USD)
	ACHTransferSimulationTransactionSourceCheckDepositReturnCurrencyUsd ACHTransferSimulationTransactionSourceCheckDepositReturnCurrency = "USD"
)

type ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReason string

const (
	// The check doesn't allow ACH conversion.
	ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReasonACHConversionNotSupported ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "ach_conversion_not_supported"
	// The account is closed.
	ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReasonClosedAccount ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "closed_account"
	// The check has already been deposited.
	ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReasonDuplicateSubmission ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "duplicate_submission"
	// Insufficient funds
	ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReasonInsufficientFunds ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "insufficient_funds"
	// No account was found matching the check details.
	ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReasonNoAccount ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "no_account"
	// The check was not authorized.
	ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReasonNotAuthorized ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "not_authorized"
	// The check is too old.
	ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReasonStaleDated ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "stale_dated"
	// The payment has been stopped by the account holder.
	ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReasonStopPayment ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "stop_payment"
	// The reason for the return is unknown.
	ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReasonUnknownReason ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "unknown_reason"
	// The image doesn't match the details submitted.
	ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReasonUnmatchedDetails ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "unmatched_details"
	// The image could not be read.
	ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReasonUnreadableImage ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "unreadable_image"
	// The check endorsement was irregular.
	ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReasonEndorsementIrregular ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "endorsement_irregular"
)

// A Check Transfer Deposit object. This field will be present in the JSON response
// if and only if `category` is equal to `check_transfer_deposit`.
type ACHTransferSimulationTransactionSourceCheckTransferDeposit struct {
	// The identifier of the API File object containing an image of the back of the
	// deposited check.
	BackImageFileID string `json:"back_image_file_id,required,nullable"`
	// When the check was deposited.
	DepositedAt time.Time `json:"deposited_at,required" format:"date-time"`
	// The identifier of the API File object containing an image of the front of the
	// deposited check.
	FrontImageFileID string `json:"front_image_file_id,required,nullable"`
	// The identifier of the Transaction object created when the check was deposited.
	TransactionID string `json:"transaction_id,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `check_transfer_deposit`.
	Type ACHTransferSimulationTransactionSourceCheckTransferDepositType `json:"type,required"`
	JSON achTransferSimulationTransactionSourceCheckTransferDepositJSON
}

// achTransferSimulationTransactionSourceCheckTransferDepositJSON contains the JSON
// metadata for the struct
// [ACHTransferSimulationTransactionSourceCheckTransferDeposit]
type achTransferSimulationTransactionSourceCheckTransferDepositJSON struct {
	BackImageFileID  apijson.Field
	DepositedAt      apijson.Field
	FrontImageFileID apijson.Field
	TransactionID    apijson.Field
	Type             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ACHTransferSimulationTransactionSourceCheckTransferDeposit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A constant representing the object's type. For this resource it will always be
// `check_transfer_deposit`.
type ACHTransferSimulationTransactionSourceCheckTransferDepositType string

const (
	ACHTransferSimulationTransactionSourceCheckTransferDepositTypeCheckTransferDeposit ACHTransferSimulationTransactionSourceCheckTransferDepositType = "check_transfer_deposit"
)

// A Check Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `check_transfer_intention`.
type ACHTransferSimulationTransactionSourceCheckTransferIntention struct {
	// The city of the check's destination.
	AddressCity string `json:"address_city,required,nullable"`
	// The street address of the check's destination.
	AddressLine1 string `json:"address_line1,required,nullable"`
	// The second line of the address of the check's destination.
	AddressLine2 string `json:"address_line2,required,nullable"`
	// The state of the check's destination.
	AddressState string `json:"address_state,required,nullable"`
	// The postal code of the check's destination.
	AddressZip string `json:"address_zip,required,nullable"`
	// The transfer amount in USD cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
	// currency.
	Currency ACHTransferSimulationTransactionSourceCheckTransferIntentionCurrency `json:"currency,required"`
	// The name that will be printed on the check.
	RecipientName string `json:"recipient_name,required,nullable"`
	// The identifier of the Check Transfer with which this is associated.
	TransferID string `json:"transfer_id,required"`
	JSON       achTransferSimulationTransactionSourceCheckTransferIntentionJSON
}

// achTransferSimulationTransactionSourceCheckTransferIntentionJSON contains the
// JSON metadata for the struct
// [ACHTransferSimulationTransactionSourceCheckTransferIntention]
type achTransferSimulationTransactionSourceCheckTransferIntentionJSON struct {
	AddressCity   apijson.Field
	AddressLine1  apijson.Field
	AddressLine2  apijson.Field
	AddressState  apijson.Field
	AddressZip    apijson.Field
	Amount        apijson.Field
	Currency      apijson.Field
	RecipientName apijson.Field
	TransferID    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ACHTransferSimulationTransactionSourceCheckTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
// currency.
type ACHTransferSimulationTransactionSourceCheckTransferIntentionCurrency string

const (
	// Canadian Dollar (CAD)
	ACHTransferSimulationTransactionSourceCheckTransferIntentionCurrencyCad ACHTransferSimulationTransactionSourceCheckTransferIntentionCurrency = "CAD"
	// Swiss Franc (CHF)
	ACHTransferSimulationTransactionSourceCheckTransferIntentionCurrencyChf ACHTransferSimulationTransactionSourceCheckTransferIntentionCurrency = "CHF"
	// Euro (EUR)
	ACHTransferSimulationTransactionSourceCheckTransferIntentionCurrencyEur ACHTransferSimulationTransactionSourceCheckTransferIntentionCurrency = "EUR"
	// British Pound (GBP)
	ACHTransferSimulationTransactionSourceCheckTransferIntentionCurrencyGbp ACHTransferSimulationTransactionSourceCheckTransferIntentionCurrency = "GBP"
	// Japanese Yen (JPY)
	ACHTransferSimulationTransactionSourceCheckTransferIntentionCurrencyJpy ACHTransferSimulationTransactionSourceCheckTransferIntentionCurrency = "JPY"
	// US Dollar (USD)
	ACHTransferSimulationTransactionSourceCheckTransferIntentionCurrencyUsd ACHTransferSimulationTransactionSourceCheckTransferIntentionCurrency = "USD"
)

// A Check Transfer Stop Payment Request object. This field will be present in the
// JSON response if and only if `category` is equal to
// `check_transfer_stop_payment_request`.
type ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequest struct {
	// The reason why this transfer was stopped.
	Reason ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequestReason `json:"reason,required"`
	// The time the stop-payment was requested.
	RequestedAt time.Time `json:"requested_at,required" format:"date-time"`
	// The ID of the check transfer that was stopped.
	TransferID string `json:"transfer_id,required"`
	// A constant representing the object's type. For this resource it will always be
	// `check_transfer_stop_payment_request`.
	Type ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequestType `json:"type,required"`
	JSON achTransferSimulationTransactionSourceCheckTransferStopPaymentRequestJSON
}

// achTransferSimulationTransactionSourceCheckTransferStopPaymentRequestJSON
// contains the JSON metadata for the struct
// [ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequest]
type achTransferSimulationTransactionSourceCheckTransferStopPaymentRequestJSON struct {
	Reason      apijson.Field
	RequestedAt apijson.Field
	TransferID  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The reason why this transfer was stopped.
type ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequestReason string

const (
	// The check could not be delivered.
	ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequestReasonMailDeliveryFailed ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequestReason = "mail_delivery_failed"
	// The check was canceled by an Increase operator who will provide details
	// out-of-band.
	ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequestReasonRejectedByIncrease ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequestReason = "rejected_by_increase"
	// The check was stopped for another reason.
	ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequestReasonUnknown ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequestReason = "unknown"
)

// A constant representing the object's type. For this resource it will always be
// `check_transfer_stop_payment_request`.
type ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequestType string

const (
	ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequestTypeCheckTransferStopPaymentRequest ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequestType = "check_transfer_stop_payment_request"
)

// A Fee Payment object. This field will be present in the JSON response if and
// only if `category` is equal to `fee_payment`.
type ACHTransferSimulationTransactionSourceFeePayment struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency ACHTransferSimulationTransactionSourceFeePaymentCurrency `json:"currency,required"`
	// The start of this payment's fee period, usually the first day of a month.
	FeePeriodStart time.Time `json:"fee_period_start,required" format:"date"`
	JSON           achTransferSimulationTransactionSourceFeePaymentJSON
}

// achTransferSimulationTransactionSourceFeePaymentJSON contains the JSON metadata
// for the struct [ACHTransferSimulationTransactionSourceFeePayment]
type achTransferSimulationTransactionSourceFeePaymentJSON struct {
	Amount         apijson.Field
	Currency       apijson.Field
	FeePeriodStart apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ACHTransferSimulationTransactionSourceFeePayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
// currency.
type ACHTransferSimulationTransactionSourceFeePaymentCurrency string

const (
	// Canadian Dollar (CAD)
	ACHTransferSimulationTransactionSourceFeePaymentCurrencyCad ACHTransferSimulationTransactionSourceFeePaymentCurrency = "CAD"
	// Swiss Franc (CHF)
	ACHTransferSimulationTransactionSourceFeePaymentCurrencyChf ACHTransferSimulationTransactionSourceFeePaymentCurrency = "CHF"
	// Euro (EUR)
	ACHTransferSimulationTransactionSourceFeePaymentCurrencyEur ACHTransferSimulationTransactionSourceFeePaymentCurrency = "EUR"
	// British Pound (GBP)
	ACHTransferSimulationTransactionSourceFeePaymentCurrencyGbp ACHTransferSimulationTransactionSourceFeePaymentCurrency = "GBP"
	// Japanese Yen (JPY)
	ACHTransferSimulationTransactionSourceFeePaymentCurrencyJpy ACHTransferSimulationTransactionSourceFeePaymentCurrency = "JPY"
	// US Dollar (USD)
	ACHTransferSimulationTransactionSourceFeePaymentCurrencyUsd ACHTransferSimulationTransactionSourceFeePaymentCurrency = "USD"
)

// An Inbound ACH Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `inbound_ach_transfer`.
type ACHTransferSimulationTransactionSourceInboundACHTransfer struct {
	// The amount in the minor unit of the destination account currency. For dollars,
	// for example, this is cents.
	Amount                             int64  `json:"amount,required"`
	OriginatorCompanyDescriptiveDate   string `json:"originator_company_descriptive_date,required,nullable"`
	OriginatorCompanyDiscretionaryData string `json:"originator_company_discretionary_data,required,nullable"`
	OriginatorCompanyEntryDescription  string `json:"originator_company_entry_description,required"`
	OriginatorCompanyID                string `json:"originator_company_id,required"`
	OriginatorCompanyName              string `json:"originator_company_name,required"`
	ReceiverIDNumber                   string `json:"receiver_id_number,required,nullable"`
	ReceiverName                       string `json:"receiver_name,required,nullable"`
	TraceNumber                        string `json:"trace_number,required"`
	JSON                               achTransferSimulationTransactionSourceInboundACHTransferJSON
}

// achTransferSimulationTransactionSourceInboundACHTransferJSON contains the JSON
// metadata for the struct
// [ACHTransferSimulationTransactionSourceInboundACHTransfer]
type achTransferSimulationTransactionSourceInboundACHTransferJSON struct {
	Amount                             apijson.Field
	OriginatorCompanyDescriptiveDate   apijson.Field
	OriginatorCompanyDiscretionaryData apijson.Field
	OriginatorCompanyEntryDescription  apijson.Field
	OriginatorCompanyID                apijson.Field
	OriginatorCompanyName              apijson.Field
	ReceiverIDNumber                   apijson.Field
	ReceiverName                       apijson.Field
	TraceNumber                        apijson.Field
	raw                                string
	ExtraFields                        map[string]apijson.Field
}

func (r *ACHTransferSimulationTransactionSourceInboundACHTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// An Inbound Check object. This field will be present in the JSON response if and
// only if `category` is equal to `inbound_check`.
type ACHTransferSimulationTransactionSourceInboundCheck struct {
	// The amount in the minor unit of the destination account currency. For dollars,
	// for example, this is cents.
	Amount                int64  `json:"amount,required"`
	CheckFrontImageFileID string `json:"check_front_image_file_id,required,nullable"`
	CheckNumber           string `json:"check_number,required,nullable"`
	CheckRearImageFileID  string `json:"check_rear_image_file_id,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency ACHTransferSimulationTransactionSourceInboundCheckCurrency `json:"currency,required"`
	JSON     achTransferSimulationTransactionSourceInboundCheckJSON
}

// achTransferSimulationTransactionSourceInboundCheckJSON contains the JSON
// metadata for the struct [ACHTransferSimulationTransactionSourceInboundCheck]
type achTransferSimulationTransactionSourceInboundCheckJSON struct {
	Amount                apijson.Field
	CheckFrontImageFileID apijson.Field
	CheckNumber           apijson.Field
	CheckRearImageFileID  apijson.Field
	Currency              apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ACHTransferSimulationTransactionSourceInboundCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
type ACHTransferSimulationTransactionSourceInboundCheckCurrency string

const (
	// Canadian Dollar (CAD)
	ACHTransferSimulationTransactionSourceInboundCheckCurrencyCad ACHTransferSimulationTransactionSourceInboundCheckCurrency = "CAD"
	// Swiss Franc (CHF)
	ACHTransferSimulationTransactionSourceInboundCheckCurrencyChf ACHTransferSimulationTransactionSourceInboundCheckCurrency = "CHF"
	// Euro (EUR)
	ACHTransferSimulationTransactionSourceInboundCheckCurrencyEur ACHTransferSimulationTransactionSourceInboundCheckCurrency = "EUR"
	// British Pound (GBP)
	ACHTransferSimulationTransactionSourceInboundCheckCurrencyGbp ACHTransferSimulationTransactionSourceInboundCheckCurrency = "GBP"
	// Japanese Yen (JPY)
	ACHTransferSimulationTransactionSourceInboundCheckCurrencyJpy ACHTransferSimulationTransactionSourceInboundCheckCurrency = "JPY"
	// US Dollar (USD)
	ACHTransferSimulationTransactionSourceInboundCheckCurrencyUsd ACHTransferSimulationTransactionSourceInboundCheckCurrency = "USD"
)

// An Inbound International ACH Transfer object. This field will be present in the
// JSON response if and only if `category` is equal to
// `inbound_international_ach_transfer`.
type ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer struct {
	// The amount in the minor unit of the destination account currency. For dollars,
	// for example, this is cents.
	Amount                                                 int64  `json:"amount,required"`
	DestinationCountryCode                                 string `json:"destination_country_code,required"`
	DestinationCurrencyCode                                string `json:"destination_currency_code,required"`
	ForeignExchangeIndicator                               string `json:"foreign_exchange_indicator,required"`
	ForeignExchangeReference                               string `json:"foreign_exchange_reference,required,nullable"`
	ForeignExchangeReferenceIndicator                      string `json:"foreign_exchange_reference_indicator,required"`
	ForeignPaymentAmount                                   int64  `json:"foreign_payment_amount,required"`
	ForeignTraceNumber                                     string `json:"foreign_trace_number,required,nullable"`
	InternationalTransactionTypeCode                       string `json:"international_transaction_type_code,required"`
	OriginatingCurrencyCode                                string `json:"originating_currency_code,required"`
	OriginatingDepositoryFinancialInstitutionBranchCountry string `json:"originating_depository_financial_institution_branch_country,required"`
	OriginatingDepositoryFinancialInstitutionID            string `json:"originating_depository_financial_institution_id,required"`
	OriginatingDepositoryFinancialInstitutionIDQualifier   string `json:"originating_depository_financial_institution_id_qualifier,required"`
	OriginatingDepositoryFinancialInstitutionName          string `json:"originating_depository_financial_institution_name,required"`
	OriginatorCity                                         string `json:"originator_city,required"`
	OriginatorCompanyEntryDescription                      string `json:"originator_company_entry_description,required"`
	OriginatorCountry                                      string `json:"originator_country,required"`
	OriginatorIdentification                               string `json:"originator_identification,required"`
	OriginatorName                                         string `json:"originator_name,required"`
	OriginatorPostalCode                                   string `json:"originator_postal_code,required,nullable"`
	OriginatorStateOrProvince                              string `json:"originator_state_or_province,required,nullable"`
	OriginatorStreetAddress                                string `json:"originator_street_address,required"`
	PaymentRelatedInformation                              string `json:"payment_related_information,required,nullable"`
	PaymentRelatedInformation2                             string `json:"payment_related_information2,required,nullable"`
	ReceiverCity                                           string `json:"receiver_city,required"`
	ReceiverCountry                                        string `json:"receiver_country,required"`
	ReceiverIdentificationNumber                           string `json:"receiver_identification_number,required,nullable"`
	ReceiverPostalCode                                     string `json:"receiver_postal_code,required,nullable"`
	ReceiverStateOrProvince                                string `json:"receiver_state_or_province,required,nullable"`
	ReceiverStreetAddress                                  string `json:"receiver_street_address,required"`
	ReceivingCompanyOrIndividualName                       string `json:"receiving_company_or_individual_name,required"`
	ReceivingDepositoryFinancialInstitutionCountry         string `json:"receiving_depository_financial_institution_country,required"`
	ReceivingDepositoryFinancialInstitutionID              string `json:"receiving_depository_financial_institution_id,required"`
	ReceivingDepositoryFinancialInstitutionIDQualifier     string `json:"receiving_depository_financial_institution_id_qualifier,required"`
	ReceivingDepositoryFinancialInstitutionName            string `json:"receiving_depository_financial_institution_name,required"`
	TraceNumber                                            string `json:"trace_number,required"`
	JSON                                                   achTransferSimulationTransactionSourceInboundInternationalACHTransferJSON
}

// achTransferSimulationTransactionSourceInboundInternationalACHTransferJSON
// contains the JSON metadata for the struct
// [ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer]
type achTransferSimulationTransactionSourceInboundInternationalACHTransferJSON struct {
	Amount                                                 apijson.Field
	DestinationCountryCode                                 apijson.Field
	DestinationCurrencyCode                                apijson.Field
	ForeignExchangeIndicator                               apijson.Field
	ForeignExchangeReference                               apijson.Field
	ForeignExchangeReferenceIndicator                      apijson.Field
	ForeignPaymentAmount                                   apijson.Field
	ForeignTraceNumber                                     apijson.Field
	InternationalTransactionTypeCode                       apijson.Field
	OriginatingCurrencyCode                                apijson.Field
	OriginatingDepositoryFinancialInstitutionBranchCountry apijson.Field
	OriginatingDepositoryFinancialInstitutionID            apijson.Field
	OriginatingDepositoryFinancialInstitutionIDQualifier   apijson.Field
	OriginatingDepositoryFinancialInstitutionName          apijson.Field
	OriginatorCity                                         apijson.Field
	OriginatorCompanyEntryDescription                      apijson.Field
	OriginatorCountry                                      apijson.Field
	OriginatorIdentification                               apijson.Field
	OriginatorName                                         apijson.Field
	OriginatorPostalCode                                   apijson.Field
	OriginatorStateOrProvince                              apijson.Field
	OriginatorStreetAddress                                apijson.Field
	PaymentRelatedInformation                              apijson.Field
	PaymentRelatedInformation2                             apijson.Field
	ReceiverCity                                           apijson.Field
	ReceiverCountry                                        apijson.Field
	ReceiverIdentificationNumber                           apijson.Field
	ReceiverPostalCode                                     apijson.Field
	ReceiverStateOrProvince                                apijson.Field
	ReceiverStreetAddress                                  apijson.Field
	ReceivingCompanyOrIndividualName                       apijson.Field
	ReceivingDepositoryFinancialInstitutionCountry         apijson.Field
	ReceivingDepositoryFinancialInstitutionID              apijson.Field
	ReceivingDepositoryFinancialInstitutionIDQualifier     apijson.Field
	ReceivingDepositoryFinancialInstitutionName            apijson.Field
	TraceNumber                                            apijson.Field
	raw                                                    string
	ExtraFields                                            map[string]apijson.Field
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// An Inbound Real-Time Payments Transfer Confirmation object. This field will be
// present in the JSON response if and only if `category` is equal to
// `inbound_real_time_payments_transfer_confirmation`.
type ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation struct {
	// The amount in the minor unit of the transfer's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The name the sender of the transfer specified as the recipient of the transfer.
	CreditorName string `json:"creditor_name,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the transfer's
	// currency. This will always be "USD" for a Real-Time Payments transfer.
	Currency ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency `json:"currency,required"`
	// The account number of the account that sent the transfer.
	DebtorAccountNumber string `json:"debtor_account_number,required"`
	// The name provided by the sender of the transfer.
	DebtorName string `json:"debtor_name,required"`
	// The routing number of the account that sent the transfer.
	DebtorRoutingNumber string `json:"debtor_routing_number,required"`
	// Additional information included with the transfer.
	RemittanceInformation string `json:"remittance_information,required,nullable"`
	// The Real-Time Payments network identification of the transfer
	TransactionIdentification string `json:"transaction_identification,required"`
	JSON                      achTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationJSON
}

// achTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationJSON
// contains the JSON metadata for the struct
// [ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation]
type achTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationJSON struct {
	Amount                    apijson.Field
	CreditorName              apijson.Field
	Currency                  apijson.Field
	DebtorAccountNumber       apijson.Field
	DebtorName                apijson.Field
	DebtorRoutingNumber       apijson.Field
	RemittanceInformation     apijson.Field
	TransactionIdentification apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the transfer's
// currency. This will always be "USD" for a Real-Time Payments transfer.
type ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency string

const (
	// Canadian Dollar (CAD)
	ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyCad ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "CAD"
	// Swiss Franc (CHF)
	ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyChf ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "CHF"
	// Euro (EUR)
	ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyEur ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "EUR"
	// British Pound (GBP)
	ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyGbp ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "GBP"
	// Japanese Yen (JPY)
	ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyJpy ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "JPY"
	// US Dollar (USD)
	ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyUsd ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "USD"
)

// An Inbound Wire Drawdown Payment object. This field will be present in the JSON
// response if and only if `category` is equal to `inbound_wire_drawdown_payment`.
type ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount                             int64  `json:"amount,required"`
	BeneficiaryAddressLine1            string `json:"beneficiary_address_line1,required,nullable"`
	BeneficiaryAddressLine2            string `json:"beneficiary_address_line2,required,nullable"`
	BeneficiaryAddressLine3            string `json:"beneficiary_address_line3,required,nullable"`
	BeneficiaryName                    string `json:"beneficiary_name,required,nullable"`
	BeneficiaryReference               string `json:"beneficiary_reference,required,nullable"`
	Description                        string `json:"description,required"`
	InputMessageAccountabilityData     string `json:"input_message_accountability_data,required,nullable"`
	OriginatorAddressLine1             string `json:"originator_address_line1,required,nullable"`
	OriginatorAddressLine2             string `json:"originator_address_line2,required,nullable"`
	OriginatorAddressLine3             string `json:"originator_address_line3,required,nullable"`
	OriginatorName                     string `json:"originator_name,required,nullable"`
	OriginatorToBeneficiaryInformation string `json:"originator_to_beneficiary_information,required,nullable"`
	JSON                               achTransferSimulationTransactionSourceInboundWireDrawdownPaymentJSON
}

// achTransferSimulationTransactionSourceInboundWireDrawdownPaymentJSON contains
// the JSON metadata for the struct
// [ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment]
type achTransferSimulationTransactionSourceInboundWireDrawdownPaymentJSON struct {
	Amount                             apijson.Field
	BeneficiaryAddressLine1            apijson.Field
	BeneficiaryAddressLine2            apijson.Field
	BeneficiaryAddressLine3            apijson.Field
	BeneficiaryName                    apijson.Field
	BeneficiaryReference               apijson.Field
	Description                        apijson.Field
	InputMessageAccountabilityData     apijson.Field
	OriginatorAddressLine1             apijson.Field
	OriginatorAddressLine2             apijson.Field
	OriginatorAddressLine3             apijson.Field
	OriginatorName                     apijson.Field
	OriginatorToBeneficiaryInformation apijson.Field
	raw                                string
	ExtraFields                        map[string]apijson.Field
}

func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// An Inbound Wire Drawdown Payment Reversal object. This field will be present in
// the JSON response if and only if `category` is equal to
// `inbound_wire_drawdown_payment_reversal`.
type ACHTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal struct {
	// The amount that was reversed.
	Amount int64 `json:"amount,required"`
	// The description on the reversal message from Fedwire.
	Description string `json:"description,required"`
	// The Fedwire cycle date for the wire reversal.
	InputCycleDate time.Time `json:"input_cycle_date,required" format:"date"`
	// The Fedwire transaction identifier.
	InputMessageAccountabilityData string `json:"input_message_accountability_data,required"`
	// The Fedwire sequence number.
	InputSequenceNumber string `json:"input_sequence_number,required"`
	// The Fedwire input source identifier.
	InputSource string `json:"input_source,required"`
	// The Fedwire cycle date for the wire transfer that was reversed.
	PreviousMessageInputCycleDate time.Time `json:"previous_message_input_cycle_date,required" format:"date"`
	// The Fedwire transaction identifier for the wire transfer that was reversed.
	PreviousMessageInputMessageAccountabilityData string `json:"previous_message_input_message_accountability_data,required"`
	// The Fedwire sequence number for the wire transfer that was reversed.
	PreviousMessageInputSequenceNumber string `json:"previous_message_input_sequence_number,required"`
	// The Fedwire input source identifier for the wire transfer that was reversed.
	PreviousMessageInputSource string `json:"previous_message_input_source,required"`
	JSON                       achTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversalJSON
}

// achTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversalJSON
// contains the JSON metadata for the struct
// [ACHTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal]
type achTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversalJSON struct {
	Amount                                        apijson.Field
	Description                                   apijson.Field
	InputCycleDate                                apijson.Field
	InputMessageAccountabilityData                apijson.Field
	InputSequenceNumber                           apijson.Field
	InputSource                                   apijson.Field
	PreviousMessageInputCycleDate                 apijson.Field
	PreviousMessageInputMessageAccountabilityData apijson.Field
	PreviousMessageInputSequenceNumber            apijson.Field
	PreviousMessageInputSource                    apijson.Field
	raw                                           string
	ExtraFields                                   map[string]apijson.Field
}

func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// An Inbound Wire Reversal object. This field will be present in the JSON response
// if and only if `category` is equal to `inbound_wire_reversal`.
type ACHTransferSimulationTransactionSourceInboundWireReversal struct {
	// The amount that was reversed.
	Amount int64 `json:"amount,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the reversal was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The description on the reversal message from Fedwire.
	Description string `json:"description,required"`
	// Additional financial institution information included in the wire reversal.
	FinancialInstitutionToFinancialInstitutionInformation string `json:"financial_institution_to_financial_institution_information,required,nullable"`
	// The Fedwire cycle date for the wire reversal.
	InputCycleDate time.Time `json:"input_cycle_date,required" format:"date"`
	// The Fedwire transaction identifier.
	InputMessageAccountabilityData string `json:"input_message_accountability_data,required"`
	// The Fedwire sequence number.
	InputSequenceNumber string `json:"input_sequence_number,required"`
	// The Fedwire input source identifier.
	InputSource string `json:"input_source,required"`
	// The Fedwire cycle date for the wire transfer that was reversed.
	PreviousMessageInputCycleDate time.Time `json:"previous_message_input_cycle_date,required" format:"date"`
	// The Fedwire transaction identifier for the wire transfer that was reversed.
	PreviousMessageInputMessageAccountabilityData string `json:"previous_message_input_message_accountability_data,required"`
	// The Fedwire sequence number for the wire transfer that was reversed.
	PreviousMessageInputSequenceNumber string `json:"previous_message_input_sequence_number,required"`
	// The Fedwire input source identifier for the wire transfer that was reversed.
	PreviousMessageInputSource string `json:"previous_message_input_source,required"`
	// Information included in the wire reversal for the receiving financial
	// institution.
	ReceiverFinancialInstitutionInformation string `json:"receiver_financial_institution_information,required,nullable"`
	// The ID for the Transaction associated with the transfer reversal.
	TransactionID string `json:"transaction_id,required"`
	// The ID for the Wire Transfer that is being reversed.
	WireTransferID string `json:"wire_transfer_id,required"`
	JSON           achTransferSimulationTransactionSourceInboundWireReversalJSON
}

// achTransferSimulationTransactionSourceInboundWireReversalJSON contains the JSON
// metadata for the struct
// [ACHTransferSimulationTransactionSourceInboundWireReversal]
type achTransferSimulationTransactionSourceInboundWireReversalJSON struct {
	Amount                                                apijson.Field
	CreatedAt                                             apijson.Field
	Description                                           apijson.Field
	FinancialInstitutionToFinancialInstitutionInformation apijson.Field
	InputCycleDate                                        apijson.Field
	InputMessageAccountabilityData                        apijson.Field
	InputSequenceNumber                                   apijson.Field
	InputSource                                           apijson.Field
	PreviousMessageInputCycleDate                         apijson.Field
	PreviousMessageInputMessageAccountabilityData         apijson.Field
	PreviousMessageInputSequenceNumber                    apijson.Field
	PreviousMessageInputSource                            apijson.Field
	ReceiverFinancialInstitutionInformation               apijson.Field
	TransactionID                                         apijson.Field
	WireTransferID                                        apijson.Field
	raw                                                   string
	ExtraFields                                           map[string]apijson.Field
}

func (r *ACHTransferSimulationTransactionSourceInboundWireReversal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// An Inbound Wire Transfer object. This field will be present in the JSON response
// if and only if `category` is equal to `inbound_wire_transfer`.
type ACHTransferSimulationTransactionSourceInboundWireTransfer struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount                                  int64  `json:"amount,required"`
	BeneficiaryAddressLine1                 string `json:"beneficiary_address_line1,required,nullable"`
	BeneficiaryAddressLine2                 string `json:"beneficiary_address_line2,required,nullable"`
	BeneficiaryAddressLine3                 string `json:"beneficiary_address_line3,required,nullable"`
	BeneficiaryName                         string `json:"beneficiary_name,required,nullable"`
	BeneficiaryReference                    string `json:"beneficiary_reference,required,nullable"`
	Description                             string `json:"description,required"`
	InputMessageAccountabilityData          string `json:"input_message_accountability_data,required,nullable"`
	OriginatorAddressLine1                  string `json:"originator_address_line1,required,nullable"`
	OriginatorAddressLine2                  string `json:"originator_address_line2,required,nullable"`
	OriginatorAddressLine3                  string `json:"originator_address_line3,required,nullable"`
	OriginatorName                          string `json:"originator_name,required,nullable"`
	OriginatorToBeneficiaryInformation      string `json:"originator_to_beneficiary_information,required,nullable"`
	OriginatorToBeneficiaryInformationLine1 string `json:"originator_to_beneficiary_information_line1,required,nullable"`
	OriginatorToBeneficiaryInformationLine2 string `json:"originator_to_beneficiary_information_line2,required,nullable"`
	OriginatorToBeneficiaryInformationLine3 string `json:"originator_to_beneficiary_information_line3,required,nullable"`
	OriginatorToBeneficiaryInformationLine4 string `json:"originator_to_beneficiary_information_line4,required,nullable"`
	JSON                                    achTransferSimulationTransactionSourceInboundWireTransferJSON
}

// achTransferSimulationTransactionSourceInboundWireTransferJSON contains the JSON
// metadata for the struct
// [ACHTransferSimulationTransactionSourceInboundWireTransfer]
type achTransferSimulationTransactionSourceInboundWireTransferJSON struct {
	Amount                                  apijson.Field
	BeneficiaryAddressLine1                 apijson.Field
	BeneficiaryAddressLine2                 apijson.Field
	BeneficiaryAddressLine3                 apijson.Field
	BeneficiaryName                         apijson.Field
	BeneficiaryReference                    apijson.Field
	Description                             apijson.Field
	InputMessageAccountabilityData          apijson.Field
	OriginatorAddressLine1                  apijson.Field
	OriginatorAddressLine2                  apijson.Field
	OriginatorAddressLine3                  apijson.Field
	OriginatorName                          apijson.Field
	OriginatorToBeneficiaryInformation      apijson.Field
	OriginatorToBeneficiaryInformationLine1 apijson.Field
	OriginatorToBeneficiaryInformationLine2 apijson.Field
	OriginatorToBeneficiaryInformationLine3 apijson.Field
	OriginatorToBeneficiaryInformationLine4 apijson.Field
	raw                                     string
	ExtraFields                             map[string]apijson.Field
}

func (r *ACHTransferSimulationTransactionSourceInboundWireTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// An Interest Payment object. This field will be present in the JSON response if
// and only if `category` is equal to `interest_payment`.
type ACHTransferSimulationTransactionSourceInterestPayment struct {
	// The account on which the interest was accrued.
	AccruedOnAccountID string `json:"accrued_on_account_id,required,nullable"`
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency ACHTransferSimulationTransactionSourceInterestPaymentCurrency `json:"currency,required"`
	// The end of the period for which this transaction paid interest.
	PeriodEnd time.Time `json:"period_end,required" format:"date-time"`
	// The start of the period for which this transaction paid interest.
	PeriodStart time.Time `json:"period_start,required" format:"date-time"`
	JSON        achTransferSimulationTransactionSourceInterestPaymentJSON
}

// achTransferSimulationTransactionSourceInterestPaymentJSON contains the JSON
// metadata for the struct [ACHTransferSimulationTransactionSourceInterestPayment]
type achTransferSimulationTransactionSourceInterestPaymentJSON struct {
	AccruedOnAccountID apijson.Field
	Amount             apijson.Field
	Currency           apijson.Field
	PeriodEnd          apijson.Field
	PeriodStart        apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ACHTransferSimulationTransactionSourceInterestPayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
// currency.
type ACHTransferSimulationTransactionSourceInterestPaymentCurrency string

const (
	// Canadian Dollar (CAD)
	ACHTransferSimulationTransactionSourceInterestPaymentCurrencyCad ACHTransferSimulationTransactionSourceInterestPaymentCurrency = "CAD"
	// Swiss Franc (CHF)
	ACHTransferSimulationTransactionSourceInterestPaymentCurrencyChf ACHTransferSimulationTransactionSourceInterestPaymentCurrency = "CHF"
	// Euro (EUR)
	ACHTransferSimulationTransactionSourceInterestPaymentCurrencyEur ACHTransferSimulationTransactionSourceInterestPaymentCurrency = "EUR"
	// British Pound (GBP)
	ACHTransferSimulationTransactionSourceInterestPaymentCurrencyGbp ACHTransferSimulationTransactionSourceInterestPaymentCurrency = "GBP"
	// Japanese Yen (JPY)
	ACHTransferSimulationTransactionSourceInterestPaymentCurrencyJpy ACHTransferSimulationTransactionSourceInterestPaymentCurrency = "JPY"
	// US Dollar (USD)
	ACHTransferSimulationTransactionSourceInterestPaymentCurrencyUsd ACHTransferSimulationTransactionSourceInterestPaymentCurrency = "USD"
)

// An Internal Source object. This field will be present in the JSON response if
// and only if `category` is equal to `internal_source`.
type ACHTransferSimulationTransactionSourceInternalSource struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency ACHTransferSimulationTransactionSourceInternalSourceCurrency `json:"currency,required"`
	Reason   ACHTransferSimulationTransactionSourceInternalSourceReason   `json:"reason,required"`
	JSON     achTransferSimulationTransactionSourceInternalSourceJSON
}

// achTransferSimulationTransactionSourceInternalSourceJSON contains the JSON
// metadata for the struct [ACHTransferSimulationTransactionSourceInternalSource]
type achTransferSimulationTransactionSourceInternalSourceJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACHTransferSimulationTransactionSourceInternalSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
// currency.
type ACHTransferSimulationTransactionSourceInternalSourceCurrency string

const (
	// Canadian Dollar (CAD)
	ACHTransferSimulationTransactionSourceInternalSourceCurrencyCad ACHTransferSimulationTransactionSourceInternalSourceCurrency = "CAD"
	// Swiss Franc (CHF)
	ACHTransferSimulationTransactionSourceInternalSourceCurrencyChf ACHTransferSimulationTransactionSourceInternalSourceCurrency = "CHF"
	// Euro (EUR)
	ACHTransferSimulationTransactionSourceInternalSourceCurrencyEur ACHTransferSimulationTransactionSourceInternalSourceCurrency = "EUR"
	// British Pound (GBP)
	ACHTransferSimulationTransactionSourceInternalSourceCurrencyGbp ACHTransferSimulationTransactionSourceInternalSourceCurrency = "GBP"
	// Japanese Yen (JPY)
	ACHTransferSimulationTransactionSourceInternalSourceCurrencyJpy ACHTransferSimulationTransactionSourceInternalSourceCurrency = "JPY"
	// US Dollar (USD)
	ACHTransferSimulationTransactionSourceInternalSourceCurrencyUsd ACHTransferSimulationTransactionSourceInternalSourceCurrency = "USD"
)

type ACHTransferSimulationTransactionSourceInternalSourceReason string

const (
	// Account closure
	ACHTransferSimulationTransactionSourceInternalSourceReasonAccountClosure ACHTransferSimulationTransactionSourceInternalSourceReason = "account_closure"
	// Bank migration
	ACHTransferSimulationTransactionSourceInternalSourceReasonBankMigration ACHTransferSimulationTransactionSourceInternalSourceReason = "bank_migration"
	// Cashback
	ACHTransferSimulationTransactionSourceInternalSourceReasonCashback ACHTransferSimulationTransactionSourceInternalSourceReason = "cashback"
	// Check adjustment
	ACHTransferSimulationTransactionSourceInternalSourceReasonCheckAdjustment ACHTransferSimulationTransactionSourceInternalSourceReason = "check_adjustment"
	// Collection receivable
	ACHTransferSimulationTransactionSourceInternalSourceReasonCollectionReceivable ACHTransferSimulationTransactionSourceInternalSourceReason = "collection_receivable"
	// Empyreal adjustment
	ACHTransferSimulationTransactionSourceInternalSourceReasonEmpyrealAdjustment ACHTransferSimulationTransactionSourceInternalSourceReason = "empyreal_adjustment"
	// Error
	ACHTransferSimulationTransactionSourceInternalSourceReasonError ACHTransferSimulationTransactionSourceInternalSourceReason = "error"
	// Error correction
	ACHTransferSimulationTransactionSourceInternalSourceReasonErrorCorrection ACHTransferSimulationTransactionSourceInternalSourceReason = "error_correction"
	// Fees
	ACHTransferSimulationTransactionSourceInternalSourceReasonFees ACHTransferSimulationTransactionSourceInternalSourceReason = "fees"
	// Interest
	ACHTransferSimulationTransactionSourceInternalSourceReasonInterest ACHTransferSimulationTransactionSourceInternalSourceReason = "interest"
	// Negative balance forgiveness
	ACHTransferSimulationTransactionSourceInternalSourceReasonNegativeBalanceForgiveness ACHTransferSimulationTransactionSourceInternalSourceReason = "negative_balance_forgiveness"
	// Sample funds
	ACHTransferSimulationTransactionSourceInternalSourceReasonSampleFunds ACHTransferSimulationTransactionSourceInternalSourceReason = "sample_funds"
	// Sample funds return
	ACHTransferSimulationTransactionSourceInternalSourceReasonSampleFundsReturn ACHTransferSimulationTransactionSourceInternalSourceReason = "sample_funds_return"
)

// A Real-Time Payments Transfer Acknowledgement object. This field will be present
// in the JSON response if and only if `category` is equal to
// `real_time_payments_transfer_acknowledgement`.
type ACHTransferSimulationTransactionSourceRealTimePaymentsTransferAcknowledgement struct {
	// The transfer amount in USD cents.
	Amount int64 `json:"amount,required"`
	// The destination account number.
	DestinationAccountNumber string `json:"destination_account_number,required"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	DestinationRoutingNumber string `json:"destination_routing_number,required"`
	// Unstructured information that will show on the recipient's bank statement.
	RemittanceInformation string `json:"remittance_information,required"`
	// The identifier of the Real-Time Payments Transfer that led to this Transaction.
	TransferID string `json:"transfer_id,required"`
	JSON       achTransferSimulationTransactionSourceRealTimePaymentsTransferAcknowledgementJSON
}

// achTransferSimulationTransactionSourceRealTimePaymentsTransferAcknowledgementJSON
// contains the JSON metadata for the struct
// [ACHTransferSimulationTransactionSourceRealTimePaymentsTransferAcknowledgement]
type achTransferSimulationTransactionSourceRealTimePaymentsTransferAcknowledgementJSON struct {
	Amount                   apijson.Field
	DestinationAccountNumber apijson.Field
	DestinationRoutingNumber apijson.Field
	RemittanceInformation    apijson.Field
	TransferID               apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *ACHTransferSimulationTransactionSourceRealTimePaymentsTransferAcknowledgement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Sample Funds object. This field will be present in the JSON response if and
// only if `category` is equal to `sample_funds`.
type ACHTransferSimulationTransactionSourceSampleFunds struct {
	// Where the sample funds came from.
	Originator string `json:"originator,required"`
	JSON       achTransferSimulationTransactionSourceSampleFundsJSON
}

// achTransferSimulationTransactionSourceSampleFundsJSON contains the JSON metadata
// for the struct [ACHTransferSimulationTransactionSourceSampleFunds]
type achTransferSimulationTransactionSourceSampleFundsJSON struct {
	Originator  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACHTransferSimulationTransactionSourceSampleFunds) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Wire Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `wire_transfer_intention`.
type ACHTransferSimulationTransactionSourceWireTransferIntention struct {
	// The destination account number.
	AccountNumber string `json:"account_number,required"`
	// The transfer amount in USD cents.
	Amount int64 `json:"amount,required"`
	// The message that will show on the recipient's bank statement.
	MessageToRecipient string `json:"message_to_recipient,required"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber string `json:"routing_number,required"`
	TransferID    string `json:"transfer_id,required"`
	JSON          achTransferSimulationTransactionSourceWireTransferIntentionJSON
}

// achTransferSimulationTransactionSourceWireTransferIntentionJSON contains the
// JSON metadata for the struct
// [ACHTransferSimulationTransactionSourceWireTransferIntention]
type achTransferSimulationTransactionSourceWireTransferIntentionJSON struct {
	AccountNumber      apijson.Field
	Amount             apijson.Field
	MessageToRecipient apijson.Field
	RoutingNumber      apijson.Field
	TransferID         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ACHTransferSimulationTransactionSourceWireTransferIntention) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Wire Transfer Rejection object. This field will be present in the JSON
// response if and only if `category` is equal to `wire_transfer_rejection`.
type ACHTransferSimulationTransactionSourceWireTransferRejection struct {
	TransferID string `json:"transfer_id,required"`
	JSON       achTransferSimulationTransactionSourceWireTransferRejectionJSON
}

// achTransferSimulationTransactionSourceWireTransferRejectionJSON contains the
// JSON metadata for the struct
// [ACHTransferSimulationTransactionSourceWireTransferRejection]
type achTransferSimulationTransactionSourceWireTransferRejectionJSON struct {
	TransferID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACHTransferSimulationTransactionSourceWireTransferRejection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A constant representing the object's type. For this resource it will always be
// `transaction`.
type ACHTransferSimulationTransactionType string

const (
	ACHTransferSimulationTransactionTypeTransaction ACHTransferSimulationTransactionType = "transaction"
)

// A constant representing the object's type. For this resource it will always be
// `inbound_ach_transfer_simulation_result`.
type ACHTransferSimulationType string

const (
	ACHTransferSimulationTypeInboundACHTransferSimulationResult ACHTransferSimulationType = "inbound_ach_transfer_simulation_result"
)

type SimulationACHTransferNewInboundParams struct {
	// The identifier of the Account Number the inbound ACH Transfer is for.
	AccountNumberID param.Field[string] `json:"account_number_id,required"`
	// The transfer amount in cents. A positive amount originates a credit transfer
	// pushing funds to the receiving account. A negative amount originates a debit
	// transfer pulling funds from the receiving account.
	Amount param.Field[int64] `json:"amount,required"`
	// The description of the date of the transfer.
	CompanyDescriptiveDate param.Field[string] `json:"company_descriptive_date"`
	// Data associated with the transfer set by the sender.
	CompanyDiscretionaryData param.Field[string] `json:"company_discretionary_data"`
	// The description of the transfer set by the sender.
	CompanyEntryDescription param.Field[string] `json:"company_entry_description"`
	// The sender's company id.
	CompanyID param.Field[string] `json:"company_id"`
	// The name of the sender.
	CompanyName param.Field[string] `json:"company_name"`
}

func (r SimulationACHTransferNewInboundParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SimulationACHTransferReturnParams struct {
	// The reason why the Federal Reserve or destination bank returned this transfer.
	// Defaults to `no_account`.
	Reason param.Field[SimulationACHTransferReturnParamsReason] `json:"reason"`
}

func (r SimulationACHTransferReturnParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The reason why the Federal Reserve or destination bank returned this transfer.
// Defaults to `no_account`.
type SimulationACHTransferReturnParamsReason string

const (
	// Code R01. Insufficient funds in the source account.
	SimulationACHTransferReturnParamsReasonInsufficientFund SimulationACHTransferReturnParamsReason = "insufficient_fund"
	// Code R03. The account does not exist or the receiving bank was unable to locate
	// it.
	SimulationACHTransferReturnParamsReasonNoAccount SimulationACHTransferReturnParamsReason = "no_account"
	// Code R02. The account is closed.
	SimulationACHTransferReturnParamsReasonAccountClosed SimulationACHTransferReturnParamsReason = "account_closed"
	// Code R04. The account number is invalid at the receiving bank.
	SimulationACHTransferReturnParamsReasonInvalidAccountNumberStructure SimulationACHTransferReturnParamsReason = "invalid_account_number_structure"
	// Code R16. The account was frozen per the Office of Foreign Assets Control.
	SimulationACHTransferReturnParamsReasonAccountFrozenEntryReturnedPerOfacInstruction SimulationACHTransferReturnParamsReason = "account_frozen_entry_returned_per_ofac_instruction"
	// Code R23. The receiving bank account refused a credit transfer.
	SimulationACHTransferReturnParamsReasonCreditEntryRefusedByReceiver SimulationACHTransferReturnParamsReason = "credit_entry_refused_by_receiver"
	// Code R05. The receiving bank rejected because of an incorrect Standard Entry
	// Class code.
	SimulationACHTransferReturnParamsReasonUnauthorizedDebitToConsumerAccountUsingCorporateSecCode SimulationACHTransferReturnParamsReason = "unauthorized_debit_to_consumer_account_using_corporate_sec_code"
	// Code R29. The corporate customer reversed the transfer.
	SimulationACHTransferReturnParamsReasonCorporateCustomerAdvisedNotAuthorized SimulationACHTransferReturnParamsReason = "corporate_customer_advised_not_authorized"
	// Code R08. The receiving bank stopped payment on this transfer.
	SimulationACHTransferReturnParamsReasonPaymentStopped SimulationACHTransferReturnParamsReason = "payment_stopped"
	// Code R20. The receiving bank account does not perform transfers.
	SimulationACHTransferReturnParamsReasonNonTransactionAccount SimulationACHTransferReturnParamsReason = "non_transaction_account"
	// Code R09. The receiving bank account does not have enough available balance for
	// the transfer.
	SimulationACHTransferReturnParamsReasonUncollectedFunds SimulationACHTransferReturnParamsReason = "uncollected_funds"
	// Code R28. The routing number is incorrect.
	SimulationACHTransferReturnParamsReasonRoutingNumberCheckDigitError SimulationACHTransferReturnParamsReason = "routing_number_check_digit_error"
	// Code R10. The customer reversed the transfer.
	SimulationACHTransferReturnParamsReasonCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete SimulationACHTransferReturnParamsReason = "customer_advised_unauthorized_improper_ineligible_or_incomplete"
	// Code R19. The amount field is incorrect or too large.
	SimulationACHTransferReturnParamsReasonAmountFieldError SimulationACHTransferReturnParamsReason = "amount_field_error"
	// Code R07. The customer who initiated the transfer revoked authorization.
	SimulationACHTransferReturnParamsReasonAuthorizationRevokedByCustomer SimulationACHTransferReturnParamsReason = "authorization_revoked_by_customer"
	// Code R13. The routing number is invalid.
	SimulationACHTransferReturnParamsReasonInvalidACHRoutingNumber SimulationACHTransferReturnParamsReason = "invalid_ach_routing_number"
	// Code R17. The receiving bank is unable to process a field in the transfer.
	SimulationACHTransferReturnParamsReasonFileRecordEditCriteria SimulationACHTransferReturnParamsReason = "file_record_edit_criteria"
	// Code R45. The individual name field was invalid.
	SimulationACHTransferReturnParamsReasonEnrInvalidIndividualName SimulationACHTransferReturnParamsReason = "enr_invalid_individual_name"
	// Code R06. The originating financial institution reversed the transfer.
	SimulationACHTransferReturnParamsReasonReturnedPerOdfiRequest SimulationACHTransferReturnParamsReason = "returned_per_odfi_request"
	// Code R34. The receiving bank's regulatory supervisor has limited their
	// participation.
	SimulationACHTransferReturnParamsReasonLimitedParticipationDfi SimulationACHTransferReturnParamsReason = "limited_participation_dfi"
	// Code R85. The outbound international ACH transfer was incorrect.
	SimulationACHTransferReturnParamsReasonIncorrectlyCodedOutboundInternationalPayment SimulationACHTransferReturnParamsReason = "incorrectly_coded_outbound_international_payment"
	// Code R12. A rare return reason. The account was sold to another bank.
	SimulationACHTransferReturnParamsReasonAccountSoldToAnotherDfi SimulationACHTransferReturnParamsReason = "account_sold_to_another_dfi"
	// Code R25. The addenda record is incorrect or missing.
	SimulationACHTransferReturnParamsReasonAddendaError SimulationACHTransferReturnParamsReason = "addenda_error"
	// Code R15. A rare return reason. The account holder is deceased.
	SimulationACHTransferReturnParamsReasonBeneficiaryOrAccountHolderDeceased SimulationACHTransferReturnParamsReason = "beneficiary_or_account_holder_deceased"
	// Code R11. A rare return reason. The customer authorized some payment to the
	// sender, but this payment was not in error.
	SimulationACHTransferReturnParamsReasonCustomerAdvisedNotWithinAuthorizationTerms SimulationACHTransferReturnParamsReason = "customer_advised_not_within_authorization_terms"
	// Code R74. A rare return reason. Sent in response to a return that was returned
	// with code `field_error`. The latest return should include the corrected
	// field(s).
	SimulationACHTransferReturnParamsReasonCorrectedReturn SimulationACHTransferReturnParamsReason = "corrected_return"
	// Code R24. A rare return reason. The receiving bank received an exact duplicate
	// entry with the same trace number and amount.
	SimulationACHTransferReturnParamsReasonDuplicateEntry SimulationACHTransferReturnParamsReason = "duplicate_entry"
	// Code R67. A rare return reason. The return this message refers to was a
	// duplicate.
	SimulationACHTransferReturnParamsReasonDuplicateReturn SimulationACHTransferReturnParamsReason = "duplicate_return"
	// Code R47. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	SimulationACHTransferReturnParamsReasonEnrDuplicateEnrollment SimulationACHTransferReturnParamsReason = "enr_duplicate_enrollment"
	// Code R43. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	SimulationACHTransferReturnParamsReasonEnrInvalidDfiAccountNumber SimulationACHTransferReturnParamsReason = "enr_invalid_dfi_account_number"
	// Code R44. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	SimulationACHTransferReturnParamsReasonEnrInvalidIndividualIDNumber SimulationACHTransferReturnParamsReason = "enr_invalid_individual_id_number"
	// Code R46. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	SimulationACHTransferReturnParamsReasonEnrInvalidRepresentativePayeeIndicator SimulationACHTransferReturnParamsReason = "enr_invalid_representative_payee_indicator"
	// Code R41. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	SimulationACHTransferReturnParamsReasonEnrInvalidTransactionCode SimulationACHTransferReturnParamsReason = "enr_invalid_transaction_code"
	// Code R40. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	SimulationACHTransferReturnParamsReasonEnrReturnOfEnrEntry SimulationACHTransferReturnParamsReason = "enr_return_of_enr_entry"
	// Code R42. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	SimulationACHTransferReturnParamsReasonEnrRoutingNumberCheckDigitError SimulationACHTransferReturnParamsReason = "enr_routing_number_check_digit_error"
	// Code R84. A rare return reason. The International ACH Transfer cannot be
	// processed by the gateway.
	SimulationACHTransferReturnParamsReasonEntryNotProcessedByGateway SimulationACHTransferReturnParamsReason = "entry_not_processed_by_gateway"
	// Code R69. A rare return reason. One or more of the fields in the ACH were
	// malformed.
	SimulationACHTransferReturnParamsReasonFieldError SimulationACHTransferReturnParamsReason = "field_error"
	// Code R83. A rare return reason. The Foreign receiving bank was unable to settle
	// this ACH transfer.
	SimulationACHTransferReturnParamsReasonForeignReceivingDfiUnableToSettle SimulationACHTransferReturnParamsReason = "foreign_receiving_dfi_unable_to_settle"
	// Code R80. A rare return reason. The International ACH Transfer is malformed.
	SimulationACHTransferReturnParamsReasonIatEntryCodingError SimulationACHTransferReturnParamsReason = "iat_entry_coding_error"
	// Code R18. A rare return reason. The ACH has an improper effective entry date
	// field.
	SimulationACHTransferReturnParamsReasonImproperEffectiveEntryDate SimulationACHTransferReturnParamsReason = "improper_effective_entry_date"
	// Code R39. A rare return reason. The source document related to this ACH, usually
	// an ACH check conversion, was presented to the bank.
	SimulationACHTransferReturnParamsReasonImproperSourceDocumentSourceDocumentPresented SimulationACHTransferReturnParamsReason = "improper_source_document_source_document_presented"
	// Code R21. A rare return reason. The Company ID field of the ACH was invalid.
	SimulationACHTransferReturnParamsReasonInvalidCompanyID SimulationACHTransferReturnParamsReason = "invalid_company_id"
	// Code R82. A rare return reason. The foreign receiving bank identifier for an
	// International ACH Transfer was invalid.
	SimulationACHTransferReturnParamsReasonInvalidForeignReceivingDfiIdentification SimulationACHTransferReturnParamsReason = "invalid_foreign_receiving_dfi_identification"
	// Code R22. A rare return reason. The Individual ID number field of the ACH was
	// invalid.
	SimulationACHTransferReturnParamsReasonInvalidIndividualIDNumber SimulationACHTransferReturnParamsReason = "invalid_individual_id_number"
	// Code R53. A rare return reason. Both the Represented Check ("RCK") entry and the
	// original check were presented to the bank.
	SimulationACHTransferReturnParamsReasonItemAndRckEntryPresentedForPayment SimulationACHTransferReturnParamsReason = "item_and_rck_entry_presented_for_payment"
	// Code R51. A rare return reason. The Represented Check ("RCK") entry is
	// ineligible.
	SimulationACHTransferReturnParamsReasonItemRelatedToRckEntryIsIneligible SimulationACHTransferReturnParamsReason = "item_related_to_rck_entry_is_ineligible"
	// Code R26. A rare return reason. The ACH is missing a required field.
	SimulationACHTransferReturnParamsReasonMandatoryFieldError SimulationACHTransferReturnParamsReason = "mandatory_field_error"
	// Code R71. A rare return reason. The receiving bank does not recognize the
	// routing number in a dishonored return entry.
	SimulationACHTransferReturnParamsReasonMisroutedDishonoredReturn SimulationACHTransferReturnParamsReason = "misrouted_dishonored_return"
	// Code R61. A rare return reason. The receiving bank does not recognize the
	// routing number in a return entry.
	SimulationACHTransferReturnParamsReasonMisroutedReturn SimulationACHTransferReturnParamsReason = "misrouted_return"
	// Code R76. A rare return reason. Sent in response to a return, the bank does not
	// find the errors alleged by the returning bank.
	SimulationACHTransferReturnParamsReasonNoErrorsFound SimulationACHTransferReturnParamsReason = "no_errors_found"
	// Code R77. A rare return reason. The receiving bank does not accept the return of
	// the erroneous debit. The funds are not available at the receiving bank.
	SimulationACHTransferReturnParamsReasonNonAcceptanceOfR62DishonoredReturn SimulationACHTransferReturnParamsReason = "non_acceptance_of_r62_dishonored_return"
	// Code R81. A rare return reason. The receiving bank does not accept International
	// ACH Transfers.
	SimulationACHTransferReturnParamsReasonNonParticipantInIatProgram SimulationACHTransferReturnParamsReason = "non_participant_in_iat_program"
	// Code R31. A rare return reason. A return that has been agreed to be accepted by
	// the receiving bank, despite falling outside of the usual return timeframe.
	SimulationACHTransferReturnParamsReasonPermissibleReturnEntry SimulationACHTransferReturnParamsReason = "permissible_return_entry"
	// Code R70. A rare return reason. The receiving bank had not approved this return.
	SimulationACHTransferReturnParamsReasonPermissibleReturnEntryNotAccepted SimulationACHTransferReturnParamsReason = "permissible_return_entry_not_accepted"
	// Code R32. A rare return reason. The receiving bank could not settle this
	// transaction.
	SimulationACHTransferReturnParamsReasonRdfiNonSettlement SimulationACHTransferReturnParamsReason = "rdfi_non_settlement"
	// Code R30. A rare return reason. The receiving bank does not accept Check
	// Truncation ACH transfers.
	SimulationACHTransferReturnParamsReasonRdfiParticipantInCheckTruncationProgram SimulationACHTransferReturnParamsReason = "rdfi_participant_in_check_truncation_program"
	// Code R14. A rare return reason. The payee is deceased.
	SimulationACHTransferReturnParamsReasonRepresentativePayeeDeceasedOrUnableToContinueInThatCapacity SimulationACHTransferReturnParamsReason = "representative_payee_deceased_or_unable_to_continue_in_that_capacity"
	// Code R75. A rare return reason. The originating bank disputes that an earlier
	// `duplicate_entry` return was actually a duplicate.
	SimulationACHTransferReturnParamsReasonReturnNotADuplicate SimulationACHTransferReturnParamsReason = "return_not_a_duplicate"
	// Code R62. A rare return reason. The originating bank made a mistake earlier and
	// this return corrects it.
	SimulationACHTransferReturnParamsReasonReturnOfErroneousOrReversingDebit SimulationACHTransferReturnParamsReason = "return_of_erroneous_or_reversing_debit"
	// Code R36. A rare return reason. Return of a malformed credit entry.
	SimulationACHTransferReturnParamsReasonReturnOfImproperCreditEntry SimulationACHTransferReturnParamsReason = "return_of_improper_credit_entry"
	// Code R35. A rare return reason. Return of a malformed debit entry.
	SimulationACHTransferReturnParamsReasonReturnOfImproperDebitEntry SimulationACHTransferReturnParamsReason = "return_of_improper_debit_entry"
	// Code R33. A rare return reason. Return of a Destroyed Check ("XKC") entry.
	SimulationACHTransferReturnParamsReasonReturnOfXckEntry SimulationACHTransferReturnParamsReason = "return_of_xck_entry"
	// Code R37. A rare return reason. The source document related to this ACH, usually
	// an ACH check conversion, was presented to the bank.
	SimulationACHTransferReturnParamsReasonSourceDocumentPresentedForPayment SimulationACHTransferReturnParamsReason = "source_document_presented_for_payment"
	// Code R50. A rare return reason. State law prevents the bank from accepting the
	// Represented Check ("RCK") entry.
	SimulationACHTransferReturnParamsReasonStateLawAffectingRckAcceptance SimulationACHTransferReturnParamsReason = "state_law_affecting_rck_acceptance"
	// Code R52. A rare return reason. A stop payment was issued on a Represented Check
	// ("RCK") entry.
	SimulationACHTransferReturnParamsReasonStopPaymentOnItemRelatedToRckEntry SimulationACHTransferReturnParamsReason = "stop_payment_on_item_related_to_rck_entry"
	// Code R38. A rare return reason. The source attached to the ACH, usually an ACH
	// check conversion, includes a stop payment.
	SimulationACHTransferReturnParamsReasonStopPaymentOnSourceDocument SimulationACHTransferReturnParamsReason = "stop_payment_on_source_document"
	// Code R73. A rare return reason. The bank receiving an `untimely_return` believes
	// it was on time.
	SimulationACHTransferReturnParamsReasonTimelyOriginalReturn SimulationACHTransferReturnParamsReason = "timely_original_return"
	// Code R27. A rare return reason. An ACH Return's trace number does not match an
	// originated ACH.
	SimulationACHTransferReturnParamsReasonTraceNumberError SimulationACHTransferReturnParamsReason = "trace_number_error"
	// Code R72. A rare return reason. The dishonored return was sent too late.
	SimulationACHTransferReturnParamsReasonUntimelyDishonoredReturn SimulationACHTransferReturnParamsReason = "untimely_dishonored_return"
	// Code R68. A rare return reason. The return was sent too late.
	SimulationACHTransferReturnParamsReasonUntimelyReturn SimulationACHTransferReturnParamsReason = "untimely_return"
)

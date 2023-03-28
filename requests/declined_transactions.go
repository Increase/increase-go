package requests

import (
	"fmt"
	"net/url"
	"time"

	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/fields"
)

type DeclinedTransaction struct {
	// The identifier for the Account the Declined Transaction belongs to.
	AccountID fields.Field[string] `json:"account_id,required"`
	// The Declined Transaction amount in the minor unit of its currency. For dollars,
	// for example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Declined
	// Transaction's currency. This will match the currency on the Declined
	// Transcation's Account.
	Currency fields.Field[DeclinedTransactionCurrency] `json:"currency,required"`
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
	RouteType fields.Field[DeclinedTransactionRouteType] `json:"route_type,required,nullable"`
	// This is an object giving more details on the network-level event that caused the
	// Declined Transaction. For example, for a card transaction this lists the
	// merchant's industry and location. Note that for backwards compatibility reasons,
	// additional undocumented keys may appear in this object. These should be treated
	// as deprecated and will be removed in the future.
	Source fields.Field[DeclinedTransactionSource] `json:"source,required"`
	// A constant representing the object's type. For this resource it will always be
	// `declined_transaction`.
	Type fields.Field[DeclinedTransactionType] `json:"type,required"`
}

// MarshalJSON serializes DeclinedTransaction into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *DeclinedTransaction) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r DeclinedTransaction) String() (result string) {
	return fmt.Sprintf("&DeclinedTransaction{AccountID:%s Amount:%s Currency:%s CreatedAt:%s Description:%s ID:%s RouteID:%s RouteType:%s Source:%s Type:%s}", r.AccountID, r.Amount, r.Currency, r.CreatedAt, r.Description, r.ID, r.RouteID, r.RouteType, r.Source, r.Type)
}

type DeclinedTransactionCurrency string

const (
	DeclinedTransactionCurrencyCad DeclinedTransactionCurrency = "CAD"
	DeclinedTransactionCurrencyChf DeclinedTransactionCurrency = "CHF"
	DeclinedTransactionCurrencyEur DeclinedTransactionCurrency = "EUR"
	DeclinedTransactionCurrencyGbp DeclinedTransactionCurrency = "GBP"
	DeclinedTransactionCurrencyJpy DeclinedTransactionCurrency = "JPY"
	DeclinedTransactionCurrencyUsd DeclinedTransactionCurrency = "USD"
)

type DeclinedTransactionRouteType string

const (
	DeclinedTransactionRouteTypeAccountNumber DeclinedTransactionRouteType = "account_number"
	DeclinedTransactionRouteTypeCard          DeclinedTransactionRouteType = "card"
)

type DeclinedTransactionSource struct {
	// The type of decline that took place. We may add additional possible values for
	// this enum over time; your application should be able to handle such additions
	// gracefully.
	Category fields.Field[DeclinedTransactionSourceCategory] `json:"category,required"`
	// A ACH Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `ach_decline`.
	ACHDecline fields.Field[DeclinedTransactionSourceACHDecline] `json:"ach_decline,required,nullable"`
	// A Card Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_decline`.
	CardDecline fields.Field[DeclinedTransactionSourceCardDecline] `json:"card_decline,required,nullable"`
	// A Check Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `check_decline`.
	CheckDecline fields.Field[DeclinedTransactionSourceCheckDecline] `json:"check_decline,required,nullable"`
	// A Inbound Real Time Payments Transfer Decline object. This field will be present
	// in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_decline`.
	InboundRealTimePaymentsTransferDecline fields.Field[DeclinedTransactionSourceInboundRealTimePaymentsTransferDecline] `json:"inbound_real_time_payments_transfer_decline,required,nullable"`
	// A International ACH Decline object. This field will be present in the JSON
	// response if and only if `category` is equal to `international_ach_decline`.
	InternationalACHDecline fields.Field[DeclinedTransactionSourceInternationalACHDecline] `json:"international_ach_decline,required,nullable"`
	// A Deprecated Card Decline object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_route_decline`.
	CardRouteDecline fields.Field[DeclinedTransactionSourceCardRouteDecline] `json:"card_route_decline,required,nullable"`
}

// MarshalJSON serializes DeclinedTransactionSource into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *DeclinedTransactionSource) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r DeclinedTransactionSource) String() (result string) {
	return fmt.Sprintf("&DeclinedTransactionSource{Category:%s ACHDecline:%s CardDecline:%s CheckDecline:%s InboundRealTimePaymentsTransferDecline:%s InternationalACHDecline:%s CardRouteDecline:%s}", r.Category, r.ACHDecline, r.CardDecline, r.CheckDecline, r.InboundRealTimePaymentsTransferDecline, r.InternationalACHDecline, r.CardRouteDecline)
}

type DeclinedTransactionSourceCategory string

const (
	DeclinedTransactionSourceCategoryACHDecline                             DeclinedTransactionSourceCategory = "ach_decline"
	DeclinedTransactionSourceCategoryCardDecline                            DeclinedTransactionSourceCategory = "card_decline"
	DeclinedTransactionSourceCategoryCheckDecline                           DeclinedTransactionSourceCategory = "check_decline"
	DeclinedTransactionSourceCategoryInboundRealTimePaymentsTransferDecline DeclinedTransactionSourceCategory = "inbound_real_time_payments_transfer_decline"
	DeclinedTransactionSourceCategoryInternationalACHDecline                DeclinedTransactionSourceCategory = "international_ach_decline"
	DeclinedTransactionSourceCategoryCardRouteDecline                       DeclinedTransactionSourceCategory = "card_route_decline"
	DeclinedTransactionSourceCategoryOther                                  DeclinedTransactionSourceCategory = "other"
)

type DeclinedTransactionSourceACHDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount                             fields.Field[int64]  `json:"amount,required"`
	OriginatorCompanyName              fields.Field[string] `json:"originator_company_name,required"`
	OriginatorCompanyDescriptiveDate   fields.Field[string] `json:"originator_company_descriptive_date,required,nullable"`
	OriginatorCompanyDiscretionaryData fields.Field[string] `json:"originator_company_discretionary_data,required,nullable"`
	OriginatorCompanyID                fields.Field[string] `json:"originator_company_id,required"`
	// Why the ACH transfer was declined.
	Reason           fields.Field[DeclinedTransactionSourceACHDeclineReason] `json:"reason,required"`
	ReceiverIDNumber fields.Field[string]                                    `json:"receiver_id_number,required,nullable"`
	ReceiverName     fields.Field[string]                                    `json:"receiver_name,required,nullable"`
	TraceNumber      fields.Field[string]                                    `json:"trace_number,required"`
}

// MarshalJSON serializes DeclinedTransactionSourceACHDecline into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *DeclinedTransactionSourceACHDecline) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r DeclinedTransactionSourceACHDecline) String() (result string) {
	return fmt.Sprintf("&DeclinedTransactionSourceACHDecline{Amount:%s OriginatorCompanyName:%s OriginatorCompanyDescriptiveDate:%s OriginatorCompanyDiscretionaryData:%s OriginatorCompanyID:%s Reason:%s ReceiverIDNumber:%s ReceiverName:%s TraceNumber:%s}", r.Amount, r.OriginatorCompanyName, r.OriginatorCompanyDescriptiveDate, r.OriginatorCompanyDiscretionaryData, r.OriginatorCompanyID, r.Reason, r.ReceiverIDNumber, r.ReceiverName, r.TraceNumber)
}

type DeclinedTransactionSourceACHDeclineReason string

const (
	DeclinedTransactionSourceACHDeclineReasonACHRouteCanceled             DeclinedTransactionSourceACHDeclineReason = "ach_route_canceled"
	DeclinedTransactionSourceACHDeclineReasonACHRouteDisabled             DeclinedTransactionSourceACHDeclineReason = "ach_route_disabled"
	DeclinedTransactionSourceACHDeclineReasonBreachesLimit                DeclinedTransactionSourceACHDeclineReason = "breaches_limit"
	DeclinedTransactionSourceACHDeclineReasonCreditEntryRefusedByReceiver DeclinedTransactionSourceACHDeclineReason = "credit_entry_refused_by_receiver"
	DeclinedTransactionSourceACHDeclineReasonDuplicateReturn              DeclinedTransactionSourceACHDeclineReason = "duplicate_return"
	DeclinedTransactionSourceACHDeclineReasonEntityNotActive              DeclinedTransactionSourceACHDeclineReason = "entity_not_active"
	DeclinedTransactionSourceACHDeclineReasonTransactionNotAllowed        DeclinedTransactionSourceACHDeclineReason = "transaction_not_allowed"
	DeclinedTransactionSourceACHDeclineReasonGroupLocked                  DeclinedTransactionSourceACHDeclineReason = "group_locked"
	DeclinedTransactionSourceACHDeclineReasonInsufficientFunds            DeclinedTransactionSourceACHDeclineReason = "insufficient_funds"
	DeclinedTransactionSourceACHDeclineReasonNoACHRoute                   DeclinedTransactionSourceACHDeclineReason = "no_ach_route"
	DeclinedTransactionSourceACHDeclineReasonOriginatorRequest            DeclinedTransactionSourceACHDeclineReason = "originator_request"
)

type DeclinedTransactionSourceCardDecline struct {
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
	Network fields.Field[DeclinedTransactionSourceCardDeclineNetwork] `json:"network,required"`
	// Fields specific to the `network`
	NetworkDetails fields.Field[DeclinedTransactionSourceCardDeclineNetworkDetails] `json:"network_details,required"`
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency fields.Field[DeclinedTransactionSourceCardDeclineCurrency] `json:"currency,required"`
	// Why the transaction was declined.
	Reason fields.Field[DeclinedTransactionSourceCardDeclineReason] `json:"reason,required"`
	// The state the merchant resides in.
	MerchantState fields.Field[string] `json:"merchant_state,required,nullable"`
	// The identifier of the Real-Time Decision sent to approve or decline this
	// transaction.
	RealTimeDecisionID fields.Field[string] `json:"real_time_decision_id,required,nullable"`
	// If the authorization was attempted using a Digital Wallet Token (such as an
	// Apple Pay purchase), the identifier of the token that was used.
	DigitalWalletTokenID fields.Field[string] `json:"digital_wallet_token_id,required,nullable"`
}

// MarshalJSON serializes DeclinedTransactionSourceCardDecline into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *DeclinedTransactionSourceCardDecline) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r DeclinedTransactionSourceCardDecline) String() (result string) {
	return fmt.Sprintf("&DeclinedTransactionSourceCardDecline{MerchantAcceptorID:%s MerchantDescriptor:%s MerchantCategoryCode:%s MerchantCity:%s MerchantCountry:%s Network:%s NetworkDetails:%s Amount:%s Currency:%s Reason:%s MerchantState:%s RealTimeDecisionID:%s DigitalWalletTokenID:%s}", r.MerchantAcceptorID, r.MerchantDescriptor, r.MerchantCategoryCode, r.MerchantCity, r.MerchantCountry, r.Network, r.NetworkDetails, r.Amount, r.Currency, r.Reason, r.MerchantState, r.RealTimeDecisionID, r.DigitalWalletTokenID)
}

type DeclinedTransactionSourceCardDeclineNetwork string

const (
	DeclinedTransactionSourceCardDeclineNetworkVisa DeclinedTransactionSourceCardDeclineNetwork = "visa"
)

type DeclinedTransactionSourceCardDeclineNetworkDetails struct {
	// Fields specific to the `visa` network
	Visa fields.Field[DeclinedTransactionSourceCardDeclineNetworkDetailsVisa] `json:"visa,required"`
}

// MarshalJSON serializes DeclinedTransactionSourceCardDeclineNetworkDetails into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *DeclinedTransactionSourceCardDeclineNetworkDetails) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r DeclinedTransactionSourceCardDeclineNetworkDetails) String() (result string) {
	return fmt.Sprintf("&DeclinedTransactionSourceCardDeclineNetworkDetails{Visa:%s}", r.Visa)
}

type DeclinedTransactionSourceCardDeclineNetworkDetailsVisa struct {
	// For electronic commerce transactions, this identifies the level of security used
	// in obtaining the customer's payment credential. For mail or telephone order
	// transactions, identifies the type of mail or telephone order.
	ElectronicCommerceIndicator fields.Field[DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator] `json:"electronic_commerce_indicator,required,nullable"`
	// The method used to enter the cardholder's primary account number and card
	// expiration date
	PointOfServiceEntryMode fields.Field[PointOfServiceEntryMode] `json:"point_of_service_entry_mode,required,nullable"`
}

// MarshalJSON serializes DeclinedTransactionSourceCardDeclineNetworkDetailsVisa
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *DeclinedTransactionSourceCardDeclineNetworkDetailsVisa) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r DeclinedTransactionSourceCardDeclineNetworkDetailsVisa) String() (result string) {
	return fmt.Sprintf("&DeclinedTransactionSourceCardDeclineNetworkDetailsVisa{ElectronicCommerceIndicator:%s PointOfServiceEntryMode:%s}", r.ElectronicCommerceIndicator, r.PointOfServiceEntryMode)
}

type DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator string

const (
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorMailPhoneOrder                                           DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "mail_phone_order"
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorRecurring                                                DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "recurring"
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorInstallment                                              DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "installment"
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorUnknownMailPhoneOrder                                    DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "unknown_mail_phone_order"
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorSecureElectronicCommerce                                 DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "secure_electronic_commerce"
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransactionAt_3dsCapableMerchant DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction_at_3ds_capable_merchant"
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransaction                      DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction"
	DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonSecureTransaction                                     DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "non_secure_transaction"
)

type DeclinedTransactionSourceCardDeclineCurrency string

const (
	DeclinedTransactionSourceCardDeclineCurrencyCad DeclinedTransactionSourceCardDeclineCurrency = "CAD"
	DeclinedTransactionSourceCardDeclineCurrencyChf DeclinedTransactionSourceCardDeclineCurrency = "CHF"
	DeclinedTransactionSourceCardDeclineCurrencyEur DeclinedTransactionSourceCardDeclineCurrency = "EUR"
	DeclinedTransactionSourceCardDeclineCurrencyGbp DeclinedTransactionSourceCardDeclineCurrency = "GBP"
	DeclinedTransactionSourceCardDeclineCurrencyJpy DeclinedTransactionSourceCardDeclineCurrency = "JPY"
	DeclinedTransactionSourceCardDeclineCurrencyUsd DeclinedTransactionSourceCardDeclineCurrency = "USD"
)

type DeclinedTransactionSourceCardDeclineReason string

const (
	DeclinedTransactionSourceCardDeclineReasonCardNotActive               DeclinedTransactionSourceCardDeclineReason = "card_not_active"
	DeclinedTransactionSourceCardDeclineReasonEntityNotActive             DeclinedTransactionSourceCardDeclineReason = "entity_not_active"
	DeclinedTransactionSourceCardDeclineReasonGroupLocked                 DeclinedTransactionSourceCardDeclineReason = "group_locked"
	DeclinedTransactionSourceCardDeclineReasonInsufficientFunds           DeclinedTransactionSourceCardDeclineReason = "insufficient_funds"
	DeclinedTransactionSourceCardDeclineReasonCvv2Mismatch                DeclinedTransactionSourceCardDeclineReason = "cvv2_mismatch"
	DeclinedTransactionSourceCardDeclineReasonTransactionNotAllowed       DeclinedTransactionSourceCardDeclineReason = "transaction_not_allowed"
	DeclinedTransactionSourceCardDeclineReasonBreachesLimit               DeclinedTransactionSourceCardDeclineReason = "breaches_limit"
	DeclinedTransactionSourceCardDeclineReasonWebhookDeclined             DeclinedTransactionSourceCardDeclineReason = "webhook_declined"
	DeclinedTransactionSourceCardDeclineReasonWebhookTimedOut             DeclinedTransactionSourceCardDeclineReason = "webhook_timed_out"
	DeclinedTransactionSourceCardDeclineReasonDeclinedByStandInProcessing DeclinedTransactionSourceCardDeclineReason = "declined_by_stand_in_processing"
	DeclinedTransactionSourceCardDeclineReasonInvalidPhysicalCard         DeclinedTransactionSourceCardDeclineReason = "invalid_physical_card"
)

type DeclinedTransactionSourceCheckDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount        fields.Field[int64]  `json:"amount,required"`
	AuxiliaryOnUs fields.Field[string] `json:"auxiliary_on_us,required,nullable"`
	// Why the check was declined.
	Reason fields.Field[DeclinedTransactionSourceCheckDeclineReason] `json:"reason,required"`
}

// MarshalJSON serializes DeclinedTransactionSourceCheckDecline into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *DeclinedTransactionSourceCheckDecline) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r DeclinedTransactionSourceCheckDecline) String() (result string) {
	return fmt.Sprintf("&DeclinedTransactionSourceCheckDecline{Amount:%s AuxiliaryOnUs:%s Reason:%s}", r.Amount, r.AuxiliaryOnUs, r.Reason)
}

type DeclinedTransactionSourceCheckDeclineReason string

const (
	DeclinedTransactionSourceCheckDeclineReasonACHRouteCanceled      DeclinedTransactionSourceCheckDeclineReason = "ach_route_canceled"
	DeclinedTransactionSourceCheckDeclineReasonACHRouteDisabled      DeclinedTransactionSourceCheckDeclineReason = "ach_route_disabled"
	DeclinedTransactionSourceCheckDeclineReasonBreachesLimit         DeclinedTransactionSourceCheckDeclineReason = "breaches_limit"
	DeclinedTransactionSourceCheckDeclineReasonEntityNotActive       DeclinedTransactionSourceCheckDeclineReason = "entity_not_active"
	DeclinedTransactionSourceCheckDeclineReasonGroupLocked           DeclinedTransactionSourceCheckDeclineReason = "group_locked"
	DeclinedTransactionSourceCheckDeclineReasonInsufficientFunds     DeclinedTransactionSourceCheckDeclineReason = "insufficient_funds"
	DeclinedTransactionSourceCheckDeclineReasonUnableToLocateAccount DeclinedTransactionSourceCheckDeclineReason = "unable_to_locate_account"
	DeclinedTransactionSourceCheckDeclineReasonUnableToProcess       DeclinedTransactionSourceCheckDeclineReason = "unable_to_process"
	DeclinedTransactionSourceCheckDeclineReasonReferToImage          DeclinedTransactionSourceCheckDeclineReason = "refer_to_image"
	DeclinedTransactionSourceCheckDeclineReasonStopPaymentRequested  DeclinedTransactionSourceCheckDeclineReason = "stop_payment_requested"
	DeclinedTransactionSourceCheckDeclineReasonReturned              DeclinedTransactionSourceCheckDeclineReason = "returned"
	DeclinedTransactionSourceCheckDeclineReasonDuplicatePresentment  DeclinedTransactionSourceCheckDeclineReason = "duplicate_presentment"
)

type DeclinedTransactionSourceInboundRealTimePaymentsTransferDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the declined
	// transfer's currency. This will always be "USD" for a Real Time Payments
	// transfer.
	Currency fields.Field[DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency] `json:"currency,required"`
	// Why the transfer was declined.
	Reason fields.Field[DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason] `json:"reason,required"`
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
// DeclinedTransactionSourceInboundRealTimePaymentsTransferDecline into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *DeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r DeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) String() (result string) {
	return fmt.Sprintf("&DeclinedTransactionSourceInboundRealTimePaymentsTransferDecline{Amount:%s Currency:%s Reason:%s CreditorName:%s DebtorName:%s DebtorAccountNumber:%s DebtorRoutingNumber:%s TransactionIdentification:%s RemittanceInformation:%s}", r.Amount, r.Currency, r.Reason, r.CreditorName, r.DebtorName, r.DebtorAccountNumber, r.DebtorRoutingNumber, r.TransactionIdentification, r.RemittanceInformation)
}

type DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency string

const (
	DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyCad DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "CAD"
	DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyChf DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "CHF"
	DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyEur DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "EUR"
	DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyGbp DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "GBP"
	DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyJpy DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "JPY"
	DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyUsd DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "USD"
)

type DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason string

const (
	DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonAccountNumberCanceled      DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "account_number_canceled"
	DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonAccountNumberDisabled      DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "account_number_disabled"
	DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonGroupLocked                DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "group_locked"
	DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonEntityNotActive            DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "entity_not_active"
	DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonRealTimePaymentsNotEnabled DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "real_time_payments_not_enabled"
)

type DeclinedTransactionSourceInternationalACHDecline struct {
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

// MarshalJSON serializes DeclinedTransactionSourceInternationalACHDecline into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *DeclinedTransactionSourceInternationalACHDecline) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r DeclinedTransactionSourceInternationalACHDecline) String() (result string) {
	return fmt.Sprintf("&DeclinedTransactionSourceInternationalACHDecline{Amount:%s ForeignExchangeIndicator:%s ForeignExchangeReferenceIndicator:%s ForeignExchangeReference:%s DestinationCountryCode:%s DestinationCurrencyCode:%s ForeignPaymentAmount:%s ForeignTraceNumber:%s InternationalTransactionTypeCode:%s OriginatingCurrencyCode:%s OriginatingDepositoryFinancialInstitutionName:%s OriginatingDepositoryFinancialInstitutionIDQualifier:%s OriginatingDepositoryFinancialInstitutionID:%s OriginatingDepositoryFinancialInstitutionBranchCountry:%s OriginatorCity:%s OriginatorCompanyEntryDescription:%s OriginatorCountry:%s OriginatorIdentification:%s OriginatorName:%s OriginatorPostalCode:%s OriginatorStreetAddress:%s OriginatorStateOrProvince:%s PaymentRelatedInformation:%s PaymentRelatedInformation2:%s ReceiverIdentificationNumber:%s ReceiverStreetAddress:%s ReceiverCity:%s ReceiverStateOrProvince:%s ReceiverCountry:%s ReceiverPostalCode:%s ReceivingCompanyOrIndividualName:%s ReceivingDepositoryFinancialInstitutionName:%s ReceivingDepositoryFinancialInstitutionIDQualifier:%s ReceivingDepositoryFinancialInstitutionID:%s ReceivingDepositoryFinancialInstitutionCountry:%s TraceNumber:%s}", r.Amount, r.ForeignExchangeIndicator, r.ForeignExchangeReferenceIndicator, r.ForeignExchangeReference, r.DestinationCountryCode, r.DestinationCurrencyCode, r.ForeignPaymentAmount, r.ForeignTraceNumber, r.InternationalTransactionTypeCode, r.OriginatingCurrencyCode, r.OriginatingDepositoryFinancialInstitutionName, r.OriginatingDepositoryFinancialInstitutionIDQualifier, r.OriginatingDepositoryFinancialInstitutionID, r.OriginatingDepositoryFinancialInstitutionBranchCountry, r.OriginatorCity, r.OriginatorCompanyEntryDescription, r.OriginatorCountry, r.OriginatorIdentification, r.OriginatorName, r.OriginatorPostalCode, r.OriginatorStreetAddress, r.OriginatorStateOrProvince, r.PaymentRelatedInformation, r.PaymentRelatedInformation2, r.ReceiverIdentificationNumber, r.ReceiverStreetAddress, r.ReceiverCity, r.ReceiverStateOrProvince, r.ReceiverCountry, r.ReceiverPostalCode, r.ReceivingCompanyOrIndividualName, r.ReceivingDepositoryFinancialInstitutionName, r.ReceivingDepositoryFinancialInstitutionIDQualifier, r.ReceivingDepositoryFinancialInstitutionID, r.ReceivingDepositoryFinancialInstitutionCountry, r.TraceNumber)
}

type DeclinedTransactionSourceCardRouteDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency             fields.Field[DeclinedTransactionSourceCardRouteDeclineCurrency] `json:"currency,required"`
	MerchantAcceptorID   fields.Field[string]                                            `json:"merchant_acceptor_id,required"`
	MerchantCity         fields.Field[string]                                            `json:"merchant_city,required,nullable"`
	MerchantCountry      fields.Field[string]                                            `json:"merchant_country,required"`
	MerchantDescriptor   fields.Field[string]                                            `json:"merchant_descriptor,required"`
	MerchantState        fields.Field[string]                                            `json:"merchant_state,required,nullable"`
	MerchantCategoryCode fields.Field[string]                                            `json:"merchant_category_code,required,nullable"`
}

// MarshalJSON serializes DeclinedTransactionSourceCardRouteDecline into an array
// of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *DeclinedTransactionSourceCardRouteDecline) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r DeclinedTransactionSourceCardRouteDecline) String() (result string) {
	return fmt.Sprintf("&DeclinedTransactionSourceCardRouteDecline{Amount:%s Currency:%s MerchantAcceptorID:%s MerchantCity:%s MerchantCountry:%s MerchantDescriptor:%s MerchantState:%s MerchantCategoryCode:%s}", r.Amount, r.Currency, r.MerchantAcceptorID, r.MerchantCity, r.MerchantCountry, r.MerchantDescriptor, r.MerchantState, r.MerchantCategoryCode)
}

type DeclinedTransactionSourceCardRouteDeclineCurrency string

const (
	DeclinedTransactionSourceCardRouteDeclineCurrencyCad DeclinedTransactionSourceCardRouteDeclineCurrency = "CAD"
	DeclinedTransactionSourceCardRouteDeclineCurrencyChf DeclinedTransactionSourceCardRouteDeclineCurrency = "CHF"
	DeclinedTransactionSourceCardRouteDeclineCurrencyEur DeclinedTransactionSourceCardRouteDeclineCurrency = "EUR"
	DeclinedTransactionSourceCardRouteDeclineCurrencyGbp DeclinedTransactionSourceCardRouteDeclineCurrency = "GBP"
	DeclinedTransactionSourceCardRouteDeclineCurrencyJpy DeclinedTransactionSourceCardRouteDeclineCurrency = "JPY"
	DeclinedTransactionSourceCardRouteDeclineCurrencyUsd DeclinedTransactionSourceCardRouteDeclineCurrency = "USD"
)

type DeclinedTransactionType string

const (
	DeclinedTransactionTypeDeclinedTransaction DeclinedTransactionType = "declined_transaction"
)

type DeclinedTransactionListParams struct {
	// Return the page of entries after this one.
	Cursor fields.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit fields.Field[int64] `query:"limit"`
	// Filter Declined Transactions to ones belonging to the specified Account.
	AccountID fields.Field[string] `query:"account_id"`
	// Filter Declined Transactions to those belonging to the specified route.
	RouteID   fields.Field[string]                                 `query:"route_id"`
	CreatedAt fields.Field[DeclinedTransactionListParamsCreatedAt] `query:"created_at"`
}

// URLQuery serializes DeclinedTransactionListParams into a url.Values of the query
// parameters associated with this value
func (r *DeclinedTransactionListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r DeclinedTransactionListParams) String() (result string) {
	return fmt.Sprintf("&DeclinedTransactionListParams{Cursor:%s Limit:%s AccountID:%s RouteID:%s CreatedAt:%s}", r.Cursor, r.Limit, r.AccountID, r.RouteID, r.CreatedAt)
}

type DeclinedTransactionListParamsCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After fields.Field[time.Time] `query:"after" format:"date-time"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before fields.Field[time.Time] `query:"before" format:"date-time"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter fields.Field[time.Time] `query:"on_or_after" format:"date-time"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore fields.Field[time.Time] `query:"on_or_before" format:"date-time"`
}

// URLQuery serializes DeclinedTransactionListParamsCreatedAt into a url.Values of
// the query parameters associated with this value
func (r *DeclinedTransactionListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r DeclinedTransactionListParamsCreatedAt) String() (result string) {
	return fmt.Sprintf("&DeclinedTransactionListParamsCreatedAt{After:%s Before:%s OnOrAfter:%s OnOrBefore:%s}", r.After, r.Before, r.OnOrAfter, r.OnOrBefore)
}

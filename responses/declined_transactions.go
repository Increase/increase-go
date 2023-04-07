package responses

import (
	"time"

	pjson "github.com/increase/increase-go/core/json"
)

type DeclinedTransaction struct {
	// The identifier for the Account the Declined Transaction belongs to.
	AccountID string `json:"account_id,required"`
	// The Declined Transaction amount in the minor unit of its currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Declined
	// Transaction's currency. This will match the currency on the Declined
	// Transcation's Account.
	Currency DeclinedTransactionCurrency `json:"currency,required"`
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
	RouteType DeclinedTransactionRouteType `json:"route_type,required,nullable"`
	// This is an object giving more details on the network-level event that caused the
	// Declined Transaction. For example, for a card transaction this lists the
	// merchant's industry and location. Note that for backwards compatibility reasons,
	// additional undocumented keys may appear in this object. These should be treated
	// as deprecated and will be removed in the future.
	Source DeclinedTransactionSource `json:"source,required"`
	// A constant representing the object's type. For this resource it will always be
	// `declined_transaction`.
	Type DeclinedTransactionType `json:"type,required"`
	JSON DeclinedTransactionJSON
}

type DeclinedTransactionJSON struct {
	AccountID   pjson.Metadata
	Amount      pjson.Metadata
	Currency    pjson.Metadata
	CreatedAt   pjson.Metadata
	Description pjson.Metadata
	ID          pjson.Metadata
	RouteID     pjson.Metadata
	RouteType   pjson.Metadata
	Source      pjson.Metadata
	Type        pjson.Metadata
	Raw         []byte
	Extras      map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into DeclinedTransaction using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *DeclinedTransaction) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	Category DeclinedTransactionSourceCategory `json:"category,required"`
	// A ACH Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `ach_decline`.
	ACHDecline DeclinedTransactionSourceACHDecline `json:"ach_decline,required,nullable"`
	// A Card Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_decline`.
	CardDecline DeclinedTransactionSourceCardDecline `json:"card_decline,required,nullable"`
	// A Check Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `check_decline`.
	CheckDecline DeclinedTransactionSourceCheckDecline `json:"check_decline,required,nullable"`
	// A Inbound Real Time Payments Transfer Decline object. This field will be present
	// in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_decline`.
	InboundRealTimePaymentsTransferDecline DeclinedTransactionSourceInboundRealTimePaymentsTransferDecline `json:"inbound_real_time_payments_transfer_decline,required,nullable"`
	// A International ACH Decline object. This field will be present in the JSON
	// response if and only if `category` is equal to `international_ach_decline`.
	InternationalACHDecline DeclinedTransactionSourceInternationalACHDecline `json:"international_ach_decline,required,nullable"`
	// A Deprecated Card Decline object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_route_decline`.
	CardRouteDecline DeclinedTransactionSourceCardRouteDecline `json:"card_route_decline,required,nullable"`
	JSON             DeclinedTransactionSourceJSON
}

type DeclinedTransactionSourceJSON struct {
	Category                               pjson.Metadata
	ACHDecline                             pjson.Metadata
	CardDecline                            pjson.Metadata
	CheckDecline                           pjson.Metadata
	InboundRealTimePaymentsTransferDecline pjson.Metadata
	InternationalACHDecline                pjson.Metadata
	CardRouteDecline                       pjson.Metadata
	Raw                                    []byte
	Extras                                 map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into DeclinedTransactionSource
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *DeclinedTransactionSource) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	Amount                             int64  `json:"amount,required"`
	OriginatorCompanyName              string `json:"originator_company_name,required"`
	OriginatorCompanyDescriptiveDate   string `json:"originator_company_descriptive_date,required,nullable"`
	OriginatorCompanyDiscretionaryData string `json:"originator_company_discretionary_data,required,nullable"`
	OriginatorCompanyID                string `json:"originator_company_id,required"`
	// Why the ACH transfer was declined.
	Reason           DeclinedTransactionSourceACHDeclineReason `json:"reason,required"`
	ReceiverIDNumber string                                    `json:"receiver_id_number,required,nullable"`
	ReceiverName     string                                    `json:"receiver_name,required,nullable"`
	TraceNumber      string                                    `json:"trace_number,required"`
	JSON             DeclinedTransactionSourceACHDeclineJSON
}

type DeclinedTransactionSourceACHDeclineJSON struct {
	Amount                             pjson.Metadata
	OriginatorCompanyName              pjson.Metadata
	OriginatorCompanyDescriptiveDate   pjson.Metadata
	OriginatorCompanyDiscretionaryData pjson.Metadata
	OriginatorCompanyID                pjson.Metadata
	Reason                             pjson.Metadata
	ReceiverIDNumber                   pjson.Metadata
	ReceiverName                       pjson.Metadata
	TraceNumber                        pjson.Metadata
	Raw                                []byte
	Extras                             map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// DeclinedTransactionSourceACHDecline using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *DeclinedTransactionSourceACHDecline) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	Network DeclinedTransactionSourceCardDeclineNetwork `json:"network,required"`
	// Fields specific to the `network`
	NetworkDetails DeclinedTransactionSourceCardDeclineNetworkDetails `json:"network_details,required"`
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency DeclinedTransactionSourceCardDeclineCurrency `json:"currency,required"`
	// Why the transaction was declined.
	Reason DeclinedTransactionSourceCardDeclineReason `json:"reason,required"`
	// The state the merchant resides in.
	MerchantState string `json:"merchant_state,required,nullable"`
	// The identifier of the Real-Time Decision sent to approve or decline this
	// transaction.
	RealTimeDecisionID string `json:"real_time_decision_id,required,nullable"`
	// If the authorization was attempted using a Digital Wallet Token (such as an
	// Apple Pay purchase), the identifier of the token that was used.
	DigitalWalletTokenID string `json:"digital_wallet_token_id,required,nullable"`
	JSON                 DeclinedTransactionSourceCardDeclineJSON
}

type DeclinedTransactionSourceCardDeclineJSON struct {
	MerchantAcceptorID   pjson.Metadata
	MerchantDescriptor   pjson.Metadata
	MerchantCategoryCode pjson.Metadata
	MerchantCity         pjson.Metadata
	MerchantCountry      pjson.Metadata
	Network              pjson.Metadata
	NetworkDetails       pjson.Metadata
	Amount               pjson.Metadata
	Currency             pjson.Metadata
	Reason               pjson.Metadata
	MerchantState        pjson.Metadata
	RealTimeDecisionID   pjson.Metadata
	DigitalWalletTokenID pjson.Metadata
	Raw                  []byte
	Extras               map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// DeclinedTransactionSourceCardDecline using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *DeclinedTransactionSourceCardDecline) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type DeclinedTransactionSourceCardDeclineNetwork string

const (
	DeclinedTransactionSourceCardDeclineNetworkVisa DeclinedTransactionSourceCardDeclineNetwork = "visa"
)

type DeclinedTransactionSourceCardDeclineNetworkDetails struct {
	// Fields specific to the `visa` network
	Visa DeclinedTransactionSourceCardDeclineNetworkDetailsVisa `json:"visa,required"`
	JSON DeclinedTransactionSourceCardDeclineNetworkDetailsJSON
}

type DeclinedTransactionSourceCardDeclineNetworkDetailsJSON struct {
	Visa   pjson.Metadata
	Raw    []byte
	Extras map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// DeclinedTransactionSourceCardDeclineNetworkDetails using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *DeclinedTransactionSourceCardDeclineNetworkDetails) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type DeclinedTransactionSourceCardDeclineNetworkDetailsVisa struct {
	// For electronic commerce transactions, this identifies the level of security used
	// in obtaining the customer's payment credential. For mail or telephone order
	// transactions, identifies the type of mail or telephone order.
	ElectronicCommerceIndicator DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator `json:"electronic_commerce_indicator,required,nullable"`
	// The method used to enter the cardholder's primary account number and card
	// expiration date
	PointOfServiceEntryMode PointOfServiceEntryMode `json:"point_of_service_entry_mode,required,nullable"`
	JSON                    DeclinedTransactionSourceCardDeclineNetworkDetailsVisaJSON
}

type DeclinedTransactionSourceCardDeclineNetworkDetailsVisaJSON struct {
	ElectronicCommerceIndicator pjson.Metadata
	PointOfServiceEntryMode     pjson.Metadata
	Raw                         []byte
	Extras                      map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// DeclinedTransactionSourceCardDeclineNetworkDetailsVisa using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *DeclinedTransactionSourceCardDeclineNetworkDetailsVisa) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	Amount        int64  `json:"amount,required"`
	AuxiliaryOnUs string `json:"auxiliary_on_us,required,nullable"`
	// Why the check was declined.
	Reason DeclinedTransactionSourceCheckDeclineReason `json:"reason,required"`
	JSON   DeclinedTransactionSourceCheckDeclineJSON
}

type DeclinedTransactionSourceCheckDeclineJSON struct {
	Amount        pjson.Metadata
	AuxiliaryOnUs pjson.Metadata
	Reason        pjson.Metadata
	Raw           []byte
	Extras        map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// DeclinedTransactionSourceCheckDecline using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *DeclinedTransactionSourceCheckDecline) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the declined
	// transfer's currency. This will always be "USD" for a Real Time Payments
	// transfer.
	Currency DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency `json:"currency,required"`
	// Why the transfer was declined.
	Reason DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason `json:"reason,required"`
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
	JSON                  DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineJSON
}

type DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineJSON struct {
	Amount                    pjson.Metadata
	Currency                  pjson.Metadata
	Reason                    pjson.Metadata
	CreditorName              pjson.Metadata
	DebtorName                pjson.Metadata
	DebtorAccountNumber       pjson.Metadata
	DebtorRoutingNumber       pjson.Metadata
	TransactionIdentification pjson.Metadata
	RemittanceInformation     pjson.Metadata
	Raw                       []byte
	Extras                    map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// DeclinedTransactionSourceInboundRealTimePaymentsTransferDecline using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *DeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	JSON                                                   DeclinedTransactionSourceInternationalACHDeclineJSON
}

type DeclinedTransactionSourceInternationalACHDeclineJSON struct {
	Amount                                                 pjson.Metadata
	ForeignExchangeIndicator                               pjson.Metadata
	ForeignExchangeReferenceIndicator                      pjson.Metadata
	ForeignExchangeReference                               pjson.Metadata
	DestinationCountryCode                                 pjson.Metadata
	DestinationCurrencyCode                                pjson.Metadata
	ForeignPaymentAmount                                   pjson.Metadata
	ForeignTraceNumber                                     pjson.Metadata
	InternationalTransactionTypeCode                       pjson.Metadata
	OriginatingCurrencyCode                                pjson.Metadata
	OriginatingDepositoryFinancialInstitutionName          pjson.Metadata
	OriginatingDepositoryFinancialInstitutionIDQualifier   pjson.Metadata
	OriginatingDepositoryFinancialInstitutionID            pjson.Metadata
	OriginatingDepositoryFinancialInstitutionBranchCountry pjson.Metadata
	OriginatorCity                                         pjson.Metadata
	OriginatorCompanyEntryDescription                      pjson.Metadata
	OriginatorCountry                                      pjson.Metadata
	OriginatorIdentification                               pjson.Metadata
	OriginatorName                                         pjson.Metadata
	OriginatorPostalCode                                   pjson.Metadata
	OriginatorStreetAddress                                pjson.Metadata
	OriginatorStateOrProvince                              pjson.Metadata
	PaymentRelatedInformation                              pjson.Metadata
	PaymentRelatedInformation2                             pjson.Metadata
	ReceiverIdentificationNumber                           pjson.Metadata
	ReceiverStreetAddress                                  pjson.Metadata
	ReceiverCity                                           pjson.Metadata
	ReceiverStateOrProvince                                pjson.Metadata
	ReceiverCountry                                        pjson.Metadata
	ReceiverPostalCode                                     pjson.Metadata
	ReceivingCompanyOrIndividualName                       pjson.Metadata
	ReceivingDepositoryFinancialInstitutionName            pjson.Metadata
	ReceivingDepositoryFinancialInstitutionIDQualifier     pjson.Metadata
	ReceivingDepositoryFinancialInstitutionID              pjson.Metadata
	ReceivingDepositoryFinancialInstitutionCountry         pjson.Metadata
	TraceNumber                                            pjson.Metadata
	Raw                                                    []byte
	Extras                                                 map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// DeclinedTransactionSourceInternationalACHDecline using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *DeclinedTransactionSourceInternationalACHDecline) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type DeclinedTransactionSourceCardRouteDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency             DeclinedTransactionSourceCardRouteDeclineCurrency `json:"currency,required"`
	MerchantAcceptorID   string                                            `json:"merchant_acceptor_id,required"`
	MerchantCity         string                                            `json:"merchant_city,required,nullable"`
	MerchantCountry      string                                            `json:"merchant_country,required"`
	MerchantDescriptor   string                                            `json:"merchant_descriptor,required"`
	MerchantState        string                                            `json:"merchant_state,required,nullable"`
	MerchantCategoryCode string                                            `json:"merchant_category_code,required,nullable"`
	JSON                 DeclinedTransactionSourceCardRouteDeclineJSON
}

type DeclinedTransactionSourceCardRouteDeclineJSON struct {
	Amount               pjson.Metadata
	Currency             pjson.Metadata
	MerchantAcceptorID   pjson.Metadata
	MerchantCity         pjson.Metadata
	MerchantCountry      pjson.Metadata
	MerchantDescriptor   pjson.Metadata
	MerchantState        pjson.Metadata
	MerchantCategoryCode pjson.Metadata
	Raw                  []byte
	Extras               map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// DeclinedTransactionSourceCardRouteDecline using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *DeclinedTransactionSourceCardRouteDecline) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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

type DeclinedTransactionListResponse struct {
	// The contents of the list.
	Data []DeclinedTransaction `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       DeclinedTransactionListResponseJSON
}

type DeclinedTransactionListResponseJSON struct {
	Data       pjson.Metadata
	NextCursor pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// DeclinedTransactionListResponse using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *DeclinedTransactionListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

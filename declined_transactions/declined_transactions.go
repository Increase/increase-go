package declined_transactions

import "context"
import "increase/core"
import "fmt"
import "increase/pagination"

type DeclinedTransactionService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewDeclinedTransactionService(requester core.Requester) (r *DeclinedTransactionService) {
	r = &DeclinedTransactionService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

//
type DeclinedTransaction struct {
	// The identifier for the Account the Declined Transaction belongs to.
	AccountID string `json:"account_id"`
	// The Declined Transaction amount in the minor unit of its currency. For dollars,
	// for example, this is cents.
	Amount int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Declined
	// Transaction's currency. This will match the currency on the Declined
	// Transcation's Account.
	Currency DeclinedTransactionCurrency `json:"currency"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the
	// Transaction occured.
	CreatedAt string `json:"created_at"`
	// This is the description the vendor provides.
	Description string `json:"description"`
	// The Declined Transaction identifier.
	ID string `json:"id"`
	// The identifier for the route this Declined Transaction came through. Routes are
	// things like cards and ACH details.
	RouteID string `json:"route_id"`
	// The type of the route this Declined Transaction came through.
	RouteType *string `json:"route_type"`
	// This is an object giving more details on the network-level event that caused the
	// Declined Transaction. For example, for a card transaction this lists the
	// merchant's industry and location. Note that for backwards compatibility reasons,
	// additional undocumented keys may appear in this object. These should be treated
	// as deprecated and will be removed in the future.
	Source DeclinedTransactionSource `json:"source"`
	// A constant representing the object's type. For this resource it will always be
	// `declined_transaction`.
	Type DeclinedTransactionType `json:"type"`
}

// The type of the route this Declined Transaction came through.
func (r *DeclinedTransaction) GetRouteType() (RouteType string) {
	if r != nil && r.RouteType != nil {
		RouteType = *r.RouteType
	}
	return
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

//
type DeclinedTransactionSource struct {
	// The type of decline that took place. We may add additional possible values for
	// this enum over time; your application should be able to handle such additions
	// gracefully.
	Category DeclinedTransactionSourceCategory `json:"category"`
	// A ACH Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `ach_decline`.
	ACHDecline *DeclinedTransactionSourceACHDecline `json:"ach_decline"`
	// A Card Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_decline`.
	CardDecline *DeclinedTransactionSourceCardDecline `json:"card_decline"`
	// A Check Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `check_decline`.
	CheckDecline *DeclinedTransactionSourceCheckDecline `json:"check_decline"`
	// A Inbound Real Time Payments Transfer Decline object. This field will be present
	// in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_decline`.
	InboundRealTimePaymentsTransferDecline *DeclinedTransactionSourceInboundRealTimePaymentsTransferDecline `json:"inbound_real_time_payments_transfer_decline"`
	// A International ACH Decline object. This field will be present in the JSON
	// response if and only if `category` is equal to `international_ach_decline`.
	InternationalACHDecline *DeclinedTransactionSourceInternationalACHDecline `json:"international_ach_decline"`
	// A Deprecated Card Decline object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_route_decline`.
	CardRouteDecline *DeclinedTransactionSourceCardRouteDecline `json:"card_route_decline"`
}

// A ACH Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `ach_decline`.
func (r *DeclinedTransactionSource) GetACHDecline() (ACHDecline DeclinedTransactionSourceACHDecline) {
	if r != nil && r.ACHDecline != nil {
		ACHDecline = *r.ACHDecline
	}
	return
}

// A Card Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `card_decline`.
func (r *DeclinedTransactionSource) GetCardDecline() (CardDecline DeclinedTransactionSourceCardDecline) {
	if r != nil && r.CardDecline != nil {
		CardDecline = *r.CardDecline
	}
	return
}

// A Check Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `check_decline`.
func (r *DeclinedTransactionSource) GetCheckDecline() (CheckDecline DeclinedTransactionSourceCheckDecline) {
	if r != nil && r.CheckDecline != nil {
		CheckDecline = *r.CheckDecline
	}
	return
}

// A Inbound Real Time Payments Transfer Decline object. This field will be present
// in the JSON response if and only if `category` is equal to
// `inbound_real_time_payments_transfer_decline`.
func (r *DeclinedTransactionSource) GetInboundRealTimePaymentsTransferDecline() (InboundRealTimePaymentsTransferDecline DeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) {
	if r != nil && r.InboundRealTimePaymentsTransferDecline != nil {
		InboundRealTimePaymentsTransferDecline = *r.InboundRealTimePaymentsTransferDecline
	}
	return
}

// A International ACH Decline object. This field will be present in the JSON
// response if and only if `category` is equal to `international_ach_decline`.
func (r *DeclinedTransactionSource) GetInternationalACHDecline() (InternationalACHDecline DeclinedTransactionSourceInternationalACHDecline) {
	if r != nil && r.InternationalACHDecline != nil {
		InternationalACHDecline = *r.InternationalACHDecline
	}
	return
}

// A Deprecated Card Decline object. This field will be present in the JSON
// response if and only if `category` is equal to `card_route_decline`.
func (r *DeclinedTransactionSource) GetCardRouteDecline() (CardRouteDecline DeclinedTransactionSourceCardRouteDecline) {
	if r != nil && r.CardRouteDecline != nil {
		CardRouteDecline = *r.CardRouteDecline
	}
	return
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

//
type DeclinedTransactionSourceACHDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount int `json:"amount"`
	//
	OriginatorCompanyName string `json:"originator_company_name"`
	//
	OriginatorCompanyDescriptiveDate *string `json:"originator_company_descriptive_date"`
	//
	OriginatorCompanyDiscretionaryData *string `json:"originator_company_discretionary_data"`
	//
	OriginatorCompanyID string `json:"originator_company_id"`
	// Why the ACH transfer was declined.
	Reason DeclinedTransactionSourceACHDeclineReason `json:"reason"`
	//
	ReceiverIDNumber *string `json:"receiver_id_number"`
	//
	ReceiverName *string `json:"receiver_name"`
	//
	TraceNumber string `json:"trace_number"`
}

func (r *DeclinedTransactionSourceACHDecline) GetOriginatorCompanyDescriptiveDate() (OriginatorCompanyDescriptiveDate string) {
	if r != nil && r.OriginatorCompanyDescriptiveDate != nil {
		OriginatorCompanyDescriptiveDate = *r.OriginatorCompanyDescriptiveDate
	}
	return
}

func (r *DeclinedTransactionSourceACHDecline) GetOriginatorCompanyDiscretionaryData() (OriginatorCompanyDiscretionaryData string) {
	if r != nil && r.OriginatorCompanyDiscretionaryData != nil {
		OriginatorCompanyDiscretionaryData = *r.OriginatorCompanyDiscretionaryData
	}
	return
}

func (r *DeclinedTransactionSourceACHDecline) GetReceiverIDNumber() (ReceiverIDNumber string) {
	if r != nil && r.ReceiverIDNumber != nil {
		ReceiverIDNumber = *r.ReceiverIDNumber
	}
	return
}

func (r *DeclinedTransactionSourceACHDecline) GetReceiverName() (ReceiverName string) {
	if r != nil && r.ReceiverName != nil {
		ReceiverName = *r.ReceiverName
	}
	return
}

type DeclinedTransactionSourceACHDeclineReason string

const (
	DeclinedTransactionSourceACHDeclineReasonACHRouteCanceled             DeclinedTransactionSourceACHDeclineReason = "ach_route_canceled"
	DeclinedTransactionSourceACHDeclineReasonACHRouteDisabled             DeclinedTransactionSourceACHDeclineReason = "ach_route_disabled"
	DeclinedTransactionSourceACHDeclineReasonNoACHRoute                   DeclinedTransactionSourceACHDeclineReason = "no_ach_route"
	DeclinedTransactionSourceACHDeclineReasonBreachesLimit                DeclinedTransactionSourceACHDeclineReason = "breaches_limit"
	DeclinedTransactionSourceACHDeclineReasonCreditEntryRefusedByReceiver DeclinedTransactionSourceACHDeclineReason = "credit_entry_refused_by_receiver"
	DeclinedTransactionSourceACHDeclineReasonGroupLocked                  DeclinedTransactionSourceACHDeclineReason = "group_locked"
	DeclinedTransactionSourceACHDeclineReasonEntityNotActive              DeclinedTransactionSourceACHDeclineReason = "entity_not_active"
	DeclinedTransactionSourceACHDeclineReasonInsufficientFunds            DeclinedTransactionSourceACHDeclineReason = "insufficient_funds"
	DeclinedTransactionSourceACHDeclineReasonOriginatorRequest            DeclinedTransactionSourceACHDeclineReason = "originator_request"
)

//
type DeclinedTransactionSourceCardDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency DeclinedTransactionSourceCardDeclineCurrency `json:"currency"`
	//
	MerchantAcceptorID string `json:"merchant_acceptor_id"`
	//
	MerchantCity *string `json:"merchant_city"`
	//
	MerchantCountry *string `json:"merchant_country"`
	//
	MerchantDescriptor string `json:"merchant_descriptor"`
	//
	MerchantState *string `json:"merchant_state"`
	//
	MerchantCategoryCode *string `json:"merchant_category_code"`
	// Why the transaction was declined.
	Reason DeclinedTransactionSourceCardDeclineReason `json:"reason"`
	// The identifier of the Real-Time Decision sent to approve or decline this
	// transaction.
	RealTimeDecisionID *string `json:"real_time_decision_id"`
	// If the authorization was attempted using a Digital Wallet Token (such as an
	// Apple Pay purchase), the identifier of the token that was used.
	DigitalWalletTokenID *string `json:"digital_wallet_token_id"`
}

func (r *DeclinedTransactionSourceCardDecline) GetMerchantCity() (MerchantCity string) {
	if r != nil && r.MerchantCity != nil {
		MerchantCity = *r.MerchantCity
	}
	return
}

func (r *DeclinedTransactionSourceCardDecline) GetMerchantCountry() (MerchantCountry string) {
	if r != nil && r.MerchantCountry != nil {
		MerchantCountry = *r.MerchantCountry
	}
	return
}

func (r *DeclinedTransactionSourceCardDecline) GetMerchantState() (MerchantState string) {
	if r != nil && r.MerchantState != nil {
		MerchantState = *r.MerchantState
	}
	return
}

func (r *DeclinedTransactionSourceCardDecline) GetMerchantCategoryCode() (MerchantCategoryCode string) {
	if r != nil && r.MerchantCategoryCode != nil {
		MerchantCategoryCode = *r.MerchantCategoryCode
	}
	return
}

// The identifier of the Real-Time Decision sent to approve or decline this
// transaction.
func (r *DeclinedTransactionSourceCardDecline) GetRealTimeDecisionID() (RealTimeDecisionID string) {
	if r != nil && r.RealTimeDecisionID != nil {
		RealTimeDecisionID = *r.RealTimeDecisionID
	}
	return
}

// If the authorization was attempted using a Digital Wallet Token (such as an
// Apple Pay purchase), the identifier of the token that was used.
func (r *DeclinedTransactionSourceCardDecline) GetDigitalWalletTokenID() (DigitalWalletTokenID string) {
	if r != nil && r.DigitalWalletTokenID != nil {
		DigitalWalletTokenID = *r.DigitalWalletTokenID
	}
	return
}

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
	DeclinedTransactionSourceCardDeclineReasonCardNotActive     DeclinedTransactionSourceCardDeclineReason = "card_not_active"
	DeclinedTransactionSourceCardDeclineReasonEntityNotActive   DeclinedTransactionSourceCardDeclineReason = "entity_not_active"
	DeclinedTransactionSourceCardDeclineReasonGroupLocked       DeclinedTransactionSourceCardDeclineReason = "group_locked"
	DeclinedTransactionSourceCardDeclineReasonInsufficientFunds DeclinedTransactionSourceCardDeclineReason = "insufficient_funds"
	DeclinedTransactionSourceCardDeclineReasonBreachesLimit     DeclinedTransactionSourceCardDeclineReason = "breaches_limit"
	DeclinedTransactionSourceCardDeclineReasonWebhookDeclined   DeclinedTransactionSourceCardDeclineReason = "webhook_declined"
	DeclinedTransactionSourceCardDeclineReasonWebhookTimedOut   DeclinedTransactionSourceCardDeclineReason = "webhook_timed_out"
)

//
type DeclinedTransactionSourceCheckDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount int `json:"amount"`
	//
	AuxiliaryOnUs *string `json:"auxiliary_on_us"`
	// Why the check was declined.
	Reason DeclinedTransactionSourceCheckDeclineReason `json:"reason"`
}

func (r *DeclinedTransactionSourceCheckDecline) GetAuxiliaryOnUs() (AuxiliaryOnUs string) {
	if r != nil && r.AuxiliaryOnUs != nil {
		AuxiliaryOnUs = *r.AuxiliaryOnUs
	}
	return
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
)

//
type DeclinedTransactionSourceInboundRealTimePaymentsTransferDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the declined
	// transfer's currency. This will always be "USD" for a Real Time Payments
	// transfer.
	Currency DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency `json:"currency"`
	// Why the transfer was declined.
	Reason DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason `json:"reason"`
	// The name the sender of the transfer specified as the recipient of the transfer.
	CreditorName string `json:"creditor_name"`
	// The name provided by the sender of the transfer.
	DebtorName string `json:"debtor_name"`
	// The account number of the account that sent the transfer.
	DebtorAccountNumber string `json:"debtor_account_number"`
	// The routing number of the account that sent the transfer.
	DebtorRoutingNumber string `json:"debtor_routing_number"`
	// The Real Time Payments network identification of the declined transfer.
	TransactionIdentification string `json:"transaction_identification"`
	// Additional information included with the transfer.
	RemittanceInformation *string `json:"remittance_information"`
}

// Additional information included with the transfer.
func (r *DeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) GetRemittanceInformation() (RemittanceInformation string) {
	if r != nil && r.RemittanceInformation != nil {
		RemittanceInformation = *r.RemittanceInformation
	}
	return
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

//
type DeclinedTransactionSourceInternationalACHDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount int `json:"amount"`
	//
	ForeignExchangeIndicator string `json:"foreign_exchange_indicator"`
	//
	ForeignExchangeReferenceIndicator string `json:"foreign_exchange_reference_indicator"`
	//
	ForeignExchangeReference *string `json:"foreign_exchange_reference"`
	//
	DestinationCountryCode string `json:"destination_country_code"`
	//
	DestinationCurrencyCode string `json:"destination_currency_code"`
	//
	ForeignPaymentAmount int `json:"foreign_payment_amount"`
	//
	ForeignTraceNumber *string `json:"foreign_trace_number"`
	//
	InternationalTransactionTypeCode string `json:"international_transaction_type_code"`
	//
	OriginatingCurrencyCode string `json:"originating_currency_code"`
	//
	OriginatingDepositoryFinancialInstitutionName string `json:"originating_depository_financial_institution_name"`
	//
	OriginatingDepositoryFinancialInstitutionIDQualifier string `json:"originating_depository_financial_institution_id_qualifier"`
	//
	OriginatingDepositoryFinancialInstitutionID string `json:"originating_depository_financial_institution_id"`
	//
	OriginatingDepositoryFinancialInstitutionBranchCountry string `json:"originating_depository_financial_institution_branch_country"`
	//
	OriginatorCity string `json:"originator_city"`
	//
	OriginatorCompanyEntryDescription string `json:"originator_company_entry_description"`
	//
	OriginatorCountry string `json:"originator_country"`
	//
	OriginatorIdentification string `json:"originator_identification"`
	//
	OriginatorName string `json:"originator_name"`
	//
	OriginatorPostalCode *string `json:"originator_postal_code"`
	//
	OriginatorStreetAddress string `json:"originator_street_address"`
	//
	OriginatorStateOrProvince *string `json:"originator_state_or_province"`
	//
	PaymentRelatedInformation *string `json:"payment_related_information"`
	//
	PaymentRelatedInformation2 *string `json:"payment_related_information2"`
	//
	ReceiverIdentificationNumber *string `json:"receiver_identification_number"`
	//
	ReceiverStreetAddress string `json:"receiver_street_address"`
	//
	ReceiverCity string `json:"receiver_city"`
	//
	ReceiverStateOrProvince *string `json:"receiver_state_or_province"`
	//
	ReceiverCountry string `json:"receiver_country"`
	//
	ReceiverPostalCode *string `json:"receiver_postal_code"`
	//
	ReceivingCompanyOrIndividualName string `json:"receiving_company_or_individual_name"`
	//
	ReceivingDepositoryFinancialInstitutionName string `json:"receiving_depository_financial_institution_name"`
	//
	ReceivingDepositoryFinancialInstitutionIDQualifier string `json:"receiving_depository_financial_institution_id_qualifier"`
	//
	ReceivingDepositoryFinancialInstitutionID string `json:"receiving_depository_financial_institution_id"`
	//
	ReceivingDepositoryFinancialInstitutionCountry string `json:"receiving_depository_financial_institution_country"`
	//
	TraceNumber string `json:"trace_number"`
}

func (r *DeclinedTransactionSourceInternationalACHDecline) GetForeignExchangeReference() (ForeignExchangeReference string) {
	if r != nil && r.ForeignExchangeReference != nil {
		ForeignExchangeReference = *r.ForeignExchangeReference
	}
	return
}

func (r *DeclinedTransactionSourceInternationalACHDecline) GetForeignTraceNumber() (ForeignTraceNumber string) {
	if r != nil && r.ForeignTraceNumber != nil {
		ForeignTraceNumber = *r.ForeignTraceNumber
	}
	return
}

func (r *DeclinedTransactionSourceInternationalACHDecline) GetOriginatorPostalCode() (OriginatorPostalCode string) {
	if r != nil && r.OriginatorPostalCode != nil {
		OriginatorPostalCode = *r.OriginatorPostalCode
	}
	return
}

func (r *DeclinedTransactionSourceInternationalACHDecline) GetOriginatorStateOrProvince() (OriginatorStateOrProvince string) {
	if r != nil && r.OriginatorStateOrProvince != nil {
		OriginatorStateOrProvince = *r.OriginatorStateOrProvince
	}
	return
}

func (r *DeclinedTransactionSourceInternationalACHDecline) GetPaymentRelatedInformation() (PaymentRelatedInformation string) {
	if r != nil && r.PaymentRelatedInformation != nil {
		PaymentRelatedInformation = *r.PaymentRelatedInformation
	}
	return
}

func (r *DeclinedTransactionSourceInternationalACHDecline) GetPaymentRelatedInformation2() (PaymentRelatedInformation2 string) {
	if r != nil && r.PaymentRelatedInformation2 != nil {
		PaymentRelatedInformation2 = *r.PaymentRelatedInformation2
	}
	return
}

func (r *DeclinedTransactionSourceInternationalACHDecline) GetReceiverIdentificationNumber() (ReceiverIdentificationNumber string) {
	if r != nil && r.ReceiverIdentificationNumber != nil {
		ReceiverIdentificationNumber = *r.ReceiverIdentificationNumber
	}
	return
}

func (r *DeclinedTransactionSourceInternationalACHDecline) GetReceiverStateOrProvince() (ReceiverStateOrProvince string) {
	if r != nil && r.ReceiverStateOrProvince != nil {
		ReceiverStateOrProvince = *r.ReceiverStateOrProvince
	}
	return
}

func (r *DeclinedTransactionSourceInternationalACHDecline) GetReceiverPostalCode() (ReceiverPostalCode string) {
	if r != nil && r.ReceiverPostalCode != nil {
		ReceiverPostalCode = *r.ReceiverPostalCode
	}
	return
}

//
type DeclinedTransactionSourceCardRouteDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency DeclinedTransactionSourceCardRouteDeclineCurrency `json:"currency"`
	//
	MerchantAcceptorID string `json:"merchant_acceptor_id"`
	//
	MerchantCity *string `json:"merchant_city"`
	//
	MerchantCountry string `json:"merchant_country"`
	//
	MerchantDescriptor string `json:"merchant_descriptor"`
	//
	MerchantState *string `json:"merchant_state"`
	//
	MerchantCategoryCode *string `json:"merchant_category_code"`
}

func (r *DeclinedTransactionSourceCardRouteDecline) GetMerchantCity() (MerchantCity string) {
	if r != nil && r.MerchantCity != nil {
		MerchantCity = *r.MerchantCity
	}
	return
}

func (r *DeclinedTransactionSourceCardRouteDecline) GetMerchantState() (MerchantState string) {
	if r != nil && r.MerchantState != nil {
		MerchantState = *r.MerchantState
	}
	return
}

func (r *DeclinedTransactionSourceCardRouteDecline) GetMerchantCategoryCode() (MerchantCategoryCode string) {
	if r != nil && r.MerchantCategoryCode != nil {
		MerchantCategoryCode = *r.MerchantCategoryCode
	}
	return
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

type ListDeclinedTransactionsQuery struct {
	// Return the page of entries after this one.
	Cursor string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit int `query:"limit"`
	// Filter Declined Transactions to ones belonging to the specified Account.
	AccountID string `query:"account_id"`
	// Filter Declined Transactions to those belonging to the specified route.
	RouteID   string                                 `query:"route_id"`
	CreatedAt ListDeclinedTransactionsQueryCreatedAt `query:"created_at"`
}

type ListDeclinedTransactionsQueryCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After string `json:"after,omitempty"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before string `json:"before,omitempty"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter string `json:"on_or_after,omitempty"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore string `json:"on_or_before,omitempty"`
}

//
type DeclinedTransactionList struct {
	// The contents of the list.
	Data []DeclinedTransaction `json:"data"`
	// A pointer to a place in the list.
	NextCursor *string `json:"next_cursor"`
}

// A pointer to a place in the list.
func (r *DeclinedTransactionList) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

func (r *DeclinedTransactionService) Retrieve(ctx context.Context, declined_transaction_id string, opts ...*core.RequestOpts) (res *DeclinedTransaction, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/declined_transactions/%s", declined_transaction_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)

	return
}

type DeclinedTransactionsPage struct {
	*pagination.Page[DeclinedTransaction]
}

func (r *DeclinedTransactionsPage) DeclinedTransaction() *DeclinedTransaction {
	return r.Current()
}

func (r *DeclinedTransactionsPage) GetNextPage() (*DeclinedTransactionsPage, error) {
	if page, err := r.Page.GetNextPage(); err != nil {
		return nil, err
	} else {
		return &DeclinedTransactionsPage{page}, nil
	}
}

func (r *DeclinedTransactionService) List(ctx context.Context, query *ListDeclinedTransactionsQuery, opts ...*core.RequestOpts) (res *DeclinedTransactionsPage, err error) {
	page := &DeclinedTransactionsPage{
		Page: &pagination.Page[DeclinedTransaction]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/declined_transactions",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

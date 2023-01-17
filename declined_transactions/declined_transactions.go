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

type PreloadedDeclinedTransactionService struct {
	DeclinedTransactions *DeclinedTransactionService
}

func (r *PreloadedDeclinedTransactionService) Init(service *DeclinedTransactionService) {
	r.DeclinedTransactions = service
}

func NewPreloadedDeclinedTransactionService(service *DeclinedTransactionService) (r *PreloadedDeclinedTransactionService) {
	r = &PreloadedDeclinedTransactionService{}
	r.Init(service)
	return
}

//
type DeclinedTransaction struct {
	// The identifier for the Account the Declined Transaction belongs to.
	AccountID *string `json:"account_id"`
	// The Declined Transaction amount in the minor unit of its currency. For dollars,
	// for example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Declined
	// Transaction's currency. This will match the currency on the Declined
	// Transcation's Account.
	Currency *DeclinedTransactionCurrency `json:"currency"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the
	// Transaction occured.
	CreatedAt *string `json:"created_at"`
	// This is the description the vendor provides.
	Description *string `json:"description"`
	// The Declined Transaction identifier.
	ID *string `json:"id"`
	// The identifier for the route this Declined Transaction came through. Routes are
	// things like cards and ACH details.
	RouteID *string `json:"route_id"`
	// The type of the route this Declined Transaction came through.
	RouteType *string `json:"route_type"`
	// This is an object giving more details on the network-level event that caused the
	// Declined Transaction. For example, for a card transaction this lists the
	// merchant's industry and location. Note that for backwards compatibility reasons,
	// additional undocumented keys may appear in this object. These should be treated
	// as deprecated and will be removed in the future.
	Source *DeclinedTransactionSource `json:"source"`
	// A constant representing the object's type. For this resource it will always be
	// `declined_transaction`.
	Type *DeclinedTransactionType `json:"type"`
}

// The identifier for the Account the Declined Transaction belongs to.
func (r *DeclinedTransaction) GetAccountID() (AccountID string) {
	if r != nil && r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

// The Declined Transaction amount in the minor unit of its currency. For dollars,
// for example, this is cents.
func (r *DeclinedTransaction) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Declined
// Transaction's currency. This will match the currency on the Declined
// Transcation's Account.
func (r *DeclinedTransaction) GetCurrency() (Currency DeclinedTransactionCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the
// Transaction occured.
func (r *DeclinedTransaction) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// This is the description the vendor provides.
func (r *DeclinedTransaction) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The Declined Transaction identifier.
func (r *DeclinedTransaction) GetID() (ID string) {
	if r != nil && r.ID != nil {
		ID = *r.ID
	}
	return
}

// The identifier for the route this Declined Transaction came through. Routes are
// things like cards and ACH details.
func (r *DeclinedTransaction) GetRouteID() (RouteID string) {
	if r != nil && r.RouteID != nil {
		RouteID = *r.RouteID
	}
	return
}

// The type of the route this Declined Transaction came through.
func (r *DeclinedTransaction) GetRouteType() (RouteType string) {
	if r != nil && r.RouteType != nil {
		RouteType = *r.RouteType
	}
	return
}

// This is an object giving more details on the network-level event that caused the
// Declined Transaction. For example, for a card transaction this lists the
// merchant's industry and location. Note that for backwards compatibility reasons,
// additional undocumented keys may appear in this object. These should be treated
// as deprecated and will be removed in the future.
func (r *DeclinedTransaction) GetSource() (Source DeclinedTransactionSource) {
	if r != nil && r.Source != nil {
		Source = *r.Source
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `declined_transaction`.
func (r *DeclinedTransaction) GetType() (Type DeclinedTransactionType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
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
	Category *DeclinedTransactionSourceCategory `json:"category"`
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

// The type of decline that took place. We may add additional possible values for
// this enum over time; your application should be able to handle such additions
// gracefully.
func (r *DeclinedTransactionSource) GetCategory() (Category DeclinedTransactionSourceCategory) {
	if r != nil && r.Category != nil {
		Category = *r.Category
	}
	return
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
	Amount *int `json:"amount"`
	//
	OriginatorCompanyName *string `json:"originator_company_name"`
	//
	OriginatorCompanyDescriptiveDate *string `json:"originator_company_descriptive_date"`
	//
	OriginatorCompanyDiscretionaryData *string `json:"originator_company_discretionary_data"`
	//
	OriginatorCompanyID *string `json:"originator_company_id"`
	// Why the ACH transfer was declined.
	Reason *DeclinedTransactionSourceACHDeclineReason `json:"reason"`
	//
	ReceiverIDNumber *string `json:"receiver_id_number"`
	//
	ReceiverName *string `json:"receiver_name"`
	//
	TraceNumber *string `json:"trace_number"`
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *DeclinedTransactionSourceACHDecline) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *DeclinedTransactionSourceACHDecline) GetOriginatorCompanyName() (OriginatorCompanyName string) {
	if r != nil && r.OriginatorCompanyName != nil {
		OriginatorCompanyName = *r.OriginatorCompanyName
	}
	return
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

func (r *DeclinedTransactionSourceACHDecline) GetOriginatorCompanyID() (OriginatorCompanyID string) {
	if r != nil && r.OriginatorCompanyID != nil {
		OriginatorCompanyID = *r.OriginatorCompanyID
	}
	return
}

// Why the ACH transfer was declined.
func (r *DeclinedTransactionSourceACHDecline) GetReason() (Reason DeclinedTransactionSourceACHDeclineReason) {
	if r != nil && r.Reason != nil {
		Reason = *r.Reason
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

func (r *DeclinedTransactionSourceACHDecline) GetTraceNumber() (TraceNumber string) {
	if r != nil && r.TraceNumber != nil {
		TraceNumber = *r.TraceNumber
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
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency *DeclinedTransactionSourceCardDeclineCurrency `json:"currency"`
	//
	MerchantAcceptorID *string `json:"merchant_acceptor_id"`
	//
	MerchantCity *string `json:"merchant_city"`
	//
	MerchantCountry *string `json:"merchant_country"`
	//
	MerchantDescriptor *string `json:"merchant_descriptor"`
	//
	MerchantState *string `json:"merchant_state"`
	//
	MerchantCategoryCode *string `json:"merchant_category_code"`
	// Why the transaction was declined.
	Reason *DeclinedTransactionSourceCardDeclineReason `json:"reason"`
	// The identifier of the Real-Time Decision sent to approve or decline this
	// transaction.
	RealTimeDecisionID *string `json:"real_time_decision_id"`
	// If the authorization was attempted using a Digital Wallet Token (such as an
	// Apple Pay purchase), the identifier of the token that was used.
	DigitalWalletTokenID *string `json:"digital_wallet_token_id"`
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *DeclinedTransactionSourceCardDecline) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
// account currency.
func (r *DeclinedTransactionSourceCardDecline) GetCurrency() (Currency DeclinedTransactionSourceCardDeclineCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

func (r *DeclinedTransactionSourceCardDecline) GetMerchantAcceptorID() (MerchantAcceptorID string) {
	if r != nil && r.MerchantAcceptorID != nil {
		MerchantAcceptorID = *r.MerchantAcceptorID
	}
	return
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

func (r *DeclinedTransactionSourceCardDecline) GetMerchantDescriptor() (MerchantDescriptor string) {
	if r != nil && r.MerchantDescriptor != nil {
		MerchantDescriptor = *r.MerchantDescriptor
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

// Why the transaction was declined.
func (r *DeclinedTransactionSourceCardDecline) GetReason() (Reason DeclinedTransactionSourceCardDeclineReason) {
	if r != nil && r.Reason != nil {
		Reason = *r.Reason
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
	Amount *int `json:"amount"`
	//
	AuxiliaryOnUs *string `json:"auxiliary_on_us"`
	// Why the check was declined.
	Reason *DeclinedTransactionSourceCheckDeclineReason `json:"reason"`
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *DeclinedTransactionSourceCheckDecline) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *DeclinedTransactionSourceCheckDecline) GetAuxiliaryOnUs() (AuxiliaryOnUs string) {
	if r != nil && r.AuxiliaryOnUs != nil {
		AuxiliaryOnUs = *r.AuxiliaryOnUs
	}
	return
}

// Why the check was declined.
func (r *DeclinedTransactionSourceCheckDecline) GetReason() (Reason DeclinedTransactionSourceCheckDeclineReason) {
	if r != nil && r.Reason != nil {
		Reason = *r.Reason
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
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the declined
	// transfer's currency. This will always be "USD" for a Real Time Payments
	// transfer.
	Currency *DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency `json:"currency"`
	// Why the transfer was declined.
	Reason *DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason `json:"reason"`
	// The name the sender of the transfer specified as the recipient of the transfer.
	CreditorName *string `json:"creditor_name"`
	// The name provided by the sender of the transfer.
	DebtorName *string `json:"debtor_name"`
	// The account number of the account that sent the transfer.
	DebtorAccountNumber *string `json:"debtor_account_number"`
	// The routing number of the account that sent the transfer.
	DebtorRoutingNumber *string `json:"debtor_routing_number"`
	// The Real Time Payments network identification of the declined transfer.
	TransactionIdentification *string `json:"transaction_identification"`
	// Additional information included with the transfer.
	RemittanceInformation *string `json:"remittance_information"`
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *DeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the declined
// transfer's currency. This will always be "USD" for a Real Time Payments
// transfer.
func (r *DeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) GetCurrency() (Currency DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// Why the transfer was declined.
func (r *DeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) GetReason() (Reason DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason) {
	if r != nil && r.Reason != nil {
		Reason = *r.Reason
	}
	return
}

// The name the sender of the transfer specified as the recipient of the transfer.
func (r *DeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) GetCreditorName() (CreditorName string) {
	if r != nil && r.CreditorName != nil {
		CreditorName = *r.CreditorName
	}
	return
}

// The name provided by the sender of the transfer.
func (r *DeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) GetDebtorName() (DebtorName string) {
	if r != nil && r.DebtorName != nil {
		DebtorName = *r.DebtorName
	}
	return
}

// The account number of the account that sent the transfer.
func (r *DeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) GetDebtorAccountNumber() (DebtorAccountNumber string) {
	if r != nil && r.DebtorAccountNumber != nil {
		DebtorAccountNumber = *r.DebtorAccountNumber
	}
	return
}

// The routing number of the account that sent the transfer.
func (r *DeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) GetDebtorRoutingNumber() (DebtorRoutingNumber string) {
	if r != nil && r.DebtorRoutingNumber != nil {
		DebtorRoutingNumber = *r.DebtorRoutingNumber
	}
	return
}

// The Real Time Payments network identification of the declined transfer.
func (r *DeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) GetTransactionIdentification() (TransactionIdentification string) {
	if r != nil && r.TransactionIdentification != nil {
		TransactionIdentification = *r.TransactionIdentification
	}
	return
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
	Amount *int `json:"amount"`
	//
	ForeignExchangeIndicator *string `json:"foreign_exchange_indicator"`
	//
	ForeignExchangeReferenceIndicator *string `json:"foreign_exchange_reference_indicator"`
	//
	ForeignExchangeReference *string `json:"foreign_exchange_reference"`
	//
	DestinationCountryCode *string `json:"destination_country_code"`
	//
	DestinationCurrencyCode *string `json:"destination_currency_code"`
	//
	ForeignPaymentAmount *int `json:"foreign_payment_amount"`
	//
	ForeignTraceNumber *string `json:"foreign_trace_number"`
	//
	InternationalTransactionTypeCode *string `json:"international_transaction_type_code"`
	//
	OriginatingCurrencyCode *string `json:"originating_currency_code"`
	//
	OriginatingDepositoryFinancialInstitutionName *string `json:"originating_depository_financial_institution_name"`
	//
	OriginatingDepositoryFinancialInstitutionIDQualifier *string `json:"originating_depository_financial_institution_id_qualifier"`
	//
	OriginatingDepositoryFinancialInstitutionID *string `json:"originating_depository_financial_institution_id"`
	//
	OriginatingDepositoryFinancialInstitutionBranchCountry *string `json:"originating_depository_financial_institution_branch_country"`
	//
	OriginatorCity *string `json:"originator_city"`
	//
	OriginatorCompanyEntryDescription *string `json:"originator_company_entry_description"`
	//
	OriginatorCountry *string `json:"originator_country"`
	//
	OriginatorIdentification *string `json:"originator_identification"`
	//
	OriginatorName *string `json:"originator_name"`
	//
	OriginatorPostalCode *string `json:"originator_postal_code"`
	//
	OriginatorStreetAddress *string `json:"originator_street_address"`
	//
	OriginatorStateOrProvince *string `json:"originator_state_or_province"`
	//
	PaymentRelatedInformation *string `json:"payment_related_information"`
	//
	PaymentRelatedInformation2 *string `json:"payment_related_information2"`
	//
	ReceiverIdentificationNumber *string `json:"receiver_identification_number"`
	//
	ReceiverStreetAddress *string `json:"receiver_street_address"`
	//
	ReceiverCity *string `json:"receiver_city"`
	//
	ReceiverStateOrProvince *string `json:"receiver_state_or_province"`
	//
	ReceiverCountry *string `json:"receiver_country"`
	//
	ReceiverPostalCode *string `json:"receiver_postal_code"`
	//
	ReceivingCompanyOrIndividualName *string `json:"receiving_company_or_individual_name"`
	//
	ReceivingDepositoryFinancialInstitutionName *string `json:"receiving_depository_financial_institution_name"`
	//
	ReceivingDepositoryFinancialInstitutionIDQualifier *string `json:"receiving_depository_financial_institution_id_qualifier"`
	//
	ReceivingDepositoryFinancialInstitutionID *string `json:"receiving_depository_financial_institution_id"`
	//
	ReceivingDepositoryFinancialInstitutionCountry *string `json:"receiving_depository_financial_institution_country"`
	//
	TraceNumber *string `json:"trace_number"`
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *DeclinedTransactionSourceInternationalACHDecline) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *DeclinedTransactionSourceInternationalACHDecline) GetForeignExchangeIndicator() (ForeignExchangeIndicator string) {
	if r != nil && r.ForeignExchangeIndicator != nil {
		ForeignExchangeIndicator = *r.ForeignExchangeIndicator
	}
	return
}

func (r *DeclinedTransactionSourceInternationalACHDecline) GetForeignExchangeReferenceIndicator() (ForeignExchangeReferenceIndicator string) {
	if r != nil && r.ForeignExchangeReferenceIndicator != nil {
		ForeignExchangeReferenceIndicator = *r.ForeignExchangeReferenceIndicator
	}
	return
}

func (r *DeclinedTransactionSourceInternationalACHDecline) GetForeignExchangeReference() (ForeignExchangeReference string) {
	if r != nil && r.ForeignExchangeReference != nil {
		ForeignExchangeReference = *r.ForeignExchangeReference
	}
	return
}

func (r *DeclinedTransactionSourceInternationalACHDecline) GetDestinationCountryCode() (DestinationCountryCode string) {
	if r != nil && r.DestinationCountryCode != nil {
		DestinationCountryCode = *r.DestinationCountryCode
	}
	return
}

func (r *DeclinedTransactionSourceInternationalACHDecline) GetDestinationCurrencyCode() (DestinationCurrencyCode string) {
	if r != nil && r.DestinationCurrencyCode != nil {
		DestinationCurrencyCode = *r.DestinationCurrencyCode
	}
	return
}

func (r *DeclinedTransactionSourceInternationalACHDecline) GetForeignPaymentAmount() (ForeignPaymentAmount int) {
	if r != nil && r.ForeignPaymentAmount != nil {
		ForeignPaymentAmount = *r.ForeignPaymentAmount
	}
	return
}

func (r *DeclinedTransactionSourceInternationalACHDecline) GetForeignTraceNumber() (ForeignTraceNumber string) {
	if r != nil && r.ForeignTraceNumber != nil {
		ForeignTraceNumber = *r.ForeignTraceNumber
	}
	return
}

func (r *DeclinedTransactionSourceInternationalACHDecline) GetInternationalTransactionTypeCode() (InternationalTransactionTypeCode string) {
	if r != nil && r.InternationalTransactionTypeCode != nil {
		InternationalTransactionTypeCode = *r.InternationalTransactionTypeCode
	}
	return
}

func (r *DeclinedTransactionSourceInternationalACHDecline) GetOriginatingCurrencyCode() (OriginatingCurrencyCode string) {
	if r != nil && r.OriginatingCurrencyCode != nil {
		OriginatingCurrencyCode = *r.OriginatingCurrencyCode
	}
	return
}

func (r *DeclinedTransactionSourceInternationalACHDecline) GetOriginatingDepositoryFinancialInstitutionName() (OriginatingDepositoryFinancialInstitutionName string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionName != nil {
		OriginatingDepositoryFinancialInstitutionName = *r.OriginatingDepositoryFinancialInstitutionName
	}
	return
}

func (r *DeclinedTransactionSourceInternationalACHDecline) GetOriginatingDepositoryFinancialInstitutionIDQualifier() (OriginatingDepositoryFinancialInstitutionIDQualifier string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionIDQualifier != nil {
		OriginatingDepositoryFinancialInstitutionIDQualifier = *r.OriginatingDepositoryFinancialInstitutionIDQualifier
	}
	return
}

func (r *DeclinedTransactionSourceInternationalACHDecline) GetOriginatingDepositoryFinancialInstitutionID() (OriginatingDepositoryFinancialInstitutionID string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionID != nil {
		OriginatingDepositoryFinancialInstitutionID = *r.OriginatingDepositoryFinancialInstitutionID
	}
	return
}

func (r *DeclinedTransactionSourceInternationalACHDecline) GetOriginatingDepositoryFinancialInstitutionBranchCountry() (OriginatingDepositoryFinancialInstitutionBranchCountry string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionBranchCountry != nil {
		OriginatingDepositoryFinancialInstitutionBranchCountry = *r.OriginatingDepositoryFinancialInstitutionBranchCountry
	}
	return
}

func (r *DeclinedTransactionSourceInternationalACHDecline) GetOriginatorCity() (OriginatorCity string) {
	if r != nil && r.OriginatorCity != nil {
		OriginatorCity = *r.OriginatorCity
	}
	return
}

func (r *DeclinedTransactionSourceInternationalACHDecline) GetOriginatorCompanyEntryDescription() (OriginatorCompanyEntryDescription string) {
	if r != nil && r.OriginatorCompanyEntryDescription != nil {
		OriginatorCompanyEntryDescription = *r.OriginatorCompanyEntryDescription
	}
	return
}

func (r *DeclinedTransactionSourceInternationalACHDecline) GetOriginatorCountry() (OriginatorCountry string) {
	if r != nil && r.OriginatorCountry != nil {
		OriginatorCountry = *r.OriginatorCountry
	}
	return
}

func (r *DeclinedTransactionSourceInternationalACHDecline) GetOriginatorIdentification() (OriginatorIdentification string) {
	if r != nil && r.OriginatorIdentification != nil {
		OriginatorIdentification = *r.OriginatorIdentification
	}
	return
}

func (r *DeclinedTransactionSourceInternationalACHDecline) GetOriginatorName() (OriginatorName string) {
	if r != nil && r.OriginatorName != nil {
		OriginatorName = *r.OriginatorName
	}
	return
}

func (r *DeclinedTransactionSourceInternationalACHDecline) GetOriginatorPostalCode() (OriginatorPostalCode string) {
	if r != nil && r.OriginatorPostalCode != nil {
		OriginatorPostalCode = *r.OriginatorPostalCode
	}
	return
}

func (r *DeclinedTransactionSourceInternationalACHDecline) GetOriginatorStreetAddress() (OriginatorStreetAddress string) {
	if r != nil && r.OriginatorStreetAddress != nil {
		OriginatorStreetAddress = *r.OriginatorStreetAddress
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

func (r *DeclinedTransactionSourceInternationalACHDecline) GetReceiverStreetAddress() (ReceiverStreetAddress string) {
	if r != nil && r.ReceiverStreetAddress != nil {
		ReceiverStreetAddress = *r.ReceiverStreetAddress
	}
	return
}

func (r *DeclinedTransactionSourceInternationalACHDecline) GetReceiverCity() (ReceiverCity string) {
	if r != nil && r.ReceiverCity != nil {
		ReceiverCity = *r.ReceiverCity
	}
	return
}

func (r *DeclinedTransactionSourceInternationalACHDecline) GetReceiverStateOrProvince() (ReceiverStateOrProvince string) {
	if r != nil && r.ReceiverStateOrProvince != nil {
		ReceiverStateOrProvince = *r.ReceiverStateOrProvince
	}
	return
}

func (r *DeclinedTransactionSourceInternationalACHDecline) GetReceiverCountry() (ReceiverCountry string) {
	if r != nil && r.ReceiverCountry != nil {
		ReceiverCountry = *r.ReceiverCountry
	}
	return
}

func (r *DeclinedTransactionSourceInternationalACHDecline) GetReceiverPostalCode() (ReceiverPostalCode string) {
	if r != nil && r.ReceiverPostalCode != nil {
		ReceiverPostalCode = *r.ReceiverPostalCode
	}
	return
}

func (r *DeclinedTransactionSourceInternationalACHDecline) GetReceivingCompanyOrIndividualName() (ReceivingCompanyOrIndividualName string) {
	if r != nil && r.ReceivingCompanyOrIndividualName != nil {
		ReceivingCompanyOrIndividualName = *r.ReceivingCompanyOrIndividualName
	}
	return
}

func (r *DeclinedTransactionSourceInternationalACHDecline) GetReceivingDepositoryFinancialInstitutionName() (ReceivingDepositoryFinancialInstitutionName string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionName != nil {
		ReceivingDepositoryFinancialInstitutionName = *r.ReceivingDepositoryFinancialInstitutionName
	}
	return
}

func (r *DeclinedTransactionSourceInternationalACHDecline) GetReceivingDepositoryFinancialInstitutionIDQualifier() (ReceivingDepositoryFinancialInstitutionIDQualifier string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionIDQualifier != nil {
		ReceivingDepositoryFinancialInstitutionIDQualifier = *r.ReceivingDepositoryFinancialInstitutionIDQualifier
	}
	return
}

func (r *DeclinedTransactionSourceInternationalACHDecline) GetReceivingDepositoryFinancialInstitutionID() (ReceivingDepositoryFinancialInstitutionID string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionID != nil {
		ReceivingDepositoryFinancialInstitutionID = *r.ReceivingDepositoryFinancialInstitutionID
	}
	return
}

func (r *DeclinedTransactionSourceInternationalACHDecline) GetReceivingDepositoryFinancialInstitutionCountry() (ReceivingDepositoryFinancialInstitutionCountry string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionCountry != nil {
		ReceivingDepositoryFinancialInstitutionCountry = *r.ReceivingDepositoryFinancialInstitutionCountry
	}
	return
}

func (r *DeclinedTransactionSourceInternationalACHDecline) GetTraceNumber() (TraceNumber string) {
	if r != nil && r.TraceNumber != nil {
		TraceNumber = *r.TraceNumber
	}
	return
}

//
type DeclinedTransactionSourceCardRouteDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency *DeclinedTransactionSourceCardRouteDeclineCurrency `json:"currency"`
	//
	MerchantAcceptorID *string `json:"merchant_acceptor_id"`
	//
	MerchantCity *string `json:"merchant_city"`
	//
	MerchantCountry *string `json:"merchant_country"`
	//
	MerchantDescriptor *string `json:"merchant_descriptor"`
	//
	MerchantState *string `json:"merchant_state"`
	//
	MerchantCategoryCode *string `json:"merchant_category_code"`
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *DeclinedTransactionSourceCardRouteDecline) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
// account currency.
func (r *DeclinedTransactionSourceCardRouteDecline) GetCurrency() (Currency DeclinedTransactionSourceCardRouteDeclineCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

func (r *DeclinedTransactionSourceCardRouteDecline) GetMerchantAcceptorID() (MerchantAcceptorID string) {
	if r != nil && r.MerchantAcceptorID != nil {
		MerchantAcceptorID = *r.MerchantAcceptorID
	}
	return
}

func (r *DeclinedTransactionSourceCardRouteDecline) GetMerchantCity() (MerchantCity string) {
	if r != nil && r.MerchantCity != nil {
		MerchantCity = *r.MerchantCity
	}
	return
}

func (r *DeclinedTransactionSourceCardRouteDecline) GetMerchantCountry() (MerchantCountry string) {
	if r != nil && r.MerchantCountry != nil {
		MerchantCountry = *r.MerchantCountry
	}
	return
}

func (r *DeclinedTransactionSourceCardRouteDecline) GetMerchantDescriptor() (MerchantDescriptor string) {
	if r != nil && r.MerchantDescriptor != nil {
		MerchantDescriptor = *r.MerchantDescriptor
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
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int `query:"limit"`
	// Filter Declined Transactions to ones belonging to the specified Account.
	AccountID *string `query:"account_id"`
	// Filter Declined Transactions to those belonging to the specified route.
	RouteID   *string                                 `query:"route_id"`
	CreatedAt *ListDeclinedTransactionsQueryCreatedAt `query:"created_at"`
}

// Return the page of entries after this one.
func (r *ListDeclinedTransactionsQuery) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *ListDeclinedTransactionsQuery) GetLimit() (Limit int) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// Filter Declined Transactions to ones belonging to the specified Account.
func (r *ListDeclinedTransactionsQuery) GetAccountID() (AccountID string) {
	if r != nil && r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

// Filter Declined Transactions to those belonging to the specified route.
func (r *ListDeclinedTransactionsQuery) GetRouteID() (RouteID string) {
	if r != nil && r.RouteID != nil {
		RouteID = *r.RouteID
	}
	return
}

func (r *ListDeclinedTransactionsQuery) GetCreatedAt() (CreatedAt ListDeclinedTransactionsQueryCreatedAt) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

type ListDeclinedTransactionsQueryCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After *string `json:"after,omitempty"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before *string `json:"before,omitempty"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter *string `json:"on_or_after,omitempty"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore *string `json:"on_or_before,omitempty"`
}

// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r *ListDeclinedTransactionsQueryCreatedAt) GetAfter() (After string) {
	if r != nil && r.After != nil {
		After = *r.After
	}
	return
}

// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r *ListDeclinedTransactionsQueryCreatedAt) GetBefore() (Before string) {
	if r != nil && r.Before != nil {
		Before = *r.Before
	}
	return
}

// Return results on or after this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r *ListDeclinedTransactionsQueryCreatedAt) GetOnOrAfter() (OnOrAfter string) {
	if r != nil && r.OnOrAfter != nil {
		OnOrAfter = *r.OnOrAfter
	}
	return
}

// Return results on or before this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r *ListDeclinedTransactionsQueryCreatedAt) GetOnOrBefore() (OnOrBefore string) {
	if r != nil && r.OnOrBefore != nil {
		OnOrBefore = *r.OnOrBefore
	}
	return
}

//
type DeclinedTransactionList struct {
	// The contents of the list.
	Data *[]DeclinedTransaction `json:"data"`
	// A pointer to a place in the list.
	NextCursor *string `json:"next_cursor"`
}

// The contents of the list.
func (r *DeclinedTransactionList) GetData() (Data []DeclinedTransaction) {
	if r != nil && r.Data != nil {
		Data = *r.Data
	}
	return
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

func (r *PreloadedDeclinedTransactionService) Retrieve(ctx context.Context, declined_transaction_id string, opts ...*core.RequestOpts) (res *DeclinedTransaction, err error) {
	err = r.DeclinedTransactions.get(
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

func (r *PreloadedDeclinedTransactionService) List(ctx context.Context, query *ListDeclinedTransactionsQuery, opts ...*core.RequestOpts) (res *DeclinedTransactionsPage, err error) {
	page := &DeclinedTransactionsPage{
		Page: &pagination.Page[DeclinedTransaction]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/declined_transactions",
			},
			Requester: r.DeclinedTransactions.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

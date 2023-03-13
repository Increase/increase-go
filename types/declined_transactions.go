package types

import (
	"fmt"
	"net/url"

	"github.com/increase/increase-go/core"
	"github.com/increase/increase-go/core/pjson"
	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/pagination"
)

type DeclinedTransaction struct {
	// The identifier for the Account the Declined Transaction belongs to.
	AccountID *string `pjson:"account_id"`
	// The Declined Transaction amount in the minor unit of its currency. For dollars,
	// for example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Declined
	// Transaction's currency. This will match the currency on the Declined
	// Transcation's Account.
	Currency *DeclinedTransactionCurrency `pjson:"currency"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the
	// Transaction occured.
	CreatedAt *string `pjson:"created_at"`
	// This is the description the vendor provides.
	Description *string `pjson:"description"`
	// The Declined Transaction identifier.
	ID *string `pjson:"id"`
	// The identifier for the route this Declined Transaction came through. Routes are
	// things like cards and ACH details.
	RouteID *string `pjson:"route_id"`
	// The type of the route this Declined Transaction came through.
	RouteType *string `pjson:"route_type"`
	// This is an object giving more details on the network-level event that caused the
	// Declined Transaction. For example, for a card transaction this lists the
	// merchant's industry and location. Note that for backwards compatibility reasons,
	// additional undocumented keys may appear in this object. These should be treated
	// as deprecated and will be removed in the future.
	Source *DeclinedTransactionSource `pjson:"source"`
	// A constant representing the object's type. For this resource it will always be
	// `declined_transaction`.
	Type       *DeclinedTransactionType `pjson:"type"`
	jsonFields map[string]interface{}   `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into DeclinedTransaction using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *DeclinedTransaction) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes DeclinedTransaction into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *DeclinedTransaction) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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
func (r *DeclinedTransaction) GetAmount() (Amount int64) {
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

func (r DeclinedTransaction) String() (result string) {
	return fmt.Sprintf("&DeclinedTransaction{AccountID:%s Amount:%s Currency:%s CreatedAt:%s Description:%s ID:%s RouteID:%s RouteType:%s Source:%s Type:%s}", core.FmtP(r.AccountID), core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.CreatedAt), core.FmtP(r.Description), core.FmtP(r.ID), core.FmtP(r.RouteID), core.FmtP(r.RouteType), r.Source, core.FmtP(r.Type))
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

type DeclinedTransactionSource struct {
	// The type of decline that took place. We may add additional possible values for
	// this enum over time; your application should be able to handle such additions
	// gracefully.
	Category *DeclinedTransactionSourceCategory `pjson:"category"`
	// A ACH Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `ach_decline`.
	ACHDecline *DeclinedTransactionSourceACHDecline `pjson:"ach_decline"`
	// A Card Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_decline`.
	CardDecline *DeclinedTransactionSourceCardDecline `pjson:"card_decline"`
	// A Check Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `check_decline`.
	CheckDecline *DeclinedTransactionSourceCheckDecline `pjson:"check_decline"`
	// A Inbound Real Time Payments Transfer Decline object. This field will be present
	// in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_decline`.
	InboundRealTimePaymentsTransferDecline *DeclinedTransactionSourceInboundRealTimePaymentsTransferDecline `pjson:"inbound_real_time_payments_transfer_decline"`
	// A International ACH Decline object. This field will be present in the JSON
	// response if and only if `category` is equal to `international_ach_decline`.
	InternationalACHDecline *DeclinedTransactionSourceInternationalACHDecline `pjson:"international_ach_decline"`
	// A Deprecated Card Decline object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_route_decline`.
	CardRouteDecline *DeclinedTransactionSourceCardRouteDecline `pjson:"card_route_decline"`
	jsonFields       map[string]interface{}                     `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into DeclinedTransactionSource
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *DeclinedTransactionSource) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes DeclinedTransactionSource into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *DeclinedTransactionSource) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
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

func (r DeclinedTransactionSource) String() (result string) {
	return fmt.Sprintf("&DeclinedTransactionSource{Category:%s ACHDecline:%s CardDecline:%s CheckDecline:%s InboundRealTimePaymentsTransferDecline:%s InternationalACHDecline:%s CardRouteDecline:%s}", core.FmtP(r.Category), r.ACHDecline, r.CardDecline, r.CheckDecline, r.InboundRealTimePaymentsTransferDecline, r.InternationalACHDecline, r.CardRouteDecline)
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
	Amount                             *int64  `pjson:"amount"`
	OriginatorCompanyName              *string `pjson:"originator_company_name"`
	OriginatorCompanyDescriptiveDate   *string `pjson:"originator_company_descriptive_date"`
	OriginatorCompanyDiscretionaryData *string `pjson:"originator_company_discretionary_data"`
	OriginatorCompanyID                *string `pjson:"originator_company_id"`
	// Why the ACH transfer was declined.
	Reason           *DeclinedTransactionSourceACHDeclineReason `pjson:"reason"`
	ReceiverIDNumber *string                                    `pjson:"receiver_id_number"`
	ReceiverName     *string                                    `pjson:"receiver_name"`
	TraceNumber      *string                                    `pjson:"trace_number"`
	jsonFields       map[string]interface{}                     `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// DeclinedTransactionSourceACHDecline using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *DeclinedTransactionSourceACHDecline) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes DeclinedTransactionSourceACHDecline into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *DeclinedTransactionSourceACHDecline) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *DeclinedTransactionSourceACHDecline) GetAmount() (Amount int64) {
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

func (r DeclinedTransactionSourceACHDecline) String() (result string) {
	return fmt.Sprintf("&DeclinedTransactionSourceACHDecline{Amount:%s OriginatorCompanyName:%s OriginatorCompanyDescriptiveDate:%s OriginatorCompanyDiscretionaryData:%s OriginatorCompanyID:%s Reason:%s ReceiverIDNumber:%s ReceiverName:%s TraceNumber:%s}", core.FmtP(r.Amount), core.FmtP(r.OriginatorCompanyName), core.FmtP(r.OriginatorCompanyDescriptiveDate), core.FmtP(r.OriginatorCompanyDiscretionaryData), core.FmtP(r.OriginatorCompanyID), core.FmtP(r.Reason), core.FmtP(r.ReceiverIDNumber), core.FmtP(r.ReceiverName), core.FmtP(r.TraceNumber))
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
	MerchantAcceptorID *string `pjson:"merchant_acceptor_id"`
	// The merchant descriptor of the merchant the card is transacting with.
	MerchantDescriptor *string `pjson:"merchant_descriptor"`
	// The Merchant Category Code (commonly abbreviated as MCC) of the merchant the
	// card is transacting with.
	MerchantCategoryCode *string `pjson:"merchant_category_code"`
	// The city the merchant resides in.
	MerchantCity *string `pjson:"merchant_city"`
	// The country the merchant resides in.
	MerchantCountry *string `pjson:"merchant_country"`
	// The payment network used to process this card authorization
	Network *DeclinedTransactionSourceCardDeclineNetwork `pjson:"network"`
	// Fields specific to the `network`
	NetworkDetails *DeclinedTransactionSourceCardDeclineNetworkDetails `pjson:"network_details"`
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency *DeclinedTransactionSourceCardDeclineCurrency `pjson:"currency"`
	// Why the transaction was declined.
	Reason *DeclinedTransactionSourceCardDeclineReason `pjson:"reason"`
	// The state the merchant resides in.
	MerchantState *string `pjson:"merchant_state"`
	// The identifier of the Real-Time Decision sent to approve or decline this
	// transaction.
	RealTimeDecisionID *string `pjson:"real_time_decision_id"`
	// If the authorization was attempted using a Digital Wallet Token (such as an
	// Apple Pay purchase), the identifier of the token that was used.
	DigitalWalletTokenID *string                `pjson:"digital_wallet_token_id"`
	jsonFields           map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// DeclinedTransactionSourceCardDecline using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *DeclinedTransactionSourceCardDecline) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes DeclinedTransactionSourceCardDecline into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *DeclinedTransactionSourceCardDecline) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The merchant identifier (commonly abbreviated as MID) of the merchant the card
// is transacting with.
func (r *DeclinedTransactionSourceCardDecline) GetMerchantAcceptorID() (MerchantAcceptorID string) {
	if r != nil && r.MerchantAcceptorID != nil {
		MerchantAcceptorID = *r.MerchantAcceptorID
	}
	return
}

// The merchant descriptor of the merchant the card is transacting with.
func (r *DeclinedTransactionSourceCardDecline) GetMerchantDescriptor() (MerchantDescriptor string) {
	if r != nil && r.MerchantDescriptor != nil {
		MerchantDescriptor = *r.MerchantDescriptor
	}
	return
}

// The Merchant Category Code (commonly abbreviated as MCC) of the merchant the
// card is transacting with.
func (r *DeclinedTransactionSourceCardDecline) GetMerchantCategoryCode() (MerchantCategoryCode string) {
	if r != nil && r.MerchantCategoryCode != nil {
		MerchantCategoryCode = *r.MerchantCategoryCode
	}
	return
}

// The city the merchant resides in.
func (r *DeclinedTransactionSourceCardDecline) GetMerchantCity() (MerchantCity string) {
	if r != nil && r.MerchantCity != nil {
		MerchantCity = *r.MerchantCity
	}
	return
}

// The country the merchant resides in.
func (r *DeclinedTransactionSourceCardDecline) GetMerchantCountry() (MerchantCountry string) {
	if r != nil && r.MerchantCountry != nil {
		MerchantCountry = *r.MerchantCountry
	}
	return
}

// The payment network used to process this card authorization
func (r *DeclinedTransactionSourceCardDecline) GetNetwork() (Network DeclinedTransactionSourceCardDeclineNetwork) {
	if r != nil && r.Network != nil {
		Network = *r.Network
	}
	return
}

// Fields specific to the `network`
func (r *DeclinedTransactionSourceCardDecline) GetNetworkDetails() (NetworkDetails DeclinedTransactionSourceCardDeclineNetworkDetails) {
	if r != nil && r.NetworkDetails != nil {
		NetworkDetails = *r.NetworkDetails
	}
	return
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *DeclinedTransactionSourceCardDecline) GetAmount() (Amount int64) {
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

// Why the transaction was declined.
func (r *DeclinedTransactionSourceCardDecline) GetReason() (Reason DeclinedTransactionSourceCardDeclineReason) {
	if r != nil && r.Reason != nil {
		Reason = *r.Reason
	}
	return
}

// The state the merchant resides in.
func (r *DeclinedTransactionSourceCardDecline) GetMerchantState() (MerchantState string) {
	if r != nil && r.MerchantState != nil {
		MerchantState = *r.MerchantState
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

func (r DeclinedTransactionSourceCardDecline) String() (result string) {
	return fmt.Sprintf("&DeclinedTransactionSourceCardDecline{MerchantAcceptorID:%s MerchantDescriptor:%s MerchantCategoryCode:%s MerchantCity:%s MerchantCountry:%s Network:%s NetworkDetails:%s Amount:%s Currency:%s Reason:%s MerchantState:%s RealTimeDecisionID:%s DigitalWalletTokenID:%s}", core.FmtP(r.MerchantAcceptorID), core.FmtP(r.MerchantDescriptor), core.FmtP(r.MerchantCategoryCode), core.FmtP(r.MerchantCity), core.FmtP(r.MerchantCountry), core.FmtP(r.Network), r.NetworkDetails, core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.Reason), core.FmtP(r.MerchantState), core.FmtP(r.RealTimeDecisionID), core.FmtP(r.DigitalWalletTokenID))
}

type DeclinedTransactionSourceCardDeclineNetwork string

const (
	DeclinedTransactionSourceCardDeclineNetworkVisa DeclinedTransactionSourceCardDeclineNetwork = "visa"
)

type DeclinedTransactionSourceCardDeclineNetworkDetails struct {
	// Fields specific to the `visa` network
	Visa       *DeclinedTransactionSourceCardDeclineNetworkDetailsVisa `pjson:"visa"`
	jsonFields map[string]interface{}                                  `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// DeclinedTransactionSourceCardDeclineNetworkDetails using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *DeclinedTransactionSourceCardDeclineNetworkDetails) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes DeclinedTransactionSourceCardDeclineNetworkDetails into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *DeclinedTransactionSourceCardDeclineNetworkDetails) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// Fields specific to the `visa` network
func (r *DeclinedTransactionSourceCardDeclineNetworkDetails) GetVisa() (Visa DeclinedTransactionSourceCardDeclineNetworkDetailsVisa) {
	if r != nil && r.Visa != nil {
		Visa = *r.Visa
	}
	return
}

func (r DeclinedTransactionSourceCardDeclineNetworkDetails) String() (result string) {
	return fmt.Sprintf("&DeclinedTransactionSourceCardDeclineNetworkDetails{Visa:%s}", r.Visa)
}

type DeclinedTransactionSourceCardDeclineNetworkDetailsVisa struct {
	// For electronic commerce transactions, this identifies the level of security used
	// in obtaining the customer's payment credential. For mail or telephone order
	// transactions, identifies the type of mail or telephone order.
	ElectronicCommerceIndicator *DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator `pjson:"electronic_commerce_indicator"`
	// The method used to enter the cardholder's primary account number and card
	// expiration date
	PointOfServiceEntryMode *PointOfServiceEntryMode `pjson:"point_of_service_entry_mode"`
	jsonFields              map[string]interface{}   `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// DeclinedTransactionSourceCardDeclineNetworkDetailsVisa using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *DeclinedTransactionSourceCardDeclineNetworkDetailsVisa) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes DeclinedTransactionSourceCardDeclineNetworkDetailsVisa
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *DeclinedTransactionSourceCardDeclineNetworkDetailsVisa) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// For electronic commerce transactions, this identifies the level of security used
// in obtaining the customer's payment credential. For mail or telephone order
// transactions, identifies the type of mail or telephone order.
func (r *DeclinedTransactionSourceCardDeclineNetworkDetailsVisa) GetElectronicCommerceIndicator() (ElectronicCommerceIndicator DeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator) {
	if r != nil && r.ElectronicCommerceIndicator != nil {
		ElectronicCommerceIndicator = *r.ElectronicCommerceIndicator
	}
	return
}

// The method used to enter the cardholder's primary account number and card
// expiration date
func (r *DeclinedTransactionSourceCardDeclineNetworkDetailsVisa) GetPointOfServiceEntryMode() (PointOfServiceEntryMode PointOfServiceEntryMode) {
	if r != nil && r.PointOfServiceEntryMode != nil {
		PointOfServiceEntryMode = *r.PointOfServiceEntryMode
	}
	return
}

func (r DeclinedTransactionSourceCardDeclineNetworkDetailsVisa) String() (result string) {
	return fmt.Sprintf("&DeclinedTransactionSourceCardDeclineNetworkDetailsVisa{ElectronicCommerceIndicator:%s PointOfServiceEntryMode:%s}", core.FmtP(r.ElectronicCommerceIndicator), core.FmtP(r.PointOfServiceEntryMode))
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
)

type DeclinedTransactionSourceCheckDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount        *int64  `pjson:"amount"`
	AuxiliaryOnUs *string `pjson:"auxiliary_on_us"`
	// Why the check was declined.
	Reason     *DeclinedTransactionSourceCheckDeclineReason `pjson:"reason"`
	jsonFields map[string]interface{}                       `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// DeclinedTransactionSourceCheckDecline using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *DeclinedTransactionSourceCheckDecline) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes DeclinedTransactionSourceCheckDecline into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *DeclinedTransactionSourceCheckDecline) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *DeclinedTransactionSourceCheckDecline) GetAmount() (Amount int64) {
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

func (r DeclinedTransactionSourceCheckDecline) String() (result string) {
	return fmt.Sprintf("&DeclinedTransactionSourceCheckDecline{Amount:%s AuxiliaryOnUs:%s Reason:%s}", core.FmtP(r.Amount), core.FmtP(r.AuxiliaryOnUs), core.FmtP(r.Reason))
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
)

type DeclinedTransactionSourceInboundRealTimePaymentsTransferDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the declined
	// transfer's currency. This will always be "USD" for a Real Time Payments
	// transfer.
	Currency *DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency `pjson:"currency"`
	// Why the transfer was declined.
	Reason *DeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason `pjson:"reason"`
	// The name the sender of the transfer specified as the recipient of the transfer.
	CreditorName *string `pjson:"creditor_name"`
	// The name provided by the sender of the transfer.
	DebtorName *string `pjson:"debtor_name"`
	// The account number of the account that sent the transfer.
	DebtorAccountNumber *string `pjson:"debtor_account_number"`
	// The routing number of the account that sent the transfer.
	DebtorRoutingNumber *string `pjson:"debtor_routing_number"`
	// The Real Time Payments network identification of the declined transfer.
	TransactionIdentification *string `pjson:"transaction_identification"`
	// Additional information included with the transfer.
	RemittanceInformation *string                `pjson:"remittance_information"`
	jsonFields            map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// DeclinedTransactionSourceInboundRealTimePaymentsTransferDecline using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *DeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// DeclinedTransactionSourceInboundRealTimePaymentsTransferDecline into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *DeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *DeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) GetAmount() (Amount int64) {
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

func (r DeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) String() (result string) {
	return fmt.Sprintf("&DeclinedTransactionSourceInboundRealTimePaymentsTransferDecline{Amount:%s Currency:%s Reason:%s CreditorName:%s DebtorName:%s DebtorAccountNumber:%s DebtorRoutingNumber:%s TransactionIdentification:%s RemittanceInformation:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.Reason), core.FmtP(r.CreditorName), core.FmtP(r.DebtorName), core.FmtP(r.DebtorAccountNumber), core.FmtP(r.DebtorRoutingNumber), core.FmtP(r.TransactionIdentification), core.FmtP(r.RemittanceInformation))
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
	Amount                                                 *int64                 `pjson:"amount"`
	ForeignExchangeIndicator                               *string                `pjson:"foreign_exchange_indicator"`
	ForeignExchangeReferenceIndicator                      *string                `pjson:"foreign_exchange_reference_indicator"`
	ForeignExchangeReference                               *string                `pjson:"foreign_exchange_reference"`
	DestinationCountryCode                                 *string                `pjson:"destination_country_code"`
	DestinationCurrencyCode                                *string                `pjson:"destination_currency_code"`
	ForeignPaymentAmount                                   *int64                 `pjson:"foreign_payment_amount"`
	ForeignTraceNumber                                     *string                `pjson:"foreign_trace_number"`
	InternationalTransactionTypeCode                       *string                `pjson:"international_transaction_type_code"`
	OriginatingCurrencyCode                                *string                `pjson:"originating_currency_code"`
	OriginatingDepositoryFinancialInstitutionName          *string                `pjson:"originating_depository_financial_institution_name"`
	OriginatingDepositoryFinancialInstitutionIDQualifier   *string                `pjson:"originating_depository_financial_institution_id_qualifier"`
	OriginatingDepositoryFinancialInstitutionID            *string                `pjson:"originating_depository_financial_institution_id"`
	OriginatingDepositoryFinancialInstitutionBranchCountry *string                `pjson:"originating_depository_financial_institution_branch_country"`
	OriginatorCity                                         *string                `pjson:"originator_city"`
	OriginatorCompanyEntryDescription                      *string                `pjson:"originator_company_entry_description"`
	OriginatorCountry                                      *string                `pjson:"originator_country"`
	OriginatorIdentification                               *string                `pjson:"originator_identification"`
	OriginatorName                                         *string                `pjson:"originator_name"`
	OriginatorPostalCode                                   *string                `pjson:"originator_postal_code"`
	OriginatorStreetAddress                                *string                `pjson:"originator_street_address"`
	OriginatorStateOrProvince                              *string                `pjson:"originator_state_or_province"`
	PaymentRelatedInformation                              *string                `pjson:"payment_related_information"`
	PaymentRelatedInformation2                             *string                `pjson:"payment_related_information2"`
	ReceiverIdentificationNumber                           *string                `pjson:"receiver_identification_number"`
	ReceiverStreetAddress                                  *string                `pjson:"receiver_street_address"`
	ReceiverCity                                           *string                `pjson:"receiver_city"`
	ReceiverStateOrProvince                                *string                `pjson:"receiver_state_or_province"`
	ReceiverCountry                                        *string                `pjson:"receiver_country"`
	ReceiverPostalCode                                     *string                `pjson:"receiver_postal_code"`
	ReceivingCompanyOrIndividualName                       *string                `pjson:"receiving_company_or_individual_name"`
	ReceivingDepositoryFinancialInstitutionName            *string                `pjson:"receiving_depository_financial_institution_name"`
	ReceivingDepositoryFinancialInstitutionIDQualifier     *string                `pjson:"receiving_depository_financial_institution_id_qualifier"`
	ReceivingDepositoryFinancialInstitutionID              *string                `pjson:"receiving_depository_financial_institution_id"`
	ReceivingDepositoryFinancialInstitutionCountry         *string                `pjson:"receiving_depository_financial_institution_country"`
	TraceNumber                                            *string                `pjson:"trace_number"`
	jsonFields                                             map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// DeclinedTransactionSourceInternationalACHDecline using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *DeclinedTransactionSourceInternationalACHDecline) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes DeclinedTransactionSourceInternationalACHDecline into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *DeclinedTransactionSourceInternationalACHDecline) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *DeclinedTransactionSourceInternationalACHDecline) GetAmount() (Amount int64) {
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

func (r *DeclinedTransactionSourceInternationalACHDecline) GetForeignPaymentAmount() (ForeignPaymentAmount int64) {
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

func (r DeclinedTransactionSourceInternationalACHDecline) String() (result string) {
	return fmt.Sprintf("&DeclinedTransactionSourceInternationalACHDecline{Amount:%s ForeignExchangeIndicator:%s ForeignExchangeReferenceIndicator:%s ForeignExchangeReference:%s DestinationCountryCode:%s DestinationCurrencyCode:%s ForeignPaymentAmount:%s ForeignTraceNumber:%s InternationalTransactionTypeCode:%s OriginatingCurrencyCode:%s OriginatingDepositoryFinancialInstitutionName:%s OriginatingDepositoryFinancialInstitutionIDQualifier:%s OriginatingDepositoryFinancialInstitutionID:%s OriginatingDepositoryFinancialInstitutionBranchCountry:%s OriginatorCity:%s OriginatorCompanyEntryDescription:%s OriginatorCountry:%s OriginatorIdentification:%s OriginatorName:%s OriginatorPostalCode:%s OriginatorStreetAddress:%s OriginatorStateOrProvince:%s PaymentRelatedInformation:%s PaymentRelatedInformation2:%s ReceiverIdentificationNumber:%s ReceiverStreetAddress:%s ReceiverCity:%s ReceiverStateOrProvince:%s ReceiverCountry:%s ReceiverPostalCode:%s ReceivingCompanyOrIndividualName:%s ReceivingDepositoryFinancialInstitutionName:%s ReceivingDepositoryFinancialInstitutionIDQualifier:%s ReceivingDepositoryFinancialInstitutionID:%s ReceivingDepositoryFinancialInstitutionCountry:%s TraceNumber:%s}", core.FmtP(r.Amount), core.FmtP(r.ForeignExchangeIndicator), core.FmtP(r.ForeignExchangeReferenceIndicator), core.FmtP(r.ForeignExchangeReference), core.FmtP(r.DestinationCountryCode), core.FmtP(r.DestinationCurrencyCode), core.FmtP(r.ForeignPaymentAmount), core.FmtP(r.ForeignTraceNumber), core.FmtP(r.InternationalTransactionTypeCode), core.FmtP(r.OriginatingCurrencyCode), core.FmtP(r.OriginatingDepositoryFinancialInstitutionName), core.FmtP(r.OriginatingDepositoryFinancialInstitutionIDQualifier), core.FmtP(r.OriginatingDepositoryFinancialInstitutionID), core.FmtP(r.OriginatingDepositoryFinancialInstitutionBranchCountry), core.FmtP(r.OriginatorCity), core.FmtP(r.OriginatorCompanyEntryDescription), core.FmtP(r.OriginatorCountry), core.FmtP(r.OriginatorIdentification), core.FmtP(r.OriginatorName), core.FmtP(r.OriginatorPostalCode), core.FmtP(r.OriginatorStreetAddress), core.FmtP(r.OriginatorStateOrProvince), core.FmtP(r.PaymentRelatedInformation), core.FmtP(r.PaymentRelatedInformation2), core.FmtP(r.ReceiverIdentificationNumber), core.FmtP(r.ReceiverStreetAddress), core.FmtP(r.ReceiverCity), core.FmtP(r.ReceiverStateOrProvince), core.FmtP(r.ReceiverCountry), core.FmtP(r.ReceiverPostalCode), core.FmtP(r.ReceivingCompanyOrIndividualName), core.FmtP(r.ReceivingDepositoryFinancialInstitutionName), core.FmtP(r.ReceivingDepositoryFinancialInstitutionIDQualifier), core.FmtP(r.ReceivingDepositoryFinancialInstitutionID), core.FmtP(r.ReceivingDepositoryFinancialInstitutionCountry), core.FmtP(r.TraceNumber))
}

type DeclinedTransactionSourceCardRouteDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency             *DeclinedTransactionSourceCardRouteDeclineCurrency `pjson:"currency"`
	MerchantAcceptorID   *string                                            `pjson:"merchant_acceptor_id"`
	MerchantCity         *string                                            `pjson:"merchant_city"`
	MerchantCountry      *string                                            `pjson:"merchant_country"`
	MerchantDescriptor   *string                                            `pjson:"merchant_descriptor"`
	MerchantState        *string                                            `pjson:"merchant_state"`
	MerchantCategoryCode *string                                            `pjson:"merchant_category_code"`
	jsonFields           map[string]interface{}                             `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// DeclinedTransactionSourceCardRouteDecline using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *DeclinedTransactionSourceCardRouteDecline) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes DeclinedTransactionSourceCardRouteDecline into an array
// of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *DeclinedTransactionSourceCardRouteDecline) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *DeclinedTransactionSourceCardRouteDecline) GetAmount() (Amount int64) {
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

func (r DeclinedTransactionSourceCardRouteDecline) String() (result string) {
	return fmt.Sprintf("&DeclinedTransactionSourceCardRouteDecline{Amount:%s Currency:%s MerchantAcceptorID:%s MerchantCity:%s MerchantCountry:%s MerchantDescriptor:%s MerchantState:%s MerchantCategoryCode:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.MerchantAcceptorID), core.FmtP(r.MerchantCity), core.FmtP(r.MerchantCountry), core.FmtP(r.MerchantDescriptor), core.FmtP(r.MerchantState), core.FmtP(r.MerchantCategoryCode))
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
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int64 `query:"limit"`
	// Filter Declined Transactions to ones belonging to the specified Account.
	AccountID *string `query:"account_id"`
	// Filter Declined Transactions to those belonging to the specified route.
	RouteID    *string                                  `query:"route_id"`
	CreatedAt  *DeclinedTransactionsListParamsCreatedAt `query:"created_at"`
	jsonFields map[string]interface{}                   `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into DeclinedTransactionListParams
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *DeclinedTransactionListParams) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes DeclinedTransactionListParams into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *DeclinedTransactionListParams) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes DeclinedTransactionListParams into a url.Values of the query
// parameters associated with this value
func (r *DeclinedTransactionListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// Return the page of entries after this one.
func (r *DeclinedTransactionListParams) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *DeclinedTransactionListParams) GetLimit() (Limit int64) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// Filter Declined Transactions to ones belonging to the specified Account.
func (r *DeclinedTransactionListParams) GetAccountID() (AccountID string) {
	if r != nil && r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

// Filter Declined Transactions to those belonging to the specified route.
func (r *DeclinedTransactionListParams) GetRouteID() (RouteID string) {
	if r != nil && r.RouteID != nil {
		RouteID = *r.RouteID
	}
	return
}

func (r *DeclinedTransactionListParams) GetCreatedAt() (CreatedAt DeclinedTransactionsListParamsCreatedAt) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

func (r DeclinedTransactionListParams) String() (result string) {
	return fmt.Sprintf("&DeclinedTransactionListParams{Cursor:%s Limit:%s AccountID:%s RouteID:%s CreatedAt:%s}", core.FmtP(r.Cursor), core.FmtP(r.Limit), core.FmtP(r.AccountID), core.FmtP(r.RouteID), r.CreatedAt)
}

type DeclinedTransactionsListParamsCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After *string `pjson:"after"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before *string `pjson:"before"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter *string `pjson:"on_or_after"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore *string                `pjson:"on_or_before"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// DeclinedTransactionsListParamsCreatedAt using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *DeclinedTransactionsListParamsCreatedAt) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes DeclinedTransactionsListParamsCreatedAt into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *DeclinedTransactionsListParamsCreatedAt) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes DeclinedTransactionsListParamsCreatedAt into a url.Values of
// the query parameters associated with this value
func (r *DeclinedTransactionsListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r *DeclinedTransactionsListParamsCreatedAt) GetAfter() (After string) {
	if r != nil && r.After != nil {
		After = *r.After
	}
	return
}

// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r *DeclinedTransactionsListParamsCreatedAt) GetBefore() (Before string) {
	if r != nil && r.Before != nil {
		Before = *r.Before
	}
	return
}

// Return results on or after this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r *DeclinedTransactionsListParamsCreatedAt) GetOnOrAfter() (OnOrAfter string) {
	if r != nil && r.OnOrAfter != nil {
		OnOrAfter = *r.OnOrAfter
	}
	return
}

// Return results on or before this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r *DeclinedTransactionsListParamsCreatedAt) GetOnOrBefore() (OnOrBefore string) {
	if r != nil && r.OnOrBefore != nil {
		OnOrBefore = *r.OnOrBefore
	}
	return
}

func (r DeclinedTransactionsListParamsCreatedAt) String() (result string) {
	return fmt.Sprintf("&DeclinedTransactionsListParamsCreatedAt{After:%s Before:%s OnOrAfter:%s OnOrBefore:%s}", core.FmtP(r.After), core.FmtP(r.Before), core.FmtP(r.OnOrAfter), core.FmtP(r.OnOrBefore))
}

type DeclinedTransactionList struct {
	// The contents of the list.
	Data *[]DeclinedTransaction `pjson:"data"`
	// A pointer to a place in the list.
	NextCursor *string                `pjson:"next_cursor"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into DeclinedTransactionList using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *DeclinedTransactionList) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes DeclinedTransactionList into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *DeclinedTransactionList) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes DeclinedTransactionList into a url.Values of the query
// parameters associated with this value
func (r *DeclinedTransactionList) URLQuery() (v url.Values) {
	return query.Marshal(r)
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

func (r DeclinedTransactionList) String() (result string) {
	return fmt.Sprintf("&DeclinedTransactionList{Data:%s NextCursor:%s}", core.Fmt(r.Data), core.FmtP(r.NextCursor))
}

type DeclinedTransactionsPage struct {
	*pagination.Page[DeclinedTransaction]
}

func (r *DeclinedTransactionsPage) DeclinedTransaction() *DeclinedTransaction {
	return r.Current()
}

func (r *DeclinedTransactionsPage) NextPage() (*DeclinedTransactionsPage, error) {
	if page, err := r.Page.NextPage(); err != nil {
		return nil, err
	} else {
		return &DeclinedTransactionsPage{page}, nil
	}
}

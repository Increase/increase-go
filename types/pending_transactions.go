package types

import (
	"fmt"
	"increase/core"
	"increase/core/pjson"
	"increase/core/query"
	"increase/pagination"
	"net/url"
)

type PendingTransaction struct {
	// The identifier for the account this Pending Transaction belongs to.
	AccountID *string `pjson:"account_id"`
	// The Pending Transaction amount in the minor unit of its currency. For dollars,
	// for example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Pending
	// Transaction's currency. This will match the currency on the Pending
	// Transcation's Account.
	Currency *PendingTransactionCurrency `pjson:"currency"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the Pending
	// Transaction occured.
	CreatedAt *string `pjson:"created_at"`
	// For a Pending Transaction related to a transfer, this is the description you
	// provide. For a Pending Transaction related to a payment, this is the description
	// the vendor provides.
	Description *string `pjson:"description"`
	// The Pending Transaction identifier.
	ID *string `pjson:"id"`
	// The identifier for the route this Pending Transaction came through. Routes are
	// things like cards and ACH details.
	RouteID *string `pjson:"route_id"`
	// The type of the route this Pending Transaction came through.
	RouteType *string `pjson:"route_type"`
	// This is an object giving more details on the network-level event that caused the
	// Pending Transaction. For example, for a card transaction this lists the
	// merchant's industry and location.
	Source *PendingTransactionSource `pjson:"source"`
	// Whether the Pending Transaction has been confirmed and has an associated
	// Transaction.
	Status *PendingTransactionStatus `pjson:"status"`
	// A constant representing the object's type. For this resource it will always be
	// `pending_transaction`.
	Type       *PendingTransactionType `pjson:"type"`
	jsonFields map[string]interface{}  `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into PendingTransaction using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *PendingTransaction) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes PendingTransaction into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *PendingTransaction) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The identifier for the account this Pending Transaction belongs to.
func (r *PendingTransaction) GetAccountID() (AccountID string) {
	if r != nil && r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

// The Pending Transaction amount in the minor unit of its currency. For dollars,
// for example, this is cents.
func (r *PendingTransaction) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Pending
// Transaction's currency. This will match the currency on the Pending
// Transcation's Account.
func (r *PendingTransaction) GetCurrency() (Currency PendingTransactionCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the Pending
// Transaction occured.
func (r *PendingTransaction) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// For a Pending Transaction related to a transfer, this is the description you
// provide. For a Pending Transaction related to a payment, this is the description
// the vendor provides.
func (r *PendingTransaction) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The Pending Transaction identifier.
func (r *PendingTransaction) GetID() (ID string) {
	if r != nil && r.ID != nil {
		ID = *r.ID
	}
	return
}

// The identifier for the route this Pending Transaction came through. Routes are
// things like cards and ACH details.
func (r *PendingTransaction) GetRouteID() (RouteID string) {
	if r != nil && r.RouteID != nil {
		RouteID = *r.RouteID
	}
	return
}

// The type of the route this Pending Transaction came through.
func (r *PendingTransaction) GetRouteType() (RouteType string) {
	if r != nil && r.RouteType != nil {
		RouteType = *r.RouteType
	}
	return
}

// This is an object giving more details on the network-level event that caused the
// Pending Transaction. For example, for a card transaction this lists the
// merchant's industry and location.
func (r *PendingTransaction) GetSource() (Source PendingTransactionSource) {
	if r != nil && r.Source != nil {
		Source = *r.Source
	}
	return
}

// Whether the Pending Transaction has been confirmed and has an associated
// Transaction.
func (r *PendingTransaction) GetStatus() (Status PendingTransactionStatus) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `pending_transaction`.
func (r *PendingTransaction) GetType() (Type PendingTransactionType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

func (r PendingTransaction) String() (result string) {
	return fmt.Sprintf("&PendingTransaction{AccountID:%s Amount:%s Currency:%s CreatedAt:%s Description:%s ID:%s RouteID:%s RouteType:%s Source:%s Status:%s Type:%s}", core.FmtP(r.AccountID), core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.CreatedAt), core.FmtP(r.Description), core.FmtP(r.ID), core.FmtP(r.RouteID), core.FmtP(r.RouteType), r.Source, core.FmtP(r.Status), core.FmtP(r.Type))
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
	Category *PendingTransactionSourceCategory `pjson:"category"`
	// A Account Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `account_transfer_instruction`.
	AccountTransferInstruction *PendingTransactionSourceAccountTransferInstruction `pjson:"account_transfer_instruction"`
	// A ACH Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `ach_transfer_instruction`.
	ACHTransferInstruction *PendingTransactionSourceACHTransferInstruction `pjson:"ach_transfer_instruction"`
	// A Card Authorization object. This field will be present in the JSON response if
	// and only if `category` is equal to `card_authorization`.
	CardAuthorization *PendingTransactionSourceCardAuthorization `pjson:"card_authorization"`
	// A Check Deposit Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_deposit_instruction`.
	CheckDepositInstruction *PendingTransactionSourceCheckDepositInstruction `pjson:"check_deposit_instruction"`
	// A Check Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_instruction`.
	CheckTransferInstruction *PendingTransactionSourceCheckTransferInstruction `pjson:"check_transfer_instruction"`
	// A Inbound Funds Hold object. This field will be present in the JSON response if
	// and only if `category` is equal to `inbound_funds_hold`.
	InboundFundsHold *PendingTransactionSourceInboundFundsHold `pjson:"inbound_funds_hold"`
	// A Deprecated Card Authorization object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_route_authorization`.
	CardRouteAuthorization *PendingTransactionSourceCardRouteAuthorization `pjson:"card_route_authorization"`
	// A Wire Drawdown Payment Instruction object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `wire_drawdown_payment_instruction`.
	WireDrawdownPaymentInstruction *PendingTransactionSourceWireDrawdownPaymentInstruction `pjson:"wire_drawdown_payment_instruction"`
	// A Wire Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_instruction`.
	WireTransferInstruction *PendingTransactionSourceWireTransferInstruction `pjson:"wire_transfer_instruction"`
	jsonFields              map[string]interface{}                           `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into PendingTransactionSource
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *PendingTransactionSource) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes PendingTransactionSource into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *PendingTransactionSource) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The type of transaction that took place. We may add additional possible values
// for this enum over time; your application should be able to handle such
// additions gracefully.
func (r *PendingTransactionSource) GetCategory() (Category PendingTransactionSourceCategory) {
	if r != nil && r.Category != nil {
		Category = *r.Category
	}
	return
}

// A Account Transfer Instruction object. This field will be present in the JSON
// response if and only if `category` is equal to `account_transfer_instruction`.
func (r *PendingTransactionSource) GetAccountTransferInstruction() (AccountTransferInstruction PendingTransactionSourceAccountTransferInstruction) {
	if r != nil && r.AccountTransferInstruction != nil {
		AccountTransferInstruction = *r.AccountTransferInstruction
	}
	return
}

// A ACH Transfer Instruction object. This field will be present in the JSON
// response if and only if `category` is equal to `ach_transfer_instruction`.
func (r *PendingTransactionSource) GetACHTransferInstruction() (ACHTransferInstruction PendingTransactionSourceACHTransferInstruction) {
	if r != nil && r.ACHTransferInstruction != nil {
		ACHTransferInstruction = *r.ACHTransferInstruction
	}
	return
}

// A Card Authorization object. This field will be present in the JSON response if
// and only if `category` is equal to `card_authorization`.
func (r *PendingTransactionSource) GetCardAuthorization() (CardAuthorization PendingTransactionSourceCardAuthorization) {
	if r != nil && r.CardAuthorization != nil {
		CardAuthorization = *r.CardAuthorization
	}
	return
}

// A Check Deposit Instruction object. This field will be present in the JSON
// response if and only if `category` is equal to `check_deposit_instruction`.
func (r *PendingTransactionSource) GetCheckDepositInstruction() (CheckDepositInstruction PendingTransactionSourceCheckDepositInstruction) {
	if r != nil && r.CheckDepositInstruction != nil {
		CheckDepositInstruction = *r.CheckDepositInstruction
	}
	return
}

// A Check Transfer Instruction object. This field will be present in the JSON
// response if and only if `category` is equal to `check_transfer_instruction`.
func (r *PendingTransactionSource) GetCheckTransferInstruction() (CheckTransferInstruction PendingTransactionSourceCheckTransferInstruction) {
	if r != nil && r.CheckTransferInstruction != nil {
		CheckTransferInstruction = *r.CheckTransferInstruction
	}
	return
}

// A Inbound Funds Hold object. This field will be present in the JSON response if
// and only if `category` is equal to `inbound_funds_hold`.
func (r *PendingTransactionSource) GetInboundFundsHold() (InboundFundsHold PendingTransactionSourceInboundFundsHold) {
	if r != nil && r.InboundFundsHold != nil {
		InboundFundsHold = *r.InboundFundsHold
	}
	return
}

// A Deprecated Card Authorization object. This field will be present in the JSON
// response if and only if `category` is equal to `card_route_authorization`.
func (r *PendingTransactionSource) GetCardRouteAuthorization() (CardRouteAuthorization PendingTransactionSourceCardRouteAuthorization) {
	if r != nil && r.CardRouteAuthorization != nil {
		CardRouteAuthorization = *r.CardRouteAuthorization
	}
	return
}

// A Wire Drawdown Payment Instruction object. This field will be present in the
// JSON response if and only if `category` is equal to
// `wire_drawdown_payment_instruction`.
func (r *PendingTransactionSource) GetWireDrawdownPaymentInstruction() (WireDrawdownPaymentInstruction PendingTransactionSourceWireDrawdownPaymentInstruction) {
	if r != nil && r.WireDrawdownPaymentInstruction != nil {
		WireDrawdownPaymentInstruction = *r.WireDrawdownPaymentInstruction
	}
	return
}

// A Wire Transfer Instruction object. This field will be present in the JSON
// response if and only if `category` is equal to `wire_transfer_instruction`.
func (r *PendingTransactionSource) GetWireTransferInstruction() (WireTransferInstruction PendingTransactionSourceWireTransferInstruction) {
	if r != nil && r.WireTransferInstruction != nil {
		WireTransferInstruction = *r.WireTransferInstruction
	}
	return
}

func (r PendingTransactionSource) String() (result string) {
	return fmt.Sprintf("&PendingTransactionSource{Category:%s AccountTransferInstruction:%s ACHTransferInstruction:%s CardAuthorization:%s CheckDepositInstruction:%s CheckTransferInstruction:%s InboundFundsHold:%s CardRouteAuthorization:%s WireDrawdownPaymentInstruction:%s WireTransferInstruction:%s}", core.FmtP(r.Category), r.AccountTransferInstruction, r.ACHTransferInstruction, r.CardAuthorization, r.CheckDepositInstruction, r.CheckTransferInstruction, r.InboundFundsHold, r.CardRouteAuthorization, r.WireDrawdownPaymentInstruction, r.WireTransferInstruction)
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
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency *PendingTransactionSourceAccountTransferInstructionCurrency `pjson:"currency"`
	// The identifier of the Account Transfer that led to this Pending Transaction.
	TransferID *string                `pjson:"transfer_id"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// PendingTransactionSourceAccountTransferInstruction using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *PendingTransactionSourceAccountTransferInstruction) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes PendingTransactionSourceAccountTransferInstruction into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *PendingTransactionSourceAccountTransferInstruction) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *PendingTransactionSourceAccountTransferInstruction) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
// account currency.
func (r *PendingTransactionSourceAccountTransferInstruction) GetCurrency() (Currency PendingTransactionSourceAccountTransferInstructionCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The identifier of the Account Transfer that led to this Pending Transaction.
func (r *PendingTransactionSourceAccountTransferInstruction) GetTransferID() (TransferID string) {
	if r != nil && r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

func (r PendingTransactionSourceAccountTransferInstruction) String() (result string) {
	return fmt.Sprintf("&PendingTransactionSourceAccountTransferInstruction{Amount:%s Currency:%s TransferID:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.TransferID))
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
	Amount *int64 `pjson:"amount"`
	// The identifier of the ACH Transfer that led to this Pending Transaction.
	TransferID *string                `pjson:"transfer_id"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// PendingTransactionSourceACHTransferInstruction using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *PendingTransactionSourceACHTransferInstruction) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes PendingTransactionSourceACHTransferInstruction into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *PendingTransactionSourceACHTransferInstruction) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *PendingTransactionSourceACHTransferInstruction) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The identifier of the ACH Transfer that led to this Pending Transaction.
func (r *PendingTransactionSourceACHTransferInstruction) GetTransferID() (TransferID string) {
	if r != nil && r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

func (r PendingTransactionSourceACHTransferInstruction) String() (result string) {
	return fmt.Sprintf("&PendingTransactionSourceACHTransferInstruction{Amount:%s TransferID:%s}", core.FmtP(r.Amount), core.FmtP(r.TransferID))
}

type PendingTransactionSourceCardAuthorization struct {
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
	Network *PendingTransactionSourceCardAuthorizationNetwork `pjson:"network"`
	// Fields specific to the `network`
	NetworkDetails *PendingTransactionSourceCardAuthorizationNetworkDetails `pjson:"network_details"`
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *PendingTransactionSourceCardAuthorizationCurrency `pjson:"currency"`
	// The identifier of the Real-Time Decision sent to approve or decline this
	// transaction.
	RealTimeDecisionID *string `pjson:"real_time_decision_id"`
	// If the authorization was made via a Digital Wallet Token (such as an Apple Pay
	// purchase), the identifier of the token that was used.
	DigitalWalletTokenID *string                `pjson:"digital_wallet_token_id"`
	jsonFields           map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// PendingTransactionSourceCardAuthorization using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *PendingTransactionSourceCardAuthorization) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes PendingTransactionSourceCardAuthorization into an array
// of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *PendingTransactionSourceCardAuthorization) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The merchant identifier (commonly abbreviated as MID) of the merchant the card
// is transacting with.
func (r *PendingTransactionSourceCardAuthorization) GetMerchantAcceptorID() (MerchantAcceptorID string) {
	if r != nil && r.MerchantAcceptorID != nil {
		MerchantAcceptorID = *r.MerchantAcceptorID
	}
	return
}

// The merchant descriptor of the merchant the card is transacting with.
func (r *PendingTransactionSourceCardAuthorization) GetMerchantDescriptor() (MerchantDescriptor string) {
	if r != nil && r.MerchantDescriptor != nil {
		MerchantDescriptor = *r.MerchantDescriptor
	}
	return
}

// The Merchant Category Code (commonly abbreviated as MCC) of the merchant the
// card is transacting with.
func (r *PendingTransactionSourceCardAuthorization) GetMerchantCategoryCode() (MerchantCategoryCode string) {
	if r != nil && r.MerchantCategoryCode != nil {
		MerchantCategoryCode = *r.MerchantCategoryCode
	}
	return
}

// The city the merchant resides in.
func (r *PendingTransactionSourceCardAuthorization) GetMerchantCity() (MerchantCity string) {
	if r != nil && r.MerchantCity != nil {
		MerchantCity = *r.MerchantCity
	}
	return
}

// The country the merchant resides in.
func (r *PendingTransactionSourceCardAuthorization) GetMerchantCountry() (MerchantCountry string) {
	if r != nil && r.MerchantCountry != nil {
		MerchantCountry = *r.MerchantCountry
	}
	return
}

// The payment network used to process this card authorization
func (r *PendingTransactionSourceCardAuthorization) GetNetwork() (Network PendingTransactionSourceCardAuthorizationNetwork) {
	if r != nil && r.Network != nil {
		Network = *r.Network
	}
	return
}

// Fields specific to the `network`
func (r *PendingTransactionSourceCardAuthorization) GetNetworkDetails() (NetworkDetails PendingTransactionSourceCardAuthorizationNetworkDetails) {
	if r != nil && r.NetworkDetails != nil {
		NetworkDetails = *r.NetworkDetails
	}
	return
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *PendingTransactionSourceCardAuthorization) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *PendingTransactionSourceCardAuthorization) GetCurrency() (Currency PendingTransactionSourceCardAuthorizationCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The identifier of the Real-Time Decision sent to approve or decline this
// transaction.
func (r *PendingTransactionSourceCardAuthorization) GetRealTimeDecisionID() (RealTimeDecisionID string) {
	if r != nil && r.RealTimeDecisionID != nil {
		RealTimeDecisionID = *r.RealTimeDecisionID
	}
	return
}

// If the authorization was made via a Digital Wallet Token (such as an Apple Pay
// purchase), the identifier of the token that was used.
func (r *PendingTransactionSourceCardAuthorization) GetDigitalWalletTokenID() (DigitalWalletTokenID string) {
	if r != nil && r.DigitalWalletTokenID != nil {
		DigitalWalletTokenID = *r.DigitalWalletTokenID
	}
	return
}

func (r PendingTransactionSourceCardAuthorization) String() (result string) {
	return fmt.Sprintf("&PendingTransactionSourceCardAuthorization{MerchantAcceptorID:%s MerchantDescriptor:%s MerchantCategoryCode:%s MerchantCity:%s MerchantCountry:%s Network:%s NetworkDetails:%s Amount:%s Currency:%s RealTimeDecisionID:%s DigitalWalletTokenID:%s}", core.FmtP(r.MerchantAcceptorID), core.FmtP(r.MerchantDescriptor), core.FmtP(r.MerchantCategoryCode), core.FmtP(r.MerchantCity), core.FmtP(r.MerchantCountry), core.FmtP(r.Network), r.NetworkDetails, core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.RealTimeDecisionID), core.FmtP(r.DigitalWalletTokenID))
}

type PendingTransactionSourceCardAuthorizationNetwork string

const (
	PendingTransactionSourceCardAuthorizationNetworkVisa PendingTransactionSourceCardAuthorizationNetwork = "visa"
)

type PendingTransactionSourceCardAuthorizationNetworkDetails struct {
	// Fields specific to the `visa` network
	Visa       *PendingTransactionSourceCardAuthorizationNetworkDetailsVisa `pjson:"visa"`
	jsonFields map[string]interface{}                                       `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// PendingTransactionSourceCardAuthorizationNetworkDetails using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *PendingTransactionSourceCardAuthorizationNetworkDetails) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes PendingTransactionSourceCardAuthorizationNetworkDetails
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *PendingTransactionSourceCardAuthorizationNetworkDetails) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// Fields specific to the `visa` network
func (r *PendingTransactionSourceCardAuthorizationNetworkDetails) GetVisa() (Visa PendingTransactionSourceCardAuthorizationNetworkDetailsVisa) {
	if r != nil && r.Visa != nil {
		Visa = *r.Visa
	}
	return
}

func (r PendingTransactionSourceCardAuthorizationNetworkDetails) String() (result string) {
	return fmt.Sprintf("&PendingTransactionSourceCardAuthorizationNetworkDetails{Visa:%s}", r.Visa)
}

type PendingTransactionSourceCardAuthorizationNetworkDetailsVisa struct {
	// For electronic commerce transactions, this identifies the level of security used
	// in obtaining the customer's payment credential. For mail or telephone order
	// transactions, identifies the type of mail or telephone order.
	ElectronicCommerceIndicator *PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator `pjson:"electronic_commerce_indicator"`
	// The method used to enter the cardholder's primary account number and card
	// expiration date
	PointOfServiceEntryMode *PointOfServiceEntryMode `pjson:"point_of_service_entry_mode"`
	jsonFields              map[string]interface{}   `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// PendingTransactionSourceCardAuthorizationNetworkDetailsVisa using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *PendingTransactionSourceCardAuthorizationNetworkDetailsVisa) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// PendingTransactionSourceCardAuthorizationNetworkDetailsVisa into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *PendingTransactionSourceCardAuthorizationNetworkDetailsVisa) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// For electronic commerce transactions, this identifies the level of security used
// in obtaining the customer's payment credential. For mail or telephone order
// transactions, identifies the type of mail or telephone order.
func (r *PendingTransactionSourceCardAuthorizationNetworkDetailsVisa) GetElectronicCommerceIndicator() (ElectronicCommerceIndicator PendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator) {
	if r != nil && r.ElectronicCommerceIndicator != nil {
		ElectronicCommerceIndicator = *r.ElectronicCommerceIndicator
	}
	return
}

// The method used to enter the cardholder's primary account number and card
// expiration date
func (r *PendingTransactionSourceCardAuthorizationNetworkDetailsVisa) GetPointOfServiceEntryMode() (PointOfServiceEntryMode PointOfServiceEntryMode) {
	if r != nil && r.PointOfServiceEntryMode != nil {
		PointOfServiceEntryMode = *r.PointOfServiceEntryMode
	}
	return
}

func (r PendingTransactionSourceCardAuthorizationNetworkDetailsVisa) String() (result string) {
	return fmt.Sprintf("&PendingTransactionSourceCardAuthorizationNetworkDetailsVisa{ElectronicCommerceIndicator:%s PointOfServiceEntryMode:%s}", core.FmtP(r.ElectronicCommerceIndicator), core.FmtP(r.PointOfServiceEntryMode))
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
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *PendingTransactionSourceCheckDepositInstructionCurrency `pjson:"currency"`
	// The identifier of the File containing the image of the front of the check that
	// was deposited.
	FrontImageFileID *string `pjson:"front_image_file_id"`
	// The identifier of the File containing the image of the back of the check that
	// was deposited.
	BackImageFileID *string `pjson:"back_image_file_id"`
	// The identifier of the Check Deposit.
	CheckDepositID *string                `pjson:"check_deposit_id"`
	jsonFields     map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// PendingTransactionSourceCheckDepositInstruction using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *PendingTransactionSourceCheckDepositInstruction) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes PendingTransactionSourceCheckDepositInstruction into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *PendingTransactionSourceCheckDepositInstruction) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *PendingTransactionSourceCheckDepositInstruction) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *PendingTransactionSourceCheckDepositInstruction) GetCurrency() (Currency PendingTransactionSourceCheckDepositInstructionCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The identifier of the File containing the image of the front of the check that
// was deposited.
func (r *PendingTransactionSourceCheckDepositInstruction) GetFrontImageFileID() (FrontImageFileID string) {
	if r != nil && r.FrontImageFileID != nil {
		FrontImageFileID = *r.FrontImageFileID
	}
	return
}

// The identifier of the File containing the image of the back of the check that
// was deposited.
func (r *PendingTransactionSourceCheckDepositInstruction) GetBackImageFileID() (BackImageFileID string) {
	if r != nil && r.BackImageFileID != nil {
		BackImageFileID = *r.BackImageFileID
	}
	return
}

// The identifier of the Check Deposit.
func (r *PendingTransactionSourceCheckDepositInstruction) GetCheckDepositID() (CheckDepositID string) {
	if r != nil && r.CheckDepositID != nil {
		CheckDepositID = *r.CheckDepositID
	}
	return
}

func (r PendingTransactionSourceCheckDepositInstruction) String() (result string) {
	return fmt.Sprintf("&PendingTransactionSourceCheckDepositInstruction{Amount:%s Currency:%s FrontImageFileID:%s BackImageFileID:%s CheckDepositID:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.FrontImageFileID), core.FmtP(r.BackImageFileID), core.FmtP(r.CheckDepositID))
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
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
	// currency.
	Currency *PendingTransactionSourceCheckTransferInstructionCurrency `pjson:"currency"`
	// The identifier of the Check Transfer that led to this Pending Transaction.
	TransferID *string                `pjson:"transfer_id"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// PendingTransactionSourceCheckTransferInstruction using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *PendingTransactionSourceCheckTransferInstruction) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes PendingTransactionSourceCheckTransferInstruction into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *PendingTransactionSourceCheckTransferInstruction) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *PendingTransactionSourceCheckTransferInstruction) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
// currency.
func (r *PendingTransactionSourceCheckTransferInstruction) GetCurrency() (Currency PendingTransactionSourceCheckTransferInstructionCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The identifier of the Check Transfer that led to this Pending Transaction.
func (r *PendingTransactionSourceCheckTransferInstruction) GetTransferID() (TransferID string) {
	if r != nil && r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

func (r PendingTransactionSourceCheckTransferInstruction) String() (result string) {
	return fmt.Sprintf("&PendingTransactionSourceCheckTransferInstruction{Amount:%s Currency:%s TransferID:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.TransferID))
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
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the hold's
	// currency.
	Currency *PendingTransactionSourceInboundFundsHoldCurrency `pjson:"currency"`
	// When the hold will be released automatically. Certain conditions may cause it to
	// be released before this time.
	AutomaticallyReleasesAt *string `pjson:"automatically_releases_at"`
	// When the hold was released (if it has been released).
	ReleasedAt *string `pjson:"released_at"`
	// The status of the hold.
	Status *PendingTransactionSourceInboundFundsHoldStatus `pjson:"status"`
	// The ID of the Transaction for which funds were held.
	HeldTransactionID *string                `pjson:"held_transaction_id"`
	jsonFields        map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// PendingTransactionSourceInboundFundsHold using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *PendingTransactionSourceInboundFundsHold) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes PendingTransactionSourceInboundFundsHold into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *PendingTransactionSourceInboundFundsHold) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The held amount in the minor unit of the account's currency. For dollars, for
// example, this is cents.
func (r *PendingTransactionSourceInboundFundsHold) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the hold's
// currency.
func (r *PendingTransactionSourceInboundFundsHold) GetCurrency() (Currency PendingTransactionSourceInboundFundsHoldCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// When the hold will be released automatically. Certain conditions may cause it to
// be released before this time.
func (r *PendingTransactionSourceInboundFundsHold) GetAutomaticallyReleasesAt() (AutomaticallyReleasesAt string) {
	if r != nil && r.AutomaticallyReleasesAt != nil {
		AutomaticallyReleasesAt = *r.AutomaticallyReleasesAt
	}
	return
}

// When the hold was released (if it has been released).
func (r *PendingTransactionSourceInboundFundsHold) GetReleasedAt() (ReleasedAt string) {
	if r != nil && r.ReleasedAt != nil {
		ReleasedAt = *r.ReleasedAt
	}
	return
}

// The status of the hold.
func (r *PendingTransactionSourceInboundFundsHold) GetStatus() (Status PendingTransactionSourceInboundFundsHoldStatus) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

// The ID of the Transaction for which funds were held.
func (r *PendingTransactionSourceInboundFundsHold) GetHeldTransactionID() (HeldTransactionID string) {
	if r != nil && r.HeldTransactionID != nil {
		HeldTransactionID = *r.HeldTransactionID
	}
	return
}

func (r PendingTransactionSourceInboundFundsHold) String() (result string) {
	return fmt.Sprintf("&PendingTransactionSourceInboundFundsHold{Amount:%s Currency:%s AutomaticallyReleasesAt:%s ReleasedAt:%s Status:%s HeldTransactionID:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.AutomaticallyReleasesAt), core.FmtP(r.ReleasedAt), core.FmtP(r.Status), core.FmtP(r.HeldTransactionID))
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
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency             *PendingTransactionSourceCardRouteAuthorizationCurrency `pjson:"currency"`
	MerchantAcceptorID   *string                                                 `pjson:"merchant_acceptor_id"`
	MerchantCity         *string                                                 `pjson:"merchant_city"`
	MerchantCountry      *string                                                 `pjson:"merchant_country"`
	MerchantDescriptor   *string                                                 `pjson:"merchant_descriptor"`
	MerchantCategoryCode *string                                                 `pjson:"merchant_category_code"`
	MerchantState        *string                                                 `pjson:"merchant_state"`
	jsonFields           map[string]interface{}                                  `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// PendingTransactionSourceCardRouteAuthorization using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *PendingTransactionSourceCardRouteAuthorization) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes PendingTransactionSourceCardRouteAuthorization into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *PendingTransactionSourceCardRouteAuthorization) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *PendingTransactionSourceCardRouteAuthorization) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *PendingTransactionSourceCardRouteAuthorization) GetCurrency() (Currency PendingTransactionSourceCardRouteAuthorizationCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

func (r *PendingTransactionSourceCardRouteAuthorization) GetMerchantAcceptorID() (MerchantAcceptorID string) {
	if r != nil && r.MerchantAcceptorID != nil {
		MerchantAcceptorID = *r.MerchantAcceptorID
	}
	return
}

func (r *PendingTransactionSourceCardRouteAuthorization) GetMerchantCity() (MerchantCity string) {
	if r != nil && r.MerchantCity != nil {
		MerchantCity = *r.MerchantCity
	}
	return
}

func (r *PendingTransactionSourceCardRouteAuthorization) GetMerchantCountry() (MerchantCountry string) {
	if r != nil && r.MerchantCountry != nil {
		MerchantCountry = *r.MerchantCountry
	}
	return
}

func (r *PendingTransactionSourceCardRouteAuthorization) GetMerchantDescriptor() (MerchantDescriptor string) {
	if r != nil && r.MerchantDescriptor != nil {
		MerchantDescriptor = *r.MerchantDescriptor
	}
	return
}

func (r *PendingTransactionSourceCardRouteAuthorization) GetMerchantCategoryCode() (MerchantCategoryCode string) {
	if r != nil && r.MerchantCategoryCode != nil {
		MerchantCategoryCode = *r.MerchantCategoryCode
	}
	return
}

func (r *PendingTransactionSourceCardRouteAuthorization) GetMerchantState() (MerchantState string) {
	if r != nil && r.MerchantState != nil {
		MerchantState = *r.MerchantState
	}
	return
}

func (r PendingTransactionSourceCardRouteAuthorization) String() (result string) {
	return fmt.Sprintf("&PendingTransactionSourceCardRouteAuthorization{Amount:%s Currency:%s MerchantAcceptorID:%s MerchantCity:%s MerchantCountry:%s MerchantDescriptor:%s MerchantCategoryCode:%s MerchantState:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.MerchantAcceptorID), core.FmtP(r.MerchantCity), core.FmtP(r.MerchantCountry), core.FmtP(r.MerchantDescriptor), core.FmtP(r.MerchantCategoryCode), core.FmtP(r.MerchantState))
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
	Amount             *int64                 `pjson:"amount"`
	AccountNumber      *string                `pjson:"account_number"`
	RoutingNumber      *string                `pjson:"routing_number"`
	MessageToRecipient *string                `pjson:"message_to_recipient"`
	jsonFields         map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// PendingTransactionSourceWireDrawdownPaymentInstruction using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *PendingTransactionSourceWireDrawdownPaymentInstruction) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes PendingTransactionSourceWireDrawdownPaymentInstruction
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *PendingTransactionSourceWireDrawdownPaymentInstruction) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *PendingTransactionSourceWireDrawdownPaymentInstruction) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *PendingTransactionSourceWireDrawdownPaymentInstruction) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

func (r *PendingTransactionSourceWireDrawdownPaymentInstruction) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

func (r *PendingTransactionSourceWireDrawdownPaymentInstruction) GetMessageToRecipient() (MessageToRecipient string) {
	if r != nil && r.MessageToRecipient != nil {
		MessageToRecipient = *r.MessageToRecipient
	}
	return
}

func (r PendingTransactionSourceWireDrawdownPaymentInstruction) String() (result string) {
	return fmt.Sprintf("&PendingTransactionSourceWireDrawdownPaymentInstruction{Amount:%s AccountNumber:%s RoutingNumber:%s MessageToRecipient:%s}", core.FmtP(r.Amount), core.FmtP(r.AccountNumber), core.FmtP(r.RoutingNumber), core.FmtP(r.MessageToRecipient))
}

type PendingTransactionSourceWireTransferInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount             *int64                 `pjson:"amount"`
	AccountNumber      *string                `pjson:"account_number"`
	RoutingNumber      *string                `pjson:"routing_number"`
	MessageToRecipient *string                `pjson:"message_to_recipient"`
	TransferID         *string                `pjson:"transfer_id"`
	jsonFields         map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// PendingTransactionSourceWireTransferInstruction using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *PendingTransactionSourceWireTransferInstruction) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes PendingTransactionSourceWireTransferInstruction into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *PendingTransactionSourceWireTransferInstruction) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *PendingTransactionSourceWireTransferInstruction) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *PendingTransactionSourceWireTransferInstruction) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

func (r *PendingTransactionSourceWireTransferInstruction) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

func (r *PendingTransactionSourceWireTransferInstruction) GetMessageToRecipient() (MessageToRecipient string) {
	if r != nil && r.MessageToRecipient != nil {
		MessageToRecipient = *r.MessageToRecipient
	}
	return
}

func (r *PendingTransactionSourceWireTransferInstruction) GetTransferID() (TransferID string) {
	if r != nil && r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

func (r PendingTransactionSourceWireTransferInstruction) String() (result string) {
	return fmt.Sprintf("&PendingTransactionSourceWireTransferInstruction{Amount:%s AccountNumber:%s RoutingNumber:%s MessageToRecipient:%s TransferID:%s}", core.FmtP(r.Amount), core.FmtP(r.AccountNumber), core.FmtP(r.RoutingNumber), core.FmtP(r.MessageToRecipient), core.FmtP(r.TransferID))
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
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int64 `query:"limit"`
	// Filter pending transactions to those belonging to the specified Account.
	AccountID *string `query:"account_id"`
	// Filter pending transactions to those belonging to the specified Route.
	RouteID *string `query:"route_id"`
	// Filter pending transactions to those caused by the specified source.
	SourceID   *string                              `query:"source_id"`
	Status     *PendingTransactionsListParamsStatus `query:"status"`
	jsonFields map[string]interface{}               `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into PendingTransactionListParams
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *PendingTransactionListParams) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes PendingTransactionListParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *PendingTransactionListParams) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes PendingTransactionListParams into a url.Values of the query
// parameters associated with this value
func (r *PendingTransactionListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// Return the page of entries after this one.
func (r *PendingTransactionListParams) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *PendingTransactionListParams) GetLimit() (Limit int64) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// Filter pending transactions to those belonging to the specified Account.
func (r *PendingTransactionListParams) GetAccountID() (AccountID string) {
	if r != nil && r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

// Filter pending transactions to those belonging to the specified Route.
func (r *PendingTransactionListParams) GetRouteID() (RouteID string) {
	if r != nil && r.RouteID != nil {
		RouteID = *r.RouteID
	}
	return
}

// Filter pending transactions to those caused by the specified source.
func (r *PendingTransactionListParams) GetSourceID() (SourceID string) {
	if r != nil && r.SourceID != nil {
		SourceID = *r.SourceID
	}
	return
}

func (r *PendingTransactionListParams) GetStatus() (Status PendingTransactionsListParamsStatus) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

func (r PendingTransactionListParams) String() (result string) {
	return fmt.Sprintf("&PendingTransactionListParams{Cursor:%s Limit:%s AccountID:%s RouteID:%s SourceID:%s Status:%s}", core.FmtP(r.Cursor), core.FmtP(r.Limit), core.FmtP(r.AccountID), core.FmtP(r.RouteID), core.FmtP(r.SourceID), r.Status)
}

type PendingTransactionsListParamsStatus struct {
	// Return results whose value is in the provided list. For GET requests, this
	// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
	In         *[]PendingTransactionsListParamsStatusIn `pjson:"in"`
	jsonFields map[string]interface{}                   `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// PendingTransactionsListParamsStatus using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *PendingTransactionsListParamsStatus) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes PendingTransactionsListParamsStatus into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *PendingTransactionsListParamsStatus) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes PendingTransactionsListParamsStatus into a url.Values of the
// query parameters associated with this value
func (r *PendingTransactionsListParamsStatus) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// Return results whose value is in the provided list. For GET requests, this
// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
func (r *PendingTransactionsListParamsStatus) GetIn() (In []PendingTransactionsListParamsStatusIn) {
	if r != nil && r.In != nil {
		In = *r.In
	}
	return
}

func (r PendingTransactionsListParamsStatus) String() (result string) {
	return fmt.Sprintf("&PendingTransactionsListParamsStatus{In:%s}", core.Fmt(r.In))
}

type PendingTransactionsListParamsStatusIn string

const (
	PendingTransactionsListParamsStatusInPending  PendingTransactionsListParamsStatusIn = "pending"
	PendingTransactionsListParamsStatusInComplete PendingTransactionsListParamsStatusIn = "complete"
)

type PendingTransactionList struct {
	// The contents of the list.
	Data *[]PendingTransaction `pjson:"data"`
	// A pointer to a place in the list.
	NextCursor *string                `pjson:"next_cursor"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into PendingTransactionList using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *PendingTransactionList) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes PendingTransactionList into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *PendingTransactionList) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes PendingTransactionList into a url.Values of the query
// parameters associated with this value
func (r *PendingTransactionList) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// The contents of the list.
func (r *PendingTransactionList) GetData() (Data []PendingTransaction) {
	if r != nil && r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r *PendingTransactionList) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

func (r PendingTransactionList) String() (result string) {
	return fmt.Sprintf("&PendingTransactionList{Data:%s NextCursor:%s}", core.Fmt(r.Data), core.FmtP(r.NextCursor))
}

type PendingTransactionsPage struct {
	*pagination.Page[PendingTransaction]
}

func (r *PendingTransactionsPage) PendingTransaction() *PendingTransaction {
	return r.Current()
}

func (r *PendingTransactionsPage) NextPage() (*PendingTransactionsPage, error) {
	if page, err := r.Page.NextPage(); err != nil {
		return nil, err
	} else {
		return &PendingTransactionsPage{page}, nil
	}
}

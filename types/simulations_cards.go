package types

import (
	"fmt"
	"increase/core"
	"increase/core/pjson"
)

type CardAuthorizationSimulation struct {
	// If the authorization attempt succeeds, this will contain the resulting Pending
	// Transaction object. The Pending Transaction's `source` will be of
	// `category: card_authorization`.
	PendingTransaction *CardAuthorizationSimulationPendingTransaction `pjson:"pending_transaction"`
	// If the authorization attempt fails, this will contain the resulting
	// [Declined Transaction](#declined-transactions) object. The Declined
	// Transaction's `source` will be of `category: card_decline`.
	DeclinedTransaction *CardAuthorizationSimulationDeclinedTransaction `pjson:"declined_transaction"`
	// A constant representing the object's type. For this resource it will always be
	// `inbound_card_authorization_simulation_result`.
	Type       *CardAuthorizationSimulationType `pjson:"type"`
	jsonFields map[string]interface{}           `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into CardAuthorizationSimulation
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *CardAuthorizationSimulation) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CardAuthorizationSimulation into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CardAuthorizationSimulation) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// If the authorization attempt succeeds, this will contain the resulting Pending
// Transaction object. The Pending Transaction's `source` will be of
// `category: card_authorization`.
func (r *CardAuthorizationSimulation) GetPendingTransaction() (PendingTransaction CardAuthorizationSimulationPendingTransaction) {
	if r != nil && r.PendingTransaction != nil {
		PendingTransaction = *r.PendingTransaction
	}
	return
}

// If the authorization attempt fails, this will contain the resulting
// [Declined Transaction](#declined-transactions) object. The Declined
// Transaction's `source` will be of `category: card_decline`.
func (r *CardAuthorizationSimulation) GetDeclinedTransaction() (DeclinedTransaction CardAuthorizationSimulationDeclinedTransaction) {
	if r != nil && r.DeclinedTransaction != nil {
		DeclinedTransaction = *r.DeclinedTransaction
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `inbound_card_authorization_simulation_result`.
func (r *CardAuthorizationSimulation) GetType() (Type CardAuthorizationSimulationType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

func (r CardAuthorizationSimulation) String() (result string) {
	return fmt.Sprintf("&CardAuthorizationSimulation{PendingTransaction:%s DeclinedTransaction:%s Type:%s}", r.PendingTransaction, r.DeclinedTransaction, core.FmtP(r.Type))
}

type CardAuthorizationSimulationPendingTransaction struct {
	// The identifier for the account this Pending Transaction belongs to.
	AccountID *string `pjson:"account_id"`
	// The Pending Transaction amount in the minor unit of its currency. For dollars,
	// for example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Pending
	// Transaction's currency. This will match the currency on the Pending
	// Transcation's Account.
	Currency *CardAuthorizationSimulationPendingTransactionCurrency `pjson:"currency"`
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
	Source *CardAuthorizationSimulationPendingTransactionSource `pjson:"source"`
	// Whether the Pending Transaction has been confirmed and has an associated
	// Transaction.
	Status *CardAuthorizationSimulationPendingTransactionStatus `pjson:"status"`
	// A constant representing the object's type. For this resource it will always be
	// `pending_transaction`.
	Type       *CardAuthorizationSimulationPendingTransactionType `pjson:"type"`
	jsonFields map[string]interface{}                             `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CardAuthorizationSimulationPendingTransaction using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *CardAuthorizationSimulationPendingTransaction) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CardAuthorizationSimulationPendingTransaction into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CardAuthorizationSimulationPendingTransaction) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The identifier for the account this Pending Transaction belongs to.
func (r *CardAuthorizationSimulationPendingTransaction) GetAccountID() (AccountID string) {
	if r != nil && r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

// The Pending Transaction amount in the minor unit of its currency. For dollars,
// for example, this is cents.
func (r *CardAuthorizationSimulationPendingTransaction) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Pending
// Transaction's currency. This will match the currency on the Pending
// Transcation's Account.
func (r *CardAuthorizationSimulationPendingTransaction) GetCurrency() (Currency CardAuthorizationSimulationPendingTransactionCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the Pending
// Transaction occured.
func (r *CardAuthorizationSimulationPendingTransaction) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// For a Pending Transaction related to a transfer, this is the description you
// provide. For a Pending Transaction related to a payment, this is the description
// the vendor provides.
func (r *CardAuthorizationSimulationPendingTransaction) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The Pending Transaction identifier.
func (r *CardAuthorizationSimulationPendingTransaction) GetID() (ID string) {
	if r != nil && r.ID != nil {
		ID = *r.ID
	}
	return
}

// The identifier for the route this Pending Transaction came through. Routes are
// things like cards and ACH details.
func (r *CardAuthorizationSimulationPendingTransaction) GetRouteID() (RouteID string) {
	if r != nil && r.RouteID != nil {
		RouteID = *r.RouteID
	}
	return
}

// The type of the route this Pending Transaction came through.
func (r *CardAuthorizationSimulationPendingTransaction) GetRouteType() (RouteType string) {
	if r != nil && r.RouteType != nil {
		RouteType = *r.RouteType
	}
	return
}

// This is an object giving more details on the network-level event that caused the
// Pending Transaction. For example, for a card transaction this lists the
// merchant's industry and location.
func (r *CardAuthorizationSimulationPendingTransaction) GetSource() (Source CardAuthorizationSimulationPendingTransactionSource) {
	if r != nil && r.Source != nil {
		Source = *r.Source
	}
	return
}

// Whether the Pending Transaction has been confirmed and has an associated
// Transaction.
func (r *CardAuthorizationSimulationPendingTransaction) GetStatus() (Status CardAuthorizationSimulationPendingTransactionStatus) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `pending_transaction`.
func (r *CardAuthorizationSimulationPendingTransaction) GetType() (Type CardAuthorizationSimulationPendingTransactionType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

func (r CardAuthorizationSimulationPendingTransaction) String() (result string) {
	return fmt.Sprintf("&CardAuthorizationSimulationPendingTransaction{AccountID:%s Amount:%s Currency:%s CreatedAt:%s Description:%s ID:%s RouteID:%s RouteType:%s Source:%s Status:%s Type:%s}", core.FmtP(r.AccountID), core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.CreatedAt), core.FmtP(r.Description), core.FmtP(r.ID), core.FmtP(r.RouteID), core.FmtP(r.RouteType), r.Source, core.FmtP(r.Status), core.FmtP(r.Type))
}

type CardAuthorizationSimulationPendingTransactionCurrency string

const (
	CardAuthorizationSimulationPendingTransactionCurrencyCad CardAuthorizationSimulationPendingTransactionCurrency = "CAD"
	CardAuthorizationSimulationPendingTransactionCurrencyChf CardAuthorizationSimulationPendingTransactionCurrency = "CHF"
	CardAuthorizationSimulationPendingTransactionCurrencyEur CardAuthorizationSimulationPendingTransactionCurrency = "EUR"
	CardAuthorizationSimulationPendingTransactionCurrencyGbp CardAuthorizationSimulationPendingTransactionCurrency = "GBP"
	CardAuthorizationSimulationPendingTransactionCurrencyJpy CardAuthorizationSimulationPendingTransactionCurrency = "JPY"
	CardAuthorizationSimulationPendingTransactionCurrencyUsd CardAuthorizationSimulationPendingTransactionCurrency = "USD"
)

type CardAuthorizationSimulationPendingTransactionSource struct {
	// The type of transaction that took place. We may add additional possible values
	// for this enum over time; your application should be able to handle such
	// additions gracefully.
	Category *CardAuthorizationSimulationPendingTransactionSourceCategory `pjson:"category"`
	// A Account Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `account_transfer_instruction`.
	AccountTransferInstruction *CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstruction `pjson:"account_transfer_instruction"`
	// A ACH Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `ach_transfer_instruction`.
	ACHTransferInstruction *CardAuthorizationSimulationPendingTransactionSourceACHTransferInstruction `pjson:"ach_transfer_instruction"`
	// A Card Authorization object. This field will be present in the JSON response if
	// and only if `category` is equal to `card_authorization`.
	CardAuthorization *CardAuthorizationSimulationPendingTransactionSourceCardAuthorization `pjson:"card_authorization"`
	// A Check Deposit Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_deposit_instruction`.
	CheckDepositInstruction *CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstruction `pjson:"check_deposit_instruction"`
	// A Check Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_instruction`.
	CheckTransferInstruction *CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstruction `pjson:"check_transfer_instruction"`
	// A Inbound Funds Hold object. This field will be present in the JSON response if
	// and only if `category` is equal to `inbound_funds_hold`.
	InboundFundsHold *CardAuthorizationSimulationPendingTransactionSourceInboundFundsHold `pjson:"inbound_funds_hold"`
	// A Deprecated Card Authorization object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_route_authorization`.
	CardRouteAuthorization *CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorization `pjson:"card_route_authorization"`
	// A Wire Drawdown Payment Instruction object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `wire_drawdown_payment_instruction`.
	WireDrawdownPaymentInstruction *CardAuthorizationSimulationPendingTransactionSourceWireDrawdownPaymentInstruction `pjson:"wire_drawdown_payment_instruction"`
	// A Wire Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_instruction`.
	WireTransferInstruction *CardAuthorizationSimulationPendingTransactionSourceWireTransferInstruction `pjson:"wire_transfer_instruction"`
	jsonFields              map[string]interface{}                                                      `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CardAuthorizationSimulationPendingTransactionSource using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *CardAuthorizationSimulationPendingTransactionSource) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CardAuthorizationSimulationPendingTransactionSource into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CardAuthorizationSimulationPendingTransactionSource) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The type of transaction that took place. We may add additional possible values
// for this enum over time; your application should be able to handle such
// additions gracefully.
func (r *CardAuthorizationSimulationPendingTransactionSource) GetCategory() (Category CardAuthorizationSimulationPendingTransactionSourceCategory) {
	if r != nil && r.Category != nil {
		Category = *r.Category
	}
	return
}

// A Account Transfer Instruction object. This field will be present in the JSON
// response if and only if `category` is equal to `account_transfer_instruction`.
func (r *CardAuthorizationSimulationPendingTransactionSource) GetAccountTransferInstruction() (AccountTransferInstruction CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstruction) {
	if r != nil && r.AccountTransferInstruction != nil {
		AccountTransferInstruction = *r.AccountTransferInstruction
	}
	return
}

// A ACH Transfer Instruction object. This field will be present in the JSON
// response if and only if `category` is equal to `ach_transfer_instruction`.
func (r *CardAuthorizationSimulationPendingTransactionSource) GetACHTransferInstruction() (ACHTransferInstruction CardAuthorizationSimulationPendingTransactionSourceACHTransferInstruction) {
	if r != nil && r.ACHTransferInstruction != nil {
		ACHTransferInstruction = *r.ACHTransferInstruction
	}
	return
}

// A Card Authorization object. This field will be present in the JSON response if
// and only if `category` is equal to `card_authorization`.
func (r *CardAuthorizationSimulationPendingTransactionSource) GetCardAuthorization() (CardAuthorization CardAuthorizationSimulationPendingTransactionSourceCardAuthorization) {
	if r != nil && r.CardAuthorization != nil {
		CardAuthorization = *r.CardAuthorization
	}
	return
}

// A Check Deposit Instruction object. This field will be present in the JSON
// response if and only if `category` is equal to `check_deposit_instruction`.
func (r *CardAuthorizationSimulationPendingTransactionSource) GetCheckDepositInstruction() (CheckDepositInstruction CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstruction) {
	if r != nil && r.CheckDepositInstruction != nil {
		CheckDepositInstruction = *r.CheckDepositInstruction
	}
	return
}

// A Check Transfer Instruction object. This field will be present in the JSON
// response if and only if `category` is equal to `check_transfer_instruction`.
func (r *CardAuthorizationSimulationPendingTransactionSource) GetCheckTransferInstruction() (CheckTransferInstruction CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstruction) {
	if r != nil && r.CheckTransferInstruction != nil {
		CheckTransferInstruction = *r.CheckTransferInstruction
	}
	return
}

// A Inbound Funds Hold object. This field will be present in the JSON response if
// and only if `category` is equal to `inbound_funds_hold`.
func (r *CardAuthorizationSimulationPendingTransactionSource) GetInboundFundsHold() (InboundFundsHold CardAuthorizationSimulationPendingTransactionSourceInboundFundsHold) {
	if r != nil && r.InboundFundsHold != nil {
		InboundFundsHold = *r.InboundFundsHold
	}
	return
}

// A Deprecated Card Authorization object. This field will be present in the JSON
// response if and only if `category` is equal to `card_route_authorization`.
func (r *CardAuthorizationSimulationPendingTransactionSource) GetCardRouteAuthorization() (CardRouteAuthorization CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorization) {
	if r != nil && r.CardRouteAuthorization != nil {
		CardRouteAuthorization = *r.CardRouteAuthorization
	}
	return
}

// A Wire Drawdown Payment Instruction object. This field will be present in the
// JSON response if and only if `category` is equal to
// `wire_drawdown_payment_instruction`.
func (r *CardAuthorizationSimulationPendingTransactionSource) GetWireDrawdownPaymentInstruction() (WireDrawdownPaymentInstruction CardAuthorizationSimulationPendingTransactionSourceWireDrawdownPaymentInstruction) {
	if r != nil && r.WireDrawdownPaymentInstruction != nil {
		WireDrawdownPaymentInstruction = *r.WireDrawdownPaymentInstruction
	}
	return
}

// A Wire Transfer Instruction object. This field will be present in the JSON
// response if and only if `category` is equal to `wire_transfer_instruction`.
func (r *CardAuthorizationSimulationPendingTransactionSource) GetWireTransferInstruction() (WireTransferInstruction CardAuthorizationSimulationPendingTransactionSourceWireTransferInstruction) {
	if r != nil && r.WireTransferInstruction != nil {
		WireTransferInstruction = *r.WireTransferInstruction
	}
	return
}

func (r CardAuthorizationSimulationPendingTransactionSource) String() (result string) {
	return fmt.Sprintf("&CardAuthorizationSimulationPendingTransactionSource{Category:%s AccountTransferInstruction:%s ACHTransferInstruction:%s CardAuthorization:%s CheckDepositInstruction:%s CheckTransferInstruction:%s InboundFundsHold:%s CardRouteAuthorization:%s WireDrawdownPaymentInstruction:%s WireTransferInstruction:%s}", core.FmtP(r.Category), r.AccountTransferInstruction, r.ACHTransferInstruction, r.CardAuthorization, r.CheckDepositInstruction, r.CheckTransferInstruction, r.InboundFundsHold, r.CardRouteAuthorization, r.WireDrawdownPaymentInstruction, r.WireTransferInstruction)
}

type CardAuthorizationSimulationPendingTransactionSourceCategory string

const (
	CardAuthorizationSimulationPendingTransactionSourceCategoryAccountTransferInstruction          CardAuthorizationSimulationPendingTransactionSourceCategory = "account_transfer_instruction"
	CardAuthorizationSimulationPendingTransactionSourceCategoryACHTransferInstruction              CardAuthorizationSimulationPendingTransactionSourceCategory = "ach_transfer_instruction"
	CardAuthorizationSimulationPendingTransactionSourceCategoryCardAuthorization                   CardAuthorizationSimulationPendingTransactionSourceCategory = "card_authorization"
	CardAuthorizationSimulationPendingTransactionSourceCategoryCheckDepositInstruction             CardAuthorizationSimulationPendingTransactionSourceCategory = "check_deposit_instruction"
	CardAuthorizationSimulationPendingTransactionSourceCategoryCheckTransferInstruction            CardAuthorizationSimulationPendingTransactionSourceCategory = "check_transfer_instruction"
	CardAuthorizationSimulationPendingTransactionSourceCategoryInboundFundsHold                    CardAuthorizationSimulationPendingTransactionSourceCategory = "inbound_funds_hold"
	CardAuthorizationSimulationPendingTransactionSourceCategoryCardRouteAuthorization              CardAuthorizationSimulationPendingTransactionSourceCategory = "card_route_authorization"
	CardAuthorizationSimulationPendingTransactionSourceCategoryRealTimePaymentsTransferInstruction CardAuthorizationSimulationPendingTransactionSourceCategory = "real_time_payments_transfer_instruction"
	CardAuthorizationSimulationPendingTransactionSourceCategoryWireDrawdownPaymentInstruction      CardAuthorizationSimulationPendingTransactionSourceCategory = "wire_drawdown_payment_instruction"
	CardAuthorizationSimulationPendingTransactionSourceCategoryWireTransferInstruction             CardAuthorizationSimulationPendingTransactionSourceCategory = "wire_transfer_instruction"
	CardAuthorizationSimulationPendingTransactionSourceCategoryOther                               CardAuthorizationSimulationPendingTransactionSourceCategory = "other"
)

type CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency *CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrency `pjson:"currency"`
	// The identifier of the Account Transfer that led to this Pending Transaction.
	TransferID *string                `pjson:"transfer_id"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstruction
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstruction) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstruction
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstruction) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstruction) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
// account currency.
func (r *CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstruction) GetCurrency() (Currency CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The identifier of the Account Transfer that led to this Pending Transaction.
func (r *CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstruction) GetTransferID() (TransferID string) {
	if r != nil && r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

func (r CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstruction) String() (result string) {
	return fmt.Sprintf("&CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstruction{Amount:%s Currency:%s TransferID:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.TransferID))
}

type CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrency string

const (
	CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrencyCad CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrency = "CAD"
	CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrencyChf CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrency = "CHF"
	CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrencyEur CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrency = "EUR"
	CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrencyGbp CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrency = "GBP"
	CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrencyJpy CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrency = "JPY"
	CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrencyUsd CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrency = "USD"
)

type CardAuthorizationSimulationPendingTransactionSourceACHTransferInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The identifier of the ACH Transfer that led to this Pending Transaction.
	TransferID *string                `pjson:"transfer_id"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CardAuthorizationSimulationPendingTransactionSourceACHTransferInstruction using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CardAuthorizationSimulationPendingTransactionSourceACHTransferInstruction) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// CardAuthorizationSimulationPendingTransactionSourceACHTransferInstruction into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CardAuthorizationSimulationPendingTransactionSourceACHTransferInstruction) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *CardAuthorizationSimulationPendingTransactionSourceACHTransferInstruction) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The identifier of the ACH Transfer that led to this Pending Transaction.
func (r *CardAuthorizationSimulationPendingTransactionSourceACHTransferInstruction) GetTransferID() (TransferID string) {
	if r != nil && r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

func (r CardAuthorizationSimulationPendingTransactionSourceACHTransferInstruction) String() (result string) {
	return fmt.Sprintf("&CardAuthorizationSimulationPendingTransactionSourceACHTransferInstruction{Amount:%s TransferID:%s}", core.FmtP(r.Amount), core.FmtP(r.TransferID))
}

type CardAuthorizationSimulationPendingTransactionSourceCardAuthorization struct {
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
	Network *CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetwork `pjson:"network"`
	// Fields specific to the `network`
	NetworkDetails *CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetails `pjson:"network_details"`
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrency `pjson:"currency"`
	// The identifier of the Real-Time Decision sent to approve or decline this
	// transaction.
	RealTimeDecisionID *string `pjson:"real_time_decision_id"`
	// If the authorization was made via a Digital Wallet Token (such as an Apple Pay
	// purchase), the identifier of the token that was used.
	DigitalWalletTokenID *string                `pjson:"digital_wallet_token_id"`
	jsonFields           map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CardAuthorizationSimulationPendingTransactionSourceCardAuthorization using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CardAuthorizationSimulationPendingTransactionSourceCardAuthorization) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// CardAuthorizationSimulationPendingTransactionSourceCardAuthorization into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CardAuthorizationSimulationPendingTransactionSourceCardAuthorization) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The merchant identifier (commonly abbreviated as MID) of the merchant the card
// is transacting with.
func (r *CardAuthorizationSimulationPendingTransactionSourceCardAuthorization) GetMerchantAcceptorID() (MerchantAcceptorID string) {
	if r != nil && r.MerchantAcceptorID != nil {
		MerchantAcceptorID = *r.MerchantAcceptorID
	}
	return
}

// The merchant descriptor of the merchant the card is transacting with.
func (r *CardAuthorizationSimulationPendingTransactionSourceCardAuthorization) GetMerchantDescriptor() (MerchantDescriptor string) {
	if r != nil && r.MerchantDescriptor != nil {
		MerchantDescriptor = *r.MerchantDescriptor
	}
	return
}

// The Merchant Category Code (commonly abbreviated as MCC) of the merchant the
// card is transacting with.
func (r *CardAuthorizationSimulationPendingTransactionSourceCardAuthorization) GetMerchantCategoryCode() (MerchantCategoryCode string) {
	if r != nil && r.MerchantCategoryCode != nil {
		MerchantCategoryCode = *r.MerchantCategoryCode
	}
	return
}

// The city the merchant resides in.
func (r *CardAuthorizationSimulationPendingTransactionSourceCardAuthorization) GetMerchantCity() (MerchantCity string) {
	if r != nil && r.MerchantCity != nil {
		MerchantCity = *r.MerchantCity
	}
	return
}

// The country the merchant resides in.
func (r *CardAuthorizationSimulationPendingTransactionSourceCardAuthorization) GetMerchantCountry() (MerchantCountry string) {
	if r != nil && r.MerchantCountry != nil {
		MerchantCountry = *r.MerchantCountry
	}
	return
}

// The payment network used to process this card authorization
func (r *CardAuthorizationSimulationPendingTransactionSourceCardAuthorization) GetNetwork() (Network CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetwork) {
	if r != nil && r.Network != nil {
		Network = *r.Network
	}
	return
}

// Fields specific to the `network`
func (r *CardAuthorizationSimulationPendingTransactionSourceCardAuthorization) GetNetworkDetails() (NetworkDetails CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetails) {
	if r != nil && r.NetworkDetails != nil {
		NetworkDetails = *r.NetworkDetails
	}
	return
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *CardAuthorizationSimulationPendingTransactionSourceCardAuthorization) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *CardAuthorizationSimulationPendingTransactionSourceCardAuthorization) GetCurrency() (Currency CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The identifier of the Real-Time Decision sent to approve or decline this
// transaction.
func (r *CardAuthorizationSimulationPendingTransactionSourceCardAuthorization) GetRealTimeDecisionID() (RealTimeDecisionID string) {
	if r != nil && r.RealTimeDecisionID != nil {
		RealTimeDecisionID = *r.RealTimeDecisionID
	}
	return
}

// If the authorization was made via a Digital Wallet Token (such as an Apple Pay
// purchase), the identifier of the token that was used.
func (r *CardAuthorizationSimulationPendingTransactionSourceCardAuthorization) GetDigitalWalletTokenID() (DigitalWalletTokenID string) {
	if r != nil && r.DigitalWalletTokenID != nil {
		DigitalWalletTokenID = *r.DigitalWalletTokenID
	}
	return
}

func (r CardAuthorizationSimulationPendingTransactionSourceCardAuthorization) String() (result string) {
	return fmt.Sprintf("&CardAuthorizationSimulationPendingTransactionSourceCardAuthorization{MerchantAcceptorID:%s MerchantDescriptor:%s MerchantCategoryCode:%s MerchantCity:%s MerchantCountry:%s Network:%s NetworkDetails:%s Amount:%s Currency:%s RealTimeDecisionID:%s DigitalWalletTokenID:%s}", core.FmtP(r.MerchantAcceptorID), core.FmtP(r.MerchantDescriptor), core.FmtP(r.MerchantCategoryCode), core.FmtP(r.MerchantCity), core.FmtP(r.MerchantCountry), core.FmtP(r.Network), r.NetworkDetails, core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.RealTimeDecisionID), core.FmtP(r.DigitalWalletTokenID))
}

type CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetwork string

const (
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkVisa CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetwork = "visa"
)

type CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetails struct {
	// Fields specific to the `visa` network
	Visa       *CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisa `pjson:"visa"`
	jsonFields map[string]interface{}                                                                  `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetails
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetails) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetails
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetails) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// Fields specific to the `visa` network
func (r *CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetails) GetVisa() (Visa CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisa) {
	if r != nil && r.Visa != nil {
		Visa = *r.Visa
	}
	return
}

func (r CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetails) String() (result string) {
	return fmt.Sprintf("&CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetails{Visa:%s}", r.Visa)
}

type CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisa struct {
	// For electronic commerce transactions, this identifies the level of security used
	// in obtaining the customer's payment credential. For mail or telephone order
	// transactions, identifies the type of mail or telephone order.
	ElectronicCommerceIndicator *CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator `pjson:"electronic_commerce_indicator"`
	// The method used to enter the cardholder's primary account number and card
	// expiration date
	PointOfServiceEntryMode *PointOfServiceEntryMode `pjson:"point_of_service_entry_mode"`
	jsonFields              map[string]interface{}   `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisa
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisa) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisa
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisa) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// For electronic commerce transactions, this identifies the level of security used
// in obtaining the customer's payment credential. For mail or telephone order
// transactions, identifies the type of mail or telephone order.
func (r *CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisa) GetElectronicCommerceIndicator() (ElectronicCommerceIndicator CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator) {
	if r != nil && r.ElectronicCommerceIndicator != nil {
		ElectronicCommerceIndicator = *r.ElectronicCommerceIndicator
	}
	return
}

// The method used to enter the cardholder's primary account number and card
// expiration date
func (r *CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisa) GetPointOfServiceEntryMode() (PointOfServiceEntryMode PointOfServiceEntryMode) {
	if r != nil && r.PointOfServiceEntryMode != nil {
		PointOfServiceEntryMode = *r.PointOfServiceEntryMode
	}
	return
}

func (r CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisa) String() (result string) {
	return fmt.Sprintf("&CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisa{ElectronicCommerceIndicator:%s PointOfServiceEntryMode:%s}", core.FmtP(r.ElectronicCommerceIndicator), core.FmtP(r.PointOfServiceEntryMode))
}

type CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator string

const (
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorMailPhoneOrder                                           CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "mail_phone_order"
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorRecurring                                                CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "recurring"
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorInstallment                                              CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "installment"
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorUnknownMailPhoneOrder                                    CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "unknown_mail_phone_order"
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorSecureElectronicCommerce                                 CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "secure_electronic_commerce"
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransactionAt_3dsCapableMerchant CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction_at_3ds_capable_merchant"
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransaction                      CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction"
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonSecureTransaction                                     CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "non_secure_transaction"
)

type CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrency string

const (
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrencyCad CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrency = "CAD"
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrencyChf CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrency = "CHF"
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrencyEur CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrency = "EUR"
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrencyGbp CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrency = "GBP"
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrencyJpy CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrency = "JPY"
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrencyUsd CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrency = "USD"
)

type CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrency `pjson:"currency"`
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
// CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstruction using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstruction) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstruction into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstruction) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstruction) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstruction) GetCurrency() (Currency CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The identifier of the File containing the image of the front of the check that
// was deposited.
func (r *CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstruction) GetFrontImageFileID() (FrontImageFileID string) {
	if r != nil && r.FrontImageFileID != nil {
		FrontImageFileID = *r.FrontImageFileID
	}
	return
}

// The identifier of the File containing the image of the back of the check that
// was deposited.
func (r *CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstruction) GetBackImageFileID() (BackImageFileID string) {
	if r != nil && r.BackImageFileID != nil {
		BackImageFileID = *r.BackImageFileID
	}
	return
}

// The identifier of the Check Deposit.
func (r *CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstruction) GetCheckDepositID() (CheckDepositID string) {
	if r != nil && r.CheckDepositID != nil {
		CheckDepositID = *r.CheckDepositID
	}
	return
}

func (r CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstruction) String() (result string) {
	return fmt.Sprintf("&CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstruction{Amount:%s Currency:%s FrontImageFileID:%s BackImageFileID:%s CheckDepositID:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.FrontImageFileID), core.FmtP(r.BackImageFileID), core.FmtP(r.CheckDepositID))
}

type CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrency string

const (
	CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrencyCad CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrency = "CAD"
	CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrencyChf CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrency = "CHF"
	CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrencyEur CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrency = "EUR"
	CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrencyGbp CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrency = "GBP"
	CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrencyJpy CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrency = "JPY"
	CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrencyUsd CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrency = "USD"
)

type CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
	// currency.
	Currency *CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrency `pjson:"currency"`
	// The identifier of the Check Transfer that led to this Pending Transaction.
	TransferID *string                `pjson:"transfer_id"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstruction
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstruction) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstruction into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstruction) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstruction) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
// currency.
func (r *CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstruction) GetCurrency() (Currency CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The identifier of the Check Transfer that led to this Pending Transaction.
func (r *CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstruction) GetTransferID() (TransferID string) {
	if r != nil && r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

func (r CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstruction) String() (result string) {
	return fmt.Sprintf("&CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstruction{Amount:%s Currency:%s TransferID:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.TransferID))
}

type CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrency string

const (
	CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrencyCad CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrency = "CAD"
	CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrencyChf CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrency = "CHF"
	CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrencyEur CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrency = "EUR"
	CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrencyGbp CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrency = "GBP"
	CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrencyJpy CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrency = "JPY"
	CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrencyUsd CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrency = "USD"
)

type CardAuthorizationSimulationPendingTransactionSourceInboundFundsHold struct {
	// The held amount in the minor unit of the account's currency. For dollars, for
	// example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the hold's
	// currency.
	Currency *CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrency `pjson:"currency"`
	// When the hold will be released automatically. Certain conditions may cause it to
	// be released before this time.
	AutomaticallyReleasesAt *string `pjson:"automatically_releases_at"`
	// When the hold was released (if it has been released).
	ReleasedAt *string `pjson:"released_at"`
	// The status of the hold.
	Status *CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldStatus `pjson:"status"`
	// The ID of the Transaction for which funds were held.
	HeldTransactionID *string                `pjson:"held_transaction_id"`
	jsonFields        map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CardAuthorizationSimulationPendingTransactionSourceInboundFundsHold using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CardAuthorizationSimulationPendingTransactionSourceInboundFundsHold) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// CardAuthorizationSimulationPendingTransactionSourceInboundFundsHold into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CardAuthorizationSimulationPendingTransactionSourceInboundFundsHold) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The held amount in the minor unit of the account's currency. For dollars, for
// example, this is cents.
func (r *CardAuthorizationSimulationPendingTransactionSourceInboundFundsHold) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the hold's
// currency.
func (r *CardAuthorizationSimulationPendingTransactionSourceInboundFundsHold) GetCurrency() (Currency CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// When the hold will be released automatically. Certain conditions may cause it to
// be released before this time.
func (r *CardAuthorizationSimulationPendingTransactionSourceInboundFundsHold) GetAutomaticallyReleasesAt() (AutomaticallyReleasesAt string) {
	if r != nil && r.AutomaticallyReleasesAt != nil {
		AutomaticallyReleasesAt = *r.AutomaticallyReleasesAt
	}
	return
}

// When the hold was released (if it has been released).
func (r *CardAuthorizationSimulationPendingTransactionSourceInboundFundsHold) GetReleasedAt() (ReleasedAt string) {
	if r != nil && r.ReleasedAt != nil {
		ReleasedAt = *r.ReleasedAt
	}
	return
}

// The status of the hold.
func (r *CardAuthorizationSimulationPendingTransactionSourceInboundFundsHold) GetStatus() (Status CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldStatus) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

// The ID of the Transaction for which funds were held.
func (r *CardAuthorizationSimulationPendingTransactionSourceInboundFundsHold) GetHeldTransactionID() (HeldTransactionID string) {
	if r != nil && r.HeldTransactionID != nil {
		HeldTransactionID = *r.HeldTransactionID
	}
	return
}

func (r CardAuthorizationSimulationPendingTransactionSourceInboundFundsHold) String() (result string) {
	return fmt.Sprintf("&CardAuthorizationSimulationPendingTransactionSourceInboundFundsHold{Amount:%s Currency:%s AutomaticallyReleasesAt:%s ReleasedAt:%s Status:%s HeldTransactionID:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.AutomaticallyReleasesAt), core.FmtP(r.ReleasedAt), core.FmtP(r.Status), core.FmtP(r.HeldTransactionID))
}

type CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrency string

const (
	CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrencyCad CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrency = "CAD"
	CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrencyChf CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrency = "CHF"
	CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrencyEur CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrency = "EUR"
	CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrencyGbp CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrency = "GBP"
	CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrencyJpy CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrency = "JPY"
	CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrencyUsd CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldCurrency = "USD"
)

type CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldStatus string

const (
	CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldStatusHeld     CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldStatus = "held"
	CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldStatusComplete CardAuthorizationSimulationPendingTransactionSourceInboundFundsHoldStatus = "complete"
)

type CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorization struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency             *CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationCurrency `pjson:"currency"`
	MerchantAcceptorID   *string                                                                            `pjson:"merchant_acceptor_id"`
	MerchantCity         *string                                                                            `pjson:"merchant_city"`
	MerchantCountry      *string                                                                            `pjson:"merchant_country"`
	MerchantDescriptor   *string                                                                            `pjson:"merchant_descriptor"`
	MerchantCategoryCode *string                                                                            `pjson:"merchant_category_code"`
	MerchantState        *string                                                                            `pjson:"merchant_state"`
	jsonFields           map[string]interface{}                                                             `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorization using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorization) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorization into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorization) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorization) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorization) GetCurrency() (Currency CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

func (r *CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorization) GetMerchantAcceptorID() (MerchantAcceptorID string) {
	if r != nil && r.MerchantAcceptorID != nil {
		MerchantAcceptorID = *r.MerchantAcceptorID
	}
	return
}

func (r *CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorization) GetMerchantCity() (MerchantCity string) {
	if r != nil && r.MerchantCity != nil {
		MerchantCity = *r.MerchantCity
	}
	return
}

func (r *CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorization) GetMerchantCountry() (MerchantCountry string) {
	if r != nil && r.MerchantCountry != nil {
		MerchantCountry = *r.MerchantCountry
	}
	return
}

func (r *CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorization) GetMerchantDescriptor() (MerchantDescriptor string) {
	if r != nil && r.MerchantDescriptor != nil {
		MerchantDescriptor = *r.MerchantDescriptor
	}
	return
}

func (r *CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorization) GetMerchantCategoryCode() (MerchantCategoryCode string) {
	if r != nil && r.MerchantCategoryCode != nil {
		MerchantCategoryCode = *r.MerchantCategoryCode
	}
	return
}

func (r *CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorization) GetMerchantState() (MerchantState string) {
	if r != nil && r.MerchantState != nil {
		MerchantState = *r.MerchantState
	}
	return
}

func (r CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorization) String() (result string) {
	return fmt.Sprintf("&CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorization{Amount:%s Currency:%s MerchantAcceptorID:%s MerchantCity:%s MerchantCountry:%s MerchantDescriptor:%s MerchantCategoryCode:%s MerchantState:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.MerchantAcceptorID), core.FmtP(r.MerchantCity), core.FmtP(r.MerchantCountry), core.FmtP(r.MerchantDescriptor), core.FmtP(r.MerchantCategoryCode), core.FmtP(r.MerchantState))
}

type CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationCurrency string

const (
	CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationCurrencyCad CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationCurrency = "CAD"
	CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationCurrencyChf CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationCurrency = "CHF"
	CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationCurrencyEur CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationCurrency = "EUR"
	CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationCurrencyGbp CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationCurrency = "GBP"
	CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationCurrencyJpy CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationCurrency = "JPY"
	CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationCurrencyUsd CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationCurrency = "USD"
)

type CardAuthorizationSimulationPendingTransactionSourceWireDrawdownPaymentInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount             *int64                 `pjson:"amount"`
	AccountNumber      *string                `pjson:"account_number"`
	RoutingNumber      *string                `pjson:"routing_number"`
	MessageToRecipient *string                `pjson:"message_to_recipient"`
	jsonFields         map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CardAuthorizationSimulationPendingTransactionSourceWireDrawdownPaymentInstruction
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *CardAuthorizationSimulationPendingTransactionSourceWireDrawdownPaymentInstruction) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// CardAuthorizationSimulationPendingTransactionSourceWireDrawdownPaymentInstruction
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *CardAuthorizationSimulationPendingTransactionSourceWireDrawdownPaymentInstruction) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *CardAuthorizationSimulationPendingTransactionSourceWireDrawdownPaymentInstruction) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *CardAuthorizationSimulationPendingTransactionSourceWireDrawdownPaymentInstruction) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

func (r *CardAuthorizationSimulationPendingTransactionSourceWireDrawdownPaymentInstruction) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

func (r *CardAuthorizationSimulationPendingTransactionSourceWireDrawdownPaymentInstruction) GetMessageToRecipient() (MessageToRecipient string) {
	if r != nil && r.MessageToRecipient != nil {
		MessageToRecipient = *r.MessageToRecipient
	}
	return
}

func (r CardAuthorizationSimulationPendingTransactionSourceWireDrawdownPaymentInstruction) String() (result string) {
	return fmt.Sprintf("&CardAuthorizationSimulationPendingTransactionSourceWireDrawdownPaymentInstruction{Amount:%s AccountNumber:%s RoutingNumber:%s MessageToRecipient:%s}", core.FmtP(r.Amount), core.FmtP(r.AccountNumber), core.FmtP(r.RoutingNumber), core.FmtP(r.MessageToRecipient))
}

type CardAuthorizationSimulationPendingTransactionSourceWireTransferInstruction struct {
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
// CardAuthorizationSimulationPendingTransactionSourceWireTransferInstruction using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CardAuthorizationSimulationPendingTransactionSourceWireTransferInstruction) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// CardAuthorizationSimulationPendingTransactionSourceWireTransferInstruction into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CardAuthorizationSimulationPendingTransactionSourceWireTransferInstruction) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *CardAuthorizationSimulationPendingTransactionSourceWireTransferInstruction) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *CardAuthorizationSimulationPendingTransactionSourceWireTransferInstruction) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

func (r *CardAuthorizationSimulationPendingTransactionSourceWireTransferInstruction) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

func (r *CardAuthorizationSimulationPendingTransactionSourceWireTransferInstruction) GetMessageToRecipient() (MessageToRecipient string) {
	if r != nil && r.MessageToRecipient != nil {
		MessageToRecipient = *r.MessageToRecipient
	}
	return
}

func (r *CardAuthorizationSimulationPendingTransactionSourceWireTransferInstruction) GetTransferID() (TransferID string) {
	if r != nil && r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

func (r CardAuthorizationSimulationPendingTransactionSourceWireTransferInstruction) String() (result string) {
	return fmt.Sprintf("&CardAuthorizationSimulationPendingTransactionSourceWireTransferInstruction{Amount:%s AccountNumber:%s RoutingNumber:%s MessageToRecipient:%s TransferID:%s}", core.FmtP(r.Amount), core.FmtP(r.AccountNumber), core.FmtP(r.RoutingNumber), core.FmtP(r.MessageToRecipient), core.FmtP(r.TransferID))
}

type CardAuthorizationSimulationPendingTransactionStatus string

const (
	CardAuthorizationSimulationPendingTransactionStatusPending  CardAuthorizationSimulationPendingTransactionStatus = "pending"
	CardAuthorizationSimulationPendingTransactionStatusComplete CardAuthorizationSimulationPendingTransactionStatus = "complete"
)

type CardAuthorizationSimulationPendingTransactionType string

const (
	CardAuthorizationSimulationPendingTransactionTypePendingTransaction CardAuthorizationSimulationPendingTransactionType = "pending_transaction"
)

type CardAuthorizationSimulationDeclinedTransaction struct {
	// The identifier for the Account the Declined Transaction belongs to.
	AccountID *string `pjson:"account_id"`
	// The Declined Transaction amount in the minor unit of its currency. For dollars,
	// for example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Declined
	// Transaction's currency. This will match the currency on the Declined
	// Transcation's Account.
	Currency *CardAuthorizationSimulationDeclinedTransactionCurrency `pjson:"currency"`
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
	Source *CardAuthorizationSimulationDeclinedTransactionSource `pjson:"source"`
	// A constant representing the object's type. For this resource it will always be
	// `declined_transaction`.
	Type       *CardAuthorizationSimulationDeclinedTransactionType `pjson:"type"`
	jsonFields map[string]interface{}                              `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CardAuthorizationSimulationDeclinedTransaction using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *CardAuthorizationSimulationDeclinedTransaction) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CardAuthorizationSimulationDeclinedTransaction into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CardAuthorizationSimulationDeclinedTransaction) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The identifier for the Account the Declined Transaction belongs to.
func (r *CardAuthorizationSimulationDeclinedTransaction) GetAccountID() (AccountID string) {
	if r != nil && r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

// The Declined Transaction amount in the minor unit of its currency. For dollars,
// for example, this is cents.
func (r *CardAuthorizationSimulationDeclinedTransaction) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Declined
// Transaction's currency. This will match the currency on the Declined
// Transcation's Account.
func (r *CardAuthorizationSimulationDeclinedTransaction) GetCurrency() (Currency CardAuthorizationSimulationDeclinedTransactionCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the
// Transaction occured.
func (r *CardAuthorizationSimulationDeclinedTransaction) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// This is the description the vendor provides.
func (r *CardAuthorizationSimulationDeclinedTransaction) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The Declined Transaction identifier.
func (r *CardAuthorizationSimulationDeclinedTransaction) GetID() (ID string) {
	if r != nil && r.ID != nil {
		ID = *r.ID
	}
	return
}

// The identifier for the route this Declined Transaction came through. Routes are
// things like cards and ACH details.
func (r *CardAuthorizationSimulationDeclinedTransaction) GetRouteID() (RouteID string) {
	if r != nil && r.RouteID != nil {
		RouteID = *r.RouteID
	}
	return
}

// The type of the route this Declined Transaction came through.
func (r *CardAuthorizationSimulationDeclinedTransaction) GetRouteType() (RouteType string) {
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
func (r *CardAuthorizationSimulationDeclinedTransaction) GetSource() (Source CardAuthorizationSimulationDeclinedTransactionSource) {
	if r != nil && r.Source != nil {
		Source = *r.Source
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `declined_transaction`.
func (r *CardAuthorizationSimulationDeclinedTransaction) GetType() (Type CardAuthorizationSimulationDeclinedTransactionType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

func (r CardAuthorizationSimulationDeclinedTransaction) String() (result string) {
	return fmt.Sprintf("&CardAuthorizationSimulationDeclinedTransaction{AccountID:%s Amount:%s Currency:%s CreatedAt:%s Description:%s ID:%s RouteID:%s RouteType:%s Source:%s Type:%s}", core.FmtP(r.AccountID), core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.CreatedAt), core.FmtP(r.Description), core.FmtP(r.ID), core.FmtP(r.RouteID), core.FmtP(r.RouteType), r.Source, core.FmtP(r.Type))
}

type CardAuthorizationSimulationDeclinedTransactionCurrency string

const (
	CardAuthorizationSimulationDeclinedTransactionCurrencyCad CardAuthorizationSimulationDeclinedTransactionCurrency = "CAD"
	CardAuthorizationSimulationDeclinedTransactionCurrencyChf CardAuthorizationSimulationDeclinedTransactionCurrency = "CHF"
	CardAuthorizationSimulationDeclinedTransactionCurrencyEur CardAuthorizationSimulationDeclinedTransactionCurrency = "EUR"
	CardAuthorizationSimulationDeclinedTransactionCurrencyGbp CardAuthorizationSimulationDeclinedTransactionCurrency = "GBP"
	CardAuthorizationSimulationDeclinedTransactionCurrencyJpy CardAuthorizationSimulationDeclinedTransactionCurrency = "JPY"
	CardAuthorizationSimulationDeclinedTransactionCurrencyUsd CardAuthorizationSimulationDeclinedTransactionCurrency = "USD"
)

type CardAuthorizationSimulationDeclinedTransactionSource struct {
	// The type of decline that took place. We may add additional possible values for
	// this enum over time; your application should be able to handle such additions
	// gracefully.
	Category *CardAuthorizationSimulationDeclinedTransactionSourceCategory `pjson:"category"`
	// A ACH Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `ach_decline`.
	ACHDecline *CardAuthorizationSimulationDeclinedTransactionSourceACHDecline `pjson:"ach_decline"`
	// A Card Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_decline`.
	CardDecline *CardAuthorizationSimulationDeclinedTransactionSourceCardDecline `pjson:"card_decline"`
	// A Check Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `check_decline`.
	CheckDecline *CardAuthorizationSimulationDeclinedTransactionSourceCheckDecline `pjson:"check_decline"`
	// A Inbound Real Time Payments Transfer Decline object. This field will be present
	// in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_decline`.
	InboundRealTimePaymentsTransferDecline *CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline `pjson:"inbound_real_time_payments_transfer_decline"`
	// A International ACH Decline object. This field will be present in the JSON
	// response if and only if `category` is equal to `international_ach_decline`.
	InternationalACHDecline *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline `pjson:"international_ach_decline"`
	// A Deprecated Card Decline object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_route_decline`.
	CardRouteDecline *CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDecline `pjson:"card_route_decline"`
	jsonFields       map[string]interface{}                                                `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CardAuthorizationSimulationDeclinedTransactionSource using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *CardAuthorizationSimulationDeclinedTransactionSource) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CardAuthorizationSimulationDeclinedTransactionSource into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CardAuthorizationSimulationDeclinedTransactionSource) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The type of decline that took place. We may add additional possible values for
// this enum over time; your application should be able to handle such additions
// gracefully.
func (r *CardAuthorizationSimulationDeclinedTransactionSource) GetCategory() (Category CardAuthorizationSimulationDeclinedTransactionSourceCategory) {
	if r != nil && r.Category != nil {
		Category = *r.Category
	}
	return
}

// A ACH Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `ach_decline`.
func (r *CardAuthorizationSimulationDeclinedTransactionSource) GetACHDecline() (ACHDecline CardAuthorizationSimulationDeclinedTransactionSourceACHDecline) {
	if r != nil && r.ACHDecline != nil {
		ACHDecline = *r.ACHDecline
	}
	return
}

// A Card Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `card_decline`.
func (r *CardAuthorizationSimulationDeclinedTransactionSource) GetCardDecline() (CardDecline CardAuthorizationSimulationDeclinedTransactionSourceCardDecline) {
	if r != nil && r.CardDecline != nil {
		CardDecline = *r.CardDecline
	}
	return
}

// A Check Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `check_decline`.
func (r *CardAuthorizationSimulationDeclinedTransactionSource) GetCheckDecline() (CheckDecline CardAuthorizationSimulationDeclinedTransactionSourceCheckDecline) {
	if r != nil && r.CheckDecline != nil {
		CheckDecline = *r.CheckDecline
	}
	return
}

// A Inbound Real Time Payments Transfer Decline object. This field will be present
// in the JSON response if and only if `category` is equal to
// `inbound_real_time_payments_transfer_decline`.
func (r *CardAuthorizationSimulationDeclinedTransactionSource) GetInboundRealTimePaymentsTransferDecline() (InboundRealTimePaymentsTransferDecline CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) {
	if r != nil && r.InboundRealTimePaymentsTransferDecline != nil {
		InboundRealTimePaymentsTransferDecline = *r.InboundRealTimePaymentsTransferDecline
	}
	return
}

// A International ACH Decline object. This field will be present in the JSON
// response if and only if `category` is equal to `international_ach_decline`.
func (r *CardAuthorizationSimulationDeclinedTransactionSource) GetInternationalACHDecline() (InternationalACHDecline CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) {
	if r != nil && r.InternationalACHDecline != nil {
		InternationalACHDecline = *r.InternationalACHDecline
	}
	return
}

// A Deprecated Card Decline object. This field will be present in the JSON
// response if and only if `category` is equal to `card_route_decline`.
func (r *CardAuthorizationSimulationDeclinedTransactionSource) GetCardRouteDecline() (CardRouteDecline CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDecline) {
	if r != nil && r.CardRouteDecline != nil {
		CardRouteDecline = *r.CardRouteDecline
	}
	return
}

func (r CardAuthorizationSimulationDeclinedTransactionSource) String() (result string) {
	return fmt.Sprintf("&CardAuthorizationSimulationDeclinedTransactionSource{Category:%s ACHDecline:%s CardDecline:%s CheckDecline:%s InboundRealTimePaymentsTransferDecline:%s InternationalACHDecline:%s CardRouteDecline:%s}", core.FmtP(r.Category), r.ACHDecline, r.CardDecline, r.CheckDecline, r.InboundRealTimePaymentsTransferDecline, r.InternationalACHDecline, r.CardRouteDecline)
}

type CardAuthorizationSimulationDeclinedTransactionSourceCategory string

const (
	CardAuthorizationSimulationDeclinedTransactionSourceCategoryACHDecline                             CardAuthorizationSimulationDeclinedTransactionSourceCategory = "ach_decline"
	CardAuthorizationSimulationDeclinedTransactionSourceCategoryCardDecline                            CardAuthorizationSimulationDeclinedTransactionSourceCategory = "card_decline"
	CardAuthorizationSimulationDeclinedTransactionSourceCategoryCheckDecline                           CardAuthorizationSimulationDeclinedTransactionSourceCategory = "check_decline"
	CardAuthorizationSimulationDeclinedTransactionSourceCategoryInboundRealTimePaymentsTransferDecline CardAuthorizationSimulationDeclinedTransactionSourceCategory = "inbound_real_time_payments_transfer_decline"
	CardAuthorizationSimulationDeclinedTransactionSourceCategoryInternationalACHDecline                CardAuthorizationSimulationDeclinedTransactionSourceCategory = "international_ach_decline"
	CardAuthorizationSimulationDeclinedTransactionSourceCategoryCardRouteDecline                       CardAuthorizationSimulationDeclinedTransactionSourceCategory = "card_route_decline"
	CardAuthorizationSimulationDeclinedTransactionSourceCategoryOther                                  CardAuthorizationSimulationDeclinedTransactionSourceCategory = "other"
)

type CardAuthorizationSimulationDeclinedTransactionSourceACHDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount                             *int64  `pjson:"amount"`
	OriginatorCompanyName              *string `pjson:"originator_company_name"`
	OriginatorCompanyDescriptiveDate   *string `pjson:"originator_company_descriptive_date"`
	OriginatorCompanyDiscretionaryData *string `pjson:"originator_company_discretionary_data"`
	OriginatorCompanyID                *string `pjson:"originator_company_id"`
	// Why the ACH transfer was declined.
	Reason           *CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason `pjson:"reason"`
	ReceiverIDNumber *string                                                               `pjson:"receiver_id_number"`
	ReceiverName     *string                                                               `pjson:"receiver_name"`
	TraceNumber      *string                                                               `pjson:"trace_number"`
	jsonFields       map[string]interface{}                                                `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CardAuthorizationSimulationDeclinedTransactionSourceACHDecline using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceACHDecline) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// CardAuthorizationSimulationDeclinedTransactionSourceACHDecline into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceACHDecline) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceACHDecline) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceACHDecline) GetOriginatorCompanyName() (OriginatorCompanyName string) {
	if r != nil && r.OriginatorCompanyName != nil {
		OriginatorCompanyName = *r.OriginatorCompanyName
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceACHDecline) GetOriginatorCompanyDescriptiveDate() (OriginatorCompanyDescriptiveDate string) {
	if r != nil && r.OriginatorCompanyDescriptiveDate != nil {
		OriginatorCompanyDescriptiveDate = *r.OriginatorCompanyDescriptiveDate
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceACHDecline) GetOriginatorCompanyDiscretionaryData() (OriginatorCompanyDiscretionaryData string) {
	if r != nil && r.OriginatorCompanyDiscretionaryData != nil {
		OriginatorCompanyDiscretionaryData = *r.OriginatorCompanyDiscretionaryData
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceACHDecline) GetOriginatorCompanyID() (OriginatorCompanyID string) {
	if r != nil && r.OriginatorCompanyID != nil {
		OriginatorCompanyID = *r.OriginatorCompanyID
	}
	return
}

// Why the ACH transfer was declined.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceACHDecline) GetReason() (Reason CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason) {
	if r != nil && r.Reason != nil {
		Reason = *r.Reason
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceACHDecline) GetReceiverIDNumber() (ReceiverIDNumber string) {
	if r != nil && r.ReceiverIDNumber != nil {
		ReceiverIDNumber = *r.ReceiverIDNumber
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceACHDecline) GetReceiverName() (ReceiverName string) {
	if r != nil && r.ReceiverName != nil {
		ReceiverName = *r.ReceiverName
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceACHDecline) GetTraceNumber() (TraceNumber string) {
	if r != nil && r.TraceNumber != nil {
		TraceNumber = *r.TraceNumber
	}
	return
}

func (r CardAuthorizationSimulationDeclinedTransactionSourceACHDecline) String() (result string) {
	return fmt.Sprintf("&CardAuthorizationSimulationDeclinedTransactionSourceACHDecline{Amount:%s OriginatorCompanyName:%s OriginatorCompanyDescriptiveDate:%s OriginatorCompanyDiscretionaryData:%s OriginatorCompanyID:%s Reason:%s ReceiverIDNumber:%s ReceiverName:%s TraceNumber:%s}", core.FmtP(r.Amount), core.FmtP(r.OriginatorCompanyName), core.FmtP(r.OriginatorCompanyDescriptiveDate), core.FmtP(r.OriginatorCompanyDiscretionaryData), core.FmtP(r.OriginatorCompanyID), core.FmtP(r.Reason), core.FmtP(r.ReceiverIDNumber), core.FmtP(r.ReceiverName), core.FmtP(r.TraceNumber))
}

type CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason string

const (
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonACHRouteCanceled             CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "ach_route_canceled"
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonACHRouteDisabled             CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "ach_route_disabled"
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonBreachesLimit                CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "breaches_limit"
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonCreditEntryRefusedByReceiver CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "credit_entry_refused_by_receiver"
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonDuplicateReturn              CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "duplicate_return"
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonEntityNotActive              CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "entity_not_active"
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonTransactionNotAllowed        CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "transaction_not_allowed"
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonGroupLocked                  CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "group_locked"
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonInsufficientFunds            CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "insufficient_funds"
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonNoACHRoute                   CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "no_ach_route"
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonOriginatorRequest            CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "originator_request"
)

type CardAuthorizationSimulationDeclinedTransactionSourceCardDecline struct {
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
	Network *CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetwork `pjson:"network"`
	// Fields specific to the `network`
	NetworkDetails *CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetails `pjson:"network_details"`
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency *CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrency `pjson:"currency"`
	// Why the transaction was declined.
	Reason *CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason `pjson:"reason"`
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
// CardAuthorizationSimulationDeclinedTransactionSourceCardDecline using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDecline) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// CardAuthorizationSimulationDeclinedTransactionSourceCardDecline into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDecline) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The merchant identifier (commonly abbreviated as MID) of the merchant the card
// is transacting with.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDecline) GetMerchantAcceptorID() (MerchantAcceptorID string) {
	if r != nil && r.MerchantAcceptorID != nil {
		MerchantAcceptorID = *r.MerchantAcceptorID
	}
	return
}

// The merchant descriptor of the merchant the card is transacting with.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDecline) GetMerchantDescriptor() (MerchantDescriptor string) {
	if r != nil && r.MerchantDescriptor != nil {
		MerchantDescriptor = *r.MerchantDescriptor
	}
	return
}

// The Merchant Category Code (commonly abbreviated as MCC) of the merchant the
// card is transacting with.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDecline) GetMerchantCategoryCode() (MerchantCategoryCode string) {
	if r != nil && r.MerchantCategoryCode != nil {
		MerchantCategoryCode = *r.MerchantCategoryCode
	}
	return
}

// The city the merchant resides in.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDecline) GetMerchantCity() (MerchantCity string) {
	if r != nil && r.MerchantCity != nil {
		MerchantCity = *r.MerchantCity
	}
	return
}

// The country the merchant resides in.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDecline) GetMerchantCountry() (MerchantCountry string) {
	if r != nil && r.MerchantCountry != nil {
		MerchantCountry = *r.MerchantCountry
	}
	return
}

// The payment network used to process this card authorization
func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDecline) GetNetwork() (Network CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetwork) {
	if r != nil && r.Network != nil {
		Network = *r.Network
	}
	return
}

// Fields specific to the `network`
func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDecline) GetNetworkDetails() (NetworkDetails CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetails) {
	if r != nil && r.NetworkDetails != nil {
		NetworkDetails = *r.NetworkDetails
	}
	return
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDecline) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
// account currency.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDecline) GetCurrency() (Currency CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// Why the transaction was declined.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDecline) GetReason() (Reason CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason) {
	if r != nil && r.Reason != nil {
		Reason = *r.Reason
	}
	return
}

// The state the merchant resides in.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDecline) GetMerchantState() (MerchantState string) {
	if r != nil && r.MerchantState != nil {
		MerchantState = *r.MerchantState
	}
	return
}

// The identifier of the Real-Time Decision sent to approve or decline this
// transaction.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDecline) GetRealTimeDecisionID() (RealTimeDecisionID string) {
	if r != nil && r.RealTimeDecisionID != nil {
		RealTimeDecisionID = *r.RealTimeDecisionID
	}
	return
}

// If the authorization was attempted using a Digital Wallet Token (such as an
// Apple Pay purchase), the identifier of the token that was used.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDecline) GetDigitalWalletTokenID() (DigitalWalletTokenID string) {
	if r != nil && r.DigitalWalletTokenID != nil {
		DigitalWalletTokenID = *r.DigitalWalletTokenID
	}
	return
}

func (r CardAuthorizationSimulationDeclinedTransactionSourceCardDecline) String() (result string) {
	return fmt.Sprintf("&CardAuthorizationSimulationDeclinedTransactionSourceCardDecline{MerchantAcceptorID:%s MerchantDescriptor:%s MerchantCategoryCode:%s MerchantCity:%s MerchantCountry:%s Network:%s NetworkDetails:%s Amount:%s Currency:%s Reason:%s MerchantState:%s RealTimeDecisionID:%s DigitalWalletTokenID:%s}", core.FmtP(r.MerchantAcceptorID), core.FmtP(r.MerchantDescriptor), core.FmtP(r.MerchantCategoryCode), core.FmtP(r.MerchantCity), core.FmtP(r.MerchantCountry), core.FmtP(r.Network), r.NetworkDetails, core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.Reason), core.FmtP(r.MerchantState), core.FmtP(r.RealTimeDecisionID), core.FmtP(r.DigitalWalletTokenID))
}

type CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetwork string

const (
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkVisa CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetwork = "visa"
)

type CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetails struct {
	// Fields specific to the `visa` network
	Visa       *CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa `pjson:"visa"`
	jsonFields map[string]interface{}                                                             `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetails
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetails) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetails
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetails) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// Fields specific to the `visa` network
func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetails) GetVisa() (Visa CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa) {
	if r != nil && r.Visa != nil {
		Visa = *r.Visa
	}
	return
}

func (r CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetails) String() (result string) {
	return fmt.Sprintf("&CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetails{Visa:%s}", r.Visa)
}

type CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa struct {
	// For electronic commerce transactions, this identifies the level of security used
	// in obtaining the customer's payment credential. For mail or telephone order
	// transactions, identifies the type of mail or telephone order.
	ElectronicCommerceIndicator *CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator `pjson:"electronic_commerce_indicator"`
	// The method used to enter the cardholder's primary account number and card
	// expiration date
	PointOfServiceEntryMode *PointOfServiceEntryMode `pjson:"point_of_service_entry_mode"`
	jsonFields              map[string]interface{}   `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// For electronic commerce transactions, this identifies the level of security used
// in obtaining the customer's payment credential. For mail or telephone order
// transactions, identifies the type of mail or telephone order.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa) GetElectronicCommerceIndicator() (ElectronicCommerceIndicator CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator) {
	if r != nil && r.ElectronicCommerceIndicator != nil {
		ElectronicCommerceIndicator = *r.ElectronicCommerceIndicator
	}
	return
}

// The method used to enter the cardholder's primary account number and card
// expiration date
func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa) GetPointOfServiceEntryMode() (PointOfServiceEntryMode PointOfServiceEntryMode) {
	if r != nil && r.PointOfServiceEntryMode != nil {
		PointOfServiceEntryMode = *r.PointOfServiceEntryMode
	}
	return
}

func (r CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa) String() (result string) {
	return fmt.Sprintf("&CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisa{ElectronicCommerceIndicator:%s PointOfServiceEntryMode:%s}", core.FmtP(r.ElectronicCommerceIndicator), core.FmtP(r.PointOfServiceEntryMode))
}

type CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator string

const (
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorMailPhoneOrder                                           CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "mail_phone_order"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorRecurring                                                CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "recurring"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorInstallment                                              CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "installment"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorUnknownMailPhoneOrder                                    CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "unknown_mail_phone_order"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorSecureElectronicCommerce                                 CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "secure_electronic_commerce"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransactionAt_3dsCapableMerchant CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction_at_3ds_capable_merchant"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransaction                      CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicatorNonSecureTransaction                                     CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineNetworkDetailsVisaElectronicCommerceIndicator = "non_secure_transaction"
)

type CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrency string

const (
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrencyCad CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrency = "CAD"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrencyChf CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrency = "CHF"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrencyEur CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrency = "EUR"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrencyGbp CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrency = "GBP"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrencyJpy CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrency = "JPY"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrencyUsd CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrency = "USD"
)

type CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason string

const (
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonCardNotActive               CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "card_not_active"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonEntityNotActive             CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "entity_not_active"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonGroupLocked                 CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "group_locked"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonInsufficientFunds           CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "insufficient_funds"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonCvv2Mismatch                CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "cvv2_mismatch"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonTransactionNotAllowed       CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "transaction_not_allowed"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonBreachesLimit               CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "breaches_limit"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonWebhookDeclined             CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "webhook_declined"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonWebhookTimedOut             CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "webhook_timed_out"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonDeclinedByStandInProcessing CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "declined_by_stand_in_processing"
)

type CardAuthorizationSimulationDeclinedTransactionSourceCheckDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount        *int64  `pjson:"amount"`
	AuxiliaryOnUs *string `pjson:"auxiliary_on_us"`
	// Why the check was declined.
	Reason     *CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason `pjson:"reason"`
	jsonFields map[string]interface{}                                                  `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CardAuthorizationSimulationDeclinedTransactionSourceCheckDecline using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceCheckDecline) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// CardAuthorizationSimulationDeclinedTransactionSourceCheckDecline into an array
// of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceCheckDecline) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceCheckDecline) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceCheckDecline) GetAuxiliaryOnUs() (AuxiliaryOnUs string) {
	if r != nil && r.AuxiliaryOnUs != nil {
		AuxiliaryOnUs = *r.AuxiliaryOnUs
	}
	return
}

// Why the check was declined.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceCheckDecline) GetReason() (Reason CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason) {
	if r != nil && r.Reason != nil {
		Reason = *r.Reason
	}
	return
}

func (r CardAuthorizationSimulationDeclinedTransactionSourceCheckDecline) String() (result string) {
	return fmt.Sprintf("&CardAuthorizationSimulationDeclinedTransactionSourceCheckDecline{Amount:%s AuxiliaryOnUs:%s Reason:%s}", core.FmtP(r.Amount), core.FmtP(r.AuxiliaryOnUs), core.FmtP(r.Reason))
}

type CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason string

const (
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonACHRouteCanceled      CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "ach_route_canceled"
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonACHRouteDisabled      CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "ach_route_disabled"
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonBreachesLimit         CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "breaches_limit"
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonEntityNotActive       CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "entity_not_active"
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonGroupLocked           CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "group_locked"
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonInsufficientFunds     CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "insufficient_funds"
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonUnableToLocateAccount CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "unable_to_locate_account"
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonUnableToProcess       CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "unable_to_process"
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonReferToImage          CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "refer_to_image"
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonStopPaymentRequested  CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "stop_payment_requested"
	CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReasonReturned              CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason = "returned"
)

type CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the declined
	// transfer's currency. This will always be "USD" for a Real Time Payments
	// transfer.
	Currency *CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency `pjson:"currency"`
	// Why the transfer was declined.
	Reason *CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason `pjson:"reason"`
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
// CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline
// into an array of bytes using the gjson library. Members of the `jsonFields`
// field are serialized into the top-level, and will overwrite known members of the
// same name.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the declined
// transfer's currency. This will always be "USD" for a Real Time Payments
// transfer.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) GetCurrency() (Currency CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// Why the transfer was declined.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) GetReason() (Reason CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason) {
	if r != nil && r.Reason != nil {
		Reason = *r.Reason
	}
	return
}

// The name the sender of the transfer specified as the recipient of the transfer.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) GetCreditorName() (CreditorName string) {
	if r != nil && r.CreditorName != nil {
		CreditorName = *r.CreditorName
	}
	return
}

// The name provided by the sender of the transfer.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) GetDebtorName() (DebtorName string) {
	if r != nil && r.DebtorName != nil {
		DebtorName = *r.DebtorName
	}
	return
}

// The account number of the account that sent the transfer.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) GetDebtorAccountNumber() (DebtorAccountNumber string) {
	if r != nil && r.DebtorAccountNumber != nil {
		DebtorAccountNumber = *r.DebtorAccountNumber
	}
	return
}

// The routing number of the account that sent the transfer.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) GetDebtorRoutingNumber() (DebtorRoutingNumber string) {
	if r != nil && r.DebtorRoutingNumber != nil {
		DebtorRoutingNumber = *r.DebtorRoutingNumber
	}
	return
}

// The Real Time Payments network identification of the declined transfer.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) GetTransactionIdentification() (TransactionIdentification string) {
	if r != nil && r.TransactionIdentification != nil {
		TransactionIdentification = *r.TransactionIdentification
	}
	return
}

// Additional information included with the transfer.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) GetRemittanceInformation() (RemittanceInformation string) {
	if r != nil && r.RemittanceInformation != nil {
		RemittanceInformation = *r.RemittanceInformation
	}
	return
}

func (r CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) String() (result string) {
	return fmt.Sprintf("&CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline{Amount:%s Currency:%s Reason:%s CreditorName:%s DebtorName:%s DebtorAccountNumber:%s DebtorRoutingNumber:%s TransactionIdentification:%s RemittanceInformation:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.Reason), core.FmtP(r.CreditorName), core.FmtP(r.DebtorName), core.FmtP(r.DebtorAccountNumber), core.FmtP(r.DebtorRoutingNumber), core.FmtP(r.TransactionIdentification), core.FmtP(r.RemittanceInformation))
}

type CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency string

const (
	CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyCad CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "CAD"
	CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyChf CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "CHF"
	CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyEur CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "EUR"
	CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyGbp CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "GBP"
	CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyJpy CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "JPY"
	CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyUsd CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "USD"
)

type CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason string

const (
	CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonAccountNumberCanceled      CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "account_number_canceled"
	CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonAccountNumberDisabled      CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "account_number_disabled"
	CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonGroupLocked                CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "group_locked"
	CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonEntityNotActive            CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "entity_not_active"
	CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonRealTimePaymentsNotEnabled CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "real_time_payments_not_enabled"
)

type CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline struct {
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
// CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline into
// an array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) GetForeignExchangeIndicator() (ForeignExchangeIndicator string) {
	if r != nil && r.ForeignExchangeIndicator != nil {
		ForeignExchangeIndicator = *r.ForeignExchangeIndicator
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) GetForeignExchangeReferenceIndicator() (ForeignExchangeReferenceIndicator string) {
	if r != nil && r.ForeignExchangeReferenceIndicator != nil {
		ForeignExchangeReferenceIndicator = *r.ForeignExchangeReferenceIndicator
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) GetForeignExchangeReference() (ForeignExchangeReference string) {
	if r != nil && r.ForeignExchangeReference != nil {
		ForeignExchangeReference = *r.ForeignExchangeReference
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) GetDestinationCountryCode() (DestinationCountryCode string) {
	if r != nil && r.DestinationCountryCode != nil {
		DestinationCountryCode = *r.DestinationCountryCode
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) GetDestinationCurrencyCode() (DestinationCurrencyCode string) {
	if r != nil && r.DestinationCurrencyCode != nil {
		DestinationCurrencyCode = *r.DestinationCurrencyCode
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) GetForeignPaymentAmount() (ForeignPaymentAmount int64) {
	if r != nil && r.ForeignPaymentAmount != nil {
		ForeignPaymentAmount = *r.ForeignPaymentAmount
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) GetForeignTraceNumber() (ForeignTraceNumber string) {
	if r != nil && r.ForeignTraceNumber != nil {
		ForeignTraceNumber = *r.ForeignTraceNumber
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) GetInternationalTransactionTypeCode() (InternationalTransactionTypeCode string) {
	if r != nil && r.InternationalTransactionTypeCode != nil {
		InternationalTransactionTypeCode = *r.InternationalTransactionTypeCode
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) GetOriginatingCurrencyCode() (OriginatingCurrencyCode string) {
	if r != nil && r.OriginatingCurrencyCode != nil {
		OriginatingCurrencyCode = *r.OriginatingCurrencyCode
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) GetOriginatingDepositoryFinancialInstitutionName() (OriginatingDepositoryFinancialInstitutionName string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionName != nil {
		OriginatingDepositoryFinancialInstitutionName = *r.OriginatingDepositoryFinancialInstitutionName
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) GetOriginatingDepositoryFinancialInstitutionIDQualifier() (OriginatingDepositoryFinancialInstitutionIDQualifier string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionIDQualifier != nil {
		OriginatingDepositoryFinancialInstitutionIDQualifier = *r.OriginatingDepositoryFinancialInstitutionIDQualifier
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) GetOriginatingDepositoryFinancialInstitutionID() (OriginatingDepositoryFinancialInstitutionID string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionID != nil {
		OriginatingDepositoryFinancialInstitutionID = *r.OriginatingDepositoryFinancialInstitutionID
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) GetOriginatingDepositoryFinancialInstitutionBranchCountry() (OriginatingDepositoryFinancialInstitutionBranchCountry string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionBranchCountry != nil {
		OriginatingDepositoryFinancialInstitutionBranchCountry = *r.OriginatingDepositoryFinancialInstitutionBranchCountry
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) GetOriginatorCity() (OriginatorCity string) {
	if r != nil && r.OriginatorCity != nil {
		OriginatorCity = *r.OriginatorCity
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) GetOriginatorCompanyEntryDescription() (OriginatorCompanyEntryDescription string) {
	if r != nil && r.OriginatorCompanyEntryDescription != nil {
		OriginatorCompanyEntryDescription = *r.OriginatorCompanyEntryDescription
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) GetOriginatorCountry() (OriginatorCountry string) {
	if r != nil && r.OriginatorCountry != nil {
		OriginatorCountry = *r.OriginatorCountry
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) GetOriginatorIdentification() (OriginatorIdentification string) {
	if r != nil && r.OriginatorIdentification != nil {
		OriginatorIdentification = *r.OriginatorIdentification
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) GetOriginatorName() (OriginatorName string) {
	if r != nil && r.OriginatorName != nil {
		OriginatorName = *r.OriginatorName
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) GetOriginatorPostalCode() (OriginatorPostalCode string) {
	if r != nil && r.OriginatorPostalCode != nil {
		OriginatorPostalCode = *r.OriginatorPostalCode
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) GetOriginatorStreetAddress() (OriginatorStreetAddress string) {
	if r != nil && r.OriginatorStreetAddress != nil {
		OriginatorStreetAddress = *r.OriginatorStreetAddress
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) GetOriginatorStateOrProvince() (OriginatorStateOrProvince string) {
	if r != nil && r.OriginatorStateOrProvince != nil {
		OriginatorStateOrProvince = *r.OriginatorStateOrProvince
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) GetPaymentRelatedInformation() (PaymentRelatedInformation string) {
	if r != nil && r.PaymentRelatedInformation != nil {
		PaymentRelatedInformation = *r.PaymentRelatedInformation
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) GetPaymentRelatedInformation2() (PaymentRelatedInformation2 string) {
	if r != nil && r.PaymentRelatedInformation2 != nil {
		PaymentRelatedInformation2 = *r.PaymentRelatedInformation2
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) GetReceiverIdentificationNumber() (ReceiverIdentificationNumber string) {
	if r != nil && r.ReceiverIdentificationNumber != nil {
		ReceiverIdentificationNumber = *r.ReceiverIdentificationNumber
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) GetReceiverStreetAddress() (ReceiverStreetAddress string) {
	if r != nil && r.ReceiverStreetAddress != nil {
		ReceiverStreetAddress = *r.ReceiverStreetAddress
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) GetReceiverCity() (ReceiverCity string) {
	if r != nil && r.ReceiverCity != nil {
		ReceiverCity = *r.ReceiverCity
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) GetReceiverStateOrProvince() (ReceiverStateOrProvince string) {
	if r != nil && r.ReceiverStateOrProvince != nil {
		ReceiverStateOrProvince = *r.ReceiverStateOrProvince
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) GetReceiverCountry() (ReceiverCountry string) {
	if r != nil && r.ReceiverCountry != nil {
		ReceiverCountry = *r.ReceiverCountry
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) GetReceiverPostalCode() (ReceiverPostalCode string) {
	if r != nil && r.ReceiverPostalCode != nil {
		ReceiverPostalCode = *r.ReceiverPostalCode
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) GetReceivingCompanyOrIndividualName() (ReceivingCompanyOrIndividualName string) {
	if r != nil && r.ReceivingCompanyOrIndividualName != nil {
		ReceivingCompanyOrIndividualName = *r.ReceivingCompanyOrIndividualName
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) GetReceivingDepositoryFinancialInstitutionName() (ReceivingDepositoryFinancialInstitutionName string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionName != nil {
		ReceivingDepositoryFinancialInstitutionName = *r.ReceivingDepositoryFinancialInstitutionName
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) GetReceivingDepositoryFinancialInstitutionIDQualifier() (ReceivingDepositoryFinancialInstitutionIDQualifier string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionIDQualifier != nil {
		ReceivingDepositoryFinancialInstitutionIDQualifier = *r.ReceivingDepositoryFinancialInstitutionIDQualifier
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) GetReceivingDepositoryFinancialInstitutionID() (ReceivingDepositoryFinancialInstitutionID string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionID != nil {
		ReceivingDepositoryFinancialInstitutionID = *r.ReceivingDepositoryFinancialInstitutionID
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) GetReceivingDepositoryFinancialInstitutionCountry() (ReceivingDepositoryFinancialInstitutionCountry string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionCountry != nil {
		ReceivingDepositoryFinancialInstitutionCountry = *r.ReceivingDepositoryFinancialInstitutionCountry
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) GetTraceNumber() (TraceNumber string) {
	if r != nil && r.TraceNumber != nil {
		TraceNumber = *r.TraceNumber
	}
	return
}

func (r CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) String() (result string) {
	return fmt.Sprintf("&CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline{Amount:%s ForeignExchangeIndicator:%s ForeignExchangeReferenceIndicator:%s ForeignExchangeReference:%s DestinationCountryCode:%s DestinationCurrencyCode:%s ForeignPaymentAmount:%s ForeignTraceNumber:%s InternationalTransactionTypeCode:%s OriginatingCurrencyCode:%s OriginatingDepositoryFinancialInstitutionName:%s OriginatingDepositoryFinancialInstitutionIDQualifier:%s OriginatingDepositoryFinancialInstitutionID:%s OriginatingDepositoryFinancialInstitutionBranchCountry:%s OriginatorCity:%s OriginatorCompanyEntryDescription:%s OriginatorCountry:%s OriginatorIdentification:%s OriginatorName:%s OriginatorPostalCode:%s OriginatorStreetAddress:%s OriginatorStateOrProvince:%s PaymentRelatedInformation:%s PaymentRelatedInformation2:%s ReceiverIdentificationNumber:%s ReceiverStreetAddress:%s ReceiverCity:%s ReceiverStateOrProvince:%s ReceiverCountry:%s ReceiverPostalCode:%s ReceivingCompanyOrIndividualName:%s ReceivingDepositoryFinancialInstitutionName:%s ReceivingDepositoryFinancialInstitutionIDQualifier:%s ReceivingDepositoryFinancialInstitutionID:%s ReceivingDepositoryFinancialInstitutionCountry:%s TraceNumber:%s}", core.FmtP(r.Amount), core.FmtP(r.ForeignExchangeIndicator), core.FmtP(r.ForeignExchangeReferenceIndicator), core.FmtP(r.ForeignExchangeReference), core.FmtP(r.DestinationCountryCode), core.FmtP(r.DestinationCurrencyCode), core.FmtP(r.ForeignPaymentAmount), core.FmtP(r.ForeignTraceNumber), core.FmtP(r.InternationalTransactionTypeCode), core.FmtP(r.OriginatingCurrencyCode), core.FmtP(r.OriginatingDepositoryFinancialInstitutionName), core.FmtP(r.OriginatingDepositoryFinancialInstitutionIDQualifier), core.FmtP(r.OriginatingDepositoryFinancialInstitutionID), core.FmtP(r.OriginatingDepositoryFinancialInstitutionBranchCountry), core.FmtP(r.OriginatorCity), core.FmtP(r.OriginatorCompanyEntryDescription), core.FmtP(r.OriginatorCountry), core.FmtP(r.OriginatorIdentification), core.FmtP(r.OriginatorName), core.FmtP(r.OriginatorPostalCode), core.FmtP(r.OriginatorStreetAddress), core.FmtP(r.OriginatorStateOrProvince), core.FmtP(r.PaymentRelatedInformation), core.FmtP(r.PaymentRelatedInformation2), core.FmtP(r.ReceiverIdentificationNumber), core.FmtP(r.ReceiverStreetAddress), core.FmtP(r.ReceiverCity), core.FmtP(r.ReceiverStateOrProvince), core.FmtP(r.ReceiverCountry), core.FmtP(r.ReceiverPostalCode), core.FmtP(r.ReceivingCompanyOrIndividualName), core.FmtP(r.ReceivingDepositoryFinancialInstitutionName), core.FmtP(r.ReceivingDepositoryFinancialInstitutionIDQualifier), core.FmtP(r.ReceivingDepositoryFinancialInstitutionID), core.FmtP(r.ReceivingDepositoryFinancialInstitutionCountry), core.FmtP(r.TraceNumber))
}

type CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency             *CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDeclineCurrency `pjson:"currency"`
	MerchantAcceptorID   *string                                                                       `pjson:"merchant_acceptor_id"`
	MerchantCity         *string                                                                       `pjson:"merchant_city"`
	MerchantCountry      *string                                                                       `pjson:"merchant_country"`
	MerchantDescriptor   *string                                                                       `pjson:"merchant_descriptor"`
	MerchantState        *string                                                                       `pjson:"merchant_state"`
	MerchantCategoryCode *string                                                                       `pjson:"merchant_category_code"`
	jsonFields           map[string]interface{}                                                        `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDecline using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDecline) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes
// CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDecline into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDecline) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDecline) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
// account currency.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDecline) GetCurrency() (Currency CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDeclineCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDecline) GetMerchantAcceptorID() (MerchantAcceptorID string) {
	if r != nil && r.MerchantAcceptorID != nil {
		MerchantAcceptorID = *r.MerchantAcceptorID
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDecline) GetMerchantCity() (MerchantCity string) {
	if r != nil && r.MerchantCity != nil {
		MerchantCity = *r.MerchantCity
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDecline) GetMerchantCountry() (MerchantCountry string) {
	if r != nil && r.MerchantCountry != nil {
		MerchantCountry = *r.MerchantCountry
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDecline) GetMerchantDescriptor() (MerchantDescriptor string) {
	if r != nil && r.MerchantDescriptor != nil {
		MerchantDescriptor = *r.MerchantDescriptor
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDecline) GetMerchantState() (MerchantState string) {
	if r != nil && r.MerchantState != nil {
		MerchantState = *r.MerchantState
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDecline) GetMerchantCategoryCode() (MerchantCategoryCode string) {
	if r != nil && r.MerchantCategoryCode != nil {
		MerchantCategoryCode = *r.MerchantCategoryCode
	}
	return
}

func (r CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDecline) String() (result string) {
	return fmt.Sprintf("&CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDecline{Amount:%s Currency:%s MerchantAcceptorID:%s MerchantCity:%s MerchantCountry:%s MerchantDescriptor:%s MerchantState:%s MerchantCategoryCode:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.MerchantAcceptorID), core.FmtP(r.MerchantCity), core.FmtP(r.MerchantCountry), core.FmtP(r.MerchantDescriptor), core.FmtP(r.MerchantState), core.FmtP(r.MerchantCategoryCode))
}

type CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDeclineCurrency string

const (
	CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDeclineCurrencyCad CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDeclineCurrency = "CAD"
	CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDeclineCurrencyChf CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDeclineCurrency = "CHF"
	CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDeclineCurrencyEur CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDeclineCurrency = "EUR"
	CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDeclineCurrencyGbp CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDeclineCurrency = "GBP"
	CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDeclineCurrencyJpy CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDeclineCurrency = "JPY"
	CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDeclineCurrencyUsd CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDeclineCurrency = "USD"
)

type CardAuthorizationSimulationDeclinedTransactionType string

const (
	CardAuthorizationSimulationDeclinedTransactionTypeDeclinedTransaction CardAuthorizationSimulationDeclinedTransactionType = "declined_transaction"
)

type CardAuthorizationSimulationType string

const (
	CardAuthorizationSimulationTypeInboundCardAuthorizationSimulationResult CardAuthorizationSimulationType = "inbound_card_authorization_simulation_result"
)

type SimulateAnAuthorizationOnACardParameters struct {
	// The authorization amount in cents.
	Amount *int64 `pjson:"amount"`
	// The identifier of the Card to be authorized.
	CardID *string `pjson:"card_id"`
	// The identifier of the Digital Wallet Token to be authorized.
	DigitalWalletTokenID *string                `pjson:"digital_wallet_token_id"`
	jsonFields           map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// SimulateAnAuthorizationOnACardParameters using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *SimulateAnAuthorizationOnACardParameters) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes SimulateAnAuthorizationOnACardParameters into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *SimulateAnAuthorizationOnACardParameters) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The authorization amount in cents.
func (r *SimulateAnAuthorizationOnACardParameters) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The identifier of the Card to be authorized.
func (r *SimulateAnAuthorizationOnACardParameters) GetCardID() (CardID string) {
	if r != nil && r.CardID != nil {
		CardID = *r.CardID
	}
	return
}

// The identifier of the Digital Wallet Token to be authorized.
func (r *SimulateAnAuthorizationOnACardParameters) GetDigitalWalletTokenID() (DigitalWalletTokenID string) {
	if r != nil && r.DigitalWalletTokenID != nil {
		DigitalWalletTokenID = *r.DigitalWalletTokenID
	}
	return
}

func (r SimulateAnAuthorizationOnACardParameters) String() (result string) {
	return fmt.Sprintf("&SimulateAnAuthorizationOnACardParameters{Amount:%s CardID:%s DigitalWalletTokenID:%s}", core.FmtP(r.Amount), core.FmtP(r.CardID), core.FmtP(r.DigitalWalletTokenID))
}

type SimulateSettlingACardAuthorizationParameters struct {
	// The identifier of the Card to create a settlement on.
	CardID *string `pjson:"card_id"`
	// The identifier of the Pending Transaction for the Card Authorization you wish to
	// settle.
	PendingTransactionID *string `pjson:"pending_transaction_id"`
	// The amount to be settled. This defaults to the amount of the Pending Transaction
	// being settled.
	Amount     *int64                 `pjson:"amount"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// SimulateSettlingACardAuthorizationParameters using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *SimulateSettlingACardAuthorizationParameters) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes SimulateSettlingACardAuthorizationParameters into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *SimulateSettlingACardAuthorizationParameters) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The identifier of the Card to create a settlement on.
func (r *SimulateSettlingACardAuthorizationParameters) GetCardID() (CardID string) {
	if r != nil && r.CardID != nil {
		CardID = *r.CardID
	}
	return
}

// The identifier of the Pending Transaction for the Card Authorization you wish to
// settle.
func (r *SimulateSettlingACardAuthorizationParameters) GetPendingTransactionID() (PendingTransactionID string) {
	if r != nil && r.PendingTransactionID != nil {
		PendingTransactionID = *r.PendingTransactionID
	}
	return
}

// The amount to be settled. This defaults to the amount of the Pending Transaction
// being settled.
func (r *SimulateSettlingACardAuthorizationParameters) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r SimulateSettlingACardAuthorizationParameters) String() (result string) {
	return fmt.Sprintf("&SimulateSettlingACardAuthorizationParameters{CardID:%s PendingTransactionID:%s Amount:%s}", core.FmtP(r.CardID), core.FmtP(r.PendingTransactionID), core.FmtP(r.Amount))
}

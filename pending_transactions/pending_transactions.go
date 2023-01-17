package pending_transactions

import "context"
import "increase/core"
import "fmt"
import "increase/pagination"

type PendingTransactionService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewPendingTransactionService(requester core.Requester) (r *PendingTransactionService) {
	r = &PendingTransactionService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

type PreloadedPendingTransactionService struct {
	PendingTransactions *PendingTransactionService
}

func (r *PreloadedPendingTransactionService) Init(service *PendingTransactionService) {
	r.PendingTransactions = service
}

func NewPreloadedPendingTransactionService(service *PendingTransactionService) (r *PreloadedPendingTransactionService) {
	r = &PreloadedPendingTransactionService{}
	r.Init(service)
	return
}

//
type PendingTransaction struct {
	// The identifier for the account this Pending Transaction belongs to.
	AccountID *string `json:"account_id"`
	// The Pending Transaction amount in the minor unit of its currency. For dollars,
	// for example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Pending
	// Transaction's currency. This will match the currency on the Pending
	// Transcation's Account.
	Currency *PendingTransactionCurrency `json:"currency"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the Pending
	// Transaction occured.
	CreatedAt *string `json:"created_at"`
	// For a Pending Transaction related to a transfer, this is the description you
	// provide. For a Pending Transaction related to a payment, this is the description
	// the vendor provides.
	Description *string `json:"description"`
	// The Pending Transaction identifier.
	ID *string `json:"id"`
	// The identifier for the route this Pending Transaction came through. Routes are
	// things like cards and ACH details.
	RouteID *string `json:"route_id"`
	// The type of the route this Pending Transaction came through.
	RouteType *string `json:"route_type"`
	// This is an object giving more details on the network-level event that caused the
	// Pending Transaction. For example, for a card transaction this lists the
	// merchant's industry and location.
	Source *PendingTransactionSource `json:"source"`
	// Whether the Pending Transaction has been confirmed and has an associated
	// Transaction.
	Status *PendingTransactionStatus `json:"status"`
	// A constant representing the object's type. For this resource it will always be
	// `pending_transaction`.
	Type *PendingTransactionType `json:"type"`
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
func (r *PendingTransaction) GetAmount() (Amount int) {
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

type PendingTransactionCurrency string

const (
	PendingTransactionCurrencyCad PendingTransactionCurrency = "CAD"
	PendingTransactionCurrencyChf PendingTransactionCurrency = "CHF"
	PendingTransactionCurrencyEur PendingTransactionCurrency = "EUR"
	PendingTransactionCurrencyGbp PendingTransactionCurrency = "GBP"
	PendingTransactionCurrencyJpy PendingTransactionCurrency = "JPY"
	PendingTransactionCurrencyUsd PendingTransactionCurrency = "USD"
)

//
type PendingTransactionSource struct {
	// The type of transaction that took place. We may add additional possible values
	// for this enum over time; your application should be able to handle such
	// additions gracefully.
	Category *PendingTransactionSourceCategory `json:"category"`
	// A Account Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `account_transfer_instruction`.
	AccountTransferInstruction *PendingTransactionSourceAccountTransferInstruction `json:"account_transfer_instruction"`
	// A ACH Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `ach_transfer_instruction`.
	ACHTransferInstruction *PendingTransactionSourceACHTransferInstruction `json:"ach_transfer_instruction"`
	// A Card Authorization object. This field will be present in the JSON response if
	// and only if `category` is equal to `card_authorization`.
	CardAuthorization *PendingTransactionSourceCardAuthorization `json:"card_authorization"`
	// A Check Deposit Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_deposit_instruction`.
	CheckDepositInstruction *PendingTransactionSourceCheckDepositInstruction `json:"check_deposit_instruction"`
	// A Check Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_instruction`.
	CheckTransferInstruction *PendingTransactionSourceCheckTransferInstruction `json:"check_transfer_instruction"`
	// A Deprecated Card Authorization object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_route_authorization`.
	CardRouteAuthorization *PendingTransactionSourceCardRouteAuthorization `json:"card_route_authorization"`
	// A Wire Drawdown Payment Instruction object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `wire_drawdown_payment_instruction`.
	WireDrawdownPaymentInstruction *PendingTransactionSourceWireDrawdownPaymentInstruction `json:"wire_drawdown_payment_instruction"`
	// A Wire Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_instruction`.
	WireTransferInstruction *PendingTransactionSourceWireTransferInstruction `json:"wire_transfer_instruction"`
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

type PendingTransactionSourceCategory string

const (
	PendingTransactionSourceCategoryAccountTransferInstruction          PendingTransactionSourceCategory = "account_transfer_instruction"
	PendingTransactionSourceCategoryACHTransferInstruction              PendingTransactionSourceCategory = "ach_transfer_instruction"
	PendingTransactionSourceCategoryCardAuthorization                   PendingTransactionSourceCategory = "card_authorization"
	PendingTransactionSourceCategoryCheckDepositInstruction             PendingTransactionSourceCategory = "check_deposit_instruction"
	PendingTransactionSourceCategoryCheckTransferInstruction            PendingTransactionSourceCategory = "check_transfer_instruction"
	PendingTransactionSourceCategoryCardRouteAuthorization              PendingTransactionSourceCategory = "card_route_authorization"
	PendingTransactionSourceCategoryRealTimePaymentsTransferInstruction PendingTransactionSourceCategory = "real_time_payments_transfer_instruction"
	PendingTransactionSourceCategoryWireDrawdownPaymentInstruction      PendingTransactionSourceCategory = "wire_drawdown_payment_instruction"
	PendingTransactionSourceCategoryWireTransferInstruction             PendingTransactionSourceCategory = "wire_transfer_instruction"
	PendingTransactionSourceCategoryOther                               PendingTransactionSourceCategory = "other"
)

//
type PendingTransactionSourceAccountTransferInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency *PendingTransactionSourceAccountTransferInstructionCurrency `json:"currency"`
	// The identifier of the Account Transfer that led to this Pending Transaction.
	TransferID *string `json:"transfer_id"`
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *PendingTransactionSourceAccountTransferInstruction) GetAmount() (Amount int) {
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

type PendingTransactionSourceAccountTransferInstructionCurrency string

const (
	PendingTransactionSourceAccountTransferInstructionCurrencyCad PendingTransactionSourceAccountTransferInstructionCurrency = "CAD"
	PendingTransactionSourceAccountTransferInstructionCurrencyChf PendingTransactionSourceAccountTransferInstructionCurrency = "CHF"
	PendingTransactionSourceAccountTransferInstructionCurrencyEur PendingTransactionSourceAccountTransferInstructionCurrency = "EUR"
	PendingTransactionSourceAccountTransferInstructionCurrencyGbp PendingTransactionSourceAccountTransferInstructionCurrency = "GBP"
	PendingTransactionSourceAccountTransferInstructionCurrencyJpy PendingTransactionSourceAccountTransferInstructionCurrency = "JPY"
	PendingTransactionSourceAccountTransferInstructionCurrencyUsd PendingTransactionSourceAccountTransferInstructionCurrency = "USD"
)

//
type PendingTransactionSourceACHTransferInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int `json:"amount"`
	// The identifier of the ACH Transfer that led to this Pending Transaction.
	TransferID *string `json:"transfer_id"`
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *PendingTransactionSourceACHTransferInstruction) GetAmount() (Amount int) {
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

//
type PendingTransactionSourceCardAuthorization struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *PendingTransactionSourceCardAuthorizationCurrency `json:"currency"`
	//
	MerchantAcceptorID *string `json:"merchant_acceptor_id"`
	//
	MerchantCity *string `json:"merchant_city"`
	//
	MerchantCountry *string `json:"merchant_country"`
	//
	MerchantDescriptor *string `json:"merchant_descriptor"`
	//
	MerchantCategoryCode *string `json:"merchant_category_code"`
	// The identifier of the Real-Time Decision sent to approve or decline this
	// transaction.
	RealTimeDecisionID *string `json:"real_time_decision_id"`
	// If the authorization was made via a Digital Wallet Token (such as an Apple Pay
	// purchase), the identifier of the token that was used.
	DigitalWalletTokenID *string `json:"digital_wallet_token_id"`
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *PendingTransactionSourceCardAuthorization) GetAmount() (Amount int) {
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

func (r *PendingTransactionSourceCardAuthorization) GetMerchantAcceptorID() (MerchantAcceptorID string) {
	if r != nil && r.MerchantAcceptorID != nil {
		MerchantAcceptorID = *r.MerchantAcceptorID
	}
	return
}

func (r *PendingTransactionSourceCardAuthorization) GetMerchantCity() (MerchantCity string) {
	if r != nil && r.MerchantCity != nil {
		MerchantCity = *r.MerchantCity
	}
	return
}

func (r *PendingTransactionSourceCardAuthorization) GetMerchantCountry() (MerchantCountry string) {
	if r != nil && r.MerchantCountry != nil {
		MerchantCountry = *r.MerchantCountry
	}
	return
}

func (r *PendingTransactionSourceCardAuthorization) GetMerchantDescriptor() (MerchantDescriptor string) {
	if r != nil && r.MerchantDescriptor != nil {
		MerchantDescriptor = *r.MerchantDescriptor
	}
	return
}

func (r *PendingTransactionSourceCardAuthorization) GetMerchantCategoryCode() (MerchantCategoryCode string) {
	if r != nil && r.MerchantCategoryCode != nil {
		MerchantCategoryCode = *r.MerchantCategoryCode
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

type PendingTransactionSourceCardAuthorizationCurrency string

const (
	PendingTransactionSourceCardAuthorizationCurrencyCad PendingTransactionSourceCardAuthorizationCurrency = "CAD"
	PendingTransactionSourceCardAuthorizationCurrencyChf PendingTransactionSourceCardAuthorizationCurrency = "CHF"
	PendingTransactionSourceCardAuthorizationCurrencyEur PendingTransactionSourceCardAuthorizationCurrency = "EUR"
	PendingTransactionSourceCardAuthorizationCurrencyGbp PendingTransactionSourceCardAuthorizationCurrency = "GBP"
	PendingTransactionSourceCardAuthorizationCurrencyJpy PendingTransactionSourceCardAuthorizationCurrency = "JPY"
	PendingTransactionSourceCardAuthorizationCurrencyUsd PendingTransactionSourceCardAuthorizationCurrency = "USD"
)

//
type PendingTransactionSourceCheckDepositInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *PendingTransactionSourceCheckDepositInstructionCurrency `json:"currency"`
	// The identifier of the File containing the image of the front of the check that
	// was deposited.
	FrontImageFileID *string `json:"front_image_file_id"`
	// The identifier of the File containing the image of the back of the check that
	// was deposited.
	BackImageFileID *string `json:"back_image_file_id"`
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *PendingTransactionSourceCheckDepositInstruction) GetAmount() (Amount int) {
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

type PendingTransactionSourceCheckDepositInstructionCurrency string

const (
	PendingTransactionSourceCheckDepositInstructionCurrencyCad PendingTransactionSourceCheckDepositInstructionCurrency = "CAD"
	PendingTransactionSourceCheckDepositInstructionCurrencyChf PendingTransactionSourceCheckDepositInstructionCurrency = "CHF"
	PendingTransactionSourceCheckDepositInstructionCurrencyEur PendingTransactionSourceCheckDepositInstructionCurrency = "EUR"
	PendingTransactionSourceCheckDepositInstructionCurrencyGbp PendingTransactionSourceCheckDepositInstructionCurrency = "GBP"
	PendingTransactionSourceCheckDepositInstructionCurrencyJpy PendingTransactionSourceCheckDepositInstructionCurrency = "JPY"
	PendingTransactionSourceCheckDepositInstructionCurrencyUsd PendingTransactionSourceCheckDepositInstructionCurrency = "USD"
)

//
type PendingTransactionSourceCheckTransferInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
	// currency.
	Currency *PendingTransactionSourceCheckTransferInstructionCurrency `json:"currency"`
	// The identifier of the Check Transfer that led to this Pending Transaction.
	TransferID *string `json:"transfer_id"`
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *PendingTransactionSourceCheckTransferInstruction) GetAmount() (Amount int) {
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

type PendingTransactionSourceCheckTransferInstructionCurrency string

const (
	PendingTransactionSourceCheckTransferInstructionCurrencyCad PendingTransactionSourceCheckTransferInstructionCurrency = "CAD"
	PendingTransactionSourceCheckTransferInstructionCurrencyChf PendingTransactionSourceCheckTransferInstructionCurrency = "CHF"
	PendingTransactionSourceCheckTransferInstructionCurrencyEur PendingTransactionSourceCheckTransferInstructionCurrency = "EUR"
	PendingTransactionSourceCheckTransferInstructionCurrencyGbp PendingTransactionSourceCheckTransferInstructionCurrency = "GBP"
	PendingTransactionSourceCheckTransferInstructionCurrencyJpy PendingTransactionSourceCheckTransferInstructionCurrency = "JPY"
	PendingTransactionSourceCheckTransferInstructionCurrencyUsd PendingTransactionSourceCheckTransferInstructionCurrency = "USD"
)

//
type PendingTransactionSourceCardRouteAuthorization struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *PendingTransactionSourceCardRouteAuthorizationCurrency `json:"currency"`
	//
	MerchantAcceptorID *string `json:"merchant_acceptor_id"`
	//
	MerchantCity *string `json:"merchant_city"`
	//
	MerchantCountry *string `json:"merchant_country"`
	//
	MerchantDescriptor *string `json:"merchant_descriptor"`
	//
	MerchantCategoryCode *string `json:"merchant_category_code"`
	//
	MerchantState *string `json:"merchant_state"`
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *PendingTransactionSourceCardRouteAuthorization) GetAmount() (Amount int) {
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

type PendingTransactionSourceCardRouteAuthorizationCurrency string

const (
	PendingTransactionSourceCardRouteAuthorizationCurrencyCad PendingTransactionSourceCardRouteAuthorizationCurrency = "CAD"
	PendingTransactionSourceCardRouteAuthorizationCurrencyChf PendingTransactionSourceCardRouteAuthorizationCurrency = "CHF"
	PendingTransactionSourceCardRouteAuthorizationCurrencyEur PendingTransactionSourceCardRouteAuthorizationCurrency = "EUR"
	PendingTransactionSourceCardRouteAuthorizationCurrencyGbp PendingTransactionSourceCardRouteAuthorizationCurrency = "GBP"
	PendingTransactionSourceCardRouteAuthorizationCurrencyJpy PendingTransactionSourceCardRouteAuthorizationCurrency = "JPY"
	PendingTransactionSourceCardRouteAuthorizationCurrencyUsd PendingTransactionSourceCardRouteAuthorizationCurrency = "USD"
)

//
type PendingTransactionSourceWireDrawdownPaymentInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int `json:"amount"`
	//
	AccountNumber *string `json:"account_number"`
	//
	RoutingNumber *string `json:"routing_number"`
	//
	MessageToRecipient *string `json:"message_to_recipient"`
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *PendingTransactionSourceWireDrawdownPaymentInstruction) GetAmount() (Amount int) {
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

//
type PendingTransactionSourceWireTransferInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int `json:"amount"`
	//
	AccountNumber *string `json:"account_number"`
	//
	RoutingNumber *string `json:"routing_number"`
	//
	MessageToRecipient *string `json:"message_to_recipient"`
	//
	TransferID *string `json:"transfer_id"`
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *PendingTransactionSourceWireTransferInstruction) GetAmount() (Amount int) {
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

type PendingTransactionStatus string

const (
	PendingTransactionStatusPending  PendingTransactionStatus = "pending"
	PendingTransactionStatusComplete PendingTransactionStatus = "complete"
)

type PendingTransactionType string

const (
	PendingTransactionTypePendingTransaction PendingTransactionType = "pending_transaction"
)

type ListPendingTransactionsQuery struct {
	// Return the page of entries after this one.
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int `query:"limit"`
	// Filter pending transactions to those belonging to the specified Account.
	AccountID *string `query:"account_id"`
	// Filter pending transactions to those belonging to the specified Route.
	RouteID *string `query:"route_id"`
	// Filter pending transactions to those caused by the specified source.
	SourceID *string                             `query:"source_id"`
	Status   *ListPendingTransactionsQueryStatus `query:"status"`
}

// Return the page of entries after this one.
func (r *ListPendingTransactionsQuery) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *ListPendingTransactionsQuery) GetLimit() (Limit int) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// Filter pending transactions to those belonging to the specified Account.
func (r *ListPendingTransactionsQuery) GetAccountID() (AccountID string) {
	if r != nil && r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

// Filter pending transactions to those belonging to the specified Route.
func (r *ListPendingTransactionsQuery) GetRouteID() (RouteID string) {
	if r != nil && r.RouteID != nil {
		RouteID = *r.RouteID
	}
	return
}

// Filter pending transactions to those caused by the specified source.
func (r *ListPendingTransactionsQuery) GetSourceID() (SourceID string) {
	if r != nil && r.SourceID != nil {
		SourceID = *r.SourceID
	}
	return
}

func (r *ListPendingTransactionsQuery) GetStatus() (Status ListPendingTransactionsQueryStatus) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

type ListPendingTransactionsQueryStatus struct {
	// Return results whose value is in the provided list. For GET requests, this
	// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
	In *[]ListPendingTransactionsQueryStatusIn `json:"in,omitempty"`
}

// Return results whose value is in the provided list. For GET requests, this
// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
func (r *ListPendingTransactionsQueryStatus) GetIn() (In []ListPendingTransactionsQueryStatusIn) {
	if r != nil && r.In != nil {
		In = *r.In
	}
	return
}

type ListPendingTransactionsQueryStatusIn string

const (
	ListPendingTransactionsQueryStatusInPending  ListPendingTransactionsQueryStatusIn = "pending"
	ListPendingTransactionsQueryStatusInComplete ListPendingTransactionsQueryStatusIn = "complete"
)

//
type PendingTransactionList struct {
	// The contents of the list.
	Data *[]PendingTransaction `json:"data"`
	// A pointer to a place in the list.
	NextCursor *string `json:"next_cursor"`
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

func (r *PendingTransactionService) Retrieve(ctx context.Context, pending_transaction_id string, opts ...*core.RequestOpts) (res *PendingTransaction, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/pending_transactions/%s", pending_transaction_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *PreloadedPendingTransactionService) Retrieve(ctx context.Context, pending_transaction_id string, opts ...*core.RequestOpts) (res *PendingTransaction, err error) {
	err = r.PendingTransactions.get(
		ctx,
		fmt.Sprintf("/pending_transactions/%s", pending_transaction_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

type PendingTransactionsPage struct {
	*pagination.Page[PendingTransaction]
}

func (r *PendingTransactionsPage) PendingTransaction() *PendingTransaction {
	return r.Current()
}

func (r *PendingTransactionsPage) GetNextPage() (*PendingTransactionsPage, error) {
	if page, err := r.Page.GetNextPage(); err != nil {
		return nil, err
	} else {
		return &PendingTransactionsPage{page}, nil
	}
}

func (r *PendingTransactionService) List(ctx context.Context, query *ListPendingTransactionsQuery, opts ...*core.RequestOpts) (res *PendingTransactionsPage, err error) {
	page := &PendingTransactionsPage{
		Page: &pagination.Page[PendingTransaction]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/pending_transactions",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

func (r *PreloadedPendingTransactionService) List(ctx context.Context, query *ListPendingTransactionsQuery, opts ...*core.RequestOpts) (res *PendingTransactionsPage, err error) {
	page := &PendingTransactionsPage{
		Page: &pagination.Page[PendingTransaction]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/pending_transactions",
			},
			Requester: r.PendingTransactions.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

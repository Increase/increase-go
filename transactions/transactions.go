package transactions

import "context"
import "increase/core"
import "fmt"
import "increase/pagination"

type TransactionService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewTransactionService(requester core.Requester) (r *TransactionService) {
	r = &TransactionService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

type PreloadedTransactionService struct {
	Transactions *TransactionService
}

func (r *PreloadedTransactionService) Init(service *TransactionService) {
	r.Transactions = service
}

func NewPreloadedTransactionService(service *TransactionService) (r *PreloadedTransactionService) {
	r = &PreloadedTransactionService{}
	r.Init(service)
	return
}

//
type Transaction struct {
	// The identifier for the Account the Transaction belongs to.
	AccountID *string `json:"account_id"`
	// The Transaction amount in the minor unit of its currency. For dollars, for
	// example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// Transaction's currency. This will match the currency on the Transcation's
	// Account.
	Currency *TransactionCurrency `json:"currency"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the
	// Transaction occured.
	CreatedAt *string `json:"created_at"`
	// For a Transaction related to a transfer, this is the description you provide.
	// For a Transaction related to a payment, this is the description the vendor
	// provides.
	Description *string `json:"description"`
	// The Transaction identifier.
	ID *string `json:"id"`
	// The identifier for the route this Transaction came through. Routes are things
	// like cards and ACH details.
	RouteID *string `json:"route_id"`
	// The type of the route this Transaction came through.
	RouteType *string `json:"route_type"`
	// This is an object giving more details on the network-level event that caused the
	// Transaction. Note that for backwards compatibility reasons, additional
	// undocumented keys may appear in this object. These should be treated as
	// deprecated and will be removed in the future.
	Source *TransactionSource `json:"source"`
	// A constant representing the object's type. For this resource it will always be
	// `transaction`.
	Type *TransactionType `json:"type"`
}

// The identifier for the Account the Transaction belongs to.
func (r *Transaction) GetAccountID() (AccountID string) {
	if r != nil && r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

// The Transaction amount in the minor unit of its currency. For dollars, for
// example, this is cents.
func (r *Transaction) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// Transaction's currency. This will match the currency on the Transcation's
// Account.
func (r *Transaction) GetCurrency() (Currency TransactionCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the
// Transaction occured.
func (r *Transaction) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// For a Transaction related to a transfer, this is the description you provide.
// For a Transaction related to a payment, this is the description the vendor
// provides.
func (r *Transaction) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The Transaction identifier.
func (r *Transaction) GetID() (ID string) {
	if r != nil && r.ID != nil {
		ID = *r.ID
	}
	return
}

// The identifier for the route this Transaction came through. Routes are things
// like cards and ACH details.
func (r *Transaction) GetRouteID() (RouteID string) {
	if r != nil && r.RouteID != nil {
		RouteID = *r.RouteID
	}
	return
}

// The type of the route this Transaction came through.
func (r *Transaction) GetRouteType() (RouteType string) {
	if r != nil && r.RouteType != nil {
		RouteType = *r.RouteType
	}
	return
}

// This is an object giving more details on the network-level event that caused the
// Transaction. Note that for backwards compatibility reasons, additional
// undocumented keys may appear in this object. These should be treated as
// deprecated and will be removed in the future.
func (r *Transaction) GetSource() (Source TransactionSource) {
	if r != nil && r.Source != nil {
		Source = *r.Source
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `transaction`.
func (r *Transaction) GetType() (Type TransactionType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

type TransactionCurrency string

const (
	TransactionCurrencyCad TransactionCurrency = "CAD"
	TransactionCurrencyChf TransactionCurrency = "CHF"
	TransactionCurrencyEur TransactionCurrency = "EUR"
	TransactionCurrencyGbp TransactionCurrency = "GBP"
	TransactionCurrencyJpy TransactionCurrency = "JPY"
	TransactionCurrencyUsd TransactionCurrency = "USD"
)

//
type TransactionSource struct {
	// The type of transaction that took place. We may add additional possible values
	// for this enum over time; your application should be able to handle such
	// additions gracefully.
	Category *TransactionSourceCategory `json:"category"`
	// A Account Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `account_transfer_intention`.
	AccountTransferIntention *TransactionSourceAccountTransferIntention `json:"account_transfer_intention"`
	// A ACH Check Conversion Return object. This field will be present in the JSON
	// response if and only if `category` is equal to `ach_check_conversion_return`.
	ACHCheckConversionReturn *TransactionSourceACHCheckConversionReturn `json:"ach_check_conversion_return"`
	// A ACH Check Conversion object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_check_conversion`.
	ACHCheckConversion *TransactionSourceACHCheckConversion `json:"ach_check_conversion"`
	// A ACH Transfer Intention object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_transfer_intention`.
	ACHTransferIntention *TransactionSourceACHTransferIntention `json:"ach_transfer_intention"`
	// A ACH Transfer Rejection object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_transfer_rejection`.
	ACHTransferRejection *TransactionSourceACHTransferRejection `json:"ach_transfer_rejection"`
	// A ACH Transfer Return object. This field will be present in the JSON response if
	// and only if `category` is equal to `ach_transfer_return`.
	ACHTransferReturn *TransactionSourceACHTransferReturn `json:"ach_transfer_return"`
	// A Card Dispute Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_dispute_acceptance`.
	CardDisputeAcceptance *TransactionSourceCardDisputeAcceptance `json:"card_dispute_acceptance"`
	// A Card Refund object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_refund`.
	CardRefund *TransactionSourceCardRefund `json:"card_refund"`
	// A Card Settlement object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_settlement`.
	CardSettlement *TransactionSourceCardSettlement `json:"card_settlement"`
	// A Check Deposit Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_deposit_acceptance`.
	CheckDepositAcceptance *TransactionSourceCheckDepositAcceptance `json:"check_deposit_acceptance"`
	// A Check Deposit Return object. This field will be present in the JSON response
	// if and only if `category` is equal to `check_deposit_return`.
	CheckDepositReturn *TransactionSourceCheckDepositReturn `json:"check_deposit_return"`
	// A Check Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_intention`.
	CheckTransferIntention *TransactionSourceCheckTransferIntention `json:"check_transfer_intention"`
	// A Check Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_rejection`.
	CheckTransferRejection *TransactionSourceCheckTransferRejection `json:"check_transfer_rejection"`
	// A Check Transfer Stop Payment Request object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `check_transfer_stop_payment_request`.
	CheckTransferStopPaymentRequest *TransactionSourceCheckTransferStopPaymentRequest `json:"check_transfer_stop_payment_request"`
	// A Dispute Resolution object. This field will be present in the JSON response if
	// and only if `category` is equal to `dispute_resolution`.
	DisputeResolution *TransactionSourceDisputeResolution `json:"dispute_resolution"`
	// A Empyreal Cash Deposit object. This field will be present in the JSON response
	// if and only if `category` is equal to `empyreal_cash_deposit`.
	EmpyrealCashDeposit *TransactionSourceEmpyrealCashDeposit `json:"empyreal_cash_deposit"`
	// A Inbound ACH Transfer object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_ach_transfer`.
	InboundACHTransfer *TransactionSourceInboundACHTransfer `json:"inbound_ach_transfer"`
	// A Inbound Check object. This field will be present in the JSON response if and
	// only if `category` is equal to `inbound_check`.
	InboundCheck *TransactionSourceInboundCheck `json:"inbound_check"`
	// A Inbound International ACH Transfer object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `inbound_international_ach_transfer`.
	InboundInternationalACHTransfer *TransactionSourceInboundInternationalACHTransfer `json:"inbound_international_ach_transfer"`
	// A Inbound Real Time Payments Transfer Confirmation object. This field will be
	// present in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_confirmation`.
	InboundRealTimePaymentsTransferConfirmation *TransactionSourceInboundRealTimePaymentsTransferConfirmation `json:"inbound_real_time_payments_transfer_confirmation"`
	// A Inbound Wire Drawdown Payment Reversal object. This field will be present in
	// the JSON response if and only if `category` is equal to
	// `inbound_wire_drawdown_payment_reversal`.
	InboundWireDrawdownPaymentReversal *TransactionSourceInboundWireDrawdownPaymentReversal `json:"inbound_wire_drawdown_payment_reversal"`
	// A Inbound Wire Drawdown Payment object. This field will be present in the JSON
	// response if and only if `category` is equal to `inbound_wire_drawdown_payment`.
	InboundWireDrawdownPayment *TransactionSourceInboundWireDrawdownPayment `json:"inbound_wire_drawdown_payment"`
	// A Inbound Wire Reversal object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_wire_reversal`.
	InboundWireReversal *TransactionSourceInboundWireReversal `json:"inbound_wire_reversal"`
	// A Inbound Wire Transfer object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_wire_transfer`.
	InboundWireTransfer *TransactionSourceInboundWireTransfer `json:"inbound_wire_transfer"`
	// A Internal Source object. This field will be present in the JSON response if and
	// only if `category` is equal to `internal_source`.
	InternalSource *TransactionSourceInternalSource `json:"internal_source"`
	// A Deprecated Card Refund object. This field will be present in the JSON response
	// if and only if `category` is equal to `card_route_refund`.
	CardRouteRefund *TransactionSourceCardRouteRefund `json:"card_route_refund"`
	// A Deprecated Card Settlement object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_route_settlement`.
	CardRouteSettlement *TransactionSourceCardRouteSettlement `json:"card_route_settlement"`
	// A Sample Funds object. This field will be present in the JSON response if and
	// only if `category` is equal to `sample_funds`.
	SampleFunds *TransactionSourceSampleFunds `json:"sample_funds"`
	// A Wire Drawdown Payment Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to
	// `wire_drawdown_payment_intention`.
	WireDrawdownPaymentIntention *TransactionSourceWireDrawdownPaymentIntention `json:"wire_drawdown_payment_intention"`
	// A Wire Drawdown Payment Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to
	// `wire_drawdown_payment_rejection`.
	WireDrawdownPaymentRejection *TransactionSourceWireDrawdownPaymentRejection `json:"wire_drawdown_payment_rejection"`
	// A Wire Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_intention`.
	WireTransferIntention *TransactionSourceWireTransferIntention `json:"wire_transfer_intention"`
	// A Wire Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_rejection`.
	WireTransferRejection *TransactionSourceWireTransferRejection `json:"wire_transfer_rejection"`
}

// The type of transaction that took place. We may add additional possible values
// for this enum over time; your application should be able to handle such
// additions gracefully.
func (r *TransactionSource) GetCategory() (Category TransactionSourceCategory) {
	if r != nil && r.Category != nil {
		Category = *r.Category
	}
	return
}

// A Account Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `account_transfer_intention`.
func (r *TransactionSource) GetAccountTransferIntention() (AccountTransferIntention TransactionSourceAccountTransferIntention) {
	if r != nil && r.AccountTransferIntention != nil {
		AccountTransferIntention = *r.AccountTransferIntention
	}
	return
}

// A ACH Check Conversion Return object. This field will be present in the JSON
// response if and only if `category` is equal to `ach_check_conversion_return`.
func (r *TransactionSource) GetACHCheckConversionReturn() (ACHCheckConversionReturn TransactionSourceACHCheckConversionReturn) {
	if r != nil && r.ACHCheckConversionReturn != nil {
		ACHCheckConversionReturn = *r.ACHCheckConversionReturn
	}
	return
}

// A ACH Check Conversion object. This field will be present in the JSON response
// if and only if `category` is equal to `ach_check_conversion`.
func (r *TransactionSource) GetACHCheckConversion() (ACHCheckConversion TransactionSourceACHCheckConversion) {
	if r != nil && r.ACHCheckConversion != nil {
		ACHCheckConversion = *r.ACHCheckConversion
	}
	return
}

// A ACH Transfer Intention object. This field will be present in the JSON response
// if and only if `category` is equal to `ach_transfer_intention`.
func (r *TransactionSource) GetACHTransferIntention() (ACHTransferIntention TransactionSourceACHTransferIntention) {
	if r != nil && r.ACHTransferIntention != nil {
		ACHTransferIntention = *r.ACHTransferIntention
	}
	return
}

// A ACH Transfer Rejection object. This field will be present in the JSON response
// if and only if `category` is equal to `ach_transfer_rejection`.
func (r *TransactionSource) GetACHTransferRejection() (ACHTransferRejection TransactionSourceACHTransferRejection) {
	if r != nil && r.ACHTransferRejection != nil {
		ACHTransferRejection = *r.ACHTransferRejection
	}
	return
}

// A ACH Transfer Return object. This field will be present in the JSON response if
// and only if `category` is equal to `ach_transfer_return`.
func (r *TransactionSource) GetACHTransferReturn() (ACHTransferReturn TransactionSourceACHTransferReturn) {
	if r != nil && r.ACHTransferReturn != nil {
		ACHTransferReturn = *r.ACHTransferReturn
	}
	return
}

// A Card Dispute Acceptance object. This field will be present in the JSON
// response if and only if `category` is equal to `card_dispute_acceptance`.
func (r *TransactionSource) GetCardDisputeAcceptance() (CardDisputeAcceptance TransactionSourceCardDisputeAcceptance) {
	if r != nil && r.CardDisputeAcceptance != nil {
		CardDisputeAcceptance = *r.CardDisputeAcceptance
	}
	return
}

// A Card Refund object. This field will be present in the JSON response if and
// only if `category` is equal to `card_refund`.
func (r *TransactionSource) GetCardRefund() (CardRefund TransactionSourceCardRefund) {
	if r != nil && r.CardRefund != nil {
		CardRefund = *r.CardRefund
	}
	return
}

// A Card Settlement object. This field will be present in the JSON response if and
// only if `category` is equal to `card_settlement`.
func (r *TransactionSource) GetCardSettlement() (CardSettlement TransactionSourceCardSettlement) {
	if r != nil && r.CardSettlement != nil {
		CardSettlement = *r.CardSettlement
	}
	return
}

// A Check Deposit Acceptance object. This field will be present in the JSON
// response if and only if `category` is equal to `check_deposit_acceptance`.
func (r *TransactionSource) GetCheckDepositAcceptance() (CheckDepositAcceptance TransactionSourceCheckDepositAcceptance) {
	if r != nil && r.CheckDepositAcceptance != nil {
		CheckDepositAcceptance = *r.CheckDepositAcceptance
	}
	return
}

// A Check Deposit Return object. This field will be present in the JSON response
// if and only if `category` is equal to `check_deposit_return`.
func (r *TransactionSource) GetCheckDepositReturn() (CheckDepositReturn TransactionSourceCheckDepositReturn) {
	if r != nil && r.CheckDepositReturn != nil {
		CheckDepositReturn = *r.CheckDepositReturn
	}
	return
}

// A Check Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `check_transfer_intention`.
func (r *TransactionSource) GetCheckTransferIntention() (CheckTransferIntention TransactionSourceCheckTransferIntention) {
	if r != nil && r.CheckTransferIntention != nil {
		CheckTransferIntention = *r.CheckTransferIntention
	}
	return
}

// A Check Transfer Rejection object. This field will be present in the JSON
// response if and only if `category` is equal to `check_transfer_rejection`.
func (r *TransactionSource) GetCheckTransferRejection() (CheckTransferRejection TransactionSourceCheckTransferRejection) {
	if r != nil && r.CheckTransferRejection != nil {
		CheckTransferRejection = *r.CheckTransferRejection
	}
	return
}

// A Check Transfer Stop Payment Request object. This field will be present in the
// JSON response if and only if `category` is equal to
// `check_transfer_stop_payment_request`.
func (r *TransactionSource) GetCheckTransferStopPaymentRequest() (CheckTransferStopPaymentRequest TransactionSourceCheckTransferStopPaymentRequest) {
	if r != nil && r.CheckTransferStopPaymentRequest != nil {
		CheckTransferStopPaymentRequest = *r.CheckTransferStopPaymentRequest
	}
	return
}

// A Dispute Resolution object. This field will be present in the JSON response if
// and only if `category` is equal to `dispute_resolution`.
func (r *TransactionSource) GetDisputeResolution() (DisputeResolution TransactionSourceDisputeResolution) {
	if r != nil && r.DisputeResolution != nil {
		DisputeResolution = *r.DisputeResolution
	}
	return
}

// A Empyreal Cash Deposit object. This field will be present in the JSON response
// if and only if `category` is equal to `empyreal_cash_deposit`.
func (r *TransactionSource) GetEmpyrealCashDeposit() (EmpyrealCashDeposit TransactionSourceEmpyrealCashDeposit) {
	if r != nil && r.EmpyrealCashDeposit != nil {
		EmpyrealCashDeposit = *r.EmpyrealCashDeposit
	}
	return
}

// A Inbound ACH Transfer object. This field will be present in the JSON response
// if and only if `category` is equal to `inbound_ach_transfer`.
func (r *TransactionSource) GetInboundACHTransfer() (InboundACHTransfer TransactionSourceInboundACHTransfer) {
	if r != nil && r.InboundACHTransfer != nil {
		InboundACHTransfer = *r.InboundACHTransfer
	}
	return
}

// A Inbound Check object. This field will be present in the JSON response if and
// only if `category` is equal to `inbound_check`.
func (r *TransactionSource) GetInboundCheck() (InboundCheck TransactionSourceInboundCheck) {
	if r != nil && r.InboundCheck != nil {
		InboundCheck = *r.InboundCheck
	}
	return
}

// A Inbound International ACH Transfer object. This field will be present in the
// JSON response if and only if `category` is equal to
// `inbound_international_ach_transfer`.
func (r *TransactionSource) GetInboundInternationalACHTransfer() (InboundInternationalACHTransfer TransactionSourceInboundInternationalACHTransfer) {
	if r != nil && r.InboundInternationalACHTransfer != nil {
		InboundInternationalACHTransfer = *r.InboundInternationalACHTransfer
	}
	return
}

// A Inbound Real Time Payments Transfer Confirmation object. This field will be
// present in the JSON response if and only if `category` is equal to
// `inbound_real_time_payments_transfer_confirmation`.
func (r *TransactionSource) GetInboundRealTimePaymentsTransferConfirmation() (InboundRealTimePaymentsTransferConfirmation TransactionSourceInboundRealTimePaymentsTransferConfirmation) {
	if r != nil && r.InboundRealTimePaymentsTransferConfirmation != nil {
		InboundRealTimePaymentsTransferConfirmation = *r.InboundRealTimePaymentsTransferConfirmation
	}
	return
}

// A Inbound Wire Drawdown Payment Reversal object. This field will be present in
// the JSON response if and only if `category` is equal to
// `inbound_wire_drawdown_payment_reversal`.
func (r *TransactionSource) GetInboundWireDrawdownPaymentReversal() (InboundWireDrawdownPaymentReversal TransactionSourceInboundWireDrawdownPaymentReversal) {
	if r != nil && r.InboundWireDrawdownPaymentReversal != nil {
		InboundWireDrawdownPaymentReversal = *r.InboundWireDrawdownPaymentReversal
	}
	return
}

// A Inbound Wire Drawdown Payment object. This field will be present in the JSON
// response if and only if `category` is equal to `inbound_wire_drawdown_payment`.
func (r *TransactionSource) GetInboundWireDrawdownPayment() (InboundWireDrawdownPayment TransactionSourceInboundWireDrawdownPayment) {
	if r != nil && r.InboundWireDrawdownPayment != nil {
		InboundWireDrawdownPayment = *r.InboundWireDrawdownPayment
	}
	return
}

// A Inbound Wire Reversal object. This field will be present in the JSON response
// if and only if `category` is equal to `inbound_wire_reversal`.
func (r *TransactionSource) GetInboundWireReversal() (InboundWireReversal TransactionSourceInboundWireReversal) {
	if r != nil && r.InboundWireReversal != nil {
		InboundWireReversal = *r.InboundWireReversal
	}
	return
}

// A Inbound Wire Transfer object. This field will be present in the JSON response
// if and only if `category` is equal to `inbound_wire_transfer`.
func (r *TransactionSource) GetInboundWireTransfer() (InboundWireTransfer TransactionSourceInboundWireTransfer) {
	if r != nil && r.InboundWireTransfer != nil {
		InboundWireTransfer = *r.InboundWireTransfer
	}
	return
}

// A Internal Source object. This field will be present in the JSON response if and
// only if `category` is equal to `internal_source`.
func (r *TransactionSource) GetInternalSource() (InternalSource TransactionSourceInternalSource) {
	if r != nil && r.InternalSource != nil {
		InternalSource = *r.InternalSource
	}
	return
}

// A Deprecated Card Refund object. This field will be present in the JSON response
// if and only if `category` is equal to `card_route_refund`.
func (r *TransactionSource) GetCardRouteRefund() (CardRouteRefund TransactionSourceCardRouteRefund) {
	if r != nil && r.CardRouteRefund != nil {
		CardRouteRefund = *r.CardRouteRefund
	}
	return
}

// A Deprecated Card Settlement object. This field will be present in the JSON
// response if and only if `category` is equal to `card_route_settlement`.
func (r *TransactionSource) GetCardRouteSettlement() (CardRouteSettlement TransactionSourceCardRouteSettlement) {
	if r != nil && r.CardRouteSettlement != nil {
		CardRouteSettlement = *r.CardRouteSettlement
	}
	return
}

// A Sample Funds object. This field will be present in the JSON response if and
// only if `category` is equal to `sample_funds`.
func (r *TransactionSource) GetSampleFunds() (SampleFunds TransactionSourceSampleFunds) {
	if r != nil && r.SampleFunds != nil {
		SampleFunds = *r.SampleFunds
	}
	return
}

// A Wire Drawdown Payment Intention object. This field will be present in the JSON
// response if and only if `category` is equal to
// `wire_drawdown_payment_intention`.
func (r *TransactionSource) GetWireDrawdownPaymentIntention() (WireDrawdownPaymentIntention TransactionSourceWireDrawdownPaymentIntention) {
	if r != nil && r.WireDrawdownPaymentIntention != nil {
		WireDrawdownPaymentIntention = *r.WireDrawdownPaymentIntention
	}
	return
}

// A Wire Drawdown Payment Rejection object. This field will be present in the JSON
// response if and only if `category` is equal to
// `wire_drawdown_payment_rejection`.
func (r *TransactionSource) GetWireDrawdownPaymentRejection() (WireDrawdownPaymentRejection TransactionSourceWireDrawdownPaymentRejection) {
	if r != nil && r.WireDrawdownPaymentRejection != nil {
		WireDrawdownPaymentRejection = *r.WireDrawdownPaymentRejection
	}
	return
}

// A Wire Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `wire_transfer_intention`.
func (r *TransactionSource) GetWireTransferIntention() (WireTransferIntention TransactionSourceWireTransferIntention) {
	if r != nil && r.WireTransferIntention != nil {
		WireTransferIntention = *r.WireTransferIntention
	}
	return
}

// A Wire Transfer Rejection object. This field will be present in the JSON
// response if and only if `category` is equal to `wire_transfer_rejection`.
func (r *TransactionSource) GetWireTransferRejection() (WireTransferRejection TransactionSourceWireTransferRejection) {
	if r != nil && r.WireTransferRejection != nil {
		WireTransferRejection = *r.WireTransferRejection
	}
	return
}

type TransactionSourceCategory string

const (
	TransactionSourceCategoryAccountTransferIntention                    TransactionSourceCategory = "account_transfer_intention"
	TransactionSourceCategoryACHCheckConversionReturn                    TransactionSourceCategory = "ach_check_conversion_return"
	TransactionSourceCategoryACHCheckConversion                          TransactionSourceCategory = "ach_check_conversion"
	TransactionSourceCategoryACHTransferIntention                        TransactionSourceCategory = "ach_transfer_intention"
	TransactionSourceCategoryACHTransferRejection                        TransactionSourceCategory = "ach_transfer_rejection"
	TransactionSourceCategoryACHTransferReturn                           TransactionSourceCategory = "ach_transfer_return"
	TransactionSourceCategoryCardDisputeAcceptance                       TransactionSourceCategory = "card_dispute_acceptance"
	TransactionSourceCategoryCardRefund                                  TransactionSourceCategory = "card_refund"
	TransactionSourceCategoryCardSettlement                              TransactionSourceCategory = "card_settlement"
	TransactionSourceCategoryCheckDepositAcceptance                      TransactionSourceCategory = "check_deposit_acceptance"
	TransactionSourceCategoryCheckDepositReturn                          TransactionSourceCategory = "check_deposit_return"
	TransactionSourceCategoryCheckTransferIntention                      TransactionSourceCategory = "check_transfer_intention"
	TransactionSourceCategoryCheckTransferRejection                      TransactionSourceCategory = "check_transfer_rejection"
	TransactionSourceCategoryCheckTransferStopPaymentRequest             TransactionSourceCategory = "check_transfer_stop_payment_request"
	TransactionSourceCategoryDisputeResolution                           TransactionSourceCategory = "dispute_resolution"
	TransactionSourceCategoryEmpyrealCashDeposit                         TransactionSourceCategory = "empyreal_cash_deposit"
	TransactionSourceCategoryInboundACHTransfer                          TransactionSourceCategory = "inbound_ach_transfer"
	TransactionSourceCategoryInboundCheck                                TransactionSourceCategory = "inbound_check"
	TransactionSourceCategoryInboundInternationalACHTransfer             TransactionSourceCategory = "inbound_international_ach_transfer"
	TransactionSourceCategoryInboundRealTimePaymentsTransferConfirmation TransactionSourceCategory = "inbound_real_time_payments_transfer_confirmation"
	TransactionSourceCategoryInboundWireDrawdownPaymentReversal          TransactionSourceCategory = "inbound_wire_drawdown_payment_reversal"
	TransactionSourceCategoryInboundWireDrawdownPayment                  TransactionSourceCategory = "inbound_wire_drawdown_payment"
	TransactionSourceCategoryInboundWireReversal                         TransactionSourceCategory = "inbound_wire_reversal"
	TransactionSourceCategoryInboundWireTransfer                         TransactionSourceCategory = "inbound_wire_transfer"
	TransactionSourceCategoryInternalSource                              TransactionSourceCategory = "internal_source"
	TransactionSourceCategoryCardRouteRefund                             TransactionSourceCategory = "card_route_refund"
	TransactionSourceCategoryCardRouteSettlement                         TransactionSourceCategory = "card_route_settlement"
	TransactionSourceCategoryRealTimePaymentsTransferAcknowledgement     TransactionSourceCategory = "real_time_payments_transfer_acknowledgement"
	TransactionSourceCategorySampleFunds                                 TransactionSourceCategory = "sample_funds"
	TransactionSourceCategoryWireDrawdownPaymentIntention                TransactionSourceCategory = "wire_drawdown_payment_intention"
	TransactionSourceCategoryWireDrawdownPaymentRejection                TransactionSourceCategory = "wire_drawdown_payment_rejection"
	TransactionSourceCategoryWireTransferIntention                       TransactionSourceCategory = "wire_transfer_intention"
	TransactionSourceCategoryWireTransferRejection                       TransactionSourceCategory = "wire_transfer_rejection"
	TransactionSourceCategoryOther                                       TransactionSourceCategory = "other"
)

//
type TransactionSourceAccountTransferIntention struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency *TransactionSourceAccountTransferIntentionCurrency `json:"currency"`
	// The description you chose to give the transfer.
	Description *string `json:"description"`
	// The identifier of the Account to where the Account Transfer was sent.
	DestinationAccountID *string `json:"destination_account_id"`
	// The identifier of the Account from where the Account Transfer was sent.
	SourceAccountID *string `json:"source_account_id"`
	// The identifier of the Account Transfer that led to this Pending Transaction.
	TransferID *string `json:"transfer_id"`
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *TransactionSourceAccountTransferIntention) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
// account currency.
func (r *TransactionSourceAccountTransferIntention) GetCurrency() (Currency TransactionSourceAccountTransferIntentionCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The description you chose to give the transfer.
func (r *TransactionSourceAccountTransferIntention) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The identifier of the Account to where the Account Transfer was sent.
func (r *TransactionSourceAccountTransferIntention) GetDestinationAccountID() (DestinationAccountID string) {
	if r != nil && r.DestinationAccountID != nil {
		DestinationAccountID = *r.DestinationAccountID
	}
	return
}

// The identifier of the Account from where the Account Transfer was sent.
func (r *TransactionSourceAccountTransferIntention) GetSourceAccountID() (SourceAccountID string) {
	if r != nil && r.SourceAccountID != nil {
		SourceAccountID = *r.SourceAccountID
	}
	return
}

// The identifier of the Account Transfer that led to this Pending Transaction.
func (r *TransactionSourceAccountTransferIntention) GetTransferID() (TransferID string) {
	if r != nil && r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

type TransactionSourceAccountTransferIntentionCurrency string

const (
	TransactionSourceAccountTransferIntentionCurrencyCad TransactionSourceAccountTransferIntentionCurrency = "CAD"
	TransactionSourceAccountTransferIntentionCurrencyChf TransactionSourceAccountTransferIntentionCurrency = "CHF"
	TransactionSourceAccountTransferIntentionCurrencyEur TransactionSourceAccountTransferIntentionCurrency = "EUR"
	TransactionSourceAccountTransferIntentionCurrencyGbp TransactionSourceAccountTransferIntentionCurrency = "GBP"
	TransactionSourceAccountTransferIntentionCurrencyJpy TransactionSourceAccountTransferIntentionCurrency = "JPY"
	TransactionSourceAccountTransferIntentionCurrencyUsd TransactionSourceAccountTransferIntentionCurrency = "USD"
)

//
type TransactionSourceACHCheckConversionReturn struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int `json:"amount"`
	// Why the transfer was returned.
	ReturnReasonCode *string `json:"return_reason_code"`
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *TransactionSourceACHCheckConversionReturn) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// Why the transfer was returned.
func (r *TransactionSourceACHCheckConversionReturn) GetReturnReasonCode() (ReturnReasonCode string) {
	if r != nil && r.ReturnReasonCode != nil {
		ReturnReasonCode = *r.ReturnReasonCode
	}
	return
}

//
type TransactionSourceACHCheckConversion struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int `json:"amount"`
	// The identifier of the File containing an image of the returned check.
	FileID *string `json:"file_id"`
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *TransactionSourceACHCheckConversion) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The identifier of the File containing an image of the returned check.
func (r *TransactionSourceACHCheckConversion) GetFileID() (FileID string) {
	if r != nil && r.FileID != nil {
		FileID = *r.FileID
	}
	return
}

//
type TransactionSourceACHTransferIntention struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int `json:"amount"`
	//
	AccountNumber *string `json:"account_number"`
	//
	RoutingNumber *string `json:"routing_number"`
	//
	StatementDescriptor *string `json:"statement_descriptor"`
	// The identifier of the ACH Transfer that led to this Transaction.
	TransferID *string `json:"transfer_id"`
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *TransactionSourceACHTransferIntention) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *TransactionSourceACHTransferIntention) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

func (r *TransactionSourceACHTransferIntention) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

func (r *TransactionSourceACHTransferIntention) GetStatementDescriptor() (StatementDescriptor string) {
	if r != nil && r.StatementDescriptor != nil {
		StatementDescriptor = *r.StatementDescriptor
	}
	return
}

// The identifier of the ACH Transfer that led to this Transaction.
func (r *TransactionSourceACHTransferIntention) GetTransferID() (TransferID string) {
	if r != nil && r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

//
type TransactionSourceACHTransferRejection struct {
	// The identifier of the ACH Transfer that led to this Transaction.
	TransferID *string `json:"transfer_id"`
}

// The identifier of the ACH Transfer that led to this Transaction.
func (r *TransactionSourceACHTransferRejection) GetTransferID() (TransferID string) {
	if r != nil && r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

//
type TransactionSourceACHTransferReturn struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt *string `json:"created_at"`
	// Why the ACH Transfer was returned.
	ReturnReasonCode *TransactionSourceACHTransferReturnReturnReasonCode `json:"return_reason_code"`
	// The identifier of the ACH Transfer associated with this return.
	TransferID *string `json:"transfer_id"`
	// The identifier of the Tranasaction associated with this return.
	TransactionID *string `json:"transaction_id"`
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the transfer was created.
func (r *TransactionSourceACHTransferReturn) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// Why the ACH Transfer was returned.
func (r *TransactionSourceACHTransferReturn) GetReturnReasonCode() (ReturnReasonCode TransactionSourceACHTransferReturnReturnReasonCode) {
	if r != nil && r.ReturnReasonCode != nil {
		ReturnReasonCode = *r.ReturnReasonCode
	}
	return
}

// The identifier of the ACH Transfer associated with this return.
func (r *TransactionSourceACHTransferReturn) GetTransferID() (TransferID string) {
	if r != nil && r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

// The identifier of the Tranasaction associated with this return.
func (r *TransactionSourceACHTransferReturn) GetTransactionID() (TransactionID string) {
	if r != nil && r.TransactionID != nil {
		TransactionID = *r.TransactionID
	}
	return
}

type TransactionSourceACHTransferReturnReturnReasonCode string

const (
	TransactionSourceACHTransferReturnReturnReasonCodeInsufficientFund                                          TransactionSourceACHTransferReturnReturnReasonCode = "insufficient_fund"
	TransactionSourceACHTransferReturnReturnReasonCodeNoAccount                                                 TransactionSourceACHTransferReturnReturnReasonCode = "no_account"
	TransactionSourceACHTransferReturnReturnReasonCodeAccountClosed                                             TransactionSourceACHTransferReturnReturnReasonCode = "account_closed"
	TransactionSourceACHTransferReturnReturnReasonCodeInvalidAccountNumberStructure                             TransactionSourceACHTransferReturnReturnReasonCode = "invalid_account_number_structure"
	TransactionSourceACHTransferReturnReturnReasonCodeAccountFrozenEntryReturnedPerOfacInstruction              TransactionSourceACHTransferReturnReturnReasonCode = "account_frozen_entry_returned_per_ofac_instruction"
	TransactionSourceACHTransferReturnReturnReasonCodeCreditEntryRefusedByReceiver                              TransactionSourceACHTransferReturnReturnReasonCode = "credit_entry_refused_by_receiver"
	TransactionSourceACHTransferReturnReturnReasonCodeUnauthorizedDebitToConsumerAccountUsingCorporateSecCode   TransactionSourceACHTransferReturnReturnReasonCode = "unauthorized_debit_to_consumer_account_using_corporate_sec_code"
	TransactionSourceACHTransferReturnReturnReasonCodeCorporateCustomerAdvisedNotAuthorized                     TransactionSourceACHTransferReturnReturnReasonCode = "corporate_customer_advised_not_authorized"
	TransactionSourceACHTransferReturnReturnReasonCodePaymentStopped                                            TransactionSourceACHTransferReturnReturnReasonCode = "payment_stopped"
	TransactionSourceACHTransferReturnReturnReasonCodeNonTransactionAccount                                     TransactionSourceACHTransferReturnReturnReasonCode = "non_transaction_account"
	TransactionSourceACHTransferReturnReturnReasonCodeUncollectedFunds                                          TransactionSourceACHTransferReturnReturnReasonCode = "uncollected_funds"
	TransactionSourceACHTransferReturnReturnReasonCodeRoutingNumberCheckDigitError                              TransactionSourceACHTransferReturnReturnReasonCode = "routing_number_check_digit_error"
	TransactionSourceACHTransferReturnReturnReasonCodeCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete TransactionSourceACHTransferReturnReturnReasonCode = "customer_advised_unauthorized_improper_ineligible_or_incomplete"
	TransactionSourceACHTransferReturnReturnReasonCodeAmountFieldError                                          TransactionSourceACHTransferReturnReturnReasonCode = "amount_field_error"
	TransactionSourceACHTransferReturnReturnReasonCodeAuthorizationRevokedByCustomer                            TransactionSourceACHTransferReturnReturnReasonCode = "authorization_revoked_by_customer"
	TransactionSourceACHTransferReturnReturnReasonCodeInvalidACHRoutingNumber                                   TransactionSourceACHTransferReturnReturnReasonCode = "invalid_ach_routing_number"
	TransactionSourceACHTransferReturnReturnReasonCodeFileRecordEditCriteria                                    TransactionSourceACHTransferReturnReturnReasonCode = "file_record_edit_criteria"
	TransactionSourceACHTransferReturnReturnReasonCodeEnrInvalidIndividualName                                  TransactionSourceACHTransferReturnReturnReasonCode = "enr_invalid_individual_name"
	TransactionSourceACHTransferReturnReturnReasonCodeReturnedPerOdfiRequest                                    TransactionSourceACHTransferReturnReturnReasonCode = "returned_per_odfi_request"
	TransactionSourceACHTransferReturnReturnReasonCodeAddendaError                                              TransactionSourceACHTransferReturnReturnReasonCode = "addenda_error"
	TransactionSourceACHTransferReturnReturnReasonCodeLimitedParticipationDfi                                   TransactionSourceACHTransferReturnReturnReasonCode = "limited_participation_dfi"
	TransactionSourceACHTransferReturnReturnReasonCodeOther                                                     TransactionSourceACHTransferReturnReturnReasonCode = "other"
)

//
type TransactionSourceCardDisputeAcceptance struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card Dispute was accepted.
	AcceptedAt *string `json:"accepted_at"`
	// The identifier of the Card Dispute that was accepted.
	CardDisputeID *string `json:"card_dispute_id"`
	// The identifier of the Transaction that was created to return the disputed funds
	// to your account.
	TransactionID *string `json:"transaction_id"`
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the Card Dispute was accepted.
func (r *TransactionSourceCardDisputeAcceptance) GetAcceptedAt() (AcceptedAt string) {
	if r != nil && r.AcceptedAt != nil {
		AcceptedAt = *r.AcceptedAt
	}
	return
}

// The identifier of the Card Dispute that was accepted.
func (r *TransactionSourceCardDisputeAcceptance) GetCardDisputeID() (CardDisputeID string) {
	if r != nil && r.CardDisputeID != nil {
		CardDisputeID = *r.CardDisputeID
	}
	return
}

// The identifier of the Transaction that was created to return the disputed funds
// to your account.
func (r *TransactionSourceCardDisputeAcceptance) GetTransactionID() (TransactionID string) {
	if r != nil && r.TransactionID != nil {
		TransactionID = *r.TransactionID
	}
	return
}

//
type TransactionSourceCardRefund struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *TransactionSourceCardRefundCurrency `json:"currency"`
	// The identifier for the Transaction this refunds, if any.
	CardSettlementTransactionID *string `json:"card_settlement_transaction_id"`
	// A constant representing the object's type. For this resource it will always be
	// `card_refund`.
	Type *TransactionSourceCardRefundType `json:"type"`
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *TransactionSourceCardRefund) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *TransactionSourceCardRefund) GetCurrency() (Currency TransactionSourceCardRefundCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The identifier for the Transaction this refunds, if any.
func (r *TransactionSourceCardRefund) GetCardSettlementTransactionID() (CardSettlementTransactionID string) {
	if r != nil && r.CardSettlementTransactionID != nil {
		CardSettlementTransactionID = *r.CardSettlementTransactionID
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `card_refund`.
func (r *TransactionSourceCardRefund) GetType() (Type TransactionSourceCardRefundType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

type TransactionSourceCardRefundCurrency string

const (
	TransactionSourceCardRefundCurrencyCad TransactionSourceCardRefundCurrency = "CAD"
	TransactionSourceCardRefundCurrencyChf TransactionSourceCardRefundCurrency = "CHF"
	TransactionSourceCardRefundCurrencyEur TransactionSourceCardRefundCurrency = "EUR"
	TransactionSourceCardRefundCurrencyGbp TransactionSourceCardRefundCurrency = "GBP"
	TransactionSourceCardRefundCurrencyJpy TransactionSourceCardRefundCurrency = "JPY"
	TransactionSourceCardRefundCurrencyUsd TransactionSourceCardRefundCurrency = "USD"
)

type TransactionSourceCardRefundType string

const (
	TransactionSourceCardRefundTypeCardRefund TransactionSourceCardRefundType = "card_refund"
)

//
type TransactionSourceCardSettlement struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *TransactionSourceCardSettlementCurrency `json:"currency"`
	//
	MerchantCity *string `json:"merchant_city"`
	//
	MerchantCountry *string `json:"merchant_country"`
	//
	MerchantName *string `json:"merchant_name"`
	//
	MerchantCategoryCode *string `json:"merchant_category_code"`
	//
	MerchantState *string `json:"merchant_state"`
	// The identifier of the Pending Transaction associated with this Transaction.
	PendingTransactionID *string `json:"pending_transaction_id"`
	// A constant representing the object's type. For this resource it will always be
	// `card_settlement`.
	Type *TransactionSourceCardSettlementType `json:"type"`
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *TransactionSourceCardSettlement) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *TransactionSourceCardSettlement) GetCurrency() (Currency TransactionSourceCardSettlementCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

func (r *TransactionSourceCardSettlement) GetMerchantCity() (MerchantCity string) {
	if r != nil && r.MerchantCity != nil {
		MerchantCity = *r.MerchantCity
	}
	return
}

func (r *TransactionSourceCardSettlement) GetMerchantCountry() (MerchantCountry string) {
	if r != nil && r.MerchantCountry != nil {
		MerchantCountry = *r.MerchantCountry
	}
	return
}

func (r *TransactionSourceCardSettlement) GetMerchantName() (MerchantName string) {
	if r != nil && r.MerchantName != nil {
		MerchantName = *r.MerchantName
	}
	return
}

func (r *TransactionSourceCardSettlement) GetMerchantCategoryCode() (MerchantCategoryCode string) {
	if r != nil && r.MerchantCategoryCode != nil {
		MerchantCategoryCode = *r.MerchantCategoryCode
	}
	return
}

func (r *TransactionSourceCardSettlement) GetMerchantState() (MerchantState string) {
	if r != nil && r.MerchantState != nil {
		MerchantState = *r.MerchantState
	}
	return
}

// The identifier of the Pending Transaction associated with this Transaction.
func (r *TransactionSourceCardSettlement) GetPendingTransactionID() (PendingTransactionID string) {
	if r != nil && r.PendingTransactionID != nil {
		PendingTransactionID = *r.PendingTransactionID
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `card_settlement`.
func (r *TransactionSourceCardSettlement) GetType() (Type TransactionSourceCardSettlementType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

type TransactionSourceCardSettlementCurrency string

const (
	TransactionSourceCardSettlementCurrencyCad TransactionSourceCardSettlementCurrency = "CAD"
	TransactionSourceCardSettlementCurrencyChf TransactionSourceCardSettlementCurrency = "CHF"
	TransactionSourceCardSettlementCurrencyEur TransactionSourceCardSettlementCurrency = "EUR"
	TransactionSourceCardSettlementCurrencyGbp TransactionSourceCardSettlementCurrency = "GBP"
	TransactionSourceCardSettlementCurrencyJpy TransactionSourceCardSettlementCurrency = "JPY"
	TransactionSourceCardSettlementCurrencyUsd TransactionSourceCardSettlementCurrency = "USD"
)

type TransactionSourceCardSettlementType string

const (
	TransactionSourceCardSettlementTypeCardSettlement TransactionSourceCardSettlementType = "card_settlement"
)

//
type TransactionSourceCheckDepositAcceptance struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *TransactionSourceCheckDepositAcceptanceCurrency `json:"currency"`
	// The ID of the Check Deposit that led to the Transaction.
	CheckDepositID *string `json:"check_deposit_id"`
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *TransactionSourceCheckDepositAcceptance) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *TransactionSourceCheckDepositAcceptance) GetCurrency() (Currency TransactionSourceCheckDepositAcceptanceCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The ID of the Check Deposit that led to the Transaction.
func (r *TransactionSourceCheckDepositAcceptance) GetCheckDepositID() (CheckDepositID string) {
	if r != nil && r.CheckDepositID != nil {
		CheckDepositID = *r.CheckDepositID
	}
	return
}

type TransactionSourceCheckDepositAcceptanceCurrency string

const (
	TransactionSourceCheckDepositAcceptanceCurrencyCad TransactionSourceCheckDepositAcceptanceCurrency = "CAD"
	TransactionSourceCheckDepositAcceptanceCurrencyChf TransactionSourceCheckDepositAcceptanceCurrency = "CHF"
	TransactionSourceCheckDepositAcceptanceCurrencyEur TransactionSourceCheckDepositAcceptanceCurrency = "EUR"
	TransactionSourceCheckDepositAcceptanceCurrencyGbp TransactionSourceCheckDepositAcceptanceCurrency = "GBP"
	TransactionSourceCheckDepositAcceptanceCurrencyJpy TransactionSourceCheckDepositAcceptanceCurrency = "JPY"
	TransactionSourceCheckDepositAcceptanceCurrencyUsd TransactionSourceCheckDepositAcceptanceCurrency = "USD"
)

//
type TransactionSourceCheckDepositReturn struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check deposit was returned.
	ReturnedAt *string `json:"returned_at"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *TransactionSourceCheckDepositReturnCurrency `json:"currency"`
	// The identifier of the Check Deposit that was returned.
	CheckDepositID *string `json:"check_deposit_id"`
	// The identifier of the transaction that reversed the original check deposit
	// transaction.
	TransactionID *string `json:"transaction_id"`
	//
	ReturnReason *TransactionSourceCheckDepositReturnReturnReason `json:"return_reason"`
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *TransactionSourceCheckDepositReturn) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the check deposit was returned.
func (r *TransactionSourceCheckDepositReturn) GetReturnedAt() (ReturnedAt string) {
	if r != nil && r.ReturnedAt != nil {
		ReturnedAt = *r.ReturnedAt
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *TransactionSourceCheckDepositReturn) GetCurrency() (Currency TransactionSourceCheckDepositReturnCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The identifier of the Check Deposit that was returned.
func (r *TransactionSourceCheckDepositReturn) GetCheckDepositID() (CheckDepositID string) {
	if r != nil && r.CheckDepositID != nil {
		CheckDepositID = *r.CheckDepositID
	}
	return
}

// The identifier of the transaction that reversed the original check deposit
// transaction.
func (r *TransactionSourceCheckDepositReturn) GetTransactionID() (TransactionID string) {
	if r != nil && r.TransactionID != nil {
		TransactionID = *r.TransactionID
	}
	return
}

func (r *TransactionSourceCheckDepositReturn) GetReturnReason() (ReturnReason TransactionSourceCheckDepositReturnReturnReason) {
	if r != nil && r.ReturnReason != nil {
		ReturnReason = *r.ReturnReason
	}
	return
}

type TransactionSourceCheckDepositReturnCurrency string

const (
	TransactionSourceCheckDepositReturnCurrencyCad TransactionSourceCheckDepositReturnCurrency = "CAD"
	TransactionSourceCheckDepositReturnCurrencyChf TransactionSourceCheckDepositReturnCurrency = "CHF"
	TransactionSourceCheckDepositReturnCurrencyEur TransactionSourceCheckDepositReturnCurrency = "EUR"
	TransactionSourceCheckDepositReturnCurrencyGbp TransactionSourceCheckDepositReturnCurrency = "GBP"
	TransactionSourceCheckDepositReturnCurrencyJpy TransactionSourceCheckDepositReturnCurrency = "JPY"
	TransactionSourceCheckDepositReturnCurrencyUsd TransactionSourceCheckDepositReturnCurrency = "USD"
)

type TransactionSourceCheckDepositReturnReturnReason string

const (
	TransactionSourceCheckDepositReturnReturnReasonACHConversionNotSupported TransactionSourceCheckDepositReturnReturnReason = "ach_conversion_not_supported"
	TransactionSourceCheckDepositReturnReturnReasonDuplicateSubmission       TransactionSourceCheckDepositReturnReturnReason = "duplicate_submission"
	TransactionSourceCheckDepositReturnReturnReasonInsufficientFunds         TransactionSourceCheckDepositReturnReturnReason = "insufficient_funds"
	TransactionSourceCheckDepositReturnReturnReasonNoAccount                 TransactionSourceCheckDepositReturnReturnReason = "no_account"
	TransactionSourceCheckDepositReturnReturnReasonNotAuthorized             TransactionSourceCheckDepositReturnReturnReason = "not_authorized"
	TransactionSourceCheckDepositReturnReturnReasonStaleDated                TransactionSourceCheckDepositReturnReturnReason = "stale_dated"
	TransactionSourceCheckDepositReturnReturnReasonStopPayment               TransactionSourceCheckDepositReturnReturnReason = "stop_payment"
	TransactionSourceCheckDepositReturnReturnReasonUnknownReason             TransactionSourceCheckDepositReturnReturnReason = "unknown_reason"
	TransactionSourceCheckDepositReturnReturnReasonUnmatchedDetails          TransactionSourceCheckDepositReturnReturnReason = "unmatched_details"
	TransactionSourceCheckDepositReturnReturnReasonUnreadableImage           TransactionSourceCheckDepositReturnReturnReason = "unreadable_image"
)

//
type TransactionSourceCheckTransferIntention struct {
	// The street address of the check's destination.
	AddressLine1 *string `json:"address_line1"`
	// The second line of the address of the check's destination.
	AddressLine2 *string `json:"address_line2"`
	// The city of the check's destination.
	AddressCity *string `json:"address_city"`
	// The state of the check's destination.
	AddressState *string `json:"address_state"`
	// The postal code of the check's destination.
	AddressZip *string `json:"address_zip"`
	// The transfer amount in USD cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
	// currency.
	Currency *TransactionSourceCheckTransferIntentionCurrency `json:"currency"`
	// The name that will be printed on the check.
	RecipientName *string `json:"recipient_name"`
	// The identifier of the Check Transfer with which this is associated.
	TransferID *string `json:"transfer_id"`
}

// The street address of the check's destination.
func (r *TransactionSourceCheckTransferIntention) GetAddressLine1() (AddressLine1 string) {
	if r != nil && r.AddressLine1 != nil {
		AddressLine1 = *r.AddressLine1
	}
	return
}

// The second line of the address of the check's destination.
func (r *TransactionSourceCheckTransferIntention) GetAddressLine2() (AddressLine2 string) {
	if r != nil && r.AddressLine2 != nil {
		AddressLine2 = *r.AddressLine2
	}
	return
}

// The city of the check's destination.
func (r *TransactionSourceCheckTransferIntention) GetAddressCity() (AddressCity string) {
	if r != nil && r.AddressCity != nil {
		AddressCity = *r.AddressCity
	}
	return
}

// The state of the check's destination.
func (r *TransactionSourceCheckTransferIntention) GetAddressState() (AddressState string) {
	if r != nil && r.AddressState != nil {
		AddressState = *r.AddressState
	}
	return
}

// The postal code of the check's destination.
func (r *TransactionSourceCheckTransferIntention) GetAddressZip() (AddressZip string) {
	if r != nil && r.AddressZip != nil {
		AddressZip = *r.AddressZip
	}
	return
}

// The transfer amount in USD cents.
func (r *TransactionSourceCheckTransferIntention) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
// currency.
func (r *TransactionSourceCheckTransferIntention) GetCurrency() (Currency TransactionSourceCheckTransferIntentionCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The name that will be printed on the check.
func (r *TransactionSourceCheckTransferIntention) GetRecipientName() (RecipientName string) {
	if r != nil && r.RecipientName != nil {
		RecipientName = *r.RecipientName
	}
	return
}

// The identifier of the Check Transfer with which this is associated.
func (r *TransactionSourceCheckTransferIntention) GetTransferID() (TransferID string) {
	if r != nil && r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

type TransactionSourceCheckTransferIntentionCurrency string

const (
	TransactionSourceCheckTransferIntentionCurrencyCad TransactionSourceCheckTransferIntentionCurrency = "CAD"
	TransactionSourceCheckTransferIntentionCurrencyChf TransactionSourceCheckTransferIntentionCurrency = "CHF"
	TransactionSourceCheckTransferIntentionCurrencyEur TransactionSourceCheckTransferIntentionCurrency = "EUR"
	TransactionSourceCheckTransferIntentionCurrencyGbp TransactionSourceCheckTransferIntentionCurrency = "GBP"
	TransactionSourceCheckTransferIntentionCurrencyJpy TransactionSourceCheckTransferIntentionCurrency = "JPY"
	TransactionSourceCheckTransferIntentionCurrencyUsd TransactionSourceCheckTransferIntentionCurrency = "USD"
)

//
type TransactionSourceCheckTransferRejection struct {
	// The identifier of the Check Transfer that led to this Transaction.
	TransferID *string `json:"transfer_id"`
}

// The identifier of the Check Transfer that led to this Transaction.
func (r *TransactionSourceCheckTransferRejection) GetTransferID() (TransferID string) {
	if r != nil && r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

//
type TransactionSourceCheckTransferStopPaymentRequest struct {
	// The ID of the check transfer that was stopped.
	TransferID *string `json:"transfer_id"`
	// The transaction ID of the corresponding credit transaction.
	TransactionID *string `json:"transaction_id"`
	// The time the stop-payment was requested.
	RequestedAt *string `json:"requested_at"`
	// A constant representing the object's type. For this resource it will always be
	// `check_transfer_stop_payment_request`.
	Type *TransactionSourceCheckTransferStopPaymentRequestType `json:"type"`
}

// The ID of the check transfer that was stopped.
func (r *TransactionSourceCheckTransferStopPaymentRequest) GetTransferID() (TransferID string) {
	if r != nil && r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

// The transaction ID of the corresponding credit transaction.
func (r *TransactionSourceCheckTransferStopPaymentRequest) GetTransactionID() (TransactionID string) {
	if r != nil && r.TransactionID != nil {
		TransactionID = *r.TransactionID
	}
	return
}

// The time the stop-payment was requested.
func (r *TransactionSourceCheckTransferStopPaymentRequest) GetRequestedAt() (RequestedAt string) {
	if r != nil && r.RequestedAt != nil {
		RequestedAt = *r.RequestedAt
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `check_transfer_stop_payment_request`.
func (r *TransactionSourceCheckTransferStopPaymentRequest) GetType() (Type TransactionSourceCheckTransferStopPaymentRequestType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

type TransactionSourceCheckTransferStopPaymentRequestType string

const (
	TransactionSourceCheckTransferStopPaymentRequestTypeCheckTransferStopPaymentRequest TransactionSourceCheckTransferStopPaymentRequestType = "check_transfer_stop_payment_request"
)

//
type TransactionSourceDisputeResolution struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *TransactionSourceDisputeResolutionCurrency `json:"currency"`
	// The identifier of the Transaction that was disputed.
	DisputedTransactionID *string `json:"disputed_transaction_id"`
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *TransactionSourceDisputeResolution) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *TransactionSourceDisputeResolution) GetCurrency() (Currency TransactionSourceDisputeResolutionCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The identifier of the Transaction that was disputed.
func (r *TransactionSourceDisputeResolution) GetDisputedTransactionID() (DisputedTransactionID string) {
	if r != nil && r.DisputedTransactionID != nil {
		DisputedTransactionID = *r.DisputedTransactionID
	}
	return
}

type TransactionSourceDisputeResolutionCurrency string

const (
	TransactionSourceDisputeResolutionCurrencyCad TransactionSourceDisputeResolutionCurrency = "CAD"
	TransactionSourceDisputeResolutionCurrencyChf TransactionSourceDisputeResolutionCurrency = "CHF"
	TransactionSourceDisputeResolutionCurrencyEur TransactionSourceDisputeResolutionCurrency = "EUR"
	TransactionSourceDisputeResolutionCurrencyGbp TransactionSourceDisputeResolutionCurrency = "GBP"
	TransactionSourceDisputeResolutionCurrencyJpy TransactionSourceDisputeResolutionCurrency = "JPY"
	TransactionSourceDisputeResolutionCurrencyUsd TransactionSourceDisputeResolutionCurrency = "USD"
)

//
type TransactionSourceEmpyrealCashDeposit struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int `json:"amount"`
	//
	BagID *string `json:"bag_id"`
	//
	DepositDate *string `json:"deposit_date"`
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *TransactionSourceEmpyrealCashDeposit) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *TransactionSourceEmpyrealCashDeposit) GetBagID() (BagID string) {
	if r != nil && r.BagID != nil {
		BagID = *r.BagID
	}
	return
}

func (r *TransactionSourceEmpyrealCashDeposit) GetDepositDate() (DepositDate string) {
	if r != nil && r.DepositDate != nil {
		DepositDate = *r.DepositDate
	}
	return
}

//
type TransactionSourceInboundACHTransfer struct {
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
	OriginatorCompanyEntryDescription *string `json:"originator_company_entry_description"`
	//
	OriginatorCompanyID *string `json:"originator_company_id"`
	//
	ReceiverIDNumber *string `json:"receiver_id_number"`
	//
	ReceiverName *string `json:"receiver_name"`
	//
	TraceNumber *string `json:"trace_number"`
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *TransactionSourceInboundACHTransfer) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *TransactionSourceInboundACHTransfer) GetOriginatorCompanyName() (OriginatorCompanyName string) {
	if r != nil && r.OriginatorCompanyName != nil {
		OriginatorCompanyName = *r.OriginatorCompanyName
	}
	return
}

func (r *TransactionSourceInboundACHTransfer) GetOriginatorCompanyDescriptiveDate() (OriginatorCompanyDescriptiveDate string) {
	if r != nil && r.OriginatorCompanyDescriptiveDate != nil {
		OriginatorCompanyDescriptiveDate = *r.OriginatorCompanyDescriptiveDate
	}
	return
}

func (r *TransactionSourceInboundACHTransfer) GetOriginatorCompanyDiscretionaryData() (OriginatorCompanyDiscretionaryData string) {
	if r != nil && r.OriginatorCompanyDiscretionaryData != nil {
		OriginatorCompanyDiscretionaryData = *r.OriginatorCompanyDiscretionaryData
	}
	return
}

func (r *TransactionSourceInboundACHTransfer) GetOriginatorCompanyEntryDescription() (OriginatorCompanyEntryDescription string) {
	if r != nil && r.OriginatorCompanyEntryDescription != nil {
		OriginatorCompanyEntryDescription = *r.OriginatorCompanyEntryDescription
	}
	return
}

func (r *TransactionSourceInboundACHTransfer) GetOriginatorCompanyID() (OriginatorCompanyID string) {
	if r != nil && r.OriginatorCompanyID != nil {
		OriginatorCompanyID = *r.OriginatorCompanyID
	}
	return
}

func (r *TransactionSourceInboundACHTransfer) GetReceiverIDNumber() (ReceiverIDNumber string) {
	if r != nil && r.ReceiverIDNumber != nil {
		ReceiverIDNumber = *r.ReceiverIDNumber
	}
	return
}

func (r *TransactionSourceInboundACHTransfer) GetReceiverName() (ReceiverName string) {
	if r != nil && r.ReceiverName != nil {
		ReceiverName = *r.ReceiverName
	}
	return
}

func (r *TransactionSourceInboundACHTransfer) GetTraceNumber() (TraceNumber string) {
	if r != nil && r.TraceNumber != nil {
		TraceNumber = *r.TraceNumber
	}
	return
}

//
type TransactionSourceInboundCheck struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *TransactionSourceInboundCheckCurrency `json:"currency"`
	//
	CheckNumber *string `json:"check_number"`
	//
	CheckFrontImageFileID *string `json:"check_front_image_file_id"`
	//
	CheckRearImageFileID *string `json:"check_rear_image_file_id"`
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *TransactionSourceInboundCheck) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *TransactionSourceInboundCheck) GetCurrency() (Currency TransactionSourceInboundCheckCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

func (r *TransactionSourceInboundCheck) GetCheckNumber() (CheckNumber string) {
	if r != nil && r.CheckNumber != nil {
		CheckNumber = *r.CheckNumber
	}
	return
}

func (r *TransactionSourceInboundCheck) GetCheckFrontImageFileID() (CheckFrontImageFileID string) {
	if r != nil && r.CheckFrontImageFileID != nil {
		CheckFrontImageFileID = *r.CheckFrontImageFileID
	}
	return
}

func (r *TransactionSourceInboundCheck) GetCheckRearImageFileID() (CheckRearImageFileID string) {
	if r != nil && r.CheckRearImageFileID != nil {
		CheckRearImageFileID = *r.CheckRearImageFileID
	}
	return
}

type TransactionSourceInboundCheckCurrency string

const (
	TransactionSourceInboundCheckCurrencyCad TransactionSourceInboundCheckCurrency = "CAD"
	TransactionSourceInboundCheckCurrencyChf TransactionSourceInboundCheckCurrency = "CHF"
	TransactionSourceInboundCheckCurrencyEur TransactionSourceInboundCheckCurrency = "EUR"
	TransactionSourceInboundCheckCurrencyGbp TransactionSourceInboundCheckCurrency = "GBP"
	TransactionSourceInboundCheckCurrencyJpy TransactionSourceInboundCheckCurrency = "JPY"
	TransactionSourceInboundCheckCurrencyUsd TransactionSourceInboundCheckCurrency = "USD"
)

//
type TransactionSourceInboundInternationalACHTransfer struct {
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
func (r *TransactionSourceInboundInternationalACHTransfer) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetForeignExchangeIndicator() (ForeignExchangeIndicator string) {
	if r != nil && r.ForeignExchangeIndicator != nil {
		ForeignExchangeIndicator = *r.ForeignExchangeIndicator
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetForeignExchangeReferenceIndicator() (ForeignExchangeReferenceIndicator string) {
	if r != nil && r.ForeignExchangeReferenceIndicator != nil {
		ForeignExchangeReferenceIndicator = *r.ForeignExchangeReferenceIndicator
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetForeignExchangeReference() (ForeignExchangeReference string) {
	if r != nil && r.ForeignExchangeReference != nil {
		ForeignExchangeReference = *r.ForeignExchangeReference
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetDestinationCountryCode() (DestinationCountryCode string) {
	if r != nil && r.DestinationCountryCode != nil {
		DestinationCountryCode = *r.DestinationCountryCode
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetDestinationCurrencyCode() (DestinationCurrencyCode string) {
	if r != nil && r.DestinationCurrencyCode != nil {
		DestinationCurrencyCode = *r.DestinationCurrencyCode
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetForeignPaymentAmount() (ForeignPaymentAmount int) {
	if r != nil && r.ForeignPaymentAmount != nil {
		ForeignPaymentAmount = *r.ForeignPaymentAmount
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetForeignTraceNumber() (ForeignTraceNumber string) {
	if r != nil && r.ForeignTraceNumber != nil {
		ForeignTraceNumber = *r.ForeignTraceNumber
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetInternationalTransactionTypeCode() (InternationalTransactionTypeCode string) {
	if r != nil && r.InternationalTransactionTypeCode != nil {
		InternationalTransactionTypeCode = *r.InternationalTransactionTypeCode
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetOriginatingCurrencyCode() (OriginatingCurrencyCode string) {
	if r != nil && r.OriginatingCurrencyCode != nil {
		OriginatingCurrencyCode = *r.OriginatingCurrencyCode
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetOriginatingDepositoryFinancialInstitutionName() (OriginatingDepositoryFinancialInstitutionName string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionName != nil {
		OriginatingDepositoryFinancialInstitutionName = *r.OriginatingDepositoryFinancialInstitutionName
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetOriginatingDepositoryFinancialInstitutionIDQualifier() (OriginatingDepositoryFinancialInstitutionIDQualifier string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionIDQualifier != nil {
		OriginatingDepositoryFinancialInstitutionIDQualifier = *r.OriginatingDepositoryFinancialInstitutionIDQualifier
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetOriginatingDepositoryFinancialInstitutionID() (OriginatingDepositoryFinancialInstitutionID string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionID != nil {
		OriginatingDepositoryFinancialInstitutionID = *r.OriginatingDepositoryFinancialInstitutionID
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetOriginatingDepositoryFinancialInstitutionBranchCountry() (OriginatingDepositoryFinancialInstitutionBranchCountry string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionBranchCountry != nil {
		OriginatingDepositoryFinancialInstitutionBranchCountry = *r.OriginatingDepositoryFinancialInstitutionBranchCountry
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetOriginatorCity() (OriginatorCity string) {
	if r != nil && r.OriginatorCity != nil {
		OriginatorCity = *r.OriginatorCity
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetOriginatorCompanyEntryDescription() (OriginatorCompanyEntryDescription string) {
	if r != nil && r.OriginatorCompanyEntryDescription != nil {
		OriginatorCompanyEntryDescription = *r.OriginatorCompanyEntryDescription
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetOriginatorCountry() (OriginatorCountry string) {
	if r != nil && r.OriginatorCountry != nil {
		OriginatorCountry = *r.OriginatorCountry
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetOriginatorIdentification() (OriginatorIdentification string) {
	if r != nil && r.OriginatorIdentification != nil {
		OriginatorIdentification = *r.OriginatorIdentification
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetOriginatorName() (OriginatorName string) {
	if r != nil && r.OriginatorName != nil {
		OriginatorName = *r.OriginatorName
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetOriginatorPostalCode() (OriginatorPostalCode string) {
	if r != nil && r.OriginatorPostalCode != nil {
		OriginatorPostalCode = *r.OriginatorPostalCode
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetOriginatorStreetAddress() (OriginatorStreetAddress string) {
	if r != nil && r.OriginatorStreetAddress != nil {
		OriginatorStreetAddress = *r.OriginatorStreetAddress
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetOriginatorStateOrProvince() (OriginatorStateOrProvince string) {
	if r != nil && r.OriginatorStateOrProvince != nil {
		OriginatorStateOrProvince = *r.OriginatorStateOrProvince
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetPaymentRelatedInformation() (PaymentRelatedInformation string) {
	if r != nil && r.PaymentRelatedInformation != nil {
		PaymentRelatedInformation = *r.PaymentRelatedInformation
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetPaymentRelatedInformation2() (PaymentRelatedInformation2 string) {
	if r != nil && r.PaymentRelatedInformation2 != nil {
		PaymentRelatedInformation2 = *r.PaymentRelatedInformation2
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetReceiverIdentificationNumber() (ReceiverIdentificationNumber string) {
	if r != nil && r.ReceiverIdentificationNumber != nil {
		ReceiverIdentificationNumber = *r.ReceiverIdentificationNumber
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetReceiverStreetAddress() (ReceiverStreetAddress string) {
	if r != nil && r.ReceiverStreetAddress != nil {
		ReceiverStreetAddress = *r.ReceiverStreetAddress
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetReceiverCity() (ReceiverCity string) {
	if r != nil && r.ReceiverCity != nil {
		ReceiverCity = *r.ReceiverCity
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetReceiverStateOrProvince() (ReceiverStateOrProvince string) {
	if r != nil && r.ReceiverStateOrProvince != nil {
		ReceiverStateOrProvince = *r.ReceiverStateOrProvince
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetReceiverCountry() (ReceiverCountry string) {
	if r != nil && r.ReceiverCountry != nil {
		ReceiverCountry = *r.ReceiverCountry
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetReceiverPostalCode() (ReceiverPostalCode string) {
	if r != nil && r.ReceiverPostalCode != nil {
		ReceiverPostalCode = *r.ReceiverPostalCode
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetReceivingCompanyOrIndividualName() (ReceivingCompanyOrIndividualName string) {
	if r != nil && r.ReceivingCompanyOrIndividualName != nil {
		ReceivingCompanyOrIndividualName = *r.ReceivingCompanyOrIndividualName
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetReceivingDepositoryFinancialInstitutionName() (ReceivingDepositoryFinancialInstitutionName string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionName != nil {
		ReceivingDepositoryFinancialInstitutionName = *r.ReceivingDepositoryFinancialInstitutionName
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetReceivingDepositoryFinancialInstitutionIDQualifier() (ReceivingDepositoryFinancialInstitutionIDQualifier string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionIDQualifier != nil {
		ReceivingDepositoryFinancialInstitutionIDQualifier = *r.ReceivingDepositoryFinancialInstitutionIDQualifier
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetReceivingDepositoryFinancialInstitutionID() (ReceivingDepositoryFinancialInstitutionID string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionID != nil {
		ReceivingDepositoryFinancialInstitutionID = *r.ReceivingDepositoryFinancialInstitutionID
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetReceivingDepositoryFinancialInstitutionCountry() (ReceivingDepositoryFinancialInstitutionCountry string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionCountry != nil {
		ReceivingDepositoryFinancialInstitutionCountry = *r.ReceivingDepositoryFinancialInstitutionCountry
	}
	return
}

func (r *TransactionSourceInboundInternationalACHTransfer) GetTraceNumber() (TraceNumber string) {
	if r != nil && r.TraceNumber != nil {
		TraceNumber = *r.TraceNumber
	}
	return
}

//
type TransactionSourceInboundRealTimePaymentsTransferConfirmation struct {
	// The amount in the minor unit of the transfer's currency. For dollars, for
	// example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the transfer's
	// currency. This will always be "USD" for a Real Time Payments transfer.
	Currency *TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency `json:"currency"`
	// The name the sender of the transfer specified as the recipient of the transfer.
	CreditorName *string `json:"creditor_name"`
	// The name provided by the sender of the transfer.
	DebtorName *string `json:"debtor_name"`
	// The account number of the account that sent the transfer.
	DebtorAccountNumber *string `json:"debtor_account_number"`
	// The routing number of the account that sent the transfer.
	DebtorRoutingNumber *string `json:"debtor_routing_number"`
	// The Real Time Payments network identification of the transfer
	TransactionIdentification *string `json:"transaction_identification"`
	// Additional information included with the transfer.
	RemittanceInformation *string `json:"remittance_information"`
}

// The amount in the minor unit of the transfer's currency. For dollars, for
// example, this is cents.
func (r *TransactionSourceInboundRealTimePaymentsTransferConfirmation) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the transfer's
// currency. This will always be "USD" for a Real Time Payments transfer.
func (r *TransactionSourceInboundRealTimePaymentsTransferConfirmation) GetCurrency() (Currency TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The name the sender of the transfer specified as the recipient of the transfer.
func (r *TransactionSourceInboundRealTimePaymentsTransferConfirmation) GetCreditorName() (CreditorName string) {
	if r != nil && r.CreditorName != nil {
		CreditorName = *r.CreditorName
	}
	return
}

// The name provided by the sender of the transfer.
func (r *TransactionSourceInboundRealTimePaymentsTransferConfirmation) GetDebtorName() (DebtorName string) {
	if r != nil && r.DebtorName != nil {
		DebtorName = *r.DebtorName
	}
	return
}

// The account number of the account that sent the transfer.
func (r *TransactionSourceInboundRealTimePaymentsTransferConfirmation) GetDebtorAccountNumber() (DebtorAccountNumber string) {
	if r != nil && r.DebtorAccountNumber != nil {
		DebtorAccountNumber = *r.DebtorAccountNumber
	}
	return
}

// The routing number of the account that sent the transfer.
func (r *TransactionSourceInboundRealTimePaymentsTransferConfirmation) GetDebtorRoutingNumber() (DebtorRoutingNumber string) {
	if r != nil && r.DebtorRoutingNumber != nil {
		DebtorRoutingNumber = *r.DebtorRoutingNumber
	}
	return
}

// The Real Time Payments network identification of the transfer
func (r *TransactionSourceInboundRealTimePaymentsTransferConfirmation) GetTransactionIdentification() (TransactionIdentification string) {
	if r != nil && r.TransactionIdentification != nil {
		TransactionIdentification = *r.TransactionIdentification
	}
	return
}

// Additional information included with the transfer.
func (r *TransactionSourceInboundRealTimePaymentsTransferConfirmation) GetRemittanceInformation() (RemittanceInformation string) {
	if r != nil && r.RemittanceInformation != nil {
		RemittanceInformation = *r.RemittanceInformation
	}
	return
}

type TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency string

const (
	TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyCad TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "CAD"
	TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyChf TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "CHF"
	TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyEur TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "EUR"
	TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyGbp TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "GBP"
	TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyJpy TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "JPY"
	TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyUsd TransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "USD"
)

//
type TransactionSourceInboundWireDrawdownPaymentReversal struct {
	// The amount that was reversed.
	Amount *int `json:"amount"`
	// The description on the reversal message from Fedwire.
	Description *string `json:"description"`
	// The Fedwire cycle date for the wire reversal.
	InputCycleDate *string `json:"input_cycle_date"`
	// The Fedwire sequence number.
	InputSequenceNumber *string `json:"input_sequence_number"`
	// The Fedwire input source identifier.
	InputSource *string `json:"input_source"`
	// The Fedwire transaction identifier.
	InputMessageAccountabilityData *string `json:"input_message_accountability_data"`
	// The Fedwire transaction identifier for the wire transfer that was reversed.
	PreviousMessageInputMessageAccountabilityData *string `json:"previous_message_input_message_accountability_data"`
	// The Fedwire cycle date for the wire transfer that was reversed.
	PreviousMessageInputCycleDate *string `json:"previous_message_input_cycle_date"`
	// The Fedwire sequence number for the wire transfer that was reversed.
	PreviousMessageInputSequenceNumber *string `json:"previous_message_input_sequence_number"`
	// The Fedwire input source identifier for the wire transfer that was reversed.
	PreviousMessageInputSource *string `json:"previous_message_input_source"`
}

// The amount that was reversed.
func (r *TransactionSourceInboundWireDrawdownPaymentReversal) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The description on the reversal message from Fedwire.
func (r *TransactionSourceInboundWireDrawdownPaymentReversal) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The Fedwire cycle date for the wire reversal.
func (r *TransactionSourceInboundWireDrawdownPaymentReversal) GetInputCycleDate() (InputCycleDate string) {
	if r != nil && r.InputCycleDate != nil {
		InputCycleDate = *r.InputCycleDate
	}
	return
}

// The Fedwire sequence number.
func (r *TransactionSourceInboundWireDrawdownPaymentReversal) GetInputSequenceNumber() (InputSequenceNumber string) {
	if r != nil && r.InputSequenceNumber != nil {
		InputSequenceNumber = *r.InputSequenceNumber
	}
	return
}

// The Fedwire input source identifier.
func (r *TransactionSourceInboundWireDrawdownPaymentReversal) GetInputSource() (InputSource string) {
	if r != nil && r.InputSource != nil {
		InputSource = *r.InputSource
	}
	return
}

// The Fedwire transaction identifier.
func (r *TransactionSourceInboundWireDrawdownPaymentReversal) GetInputMessageAccountabilityData() (InputMessageAccountabilityData string) {
	if r != nil && r.InputMessageAccountabilityData != nil {
		InputMessageAccountabilityData = *r.InputMessageAccountabilityData
	}
	return
}

// The Fedwire transaction identifier for the wire transfer that was reversed.
func (r *TransactionSourceInboundWireDrawdownPaymentReversal) GetPreviousMessageInputMessageAccountabilityData() (PreviousMessageInputMessageAccountabilityData string) {
	if r != nil && r.PreviousMessageInputMessageAccountabilityData != nil {
		PreviousMessageInputMessageAccountabilityData = *r.PreviousMessageInputMessageAccountabilityData
	}
	return
}

// The Fedwire cycle date for the wire transfer that was reversed.
func (r *TransactionSourceInboundWireDrawdownPaymentReversal) GetPreviousMessageInputCycleDate() (PreviousMessageInputCycleDate string) {
	if r != nil && r.PreviousMessageInputCycleDate != nil {
		PreviousMessageInputCycleDate = *r.PreviousMessageInputCycleDate
	}
	return
}

// The Fedwire sequence number for the wire transfer that was reversed.
func (r *TransactionSourceInboundWireDrawdownPaymentReversal) GetPreviousMessageInputSequenceNumber() (PreviousMessageInputSequenceNumber string) {
	if r != nil && r.PreviousMessageInputSequenceNumber != nil {
		PreviousMessageInputSequenceNumber = *r.PreviousMessageInputSequenceNumber
	}
	return
}

// The Fedwire input source identifier for the wire transfer that was reversed.
func (r *TransactionSourceInboundWireDrawdownPaymentReversal) GetPreviousMessageInputSource() (PreviousMessageInputSource string) {
	if r != nil && r.PreviousMessageInputSource != nil {
		PreviousMessageInputSource = *r.PreviousMessageInputSource
	}
	return
}

//
type TransactionSourceInboundWireDrawdownPayment struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int `json:"amount"`
	//
	BeneficiaryAddressLine1 *string `json:"beneficiary_address_line1"`
	//
	BeneficiaryAddressLine2 *string `json:"beneficiary_address_line2"`
	//
	BeneficiaryAddressLine3 *string `json:"beneficiary_address_line3"`
	//
	BeneficiaryName *string `json:"beneficiary_name"`
	//
	BeneficiaryReference *string `json:"beneficiary_reference"`
	//
	Description *string `json:"description"`
	//
	InputMessageAccountabilityData *string `json:"input_message_accountability_data"`
	//
	OriginatorAddressLine1 *string `json:"originator_address_line1"`
	//
	OriginatorAddressLine2 *string `json:"originator_address_line2"`
	//
	OriginatorAddressLine3 *string `json:"originator_address_line3"`
	//
	OriginatorName *string `json:"originator_name"`
	//
	OriginatorToBeneficiaryInformation *string `json:"originator_to_beneficiary_information"`
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *TransactionSourceInboundWireDrawdownPayment) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *TransactionSourceInboundWireDrawdownPayment) GetBeneficiaryAddressLine1() (BeneficiaryAddressLine1 string) {
	if r != nil && r.BeneficiaryAddressLine1 != nil {
		BeneficiaryAddressLine1 = *r.BeneficiaryAddressLine1
	}
	return
}

func (r *TransactionSourceInboundWireDrawdownPayment) GetBeneficiaryAddressLine2() (BeneficiaryAddressLine2 string) {
	if r != nil && r.BeneficiaryAddressLine2 != nil {
		BeneficiaryAddressLine2 = *r.BeneficiaryAddressLine2
	}
	return
}

func (r *TransactionSourceInboundWireDrawdownPayment) GetBeneficiaryAddressLine3() (BeneficiaryAddressLine3 string) {
	if r != nil && r.BeneficiaryAddressLine3 != nil {
		BeneficiaryAddressLine3 = *r.BeneficiaryAddressLine3
	}
	return
}

func (r *TransactionSourceInboundWireDrawdownPayment) GetBeneficiaryName() (BeneficiaryName string) {
	if r != nil && r.BeneficiaryName != nil {
		BeneficiaryName = *r.BeneficiaryName
	}
	return
}

func (r *TransactionSourceInboundWireDrawdownPayment) GetBeneficiaryReference() (BeneficiaryReference string) {
	if r != nil && r.BeneficiaryReference != nil {
		BeneficiaryReference = *r.BeneficiaryReference
	}
	return
}

func (r *TransactionSourceInboundWireDrawdownPayment) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

func (r *TransactionSourceInboundWireDrawdownPayment) GetInputMessageAccountabilityData() (InputMessageAccountabilityData string) {
	if r != nil && r.InputMessageAccountabilityData != nil {
		InputMessageAccountabilityData = *r.InputMessageAccountabilityData
	}
	return
}

func (r *TransactionSourceInboundWireDrawdownPayment) GetOriginatorAddressLine1() (OriginatorAddressLine1 string) {
	if r != nil && r.OriginatorAddressLine1 != nil {
		OriginatorAddressLine1 = *r.OriginatorAddressLine1
	}
	return
}

func (r *TransactionSourceInboundWireDrawdownPayment) GetOriginatorAddressLine2() (OriginatorAddressLine2 string) {
	if r != nil && r.OriginatorAddressLine2 != nil {
		OriginatorAddressLine2 = *r.OriginatorAddressLine2
	}
	return
}

func (r *TransactionSourceInboundWireDrawdownPayment) GetOriginatorAddressLine3() (OriginatorAddressLine3 string) {
	if r != nil && r.OriginatorAddressLine3 != nil {
		OriginatorAddressLine3 = *r.OriginatorAddressLine3
	}
	return
}

func (r *TransactionSourceInboundWireDrawdownPayment) GetOriginatorName() (OriginatorName string) {
	if r != nil && r.OriginatorName != nil {
		OriginatorName = *r.OriginatorName
	}
	return
}

func (r *TransactionSourceInboundWireDrawdownPayment) GetOriginatorToBeneficiaryInformation() (OriginatorToBeneficiaryInformation string) {
	if r != nil && r.OriginatorToBeneficiaryInformation != nil {
		OriginatorToBeneficiaryInformation = *r.OriginatorToBeneficiaryInformation
	}
	return
}

//
type TransactionSourceInboundWireReversal struct {
	// The amount that was reversed.
	Amount *int `json:"amount"`
	// The description on the reversal message from Fedwire.
	Description *string `json:"description"`
	// The Fedwire cycle date for the wire reversal.
	InputCycleDate *string `json:"input_cycle_date"`
	// The Fedwire sequence number.
	InputSequenceNumber *string `json:"input_sequence_number"`
	// The Fedwire input source identifier.
	InputSource *string `json:"input_source"`
	// The Fedwire transaction identifier.
	InputMessageAccountabilityData *string `json:"input_message_accountability_data"`
	// The Fedwire transaction identifier for the wire transfer that was reversed.
	PreviousMessageInputMessageAccountabilityData *string `json:"previous_message_input_message_accountability_data"`
	// The Fedwire cycle date for the wire transfer that was reversed.
	PreviousMessageInputCycleDate *string `json:"previous_message_input_cycle_date"`
	// The Fedwire sequence number for the wire transfer that was reversed.
	PreviousMessageInputSequenceNumber *string `json:"previous_message_input_sequence_number"`
	// The Fedwire input source identifier for the wire transfer that was reversed.
	PreviousMessageInputSource *string `json:"previous_message_input_source"`
	// Information included in the wire reversal for the receiving financial
	// institution.
	ReceiverFinancialInstitutionInformation *string `json:"receiver_financial_institution_information"`
	// Additional financial institution information included in the wire reversal.
	FinancialInstitutionToFinancialInstitutionInformation *string `json:"financial_institution_to_financial_institution_information"`
}

// The amount that was reversed.
func (r *TransactionSourceInboundWireReversal) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The description on the reversal message from Fedwire.
func (r *TransactionSourceInboundWireReversal) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The Fedwire cycle date for the wire reversal.
func (r *TransactionSourceInboundWireReversal) GetInputCycleDate() (InputCycleDate string) {
	if r != nil && r.InputCycleDate != nil {
		InputCycleDate = *r.InputCycleDate
	}
	return
}

// The Fedwire sequence number.
func (r *TransactionSourceInboundWireReversal) GetInputSequenceNumber() (InputSequenceNumber string) {
	if r != nil && r.InputSequenceNumber != nil {
		InputSequenceNumber = *r.InputSequenceNumber
	}
	return
}

// The Fedwire input source identifier.
func (r *TransactionSourceInboundWireReversal) GetInputSource() (InputSource string) {
	if r != nil && r.InputSource != nil {
		InputSource = *r.InputSource
	}
	return
}

// The Fedwire transaction identifier.
func (r *TransactionSourceInboundWireReversal) GetInputMessageAccountabilityData() (InputMessageAccountabilityData string) {
	if r != nil && r.InputMessageAccountabilityData != nil {
		InputMessageAccountabilityData = *r.InputMessageAccountabilityData
	}
	return
}

// The Fedwire transaction identifier for the wire transfer that was reversed.
func (r *TransactionSourceInboundWireReversal) GetPreviousMessageInputMessageAccountabilityData() (PreviousMessageInputMessageAccountabilityData string) {
	if r != nil && r.PreviousMessageInputMessageAccountabilityData != nil {
		PreviousMessageInputMessageAccountabilityData = *r.PreviousMessageInputMessageAccountabilityData
	}
	return
}

// The Fedwire cycle date for the wire transfer that was reversed.
func (r *TransactionSourceInboundWireReversal) GetPreviousMessageInputCycleDate() (PreviousMessageInputCycleDate string) {
	if r != nil && r.PreviousMessageInputCycleDate != nil {
		PreviousMessageInputCycleDate = *r.PreviousMessageInputCycleDate
	}
	return
}

// The Fedwire sequence number for the wire transfer that was reversed.
func (r *TransactionSourceInboundWireReversal) GetPreviousMessageInputSequenceNumber() (PreviousMessageInputSequenceNumber string) {
	if r != nil && r.PreviousMessageInputSequenceNumber != nil {
		PreviousMessageInputSequenceNumber = *r.PreviousMessageInputSequenceNumber
	}
	return
}

// The Fedwire input source identifier for the wire transfer that was reversed.
func (r *TransactionSourceInboundWireReversal) GetPreviousMessageInputSource() (PreviousMessageInputSource string) {
	if r != nil && r.PreviousMessageInputSource != nil {
		PreviousMessageInputSource = *r.PreviousMessageInputSource
	}
	return
}

// Information included in the wire reversal for the receiving financial
// institution.
func (r *TransactionSourceInboundWireReversal) GetReceiverFinancialInstitutionInformation() (ReceiverFinancialInstitutionInformation string) {
	if r != nil && r.ReceiverFinancialInstitutionInformation != nil {
		ReceiverFinancialInstitutionInformation = *r.ReceiverFinancialInstitutionInformation
	}
	return
}

// Additional financial institution information included in the wire reversal.
func (r *TransactionSourceInboundWireReversal) GetFinancialInstitutionToFinancialInstitutionInformation() (FinancialInstitutionToFinancialInstitutionInformation string) {
	if r != nil && r.FinancialInstitutionToFinancialInstitutionInformation != nil {
		FinancialInstitutionToFinancialInstitutionInformation = *r.FinancialInstitutionToFinancialInstitutionInformation
	}
	return
}

//
type TransactionSourceInboundWireTransfer struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int `json:"amount"`
	//
	BeneficiaryAddressLine1 *string `json:"beneficiary_address_line1"`
	//
	BeneficiaryAddressLine2 *string `json:"beneficiary_address_line2"`
	//
	BeneficiaryAddressLine3 *string `json:"beneficiary_address_line3"`
	//
	BeneficiaryName *string `json:"beneficiary_name"`
	//
	BeneficiaryReference *string `json:"beneficiary_reference"`
	//
	Description *string `json:"description"`
	//
	InputMessageAccountabilityData *string `json:"input_message_accountability_data"`
	//
	OriginatorAddressLine1 *string `json:"originator_address_line1"`
	//
	OriginatorAddressLine2 *string `json:"originator_address_line2"`
	//
	OriginatorAddressLine3 *string `json:"originator_address_line3"`
	//
	OriginatorName *string `json:"originator_name"`
	//
	OriginatorToBeneficiaryInformation *string `json:"originator_to_beneficiary_information"`
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *TransactionSourceInboundWireTransfer) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *TransactionSourceInboundWireTransfer) GetBeneficiaryAddressLine1() (BeneficiaryAddressLine1 string) {
	if r != nil && r.BeneficiaryAddressLine1 != nil {
		BeneficiaryAddressLine1 = *r.BeneficiaryAddressLine1
	}
	return
}

func (r *TransactionSourceInboundWireTransfer) GetBeneficiaryAddressLine2() (BeneficiaryAddressLine2 string) {
	if r != nil && r.BeneficiaryAddressLine2 != nil {
		BeneficiaryAddressLine2 = *r.BeneficiaryAddressLine2
	}
	return
}

func (r *TransactionSourceInboundWireTransfer) GetBeneficiaryAddressLine3() (BeneficiaryAddressLine3 string) {
	if r != nil && r.BeneficiaryAddressLine3 != nil {
		BeneficiaryAddressLine3 = *r.BeneficiaryAddressLine3
	}
	return
}

func (r *TransactionSourceInboundWireTransfer) GetBeneficiaryName() (BeneficiaryName string) {
	if r != nil && r.BeneficiaryName != nil {
		BeneficiaryName = *r.BeneficiaryName
	}
	return
}

func (r *TransactionSourceInboundWireTransfer) GetBeneficiaryReference() (BeneficiaryReference string) {
	if r != nil && r.BeneficiaryReference != nil {
		BeneficiaryReference = *r.BeneficiaryReference
	}
	return
}

func (r *TransactionSourceInboundWireTransfer) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

func (r *TransactionSourceInboundWireTransfer) GetInputMessageAccountabilityData() (InputMessageAccountabilityData string) {
	if r != nil && r.InputMessageAccountabilityData != nil {
		InputMessageAccountabilityData = *r.InputMessageAccountabilityData
	}
	return
}

func (r *TransactionSourceInboundWireTransfer) GetOriginatorAddressLine1() (OriginatorAddressLine1 string) {
	if r != nil && r.OriginatorAddressLine1 != nil {
		OriginatorAddressLine1 = *r.OriginatorAddressLine1
	}
	return
}

func (r *TransactionSourceInboundWireTransfer) GetOriginatorAddressLine2() (OriginatorAddressLine2 string) {
	if r != nil && r.OriginatorAddressLine2 != nil {
		OriginatorAddressLine2 = *r.OriginatorAddressLine2
	}
	return
}

func (r *TransactionSourceInboundWireTransfer) GetOriginatorAddressLine3() (OriginatorAddressLine3 string) {
	if r != nil && r.OriginatorAddressLine3 != nil {
		OriginatorAddressLine3 = *r.OriginatorAddressLine3
	}
	return
}

func (r *TransactionSourceInboundWireTransfer) GetOriginatorName() (OriginatorName string) {
	if r != nil && r.OriginatorName != nil {
		OriginatorName = *r.OriginatorName
	}
	return
}

func (r *TransactionSourceInboundWireTransfer) GetOriginatorToBeneficiaryInformation() (OriginatorToBeneficiaryInformation string) {
	if r != nil && r.OriginatorToBeneficiaryInformation != nil {
		OriginatorToBeneficiaryInformation = *r.OriginatorToBeneficiaryInformation
	}
	return
}

//
type TransactionSourceInternalSource struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency *TransactionSourceInternalSourceCurrency `json:"currency"`
	//
	Reason *TransactionSourceInternalSourceReason `json:"reason"`
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *TransactionSourceInternalSource) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
// currency.
func (r *TransactionSourceInternalSource) GetCurrency() (Currency TransactionSourceInternalSourceCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

func (r *TransactionSourceInternalSource) GetReason() (Reason TransactionSourceInternalSourceReason) {
	if r != nil && r.Reason != nil {
		Reason = *r.Reason
	}
	return
}

type TransactionSourceInternalSourceCurrency string

const (
	TransactionSourceInternalSourceCurrencyCad TransactionSourceInternalSourceCurrency = "CAD"
	TransactionSourceInternalSourceCurrencyChf TransactionSourceInternalSourceCurrency = "CHF"
	TransactionSourceInternalSourceCurrencyEur TransactionSourceInternalSourceCurrency = "EUR"
	TransactionSourceInternalSourceCurrencyGbp TransactionSourceInternalSourceCurrency = "GBP"
	TransactionSourceInternalSourceCurrencyJpy TransactionSourceInternalSourceCurrency = "JPY"
	TransactionSourceInternalSourceCurrencyUsd TransactionSourceInternalSourceCurrency = "USD"
)

type TransactionSourceInternalSourceReason string

const (
	TransactionSourceInternalSourceReasonCashback           TransactionSourceInternalSourceReason = "cashback"
	TransactionSourceInternalSourceReasonEmpyrealAdjustment TransactionSourceInternalSourceReason = "empyreal_adjustment"
	TransactionSourceInternalSourceReasonError              TransactionSourceInternalSourceReason = "error"
	TransactionSourceInternalSourceReasonErrorCorrection    TransactionSourceInternalSourceReason = "error_correction"
	TransactionSourceInternalSourceReasonFees               TransactionSourceInternalSourceReason = "fees"
	TransactionSourceInternalSourceReasonInterest           TransactionSourceInternalSourceReason = "interest"
	TransactionSourceInternalSourceReasonSampleFunds        TransactionSourceInternalSourceReason = "sample_funds"
	TransactionSourceInternalSourceReasonSampleFundsReturn  TransactionSourceInternalSourceReason = "sample_funds_return"
)

//
type TransactionSourceCardRouteRefund struct {
	// The refunded amount in the minor unit of the refunded currency. For dollars, for
	// example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the refund
	// currency.
	Currency *TransactionSourceCardRouteRefundCurrency `json:"currency"`
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

// The refunded amount in the minor unit of the refunded currency. For dollars, for
// example, this is cents.
func (r *TransactionSourceCardRouteRefund) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the refund
// currency.
func (r *TransactionSourceCardRouteRefund) GetCurrency() (Currency TransactionSourceCardRouteRefundCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

func (r *TransactionSourceCardRouteRefund) GetMerchantAcceptorID() (MerchantAcceptorID string) {
	if r != nil && r.MerchantAcceptorID != nil {
		MerchantAcceptorID = *r.MerchantAcceptorID
	}
	return
}

func (r *TransactionSourceCardRouteRefund) GetMerchantCity() (MerchantCity string) {
	if r != nil && r.MerchantCity != nil {
		MerchantCity = *r.MerchantCity
	}
	return
}

func (r *TransactionSourceCardRouteRefund) GetMerchantCountry() (MerchantCountry string) {
	if r != nil && r.MerchantCountry != nil {
		MerchantCountry = *r.MerchantCountry
	}
	return
}

func (r *TransactionSourceCardRouteRefund) GetMerchantDescriptor() (MerchantDescriptor string) {
	if r != nil && r.MerchantDescriptor != nil {
		MerchantDescriptor = *r.MerchantDescriptor
	}
	return
}

func (r *TransactionSourceCardRouteRefund) GetMerchantState() (MerchantState string) {
	if r != nil && r.MerchantState != nil {
		MerchantState = *r.MerchantState
	}
	return
}

func (r *TransactionSourceCardRouteRefund) GetMerchantCategoryCode() (MerchantCategoryCode string) {
	if r != nil && r.MerchantCategoryCode != nil {
		MerchantCategoryCode = *r.MerchantCategoryCode
	}
	return
}

type TransactionSourceCardRouteRefundCurrency string

const (
	TransactionSourceCardRouteRefundCurrencyCad TransactionSourceCardRouteRefundCurrency = "CAD"
	TransactionSourceCardRouteRefundCurrencyChf TransactionSourceCardRouteRefundCurrency = "CHF"
	TransactionSourceCardRouteRefundCurrencyEur TransactionSourceCardRouteRefundCurrency = "EUR"
	TransactionSourceCardRouteRefundCurrencyGbp TransactionSourceCardRouteRefundCurrency = "GBP"
	TransactionSourceCardRouteRefundCurrencyJpy TransactionSourceCardRouteRefundCurrency = "JPY"
	TransactionSourceCardRouteRefundCurrencyUsd TransactionSourceCardRouteRefundCurrency = "USD"
)

//
type TransactionSourceCardRouteSettlement struct {
	// The settled amount in the minor unit of the settlement currency. For dollars,
	// for example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the settlement
	// currency.
	Currency *TransactionSourceCardRouteSettlementCurrency `json:"currency"`
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

// The settled amount in the minor unit of the settlement currency. For dollars,
// for example, this is cents.
func (r *TransactionSourceCardRouteSettlement) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the settlement
// currency.
func (r *TransactionSourceCardRouteSettlement) GetCurrency() (Currency TransactionSourceCardRouteSettlementCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

func (r *TransactionSourceCardRouteSettlement) GetMerchantAcceptorID() (MerchantAcceptorID string) {
	if r != nil && r.MerchantAcceptorID != nil {
		MerchantAcceptorID = *r.MerchantAcceptorID
	}
	return
}

func (r *TransactionSourceCardRouteSettlement) GetMerchantCity() (MerchantCity string) {
	if r != nil && r.MerchantCity != nil {
		MerchantCity = *r.MerchantCity
	}
	return
}

func (r *TransactionSourceCardRouteSettlement) GetMerchantCountry() (MerchantCountry string) {
	if r != nil && r.MerchantCountry != nil {
		MerchantCountry = *r.MerchantCountry
	}
	return
}

func (r *TransactionSourceCardRouteSettlement) GetMerchantDescriptor() (MerchantDescriptor string) {
	if r != nil && r.MerchantDescriptor != nil {
		MerchantDescriptor = *r.MerchantDescriptor
	}
	return
}

func (r *TransactionSourceCardRouteSettlement) GetMerchantState() (MerchantState string) {
	if r != nil && r.MerchantState != nil {
		MerchantState = *r.MerchantState
	}
	return
}

func (r *TransactionSourceCardRouteSettlement) GetMerchantCategoryCode() (MerchantCategoryCode string) {
	if r != nil && r.MerchantCategoryCode != nil {
		MerchantCategoryCode = *r.MerchantCategoryCode
	}
	return
}

type TransactionSourceCardRouteSettlementCurrency string

const (
	TransactionSourceCardRouteSettlementCurrencyCad TransactionSourceCardRouteSettlementCurrency = "CAD"
	TransactionSourceCardRouteSettlementCurrencyChf TransactionSourceCardRouteSettlementCurrency = "CHF"
	TransactionSourceCardRouteSettlementCurrencyEur TransactionSourceCardRouteSettlementCurrency = "EUR"
	TransactionSourceCardRouteSettlementCurrencyGbp TransactionSourceCardRouteSettlementCurrency = "GBP"
	TransactionSourceCardRouteSettlementCurrencyJpy TransactionSourceCardRouteSettlementCurrency = "JPY"
	TransactionSourceCardRouteSettlementCurrencyUsd TransactionSourceCardRouteSettlementCurrency = "USD"
)

//
type TransactionSourceSampleFunds struct {
	// Where the sample funds came from.
	Originator *string `json:"originator"`
}

// Where the sample funds came from.
func (r *TransactionSourceSampleFunds) GetOriginator() (Originator string) {
	if r != nil && r.Originator != nil {
		Originator = *r.Originator
	}
	return
}

//
type TransactionSourceWireDrawdownPaymentIntention struct {
	// The transfer amount in USD cents.
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

// The transfer amount in USD cents.
func (r *TransactionSourceWireDrawdownPaymentIntention) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *TransactionSourceWireDrawdownPaymentIntention) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

func (r *TransactionSourceWireDrawdownPaymentIntention) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

func (r *TransactionSourceWireDrawdownPaymentIntention) GetMessageToRecipient() (MessageToRecipient string) {
	if r != nil && r.MessageToRecipient != nil {
		MessageToRecipient = *r.MessageToRecipient
	}
	return
}

func (r *TransactionSourceWireDrawdownPaymentIntention) GetTransferID() (TransferID string) {
	if r != nil && r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

//
type TransactionSourceWireDrawdownPaymentRejection struct {
	//
	TransferID *string `json:"transfer_id"`
}

func (r *TransactionSourceWireDrawdownPaymentRejection) GetTransferID() (TransferID string) {
	if r != nil && r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

//
type TransactionSourceWireTransferIntention struct {
	// The transfer amount in USD cents.
	Amount *int `json:"amount"`
	// The destination account number.
	AccountNumber *string `json:"account_number"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber *string `json:"routing_number"`
	// The message that will show on the recipient's bank statement.
	MessageToRecipient *string `json:"message_to_recipient"`
	//
	TransferID *string `json:"transfer_id"`
}

// The transfer amount in USD cents.
func (r *TransactionSourceWireTransferIntention) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The destination account number.
func (r *TransactionSourceWireTransferIntention) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

// The American Bankers' Association (ABA) Routing Transit Number (RTN).
func (r *TransactionSourceWireTransferIntention) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

// The message that will show on the recipient's bank statement.
func (r *TransactionSourceWireTransferIntention) GetMessageToRecipient() (MessageToRecipient string) {
	if r != nil && r.MessageToRecipient != nil {
		MessageToRecipient = *r.MessageToRecipient
	}
	return
}

func (r *TransactionSourceWireTransferIntention) GetTransferID() (TransferID string) {
	if r != nil && r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

//
type TransactionSourceWireTransferRejection struct {
	//
	TransferID *string `json:"transfer_id"`
}

func (r *TransactionSourceWireTransferRejection) GetTransferID() (TransferID string) {
	if r != nil && r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

type TransactionType string

const (
	TransactionTypeTransaction TransactionType = "transaction"
)

type ListTransactionsQuery struct {
	// Return the page of entries after this one.
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int `query:"limit"`
	// Filter Transactions for those belonging to the specified Account.
	AccountID *string `query:"account_id"`
	// Filter Transactions for those belonging to the specified route.
	RouteID   *string                         `query:"route_id"`
	CreatedAt *ListTransactionsQueryCreatedAt `query:"created_at"`
}

// Return the page of entries after this one.
func (r *ListTransactionsQuery) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *ListTransactionsQuery) GetLimit() (Limit int) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// Filter Transactions for those belonging to the specified Account.
func (r *ListTransactionsQuery) GetAccountID() (AccountID string) {
	if r != nil && r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

// Filter Transactions for those belonging to the specified route.
func (r *ListTransactionsQuery) GetRouteID() (RouteID string) {
	if r != nil && r.RouteID != nil {
		RouteID = *r.RouteID
	}
	return
}

func (r *ListTransactionsQuery) GetCreatedAt() (CreatedAt ListTransactionsQueryCreatedAt) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

type ListTransactionsQueryCreatedAt struct {
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
func (r *ListTransactionsQueryCreatedAt) GetAfter() (After string) {
	if r != nil && r.After != nil {
		After = *r.After
	}
	return
}

// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r *ListTransactionsQueryCreatedAt) GetBefore() (Before string) {
	if r != nil && r.Before != nil {
		Before = *r.Before
	}
	return
}

// Return results on or after this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r *ListTransactionsQueryCreatedAt) GetOnOrAfter() (OnOrAfter string) {
	if r != nil && r.OnOrAfter != nil {
		OnOrAfter = *r.OnOrAfter
	}
	return
}

// Return results on or before this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r *ListTransactionsQueryCreatedAt) GetOnOrBefore() (OnOrBefore string) {
	if r != nil && r.OnOrBefore != nil {
		OnOrBefore = *r.OnOrBefore
	}
	return
}

//
type TransactionList struct {
	// The contents of the list.
	Data *[]Transaction `json:"data"`
	// A pointer to a place in the list.
	NextCursor *string `json:"next_cursor"`
}

// The contents of the list.
func (r *TransactionList) GetData() (Data []Transaction) {
	if r != nil && r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r *TransactionList) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

func (r *TransactionService) Retrieve(ctx context.Context, transaction_id string, opts ...*core.RequestOpts) (res *Transaction, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/transactions/%s", transaction_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *PreloadedTransactionService) Retrieve(ctx context.Context, transaction_id string, opts ...*core.RequestOpts) (res *Transaction, err error) {
	err = r.Transactions.get(
		ctx,
		fmt.Sprintf("/transactions/%s", transaction_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

type TransactionsPage struct {
	*pagination.Page[Transaction]
}

func (r *TransactionsPage) Transaction() *Transaction {
	return r.Current()
}

func (r *TransactionsPage) GetNextPage() (*TransactionsPage, error) {
	if page, err := r.Page.GetNextPage(); err != nil {
		return nil, err
	} else {
		return &TransactionsPage{page}, nil
	}
}

func (r *TransactionService) List(ctx context.Context, query *ListTransactionsQuery, opts ...*core.RequestOpts) (res *TransactionsPage, err error) {
	page := &TransactionsPage{
		Page: &pagination.Page[Transaction]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/transactions",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

func (r *PreloadedTransactionService) List(ctx context.Context, query *ListTransactionsQuery, opts ...*core.RequestOpts) (res *TransactionsPage, err error) {
	page := &TransactionsPage{
		Page: &pagination.Page[Transaction]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/transactions",
			},
			Requester: r.Transactions.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

package simulations

import "context"
import "increase/core"

type SimulationService struct {
	Requester                  core.Requester
	get                        func(context.Context, string, *core.CoreRequest, interface{}) error
	post                       func(context.Context, string, *core.CoreRequest, interface{}) error
	patch                      func(context.Context, string, *core.CoreRequest, interface{}) error
	put                        func(context.Context, string, *core.CoreRequest, interface{}) error
	delete                     func(context.Context, string, *core.CoreRequest, interface{}) error
	AccountTransfers           *AccountTransferService
	AccountStatements          *AccountStatementService
	ACHTransfers               *ACHTransferService
	CardDisputes               *CardDisputeService
	CheckTransfers             *CheckTransferService
	DigitalWalletTokenRequests *DigitalWalletTokenRequestService
	CheckDeposits              *CheckDepositService
	WireTransfers              *WireTransferService
	Cards                      *CardService
	RealTimePaymentsTransfers  *RealTimePaymentsTransferService
}

func NewSimulationService(requester core.Requester) (r *SimulationService) {
	r = &SimulationService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	r.AccountTransfers = NewAccountTransferService(requester)
	r.AccountStatements = NewAccountStatementService(requester)
	r.ACHTransfers = NewACHTransferService(requester)
	r.CardDisputes = NewCardDisputeService(requester)
	r.CheckTransfers = NewCheckTransferService(requester)
	r.DigitalWalletTokenRequests = NewDigitalWalletTokenRequestService(requester)
	r.CheckDeposits = NewCheckDepositService(requester)
	r.WireTransfers = NewWireTransferService(requester)
	r.Cards = NewCardService(requester)
	r.RealTimePaymentsTransfers = NewRealTimePaymentsTransferService(requester)
	return
}

type PreloadedSimulationService struct {
	Simulations                *SimulationService
	AccountTransfers           *PreloadedAccountTransferService
	AccountStatements          *PreloadedAccountStatementService
	ACHTransfers               *PreloadedACHTransferService
	CardDisputes               *PreloadedCardDisputeService
	CheckTransfers             *PreloadedCheckTransferService
	DigitalWalletTokenRequests *PreloadedDigitalWalletTokenRequestService
	CheckDeposits              *PreloadedCheckDepositService
	WireTransfers              *PreloadedWireTransferService
	Cards                      *PreloadedCardService
	RealTimePaymentsTransfers  *PreloadedRealTimePaymentsTransferService
}

func (r *PreloadedSimulationService) Init(service *SimulationService) {
	r.Simulations = service
	r.AccountTransfers = &PreloadedAccountTransferService{}
	r.AccountTransfers.Init(r.Simulations.AccountTransfers)
	r.AccountStatements = &PreloadedAccountStatementService{}
	r.AccountStatements.Init(r.Simulations.AccountStatements)
	r.ACHTransfers = &PreloadedACHTransferService{}
	r.ACHTransfers.Init(r.Simulations.ACHTransfers)
	r.CardDisputes = &PreloadedCardDisputeService{}
	r.CardDisputes.Init(r.Simulations.CardDisputes)
	r.CheckTransfers = &PreloadedCheckTransferService{}
	r.CheckTransfers.Init(r.Simulations.CheckTransfers)
	r.DigitalWalletTokenRequests = &PreloadedDigitalWalletTokenRequestService{}
	r.DigitalWalletTokenRequests.Init(r.Simulations.DigitalWalletTokenRequests)
	r.CheckDeposits = &PreloadedCheckDepositService{}
	r.CheckDeposits.Init(r.Simulations.CheckDeposits)
	r.WireTransfers = &PreloadedWireTransferService{}
	r.WireTransfers.Init(r.Simulations.WireTransfers)
	r.Cards = &PreloadedCardService{}
	r.Cards.Init(r.Simulations.Cards)
	r.RealTimePaymentsTransfers = &PreloadedRealTimePaymentsTransferService{}
	r.RealTimePaymentsTransfers.Init(r.Simulations.RealTimePaymentsTransfers)
}

func NewPreloadedSimulationService(service *SimulationService) (r *PreloadedSimulationService) {
	r = &PreloadedSimulationService{}
	r.Init(service)
	return
}

type SimulateAnAccountStatementBeingCreatedParameters struct {
	// The identifier of the Account the statement is for.
	AccountId *string `json:"account_id"`
}

// The identifier of the Account the statement is for.
func (r *SimulateAnAccountStatementBeingCreatedParameters) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

//
type ACHTransferSimulation struct {
	// If the ACH Transfer attempt succeeds, this will contain the resulting
	// [Transaction](#transactions) object. The Transaction's `source` will be of
	// `category: inbound_ach_transfer`.
	Transaction *ACHTransferSimulationTransaction `json:"transaction"`
	// If the ACH Transfer attempt fails, this will contain the resulting
	// [Declined Transaction](#declined-transactions) object. The Declined
	// Transaction's `source` will be of `category: inbound_ach_transfer`.
	DeclinedTransaction *ACHTransferSimulationDeclinedTransaction `json:"declined_transaction"`
	// A constant representing the object's type. For this resource it will always be
	// `inbound_ach_transfer_simulation_result`.
	Type *ACHTransferSimulationType `json:"type"`
}

// If the ACH Transfer attempt succeeds, this will contain the resulting
// [Transaction](#transactions) object. The Transaction's `source` will be of
// `category: inbound_ach_transfer`.
func (r *ACHTransferSimulation) GetTransaction() (Transaction ACHTransferSimulationTransaction) {
	if r != nil && r.Transaction != nil {
		Transaction = *r.Transaction
	}
	return
}

// If the ACH Transfer attempt fails, this will contain the resulting
// [Declined Transaction](#declined-transactions) object. The Declined
// Transaction's `source` will be of `category: inbound_ach_transfer`.
func (r *ACHTransferSimulation) GetDeclinedTransaction() (DeclinedTransaction ACHTransferSimulationDeclinedTransaction) {
	if r != nil && r.DeclinedTransaction != nil {
		DeclinedTransaction = *r.DeclinedTransaction
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `inbound_ach_transfer_simulation_result`.
func (r *ACHTransferSimulation) GetType() (Type ACHTransferSimulationType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

//
type ACHTransferSimulationTransaction struct {
	// The identifier for the Account the Transaction belongs to.
	AccountId *string `json:"account_id"`
	// The Transaction amount in the minor unit of its currency. For dollars, for
	// example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// Transaction's currency. This will match the currency on the Transcation's
	// Account.
	Currency *ACHTransferSimulationTransactionCurrency `json:"currency"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the
	// Transaction occured.
	CreatedAt *string `json:"created_at"`
	// For a Transaction related to a transfer, this is the description you provide.
	// For a Transaction related to a payment, this is the description the vendor
	// provides.
	Description *string `json:"description"`
	// The Transaction identifier.
	Id *string `json:"id"`
	// The identifier for the route this Transaction came through. Routes are things
	// like cards and ACH details.
	RouteId *string `json:"route_id"`
	// The type of the route this Transaction came through.
	RouteType *string `json:"route_type"`
	// This is an object giving more details on the network-level event that caused the
	// Transaction. Note that for backwards compatibility reasons, additional
	// undocumented keys may appear in this object. These should be treated as
	// deprecated and will be removed in the future.
	Source *ACHTransferSimulationTransactionSource `json:"source"`
	// A constant representing the object's type. For this resource it will always be
	// `transaction`.
	Type *ACHTransferSimulationTransactionType `json:"type"`
}

// The identifier for the Account the Transaction belongs to.
func (r *ACHTransferSimulationTransaction) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

// The Transaction amount in the minor unit of its currency. For dollars, for
// example, this is cents.
func (r *ACHTransferSimulationTransaction) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// Transaction's currency. This will match the currency on the Transcation's
// Account.
func (r *ACHTransferSimulationTransaction) GetCurrency() (Currency ACHTransferSimulationTransactionCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the
// Transaction occured.
func (r *ACHTransferSimulationTransaction) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// For a Transaction related to a transfer, this is the description you provide.
// For a Transaction related to a payment, this is the description the vendor
// provides.
func (r *ACHTransferSimulationTransaction) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The Transaction identifier.
func (r *ACHTransferSimulationTransaction) GetId() (Id string) {
	if r != nil && r.Id != nil {
		Id = *r.Id
	}
	return
}

// The identifier for the route this Transaction came through. Routes are things
// like cards and ACH details.
func (r *ACHTransferSimulationTransaction) GetRouteId() (RouteId string) {
	if r != nil && r.RouteId != nil {
		RouteId = *r.RouteId
	}
	return
}

// The type of the route this Transaction came through.
func (r *ACHTransferSimulationTransaction) GetRouteType() (RouteType string) {
	if r != nil && r.RouteType != nil {
		RouteType = *r.RouteType
	}
	return
}

// This is an object giving more details on the network-level event that caused the
// Transaction. Note that for backwards compatibility reasons, additional
// undocumented keys may appear in this object. These should be treated as
// deprecated and will be removed in the future.
func (r *ACHTransferSimulationTransaction) GetSource() (Source ACHTransferSimulationTransactionSource) {
	if r != nil && r.Source != nil {
		Source = *r.Source
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `transaction`.
func (r *ACHTransferSimulationTransaction) GetType() (Type ACHTransferSimulationTransactionType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

type ACHTransferSimulationTransactionCurrency string

const (
	ACHTransferSimulationTransactionCurrencyCad ACHTransferSimulationTransactionCurrency = "CAD"
	ACHTransferSimulationTransactionCurrencyChf ACHTransferSimulationTransactionCurrency = "CHF"
	ACHTransferSimulationTransactionCurrencyEur ACHTransferSimulationTransactionCurrency = "EUR"
	ACHTransferSimulationTransactionCurrencyGbp ACHTransferSimulationTransactionCurrency = "GBP"
	ACHTransferSimulationTransactionCurrencyJpy ACHTransferSimulationTransactionCurrency = "JPY"
	ACHTransferSimulationTransactionCurrencyUsd ACHTransferSimulationTransactionCurrency = "USD"
)

//
type ACHTransferSimulationTransactionSource struct {
	// The type of transaction that took place. We may add additional possible values
	// for this enum over time; your application should be able to handle such
	// additions gracefully.
	Category *ACHTransferSimulationTransactionSourceCategory `json:"category"`
	// A Account Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `account_transfer_intention`.
	AccountTransferIntention *ACHTransferSimulationTransactionSourceAccountTransferIntention `json:"account_transfer_intention"`
	// A ACH Check Conversion Return object. This field will be present in the JSON
	// response if and only if `category` is equal to `ach_check_conversion_return`.
	ACHCheckConversionReturn *ACHTransferSimulationTransactionSourceACHCheckConversionReturn `json:"ach_check_conversion_return"`
	// A ACH Check Conversion object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_check_conversion`.
	ACHCheckConversion *ACHTransferSimulationTransactionSourceACHCheckConversion `json:"ach_check_conversion"`
	// A ACH Transfer Intention object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_transfer_intention`.
	ACHTransferIntention *ACHTransferSimulationTransactionSourceACHTransferIntention `json:"ach_transfer_intention"`
	// A ACH Transfer Rejection object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_transfer_rejection`.
	ACHTransferRejection *ACHTransferSimulationTransactionSourceACHTransferRejection `json:"ach_transfer_rejection"`
	// A ACH Transfer Return object. This field will be present in the JSON response if
	// and only if `category` is equal to `ach_transfer_return`.
	ACHTransferReturn *ACHTransferSimulationTransactionSourceACHTransferReturn `json:"ach_transfer_return"`
	// A Card Dispute Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_dispute_acceptance`.
	CardDisputeAcceptance *ACHTransferSimulationTransactionSourceCardDisputeAcceptance `json:"card_dispute_acceptance"`
	// A Card Refund object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_refund`.
	CardRefund *ACHTransferSimulationTransactionSourceCardRefund `json:"card_refund"`
	// A Card Settlement object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_settlement`.
	CardSettlement *ACHTransferSimulationTransactionSourceCardSettlement `json:"card_settlement"`
	// A Check Deposit Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_deposit_acceptance`.
	CheckDepositAcceptance *ACHTransferSimulationTransactionSourceCheckDepositAcceptance `json:"check_deposit_acceptance"`
	// A Check Deposit Return object. This field will be present in the JSON response
	// if and only if `category` is equal to `check_deposit_return`.
	CheckDepositReturn *ACHTransferSimulationTransactionSourceCheckDepositReturn `json:"check_deposit_return"`
	// A Check Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_intention`.
	CheckTransferIntention *ACHTransferSimulationTransactionSourceCheckTransferIntention `json:"check_transfer_intention"`
	// A Check Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_rejection`.
	CheckTransferRejection *ACHTransferSimulationTransactionSourceCheckTransferRejection `json:"check_transfer_rejection"`
	// A Check Transfer Stop Payment Request object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `check_transfer_stop_payment_request`.
	CheckTransferStopPaymentRequest *ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequest `json:"check_transfer_stop_payment_request"`
	// A Dispute Resolution object. This field will be present in the JSON response if
	// and only if `category` is equal to `dispute_resolution`.
	DisputeResolution *ACHTransferSimulationTransactionSourceDisputeResolution `json:"dispute_resolution"`
	// A Empyreal Cash Deposit object. This field will be present in the JSON response
	// if and only if `category` is equal to `empyreal_cash_deposit`.
	EmpyrealCashDeposit *ACHTransferSimulationTransactionSourceEmpyrealCashDeposit `json:"empyreal_cash_deposit"`
	// A Inbound ACH Transfer object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_ach_transfer`.
	InboundACHTransfer *ACHTransferSimulationTransactionSourceInboundACHTransfer `json:"inbound_ach_transfer"`
	// A Inbound Check object. This field will be present in the JSON response if and
	// only if `category` is equal to `inbound_check`.
	InboundCheck *ACHTransferSimulationTransactionSourceInboundCheck `json:"inbound_check"`
	// A Inbound International ACH Transfer object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `inbound_international_ach_transfer`.
	InboundInternationalACHTransfer *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer `json:"inbound_international_ach_transfer"`
	// A Inbound Real Time Payments Transfer Confirmation object. This field will be
	// present in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_confirmation`.
	InboundRealTimePaymentsTransferConfirmation *ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation `json:"inbound_real_time_payments_transfer_confirmation"`
	// A Inbound Wire Drawdown Payment Reversal object. This field will be present in
	// the JSON response if and only if `category` is equal to
	// `inbound_wire_drawdown_payment_reversal`.
	InboundWireDrawdownPaymentReversal *ACHTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal `json:"inbound_wire_drawdown_payment_reversal"`
	// A Inbound Wire Drawdown Payment object. This field will be present in the JSON
	// response if and only if `category` is equal to `inbound_wire_drawdown_payment`.
	InboundWireDrawdownPayment *ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment `json:"inbound_wire_drawdown_payment"`
	// A Inbound Wire Reversal object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_wire_reversal`.
	InboundWireReversal *ACHTransferSimulationTransactionSourceInboundWireReversal `json:"inbound_wire_reversal"`
	// A Inbound Wire Transfer object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_wire_transfer`.
	InboundWireTransfer *ACHTransferSimulationTransactionSourceInboundWireTransfer `json:"inbound_wire_transfer"`
	// A Internal Source object. This field will be present in the JSON response if and
	// only if `category` is equal to `internal_source`.
	InternalSource *ACHTransferSimulationTransactionSourceInternalSource `json:"internal_source"`
	// A Deprecated Card Refund object. This field will be present in the JSON response
	// if and only if `category` is equal to `card_route_refund`.
	CardRouteRefund *ACHTransferSimulationTransactionSourceCardRouteRefund `json:"card_route_refund"`
	// A Deprecated Card Settlement object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_route_settlement`.
	CardRouteSettlement *ACHTransferSimulationTransactionSourceCardRouteSettlement `json:"card_route_settlement"`
	// A Sample Funds object. This field will be present in the JSON response if and
	// only if `category` is equal to `sample_funds`.
	SampleFunds *ACHTransferSimulationTransactionSourceSampleFunds `json:"sample_funds"`
	// A Wire Drawdown Payment Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to
	// `wire_drawdown_payment_intention`.
	WireDrawdownPaymentIntention *ACHTransferSimulationTransactionSourceWireDrawdownPaymentIntention `json:"wire_drawdown_payment_intention"`
	// A Wire Drawdown Payment Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to
	// `wire_drawdown_payment_rejection`.
	WireDrawdownPaymentRejection *ACHTransferSimulationTransactionSourceWireDrawdownPaymentRejection `json:"wire_drawdown_payment_rejection"`
	// A Wire Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_intention`.
	WireTransferIntention *ACHTransferSimulationTransactionSourceWireTransferIntention `json:"wire_transfer_intention"`
	// A Wire Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_rejection`.
	WireTransferRejection *ACHTransferSimulationTransactionSourceWireTransferRejection `json:"wire_transfer_rejection"`
}

// The type of transaction that took place. We may add additional possible values
// for this enum over time; your application should be able to handle such
// additions gracefully.
func (r *ACHTransferSimulationTransactionSource) GetCategory() (Category ACHTransferSimulationTransactionSourceCategory) {
	if r != nil && r.Category != nil {
		Category = *r.Category
	}
	return
}

// A Account Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `account_transfer_intention`.
func (r *ACHTransferSimulationTransactionSource) GetAccountTransferIntention() (AccountTransferIntention ACHTransferSimulationTransactionSourceAccountTransferIntention) {
	if r != nil && r.AccountTransferIntention != nil {
		AccountTransferIntention = *r.AccountTransferIntention
	}
	return
}

// A ACH Check Conversion Return object. This field will be present in the JSON
// response if and only if `category` is equal to `ach_check_conversion_return`.
func (r *ACHTransferSimulationTransactionSource) GetACHCheckConversionReturn() (ACHCheckConversionReturn ACHTransferSimulationTransactionSourceACHCheckConversionReturn) {
	if r != nil && r.ACHCheckConversionReturn != nil {
		ACHCheckConversionReturn = *r.ACHCheckConversionReturn
	}
	return
}

// A ACH Check Conversion object. This field will be present in the JSON response
// if and only if `category` is equal to `ach_check_conversion`.
func (r *ACHTransferSimulationTransactionSource) GetACHCheckConversion() (ACHCheckConversion ACHTransferSimulationTransactionSourceACHCheckConversion) {
	if r != nil && r.ACHCheckConversion != nil {
		ACHCheckConversion = *r.ACHCheckConversion
	}
	return
}

// A ACH Transfer Intention object. This field will be present in the JSON response
// if and only if `category` is equal to `ach_transfer_intention`.
func (r *ACHTransferSimulationTransactionSource) GetACHTransferIntention() (ACHTransferIntention ACHTransferSimulationTransactionSourceACHTransferIntention) {
	if r != nil && r.ACHTransferIntention != nil {
		ACHTransferIntention = *r.ACHTransferIntention
	}
	return
}

// A ACH Transfer Rejection object. This field will be present in the JSON response
// if and only if `category` is equal to `ach_transfer_rejection`.
func (r *ACHTransferSimulationTransactionSource) GetACHTransferRejection() (ACHTransferRejection ACHTransferSimulationTransactionSourceACHTransferRejection) {
	if r != nil && r.ACHTransferRejection != nil {
		ACHTransferRejection = *r.ACHTransferRejection
	}
	return
}

// A ACH Transfer Return object. This field will be present in the JSON response if
// and only if `category` is equal to `ach_transfer_return`.
func (r *ACHTransferSimulationTransactionSource) GetACHTransferReturn() (ACHTransferReturn ACHTransferSimulationTransactionSourceACHTransferReturn) {
	if r != nil && r.ACHTransferReturn != nil {
		ACHTransferReturn = *r.ACHTransferReturn
	}
	return
}

// A Card Dispute Acceptance object. This field will be present in the JSON
// response if and only if `category` is equal to `card_dispute_acceptance`.
func (r *ACHTransferSimulationTransactionSource) GetCardDisputeAcceptance() (CardDisputeAcceptance ACHTransferSimulationTransactionSourceCardDisputeAcceptance) {
	if r != nil && r.CardDisputeAcceptance != nil {
		CardDisputeAcceptance = *r.CardDisputeAcceptance
	}
	return
}

// A Card Refund object. This field will be present in the JSON response if and
// only if `category` is equal to `card_refund`.
func (r *ACHTransferSimulationTransactionSource) GetCardRefund() (CardRefund ACHTransferSimulationTransactionSourceCardRefund) {
	if r != nil && r.CardRefund != nil {
		CardRefund = *r.CardRefund
	}
	return
}

// A Card Settlement object. This field will be present in the JSON response if and
// only if `category` is equal to `card_settlement`.
func (r *ACHTransferSimulationTransactionSource) GetCardSettlement() (CardSettlement ACHTransferSimulationTransactionSourceCardSettlement) {
	if r != nil && r.CardSettlement != nil {
		CardSettlement = *r.CardSettlement
	}
	return
}

// A Check Deposit Acceptance object. This field will be present in the JSON
// response if and only if `category` is equal to `check_deposit_acceptance`.
func (r *ACHTransferSimulationTransactionSource) GetCheckDepositAcceptance() (CheckDepositAcceptance ACHTransferSimulationTransactionSourceCheckDepositAcceptance) {
	if r != nil && r.CheckDepositAcceptance != nil {
		CheckDepositAcceptance = *r.CheckDepositAcceptance
	}
	return
}

// A Check Deposit Return object. This field will be present in the JSON response
// if and only if `category` is equal to `check_deposit_return`.
func (r *ACHTransferSimulationTransactionSource) GetCheckDepositReturn() (CheckDepositReturn ACHTransferSimulationTransactionSourceCheckDepositReturn) {
	if r != nil && r.CheckDepositReturn != nil {
		CheckDepositReturn = *r.CheckDepositReturn
	}
	return
}

// A Check Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `check_transfer_intention`.
func (r *ACHTransferSimulationTransactionSource) GetCheckTransferIntention() (CheckTransferIntention ACHTransferSimulationTransactionSourceCheckTransferIntention) {
	if r != nil && r.CheckTransferIntention != nil {
		CheckTransferIntention = *r.CheckTransferIntention
	}
	return
}

// A Check Transfer Rejection object. This field will be present in the JSON
// response if and only if `category` is equal to `check_transfer_rejection`.
func (r *ACHTransferSimulationTransactionSource) GetCheckTransferRejection() (CheckTransferRejection ACHTransferSimulationTransactionSourceCheckTransferRejection) {
	if r != nil && r.CheckTransferRejection != nil {
		CheckTransferRejection = *r.CheckTransferRejection
	}
	return
}

// A Check Transfer Stop Payment Request object. This field will be present in the
// JSON response if and only if `category` is equal to
// `check_transfer_stop_payment_request`.
func (r *ACHTransferSimulationTransactionSource) GetCheckTransferStopPaymentRequest() (CheckTransferStopPaymentRequest ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequest) {
	if r != nil && r.CheckTransferStopPaymentRequest != nil {
		CheckTransferStopPaymentRequest = *r.CheckTransferStopPaymentRequest
	}
	return
}

// A Dispute Resolution object. This field will be present in the JSON response if
// and only if `category` is equal to `dispute_resolution`.
func (r *ACHTransferSimulationTransactionSource) GetDisputeResolution() (DisputeResolution ACHTransferSimulationTransactionSourceDisputeResolution) {
	if r != nil && r.DisputeResolution != nil {
		DisputeResolution = *r.DisputeResolution
	}
	return
}

// A Empyreal Cash Deposit object. This field will be present in the JSON response
// if and only if `category` is equal to `empyreal_cash_deposit`.
func (r *ACHTransferSimulationTransactionSource) GetEmpyrealCashDeposit() (EmpyrealCashDeposit ACHTransferSimulationTransactionSourceEmpyrealCashDeposit) {
	if r != nil && r.EmpyrealCashDeposit != nil {
		EmpyrealCashDeposit = *r.EmpyrealCashDeposit
	}
	return
}

// A Inbound ACH Transfer object. This field will be present in the JSON response
// if and only if `category` is equal to `inbound_ach_transfer`.
func (r *ACHTransferSimulationTransactionSource) GetInboundACHTransfer() (InboundACHTransfer ACHTransferSimulationTransactionSourceInboundACHTransfer) {
	if r != nil && r.InboundACHTransfer != nil {
		InboundACHTransfer = *r.InboundACHTransfer
	}
	return
}

// A Inbound Check object. This field will be present in the JSON response if and
// only if `category` is equal to `inbound_check`.
func (r *ACHTransferSimulationTransactionSource) GetInboundCheck() (InboundCheck ACHTransferSimulationTransactionSourceInboundCheck) {
	if r != nil && r.InboundCheck != nil {
		InboundCheck = *r.InboundCheck
	}
	return
}

// A Inbound International ACH Transfer object. This field will be present in the
// JSON response if and only if `category` is equal to
// `inbound_international_ach_transfer`.
func (r *ACHTransferSimulationTransactionSource) GetInboundInternationalACHTransfer() (InboundInternationalACHTransfer ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) {
	if r != nil && r.InboundInternationalACHTransfer != nil {
		InboundInternationalACHTransfer = *r.InboundInternationalACHTransfer
	}
	return
}

// A Inbound Real Time Payments Transfer Confirmation object. This field will be
// present in the JSON response if and only if `category` is equal to
// `inbound_real_time_payments_transfer_confirmation`.
func (r *ACHTransferSimulationTransactionSource) GetInboundRealTimePaymentsTransferConfirmation() (InboundRealTimePaymentsTransferConfirmation ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) {
	if r != nil && r.InboundRealTimePaymentsTransferConfirmation != nil {
		InboundRealTimePaymentsTransferConfirmation = *r.InboundRealTimePaymentsTransferConfirmation
	}
	return
}

// A Inbound Wire Drawdown Payment Reversal object. This field will be present in
// the JSON response if and only if `category` is equal to
// `inbound_wire_drawdown_payment_reversal`.
func (r *ACHTransferSimulationTransactionSource) GetInboundWireDrawdownPaymentReversal() (InboundWireDrawdownPaymentReversal ACHTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) {
	if r != nil && r.InboundWireDrawdownPaymentReversal != nil {
		InboundWireDrawdownPaymentReversal = *r.InboundWireDrawdownPaymentReversal
	}
	return
}

// A Inbound Wire Drawdown Payment object. This field will be present in the JSON
// response if and only if `category` is equal to `inbound_wire_drawdown_payment`.
func (r *ACHTransferSimulationTransactionSource) GetInboundWireDrawdownPayment() (InboundWireDrawdownPayment ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment) {
	if r != nil && r.InboundWireDrawdownPayment != nil {
		InboundWireDrawdownPayment = *r.InboundWireDrawdownPayment
	}
	return
}

// A Inbound Wire Reversal object. This field will be present in the JSON response
// if and only if `category` is equal to `inbound_wire_reversal`.
func (r *ACHTransferSimulationTransactionSource) GetInboundWireReversal() (InboundWireReversal ACHTransferSimulationTransactionSourceInboundWireReversal) {
	if r != nil && r.InboundWireReversal != nil {
		InboundWireReversal = *r.InboundWireReversal
	}
	return
}

// A Inbound Wire Transfer object. This field will be present in the JSON response
// if and only if `category` is equal to `inbound_wire_transfer`.
func (r *ACHTransferSimulationTransactionSource) GetInboundWireTransfer() (InboundWireTransfer ACHTransferSimulationTransactionSourceInboundWireTransfer) {
	if r != nil && r.InboundWireTransfer != nil {
		InboundWireTransfer = *r.InboundWireTransfer
	}
	return
}

// A Internal Source object. This field will be present in the JSON response if and
// only if `category` is equal to `internal_source`.
func (r *ACHTransferSimulationTransactionSource) GetInternalSource() (InternalSource ACHTransferSimulationTransactionSourceInternalSource) {
	if r != nil && r.InternalSource != nil {
		InternalSource = *r.InternalSource
	}
	return
}

// A Deprecated Card Refund object. This field will be present in the JSON response
// if and only if `category` is equal to `card_route_refund`.
func (r *ACHTransferSimulationTransactionSource) GetCardRouteRefund() (CardRouteRefund ACHTransferSimulationTransactionSourceCardRouteRefund) {
	if r != nil && r.CardRouteRefund != nil {
		CardRouteRefund = *r.CardRouteRefund
	}
	return
}

// A Deprecated Card Settlement object. This field will be present in the JSON
// response if and only if `category` is equal to `card_route_settlement`.
func (r *ACHTransferSimulationTransactionSource) GetCardRouteSettlement() (CardRouteSettlement ACHTransferSimulationTransactionSourceCardRouteSettlement) {
	if r != nil && r.CardRouteSettlement != nil {
		CardRouteSettlement = *r.CardRouteSettlement
	}
	return
}

// A Sample Funds object. This field will be present in the JSON response if and
// only if `category` is equal to `sample_funds`.
func (r *ACHTransferSimulationTransactionSource) GetSampleFunds() (SampleFunds ACHTransferSimulationTransactionSourceSampleFunds) {
	if r != nil && r.SampleFunds != nil {
		SampleFunds = *r.SampleFunds
	}
	return
}

// A Wire Drawdown Payment Intention object. This field will be present in the JSON
// response if and only if `category` is equal to
// `wire_drawdown_payment_intention`.
func (r *ACHTransferSimulationTransactionSource) GetWireDrawdownPaymentIntention() (WireDrawdownPaymentIntention ACHTransferSimulationTransactionSourceWireDrawdownPaymentIntention) {
	if r != nil && r.WireDrawdownPaymentIntention != nil {
		WireDrawdownPaymentIntention = *r.WireDrawdownPaymentIntention
	}
	return
}

// A Wire Drawdown Payment Rejection object. This field will be present in the JSON
// response if and only if `category` is equal to
// `wire_drawdown_payment_rejection`.
func (r *ACHTransferSimulationTransactionSource) GetWireDrawdownPaymentRejection() (WireDrawdownPaymentRejection ACHTransferSimulationTransactionSourceWireDrawdownPaymentRejection) {
	if r != nil && r.WireDrawdownPaymentRejection != nil {
		WireDrawdownPaymentRejection = *r.WireDrawdownPaymentRejection
	}
	return
}

// A Wire Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `wire_transfer_intention`.
func (r *ACHTransferSimulationTransactionSource) GetWireTransferIntention() (WireTransferIntention ACHTransferSimulationTransactionSourceWireTransferIntention) {
	if r != nil && r.WireTransferIntention != nil {
		WireTransferIntention = *r.WireTransferIntention
	}
	return
}

// A Wire Transfer Rejection object. This field will be present in the JSON
// response if and only if `category` is equal to `wire_transfer_rejection`.
func (r *ACHTransferSimulationTransactionSource) GetWireTransferRejection() (WireTransferRejection ACHTransferSimulationTransactionSourceWireTransferRejection) {
	if r != nil && r.WireTransferRejection != nil {
		WireTransferRejection = *r.WireTransferRejection
	}
	return
}

type ACHTransferSimulationTransactionSourceCategory string

const (
	ACHTransferSimulationTransactionSourceCategoryAccountTransferIntention                    ACHTransferSimulationTransactionSourceCategory = "account_transfer_intention"
	ACHTransferSimulationTransactionSourceCategoryACHCheckConversionReturn                    ACHTransferSimulationTransactionSourceCategory = "ach_check_conversion_return"
	ACHTransferSimulationTransactionSourceCategoryACHCheckConversion                          ACHTransferSimulationTransactionSourceCategory = "ach_check_conversion"
	ACHTransferSimulationTransactionSourceCategoryACHTransferIntention                        ACHTransferSimulationTransactionSourceCategory = "ach_transfer_intention"
	ACHTransferSimulationTransactionSourceCategoryACHTransferRejection                        ACHTransferSimulationTransactionSourceCategory = "ach_transfer_rejection"
	ACHTransferSimulationTransactionSourceCategoryACHTransferReturn                           ACHTransferSimulationTransactionSourceCategory = "ach_transfer_return"
	ACHTransferSimulationTransactionSourceCategoryCardDisputeAcceptance                       ACHTransferSimulationTransactionSourceCategory = "card_dispute_acceptance"
	ACHTransferSimulationTransactionSourceCategoryCardRefund                                  ACHTransferSimulationTransactionSourceCategory = "card_refund"
	ACHTransferSimulationTransactionSourceCategoryCardSettlement                              ACHTransferSimulationTransactionSourceCategory = "card_settlement"
	ACHTransferSimulationTransactionSourceCategoryCheckDepositAcceptance                      ACHTransferSimulationTransactionSourceCategory = "check_deposit_acceptance"
	ACHTransferSimulationTransactionSourceCategoryCheckDepositReturn                          ACHTransferSimulationTransactionSourceCategory = "check_deposit_return"
	ACHTransferSimulationTransactionSourceCategoryCheckTransferIntention                      ACHTransferSimulationTransactionSourceCategory = "check_transfer_intention"
	ACHTransferSimulationTransactionSourceCategoryCheckTransferRejection                      ACHTransferSimulationTransactionSourceCategory = "check_transfer_rejection"
	ACHTransferSimulationTransactionSourceCategoryCheckTransferStopPaymentRequest             ACHTransferSimulationTransactionSourceCategory = "check_transfer_stop_payment_request"
	ACHTransferSimulationTransactionSourceCategoryDisputeResolution                           ACHTransferSimulationTransactionSourceCategory = "dispute_resolution"
	ACHTransferSimulationTransactionSourceCategoryEmpyrealCashDeposit                         ACHTransferSimulationTransactionSourceCategory = "empyreal_cash_deposit"
	ACHTransferSimulationTransactionSourceCategoryInboundACHTransfer                          ACHTransferSimulationTransactionSourceCategory = "inbound_ach_transfer"
	ACHTransferSimulationTransactionSourceCategoryInboundCheck                                ACHTransferSimulationTransactionSourceCategory = "inbound_check"
	ACHTransferSimulationTransactionSourceCategoryInboundInternationalACHTransfer             ACHTransferSimulationTransactionSourceCategory = "inbound_international_ach_transfer"
	ACHTransferSimulationTransactionSourceCategoryInboundRealTimePaymentsTransferConfirmation ACHTransferSimulationTransactionSourceCategory = "inbound_real_time_payments_transfer_confirmation"
	ACHTransferSimulationTransactionSourceCategoryInboundWireDrawdownPaymentReversal          ACHTransferSimulationTransactionSourceCategory = "inbound_wire_drawdown_payment_reversal"
	ACHTransferSimulationTransactionSourceCategoryInboundWireDrawdownPayment                  ACHTransferSimulationTransactionSourceCategory = "inbound_wire_drawdown_payment"
	ACHTransferSimulationTransactionSourceCategoryInboundWireReversal                         ACHTransferSimulationTransactionSourceCategory = "inbound_wire_reversal"
	ACHTransferSimulationTransactionSourceCategoryInboundWireTransfer                         ACHTransferSimulationTransactionSourceCategory = "inbound_wire_transfer"
	ACHTransferSimulationTransactionSourceCategoryInternalSource                              ACHTransferSimulationTransactionSourceCategory = "internal_source"
	ACHTransferSimulationTransactionSourceCategoryCardRouteRefund                             ACHTransferSimulationTransactionSourceCategory = "card_route_refund"
	ACHTransferSimulationTransactionSourceCategoryCardRouteSettlement                         ACHTransferSimulationTransactionSourceCategory = "card_route_settlement"
	ACHTransferSimulationTransactionSourceCategoryRealTimePaymentsTransferAcknowledgement     ACHTransferSimulationTransactionSourceCategory = "real_time_payments_transfer_acknowledgement"
	ACHTransferSimulationTransactionSourceCategorySampleFunds                                 ACHTransferSimulationTransactionSourceCategory = "sample_funds"
	ACHTransferSimulationTransactionSourceCategoryWireDrawdownPaymentIntention                ACHTransferSimulationTransactionSourceCategory = "wire_drawdown_payment_intention"
	ACHTransferSimulationTransactionSourceCategoryWireDrawdownPaymentRejection                ACHTransferSimulationTransactionSourceCategory = "wire_drawdown_payment_rejection"
	ACHTransferSimulationTransactionSourceCategoryWireTransferIntention                       ACHTransferSimulationTransactionSourceCategory = "wire_transfer_intention"
	ACHTransferSimulationTransactionSourceCategoryWireTransferRejection                       ACHTransferSimulationTransactionSourceCategory = "wire_transfer_rejection"
	ACHTransferSimulationTransactionSourceCategoryOther                                       ACHTransferSimulationTransactionSourceCategory = "other"
)

//
type ACHTransferSimulationTransactionSourceAccountTransferIntention struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency *ACHTransferSimulationTransactionSourceAccountTransferIntentionCurrency `json:"currency"`
	// The description you chose to give the transfer.
	Description *string `json:"description"`
	// The identifier of the Account to where the Account Transfer was sent.
	DestinationAccountId *string `json:"destination_account_id"`
	// The identifier of the Account from where the Account Transfer was sent.
	SourceAccountId *string `json:"source_account_id"`
	// The identifier of the Account Transfer that led to this Pending Transaction.
	TransferId *string `json:"transfer_id"`
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *ACHTransferSimulationTransactionSourceAccountTransferIntention) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
// account currency.
func (r *ACHTransferSimulationTransactionSourceAccountTransferIntention) GetCurrency() (Currency ACHTransferSimulationTransactionSourceAccountTransferIntentionCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The description you chose to give the transfer.
func (r *ACHTransferSimulationTransactionSourceAccountTransferIntention) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The identifier of the Account to where the Account Transfer was sent.
func (r *ACHTransferSimulationTransactionSourceAccountTransferIntention) GetDestinationAccountId() (DestinationAccountId string) {
	if r != nil && r.DestinationAccountId != nil {
		DestinationAccountId = *r.DestinationAccountId
	}
	return
}

// The identifier of the Account from where the Account Transfer was sent.
func (r *ACHTransferSimulationTransactionSourceAccountTransferIntention) GetSourceAccountId() (SourceAccountId string) {
	if r != nil && r.SourceAccountId != nil {
		SourceAccountId = *r.SourceAccountId
	}
	return
}

// The identifier of the Account Transfer that led to this Pending Transaction.
func (r *ACHTransferSimulationTransactionSourceAccountTransferIntention) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}

type ACHTransferSimulationTransactionSourceAccountTransferIntentionCurrency string

const (
	ACHTransferSimulationTransactionSourceAccountTransferIntentionCurrencyCad ACHTransferSimulationTransactionSourceAccountTransferIntentionCurrency = "CAD"
	ACHTransferSimulationTransactionSourceAccountTransferIntentionCurrencyChf ACHTransferSimulationTransactionSourceAccountTransferIntentionCurrency = "CHF"
	ACHTransferSimulationTransactionSourceAccountTransferIntentionCurrencyEur ACHTransferSimulationTransactionSourceAccountTransferIntentionCurrency = "EUR"
	ACHTransferSimulationTransactionSourceAccountTransferIntentionCurrencyGbp ACHTransferSimulationTransactionSourceAccountTransferIntentionCurrency = "GBP"
	ACHTransferSimulationTransactionSourceAccountTransferIntentionCurrencyJpy ACHTransferSimulationTransactionSourceAccountTransferIntentionCurrency = "JPY"
	ACHTransferSimulationTransactionSourceAccountTransferIntentionCurrencyUsd ACHTransferSimulationTransactionSourceAccountTransferIntentionCurrency = "USD"
)

//
type ACHTransferSimulationTransactionSourceACHCheckConversionReturn struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int `json:"amount"`
	// Why the transfer was returned.
	ReturnReasonCode *string `json:"return_reason_code"`
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *ACHTransferSimulationTransactionSourceACHCheckConversionReturn) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// Why the transfer was returned.
func (r *ACHTransferSimulationTransactionSourceACHCheckConversionReturn) GetReturnReasonCode() (ReturnReasonCode string) {
	if r != nil && r.ReturnReasonCode != nil {
		ReturnReasonCode = *r.ReturnReasonCode
	}
	return
}

//
type ACHTransferSimulationTransactionSourceACHCheckConversion struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int `json:"amount"`
	// The identifier of the File containing an image of the returned check.
	FileId *string `json:"file_id"`
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *ACHTransferSimulationTransactionSourceACHCheckConversion) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The identifier of the File containing an image of the returned check.
func (r *ACHTransferSimulationTransactionSourceACHCheckConversion) GetFileId() (FileId string) {
	if r != nil && r.FileId != nil {
		FileId = *r.FileId
	}
	return
}

//
type ACHTransferSimulationTransactionSourceACHTransferIntention struct {
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
	TransferId *string `json:"transfer_id"`
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *ACHTransferSimulationTransactionSourceACHTransferIntention) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceACHTransferIntention) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceACHTransferIntention) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceACHTransferIntention) GetStatementDescriptor() (StatementDescriptor string) {
	if r != nil && r.StatementDescriptor != nil {
		StatementDescriptor = *r.StatementDescriptor
	}
	return
}

// The identifier of the ACH Transfer that led to this Transaction.
func (r *ACHTransferSimulationTransactionSourceACHTransferIntention) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}

//
type ACHTransferSimulationTransactionSourceACHTransferRejection struct {
	// The identifier of the ACH Transfer that led to this Transaction.
	TransferId *string `json:"transfer_id"`
}

// The identifier of the ACH Transfer that led to this Transaction.
func (r *ACHTransferSimulationTransactionSourceACHTransferRejection) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}

//
type ACHTransferSimulationTransactionSourceACHTransferReturn struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt *string `json:"created_at"`
	// Why the ACH Transfer was returned.
	ReturnReasonCode *ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode `json:"return_reason_code"`
	// The identifier of the ACH Transfer associated with this return.
	TransferId *string `json:"transfer_id"`
	// The identifier of the Tranasaction associated with this return.
	TransactionId *string `json:"transaction_id"`
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the transfer was created.
func (r *ACHTransferSimulationTransactionSourceACHTransferReturn) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// Why the ACH Transfer was returned.
func (r *ACHTransferSimulationTransactionSourceACHTransferReturn) GetReturnReasonCode() (ReturnReasonCode ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode) {
	if r != nil && r.ReturnReasonCode != nil {
		ReturnReasonCode = *r.ReturnReasonCode
	}
	return
}

// The identifier of the ACH Transfer associated with this return.
func (r *ACHTransferSimulationTransactionSourceACHTransferReturn) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}

// The identifier of the Tranasaction associated with this return.
func (r *ACHTransferSimulationTransactionSourceACHTransferReturn) GetTransactionId() (TransactionId string) {
	if r != nil && r.TransactionId != nil {
		TransactionId = *r.TransactionId
	}
	return
}

type ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode string

const (
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeInsufficientFund                                          ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "insufficient_fund"
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeNoAccount                                                 ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "no_account"
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeAccountClosed                                             ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "account_closed"
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeInvalidAccountNumberStructure                             ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "invalid_account_number_structure"
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeAccountFrozenEntryReturnedPerOfacInstruction              ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "account_frozen_entry_returned_per_ofac_instruction"
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeCreditEntryRefusedByReceiver                              ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "credit_entry_refused_by_receiver"
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeUnauthorizedDebitToConsumerAccountUsingCorporateSecCode   ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "unauthorized_debit_to_consumer_account_using_corporate_sec_code"
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeCorporateCustomerAdvisedNotAuthorized                     ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "corporate_customer_advised_not_authorized"
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodePaymentStopped                                            ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "payment_stopped"
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeNonTransactionAccount                                     ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "non_transaction_account"
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeUncollectedFunds                                          ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "uncollected_funds"
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeRoutingNumberCheckDigitError                              ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "routing_number_check_digit_error"
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "customer_advised_unauthorized_improper_ineligible_or_incomplete"
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeAmountFieldError                                          ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "amount_field_error"
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeAuthorizationRevokedByCustomer                            ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "authorization_revoked_by_customer"
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeInvalidACHRoutingNumber                                   ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "invalid_ach_routing_number"
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeFileRecordEditCriteria                                    ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "file_record_edit_criteria"
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeEnrInvalidIndividualName                                  ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "enr_invalid_individual_name"
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeReturnedPerOdfiRequest                                    ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "returned_per_odfi_request"
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeAddendaError                                              ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "addenda_error"
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeLimitedParticipationDfi                                   ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "limited_participation_dfi"
	ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeOther                                                     ACHTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "other"
)

//
type ACHTransferSimulationTransactionSourceCardDisputeAcceptance struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card Dispute was accepted.
	AcceptedAt *string `json:"accepted_at"`
	// The identifier of the Card Dispute that was accepted.
	CardDisputeId *string `json:"card_dispute_id"`
	// The identifier of the Transaction that was created to return the disputed funds
	// to your account.
	TransactionId *string `json:"transaction_id"`
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the Card Dispute was accepted.
func (r *ACHTransferSimulationTransactionSourceCardDisputeAcceptance) GetAcceptedAt() (AcceptedAt string) {
	if r != nil && r.AcceptedAt != nil {
		AcceptedAt = *r.AcceptedAt
	}
	return
}

// The identifier of the Card Dispute that was accepted.
func (r *ACHTransferSimulationTransactionSourceCardDisputeAcceptance) GetCardDisputeId() (CardDisputeId string) {
	if r != nil && r.CardDisputeId != nil {
		CardDisputeId = *r.CardDisputeId
	}
	return
}

// The identifier of the Transaction that was created to return the disputed funds
// to your account.
func (r *ACHTransferSimulationTransactionSourceCardDisputeAcceptance) GetTransactionId() (TransactionId string) {
	if r != nil && r.TransactionId != nil {
		TransactionId = *r.TransactionId
	}
	return
}

//
type ACHTransferSimulationTransactionSourceCardRefund struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *ACHTransferSimulationTransactionSourceCardRefundCurrency `json:"currency"`
	// The identifier for the Transaction this refunds, if any.
	CardSettlementTransactionId *string `json:"card_settlement_transaction_id"`
	// A constant representing the object's type. For this resource it will always be
	// `card_refund`.
	Type *ACHTransferSimulationTransactionSourceCardRefundType `json:"type"`
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *ACHTransferSimulationTransactionSourceCardRefund) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *ACHTransferSimulationTransactionSourceCardRefund) GetCurrency() (Currency ACHTransferSimulationTransactionSourceCardRefundCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The identifier for the Transaction this refunds, if any.
func (r *ACHTransferSimulationTransactionSourceCardRefund) GetCardSettlementTransactionId() (CardSettlementTransactionId string) {
	if r != nil && r.CardSettlementTransactionId != nil {
		CardSettlementTransactionId = *r.CardSettlementTransactionId
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `card_refund`.
func (r *ACHTransferSimulationTransactionSourceCardRefund) GetType() (Type ACHTransferSimulationTransactionSourceCardRefundType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

type ACHTransferSimulationTransactionSourceCardRefundCurrency string

const (
	ACHTransferSimulationTransactionSourceCardRefundCurrencyCad ACHTransferSimulationTransactionSourceCardRefundCurrency = "CAD"
	ACHTransferSimulationTransactionSourceCardRefundCurrencyChf ACHTransferSimulationTransactionSourceCardRefundCurrency = "CHF"
	ACHTransferSimulationTransactionSourceCardRefundCurrencyEur ACHTransferSimulationTransactionSourceCardRefundCurrency = "EUR"
	ACHTransferSimulationTransactionSourceCardRefundCurrencyGbp ACHTransferSimulationTransactionSourceCardRefundCurrency = "GBP"
	ACHTransferSimulationTransactionSourceCardRefundCurrencyJpy ACHTransferSimulationTransactionSourceCardRefundCurrency = "JPY"
	ACHTransferSimulationTransactionSourceCardRefundCurrencyUsd ACHTransferSimulationTransactionSourceCardRefundCurrency = "USD"
)

type ACHTransferSimulationTransactionSourceCardRefundType string

const (
	ACHTransferSimulationTransactionSourceCardRefundTypeCardRefund ACHTransferSimulationTransactionSourceCardRefundType = "card_refund"
)

//
type ACHTransferSimulationTransactionSourceCardSettlement struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *ACHTransferSimulationTransactionSourceCardSettlementCurrency `json:"currency"`
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
	PendingTransactionId *string `json:"pending_transaction_id"`
	// A constant representing the object's type. For this resource it will always be
	// `card_settlement`.
	Type *ACHTransferSimulationTransactionSourceCardSettlementType `json:"type"`
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *ACHTransferSimulationTransactionSourceCardSettlement) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *ACHTransferSimulationTransactionSourceCardSettlement) GetCurrency() (Currency ACHTransferSimulationTransactionSourceCardSettlementCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceCardSettlement) GetMerchantCity() (MerchantCity string) {
	if r != nil && r.MerchantCity != nil {
		MerchantCity = *r.MerchantCity
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceCardSettlement) GetMerchantCountry() (MerchantCountry string) {
	if r != nil && r.MerchantCountry != nil {
		MerchantCountry = *r.MerchantCountry
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceCardSettlement) GetMerchantName() (MerchantName string) {
	if r != nil && r.MerchantName != nil {
		MerchantName = *r.MerchantName
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceCardSettlement) GetMerchantCategoryCode() (MerchantCategoryCode string) {
	if r != nil && r.MerchantCategoryCode != nil {
		MerchantCategoryCode = *r.MerchantCategoryCode
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceCardSettlement) GetMerchantState() (MerchantState string) {
	if r != nil && r.MerchantState != nil {
		MerchantState = *r.MerchantState
	}
	return
}

// The identifier of the Pending Transaction associated with this Transaction.
func (r *ACHTransferSimulationTransactionSourceCardSettlement) GetPendingTransactionId() (PendingTransactionId string) {
	if r != nil && r.PendingTransactionId != nil {
		PendingTransactionId = *r.PendingTransactionId
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `card_settlement`.
func (r *ACHTransferSimulationTransactionSourceCardSettlement) GetType() (Type ACHTransferSimulationTransactionSourceCardSettlementType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

type ACHTransferSimulationTransactionSourceCardSettlementCurrency string

const (
	ACHTransferSimulationTransactionSourceCardSettlementCurrencyCad ACHTransferSimulationTransactionSourceCardSettlementCurrency = "CAD"
	ACHTransferSimulationTransactionSourceCardSettlementCurrencyChf ACHTransferSimulationTransactionSourceCardSettlementCurrency = "CHF"
	ACHTransferSimulationTransactionSourceCardSettlementCurrencyEur ACHTransferSimulationTransactionSourceCardSettlementCurrency = "EUR"
	ACHTransferSimulationTransactionSourceCardSettlementCurrencyGbp ACHTransferSimulationTransactionSourceCardSettlementCurrency = "GBP"
	ACHTransferSimulationTransactionSourceCardSettlementCurrencyJpy ACHTransferSimulationTransactionSourceCardSettlementCurrency = "JPY"
	ACHTransferSimulationTransactionSourceCardSettlementCurrencyUsd ACHTransferSimulationTransactionSourceCardSettlementCurrency = "USD"
)

type ACHTransferSimulationTransactionSourceCardSettlementType string

const (
	ACHTransferSimulationTransactionSourceCardSettlementTypeCardSettlement ACHTransferSimulationTransactionSourceCardSettlementType = "card_settlement"
)

//
type ACHTransferSimulationTransactionSourceCheckDepositAcceptance struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *ACHTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency `json:"currency"`
	// The ID of the Check Deposit that led to the Transaction.
	CheckDepositId *string `json:"check_deposit_id"`
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *ACHTransferSimulationTransactionSourceCheckDepositAcceptance) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *ACHTransferSimulationTransactionSourceCheckDepositAcceptance) GetCurrency() (Currency ACHTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The ID of the Check Deposit that led to the Transaction.
func (r *ACHTransferSimulationTransactionSourceCheckDepositAcceptance) GetCheckDepositId() (CheckDepositId string) {
	if r != nil && r.CheckDepositId != nil {
		CheckDepositId = *r.CheckDepositId
	}
	return
}

type ACHTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency string

const (
	ACHTransferSimulationTransactionSourceCheckDepositAcceptanceCurrencyCad ACHTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency = "CAD"
	ACHTransferSimulationTransactionSourceCheckDepositAcceptanceCurrencyChf ACHTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency = "CHF"
	ACHTransferSimulationTransactionSourceCheckDepositAcceptanceCurrencyEur ACHTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency = "EUR"
	ACHTransferSimulationTransactionSourceCheckDepositAcceptanceCurrencyGbp ACHTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency = "GBP"
	ACHTransferSimulationTransactionSourceCheckDepositAcceptanceCurrencyJpy ACHTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency = "JPY"
	ACHTransferSimulationTransactionSourceCheckDepositAcceptanceCurrencyUsd ACHTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency = "USD"
)

//
type ACHTransferSimulationTransactionSourceCheckDepositReturn struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check deposit was returned.
	ReturnedAt *string `json:"returned_at"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *ACHTransferSimulationTransactionSourceCheckDepositReturnCurrency `json:"currency"`
	// The identifier of the Check Deposit that was returned.
	CheckDepositId *string `json:"check_deposit_id"`
	// The identifier of the transaction that reversed the original check deposit
	// transaction.
	TransactionId *string `json:"transaction_id"`
	//
	ReturnReason *ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReason `json:"return_reason"`
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *ACHTransferSimulationTransactionSourceCheckDepositReturn) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the check deposit was returned.
func (r *ACHTransferSimulationTransactionSourceCheckDepositReturn) GetReturnedAt() (ReturnedAt string) {
	if r != nil && r.ReturnedAt != nil {
		ReturnedAt = *r.ReturnedAt
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *ACHTransferSimulationTransactionSourceCheckDepositReturn) GetCurrency() (Currency ACHTransferSimulationTransactionSourceCheckDepositReturnCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The identifier of the Check Deposit that was returned.
func (r *ACHTransferSimulationTransactionSourceCheckDepositReturn) GetCheckDepositId() (CheckDepositId string) {
	if r != nil && r.CheckDepositId != nil {
		CheckDepositId = *r.CheckDepositId
	}
	return
}

// The identifier of the transaction that reversed the original check deposit
// transaction.
func (r *ACHTransferSimulationTransactionSourceCheckDepositReturn) GetTransactionId() (TransactionId string) {
	if r != nil && r.TransactionId != nil {
		TransactionId = *r.TransactionId
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceCheckDepositReturn) GetReturnReason() (ReturnReason ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReason) {
	if r != nil && r.ReturnReason != nil {
		ReturnReason = *r.ReturnReason
	}
	return
}

type ACHTransferSimulationTransactionSourceCheckDepositReturnCurrency string

const (
	ACHTransferSimulationTransactionSourceCheckDepositReturnCurrencyCad ACHTransferSimulationTransactionSourceCheckDepositReturnCurrency = "CAD"
	ACHTransferSimulationTransactionSourceCheckDepositReturnCurrencyChf ACHTransferSimulationTransactionSourceCheckDepositReturnCurrency = "CHF"
	ACHTransferSimulationTransactionSourceCheckDepositReturnCurrencyEur ACHTransferSimulationTransactionSourceCheckDepositReturnCurrency = "EUR"
	ACHTransferSimulationTransactionSourceCheckDepositReturnCurrencyGbp ACHTransferSimulationTransactionSourceCheckDepositReturnCurrency = "GBP"
	ACHTransferSimulationTransactionSourceCheckDepositReturnCurrencyJpy ACHTransferSimulationTransactionSourceCheckDepositReturnCurrency = "JPY"
	ACHTransferSimulationTransactionSourceCheckDepositReturnCurrencyUsd ACHTransferSimulationTransactionSourceCheckDepositReturnCurrency = "USD"
)

type ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReason string

const (
	ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReasonACHConversionNotSupported ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "ach_conversion_not_supported"
	ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReasonDuplicateSubmission       ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "duplicate_submission"
	ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReasonInsufficientFunds         ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "insufficient_funds"
	ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReasonNoAccount                 ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "no_account"
	ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReasonNotAuthorized             ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "not_authorized"
	ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReasonStaleDated                ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "stale_dated"
	ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReasonStopPayment               ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "stop_payment"
	ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReasonUnknownReason             ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "unknown_reason"
	ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReasonUnmatchedDetails          ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "unmatched_details"
	ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReasonUnreadableImage           ACHTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "unreadable_image"
)

//
type ACHTransferSimulationTransactionSourceCheckTransferIntention struct {
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
	Currency *ACHTransferSimulationTransactionSourceCheckTransferIntentionCurrency `json:"currency"`
	// The name that will be printed on the check.
	RecipientName *string `json:"recipient_name"`
	// The identifier of the Check Transfer with which this is associated.
	TransferId *string `json:"transfer_id"`
}

// The street address of the check's destination.
func (r *ACHTransferSimulationTransactionSourceCheckTransferIntention) GetAddressLine1() (AddressLine1 string) {
	if r != nil && r.AddressLine1 != nil {
		AddressLine1 = *r.AddressLine1
	}
	return
}

// The second line of the address of the check's destination.
func (r *ACHTransferSimulationTransactionSourceCheckTransferIntention) GetAddressLine2() (AddressLine2 string) {
	if r != nil && r.AddressLine2 != nil {
		AddressLine2 = *r.AddressLine2
	}
	return
}

// The city of the check's destination.
func (r *ACHTransferSimulationTransactionSourceCheckTransferIntention) GetAddressCity() (AddressCity string) {
	if r != nil && r.AddressCity != nil {
		AddressCity = *r.AddressCity
	}
	return
}

// The state of the check's destination.
func (r *ACHTransferSimulationTransactionSourceCheckTransferIntention) GetAddressState() (AddressState string) {
	if r != nil && r.AddressState != nil {
		AddressState = *r.AddressState
	}
	return
}

// The postal code of the check's destination.
func (r *ACHTransferSimulationTransactionSourceCheckTransferIntention) GetAddressZip() (AddressZip string) {
	if r != nil && r.AddressZip != nil {
		AddressZip = *r.AddressZip
	}
	return
}

// The transfer amount in USD cents.
func (r *ACHTransferSimulationTransactionSourceCheckTransferIntention) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
// currency.
func (r *ACHTransferSimulationTransactionSourceCheckTransferIntention) GetCurrency() (Currency ACHTransferSimulationTransactionSourceCheckTransferIntentionCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The name that will be printed on the check.
func (r *ACHTransferSimulationTransactionSourceCheckTransferIntention) GetRecipientName() (RecipientName string) {
	if r != nil && r.RecipientName != nil {
		RecipientName = *r.RecipientName
	}
	return
}

// The identifier of the Check Transfer with which this is associated.
func (r *ACHTransferSimulationTransactionSourceCheckTransferIntention) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}

type ACHTransferSimulationTransactionSourceCheckTransferIntentionCurrency string

const (
	ACHTransferSimulationTransactionSourceCheckTransferIntentionCurrencyCad ACHTransferSimulationTransactionSourceCheckTransferIntentionCurrency = "CAD"
	ACHTransferSimulationTransactionSourceCheckTransferIntentionCurrencyChf ACHTransferSimulationTransactionSourceCheckTransferIntentionCurrency = "CHF"
	ACHTransferSimulationTransactionSourceCheckTransferIntentionCurrencyEur ACHTransferSimulationTransactionSourceCheckTransferIntentionCurrency = "EUR"
	ACHTransferSimulationTransactionSourceCheckTransferIntentionCurrencyGbp ACHTransferSimulationTransactionSourceCheckTransferIntentionCurrency = "GBP"
	ACHTransferSimulationTransactionSourceCheckTransferIntentionCurrencyJpy ACHTransferSimulationTransactionSourceCheckTransferIntentionCurrency = "JPY"
	ACHTransferSimulationTransactionSourceCheckTransferIntentionCurrencyUsd ACHTransferSimulationTransactionSourceCheckTransferIntentionCurrency = "USD"
)

//
type ACHTransferSimulationTransactionSourceCheckTransferRejection struct {
	// The identifier of the Check Transfer that led to this Transaction.
	TransferId *string `json:"transfer_id"`
}

// The identifier of the Check Transfer that led to this Transaction.
func (r *ACHTransferSimulationTransactionSourceCheckTransferRejection) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}

//
type ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequest struct {
	// The ID of the check transfer that was stopped.
	TransferId *string `json:"transfer_id"`
	// The transaction ID of the corresponding credit transaction.
	TransactionId *string `json:"transaction_id"`
	// The time the stop-payment was requested.
	RequestedAt *string `json:"requested_at"`
	// A constant representing the object's type. For this resource it will always be
	// `check_transfer_stop_payment_request`.
	Type *ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequestType `json:"type"`
}

// The ID of the check transfer that was stopped.
func (r *ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequest) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}

// The transaction ID of the corresponding credit transaction.
func (r *ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequest) GetTransactionId() (TransactionId string) {
	if r != nil && r.TransactionId != nil {
		TransactionId = *r.TransactionId
	}
	return
}

// The time the stop-payment was requested.
func (r *ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequest) GetRequestedAt() (RequestedAt string) {
	if r != nil && r.RequestedAt != nil {
		RequestedAt = *r.RequestedAt
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `check_transfer_stop_payment_request`.
func (r *ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequest) GetType() (Type ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequestType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

type ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequestType string

const (
	ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequestTypeCheckTransferStopPaymentRequest ACHTransferSimulationTransactionSourceCheckTransferStopPaymentRequestType = "check_transfer_stop_payment_request"
)

//
type ACHTransferSimulationTransactionSourceDisputeResolution struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *ACHTransferSimulationTransactionSourceDisputeResolutionCurrency `json:"currency"`
	// The identifier of the Transaction that was disputed.
	DisputedTransactionId *string `json:"disputed_transaction_id"`
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *ACHTransferSimulationTransactionSourceDisputeResolution) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *ACHTransferSimulationTransactionSourceDisputeResolution) GetCurrency() (Currency ACHTransferSimulationTransactionSourceDisputeResolutionCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The identifier of the Transaction that was disputed.
func (r *ACHTransferSimulationTransactionSourceDisputeResolution) GetDisputedTransactionId() (DisputedTransactionId string) {
	if r != nil && r.DisputedTransactionId != nil {
		DisputedTransactionId = *r.DisputedTransactionId
	}
	return
}

type ACHTransferSimulationTransactionSourceDisputeResolutionCurrency string

const (
	ACHTransferSimulationTransactionSourceDisputeResolutionCurrencyCad ACHTransferSimulationTransactionSourceDisputeResolutionCurrency = "CAD"
	ACHTransferSimulationTransactionSourceDisputeResolutionCurrencyChf ACHTransferSimulationTransactionSourceDisputeResolutionCurrency = "CHF"
	ACHTransferSimulationTransactionSourceDisputeResolutionCurrencyEur ACHTransferSimulationTransactionSourceDisputeResolutionCurrency = "EUR"
	ACHTransferSimulationTransactionSourceDisputeResolutionCurrencyGbp ACHTransferSimulationTransactionSourceDisputeResolutionCurrency = "GBP"
	ACHTransferSimulationTransactionSourceDisputeResolutionCurrencyJpy ACHTransferSimulationTransactionSourceDisputeResolutionCurrency = "JPY"
	ACHTransferSimulationTransactionSourceDisputeResolutionCurrencyUsd ACHTransferSimulationTransactionSourceDisputeResolutionCurrency = "USD"
)

//
type ACHTransferSimulationTransactionSourceEmpyrealCashDeposit struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int `json:"amount"`
	//
	BagId *string `json:"bag_id"`
	//
	DepositDate *string `json:"deposit_date"`
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *ACHTransferSimulationTransactionSourceEmpyrealCashDeposit) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceEmpyrealCashDeposit) GetBagId() (BagId string) {
	if r != nil && r.BagId != nil {
		BagId = *r.BagId
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceEmpyrealCashDeposit) GetDepositDate() (DepositDate string) {
	if r != nil && r.DepositDate != nil {
		DepositDate = *r.DepositDate
	}
	return
}

//
type ACHTransferSimulationTransactionSourceInboundACHTransfer struct {
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
	OriginatorCompanyId *string `json:"originator_company_id"`
	//
	ReceiverIdNumber *string `json:"receiver_id_number"`
	//
	ReceiverName *string `json:"receiver_name"`
	//
	TraceNumber *string `json:"trace_number"`
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *ACHTransferSimulationTransactionSourceInboundACHTransfer) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundACHTransfer) GetOriginatorCompanyName() (OriginatorCompanyName string) {
	if r != nil && r.OriginatorCompanyName != nil {
		OriginatorCompanyName = *r.OriginatorCompanyName
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundACHTransfer) GetOriginatorCompanyDescriptiveDate() (OriginatorCompanyDescriptiveDate string) {
	if r != nil && r.OriginatorCompanyDescriptiveDate != nil {
		OriginatorCompanyDescriptiveDate = *r.OriginatorCompanyDescriptiveDate
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundACHTransfer) GetOriginatorCompanyDiscretionaryData() (OriginatorCompanyDiscretionaryData string) {
	if r != nil && r.OriginatorCompanyDiscretionaryData != nil {
		OriginatorCompanyDiscretionaryData = *r.OriginatorCompanyDiscretionaryData
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundACHTransfer) GetOriginatorCompanyEntryDescription() (OriginatorCompanyEntryDescription string) {
	if r != nil && r.OriginatorCompanyEntryDescription != nil {
		OriginatorCompanyEntryDescription = *r.OriginatorCompanyEntryDescription
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundACHTransfer) GetOriginatorCompanyId() (OriginatorCompanyId string) {
	if r != nil && r.OriginatorCompanyId != nil {
		OriginatorCompanyId = *r.OriginatorCompanyId
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundACHTransfer) GetReceiverIdNumber() (ReceiverIdNumber string) {
	if r != nil && r.ReceiverIdNumber != nil {
		ReceiverIdNumber = *r.ReceiverIdNumber
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundACHTransfer) GetReceiverName() (ReceiverName string) {
	if r != nil && r.ReceiverName != nil {
		ReceiverName = *r.ReceiverName
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundACHTransfer) GetTraceNumber() (TraceNumber string) {
	if r != nil && r.TraceNumber != nil {
		TraceNumber = *r.TraceNumber
	}
	return
}

//
type ACHTransferSimulationTransactionSourceInboundCheck struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *ACHTransferSimulationTransactionSourceInboundCheckCurrency `json:"currency"`
	//
	CheckNumber *string `json:"check_number"`
	//
	CheckFrontImageFileId *string `json:"check_front_image_file_id"`
	//
	CheckRearImageFileId *string `json:"check_rear_image_file_id"`
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *ACHTransferSimulationTransactionSourceInboundCheck) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *ACHTransferSimulationTransactionSourceInboundCheck) GetCurrency() (Currency ACHTransferSimulationTransactionSourceInboundCheckCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundCheck) GetCheckNumber() (CheckNumber string) {
	if r != nil && r.CheckNumber != nil {
		CheckNumber = *r.CheckNumber
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundCheck) GetCheckFrontImageFileId() (CheckFrontImageFileId string) {
	if r != nil && r.CheckFrontImageFileId != nil {
		CheckFrontImageFileId = *r.CheckFrontImageFileId
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundCheck) GetCheckRearImageFileId() (CheckRearImageFileId string) {
	if r != nil && r.CheckRearImageFileId != nil {
		CheckRearImageFileId = *r.CheckRearImageFileId
	}
	return
}

type ACHTransferSimulationTransactionSourceInboundCheckCurrency string

const (
	ACHTransferSimulationTransactionSourceInboundCheckCurrencyCad ACHTransferSimulationTransactionSourceInboundCheckCurrency = "CAD"
	ACHTransferSimulationTransactionSourceInboundCheckCurrencyChf ACHTransferSimulationTransactionSourceInboundCheckCurrency = "CHF"
	ACHTransferSimulationTransactionSourceInboundCheckCurrencyEur ACHTransferSimulationTransactionSourceInboundCheckCurrency = "EUR"
	ACHTransferSimulationTransactionSourceInboundCheckCurrencyGbp ACHTransferSimulationTransactionSourceInboundCheckCurrency = "GBP"
	ACHTransferSimulationTransactionSourceInboundCheckCurrencyJpy ACHTransferSimulationTransactionSourceInboundCheckCurrency = "JPY"
	ACHTransferSimulationTransactionSourceInboundCheckCurrencyUsd ACHTransferSimulationTransactionSourceInboundCheckCurrency = "USD"
)

//
type ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer struct {
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
	OriginatingDepositoryFinancialInstitutionIdQualifier *string `json:"originating_depository_financial_institution_id_qualifier"`
	//
	OriginatingDepositoryFinancialInstitutionId *string `json:"originating_depository_financial_institution_id"`
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
	ReceivingDepositoryFinancialInstitutionIdQualifier *string `json:"receiving_depository_financial_institution_id_qualifier"`
	//
	ReceivingDepositoryFinancialInstitutionId *string `json:"receiving_depository_financial_institution_id"`
	//
	ReceivingDepositoryFinancialInstitutionCountry *string `json:"receiving_depository_financial_institution_country"`
	//
	TraceNumber *string `json:"trace_number"`
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetForeignExchangeIndicator() (ForeignExchangeIndicator string) {
	if r != nil && r.ForeignExchangeIndicator != nil {
		ForeignExchangeIndicator = *r.ForeignExchangeIndicator
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetForeignExchangeReferenceIndicator() (ForeignExchangeReferenceIndicator string) {
	if r != nil && r.ForeignExchangeReferenceIndicator != nil {
		ForeignExchangeReferenceIndicator = *r.ForeignExchangeReferenceIndicator
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetForeignExchangeReference() (ForeignExchangeReference string) {
	if r != nil && r.ForeignExchangeReference != nil {
		ForeignExchangeReference = *r.ForeignExchangeReference
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetDestinationCountryCode() (DestinationCountryCode string) {
	if r != nil && r.DestinationCountryCode != nil {
		DestinationCountryCode = *r.DestinationCountryCode
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetDestinationCurrencyCode() (DestinationCurrencyCode string) {
	if r != nil && r.DestinationCurrencyCode != nil {
		DestinationCurrencyCode = *r.DestinationCurrencyCode
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetForeignPaymentAmount() (ForeignPaymentAmount int) {
	if r != nil && r.ForeignPaymentAmount != nil {
		ForeignPaymentAmount = *r.ForeignPaymentAmount
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetForeignTraceNumber() (ForeignTraceNumber string) {
	if r != nil && r.ForeignTraceNumber != nil {
		ForeignTraceNumber = *r.ForeignTraceNumber
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetInternationalTransactionTypeCode() (InternationalTransactionTypeCode string) {
	if r != nil && r.InternationalTransactionTypeCode != nil {
		InternationalTransactionTypeCode = *r.InternationalTransactionTypeCode
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatingCurrencyCode() (OriginatingCurrencyCode string) {
	if r != nil && r.OriginatingCurrencyCode != nil {
		OriginatingCurrencyCode = *r.OriginatingCurrencyCode
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatingDepositoryFinancialInstitutionName() (OriginatingDepositoryFinancialInstitutionName string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionName != nil {
		OriginatingDepositoryFinancialInstitutionName = *r.OriginatingDepositoryFinancialInstitutionName
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatingDepositoryFinancialInstitutionIdQualifier() (OriginatingDepositoryFinancialInstitutionIdQualifier string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionIdQualifier != nil {
		OriginatingDepositoryFinancialInstitutionIdQualifier = *r.OriginatingDepositoryFinancialInstitutionIdQualifier
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatingDepositoryFinancialInstitutionId() (OriginatingDepositoryFinancialInstitutionId string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionId != nil {
		OriginatingDepositoryFinancialInstitutionId = *r.OriginatingDepositoryFinancialInstitutionId
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatingDepositoryFinancialInstitutionBranchCountry() (OriginatingDepositoryFinancialInstitutionBranchCountry string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionBranchCountry != nil {
		OriginatingDepositoryFinancialInstitutionBranchCountry = *r.OriginatingDepositoryFinancialInstitutionBranchCountry
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatorCity() (OriginatorCity string) {
	if r != nil && r.OriginatorCity != nil {
		OriginatorCity = *r.OriginatorCity
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatorCompanyEntryDescription() (OriginatorCompanyEntryDescription string) {
	if r != nil && r.OriginatorCompanyEntryDescription != nil {
		OriginatorCompanyEntryDescription = *r.OriginatorCompanyEntryDescription
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatorCountry() (OriginatorCountry string) {
	if r != nil && r.OriginatorCountry != nil {
		OriginatorCountry = *r.OriginatorCountry
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatorIdentification() (OriginatorIdentification string) {
	if r != nil && r.OriginatorIdentification != nil {
		OriginatorIdentification = *r.OriginatorIdentification
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatorName() (OriginatorName string) {
	if r != nil && r.OriginatorName != nil {
		OriginatorName = *r.OriginatorName
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatorPostalCode() (OriginatorPostalCode string) {
	if r != nil && r.OriginatorPostalCode != nil {
		OriginatorPostalCode = *r.OriginatorPostalCode
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatorStreetAddress() (OriginatorStreetAddress string) {
	if r != nil && r.OriginatorStreetAddress != nil {
		OriginatorStreetAddress = *r.OriginatorStreetAddress
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatorStateOrProvince() (OriginatorStateOrProvince string) {
	if r != nil && r.OriginatorStateOrProvince != nil {
		OriginatorStateOrProvince = *r.OriginatorStateOrProvince
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetPaymentRelatedInformation() (PaymentRelatedInformation string) {
	if r != nil && r.PaymentRelatedInformation != nil {
		PaymentRelatedInformation = *r.PaymentRelatedInformation
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetPaymentRelatedInformation2() (PaymentRelatedInformation2 string) {
	if r != nil && r.PaymentRelatedInformation2 != nil {
		PaymentRelatedInformation2 = *r.PaymentRelatedInformation2
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetReceiverIdentificationNumber() (ReceiverIdentificationNumber string) {
	if r != nil && r.ReceiverIdentificationNumber != nil {
		ReceiverIdentificationNumber = *r.ReceiverIdentificationNumber
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetReceiverStreetAddress() (ReceiverStreetAddress string) {
	if r != nil && r.ReceiverStreetAddress != nil {
		ReceiverStreetAddress = *r.ReceiverStreetAddress
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetReceiverCity() (ReceiverCity string) {
	if r != nil && r.ReceiverCity != nil {
		ReceiverCity = *r.ReceiverCity
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetReceiverStateOrProvince() (ReceiverStateOrProvince string) {
	if r != nil && r.ReceiverStateOrProvince != nil {
		ReceiverStateOrProvince = *r.ReceiverStateOrProvince
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetReceiverCountry() (ReceiverCountry string) {
	if r != nil && r.ReceiverCountry != nil {
		ReceiverCountry = *r.ReceiverCountry
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetReceiverPostalCode() (ReceiverPostalCode string) {
	if r != nil && r.ReceiverPostalCode != nil {
		ReceiverPostalCode = *r.ReceiverPostalCode
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetReceivingCompanyOrIndividualName() (ReceivingCompanyOrIndividualName string) {
	if r != nil && r.ReceivingCompanyOrIndividualName != nil {
		ReceivingCompanyOrIndividualName = *r.ReceivingCompanyOrIndividualName
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetReceivingDepositoryFinancialInstitutionName() (ReceivingDepositoryFinancialInstitutionName string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionName != nil {
		ReceivingDepositoryFinancialInstitutionName = *r.ReceivingDepositoryFinancialInstitutionName
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetReceivingDepositoryFinancialInstitutionIdQualifier() (ReceivingDepositoryFinancialInstitutionIdQualifier string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionIdQualifier != nil {
		ReceivingDepositoryFinancialInstitutionIdQualifier = *r.ReceivingDepositoryFinancialInstitutionIdQualifier
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetReceivingDepositoryFinancialInstitutionId() (ReceivingDepositoryFinancialInstitutionId string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionId != nil {
		ReceivingDepositoryFinancialInstitutionId = *r.ReceivingDepositoryFinancialInstitutionId
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetReceivingDepositoryFinancialInstitutionCountry() (ReceivingDepositoryFinancialInstitutionCountry string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionCountry != nil {
		ReceivingDepositoryFinancialInstitutionCountry = *r.ReceivingDepositoryFinancialInstitutionCountry
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetTraceNumber() (TraceNumber string) {
	if r != nil && r.TraceNumber != nil {
		TraceNumber = *r.TraceNumber
	}
	return
}

//
type ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation struct {
	// The amount in the minor unit of the transfer's currency. For dollars, for
	// example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the transfer's
	// currency. This will always be "USD" for a Real Time Payments transfer.
	Currency *ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency `json:"currency"`
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
func (r *ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the transfer's
// currency. This will always be "USD" for a Real Time Payments transfer.
func (r *ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) GetCurrency() (Currency ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The name the sender of the transfer specified as the recipient of the transfer.
func (r *ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) GetCreditorName() (CreditorName string) {
	if r != nil && r.CreditorName != nil {
		CreditorName = *r.CreditorName
	}
	return
}

// The name provided by the sender of the transfer.
func (r *ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) GetDebtorName() (DebtorName string) {
	if r != nil && r.DebtorName != nil {
		DebtorName = *r.DebtorName
	}
	return
}

// The account number of the account that sent the transfer.
func (r *ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) GetDebtorAccountNumber() (DebtorAccountNumber string) {
	if r != nil && r.DebtorAccountNumber != nil {
		DebtorAccountNumber = *r.DebtorAccountNumber
	}
	return
}

// The routing number of the account that sent the transfer.
func (r *ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) GetDebtorRoutingNumber() (DebtorRoutingNumber string) {
	if r != nil && r.DebtorRoutingNumber != nil {
		DebtorRoutingNumber = *r.DebtorRoutingNumber
	}
	return
}

// The Real Time Payments network identification of the transfer
func (r *ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) GetTransactionIdentification() (TransactionIdentification string) {
	if r != nil && r.TransactionIdentification != nil {
		TransactionIdentification = *r.TransactionIdentification
	}
	return
}

// Additional information included with the transfer.
func (r *ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) GetRemittanceInformation() (RemittanceInformation string) {
	if r != nil && r.RemittanceInformation != nil {
		RemittanceInformation = *r.RemittanceInformation
	}
	return
}

type ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency string

const (
	ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyCad ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "CAD"
	ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyChf ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "CHF"
	ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyEur ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "EUR"
	ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyGbp ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "GBP"
	ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyJpy ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "JPY"
	ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyUsd ACHTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "USD"
)

//
type ACHTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal struct {
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
func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The description on the reversal message from Fedwire.
func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The Fedwire cycle date for the wire reversal.
func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) GetInputCycleDate() (InputCycleDate string) {
	if r != nil && r.InputCycleDate != nil {
		InputCycleDate = *r.InputCycleDate
	}
	return
}

// The Fedwire sequence number.
func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) GetInputSequenceNumber() (InputSequenceNumber string) {
	if r != nil && r.InputSequenceNumber != nil {
		InputSequenceNumber = *r.InputSequenceNumber
	}
	return
}

// The Fedwire input source identifier.
func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) GetInputSource() (InputSource string) {
	if r != nil && r.InputSource != nil {
		InputSource = *r.InputSource
	}
	return
}

// The Fedwire transaction identifier.
func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) GetInputMessageAccountabilityData() (InputMessageAccountabilityData string) {
	if r != nil && r.InputMessageAccountabilityData != nil {
		InputMessageAccountabilityData = *r.InputMessageAccountabilityData
	}
	return
}

// The Fedwire transaction identifier for the wire transfer that was reversed.
func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) GetPreviousMessageInputMessageAccountabilityData() (PreviousMessageInputMessageAccountabilityData string) {
	if r != nil && r.PreviousMessageInputMessageAccountabilityData != nil {
		PreviousMessageInputMessageAccountabilityData = *r.PreviousMessageInputMessageAccountabilityData
	}
	return
}

// The Fedwire cycle date for the wire transfer that was reversed.
func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) GetPreviousMessageInputCycleDate() (PreviousMessageInputCycleDate string) {
	if r != nil && r.PreviousMessageInputCycleDate != nil {
		PreviousMessageInputCycleDate = *r.PreviousMessageInputCycleDate
	}
	return
}

// The Fedwire sequence number for the wire transfer that was reversed.
func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) GetPreviousMessageInputSequenceNumber() (PreviousMessageInputSequenceNumber string) {
	if r != nil && r.PreviousMessageInputSequenceNumber != nil {
		PreviousMessageInputSequenceNumber = *r.PreviousMessageInputSequenceNumber
	}
	return
}

// The Fedwire input source identifier for the wire transfer that was reversed.
func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) GetPreviousMessageInputSource() (PreviousMessageInputSource string) {
	if r != nil && r.PreviousMessageInputSource != nil {
		PreviousMessageInputSource = *r.PreviousMessageInputSource
	}
	return
}

//
type ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment struct {
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
func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetBeneficiaryAddressLine1() (BeneficiaryAddressLine1 string) {
	if r != nil && r.BeneficiaryAddressLine1 != nil {
		BeneficiaryAddressLine1 = *r.BeneficiaryAddressLine1
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetBeneficiaryAddressLine2() (BeneficiaryAddressLine2 string) {
	if r != nil && r.BeneficiaryAddressLine2 != nil {
		BeneficiaryAddressLine2 = *r.BeneficiaryAddressLine2
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetBeneficiaryAddressLine3() (BeneficiaryAddressLine3 string) {
	if r != nil && r.BeneficiaryAddressLine3 != nil {
		BeneficiaryAddressLine3 = *r.BeneficiaryAddressLine3
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetBeneficiaryName() (BeneficiaryName string) {
	if r != nil && r.BeneficiaryName != nil {
		BeneficiaryName = *r.BeneficiaryName
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetBeneficiaryReference() (BeneficiaryReference string) {
	if r != nil && r.BeneficiaryReference != nil {
		BeneficiaryReference = *r.BeneficiaryReference
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetInputMessageAccountabilityData() (InputMessageAccountabilityData string) {
	if r != nil && r.InputMessageAccountabilityData != nil {
		InputMessageAccountabilityData = *r.InputMessageAccountabilityData
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetOriginatorAddressLine1() (OriginatorAddressLine1 string) {
	if r != nil && r.OriginatorAddressLine1 != nil {
		OriginatorAddressLine1 = *r.OriginatorAddressLine1
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetOriginatorAddressLine2() (OriginatorAddressLine2 string) {
	if r != nil && r.OriginatorAddressLine2 != nil {
		OriginatorAddressLine2 = *r.OriginatorAddressLine2
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetOriginatorAddressLine3() (OriginatorAddressLine3 string) {
	if r != nil && r.OriginatorAddressLine3 != nil {
		OriginatorAddressLine3 = *r.OriginatorAddressLine3
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetOriginatorName() (OriginatorName string) {
	if r != nil && r.OriginatorName != nil {
		OriginatorName = *r.OriginatorName
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetOriginatorToBeneficiaryInformation() (OriginatorToBeneficiaryInformation string) {
	if r != nil && r.OriginatorToBeneficiaryInformation != nil {
		OriginatorToBeneficiaryInformation = *r.OriginatorToBeneficiaryInformation
	}
	return
}

//
type ACHTransferSimulationTransactionSourceInboundWireReversal struct {
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
func (r *ACHTransferSimulationTransactionSourceInboundWireReversal) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The description on the reversal message from Fedwire.
func (r *ACHTransferSimulationTransactionSourceInboundWireReversal) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The Fedwire cycle date for the wire reversal.
func (r *ACHTransferSimulationTransactionSourceInboundWireReversal) GetInputCycleDate() (InputCycleDate string) {
	if r != nil && r.InputCycleDate != nil {
		InputCycleDate = *r.InputCycleDate
	}
	return
}

// The Fedwire sequence number.
func (r *ACHTransferSimulationTransactionSourceInboundWireReversal) GetInputSequenceNumber() (InputSequenceNumber string) {
	if r != nil && r.InputSequenceNumber != nil {
		InputSequenceNumber = *r.InputSequenceNumber
	}
	return
}

// The Fedwire input source identifier.
func (r *ACHTransferSimulationTransactionSourceInboundWireReversal) GetInputSource() (InputSource string) {
	if r != nil && r.InputSource != nil {
		InputSource = *r.InputSource
	}
	return
}

// The Fedwire transaction identifier.
func (r *ACHTransferSimulationTransactionSourceInboundWireReversal) GetInputMessageAccountabilityData() (InputMessageAccountabilityData string) {
	if r != nil && r.InputMessageAccountabilityData != nil {
		InputMessageAccountabilityData = *r.InputMessageAccountabilityData
	}
	return
}

// The Fedwire transaction identifier for the wire transfer that was reversed.
func (r *ACHTransferSimulationTransactionSourceInboundWireReversal) GetPreviousMessageInputMessageAccountabilityData() (PreviousMessageInputMessageAccountabilityData string) {
	if r != nil && r.PreviousMessageInputMessageAccountabilityData != nil {
		PreviousMessageInputMessageAccountabilityData = *r.PreviousMessageInputMessageAccountabilityData
	}
	return
}

// The Fedwire cycle date for the wire transfer that was reversed.
func (r *ACHTransferSimulationTransactionSourceInboundWireReversal) GetPreviousMessageInputCycleDate() (PreviousMessageInputCycleDate string) {
	if r != nil && r.PreviousMessageInputCycleDate != nil {
		PreviousMessageInputCycleDate = *r.PreviousMessageInputCycleDate
	}
	return
}

// The Fedwire sequence number for the wire transfer that was reversed.
func (r *ACHTransferSimulationTransactionSourceInboundWireReversal) GetPreviousMessageInputSequenceNumber() (PreviousMessageInputSequenceNumber string) {
	if r != nil && r.PreviousMessageInputSequenceNumber != nil {
		PreviousMessageInputSequenceNumber = *r.PreviousMessageInputSequenceNumber
	}
	return
}

// The Fedwire input source identifier for the wire transfer that was reversed.
func (r *ACHTransferSimulationTransactionSourceInboundWireReversal) GetPreviousMessageInputSource() (PreviousMessageInputSource string) {
	if r != nil && r.PreviousMessageInputSource != nil {
		PreviousMessageInputSource = *r.PreviousMessageInputSource
	}
	return
}

// Information included in the wire reversal for the receiving financial
// institution.
func (r *ACHTransferSimulationTransactionSourceInboundWireReversal) GetReceiverFinancialInstitutionInformation() (ReceiverFinancialInstitutionInformation string) {
	if r != nil && r.ReceiverFinancialInstitutionInformation != nil {
		ReceiverFinancialInstitutionInformation = *r.ReceiverFinancialInstitutionInformation
	}
	return
}

// Additional financial institution information included in the wire reversal.
func (r *ACHTransferSimulationTransactionSourceInboundWireReversal) GetFinancialInstitutionToFinancialInstitutionInformation() (FinancialInstitutionToFinancialInstitutionInformation string) {
	if r != nil && r.FinancialInstitutionToFinancialInstitutionInformation != nil {
		FinancialInstitutionToFinancialInstitutionInformation = *r.FinancialInstitutionToFinancialInstitutionInformation
	}
	return
}

//
type ACHTransferSimulationTransactionSourceInboundWireTransfer struct {
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
func (r *ACHTransferSimulationTransactionSourceInboundWireTransfer) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireTransfer) GetBeneficiaryAddressLine1() (BeneficiaryAddressLine1 string) {
	if r != nil && r.BeneficiaryAddressLine1 != nil {
		BeneficiaryAddressLine1 = *r.BeneficiaryAddressLine1
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireTransfer) GetBeneficiaryAddressLine2() (BeneficiaryAddressLine2 string) {
	if r != nil && r.BeneficiaryAddressLine2 != nil {
		BeneficiaryAddressLine2 = *r.BeneficiaryAddressLine2
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireTransfer) GetBeneficiaryAddressLine3() (BeneficiaryAddressLine3 string) {
	if r != nil && r.BeneficiaryAddressLine3 != nil {
		BeneficiaryAddressLine3 = *r.BeneficiaryAddressLine3
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireTransfer) GetBeneficiaryName() (BeneficiaryName string) {
	if r != nil && r.BeneficiaryName != nil {
		BeneficiaryName = *r.BeneficiaryName
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireTransfer) GetBeneficiaryReference() (BeneficiaryReference string) {
	if r != nil && r.BeneficiaryReference != nil {
		BeneficiaryReference = *r.BeneficiaryReference
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireTransfer) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireTransfer) GetInputMessageAccountabilityData() (InputMessageAccountabilityData string) {
	if r != nil && r.InputMessageAccountabilityData != nil {
		InputMessageAccountabilityData = *r.InputMessageAccountabilityData
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireTransfer) GetOriginatorAddressLine1() (OriginatorAddressLine1 string) {
	if r != nil && r.OriginatorAddressLine1 != nil {
		OriginatorAddressLine1 = *r.OriginatorAddressLine1
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireTransfer) GetOriginatorAddressLine2() (OriginatorAddressLine2 string) {
	if r != nil && r.OriginatorAddressLine2 != nil {
		OriginatorAddressLine2 = *r.OriginatorAddressLine2
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireTransfer) GetOriginatorAddressLine3() (OriginatorAddressLine3 string) {
	if r != nil && r.OriginatorAddressLine3 != nil {
		OriginatorAddressLine3 = *r.OriginatorAddressLine3
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireTransfer) GetOriginatorName() (OriginatorName string) {
	if r != nil && r.OriginatorName != nil {
		OriginatorName = *r.OriginatorName
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInboundWireTransfer) GetOriginatorToBeneficiaryInformation() (OriginatorToBeneficiaryInformation string) {
	if r != nil && r.OriginatorToBeneficiaryInformation != nil {
		OriginatorToBeneficiaryInformation = *r.OriginatorToBeneficiaryInformation
	}
	return
}

//
type ACHTransferSimulationTransactionSourceInternalSource struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency *ACHTransferSimulationTransactionSourceInternalSourceCurrency `json:"currency"`
	//
	Reason *ACHTransferSimulationTransactionSourceInternalSourceReason `json:"reason"`
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *ACHTransferSimulationTransactionSourceInternalSource) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
// currency.
func (r *ACHTransferSimulationTransactionSourceInternalSource) GetCurrency() (Currency ACHTransferSimulationTransactionSourceInternalSourceCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceInternalSource) GetReason() (Reason ACHTransferSimulationTransactionSourceInternalSourceReason) {
	if r != nil && r.Reason != nil {
		Reason = *r.Reason
	}
	return
}

type ACHTransferSimulationTransactionSourceInternalSourceCurrency string

const (
	ACHTransferSimulationTransactionSourceInternalSourceCurrencyCad ACHTransferSimulationTransactionSourceInternalSourceCurrency = "CAD"
	ACHTransferSimulationTransactionSourceInternalSourceCurrencyChf ACHTransferSimulationTransactionSourceInternalSourceCurrency = "CHF"
	ACHTransferSimulationTransactionSourceInternalSourceCurrencyEur ACHTransferSimulationTransactionSourceInternalSourceCurrency = "EUR"
	ACHTransferSimulationTransactionSourceInternalSourceCurrencyGbp ACHTransferSimulationTransactionSourceInternalSourceCurrency = "GBP"
	ACHTransferSimulationTransactionSourceInternalSourceCurrencyJpy ACHTransferSimulationTransactionSourceInternalSourceCurrency = "JPY"
	ACHTransferSimulationTransactionSourceInternalSourceCurrencyUsd ACHTransferSimulationTransactionSourceInternalSourceCurrency = "USD"
)

type ACHTransferSimulationTransactionSourceInternalSourceReason string

const (
	ACHTransferSimulationTransactionSourceInternalSourceReasonCashback           ACHTransferSimulationTransactionSourceInternalSourceReason = "cashback"
	ACHTransferSimulationTransactionSourceInternalSourceReasonEmpyrealAdjustment ACHTransferSimulationTransactionSourceInternalSourceReason = "empyreal_adjustment"
	ACHTransferSimulationTransactionSourceInternalSourceReasonError              ACHTransferSimulationTransactionSourceInternalSourceReason = "error"
	ACHTransferSimulationTransactionSourceInternalSourceReasonErrorCorrection    ACHTransferSimulationTransactionSourceInternalSourceReason = "error_correction"
	ACHTransferSimulationTransactionSourceInternalSourceReasonFees               ACHTransferSimulationTransactionSourceInternalSourceReason = "fees"
	ACHTransferSimulationTransactionSourceInternalSourceReasonInterest           ACHTransferSimulationTransactionSourceInternalSourceReason = "interest"
	ACHTransferSimulationTransactionSourceInternalSourceReasonSampleFunds        ACHTransferSimulationTransactionSourceInternalSourceReason = "sample_funds"
	ACHTransferSimulationTransactionSourceInternalSourceReasonSampleFundsReturn  ACHTransferSimulationTransactionSourceInternalSourceReason = "sample_funds_return"
)

//
type ACHTransferSimulationTransactionSourceCardRouteRefund struct {
	// The refunded amount in the minor unit of the refunded currency. For dollars, for
	// example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the refund
	// currency.
	Currency *ACHTransferSimulationTransactionSourceCardRouteRefundCurrency `json:"currency"`
	//
	MerchantAcceptorId *string `json:"merchant_acceptor_id"`
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
func (r *ACHTransferSimulationTransactionSourceCardRouteRefund) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the refund
// currency.
func (r *ACHTransferSimulationTransactionSourceCardRouteRefund) GetCurrency() (Currency ACHTransferSimulationTransactionSourceCardRouteRefundCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceCardRouteRefund) GetMerchantAcceptorId() (MerchantAcceptorId string) {
	if r != nil && r.MerchantAcceptorId != nil {
		MerchantAcceptorId = *r.MerchantAcceptorId
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceCardRouteRefund) GetMerchantCity() (MerchantCity string) {
	if r != nil && r.MerchantCity != nil {
		MerchantCity = *r.MerchantCity
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceCardRouteRefund) GetMerchantCountry() (MerchantCountry string) {
	if r != nil && r.MerchantCountry != nil {
		MerchantCountry = *r.MerchantCountry
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceCardRouteRefund) GetMerchantDescriptor() (MerchantDescriptor string) {
	if r != nil && r.MerchantDescriptor != nil {
		MerchantDescriptor = *r.MerchantDescriptor
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceCardRouteRefund) GetMerchantState() (MerchantState string) {
	if r != nil && r.MerchantState != nil {
		MerchantState = *r.MerchantState
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceCardRouteRefund) GetMerchantCategoryCode() (MerchantCategoryCode string) {
	if r != nil && r.MerchantCategoryCode != nil {
		MerchantCategoryCode = *r.MerchantCategoryCode
	}
	return
}

type ACHTransferSimulationTransactionSourceCardRouteRefundCurrency string

const (
	ACHTransferSimulationTransactionSourceCardRouteRefundCurrencyCad ACHTransferSimulationTransactionSourceCardRouteRefundCurrency = "CAD"
	ACHTransferSimulationTransactionSourceCardRouteRefundCurrencyChf ACHTransferSimulationTransactionSourceCardRouteRefundCurrency = "CHF"
	ACHTransferSimulationTransactionSourceCardRouteRefundCurrencyEur ACHTransferSimulationTransactionSourceCardRouteRefundCurrency = "EUR"
	ACHTransferSimulationTransactionSourceCardRouteRefundCurrencyGbp ACHTransferSimulationTransactionSourceCardRouteRefundCurrency = "GBP"
	ACHTransferSimulationTransactionSourceCardRouteRefundCurrencyJpy ACHTransferSimulationTransactionSourceCardRouteRefundCurrency = "JPY"
	ACHTransferSimulationTransactionSourceCardRouteRefundCurrencyUsd ACHTransferSimulationTransactionSourceCardRouteRefundCurrency = "USD"
)

//
type ACHTransferSimulationTransactionSourceCardRouteSettlement struct {
	// The settled amount in the minor unit of the settlement currency. For dollars,
	// for example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the settlement
	// currency.
	Currency *ACHTransferSimulationTransactionSourceCardRouteSettlementCurrency `json:"currency"`
	//
	MerchantAcceptorId *string `json:"merchant_acceptor_id"`
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
func (r *ACHTransferSimulationTransactionSourceCardRouteSettlement) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the settlement
// currency.
func (r *ACHTransferSimulationTransactionSourceCardRouteSettlement) GetCurrency() (Currency ACHTransferSimulationTransactionSourceCardRouteSettlementCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceCardRouteSettlement) GetMerchantAcceptorId() (MerchantAcceptorId string) {
	if r != nil && r.MerchantAcceptorId != nil {
		MerchantAcceptorId = *r.MerchantAcceptorId
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceCardRouteSettlement) GetMerchantCity() (MerchantCity string) {
	if r != nil && r.MerchantCity != nil {
		MerchantCity = *r.MerchantCity
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceCardRouteSettlement) GetMerchantCountry() (MerchantCountry string) {
	if r != nil && r.MerchantCountry != nil {
		MerchantCountry = *r.MerchantCountry
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceCardRouteSettlement) GetMerchantDescriptor() (MerchantDescriptor string) {
	if r != nil && r.MerchantDescriptor != nil {
		MerchantDescriptor = *r.MerchantDescriptor
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceCardRouteSettlement) GetMerchantState() (MerchantState string) {
	if r != nil && r.MerchantState != nil {
		MerchantState = *r.MerchantState
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceCardRouteSettlement) GetMerchantCategoryCode() (MerchantCategoryCode string) {
	if r != nil && r.MerchantCategoryCode != nil {
		MerchantCategoryCode = *r.MerchantCategoryCode
	}
	return
}

type ACHTransferSimulationTransactionSourceCardRouteSettlementCurrency string

const (
	ACHTransferSimulationTransactionSourceCardRouteSettlementCurrencyCad ACHTransferSimulationTransactionSourceCardRouteSettlementCurrency = "CAD"
	ACHTransferSimulationTransactionSourceCardRouteSettlementCurrencyChf ACHTransferSimulationTransactionSourceCardRouteSettlementCurrency = "CHF"
	ACHTransferSimulationTransactionSourceCardRouteSettlementCurrencyEur ACHTransferSimulationTransactionSourceCardRouteSettlementCurrency = "EUR"
	ACHTransferSimulationTransactionSourceCardRouteSettlementCurrencyGbp ACHTransferSimulationTransactionSourceCardRouteSettlementCurrency = "GBP"
	ACHTransferSimulationTransactionSourceCardRouteSettlementCurrencyJpy ACHTransferSimulationTransactionSourceCardRouteSettlementCurrency = "JPY"
	ACHTransferSimulationTransactionSourceCardRouteSettlementCurrencyUsd ACHTransferSimulationTransactionSourceCardRouteSettlementCurrency = "USD"
)

//
type ACHTransferSimulationTransactionSourceSampleFunds struct {
	// Where the sample funds came from.
	Originator *string `json:"originator"`
}

// Where the sample funds came from.
func (r *ACHTransferSimulationTransactionSourceSampleFunds) GetOriginator() (Originator string) {
	if r != nil && r.Originator != nil {
		Originator = *r.Originator
	}
	return
}

//
type ACHTransferSimulationTransactionSourceWireDrawdownPaymentIntention struct {
	// The transfer amount in USD cents.
	Amount *int `json:"amount"`
	//
	AccountNumber *string `json:"account_number"`
	//
	RoutingNumber *string `json:"routing_number"`
	//
	MessageToRecipient *string `json:"message_to_recipient"`
	//
	TransferId *string `json:"transfer_id"`
}

// The transfer amount in USD cents.
func (r *ACHTransferSimulationTransactionSourceWireDrawdownPaymentIntention) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceWireDrawdownPaymentIntention) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceWireDrawdownPaymentIntention) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceWireDrawdownPaymentIntention) GetMessageToRecipient() (MessageToRecipient string) {
	if r != nil && r.MessageToRecipient != nil {
		MessageToRecipient = *r.MessageToRecipient
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceWireDrawdownPaymentIntention) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}

//
type ACHTransferSimulationTransactionSourceWireDrawdownPaymentRejection struct {
	//
	TransferId *string `json:"transfer_id"`
}

func (r *ACHTransferSimulationTransactionSourceWireDrawdownPaymentRejection) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}

//
type ACHTransferSimulationTransactionSourceWireTransferIntention struct {
	// The transfer amount in USD cents.
	Amount *int `json:"amount"`
	// The destination account number.
	AccountNumber *string `json:"account_number"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber *string `json:"routing_number"`
	// The message that will show on the recipient's bank statement.
	MessageToRecipient *string `json:"message_to_recipient"`
	//
	TransferId *string `json:"transfer_id"`
}

// The transfer amount in USD cents.
func (r *ACHTransferSimulationTransactionSourceWireTransferIntention) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The destination account number.
func (r *ACHTransferSimulationTransactionSourceWireTransferIntention) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

// The American Bankers' Association (ABA) Routing Transit Number (RTN).
func (r *ACHTransferSimulationTransactionSourceWireTransferIntention) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

// The message that will show on the recipient's bank statement.
func (r *ACHTransferSimulationTransactionSourceWireTransferIntention) GetMessageToRecipient() (MessageToRecipient string) {
	if r != nil && r.MessageToRecipient != nil {
		MessageToRecipient = *r.MessageToRecipient
	}
	return
}

func (r *ACHTransferSimulationTransactionSourceWireTransferIntention) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}

//
type ACHTransferSimulationTransactionSourceWireTransferRejection struct {
	//
	TransferId *string `json:"transfer_id"`
}

func (r *ACHTransferSimulationTransactionSourceWireTransferRejection) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}

type ACHTransferSimulationTransactionType string

const (
	ACHTransferSimulationTransactionTypeTransaction ACHTransferSimulationTransactionType = "transaction"
)

//
type ACHTransferSimulationDeclinedTransaction struct {
	// The identifier for the Account the Declined Transaction belongs to.
	AccountId *string `json:"account_id"`
	// The Declined Transaction amount in the minor unit of its currency. For dollars,
	// for example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Declined
	// Transaction's currency. This will match the currency on the Declined
	// Transcation's Account.
	Currency *ACHTransferSimulationDeclinedTransactionCurrency `json:"currency"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the
	// Transaction occured.
	CreatedAt *string `json:"created_at"`
	// This is the description the vendor provides.
	Description *string `json:"description"`
	// The Declined Transaction identifier.
	Id *string `json:"id"`
	// The identifier for the route this Declined Transaction came through. Routes are
	// things like cards and ACH details.
	RouteId *string `json:"route_id"`
	// The type of the route this Declined Transaction came through.
	RouteType *string `json:"route_type"`
	// This is an object giving more details on the network-level event that caused the
	// Declined Transaction. For example, for a card transaction this lists the
	// merchant's industry and location. Note that for backwards compatibility reasons,
	// additional undocumented keys may appear in this object. These should be treated
	// as deprecated and will be removed in the future.
	Source *ACHTransferSimulationDeclinedTransactionSource `json:"source"`
	// A constant representing the object's type. For this resource it will always be
	// `declined_transaction`.
	Type *ACHTransferSimulationDeclinedTransactionType `json:"type"`
}

// The identifier for the Account the Declined Transaction belongs to.
func (r *ACHTransferSimulationDeclinedTransaction) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

// The Declined Transaction amount in the minor unit of its currency. For dollars,
// for example, this is cents.
func (r *ACHTransferSimulationDeclinedTransaction) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Declined
// Transaction's currency. This will match the currency on the Declined
// Transcation's Account.
func (r *ACHTransferSimulationDeclinedTransaction) GetCurrency() (Currency ACHTransferSimulationDeclinedTransactionCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the
// Transaction occured.
func (r *ACHTransferSimulationDeclinedTransaction) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// This is the description the vendor provides.
func (r *ACHTransferSimulationDeclinedTransaction) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The Declined Transaction identifier.
func (r *ACHTransferSimulationDeclinedTransaction) GetId() (Id string) {
	if r != nil && r.Id != nil {
		Id = *r.Id
	}
	return
}

// The identifier for the route this Declined Transaction came through. Routes are
// things like cards and ACH details.
func (r *ACHTransferSimulationDeclinedTransaction) GetRouteId() (RouteId string) {
	if r != nil && r.RouteId != nil {
		RouteId = *r.RouteId
	}
	return
}

// The type of the route this Declined Transaction came through.
func (r *ACHTransferSimulationDeclinedTransaction) GetRouteType() (RouteType string) {
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
func (r *ACHTransferSimulationDeclinedTransaction) GetSource() (Source ACHTransferSimulationDeclinedTransactionSource) {
	if r != nil && r.Source != nil {
		Source = *r.Source
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `declined_transaction`.
func (r *ACHTransferSimulationDeclinedTransaction) GetType() (Type ACHTransferSimulationDeclinedTransactionType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

type ACHTransferSimulationDeclinedTransactionCurrency string

const (
	ACHTransferSimulationDeclinedTransactionCurrencyCad ACHTransferSimulationDeclinedTransactionCurrency = "CAD"
	ACHTransferSimulationDeclinedTransactionCurrencyChf ACHTransferSimulationDeclinedTransactionCurrency = "CHF"
	ACHTransferSimulationDeclinedTransactionCurrencyEur ACHTransferSimulationDeclinedTransactionCurrency = "EUR"
	ACHTransferSimulationDeclinedTransactionCurrencyGbp ACHTransferSimulationDeclinedTransactionCurrency = "GBP"
	ACHTransferSimulationDeclinedTransactionCurrencyJpy ACHTransferSimulationDeclinedTransactionCurrency = "JPY"
	ACHTransferSimulationDeclinedTransactionCurrencyUsd ACHTransferSimulationDeclinedTransactionCurrency = "USD"
)

//
type ACHTransferSimulationDeclinedTransactionSource struct {
	// The type of decline that took place. We may add additional possible values for
	// this enum over time; your application should be able to handle such additions
	// gracefully.
	Category *ACHTransferSimulationDeclinedTransactionSourceCategory `json:"category"`
	// A ACH Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `ach_decline`.
	ACHDecline *ACHTransferSimulationDeclinedTransactionSourceACHDecline `json:"ach_decline"`
	// A Card Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_decline`.
	CardDecline *ACHTransferSimulationDeclinedTransactionSourceCardDecline `json:"card_decline"`
	// A Check Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `check_decline`.
	CheckDecline *ACHTransferSimulationDeclinedTransactionSourceCheckDecline `json:"check_decline"`
	// A Inbound Real Time Payments Transfer Decline object. This field will be present
	// in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_decline`.
	InboundRealTimePaymentsTransferDecline *ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline `json:"inbound_real_time_payments_transfer_decline"`
	// A International ACH Decline object. This field will be present in the JSON
	// response if and only if `category` is equal to `international_ach_decline`.
	InternationalACHDecline *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline `json:"international_ach_decline"`
	// A Deprecated Card Decline object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_route_decline`.
	CardRouteDecline *ACHTransferSimulationDeclinedTransactionSourceCardRouteDecline `json:"card_route_decline"`
}

// The type of decline that took place. We may add additional possible values for
// this enum over time; your application should be able to handle such additions
// gracefully.
func (r *ACHTransferSimulationDeclinedTransactionSource) GetCategory() (Category ACHTransferSimulationDeclinedTransactionSourceCategory) {
	if r != nil && r.Category != nil {
		Category = *r.Category
	}
	return
}

// A ACH Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `ach_decline`.
func (r *ACHTransferSimulationDeclinedTransactionSource) GetACHDecline() (ACHDecline ACHTransferSimulationDeclinedTransactionSourceACHDecline) {
	if r != nil && r.ACHDecline != nil {
		ACHDecline = *r.ACHDecline
	}
	return
}

// A Card Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `card_decline`.
func (r *ACHTransferSimulationDeclinedTransactionSource) GetCardDecline() (CardDecline ACHTransferSimulationDeclinedTransactionSourceCardDecline) {
	if r != nil && r.CardDecline != nil {
		CardDecline = *r.CardDecline
	}
	return
}

// A Check Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `check_decline`.
func (r *ACHTransferSimulationDeclinedTransactionSource) GetCheckDecline() (CheckDecline ACHTransferSimulationDeclinedTransactionSourceCheckDecline) {
	if r != nil && r.CheckDecline != nil {
		CheckDecline = *r.CheckDecline
	}
	return
}

// A Inbound Real Time Payments Transfer Decline object. This field will be present
// in the JSON response if and only if `category` is equal to
// `inbound_real_time_payments_transfer_decline`.
func (r *ACHTransferSimulationDeclinedTransactionSource) GetInboundRealTimePaymentsTransferDecline() (InboundRealTimePaymentsTransferDecline ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) {
	if r != nil && r.InboundRealTimePaymentsTransferDecline != nil {
		InboundRealTimePaymentsTransferDecline = *r.InboundRealTimePaymentsTransferDecline
	}
	return
}

// A International ACH Decline object. This field will be present in the JSON
// response if and only if `category` is equal to `international_ach_decline`.
func (r *ACHTransferSimulationDeclinedTransactionSource) GetInternationalACHDecline() (InternationalACHDecline ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) {
	if r != nil && r.InternationalACHDecline != nil {
		InternationalACHDecline = *r.InternationalACHDecline
	}
	return
}

// A Deprecated Card Decline object. This field will be present in the JSON
// response if and only if `category` is equal to `card_route_decline`.
func (r *ACHTransferSimulationDeclinedTransactionSource) GetCardRouteDecline() (CardRouteDecline ACHTransferSimulationDeclinedTransactionSourceCardRouteDecline) {
	if r != nil && r.CardRouteDecline != nil {
		CardRouteDecline = *r.CardRouteDecline
	}
	return
}

type ACHTransferSimulationDeclinedTransactionSourceCategory string

const (
	ACHTransferSimulationDeclinedTransactionSourceCategoryACHDecline                             ACHTransferSimulationDeclinedTransactionSourceCategory = "ach_decline"
	ACHTransferSimulationDeclinedTransactionSourceCategoryCardDecline                            ACHTransferSimulationDeclinedTransactionSourceCategory = "card_decline"
	ACHTransferSimulationDeclinedTransactionSourceCategoryCheckDecline                           ACHTransferSimulationDeclinedTransactionSourceCategory = "check_decline"
	ACHTransferSimulationDeclinedTransactionSourceCategoryInboundRealTimePaymentsTransferDecline ACHTransferSimulationDeclinedTransactionSourceCategory = "inbound_real_time_payments_transfer_decline"
	ACHTransferSimulationDeclinedTransactionSourceCategoryInternationalACHDecline                ACHTransferSimulationDeclinedTransactionSourceCategory = "international_ach_decline"
	ACHTransferSimulationDeclinedTransactionSourceCategoryCardRouteDecline                       ACHTransferSimulationDeclinedTransactionSourceCategory = "card_route_decline"
	ACHTransferSimulationDeclinedTransactionSourceCategoryOther                                  ACHTransferSimulationDeclinedTransactionSourceCategory = "other"
)

//
type ACHTransferSimulationDeclinedTransactionSourceACHDecline struct {
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
	OriginatorCompanyId *string `json:"originator_company_id"`
	// Why the ACH transfer was declined.
	Reason *ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason `json:"reason"`
	//
	ReceiverIdNumber *string `json:"receiver_id_number"`
	//
	ReceiverName *string `json:"receiver_name"`
	//
	TraceNumber *string `json:"trace_number"`
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *ACHTransferSimulationDeclinedTransactionSourceACHDecline) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceACHDecline) GetOriginatorCompanyName() (OriginatorCompanyName string) {
	if r != nil && r.OriginatorCompanyName != nil {
		OriginatorCompanyName = *r.OriginatorCompanyName
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceACHDecline) GetOriginatorCompanyDescriptiveDate() (OriginatorCompanyDescriptiveDate string) {
	if r != nil && r.OriginatorCompanyDescriptiveDate != nil {
		OriginatorCompanyDescriptiveDate = *r.OriginatorCompanyDescriptiveDate
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceACHDecline) GetOriginatorCompanyDiscretionaryData() (OriginatorCompanyDiscretionaryData string) {
	if r != nil && r.OriginatorCompanyDiscretionaryData != nil {
		OriginatorCompanyDiscretionaryData = *r.OriginatorCompanyDiscretionaryData
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceACHDecline) GetOriginatorCompanyId() (OriginatorCompanyId string) {
	if r != nil && r.OriginatorCompanyId != nil {
		OriginatorCompanyId = *r.OriginatorCompanyId
	}
	return
}

// Why the ACH transfer was declined.
func (r *ACHTransferSimulationDeclinedTransactionSourceACHDecline) GetReason() (Reason ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason) {
	if r != nil && r.Reason != nil {
		Reason = *r.Reason
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceACHDecline) GetReceiverIdNumber() (ReceiverIdNumber string) {
	if r != nil && r.ReceiverIdNumber != nil {
		ReceiverIdNumber = *r.ReceiverIdNumber
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceACHDecline) GetReceiverName() (ReceiverName string) {
	if r != nil && r.ReceiverName != nil {
		ReceiverName = *r.ReceiverName
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceACHDecline) GetTraceNumber() (TraceNumber string) {
	if r != nil && r.TraceNumber != nil {
		TraceNumber = *r.TraceNumber
	}
	return
}

type ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason string

const (
	ACHTransferSimulationDeclinedTransactionSourceACHDeclineReasonACHRouteCanceled             ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason = "ach_route_canceled"
	ACHTransferSimulationDeclinedTransactionSourceACHDeclineReasonACHRouteDisabled             ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason = "ach_route_disabled"
	ACHTransferSimulationDeclinedTransactionSourceACHDeclineReasonNoACHRoute                   ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason = "no_ach_route"
	ACHTransferSimulationDeclinedTransactionSourceACHDeclineReasonBreachesLimit                ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason = "breaches_limit"
	ACHTransferSimulationDeclinedTransactionSourceACHDeclineReasonCreditEntryRefusedByReceiver ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason = "credit_entry_refused_by_receiver"
	ACHTransferSimulationDeclinedTransactionSourceACHDeclineReasonGroupLocked                  ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason = "group_locked"
	ACHTransferSimulationDeclinedTransactionSourceACHDeclineReasonEntityNotActive              ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason = "entity_not_active"
	ACHTransferSimulationDeclinedTransactionSourceACHDeclineReasonInsufficientFunds            ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason = "insufficient_funds"
	ACHTransferSimulationDeclinedTransactionSourceACHDeclineReasonOriginatorRequest            ACHTransferSimulationDeclinedTransactionSourceACHDeclineReason = "originator_request"
)

//
type ACHTransferSimulationDeclinedTransactionSourceCardDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency *ACHTransferSimulationDeclinedTransactionSourceCardDeclineCurrency `json:"currency"`
	//
	MerchantAcceptorId *string `json:"merchant_acceptor_id"`
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
	Reason *ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason `json:"reason"`
	// The identifier of the Real-Time Decision sent to approve or decline this
	// transaction.
	RealTimeDecisionId *string `json:"real_time_decision_id"`
	// If the authorization was attempted using a Digital Wallet Token (such as an
	// Apple Pay purchase), the identifier of the token that was used.
	DigitalWalletTokenId *string `json:"digital_wallet_token_id"`
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *ACHTransferSimulationDeclinedTransactionSourceCardDecline) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
// account currency.
func (r *ACHTransferSimulationDeclinedTransactionSourceCardDecline) GetCurrency() (Currency ACHTransferSimulationDeclinedTransactionSourceCardDeclineCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceCardDecline) GetMerchantAcceptorId() (MerchantAcceptorId string) {
	if r != nil && r.MerchantAcceptorId != nil {
		MerchantAcceptorId = *r.MerchantAcceptorId
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceCardDecline) GetMerchantCity() (MerchantCity string) {
	if r != nil && r.MerchantCity != nil {
		MerchantCity = *r.MerchantCity
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceCardDecline) GetMerchantCountry() (MerchantCountry string) {
	if r != nil && r.MerchantCountry != nil {
		MerchantCountry = *r.MerchantCountry
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceCardDecline) GetMerchantDescriptor() (MerchantDescriptor string) {
	if r != nil && r.MerchantDescriptor != nil {
		MerchantDescriptor = *r.MerchantDescriptor
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceCardDecline) GetMerchantState() (MerchantState string) {
	if r != nil && r.MerchantState != nil {
		MerchantState = *r.MerchantState
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceCardDecline) GetMerchantCategoryCode() (MerchantCategoryCode string) {
	if r != nil && r.MerchantCategoryCode != nil {
		MerchantCategoryCode = *r.MerchantCategoryCode
	}
	return
}

// Why the transaction was declined.
func (r *ACHTransferSimulationDeclinedTransactionSourceCardDecline) GetReason() (Reason ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason) {
	if r != nil && r.Reason != nil {
		Reason = *r.Reason
	}
	return
}

// The identifier of the Real-Time Decision sent to approve or decline this
// transaction.
func (r *ACHTransferSimulationDeclinedTransactionSourceCardDecline) GetRealTimeDecisionId() (RealTimeDecisionId string) {
	if r != nil && r.RealTimeDecisionId != nil {
		RealTimeDecisionId = *r.RealTimeDecisionId
	}
	return
}

// If the authorization was attempted using a Digital Wallet Token (such as an
// Apple Pay purchase), the identifier of the token that was used.
func (r *ACHTransferSimulationDeclinedTransactionSourceCardDecline) GetDigitalWalletTokenId() (DigitalWalletTokenId string) {
	if r != nil && r.DigitalWalletTokenId != nil {
		DigitalWalletTokenId = *r.DigitalWalletTokenId
	}
	return
}

type ACHTransferSimulationDeclinedTransactionSourceCardDeclineCurrency string

const (
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineCurrencyCad ACHTransferSimulationDeclinedTransactionSourceCardDeclineCurrency = "CAD"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineCurrencyChf ACHTransferSimulationDeclinedTransactionSourceCardDeclineCurrency = "CHF"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineCurrencyEur ACHTransferSimulationDeclinedTransactionSourceCardDeclineCurrency = "EUR"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineCurrencyGbp ACHTransferSimulationDeclinedTransactionSourceCardDeclineCurrency = "GBP"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineCurrencyJpy ACHTransferSimulationDeclinedTransactionSourceCardDeclineCurrency = "JPY"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineCurrencyUsd ACHTransferSimulationDeclinedTransactionSourceCardDeclineCurrency = "USD"
)

type ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason string

const (
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineReasonCardNotActive     ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason = "card_not_active"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineReasonEntityNotActive   ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason = "entity_not_active"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineReasonGroupLocked       ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason = "group_locked"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineReasonInsufficientFunds ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason = "insufficient_funds"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineReasonBreachesLimit     ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason = "breaches_limit"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineReasonWebhookDeclined   ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason = "webhook_declined"
	ACHTransferSimulationDeclinedTransactionSourceCardDeclineReasonWebhookTimedOut   ACHTransferSimulationDeclinedTransactionSourceCardDeclineReason = "webhook_timed_out"
)

//
type ACHTransferSimulationDeclinedTransactionSourceCheckDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount *int `json:"amount"`
	//
	AuxiliaryOnUs *string `json:"auxiliary_on_us"`
	// Why the check was declined.
	Reason *ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReason `json:"reason"`
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *ACHTransferSimulationDeclinedTransactionSourceCheckDecline) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceCheckDecline) GetAuxiliaryOnUs() (AuxiliaryOnUs string) {
	if r != nil && r.AuxiliaryOnUs != nil {
		AuxiliaryOnUs = *r.AuxiliaryOnUs
	}
	return
}

// Why the check was declined.
func (r *ACHTransferSimulationDeclinedTransactionSourceCheckDecline) GetReason() (Reason ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReason) {
	if r != nil && r.Reason != nil {
		Reason = *r.Reason
	}
	return
}

type ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReason string

const (
	ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReasonACHRouteCanceled      ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReason = "ach_route_canceled"
	ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReasonACHRouteDisabled      ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReason = "ach_route_disabled"
	ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReasonBreachesLimit         ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReason = "breaches_limit"
	ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReasonEntityNotActive       ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReason = "entity_not_active"
	ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReasonGroupLocked           ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReason = "group_locked"
	ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReasonInsufficientFunds     ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReason = "insufficient_funds"
	ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReasonUnableToLocateAccount ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReason = "unable_to_locate_account"
	ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReasonUnableToProcess       ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReason = "unable_to_process"
	ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReasonReferToImage          ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReason = "refer_to_image"
	ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReasonStopPaymentRequested  ACHTransferSimulationDeclinedTransactionSourceCheckDeclineReason = "stop_payment_requested"
)

//
type ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the declined
	// transfer's currency. This will always be "USD" for a Real Time Payments
	// transfer.
	Currency *ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency `json:"currency"`
	// Why the transfer was declined.
	Reason *ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason `json:"reason"`
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
func (r *ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the declined
// transfer's currency. This will always be "USD" for a Real Time Payments
// transfer.
func (r *ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) GetCurrency() (Currency ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// Why the transfer was declined.
func (r *ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) GetReason() (Reason ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason) {
	if r != nil && r.Reason != nil {
		Reason = *r.Reason
	}
	return
}

// The name the sender of the transfer specified as the recipient of the transfer.
func (r *ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) GetCreditorName() (CreditorName string) {
	if r != nil && r.CreditorName != nil {
		CreditorName = *r.CreditorName
	}
	return
}

// The name provided by the sender of the transfer.
func (r *ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) GetDebtorName() (DebtorName string) {
	if r != nil && r.DebtorName != nil {
		DebtorName = *r.DebtorName
	}
	return
}

// The account number of the account that sent the transfer.
func (r *ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) GetDebtorAccountNumber() (DebtorAccountNumber string) {
	if r != nil && r.DebtorAccountNumber != nil {
		DebtorAccountNumber = *r.DebtorAccountNumber
	}
	return
}

// The routing number of the account that sent the transfer.
func (r *ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) GetDebtorRoutingNumber() (DebtorRoutingNumber string) {
	if r != nil && r.DebtorRoutingNumber != nil {
		DebtorRoutingNumber = *r.DebtorRoutingNumber
	}
	return
}

// The Real Time Payments network identification of the declined transfer.
func (r *ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) GetTransactionIdentification() (TransactionIdentification string) {
	if r != nil && r.TransactionIdentification != nil {
		TransactionIdentification = *r.TransactionIdentification
	}
	return
}

// Additional information included with the transfer.
func (r *ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) GetRemittanceInformation() (RemittanceInformation string) {
	if r != nil && r.RemittanceInformation != nil {
		RemittanceInformation = *r.RemittanceInformation
	}
	return
}

type ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency string

const (
	ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyCad ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "CAD"
	ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyChf ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "CHF"
	ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyEur ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "EUR"
	ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyGbp ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "GBP"
	ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyJpy ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "JPY"
	ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyUsd ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "USD"
)

type ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason string

const (
	ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonAccountNumberCanceled      ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "account_number_canceled"
	ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonAccountNumberDisabled      ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "account_number_disabled"
	ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonGroupLocked                ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "group_locked"
	ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonEntityNotActive            ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "entity_not_active"
	ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonRealTimePaymentsNotEnabled ACHTransferSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "real_time_payments_not_enabled"
)

//
type ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline struct {
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
	OriginatingDepositoryFinancialInstitutionIdQualifier *string `json:"originating_depository_financial_institution_id_qualifier"`
	//
	OriginatingDepositoryFinancialInstitutionId *string `json:"originating_depository_financial_institution_id"`
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
	ReceivingDepositoryFinancialInstitutionIdQualifier *string `json:"receiving_depository_financial_institution_id_qualifier"`
	//
	ReceivingDepositoryFinancialInstitutionId *string `json:"receiving_depository_financial_institution_id"`
	//
	ReceivingDepositoryFinancialInstitutionCountry *string `json:"receiving_depository_financial_institution_country"`
	//
	TraceNumber *string `json:"trace_number"`
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetForeignExchangeIndicator() (ForeignExchangeIndicator string) {
	if r != nil && r.ForeignExchangeIndicator != nil {
		ForeignExchangeIndicator = *r.ForeignExchangeIndicator
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetForeignExchangeReferenceIndicator() (ForeignExchangeReferenceIndicator string) {
	if r != nil && r.ForeignExchangeReferenceIndicator != nil {
		ForeignExchangeReferenceIndicator = *r.ForeignExchangeReferenceIndicator
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetForeignExchangeReference() (ForeignExchangeReference string) {
	if r != nil && r.ForeignExchangeReference != nil {
		ForeignExchangeReference = *r.ForeignExchangeReference
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetDestinationCountryCode() (DestinationCountryCode string) {
	if r != nil && r.DestinationCountryCode != nil {
		DestinationCountryCode = *r.DestinationCountryCode
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetDestinationCurrencyCode() (DestinationCurrencyCode string) {
	if r != nil && r.DestinationCurrencyCode != nil {
		DestinationCurrencyCode = *r.DestinationCurrencyCode
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetForeignPaymentAmount() (ForeignPaymentAmount int) {
	if r != nil && r.ForeignPaymentAmount != nil {
		ForeignPaymentAmount = *r.ForeignPaymentAmount
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetForeignTraceNumber() (ForeignTraceNumber string) {
	if r != nil && r.ForeignTraceNumber != nil {
		ForeignTraceNumber = *r.ForeignTraceNumber
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetInternationalTransactionTypeCode() (InternationalTransactionTypeCode string) {
	if r != nil && r.InternationalTransactionTypeCode != nil {
		InternationalTransactionTypeCode = *r.InternationalTransactionTypeCode
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetOriginatingCurrencyCode() (OriginatingCurrencyCode string) {
	if r != nil && r.OriginatingCurrencyCode != nil {
		OriginatingCurrencyCode = *r.OriginatingCurrencyCode
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetOriginatingDepositoryFinancialInstitutionName() (OriginatingDepositoryFinancialInstitutionName string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionName != nil {
		OriginatingDepositoryFinancialInstitutionName = *r.OriginatingDepositoryFinancialInstitutionName
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetOriginatingDepositoryFinancialInstitutionIdQualifier() (OriginatingDepositoryFinancialInstitutionIdQualifier string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionIdQualifier != nil {
		OriginatingDepositoryFinancialInstitutionIdQualifier = *r.OriginatingDepositoryFinancialInstitutionIdQualifier
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetOriginatingDepositoryFinancialInstitutionId() (OriginatingDepositoryFinancialInstitutionId string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionId != nil {
		OriginatingDepositoryFinancialInstitutionId = *r.OriginatingDepositoryFinancialInstitutionId
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetOriginatingDepositoryFinancialInstitutionBranchCountry() (OriginatingDepositoryFinancialInstitutionBranchCountry string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionBranchCountry != nil {
		OriginatingDepositoryFinancialInstitutionBranchCountry = *r.OriginatingDepositoryFinancialInstitutionBranchCountry
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetOriginatorCity() (OriginatorCity string) {
	if r != nil && r.OriginatorCity != nil {
		OriginatorCity = *r.OriginatorCity
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetOriginatorCompanyEntryDescription() (OriginatorCompanyEntryDescription string) {
	if r != nil && r.OriginatorCompanyEntryDescription != nil {
		OriginatorCompanyEntryDescription = *r.OriginatorCompanyEntryDescription
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetOriginatorCountry() (OriginatorCountry string) {
	if r != nil && r.OriginatorCountry != nil {
		OriginatorCountry = *r.OriginatorCountry
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetOriginatorIdentification() (OriginatorIdentification string) {
	if r != nil && r.OriginatorIdentification != nil {
		OriginatorIdentification = *r.OriginatorIdentification
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetOriginatorName() (OriginatorName string) {
	if r != nil && r.OriginatorName != nil {
		OriginatorName = *r.OriginatorName
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetOriginatorPostalCode() (OriginatorPostalCode string) {
	if r != nil && r.OriginatorPostalCode != nil {
		OriginatorPostalCode = *r.OriginatorPostalCode
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetOriginatorStreetAddress() (OriginatorStreetAddress string) {
	if r != nil && r.OriginatorStreetAddress != nil {
		OriginatorStreetAddress = *r.OriginatorStreetAddress
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetOriginatorStateOrProvince() (OriginatorStateOrProvince string) {
	if r != nil && r.OriginatorStateOrProvince != nil {
		OriginatorStateOrProvince = *r.OriginatorStateOrProvince
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetPaymentRelatedInformation() (PaymentRelatedInformation string) {
	if r != nil && r.PaymentRelatedInformation != nil {
		PaymentRelatedInformation = *r.PaymentRelatedInformation
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetPaymentRelatedInformation2() (PaymentRelatedInformation2 string) {
	if r != nil && r.PaymentRelatedInformation2 != nil {
		PaymentRelatedInformation2 = *r.PaymentRelatedInformation2
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetReceiverIdentificationNumber() (ReceiverIdentificationNumber string) {
	if r != nil && r.ReceiverIdentificationNumber != nil {
		ReceiverIdentificationNumber = *r.ReceiverIdentificationNumber
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetReceiverStreetAddress() (ReceiverStreetAddress string) {
	if r != nil && r.ReceiverStreetAddress != nil {
		ReceiverStreetAddress = *r.ReceiverStreetAddress
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetReceiverCity() (ReceiverCity string) {
	if r != nil && r.ReceiverCity != nil {
		ReceiverCity = *r.ReceiverCity
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetReceiverStateOrProvince() (ReceiverStateOrProvince string) {
	if r != nil && r.ReceiverStateOrProvince != nil {
		ReceiverStateOrProvince = *r.ReceiverStateOrProvince
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetReceiverCountry() (ReceiverCountry string) {
	if r != nil && r.ReceiverCountry != nil {
		ReceiverCountry = *r.ReceiverCountry
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetReceiverPostalCode() (ReceiverPostalCode string) {
	if r != nil && r.ReceiverPostalCode != nil {
		ReceiverPostalCode = *r.ReceiverPostalCode
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetReceivingCompanyOrIndividualName() (ReceivingCompanyOrIndividualName string) {
	if r != nil && r.ReceivingCompanyOrIndividualName != nil {
		ReceivingCompanyOrIndividualName = *r.ReceivingCompanyOrIndividualName
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetReceivingDepositoryFinancialInstitutionName() (ReceivingDepositoryFinancialInstitutionName string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionName != nil {
		ReceivingDepositoryFinancialInstitutionName = *r.ReceivingDepositoryFinancialInstitutionName
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetReceivingDepositoryFinancialInstitutionIdQualifier() (ReceivingDepositoryFinancialInstitutionIdQualifier string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionIdQualifier != nil {
		ReceivingDepositoryFinancialInstitutionIdQualifier = *r.ReceivingDepositoryFinancialInstitutionIdQualifier
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetReceivingDepositoryFinancialInstitutionId() (ReceivingDepositoryFinancialInstitutionId string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionId != nil {
		ReceivingDepositoryFinancialInstitutionId = *r.ReceivingDepositoryFinancialInstitutionId
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetReceivingDepositoryFinancialInstitutionCountry() (ReceivingDepositoryFinancialInstitutionCountry string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionCountry != nil {
		ReceivingDepositoryFinancialInstitutionCountry = *r.ReceivingDepositoryFinancialInstitutionCountry
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceInternationalACHDecline) GetTraceNumber() (TraceNumber string) {
	if r != nil && r.TraceNumber != nil {
		TraceNumber = *r.TraceNumber
	}
	return
}

//
type ACHTransferSimulationDeclinedTransactionSourceCardRouteDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency *ACHTransferSimulationDeclinedTransactionSourceCardRouteDeclineCurrency `json:"currency"`
	//
	MerchantAcceptorId *string `json:"merchant_acceptor_id"`
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
func (r *ACHTransferSimulationDeclinedTransactionSourceCardRouteDecline) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
// account currency.
func (r *ACHTransferSimulationDeclinedTransactionSourceCardRouteDecline) GetCurrency() (Currency ACHTransferSimulationDeclinedTransactionSourceCardRouteDeclineCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceCardRouteDecline) GetMerchantAcceptorId() (MerchantAcceptorId string) {
	if r != nil && r.MerchantAcceptorId != nil {
		MerchantAcceptorId = *r.MerchantAcceptorId
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceCardRouteDecline) GetMerchantCity() (MerchantCity string) {
	if r != nil && r.MerchantCity != nil {
		MerchantCity = *r.MerchantCity
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceCardRouteDecline) GetMerchantCountry() (MerchantCountry string) {
	if r != nil && r.MerchantCountry != nil {
		MerchantCountry = *r.MerchantCountry
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceCardRouteDecline) GetMerchantDescriptor() (MerchantDescriptor string) {
	if r != nil && r.MerchantDescriptor != nil {
		MerchantDescriptor = *r.MerchantDescriptor
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceCardRouteDecline) GetMerchantState() (MerchantState string) {
	if r != nil && r.MerchantState != nil {
		MerchantState = *r.MerchantState
	}
	return
}

func (r *ACHTransferSimulationDeclinedTransactionSourceCardRouteDecline) GetMerchantCategoryCode() (MerchantCategoryCode string) {
	if r != nil && r.MerchantCategoryCode != nil {
		MerchantCategoryCode = *r.MerchantCategoryCode
	}
	return
}

type ACHTransferSimulationDeclinedTransactionSourceCardRouteDeclineCurrency string

const (
	ACHTransferSimulationDeclinedTransactionSourceCardRouteDeclineCurrencyCad ACHTransferSimulationDeclinedTransactionSourceCardRouteDeclineCurrency = "CAD"
	ACHTransferSimulationDeclinedTransactionSourceCardRouteDeclineCurrencyChf ACHTransferSimulationDeclinedTransactionSourceCardRouteDeclineCurrency = "CHF"
	ACHTransferSimulationDeclinedTransactionSourceCardRouteDeclineCurrencyEur ACHTransferSimulationDeclinedTransactionSourceCardRouteDeclineCurrency = "EUR"
	ACHTransferSimulationDeclinedTransactionSourceCardRouteDeclineCurrencyGbp ACHTransferSimulationDeclinedTransactionSourceCardRouteDeclineCurrency = "GBP"
	ACHTransferSimulationDeclinedTransactionSourceCardRouteDeclineCurrencyJpy ACHTransferSimulationDeclinedTransactionSourceCardRouteDeclineCurrency = "JPY"
	ACHTransferSimulationDeclinedTransactionSourceCardRouteDeclineCurrencyUsd ACHTransferSimulationDeclinedTransactionSourceCardRouteDeclineCurrency = "USD"
)

type ACHTransferSimulationDeclinedTransactionType string

const (
	ACHTransferSimulationDeclinedTransactionTypeDeclinedTransaction ACHTransferSimulationDeclinedTransactionType = "declined_transaction"
)

type ACHTransferSimulationType string

const (
	ACHTransferSimulationTypeInboundACHTransferSimulationResult ACHTransferSimulationType = "inbound_ach_transfer_simulation_result"
)

type SimulateAnACHTransferToYourAccountParameters struct {
	// The identifier of the Account Number the inbound ACH Transfer is for.
	AccountNumberId *string `json:"account_number_id"`
	// The transfer amount in cents. A positive amount originates a credit transfer
	// pushing funds to the receiving account. A negative amount originates a debit
	// transfer pulling funds from the receiving account.
	Amount *int `json:"amount"`
}

// The identifier of the Account Number the inbound ACH Transfer is for.
func (r *SimulateAnACHTransferToYourAccountParameters) GetAccountNumberId() (AccountNumberId string) {
	if r != nil && r.AccountNumberId != nil {
		AccountNumberId = *r.AccountNumberId
	}
	return
}

// The transfer amount in cents. A positive amount originates a credit transfer
// pushing funds to the receiving account. A negative amount originates a debit
// transfer pulling funds from the receiving account.
func (r *SimulateAnACHTransferToYourAccountParameters) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

type SimulatesAdvancingTheStateOfACardDisputeParameters struct {
	// The status to move the dispute to.
	Status *SimulatesAdvancingTheStateOfACardDisputeParametersStatus `json:"status"`
	// Why the dispute was rejected. Not required for accepting disputes.
	Explanation *string `json:"explanation,omitempty"`
}

// The status to move the dispute to.
func (r *SimulatesAdvancingTheStateOfACardDisputeParameters) GetStatus() (Status SimulatesAdvancingTheStateOfACardDisputeParametersStatus) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

// Why the dispute was rejected. Not required for accepting disputes.
func (r *SimulatesAdvancingTheStateOfACardDisputeParameters) GetExplanation() (Explanation string) {
	if r != nil && r.Explanation != nil {
		Explanation = *r.Explanation
	}
	return
}

type SimulatesAdvancingTheStateOfACardDisputeParametersStatus string

const (
	SimulatesAdvancingTheStateOfACardDisputeParametersStatusAccepted SimulatesAdvancingTheStateOfACardDisputeParametersStatus = "accepted"
	SimulatesAdvancingTheStateOfACardDisputeParametersStatusRejected SimulatesAdvancingTheStateOfACardDisputeParametersStatus = "rejected"
)

//
type DigitalWalletTokenRequestCreateResponse struct {
	// If the simulated tokenization attempt was declined, this field contains details
	// as to why.
	DeclineReason *DigitalWalletTokenRequestCreateResponseDeclineReason `json:"decline_reason"`
	// If the simulated tokenization attempt was accepted, this field contains the id
	// of the Digital Wallet Token that was created.
	DigitalWalletTokenId *string `json:"digital_wallet_token_id"`
	// A constant representing the object's type. For this resource it will always be
	// `inbound_digital_wallet_token_request_simulation_result`.
	Type *DigitalWalletTokenRequestCreateResponseType `json:"type"`
}

// If the simulated tokenization attempt was declined, this field contains details
// as to why.
func (r *DigitalWalletTokenRequestCreateResponse) GetDeclineReason() (DeclineReason DigitalWalletTokenRequestCreateResponseDeclineReason) {
	if r != nil && r.DeclineReason != nil {
		DeclineReason = *r.DeclineReason
	}
	return
}

// If the simulated tokenization attempt was accepted, this field contains the id
// of the Digital Wallet Token that was created.
func (r *DigitalWalletTokenRequestCreateResponse) GetDigitalWalletTokenId() (DigitalWalletTokenId string) {
	if r != nil && r.DigitalWalletTokenId != nil {
		DigitalWalletTokenId = *r.DigitalWalletTokenId
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `inbound_digital_wallet_token_request_simulation_result`.
func (r *DigitalWalletTokenRequestCreateResponse) GetType() (Type DigitalWalletTokenRequestCreateResponseType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

type DigitalWalletTokenRequestCreateResponseDeclineReason string

const (
	DigitalWalletTokenRequestCreateResponseDeclineReasonCardNotActive        DigitalWalletTokenRequestCreateResponseDeclineReason = "card_not_active"
	DigitalWalletTokenRequestCreateResponseDeclineReasonNoVerificationMethod DigitalWalletTokenRequestCreateResponseDeclineReason = "no_verification_method"
	DigitalWalletTokenRequestCreateResponseDeclineReasonWebhookTimedOut      DigitalWalletTokenRequestCreateResponseDeclineReason = "webhook_timed_out"
	DigitalWalletTokenRequestCreateResponseDeclineReasonWebhookDeclined      DigitalWalletTokenRequestCreateResponseDeclineReason = "webhook_declined"
)

type DigitalWalletTokenRequestCreateResponseType string

const (
	DigitalWalletTokenRequestCreateResponseTypeInboundDigitalWalletTokenRequestSimulationResult DigitalWalletTokenRequestCreateResponseType = "inbound_digital_wallet_token_request_simulation_result"
)

type SimulateDigitalWalletActivityOnACardParameters struct {
	// The identifier of the Card to be authorized.
	CardId *string `json:"card_id"`
}

// The identifier of the Card to be authorized.
func (r *SimulateDigitalWalletActivityOnACardParameters) GetCardId() (CardId string) {
	if r != nil && r.CardId != nil {
		CardId = *r.CardId
	}
	return
}

//
type WireTransferSimulation struct {
	// If the Wire Transfer attempt succeeds, this will contain the resulting
	// [Transaction](#transactions) object. The Transaction's `source` will be of
	// `category: inbound_wire_transfer`.
	Transaction *WireTransferSimulationTransaction `json:"transaction"`
	// A constant representing the object's type. For this resource it will always be
	// `inbound_wire_transfer_simulation_result`.
	Type *WireTransferSimulationType `json:"type"`
}

// If the Wire Transfer attempt succeeds, this will contain the resulting
// [Transaction](#transactions) object. The Transaction's `source` will be of
// `category: inbound_wire_transfer`.
func (r *WireTransferSimulation) GetTransaction() (Transaction WireTransferSimulationTransaction) {
	if r != nil && r.Transaction != nil {
		Transaction = *r.Transaction
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `inbound_wire_transfer_simulation_result`.
func (r *WireTransferSimulation) GetType() (Type WireTransferSimulationType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

//
type WireTransferSimulationTransaction struct {
	// The identifier for the Account the Transaction belongs to.
	AccountId *string `json:"account_id"`
	// The Transaction amount in the minor unit of its currency. For dollars, for
	// example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// Transaction's currency. This will match the currency on the Transcation's
	// Account.
	Currency *WireTransferSimulationTransactionCurrency `json:"currency"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the
	// Transaction occured.
	CreatedAt *string `json:"created_at"`
	// For a Transaction related to a transfer, this is the description you provide.
	// For a Transaction related to a payment, this is the description the vendor
	// provides.
	Description *string `json:"description"`
	// The Transaction identifier.
	Id *string `json:"id"`
	// The identifier for the route this Transaction came through. Routes are things
	// like cards and ACH details.
	RouteId *string `json:"route_id"`
	// The type of the route this Transaction came through.
	RouteType *string `json:"route_type"`
	// This is an object giving more details on the network-level event that caused the
	// Transaction. Note that for backwards compatibility reasons, additional
	// undocumented keys may appear in this object. These should be treated as
	// deprecated and will be removed in the future.
	Source *WireTransferSimulationTransactionSource `json:"source"`
	// A constant representing the object's type. For this resource it will always be
	// `transaction`.
	Type *WireTransferSimulationTransactionType `json:"type"`
}

// The identifier for the Account the Transaction belongs to.
func (r *WireTransferSimulationTransaction) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

// The Transaction amount in the minor unit of its currency. For dollars, for
// example, this is cents.
func (r *WireTransferSimulationTransaction) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// Transaction's currency. This will match the currency on the Transcation's
// Account.
func (r *WireTransferSimulationTransaction) GetCurrency() (Currency WireTransferSimulationTransactionCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the
// Transaction occured.
func (r *WireTransferSimulationTransaction) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// For a Transaction related to a transfer, this is the description you provide.
// For a Transaction related to a payment, this is the description the vendor
// provides.
func (r *WireTransferSimulationTransaction) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The Transaction identifier.
func (r *WireTransferSimulationTransaction) GetId() (Id string) {
	if r != nil && r.Id != nil {
		Id = *r.Id
	}
	return
}

// The identifier for the route this Transaction came through. Routes are things
// like cards and ACH details.
func (r *WireTransferSimulationTransaction) GetRouteId() (RouteId string) {
	if r != nil && r.RouteId != nil {
		RouteId = *r.RouteId
	}
	return
}

// The type of the route this Transaction came through.
func (r *WireTransferSimulationTransaction) GetRouteType() (RouteType string) {
	if r != nil && r.RouteType != nil {
		RouteType = *r.RouteType
	}
	return
}

// This is an object giving more details on the network-level event that caused the
// Transaction. Note that for backwards compatibility reasons, additional
// undocumented keys may appear in this object. These should be treated as
// deprecated and will be removed in the future.
func (r *WireTransferSimulationTransaction) GetSource() (Source WireTransferSimulationTransactionSource) {
	if r != nil && r.Source != nil {
		Source = *r.Source
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `transaction`.
func (r *WireTransferSimulationTransaction) GetType() (Type WireTransferSimulationTransactionType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

type WireTransferSimulationTransactionCurrency string

const (
	WireTransferSimulationTransactionCurrencyCad WireTransferSimulationTransactionCurrency = "CAD"
	WireTransferSimulationTransactionCurrencyChf WireTransferSimulationTransactionCurrency = "CHF"
	WireTransferSimulationTransactionCurrencyEur WireTransferSimulationTransactionCurrency = "EUR"
	WireTransferSimulationTransactionCurrencyGbp WireTransferSimulationTransactionCurrency = "GBP"
	WireTransferSimulationTransactionCurrencyJpy WireTransferSimulationTransactionCurrency = "JPY"
	WireTransferSimulationTransactionCurrencyUsd WireTransferSimulationTransactionCurrency = "USD"
)

//
type WireTransferSimulationTransactionSource struct {
	// The type of transaction that took place. We may add additional possible values
	// for this enum over time; your application should be able to handle such
	// additions gracefully.
	Category *WireTransferSimulationTransactionSourceCategory `json:"category"`
	// A Account Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `account_transfer_intention`.
	AccountTransferIntention *WireTransferSimulationTransactionSourceAccountTransferIntention `json:"account_transfer_intention"`
	// A ACH Check Conversion Return object. This field will be present in the JSON
	// response if and only if `category` is equal to `ach_check_conversion_return`.
	ACHCheckConversionReturn *WireTransferSimulationTransactionSourceACHCheckConversionReturn `json:"ach_check_conversion_return"`
	// A ACH Check Conversion object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_check_conversion`.
	ACHCheckConversion *WireTransferSimulationTransactionSourceACHCheckConversion `json:"ach_check_conversion"`
	// A ACH Transfer Intention object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_transfer_intention`.
	ACHTransferIntention *WireTransferSimulationTransactionSourceACHTransferIntention `json:"ach_transfer_intention"`
	// A ACH Transfer Rejection object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_transfer_rejection`.
	ACHTransferRejection *WireTransferSimulationTransactionSourceACHTransferRejection `json:"ach_transfer_rejection"`
	// A ACH Transfer Return object. This field will be present in the JSON response if
	// and only if `category` is equal to `ach_transfer_return`.
	ACHTransferReturn *WireTransferSimulationTransactionSourceACHTransferReturn `json:"ach_transfer_return"`
	// A Card Dispute Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_dispute_acceptance`.
	CardDisputeAcceptance *WireTransferSimulationTransactionSourceCardDisputeAcceptance `json:"card_dispute_acceptance"`
	// A Card Refund object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_refund`.
	CardRefund *WireTransferSimulationTransactionSourceCardRefund `json:"card_refund"`
	// A Card Settlement object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_settlement`.
	CardSettlement *WireTransferSimulationTransactionSourceCardSettlement `json:"card_settlement"`
	// A Check Deposit Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_deposit_acceptance`.
	CheckDepositAcceptance *WireTransferSimulationTransactionSourceCheckDepositAcceptance `json:"check_deposit_acceptance"`
	// A Check Deposit Return object. This field will be present in the JSON response
	// if and only if `category` is equal to `check_deposit_return`.
	CheckDepositReturn *WireTransferSimulationTransactionSourceCheckDepositReturn `json:"check_deposit_return"`
	// A Check Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_intention`.
	CheckTransferIntention *WireTransferSimulationTransactionSourceCheckTransferIntention `json:"check_transfer_intention"`
	// A Check Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_rejection`.
	CheckTransferRejection *WireTransferSimulationTransactionSourceCheckTransferRejection `json:"check_transfer_rejection"`
	// A Check Transfer Stop Payment Request object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `check_transfer_stop_payment_request`.
	CheckTransferStopPaymentRequest *WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequest `json:"check_transfer_stop_payment_request"`
	// A Dispute Resolution object. This field will be present in the JSON response if
	// and only if `category` is equal to `dispute_resolution`.
	DisputeResolution *WireTransferSimulationTransactionSourceDisputeResolution `json:"dispute_resolution"`
	// A Empyreal Cash Deposit object. This field will be present in the JSON response
	// if and only if `category` is equal to `empyreal_cash_deposit`.
	EmpyrealCashDeposit *WireTransferSimulationTransactionSourceEmpyrealCashDeposit `json:"empyreal_cash_deposit"`
	// A Inbound ACH Transfer object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_ach_transfer`.
	InboundACHTransfer *WireTransferSimulationTransactionSourceInboundACHTransfer `json:"inbound_ach_transfer"`
	// A Inbound Check object. This field will be present in the JSON response if and
	// only if `category` is equal to `inbound_check`.
	InboundCheck *WireTransferSimulationTransactionSourceInboundCheck `json:"inbound_check"`
	// A Inbound International ACH Transfer object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `inbound_international_ach_transfer`.
	InboundInternationalACHTransfer *WireTransferSimulationTransactionSourceInboundInternationalACHTransfer `json:"inbound_international_ach_transfer"`
	// A Inbound Real Time Payments Transfer Confirmation object. This field will be
	// present in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_confirmation`.
	InboundRealTimePaymentsTransferConfirmation *WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation `json:"inbound_real_time_payments_transfer_confirmation"`
	// A Inbound Wire Drawdown Payment Reversal object. This field will be present in
	// the JSON response if and only if `category` is equal to
	// `inbound_wire_drawdown_payment_reversal`.
	InboundWireDrawdownPaymentReversal *WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal `json:"inbound_wire_drawdown_payment_reversal"`
	// A Inbound Wire Drawdown Payment object. This field will be present in the JSON
	// response if and only if `category` is equal to `inbound_wire_drawdown_payment`.
	InboundWireDrawdownPayment *WireTransferSimulationTransactionSourceInboundWireDrawdownPayment `json:"inbound_wire_drawdown_payment"`
	// A Inbound Wire Reversal object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_wire_reversal`.
	InboundWireReversal *WireTransferSimulationTransactionSourceInboundWireReversal `json:"inbound_wire_reversal"`
	// A Inbound Wire Transfer object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_wire_transfer`.
	InboundWireTransfer *WireTransferSimulationTransactionSourceInboundWireTransfer `json:"inbound_wire_transfer"`
	// A Internal Source object. This field will be present in the JSON response if and
	// only if `category` is equal to `internal_source`.
	InternalSource *WireTransferSimulationTransactionSourceInternalSource `json:"internal_source"`
	// A Deprecated Card Refund object. This field will be present in the JSON response
	// if and only if `category` is equal to `card_route_refund`.
	CardRouteRefund *WireTransferSimulationTransactionSourceCardRouteRefund `json:"card_route_refund"`
	// A Deprecated Card Settlement object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_route_settlement`.
	CardRouteSettlement *WireTransferSimulationTransactionSourceCardRouteSettlement `json:"card_route_settlement"`
	// A Sample Funds object. This field will be present in the JSON response if and
	// only if `category` is equal to `sample_funds`.
	SampleFunds *WireTransferSimulationTransactionSourceSampleFunds `json:"sample_funds"`
	// A Wire Drawdown Payment Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to
	// `wire_drawdown_payment_intention`.
	WireDrawdownPaymentIntention *WireTransferSimulationTransactionSourceWireDrawdownPaymentIntention `json:"wire_drawdown_payment_intention"`
	// A Wire Drawdown Payment Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to
	// `wire_drawdown_payment_rejection`.
	WireDrawdownPaymentRejection *WireTransferSimulationTransactionSourceWireDrawdownPaymentRejection `json:"wire_drawdown_payment_rejection"`
	// A Wire Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_intention`.
	WireTransferIntention *WireTransferSimulationTransactionSourceWireTransferIntention `json:"wire_transfer_intention"`
	// A Wire Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_rejection`.
	WireTransferRejection *WireTransferSimulationTransactionSourceWireTransferRejection `json:"wire_transfer_rejection"`
}

// The type of transaction that took place. We may add additional possible values
// for this enum over time; your application should be able to handle such
// additions gracefully.
func (r *WireTransferSimulationTransactionSource) GetCategory() (Category WireTransferSimulationTransactionSourceCategory) {
	if r != nil && r.Category != nil {
		Category = *r.Category
	}
	return
}

// A Account Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `account_transfer_intention`.
func (r *WireTransferSimulationTransactionSource) GetAccountTransferIntention() (AccountTransferIntention WireTransferSimulationTransactionSourceAccountTransferIntention) {
	if r != nil && r.AccountTransferIntention != nil {
		AccountTransferIntention = *r.AccountTransferIntention
	}
	return
}

// A ACH Check Conversion Return object. This field will be present in the JSON
// response if and only if `category` is equal to `ach_check_conversion_return`.
func (r *WireTransferSimulationTransactionSource) GetACHCheckConversionReturn() (ACHCheckConversionReturn WireTransferSimulationTransactionSourceACHCheckConversionReturn) {
	if r != nil && r.ACHCheckConversionReturn != nil {
		ACHCheckConversionReturn = *r.ACHCheckConversionReturn
	}
	return
}

// A ACH Check Conversion object. This field will be present in the JSON response
// if and only if `category` is equal to `ach_check_conversion`.
func (r *WireTransferSimulationTransactionSource) GetACHCheckConversion() (ACHCheckConversion WireTransferSimulationTransactionSourceACHCheckConversion) {
	if r != nil && r.ACHCheckConversion != nil {
		ACHCheckConversion = *r.ACHCheckConversion
	}
	return
}

// A ACH Transfer Intention object. This field will be present in the JSON response
// if and only if `category` is equal to `ach_transfer_intention`.
func (r *WireTransferSimulationTransactionSource) GetACHTransferIntention() (ACHTransferIntention WireTransferSimulationTransactionSourceACHTransferIntention) {
	if r != nil && r.ACHTransferIntention != nil {
		ACHTransferIntention = *r.ACHTransferIntention
	}
	return
}

// A ACH Transfer Rejection object. This field will be present in the JSON response
// if and only if `category` is equal to `ach_transfer_rejection`.
func (r *WireTransferSimulationTransactionSource) GetACHTransferRejection() (ACHTransferRejection WireTransferSimulationTransactionSourceACHTransferRejection) {
	if r != nil && r.ACHTransferRejection != nil {
		ACHTransferRejection = *r.ACHTransferRejection
	}
	return
}

// A ACH Transfer Return object. This field will be present in the JSON response if
// and only if `category` is equal to `ach_transfer_return`.
func (r *WireTransferSimulationTransactionSource) GetACHTransferReturn() (ACHTransferReturn WireTransferSimulationTransactionSourceACHTransferReturn) {
	if r != nil && r.ACHTransferReturn != nil {
		ACHTransferReturn = *r.ACHTransferReturn
	}
	return
}

// A Card Dispute Acceptance object. This field will be present in the JSON
// response if and only if `category` is equal to `card_dispute_acceptance`.
func (r *WireTransferSimulationTransactionSource) GetCardDisputeAcceptance() (CardDisputeAcceptance WireTransferSimulationTransactionSourceCardDisputeAcceptance) {
	if r != nil && r.CardDisputeAcceptance != nil {
		CardDisputeAcceptance = *r.CardDisputeAcceptance
	}
	return
}

// A Card Refund object. This field will be present in the JSON response if and
// only if `category` is equal to `card_refund`.
func (r *WireTransferSimulationTransactionSource) GetCardRefund() (CardRefund WireTransferSimulationTransactionSourceCardRefund) {
	if r != nil && r.CardRefund != nil {
		CardRefund = *r.CardRefund
	}
	return
}

// A Card Settlement object. This field will be present in the JSON response if and
// only if `category` is equal to `card_settlement`.
func (r *WireTransferSimulationTransactionSource) GetCardSettlement() (CardSettlement WireTransferSimulationTransactionSourceCardSettlement) {
	if r != nil && r.CardSettlement != nil {
		CardSettlement = *r.CardSettlement
	}
	return
}

// A Check Deposit Acceptance object. This field will be present in the JSON
// response if and only if `category` is equal to `check_deposit_acceptance`.
func (r *WireTransferSimulationTransactionSource) GetCheckDepositAcceptance() (CheckDepositAcceptance WireTransferSimulationTransactionSourceCheckDepositAcceptance) {
	if r != nil && r.CheckDepositAcceptance != nil {
		CheckDepositAcceptance = *r.CheckDepositAcceptance
	}
	return
}

// A Check Deposit Return object. This field will be present in the JSON response
// if and only if `category` is equal to `check_deposit_return`.
func (r *WireTransferSimulationTransactionSource) GetCheckDepositReturn() (CheckDepositReturn WireTransferSimulationTransactionSourceCheckDepositReturn) {
	if r != nil && r.CheckDepositReturn != nil {
		CheckDepositReturn = *r.CheckDepositReturn
	}
	return
}

// A Check Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `check_transfer_intention`.
func (r *WireTransferSimulationTransactionSource) GetCheckTransferIntention() (CheckTransferIntention WireTransferSimulationTransactionSourceCheckTransferIntention) {
	if r != nil && r.CheckTransferIntention != nil {
		CheckTransferIntention = *r.CheckTransferIntention
	}
	return
}

// A Check Transfer Rejection object. This field will be present in the JSON
// response if and only if `category` is equal to `check_transfer_rejection`.
func (r *WireTransferSimulationTransactionSource) GetCheckTransferRejection() (CheckTransferRejection WireTransferSimulationTransactionSourceCheckTransferRejection) {
	if r != nil && r.CheckTransferRejection != nil {
		CheckTransferRejection = *r.CheckTransferRejection
	}
	return
}

// A Check Transfer Stop Payment Request object. This field will be present in the
// JSON response if and only if `category` is equal to
// `check_transfer_stop_payment_request`.
func (r *WireTransferSimulationTransactionSource) GetCheckTransferStopPaymentRequest() (CheckTransferStopPaymentRequest WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequest) {
	if r != nil && r.CheckTransferStopPaymentRequest != nil {
		CheckTransferStopPaymentRequest = *r.CheckTransferStopPaymentRequest
	}
	return
}

// A Dispute Resolution object. This field will be present in the JSON response if
// and only if `category` is equal to `dispute_resolution`.
func (r *WireTransferSimulationTransactionSource) GetDisputeResolution() (DisputeResolution WireTransferSimulationTransactionSourceDisputeResolution) {
	if r != nil && r.DisputeResolution != nil {
		DisputeResolution = *r.DisputeResolution
	}
	return
}

// A Empyreal Cash Deposit object. This field will be present in the JSON response
// if and only if `category` is equal to `empyreal_cash_deposit`.
func (r *WireTransferSimulationTransactionSource) GetEmpyrealCashDeposit() (EmpyrealCashDeposit WireTransferSimulationTransactionSourceEmpyrealCashDeposit) {
	if r != nil && r.EmpyrealCashDeposit != nil {
		EmpyrealCashDeposit = *r.EmpyrealCashDeposit
	}
	return
}

// A Inbound ACH Transfer object. This field will be present in the JSON response
// if and only if `category` is equal to `inbound_ach_transfer`.
func (r *WireTransferSimulationTransactionSource) GetInboundACHTransfer() (InboundACHTransfer WireTransferSimulationTransactionSourceInboundACHTransfer) {
	if r != nil && r.InboundACHTransfer != nil {
		InboundACHTransfer = *r.InboundACHTransfer
	}
	return
}

// A Inbound Check object. This field will be present in the JSON response if and
// only if `category` is equal to `inbound_check`.
func (r *WireTransferSimulationTransactionSource) GetInboundCheck() (InboundCheck WireTransferSimulationTransactionSourceInboundCheck) {
	if r != nil && r.InboundCheck != nil {
		InboundCheck = *r.InboundCheck
	}
	return
}

// A Inbound International ACH Transfer object. This field will be present in the
// JSON response if and only if `category` is equal to
// `inbound_international_ach_transfer`.
func (r *WireTransferSimulationTransactionSource) GetInboundInternationalACHTransfer() (InboundInternationalACHTransfer WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) {
	if r != nil && r.InboundInternationalACHTransfer != nil {
		InboundInternationalACHTransfer = *r.InboundInternationalACHTransfer
	}
	return
}

// A Inbound Real Time Payments Transfer Confirmation object. This field will be
// present in the JSON response if and only if `category` is equal to
// `inbound_real_time_payments_transfer_confirmation`.
func (r *WireTransferSimulationTransactionSource) GetInboundRealTimePaymentsTransferConfirmation() (InboundRealTimePaymentsTransferConfirmation WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) {
	if r != nil && r.InboundRealTimePaymentsTransferConfirmation != nil {
		InboundRealTimePaymentsTransferConfirmation = *r.InboundRealTimePaymentsTransferConfirmation
	}
	return
}

// A Inbound Wire Drawdown Payment Reversal object. This field will be present in
// the JSON response if and only if `category` is equal to
// `inbound_wire_drawdown_payment_reversal`.
func (r *WireTransferSimulationTransactionSource) GetInboundWireDrawdownPaymentReversal() (InboundWireDrawdownPaymentReversal WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) {
	if r != nil && r.InboundWireDrawdownPaymentReversal != nil {
		InboundWireDrawdownPaymentReversal = *r.InboundWireDrawdownPaymentReversal
	}
	return
}

// A Inbound Wire Drawdown Payment object. This field will be present in the JSON
// response if and only if `category` is equal to `inbound_wire_drawdown_payment`.
func (r *WireTransferSimulationTransactionSource) GetInboundWireDrawdownPayment() (InboundWireDrawdownPayment WireTransferSimulationTransactionSourceInboundWireDrawdownPayment) {
	if r != nil && r.InboundWireDrawdownPayment != nil {
		InboundWireDrawdownPayment = *r.InboundWireDrawdownPayment
	}
	return
}

// A Inbound Wire Reversal object. This field will be present in the JSON response
// if and only if `category` is equal to `inbound_wire_reversal`.
func (r *WireTransferSimulationTransactionSource) GetInboundWireReversal() (InboundWireReversal WireTransferSimulationTransactionSourceInboundWireReversal) {
	if r != nil && r.InboundWireReversal != nil {
		InboundWireReversal = *r.InboundWireReversal
	}
	return
}

// A Inbound Wire Transfer object. This field will be present in the JSON response
// if and only if `category` is equal to `inbound_wire_transfer`.
func (r *WireTransferSimulationTransactionSource) GetInboundWireTransfer() (InboundWireTransfer WireTransferSimulationTransactionSourceInboundWireTransfer) {
	if r != nil && r.InboundWireTransfer != nil {
		InboundWireTransfer = *r.InboundWireTransfer
	}
	return
}

// A Internal Source object. This field will be present in the JSON response if and
// only if `category` is equal to `internal_source`.
func (r *WireTransferSimulationTransactionSource) GetInternalSource() (InternalSource WireTransferSimulationTransactionSourceInternalSource) {
	if r != nil && r.InternalSource != nil {
		InternalSource = *r.InternalSource
	}
	return
}

// A Deprecated Card Refund object. This field will be present in the JSON response
// if and only if `category` is equal to `card_route_refund`.
func (r *WireTransferSimulationTransactionSource) GetCardRouteRefund() (CardRouteRefund WireTransferSimulationTransactionSourceCardRouteRefund) {
	if r != nil && r.CardRouteRefund != nil {
		CardRouteRefund = *r.CardRouteRefund
	}
	return
}

// A Deprecated Card Settlement object. This field will be present in the JSON
// response if and only if `category` is equal to `card_route_settlement`.
func (r *WireTransferSimulationTransactionSource) GetCardRouteSettlement() (CardRouteSettlement WireTransferSimulationTransactionSourceCardRouteSettlement) {
	if r != nil && r.CardRouteSettlement != nil {
		CardRouteSettlement = *r.CardRouteSettlement
	}
	return
}

// A Sample Funds object. This field will be present in the JSON response if and
// only if `category` is equal to `sample_funds`.
func (r *WireTransferSimulationTransactionSource) GetSampleFunds() (SampleFunds WireTransferSimulationTransactionSourceSampleFunds) {
	if r != nil && r.SampleFunds != nil {
		SampleFunds = *r.SampleFunds
	}
	return
}

// A Wire Drawdown Payment Intention object. This field will be present in the JSON
// response if and only if `category` is equal to
// `wire_drawdown_payment_intention`.
func (r *WireTransferSimulationTransactionSource) GetWireDrawdownPaymentIntention() (WireDrawdownPaymentIntention WireTransferSimulationTransactionSourceWireDrawdownPaymentIntention) {
	if r != nil && r.WireDrawdownPaymentIntention != nil {
		WireDrawdownPaymentIntention = *r.WireDrawdownPaymentIntention
	}
	return
}

// A Wire Drawdown Payment Rejection object. This field will be present in the JSON
// response if and only if `category` is equal to
// `wire_drawdown_payment_rejection`.
func (r *WireTransferSimulationTransactionSource) GetWireDrawdownPaymentRejection() (WireDrawdownPaymentRejection WireTransferSimulationTransactionSourceWireDrawdownPaymentRejection) {
	if r != nil && r.WireDrawdownPaymentRejection != nil {
		WireDrawdownPaymentRejection = *r.WireDrawdownPaymentRejection
	}
	return
}

// A Wire Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `wire_transfer_intention`.
func (r *WireTransferSimulationTransactionSource) GetWireTransferIntention() (WireTransferIntention WireTransferSimulationTransactionSourceWireTransferIntention) {
	if r != nil && r.WireTransferIntention != nil {
		WireTransferIntention = *r.WireTransferIntention
	}
	return
}

// A Wire Transfer Rejection object. This field will be present in the JSON
// response if and only if `category` is equal to `wire_transfer_rejection`.
func (r *WireTransferSimulationTransactionSource) GetWireTransferRejection() (WireTransferRejection WireTransferSimulationTransactionSourceWireTransferRejection) {
	if r != nil && r.WireTransferRejection != nil {
		WireTransferRejection = *r.WireTransferRejection
	}
	return
}

type WireTransferSimulationTransactionSourceCategory string

const (
	WireTransferSimulationTransactionSourceCategoryAccountTransferIntention                    WireTransferSimulationTransactionSourceCategory = "account_transfer_intention"
	WireTransferSimulationTransactionSourceCategoryACHCheckConversionReturn                    WireTransferSimulationTransactionSourceCategory = "ach_check_conversion_return"
	WireTransferSimulationTransactionSourceCategoryACHCheckConversion                          WireTransferSimulationTransactionSourceCategory = "ach_check_conversion"
	WireTransferSimulationTransactionSourceCategoryACHTransferIntention                        WireTransferSimulationTransactionSourceCategory = "ach_transfer_intention"
	WireTransferSimulationTransactionSourceCategoryACHTransferRejection                        WireTransferSimulationTransactionSourceCategory = "ach_transfer_rejection"
	WireTransferSimulationTransactionSourceCategoryACHTransferReturn                           WireTransferSimulationTransactionSourceCategory = "ach_transfer_return"
	WireTransferSimulationTransactionSourceCategoryCardDisputeAcceptance                       WireTransferSimulationTransactionSourceCategory = "card_dispute_acceptance"
	WireTransferSimulationTransactionSourceCategoryCardRefund                                  WireTransferSimulationTransactionSourceCategory = "card_refund"
	WireTransferSimulationTransactionSourceCategoryCardSettlement                              WireTransferSimulationTransactionSourceCategory = "card_settlement"
	WireTransferSimulationTransactionSourceCategoryCheckDepositAcceptance                      WireTransferSimulationTransactionSourceCategory = "check_deposit_acceptance"
	WireTransferSimulationTransactionSourceCategoryCheckDepositReturn                          WireTransferSimulationTransactionSourceCategory = "check_deposit_return"
	WireTransferSimulationTransactionSourceCategoryCheckTransferIntention                      WireTransferSimulationTransactionSourceCategory = "check_transfer_intention"
	WireTransferSimulationTransactionSourceCategoryCheckTransferRejection                      WireTransferSimulationTransactionSourceCategory = "check_transfer_rejection"
	WireTransferSimulationTransactionSourceCategoryCheckTransferStopPaymentRequest             WireTransferSimulationTransactionSourceCategory = "check_transfer_stop_payment_request"
	WireTransferSimulationTransactionSourceCategoryDisputeResolution                           WireTransferSimulationTransactionSourceCategory = "dispute_resolution"
	WireTransferSimulationTransactionSourceCategoryEmpyrealCashDeposit                         WireTransferSimulationTransactionSourceCategory = "empyreal_cash_deposit"
	WireTransferSimulationTransactionSourceCategoryInboundACHTransfer                          WireTransferSimulationTransactionSourceCategory = "inbound_ach_transfer"
	WireTransferSimulationTransactionSourceCategoryInboundCheck                                WireTransferSimulationTransactionSourceCategory = "inbound_check"
	WireTransferSimulationTransactionSourceCategoryInboundInternationalACHTransfer             WireTransferSimulationTransactionSourceCategory = "inbound_international_ach_transfer"
	WireTransferSimulationTransactionSourceCategoryInboundRealTimePaymentsTransferConfirmation WireTransferSimulationTransactionSourceCategory = "inbound_real_time_payments_transfer_confirmation"
	WireTransferSimulationTransactionSourceCategoryInboundWireDrawdownPaymentReversal          WireTransferSimulationTransactionSourceCategory = "inbound_wire_drawdown_payment_reversal"
	WireTransferSimulationTransactionSourceCategoryInboundWireDrawdownPayment                  WireTransferSimulationTransactionSourceCategory = "inbound_wire_drawdown_payment"
	WireTransferSimulationTransactionSourceCategoryInboundWireReversal                         WireTransferSimulationTransactionSourceCategory = "inbound_wire_reversal"
	WireTransferSimulationTransactionSourceCategoryInboundWireTransfer                         WireTransferSimulationTransactionSourceCategory = "inbound_wire_transfer"
	WireTransferSimulationTransactionSourceCategoryInternalSource                              WireTransferSimulationTransactionSourceCategory = "internal_source"
	WireTransferSimulationTransactionSourceCategoryCardRouteRefund                             WireTransferSimulationTransactionSourceCategory = "card_route_refund"
	WireTransferSimulationTransactionSourceCategoryCardRouteSettlement                         WireTransferSimulationTransactionSourceCategory = "card_route_settlement"
	WireTransferSimulationTransactionSourceCategoryRealTimePaymentsTransferAcknowledgement     WireTransferSimulationTransactionSourceCategory = "real_time_payments_transfer_acknowledgement"
	WireTransferSimulationTransactionSourceCategorySampleFunds                                 WireTransferSimulationTransactionSourceCategory = "sample_funds"
	WireTransferSimulationTransactionSourceCategoryWireDrawdownPaymentIntention                WireTransferSimulationTransactionSourceCategory = "wire_drawdown_payment_intention"
	WireTransferSimulationTransactionSourceCategoryWireDrawdownPaymentRejection                WireTransferSimulationTransactionSourceCategory = "wire_drawdown_payment_rejection"
	WireTransferSimulationTransactionSourceCategoryWireTransferIntention                       WireTransferSimulationTransactionSourceCategory = "wire_transfer_intention"
	WireTransferSimulationTransactionSourceCategoryWireTransferRejection                       WireTransferSimulationTransactionSourceCategory = "wire_transfer_rejection"
	WireTransferSimulationTransactionSourceCategoryOther                                       WireTransferSimulationTransactionSourceCategory = "other"
)

//
type WireTransferSimulationTransactionSourceAccountTransferIntention struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency *WireTransferSimulationTransactionSourceAccountTransferIntentionCurrency `json:"currency"`
	// The description you chose to give the transfer.
	Description *string `json:"description"`
	// The identifier of the Account to where the Account Transfer was sent.
	DestinationAccountId *string `json:"destination_account_id"`
	// The identifier of the Account from where the Account Transfer was sent.
	SourceAccountId *string `json:"source_account_id"`
	// The identifier of the Account Transfer that led to this Pending Transaction.
	TransferId *string `json:"transfer_id"`
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *WireTransferSimulationTransactionSourceAccountTransferIntention) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
// account currency.
func (r *WireTransferSimulationTransactionSourceAccountTransferIntention) GetCurrency() (Currency WireTransferSimulationTransactionSourceAccountTransferIntentionCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The description you chose to give the transfer.
func (r *WireTransferSimulationTransactionSourceAccountTransferIntention) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The identifier of the Account to where the Account Transfer was sent.
func (r *WireTransferSimulationTransactionSourceAccountTransferIntention) GetDestinationAccountId() (DestinationAccountId string) {
	if r != nil && r.DestinationAccountId != nil {
		DestinationAccountId = *r.DestinationAccountId
	}
	return
}

// The identifier of the Account from where the Account Transfer was sent.
func (r *WireTransferSimulationTransactionSourceAccountTransferIntention) GetSourceAccountId() (SourceAccountId string) {
	if r != nil && r.SourceAccountId != nil {
		SourceAccountId = *r.SourceAccountId
	}
	return
}

// The identifier of the Account Transfer that led to this Pending Transaction.
func (r *WireTransferSimulationTransactionSourceAccountTransferIntention) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}

type WireTransferSimulationTransactionSourceAccountTransferIntentionCurrency string

const (
	WireTransferSimulationTransactionSourceAccountTransferIntentionCurrencyCad WireTransferSimulationTransactionSourceAccountTransferIntentionCurrency = "CAD"
	WireTransferSimulationTransactionSourceAccountTransferIntentionCurrencyChf WireTransferSimulationTransactionSourceAccountTransferIntentionCurrency = "CHF"
	WireTransferSimulationTransactionSourceAccountTransferIntentionCurrencyEur WireTransferSimulationTransactionSourceAccountTransferIntentionCurrency = "EUR"
	WireTransferSimulationTransactionSourceAccountTransferIntentionCurrencyGbp WireTransferSimulationTransactionSourceAccountTransferIntentionCurrency = "GBP"
	WireTransferSimulationTransactionSourceAccountTransferIntentionCurrencyJpy WireTransferSimulationTransactionSourceAccountTransferIntentionCurrency = "JPY"
	WireTransferSimulationTransactionSourceAccountTransferIntentionCurrencyUsd WireTransferSimulationTransactionSourceAccountTransferIntentionCurrency = "USD"
)

//
type WireTransferSimulationTransactionSourceACHCheckConversionReturn struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int `json:"amount"`
	// Why the transfer was returned.
	ReturnReasonCode *string `json:"return_reason_code"`
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *WireTransferSimulationTransactionSourceACHCheckConversionReturn) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// Why the transfer was returned.
func (r *WireTransferSimulationTransactionSourceACHCheckConversionReturn) GetReturnReasonCode() (ReturnReasonCode string) {
	if r != nil && r.ReturnReasonCode != nil {
		ReturnReasonCode = *r.ReturnReasonCode
	}
	return
}

//
type WireTransferSimulationTransactionSourceACHCheckConversion struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int `json:"amount"`
	// The identifier of the File containing an image of the returned check.
	FileId *string `json:"file_id"`
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *WireTransferSimulationTransactionSourceACHCheckConversion) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The identifier of the File containing an image of the returned check.
func (r *WireTransferSimulationTransactionSourceACHCheckConversion) GetFileId() (FileId string) {
	if r != nil && r.FileId != nil {
		FileId = *r.FileId
	}
	return
}

//
type WireTransferSimulationTransactionSourceACHTransferIntention struct {
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
	TransferId *string `json:"transfer_id"`
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *WireTransferSimulationTransactionSourceACHTransferIntention) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *WireTransferSimulationTransactionSourceACHTransferIntention) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

func (r *WireTransferSimulationTransactionSourceACHTransferIntention) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

func (r *WireTransferSimulationTransactionSourceACHTransferIntention) GetStatementDescriptor() (StatementDescriptor string) {
	if r != nil && r.StatementDescriptor != nil {
		StatementDescriptor = *r.StatementDescriptor
	}
	return
}

// The identifier of the ACH Transfer that led to this Transaction.
func (r *WireTransferSimulationTransactionSourceACHTransferIntention) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}

//
type WireTransferSimulationTransactionSourceACHTransferRejection struct {
	// The identifier of the ACH Transfer that led to this Transaction.
	TransferId *string `json:"transfer_id"`
}

// The identifier of the ACH Transfer that led to this Transaction.
func (r *WireTransferSimulationTransactionSourceACHTransferRejection) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}

//
type WireTransferSimulationTransactionSourceACHTransferReturn struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt *string `json:"created_at"`
	// Why the ACH Transfer was returned.
	ReturnReasonCode *WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode `json:"return_reason_code"`
	// The identifier of the ACH Transfer associated with this return.
	TransferId *string `json:"transfer_id"`
	// The identifier of the Tranasaction associated with this return.
	TransactionId *string `json:"transaction_id"`
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the transfer was created.
func (r *WireTransferSimulationTransactionSourceACHTransferReturn) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// Why the ACH Transfer was returned.
func (r *WireTransferSimulationTransactionSourceACHTransferReturn) GetReturnReasonCode() (ReturnReasonCode WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode) {
	if r != nil && r.ReturnReasonCode != nil {
		ReturnReasonCode = *r.ReturnReasonCode
	}
	return
}

// The identifier of the ACH Transfer associated with this return.
func (r *WireTransferSimulationTransactionSourceACHTransferReturn) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}

// The identifier of the Tranasaction associated with this return.
func (r *WireTransferSimulationTransactionSourceACHTransferReturn) GetTransactionId() (TransactionId string) {
	if r != nil && r.TransactionId != nil {
		TransactionId = *r.TransactionId
	}
	return
}

type WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode string

const (
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeInsufficientFund                                          WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "insufficient_fund"
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeNoAccount                                                 WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "no_account"
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeAccountClosed                                             WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "account_closed"
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeInvalidAccountNumberStructure                             WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "invalid_account_number_structure"
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeAccountFrozenEntryReturnedPerOfacInstruction              WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "account_frozen_entry_returned_per_ofac_instruction"
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeCreditEntryRefusedByReceiver                              WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "credit_entry_refused_by_receiver"
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeUnauthorizedDebitToConsumerAccountUsingCorporateSecCode   WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "unauthorized_debit_to_consumer_account_using_corporate_sec_code"
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeCorporateCustomerAdvisedNotAuthorized                     WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "corporate_customer_advised_not_authorized"
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodePaymentStopped                                            WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "payment_stopped"
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeNonTransactionAccount                                     WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "non_transaction_account"
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeUncollectedFunds                                          WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "uncollected_funds"
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeRoutingNumberCheckDigitError                              WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "routing_number_check_digit_error"
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "customer_advised_unauthorized_improper_ineligible_or_incomplete"
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeAmountFieldError                                          WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "amount_field_error"
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeAuthorizationRevokedByCustomer                            WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "authorization_revoked_by_customer"
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeInvalidACHRoutingNumber                                   WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "invalid_ach_routing_number"
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeFileRecordEditCriteria                                    WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "file_record_edit_criteria"
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeEnrInvalidIndividualName                                  WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "enr_invalid_individual_name"
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeReturnedPerOdfiRequest                                    WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "returned_per_odfi_request"
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeAddendaError                                              WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "addenda_error"
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeLimitedParticipationDfi                                   WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "limited_participation_dfi"
	WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCodeOther                                                     WireTransferSimulationTransactionSourceACHTransferReturnReturnReasonCode = "other"
)

//
type WireTransferSimulationTransactionSourceCardDisputeAcceptance struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card Dispute was accepted.
	AcceptedAt *string `json:"accepted_at"`
	// The identifier of the Card Dispute that was accepted.
	CardDisputeId *string `json:"card_dispute_id"`
	// The identifier of the Transaction that was created to return the disputed funds
	// to your account.
	TransactionId *string `json:"transaction_id"`
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the Card Dispute was accepted.
func (r *WireTransferSimulationTransactionSourceCardDisputeAcceptance) GetAcceptedAt() (AcceptedAt string) {
	if r != nil && r.AcceptedAt != nil {
		AcceptedAt = *r.AcceptedAt
	}
	return
}

// The identifier of the Card Dispute that was accepted.
func (r *WireTransferSimulationTransactionSourceCardDisputeAcceptance) GetCardDisputeId() (CardDisputeId string) {
	if r != nil && r.CardDisputeId != nil {
		CardDisputeId = *r.CardDisputeId
	}
	return
}

// The identifier of the Transaction that was created to return the disputed funds
// to your account.
func (r *WireTransferSimulationTransactionSourceCardDisputeAcceptance) GetTransactionId() (TransactionId string) {
	if r != nil && r.TransactionId != nil {
		TransactionId = *r.TransactionId
	}
	return
}

//
type WireTransferSimulationTransactionSourceCardRefund struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *WireTransferSimulationTransactionSourceCardRefundCurrency `json:"currency"`
	// The identifier for the Transaction this refunds, if any.
	CardSettlementTransactionId *string `json:"card_settlement_transaction_id"`
	// A constant representing the object's type. For this resource it will always be
	// `card_refund`.
	Type *WireTransferSimulationTransactionSourceCardRefundType `json:"type"`
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *WireTransferSimulationTransactionSourceCardRefund) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *WireTransferSimulationTransactionSourceCardRefund) GetCurrency() (Currency WireTransferSimulationTransactionSourceCardRefundCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The identifier for the Transaction this refunds, if any.
func (r *WireTransferSimulationTransactionSourceCardRefund) GetCardSettlementTransactionId() (CardSettlementTransactionId string) {
	if r != nil && r.CardSettlementTransactionId != nil {
		CardSettlementTransactionId = *r.CardSettlementTransactionId
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `card_refund`.
func (r *WireTransferSimulationTransactionSourceCardRefund) GetType() (Type WireTransferSimulationTransactionSourceCardRefundType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

type WireTransferSimulationTransactionSourceCardRefundCurrency string

const (
	WireTransferSimulationTransactionSourceCardRefundCurrencyCad WireTransferSimulationTransactionSourceCardRefundCurrency = "CAD"
	WireTransferSimulationTransactionSourceCardRefundCurrencyChf WireTransferSimulationTransactionSourceCardRefundCurrency = "CHF"
	WireTransferSimulationTransactionSourceCardRefundCurrencyEur WireTransferSimulationTransactionSourceCardRefundCurrency = "EUR"
	WireTransferSimulationTransactionSourceCardRefundCurrencyGbp WireTransferSimulationTransactionSourceCardRefundCurrency = "GBP"
	WireTransferSimulationTransactionSourceCardRefundCurrencyJpy WireTransferSimulationTransactionSourceCardRefundCurrency = "JPY"
	WireTransferSimulationTransactionSourceCardRefundCurrencyUsd WireTransferSimulationTransactionSourceCardRefundCurrency = "USD"
)

type WireTransferSimulationTransactionSourceCardRefundType string

const (
	WireTransferSimulationTransactionSourceCardRefundTypeCardRefund WireTransferSimulationTransactionSourceCardRefundType = "card_refund"
)

//
type WireTransferSimulationTransactionSourceCardSettlement struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *WireTransferSimulationTransactionSourceCardSettlementCurrency `json:"currency"`
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
	PendingTransactionId *string `json:"pending_transaction_id"`
	// A constant representing the object's type. For this resource it will always be
	// `card_settlement`.
	Type *WireTransferSimulationTransactionSourceCardSettlementType `json:"type"`
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *WireTransferSimulationTransactionSourceCardSettlement) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *WireTransferSimulationTransactionSourceCardSettlement) GetCurrency() (Currency WireTransferSimulationTransactionSourceCardSettlementCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

func (r *WireTransferSimulationTransactionSourceCardSettlement) GetMerchantCity() (MerchantCity string) {
	if r != nil && r.MerchantCity != nil {
		MerchantCity = *r.MerchantCity
	}
	return
}

func (r *WireTransferSimulationTransactionSourceCardSettlement) GetMerchantCountry() (MerchantCountry string) {
	if r != nil && r.MerchantCountry != nil {
		MerchantCountry = *r.MerchantCountry
	}
	return
}

func (r *WireTransferSimulationTransactionSourceCardSettlement) GetMerchantName() (MerchantName string) {
	if r != nil && r.MerchantName != nil {
		MerchantName = *r.MerchantName
	}
	return
}

func (r *WireTransferSimulationTransactionSourceCardSettlement) GetMerchantCategoryCode() (MerchantCategoryCode string) {
	if r != nil && r.MerchantCategoryCode != nil {
		MerchantCategoryCode = *r.MerchantCategoryCode
	}
	return
}

func (r *WireTransferSimulationTransactionSourceCardSettlement) GetMerchantState() (MerchantState string) {
	if r != nil && r.MerchantState != nil {
		MerchantState = *r.MerchantState
	}
	return
}

// The identifier of the Pending Transaction associated with this Transaction.
func (r *WireTransferSimulationTransactionSourceCardSettlement) GetPendingTransactionId() (PendingTransactionId string) {
	if r != nil && r.PendingTransactionId != nil {
		PendingTransactionId = *r.PendingTransactionId
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `card_settlement`.
func (r *WireTransferSimulationTransactionSourceCardSettlement) GetType() (Type WireTransferSimulationTransactionSourceCardSettlementType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

type WireTransferSimulationTransactionSourceCardSettlementCurrency string

const (
	WireTransferSimulationTransactionSourceCardSettlementCurrencyCad WireTransferSimulationTransactionSourceCardSettlementCurrency = "CAD"
	WireTransferSimulationTransactionSourceCardSettlementCurrencyChf WireTransferSimulationTransactionSourceCardSettlementCurrency = "CHF"
	WireTransferSimulationTransactionSourceCardSettlementCurrencyEur WireTransferSimulationTransactionSourceCardSettlementCurrency = "EUR"
	WireTransferSimulationTransactionSourceCardSettlementCurrencyGbp WireTransferSimulationTransactionSourceCardSettlementCurrency = "GBP"
	WireTransferSimulationTransactionSourceCardSettlementCurrencyJpy WireTransferSimulationTransactionSourceCardSettlementCurrency = "JPY"
	WireTransferSimulationTransactionSourceCardSettlementCurrencyUsd WireTransferSimulationTransactionSourceCardSettlementCurrency = "USD"
)

type WireTransferSimulationTransactionSourceCardSettlementType string

const (
	WireTransferSimulationTransactionSourceCardSettlementTypeCardSettlement WireTransferSimulationTransactionSourceCardSettlementType = "card_settlement"
)

//
type WireTransferSimulationTransactionSourceCheckDepositAcceptance struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *WireTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency `json:"currency"`
	// The ID of the Check Deposit that led to the Transaction.
	CheckDepositId *string `json:"check_deposit_id"`
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *WireTransferSimulationTransactionSourceCheckDepositAcceptance) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *WireTransferSimulationTransactionSourceCheckDepositAcceptance) GetCurrency() (Currency WireTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The ID of the Check Deposit that led to the Transaction.
func (r *WireTransferSimulationTransactionSourceCheckDepositAcceptance) GetCheckDepositId() (CheckDepositId string) {
	if r != nil && r.CheckDepositId != nil {
		CheckDepositId = *r.CheckDepositId
	}
	return
}

type WireTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency string

const (
	WireTransferSimulationTransactionSourceCheckDepositAcceptanceCurrencyCad WireTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency = "CAD"
	WireTransferSimulationTransactionSourceCheckDepositAcceptanceCurrencyChf WireTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency = "CHF"
	WireTransferSimulationTransactionSourceCheckDepositAcceptanceCurrencyEur WireTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency = "EUR"
	WireTransferSimulationTransactionSourceCheckDepositAcceptanceCurrencyGbp WireTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency = "GBP"
	WireTransferSimulationTransactionSourceCheckDepositAcceptanceCurrencyJpy WireTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency = "JPY"
	WireTransferSimulationTransactionSourceCheckDepositAcceptanceCurrencyUsd WireTransferSimulationTransactionSourceCheckDepositAcceptanceCurrency = "USD"
)

//
type WireTransferSimulationTransactionSourceCheckDepositReturn struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check deposit was returned.
	ReturnedAt *string `json:"returned_at"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *WireTransferSimulationTransactionSourceCheckDepositReturnCurrency `json:"currency"`
	// The identifier of the Check Deposit that was returned.
	CheckDepositId *string `json:"check_deposit_id"`
	// The identifier of the transaction that reversed the original check deposit
	// transaction.
	TransactionId *string `json:"transaction_id"`
	//
	ReturnReason *WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason `json:"return_reason"`
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *WireTransferSimulationTransactionSourceCheckDepositReturn) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the check deposit was returned.
func (r *WireTransferSimulationTransactionSourceCheckDepositReturn) GetReturnedAt() (ReturnedAt string) {
	if r != nil && r.ReturnedAt != nil {
		ReturnedAt = *r.ReturnedAt
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *WireTransferSimulationTransactionSourceCheckDepositReturn) GetCurrency() (Currency WireTransferSimulationTransactionSourceCheckDepositReturnCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The identifier of the Check Deposit that was returned.
func (r *WireTransferSimulationTransactionSourceCheckDepositReturn) GetCheckDepositId() (CheckDepositId string) {
	if r != nil && r.CheckDepositId != nil {
		CheckDepositId = *r.CheckDepositId
	}
	return
}

// The identifier of the transaction that reversed the original check deposit
// transaction.
func (r *WireTransferSimulationTransactionSourceCheckDepositReturn) GetTransactionId() (TransactionId string) {
	if r != nil && r.TransactionId != nil {
		TransactionId = *r.TransactionId
	}
	return
}

func (r *WireTransferSimulationTransactionSourceCheckDepositReturn) GetReturnReason() (ReturnReason WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason) {
	if r != nil && r.ReturnReason != nil {
		ReturnReason = *r.ReturnReason
	}
	return
}

type WireTransferSimulationTransactionSourceCheckDepositReturnCurrency string

const (
	WireTransferSimulationTransactionSourceCheckDepositReturnCurrencyCad WireTransferSimulationTransactionSourceCheckDepositReturnCurrency = "CAD"
	WireTransferSimulationTransactionSourceCheckDepositReturnCurrencyChf WireTransferSimulationTransactionSourceCheckDepositReturnCurrency = "CHF"
	WireTransferSimulationTransactionSourceCheckDepositReturnCurrencyEur WireTransferSimulationTransactionSourceCheckDepositReturnCurrency = "EUR"
	WireTransferSimulationTransactionSourceCheckDepositReturnCurrencyGbp WireTransferSimulationTransactionSourceCheckDepositReturnCurrency = "GBP"
	WireTransferSimulationTransactionSourceCheckDepositReturnCurrencyJpy WireTransferSimulationTransactionSourceCheckDepositReturnCurrency = "JPY"
	WireTransferSimulationTransactionSourceCheckDepositReturnCurrencyUsd WireTransferSimulationTransactionSourceCheckDepositReturnCurrency = "USD"
)

type WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason string

const (
	WireTransferSimulationTransactionSourceCheckDepositReturnReturnReasonACHConversionNotSupported WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "ach_conversion_not_supported"
	WireTransferSimulationTransactionSourceCheckDepositReturnReturnReasonDuplicateSubmission       WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "duplicate_submission"
	WireTransferSimulationTransactionSourceCheckDepositReturnReturnReasonInsufficientFunds         WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "insufficient_funds"
	WireTransferSimulationTransactionSourceCheckDepositReturnReturnReasonNoAccount                 WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "no_account"
	WireTransferSimulationTransactionSourceCheckDepositReturnReturnReasonNotAuthorized             WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "not_authorized"
	WireTransferSimulationTransactionSourceCheckDepositReturnReturnReasonStaleDated                WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "stale_dated"
	WireTransferSimulationTransactionSourceCheckDepositReturnReturnReasonStopPayment               WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "stop_payment"
	WireTransferSimulationTransactionSourceCheckDepositReturnReturnReasonUnknownReason             WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "unknown_reason"
	WireTransferSimulationTransactionSourceCheckDepositReturnReturnReasonUnmatchedDetails          WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "unmatched_details"
	WireTransferSimulationTransactionSourceCheckDepositReturnReturnReasonUnreadableImage           WireTransferSimulationTransactionSourceCheckDepositReturnReturnReason = "unreadable_image"
)

//
type WireTransferSimulationTransactionSourceCheckTransferIntention struct {
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
	Currency *WireTransferSimulationTransactionSourceCheckTransferIntentionCurrency `json:"currency"`
	// The name that will be printed on the check.
	RecipientName *string `json:"recipient_name"`
	// The identifier of the Check Transfer with which this is associated.
	TransferId *string `json:"transfer_id"`
}

// The street address of the check's destination.
func (r *WireTransferSimulationTransactionSourceCheckTransferIntention) GetAddressLine1() (AddressLine1 string) {
	if r != nil && r.AddressLine1 != nil {
		AddressLine1 = *r.AddressLine1
	}
	return
}

// The second line of the address of the check's destination.
func (r *WireTransferSimulationTransactionSourceCheckTransferIntention) GetAddressLine2() (AddressLine2 string) {
	if r != nil && r.AddressLine2 != nil {
		AddressLine2 = *r.AddressLine2
	}
	return
}

// The city of the check's destination.
func (r *WireTransferSimulationTransactionSourceCheckTransferIntention) GetAddressCity() (AddressCity string) {
	if r != nil && r.AddressCity != nil {
		AddressCity = *r.AddressCity
	}
	return
}

// The state of the check's destination.
func (r *WireTransferSimulationTransactionSourceCheckTransferIntention) GetAddressState() (AddressState string) {
	if r != nil && r.AddressState != nil {
		AddressState = *r.AddressState
	}
	return
}

// The postal code of the check's destination.
func (r *WireTransferSimulationTransactionSourceCheckTransferIntention) GetAddressZip() (AddressZip string) {
	if r != nil && r.AddressZip != nil {
		AddressZip = *r.AddressZip
	}
	return
}

// The transfer amount in USD cents.
func (r *WireTransferSimulationTransactionSourceCheckTransferIntention) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
// currency.
func (r *WireTransferSimulationTransactionSourceCheckTransferIntention) GetCurrency() (Currency WireTransferSimulationTransactionSourceCheckTransferIntentionCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The name that will be printed on the check.
func (r *WireTransferSimulationTransactionSourceCheckTransferIntention) GetRecipientName() (RecipientName string) {
	if r != nil && r.RecipientName != nil {
		RecipientName = *r.RecipientName
	}
	return
}

// The identifier of the Check Transfer with which this is associated.
func (r *WireTransferSimulationTransactionSourceCheckTransferIntention) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}

type WireTransferSimulationTransactionSourceCheckTransferIntentionCurrency string

const (
	WireTransferSimulationTransactionSourceCheckTransferIntentionCurrencyCad WireTransferSimulationTransactionSourceCheckTransferIntentionCurrency = "CAD"
	WireTransferSimulationTransactionSourceCheckTransferIntentionCurrencyChf WireTransferSimulationTransactionSourceCheckTransferIntentionCurrency = "CHF"
	WireTransferSimulationTransactionSourceCheckTransferIntentionCurrencyEur WireTransferSimulationTransactionSourceCheckTransferIntentionCurrency = "EUR"
	WireTransferSimulationTransactionSourceCheckTransferIntentionCurrencyGbp WireTransferSimulationTransactionSourceCheckTransferIntentionCurrency = "GBP"
	WireTransferSimulationTransactionSourceCheckTransferIntentionCurrencyJpy WireTransferSimulationTransactionSourceCheckTransferIntentionCurrency = "JPY"
	WireTransferSimulationTransactionSourceCheckTransferIntentionCurrencyUsd WireTransferSimulationTransactionSourceCheckTransferIntentionCurrency = "USD"
)

//
type WireTransferSimulationTransactionSourceCheckTransferRejection struct {
	// The identifier of the Check Transfer that led to this Transaction.
	TransferId *string `json:"transfer_id"`
}

// The identifier of the Check Transfer that led to this Transaction.
func (r *WireTransferSimulationTransactionSourceCheckTransferRejection) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}

//
type WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequest struct {
	// The ID of the check transfer that was stopped.
	TransferId *string `json:"transfer_id"`
	// The transaction ID of the corresponding credit transaction.
	TransactionId *string `json:"transaction_id"`
	// The time the stop-payment was requested.
	RequestedAt *string `json:"requested_at"`
	// A constant representing the object's type. For this resource it will always be
	// `check_transfer_stop_payment_request`.
	Type *WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequestType `json:"type"`
}

// The ID of the check transfer that was stopped.
func (r *WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequest) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}

// The transaction ID of the corresponding credit transaction.
func (r *WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequest) GetTransactionId() (TransactionId string) {
	if r != nil && r.TransactionId != nil {
		TransactionId = *r.TransactionId
	}
	return
}

// The time the stop-payment was requested.
func (r *WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequest) GetRequestedAt() (RequestedAt string) {
	if r != nil && r.RequestedAt != nil {
		RequestedAt = *r.RequestedAt
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `check_transfer_stop_payment_request`.
func (r *WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequest) GetType() (Type WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequestType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

type WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequestType string

const (
	WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequestTypeCheckTransferStopPaymentRequest WireTransferSimulationTransactionSourceCheckTransferStopPaymentRequestType = "check_transfer_stop_payment_request"
)

//
type WireTransferSimulationTransactionSourceDisputeResolution struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *WireTransferSimulationTransactionSourceDisputeResolutionCurrency `json:"currency"`
	// The identifier of the Transaction that was disputed.
	DisputedTransactionId *string `json:"disputed_transaction_id"`
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *WireTransferSimulationTransactionSourceDisputeResolution) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *WireTransferSimulationTransactionSourceDisputeResolution) GetCurrency() (Currency WireTransferSimulationTransactionSourceDisputeResolutionCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The identifier of the Transaction that was disputed.
func (r *WireTransferSimulationTransactionSourceDisputeResolution) GetDisputedTransactionId() (DisputedTransactionId string) {
	if r != nil && r.DisputedTransactionId != nil {
		DisputedTransactionId = *r.DisputedTransactionId
	}
	return
}

type WireTransferSimulationTransactionSourceDisputeResolutionCurrency string

const (
	WireTransferSimulationTransactionSourceDisputeResolutionCurrencyCad WireTransferSimulationTransactionSourceDisputeResolutionCurrency = "CAD"
	WireTransferSimulationTransactionSourceDisputeResolutionCurrencyChf WireTransferSimulationTransactionSourceDisputeResolutionCurrency = "CHF"
	WireTransferSimulationTransactionSourceDisputeResolutionCurrencyEur WireTransferSimulationTransactionSourceDisputeResolutionCurrency = "EUR"
	WireTransferSimulationTransactionSourceDisputeResolutionCurrencyGbp WireTransferSimulationTransactionSourceDisputeResolutionCurrency = "GBP"
	WireTransferSimulationTransactionSourceDisputeResolutionCurrencyJpy WireTransferSimulationTransactionSourceDisputeResolutionCurrency = "JPY"
	WireTransferSimulationTransactionSourceDisputeResolutionCurrencyUsd WireTransferSimulationTransactionSourceDisputeResolutionCurrency = "USD"
)

//
type WireTransferSimulationTransactionSourceEmpyrealCashDeposit struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int `json:"amount"`
	//
	BagId *string `json:"bag_id"`
	//
	DepositDate *string `json:"deposit_date"`
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *WireTransferSimulationTransactionSourceEmpyrealCashDeposit) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *WireTransferSimulationTransactionSourceEmpyrealCashDeposit) GetBagId() (BagId string) {
	if r != nil && r.BagId != nil {
		BagId = *r.BagId
	}
	return
}

func (r *WireTransferSimulationTransactionSourceEmpyrealCashDeposit) GetDepositDate() (DepositDate string) {
	if r != nil && r.DepositDate != nil {
		DepositDate = *r.DepositDate
	}
	return
}

//
type WireTransferSimulationTransactionSourceInboundACHTransfer struct {
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
	OriginatorCompanyId *string `json:"originator_company_id"`
	//
	ReceiverIdNumber *string `json:"receiver_id_number"`
	//
	ReceiverName *string `json:"receiver_name"`
	//
	TraceNumber *string `json:"trace_number"`
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *WireTransferSimulationTransactionSourceInboundACHTransfer) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundACHTransfer) GetOriginatorCompanyName() (OriginatorCompanyName string) {
	if r != nil && r.OriginatorCompanyName != nil {
		OriginatorCompanyName = *r.OriginatorCompanyName
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundACHTransfer) GetOriginatorCompanyDescriptiveDate() (OriginatorCompanyDescriptiveDate string) {
	if r != nil && r.OriginatorCompanyDescriptiveDate != nil {
		OriginatorCompanyDescriptiveDate = *r.OriginatorCompanyDescriptiveDate
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundACHTransfer) GetOriginatorCompanyDiscretionaryData() (OriginatorCompanyDiscretionaryData string) {
	if r != nil && r.OriginatorCompanyDiscretionaryData != nil {
		OriginatorCompanyDiscretionaryData = *r.OriginatorCompanyDiscretionaryData
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundACHTransfer) GetOriginatorCompanyEntryDescription() (OriginatorCompanyEntryDescription string) {
	if r != nil && r.OriginatorCompanyEntryDescription != nil {
		OriginatorCompanyEntryDescription = *r.OriginatorCompanyEntryDescription
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundACHTransfer) GetOriginatorCompanyId() (OriginatorCompanyId string) {
	if r != nil && r.OriginatorCompanyId != nil {
		OriginatorCompanyId = *r.OriginatorCompanyId
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundACHTransfer) GetReceiverIdNumber() (ReceiverIdNumber string) {
	if r != nil && r.ReceiverIdNumber != nil {
		ReceiverIdNumber = *r.ReceiverIdNumber
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundACHTransfer) GetReceiverName() (ReceiverName string) {
	if r != nil && r.ReceiverName != nil {
		ReceiverName = *r.ReceiverName
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundACHTransfer) GetTraceNumber() (TraceNumber string) {
	if r != nil && r.TraceNumber != nil {
		TraceNumber = *r.TraceNumber
	}
	return
}

//
type WireTransferSimulationTransactionSourceInboundCheck struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *WireTransferSimulationTransactionSourceInboundCheckCurrency `json:"currency"`
	//
	CheckNumber *string `json:"check_number"`
	//
	CheckFrontImageFileId *string `json:"check_front_image_file_id"`
	//
	CheckRearImageFileId *string `json:"check_rear_image_file_id"`
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *WireTransferSimulationTransactionSourceInboundCheck) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *WireTransferSimulationTransactionSourceInboundCheck) GetCurrency() (Currency WireTransferSimulationTransactionSourceInboundCheckCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundCheck) GetCheckNumber() (CheckNumber string) {
	if r != nil && r.CheckNumber != nil {
		CheckNumber = *r.CheckNumber
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundCheck) GetCheckFrontImageFileId() (CheckFrontImageFileId string) {
	if r != nil && r.CheckFrontImageFileId != nil {
		CheckFrontImageFileId = *r.CheckFrontImageFileId
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundCheck) GetCheckRearImageFileId() (CheckRearImageFileId string) {
	if r != nil && r.CheckRearImageFileId != nil {
		CheckRearImageFileId = *r.CheckRearImageFileId
	}
	return
}

type WireTransferSimulationTransactionSourceInboundCheckCurrency string

const (
	WireTransferSimulationTransactionSourceInboundCheckCurrencyCad WireTransferSimulationTransactionSourceInboundCheckCurrency = "CAD"
	WireTransferSimulationTransactionSourceInboundCheckCurrencyChf WireTransferSimulationTransactionSourceInboundCheckCurrency = "CHF"
	WireTransferSimulationTransactionSourceInboundCheckCurrencyEur WireTransferSimulationTransactionSourceInboundCheckCurrency = "EUR"
	WireTransferSimulationTransactionSourceInboundCheckCurrencyGbp WireTransferSimulationTransactionSourceInboundCheckCurrency = "GBP"
	WireTransferSimulationTransactionSourceInboundCheckCurrencyJpy WireTransferSimulationTransactionSourceInboundCheckCurrency = "JPY"
	WireTransferSimulationTransactionSourceInboundCheckCurrencyUsd WireTransferSimulationTransactionSourceInboundCheckCurrency = "USD"
)

//
type WireTransferSimulationTransactionSourceInboundInternationalACHTransfer struct {
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
	OriginatingDepositoryFinancialInstitutionIdQualifier *string `json:"originating_depository_financial_institution_id_qualifier"`
	//
	OriginatingDepositoryFinancialInstitutionId *string `json:"originating_depository_financial_institution_id"`
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
	ReceivingDepositoryFinancialInstitutionIdQualifier *string `json:"receiving_depository_financial_institution_id_qualifier"`
	//
	ReceivingDepositoryFinancialInstitutionId *string `json:"receiving_depository_financial_institution_id"`
	//
	ReceivingDepositoryFinancialInstitutionCountry *string `json:"receiving_depository_financial_institution_country"`
	//
	TraceNumber *string `json:"trace_number"`
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetForeignExchangeIndicator() (ForeignExchangeIndicator string) {
	if r != nil && r.ForeignExchangeIndicator != nil {
		ForeignExchangeIndicator = *r.ForeignExchangeIndicator
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetForeignExchangeReferenceIndicator() (ForeignExchangeReferenceIndicator string) {
	if r != nil && r.ForeignExchangeReferenceIndicator != nil {
		ForeignExchangeReferenceIndicator = *r.ForeignExchangeReferenceIndicator
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetForeignExchangeReference() (ForeignExchangeReference string) {
	if r != nil && r.ForeignExchangeReference != nil {
		ForeignExchangeReference = *r.ForeignExchangeReference
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetDestinationCountryCode() (DestinationCountryCode string) {
	if r != nil && r.DestinationCountryCode != nil {
		DestinationCountryCode = *r.DestinationCountryCode
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetDestinationCurrencyCode() (DestinationCurrencyCode string) {
	if r != nil && r.DestinationCurrencyCode != nil {
		DestinationCurrencyCode = *r.DestinationCurrencyCode
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetForeignPaymentAmount() (ForeignPaymentAmount int) {
	if r != nil && r.ForeignPaymentAmount != nil {
		ForeignPaymentAmount = *r.ForeignPaymentAmount
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetForeignTraceNumber() (ForeignTraceNumber string) {
	if r != nil && r.ForeignTraceNumber != nil {
		ForeignTraceNumber = *r.ForeignTraceNumber
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetInternationalTransactionTypeCode() (InternationalTransactionTypeCode string) {
	if r != nil && r.InternationalTransactionTypeCode != nil {
		InternationalTransactionTypeCode = *r.InternationalTransactionTypeCode
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatingCurrencyCode() (OriginatingCurrencyCode string) {
	if r != nil && r.OriginatingCurrencyCode != nil {
		OriginatingCurrencyCode = *r.OriginatingCurrencyCode
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatingDepositoryFinancialInstitutionName() (OriginatingDepositoryFinancialInstitutionName string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionName != nil {
		OriginatingDepositoryFinancialInstitutionName = *r.OriginatingDepositoryFinancialInstitutionName
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatingDepositoryFinancialInstitutionIdQualifier() (OriginatingDepositoryFinancialInstitutionIdQualifier string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionIdQualifier != nil {
		OriginatingDepositoryFinancialInstitutionIdQualifier = *r.OriginatingDepositoryFinancialInstitutionIdQualifier
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatingDepositoryFinancialInstitutionId() (OriginatingDepositoryFinancialInstitutionId string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionId != nil {
		OriginatingDepositoryFinancialInstitutionId = *r.OriginatingDepositoryFinancialInstitutionId
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatingDepositoryFinancialInstitutionBranchCountry() (OriginatingDepositoryFinancialInstitutionBranchCountry string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionBranchCountry != nil {
		OriginatingDepositoryFinancialInstitutionBranchCountry = *r.OriginatingDepositoryFinancialInstitutionBranchCountry
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatorCity() (OriginatorCity string) {
	if r != nil && r.OriginatorCity != nil {
		OriginatorCity = *r.OriginatorCity
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatorCompanyEntryDescription() (OriginatorCompanyEntryDescription string) {
	if r != nil && r.OriginatorCompanyEntryDescription != nil {
		OriginatorCompanyEntryDescription = *r.OriginatorCompanyEntryDescription
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatorCountry() (OriginatorCountry string) {
	if r != nil && r.OriginatorCountry != nil {
		OriginatorCountry = *r.OriginatorCountry
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatorIdentification() (OriginatorIdentification string) {
	if r != nil && r.OriginatorIdentification != nil {
		OriginatorIdentification = *r.OriginatorIdentification
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatorName() (OriginatorName string) {
	if r != nil && r.OriginatorName != nil {
		OriginatorName = *r.OriginatorName
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatorPostalCode() (OriginatorPostalCode string) {
	if r != nil && r.OriginatorPostalCode != nil {
		OriginatorPostalCode = *r.OriginatorPostalCode
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatorStreetAddress() (OriginatorStreetAddress string) {
	if r != nil && r.OriginatorStreetAddress != nil {
		OriginatorStreetAddress = *r.OriginatorStreetAddress
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetOriginatorStateOrProvince() (OriginatorStateOrProvince string) {
	if r != nil && r.OriginatorStateOrProvince != nil {
		OriginatorStateOrProvince = *r.OriginatorStateOrProvince
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetPaymentRelatedInformation() (PaymentRelatedInformation string) {
	if r != nil && r.PaymentRelatedInformation != nil {
		PaymentRelatedInformation = *r.PaymentRelatedInformation
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetPaymentRelatedInformation2() (PaymentRelatedInformation2 string) {
	if r != nil && r.PaymentRelatedInformation2 != nil {
		PaymentRelatedInformation2 = *r.PaymentRelatedInformation2
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetReceiverIdentificationNumber() (ReceiverIdentificationNumber string) {
	if r != nil && r.ReceiverIdentificationNumber != nil {
		ReceiverIdentificationNumber = *r.ReceiverIdentificationNumber
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetReceiverStreetAddress() (ReceiverStreetAddress string) {
	if r != nil && r.ReceiverStreetAddress != nil {
		ReceiverStreetAddress = *r.ReceiverStreetAddress
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetReceiverCity() (ReceiverCity string) {
	if r != nil && r.ReceiverCity != nil {
		ReceiverCity = *r.ReceiverCity
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetReceiverStateOrProvince() (ReceiverStateOrProvince string) {
	if r != nil && r.ReceiverStateOrProvince != nil {
		ReceiverStateOrProvince = *r.ReceiverStateOrProvince
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetReceiverCountry() (ReceiverCountry string) {
	if r != nil && r.ReceiverCountry != nil {
		ReceiverCountry = *r.ReceiverCountry
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetReceiverPostalCode() (ReceiverPostalCode string) {
	if r != nil && r.ReceiverPostalCode != nil {
		ReceiverPostalCode = *r.ReceiverPostalCode
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetReceivingCompanyOrIndividualName() (ReceivingCompanyOrIndividualName string) {
	if r != nil && r.ReceivingCompanyOrIndividualName != nil {
		ReceivingCompanyOrIndividualName = *r.ReceivingCompanyOrIndividualName
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetReceivingDepositoryFinancialInstitutionName() (ReceivingDepositoryFinancialInstitutionName string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionName != nil {
		ReceivingDepositoryFinancialInstitutionName = *r.ReceivingDepositoryFinancialInstitutionName
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetReceivingDepositoryFinancialInstitutionIdQualifier() (ReceivingDepositoryFinancialInstitutionIdQualifier string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionIdQualifier != nil {
		ReceivingDepositoryFinancialInstitutionIdQualifier = *r.ReceivingDepositoryFinancialInstitutionIdQualifier
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetReceivingDepositoryFinancialInstitutionId() (ReceivingDepositoryFinancialInstitutionId string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionId != nil {
		ReceivingDepositoryFinancialInstitutionId = *r.ReceivingDepositoryFinancialInstitutionId
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetReceivingDepositoryFinancialInstitutionCountry() (ReceivingDepositoryFinancialInstitutionCountry string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionCountry != nil {
		ReceivingDepositoryFinancialInstitutionCountry = *r.ReceivingDepositoryFinancialInstitutionCountry
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundInternationalACHTransfer) GetTraceNumber() (TraceNumber string) {
	if r != nil && r.TraceNumber != nil {
		TraceNumber = *r.TraceNumber
	}
	return
}

//
type WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation struct {
	// The amount in the minor unit of the transfer's currency. For dollars, for
	// example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the transfer's
	// currency. This will always be "USD" for a Real Time Payments transfer.
	Currency *WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency `json:"currency"`
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
func (r *WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the transfer's
// currency. This will always be "USD" for a Real Time Payments transfer.
func (r *WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) GetCurrency() (Currency WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The name the sender of the transfer specified as the recipient of the transfer.
func (r *WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) GetCreditorName() (CreditorName string) {
	if r != nil && r.CreditorName != nil {
		CreditorName = *r.CreditorName
	}
	return
}

// The name provided by the sender of the transfer.
func (r *WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) GetDebtorName() (DebtorName string) {
	if r != nil && r.DebtorName != nil {
		DebtorName = *r.DebtorName
	}
	return
}

// The account number of the account that sent the transfer.
func (r *WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) GetDebtorAccountNumber() (DebtorAccountNumber string) {
	if r != nil && r.DebtorAccountNumber != nil {
		DebtorAccountNumber = *r.DebtorAccountNumber
	}
	return
}

// The routing number of the account that sent the transfer.
func (r *WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) GetDebtorRoutingNumber() (DebtorRoutingNumber string) {
	if r != nil && r.DebtorRoutingNumber != nil {
		DebtorRoutingNumber = *r.DebtorRoutingNumber
	}
	return
}

// The Real Time Payments network identification of the transfer
func (r *WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) GetTransactionIdentification() (TransactionIdentification string) {
	if r != nil && r.TransactionIdentification != nil {
		TransactionIdentification = *r.TransactionIdentification
	}
	return
}

// Additional information included with the transfer.
func (r *WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmation) GetRemittanceInformation() (RemittanceInformation string) {
	if r != nil && r.RemittanceInformation != nil {
		RemittanceInformation = *r.RemittanceInformation
	}
	return
}

type WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency string

const (
	WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyCad WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "CAD"
	WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyChf WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "CHF"
	WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyEur WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "EUR"
	WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyGbp WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "GBP"
	WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyJpy WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "JPY"
	WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyUsd WireTransferSimulationTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "USD"
)

//
type WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal struct {
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
func (r *WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The description on the reversal message from Fedwire.
func (r *WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The Fedwire cycle date for the wire reversal.
func (r *WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) GetInputCycleDate() (InputCycleDate string) {
	if r != nil && r.InputCycleDate != nil {
		InputCycleDate = *r.InputCycleDate
	}
	return
}

// The Fedwire sequence number.
func (r *WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) GetInputSequenceNumber() (InputSequenceNumber string) {
	if r != nil && r.InputSequenceNumber != nil {
		InputSequenceNumber = *r.InputSequenceNumber
	}
	return
}

// The Fedwire input source identifier.
func (r *WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) GetInputSource() (InputSource string) {
	if r != nil && r.InputSource != nil {
		InputSource = *r.InputSource
	}
	return
}

// The Fedwire transaction identifier.
func (r *WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) GetInputMessageAccountabilityData() (InputMessageAccountabilityData string) {
	if r != nil && r.InputMessageAccountabilityData != nil {
		InputMessageAccountabilityData = *r.InputMessageAccountabilityData
	}
	return
}

// The Fedwire transaction identifier for the wire transfer that was reversed.
func (r *WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) GetPreviousMessageInputMessageAccountabilityData() (PreviousMessageInputMessageAccountabilityData string) {
	if r != nil && r.PreviousMessageInputMessageAccountabilityData != nil {
		PreviousMessageInputMessageAccountabilityData = *r.PreviousMessageInputMessageAccountabilityData
	}
	return
}

// The Fedwire cycle date for the wire transfer that was reversed.
func (r *WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) GetPreviousMessageInputCycleDate() (PreviousMessageInputCycleDate string) {
	if r != nil && r.PreviousMessageInputCycleDate != nil {
		PreviousMessageInputCycleDate = *r.PreviousMessageInputCycleDate
	}
	return
}

// The Fedwire sequence number for the wire transfer that was reversed.
func (r *WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) GetPreviousMessageInputSequenceNumber() (PreviousMessageInputSequenceNumber string) {
	if r != nil && r.PreviousMessageInputSequenceNumber != nil {
		PreviousMessageInputSequenceNumber = *r.PreviousMessageInputSequenceNumber
	}
	return
}

// The Fedwire input source identifier for the wire transfer that was reversed.
func (r *WireTransferSimulationTransactionSourceInboundWireDrawdownPaymentReversal) GetPreviousMessageInputSource() (PreviousMessageInputSource string) {
	if r != nil && r.PreviousMessageInputSource != nil {
		PreviousMessageInputSource = *r.PreviousMessageInputSource
	}
	return
}

//
type WireTransferSimulationTransactionSourceInboundWireDrawdownPayment struct {
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
func (r *WireTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetBeneficiaryAddressLine1() (BeneficiaryAddressLine1 string) {
	if r != nil && r.BeneficiaryAddressLine1 != nil {
		BeneficiaryAddressLine1 = *r.BeneficiaryAddressLine1
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetBeneficiaryAddressLine2() (BeneficiaryAddressLine2 string) {
	if r != nil && r.BeneficiaryAddressLine2 != nil {
		BeneficiaryAddressLine2 = *r.BeneficiaryAddressLine2
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetBeneficiaryAddressLine3() (BeneficiaryAddressLine3 string) {
	if r != nil && r.BeneficiaryAddressLine3 != nil {
		BeneficiaryAddressLine3 = *r.BeneficiaryAddressLine3
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetBeneficiaryName() (BeneficiaryName string) {
	if r != nil && r.BeneficiaryName != nil {
		BeneficiaryName = *r.BeneficiaryName
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetBeneficiaryReference() (BeneficiaryReference string) {
	if r != nil && r.BeneficiaryReference != nil {
		BeneficiaryReference = *r.BeneficiaryReference
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetInputMessageAccountabilityData() (InputMessageAccountabilityData string) {
	if r != nil && r.InputMessageAccountabilityData != nil {
		InputMessageAccountabilityData = *r.InputMessageAccountabilityData
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetOriginatorAddressLine1() (OriginatorAddressLine1 string) {
	if r != nil && r.OriginatorAddressLine1 != nil {
		OriginatorAddressLine1 = *r.OriginatorAddressLine1
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetOriginatorAddressLine2() (OriginatorAddressLine2 string) {
	if r != nil && r.OriginatorAddressLine2 != nil {
		OriginatorAddressLine2 = *r.OriginatorAddressLine2
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetOriginatorAddressLine3() (OriginatorAddressLine3 string) {
	if r != nil && r.OriginatorAddressLine3 != nil {
		OriginatorAddressLine3 = *r.OriginatorAddressLine3
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetOriginatorName() (OriginatorName string) {
	if r != nil && r.OriginatorName != nil {
		OriginatorName = *r.OriginatorName
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundWireDrawdownPayment) GetOriginatorToBeneficiaryInformation() (OriginatorToBeneficiaryInformation string) {
	if r != nil && r.OriginatorToBeneficiaryInformation != nil {
		OriginatorToBeneficiaryInformation = *r.OriginatorToBeneficiaryInformation
	}
	return
}

//
type WireTransferSimulationTransactionSourceInboundWireReversal struct {
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
func (r *WireTransferSimulationTransactionSourceInboundWireReversal) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The description on the reversal message from Fedwire.
func (r *WireTransferSimulationTransactionSourceInboundWireReversal) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The Fedwire cycle date for the wire reversal.
func (r *WireTransferSimulationTransactionSourceInboundWireReversal) GetInputCycleDate() (InputCycleDate string) {
	if r != nil && r.InputCycleDate != nil {
		InputCycleDate = *r.InputCycleDate
	}
	return
}

// The Fedwire sequence number.
func (r *WireTransferSimulationTransactionSourceInboundWireReversal) GetInputSequenceNumber() (InputSequenceNumber string) {
	if r != nil && r.InputSequenceNumber != nil {
		InputSequenceNumber = *r.InputSequenceNumber
	}
	return
}

// The Fedwire input source identifier.
func (r *WireTransferSimulationTransactionSourceInboundWireReversal) GetInputSource() (InputSource string) {
	if r != nil && r.InputSource != nil {
		InputSource = *r.InputSource
	}
	return
}

// The Fedwire transaction identifier.
func (r *WireTransferSimulationTransactionSourceInboundWireReversal) GetInputMessageAccountabilityData() (InputMessageAccountabilityData string) {
	if r != nil && r.InputMessageAccountabilityData != nil {
		InputMessageAccountabilityData = *r.InputMessageAccountabilityData
	}
	return
}

// The Fedwire transaction identifier for the wire transfer that was reversed.
func (r *WireTransferSimulationTransactionSourceInboundWireReversal) GetPreviousMessageInputMessageAccountabilityData() (PreviousMessageInputMessageAccountabilityData string) {
	if r != nil && r.PreviousMessageInputMessageAccountabilityData != nil {
		PreviousMessageInputMessageAccountabilityData = *r.PreviousMessageInputMessageAccountabilityData
	}
	return
}

// The Fedwire cycle date for the wire transfer that was reversed.
func (r *WireTransferSimulationTransactionSourceInboundWireReversal) GetPreviousMessageInputCycleDate() (PreviousMessageInputCycleDate string) {
	if r != nil && r.PreviousMessageInputCycleDate != nil {
		PreviousMessageInputCycleDate = *r.PreviousMessageInputCycleDate
	}
	return
}

// The Fedwire sequence number for the wire transfer that was reversed.
func (r *WireTransferSimulationTransactionSourceInboundWireReversal) GetPreviousMessageInputSequenceNumber() (PreviousMessageInputSequenceNumber string) {
	if r != nil && r.PreviousMessageInputSequenceNumber != nil {
		PreviousMessageInputSequenceNumber = *r.PreviousMessageInputSequenceNumber
	}
	return
}

// The Fedwire input source identifier for the wire transfer that was reversed.
func (r *WireTransferSimulationTransactionSourceInboundWireReversal) GetPreviousMessageInputSource() (PreviousMessageInputSource string) {
	if r != nil && r.PreviousMessageInputSource != nil {
		PreviousMessageInputSource = *r.PreviousMessageInputSource
	}
	return
}

// Information included in the wire reversal for the receiving financial
// institution.
func (r *WireTransferSimulationTransactionSourceInboundWireReversal) GetReceiverFinancialInstitutionInformation() (ReceiverFinancialInstitutionInformation string) {
	if r != nil && r.ReceiverFinancialInstitutionInformation != nil {
		ReceiverFinancialInstitutionInformation = *r.ReceiverFinancialInstitutionInformation
	}
	return
}

// Additional financial institution information included in the wire reversal.
func (r *WireTransferSimulationTransactionSourceInboundWireReversal) GetFinancialInstitutionToFinancialInstitutionInformation() (FinancialInstitutionToFinancialInstitutionInformation string) {
	if r != nil && r.FinancialInstitutionToFinancialInstitutionInformation != nil {
		FinancialInstitutionToFinancialInstitutionInformation = *r.FinancialInstitutionToFinancialInstitutionInformation
	}
	return
}

//
type WireTransferSimulationTransactionSourceInboundWireTransfer struct {
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
func (r *WireTransferSimulationTransactionSourceInboundWireTransfer) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundWireTransfer) GetBeneficiaryAddressLine1() (BeneficiaryAddressLine1 string) {
	if r != nil && r.BeneficiaryAddressLine1 != nil {
		BeneficiaryAddressLine1 = *r.BeneficiaryAddressLine1
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundWireTransfer) GetBeneficiaryAddressLine2() (BeneficiaryAddressLine2 string) {
	if r != nil && r.BeneficiaryAddressLine2 != nil {
		BeneficiaryAddressLine2 = *r.BeneficiaryAddressLine2
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundWireTransfer) GetBeneficiaryAddressLine3() (BeneficiaryAddressLine3 string) {
	if r != nil && r.BeneficiaryAddressLine3 != nil {
		BeneficiaryAddressLine3 = *r.BeneficiaryAddressLine3
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundWireTransfer) GetBeneficiaryName() (BeneficiaryName string) {
	if r != nil && r.BeneficiaryName != nil {
		BeneficiaryName = *r.BeneficiaryName
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundWireTransfer) GetBeneficiaryReference() (BeneficiaryReference string) {
	if r != nil && r.BeneficiaryReference != nil {
		BeneficiaryReference = *r.BeneficiaryReference
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundWireTransfer) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundWireTransfer) GetInputMessageAccountabilityData() (InputMessageAccountabilityData string) {
	if r != nil && r.InputMessageAccountabilityData != nil {
		InputMessageAccountabilityData = *r.InputMessageAccountabilityData
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundWireTransfer) GetOriginatorAddressLine1() (OriginatorAddressLine1 string) {
	if r != nil && r.OriginatorAddressLine1 != nil {
		OriginatorAddressLine1 = *r.OriginatorAddressLine1
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundWireTransfer) GetOriginatorAddressLine2() (OriginatorAddressLine2 string) {
	if r != nil && r.OriginatorAddressLine2 != nil {
		OriginatorAddressLine2 = *r.OriginatorAddressLine2
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundWireTransfer) GetOriginatorAddressLine3() (OriginatorAddressLine3 string) {
	if r != nil && r.OriginatorAddressLine3 != nil {
		OriginatorAddressLine3 = *r.OriginatorAddressLine3
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundWireTransfer) GetOriginatorName() (OriginatorName string) {
	if r != nil && r.OriginatorName != nil {
		OriginatorName = *r.OriginatorName
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInboundWireTransfer) GetOriginatorToBeneficiaryInformation() (OriginatorToBeneficiaryInformation string) {
	if r != nil && r.OriginatorToBeneficiaryInformation != nil {
		OriginatorToBeneficiaryInformation = *r.OriginatorToBeneficiaryInformation
	}
	return
}

//
type WireTransferSimulationTransactionSourceInternalSource struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency *WireTransferSimulationTransactionSourceInternalSourceCurrency `json:"currency"`
	//
	Reason *WireTransferSimulationTransactionSourceInternalSourceReason `json:"reason"`
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *WireTransferSimulationTransactionSourceInternalSource) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
// currency.
func (r *WireTransferSimulationTransactionSourceInternalSource) GetCurrency() (Currency WireTransferSimulationTransactionSourceInternalSourceCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

func (r *WireTransferSimulationTransactionSourceInternalSource) GetReason() (Reason WireTransferSimulationTransactionSourceInternalSourceReason) {
	if r != nil && r.Reason != nil {
		Reason = *r.Reason
	}
	return
}

type WireTransferSimulationTransactionSourceInternalSourceCurrency string

const (
	WireTransferSimulationTransactionSourceInternalSourceCurrencyCad WireTransferSimulationTransactionSourceInternalSourceCurrency = "CAD"
	WireTransferSimulationTransactionSourceInternalSourceCurrencyChf WireTransferSimulationTransactionSourceInternalSourceCurrency = "CHF"
	WireTransferSimulationTransactionSourceInternalSourceCurrencyEur WireTransferSimulationTransactionSourceInternalSourceCurrency = "EUR"
	WireTransferSimulationTransactionSourceInternalSourceCurrencyGbp WireTransferSimulationTransactionSourceInternalSourceCurrency = "GBP"
	WireTransferSimulationTransactionSourceInternalSourceCurrencyJpy WireTransferSimulationTransactionSourceInternalSourceCurrency = "JPY"
	WireTransferSimulationTransactionSourceInternalSourceCurrencyUsd WireTransferSimulationTransactionSourceInternalSourceCurrency = "USD"
)

type WireTransferSimulationTransactionSourceInternalSourceReason string

const (
	WireTransferSimulationTransactionSourceInternalSourceReasonCashback           WireTransferSimulationTransactionSourceInternalSourceReason = "cashback"
	WireTransferSimulationTransactionSourceInternalSourceReasonEmpyrealAdjustment WireTransferSimulationTransactionSourceInternalSourceReason = "empyreal_adjustment"
	WireTransferSimulationTransactionSourceInternalSourceReasonError              WireTransferSimulationTransactionSourceInternalSourceReason = "error"
	WireTransferSimulationTransactionSourceInternalSourceReasonErrorCorrection    WireTransferSimulationTransactionSourceInternalSourceReason = "error_correction"
	WireTransferSimulationTransactionSourceInternalSourceReasonFees               WireTransferSimulationTransactionSourceInternalSourceReason = "fees"
	WireTransferSimulationTransactionSourceInternalSourceReasonInterest           WireTransferSimulationTransactionSourceInternalSourceReason = "interest"
	WireTransferSimulationTransactionSourceInternalSourceReasonSampleFunds        WireTransferSimulationTransactionSourceInternalSourceReason = "sample_funds"
	WireTransferSimulationTransactionSourceInternalSourceReasonSampleFundsReturn  WireTransferSimulationTransactionSourceInternalSourceReason = "sample_funds_return"
)

//
type WireTransferSimulationTransactionSourceCardRouteRefund struct {
	// The refunded amount in the minor unit of the refunded currency. For dollars, for
	// example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the refund
	// currency.
	Currency *WireTransferSimulationTransactionSourceCardRouteRefundCurrency `json:"currency"`
	//
	MerchantAcceptorId *string `json:"merchant_acceptor_id"`
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
func (r *WireTransferSimulationTransactionSourceCardRouteRefund) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the refund
// currency.
func (r *WireTransferSimulationTransactionSourceCardRouteRefund) GetCurrency() (Currency WireTransferSimulationTransactionSourceCardRouteRefundCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

func (r *WireTransferSimulationTransactionSourceCardRouteRefund) GetMerchantAcceptorId() (MerchantAcceptorId string) {
	if r != nil && r.MerchantAcceptorId != nil {
		MerchantAcceptorId = *r.MerchantAcceptorId
	}
	return
}

func (r *WireTransferSimulationTransactionSourceCardRouteRefund) GetMerchantCity() (MerchantCity string) {
	if r != nil && r.MerchantCity != nil {
		MerchantCity = *r.MerchantCity
	}
	return
}

func (r *WireTransferSimulationTransactionSourceCardRouteRefund) GetMerchantCountry() (MerchantCountry string) {
	if r != nil && r.MerchantCountry != nil {
		MerchantCountry = *r.MerchantCountry
	}
	return
}

func (r *WireTransferSimulationTransactionSourceCardRouteRefund) GetMerchantDescriptor() (MerchantDescriptor string) {
	if r != nil && r.MerchantDescriptor != nil {
		MerchantDescriptor = *r.MerchantDescriptor
	}
	return
}

func (r *WireTransferSimulationTransactionSourceCardRouteRefund) GetMerchantState() (MerchantState string) {
	if r != nil && r.MerchantState != nil {
		MerchantState = *r.MerchantState
	}
	return
}

func (r *WireTransferSimulationTransactionSourceCardRouteRefund) GetMerchantCategoryCode() (MerchantCategoryCode string) {
	if r != nil && r.MerchantCategoryCode != nil {
		MerchantCategoryCode = *r.MerchantCategoryCode
	}
	return
}

type WireTransferSimulationTransactionSourceCardRouteRefundCurrency string

const (
	WireTransferSimulationTransactionSourceCardRouteRefundCurrencyCad WireTransferSimulationTransactionSourceCardRouteRefundCurrency = "CAD"
	WireTransferSimulationTransactionSourceCardRouteRefundCurrencyChf WireTransferSimulationTransactionSourceCardRouteRefundCurrency = "CHF"
	WireTransferSimulationTransactionSourceCardRouteRefundCurrencyEur WireTransferSimulationTransactionSourceCardRouteRefundCurrency = "EUR"
	WireTransferSimulationTransactionSourceCardRouteRefundCurrencyGbp WireTransferSimulationTransactionSourceCardRouteRefundCurrency = "GBP"
	WireTransferSimulationTransactionSourceCardRouteRefundCurrencyJpy WireTransferSimulationTransactionSourceCardRouteRefundCurrency = "JPY"
	WireTransferSimulationTransactionSourceCardRouteRefundCurrencyUsd WireTransferSimulationTransactionSourceCardRouteRefundCurrency = "USD"
)

//
type WireTransferSimulationTransactionSourceCardRouteSettlement struct {
	// The settled amount in the minor unit of the settlement currency. For dollars,
	// for example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the settlement
	// currency.
	Currency *WireTransferSimulationTransactionSourceCardRouteSettlementCurrency `json:"currency"`
	//
	MerchantAcceptorId *string `json:"merchant_acceptor_id"`
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
func (r *WireTransferSimulationTransactionSourceCardRouteSettlement) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the settlement
// currency.
func (r *WireTransferSimulationTransactionSourceCardRouteSettlement) GetCurrency() (Currency WireTransferSimulationTransactionSourceCardRouteSettlementCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

func (r *WireTransferSimulationTransactionSourceCardRouteSettlement) GetMerchantAcceptorId() (MerchantAcceptorId string) {
	if r != nil && r.MerchantAcceptorId != nil {
		MerchantAcceptorId = *r.MerchantAcceptorId
	}
	return
}

func (r *WireTransferSimulationTransactionSourceCardRouteSettlement) GetMerchantCity() (MerchantCity string) {
	if r != nil && r.MerchantCity != nil {
		MerchantCity = *r.MerchantCity
	}
	return
}

func (r *WireTransferSimulationTransactionSourceCardRouteSettlement) GetMerchantCountry() (MerchantCountry string) {
	if r != nil && r.MerchantCountry != nil {
		MerchantCountry = *r.MerchantCountry
	}
	return
}

func (r *WireTransferSimulationTransactionSourceCardRouteSettlement) GetMerchantDescriptor() (MerchantDescriptor string) {
	if r != nil && r.MerchantDescriptor != nil {
		MerchantDescriptor = *r.MerchantDescriptor
	}
	return
}

func (r *WireTransferSimulationTransactionSourceCardRouteSettlement) GetMerchantState() (MerchantState string) {
	if r != nil && r.MerchantState != nil {
		MerchantState = *r.MerchantState
	}
	return
}

func (r *WireTransferSimulationTransactionSourceCardRouteSettlement) GetMerchantCategoryCode() (MerchantCategoryCode string) {
	if r != nil && r.MerchantCategoryCode != nil {
		MerchantCategoryCode = *r.MerchantCategoryCode
	}
	return
}

type WireTransferSimulationTransactionSourceCardRouteSettlementCurrency string

const (
	WireTransferSimulationTransactionSourceCardRouteSettlementCurrencyCad WireTransferSimulationTransactionSourceCardRouteSettlementCurrency = "CAD"
	WireTransferSimulationTransactionSourceCardRouteSettlementCurrencyChf WireTransferSimulationTransactionSourceCardRouteSettlementCurrency = "CHF"
	WireTransferSimulationTransactionSourceCardRouteSettlementCurrencyEur WireTransferSimulationTransactionSourceCardRouteSettlementCurrency = "EUR"
	WireTransferSimulationTransactionSourceCardRouteSettlementCurrencyGbp WireTransferSimulationTransactionSourceCardRouteSettlementCurrency = "GBP"
	WireTransferSimulationTransactionSourceCardRouteSettlementCurrencyJpy WireTransferSimulationTransactionSourceCardRouteSettlementCurrency = "JPY"
	WireTransferSimulationTransactionSourceCardRouteSettlementCurrencyUsd WireTransferSimulationTransactionSourceCardRouteSettlementCurrency = "USD"
)

//
type WireTransferSimulationTransactionSourceSampleFunds struct {
	// Where the sample funds came from.
	Originator *string `json:"originator"`
}

// Where the sample funds came from.
func (r *WireTransferSimulationTransactionSourceSampleFunds) GetOriginator() (Originator string) {
	if r != nil && r.Originator != nil {
		Originator = *r.Originator
	}
	return
}

//
type WireTransferSimulationTransactionSourceWireDrawdownPaymentIntention struct {
	// The transfer amount in USD cents.
	Amount *int `json:"amount"`
	//
	AccountNumber *string `json:"account_number"`
	//
	RoutingNumber *string `json:"routing_number"`
	//
	MessageToRecipient *string `json:"message_to_recipient"`
	//
	TransferId *string `json:"transfer_id"`
}

// The transfer amount in USD cents.
func (r *WireTransferSimulationTransactionSourceWireDrawdownPaymentIntention) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *WireTransferSimulationTransactionSourceWireDrawdownPaymentIntention) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

func (r *WireTransferSimulationTransactionSourceWireDrawdownPaymentIntention) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

func (r *WireTransferSimulationTransactionSourceWireDrawdownPaymentIntention) GetMessageToRecipient() (MessageToRecipient string) {
	if r != nil && r.MessageToRecipient != nil {
		MessageToRecipient = *r.MessageToRecipient
	}
	return
}

func (r *WireTransferSimulationTransactionSourceWireDrawdownPaymentIntention) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}

//
type WireTransferSimulationTransactionSourceWireDrawdownPaymentRejection struct {
	//
	TransferId *string `json:"transfer_id"`
}

func (r *WireTransferSimulationTransactionSourceWireDrawdownPaymentRejection) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}

//
type WireTransferSimulationTransactionSourceWireTransferIntention struct {
	// The transfer amount in USD cents.
	Amount *int `json:"amount"`
	// The destination account number.
	AccountNumber *string `json:"account_number"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber *string `json:"routing_number"`
	// The message that will show on the recipient's bank statement.
	MessageToRecipient *string `json:"message_to_recipient"`
	//
	TransferId *string `json:"transfer_id"`
}

// The transfer amount in USD cents.
func (r *WireTransferSimulationTransactionSourceWireTransferIntention) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The destination account number.
func (r *WireTransferSimulationTransactionSourceWireTransferIntention) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

// The American Bankers' Association (ABA) Routing Transit Number (RTN).
func (r *WireTransferSimulationTransactionSourceWireTransferIntention) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

// The message that will show on the recipient's bank statement.
func (r *WireTransferSimulationTransactionSourceWireTransferIntention) GetMessageToRecipient() (MessageToRecipient string) {
	if r != nil && r.MessageToRecipient != nil {
		MessageToRecipient = *r.MessageToRecipient
	}
	return
}

func (r *WireTransferSimulationTransactionSourceWireTransferIntention) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}

//
type WireTransferSimulationTransactionSourceWireTransferRejection struct {
	//
	TransferId *string `json:"transfer_id"`
}

func (r *WireTransferSimulationTransactionSourceWireTransferRejection) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}

type WireTransferSimulationTransactionType string

const (
	WireTransferSimulationTransactionTypeTransaction WireTransferSimulationTransactionType = "transaction"
)

type WireTransferSimulationType string

const (
	WireTransferSimulationTypeInboundWireTransferSimulationResult WireTransferSimulationType = "inbound_wire_transfer_simulation_result"
)

type SimulateAWireTransferToYourAccountParameters struct {
	// The identifier of the Account Number the inbound Wire Transfer is for.
	AccountNumberId *string `json:"account_number_id"`
	// The transfer amount in cents. Must be positive.
	Amount *int `json:"amount"`
}

// The identifier of the Account Number the inbound Wire Transfer is for.
func (r *SimulateAWireTransferToYourAccountParameters) GetAccountNumberId() (AccountNumberId string) {
	if r != nil && r.AccountNumberId != nil {
		AccountNumberId = *r.AccountNumberId
	}
	return
}

// The transfer amount in cents. Must be positive.
func (r *SimulateAWireTransferToYourAccountParameters) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

//
type CardAuthorizationSimulation struct {
	// If the authorization attempt succeeds, this will contain the resulting Pending
	// Transaction object. The Pending Transaction's `source` will be of
	// `category: card_authorization`.
	PendingTransaction *CardAuthorizationSimulationPendingTransaction `json:"pending_transaction"`
	// If the authorization attempt fails, this will contain the resulting
	// [Declined Transaction](#declined-transactions) object. The Declined
	// Transaction's `source` will be of `category: card_decline`.
	DeclinedTransaction *CardAuthorizationSimulationDeclinedTransaction `json:"declined_transaction"`
	// A constant representing the object's type. For this resource it will always be
	// `inbound_card_authorization_simulation_result`.
	Type *CardAuthorizationSimulationType `json:"type"`
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

//
type CardAuthorizationSimulationPendingTransaction struct {
	// The identifier for the account this Pending Transaction belongs to.
	AccountId *string `json:"account_id"`
	// The Pending Transaction amount in the minor unit of its currency. For dollars,
	// for example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Pending
	// Transaction's currency. This will match the currency on the Pending
	// Transcation's Account.
	Currency *CardAuthorizationSimulationPendingTransactionCurrency `json:"currency"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the Pending
	// Transaction occured.
	CreatedAt *string `json:"created_at"`
	// For a Pending Transaction related to a transfer, this is the description you
	// provide. For a Pending Transaction related to a payment, this is the description
	// the vendor provides.
	Description *string `json:"description"`
	// The Pending Transaction identifier.
	Id *string `json:"id"`
	// The identifier for the route this Pending Transaction came through. Routes are
	// things like cards and ACH details.
	RouteId *string `json:"route_id"`
	// The type of the route this Pending Transaction came through.
	RouteType *string `json:"route_type"`
	// This is an object giving more details on the network-level event that caused the
	// Pending Transaction. For example, for a card transaction this lists the
	// merchant's industry and location.
	Source *CardAuthorizationSimulationPendingTransactionSource `json:"source"`
	// Whether the Pending Transaction has been confirmed and has an associated
	// Transaction.
	Status *CardAuthorizationSimulationPendingTransactionStatus `json:"status"`
	// A constant representing the object's type. For this resource it will always be
	// `pending_transaction`.
	Type *CardAuthorizationSimulationPendingTransactionType `json:"type"`
}

// The identifier for the account this Pending Transaction belongs to.
func (r *CardAuthorizationSimulationPendingTransaction) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

// The Pending Transaction amount in the minor unit of its currency. For dollars,
// for example, this is cents.
func (r *CardAuthorizationSimulationPendingTransaction) GetAmount() (Amount int) {
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
func (r *CardAuthorizationSimulationPendingTransaction) GetId() (Id string) {
	if r != nil && r.Id != nil {
		Id = *r.Id
	}
	return
}

// The identifier for the route this Pending Transaction came through. Routes are
// things like cards and ACH details.
func (r *CardAuthorizationSimulationPendingTransaction) GetRouteId() (RouteId string) {
	if r != nil && r.RouteId != nil {
		RouteId = *r.RouteId
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

type CardAuthorizationSimulationPendingTransactionCurrency string

const (
	CardAuthorizationSimulationPendingTransactionCurrencyCad CardAuthorizationSimulationPendingTransactionCurrency = "CAD"
	CardAuthorizationSimulationPendingTransactionCurrencyChf CardAuthorizationSimulationPendingTransactionCurrency = "CHF"
	CardAuthorizationSimulationPendingTransactionCurrencyEur CardAuthorizationSimulationPendingTransactionCurrency = "EUR"
	CardAuthorizationSimulationPendingTransactionCurrencyGbp CardAuthorizationSimulationPendingTransactionCurrency = "GBP"
	CardAuthorizationSimulationPendingTransactionCurrencyJpy CardAuthorizationSimulationPendingTransactionCurrency = "JPY"
	CardAuthorizationSimulationPendingTransactionCurrencyUsd CardAuthorizationSimulationPendingTransactionCurrency = "USD"
)

//
type CardAuthorizationSimulationPendingTransactionSource struct {
	// The type of transaction that took place. We may add additional possible values
	// for this enum over time; your application should be able to handle such
	// additions gracefully.
	Category *CardAuthorizationSimulationPendingTransactionSourceCategory `json:"category"`
	// A Account Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `account_transfer_instruction`.
	AccountTransferInstruction *CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstruction `json:"account_transfer_instruction"`
	// A ACH Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `ach_transfer_instruction`.
	ACHTransferInstruction *CardAuthorizationSimulationPendingTransactionSourceACHTransferInstruction `json:"ach_transfer_instruction"`
	// A Card Authorization object. This field will be present in the JSON response if
	// and only if `category` is equal to `card_authorization`.
	CardAuthorization *CardAuthorizationSimulationPendingTransactionSourceCardAuthorization `json:"card_authorization"`
	// A Check Deposit Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_deposit_instruction`.
	CheckDepositInstruction *CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstruction `json:"check_deposit_instruction"`
	// A Check Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_instruction`.
	CheckTransferInstruction *CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstruction `json:"check_transfer_instruction"`
	// A Deprecated Card Authorization object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_route_authorization`.
	CardRouteAuthorization *CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorization `json:"card_route_authorization"`
	// A Wire Drawdown Payment Instruction object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `wire_drawdown_payment_instruction`.
	WireDrawdownPaymentInstruction *CardAuthorizationSimulationPendingTransactionSourceWireDrawdownPaymentInstruction `json:"wire_drawdown_payment_instruction"`
	// A Wire Transfer Instruction object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_instruction`.
	WireTransferInstruction *CardAuthorizationSimulationPendingTransactionSourceWireTransferInstruction `json:"wire_transfer_instruction"`
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

type CardAuthorizationSimulationPendingTransactionSourceCategory string

const (
	CardAuthorizationSimulationPendingTransactionSourceCategoryAccountTransferInstruction          CardAuthorizationSimulationPendingTransactionSourceCategory = "account_transfer_instruction"
	CardAuthorizationSimulationPendingTransactionSourceCategoryACHTransferInstruction              CardAuthorizationSimulationPendingTransactionSourceCategory = "ach_transfer_instruction"
	CardAuthorizationSimulationPendingTransactionSourceCategoryCardAuthorization                   CardAuthorizationSimulationPendingTransactionSourceCategory = "card_authorization"
	CardAuthorizationSimulationPendingTransactionSourceCategoryCheckDepositInstruction             CardAuthorizationSimulationPendingTransactionSourceCategory = "check_deposit_instruction"
	CardAuthorizationSimulationPendingTransactionSourceCategoryCheckTransferInstruction            CardAuthorizationSimulationPendingTransactionSourceCategory = "check_transfer_instruction"
	CardAuthorizationSimulationPendingTransactionSourceCategoryCardRouteAuthorization              CardAuthorizationSimulationPendingTransactionSourceCategory = "card_route_authorization"
	CardAuthorizationSimulationPendingTransactionSourceCategoryRealTimePaymentsTransferInstruction CardAuthorizationSimulationPendingTransactionSourceCategory = "real_time_payments_transfer_instruction"
	CardAuthorizationSimulationPendingTransactionSourceCategoryWireDrawdownPaymentInstruction      CardAuthorizationSimulationPendingTransactionSourceCategory = "wire_drawdown_payment_instruction"
	CardAuthorizationSimulationPendingTransactionSourceCategoryWireTransferInstruction             CardAuthorizationSimulationPendingTransactionSourceCategory = "wire_transfer_instruction"
	CardAuthorizationSimulationPendingTransactionSourceCategoryOther                               CardAuthorizationSimulationPendingTransactionSourceCategory = "other"
)

//
type CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency *CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstructionCurrency `json:"currency"`
	// The identifier of the Account Transfer that led to this Pending Transaction.
	TransferId *string `json:"transfer_id"`
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstruction) GetAmount() (Amount int) {
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
func (r *CardAuthorizationSimulationPendingTransactionSourceAccountTransferInstruction) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
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

//
type CardAuthorizationSimulationPendingTransactionSourceACHTransferInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int `json:"amount"`
	// The identifier of the ACH Transfer that led to this Pending Transaction.
	TransferId *string `json:"transfer_id"`
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *CardAuthorizationSimulationPendingTransactionSourceACHTransferInstruction) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The identifier of the ACH Transfer that led to this Pending Transaction.
func (r *CardAuthorizationSimulationPendingTransactionSourceACHTransferInstruction) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}

//
type CardAuthorizationSimulationPendingTransactionSourceCardAuthorization struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrency `json:"currency"`
	//
	MerchantAcceptorId *string `json:"merchant_acceptor_id"`
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
	RealTimeDecisionId *string `json:"real_time_decision_id"`
	// If the authorization was made via a Digital Wallet Token (such as an Apple Pay
	// purchase), the identifier of the token that was used.
	DigitalWalletTokenId *string `json:"digital_wallet_token_id"`
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *CardAuthorizationSimulationPendingTransactionSourceCardAuthorization) GetAmount() (Amount int) {
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

func (r *CardAuthorizationSimulationPendingTransactionSourceCardAuthorization) GetMerchantAcceptorId() (MerchantAcceptorId string) {
	if r != nil && r.MerchantAcceptorId != nil {
		MerchantAcceptorId = *r.MerchantAcceptorId
	}
	return
}

func (r *CardAuthorizationSimulationPendingTransactionSourceCardAuthorization) GetMerchantCity() (MerchantCity string) {
	if r != nil && r.MerchantCity != nil {
		MerchantCity = *r.MerchantCity
	}
	return
}

func (r *CardAuthorizationSimulationPendingTransactionSourceCardAuthorization) GetMerchantCountry() (MerchantCountry string) {
	if r != nil && r.MerchantCountry != nil {
		MerchantCountry = *r.MerchantCountry
	}
	return
}

func (r *CardAuthorizationSimulationPendingTransactionSourceCardAuthorization) GetMerchantDescriptor() (MerchantDescriptor string) {
	if r != nil && r.MerchantDescriptor != nil {
		MerchantDescriptor = *r.MerchantDescriptor
	}
	return
}

func (r *CardAuthorizationSimulationPendingTransactionSourceCardAuthorization) GetMerchantCategoryCode() (MerchantCategoryCode string) {
	if r != nil && r.MerchantCategoryCode != nil {
		MerchantCategoryCode = *r.MerchantCategoryCode
	}
	return
}

// The identifier of the Real-Time Decision sent to approve or decline this
// transaction.
func (r *CardAuthorizationSimulationPendingTransactionSourceCardAuthorization) GetRealTimeDecisionId() (RealTimeDecisionId string) {
	if r != nil && r.RealTimeDecisionId != nil {
		RealTimeDecisionId = *r.RealTimeDecisionId
	}
	return
}

// If the authorization was made via a Digital Wallet Token (such as an Apple Pay
// purchase), the identifier of the token that was used.
func (r *CardAuthorizationSimulationPendingTransactionSourceCardAuthorization) GetDigitalWalletTokenId() (DigitalWalletTokenId string) {
	if r != nil && r.DigitalWalletTokenId != nil {
		DigitalWalletTokenId = *r.DigitalWalletTokenId
	}
	return
}

type CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrency string

const (
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrencyCad CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrency = "CAD"
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrencyChf CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrency = "CHF"
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrencyEur CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrency = "EUR"
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrencyGbp CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrency = "GBP"
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrencyJpy CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrency = "JPY"
	CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrencyUsd CardAuthorizationSimulationPendingTransactionSourceCardAuthorizationCurrency = "USD"
)

//
type CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstructionCurrency `json:"currency"`
	// The identifier of the File containing the image of the front of the check that
	// was deposited.
	FrontImageFileId *string `json:"front_image_file_id"`
	// The identifier of the File containing the image of the back of the check that
	// was deposited.
	BackImageFileId *string `json:"back_image_file_id"`
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstruction) GetAmount() (Amount int) {
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
func (r *CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstruction) GetFrontImageFileId() (FrontImageFileId string) {
	if r != nil && r.FrontImageFileId != nil {
		FrontImageFileId = *r.FrontImageFileId
	}
	return
}

// The identifier of the File containing the image of the back of the check that
// was deposited.
func (r *CardAuthorizationSimulationPendingTransactionSourceCheckDepositInstruction) GetBackImageFileId() (BackImageFileId string) {
	if r != nil && r.BackImageFileId != nil {
		BackImageFileId = *r.BackImageFileId
	}
	return
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

//
type CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstruction struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
	// currency.
	Currency *CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstructionCurrency `json:"currency"`
	// The identifier of the Check Transfer that led to this Pending Transaction.
	TransferId *string `json:"transfer_id"`
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstruction) GetAmount() (Amount int) {
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
func (r *CardAuthorizationSimulationPendingTransactionSourceCheckTransferInstruction) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
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

//
type CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorization struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationCurrency `json:"currency"`
	//
	MerchantAcceptorId *string `json:"merchant_acceptor_id"`
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
func (r *CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorization) GetAmount() (Amount int) {
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

func (r *CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorization) GetMerchantAcceptorId() (MerchantAcceptorId string) {
	if r != nil && r.MerchantAcceptorId != nil {
		MerchantAcceptorId = *r.MerchantAcceptorId
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

type CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationCurrency string

const (
	CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationCurrencyCad CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationCurrency = "CAD"
	CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationCurrencyChf CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationCurrency = "CHF"
	CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationCurrencyEur CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationCurrency = "EUR"
	CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationCurrencyGbp CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationCurrency = "GBP"
	CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationCurrencyJpy CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationCurrency = "JPY"
	CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationCurrencyUsd CardAuthorizationSimulationPendingTransactionSourceCardRouteAuthorizationCurrency = "USD"
)

//
type CardAuthorizationSimulationPendingTransactionSourceWireDrawdownPaymentInstruction struct {
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
func (r *CardAuthorizationSimulationPendingTransactionSourceWireDrawdownPaymentInstruction) GetAmount() (Amount int) {
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

//
type CardAuthorizationSimulationPendingTransactionSourceWireTransferInstruction struct {
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
	TransferId *string `json:"transfer_id"`
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *CardAuthorizationSimulationPendingTransactionSourceWireTransferInstruction) GetAmount() (Amount int) {
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

func (r *CardAuthorizationSimulationPendingTransactionSourceWireTransferInstruction) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
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

//
type CardAuthorizationSimulationDeclinedTransaction struct {
	// The identifier for the Account the Declined Transaction belongs to.
	AccountId *string `json:"account_id"`
	// The Declined Transaction amount in the minor unit of its currency. For dollars,
	// for example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Declined
	// Transaction's currency. This will match the currency on the Declined
	// Transcation's Account.
	Currency *CardAuthorizationSimulationDeclinedTransactionCurrency `json:"currency"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the
	// Transaction occured.
	CreatedAt *string `json:"created_at"`
	// This is the description the vendor provides.
	Description *string `json:"description"`
	// The Declined Transaction identifier.
	Id *string `json:"id"`
	// The identifier for the route this Declined Transaction came through. Routes are
	// things like cards and ACH details.
	RouteId *string `json:"route_id"`
	// The type of the route this Declined Transaction came through.
	RouteType *string `json:"route_type"`
	// This is an object giving more details on the network-level event that caused the
	// Declined Transaction. For example, for a card transaction this lists the
	// merchant's industry and location. Note that for backwards compatibility reasons,
	// additional undocumented keys may appear in this object. These should be treated
	// as deprecated and will be removed in the future.
	Source *CardAuthorizationSimulationDeclinedTransactionSource `json:"source"`
	// A constant representing the object's type. For this resource it will always be
	// `declined_transaction`.
	Type *CardAuthorizationSimulationDeclinedTransactionType `json:"type"`
}

// The identifier for the Account the Declined Transaction belongs to.
func (r *CardAuthorizationSimulationDeclinedTransaction) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

// The Declined Transaction amount in the minor unit of its currency. For dollars,
// for example, this is cents.
func (r *CardAuthorizationSimulationDeclinedTransaction) GetAmount() (Amount int) {
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
func (r *CardAuthorizationSimulationDeclinedTransaction) GetId() (Id string) {
	if r != nil && r.Id != nil {
		Id = *r.Id
	}
	return
}

// The identifier for the route this Declined Transaction came through. Routes are
// things like cards and ACH details.
func (r *CardAuthorizationSimulationDeclinedTransaction) GetRouteId() (RouteId string) {
	if r != nil && r.RouteId != nil {
		RouteId = *r.RouteId
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

type CardAuthorizationSimulationDeclinedTransactionCurrency string

const (
	CardAuthorizationSimulationDeclinedTransactionCurrencyCad CardAuthorizationSimulationDeclinedTransactionCurrency = "CAD"
	CardAuthorizationSimulationDeclinedTransactionCurrencyChf CardAuthorizationSimulationDeclinedTransactionCurrency = "CHF"
	CardAuthorizationSimulationDeclinedTransactionCurrencyEur CardAuthorizationSimulationDeclinedTransactionCurrency = "EUR"
	CardAuthorizationSimulationDeclinedTransactionCurrencyGbp CardAuthorizationSimulationDeclinedTransactionCurrency = "GBP"
	CardAuthorizationSimulationDeclinedTransactionCurrencyJpy CardAuthorizationSimulationDeclinedTransactionCurrency = "JPY"
	CardAuthorizationSimulationDeclinedTransactionCurrencyUsd CardAuthorizationSimulationDeclinedTransactionCurrency = "USD"
)

//
type CardAuthorizationSimulationDeclinedTransactionSource struct {
	// The type of decline that took place. We may add additional possible values for
	// this enum over time; your application should be able to handle such additions
	// gracefully.
	Category *CardAuthorizationSimulationDeclinedTransactionSourceCategory `json:"category"`
	// A ACH Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `ach_decline`.
	ACHDecline *CardAuthorizationSimulationDeclinedTransactionSourceACHDecline `json:"ach_decline"`
	// A Card Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_decline`.
	CardDecline *CardAuthorizationSimulationDeclinedTransactionSourceCardDecline `json:"card_decline"`
	// A Check Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `check_decline`.
	CheckDecline *CardAuthorizationSimulationDeclinedTransactionSourceCheckDecline `json:"check_decline"`
	// A Inbound Real Time Payments Transfer Decline object. This field will be present
	// in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_decline`.
	InboundRealTimePaymentsTransferDecline *CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline `json:"inbound_real_time_payments_transfer_decline"`
	// A International ACH Decline object. This field will be present in the JSON
	// response if and only if `category` is equal to `international_ach_decline`.
	InternationalACHDecline *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline `json:"international_ach_decline"`
	// A Deprecated Card Decline object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_route_decline`.
	CardRouteDecline *CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDecline `json:"card_route_decline"`
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

//
type CardAuthorizationSimulationDeclinedTransactionSourceACHDecline struct {
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
	OriginatorCompanyId *string `json:"originator_company_id"`
	// Why the ACH transfer was declined.
	Reason *CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason `json:"reason"`
	//
	ReceiverIdNumber *string `json:"receiver_id_number"`
	//
	ReceiverName *string `json:"receiver_name"`
	//
	TraceNumber *string `json:"trace_number"`
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceACHDecline) GetAmount() (Amount int) {
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

func (r *CardAuthorizationSimulationDeclinedTransactionSourceACHDecline) GetOriginatorCompanyId() (OriginatorCompanyId string) {
	if r != nil && r.OriginatorCompanyId != nil {
		OriginatorCompanyId = *r.OriginatorCompanyId
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

func (r *CardAuthorizationSimulationDeclinedTransactionSourceACHDecline) GetReceiverIdNumber() (ReceiverIdNumber string) {
	if r != nil && r.ReceiverIdNumber != nil {
		ReceiverIdNumber = *r.ReceiverIdNumber
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

type CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason string

const (
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonACHRouteCanceled             CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "ach_route_canceled"
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonACHRouteDisabled             CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "ach_route_disabled"
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonNoACHRoute                   CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "no_ach_route"
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonBreachesLimit                CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "breaches_limit"
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonCreditEntryRefusedByReceiver CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "credit_entry_refused_by_receiver"
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonGroupLocked                  CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "group_locked"
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonEntityNotActive              CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "entity_not_active"
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonInsufficientFunds            CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "insufficient_funds"
	CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReasonOriginatorRequest            CardAuthorizationSimulationDeclinedTransactionSourceACHDeclineReason = "originator_request"
)

//
type CardAuthorizationSimulationDeclinedTransactionSourceCardDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency *CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineCurrency `json:"currency"`
	//
	MerchantAcceptorId *string `json:"merchant_acceptor_id"`
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
	Reason *CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason `json:"reason"`
	// The identifier of the Real-Time Decision sent to approve or decline this
	// transaction.
	RealTimeDecisionId *string `json:"real_time_decision_id"`
	// If the authorization was attempted using a Digital Wallet Token (such as an
	// Apple Pay purchase), the identifier of the token that was used.
	DigitalWalletTokenId *string `json:"digital_wallet_token_id"`
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDecline) GetAmount() (Amount int) {
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

func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDecline) GetMerchantAcceptorId() (MerchantAcceptorId string) {
	if r != nil && r.MerchantAcceptorId != nil {
		MerchantAcceptorId = *r.MerchantAcceptorId
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDecline) GetMerchantCity() (MerchantCity string) {
	if r != nil && r.MerchantCity != nil {
		MerchantCity = *r.MerchantCity
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDecline) GetMerchantCountry() (MerchantCountry string) {
	if r != nil && r.MerchantCountry != nil {
		MerchantCountry = *r.MerchantCountry
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDecline) GetMerchantDescriptor() (MerchantDescriptor string) {
	if r != nil && r.MerchantDescriptor != nil {
		MerchantDescriptor = *r.MerchantDescriptor
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDecline) GetMerchantState() (MerchantState string) {
	if r != nil && r.MerchantState != nil {
		MerchantState = *r.MerchantState
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDecline) GetMerchantCategoryCode() (MerchantCategoryCode string) {
	if r != nil && r.MerchantCategoryCode != nil {
		MerchantCategoryCode = *r.MerchantCategoryCode
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

// The identifier of the Real-Time Decision sent to approve or decline this
// transaction.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDecline) GetRealTimeDecisionId() (RealTimeDecisionId string) {
	if r != nil && r.RealTimeDecisionId != nil {
		RealTimeDecisionId = *r.RealTimeDecisionId
	}
	return
}

// If the authorization was attempted using a Digital Wallet Token (such as an
// Apple Pay purchase), the identifier of the token that was used.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardDecline) GetDigitalWalletTokenId() (DigitalWalletTokenId string) {
	if r != nil && r.DigitalWalletTokenId != nil {
		DigitalWalletTokenId = *r.DigitalWalletTokenId
	}
	return
}

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
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonCardNotActive     CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "card_not_active"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonEntityNotActive   CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "entity_not_active"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonGroupLocked       CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "group_locked"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonInsufficientFunds CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "insufficient_funds"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonBreachesLimit     CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "breaches_limit"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonWebhookDeclined   CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "webhook_declined"
	CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReasonWebhookTimedOut   CardAuthorizationSimulationDeclinedTransactionSourceCardDeclineReason = "webhook_timed_out"
)

//
type CardAuthorizationSimulationDeclinedTransactionSourceCheckDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount *int `json:"amount"`
	//
	AuxiliaryOnUs *string `json:"auxiliary_on_us"`
	// Why the check was declined.
	Reason *CardAuthorizationSimulationDeclinedTransactionSourceCheckDeclineReason `json:"reason"`
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceCheckDecline) GetAmount() (Amount int) {
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
)

//
type CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the declined
	// transfer's currency. This will always be "USD" for a Real Time Payments
	// transfer.
	Currency *CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency `json:"currency"`
	// Why the transfer was declined.
	Reason *CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason `json:"reason"`
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
func (r *CardAuthorizationSimulationDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) GetAmount() (Amount int) {
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

//
type CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline struct {
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
	OriginatingDepositoryFinancialInstitutionIdQualifier *string `json:"originating_depository_financial_institution_id_qualifier"`
	//
	OriginatingDepositoryFinancialInstitutionId *string `json:"originating_depository_financial_institution_id"`
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
	ReceivingDepositoryFinancialInstitutionIdQualifier *string `json:"receiving_depository_financial_institution_id_qualifier"`
	//
	ReceivingDepositoryFinancialInstitutionId *string `json:"receiving_depository_financial_institution_id"`
	//
	ReceivingDepositoryFinancialInstitutionCountry *string `json:"receiving_depository_financial_institution_country"`
	//
	TraceNumber *string `json:"trace_number"`
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) GetAmount() (Amount int) {
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

func (r *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) GetForeignPaymentAmount() (ForeignPaymentAmount int) {
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

func (r *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) GetOriginatingDepositoryFinancialInstitutionIdQualifier() (OriginatingDepositoryFinancialInstitutionIdQualifier string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionIdQualifier != nil {
		OriginatingDepositoryFinancialInstitutionIdQualifier = *r.OriginatingDepositoryFinancialInstitutionIdQualifier
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) GetOriginatingDepositoryFinancialInstitutionId() (OriginatingDepositoryFinancialInstitutionId string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionId != nil {
		OriginatingDepositoryFinancialInstitutionId = *r.OriginatingDepositoryFinancialInstitutionId
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

func (r *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) GetReceivingDepositoryFinancialInstitutionIdQualifier() (ReceivingDepositoryFinancialInstitutionIdQualifier string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionIdQualifier != nil {
		ReceivingDepositoryFinancialInstitutionIdQualifier = *r.ReceivingDepositoryFinancialInstitutionIdQualifier
	}
	return
}

func (r *CardAuthorizationSimulationDeclinedTransactionSourceInternationalACHDecline) GetReceivingDepositoryFinancialInstitutionId() (ReceivingDepositoryFinancialInstitutionId string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionId != nil {
		ReceivingDepositoryFinancialInstitutionId = *r.ReceivingDepositoryFinancialInstitutionId
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

//
type CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency *CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDeclineCurrency `json:"currency"`
	//
	MerchantAcceptorId *string `json:"merchant_acceptor_id"`
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
func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDecline) GetAmount() (Amount int) {
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

func (r *CardAuthorizationSimulationDeclinedTransactionSourceCardRouteDecline) GetMerchantAcceptorId() (MerchantAcceptorId string) {
	if r != nil && r.MerchantAcceptorId != nil {
		MerchantAcceptorId = *r.MerchantAcceptorId
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
	Amount *int `json:"amount"`
	// The identifier of the Card to be authorized.
	CardId *string `json:"card_id,omitempty"`
	// The identifier of the Digital Wallet Token to be authorized.
	DigitalWalletTokenId *string `json:"digital_wallet_token_id,omitempty"`
}

// The authorization amount in cents.
func (r *SimulateAnAuthorizationOnACardParameters) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The identifier of the Card to be authorized.
func (r *SimulateAnAuthorizationOnACardParameters) GetCardId() (CardId string) {
	if r != nil && r.CardId != nil {
		CardId = *r.CardId
	}
	return
}

// The identifier of the Digital Wallet Token to be authorized.
func (r *SimulateAnAuthorizationOnACardParameters) GetDigitalWalletTokenId() (DigitalWalletTokenId string) {
	if r != nil && r.DigitalWalletTokenId != nil {
		DigitalWalletTokenId = *r.DigitalWalletTokenId
	}
	return
}

type SimulateSettlingACardAuthorizationParameters struct {
	// The identifier of the Card to create a settlement on.
	CardId *string `json:"card_id"`
	// The identifier of the Pending Transaction for the Card Authorization you wish to
	// settle.
	PendingTransactionId *string `json:"pending_transaction_id"`
	// The amount to be settled. This defaults to the amount of the Pending Transaction
	// being settled.
	Amount *int `json:"amount,omitempty"`
}

// The identifier of the Card to create a settlement on.
func (r *SimulateSettlingACardAuthorizationParameters) GetCardId() (CardId string) {
	if r != nil && r.CardId != nil {
		CardId = *r.CardId
	}
	return
}

// The identifier of the Pending Transaction for the Card Authorization you wish to
// settle.
func (r *SimulateSettlingACardAuthorizationParameters) GetPendingTransactionId() (PendingTransactionId string) {
	if r != nil && r.PendingTransactionId != nil {
		PendingTransactionId = *r.PendingTransactionId
	}
	return
}

// The amount to be settled. This defaults to the amount of the Pending Transaction
// being settled.
func (r *SimulateSettlingACardAuthorizationParameters) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

//
type InboundRealTimePaymentsTransferSimulationResult struct {
	// If the Real Time Payments Transfer attempt succeeds, this will contain the
	// resulting [Transaction](#transactions) object. The Transaction's `source` will
	// be of `category: inbound_real_time_payments_transfer_confirmation`.
	Transaction *InboundRealTimePaymentsTransferSimulationResultTransaction `json:"transaction"`
	// If the Real Time Payments Transfer attempt fails, this will contain the
	// resulting [Declined Transaction](#declined-transactions) object. The Declined
	// Transaction's `source` will be of
	// `category: inbound_real_time_payments_transfer_decline`.
	DeclinedTransaction *InboundRealTimePaymentsTransferSimulationResultDeclinedTransaction `json:"declined_transaction"`
	// A constant representing the object's type. For this resource it will always be
	// `inbound_real_time_payments_transfer_simulation_result`.
	Type *InboundRealTimePaymentsTransferSimulationResultType `json:"type"`
}

// If the Real Time Payments Transfer attempt succeeds, this will contain the
// resulting [Transaction](#transactions) object. The Transaction's `source` will
// be of `category: inbound_real_time_payments_transfer_confirmation`.
func (r *InboundRealTimePaymentsTransferSimulationResult) GetTransaction() (Transaction InboundRealTimePaymentsTransferSimulationResultTransaction) {
	if r != nil && r.Transaction != nil {
		Transaction = *r.Transaction
	}
	return
}

// If the Real Time Payments Transfer attempt fails, this will contain the
// resulting [Declined Transaction](#declined-transactions) object. The Declined
// Transaction's `source` will be of
// `category: inbound_real_time_payments_transfer_decline`.
func (r *InboundRealTimePaymentsTransferSimulationResult) GetDeclinedTransaction() (DeclinedTransaction InboundRealTimePaymentsTransferSimulationResultDeclinedTransaction) {
	if r != nil && r.DeclinedTransaction != nil {
		DeclinedTransaction = *r.DeclinedTransaction
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `inbound_real_time_payments_transfer_simulation_result`.
func (r *InboundRealTimePaymentsTransferSimulationResult) GetType() (Type InboundRealTimePaymentsTransferSimulationResultType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

//
type InboundRealTimePaymentsTransferSimulationResultTransaction struct {
	// The identifier for the Account the Transaction belongs to.
	AccountId *string `json:"account_id"`
	// The Transaction amount in the minor unit of its currency. For dollars, for
	// example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// Transaction's currency. This will match the currency on the Transcation's
	// Account.
	Currency *InboundRealTimePaymentsTransferSimulationResultTransactionCurrency `json:"currency"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the
	// Transaction occured.
	CreatedAt *string `json:"created_at"`
	// For a Transaction related to a transfer, this is the description you provide.
	// For a Transaction related to a payment, this is the description the vendor
	// provides.
	Description *string `json:"description"`
	// The Transaction identifier.
	Id *string `json:"id"`
	// The identifier for the route this Transaction came through. Routes are things
	// like cards and ACH details.
	RouteId *string `json:"route_id"`
	// The type of the route this Transaction came through.
	RouteType *string `json:"route_type"`
	// This is an object giving more details on the network-level event that caused the
	// Transaction. Note that for backwards compatibility reasons, additional
	// undocumented keys may appear in this object. These should be treated as
	// deprecated and will be removed in the future.
	Source *InboundRealTimePaymentsTransferSimulationResultTransactionSource `json:"source"`
	// A constant representing the object's type. For this resource it will always be
	// `transaction`.
	Type *InboundRealTimePaymentsTransferSimulationResultTransactionType `json:"type"`
}

// The identifier for the Account the Transaction belongs to.
func (r *InboundRealTimePaymentsTransferSimulationResultTransaction) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

// The Transaction amount in the minor unit of its currency. For dollars, for
// example, this is cents.
func (r *InboundRealTimePaymentsTransferSimulationResultTransaction) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// Transaction's currency. This will match the currency on the Transcation's
// Account.
func (r *InboundRealTimePaymentsTransferSimulationResultTransaction) GetCurrency() (Currency InboundRealTimePaymentsTransferSimulationResultTransactionCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the
// Transaction occured.
func (r *InboundRealTimePaymentsTransferSimulationResultTransaction) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// For a Transaction related to a transfer, this is the description you provide.
// For a Transaction related to a payment, this is the description the vendor
// provides.
func (r *InboundRealTimePaymentsTransferSimulationResultTransaction) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The Transaction identifier.
func (r *InboundRealTimePaymentsTransferSimulationResultTransaction) GetId() (Id string) {
	if r != nil && r.Id != nil {
		Id = *r.Id
	}
	return
}

// The identifier for the route this Transaction came through. Routes are things
// like cards and ACH details.
func (r *InboundRealTimePaymentsTransferSimulationResultTransaction) GetRouteId() (RouteId string) {
	if r != nil && r.RouteId != nil {
		RouteId = *r.RouteId
	}
	return
}

// The type of the route this Transaction came through.
func (r *InboundRealTimePaymentsTransferSimulationResultTransaction) GetRouteType() (RouteType string) {
	if r != nil && r.RouteType != nil {
		RouteType = *r.RouteType
	}
	return
}

// This is an object giving more details on the network-level event that caused the
// Transaction. Note that for backwards compatibility reasons, additional
// undocumented keys may appear in this object. These should be treated as
// deprecated and will be removed in the future.
func (r *InboundRealTimePaymentsTransferSimulationResultTransaction) GetSource() (Source InboundRealTimePaymentsTransferSimulationResultTransactionSource) {
	if r != nil && r.Source != nil {
		Source = *r.Source
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `transaction`.
func (r *InboundRealTimePaymentsTransferSimulationResultTransaction) GetType() (Type InboundRealTimePaymentsTransferSimulationResultTransactionType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

type InboundRealTimePaymentsTransferSimulationResultTransactionCurrency string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionCurrencyCad InboundRealTimePaymentsTransferSimulationResultTransactionCurrency = "CAD"
	InboundRealTimePaymentsTransferSimulationResultTransactionCurrencyChf InboundRealTimePaymentsTransferSimulationResultTransactionCurrency = "CHF"
	InboundRealTimePaymentsTransferSimulationResultTransactionCurrencyEur InboundRealTimePaymentsTransferSimulationResultTransactionCurrency = "EUR"
	InboundRealTimePaymentsTransferSimulationResultTransactionCurrencyGbp InboundRealTimePaymentsTransferSimulationResultTransactionCurrency = "GBP"
	InboundRealTimePaymentsTransferSimulationResultTransactionCurrencyJpy InboundRealTimePaymentsTransferSimulationResultTransactionCurrency = "JPY"
	InboundRealTimePaymentsTransferSimulationResultTransactionCurrencyUsd InboundRealTimePaymentsTransferSimulationResultTransactionCurrency = "USD"
)

//
type InboundRealTimePaymentsTransferSimulationResultTransactionSource struct {
	// The type of transaction that took place. We may add additional possible values
	// for this enum over time; your application should be able to handle such
	// additions gracefully.
	Category *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory `json:"category"`
	// A Account Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `account_transfer_intention`.
	AccountTransferIntention *InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntention `json:"account_transfer_intention"`
	// A ACH Check Conversion Return object. This field will be present in the JSON
	// response if and only if `category` is equal to `ach_check_conversion_return`.
	ACHCheckConversionReturn *InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHCheckConversionReturn `json:"ach_check_conversion_return"`
	// A ACH Check Conversion object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_check_conversion`.
	ACHCheckConversion *InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHCheckConversion `json:"ach_check_conversion"`
	// A ACH Transfer Intention object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_transfer_intention`.
	ACHTransferIntention *InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferIntention `json:"ach_transfer_intention"`
	// A ACH Transfer Rejection object. This field will be present in the JSON response
	// if and only if `category` is equal to `ach_transfer_rejection`.
	ACHTransferRejection *InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferRejection `json:"ach_transfer_rejection"`
	// A ACH Transfer Return object. This field will be present in the JSON response if
	// and only if `category` is equal to `ach_transfer_return`.
	ACHTransferReturn *InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturn `json:"ach_transfer_return"`
	// A Card Dispute Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_dispute_acceptance`.
	CardDisputeAcceptance *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardDisputeAcceptance `json:"card_dispute_acceptance"`
	// A Card Refund object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_refund`.
	CardRefund *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefund `json:"card_refund"`
	// A Card Settlement object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_settlement`.
	CardSettlement *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlement `json:"card_settlement"`
	// A Check Deposit Acceptance object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_deposit_acceptance`.
	CheckDepositAcceptance *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptance `json:"check_deposit_acceptance"`
	// A Check Deposit Return object. This field will be present in the JSON response
	// if and only if `category` is equal to `check_deposit_return`.
	CheckDepositReturn *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturn `json:"check_deposit_return"`
	// A Check Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_intention`.
	CheckTransferIntention *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntention `json:"check_transfer_intention"`
	// A Check Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `check_transfer_rejection`.
	CheckTransferRejection *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferRejection `json:"check_transfer_rejection"`
	// A Check Transfer Stop Payment Request object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `check_transfer_stop_payment_request`.
	CheckTransferStopPaymentRequest *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequest `json:"check_transfer_stop_payment_request"`
	// A Dispute Resolution object. This field will be present in the JSON response if
	// and only if `category` is equal to `dispute_resolution`.
	DisputeResolution *InboundRealTimePaymentsTransferSimulationResultTransactionSourceDisputeResolution `json:"dispute_resolution"`
	// A Empyreal Cash Deposit object. This field will be present in the JSON response
	// if and only if `category` is equal to `empyreal_cash_deposit`.
	EmpyrealCashDeposit *InboundRealTimePaymentsTransferSimulationResultTransactionSourceEmpyrealCashDeposit `json:"empyreal_cash_deposit"`
	// A Inbound ACH Transfer object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_ach_transfer`.
	InboundACHTransfer *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundACHTransfer `json:"inbound_ach_transfer"`
	// A Inbound Check object. This field will be present in the JSON response if and
	// only if `category` is equal to `inbound_check`.
	InboundCheck *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheck `json:"inbound_check"`
	// A Inbound International ACH Transfer object. This field will be present in the
	// JSON response if and only if `category` is equal to
	// `inbound_international_ach_transfer`.
	InboundInternationalACHTransfer *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer `json:"inbound_international_ach_transfer"`
	// A Inbound Real Time Payments Transfer Confirmation object. This field will be
	// present in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_confirmation`.
	InboundRealTimePaymentsTransferConfirmation *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmation `json:"inbound_real_time_payments_transfer_confirmation"`
	// A Inbound Wire Drawdown Payment Reversal object. This field will be present in
	// the JSON response if and only if `category` is equal to
	// `inbound_wire_drawdown_payment_reversal`.
	InboundWireDrawdownPaymentReversal *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPaymentReversal `json:"inbound_wire_drawdown_payment_reversal"`
	// A Inbound Wire Drawdown Payment object. This field will be present in the JSON
	// response if and only if `category` is equal to `inbound_wire_drawdown_payment`.
	InboundWireDrawdownPayment *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPayment `json:"inbound_wire_drawdown_payment"`
	// A Inbound Wire Reversal object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_wire_reversal`.
	InboundWireReversal *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireReversal `json:"inbound_wire_reversal"`
	// A Inbound Wire Transfer object. This field will be present in the JSON response
	// if and only if `category` is equal to `inbound_wire_transfer`.
	InboundWireTransfer *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireTransfer `json:"inbound_wire_transfer"`
	// A Internal Source object. This field will be present in the JSON response if and
	// only if `category` is equal to `internal_source`.
	InternalSource *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSource `json:"internal_source"`
	// A Deprecated Card Refund object. This field will be present in the JSON response
	// if and only if `category` is equal to `card_route_refund`.
	CardRouteRefund *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefund `json:"card_route_refund"`
	// A Deprecated Card Settlement object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_route_settlement`.
	CardRouteSettlement *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlement `json:"card_route_settlement"`
	// A Sample Funds object. This field will be present in the JSON response if and
	// only if `category` is equal to `sample_funds`.
	SampleFunds *InboundRealTimePaymentsTransferSimulationResultTransactionSourceSampleFunds `json:"sample_funds"`
	// A Wire Drawdown Payment Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to
	// `wire_drawdown_payment_intention`.
	WireDrawdownPaymentIntention *InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireDrawdownPaymentIntention `json:"wire_drawdown_payment_intention"`
	// A Wire Drawdown Payment Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to
	// `wire_drawdown_payment_rejection`.
	WireDrawdownPaymentRejection *InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireDrawdownPaymentRejection `json:"wire_drawdown_payment_rejection"`
	// A Wire Transfer Intention object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_intention`.
	WireTransferIntention *InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferIntention `json:"wire_transfer_intention"`
	// A Wire Transfer Rejection object. This field will be present in the JSON
	// response if and only if `category` is equal to `wire_transfer_rejection`.
	WireTransferRejection *InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferRejection `json:"wire_transfer_rejection"`
}

// The type of transaction that took place. We may add additional possible values
// for this enum over time; your application should be able to handle such
// additions gracefully.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSource) GetCategory() (Category InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory) {
	if r != nil && r.Category != nil {
		Category = *r.Category
	}
	return
}

// A Account Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `account_transfer_intention`.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSource) GetAccountTransferIntention() (AccountTransferIntention InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntention) {
	if r != nil && r.AccountTransferIntention != nil {
		AccountTransferIntention = *r.AccountTransferIntention
	}
	return
}

// A ACH Check Conversion Return object. This field will be present in the JSON
// response if and only if `category` is equal to `ach_check_conversion_return`.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSource) GetACHCheckConversionReturn() (ACHCheckConversionReturn InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHCheckConversionReturn) {
	if r != nil && r.ACHCheckConversionReturn != nil {
		ACHCheckConversionReturn = *r.ACHCheckConversionReturn
	}
	return
}

// A ACH Check Conversion object. This field will be present in the JSON response
// if and only if `category` is equal to `ach_check_conversion`.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSource) GetACHCheckConversion() (ACHCheckConversion InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHCheckConversion) {
	if r != nil && r.ACHCheckConversion != nil {
		ACHCheckConversion = *r.ACHCheckConversion
	}
	return
}

// A ACH Transfer Intention object. This field will be present in the JSON response
// if and only if `category` is equal to `ach_transfer_intention`.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSource) GetACHTransferIntention() (ACHTransferIntention InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferIntention) {
	if r != nil && r.ACHTransferIntention != nil {
		ACHTransferIntention = *r.ACHTransferIntention
	}
	return
}

// A ACH Transfer Rejection object. This field will be present in the JSON response
// if and only if `category` is equal to `ach_transfer_rejection`.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSource) GetACHTransferRejection() (ACHTransferRejection InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferRejection) {
	if r != nil && r.ACHTransferRejection != nil {
		ACHTransferRejection = *r.ACHTransferRejection
	}
	return
}

// A ACH Transfer Return object. This field will be present in the JSON response if
// and only if `category` is equal to `ach_transfer_return`.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSource) GetACHTransferReturn() (ACHTransferReturn InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturn) {
	if r != nil && r.ACHTransferReturn != nil {
		ACHTransferReturn = *r.ACHTransferReturn
	}
	return
}

// A Card Dispute Acceptance object. This field will be present in the JSON
// response if and only if `category` is equal to `card_dispute_acceptance`.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSource) GetCardDisputeAcceptance() (CardDisputeAcceptance InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardDisputeAcceptance) {
	if r != nil && r.CardDisputeAcceptance != nil {
		CardDisputeAcceptance = *r.CardDisputeAcceptance
	}
	return
}

// A Card Refund object. This field will be present in the JSON response if and
// only if `category` is equal to `card_refund`.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSource) GetCardRefund() (CardRefund InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefund) {
	if r != nil && r.CardRefund != nil {
		CardRefund = *r.CardRefund
	}
	return
}

// A Card Settlement object. This field will be present in the JSON response if and
// only if `category` is equal to `card_settlement`.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSource) GetCardSettlement() (CardSettlement InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlement) {
	if r != nil && r.CardSettlement != nil {
		CardSettlement = *r.CardSettlement
	}
	return
}

// A Check Deposit Acceptance object. This field will be present in the JSON
// response if and only if `category` is equal to `check_deposit_acceptance`.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSource) GetCheckDepositAcceptance() (CheckDepositAcceptance InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptance) {
	if r != nil && r.CheckDepositAcceptance != nil {
		CheckDepositAcceptance = *r.CheckDepositAcceptance
	}
	return
}

// A Check Deposit Return object. This field will be present in the JSON response
// if and only if `category` is equal to `check_deposit_return`.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSource) GetCheckDepositReturn() (CheckDepositReturn InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturn) {
	if r != nil && r.CheckDepositReturn != nil {
		CheckDepositReturn = *r.CheckDepositReturn
	}
	return
}

// A Check Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `check_transfer_intention`.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSource) GetCheckTransferIntention() (CheckTransferIntention InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntention) {
	if r != nil && r.CheckTransferIntention != nil {
		CheckTransferIntention = *r.CheckTransferIntention
	}
	return
}

// A Check Transfer Rejection object. This field will be present in the JSON
// response if and only if `category` is equal to `check_transfer_rejection`.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSource) GetCheckTransferRejection() (CheckTransferRejection InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferRejection) {
	if r != nil && r.CheckTransferRejection != nil {
		CheckTransferRejection = *r.CheckTransferRejection
	}
	return
}

// A Check Transfer Stop Payment Request object. This field will be present in the
// JSON response if and only if `category` is equal to
// `check_transfer_stop_payment_request`.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSource) GetCheckTransferStopPaymentRequest() (CheckTransferStopPaymentRequest InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequest) {
	if r != nil && r.CheckTransferStopPaymentRequest != nil {
		CheckTransferStopPaymentRequest = *r.CheckTransferStopPaymentRequest
	}
	return
}

// A Dispute Resolution object. This field will be present in the JSON response if
// and only if `category` is equal to `dispute_resolution`.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSource) GetDisputeResolution() (DisputeResolution InboundRealTimePaymentsTransferSimulationResultTransactionSourceDisputeResolution) {
	if r != nil && r.DisputeResolution != nil {
		DisputeResolution = *r.DisputeResolution
	}
	return
}

// A Empyreal Cash Deposit object. This field will be present in the JSON response
// if and only if `category` is equal to `empyreal_cash_deposit`.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSource) GetEmpyrealCashDeposit() (EmpyrealCashDeposit InboundRealTimePaymentsTransferSimulationResultTransactionSourceEmpyrealCashDeposit) {
	if r != nil && r.EmpyrealCashDeposit != nil {
		EmpyrealCashDeposit = *r.EmpyrealCashDeposit
	}
	return
}

// A Inbound ACH Transfer object. This field will be present in the JSON response
// if and only if `category` is equal to `inbound_ach_transfer`.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSource) GetInboundACHTransfer() (InboundACHTransfer InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundACHTransfer) {
	if r != nil && r.InboundACHTransfer != nil {
		InboundACHTransfer = *r.InboundACHTransfer
	}
	return
}

// A Inbound Check object. This field will be present in the JSON response if and
// only if `category` is equal to `inbound_check`.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSource) GetInboundCheck() (InboundCheck InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheck) {
	if r != nil && r.InboundCheck != nil {
		InboundCheck = *r.InboundCheck
	}
	return
}

// A Inbound International ACH Transfer object. This field will be present in the
// JSON response if and only if `category` is equal to
// `inbound_international_ach_transfer`.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSource) GetInboundInternationalACHTransfer() (InboundInternationalACHTransfer InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer) {
	if r != nil && r.InboundInternationalACHTransfer != nil {
		InboundInternationalACHTransfer = *r.InboundInternationalACHTransfer
	}
	return
}

// A Inbound Real Time Payments Transfer Confirmation object. This field will be
// present in the JSON response if and only if `category` is equal to
// `inbound_real_time_payments_transfer_confirmation`.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSource) GetInboundRealTimePaymentsTransferConfirmation() (InboundRealTimePaymentsTransferConfirmation InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmation) {
	if r != nil && r.InboundRealTimePaymentsTransferConfirmation != nil {
		InboundRealTimePaymentsTransferConfirmation = *r.InboundRealTimePaymentsTransferConfirmation
	}
	return
}

// A Inbound Wire Drawdown Payment Reversal object. This field will be present in
// the JSON response if and only if `category` is equal to
// `inbound_wire_drawdown_payment_reversal`.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSource) GetInboundWireDrawdownPaymentReversal() (InboundWireDrawdownPaymentReversal InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPaymentReversal) {
	if r != nil && r.InboundWireDrawdownPaymentReversal != nil {
		InboundWireDrawdownPaymentReversal = *r.InboundWireDrawdownPaymentReversal
	}
	return
}

// A Inbound Wire Drawdown Payment object. This field will be present in the JSON
// response if and only if `category` is equal to `inbound_wire_drawdown_payment`.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSource) GetInboundWireDrawdownPayment() (InboundWireDrawdownPayment InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPayment) {
	if r != nil && r.InboundWireDrawdownPayment != nil {
		InboundWireDrawdownPayment = *r.InboundWireDrawdownPayment
	}
	return
}

// A Inbound Wire Reversal object. This field will be present in the JSON response
// if and only if `category` is equal to `inbound_wire_reversal`.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSource) GetInboundWireReversal() (InboundWireReversal InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireReversal) {
	if r != nil && r.InboundWireReversal != nil {
		InboundWireReversal = *r.InboundWireReversal
	}
	return
}

// A Inbound Wire Transfer object. This field will be present in the JSON response
// if and only if `category` is equal to `inbound_wire_transfer`.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSource) GetInboundWireTransfer() (InboundWireTransfer InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireTransfer) {
	if r != nil && r.InboundWireTransfer != nil {
		InboundWireTransfer = *r.InboundWireTransfer
	}
	return
}

// A Internal Source object. This field will be present in the JSON response if and
// only if `category` is equal to `internal_source`.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSource) GetInternalSource() (InternalSource InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSource) {
	if r != nil && r.InternalSource != nil {
		InternalSource = *r.InternalSource
	}
	return
}

// A Deprecated Card Refund object. This field will be present in the JSON response
// if and only if `category` is equal to `card_route_refund`.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSource) GetCardRouteRefund() (CardRouteRefund InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefund) {
	if r != nil && r.CardRouteRefund != nil {
		CardRouteRefund = *r.CardRouteRefund
	}
	return
}

// A Deprecated Card Settlement object. This field will be present in the JSON
// response if and only if `category` is equal to `card_route_settlement`.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSource) GetCardRouteSettlement() (CardRouteSettlement InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlement) {
	if r != nil && r.CardRouteSettlement != nil {
		CardRouteSettlement = *r.CardRouteSettlement
	}
	return
}

// A Sample Funds object. This field will be present in the JSON response if and
// only if `category` is equal to `sample_funds`.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSource) GetSampleFunds() (SampleFunds InboundRealTimePaymentsTransferSimulationResultTransactionSourceSampleFunds) {
	if r != nil && r.SampleFunds != nil {
		SampleFunds = *r.SampleFunds
	}
	return
}

// A Wire Drawdown Payment Intention object. This field will be present in the JSON
// response if and only if `category` is equal to
// `wire_drawdown_payment_intention`.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSource) GetWireDrawdownPaymentIntention() (WireDrawdownPaymentIntention InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireDrawdownPaymentIntention) {
	if r != nil && r.WireDrawdownPaymentIntention != nil {
		WireDrawdownPaymentIntention = *r.WireDrawdownPaymentIntention
	}
	return
}

// A Wire Drawdown Payment Rejection object. This field will be present in the JSON
// response if and only if `category` is equal to
// `wire_drawdown_payment_rejection`.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSource) GetWireDrawdownPaymentRejection() (WireDrawdownPaymentRejection InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireDrawdownPaymentRejection) {
	if r != nil && r.WireDrawdownPaymentRejection != nil {
		WireDrawdownPaymentRejection = *r.WireDrawdownPaymentRejection
	}
	return
}

// A Wire Transfer Intention object. This field will be present in the JSON
// response if and only if `category` is equal to `wire_transfer_intention`.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSource) GetWireTransferIntention() (WireTransferIntention InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferIntention) {
	if r != nil && r.WireTransferIntention != nil {
		WireTransferIntention = *r.WireTransferIntention
	}
	return
}

// A Wire Transfer Rejection object. This field will be present in the JSON
// response if and only if `category` is equal to `wire_transfer_rejection`.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSource) GetWireTransferRejection() (WireTransferRejection InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferRejection) {
	if r != nil && r.WireTransferRejection != nil {
		WireTransferRejection = *r.WireTransferRejection
	}
	return
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryAccountTransferIntention                    InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "account_transfer_intention"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryACHCheckConversionReturn                    InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "ach_check_conversion_return"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryACHCheckConversion                          InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "ach_check_conversion"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryACHTransferIntention                        InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "ach_transfer_intention"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryACHTransferRejection                        InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "ach_transfer_rejection"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryACHTransferReturn                           InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "ach_transfer_return"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryCardDisputeAcceptance                       InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "card_dispute_acceptance"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryCardRefund                                  InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "card_refund"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryCardSettlement                              InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "card_settlement"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryCheckDepositAcceptance                      InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "check_deposit_acceptance"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryCheckDepositReturn                          InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "check_deposit_return"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryCheckTransferIntention                      InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "check_transfer_intention"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryCheckTransferRejection                      InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "check_transfer_rejection"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryCheckTransferStopPaymentRequest             InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "check_transfer_stop_payment_request"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryDisputeResolution                           InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "dispute_resolution"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryEmpyrealCashDeposit                         InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "empyreal_cash_deposit"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryInboundACHTransfer                          InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "inbound_ach_transfer"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryInboundCheck                                InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "inbound_check"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryInboundInternationalACHTransfer             InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "inbound_international_ach_transfer"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryInboundRealTimePaymentsTransferConfirmation InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "inbound_real_time_payments_transfer_confirmation"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryInboundWireDrawdownPaymentReversal          InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "inbound_wire_drawdown_payment_reversal"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryInboundWireDrawdownPayment                  InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "inbound_wire_drawdown_payment"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryInboundWireReversal                         InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "inbound_wire_reversal"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryInboundWireTransfer                         InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "inbound_wire_transfer"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryInternalSource                              InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "internal_source"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryCardRouteRefund                             InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "card_route_refund"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryCardRouteSettlement                         InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "card_route_settlement"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryRealTimePaymentsTransferAcknowledgement     InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "real_time_payments_transfer_acknowledgement"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategorySampleFunds                                 InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "sample_funds"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryWireDrawdownPaymentIntention                InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "wire_drawdown_payment_intention"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryWireDrawdownPaymentRejection                InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "wire_drawdown_payment_rejection"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryWireTransferIntention                       InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "wire_transfer_intention"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryWireTransferRejection                       InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "wire_transfer_rejection"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategoryOther                                       InboundRealTimePaymentsTransferSimulationResultTransactionSourceCategory = "other"
)

//
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntention struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency *InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionCurrency `json:"currency"`
	// The description you chose to give the transfer.
	Description *string `json:"description"`
	// The identifier of the Account to where the Account Transfer was sent.
	DestinationAccountId *string `json:"destination_account_id"`
	// The identifier of the Account from where the Account Transfer was sent.
	SourceAccountId *string `json:"source_account_id"`
	// The identifier of the Account Transfer that led to this Pending Transaction.
	TransferId *string `json:"transfer_id"`
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntention) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
// account currency.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntention) GetCurrency() (Currency InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The description you chose to give the transfer.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntention) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The identifier of the Account to where the Account Transfer was sent.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntention) GetDestinationAccountId() (DestinationAccountId string) {
	if r != nil && r.DestinationAccountId != nil {
		DestinationAccountId = *r.DestinationAccountId
	}
	return
}

// The identifier of the Account from where the Account Transfer was sent.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntention) GetSourceAccountId() (SourceAccountId string) {
	if r != nil && r.SourceAccountId != nil {
		SourceAccountId = *r.SourceAccountId
	}
	return
}

// The identifier of the Account Transfer that led to this Pending Transaction.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntention) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionCurrency string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionCurrencyCad InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionCurrency = "CAD"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionCurrencyChf InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionCurrency = "CHF"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionCurrencyEur InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionCurrency = "EUR"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionCurrencyGbp InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionCurrency = "GBP"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionCurrencyJpy InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionCurrency = "JPY"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionCurrencyUsd InboundRealTimePaymentsTransferSimulationResultTransactionSourceAccountTransferIntentionCurrency = "USD"
)

//
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHCheckConversionReturn struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int `json:"amount"`
	// Why the transfer was returned.
	ReturnReasonCode *string `json:"return_reason_code"`
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHCheckConversionReturn) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// Why the transfer was returned.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHCheckConversionReturn) GetReturnReasonCode() (ReturnReasonCode string) {
	if r != nil && r.ReturnReasonCode != nil {
		ReturnReasonCode = *r.ReturnReasonCode
	}
	return
}

//
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHCheckConversion struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int `json:"amount"`
	// The identifier of the File containing an image of the returned check.
	FileId *string `json:"file_id"`
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHCheckConversion) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The identifier of the File containing an image of the returned check.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHCheckConversion) GetFileId() (FileId string) {
	if r != nil && r.FileId != nil {
		FileId = *r.FileId
	}
	return
}

//
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferIntention struct {
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
	TransferId *string `json:"transfer_id"`
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferIntention) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferIntention) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferIntention) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferIntention) GetStatementDescriptor() (StatementDescriptor string) {
	if r != nil && r.StatementDescriptor != nil {
		StatementDescriptor = *r.StatementDescriptor
	}
	return
}

// The identifier of the ACH Transfer that led to this Transaction.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferIntention) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}

//
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferRejection struct {
	// The identifier of the ACH Transfer that led to this Transaction.
	TransferId *string `json:"transfer_id"`
}

// The identifier of the ACH Transfer that led to this Transaction.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferRejection) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}

//
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturn struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt *string `json:"created_at"`
	// Why the ACH Transfer was returned.
	ReturnReasonCode *InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode `json:"return_reason_code"`
	// The identifier of the ACH Transfer associated with this return.
	TransferId *string `json:"transfer_id"`
	// The identifier of the Tranasaction associated with this return.
	TransactionId *string `json:"transaction_id"`
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the transfer was created.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturn) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// Why the ACH Transfer was returned.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturn) GetReturnReasonCode() (ReturnReasonCode InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode) {
	if r != nil && r.ReturnReasonCode != nil {
		ReturnReasonCode = *r.ReturnReasonCode
	}
	return
}

// The identifier of the ACH Transfer associated with this return.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturn) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}

// The identifier of the Tranasaction associated with this return.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturn) GetTransactionId() (TransactionId string) {
	if r != nil && r.TransactionId != nil {
		TransactionId = *r.TransactionId
	}
	return
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeInsufficientFund                                          InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "insufficient_fund"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeNoAccount                                                 InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "no_account"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeAccountClosed                                             InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "account_closed"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeInvalidAccountNumberStructure                             InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "invalid_account_number_structure"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeAccountFrozenEntryReturnedPerOfacInstruction              InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "account_frozen_entry_returned_per_ofac_instruction"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeCreditEntryRefusedByReceiver                              InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "credit_entry_refused_by_receiver"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeUnauthorizedDebitToConsumerAccountUsingCorporateSecCode   InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "unauthorized_debit_to_consumer_account_using_corporate_sec_code"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeCorporateCustomerAdvisedNotAuthorized                     InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "corporate_customer_advised_not_authorized"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodePaymentStopped                                            InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "payment_stopped"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeNonTransactionAccount                                     InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "non_transaction_account"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeUncollectedFunds                                          InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "uncollected_funds"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeRoutingNumberCheckDigitError                              InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "routing_number_check_digit_error"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "customer_advised_unauthorized_improper_ineligible_or_incomplete"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeAmountFieldError                                          InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "amount_field_error"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeAuthorizationRevokedByCustomer                            InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "authorization_revoked_by_customer"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeInvalidACHRoutingNumber                                   InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "invalid_ach_routing_number"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeFileRecordEditCriteria                                    InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "file_record_edit_criteria"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeEnrInvalidIndividualName                                  InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "enr_invalid_individual_name"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeReturnedPerOdfiRequest                                    InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "returned_per_odfi_request"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeAddendaError                                              InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "addenda_error"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeLimitedParticipationDfi                                   InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "limited_participation_dfi"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCodeOther                                                     InboundRealTimePaymentsTransferSimulationResultTransactionSourceACHTransferReturnReturnReasonCode = "other"
)

//
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardDisputeAcceptance struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card Dispute was accepted.
	AcceptedAt *string `json:"accepted_at"`
	// The identifier of the Card Dispute that was accepted.
	CardDisputeId *string `json:"card_dispute_id"`
	// The identifier of the Transaction that was created to return the disputed funds
	// to your account.
	TransactionId *string `json:"transaction_id"`
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the Card Dispute was accepted.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardDisputeAcceptance) GetAcceptedAt() (AcceptedAt string) {
	if r != nil && r.AcceptedAt != nil {
		AcceptedAt = *r.AcceptedAt
	}
	return
}

// The identifier of the Card Dispute that was accepted.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardDisputeAcceptance) GetCardDisputeId() (CardDisputeId string) {
	if r != nil && r.CardDisputeId != nil {
		CardDisputeId = *r.CardDisputeId
	}
	return
}

// The identifier of the Transaction that was created to return the disputed funds
// to your account.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardDisputeAcceptance) GetTransactionId() (TransactionId string) {
	if r != nil && r.TransactionId != nil {
		TransactionId = *r.TransactionId
	}
	return
}

//
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefund struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundCurrency `json:"currency"`
	// The identifier for the Transaction this refunds, if any.
	CardSettlementTransactionId *string `json:"card_settlement_transaction_id"`
	// A constant representing the object's type. For this resource it will always be
	// `card_refund`.
	Type *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundType `json:"type"`
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefund) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefund) GetCurrency() (Currency InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The identifier for the Transaction this refunds, if any.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefund) GetCardSettlementTransactionId() (CardSettlementTransactionId string) {
	if r != nil && r.CardSettlementTransactionId != nil {
		CardSettlementTransactionId = *r.CardSettlementTransactionId
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `card_refund`.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefund) GetType() (Type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundCurrency string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundCurrencyCad InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundCurrency = "CAD"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundCurrencyChf InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundCurrency = "CHF"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundCurrencyEur InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundCurrency = "EUR"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundCurrencyGbp InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundCurrency = "GBP"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundCurrencyJpy InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundCurrency = "JPY"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundCurrencyUsd InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundCurrency = "USD"
)

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundType string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundTypeCardRefund InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRefundType = "card_refund"
)

//
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlement struct {
	// The pending amount in the minor unit of the transaction's currency. For dollars,
	// for example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementCurrency `json:"currency"`
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
	PendingTransactionId *string `json:"pending_transaction_id"`
	// A constant representing the object's type. For this resource it will always be
	// `card_settlement`.
	Type *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementType `json:"type"`
}

// The pending amount in the minor unit of the transaction's currency. For dollars,
// for example, this is cents.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlement) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlement) GetCurrency() (Currency InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlement) GetMerchantCity() (MerchantCity string) {
	if r != nil && r.MerchantCity != nil {
		MerchantCity = *r.MerchantCity
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlement) GetMerchantCountry() (MerchantCountry string) {
	if r != nil && r.MerchantCountry != nil {
		MerchantCountry = *r.MerchantCountry
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlement) GetMerchantName() (MerchantName string) {
	if r != nil && r.MerchantName != nil {
		MerchantName = *r.MerchantName
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlement) GetMerchantCategoryCode() (MerchantCategoryCode string) {
	if r != nil && r.MerchantCategoryCode != nil {
		MerchantCategoryCode = *r.MerchantCategoryCode
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlement) GetMerchantState() (MerchantState string) {
	if r != nil && r.MerchantState != nil {
		MerchantState = *r.MerchantState
	}
	return
}

// The identifier of the Pending Transaction associated with this Transaction.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlement) GetPendingTransactionId() (PendingTransactionId string) {
	if r != nil && r.PendingTransactionId != nil {
		PendingTransactionId = *r.PendingTransactionId
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `card_settlement`.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlement) GetType() (Type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementCurrency string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementCurrencyCad InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementCurrency = "CAD"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementCurrencyChf InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementCurrency = "CHF"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementCurrencyEur InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementCurrency = "EUR"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementCurrencyGbp InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementCurrency = "GBP"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementCurrencyJpy InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementCurrency = "JPY"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementCurrencyUsd InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementCurrency = "USD"
)

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementType string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementTypeCardSettlement InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardSettlementType = "card_settlement"
)

//
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptance struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceCurrency `json:"currency"`
	// The ID of the Check Deposit that led to the Transaction.
	CheckDepositId *string `json:"check_deposit_id"`
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptance) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptance) GetCurrency() (Currency InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The ID of the Check Deposit that led to the Transaction.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptance) GetCheckDepositId() (CheckDepositId string) {
	if r != nil && r.CheckDepositId != nil {
		CheckDepositId = *r.CheckDepositId
	}
	return
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceCurrency string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceCurrencyCad InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceCurrency = "CAD"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceCurrencyChf InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceCurrency = "CHF"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceCurrencyEur InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceCurrency = "EUR"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceCurrencyGbp InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceCurrency = "GBP"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceCurrencyJpy InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceCurrency = "JPY"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceCurrencyUsd InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositAcceptanceCurrency = "USD"
)

//
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturn struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check deposit was returned.
	ReturnedAt *string `json:"returned_at"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnCurrency `json:"currency"`
	// The identifier of the Check Deposit that was returned.
	CheckDepositId *string `json:"check_deposit_id"`
	// The identifier of the transaction that reversed the original check deposit
	// transaction.
	TransactionId *string `json:"transaction_id"`
	//
	ReturnReason *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReason `json:"return_reason"`
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturn) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the check deposit was returned.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturn) GetReturnedAt() (ReturnedAt string) {
	if r != nil && r.ReturnedAt != nil {
		ReturnedAt = *r.ReturnedAt
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturn) GetCurrency() (Currency InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The identifier of the Check Deposit that was returned.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturn) GetCheckDepositId() (CheckDepositId string) {
	if r != nil && r.CheckDepositId != nil {
		CheckDepositId = *r.CheckDepositId
	}
	return
}

// The identifier of the transaction that reversed the original check deposit
// transaction.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturn) GetTransactionId() (TransactionId string) {
	if r != nil && r.TransactionId != nil {
		TransactionId = *r.TransactionId
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturn) GetReturnReason() (ReturnReason InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReason) {
	if r != nil && r.ReturnReason != nil {
		ReturnReason = *r.ReturnReason
	}
	return
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnCurrency string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnCurrencyCad InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnCurrency = "CAD"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnCurrencyChf InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnCurrency = "CHF"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnCurrencyEur InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnCurrency = "EUR"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnCurrencyGbp InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnCurrency = "GBP"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnCurrencyJpy InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnCurrency = "JPY"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnCurrencyUsd InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnCurrency = "USD"
)

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReason string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReasonACHConversionNotSupported InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReason = "ach_conversion_not_supported"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReasonDuplicateSubmission       InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReason = "duplicate_submission"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReasonInsufficientFunds         InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReason = "insufficient_funds"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReasonNoAccount                 InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReason = "no_account"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReasonNotAuthorized             InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReason = "not_authorized"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReasonStaleDated                InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReason = "stale_dated"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReasonStopPayment               InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReason = "stop_payment"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReasonUnknownReason             InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReason = "unknown_reason"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReasonUnmatchedDetails          InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReason = "unmatched_details"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReasonUnreadableImage           InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckDepositReturnReturnReason = "unreadable_image"
)

//
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntention struct {
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
	Currency *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionCurrency `json:"currency"`
	// The name that will be printed on the check.
	RecipientName *string `json:"recipient_name"`
	// The identifier of the Check Transfer with which this is associated.
	TransferId *string `json:"transfer_id"`
}

// The street address of the check's destination.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntention) GetAddressLine1() (AddressLine1 string) {
	if r != nil && r.AddressLine1 != nil {
		AddressLine1 = *r.AddressLine1
	}
	return
}

// The second line of the address of the check's destination.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntention) GetAddressLine2() (AddressLine2 string) {
	if r != nil && r.AddressLine2 != nil {
		AddressLine2 = *r.AddressLine2
	}
	return
}

// The city of the check's destination.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntention) GetAddressCity() (AddressCity string) {
	if r != nil && r.AddressCity != nil {
		AddressCity = *r.AddressCity
	}
	return
}

// The state of the check's destination.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntention) GetAddressState() (AddressState string) {
	if r != nil && r.AddressState != nil {
		AddressState = *r.AddressState
	}
	return
}

// The postal code of the check's destination.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntention) GetAddressZip() (AddressZip string) {
	if r != nil && r.AddressZip != nil {
		AddressZip = *r.AddressZip
	}
	return
}

// The transfer amount in USD cents.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntention) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
// currency.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntention) GetCurrency() (Currency InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The name that will be printed on the check.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntention) GetRecipientName() (RecipientName string) {
	if r != nil && r.RecipientName != nil {
		RecipientName = *r.RecipientName
	}
	return
}

// The identifier of the Check Transfer with which this is associated.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntention) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionCurrency string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionCurrencyCad InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionCurrency = "CAD"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionCurrencyChf InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionCurrency = "CHF"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionCurrencyEur InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionCurrency = "EUR"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionCurrencyGbp InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionCurrency = "GBP"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionCurrencyJpy InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionCurrency = "JPY"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionCurrencyUsd InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferIntentionCurrency = "USD"
)

//
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferRejection struct {
	// The identifier of the Check Transfer that led to this Transaction.
	TransferId *string `json:"transfer_id"`
}

// The identifier of the Check Transfer that led to this Transaction.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferRejection) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}

//
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequest struct {
	// The ID of the check transfer that was stopped.
	TransferId *string `json:"transfer_id"`
	// The transaction ID of the corresponding credit transaction.
	TransactionId *string `json:"transaction_id"`
	// The time the stop-payment was requested.
	RequestedAt *string `json:"requested_at"`
	// A constant representing the object's type. For this resource it will always be
	// `check_transfer_stop_payment_request`.
	Type *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequestType `json:"type"`
}

// The ID of the check transfer that was stopped.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequest) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}

// The transaction ID of the corresponding credit transaction.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequest) GetTransactionId() (TransactionId string) {
	if r != nil && r.TransactionId != nil {
		TransactionId = *r.TransactionId
	}
	return
}

// The time the stop-payment was requested.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequest) GetRequestedAt() (RequestedAt string) {
	if r != nil && r.RequestedAt != nil {
		RequestedAt = *r.RequestedAt
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `check_transfer_stop_payment_request`.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequest) GetType() (Type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequestType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequestType string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequestTypeCheckTransferStopPaymentRequest InboundRealTimePaymentsTransferSimulationResultTransactionSourceCheckTransferStopPaymentRequestType = "check_transfer_stop_payment_request"
)

//
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceDisputeResolution struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *InboundRealTimePaymentsTransferSimulationResultTransactionSourceDisputeResolutionCurrency `json:"currency"`
	// The identifier of the Transaction that was disputed.
	DisputedTransactionId *string `json:"disputed_transaction_id"`
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceDisputeResolution) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceDisputeResolution) GetCurrency() (Currency InboundRealTimePaymentsTransferSimulationResultTransactionSourceDisputeResolutionCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The identifier of the Transaction that was disputed.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceDisputeResolution) GetDisputedTransactionId() (DisputedTransactionId string) {
	if r != nil && r.DisputedTransactionId != nil {
		DisputedTransactionId = *r.DisputedTransactionId
	}
	return
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceDisputeResolutionCurrency string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceDisputeResolutionCurrencyCad InboundRealTimePaymentsTransferSimulationResultTransactionSourceDisputeResolutionCurrency = "CAD"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceDisputeResolutionCurrencyChf InboundRealTimePaymentsTransferSimulationResultTransactionSourceDisputeResolutionCurrency = "CHF"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceDisputeResolutionCurrencyEur InboundRealTimePaymentsTransferSimulationResultTransactionSourceDisputeResolutionCurrency = "EUR"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceDisputeResolutionCurrencyGbp InboundRealTimePaymentsTransferSimulationResultTransactionSourceDisputeResolutionCurrency = "GBP"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceDisputeResolutionCurrencyJpy InboundRealTimePaymentsTransferSimulationResultTransactionSourceDisputeResolutionCurrency = "JPY"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceDisputeResolutionCurrencyUsd InboundRealTimePaymentsTransferSimulationResultTransactionSourceDisputeResolutionCurrency = "USD"
)

//
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceEmpyrealCashDeposit struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int `json:"amount"`
	//
	BagId *string `json:"bag_id"`
	//
	DepositDate *string `json:"deposit_date"`
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceEmpyrealCashDeposit) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceEmpyrealCashDeposit) GetBagId() (BagId string) {
	if r != nil && r.BagId != nil {
		BagId = *r.BagId
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceEmpyrealCashDeposit) GetDepositDate() (DepositDate string) {
	if r != nil && r.DepositDate != nil {
		DepositDate = *r.DepositDate
	}
	return
}

//
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundACHTransfer struct {
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
	OriginatorCompanyId *string `json:"originator_company_id"`
	//
	ReceiverIdNumber *string `json:"receiver_id_number"`
	//
	ReceiverName *string `json:"receiver_name"`
	//
	TraceNumber *string `json:"trace_number"`
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundACHTransfer) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundACHTransfer) GetOriginatorCompanyName() (OriginatorCompanyName string) {
	if r != nil && r.OriginatorCompanyName != nil {
		OriginatorCompanyName = *r.OriginatorCompanyName
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundACHTransfer) GetOriginatorCompanyDescriptiveDate() (OriginatorCompanyDescriptiveDate string) {
	if r != nil && r.OriginatorCompanyDescriptiveDate != nil {
		OriginatorCompanyDescriptiveDate = *r.OriginatorCompanyDescriptiveDate
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundACHTransfer) GetOriginatorCompanyDiscretionaryData() (OriginatorCompanyDiscretionaryData string) {
	if r != nil && r.OriginatorCompanyDiscretionaryData != nil {
		OriginatorCompanyDiscretionaryData = *r.OriginatorCompanyDiscretionaryData
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundACHTransfer) GetOriginatorCompanyEntryDescription() (OriginatorCompanyEntryDescription string) {
	if r != nil && r.OriginatorCompanyEntryDescription != nil {
		OriginatorCompanyEntryDescription = *r.OriginatorCompanyEntryDescription
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundACHTransfer) GetOriginatorCompanyId() (OriginatorCompanyId string) {
	if r != nil && r.OriginatorCompanyId != nil {
		OriginatorCompanyId = *r.OriginatorCompanyId
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundACHTransfer) GetReceiverIdNumber() (ReceiverIdNumber string) {
	if r != nil && r.ReceiverIdNumber != nil {
		ReceiverIdNumber = *r.ReceiverIdNumber
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundACHTransfer) GetReceiverName() (ReceiverName string) {
	if r != nil && r.ReceiverName != nil {
		ReceiverName = *r.ReceiverName
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundACHTransfer) GetTraceNumber() (TraceNumber string) {
	if r != nil && r.TraceNumber != nil {
		TraceNumber = *r.TraceNumber
	}
	return
}

//
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheck struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckCurrency `json:"currency"`
	//
	CheckNumber *string `json:"check_number"`
	//
	CheckFrontImageFileId *string `json:"check_front_image_file_id"`
	//
	CheckRearImageFileId *string `json:"check_rear_image_file_id"`
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheck) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheck) GetCurrency() (Currency InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheck) GetCheckNumber() (CheckNumber string) {
	if r != nil && r.CheckNumber != nil {
		CheckNumber = *r.CheckNumber
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheck) GetCheckFrontImageFileId() (CheckFrontImageFileId string) {
	if r != nil && r.CheckFrontImageFileId != nil {
		CheckFrontImageFileId = *r.CheckFrontImageFileId
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheck) GetCheckRearImageFileId() (CheckRearImageFileId string) {
	if r != nil && r.CheckRearImageFileId != nil {
		CheckRearImageFileId = *r.CheckRearImageFileId
	}
	return
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckCurrency string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckCurrencyCad InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckCurrency = "CAD"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckCurrencyChf InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckCurrency = "CHF"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckCurrencyEur InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckCurrency = "EUR"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckCurrencyGbp InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckCurrency = "GBP"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckCurrencyJpy InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckCurrency = "JPY"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckCurrencyUsd InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundCheckCurrency = "USD"
)

//
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer struct {
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
	OriginatingDepositoryFinancialInstitutionIdQualifier *string `json:"originating_depository_financial_institution_id_qualifier"`
	//
	OriginatingDepositoryFinancialInstitutionId *string `json:"originating_depository_financial_institution_id"`
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
	ReceivingDepositoryFinancialInstitutionIdQualifier *string `json:"receiving_depository_financial_institution_id_qualifier"`
	//
	ReceivingDepositoryFinancialInstitutionId *string `json:"receiving_depository_financial_institution_id"`
	//
	ReceivingDepositoryFinancialInstitutionCountry *string `json:"receiving_depository_financial_institution_country"`
	//
	TraceNumber *string `json:"trace_number"`
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer) GetForeignExchangeIndicator() (ForeignExchangeIndicator string) {
	if r != nil && r.ForeignExchangeIndicator != nil {
		ForeignExchangeIndicator = *r.ForeignExchangeIndicator
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer) GetForeignExchangeReferenceIndicator() (ForeignExchangeReferenceIndicator string) {
	if r != nil && r.ForeignExchangeReferenceIndicator != nil {
		ForeignExchangeReferenceIndicator = *r.ForeignExchangeReferenceIndicator
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer) GetForeignExchangeReference() (ForeignExchangeReference string) {
	if r != nil && r.ForeignExchangeReference != nil {
		ForeignExchangeReference = *r.ForeignExchangeReference
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer) GetDestinationCountryCode() (DestinationCountryCode string) {
	if r != nil && r.DestinationCountryCode != nil {
		DestinationCountryCode = *r.DestinationCountryCode
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer) GetDestinationCurrencyCode() (DestinationCurrencyCode string) {
	if r != nil && r.DestinationCurrencyCode != nil {
		DestinationCurrencyCode = *r.DestinationCurrencyCode
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer) GetForeignPaymentAmount() (ForeignPaymentAmount int) {
	if r != nil && r.ForeignPaymentAmount != nil {
		ForeignPaymentAmount = *r.ForeignPaymentAmount
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer) GetForeignTraceNumber() (ForeignTraceNumber string) {
	if r != nil && r.ForeignTraceNumber != nil {
		ForeignTraceNumber = *r.ForeignTraceNumber
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer) GetInternationalTransactionTypeCode() (InternationalTransactionTypeCode string) {
	if r != nil && r.InternationalTransactionTypeCode != nil {
		InternationalTransactionTypeCode = *r.InternationalTransactionTypeCode
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer) GetOriginatingCurrencyCode() (OriginatingCurrencyCode string) {
	if r != nil && r.OriginatingCurrencyCode != nil {
		OriginatingCurrencyCode = *r.OriginatingCurrencyCode
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer) GetOriginatingDepositoryFinancialInstitutionName() (OriginatingDepositoryFinancialInstitutionName string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionName != nil {
		OriginatingDepositoryFinancialInstitutionName = *r.OriginatingDepositoryFinancialInstitutionName
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer) GetOriginatingDepositoryFinancialInstitutionIdQualifier() (OriginatingDepositoryFinancialInstitutionIdQualifier string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionIdQualifier != nil {
		OriginatingDepositoryFinancialInstitutionIdQualifier = *r.OriginatingDepositoryFinancialInstitutionIdQualifier
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer) GetOriginatingDepositoryFinancialInstitutionId() (OriginatingDepositoryFinancialInstitutionId string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionId != nil {
		OriginatingDepositoryFinancialInstitutionId = *r.OriginatingDepositoryFinancialInstitutionId
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer) GetOriginatingDepositoryFinancialInstitutionBranchCountry() (OriginatingDepositoryFinancialInstitutionBranchCountry string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionBranchCountry != nil {
		OriginatingDepositoryFinancialInstitutionBranchCountry = *r.OriginatingDepositoryFinancialInstitutionBranchCountry
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer) GetOriginatorCity() (OriginatorCity string) {
	if r != nil && r.OriginatorCity != nil {
		OriginatorCity = *r.OriginatorCity
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer) GetOriginatorCompanyEntryDescription() (OriginatorCompanyEntryDescription string) {
	if r != nil && r.OriginatorCompanyEntryDescription != nil {
		OriginatorCompanyEntryDescription = *r.OriginatorCompanyEntryDescription
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer) GetOriginatorCountry() (OriginatorCountry string) {
	if r != nil && r.OriginatorCountry != nil {
		OriginatorCountry = *r.OriginatorCountry
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer) GetOriginatorIdentification() (OriginatorIdentification string) {
	if r != nil && r.OriginatorIdentification != nil {
		OriginatorIdentification = *r.OriginatorIdentification
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer) GetOriginatorName() (OriginatorName string) {
	if r != nil && r.OriginatorName != nil {
		OriginatorName = *r.OriginatorName
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer) GetOriginatorPostalCode() (OriginatorPostalCode string) {
	if r != nil && r.OriginatorPostalCode != nil {
		OriginatorPostalCode = *r.OriginatorPostalCode
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer) GetOriginatorStreetAddress() (OriginatorStreetAddress string) {
	if r != nil && r.OriginatorStreetAddress != nil {
		OriginatorStreetAddress = *r.OriginatorStreetAddress
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer) GetOriginatorStateOrProvince() (OriginatorStateOrProvince string) {
	if r != nil && r.OriginatorStateOrProvince != nil {
		OriginatorStateOrProvince = *r.OriginatorStateOrProvince
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer) GetPaymentRelatedInformation() (PaymentRelatedInformation string) {
	if r != nil && r.PaymentRelatedInformation != nil {
		PaymentRelatedInformation = *r.PaymentRelatedInformation
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer) GetPaymentRelatedInformation2() (PaymentRelatedInformation2 string) {
	if r != nil && r.PaymentRelatedInformation2 != nil {
		PaymentRelatedInformation2 = *r.PaymentRelatedInformation2
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer) GetReceiverIdentificationNumber() (ReceiverIdentificationNumber string) {
	if r != nil && r.ReceiverIdentificationNumber != nil {
		ReceiverIdentificationNumber = *r.ReceiverIdentificationNumber
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer) GetReceiverStreetAddress() (ReceiverStreetAddress string) {
	if r != nil && r.ReceiverStreetAddress != nil {
		ReceiverStreetAddress = *r.ReceiverStreetAddress
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer) GetReceiverCity() (ReceiverCity string) {
	if r != nil && r.ReceiverCity != nil {
		ReceiverCity = *r.ReceiverCity
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer) GetReceiverStateOrProvince() (ReceiverStateOrProvince string) {
	if r != nil && r.ReceiverStateOrProvince != nil {
		ReceiverStateOrProvince = *r.ReceiverStateOrProvince
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer) GetReceiverCountry() (ReceiverCountry string) {
	if r != nil && r.ReceiverCountry != nil {
		ReceiverCountry = *r.ReceiverCountry
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer) GetReceiverPostalCode() (ReceiverPostalCode string) {
	if r != nil && r.ReceiverPostalCode != nil {
		ReceiverPostalCode = *r.ReceiverPostalCode
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer) GetReceivingCompanyOrIndividualName() (ReceivingCompanyOrIndividualName string) {
	if r != nil && r.ReceivingCompanyOrIndividualName != nil {
		ReceivingCompanyOrIndividualName = *r.ReceivingCompanyOrIndividualName
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer) GetReceivingDepositoryFinancialInstitutionName() (ReceivingDepositoryFinancialInstitutionName string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionName != nil {
		ReceivingDepositoryFinancialInstitutionName = *r.ReceivingDepositoryFinancialInstitutionName
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer) GetReceivingDepositoryFinancialInstitutionIdQualifier() (ReceivingDepositoryFinancialInstitutionIdQualifier string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionIdQualifier != nil {
		ReceivingDepositoryFinancialInstitutionIdQualifier = *r.ReceivingDepositoryFinancialInstitutionIdQualifier
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer) GetReceivingDepositoryFinancialInstitutionId() (ReceivingDepositoryFinancialInstitutionId string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionId != nil {
		ReceivingDepositoryFinancialInstitutionId = *r.ReceivingDepositoryFinancialInstitutionId
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer) GetReceivingDepositoryFinancialInstitutionCountry() (ReceivingDepositoryFinancialInstitutionCountry string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionCountry != nil {
		ReceivingDepositoryFinancialInstitutionCountry = *r.ReceivingDepositoryFinancialInstitutionCountry
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundInternationalACHTransfer) GetTraceNumber() (TraceNumber string) {
	if r != nil && r.TraceNumber != nil {
		TraceNumber = *r.TraceNumber
	}
	return
}

//
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmation struct {
	// The amount in the minor unit of the transfer's currency. For dollars, for
	// example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the transfer's
	// currency. This will always be "USD" for a Real Time Payments transfer.
	Currency *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency `json:"currency"`
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
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmation) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the transfer's
// currency. This will always be "USD" for a Real Time Payments transfer.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmation) GetCurrency() (Currency InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The name the sender of the transfer specified as the recipient of the transfer.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmation) GetCreditorName() (CreditorName string) {
	if r != nil && r.CreditorName != nil {
		CreditorName = *r.CreditorName
	}
	return
}

// The name provided by the sender of the transfer.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmation) GetDebtorName() (DebtorName string) {
	if r != nil && r.DebtorName != nil {
		DebtorName = *r.DebtorName
	}
	return
}

// The account number of the account that sent the transfer.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmation) GetDebtorAccountNumber() (DebtorAccountNumber string) {
	if r != nil && r.DebtorAccountNumber != nil {
		DebtorAccountNumber = *r.DebtorAccountNumber
	}
	return
}

// The routing number of the account that sent the transfer.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmation) GetDebtorRoutingNumber() (DebtorRoutingNumber string) {
	if r != nil && r.DebtorRoutingNumber != nil {
		DebtorRoutingNumber = *r.DebtorRoutingNumber
	}
	return
}

// The Real Time Payments network identification of the transfer
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmation) GetTransactionIdentification() (TransactionIdentification string) {
	if r != nil && r.TransactionIdentification != nil {
		TransactionIdentification = *r.TransactionIdentification
	}
	return
}

// Additional information included with the transfer.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmation) GetRemittanceInformation() (RemittanceInformation string) {
	if r != nil && r.RemittanceInformation != nil {
		RemittanceInformation = *r.RemittanceInformation
	}
	return
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyCad InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "CAD"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyChf InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "CHF"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyEur InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "EUR"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyGbp InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "GBP"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyJpy InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "JPY"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrencyUsd InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundRealTimePaymentsTransferConfirmationCurrency = "USD"
)

//
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPaymentReversal struct {
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
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPaymentReversal) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The description on the reversal message from Fedwire.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPaymentReversal) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The Fedwire cycle date for the wire reversal.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPaymentReversal) GetInputCycleDate() (InputCycleDate string) {
	if r != nil && r.InputCycleDate != nil {
		InputCycleDate = *r.InputCycleDate
	}
	return
}

// The Fedwire sequence number.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPaymentReversal) GetInputSequenceNumber() (InputSequenceNumber string) {
	if r != nil && r.InputSequenceNumber != nil {
		InputSequenceNumber = *r.InputSequenceNumber
	}
	return
}

// The Fedwire input source identifier.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPaymentReversal) GetInputSource() (InputSource string) {
	if r != nil && r.InputSource != nil {
		InputSource = *r.InputSource
	}
	return
}

// The Fedwire transaction identifier.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPaymentReversal) GetInputMessageAccountabilityData() (InputMessageAccountabilityData string) {
	if r != nil && r.InputMessageAccountabilityData != nil {
		InputMessageAccountabilityData = *r.InputMessageAccountabilityData
	}
	return
}

// The Fedwire transaction identifier for the wire transfer that was reversed.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPaymentReversal) GetPreviousMessageInputMessageAccountabilityData() (PreviousMessageInputMessageAccountabilityData string) {
	if r != nil && r.PreviousMessageInputMessageAccountabilityData != nil {
		PreviousMessageInputMessageAccountabilityData = *r.PreviousMessageInputMessageAccountabilityData
	}
	return
}

// The Fedwire cycle date for the wire transfer that was reversed.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPaymentReversal) GetPreviousMessageInputCycleDate() (PreviousMessageInputCycleDate string) {
	if r != nil && r.PreviousMessageInputCycleDate != nil {
		PreviousMessageInputCycleDate = *r.PreviousMessageInputCycleDate
	}
	return
}

// The Fedwire sequence number for the wire transfer that was reversed.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPaymentReversal) GetPreviousMessageInputSequenceNumber() (PreviousMessageInputSequenceNumber string) {
	if r != nil && r.PreviousMessageInputSequenceNumber != nil {
		PreviousMessageInputSequenceNumber = *r.PreviousMessageInputSequenceNumber
	}
	return
}

// The Fedwire input source identifier for the wire transfer that was reversed.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPaymentReversal) GetPreviousMessageInputSource() (PreviousMessageInputSource string) {
	if r != nil && r.PreviousMessageInputSource != nil {
		PreviousMessageInputSource = *r.PreviousMessageInputSource
	}
	return
}

//
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPayment struct {
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
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPayment) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPayment) GetBeneficiaryAddressLine1() (BeneficiaryAddressLine1 string) {
	if r != nil && r.BeneficiaryAddressLine1 != nil {
		BeneficiaryAddressLine1 = *r.BeneficiaryAddressLine1
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPayment) GetBeneficiaryAddressLine2() (BeneficiaryAddressLine2 string) {
	if r != nil && r.BeneficiaryAddressLine2 != nil {
		BeneficiaryAddressLine2 = *r.BeneficiaryAddressLine2
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPayment) GetBeneficiaryAddressLine3() (BeneficiaryAddressLine3 string) {
	if r != nil && r.BeneficiaryAddressLine3 != nil {
		BeneficiaryAddressLine3 = *r.BeneficiaryAddressLine3
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPayment) GetBeneficiaryName() (BeneficiaryName string) {
	if r != nil && r.BeneficiaryName != nil {
		BeneficiaryName = *r.BeneficiaryName
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPayment) GetBeneficiaryReference() (BeneficiaryReference string) {
	if r != nil && r.BeneficiaryReference != nil {
		BeneficiaryReference = *r.BeneficiaryReference
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPayment) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPayment) GetInputMessageAccountabilityData() (InputMessageAccountabilityData string) {
	if r != nil && r.InputMessageAccountabilityData != nil {
		InputMessageAccountabilityData = *r.InputMessageAccountabilityData
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPayment) GetOriginatorAddressLine1() (OriginatorAddressLine1 string) {
	if r != nil && r.OriginatorAddressLine1 != nil {
		OriginatorAddressLine1 = *r.OriginatorAddressLine1
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPayment) GetOriginatorAddressLine2() (OriginatorAddressLine2 string) {
	if r != nil && r.OriginatorAddressLine2 != nil {
		OriginatorAddressLine2 = *r.OriginatorAddressLine2
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPayment) GetOriginatorAddressLine3() (OriginatorAddressLine3 string) {
	if r != nil && r.OriginatorAddressLine3 != nil {
		OriginatorAddressLine3 = *r.OriginatorAddressLine3
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPayment) GetOriginatorName() (OriginatorName string) {
	if r != nil && r.OriginatorName != nil {
		OriginatorName = *r.OriginatorName
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireDrawdownPayment) GetOriginatorToBeneficiaryInformation() (OriginatorToBeneficiaryInformation string) {
	if r != nil && r.OriginatorToBeneficiaryInformation != nil {
		OriginatorToBeneficiaryInformation = *r.OriginatorToBeneficiaryInformation
	}
	return
}

//
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireReversal struct {
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
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireReversal) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The description on the reversal message from Fedwire.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireReversal) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The Fedwire cycle date for the wire reversal.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireReversal) GetInputCycleDate() (InputCycleDate string) {
	if r != nil && r.InputCycleDate != nil {
		InputCycleDate = *r.InputCycleDate
	}
	return
}

// The Fedwire sequence number.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireReversal) GetInputSequenceNumber() (InputSequenceNumber string) {
	if r != nil && r.InputSequenceNumber != nil {
		InputSequenceNumber = *r.InputSequenceNumber
	}
	return
}

// The Fedwire input source identifier.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireReversal) GetInputSource() (InputSource string) {
	if r != nil && r.InputSource != nil {
		InputSource = *r.InputSource
	}
	return
}

// The Fedwire transaction identifier.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireReversal) GetInputMessageAccountabilityData() (InputMessageAccountabilityData string) {
	if r != nil && r.InputMessageAccountabilityData != nil {
		InputMessageAccountabilityData = *r.InputMessageAccountabilityData
	}
	return
}

// The Fedwire transaction identifier for the wire transfer that was reversed.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireReversal) GetPreviousMessageInputMessageAccountabilityData() (PreviousMessageInputMessageAccountabilityData string) {
	if r != nil && r.PreviousMessageInputMessageAccountabilityData != nil {
		PreviousMessageInputMessageAccountabilityData = *r.PreviousMessageInputMessageAccountabilityData
	}
	return
}

// The Fedwire cycle date for the wire transfer that was reversed.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireReversal) GetPreviousMessageInputCycleDate() (PreviousMessageInputCycleDate string) {
	if r != nil && r.PreviousMessageInputCycleDate != nil {
		PreviousMessageInputCycleDate = *r.PreviousMessageInputCycleDate
	}
	return
}

// The Fedwire sequence number for the wire transfer that was reversed.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireReversal) GetPreviousMessageInputSequenceNumber() (PreviousMessageInputSequenceNumber string) {
	if r != nil && r.PreviousMessageInputSequenceNumber != nil {
		PreviousMessageInputSequenceNumber = *r.PreviousMessageInputSequenceNumber
	}
	return
}

// The Fedwire input source identifier for the wire transfer that was reversed.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireReversal) GetPreviousMessageInputSource() (PreviousMessageInputSource string) {
	if r != nil && r.PreviousMessageInputSource != nil {
		PreviousMessageInputSource = *r.PreviousMessageInputSource
	}
	return
}

// Information included in the wire reversal for the receiving financial
// institution.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireReversal) GetReceiverFinancialInstitutionInformation() (ReceiverFinancialInstitutionInformation string) {
	if r != nil && r.ReceiverFinancialInstitutionInformation != nil {
		ReceiverFinancialInstitutionInformation = *r.ReceiverFinancialInstitutionInformation
	}
	return
}

// Additional financial institution information included in the wire reversal.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireReversal) GetFinancialInstitutionToFinancialInstitutionInformation() (FinancialInstitutionToFinancialInstitutionInformation string) {
	if r != nil && r.FinancialInstitutionToFinancialInstitutionInformation != nil {
		FinancialInstitutionToFinancialInstitutionInformation = *r.FinancialInstitutionToFinancialInstitutionInformation
	}
	return
}

//
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireTransfer struct {
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
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireTransfer) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireTransfer) GetBeneficiaryAddressLine1() (BeneficiaryAddressLine1 string) {
	if r != nil && r.BeneficiaryAddressLine1 != nil {
		BeneficiaryAddressLine1 = *r.BeneficiaryAddressLine1
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireTransfer) GetBeneficiaryAddressLine2() (BeneficiaryAddressLine2 string) {
	if r != nil && r.BeneficiaryAddressLine2 != nil {
		BeneficiaryAddressLine2 = *r.BeneficiaryAddressLine2
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireTransfer) GetBeneficiaryAddressLine3() (BeneficiaryAddressLine3 string) {
	if r != nil && r.BeneficiaryAddressLine3 != nil {
		BeneficiaryAddressLine3 = *r.BeneficiaryAddressLine3
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireTransfer) GetBeneficiaryName() (BeneficiaryName string) {
	if r != nil && r.BeneficiaryName != nil {
		BeneficiaryName = *r.BeneficiaryName
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireTransfer) GetBeneficiaryReference() (BeneficiaryReference string) {
	if r != nil && r.BeneficiaryReference != nil {
		BeneficiaryReference = *r.BeneficiaryReference
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireTransfer) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireTransfer) GetInputMessageAccountabilityData() (InputMessageAccountabilityData string) {
	if r != nil && r.InputMessageAccountabilityData != nil {
		InputMessageAccountabilityData = *r.InputMessageAccountabilityData
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireTransfer) GetOriginatorAddressLine1() (OriginatorAddressLine1 string) {
	if r != nil && r.OriginatorAddressLine1 != nil {
		OriginatorAddressLine1 = *r.OriginatorAddressLine1
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireTransfer) GetOriginatorAddressLine2() (OriginatorAddressLine2 string) {
	if r != nil && r.OriginatorAddressLine2 != nil {
		OriginatorAddressLine2 = *r.OriginatorAddressLine2
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireTransfer) GetOriginatorAddressLine3() (OriginatorAddressLine3 string) {
	if r != nil && r.OriginatorAddressLine3 != nil {
		OriginatorAddressLine3 = *r.OriginatorAddressLine3
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireTransfer) GetOriginatorName() (OriginatorName string) {
	if r != nil && r.OriginatorName != nil {
		OriginatorName = *r.OriginatorName
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInboundWireTransfer) GetOriginatorToBeneficiaryInformation() (OriginatorToBeneficiaryInformation string) {
	if r != nil && r.OriginatorToBeneficiaryInformation != nil {
		OriginatorToBeneficiaryInformation = *r.OriginatorToBeneficiaryInformation
	}
	return
}

//
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSource struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
	// currency.
	Currency *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceCurrency `json:"currency"`
	//
	Reason *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReason `json:"reason"`
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSource) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transaction
// currency.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSource) GetCurrency() (Currency InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSource) GetReason() (Reason InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReason) {
	if r != nil && r.Reason != nil {
		Reason = *r.Reason
	}
	return
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceCurrency string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceCurrencyCad InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceCurrency = "CAD"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceCurrencyChf InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceCurrency = "CHF"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceCurrencyEur InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceCurrency = "EUR"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceCurrencyGbp InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceCurrency = "GBP"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceCurrencyJpy InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceCurrency = "JPY"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceCurrencyUsd InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceCurrency = "USD"
)

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReason string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReasonCashback           InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReason = "cashback"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReasonEmpyrealAdjustment InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReason = "empyreal_adjustment"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReasonError              InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReason = "error"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReasonErrorCorrection    InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReason = "error_correction"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReasonFees               InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReason = "fees"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReasonInterest           InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReason = "interest"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReasonSampleFunds        InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReason = "sample_funds"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReasonSampleFundsReturn  InboundRealTimePaymentsTransferSimulationResultTransactionSourceInternalSourceReason = "sample_funds_return"
)

//
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefund struct {
	// The refunded amount in the minor unit of the refunded currency. For dollars, for
	// example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the refund
	// currency.
	Currency *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefundCurrency `json:"currency"`
	//
	MerchantAcceptorId *string `json:"merchant_acceptor_id"`
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
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefund) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the refund
// currency.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefund) GetCurrency() (Currency InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefundCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefund) GetMerchantAcceptorId() (MerchantAcceptorId string) {
	if r != nil && r.MerchantAcceptorId != nil {
		MerchantAcceptorId = *r.MerchantAcceptorId
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefund) GetMerchantCity() (MerchantCity string) {
	if r != nil && r.MerchantCity != nil {
		MerchantCity = *r.MerchantCity
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefund) GetMerchantCountry() (MerchantCountry string) {
	if r != nil && r.MerchantCountry != nil {
		MerchantCountry = *r.MerchantCountry
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefund) GetMerchantDescriptor() (MerchantDescriptor string) {
	if r != nil && r.MerchantDescriptor != nil {
		MerchantDescriptor = *r.MerchantDescriptor
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefund) GetMerchantState() (MerchantState string) {
	if r != nil && r.MerchantState != nil {
		MerchantState = *r.MerchantState
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefund) GetMerchantCategoryCode() (MerchantCategoryCode string) {
	if r != nil && r.MerchantCategoryCode != nil {
		MerchantCategoryCode = *r.MerchantCategoryCode
	}
	return
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefundCurrency string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefundCurrencyCad InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefundCurrency = "CAD"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefundCurrencyChf InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefundCurrency = "CHF"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefundCurrencyEur InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefundCurrency = "EUR"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefundCurrencyGbp InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefundCurrency = "GBP"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefundCurrencyJpy InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefundCurrency = "JPY"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefundCurrencyUsd InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteRefundCurrency = "USD"
)

//
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlement struct {
	// The settled amount in the minor unit of the settlement currency. For dollars,
	// for example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the settlement
	// currency.
	Currency *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlementCurrency `json:"currency"`
	//
	MerchantAcceptorId *string `json:"merchant_acceptor_id"`
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
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlement) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the settlement
// currency.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlement) GetCurrency() (Currency InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlementCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlement) GetMerchantAcceptorId() (MerchantAcceptorId string) {
	if r != nil && r.MerchantAcceptorId != nil {
		MerchantAcceptorId = *r.MerchantAcceptorId
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlement) GetMerchantCity() (MerchantCity string) {
	if r != nil && r.MerchantCity != nil {
		MerchantCity = *r.MerchantCity
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlement) GetMerchantCountry() (MerchantCountry string) {
	if r != nil && r.MerchantCountry != nil {
		MerchantCountry = *r.MerchantCountry
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlement) GetMerchantDescriptor() (MerchantDescriptor string) {
	if r != nil && r.MerchantDescriptor != nil {
		MerchantDescriptor = *r.MerchantDescriptor
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlement) GetMerchantState() (MerchantState string) {
	if r != nil && r.MerchantState != nil {
		MerchantState = *r.MerchantState
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlement) GetMerchantCategoryCode() (MerchantCategoryCode string) {
	if r != nil && r.MerchantCategoryCode != nil {
		MerchantCategoryCode = *r.MerchantCategoryCode
	}
	return
}

type InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlementCurrency string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlementCurrencyCad InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlementCurrency = "CAD"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlementCurrencyChf InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlementCurrency = "CHF"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlementCurrencyEur InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlementCurrency = "EUR"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlementCurrencyGbp InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlementCurrency = "GBP"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlementCurrencyJpy InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlementCurrency = "JPY"
	InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlementCurrencyUsd InboundRealTimePaymentsTransferSimulationResultTransactionSourceCardRouteSettlementCurrency = "USD"
)

//
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceSampleFunds struct {
	// Where the sample funds came from.
	Originator *string `json:"originator"`
}

// Where the sample funds came from.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceSampleFunds) GetOriginator() (Originator string) {
	if r != nil && r.Originator != nil {
		Originator = *r.Originator
	}
	return
}

//
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireDrawdownPaymentIntention struct {
	// The transfer amount in USD cents.
	Amount *int `json:"amount"`
	//
	AccountNumber *string `json:"account_number"`
	//
	RoutingNumber *string `json:"routing_number"`
	//
	MessageToRecipient *string `json:"message_to_recipient"`
	//
	TransferId *string `json:"transfer_id"`
}

// The transfer amount in USD cents.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireDrawdownPaymentIntention) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireDrawdownPaymentIntention) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireDrawdownPaymentIntention) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireDrawdownPaymentIntention) GetMessageToRecipient() (MessageToRecipient string) {
	if r != nil && r.MessageToRecipient != nil {
		MessageToRecipient = *r.MessageToRecipient
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireDrawdownPaymentIntention) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}

//
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireDrawdownPaymentRejection struct {
	//
	TransferId *string `json:"transfer_id"`
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireDrawdownPaymentRejection) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}

//
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferIntention struct {
	// The transfer amount in USD cents.
	Amount *int `json:"amount"`
	// The destination account number.
	AccountNumber *string `json:"account_number"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber *string `json:"routing_number"`
	// The message that will show on the recipient's bank statement.
	MessageToRecipient *string `json:"message_to_recipient"`
	//
	TransferId *string `json:"transfer_id"`
}

// The transfer amount in USD cents.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferIntention) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The destination account number.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferIntention) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

// The American Bankers' Association (ABA) Routing Transit Number (RTN).
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferIntention) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

// The message that will show on the recipient's bank statement.
func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferIntention) GetMessageToRecipient() (MessageToRecipient string) {
	if r != nil && r.MessageToRecipient != nil {
		MessageToRecipient = *r.MessageToRecipient
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferIntention) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}

//
type InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferRejection struct {
	//
	TransferId *string `json:"transfer_id"`
}

func (r *InboundRealTimePaymentsTransferSimulationResultTransactionSourceWireTransferRejection) GetTransferId() (TransferId string) {
	if r != nil && r.TransferId != nil {
		TransferId = *r.TransferId
	}
	return
}

type InboundRealTimePaymentsTransferSimulationResultTransactionType string

const (
	InboundRealTimePaymentsTransferSimulationResultTransactionTypeTransaction InboundRealTimePaymentsTransferSimulationResultTransactionType = "transaction"
)

//
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransaction struct {
	// The identifier for the Account the Declined Transaction belongs to.
	AccountId *string `json:"account_id"`
	// The Declined Transaction amount in the minor unit of its currency. For dollars,
	// for example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Declined
	// Transaction's currency. This will match the currency on the Declined
	// Transcation's Account.
	Currency *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionCurrency `json:"currency"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the
	// Transaction occured.
	CreatedAt *string `json:"created_at"`
	// This is the description the vendor provides.
	Description *string `json:"description"`
	// The Declined Transaction identifier.
	Id *string `json:"id"`
	// The identifier for the route this Declined Transaction came through. Routes are
	// things like cards and ACH details.
	RouteId *string `json:"route_id"`
	// The type of the route this Declined Transaction came through.
	RouteType *string `json:"route_type"`
	// This is an object giving more details on the network-level event that caused the
	// Declined Transaction. For example, for a card transaction this lists the
	// merchant's industry and location. Note that for backwards compatibility reasons,
	// additional undocumented keys may appear in this object. These should be treated
	// as deprecated and will be removed in the future.
	Source *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSource `json:"source"`
	// A constant representing the object's type. For this resource it will always be
	// `declined_transaction`.
	Type *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionType `json:"type"`
}

// The identifier for the Account the Declined Transaction belongs to.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransaction) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

// The Declined Transaction amount in the minor unit of its currency. For dollars,
// for example, this is cents.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransaction) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the Declined
// Transaction's currency. This will match the currency on the Declined
// Transcation's Account.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransaction) GetCurrency() (Currency InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date on which the
// Transaction occured.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransaction) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// This is the description the vendor provides.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransaction) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// The Declined Transaction identifier.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransaction) GetId() (Id string) {
	if r != nil && r.Id != nil {
		Id = *r.Id
	}
	return
}

// The identifier for the route this Declined Transaction came through. Routes are
// things like cards and ACH details.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransaction) GetRouteId() (RouteId string) {
	if r != nil && r.RouteId != nil {
		RouteId = *r.RouteId
	}
	return
}

// The type of the route this Declined Transaction came through.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransaction) GetRouteType() (RouteType string) {
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
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransaction) GetSource() (Source InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSource) {
	if r != nil && r.Source != nil {
		Source = *r.Source
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `declined_transaction`.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransaction) GetType() (Type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionCurrency string

const (
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionCurrencyCad InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionCurrency = "CAD"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionCurrencyChf InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionCurrency = "CHF"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionCurrencyEur InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionCurrency = "EUR"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionCurrencyGbp InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionCurrency = "GBP"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionCurrencyJpy InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionCurrency = "JPY"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionCurrencyUsd InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionCurrency = "USD"
)

//
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSource struct {
	// The type of decline that took place. We may add additional possible values for
	// this enum over time; your application should be able to handle such additions
	// gracefully.
	Category *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategory `json:"category"`
	// A ACH Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `ach_decline`.
	ACHDecline *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDecline `json:"ach_decline"`
	// A Card Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `card_decline`.
	CardDecline *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDecline `json:"card_decline"`
	// A Check Decline object. This field will be present in the JSON response if and
	// only if `category` is equal to `check_decline`.
	CheckDecline *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDecline `json:"check_decline"`
	// A Inbound Real Time Payments Transfer Decline object. This field will be present
	// in the JSON response if and only if `category` is equal to
	// `inbound_real_time_payments_transfer_decline`.
	InboundRealTimePaymentsTransferDecline *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline `json:"inbound_real_time_payments_transfer_decline"`
	// A International ACH Decline object. This field will be present in the JSON
	// response if and only if `category` is equal to `international_ach_decline`.
	InternationalACHDecline *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline `json:"international_ach_decline"`
	// A Deprecated Card Decline object. This field will be present in the JSON
	// response if and only if `category` is equal to `card_route_decline`.
	CardRouteDecline *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDecline `json:"card_route_decline"`
}

// The type of decline that took place. We may add additional possible values for
// this enum over time; your application should be able to handle such additions
// gracefully.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSource) GetCategory() (Category InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategory) {
	if r != nil && r.Category != nil {
		Category = *r.Category
	}
	return
}

// A ACH Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `ach_decline`.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSource) GetACHDecline() (ACHDecline InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDecline) {
	if r != nil && r.ACHDecline != nil {
		ACHDecline = *r.ACHDecline
	}
	return
}

// A Card Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `card_decline`.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSource) GetCardDecline() (CardDecline InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDecline) {
	if r != nil && r.CardDecline != nil {
		CardDecline = *r.CardDecline
	}
	return
}

// A Check Decline object. This field will be present in the JSON response if and
// only if `category` is equal to `check_decline`.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSource) GetCheckDecline() (CheckDecline InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDecline) {
	if r != nil && r.CheckDecline != nil {
		CheckDecline = *r.CheckDecline
	}
	return
}

// A Inbound Real Time Payments Transfer Decline object. This field will be present
// in the JSON response if and only if `category` is equal to
// `inbound_real_time_payments_transfer_decline`.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSource) GetInboundRealTimePaymentsTransferDecline() (InboundRealTimePaymentsTransferDecline InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) {
	if r != nil && r.InboundRealTimePaymentsTransferDecline != nil {
		InboundRealTimePaymentsTransferDecline = *r.InboundRealTimePaymentsTransferDecline
	}
	return
}

// A International ACH Decline object. This field will be present in the JSON
// response if and only if `category` is equal to `international_ach_decline`.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSource) GetInternationalACHDecline() (InternationalACHDecline InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline) {
	if r != nil && r.InternationalACHDecline != nil {
		InternationalACHDecline = *r.InternationalACHDecline
	}
	return
}

// A Deprecated Card Decline object. This field will be present in the JSON
// response if and only if `category` is equal to `card_route_decline`.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSource) GetCardRouteDecline() (CardRouteDecline InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDecline) {
	if r != nil && r.CardRouteDecline != nil {
		CardRouteDecline = *r.CardRouteDecline
	}
	return
}

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategory string

const (
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategoryACHDecline                             InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategory = "ach_decline"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategoryCardDecline                            InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategory = "card_decline"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategoryCheckDecline                           InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategory = "check_decline"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategoryInboundRealTimePaymentsTransferDecline InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategory = "inbound_real_time_payments_transfer_decline"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategoryInternationalACHDecline                InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategory = "international_ach_decline"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategoryCardRouteDecline                       InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategory = "card_route_decline"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategoryOther                                  InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCategory = "other"
)

//
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDecline struct {
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
	OriginatorCompanyId *string `json:"originator_company_id"`
	// Why the ACH transfer was declined.
	Reason *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason `json:"reason"`
	//
	ReceiverIdNumber *string `json:"receiver_id_number"`
	//
	ReceiverName *string `json:"receiver_name"`
	//
	TraceNumber *string `json:"trace_number"`
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDecline) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDecline) GetOriginatorCompanyName() (OriginatorCompanyName string) {
	if r != nil && r.OriginatorCompanyName != nil {
		OriginatorCompanyName = *r.OriginatorCompanyName
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDecline) GetOriginatorCompanyDescriptiveDate() (OriginatorCompanyDescriptiveDate string) {
	if r != nil && r.OriginatorCompanyDescriptiveDate != nil {
		OriginatorCompanyDescriptiveDate = *r.OriginatorCompanyDescriptiveDate
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDecline) GetOriginatorCompanyDiscretionaryData() (OriginatorCompanyDiscretionaryData string) {
	if r != nil && r.OriginatorCompanyDiscretionaryData != nil {
		OriginatorCompanyDiscretionaryData = *r.OriginatorCompanyDiscretionaryData
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDecline) GetOriginatorCompanyId() (OriginatorCompanyId string) {
	if r != nil && r.OriginatorCompanyId != nil {
		OriginatorCompanyId = *r.OriginatorCompanyId
	}
	return
}

// Why the ACH transfer was declined.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDecline) GetReason() (Reason InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason) {
	if r != nil && r.Reason != nil {
		Reason = *r.Reason
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDecline) GetReceiverIdNumber() (ReceiverIdNumber string) {
	if r != nil && r.ReceiverIdNumber != nil {
		ReceiverIdNumber = *r.ReceiverIdNumber
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDecline) GetReceiverName() (ReceiverName string) {
	if r != nil && r.ReceiverName != nil {
		ReceiverName = *r.ReceiverName
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDecline) GetTraceNumber() (TraceNumber string) {
	if r != nil && r.TraceNumber != nil {
		TraceNumber = *r.TraceNumber
	}
	return
}

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason string

const (
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonACHRouteCanceled             InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "ach_route_canceled"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonACHRouteDisabled             InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "ach_route_disabled"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonNoACHRoute                   InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "no_ach_route"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonBreachesLimit                InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "breaches_limit"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonCreditEntryRefusedByReceiver InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "credit_entry_refused_by_receiver"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonGroupLocked                  InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "group_locked"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonEntityNotActive              InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "entity_not_active"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonInsufficientFunds            InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "insufficient_funds"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReasonOriginatorRequest            InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceACHDeclineReason = "originator_request"
)

//
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineCurrency `json:"currency"`
	//
	MerchantAcceptorId *string `json:"merchant_acceptor_id"`
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
	Reason *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason `json:"reason"`
	// The identifier of the Real-Time Decision sent to approve or decline this
	// transaction.
	RealTimeDecisionId *string `json:"real_time_decision_id"`
	// If the authorization was attempted using a Digital Wallet Token (such as an
	// Apple Pay purchase), the identifier of the token that was used.
	DigitalWalletTokenId *string `json:"digital_wallet_token_id"`
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDecline) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
// account currency.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDecline) GetCurrency() (Currency InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDecline) GetMerchantAcceptorId() (MerchantAcceptorId string) {
	if r != nil && r.MerchantAcceptorId != nil {
		MerchantAcceptorId = *r.MerchantAcceptorId
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDecline) GetMerchantCity() (MerchantCity string) {
	if r != nil && r.MerchantCity != nil {
		MerchantCity = *r.MerchantCity
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDecline) GetMerchantCountry() (MerchantCountry string) {
	if r != nil && r.MerchantCountry != nil {
		MerchantCountry = *r.MerchantCountry
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDecline) GetMerchantDescriptor() (MerchantDescriptor string) {
	if r != nil && r.MerchantDescriptor != nil {
		MerchantDescriptor = *r.MerchantDescriptor
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDecline) GetMerchantState() (MerchantState string) {
	if r != nil && r.MerchantState != nil {
		MerchantState = *r.MerchantState
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDecline) GetMerchantCategoryCode() (MerchantCategoryCode string) {
	if r != nil && r.MerchantCategoryCode != nil {
		MerchantCategoryCode = *r.MerchantCategoryCode
	}
	return
}

// Why the transaction was declined.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDecline) GetReason() (Reason InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason) {
	if r != nil && r.Reason != nil {
		Reason = *r.Reason
	}
	return
}

// The identifier of the Real-Time Decision sent to approve or decline this
// transaction.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDecline) GetRealTimeDecisionId() (RealTimeDecisionId string) {
	if r != nil && r.RealTimeDecisionId != nil {
		RealTimeDecisionId = *r.RealTimeDecisionId
	}
	return
}

// If the authorization was attempted using a Digital Wallet Token (such as an
// Apple Pay purchase), the identifier of the token that was used.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDecline) GetDigitalWalletTokenId() (DigitalWalletTokenId string) {
	if r != nil && r.DigitalWalletTokenId != nil {
		DigitalWalletTokenId = *r.DigitalWalletTokenId
	}
	return
}

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineCurrency string

const (
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineCurrencyCad InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineCurrency = "CAD"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineCurrencyChf InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineCurrency = "CHF"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineCurrencyEur InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineCurrency = "EUR"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineCurrencyGbp InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineCurrency = "GBP"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineCurrencyJpy InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineCurrency = "JPY"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineCurrencyUsd InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineCurrency = "USD"
)

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason string

const (
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReasonCardNotActive     InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason = "card_not_active"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReasonEntityNotActive   InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason = "entity_not_active"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReasonGroupLocked       InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason = "group_locked"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReasonInsufficientFunds InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason = "insufficient_funds"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReasonBreachesLimit     InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason = "breaches_limit"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReasonWebhookDeclined   InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason = "webhook_declined"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReasonWebhookTimedOut   InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardDeclineReason = "webhook_timed_out"
)

//
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount *int `json:"amount"`
	//
	AuxiliaryOnUs *string `json:"auxiliary_on_us"`
	// Why the check was declined.
	Reason *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason `json:"reason"`
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDecline) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDecline) GetAuxiliaryOnUs() (AuxiliaryOnUs string) {
	if r != nil && r.AuxiliaryOnUs != nil {
		AuxiliaryOnUs = *r.AuxiliaryOnUs
	}
	return
}

// Why the check was declined.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDecline) GetReason() (Reason InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason) {
	if r != nil && r.Reason != nil {
		Reason = *r.Reason
	}
	return
}

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason string

const (
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonACHRouteCanceled      InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "ach_route_canceled"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonACHRouteDisabled      InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "ach_route_disabled"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonBreachesLimit         InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "breaches_limit"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonEntityNotActive       InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "entity_not_active"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonGroupLocked           InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "group_locked"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonInsufficientFunds     InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "insufficient_funds"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonUnableToLocateAccount InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "unable_to_locate_account"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonUnableToProcess       InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "unable_to_process"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonReferToImage          InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "refer_to_image"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReasonStopPaymentRequested  InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCheckDeclineReason = "stop_payment_requested"
)

//
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the declined
	// transfer's currency. This will always be "USD" for a Real Time Payments
	// transfer.
	Currency *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency `json:"currency"`
	// Why the transfer was declined.
	Reason *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason `json:"reason"`
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
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code of the declined
// transfer's currency. This will always be "USD" for a Real Time Payments
// transfer.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) GetCurrency() (Currency InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// Why the transfer was declined.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) GetReason() (Reason InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason) {
	if r != nil && r.Reason != nil {
		Reason = *r.Reason
	}
	return
}

// The name the sender of the transfer specified as the recipient of the transfer.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) GetCreditorName() (CreditorName string) {
	if r != nil && r.CreditorName != nil {
		CreditorName = *r.CreditorName
	}
	return
}

// The name provided by the sender of the transfer.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) GetDebtorName() (DebtorName string) {
	if r != nil && r.DebtorName != nil {
		DebtorName = *r.DebtorName
	}
	return
}

// The account number of the account that sent the transfer.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) GetDebtorAccountNumber() (DebtorAccountNumber string) {
	if r != nil && r.DebtorAccountNumber != nil {
		DebtorAccountNumber = *r.DebtorAccountNumber
	}
	return
}

// The routing number of the account that sent the transfer.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) GetDebtorRoutingNumber() (DebtorRoutingNumber string) {
	if r != nil && r.DebtorRoutingNumber != nil {
		DebtorRoutingNumber = *r.DebtorRoutingNumber
	}
	return
}

// The Real Time Payments network identification of the declined transfer.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) GetTransactionIdentification() (TransactionIdentification string) {
	if r != nil && r.TransactionIdentification != nil {
		TransactionIdentification = *r.TransactionIdentification
	}
	return
}

// Additional information included with the transfer.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDecline) GetRemittanceInformation() (RemittanceInformation string) {
	if r != nil && r.RemittanceInformation != nil {
		RemittanceInformation = *r.RemittanceInformation
	}
	return
}

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency string

const (
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyCad InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "CAD"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyChf InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "CHF"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyEur InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "EUR"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyGbp InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "GBP"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyJpy InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "JPY"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrencyUsd InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineCurrency = "USD"
)

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason string

const (
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonAccountNumberCanceled      InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "account_number_canceled"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonAccountNumberDisabled      InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "account_number_disabled"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonGroupLocked                InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "group_locked"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonEntityNotActive            InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "entity_not_active"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReasonRealTimePaymentsNotEnabled InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInboundRealTimePaymentsTransferDeclineReason = "real_time_payments_not_enabled"
)

//
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline struct {
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
	OriginatingDepositoryFinancialInstitutionIdQualifier *string `json:"originating_depository_financial_institution_id_qualifier"`
	//
	OriginatingDepositoryFinancialInstitutionId *string `json:"originating_depository_financial_institution_id"`
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
	ReceivingDepositoryFinancialInstitutionIdQualifier *string `json:"receiving_depository_financial_institution_id_qualifier"`
	//
	ReceivingDepositoryFinancialInstitutionId *string `json:"receiving_depository_financial_institution_id"`
	//
	ReceivingDepositoryFinancialInstitutionCountry *string `json:"receiving_depository_financial_institution_country"`
	//
	TraceNumber *string `json:"trace_number"`
}

// The declined amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline) GetForeignExchangeIndicator() (ForeignExchangeIndicator string) {
	if r != nil && r.ForeignExchangeIndicator != nil {
		ForeignExchangeIndicator = *r.ForeignExchangeIndicator
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline) GetForeignExchangeReferenceIndicator() (ForeignExchangeReferenceIndicator string) {
	if r != nil && r.ForeignExchangeReferenceIndicator != nil {
		ForeignExchangeReferenceIndicator = *r.ForeignExchangeReferenceIndicator
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline) GetForeignExchangeReference() (ForeignExchangeReference string) {
	if r != nil && r.ForeignExchangeReference != nil {
		ForeignExchangeReference = *r.ForeignExchangeReference
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline) GetDestinationCountryCode() (DestinationCountryCode string) {
	if r != nil && r.DestinationCountryCode != nil {
		DestinationCountryCode = *r.DestinationCountryCode
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline) GetDestinationCurrencyCode() (DestinationCurrencyCode string) {
	if r != nil && r.DestinationCurrencyCode != nil {
		DestinationCurrencyCode = *r.DestinationCurrencyCode
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline) GetForeignPaymentAmount() (ForeignPaymentAmount int) {
	if r != nil && r.ForeignPaymentAmount != nil {
		ForeignPaymentAmount = *r.ForeignPaymentAmount
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline) GetForeignTraceNumber() (ForeignTraceNumber string) {
	if r != nil && r.ForeignTraceNumber != nil {
		ForeignTraceNumber = *r.ForeignTraceNumber
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline) GetInternationalTransactionTypeCode() (InternationalTransactionTypeCode string) {
	if r != nil && r.InternationalTransactionTypeCode != nil {
		InternationalTransactionTypeCode = *r.InternationalTransactionTypeCode
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline) GetOriginatingCurrencyCode() (OriginatingCurrencyCode string) {
	if r != nil && r.OriginatingCurrencyCode != nil {
		OriginatingCurrencyCode = *r.OriginatingCurrencyCode
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline) GetOriginatingDepositoryFinancialInstitutionName() (OriginatingDepositoryFinancialInstitutionName string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionName != nil {
		OriginatingDepositoryFinancialInstitutionName = *r.OriginatingDepositoryFinancialInstitutionName
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline) GetOriginatingDepositoryFinancialInstitutionIdQualifier() (OriginatingDepositoryFinancialInstitutionIdQualifier string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionIdQualifier != nil {
		OriginatingDepositoryFinancialInstitutionIdQualifier = *r.OriginatingDepositoryFinancialInstitutionIdQualifier
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline) GetOriginatingDepositoryFinancialInstitutionId() (OriginatingDepositoryFinancialInstitutionId string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionId != nil {
		OriginatingDepositoryFinancialInstitutionId = *r.OriginatingDepositoryFinancialInstitutionId
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline) GetOriginatingDepositoryFinancialInstitutionBranchCountry() (OriginatingDepositoryFinancialInstitutionBranchCountry string) {
	if r != nil && r.OriginatingDepositoryFinancialInstitutionBranchCountry != nil {
		OriginatingDepositoryFinancialInstitutionBranchCountry = *r.OriginatingDepositoryFinancialInstitutionBranchCountry
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline) GetOriginatorCity() (OriginatorCity string) {
	if r != nil && r.OriginatorCity != nil {
		OriginatorCity = *r.OriginatorCity
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline) GetOriginatorCompanyEntryDescription() (OriginatorCompanyEntryDescription string) {
	if r != nil && r.OriginatorCompanyEntryDescription != nil {
		OriginatorCompanyEntryDescription = *r.OriginatorCompanyEntryDescription
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline) GetOriginatorCountry() (OriginatorCountry string) {
	if r != nil && r.OriginatorCountry != nil {
		OriginatorCountry = *r.OriginatorCountry
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline) GetOriginatorIdentification() (OriginatorIdentification string) {
	if r != nil && r.OriginatorIdentification != nil {
		OriginatorIdentification = *r.OriginatorIdentification
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline) GetOriginatorName() (OriginatorName string) {
	if r != nil && r.OriginatorName != nil {
		OriginatorName = *r.OriginatorName
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline) GetOriginatorPostalCode() (OriginatorPostalCode string) {
	if r != nil && r.OriginatorPostalCode != nil {
		OriginatorPostalCode = *r.OriginatorPostalCode
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline) GetOriginatorStreetAddress() (OriginatorStreetAddress string) {
	if r != nil && r.OriginatorStreetAddress != nil {
		OriginatorStreetAddress = *r.OriginatorStreetAddress
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline) GetOriginatorStateOrProvince() (OriginatorStateOrProvince string) {
	if r != nil && r.OriginatorStateOrProvince != nil {
		OriginatorStateOrProvince = *r.OriginatorStateOrProvince
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline) GetPaymentRelatedInformation() (PaymentRelatedInformation string) {
	if r != nil && r.PaymentRelatedInformation != nil {
		PaymentRelatedInformation = *r.PaymentRelatedInformation
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline) GetPaymentRelatedInformation2() (PaymentRelatedInformation2 string) {
	if r != nil && r.PaymentRelatedInformation2 != nil {
		PaymentRelatedInformation2 = *r.PaymentRelatedInformation2
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline) GetReceiverIdentificationNumber() (ReceiverIdentificationNumber string) {
	if r != nil && r.ReceiverIdentificationNumber != nil {
		ReceiverIdentificationNumber = *r.ReceiverIdentificationNumber
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline) GetReceiverStreetAddress() (ReceiverStreetAddress string) {
	if r != nil && r.ReceiverStreetAddress != nil {
		ReceiverStreetAddress = *r.ReceiverStreetAddress
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline) GetReceiverCity() (ReceiverCity string) {
	if r != nil && r.ReceiverCity != nil {
		ReceiverCity = *r.ReceiverCity
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline) GetReceiverStateOrProvince() (ReceiverStateOrProvince string) {
	if r != nil && r.ReceiverStateOrProvince != nil {
		ReceiverStateOrProvince = *r.ReceiverStateOrProvince
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline) GetReceiverCountry() (ReceiverCountry string) {
	if r != nil && r.ReceiverCountry != nil {
		ReceiverCountry = *r.ReceiverCountry
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline) GetReceiverPostalCode() (ReceiverPostalCode string) {
	if r != nil && r.ReceiverPostalCode != nil {
		ReceiverPostalCode = *r.ReceiverPostalCode
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline) GetReceivingCompanyOrIndividualName() (ReceivingCompanyOrIndividualName string) {
	if r != nil && r.ReceivingCompanyOrIndividualName != nil {
		ReceivingCompanyOrIndividualName = *r.ReceivingCompanyOrIndividualName
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline) GetReceivingDepositoryFinancialInstitutionName() (ReceivingDepositoryFinancialInstitutionName string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionName != nil {
		ReceivingDepositoryFinancialInstitutionName = *r.ReceivingDepositoryFinancialInstitutionName
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline) GetReceivingDepositoryFinancialInstitutionIdQualifier() (ReceivingDepositoryFinancialInstitutionIdQualifier string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionIdQualifier != nil {
		ReceivingDepositoryFinancialInstitutionIdQualifier = *r.ReceivingDepositoryFinancialInstitutionIdQualifier
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline) GetReceivingDepositoryFinancialInstitutionId() (ReceivingDepositoryFinancialInstitutionId string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionId != nil {
		ReceivingDepositoryFinancialInstitutionId = *r.ReceivingDepositoryFinancialInstitutionId
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline) GetReceivingDepositoryFinancialInstitutionCountry() (ReceivingDepositoryFinancialInstitutionCountry string) {
	if r != nil && r.ReceivingDepositoryFinancialInstitutionCountry != nil {
		ReceivingDepositoryFinancialInstitutionCountry = *r.ReceivingDepositoryFinancialInstitutionCountry
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceInternationalACHDecline) GetTraceNumber() (TraceNumber string) {
	if r != nil && r.TraceNumber != nil {
		TraceNumber = *r.TraceNumber
	}
	return
}

//
type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDecline struct {
	// The declined amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount *int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
	// account currency.
	Currency *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDeclineCurrency `json:"currency"`
	//
	MerchantAcceptorId *string `json:"merchant_acceptor_id"`
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
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDecline) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the destination
// account currency.
func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDecline) GetCurrency() (Currency InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDeclineCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDecline) GetMerchantAcceptorId() (MerchantAcceptorId string) {
	if r != nil && r.MerchantAcceptorId != nil {
		MerchantAcceptorId = *r.MerchantAcceptorId
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDecline) GetMerchantCity() (MerchantCity string) {
	if r != nil && r.MerchantCity != nil {
		MerchantCity = *r.MerchantCity
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDecline) GetMerchantCountry() (MerchantCountry string) {
	if r != nil && r.MerchantCountry != nil {
		MerchantCountry = *r.MerchantCountry
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDecline) GetMerchantDescriptor() (MerchantDescriptor string) {
	if r != nil && r.MerchantDescriptor != nil {
		MerchantDescriptor = *r.MerchantDescriptor
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDecline) GetMerchantState() (MerchantState string) {
	if r != nil && r.MerchantState != nil {
		MerchantState = *r.MerchantState
	}
	return
}

func (r *InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDecline) GetMerchantCategoryCode() (MerchantCategoryCode string) {
	if r != nil && r.MerchantCategoryCode != nil {
		MerchantCategoryCode = *r.MerchantCategoryCode
	}
	return
}

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDeclineCurrency string

const (
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDeclineCurrencyCad InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDeclineCurrency = "CAD"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDeclineCurrencyChf InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDeclineCurrency = "CHF"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDeclineCurrencyEur InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDeclineCurrency = "EUR"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDeclineCurrencyGbp InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDeclineCurrency = "GBP"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDeclineCurrencyJpy InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDeclineCurrency = "JPY"
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDeclineCurrencyUsd InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionSourceCardRouteDeclineCurrency = "USD"
)

type InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionType string

const (
	InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionTypeDeclinedTransaction InboundRealTimePaymentsTransferSimulationResultDeclinedTransactionType = "declined_transaction"
)

type InboundRealTimePaymentsTransferSimulationResultType string

const (
	InboundRealTimePaymentsTransferSimulationResultTypeInboundRealTimePaymentsTransferSimulationResult InboundRealTimePaymentsTransferSimulationResultType = "inbound_real_time_payments_transfer_simulation_result"
)

type SimulateARealTimePaymentsTransferToYourAccountParameters struct {
	// The identifier of the Account Number the inbound Real Time Payments Transfer is
	// for.
	AccountNumberId *string `json:"account_number_id"`
	// The transfer amount in USD cents. Must be positive.
	Amount *int `json:"amount"`
}

// The identifier of the Account Number the inbound Real Time Payments Transfer is
// for.
func (r *SimulateARealTimePaymentsTransferToYourAccountParameters) GetAccountNumberId() (AccountNumberId string) {
	if r != nil && r.AccountNumberId != nil {
		AccountNumberId = *r.AccountNumberId
	}
	return
}

// The transfer amount in USD cents. Must be positive.
func (r *SimulateARealTimePaymentsTransferToYourAccountParameters) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

package requests

import (
	"net/url"
	"time"

	"github.com/increase/increase-go/core/field"
	"github.com/increase/increase-go/core/query"
)

type TransactionListParams struct {
	// Return the page of entries after this one.
	Cursor field.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit field.Field[int64] `query:"limit"`
	// Filter Transactions for those belonging to the specified Account.
	AccountID field.Field[string] `query:"account_id"`
	// Filter Transactions for those belonging to the specified route.
	RouteID   field.Field[string]                         `query:"route_id"`
	CreatedAt field.Field[TransactionListParamsCreatedAt] `query:"created_at"`
	Category  field.Field[TransactionListParamsCategory]  `query:"category"`
}

// URLQuery serializes TransactionListParams into a url.Values of the query
// parameters associated with this value
func (r *TransactionListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type TransactionListParamsCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After field.Field[time.Time] `query:"after" format:"date-time"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before field.Field[time.Time] `query:"before" format:"date-time"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter field.Field[time.Time] `query:"on_or_after" format:"date-time"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore field.Field[time.Time] `query:"on_or_before" format:"date-time"`
}

// URLQuery serializes TransactionListParamsCreatedAt into a url.Values of the
// query parameters associated with this value
func (r *TransactionListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type TransactionListParamsCategory struct {
	// Return results whose value is in the provided list. For GET requests, this
	// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
	In field.Field[[]TransactionListParamsCategoryIn] `query:"in"`
}

// URLQuery serializes TransactionListParamsCategory into a url.Values of the query
// parameters associated with this value
func (r *TransactionListParamsCategory) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type TransactionListParamsCategoryIn string

const (
	TransactionListParamsCategoryInAccountTransferIntention                    TransactionListParamsCategoryIn = "account_transfer_intention"
	TransactionListParamsCategoryInACHCheckConversionReturn                    TransactionListParamsCategoryIn = "ach_check_conversion_return"
	TransactionListParamsCategoryInACHCheckConversion                          TransactionListParamsCategoryIn = "ach_check_conversion"
	TransactionListParamsCategoryInACHTransferIntention                        TransactionListParamsCategoryIn = "ach_transfer_intention"
	TransactionListParamsCategoryInACHTransferRejection                        TransactionListParamsCategoryIn = "ach_transfer_rejection"
	TransactionListParamsCategoryInACHTransferReturn                           TransactionListParamsCategoryIn = "ach_transfer_return"
	TransactionListParamsCategoryInCardDisputeAcceptance                       TransactionListParamsCategoryIn = "card_dispute_acceptance"
	TransactionListParamsCategoryInCardRefund                                  TransactionListParamsCategoryIn = "card_refund"
	TransactionListParamsCategoryInCardSettlement                              TransactionListParamsCategoryIn = "card_settlement"
	TransactionListParamsCategoryInCheckDepositAcceptance                      TransactionListParamsCategoryIn = "check_deposit_acceptance"
	TransactionListParamsCategoryInCheckDepositReturn                          TransactionListParamsCategoryIn = "check_deposit_return"
	TransactionListParamsCategoryInCheckTransferIntention                      TransactionListParamsCategoryIn = "check_transfer_intention"
	TransactionListParamsCategoryInCheckTransferReturn                         TransactionListParamsCategoryIn = "check_transfer_return"
	TransactionListParamsCategoryInCheckTransferRejection                      TransactionListParamsCategoryIn = "check_transfer_rejection"
	TransactionListParamsCategoryInCheckTransferStopPaymentRequest             TransactionListParamsCategoryIn = "check_transfer_stop_payment_request"
	TransactionListParamsCategoryInDisputeResolution                           TransactionListParamsCategoryIn = "dispute_resolution"
	TransactionListParamsCategoryInEmpyrealCashDeposit                         TransactionListParamsCategoryIn = "empyreal_cash_deposit"
	TransactionListParamsCategoryInInboundACHTransfer                          TransactionListParamsCategoryIn = "inbound_ach_transfer"
	TransactionListParamsCategoryInInboundACHTransferReturnIntention           TransactionListParamsCategoryIn = "inbound_ach_transfer_return_intention"
	TransactionListParamsCategoryInInboundCheck                                TransactionListParamsCategoryIn = "inbound_check"
	TransactionListParamsCategoryInInboundInternationalACHTransfer             TransactionListParamsCategoryIn = "inbound_international_ach_transfer"
	TransactionListParamsCategoryInInboundRealTimePaymentsTransferConfirmation TransactionListParamsCategoryIn = "inbound_real_time_payments_transfer_confirmation"
	TransactionListParamsCategoryInInboundWireDrawdownPaymentReversal          TransactionListParamsCategoryIn = "inbound_wire_drawdown_payment_reversal"
	TransactionListParamsCategoryInInboundWireDrawdownPayment                  TransactionListParamsCategoryIn = "inbound_wire_drawdown_payment"
	TransactionListParamsCategoryInInboundWireReversal                         TransactionListParamsCategoryIn = "inbound_wire_reversal"
	TransactionListParamsCategoryInInboundWireTransfer                         TransactionListParamsCategoryIn = "inbound_wire_transfer"
	TransactionListParamsCategoryInInterestPayment                             TransactionListParamsCategoryIn = "interest_payment"
	TransactionListParamsCategoryInInternalGeneralLedgerTransaction            TransactionListParamsCategoryIn = "internal_general_ledger_transaction"
	TransactionListParamsCategoryInInternalSource                              TransactionListParamsCategoryIn = "internal_source"
	TransactionListParamsCategoryInCardRouteRefund                             TransactionListParamsCategoryIn = "card_route_refund"
	TransactionListParamsCategoryInCardRouteSettlement                         TransactionListParamsCategoryIn = "card_route_settlement"
	TransactionListParamsCategoryInRealTimePaymentsTransferAcknowledgement     TransactionListParamsCategoryIn = "real_time_payments_transfer_acknowledgement"
	TransactionListParamsCategoryInSampleFunds                                 TransactionListParamsCategoryIn = "sample_funds"
	TransactionListParamsCategoryInWireDrawdownPaymentIntention                TransactionListParamsCategoryIn = "wire_drawdown_payment_intention"
	TransactionListParamsCategoryInWireDrawdownPaymentRejection                TransactionListParamsCategoryIn = "wire_drawdown_payment_rejection"
	TransactionListParamsCategoryInWireTransferIntention                       TransactionListParamsCategoryIn = "wire_transfer_intention"
	TransactionListParamsCategoryInWireTransferRejection                       TransactionListParamsCategoryIn = "wire_transfer_rejection"
	TransactionListParamsCategoryInOther                                       TransactionListParamsCategoryIn = "other"
)

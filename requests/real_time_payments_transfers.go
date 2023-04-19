package requests

import (
	"net/url"
	"time"

	"github.com/increase/increase-go/core/field"
	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
)

type RealTimePaymentsTransferNewParams struct {
	// The identifier of the Account Number from which to send the transfer.
	SourceAccountNumberID field.Field[string] `json:"source_account_number_id,required"`
	// The destination account number.
	DestinationAccountNumber field.Field[string] `json:"destination_account_number"`
	// The destination American Bankers' Association (ABA) Routing Transit Number
	// (RTN).
	DestinationRoutingNumber field.Field[string] `json:"destination_routing_number"`
	// The ID of an External Account to initiate a transfer to. If this parameter is
	// provided, `destination_account_number` and `destination_routing_number` must be
	// absent.
	ExternalAccountID field.Field[string] `json:"external_account_id"`
	// The transfer amount in USD cents. For Real Time Payments transfers, must be
	// positive.
	Amount field.Field[int64] `json:"amount,required"`
	// The name of the transfer's recipient.
	CreditorName field.Field[string] `json:"creditor_name,required"`
	// Unstructured information that will show on the recipient's bank statement.
	RemittanceInformation field.Field[string] `json:"remittance_information,required"`
	// Whether the transfer requires explicit approval via the dashboard or API.
	RequireApproval field.Field[bool] `json:"require_approval"`
}

// MarshalJSON serializes RealTimePaymentsTransferNewParams into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r RealTimePaymentsTransferNewParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type RealTimePaymentsTransferListParams struct {
	// Return the page of entries after this one.
	Cursor field.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit field.Field[int64] `query:"limit"`
	// Filter Real Time Payments Transfers to those belonging to the specified Account.
	AccountID field.Field[string] `query:"account_id"`
	// Filter Real Time Payments Transfers to those made to the specified External
	// Account.
	ExternalAccountID field.Field[string]                                      `query:"external_account_id"`
	CreatedAt         field.Field[RealTimePaymentsTransferListParamsCreatedAt] `query:"created_at"`
}

// URLQuery serializes RealTimePaymentsTransferListParams into a url.Values of the
// query parameters associated with this value
func (r RealTimePaymentsTransferListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type RealTimePaymentsTransferListParamsCreatedAt struct {
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

// URLQuery serializes RealTimePaymentsTransferListParamsCreatedAt into a
// url.Values of the query parameters associated with this value
func (r RealTimePaymentsTransferListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

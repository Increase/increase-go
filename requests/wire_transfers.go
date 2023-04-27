package requests

import (
	"net/url"
	"time"

	"github.com/increase/increase-go/internal/field"
	apijson "github.com/increase/increase-go/internal/json"
	"github.com/increase/increase-go/internal/query"
)

type WireTransferNewParams struct {
	// The identifier for the account that will send the transfer.
	AccountID field.Field[string] `json:"account_id,required"`
	// The account number for the destination account.
	AccountNumber field.Field[string] `json:"account_number"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
	// destination account.
	RoutingNumber field.Field[string] `json:"routing_number"`
	// The ID of an External Account to initiate a transfer to. If this parameter is
	// provided, `account_number` and `routing_number` must be absent.
	ExternalAccountID field.Field[string] `json:"external_account_id"`
	// The transfer amount in cents.
	Amount field.Field[int64] `json:"amount,required"`
	// The message that will show on the recipient's bank statement.
	MessageToRecipient field.Field[string] `json:"message_to_recipient,required"`
	// The beneficiary's name.
	BeneficiaryName field.Field[string] `json:"beneficiary_name,required"`
	// The beneficiary's address line 1.
	BeneficiaryAddressLine1 field.Field[string] `json:"beneficiary_address_line1"`
	// The beneficiary's address line 2.
	BeneficiaryAddressLine2 field.Field[string] `json:"beneficiary_address_line2"`
	// The beneficiary's address line 3.
	BeneficiaryAddressLine3 field.Field[string] `json:"beneficiary_address_line3"`
	// Whether the transfer requires explicit approval via the dashboard or API.
	RequireApproval field.Field[bool] `json:"require_approval"`
}

// MarshalJSON serializes WireTransferNewParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r WireTransferNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WireTransferListParams struct {
	// Return the page of entries after this one.
	Cursor field.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit field.Field[int64] `query:"limit"`
	// Filter Wire Transfers to those belonging to the specified Account.
	AccountID field.Field[string] `query:"account_id"`
	// Filter Wire Transfers to those made to the specified External Account.
	ExternalAccountID field.Field[string]                          `query:"external_account_id"`
	CreatedAt         field.Field[WireTransferListParamsCreatedAt] `query:"created_at"`
}

// URLQuery serializes WireTransferListParams into a url.Values of the query
// parameters associated with this value
func (r WireTransferListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type WireTransferListParamsCreatedAt struct {
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

// URLQuery serializes WireTransferListParamsCreatedAt into a url.Values of the
// query parameters associated with this value
func (r WireTransferListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

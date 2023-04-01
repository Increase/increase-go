package requests

import (
	"fmt"
	"net/url"
	"time"

	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/fields"
)

type CreateAWireTransferParameters struct {
	// The identifier for the account that will send the transfer.
	AccountID fields.Field[string] `json:"account_id,required"`
	// The account number for the destination account.
	AccountNumber fields.Field[string] `json:"account_number"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
	// destination account.
	RoutingNumber fields.Field[string] `json:"routing_number"`
	// The ID of an External Account to initiate a transfer to. If this parameter is
	// provided, `account_number` and `routing_number` must be absent.
	ExternalAccountID fields.Field[string] `json:"external_account_id"`
	// The transfer amount in cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The message that will show on the recipient's bank statement.
	MessageToRecipient fields.Field[string] `json:"message_to_recipient,required"`
	// The beneficiary's name.
	BeneficiaryName fields.Field[string] `json:"beneficiary_name,required"`
	// The beneficiary's address line 1.
	BeneficiaryAddressLine1 fields.Field[string] `json:"beneficiary_address_line1"`
	// The beneficiary's address line 2.
	BeneficiaryAddressLine2 fields.Field[string] `json:"beneficiary_address_line2"`
	// The beneficiary's address line 3.
	BeneficiaryAddressLine3 fields.Field[string] `json:"beneficiary_address_line3"`
	// Whether the transfer requires explicit approval via the dashboard or API.
	RequireApproval fields.Field[bool] `json:"require_approval"`
}

// MarshalJSON serializes CreateAWireTransferParameters into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *CreateAWireTransferParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CreateAWireTransferParameters) String() (result string) {
	return fmt.Sprintf("&CreateAWireTransferParameters{AccountID:%s AccountNumber:%s RoutingNumber:%s ExternalAccountID:%s Amount:%s MessageToRecipient:%s BeneficiaryName:%s BeneficiaryAddressLine1:%s BeneficiaryAddressLine2:%s BeneficiaryAddressLine3:%s RequireApproval:%s}", r.AccountID, r.AccountNumber, r.RoutingNumber, r.ExternalAccountID, r.Amount, r.MessageToRecipient, r.BeneficiaryName, r.BeneficiaryAddressLine1, r.BeneficiaryAddressLine2, r.BeneficiaryAddressLine3, r.RequireApproval)
}

type WireTransferListParams struct {
	// Return the page of entries after this one.
	Cursor fields.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit fields.Field[int64] `query:"limit"`
	// Filter Wire Transfers to those belonging to the specified Account.
	AccountID fields.Field[string] `query:"account_id"`
	// Filter Wire Transfers to those made to the specified External Account.
	ExternalAccountID fields.Field[string]                          `query:"external_account_id"`
	CreatedAt         fields.Field[WireTransferListParamsCreatedAt] `query:"created_at"`
}

// URLQuery serializes WireTransferListParams into a url.Values of the query
// parameters associated with this value
func (r *WireTransferListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r WireTransferListParams) String() (result string) {
	return fmt.Sprintf("&WireTransferListParams{Cursor:%s Limit:%s AccountID:%s ExternalAccountID:%s CreatedAt:%s}", r.Cursor, r.Limit, r.AccountID, r.ExternalAccountID, r.CreatedAt)
}

type WireTransferListParamsCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After fields.Field[time.Time] `query:"after" format:"date-time"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before fields.Field[time.Time] `query:"before" format:"date-time"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter fields.Field[time.Time] `query:"on_or_after" format:"date-time"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore fields.Field[time.Time] `query:"on_or_before" format:"date-time"`
}

// URLQuery serializes WireTransferListParamsCreatedAt into a url.Values of the
// query parameters associated with this value
func (r *WireTransferListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r WireTransferListParamsCreatedAt) String() (result string) {
	return fmt.Sprintf("&WireTransferListParamsCreatedAt{After:%s Before:%s OnOrAfter:%s OnOrBefore:%s}", r.After, r.Before, r.OnOrAfter, r.OnOrBefore)
}

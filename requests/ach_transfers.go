package requests

import (
	"net/url"
	"time"

	"github.com/increase/increase-go/internal/field"
	apijson "github.com/increase/increase-go/internal/json"
	"github.com/increase/increase-go/internal/query"
)

type ACHTransferNewParams struct {
	// The Increase identifier for the account that will send the transfer.
	AccountID field.Field[string] `json:"account_id,required"`
	// The account number for the destination account.
	AccountNumber field.Field[string] `json:"account_number"`
	// Additional information that will be sent to the recipient. This is included in
	// the transfer data sent to the receiving bank.
	Addendum field.Field[string] `json:"addendum"`
	// The transfer amount in cents. A positive amount originates a credit transfer
	// pushing funds to the receiving account. A negative amount originates a debit
	// transfer pulling funds from the receiving account.
	Amount field.Field[int64] `json:"amount,required"`
	// The description of the date of the transfer, usually in the format `YYMMDD`.
	// This is included in the transfer data sent to the receiving bank.
	CompanyDescriptiveDate field.Field[string] `json:"company_descriptive_date"`
	// The data you choose to associate with the transfer. This is included in the
	// transfer data sent to the receiving bank.
	CompanyDiscretionaryData field.Field[string] `json:"company_discretionary_data"`
	// A description of the transfer. This is included in the transfer data sent to the
	// receiving bank.
	CompanyEntryDescription field.Field[string] `json:"company_entry_description"`
	// The name by which the recipient knows you. This is included in the transfer data
	// sent to the receiving bank.
	CompanyName field.Field[string] `json:"company_name"`
	// The transfer effective date in
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	EffectiveDate field.Field[time.Time] `json:"effective_date" format:"date"`
	// The ID of an External Account to initiate a transfer to. If this parameter is
	// provided, `account_number`, `routing_number`, and `funding` must be absent.
	ExternalAccountID field.Field[string] `json:"external_account_id"`
	// The type of the account to which the transfer will be sent.
	Funding field.Field[ACHTransferNewParamsFunding] `json:"funding"`
	// Your identifer for the transfer recipient.
	IndividualID field.Field[string] `json:"individual_id"`
	// The name of the transfer recipient. This value is informational and not verified
	// by the recipient's bank.
	IndividualName field.Field[string] `json:"individual_name"`
	// Whether the transfer requires explicit approval via the dashboard or API.
	RequireApproval field.Field[bool] `json:"require_approval"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
	// destination account.
	RoutingNumber field.Field[string] `json:"routing_number"`
	// The Standard Entry Class (SEC) code to use for the transfer.
	StandardEntryClassCode field.Field[ACHTransferNewParamsStandardEntryClassCode] `json:"standard_entry_class_code"`
	// A description you choose to give the transfer. This will be saved with the
	// transfer details, displayed in the dashboard, and returned by the API. If
	// `individual_name` and `company_name` are not explicitly set by this API, the
	// `statement_descriptor` will be sent in those fields to the receiving bank to
	// help the customer recognize the transfer. You are highly encouraged to pass
	// `individual_name` and `company_name` instead of relying on this fallback.
	StatementDescriptor field.Field[string] `json:"statement_descriptor,required"`
}

// MarshalJSON serializes ACHTransferNewParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r ACHTransferNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ACHTransferNewParamsFunding string

const (
	ACHTransferNewParamsFundingChecking ACHTransferNewParamsFunding = "checking"
	ACHTransferNewParamsFundingSavings  ACHTransferNewParamsFunding = "savings"
)

type ACHTransferNewParamsStandardEntryClassCode string

const (
	ACHTransferNewParamsStandardEntryClassCodeCorporateCreditOrDebit        ACHTransferNewParamsStandardEntryClassCode = "corporate_credit_or_debit"
	ACHTransferNewParamsStandardEntryClassCodePrearrangedPaymentsAndDeposit ACHTransferNewParamsStandardEntryClassCode = "prearranged_payments_and_deposit"
	ACHTransferNewParamsStandardEntryClassCodeInternetInitiated             ACHTransferNewParamsStandardEntryClassCode = "internet_initiated"
)

type ACHTransferListParams struct {
	// Return the page of entries after this one.
	Cursor field.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit field.Field[int64] `query:"limit"`
	// Filter ACH Transfers to those that originated from the specified Account.
	AccountID field.Field[string] `query:"account_id"`
	// Filter ACH Transfers to those made to the specified External Account.
	ExternalAccountID field.Field[string]                         `query:"external_account_id"`
	CreatedAt         field.Field[ACHTransferListParamsCreatedAt] `query:"created_at"`
}

// URLQuery serializes ACHTransferListParams into a url.Values of the query
// parameters associated with this value
func (r ACHTransferListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type ACHTransferListParamsCreatedAt struct {
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

// URLQuery serializes ACHTransferListParamsCreatedAt into a url.Values of the
// query parameters associated with this value
func (r ACHTransferListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

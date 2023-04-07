package requests

import (
	"net/url"
	"time"

	"github.com/increase/increase-go/core/field"
	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
)

type ACHPrenotificationNewParams struct {
	// The account number for the destination account.
	AccountNumber field.Field[string] `json:"account_number,required"`
	// Additional information that will be sent to the recipient.
	Addendum field.Field[string] `json:"addendum"`
	// The description of the date of the transfer.
	CompanyDescriptiveDate field.Field[string] `json:"company_descriptive_date"`
	// The data you choose to associate with the transfer.
	CompanyDiscretionaryData field.Field[string] `json:"company_discretionary_data"`
	// The description of the transfer you wish to be shown to the recipient.
	CompanyEntryDescription field.Field[string] `json:"company_entry_description"`
	// The name by which the recipient knows you.
	CompanyName field.Field[string] `json:"company_name"`
	// Whether the Prenotification is for a future debit or credit.
	CreditDebitIndicator field.Field[ACHPrenotificationNewParamsCreditDebitIndicator] `json:"credit_debit_indicator"`
	// The transfer effective date in
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	EffectiveDate field.Field[time.Time] `json:"effective_date" format:"date"`
	// Your identifer for the transfer recipient.
	IndividualID field.Field[string] `json:"individual_id"`
	// The name of the transfer recipient. This value is information and not verified
	// by the recipient's bank.
	IndividualName field.Field[string] `json:"individual_name"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
	// destination account.
	RoutingNumber field.Field[string] `json:"routing_number,required"`
	// The Standard Entry Class (SEC) code to use for the ACH Prenotification.
	StandardEntryClassCode field.Field[ACHPrenotificationNewParamsStandardEntryClassCode] `json:"standard_entry_class_code"`
}

// MarshalJSON serializes ACHPrenotificationNewParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *ACHPrenotificationNewParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type ACHPrenotificationNewParamsCreditDebitIndicator string

const (
	ACHPrenotificationNewParamsCreditDebitIndicatorCredit ACHPrenotificationNewParamsCreditDebitIndicator = "credit"
	ACHPrenotificationNewParamsCreditDebitIndicatorDebit  ACHPrenotificationNewParamsCreditDebitIndicator = "debit"
)

type ACHPrenotificationNewParamsStandardEntryClassCode string

const (
	ACHPrenotificationNewParamsStandardEntryClassCodeCorporateCreditOrDebit        ACHPrenotificationNewParamsStandardEntryClassCode = "corporate_credit_or_debit"
	ACHPrenotificationNewParamsStandardEntryClassCodePrearrangedPaymentsAndDeposit ACHPrenotificationNewParamsStandardEntryClassCode = "prearranged_payments_and_deposit"
	ACHPrenotificationNewParamsStandardEntryClassCodeInternetInitiated             ACHPrenotificationNewParamsStandardEntryClassCode = "internet_initiated"
)

type ACHPrenotificationListParams struct {
	// Return the page of entries after this one.
	Cursor field.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit     field.Field[int64]                                 `query:"limit"`
	CreatedAt field.Field[ACHPrenotificationListParamsCreatedAt] `query:"created_at"`
}

// URLQuery serializes ACHPrenotificationListParams into a url.Values of the query
// parameters associated with this value
func (r *ACHPrenotificationListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type ACHPrenotificationListParamsCreatedAt struct {
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

// URLQuery serializes ACHPrenotificationListParamsCreatedAt into a url.Values of
// the query parameters associated with this value
func (r *ACHPrenotificationListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

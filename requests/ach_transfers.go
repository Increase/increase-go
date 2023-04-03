package requests

import (
	"fmt"
	"net/url"
	"time"

	"github.com/increase/increase-go/core/fields"
	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
)

type CreateAnACHTransferParameters struct {
	// The Increase identifier for the account that will send the transfer.
	AccountID fields.Field[string] `json:"account_id,required"`
	// The account number for the destination account.
	AccountNumber fields.Field[string] `json:"account_number"`
	// Additional information that will be sent to the recipient. This is included in
	// the transfer data sent to the receiving bank.
	Addendum fields.Field[string] `json:"addendum"`
	// The transfer amount in cents. A positive amount originates a credit transfer
	// pushing funds to the receiving account. A negative amount originates a debit
	// transfer pulling funds from the receiving account.
	Amount fields.Field[int64] `json:"amount,required"`
	// The description of the date of the transfer, usually in the format `YYYYMMDD`.
	// This is included in the transfer data sent to the receiving bank.
	CompanyDescriptiveDate fields.Field[string] `json:"company_descriptive_date"`
	// The data you choose to associate with the transfer. This is included in the
	// transfer data sent to the receiving bank.
	CompanyDiscretionaryData fields.Field[string] `json:"company_discretionary_data"`
	// A description of the transfer. This is included in the transfer data sent to the
	// receiving bank.
	CompanyEntryDescription fields.Field[string] `json:"company_entry_description"`
	// The name by which the recipient knows you. This is included in the transfer data
	// sent to the receiving bank.
	CompanyName fields.Field[string] `json:"company_name"`
	// The transfer effective date in
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	EffectiveDate fields.Field[time.Time] `json:"effective_date" format:"date"`
	// The ID of an External Account to initiate a transfer to. If this parameter is
	// provided, `account_number`, `routing_number`, and `funding` must be absent.
	ExternalAccountID fields.Field[string] `json:"external_account_id"`
	// The type of the account to which the transfer will be sent.
	Funding fields.Field[CreateAnACHTransferParametersFunding] `json:"funding"`
	// Your identifer for the transfer recipient.
	IndividualID fields.Field[string] `json:"individual_id"`
	// The name of the transfer recipient. This value is informational and not verified
	// by the recipient's bank.
	IndividualName fields.Field[string] `json:"individual_name"`
	// Whether the transfer requires explicit approval via the dashboard or API.
	RequireApproval fields.Field[bool] `json:"require_approval"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
	// destination account.
	RoutingNumber fields.Field[string] `json:"routing_number"`
	// The Standard Entry Class (SEC) code to use for the transfer.
	StandardEntryClassCode fields.Field[CreateAnACHTransferParametersStandardEntryClassCode] `json:"standard_entry_class_code"`
	// A description you choose to give the transfer. This will be saved with the
	// transfer details, displayed in the dashboard, and returned by the API. If
	// `individual_name` and `company_name` are not explicitly set by this API, the
	// `statement_descriptor` will be sent in those fields to the receiving bank to
	// help the customer recognize the transfer. You are highly encouraged to pass
	// `individual_name` and `company_name` instead of relying on this fallback.
	StatementDescriptor fields.Field[string] `json:"statement_descriptor,required"`
}

// MarshalJSON serializes CreateAnACHTransferParameters into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *CreateAnACHTransferParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CreateAnACHTransferParameters) String() (result string) {
	return fmt.Sprintf("&CreateAnACHTransferParameters{AccountID:%s AccountNumber:%s Addendum:%s Amount:%s CompanyDescriptiveDate:%s CompanyDiscretionaryData:%s CompanyEntryDescription:%s CompanyName:%s EffectiveDate:%s ExternalAccountID:%s Funding:%s IndividualID:%s IndividualName:%s RequireApproval:%s RoutingNumber:%s StandardEntryClassCode:%s StatementDescriptor:%s}", r.AccountID, r.AccountNumber, r.Addendum, r.Amount, r.CompanyDescriptiveDate, r.CompanyDiscretionaryData, r.CompanyEntryDescription, r.CompanyName, r.EffectiveDate, r.ExternalAccountID, r.Funding, r.IndividualID, r.IndividualName, r.RequireApproval, r.RoutingNumber, r.StandardEntryClassCode, r.StatementDescriptor)
}

type CreateAnACHTransferParametersFunding string

const (
	CreateAnACHTransferParametersFundingChecking CreateAnACHTransferParametersFunding = "checking"
	CreateAnACHTransferParametersFundingSavings  CreateAnACHTransferParametersFunding = "savings"
)

type CreateAnACHTransferParametersStandardEntryClassCode string

const (
	CreateAnACHTransferParametersStandardEntryClassCodeCorporateCreditOrDebit        CreateAnACHTransferParametersStandardEntryClassCode = "corporate_credit_or_debit"
	CreateAnACHTransferParametersStandardEntryClassCodePrearrangedPaymentsAndDeposit CreateAnACHTransferParametersStandardEntryClassCode = "prearranged_payments_and_deposit"
	CreateAnACHTransferParametersStandardEntryClassCodeInternetInitiated             CreateAnACHTransferParametersStandardEntryClassCode = "internet_initiated"
)

type ACHTransferListParams struct {
	// Return the page of entries after this one.
	Cursor fields.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit fields.Field[int64] `query:"limit"`
	// Filter ACH Transfers to those that originated from the specified Account.
	AccountID fields.Field[string] `query:"account_id"`
	// Filter ACH Transfers to those made to the specified External Account.
	ExternalAccountID fields.Field[string]                         `query:"external_account_id"`
	CreatedAt         fields.Field[ACHTransferListParamsCreatedAt] `query:"created_at"`
}

// URLQuery serializes ACHTransferListParams into a url.Values of the query
// parameters associated with this value
func (r *ACHTransferListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r ACHTransferListParams) String() (result string) {
	return fmt.Sprintf("&ACHTransferListParams{Cursor:%s Limit:%s AccountID:%s ExternalAccountID:%s CreatedAt:%s}", r.Cursor, r.Limit, r.AccountID, r.ExternalAccountID, r.CreatedAt)
}

type ACHTransferListParamsCreatedAt struct {
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

// URLQuery serializes ACHTransferListParamsCreatedAt into a url.Values of the
// query parameters associated with this value
func (r *ACHTransferListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r ACHTransferListParamsCreatedAt) String() (result string) {
	return fmt.Sprintf("&ACHTransferListParamsCreatedAt{After:%s Before:%s OnOrAfter:%s OnOrBefore:%s}", r.After, r.Before, r.OnOrAfter, r.OnOrBefore)
}

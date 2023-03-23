package requests

import (
	"fmt"
	"net/url"
	"time"

	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/fields"
)

type ACHPrenotification struct {
	// The ACH Prenotification's identifier.
	ID fields.Field[string] `json:"id,required"`
	// The destination account number.
	AccountNumber fields.Field[string] `json:"account_number,required"`
	// Additional information for the recipient.
	Addendum fields.Field[string] `json:"addendum,required,nullable"`
	// The description of the date of the notification.
	CompanyDescriptiveDate fields.Field[string] `json:"company_descriptive_date,required,nullable"`
	// Optional data associated with the notification.
	CompanyDiscretionaryData fields.Field[string] `json:"company_discretionary_data,required,nullable"`
	// The description of the notification.
	CompanyEntryDescription fields.Field[string] `json:"company_entry_description,required,nullable"`
	// The name by which you know the company.
	CompanyName fields.Field[string] `json:"company_name,required,nullable"`
	// If the notification is for a future credit or debit.
	CreditDebitIndicator fields.Field[ACHPrenotificationCreditDebitIndicator] `json:"credit_debit_indicator,required,nullable"`
	// The effective date in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	EffectiveDate fields.Field[time.Time] `json:"effective_date,required,nullable" format:"date-time"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber fields.Field[string] `json:"routing_number,required"`
	// If your prenotification is returned, this will contain details of the return.
	PrenotificationReturn fields.Field[ACHPrenotificationPrenotificationReturn] `json:"prenotification_return,required,nullable"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the prenotification was created.
	CreatedAt fields.Field[time.Time] `json:"created_at,required" format:"date-time"`
	// The lifecycle status of the ACH Prenotification.
	Status fields.Field[ACHPrenotificationStatus] `json:"status,required"`
	// A constant representing the object's type. For this resource it will always be
	// `ach_prenotification`.
	Type fields.Field[ACHPrenotificationType] `json:"type,required"`
}

// MarshalJSON serializes ACHPrenotification into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *ACHPrenotification) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHPrenotification) String() (result string) {
	return fmt.Sprintf("&ACHPrenotification{ID:%s AccountNumber:%s Addendum:%s CompanyDescriptiveDate:%s CompanyDiscretionaryData:%s CompanyEntryDescription:%s CompanyName:%s CreditDebitIndicator:%s EffectiveDate:%s RoutingNumber:%s PrenotificationReturn:%s CreatedAt:%s Status:%s Type:%s}", r.ID, r.AccountNumber, r.Addendum, r.CompanyDescriptiveDate, r.CompanyDiscretionaryData, r.CompanyEntryDescription, r.CompanyName, r.CreditDebitIndicator, r.EffectiveDate, r.RoutingNumber, r.PrenotificationReturn, r.CreatedAt, r.Status, r.Type)
}

type ACHPrenotificationCreditDebitIndicator string

const (
	ACHPrenotificationCreditDebitIndicatorCredit ACHPrenotificationCreditDebitIndicator = "credit"
	ACHPrenotificationCreditDebitIndicatorDebit  ACHPrenotificationCreditDebitIndicator = "debit"
)

type ACHPrenotificationPrenotificationReturn struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Prenotification was returned.
	CreatedAt fields.Field[time.Time] `json:"created_at,required" format:"date-time"`
	// Why the Prenotification was returned.
	ReturnReasonCode fields.Field[string] `json:"return_reason_code,required"`
}

// MarshalJSON serializes ACHPrenotificationPrenotificationReturn into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *ACHPrenotificationPrenotificationReturn) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ACHPrenotificationPrenotificationReturn) String() (result string) {
	return fmt.Sprintf("&ACHPrenotificationPrenotificationReturn{CreatedAt:%s ReturnReasonCode:%s}", r.CreatedAt, r.ReturnReasonCode)
}

type ACHPrenotificationStatus string

const (
	ACHPrenotificationStatusPendingSubmitting ACHPrenotificationStatus = "pending_submitting"
	ACHPrenotificationStatusRequiresAttention ACHPrenotificationStatus = "requires_attention"
	ACHPrenotificationStatusReturned          ACHPrenotificationStatus = "returned"
	ACHPrenotificationStatusSubmitted         ACHPrenotificationStatus = "submitted"
)

type ACHPrenotificationType string

const (
	ACHPrenotificationTypeACHPrenotification ACHPrenotificationType = "ach_prenotification"
)

type CreateAnACHPrenotificationParameters struct {
	// The account number for the destination account.
	AccountNumber fields.Field[string] `json:"account_number,required"`
	// Additional information that will be sent to the recipient.
	Addendum fields.Field[string] `json:"addendum"`
	// The description of the date of the transfer.
	CompanyDescriptiveDate fields.Field[string] `json:"company_descriptive_date"`
	// The data you choose to associate with the transfer.
	CompanyDiscretionaryData fields.Field[string] `json:"company_discretionary_data"`
	// The description of the transfer you wish to be shown to the recipient.
	CompanyEntryDescription fields.Field[string] `json:"company_entry_description"`
	// The name by which the recipient knows you.
	CompanyName fields.Field[string] `json:"company_name"`
	// Whether the Prenotification is for a future debit or credit.
	CreditDebitIndicator fields.Field[CreateAnACHPrenotificationParametersCreditDebitIndicator] `json:"credit_debit_indicator"`
	// The transfer effective date in
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	EffectiveDate fields.Field[time.Time] `json:"effective_date" format:"date"`
	// Your identifer for the transfer recipient.
	IndividualID fields.Field[string] `json:"individual_id"`
	// The name of the transfer recipient. This value is information and not verified
	// by the recipient's bank.
	IndividualName fields.Field[string] `json:"individual_name"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
	// destination account.
	RoutingNumber fields.Field[string] `json:"routing_number,required"`
	// The Standard Entry Class (SEC) code to use for the ACH Prenotification.
	StandardEntryClassCode fields.Field[CreateAnACHPrenotificationParametersStandardEntryClassCode] `json:"standard_entry_class_code"`
}

// MarshalJSON serializes CreateAnACHPrenotificationParameters into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *CreateAnACHPrenotificationParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CreateAnACHPrenotificationParameters) String() (result string) {
	return fmt.Sprintf("&CreateAnACHPrenotificationParameters{AccountNumber:%s Addendum:%s CompanyDescriptiveDate:%s CompanyDiscretionaryData:%s CompanyEntryDescription:%s CompanyName:%s CreditDebitIndicator:%s EffectiveDate:%s IndividualID:%s IndividualName:%s RoutingNumber:%s StandardEntryClassCode:%s}", r.AccountNumber, r.Addendum, r.CompanyDescriptiveDate, r.CompanyDiscretionaryData, r.CompanyEntryDescription, r.CompanyName, r.CreditDebitIndicator, r.EffectiveDate, r.IndividualID, r.IndividualName, r.RoutingNumber, r.StandardEntryClassCode)
}

type CreateAnACHPrenotificationParametersCreditDebitIndicator string

const (
	CreateAnACHPrenotificationParametersCreditDebitIndicatorCredit CreateAnACHPrenotificationParametersCreditDebitIndicator = "credit"
	CreateAnACHPrenotificationParametersCreditDebitIndicatorDebit  CreateAnACHPrenotificationParametersCreditDebitIndicator = "debit"
)

type CreateAnACHPrenotificationParametersStandardEntryClassCode string

const (
	CreateAnACHPrenotificationParametersStandardEntryClassCodeCorporateCreditOrDebit        CreateAnACHPrenotificationParametersStandardEntryClassCode = "corporate_credit_or_debit"
	CreateAnACHPrenotificationParametersStandardEntryClassCodePrearrangedPaymentsAndDeposit CreateAnACHPrenotificationParametersStandardEntryClassCode = "prearranged_payments_and_deposit"
	CreateAnACHPrenotificationParametersStandardEntryClassCodeInternetInitiated             CreateAnACHPrenotificationParametersStandardEntryClassCode = "internet_initiated"
)

type ACHPrenotificationListParams struct {
	// Return the page of entries after this one.
	Cursor fields.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit     fields.Field[int64]                                 `query:"limit"`
	CreatedAt fields.Field[ACHPrenotificationListParamsCreatedAt] `query:"created_at"`
}

// URLQuery serializes ACHPrenotificationListParams into a url.Values of the query
// parameters associated with this value
func (r *ACHPrenotificationListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r ACHPrenotificationListParams) String() (result string) {
	return fmt.Sprintf("&ACHPrenotificationListParams{Cursor:%s Limit:%s CreatedAt:%s}", r.Cursor, r.Limit, r.CreatedAt)
}

type ACHPrenotificationListParamsCreatedAt struct {
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

// URLQuery serializes ACHPrenotificationListParamsCreatedAt into a url.Values of
// the query parameters associated with this value
func (r *ACHPrenotificationListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r ACHPrenotificationListParamsCreatedAt) String() (result string) {
	return fmt.Sprintf("&ACHPrenotificationListParamsCreatedAt{After:%s Before:%s OnOrAfter:%s OnOrBefore:%s}", r.After, r.Before, r.OnOrAfter, r.OnOrBefore)
}

package types

import (
	"fmt"
	"net/url"

	"github.com/increase/increase-go/core"
	"github.com/increase/increase-go/core/pjson"
	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/pagination"
)

type ACHPrenotification struct {
	// The ACH Prenotification's identifier.
	ID *string `pjson:"id"`
	// The destination account number.
	AccountNumber *string `pjson:"account_number"`
	// Additional information for the recipient.
	Addendum *string `pjson:"addendum"`
	// The description of the date of the notification.
	CompanyDescriptiveDate *string `pjson:"company_descriptive_date"`
	// Optional data associated with the notification.
	CompanyDiscretionaryData *string `pjson:"company_discretionary_data"`
	// The description of the notification.
	CompanyEntryDescription *string `pjson:"company_entry_description"`
	// The name by which you know the company.
	CompanyName *string `pjson:"company_name"`
	// If the notification is for a future credit or debit.
	CreditDebitIndicator *ACHPrenotificationCreditDebitIndicator `pjson:"credit_debit_indicator"`
	// The effective date in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	EffectiveDate *string `pjson:"effective_date"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber *string `pjson:"routing_number"`
	// If your prenotification is returned, this will contain details of the return.
	PrenotificationReturn *ACHPrenotificationPrenotificationReturn `pjson:"prenotification_return"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the prenotification was created.
	CreatedAt *string `pjson:"created_at"`
	// The lifecycle status of the ACH Prenotification.
	Status *ACHPrenotificationStatus `pjson:"status"`
	// A constant representing the object's type. For this resource it will always be
	// `ach_prenotification`.
	Type       *ACHPrenotificationType `pjson:"type"`
	jsonFields map[string]interface{}  `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into ACHPrenotification using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ACHPrenotification) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes ACHPrenotification into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *ACHPrenotification) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The ACH Prenotification's identifier.
func (r ACHPrenotification) GetID() (ID string) {
	if r.ID != nil {
		ID = *r.ID
	}
	return
}

// The destination account number.
func (r ACHPrenotification) GetAccountNumber() (AccountNumber string) {
	if r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

// Additional information for the recipient.
func (r ACHPrenotification) GetAddendum() (Addendum string) {
	if r.Addendum != nil {
		Addendum = *r.Addendum
	}
	return
}

// The description of the date of the notification.
func (r ACHPrenotification) GetCompanyDescriptiveDate() (CompanyDescriptiveDate string) {
	if r.CompanyDescriptiveDate != nil {
		CompanyDescriptiveDate = *r.CompanyDescriptiveDate
	}
	return
}

// Optional data associated with the notification.
func (r ACHPrenotification) GetCompanyDiscretionaryData() (CompanyDiscretionaryData string) {
	if r.CompanyDiscretionaryData != nil {
		CompanyDiscretionaryData = *r.CompanyDiscretionaryData
	}
	return
}

// The description of the notification.
func (r ACHPrenotification) GetCompanyEntryDescription() (CompanyEntryDescription string) {
	if r.CompanyEntryDescription != nil {
		CompanyEntryDescription = *r.CompanyEntryDescription
	}
	return
}

// The name by which you know the company.
func (r ACHPrenotification) GetCompanyName() (CompanyName string) {
	if r.CompanyName != nil {
		CompanyName = *r.CompanyName
	}
	return
}

// If the notification is for a future credit or debit.
func (r ACHPrenotification) GetCreditDebitIndicator() (CreditDebitIndicator ACHPrenotificationCreditDebitIndicator) {
	if r.CreditDebitIndicator != nil {
		CreditDebitIndicator = *r.CreditDebitIndicator
	}
	return
}

// The effective date in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
func (r ACHPrenotification) GetEffectiveDate() (EffectiveDate string) {
	if r.EffectiveDate != nil {
		EffectiveDate = *r.EffectiveDate
	}
	return
}

// The American Bankers' Association (ABA) Routing Transit Number (RTN).
func (r ACHPrenotification) GetRoutingNumber() (RoutingNumber string) {
	if r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

// If your prenotification is returned, this will contain details of the return.
func (r ACHPrenotification) GetPrenotificationReturn() (PrenotificationReturn ACHPrenotificationPrenotificationReturn) {
	if r.PrenotificationReturn != nil {
		PrenotificationReturn = *r.PrenotificationReturn
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the prenotification was created.
func (r ACHPrenotification) GetCreatedAt() (CreatedAt string) {
	if r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The lifecycle status of the ACH Prenotification.
func (r ACHPrenotification) GetStatus() (Status ACHPrenotificationStatus) {
	if r.Status != nil {
		Status = *r.Status
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `ach_prenotification`.
func (r ACHPrenotification) GetType() (Type ACHPrenotificationType) {
	if r.Type != nil {
		Type = *r.Type
	}
	return
}

func (r ACHPrenotification) String() (result string) {
	return fmt.Sprintf("&ACHPrenotification{ID:%s AccountNumber:%s Addendum:%s CompanyDescriptiveDate:%s CompanyDiscretionaryData:%s CompanyEntryDescription:%s CompanyName:%s CreditDebitIndicator:%s EffectiveDate:%s RoutingNumber:%s PrenotificationReturn:%s CreatedAt:%s Status:%s Type:%s}", core.FmtP(r.ID), core.FmtP(r.AccountNumber), core.FmtP(r.Addendum), core.FmtP(r.CompanyDescriptiveDate), core.FmtP(r.CompanyDiscretionaryData), core.FmtP(r.CompanyEntryDescription), core.FmtP(r.CompanyName), core.FmtP(r.CreditDebitIndicator), core.FmtP(r.EffectiveDate), core.FmtP(r.RoutingNumber), r.PrenotificationReturn, core.FmtP(r.CreatedAt), core.FmtP(r.Status), core.FmtP(r.Type))
}

type ACHPrenotificationCreditDebitIndicator string

const (
	ACHPrenotificationCreditDebitIndicatorCredit ACHPrenotificationCreditDebitIndicator = "credit"
	ACHPrenotificationCreditDebitIndicatorDebit  ACHPrenotificationCreditDebitIndicator = "debit"
)

type ACHPrenotificationPrenotificationReturn struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Prenotification was returned.
	CreatedAt *string `pjson:"created_at"`
	// Why the Prenotification was returned.
	ReturnReasonCode *string                `pjson:"return_reason_code"`
	jsonFields       map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// ACHPrenotificationPrenotificationReturn using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHPrenotificationPrenotificationReturn) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes ACHPrenotificationPrenotificationReturn into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *ACHPrenotificationPrenotificationReturn) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the Prenotification was returned.
func (r ACHPrenotificationPrenotificationReturn) GetCreatedAt() (CreatedAt string) {
	if r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// Why the Prenotification was returned.
func (r ACHPrenotificationPrenotificationReturn) GetReturnReasonCode() (ReturnReasonCode string) {
	if r.ReturnReasonCode != nil {
		ReturnReasonCode = *r.ReturnReasonCode
	}
	return
}

func (r ACHPrenotificationPrenotificationReturn) String() (result string) {
	return fmt.Sprintf("&ACHPrenotificationPrenotificationReturn{CreatedAt:%s ReturnReasonCode:%s}", core.FmtP(r.CreatedAt), core.FmtP(r.ReturnReasonCode))
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
	AccountNumber *string `pjson:"account_number"`
	// Additional information that will be sent to the recipient.
	Addendum *string `pjson:"addendum"`
	// The description of the date of the transfer.
	CompanyDescriptiveDate *string `pjson:"company_descriptive_date"`
	// The data you choose to associate with the transfer.
	CompanyDiscretionaryData *string `pjson:"company_discretionary_data"`
	// The description of the transfer you wish to be shown to the recipient.
	CompanyEntryDescription *string `pjson:"company_entry_description"`
	// The name by which the recipient knows you.
	CompanyName *string `pjson:"company_name"`
	// Whether the Prenotification is for a future debit or credit.
	CreditDebitIndicator *CreateAnACHPrenotificationParametersCreditDebitIndicator `pjson:"credit_debit_indicator"`
	// The transfer effective date in
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	EffectiveDate *string `pjson:"effective_date"`
	// Your identifer for the transfer recipient.
	IndividualID *string `pjson:"individual_id"`
	// The name of the transfer recipient. This value is information and not verified
	// by the recipient's bank.
	IndividualName *string `pjson:"individual_name"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
	// destination account.
	RoutingNumber *string `pjson:"routing_number"`
	// The Standard Entry Class (SEC) code to use for the ACH Prenotification.
	StandardEntryClassCode *CreateAnACHPrenotificationParametersStandardEntryClassCode `pjson:"standard_entry_class_code"`
	jsonFields             map[string]interface{}                                      `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateAnACHPrenotificationParameters using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *CreateAnACHPrenotificationParameters) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CreateAnACHPrenotificationParameters into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *CreateAnACHPrenotificationParameters) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The account number for the destination account.
func (r CreateAnACHPrenotificationParameters) GetAccountNumber() (AccountNumber string) {
	if r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

// Additional information that will be sent to the recipient.
func (r CreateAnACHPrenotificationParameters) GetAddendum() (Addendum string) {
	if r.Addendum != nil {
		Addendum = *r.Addendum
	}
	return
}

// The description of the date of the transfer.
func (r CreateAnACHPrenotificationParameters) GetCompanyDescriptiveDate() (CompanyDescriptiveDate string) {
	if r.CompanyDescriptiveDate != nil {
		CompanyDescriptiveDate = *r.CompanyDescriptiveDate
	}
	return
}

// The data you choose to associate with the transfer.
func (r CreateAnACHPrenotificationParameters) GetCompanyDiscretionaryData() (CompanyDiscretionaryData string) {
	if r.CompanyDiscretionaryData != nil {
		CompanyDiscretionaryData = *r.CompanyDiscretionaryData
	}
	return
}

// The description of the transfer you wish to be shown to the recipient.
func (r CreateAnACHPrenotificationParameters) GetCompanyEntryDescription() (CompanyEntryDescription string) {
	if r.CompanyEntryDescription != nil {
		CompanyEntryDescription = *r.CompanyEntryDescription
	}
	return
}

// The name by which the recipient knows you.
func (r CreateAnACHPrenotificationParameters) GetCompanyName() (CompanyName string) {
	if r.CompanyName != nil {
		CompanyName = *r.CompanyName
	}
	return
}

// Whether the Prenotification is for a future debit or credit.
func (r CreateAnACHPrenotificationParameters) GetCreditDebitIndicator() (CreditDebitIndicator CreateAnACHPrenotificationParametersCreditDebitIndicator) {
	if r.CreditDebitIndicator != nil {
		CreditDebitIndicator = *r.CreditDebitIndicator
	}
	return
}

// The transfer effective date in
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
func (r CreateAnACHPrenotificationParameters) GetEffectiveDate() (EffectiveDate string) {
	if r.EffectiveDate != nil {
		EffectiveDate = *r.EffectiveDate
	}
	return
}

// Your identifer for the transfer recipient.
func (r CreateAnACHPrenotificationParameters) GetIndividualID() (IndividualID string) {
	if r.IndividualID != nil {
		IndividualID = *r.IndividualID
	}
	return
}

// The name of the transfer recipient. This value is information and not verified
// by the recipient's bank.
func (r CreateAnACHPrenotificationParameters) GetIndividualName() (IndividualName string) {
	if r.IndividualName != nil {
		IndividualName = *r.IndividualName
	}
	return
}

// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
// destination account.
func (r CreateAnACHPrenotificationParameters) GetRoutingNumber() (RoutingNumber string) {
	if r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

// The Standard Entry Class (SEC) code to use for the ACH Prenotification.
func (r CreateAnACHPrenotificationParameters) GetStandardEntryClassCode() (StandardEntryClassCode CreateAnACHPrenotificationParametersStandardEntryClassCode) {
	if r.StandardEntryClassCode != nil {
		StandardEntryClassCode = *r.StandardEntryClassCode
	}
	return
}

func (r CreateAnACHPrenotificationParameters) String() (result string) {
	return fmt.Sprintf("&CreateAnACHPrenotificationParameters{AccountNumber:%s Addendum:%s CompanyDescriptiveDate:%s CompanyDiscretionaryData:%s CompanyEntryDescription:%s CompanyName:%s CreditDebitIndicator:%s EffectiveDate:%s IndividualID:%s IndividualName:%s RoutingNumber:%s StandardEntryClassCode:%s}", core.FmtP(r.AccountNumber), core.FmtP(r.Addendum), core.FmtP(r.CompanyDescriptiveDate), core.FmtP(r.CompanyDiscretionaryData), core.FmtP(r.CompanyEntryDescription), core.FmtP(r.CompanyName), core.FmtP(r.CreditDebitIndicator), core.FmtP(r.EffectiveDate), core.FmtP(r.IndividualID), core.FmtP(r.IndividualName), core.FmtP(r.RoutingNumber), core.FmtP(r.StandardEntryClassCode))
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
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit      *int64                                  `query:"limit"`
	CreatedAt  *ACHPrenotificationsListParamsCreatedAt `query:"created_at"`
	jsonFields map[string]interface{}                  `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into ACHPrenotificationListParams
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *ACHPrenotificationListParams) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes ACHPrenotificationListParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *ACHPrenotificationListParams) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes ACHPrenotificationListParams into a url.Values of the query
// parameters associated with this value
func (r *ACHPrenotificationListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// Return the page of entries after this one.
func (r ACHPrenotificationListParams) GetCursor() (Cursor string) {
	if r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r ACHPrenotificationListParams) GetLimit() (Limit int64) {
	if r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

func (r ACHPrenotificationListParams) GetCreatedAt() (CreatedAt ACHPrenotificationsListParamsCreatedAt) {
	if r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

func (r ACHPrenotificationListParams) String() (result string) {
	return fmt.Sprintf("&ACHPrenotificationListParams{Cursor:%s Limit:%s CreatedAt:%s}", core.FmtP(r.Cursor), core.FmtP(r.Limit), r.CreatedAt)
}

type ACHPrenotificationsListParamsCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After *string `pjson:"after"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before *string `pjson:"before"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter *string `pjson:"on_or_after"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore *string                `pjson:"on_or_before"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// ACHPrenotificationsListParamsCreatedAt using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHPrenotificationsListParamsCreatedAt) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes ACHPrenotificationsListParamsCreatedAt into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *ACHPrenotificationsListParamsCreatedAt) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes ACHPrenotificationsListParamsCreatedAt into a url.Values of
// the query parameters associated with this value
func (r *ACHPrenotificationsListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r ACHPrenotificationsListParamsCreatedAt) GetAfter() (After string) {
	if r.After != nil {
		After = *r.After
	}
	return
}

// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r ACHPrenotificationsListParamsCreatedAt) GetBefore() (Before string) {
	if r.Before != nil {
		Before = *r.Before
	}
	return
}

// Return results on or after this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r ACHPrenotificationsListParamsCreatedAt) GetOnOrAfter() (OnOrAfter string) {
	if r.OnOrAfter != nil {
		OnOrAfter = *r.OnOrAfter
	}
	return
}

// Return results on or before this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r ACHPrenotificationsListParamsCreatedAt) GetOnOrBefore() (OnOrBefore string) {
	if r.OnOrBefore != nil {
		OnOrBefore = *r.OnOrBefore
	}
	return
}

func (r ACHPrenotificationsListParamsCreatedAt) String() (result string) {
	return fmt.Sprintf("&ACHPrenotificationsListParamsCreatedAt{After:%s Before:%s OnOrAfter:%s OnOrBefore:%s}", core.FmtP(r.After), core.FmtP(r.Before), core.FmtP(r.OnOrAfter), core.FmtP(r.OnOrBefore))
}

type ACHPrenotificationList struct {
	// The contents of the list.
	Data *[]ACHPrenotification `pjson:"data"`
	// A pointer to a place in the list.
	NextCursor *string                `pjson:"next_cursor"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into ACHPrenotificationList using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ACHPrenotificationList) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes ACHPrenotificationList into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *ACHPrenotificationList) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes ACHPrenotificationList into a url.Values of the query
// parameters associated with this value
func (r *ACHPrenotificationList) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// The contents of the list.
func (r ACHPrenotificationList) GetData() (Data []ACHPrenotification) {
	if r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r ACHPrenotificationList) GetNextCursor() (NextCursor string) {
	if r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

func (r ACHPrenotificationList) String() (result string) {
	return fmt.Sprintf("&ACHPrenotificationList{Data:%s NextCursor:%s}", core.Fmt(r.Data), core.FmtP(r.NextCursor))
}

type ACHPrenotificationsPage struct {
	*pagination.Page[ACHPrenotification]
}

func (r *ACHPrenotificationsPage) ACHPrenotification() *ACHPrenotification {
	return r.Current()
}

func (r *ACHPrenotificationsPage) NextPage() (*ACHPrenotificationsPage, error) {
	if page, err := r.Page.NextPage(); err != nil {
		return nil, err
	} else {
		return &ACHPrenotificationsPage{page}, nil
	}
}

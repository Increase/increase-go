package responses

import (
	"time"

	apijson "github.com/increase/increase-go/core/json"
)

type ACHPrenotification struct {
	// The ACH Prenotification's identifier.
	ID string `json:"id,required"`
	// The destination account number.
	AccountNumber string `json:"account_number,required"`
	// Additional information for the recipient.
	Addendum string `json:"addendum,required,nullable"`
	// The description of the date of the notification.
	CompanyDescriptiveDate string `json:"company_descriptive_date,required,nullable"`
	// Optional data associated with the notification.
	CompanyDiscretionaryData string `json:"company_discretionary_data,required,nullable"`
	// The description of the notification.
	CompanyEntryDescription string `json:"company_entry_description,required,nullable"`
	// The name by which you know the company.
	CompanyName string `json:"company_name,required,nullable"`
	// If the notification is for a future credit or debit.
	CreditDebitIndicator ACHPrenotificationCreditDebitIndicator `json:"credit_debit_indicator,required,nullable"`
	// The effective date in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	EffectiveDate time.Time `json:"effective_date,required,nullable" format:"date-time"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber string `json:"routing_number,required"`
	// If your prenotification is returned, this will contain details of the return.
	PrenotificationReturn ACHPrenotificationPrenotificationReturn `json:"prenotification_return,required,nullable"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the prenotification was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The lifecycle status of the ACH Prenotification.
	Status ACHPrenotificationStatus `json:"status,required"`
	// A constant representing the object's type. For this resource it will always be
	// `ach_prenotification`.
	Type ACHPrenotificationType `json:"type,required"`
	JSON ACHPrenotificationJSON
}

type ACHPrenotificationJSON struct {
	ID                       apijson.Metadata
	AccountNumber            apijson.Metadata
	Addendum                 apijson.Metadata
	CompanyDescriptiveDate   apijson.Metadata
	CompanyDiscretionaryData apijson.Metadata
	CompanyEntryDescription  apijson.Metadata
	CompanyName              apijson.Metadata
	CreditDebitIndicator     apijson.Metadata
	EffectiveDate            apijson.Metadata
	RoutingNumber            apijson.Metadata
	PrenotificationReturn    apijson.Metadata
	CreatedAt                apijson.Metadata
	Status                   apijson.Metadata
	Type                     apijson.Metadata
	Raw                      []byte
	Extras                   map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into ACHPrenotification using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ACHPrenotification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ACHPrenotificationCreditDebitIndicator string

const (
	ACHPrenotificationCreditDebitIndicatorCredit ACHPrenotificationCreditDebitIndicator = "credit"
	ACHPrenotificationCreditDebitIndicatorDebit  ACHPrenotificationCreditDebitIndicator = "debit"
)

type ACHPrenotificationPrenotificationReturn struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Prenotification was returned.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Why the Prenotification was returned.
	ReturnReasonCode string `json:"return_reason_code,required"`
	JSON             ACHPrenotificationPrenotificationReturnJSON
}

type ACHPrenotificationPrenotificationReturnJSON struct {
	CreatedAt        apijson.Metadata
	ReturnReasonCode apijson.Metadata
	Raw              []byte
	Extras           map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// ACHPrenotificationPrenotificationReturn using the internal json library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *ACHPrenotificationPrenotificationReturn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

type ACHPrenotificationListResponse struct {
	// The contents of the list.
	Data []ACHPrenotification `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       ACHPrenotificationListResponseJSON
}

type ACHPrenotificationListResponseJSON struct {
	Data       apijson.Metadata
	NextCursor apijson.Metadata
	Raw        []byte
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// ACHPrenotificationListResponse using the internal json library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *ACHPrenotificationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

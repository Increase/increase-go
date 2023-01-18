package ach_prenotifications

import "context"
import "increase/core"
import "fmt"
import "increase/pagination"

type ACHPrenotificationService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewACHPrenotificationService(requester core.Requester) (r *ACHPrenotificationService) {
	r = &ACHPrenotificationService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

//
type ACHPrenotification struct {
	// The ACH Prenotification's identifier.
	ID string `json:"id"`
	// The destination account number.
	AccountNumber string `json:"account_number"`
	// Additional information for the recipient.
	Addendum *string `json:"addendum"`
	// The description of the date of the notification.
	CompanyDescriptiveDate *string `json:"company_descriptive_date"`
	// Optional data associated with the notification.
	CompanyDiscretionaryData *string `json:"company_discretionary_data"`
	// The description of the notification.
	CompanyEntryDescription *string `json:"company_entry_description"`
	// The name by which you know the company.
	CompanyName *string `json:"company_name"`
	// If the notification is for a future credit or debit.
	CreditDebitIndicator *ACHPrenotificationCreditDebitIndicator `json:"credit_debit_indicator"`
	// The effective date in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	EffectiveDate *string `json:"effective_date"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber string `json:"routing_number"`
	// If your prenotification is returned, this will contain details of the return.
	PrenotificationReturn *ACHPrenotificationPrenotificationReturn `json:"prenotification_return"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the prenotification was created.
	CreatedAt string `json:"created_at"`
	// The lifecycle status of the ACH Prenotification.
	Status ACHPrenotificationStatus `json:"status"`
	// A constant representing the object's type. For this resource it will always be
	// `ach_prenotification`.
	Type ACHPrenotificationType `json:"type"`
}

// Additional information for the recipient.
func (r *ACHPrenotification) GetAddendum() (Addendum string) {
	if r != nil && r.Addendum != nil {
		Addendum = *r.Addendum
	}
	return
}

// The description of the date of the notification.
func (r *ACHPrenotification) GetCompanyDescriptiveDate() (CompanyDescriptiveDate string) {
	if r != nil && r.CompanyDescriptiveDate != nil {
		CompanyDescriptiveDate = *r.CompanyDescriptiveDate
	}
	return
}

// Optional data associated with the notification.
func (r *ACHPrenotification) GetCompanyDiscretionaryData() (CompanyDiscretionaryData string) {
	if r != nil && r.CompanyDiscretionaryData != nil {
		CompanyDiscretionaryData = *r.CompanyDiscretionaryData
	}
	return
}

// The description of the notification.
func (r *ACHPrenotification) GetCompanyEntryDescription() (CompanyEntryDescription string) {
	if r != nil && r.CompanyEntryDescription != nil {
		CompanyEntryDescription = *r.CompanyEntryDescription
	}
	return
}

// The name by which you know the company.
func (r *ACHPrenotification) GetCompanyName() (CompanyName string) {
	if r != nil && r.CompanyName != nil {
		CompanyName = *r.CompanyName
	}
	return
}

// If the notification is for a future credit or debit.
func (r *ACHPrenotification) GetCreditDebitIndicator() (CreditDebitIndicator ACHPrenotificationCreditDebitIndicator) {
	if r != nil && r.CreditDebitIndicator != nil {
		CreditDebitIndicator = *r.CreditDebitIndicator
	}
	return
}

// The effective date in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
func (r *ACHPrenotification) GetEffectiveDate() (EffectiveDate string) {
	if r != nil && r.EffectiveDate != nil {
		EffectiveDate = *r.EffectiveDate
	}
	return
}

// If your prenotification is returned, this will contain details of the return.
func (r *ACHPrenotification) GetPrenotificationReturn() (PrenotificationReturn ACHPrenotificationPrenotificationReturn) {
	if r != nil && r.PrenotificationReturn != nil {
		PrenotificationReturn = *r.PrenotificationReturn
	}
	return
}

type ACHPrenotificationCreditDebitIndicator string

const (
	ACHPrenotificationCreditDebitIndicatorCredit ACHPrenotificationCreditDebitIndicator = "credit"
	ACHPrenotificationCreditDebitIndicatorDebit  ACHPrenotificationCreditDebitIndicator = "debit"
)

//
type ACHPrenotificationPrenotificationReturn struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Prenotification was returned.
	CreatedAt string `json:"created_at"`
	// Why the Prenotification was returned.
	ReturnReasonCode string `json:"return_reason_code"`
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
	AccountNumber string `json:"account_number"`
	// Additional information that will be sent to the recipient.
	Addendum string `json:"addendum,omitempty"`
	// The description of the date of the transfer.
	CompanyDescriptiveDate string `json:"company_descriptive_date,omitempty"`
	// The data you choose to associate with the transfer.
	CompanyDiscretionaryData string `json:"company_discretionary_data,omitempty"`
	// The description of the transfer you wish to be shown to the recipient.
	CompanyEntryDescription string `json:"company_entry_description,omitempty"`
	// The name by which the recipient knows you.
	CompanyName string `json:"company_name,omitempty"`
	// Whether the Prenotification is for a future debit or credit.
	CreditDebitIndicator CreateAnACHPrenotificationParametersCreditDebitIndicator `json:"credit_debit_indicator,omitempty"`
	// The transfer effective date in
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	EffectiveDate string `json:"effective_date,omitempty"`
	// Your identifer for the transfer recipient.
	IndividualID string `json:"individual_id,omitempty"`
	// The name of the transfer recipient. This value is information and not verified
	// by the recipient's bank.
	IndividualName string `json:"individual_name,omitempty"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
	// destination account.
	RoutingNumber string `json:"routing_number"`
	// The Standard Entry Class (SEC) code to use for the ACH Prenotification.
	StandardEntryClassCode CreateAnACHPrenotificationParametersStandardEntryClassCode `json:"standard_entry_class_code,omitempty"`
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

type ListACHPrenotificationsQuery struct {
	// Return the page of entries after this one.
	Cursor string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit     int                                   `query:"limit"`
	CreatedAt ListACHPrenotificationsQueryCreatedAt `query:"created_at"`
}

type ListACHPrenotificationsQueryCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After string `json:"after,omitempty"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before string `json:"before,omitempty"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter string `json:"on_or_after,omitempty"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore string `json:"on_or_before,omitempty"`
}

//
type ACHPrenotificationList struct {
	// The contents of the list.
	Data []ACHPrenotification `json:"data"`
	// A pointer to a place in the list.
	NextCursor *string `json:"next_cursor"`
}

// A pointer to a place in the list.
func (r *ACHPrenotificationList) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

func (r *ACHPrenotificationService) Create(ctx context.Context, body *CreateAnACHPrenotificationParameters, opts ...*core.RequestOpts) (res *ACHPrenotification, err error) {
	err = r.post(
		ctx,
		"/ach_prenotifications",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)

	return
}

func (r *ACHPrenotificationService) Retrieve(ctx context.Context, ach_prenotification_id string, opts ...*core.RequestOpts) (res *ACHPrenotification, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/ach_prenotifications/%s", ach_prenotification_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)

	return
}

type ACHPrenotificationsPage struct {
	*pagination.Page[ACHPrenotification]
}

func (r *ACHPrenotificationsPage) ACHPrenotification() *ACHPrenotification {
	return r.Current()
}

func (r *ACHPrenotificationsPage) GetNextPage() (*ACHPrenotificationsPage, error) {
	if page, err := r.Page.GetNextPage(); err != nil {
		return nil, err
	} else {
		return &ACHPrenotificationsPage{page}, nil
	}
}

func (r *ACHPrenotificationService) List(ctx context.Context, query *ListACHPrenotificationsQuery, opts ...*core.RequestOpts) (res *ACHPrenotificationsPage, err error) {
	page := &ACHPrenotificationsPage{
		Page: &pagination.Page[ACHPrenotification]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/ach_prenotifications",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

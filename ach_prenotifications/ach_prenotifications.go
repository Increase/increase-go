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

type PreloadedACHPrenotificationService struct {
	ACHPrenotifications *ACHPrenotificationService
}

func (r *PreloadedACHPrenotificationService) Init(service *ACHPrenotificationService) {
	r.ACHPrenotifications = service
}

func NewPreloadedACHPrenotificationService(service *ACHPrenotificationService) (r *PreloadedACHPrenotificationService) {
	r = &PreloadedACHPrenotificationService{}
	r.Init(service)
	return
}

//
type ACHPrenotification struct {
	// The ACH Prenotification's identifier.
	ID *string `json:"id"`
	// The destination account number.
	AccountNumber *string `json:"account_number"`
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
	RoutingNumber *string `json:"routing_number"`
	// If your prenotification is returned, this will contain details of the return.
	PrenotificationReturn *ACHPrenotificationPrenotificationReturn `json:"prenotification_return"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the prenotification was created.
	CreatedAt *string `json:"created_at"`
	// The lifecycle status of the ACH Prenotification.
	Status *ACHPrenotificationStatus `json:"status"`
	// A constant representing the object's type. For this resource it will always be
	// `ach_prenotification`.
	Type *ACHPrenotificationType `json:"type"`
}

// The ACH Prenotification's identifier.
func (r *ACHPrenotification) GetID() (ID string) {
	if r != nil && r.ID != nil {
		ID = *r.ID
	}
	return
}

// The destination account number.
func (r *ACHPrenotification) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
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

// The American Bankers' Association (ABA) Routing Transit Number (RTN).
func (r *ACHPrenotification) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
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

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the prenotification was created.
func (r *ACHPrenotification) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The lifecycle status of the ACH Prenotification.
func (r *ACHPrenotification) GetStatus() (Status ACHPrenotificationStatus) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `ach_prenotification`.
func (r *ACHPrenotification) GetType() (Type ACHPrenotificationType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
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
	CreatedAt *string `json:"created_at"`
	// Why the Prenotification was returned.
	ReturnReasonCode *string `json:"return_reason_code"`
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the Prenotification was returned.
func (r *ACHPrenotificationPrenotificationReturn) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// Why the Prenotification was returned.
func (r *ACHPrenotificationPrenotificationReturn) GetReturnReasonCode() (ReturnReasonCode string) {
	if r != nil && r.ReturnReasonCode != nil {
		ReturnReasonCode = *r.ReturnReasonCode
	}
	return
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
	AccountNumber *string `json:"account_number"`
	// Additional information that will be sent to the recipient.
	Addendum *string `json:"addendum,omitempty"`
	// The description of the date of the transfer.
	CompanyDescriptiveDate *string `json:"company_descriptive_date,omitempty"`
	// The data you choose to associate with the transfer.
	CompanyDiscretionaryData *string `json:"company_discretionary_data,omitempty"`
	// The description of the transfer you wish to be shown to the recipient.
	CompanyEntryDescription *string `json:"company_entry_description,omitempty"`
	// The name by which the recipient knows you.
	CompanyName *string `json:"company_name,omitempty"`
	// Whether the Prenotification is for a future debit or credit.
	CreditDebitIndicator *CreateAnACHPrenotificationParametersCreditDebitIndicator `json:"credit_debit_indicator,omitempty"`
	// The transfer effective date in
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	EffectiveDate *string `json:"effective_date,omitempty"`
	// Your identifer for the transfer recipient.
	IndividualID *string `json:"individual_id,omitempty"`
	// The name of the transfer recipient. This value is information and not verified
	// by the recipient's bank.
	IndividualName *string `json:"individual_name,omitempty"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
	// destination account.
	RoutingNumber *string `json:"routing_number"`
	// The Standard Entry Class (SEC) code to use for the ACH Prenotification.
	StandardEntryClassCode *CreateAnACHPrenotificationParametersStandardEntryClassCode `json:"standard_entry_class_code,omitempty"`
}

// The account number for the destination account.
func (r *CreateAnACHPrenotificationParameters) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

// Additional information that will be sent to the recipient.
func (r *CreateAnACHPrenotificationParameters) GetAddendum() (Addendum string) {
	if r != nil && r.Addendum != nil {
		Addendum = *r.Addendum
	}
	return
}

// The description of the date of the transfer.
func (r *CreateAnACHPrenotificationParameters) GetCompanyDescriptiveDate() (CompanyDescriptiveDate string) {
	if r != nil && r.CompanyDescriptiveDate != nil {
		CompanyDescriptiveDate = *r.CompanyDescriptiveDate
	}
	return
}

// The data you choose to associate with the transfer.
func (r *CreateAnACHPrenotificationParameters) GetCompanyDiscretionaryData() (CompanyDiscretionaryData string) {
	if r != nil && r.CompanyDiscretionaryData != nil {
		CompanyDiscretionaryData = *r.CompanyDiscretionaryData
	}
	return
}

// The description of the transfer you wish to be shown to the recipient.
func (r *CreateAnACHPrenotificationParameters) GetCompanyEntryDescription() (CompanyEntryDescription string) {
	if r != nil && r.CompanyEntryDescription != nil {
		CompanyEntryDescription = *r.CompanyEntryDescription
	}
	return
}

// The name by which the recipient knows you.
func (r *CreateAnACHPrenotificationParameters) GetCompanyName() (CompanyName string) {
	if r != nil && r.CompanyName != nil {
		CompanyName = *r.CompanyName
	}
	return
}

// Whether the Prenotification is for a future debit or credit.
func (r *CreateAnACHPrenotificationParameters) GetCreditDebitIndicator() (CreditDebitIndicator CreateAnACHPrenotificationParametersCreditDebitIndicator) {
	if r != nil && r.CreditDebitIndicator != nil {
		CreditDebitIndicator = *r.CreditDebitIndicator
	}
	return
}

// The transfer effective date in
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
func (r *CreateAnACHPrenotificationParameters) GetEffectiveDate() (EffectiveDate string) {
	if r != nil && r.EffectiveDate != nil {
		EffectiveDate = *r.EffectiveDate
	}
	return
}

// Your identifer for the transfer recipient.
func (r *CreateAnACHPrenotificationParameters) GetIndividualID() (IndividualID string) {
	if r != nil && r.IndividualID != nil {
		IndividualID = *r.IndividualID
	}
	return
}

// The name of the transfer recipient. This value is information and not verified
// by the recipient's bank.
func (r *CreateAnACHPrenotificationParameters) GetIndividualName() (IndividualName string) {
	if r != nil && r.IndividualName != nil {
		IndividualName = *r.IndividualName
	}
	return
}

// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
// destination account.
func (r *CreateAnACHPrenotificationParameters) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

// The Standard Entry Class (SEC) code to use for the ACH Prenotification.
func (r *CreateAnACHPrenotificationParameters) GetStandardEntryClassCode() (StandardEntryClassCode CreateAnACHPrenotificationParametersStandardEntryClassCode) {
	if r != nil && r.StandardEntryClassCode != nil {
		StandardEntryClassCode = *r.StandardEntryClassCode
	}
	return
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
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit     *int                                   `query:"limit"`
	CreatedAt *ListACHPrenotificationsQueryCreatedAt `query:"created_at"`
}

// Return the page of entries after this one.
func (r *ListACHPrenotificationsQuery) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *ListACHPrenotificationsQuery) GetLimit() (Limit int) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

func (r *ListACHPrenotificationsQuery) GetCreatedAt() (CreatedAt ListACHPrenotificationsQueryCreatedAt) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

type ListACHPrenotificationsQueryCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After *string `json:"after,omitempty"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before *string `json:"before,omitempty"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter *string `json:"on_or_after,omitempty"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore *string `json:"on_or_before,omitempty"`
}

// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r *ListACHPrenotificationsQueryCreatedAt) GetAfter() (After string) {
	if r != nil && r.After != nil {
		After = *r.After
	}
	return
}

// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r *ListACHPrenotificationsQueryCreatedAt) GetBefore() (Before string) {
	if r != nil && r.Before != nil {
		Before = *r.Before
	}
	return
}

// Return results on or after this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r *ListACHPrenotificationsQueryCreatedAt) GetOnOrAfter() (OnOrAfter string) {
	if r != nil && r.OnOrAfter != nil {
		OnOrAfter = *r.OnOrAfter
	}
	return
}

// Return results on or before this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r *ListACHPrenotificationsQueryCreatedAt) GetOnOrBefore() (OnOrBefore string) {
	if r != nil && r.OnOrBefore != nil {
		OnOrBefore = *r.OnOrBefore
	}
	return
}

//
type ACHPrenotificationList struct {
	// The contents of the list.
	Data *[]ACHPrenotification `json:"data"`
	// A pointer to a place in the list.
	NextCursor *string `json:"next_cursor"`
}

// The contents of the list.
func (r *ACHPrenotificationList) GetData() (Data []ACHPrenotification) {
	if r != nil && r.Data != nil {
		Data = *r.Data
	}
	return
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

func (r *PreloadedACHPrenotificationService) Create(ctx context.Context, body *CreateAnACHPrenotificationParameters, opts ...*core.RequestOpts) (res *ACHPrenotification, err error) {
	err = r.ACHPrenotifications.post(
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

func (r *PreloadedACHPrenotificationService) Retrieve(ctx context.Context, ach_prenotification_id string, opts ...*core.RequestOpts) (res *ACHPrenotification, err error) {
	err = r.ACHPrenotifications.get(
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

func (r *PreloadedACHPrenotificationService) List(ctx context.Context, query *ListACHPrenotificationsQuery, opts ...*core.RequestOpts) (res *ACHPrenotificationsPage, err error) {
	page := &ACHPrenotificationsPage{
		Page: &pagination.Page[ACHPrenotification]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/ach_prenotifications",
			},
			Requester: r.ACHPrenotifications.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

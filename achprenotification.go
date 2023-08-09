// File generated from our OpenAPI spec by Stainless.

package increase

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/apiquery"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/internal/shared"
	"github.com/increase/increase-go/option"
)

// ACHPrenotificationService contains methods and other services that help with
// interacting with the increase API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewACHPrenotificationService] method
// instead.
type ACHPrenotificationService struct {
	Options []option.RequestOption
}

// NewACHPrenotificationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewACHPrenotificationService(opts ...option.RequestOption) (r *ACHPrenotificationService) {
	r = &ACHPrenotificationService{}
	r.Options = opts
	return
}

// Create an ACH Prenotification
func (r *ACHPrenotificationService) New(ctx context.Context, body ACHPrenotificationNewParams, opts ...option.RequestOption) (res *ACHPrenotification, err error) {
	opts = append(r.Options[:], opts...)
	path := "ach_prenotifications"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve an ACH Prenotification
func (r *ACHPrenotificationService) Get(ctx context.Context, achPrenotificationID string, opts ...option.RequestOption) (res *ACHPrenotification, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("ach_prenotifications/%s", achPrenotificationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List ACH Prenotifications
func (r *ACHPrenotificationService) List(ctx context.Context, query ACHPrenotificationListParams, opts ...option.RequestOption) (res *shared.Page[ACHPrenotification], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "ach_prenotifications"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List ACH Prenotifications
func (r *ACHPrenotificationService) ListAutoPaging(ctx context.Context, query ACHPrenotificationListParams, opts ...option.RequestOption) *shared.PageAutoPager[ACHPrenotification] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// ACH Prenotifications are one way you can verify account and routing numbers by
// Automated Clearing House (ACH).
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
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the prenotification was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// If the notification is for a future credit or debit.
	CreditDebitIndicator ACHPrenotificationCreditDebitIndicator `json:"credit_debit_indicator,required,nullable"`
	// The effective date in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	EffectiveDate time.Time `json:"effective_date,required,nullable" format:"date-time"`
	// If the receiving bank notifies that future transfers should use different
	// details, this will contain those details.
	NotificationsOfChange []ACHPrenotificationNotificationsOfChange `json:"notifications_of_change,required"`
	// If your prenotification is returned, this will contain details of the return.
	PrenotificationReturn ACHPrenotificationPrenotificationReturn `json:"prenotification_return,required,nullable"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber string `json:"routing_number,required"`
	// The lifecycle status of the ACH Prenotification.
	Status ACHPrenotificationStatus `json:"status,required"`
	// A constant representing the object's type. For this resource it will always be
	// `ach_prenotification`.
	Type ACHPrenotificationType `json:"type,required"`
	JSON achPrenotificationJSON
}

// achPrenotificationJSON contains the JSON metadata for the struct
// [ACHPrenotification]
type achPrenotificationJSON struct {
	ID                       apijson.Field
	AccountNumber            apijson.Field
	Addendum                 apijson.Field
	CompanyDescriptiveDate   apijson.Field
	CompanyDiscretionaryData apijson.Field
	CompanyEntryDescription  apijson.Field
	CompanyName              apijson.Field
	CreatedAt                apijson.Field
	CreditDebitIndicator     apijson.Field
	EffectiveDate            apijson.Field
	NotificationsOfChange    apijson.Field
	PrenotificationReturn    apijson.Field
	RoutingNumber            apijson.Field
	Status                   apijson.Field
	Type                     apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *ACHPrenotification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// If the notification is for a future credit or debit.
type ACHPrenotificationCreditDebitIndicator string

const (
	// The Prenotification is for an anticipated credit.
	ACHPrenotificationCreditDebitIndicatorCredit ACHPrenotificationCreditDebitIndicator = "credit"
	// The Prenotification is for an anticipated debit.
	ACHPrenotificationCreditDebitIndicatorDebit ACHPrenotificationCreditDebitIndicator = "debit"
)

type ACHPrenotificationNotificationsOfChange struct {
	// The type of change that occurred.
	ChangeCode ACHPrenotificationNotificationsOfChangeChangeCode `json:"change_code,required"`
	// The corrected data.
	CorrectedData string `json:"corrected_data,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the notification occurred.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	JSON      achPrenotificationNotificationsOfChangeJSON
}

// achPrenotificationNotificationsOfChangeJSON contains the JSON metadata for the
// struct [ACHPrenotificationNotificationsOfChange]
type achPrenotificationNotificationsOfChangeJSON struct {
	ChangeCode    apijson.Field
	CorrectedData apijson.Field
	CreatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ACHPrenotificationNotificationsOfChange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of change that occurred.
type ACHPrenotificationNotificationsOfChangeChangeCode string

const (
	// The account number was incorrect.
	ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectAccountNumber ACHPrenotificationNotificationsOfChangeChangeCode = "incorrect_account_number"
	// The routing number was incorrect.
	ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectRoutingNumber ACHPrenotificationNotificationsOfChangeChangeCode = "incorrect_routing_number"
	// Both the routing number and the account number were incorrect.
	ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectRoutingNumberAndAccountNumber ACHPrenotificationNotificationsOfChangeChangeCode = "incorrect_routing_number_and_account_number"
	// The transaction code was incorrect.
	ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectTransactionCode ACHPrenotificationNotificationsOfChangeChangeCode = "incorrect_transaction_code"
	// The account number and the transaction code were incorrect.
	ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectAccountNumberAndTransactionCode ACHPrenotificationNotificationsOfChangeChangeCode = "incorrect_account_number_and_transaction_code"
	// The routing number, account number, and transaction code were incorrect.
	ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectRoutingNumberAccountNumberAndTransactionCode ACHPrenotificationNotificationsOfChangeChangeCode = "incorrect_routing_number_account_number_and_transaction_code"
	// The receiving depository financial institution identification was incorrect.
	ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectReceivingDepositoryFinancialInstitutionIdentification ACHPrenotificationNotificationsOfChangeChangeCode = "incorrect_receiving_depository_financial_institution_identification"
	// The individual identification number was incorrect.
	ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectIndividualIdentificationNumber ACHPrenotificationNotificationsOfChangeChangeCode = "incorrect_individual_identification_number"
	// The addenda had an incorrect format.
	ACHPrenotificationNotificationsOfChangeChangeCodeAddendaFormatError ACHPrenotificationNotificationsOfChangeChangeCode = "addenda_format_error"
	// The standard entry class code was incorrect for an outbound international
	// payment.
	ACHPrenotificationNotificationsOfChangeChangeCodeIncorrectStandardEntryClassCodeForOutboundInternationalPayment ACHPrenotificationNotificationsOfChangeChangeCode = "incorrect_standard_entry_class_code_for_outbound_international_payment"
)

// If your prenotification is returned, this will contain details of the return.
type ACHPrenotificationPrenotificationReturn struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Prenotification was returned.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Why the Prenotification was returned.
	ReturnReasonCode string `json:"return_reason_code,required"`
	JSON             achPrenotificationPrenotificationReturnJSON
}

// achPrenotificationPrenotificationReturnJSON contains the JSON metadata for the
// struct [ACHPrenotificationPrenotificationReturn]
type achPrenotificationPrenotificationReturnJSON struct {
	CreatedAt        apijson.Field
	ReturnReasonCode apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ACHPrenotificationPrenotificationReturn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The lifecycle status of the ACH Prenotification.
type ACHPrenotificationStatus string

const (
	// The Prenotification is pending submission.
	ACHPrenotificationStatusPendingSubmitting ACHPrenotificationStatus = "pending_submitting"
	// The Prenotification requires attention.
	ACHPrenotificationStatusRequiresAttention ACHPrenotificationStatus = "requires_attention"
	// The Prenotification has been returned.
	ACHPrenotificationStatusReturned ACHPrenotificationStatus = "returned"
	// The Prentification is complete.
	ACHPrenotificationStatusSubmitted ACHPrenotificationStatus = "submitted"
)

// A constant representing the object's type. For this resource it will always be
// `ach_prenotification`.
type ACHPrenotificationType string

const (
	ACHPrenotificationTypeACHPrenotification ACHPrenotificationType = "ach_prenotification"
)

type ACHPrenotificationNewParams struct {
	// The account number for the destination account.
	AccountNumber param.Field[string] `json:"account_number,required"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
	// destination account.
	RoutingNumber param.Field[string] `json:"routing_number,required"`
	// Additional information that will be sent to the recipient.
	Addendum param.Field[string] `json:"addendum"`
	// The description of the date of the transfer.
	CompanyDescriptiveDate param.Field[string] `json:"company_descriptive_date"`
	// The data you choose to associate with the transfer.
	CompanyDiscretionaryData param.Field[string] `json:"company_discretionary_data"`
	// The description of the transfer you wish to be shown to the recipient.
	CompanyEntryDescription param.Field[string] `json:"company_entry_description"`
	// The name by which the recipient knows you.
	CompanyName param.Field[string] `json:"company_name"`
	// Whether the Prenotification is for a future debit or credit.
	CreditDebitIndicator param.Field[ACHPrenotificationNewParamsCreditDebitIndicator] `json:"credit_debit_indicator"`
	// The transfer effective date in
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	EffectiveDate param.Field[time.Time] `json:"effective_date" format:"date"`
	// Your identifer for the transfer recipient.
	IndividualID param.Field[string] `json:"individual_id"`
	// The name of the transfer recipient. This value is information and not verified
	// by the recipient's bank.
	IndividualName param.Field[string] `json:"individual_name"`
	// The Standard Entry Class (SEC) code to use for the ACH Prenotification.
	StandardEntryClassCode param.Field[ACHPrenotificationNewParamsStandardEntryClassCode] `json:"standard_entry_class_code"`
}

func (r ACHPrenotificationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether the Prenotification is for a future debit or credit.
type ACHPrenotificationNewParamsCreditDebitIndicator string

const (
	// The Prenotification is for an anticipated credit.
	ACHPrenotificationNewParamsCreditDebitIndicatorCredit ACHPrenotificationNewParamsCreditDebitIndicator = "credit"
	// The Prenotification is for an anticipated debit.
	ACHPrenotificationNewParamsCreditDebitIndicatorDebit ACHPrenotificationNewParamsCreditDebitIndicator = "debit"
)

// The Standard Entry Class (SEC) code to use for the ACH Prenotification.
type ACHPrenotificationNewParamsStandardEntryClassCode string

const (
	// Corporate Credit and Debit (CCD).
	ACHPrenotificationNewParamsStandardEntryClassCodeCorporateCreditOrDebit ACHPrenotificationNewParamsStandardEntryClassCode = "corporate_credit_or_debit"
	// Prearranged Payments and Deposits (PPD).
	ACHPrenotificationNewParamsStandardEntryClassCodePrearrangedPaymentsAndDeposit ACHPrenotificationNewParamsStandardEntryClassCode = "prearranged_payments_and_deposit"
	// Internet Initiated (WEB).
	ACHPrenotificationNewParamsStandardEntryClassCodeInternetInitiated ACHPrenotificationNewParamsStandardEntryClassCode = "internet_initiated"
)

type ACHPrenotificationListParams struct {
	CreatedAt param.Field[ACHPrenotificationListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [ACHPrenotificationListParams]'s query parameters as
// `url.Values`.
func (r ACHPrenotificationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ACHPrenotificationListParamsCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After param.Field[time.Time] `query:"after" format:"date-time"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before param.Field[time.Time] `query:"before" format:"date-time"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter param.Field[time.Time] `query:"on_or_after" format:"date-time"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore param.Field[time.Time] `query:"on_or_before" format:"date-time"`
}

// URLQuery serializes [ACHPrenotificationListParamsCreatedAt]'s query parameters
// as `url.Values`.
func (r ACHPrenotificationListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

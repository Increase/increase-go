package increase

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/apiquery"
	"github.com/increase/increase-go/internal/field"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/internal/shared"
	"github.com/increase/increase-go/option"
)

// ACHTransferService contains methods and other services that help with
// interacting with the increase API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewACHTransferService] method
// instead.
type ACHTransferService struct {
	Options []option.RequestOption
}

// NewACHTransferService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewACHTransferService(opts ...option.RequestOption) (r *ACHTransferService) {
	r = &ACHTransferService{}
	r.Options = opts
	return
}

// Create an ACH Transfer
func (r *ACHTransferService) New(ctx context.Context, body ACHTransferNewParams, opts ...option.RequestOption) (res *ACHTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := "ach_transfers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve an ACH Transfer
func (r *ACHTransferService) Get(ctx context.Context, ach_transfer_id string, opts ...option.RequestOption) (res *ACHTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("ach_transfers/%s", ach_transfer_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List ACH Transfers
func (r *ACHTransferService) List(ctx context.Context, query ACHTransferListParams, opts ...option.RequestOption) (res *shared.Page[ACHTransfer], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "ach_transfers"
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

// List ACH Transfers
func (r *ACHTransferService) ListAutoPaging(ctx context.Context, query ACHTransferListParams, opts ...option.RequestOption) *shared.PageAutoPager[ACHTransfer] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Approves an ACH Transfer in a pending_approval state.
func (r *ACHTransferService) Approve(ctx context.Context, ach_transfer_id string, opts ...option.RequestOption) (res *ACHTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("ach_transfers/%s/approve", ach_transfer_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Cancels an ACH Transfer in a pending_approval state.
func (r *ACHTransferService) Cancel(ctx context.Context, ach_transfer_id string, opts ...option.RequestOption) (res *ACHTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("ach_transfers/%s/cancel", ach_transfer_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// ACH transfers move funds between your Increase account and any other account
// accessible by the Automated Clearing House (ACH).
type ACHTransfer struct {
	// The Account to which the transfer belongs.
	AccountID string `json:"account_id,required"`
	// The destination account number.
	AccountNumber string `json:"account_number,required"`
	// Additional information that will be sent to the recipient.
	Addendum string `json:"addendum,required,nullable"`
	// The transfer amount in USD cents. A positive amount indicates a credit transfer
	// pushing funds to the receiving account. A negative amount indicates a debit
	// transfer pulling funds from the receiving account.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transfer's
	// currency. For ACH transfers this is always equal to `usd`.
	Currency ACHTransferCurrency `json:"currency,required"`
	// If your account requires approvals for transfers and the transfer was approved,
	// this will contain details of the approval.
	Approval ACHTransferApproval `json:"approval,required,nullable"`
	// If your account requires approvals for transfers and the transfer was not
	// approved, this will contain details of the cancellation.
	Cancellation ACHTransferCancellation `json:"cancellation,required,nullable"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The identifier of the External Account the transfer was made to, if any.
	ExternalAccountID string `json:"external_account_id,required,nullable"`
	// The ACH transfer's identifier.
	ID string `json:"id,required"`
	// The transfer's network.
	Network ACHTransferNetwork `json:"network,required"`
	// If the receiving bank accepts the transfer but notifies that future transfers
	// should use different details, this will contain those details.
	NotificationsOfChange []ACHTransferNotificationsOfChange `json:"notifications_of_change,required"`
	// If your transfer is returned, this will contain details of the return.
	Return ACHTransferReturn `json:"return,required,nullable"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber string `json:"routing_number,required"`
	// The descriptor that will show on the recipient's bank statement.
	StatementDescriptor string `json:"statement_descriptor,required"`
	// The lifecycle status of the transfer.
	Status ACHTransferStatus `json:"status,required"`
	// After the transfer is submitted to FedACH, this will contain supplemental
	// details.
	Submission ACHTransferSubmission `json:"submission,required,nullable"`
	// The ID for the transaction funding the transfer.
	TransactionID string `json:"transaction_id,required,nullable"`
	// The description of the date of the transfer.
	CompanyDescriptiveDate string `json:"company_descriptive_date,required,nullable"`
	// The data you chose to associate with the transfer.
	CompanyDiscretionaryData string `json:"company_discretionary_data,required,nullable"`
	// The description of the transfer you set to be shown to the recipient.
	CompanyEntryDescription string `json:"company_entry_description,required,nullable"`
	// The name by which the recipient knows you.
	CompanyName string `json:"company_name,required,nullable"`
	// The type of the account to which the transfer will be sent.
	Funding ACHTransferFunding `json:"funding,required"`
	// Your identifer for the transfer recipient.
	IndividualID string `json:"individual_id,required,nullable"`
	// The name of the transfer recipient. This value is information and not verified
	// by the recipient's bank.
	IndividualName string `json:"individual_name,required,nullable"`
	// The transfer effective date in
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	EffectiveDate time.Time `json:"effective_date,required,nullable" format:"date"`
	// The Standard Entry Class (SEC) code to use for the transfer.
	StandardEntryClassCode ACHTransferStandardEntryClassCode `json:"standard_entry_class_code,required"`
	// A constant representing the object's type. For this resource it will always be
	// `ach_transfer`.
	Type ACHTransferType `json:"type,required"`
	JSON achTransferJSON
}

// achTransferJSON contains the JSON metadata for the struct [ACHTransfer]
type achTransferJSON struct {
	AccountID                apijson.Field
	AccountNumber            apijson.Field
	Addendum                 apijson.Field
	Amount                   apijson.Field
	Currency                 apijson.Field
	Approval                 apijson.Field
	Cancellation             apijson.Field
	CreatedAt                apijson.Field
	ExternalAccountID        apijson.Field
	ID                       apijson.Field
	Network                  apijson.Field
	NotificationsOfChange    apijson.Field
	Return                   apijson.Field
	RoutingNumber            apijson.Field
	StatementDescriptor      apijson.Field
	Status                   apijson.Field
	Submission               apijson.Field
	TransactionID            apijson.Field
	CompanyDescriptiveDate   apijson.Field
	CompanyDiscretionaryData apijson.Field
	CompanyEntryDescription  apijson.Field
	CompanyName              apijson.Field
	Funding                  apijson.Field
	IndividualID             apijson.Field
	IndividualName           apijson.Field
	EffectiveDate            apijson.Field
	StandardEntryClassCode   apijson.Field
	Type                     apijson.Field
	raw                      string
	Extras                   map[string]apijson.Field
}

func (r *ACHTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ACHTransferCurrency string

const (
	ACHTransferCurrencyCad ACHTransferCurrency = "CAD"
	ACHTransferCurrencyChf ACHTransferCurrency = "CHF"
	ACHTransferCurrencyEur ACHTransferCurrency = "EUR"
	ACHTransferCurrencyGbp ACHTransferCurrency = "GBP"
	ACHTransferCurrencyJpy ACHTransferCurrency = "JPY"
	ACHTransferCurrencyUsd ACHTransferCurrency = "USD"
)

// If your account requires approvals for transfers and the transfer was approved,
// this will contain details of the approval.
type ACHTransferApproval struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was approved.
	ApprovedAt time.Time `json:"approved_at,required" format:"date-time"`
	// If the Transfer was approved by a user in the dashboard, the email address of
	// that user.
	ApprovedBy string `json:"approved_by,required,nullable"`
	JSON       achTransferApprovalJSON
}

// achTransferApprovalJSON contains the JSON metadata for the struct
// [ACHTransferApproval]
type achTransferApprovalJSON struct {
	ApprovedAt apijson.Field
	ApprovedBy apijson.Field
	raw        string
	Extras     map[string]apijson.Field
}

func (r *ACHTransferApproval) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// If your account requires approvals for transfers and the transfer was not
// approved, this will contain details of the cancellation.
type ACHTransferCancellation struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Transfer was canceled.
	CanceledAt time.Time `json:"canceled_at,required" format:"date-time"`
	// If the Transfer was canceled by a user in the dashboard, the email address of
	// that user.
	CanceledBy string `json:"canceled_by,required,nullable"`
	JSON       achTransferCancellationJSON
}

// achTransferCancellationJSON contains the JSON metadata for the struct
// [ACHTransferCancellation]
type achTransferCancellationJSON struct {
	CanceledAt apijson.Field
	CanceledBy apijson.Field
	raw        string
	Extras     map[string]apijson.Field
}

func (r *ACHTransferCancellation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ACHTransferNetwork string

const (
	ACHTransferNetworkACH ACHTransferNetwork = "ach"
)

type ACHTransferNotificationsOfChange struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the notification occurred.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The type of change that occurred.
	ChangeCode string `json:"change_code,required"`
	// The corrected data.
	CorrectedData string `json:"corrected_data,required"`
	JSON          achTransferNotificationsOfChangeJSON
}

// achTransferNotificationsOfChangeJSON contains the JSON metadata for the struct
// [ACHTransferNotificationsOfChange]
type achTransferNotificationsOfChangeJSON struct {
	CreatedAt     apijson.Field
	ChangeCode    apijson.Field
	CorrectedData apijson.Field
	raw           string
	Extras        map[string]apijson.Field
}

func (r *ACHTransferNotificationsOfChange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// If your transfer is returned, this will contain details of the return.
type ACHTransferReturn struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Why the ACH Transfer was returned.
	ReturnReasonCode ACHTransferReturnReturnReasonCode `json:"return_reason_code,required"`
	// The identifier of the ACH Transfer associated with this return.
	TransferID string `json:"transfer_id,required"`
	// The identifier of the Tranasaction associated with this return.
	TransactionID string `json:"transaction_id,required"`
	JSON          achTransferReturnJSON
}

// achTransferReturnJSON contains the JSON metadata for the struct
// [ACHTransferReturn]
type achTransferReturnJSON struct {
	CreatedAt        apijson.Field
	ReturnReasonCode apijson.Field
	TransferID       apijson.Field
	TransactionID    apijson.Field
	raw              string
	Extras           map[string]apijson.Field
}

func (r *ACHTransferReturn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ACHTransferReturnReturnReasonCode string

const (
	ACHTransferReturnReturnReasonCodeInsufficientFund                                          ACHTransferReturnReturnReasonCode = "insufficient_fund"
	ACHTransferReturnReturnReasonCodeNoAccount                                                 ACHTransferReturnReturnReasonCode = "no_account"
	ACHTransferReturnReturnReasonCodeAccountClosed                                             ACHTransferReturnReturnReasonCode = "account_closed"
	ACHTransferReturnReturnReasonCodeInvalidAccountNumberStructure                             ACHTransferReturnReturnReasonCode = "invalid_account_number_structure"
	ACHTransferReturnReturnReasonCodeAccountFrozenEntryReturnedPerOfacInstruction              ACHTransferReturnReturnReasonCode = "account_frozen_entry_returned_per_ofac_instruction"
	ACHTransferReturnReturnReasonCodeCreditEntryRefusedByReceiver                              ACHTransferReturnReturnReasonCode = "credit_entry_refused_by_receiver"
	ACHTransferReturnReturnReasonCodeUnauthorizedDebitToConsumerAccountUsingCorporateSecCode   ACHTransferReturnReturnReasonCode = "unauthorized_debit_to_consumer_account_using_corporate_sec_code"
	ACHTransferReturnReturnReasonCodeCorporateCustomerAdvisedNotAuthorized                     ACHTransferReturnReturnReasonCode = "corporate_customer_advised_not_authorized"
	ACHTransferReturnReturnReasonCodePaymentStopped                                            ACHTransferReturnReturnReasonCode = "payment_stopped"
	ACHTransferReturnReturnReasonCodeNonTransactionAccount                                     ACHTransferReturnReturnReasonCode = "non_transaction_account"
	ACHTransferReturnReturnReasonCodeUncollectedFunds                                          ACHTransferReturnReturnReasonCode = "uncollected_funds"
	ACHTransferReturnReturnReasonCodeRoutingNumberCheckDigitError                              ACHTransferReturnReturnReasonCode = "routing_number_check_digit_error"
	ACHTransferReturnReturnReasonCodeCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete ACHTransferReturnReturnReasonCode = "customer_advised_unauthorized_improper_ineligible_or_incomplete"
	ACHTransferReturnReturnReasonCodeAmountFieldError                                          ACHTransferReturnReturnReasonCode = "amount_field_error"
	ACHTransferReturnReturnReasonCodeAuthorizationRevokedByCustomer                            ACHTransferReturnReturnReasonCode = "authorization_revoked_by_customer"
	ACHTransferReturnReturnReasonCodeInvalidACHRoutingNumber                                   ACHTransferReturnReturnReasonCode = "invalid_ach_routing_number"
	ACHTransferReturnReturnReasonCodeFileRecordEditCriteria                                    ACHTransferReturnReturnReasonCode = "file_record_edit_criteria"
	ACHTransferReturnReturnReasonCodeEnrInvalidIndividualName                                  ACHTransferReturnReturnReasonCode = "enr_invalid_individual_name"
	ACHTransferReturnReturnReasonCodeReturnedPerOdfiRequest                                    ACHTransferReturnReturnReasonCode = "returned_per_odfi_request"
	ACHTransferReturnReturnReasonCodeAddendaError                                              ACHTransferReturnReturnReasonCode = "addenda_error"
	ACHTransferReturnReturnReasonCodeLimitedParticipationDfi                                   ACHTransferReturnReturnReasonCode = "limited_participation_dfi"
	ACHTransferReturnReturnReasonCodeIncorrectlyCodedOutboundInternationalPayment              ACHTransferReturnReturnReasonCode = "incorrectly_coded_outbound_international_payment"
	ACHTransferReturnReturnReasonCodeOther                                                     ACHTransferReturnReturnReasonCode = "other"
)

type ACHTransferStatus string

const (
	ACHTransferStatusPendingApproval   ACHTransferStatus = "pending_approval"
	ACHTransferStatusCanceled          ACHTransferStatus = "canceled"
	ACHTransferStatusPendingReviewing  ACHTransferStatus = "pending_reviewing"
	ACHTransferStatusPendingSubmission ACHTransferStatus = "pending_submission"
	ACHTransferStatusSubmitted         ACHTransferStatus = "submitted"
	ACHTransferStatusReturned          ACHTransferStatus = "returned"
	ACHTransferStatusRequiresAttention ACHTransferStatus = "requires_attention"
	ACHTransferStatusRejected          ACHTransferStatus = "rejected"
)

// After the transfer is submitted to FedACH, this will contain supplemental
// details.
type ACHTransferSubmission struct {
	// The trace number for the submission.
	TraceNumber string `json:"trace_number,required"`
	// When the ACH transfer was sent to FedACH.
	SubmittedAt time.Time `json:"submitted_at,required" format:"date-time"`
	JSON        achTransferSubmissionJSON
}

// achTransferSubmissionJSON contains the JSON metadata for the struct
// [ACHTransferSubmission]
type achTransferSubmissionJSON struct {
	TraceNumber apijson.Field
	SubmittedAt apijson.Field
	raw         string
	Extras      map[string]apijson.Field
}

func (r *ACHTransferSubmission) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ACHTransferFunding string

const (
	ACHTransferFundingChecking ACHTransferFunding = "checking"
	ACHTransferFundingSavings  ACHTransferFunding = "savings"
)

type ACHTransferStandardEntryClassCode string

const (
	ACHTransferStandardEntryClassCodeCorporateCreditOrDebit        ACHTransferStandardEntryClassCode = "corporate_credit_or_debit"
	ACHTransferStandardEntryClassCodePrearrangedPaymentsAndDeposit ACHTransferStandardEntryClassCode = "prearranged_payments_and_deposit"
	ACHTransferStandardEntryClassCodeInternetInitiated             ACHTransferStandardEntryClassCode = "internet_initiated"
)

type ACHTransferType string

const (
	ACHTransferTypeACHTransfer ACHTransferType = "ach_transfer"
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

// URLQuery serializes [ACHTransferListParams]'s query parameters as `url.Values`.
func (r ACHTransferListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
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

// URLQuery serializes [ACHTransferListParamsCreatedAt]'s query parameters as
// `url.Values`.
func (r ACHTransferListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

// A list of ACH Transfer objects
type ACHTransferListResponse struct {
	// The contents of the list.
	Data []ACHTransfer `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       achTransferListResponseJSON
}

// achTransferListResponseJSON contains the JSON metadata for the struct
// [ACHTransferListResponse]
type achTransferListResponseJSON struct {
	Data       apijson.Field
	NextCursor apijson.Field
	raw        string
	Extras     map[string]apijson.Field
}

func (r *ACHTransferListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

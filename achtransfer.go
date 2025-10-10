// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Increase/increase-go/internal/apijson"
	"github.com/Increase/increase-go/internal/apiquery"
	"github.com/Increase/increase-go/internal/param"
	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
	"github.com/Increase/increase-go/packages/pagination"
)

// ACHTransferService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewACHTransferService] method instead.
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
	opts = slices.Concat(r.Options, opts)
	path := "ach_transfers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve an ACH Transfer
func (r *ACHTransferService) Get(ctx context.Context, achTransferID string, opts ...option.RequestOption) (res *ACHTransfer, err error) {
	opts = slices.Concat(r.Options, opts)
	if achTransferID == "" {
		err = errors.New("missing required ach_transfer_id parameter")
		return
	}
	path := fmt.Sprintf("ach_transfers/%s", achTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List ACH Transfers
func (r *ACHTransferService) List(ctx context.Context, query ACHTransferListParams, opts ...option.RequestOption) (res *pagination.Page[ACHTransfer], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
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
func (r *ACHTransferService) ListAutoPaging(ctx context.Context, query ACHTransferListParams, opts ...option.RequestOption) *pagination.PageAutoPager[ACHTransfer] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Approves an ACH Transfer in a pending_approval state.
func (r *ACHTransferService) Approve(ctx context.Context, achTransferID string, opts ...option.RequestOption) (res *ACHTransfer, err error) {
	opts = slices.Concat(r.Options, opts)
	if achTransferID == "" {
		err = errors.New("missing required ach_transfer_id parameter")
		return
	}
	path := fmt.Sprintf("ach_transfers/%s/approve", achTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Cancels an ACH Transfer in a pending_approval state.
func (r *ACHTransferService) Cancel(ctx context.Context, achTransferID string, opts ...option.RequestOption) (res *ACHTransfer, err error) {
	opts = slices.Concat(r.Options, opts)
	if achTransferID == "" {
		err = errors.New("missing required ach_transfer_id parameter")
		return
	}
	path := fmt.Sprintf("ach_transfers/%s/cancel", achTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// ACH transfers move funds between your Increase account and any other account
// accessible by the Automated Clearing House (ACH).
type ACHTransfer struct {
	// The ACH transfer's identifier.
	ID string `json:"id,required"`
	// The Account to which the transfer belongs.
	AccountID string `json:"account_id,required"`
	// The destination account number.
	AccountNumber string `json:"account_number,required"`
	// After the transfer is acknowledged by FedACH, this will contain supplemental
	// details. The Federal Reserve sends an acknowledgement message for each file that
	// Increase submits.
	Acknowledgement ACHTransferAcknowledgement `json:"acknowledgement,required,nullable"`
	// Additional information that will be sent to the recipient.
	Addenda ACHTransferAddenda `json:"addenda,required,nullable"`
	// The transfer amount in USD cents. A positive amount indicates a credit transfer
	// pushing funds to the receiving account. A negative amount indicates a debit
	// transfer pulling funds from the receiving account.
	Amount int64 `json:"amount,required"`
	// If your account requires approvals for transfers and the transfer was approved,
	// this will contain details of the approval.
	Approval ACHTransferApproval `json:"approval,required,nullable"`
	// If your account requires approvals for transfers and the transfer was not
	// approved, this will contain details of the cancellation.
	Cancellation ACHTransferCancellation `json:"cancellation,required,nullable"`
	// The description of the date of the transfer.
	CompanyDescriptiveDate string `json:"company_descriptive_date,required,nullable"`
	// The data you chose to associate with the transfer.
	CompanyDiscretionaryData string `json:"company_discretionary_data,required,nullable"`
	// The description of the transfer you set to be shown to the recipient.
	CompanyEntryDescription string `json:"company_entry_description,required,nullable"`
	// The company ID associated with the transfer.
	CompanyID string `json:"company_id,required"`
	// The name by which the recipient knows you.
	CompanyName string `json:"company_name,required,nullable"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// What object created the transfer, either via the API or the dashboard.
	CreatedBy ACHTransferCreatedBy `json:"created_by,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transfer's
	// currency. For ACH transfers this is always equal to `usd`.
	Currency ACHTransferCurrency `json:"currency,required"`
	// The type of entity that owns the account to which the ACH Transfer is being
	// sent.
	DestinationAccountHolder ACHTransferDestinationAccountHolder `json:"destination_account_holder,required"`
	// The identifier of the External Account the transfer was made to, if any.
	ExternalAccountID string `json:"external_account_id,required,nullable"`
	// The type of the account to which the transfer will be sent.
	Funding ACHTransferFunding `json:"funding,required"`
	// The idempotency key you chose for this object. This value is unique across
	// Increase and is used to ensure that a request is only processed once. Learn more
	// about [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey string `json:"idempotency_key,required,nullable"`
	// Increase will sometimes hold the funds for ACH debit transfers. If funds are
	// held, this sub-object will contain details of the hold.
	InboundFundsHold ACHTransferInboundFundsHold `json:"inbound_funds_hold,required,nullable"`
	// Your identifier for the transfer recipient.
	IndividualID string `json:"individual_id,required,nullable"`
	// The name of the transfer recipient. This value is information and not verified
	// by the recipient's bank.
	IndividualName string `json:"individual_name,required,nullable"`
	// The transfer's network.
	Network ACHTransferNetwork `json:"network,required"`
	// If the receiving bank accepts the transfer but notifies that future transfers
	// should use different details, this will contain those details.
	NotificationsOfChange []ACHTransferNotificationsOfChange `json:"notifications_of_change,required"`
	// The ID for the pending transaction representing the transfer. A pending
	// transaction is created when the transfer
	// [requires approval](https://increase.com/documentation/transfer-approvals#transfer-approvals)
	// by someone else in your organization.
	PendingTransactionID string `json:"pending_transaction_id,required,nullable"`
	// Configuration for how the effective date of the transfer will be set. This
	// determines same-day vs future-dated settlement timing. If not set, defaults to a
	// `settlement_schedule` of `same_day`. If set, exactly one of the child attributes
	// must be set.
	PreferredEffectiveDate ACHTransferPreferredEffectiveDate `json:"preferred_effective_date,required"`
	// If your transfer is returned, this will contain details of the return.
	Return ACHTransferReturn `json:"return,required,nullable"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber string `json:"routing_number,required"`
	// A subhash containing information about when and how the transfer settled at the
	// Federal Reserve.
	Settlement ACHTransferSettlement `json:"settlement,required,nullable"`
	// The Standard Entry Class (SEC) code to use for the transfer.
	StandardEntryClassCode ACHTransferStandardEntryClassCode `json:"standard_entry_class_code,required"`
	// The descriptor that will show on the recipient's bank statement.
	StatementDescriptor string `json:"statement_descriptor,required"`
	// The lifecycle status of the transfer.
	Status ACHTransferStatus `json:"status,required"`
	// After the transfer is submitted to FedACH, this will contain supplemental
	// details. Increase batches transfers and submits a file to the Federal Reserve
	// roughly every 30 minutes. The Federal Reserve processes ACH transfers during
	// weekdays according to their
	// [posted schedule](https://www.frbservices.org/resources/resource-centers/same-day-ach/fedach-processing-schedule.html).
	Submission ACHTransferSubmission `json:"submission,required,nullable"`
	// The ID for the transaction funding the transfer.
	TransactionID string `json:"transaction_id,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `ach_transfer`.
	Type ACHTransferType `json:"type,required"`
	JSON achTransferJSON `json:"-"`
}

// achTransferJSON contains the JSON metadata for the struct [ACHTransfer]
type achTransferJSON struct {
	ID                       apijson.Field
	AccountID                apijson.Field
	AccountNumber            apijson.Field
	Acknowledgement          apijson.Field
	Addenda                  apijson.Field
	Amount                   apijson.Field
	Approval                 apijson.Field
	Cancellation             apijson.Field
	CompanyDescriptiveDate   apijson.Field
	CompanyDiscretionaryData apijson.Field
	CompanyEntryDescription  apijson.Field
	CompanyID                apijson.Field
	CompanyName              apijson.Field
	CreatedAt                apijson.Field
	CreatedBy                apijson.Field
	Currency                 apijson.Field
	DestinationAccountHolder apijson.Field
	ExternalAccountID        apijson.Field
	Funding                  apijson.Field
	IdempotencyKey           apijson.Field
	InboundFundsHold         apijson.Field
	IndividualID             apijson.Field
	IndividualName           apijson.Field
	Network                  apijson.Field
	NotificationsOfChange    apijson.Field
	PendingTransactionID     apijson.Field
	PreferredEffectiveDate   apijson.Field
	Return                   apijson.Field
	RoutingNumber            apijson.Field
	Settlement               apijson.Field
	StandardEntryClassCode   apijson.Field
	StatementDescriptor      apijson.Field
	Status                   apijson.Field
	Submission               apijson.Field
	TransactionID            apijson.Field
	Type                     apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *ACHTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r achTransferJSON) RawJSON() string {
	return r.raw
}

// After the transfer is acknowledged by FedACH, this will contain supplemental
// details. The Federal Reserve sends an acknowledgement message for each file that
// Increase submits.
type ACHTransferAcknowledgement struct {
	// When the Federal Reserve acknowledged the submitted file containing this
	// transfer.
	AcknowledgedAt string                         `json:"acknowledged_at,required"`
	JSON           achTransferAcknowledgementJSON `json:"-"`
}

// achTransferAcknowledgementJSON contains the JSON metadata for the struct
// [ACHTransferAcknowledgement]
type achTransferAcknowledgementJSON struct {
	AcknowledgedAt apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ACHTransferAcknowledgement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r achTransferAcknowledgementJSON) RawJSON() string {
	return r.raw
}

// Additional information that will be sent to the recipient.
type ACHTransferAddenda struct {
	// The type of the resource. We may add additional possible values for this enum
	// over time; your application should be able to handle such additions gracefully.
	Category ACHTransferAddendaCategory `json:"category,required"`
	// Unstructured `payment_related_information` passed through with the transfer.
	Freeform ACHTransferAddendaFreeform `json:"freeform,required,nullable"`
	// Structured ASC X12 820 remittance advice records. Please reach out to
	// [support@increase.com](mailto:support@increase.com) for more information.
	PaymentOrderRemittanceAdvice ACHTransferAddendaPaymentOrderRemittanceAdvice `json:"payment_order_remittance_advice,required,nullable"`
	JSON                         achTransferAddendaJSON                         `json:"-"`
}

// achTransferAddendaJSON contains the JSON metadata for the struct
// [ACHTransferAddenda]
type achTransferAddendaJSON struct {
	Category                     apijson.Field
	Freeform                     apijson.Field
	PaymentOrderRemittanceAdvice apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ACHTransferAddenda) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r achTransferAddendaJSON) RawJSON() string {
	return r.raw
}

// The type of the resource. We may add additional possible values for this enum
// over time; your application should be able to handle such additions gracefully.
type ACHTransferAddendaCategory string

const (
	ACHTransferAddendaCategoryFreeform                     ACHTransferAddendaCategory = "freeform"
	ACHTransferAddendaCategoryPaymentOrderRemittanceAdvice ACHTransferAddendaCategory = "payment_order_remittance_advice"
	ACHTransferAddendaCategoryOther                        ACHTransferAddendaCategory = "other"
)

func (r ACHTransferAddendaCategory) IsKnown() bool {
	switch r {
	case ACHTransferAddendaCategoryFreeform, ACHTransferAddendaCategoryPaymentOrderRemittanceAdvice, ACHTransferAddendaCategoryOther:
		return true
	}
	return false
}

// Unstructured `payment_related_information` passed through with the transfer.
type ACHTransferAddendaFreeform struct {
	// Each entry represents an addendum sent with the transfer.
	Entries []ACHTransferAddendaFreeformEntry `json:"entries,required"`
	JSON    achTransferAddendaFreeformJSON    `json:"-"`
}

// achTransferAddendaFreeformJSON contains the JSON metadata for the struct
// [ACHTransferAddendaFreeform]
type achTransferAddendaFreeformJSON struct {
	Entries     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACHTransferAddendaFreeform) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r achTransferAddendaFreeformJSON) RawJSON() string {
	return r.raw
}

type ACHTransferAddendaFreeformEntry struct {
	// The payment related information passed in the addendum.
	PaymentRelatedInformation string                              `json:"payment_related_information,required"`
	JSON                      achTransferAddendaFreeformEntryJSON `json:"-"`
}

// achTransferAddendaFreeformEntryJSON contains the JSON metadata for the struct
// [ACHTransferAddendaFreeformEntry]
type achTransferAddendaFreeformEntryJSON struct {
	PaymentRelatedInformation apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *ACHTransferAddendaFreeformEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r achTransferAddendaFreeformEntryJSON) RawJSON() string {
	return r.raw
}

// Structured ASC X12 820 remittance advice records. Please reach out to
// [support@increase.com](mailto:support@increase.com) for more information.
type ACHTransferAddendaPaymentOrderRemittanceAdvice struct {
	// ASC X12 RMR records for this specific transfer.
	Invoices []ACHTransferAddendaPaymentOrderRemittanceAdviceInvoice `json:"invoices,required"`
	JSON     achTransferAddendaPaymentOrderRemittanceAdviceJSON      `json:"-"`
}

// achTransferAddendaPaymentOrderRemittanceAdviceJSON contains the JSON metadata
// for the struct [ACHTransferAddendaPaymentOrderRemittanceAdvice]
type achTransferAddendaPaymentOrderRemittanceAdviceJSON struct {
	Invoices    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACHTransferAddendaPaymentOrderRemittanceAdvice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r achTransferAddendaPaymentOrderRemittanceAdviceJSON) RawJSON() string {
	return r.raw
}

type ACHTransferAddendaPaymentOrderRemittanceAdviceInvoice struct {
	// The invoice number for this reference, determined in advance with the receiver.
	InvoiceNumber string `json:"invoice_number,required"`
	// The amount that was paid for this invoice in the minor unit of its currency. For
	// dollars, for example, this is cents.
	PaidAmount int64                                                     `json:"paid_amount,required"`
	JSON       achTransferAddendaPaymentOrderRemittanceAdviceInvoiceJSON `json:"-"`
}

// achTransferAddendaPaymentOrderRemittanceAdviceInvoiceJSON contains the JSON
// metadata for the struct [ACHTransferAddendaPaymentOrderRemittanceAdviceInvoice]
type achTransferAddendaPaymentOrderRemittanceAdviceInvoiceJSON struct {
	InvoiceNumber apijson.Field
	PaidAmount    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ACHTransferAddendaPaymentOrderRemittanceAdviceInvoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r achTransferAddendaPaymentOrderRemittanceAdviceInvoiceJSON) RawJSON() string {
	return r.raw
}

// If your account requires approvals for transfers and the transfer was approved,
// this will contain details of the approval.
type ACHTransferApproval struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was approved.
	ApprovedAt time.Time `json:"approved_at,required" format:"date-time"`
	// If the Transfer was approved by a user in the dashboard, the email address of
	// that user.
	ApprovedBy string                  `json:"approved_by,required,nullable"`
	JSON       achTransferApprovalJSON `json:"-"`
}

// achTransferApprovalJSON contains the JSON metadata for the struct
// [ACHTransferApproval]
type achTransferApprovalJSON struct {
	ApprovedAt  apijson.Field
	ApprovedBy  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACHTransferApproval) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r achTransferApprovalJSON) RawJSON() string {
	return r.raw
}

// If your account requires approvals for transfers and the transfer was not
// approved, this will contain details of the cancellation.
type ACHTransferCancellation struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Transfer was canceled.
	CanceledAt time.Time `json:"canceled_at,required" format:"date-time"`
	// If the Transfer was canceled by a user in the dashboard, the email address of
	// that user.
	CanceledBy string                      `json:"canceled_by,required,nullable"`
	JSON       achTransferCancellationJSON `json:"-"`
}

// achTransferCancellationJSON contains the JSON metadata for the struct
// [ACHTransferCancellation]
type achTransferCancellationJSON struct {
	CanceledAt  apijson.Field
	CanceledBy  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACHTransferCancellation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r achTransferCancellationJSON) RawJSON() string {
	return r.raw
}

// What object created the transfer, either via the API or the dashboard.
type ACHTransferCreatedBy struct {
	// If present, details about the API key that created the transfer.
	APIKey ACHTransferCreatedByAPIKey `json:"api_key,required,nullable"`
	// The type of object that created this transfer.
	Category ACHTransferCreatedByCategory `json:"category,required"`
	// If present, details about the OAuth Application that created the transfer.
	OAuthApplication ACHTransferCreatedByOAuthApplication `json:"oauth_application,required,nullable"`
	// If present, details about the User that created the transfer.
	User ACHTransferCreatedByUser `json:"user,required,nullable"`
	JSON achTransferCreatedByJSON `json:"-"`
}

// achTransferCreatedByJSON contains the JSON metadata for the struct
// [ACHTransferCreatedBy]
type achTransferCreatedByJSON struct {
	APIKey           apijson.Field
	Category         apijson.Field
	OAuthApplication apijson.Field
	User             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ACHTransferCreatedBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r achTransferCreatedByJSON) RawJSON() string {
	return r.raw
}

// If present, details about the API key that created the transfer.
type ACHTransferCreatedByAPIKey struct {
	// The description set for the API key when it was created.
	Description string                         `json:"description,required,nullable"`
	JSON        achTransferCreatedByAPIKeyJSON `json:"-"`
}

// achTransferCreatedByAPIKeyJSON contains the JSON metadata for the struct
// [ACHTransferCreatedByAPIKey]
type achTransferCreatedByAPIKeyJSON struct {
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACHTransferCreatedByAPIKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r achTransferCreatedByAPIKeyJSON) RawJSON() string {
	return r.raw
}

// The type of object that created this transfer.
type ACHTransferCreatedByCategory string

const (
	ACHTransferCreatedByCategoryAPIKey           ACHTransferCreatedByCategory = "api_key"
	ACHTransferCreatedByCategoryOAuthApplication ACHTransferCreatedByCategory = "oauth_application"
	ACHTransferCreatedByCategoryUser             ACHTransferCreatedByCategory = "user"
)

func (r ACHTransferCreatedByCategory) IsKnown() bool {
	switch r {
	case ACHTransferCreatedByCategoryAPIKey, ACHTransferCreatedByCategoryOAuthApplication, ACHTransferCreatedByCategoryUser:
		return true
	}
	return false
}

// If present, details about the OAuth Application that created the transfer.
type ACHTransferCreatedByOAuthApplication struct {
	// The name of the OAuth Application.
	Name string                                   `json:"name,required"`
	JSON achTransferCreatedByOAuthApplicationJSON `json:"-"`
}

// achTransferCreatedByOAuthApplicationJSON contains the JSON metadata for the
// struct [ACHTransferCreatedByOAuthApplication]
type achTransferCreatedByOAuthApplicationJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACHTransferCreatedByOAuthApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r achTransferCreatedByOAuthApplicationJSON) RawJSON() string {
	return r.raw
}

// If present, details about the User that created the transfer.
type ACHTransferCreatedByUser struct {
	// The email address of the User.
	Email string                       `json:"email,required"`
	JSON  achTransferCreatedByUserJSON `json:"-"`
}

// achTransferCreatedByUserJSON contains the JSON metadata for the struct
// [ACHTransferCreatedByUser]
type achTransferCreatedByUserJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACHTransferCreatedByUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r achTransferCreatedByUserJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transfer's
// currency. For ACH transfers this is always equal to `usd`.
type ACHTransferCurrency string

const (
	ACHTransferCurrencyCad ACHTransferCurrency = "CAD"
	ACHTransferCurrencyChf ACHTransferCurrency = "CHF"
	ACHTransferCurrencyEur ACHTransferCurrency = "EUR"
	ACHTransferCurrencyGbp ACHTransferCurrency = "GBP"
	ACHTransferCurrencyJpy ACHTransferCurrency = "JPY"
	ACHTransferCurrencyUsd ACHTransferCurrency = "USD"
)

func (r ACHTransferCurrency) IsKnown() bool {
	switch r {
	case ACHTransferCurrencyCad, ACHTransferCurrencyChf, ACHTransferCurrencyEur, ACHTransferCurrencyGbp, ACHTransferCurrencyJpy, ACHTransferCurrencyUsd:
		return true
	}
	return false
}

// The type of entity that owns the account to which the ACH Transfer is being
// sent.
type ACHTransferDestinationAccountHolder string

const (
	ACHTransferDestinationAccountHolderBusiness   ACHTransferDestinationAccountHolder = "business"
	ACHTransferDestinationAccountHolderIndividual ACHTransferDestinationAccountHolder = "individual"
	ACHTransferDestinationAccountHolderUnknown    ACHTransferDestinationAccountHolder = "unknown"
)

func (r ACHTransferDestinationAccountHolder) IsKnown() bool {
	switch r {
	case ACHTransferDestinationAccountHolderBusiness, ACHTransferDestinationAccountHolderIndividual, ACHTransferDestinationAccountHolderUnknown:
		return true
	}
	return false
}

// The type of the account to which the transfer will be sent.
type ACHTransferFunding string

const (
	ACHTransferFundingChecking      ACHTransferFunding = "checking"
	ACHTransferFundingSavings       ACHTransferFunding = "savings"
	ACHTransferFundingGeneralLedger ACHTransferFunding = "general_ledger"
)

func (r ACHTransferFunding) IsKnown() bool {
	switch r {
	case ACHTransferFundingChecking, ACHTransferFundingSavings, ACHTransferFundingGeneralLedger:
		return true
	}
	return false
}

// Increase will sometimes hold the funds for ACH debit transfers. If funds are
// held, this sub-object will contain details of the hold.
type ACHTransferInboundFundsHold struct {
	// The held amount in the minor unit of the account's currency. For dollars, for
	// example, this is cents.
	Amount int64 `json:"amount,required"`
	// When the hold will be released automatically. Certain conditions may cause it to
	// be released before this time.
	AutomaticallyReleasesAt time.Time `json:"automatically_releases_at,required" format:"date-time"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the hold
	// was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the hold's
	// currency.
	Currency ACHTransferInboundFundsHoldCurrency `json:"currency,required"`
	// The ID of the Transaction for which funds were held.
	HeldTransactionID string `json:"held_transaction_id,required,nullable"`
	// The ID of the Pending Transaction representing the held funds.
	PendingTransactionID string `json:"pending_transaction_id,required,nullable"`
	// When the hold was released (if it has been released).
	ReleasedAt time.Time `json:"released_at,required,nullable" format:"date-time"`
	// The status of the hold.
	Status ACHTransferInboundFundsHoldStatus `json:"status,required"`
	// A constant representing the object's type. For this resource it will always be
	// `inbound_funds_hold`.
	Type ACHTransferInboundFundsHoldType `json:"type,required"`
	JSON achTransferInboundFundsHoldJSON `json:"-"`
}

// achTransferInboundFundsHoldJSON contains the JSON metadata for the struct
// [ACHTransferInboundFundsHold]
type achTransferInboundFundsHoldJSON struct {
	Amount                  apijson.Field
	AutomaticallyReleasesAt apijson.Field
	CreatedAt               apijson.Field
	Currency                apijson.Field
	HeldTransactionID       apijson.Field
	PendingTransactionID    apijson.Field
	ReleasedAt              apijson.Field
	Status                  apijson.Field
	Type                    apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ACHTransferInboundFundsHold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r achTransferInboundFundsHoldJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the hold's
// currency.
type ACHTransferInboundFundsHoldCurrency string

const (
	ACHTransferInboundFundsHoldCurrencyCad ACHTransferInboundFundsHoldCurrency = "CAD"
	ACHTransferInboundFundsHoldCurrencyChf ACHTransferInboundFundsHoldCurrency = "CHF"
	ACHTransferInboundFundsHoldCurrencyEur ACHTransferInboundFundsHoldCurrency = "EUR"
	ACHTransferInboundFundsHoldCurrencyGbp ACHTransferInboundFundsHoldCurrency = "GBP"
	ACHTransferInboundFundsHoldCurrencyJpy ACHTransferInboundFundsHoldCurrency = "JPY"
	ACHTransferInboundFundsHoldCurrencyUsd ACHTransferInboundFundsHoldCurrency = "USD"
)

func (r ACHTransferInboundFundsHoldCurrency) IsKnown() bool {
	switch r {
	case ACHTransferInboundFundsHoldCurrencyCad, ACHTransferInboundFundsHoldCurrencyChf, ACHTransferInboundFundsHoldCurrencyEur, ACHTransferInboundFundsHoldCurrencyGbp, ACHTransferInboundFundsHoldCurrencyJpy, ACHTransferInboundFundsHoldCurrencyUsd:
		return true
	}
	return false
}

// The status of the hold.
type ACHTransferInboundFundsHoldStatus string

const (
	ACHTransferInboundFundsHoldStatusHeld     ACHTransferInboundFundsHoldStatus = "held"
	ACHTransferInboundFundsHoldStatusComplete ACHTransferInboundFundsHoldStatus = "complete"
)

func (r ACHTransferInboundFundsHoldStatus) IsKnown() bool {
	switch r {
	case ACHTransferInboundFundsHoldStatusHeld, ACHTransferInboundFundsHoldStatusComplete:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `inbound_funds_hold`.
type ACHTransferInboundFundsHoldType string

const (
	ACHTransferInboundFundsHoldTypeInboundFundsHold ACHTransferInboundFundsHoldType = "inbound_funds_hold"
)

func (r ACHTransferInboundFundsHoldType) IsKnown() bool {
	switch r {
	case ACHTransferInboundFundsHoldTypeInboundFundsHold:
		return true
	}
	return false
}

// The transfer's network.
type ACHTransferNetwork string

const (
	ACHTransferNetworkACH ACHTransferNetwork = "ach"
)

func (r ACHTransferNetwork) IsKnown() bool {
	switch r {
	case ACHTransferNetworkACH:
		return true
	}
	return false
}

type ACHTransferNotificationsOfChange struct {
	// The required type of change that is being signaled by the receiving financial
	// institution.
	ChangeCode ACHTransferNotificationsOfChangeChangeCode `json:"change_code,required"`
	// The corrected data that should be used in future ACHs to this account. This may
	// contain the suggested new account number or routing number. When the
	// `change_code` is `incorrect_transaction_code`, this field contains an integer.
	// Numbers starting with a 2 encourage changing the `funding` parameter to
	// checking; numbers starting with a 3 encourage changing to savings.
	CorrectedData string `json:"corrected_data,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the notification occurred.
	CreatedAt time.Time                            `json:"created_at,required" format:"date-time"`
	JSON      achTransferNotificationsOfChangeJSON `json:"-"`
}

// achTransferNotificationsOfChangeJSON contains the JSON metadata for the struct
// [ACHTransferNotificationsOfChange]
type achTransferNotificationsOfChangeJSON struct {
	ChangeCode    apijson.Field
	CorrectedData apijson.Field
	CreatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ACHTransferNotificationsOfChange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r achTransferNotificationsOfChangeJSON) RawJSON() string {
	return r.raw
}

// The required type of change that is being signaled by the receiving financial
// institution.
type ACHTransferNotificationsOfChangeChangeCode string

const (
	ACHTransferNotificationsOfChangeChangeCodeIncorrectAccountNumber                                                      ACHTransferNotificationsOfChangeChangeCode = "incorrect_account_number"
	ACHTransferNotificationsOfChangeChangeCodeIncorrectRoutingNumber                                                      ACHTransferNotificationsOfChangeChangeCode = "incorrect_routing_number"
	ACHTransferNotificationsOfChangeChangeCodeIncorrectRoutingNumberAndAccountNumber                                      ACHTransferNotificationsOfChangeChangeCode = "incorrect_routing_number_and_account_number"
	ACHTransferNotificationsOfChangeChangeCodeIncorrectTransactionCode                                                    ACHTransferNotificationsOfChangeChangeCode = "incorrect_transaction_code"
	ACHTransferNotificationsOfChangeChangeCodeIncorrectAccountNumberAndTransactionCode                                    ACHTransferNotificationsOfChangeChangeCode = "incorrect_account_number_and_transaction_code"
	ACHTransferNotificationsOfChangeChangeCodeIncorrectRoutingNumberAccountNumberAndTransactionCode                       ACHTransferNotificationsOfChangeChangeCode = "incorrect_routing_number_account_number_and_transaction_code"
	ACHTransferNotificationsOfChangeChangeCodeIncorrectReceivingDepositoryFinancialInstitutionIdentification              ACHTransferNotificationsOfChangeChangeCode = "incorrect_receiving_depository_financial_institution_identification"
	ACHTransferNotificationsOfChangeChangeCodeIncorrectIndividualIdentificationNumber                                     ACHTransferNotificationsOfChangeChangeCode = "incorrect_individual_identification_number"
	ACHTransferNotificationsOfChangeChangeCodeAddendaFormatError                                                          ACHTransferNotificationsOfChangeChangeCode = "addenda_format_error"
	ACHTransferNotificationsOfChangeChangeCodeIncorrectStandardEntryClassCodeForOutboundInternationalPayment              ACHTransferNotificationsOfChangeChangeCode = "incorrect_standard_entry_class_code_for_outbound_international_payment"
	ACHTransferNotificationsOfChangeChangeCodeMisroutedNotificationOfChange                                               ACHTransferNotificationsOfChangeChangeCode = "misrouted_notification_of_change"
	ACHTransferNotificationsOfChangeChangeCodeIncorrectTraceNumber                                                        ACHTransferNotificationsOfChangeChangeCode = "incorrect_trace_number"
	ACHTransferNotificationsOfChangeChangeCodeIncorrectCompanyIdentificationNumber                                        ACHTransferNotificationsOfChangeChangeCode = "incorrect_company_identification_number"
	ACHTransferNotificationsOfChangeChangeCodeIncorrectIdentificationNumber                                               ACHTransferNotificationsOfChangeChangeCode = "incorrect_identification_number"
	ACHTransferNotificationsOfChangeChangeCodeIncorrectlyFormattedCorrectedData                                           ACHTransferNotificationsOfChangeChangeCode = "incorrectly_formatted_corrected_data"
	ACHTransferNotificationsOfChangeChangeCodeIncorrectDiscretionaryData                                                  ACHTransferNotificationsOfChangeChangeCode = "incorrect_discretionary_data"
	ACHTransferNotificationsOfChangeChangeCodeRoutingNumberNotFromOriginalEntryDetailRecord                               ACHTransferNotificationsOfChangeChangeCode = "routing_number_not_from_original_entry_detail_record"
	ACHTransferNotificationsOfChangeChangeCodeDepositoryFinancialInstitutionAccountNumberNotFromOriginalEntryDetailRecord ACHTransferNotificationsOfChangeChangeCode = "depository_financial_institution_account_number_not_from_original_entry_detail_record"
	ACHTransferNotificationsOfChangeChangeCodeIncorrectTransactionCodeByOriginatingDepositoryFinancialInstitution         ACHTransferNotificationsOfChangeChangeCode = "incorrect_transaction_code_by_originating_depository_financial_institution"
)

func (r ACHTransferNotificationsOfChangeChangeCode) IsKnown() bool {
	switch r {
	case ACHTransferNotificationsOfChangeChangeCodeIncorrectAccountNumber, ACHTransferNotificationsOfChangeChangeCodeIncorrectRoutingNumber, ACHTransferNotificationsOfChangeChangeCodeIncorrectRoutingNumberAndAccountNumber, ACHTransferNotificationsOfChangeChangeCodeIncorrectTransactionCode, ACHTransferNotificationsOfChangeChangeCodeIncorrectAccountNumberAndTransactionCode, ACHTransferNotificationsOfChangeChangeCodeIncorrectRoutingNumberAccountNumberAndTransactionCode, ACHTransferNotificationsOfChangeChangeCodeIncorrectReceivingDepositoryFinancialInstitutionIdentification, ACHTransferNotificationsOfChangeChangeCodeIncorrectIndividualIdentificationNumber, ACHTransferNotificationsOfChangeChangeCodeAddendaFormatError, ACHTransferNotificationsOfChangeChangeCodeIncorrectStandardEntryClassCodeForOutboundInternationalPayment, ACHTransferNotificationsOfChangeChangeCodeMisroutedNotificationOfChange, ACHTransferNotificationsOfChangeChangeCodeIncorrectTraceNumber, ACHTransferNotificationsOfChangeChangeCodeIncorrectCompanyIdentificationNumber, ACHTransferNotificationsOfChangeChangeCodeIncorrectIdentificationNumber, ACHTransferNotificationsOfChangeChangeCodeIncorrectlyFormattedCorrectedData, ACHTransferNotificationsOfChangeChangeCodeIncorrectDiscretionaryData, ACHTransferNotificationsOfChangeChangeCodeRoutingNumberNotFromOriginalEntryDetailRecord, ACHTransferNotificationsOfChangeChangeCodeDepositoryFinancialInstitutionAccountNumberNotFromOriginalEntryDetailRecord, ACHTransferNotificationsOfChangeChangeCodeIncorrectTransactionCodeByOriginatingDepositoryFinancialInstitution:
		return true
	}
	return false
}

// Configuration for how the effective date of the transfer will be set. This
// determines same-day vs future-dated settlement timing. If not set, defaults to a
// `settlement_schedule` of `same_day`. If set, exactly one of the child attributes
// must be set.
type ACHTransferPreferredEffectiveDate struct {
	// A specific date in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format to
	// use as the effective date when submitting this transfer.
	Date time.Time `json:"date,required,nullable" format:"date"`
	// A schedule by which Increase will choose an effective date for the transfer.
	SettlementSchedule ACHTransferPreferredEffectiveDateSettlementSchedule `json:"settlement_schedule,required,nullable"`
	JSON               achTransferPreferredEffectiveDateJSON               `json:"-"`
}

// achTransferPreferredEffectiveDateJSON contains the JSON metadata for the struct
// [ACHTransferPreferredEffectiveDate]
type achTransferPreferredEffectiveDateJSON struct {
	Date               apijson.Field
	SettlementSchedule apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ACHTransferPreferredEffectiveDate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r achTransferPreferredEffectiveDateJSON) RawJSON() string {
	return r.raw
}

// A schedule by which Increase will choose an effective date for the transfer.
type ACHTransferPreferredEffectiveDateSettlementSchedule string

const (
	ACHTransferPreferredEffectiveDateSettlementScheduleSameDay     ACHTransferPreferredEffectiveDateSettlementSchedule = "same_day"
	ACHTransferPreferredEffectiveDateSettlementScheduleFutureDated ACHTransferPreferredEffectiveDateSettlementSchedule = "future_dated"
)

func (r ACHTransferPreferredEffectiveDateSettlementSchedule) IsKnown() bool {
	switch r {
	case ACHTransferPreferredEffectiveDateSettlementScheduleSameDay, ACHTransferPreferredEffectiveDateSettlementScheduleFutureDated:
		return true
	}
	return false
}

// If your transfer is returned, this will contain details of the return.
type ACHTransferReturn struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The three character ACH return code, in the range R01 to R85.
	RawReturnReasonCode string `json:"raw_return_reason_code,required"`
	// Why the ACH Transfer was returned. This reason code is sent by the receiving
	// bank back to Increase.
	ReturnReasonCode ACHTransferReturnReturnReasonCode `json:"return_reason_code,required"`
	// A 15 digit number that was generated by the bank that initiated the return. The
	// trace number of the return is different than that of the original transfer. ACH
	// trace numbers are not unique, but along with the amount and date this number can
	// be used to identify the ACH return at the bank that initiated it.
	TraceNumber string `json:"trace_number,required"`
	// The identifier of the Transaction associated with this return.
	TransactionID string `json:"transaction_id,required"`
	// The identifier of the ACH Transfer associated with this return.
	TransferID string                `json:"transfer_id,required"`
	JSON       achTransferReturnJSON `json:"-"`
}

// achTransferReturnJSON contains the JSON metadata for the struct
// [ACHTransferReturn]
type achTransferReturnJSON struct {
	CreatedAt           apijson.Field
	RawReturnReasonCode apijson.Field
	ReturnReasonCode    apijson.Field
	TraceNumber         apijson.Field
	TransactionID       apijson.Field
	TransferID          apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ACHTransferReturn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r achTransferReturnJSON) RawJSON() string {
	return r.raw
}

// Why the ACH Transfer was returned. This reason code is sent by the receiving
// bank back to Increase.
type ACHTransferReturnReturnReasonCode string

const (
	ACHTransferReturnReturnReasonCodeInsufficientFund                                            ACHTransferReturnReturnReasonCode = "insufficient_fund"
	ACHTransferReturnReturnReasonCodeNoAccount                                                   ACHTransferReturnReturnReasonCode = "no_account"
	ACHTransferReturnReturnReasonCodeAccountClosed                                               ACHTransferReturnReturnReasonCode = "account_closed"
	ACHTransferReturnReturnReasonCodeInvalidAccountNumberStructure                               ACHTransferReturnReturnReasonCode = "invalid_account_number_structure"
	ACHTransferReturnReturnReasonCodeAccountFrozenEntryReturnedPerOfacInstruction                ACHTransferReturnReturnReasonCode = "account_frozen_entry_returned_per_ofac_instruction"
	ACHTransferReturnReturnReasonCodeCreditEntryRefusedByReceiver                                ACHTransferReturnReturnReasonCode = "credit_entry_refused_by_receiver"
	ACHTransferReturnReturnReasonCodeUnauthorizedDebitToConsumerAccountUsingCorporateSecCode     ACHTransferReturnReturnReasonCode = "unauthorized_debit_to_consumer_account_using_corporate_sec_code"
	ACHTransferReturnReturnReasonCodeCorporateCustomerAdvisedNotAuthorized                       ACHTransferReturnReturnReasonCode = "corporate_customer_advised_not_authorized"
	ACHTransferReturnReturnReasonCodePaymentStopped                                              ACHTransferReturnReturnReasonCode = "payment_stopped"
	ACHTransferReturnReturnReasonCodeNonTransactionAccount                                       ACHTransferReturnReturnReasonCode = "non_transaction_account"
	ACHTransferReturnReturnReasonCodeUncollectedFunds                                            ACHTransferReturnReturnReasonCode = "uncollected_funds"
	ACHTransferReturnReturnReasonCodeRoutingNumberCheckDigitError                                ACHTransferReturnReturnReasonCode = "routing_number_check_digit_error"
	ACHTransferReturnReturnReasonCodeCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete   ACHTransferReturnReturnReasonCode = "customer_advised_unauthorized_improper_ineligible_or_incomplete"
	ACHTransferReturnReturnReasonCodeAmountFieldError                                            ACHTransferReturnReturnReasonCode = "amount_field_error"
	ACHTransferReturnReturnReasonCodeAuthorizationRevokedByCustomer                              ACHTransferReturnReturnReasonCode = "authorization_revoked_by_customer"
	ACHTransferReturnReturnReasonCodeInvalidACHRoutingNumber                                     ACHTransferReturnReturnReasonCode = "invalid_ach_routing_number"
	ACHTransferReturnReturnReasonCodeFileRecordEditCriteria                                      ACHTransferReturnReturnReasonCode = "file_record_edit_criteria"
	ACHTransferReturnReturnReasonCodeEnrInvalidIndividualName                                    ACHTransferReturnReturnReasonCode = "enr_invalid_individual_name"
	ACHTransferReturnReturnReasonCodeReturnedPerOdfiRequest                                      ACHTransferReturnReturnReasonCode = "returned_per_odfi_request"
	ACHTransferReturnReturnReasonCodeLimitedParticipationDfi                                     ACHTransferReturnReturnReasonCode = "limited_participation_dfi"
	ACHTransferReturnReturnReasonCodeIncorrectlyCodedOutboundInternationalPayment                ACHTransferReturnReturnReasonCode = "incorrectly_coded_outbound_international_payment"
	ACHTransferReturnReturnReasonCodeAccountSoldToAnotherDfi                                     ACHTransferReturnReturnReasonCode = "account_sold_to_another_dfi"
	ACHTransferReturnReturnReasonCodeAddendaError                                                ACHTransferReturnReturnReasonCode = "addenda_error"
	ACHTransferReturnReturnReasonCodeBeneficiaryOrAccountHolderDeceased                          ACHTransferReturnReturnReasonCode = "beneficiary_or_account_holder_deceased"
	ACHTransferReturnReturnReasonCodeCustomerAdvisedNotWithinAuthorizationTerms                  ACHTransferReturnReturnReasonCode = "customer_advised_not_within_authorization_terms"
	ACHTransferReturnReturnReasonCodeCorrectedReturn                                             ACHTransferReturnReturnReasonCode = "corrected_return"
	ACHTransferReturnReturnReasonCodeDuplicateEntry                                              ACHTransferReturnReturnReasonCode = "duplicate_entry"
	ACHTransferReturnReturnReasonCodeDuplicateReturn                                             ACHTransferReturnReturnReasonCode = "duplicate_return"
	ACHTransferReturnReturnReasonCodeEnrDuplicateEnrollment                                      ACHTransferReturnReturnReasonCode = "enr_duplicate_enrollment"
	ACHTransferReturnReturnReasonCodeEnrInvalidDfiAccountNumber                                  ACHTransferReturnReturnReasonCode = "enr_invalid_dfi_account_number"
	ACHTransferReturnReturnReasonCodeEnrInvalidIndividualIDNumber                                ACHTransferReturnReturnReasonCode = "enr_invalid_individual_id_number"
	ACHTransferReturnReturnReasonCodeEnrInvalidRepresentativePayeeIndicator                      ACHTransferReturnReturnReasonCode = "enr_invalid_representative_payee_indicator"
	ACHTransferReturnReturnReasonCodeEnrInvalidTransactionCode                                   ACHTransferReturnReturnReasonCode = "enr_invalid_transaction_code"
	ACHTransferReturnReturnReasonCodeEnrReturnOfEnrEntry                                         ACHTransferReturnReturnReasonCode = "enr_return_of_enr_entry"
	ACHTransferReturnReturnReasonCodeEnrRoutingNumberCheckDigitError                             ACHTransferReturnReturnReasonCode = "enr_routing_number_check_digit_error"
	ACHTransferReturnReturnReasonCodeEntryNotProcessedByGateway                                  ACHTransferReturnReturnReasonCode = "entry_not_processed_by_gateway"
	ACHTransferReturnReturnReasonCodeFieldError                                                  ACHTransferReturnReturnReasonCode = "field_error"
	ACHTransferReturnReturnReasonCodeForeignReceivingDfiUnableToSettle                           ACHTransferReturnReturnReasonCode = "foreign_receiving_dfi_unable_to_settle"
	ACHTransferReturnReturnReasonCodeIatEntryCodingError                                         ACHTransferReturnReturnReasonCode = "iat_entry_coding_error"
	ACHTransferReturnReturnReasonCodeImproperEffectiveEntryDate                                  ACHTransferReturnReturnReasonCode = "improper_effective_entry_date"
	ACHTransferReturnReturnReasonCodeImproperSourceDocumentSourceDocumentPresented               ACHTransferReturnReturnReasonCode = "improper_source_document_source_document_presented"
	ACHTransferReturnReturnReasonCodeInvalidCompanyID                                            ACHTransferReturnReturnReasonCode = "invalid_company_id"
	ACHTransferReturnReturnReasonCodeInvalidForeignReceivingDfiIdentification                    ACHTransferReturnReturnReasonCode = "invalid_foreign_receiving_dfi_identification"
	ACHTransferReturnReturnReasonCodeInvalidIndividualIDNumber                                   ACHTransferReturnReturnReasonCode = "invalid_individual_id_number"
	ACHTransferReturnReturnReasonCodeItemAndRckEntryPresentedForPayment                          ACHTransferReturnReturnReasonCode = "item_and_rck_entry_presented_for_payment"
	ACHTransferReturnReturnReasonCodeItemRelatedToRckEntryIsIneligible                           ACHTransferReturnReturnReasonCode = "item_related_to_rck_entry_is_ineligible"
	ACHTransferReturnReturnReasonCodeMandatoryFieldError                                         ACHTransferReturnReturnReasonCode = "mandatory_field_error"
	ACHTransferReturnReturnReasonCodeMisroutedDishonoredReturn                                   ACHTransferReturnReturnReasonCode = "misrouted_dishonored_return"
	ACHTransferReturnReturnReasonCodeMisroutedReturn                                             ACHTransferReturnReturnReasonCode = "misrouted_return"
	ACHTransferReturnReturnReasonCodeNoErrorsFound                                               ACHTransferReturnReturnReasonCode = "no_errors_found"
	ACHTransferReturnReturnReasonCodeNonAcceptanceOfR62DishonoredReturn                          ACHTransferReturnReturnReasonCode = "non_acceptance_of_r62_dishonored_return"
	ACHTransferReturnReturnReasonCodeNonParticipantInIatProgram                                  ACHTransferReturnReturnReasonCode = "non_participant_in_iat_program"
	ACHTransferReturnReturnReasonCodePermissibleReturnEntry                                      ACHTransferReturnReturnReasonCode = "permissible_return_entry"
	ACHTransferReturnReturnReasonCodePermissibleReturnEntryNotAccepted                           ACHTransferReturnReturnReasonCode = "permissible_return_entry_not_accepted"
	ACHTransferReturnReturnReasonCodeRdfiNonSettlement                                           ACHTransferReturnReturnReasonCode = "rdfi_non_settlement"
	ACHTransferReturnReturnReasonCodeRdfiParticipantInCheckTruncationProgram                     ACHTransferReturnReturnReasonCode = "rdfi_participant_in_check_truncation_program"
	ACHTransferReturnReturnReasonCodeRepresentativePayeeDeceasedOrUnableToContinueInThatCapacity ACHTransferReturnReturnReasonCode = "representative_payee_deceased_or_unable_to_continue_in_that_capacity"
	ACHTransferReturnReturnReasonCodeReturnNotADuplicate                                         ACHTransferReturnReturnReasonCode = "return_not_a_duplicate"
	ACHTransferReturnReturnReasonCodeReturnOfErroneousOrReversingDebit                           ACHTransferReturnReturnReasonCode = "return_of_erroneous_or_reversing_debit"
	ACHTransferReturnReturnReasonCodeReturnOfImproperCreditEntry                                 ACHTransferReturnReturnReasonCode = "return_of_improper_credit_entry"
	ACHTransferReturnReturnReasonCodeReturnOfImproperDebitEntry                                  ACHTransferReturnReturnReasonCode = "return_of_improper_debit_entry"
	ACHTransferReturnReturnReasonCodeReturnOfXckEntry                                            ACHTransferReturnReturnReasonCode = "return_of_xck_entry"
	ACHTransferReturnReturnReasonCodeSourceDocumentPresentedForPayment                           ACHTransferReturnReturnReasonCode = "source_document_presented_for_payment"
	ACHTransferReturnReturnReasonCodeStateLawAffectingRckAcceptance                              ACHTransferReturnReturnReasonCode = "state_law_affecting_rck_acceptance"
	ACHTransferReturnReturnReasonCodeStopPaymentOnItemRelatedToRckEntry                          ACHTransferReturnReturnReasonCode = "stop_payment_on_item_related_to_rck_entry"
	ACHTransferReturnReturnReasonCodeStopPaymentOnSourceDocument                                 ACHTransferReturnReturnReasonCode = "stop_payment_on_source_document"
	ACHTransferReturnReturnReasonCodeTimelyOriginalReturn                                        ACHTransferReturnReturnReasonCode = "timely_original_return"
	ACHTransferReturnReturnReasonCodeTraceNumberError                                            ACHTransferReturnReturnReasonCode = "trace_number_error"
	ACHTransferReturnReturnReasonCodeUntimelyDishonoredReturn                                    ACHTransferReturnReturnReasonCode = "untimely_dishonored_return"
	ACHTransferReturnReturnReasonCodeUntimelyReturn                                              ACHTransferReturnReturnReasonCode = "untimely_return"
)

func (r ACHTransferReturnReturnReasonCode) IsKnown() bool {
	switch r {
	case ACHTransferReturnReturnReasonCodeInsufficientFund, ACHTransferReturnReturnReasonCodeNoAccount, ACHTransferReturnReturnReasonCodeAccountClosed, ACHTransferReturnReturnReasonCodeInvalidAccountNumberStructure, ACHTransferReturnReturnReasonCodeAccountFrozenEntryReturnedPerOfacInstruction, ACHTransferReturnReturnReasonCodeCreditEntryRefusedByReceiver, ACHTransferReturnReturnReasonCodeUnauthorizedDebitToConsumerAccountUsingCorporateSecCode, ACHTransferReturnReturnReasonCodeCorporateCustomerAdvisedNotAuthorized, ACHTransferReturnReturnReasonCodePaymentStopped, ACHTransferReturnReturnReasonCodeNonTransactionAccount, ACHTransferReturnReturnReasonCodeUncollectedFunds, ACHTransferReturnReturnReasonCodeRoutingNumberCheckDigitError, ACHTransferReturnReturnReasonCodeCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete, ACHTransferReturnReturnReasonCodeAmountFieldError, ACHTransferReturnReturnReasonCodeAuthorizationRevokedByCustomer, ACHTransferReturnReturnReasonCodeInvalidACHRoutingNumber, ACHTransferReturnReturnReasonCodeFileRecordEditCriteria, ACHTransferReturnReturnReasonCodeEnrInvalidIndividualName, ACHTransferReturnReturnReasonCodeReturnedPerOdfiRequest, ACHTransferReturnReturnReasonCodeLimitedParticipationDfi, ACHTransferReturnReturnReasonCodeIncorrectlyCodedOutboundInternationalPayment, ACHTransferReturnReturnReasonCodeAccountSoldToAnotherDfi, ACHTransferReturnReturnReasonCodeAddendaError, ACHTransferReturnReturnReasonCodeBeneficiaryOrAccountHolderDeceased, ACHTransferReturnReturnReasonCodeCustomerAdvisedNotWithinAuthorizationTerms, ACHTransferReturnReturnReasonCodeCorrectedReturn, ACHTransferReturnReturnReasonCodeDuplicateEntry, ACHTransferReturnReturnReasonCodeDuplicateReturn, ACHTransferReturnReturnReasonCodeEnrDuplicateEnrollment, ACHTransferReturnReturnReasonCodeEnrInvalidDfiAccountNumber, ACHTransferReturnReturnReasonCodeEnrInvalidIndividualIDNumber, ACHTransferReturnReturnReasonCodeEnrInvalidRepresentativePayeeIndicator, ACHTransferReturnReturnReasonCodeEnrInvalidTransactionCode, ACHTransferReturnReturnReasonCodeEnrReturnOfEnrEntry, ACHTransferReturnReturnReasonCodeEnrRoutingNumberCheckDigitError, ACHTransferReturnReturnReasonCodeEntryNotProcessedByGateway, ACHTransferReturnReturnReasonCodeFieldError, ACHTransferReturnReturnReasonCodeForeignReceivingDfiUnableToSettle, ACHTransferReturnReturnReasonCodeIatEntryCodingError, ACHTransferReturnReturnReasonCodeImproperEffectiveEntryDate, ACHTransferReturnReturnReasonCodeImproperSourceDocumentSourceDocumentPresented, ACHTransferReturnReturnReasonCodeInvalidCompanyID, ACHTransferReturnReturnReasonCodeInvalidForeignReceivingDfiIdentification, ACHTransferReturnReturnReasonCodeInvalidIndividualIDNumber, ACHTransferReturnReturnReasonCodeItemAndRckEntryPresentedForPayment, ACHTransferReturnReturnReasonCodeItemRelatedToRckEntryIsIneligible, ACHTransferReturnReturnReasonCodeMandatoryFieldError, ACHTransferReturnReturnReasonCodeMisroutedDishonoredReturn, ACHTransferReturnReturnReasonCodeMisroutedReturn, ACHTransferReturnReturnReasonCodeNoErrorsFound, ACHTransferReturnReturnReasonCodeNonAcceptanceOfR62DishonoredReturn, ACHTransferReturnReturnReasonCodeNonParticipantInIatProgram, ACHTransferReturnReturnReasonCodePermissibleReturnEntry, ACHTransferReturnReturnReasonCodePermissibleReturnEntryNotAccepted, ACHTransferReturnReturnReasonCodeRdfiNonSettlement, ACHTransferReturnReturnReasonCodeRdfiParticipantInCheckTruncationProgram, ACHTransferReturnReturnReasonCodeRepresentativePayeeDeceasedOrUnableToContinueInThatCapacity, ACHTransferReturnReturnReasonCodeReturnNotADuplicate, ACHTransferReturnReturnReasonCodeReturnOfErroneousOrReversingDebit, ACHTransferReturnReturnReasonCodeReturnOfImproperCreditEntry, ACHTransferReturnReturnReasonCodeReturnOfImproperDebitEntry, ACHTransferReturnReturnReasonCodeReturnOfXckEntry, ACHTransferReturnReturnReasonCodeSourceDocumentPresentedForPayment, ACHTransferReturnReturnReasonCodeStateLawAffectingRckAcceptance, ACHTransferReturnReturnReasonCodeStopPaymentOnItemRelatedToRckEntry, ACHTransferReturnReturnReasonCodeStopPaymentOnSourceDocument, ACHTransferReturnReturnReasonCodeTimelyOriginalReturn, ACHTransferReturnReturnReasonCodeTraceNumberError, ACHTransferReturnReturnReasonCodeUntimelyDishonoredReturn, ACHTransferReturnReturnReasonCodeUntimelyReturn:
		return true
	}
	return false
}

// A subhash containing information about when and how the transfer settled at the
// Federal Reserve.
type ACHTransferSettlement struct {
	// When the funds for this transfer have settled at the destination bank at the
	// Federal Reserve.
	SettledAt time.Time                 `json:"settled_at,required" format:"date-time"`
	JSON      achTransferSettlementJSON `json:"-"`
}

// achTransferSettlementJSON contains the JSON metadata for the struct
// [ACHTransferSettlement]
type achTransferSettlementJSON struct {
	SettledAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACHTransferSettlement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r achTransferSettlementJSON) RawJSON() string {
	return r.raw
}

// The Standard Entry Class (SEC) code to use for the transfer.
type ACHTransferStandardEntryClassCode string

const (
	ACHTransferStandardEntryClassCodeCorporateCreditOrDebit        ACHTransferStandardEntryClassCode = "corporate_credit_or_debit"
	ACHTransferStandardEntryClassCodeCorporateTradeExchange        ACHTransferStandardEntryClassCode = "corporate_trade_exchange"
	ACHTransferStandardEntryClassCodePrearrangedPaymentsAndDeposit ACHTransferStandardEntryClassCode = "prearranged_payments_and_deposit"
	ACHTransferStandardEntryClassCodeInternetInitiated             ACHTransferStandardEntryClassCode = "internet_initiated"
)

func (r ACHTransferStandardEntryClassCode) IsKnown() bool {
	switch r {
	case ACHTransferStandardEntryClassCodeCorporateCreditOrDebit, ACHTransferStandardEntryClassCodeCorporateTradeExchange, ACHTransferStandardEntryClassCodePrearrangedPaymentsAndDeposit, ACHTransferStandardEntryClassCodeInternetInitiated:
		return true
	}
	return false
}

// The lifecycle status of the transfer.
type ACHTransferStatus string

const (
	ACHTransferStatusPendingApproval                    ACHTransferStatus = "pending_approval"
	ACHTransferStatusPendingTransferSessionConfirmation ACHTransferStatus = "pending_transfer_session_confirmation"
	ACHTransferStatusCanceled                           ACHTransferStatus = "canceled"
	ACHTransferStatusPendingSubmission                  ACHTransferStatus = "pending_submission"
	ACHTransferStatusPendingReviewing                   ACHTransferStatus = "pending_reviewing"
	ACHTransferStatusRequiresAttention                  ACHTransferStatus = "requires_attention"
	ACHTransferStatusRejected                           ACHTransferStatus = "rejected"
	ACHTransferStatusSubmitted                          ACHTransferStatus = "submitted"
	ACHTransferStatusReturned                           ACHTransferStatus = "returned"
)

func (r ACHTransferStatus) IsKnown() bool {
	switch r {
	case ACHTransferStatusPendingApproval, ACHTransferStatusPendingTransferSessionConfirmation, ACHTransferStatusCanceled, ACHTransferStatusPendingSubmission, ACHTransferStatusPendingReviewing, ACHTransferStatusRequiresAttention, ACHTransferStatusRejected, ACHTransferStatusSubmitted, ACHTransferStatusReturned:
		return true
	}
	return false
}

// After the transfer is submitted to FedACH, this will contain supplemental
// details. Increase batches transfers and submits a file to the Federal Reserve
// roughly every 30 minutes. The Federal Reserve processes ACH transfers during
// weekdays according to their
// [posted schedule](https://www.frbservices.org/resources/resource-centers/same-day-ach/fedach-processing-schedule.html).
type ACHTransferSubmission struct {
	// The timestamp by which any administrative returns are expected to be received
	// by. This follows the NACHA guidelines for return windows, which are: "In
	// general, return entries must be received by the RDFI’s ACH Operator by its
	// deposit deadline for the return entry to be made available to the ODFI no later
	// than the opening of business on the second banking day following the Settlement
	// Date of the original entry.".
	AdministrativeReturnsExpectedBy time.Time `json:"administrative_returns_expected_by,required" format:"date-time"`
	// The ACH transfer's effective date as sent to the Federal Reserve. If a specific
	// date was configured using `preferred_effective_date`, this will match that
	// value. Otherwise, it will be the date selected (following the specified
	// settlement schedule) at the time the transfer was submitted.
	EffectiveDate time.Time `json:"effective_date,required" format:"date"`
	// When the transfer is expected to settle in the recipient's account. Credits may
	// be available sooner, at the receiving banks discretion. The FedACH schedule is
	// published
	// [here](https://www.frbservices.org/resources/resource-centers/same-day-ach/fedach-processing-schedule.html).
	ExpectedFundsSettlementAt time.Time `json:"expected_funds_settlement_at,required" format:"date-time"`
	// The settlement schedule the transfer is expected to follow. This expectation
	// takes into account the `effective_date`, `submitted_at`, and the amount of the
	// transfer.
	ExpectedSettlementSchedule ACHTransferSubmissionExpectedSettlementSchedule `json:"expected_settlement_schedule,required"`
	// When the ACH transfer was sent to FedACH.
	SubmittedAt time.Time `json:"submitted_at,required" format:"date-time"`
	// A 15 digit number recorded in the Nacha file and transmitted to the receiving
	// bank. Along with the amount, date, and originating routing number, this can be
	// used to identify the ACH transfer at the receiving bank. ACH trace numbers are
	// not unique, but are
	// [used to correlate returns](https://increase.com/documentation/ach-returns#ach-returns).
	TraceNumber string                    `json:"trace_number,required"`
	JSON        achTransferSubmissionJSON `json:"-"`
}

// achTransferSubmissionJSON contains the JSON metadata for the struct
// [ACHTransferSubmission]
type achTransferSubmissionJSON struct {
	AdministrativeReturnsExpectedBy apijson.Field
	EffectiveDate                   apijson.Field
	ExpectedFundsSettlementAt       apijson.Field
	ExpectedSettlementSchedule      apijson.Field
	SubmittedAt                     apijson.Field
	TraceNumber                     apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *ACHTransferSubmission) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r achTransferSubmissionJSON) RawJSON() string {
	return r.raw
}

// The settlement schedule the transfer is expected to follow. This expectation
// takes into account the `effective_date`, `submitted_at`, and the amount of the
// transfer.
type ACHTransferSubmissionExpectedSettlementSchedule string

const (
	ACHTransferSubmissionExpectedSettlementScheduleSameDay     ACHTransferSubmissionExpectedSettlementSchedule = "same_day"
	ACHTransferSubmissionExpectedSettlementScheduleFutureDated ACHTransferSubmissionExpectedSettlementSchedule = "future_dated"
)

func (r ACHTransferSubmissionExpectedSettlementSchedule) IsKnown() bool {
	switch r {
	case ACHTransferSubmissionExpectedSettlementScheduleSameDay, ACHTransferSubmissionExpectedSettlementScheduleFutureDated:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `ach_transfer`.
type ACHTransferType string

const (
	ACHTransferTypeACHTransfer ACHTransferType = "ach_transfer"
)

func (r ACHTransferType) IsKnown() bool {
	switch r {
	case ACHTransferTypeACHTransfer:
		return true
	}
	return false
}

type ACHTransferNewParams struct {
	// The Increase identifier for the account that will send the transfer.
	AccountID param.Field[string] `json:"account_id,required"`
	// The transfer amount in USD cents. A positive amount originates a credit transfer
	// pushing funds to the receiving account. A negative amount originates a debit
	// transfer pulling funds from the receiving account.
	Amount param.Field[int64] `json:"amount,required"`
	// A description you choose to give the transfer. This will be saved with the
	// transfer details, displayed in the dashboard, and returned by the API. If
	// `individual_name` and `company_name` are not explicitly set by this API, the
	// `statement_descriptor` will be sent in those fields to the receiving bank to
	// help the customer recognize the transfer. You are highly encouraged to pass
	// `individual_name` and `company_name` instead of relying on this fallback.
	StatementDescriptor param.Field[string] `json:"statement_descriptor,required"`
	// The account number for the destination account.
	AccountNumber param.Field[string] `json:"account_number"`
	// Additional information that will be sent to the recipient. This is included in
	// the transfer data sent to the receiving bank.
	Addenda param.Field[ACHTransferNewParamsAddenda] `json:"addenda"`
	// The description of the date of the transfer, usually in the format `YYMMDD`.
	// This is included in the transfer data sent to the receiving bank.
	CompanyDescriptiveDate param.Field[string] `json:"company_descriptive_date"`
	// The data you choose to associate with the transfer. This is included in the
	// transfer data sent to the receiving bank.
	CompanyDiscretionaryData param.Field[string] `json:"company_discretionary_data"`
	// A description of the transfer. This is included in the transfer data sent to the
	// receiving bank.
	CompanyEntryDescription param.Field[string] `json:"company_entry_description"`
	// The name by which the recipient knows you. This is included in the transfer data
	// sent to the receiving bank.
	CompanyName param.Field[string] `json:"company_name"`
	// The type of entity that owns the account to which the ACH Transfer is being
	// sent.
	DestinationAccountHolder param.Field[ACHTransferNewParamsDestinationAccountHolder] `json:"destination_account_holder"`
	// The ID of an External Account to initiate a transfer to. If this parameter is
	// provided, `account_number`, `routing_number`, and `funding` must be absent.
	ExternalAccountID param.Field[string] `json:"external_account_id"`
	// The type of the account to which the transfer will be sent.
	Funding param.Field[ACHTransferNewParamsFunding] `json:"funding"`
	// Your identifier for the transfer recipient.
	IndividualID param.Field[string] `json:"individual_id"`
	// The name of the transfer recipient. This value is informational and not verified
	// by the recipient's bank.
	IndividualName param.Field[string] `json:"individual_name"`
	// Configuration for how the effective date of the transfer will be set. This
	// determines same-day vs future-dated settlement timing. If not set, defaults to a
	// `settlement_schedule` of `same_day`. If set, exactly one of the child attributes
	// must be set.
	PreferredEffectiveDate param.Field[ACHTransferNewParamsPreferredEffectiveDate] `json:"preferred_effective_date"`
	// Whether the transfer requires explicit approval via the dashboard or API.
	RequireApproval param.Field[bool] `json:"require_approval"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
	// destination account.
	RoutingNumber param.Field[string] `json:"routing_number"`
	// The Standard Entry Class (SEC) code to use for the transfer.
	StandardEntryClassCode param.Field[ACHTransferNewParamsStandardEntryClassCode] `json:"standard_entry_class_code"`
	// The timing of the transaction.
	TransactionTiming param.Field[ACHTransferNewParamsTransactionTiming] `json:"transaction_timing"`
}

func (r ACHTransferNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Additional information that will be sent to the recipient. This is included in
// the transfer data sent to the receiving bank.
type ACHTransferNewParamsAddenda struct {
	// The type of addenda to pass with the transfer.
	Category param.Field[ACHTransferNewParamsAddendaCategory] `json:"category,required"`
	// Unstructured `payment_related_information` passed through with the transfer.
	// Required if and only if `category` is `freeform`.
	Freeform param.Field[ACHTransferNewParamsAddendaFreeform] `json:"freeform"`
	// Structured ASC X12 820 remittance advice records. Please reach out to
	// [support@increase.com](mailto:support@increase.com) for more information.
	// Required if and only if `category` is `payment_order_remittance_advice`.
	PaymentOrderRemittanceAdvice param.Field[ACHTransferNewParamsAddendaPaymentOrderRemittanceAdvice] `json:"payment_order_remittance_advice"`
}

func (r ACHTransferNewParamsAddenda) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of addenda to pass with the transfer.
type ACHTransferNewParamsAddendaCategory string

const (
	ACHTransferNewParamsAddendaCategoryFreeform                     ACHTransferNewParamsAddendaCategory = "freeform"
	ACHTransferNewParamsAddendaCategoryPaymentOrderRemittanceAdvice ACHTransferNewParamsAddendaCategory = "payment_order_remittance_advice"
)

func (r ACHTransferNewParamsAddendaCategory) IsKnown() bool {
	switch r {
	case ACHTransferNewParamsAddendaCategoryFreeform, ACHTransferNewParamsAddendaCategoryPaymentOrderRemittanceAdvice:
		return true
	}
	return false
}

// Unstructured `payment_related_information` passed through with the transfer.
// Required if and only if `category` is `freeform`.
type ACHTransferNewParamsAddendaFreeform struct {
	// Each entry represents an addendum sent with the transfer. In general, you should
	// send at most one addendum–most ACH recipients cannot access beyond the first 80
	// characters sent. Please reach out to
	// [support@increase.com](mailto:support@increase.com) to send 2 or more addenda to
	// a recipient expecting a specific addendum format.
	Entries param.Field[[]ACHTransferNewParamsAddendaFreeformEntry] `json:"entries,required"`
}

func (r ACHTransferNewParamsAddendaFreeform) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ACHTransferNewParamsAddendaFreeformEntry struct {
	// The payment related information passed in the addendum.
	PaymentRelatedInformation param.Field[string] `json:"payment_related_information,required"`
}

func (r ACHTransferNewParamsAddendaFreeformEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Structured ASC X12 820 remittance advice records. Please reach out to
// [support@increase.com](mailto:support@increase.com) for more information.
// Required if and only if `category` is `payment_order_remittance_advice`.
type ACHTransferNewParamsAddendaPaymentOrderRemittanceAdvice struct {
	// ASC X12 RMR records for this specific transfer.
	Invoices param.Field[[]ACHTransferNewParamsAddendaPaymentOrderRemittanceAdviceInvoice] `json:"invoices,required"`
}

func (r ACHTransferNewParamsAddendaPaymentOrderRemittanceAdvice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ACHTransferNewParamsAddendaPaymentOrderRemittanceAdviceInvoice struct {
	// The invoice number for this reference, determined in advance with the receiver.
	InvoiceNumber param.Field[string] `json:"invoice_number,required"`
	// The amount that was paid for this invoice in the minor unit of its currency. For
	// dollars, for example, this is cents.
	PaidAmount param.Field[int64] `json:"paid_amount,required"`
}

func (r ACHTransferNewParamsAddendaPaymentOrderRemittanceAdviceInvoice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of entity that owns the account to which the ACH Transfer is being
// sent.
type ACHTransferNewParamsDestinationAccountHolder string

const (
	ACHTransferNewParamsDestinationAccountHolderBusiness   ACHTransferNewParamsDestinationAccountHolder = "business"
	ACHTransferNewParamsDestinationAccountHolderIndividual ACHTransferNewParamsDestinationAccountHolder = "individual"
	ACHTransferNewParamsDestinationAccountHolderUnknown    ACHTransferNewParamsDestinationAccountHolder = "unknown"
)

func (r ACHTransferNewParamsDestinationAccountHolder) IsKnown() bool {
	switch r {
	case ACHTransferNewParamsDestinationAccountHolderBusiness, ACHTransferNewParamsDestinationAccountHolderIndividual, ACHTransferNewParamsDestinationAccountHolderUnknown:
		return true
	}
	return false
}

// The type of the account to which the transfer will be sent.
type ACHTransferNewParamsFunding string

const (
	ACHTransferNewParamsFundingChecking      ACHTransferNewParamsFunding = "checking"
	ACHTransferNewParamsFundingSavings       ACHTransferNewParamsFunding = "savings"
	ACHTransferNewParamsFundingGeneralLedger ACHTransferNewParamsFunding = "general_ledger"
)

func (r ACHTransferNewParamsFunding) IsKnown() bool {
	switch r {
	case ACHTransferNewParamsFundingChecking, ACHTransferNewParamsFundingSavings, ACHTransferNewParamsFundingGeneralLedger:
		return true
	}
	return false
}

// Configuration for how the effective date of the transfer will be set. This
// determines same-day vs future-dated settlement timing. If not set, defaults to a
// `settlement_schedule` of `same_day`. If set, exactly one of the child attributes
// must be set.
type ACHTransferNewParamsPreferredEffectiveDate struct {
	// A specific date in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format to
	// use as the effective date when submitting this transfer.
	Date param.Field[time.Time] `json:"date" format:"date"`
	// A schedule by which Increase will choose an effective date for the transfer.
	SettlementSchedule param.Field[ACHTransferNewParamsPreferredEffectiveDateSettlementSchedule] `json:"settlement_schedule"`
}

func (r ACHTransferNewParamsPreferredEffectiveDate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A schedule by which Increase will choose an effective date for the transfer.
type ACHTransferNewParamsPreferredEffectiveDateSettlementSchedule string

const (
	ACHTransferNewParamsPreferredEffectiveDateSettlementScheduleSameDay     ACHTransferNewParamsPreferredEffectiveDateSettlementSchedule = "same_day"
	ACHTransferNewParamsPreferredEffectiveDateSettlementScheduleFutureDated ACHTransferNewParamsPreferredEffectiveDateSettlementSchedule = "future_dated"
)

func (r ACHTransferNewParamsPreferredEffectiveDateSettlementSchedule) IsKnown() bool {
	switch r {
	case ACHTransferNewParamsPreferredEffectiveDateSettlementScheduleSameDay, ACHTransferNewParamsPreferredEffectiveDateSettlementScheduleFutureDated:
		return true
	}
	return false
}

// The Standard Entry Class (SEC) code to use for the transfer.
type ACHTransferNewParamsStandardEntryClassCode string

const (
	ACHTransferNewParamsStandardEntryClassCodeCorporateCreditOrDebit        ACHTransferNewParamsStandardEntryClassCode = "corporate_credit_or_debit"
	ACHTransferNewParamsStandardEntryClassCodeCorporateTradeExchange        ACHTransferNewParamsStandardEntryClassCode = "corporate_trade_exchange"
	ACHTransferNewParamsStandardEntryClassCodePrearrangedPaymentsAndDeposit ACHTransferNewParamsStandardEntryClassCode = "prearranged_payments_and_deposit"
	ACHTransferNewParamsStandardEntryClassCodeInternetInitiated             ACHTransferNewParamsStandardEntryClassCode = "internet_initiated"
)

func (r ACHTransferNewParamsStandardEntryClassCode) IsKnown() bool {
	switch r {
	case ACHTransferNewParamsStandardEntryClassCodeCorporateCreditOrDebit, ACHTransferNewParamsStandardEntryClassCodeCorporateTradeExchange, ACHTransferNewParamsStandardEntryClassCodePrearrangedPaymentsAndDeposit, ACHTransferNewParamsStandardEntryClassCodeInternetInitiated:
		return true
	}
	return false
}

// The timing of the transaction.
type ACHTransferNewParamsTransactionTiming string

const (
	ACHTransferNewParamsTransactionTimingSynchronous  ACHTransferNewParamsTransactionTiming = "synchronous"
	ACHTransferNewParamsTransactionTimingAsynchronous ACHTransferNewParamsTransactionTiming = "asynchronous"
)

func (r ACHTransferNewParamsTransactionTiming) IsKnown() bool {
	switch r {
	case ACHTransferNewParamsTransactionTimingSynchronous, ACHTransferNewParamsTransactionTimingAsynchronous:
		return true
	}
	return false
}

type ACHTransferListParams struct {
	// Filter ACH Transfers to those that originated from the specified Account.
	AccountID param.Field[string]                         `query:"account_id"`
	CreatedAt param.Field[ACHTransferListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Filter ACH Transfers to those made to the specified External Account.
	ExternalAccountID param.Field[string] `query:"external_account_id"`
	// Filter records to the one with the specified `idempotency_key` you chose for
	// that object. This value is unique across Increase and is used to ensure that a
	// request is only processed once. Learn more about
	// [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey param.Field[string] `query:"idempotency_key"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit  param.Field[int64]                       `query:"limit"`
	Status param.Field[ACHTransferListParamsStatus] `query:"status"`
}

// URLQuery serializes [ACHTransferListParams]'s query parameters as `url.Values`.
func (r ACHTransferListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ACHTransferListParamsCreatedAt struct {
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

// URLQuery serializes [ACHTransferListParamsCreatedAt]'s query parameters as
// `url.Values`.
func (r ACHTransferListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ACHTransferListParamsStatus struct {
	// Return results whose value is in the provided list. For GET requests, this
	// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
	In param.Field[[]ACHTransferListParamsStatusIn] `query:"in"`
}

// URLQuery serializes [ACHTransferListParamsStatus]'s query parameters as
// `url.Values`.
func (r ACHTransferListParamsStatus) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ACHTransferListParamsStatusIn string

const (
	ACHTransferListParamsStatusInPendingApproval                    ACHTransferListParamsStatusIn = "pending_approval"
	ACHTransferListParamsStatusInPendingTransferSessionConfirmation ACHTransferListParamsStatusIn = "pending_transfer_session_confirmation"
	ACHTransferListParamsStatusInCanceled                           ACHTransferListParamsStatusIn = "canceled"
	ACHTransferListParamsStatusInPendingSubmission                  ACHTransferListParamsStatusIn = "pending_submission"
	ACHTransferListParamsStatusInPendingReviewing                   ACHTransferListParamsStatusIn = "pending_reviewing"
	ACHTransferListParamsStatusInRequiresAttention                  ACHTransferListParamsStatusIn = "requires_attention"
	ACHTransferListParamsStatusInRejected                           ACHTransferListParamsStatusIn = "rejected"
	ACHTransferListParamsStatusInSubmitted                          ACHTransferListParamsStatusIn = "submitted"
	ACHTransferListParamsStatusInReturned                           ACHTransferListParamsStatusIn = "returned"
)

func (r ACHTransferListParamsStatusIn) IsKnown() bool {
	switch r {
	case ACHTransferListParamsStatusInPendingApproval, ACHTransferListParamsStatusInPendingTransferSessionConfirmation, ACHTransferListParamsStatusInCanceled, ACHTransferListParamsStatusInPendingSubmission, ACHTransferListParamsStatusInPendingReviewing, ACHTransferListParamsStatusInRequiresAttention, ACHTransferListParamsStatusInRejected, ACHTransferListParamsStatusInSubmitted, ACHTransferListParamsStatusInReturned:
		return true
	}
	return false
}

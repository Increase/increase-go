// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/apiquery"
	"github.com/increase/increase-go/internal/pagination"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/option"
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
	opts = append(r.Options[:], opts...)
	path := "ach_transfers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve an ACH Transfer
func (r *ACHTransferService) Get(ctx context.Context, achTransferID string, opts ...option.RequestOption) (res *ACHTransfer, err error) {
	opts = append(r.Options[:], opts...)
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
func (r *ACHTransferService) ListAutoPaging(ctx context.Context, query ACHTransferListParams, opts ...option.RequestOption) *pagination.PageAutoPager[ACHTransfer] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Approves an ACH Transfer in a pending_approval state.
func (r *ACHTransferService) Approve(ctx context.Context, achTransferID string, opts ...option.RequestOption) (res *ACHTransfer, err error) {
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options[:], opts...)
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
	// `settlement_schedule` of `same_day`. If set, exactly one of the child atributes
	// must be set.
	PreferredEffectiveDate ACHTransferPreferredEffectiveDate `json:"preferred_effective_date,required"`
	// If your transfer is returned, this will contain details of the return.
	Return ACHTransferReturn `json:"return,required,nullable"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber string `json:"routing_number,required"`
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
	CompanyName              apijson.Field
	CreatedAt                apijson.Field
	CreatedBy                apijson.Field
	Currency                 apijson.Field
	DestinationAccountHolder apijson.Field
	ExternalAccountID        apijson.Field
	Funding                  apijson.Field
	IdempotencyKey           apijson.Field
	IndividualID             apijson.Field
	IndividualName           apijson.Field
	Network                  apijson.Field
	NotificationsOfChange    apijson.Field
	PendingTransactionID     apijson.Field
	PreferredEffectiveDate   apijson.Field
	Return                   apijson.Field
	RoutingNumber            apijson.Field
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
	// Unstructured `payment_related_information` passed through with the transfer.
	ACHTransferAddendaCategoryFreeform ACHTransferAddendaCategory = "freeform"
	// Structured ASC X12 820 remittance advice records. Please reach out to
	// [support@increase.com](mailto:support@increase.com) for more information.
	ACHTransferAddendaCategoryPaymentOrderRemittanceAdvice ACHTransferAddendaCategory = "payment_order_remittance_advice"
	// Unknown addenda type.
	ACHTransferAddendaCategoryOther ACHTransferAddendaCategory = "other"
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
	// An API key. Details will be under the `api_key` object.
	ACHTransferCreatedByCategoryAPIKey ACHTransferCreatedByCategory = "api_key"
	// An OAuth application you connected to Increase. Details will be under the
	// `oauth_application` object.
	ACHTransferCreatedByCategoryOAuthApplication ACHTransferCreatedByCategory = "oauth_application"
	// A User in the Increase dashboard. Details will be under the `user` object.
	ACHTransferCreatedByCategoryUser ACHTransferCreatedByCategory = "user"
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
	// Canadian Dollar (CAD)
	ACHTransferCurrencyCad ACHTransferCurrency = "CAD"
	// Swiss Franc (CHF)
	ACHTransferCurrencyChf ACHTransferCurrency = "CHF"
	// Euro (EUR)
	ACHTransferCurrencyEur ACHTransferCurrency = "EUR"
	// British Pound (GBP)
	ACHTransferCurrencyGbp ACHTransferCurrency = "GBP"
	// Japanese Yen (JPY)
	ACHTransferCurrencyJpy ACHTransferCurrency = "JPY"
	// US Dollar (USD)
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
	// The External Account is owned by a business.
	ACHTransferDestinationAccountHolderBusiness ACHTransferDestinationAccountHolder = "business"
	// The External Account is owned by an individual.
	ACHTransferDestinationAccountHolderIndividual ACHTransferDestinationAccountHolder = "individual"
	// It's unknown what kind of entity owns the External Account.
	ACHTransferDestinationAccountHolderUnknown ACHTransferDestinationAccountHolder = "unknown"
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
	// A checking account.
	ACHTransferFundingChecking ACHTransferFunding = "checking"
	// A savings account.
	ACHTransferFundingSavings ACHTransferFunding = "savings"
)

func (r ACHTransferFunding) IsKnown() bool {
	switch r {
	case ACHTransferFundingChecking, ACHTransferFundingSavings:
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
	// The account number was incorrect.
	ACHTransferNotificationsOfChangeChangeCodeIncorrectAccountNumber ACHTransferNotificationsOfChangeChangeCode = "incorrect_account_number"
	// The routing number was incorrect.
	ACHTransferNotificationsOfChangeChangeCodeIncorrectRoutingNumber ACHTransferNotificationsOfChangeChangeCode = "incorrect_routing_number"
	// Both the routing number and the account number were incorrect.
	ACHTransferNotificationsOfChangeChangeCodeIncorrectRoutingNumberAndAccountNumber ACHTransferNotificationsOfChangeChangeCode = "incorrect_routing_number_and_account_number"
	// The transaction code was incorrect. Try changing the `funding` parameter from
	// checking to savings or vice-versa.
	ACHTransferNotificationsOfChangeChangeCodeIncorrectTransactionCode ACHTransferNotificationsOfChangeChangeCode = "incorrect_transaction_code"
	// The account number and the transaction code were incorrect.
	ACHTransferNotificationsOfChangeChangeCodeIncorrectAccountNumberAndTransactionCode ACHTransferNotificationsOfChangeChangeCode = "incorrect_account_number_and_transaction_code"
	// The routing number, account number, and transaction code were incorrect.
	ACHTransferNotificationsOfChangeChangeCodeIncorrectRoutingNumberAccountNumberAndTransactionCode ACHTransferNotificationsOfChangeChangeCode = "incorrect_routing_number_account_number_and_transaction_code"
	// The receiving depository financial institution identification was incorrect.
	ACHTransferNotificationsOfChangeChangeCodeIncorrectReceivingDepositoryFinancialInstitutionIdentification ACHTransferNotificationsOfChangeChangeCode = "incorrect_receiving_depository_financial_institution_identification"
	// The individual identification number was incorrect.
	ACHTransferNotificationsOfChangeChangeCodeIncorrectIndividualIdentificationNumber ACHTransferNotificationsOfChangeChangeCode = "incorrect_individual_identification_number"
	// The addenda had an incorrect format.
	ACHTransferNotificationsOfChangeChangeCodeAddendaFormatError ACHTransferNotificationsOfChangeChangeCode = "addenda_format_error"
	// The standard entry class code was incorrect for an outbound international
	// payment.
	ACHTransferNotificationsOfChangeChangeCodeIncorrectStandardEntryClassCodeForOutboundInternationalPayment ACHTransferNotificationsOfChangeChangeCode = "incorrect_standard_entry_class_code_for_outbound_international_payment"
	// The notification of change was misrouted.
	ACHTransferNotificationsOfChangeChangeCodeMisroutedNotificationOfChange ACHTransferNotificationsOfChangeChangeCode = "misrouted_notification_of_change"
	// The trace number was incorrect.
	ACHTransferNotificationsOfChangeChangeCodeIncorrectTraceNumber ACHTransferNotificationsOfChangeChangeCode = "incorrect_trace_number"
	// The company identification number was incorrect.
	ACHTransferNotificationsOfChangeChangeCodeIncorrectCompanyIdentificationNumber ACHTransferNotificationsOfChangeChangeCode = "incorrect_company_identification_number"
	// The individual identification number or identification number was incorrect.
	ACHTransferNotificationsOfChangeChangeCodeIncorrectIdentificationNumber ACHTransferNotificationsOfChangeChangeCode = "incorrect_identification_number"
	// The corrected data was incorrectly formatted.
	ACHTransferNotificationsOfChangeChangeCodeIncorrectlyFormattedCorrectedData ACHTransferNotificationsOfChangeChangeCode = "incorrectly_formatted_corrected_data"
	// The discretionary data was incorrect.
	ACHTransferNotificationsOfChangeChangeCodeIncorrectDiscretionaryData ACHTransferNotificationsOfChangeChangeCode = "incorrect_discretionary_data"
	// The routing number was not from the original entry detail record.
	ACHTransferNotificationsOfChangeChangeCodeRoutingNumberNotFromOriginalEntryDetailRecord ACHTransferNotificationsOfChangeChangeCode = "routing_number_not_from_original_entry_detail_record"
	// The depository financial institution account number was not from the original
	// entry detail record.
	ACHTransferNotificationsOfChangeChangeCodeDepositoryFinancialInstitutionAccountNumberNotFromOriginalEntryDetailRecord ACHTransferNotificationsOfChangeChangeCode = "depository_financial_institution_account_number_not_from_original_entry_detail_record"
	// The transaction code was incorrect, initiated by the originating depository
	// financial institution.
	ACHTransferNotificationsOfChangeChangeCodeIncorrectTransactionCodeByOriginatingDepositoryFinancialInstitution ACHTransferNotificationsOfChangeChangeCode = "incorrect_transaction_code_by_originating_depository_financial_institution"
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
// `settlement_schedule` of `same_day`. If set, exactly one of the child atributes
// must be set.
type ACHTransferPreferredEffectiveDate struct {
	// A specific date in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format to
	// use as the effective date when submitting this transfer.
	Date time.Time `json:"date,required,nullable" format:"date"`
	// A schedule by which Increase whill choose an effective date for the transfer.
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

// A schedule by which Increase whill choose an effective date for the transfer.
type ACHTransferPreferredEffectiveDateSettlementSchedule string

const (
	// The chosen effective date will be the same as the ACH processing date on which
	// the transfer is submitted. This is necessary, but not sufficient for the
	// transfer to be settled same-day: it must also be submitted before the last
	// same-day cutoff and be less than or equal to $1,000.000.00.
	ACHTransferPreferredEffectiveDateSettlementScheduleSameDay ACHTransferPreferredEffectiveDateSettlementSchedule = "same_day"
	// The chosen effective date will be the business day following the ACH processing
	// date on which the transfer is submitted. The transfer will be settled on that
	// future day.
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
	// Code R01. Insufficient funds in the receiving account. Sometimes abbreviated to
	// NSF.
	ACHTransferReturnReturnReasonCodeInsufficientFund ACHTransferReturnReturnReasonCode = "insufficient_fund"
	// Code R03. The account does not exist or the receiving bank was unable to locate
	// it.
	ACHTransferReturnReturnReasonCodeNoAccount ACHTransferReturnReturnReasonCode = "no_account"
	// Code R02. The account is closed at the receiving bank.
	ACHTransferReturnReturnReasonCodeAccountClosed ACHTransferReturnReturnReasonCode = "account_closed"
	// Code R04. The account number is invalid at the receiving bank.
	ACHTransferReturnReturnReasonCodeInvalidAccountNumberStructure ACHTransferReturnReturnReasonCode = "invalid_account_number_structure"
	// Code R16. The account at the receiving bank was frozen per the Office of Foreign
	// Assets Control.
	ACHTransferReturnReturnReasonCodeAccountFrozenEntryReturnedPerOfacInstruction ACHTransferReturnReturnReasonCode = "account_frozen_entry_returned_per_ofac_instruction"
	// Code R23. The receiving bank account refused a credit transfer.
	ACHTransferReturnReturnReasonCodeCreditEntryRefusedByReceiver ACHTransferReturnReturnReasonCode = "credit_entry_refused_by_receiver"
	// Code R05. The receiving bank rejected because of an incorrect Standard Entry
	// Class code.
	ACHTransferReturnReturnReasonCodeUnauthorizedDebitToConsumerAccountUsingCorporateSecCode ACHTransferReturnReturnReasonCode = "unauthorized_debit_to_consumer_account_using_corporate_sec_code"
	// Code R29. The corporate customer at the receiving bank reversed the transfer.
	ACHTransferReturnReturnReasonCodeCorporateCustomerAdvisedNotAuthorized ACHTransferReturnReturnReasonCode = "corporate_customer_advised_not_authorized"
	// Code R08. The receiving bank stopped payment on this transfer.
	ACHTransferReturnReturnReasonCodePaymentStopped ACHTransferReturnReturnReasonCode = "payment_stopped"
	// Code R20. The receiving bank account does not perform transfers.
	ACHTransferReturnReturnReasonCodeNonTransactionAccount ACHTransferReturnReturnReasonCode = "non_transaction_account"
	// Code R09. The receiving bank account does not have enough available balance for
	// the transfer.
	ACHTransferReturnReturnReasonCodeUncollectedFunds ACHTransferReturnReturnReasonCode = "uncollected_funds"
	// Code R28. The routing number is incorrect.
	ACHTransferReturnReturnReasonCodeRoutingNumberCheckDigitError ACHTransferReturnReturnReasonCode = "routing_number_check_digit_error"
	// Code R10. The customer at the receiving bank reversed the transfer.
	ACHTransferReturnReturnReasonCodeCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete ACHTransferReturnReturnReasonCode = "customer_advised_unauthorized_improper_ineligible_or_incomplete"
	// Code R19. The amount field is incorrect or too large.
	ACHTransferReturnReturnReasonCodeAmountFieldError ACHTransferReturnReturnReasonCode = "amount_field_error"
	// Code R07. The customer at the receiving institution informed their bank that
	// they have revoked authorization for a previously authorized transfer.
	ACHTransferReturnReturnReasonCodeAuthorizationRevokedByCustomer ACHTransferReturnReturnReasonCode = "authorization_revoked_by_customer"
	// Code R13. The routing number is invalid.
	ACHTransferReturnReturnReasonCodeInvalidACHRoutingNumber ACHTransferReturnReturnReasonCode = "invalid_ach_routing_number"
	// Code R17. The receiving bank is unable to process a field in the transfer.
	ACHTransferReturnReturnReasonCodeFileRecordEditCriteria ACHTransferReturnReturnReasonCode = "file_record_edit_criteria"
	// Code R45. The individual name field was invalid.
	ACHTransferReturnReturnReasonCodeEnrInvalidIndividualName ACHTransferReturnReturnReasonCode = "enr_invalid_individual_name"
	// Code R06. The originating financial institution asked for this transfer to be
	// returned. The receiving bank is complying with the request.
	ACHTransferReturnReturnReasonCodeReturnedPerOdfiRequest ACHTransferReturnReturnReasonCode = "returned_per_odfi_request"
	// Code R34. The receiving bank's regulatory supervisor has limited their
	// participation in the ACH network.
	ACHTransferReturnReturnReasonCodeLimitedParticipationDfi ACHTransferReturnReturnReasonCode = "limited_participation_dfi"
	// Code R85. The outbound international ACH transfer was incorrect.
	ACHTransferReturnReturnReasonCodeIncorrectlyCodedOutboundInternationalPayment ACHTransferReturnReturnReasonCode = "incorrectly_coded_outbound_international_payment"
	// Code R12. A rare return reason. The account was sold to another bank.
	ACHTransferReturnReturnReasonCodeAccountSoldToAnotherDfi ACHTransferReturnReturnReasonCode = "account_sold_to_another_dfi"
	// Code R25. The addenda record is incorrect or missing.
	ACHTransferReturnReturnReasonCodeAddendaError ACHTransferReturnReturnReasonCode = "addenda_error"
	// Code R15. A rare return reason. The account holder is deceased.
	ACHTransferReturnReturnReasonCodeBeneficiaryOrAccountHolderDeceased ACHTransferReturnReturnReasonCode = "beneficiary_or_account_holder_deceased"
	// Code R11. A rare return reason. The customer authorized some payment to the
	// sender, but this payment was not in error.
	ACHTransferReturnReturnReasonCodeCustomerAdvisedNotWithinAuthorizationTerms ACHTransferReturnReturnReasonCode = "customer_advised_not_within_authorization_terms"
	// Code R74. A rare return reason. Sent in response to a return that was returned
	// with code `field_error`. The latest return should include the corrected
	// field(s).
	ACHTransferReturnReturnReasonCodeCorrectedReturn ACHTransferReturnReturnReasonCode = "corrected_return"
	// Code R24. A rare return reason. The receiving bank received an exact duplicate
	// entry with the same trace number and amount.
	ACHTransferReturnReturnReasonCodeDuplicateEntry ACHTransferReturnReturnReasonCode = "duplicate_entry"
	// Code R67. A rare return reason. The return this message refers to was a
	// duplicate.
	ACHTransferReturnReturnReasonCodeDuplicateReturn ACHTransferReturnReturnReasonCode = "duplicate_return"
	// Code R47. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	ACHTransferReturnReturnReasonCodeEnrDuplicateEnrollment ACHTransferReturnReturnReasonCode = "enr_duplicate_enrollment"
	// Code R43. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	ACHTransferReturnReturnReasonCodeEnrInvalidDfiAccountNumber ACHTransferReturnReturnReasonCode = "enr_invalid_dfi_account_number"
	// Code R44. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	ACHTransferReturnReturnReasonCodeEnrInvalidIndividualIDNumber ACHTransferReturnReturnReasonCode = "enr_invalid_individual_id_number"
	// Code R46. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	ACHTransferReturnReturnReasonCodeEnrInvalidRepresentativePayeeIndicator ACHTransferReturnReturnReasonCode = "enr_invalid_representative_payee_indicator"
	// Code R41. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	ACHTransferReturnReturnReasonCodeEnrInvalidTransactionCode ACHTransferReturnReturnReasonCode = "enr_invalid_transaction_code"
	// Code R40. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	ACHTransferReturnReturnReasonCodeEnrReturnOfEnrEntry ACHTransferReturnReturnReasonCode = "enr_return_of_enr_entry"
	// Code R42. A rare return reason. Only used for US Government agency non-monetary
	// automatic enrollment messages.
	ACHTransferReturnReturnReasonCodeEnrRoutingNumberCheckDigitError ACHTransferReturnReturnReasonCode = "enr_routing_number_check_digit_error"
	// Code R84. A rare return reason. The International ACH Transfer cannot be
	// processed by the gateway.
	ACHTransferReturnReturnReasonCodeEntryNotProcessedByGateway ACHTransferReturnReturnReasonCode = "entry_not_processed_by_gateway"
	// Code R69. A rare return reason. One or more of the fields in the ACH were
	// malformed.
	ACHTransferReturnReturnReasonCodeFieldError ACHTransferReturnReturnReasonCode = "field_error"
	// Code R83. A rare return reason. The Foreign receiving bank was unable to settle
	// this ACH transfer.
	ACHTransferReturnReturnReasonCodeForeignReceivingDfiUnableToSettle ACHTransferReturnReturnReasonCode = "foreign_receiving_dfi_unable_to_settle"
	// Code R80. A rare return reason. The International ACH Transfer is malformed.
	ACHTransferReturnReturnReasonCodeIatEntryCodingError ACHTransferReturnReturnReasonCode = "iat_entry_coding_error"
	// Code R18. A rare return reason. The ACH has an improper effective entry date
	// field.
	ACHTransferReturnReturnReasonCodeImproperEffectiveEntryDate ACHTransferReturnReturnReasonCode = "improper_effective_entry_date"
	// Code R39. A rare return reason. The source document related to this ACH, usually
	// an ACH check conversion, was presented to the bank.
	ACHTransferReturnReturnReasonCodeImproperSourceDocumentSourceDocumentPresented ACHTransferReturnReturnReasonCode = "improper_source_document_source_document_presented"
	// Code R21. A rare return reason. The Company ID field of the ACH was invalid.
	ACHTransferReturnReturnReasonCodeInvalidCompanyID ACHTransferReturnReturnReasonCode = "invalid_company_id"
	// Code R82. A rare return reason. The foreign receiving bank identifier for an
	// International ACH Transfer was invalid.
	ACHTransferReturnReturnReasonCodeInvalidForeignReceivingDfiIdentification ACHTransferReturnReturnReasonCode = "invalid_foreign_receiving_dfi_identification"
	// Code R22. A rare return reason. The Individual ID number field of the ACH was
	// invalid.
	ACHTransferReturnReturnReasonCodeInvalidIndividualIDNumber ACHTransferReturnReturnReasonCode = "invalid_individual_id_number"
	// Code R53. A rare return reason. Both the Represented Check ("RCK") entry and the
	// original check were presented to the bank.
	ACHTransferReturnReturnReasonCodeItemAndRckEntryPresentedForPayment ACHTransferReturnReturnReasonCode = "item_and_rck_entry_presented_for_payment"
	// Code R51. A rare return reason. The Represented Check ("RCK") entry is
	// ineligible.
	ACHTransferReturnReturnReasonCodeItemRelatedToRckEntryIsIneligible ACHTransferReturnReturnReasonCode = "item_related_to_rck_entry_is_ineligible"
	// Code R26. A rare return reason. The ACH is missing a required field.
	ACHTransferReturnReturnReasonCodeMandatoryFieldError ACHTransferReturnReturnReasonCode = "mandatory_field_error"
	// Code R71. A rare return reason. The receiving bank does not recognize the
	// routing number in a dishonored return entry.
	ACHTransferReturnReturnReasonCodeMisroutedDishonoredReturn ACHTransferReturnReturnReasonCode = "misrouted_dishonored_return"
	// Code R61. A rare return reason. The receiving bank does not recognize the
	// routing number in a return entry.
	ACHTransferReturnReturnReasonCodeMisroutedReturn ACHTransferReturnReturnReasonCode = "misrouted_return"
	// Code R76. A rare return reason. Sent in response to a return, the bank does not
	// find the errors alleged by the returning bank.
	ACHTransferReturnReturnReasonCodeNoErrorsFound ACHTransferReturnReturnReasonCode = "no_errors_found"
	// Code R77. A rare return reason. The receiving bank does not accept the return of
	// the erroneous debit. The funds are not available at the receiving bank.
	ACHTransferReturnReturnReasonCodeNonAcceptanceOfR62DishonoredReturn ACHTransferReturnReturnReasonCode = "non_acceptance_of_r62_dishonored_return"
	// Code R81. A rare return reason. The receiving bank does not accept International
	// ACH Transfers.
	ACHTransferReturnReturnReasonCodeNonParticipantInIatProgram ACHTransferReturnReturnReasonCode = "non_participant_in_iat_program"
	// Code R31. A rare return reason. A return that has been agreed to be accepted by
	// the receiving bank, despite falling outside of the usual return timeframe.
	ACHTransferReturnReturnReasonCodePermissibleReturnEntry ACHTransferReturnReturnReasonCode = "permissible_return_entry"
	// Code R70. A rare return reason. The receiving bank had not approved this return.
	ACHTransferReturnReturnReasonCodePermissibleReturnEntryNotAccepted ACHTransferReturnReturnReasonCode = "permissible_return_entry_not_accepted"
	// Code R32. A rare return reason. The receiving bank could not settle this
	// transaction.
	ACHTransferReturnReturnReasonCodeRdfiNonSettlement ACHTransferReturnReturnReasonCode = "rdfi_non_settlement"
	// Code R30. A rare return reason. The receiving bank does not accept Check
	// Truncation ACH transfers.
	ACHTransferReturnReturnReasonCodeRdfiParticipantInCheckTruncationProgram ACHTransferReturnReturnReasonCode = "rdfi_participant_in_check_truncation_program"
	// Code R14. A rare return reason. The payee is deceased.
	ACHTransferReturnReturnReasonCodeRepresentativePayeeDeceasedOrUnableToContinueInThatCapacity ACHTransferReturnReturnReasonCode = "representative_payee_deceased_or_unable_to_continue_in_that_capacity"
	// Code R75. A rare return reason. The originating bank disputes that an earlier
	// `duplicate_entry` return was actually a duplicate.
	ACHTransferReturnReturnReasonCodeReturnNotADuplicate ACHTransferReturnReturnReasonCode = "return_not_a_duplicate"
	// Code R62. A rare return reason. The originating financial institution made a
	// mistake and this return corrects it.
	ACHTransferReturnReturnReasonCodeReturnOfErroneousOrReversingDebit ACHTransferReturnReturnReasonCode = "return_of_erroneous_or_reversing_debit"
	// Code R36. A rare return reason. Return of a malformed credit entry.
	ACHTransferReturnReturnReasonCodeReturnOfImproperCreditEntry ACHTransferReturnReturnReasonCode = "return_of_improper_credit_entry"
	// Code R35. A rare return reason. Return of a malformed debit entry.
	ACHTransferReturnReturnReasonCodeReturnOfImproperDebitEntry ACHTransferReturnReturnReasonCode = "return_of_improper_debit_entry"
	// Code R33. A rare return reason. Return of a Destroyed Check ("XKC") entry.
	ACHTransferReturnReturnReasonCodeReturnOfXckEntry ACHTransferReturnReturnReasonCode = "return_of_xck_entry"
	// Code R37. A rare return reason. The source document related to this ACH, usually
	// an ACH check conversion, was presented to the bank.
	ACHTransferReturnReturnReasonCodeSourceDocumentPresentedForPayment ACHTransferReturnReturnReasonCode = "source_document_presented_for_payment"
	// Code R50. A rare return reason. State law prevents the bank from accepting the
	// Represented Check ("RCK") entry.
	ACHTransferReturnReturnReasonCodeStateLawAffectingRckAcceptance ACHTransferReturnReturnReasonCode = "state_law_affecting_rck_acceptance"
	// Code R52. A rare return reason. A stop payment was issued on a Represented Check
	// ("RCK") entry.
	ACHTransferReturnReturnReasonCodeStopPaymentOnItemRelatedToRckEntry ACHTransferReturnReturnReasonCode = "stop_payment_on_item_related_to_rck_entry"
	// Code R38. A rare return reason. The source attached to the ACH, usually an ACH
	// check conversion, includes a stop payment.
	ACHTransferReturnReturnReasonCodeStopPaymentOnSourceDocument ACHTransferReturnReturnReasonCode = "stop_payment_on_source_document"
	// Code R73. A rare return reason. The bank receiving an `untimely_return` believes
	// it was on time.
	ACHTransferReturnReturnReasonCodeTimelyOriginalReturn ACHTransferReturnReturnReasonCode = "timely_original_return"
	// Code R27. A rare return reason. An ACH return's trace number does not match an
	// originated ACH.
	ACHTransferReturnReturnReasonCodeTraceNumberError ACHTransferReturnReturnReasonCode = "trace_number_error"
	// Code R72. A rare return reason. The dishonored return was sent too late.
	ACHTransferReturnReturnReasonCodeUntimelyDishonoredReturn ACHTransferReturnReturnReasonCode = "untimely_dishonored_return"
	// Code R68. A rare return reason. The return was sent too late.
	ACHTransferReturnReturnReasonCodeUntimelyReturn ACHTransferReturnReturnReasonCode = "untimely_return"
)

func (r ACHTransferReturnReturnReasonCode) IsKnown() bool {
	switch r {
	case ACHTransferReturnReturnReasonCodeInsufficientFund, ACHTransferReturnReturnReasonCodeNoAccount, ACHTransferReturnReturnReasonCodeAccountClosed, ACHTransferReturnReturnReasonCodeInvalidAccountNumberStructure, ACHTransferReturnReturnReasonCodeAccountFrozenEntryReturnedPerOfacInstruction, ACHTransferReturnReturnReasonCodeCreditEntryRefusedByReceiver, ACHTransferReturnReturnReasonCodeUnauthorizedDebitToConsumerAccountUsingCorporateSecCode, ACHTransferReturnReturnReasonCodeCorporateCustomerAdvisedNotAuthorized, ACHTransferReturnReturnReasonCodePaymentStopped, ACHTransferReturnReturnReasonCodeNonTransactionAccount, ACHTransferReturnReturnReasonCodeUncollectedFunds, ACHTransferReturnReturnReasonCodeRoutingNumberCheckDigitError, ACHTransferReturnReturnReasonCodeCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete, ACHTransferReturnReturnReasonCodeAmountFieldError, ACHTransferReturnReturnReasonCodeAuthorizationRevokedByCustomer, ACHTransferReturnReturnReasonCodeInvalidACHRoutingNumber, ACHTransferReturnReturnReasonCodeFileRecordEditCriteria, ACHTransferReturnReturnReasonCodeEnrInvalidIndividualName, ACHTransferReturnReturnReasonCodeReturnedPerOdfiRequest, ACHTransferReturnReturnReasonCodeLimitedParticipationDfi, ACHTransferReturnReturnReasonCodeIncorrectlyCodedOutboundInternationalPayment, ACHTransferReturnReturnReasonCodeAccountSoldToAnotherDfi, ACHTransferReturnReturnReasonCodeAddendaError, ACHTransferReturnReturnReasonCodeBeneficiaryOrAccountHolderDeceased, ACHTransferReturnReturnReasonCodeCustomerAdvisedNotWithinAuthorizationTerms, ACHTransferReturnReturnReasonCodeCorrectedReturn, ACHTransferReturnReturnReasonCodeDuplicateEntry, ACHTransferReturnReturnReasonCodeDuplicateReturn, ACHTransferReturnReturnReasonCodeEnrDuplicateEnrollment, ACHTransferReturnReturnReasonCodeEnrInvalidDfiAccountNumber, ACHTransferReturnReturnReasonCodeEnrInvalidIndividualIDNumber, ACHTransferReturnReturnReasonCodeEnrInvalidRepresentativePayeeIndicator, ACHTransferReturnReturnReasonCodeEnrInvalidTransactionCode, ACHTransferReturnReturnReasonCodeEnrReturnOfEnrEntry, ACHTransferReturnReturnReasonCodeEnrRoutingNumberCheckDigitError, ACHTransferReturnReturnReasonCodeEntryNotProcessedByGateway, ACHTransferReturnReturnReasonCodeFieldError, ACHTransferReturnReturnReasonCodeForeignReceivingDfiUnableToSettle, ACHTransferReturnReturnReasonCodeIatEntryCodingError, ACHTransferReturnReturnReasonCodeImproperEffectiveEntryDate, ACHTransferReturnReturnReasonCodeImproperSourceDocumentSourceDocumentPresented, ACHTransferReturnReturnReasonCodeInvalidCompanyID, ACHTransferReturnReturnReasonCodeInvalidForeignReceivingDfiIdentification, ACHTransferReturnReturnReasonCodeInvalidIndividualIDNumber, ACHTransferReturnReturnReasonCodeItemAndRckEntryPresentedForPayment, ACHTransferReturnReturnReasonCodeItemRelatedToRckEntryIsIneligible, ACHTransferReturnReturnReasonCodeMandatoryFieldError, ACHTransferReturnReturnReasonCodeMisroutedDishonoredReturn, ACHTransferReturnReturnReasonCodeMisroutedReturn, ACHTransferReturnReturnReasonCodeNoErrorsFound, ACHTransferReturnReturnReasonCodeNonAcceptanceOfR62DishonoredReturn, ACHTransferReturnReturnReasonCodeNonParticipantInIatProgram, ACHTransferReturnReturnReasonCodePermissibleReturnEntry, ACHTransferReturnReturnReasonCodePermissibleReturnEntryNotAccepted, ACHTransferReturnReturnReasonCodeRdfiNonSettlement, ACHTransferReturnReturnReasonCodeRdfiParticipantInCheckTruncationProgram, ACHTransferReturnReturnReasonCodeRepresentativePayeeDeceasedOrUnableToContinueInThatCapacity, ACHTransferReturnReturnReasonCodeReturnNotADuplicate, ACHTransferReturnReturnReasonCodeReturnOfErroneousOrReversingDebit, ACHTransferReturnReturnReasonCodeReturnOfImproperCreditEntry, ACHTransferReturnReturnReasonCodeReturnOfImproperDebitEntry, ACHTransferReturnReturnReasonCodeReturnOfXckEntry, ACHTransferReturnReturnReasonCodeSourceDocumentPresentedForPayment, ACHTransferReturnReturnReasonCodeStateLawAffectingRckAcceptance, ACHTransferReturnReturnReasonCodeStopPaymentOnItemRelatedToRckEntry, ACHTransferReturnReturnReasonCodeStopPaymentOnSourceDocument, ACHTransferReturnReturnReasonCodeTimelyOriginalReturn, ACHTransferReturnReturnReasonCodeTraceNumberError, ACHTransferReturnReturnReasonCodeUntimelyDishonoredReturn, ACHTransferReturnReturnReasonCodeUntimelyReturn:
		return true
	}
	return false
}

// The Standard Entry Class (SEC) code to use for the transfer.
type ACHTransferStandardEntryClassCode string

const (
	// Corporate Credit and Debit (CCD).
	ACHTransferStandardEntryClassCodeCorporateCreditOrDebit ACHTransferStandardEntryClassCode = "corporate_credit_or_debit"
	// Corporate Trade Exchange (CTX).
	ACHTransferStandardEntryClassCodeCorporateTradeExchange ACHTransferStandardEntryClassCode = "corporate_trade_exchange"
	// Prearranged Payments and Deposits (PPD).
	ACHTransferStandardEntryClassCodePrearrangedPaymentsAndDeposit ACHTransferStandardEntryClassCode = "prearranged_payments_and_deposit"
	// Internet Initiated (WEB).
	ACHTransferStandardEntryClassCodeInternetInitiated ACHTransferStandardEntryClassCode = "internet_initiated"
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
	// The transfer is pending approval.
	ACHTransferStatusPendingApproval ACHTransferStatus = "pending_approval"
	// The transfer has been canceled.
	ACHTransferStatusCanceled ACHTransferStatus = "canceled"
	// The transfer is pending review by Increase.
	ACHTransferStatusPendingReviewing ACHTransferStatus = "pending_reviewing"
	// The transfer is pending submission to the Federal Reserve.
	ACHTransferStatusPendingSubmission ACHTransferStatus = "pending_submission"
	// The transfer is complete.
	ACHTransferStatusSubmitted ACHTransferStatus = "submitted"
	// The transfer has been returned.
	ACHTransferStatusReturned ACHTransferStatus = "returned"
	// The transfer requires attention from an Increase operator.
	ACHTransferStatusRequiresAttention ACHTransferStatus = "requires_attention"
	// The transfer has been rejected.
	ACHTransferStatusRejected ACHTransferStatus = "rejected"
)

func (r ACHTransferStatus) IsKnown() bool {
	switch r {
	case ACHTransferStatusPendingApproval, ACHTransferStatusCanceled, ACHTransferStatusPendingReviewing, ACHTransferStatusPendingSubmission, ACHTransferStatusSubmitted, ACHTransferStatusReturned, ACHTransferStatusRequiresAttention, ACHTransferStatusRejected:
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
	EffectiveDate              apijson.Field
	ExpectedFundsSettlementAt  apijson.Field
	ExpectedSettlementSchedule apijson.Field
	SubmittedAt                apijson.Field
	TraceNumber                apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
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
	// The transfer is expected to settle same-day.
	ACHTransferSubmissionExpectedSettlementScheduleSameDay ACHTransferSubmissionExpectedSettlementSchedule = "same_day"
	// The transfer is expected to settle on a future date.
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
	// The transfer amount in cents. A positive amount originates a credit transfer
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
	// The transfer effective date in
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	EffectiveDate param.Field[time.Time] `json:"effective_date" format:"date"`
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
	// `settlement_schedule` of `same_day`. If set, exactly one of the child atributes
	// must be set.
	PreferredEffectiveDate param.Field[ACHTransferNewParamsPreferredEffectiveDate] `json:"preferred_effective_date"`
	// Whether the transfer requires explicit approval via the dashboard or API.
	RequireApproval param.Field[bool] `json:"require_approval"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
	// destination account.
	RoutingNumber param.Field[string] `json:"routing_number"`
	// The Standard Entry Class (SEC) code to use for the transfer.
	StandardEntryClassCode param.Field[ACHTransferNewParamsStandardEntryClassCode] `json:"standard_entry_class_code"`
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
	Freeform param.Field[ACHTransferNewParamsAddendaFreeform] `json:"freeform"`
	// Structured ASC X12 820 remittance advice records. Please reach out to
	// [support@increase.com](mailto:support@increase.com) for more information.
	PaymentOrderRemittanceAdvice param.Field[ACHTransferNewParamsAddendaPaymentOrderRemittanceAdvice] `json:"payment_order_remittance_advice"`
}

func (r ACHTransferNewParamsAddenda) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of addenda to pass with the transfer.
type ACHTransferNewParamsAddendaCategory string

const (
	// Unstructured `payment_related_information` passed through with the transfer.
	ACHTransferNewParamsAddendaCategoryFreeform ACHTransferNewParamsAddendaCategory = "freeform"
	// Structured ASC X12 820 remittance advice records. Please reach out to
	// [support@increase.com](mailto:support@increase.com) for more information.
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
type ACHTransferNewParamsAddendaFreeform struct {
	// Each entry represents an addendum sent with the transfer. Please reach out to
	// [support@increase.com](mailto:support@increase.com) to send more than one
	// addendum.
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
	// The External Account is owned by a business.
	ACHTransferNewParamsDestinationAccountHolderBusiness ACHTransferNewParamsDestinationAccountHolder = "business"
	// The External Account is owned by an individual.
	ACHTransferNewParamsDestinationAccountHolderIndividual ACHTransferNewParamsDestinationAccountHolder = "individual"
	// It's unknown what kind of entity owns the External Account.
	ACHTransferNewParamsDestinationAccountHolderUnknown ACHTransferNewParamsDestinationAccountHolder = "unknown"
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
	// A checking account.
	ACHTransferNewParamsFundingChecking ACHTransferNewParamsFunding = "checking"
	// A savings account.
	ACHTransferNewParamsFundingSavings ACHTransferNewParamsFunding = "savings"
)

func (r ACHTransferNewParamsFunding) IsKnown() bool {
	switch r {
	case ACHTransferNewParamsFundingChecking, ACHTransferNewParamsFundingSavings:
		return true
	}
	return false
}

// Configuration for how the effective date of the transfer will be set. This
// determines same-day vs future-dated settlement timing. If not set, defaults to a
// `settlement_schedule` of `same_day`. If set, exactly one of the child atributes
// must be set.
type ACHTransferNewParamsPreferredEffectiveDate struct {
	// A specific date in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format to
	// use as the effective date when submitting this transfer.
	Date param.Field[time.Time] `json:"date" format:"date"`
	// A schedule by which Increase whill choose an effective date for the transfer.
	SettlementSchedule param.Field[ACHTransferNewParamsPreferredEffectiveDateSettlementSchedule] `json:"settlement_schedule"`
}

func (r ACHTransferNewParamsPreferredEffectiveDate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A schedule by which Increase whill choose an effective date for the transfer.
type ACHTransferNewParamsPreferredEffectiveDateSettlementSchedule string

const (
	// The chosen effective date will be the same as the ACH processing date on which
	// the transfer is submitted. This is necessary, but not sufficient for the
	// transfer to be settled same-day: it must also be submitted before the last
	// same-day cutoff and be less than or equal to $1,000.000.00.
	ACHTransferNewParamsPreferredEffectiveDateSettlementScheduleSameDay ACHTransferNewParamsPreferredEffectiveDateSettlementSchedule = "same_day"
	// The chosen effective date will be the business day following the ACH processing
	// date on which the transfer is submitted. The transfer will be settled on that
	// future day.
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
	// Corporate Credit and Debit (CCD).
	ACHTransferNewParamsStandardEntryClassCodeCorporateCreditOrDebit ACHTransferNewParamsStandardEntryClassCode = "corporate_credit_or_debit"
	// Corporate Trade Exchange (CTX).
	ACHTransferNewParamsStandardEntryClassCodeCorporateTradeExchange ACHTransferNewParamsStandardEntryClassCode = "corporate_trade_exchange"
	// Prearranged Payments and Deposits (PPD).
	ACHTransferNewParamsStandardEntryClassCodePrearrangedPaymentsAndDeposit ACHTransferNewParamsStandardEntryClassCode = "prearranged_payments_and_deposit"
	// Internet Initiated (WEB).
	ACHTransferNewParamsStandardEntryClassCodeInternetInitiated ACHTransferNewParamsStandardEntryClassCode = "internet_initiated"
)

func (r ACHTransferNewParamsStandardEntryClassCode) IsKnown() bool {
	switch r {
	case ACHTransferNewParamsStandardEntryClassCodeCorporateCreditOrDebit, ACHTransferNewParamsStandardEntryClassCodeCorporateTradeExchange, ACHTransferNewParamsStandardEntryClassCodePrearrangedPaymentsAndDeposit, ACHTransferNewParamsStandardEntryClassCodeInternetInitiated:
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
	Limit param.Field[int64] `query:"limit"`
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

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
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

// CheckTransferService contains methods and other services that help with
// interacting with the increase API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewCheckTransferService] method
// instead.
type CheckTransferService struct {
	Options []option.RequestOption
}

// NewCheckTransferService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCheckTransferService(opts ...option.RequestOption) (r *CheckTransferService) {
	r = &CheckTransferService{}
	r.Options = opts
	return
}

// Create a Check Transfer
func (r *CheckTransferService) New(ctx context.Context, body CheckTransferNewParams, opts ...option.RequestOption) (res *CheckTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := "check_transfers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a Check Transfer
func (r *CheckTransferService) Get(ctx context.Context, checkTransferID string, opts ...option.RequestOption) (res *CheckTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("check_transfers/%s", checkTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Check Transfers
func (r *CheckTransferService) List(ctx context.Context, query CheckTransferListParams, opts ...option.RequestOption) (res *pagination.Page[CheckTransfer], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "check_transfers"
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

// List Check Transfers
func (r *CheckTransferService) ListAutoPaging(ctx context.Context, query CheckTransferListParams, opts ...option.RequestOption) *pagination.PageAutoPager[CheckTransfer] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Approve a Check Transfer
func (r *CheckTransferService) Approve(ctx context.Context, checkTransferID string, opts ...option.RequestOption) (res *CheckTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("check_transfers/%s/approve", checkTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Cancel a pending Check Transfer
func (r *CheckTransferService) Cancel(ctx context.Context, checkTransferID string, opts ...option.RequestOption) (res *CheckTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("check_transfers/%s/cancel", checkTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Request a stop payment on a Check Transfer
func (r *CheckTransferService) StopPayment(ctx context.Context, checkTransferID string, body CheckTransferStopPaymentParams, opts ...option.RequestOption) (res *CheckTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("check_transfers/%s/stop_payment", checkTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Check Transfers move funds from your Increase account by mailing a physical
// check.
type CheckTransfer struct {
	// The Check transfer's identifier.
	ID string `json:"id,required"`
	// The identifier of the Account from which funds will be transferred.
	AccountID string `json:"account_id,required"`
	// The account number printed on the check.
	AccountNumber string `json:"account_number,required"`
	// The transfer amount in USD cents.
	Amount int64 `json:"amount,required"`
	// If your account requires approvals for transfers and the transfer was approved,
	// this will contain details of the approval.
	Approval CheckTransferApproval `json:"approval,required,nullable"`
	// If your account requires approvals for transfers and the transfer was not
	// approved, this will contain details of the cancellation.
	Cancellation CheckTransferCancellation `json:"cancellation,required,nullable"`
	// The check number printed on the check.
	CheckNumber string `json:"check_number,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
	// currency.
	Currency CheckTransferCurrency `json:"currency,required"`
	// After a check transfer is deposited, this will contain supplemental details.
	Deposit CheckTransferDeposit `json:"deposit,required,nullable"`
	// Whether Increase will print and mail the check or if you will do it yourself.
	FulfillmentMethod CheckTransferFulfillmentMethod `json:"fulfillment_method,required"`
	// The idempotency key you chose for this object. This value is unique across
	// Increase and is used to ensure that a request is only processed once. Learn more
	// about [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey string `json:"idempotency_key,required,nullable"`
	// If the check has been mailed by Increase, this will contain details of the
	// shipment.
	Mailing CheckTransferMailing `json:"mailing,required,nullable"`
	// The ID for the pending transaction representing the transfer. A pending
	// transaction is created when the transfer
	// [requires approval](https://increase.com/documentation/transfer-approvals#transfer-approvals)
	// by someone else in your organization.
	PendingTransactionID string `json:"pending_transaction_id,required,nullable"`
	// Details relating to the physical check that Increase will print and mail. Will
	// be present if and only if `fulfillment_method` is equal to `physical_check`.
	PhysicalCheck CheckTransferPhysicalCheck `json:"physical_check,required,nullable"`
	// The routing number printed on the check.
	RoutingNumber string `json:"routing_number,required"`
	// The identifier of the Account Number from which to send the transfer and print
	// on the check.
	SourceAccountNumberID string `json:"source_account_number_id,required,nullable"`
	// The lifecycle status of the transfer.
	Status CheckTransferStatus `json:"status,required"`
	// After a stop-payment is requested on the check, this will contain supplemental
	// details.
	StopPaymentRequest CheckTransferStopPaymentRequest `json:"stop_payment_request,required,nullable"`
	// After the transfer is submitted, this will contain supplemental details.
	Submission CheckTransferSubmission `json:"submission,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `check_transfer`.
	Type CheckTransferType `json:"type,required"`
	JSON checkTransferJSON `json:"-"`
}

// checkTransferJSON contains the JSON metadata for the struct [CheckTransfer]
type checkTransferJSON struct {
	ID                    apijson.Field
	AccountID             apijson.Field
	AccountNumber         apijson.Field
	Amount                apijson.Field
	Approval              apijson.Field
	Cancellation          apijson.Field
	CheckNumber           apijson.Field
	CreatedAt             apijson.Field
	Currency              apijson.Field
	Deposit               apijson.Field
	FulfillmentMethod     apijson.Field
	IdempotencyKey        apijson.Field
	Mailing               apijson.Field
	PendingTransactionID  apijson.Field
	PhysicalCheck         apijson.Field
	RoutingNumber         apijson.Field
	SourceAccountNumberID apijson.Field
	Status                apijson.Field
	StopPaymentRequest    apijson.Field
	Submission            apijson.Field
	Type                  apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *CheckTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkTransferJSON) RawJSON() string {
	return r.raw
}

// If your account requires approvals for transfers and the transfer was approved,
// this will contain details of the approval.
type CheckTransferApproval struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was approved.
	ApprovedAt time.Time `json:"approved_at,required" format:"date-time"`
	// If the Transfer was approved by a user in the dashboard, the email address of
	// that user.
	ApprovedBy string                    `json:"approved_by,required,nullable"`
	JSON       checkTransferApprovalJSON `json:"-"`
}

// checkTransferApprovalJSON contains the JSON metadata for the struct
// [CheckTransferApproval]
type checkTransferApprovalJSON struct {
	ApprovedAt  apijson.Field
	ApprovedBy  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CheckTransferApproval) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkTransferApprovalJSON) RawJSON() string {
	return r.raw
}

// If your account requires approvals for transfers and the transfer was not
// approved, this will contain details of the cancellation.
type CheckTransferCancellation struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Transfer was canceled.
	CanceledAt time.Time `json:"canceled_at,required" format:"date-time"`
	// If the Transfer was canceled by a user in the dashboard, the email address of
	// that user.
	CanceledBy string                        `json:"canceled_by,required,nullable"`
	JSON       checkTransferCancellationJSON `json:"-"`
}

// checkTransferCancellationJSON contains the JSON metadata for the struct
// [CheckTransferCancellation]
type checkTransferCancellationJSON struct {
	CanceledAt  apijson.Field
	CanceledBy  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CheckTransferCancellation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkTransferCancellationJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
// currency.
type CheckTransferCurrency string

const (
	// Canadian Dollar (CAD)
	CheckTransferCurrencyCad CheckTransferCurrency = "CAD"
	// Swiss Franc (CHF)
	CheckTransferCurrencyChf CheckTransferCurrency = "CHF"
	// Euro (EUR)
	CheckTransferCurrencyEur CheckTransferCurrency = "EUR"
	// British Pound (GBP)
	CheckTransferCurrencyGbp CheckTransferCurrency = "GBP"
	// Japanese Yen (JPY)
	CheckTransferCurrencyJpy CheckTransferCurrency = "JPY"
	// US Dollar (USD)
	CheckTransferCurrencyUsd CheckTransferCurrency = "USD"
)

func (r CheckTransferCurrency) IsKnown() bool {
	switch r {
	case CheckTransferCurrencyCad, CheckTransferCurrencyChf, CheckTransferCurrencyEur, CheckTransferCurrencyGbp, CheckTransferCurrencyJpy, CheckTransferCurrencyUsd:
		return true
	}
	return false
}

// After a check transfer is deposited, this will contain supplemental details.
type CheckTransferDeposit struct {
	// The identifier of the API File object containing an image of the back of the
	// deposited check.
	BackImageFileID string `json:"back_image_file_id,required,nullable"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
	// bank depositing this check. In some rare cases, this is not transmitted via
	// Check21 and the value will be null.
	BankOfFirstDepositRoutingNumber string `json:"bank_of_first_deposit_routing_number,required,nullable"`
	// When the check was deposited.
	DepositedAt time.Time `json:"deposited_at,required" format:"date-time"`
	// The identifier of the API File object containing an image of the front of the
	// deposited check.
	FrontImageFileID string `json:"front_image_file_id,required,nullable"`
	// The identifier of the Transaction object created when the check was deposited.
	TransactionID string `json:"transaction_id,required,nullable"`
	// The identifier of the Check Transfer object that was deposited.
	TransferID string `json:"transfer_id,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `check_transfer_deposit`.
	Type CheckTransferDepositType `json:"type,required"`
	JSON checkTransferDepositJSON `json:"-"`
}

// checkTransferDepositJSON contains the JSON metadata for the struct
// [CheckTransferDeposit]
type checkTransferDepositJSON struct {
	BackImageFileID                 apijson.Field
	BankOfFirstDepositRoutingNumber apijson.Field
	DepositedAt                     apijson.Field
	FrontImageFileID                apijson.Field
	TransactionID                   apijson.Field
	TransferID                      apijson.Field
	Type                            apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *CheckTransferDeposit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkTransferDepositJSON) RawJSON() string {
	return r.raw
}

// A constant representing the object's type. For this resource it will always be
// `check_transfer_deposit`.
type CheckTransferDepositType string

const (
	CheckTransferDepositTypeCheckTransferDeposit CheckTransferDepositType = "check_transfer_deposit"
)

func (r CheckTransferDepositType) IsKnown() bool {
	switch r {
	case CheckTransferDepositTypeCheckTransferDeposit:
		return true
	}
	return false
}

// Whether Increase will print and mail the check or if you will do it yourself.
type CheckTransferFulfillmentMethod string

const (
	// Increase will print and mail a physical check.
	CheckTransferFulfillmentMethodPhysicalCheck CheckTransferFulfillmentMethod = "physical_check"
	// Increase will not print a check; you are responsible for printing and mailing a
	// check with the provided account number, routing number, check number, and
	// amount.
	CheckTransferFulfillmentMethodThirdParty CheckTransferFulfillmentMethod = "third_party"
)

func (r CheckTransferFulfillmentMethod) IsKnown() bool {
	switch r {
	case CheckTransferFulfillmentMethodPhysicalCheck, CheckTransferFulfillmentMethodThirdParty:
		return true
	}
	return false
}

// If the check has been mailed by Increase, this will contain details of the
// shipment.
type CheckTransferMailing struct {
	// The ID of the file corresponding to an image of the check that was mailed, if
	// available.
	ImageID string `json:"image_id,required,nullable"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check was mailed.
	MailedAt time.Time                `json:"mailed_at,required" format:"date-time"`
	JSON     checkTransferMailingJSON `json:"-"`
}

// checkTransferMailingJSON contains the JSON metadata for the struct
// [CheckTransferMailing]
type checkTransferMailingJSON struct {
	ImageID     apijson.Field
	MailedAt    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CheckTransferMailing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkTransferMailingJSON) RawJSON() string {
	return r.raw
}

// Details relating to the physical check that Increase will print and mail. Will
// be present if and only if `fulfillment_method` is equal to `physical_check`.
type CheckTransferPhysicalCheck struct {
	// Details for where Increase will mail the check.
	MailingAddress CheckTransferPhysicalCheckMailingAddress `json:"mailing_address,required"`
	// The descriptor that will be printed on the memo field on the check.
	Memo string `json:"memo,required,nullable"`
	// The descriptor that will be printed on the letter included with the check.
	Note string `json:"note,required,nullable"`
	// The name that will be printed on the check.
	RecipientName string `json:"recipient_name,required"`
	// The return address to be printed on the check.
	ReturnAddress CheckTransferPhysicalCheckReturnAddress `json:"return_address,required,nullable"`
	JSON          checkTransferPhysicalCheckJSON          `json:"-"`
}

// checkTransferPhysicalCheckJSON contains the JSON metadata for the struct
// [CheckTransferPhysicalCheck]
type checkTransferPhysicalCheckJSON struct {
	MailingAddress apijson.Field
	Memo           apijson.Field
	Note           apijson.Field
	RecipientName  apijson.Field
	ReturnAddress  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CheckTransferPhysicalCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkTransferPhysicalCheckJSON) RawJSON() string {
	return r.raw
}

// Details for where Increase will mail the check.
type CheckTransferPhysicalCheckMailingAddress struct {
	// The city of the check's destination.
	City string `json:"city,required,nullable"`
	// The street address of the check's destination.
	Line1 string `json:"line1,required,nullable"`
	// The second line of the address of the check's destination.
	Line2 string `json:"line2,required,nullable"`
	// The name component of the check's mailing address.
	Name string `json:"name,required,nullable"`
	// The postal code of the check's destination.
	PostalCode string `json:"postal_code,required,nullable"`
	// The state of the check's destination.
	State string                                       `json:"state,required,nullable"`
	JSON  checkTransferPhysicalCheckMailingAddressJSON `json:"-"`
}

// checkTransferPhysicalCheckMailingAddressJSON contains the JSON metadata for the
// struct [CheckTransferPhysicalCheckMailingAddress]
type checkTransferPhysicalCheckMailingAddressJSON struct {
	City        apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	Name        apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CheckTransferPhysicalCheckMailingAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkTransferPhysicalCheckMailingAddressJSON) RawJSON() string {
	return r.raw
}

// The return address to be printed on the check.
type CheckTransferPhysicalCheckReturnAddress struct {
	// The city of the check's destination.
	City string `json:"city,required,nullable"`
	// The street address of the check's destination.
	Line1 string `json:"line1,required,nullable"`
	// The second line of the address of the check's destination.
	Line2 string `json:"line2,required,nullable"`
	// The name component of the check's return address.
	Name string `json:"name,required,nullable"`
	// The postal code of the check's destination.
	PostalCode string `json:"postal_code,required,nullable"`
	// The state of the check's destination.
	State string                                      `json:"state,required,nullable"`
	JSON  checkTransferPhysicalCheckReturnAddressJSON `json:"-"`
}

// checkTransferPhysicalCheckReturnAddressJSON contains the JSON metadata for the
// struct [CheckTransferPhysicalCheckReturnAddress]
type checkTransferPhysicalCheckReturnAddressJSON struct {
	City        apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	Name        apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CheckTransferPhysicalCheckReturnAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkTransferPhysicalCheckReturnAddressJSON) RawJSON() string {
	return r.raw
}

// The lifecycle status of the transfer.
type CheckTransferStatus string

const (
	// The transfer is awaiting approval.
	CheckTransferStatusPendingApproval CheckTransferStatus = "pending_approval"
	// The transfer is pending submission.
	CheckTransferStatusPendingSubmission CheckTransferStatus = "pending_submission"
	// The check is queued for mailing.
	CheckTransferStatusPendingMailing CheckTransferStatus = "pending_mailing"
	// The check has been mailed.
	CheckTransferStatusMailed CheckTransferStatus = "mailed"
	// The transfer has been canceled.
	CheckTransferStatusCanceled CheckTransferStatus = "canceled"
	// The check has been deposited.
	CheckTransferStatusDeposited CheckTransferStatus = "deposited"
	// A stop-payment was requested for this check.
	CheckTransferStatusStopped CheckTransferStatus = "stopped"
	// The transfer has been rejected.
	CheckTransferStatusRejected CheckTransferStatus = "rejected"
	// The transfer requires attention from an Increase operator.
	CheckTransferStatusRequiresAttention CheckTransferStatus = "requires_attention"
	// The transfer has been returned.
	CheckTransferStatusReturned CheckTransferStatus = "returned"
)

func (r CheckTransferStatus) IsKnown() bool {
	switch r {
	case CheckTransferStatusPendingApproval, CheckTransferStatusPendingSubmission, CheckTransferStatusPendingMailing, CheckTransferStatusMailed, CheckTransferStatusCanceled, CheckTransferStatusDeposited, CheckTransferStatusStopped, CheckTransferStatusRejected, CheckTransferStatusRequiresAttention, CheckTransferStatusReturned:
		return true
	}
	return false
}

// After a stop-payment is requested on the check, this will contain supplemental
// details.
type CheckTransferStopPaymentRequest struct {
	// The reason why this transfer was stopped.
	Reason CheckTransferStopPaymentRequestReason `json:"reason,required"`
	// The time the stop-payment was requested.
	RequestedAt time.Time `json:"requested_at,required" format:"date-time"`
	// The ID of the check transfer that was stopped.
	TransferID string `json:"transfer_id,required"`
	// A constant representing the object's type. For this resource it will always be
	// `check_transfer_stop_payment_request`.
	Type CheckTransferStopPaymentRequestType `json:"type,required"`
	JSON checkTransferStopPaymentRequestJSON `json:"-"`
}

// checkTransferStopPaymentRequestJSON contains the JSON metadata for the struct
// [CheckTransferStopPaymentRequest]
type checkTransferStopPaymentRequestJSON struct {
	Reason      apijson.Field
	RequestedAt apijson.Field
	TransferID  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CheckTransferStopPaymentRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkTransferStopPaymentRequestJSON) RawJSON() string {
	return r.raw
}

// The reason why this transfer was stopped.
type CheckTransferStopPaymentRequestReason string

const (
	// The check could not be delivered.
	CheckTransferStopPaymentRequestReasonMailDeliveryFailed CheckTransferStopPaymentRequestReason = "mail_delivery_failed"
	// The check was canceled by an Increase operator who will provide details
	// out-of-band.
	CheckTransferStopPaymentRequestReasonRejectedByIncrease CheckTransferStopPaymentRequestReason = "rejected_by_increase"
	// The check was not authorized.
	CheckTransferStopPaymentRequestReasonNotAuthorized CheckTransferStopPaymentRequestReason = "not_authorized"
	// The check was stopped for another reason.
	CheckTransferStopPaymentRequestReasonUnknown CheckTransferStopPaymentRequestReason = "unknown"
)

func (r CheckTransferStopPaymentRequestReason) IsKnown() bool {
	switch r {
	case CheckTransferStopPaymentRequestReasonMailDeliveryFailed, CheckTransferStopPaymentRequestReasonRejectedByIncrease, CheckTransferStopPaymentRequestReasonNotAuthorized, CheckTransferStopPaymentRequestReasonUnknown:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `check_transfer_stop_payment_request`.
type CheckTransferStopPaymentRequestType string

const (
	CheckTransferStopPaymentRequestTypeCheckTransferStopPaymentRequest CheckTransferStopPaymentRequestType = "check_transfer_stop_payment_request"
)

func (r CheckTransferStopPaymentRequestType) IsKnown() bool {
	switch r {
	case CheckTransferStopPaymentRequestTypeCheckTransferStopPaymentRequest:
		return true
	}
	return false
}

// After the transfer is submitted, this will contain supplemental details.
type CheckTransferSubmission struct {
	// When this check transfer was submitted to our check printer.
	SubmittedAt time.Time                   `json:"submitted_at,required" format:"date-time"`
	JSON        checkTransferSubmissionJSON `json:"-"`
}

// checkTransferSubmissionJSON contains the JSON metadata for the struct
// [CheckTransferSubmission]
type checkTransferSubmissionJSON struct {
	SubmittedAt apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CheckTransferSubmission) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkTransferSubmissionJSON) RawJSON() string {
	return r.raw
}

// A constant representing the object's type. For this resource it will always be
// `check_transfer`.
type CheckTransferType string

const (
	CheckTransferTypeCheckTransfer CheckTransferType = "check_transfer"
)

func (r CheckTransferType) IsKnown() bool {
	switch r {
	case CheckTransferTypeCheckTransfer:
		return true
	}
	return false
}

type CheckTransferNewParams struct {
	// The identifier for the account that will send the transfer.
	AccountID param.Field[string] `json:"account_id,required"`
	// The transfer amount in cents.
	Amount param.Field[int64] `json:"amount,required"`
	// The identifier of the Account Number from which to send the transfer and print
	// on the check.
	SourceAccountNumberID param.Field[string] `json:"source_account_number_id,required"`
	// Whether Increase will print and mail the check or if you will do it yourself.
	FulfillmentMethod param.Field[CheckTransferNewParamsFulfillmentMethod] `json:"fulfillment_method"`
	// Details relating to the physical check that Increase will print and mail. This
	// is required if `fulfillment_method` is equal to `physical_check`. It must not be
	// included if any other `fulfillment_method` is provided.
	PhysicalCheck param.Field[CheckTransferNewParamsPhysicalCheck] `json:"physical_check"`
	// Whether the transfer requires explicit approval via the dashboard or API.
	RequireApproval param.Field[bool] `json:"require_approval"`
}

func (r CheckTransferNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether Increase will print and mail the check or if you will do it yourself.
type CheckTransferNewParamsFulfillmentMethod string

const (
	// Increase will print and mail a physical check.
	CheckTransferNewParamsFulfillmentMethodPhysicalCheck CheckTransferNewParamsFulfillmentMethod = "physical_check"
	// Increase will not print a check; you are responsible for printing and mailing a
	// check with the provided account number, routing number, check number, and
	// amount.
	CheckTransferNewParamsFulfillmentMethodThirdParty CheckTransferNewParamsFulfillmentMethod = "third_party"
)

func (r CheckTransferNewParamsFulfillmentMethod) IsKnown() bool {
	switch r {
	case CheckTransferNewParamsFulfillmentMethodPhysicalCheck, CheckTransferNewParamsFulfillmentMethodThirdParty:
		return true
	}
	return false
}

// Details relating to the physical check that Increase will print and mail. This
// is required if `fulfillment_method` is equal to `physical_check`. It must not be
// included if any other `fulfillment_method` is provided.
type CheckTransferNewParamsPhysicalCheck struct {
	// Details for where Increase will mail the check.
	MailingAddress param.Field[CheckTransferNewParamsPhysicalCheckMailingAddress] `json:"mailing_address,required"`
	// The descriptor that will be printed on the memo field on the check.
	Memo param.Field[string] `json:"memo,required"`
	// The name that will be printed on the check in the 'To:' field.
	RecipientName param.Field[string] `json:"recipient_name,required"`
	// The descriptor that will be printed on the letter included with the check.
	Note param.Field[string] `json:"note"`
	// The return address to be printed on the check. If omitted this will default to
	// the address of the Entity of the Account used to make the Check Transfer.
	ReturnAddress param.Field[CheckTransferNewParamsPhysicalCheckReturnAddress] `json:"return_address"`
}

func (r CheckTransferNewParamsPhysicalCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Details for where Increase will mail the check.
type CheckTransferNewParamsPhysicalCheckMailingAddress struct {
	// The city component of the check's destination address.
	City param.Field[string] `json:"city,required"`
	// The first line of the address component of the check's destination address.
	Line1 param.Field[string] `json:"line1,required"`
	// The postal code component of the check's destination address.
	PostalCode param.Field[string] `json:"postal_code,required"`
	// The US state component of the check's destination address.
	State param.Field[string] `json:"state,required"`
	// The second line of the address component of the check's destination address.
	Line2 param.Field[string] `json:"line2"`
	// The name component of the check's destination address. Defaults to the provided
	// `recipient_name` parameter.
	Name param.Field[string] `json:"name"`
}

func (r CheckTransferNewParamsPhysicalCheckMailingAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The return address to be printed on the check. If omitted this will default to
// the address of the Entity of the Account used to make the Check Transfer.
type CheckTransferNewParamsPhysicalCheckReturnAddress struct {
	// The city of the return address.
	City param.Field[string] `json:"city,required"`
	// The first line of the return address.
	Line1 param.Field[string] `json:"line1,required"`
	// The name of the return address.
	Name param.Field[string] `json:"name,required"`
	// The postal code of the return address.
	PostalCode param.Field[string] `json:"postal_code,required"`
	// The US state of the return address.
	State param.Field[string] `json:"state,required"`
	// The second line of the return address.
	Line2 param.Field[string] `json:"line2"`
}

func (r CheckTransferNewParamsPhysicalCheckReturnAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CheckTransferListParams struct {
	// Filter Check Transfers to those that originated from the specified Account.
	AccountID param.Field[string]                           `query:"account_id"`
	CreatedAt param.Field[CheckTransferListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Filter records to the one with the specified `idempotency_key` you chose for
	// that object. This value is unique across Increase and is used to ensure that a
	// request is only processed once. Learn more about
	// [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey param.Field[string] `query:"idempotency_key"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [CheckTransferListParams]'s query parameters as
// `url.Values`.
func (r CheckTransferListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type CheckTransferListParamsCreatedAt struct {
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

// URLQuery serializes [CheckTransferListParamsCreatedAt]'s query parameters as
// `url.Values`.
func (r CheckTransferListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type CheckTransferStopPaymentParams struct {
	// The reason why this transfer should be stopped.
	Reason param.Field[CheckTransferStopPaymentParamsReason] `json:"reason"`
}

func (r CheckTransferStopPaymentParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The reason why this transfer should be stopped.
type CheckTransferStopPaymentParamsReason string

const (
	// The check could not be delivered.
	CheckTransferStopPaymentParamsReasonMailDeliveryFailed CheckTransferStopPaymentParamsReason = "mail_delivery_failed"
	// The check was not authorized.
	CheckTransferStopPaymentParamsReasonNotAuthorized CheckTransferStopPaymentParamsReason = "not_authorized"
	// The check was stopped for another reason.
	CheckTransferStopPaymentParamsReasonUnknown CheckTransferStopPaymentParamsReason = "unknown"
)

func (r CheckTransferStopPaymentParamsReason) IsKnown() bool {
	switch r {
	case CheckTransferStopPaymentParamsReasonMailDeliveryFailed, CheckTransferStopPaymentParamsReasonNotAuthorized, CheckTransferStopPaymentParamsReasonUnknown:
		return true
	}
	return false
}

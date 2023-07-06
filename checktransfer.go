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
func (r *CheckTransferService) List(ctx context.Context, query CheckTransferListParams, opts ...option.RequestOption) (res *shared.Page[CheckTransfer], err error) {
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
func (r *CheckTransferService) ListAutoPaging(ctx context.Context, query CheckTransferListParams, opts ...option.RequestOption) *shared.PageAutoPager[CheckTransfer] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
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
	// The city of the check's destination.
	AddressCity string `json:"address_city,required,nullable"`
	// The street address of the check's destination.
	AddressLine1 string `json:"address_line1,required,nullable"`
	// The second line of the address of the check's destination.
	AddressLine2 string `json:"address_line2,required,nullable"`
	// The state of the check's destination.
	AddressState string `json:"address_state,required,nullable"`
	// The postal code of the check's destination.
	AddressZip string `json:"address_zip,required,nullable"`
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
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check was mailed.
	MailedAt time.Time `json:"mailed_at,required,nullable" format:"date-time"`
	// The descriptor that will be printed on the memo field on the check.
	Message string `json:"message,required,nullable"`
	// The descriptor that will be printed on the letter included with the check.
	Note string `json:"note,required,nullable"`
	// The identifier of the Pending Transaction associated with the check's creation.
	PendingTransactionID string `json:"pending_transaction_id,required,nullable"`
	// The name that will be printed on the check.
	RecipientName string `json:"recipient_name,required,nullable"`
	// The return address to be printed on the check.
	ReturnAddress CheckTransferReturnAddress `json:"return_address,required,nullable"`
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
	JSON checkTransferJSON
}

// checkTransferJSON contains the JSON metadata for the struct [CheckTransfer]
type checkTransferJSON struct {
	ID                    apijson.Field
	AccountID             apijson.Field
	AccountNumber         apijson.Field
	AddressCity           apijson.Field
	AddressLine1          apijson.Field
	AddressLine2          apijson.Field
	AddressState          apijson.Field
	AddressZip            apijson.Field
	Amount                apijson.Field
	Approval              apijson.Field
	Cancellation          apijson.Field
	CheckNumber           apijson.Field
	CreatedAt             apijson.Field
	Currency              apijson.Field
	Deposit               apijson.Field
	MailedAt              apijson.Field
	Message               apijson.Field
	Note                  apijson.Field
	PendingTransactionID  apijson.Field
	RecipientName         apijson.Field
	ReturnAddress         apijson.Field
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

// If your account requires approvals for transfers and the transfer was approved,
// this will contain details of the approval.
type CheckTransferApproval struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was approved.
	ApprovedAt time.Time `json:"approved_at,required" format:"date-time"`
	// If the Transfer was approved by a user in the dashboard, the email address of
	// that user.
	ApprovedBy string `json:"approved_by,required,nullable"`
	JSON       checkTransferApprovalJSON
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

// If your account requires approvals for transfers and the transfer was not
// approved, this will contain details of the cancellation.
type CheckTransferCancellation struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Transfer was canceled.
	CanceledAt time.Time `json:"canceled_at,required" format:"date-time"`
	// If the Transfer was canceled by a user in the dashboard, the email address of
	// that user.
	CanceledBy string `json:"canceled_by,required,nullable"`
	JSON       checkTransferCancellationJSON
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

// After a check transfer is deposited, this will contain supplemental details.
type CheckTransferDeposit struct {
	// The identifier of the API File object containing an image of the back of the
	// deposited check.
	BackImageFileID string `json:"back_image_file_id,required,nullable"`
	// When the check was deposited.
	DepositedAt time.Time `json:"deposited_at,required" format:"date-time"`
	// The identifier of the API File object containing an image of the front of the
	// deposited check.
	FrontImageFileID string `json:"front_image_file_id,required,nullable"`
	// The identifier of the Transaction object created when the check was deposited.
	TransactionID string `json:"transaction_id,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `check_transfer_deposit`.
	Type CheckTransferDepositType `json:"type,required"`
	JSON checkTransferDepositJSON
}

// checkTransferDepositJSON contains the JSON metadata for the struct
// [CheckTransferDeposit]
type checkTransferDepositJSON struct {
	BackImageFileID  apijson.Field
	DepositedAt      apijson.Field
	FrontImageFileID apijson.Field
	TransactionID    apijson.Field
	Type             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CheckTransferDeposit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A constant representing the object's type. For this resource it will always be
// `check_transfer_deposit`.
type CheckTransferDepositType string

const (
	CheckTransferDepositTypeCheckTransferDeposit CheckTransferDepositType = "check_transfer_deposit"
)

// The return address to be printed on the check.
type CheckTransferReturnAddress struct {
	// The city of the address.
	City string `json:"city,required,nullable"`
	// The first line of the address.
	Line1 string `json:"line1,required,nullable"`
	// The second line of the address.
	Line2 string `json:"line2,required,nullable"`
	// The name of the address.
	Name string `json:"name,required,nullable"`
	// The US state of the address.
	State string `json:"state,required,nullable"`
	// The postal code of the address.
	Zip  string `json:"zip,required,nullable"`
	JSON checkTransferReturnAddressJSON
}

// checkTransferReturnAddressJSON contains the JSON metadata for the struct
// [CheckTransferReturnAddress]
type checkTransferReturnAddressJSON struct {
	City        apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	Name        apijson.Field
	State       apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CheckTransferReturnAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The lifecycle status of the transfer.
type CheckTransferStatus string

const (
	// The transfer is awaiting approval.
	CheckTransferStatusPendingApproval CheckTransferStatus = "pending_approval"
	// The transfer is pending submission.
	CheckTransferStatusPendingSubmission CheckTransferStatus = "pending_submission"
	// The transfer is complete.
	CheckTransferStatusSubmitted CheckTransferStatus = "submitted"
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
)

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
	JSON checkTransferStopPaymentRequestJSON
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

// The reason why this transfer was stopped.
type CheckTransferStopPaymentRequestReason string

const (
	// The check could not be delivered.
	CheckTransferStopPaymentRequestReasonMailDeliveryFailed CheckTransferStopPaymentRequestReason = "mail_delivery_failed"
	// The check was stopped for another reason.
	CheckTransferStopPaymentRequestReasonUnknown CheckTransferStopPaymentRequestReason = "unknown"
)

// A constant representing the object's type. For this resource it will always be
// `check_transfer_stop_payment_request`.
type CheckTransferStopPaymentRequestType string

const (
	CheckTransferStopPaymentRequestTypeCheckTransferStopPaymentRequest CheckTransferStopPaymentRequestType = "check_transfer_stop_payment_request"
)

// After the transfer is submitted, this will contain supplemental details.
type CheckTransferSubmission struct {
	// When this check transfer was submitted to our check printer.
	SubmittedAt time.Time `json:"submitted_at,required" format:"date-time"`
	JSON        checkTransferSubmissionJSON
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

// A constant representing the object's type. For this resource it will always be
// `check_transfer`.
type CheckTransferType string

const (
	CheckTransferTypeCheckTransfer CheckTransferType = "check_transfer"
)

type CheckTransferNewParams struct {
	// The identifier for the account that will send the transfer.
	AccountID param.Field[string] `json:"account_id,required"`
	// The city of the check's destination.
	AddressCity param.Field[string] `json:"address_city,required"`
	// The street address of the check's destination.
	AddressLine1 param.Field[string] `json:"address_line1,required"`
	// The state of the check's destination.
	AddressState param.Field[string] `json:"address_state,required"`
	// The postal code of the check's destination.
	AddressZip param.Field[string] `json:"address_zip,required"`
	// The transfer amount in cents.
	Amount param.Field[int64] `json:"amount,required"`
	// The descriptor that will be printed on the memo field on the check.
	Message param.Field[string] `json:"message,required"`
	// The name that will be printed on the check.
	RecipientName param.Field[string] `json:"recipient_name,required"`
	// The second line of the address of the check's destination.
	AddressLine2 param.Field[string] `json:"address_line2"`
	// The descriptor that will be printed on the letter included with the check.
	Note param.Field[string] `json:"note"`
	// Whether the transfer requires explicit approval via the dashboard or API.
	RequireApproval param.Field[bool] `json:"require_approval"`
	// The return address to be printed on the check. If omitted this will default to
	// the address of the Entity of the Account used to make the Check Transfer.
	ReturnAddress param.Field[CheckTransferNewParamsReturnAddress] `json:"return_address"`
	// The identifier of the Account Number from which to send the transfer and print
	// on the check.
	SourceAccountNumberID param.Field[string] `json:"source_account_number_id"`
}

func (r CheckTransferNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The return address to be printed on the check. If omitted this will default to
// the address of the Entity of the Account used to make the Check Transfer.
type CheckTransferNewParamsReturnAddress struct {
	// The city of the return address.
	City param.Field[string] `json:"city,required"`
	// The first line of the return address.
	Line1 param.Field[string] `json:"line1,required"`
	// The name of the return address.
	Name param.Field[string] `json:"name,required"`
	// The US state of the return address.
	State param.Field[string] `json:"state,required"`
	// The postal code of the return address.
	Zip param.Field[string] `json:"zip,required"`
	// The second line of the return address.
	Line2 param.Field[string] `json:"line2"`
}

func (r CheckTransferNewParamsReturnAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CheckTransferListParams struct {
	// Filter Check Transfers to those that originated from the specified Account.
	AccountID param.Field[string]                           `query:"account_id"`
	CreatedAt param.Field[CheckTransferListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
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
	// The check was stopped for another reason.
	CheckTransferStopPaymentParamsReasonUnknown CheckTransferStopPaymentParamsReason = "unknown"
)

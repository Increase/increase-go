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

type CheckTransferService struct {
	Options []option.RequestOption
}

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
func (r *CheckTransferService) Get(ctx context.Context, check_transfer_id string, opts ...option.RequestOption) (res *CheckTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("check_transfers/%s", check_transfer_id)
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
func (r *CheckTransferService) Approve(ctx context.Context, check_transfer_id string, opts ...option.RequestOption) (res *CheckTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("check_transfers/%s/approve", check_transfer_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Cancel a pending Check Transfer
func (r *CheckTransferService) Cancel(ctx context.Context, check_transfer_id string, opts ...option.RequestOption) (res *CheckTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("check_transfers/%s/cancel", check_transfer_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Request a stop payment on a Check Transfer
func (r *CheckTransferService) StopPayment(ctx context.Context, check_transfer_id string, opts ...option.RequestOption) (res *CheckTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("check_transfers/%s/stop_payment", check_transfer_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type CheckTransfer struct {
	// The identifier of the Account from which funds will be transferred.
	AccountID string `json:"account_id,required"`
	// The street address of the check's destination.
	AddressLine1 string `json:"address_line1,required"`
	// The second line of the address of the check's destination.
	AddressLine2 string `json:"address_line2,required,nullable"`
	// The city of the check's destination.
	AddressCity string `json:"address_city,required"`
	// The state of the check's destination.
	AddressState string `json:"address_state,required"`
	// The postal code of the check's destination.
	AddressZip string `json:"address_zip,required"`
	// The return address to be printed on the check.
	ReturnAddress CheckTransferReturnAddress `json:"return_address,required,nullable"`
	// The transfer amount in USD cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
	// currency.
	Currency CheckTransferCurrency `json:"currency,required"`
	// If your account requires approvals for transfers and the transfer was approved,
	// this will contain details of the approval.
	Approval CheckTransferApproval `json:"approval,required,nullable"`
	// If your account requires approvals for transfers and the transfer was not
	// approved, this will contain details of the cancellation.
	Cancellation CheckTransferCancellation `json:"cancellation,required,nullable"`
	// The Check transfer's identifier.
	ID string `json:"id,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check was mailed.
	MailedAt time.Time `json:"mailed_at,required,nullable" format:"date-time"`
	// The descriptor that will be printed on the memo field on the check.
	Message string `json:"message,required"`
	// The descriptor that will be printed on the letter included with the check.
	Note string `json:"note,required,nullable"`
	// The name that will be printed on the check.
	RecipientName string `json:"recipient_name,required"`
	// The lifecycle status of the transfer.
	Status CheckTransferStatus `json:"status,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check was submitted.
	SubmittedAt time.Time `json:"submitted_at,required,nullable" format:"date-time"`
	// After the transfer is submitted, this will contain supplemental details.
	Submission CheckTransferSubmission `json:"submission,required,nullable"`
	// The ID for the transaction caused by the transfer.
	TransactionID string `json:"transaction_id,required,nullable"`
	// After a stop-payment is requested on the check, this will contain supplemental
	// details.
	StopPaymentRequest CheckTransferStopPaymentRequest `json:"stop_payment_request,required,nullable"`
	// After a check transfer is deposited, this will contain supplemental details.
	Deposit CheckTransferDeposit `json:"deposit,required,nullable"`
	// After a check transfer is returned, this will contain supplemental details. A
	// check transfer is returned when the receiver mails a never deposited check back
	// to the bank printed on the check.
	ReturnDetails CheckTransferReturnDetails `json:"return_details,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `check_transfer`.
	Type CheckTransferType `json:"type,required"`
	JSON CheckTransferJSON
}

type CheckTransferJSON struct {
	AccountID          apijson.Metadata
	AddressLine1       apijson.Metadata
	AddressLine2       apijson.Metadata
	AddressCity        apijson.Metadata
	AddressState       apijson.Metadata
	AddressZip         apijson.Metadata
	ReturnAddress      apijson.Metadata
	Amount             apijson.Metadata
	CreatedAt          apijson.Metadata
	Currency           apijson.Metadata
	Approval           apijson.Metadata
	Cancellation       apijson.Metadata
	ID                 apijson.Metadata
	MailedAt           apijson.Metadata
	Message            apijson.Metadata
	Note               apijson.Metadata
	RecipientName      apijson.Metadata
	Status             apijson.Metadata
	SubmittedAt        apijson.Metadata
	Submission         apijson.Metadata
	TransactionID      apijson.Metadata
	StopPaymentRequest apijson.Metadata
	Deposit            apijson.Metadata
	ReturnDetails      apijson.Metadata
	Type               apijson.Metadata
	raw                string
	Extras             map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CheckTransfer using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CheckTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CheckTransferReturnAddress struct {
	// The name of the address.
	Name string `json:"name,required,nullable"`
	// The first line of the address.
	Line1 string `json:"line1,required,nullable"`
	// The second line of the address.
	Line2 string `json:"line2,required,nullable"`
	// The city of the address.
	City string `json:"city,required,nullable"`
	// The US state of the address.
	State string `json:"state,required,nullable"`
	// The postal code of the address.
	Zip  string `json:"zip,required,nullable"`
	JSON CheckTransferReturnAddressJSON
}

type CheckTransferReturnAddressJSON struct {
	Name   apijson.Metadata
	Line1  apijson.Metadata
	Line2  apijson.Metadata
	City   apijson.Metadata
	State  apijson.Metadata
	Zip    apijson.Metadata
	raw    string
	Extras map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CheckTransferReturnAddress
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *CheckTransferReturnAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CheckTransferCurrency string

const (
	CheckTransferCurrencyCad CheckTransferCurrency = "CAD"
	CheckTransferCurrencyChf CheckTransferCurrency = "CHF"
	CheckTransferCurrencyEur CheckTransferCurrency = "EUR"
	CheckTransferCurrencyGbp CheckTransferCurrency = "GBP"
	CheckTransferCurrencyJpy CheckTransferCurrency = "JPY"
	CheckTransferCurrencyUsd CheckTransferCurrency = "USD"
)

type CheckTransferApproval struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was approved.
	ApprovedAt time.Time `json:"approved_at,required" format:"date-time"`
	// If the Transfer was approved by a user in the dashboard, the email address of
	// that user.
	ApprovedBy string `json:"approved_by,required,nullable"`
	JSON       CheckTransferApprovalJSON
}

type CheckTransferApprovalJSON struct {
	ApprovedAt apijson.Metadata
	ApprovedBy apijson.Metadata
	raw        string
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CheckTransferApproval using
// the internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CheckTransferApproval) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CheckTransferCancellation struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Transfer was canceled.
	CanceledAt time.Time `json:"canceled_at,required" format:"date-time"`
	// If the Transfer was canceled by a user in the dashboard, the email address of
	// that user.
	CanceledBy string `json:"canceled_by,required,nullable"`
	JSON       CheckTransferCancellationJSON
}

type CheckTransferCancellationJSON struct {
	CanceledAt apijson.Metadata
	CanceledBy apijson.Metadata
	raw        string
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CheckTransferCancellation
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *CheckTransferCancellation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CheckTransferStatus string

const (
	CheckTransferStatusPendingApproval   CheckTransferStatus = "pending_approval"
	CheckTransferStatusPendingSubmission CheckTransferStatus = "pending_submission"
	CheckTransferStatusSubmitted         CheckTransferStatus = "submitted"
	CheckTransferStatusPendingMailing    CheckTransferStatus = "pending_mailing"
	CheckTransferStatusMailed            CheckTransferStatus = "mailed"
	CheckTransferStatusCanceled          CheckTransferStatus = "canceled"
	CheckTransferStatusDeposited         CheckTransferStatus = "deposited"
	CheckTransferStatusStopped           CheckTransferStatus = "stopped"
	CheckTransferStatusReturned          CheckTransferStatus = "returned"
	CheckTransferStatusRejected          CheckTransferStatus = "rejected"
	CheckTransferStatusRequiresAttention CheckTransferStatus = "requires_attention"
)

type CheckTransferSubmission struct {
	// When this check transfer was submitted to our check printer.
	SubmittedAt time.Time `json:"submitted_at,required" format:"date-time"`
	// The identitying number of the check.
	CheckNumber string `json:"check_number,required"`
	JSON        CheckTransferSubmissionJSON
}

type CheckTransferSubmissionJSON struct {
	SubmittedAt apijson.Metadata
	CheckNumber apijson.Metadata
	raw         string
	Extras      map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CheckTransferSubmission using
// the internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CheckTransferSubmission) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CheckTransferStopPaymentRequest struct {
	// The ID of the check transfer that was stopped.
	TransferID string `json:"transfer_id,required"`
	// The transaction ID of the corresponding credit transaction.
	TransactionID string `json:"transaction_id,required"`
	// The time the stop-payment was requested.
	RequestedAt time.Time `json:"requested_at,required" format:"date-time"`
	// A constant representing the object's type. For this resource it will always be
	// `check_transfer_stop_payment_request`.
	Type CheckTransferStopPaymentRequestType `json:"type,required"`
	JSON CheckTransferStopPaymentRequestJSON
}

type CheckTransferStopPaymentRequestJSON struct {
	TransferID    apijson.Metadata
	TransactionID apijson.Metadata
	RequestedAt   apijson.Metadata
	Type          apijson.Metadata
	raw           string
	Extras        map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// CheckTransferStopPaymentRequest using the internal json library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *CheckTransferStopPaymentRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CheckTransferStopPaymentRequestType string

const (
	CheckTransferStopPaymentRequestTypeCheckTransferStopPaymentRequest CheckTransferStopPaymentRequestType = "check_transfer_stop_payment_request"
)

type CheckTransferDeposit struct {
	// When the check was deposited.
	DepositedAt time.Time `json:"deposited_at,required" format:"date-time"`
	// The ID for the File containing the image of the front of the check.
	FrontImageFileID string `json:"front_image_file_id,required,nullable"`
	// The ID for the File containing the image of the rear of the check.
	BackImageFileID string `json:"back_image_file_id,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `check_transfer_deposit`.
	Type CheckTransferDepositType `json:"type,required"`
	JSON CheckTransferDepositJSON
}

type CheckTransferDepositJSON struct {
	DepositedAt      apijson.Metadata
	FrontImageFileID apijson.Metadata
	BackImageFileID  apijson.Metadata
	Type             apijson.Metadata
	raw              string
	Extras           map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CheckTransferDeposit using
// the internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CheckTransferDeposit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CheckTransferDepositType string

const (
	CheckTransferDepositTypeCheckTransferDeposit CheckTransferDepositType = "check_transfer_deposit"
)

type CheckTransferReturnDetails struct {
	// The identifier of the returned Check Transfer.
	TransferID string `json:"transfer_id,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check was returned.
	ReturnedAt time.Time `json:"returned_at,required" format:"date-time"`
	// If available, a document with additional information about the return.
	FileID string `json:"file_id,required,nullable"`
	// The reason why the check was returned.
	Reason CheckTransferReturnDetailsReason `json:"reason,required"`
	// The identifier of the Transaction that was created to credit you for the
	// returned check.
	TransactionID string `json:"transaction_id,required,nullable"`
	JSON          CheckTransferReturnDetailsJSON
}

type CheckTransferReturnDetailsJSON struct {
	TransferID    apijson.Metadata
	ReturnedAt    apijson.Metadata
	FileID        apijson.Metadata
	Reason        apijson.Metadata
	TransactionID apijson.Metadata
	raw           string
	Extras        map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CheckTransferReturnDetails
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *CheckTransferReturnDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CheckTransferReturnDetailsReason string

const (
	CheckTransferReturnDetailsReasonMailDeliveryFailure CheckTransferReturnDetailsReason = "mail_delivery_failure"
	CheckTransferReturnDetailsReasonRefusedByRecipient  CheckTransferReturnDetailsReason = "refused_by_recipient"
)

type CheckTransferType string

const (
	CheckTransferTypeCheckTransfer CheckTransferType = "check_transfer"
)

type CheckTransferNewParams struct {
	// The identifier for the account that will send the transfer.
	AccountID field.Field[string] `json:"account_id,required"`
	// The street address of the check's destination.
	AddressLine1 field.Field[string] `json:"address_line1,required"`
	// The second line of the address of the check's destination.
	AddressLine2 field.Field[string] `json:"address_line2"`
	// The city of the check's destination.
	AddressCity field.Field[string] `json:"address_city,required"`
	// The state of the check's destination.
	AddressState field.Field[string] `json:"address_state,required"`
	// The postal code of the check's destination.
	AddressZip field.Field[string] `json:"address_zip,required"`
	// The return address to be printed on the check. If omitted this will default to
	// the address of the Entity of the Account used to make the Check Transfer.
	ReturnAddress field.Field[CheckTransferNewParamsReturnAddress] `json:"return_address"`
	// The transfer amount in cents.
	Amount field.Field[int64] `json:"amount,required"`
	// The descriptor that will be printed on the memo field on the check.
	Message field.Field[string] `json:"message,required"`
	// The descriptor that will be printed on the letter included with the check.
	Note field.Field[string] `json:"note"`
	// The name that will be printed on the check.
	RecipientName field.Field[string] `json:"recipient_name,required"`
	// Whether the transfer requires explicit approval via the dashboard or API.
	RequireApproval field.Field[bool] `json:"require_approval"`
}

// MarshalJSON serializes CheckTransferNewParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r CheckTransferNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CheckTransferNewParamsReturnAddress struct {
	// The name of the return address.
	Name field.Field[string] `json:"name,required"`
	// The first line of the return address.
	Line1 field.Field[string] `json:"line1,required"`
	// The second line of the return address.
	Line2 field.Field[string] `json:"line2"`
	// The city of the return address.
	City field.Field[string] `json:"city,required"`
	// The US state of the return address.
	State field.Field[string] `json:"state,required"`
	// The postal code of the return address.
	Zip field.Field[string] `json:"zip,required"`
}

type CheckTransferListParams struct {
	// Return the page of entries after this one.
	Cursor field.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit field.Field[int64] `query:"limit"`
	// Filter Check Transfers to those that originated from the specified Account.
	AccountID field.Field[string]                           `query:"account_id"`
	CreatedAt field.Field[CheckTransferListParamsCreatedAt] `query:"created_at"`
}

// URLQuery serializes CheckTransferListParams into a url.Values of the query
// parameters associated with this value
func (r CheckTransferListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

type CheckTransferListParamsCreatedAt struct {
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

// URLQuery serializes CheckTransferListParamsCreatedAt into a url.Values of the
// query parameters associated with this value
func (r CheckTransferListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

type CheckTransferListResponse struct {
	// The contents of the list.
	Data []CheckTransfer `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       CheckTransferListResponseJSON
}

type CheckTransferListResponseJSON struct {
	Data       apijson.Metadata
	NextCursor apijson.Metadata
	raw        string
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CheckTransferListResponse
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *CheckTransferListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

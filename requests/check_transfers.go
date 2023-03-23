package requests

import (
	"fmt"
	"net/url"
	"time"

	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/fields"
)

type CheckTransfer struct {
	// The identifier of the Account from which funds will be transferred.
	AccountID fields.Field[string] `json:"account_id,required"`
	// The street address of the check's destination.
	AddressLine1 fields.Field[string] `json:"address_line1,required"`
	// The second line of the address of the check's destination.
	AddressLine2 fields.Field[string] `json:"address_line2,required,nullable"`
	// The city of the check's destination.
	AddressCity fields.Field[string] `json:"address_city,required"`
	// The state of the check's destination.
	AddressState fields.Field[string] `json:"address_state,required"`
	// The postal code of the check's destination.
	AddressZip fields.Field[string] `json:"address_zip,required"`
	// The return address to be printed on the check.
	ReturnAddress fields.Field[CheckTransferReturnAddress] `json:"return_address,required,nullable"`
	// The transfer amount in USD cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt fields.Field[time.Time] `json:"created_at,required" format:"date-time"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
	// currency.
	Currency fields.Field[CheckTransferCurrency] `json:"currency,required"`
	// The Check transfer's identifier.
	ID fields.Field[string] `json:"id,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check was mailed.
	MailedAt fields.Field[time.Time] `json:"mailed_at,required,nullable" format:"date-time"`
	// The descriptor that will be printed on the memo field on the check.
	Message fields.Field[string] `json:"message,required"`
	// The descriptor that will be printed on the letter included with the check.
	Note fields.Field[string] `json:"note,required,nullable"`
	// The name that will be printed on the check.
	RecipientName fields.Field[string] `json:"recipient_name,required"`
	// The lifecycle status of the transfer.
	Status fields.Field[CheckTransferStatus] `json:"status,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check was submitted.
	SubmittedAt fields.Field[time.Time] `json:"submitted_at,required,nullable" format:"date-time"`
	// After the transfer is submitted, this will contain supplemental details.
	Submission fields.Field[CheckTransferSubmission] `json:"submission,required,nullable"`
	// If the transfer was created from a template, this will be the template's ID.
	TemplateID fields.Field[string] `json:"template_id,required,nullable"`
	// The ID for the transaction caused by the transfer.
	TransactionID fields.Field[string] `json:"transaction_id,required,nullable"`
	// After a stop-payment is requested on the check, this will contain supplemental
	// details.
	StopPaymentRequest fields.Field[CheckTransferStopPaymentRequest] `json:"stop_payment_request,required,nullable"`
	// After a check transfer is deposited, this will contain supplemental details.
	Deposit fields.Field[CheckTransferDeposit] `json:"deposit,required,nullable"`
	// After a check transfer is returned, this will contain supplemental details. A
	// check transfer is returned when the receiver mails a never deposited check back
	// to the bank printed on the check.
	ReturnDetails fields.Field[CheckTransferReturnDetails] `json:"return_details,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `check_transfer`.
	Type fields.Field[CheckTransferType] `json:"type,required"`
}

// MarshalJSON serializes CheckTransfer into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *CheckTransfer) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CheckTransfer) String() (result string) {
	return fmt.Sprintf("&CheckTransfer{AccountID:%s AddressLine1:%s AddressLine2:%s AddressCity:%s AddressState:%s AddressZip:%s ReturnAddress:%s Amount:%s CreatedAt:%s Currency:%s ID:%s MailedAt:%s Message:%s Note:%s RecipientName:%s Status:%s SubmittedAt:%s Submission:%s TemplateID:%s TransactionID:%s StopPaymentRequest:%s Deposit:%s ReturnDetails:%s Type:%s}", r.AccountID, r.AddressLine1, r.AddressLine2, r.AddressCity, r.AddressState, r.AddressZip, r.ReturnAddress, r.Amount, r.CreatedAt, r.Currency, r.ID, r.MailedAt, r.Message, r.Note, r.RecipientName, r.Status, r.SubmittedAt, r.Submission, r.TemplateID, r.TransactionID, r.StopPaymentRequest, r.Deposit, r.ReturnDetails, r.Type)
}

type CheckTransferReturnAddress struct {
	// The name of the address.
	Name fields.Field[string] `json:"name,required,nullable"`
	// The first line of the address.
	Line1 fields.Field[string] `json:"line1,required,nullable"`
	// The second line of the address.
	Line2 fields.Field[string] `json:"line2,required,nullable"`
	// The city of the address.
	City fields.Field[string] `json:"city,required,nullable"`
	// The US state of the address.
	State fields.Field[string] `json:"state,required,nullable"`
	// The postal code of the address.
	Zip fields.Field[string] `json:"zip,required,nullable"`
}

// MarshalJSON serializes CheckTransferReturnAddress into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CheckTransferReturnAddress) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CheckTransferReturnAddress) String() (result string) {
	return fmt.Sprintf("&CheckTransferReturnAddress{Name:%s Line1:%s Line2:%s City:%s State:%s Zip:%s}", r.Name, r.Line1, r.Line2, r.City, r.State, r.Zip)
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

type CheckTransferStatus string

const (
	CheckTransferStatusPendingApproval   CheckTransferStatus = "pending_approval"
	CheckTransferStatusPendingSubmission CheckTransferStatus = "pending_submission"
	CheckTransferStatusSubmitting        CheckTransferStatus = "submitting"
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
	// The identitying number of the check.
	CheckNumber fields.Field[string] `json:"check_number,required"`
}

// MarshalJSON serializes CheckTransferSubmission into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CheckTransferSubmission) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CheckTransferSubmission) String() (result string) {
	return fmt.Sprintf("&CheckTransferSubmission{CheckNumber:%s}", r.CheckNumber)
}

type CheckTransferStopPaymentRequest struct {
	// The ID of the check transfer that was stopped.
	TransferID fields.Field[string] `json:"transfer_id,required"`
	// The transaction ID of the corresponding credit transaction.
	TransactionID fields.Field[string] `json:"transaction_id,required"`
	// The time the stop-payment was requested.
	RequestedAt fields.Field[time.Time] `json:"requested_at,required" format:"date-time"`
	// A constant representing the object's type. For this resource it will always be
	// `check_transfer_stop_payment_request`.
	Type fields.Field[CheckTransferStopPaymentRequestType] `json:"type,required"`
}

// MarshalJSON serializes CheckTransferStopPaymentRequest into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *CheckTransferStopPaymentRequest) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CheckTransferStopPaymentRequest) String() (result string) {
	return fmt.Sprintf("&CheckTransferStopPaymentRequest{TransferID:%s TransactionID:%s RequestedAt:%s Type:%s}", r.TransferID, r.TransactionID, r.RequestedAt, r.Type)
}

type CheckTransferStopPaymentRequestType string

const (
	CheckTransferStopPaymentRequestTypeCheckTransferStopPaymentRequest CheckTransferStopPaymentRequestType = "check_transfer_stop_payment_request"
)

type CheckTransferDeposit struct {
	// The ID for the File containing the image of the front of the check.
	FrontImageFileID fields.Field[string] `json:"front_image_file_id,required,nullable"`
	// The ID for the File containing the image of the rear of the check.
	BackImageFileID fields.Field[string] `json:"back_image_file_id,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `check_transfer_deposit`.
	Type fields.Field[CheckTransferDepositType] `json:"type,required"`
}

// MarshalJSON serializes CheckTransferDeposit into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CheckTransferDeposit) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CheckTransferDeposit) String() (result string) {
	return fmt.Sprintf("&CheckTransferDeposit{FrontImageFileID:%s BackImageFileID:%s Type:%s}", r.FrontImageFileID, r.BackImageFileID, r.Type)
}

type CheckTransferDepositType string

const (
	CheckTransferDepositTypeCheckTransferDeposit CheckTransferDepositType = "check_transfer_deposit"
)

type CheckTransferReturnDetails struct {
	// The identifier of the returned Check Transfer.
	TransferID fields.Field[string] `json:"transfer_id,required"`
	// If available, a document with additional information about the return.
	FileID fields.Field[string] `json:"file_id,required,nullable"`
	// The reason why the check was returned.
	Reason fields.Field[CheckTransferReturnDetailsReason] `json:"reason,required"`
}

// MarshalJSON serializes CheckTransferReturnDetails into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CheckTransferReturnDetails) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CheckTransferReturnDetails) String() (result string) {
	return fmt.Sprintf("&CheckTransferReturnDetails{TransferID:%s FileID:%s Reason:%s}", r.TransferID, r.FileID, r.Reason)
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

type CreateACheckTransferParameters struct {
	// The identifier for the account that will send the transfer.
	AccountID fields.Field[string] `json:"account_id,required"`
	// The street address of the check's destination.
	AddressLine1 fields.Field[string] `json:"address_line1,required"`
	// The second line of the address of the check's destination.
	AddressLine2 fields.Field[string] `json:"address_line2"`
	// The city of the check's destination.
	AddressCity fields.Field[string] `json:"address_city,required"`
	// The state of the check's destination.
	AddressState fields.Field[string] `json:"address_state,required"`
	// The postal code of the check's destination.
	AddressZip fields.Field[string] `json:"address_zip,required"`
	// The return address to be printed on the check. If omitted this will default to
	// the address of the Entity of the Account used to make the Check Transfer.
	ReturnAddress fields.Field[CreateACheckTransferParametersReturnAddress] `json:"return_address"`
	// The transfer amount in cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The descriptor that will be printed on the memo field on the check.
	Message fields.Field[string] `json:"message,required"`
	// The descriptor that will be printed on the letter included with the check.
	Note fields.Field[string] `json:"note"`
	// The name that will be printed on the check.
	RecipientName fields.Field[string] `json:"recipient_name,required"`
	// Whether the transfer requires explicit approval via the dashboard or API.
	RequireApproval fields.Field[bool] `json:"require_approval"`
}

// MarshalJSON serializes CreateACheckTransferParameters into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *CreateACheckTransferParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CreateACheckTransferParameters) String() (result string) {
	return fmt.Sprintf("&CreateACheckTransferParameters{AccountID:%s AddressLine1:%s AddressLine2:%s AddressCity:%s AddressState:%s AddressZip:%s ReturnAddress:%s Amount:%s Message:%s Note:%s RecipientName:%s RequireApproval:%s}", r.AccountID, r.AddressLine1, r.AddressLine2, r.AddressCity, r.AddressState, r.AddressZip, r.ReturnAddress, r.Amount, r.Message, r.Note, r.RecipientName, r.RequireApproval)
}

type CreateACheckTransferParametersReturnAddress struct {
	// The name of the return address.
	Name fields.Field[string] `json:"name,required"`
	// The first line of the return address.
	Line1 fields.Field[string] `json:"line1,required"`
	// The second line of the return address.
	Line2 fields.Field[string] `json:"line2"`
	// The city of the return address.
	City fields.Field[string] `json:"city,required"`
	// The US state of the return address.
	State fields.Field[string] `json:"state,required"`
	// The postal code of the return address.
	Zip fields.Field[string] `json:"zip,required"`
}

func (r CreateACheckTransferParametersReturnAddress) String() (result string) {
	return fmt.Sprintf("&CreateACheckTransferParametersReturnAddress{Name:%s Line1:%s Line2:%s City:%s State:%s Zip:%s}", r.Name, r.Line1, r.Line2, r.City, r.State, r.Zip)
}

type CheckTransferListParams struct {
	// Return the page of entries after this one.
	Cursor fields.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit fields.Field[int64] `query:"limit"`
	// Filter Check Transfers to those that originated from the specified Account.
	AccountID fields.Field[string]                           `query:"account_id"`
	CreatedAt fields.Field[CheckTransferListParamsCreatedAt] `query:"created_at"`
}

// URLQuery serializes CheckTransferListParams into a url.Values of the query
// parameters associated with this value
func (r *CheckTransferListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r CheckTransferListParams) String() (result string) {
	return fmt.Sprintf("&CheckTransferListParams{Cursor:%s Limit:%s AccountID:%s CreatedAt:%s}", r.Cursor, r.Limit, r.AccountID, r.CreatedAt)
}

type CheckTransferListParamsCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After fields.Field[time.Time] `query:"after" format:"date-time"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before fields.Field[time.Time] `query:"before" format:"date-time"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter fields.Field[time.Time] `query:"on_or_after" format:"date-time"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore fields.Field[time.Time] `query:"on_or_before" format:"date-time"`
}

// URLQuery serializes CheckTransferListParamsCreatedAt into a url.Values of the
// query parameters associated with this value
func (r *CheckTransferListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r CheckTransferListParamsCreatedAt) String() (result string) {
	return fmt.Sprintf("&CheckTransferListParamsCreatedAt{After:%s Before:%s OnOrAfter:%s OnOrBefore:%s}", r.After, r.Before, r.OnOrAfter, r.OnOrBefore)
}

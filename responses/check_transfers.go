package responses

import (
	"time"

	pjson "github.com/increase/increase-go/core/json"
)

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
	// If the transfer was created from a template, this will be the template's ID.
	TemplateID string `json:"template_id,required,nullable"`
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
	AccountID          pjson.Metadata
	AddressLine1       pjson.Metadata
	AddressLine2       pjson.Metadata
	AddressCity        pjson.Metadata
	AddressState       pjson.Metadata
	AddressZip         pjson.Metadata
	ReturnAddress      pjson.Metadata
	Amount             pjson.Metadata
	CreatedAt          pjson.Metadata
	Currency           pjson.Metadata
	ID                 pjson.Metadata
	MailedAt           pjson.Metadata
	Message            pjson.Metadata
	Note               pjson.Metadata
	RecipientName      pjson.Metadata
	Status             pjson.Metadata
	SubmittedAt        pjson.Metadata
	Submission         pjson.Metadata
	TemplateID         pjson.Metadata
	TransactionID      pjson.Metadata
	StopPaymentRequest pjson.Metadata
	Deposit            pjson.Metadata
	ReturnDetails      pjson.Metadata
	Type               pjson.Metadata
	Raw                []byte
	Extras             map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CheckTransfer using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CheckTransfer) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	Name   pjson.Metadata
	Line1  pjson.Metadata
	Line2  pjson.Metadata
	City   pjson.Metadata
	State  pjson.Metadata
	Zip    pjson.Metadata
	Raw    []byte
	Extras map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CheckTransferReturnAddress
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *CheckTransferReturnAddress) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	CheckNumber string `json:"check_number,required"`
	JSON        CheckTransferSubmissionJSON
}

type CheckTransferSubmissionJSON struct {
	CheckNumber pjson.Metadata
	Raw         []byte
	Extras      map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CheckTransferSubmission using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CheckTransferSubmission) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	TransferID    pjson.Metadata
	TransactionID pjson.Metadata
	RequestedAt   pjson.Metadata
	Type          pjson.Metadata
	Raw           []byte
	Extras        map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// CheckTransferStopPaymentRequest using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *CheckTransferStopPaymentRequest) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type CheckTransferStopPaymentRequestType string

const (
	CheckTransferStopPaymentRequestTypeCheckTransferStopPaymentRequest CheckTransferStopPaymentRequestType = "check_transfer_stop_payment_request"
)

type CheckTransferDeposit struct {
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
	FrontImageFileID pjson.Metadata
	BackImageFileID  pjson.Metadata
	Type             pjson.Metadata
	Raw              []byte
	Extras           map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CheckTransferDeposit using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CheckTransferDeposit) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type CheckTransferDepositType string

const (
	CheckTransferDepositTypeCheckTransferDeposit CheckTransferDepositType = "check_transfer_deposit"
)

type CheckTransferReturnDetails struct {
	// The identifier of the returned Check Transfer.
	TransferID string `json:"transfer_id,required"`
	// If available, a document with additional information about the return.
	FileID string `json:"file_id,required,nullable"`
	// The reason why the check was returned.
	Reason CheckTransferReturnDetailsReason `json:"reason,required"`
	JSON   CheckTransferReturnDetailsJSON
}

type CheckTransferReturnDetailsJSON struct {
	TransferID pjson.Metadata
	FileID     pjson.Metadata
	Reason     pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CheckTransferReturnDetails
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *CheckTransferReturnDetails) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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

type CheckTransferListResponse struct {
	// The contents of the list.
	Data []CheckTransfer `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       CheckTransferListResponseJSON
}

type CheckTransferListResponseJSON struct {
	Data       pjson.Metadata
	NextCursor pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CheckTransferListResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *CheckTransferListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

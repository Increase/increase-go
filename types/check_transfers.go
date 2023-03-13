package types

import (
	"fmt"
	"net/url"

	"github.com/increase/increase-go/core"
	"github.com/increase/increase-go/core/pjson"
	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/pagination"
)

type CheckTransfer struct {
	// The identifier of the Account from which funds will be transferred.
	AccountID *string `pjson:"account_id"`
	// The street address of the check's destination.
	AddressLine1 *string `pjson:"address_line1"`
	// The second line of the address of the check's destination.
	AddressLine2 *string `pjson:"address_line2"`
	// The city of the check's destination.
	AddressCity *string `pjson:"address_city"`
	// The state of the check's destination.
	AddressState *string `pjson:"address_state"`
	// The postal code of the check's destination.
	AddressZip *string `pjson:"address_zip"`
	// The return address to be printed on the check.
	ReturnAddress *CheckTransferReturnAddress `pjson:"return_address"`
	// The transfer amount in USD cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt *string `pjson:"created_at"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
	// currency.
	Currency *CheckTransferCurrency `pjson:"currency"`
	// The Check transfer's identifier.
	ID *string `pjson:"id"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check was mailed.
	MailedAt *string `pjson:"mailed_at"`
	// The descriptor that will be printed on the memo field on the check.
	Message *string `pjson:"message"`
	// The descriptor that will be printed on the letter included with the check.
	Note *string `pjson:"note"`
	// The name that will be printed on the check.
	RecipientName *string `pjson:"recipient_name"`
	// The lifecycle status of the transfer.
	Status *CheckTransferStatus `pjson:"status"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check was submitted.
	SubmittedAt *string `pjson:"submitted_at"`
	// After the transfer is submitted, this will contain supplemental details.
	Submission *CheckTransferSubmission `pjson:"submission"`
	// If the transfer was created from a template, this will be the template's ID.
	TemplateID *string `pjson:"template_id"`
	// The ID for the transaction caused by the transfer.
	TransactionID *string `pjson:"transaction_id"`
	// After a stop-payment is requested on the check, this will contain supplemental
	// details.
	StopPaymentRequest *CheckTransferStopPaymentRequest `pjson:"stop_payment_request"`
	// After a check transfer is deposited, this will contain supplemental details.
	Deposit *CheckTransferDeposit `pjson:"deposit"`
	// A constant representing the object's type. For this resource it will always be
	// `check_transfer`.
	Type       *CheckTransferType     `pjson:"type"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into CheckTransfer using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CheckTransfer) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CheckTransfer into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *CheckTransfer) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The identifier of the Account from which funds will be transferred.
func (r CheckTransfer) GetAccountID() (AccountID string) {
	if r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

// The street address of the check's destination.
func (r CheckTransfer) GetAddressLine1() (AddressLine1 string) {
	if r.AddressLine1 != nil {
		AddressLine1 = *r.AddressLine1
	}
	return
}

// The second line of the address of the check's destination.
func (r CheckTransfer) GetAddressLine2() (AddressLine2 string) {
	if r.AddressLine2 != nil {
		AddressLine2 = *r.AddressLine2
	}
	return
}

// The city of the check's destination.
func (r CheckTransfer) GetAddressCity() (AddressCity string) {
	if r.AddressCity != nil {
		AddressCity = *r.AddressCity
	}
	return
}

// The state of the check's destination.
func (r CheckTransfer) GetAddressState() (AddressState string) {
	if r.AddressState != nil {
		AddressState = *r.AddressState
	}
	return
}

// The postal code of the check's destination.
func (r CheckTransfer) GetAddressZip() (AddressZip string) {
	if r.AddressZip != nil {
		AddressZip = *r.AddressZip
	}
	return
}

// The return address to be printed on the check.
func (r CheckTransfer) GetReturnAddress() (ReturnAddress CheckTransferReturnAddress) {
	if r.ReturnAddress != nil {
		ReturnAddress = *r.ReturnAddress
	}
	return
}

// The transfer amount in USD cents.
func (r CheckTransfer) GetAmount() (Amount int64) {
	if r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the transfer was created.
func (r CheckTransfer) GetCreatedAt() (CreatedAt string) {
	if r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
// currency.
func (r CheckTransfer) GetCurrency() (Currency CheckTransferCurrency) {
	if r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The Check transfer's identifier.
func (r CheckTransfer) GetID() (ID string) {
	if r.ID != nil {
		ID = *r.ID
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the check was mailed.
func (r CheckTransfer) GetMailedAt() (MailedAt string) {
	if r.MailedAt != nil {
		MailedAt = *r.MailedAt
	}
	return
}

// The descriptor that will be printed on the memo field on the check.
func (r CheckTransfer) GetMessage() (Message string) {
	if r.Message != nil {
		Message = *r.Message
	}
	return
}

// The descriptor that will be printed on the letter included with the check.
func (r CheckTransfer) GetNote() (Note string) {
	if r.Note != nil {
		Note = *r.Note
	}
	return
}

// The name that will be printed on the check.
func (r CheckTransfer) GetRecipientName() (RecipientName string) {
	if r.RecipientName != nil {
		RecipientName = *r.RecipientName
	}
	return
}

// The lifecycle status of the transfer.
func (r CheckTransfer) GetStatus() (Status CheckTransferStatus) {
	if r.Status != nil {
		Status = *r.Status
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the check was submitted.
func (r CheckTransfer) GetSubmittedAt() (SubmittedAt string) {
	if r.SubmittedAt != nil {
		SubmittedAt = *r.SubmittedAt
	}
	return
}

// After the transfer is submitted, this will contain supplemental details.
func (r CheckTransfer) GetSubmission() (Submission CheckTransferSubmission) {
	if r.Submission != nil {
		Submission = *r.Submission
	}
	return
}

// If the transfer was created from a template, this will be the template's ID.
func (r CheckTransfer) GetTemplateID() (TemplateID string) {
	if r.TemplateID != nil {
		TemplateID = *r.TemplateID
	}
	return
}

// The ID for the transaction caused by the transfer.
func (r CheckTransfer) GetTransactionID() (TransactionID string) {
	if r.TransactionID != nil {
		TransactionID = *r.TransactionID
	}
	return
}

// After a stop-payment is requested on the check, this will contain supplemental
// details.
func (r CheckTransfer) GetStopPaymentRequest() (StopPaymentRequest CheckTransferStopPaymentRequest) {
	if r.StopPaymentRequest != nil {
		StopPaymentRequest = *r.StopPaymentRequest
	}
	return
}

// After a check transfer is deposited, this will contain supplemental details.
func (r CheckTransfer) GetDeposit() (Deposit CheckTransferDeposit) {
	if r.Deposit != nil {
		Deposit = *r.Deposit
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `check_transfer`.
func (r CheckTransfer) GetType() (Type CheckTransferType) {
	if r.Type != nil {
		Type = *r.Type
	}
	return
}

func (r CheckTransfer) String() (result string) {
	return fmt.Sprintf("&CheckTransfer{AccountID:%s AddressLine1:%s AddressLine2:%s AddressCity:%s AddressState:%s AddressZip:%s ReturnAddress:%s Amount:%s CreatedAt:%s Currency:%s ID:%s MailedAt:%s Message:%s Note:%s RecipientName:%s Status:%s SubmittedAt:%s Submission:%s TemplateID:%s TransactionID:%s StopPaymentRequest:%s Deposit:%s Type:%s}", core.FmtP(r.AccountID), core.FmtP(r.AddressLine1), core.FmtP(r.AddressLine2), core.FmtP(r.AddressCity), core.FmtP(r.AddressState), core.FmtP(r.AddressZip), r.ReturnAddress, core.FmtP(r.Amount), core.FmtP(r.CreatedAt), core.FmtP(r.Currency), core.FmtP(r.ID), core.FmtP(r.MailedAt), core.FmtP(r.Message), core.FmtP(r.Note), core.FmtP(r.RecipientName), core.FmtP(r.Status), core.FmtP(r.SubmittedAt), r.Submission, core.FmtP(r.TemplateID), core.FmtP(r.TransactionID), r.StopPaymentRequest, r.Deposit, core.FmtP(r.Type))
}

type CheckTransferReturnAddress struct {
	// The name of the address.
	Name *string `pjson:"name"`
	// The first line of the address.
	Line1 *string `pjson:"line1"`
	// The second line of the address.
	Line2 *string `pjson:"line2"`
	// The city of the address.
	City *string `pjson:"city"`
	// The US state of the address.
	State *string `pjson:"state"`
	// The postal code of the address.
	Zip        *string                `pjson:"zip"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into CheckTransferReturnAddress
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *CheckTransferReturnAddress) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CheckTransferReturnAddress into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CheckTransferReturnAddress) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The name of the address.
func (r CheckTransferReturnAddress) GetName() (Name string) {
	if r.Name != nil {
		Name = *r.Name
	}
	return
}

// The first line of the address.
func (r CheckTransferReturnAddress) GetLine1() (Line1 string) {
	if r.Line1 != nil {
		Line1 = *r.Line1
	}
	return
}

// The second line of the address.
func (r CheckTransferReturnAddress) GetLine2() (Line2 string) {
	if r.Line2 != nil {
		Line2 = *r.Line2
	}
	return
}

// The city of the address.
func (r CheckTransferReturnAddress) GetCity() (City string) {
	if r.City != nil {
		City = *r.City
	}
	return
}

// The US state of the address.
func (r CheckTransferReturnAddress) GetState() (State string) {
	if r.State != nil {
		State = *r.State
	}
	return
}

// The postal code of the address.
func (r CheckTransferReturnAddress) GetZip() (Zip string) {
	if r.Zip != nil {
		Zip = *r.Zip
	}
	return
}

func (r CheckTransferReturnAddress) String() (result string) {
	return fmt.Sprintf("&CheckTransferReturnAddress{Name:%s Line1:%s Line2:%s City:%s State:%s Zip:%s}", core.FmtP(r.Name), core.FmtP(r.Line1), core.FmtP(r.Line2), core.FmtP(r.City), core.FmtP(r.State), core.FmtP(r.Zip))
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
	CheckNumber *string                `pjson:"check_number"`
	jsonFields  map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into CheckTransferSubmission using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CheckTransferSubmission) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CheckTransferSubmission into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CheckTransferSubmission) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The identitying number of the check.
func (r CheckTransferSubmission) GetCheckNumber() (CheckNumber string) {
	if r.CheckNumber != nil {
		CheckNumber = *r.CheckNumber
	}
	return
}

func (r CheckTransferSubmission) String() (result string) {
	return fmt.Sprintf("&CheckTransferSubmission{CheckNumber:%s}", core.FmtP(r.CheckNumber))
}

type CheckTransferStopPaymentRequest struct {
	// The ID of the check transfer that was stopped.
	TransferID *string `pjson:"transfer_id"`
	// The transaction ID of the corresponding credit transaction.
	TransactionID *string `pjson:"transaction_id"`
	// The time the stop-payment was requested.
	RequestedAt *string `pjson:"requested_at"`
	// A constant representing the object's type. For this resource it will always be
	// `check_transfer_stop_payment_request`.
	Type       *CheckTransferStopPaymentRequestType `pjson:"type"`
	jsonFields map[string]interface{}               `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CheckTransferStopPaymentRequest using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *CheckTransferStopPaymentRequest) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CheckTransferStopPaymentRequest into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *CheckTransferStopPaymentRequest) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The ID of the check transfer that was stopped.
func (r CheckTransferStopPaymentRequest) GetTransferID() (TransferID string) {
	if r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

// The transaction ID of the corresponding credit transaction.
func (r CheckTransferStopPaymentRequest) GetTransactionID() (TransactionID string) {
	if r.TransactionID != nil {
		TransactionID = *r.TransactionID
	}
	return
}

// The time the stop-payment was requested.
func (r CheckTransferStopPaymentRequest) GetRequestedAt() (RequestedAt string) {
	if r.RequestedAt != nil {
		RequestedAt = *r.RequestedAt
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `check_transfer_stop_payment_request`.
func (r CheckTransferStopPaymentRequest) GetType() (Type CheckTransferStopPaymentRequestType) {
	if r.Type != nil {
		Type = *r.Type
	}
	return
}

func (r CheckTransferStopPaymentRequest) String() (result string) {
	return fmt.Sprintf("&CheckTransferStopPaymentRequest{TransferID:%s TransactionID:%s RequestedAt:%s Type:%s}", core.FmtP(r.TransferID), core.FmtP(r.TransactionID), core.FmtP(r.RequestedAt), core.FmtP(r.Type))
}

type CheckTransferStopPaymentRequestType string

const (
	CheckTransferStopPaymentRequestTypeCheckTransferStopPaymentRequest CheckTransferStopPaymentRequestType = "check_transfer_stop_payment_request"
)

type CheckTransferDeposit struct {
	// The ID for the File containing the image of the front of the check.
	FrontImageFileID *string `pjson:"front_image_file_id"`
	// The ID for the File containing the image of the rear of the check.
	BackImageFileID *string `pjson:"back_image_file_id"`
	// A constant representing the object's type. For this resource it will always be
	// `check_transfer_deposit`.
	Type       *CheckTransferDepositType `pjson:"type"`
	jsonFields map[string]interface{}    `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into CheckTransferDeposit using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CheckTransferDeposit) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CheckTransferDeposit into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CheckTransferDeposit) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The ID for the File containing the image of the front of the check.
func (r CheckTransferDeposit) GetFrontImageFileID() (FrontImageFileID string) {
	if r.FrontImageFileID != nil {
		FrontImageFileID = *r.FrontImageFileID
	}
	return
}

// The ID for the File containing the image of the rear of the check.
func (r CheckTransferDeposit) GetBackImageFileID() (BackImageFileID string) {
	if r.BackImageFileID != nil {
		BackImageFileID = *r.BackImageFileID
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `check_transfer_deposit`.
func (r CheckTransferDeposit) GetType() (Type CheckTransferDepositType) {
	if r.Type != nil {
		Type = *r.Type
	}
	return
}

func (r CheckTransferDeposit) String() (result string) {
	return fmt.Sprintf("&CheckTransferDeposit{FrontImageFileID:%s BackImageFileID:%s Type:%s}", core.FmtP(r.FrontImageFileID), core.FmtP(r.BackImageFileID), core.FmtP(r.Type))
}

type CheckTransferDepositType string

const (
	CheckTransferDepositTypeCheckTransferDeposit CheckTransferDepositType = "check_transfer_deposit"
)

type CheckTransferType string

const (
	CheckTransferTypeCheckTransfer CheckTransferType = "check_transfer"
)

type CreateACheckTransferParameters struct {
	// The identifier for the account that will send the transfer.
	AccountID *string `pjson:"account_id"`
	// The street address of the check's destination.
	AddressLine1 *string `pjson:"address_line1"`
	// The second line of the address of the check's destination.
	AddressLine2 *string `pjson:"address_line2"`
	// The city of the check's destination.
	AddressCity *string `pjson:"address_city"`
	// The state of the check's destination.
	AddressState *string `pjson:"address_state"`
	// The postal code of the check's destination.
	AddressZip *string `pjson:"address_zip"`
	// The return address to be printed on the check. If omitted this will default to
	// the address of the Entity of the Account used to make the Check Transfer.
	ReturnAddress *CreateACheckTransferParametersReturnAddress `pjson:"return_address"`
	// The transfer amount in cents.
	Amount *int64 `pjson:"amount"`
	// The descriptor that will be printed on the memo field on the check.
	Message *string `pjson:"message"`
	// The descriptor that will be printed on the letter included with the check.
	Note *string `pjson:"note"`
	// The name that will be printed on the check.
	RecipientName *string `pjson:"recipient_name"`
	// Whether the transfer requires explicit approval via the dashboard or API.
	RequireApproval *bool                  `pjson:"require_approval"`
	jsonFields      map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateACheckTransferParameters using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *CreateACheckTransferParameters) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CreateACheckTransferParameters into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *CreateACheckTransferParameters) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The identifier for the account that will send the transfer.
func (r CreateACheckTransferParameters) GetAccountID() (AccountID string) {
	if r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

// The street address of the check's destination.
func (r CreateACheckTransferParameters) GetAddressLine1() (AddressLine1 string) {
	if r.AddressLine1 != nil {
		AddressLine1 = *r.AddressLine1
	}
	return
}

// The second line of the address of the check's destination.
func (r CreateACheckTransferParameters) GetAddressLine2() (AddressLine2 string) {
	if r.AddressLine2 != nil {
		AddressLine2 = *r.AddressLine2
	}
	return
}

// The city of the check's destination.
func (r CreateACheckTransferParameters) GetAddressCity() (AddressCity string) {
	if r.AddressCity != nil {
		AddressCity = *r.AddressCity
	}
	return
}

// The state of the check's destination.
func (r CreateACheckTransferParameters) GetAddressState() (AddressState string) {
	if r.AddressState != nil {
		AddressState = *r.AddressState
	}
	return
}

// The postal code of the check's destination.
func (r CreateACheckTransferParameters) GetAddressZip() (AddressZip string) {
	if r.AddressZip != nil {
		AddressZip = *r.AddressZip
	}
	return
}

// The return address to be printed on the check. If omitted this will default to
// the address of the Entity of the Account used to make the Check Transfer.
func (r CreateACheckTransferParameters) GetReturnAddress() (ReturnAddress CreateACheckTransferParametersReturnAddress) {
	if r.ReturnAddress != nil {
		ReturnAddress = *r.ReturnAddress
	}
	return
}

// The transfer amount in cents.
func (r CreateACheckTransferParameters) GetAmount() (Amount int64) {
	if r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The descriptor that will be printed on the memo field on the check.
func (r CreateACheckTransferParameters) GetMessage() (Message string) {
	if r.Message != nil {
		Message = *r.Message
	}
	return
}

// The descriptor that will be printed on the letter included with the check.
func (r CreateACheckTransferParameters) GetNote() (Note string) {
	if r.Note != nil {
		Note = *r.Note
	}
	return
}

// The name that will be printed on the check.
func (r CreateACheckTransferParameters) GetRecipientName() (RecipientName string) {
	if r.RecipientName != nil {
		RecipientName = *r.RecipientName
	}
	return
}

// Whether the transfer requires explicit approval via the dashboard or API.
func (r CreateACheckTransferParameters) GetRequireApproval() (RequireApproval bool) {
	if r.RequireApproval != nil {
		RequireApproval = *r.RequireApproval
	}
	return
}

func (r CreateACheckTransferParameters) String() (result string) {
	return fmt.Sprintf("&CreateACheckTransferParameters{AccountID:%s AddressLine1:%s AddressLine2:%s AddressCity:%s AddressState:%s AddressZip:%s ReturnAddress:%s Amount:%s Message:%s Note:%s RecipientName:%s RequireApproval:%s}", core.FmtP(r.AccountID), core.FmtP(r.AddressLine1), core.FmtP(r.AddressLine2), core.FmtP(r.AddressCity), core.FmtP(r.AddressState), core.FmtP(r.AddressZip), r.ReturnAddress, core.FmtP(r.Amount), core.FmtP(r.Message), core.FmtP(r.Note), core.FmtP(r.RecipientName), core.FmtP(r.RequireApproval))
}

type CreateACheckTransferParametersReturnAddress struct {
	// The name of the return address.
	Name *string `pjson:"name"`
	// The first line of the return address.
	Line1 *string `pjson:"line1"`
	// The second line of the return address.
	Line2 *string `pjson:"line2"`
	// The city of the return address.
	City *string `pjson:"city"`
	// The US state of the return address.
	State *string `pjson:"state"`
	// The postal code of the return address.
	Zip        *string                `pjson:"zip"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateACheckTransferParametersReturnAddress using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *CreateACheckTransferParametersReturnAddress) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CreateACheckTransferParametersReturnAddress into an array
// of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CreateACheckTransferParametersReturnAddress) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The name of the return address.
func (r CreateACheckTransferParametersReturnAddress) GetName() (Name string) {
	if r.Name != nil {
		Name = *r.Name
	}
	return
}

// The first line of the return address.
func (r CreateACheckTransferParametersReturnAddress) GetLine1() (Line1 string) {
	if r.Line1 != nil {
		Line1 = *r.Line1
	}
	return
}

// The second line of the return address.
func (r CreateACheckTransferParametersReturnAddress) GetLine2() (Line2 string) {
	if r.Line2 != nil {
		Line2 = *r.Line2
	}
	return
}

// The city of the return address.
func (r CreateACheckTransferParametersReturnAddress) GetCity() (City string) {
	if r.City != nil {
		City = *r.City
	}
	return
}

// The US state of the return address.
func (r CreateACheckTransferParametersReturnAddress) GetState() (State string) {
	if r.State != nil {
		State = *r.State
	}
	return
}

// The postal code of the return address.
func (r CreateACheckTransferParametersReturnAddress) GetZip() (Zip string) {
	if r.Zip != nil {
		Zip = *r.Zip
	}
	return
}

func (r CreateACheckTransferParametersReturnAddress) String() (result string) {
	return fmt.Sprintf("&CreateACheckTransferParametersReturnAddress{Name:%s Line1:%s Line2:%s City:%s State:%s Zip:%s}", core.FmtP(r.Name), core.FmtP(r.Line1), core.FmtP(r.Line2), core.FmtP(r.City), core.FmtP(r.State), core.FmtP(r.Zip))
}

type CheckTransferListParams struct {
	// Return the page of entries after this one.
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int64 `query:"limit"`
	// Filter Check Transfers to those that originated from the specified Account.
	AccountID  *string                            `query:"account_id"`
	CreatedAt  *CheckTransfersListParamsCreatedAt `query:"created_at"`
	jsonFields map[string]interface{}             `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into CheckTransferListParams using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CheckTransferListParams) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CheckTransferListParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CheckTransferListParams) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes CheckTransferListParams into a url.Values of the query
// parameters associated with this value
func (r *CheckTransferListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// Return the page of entries after this one.
func (r CheckTransferListParams) GetCursor() (Cursor string) {
	if r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r CheckTransferListParams) GetLimit() (Limit int64) {
	if r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// Filter Check Transfers to those that originated from the specified Account.
func (r CheckTransferListParams) GetAccountID() (AccountID string) {
	if r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

func (r CheckTransferListParams) GetCreatedAt() (CreatedAt CheckTransfersListParamsCreatedAt) {
	if r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

func (r CheckTransferListParams) String() (result string) {
	return fmt.Sprintf("&CheckTransferListParams{Cursor:%s Limit:%s AccountID:%s CreatedAt:%s}", core.FmtP(r.Cursor), core.FmtP(r.Limit), core.FmtP(r.AccountID), r.CreatedAt)
}

type CheckTransfersListParamsCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After *string `pjson:"after"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before *string `pjson:"before"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter *string `pjson:"on_or_after"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore *string                `pjson:"on_or_before"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CheckTransfersListParamsCreatedAt using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *CheckTransfersListParamsCreatedAt) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CheckTransfersListParamsCreatedAt into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *CheckTransfersListParamsCreatedAt) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes CheckTransfersListParamsCreatedAt into a url.Values of the
// query parameters associated with this value
func (r *CheckTransfersListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r CheckTransfersListParamsCreatedAt) GetAfter() (After string) {
	if r.After != nil {
		After = *r.After
	}
	return
}

// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r CheckTransfersListParamsCreatedAt) GetBefore() (Before string) {
	if r.Before != nil {
		Before = *r.Before
	}
	return
}

// Return results on or after this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r CheckTransfersListParamsCreatedAt) GetOnOrAfter() (OnOrAfter string) {
	if r.OnOrAfter != nil {
		OnOrAfter = *r.OnOrAfter
	}
	return
}

// Return results on or before this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r CheckTransfersListParamsCreatedAt) GetOnOrBefore() (OnOrBefore string) {
	if r.OnOrBefore != nil {
		OnOrBefore = *r.OnOrBefore
	}
	return
}

func (r CheckTransfersListParamsCreatedAt) String() (result string) {
	return fmt.Sprintf("&CheckTransfersListParamsCreatedAt{After:%s Before:%s OnOrAfter:%s OnOrBefore:%s}", core.FmtP(r.After), core.FmtP(r.Before), core.FmtP(r.OnOrAfter), core.FmtP(r.OnOrBefore))
}

type CheckTransferList struct {
	// The contents of the list.
	Data *[]CheckTransfer `pjson:"data"`
	// A pointer to a place in the list.
	NextCursor *string                `pjson:"next_cursor"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into CheckTransferList using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CheckTransferList) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CheckTransferList into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *CheckTransferList) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes CheckTransferList into a url.Values of the query parameters
// associated with this value
func (r *CheckTransferList) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// The contents of the list.
func (r CheckTransferList) GetData() (Data []CheckTransfer) {
	if r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r CheckTransferList) GetNextCursor() (NextCursor string) {
	if r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

func (r CheckTransferList) String() (result string) {
	return fmt.Sprintf("&CheckTransferList{Data:%s NextCursor:%s}", core.Fmt(r.Data), core.FmtP(r.NextCursor))
}

type CheckTransfersPage struct {
	*pagination.Page[CheckTransfer]
}

func (r *CheckTransfersPage) CheckTransfer() *CheckTransfer {
	return r.Current()
}

func (r *CheckTransfersPage) NextPage() (*CheckTransfersPage, error) {
	if page, err := r.Page.NextPage(); err != nil {
		return nil, err
	} else {
		return &CheckTransfersPage{page}, nil
	}
}

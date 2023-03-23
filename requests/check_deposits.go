package requests

import (
	"fmt"
	"net/url"
	"time"

	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/fields"
)

type CheckDeposit struct {
	// The deposit's identifier.
	ID fields.Field[string] `json:"id,required"`
	// The deposited amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt fields.Field[time.Time] `json:"created_at,required" format:"date-time"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the deposit.
	Currency fields.Field[CheckDepositCurrency] `json:"currency,required"`
	// The status of the Check Deposit.
	Status fields.Field[CheckDepositStatus] `json:"status,required"`
	// The Account the check was deposited into.
	AccountID fields.Field[string] `json:"account_id,required"`
	// The ID for the File containing the image of the front of the check.
	FrontImageFileID fields.Field[string] `json:"front_image_file_id,required"`
	// The ID for the File containing the image of the back of the check.
	BackImageFileID fields.Field[string] `json:"back_image_file_id,required,nullable"`
	// The ID for the Transaction created by the deposit.
	TransactionID fields.Field[string] `json:"transaction_id,required,nullable"`
	// If your deposit is successfully parsed and accepted by Increase, this will
	// contain details of the parsed check.
	DepositAcceptance fields.Field[CheckDepositDepositAcceptance] `json:"deposit_acceptance,required,nullable"`
	// If your deposit is rejected by Increase, this will contain details as to why it
	// was rejected.
	DepositRejection fields.Field[CheckDepositDepositRejection] `json:"deposit_rejection,required,nullable"`
	// If your deposit is returned, this will contain details as to why it was
	// returned.
	DepositReturn fields.Field[CheckDepositDepositReturn] `json:"deposit_return,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `check_deposit`.
	Type fields.Field[CheckDepositType] `json:"type,required"`
}

// MarshalJSON serializes CheckDeposit into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *CheckDeposit) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CheckDeposit) String() (result string) {
	return fmt.Sprintf("&CheckDeposit{ID:%s Amount:%s CreatedAt:%s Currency:%s Status:%s AccountID:%s FrontImageFileID:%s BackImageFileID:%s TransactionID:%s DepositAcceptance:%s DepositRejection:%s DepositReturn:%s Type:%s}", r.ID, r.Amount, r.CreatedAt, r.Currency, r.Status, r.AccountID, r.FrontImageFileID, r.BackImageFileID, r.TransactionID, r.DepositAcceptance, r.DepositRejection, r.DepositReturn, r.Type)
}

type CheckDepositCurrency string

const (
	CheckDepositCurrencyCad CheckDepositCurrency = "CAD"
	CheckDepositCurrencyChf CheckDepositCurrency = "CHF"
	CheckDepositCurrencyEur CheckDepositCurrency = "EUR"
	CheckDepositCurrencyGbp CheckDepositCurrency = "GBP"
	CheckDepositCurrencyJpy CheckDepositCurrency = "JPY"
	CheckDepositCurrencyUsd CheckDepositCurrency = "USD"
)

type CheckDepositStatus string

const (
	CheckDepositStatusPending   CheckDepositStatus = "pending"
	CheckDepositStatusSubmitted CheckDepositStatus = "submitted"
	CheckDepositStatusRejected  CheckDepositStatus = "rejected"
	CheckDepositStatusReturned  CheckDepositStatus = "returned"
)

type CheckDepositDepositAcceptance struct {
	// The amount to be deposited in the minor unit of the transaction's currency. For
	// dollars, for example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency fields.Field[CheckDepositDepositAcceptanceCurrency] `json:"currency,required"`
	// The account number printed on the check.
	AccountNumber fields.Field[string] `json:"account_number,required"`
	// The routing number printed on the check.
	RoutingNumber fields.Field[string] `json:"routing_number,required"`
	// An additional line of metadata printed on the check. This typically includes the
	// check number for business checks.
	AuxiliaryOnUs fields.Field[string] `json:"auxiliary_on_us,required,nullable"`
	// The check serial number, if present, for consumer checks. For business checks,
	// the serial number is usually in the `auxiliary_on_us` field.
	SerialNumber fields.Field[string] `json:"serial_number,required,nullable"`
	// The ID of the Check Deposit that was accepted.
	CheckDepositID fields.Field[string] `json:"check_deposit_id,required"`
}

// MarshalJSON serializes CheckDepositDepositAcceptance into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *CheckDepositDepositAcceptance) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CheckDepositDepositAcceptance) String() (result string) {
	return fmt.Sprintf("&CheckDepositDepositAcceptance{Amount:%s Currency:%s AccountNumber:%s RoutingNumber:%s AuxiliaryOnUs:%s SerialNumber:%s CheckDepositID:%s}", r.Amount, r.Currency, r.AccountNumber, r.RoutingNumber, r.AuxiliaryOnUs, r.SerialNumber, r.CheckDepositID)
}

type CheckDepositDepositAcceptanceCurrency string

const (
	CheckDepositDepositAcceptanceCurrencyCad CheckDepositDepositAcceptanceCurrency = "CAD"
	CheckDepositDepositAcceptanceCurrencyChf CheckDepositDepositAcceptanceCurrency = "CHF"
	CheckDepositDepositAcceptanceCurrencyEur CheckDepositDepositAcceptanceCurrency = "EUR"
	CheckDepositDepositAcceptanceCurrencyGbp CheckDepositDepositAcceptanceCurrency = "GBP"
	CheckDepositDepositAcceptanceCurrencyJpy CheckDepositDepositAcceptanceCurrency = "JPY"
	CheckDepositDepositAcceptanceCurrencyUsd CheckDepositDepositAcceptanceCurrency = "USD"
)

type CheckDepositDepositRejection struct {
	// The rejected amount in the minor unit of check's currency. For dollars, for
	// example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
	// currency.
	Currency fields.Field[CheckDepositDepositRejectionCurrency] `json:"currency,required"`
	// Why the check deposit was rejected.
	Reason fields.Field[CheckDepositDepositRejectionReason] `json:"reason,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check deposit was rejected.
	RejectedAt fields.Field[time.Time] `json:"rejected_at,required" format:"date-time"`
}

// MarshalJSON serializes CheckDepositDepositRejection into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CheckDepositDepositRejection) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CheckDepositDepositRejection) String() (result string) {
	return fmt.Sprintf("&CheckDepositDepositRejection{Amount:%s Currency:%s Reason:%s RejectedAt:%s}", r.Amount, r.Currency, r.Reason, r.RejectedAt)
}

type CheckDepositDepositRejectionCurrency string

const (
	CheckDepositDepositRejectionCurrencyCad CheckDepositDepositRejectionCurrency = "CAD"
	CheckDepositDepositRejectionCurrencyChf CheckDepositDepositRejectionCurrency = "CHF"
	CheckDepositDepositRejectionCurrencyEur CheckDepositDepositRejectionCurrency = "EUR"
	CheckDepositDepositRejectionCurrencyGbp CheckDepositDepositRejectionCurrency = "GBP"
	CheckDepositDepositRejectionCurrencyJpy CheckDepositDepositRejectionCurrency = "JPY"
	CheckDepositDepositRejectionCurrencyUsd CheckDepositDepositRejectionCurrency = "USD"
)

type CheckDepositDepositRejectionReason string

const (
	CheckDepositDepositRejectionReasonIncompleteImage             CheckDepositDepositRejectionReason = "incomplete_image"
	CheckDepositDepositRejectionReasonDuplicate                   CheckDepositDepositRejectionReason = "duplicate"
	CheckDepositDepositRejectionReasonPoorImageQuality            CheckDepositDepositRejectionReason = "poor_image_quality"
	CheckDepositDepositRejectionReasonIncorrectAmount             CheckDepositDepositRejectionReason = "incorrect_amount"
	CheckDepositDepositRejectionReasonIncorrectRecipient          CheckDepositDepositRejectionReason = "incorrect_recipient"
	CheckDepositDepositRejectionReasonNotEligibleForMobileDeposit CheckDepositDepositRejectionReason = "not_eligible_for_mobile_deposit"
	CheckDepositDepositRejectionReasonUnknown                     CheckDepositDepositRejectionReason = "unknown"
)

type CheckDepositDepositReturn struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check deposit was returned.
	ReturnedAt fields.Field[time.Time] `json:"returned_at,required" format:"date-time"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency fields.Field[CheckDepositDepositReturnCurrency] `json:"currency,required"`
	// The identifier of the Check Deposit that was returned.
	CheckDepositID fields.Field[string] `json:"check_deposit_id,required"`
	// The identifier of the transaction that reversed the original check deposit
	// transaction.
	TransactionID fields.Field[string]                                `json:"transaction_id,required"`
	ReturnReason  fields.Field[CheckDepositDepositReturnReturnReason] `json:"return_reason,required"`
}

// MarshalJSON serializes CheckDepositDepositReturn into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CheckDepositDepositReturn) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CheckDepositDepositReturn) String() (result string) {
	return fmt.Sprintf("&CheckDepositDepositReturn{Amount:%s ReturnedAt:%s Currency:%s CheckDepositID:%s TransactionID:%s ReturnReason:%s}", r.Amount, r.ReturnedAt, r.Currency, r.CheckDepositID, r.TransactionID, r.ReturnReason)
}

type CheckDepositDepositReturnCurrency string

const (
	CheckDepositDepositReturnCurrencyCad CheckDepositDepositReturnCurrency = "CAD"
	CheckDepositDepositReturnCurrencyChf CheckDepositDepositReturnCurrency = "CHF"
	CheckDepositDepositReturnCurrencyEur CheckDepositDepositReturnCurrency = "EUR"
	CheckDepositDepositReturnCurrencyGbp CheckDepositDepositReturnCurrency = "GBP"
	CheckDepositDepositReturnCurrencyJpy CheckDepositDepositReturnCurrency = "JPY"
	CheckDepositDepositReturnCurrencyUsd CheckDepositDepositReturnCurrency = "USD"
)

type CheckDepositDepositReturnReturnReason string

const (
	CheckDepositDepositReturnReturnReasonACHConversionNotSupported CheckDepositDepositReturnReturnReason = "ach_conversion_not_supported"
	CheckDepositDepositReturnReturnReasonClosedAccount             CheckDepositDepositReturnReturnReason = "closed_account"
	CheckDepositDepositReturnReturnReasonDuplicateSubmission       CheckDepositDepositReturnReturnReason = "duplicate_submission"
	CheckDepositDepositReturnReturnReasonInsufficientFunds         CheckDepositDepositReturnReturnReason = "insufficient_funds"
	CheckDepositDepositReturnReturnReasonNoAccount                 CheckDepositDepositReturnReturnReason = "no_account"
	CheckDepositDepositReturnReturnReasonNotAuthorized             CheckDepositDepositReturnReturnReason = "not_authorized"
	CheckDepositDepositReturnReturnReasonStaleDated                CheckDepositDepositReturnReturnReason = "stale_dated"
	CheckDepositDepositReturnReturnReasonStopPayment               CheckDepositDepositReturnReturnReason = "stop_payment"
	CheckDepositDepositReturnReturnReasonUnknownReason             CheckDepositDepositReturnReturnReason = "unknown_reason"
	CheckDepositDepositReturnReturnReasonUnmatchedDetails          CheckDepositDepositReturnReturnReason = "unmatched_details"
	CheckDepositDepositReturnReturnReasonUnreadableImage           CheckDepositDepositReturnReturnReason = "unreadable_image"
)

type CheckDepositType string

const (
	CheckDepositTypeCheckDeposit CheckDepositType = "check_deposit"
)

type CreateACheckDepositParameters struct {
	// The identifier for the Account to deposit the check in.
	AccountID fields.Field[string] `json:"account_id,required"`
	// The deposit amount in the minor unit of the account currency. For dollars, for
	// example, this is cents.
	Amount fields.Field[int64] `json:"amount,required"`
	// The currency to use for the deposit.
	Currency fields.Field[string] `json:"currency,required"`
	// The File containing the check's front image.
	FrontImageFileID fields.Field[string] `json:"front_image_file_id,required"`
	// The File containing the check's back image.
	BackImageFileID fields.Field[string] `json:"back_image_file_id,required"`
}

// MarshalJSON serializes CreateACheckDepositParameters into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *CreateACheckDepositParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CreateACheckDepositParameters) String() (result string) {
	return fmt.Sprintf("&CreateACheckDepositParameters{AccountID:%s Amount:%s Currency:%s FrontImageFileID:%s BackImageFileID:%s}", r.AccountID, r.Amount, r.Currency, r.FrontImageFileID, r.BackImageFileID)
}

type CheckDepositListParams struct {
	// Return the page of entries after this one.
	Cursor fields.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit fields.Field[int64] `query:"limit"`
	// Filter Check Deposits to those belonging to the specified Account.
	AccountID fields.Field[string]                          `query:"account_id"`
	CreatedAt fields.Field[CheckDepositListParamsCreatedAt] `query:"created_at"`
}

// URLQuery serializes CheckDepositListParams into a url.Values of the query
// parameters associated with this value
func (r *CheckDepositListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r CheckDepositListParams) String() (result string) {
	return fmt.Sprintf("&CheckDepositListParams{Cursor:%s Limit:%s AccountID:%s CreatedAt:%s}", r.Cursor, r.Limit, r.AccountID, r.CreatedAt)
}

type CheckDepositListParamsCreatedAt struct {
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

// URLQuery serializes CheckDepositListParamsCreatedAt into a url.Values of the
// query parameters associated with this value
func (r *CheckDepositListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r CheckDepositListParamsCreatedAt) String() (result string) {
	return fmt.Sprintf("&CheckDepositListParamsCreatedAt{After:%s Before:%s OnOrAfter:%s OnOrBefore:%s}", r.After, r.Before, r.OnOrAfter, r.OnOrBefore)
}

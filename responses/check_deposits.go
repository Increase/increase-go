package responses

import (
	"time"

	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/pagination"
)

type CheckDeposit struct {
	// The deposit's identifier.
	ID string `json:"id,required"`
	// The deposited amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the deposit.
	Currency CheckDepositCurrency `json:"currency,required"`
	// The status of the Check Deposit.
	Status CheckDepositStatus `json:"status,required"`
	// The Account the check was deposited into.
	AccountID string `json:"account_id,required"`
	// The ID for the File containing the image of the front of the check.
	FrontImageFileID string `json:"front_image_file_id,required"`
	// The ID for the File containing the image of the back of the check.
	BackImageFileID string `json:"back_image_file_id,required,nullable"`
	// The ID for the Transaction created by the deposit.
	TransactionID string `json:"transaction_id,required,nullable"`
	// If your deposit is successfully parsed and accepted by Increase, this will
	// contain details of the parsed check.
	DepositAcceptance CheckDepositDepositAcceptance `json:"deposit_acceptance,required,nullable"`
	// If your deposit is rejected by Increase, this will contain details as to why it
	// was rejected.
	DepositRejection CheckDepositDepositRejection `json:"deposit_rejection,required,nullable"`
	// If your deposit is returned, this will contain details as to why it was
	// returned.
	DepositReturn CheckDepositDepositReturn `json:"deposit_return,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `check_deposit`.
	Type CheckDepositType `json:"type,required"`
	JSON CheckDepositJSON
}

type CheckDepositJSON struct {
	ID                pjson.Metadata
	Amount            pjson.Metadata
	CreatedAt         pjson.Metadata
	Currency          pjson.Metadata
	Status            pjson.Metadata
	AccountID         pjson.Metadata
	FrontImageFileID  pjson.Metadata
	BackImageFileID   pjson.Metadata
	TransactionID     pjson.Metadata
	DepositAcceptance pjson.Metadata
	DepositRejection  pjson.Metadata
	DepositReturn     pjson.Metadata
	Type              pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CheckDeposit using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CheckDeposit) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency CheckDepositDepositAcceptanceCurrency `json:"currency,required"`
	// The account number printed on the check.
	AccountNumber string `json:"account_number,required"`
	// The routing number printed on the check.
	RoutingNumber string `json:"routing_number,required"`
	// An additional line of metadata printed on the check. This typically includes the
	// check number for business checks.
	AuxiliaryOnUs string `json:"auxiliary_on_us,required,nullable"`
	// The check serial number, if present, for consumer checks. For business checks,
	// the serial number is usually in the `auxiliary_on_us` field.
	SerialNumber string `json:"serial_number,required,nullable"`
	// The ID of the Check Deposit that was accepted.
	CheckDepositID string `json:"check_deposit_id,required"`
	JSON           CheckDepositDepositAcceptanceJSON
}

type CheckDepositDepositAcceptanceJSON struct {
	Amount         pjson.Metadata
	Currency       pjson.Metadata
	AccountNumber  pjson.Metadata
	RoutingNumber  pjson.Metadata
	AuxiliaryOnUs  pjson.Metadata
	SerialNumber   pjson.Metadata
	CheckDepositID pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CheckDepositDepositAcceptance
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *CheckDepositDepositAcceptance) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
	// currency.
	Currency CheckDepositDepositRejectionCurrency `json:"currency,required"`
	// Why the check deposit was rejected.
	Reason CheckDepositDepositRejectionReason `json:"reason,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check deposit was rejected.
	RejectedAt time.Time `json:"rejected_at,required" format:"date-time"`
	JSON       CheckDepositDepositRejectionJSON
}

type CheckDepositDepositRejectionJSON struct {
	Amount     pjson.Metadata
	Currency   pjson.Metadata
	Reason     pjson.Metadata
	RejectedAt pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CheckDepositDepositRejection
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *CheckDepositDepositRejection) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	Amount int64 `json:"amount,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check deposit was returned.
	ReturnedAt time.Time `json:"returned_at,required" format:"date-time"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency CheckDepositDepositReturnCurrency `json:"currency,required"`
	// The identifier of the Check Deposit that was returned.
	CheckDepositID string `json:"check_deposit_id,required"`
	// The identifier of the transaction that reversed the original check deposit
	// transaction.
	TransactionID string                                `json:"transaction_id,required"`
	ReturnReason  CheckDepositDepositReturnReturnReason `json:"return_reason,required"`
	JSON          CheckDepositDepositReturnJSON
}

type CheckDepositDepositReturnJSON struct {
	Amount         pjson.Metadata
	ReturnedAt     pjson.Metadata
	Currency       pjson.Metadata
	CheckDepositID pjson.Metadata
	TransactionID  pjson.Metadata
	ReturnReason   pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CheckDepositDepositReturn
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *CheckDepositDepositReturn) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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

type CheckDepositList struct {
	// The contents of the list.
	Data []CheckDeposit `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       CheckDepositListJSON
}

type CheckDepositListJSON struct {
	Data       pjson.Metadata
	NextCursor pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CheckDepositList using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CheckDepositList) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type CheckDepositsPage struct {
	*pagination.Page[CheckDeposit]
}

func (r *CheckDepositsPage) CheckDeposit() *CheckDeposit {
	return r.Current()
}

func (r *CheckDepositsPage) NextPage() (*CheckDepositsPage, error) {
	if page, err := r.Page.NextPage(); err != nil {
		return nil, err
	} else {
		return &CheckDepositsPage{page}, nil
	}
}

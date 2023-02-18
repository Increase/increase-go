package types

import (
	"fmt"
	"increase/core"
	"increase/core/pjson"
	"increase/pagination"
)

type CheckDeposit struct {
	// The deposit's identifier.
	ID *string `pjson:"id"`
	// The deposited amount in the minor unit of the destination account currency. For
	// dollars, for example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt *string `pjson:"created_at"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the deposit.
	Currency *CheckDepositCurrency `pjson:"currency"`
	// The status of the Check Deposit.
	Status *CheckDepositStatus `pjson:"status"`
	// The Account the check was deposited into.
	AccountID *string `pjson:"account_id"`
	// The ID for the File containing the image of the front of the check.
	FrontImageFileID *string `pjson:"front_image_file_id"`
	// The ID for the File containing the image of the back of the check.
	BackImageFileID *string `pjson:"back_image_file_id"`
	// The ID for the Transaction created by the deposit.
	TransactionID *string `pjson:"transaction_id"`
	// If your deposit is successfully parsed and accepted by Increase, this will
	// contain details of the parsed check.
	DepositAcceptance *CheckDepositDepositAcceptance `pjson:"deposit_acceptance"`
	// If your deposit is rejected by Increase, this will contain details as to why it
	// was rejected.
	DepositRejection *CheckDepositDepositRejection `pjson:"deposit_rejection"`
	// If your deposit is returned, this will contain details as to why it was
	// returned.
	DepositReturn *CheckDepositDepositReturn `pjson:"deposit_return"`
	// A constant representing the object's type. For this resource it will always be
	// `check_deposit`.
	Type       *CheckDepositType      `pjson:"type"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into CheckDeposit using the
// internal pjson library. Unrecognized fields are stored in the `Extras` property.
func (r *CheckDeposit) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CheckDeposit into an array of bytes using the gjson
// library. Members of the `Extras` field are serialized into the top-level, and
// will overwrite known members of the same name.
func (r *CheckDeposit) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The deposit's identifier.
func (r *CheckDeposit) GetID() (ID string) {
	if r != nil && r.ID != nil {
		ID = *r.ID
	}
	return
}

// The deposited amount in the minor unit of the destination account currency. For
// dollars, for example, this is cents.
func (r *CheckDeposit) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the transfer was created.
func (r *CheckDeposit) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the deposit.
func (r *CheckDeposit) GetCurrency() (Currency CheckDepositCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The status of the Check Deposit.
func (r *CheckDeposit) GetStatus() (Status CheckDepositStatus) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

// The Account the check was deposited into.
func (r *CheckDeposit) GetAccountID() (AccountID string) {
	if r != nil && r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

// The ID for the File containing the image of the front of the check.
func (r *CheckDeposit) GetFrontImageFileID() (FrontImageFileID string) {
	if r != nil && r.FrontImageFileID != nil {
		FrontImageFileID = *r.FrontImageFileID
	}
	return
}

// The ID for the File containing the image of the back of the check.
func (r *CheckDeposit) GetBackImageFileID() (BackImageFileID string) {
	if r != nil && r.BackImageFileID != nil {
		BackImageFileID = *r.BackImageFileID
	}
	return
}

// The ID for the Transaction created by the deposit.
func (r *CheckDeposit) GetTransactionID() (TransactionID string) {
	if r != nil && r.TransactionID != nil {
		TransactionID = *r.TransactionID
	}
	return
}

// If your deposit is successfully parsed and accepted by Increase, this will
// contain details of the parsed check.
func (r *CheckDeposit) GetDepositAcceptance() (DepositAcceptance CheckDepositDepositAcceptance) {
	if r != nil && r.DepositAcceptance != nil {
		DepositAcceptance = *r.DepositAcceptance
	}
	return
}

// If your deposit is rejected by Increase, this will contain details as to why it
// was rejected.
func (r *CheckDeposit) GetDepositRejection() (DepositRejection CheckDepositDepositRejection) {
	if r != nil && r.DepositRejection != nil {
		DepositRejection = *r.DepositRejection
	}
	return
}

// If your deposit is returned, this will contain details as to why it was
// returned.
func (r *CheckDeposit) GetDepositReturn() (DepositReturn CheckDepositDepositReturn) {
	if r != nil && r.DepositReturn != nil {
		DepositReturn = *r.DepositReturn
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `check_deposit`.
func (r *CheckDeposit) GetType() (Type CheckDepositType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

func (r CheckDeposit) String() (result string) {
	return fmt.Sprintf("&CheckDeposit{ID:%s Amount:%s CreatedAt:%s Currency:%s Status:%s AccountID:%s FrontImageFileID:%s BackImageFileID:%s TransactionID:%s DepositAcceptance:%s DepositRejection:%s DepositReturn:%s Type:%s}", core.FmtP(r.ID), core.FmtP(r.Amount), core.FmtP(r.CreatedAt), core.FmtP(r.Currency), core.FmtP(r.Status), core.FmtP(r.AccountID), core.FmtP(r.FrontImageFileID), core.FmtP(r.BackImageFileID), core.FmtP(r.TransactionID), r.DepositAcceptance, r.DepositRejection, r.DepositReturn, core.FmtP(r.Type))
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
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *CheckDepositDepositAcceptanceCurrency `pjson:"currency"`
	// The account number printed on the check.
	AccountNumber *string `pjson:"account_number"`
	// The routing number printed on the check.
	RoutingNumber *string `pjson:"routing_number"`
	// An additional line of metadata printed on the check. This typically includes the
	// check number for business checks.
	AuxiliaryOnUs *string `pjson:"auxiliary_on_us"`
	// The check serial number, if present, for consumer checks. For business checks,
	// the serial number is usually in the `auxiliary_on_us` field.
	SerialNumber *string `pjson:"serial_number"`
	// The ID of the Check Deposit that was accepted.
	CheckDepositID *string                `pjson:"check_deposit_id"`
	jsonFields     map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into CheckDepositDepositAcceptance
// using the internal pjson library. Unrecognized fields are stored in the `Extras`
// property.
func (r *CheckDepositDepositAcceptance) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CheckDepositDepositAcceptance into an array of bytes
// using the gjson library. Members of the `Extras` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CheckDepositDepositAcceptance) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount to be deposited in the minor unit of the transaction's currency. For
// dollars, for example, this is cents.
func (r *CheckDepositDepositAcceptance) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *CheckDepositDepositAcceptance) GetCurrency() (Currency CheckDepositDepositAcceptanceCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The account number printed on the check.
func (r *CheckDepositDepositAcceptance) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

// The routing number printed on the check.
func (r *CheckDepositDepositAcceptance) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

// An additional line of metadata printed on the check. This typically includes the
// check number for business checks.
func (r *CheckDepositDepositAcceptance) GetAuxiliaryOnUs() (AuxiliaryOnUs string) {
	if r != nil && r.AuxiliaryOnUs != nil {
		AuxiliaryOnUs = *r.AuxiliaryOnUs
	}
	return
}

// The check serial number, if present, for consumer checks. For business checks,
// the serial number is usually in the `auxiliary_on_us` field.
func (r *CheckDepositDepositAcceptance) GetSerialNumber() (SerialNumber string) {
	if r != nil && r.SerialNumber != nil {
		SerialNumber = *r.SerialNumber
	}
	return
}

// The ID of the Check Deposit that was accepted.
func (r *CheckDepositDepositAcceptance) GetCheckDepositID() (CheckDepositID string) {
	if r != nil && r.CheckDepositID != nil {
		CheckDepositID = *r.CheckDepositID
	}
	return
}

func (r CheckDepositDepositAcceptance) String() (result string) {
	return fmt.Sprintf("&CheckDepositDepositAcceptance{Amount:%s Currency:%s AccountNumber:%s RoutingNumber:%s AuxiliaryOnUs:%s SerialNumber:%s CheckDepositID:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.AccountNumber), core.FmtP(r.RoutingNumber), core.FmtP(r.AuxiliaryOnUs), core.FmtP(r.SerialNumber), core.FmtP(r.CheckDepositID))
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
	Amount *int64 `pjson:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
	// currency.
	Currency *CheckDepositDepositRejectionCurrency `pjson:"currency"`
	// Why the check deposit was rejected.
	Reason *CheckDepositDepositRejectionReason `pjson:"reason"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check deposit was rejected.
	RejectedAt *string                `pjson:"rejected_at"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into CheckDepositDepositRejection
// using the internal pjson library. Unrecognized fields are stored in the `Extras`
// property.
func (r *CheckDepositDepositRejection) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CheckDepositDepositRejection into an array of bytes using
// the gjson library. Members of the `Extras` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CheckDepositDepositRejection) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The rejected amount in the minor unit of check's currency. For dollars, for
// example, this is cents.
func (r *CheckDepositDepositRejection) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
// currency.
func (r *CheckDepositDepositRejection) GetCurrency() (Currency CheckDepositDepositRejectionCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// Why the check deposit was rejected.
func (r *CheckDepositDepositRejection) GetReason() (Reason CheckDepositDepositRejectionReason) {
	if r != nil && r.Reason != nil {
		Reason = *r.Reason
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the check deposit was rejected.
func (r *CheckDepositDepositRejection) GetRejectedAt() (RejectedAt string) {
	if r != nil && r.RejectedAt != nil {
		RejectedAt = *r.RejectedAt
	}
	return
}

func (r CheckDepositDepositRejection) String() (result string) {
	return fmt.Sprintf("&CheckDepositDepositRejection{Amount:%s Currency:%s Reason:%s RejectedAt:%s}", core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.Reason), core.FmtP(r.RejectedAt))
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
	Amount *int64 `pjson:"amount"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check deposit was returned.
	ReturnedAt *string `pjson:"returned_at"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
	// transaction's currency.
	Currency *CheckDepositDepositReturnCurrency `pjson:"currency"`
	// The identifier of the Check Deposit that was returned.
	CheckDepositID *string `pjson:"check_deposit_id"`
	// The identifier of the transaction that reversed the original check deposit
	// transaction.
	TransactionID *string                                `pjson:"transaction_id"`
	ReturnReason  *CheckDepositDepositReturnReturnReason `pjson:"return_reason"`
	jsonFields    map[string]interface{}                 `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into CheckDepositDepositReturn
// using the internal pjson library. Unrecognized fields are stored in the `Extras`
// property.
func (r *CheckDepositDepositReturn) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CheckDepositDepositReturn into an array of bytes using
// the gjson library. Members of the `Extras` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CheckDepositDepositReturn) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *CheckDepositDepositReturn) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the check deposit was returned.
func (r *CheckDepositDepositReturn) GetReturnedAt() (ReturnedAt string) {
	if r != nil && r.ReturnedAt != nil {
		ReturnedAt = *r.ReturnedAt
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the
// transaction's currency.
func (r *CheckDepositDepositReturn) GetCurrency() (Currency CheckDepositDepositReturnCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The identifier of the Check Deposit that was returned.
func (r *CheckDepositDepositReturn) GetCheckDepositID() (CheckDepositID string) {
	if r != nil && r.CheckDepositID != nil {
		CheckDepositID = *r.CheckDepositID
	}
	return
}

// The identifier of the transaction that reversed the original check deposit
// transaction.
func (r *CheckDepositDepositReturn) GetTransactionID() (TransactionID string) {
	if r != nil && r.TransactionID != nil {
		TransactionID = *r.TransactionID
	}
	return
}

func (r *CheckDepositDepositReturn) GetReturnReason() (ReturnReason CheckDepositDepositReturnReturnReason) {
	if r != nil && r.ReturnReason != nil {
		ReturnReason = *r.ReturnReason
	}
	return
}

func (r CheckDepositDepositReturn) String() (result string) {
	return fmt.Sprintf("&CheckDepositDepositReturn{Amount:%s ReturnedAt:%s Currency:%s CheckDepositID:%s TransactionID:%s ReturnReason:%s}", core.FmtP(r.Amount), core.FmtP(r.ReturnedAt), core.FmtP(r.Currency), core.FmtP(r.CheckDepositID), core.FmtP(r.TransactionID), core.FmtP(r.ReturnReason))
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
	AccountID *string `pjson:"account_id"`
	// The deposit amount in the minor unit of the account currency. For dollars, for
	// example, this is cents.
	Amount *int64 `pjson:"amount"`
	// The currency to use for the deposit.
	Currency *string `pjson:"currency"`
	// The File containing the check's front image.
	FrontImageFileID *string `pjson:"front_image_file_id"`
	// The File containing the check's back image.
	BackImageFileID *string                `pjson:"back_image_file_id"`
	jsonFields      map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into CreateACheckDepositParameters
// using the internal pjson library. Unrecognized fields are stored in the `Extras`
// property.
func (r *CreateACheckDepositParameters) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CreateACheckDepositParameters into an array of bytes
// using the gjson library. Members of the `Extras` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CreateACheckDepositParameters) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The identifier for the Account to deposit the check in.
func (r *CreateACheckDepositParameters) GetAccountID() (AccountID string) {
	if r != nil && r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

// The deposit amount in the minor unit of the account currency. For dollars, for
// example, this is cents.
func (r *CreateACheckDepositParameters) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The currency to use for the deposit.
func (r *CreateACheckDepositParameters) GetCurrency() (Currency string) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The File containing the check's front image.
func (r *CreateACheckDepositParameters) GetFrontImageFileID() (FrontImageFileID string) {
	if r != nil && r.FrontImageFileID != nil {
		FrontImageFileID = *r.FrontImageFileID
	}
	return
}

// The File containing the check's back image.
func (r *CreateACheckDepositParameters) GetBackImageFileID() (BackImageFileID string) {
	if r != nil && r.BackImageFileID != nil {
		BackImageFileID = *r.BackImageFileID
	}
	return
}

func (r CreateACheckDepositParameters) String() (result string) {
	return fmt.Sprintf("&CreateACheckDepositParameters{AccountID:%s Amount:%s Currency:%s FrontImageFileID:%s BackImageFileID:%s}", core.FmtP(r.AccountID), core.FmtP(r.Amount), core.FmtP(r.Currency), core.FmtP(r.FrontImageFileID), core.FmtP(r.BackImageFileID))
}

type CheckDepositListParams struct {
	// Return the page of entries after this one.
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int64 `query:"limit"`
	// Filter Check Deposits to those belonging to the specified Account.
	AccountID  *string                           `query:"account_id"`
	CreatedAt  *CheckDepositsListParamsCreatedAt `query:"created_at"`
	jsonFields map[string]interface{}            `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into CheckDepositListParams using
// the internal pjson library. Unrecognized fields are stored in the `Extras`
// property.
func (r *CheckDepositListParams) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CheckDepositListParams into an array of bytes using the
// gjson library. Members of the `Extras` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *CheckDepositListParams) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// Return the page of entries after this one.
func (r *CheckDepositListParams) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *CheckDepositListParams) GetLimit() (Limit int64) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// Filter Check Deposits to those belonging to the specified Account.
func (r *CheckDepositListParams) GetAccountID() (AccountID string) {
	if r != nil && r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

func (r *CheckDepositListParams) GetCreatedAt() (CreatedAt CheckDepositsListParamsCreatedAt) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

func (r CheckDepositListParams) String() (result string) {
	return fmt.Sprintf("&CheckDepositListParams{Cursor:%s Limit:%s AccountID:%s CreatedAt:%s}", core.FmtP(r.Cursor), core.FmtP(r.Limit), core.FmtP(r.AccountID), r.CreatedAt)
}

type CheckDepositsListParamsCreatedAt struct {
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
// CheckDepositsListParamsCreatedAt using the internal pjson library. Unrecognized
// fields are stored in the `Extras` property.
func (r *CheckDepositsListParamsCreatedAt) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CheckDepositsListParamsCreatedAt into an array of bytes
// using the gjson library. Members of the `Extras` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CheckDepositsListParamsCreatedAt) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r *CheckDepositsListParamsCreatedAt) GetAfter() (After string) {
	if r != nil && r.After != nil {
		After = *r.After
	}
	return
}

// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r *CheckDepositsListParamsCreatedAt) GetBefore() (Before string) {
	if r != nil && r.Before != nil {
		Before = *r.Before
	}
	return
}

// Return results on or after this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r *CheckDepositsListParamsCreatedAt) GetOnOrAfter() (OnOrAfter string) {
	if r != nil && r.OnOrAfter != nil {
		OnOrAfter = *r.OnOrAfter
	}
	return
}

// Return results on or before this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r *CheckDepositsListParamsCreatedAt) GetOnOrBefore() (OnOrBefore string) {
	if r != nil && r.OnOrBefore != nil {
		OnOrBefore = *r.OnOrBefore
	}
	return
}

func (r CheckDepositsListParamsCreatedAt) String() (result string) {
	return fmt.Sprintf("&CheckDepositsListParamsCreatedAt{After:%s Before:%s OnOrAfter:%s OnOrBefore:%s}", core.FmtP(r.After), core.FmtP(r.Before), core.FmtP(r.OnOrAfter), core.FmtP(r.OnOrBefore))
}

type CheckDepositList struct {
	// The contents of the list.
	Data *[]CheckDeposit `pjson:"data"`
	// A pointer to a place in the list.
	NextCursor *string                `pjson:"next_cursor"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into CheckDepositList using the
// internal pjson library. Unrecognized fields are stored in the `Extras` property.
func (r *CheckDepositList) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CheckDepositList into an array of bytes using the gjson
// library. Members of the `Extras` field are serialized into the top-level, and
// will overwrite known members of the same name.
func (r *CheckDepositList) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The contents of the list.
func (r *CheckDepositList) GetData() (Data []CheckDeposit) {
	if r != nil && r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r *CheckDepositList) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

func (r CheckDepositList) String() (result string) {
	return fmt.Sprintf("&CheckDepositList{Data:%s NextCursor:%s}", core.Fmt(r.Data), core.FmtP(r.NextCursor))
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

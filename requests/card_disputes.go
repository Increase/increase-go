package requests

import (
	"fmt"
	"net/url"
	"time"

	"github.com/increase/increase-go/core"
	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/fields"
)

type CardDispute struct {
	// The Card Dispute identifier.
	ID fields.Field[string] `json:"id,required"`
	// Why you disputed the Transaction in question.
	Explanation fields.Field[string] `json:"explanation,required"`
	// The results of the Dispute investigation.
	Status fields.Field[CardDisputeStatus] `json:"status,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card Dispute was created.
	CreatedAt fields.Field[time.Time] `json:"created_at,required" format:"date-time"`
	// The identifier of the Transaction that was disputed.
	DisputedTransactionID fields.Field[string] `json:"disputed_transaction_id,required"`
	// If the Card Dispute's status is `accepted`, this will contain details of the
	// successful dispute.
	Acceptance fields.Field[CardDisputeAcceptance] `json:"acceptance,required,nullable"`
	// If the Card Dispute's status is `rejected`, this will contain details of the
	// unsuccessful dispute.
	Rejection fields.Field[CardDisputeRejection] `json:"rejection,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `card_dispute`.
	Type fields.Field[CardDisputeType] `json:"type,required"`
}

// MarshalJSON serializes CardDispute into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *CardDispute) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CardDispute) String() (result string) {
	return fmt.Sprintf("&CardDispute{ID:%s Explanation:%s Status:%s CreatedAt:%s DisputedTransactionID:%s Acceptance:%s Rejection:%s Type:%s}", r.ID, r.Explanation, r.Status, r.CreatedAt, r.DisputedTransactionID, r.Acceptance, r.Rejection, r.Type)
}

type CardDisputeStatus string

const (
	CardDisputeStatusPendingReviewing CardDisputeStatus = "pending_reviewing"
	CardDisputeStatusAccepted         CardDisputeStatus = "accepted"
	CardDisputeStatusRejected         CardDisputeStatus = "rejected"
)

type CardDisputeAcceptance struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card Dispute was accepted.
	AcceptedAt fields.Field[time.Time] `json:"accepted_at,required" format:"date-time"`
	// The identifier of the Card Dispute that was accepted.
	CardDisputeID fields.Field[string] `json:"card_dispute_id,required"`
	// The identifier of the Transaction that was created to return the disputed funds
	// to your account.
	TransactionID fields.Field[string] `json:"transaction_id,required,nullable"`
}

// MarshalJSON serializes CardDisputeAcceptance into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CardDisputeAcceptance) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CardDisputeAcceptance) String() (result string) {
	return fmt.Sprintf("&CardDisputeAcceptance{AcceptedAt:%s CardDisputeID:%s TransactionID:%s}", r.AcceptedAt, r.CardDisputeID, r.TransactionID)
}

type CardDisputeRejection struct {
	// Why the Card Dispute was rejected.
	Explanation fields.Field[string] `json:"explanation,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card Dispute was rejected.
	RejectedAt fields.Field[time.Time] `json:"rejected_at,required" format:"date-time"`
	// The identifier of the Card Dispute that was rejected.
	CardDisputeID fields.Field[string] `json:"card_dispute_id,required"`
}

// MarshalJSON serializes CardDisputeRejection into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CardDisputeRejection) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CardDisputeRejection) String() (result string) {
	return fmt.Sprintf("&CardDisputeRejection{Explanation:%s RejectedAt:%s CardDisputeID:%s}", r.Explanation, r.RejectedAt, r.CardDisputeID)
}

type CardDisputeType string

const (
	CardDisputeTypeCardDispute CardDisputeType = "card_dispute"
)

type CreateACardDisputeParameters struct {
	// The Transaction you wish to dispute. This Transaction must have a `source_type`
	// of `card_settlement`.
	DisputedTransactionID fields.Field[string] `json:"disputed_transaction_id,required"`
	// Why you are disputing this Transaction.
	Explanation fields.Field[string] `json:"explanation,required"`
}

// MarshalJSON serializes CreateACardDisputeParameters into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CreateACardDisputeParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CreateACardDisputeParameters) String() (result string) {
	return fmt.Sprintf("&CreateACardDisputeParameters{DisputedTransactionID:%s Explanation:%s}", r.DisputedTransactionID, r.Explanation)
}

type CardDisputeListParams struct {
	// Return the page of entries after this one.
	Cursor fields.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit     fields.Field[int64]                          `query:"limit"`
	CreatedAt fields.Field[CardDisputeListParamsCreatedAt] `query:"created_at"`
	Status    fields.Field[CardDisputeListParamsStatus]    `query:"status"`
}

// URLQuery serializes CardDisputeListParams into a url.Values of the query
// parameters associated with this value
func (r *CardDisputeListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r CardDisputeListParams) String() (result string) {
	return fmt.Sprintf("&CardDisputeListParams{Cursor:%s Limit:%s CreatedAt:%s Status:%s}", r.Cursor, r.Limit, r.CreatedAt, r.Status)
}

type CardDisputeListParamsCreatedAt struct {
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

// URLQuery serializes CardDisputeListParamsCreatedAt into a url.Values of the
// query parameters associated with this value
func (r *CardDisputeListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r CardDisputeListParamsCreatedAt) String() (result string) {
	return fmt.Sprintf("&CardDisputeListParamsCreatedAt{After:%s Before:%s OnOrAfter:%s OnOrBefore:%s}", r.After, r.Before, r.OnOrAfter, r.OnOrBefore)
}

type CardDisputeListParamsStatus struct {
	// Return results whose value is in the provided list. For GET requests, this
	// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
	In fields.Field[[]CardDisputeListParamsStatusIn] `query:"in"`
}

// URLQuery serializes CardDisputeListParamsStatus into a url.Values of the query
// parameters associated with this value
func (r *CardDisputeListParamsStatus) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r CardDisputeListParamsStatus) String() (result string) {
	return fmt.Sprintf("&CardDisputeListParamsStatus{In:%s}", core.Fmt(r.In))
}

type CardDisputeListParamsStatusIn string

const (
	CardDisputeListParamsStatusInPendingReviewing CardDisputeListParamsStatusIn = "pending_reviewing"
	CardDisputeListParamsStatusInAccepted         CardDisputeListParamsStatusIn = "accepted"
	CardDisputeListParamsStatusInRejected         CardDisputeListParamsStatusIn = "rejected"
)

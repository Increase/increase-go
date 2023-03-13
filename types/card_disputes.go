package types

import (
	"fmt"
	"net/url"

	"github.com/increase/increase-go/core"
	"github.com/increase/increase-go/core/pjson"
	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/pagination"
)

type CardDispute struct {
	// The Card Dispute identifier.
	ID *string `pjson:"id"`
	// Why you disputed the Transaction in question.
	Explanation *string `pjson:"explanation"`
	// The results of the Dispute investigation.
	Status *CardDisputeStatus `pjson:"status"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card Dispute was created.
	CreatedAt *string `pjson:"created_at"`
	// The identifier of the Transaction that was disputed.
	DisputedTransactionID *string `pjson:"disputed_transaction_id"`
	// If the Card Dispute's status is `accepted`, this will contain details of the
	// successful dispute.
	Acceptance *CardDisputeAcceptance `pjson:"acceptance"`
	// If the Card Dispute's status is `rejected`, this will contain details of the
	// unsuccessful dispute.
	Rejection *CardDisputeRejection `pjson:"rejection"`
	// A constant representing the object's type. For this resource it will always be
	// `card_dispute`.
	Type       *CardDisputeType       `pjson:"type"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into CardDispute using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CardDispute) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CardDispute into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *CardDispute) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The Card Dispute identifier.
func (r *CardDispute) GetID() (ID string) {
	if r != nil && r.ID != nil {
		ID = *r.ID
	}
	return
}

// Why you disputed the Transaction in question.
func (r *CardDispute) GetExplanation() (Explanation string) {
	if r != nil && r.Explanation != nil {
		Explanation = *r.Explanation
	}
	return
}

// The results of the Dispute investigation.
func (r *CardDispute) GetStatus() (Status CardDisputeStatus) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the Card Dispute was created.
func (r *CardDispute) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The identifier of the Transaction that was disputed.
func (r *CardDispute) GetDisputedTransactionID() (DisputedTransactionID string) {
	if r != nil && r.DisputedTransactionID != nil {
		DisputedTransactionID = *r.DisputedTransactionID
	}
	return
}

// If the Card Dispute's status is `accepted`, this will contain details of the
// successful dispute.
func (r *CardDispute) GetAcceptance() (Acceptance CardDisputeAcceptance) {
	if r != nil && r.Acceptance != nil {
		Acceptance = *r.Acceptance
	}
	return
}

// If the Card Dispute's status is `rejected`, this will contain details of the
// unsuccessful dispute.
func (r *CardDispute) GetRejection() (Rejection CardDisputeRejection) {
	if r != nil && r.Rejection != nil {
		Rejection = *r.Rejection
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `card_dispute`.
func (r *CardDispute) GetType() (Type CardDisputeType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

func (r CardDispute) String() (result string) {
	return fmt.Sprintf("&CardDispute{ID:%s Explanation:%s Status:%s CreatedAt:%s DisputedTransactionID:%s Acceptance:%s Rejection:%s Type:%s}", core.FmtP(r.ID), core.FmtP(r.Explanation), core.FmtP(r.Status), core.FmtP(r.CreatedAt), core.FmtP(r.DisputedTransactionID), r.Acceptance, r.Rejection, core.FmtP(r.Type))
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
	AcceptedAt *string `pjson:"accepted_at"`
	// The identifier of the Card Dispute that was accepted.
	CardDisputeID *string `pjson:"card_dispute_id"`
	// The identifier of the Transaction that was created to return the disputed funds
	// to your account.
	TransactionID *string                `pjson:"transaction_id"`
	jsonFields    map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into CardDisputeAcceptance using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CardDisputeAcceptance) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CardDisputeAcceptance into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CardDisputeAcceptance) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the Card Dispute was accepted.
func (r *CardDisputeAcceptance) GetAcceptedAt() (AcceptedAt string) {
	if r != nil && r.AcceptedAt != nil {
		AcceptedAt = *r.AcceptedAt
	}
	return
}

// The identifier of the Card Dispute that was accepted.
func (r *CardDisputeAcceptance) GetCardDisputeID() (CardDisputeID string) {
	if r != nil && r.CardDisputeID != nil {
		CardDisputeID = *r.CardDisputeID
	}
	return
}

// The identifier of the Transaction that was created to return the disputed funds
// to your account.
func (r *CardDisputeAcceptance) GetTransactionID() (TransactionID string) {
	if r != nil && r.TransactionID != nil {
		TransactionID = *r.TransactionID
	}
	return
}

func (r CardDisputeAcceptance) String() (result string) {
	return fmt.Sprintf("&CardDisputeAcceptance{AcceptedAt:%s CardDisputeID:%s TransactionID:%s}", core.FmtP(r.AcceptedAt), core.FmtP(r.CardDisputeID), core.FmtP(r.TransactionID))
}

type CardDisputeRejection struct {
	// Why the Card Dispute was rejected.
	Explanation *string `pjson:"explanation"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card Dispute was rejected.
	RejectedAt *string `pjson:"rejected_at"`
	// The identifier of the Card Dispute that was rejected.
	CardDisputeID *string                `pjson:"card_dispute_id"`
	jsonFields    map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into CardDisputeRejection using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CardDisputeRejection) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CardDisputeRejection into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CardDisputeRejection) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// Why the Card Dispute was rejected.
func (r *CardDisputeRejection) GetExplanation() (Explanation string) {
	if r != nil && r.Explanation != nil {
		Explanation = *r.Explanation
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the Card Dispute was rejected.
func (r *CardDisputeRejection) GetRejectedAt() (RejectedAt string) {
	if r != nil && r.RejectedAt != nil {
		RejectedAt = *r.RejectedAt
	}
	return
}

// The identifier of the Card Dispute that was rejected.
func (r *CardDisputeRejection) GetCardDisputeID() (CardDisputeID string) {
	if r != nil && r.CardDisputeID != nil {
		CardDisputeID = *r.CardDisputeID
	}
	return
}

func (r CardDisputeRejection) String() (result string) {
	return fmt.Sprintf("&CardDisputeRejection{Explanation:%s RejectedAt:%s CardDisputeID:%s}", core.FmtP(r.Explanation), core.FmtP(r.RejectedAt), core.FmtP(r.CardDisputeID))
}

type CardDisputeType string

const (
	CardDisputeTypeCardDispute CardDisputeType = "card_dispute"
)

type CreateACardDisputeParameters struct {
	// The Transaction you wish to dispute. This Transaction must have a `source_type`
	// of `card_settlement`.
	DisputedTransactionID *string `pjson:"disputed_transaction_id"`
	// Why you are disputing this Transaction.
	Explanation *string                `pjson:"explanation"`
	jsonFields  map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into CreateACardDisputeParameters
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *CreateACardDisputeParameters) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CreateACardDisputeParameters into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CreateACardDisputeParameters) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The Transaction you wish to dispute. This Transaction must have a `source_type`
// of `card_settlement`.
func (r *CreateACardDisputeParameters) GetDisputedTransactionID() (DisputedTransactionID string) {
	if r != nil && r.DisputedTransactionID != nil {
		DisputedTransactionID = *r.DisputedTransactionID
	}
	return
}

// Why you are disputing this Transaction.
func (r *CreateACardDisputeParameters) GetExplanation() (Explanation string) {
	if r != nil && r.Explanation != nil {
		Explanation = *r.Explanation
	}
	return
}

func (r CreateACardDisputeParameters) String() (result string) {
	return fmt.Sprintf("&CreateACardDisputeParameters{DisputedTransactionID:%s Explanation:%s}", core.FmtP(r.DisputedTransactionID), core.FmtP(r.Explanation))
}

type CardDisputeListParams struct {
	// Return the page of entries after this one.
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit      *int64                           `query:"limit"`
	CreatedAt  *CardDisputesListParamsCreatedAt `query:"created_at"`
	Status     *CardDisputesListParamsStatus    `query:"status"`
	jsonFields map[string]interface{}           `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into CardDisputeListParams using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CardDisputeListParams) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CardDisputeListParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CardDisputeListParams) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes CardDisputeListParams into a url.Values of the query
// parameters associated with this value
func (r *CardDisputeListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// Return the page of entries after this one.
func (r *CardDisputeListParams) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *CardDisputeListParams) GetLimit() (Limit int64) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

func (r *CardDisputeListParams) GetCreatedAt() (CreatedAt CardDisputesListParamsCreatedAt) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

func (r *CardDisputeListParams) GetStatus() (Status CardDisputesListParamsStatus) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

func (r CardDisputeListParams) String() (result string) {
	return fmt.Sprintf("&CardDisputeListParams{Cursor:%s Limit:%s CreatedAt:%s Status:%s}", core.FmtP(r.Cursor), core.FmtP(r.Limit), r.CreatedAt, r.Status)
}

type CardDisputesListParamsCreatedAt struct {
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
// CardDisputesListParamsCreatedAt using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *CardDisputesListParamsCreatedAt) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CardDisputesListParamsCreatedAt into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *CardDisputesListParamsCreatedAt) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes CardDisputesListParamsCreatedAt into a url.Values of the
// query parameters associated with this value
func (r *CardDisputesListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r *CardDisputesListParamsCreatedAt) GetAfter() (After string) {
	if r != nil && r.After != nil {
		After = *r.After
	}
	return
}

// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r *CardDisputesListParamsCreatedAt) GetBefore() (Before string) {
	if r != nil && r.Before != nil {
		Before = *r.Before
	}
	return
}

// Return results on or after this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r *CardDisputesListParamsCreatedAt) GetOnOrAfter() (OnOrAfter string) {
	if r != nil && r.OnOrAfter != nil {
		OnOrAfter = *r.OnOrAfter
	}
	return
}

// Return results on or before this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r *CardDisputesListParamsCreatedAt) GetOnOrBefore() (OnOrBefore string) {
	if r != nil && r.OnOrBefore != nil {
		OnOrBefore = *r.OnOrBefore
	}
	return
}

func (r CardDisputesListParamsCreatedAt) String() (result string) {
	return fmt.Sprintf("&CardDisputesListParamsCreatedAt{After:%s Before:%s OnOrAfter:%s OnOrBefore:%s}", core.FmtP(r.After), core.FmtP(r.Before), core.FmtP(r.OnOrAfter), core.FmtP(r.OnOrBefore))
}

type CardDisputesListParamsStatus struct {
	// Return results whose value is in the provided list. For GET requests, this
	// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
	In         *[]CardDisputesListParamsStatusIn `pjson:"in"`
	jsonFields map[string]interface{}            `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into CardDisputesListParamsStatus
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *CardDisputesListParamsStatus) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CardDisputesListParamsStatus into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CardDisputesListParamsStatus) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes CardDisputesListParamsStatus into a url.Values of the query
// parameters associated with this value
func (r *CardDisputesListParamsStatus) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// Return results whose value is in the provided list. For GET requests, this
// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
func (r *CardDisputesListParamsStatus) GetIn() (In []CardDisputesListParamsStatusIn) {
	if r != nil && r.In != nil {
		In = *r.In
	}
	return
}

func (r CardDisputesListParamsStatus) String() (result string) {
	return fmt.Sprintf("&CardDisputesListParamsStatus{In:%s}", core.Fmt(r.In))
}

type CardDisputesListParamsStatusIn string

const (
	CardDisputesListParamsStatusInPendingReviewing CardDisputesListParamsStatusIn = "pending_reviewing"
	CardDisputesListParamsStatusInAccepted         CardDisputesListParamsStatusIn = "accepted"
	CardDisputesListParamsStatusInRejected         CardDisputesListParamsStatusIn = "rejected"
)

type CardDisputeList struct {
	// The contents of the list.
	Data *[]CardDispute `pjson:"data"`
	// A pointer to a place in the list.
	NextCursor *string                `pjson:"next_cursor"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into CardDisputeList using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CardDisputeList) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CardDisputeList into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *CardDisputeList) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes CardDisputeList into a url.Values of the query parameters
// associated with this value
func (r *CardDisputeList) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// The contents of the list.
func (r *CardDisputeList) GetData() (Data []CardDispute) {
	if r != nil && r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r *CardDisputeList) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

func (r CardDisputeList) String() (result string) {
	return fmt.Sprintf("&CardDisputeList{Data:%s NextCursor:%s}", core.Fmt(r.Data), core.FmtP(r.NextCursor))
}

type CardDisputesPage struct {
	*pagination.Page[CardDispute]
}

func (r *CardDisputesPage) CardDispute() *CardDispute {
	return r.Current()
}

func (r *CardDisputesPage) NextPage() (*CardDisputesPage, error) {
	if page, err := r.Page.NextPage(); err != nil {
		return nil, err
	} else {
		return &CardDisputesPage{page}, nil
	}
}

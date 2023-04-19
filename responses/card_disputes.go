package responses

import (
	"time"

	pjson "github.com/increase/increase-go/core/json"
)

type CardDispute struct {
	// The Card Dispute identifier.
	ID string `json:"id,required"`
	// Why you disputed the Transaction in question.
	Explanation string `json:"explanation,required"`
	// The results of the Dispute investigation.
	Status CardDisputeStatus `json:"status,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card Dispute was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The identifier of the Transaction that was disputed.
	DisputedTransactionID string `json:"disputed_transaction_id,required"`
	// If the Card Dispute's status is `accepted`, this will contain details of the
	// successful dispute.
	Acceptance CardDisputeAcceptance `json:"acceptance,required,nullable"`
	// If the Card Dispute's status is `rejected`, this will contain details of the
	// unsuccessful dispute.
	Rejection CardDisputeRejection `json:"rejection,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `card_dispute`.
	Type CardDisputeType `json:"type,required"`
	JSON CardDisputeJSON
}

type CardDisputeJSON struct {
	ID                    pjson.Metadata
	Explanation           pjson.Metadata
	Status                pjson.Metadata
	CreatedAt             pjson.Metadata
	DisputedTransactionID pjson.Metadata
	Acceptance            pjson.Metadata
	Rejection             pjson.Metadata
	Type                  pjson.Metadata
	Raw                   []byte
	Extras                map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CardDispute using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CardDispute) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	AcceptedAt time.Time `json:"accepted_at,required" format:"date-time"`
	// The identifier of the Card Dispute that was accepted.
	CardDisputeID string `json:"card_dispute_id,required"`
	// The identifier of the Transaction that was created to return the disputed funds
	// to your account.
	TransactionID string `json:"transaction_id,required"`
	JSON          CardDisputeAcceptanceJSON
}

type CardDisputeAcceptanceJSON struct {
	AcceptedAt    pjson.Metadata
	CardDisputeID pjson.Metadata
	TransactionID pjson.Metadata
	Raw           []byte
	Extras        map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CardDisputeAcceptance using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CardDisputeAcceptance) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type CardDisputeRejection struct {
	// Why the Card Dispute was rejected.
	Explanation string `json:"explanation,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card Dispute was rejected.
	RejectedAt time.Time `json:"rejected_at,required" format:"date-time"`
	// The identifier of the Card Dispute that was rejected.
	CardDisputeID string `json:"card_dispute_id,required"`
	JSON          CardDisputeRejectionJSON
}

type CardDisputeRejectionJSON struct {
	Explanation   pjson.Metadata
	RejectedAt    pjson.Metadata
	CardDisputeID pjson.Metadata
	Raw           []byte
	Extras        map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CardDisputeRejection using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CardDisputeRejection) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type CardDisputeType string

const (
	CardDisputeTypeCardDispute CardDisputeType = "card_dispute"
)

type CardDisputeListResponse struct {
	// The contents of the list.
	Data []CardDispute `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       CardDisputeListResponseJSON
}

type CardDisputeListResponseJSON struct {
	Data       pjson.Metadata
	NextCursor pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CardDisputeListResponse using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CardDisputeListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

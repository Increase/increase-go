package responses

import (
	"time"

	apijson "github.com/increase/increase-go/core/json"
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
	ID                    apijson.Metadata
	Explanation           apijson.Metadata
	Status                apijson.Metadata
	CreatedAt             apijson.Metadata
	DisputedTransactionID apijson.Metadata
	Acceptance            apijson.Metadata
	Rejection             apijson.Metadata
	Type                  apijson.Metadata
	Raw                   []byte
	Extras                map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CardDispute using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CardDispute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	AcceptedAt    apijson.Metadata
	CardDisputeID apijson.Metadata
	TransactionID apijson.Metadata
	Raw           []byte
	Extras        map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CardDisputeAcceptance using
// the internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CardDisputeAcceptance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	Explanation   apijson.Metadata
	RejectedAt    apijson.Metadata
	CardDisputeID apijson.Metadata
	Raw           []byte
	Extras        map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CardDisputeRejection using
// the internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CardDisputeRejection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	Data       apijson.Metadata
	NextCursor apijson.Metadata
	Raw        []byte
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CardDisputeListResponse using
// the internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CardDisputeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

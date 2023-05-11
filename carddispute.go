package increase

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/apiquery"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/internal/shared"
	"github.com/increase/increase-go/option"
)

// CardDisputeService contains methods and other services that help with
// interacting with the increase API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewCardDisputeService] method
// instead.
type CardDisputeService struct {
	Options []option.RequestOption
}

// NewCardDisputeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCardDisputeService(opts ...option.RequestOption) (r *CardDisputeService) {
	r = &CardDisputeService{}
	r.Options = opts
	return
}

// Create a Card Dispute
func (r *CardDisputeService) New(ctx context.Context, body CardDisputeNewParams, opts ...option.RequestOption) (res *CardDispute, err error) {
	opts = append(r.Options[:], opts...)
	path := "card_disputes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a Card Dispute
func (r *CardDisputeService) Get(ctx context.Context, card_dispute_id string, opts ...option.RequestOption) (res *CardDispute, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("card_disputes/%s", card_dispute_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Card Disputes
func (r *CardDisputeService) List(ctx context.Context, query CardDisputeListParams, opts ...option.RequestOption) (res *shared.Page[CardDispute], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "card_disputes"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List Card Disputes
func (r *CardDisputeService) ListAutoPaging(ctx context.Context, query CardDisputeListParams, opts ...option.RequestOption) *shared.PageAutoPager[CardDispute] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// If unauthorized activity occurs on a card, you can create a Card Dispute and
// we'll return the funds if appropriate.
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
	JSON cardDisputeJSON
}

// cardDisputeJSON contains the JSON metadata for the struct [CardDispute]
type cardDisputeJSON struct {
	ID                    apijson.Field
	Explanation           apijson.Field
	Status                apijson.Field
	CreatedAt             apijson.Field
	DisputedTransactionID apijson.Field
	Acceptance            apijson.Field
	Rejection             apijson.Field
	Type                  apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *CardDispute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CardDisputeStatus string

const (
	CardDisputeStatusPendingReviewing CardDisputeStatus = "pending_reviewing"
	CardDisputeStatusAccepted         CardDisputeStatus = "accepted"
	CardDisputeStatusRejected         CardDisputeStatus = "rejected"
)

// If the Card Dispute's status is `accepted`, this will contain details of the
// successful dispute.
type CardDisputeAcceptance struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card Dispute was accepted.
	AcceptedAt time.Time `json:"accepted_at,required" format:"date-time"`
	// The identifier of the Card Dispute that was accepted.
	CardDisputeID string `json:"card_dispute_id,required"`
	// The identifier of the Transaction that was created to return the disputed funds
	// to your account.
	TransactionID string `json:"transaction_id,required"`
	JSON          cardDisputeAcceptanceJSON
}

// cardDisputeAcceptanceJSON contains the JSON metadata for the struct
// [CardDisputeAcceptance]
type cardDisputeAcceptanceJSON struct {
	AcceptedAt    apijson.Field
	CardDisputeID apijson.Field
	TransactionID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CardDisputeAcceptance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// If the Card Dispute's status is `rejected`, this will contain details of the
// unsuccessful dispute.
type CardDisputeRejection struct {
	// Why the Card Dispute was rejected.
	Explanation string `json:"explanation,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card Dispute was rejected.
	RejectedAt time.Time `json:"rejected_at,required" format:"date-time"`
	// The identifier of the Card Dispute that was rejected.
	CardDisputeID string `json:"card_dispute_id,required"`
	JSON          cardDisputeRejectionJSON
}

// cardDisputeRejectionJSON contains the JSON metadata for the struct
// [CardDisputeRejection]
type cardDisputeRejectionJSON struct {
	Explanation   apijson.Field
	RejectedAt    apijson.Field
	CardDisputeID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CardDisputeRejection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CardDisputeType string

const (
	CardDisputeTypeCardDispute CardDisputeType = "card_dispute"
)

type CardDisputeNewParams struct {
	// The Transaction you wish to dispute. This Transaction must have a `source_type`
	// of `card_settlement`.
	DisputedTransactionID param.Field[string] `json:"disputed_transaction_id,required"`
	// Why you are disputing this Transaction.
	Explanation param.Field[string] `json:"explanation,required"`
}

func (r CardDisputeNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CardDisputeListParams struct {
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit     param.Field[int64]                          `query:"limit"`
	CreatedAt param.Field[CardDisputeListParamsCreatedAt] `query:"created_at"`
	Status    param.Field[CardDisputeListParamsStatus]    `query:"status"`
}

// URLQuery serializes [CardDisputeListParams]'s query parameters as `url.Values`.
func (r CardDisputeListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type CardDisputeListParamsCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After param.Field[time.Time] `query:"after" format:"date-time"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before param.Field[time.Time] `query:"before" format:"date-time"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter param.Field[time.Time] `query:"on_or_after" format:"date-time"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore param.Field[time.Time] `query:"on_or_before" format:"date-time"`
}

// URLQuery serializes [CardDisputeListParamsCreatedAt]'s query parameters as
// `url.Values`.
func (r CardDisputeListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type CardDisputeListParamsStatus struct {
	// Filter Card Disputes for those with the specified status or statuses. For GET
	// requests, this should be encoded as a comma-delimited string, such as
	// `?in=one,two,three`.
	In param.Field[[]CardDisputeListParamsStatusIn] `query:"in"`
}

// URLQuery serializes [CardDisputeListParamsStatus]'s query parameters as
// `url.Values`.
func (r CardDisputeListParamsStatus) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type CardDisputeListParamsStatusIn string

const (
	CardDisputeListParamsStatusInPendingReviewing CardDisputeListParamsStatusIn = "pending_reviewing"
	CardDisputeListParamsStatusInAccepted         CardDisputeListParamsStatusIn = "accepted"
	CardDisputeListParamsStatusInRejected         CardDisputeListParamsStatusIn = "rejected"
)

// A list of Card Dispute objects
type CardDisputeListResponse struct {
	// The contents of the list.
	Data []CardDispute `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       cardDisputeListResponseJSON
}

// cardDisputeListResponseJSON contains the JSON metadata for the struct
// [CardDisputeListResponse]
type cardDisputeListResponseJSON struct {
	Data        apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

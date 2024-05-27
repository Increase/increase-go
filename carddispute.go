// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/apiquery"
	"github.com/increase/increase-go/internal/pagination"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/option"
)

// CardDisputeService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCardDisputeService] method instead.
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
func (r *CardDisputeService) Get(ctx context.Context, cardDisputeID string, opts ...option.RequestOption) (res *CardDispute, err error) {
	opts = append(r.Options[:], opts...)
	if cardDisputeID == "" {
		err = errors.New("missing required card_dispute_id parameter")
		return
	}
	path := fmt.Sprintf("card_disputes/%s", cardDisputeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Card Disputes
func (r *CardDisputeService) List(ctx context.Context, query CardDisputeListParams, opts ...option.RequestOption) (res *pagination.Page[CardDispute], err error) {
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
func (r *CardDisputeService) ListAutoPaging(ctx context.Context, query CardDisputeListParams, opts ...option.RequestOption) *pagination.PageAutoPager[CardDispute] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// If unauthorized activity occurs on a card, you can create a Card Dispute and
// we'll return the funds if appropriate.
type CardDispute struct {
	// The Card Dispute identifier.
	ID string `json:"id,required"`
	// If the Card Dispute's status is `accepted`, this will contain details of the
	// successful dispute.
	Acceptance CardDisputeAcceptance `json:"acceptance,required,nullable"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card Dispute was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The identifier of the Transaction that was disputed.
	DisputedTransactionID string `json:"disputed_transaction_id,required"`
	// Why you disputed the Transaction in question.
	Explanation string `json:"explanation,required"`
	// The idempotency key you chose for this object. This value is unique across
	// Increase and is used to ensure that a request is only processed once. Learn more
	// about [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey string `json:"idempotency_key,required,nullable"`
	// If the Card Dispute's status is `rejected`, this will contain details of the
	// unsuccessful dispute.
	Rejection CardDisputeRejection `json:"rejection,required,nullable"`
	// The results of the Dispute investigation.
	Status CardDisputeStatus `json:"status,required"`
	// A constant representing the object's type. For this resource it will always be
	// `card_dispute`.
	Type CardDisputeType `json:"type,required"`
	JSON cardDisputeJSON `json:"-"`
}

// cardDisputeJSON contains the JSON metadata for the struct [CardDispute]
type cardDisputeJSON struct {
	ID                    apijson.Field
	Acceptance            apijson.Field
	CreatedAt             apijson.Field
	DisputedTransactionID apijson.Field
	Explanation           apijson.Field
	IdempotencyKey        apijson.Field
	Rejection             apijson.Field
	Status                apijson.Field
	Type                  apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *CardDispute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeJSON) RawJSON() string {
	return r.raw
}

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
	TransactionID string                    `json:"transaction_id,required"`
	JSON          cardDisputeAcceptanceJSON `json:"-"`
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

func (r cardDisputeAcceptanceJSON) RawJSON() string {
	return r.raw
}

// If the Card Dispute's status is `rejected`, this will contain details of the
// unsuccessful dispute.
type CardDisputeRejection struct {
	// The identifier of the Card Dispute that was rejected.
	CardDisputeID string `json:"card_dispute_id,required"`
	// Why the Card Dispute was rejected.
	Explanation string `json:"explanation,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card Dispute was rejected.
	RejectedAt time.Time                `json:"rejected_at,required" format:"date-time"`
	JSON       cardDisputeRejectionJSON `json:"-"`
}

// cardDisputeRejectionJSON contains the JSON metadata for the struct
// [CardDisputeRejection]
type cardDisputeRejectionJSON struct {
	CardDisputeID apijson.Field
	Explanation   apijson.Field
	RejectedAt    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CardDisputeRejection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeRejectionJSON) RawJSON() string {
	return r.raw
}

// The results of the Dispute investigation.
type CardDisputeStatus string

const (
	// The Card Dispute is pending review.
	CardDisputeStatusPendingReviewing CardDisputeStatus = "pending_reviewing"
	// The Card Dispute has been accepted and your funds have been returned.
	CardDisputeStatusAccepted CardDisputeStatus = "accepted"
	// The Card Dispute has been rejected.
	CardDisputeStatusRejected CardDisputeStatus = "rejected"
)

func (r CardDisputeStatus) IsKnown() bool {
	switch r {
	case CardDisputeStatusPendingReviewing, CardDisputeStatusAccepted, CardDisputeStatusRejected:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `card_dispute`.
type CardDisputeType string

const (
	CardDisputeTypeCardDispute CardDisputeType = "card_dispute"
)

func (r CardDisputeType) IsKnown() bool {
	switch r {
	case CardDisputeTypeCardDispute:
		return true
	}
	return false
}

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
	CreatedAt param.Field[CardDisputeListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Filter records to the one with the specified `idempotency_key` you chose for
	// that object. This value is unique across Increase and is used to ensure that a
	// request is only processed once. Learn more about
	// [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey param.Field[string] `query:"idempotency_key"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit  param.Field[int64]                       `query:"limit"`
	Status param.Field[CardDisputeListParamsStatus] `query:"status"`
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
	// The Card Dispute is pending review.
	CardDisputeListParamsStatusInPendingReviewing CardDisputeListParamsStatusIn = "pending_reviewing"
	// The Card Dispute has been accepted and your funds have been returned.
	CardDisputeListParamsStatusInAccepted CardDisputeListParamsStatusIn = "accepted"
	// The Card Dispute has been rejected.
	CardDisputeListParamsStatusInRejected CardDisputeListParamsStatusIn = "rejected"
)

func (r CardDisputeListParamsStatusIn) IsKnown() bool {
	switch r {
	case CardDisputeListParamsStatusInPendingReviewing, CardDisputeListParamsStatusInAccepted, CardDisputeListParamsStatusInRejected:
		return true
	}
	return false
}

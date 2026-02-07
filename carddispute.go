// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Increase/increase-go/internal/apijson"
	"github.com/Increase/increase-go/internal/apiquery"
	"github.com/Increase/increase-go/internal/param"
	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
	"github.com/Increase/increase-go/packages/pagination"
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
	opts = slices.Concat(r.Options, opts)
	path := "card_disputes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a Card Dispute
func (r *CardDisputeService) Get(ctx context.Context, cardDisputeID string, opts ...option.RequestOption) (res *CardDispute, err error) {
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
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

// Submit a User Submission for a Card Dispute
func (r *CardDisputeService) SubmitUserSubmission(ctx context.Context, cardDisputeID string, body CardDisputeSubmitUserSubmissionParams, opts ...option.RequestOption) (res *CardDispute, err error) {
	opts = slices.Concat(r.Options, opts)
	if cardDisputeID == "" {
		err = errors.New("missing required card_dispute_id parameter")
		return
	}
	path := fmt.Sprintf("card_disputes/%s/submit_user_submission", cardDisputeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Withdraw a Card Dispute
func (r *CardDisputeService) Withdraw(ctx context.Context, cardDisputeID string, body CardDisputeWithdrawParams, opts ...option.RequestOption) (res *CardDispute, err error) {
	opts = slices.Concat(r.Options, opts)
	if cardDisputeID == "" {
		err = errors.New("missing required card_dispute_id parameter")
		return
	}
	path := fmt.Sprintf("card_disputes/%s/withdraw", cardDisputeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// If unauthorized activity occurs on a card, you can create a Card Dispute and
// we'll work with the card networks to return the funds if appropriate.
type CardDispute struct {
	// The Card Dispute identifier.
	ID string `json:"id,required"`
	// The amount of the dispute.
	Amount int64 `json:"amount,required"`
	// The Card that the Card Dispute is associated with.
	CardID string `json:"card_id,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card Dispute was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The identifier of the Transaction that was disputed.
	DisputedTransactionID string `json:"disputed_transaction_id,required"`
	// The idempotency key you chose for this object. This value is unique across
	// Increase and is used to ensure that a request is only processed once. Learn more
	// about [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey string `json:"idempotency_key,required,nullable"`
	// If the Card Dispute's status is `lost`, this will contain details of the lost
	// dispute.
	Loss CardDisputeLoss `json:"loss,required,nullable"`
	// The network that the Card Dispute is associated with.
	Network CardDisputeNetwork `json:"network,required"`
	// The status of the Card Dispute.
	Status CardDisputeStatus `json:"status,required"`
	// A constant representing the object's type. For this resource it will always be
	// `card_dispute`.
	Type CardDisputeType `json:"type,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the user submission is required by. Present only if status is
	// `user_submission_required` and a user submission is required by a certain time.
	// Otherwise, this will be `nil`.
	UserSubmissionRequiredBy time.Time `json:"user_submission_required_by,required,nullable" format:"date-time"`
	// Card Dispute information for card payments processed over Visa's network. This
	// field will be present in the JSON response if and only if `network` is equal to
	// `visa`.
	Visa CardDisputeVisa `json:"visa,required,nullable"`
	// If the Card Dispute's status is `won`, this will contain details of the won
	// dispute.
	Win CardDisputeWin `json:"win,required,nullable"`
	// If the Card Dispute has been withdrawn, this will contain details of the
	// withdrawal.
	Withdrawal CardDisputeWithdrawal `json:"withdrawal,required,nullable"`
	JSON       cardDisputeJSON       `json:"-"`
}

// cardDisputeJSON contains the JSON metadata for the struct [CardDispute]
type cardDisputeJSON struct {
	ID                       apijson.Field
	Amount                   apijson.Field
	CardID                   apijson.Field
	CreatedAt                apijson.Field
	DisputedTransactionID    apijson.Field
	IdempotencyKey           apijson.Field
	Loss                     apijson.Field
	Network                  apijson.Field
	Status                   apijson.Field
	Type                     apijson.Field
	UserSubmissionRequiredBy apijson.Field
	Visa                     apijson.Field
	Win                      apijson.Field
	Withdrawal               apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *CardDispute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeJSON) RawJSON() string {
	return r.raw
}

// If the Card Dispute's status is `lost`, this will contain details of the lost
// dispute.
type CardDisputeLoss struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card Dispute was lost.
	LostAt time.Time `json:"lost_at,required" format:"date-time"`
	// The reason the Card Dispute was lost.
	Reason CardDisputeLossReason `json:"reason,required"`
	JSON   cardDisputeLossJSON   `json:"-"`
}

// cardDisputeLossJSON contains the JSON metadata for the struct [CardDisputeLoss]
type cardDisputeLossJSON struct {
	LostAt      apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeLoss) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeLossJSON) RawJSON() string {
	return r.raw
}

// The reason the Card Dispute was lost.
type CardDisputeLossReason string

const (
	CardDisputeLossReasonUserWithdrawn CardDisputeLossReason = "user_withdrawn"
	CardDisputeLossReasonLoss          CardDisputeLossReason = "loss"
)

func (r CardDisputeLossReason) IsKnown() bool {
	switch r {
	case CardDisputeLossReasonUserWithdrawn, CardDisputeLossReasonLoss:
		return true
	}
	return false
}

// The network that the Card Dispute is associated with.
type CardDisputeNetwork string

const (
	CardDisputeNetworkVisa  CardDisputeNetwork = "visa"
	CardDisputeNetworkPulse CardDisputeNetwork = "pulse"
)

func (r CardDisputeNetwork) IsKnown() bool {
	switch r {
	case CardDisputeNetworkVisa, CardDisputeNetworkPulse:
		return true
	}
	return false
}

// The status of the Card Dispute.
type CardDisputeStatus string

const (
	CardDisputeStatusUserSubmissionRequired          CardDisputeStatus = "user_submission_required"
	CardDisputeStatusPendingUserSubmissionReviewing  CardDisputeStatus = "pending_user_submission_reviewing"
	CardDisputeStatusPendingUserSubmissionSubmitting CardDisputeStatus = "pending_user_submission_submitting"
	CardDisputeStatusPendingUserWithdrawalSubmitting CardDisputeStatus = "pending_user_withdrawal_submitting"
	CardDisputeStatusPendingResponse                 CardDisputeStatus = "pending_response"
	CardDisputeStatusLost                            CardDisputeStatus = "lost"
	CardDisputeStatusWon                             CardDisputeStatus = "won"
)

func (r CardDisputeStatus) IsKnown() bool {
	switch r {
	case CardDisputeStatusUserSubmissionRequired, CardDisputeStatusPendingUserSubmissionReviewing, CardDisputeStatusPendingUserSubmissionSubmitting, CardDisputeStatusPendingUserWithdrawalSubmitting, CardDisputeStatusPendingResponse, CardDisputeStatusLost, CardDisputeStatusWon:
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

// Card Dispute information for card payments processed over Visa's network. This
// field will be present in the JSON response if and only if `network` is equal to
// `visa`.
type CardDisputeVisa struct {
	// The network events for the Card Dispute.
	NetworkEvents []CardDisputeVisaNetworkEvent `json:"network_events,required"`
	// The category of the currently required user submission if the user wishes to
	// proceed with the dispute. Present if and only if status is
	// `user_submission_required`. Otherwise, this will be `nil`.
	RequiredUserSubmissionCategory CardDisputeVisaRequiredUserSubmissionCategory `json:"required_user_submission_category,required,nullable"`
	// The user submissions for the Card Dispute.
	UserSubmissions []CardDisputeVisaUserSubmission `json:"user_submissions,required"`
	JSON            cardDisputeVisaJSON             `json:"-"`
}

// cardDisputeVisaJSON contains the JSON metadata for the struct [CardDisputeVisa]
type cardDisputeVisaJSON struct {
	NetworkEvents                  apijson.Field
	RequiredUserSubmissionCategory apijson.Field
	UserSubmissions                apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *CardDisputeVisa) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaJSON) RawJSON() string {
	return r.raw
}

type CardDisputeVisaNetworkEvent struct {
	// The files attached to the Visa Card Dispute User Submission.
	AttachmentFiles []CardDisputeVisaNetworkEventsAttachmentFile `json:"attachment_files,required"`
	// The category of the user submission. We may add additional possible values for
	// this enum over time; your application should be able to handle such additions
	// gracefully.
	Category CardDisputeVisaNetworkEventsCategory `json:"category,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Visa Card Dispute Network Event was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The dispute financial transaction that resulted from the network event, if any.
	DisputeFinancialTransactionID string `json:"dispute_financial_transaction_id,required,nullable"`
	// A Card Dispute Chargeback Accepted Visa Network Event object. This field will be
	// present in the JSON response if and only if `category` is equal to
	// `chargeback_accepted`. Contains the details specific to a chargeback accepted
	// Visa Card Dispute Network Event, which represents that a chargeback has been
	// accepted by the merchant.
	ChargebackAccepted CardDisputeVisaNetworkEventsChargebackAccepted `json:"chargeback_accepted,nullable"`
	// A Card Dispute Chargeback Submitted Visa Network Event object. This field will
	// be present in the JSON response if and only if `category` is equal to
	// `chargeback_submitted`. Contains the details specific to a chargeback submitted
	// Visa Card Dispute Network Event, which represents that a chargeback has been
	// submitted to the network.
	ChargebackSubmitted CardDisputeVisaNetworkEventsChargebackSubmitted `json:"chargeback_submitted,nullable"`
	// A Card Dispute Chargeback Timed Out Visa Network Event object. This field will
	// be present in the JSON response if and only if `category` is equal to
	// `chargeback_timed_out`. Contains the details specific to a chargeback timed out
	// Visa Card Dispute Network Event, which represents that the chargeback has timed
	// out in the user's favor.
	ChargebackTimedOut CardDisputeVisaNetworkEventsChargebackTimedOut `json:"chargeback_timed_out,nullable"`
	// A Card Dispute Merchant Pre-Arbitration Decline Submitted Visa Network Event
	// object. This field will be present in the JSON response if and only if
	// `category` is equal to `merchant_prearbitration_decline_submitted`. Contains the
	// details specific to a merchant prearbitration decline submitted Visa Card
	// Dispute Network Event, which represents that the user has declined the
	// merchant's request for a prearbitration request decision in their favor.
	MerchantPrearbitrationDeclineSubmitted CardDisputeVisaNetworkEventsMerchantPrearbitrationDeclineSubmitted `json:"merchant_prearbitration_decline_submitted,nullable"`
	// A Card Dispute Merchant Pre-Arbitration Received Visa Network Event object. This
	// field will be present in the JSON response if and only if `category` is equal to
	// `merchant_prearbitration_received`. Contains the details specific to a merchant
	// prearbitration received Visa Card Dispute Network Event, which represents that
	// the merchant has issued a prearbitration request in the user's favor.
	MerchantPrearbitrationReceived CardDisputeVisaNetworkEventsMerchantPrearbitrationReceived `json:"merchant_prearbitration_received,nullable"`
	// A Card Dispute Merchant Pre-Arbitration Timed Out Visa Network Event object.
	// This field will be present in the JSON response if and only if `category` is
	// equal to `merchant_prearbitration_timed_out`. Contains the details specific to a
	// merchant prearbitration timed out Visa Card Dispute Network Event, which
	// represents that the user has timed out responding to the merchant's
	// prearbitration request.
	MerchantPrearbitrationTimedOut CardDisputeVisaNetworkEventsMerchantPrearbitrationTimedOut `json:"merchant_prearbitration_timed_out,nullable"`
	// A Card Dispute Re-presented Visa Network Event object. This field will be
	// present in the JSON response if and only if `category` is equal to
	// `represented`. Contains the details specific to a re-presented Visa Card Dispute
	// Network Event, which represents that the merchant has declined the user's
	// chargeback and has re-presented the payment.
	Represented CardDisputeVisaNetworkEventsRepresented `json:"represented,nullable"`
	// A Card Dispute Re-presentment Timed Out Visa Network Event object. This field
	// will be present in the JSON response if and only if `category` is equal to
	// `representment_timed_out`. Contains the details specific to a re-presentment
	// time-out Visa Card Dispute Network Event, which represents that the user did not
	// respond to the re-presentment by the merchant within the time limit.
	RepresentmentTimedOut CardDisputeVisaNetworkEventsRepresentmentTimedOut `json:"representment_timed_out,nullable"`
	// A Card Dispute User Pre-Arbitration Accepted Visa Network Event object. This
	// field will be present in the JSON response if and only if `category` is equal to
	// `user_prearbitration_accepted`. Contains the details specific to a user
	// prearbitration accepted Visa Card Dispute Network Event, which represents that
	// the merchant has accepted the user's prearbitration request in the user's favor.
	UserPrearbitrationAccepted CardDisputeVisaNetworkEventsUserPrearbitrationAccepted `json:"user_prearbitration_accepted,nullable"`
	// A Card Dispute User Pre-Arbitration Declined Visa Network Event object. This
	// field will be present in the JSON response if and only if `category` is equal to
	// `user_prearbitration_declined`. Contains the details specific to a user
	// prearbitration declined Visa Card Dispute Network Event, which represents that
	// the merchant has declined the user's prearbitration request.
	UserPrearbitrationDeclined CardDisputeVisaNetworkEventsUserPrearbitrationDeclined `json:"user_prearbitration_declined,nullable"`
	// A Card Dispute User Pre-Arbitration Submitted Visa Network Event object. This
	// field will be present in the JSON response if and only if `category` is equal to
	// `user_prearbitration_submitted`. Contains the details specific to a user
	// prearbitration submitted Visa Card Dispute Network Event, which represents that
	// the user's request for prearbitration has been submitted to the network.
	UserPrearbitrationSubmitted CardDisputeVisaNetworkEventsUserPrearbitrationSubmitted `json:"user_prearbitration_submitted,nullable"`
	// A Card Dispute User Pre-Arbitration Timed Out Visa Network Event object. This
	// field will be present in the JSON response if and only if `category` is equal to
	// `user_prearbitration_timed_out`. Contains the details specific to a user
	// prearbitration timed out Visa Card Dispute Network Event, which represents that
	// the merchant has timed out responding to the user's prearbitration request.
	UserPrearbitrationTimedOut CardDisputeVisaNetworkEventsUserPrearbitrationTimedOut `json:"user_prearbitration_timed_out,nullable"`
	// A Card Dispute User Withdrawal Submitted Visa Network Event object. This field
	// will be present in the JSON response if and only if `category` is equal to
	// `user_withdrawal_submitted`. Contains the details specific to a user withdrawal
	// submitted Visa Card Dispute Network Event, which represents that the user's
	// request to withdraw the dispute has been submitted to the network.
	UserWithdrawalSubmitted CardDisputeVisaNetworkEventsUserWithdrawalSubmitted `json:"user_withdrawal_submitted,nullable"`
	JSON                    cardDisputeVisaNetworkEventJSON                     `json:"-"`
}

// cardDisputeVisaNetworkEventJSON contains the JSON metadata for the struct
// [CardDisputeVisaNetworkEvent]
type cardDisputeVisaNetworkEventJSON struct {
	AttachmentFiles                        apijson.Field
	Category                               apijson.Field
	CreatedAt                              apijson.Field
	DisputeFinancialTransactionID          apijson.Field
	ChargebackAccepted                     apijson.Field
	ChargebackSubmitted                    apijson.Field
	ChargebackTimedOut                     apijson.Field
	MerchantPrearbitrationDeclineSubmitted apijson.Field
	MerchantPrearbitrationReceived         apijson.Field
	MerchantPrearbitrationTimedOut         apijson.Field
	Represented                            apijson.Field
	RepresentmentTimedOut                  apijson.Field
	UserPrearbitrationAccepted             apijson.Field
	UserPrearbitrationDeclined             apijson.Field
	UserPrearbitrationSubmitted            apijson.Field
	UserPrearbitrationTimedOut             apijson.Field
	UserWithdrawalSubmitted                apijson.Field
	raw                                    string
	ExtraFields                            map[string]apijson.Field
}

func (r *CardDisputeVisaNetworkEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaNetworkEventJSON) RawJSON() string {
	return r.raw
}

type CardDisputeVisaNetworkEventsAttachmentFile struct {
	// The ID of the file attached to the Card Dispute.
	FileID string                                         `json:"file_id,required"`
	JSON   cardDisputeVisaNetworkEventsAttachmentFileJSON `json:"-"`
}

// cardDisputeVisaNetworkEventsAttachmentFileJSON contains the JSON metadata for
// the struct [CardDisputeVisaNetworkEventsAttachmentFile]
type cardDisputeVisaNetworkEventsAttachmentFileJSON struct {
	FileID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeVisaNetworkEventsAttachmentFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaNetworkEventsAttachmentFileJSON) RawJSON() string {
	return r.raw
}

// The category of the user submission. We may add additional possible values for
// this enum over time; your application should be able to handle such additions
// gracefully.
type CardDisputeVisaNetworkEventsCategory string

const (
	CardDisputeVisaNetworkEventsCategoryChargebackAccepted                     CardDisputeVisaNetworkEventsCategory = "chargeback_accepted"
	CardDisputeVisaNetworkEventsCategoryChargebackSubmitted                    CardDisputeVisaNetworkEventsCategory = "chargeback_submitted"
	CardDisputeVisaNetworkEventsCategoryChargebackTimedOut                     CardDisputeVisaNetworkEventsCategory = "chargeback_timed_out"
	CardDisputeVisaNetworkEventsCategoryMerchantPrearbitrationDeclineSubmitted CardDisputeVisaNetworkEventsCategory = "merchant_prearbitration_decline_submitted"
	CardDisputeVisaNetworkEventsCategoryMerchantPrearbitrationReceived         CardDisputeVisaNetworkEventsCategory = "merchant_prearbitration_received"
	CardDisputeVisaNetworkEventsCategoryMerchantPrearbitrationTimedOut         CardDisputeVisaNetworkEventsCategory = "merchant_prearbitration_timed_out"
	CardDisputeVisaNetworkEventsCategoryRepresented                            CardDisputeVisaNetworkEventsCategory = "represented"
	CardDisputeVisaNetworkEventsCategoryRepresentmentTimedOut                  CardDisputeVisaNetworkEventsCategory = "representment_timed_out"
	CardDisputeVisaNetworkEventsCategoryUserPrearbitrationAccepted             CardDisputeVisaNetworkEventsCategory = "user_prearbitration_accepted"
	CardDisputeVisaNetworkEventsCategoryUserPrearbitrationDeclined             CardDisputeVisaNetworkEventsCategory = "user_prearbitration_declined"
	CardDisputeVisaNetworkEventsCategoryUserPrearbitrationSubmitted            CardDisputeVisaNetworkEventsCategory = "user_prearbitration_submitted"
	CardDisputeVisaNetworkEventsCategoryUserPrearbitrationTimedOut             CardDisputeVisaNetworkEventsCategory = "user_prearbitration_timed_out"
	CardDisputeVisaNetworkEventsCategoryUserWithdrawalSubmitted                CardDisputeVisaNetworkEventsCategory = "user_withdrawal_submitted"
)

func (r CardDisputeVisaNetworkEventsCategory) IsKnown() bool {
	switch r {
	case CardDisputeVisaNetworkEventsCategoryChargebackAccepted, CardDisputeVisaNetworkEventsCategoryChargebackSubmitted, CardDisputeVisaNetworkEventsCategoryChargebackTimedOut, CardDisputeVisaNetworkEventsCategoryMerchantPrearbitrationDeclineSubmitted, CardDisputeVisaNetworkEventsCategoryMerchantPrearbitrationReceived, CardDisputeVisaNetworkEventsCategoryMerchantPrearbitrationTimedOut, CardDisputeVisaNetworkEventsCategoryRepresented, CardDisputeVisaNetworkEventsCategoryRepresentmentTimedOut, CardDisputeVisaNetworkEventsCategoryUserPrearbitrationAccepted, CardDisputeVisaNetworkEventsCategoryUserPrearbitrationDeclined, CardDisputeVisaNetworkEventsCategoryUserPrearbitrationSubmitted, CardDisputeVisaNetworkEventsCategoryUserPrearbitrationTimedOut, CardDisputeVisaNetworkEventsCategoryUserWithdrawalSubmitted:
		return true
	}
	return false
}

// A Card Dispute Chargeback Accepted Visa Network Event object. This field will be
// present in the JSON response if and only if `category` is equal to
// `chargeback_accepted`. Contains the details specific to a chargeback accepted
// Visa Card Dispute Network Event, which represents that a chargeback has been
// accepted by the merchant.
type CardDisputeVisaNetworkEventsChargebackAccepted struct {
	JSON cardDisputeVisaNetworkEventsChargebackAcceptedJSON `json:"-"`
}

// cardDisputeVisaNetworkEventsChargebackAcceptedJSON contains the JSON metadata
// for the struct [CardDisputeVisaNetworkEventsChargebackAccepted]
type cardDisputeVisaNetworkEventsChargebackAcceptedJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeVisaNetworkEventsChargebackAccepted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaNetworkEventsChargebackAcceptedJSON) RawJSON() string {
	return r.raw
}

// A Card Dispute Chargeback Submitted Visa Network Event object. This field will
// be present in the JSON response if and only if `category` is equal to
// `chargeback_submitted`. Contains the details specific to a chargeback submitted
// Visa Card Dispute Network Event, which represents that a chargeback has been
// submitted to the network.
type CardDisputeVisaNetworkEventsChargebackSubmitted struct {
	JSON cardDisputeVisaNetworkEventsChargebackSubmittedJSON `json:"-"`
}

// cardDisputeVisaNetworkEventsChargebackSubmittedJSON contains the JSON metadata
// for the struct [CardDisputeVisaNetworkEventsChargebackSubmitted]
type cardDisputeVisaNetworkEventsChargebackSubmittedJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeVisaNetworkEventsChargebackSubmitted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaNetworkEventsChargebackSubmittedJSON) RawJSON() string {
	return r.raw
}

// A Card Dispute Chargeback Timed Out Visa Network Event object. This field will
// be present in the JSON response if and only if `category` is equal to
// `chargeback_timed_out`. Contains the details specific to a chargeback timed out
// Visa Card Dispute Network Event, which represents that the chargeback has timed
// out in the user's favor.
type CardDisputeVisaNetworkEventsChargebackTimedOut struct {
	JSON cardDisputeVisaNetworkEventsChargebackTimedOutJSON `json:"-"`
}

// cardDisputeVisaNetworkEventsChargebackTimedOutJSON contains the JSON metadata
// for the struct [CardDisputeVisaNetworkEventsChargebackTimedOut]
type cardDisputeVisaNetworkEventsChargebackTimedOutJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeVisaNetworkEventsChargebackTimedOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaNetworkEventsChargebackTimedOutJSON) RawJSON() string {
	return r.raw
}

// A Card Dispute Merchant Pre-Arbitration Decline Submitted Visa Network Event
// object. This field will be present in the JSON response if and only if
// `category` is equal to `merchant_prearbitration_decline_submitted`. Contains the
// details specific to a merchant prearbitration decline submitted Visa Card
// Dispute Network Event, which represents that the user has declined the
// merchant's request for a prearbitration request decision in their favor.
type CardDisputeVisaNetworkEventsMerchantPrearbitrationDeclineSubmitted struct {
	JSON cardDisputeVisaNetworkEventsMerchantPrearbitrationDeclineSubmittedJSON `json:"-"`
}

// cardDisputeVisaNetworkEventsMerchantPrearbitrationDeclineSubmittedJSON contains
// the JSON metadata for the struct
// [CardDisputeVisaNetworkEventsMerchantPrearbitrationDeclineSubmitted]
type cardDisputeVisaNetworkEventsMerchantPrearbitrationDeclineSubmittedJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeVisaNetworkEventsMerchantPrearbitrationDeclineSubmitted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaNetworkEventsMerchantPrearbitrationDeclineSubmittedJSON) RawJSON() string {
	return r.raw
}

// A Card Dispute Merchant Pre-Arbitration Received Visa Network Event object. This
// field will be present in the JSON response if and only if `category` is equal to
// `merchant_prearbitration_received`. Contains the details specific to a merchant
// prearbitration received Visa Card Dispute Network Event, which represents that
// the merchant has issued a prearbitration request in the user's favor.
type CardDisputeVisaNetworkEventsMerchantPrearbitrationReceived struct {
	// Cardholder no longer disputes details. Present if and only if `reason` is
	// `cardholder_no_longer_disputes`.
	CardholderNoLongerDisputes CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCardholderNoLongerDisputes `json:"cardholder_no_longer_disputes,required,nullable"`
	// Compelling evidence details. Present if and only if `reason` is
	// `compelling_evidence`.
	CompellingEvidence CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidence `json:"compelling_evidence,required,nullable"`
	// Credit or reversal processed details. Present if and only if `reason` is
	// `credit_or_reversal_processed`.
	CreditOrReversalProcessed CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCreditOrReversalProcessed `json:"credit_or_reversal_processed,required,nullable"`
	// Delayed charge transaction details. Present if and only if `reason` is
	// `delayed_charge_transaction`.
	DelayedChargeTransaction CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedDelayedChargeTransaction `json:"delayed_charge_transaction,required,nullable"`
	// Evidence of imprint details. Present if and only if `reason` is
	// `evidence_of_imprint`.
	EvidenceOfImprint CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedEvidenceOfImprint `json:"evidence_of_imprint,required,nullable"`
	// Invalid dispute details. Present if and only if `reason` is `invalid_dispute`.
	InvalidDispute CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedInvalidDispute `json:"invalid_dispute,required,nullable"`
	// Non-fiat currency or non-fungible token received details. Present if and only if
	// `reason` is `non_fiat_currency_or_non_fungible_token_received`.
	NonFiatCurrencyOrNonFungibleTokenReceived CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedNonFiatCurrencyOrNonFungibleTokenReceived `json:"non_fiat_currency_or_non_fungible_token_received,required,nullable"`
	// Prior undisputed non-fraud transactions details. Present if and only if `reason`
	// is `prior_undisputed_non_fraud_transactions`.
	PriorUndisputedNonFraudTransactions CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedPriorUndisputedNonFraudTransactions `json:"prior_undisputed_non_fraud_transactions,required,nullable"`
	// The reason the merchant re-presented the dispute.
	Reason CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedReason `json:"reason,required"`
	JSON   cardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedJSON   `json:"-"`
}

// cardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedJSON contains the JSON
// metadata for the struct
// [CardDisputeVisaNetworkEventsMerchantPrearbitrationReceived]
type cardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedJSON struct {
	CardholderNoLongerDisputes                apijson.Field
	CompellingEvidence                        apijson.Field
	CreditOrReversalProcessed                 apijson.Field
	DelayedChargeTransaction                  apijson.Field
	EvidenceOfImprint                         apijson.Field
	InvalidDispute                            apijson.Field
	NonFiatCurrencyOrNonFungibleTokenReceived apijson.Field
	PriorUndisputedNonFraudTransactions       apijson.Field
	Reason                                    apijson.Field
	raw                                       string
	ExtraFields                               map[string]apijson.Field
}

func (r *CardDisputeVisaNetworkEventsMerchantPrearbitrationReceived) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedJSON) RawJSON() string {
	return r.raw
}

// Cardholder no longer disputes details. Present if and only if `reason` is
// `cardholder_no_longer_disputes`.
type CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCardholderNoLongerDisputes struct {
	// Explanation for why the merchant believes the cardholder no longer disputes the
	// transaction.
	Explanation string                                                                                   `json:"explanation,required,nullable"`
	JSON        cardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCardholderNoLongerDisputesJSON `json:"-"`
}

// cardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCardholderNoLongerDisputesJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCardholderNoLongerDisputes]
type cardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCardholderNoLongerDisputesJSON struct {
	Explanation apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCardholderNoLongerDisputes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCardholderNoLongerDisputesJSON) RawJSON() string {
	return r.raw
}

// Compelling evidence details. Present if and only if `reason` is
// `compelling_evidence`.
type CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidence struct {
	// The category of compelling evidence provided by the merchant.
	Category CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceCategory `json:"category,required"`
	// Explanation of the compelling evidence provided by the merchant.
	Explanation string                                                                           `json:"explanation,required,nullable"`
	JSON        cardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceJSON `json:"-"`
}

// cardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidence]
type cardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceJSON struct {
	Category    apijson.Field
	Explanation apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceJSON) RawJSON() string {
	return r.raw
}

// The category of compelling evidence provided by the merchant.
type CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceCategory string

const (
	CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceCategoryAuthorizedSigner                                                       CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceCategory = "authorized_signer"
	CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceCategoryDelivery                                                               CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceCategory = "delivery"
	CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceCategoryDeliveryAtPlaceOfEmployment                                            CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceCategory = "delivery_at_place_of_employment"
	CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceCategoryDigitalGoodsDownload                                                   CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceCategory = "digital_goods_download"
	CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceCategoryDynamicCurrencyConversionActivelyChosen                                CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceCategory = "dynamic_currency_conversion_actively_chosen"
	CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceCategoryFlightManifestAndPurchaseItinerary                                     CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceCategory = "flight_manifest_and_purchase_itinerary"
	CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceCategoryHouseholdMemberSigner                                                  CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceCategory = "household_member_signer"
	CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceCategoryLegitimateSpendAcrossPaymentTypesForSameMerchandise                    CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceCategory = "legitimate_spend_across_payment_types_for_same_merchandise"
	CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceCategoryMerchandiseUse                                                         CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceCategory = "merchandise_use"
	CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceCategoryPassengerTransportTicketUse                                            CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceCategory = "passenger_transport_ticket_use"
	CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceCategoryRecurringTransactionWithBindingContractOrPreviousUndisputedTransaction CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceCategory = "recurring_transaction_with_binding_contract_or_previous_undisputed_transaction"
	CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceCategorySignedDeliveryOrPickupForm                                             CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceCategory = "signed_delivery_or_pickup_form"
	CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceCategorySignedMailOrderPhoneOrderForm                                          CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceCategory = "signed_mail_order_phone_order_form"
	CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceCategoryTravelAndExpenseLoyaltyTransaction                                     CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceCategory = "travel_and_expense_loyalty_transaction"
	CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceCategoryTravelAndExpenseSubsequentPurchase                                     CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceCategory = "travel_and_expense_subsequent_purchase"
)

func (r CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceCategory) IsKnown() bool {
	switch r {
	case CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceCategoryAuthorizedSigner, CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceCategoryDelivery, CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceCategoryDeliveryAtPlaceOfEmployment, CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceCategoryDigitalGoodsDownload, CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceCategoryDynamicCurrencyConversionActivelyChosen, CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceCategoryFlightManifestAndPurchaseItinerary, CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceCategoryHouseholdMemberSigner, CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceCategoryLegitimateSpendAcrossPaymentTypesForSameMerchandise, CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceCategoryMerchandiseUse, CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceCategoryPassengerTransportTicketUse, CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceCategoryRecurringTransactionWithBindingContractOrPreviousUndisputedTransaction, CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceCategorySignedDeliveryOrPickupForm, CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceCategorySignedMailOrderPhoneOrderForm, CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceCategoryTravelAndExpenseLoyaltyTransaction, CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCompellingEvidenceCategoryTravelAndExpenseSubsequentPurchase:
		return true
	}
	return false
}

// Credit or reversal processed details. Present if and only if `reason` is
// `credit_or_reversal_processed`.
type CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCreditOrReversalProcessed struct {
	// The amount of the credit or reversal in the minor unit of its currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the credit or
	// reversal's currency.
	Currency string `json:"currency,required"`
	// Explanation for why the merchant believes the credit or reversal was processed.
	Explanation string `json:"explanation,required,nullable"`
	// The date the credit or reversal was processed.
	ProcessedAt time.Time                                                                               `json:"processed_at,required" format:"date"`
	JSON        cardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCreditOrReversalProcessedJSON `json:"-"`
}

// cardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCreditOrReversalProcessedJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCreditOrReversalProcessed]
type cardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCreditOrReversalProcessedJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	Explanation apijson.Field
	ProcessedAt apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCreditOrReversalProcessed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedCreditOrReversalProcessedJSON) RawJSON() string {
	return r.raw
}

// Delayed charge transaction details. Present if and only if `reason` is
// `delayed_charge_transaction`.
type CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedDelayedChargeTransaction struct {
	// Additional details about the delayed charge transaction.
	Explanation string                                                                                 `json:"explanation,required,nullable"`
	JSON        cardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedDelayedChargeTransactionJSON `json:"-"`
}

// cardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedDelayedChargeTransactionJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedDelayedChargeTransaction]
type cardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedDelayedChargeTransactionJSON struct {
	Explanation apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedDelayedChargeTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedDelayedChargeTransactionJSON) RawJSON() string {
	return r.raw
}

// Evidence of imprint details. Present if and only if `reason` is
// `evidence_of_imprint`.
type CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedEvidenceOfImprint struct {
	// Explanation of the evidence of imprint.
	Explanation string                                                                          `json:"explanation,required,nullable"`
	JSON        cardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedEvidenceOfImprintJSON `json:"-"`
}

// cardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedEvidenceOfImprintJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedEvidenceOfImprint]
type cardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedEvidenceOfImprintJSON struct {
	Explanation apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedEvidenceOfImprint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedEvidenceOfImprintJSON) RawJSON() string {
	return r.raw
}

// Invalid dispute details. Present if and only if `reason` is `invalid_dispute`.
type CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedInvalidDispute struct {
	// Explanation for why the dispute is considered invalid by the merchant.
	Explanation string `json:"explanation,required,nullable"`
	// The reason a merchant considers the dispute invalid.
	Reason CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedInvalidDisputeReason `json:"reason,required"`
	JSON   cardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedInvalidDisputeJSON   `json:"-"`
}

// cardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedInvalidDisputeJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedInvalidDispute]
type cardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedInvalidDisputeJSON struct {
	Explanation apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedInvalidDispute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedInvalidDisputeJSON) RawJSON() string {
	return r.raw
}

// The reason a merchant considers the dispute invalid.
type CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedInvalidDisputeReason string

const (
	CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedInvalidDisputeReasonOther                                  CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedInvalidDisputeReason = "other"
	CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedInvalidDisputeReasonSpecialAuthorizationProceduresFollowed CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedInvalidDisputeReason = "special_authorization_procedures_followed"
)

func (r CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedInvalidDisputeReason) IsKnown() bool {
	switch r {
	case CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedInvalidDisputeReasonOther, CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedInvalidDisputeReasonSpecialAuthorizationProceduresFollowed:
		return true
	}
	return false
}

// Non-fiat currency or non-fungible token received details. Present if and only if
// `reason` is `non_fiat_currency_or_non_fungible_token_received`.
type CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedNonFiatCurrencyOrNonFungibleTokenReceived struct {
	// Blockchain transaction hash.
	BlockchainTransactionHash string `json:"blockchain_transaction_hash,required"`
	// Destination wallet address.
	DestinationWalletAddress string `json:"destination_wallet_address,required"`
	// Prior approved transactions.
	PriorApprovedTransactions string                                                                                                  `json:"prior_approved_transactions,required,nullable"`
	JSON                      cardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedNonFiatCurrencyOrNonFungibleTokenReceivedJSON `json:"-"`
}

// cardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedNonFiatCurrencyOrNonFungibleTokenReceivedJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedNonFiatCurrencyOrNonFungibleTokenReceived]
type cardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedNonFiatCurrencyOrNonFungibleTokenReceivedJSON struct {
	BlockchainTransactionHash apijson.Field
	DestinationWalletAddress  apijson.Field
	PriorApprovedTransactions apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedNonFiatCurrencyOrNonFungibleTokenReceived) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedNonFiatCurrencyOrNonFungibleTokenReceivedJSON) RawJSON() string {
	return r.raw
}

// Prior undisputed non-fraud transactions details. Present if and only if `reason`
// is `prior_undisputed_non_fraud_transactions`.
type CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedPriorUndisputedNonFraudTransactions struct {
	// Explanation of the prior undisputed non-fraud transactions provided by the
	// merchant.
	Explanation string                                                                                            `json:"explanation,required,nullable"`
	JSON        cardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedPriorUndisputedNonFraudTransactionsJSON `json:"-"`
}

// cardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedPriorUndisputedNonFraudTransactionsJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedPriorUndisputedNonFraudTransactions]
type cardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedPriorUndisputedNonFraudTransactionsJSON struct {
	Explanation apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedPriorUndisputedNonFraudTransactions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedPriorUndisputedNonFraudTransactionsJSON) RawJSON() string {
	return r.raw
}

// The reason the merchant re-presented the dispute.
type CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedReason string

const (
	CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedReasonCardholderNoLongerDisputes                CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedReason = "cardholder_no_longer_disputes"
	CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedReasonCompellingEvidence                        CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedReason = "compelling_evidence"
	CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedReasonCreditOrReversalProcessed                 CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedReason = "credit_or_reversal_processed"
	CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedReasonDelayedChargeTransaction                  CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedReason = "delayed_charge_transaction"
	CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedReasonEvidenceOfImprint                         CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedReason = "evidence_of_imprint"
	CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedReasonInvalidDispute                            CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedReason = "invalid_dispute"
	CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedReasonNonFiatCurrencyOrNonFungibleTokenReceived CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedReason = "non_fiat_currency_or_non_fungible_token_received"
	CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedReasonPriorUndisputedNonFraudTransactions       CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedReason = "prior_undisputed_non_fraud_transactions"
)

func (r CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedReason) IsKnown() bool {
	switch r {
	case CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedReasonCardholderNoLongerDisputes, CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedReasonCompellingEvidence, CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedReasonCreditOrReversalProcessed, CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedReasonDelayedChargeTransaction, CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedReasonEvidenceOfImprint, CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedReasonInvalidDispute, CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedReasonNonFiatCurrencyOrNonFungibleTokenReceived, CardDisputeVisaNetworkEventsMerchantPrearbitrationReceivedReasonPriorUndisputedNonFraudTransactions:
		return true
	}
	return false
}

// A Card Dispute Merchant Pre-Arbitration Timed Out Visa Network Event object.
// This field will be present in the JSON response if and only if `category` is
// equal to `merchant_prearbitration_timed_out`. Contains the details specific to a
// merchant prearbitration timed out Visa Card Dispute Network Event, which
// represents that the user has timed out responding to the merchant's
// prearbitration request.
type CardDisputeVisaNetworkEventsMerchantPrearbitrationTimedOut struct {
	JSON cardDisputeVisaNetworkEventsMerchantPrearbitrationTimedOutJSON `json:"-"`
}

// cardDisputeVisaNetworkEventsMerchantPrearbitrationTimedOutJSON contains the JSON
// metadata for the struct
// [CardDisputeVisaNetworkEventsMerchantPrearbitrationTimedOut]
type cardDisputeVisaNetworkEventsMerchantPrearbitrationTimedOutJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeVisaNetworkEventsMerchantPrearbitrationTimedOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaNetworkEventsMerchantPrearbitrationTimedOutJSON) RawJSON() string {
	return r.raw
}

// A Card Dispute Re-presented Visa Network Event object. This field will be
// present in the JSON response if and only if `category` is equal to
// `represented`. Contains the details specific to a re-presented Visa Card Dispute
// Network Event, which represents that the merchant has declined the user's
// chargeback and has re-presented the payment.
type CardDisputeVisaNetworkEventsRepresented struct {
	// Cardholder no longer disputes details. Present if and only if `reason` is
	// `cardholder_no_longer_disputes`.
	CardholderNoLongerDisputes CardDisputeVisaNetworkEventsRepresentedCardholderNoLongerDisputes `json:"cardholder_no_longer_disputes,required,nullable"`
	// Credit or reversal processed details. Present if and only if `reason` is
	// `credit_or_reversal_processed`.
	CreditOrReversalProcessed CardDisputeVisaNetworkEventsRepresentedCreditOrReversalProcessed `json:"credit_or_reversal_processed,required,nullable"`
	// Invalid dispute details. Present if and only if `reason` is `invalid_dispute`.
	InvalidDispute CardDisputeVisaNetworkEventsRepresentedInvalidDispute `json:"invalid_dispute,required,nullable"`
	// Non-fiat currency or non-fungible token as described details. Present if and
	// only if `reason` is `non_fiat_currency_or_non_fungible_token_as_described`.
	NonFiatCurrencyOrNonFungibleTokenAsDescribed CardDisputeVisaNetworkEventsRepresentedNonFiatCurrencyOrNonFungibleTokenAsDescribed `json:"non_fiat_currency_or_non_fungible_token_as_described,required,nullable"`
	// Non-fiat currency or non-fungible token received details. Present if and only if
	// `reason` is `non_fiat_currency_or_non_fungible_token_received`.
	NonFiatCurrencyOrNonFungibleTokenReceived CardDisputeVisaNetworkEventsRepresentedNonFiatCurrencyOrNonFungibleTokenReceived `json:"non_fiat_currency_or_non_fungible_token_received,required,nullable"`
	// Proof of cash disbursement details. Present if and only if `reason` is
	// `proof_of_cash_disbursement`.
	ProofOfCashDisbursement CardDisputeVisaNetworkEventsRepresentedProofOfCashDisbursement `json:"proof_of_cash_disbursement,required,nullable"`
	// The reason the merchant re-presented the dispute.
	Reason CardDisputeVisaNetworkEventsRepresentedReason `json:"reason,required"`
	// Reversal issued by merchant details. Present if and only if `reason` is
	// `reversal_issued`.
	ReversalIssued CardDisputeVisaNetworkEventsRepresentedReversalIssued `json:"reversal_issued,required,nullable"`
	JSON           cardDisputeVisaNetworkEventsRepresentedJSON           `json:"-"`
}

// cardDisputeVisaNetworkEventsRepresentedJSON contains the JSON metadata for the
// struct [CardDisputeVisaNetworkEventsRepresented]
type cardDisputeVisaNetworkEventsRepresentedJSON struct {
	CardholderNoLongerDisputes                   apijson.Field
	CreditOrReversalProcessed                    apijson.Field
	InvalidDispute                               apijson.Field
	NonFiatCurrencyOrNonFungibleTokenAsDescribed apijson.Field
	NonFiatCurrencyOrNonFungibleTokenReceived    apijson.Field
	ProofOfCashDisbursement                      apijson.Field
	Reason                                       apijson.Field
	ReversalIssued                               apijson.Field
	raw                                          string
	ExtraFields                                  map[string]apijson.Field
}

func (r *CardDisputeVisaNetworkEventsRepresented) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaNetworkEventsRepresentedJSON) RawJSON() string {
	return r.raw
}

// Cardholder no longer disputes details. Present if and only if `reason` is
// `cardholder_no_longer_disputes`.
type CardDisputeVisaNetworkEventsRepresentedCardholderNoLongerDisputes struct {
	// Explanation for why the merchant believes the cardholder no longer disputes the
	// transaction.
	Explanation string                                                                `json:"explanation,required,nullable"`
	JSON        cardDisputeVisaNetworkEventsRepresentedCardholderNoLongerDisputesJSON `json:"-"`
}

// cardDisputeVisaNetworkEventsRepresentedCardholderNoLongerDisputesJSON contains
// the JSON metadata for the struct
// [CardDisputeVisaNetworkEventsRepresentedCardholderNoLongerDisputes]
type cardDisputeVisaNetworkEventsRepresentedCardholderNoLongerDisputesJSON struct {
	Explanation apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeVisaNetworkEventsRepresentedCardholderNoLongerDisputes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaNetworkEventsRepresentedCardholderNoLongerDisputesJSON) RawJSON() string {
	return r.raw
}

// Credit or reversal processed details. Present if and only if `reason` is
// `credit_or_reversal_processed`.
type CardDisputeVisaNetworkEventsRepresentedCreditOrReversalProcessed struct {
	// The amount of the credit or reversal in the minor unit of its currency. For
	// dollars, for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the credit or
	// reversal's currency.
	Currency string `json:"currency,required"`
	// Explanation for why the merchant believes the credit or reversal was processed.
	Explanation string `json:"explanation,required,nullable"`
	// The date the credit or reversal was processed.
	ProcessedAt time.Time                                                            `json:"processed_at,required" format:"date"`
	JSON        cardDisputeVisaNetworkEventsRepresentedCreditOrReversalProcessedJSON `json:"-"`
}

// cardDisputeVisaNetworkEventsRepresentedCreditOrReversalProcessedJSON contains
// the JSON metadata for the struct
// [CardDisputeVisaNetworkEventsRepresentedCreditOrReversalProcessed]
type cardDisputeVisaNetworkEventsRepresentedCreditOrReversalProcessedJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	Explanation apijson.Field
	ProcessedAt apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeVisaNetworkEventsRepresentedCreditOrReversalProcessed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaNetworkEventsRepresentedCreditOrReversalProcessedJSON) RawJSON() string {
	return r.raw
}

// Invalid dispute details. Present if and only if `reason` is `invalid_dispute`.
type CardDisputeVisaNetworkEventsRepresentedInvalidDispute struct {
	// Explanation for why the dispute is considered invalid by the merchant.
	Explanation string `json:"explanation,required,nullable"`
	// The reason a merchant considers the dispute invalid.
	Reason CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReason `json:"reason,required"`
	JSON   cardDisputeVisaNetworkEventsRepresentedInvalidDisputeJSON   `json:"-"`
}

// cardDisputeVisaNetworkEventsRepresentedInvalidDisputeJSON contains the JSON
// metadata for the struct [CardDisputeVisaNetworkEventsRepresentedInvalidDispute]
type cardDisputeVisaNetworkEventsRepresentedInvalidDisputeJSON struct {
	Explanation apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeVisaNetworkEventsRepresentedInvalidDispute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaNetworkEventsRepresentedInvalidDisputeJSON) RawJSON() string {
	return r.raw
}

// The reason a merchant considers the dispute invalid.
type CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReason string

const (
	CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonAutomaticTellerMachineTransactionProofProvided              CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReason = "automatic_teller_machine_transaction_proof_provided"
	CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonBalanceOfPartialPrepaymentNotPaid                           CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReason = "balance_of_partial_prepayment_not_paid"
	CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonCardholderCanceledBeforeExpectedMerchandiseReceiptDate      CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReason = "cardholder_canceled_before_expected_merchandise_receipt_date"
	CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonCardholderCanceledBeforeExpectedServicesReceiptDate         CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReason = "cardholder_canceled_before_expected_services_receipt_date"
	CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonCardholderCanceledDifferentDate                             CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReason = "cardholder_canceled_different_date"
	CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonCardholderDidNotCancelAccordingToPolicy                     CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReason = "cardholder_did_not_cancel_according_to_policy"
	CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonCardholderReceivedMerchandise                               CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReason = "cardholder_received_merchandise"
	CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonCountryCodeCorrect                                          CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReason = "country_code_correct"
	CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonCreditProcessedCorrectly                                    CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReason = "credit_processed_correctly"
	CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonCurrencyCorrect                                             CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReason = "currency_correct"
	CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonDisputeIsForQuality                                         CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReason = "dispute_is_for_quality"
	CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonDisputeIsForVisaCashBackTransactionPortion                  CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReason = "dispute_is_for_visa_cash_back_transaction_portion"
	CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonDisputedAmountIsValueAddedTax                               CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReason = "disputed_amount_is_value_added_tax"
	CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonDisputedAmountIsValueAddedTaxNoCreditReceiptProvided        CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReason = "disputed_amount_is_value_added_tax_no_credit_receipt_provided"
	CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonLimitedReturnOrCancellationPolicyProperlyDisclosed          CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReason = "limited_return_or_cancellation_policy_properly_disclosed"
	CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonMerchandiseHeldAtCardholderCustomsAgency                    CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReason = "merchandise_held_at_cardholder_customs_agency"
	CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonMerchandiseMatchesDescription                               CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReason = "merchandise_matches_description"
	CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonMerchandiseNotCounterfeit                                   CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReason = "merchandise_not_counterfeit"
	CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonMerchandiseNotDamaged                                       CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReason = "merchandise_not_damaged"
	CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonMerchandiseNotDefective                                     CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReason = "merchandise_not_defective"
	CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonMerchandiseProvidedPriorToCancellationDate                  CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReason = "merchandise_provided_prior_to_cancellation_date"
	CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonMerchandiseQualityMatchesDescription                        CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReason = "merchandise_quality_matches_description"
	CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonMerchandiseReturnNotAttempted                               CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReason = "merchandise_return_not_attempted"
	CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonMerchantNotNotifiedOfClosedAccount                          CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReason = "merchant_not_notified_of_closed_account"
	CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonNameOnFlightManifestMatchesPurchase                         CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReason = "name_on_flight_manifest_matches_purchase"
	CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonNoCreditReceiptProvided                                     CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReason = "no_credit_receipt_provided"
	CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonOther                                                       CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReason = "other"
	CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonProcessingErrorIncorrect                                    CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReason = "processing_error_incorrect"
	CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonReturnedMechandiseHeldAtCustomsAgencyOutsideMerchantCountry CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReason = "returned_mechandise_held_at_customs_agency_outside_merchant_country"
	CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonServicesMatchDescription                                    CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReason = "services_match_description"
	CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonServicesProvidedPriorToCancellationDate                     CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReason = "services_provided_prior_to_cancellation_date"
	CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonServicesUsedAfterCancellationDate                           CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReason = "services_used_after_cancellation_date"
	CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonTermsOfServiceNotMisrepresented                             CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReason = "terms_of_service_not_misrepresented"
	CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonTransactionCodeCorrect                                      CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReason = "transaction_code_correct"
)

func (r CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReason) IsKnown() bool {
	switch r {
	case CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonAutomaticTellerMachineTransactionProofProvided, CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonBalanceOfPartialPrepaymentNotPaid, CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonCardholderCanceledBeforeExpectedMerchandiseReceiptDate, CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonCardholderCanceledBeforeExpectedServicesReceiptDate, CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonCardholderCanceledDifferentDate, CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonCardholderDidNotCancelAccordingToPolicy, CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonCardholderReceivedMerchandise, CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonCountryCodeCorrect, CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonCreditProcessedCorrectly, CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonCurrencyCorrect, CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonDisputeIsForQuality, CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonDisputeIsForVisaCashBackTransactionPortion, CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonDisputedAmountIsValueAddedTax, CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonDisputedAmountIsValueAddedTaxNoCreditReceiptProvided, CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonLimitedReturnOrCancellationPolicyProperlyDisclosed, CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonMerchandiseHeldAtCardholderCustomsAgency, CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonMerchandiseMatchesDescription, CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonMerchandiseNotCounterfeit, CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonMerchandiseNotDamaged, CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonMerchandiseNotDefective, CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonMerchandiseProvidedPriorToCancellationDate, CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonMerchandiseQualityMatchesDescription, CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonMerchandiseReturnNotAttempted, CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonMerchantNotNotifiedOfClosedAccount, CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonNameOnFlightManifestMatchesPurchase, CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonNoCreditReceiptProvided, CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonOther, CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonProcessingErrorIncorrect, CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonReturnedMechandiseHeldAtCustomsAgencyOutsideMerchantCountry, CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonServicesMatchDescription, CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonServicesProvidedPriorToCancellationDate, CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonServicesUsedAfterCancellationDate, CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonTermsOfServiceNotMisrepresented, CardDisputeVisaNetworkEventsRepresentedInvalidDisputeReasonTransactionCodeCorrect:
		return true
	}
	return false
}

// Non-fiat currency or non-fungible token as described details. Present if and
// only if `reason` is `non_fiat_currency_or_non_fungible_token_as_described`.
type CardDisputeVisaNetworkEventsRepresentedNonFiatCurrencyOrNonFungibleTokenAsDescribed struct {
	JSON cardDisputeVisaNetworkEventsRepresentedNonFiatCurrencyOrNonFungibleTokenAsDescribedJSON `json:"-"`
}

// cardDisputeVisaNetworkEventsRepresentedNonFiatCurrencyOrNonFungibleTokenAsDescribedJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaNetworkEventsRepresentedNonFiatCurrencyOrNonFungibleTokenAsDescribed]
type cardDisputeVisaNetworkEventsRepresentedNonFiatCurrencyOrNonFungibleTokenAsDescribedJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeVisaNetworkEventsRepresentedNonFiatCurrencyOrNonFungibleTokenAsDescribed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaNetworkEventsRepresentedNonFiatCurrencyOrNonFungibleTokenAsDescribedJSON) RawJSON() string {
	return r.raw
}

// Non-fiat currency or non-fungible token received details. Present if and only if
// `reason` is `non_fiat_currency_or_non_fungible_token_received`.
type CardDisputeVisaNetworkEventsRepresentedNonFiatCurrencyOrNonFungibleTokenReceived struct {
	// Blockchain transaction hash.
	BlockchainTransactionHash string `json:"blockchain_transaction_hash,required"`
	// Destination wallet address.
	DestinationWalletAddress string `json:"destination_wallet_address,required"`
	// Prior approved transactions.
	PriorApprovedTransactions string                                                                               `json:"prior_approved_transactions,required,nullable"`
	JSON                      cardDisputeVisaNetworkEventsRepresentedNonFiatCurrencyOrNonFungibleTokenReceivedJSON `json:"-"`
}

// cardDisputeVisaNetworkEventsRepresentedNonFiatCurrencyOrNonFungibleTokenReceivedJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaNetworkEventsRepresentedNonFiatCurrencyOrNonFungibleTokenReceived]
type cardDisputeVisaNetworkEventsRepresentedNonFiatCurrencyOrNonFungibleTokenReceivedJSON struct {
	BlockchainTransactionHash apijson.Field
	DestinationWalletAddress  apijson.Field
	PriorApprovedTransactions apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *CardDisputeVisaNetworkEventsRepresentedNonFiatCurrencyOrNonFungibleTokenReceived) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaNetworkEventsRepresentedNonFiatCurrencyOrNonFungibleTokenReceivedJSON) RawJSON() string {
	return r.raw
}

// Proof of cash disbursement details. Present if and only if `reason` is
// `proof_of_cash_disbursement`.
type CardDisputeVisaNetworkEventsRepresentedProofOfCashDisbursement struct {
	// Explanation for why the merchant believes the evidence provides proof of cash
	// disbursement.
	Explanation string                                                             `json:"explanation,required,nullable"`
	JSON        cardDisputeVisaNetworkEventsRepresentedProofOfCashDisbursementJSON `json:"-"`
}

// cardDisputeVisaNetworkEventsRepresentedProofOfCashDisbursementJSON contains the
// JSON metadata for the struct
// [CardDisputeVisaNetworkEventsRepresentedProofOfCashDisbursement]
type cardDisputeVisaNetworkEventsRepresentedProofOfCashDisbursementJSON struct {
	Explanation apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeVisaNetworkEventsRepresentedProofOfCashDisbursement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaNetworkEventsRepresentedProofOfCashDisbursementJSON) RawJSON() string {
	return r.raw
}

// The reason the merchant re-presented the dispute.
type CardDisputeVisaNetworkEventsRepresentedReason string

const (
	CardDisputeVisaNetworkEventsRepresentedReasonCardholderNoLongerDisputes                   CardDisputeVisaNetworkEventsRepresentedReason = "cardholder_no_longer_disputes"
	CardDisputeVisaNetworkEventsRepresentedReasonCreditOrReversalProcessed                    CardDisputeVisaNetworkEventsRepresentedReason = "credit_or_reversal_processed"
	CardDisputeVisaNetworkEventsRepresentedReasonInvalidDispute                               CardDisputeVisaNetworkEventsRepresentedReason = "invalid_dispute"
	CardDisputeVisaNetworkEventsRepresentedReasonNonFiatCurrencyOrNonFungibleTokenAsDescribed CardDisputeVisaNetworkEventsRepresentedReason = "non_fiat_currency_or_non_fungible_token_as_described"
	CardDisputeVisaNetworkEventsRepresentedReasonNonFiatCurrencyOrNonFungibleTokenReceived    CardDisputeVisaNetworkEventsRepresentedReason = "non_fiat_currency_or_non_fungible_token_received"
	CardDisputeVisaNetworkEventsRepresentedReasonProofOfCashDisbursement                      CardDisputeVisaNetworkEventsRepresentedReason = "proof_of_cash_disbursement"
	CardDisputeVisaNetworkEventsRepresentedReasonReversalIssued                               CardDisputeVisaNetworkEventsRepresentedReason = "reversal_issued"
)

func (r CardDisputeVisaNetworkEventsRepresentedReason) IsKnown() bool {
	switch r {
	case CardDisputeVisaNetworkEventsRepresentedReasonCardholderNoLongerDisputes, CardDisputeVisaNetworkEventsRepresentedReasonCreditOrReversalProcessed, CardDisputeVisaNetworkEventsRepresentedReasonInvalidDispute, CardDisputeVisaNetworkEventsRepresentedReasonNonFiatCurrencyOrNonFungibleTokenAsDescribed, CardDisputeVisaNetworkEventsRepresentedReasonNonFiatCurrencyOrNonFungibleTokenReceived, CardDisputeVisaNetworkEventsRepresentedReasonProofOfCashDisbursement, CardDisputeVisaNetworkEventsRepresentedReasonReversalIssued:
		return true
	}
	return false
}

// Reversal issued by merchant details. Present if and only if `reason` is
// `reversal_issued`.
type CardDisputeVisaNetworkEventsRepresentedReversalIssued struct {
	// Explanation of the reversal issued by the merchant.
	Explanation string                                                    `json:"explanation,required,nullable"`
	JSON        cardDisputeVisaNetworkEventsRepresentedReversalIssuedJSON `json:"-"`
}

// cardDisputeVisaNetworkEventsRepresentedReversalIssuedJSON contains the JSON
// metadata for the struct [CardDisputeVisaNetworkEventsRepresentedReversalIssued]
type cardDisputeVisaNetworkEventsRepresentedReversalIssuedJSON struct {
	Explanation apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeVisaNetworkEventsRepresentedReversalIssued) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaNetworkEventsRepresentedReversalIssuedJSON) RawJSON() string {
	return r.raw
}

// A Card Dispute Re-presentment Timed Out Visa Network Event object. This field
// will be present in the JSON response if and only if `category` is equal to
// `representment_timed_out`. Contains the details specific to a re-presentment
// time-out Visa Card Dispute Network Event, which represents that the user did not
// respond to the re-presentment by the merchant within the time limit.
type CardDisputeVisaNetworkEventsRepresentmentTimedOut struct {
	JSON cardDisputeVisaNetworkEventsRepresentmentTimedOutJSON `json:"-"`
}

// cardDisputeVisaNetworkEventsRepresentmentTimedOutJSON contains the JSON metadata
// for the struct [CardDisputeVisaNetworkEventsRepresentmentTimedOut]
type cardDisputeVisaNetworkEventsRepresentmentTimedOutJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeVisaNetworkEventsRepresentmentTimedOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaNetworkEventsRepresentmentTimedOutJSON) RawJSON() string {
	return r.raw
}

// A Card Dispute User Pre-Arbitration Accepted Visa Network Event object. This
// field will be present in the JSON response if and only if `category` is equal to
// `user_prearbitration_accepted`. Contains the details specific to a user
// prearbitration accepted Visa Card Dispute Network Event, which represents that
// the merchant has accepted the user's prearbitration request in the user's favor.
type CardDisputeVisaNetworkEventsUserPrearbitrationAccepted struct {
	JSON cardDisputeVisaNetworkEventsUserPrearbitrationAcceptedJSON `json:"-"`
}

// cardDisputeVisaNetworkEventsUserPrearbitrationAcceptedJSON contains the JSON
// metadata for the struct [CardDisputeVisaNetworkEventsUserPrearbitrationAccepted]
type cardDisputeVisaNetworkEventsUserPrearbitrationAcceptedJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeVisaNetworkEventsUserPrearbitrationAccepted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaNetworkEventsUserPrearbitrationAcceptedJSON) RawJSON() string {
	return r.raw
}

// A Card Dispute User Pre-Arbitration Declined Visa Network Event object. This
// field will be present in the JSON response if and only if `category` is equal to
// `user_prearbitration_declined`. Contains the details specific to a user
// prearbitration declined Visa Card Dispute Network Event, which represents that
// the merchant has declined the user's prearbitration request.
type CardDisputeVisaNetworkEventsUserPrearbitrationDeclined struct {
	JSON cardDisputeVisaNetworkEventsUserPrearbitrationDeclinedJSON `json:"-"`
}

// cardDisputeVisaNetworkEventsUserPrearbitrationDeclinedJSON contains the JSON
// metadata for the struct [CardDisputeVisaNetworkEventsUserPrearbitrationDeclined]
type cardDisputeVisaNetworkEventsUserPrearbitrationDeclinedJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeVisaNetworkEventsUserPrearbitrationDeclined) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaNetworkEventsUserPrearbitrationDeclinedJSON) RawJSON() string {
	return r.raw
}

// A Card Dispute User Pre-Arbitration Submitted Visa Network Event object. This
// field will be present in the JSON response if and only if `category` is equal to
// `user_prearbitration_submitted`. Contains the details specific to a user
// prearbitration submitted Visa Card Dispute Network Event, which represents that
// the user's request for prearbitration has been submitted to the network.
type CardDisputeVisaNetworkEventsUserPrearbitrationSubmitted struct {
	JSON cardDisputeVisaNetworkEventsUserPrearbitrationSubmittedJSON `json:"-"`
}

// cardDisputeVisaNetworkEventsUserPrearbitrationSubmittedJSON contains the JSON
// metadata for the struct
// [CardDisputeVisaNetworkEventsUserPrearbitrationSubmitted]
type cardDisputeVisaNetworkEventsUserPrearbitrationSubmittedJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeVisaNetworkEventsUserPrearbitrationSubmitted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaNetworkEventsUserPrearbitrationSubmittedJSON) RawJSON() string {
	return r.raw
}

// A Card Dispute User Pre-Arbitration Timed Out Visa Network Event object. This
// field will be present in the JSON response if and only if `category` is equal to
// `user_prearbitration_timed_out`. Contains the details specific to a user
// prearbitration timed out Visa Card Dispute Network Event, which represents that
// the merchant has timed out responding to the user's prearbitration request.
type CardDisputeVisaNetworkEventsUserPrearbitrationTimedOut struct {
	JSON cardDisputeVisaNetworkEventsUserPrearbitrationTimedOutJSON `json:"-"`
}

// cardDisputeVisaNetworkEventsUserPrearbitrationTimedOutJSON contains the JSON
// metadata for the struct [CardDisputeVisaNetworkEventsUserPrearbitrationTimedOut]
type cardDisputeVisaNetworkEventsUserPrearbitrationTimedOutJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeVisaNetworkEventsUserPrearbitrationTimedOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaNetworkEventsUserPrearbitrationTimedOutJSON) RawJSON() string {
	return r.raw
}

// A Card Dispute User Withdrawal Submitted Visa Network Event object. This field
// will be present in the JSON response if and only if `category` is equal to
// `user_withdrawal_submitted`. Contains the details specific to a user withdrawal
// submitted Visa Card Dispute Network Event, which represents that the user's
// request to withdraw the dispute has been submitted to the network.
type CardDisputeVisaNetworkEventsUserWithdrawalSubmitted struct {
	JSON cardDisputeVisaNetworkEventsUserWithdrawalSubmittedJSON `json:"-"`
}

// cardDisputeVisaNetworkEventsUserWithdrawalSubmittedJSON contains the JSON
// metadata for the struct [CardDisputeVisaNetworkEventsUserWithdrawalSubmitted]
type cardDisputeVisaNetworkEventsUserWithdrawalSubmittedJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeVisaNetworkEventsUserWithdrawalSubmitted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaNetworkEventsUserWithdrawalSubmittedJSON) RawJSON() string {
	return r.raw
}

// The category of the currently required user submission if the user wishes to
// proceed with the dispute. Present if and only if status is
// `user_submission_required`. Otherwise, this will be `nil`.
type CardDisputeVisaRequiredUserSubmissionCategory string

const (
	CardDisputeVisaRequiredUserSubmissionCategoryChargeback                    CardDisputeVisaRequiredUserSubmissionCategory = "chargeback"
	CardDisputeVisaRequiredUserSubmissionCategoryMerchantPrearbitrationDecline CardDisputeVisaRequiredUserSubmissionCategory = "merchant_prearbitration_decline"
	CardDisputeVisaRequiredUserSubmissionCategoryUserPrearbitration            CardDisputeVisaRequiredUserSubmissionCategory = "user_prearbitration"
)

func (r CardDisputeVisaRequiredUserSubmissionCategory) IsKnown() bool {
	switch r {
	case CardDisputeVisaRequiredUserSubmissionCategoryChargeback, CardDisputeVisaRequiredUserSubmissionCategoryMerchantPrearbitrationDecline, CardDisputeVisaRequiredUserSubmissionCategoryUserPrearbitration:
		return true
	}
	return false
}

type CardDisputeVisaUserSubmission struct {
	// The date and time at which the Visa Card Dispute User Submission was reviewed
	// and accepted.
	AcceptedAt time.Time `json:"accepted_at,required,nullable" format:"date-time"`
	// The amount of the dispute if it is different from the amount of a prior user
	// submission or the disputed transaction.
	Amount int64 `json:"amount,required,nullable"`
	// The files attached to the Visa Card Dispute User Submission.
	AttachmentFiles []CardDisputeVisaUserSubmissionsAttachmentFile `json:"attachment_files,required"`
	// The category of the user submission. We may add additional possible values for
	// this enum over time; your application should be able to handle such additions
	// gracefully.
	Category CardDisputeVisaUserSubmissionsCategory `json:"category,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Visa Card Dispute User Submission was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The free-form explanation provided to Increase to provide more context for the
	// user submission. This field is not sent directly to the card networks.
	Explanation string `json:"explanation,required,nullable"`
	// The date and time at which Increase requested further information from the user
	// for the Visa Card Dispute.
	FurtherInformationRequestedAt time.Time `json:"further_information_requested_at,required,nullable" format:"date-time"`
	// The reason for Increase requesting further information from the user for the
	// Visa Card Dispute.
	FurtherInformationRequestedReason string `json:"further_information_requested_reason,required,nullable"`
	// The status of the Visa Card Dispute User Submission.
	Status CardDisputeVisaUserSubmissionsStatus `json:"status,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Visa Card Dispute User Submission was updated.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// A Visa Card Dispute Chargeback User Submission Chargeback Details object. This
	// field will be present in the JSON response if and only if `category` is equal to
	// `chargeback`. Contains the details specific to a Visa chargeback User Submission
	// for a Card Dispute.
	Chargeback CardDisputeVisaUserSubmissionsChargeback `json:"chargeback,nullable"`
	// A Visa Card Dispute Merchant Pre-Arbitration Decline User Submission object.
	// This field will be present in the JSON response if and only if `category` is
	// equal to `merchant_prearbitration_decline`. Contains the details specific to a
	// merchant prearbitration decline Visa Card Dispute User Submission.
	MerchantPrearbitrationDecline CardDisputeVisaUserSubmissionsMerchantPrearbitrationDecline `json:"merchant_prearbitration_decline,nullable"`
	// A Visa Card Dispute User-Initiated Pre-Arbitration User Submission object. This
	// field will be present in the JSON response if and only if `category` is equal to
	// `user_prearbitration`. Contains the details specific to a user-initiated
	// pre-arbitration Visa Card Dispute User Submission.
	UserPrearbitration CardDisputeVisaUserSubmissionsUserPrearbitration `json:"user_prearbitration,nullable"`
	JSON               cardDisputeVisaUserSubmissionJSON                `json:"-"`
}

// cardDisputeVisaUserSubmissionJSON contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmission]
type cardDisputeVisaUserSubmissionJSON struct {
	AcceptedAt                        apijson.Field
	Amount                            apijson.Field
	AttachmentFiles                   apijson.Field
	Category                          apijson.Field
	CreatedAt                         apijson.Field
	Explanation                       apijson.Field
	FurtherInformationRequestedAt     apijson.Field
	FurtherInformationRequestedReason apijson.Field
	Status                            apijson.Field
	UpdatedAt                         apijson.Field
	Chargeback                        apijson.Field
	MerchantPrearbitrationDecline     apijson.Field
	UserPrearbitration                apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmission) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionJSON) RawJSON() string {
	return r.raw
}

type CardDisputeVisaUserSubmissionsAttachmentFile struct {
	// The ID of the file attached to the Card Dispute.
	FileID string                                           `json:"file_id,required"`
	JSON   cardDisputeVisaUserSubmissionsAttachmentFileJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsAttachmentFileJSON contains the JSON metadata for
// the struct [CardDisputeVisaUserSubmissionsAttachmentFile]
type cardDisputeVisaUserSubmissionsAttachmentFileJSON struct {
	FileID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsAttachmentFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsAttachmentFileJSON) RawJSON() string {
	return r.raw
}

// The category of the user submission. We may add additional possible values for
// this enum over time; your application should be able to handle such additions
// gracefully.
type CardDisputeVisaUserSubmissionsCategory string

const (
	CardDisputeVisaUserSubmissionsCategoryChargeback                    CardDisputeVisaUserSubmissionsCategory = "chargeback"
	CardDisputeVisaUserSubmissionsCategoryMerchantPrearbitrationDecline CardDisputeVisaUserSubmissionsCategory = "merchant_prearbitration_decline"
	CardDisputeVisaUserSubmissionsCategoryUserPrearbitration            CardDisputeVisaUserSubmissionsCategory = "user_prearbitration"
)

func (r CardDisputeVisaUserSubmissionsCategory) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsCategoryChargeback, CardDisputeVisaUserSubmissionsCategoryMerchantPrearbitrationDecline, CardDisputeVisaUserSubmissionsCategoryUserPrearbitration:
		return true
	}
	return false
}

// The status of the Visa Card Dispute User Submission.
type CardDisputeVisaUserSubmissionsStatus string

const (
	CardDisputeVisaUserSubmissionsStatusAbandoned                   CardDisputeVisaUserSubmissionsStatus = "abandoned"
	CardDisputeVisaUserSubmissionsStatusAccepted                    CardDisputeVisaUserSubmissionsStatus = "accepted"
	CardDisputeVisaUserSubmissionsStatusFurtherInformationRequested CardDisputeVisaUserSubmissionsStatus = "further_information_requested"
	CardDisputeVisaUserSubmissionsStatusPendingReviewing            CardDisputeVisaUserSubmissionsStatus = "pending_reviewing"
)

func (r CardDisputeVisaUserSubmissionsStatus) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsStatusAbandoned, CardDisputeVisaUserSubmissionsStatusAccepted, CardDisputeVisaUserSubmissionsStatusFurtherInformationRequested, CardDisputeVisaUserSubmissionsStatusPendingReviewing:
		return true
	}
	return false
}

// A Visa Card Dispute Chargeback User Submission Chargeback Details object. This
// field will be present in the JSON response if and only if `category` is equal to
// `chargeback`. Contains the details specific to a Visa chargeback User Submission
// for a Card Dispute.
type CardDisputeVisaUserSubmissionsChargeback struct {
	// Authorization. Present if and only if `category` is `authorization`.
	Authorization CardDisputeVisaUserSubmissionsChargebackAuthorization `json:"authorization,required,nullable"`
	// Category.
	Category CardDisputeVisaUserSubmissionsChargebackCategory `json:"category,required"`
	// Canceled merchandise. Present if and only if `category` is
	// `consumer_canceled_merchandise`.
	ConsumerCanceledMerchandise CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandise `json:"consumer_canceled_merchandise,required,nullable"`
	// Canceled recurring transaction. Present if and only if `category` is
	// `consumer_canceled_recurring_transaction`.
	ConsumerCanceledRecurringTransaction CardDisputeVisaUserSubmissionsChargebackConsumerCanceledRecurringTransaction `json:"consumer_canceled_recurring_transaction,required,nullable"`
	// Canceled services. Present if and only if `category` is
	// `consumer_canceled_services`.
	ConsumerCanceledServices CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServices `json:"consumer_canceled_services,required,nullable"`
	// Counterfeit merchandise. Present if and only if `category` is
	// `consumer_counterfeit_merchandise`.
	ConsumerCounterfeitMerchandise CardDisputeVisaUserSubmissionsChargebackConsumerCounterfeitMerchandise `json:"consumer_counterfeit_merchandise,required,nullable"`
	// Credit not processed. Present if and only if `category` is
	// `consumer_credit_not_processed`.
	ConsumerCreditNotProcessed CardDisputeVisaUserSubmissionsChargebackConsumerCreditNotProcessed `json:"consumer_credit_not_processed,required,nullable"`
	// Damaged or defective merchandise. Present if and only if `category` is
	// `consumer_damaged_or_defective_merchandise`.
	ConsumerDamagedOrDefectiveMerchandise CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandise `json:"consumer_damaged_or_defective_merchandise,required,nullable"`
	// Merchandise misrepresentation. Present if and only if `category` is
	// `consumer_merchandise_misrepresentation`.
	ConsumerMerchandiseMisrepresentation CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentation `json:"consumer_merchandise_misrepresentation,required,nullable"`
	// Merchandise not as described. Present if and only if `category` is
	// `consumer_merchandise_not_as_described`.
	ConsumerMerchandiseNotAsDescribed CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribed `json:"consumer_merchandise_not_as_described,required,nullable"`
	// Merchandise not received. Present if and only if `category` is
	// `consumer_merchandise_not_received`.
	ConsumerMerchandiseNotReceived CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceived `json:"consumer_merchandise_not_received,required,nullable"`
	// Non-receipt of cash. Present if and only if `category` is
	// `consumer_non_receipt_of_cash`.
	ConsumerNonReceiptOfCash CardDisputeVisaUserSubmissionsChargebackConsumerNonReceiptOfCash `json:"consumer_non_receipt_of_cash,required,nullable"`
	// Original Credit Transaction (OCT) not accepted. Present if and only if
	// `category` is `consumer_original_credit_transaction_not_accepted`.
	ConsumerOriginalCreditTransactionNotAccepted CardDisputeVisaUserSubmissionsChargebackConsumerOriginalCreditTransactionNotAccepted `json:"consumer_original_credit_transaction_not_accepted,required,nullable"`
	// Merchandise quality issue. Present if and only if `category` is
	// `consumer_quality_merchandise`.
	ConsumerQualityMerchandise CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandise `json:"consumer_quality_merchandise,required,nullable"`
	// Services quality issue. Present if and only if `category` is
	// `consumer_quality_services`.
	ConsumerQualityServices CardDisputeVisaUserSubmissionsChargebackConsumerQualityServices `json:"consumer_quality_services,required,nullable"`
	// Services misrepresentation. Present if and only if `category` is
	// `consumer_services_misrepresentation`.
	ConsumerServicesMisrepresentation CardDisputeVisaUserSubmissionsChargebackConsumerServicesMisrepresentation `json:"consumer_services_misrepresentation,required,nullable"`
	// Services not as described. Present if and only if `category` is
	// `consumer_services_not_as_described`.
	ConsumerServicesNotAsDescribed CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotAsDescribed `json:"consumer_services_not_as_described,required,nullable"`
	// Services not received. Present if and only if `category` is
	// `consumer_services_not_received`.
	ConsumerServicesNotReceived CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceived `json:"consumer_services_not_received,required,nullable"`
	// Fraud. Present if and only if `category` is `fraud`.
	Fraud CardDisputeVisaUserSubmissionsChargebackFraud `json:"fraud,required,nullable"`
	// Processing error. Present if and only if `category` is `processing_error`.
	ProcessingError CardDisputeVisaUserSubmissionsChargebackProcessingError `json:"processing_error,required,nullable"`
	JSON            cardDisputeVisaUserSubmissionsChargebackJSON            `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackJSON contains the JSON metadata for the
// struct [CardDisputeVisaUserSubmissionsChargeback]
type cardDisputeVisaUserSubmissionsChargebackJSON struct {
	Authorization                                apijson.Field
	Category                                     apijson.Field
	ConsumerCanceledMerchandise                  apijson.Field
	ConsumerCanceledRecurringTransaction         apijson.Field
	ConsumerCanceledServices                     apijson.Field
	ConsumerCounterfeitMerchandise               apijson.Field
	ConsumerCreditNotProcessed                   apijson.Field
	ConsumerDamagedOrDefectiveMerchandise        apijson.Field
	ConsumerMerchandiseMisrepresentation         apijson.Field
	ConsumerMerchandiseNotAsDescribed            apijson.Field
	ConsumerMerchandiseNotReceived               apijson.Field
	ConsumerNonReceiptOfCash                     apijson.Field
	ConsumerOriginalCreditTransactionNotAccepted apijson.Field
	ConsumerQualityMerchandise                   apijson.Field
	ConsumerQualityServices                      apijson.Field
	ConsumerServicesMisrepresentation            apijson.Field
	ConsumerServicesNotAsDescribed               apijson.Field
	ConsumerServicesNotReceived                  apijson.Field
	Fraud                                        apijson.Field
	ProcessingError                              apijson.Field
	raw                                          string
	ExtraFields                                  map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargeback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackJSON) RawJSON() string {
	return r.raw
}

// Authorization. Present if and only if `category` is `authorization`.
type CardDisputeVisaUserSubmissionsChargebackAuthorization struct {
	// Account status.
	AccountStatus CardDisputeVisaUserSubmissionsChargebackAuthorizationAccountStatus `json:"account_status,required"`
	JSON          cardDisputeVisaUserSubmissionsChargebackAuthorizationJSON          `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackAuthorizationJSON contains the JSON
// metadata for the struct [CardDisputeVisaUserSubmissionsChargebackAuthorization]
type cardDisputeVisaUserSubmissionsChargebackAuthorizationJSON struct {
	AccountStatus apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackAuthorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackAuthorizationJSON) RawJSON() string {
	return r.raw
}

// Account status.
type CardDisputeVisaUserSubmissionsChargebackAuthorizationAccountStatus string

const (
	CardDisputeVisaUserSubmissionsChargebackAuthorizationAccountStatusAccountClosed CardDisputeVisaUserSubmissionsChargebackAuthorizationAccountStatus = "account_closed"
	CardDisputeVisaUserSubmissionsChargebackAuthorizationAccountStatusCreditProblem CardDisputeVisaUserSubmissionsChargebackAuthorizationAccountStatus = "credit_problem"
	CardDisputeVisaUserSubmissionsChargebackAuthorizationAccountStatusFraud         CardDisputeVisaUserSubmissionsChargebackAuthorizationAccountStatus = "fraud"
)

func (r CardDisputeVisaUserSubmissionsChargebackAuthorizationAccountStatus) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsChargebackAuthorizationAccountStatusAccountClosed, CardDisputeVisaUserSubmissionsChargebackAuthorizationAccountStatusCreditProblem, CardDisputeVisaUserSubmissionsChargebackAuthorizationAccountStatusFraud:
		return true
	}
	return false
}

// Category.
type CardDisputeVisaUserSubmissionsChargebackCategory string

const (
	CardDisputeVisaUserSubmissionsChargebackCategoryAuthorization                                CardDisputeVisaUserSubmissionsChargebackCategory = "authorization"
	CardDisputeVisaUserSubmissionsChargebackCategoryConsumerCanceledMerchandise                  CardDisputeVisaUserSubmissionsChargebackCategory = "consumer_canceled_merchandise"
	CardDisputeVisaUserSubmissionsChargebackCategoryConsumerCanceledRecurringTransaction         CardDisputeVisaUserSubmissionsChargebackCategory = "consumer_canceled_recurring_transaction"
	CardDisputeVisaUserSubmissionsChargebackCategoryConsumerCanceledServices                     CardDisputeVisaUserSubmissionsChargebackCategory = "consumer_canceled_services"
	CardDisputeVisaUserSubmissionsChargebackCategoryConsumerCounterfeitMerchandise               CardDisputeVisaUserSubmissionsChargebackCategory = "consumer_counterfeit_merchandise"
	CardDisputeVisaUserSubmissionsChargebackCategoryConsumerCreditNotProcessed                   CardDisputeVisaUserSubmissionsChargebackCategory = "consumer_credit_not_processed"
	CardDisputeVisaUserSubmissionsChargebackCategoryConsumerDamagedOrDefectiveMerchandise        CardDisputeVisaUserSubmissionsChargebackCategory = "consumer_damaged_or_defective_merchandise"
	CardDisputeVisaUserSubmissionsChargebackCategoryConsumerMerchandiseMisrepresentation         CardDisputeVisaUserSubmissionsChargebackCategory = "consumer_merchandise_misrepresentation"
	CardDisputeVisaUserSubmissionsChargebackCategoryConsumerMerchandiseNotAsDescribed            CardDisputeVisaUserSubmissionsChargebackCategory = "consumer_merchandise_not_as_described"
	CardDisputeVisaUserSubmissionsChargebackCategoryConsumerMerchandiseNotReceived               CardDisputeVisaUserSubmissionsChargebackCategory = "consumer_merchandise_not_received"
	CardDisputeVisaUserSubmissionsChargebackCategoryConsumerNonReceiptOfCash                     CardDisputeVisaUserSubmissionsChargebackCategory = "consumer_non_receipt_of_cash"
	CardDisputeVisaUserSubmissionsChargebackCategoryConsumerOriginalCreditTransactionNotAccepted CardDisputeVisaUserSubmissionsChargebackCategory = "consumer_original_credit_transaction_not_accepted"
	CardDisputeVisaUserSubmissionsChargebackCategoryConsumerQualityMerchandise                   CardDisputeVisaUserSubmissionsChargebackCategory = "consumer_quality_merchandise"
	CardDisputeVisaUserSubmissionsChargebackCategoryConsumerQualityServices                      CardDisputeVisaUserSubmissionsChargebackCategory = "consumer_quality_services"
	CardDisputeVisaUserSubmissionsChargebackCategoryConsumerServicesMisrepresentation            CardDisputeVisaUserSubmissionsChargebackCategory = "consumer_services_misrepresentation"
	CardDisputeVisaUserSubmissionsChargebackCategoryConsumerServicesNotAsDescribed               CardDisputeVisaUserSubmissionsChargebackCategory = "consumer_services_not_as_described"
	CardDisputeVisaUserSubmissionsChargebackCategoryConsumerServicesNotReceived                  CardDisputeVisaUserSubmissionsChargebackCategory = "consumer_services_not_received"
	CardDisputeVisaUserSubmissionsChargebackCategoryFraud                                        CardDisputeVisaUserSubmissionsChargebackCategory = "fraud"
	CardDisputeVisaUserSubmissionsChargebackCategoryProcessingError                              CardDisputeVisaUserSubmissionsChargebackCategory = "processing_error"
)

func (r CardDisputeVisaUserSubmissionsChargebackCategory) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsChargebackCategoryAuthorization, CardDisputeVisaUserSubmissionsChargebackCategoryConsumerCanceledMerchandise, CardDisputeVisaUserSubmissionsChargebackCategoryConsumerCanceledRecurringTransaction, CardDisputeVisaUserSubmissionsChargebackCategoryConsumerCanceledServices, CardDisputeVisaUserSubmissionsChargebackCategoryConsumerCounterfeitMerchandise, CardDisputeVisaUserSubmissionsChargebackCategoryConsumerCreditNotProcessed, CardDisputeVisaUserSubmissionsChargebackCategoryConsumerDamagedOrDefectiveMerchandise, CardDisputeVisaUserSubmissionsChargebackCategoryConsumerMerchandiseMisrepresentation, CardDisputeVisaUserSubmissionsChargebackCategoryConsumerMerchandiseNotAsDescribed, CardDisputeVisaUserSubmissionsChargebackCategoryConsumerMerchandiseNotReceived, CardDisputeVisaUserSubmissionsChargebackCategoryConsumerNonReceiptOfCash, CardDisputeVisaUserSubmissionsChargebackCategoryConsumerOriginalCreditTransactionNotAccepted, CardDisputeVisaUserSubmissionsChargebackCategoryConsumerQualityMerchandise, CardDisputeVisaUserSubmissionsChargebackCategoryConsumerQualityServices, CardDisputeVisaUserSubmissionsChargebackCategoryConsumerServicesMisrepresentation, CardDisputeVisaUserSubmissionsChargebackCategoryConsumerServicesNotAsDescribed, CardDisputeVisaUserSubmissionsChargebackCategoryConsumerServicesNotReceived, CardDisputeVisaUserSubmissionsChargebackCategoryFraud, CardDisputeVisaUserSubmissionsChargebackCategoryProcessingError:
		return true
	}
	return false
}

// Canceled merchandise. Present if and only if `category` is
// `consumer_canceled_merchandise`.
type CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandise struct {
	// Cardholder cancellation.
	CardholderCancellation CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseCardholderCancellation `json:"cardholder_cancellation,required,nullable"`
	// Merchant resolution attempted.
	MerchantResolutionAttempted CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseMerchantResolutionAttempted `json:"merchant_resolution_attempted,required"`
	// Not returned. Present if and only if `return_outcome` is `not_returned`.
	NotReturned CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseNotReturned `json:"not_returned,required,nullable"`
	// Purchase explanation.
	PurchaseExplanation string `json:"purchase_explanation,required"`
	// Received or expected at.
	ReceivedOrExpectedAt time.Time `json:"received_or_expected_at,required" format:"date"`
	// Return attempted. Present if and only if `return_outcome` is `return_attempted`.
	ReturnAttempted CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnAttempted `json:"return_attempted,required,nullable"`
	// Return outcome.
	ReturnOutcome CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnOutcome `json:"return_outcome,required"`
	// Returned. Present if and only if `return_outcome` is `returned`.
	Returned CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturned `json:"returned,required,nullable"`
	JSON     cardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseJSON     `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseJSON contains
// the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandise]
type cardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseJSON struct {
	CardholderCancellation      apijson.Field
	MerchantResolutionAttempted apijson.Field
	NotReturned                 apijson.Field
	PurchaseExplanation         apijson.Field
	ReceivedOrExpectedAt        apijson.Field
	ReturnAttempted             apijson.Field
	ReturnOutcome               apijson.Field
	Returned                    apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandise) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseJSON) RawJSON() string {
	return r.raw
}

// Cardholder cancellation.
type CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseCardholderCancellation struct {
	// Canceled at.
	CanceledAt time.Time `json:"canceled_at,required" format:"date"`
	// Canceled prior to ship date.
	CanceledPriorToShipDate CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseCardholderCancellationCanceledPriorToShipDate `json:"canceled_prior_to_ship_date,required"`
	// Cancellation policy provided.
	CancellationPolicyProvided CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseCardholderCancellationCancellationPolicyProvided `json:"cancellation_policy_provided,required"`
	// Reason.
	Reason string                                                                                        `json:"reason,required"`
	JSON   cardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseCardholderCancellationJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseCardholderCancellationJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseCardholderCancellation]
type cardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseCardholderCancellationJSON struct {
	CanceledAt                 apijson.Field
	CanceledPriorToShipDate    apijson.Field
	CancellationPolicyProvided apijson.Field
	Reason                     apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseCardholderCancellation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseCardholderCancellationJSON) RawJSON() string {
	return r.raw
}

// Canceled prior to ship date.
type CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseCardholderCancellationCanceledPriorToShipDate string

const (
	CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseCardholderCancellationCanceledPriorToShipDateCanceledPriorToShipDate    CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseCardholderCancellationCanceledPriorToShipDate = "canceled_prior_to_ship_date"
	CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseCardholderCancellationCanceledPriorToShipDateNotCanceledPriorToShipDate CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseCardholderCancellationCanceledPriorToShipDate = "not_canceled_prior_to_ship_date"
)

func (r CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseCardholderCancellationCanceledPriorToShipDate) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseCardholderCancellationCanceledPriorToShipDateCanceledPriorToShipDate, CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseCardholderCancellationCanceledPriorToShipDateNotCanceledPriorToShipDate:
		return true
	}
	return false
}

// Cancellation policy provided.
type CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseCardholderCancellationCancellationPolicyProvided string

const (
	CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseCardholderCancellationCancellationPolicyProvidedNotProvided CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseCardholderCancellationCancellationPolicyProvided = "not_provided"
	CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseCardholderCancellationCancellationPolicyProvidedProvided    CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseCardholderCancellationCancellationPolicyProvided = "provided"
)

func (r CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseCardholderCancellationCancellationPolicyProvided) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseCardholderCancellationCancellationPolicyProvidedNotProvided, CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseCardholderCancellationCancellationPolicyProvidedProvided:
		return true
	}
	return false
}

// Merchant resolution attempted.
type CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseMerchantResolutionAttempted string

const (
	CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseMerchantResolutionAttemptedAttempted            CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseMerchantResolutionAttempted = "attempted"
	CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseMerchantResolutionAttemptedProhibitedByLocalLaw CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseMerchantResolutionAttempted = "prohibited_by_local_law"
)

func (r CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseMerchantResolutionAttempted) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseMerchantResolutionAttemptedAttempted, CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseMerchantResolutionAttemptedProhibitedByLocalLaw:
		return true
	}
	return false
}

// Not returned. Present if and only if `return_outcome` is `not_returned`.
type CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseNotReturned struct {
	JSON cardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseNotReturnedJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseNotReturnedJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseNotReturned]
type cardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseNotReturnedJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseNotReturned) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseNotReturnedJSON) RawJSON() string {
	return r.raw
}

// Return attempted. Present if and only if `return_outcome` is `return_attempted`.
type CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnAttempted struct {
	// Attempt explanation.
	AttemptExplanation string `json:"attempt_explanation,required"`
	// Attempt reason.
	AttemptReason CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnAttemptedAttemptReason `json:"attempt_reason,required"`
	// Attempted at.
	AttemptedAt time.Time `json:"attempted_at,required" format:"date"`
	// Merchandise disposition.
	MerchandiseDisposition string                                                                                 `json:"merchandise_disposition,required"`
	JSON                   cardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnAttemptedJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnAttemptedJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnAttempted]
type cardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnAttemptedJSON struct {
	AttemptExplanation     apijson.Field
	AttemptReason          apijson.Field
	AttemptedAt            apijson.Field
	MerchandiseDisposition apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnAttempted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnAttemptedJSON) RawJSON() string {
	return r.raw
}

// Attempt reason.
type CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnAttemptedAttemptReason string

const (
	CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnAttemptedAttemptReasonMerchantNotResponding         CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnAttemptedAttemptReason = "merchant_not_responding"
	CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnAttemptedAttemptReasonNoReturnAuthorizationProvided CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnAttemptedAttemptReason = "no_return_authorization_provided"
	CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnAttemptedAttemptReasonNoReturnInstructions          CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnAttemptedAttemptReason = "no_return_instructions"
	CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnAttemptedAttemptReasonRequestedNotToReturn          CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnAttemptedAttemptReason = "requested_not_to_return"
	CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnAttemptedAttemptReasonReturnNotAccepted             CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnAttemptedAttemptReason = "return_not_accepted"
)

func (r CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnAttemptedAttemptReason) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnAttemptedAttemptReasonMerchantNotResponding, CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnAttemptedAttemptReasonNoReturnAuthorizationProvided, CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnAttemptedAttemptReasonNoReturnInstructions, CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnAttemptedAttemptReasonRequestedNotToReturn, CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnAttemptedAttemptReasonReturnNotAccepted:
		return true
	}
	return false
}

// Return outcome.
type CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnOutcome string

const (
	CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnOutcomeNotReturned     CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnOutcome = "not_returned"
	CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnOutcomeReturned        CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnOutcome = "returned"
	CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnOutcomeReturnAttempted CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnOutcome = "return_attempted"
)

func (r CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnOutcome) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnOutcomeNotReturned, CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnOutcomeReturned, CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnOutcomeReturnAttempted:
		return true
	}
	return false
}

// Returned. Present if and only if `return_outcome` is `returned`.
type CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturned struct {
	// Merchant received return at.
	MerchantReceivedReturnAt time.Time `json:"merchant_received_return_at,required,nullable" format:"date"`
	// Other explanation. Required if and only if the return method is `other`.
	OtherExplanation string `json:"other_explanation,required,nullable"`
	// Return method.
	ReturnMethod CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnedReturnMethod `json:"return_method,required"`
	// Returned at.
	ReturnedAt time.Time `json:"returned_at,required" format:"date"`
	// Tracking number.
	TrackingNumber string                                                                          `json:"tracking_number,required,nullable"`
	JSON           cardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnedJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnedJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturned]
type cardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnedJSON struct {
	MerchantReceivedReturnAt apijson.Field
	OtherExplanation         apijson.Field
	ReturnMethod             apijson.Field
	ReturnedAt               apijson.Field
	TrackingNumber           apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturned) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnedJSON) RawJSON() string {
	return r.raw
}

// Return method.
type CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnedReturnMethod string

const (
	CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnedReturnMethodDhl           CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnedReturnMethod = "dhl"
	CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnedReturnMethodFaceToFace    CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnedReturnMethod = "face_to_face"
	CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnedReturnMethodFedex         CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnedReturnMethod = "fedex"
	CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnedReturnMethodOther         CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnedReturnMethod = "other"
	CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnedReturnMethodPostalService CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnedReturnMethod = "postal_service"
	CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnedReturnMethodUps           CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnedReturnMethod = "ups"
)

func (r CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnedReturnMethod) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnedReturnMethodDhl, CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnedReturnMethodFaceToFace, CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnedReturnMethodFedex, CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnedReturnMethodOther, CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnedReturnMethodPostalService, CardDisputeVisaUserSubmissionsChargebackConsumerCanceledMerchandiseReturnedReturnMethodUps:
		return true
	}
	return false
}

// Canceled recurring transaction. Present if and only if `category` is
// `consumer_canceled_recurring_transaction`.
type CardDisputeVisaUserSubmissionsChargebackConsumerCanceledRecurringTransaction struct {
	// Cancellation target.
	CancellationTarget CardDisputeVisaUserSubmissionsChargebackConsumerCanceledRecurringTransactionCancellationTarget `json:"cancellation_target,required"`
	// Merchant contact methods.
	MerchantContactMethods CardDisputeVisaUserSubmissionsChargebackConsumerCanceledRecurringTransactionMerchantContactMethods `json:"merchant_contact_methods,required"`
	// Other form of payment explanation.
	OtherFormOfPaymentExplanation string `json:"other_form_of_payment_explanation,required,nullable"`
	// Transaction or account canceled at.
	TransactionOrAccountCanceledAt time.Time                                                                        `json:"transaction_or_account_canceled_at,required" format:"date"`
	JSON                           cardDisputeVisaUserSubmissionsChargebackConsumerCanceledRecurringTransactionJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerCanceledRecurringTransactionJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerCanceledRecurringTransaction]
type cardDisputeVisaUserSubmissionsChargebackConsumerCanceledRecurringTransactionJSON struct {
	CancellationTarget             apijson.Field
	MerchantContactMethods         apijson.Field
	OtherFormOfPaymentExplanation  apijson.Field
	TransactionOrAccountCanceledAt apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerCanceledRecurringTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerCanceledRecurringTransactionJSON) RawJSON() string {
	return r.raw
}

// Cancellation target.
type CardDisputeVisaUserSubmissionsChargebackConsumerCanceledRecurringTransactionCancellationTarget string

const (
	CardDisputeVisaUserSubmissionsChargebackConsumerCanceledRecurringTransactionCancellationTargetAccount     CardDisputeVisaUserSubmissionsChargebackConsumerCanceledRecurringTransactionCancellationTarget = "account"
	CardDisputeVisaUserSubmissionsChargebackConsumerCanceledRecurringTransactionCancellationTargetTransaction CardDisputeVisaUserSubmissionsChargebackConsumerCanceledRecurringTransactionCancellationTarget = "transaction"
)

func (r CardDisputeVisaUserSubmissionsChargebackConsumerCanceledRecurringTransactionCancellationTarget) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsChargebackConsumerCanceledRecurringTransactionCancellationTargetAccount, CardDisputeVisaUserSubmissionsChargebackConsumerCanceledRecurringTransactionCancellationTargetTransaction:
		return true
	}
	return false
}

// Merchant contact methods.
type CardDisputeVisaUserSubmissionsChargebackConsumerCanceledRecurringTransactionMerchantContactMethods struct {
	// Application name.
	ApplicationName string `json:"application_name,required,nullable"`
	// Call center phone number.
	CallCenterPhoneNumber string `json:"call_center_phone_number,required,nullable"`
	// Email address.
	EmailAddress string `json:"email_address,required,nullable"`
	// In person address.
	InPersonAddress string `json:"in_person_address,required,nullable"`
	// Mailing address.
	MailingAddress string `json:"mailing_address,required,nullable"`
	// Text phone number.
	TextPhoneNumber string                                                                                                 `json:"text_phone_number,required,nullable"`
	JSON            cardDisputeVisaUserSubmissionsChargebackConsumerCanceledRecurringTransactionMerchantContactMethodsJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerCanceledRecurringTransactionMerchantContactMethodsJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerCanceledRecurringTransactionMerchantContactMethods]
type cardDisputeVisaUserSubmissionsChargebackConsumerCanceledRecurringTransactionMerchantContactMethodsJSON struct {
	ApplicationName       apijson.Field
	CallCenterPhoneNumber apijson.Field
	EmailAddress          apijson.Field
	InPersonAddress       apijson.Field
	MailingAddress        apijson.Field
	TextPhoneNumber       apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerCanceledRecurringTransactionMerchantContactMethods) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerCanceledRecurringTransactionMerchantContactMethodsJSON) RawJSON() string {
	return r.raw
}

// Canceled services. Present if and only if `category` is
// `consumer_canceled_services`.
type CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServices struct {
	// Cardholder cancellation.
	CardholderCancellation CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesCardholderCancellation `json:"cardholder_cancellation,required"`
	// Contracted at.
	ContractedAt time.Time `json:"contracted_at,required" format:"date"`
	// Guaranteed reservation explanation. Present if and only if `service_type` is
	// `guaranteed_reservation`.
	GuaranteedReservation CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesGuaranteedReservation `json:"guaranteed_reservation,required,nullable"`
	// Merchant resolution attempted.
	MerchantResolutionAttempted CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesMerchantResolutionAttempted `json:"merchant_resolution_attempted,required"`
	// Other service type explanation. Present if and only if `service_type` is
	// `other`.
	Other CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesOther `json:"other,required,nullable"`
	// Purchase explanation.
	PurchaseExplanation string `json:"purchase_explanation,required"`
	// Service type.
	ServiceType CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesServiceType `json:"service_type,required"`
	// Timeshare explanation. Present if and only if `service_type` is `timeshare`.
	Timeshare CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesTimeshare `json:"timeshare,required,nullable"`
	JSON      cardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesJSON      `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesJSON contains
// the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServices]
type cardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesJSON struct {
	CardholderCancellation      apijson.Field
	ContractedAt                apijson.Field
	GuaranteedReservation       apijson.Field
	MerchantResolutionAttempted apijson.Field
	Other                       apijson.Field
	PurchaseExplanation         apijson.Field
	ServiceType                 apijson.Field
	Timeshare                   apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServices) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesJSON) RawJSON() string {
	return r.raw
}

// Cardholder cancellation.
type CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesCardholderCancellation struct {
	// Canceled at.
	CanceledAt time.Time `json:"canceled_at,required" format:"date"`
	// Cancellation policy provided.
	CancellationPolicyProvided CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesCardholderCancellationCancellationPolicyProvided `json:"cancellation_policy_provided,required"`
	// Reason.
	Reason string                                                                                     `json:"reason,required"`
	JSON   cardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesCardholderCancellationJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesCardholderCancellationJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesCardholderCancellation]
type cardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesCardholderCancellationJSON struct {
	CanceledAt                 apijson.Field
	CancellationPolicyProvided apijson.Field
	Reason                     apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesCardholderCancellation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesCardholderCancellationJSON) RawJSON() string {
	return r.raw
}

// Cancellation policy provided.
type CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesCardholderCancellationCancellationPolicyProvided string

const (
	CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesCardholderCancellationCancellationPolicyProvidedNotProvided CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesCardholderCancellationCancellationPolicyProvided = "not_provided"
	CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesCardholderCancellationCancellationPolicyProvidedProvided    CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesCardholderCancellationCancellationPolicyProvided = "provided"
)

func (r CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesCardholderCancellationCancellationPolicyProvided) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesCardholderCancellationCancellationPolicyProvidedNotProvided, CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesCardholderCancellationCancellationPolicyProvidedProvided:
		return true
	}
	return false
}

// Guaranteed reservation explanation. Present if and only if `service_type` is
// `guaranteed_reservation`.
type CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesGuaranteedReservation struct {
	// Explanation.
	Explanation CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesGuaranteedReservationExplanation `json:"explanation,required"`
	JSON        cardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesGuaranteedReservationJSON        `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesGuaranteedReservationJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesGuaranteedReservation]
type cardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesGuaranteedReservationJSON struct {
	Explanation apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesGuaranteedReservation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesGuaranteedReservationJSON) RawJSON() string {
	return r.raw
}

// Explanation.
type CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesGuaranteedReservationExplanation string

const (
	CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesGuaranteedReservationExplanationCardholderCanceledPriorToService                         CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesGuaranteedReservationExplanation = "cardholder_canceled_prior_to_service"
	CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesGuaranteedReservationExplanationCardholderCancellationAttemptWithin24HoursOfConfirmation CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesGuaranteedReservationExplanation = "cardholder_cancellation_attempt_within_24_hours_of_confirmation"
	CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesGuaranteedReservationExplanationMerchantBilledNoShow                                     CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesGuaranteedReservationExplanation = "merchant_billed_no_show"
)

func (r CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesGuaranteedReservationExplanation) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesGuaranteedReservationExplanationCardholderCanceledPriorToService, CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesGuaranteedReservationExplanationCardholderCancellationAttemptWithin24HoursOfConfirmation, CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesGuaranteedReservationExplanationMerchantBilledNoShow:
		return true
	}
	return false
}

// Merchant resolution attempted.
type CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesMerchantResolutionAttempted string

const (
	CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesMerchantResolutionAttemptedAttempted            CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesMerchantResolutionAttempted = "attempted"
	CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesMerchantResolutionAttemptedProhibitedByLocalLaw CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesMerchantResolutionAttempted = "prohibited_by_local_law"
)

func (r CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesMerchantResolutionAttempted) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesMerchantResolutionAttemptedAttempted, CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesMerchantResolutionAttemptedProhibitedByLocalLaw:
		return true
	}
	return false
}

// Other service type explanation. Present if and only if `service_type` is
// `other`.
type CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesOther struct {
	JSON cardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesOtherJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesOtherJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesOther]
type cardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesOtherJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesOther) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesOtherJSON) RawJSON() string {
	return r.raw
}

// Service type.
type CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesServiceType string

const (
	CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesServiceTypeGuaranteedReservation CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesServiceType = "guaranteed_reservation"
	CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesServiceTypeOther                 CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesServiceType = "other"
	CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesServiceTypeTimeshare             CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesServiceType = "timeshare"
)

func (r CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesServiceType) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesServiceTypeGuaranteedReservation, CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesServiceTypeOther, CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesServiceTypeTimeshare:
		return true
	}
	return false
}

// Timeshare explanation. Present if and only if `service_type` is `timeshare`.
type CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesTimeshare struct {
	JSON cardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesTimeshareJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesTimeshareJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesTimeshare]
type cardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesTimeshareJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesTimeshare) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerCanceledServicesTimeshareJSON) RawJSON() string {
	return r.raw
}

// Counterfeit merchandise. Present if and only if `category` is
// `consumer_counterfeit_merchandise`.
type CardDisputeVisaUserSubmissionsChargebackConsumerCounterfeitMerchandise struct {
	// Counterfeit explanation.
	CounterfeitExplanation string `json:"counterfeit_explanation,required"`
	// Disposition explanation.
	DispositionExplanation string `json:"disposition_explanation,required"`
	// Order explanation.
	OrderExplanation string `json:"order_explanation,required"`
	// Received at.
	ReceivedAt time.Time                                                                  `json:"received_at,required" format:"date"`
	JSON       cardDisputeVisaUserSubmissionsChargebackConsumerCounterfeitMerchandiseJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerCounterfeitMerchandiseJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerCounterfeitMerchandise]
type cardDisputeVisaUserSubmissionsChargebackConsumerCounterfeitMerchandiseJSON struct {
	CounterfeitExplanation apijson.Field
	DispositionExplanation apijson.Field
	OrderExplanation       apijson.Field
	ReceivedAt             apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerCounterfeitMerchandise) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerCounterfeitMerchandiseJSON) RawJSON() string {
	return r.raw
}

// Credit not processed. Present if and only if `category` is
// `consumer_credit_not_processed`.
type CardDisputeVisaUserSubmissionsChargebackConsumerCreditNotProcessed struct {
	// Canceled or returned at.
	CanceledOrReturnedAt time.Time `json:"canceled_or_returned_at,required,nullable" format:"date"`
	// Credit expected at.
	CreditExpectedAt time.Time                                                              `json:"credit_expected_at,required,nullable" format:"date"`
	JSON             cardDisputeVisaUserSubmissionsChargebackConsumerCreditNotProcessedJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerCreditNotProcessedJSON contains
// the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerCreditNotProcessed]
type cardDisputeVisaUserSubmissionsChargebackConsumerCreditNotProcessedJSON struct {
	CanceledOrReturnedAt apijson.Field
	CreditExpectedAt     apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerCreditNotProcessed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerCreditNotProcessedJSON) RawJSON() string {
	return r.raw
}

// Damaged or defective merchandise. Present if and only if `category` is
// `consumer_damaged_or_defective_merchandise`.
type CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandise struct {
	// Merchant resolution attempted.
	MerchantResolutionAttempted CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseMerchantResolutionAttempted `json:"merchant_resolution_attempted,required"`
	// Not returned. Present if and only if `return_outcome` is `not_returned`.
	NotReturned CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseNotReturned `json:"not_returned,required,nullable"`
	// Order and issue explanation.
	OrderAndIssueExplanation string `json:"order_and_issue_explanation,required"`
	// Received at.
	ReceivedAt time.Time `json:"received_at,required" format:"date"`
	// Return attempted. Present if and only if `return_outcome` is `return_attempted`.
	ReturnAttempted CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttempted `json:"return_attempted,required,nullable"`
	// Return outcome.
	ReturnOutcome CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnOutcome `json:"return_outcome,required"`
	// Returned. Present if and only if `return_outcome` is `returned`.
	Returned CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturned `json:"returned,required,nullable"`
	JSON     cardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseJSON     `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandise]
type cardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseJSON struct {
	MerchantResolutionAttempted apijson.Field
	NotReturned                 apijson.Field
	OrderAndIssueExplanation    apijson.Field
	ReceivedAt                  apijson.Field
	ReturnAttempted             apijson.Field
	ReturnOutcome               apijson.Field
	Returned                    apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandise) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseJSON) RawJSON() string {
	return r.raw
}

// Merchant resolution attempted.
type CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseMerchantResolutionAttempted string

const (
	CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseMerchantResolutionAttemptedAttempted            CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseMerchantResolutionAttempted = "attempted"
	CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseMerchantResolutionAttemptedProhibitedByLocalLaw CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseMerchantResolutionAttempted = "prohibited_by_local_law"
)

func (r CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseMerchantResolutionAttempted) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseMerchantResolutionAttemptedAttempted, CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseMerchantResolutionAttemptedProhibitedByLocalLaw:
		return true
	}
	return false
}

// Not returned. Present if and only if `return_outcome` is `not_returned`.
type CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseNotReturned struct {
	JSON cardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseNotReturnedJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseNotReturnedJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseNotReturned]
type cardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseNotReturnedJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseNotReturned) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseNotReturnedJSON) RawJSON() string {
	return r.raw
}

// Return attempted. Present if and only if `return_outcome` is `return_attempted`.
type CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttempted struct {
	// Attempt explanation.
	AttemptExplanation string `json:"attempt_explanation,required"`
	// Attempt reason.
	AttemptReason CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReason `json:"attempt_reason,required"`
	// Attempted at.
	AttemptedAt time.Time `json:"attempted_at,required" format:"date"`
	// Merchandise disposition.
	MerchandiseDisposition string                                                                                           `json:"merchandise_disposition,required"`
	JSON                   cardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttemptedJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttemptedJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttempted]
type cardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttemptedJSON struct {
	AttemptExplanation     apijson.Field
	AttemptReason          apijson.Field
	AttemptedAt            apijson.Field
	MerchandiseDisposition apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttempted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttemptedJSON) RawJSON() string {
	return r.raw
}

// Attempt reason.
type CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReason string

const (
	CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReasonMerchantNotResponding         CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReason = "merchant_not_responding"
	CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReasonNoReturnAuthorizationProvided CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReason = "no_return_authorization_provided"
	CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReasonNoReturnInstructions          CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReason = "no_return_instructions"
	CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReasonRequestedNotToReturn          CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReason = "requested_not_to_return"
	CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReasonReturnNotAccepted             CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReason = "return_not_accepted"
)

func (r CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReason) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReasonMerchantNotResponding, CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReasonNoReturnAuthorizationProvided, CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReasonNoReturnInstructions, CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReasonRequestedNotToReturn, CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReasonReturnNotAccepted:
		return true
	}
	return false
}

// Return outcome.
type CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnOutcome string

const (
	CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnOutcomeNotReturned     CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnOutcome = "not_returned"
	CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnOutcomeReturned        CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnOutcome = "returned"
	CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnOutcomeReturnAttempted CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnOutcome = "return_attempted"
)

func (r CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnOutcome) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnOutcomeNotReturned, CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnOutcomeReturned, CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnOutcomeReturnAttempted:
		return true
	}
	return false
}

// Returned. Present if and only if `return_outcome` is `returned`.
type CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturned struct {
	// Merchant received return at.
	MerchantReceivedReturnAt time.Time `json:"merchant_received_return_at,required,nullable" format:"date"`
	// Other explanation. Required if and only if the return method is `other`.
	OtherExplanation string `json:"other_explanation,required,nullable"`
	// Return method.
	ReturnMethod CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethod `json:"return_method,required"`
	// Returned at.
	ReturnedAt time.Time `json:"returned_at,required" format:"date"`
	// Tracking number.
	TrackingNumber string                                                                                    `json:"tracking_number,required,nullable"`
	JSON           cardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnedJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnedJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturned]
type cardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnedJSON struct {
	MerchantReceivedReturnAt apijson.Field
	OtherExplanation         apijson.Field
	ReturnMethod             apijson.Field
	ReturnedAt               apijson.Field
	TrackingNumber           apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturned) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnedJSON) RawJSON() string {
	return r.raw
}

// Return method.
type CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethod string

const (
	CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethodDhl           CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethod = "dhl"
	CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethodFaceToFace    CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethod = "face_to_face"
	CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethodFedex         CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethod = "fedex"
	CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethodOther         CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethod = "other"
	CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethodPostalService CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethod = "postal_service"
	CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethodUps           CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethod = "ups"
)

func (r CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethod) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethodDhl, CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethodFaceToFace, CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethodFedex, CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethodOther, CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethodPostalService, CardDisputeVisaUserSubmissionsChargebackConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethodUps:
		return true
	}
	return false
}

// Merchandise misrepresentation. Present if and only if `category` is
// `consumer_merchandise_misrepresentation`.
type CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentation struct {
	// Merchant resolution attempted.
	MerchantResolutionAttempted CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationMerchantResolutionAttempted `json:"merchant_resolution_attempted,required"`
	// Misrepresentation explanation.
	MisrepresentationExplanation string `json:"misrepresentation_explanation,required"`
	// Not returned. Present if and only if `return_outcome` is `not_returned`.
	NotReturned CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationNotReturned `json:"not_returned,required,nullable"`
	// Purchase explanation.
	PurchaseExplanation string `json:"purchase_explanation,required"`
	// Received at.
	ReceivedAt time.Time `json:"received_at,required" format:"date"`
	// Return attempted. Present if and only if `return_outcome` is `return_attempted`.
	ReturnAttempted CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnAttempted `json:"return_attempted,required,nullable"`
	// Return outcome.
	ReturnOutcome CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnOutcome `json:"return_outcome,required"`
	// Returned. Present if and only if `return_outcome` is `returned`.
	Returned CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturned `json:"returned,required,nullable"`
	JSON     cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationJSON     `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentation]
type cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationJSON struct {
	MerchantResolutionAttempted  apijson.Field
	MisrepresentationExplanation apijson.Field
	NotReturned                  apijson.Field
	PurchaseExplanation          apijson.Field
	ReceivedAt                   apijson.Field
	ReturnAttempted              apijson.Field
	ReturnOutcome                apijson.Field
	Returned                     apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationJSON) RawJSON() string {
	return r.raw
}

// Merchant resolution attempted.
type CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationMerchantResolutionAttempted string

const (
	CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationMerchantResolutionAttemptedAttempted            CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationMerchantResolutionAttempted = "attempted"
	CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationMerchantResolutionAttemptedProhibitedByLocalLaw CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationMerchantResolutionAttempted = "prohibited_by_local_law"
)

func (r CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationMerchantResolutionAttempted) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationMerchantResolutionAttemptedAttempted, CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationMerchantResolutionAttemptedProhibitedByLocalLaw:
		return true
	}
	return false
}

// Not returned. Present if and only if `return_outcome` is `not_returned`.
type CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationNotReturned struct {
	JSON cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationNotReturnedJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationNotReturnedJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationNotReturned]
type cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationNotReturnedJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationNotReturned) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationNotReturnedJSON) RawJSON() string {
	return r.raw
}

// Return attempted. Present if and only if `return_outcome` is `return_attempted`.
type CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnAttempted struct {
	// Attempt explanation.
	AttemptExplanation string `json:"attempt_explanation,required"`
	// Attempt reason.
	AttemptReason CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReason `json:"attempt_reason,required"`
	// Attempted at.
	AttemptedAt time.Time `json:"attempted_at,required" format:"date"`
	// Merchandise disposition.
	MerchandiseDisposition string                                                                                          `json:"merchandise_disposition,required"`
	JSON                   cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnAttemptedJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnAttemptedJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnAttempted]
type cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnAttemptedJSON struct {
	AttemptExplanation     apijson.Field
	AttemptReason          apijson.Field
	AttemptedAt            apijson.Field
	MerchandiseDisposition apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnAttempted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnAttemptedJSON) RawJSON() string {
	return r.raw
}

// Attempt reason.
type CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReason string

const (
	CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReasonMerchantNotResponding         CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReason = "merchant_not_responding"
	CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReasonNoReturnAuthorizationProvided CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReason = "no_return_authorization_provided"
	CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReasonNoReturnInstructions          CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReason = "no_return_instructions"
	CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReasonRequestedNotToReturn          CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReason = "requested_not_to_return"
	CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReasonReturnNotAccepted             CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReason = "return_not_accepted"
)

func (r CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReason) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReasonMerchantNotResponding, CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReasonNoReturnAuthorizationProvided, CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReasonNoReturnInstructions, CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReasonRequestedNotToReturn, CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReasonReturnNotAccepted:
		return true
	}
	return false
}

// Return outcome.
type CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnOutcome string

const (
	CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnOutcomeNotReturned     CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnOutcome = "not_returned"
	CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnOutcomeReturned        CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnOutcome = "returned"
	CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnOutcomeReturnAttempted CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnOutcome = "return_attempted"
)

func (r CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnOutcome) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnOutcomeNotReturned, CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnOutcomeReturned, CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnOutcomeReturnAttempted:
		return true
	}
	return false
}

// Returned. Present if and only if `return_outcome` is `returned`.
type CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturned struct {
	// Merchant received return at.
	MerchantReceivedReturnAt time.Time `json:"merchant_received_return_at,required,nullable" format:"date"`
	// Other explanation. Required if and only if the return method is `other`.
	OtherExplanation string `json:"other_explanation,required,nullable"`
	// Return method.
	ReturnMethod CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnedReturnMethod `json:"return_method,required"`
	// Returned at.
	ReturnedAt time.Time `json:"returned_at,required" format:"date"`
	// Tracking number.
	TrackingNumber string                                                                                   `json:"tracking_number,required,nullable"`
	JSON           cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnedJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnedJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturned]
type cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnedJSON struct {
	MerchantReceivedReturnAt apijson.Field
	OtherExplanation         apijson.Field
	ReturnMethod             apijson.Field
	ReturnedAt               apijson.Field
	TrackingNumber           apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturned) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnedJSON) RawJSON() string {
	return r.raw
}

// Return method.
type CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnedReturnMethod string

const (
	CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnedReturnMethodDhl           CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnedReturnMethod = "dhl"
	CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnedReturnMethodFaceToFace    CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnedReturnMethod = "face_to_face"
	CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnedReturnMethodFedex         CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnedReturnMethod = "fedex"
	CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnedReturnMethodOther         CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnedReturnMethod = "other"
	CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnedReturnMethodPostalService CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnedReturnMethod = "postal_service"
	CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnedReturnMethodUps           CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnedReturnMethod = "ups"
)

func (r CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnedReturnMethod) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnedReturnMethodDhl, CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnedReturnMethodFaceToFace, CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnedReturnMethodFedex, CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnedReturnMethodOther, CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnedReturnMethodPostalService, CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseMisrepresentationReturnedReturnMethodUps:
		return true
	}
	return false
}

// Merchandise not as described. Present if and only if `category` is
// `consumer_merchandise_not_as_described`.
type CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribed struct {
	// Merchant resolution attempted.
	MerchantResolutionAttempted CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedMerchantResolutionAttempted `json:"merchant_resolution_attempted,required"`
	// Received at.
	ReceivedAt time.Time `json:"received_at,required" format:"date"`
	// Return attempted. Present if and only if `return_outcome` is `return_attempted`.
	ReturnAttempted CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnAttempted `json:"return_attempted,required,nullable"`
	// Return outcome.
	ReturnOutcome CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnOutcome `json:"return_outcome,required"`
	// Returned. Present if and only if `return_outcome` is `returned`.
	Returned CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturned `json:"returned,required,nullable"`
	JSON     cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedJSON     `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribed]
type cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedJSON struct {
	MerchantResolutionAttempted apijson.Field
	ReceivedAt                  apijson.Field
	ReturnAttempted             apijson.Field
	ReturnOutcome               apijson.Field
	Returned                    apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedJSON) RawJSON() string {
	return r.raw
}

// Merchant resolution attempted.
type CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedMerchantResolutionAttempted string

const (
	CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedMerchantResolutionAttemptedAttempted            CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedMerchantResolutionAttempted = "attempted"
	CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedMerchantResolutionAttemptedProhibitedByLocalLaw CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedMerchantResolutionAttempted = "prohibited_by_local_law"
)

func (r CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedMerchantResolutionAttempted) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedMerchantResolutionAttemptedAttempted, CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedMerchantResolutionAttemptedProhibitedByLocalLaw:
		return true
	}
	return false
}

// Return attempted. Present if and only if `return_outcome` is `return_attempted`.
type CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnAttempted struct {
	// Attempt explanation.
	AttemptExplanation string `json:"attempt_explanation,required"`
	// Attempt reason.
	AttemptReason CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReason `json:"attempt_reason,required"`
	// Attempted at.
	AttemptedAt time.Time `json:"attempted_at,required" format:"date"`
	// Merchandise disposition.
	MerchandiseDisposition string                                                                                       `json:"merchandise_disposition,required"`
	JSON                   cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnAttemptedJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnAttemptedJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnAttempted]
type cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnAttemptedJSON struct {
	AttemptExplanation     apijson.Field
	AttemptReason          apijson.Field
	AttemptedAt            apijson.Field
	MerchandiseDisposition apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnAttempted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnAttemptedJSON) RawJSON() string {
	return r.raw
}

// Attempt reason.
type CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReason string

const (
	CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReasonMerchantNotResponding         CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReason = "merchant_not_responding"
	CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReasonNoReturnAuthorizationProvided CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReason = "no_return_authorization_provided"
	CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReasonNoReturnInstructions          CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReason = "no_return_instructions"
	CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReasonRequestedNotToReturn          CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReason = "requested_not_to_return"
	CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReasonReturnNotAccepted             CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReason = "return_not_accepted"
)

func (r CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReason) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReasonMerchantNotResponding, CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReasonNoReturnAuthorizationProvided, CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReasonNoReturnInstructions, CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReasonRequestedNotToReturn, CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReasonReturnNotAccepted:
		return true
	}
	return false
}

// Return outcome.
type CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnOutcome string

const (
	CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnOutcomeReturned        CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnOutcome = "returned"
	CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnOutcomeReturnAttempted CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnOutcome = "return_attempted"
)

func (r CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnOutcome) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnOutcomeReturned, CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnOutcomeReturnAttempted:
		return true
	}
	return false
}

// Returned. Present if and only if `return_outcome` is `returned`.
type CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturned struct {
	// Merchant received return at.
	MerchantReceivedReturnAt time.Time `json:"merchant_received_return_at,required,nullable" format:"date"`
	// Other explanation. Required if and only if the return method is `other`.
	OtherExplanation string `json:"other_explanation,required,nullable"`
	// Return method.
	ReturnMethod CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnedReturnMethod `json:"return_method,required"`
	// Returned at.
	ReturnedAt time.Time `json:"returned_at,required" format:"date"`
	// Tracking number.
	TrackingNumber string                                                                                `json:"tracking_number,required,nullable"`
	JSON           cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnedJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnedJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturned]
type cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnedJSON struct {
	MerchantReceivedReturnAt apijson.Field
	OtherExplanation         apijson.Field
	ReturnMethod             apijson.Field
	ReturnedAt               apijson.Field
	TrackingNumber           apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturned) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnedJSON) RawJSON() string {
	return r.raw
}

// Return method.
type CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnedReturnMethod string

const (
	CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnedReturnMethodDhl           CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnedReturnMethod = "dhl"
	CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnedReturnMethodFaceToFace    CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnedReturnMethod = "face_to_face"
	CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnedReturnMethodFedex         CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnedReturnMethod = "fedex"
	CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnedReturnMethodOther         CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnedReturnMethod = "other"
	CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnedReturnMethodPostalService CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnedReturnMethod = "postal_service"
	CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnedReturnMethodUps           CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnedReturnMethod = "ups"
)

func (r CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnedReturnMethod) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnedReturnMethodDhl, CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnedReturnMethodFaceToFace, CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnedReturnMethodFedex, CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnedReturnMethodOther, CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnedReturnMethodPostalService, CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotAsDescribedReturnedReturnMethodUps:
		return true
	}
	return false
}

// Merchandise not received. Present if and only if `category` is
// `consumer_merchandise_not_received`.
type CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceived struct {
	// Cancellation outcome.
	CancellationOutcome CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedCancellationOutcome `json:"cancellation_outcome,required"`
	// Cardholder cancellation prior to expected receipt. Present if and only if
	// `cancellation_outcome` is `cardholder_cancellation_prior_to_expected_receipt`.
	CardholderCancellationPriorToExpectedReceipt CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedCardholderCancellationPriorToExpectedReceipt `json:"cardholder_cancellation_prior_to_expected_receipt,required,nullable"`
	// Delayed. Present if and only if `delivery_issue` is `delayed`.
	Delayed CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDelayed `json:"delayed,required,nullable"`
	// Delivered to wrong location. Present if and only if `delivery_issue` is
	// `delivered_to_wrong_location`.
	DeliveredToWrongLocation CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDeliveredToWrongLocation `json:"delivered_to_wrong_location,required,nullable"`
	// Delivery issue.
	DeliveryIssue CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDeliveryIssue `json:"delivery_issue,required"`
	// Last expected receipt at.
	LastExpectedReceiptAt time.Time `json:"last_expected_receipt_at,required" format:"date"`
	// Merchant cancellation. Present if and only if `cancellation_outcome` is
	// `merchant_cancellation`.
	MerchantCancellation CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedMerchantCancellation `json:"merchant_cancellation,required,nullable"`
	// Merchant resolution attempted.
	MerchantResolutionAttempted CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedMerchantResolutionAttempted `json:"merchant_resolution_attempted,required"`
	// No cancellation. Present if and only if `cancellation_outcome` is
	// `no_cancellation`.
	NoCancellation CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedNoCancellation `json:"no_cancellation,required,nullable"`
	// Purchase information and explanation.
	PurchaseInfoAndExplanation string                                                                     `json:"purchase_info_and_explanation,required"`
	JSON                       cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceived]
type cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedJSON struct {
	CancellationOutcome                          apijson.Field
	CardholderCancellationPriorToExpectedReceipt apijson.Field
	Delayed                                      apijson.Field
	DeliveredToWrongLocation                     apijson.Field
	DeliveryIssue                                apijson.Field
	LastExpectedReceiptAt                        apijson.Field
	MerchantCancellation                         apijson.Field
	MerchantResolutionAttempted                  apijson.Field
	NoCancellation                               apijson.Field
	PurchaseInfoAndExplanation                   apijson.Field
	raw                                          string
	ExtraFields                                  map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceived) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedJSON) RawJSON() string {
	return r.raw
}

// Cancellation outcome.
type CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedCancellationOutcome string

const (
	CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedCancellationOutcomeCardholderCancellationPriorToExpectedReceipt CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedCancellationOutcome = "cardholder_cancellation_prior_to_expected_receipt"
	CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedCancellationOutcomeMerchantCancellation                         CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedCancellationOutcome = "merchant_cancellation"
	CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedCancellationOutcomeNoCancellation                               CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedCancellationOutcome = "no_cancellation"
)

func (r CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedCancellationOutcome) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedCancellationOutcomeCardholderCancellationPriorToExpectedReceipt, CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedCancellationOutcomeMerchantCancellation, CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedCancellationOutcomeNoCancellation:
		return true
	}
	return false
}

// Cardholder cancellation prior to expected receipt. Present if and only if
// `cancellation_outcome` is `cardholder_cancellation_prior_to_expected_receipt`.
type CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedCardholderCancellationPriorToExpectedReceipt struct {
	// Canceled at.
	CanceledAt time.Time `json:"canceled_at,required" format:"date"`
	// Reason.
	Reason string                                                                                                                 `json:"reason,required,nullable"`
	JSON   cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedCardholderCancellationPriorToExpectedReceiptJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedCardholderCancellationPriorToExpectedReceiptJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedCardholderCancellationPriorToExpectedReceipt]
type cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedCardholderCancellationPriorToExpectedReceiptJSON struct {
	CanceledAt  apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedCardholderCancellationPriorToExpectedReceipt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedCardholderCancellationPriorToExpectedReceiptJSON) RawJSON() string {
	return r.raw
}

// Delayed. Present if and only if `delivery_issue` is `delayed`.
type CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDelayed struct {
	// Explanation.
	Explanation string `json:"explanation,required"`
	// Not returned. Present if and only if `return_outcome` is `not_returned`.
	NotReturned CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDelayedNotReturned `json:"not_returned,required,nullable"`
	// Return attempted. Present if and only if `return_outcome` is `return_attempted`.
	ReturnAttempted CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDelayedReturnAttempted `json:"return_attempted,required,nullable"`
	// Return outcome.
	ReturnOutcome CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDelayedReturnOutcome `json:"return_outcome,required"`
	// Returned. Present if and only if `return_outcome` is `returned`.
	Returned CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDelayedReturned `json:"returned,required,nullable"`
	JSON     cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDelayedJSON     `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDelayedJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDelayed]
type cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDelayedJSON struct {
	Explanation     apijson.Field
	NotReturned     apijson.Field
	ReturnAttempted apijson.Field
	ReturnOutcome   apijson.Field
	Returned        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDelayed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDelayedJSON) RawJSON() string {
	return r.raw
}

// Not returned. Present if and only if `return_outcome` is `not_returned`.
type CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDelayedNotReturned struct {
	JSON cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDelayedNotReturnedJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDelayedNotReturnedJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDelayedNotReturned]
type cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDelayedNotReturnedJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDelayedNotReturned) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDelayedNotReturnedJSON) RawJSON() string {
	return r.raw
}

// Return attempted. Present if and only if `return_outcome` is `return_attempted`.
type CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDelayedReturnAttempted struct {
	// Attempted at.
	AttemptedAt time.Time                                                                                        `json:"attempted_at,required" format:"date"`
	JSON        cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDelayedReturnAttemptedJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDelayedReturnAttemptedJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDelayedReturnAttempted]
type cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDelayedReturnAttemptedJSON struct {
	AttemptedAt apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDelayedReturnAttempted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDelayedReturnAttemptedJSON) RawJSON() string {
	return r.raw
}

// Return outcome.
type CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDelayedReturnOutcome string

const (
	CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDelayedReturnOutcomeNotReturned     CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDelayedReturnOutcome = "not_returned"
	CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDelayedReturnOutcomeReturned        CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDelayedReturnOutcome = "returned"
	CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDelayedReturnOutcomeReturnAttempted CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDelayedReturnOutcome = "return_attempted"
)

func (r CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDelayedReturnOutcome) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDelayedReturnOutcomeNotReturned, CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDelayedReturnOutcomeReturned, CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDelayedReturnOutcomeReturnAttempted:
		return true
	}
	return false
}

// Returned. Present if and only if `return_outcome` is `returned`.
type CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDelayedReturned struct {
	// Merchant received return at.
	MerchantReceivedReturnAt time.Time `json:"merchant_received_return_at,required" format:"date"`
	// Returned at.
	ReturnedAt time.Time                                                                                 `json:"returned_at,required" format:"date"`
	JSON       cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDelayedReturnedJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDelayedReturnedJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDelayedReturned]
type cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDelayedReturnedJSON struct {
	MerchantReceivedReturnAt apijson.Field
	ReturnedAt               apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDelayedReturned) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDelayedReturnedJSON) RawJSON() string {
	return r.raw
}

// Delivered to wrong location. Present if and only if `delivery_issue` is
// `delivered_to_wrong_location`.
type CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDeliveredToWrongLocation struct {
	// Agreed location.
	AgreedLocation string                                                                                             `json:"agreed_location,required"`
	JSON           cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDeliveredToWrongLocationJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDeliveredToWrongLocationJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDeliveredToWrongLocation]
type cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDeliveredToWrongLocationJSON struct {
	AgreedLocation apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDeliveredToWrongLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDeliveredToWrongLocationJSON) RawJSON() string {
	return r.raw
}

// Delivery issue.
type CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDeliveryIssue string

const (
	CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDeliveryIssueDelayed                  CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDeliveryIssue = "delayed"
	CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDeliveryIssueDeliveredToWrongLocation CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDeliveryIssue = "delivered_to_wrong_location"
)

func (r CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDeliveryIssue) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDeliveryIssueDelayed, CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedDeliveryIssueDeliveredToWrongLocation:
		return true
	}
	return false
}

// Merchant cancellation. Present if and only if `cancellation_outcome` is
// `merchant_cancellation`.
type CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedMerchantCancellation struct {
	// Canceled at.
	CanceledAt time.Time                                                                                      `json:"canceled_at,required" format:"date"`
	JSON       cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedMerchantCancellationJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedMerchantCancellationJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedMerchantCancellation]
type cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedMerchantCancellationJSON struct {
	CanceledAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedMerchantCancellation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedMerchantCancellationJSON) RawJSON() string {
	return r.raw
}

// Merchant resolution attempted.
type CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedMerchantResolutionAttempted string

const (
	CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedMerchantResolutionAttemptedAttempted            CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedMerchantResolutionAttempted = "attempted"
	CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedMerchantResolutionAttemptedProhibitedByLocalLaw CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedMerchantResolutionAttempted = "prohibited_by_local_law"
)

func (r CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedMerchantResolutionAttempted) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedMerchantResolutionAttemptedAttempted, CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedMerchantResolutionAttemptedProhibitedByLocalLaw:
		return true
	}
	return false
}

// No cancellation. Present if and only if `cancellation_outcome` is
// `no_cancellation`.
type CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedNoCancellation struct {
	JSON cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedNoCancellationJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedNoCancellationJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedNoCancellation]
type cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedNoCancellationJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedNoCancellation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerMerchandiseNotReceivedNoCancellationJSON) RawJSON() string {
	return r.raw
}

// Non-receipt of cash. Present if and only if `category` is
// `consumer_non_receipt_of_cash`.
type CardDisputeVisaUserSubmissionsChargebackConsumerNonReceiptOfCash struct {
	JSON cardDisputeVisaUserSubmissionsChargebackConsumerNonReceiptOfCashJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerNonReceiptOfCashJSON contains
// the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerNonReceiptOfCash]
type cardDisputeVisaUserSubmissionsChargebackConsumerNonReceiptOfCashJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerNonReceiptOfCash) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerNonReceiptOfCashJSON) RawJSON() string {
	return r.raw
}

// Original Credit Transaction (OCT) not accepted. Present if and only if
// `category` is `consumer_original_credit_transaction_not_accepted`.
type CardDisputeVisaUserSubmissionsChargebackConsumerOriginalCreditTransactionNotAccepted struct {
	// Explanation.
	Explanation string `json:"explanation,required"`
	// Reason.
	Reason CardDisputeVisaUserSubmissionsChargebackConsumerOriginalCreditTransactionNotAcceptedReason `json:"reason,required"`
	JSON   cardDisputeVisaUserSubmissionsChargebackConsumerOriginalCreditTransactionNotAcceptedJSON   `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerOriginalCreditTransactionNotAcceptedJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerOriginalCreditTransactionNotAccepted]
type cardDisputeVisaUserSubmissionsChargebackConsumerOriginalCreditTransactionNotAcceptedJSON struct {
	Explanation apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerOriginalCreditTransactionNotAccepted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerOriginalCreditTransactionNotAcceptedJSON) RawJSON() string {
	return r.raw
}

// Reason.
type CardDisputeVisaUserSubmissionsChargebackConsumerOriginalCreditTransactionNotAcceptedReason string

const (
	CardDisputeVisaUserSubmissionsChargebackConsumerOriginalCreditTransactionNotAcceptedReasonProhibitedByLocalLawsOrRegulation CardDisputeVisaUserSubmissionsChargebackConsumerOriginalCreditTransactionNotAcceptedReason = "prohibited_by_local_laws_or_regulation"
	CardDisputeVisaUserSubmissionsChargebackConsumerOriginalCreditTransactionNotAcceptedReasonRecipientRefused                  CardDisputeVisaUserSubmissionsChargebackConsumerOriginalCreditTransactionNotAcceptedReason = "recipient_refused"
)

func (r CardDisputeVisaUserSubmissionsChargebackConsumerOriginalCreditTransactionNotAcceptedReason) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsChargebackConsumerOriginalCreditTransactionNotAcceptedReasonProhibitedByLocalLawsOrRegulation, CardDisputeVisaUserSubmissionsChargebackConsumerOriginalCreditTransactionNotAcceptedReasonRecipientRefused:
		return true
	}
	return false
}

// Merchandise quality issue. Present if and only if `category` is
// `consumer_quality_merchandise`.
type CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandise struct {
	// Expected at.
	ExpectedAt time.Time `json:"expected_at,required" format:"date"`
	// Merchant resolution attempted.
	MerchantResolutionAttempted CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseMerchantResolutionAttempted `json:"merchant_resolution_attempted,required"`
	// Not returned. Present if and only if `return_outcome` is `not_returned`.
	NotReturned CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseNotReturned `json:"not_returned,required,nullable"`
	// Ongoing negotiations. Exclude if there is no evidence of ongoing negotiations.
	OngoingNegotiations CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseOngoingNegotiations `json:"ongoing_negotiations,required,nullable"`
	// Purchase information and quality issue.
	PurchaseInfoAndQualityIssue string `json:"purchase_info_and_quality_issue,required"`
	// Received at.
	ReceivedAt time.Time `json:"received_at,required" format:"date"`
	// Return attempted. Present if and only if `return_outcome` is `return_attempted`.
	ReturnAttempted CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnAttempted `json:"return_attempted,required,nullable"`
	// Return outcome.
	ReturnOutcome CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnOutcome `json:"return_outcome,required"`
	// Returned. Present if and only if `return_outcome` is `returned`.
	Returned CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturned `json:"returned,required,nullable"`
	JSON     cardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseJSON     `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseJSON contains
// the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandise]
type cardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseJSON struct {
	ExpectedAt                  apijson.Field
	MerchantResolutionAttempted apijson.Field
	NotReturned                 apijson.Field
	OngoingNegotiations         apijson.Field
	PurchaseInfoAndQualityIssue apijson.Field
	ReceivedAt                  apijson.Field
	ReturnAttempted             apijson.Field
	ReturnOutcome               apijson.Field
	Returned                    apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandise) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseJSON) RawJSON() string {
	return r.raw
}

// Merchant resolution attempted.
type CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseMerchantResolutionAttempted string

const (
	CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseMerchantResolutionAttemptedAttempted            CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseMerchantResolutionAttempted = "attempted"
	CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseMerchantResolutionAttemptedProhibitedByLocalLaw CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseMerchantResolutionAttempted = "prohibited_by_local_law"
)

func (r CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseMerchantResolutionAttempted) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseMerchantResolutionAttemptedAttempted, CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseMerchantResolutionAttemptedProhibitedByLocalLaw:
		return true
	}
	return false
}

// Not returned. Present if and only if `return_outcome` is `not_returned`.
type CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseNotReturned struct {
	JSON cardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseNotReturnedJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseNotReturnedJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseNotReturned]
type cardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseNotReturnedJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseNotReturned) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseNotReturnedJSON) RawJSON() string {
	return r.raw
}

// Ongoing negotiations. Exclude if there is no evidence of ongoing negotiations.
type CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseOngoingNegotiations struct {
	// Explanation of the previous ongoing negotiations between the cardholder and
	// merchant.
	Explanation string `json:"explanation,required"`
	// Date the cardholder first notified the issuer of the dispute.
	IssuerFirstNotifiedAt time.Time `json:"issuer_first_notified_at,required" format:"date"`
	// Started at.
	StartedAt time.Time                                                                                 `json:"started_at,required" format:"date"`
	JSON      cardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseOngoingNegotiationsJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseOngoingNegotiationsJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseOngoingNegotiations]
type cardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseOngoingNegotiationsJSON struct {
	Explanation           apijson.Field
	IssuerFirstNotifiedAt apijson.Field
	StartedAt             apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseOngoingNegotiations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseOngoingNegotiationsJSON) RawJSON() string {
	return r.raw
}

// Return attempted. Present if and only if `return_outcome` is `return_attempted`.
type CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnAttempted struct {
	// Attempt explanation.
	AttemptExplanation string `json:"attempt_explanation,required"`
	// Attempt reason.
	AttemptReason CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnAttemptedAttemptReason `json:"attempt_reason,required"`
	// Attempted at.
	AttemptedAt time.Time `json:"attempted_at,required" format:"date"`
	// Merchandise disposition.
	MerchandiseDisposition string                                                                                `json:"merchandise_disposition,required"`
	JSON                   cardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnAttemptedJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnAttemptedJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnAttempted]
type cardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnAttemptedJSON struct {
	AttemptExplanation     apijson.Field
	AttemptReason          apijson.Field
	AttemptedAt            apijson.Field
	MerchandiseDisposition apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnAttempted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnAttemptedJSON) RawJSON() string {
	return r.raw
}

// Attempt reason.
type CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnAttemptedAttemptReason string

const (
	CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnAttemptedAttemptReasonMerchantNotResponding         CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnAttemptedAttemptReason = "merchant_not_responding"
	CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnAttemptedAttemptReasonNoReturnAuthorizationProvided CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnAttemptedAttemptReason = "no_return_authorization_provided"
	CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnAttemptedAttemptReasonNoReturnInstructions          CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnAttemptedAttemptReason = "no_return_instructions"
	CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnAttemptedAttemptReasonRequestedNotToReturn          CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnAttemptedAttemptReason = "requested_not_to_return"
	CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnAttemptedAttemptReasonReturnNotAccepted             CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnAttemptedAttemptReason = "return_not_accepted"
)

func (r CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnAttemptedAttemptReason) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnAttemptedAttemptReasonMerchantNotResponding, CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnAttemptedAttemptReasonNoReturnAuthorizationProvided, CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnAttemptedAttemptReasonNoReturnInstructions, CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnAttemptedAttemptReasonRequestedNotToReturn, CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnAttemptedAttemptReasonReturnNotAccepted:
		return true
	}
	return false
}

// Return outcome.
type CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnOutcome string

const (
	CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnOutcomeNotReturned     CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnOutcome = "not_returned"
	CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnOutcomeReturned        CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnOutcome = "returned"
	CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnOutcomeReturnAttempted CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnOutcome = "return_attempted"
)

func (r CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnOutcome) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnOutcomeNotReturned, CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnOutcomeReturned, CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnOutcomeReturnAttempted:
		return true
	}
	return false
}

// Returned. Present if and only if `return_outcome` is `returned`.
type CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturned struct {
	// Merchant received return at.
	MerchantReceivedReturnAt time.Time `json:"merchant_received_return_at,required,nullable" format:"date"`
	// Other explanation. Required if and only if the return method is `other`.
	OtherExplanation string `json:"other_explanation,required,nullable"`
	// Return method.
	ReturnMethod CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnedReturnMethod `json:"return_method,required"`
	// Returned at.
	ReturnedAt time.Time `json:"returned_at,required" format:"date"`
	// Tracking number.
	TrackingNumber string                                                                         `json:"tracking_number,required,nullable"`
	JSON           cardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnedJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnedJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturned]
type cardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnedJSON struct {
	MerchantReceivedReturnAt apijson.Field
	OtherExplanation         apijson.Field
	ReturnMethod             apijson.Field
	ReturnedAt               apijson.Field
	TrackingNumber           apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturned) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnedJSON) RawJSON() string {
	return r.raw
}

// Return method.
type CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnedReturnMethod string

const (
	CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnedReturnMethodDhl           CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnedReturnMethod = "dhl"
	CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnedReturnMethodFaceToFace    CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnedReturnMethod = "face_to_face"
	CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnedReturnMethodFedex         CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnedReturnMethod = "fedex"
	CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnedReturnMethodOther         CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnedReturnMethod = "other"
	CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnedReturnMethodPostalService CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnedReturnMethod = "postal_service"
	CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnedReturnMethodUps           CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnedReturnMethod = "ups"
)

func (r CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnedReturnMethod) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnedReturnMethodDhl, CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnedReturnMethodFaceToFace, CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnedReturnMethodFedex, CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnedReturnMethodOther, CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnedReturnMethodPostalService, CardDisputeVisaUserSubmissionsChargebackConsumerQualityMerchandiseReturnedReturnMethodUps:
		return true
	}
	return false
}

// Services quality issue. Present if and only if `category` is
// `consumer_quality_services`.
type CardDisputeVisaUserSubmissionsChargebackConsumerQualityServices struct {
	// Cardholder cancellation.
	CardholderCancellation CardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesCardholderCancellation `json:"cardholder_cancellation,required"`
	// Cardholder paid to have work redone.
	CardholderPaidToHaveWorkRedone CardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesCardholderPaidToHaveWorkRedone `json:"cardholder_paid_to_have_work_redone,required,nullable"`
	// Non-fiat currency or non-fungible token related and not matching description.
	NonFiatCurrencyOrNonFungibleTokenRelatedAndNotMatchingDescription CardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesNonFiatCurrencyOrNonFungibleTokenRelatedAndNotMatchingDescription `json:"non_fiat_currency_or_non_fungible_token_related_and_not_matching_description,required"`
	// Ongoing negotiations. Exclude if there is no evidence of ongoing negotiations.
	OngoingNegotiations CardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesOngoingNegotiations `json:"ongoing_negotiations,required,nullable"`
	// Purchase information and quality issue.
	PurchaseInfoAndQualityIssue string `json:"purchase_info_and_quality_issue,required"`
	// Whether the dispute is related to the quality of food from an eating place or
	// restaurant. Must be provided when Merchant Category Code (MCC) is 5812, 5813
	// or 5814.
	RestaurantFoodRelated CardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesRestaurantFoodRelated `json:"restaurant_food_related,required,nullable"`
	// Services received at.
	ServicesReceivedAt time.Time                                                           `json:"services_received_at,required" format:"date"`
	JSON               cardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesJSON contains the
// JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerQualityServices]
type cardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesJSON struct {
	CardholderCancellation                                            apijson.Field
	CardholderPaidToHaveWorkRedone                                    apijson.Field
	NonFiatCurrencyOrNonFungibleTokenRelatedAndNotMatchingDescription apijson.Field
	OngoingNegotiations                                               apijson.Field
	PurchaseInfoAndQualityIssue                                       apijson.Field
	RestaurantFoodRelated                                             apijson.Field
	ServicesReceivedAt                                                apijson.Field
	raw                                                               string
	ExtraFields                                                       map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerQualityServices) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesJSON) RawJSON() string {
	return r.raw
}

// Cardholder cancellation.
type CardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesCardholderCancellation struct {
	// Accepted by merchant.
	AcceptedByMerchant CardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesCardholderCancellationAcceptedByMerchant `json:"accepted_by_merchant,required"`
	// Canceled at.
	CanceledAt time.Time `json:"canceled_at,required" format:"date"`
	// Reason.
	Reason string                                                                                    `json:"reason,required"`
	JSON   cardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesCardholderCancellationJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesCardholderCancellationJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesCardholderCancellation]
type cardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesCardholderCancellationJSON struct {
	AcceptedByMerchant apijson.Field
	CanceledAt         apijson.Field
	Reason             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesCardholderCancellation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesCardholderCancellationJSON) RawJSON() string {
	return r.raw
}

// Accepted by merchant.
type CardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesCardholderCancellationAcceptedByMerchant string

const (
	CardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesCardholderCancellationAcceptedByMerchantAccepted    CardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesCardholderCancellationAcceptedByMerchant = "accepted"
	CardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesCardholderCancellationAcceptedByMerchantNotAccepted CardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesCardholderCancellationAcceptedByMerchant = "not_accepted"
)

func (r CardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesCardholderCancellationAcceptedByMerchant) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesCardholderCancellationAcceptedByMerchantAccepted, CardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesCardholderCancellationAcceptedByMerchantNotAccepted:
		return true
	}
	return false
}

// Cardholder paid to have work redone.
type CardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesCardholderPaidToHaveWorkRedone string

const (
	CardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesCardholderPaidToHaveWorkRedoneDidNotPayToHaveWorkRedone CardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesCardholderPaidToHaveWorkRedone = "did_not_pay_to_have_work_redone"
	CardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesCardholderPaidToHaveWorkRedonePaidToHaveWorkRedone      CardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesCardholderPaidToHaveWorkRedone = "paid_to_have_work_redone"
)

func (r CardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesCardholderPaidToHaveWorkRedone) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesCardholderPaidToHaveWorkRedoneDidNotPayToHaveWorkRedone, CardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesCardholderPaidToHaveWorkRedonePaidToHaveWorkRedone:
		return true
	}
	return false
}

// Non-fiat currency or non-fungible token related and not matching description.
type CardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesNonFiatCurrencyOrNonFungibleTokenRelatedAndNotMatchingDescription string

const (
	CardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesNonFiatCurrencyOrNonFungibleTokenRelatedAndNotMatchingDescriptionNotRelated CardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesNonFiatCurrencyOrNonFungibleTokenRelatedAndNotMatchingDescription = "not_related"
	CardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesNonFiatCurrencyOrNonFungibleTokenRelatedAndNotMatchingDescriptionRelated    CardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesNonFiatCurrencyOrNonFungibleTokenRelatedAndNotMatchingDescription = "related"
)

func (r CardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesNonFiatCurrencyOrNonFungibleTokenRelatedAndNotMatchingDescription) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesNonFiatCurrencyOrNonFungibleTokenRelatedAndNotMatchingDescriptionNotRelated, CardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesNonFiatCurrencyOrNonFungibleTokenRelatedAndNotMatchingDescriptionRelated:
		return true
	}
	return false
}

// Ongoing negotiations. Exclude if there is no evidence of ongoing negotiations.
type CardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesOngoingNegotiations struct {
	// Explanation of the previous ongoing negotiations between the cardholder and
	// merchant.
	Explanation string `json:"explanation,required"`
	// Date the cardholder first notified the issuer of the dispute.
	IssuerFirstNotifiedAt time.Time `json:"issuer_first_notified_at,required" format:"date"`
	// Started at.
	StartedAt time.Time                                                                              `json:"started_at,required" format:"date"`
	JSON      cardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesOngoingNegotiationsJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesOngoingNegotiationsJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesOngoingNegotiations]
type cardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesOngoingNegotiationsJSON struct {
	Explanation           apijson.Field
	IssuerFirstNotifiedAt apijson.Field
	StartedAt             apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesOngoingNegotiations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesOngoingNegotiationsJSON) RawJSON() string {
	return r.raw
}

// Whether the dispute is related to the quality of food from an eating place or
// restaurant. Must be provided when Merchant Category Code (MCC) is 5812, 5813
// or 5814.
type CardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesRestaurantFoodRelated string

const (
	CardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesRestaurantFoodRelatedNotRelated CardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesRestaurantFoodRelated = "not_related"
	CardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesRestaurantFoodRelatedRelated    CardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesRestaurantFoodRelated = "related"
)

func (r CardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesRestaurantFoodRelated) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesRestaurantFoodRelatedNotRelated, CardDisputeVisaUserSubmissionsChargebackConsumerQualityServicesRestaurantFoodRelatedRelated:
		return true
	}
	return false
}

// Services misrepresentation. Present if and only if `category` is
// `consumer_services_misrepresentation`.
type CardDisputeVisaUserSubmissionsChargebackConsumerServicesMisrepresentation struct {
	// Cardholder cancellation.
	CardholderCancellation CardDisputeVisaUserSubmissionsChargebackConsumerServicesMisrepresentationCardholderCancellation `json:"cardholder_cancellation,required"`
	// Merchant resolution attempted.
	MerchantResolutionAttempted CardDisputeVisaUserSubmissionsChargebackConsumerServicesMisrepresentationMerchantResolutionAttempted `json:"merchant_resolution_attempted,required"`
	// Misrepresentation explanation.
	MisrepresentationExplanation string `json:"misrepresentation_explanation,required"`
	// Purchase explanation.
	PurchaseExplanation string `json:"purchase_explanation,required"`
	// Received at.
	ReceivedAt time.Time                                                                     `json:"received_at,required" format:"date"`
	JSON       cardDisputeVisaUserSubmissionsChargebackConsumerServicesMisrepresentationJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerServicesMisrepresentationJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerServicesMisrepresentation]
type cardDisputeVisaUserSubmissionsChargebackConsumerServicesMisrepresentationJSON struct {
	CardholderCancellation       apijson.Field
	MerchantResolutionAttempted  apijson.Field
	MisrepresentationExplanation apijson.Field
	PurchaseExplanation          apijson.Field
	ReceivedAt                   apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerServicesMisrepresentation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerServicesMisrepresentationJSON) RawJSON() string {
	return r.raw
}

// Cardholder cancellation.
type CardDisputeVisaUserSubmissionsChargebackConsumerServicesMisrepresentationCardholderCancellation struct {
	// Accepted by merchant.
	AcceptedByMerchant CardDisputeVisaUserSubmissionsChargebackConsumerServicesMisrepresentationCardholderCancellationAcceptedByMerchant `json:"accepted_by_merchant,required"`
	// Canceled at.
	CanceledAt time.Time `json:"canceled_at,required" format:"date"`
	// Reason.
	Reason string                                                                                              `json:"reason,required"`
	JSON   cardDisputeVisaUserSubmissionsChargebackConsumerServicesMisrepresentationCardholderCancellationJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerServicesMisrepresentationCardholderCancellationJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerServicesMisrepresentationCardholderCancellation]
type cardDisputeVisaUserSubmissionsChargebackConsumerServicesMisrepresentationCardholderCancellationJSON struct {
	AcceptedByMerchant apijson.Field
	CanceledAt         apijson.Field
	Reason             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerServicesMisrepresentationCardholderCancellation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerServicesMisrepresentationCardholderCancellationJSON) RawJSON() string {
	return r.raw
}

// Accepted by merchant.
type CardDisputeVisaUserSubmissionsChargebackConsumerServicesMisrepresentationCardholderCancellationAcceptedByMerchant string

const (
	CardDisputeVisaUserSubmissionsChargebackConsumerServicesMisrepresentationCardholderCancellationAcceptedByMerchantAccepted    CardDisputeVisaUserSubmissionsChargebackConsumerServicesMisrepresentationCardholderCancellationAcceptedByMerchant = "accepted"
	CardDisputeVisaUserSubmissionsChargebackConsumerServicesMisrepresentationCardholderCancellationAcceptedByMerchantNotAccepted CardDisputeVisaUserSubmissionsChargebackConsumerServicesMisrepresentationCardholderCancellationAcceptedByMerchant = "not_accepted"
)

func (r CardDisputeVisaUserSubmissionsChargebackConsumerServicesMisrepresentationCardholderCancellationAcceptedByMerchant) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsChargebackConsumerServicesMisrepresentationCardholderCancellationAcceptedByMerchantAccepted, CardDisputeVisaUserSubmissionsChargebackConsumerServicesMisrepresentationCardholderCancellationAcceptedByMerchantNotAccepted:
		return true
	}
	return false
}

// Merchant resolution attempted.
type CardDisputeVisaUserSubmissionsChargebackConsumerServicesMisrepresentationMerchantResolutionAttempted string

const (
	CardDisputeVisaUserSubmissionsChargebackConsumerServicesMisrepresentationMerchantResolutionAttemptedAttempted            CardDisputeVisaUserSubmissionsChargebackConsumerServicesMisrepresentationMerchantResolutionAttempted = "attempted"
	CardDisputeVisaUserSubmissionsChargebackConsumerServicesMisrepresentationMerchantResolutionAttemptedProhibitedByLocalLaw CardDisputeVisaUserSubmissionsChargebackConsumerServicesMisrepresentationMerchantResolutionAttempted = "prohibited_by_local_law"
)

func (r CardDisputeVisaUserSubmissionsChargebackConsumerServicesMisrepresentationMerchantResolutionAttempted) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsChargebackConsumerServicesMisrepresentationMerchantResolutionAttemptedAttempted, CardDisputeVisaUserSubmissionsChargebackConsumerServicesMisrepresentationMerchantResolutionAttemptedProhibitedByLocalLaw:
		return true
	}
	return false
}

// Services not as described. Present if and only if `category` is
// `consumer_services_not_as_described`.
type CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotAsDescribed struct {
	// Cardholder cancellation.
	CardholderCancellation CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotAsDescribedCardholderCancellation `json:"cardholder_cancellation,required"`
	// Explanation of what was ordered and was not as described.
	Explanation string `json:"explanation,required"`
	// Merchant resolution attempted.
	MerchantResolutionAttempted CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotAsDescribedMerchantResolutionAttempted `json:"merchant_resolution_attempted,required"`
	// Received at.
	ReceivedAt time.Time                                                                  `json:"received_at,required" format:"date"`
	JSON       cardDisputeVisaUserSubmissionsChargebackConsumerServicesNotAsDescribedJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerServicesNotAsDescribedJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotAsDescribed]
type cardDisputeVisaUserSubmissionsChargebackConsumerServicesNotAsDescribedJSON struct {
	CardholderCancellation      apijson.Field
	Explanation                 apijson.Field
	MerchantResolutionAttempted apijson.Field
	ReceivedAt                  apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotAsDescribed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerServicesNotAsDescribedJSON) RawJSON() string {
	return r.raw
}

// Cardholder cancellation.
type CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotAsDescribedCardholderCancellation struct {
	// Accepted by merchant.
	AcceptedByMerchant CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotAsDescribedCardholderCancellationAcceptedByMerchant `json:"accepted_by_merchant,required"`
	// Canceled at.
	CanceledAt time.Time `json:"canceled_at,required" format:"date"`
	// Reason.
	Reason string                                                                                           `json:"reason,required"`
	JSON   cardDisputeVisaUserSubmissionsChargebackConsumerServicesNotAsDescribedCardholderCancellationJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerServicesNotAsDescribedCardholderCancellationJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotAsDescribedCardholderCancellation]
type cardDisputeVisaUserSubmissionsChargebackConsumerServicesNotAsDescribedCardholderCancellationJSON struct {
	AcceptedByMerchant apijson.Field
	CanceledAt         apijson.Field
	Reason             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotAsDescribedCardholderCancellation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerServicesNotAsDescribedCardholderCancellationJSON) RawJSON() string {
	return r.raw
}

// Accepted by merchant.
type CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotAsDescribedCardholderCancellationAcceptedByMerchant string

const (
	CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotAsDescribedCardholderCancellationAcceptedByMerchantAccepted    CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotAsDescribedCardholderCancellationAcceptedByMerchant = "accepted"
	CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotAsDescribedCardholderCancellationAcceptedByMerchantNotAccepted CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotAsDescribedCardholderCancellationAcceptedByMerchant = "not_accepted"
)

func (r CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotAsDescribedCardholderCancellationAcceptedByMerchant) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotAsDescribedCardholderCancellationAcceptedByMerchantAccepted, CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotAsDescribedCardholderCancellationAcceptedByMerchantNotAccepted:
		return true
	}
	return false
}

// Merchant resolution attempted.
type CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotAsDescribedMerchantResolutionAttempted string

const (
	CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotAsDescribedMerchantResolutionAttemptedAttempted            CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotAsDescribedMerchantResolutionAttempted = "attempted"
	CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotAsDescribedMerchantResolutionAttemptedProhibitedByLocalLaw CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotAsDescribedMerchantResolutionAttempted = "prohibited_by_local_law"
)

func (r CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotAsDescribedMerchantResolutionAttempted) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotAsDescribedMerchantResolutionAttemptedAttempted, CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotAsDescribedMerchantResolutionAttemptedProhibitedByLocalLaw:
		return true
	}
	return false
}

// Services not received. Present if and only if `category` is
// `consumer_services_not_received`.
type CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceived struct {
	// Cancellation outcome.
	CancellationOutcome CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedCancellationOutcome `json:"cancellation_outcome,required"`
	// Cardholder cancellation prior to expected receipt. Present if and only if
	// `cancellation_outcome` is `cardholder_cancellation_prior_to_expected_receipt`.
	CardholderCancellationPriorToExpectedReceipt CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedCardholderCancellationPriorToExpectedReceipt `json:"cardholder_cancellation_prior_to_expected_receipt,required,nullable"`
	// Last expected receipt at.
	LastExpectedReceiptAt time.Time `json:"last_expected_receipt_at,required" format:"date"`
	// Merchant cancellation. Present if and only if `cancellation_outcome` is
	// `merchant_cancellation`.
	MerchantCancellation CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedMerchantCancellation `json:"merchant_cancellation,required,nullable"`
	// Merchant resolution attempted.
	MerchantResolutionAttempted CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedMerchantResolutionAttempted `json:"merchant_resolution_attempted,required"`
	// No cancellation. Present if and only if `cancellation_outcome` is
	// `no_cancellation`.
	NoCancellation CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedNoCancellation `json:"no_cancellation,required,nullable"`
	// Purchase information and explanation.
	PurchaseInfoAndExplanation string                                                                  `json:"purchase_info_and_explanation,required"`
	JSON                       cardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedJSON contains
// the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceived]
type cardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedJSON struct {
	CancellationOutcome                          apijson.Field
	CardholderCancellationPriorToExpectedReceipt apijson.Field
	LastExpectedReceiptAt                        apijson.Field
	MerchantCancellation                         apijson.Field
	MerchantResolutionAttempted                  apijson.Field
	NoCancellation                               apijson.Field
	PurchaseInfoAndExplanation                   apijson.Field
	raw                                          string
	ExtraFields                                  map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceived) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedJSON) RawJSON() string {
	return r.raw
}

// Cancellation outcome.
type CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedCancellationOutcome string

const (
	CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedCancellationOutcomeCardholderCancellationPriorToExpectedReceipt CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedCancellationOutcome = "cardholder_cancellation_prior_to_expected_receipt"
	CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedCancellationOutcomeMerchantCancellation                         CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedCancellationOutcome = "merchant_cancellation"
	CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedCancellationOutcomeNoCancellation                               CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedCancellationOutcome = "no_cancellation"
)

func (r CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedCancellationOutcome) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedCancellationOutcomeCardholderCancellationPriorToExpectedReceipt, CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedCancellationOutcomeMerchantCancellation, CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedCancellationOutcomeNoCancellation:
		return true
	}
	return false
}

// Cardholder cancellation prior to expected receipt. Present if and only if
// `cancellation_outcome` is `cardholder_cancellation_prior_to_expected_receipt`.
type CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedCardholderCancellationPriorToExpectedReceipt struct {
	// Canceled at.
	CanceledAt time.Time `json:"canceled_at,required" format:"date"`
	// Reason.
	Reason string                                                                                                              `json:"reason,required,nullable"`
	JSON   cardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedCardholderCancellationPriorToExpectedReceiptJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedCardholderCancellationPriorToExpectedReceiptJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedCardholderCancellationPriorToExpectedReceipt]
type cardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedCardholderCancellationPriorToExpectedReceiptJSON struct {
	CanceledAt  apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedCardholderCancellationPriorToExpectedReceipt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedCardholderCancellationPriorToExpectedReceiptJSON) RawJSON() string {
	return r.raw
}

// Merchant cancellation. Present if and only if `cancellation_outcome` is
// `merchant_cancellation`.
type CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedMerchantCancellation struct {
	// Canceled at.
	CanceledAt time.Time                                                                                   `json:"canceled_at,required" format:"date"`
	JSON       cardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedMerchantCancellationJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedMerchantCancellationJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedMerchantCancellation]
type cardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedMerchantCancellationJSON struct {
	CanceledAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedMerchantCancellation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedMerchantCancellationJSON) RawJSON() string {
	return r.raw
}

// Merchant resolution attempted.
type CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedMerchantResolutionAttempted string

const (
	CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedMerchantResolutionAttemptedAttempted            CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedMerchantResolutionAttempted = "attempted"
	CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedMerchantResolutionAttemptedProhibitedByLocalLaw CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedMerchantResolutionAttempted = "prohibited_by_local_law"
)

func (r CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedMerchantResolutionAttempted) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedMerchantResolutionAttemptedAttempted, CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedMerchantResolutionAttemptedProhibitedByLocalLaw:
		return true
	}
	return false
}

// No cancellation. Present if and only if `cancellation_outcome` is
// `no_cancellation`.
type CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedNoCancellation struct {
	JSON cardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedNoCancellationJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedNoCancellationJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedNoCancellation]
type cardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedNoCancellationJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedNoCancellation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackConsumerServicesNotReceivedNoCancellationJSON) RawJSON() string {
	return r.raw
}

// Fraud. Present if and only if `category` is `fraud`.
type CardDisputeVisaUserSubmissionsChargebackFraud struct {
	// Fraud type.
	FraudType CardDisputeVisaUserSubmissionsChargebackFraudFraudType `json:"fraud_type,required"`
	JSON      cardDisputeVisaUserSubmissionsChargebackFraudJSON      `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackFraudJSON contains the JSON metadata for
// the struct [CardDisputeVisaUserSubmissionsChargebackFraud]
type cardDisputeVisaUserSubmissionsChargebackFraudJSON struct {
	FraudType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackFraud) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackFraudJSON) RawJSON() string {
	return r.raw
}

// Fraud type.
type CardDisputeVisaUserSubmissionsChargebackFraudFraudType string

const (
	CardDisputeVisaUserSubmissionsChargebackFraudFraudTypeAccountOrCredentialsTakeover CardDisputeVisaUserSubmissionsChargebackFraudFraudType = "account_or_credentials_takeover"
	CardDisputeVisaUserSubmissionsChargebackFraudFraudTypeCardNotReceivedAsIssued      CardDisputeVisaUserSubmissionsChargebackFraudFraudType = "card_not_received_as_issued"
	CardDisputeVisaUserSubmissionsChargebackFraudFraudTypeFraudulentApplication        CardDisputeVisaUserSubmissionsChargebackFraudFraudType = "fraudulent_application"
	CardDisputeVisaUserSubmissionsChargebackFraudFraudTypeFraudulentUseOfAccountNumber CardDisputeVisaUserSubmissionsChargebackFraudFraudType = "fraudulent_use_of_account_number"
	CardDisputeVisaUserSubmissionsChargebackFraudFraudTypeIncorrectProcessing          CardDisputeVisaUserSubmissionsChargebackFraudFraudType = "incorrect_processing"
	CardDisputeVisaUserSubmissionsChargebackFraudFraudTypeIssuerReportedCounterfeit    CardDisputeVisaUserSubmissionsChargebackFraudFraudType = "issuer_reported_counterfeit"
	CardDisputeVisaUserSubmissionsChargebackFraudFraudTypeLost                         CardDisputeVisaUserSubmissionsChargebackFraudFraudType = "lost"
	CardDisputeVisaUserSubmissionsChargebackFraudFraudTypeManipulationOfAccountHolder  CardDisputeVisaUserSubmissionsChargebackFraudFraudType = "manipulation_of_account_holder"
	CardDisputeVisaUserSubmissionsChargebackFraudFraudTypeMerchantMisrepresentation    CardDisputeVisaUserSubmissionsChargebackFraudFraudType = "merchant_misrepresentation"
	CardDisputeVisaUserSubmissionsChargebackFraudFraudTypeMiscellaneous                CardDisputeVisaUserSubmissionsChargebackFraudFraudType = "miscellaneous"
	CardDisputeVisaUserSubmissionsChargebackFraudFraudTypeStolen                       CardDisputeVisaUserSubmissionsChargebackFraudFraudType = "stolen"
)

func (r CardDisputeVisaUserSubmissionsChargebackFraudFraudType) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsChargebackFraudFraudTypeAccountOrCredentialsTakeover, CardDisputeVisaUserSubmissionsChargebackFraudFraudTypeCardNotReceivedAsIssued, CardDisputeVisaUserSubmissionsChargebackFraudFraudTypeFraudulentApplication, CardDisputeVisaUserSubmissionsChargebackFraudFraudTypeFraudulentUseOfAccountNumber, CardDisputeVisaUserSubmissionsChargebackFraudFraudTypeIncorrectProcessing, CardDisputeVisaUserSubmissionsChargebackFraudFraudTypeIssuerReportedCounterfeit, CardDisputeVisaUserSubmissionsChargebackFraudFraudTypeLost, CardDisputeVisaUserSubmissionsChargebackFraudFraudTypeManipulationOfAccountHolder, CardDisputeVisaUserSubmissionsChargebackFraudFraudTypeMerchantMisrepresentation, CardDisputeVisaUserSubmissionsChargebackFraudFraudTypeMiscellaneous, CardDisputeVisaUserSubmissionsChargebackFraudFraudTypeStolen:
		return true
	}
	return false
}

// Processing error. Present if and only if `category` is `processing_error`.
type CardDisputeVisaUserSubmissionsChargebackProcessingError struct {
	// Duplicate transaction. Present if and only if `error_reason` is
	// `duplicate_transaction`.
	DuplicateTransaction CardDisputeVisaUserSubmissionsChargebackProcessingErrorDuplicateTransaction `json:"duplicate_transaction,required,nullable"`
	// Error reason.
	ErrorReason CardDisputeVisaUserSubmissionsChargebackProcessingErrorErrorReason `json:"error_reason,required"`
	// Incorrect amount. Present if and only if `error_reason` is `incorrect_amount`.
	IncorrectAmount CardDisputeVisaUserSubmissionsChargebackProcessingErrorIncorrectAmount `json:"incorrect_amount,required,nullable"`
	// Merchant resolution attempted.
	MerchantResolutionAttempted CardDisputeVisaUserSubmissionsChargebackProcessingErrorMerchantResolutionAttempted `json:"merchant_resolution_attempted,required"`
	// Paid by other means. Present if and only if `error_reason` is
	// `paid_by_other_means`.
	PaidByOtherMeans CardDisputeVisaUserSubmissionsChargebackProcessingErrorPaidByOtherMeans `json:"paid_by_other_means,required,nullable"`
	JSON             cardDisputeVisaUserSubmissionsChargebackProcessingErrorJSON             `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackProcessingErrorJSON contains the JSON
// metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackProcessingError]
type cardDisputeVisaUserSubmissionsChargebackProcessingErrorJSON struct {
	DuplicateTransaction        apijson.Field
	ErrorReason                 apijson.Field
	IncorrectAmount             apijson.Field
	MerchantResolutionAttempted apijson.Field
	PaidByOtherMeans            apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackProcessingError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackProcessingErrorJSON) RawJSON() string {
	return r.raw
}

// Duplicate transaction. Present if and only if `error_reason` is
// `duplicate_transaction`.
type CardDisputeVisaUserSubmissionsChargebackProcessingErrorDuplicateTransaction struct {
	// Other transaction ID.
	OtherTransactionID string                                                                          `json:"other_transaction_id,required"`
	JSON               cardDisputeVisaUserSubmissionsChargebackProcessingErrorDuplicateTransactionJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackProcessingErrorDuplicateTransactionJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackProcessingErrorDuplicateTransaction]
type cardDisputeVisaUserSubmissionsChargebackProcessingErrorDuplicateTransactionJSON struct {
	OtherTransactionID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackProcessingErrorDuplicateTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackProcessingErrorDuplicateTransactionJSON) RawJSON() string {
	return r.raw
}

// Error reason.
type CardDisputeVisaUserSubmissionsChargebackProcessingErrorErrorReason string

const (
	CardDisputeVisaUserSubmissionsChargebackProcessingErrorErrorReasonDuplicateTransaction CardDisputeVisaUserSubmissionsChargebackProcessingErrorErrorReason = "duplicate_transaction"
	CardDisputeVisaUserSubmissionsChargebackProcessingErrorErrorReasonIncorrectAmount      CardDisputeVisaUserSubmissionsChargebackProcessingErrorErrorReason = "incorrect_amount"
	CardDisputeVisaUserSubmissionsChargebackProcessingErrorErrorReasonPaidByOtherMeans     CardDisputeVisaUserSubmissionsChargebackProcessingErrorErrorReason = "paid_by_other_means"
)

func (r CardDisputeVisaUserSubmissionsChargebackProcessingErrorErrorReason) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsChargebackProcessingErrorErrorReasonDuplicateTransaction, CardDisputeVisaUserSubmissionsChargebackProcessingErrorErrorReasonIncorrectAmount, CardDisputeVisaUserSubmissionsChargebackProcessingErrorErrorReasonPaidByOtherMeans:
		return true
	}
	return false
}

// Incorrect amount. Present if and only if `error_reason` is `incorrect_amount`.
type CardDisputeVisaUserSubmissionsChargebackProcessingErrorIncorrectAmount struct {
	// Expected amount.
	ExpectedAmount int64                                                                      `json:"expected_amount,required"`
	JSON           cardDisputeVisaUserSubmissionsChargebackProcessingErrorIncorrectAmountJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackProcessingErrorIncorrectAmountJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackProcessingErrorIncorrectAmount]
type cardDisputeVisaUserSubmissionsChargebackProcessingErrorIncorrectAmountJSON struct {
	ExpectedAmount apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackProcessingErrorIncorrectAmount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackProcessingErrorIncorrectAmountJSON) RawJSON() string {
	return r.raw
}

// Merchant resolution attempted.
type CardDisputeVisaUserSubmissionsChargebackProcessingErrorMerchantResolutionAttempted string

const (
	CardDisputeVisaUserSubmissionsChargebackProcessingErrorMerchantResolutionAttemptedAttempted            CardDisputeVisaUserSubmissionsChargebackProcessingErrorMerchantResolutionAttempted = "attempted"
	CardDisputeVisaUserSubmissionsChargebackProcessingErrorMerchantResolutionAttemptedProhibitedByLocalLaw CardDisputeVisaUserSubmissionsChargebackProcessingErrorMerchantResolutionAttempted = "prohibited_by_local_law"
)

func (r CardDisputeVisaUserSubmissionsChargebackProcessingErrorMerchantResolutionAttempted) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsChargebackProcessingErrorMerchantResolutionAttemptedAttempted, CardDisputeVisaUserSubmissionsChargebackProcessingErrorMerchantResolutionAttemptedProhibitedByLocalLaw:
		return true
	}
	return false
}

// Paid by other means. Present if and only if `error_reason` is
// `paid_by_other_means`.
type CardDisputeVisaUserSubmissionsChargebackProcessingErrorPaidByOtherMeans struct {
	// Other form of payment evidence.
	OtherFormOfPaymentEvidence CardDisputeVisaUserSubmissionsChargebackProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidence `json:"other_form_of_payment_evidence,required"`
	// Other transaction ID.
	OtherTransactionID string                                                                      `json:"other_transaction_id,required,nullable"`
	JSON               cardDisputeVisaUserSubmissionsChargebackProcessingErrorPaidByOtherMeansJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsChargebackProcessingErrorPaidByOtherMeansJSON
// contains the JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsChargebackProcessingErrorPaidByOtherMeans]
type cardDisputeVisaUserSubmissionsChargebackProcessingErrorPaidByOtherMeansJSON struct {
	OtherFormOfPaymentEvidence apijson.Field
	OtherTransactionID         apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsChargebackProcessingErrorPaidByOtherMeans) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsChargebackProcessingErrorPaidByOtherMeansJSON) RawJSON() string {
	return r.raw
}

// Other form of payment evidence.
type CardDisputeVisaUserSubmissionsChargebackProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidence string

const (
	CardDisputeVisaUserSubmissionsChargebackProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidenceCanceledCheck   CardDisputeVisaUserSubmissionsChargebackProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidence = "canceled_check"
	CardDisputeVisaUserSubmissionsChargebackProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidenceCardTransaction CardDisputeVisaUserSubmissionsChargebackProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidence = "card_transaction"
	CardDisputeVisaUserSubmissionsChargebackProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidenceCashReceipt     CardDisputeVisaUserSubmissionsChargebackProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidence = "cash_receipt"
	CardDisputeVisaUserSubmissionsChargebackProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidenceOther           CardDisputeVisaUserSubmissionsChargebackProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidence = "other"
	CardDisputeVisaUserSubmissionsChargebackProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidenceStatement       CardDisputeVisaUserSubmissionsChargebackProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidence = "statement"
	CardDisputeVisaUserSubmissionsChargebackProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidenceVoucher         CardDisputeVisaUserSubmissionsChargebackProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidence = "voucher"
)

func (r CardDisputeVisaUserSubmissionsChargebackProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidence) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsChargebackProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidenceCanceledCheck, CardDisputeVisaUserSubmissionsChargebackProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidenceCardTransaction, CardDisputeVisaUserSubmissionsChargebackProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidenceCashReceipt, CardDisputeVisaUserSubmissionsChargebackProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidenceOther, CardDisputeVisaUserSubmissionsChargebackProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidenceStatement, CardDisputeVisaUserSubmissionsChargebackProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidenceVoucher:
		return true
	}
	return false
}

// A Visa Card Dispute Merchant Pre-Arbitration Decline User Submission object.
// This field will be present in the JSON response if and only if `category` is
// equal to `merchant_prearbitration_decline`. Contains the details specific to a
// merchant prearbitration decline Visa Card Dispute User Submission.
type CardDisputeVisaUserSubmissionsMerchantPrearbitrationDecline struct {
	// The reason the user declined the merchant's request for pre-arbitration in their
	// favor.
	Reason string                                                          `json:"reason,required"`
	JSON   cardDisputeVisaUserSubmissionsMerchantPrearbitrationDeclineJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsMerchantPrearbitrationDeclineJSON contains the
// JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsMerchantPrearbitrationDecline]
type cardDisputeVisaUserSubmissionsMerchantPrearbitrationDeclineJSON struct {
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsMerchantPrearbitrationDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsMerchantPrearbitrationDeclineJSON) RawJSON() string {
	return r.raw
}

// A Visa Card Dispute User-Initiated Pre-Arbitration User Submission object. This
// field will be present in the JSON response if and only if `category` is equal to
// `user_prearbitration`. Contains the details specific to a user-initiated
// pre-arbitration Visa Card Dispute User Submission.
type CardDisputeVisaUserSubmissionsUserPrearbitration struct {
	// Category change details for the pre-arbitration request, if requested.
	CategoryChange CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChange `json:"category_change,required,nullable"`
	// The reason for the pre-arbitration request.
	Reason string                                               `json:"reason,required"`
	JSON   cardDisputeVisaUserSubmissionsUserPrearbitrationJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsUserPrearbitrationJSON contains the JSON metadata
// for the struct [CardDisputeVisaUserSubmissionsUserPrearbitration]
type cardDisputeVisaUserSubmissionsUserPrearbitrationJSON struct {
	CategoryChange apijson.Field
	Reason         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsUserPrearbitration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsUserPrearbitrationJSON) RawJSON() string {
	return r.raw
}

// Category change details for the pre-arbitration request, if requested.
type CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChange struct {
	// The category the dispute is being changed to.
	Category CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategory `json:"category,required"`
	// The reason for the pre-arbitration request.
	Reason string                                                             `json:"reason,required"`
	JSON   cardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeJSON `json:"-"`
}

// cardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeJSON contains the
// JSON metadata for the struct
// [CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChange]
type cardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeJSON struct {
	Category    apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeJSON) RawJSON() string {
	return r.raw
}

// The category the dispute is being changed to.
type CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategory string

const (
	CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategoryAuthorization                                CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategory = "authorization"
	CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategoryConsumerCanceledMerchandise                  CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategory = "consumer_canceled_merchandise"
	CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategoryConsumerCanceledRecurringTransaction         CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategory = "consumer_canceled_recurring_transaction"
	CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategoryConsumerCanceledServices                     CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategory = "consumer_canceled_services"
	CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategoryConsumerCounterfeitMerchandise               CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategory = "consumer_counterfeit_merchandise"
	CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategoryConsumerCreditNotProcessed                   CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategory = "consumer_credit_not_processed"
	CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategoryConsumerDamagedOrDefectiveMerchandise        CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategory = "consumer_damaged_or_defective_merchandise"
	CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategoryConsumerMerchandiseMisrepresentation         CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategory = "consumer_merchandise_misrepresentation"
	CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategoryConsumerMerchandiseNotAsDescribed            CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategory = "consumer_merchandise_not_as_described"
	CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategoryConsumerMerchandiseNotReceived               CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategory = "consumer_merchandise_not_received"
	CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategoryConsumerNonReceiptOfCash                     CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategory = "consumer_non_receipt_of_cash"
	CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategoryConsumerOriginalCreditTransactionNotAccepted CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategory = "consumer_original_credit_transaction_not_accepted"
	CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategoryConsumerQualityMerchandise                   CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategory = "consumer_quality_merchandise"
	CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategoryConsumerQualityServices                      CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategory = "consumer_quality_services"
	CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategoryConsumerServicesMisrepresentation            CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategory = "consumer_services_misrepresentation"
	CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategoryConsumerServicesNotAsDescribed               CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategory = "consumer_services_not_as_described"
	CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategoryConsumerServicesNotReceived                  CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategory = "consumer_services_not_received"
	CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategoryFraud                                        CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategory = "fraud"
	CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategoryProcessingError                              CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategory = "processing_error"
)

func (r CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategory) IsKnown() bool {
	switch r {
	case CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategoryAuthorization, CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategoryConsumerCanceledMerchandise, CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategoryConsumerCanceledRecurringTransaction, CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategoryConsumerCanceledServices, CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategoryConsumerCounterfeitMerchandise, CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategoryConsumerCreditNotProcessed, CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategoryConsumerDamagedOrDefectiveMerchandise, CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategoryConsumerMerchandiseMisrepresentation, CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategoryConsumerMerchandiseNotAsDescribed, CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategoryConsumerMerchandiseNotReceived, CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategoryConsumerNonReceiptOfCash, CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategoryConsumerOriginalCreditTransactionNotAccepted, CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategoryConsumerQualityMerchandise, CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategoryConsumerQualityServices, CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategoryConsumerServicesMisrepresentation, CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategoryConsumerServicesNotAsDescribed, CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategoryConsumerServicesNotReceived, CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategoryFraud, CardDisputeVisaUserSubmissionsUserPrearbitrationCategoryChangeCategoryProcessingError:
		return true
	}
	return false
}

// If the Card Dispute's status is `won`, this will contain details of the won
// dispute.
type CardDisputeWin struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card Dispute was won.
	WonAt time.Time          `json:"won_at,required" format:"date-time"`
	JSON  cardDisputeWinJSON `json:"-"`
}

// cardDisputeWinJSON contains the JSON metadata for the struct [CardDisputeWin]
type cardDisputeWinJSON struct {
	WonAt       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeWin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeWinJSON) RawJSON() string {
	return r.raw
}

// If the Card Dispute has been withdrawn, this will contain details of the
// withdrawal.
type CardDisputeWithdrawal struct {
	// The explanation for the withdrawal of the Card Dispute.
	Explanation string                    `json:"explanation,required,nullable"`
	JSON        cardDisputeWithdrawalJSON `json:"-"`
}

// cardDisputeWithdrawalJSON contains the JSON metadata for the struct
// [CardDisputeWithdrawal]
type cardDisputeWithdrawalJSON struct {
	Explanation apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardDisputeWithdrawal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardDisputeWithdrawalJSON) RawJSON() string {
	return r.raw
}

type CardDisputeNewParams struct {
	// The Transaction you wish to dispute. This Transaction must have a `source_type`
	// of `card_settlement`.
	DisputedTransactionID param.Field[string] `json:"disputed_transaction_id,required"`
	// The network of the disputed transaction. Details specific to the network are
	// required under the sub-object with the same identifier as the network.
	Network param.Field[CardDisputeNewParamsNetwork] `json:"network,required"`
	// The monetary amount of the part of the transaction that is being disputed. This
	// is optional and will default to the full amount of the transaction if not
	// provided. If provided, the amount must be less than or equal to the amount of
	// the transaction.
	Amount param.Field[int64] `json:"amount"`
	// The files to be attached to the initial dispute submission.
	AttachmentFiles param.Field[[]CardDisputeNewParamsAttachmentFile] `json:"attachment_files"`
	// The free-form explanation provided to Increase to provide more context for the
	// user submission. This field is not sent directly to the card networks.
	Explanation param.Field[string] `json:"explanation"`
	// The Visa-specific parameters for the dispute. Required if and only if `network`
	// is `visa`.
	Visa param.Field[CardDisputeNewParamsVisa] `json:"visa"`
}

func (r CardDisputeNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The network of the disputed transaction. Details specific to the network are
// required under the sub-object with the same identifier as the network.
type CardDisputeNewParamsNetwork string

const (
	CardDisputeNewParamsNetworkVisa CardDisputeNewParamsNetwork = "visa"
)

func (r CardDisputeNewParamsNetwork) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsNetworkVisa:
		return true
	}
	return false
}

type CardDisputeNewParamsAttachmentFile struct {
	// The ID of the file to be attached. The file must have a `purpose` of
	// `card_dispute_attachment`.
	FileID param.Field[string] `json:"file_id,required"`
}

func (r CardDisputeNewParamsAttachmentFile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The Visa-specific parameters for the dispute. Required if and only if `network`
// is `visa`.
type CardDisputeNewParamsVisa struct {
	// Category.
	Category param.Field[CardDisputeNewParamsVisaCategory] `json:"category,required"`
	// Authorization. Required if and only if `category` is `authorization`.
	Authorization param.Field[CardDisputeNewParamsVisaAuthorization] `json:"authorization"`
	// Canceled merchandise. Required if and only if `category` is
	// `consumer_canceled_merchandise`.
	ConsumerCanceledMerchandise param.Field[CardDisputeNewParamsVisaConsumerCanceledMerchandise] `json:"consumer_canceled_merchandise"`
	// Canceled recurring transaction. Required if and only if `category` is
	// `consumer_canceled_recurring_transaction`.
	ConsumerCanceledRecurringTransaction param.Field[CardDisputeNewParamsVisaConsumerCanceledRecurringTransaction] `json:"consumer_canceled_recurring_transaction"`
	// Canceled services. Required if and only if `category` is
	// `consumer_canceled_services`.
	ConsumerCanceledServices param.Field[CardDisputeNewParamsVisaConsumerCanceledServices] `json:"consumer_canceled_services"`
	// Counterfeit merchandise. Required if and only if `category` is
	// `consumer_counterfeit_merchandise`.
	ConsumerCounterfeitMerchandise param.Field[CardDisputeNewParamsVisaConsumerCounterfeitMerchandise] `json:"consumer_counterfeit_merchandise"`
	// Credit not processed. Required if and only if `category` is
	// `consumer_credit_not_processed`.
	ConsumerCreditNotProcessed param.Field[CardDisputeNewParamsVisaConsumerCreditNotProcessed] `json:"consumer_credit_not_processed"`
	// Damaged or defective merchandise. Required if and only if `category` is
	// `consumer_damaged_or_defective_merchandise`.
	ConsumerDamagedOrDefectiveMerchandise param.Field[CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandise] `json:"consumer_damaged_or_defective_merchandise"`
	// Merchandise misrepresentation. Required if and only if `category` is
	// `consumer_merchandise_misrepresentation`.
	ConsumerMerchandiseMisrepresentation param.Field[CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentation] `json:"consumer_merchandise_misrepresentation"`
	// Merchandise not as described. Required if and only if `category` is
	// `consumer_merchandise_not_as_described`.
	ConsumerMerchandiseNotAsDescribed param.Field[CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribed] `json:"consumer_merchandise_not_as_described"`
	// Merchandise not received. Required if and only if `category` is
	// `consumer_merchandise_not_received`.
	ConsumerMerchandiseNotReceived param.Field[CardDisputeNewParamsVisaConsumerMerchandiseNotReceived] `json:"consumer_merchandise_not_received"`
	// Non-receipt of cash. Required if and only if `category` is
	// `consumer_non_receipt_of_cash`.
	ConsumerNonReceiptOfCash param.Field[CardDisputeNewParamsVisaConsumerNonReceiptOfCash] `json:"consumer_non_receipt_of_cash"`
	// Original Credit Transaction (OCT) not accepted. Required if and only if
	// `category` is `consumer_original_credit_transaction_not_accepted`.
	ConsumerOriginalCreditTransactionNotAccepted param.Field[CardDisputeNewParamsVisaConsumerOriginalCreditTransactionNotAccepted] `json:"consumer_original_credit_transaction_not_accepted"`
	// Merchandise quality issue. Required if and only if `category` is
	// `consumer_quality_merchandise`.
	ConsumerQualityMerchandise param.Field[CardDisputeNewParamsVisaConsumerQualityMerchandise] `json:"consumer_quality_merchandise"`
	// Services quality issue. Required if and only if `category` is
	// `consumer_quality_services`.
	ConsumerQualityServices param.Field[CardDisputeNewParamsVisaConsumerQualityServices] `json:"consumer_quality_services"`
	// Services misrepresentation. Required if and only if `category` is
	// `consumer_services_misrepresentation`.
	ConsumerServicesMisrepresentation param.Field[CardDisputeNewParamsVisaConsumerServicesMisrepresentation] `json:"consumer_services_misrepresentation"`
	// Services not as described. Required if and only if `category` is
	// `consumer_services_not_as_described`.
	ConsumerServicesNotAsDescribed param.Field[CardDisputeNewParamsVisaConsumerServicesNotAsDescribed] `json:"consumer_services_not_as_described"`
	// Services not received. Required if and only if `category` is
	// `consumer_services_not_received`.
	ConsumerServicesNotReceived param.Field[CardDisputeNewParamsVisaConsumerServicesNotReceived] `json:"consumer_services_not_received"`
	// Fraud. Required if and only if `category` is `fraud`.
	Fraud param.Field[CardDisputeNewParamsVisaFraud] `json:"fraud"`
	// Processing error. Required if and only if `category` is `processing_error`.
	ProcessingError param.Field[CardDisputeNewParamsVisaProcessingError] `json:"processing_error"`
}

func (r CardDisputeNewParamsVisa) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Category.
type CardDisputeNewParamsVisaCategory string

const (
	CardDisputeNewParamsVisaCategoryAuthorization                                CardDisputeNewParamsVisaCategory = "authorization"
	CardDisputeNewParamsVisaCategoryConsumerCanceledMerchandise                  CardDisputeNewParamsVisaCategory = "consumer_canceled_merchandise"
	CardDisputeNewParamsVisaCategoryConsumerCanceledRecurringTransaction         CardDisputeNewParamsVisaCategory = "consumer_canceled_recurring_transaction"
	CardDisputeNewParamsVisaCategoryConsumerCanceledServices                     CardDisputeNewParamsVisaCategory = "consumer_canceled_services"
	CardDisputeNewParamsVisaCategoryConsumerCounterfeitMerchandise               CardDisputeNewParamsVisaCategory = "consumer_counterfeit_merchandise"
	CardDisputeNewParamsVisaCategoryConsumerCreditNotProcessed                   CardDisputeNewParamsVisaCategory = "consumer_credit_not_processed"
	CardDisputeNewParamsVisaCategoryConsumerDamagedOrDefectiveMerchandise        CardDisputeNewParamsVisaCategory = "consumer_damaged_or_defective_merchandise"
	CardDisputeNewParamsVisaCategoryConsumerMerchandiseMisrepresentation         CardDisputeNewParamsVisaCategory = "consumer_merchandise_misrepresentation"
	CardDisputeNewParamsVisaCategoryConsumerMerchandiseNotAsDescribed            CardDisputeNewParamsVisaCategory = "consumer_merchandise_not_as_described"
	CardDisputeNewParamsVisaCategoryConsumerMerchandiseNotReceived               CardDisputeNewParamsVisaCategory = "consumer_merchandise_not_received"
	CardDisputeNewParamsVisaCategoryConsumerNonReceiptOfCash                     CardDisputeNewParamsVisaCategory = "consumer_non_receipt_of_cash"
	CardDisputeNewParamsVisaCategoryConsumerOriginalCreditTransactionNotAccepted CardDisputeNewParamsVisaCategory = "consumer_original_credit_transaction_not_accepted"
	CardDisputeNewParamsVisaCategoryConsumerQualityMerchandise                   CardDisputeNewParamsVisaCategory = "consumer_quality_merchandise"
	CardDisputeNewParamsVisaCategoryConsumerQualityServices                      CardDisputeNewParamsVisaCategory = "consumer_quality_services"
	CardDisputeNewParamsVisaCategoryConsumerServicesMisrepresentation            CardDisputeNewParamsVisaCategory = "consumer_services_misrepresentation"
	CardDisputeNewParamsVisaCategoryConsumerServicesNotAsDescribed               CardDisputeNewParamsVisaCategory = "consumer_services_not_as_described"
	CardDisputeNewParamsVisaCategoryConsumerServicesNotReceived                  CardDisputeNewParamsVisaCategory = "consumer_services_not_received"
	CardDisputeNewParamsVisaCategoryFraud                                        CardDisputeNewParamsVisaCategory = "fraud"
	CardDisputeNewParamsVisaCategoryProcessingError                              CardDisputeNewParamsVisaCategory = "processing_error"
)

func (r CardDisputeNewParamsVisaCategory) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsVisaCategoryAuthorization, CardDisputeNewParamsVisaCategoryConsumerCanceledMerchandise, CardDisputeNewParamsVisaCategoryConsumerCanceledRecurringTransaction, CardDisputeNewParamsVisaCategoryConsumerCanceledServices, CardDisputeNewParamsVisaCategoryConsumerCounterfeitMerchandise, CardDisputeNewParamsVisaCategoryConsumerCreditNotProcessed, CardDisputeNewParamsVisaCategoryConsumerDamagedOrDefectiveMerchandise, CardDisputeNewParamsVisaCategoryConsumerMerchandiseMisrepresentation, CardDisputeNewParamsVisaCategoryConsumerMerchandiseNotAsDescribed, CardDisputeNewParamsVisaCategoryConsumerMerchandiseNotReceived, CardDisputeNewParamsVisaCategoryConsumerNonReceiptOfCash, CardDisputeNewParamsVisaCategoryConsumerOriginalCreditTransactionNotAccepted, CardDisputeNewParamsVisaCategoryConsumerQualityMerchandise, CardDisputeNewParamsVisaCategoryConsumerQualityServices, CardDisputeNewParamsVisaCategoryConsumerServicesMisrepresentation, CardDisputeNewParamsVisaCategoryConsumerServicesNotAsDescribed, CardDisputeNewParamsVisaCategoryConsumerServicesNotReceived, CardDisputeNewParamsVisaCategoryFraud, CardDisputeNewParamsVisaCategoryProcessingError:
		return true
	}
	return false
}

// Authorization. Required if and only if `category` is `authorization`.
type CardDisputeNewParamsVisaAuthorization struct {
	// Account status.
	AccountStatus param.Field[CardDisputeNewParamsVisaAuthorizationAccountStatus] `json:"account_status,required"`
}

func (r CardDisputeNewParamsVisaAuthorization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Account status.
type CardDisputeNewParamsVisaAuthorizationAccountStatus string

const (
	CardDisputeNewParamsVisaAuthorizationAccountStatusAccountClosed CardDisputeNewParamsVisaAuthorizationAccountStatus = "account_closed"
	CardDisputeNewParamsVisaAuthorizationAccountStatusCreditProblem CardDisputeNewParamsVisaAuthorizationAccountStatus = "credit_problem"
	CardDisputeNewParamsVisaAuthorizationAccountStatusFraud         CardDisputeNewParamsVisaAuthorizationAccountStatus = "fraud"
)

func (r CardDisputeNewParamsVisaAuthorizationAccountStatus) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsVisaAuthorizationAccountStatusAccountClosed, CardDisputeNewParamsVisaAuthorizationAccountStatusCreditProblem, CardDisputeNewParamsVisaAuthorizationAccountStatusFraud:
		return true
	}
	return false
}

// Canceled merchandise. Required if and only if `category` is
// `consumer_canceled_merchandise`.
type CardDisputeNewParamsVisaConsumerCanceledMerchandise struct {
	// Merchant resolution attempted.
	MerchantResolutionAttempted param.Field[CardDisputeNewParamsVisaConsumerCanceledMerchandiseMerchantResolutionAttempted] `json:"merchant_resolution_attempted,required"`
	// Purchase explanation.
	PurchaseExplanation param.Field[string] `json:"purchase_explanation,required"`
	// Received or expected at.
	ReceivedOrExpectedAt param.Field[time.Time] `json:"received_or_expected_at,required" format:"date"`
	// Return outcome.
	ReturnOutcome param.Field[CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnOutcome] `json:"return_outcome,required"`
	// Cardholder cancellation.
	CardholderCancellation param.Field[CardDisputeNewParamsVisaConsumerCanceledMerchandiseCardholderCancellation] `json:"cardholder_cancellation"`
	// Not returned. Required if and only if `return_outcome` is `not_returned`.
	NotReturned param.Field[CardDisputeNewParamsVisaConsumerCanceledMerchandiseNotReturned] `json:"not_returned"`
	// Return attempted. Required if and only if `return_outcome` is
	// `return_attempted`.
	ReturnAttempted param.Field[CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnAttempted] `json:"return_attempted"`
	// Returned. Required if and only if `return_outcome` is `returned`.
	Returned param.Field[CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturned] `json:"returned"`
}

func (r CardDisputeNewParamsVisaConsumerCanceledMerchandise) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Merchant resolution attempted.
type CardDisputeNewParamsVisaConsumerCanceledMerchandiseMerchantResolutionAttempted string

const (
	CardDisputeNewParamsVisaConsumerCanceledMerchandiseMerchantResolutionAttemptedAttempted            CardDisputeNewParamsVisaConsumerCanceledMerchandiseMerchantResolutionAttempted = "attempted"
	CardDisputeNewParamsVisaConsumerCanceledMerchandiseMerchantResolutionAttemptedProhibitedByLocalLaw CardDisputeNewParamsVisaConsumerCanceledMerchandiseMerchantResolutionAttempted = "prohibited_by_local_law"
)

func (r CardDisputeNewParamsVisaConsumerCanceledMerchandiseMerchantResolutionAttempted) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsVisaConsumerCanceledMerchandiseMerchantResolutionAttemptedAttempted, CardDisputeNewParamsVisaConsumerCanceledMerchandiseMerchantResolutionAttemptedProhibitedByLocalLaw:
		return true
	}
	return false
}

// Return outcome.
type CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnOutcome string

const (
	CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnOutcomeNotReturned     CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnOutcome = "not_returned"
	CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnOutcomeReturned        CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnOutcome = "returned"
	CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnOutcomeReturnAttempted CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnOutcome = "return_attempted"
)

func (r CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnOutcome) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnOutcomeNotReturned, CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnOutcomeReturned, CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnOutcomeReturnAttempted:
		return true
	}
	return false
}

// Cardholder cancellation.
type CardDisputeNewParamsVisaConsumerCanceledMerchandiseCardholderCancellation struct {
	// Canceled at.
	CanceledAt param.Field[time.Time] `json:"canceled_at,required" format:"date"`
	// Canceled prior to ship date.
	CanceledPriorToShipDate param.Field[CardDisputeNewParamsVisaConsumerCanceledMerchandiseCardholderCancellationCanceledPriorToShipDate] `json:"canceled_prior_to_ship_date,required"`
	// Cancellation policy provided.
	CancellationPolicyProvided param.Field[CardDisputeNewParamsVisaConsumerCanceledMerchandiseCardholderCancellationCancellationPolicyProvided] `json:"cancellation_policy_provided,required"`
	// Reason.
	Reason param.Field[string] `json:"reason,required"`
}

func (r CardDisputeNewParamsVisaConsumerCanceledMerchandiseCardholderCancellation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Canceled prior to ship date.
type CardDisputeNewParamsVisaConsumerCanceledMerchandiseCardholderCancellationCanceledPriorToShipDate string

const (
	CardDisputeNewParamsVisaConsumerCanceledMerchandiseCardholderCancellationCanceledPriorToShipDateCanceledPriorToShipDate    CardDisputeNewParamsVisaConsumerCanceledMerchandiseCardholderCancellationCanceledPriorToShipDate = "canceled_prior_to_ship_date"
	CardDisputeNewParamsVisaConsumerCanceledMerchandiseCardholderCancellationCanceledPriorToShipDateNotCanceledPriorToShipDate CardDisputeNewParamsVisaConsumerCanceledMerchandiseCardholderCancellationCanceledPriorToShipDate = "not_canceled_prior_to_ship_date"
)

func (r CardDisputeNewParamsVisaConsumerCanceledMerchandiseCardholderCancellationCanceledPriorToShipDate) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsVisaConsumerCanceledMerchandiseCardholderCancellationCanceledPriorToShipDateCanceledPriorToShipDate, CardDisputeNewParamsVisaConsumerCanceledMerchandiseCardholderCancellationCanceledPriorToShipDateNotCanceledPriorToShipDate:
		return true
	}
	return false
}

// Cancellation policy provided.
type CardDisputeNewParamsVisaConsumerCanceledMerchandiseCardholderCancellationCancellationPolicyProvided string

const (
	CardDisputeNewParamsVisaConsumerCanceledMerchandiseCardholderCancellationCancellationPolicyProvidedNotProvided CardDisputeNewParamsVisaConsumerCanceledMerchandiseCardholderCancellationCancellationPolicyProvided = "not_provided"
	CardDisputeNewParamsVisaConsumerCanceledMerchandiseCardholderCancellationCancellationPolicyProvidedProvided    CardDisputeNewParamsVisaConsumerCanceledMerchandiseCardholderCancellationCancellationPolicyProvided = "provided"
)

func (r CardDisputeNewParamsVisaConsumerCanceledMerchandiseCardholderCancellationCancellationPolicyProvided) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsVisaConsumerCanceledMerchandiseCardholderCancellationCancellationPolicyProvidedNotProvided, CardDisputeNewParamsVisaConsumerCanceledMerchandiseCardholderCancellationCancellationPolicyProvidedProvided:
		return true
	}
	return false
}

// Not returned. Required if and only if `return_outcome` is `not_returned`.
type CardDisputeNewParamsVisaConsumerCanceledMerchandiseNotReturned struct {
}

func (r CardDisputeNewParamsVisaConsumerCanceledMerchandiseNotReturned) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Return attempted. Required if and only if `return_outcome` is
// `return_attempted`.
type CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnAttempted struct {
	// Attempt explanation.
	AttemptExplanation param.Field[string] `json:"attempt_explanation,required"`
	// Attempt reason.
	AttemptReason param.Field[CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnAttemptedAttemptReason] `json:"attempt_reason,required"`
	// Attempted at.
	AttemptedAt param.Field[time.Time] `json:"attempted_at,required" format:"date"`
	// Merchandise disposition.
	MerchandiseDisposition param.Field[string] `json:"merchandise_disposition,required"`
}

func (r CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnAttempted) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Attempt reason.
type CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnAttemptedAttemptReason string

const (
	CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnAttemptedAttemptReasonMerchantNotResponding         CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnAttemptedAttemptReason = "merchant_not_responding"
	CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnAttemptedAttemptReasonNoReturnAuthorizationProvided CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnAttemptedAttemptReason = "no_return_authorization_provided"
	CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnAttemptedAttemptReasonNoReturnInstructions          CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnAttemptedAttemptReason = "no_return_instructions"
	CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnAttemptedAttemptReasonRequestedNotToReturn          CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnAttemptedAttemptReason = "requested_not_to_return"
	CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnAttemptedAttemptReasonReturnNotAccepted             CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnAttemptedAttemptReason = "return_not_accepted"
)

func (r CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnAttemptedAttemptReason) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnAttemptedAttemptReasonMerchantNotResponding, CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnAttemptedAttemptReasonNoReturnAuthorizationProvided, CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnAttemptedAttemptReasonNoReturnInstructions, CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnAttemptedAttemptReasonRequestedNotToReturn, CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnAttemptedAttemptReasonReturnNotAccepted:
		return true
	}
	return false
}

// Returned. Required if and only if `return_outcome` is `returned`.
type CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturned struct {
	// Return method.
	ReturnMethod param.Field[CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnedReturnMethod] `json:"return_method,required"`
	// Returned at.
	ReturnedAt param.Field[time.Time] `json:"returned_at,required" format:"date"`
	// Merchant received return at.
	MerchantReceivedReturnAt param.Field[time.Time] `json:"merchant_received_return_at" format:"date"`
	// Other explanation. Required if and only if the return method is `other`.
	OtherExplanation param.Field[string] `json:"other_explanation"`
	// Tracking number.
	TrackingNumber param.Field[string] `json:"tracking_number"`
}

func (r CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturned) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Return method.
type CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnedReturnMethod string

const (
	CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnedReturnMethodDhl           CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnedReturnMethod = "dhl"
	CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnedReturnMethodFaceToFace    CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnedReturnMethod = "face_to_face"
	CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnedReturnMethodFedex         CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnedReturnMethod = "fedex"
	CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnedReturnMethodOther         CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnedReturnMethod = "other"
	CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnedReturnMethodPostalService CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnedReturnMethod = "postal_service"
	CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnedReturnMethodUps           CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnedReturnMethod = "ups"
)

func (r CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnedReturnMethod) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnedReturnMethodDhl, CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnedReturnMethodFaceToFace, CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnedReturnMethodFedex, CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnedReturnMethodOther, CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnedReturnMethodPostalService, CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnedReturnMethodUps:
		return true
	}
	return false
}

// Canceled recurring transaction. Required if and only if `category` is
// `consumer_canceled_recurring_transaction`.
type CardDisputeNewParamsVisaConsumerCanceledRecurringTransaction struct {
	// Cancellation target.
	CancellationTarget param.Field[CardDisputeNewParamsVisaConsumerCanceledRecurringTransactionCancellationTarget] `json:"cancellation_target,required"`
	// Merchant contact methods.
	MerchantContactMethods param.Field[CardDisputeNewParamsVisaConsumerCanceledRecurringTransactionMerchantContactMethods] `json:"merchant_contact_methods,required"`
	// Transaction or account canceled at.
	TransactionOrAccountCanceledAt param.Field[time.Time] `json:"transaction_or_account_canceled_at,required" format:"date"`
	// Other form of payment explanation.
	OtherFormOfPaymentExplanation param.Field[string] `json:"other_form_of_payment_explanation"`
}

func (r CardDisputeNewParamsVisaConsumerCanceledRecurringTransaction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Cancellation target.
type CardDisputeNewParamsVisaConsumerCanceledRecurringTransactionCancellationTarget string

const (
	CardDisputeNewParamsVisaConsumerCanceledRecurringTransactionCancellationTargetAccount     CardDisputeNewParamsVisaConsumerCanceledRecurringTransactionCancellationTarget = "account"
	CardDisputeNewParamsVisaConsumerCanceledRecurringTransactionCancellationTargetTransaction CardDisputeNewParamsVisaConsumerCanceledRecurringTransactionCancellationTarget = "transaction"
)

func (r CardDisputeNewParamsVisaConsumerCanceledRecurringTransactionCancellationTarget) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsVisaConsumerCanceledRecurringTransactionCancellationTargetAccount, CardDisputeNewParamsVisaConsumerCanceledRecurringTransactionCancellationTargetTransaction:
		return true
	}
	return false
}

// Merchant contact methods.
type CardDisputeNewParamsVisaConsumerCanceledRecurringTransactionMerchantContactMethods struct {
	// Application name.
	ApplicationName param.Field[string] `json:"application_name"`
	// Call center phone number.
	CallCenterPhoneNumber param.Field[string] `json:"call_center_phone_number"`
	// Email address.
	EmailAddress param.Field[string] `json:"email_address"`
	// In person address.
	InPersonAddress param.Field[string] `json:"in_person_address"`
	// Mailing address.
	MailingAddress param.Field[string] `json:"mailing_address"`
	// Text phone number.
	TextPhoneNumber param.Field[string] `json:"text_phone_number"`
}

func (r CardDisputeNewParamsVisaConsumerCanceledRecurringTransactionMerchantContactMethods) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Canceled services. Required if and only if `category` is
// `consumer_canceled_services`.
type CardDisputeNewParamsVisaConsumerCanceledServices struct {
	// Cardholder cancellation.
	CardholderCancellation param.Field[CardDisputeNewParamsVisaConsumerCanceledServicesCardholderCancellation] `json:"cardholder_cancellation,required"`
	// Contracted at.
	ContractedAt param.Field[time.Time] `json:"contracted_at,required" format:"date"`
	// Merchant resolution attempted.
	MerchantResolutionAttempted param.Field[CardDisputeNewParamsVisaConsumerCanceledServicesMerchantResolutionAttempted] `json:"merchant_resolution_attempted,required"`
	// Purchase explanation.
	PurchaseExplanation param.Field[string] `json:"purchase_explanation,required"`
	// Service type.
	ServiceType param.Field[CardDisputeNewParamsVisaConsumerCanceledServicesServiceType] `json:"service_type,required"`
	// Guaranteed reservation explanation. Required if and only if `service_type` is
	// `guaranteed_reservation`.
	GuaranteedReservation param.Field[CardDisputeNewParamsVisaConsumerCanceledServicesGuaranteedReservation] `json:"guaranteed_reservation"`
	// Other service type explanation. Required if and only if `service_type` is
	// `other`.
	Other param.Field[CardDisputeNewParamsVisaConsumerCanceledServicesOther] `json:"other"`
	// Timeshare explanation. Required if and only if `service_type` is `timeshare`.
	Timeshare param.Field[CardDisputeNewParamsVisaConsumerCanceledServicesTimeshare] `json:"timeshare"`
}

func (r CardDisputeNewParamsVisaConsumerCanceledServices) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Cardholder cancellation.
type CardDisputeNewParamsVisaConsumerCanceledServicesCardholderCancellation struct {
	// Canceled at.
	CanceledAt param.Field[time.Time] `json:"canceled_at,required" format:"date"`
	// Cancellation policy provided.
	CancellationPolicyProvided param.Field[CardDisputeNewParamsVisaConsumerCanceledServicesCardholderCancellationCancellationPolicyProvided] `json:"cancellation_policy_provided,required"`
	// Reason.
	Reason param.Field[string] `json:"reason,required"`
}

func (r CardDisputeNewParamsVisaConsumerCanceledServicesCardholderCancellation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Cancellation policy provided.
type CardDisputeNewParamsVisaConsumerCanceledServicesCardholderCancellationCancellationPolicyProvided string

const (
	CardDisputeNewParamsVisaConsumerCanceledServicesCardholderCancellationCancellationPolicyProvidedNotProvided CardDisputeNewParamsVisaConsumerCanceledServicesCardholderCancellationCancellationPolicyProvided = "not_provided"
	CardDisputeNewParamsVisaConsumerCanceledServicesCardholderCancellationCancellationPolicyProvidedProvided    CardDisputeNewParamsVisaConsumerCanceledServicesCardholderCancellationCancellationPolicyProvided = "provided"
)

func (r CardDisputeNewParamsVisaConsumerCanceledServicesCardholderCancellationCancellationPolicyProvided) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsVisaConsumerCanceledServicesCardholderCancellationCancellationPolicyProvidedNotProvided, CardDisputeNewParamsVisaConsumerCanceledServicesCardholderCancellationCancellationPolicyProvidedProvided:
		return true
	}
	return false
}

// Merchant resolution attempted.
type CardDisputeNewParamsVisaConsumerCanceledServicesMerchantResolutionAttempted string

const (
	CardDisputeNewParamsVisaConsumerCanceledServicesMerchantResolutionAttemptedAttempted            CardDisputeNewParamsVisaConsumerCanceledServicesMerchantResolutionAttempted = "attempted"
	CardDisputeNewParamsVisaConsumerCanceledServicesMerchantResolutionAttemptedProhibitedByLocalLaw CardDisputeNewParamsVisaConsumerCanceledServicesMerchantResolutionAttempted = "prohibited_by_local_law"
)

func (r CardDisputeNewParamsVisaConsumerCanceledServicesMerchantResolutionAttempted) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsVisaConsumerCanceledServicesMerchantResolutionAttemptedAttempted, CardDisputeNewParamsVisaConsumerCanceledServicesMerchantResolutionAttemptedProhibitedByLocalLaw:
		return true
	}
	return false
}

// Service type.
type CardDisputeNewParamsVisaConsumerCanceledServicesServiceType string

const (
	CardDisputeNewParamsVisaConsumerCanceledServicesServiceTypeGuaranteedReservation CardDisputeNewParamsVisaConsumerCanceledServicesServiceType = "guaranteed_reservation"
	CardDisputeNewParamsVisaConsumerCanceledServicesServiceTypeOther                 CardDisputeNewParamsVisaConsumerCanceledServicesServiceType = "other"
	CardDisputeNewParamsVisaConsumerCanceledServicesServiceTypeTimeshare             CardDisputeNewParamsVisaConsumerCanceledServicesServiceType = "timeshare"
)

func (r CardDisputeNewParamsVisaConsumerCanceledServicesServiceType) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsVisaConsumerCanceledServicesServiceTypeGuaranteedReservation, CardDisputeNewParamsVisaConsumerCanceledServicesServiceTypeOther, CardDisputeNewParamsVisaConsumerCanceledServicesServiceTypeTimeshare:
		return true
	}
	return false
}

// Guaranteed reservation explanation. Required if and only if `service_type` is
// `guaranteed_reservation`.
type CardDisputeNewParamsVisaConsumerCanceledServicesGuaranteedReservation struct {
	// Explanation.
	Explanation param.Field[CardDisputeNewParamsVisaConsumerCanceledServicesGuaranteedReservationExplanation] `json:"explanation,required"`
}

func (r CardDisputeNewParamsVisaConsumerCanceledServicesGuaranteedReservation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Explanation.
type CardDisputeNewParamsVisaConsumerCanceledServicesGuaranteedReservationExplanation string

const (
	CardDisputeNewParamsVisaConsumerCanceledServicesGuaranteedReservationExplanationCardholderCanceledPriorToService                         CardDisputeNewParamsVisaConsumerCanceledServicesGuaranteedReservationExplanation = "cardholder_canceled_prior_to_service"
	CardDisputeNewParamsVisaConsumerCanceledServicesGuaranteedReservationExplanationCardholderCancellationAttemptWithin24HoursOfConfirmation CardDisputeNewParamsVisaConsumerCanceledServicesGuaranteedReservationExplanation = "cardholder_cancellation_attempt_within_24_hours_of_confirmation"
	CardDisputeNewParamsVisaConsumerCanceledServicesGuaranteedReservationExplanationMerchantBilledNoShow                                     CardDisputeNewParamsVisaConsumerCanceledServicesGuaranteedReservationExplanation = "merchant_billed_no_show"
)

func (r CardDisputeNewParamsVisaConsumerCanceledServicesGuaranteedReservationExplanation) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsVisaConsumerCanceledServicesGuaranteedReservationExplanationCardholderCanceledPriorToService, CardDisputeNewParamsVisaConsumerCanceledServicesGuaranteedReservationExplanationCardholderCancellationAttemptWithin24HoursOfConfirmation, CardDisputeNewParamsVisaConsumerCanceledServicesGuaranteedReservationExplanationMerchantBilledNoShow:
		return true
	}
	return false
}

// Other service type explanation. Required if and only if `service_type` is
// `other`.
type CardDisputeNewParamsVisaConsumerCanceledServicesOther struct {
}

func (r CardDisputeNewParamsVisaConsumerCanceledServicesOther) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Timeshare explanation. Required if and only if `service_type` is `timeshare`.
type CardDisputeNewParamsVisaConsumerCanceledServicesTimeshare struct {
}

func (r CardDisputeNewParamsVisaConsumerCanceledServicesTimeshare) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Counterfeit merchandise. Required if and only if `category` is
// `consumer_counterfeit_merchandise`.
type CardDisputeNewParamsVisaConsumerCounterfeitMerchandise struct {
	// Counterfeit explanation.
	CounterfeitExplanation param.Field[string] `json:"counterfeit_explanation,required"`
	// Disposition explanation.
	DispositionExplanation param.Field[string] `json:"disposition_explanation,required"`
	// Order explanation.
	OrderExplanation param.Field[string] `json:"order_explanation,required"`
	// Received at.
	ReceivedAt param.Field[time.Time] `json:"received_at,required" format:"date"`
}

func (r CardDisputeNewParamsVisaConsumerCounterfeitMerchandise) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Credit not processed. Required if and only if `category` is
// `consumer_credit_not_processed`.
type CardDisputeNewParamsVisaConsumerCreditNotProcessed struct {
	// Canceled or returned at.
	CanceledOrReturnedAt param.Field[time.Time] `json:"canceled_or_returned_at" format:"date"`
	// Credit expected at.
	CreditExpectedAt param.Field[time.Time] `json:"credit_expected_at" format:"date"`
}

func (r CardDisputeNewParamsVisaConsumerCreditNotProcessed) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Damaged or defective merchandise. Required if and only if `category` is
// `consumer_damaged_or_defective_merchandise`.
type CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandise struct {
	// Merchant resolution attempted.
	MerchantResolutionAttempted param.Field[CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseMerchantResolutionAttempted] `json:"merchant_resolution_attempted,required"`
	// Order and issue explanation.
	OrderAndIssueExplanation param.Field[string] `json:"order_and_issue_explanation,required"`
	// Received at.
	ReceivedAt param.Field[time.Time] `json:"received_at,required" format:"date"`
	// Return outcome.
	ReturnOutcome param.Field[CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnOutcome] `json:"return_outcome,required"`
	// Not returned. Required if and only if `return_outcome` is `not_returned`.
	NotReturned param.Field[CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseNotReturned] `json:"not_returned"`
	// Return attempted. Required if and only if `return_outcome` is
	// `return_attempted`.
	ReturnAttempted param.Field[CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnAttempted] `json:"return_attempted"`
	// Returned. Required if and only if `return_outcome` is `returned`.
	Returned param.Field[CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturned] `json:"returned"`
}

func (r CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandise) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Merchant resolution attempted.
type CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseMerchantResolutionAttempted string

const (
	CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseMerchantResolutionAttemptedAttempted            CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseMerchantResolutionAttempted = "attempted"
	CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseMerchantResolutionAttemptedProhibitedByLocalLaw CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseMerchantResolutionAttempted = "prohibited_by_local_law"
)

func (r CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseMerchantResolutionAttempted) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseMerchantResolutionAttemptedAttempted, CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseMerchantResolutionAttemptedProhibitedByLocalLaw:
		return true
	}
	return false
}

// Return outcome.
type CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnOutcome string

const (
	CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnOutcomeNotReturned     CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnOutcome = "not_returned"
	CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnOutcomeReturned        CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnOutcome = "returned"
	CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnOutcomeReturnAttempted CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnOutcome = "return_attempted"
)

func (r CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnOutcome) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnOutcomeNotReturned, CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnOutcomeReturned, CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnOutcomeReturnAttempted:
		return true
	}
	return false
}

// Not returned. Required if and only if `return_outcome` is `not_returned`.
type CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseNotReturned struct {
}

func (r CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseNotReturned) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Return attempted. Required if and only if `return_outcome` is
// `return_attempted`.
type CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnAttempted struct {
	// Attempt explanation.
	AttemptExplanation param.Field[string] `json:"attempt_explanation,required"`
	// Attempt reason.
	AttemptReason param.Field[CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReason] `json:"attempt_reason,required"`
	// Attempted at.
	AttemptedAt param.Field[time.Time] `json:"attempted_at,required" format:"date"`
	// Merchandise disposition.
	MerchandiseDisposition param.Field[string] `json:"merchandise_disposition,required"`
}

func (r CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnAttempted) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Attempt reason.
type CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReason string

const (
	CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReasonMerchantNotResponding         CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReason = "merchant_not_responding"
	CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReasonNoReturnAuthorizationProvided CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReason = "no_return_authorization_provided"
	CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReasonNoReturnInstructions          CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReason = "no_return_instructions"
	CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReasonRequestedNotToReturn          CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReason = "requested_not_to_return"
	CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReasonReturnNotAccepted             CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReason = "return_not_accepted"
)

func (r CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReason) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReasonMerchantNotResponding, CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReasonNoReturnAuthorizationProvided, CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReasonNoReturnInstructions, CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReasonRequestedNotToReturn, CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReasonReturnNotAccepted:
		return true
	}
	return false
}

// Returned. Required if and only if `return_outcome` is `returned`.
type CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturned struct {
	// Return method.
	ReturnMethod param.Field[CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethod] `json:"return_method,required"`
	// Returned at.
	ReturnedAt param.Field[time.Time] `json:"returned_at,required" format:"date"`
	// Merchant received return at.
	MerchantReceivedReturnAt param.Field[time.Time] `json:"merchant_received_return_at" format:"date"`
	// Other explanation. Required if and only if the return method is `other`.
	OtherExplanation param.Field[string] `json:"other_explanation"`
	// Tracking number.
	TrackingNumber param.Field[string] `json:"tracking_number"`
}

func (r CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturned) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Return method.
type CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethod string

const (
	CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethodDhl           CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethod = "dhl"
	CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethodFaceToFace    CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethod = "face_to_face"
	CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethodFedex         CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethod = "fedex"
	CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethodOther         CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethod = "other"
	CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethodPostalService CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethod = "postal_service"
	CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethodUps           CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethod = "ups"
)

func (r CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethod) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethodDhl, CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethodFaceToFace, CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethodFedex, CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethodOther, CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethodPostalService, CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethodUps:
		return true
	}
	return false
}

// Merchandise misrepresentation. Required if and only if `category` is
// `consumer_merchandise_misrepresentation`.
type CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentation struct {
	// Merchant resolution attempted.
	MerchantResolutionAttempted param.Field[CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationMerchantResolutionAttempted] `json:"merchant_resolution_attempted,required"`
	// Misrepresentation explanation.
	MisrepresentationExplanation param.Field[string] `json:"misrepresentation_explanation,required"`
	// Purchase explanation.
	PurchaseExplanation param.Field[string] `json:"purchase_explanation,required"`
	// Received at.
	ReceivedAt param.Field[time.Time] `json:"received_at,required" format:"date"`
	// Return outcome.
	ReturnOutcome param.Field[CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnOutcome] `json:"return_outcome,required"`
	// Not returned. Required if and only if `return_outcome` is `not_returned`.
	NotReturned param.Field[CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationNotReturned] `json:"not_returned"`
	// Return attempted. Required if and only if `return_outcome` is
	// `return_attempted`.
	ReturnAttempted param.Field[CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnAttempted] `json:"return_attempted"`
	// Returned. Required if and only if `return_outcome` is `returned`.
	Returned param.Field[CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturned] `json:"returned"`
}

func (r CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Merchant resolution attempted.
type CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationMerchantResolutionAttempted string

const (
	CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationMerchantResolutionAttemptedAttempted            CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationMerchantResolutionAttempted = "attempted"
	CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationMerchantResolutionAttemptedProhibitedByLocalLaw CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationMerchantResolutionAttempted = "prohibited_by_local_law"
)

func (r CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationMerchantResolutionAttempted) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationMerchantResolutionAttemptedAttempted, CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationMerchantResolutionAttemptedProhibitedByLocalLaw:
		return true
	}
	return false
}

// Return outcome.
type CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnOutcome string

const (
	CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnOutcomeNotReturned     CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnOutcome = "not_returned"
	CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnOutcomeReturned        CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnOutcome = "returned"
	CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnOutcomeReturnAttempted CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnOutcome = "return_attempted"
)

func (r CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnOutcome) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnOutcomeNotReturned, CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnOutcomeReturned, CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnOutcomeReturnAttempted:
		return true
	}
	return false
}

// Not returned. Required if and only if `return_outcome` is `not_returned`.
type CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationNotReturned struct {
}

func (r CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationNotReturned) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Return attempted. Required if and only if `return_outcome` is
// `return_attempted`.
type CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnAttempted struct {
	// Attempt explanation.
	AttemptExplanation param.Field[string] `json:"attempt_explanation,required"`
	// Attempt reason.
	AttemptReason param.Field[CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReason] `json:"attempt_reason,required"`
	// Attempted at.
	AttemptedAt param.Field[time.Time] `json:"attempted_at,required" format:"date"`
	// Merchandise disposition.
	MerchandiseDisposition param.Field[string] `json:"merchandise_disposition,required"`
}

func (r CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnAttempted) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Attempt reason.
type CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReason string

const (
	CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReasonMerchantNotResponding         CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReason = "merchant_not_responding"
	CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReasonNoReturnAuthorizationProvided CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReason = "no_return_authorization_provided"
	CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReasonNoReturnInstructions          CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReason = "no_return_instructions"
	CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReasonRequestedNotToReturn          CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReason = "requested_not_to_return"
	CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReasonReturnNotAccepted             CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReason = "return_not_accepted"
)

func (r CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReason) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReasonMerchantNotResponding, CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReasonNoReturnAuthorizationProvided, CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReasonNoReturnInstructions, CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReasonRequestedNotToReturn, CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReasonReturnNotAccepted:
		return true
	}
	return false
}

// Returned. Required if and only if `return_outcome` is `returned`.
type CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturned struct {
	// Return method.
	ReturnMethod param.Field[CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnedReturnMethod] `json:"return_method,required"`
	// Returned at.
	ReturnedAt param.Field[time.Time] `json:"returned_at,required" format:"date"`
	// Merchant received return at.
	MerchantReceivedReturnAt param.Field[time.Time] `json:"merchant_received_return_at" format:"date"`
	// Other explanation. Required if and only if the return method is `other`.
	OtherExplanation param.Field[string] `json:"other_explanation"`
	// Tracking number.
	TrackingNumber param.Field[string] `json:"tracking_number"`
}

func (r CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturned) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Return method.
type CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnedReturnMethod string

const (
	CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnedReturnMethodDhl           CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnedReturnMethod = "dhl"
	CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnedReturnMethodFaceToFace    CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnedReturnMethod = "face_to_face"
	CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnedReturnMethodFedex         CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnedReturnMethod = "fedex"
	CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnedReturnMethodOther         CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnedReturnMethod = "other"
	CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnedReturnMethodPostalService CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnedReturnMethod = "postal_service"
	CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnedReturnMethodUps           CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnedReturnMethod = "ups"
)

func (r CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnedReturnMethod) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnedReturnMethodDhl, CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnedReturnMethodFaceToFace, CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnedReturnMethodFedex, CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnedReturnMethodOther, CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnedReturnMethodPostalService, CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnedReturnMethodUps:
		return true
	}
	return false
}

// Merchandise not as described. Required if and only if `category` is
// `consumer_merchandise_not_as_described`.
type CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribed struct {
	// Merchant resolution attempted.
	MerchantResolutionAttempted param.Field[CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedMerchantResolutionAttempted] `json:"merchant_resolution_attempted,required"`
	// Received at.
	ReceivedAt param.Field[time.Time] `json:"received_at,required" format:"date"`
	// Return outcome.
	ReturnOutcome param.Field[CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnOutcome] `json:"return_outcome,required"`
	// Return attempted. Required if and only if `return_outcome` is
	// `return_attempted`.
	ReturnAttempted param.Field[CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnAttempted] `json:"return_attempted"`
	// Returned. Required if and only if `return_outcome` is `returned`.
	Returned param.Field[CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturned] `json:"returned"`
}

func (r CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribed) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Merchant resolution attempted.
type CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedMerchantResolutionAttempted string

const (
	CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedMerchantResolutionAttemptedAttempted            CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedMerchantResolutionAttempted = "attempted"
	CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedMerchantResolutionAttemptedProhibitedByLocalLaw CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedMerchantResolutionAttempted = "prohibited_by_local_law"
)

func (r CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedMerchantResolutionAttempted) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedMerchantResolutionAttemptedAttempted, CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedMerchantResolutionAttemptedProhibitedByLocalLaw:
		return true
	}
	return false
}

// Return outcome.
type CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnOutcome string

const (
	CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnOutcomeReturned        CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnOutcome = "returned"
	CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnOutcomeReturnAttempted CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnOutcome = "return_attempted"
)

func (r CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnOutcome) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnOutcomeReturned, CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnOutcomeReturnAttempted:
		return true
	}
	return false
}

// Return attempted. Required if and only if `return_outcome` is
// `return_attempted`.
type CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnAttempted struct {
	// Attempt explanation.
	AttemptExplanation param.Field[string] `json:"attempt_explanation,required"`
	// Attempt reason.
	AttemptReason param.Field[CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReason] `json:"attempt_reason,required"`
	// Attempted at.
	AttemptedAt param.Field[time.Time] `json:"attempted_at,required" format:"date"`
	// Merchandise disposition.
	MerchandiseDisposition param.Field[string] `json:"merchandise_disposition,required"`
}

func (r CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnAttempted) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Attempt reason.
type CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReason string

const (
	CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReasonMerchantNotResponding         CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReason = "merchant_not_responding"
	CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReasonNoReturnAuthorizationProvided CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReason = "no_return_authorization_provided"
	CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReasonNoReturnInstructions          CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReason = "no_return_instructions"
	CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReasonRequestedNotToReturn          CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReason = "requested_not_to_return"
	CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReasonReturnNotAccepted             CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReason = "return_not_accepted"
)

func (r CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReason) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReasonMerchantNotResponding, CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReasonNoReturnAuthorizationProvided, CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReasonNoReturnInstructions, CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReasonRequestedNotToReturn, CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReasonReturnNotAccepted:
		return true
	}
	return false
}

// Returned. Required if and only if `return_outcome` is `returned`.
type CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturned struct {
	// Return method.
	ReturnMethod param.Field[CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnedReturnMethod] `json:"return_method,required"`
	// Returned at.
	ReturnedAt param.Field[time.Time] `json:"returned_at,required" format:"date"`
	// Merchant received return at.
	MerchantReceivedReturnAt param.Field[time.Time] `json:"merchant_received_return_at" format:"date"`
	// Other explanation. Required if and only if the return method is `other`.
	OtherExplanation param.Field[string] `json:"other_explanation"`
	// Tracking number.
	TrackingNumber param.Field[string] `json:"tracking_number"`
}

func (r CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturned) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Return method.
type CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnedReturnMethod string

const (
	CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnedReturnMethodDhl           CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnedReturnMethod = "dhl"
	CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnedReturnMethodFaceToFace    CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnedReturnMethod = "face_to_face"
	CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnedReturnMethodFedex         CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnedReturnMethod = "fedex"
	CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnedReturnMethodOther         CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnedReturnMethod = "other"
	CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnedReturnMethodPostalService CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnedReturnMethod = "postal_service"
	CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnedReturnMethodUps           CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnedReturnMethod = "ups"
)

func (r CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnedReturnMethod) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnedReturnMethodDhl, CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnedReturnMethodFaceToFace, CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnedReturnMethodFedex, CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnedReturnMethodOther, CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnedReturnMethodPostalService, CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnedReturnMethodUps:
		return true
	}
	return false
}

// Merchandise not received. Required if and only if `category` is
// `consumer_merchandise_not_received`.
type CardDisputeNewParamsVisaConsumerMerchandiseNotReceived struct {
	// Cancellation outcome.
	CancellationOutcome param.Field[CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedCancellationOutcome] `json:"cancellation_outcome,required"`
	// Delivery issue.
	DeliveryIssue param.Field[CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedDeliveryIssue] `json:"delivery_issue,required"`
	// Last expected receipt at.
	LastExpectedReceiptAt param.Field[time.Time] `json:"last_expected_receipt_at,required" format:"date"`
	// Merchant resolution attempted.
	MerchantResolutionAttempted param.Field[CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedMerchantResolutionAttempted] `json:"merchant_resolution_attempted,required"`
	// Purchase information and explanation.
	PurchaseInfoAndExplanation param.Field[string] `json:"purchase_info_and_explanation,required"`
	// Cardholder cancellation prior to expected receipt. Required if and only if
	// `cancellation_outcome` is `cardholder_cancellation_prior_to_expected_receipt`.
	CardholderCancellationPriorToExpectedReceipt param.Field[CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedCardholderCancellationPriorToExpectedReceipt] `json:"cardholder_cancellation_prior_to_expected_receipt"`
	// Delayed. Required if and only if `delivery_issue` is `delayed`.
	Delayed param.Field[CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedDelayed] `json:"delayed"`
	// Delivered to wrong location. Required if and only if `delivery_issue` is
	// `delivered_to_wrong_location`.
	DeliveredToWrongLocation param.Field[CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedDeliveredToWrongLocation] `json:"delivered_to_wrong_location"`
	// Merchant cancellation. Required if and only if `cancellation_outcome` is
	// `merchant_cancellation`.
	MerchantCancellation param.Field[CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedMerchantCancellation] `json:"merchant_cancellation"`
	// No cancellation. Required if and only if `cancellation_outcome` is
	// `no_cancellation`.
	NoCancellation param.Field[CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedNoCancellation] `json:"no_cancellation"`
}

func (r CardDisputeNewParamsVisaConsumerMerchandiseNotReceived) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Cancellation outcome.
type CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedCancellationOutcome string

const (
	CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedCancellationOutcomeCardholderCancellationPriorToExpectedReceipt CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedCancellationOutcome = "cardholder_cancellation_prior_to_expected_receipt"
	CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedCancellationOutcomeMerchantCancellation                         CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedCancellationOutcome = "merchant_cancellation"
	CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedCancellationOutcomeNoCancellation                               CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedCancellationOutcome = "no_cancellation"
)

func (r CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedCancellationOutcome) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedCancellationOutcomeCardholderCancellationPriorToExpectedReceipt, CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedCancellationOutcomeMerchantCancellation, CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedCancellationOutcomeNoCancellation:
		return true
	}
	return false
}

// Delivery issue.
type CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedDeliveryIssue string

const (
	CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedDeliveryIssueDelayed                  CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedDeliveryIssue = "delayed"
	CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedDeliveryIssueDeliveredToWrongLocation CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedDeliveryIssue = "delivered_to_wrong_location"
)

func (r CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedDeliveryIssue) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedDeliveryIssueDelayed, CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedDeliveryIssueDeliveredToWrongLocation:
		return true
	}
	return false
}

// Merchant resolution attempted.
type CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedMerchantResolutionAttempted string

const (
	CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedMerchantResolutionAttemptedAttempted            CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedMerchantResolutionAttempted = "attempted"
	CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedMerchantResolutionAttemptedProhibitedByLocalLaw CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedMerchantResolutionAttempted = "prohibited_by_local_law"
)

func (r CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedMerchantResolutionAttempted) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedMerchantResolutionAttemptedAttempted, CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedMerchantResolutionAttemptedProhibitedByLocalLaw:
		return true
	}
	return false
}

// Cardholder cancellation prior to expected receipt. Required if and only if
// `cancellation_outcome` is `cardholder_cancellation_prior_to_expected_receipt`.
type CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedCardholderCancellationPriorToExpectedReceipt struct {
	// Canceled at.
	CanceledAt param.Field[time.Time] `json:"canceled_at,required" format:"date"`
	// Reason.
	Reason param.Field[string] `json:"reason"`
}

func (r CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedCardholderCancellationPriorToExpectedReceipt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Delayed. Required if and only if `delivery_issue` is `delayed`.
type CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedDelayed struct {
	// Explanation.
	Explanation param.Field[string] `json:"explanation,required"`
	// Return outcome.
	ReturnOutcome param.Field[CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedDelayedReturnOutcome] `json:"return_outcome,required"`
	// Not returned. Required if and only if `return_outcome` is `not_returned`.
	NotReturned param.Field[CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedDelayedNotReturned] `json:"not_returned"`
	// Return attempted. Required if and only if `return_outcome` is
	// `return_attempted`.
	ReturnAttempted param.Field[CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedDelayedReturnAttempted] `json:"return_attempted"`
	// Returned. Required if and only if `return_outcome` is `returned`.
	Returned param.Field[CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedDelayedReturned] `json:"returned"`
}

func (r CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedDelayed) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Return outcome.
type CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedDelayedReturnOutcome string

const (
	CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedDelayedReturnOutcomeNotReturned     CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedDelayedReturnOutcome = "not_returned"
	CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedDelayedReturnOutcomeReturned        CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedDelayedReturnOutcome = "returned"
	CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedDelayedReturnOutcomeReturnAttempted CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedDelayedReturnOutcome = "return_attempted"
)

func (r CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedDelayedReturnOutcome) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedDelayedReturnOutcomeNotReturned, CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedDelayedReturnOutcomeReturned, CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedDelayedReturnOutcomeReturnAttempted:
		return true
	}
	return false
}

// Not returned. Required if and only if `return_outcome` is `not_returned`.
type CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedDelayedNotReturned struct {
}

func (r CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedDelayedNotReturned) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Return attempted. Required if and only if `return_outcome` is
// `return_attempted`.
type CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedDelayedReturnAttempted struct {
	// Attempted at.
	AttemptedAt param.Field[time.Time] `json:"attempted_at,required" format:"date"`
}

func (r CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedDelayedReturnAttempted) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Returned. Required if and only if `return_outcome` is `returned`.
type CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedDelayedReturned struct {
	// Merchant received return at.
	MerchantReceivedReturnAt param.Field[time.Time] `json:"merchant_received_return_at,required" format:"date"`
	// Returned at.
	ReturnedAt param.Field[time.Time] `json:"returned_at,required" format:"date"`
}

func (r CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedDelayedReturned) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Delivered to wrong location. Required if and only if `delivery_issue` is
// `delivered_to_wrong_location`.
type CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedDeliveredToWrongLocation struct {
	// Agreed location.
	AgreedLocation param.Field[string] `json:"agreed_location,required"`
}

func (r CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedDeliveredToWrongLocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Merchant cancellation. Required if and only if `cancellation_outcome` is
// `merchant_cancellation`.
type CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedMerchantCancellation struct {
	// Canceled at.
	CanceledAt param.Field[time.Time] `json:"canceled_at,required" format:"date"`
}

func (r CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedMerchantCancellation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// No cancellation. Required if and only if `cancellation_outcome` is
// `no_cancellation`.
type CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedNoCancellation struct {
}

func (r CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedNoCancellation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Non-receipt of cash. Required if and only if `category` is
// `consumer_non_receipt_of_cash`.
type CardDisputeNewParamsVisaConsumerNonReceiptOfCash struct {
}

func (r CardDisputeNewParamsVisaConsumerNonReceiptOfCash) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Original Credit Transaction (OCT) not accepted. Required if and only if
// `category` is `consumer_original_credit_transaction_not_accepted`.
type CardDisputeNewParamsVisaConsumerOriginalCreditTransactionNotAccepted struct {
	// Explanation.
	Explanation param.Field[string] `json:"explanation,required"`
	// Reason.
	Reason param.Field[CardDisputeNewParamsVisaConsumerOriginalCreditTransactionNotAcceptedReason] `json:"reason,required"`
}

func (r CardDisputeNewParamsVisaConsumerOriginalCreditTransactionNotAccepted) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Reason.
type CardDisputeNewParamsVisaConsumerOriginalCreditTransactionNotAcceptedReason string

const (
	CardDisputeNewParamsVisaConsumerOriginalCreditTransactionNotAcceptedReasonProhibitedByLocalLawsOrRegulation CardDisputeNewParamsVisaConsumerOriginalCreditTransactionNotAcceptedReason = "prohibited_by_local_laws_or_regulation"
	CardDisputeNewParamsVisaConsumerOriginalCreditTransactionNotAcceptedReasonRecipientRefused                  CardDisputeNewParamsVisaConsumerOriginalCreditTransactionNotAcceptedReason = "recipient_refused"
)

func (r CardDisputeNewParamsVisaConsumerOriginalCreditTransactionNotAcceptedReason) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsVisaConsumerOriginalCreditTransactionNotAcceptedReasonProhibitedByLocalLawsOrRegulation, CardDisputeNewParamsVisaConsumerOriginalCreditTransactionNotAcceptedReasonRecipientRefused:
		return true
	}
	return false
}

// Merchandise quality issue. Required if and only if `category` is
// `consumer_quality_merchandise`.
type CardDisputeNewParamsVisaConsumerQualityMerchandise struct {
	// Expected at.
	ExpectedAt param.Field[time.Time] `json:"expected_at,required" format:"date"`
	// Merchant resolution attempted.
	MerchantResolutionAttempted param.Field[CardDisputeNewParamsVisaConsumerQualityMerchandiseMerchantResolutionAttempted] `json:"merchant_resolution_attempted,required"`
	// Purchase information and quality issue.
	PurchaseInfoAndQualityIssue param.Field[string] `json:"purchase_info_and_quality_issue,required"`
	// Received at.
	ReceivedAt param.Field[time.Time] `json:"received_at,required" format:"date"`
	// Return outcome.
	ReturnOutcome param.Field[CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnOutcome] `json:"return_outcome,required"`
	// Not returned. Required if and only if `return_outcome` is `not_returned`.
	NotReturned param.Field[CardDisputeNewParamsVisaConsumerQualityMerchandiseNotReturned] `json:"not_returned"`
	// Ongoing negotiations. Exclude if there is no evidence of ongoing negotiations.
	OngoingNegotiations param.Field[CardDisputeNewParamsVisaConsumerQualityMerchandiseOngoingNegotiations] `json:"ongoing_negotiations"`
	// Return attempted. Required if and only if `return_outcome` is
	// `return_attempted`.
	ReturnAttempted param.Field[CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnAttempted] `json:"return_attempted"`
	// Returned. Required if and only if `return_outcome` is `returned`.
	Returned param.Field[CardDisputeNewParamsVisaConsumerQualityMerchandiseReturned] `json:"returned"`
}

func (r CardDisputeNewParamsVisaConsumerQualityMerchandise) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Merchant resolution attempted.
type CardDisputeNewParamsVisaConsumerQualityMerchandiseMerchantResolutionAttempted string

const (
	CardDisputeNewParamsVisaConsumerQualityMerchandiseMerchantResolutionAttemptedAttempted            CardDisputeNewParamsVisaConsumerQualityMerchandiseMerchantResolutionAttempted = "attempted"
	CardDisputeNewParamsVisaConsumerQualityMerchandiseMerchantResolutionAttemptedProhibitedByLocalLaw CardDisputeNewParamsVisaConsumerQualityMerchandiseMerchantResolutionAttempted = "prohibited_by_local_law"
)

func (r CardDisputeNewParamsVisaConsumerQualityMerchandiseMerchantResolutionAttempted) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsVisaConsumerQualityMerchandiseMerchantResolutionAttemptedAttempted, CardDisputeNewParamsVisaConsumerQualityMerchandiseMerchantResolutionAttemptedProhibitedByLocalLaw:
		return true
	}
	return false
}

// Return outcome.
type CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnOutcome string

const (
	CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnOutcomeNotReturned     CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnOutcome = "not_returned"
	CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnOutcomeReturned        CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnOutcome = "returned"
	CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnOutcomeReturnAttempted CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnOutcome = "return_attempted"
)

func (r CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnOutcome) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnOutcomeNotReturned, CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnOutcomeReturned, CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnOutcomeReturnAttempted:
		return true
	}
	return false
}

// Not returned. Required if and only if `return_outcome` is `not_returned`.
type CardDisputeNewParamsVisaConsumerQualityMerchandiseNotReturned struct {
}

func (r CardDisputeNewParamsVisaConsumerQualityMerchandiseNotReturned) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Ongoing negotiations. Exclude if there is no evidence of ongoing negotiations.
type CardDisputeNewParamsVisaConsumerQualityMerchandiseOngoingNegotiations struct {
	// Explanation of the previous ongoing negotiations between the cardholder and
	// merchant.
	Explanation param.Field[string] `json:"explanation,required"`
	// Date the cardholder first notified the issuer of the dispute.
	IssuerFirstNotifiedAt param.Field[time.Time] `json:"issuer_first_notified_at,required" format:"date"`
	// Started at.
	StartedAt param.Field[time.Time] `json:"started_at,required" format:"date"`
}

func (r CardDisputeNewParamsVisaConsumerQualityMerchandiseOngoingNegotiations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Return attempted. Required if and only if `return_outcome` is
// `return_attempted`.
type CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnAttempted struct {
	// Attempt explanation.
	AttemptExplanation param.Field[string] `json:"attempt_explanation,required"`
	// Attempt reason.
	AttemptReason param.Field[CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnAttemptedAttemptReason] `json:"attempt_reason,required"`
	// Attempted at.
	AttemptedAt param.Field[time.Time] `json:"attempted_at,required" format:"date"`
	// Merchandise disposition.
	MerchandiseDisposition param.Field[string] `json:"merchandise_disposition,required"`
}

func (r CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnAttempted) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Attempt reason.
type CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnAttemptedAttemptReason string

const (
	CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnAttemptedAttemptReasonMerchantNotResponding         CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnAttemptedAttemptReason = "merchant_not_responding"
	CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnAttemptedAttemptReasonNoReturnAuthorizationProvided CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnAttemptedAttemptReason = "no_return_authorization_provided"
	CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnAttemptedAttemptReasonNoReturnInstructions          CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnAttemptedAttemptReason = "no_return_instructions"
	CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnAttemptedAttemptReasonRequestedNotToReturn          CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnAttemptedAttemptReason = "requested_not_to_return"
	CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnAttemptedAttemptReasonReturnNotAccepted             CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnAttemptedAttemptReason = "return_not_accepted"
)

func (r CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnAttemptedAttemptReason) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnAttemptedAttemptReasonMerchantNotResponding, CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnAttemptedAttemptReasonNoReturnAuthorizationProvided, CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnAttemptedAttemptReasonNoReturnInstructions, CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnAttemptedAttemptReasonRequestedNotToReturn, CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnAttemptedAttemptReasonReturnNotAccepted:
		return true
	}
	return false
}

// Returned. Required if and only if `return_outcome` is `returned`.
type CardDisputeNewParamsVisaConsumerQualityMerchandiseReturned struct {
	// Return method.
	ReturnMethod param.Field[CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnedReturnMethod] `json:"return_method,required"`
	// Returned at.
	ReturnedAt param.Field[time.Time] `json:"returned_at,required" format:"date"`
	// Merchant received return at.
	MerchantReceivedReturnAt param.Field[time.Time] `json:"merchant_received_return_at" format:"date"`
	// Other explanation. Required if and only if the return method is `other`.
	OtherExplanation param.Field[string] `json:"other_explanation"`
	// Tracking number.
	TrackingNumber param.Field[string] `json:"tracking_number"`
}

func (r CardDisputeNewParamsVisaConsumerQualityMerchandiseReturned) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Return method.
type CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnedReturnMethod string

const (
	CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnedReturnMethodDhl           CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnedReturnMethod = "dhl"
	CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnedReturnMethodFaceToFace    CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnedReturnMethod = "face_to_face"
	CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnedReturnMethodFedex         CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnedReturnMethod = "fedex"
	CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnedReturnMethodOther         CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnedReturnMethod = "other"
	CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnedReturnMethodPostalService CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnedReturnMethod = "postal_service"
	CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnedReturnMethodUps           CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnedReturnMethod = "ups"
)

func (r CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnedReturnMethod) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnedReturnMethodDhl, CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnedReturnMethodFaceToFace, CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnedReturnMethodFedex, CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnedReturnMethodOther, CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnedReturnMethodPostalService, CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnedReturnMethodUps:
		return true
	}
	return false
}

// Services quality issue. Required if and only if `category` is
// `consumer_quality_services`.
type CardDisputeNewParamsVisaConsumerQualityServices struct {
	// Cardholder cancellation.
	CardholderCancellation param.Field[CardDisputeNewParamsVisaConsumerQualityServicesCardholderCancellation] `json:"cardholder_cancellation,required"`
	// Non-fiat currency or non-fungible token related and not matching description.
	NonFiatCurrencyOrNonFungibleTokenRelatedAndNotMatchingDescription param.Field[CardDisputeNewParamsVisaConsumerQualityServicesNonFiatCurrencyOrNonFungibleTokenRelatedAndNotMatchingDescription] `json:"non_fiat_currency_or_non_fungible_token_related_and_not_matching_description,required"`
	// Purchase information and quality issue.
	PurchaseInfoAndQualityIssue param.Field[string] `json:"purchase_info_and_quality_issue,required"`
	// Services received at.
	ServicesReceivedAt param.Field[time.Time] `json:"services_received_at,required" format:"date"`
	// Cardholder paid to have work redone.
	CardholderPaidToHaveWorkRedone param.Field[CardDisputeNewParamsVisaConsumerQualityServicesCardholderPaidToHaveWorkRedone] `json:"cardholder_paid_to_have_work_redone"`
	// Ongoing negotiations. Exclude if there is no evidence of ongoing negotiations.
	OngoingNegotiations param.Field[CardDisputeNewParamsVisaConsumerQualityServicesOngoingNegotiations] `json:"ongoing_negotiations"`
	// Whether the dispute is related to the quality of food from an eating place or
	// restaurant. Must be provided when Merchant Category Code (MCC) is 5812, 5813
	// or 5814.
	RestaurantFoodRelated param.Field[CardDisputeNewParamsVisaConsumerQualityServicesRestaurantFoodRelated] `json:"restaurant_food_related"`
}

func (r CardDisputeNewParamsVisaConsumerQualityServices) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Cardholder cancellation.
type CardDisputeNewParamsVisaConsumerQualityServicesCardholderCancellation struct {
	// Accepted by merchant.
	AcceptedByMerchant param.Field[CardDisputeNewParamsVisaConsumerQualityServicesCardholderCancellationAcceptedByMerchant] `json:"accepted_by_merchant,required"`
	// Canceled at.
	CanceledAt param.Field[time.Time] `json:"canceled_at,required" format:"date"`
	// Reason.
	Reason param.Field[string] `json:"reason,required"`
}

func (r CardDisputeNewParamsVisaConsumerQualityServicesCardholderCancellation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Accepted by merchant.
type CardDisputeNewParamsVisaConsumerQualityServicesCardholderCancellationAcceptedByMerchant string

const (
	CardDisputeNewParamsVisaConsumerQualityServicesCardholderCancellationAcceptedByMerchantAccepted    CardDisputeNewParamsVisaConsumerQualityServicesCardholderCancellationAcceptedByMerchant = "accepted"
	CardDisputeNewParamsVisaConsumerQualityServicesCardholderCancellationAcceptedByMerchantNotAccepted CardDisputeNewParamsVisaConsumerQualityServicesCardholderCancellationAcceptedByMerchant = "not_accepted"
)

func (r CardDisputeNewParamsVisaConsumerQualityServicesCardholderCancellationAcceptedByMerchant) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsVisaConsumerQualityServicesCardholderCancellationAcceptedByMerchantAccepted, CardDisputeNewParamsVisaConsumerQualityServicesCardholderCancellationAcceptedByMerchantNotAccepted:
		return true
	}
	return false
}

// Non-fiat currency or non-fungible token related and not matching description.
type CardDisputeNewParamsVisaConsumerQualityServicesNonFiatCurrencyOrNonFungibleTokenRelatedAndNotMatchingDescription string

const (
	CardDisputeNewParamsVisaConsumerQualityServicesNonFiatCurrencyOrNonFungibleTokenRelatedAndNotMatchingDescriptionNotRelated CardDisputeNewParamsVisaConsumerQualityServicesNonFiatCurrencyOrNonFungibleTokenRelatedAndNotMatchingDescription = "not_related"
	CardDisputeNewParamsVisaConsumerQualityServicesNonFiatCurrencyOrNonFungibleTokenRelatedAndNotMatchingDescriptionRelated    CardDisputeNewParamsVisaConsumerQualityServicesNonFiatCurrencyOrNonFungibleTokenRelatedAndNotMatchingDescription = "related"
)

func (r CardDisputeNewParamsVisaConsumerQualityServicesNonFiatCurrencyOrNonFungibleTokenRelatedAndNotMatchingDescription) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsVisaConsumerQualityServicesNonFiatCurrencyOrNonFungibleTokenRelatedAndNotMatchingDescriptionNotRelated, CardDisputeNewParamsVisaConsumerQualityServicesNonFiatCurrencyOrNonFungibleTokenRelatedAndNotMatchingDescriptionRelated:
		return true
	}
	return false
}

// Cardholder paid to have work redone.
type CardDisputeNewParamsVisaConsumerQualityServicesCardholderPaidToHaveWorkRedone string

const (
	CardDisputeNewParamsVisaConsumerQualityServicesCardholderPaidToHaveWorkRedoneDidNotPayToHaveWorkRedone CardDisputeNewParamsVisaConsumerQualityServicesCardholderPaidToHaveWorkRedone = "did_not_pay_to_have_work_redone"
	CardDisputeNewParamsVisaConsumerQualityServicesCardholderPaidToHaveWorkRedonePaidToHaveWorkRedone      CardDisputeNewParamsVisaConsumerQualityServicesCardholderPaidToHaveWorkRedone = "paid_to_have_work_redone"
)

func (r CardDisputeNewParamsVisaConsumerQualityServicesCardholderPaidToHaveWorkRedone) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsVisaConsumerQualityServicesCardholderPaidToHaveWorkRedoneDidNotPayToHaveWorkRedone, CardDisputeNewParamsVisaConsumerQualityServicesCardholderPaidToHaveWorkRedonePaidToHaveWorkRedone:
		return true
	}
	return false
}

// Ongoing negotiations. Exclude if there is no evidence of ongoing negotiations.
type CardDisputeNewParamsVisaConsumerQualityServicesOngoingNegotiations struct {
	// Explanation of the previous ongoing negotiations between the cardholder and
	// merchant.
	Explanation param.Field[string] `json:"explanation,required"`
	// Date the cardholder first notified the issuer of the dispute.
	IssuerFirstNotifiedAt param.Field[time.Time] `json:"issuer_first_notified_at,required" format:"date"`
	// Started at.
	StartedAt param.Field[time.Time] `json:"started_at,required" format:"date"`
}

func (r CardDisputeNewParamsVisaConsumerQualityServicesOngoingNegotiations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether the dispute is related to the quality of food from an eating place or
// restaurant. Must be provided when Merchant Category Code (MCC) is 5812, 5813
// or 5814.
type CardDisputeNewParamsVisaConsumerQualityServicesRestaurantFoodRelated string

const (
	CardDisputeNewParamsVisaConsumerQualityServicesRestaurantFoodRelatedNotRelated CardDisputeNewParamsVisaConsumerQualityServicesRestaurantFoodRelated = "not_related"
	CardDisputeNewParamsVisaConsumerQualityServicesRestaurantFoodRelatedRelated    CardDisputeNewParamsVisaConsumerQualityServicesRestaurantFoodRelated = "related"
)

func (r CardDisputeNewParamsVisaConsumerQualityServicesRestaurantFoodRelated) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsVisaConsumerQualityServicesRestaurantFoodRelatedNotRelated, CardDisputeNewParamsVisaConsumerQualityServicesRestaurantFoodRelatedRelated:
		return true
	}
	return false
}

// Services misrepresentation. Required if and only if `category` is
// `consumer_services_misrepresentation`.
type CardDisputeNewParamsVisaConsumerServicesMisrepresentation struct {
	// Cardholder cancellation.
	CardholderCancellation param.Field[CardDisputeNewParamsVisaConsumerServicesMisrepresentationCardholderCancellation] `json:"cardholder_cancellation,required"`
	// Merchant resolution attempted.
	MerchantResolutionAttempted param.Field[CardDisputeNewParamsVisaConsumerServicesMisrepresentationMerchantResolutionAttempted] `json:"merchant_resolution_attempted,required"`
	// Misrepresentation explanation.
	MisrepresentationExplanation param.Field[string] `json:"misrepresentation_explanation,required"`
	// Purchase explanation.
	PurchaseExplanation param.Field[string] `json:"purchase_explanation,required"`
	// Received at.
	ReceivedAt param.Field[time.Time] `json:"received_at,required" format:"date"`
}

func (r CardDisputeNewParamsVisaConsumerServicesMisrepresentation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Cardholder cancellation.
type CardDisputeNewParamsVisaConsumerServicesMisrepresentationCardholderCancellation struct {
	// Accepted by merchant.
	AcceptedByMerchant param.Field[CardDisputeNewParamsVisaConsumerServicesMisrepresentationCardholderCancellationAcceptedByMerchant] `json:"accepted_by_merchant,required"`
	// Canceled at.
	CanceledAt param.Field[time.Time] `json:"canceled_at,required" format:"date"`
	// Reason.
	Reason param.Field[string] `json:"reason,required"`
}

func (r CardDisputeNewParamsVisaConsumerServicesMisrepresentationCardholderCancellation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Accepted by merchant.
type CardDisputeNewParamsVisaConsumerServicesMisrepresentationCardholderCancellationAcceptedByMerchant string

const (
	CardDisputeNewParamsVisaConsumerServicesMisrepresentationCardholderCancellationAcceptedByMerchantAccepted    CardDisputeNewParamsVisaConsumerServicesMisrepresentationCardholderCancellationAcceptedByMerchant = "accepted"
	CardDisputeNewParamsVisaConsumerServicesMisrepresentationCardholderCancellationAcceptedByMerchantNotAccepted CardDisputeNewParamsVisaConsumerServicesMisrepresentationCardholderCancellationAcceptedByMerchant = "not_accepted"
)

func (r CardDisputeNewParamsVisaConsumerServicesMisrepresentationCardholderCancellationAcceptedByMerchant) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsVisaConsumerServicesMisrepresentationCardholderCancellationAcceptedByMerchantAccepted, CardDisputeNewParamsVisaConsumerServicesMisrepresentationCardholderCancellationAcceptedByMerchantNotAccepted:
		return true
	}
	return false
}

// Merchant resolution attempted.
type CardDisputeNewParamsVisaConsumerServicesMisrepresentationMerchantResolutionAttempted string

const (
	CardDisputeNewParamsVisaConsumerServicesMisrepresentationMerchantResolutionAttemptedAttempted            CardDisputeNewParamsVisaConsumerServicesMisrepresentationMerchantResolutionAttempted = "attempted"
	CardDisputeNewParamsVisaConsumerServicesMisrepresentationMerchantResolutionAttemptedProhibitedByLocalLaw CardDisputeNewParamsVisaConsumerServicesMisrepresentationMerchantResolutionAttempted = "prohibited_by_local_law"
)

func (r CardDisputeNewParamsVisaConsumerServicesMisrepresentationMerchantResolutionAttempted) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsVisaConsumerServicesMisrepresentationMerchantResolutionAttemptedAttempted, CardDisputeNewParamsVisaConsumerServicesMisrepresentationMerchantResolutionAttemptedProhibitedByLocalLaw:
		return true
	}
	return false
}

// Services not as described. Required if and only if `category` is
// `consumer_services_not_as_described`.
type CardDisputeNewParamsVisaConsumerServicesNotAsDescribed struct {
	// Cardholder cancellation.
	CardholderCancellation param.Field[CardDisputeNewParamsVisaConsumerServicesNotAsDescribedCardholderCancellation] `json:"cardholder_cancellation,required"`
	// Explanation of what was ordered and was not as described.
	Explanation param.Field[string] `json:"explanation,required"`
	// Merchant resolution attempted.
	MerchantResolutionAttempted param.Field[CardDisputeNewParamsVisaConsumerServicesNotAsDescribedMerchantResolutionAttempted] `json:"merchant_resolution_attempted,required"`
	// Received at.
	ReceivedAt param.Field[time.Time] `json:"received_at,required" format:"date"`
}

func (r CardDisputeNewParamsVisaConsumerServicesNotAsDescribed) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Cardholder cancellation.
type CardDisputeNewParamsVisaConsumerServicesNotAsDescribedCardholderCancellation struct {
	// Accepted by merchant.
	AcceptedByMerchant param.Field[CardDisputeNewParamsVisaConsumerServicesNotAsDescribedCardholderCancellationAcceptedByMerchant] `json:"accepted_by_merchant,required"`
	// Canceled at.
	CanceledAt param.Field[time.Time] `json:"canceled_at,required" format:"date"`
	// Reason.
	Reason param.Field[string] `json:"reason,required"`
}

func (r CardDisputeNewParamsVisaConsumerServicesNotAsDescribedCardholderCancellation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Accepted by merchant.
type CardDisputeNewParamsVisaConsumerServicesNotAsDescribedCardholderCancellationAcceptedByMerchant string

const (
	CardDisputeNewParamsVisaConsumerServicesNotAsDescribedCardholderCancellationAcceptedByMerchantAccepted    CardDisputeNewParamsVisaConsumerServicesNotAsDescribedCardholderCancellationAcceptedByMerchant = "accepted"
	CardDisputeNewParamsVisaConsumerServicesNotAsDescribedCardholderCancellationAcceptedByMerchantNotAccepted CardDisputeNewParamsVisaConsumerServicesNotAsDescribedCardholderCancellationAcceptedByMerchant = "not_accepted"
)

func (r CardDisputeNewParamsVisaConsumerServicesNotAsDescribedCardholderCancellationAcceptedByMerchant) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsVisaConsumerServicesNotAsDescribedCardholderCancellationAcceptedByMerchantAccepted, CardDisputeNewParamsVisaConsumerServicesNotAsDescribedCardholderCancellationAcceptedByMerchantNotAccepted:
		return true
	}
	return false
}

// Merchant resolution attempted.
type CardDisputeNewParamsVisaConsumerServicesNotAsDescribedMerchantResolutionAttempted string

const (
	CardDisputeNewParamsVisaConsumerServicesNotAsDescribedMerchantResolutionAttemptedAttempted            CardDisputeNewParamsVisaConsumerServicesNotAsDescribedMerchantResolutionAttempted = "attempted"
	CardDisputeNewParamsVisaConsumerServicesNotAsDescribedMerchantResolutionAttemptedProhibitedByLocalLaw CardDisputeNewParamsVisaConsumerServicesNotAsDescribedMerchantResolutionAttempted = "prohibited_by_local_law"
)

func (r CardDisputeNewParamsVisaConsumerServicesNotAsDescribedMerchantResolutionAttempted) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsVisaConsumerServicesNotAsDescribedMerchantResolutionAttemptedAttempted, CardDisputeNewParamsVisaConsumerServicesNotAsDescribedMerchantResolutionAttemptedProhibitedByLocalLaw:
		return true
	}
	return false
}

// Services not received. Required if and only if `category` is
// `consumer_services_not_received`.
type CardDisputeNewParamsVisaConsumerServicesNotReceived struct {
	// Cancellation outcome.
	CancellationOutcome param.Field[CardDisputeNewParamsVisaConsumerServicesNotReceivedCancellationOutcome] `json:"cancellation_outcome,required"`
	// Last expected receipt at.
	LastExpectedReceiptAt param.Field[time.Time] `json:"last_expected_receipt_at,required" format:"date"`
	// Merchant resolution attempted.
	MerchantResolutionAttempted param.Field[CardDisputeNewParamsVisaConsumerServicesNotReceivedMerchantResolutionAttempted] `json:"merchant_resolution_attempted,required"`
	// Purchase information and explanation.
	PurchaseInfoAndExplanation param.Field[string] `json:"purchase_info_and_explanation,required"`
	// Cardholder cancellation prior to expected receipt. Required if and only if
	// `cancellation_outcome` is `cardholder_cancellation_prior_to_expected_receipt`.
	CardholderCancellationPriorToExpectedReceipt param.Field[CardDisputeNewParamsVisaConsumerServicesNotReceivedCardholderCancellationPriorToExpectedReceipt] `json:"cardholder_cancellation_prior_to_expected_receipt"`
	// Merchant cancellation. Required if and only if `cancellation_outcome` is
	// `merchant_cancellation`.
	MerchantCancellation param.Field[CardDisputeNewParamsVisaConsumerServicesNotReceivedMerchantCancellation] `json:"merchant_cancellation"`
	// No cancellation. Required if and only if `cancellation_outcome` is
	// `no_cancellation`.
	NoCancellation param.Field[CardDisputeNewParamsVisaConsumerServicesNotReceivedNoCancellation] `json:"no_cancellation"`
}

func (r CardDisputeNewParamsVisaConsumerServicesNotReceived) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Cancellation outcome.
type CardDisputeNewParamsVisaConsumerServicesNotReceivedCancellationOutcome string

const (
	CardDisputeNewParamsVisaConsumerServicesNotReceivedCancellationOutcomeCardholderCancellationPriorToExpectedReceipt CardDisputeNewParamsVisaConsumerServicesNotReceivedCancellationOutcome = "cardholder_cancellation_prior_to_expected_receipt"
	CardDisputeNewParamsVisaConsumerServicesNotReceivedCancellationOutcomeMerchantCancellation                         CardDisputeNewParamsVisaConsumerServicesNotReceivedCancellationOutcome = "merchant_cancellation"
	CardDisputeNewParamsVisaConsumerServicesNotReceivedCancellationOutcomeNoCancellation                               CardDisputeNewParamsVisaConsumerServicesNotReceivedCancellationOutcome = "no_cancellation"
)

func (r CardDisputeNewParamsVisaConsumerServicesNotReceivedCancellationOutcome) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsVisaConsumerServicesNotReceivedCancellationOutcomeCardholderCancellationPriorToExpectedReceipt, CardDisputeNewParamsVisaConsumerServicesNotReceivedCancellationOutcomeMerchantCancellation, CardDisputeNewParamsVisaConsumerServicesNotReceivedCancellationOutcomeNoCancellation:
		return true
	}
	return false
}

// Merchant resolution attempted.
type CardDisputeNewParamsVisaConsumerServicesNotReceivedMerchantResolutionAttempted string

const (
	CardDisputeNewParamsVisaConsumerServicesNotReceivedMerchantResolutionAttemptedAttempted            CardDisputeNewParamsVisaConsumerServicesNotReceivedMerchantResolutionAttempted = "attempted"
	CardDisputeNewParamsVisaConsumerServicesNotReceivedMerchantResolutionAttemptedProhibitedByLocalLaw CardDisputeNewParamsVisaConsumerServicesNotReceivedMerchantResolutionAttempted = "prohibited_by_local_law"
)

func (r CardDisputeNewParamsVisaConsumerServicesNotReceivedMerchantResolutionAttempted) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsVisaConsumerServicesNotReceivedMerchantResolutionAttemptedAttempted, CardDisputeNewParamsVisaConsumerServicesNotReceivedMerchantResolutionAttemptedProhibitedByLocalLaw:
		return true
	}
	return false
}

// Cardholder cancellation prior to expected receipt. Required if and only if
// `cancellation_outcome` is `cardholder_cancellation_prior_to_expected_receipt`.
type CardDisputeNewParamsVisaConsumerServicesNotReceivedCardholderCancellationPriorToExpectedReceipt struct {
	// Canceled at.
	CanceledAt param.Field[time.Time] `json:"canceled_at,required" format:"date"`
	// Reason.
	Reason param.Field[string] `json:"reason"`
}

func (r CardDisputeNewParamsVisaConsumerServicesNotReceivedCardholderCancellationPriorToExpectedReceipt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Merchant cancellation. Required if and only if `cancellation_outcome` is
// `merchant_cancellation`.
type CardDisputeNewParamsVisaConsumerServicesNotReceivedMerchantCancellation struct {
	// Canceled at.
	CanceledAt param.Field[time.Time] `json:"canceled_at,required" format:"date"`
}

func (r CardDisputeNewParamsVisaConsumerServicesNotReceivedMerchantCancellation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// No cancellation. Required if and only if `cancellation_outcome` is
// `no_cancellation`.
type CardDisputeNewParamsVisaConsumerServicesNotReceivedNoCancellation struct {
}

func (r CardDisputeNewParamsVisaConsumerServicesNotReceivedNoCancellation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Fraud. Required if and only if `category` is `fraud`.
type CardDisputeNewParamsVisaFraud struct {
	// Fraud type.
	FraudType param.Field[CardDisputeNewParamsVisaFraudFraudType] `json:"fraud_type,required"`
}

func (r CardDisputeNewParamsVisaFraud) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Fraud type.
type CardDisputeNewParamsVisaFraudFraudType string

const (
	CardDisputeNewParamsVisaFraudFraudTypeAccountOrCredentialsTakeover CardDisputeNewParamsVisaFraudFraudType = "account_or_credentials_takeover"
	CardDisputeNewParamsVisaFraudFraudTypeCardNotReceivedAsIssued      CardDisputeNewParamsVisaFraudFraudType = "card_not_received_as_issued"
	CardDisputeNewParamsVisaFraudFraudTypeFraudulentApplication        CardDisputeNewParamsVisaFraudFraudType = "fraudulent_application"
	CardDisputeNewParamsVisaFraudFraudTypeFraudulentUseOfAccountNumber CardDisputeNewParamsVisaFraudFraudType = "fraudulent_use_of_account_number"
	CardDisputeNewParamsVisaFraudFraudTypeIncorrectProcessing          CardDisputeNewParamsVisaFraudFraudType = "incorrect_processing"
	CardDisputeNewParamsVisaFraudFraudTypeIssuerReportedCounterfeit    CardDisputeNewParamsVisaFraudFraudType = "issuer_reported_counterfeit"
	CardDisputeNewParamsVisaFraudFraudTypeLost                         CardDisputeNewParamsVisaFraudFraudType = "lost"
	CardDisputeNewParamsVisaFraudFraudTypeManipulationOfAccountHolder  CardDisputeNewParamsVisaFraudFraudType = "manipulation_of_account_holder"
	CardDisputeNewParamsVisaFraudFraudTypeMerchantMisrepresentation    CardDisputeNewParamsVisaFraudFraudType = "merchant_misrepresentation"
	CardDisputeNewParamsVisaFraudFraudTypeMiscellaneous                CardDisputeNewParamsVisaFraudFraudType = "miscellaneous"
	CardDisputeNewParamsVisaFraudFraudTypeStolen                       CardDisputeNewParamsVisaFraudFraudType = "stolen"
)

func (r CardDisputeNewParamsVisaFraudFraudType) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsVisaFraudFraudTypeAccountOrCredentialsTakeover, CardDisputeNewParamsVisaFraudFraudTypeCardNotReceivedAsIssued, CardDisputeNewParamsVisaFraudFraudTypeFraudulentApplication, CardDisputeNewParamsVisaFraudFraudTypeFraudulentUseOfAccountNumber, CardDisputeNewParamsVisaFraudFraudTypeIncorrectProcessing, CardDisputeNewParamsVisaFraudFraudTypeIssuerReportedCounterfeit, CardDisputeNewParamsVisaFraudFraudTypeLost, CardDisputeNewParamsVisaFraudFraudTypeManipulationOfAccountHolder, CardDisputeNewParamsVisaFraudFraudTypeMerchantMisrepresentation, CardDisputeNewParamsVisaFraudFraudTypeMiscellaneous, CardDisputeNewParamsVisaFraudFraudTypeStolen:
		return true
	}
	return false
}

// Processing error. Required if and only if `category` is `processing_error`.
type CardDisputeNewParamsVisaProcessingError struct {
	// Error reason.
	ErrorReason param.Field[CardDisputeNewParamsVisaProcessingErrorErrorReason] `json:"error_reason,required"`
	// Merchant resolution attempted.
	MerchantResolutionAttempted param.Field[CardDisputeNewParamsVisaProcessingErrorMerchantResolutionAttempted] `json:"merchant_resolution_attempted,required"`
	// Duplicate transaction. Required if and only if `error_reason` is
	// `duplicate_transaction`.
	DuplicateTransaction param.Field[CardDisputeNewParamsVisaProcessingErrorDuplicateTransaction] `json:"duplicate_transaction"`
	// Incorrect amount. Required if and only if `error_reason` is `incorrect_amount`.
	IncorrectAmount param.Field[CardDisputeNewParamsVisaProcessingErrorIncorrectAmount] `json:"incorrect_amount"`
	// Paid by other means. Required if and only if `error_reason` is
	// `paid_by_other_means`.
	PaidByOtherMeans param.Field[CardDisputeNewParamsVisaProcessingErrorPaidByOtherMeans] `json:"paid_by_other_means"`
}

func (r CardDisputeNewParamsVisaProcessingError) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Error reason.
type CardDisputeNewParamsVisaProcessingErrorErrorReason string

const (
	CardDisputeNewParamsVisaProcessingErrorErrorReasonDuplicateTransaction CardDisputeNewParamsVisaProcessingErrorErrorReason = "duplicate_transaction"
	CardDisputeNewParamsVisaProcessingErrorErrorReasonIncorrectAmount      CardDisputeNewParamsVisaProcessingErrorErrorReason = "incorrect_amount"
	CardDisputeNewParamsVisaProcessingErrorErrorReasonPaidByOtherMeans     CardDisputeNewParamsVisaProcessingErrorErrorReason = "paid_by_other_means"
)

func (r CardDisputeNewParamsVisaProcessingErrorErrorReason) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsVisaProcessingErrorErrorReasonDuplicateTransaction, CardDisputeNewParamsVisaProcessingErrorErrorReasonIncorrectAmount, CardDisputeNewParamsVisaProcessingErrorErrorReasonPaidByOtherMeans:
		return true
	}
	return false
}

// Merchant resolution attempted.
type CardDisputeNewParamsVisaProcessingErrorMerchantResolutionAttempted string

const (
	CardDisputeNewParamsVisaProcessingErrorMerchantResolutionAttemptedAttempted            CardDisputeNewParamsVisaProcessingErrorMerchantResolutionAttempted = "attempted"
	CardDisputeNewParamsVisaProcessingErrorMerchantResolutionAttemptedProhibitedByLocalLaw CardDisputeNewParamsVisaProcessingErrorMerchantResolutionAttempted = "prohibited_by_local_law"
)

func (r CardDisputeNewParamsVisaProcessingErrorMerchantResolutionAttempted) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsVisaProcessingErrorMerchantResolutionAttemptedAttempted, CardDisputeNewParamsVisaProcessingErrorMerchantResolutionAttemptedProhibitedByLocalLaw:
		return true
	}
	return false
}

// Duplicate transaction. Required if and only if `error_reason` is
// `duplicate_transaction`.
type CardDisputeNewParamsVisaProcessingErrorDuplicateTransaction struct {
	// Other transaction ID.
	OtherTransactionID param.Field[string] `json:"other_transaction_id,required"`
}

func (r CardDisputeNewParamsVisaProcessingErrorDuplicateTransaction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Incorrect amount. Required if and only if `error_reason` is `incorrect_amount`.
type CardDisputeNewParamsVisaProcessingErrorIncorrectAmount struct {
	// Expected amount.
	ExpectedAmount param.Field[int64] `json:"expected_amount,required"`
}

func (r CardDisputeNewParamsVisaProcessingErrorIncorrectAmount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Paid by other means. Required if and only if `error_reason` is
// `paid_by_other_means`.
type CardDisputeNewParamsVisaProcessingErrorPaidByOtherMeans struct {
	// Other form of payment evidence.
	OtherFormOfPaymentEvidence param.Field[CardDisputeNewParamsVisaProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidence] `json:"other_form_of_payment_evidence,required"`
	// Other transaction ID.
	OtherTransactionID param.Field[string] `json:"other_transaction_id"`
}

func (r CardDisputeNewParamsVisaProcessingErrorPaidByOtherMeans) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Other form of payment evidence.
type CardDisputeNewParamsVisaProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidence string

const (
	CardDisputeNewParamsVisaProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidenceCanceledCheck   CardDisputeNewParamsVisaProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidence = "canceled_check"
	CardDisputeNewParamsVisaProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidenceCardTransaction CardDisputeNewParamsVisaProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidence = "card_transaction"
	CardDisputeNewParamsVisaProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidenceCashReceipt     CardDisputeNewParamsVisaProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidence = "cash_receipt"
	CardDisputeNewParamsVisaProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidenceOther           CardDisputeNewParamsVisaProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidence = "other"
	CardDisputeNewParamsVisaProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidenceStatement       CardDisputeNewParamsVisaProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidence = "statement"
	CardDisputeNewParamsVisaProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidenceVoucher         CardDisputeNewParamsVisaProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidence = "voucher"
)

func (r CardDisputeNewParamsVisaProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidence) IsKnown() bool {
	switch r {
	case CardDisputeNewParamsVisaProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidenceCanceledCheck, CardDisputeNewParamsVisaProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidenceCardTransaction, CardDisputeNewParamsVisaProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidenceCashReceipt, CardDisputeNewParamsVisaProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidenceOther, CardDisputeNewParamsVisaProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidenceStatement, CardDisputeNewParamsVisaProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidenceVoucher:
		return true
	}
	return false
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
	CardDisputeListParamsStatusInUserSubmissionRequired          CardDisputeListParamsStatusIn = "user_submission_required"
	CardDisputeListParamsStatusInPendingUserSubmissionReviewing  CardDisputeListParamsStatusIn = "pending_user_submission_reviewing"
	CardDisputeListParamsStatusInPendingUserSubmissionSubmitting CardDisputeListParamsStatusIn = "pending_user_submission_submitting"
	CardDisputeListParamsStatusInPendingUserWithdrawalSubmitting CardDisputeListParamsStatusIn = "pending_user_withdrawal_submitting"
	CardDisputeListParamsStatusInPendingResponse                 CardDisputeListParamsStatusIn = "pending_response"
	CardDisputeListParamsStatusInLost                            CardDisputeListParamsStatusIn = "lost"
	CardDisputeListParamsStatusInWon                             CardDisputeListParamsStatusIn = "won"
)

func (r CardDisputeListParamsStatusIn) IsKnown() bool {
	switch r {
	case CardDisputeListParamsStatusInUserSubmissionRequired, CardDisputeListParamsStatusInPendingUserSubmissionReviewing, CardDisputeListParamsStatusInPendingUserSubmissionSubmitting, CardDisputeListParamsStatusInPendingUserWithdrawalSubmitting, CardDisputeListParamsStatusInPendingResponse, CardDisputeListParamsStatusInLost, CardDisputeListParamsStatusInWon:
		return true
	}
	return false
}

type CardDisputeSubmitUserSubmissionParams struct {
	// The network of the Card Dispute. Details specific to the network are required
	// under the sub-object with the same identifier as the network.
	Network param.Field[CardDisputeSubmitUserSubmissionParamsNetwork] `json:"network,required"`
	// The adjusted monetary amount of the part of the transaction that is being
	// disputed. This is optional and will default to the most recent amount provided.
	// If provided, the amount must be less than or equal to the amount of the
	// transaction.
	Amount param.Field[int64] `json:"amount"`
	// The files to be attached to the user submission.
	AttachmentFiles param.Field[[]CardDisputeSubmitUserSubmissionParamsAttachmentFile] `json:"attachment_files"`
	// The free-form explanation provided to Increase to provide more context for the
	// user submission. This field is not sent directly to the card networks.
	Explanation param.Field[string] `json:"explanation"`
	// The Visa-specific parameters for the dispute. Required if and only if `network`
	// is `visa`.
	Visa param.Field[CardDisputeSubmitUserSubmissionParamsVisa] `json:"visa"`
}

func (r CardDisputeSubmitUserSubmissionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The network of the Card Dispute. Details specific to the network are required
// under the sub-object with the same identifier as the network.
type CardDisputeSubmitUserSubmissionParamsNetwork string

const (
	CardDisputeSubmitUserSubmissionParamsNetworkVisa CardDisputeSubmitUserSubmissionParamsNetwork = "visa"
)

func (r CardDisputeSubmitUserSubmissionParamsNetwork) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsNetworkVisa:
		return true
	}
	return false
}

type CardDisputeSubmitUserSubmissionParamsAttachmentFile struct {
	// The ID of the file to be attached. The file must have a `purpose` of
	// `card_dispute_attachment`.
	FileID param.Field[string] `json:"file_id,required"`
}

func (r CardDisputeSubmitUserSubmissionParamsAttachmentFile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The Visa-specific parameters for the dispute. Required if and only if `network`
// is `visa`.
type CardDisputeSubmitUserSubmissionParamsVisa struct {
	// The category of the user submission. Details specific to the category are
	// required under the sub-object with the same identifier as the category.
	Category param.Field[CardDisputeSubmitUserSubmissionParamsVisaCategory] `json:"category,required"`
	// The chargeback details for the user submission. Required if and only if
	// `category` is `chargeback`.
	Chargeback param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargeback] `json:"chargeback"`
	// The merchant pre-arbitration decline details for the user submission. Required
	// if and only if `category` is `merchant_prearbitration_decline`.
	MerchantPrearbitrationDecline param.Field[CardDisputeSubmitUserSubmissionParamsVisaMerchantPrearbitrationDecline] `json:"merchant_prearbitration_decline"`
	// The user pre-arbitration details for the user submission. Required if and only
	// if `category` is `user_prearbitration`.
	UserPrearbitration param.Field[CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitration] `json:"user_prearbitration"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisa) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The category of the user submission. Details specific to the category are
// required under the sub-object with the same identifier as the category.
type CardDisputeSubmitUserSubmissionParamsVisaCategory string

const (
	CardDisputeSubmitUserSubmissionParamsVisaCategoryChargeback                    CardDisputeSubmitUserSubmissionParamsVisaCategory = "chargeback"
	CardDisputeSubmitUserSubmissionParamsVisaCategoryMerchantPrearbitrationDecline CardDisputeSubmitUserSubmissionParamsVisaCategory = "merchant_prearbitration_decline"
	CardDisputeSubmitUserSubmissionParamsVisaCategoryUserPrearbitration            CardDisputeSubmitUserSubmissionParamsVisaCategory = "user_prearbitration"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaCategory) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaCategoryChargeback, CardDisputeSubmitUserSubmissionParamsVisaCategoryMerchantPrearbitrationDecline, CardDisputeSubmitUserSubmissionParamsVisaCategoryUserPrearbitration:
		return true
	}
	return false
}

// The chargeback details for the user submission. Required if and only if
// `category` is `chargeback`.
type CardDisputeSubmitUserSubmissionParamsVisaChargeback struct {
	// Category.
	Category param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackCategory] `json:"category,required"`
	// Authorization. Required if and only if `category` is `authorization`.
	Authorization param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackAuthorization] `json:"authorization"`
	// Canceled merchandise. Required if and only if `category` is
	// `consumer_canceled_merchandise`.
	ConsumerCanceledMerchandise param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandise] `json:"consumer_canceled_merchandise"`
	// Canceled recurring transaction. Required if and only if `category` is
	// `consumer_canceled_recurring_transaction`.
	ConsumerCanceledRecurringTransaction param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledRecurringTransaction] `json:"consumer_canceled_recurring_transaction"`
	// Canceled services. Required if and only if `category` is
	// `consumer_canceled_services`.
	ConsumerCanceledServices param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServices] `json:"consumer_canceled_services"`
	// Counterfeit merchandise. Required if and only if `category` is
	// `consumer_counterfeit_merchandise`.
	ConsumerCounterfeitMerchandise param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCounterfeitMerchandise] `json:"consumer_counterfeit_merchandise"`
	// Credit not processed. Required if and only if `category` is
	// `consumer_credit_not_processed`.
	ConsumerCreditNotProcessed param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCreditNotProcessed] `json:"consumer_credit_not_processed"`
	// Damaged or defective merchandise. Required if and only if `category` is
	// `consumer_damaged_or_defective_merchandise`.
	ConsumerDamagedOrDefectiveMerchandise param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandise] `json:"consumer_damaged_or_defective_merchandise"`
	// Merchandise misrepresentation. Required if and only if `category` is
	// `consumer_merchandise_misrepresentation`.
	ConsumerMerchandiseMisrepresentation param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentation] `json:"consumer_merchandise_misrepresentation"`
	// Merchandise not as described. Required if and only if `category` is
	// `consumer_merchandise_not_as_described`.
	ConsumerMerchandiseNotAsDescribed param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribed] `json:"consumer_merchandise_not_as_described"`
	// Merchandise not received. Required if and only if `category` is
	// `consumer_merchandise_not_received`.
	ConsumerMerchandiseNotReceived param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceived] `json:"consumer_merchandise_not_received"`
	// Non-receipt of cash. Required if and only if `category` is
	// `consumer_non_receipt_of_cash`.
	ConsumerNonReceiptOfCash param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerNonReceiptOfCash] `json:"consumer_non_receipt_of_cash"`
	// Original Credit Transaction (OCT) not accepted. Required if and only if
	// `category` is `consumer_original_credit_transaction_not_accepted`.
	ConsumerOriginalCreditTransactionNotAccepted param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerOriginalCreditTransactionNotAccepted] `json:"consumer_original_credit_transaction_not_accepted"`
	// Merchandise quality issue. Required if and only if `category` is
	// `consumer_quality_merchandise`.
	ConsumerQualityMerchandise param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandise] `json:"consumer_quality_merchandise"`
	// Services quality issue. Required if and only if `category` is
	// `consumer_quality_services`.
	ConsumerQualityServices param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServices] `json:"consumer_quality_services"`
	// Services misrepresentation. Required if and only if `category` is
	// `consumer_services_misrepresentation`.
	ConsumerServicesMisrepresentation param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesMisrepresentation] `json:"consumer_services_misrepresentation"`
	// Services not as described. Required if and only if `category` is
	// `consumer_services_not_as_described`.
	ConsumerServicesNotAsDescribed param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotAsDescribed] `json:"consumer_services_not_as_described"`
	// Services not received. Required if and only if `category` is
	// `consumer_services_not_received`.
	ConsumerServicesNotReceived param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotReceived] `json:"consumer_services_not_received"`
	// Fraud. Required if and only if `category` is `fraud`.
	Fraud param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackFraud] `json:"fraud"`
	// Processing error. Required if and only if `category` is `processing_error`.
	ProcessingError param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingError] `json:"processing_error"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargeback) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Category.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackCategory string

const (
	CardDisputeSubmitUserSubmissionParamsVisaChargebackCategoryAuthorization                                CardDisputeSubmitUserSubmissionParamsVisaChargebackCategory = "authorization"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackCategoryConsumerCanceledMerchandise                  CardDisputeSubmitUserSubmissionParamsVisaChargebackCategory = "consumer_canceled_merchandise"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackCategoryConsumerCanceledRecurringTransaction         CardDisputeSubmitUserSubmissionParamsVisaChargebackCategory = "consumer_canceled_recurring_transaction"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackCategoryConsumerCanceledServices                     CardDisputeSubmitUserSubmissionParamsVisaChargebackCategory = "consumer_canceled_services"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackCategoryConsumerCounterfeitMerchandise               CardDisputeSubmitUserSubmissionParamsVisaChargebackCategory = "consumer_counterfeit_merchandise"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackCategoryConsumerCreditNotProcessed                   CardDisputeSubmitUserSubmissionParamsVisaChargebackCategory = "consumer_credit_not_processed"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackCategoryConsumerDamagedOrDefectiveMerchandise        CardDisputeSubmitUserSubmissionParamsVisaChargebackCategory = "consumer_damaged_or_defective_merchandise"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackCategoryConsumerMerchandiseMisrepresentation         CardDisputeSubmitUserSubmissionParamsVisaChargebackCategory = "consumer_merchandise_misrepresentation"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackCategoryConsumerMerchandiseNotAsDescribed            CardDisputeSubmitUserSubmissionParamsVisaChargebackCategory = "consumer_merchandise_not_as_described"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackCategoryConsumerMerchandiseNotReceived               CardDisputeSubmitUserSubmissionParamsVisaChargebackCategory = "consumer_merchandise_not_received"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackCategoryConsumerNonReceiptOfCash                     CardDisputeSubmitUserSubmissionParamsVisaChargebackCategory = "consumer_non_receipt_of_cash"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackCategoryConsumerOriginalCreditTransactionNotAccepted CardDisputeSubmitUserSubmissionParamsVisaChargebackCategory = "consumer_original_credit_transaction_not_accepted"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackCategoryConsumerQualityMerchandise                   CardDisputeSubmitUserSubmissionParamsVisaChargebackCategory = "consumer_quality_merchandise"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackCategoryConsumerQualityServices                      CardDisputeSubmitUserSubmissionParamsVisaChargebackCategory = "consumer_quality_services"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackCategoryConsumerServicesMisrepresentation            CardDisputeSubmitUserSubmissionParamsVisaChargebackCategory = "consumer_services_misrepresentation"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackCategoryConsumerServicesNotAsDescribed               CardDisputeSubmitUserSubmissionParamsVisaChargebackCategory = "consumer_services_not_as_described"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackCategoryConsumerServicesNotReceived                  CardDisputeSubmitUserSubmissionParamsVisaChargebackCategory = "consumer_services_not_received"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackCategoryFraud                                        CardDisputeSubmitUserSubmissionParamsVisaChargebackCategory = "fraud"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackCategoryProcessingError                              CardDisputeSubmitUserSubmissionParamsVisaChargebackCategory = "processing_error"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackCategory) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaChargebackCategoryAuthorization, CardDisputeSubmitUserSubmissionParamsVisaChargebackCategoryConsumerCanceledMerchandise, CardDisputeSubmitUserSubmissionParamsVisaChargebackCategoryConsumerCanceledRecurringTransaction, CardDisputeSubmitUserSubmissionParamsVisaChargebackCategoryConsumerCanceledServices, CardDisputeSubmitUserSubmissionParamsVisaChargebackCategoryConsumerCounterfeitMerchandise, CardDisputeSubmitUserSubmissionParamsVisaChargebackCategoryConsumerCreditNotProcessed, CardDisputeSubmitUserSubmissionParamsVisaChargebackCategoryConsumerDamagedOrDefectiveMerchandise, CardDisputeSubmitUserSubmissionParamsVisaChargebackCategoryConsumerMerchandiseMisrepresentation, CardDisputeSubmitUserSubmissionParamsVisaChargebackCategoryConsumerMerchandiseNotAsDescribed, CardDisputeSubmitUserSubmissionParamsVisaChargebackCategoryConsumerMerchandiseNotReceived, CardDisputeSubmitUserSubmissionParamsVisaChargebackCategoryConsumerNonReceiptOfCash, CardDisputeSubmitUserSubmissionParamsVisaChargebackCategoryConsumerOriginalCreditTransactionNotAccepted, CardDisputeSubmitUserSubmissionParamsVisaChargebackCategoryConsumerQualityMerchandise, CardDisputeSubmitUserSubmissionParamsVisaChargebackCategoryConsumerQualityServices, CardDisputeSubmitUserSubmissionParamsVisaChargebackCategoryConsumerServicesMisrepresentation, CardDisputeSubmitUserSubmissionParamsVisaChargebackCategoryConsumerServicesNotAsDescribed, CardDisputeSubmitUserSubmissionParamsVisaChargebackCategoryConsumerServicesNotReceived, CardDisputeSubmitUserSubmissionParamsVisaChargebackCategoryFraud, CardDisputeSubmitUserSubmissionParamsVisaChargebackCategoryProcessingError:
		return true
	}
	return false
}

// Authorization. Required if and only if `category` is `authorization`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackAuthorization struct {
	// Account status.
	AccountStatus param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackAuthorizationAccountStatus] `json:"account_status,required"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackAuthorization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Account status.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackAuthorizationAccountStatus string

const (
	CardDisputeSubmitUserSubmissionParamsVisaChargebackAuthorizationAccountStatusAccountClosed CardDisputeSubmitUserSubmissionParamsVisaChargebackAuthorizationAccountStatus = "account_closed"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackAuthorizationAccountStatusCreditProblem CardDisputeSubmitUserSubmissionParamsVisaChargebackAuthorizationAccountStatus = "credit_problem"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackAuthorizationAccountStatusFraud         CardDisputeSubmitUserSubmissionParamsVisaChargebackAuthorizationAccountStatus = "fraud"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackAuthorizationAccountStatus) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaChargebackAuthorizationAccountStatusAccountClosed, CardDisputeSubmitUserSubmissionParamsVisaChargebackAuthorizationAccountStatusCreditProblem, CardDisputeSubmitUserSubmissionParamsVisaChargebackAuthorizationAccountStatusFraud:
		return true
	}
	return false
}

// Canceled merchandise. Required if and only if `category` is
// `consumer_canceled_merchandise`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandise struct {
	// Merchant resolution attempted.
	MerchantResolutionAttempted param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseMerchantResolutionAttempted] `json:"merchant_resolution_attempted,required"`
	// Purchase explanation.
	PurchaseExplanation param.Field[string] `json:"purchase_explanation,required"`
	// Received or expected at.
	ReceivedOrExpectedAt param.Field[time.Time] `json:"received_or_expected_at,required" format:"date"`
	// Return outcome.
	ReturnOutcome param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnOutcome] `json:"return_outcome,required"`
	// Cardholder cancellation.
	CardholderCancellation param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseCardholderCancellation] `json:"cardholder_cancellation"`
	// Not returned. Required if and only if `return_outcome` is `not_returned`.
	NotReturned param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseNotReturned] `json:"not_returned"`
	// Return attempted. Required if and only if `return_outcome` is
	// `return_attempted`.
	ReturnAttempted param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnAttempted] `json:"return_attempted"`
	// Returned. Required if and only if `return_outcome` is `returned`.
	Returned param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturned] `json:"returned"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandise) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Merchant resolution attempted.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseMerchantResolutionAttempted string

const (
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseMerchantResolutionAttemptedAttempted            CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseMerchantResolutionAttempted = "attempted"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseMerchantResolutionAttemptedProhibitedByLocalLaw CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseMerchantResolutionAttempted = "prohibited_by_local_law"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseMerchantResolutionAttempted) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseMerchantResolutionAttemptedAttempted, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseMerchantResolutionAttemptedProhibitedByLocalLaw:
		return true
	}
	return false
}

// Return outcome.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnOutcome string

const (
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnOutcomeNotReturned     CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnOutcome = "not_returned"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnOutcomeReturned        CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnOutcome = "returned"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnOutcomeReturnAttempted CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnOutcome = "return_attempted"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnOutcome) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnOutcomeNotReturned, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnOutcomeReturned, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnOutcomeReturnAttempted:
		return true
	}
	return false
}

// Cardholder cancellation.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseCardholderCancellation struct {
	// Canceled at.
	CanceledAt param.Field[time.Time] `json:"canceled_at,required" format:"date"`
	// Canceled prior to ship date.
	CanceledPriorToShipDate param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseCardholderCancellationCanceledPriorToShipDate] `json:"canceled_prior_to_ship_date,required"`
	// Cancellation policy provided.
	CancellationPolicyProvided param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseCardholderCancellationCancellationPolicyProvided] `json:"cancellation_policy_provided,required"`
	// Reason.
	Reason param.Field[string] `json:"reason,required"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseCardholderCancellation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Canceled prior to ship date.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseCardholderCancellationCanceledPriorToShipDate string

const (
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseCardholderCancellationCanceledPriorToShipDateCanceledPriorToShipDate    CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseCardholderCancellationCanceledPriorToShipDate = "canceled_prior_to_ship_date"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseCardholderCancellationCanceledPriorToShipDateNotCanceledPriorToShipDate CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseCardholderCancellationCanceledPriorToShipDate = "not_canceled_prior_to_ship_date"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseCardholderCancellationCanceledPriorToShipDate) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseCardholderCancellationCanceledPriorToShipDateCanceledPriorToShipDate, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseCardholderCancellationCanceledPriorToShipDateNotCanceledPriorToShipDate:
		return true
	}
	return false
}

// Cancellation policy provided.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseCardholderCancellationCancellationPolicyProvided string

const (
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseCardholderCancellationCancellationPolicyProvidedNotProvided CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseCardholderCancellationCancellationPolicyProvided = "not_provided"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseCardholderCancellationCancellationPolicyProvidedProvided    CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseCardholderCancellationCancellationPolicyProvided = "provided"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseCardholderCancellationCancellationPolicyProvided) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseCardholderCancellationCancellationPolicyProvidedNotProvided, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseCardholderCancellationCancellationPolicyProvidedProvided:
		return true
	}
	return false
}

// Not returned. Required if and only if `return_outcome` is `not_returned`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseNotReturned struct {
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseNotReturned) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Return attempted. Required if and only if `return_outcome` is
// `return_attempted`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnAttempted struct {
	// Attempt explanation.
	AttemptExplanation param.Field[string] `json:"attempt_explanation,required"`
	// Attempt reason.
	AttemptReason param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnAttemptedAttemptReason] `json:"attempt_reason,required"`
	// Attempted at.
	AttemptedAt param.Field[time.Time] `json:"attempted_at,required" format:"date"`
	// Merchandise disposition.
	MerchandiseDisposition param.Field[string] `json:"merchandise_disposition,required"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnAttempted) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Attempt reason.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnAttemptedAttemptReason string

const (
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnAttemptedAttemptReasonMerchantNotResponding         CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnAttemptedAttemptReason = "merchant_not_responding"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnAttemptedAttemptReasonNoReturnAuthorizationProvided CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnAttemptedAttemptReason = "no_return_authorization_provided"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnAttemptedAttemptReasonNoReturnInstructions          CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnAttemptedAttemptReason = "no_return_instructions"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnAttemptedAttemptReasonRequestedNotToReturn          CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnAttemptedAttemptReason = "requested_not_to_return"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnAttemptedAttemptReasonReturnNotAccepted             CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnAttemptedAttemptReason = "return_not_accepted"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnAttemptedAttemptReason) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnAttemptedAttemptReasonMerchantNotResponding, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnAttemptedAttemptReasonNoReturnAuthorizationProvided, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnAttemptedAttemptReasonNoReturnInstructions, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnAttemptedAttemptReasonRequestedNotToReturn, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnAttemptedAttemptReasonReturnNotAccepted:
		return true
	}
	return false
}

// Returned. Required if and only if `return_outcome` is `returned`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturned struct {
	// Return method.
	ReturnMethod param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnedReturnMethod] `json:"return_method,required"`
	// Returned at.
	ReturnedAt param.Field[time.Time] `json:"returned_at,required" format:"date"`
	// Merchant received return at.
	MerchantReceivedReturnAt param.Field[time.Time] `json:"merchant_received_return_at" format:"date"`
	// Other explanation. Required if and only if the return method is `other`.
	OtherExplanation param.Field[string] `json:"other_explanation"`
	// Tracking number.
	TrackingNumber param.Field[string] `json:"tracking_number"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturned) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Return method.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnedReturnMethod string

const (
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnedReturnMethodDhl           CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnedReturnMethod = "dhl"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnedReturnMethodFaceToFace    CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnedReturnMethod = "face_to_face"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnedReturnMethodFedex         CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnedReturnMethod = "fedex"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnedReturnMethodOther         CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnedReturnMethod = "other"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnedReturnMethodPostalService CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnedReturnMethod = "postal_service"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnedReturnMethodUps           CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnedReturnMethod = "ups"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnedReturnMethod) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnedReturnMethodDhl, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnedReturnMethodFaceToFace, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnedReturnMethodFedex, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnedReturnMethodOther, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnedReturnMethodPostalService, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnedReturnMethodUps:
		return true
	}
	return false
}

// Canceled recurring transaction. Required if and only if `category` is
// `consumer_canceled_recurring_transaction`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledRecurringTransaction struct {
	// Cancellation target.
	CancellationTarget param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledRecurringTransactionCancellationTarget] `json:"cancellation_target,required"`
	// Merchant contact methods.
	MerchantContactMethods param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledRecurringTransactionMerchantContactMethods] `json:"merchant_contact_methods,required"`
	// Transaction or account canceled at.
	TransactionOrAccountCanceledAt param.Field[time.Time] `json:"transaction_or_account_canceled_at,required" format:"date"`
	// Other form of payment explanation.
	OtherFormOfPaymentExplanation param.Field[string] `json:"other_form_of_payment_explanation"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledRecurringTransaction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Cancellation target.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledRecurringTransactionCancellationTarget string

const (
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledRecurringTransactionCancellationTargetAccount     CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledRecurringTransactionCancellationTarget = "account"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledRecurringTransactionCancellationTargetTransaction CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledRecurringTransactionCancellationTarget = "transaction"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledRecurringTransactionCancellationTarget) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledRecurringTransactionCancellationTargetAccount, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledRecurringTransactionCancellationTargetTransaction:
		return true
	}
	return false
}

// Merchant contact methods.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledRecurringTransactionMerchantContactMethods struct {
	// Application name.
	ApplicationName param.Field[string] `json:"application_name"`
	// Call center phone number.
	CallCenterPhoneNumber param.Field[string] `json:"call_center_phone_number"`
	// Email address.
	EmailAddress param.Field[string] `json:"email_address"`
	// In person address.
	InPersonAddress param.Field[string] `json:"in_person_address"`
	// Mailing address.
	MailingAddress param.Field[string] `json:"mailing_address"`
	// Text phone number.
	TextPhoneNumber param.Field[string] `json:"text_phone_number"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledRecurringTransactionMerchantContactMethods) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Canceled services. Required if and only if `category` is
// `consumer_canceled_services`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServices struct {
	// Cardholder cancellation.
	CardholderCancellation param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesCardholderCancellation] `json:"cardholder_cancellation,required"`
	// Contracted at.
	ContractedAt param.Field[time.Time] `json:"contracted_at,required" format:"date"`
	// Merchant resolution attempted.
	MerchantResolutionAttempted param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesMerchantResolutionAttempted] `json:"merchant_resolution_attempted,required"`
	// Purchase explanation.
	PurchaseExplanation param.Field[string] `json:"purchase_explanation,required"`
	// Service type.
	ServiceType param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesServiceType] `json:"service_type,required"`
	// Guaranteed reservation explanation. Required if and only if `service_type` is
	// `guaranteed_reservation`.
	GuaranteedReservation param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesGuaranteedReservation] `json:"guaranteed_reservation"`
	// Other service type explanation. Required if and only if `service_type` is
	// `other`.
	Other param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesOther] `json:"other"`
	// Timeshare explanation. Required if and only if `service_type` is `timeshare`.
	Timeshare param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesTimeshare] `json:"timeshare"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServices) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Cardholder cancellation.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesCardholderCancellation struct {
	// Canceled at.
	CanceledAt param.Field[time.Time] `json:"canceled_at,required" format:"date"`
	// Cancellation policy provided.
	CancellationPolicyProvided param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesCardholderCancellationCancellationPolicyProvided] `json:"cancellation_policy_provided,required"`
	// Reason.
	Reason param.Field[string] `json:"reason,required"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesCardholderCancellation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Cancellation policy provided.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesCardholderCancellationCancellationPolicyProvided string

const (
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesCardholderCancellationCancellationPolicyProvidedNotProvided CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesCardholderCancellationCancellationPolicyProvided = "not_provided"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesCardholderCancellationCancellationPolicyProvidedProvided    CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesCardholderCancellationCancellationPolicyProvided = "provided"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesCardholderCancellationCancellationPolicyProvided) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesCardholderCancellationCancellationPolicyProvidedNotProvided, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesCardholderCancellationCancellationPolicyProvidedProvided:
		return true
	}
	return false
}

// Merchant resolution attempted.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesMerchantResolutionAttempted string

const (
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesMerchantResolutionAttemptedAttempted            CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesMerchantResolutionAttempted = "attempted"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesMerchantResolutionAttemptedProhibitedByLocalLaw CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesMerchantResolutionAttempted = "prohibited_by_local_law"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesMerchantResolutionAttempted) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesMerchantResolutionAttemptedAttempted, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesMerchantResolutionAttemptedProhibitedByLocalLaw:
		return true
	}
	return false
}

// Service type.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesServiceType string

const (
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesServiceTypeGuaranteedReservation CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesServiceType = "guaranteed_reservation"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesServiceTypeOther                 CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesServiceType = "other"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesServiceTypeTimeshare             CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesServiceType = "timeshare"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesServiceType) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesServiceTypeGuaranteedReservation, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesServiceTypeOther, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesServiceTypeTimeshare:
		return true
	}
	return false
}

// Guaranteed reservation explanation. Required if and only if `service_type` is
// `guaranteed_reservation`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesGuaranteedReservation struct {
	// Explanation.
	Explanation param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesGuaranteedReservationExplanation] `json:"explanation,required"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesGuaranteedReservation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Explanation.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesGuaranteedReservationExplanation string

const (
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesGuaranteedReservationExplanationCardholderCanceledPriorToService                         CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesGuaranteedReservationExplanation = "cardholder_canceled_prior_to_service"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesGuaranteedReservationExplanationCardholderCancellationAttemptWithin24HoursOfConfirmation CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesGuaranteedReservationExplanation = "cardholder_cancellation_attempt_within_24_hours_of_confirmation"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesGuaranteedReservationExplanationMerchantBilledNoShow                                     CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesGuaranteedReservationExplanation = "merchant_billed_no_show"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesGuaranteedReservationExplanation) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesGuaranteedReservationExplanationCardholderCanceledPriorToService, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesGuaranteedReservationExplanationCardholderCancellationAttemptWithin24HoursOfConfirmation, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesGuaranteedReservationExplanationMerchantBilledNoShow:
		return true
	}
	return false
}

// Other service type explanation. Required if and only if `service_type` is
// `other`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesOther struct {
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesOther) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Timeshare explanation. Required if and only if `service_type` is `timeshare`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesTimeshare struct {
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesTimeshare) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Counterfeit merchandise. Required if and only if `category` is
// `consumer_counterfeit_merchandise`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCounterfeitMerchandise struct {
	// Counterfeit explanation.
	CounterfeitExplanation param.Field[string] `json:"counterfeit_explanation,required"`
	// Disposition explanation.
	DispositionExplanation param.Field[string] `json:"disposition_explanation,required"`
	// Order explanation.
	OrderExplanation param.Field[string] `json:"order_explanation,required"`
	// Received at.
	ReceivedAt param.Field[time.Time] `json:"received_at,required" format:"date"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCounterfeitMerchandise) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Credit not processed. Required if and only if `category` is
// `consumer_credit_not_processed`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCreditNotProcessed struct {
	// Canceled or returned at.
	CanceledOrReturnedAt param.Field[time.Time] `json:"canceled_or_returned_at" format:"date"`
	// Credit expected at.
	CreditExpectedAt param.Field[time.Time] `json:"credit_expected_at" format:"date"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCreditNotProcessed) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Damaged or defective merchandise. Required if and only if `category` is
// `consumer_damaged_or_defective_merchandise`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandise struct {
	// Merchant resolution attempted.
	MerchantResolutionAttempted param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseMerchantResolutionAttempted] `json:"merchant_resolution_attempted,required"`
	// Order and issue explanation.
	OrderAndIssueExplanation param.Field[string] `json:"order_and_issue_explanation,required"`
	// Received at.
	ReceivedAt param.Field[time.Time] `json:"received_at,required" format:"date"`
	// Return outcome.
	ReturnOutcome param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnOutcome] `json:"return_outcome,required"`
	// Not returned. Required if and only if `return_outcome` is `not_returned`.
	NotReturned param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseNotReturned] `json:"not_returned"`
	// Return attempted. Required if and only if `return_outcome` is
	// `return_attempted`.
	ReturnAttempted param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttempted] `json:"return_attempted"`
	// Returned. Required if and only if `return_outcome` is `returned`.
	Returned param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturned] `json:"returned"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandise) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Merchant resolution attempted.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseMerchantResolutionAttempted string

const (
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseMerchantResolutionAttemptedAttempted            CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseMerchantResolutionAttempted = "attempted"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseMerchantResolutionAttemptedProhibitedByLocalLaw CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseMerchantResolutionAttempted = "prohibited_by_local_law"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseMerchantResolutionAttempted) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseMerchantResolutionAttemptedAttempted, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseMerchantResolutionAttemptedProhibitedByLocalLaw:
		return true
	}
	return false
}

// Return outcome.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnOutcome string

const (
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnOutcomeNotReturned     CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnOutcome = "not_returned"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnOutcomeReturned        CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnOutcome = "returned"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnOutcomeReturnAttempted CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnOutcome = "return_attempted"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnOutcome) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnOutcomeNotReturned, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnOutcomeReturned, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnOutcomeReturnAttempted:
		return true
	}
	return false
}

// Not returned. Required if and only if `return_outcome` is `not_returned`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseNotReturned struct {
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseNotReturned) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Return attempted. Required if and only if `return_outcome` is
// `return_attempted`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttempted struct {
	// Attempt explanation.
	AttemptExplanation param.Field[string] `json:"attempt_explanation,required"`
	// Attempt reason.
	AttemptReason param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReason] `json:"attempt_reason,required"`
	// Attempted at.
	AttemptedAt param.Field[time.Time] `json:"attempted_at,required" format:"date"`
	// Merchandise disposition.
	MerchandiseDisposition param.Field[string] `json:"merchandise_disposition,required"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttempted) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Attempt reason.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReason string

const (
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReasonMerchantNotResponding         CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReason = "merchant_not_responding"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReasonNoReturnAuthorizationProvided CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReason = "no_return_authorization_provided"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReasonNoReturnInstructions          CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReason = "no_return_instructions"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReasonRequestedNotToReturn          CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReason = "requested_not_to_return"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReasonReturnNotAccepted             CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReason = "return_not_accepted"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReason) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReasonMerchantNotResponding, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReasonNoReturnAuthorizationProvided, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReasonNoReturnInstructions, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReasonRequestedNotToReturn, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReasonReturnNotAccepted:
		return true
	}
	return false
}

// Returned. Required if and only if `return_outcome` is `returned`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturned struct {
	// Return method.
	ReturnMethod param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethod] `json:"return_method,required"`
	// Returned at.
	ReturnedAt param.Field[time.Time] `json:"returned_at,required" format:"date"`
	// Merchant received return at.
	MerchantReceivedReturnAt param.Field[time.Time] `json:"merchant_received_return_at" format:"date"`
	// Other explanation. Required if and only if the return method is `other`.
	OtherExplanation param.Field[string] `json:"other_explanation"`
	// Tracking number.
	TrackingNumber param.Field[string] `json:"tracking_number"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturned) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Return method.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethod string

const (
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethodDhl           CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethod = "dhl"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethodFaceToFace    CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethod = "face_to_face"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethodFedex         CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethod = "fedex"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethodOther         CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethod = "other"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethodPostalService CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethod = "postal_service"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethodUps           CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethod = "ups"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethod) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethodDhl, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethodFaceToFace, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethodFedex, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethodOther, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethodPostalService, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethodUps:
		return true
	}
	return false
}

// Merchandise misrepresentation. Required if and only if `category` is
// `consumer_merchandise_misrepresentation`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentation struct {
	// Merchant resolution attempted.
	MerchantResolutionAttempted param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationMerchantResolutionAttempted] `json:"merchant_resolution_attempted,required"`
	// Misrepresentation explanation.
	MisrepresentationExplanation param.Field[string] `json:"misrepresentation_explanation,required"`
	// Purchase explanation.
	PurchaseExplanation param.Field[string] `json:"purchase_explanation,required"`
	// Received at.
	ReceivedAt param.Field[time.Time] `json:"received_at,required" format:"date"`
	// Return outcome.
	ReturnOutcome param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnOutcome] `json:"return_outcome,required"`
	// Not returned. Required if and only if `return_outcome` is `not_returned`.
	NotReturned param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationNotReturned] `json:"not_returned"`
	// Return attempted. Required if and only if `return_outcome` is
	// `return_attempted`.
	ReturnAttempted param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnAttempted] `json:"return_attempted"`
	// Returned. Required if and only if `return_outcome` is `returned`.
	Returned param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturned] `json:"returned"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Merchant resolution attempted.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationMerchantResolutionAttempted string

const (
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationMerchantResolutionAttemptedAttempted            CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationMerchantResolutionAttempted = "attempted"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationMerchantResolutionAttemptedProhibitedByLocalLaw CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationMerchantResolutionAttempted = "prohibited_by_local_law"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationMerchantResolutionAttempted) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationMerchantResolutionAttemptedAttempted, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationMerchantResolutionAttemptedProhibitedByLocalLaw:
		return true
	}
	return false
}

// Return outcome.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnOutcome string

const (
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnOutcomeNotReturned     CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnOutcome = "not_returned"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnOutcomeReturned        CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnOutcome = "returned"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnOutcomeReturnAttempted CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnOutcome = "return_attempted"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnOutcome) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnOutcomeNotReturned, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnOutcomeReturned, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnOutcomeReturnAttempted:
		return true
	}
	return false
}

// Not returned. Required if and only if `return_outcome` is `not_returned`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationNotReturned struct {
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationNotReturned) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Return attempted. Required if and only if `return_outcome` is
// `return_attempted`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnAttempted struct {
	// Attempt explanation.
	AttemptExplanation param.Field[string] `json:"attempt_explanation,required"`
	// Attempt reason.
	AttemptReason param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReason] `json:"attempt_reason,required"`
	// Attempted at.
	AttemptedAt param.Field[time.Time] `json:"attempted_at,required" format:"date"`
	// Merchandise disposition.
	MerchandiseDisposition param.Field[string] `json:"merchandise_disposition,required"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnAttempted) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Attempt reason.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReason string

const (
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReasonMerchantNotResponding         CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReason = "merchant_not_responding"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReasonNoReturnAuthorizationProvided CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReason = "no_return_authorization_provided"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReasonNoReturnInstructions          CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReason = "no_return_instructions"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReasonRequestedNotToReturn          CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReason = "requested_not_to_return"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReasonReturnNotAccepted             CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReason = "return_not_accepted"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReason) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReasonMerchantNotResponding, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReasonNoReturnAuthorizationProvided, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReasonNoReturnInstructions, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReasonRequestedNotToReturn, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReasonReturnNotAccepted:
		return true
	}
	return false
}

// Returned. Required if and only if `return_outcome` is `returned`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturned struct {
	// Return method.
	ReturnMethod param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnedReturnMethod] `json:"return_method,required"`
	// Returned at.
	ReturnedAt param.Field[time.Time] `json:"returned_at,required" format:"date"`
	// Merchant received return at.
	MerchantReceivedReturnAt param.Field[time.Time] `json:"merchant_received_return_at" format:"date"`
	// Other explanation. Required if and only if the return method is `other`.
	OtherExplanation param.Field[string] `json:"other_explanation"`
	// Tracking number.
	TrackingNumber param.Field[string] `json:"tracking_number"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturned) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Return method.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnedReturnMethod string

const (
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnedReturnMethodDhl           CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnedReturnMethod = "dhl"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnedReturnMethodFaceToFace    CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnedReturnMethod = "face_to_face"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnedReturnMethodFedex         CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnedReturnMethod = "fedex"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnedReturnMethodOther         CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnedReturnMethod = "other"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnedReturnMethodPostalService CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnedReturnMethod = "postal_service"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnedReturnMethodUps           CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnedReturnMethod = "ups"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnedReturnMethod) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnedReturnMethodDhl, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnedReturnMethodFaceToFace, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnedReturnMethodFedex, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnedReturnMethodOther, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnedReturnMethodPostalService, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnedReturnMethodUps:
		return true
	}
	return false
}

// Merchandise not as described. Required if and only if `category` is
// `consumer_merchandise_not_as_described`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribed struct {
	// Merchant resolution attempted.
	MerchantResolutionAttempted param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedMerchantResolutionAttempted] `json:"merchant_resolution_attempted,required"`
	// Received at.
	ReceivedAt param.Field[time.Time] `json:"received_at,required" format:"date"`
	// Return outcome.
	ReturnOutcome param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnOutcome] `json:"return_outcome,required"`
	// Return attempted. Required if and only if `return_outcome` is
	// `return_attempted`.
	ReturnAttempted param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnAttempted] `json:"return_attempted"`
	// Returned. Required if and only if `return_outcome` is `returned`.
	Returned param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturned] `json:"returned"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribed) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Merchant resolution attempted.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedMerchantResolutionAttempted string

const (
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedMerchantResolutionAttemptedAttempted            CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedMerchantResolutionAttempted = "attempted"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedMerchantResolutionAttemptedProhibitedByLocalLaw CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedMerchantResolutionAttempted = "prohibited_by_local_law"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedMerchantResolutionAttempted) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedMerchantResolutionAttemptedAttempted, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedMerchantResolutionAttemptedProhibitedByLocalLaw:
		return true
	}
	return false
}

// Return outcome.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnOutcome string

const (
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnOutcomeReturned        CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnOutcome = "returned"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnOutcomeReturnAttempted CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnOutcome = "return_attempted"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnOutcome) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnOutcomeReturned, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnOutcomeReturnAttempted:
		return true
	}
	return false
}

// Return attempted. Required if and only if `return_outcome` is
// `return_attempted`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnAttempted struct {
	// Attempt explanation.
	AttemptExplanation param.Field[string] `json:"attempt_explanation,required"`
	// Attempt reason.
	AttemptReason param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReason] `json:"attempt_reason,required"`
	// Attempted at.
	AttemptedAt param.Field[time.Time] `json:"attempted_at,required" format:"date"`
	// Merchandise disposition.
	MerchandiseDisposition param.Field[string] `json:"merchandise_disposition,required"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnAttempted) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Attempt reason.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReason string

const (
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReasonMerchantNotResponding         CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReason = "merchant_not_responding"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReasonNoReturnAuthorizationProvided CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReason = "no_return_authorization_provided"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReasonNoReturnInstructions          CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReason = "no_return_instructions"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReasonRequestedNotToReturn          CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReason = "requested_not_to_return"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReasonReturnNotAccepted             CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReason = "return_not_accepted"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReason) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReasonMerchantNotResponding, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReasonNoReturnAuthorizationProvided, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReasonNoReturnInstructions, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReasonRequestedNotToReturn, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReasonReturnNotAccepted:
		return true
	}
	return false
}

// Returned. Required if and only if `return_outcome` is `returned`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturned struct {
	// Return method.
	ReturnMethod param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnedReturnMethod] `json:"return_method,required"`
	// Returned at.
	ReturnedAt param.Field[time.Time] `json:"returned_at,required" format:"date"`
	// Merchant received return at.
	MerchantReceivedReturnAt param.Field[time.Time] `json:"merchant_received_return_at" format:"date"`
	// Other explanation. Required if and only if the return method is `other`.
	OtherExplanation param.Field[string] `json:"other_explanation"`
	// Tracking number.
	TrackingNumber param.Field[string] `json:"tracking_number"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturned) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Return method.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnedReturnMethod string

const (
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnedReturnMethodDhl           CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnedReturnMethod = "dhl"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnedReturnMethodFaceToFace    CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnedReturnMethod = "face_to_face"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnedReturnMethodFedex         CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnedReturnMethod = "fedex"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnedReturnMethodOther         CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnedReturnMethod = "other"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnedReturnMethodPostalService CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnedReturnMethod = "postal_service"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnedReturnMethodUps           CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnedReturnMethod = "ups"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnedReturnMethod) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnedReturnMethodDhl, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnedReturnMethodFaceToFace, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnedReturnMethodFedex, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnedReturnMethodOther, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnedReturnMethodPostalService, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnedReturnMethodUps:
		return true
	}
	return false
}

// Merchandise not received. Required if and only if `category` is
// `consumer_merchandise_not_received`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceived struct {
	// Cancellation outcome.
	CancellationOutcome param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedCancellationOutcome] `json:"cancellation_outcome,required"`
	// Delivery issue.
	DeliveryIssue param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedDeliveryIssue] `json:"delivery_issue,required"`
	// Last expected receipt at.
	LastExpectedReceiptAt param.Field[time.Time] `json:"last_expected_receipt_at,required" format:"date"`
	// Merchant resolution attempted.
	MerchantResolutionAttempted param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedMerchantResolutionAttempted] `json:"merchant_resolution_attempted,required"`
	// Purchase information and explanation.
	PurchaseInfoAndExplanation param.Field[string] `json:"purchase_info_and_explanation,required"`
	// Cardholder cancellation prior to expected receipt. Required if and only if
	// `cancellation_outcome` is `cardholder_cancellation_prior_to_expected_receipt`.
	CardholderCancellationPriorToExpectedReceipt param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedCardholderCancellationPriorToExpectedReceipt] `json:"cardholder_cancellation_prior_to_expected_receipt"`
	// Delayed. Required if and only if `delivery_issue` is `delayed`.
	Delayed param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedDelayed] `json:"delayed"`
	// Delivered to wrong location. Required if and only if `delivery_issue` is
	// `delivered_to_wrong_location`.
	DeliveredToWrongLocation param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedDeliveredToWrongLocation] `json:"delivered_to_wrong_location"`
	// Merchant cancellation. Required if and only if `cancellation_outcome` is
	// `merchant_cancellation`.
	MerchantCancellation param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedMerchantCancellation] `json:"merchant_cancellation"`
	// No cancellation. Required if and only if `cancellation_outcome` is
	// `no_cancellation`.
	NoCancellation param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedNoCancellation] `json:"no_cancellation"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceived) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Cancellation outcome.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedCancellationOutcome string

const (
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedCancellationOutcomeCardholderCancellationPriorToExpectedReceipt CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedCancellationOutcome = "cardholder_cancellation_prior_to_expected_receipt"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedCancellationOutcomeMerchantCancellation                         CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedCancellationOutcome = "merchant_cancellation"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedCancellationOutcomeNoCancellation                               CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedCancellationOutcome = "no_cancellation"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedCancellationOutcome) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedCancellationOutcomeCardholderCancellationPriorToExpectedReceipt, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedCancellationOutcomeMerchantCancellation, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedCancellationOutcomeNoCancellation:
		return true
	}
	return false
}

// Delivery issue.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedDeliveryIssue string

const (
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedDeliveryIssueDelayed                  CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedDeliveryIssue = "delayed"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedDeliveryIssueDeliveredToWrongLocation CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedDeliveryIssue = "delivered_to_wrong_location"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedDeliveryIssue) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedDeliveryIssueDelayed, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedDeliveryIssueDeliveredToWrongLocation:
		return true
	}
	return false
}

// Merchant resolution attempted.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedMerchantResolutionAttempted string

const (
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedMerchantResolutionAttemptedAttempted            CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedMerchantResolutionAttempted = "attempted"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedMerchantResolutionAttemptedProhibitedByLocalLaw CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedMerchantResolutionAttempted = "prohibited_by_local_law"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedMerchantResolutionAttempted) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedMerchantResolutionAttemptedAttempted, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedMerchantResolutionAttemptedProhibitedByLocalLaw:
		return true
	}
	return false
}

// Cardholder cancellation prior to expected receipt. Required if and only if
// `cancellation_outcome` is `cardholder_cancellation_prior_to_expected_receipt`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedCardholderCancellationPriorToExpectedReceipt struct {
	// Canceled at.
	CanceledAt param.Field[time.Time] `json:"canceled_at,required" format:"date"`
	// Reason.
	Reason param.Field[string] `json:"reason"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedCardholderCancellationPriorToExpectedReceipt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Delayed. Required if and only if `delivery_issue` is `delayed`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedDelayed struct {
	// Explanation.
	Explanation param.Field[string] `json:"explanation,required"`
	// Return outcome.
	ReturnOutcome param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedDelayedReturnOutcome] `json:"return_outcome,required"`
	// Not returned. Required if and only if `return_outcome` is `not_returned`.
	NotReturned param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedDelayedNotReturned] `json:"not_returned"`
	// Return attempted. Required if and only if `return_outcome` is
	// `return_attempted`.
	ReturnAttempted param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedDelayedReturnAttempted] `json:"return_attempted"`
	// Returned. Required if and only if `return_outcome` is `returned`.
	Returned param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedDelayedReturned] `json:"returned"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedDelayed) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Return outcome.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedDelayedReturnOutcome string

const (
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedDelayedReturnOutcomeNotReturned     CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedDelayedReturnOutcome = "not_returned"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedDelayedReturnOutcomeReturned        CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedDelayedReturnOutcome = "returned"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedDelayedReturnOutcomeReturnAttempted CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedDelayedReturnOutcome = "return_attempted"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedDelayedReturnOutcome) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedDelayedReturnOutcomeNotReturned, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedDelayedReturnOutcomeReturned, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedDelayedReturnOutcomeReturnAttempted:
		return true
	}
	return false
}

// Not returned. Required if and only if `return_outcome` is `not_returned`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedDelayedNotReturned struct {
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedDelayedNotReturned) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Return attempted. Required if and only if `return_outcome` is
// `return_attempted`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedDelayedReturnAttempted struct {
	// Attempted at.
	AttemptedAt param.Field[time.Time] `json:"attempted_at,required" format:"date"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedDelayedReturnAttempted) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Returned. Required if and only if `return_outcome` is `returned`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedDelayedReturned struct {
	// Merchant received return at.
	MerchantReceivedReturnAt param.Field[time.Time] `json:"merchant_received_return_at,required" format:"date"`
	// Returned at.
	ReturnedAt param.Field[time.Time] `json:"returned_at,required" format:"date"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedDelayedReturned) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Delivered to wrong location. Required if and only if `delivery_issue` is
// `delivered_to_wrong_location`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedDeliveredToWrongLocation struct {
	// Agreed location.
	AgreedLocation param.Field[string] `json:"agreed_location,required"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedDeliveredToWrongLocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Merchant cancellation. Required if and only if `cancellation_outcome` is
// `merchant_cancellation`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedMerchantCancellation struct {
	// Canceled at.
	CanceledAt param.Field[time.Time] `json:"canceled_at,required" format:"date"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedMerchantCancellation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// No cancellation. Required if and only if `cancellation_outcome` is
// `no_cancellation`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedNoCancellation struct {
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedNoCancellation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Non-receipt of cash. Required if and only if `category` is
// `consumer_non_receipt_of_cash`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerNonReceiptOfCash struct {
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerNonReceiptOfCash) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Original Credit Transaction (OCT) not accepted. Required if and only if
// `category` is `consumer_original_credit_transaction_not_accepted`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerOriginalCreditTransactionNotAccepted struct {
	// Explanation.
	Explanation param.Field[string] `json:"explanation,required"`
	// Reason.
	Reason param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerOriginalCreditTransactionNotAcceptedReason] `json:"reason,required"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerOriginalCreditTransactionNotAccepted) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Reason.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerOriginalCreditTransactionNotAcceptedReason string

const (
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerOriginalCreditTransactionNotAcceptedReasonProhibitedByLocalLawsOrRegulation CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerOriginalCreditTransactionNotAcceptedReason = "prohibited_by_local_laws_or_regulation"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerOriginalCreditTransactionNotAcceptedReasonRecipientRefused                  CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerOriginalCreditTransactionNotAcceptedReason = "recipient_refused"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerOriginalCreditTransactionNotAcceptedReason) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerOriginalCreditTransactionNotAcceptedReasonProhibitedByLocalLawsOrRegulation, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerOriginalCreditTransactionNotAcceptedReasonRecipientRefused:
		return true
	}
	return false
}

// Merchandise quality issue. Required if and only if `category` is
// `consumer_quality_merchandise`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandise struct {
	// Expected at.
	ExpectedAt param.Field[time.Time] `json:"expected_at,required" format:"date"`
	// Merchant resolution attempted.
	MerchantResolutionAttempted param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseMerchantResolutionAttempted] `json:"merchant_resolution_attempted,required"`
	// Purchase information and quality issue.
	PurchaseInfoAndQualityIssue param.Field[string] `json:"purchase_info_and_quality_issue,required"`
	// Received at.
	ReceivedAt param.Field[time.Time] `json:"received_at,required" format:"date"`
	// Return outcome.
	ReturnOutcome param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnOutcome] `json:"return_outcome,required"`
	// Not returned. Required if and only if `return_outcome` is `not_returned`.
	NotReturned param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseNotReturned] `json:"not_returned"`
	// Ongoing negotiations. Exclude if there is no evidence of ongoing negotiations.
	OngoingNegotiations param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseOngoingNegotiations] `json:"ongoing_negotiations"`
	// Return attempted. Required if and only if `return_outcome` is
	// `return_attempted`.
	ReturnAttempted param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnAttempted] `json:"return_attempted"`
	// Returned. Required if and only if `return_outcome` is `returned`.
	Returned param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturned] `json:"returned"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandise) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Merchant resolution attempted.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseMerchantResolutionAttempted string

const (
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseMerchantResolutionAttemptedAttempted            CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseMerchantResolutionAttempted = "attempted"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseMerchantResolutionAttemptedProhibitedByLocalLaw CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseMerchantResolutionAttempted = "prohibited_by_local_law"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseMerchantResolutionAttempted) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseMerchantResolutionAttemptedAttempted, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseMerchantResolutionAttemptedProhibitedByLocalLaw:
		return true
	}
	return false
}

// Return outcome.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnOutcome string

const (
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnOutcomeNotReturned     CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnOutcome = "not_returned"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnOutcomeReturned        CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnOutcome = "returned"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnOutcomeReturnAttempted CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnOutcome = "return_attempted"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnOutcome) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnOutcomeNotReturned, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnOutcomeReturned, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnOutcomeReturnAttempted:
		return true
	}
	return false
}

// Not returned. Required if and only if `return_outcome` is `not_returned`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseNotReturned struct {
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseNotReturned) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Ongoing negotiations. Exclude if there is no evidence of ongoing negotiations.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseOngoingNegotiations struct {
	// Explanation of the previous ongoing negotiations between the cardholder and
	// merchant.
	Explanation param.Field[string] `json:"explanation,required"`
	// Date the cardholder first notified the issuer of the dispute.
	IssuerFirstNotifiedAt param.Field[time.Time] `json:"issuer_first_notified_at,required" format:"date"`
	// Started at.
	StartedAt param.Field[time.Time] `json:"started_at,required" format:"date"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseOngoingNegotiations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Return attempted. Required if and only if `return_outcome` is
// `return_attempted`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnAttempted struct {
	// Attempt explanation.
	AttemptExplanation param.Field[string] `json:"attempt_explanation,required"`
	// Attempt reason.
	AttemptReason param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnAttemptedAttemptReason] `json:"attempt_reason,required"`
	// Attempted at.
	AttemptedAt param.Field[time.Time] `json:"attempted_at,required" format:"date"`
	// Merchandise disposition.
	MerchandiseDisposition param.Field[string] `json:"merchandise_disposition,required"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnAttempted) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Attempt reason.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnAttemptedAttemptReason string

const (
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnAttemptedAttemptReasonMerchantNotResponding         CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnAttemptedAttemptReason = "merchant_not_responding"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnAttemptedAttemptReasonNoReturnAuthorizationProvided CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnAttemptedAttemptReason = "no_return_authorization_provided"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnAttemptedAttemptReasonNoReturnInstructions          CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnAttemptedAttemptReason = "no_return_instructions"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnAttemptedAttemptReasonRequestedNotToReturn          CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnAttemptedAttemptReason = "requested_not_to_return"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnAttemptedAttemptReasonReturnNotAccepted             CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnAttemptedAttemptReason = "return_not_accepted"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnAttemptedAttemptReason) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnAttemptedAttemptReasonMerchantNotResponding, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnAttemptedAttemptReasonNoReturnAuthorizationProvided, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnAttemptedAttemptReasonNoReturnInstructions, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnAttemptedAttemptReasonRequestedNotToReturn, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnAttemptedAttemptReasonReturnNotAccepted:
		return true
	}
	return false
}

// Returned. Required if and only if `return_outcome` is `returned`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturned struct {
	// Return method.
	ReturnMethod param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnedReturnMethod] `json:"return_method,required"`
	// Returned at.
	ReturnedAt param.Field[time.Time] `json:"returned_at,required" format:"date"`
	// Merchant received return at.
	MerchantReceivedReturnAt param.Field[time.Time] `json:"merchant_received_return_at" format:"date"`
	// Other explanation. Required if and only if the return method is `other`.
	OtherExplanation param.Field[string] `json:"other_explanation"`
	// Tracking number.
	TrackingNumber param.Field[string] `json:"tracking_number"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturned) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Return method.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnedReturnMethod string

const (
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnedReturnMethodDhl           CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnedReturnMethod = "dhl"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnedReturnMethodFaceToFace    CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnedReturnMethod = "face_to_face"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnedReturnMethodFedex         CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnedReturnMethod = "fedex"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnedReturnMethodOther         CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnedReturnMethod = "other"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnedReturnMethodPostalService CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnedReturnMethod = "postal_service"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnedReturnMethodUps           CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnedReturnMethod = "ups"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnedReturnMethod) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnedReturnMethodDhl, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnedReturnMethodFaceToFace, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnedReturnMethodFedex, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnedReturnMethodOther, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnedReturnMethodPostalService, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnedReturnMethodUps:
		return true
	}
	return false
}

// Services quality issue. Required if and only if `category` is
// `consumer_quality_services`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServices struct {
	// Cardholder cancellation.
	CardholderCancellation param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServicesCardholderCancellation] `json:"cardholder_cancellation,required"`
	// Non-fiat currency or non-fungible token related and not matching description.
	NonFiatCurrencyOrNonFungibleTokenRelatedAndNotMatchingDescription param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServicesNonFiatCurrencyOrNonFungibleTokenRelatedAndNotMatchingDescription] `json:"non_fiat_currency_or_non_fungible_token_related_and_not_matching_description,required"`
	// Purchase information and quality issue.
	PurchaseInfoAndQualityIssue param.Field[string] `json:"purchase_info_and_quality_issue,required"`
	// Services received at.
	ServicesReceivedAt param.Field[time.Time] `json:"services_received_at,required" format:"date"`
	// Cardholder paid to have work redone.
	CardholderPaidToHaveWorkRedone param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServicesCardholderPaidToHaveWorkRedone] `json:"cardholder_paid_to_have_work_redone"`
	// Ongoing negotiations. Exclude if there is no evidence of ongoing negotiations.
	OngoingNegotiations param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServicesOngoingNegotiations] `json:"ongoing_negotiations"`
	// Whether the dispute is related to the quality of food from an eating place or
	// restaurant. Must be provided when Merchant Category Code (MCC) is 5812, 5813
	// or 5814.
	RestaurantFoodRelated param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServicesRestaurantFoodRelated] `json:"restaurant_food_related"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServices) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Cardholder cancellation.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServicesCardholderCancellation struct {
	// Accepted by merchant.
	AcceptedByMerchant param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServicesCardholderCancellationAcceptedByMerchant] `json:"accepted_by_merchant,required"`
	// Canceled at.
	CanceledAt param.Field[time.Time] `json:"canceled_at,required" format:"date"`
	// Reason.
	Reason param.Field[string] `json:"reason,required"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServicesCardholderCancellation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Accepted by merchant.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServicesCardholderCancellationAcceptedByMerchant string

const (
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServicesCardholderCancellationAcceptedByMerchantAccepted    CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServicesCardholderCancellationAcceptedByMerchant = "accepted"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServicesCardholderCancellationAcceptedByMerchantNotAccepted CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServicesCardholderCancellationAcceptedByMerchant = "not_accepted"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServicesCardholderCancellationAcceptedByMerchant) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServicesCardholderCancellationAcceptedByMerchantAccepted, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServicesCardholderCancellationAcceptedByMerchantNotAccepted:
		return true
	}
	return false
}

// Non-fiat currency or non-fungible token related and not matching description.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServicesNonFiatCurrencyOrNonFungibleTokenRelatedAndNotMatchingDescription string

const (
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServicesNonFiatCurrencyOrNonFungibleTokenRelatedAndNotMatchingDescriptionNotRelated CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServicesNonFiatCurrencyOrNonFungibleTokenRelatedAndNotMatchingDescription = "not_related"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServicesNonFiatCurrencyOrNonFungibleTokenRelatedAndNotMatchingDescriptionRelated    CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServicesNonFiatCurrencyOrNonFungibleTokenRelatedAndNotMatchingDescription = "related"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServicesNonFiatCurrencyOrNonFungibleTokenRelatedAndNotMatchingDescription) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServicesNonFiatCurrencyOrNonFungibleTokenRelatedAndNotMatchingDescriptionNotRelated, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServicesNonFiatCurrencyOrNonFungibleTokenRelatedAndNotMatchingDescriptionRelated:
		return true
	}
	return false
}

// Cardholder paid to have work redone.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServicesCardholderPaidToHaveWorkRedone string

const (
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServicesCardholderPaidToHaveWorkRedoneDidNotPayToHaveWorkRedone CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServicesCardholderPaidToHaveWorkRedone = "did_not_pay_to_have_work_redone"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServicesCardholderPaidToHaveWorkRedonePaidToHaveWorkRedone      CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServicesCardholderPaidToHaveWorkRedone = "paid_to_have_work_redone"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServicesCardholderPaidToHaveWorkRedone) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServicesCardholderPaidToHaveWorkRedoneDidNotPayToHaveWorkRedone, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServicesCardholderPaidToHaveWorkRedonePaidToHaveWorkRedone:
		return true
	}
	return false
}

// Ongoing negotiations. Exclude if there is no evidence of ongoing negotiations.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServicesOngoingNegotiations struct {
	// Explanation of the previous ongoing negotiations between the cardholder and
	// merchant.
	Explanation param.Field[string] `json:"explanation,required"`
	// Date the cardholder first notified the issuer of the dispute.
	IssuerFirstNotifiedAt param.Field[time.Time] `json:"issuer_first_notified_at,required" format:"date"`
	// Started at.
	StartedAt param.Field[time.Time] `json:"started_at,required" format:"date"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServicesOngoingNegotiations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether the dispute is related to the quality of food from an eating place or
// restaurant. Must be provided when Merchant Category Code (MCC) is 5812, 5813
// or 5814.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServicesRestaurantFoodRelated string

const (
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServicesRestaurantFoodRelatedNotRelated CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServicesRestaurantFoodRelated = "not_related"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServicesRestaurantFoodRelatedRelated    CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServicesRestaurantFoodRelated = "related"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServicesRestaurantFoodRelated) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServicesRestaurantFoodRelatedNotRelated, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServicesRestaurantFoodRelatedRelated:
		return true
	}
	return false
}

// Services misrepresentation. Required if and only if `category` is
// `consumer_services_misrepresentation`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesMisrepresentation struct {
	// Cardholder cancellation.
	CardholderCancellation param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesMisrepresentationCardholderCancellation] `json:"cardholder_cancellation,required"`
	// Merchant resolution attempted.
	MerchantResolutionAttempted param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesMisrepresentationMerchantResolutionAttempted] `json:"merchant_resolution_attempted,required"`
	// Misrepresentation explanation.
	MisrepresentationExplanation param.Field[string] `json:"misrepresentation_explanation,required"`
	// Purchase explanation.
	PurchaseExplanation param.Field[string] `json:"purchase_explanation,required"`
	// Received at.
	ReceivedAt param.Field[time.Time] `json:"received_at,required" format:"date"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesMisrepresentation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Cardholder cancellation.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesMisrepresentationCardholderCancellation struct {
	// Accepted by merchant.
	AcceptedByMerchant param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesMisrepresentationCardholderCancellationAcceptedByMerchant] `json:"accepted_by_merchant,required"`
	// Canceled at.
	CanceledAt param.Field[time.Time] `json:"canceled_at,required" format:"date"`
	// Reason.
	Reason param.Field[string] `json:"reason,required"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesMisrepresentationCardholderCancellation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Accepted by merchant.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesMisrepresentationCardholderCancellationAcceptedByMerchant string

const (
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesMisrepresentationCardholderCancellationAcceptedByMerchantAccepted    CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesMisrepresentationCardholderCancellationAcceptedByMerchant = "accepted"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesMisrepresentationCardholderCancellationAcceptedByMerchantNotAccepted CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesMisrepresentationCardholderCancellationAcceptedByMerchant = "not_accepted"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesMisrepresentationCardholderCancellationAcceptedByMerchant) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesMisrepresentationCardholderCancellationAcceptedByMerchantAccepted, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesMisrepresentationCardholderCancellationAcceptedByMerchantNotAccepted:
		return true
	}
	return false
}

// Merchant resolution attempted.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesMisrepresentationMerchantResolutionAttempted string

const (
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesMisrepresentationMerchantResolutionAttemptedAttempted            CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesMisrepresentationMerchantResolutionAttempted = "attempted"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesMisrepresentationMerchantResolutionAttemptedProhibitedByLocalLaw CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesMisrepresentationMerchantResolutionAttempted = "prohibited_by_local_law"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesMisrepresentationMerchantResolutionAttempted) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesMisrepresentationMerchantResolutionAttemptedAttempted, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesMisrepresentationMerchantResolutionAttemptedProhibitedByLocalLaw:
		return true
	}
	return false
}

// Services not as described. Required if and only if `category` is
// `consumer_services_not_as_described`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotAsDescribed struct {
	// Cardholder cancellation.
	CardholderCancellation param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotAsDescribedCardholderCancellation] `json:"cardholder_cancellation,required"`
	// Explanation of what was ordered and was not as described.
	Explanation param.Field[string] `json:"explanation,required"`
	// Merchant resolution attempted.
	MerchantResolutionAttempted param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotAsDescribedMerchantResolutionAttempted] `json:"merchant_resolution_attempted,required"`
	// Received at.
	ReceivedAt param.Field[time.Time] `json:"received_at,required" format:"date"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotAsDescribed) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Cardholder cancellation.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotAsDescribedCardholderCancellation struct {
	// Accepted by merchant.
	AcceptedByMerchant param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotAsDescribedCardholderCancellationAcceptedByMerchant] `json:"accepted_by_merchant,required"`
	// Canceled at.
	CanceledAt param.Field[time.Time] `json:"canceled_at,required" format:"date"`
	// Reason.
	Reason param.Field[string] `json:"reason,required"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotAsDescribedCardholderCancellation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Accepted by merchant.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotAsDescribedCardholderCancellationAcceptedByMerchant string

const (
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotAsDescribedCardholderCancellationAcceptedByMerchantAccepted    CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotAsDescribedCardholderCancellationAcceptedByMerchant = "accepted"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotAsDescribedCardholderCancellationAcceptedByMerchantNotAccepted CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotAsDescribedCardholderCancellationAcceptedByMerchant = "not_accepted"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotAsDescribedCardholderCancellationAcceptedByMerchant) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotAsDescribedCardholderCancellationAcceptedByMerchantAccepted, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotAsDescribedCardholderCancellationAcceptedByMerchantNotAccepted:
		return true
	}
	return false
}

// Merchant resolution attempted.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotAsDescribedMerchantResolutionAttempted string

const (
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotAsDescribedMerchantResolutionAttemptedAttempted            CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotAsDescribedMerchantResolutionAttempted = "attempted"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotAsDescribedMerchantResolutionAttemptedProhibitedByLocalLaw CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotAsDescribedMerchantResolutionAttempted = "prohibited_by_local_law"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotAsDescribedMerchantResolutionAttempted) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotAsDescribedMerchantResolutionAttemptedAttempted, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotAsDescribedMerchantResolutionAttemptedProhibitedByLocalLaw:
		return true
	}
	return false
}

// Services not received. Required if and only if `category` is
// `consumer_services_not_received`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotReceived struct {
	// Cancellation outcome.
	CancellationOutcome param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotReceivedCancellationOutcome] `json:"cancellation_outcome,required"`
	// Last expected receipt at.
	LastExpectedReceiptAt param.Field[time.Time] `json:"last_expected_receipt_at,required" format:"date"`
	// Merchant resolution attempted.
	MerchantResolutionAttempted param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotReceivedMerchantResolutionAttempted] `json:"merchant_resolution_attempted,required"`
	// Purchase information and explanation.
	PurchaseInfoAndExplanation param.Field[string] `json:"purchase_info_and_explanation,required"`
	// Cardholder cancellation prior to expected receipt. Required if and only if
	// `cancellation_outcome` is `cardholder_cancellation_prior_to_expected_receipt`.
	CardholderCancellationPriorToExpectedReceipt param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotReceivedCardholderCancellationPriorToExpectedReceipt] `json:"cardholder_cancellation_prior_to_expected_receipt"`
	// Merchant cancellation. Required if and only if `cancellation_outcome` is
	// `merchant_cancellation`.
	MerchantCancellation param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotReceivedMerchantCancellation] `json:"merchant_cancellation"`
	// No cancellation. Required if and only if `cancellation_outcome` is
	// `no_cancellation`.
	NoCancellation param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotReceivedNoCancellation] `json:"no_cancellation"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotReceived) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Cancellation outcome.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotReceivedCancellationOutcome string

const (
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotReceivedCancellationOutcomeCardholderCancellationPriorToExpectedReceipt CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotReceivedCancellationOutcome = "cardholder_cancellation_prior_to_expected_receipt"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotReceivedCancellationOutcomeMerchantCancellation                         CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotReceivedCancellationOutcome = "merchant_cancellation"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotReceivedCancellationOutcomeNoCancellation                               CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotReceivedCancellationOutcome = "no_cancellation"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotReceivedCancellationOutcome) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotReceivedCancellationOutcomeCardholderCancellationPriorToExpectedReceipt, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotReceivedCancellationOutcomeMerchantCancellation, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotReceivedCancellationOutcomeNoCancellation:
		return true
	}
	return false
}

// Merchant resolution attempted.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotReceivedMerchantResolutionAttempted string

const (
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotReceivedMerchantResolutionAttemptedAttempted            CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotReceivedMerchantResolutionAttempted = "attempted"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotReceivedMerchantResolutionAttemptedProhibitedByLocalLaw CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotReceivedMerchantResolutionAttempted = "prohibited_by_local_law"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotReceivedMerchantResolutionAttempted) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotReceivedMerchantResolutionAttemptedAttempted, CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotReceivedMerchantResolutionAttemptedProhibitedByLocalLaw:
		return true
	}
	return false
}

// Cardholder cancellation prior to expected receipt. Required if and only if
// `cancellation_outcome` is `cardholder_cancellation_prior_to_expected_receipt`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotReceivedCardholderCancellationPriorToExpectedReceipt struct {
	// Canceled at.
	CanceledAt param.Field[time.Time] `json:"canceled_at,required" format:"date"`
	// Reason.
	Reason param.Field[string] `json:"reason"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotReceivedCardholderCancellationPriorToExpectedReceipt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Merchant cancellation. Required if and only if `cancellation_outcome` is
// `merchant_cancellation`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotReceivedMerchantCancellation struct {
	// Canceled at.
	CanceledAt param.Field[time.Time] `json:"canceled_at,required" format:"date"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotReceivedMerchantCancellation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// No cancellation. Required if and only if `cancellation_outcome` is
// `no_cancellation`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotReceivedNoCancellation struct {
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotReceivedNoCancellation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Fraud. Required if and only if `category` is `fraud`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackFraud struct {
	// Fraud type.
	FraudType param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackFraudFraudType] `json:"fraud_type,required"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackFraud) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Fraud type.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackFraudFraudType string

const (
	CardDisputeSubmitUserSubmissionParamsVisaChargebackFraudFraudTypeAccountOrCredentialsTakeover CardDisputeSubmitUserSubmissionParamsVisaChargebackFraudFraudType = "account_or_credentials_takeover"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackFraudFraudTypeCardNotReceivedAsIssued      CardDisputeSubmitUserSubmissionParamsVisaChargebackFraudFraudType = "card_not_received_as_issued"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackFraudFraudTypeFraudulentApplication        CardDisputeSubmitUserSubmissionParamsVisaChargebackFraudFraudType = "fraudulent_application"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackFraudFraudTypeFraudulentUseOfAccountNumber CardDisputeSubmitUserSubmissionParamsVisaChargebackFraudFraudType = "fraudulent_use_of_account_number"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackFraudFraudTypeIncorrectProcessing          CardDisputeSubmitUserSubmissionParamsVisaChargebackFraudFraudType = "incorrect_processing"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackFraudFraudTypeIssuerReportedCounterfeit    CardDisputeSubmitUserSubmissionParamsVisaChargebackFraudFraudType = "issuer_reported_counterfeit"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackFraudFraudTypeLost                         CardDisputeSubmitUserSubmissionParamsVisaChargebackFraudFraudType = "lost"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackFraudFraudTypeManipulationOfAccountHolder  CardDisputeSubmitUserSubmissionParamsVisaChargebackFraudFraudType = "manipulation_of_account_holder"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackFraudFraudTypeMerchantMisrepresentation    CardDisputeSubmitUserSubmissionParamsVisaChargebackFraudFraudType = "merchant_misrepresentation"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackFraudFraudTypeMiscellaneous                CardDisputeSubmitUserSubmissionParamsVisaChargebackFraudFraudType = "miscellaneous"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackFraudFraudTypeStolen                       CardDisputeSubmitUserSubmissionParamsVisaChargebackFraudFraudType = "stolen"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackFraudFraudType) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaChargebackFraudFraudTypeAccountOrCredentialsTakeover, CardDisputeSubmitUserSubmissionParamsVisaChargebackFraudFraudTypeCardNotReceivedAsIssued, CardDisputeSubmitUserSubmissionParamsVisaChargebackFraudFraudTypeFraudulentApplication, CardDisputeSubmitUserSubmissionParamsVisaChargebackFraudFraudTypeFraudulentUseOfAccountNumber, CardDisputeSubmitUserSubmissionParamsVisaChargebackFraudFraudTypeIncorrectProcessing, CardDisputeSubmitUserSubmissionParamsVisaChargebackFraudFraudTypeIssuerReportedCounterfeit, CardDisputeSubmitUserSubmissionParamsVisaChargebackFraudFraudTypeLost, CardDisputeSubmitUserSubmissionParamsVisaChargebackFraudFraudTypeManipulationOfAccountHolder, CardDisputeSubmitUserSubmissionParamsVisaChargebackFraudFraudTypeMerchantMisrepresentation, CardDisputeSubmitUserSubmissionParamsVisaChargebackFraudFraudTypeMiscellaneous, CardDisputeSubmitUserSubmissionParamsVisaChargebackFraudFraudTypeStolen:
		return true
	}
	return false
}

// Processing error. Required if and only if `category` is `processing_error`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingError struct {
	// Error reason.
	ErrorReason param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorErrorReason] `json:"error_reason,required"`
	// Merchant resolution attempted.
	MerchantResolutionAttempted param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorMerchantResolutionAttempted] `json:"merchant_resolution_attempted,required"`
	// Duplicate transaction. Required if and only if `error_reason` is
	// `duplicate_transaction`.
	DuplicateTransaction param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorDuplicateTransaction] `json:"duplicate_transaction"`
	// Incorrect amount. Required if and only if `error_reason` is `incorrect_amount`.
	IncorrectAmount param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorIncorrectAmount] `json:"incorrect_amount"`
	// Paid by other means. Required if and only if `error_reason` is
	// `paid_by_other_means`.
	PaidByOtherMeans param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorPaidByOtherMeans] `json:"paid_by_other_means"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingError) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Error reason.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorErrorReason string

const (
	CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorErrorReasonDuplicateTransaction CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorErrorReason = "duplicate_transaction"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorErrorReasonIncorrectAmount      CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorErrorReason = "incorrect_amount"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorErrorReasonPaidByOtherMeans     CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorErrorReason = "paid_by_other_means"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorErrorReason) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorErrorReasonDuplicateTransaction, CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorErrorReasonIncorrectAmount, CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorErrorReasonPaidByOtherMeans:
		return true
	}
	return false
}

// Merchant resolution attempted.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorMerchantResolutionAttempted string

const (
	CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorMerchantResolutionAttemptedAttempted            CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorMerchantResolutionAttempted = "attempted"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorMerchantResolutionAttemptedProhibitedByLocalLaw CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorMerchantResolutionAttempted = "prohibited_by_local_law"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorMerchantResolutionAttempted) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorMerchantResolutionAttemptedAttempted, CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorMerchantResolutionAttemptedProhibitedByLocalLaw:
		return true
	}
	return false
}

// Duplicate transaction. Required if and only if `error_reason` is
// `duplicate_transaction`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorDuplicateTransaction struct {
	// Other transaction ID.
	OtherTransactionID param.Field[string] `json:"other_transaction_id,required"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorDuplicateTransaction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Incorrect amount. Required if and only if `error_reason` is `incorrect_amount`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorIncorrectAmount struct {
	// Expected amount.
	ExpectedAmount param.Field[int64] `json:"expected_amount,required"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorIncorrectAmount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Paid by other means. Required if and only if `error_reason` is
// `paid_by_other_means`.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorPaidByOtherMeans struct {
	// Other form of payment evidence.
	OtherFormOfPaymentEvidence param.Field[CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidence] `json:"other_form_of_payment_evidence,required"`
	// Other transaction ID.
	OtherTransactionID param.Field[string] `json:"other_transaction_id"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorPaidByOtherMeans) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Other form of payment evidence.
type CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidence string

const (
	CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidenceCanceledCheck   CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidence = "canceled_check"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidenceCardTransaction CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidence = "card_transaction"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidenceCashReceipt     CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidence = "cash_receipt"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidenceOther           CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidence = "other"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidenceStatement       CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidence = "statement"
	CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidenceVoucher         CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidence = "voucher"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidence) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidenceCanceledCheck, CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidenceCardTransaction, CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidenceCashReceipt, CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidenceOther, CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidenceStatement, CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidenceVoucher:
		return true
	}
	return false
}

// The merchant pre-arbitration decline details for the user submission. Required
// if and only if `category` is `merchant_prearbitration_decline`.
type CardDisputeSubmitUserSubmissionParamsVisaMerchantPrearbitrationDecline struct {
	// The reason for declining the merchant's pre-arbitration request.
	Reason param.Field[string] `json:"reason,required"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaMerchantPrearbitrationDecline) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The user pre-arbitration details for the user submission. Required if and only
// if `category` is `user_prearbitration`.
type CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitration struct {
	// The reason for the pre-arbitration request.
	Reason param.Field[string] `json:"reason,required"`
	// Category change details for the pre-arbitration request. Should only be
	// populated if the category of the dispute is being changed as part of the
	// pre-arbitration request.
	CategoryChange param.Field[CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChange] `json:"category_change"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Category change details for the pre-arbitration request. Should only be
// populated if the category of the dispute is being changed as part of the
// pre-arbitration request.
type CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChange struct {
	Category param.Field[CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategory] `json:"category,required"`
	// The reason for the category change.
	Reason param.Field[string] `json:"reason,required"`
}

func (r CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChange) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategory string

const (
	CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategoryAuthorization                                CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategory = "authorization"
	CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategoryConsumerCanceledMerchandise                  CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategory = "consumer_canceled_merchandise"
	CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategoryConsumerCanceledRecurringTransaction         CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategory = "consumer_canceled_recurring_transaction"
	CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategoryConsumerCanceledServices                     CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategory = "consumer_canceled_services"
	CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategoryConsumerCounterfeitMerchandise               CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategory = "consumer_counterfeit_merchandise"
	CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategoryConsumerCreditNotProcessed                   CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategory = "consumer_credit_not_processed"
	CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategoryConsumerDamagedOrDefectiveMerchandise        CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategory = "consumer_damaged_or_defective_merchandise"
	CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategoryConsumerMerchandiseMisrepresentation         CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategory = "consumer_merchandise_misrepresentation"
	CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategoryConsumerMerchandiseNotAsDescribed            CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategory = "consumer_merchandise_not_as_described"
	CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategoryConsumerMerchandiseNotReceived               CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategory = "consumer_merchandise_not_received"
	CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategoryConsumerNonReceiptOfCash                     CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategory = "consumer_non_receipt_of_cash"
	CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategoryConsumerOriginalCreditTransactionNotAccepted CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategory = "consumer_original_credit_transaction_not_accepted"
	CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategoryConsumerQualityMerchandise                   CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategory = "consumer_quality_merchandise"
	CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategoryConsumerQualityServices                      CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategory = "consumer_quality_services"
	CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategoryConsumerServicesMisrepresentation            CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategory = "consumer_services_misrepresentation"
	CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategoryConsumerServicesNotAsDescribed               CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategory = "consumer_services_not_as_described"
	CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategoryConsumerServicesNotReceived                  CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategory = "consumer_services_not_received"
	CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategoryFraud                                        CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategory = "fraud"
	CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategoryProcessingError                              CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategory = "processing_error"
)

func (r CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategory) IsKnown() bool {
	switch r {
	case CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategoryAuthorization, CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategoryConsumerCanceledMerchandise, CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategoryConsumerCanceledRecurringTransaction, CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategoryConsumerCanceledServices, CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategoryConsumerCounterfeitMerchandise, CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategoryConsumerCreditNotProcessed, CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategoryConsumerDamagedOrDefectiveMerchandise, CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategoryConsumerMerchandiseMisrepresentation, CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategoryConsumerMerchandiseNotAsDescribed, CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategoryConsumerMerchandiseNotReceived, CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategoryConsumerNonReceiptOfCash, CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategoryConsumerOriginalCreditTransactionNotAccepted, CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategoryConsumerQualityMerchandise, CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategoryConsumerQualityServices, CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategoryConsumerServicesMisrepresentation, CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategoryConsumerServicesNotAsDescribed, CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategoryConsumerServicesNotReceived, CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategoryFraud, CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategoryProcessingError:
		return true
	}
	return false
}

type CardDisputeWithdrawParams struct {
	// The explanation for withdrawing the Card Dispute.
	Explanation param.Field[string] `json:"explanation"`
}

func (r CardDisputeWithdrawParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

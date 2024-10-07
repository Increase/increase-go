// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Increase/increase-go/internal/apijson"
	"github.com/Increase/increase-go/internal/apiquery"
	"github.com/Increase/increase-go/internal/pagination"
	"github.com/Increase/increase-go/internal/param"
	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
)

// EventSubscriptionService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEventSubscriptionService] method instead.
type EventSubscriptionService struct {
	Options []option.RequestOption
}

// NewEventSubscriptionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEventSubscriptionService(opts ...option.RequestOption) (r *EventSubscriptionService) {
	r = &EventSubscriptionService{}
	r.Options = opts
	return
}

// Create an Event Subscription
func (r *EventSubscriptionService) New(ctx context.Context, body EventSubscriptionNewParams, opts ...option.RequestOption) (res *EventSubscription, err error) {
	opts = append(r.Options[:], opts...)
	path := "event_subscriptions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve an Event Subscription
func (r *EventSubscriptionService) Get(ctx context.Context, eventSubscriptionID string, opts ...option.RequestOption) (res *EventSubscription, err error) {
	opts = append(r.Options[:], opts...)
	if eventSubscriptionID == "" {
		err = errors.New("missing required event_subscription_id parameter")
		return
	}
	path := fmt.Sprintf("event_subscriptions/%s", eventSubscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an Event Subscription
func (r *EventSubscriptionService) Update(ctx context.Context, eventSubscriptionID string, body EventSubscriptionUpdateParams, opts ...option.RequestOption) (res *EventSubscription, err error) {
	opts = append(r.Options[:], opts...)
	if eventSubscriptionID == "" {
		err = errors.New("missing required event_subscription_id parameter")
		return
	}
	path := fmt.Sprintf("event_subscriptions/%s", eventSubscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List Event Subscriptions
func (r *EventSubscriptionService) List(ctx context.Context, query EventSubscriptionListParams, opts ...option.RequestOption) (res *pagination.Page[EventSubscription], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "event_subscriptions"
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

// List Event Subscriptions
func (r *EventSubscriptionService) ListAutoPaging(ctx context.Context, query EventSubscriptionListParams, opts ...option.RequestOption) *pagination.PageAutoPager[EventSubscription] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Webhooks are event notifications we send to you by HTTPS POST requests. Event
// Subscriptions are how you configure your application to listen for them. You can
// create an Event Subscription through your
// [developer dashboard](https://dashboard.increase.com/developers/webhooks) or the
// API. For more information, see our
// [webhooks guide](https://increase.com/documentation/webhooks).
type EventSubscription struct {
	// The event subscription identifier.
	ID string `json:"id,required"`
	// The time the event subscription was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The idempotency key you chose for this object. This value is unique across
	// Increase and is used to ensure that a request is only processed once. Learn more
	// about [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey string `json:"idempotency_key,required,nullable"`
	// If specified, this subscription will only receive webhooks for Events associated
	// with this OAuth Connection.
	OAuthConnectionID string `json:"oauth_connection_id,required,nullable"`
	// If specified, this subscription will only receive webhooks for Events with the
	// specified `category`.
	SelectedEventCategory EventSubscriptionSelectedEventCategory `json:"selected_event_category,required,nullable"`
	// This indicates if we'll send notifications to this subscription.
	Status EventSubscriptionStatus `json:"status,required"`
	// A constant representing the object's type. For this resource it will always be
	// `event_subscription`.
	Type EventSubscriptionType `json:"type,required"`
	// The webhook url where we'll send notifications.
	URL  string                `json:"url,required"`
	JSON eventSubscriptionJSON `json:"-"`
}

// eventSubscriptionJSON contains the JSON metadata for the struct
// [EventSubscription]
type eventSubscriptionJSON struct {
	ID                    apijson.Field
	CreatedAt             apijson.Field
	IdempotencyKey        apijson.Field
	OAuthConnectionID     apijson.Field
	SelectedEventCategory apijson.Field
	Status                apijson.Field
	Type                  apijson.Field
	URL                   apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *EventSubscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventSubscriptionJSON) RawJSON() string {
	return r.raw
}

// If specified, this subscription will only receive webhooks for Events with the
// specified `category`.
type EventSubscriptionSelectedEventCategory string

const (
	// Occurs whenever an Account is created.
	EventSubscriptionSelectedEventCategoryAccountCreated EventSubscriptionSelectedEventCategory = "account.created"
	// Occurs whenever an Account is updated.
	EventSubscriptionSelectedEventCategoryAccountUpdated EventSubscriptionSelectedEventCategory = "account.updated"
	// Occurs whenever an Account Number is created.
	EventSubscriptionSelectedEventCategoryAccountNumberCreated EventSubscriptionSelectedEventCategory = "account_number.created"
	// Occurs whenever an Account Number is updated.
	EventSubscriptionSelectedEventCategoryAccountNumberUpdated EventSubscriptionSelectedEventCategory = "account_number.updated"
	// Occurs whenever an Account Statement is created.
	EventSubscriptionSelectedEventCategoryAccountStatementCreated EventSubscriptionSelectedEventCategory = "account_statement.created"
	// Occurs whenever an Account Transfer is created.
	EventSubscriptionSelectedEventCategoryAccountTransferCreated EventSubscriptionSelectedEventCategory = "account_transfer.created"
	// Occurs whenever an Account Transfer is updated.
	EventSubscriptionSelectedEventCategoryAccountTransferUpdated EventSubscriptionSelectedEventCategory = "account_transfer.updated"
	// Occurs whenever an ACH Prenotification is created.
	EventSubscriptionSelectedEventCategoryACHPrenotificationCreated EventSubscriptionSelectedEventCategory = "ach_prenotification.created"
	// Occurs whenever an ACH Prenotification is updated.
	EventSubscriptionSelectedEventCategoryACHPrenotificationUpdated EventSubscriptionSelectedEventCategory = "ach_prenotification.updated"
	// Occurs whenever an ACH Transfer is created.
	EventSubscriptionSelectedEventCategoryACHTransferCreated EventSubscriptionSelectedEventCategory = "ach_transfer.created"
	// Occurs whenever an ACH Transfer is updated.
	EventSubscriptionSelectedEventCategoryACHTransferUpdated EventSubscriptionSelectedEventCategory = "ach_transfer.updated"
	// Occurs whenever a Bookkeeping Account is created.
	EventSubscriptionSelectedEventCategoryBookkeepingAccountCreated EventSubscriptionSelectedEventCategory = "bookkeeping_account.created"
	// Occurs whenever a Bookkeeping Account is updated.
	EventSubscriptionSelectedEventCategoryBookkeepingAccountUpdated EventSubscriptionSelectedEventCategory = "bookkeeping_account.updated"
	// Occurs whenever a Bookkeeping Entry Set is created.
	EventSubscriptionSelectedEventCategoryBookkeepingEntrySetUpdated EventSubscriptionSelectedEventCategory = "bookkeeping_entry_set.updated"
	// Occurs whenever a Card is created.
	EventSubscriptionSelectedEventCategoryCardCreated EventSubscriptionSelectedEventCategory = "card.created"
	// Occurs whenever a Card is updated.
	EventSubscriptionSelectedEventCategoryCardUpdated EventSubscriptionSelectedEventCategory = "card.updated"
	// Occurs whenever a Card Payment is created.
	EventSubscriptionSelectedEventCategoryCardPaymentCreated EventSubscriptionSelectedEventCategory = "card_payment.created"
	// Occurs whenever a Card Payment is updated.
	EventSubscriptionSelectedEventCategoryCardPaymentUpdated EventSubscriptionSelectedEventCategory = "card_payment.updated"
	// Occurs whenever a Card Profile is created.
	EventSubscriptionSelectedEventCategoryCardProfileCreated EventSubscriptionSelectedEventCategory = "card_profile.created"
	// Occurs whenever a Card Profile is updated.
	EventSubscriptionSelectedEventCategoryCardProfileUpdated EventSubscriptionSelectedEventCategory = "card_profile.updated"
	// Occurs whenever a Card Dispute is created.
	EventSubscriptionSelectedEventCategoryCardDisputeCreated EventSubscriptionSelectedEventCategory = "card_dispute.created"
	// Occurs whenever a Card Dispute is updated.
	EventSubscriptionSelectedEventCategoryCardDisputeUpdated EventSubscriptionSelectedEventCategory = "card_dispute.updated"
	// Occurs whenever a Check Deposit is created.
	EventSubscriptionSelectedEventCategoryCheckDepositCreated EventSubscriptionSelectedEventCategory = "check_deposit.created"
	// Occurs whenever a Check Deposit is updated.
	EventSubscriptionSelectedEventCategoryCheckDepositUpdated EventSubscriptionSelectedEventCategory = "check_deposit.updated"
	// Occurs whenever a Check Transfer is created.
	EventSubscriptionSelectedEventCategoryCheckTransferCreated EventSubscriptionSelectedEventCategory = "check_transfer.created"
	// Occurs whenever a Check Transfer is updated.
	EventSubscriptionSelectedEventCategoryCheckTransferUpdated EventSubscriptionSelectedEventCategory = "check_transfer.updated"
	// Occurs whenever a Declined Transaction is created.
	EventSubscriptionSelectedEventCategoryDeclinedTransactionCreated EventSubscriptionSelectedEventCategory = "declined_transaction.created"
	// Occurs whenever a Digital Card Profile is created.
	EventSubscriptionSelectedEventCategoryDigitalCardProfileCreated EventSubscriptionSelectedEventCategory = "digital_card_profile.created"
	// Occurs whenever a Digital Card Profile is updated.
	EventSubscriptionSelectedEventCategoryDigitalCardProfileUpdated EventSubscriptionSelectedEventCategory = "digital_card_profile.updated"
	// Occurs whenever a Digital Wallet Token is created.
	EventSubscriptionSelectedEventCategoryDigitalWalletTokenCreated EventSubscriptionSelectedEventCategory = "digital_wallet_token.created"
	// Occurs whenever a Digital Wallet Token is updated.
	EventSubscriptionSelectedEventCategoryDigitalWalletTokenUpdated EventSubscriptionSelectedEventCategory = "digital_wallet_token.updated"
	// Occurs whenever a Document is created.
	EventSubscriptionSelectedEventCategoryDocumentCreated EventSubscriptionSelectedEventCategory = "document.created"
	// Occurs whenever an Entity is created.
	EventSubscriptionSelectedEventCategoryEntityCreated EventSubscriptionSelectedEventCategory = "entity.created"
	// Occurs whenever an Entity is updated.
	EventSubscriptionSelectedEventCategoryEntityUpdated EventSubscriptionSelectedEventCategory = "entity.updated"
	// Occurs whenever an Event Subscription is created.
	EventSubscriptionSelectedEventCategoryEventSubscriptionCreated EventSubscriptionSelectedEventCategory = "event_subscription.created"
	// Occurs whenever an Event Subscription is updated.
	EventSubscriptionSelectedEventCategoryEventSubscriptionUpdated EventSubscriptionSelectedEventCategory = "event_subscription.updated"
	// Occurs whenever an Export is created.
	EventSubscriptionSelectedEventCategoryExportCreated EventSubscriptionSelectedEventCategory = "export.created"
	// Occurs whenever an Export is updated.
	EventSubscriptionSelectedEventCategoryExportUpdated EventSubscriptionSelectedEventCategory = "export.updated"
	// Occurs whenever an External Account is created.
	EventSubscriptionSelectedEventCategoryExternalAccountCreated EventSubscriptionSelectedEventCategory = "external_account.created"
	// Occurs whenever an External Account is updated.
	EventSubscriptionSelectedEventCategoryExternalAccountUpdated EventSubscriptionSelectedEventCategory = "external_account.updated"
	// Occurs whenever a File is created.
	EventSubscriptionSelectedEventCategoryFileCreated EventSubscriptionSelectedEventCategory = "file.created"
	// Occurs whenever a Group is updated.
	EventSubscriptionSelectedEventCategoryGroupUpdated EventSubscriptionSelectedEventCategory = "group.updated"
	// Increase may send webhooks with this category to see if a webhook endpoint is
	// working properly.
	EventSubscriptionSelectedEventCategoryGroupHeartbeat EventSubscriptionSelectedEventCategory = "group.heartbeat"
	// Occurs whenever an Inbound ACH Transfer is created.
	EventSubscriptionSelectedEventCategoryInboundACHTransferCreated EventSubscriptionSelectedEventCategory = "inbound_ach_transfer.created"
	// Occurs whenever an Inbound ACH Transfer is updated.
	EventSubscriptionSelectedEventCategoryInboundACHTransferUpdated EventSubscriptionSelectedEventCategory = "inbound_ach_transfer.updated"
	// Occurs whenever an Inbound ACH Transfer Return is created.
	EventSubscriptionSelectedEventCategoryInboundACHTransferReturnCreated EventSubscriptionSelectedEventCategory = "inbound_ach_transfer_return.created"
	// Occurs whenever an Inbound ACH Transfer Return is updated.
	EventSubscriptionSelectedEventCategoryInboundACHTransferReturnUpdated EventSubscriptionSelectedEventCategory = "inbound_ach_transfer_return.updated"
	// Occurs whenever an Inbound Check Deposit is created.
	EventSubscriptionSelectedEventCategoryInboundCheckDepositCreated EventSubscriptionSelectedEventCategory = "inbound_check_deposit.created"
	// Occurs whenever an Inbound Check Deposit is updated.
	EventSubscriptionSelectedEventCategoryInboundCheckDepositUpdated EventSubscriptionSelectedEventCategory = "inbound_check_deposit.updated"
	// Occurs whenever an Inbound Mail Item is created.
	EventSubscriptionSelectedEventCategoryInboundMailItemCreated EventSubscriptionSelectedEventCategory = "inbound_mail_item.created"
	// Occurs whenever an Inbound Mail Item is updated.
	EventSubscriptionSelectedEventCategoryInboundMailItemUpdated EventSubscriptionSelectedEventCategory = "inbound_mail_item.updated"
	// Occurs whenever an Inbound Real-Time Payments Transfer is created.
	EventSubscriptionSelectedEventCategoryInboundRealTimePaymentsTransferCreated EventSubscriptionSelectedEventCategory = "inbound_real_time_payments_transfer.created"
	// Occurs whenever an Inbound Real-Time Payments Transfer is updated.
	EventSubscriptionSelectedEventCategoryInboundRealTimePaymentsTransferUpdated EventSubscriptionSelectedEventCategory = "inbound_real_time_payments_transfer.updated"
	// Occurs whenever an Inbound Wire Drawdown Request is created.
	EventSubscriptionSelectedEventCategoryInboundWireDrawdownRequestCreated EventSubscriptionSelectedEventCategory = "inbound_wire_drawdown_request.created"
	// Occurs whenever an Inbound Wire Transfer is created.
	EventSubscriptionSelectedEventCategoryInboundWireTransferCreated EventSubscriptionSelectedEventCategory = "inbound_wire_transfer.created"
	// Occurs whenever an Inbound Wire Transfer is updated.
	EventSubscriptionSelectedEventCategoryInboundWireTransferUpdated EventSubscriptionSelectedEventCategory = "inbound_wire_transfer.updated"
	// Occurs whenever an IntraFi Account Enrollment is created.
	EventSubscriptionSelectedEventCategoryIntrafiAccountEnrollmentCreated EventSubscriptionSelectedEventCategory = "intrafi_account_enrollment.created"
	// Occurs whenever an IntraFi Account Enrollment is updated.
	EventSubscriptionSelectedEventCategoryIntrafiAccountEnrollmentUpdated EventSubscriptionSelectedEventCategory = "intrafi_account_enrollment.updated"
	// Occurs whenever an IntraFi Exclusion is created.
	EventSubscriptionSelectedEventCategoryIntrafiExclusionCreated EventSubscriptionSelectedEventCategory = "intrafi_exclusion.created"
	// Occurs whenever an IntraFi Exclusion is updated.
	EventSubscriptionSelectedEventCategoryIntrafiExclusionUpdated EventSubscriptionSelectedEventCategory = "intrafi_exclusion.updated"
	// Occurs whenever a Lockbox is created.
	EventSubscriptionSelectedEventCategoryLockboxCreated EventSubscriptionSelectedEventCategory = "lockbox.created"
	// Occurs whenever a Lockbox is updated.
	EventSubscriptionSelectedEventCategoryLockboxUpdated EventSubscriptionSelectedEventCategory = "lockbox.updated"
	// Occurs whenever an OAuth Connection is created.
	EventSubscriptionSelectedEventCategoryOAuthConnectionCreated EventSubscriptionSelectedEventCategory = "oauth_connection.created"
	// Occurs whenever an OAuth Connection is deactivated.
	EventSubscriptionSelectedEventCategoryOAuthConnectionDeactivated EventSubscriptionSelectedEventCategory = "oauth_connection.deactivated"
	// Occurs whenever a Pending Transaction is created.
	EventSubscriptionSelectedEventCategoryPendingTransactionCreated EventSubscriptionSelectedEventCategory = "pending_transaction.created"
	// Occurs whenever a Pending Transaction is updated.
	EventSubscriptionSelectedEventCategoryPendingTransactionUpdated EventSubscriptionSelectedEventCategory = "pending_transaction.updated"
	// Occurs whenever a Physical Card is created.
	EventSubscriptionSelectedEventCategoryPhysicalCardCreated EventSubscriptionSelectedEventCategory = "physical_card.created"
	// Occurs whenever a Physical Card is updated.
	EventSubscriptionSelectedEventCategoryPhysicalCardUpdated EventSubscriptionSelectedEventCategory = "physical_card.updated"
	// Occurs whenever a Physical Card Profile is created.
	EventSubscriptionSelectedEventCategoryPhysicalCardProfileCreated EventSubscriptionSelectedEventCategory = "physical_card_profile.created"
	// Occurs whenever a Physical Card Profile is updated.
	EventSubscriptionSelectedEventCategoryPhysicalCardProfileUpdated EventSubscriptionSelectedEventCategory = "physical_card_profile.updated"
	// Occurs whenever a Proof of Authorization Request is created.
	EventSubscriptionSelectedEventCategoryProofOfAuthorizationRequestCreated EventSubscriptionSelectedEventCategory = "proof_of_authorization_request.created"
	// Occurs whenever a Proof of Authorization Request is updated.
	EventSubscriptionSelectedEventCategoryProofOfAuthorizationRequestUpdated EventSubscriptionSelectedEventCategory = "proof_of_authorization_request.updated"
	// Occurs whenever a Proof of Authorization Request Submission is created.
	EventSubscriptionSelectedEventCategoryProofOfAuthorizationRequestSubmissionCreated EventSubscriptionSelectedEventCategory = "proof_of_authorization_request_submission.created"
	// Occurs whenever a Proof of Authorization Request Submission is updated.
	EventSubscriptionSelectedEventCategoryProofOfAuthorizationRequestSubmissionUpdated EventSubscriptionSelectedEventCategory = "proof_of_authorization_request_submission.updated"
	// Occurs whenever a Real-Time Decision is created in response to a card
	// authorization.
	EventSubscriptionSelectedEventCategoryRealTimeDecisionCardAuthorizationRequested EventSubscriptionSelectedEventCategory = "real_time_decision.card_authorization_requested"
	// Occurs whenever a Real-Time Decision is created in response to a digital wallet
	// provisioning attempt.
	EventSubscriptionSelectedEventCategoryRealTimeDecisionDigitalWalletTokenRequested EventSubscriptionSelectedEventCategory = "real_time_decision.digital_wallet_token_requested"
	// Occurs whenever a Real-Time Decision is created in response to a digital wallet
	// requiring two-factor authentication.
	EventSubscriptionSelectedEventCategoryRealTimeDecisionDigitalWalletAuthenticationRequested EventSubscriptionSelectedEventCategory = "real_time_decision.digital_wallet_authentication_requested"
	// Occurs whenever a Real-Time Decision is created in response to 3DS
	// authentication.
	EventSubscriptionSelectedEventCategoryRealTimeDecisionCardAuthenticationRequested EventSubscriptionSelectedEventCategory = "real_time_decision.card_authentication_requested"
	// Occurs whenever a Real-Time Payments Transfer is created.
	EventSubscriptionSelectedEventCategoryRealTimePaymentsTransferCreated EventSubscriptionSelectedEventCategory = "real_time_payments_transfer.created"
	// Occurs whenever a Real-Time Payments Transfer is updated.
	EventSubscriptionSelectedEventCategoryRealTimePaymentsTransferUpdated EventSubscriptionSelectedEventCategory = "real_time_payments_transfer.updated"
	// Occurs whenever a Real-Time Payments Request for Payment is created.
	EventSubscriptionSelectedEventCategoryRealTimePaymentsRequestForPaymentCreated EventSubscriptionSelectedEventCategory = "real_time_payments_request_for_payment.created"
	// Occurs whenever a Real-Time Payments Request for Payment is updated.
	EventSubscriptionSelectedEventCategoryRealTimePaymentsRequestForPaymentUpdated EventSubscriptionSelectedEventCategory = "real_time_payments_request_for_payment.updated"
	// Occurs whenever a Transaction is created.
	EventSubscriptionSelectedEventCategoryTransactionCreated EventSubscriptionSelectedEventCategory = "transaction.created"
	// Occurs whenever a Wire Drawdown Request is created.
	EventSubscriptionSelectedEventCategoryWireDrawdownRequestCreated EventSubscriptionSelectedEventCategory = "wire_drawdown_request.created"
	// Occurs whenever a Wire Drawdown Request is updated.
	EventSubscriptionSelectedEventCategoryWireDrawdownRequestUpdated EventSubscriptionSelectedEventCategory = "wire_drawdown_request.updated"
	// Occurs whenever a Wire Transfer is created.
	EventSubscriptionSelectedEventCategoryWireTransferCreated EventSubscriptionSelectedEventCategory = "wire_transfer.created"
	// Occurs whenever a Wire Transfer is updated.
	EventSubscriptionSelectedEventCategoryWireTransferUpdated EventSubscriptionSelectedEventCategory = "wire_transfer.updated"
)

func (r EventSubscriptionSelectedEventCategory) IsKnown() bool {
	switch r {
	case EventSubscriptionSelectedEventCategoryAccountCreated, EventSubscriptionSelectedEventCategoryAccountUpdated, EventSubscriptionSelectedEventCategoryAccountNumberCreated, EventSubscriptionSelectedEventCategoryAccountNumberUpdated, EventSubscriptionSelectedEventCategoryAccountStatementCreated, EventSubscriptionSelectedEventCategoryAccountTransferCreated, EventSubscriptionSelectedEventCategoryAccountTransferUpdated, EventSubscriptionSelectedEventCategoryACHPrenotificationCreated, EventSubscriptionSelectedEventCategoryACHPrenotificationUpdated, EventSubscriptionSelectedEventCategoryACHTransferCreated, EventSubscriptionSelectedEventCategoryACHTransferUpdated, EventSubscriptionSelectedEventCategoryBookkeepingAccountCreated, EventSubscriptionSelectedEventCategoryBookkeepingAccountUpdated, EventSubscriptionSelectedEventCategoryBookkeepingEntrySetUpdated, EventSubscriptionSelectedEventCategoryCardCreated, EventSubscriptionSelectedEventCategoryCardUpdated, EventSubscriptionSelectedEventCategoryCardPaymentCreated, EventSubscriptionSelectedEventCategoryCardPaymentUpdated, EventSubscriptionSelectedEventCategoryCardProfileCreated, EventSubscriptionSelectedEventCategoryCardProfileUpdated, EventSubscriptionSelectedEventCategoryCardDisputeCreated, EventSubscriptionSelectedEventCategoryCardDisputeUpdated, EventSubscriptionSelectedEventCategoryCheckDepositCreated, EventSubscriptionSelectedEventCategoryCheckDepositUpdated, EventSubscriptionSelectedEventCategoryCheckTransferCreated, EventSubscriptionSelectedEventCategoryCheckTransferUpdated, EventSubscriptionSelectedEventCategoryDeclinedTransactionCreated, EventSubscriptionSelectedEventCategoryDigitalCardProfileCreated, EventSubscriptionSelectedEventCategoryDigitalCardProfileUpdated, EventSubscriptionSelectedEventCategoryDigitalWalletTokenCreated, EventSubscriptionSelectedEventCategoryDigitalWalletTokenUpdated, EventSubscriptionSelectedEventCategoryDocumentCreated, EventSubscriptionSelectedEventCategoryEntityCreated, EventSubscriptionSelectedEventCategoryEntityUpdated, EventSubscriptionSelectedEventCategoryEventSubscriptionCreated, EventSubscriptionSelectedEventCategoryEventSubscriptionUpdated, EventSubscriptionSelectedEventCategoryExportCreated, EventSubscriptionSelectedEventCategoryExportUpdated, EventSubscriptionSelectedEventCategoryExternalAccountCreated, EventSubscriptionSelectedEventCategoryExternalAccountUpdated, EventSubscriptionSelectedEventCategoryFileCreated, EventSubscriptionSelectedEventCategoryGroupUpdated, EventSubscriptionSelectedEventCategoryGroupHeartbeat, EventSubscriptionSelectedEventCategoryInboundACHTransferCreated, EventSubscriptionSelectedEventCategoryInboundACHTransferUpdated, EventSubscriptionSelectedEventCategoryInboundACHTransferReturnCreated, EventSubscriptionSelectedEventCategoryInboundACHTransferReturnUpdated, EventSubscriptionSelectedEventCategoryInboundCheckDepositCreated, EventSubscriptionSelectedEventCategoryInboundCheckDepositUpdated, EventSubscriptionSelectedEventCategoryInboundMailItemCreated, EventSubscriptionSelectedEventCategoryInboundMailItemUpdated, EventSubscriptionSelectedEventCategoryInboundRealTimePaymentsTransferCreated, EventSubscriptionSelectedEventCategoryInboundRealTimePaymentsTransferUpdated, EventSubscriptionSelectedEventCategoryInboundWireDrawdownRequestCreated, EventSubscriptionSelectedEventCategoryInboundWireTransferCreated, EventSubscriptionSelectedEventCategoryInboundWireTransferUpdated, EventSubscriptionSelectedEventCategoryIntrafiAccountEnrollmentCreated, EventSubscriptionSelectedEventCategoryIntrafiAccountEnrollmentUpdated, EventSubscriptionSelectedEventCategoryIntrafiExclusionCreated, EventSubscriptionSelectedEventCategoryIntrafiExclusionUpdated, EventSubscriptionSelectedEventCategoryLockboxCreated, EventSubscriptionSelectedEventCategoryLockboxUpdated, EventSubscriptionSelectedEventCategoryOAuthConnectionCreated, EventSubscriptionSelectedEventCategoryOAuthConnectionDeactivated, EventSubscriptionSelectedEventCategoryPendingTransactionCreated, EventSubscriptionSelectedEventCategoryPendingTransactionUpdated, EventSubscriptionSelectedEventCategoryPhysicalCardCreated, EventSubscriptionSelectedEventCategoryPhysicalCardUpdated, EventSubscriptionSelectedEventCategoryPhysicalCardProfileCreated, EventSubscriptionSelectedEventCategoryPhysicalCardProfileUpdated, EventSubscriptionSelectedEventCategoryProofOfAuthorizationRequestCreated, EventSubscriptionSelectedEventCategoryProofOfAuthorizationRequestUpdated, EventSubscriptionSelectedEventCategoryProofOfAuthorizationRequestSubmissionCreated, EventSubscriptionSelectedEventCategoryProofOfAuthorizationRequestSubmissionUpdated, EventSubscriptionSelectedEventCategoryRealTimeDecisionCardAuthorizationRequested, EventSubscriptionSelectedEventCategoryRealTimeDecisionDigitalWalletTokenRequested, EventSubscriptionSelectedEventCategoryRealTimeDecisionDigitalWalletAuthenticationRequested, EventSubscriptionSelectedEventCategoryRealTimeDecisionCardAuthenticationRequested, EventSubscriptionSelectedEventCategoryRealTimePaymentsTransferCreated, EventSubscriptionSelectedEventCategoryRealTimePaymentsTransferUpdated, EventSubscriptionSelectedEventCategoryRealTimePaymentsRequestForPaymentCreated, EventSubscriptionSelectedEventCategoryRealTimePaymentsRequestForPaymentUpdated, EventSubscriptionSelectedEventCategoryTransactionCreated, EventSubscriptionSelectedEventCategoryWireDrawdownRequestCreated, EventSubscriptionSelectedEventCategoryWireDrawdownRequestUpdated, EventSubscriptionSelectedEventCategoryWireTransferCreated, EventSubscriptionSelectedEventCategoryWireTransferUpdated:
		return true
	}
	return false
}

// This indicates if we'll send notifications to this subscription.
type EventSubscriptionStatus string

const (
	// The subscription is active and Events will be delivered normally.
	EventSubscriptionStatusActive EventSubscriptionStatus = "active"
	// The subscription is temporarily disabled and Events will not be delivered.
	EventSubscriptionStatusDisabled EventSubscriptionStatus = "disabled"
	// The subscription is permanently disabled and Events will not be delivered.
	EventSubscriptionStatusDeleted EventSubscriptionStatus = "deleted"
	// The subscription is temporarily disabled due to delivery errors and Events will
	// not be delivered.
	EventSubscriptionStatusRequiresAttention EventSubscriptionStatus = "requires_attention"
)

func (r EventSubscriptionStatus) IsKnown() bool {
	switch r {
	case EventSubscriptionStatusActive, EventSubscriptionStatusDisabled, EventSubscriptionStatusDeleted, EventSubscriptionStatusRequiresAttention:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `event_subscription`.
type EventSubscriptionType string

const (
	EventSubscriptionTypeEventSubscription EventSubscriptionType = "event_subscription"
)

func (r EventSubscriptionType) IsKnown() bool {
	switch r {
	case EventSubscriptionTypeEventSubscription:
		return true
	}
	return false
}

type EventSubscriptionNewParams struct {
	// The URL you'd like us to send webhooks to.
	URL param.Field[string] `json:"url,required"`
	// If specified, this subscription will only receive webhooks for Events associated
	// with the specified OAuth Connection.
	OAuthConnectionID param.Field[string] `json:"oauth_connection_id"`
	// If specified, this subscription will only receive webhooks for Events with the
	// specified `category`.
	SelectedEventCategory param.Field[EventSubscriptionNewParamsSelectedEventCategory] `json:"selected_event_category"`
	// The key that will be used to sign webhooks. If no value is passed, a random
	// string will be used as default.
	SharedSecret param.Field[string] `json:"shared_secret"`
}

func (r EventSubscriptionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If specified, this subscription will only receive webhooks for Events with the
// specified `category`.
type EventSubscriptionNewParamsSelectedEventCategory string

const (
	// Occurs whenever an Account is created.
	EventSubscriptionNewParamsSelectedEventCategoryAccountCreated EventSubscriptionNewParamsSelectedEventCategory = "account.created"
	// Occurs whenever an Account is updated.
	EventSubscriptionNewParamsSelectedEventCategoryAccountUpdated EventSubscriptionNewParamsSelectedEventCategory = "account.updated"
	// Occurs whenever an Account Number is created.
	EventSubscriptionNewParamsSelectedEventCategoryAccountNumberCreated EventSubscriptionNewParamsSelectedEventCategory = "account_number.created"
	// Occurs whenever an Account Number is updated.
	EventSubscriptionNewParamsSelectedEventCategoryAccountNumberUpdated EventSubscriptionNewParamsSelectedEventCategory = "account_number.updated"
	// Occurs whenever an Account Statement is created.
	EventSubscriptionNewParamsSelectedEventCategoryAccountStatementCreated EventSubscriptionNewParamsSelectedEventCategory = "account_statement.created"
	// Occurs whenever an Account Transfer is created.
	EventSubscriptionNewParamsSelectedEventCategoryAccountTransferCreated EventSubscriptionNewParamsSelectedEventCategory = "account_transfer.created"
	// Occurs whenever an Account Transfer is updated.
	EventSubscriptionNewParamsSelectedEventCategoryAccountTransferUpdated EventSubscriptionNewParamsSelectedEventCategory = "account_transfer.updated"
	// Occurs whenever an ACH Prenotification is created.
	EventSubscriptionNewParamsSelectedEventCategoryACHPrenotificationCreated EventSubscriptionNewParamsSelectedEventCategory = "ach_prenotification.created"
	// Occurs whenever an ACH Prenotification is updated.
	EventSubscriptionNewParamsSelectedEventCategoryACHPrenotificationUpdated EventSubscriptionNewParamsSelectedEventCategory = "ach_prenotification.updated"
	// Occurs whenever an ACH Transfer is created.
	EventSubscriptionNewParamsSelectedEventCategoryACHTransferCreated EventSubscriptionNewParamsSelectedEventCategory = "ach_transfer.created"
	// Occurs whenever an ACH Transfer is updated.
	EventSubscriptionNewParamsSelectedEventCategoryACHTransferUpdated EventSubscriptionNewParamsSelectedEventCategory = "ach_transfer.updated"
	// Occurs whenever a Bookkeeping Account is created.
	EventSubscriptionNewParamsSelectedEventCategoryBookkeepingAccountCreated EventSubscriptionNewParamsSelectedEventCategory = "bookkeeping_account.created"
	// Occurs whenever a Bookkeeping Account is updated.
	EventSubscriptionNewParamsSelectedEventCategoryBookkeepingAccountUpdated EventSubscriptionNewParamsSelectedEventCategory = "bookkeeping_account.updated"
	// Occurs whenever a Bookkeeping Entry Set is created.
	EventSubscriptionNewParamsSelectedEventCategoryBookkeepingEntrySetUpdated EventSubscriptionNewParamsSelectedEventCategory = "bookkeeping_entry_set.updated"
	// Occurs whenever a Card is created.
	EventSubscriptionNewParamsSelectedEventCategoryCardCreated EventSubscriptionNewParamsSelectedEventCategory = "card.created"
	// Occurs whenever a Card is updated.
	EventSubscriptionNewParamsSelectedEventCategoryCardUpdated EventSubscriptionNewParamsSelectedEventCategory = "card.updated"
	// Occurs whenever a Card Payment is created.
	EventSubscriptionNewParamsSelectedEventCategoryCardPaymentCreated EventSubscriptionNewParamsSelectedEventCategory = "card_payment.created"
	// Occurs whenever a Card Payment is updated.
	EventSubscriptionNewParamsSelectedEventCategoryCardPaymentUpdated EventSubscriptionNewParamsSelectedEventCategory = "card_payment.updated"
	// Occurs whenever a Card Profile is created.
	EventSubscriptionNewParamsSelectedEventCategoryCardProfileCreated EventSubscriptionNewParamsSelectedEventCategory = "card_profile.created"
	// Occurs whenever a Card Profile is updated.
	EventSubscriptionNewParamsSelectedEventCategoryCardProfileUpdated EventSubscriptionNewParamsSelectedEventCategory = "card_profile.updated"
	// Occurs whenever a Card Dispute is created.
	EventSubscriptionNewParamsSelectedEventCategoryCardDisputeCreated EventSubscriptionNewParamsSelectedEventCategory = "card_dispute.created"
	// Occurs whenever a Card Dispute is updated.
	EventSubscriptionNewParamsSelectedEventCategoryCardDisputeUpdated EventSubscriptionNewParamsSelectedEventCategory = "card_dispute.updated"
	// Occurs whenever a Check Deposit is created.
	EventSubscriptionNewParamsSelectedEventCategoryCheckDepositCreated EventSubscriptionNewParamsSelectedEventCategory = "check_deposit.created"
	// Occurs whenever a Check Deposit is updated.
	EventSubscriptionNewParamsSelectedEventCategoryCheckDepositUpdated EventSubscriptionNewParamsSelectedEventCategory = "check_deposit.updated"
	// Occurs whenever a Check Transfer is created.
	EventSubscriptionNewParamsSelectedEventCategoryCheckTransferCreated EventSubscriptionNewParamsSelectedEventCategory = "check_transfer.created"
	// Occurs whenever a Check Transfer is updated.
	EventSubscriptionNewParamsSelectedEventCategoryCheckTransferUpdated EventSubscriptionNewParamsSelectedEventCategory = "check_transfer.updated"
	// Occurs whenever a Declined Transaction is created.
	EventSubscriptionNewParamsSelectedEventCategoryDeclinedTransactionCreated EventSubscriptionNewParamsSelectedEventCategory = "declined_transaction.created"
	// Occurs whenever a Digital Card Profile is created.
	EventSubscriptionNewParamsSelectedEventCategoryDigitalCardProfileCreated EventSubscriptionNewParamsSelectedEventCategory = "digital_card_profile.created"
	// Occurs whenever a Digital Card Profile is updated.
	EventSubscriptionNewParamsSelectedEventCategoryDigitalCardProfileUpdated EventSubscriptionNewParamsSelectedEventCategory = "digital_card_profile.updated"
	// Occurs whenever a Digital Wallet Token is created.
	EventSubscriptionNewParamsSelectedEventCategoryDigitalWalletTokenCreated EventSubscriptionNewParamsSelectedEventCategory = "digital_wallet_token.created"
	// Occurs whenever a Digital Wallet Token is updated.
	EventSubscriptionNewParamsSelectedEventCategoryDigitalWalletTokenUpdated EventSubscriptionNewParamsSelectedEventCategory = "digital_wallet_token.updated"
	// Occurs whenever a Document is created.
	EventSubscriptionNewParamsSelectedEventCategoryDocumentCreated EventSubscriptionNewParamsSelectedEventCategory = "document.created"
	// Occurs whenever an Entity is created.
	EventSubscriptionNewParamsSelectedEventCategoryEntityCreated EventSubscriptionNewParamsSelectedEventCategory = "entity.created"
	// Occurs whenever an Entity is updated.
	EventSubscriptionNewParamsSelectedEventCategoryEntityUpdated EventSubscriptionNewParamsSelectedEventCategory = "entity.updated"
	// Occurs whenever an Event Subscription is created.
	EventSubscriptionNewParamsSelectedEventCategoryEventSubscriptionCreated EventSubscriptionNewParamsSelectedEventCategory = "event_subscription.created"
	// Occurs whenever an Event Subscription is updated.
	EventSubscriptionNewParamsSelectedEventCategoryEventSubscriptionUpdated EventSubscriptionNewParamsSelectedEventCategory = "event_subscription.updated"
	// Occurs whenever an Export is created.
	EventSubscriptionNewParamsSelectedEventCategoryExportCreated EventSubscriptionNewParamsSelectedEventCategory = "export.created"
	// Occurs whenever an Export is updated.
	EventSubscriptionNewParamsSelectedEventCategoryExportUpdated EventSubscriptionNewParamsSelectedEventCategory = "export.updated"
	// Occurs whenever an External Account is created.
	EventSubscriptionNewParamsSelectedEventCategoryExternalAccountCreated EventSubscriptionNewParamsSelectedEventCategory = "external_account.created"
	// Occurs whenever an External Account is updated.
	EventSubscriptionNewParamsSelectedEventCategoryExternalAccountUpdated EventSubscriptionNewParamsSelectedEventCategory = "external_account.updated"
	// Occurs whenever a File is created.
	EventSubscriptionNewParamsSelectedEventCategoryFileCreated EventSubscriptionNewParamsSelectedEventCategory = "file.created"
	// Occurs whenever a Group is updated.
	EventSubscriptionNewParamsSelectedEventCategoryGroupUpdated EventSubscriptionNewParamsSelectedEventCategory = "group.updated"
	// Increase may send webhooks with this category to see if a webhook endpoint is
	// working properly.
	EventSubscriptionNewParamsSelectedEventCategoryGroupHeartbeat EventSubscriptionNewParamsSelectedEventCategory = "group.heartbeat"
	// Occurs whenever an Inbound ACH Transfer is created.
	EventSubscriptionNewParamsSelectedEventCategoryInboundACHTransferCreated EventSubscriptionNewParamsSelectedEventCategory = "inbound_ach_transfer.created"
	// Occurs whenever an Inbound ACH Transfer is updated.
	EventSubscriptionNewParamsSelectedEventCategoryInboundACHTransferUpdated EventSubscriptionNewParamsSelectedEventCategory = "inbound_ach_transfer.updated"
	// Occurs whenever an Inbound ACH Transfer Return is created.
	EventSubscriptionNewParamsSelectedEventCategoryInboundACHTransferReturnCreated EventSubscriptionNewParamsSelectedEventCategory = "inbound_ach_transfer_return.created"
	// Occurs whenever an Inbound ACH Transfer Return is updated.
	EventSubscriptionNewParamsSelectedEventCategoryInboundACHTransferReturnUpdated EventSubscriptionNewParamsSelectedEventCategory = "inbound_ach_transfer_return.updated"
	// Occurs whenever an Inbound Check Deposit is created.
	EventSubscriptionNewParamsSelectedEventCategoryInboundCheckDepositCreated EventSubscriptionNewParamsSelectedEventCategory = "inbound_check_deposit.created"
	// Occurs whenever an Inbound Check Deposit is updated.
	EventSubscriptionNewParamsSelectedEventCategoryInboundCheckDepositUpdated EventSubscriptionNewParamsSelectedEventCategory = "inbound_check_deposit.updated"
	// Occurs whenever an Inbound Mail Item is created.
	EventSubscriptionNewParamsSelectedEventCategoryInboundMailItemCreated EventSubscriptionNewParamsSelectedEventCategory = "inbound_mail_item.created"
	// Occurs whenever an Inbound Mail Item is updated.
	EventSubscriptionNewParamsSelectedEventCategoryInboundMailItemUpdated EventSubscriptionNewParamsSelectedEventCategory = "inbound_mail_item.updated"
	// Occurs whenever an Inbound Real-Time Payments Transfer is created.
	EventSubscriptionNewParamsSelectedEventCategoryInboundRealTimePaymentsTransferCreated EventSubscriptionNewParamsSelectedEventCategory = "inbound_real_time_payments_transfer.created"
	// Occurs whenever an Inbound Real-Time Payments Transfer is updated.
	EventSubscriptionNewParamsSelectedEventCategoryInboundRealTimePaymentsTransferUpdated EventSubscriptionNewParamsSelectedEventCategory = "inbound_real_time_payments_transfer.updated"
	// Occurs whenever an Inbound Wire Drawdown Request is created.
	EventSubscriptionNewParamsSelectedEventCategoryInboundWireDrawdownRequestCreated EventSubscriptionNewParamsSelectedEventCategory = "inbound_wire_drawdown_request.created"
	// Occurs whenever an Inbound Wire Transfer is created.
	EventSubscriptionNewParamsSelectedEventCategoryInboundWireTransferCreated EventSubscriptionNewParamsSelectedEventCategory = "inbound_wire_transfer.created"
	// Occurs whenever an Inbound Wire Transfer is updated.
	EventSubscriptionNewParamsSelectedEventCategoryInboundWireTransferUpdated EventSubscriptionNewParamsSelectedEventCategory = "inbound_wire_transfer.updated"
	// Occurs whenever an IntraFi Account Enrollment is created.
	EventSubscriptionNewParamsSelectedEventCategoryIntrafiAccountEnrollmentCreated EventSubscriptionNewParamsSelectedEventCategory = "intrafi_account_enrollment.created"
	// Occurs whenever an IntraFi Account Enrollment is updated.
	EventSubscriptionNewParamsSelectedEventCategoryIntrafiAccountEnrollmentUpdated EventSubscriptionNewParamsSelectedEventCategory = "intrafi_account_enrollment.updated"
	// Occurs whenever an IntraFi Exclusion is created.
	EventSubscriptionNewParamsSelectedEventCategoryIntrafiExclusionCreated EventSubscriptionNewParamsSelectedEventCategory = "intrafi_exclusion.created"
	// Occurs whenever an IntraFi Exclusion is updated.
	EventSubscriptionNewParamsSelectedEventCategoryIntrafiExclusionUpdated EventSubscriptionNewParamsSelectedEventCategory = "intrafi_exclusion.updated"
	// Occurs whenever a Lockbox is created.
	EventSubscriptionNewParamsSelectedEventCategoryLockboxCreated EventSubscriptionNewParamsSelectedEventCategory = "lockbox.created"
	// Occurs whenever a Lockbox is updated.
	EventSubscriptionNewParamsSelectedEventCategoryLockboxUpdated EventSubscriptionNewParamsSelectedEventCategory = "lockbox.updated"
	// Occurs whenever an OAuth Connection is created.
	EventSubscriptionNewParamsSelectedEventCategoryOAuthConnectionCreated EventSubscriptionNewParamsSelectedEventCategory = "oauth_connection.created"
	// Occurs whenever an OAuth Connection is deactivated.
	EventSubscriptionNewParamsSelectedEventCategoryOAuthConnectionDeactivated EventSubscriptionNewParamsSelectedEventCategory = "oauth_connection.deactivated"
	// Occurs whenever a Pending Transaction is created.
	EventSubscriptionNewParamsSelectedEventCategoryPendingTransactionCreated EventSubscriptionNewParamsSelectedEventCategory = "pending_transaction.created"
	// Occurs whenever a Pending Transaction is updated.
	EventSubscriptionNewParamsSelectedEventCategoryPendingTransactionUpdated EventSubscriptionNewParamsSelectedEventCategory = "pending_transaction.updated"
	// Occurs whenever a Physical Card is created.
	EventSubscriptionNewParamsSelectedEventCategoryPhysicalCardCreated EventSubscriptionNewParamsSelectedEventCategory = "physical_card.created"
	// Occurs whenever a Physical Card is updated.
	EventSubscriptionNewParamsSelectedEventCategoryPhysicalCardUpdated EventSubscriptionNewParamsSelectedEventCategory = "physical_card.updated"
	// Occurs whenever a Physical Card Profile is created.
	EventSubscriptionNewParamsSelectedEventCategoryPhysicalCardProfileCreated EventSubscriptionNewParamsSelectedEventCategory = "physical_card_profile.created"
	// Occurs whenever a Physical Card Profile is updated.
	EventSubscriptionNewParamsSelectedEventCategoryPhysicalCardProfileUpdated EventSubscriptionNewParamsSelectedEventCategory = "physical_card_profile.updated"
	// Occurs whenever a Proof of Authorization Request is created.
	EventSubscriptionNewParamsSelectedEventCategoryProofOfAuthorizationRequestCreated EventSubscriptionNewParamsSelectedEventCategory = "proof_of_authorization_request.created"
	// Occurs whenever a Proof of Authorization Request is updated.
	EventSubscriptionNewParamsSelectedEventCategoryProofOfAuthorizationRequestUpdated EventSubscriptionNewParamsSelectedEventCategory = "proof_of_authorization_request.updated"
	// Occurs whenever a Proof of Authorization Request Submission is created.
	EventSubscriptionNewParamsSelectedEventCategoryProofOfAuthorizationRequestSubmissionCreated EventSubscriptionNewParamsSelectedEventCategory = "proof_of_authorization_request_submission.created"
	// Occurs whenever a Proof of Authorization Request Submission is updated.
	EventSubscriptionNewParamsSelectedEventCategoryProofOfAuthorizationRequestSubmissionUpdated EventSubscriptionNewParamsSelectedEventCategory = "proof_of_authorization_request_submission.updated"
	// Occurs whenever a Real-Time Decision is created in response to a card
	// authorization.
	EventSubscriptionNewParamsSelectedEventCategoryRealTimeDecisionCardAuthorizationRequested EventSubscriptionNewParamsSelectedEventCategory = "real_time_decision.card_authorization_requested"
	// Occurs whenever a Real-Time Decision is created in response to a digital wallet
	// provisioning attempt.
	EventSubscriptionNewParamsSelectedEventCategoryRealTimeDecisionDigitalWalletTokenRequested EventSubscriptionNewParamsSelectedEventCategory = "real_time_decision.digital_wallet_token_requested"
	// Occurs whenever a Real-Time Decision is created in response to a digital wallet
	// requiring two-factor authentication.
	EventSubscriptionNewParamsSelectedEventCategoryRealTimeDecisionDigitalWalletAuthenticationRequested EventSubscriptionNewParamsSelectedEventCategory = "real_time_decision.digital_wallet_authentication_requested"
	// Occurs whenever a Real-Time Decision is created in response to 3DS
	// authentication.
	EventSubscriptionNewParamsSelectedEventCategoryRealTimeDecisionCardAuthenticationRequested EventSubscriptionNewParamsSelectedEventCategory = "real_time_decision.card_authentication_requested"
	// Occurs whenever a Real-Time Payments Transfer is created.
	EventSubscriptionNewParamsSelectedEventCategoryRealTimePaymentsTransferCreated EventSubscriptionNewParamsSelectedEventCategory = "real_time_payments_transfer.created"
	// Occurs whenever a Real-Time Payments Transfer is updated.
	EventSubscriptionNewParamsSelectedEventCategoryRealTimePaymentsTransferUpdated EventSubscriptionNewParamsSelectedEventCategory = "real_time_payments_transfer.updated"
	// Occurs whenever a Real-Time Payments Request for Payment is created.
	EventSubscriptionNewParamsSelectedEventCategoryRealTimePaymentsRequestForPaymentCreated EventSubscriptionNewParamsSelectedEventCategory = "real_time_payments_request_for_payment.created"
	// Occurs whenever a Real-Time Payments Request for Payment is updated.
	EventSubscriptionNewParamsSelectedEventCategoryRealTimePaymentsRequestForPaymentUpdated EventSubscriptionNewParamsSelectedEventCategory = "real_time_payments_request_for_payment.updated"
	// Occurs whenever a Transaction is created.
	EventSubscriptionNewParamsSelectedEventCategoryTransactionCreated EventSubscriptionNewParamsSelectedEventCategory = "transaction.created"
	// Occurs whenever a Wire Drawdown Request is created.
	EventSubscriptionNewParamsSelectedEventCategoryWireDrawdownRequestCreated EventSubscriptionNewParamsSelectedEventCategory = "wire_drawdown_request.created"
	// Occurs whenever a Wire Drawdown Request is updated.
	EventSubscriptionNewParamsSelectedEventCategoryWireDrawdownRequestUpdated EventSubscriptionNewParamsSelectedEventCategory = "wire_drawdown_request.updated"
	// Occurs whenever a Wire Transfer is created.
	EventSubscriptionNewParamsSelectedEventCategoryWireTransferCreated EventSubscriptionNewParamsSelectedEventCategory = "wire_transfer.created"
	// Occurs whenever a Wire Transfer is updated.
	EventSubscriptionNewParamsSelectedEventCategoryWireTransferUpdated EventSubscriptionNewParamsSelectedEventCategory = "wire_transfer.updated"
)

func (r EventSubscriptionNewParamsSelectedEventCategory) IsKnown() bool {
	switch r {
	case EventSubscriptionNewParamsSelectedEventCategoryAccountCreated, EventSubscriptionNewParamsSelectedEventCategoryAccountUpdated, EventSubscriptionNewParamsSelectedEventCategoryAccountNumberCreated, EventSubscriptionNewParamsSelectedEventCategoryAccountNumberUpdated, EventSubscriptionNewParamsSelectedEventCategoryAccountStatementCreated, EventSubscriptionNewParamsSelectedEventCategoryAccountTransferCreated, EventSubscriptionNewParamsSelectedEventCategoryAccountTransferUpdated, EventSubscriptionNewParamsSelectedEventCategoryACHPrenotificationCreated, EventSubscriptionNewParamsSelectedEventCategoryACHPrenotificationUpdated, EventSubscriptionNewParamsSelectedEventCategoryACHTransferCreated, EventSubscriptionNewParamsSelectedEventCategoryACHTransferUpdated, EventSubscriptionNewParamsSelectedEventCategoryBookkeepingAccountCreated, EventSubscriptionNewParamsSelectedEventCategoryBookkeepingAccountUpdated, EventSubscriptionNewParamsSelectedEventCategoryBookkeepingEntrySetUpdated, EventSubscriptionNewParamsSelectedEventCategoryCardCreated, EventSubscriptionNewParamsSelectedEventCategoryCardUpdated, EventSubscriptionNewParamsSelectedEventCategoryCardPaymentCreated, EventSubscriptionNewParamsSelectedEventCategoryCardPaymentUpdated, EventSubscriptionNewParamsSelectedEventCategoryCardProfileCreated, EventSubscriptionNewParamsSelectedEventCategoryCardProfileUpdated, EventSubscriptionNewParamsSelectedEventCategoryCardDisputeCreated, EventSubscriptionNewParamsSelectedEventCategoryCardDisputeUpdated, EventSubscriptionNewParamsSelectedEventCategoryCheckDepositCreated, EventSubscriptionNewParamsSelectedEventCategoryCheckDepositUpdated, EventSubscriptionNewParamsSelectedEventCategoryCheckTransferCreated, EventSubscriptionNewParamsSelectedEventCategoryCheckTransferUpdated, EventSubscriptionNewParamsSelectedEventCategoryDeclinedTransactionCreated, EventSubscriptionNewParamsSelectedEventCategoryDigitalCardProfileCreated, EventSubscriptionNewParamsSelectedEventCategoryDigitalCardProfileUpdated, EventSubscriptionNewParamsSelectedEventCategoryDigitalWalletTokenCreated, EventSubscriptionNewParamsSelectedEventCategoryDigitalWalletTokenUpdated, EventSubscriptionNewParamsSelectedEventCategoryDocumentCreated, EventSubscriptionNewParamsSelectedEventCategoryEntityCreated, EventSubscriptionNewParamsSelectedEventCategoryEntityUpdated, EventSubscriptionNewParamsSelectedEventCategoryEventSubscriptionCreated, EventSubscriptionNewParamsSelectedEventCategoryEventSubscriptionUpdated, EventSubscriptionNewParamsSelectedEventCategoryExportCreated, EventSubscriptionNewParamsSelectedEventCategoryExportUpdated, EventSubscriptionNewParamsSelectedEventCategoryExternalAccountCreated, EventSubscriptionNewParamsSelectedEventCategoryExternalAccountUpdated, EventSubscriptionNewParamsSelectedEventCategoryFileCreated, EventSubscriptionNewParamsSelectedEventCategoryGroupUpdated, EventSubscriptionNewParamsSelectedEventCategoryGroupHeartbeat, EventSubscriptionNewParamsSelectedEventCategoryInboundACHTransferCreated, EventSubscriptionNewParamsSelectedEventCategoryInboundACHTransferUpdated, EventSubscriptionNewParamsSelectedEventCategoryInboundACHTransferReturnCreated, EventSubscriptionNewParamsSelectedEventCategoryInboundACHTransferReturnUpdated, EventSubscriptionNewParamsSelectedEventCategoryInboundCheckDepositCreated, EventSubscriptionNewParamsSelectedEventCategoryInboundCheckDepositUpdated, EventSubscriptionNewParamsSelectedEventCategoryInboundMailItemCreated, EventSubscriptionNewParamsSelectedEventCategoryInboundMailItemUpdated, EventSubscriptionNewParamsSelectedEventCategoryInboundRealTimePaymentsTransferCreated, EventSubscriptionNewParamsSelectedEventCategoryInboundRealTimePaymentsTransferUpdated, EventSubscriptionNewParamsSelectedEventCategoryInboundWireDrawdownRequestCreated, EventSubscriptionNewParamsSelectedEventCategoryInboundWireTransferCreated, EventSubscriptionNewParamsSelectedEventCategoryInboundWireTransferUpdated, EventSubscriptionNewParamsSelectedEventCategoryIntrafiAccountEnrollmentCreated, EventSubscriptionNewParamsSelectedEventCategoryIntrafiAccountEnrollmentUpdated, EventSubscriptionNewParamsSelectedEventCategoryIntrafiExclusionCreated, EventSubscriptionNewParamsSelectedEventCategoryIntrafiExclusionUpdated, EventSubscriptionNewParamsSelectedEventCategoryLockboxCreated, EventSubscriptionNewParamsSelectedEventCategoryLockboxUpdated, EventSubscriptionNewParamsSelectedEventCategoryOAuthConnectionCreated, EventSubscriptionNewParamsSelectedEventCategoryOAuthConnectionDeactivated, EventSubscriptionNewParamsSelectedEventCategoryPendingTransactionCreated, EventSubscriptionNewParamsSelectedEventCategoryPendingTransactionUpdated, EventSubscriptionNewParamsSelectedEventCategoryPhysicalCardCreated, EventSubscriptionNewParamsSelectedEventCategoryPhysicalCardUpdated, EventSubscriptionNewParamsSelectedEventCategoryPhysicalCardProfileCreated, EventSubscriptionNewParamsSelectedEventCategoryPhysicalCardProfileUpdated, EventSubscriptionNewParamsSelectedEventCategoryProofOfAuthorizationRequestCreated, EventSubscriptionNewParamsSelectedEventCategoryProofOfAuthorizationRequestUpdated, EventSubscriptionNewParamsSelectedEventCategoryProofOfAuthorizationRequestSubmissionCreated, EventSubscriptionNewParamsSelectedEventCategoryProofOfAuthorizationRequestSubmissionUpdated, EventSubscriptionNewParamsSelectedEventCategoryRealTimeDecisionCardAuthorizationRequested, EventSubscriptionNewParamsSelectedEventCategoryRealTimeDecisionDigitalWalletTokenRequested, EventSubscriptionNewParamsSelectedEventCategoryRealTimeDecisionDigitalWalletAuthenticationRequested, EventSubscriptionNewParamsSelectedEventCategoryRealTimeDecisionCardAuthenticationRequested, EventSubscriptionNewParamsSelectedEventCategoryRealTimePaymentsTransferCreated, EventSubscriptionNewParamsSelectedEventCategoryRealTimePaymentsTransferUpdated, EventSubscriptionNewParamsSelectedEventCategoryRealTimePaymentsRequestForPaymentCreated, EventSubscriptionNewParamsSelectedEventCategoryRealTimePaymentsRequestForPaymentUpdated, EventSubscriptionNewParamsSelectedEventCategoryTransactionCreated, EventSubscriptionNewParamsSelectedEventCategoryWireDrawdownRequestCreated, EventSubscriptionNewParamsSelectedEventCategoryWireDrawdownRequestUpdated, EventSubscriptionNewParamsSelectedEventCategoryWireTransferCreated, EventSubscriptionNewParamsSelectedEventCategoryWireTransferUpdated:
		return true
	}
	return false
}

type EventSubscriptionUpdateParams struct {
	// The status to update the Event Subscription with.
	Status param.Field[EventSubscriptionUpdateParamsStatus] `json:"status"`
}

func (r EventSubscriptionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The status to update the Event Subscription with.
type EventSubscriptionUpdateParamsStatus string

const (
	// The subscription is active and Events will be delivered normally.
	EventSubscriptionUpdateParamsStatusActive EventSubscriptionUpdateParamsStatus = "active"
	// The subscription is temporarily disabled and Events will not be delivered.
	EventSubscriptionUpdateParamsStatusDisabled EventSubscriptionUpdateParamsStatus = "disabled"
	// The subscription is permanently disabled and Events will not be delivered.
	EventSubscriptionUpdateParamsStatusDeleted EventSubscriptionUpdateParamsStatus = "deleted"
)

func (r EventSubscriptionUpdateParamsStatus) IsKnown() bool {
	switch r {
	case EventSubscriptionUpdateParamsStatusActive, EventSubscriptionUpdateParamsStatusDisabled, EventSubscriptionUpdateParamsStatusDeleted:
		return true
	}
	return false
}

type EventSubscriptionListParams struct {
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Filter records to the one with the specified `idempotency_key` you chose for
	// that object. This value is unique across Increase and is used to ensure that a
	// request is only processed once. Learn more about
	// [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey param.Field[string] `query:"idempotency_key"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [EventSubscriptionListParams]'s query parameters as
// `url.Values`.
func (r EventSubscriptionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

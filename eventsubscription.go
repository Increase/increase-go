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
	opts = slices.Concat(r.Options, opts)
	path := "event_subscriptions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve an Event Subscription
func (r *EventSubscriptionService) Get(ctx context.Context, eventSubscriptionID string, opts ...option.RequestOption) (res *EventSubscription, err error) {
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
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
	EventSubscriptionSelectedEventCategoryAccountCreated                                       EventSubscriptionSelectedEventCategory = "account.created"
	EventSubscriptionSelectedEventCategoryAccountUpdated                                       EventSubscriptionSelectedEventCategory = "account.updated"
	EventSubscriptionSelectedEventCategoryAccountNumberCreated                                 EventSubscriptionSelectedEventCategory = "account_number.created"
	EventSubscriptionSelectedEventCategoryAccountNumberUpdated                                 EventSubscriptionSelectedEventCategory = "account_number.updated"
	EventSubscriptionSelectedEventCategoryAccountStatementCreated                              EventSubscriptionSelectedEventCategory = "account_statement.created"
	EventSubscriptionSelectedEventCategoryAccountTransferCreated                               EventSubscriptionSelectedEventCategory = "account_transfer.created"
	EventSubscriptionSelectedEventCategoryAccountTransferUpdated                               EventSubscriptionSelectedEventCategory = "account_transfer.updated"
	EventSubscriptionSelectedEventCategoryACHPrenotificationCreated                            EventSubscriptionSelectedEventCategory = "ach_prenotification.created"
	EventSubscriptionSelectedEventCategoryACHPrenotificationUpdated                            EventSubscriptionSelectedEventCategory = "ach_prenotification.updated"
	EventSubscriptionSelectedEventCategoryACHTransferCreated                                   EventSubscriptionSelectedEventCategory = "ach_transfer.created"
	EventSubscriptionSelectedEventCategoryACHTransferUpdated                                   EventSubscriptionSelectedEventCategory = "ach_transfer.updated"
	EventSubscriptionSelectedEventCategoryBookkeepingAccountCreated                            EventSubscriptionSelectedEventCategory = "bookkeeping_account.created"
	EventSubscriptionSelectedEventCategoryBookkeepingAccountUpdated                            EventSubscriptionSelectedEventCategory = "bookkeeping_account.updated"
	EventSubscriptionSelectedEventCategoryBookkeepingEntrySetUpdated                           EventSubscriptionSelectedEventCategory = "bookkeeping_entry_set.updated"
	EventSubscriptionSelectedEventCategoryCardCreated                                          EventSubscriptionSelectedEventCategory = "card.created"
	EventSubscriptionSelectedEventCategoryCardUpdated                                          EventSubscriptionSelectedEventCategory = "card.updated"
	EventSubscriptionSelectedEventCategoryCardPaymentCreated                                   EventSubscriptionSelectedEventCategory = "card_payment.created"
	EventSubscriptionSelectedEventCategoryCardPaymentUpdated                                   EventSubscriptionSelectedEventCategory = "card_payment.updated"
	EventSubscriptionSelectedEventCategoryCardProfileCreated                                   EventSubscriptionSelectedEventCategory = "card_profile.created"
	EventSubscriptionSelectedEventCategoryCardProfileUpdated                                   EventSubscriptionSelectedEventCategory = "card_profile.updated"
	EventSubscriptionSelectedEventCategoryCardDisputeCreated                                   EventSubscriptionSelectedEventCategory = "card_dispute.created"
	EventSubscriptionSelectedEventCategoryCardDisputeUpdated                                   EventSubscriptionSelectedEventCategory = "card_dispute.updated"
	EventSubscriptionSelectedEventCategoryCheckDepositCreated                                  EventSubscriptionSelectedEventCategory = "check_deposit.created"
	EventSubscriptionSelectedEventCategoryCheckDepositUpdated                                  EventSubscriptionSelectedEventCategory = "check_deposit.updated"
	EventSubscriptionSelectedEventCategoryCheckTransferCreated                                 EventSubscriptionSelectedEventCategory = "check_transfer.created"
	EventSubscriptionSelectedEventCategoryCheckTransferUpdated                                 EventSubscriptionSelectedEventCategory = "check_transfer.updated"
	EventSubscriptionSelectedEventCategoryDeclinedTransactionCreated                           EventSubscriptionSelectedEventCategory = "declined_transaction.created"
	EventSubscriptionSelectedEventCategoryDigitalCardProfileCreated                            EventSubscriptionSelectedEventCategory = "digital_card_profile.created"
	EventSubscriptionSelectedEventCategoryDigitalCardProfileUpdated                            EventSubscriptionSelectedEventCategory = "digital_card_profile.updated"
	EventSubscriptionSelectedEventCategoryDigitalWalletTokenCreated                            EventSubscriptionSelectedEventCategory = "digital_wallet_token.created"
	EventSubscriptionSelectedEventCategoryDigitalWalletTokenUpdated                            EventSubscriptionSelectedEventCategory = "digital_wallet_token.updated"
	EventSubscriptionSelectedEventCategoryDocumentCreated                                      EventSubscriptionSelectedEventCategory = "document.created"
	EventSubscriptionSelectedEventCategoryEntityCreated                                        EventSubscriptionSelectedEventCategory = "entity.created"
	EventSubscriptionSelectedEventCategoryEntityUpdated                                        EventSubscriptionSelectedEventCategory = "entity.updated"
	EventSubscriptionSelectedEventCategoryEventSubscriptionCreated                             EventSubscriptionSelectedEventCategory = "event_subscription.created"
	EventSubscriptionSelectedEventCategoryEventSubscriptionUpdated                             EventSubscriptionSelectedEventCategory = "event_subscription.updated"
	EventSubscriptionSelectedEventCategoryExportCreated                                        EventSubscriptionSelectedEventCategory = "export.created"
	EventSubscriptionSelectedEventCategoryExportUpdated                                        EventSubscriptionSelectedEventCategory = "export.updated"
	EventSubscriptionSelectedEventCategoryExternalAccountCreated                               EventSubscriptionSelectedEventCategory = "external_account.created"
	EventSubscriptionSelectedEventCategoryExternalAccountUpdated                               EventSubscriptionSelectedEventCategory = "external_account.updated"
	EventSubscriptionSelectedEventCategoryFednowTransferCreated                                EventSubscriptionSelectedEventCategory = "fednow_transfer.created"
	EventSubscriptionSelectedEventCategoryFednowTransferUpdated                                EventSubscriptionSelectedEventCategory = "fednow_transfer.updated"
	EventSubscriptionSelectedEventCategoryFileCreated                                          EventSubscriptionSelectedEventCategory = "file.created"
	EventSubscriptionSelectedEventCategoryGroupUpdated                                         EventSubscriptionSelectedEventCategory = "group.updated"
	EventSubscriptionSelectedEventCategoryGroupHeartbeat                                       EventSubscriptionSelectedEventCategory = "group.heartbeat"
	EventSubscriptionSelectedEventCategoryInboundACHTransferCreated                            EventSubscriptionSelectedEventCategory = "inbound_ach_transfer.created"
	EventSubscriptionSelectedEventCategoryInboundACHTransferUpdated                            EventSubscriptionSelectedEventCategory = "inbound_ach_transfer.updated"
	EventSubscriptionSelectedEventCategoryInboundACHTransferReturnCreated                      EventSubscriptionSelectedEventCategory = "inbound_ach_transfer_return.created"
	EventSubscriptionSelectedEventCategoryInboundACHTransferReturnUpdated                      EventSubscriptionSelectedEventCategory = "inbound_ach_transfer_return.updated"
	EventSubscriptionSelectedEventCategoryInboundCheckDepositCreated                           EventSubscriptionSelectedEventCategory = "inbound_check_deposit.created"
	EventSubscriptionSelectedEventCategoryInboundCheckDepositUpdated                           EventSubscriptionSelectedEventCategory = "inbound_check_deposit.updated"
	EventSubscriptionSelectedEventCategoryInboundFednowTransferCreated                         EventSubscriptionSelectedEventCategory = "inbound_fednow_transfer.created"
	EventSubscriptionSelectedEventCategoryInboundFednowTransferUpdated                         EventSubscriptionSelectedEventCategory = "inbound_fednow_transfer.updated"
	EventSubscriptionSelectedEventCategoryInboundMailItemCreated                               EventSubscriptionSelectedEventCategory = "inbound_mail_item.created"
	EventSubscriptionSelectedEventCategoryInboundMailItemUpdated                               EventSubscriptionSelectedEventCategory = "inbound_mail_item.updated"
	EventSubscriptionSelectedEventCategoryInboundRealTimePaymentsTransferCreated               EventSubscriptionSelectedEventCategory = "inbound_real_time_payments_transfer.created"
	EventSubscriptionSelectedEventCategoryInboundRealTimePaymentsTransferUpdated               EventSubscriptionSelectedEventCategory = "inbound_real_time_payments_transfer.updated"
	EventSubscriptionSelectedEventCategoryInboundWireDrawdownRequestCreated                    EventSubscriptionSelectedEventCategory = "inbound_wire_drawdown_request.created"
	EventSubscriptionSelectedEventCategoryInboundWireTransferCreated                           EventSubscriptionSelectedEventCategory = "inbound_wire_transfer.created"
	EventSubscriptionSelectedEventCategoryInboundWireTransferUpdated                           EventSubscriptionSelectedEventCategory = "inbound_wire_transfer.updated"
	EventSubscriptionSelectedEventCategoryIntrafiAccountEnrollmentCreated                      EventSubscriptionSelectedEventCategory = "intrafi_account_enrollment.created"
	EventSubscriptionSelectedEventCategoryIntrafiAccountEnrollmentUpdated                      EventSubscriptionSelectedEventCategory = "intrafi_account_enrollment.updated"
	EventSubscriptionSelectedEventCategoryIntrafiExclusionCreated                              EventSubscriptionSelectedEventCategory = "intrafi_exclusion.created"
	EventSubscriptionSelectedEventCategoryIntrafiExclusionUpdated                              EventSubscriptionSelectedEventCategory = "intrafi_exclusion.updated"
	EventSubscriptionSelectedEventCategoryLegacyCardDisputeCreated                             EventSubscriptionSelectedEventCategory = "legacy_card_dispute.created"
	EventSubscriptionSelectedEventCategoryLegacyCardDisputeUpdated                             EventSubscriptionSelectedEventCategory = "legacy_card_dispute.updated"
	EventSubscriptionSelectedEventCategoryLockboxCreated                                       EventSubscriptionSelectedEventCategory = "lockbox.created"
	EventSubscriptionSelectedEventCategoryLockboxUpdated                                       EventSubscriptionSelectedEventCategory = "lockbox.updated"
	EventSubscriptionSelectedEventCategoryOAuthConnectionCreated                               EventSubscriptionSelectedEventCategory = "oauth_connection.created"
	EventSubscriptionSelectedEventCategoryOAuthConnectionDeactivated                           EventSubscriptionSelectedEventCategory = "oauth_connection.deactivated"
	EventSubscriptionSelectedEventCategoryOutboundCardPushTransferCreated                      EventSubscriptionSelectedEventCategory = "outbound_card_push_transfer.created"
	EventSubscriptionSelectedEventCategoryOutboundCardPushTransferUpdated                      EventSubscriptionSelectedEventCategory = "outbound_card_push_transfer.updated"
	EventSubscriptionSelectedEventCategoryOutboundCardValidationCreated                        EventSubscriptionSelectedEventCategory = "outbound_card_validation.created"
	EventSubscriptionSelectedEventCategoryOutboundCardValidationUpdated                        EventSubscriptionSelectedEventCategory = "outbound_card_validation.updated"
	EventSubscriptionSelectedEventCategoryCardPushTransferCreated                              EventSubscriptionSelectedEventCategory = "card_push_transfer.created"
	EventSubscriptionSelectedEventCategoryCardPushTransferUpdated                              EventSubscriptionSelectedEventCategory = "card_push_transfer.updated"
	EventSubscriptionSelectedEventCategoryCardValidationCreated                                EventSubscriptionSelectedEventCategory = "card_validation.created"
	EventSubscriptionSelectedEventCategoryCardValidationUpdated                                EventSubscriptionSelectedEventCategory = "card_validation.updated"
	EventSubscriptionSelectedEventCategoryPendingTransactionCreated                            EventSubscriptionSelectedEventCategory = "pending_transaction.created"
	EventSubscriptionSelectedEventCategoryPendingTransactionUpdated                            EventSubscriptionSelectedEventCategory = "pending_transaction.updated"
	EventSubscriptionSelectedEventCategoryPhysicalCardCreated                                  EventSubscriptionSelectedEventCategory = "physical_card.created"
	EventSubscriptionSelectedEventCategoryPhysicalCardUpdated                                  EventSubscriptionSelectedEventCategory = "physical_card.updated"
	EventSubscriptionSelectedEventCategoryPhysicalCardProfileCreated                           EventSubscriptionSelectedEventCategory = "physical_card_profile.created"
	EventSubscriptionSelectedEventCategoryPhysicalCardProfileUpdated                           EventSubscriptionSelectedEventCategory = "physical_card_profile.updated"
	EventSubscriptionSelectedEventCategoryProgramCreated                                       EventSubscriptionSelectedEventCategory = "program.created"
	EventSubscriptionSelectedEventCategoryProgramUpdated                                       EventSubscriptionSelectedEventCategory = "program.updated"
	EventSubscriptionSelectedEventCategoryProofOfAuthorizationRequestCreated                   EventSubscriptionSelectedEventCategory = "proof_of_authorization_request.created"
	EventSubscriptionSelectedEventCategoryProofOfAuthorizationRequestUpdated                   EventSubscriptionSelectedEventCategory = "proof_of_authorization_request.updated"
	EventSubscriptionSelectedEventCategoryRealTimeDecisionCardAuthorizationRequested           EventSubscriptionSelectedEventCategory = "real_time_decision.card_authorization_requested"
	EventSubscriptionSelectedEventCategoryRealTimeDecisionDigitalWalletTokenRequested          EventSubscriptionSelectedEventCategory = "real_time_decision.digital_wallet_token_requested"
	EventSubscriptionSelectedEventCategoryRealTimeDecisionDigitalWalletAuthenticationRequested EventSubscriptionSelectedEventCategory = "real_time_decision.digital_wallet_authentication_requested"
	EventSubscriptionSelectedEventCategoryRealTimeDecisionCardAuthenticationRequested          EventSubscriptionSelectedEventCategory = "real_time_decision.card_authentication_requested"
	EventSubscriptionSelectedEventCategoryRealTimeDecisionCardAuthenticationChallengeRequested EventSubscriptionSelectedEventCategory = "real_time_decision.card_authentication_challenge_requested"
	EventSubscriptionSelectedEventCategoryRealTimePaymentsTransferCreated                      EventSubscriptionSelectedEventCategory = "real_time_payments_transfer.created"
	EventSubscriptionSelectedEventCategoryRealTimePaymentsTransferUpdated                      EventSubscriptionSelectedEventCategory = "real_time_payments_transfer.updated"
	EventSubscriptionSelectedEventCategoryRealTimePaymentsRequestForPaymentCreated             EventSubscriptionSelectedEventCategory = "real_time_payments_request_for_payment.created"
	EventSubscriptionSelectedEventCategoryRealTimePaymentsRequestForPaymentUpdated             EventSubscriptionSelectedEventCategory = "real_time_payments_request_for_payment.updated"
	EventSubscriptionSelectedEventCategorySwiftTransferCreated                                 EventSubscriptionSelectedEventCategory = "swift_transfer.created"
	EventSubscriptionSelectedEventCategorySwiftTransferUpdated                                 EventSubscriptionSelectedEventCategory = "swift_transfer.updated"
	EventSubscriptionSelectedEventCategoryTransactionCreated                                   EventSubscriptionSelectedEventCategory = "transaction.created"
	EventSubscriptionSelectedEventCategoryWireDrawdownRequestCreated                           EventSubscriptionSelectedEventCategory = "wire_drawdown_request.created"
	EventSubscriptionSelectedEventCategoryWireDrawdownRequestUpdated                           EventSubscriptionSelectedEventCategory = "wire_drawdown_request.updated"
	EventSubscriptionSelectedEventCategoryWireTransferCreated                                  EventSubscriptionSelectedEventCategory = "wire_transfer.created"
	EventSubscriptionSelectedEventCategoryWireTransferUpdated                                  EventSubscriptionSelectedEventCategory = "wire_transfer.updated"
)

func (r EventSubscriptionSelectedEventCategory) IsKnown() bool {
	switch r {
	case EventSubscriptionSelectedEventCategoryAccountCreated, EventSubscriptionSelectedEventCategoryAccountUpdated, EventSubscriptionSelectedEventCategoryAccountNumberCreated, EventSubscriptionSelectedEventCategoryAccountNumberUpdated, EventSubscriptionSelectedEventCategoryAccountStatementCreated, EventSubscriptionSelectedEventCategoryAccountTransferCreated, EventSubscriptionSelectedEventCategoryAccountTransferUpdated, EventSubscriptionSelectedEventCategoryACHPrenotificationCreated, EventSubscriptionSelectedEventCategoryACHPrenotificationUpdated, EventSubscriptionSelectedEventCategoryACHTransferCreated, EventSubscriptionSelectedEventCategoryACHTransferUpdated, EventSubscriptionSelectedEventCategoryBookkeepingAccountCreated, EventSubscriptionSelectedEventCategoryBookkeepingAccountUpdated, EventSubscriptionSelectedEventCategoryBookkeepingEntrySetUpdated, EventSubscriptionSelectedEventCategoryCardCreated, EventSubscriptionSelectedEventCategoryCardUpdated, EventSubscriptionSelectedEventCategoryCardPaymentCreated, EventSubscriptionSelectedEventCategoryCardPaymentUpdated, EventSubscriptionSelectedEventCategoryCardProfileCreated, EventSubscriptionSelectedEventCategoryCardProfileUpdated, EventSubscriptionSelectedEventCategoryCardDisputeCreated, EventSubscriptionSelectedEventCategoryCardDisputeUpdated, EventSubscriptionSelectedEventCategoryCheckDepositCreated, EventSubscriptionSelectedEventCategoryCheckDepositUpdated, EventSubscriptionSelectedEventCategoryCheckTransferCreated, EventSubscriptionSelectedEventCategoryCheckTransferUpdated, EventSubscriptionSelectedEventCategoryDeclinedTransactionCreated, EventSubscriptionSelectedEventCategoryDigitalCardProfileCreated, EventSubscriptionSelectedEventCategoryDigitalCardProfileUpdated, EventSubscriptionSelectedEventCategoryDigitalWalletTokenCreated, EventSubscriptionSelectedEventCategoryDigitalWalletTokenUpdated, EventSubscriptionSelectedEventCategoryDocumentCreated, EventSubscriptionSelectedEventCategoryEntityCreated, EventSubscriptionSelectedEventCategoryEntityUpdated, EventSubscriptionSelectedEventCategoryEventSubscriptionCreated, EventSubscriptionSelectedEventCategoryEventSubscriptionUpdated, EventSubscriptionSelectedEventCategoryExportCreated, EventSubscriptionSelectedEventCategoryExportUpdated, EventSubscriptionSelectedEventCategoryExternalAccountCreated, EventSubscriptionSelectedEventCategoryExternalAccountUpdated, EventSubscriptionSelectedEventCategoryFednowTransferCreated, EventSubscriptionSelectedEventCategoryFednowTransferUpdated, EventSubscriptionSelectedEventCategoryFileCreated, EventSubscriptionSelectedEventCategoryGroupUpdated, EventSubscriptionSelectedEventCategoryGroupHeartbeat, EventSubscriptionSelectedEventCategoryInboundACHTransferCreated, EventSubscriptionSelectedEventCategoryInboundACHTransferUpdated, EventSubscriptionSelectedEventCategoryInboundACHTransferReturnCreated, EventSubscriptionSelectedEventCategoryInboundACHTransferReturnUpdated, EventSubscriptionSelectedEventCategoryInboundCheckDepositCreated, EventSubscriptionSelectedEventCategoryInboundCheckDepositUpdated, EventSubscriptionSelectedEventCategoryInboundFednowTransferCreated, EventSubscriptionSelectedEventCategoryInboundFednowTransferUpdated, EventSubscriptionSelectedEventCategoryInboundMailItemCreated, EventSubscriptionSelectedEventCategoryInboundMailItemUpdated, EventSubscriptionSelectedEventCategoryInboundRealTimePaymentsTransferCreated, EventSubscriptionSelectedEventCategoryInboundRealTimePaymentsTransferUpdated, EventSubscriptionSelectedEventCategoryInboundWireDrawdownRequestCreated, EventSubscriptionSelectedEventCategoryInboundWireTransferCreated, EventSubscriptionSelectedEventCategoryInboundWireTransferUpdated, EventSubscriptionSelectedEventCategoryIntrafiAccountEnrollmentCreated, EventSubscriptionSelectedEventCategoryIntrafiAccountEnrollmentUpdated, EventSubscriptionSelectedEventCategoryIntrafiExclusionCreated, EventSubscriptionSelectedEventCategoryIntrafiExclusionUpdated, EventSubscriptionSelectedEventCategoryLegacyCardDisputeCreated, EventSubscriptionSelectedEventCategoryLegacyCardDisputeUpdated, EventSubscriptionSelectedEventCategoryLockboxCreated, EventSubscriptionSelectedEventCategoryLockboxUpdated, EventSubscriptionSelectedEventCategoryOAuthConnectionCreated, EventSubscriptionSelectedEventCategoryOAuthConnectionDeactivated, EventSubscriptionSelectedEventCategoryOutboundCardPushTransferCreated, EventSubscriptionSelectedEventCategoryOutboundCardPushTransferUpdated, EventSubscriptionSelectedEventCategoryOutboundCardValidationCreated, EventSubscriptionSelectedEventCategoryOutboundCardValidationUpdated, EventSubscriptionSelectedEventCategoryCardPushTransferCreated, EventSubscriptionSelectedEventCategoryCardPushTransferUpdated, EventSubscriptionSelectedEventCategoryCardValidationCreated, EventSubscriptionSelectedEventCategoryCardValidationUpdated, EventSubscriptionSelectedEventCategoryPendingTransactionCreated, EventSubscriptionSelectedEventCategoryPendingTransactionUpdated, EventSubscriptionSelectedEventCategoryPhysicalCardCreated, EventSubscriptionSelectedEventCategoryPhysicalCardUpdated, EventSubscriptionSelectedEventCategoryPhysicalCardProfileCreated, EventSubscriptionSelectedEventCategoryPhysicalCardProfileUpdated, EventSubscriptionSelectedEventCategoryProgramCreated, EventSubscriptionSelectedEventCategoryProgramUpdated, EventSubscriptionSelectedEventCategoryProofOfAuthorizationRequestCreated, EventSubscriptionSelectedEventCategoryProofOfAuthorizationRequestUpdated, EventSubscriptionSelectedEventCategoryRealTimeDecisionCardAuthorizationRequested, EventSubscriptionSelectedEventCategoryRealTimeDecisionDigitalWalletTokenRequested, EventSubscriptionSelectedEventCategoryRealTimeDecisionDigitalWalletAuthenticationRequested, EventSubscriptionSelectedEventCategoryRealTimeDecisionCardAuthenticationRequested, EventSubscriptionSelectedEventCategoryRealTimeDecisionCardAuthenticationChallengeRequested, EventSubscriptionSelectedEventCategoryRealTimePaymentsTransferCreated, EventSubscriptionSelectedEventCategoryRealTimePaymentsTransferUpdated, EventSubscriptionSelectedEventCategoryRealTimePaymentsRequestForPaymentCreated, EventSubscriptionSelectedEventCategoryRealTimePaymentsRequestForPaymentUpdated, EventSubscriptionSelectedEventCategorySwiftTransferCreated, EventSubscriptionSelectedEventCategorySwiftTransferUpdated, EventSubscriptionSelectedEventCategoryTransactionCreated, EventSubscriptionSelectedEventCategoryWireDrawdownRequestCreated, EventSubscriptionSelectedEventCategoryWireDrawdownRequestUpdated, EventSubscriptionSelectedEventCategoryWireTransferCreated, EventSubscriptionSelectedEventCategoryWireTransferUpdated:
		return true
	}
	return false
}

// This indicates if we'll send notifications to this subscription.
type EventSubscriptionStatus string

const (
	EventSubscriptionStatusActive            EventSubscriptionStatus = "active"
	EventSubscriptionStatusDisabled          EventSubscriptionStatus = "disabled"
	EventSubscriptionStatusDeleted           EventSubscriptionStatus = "deleted"
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
	EventSubscriptionNewParamsSelectedEventCategoryAccountCreated                                       EventSubscriptionNewParamsSelectedEventCategory = "account.created"
	EventSubscriptionNewParamsSelectedEventCategoryAccountUpdated                                       EventSubscriptionNewParamsSelectedEventCategory = "account.updated"
	EventSubscriptionNewParamsSelectedEventCategoryAccountNumberCreated                                 EventSubscriptionNewParamsSelectedEventCategory = "account_number.created"
	EventSubscriptionNewParamsSelectedEventCategoryAccountNumberUpdated                                 EventSubscriptionNewParamsSelectedEventCategory = "account_number.updated"
	EventSubscriptionNewParamsSelectedEventCategoryAccountStatementCreated                              EventSubscriptionNewParamsSelectedEventCategory = "account_statement.created"
	EventSubscriptionNewParamsSelectedEventCategoryAccountTransferCreated                               EventSubscriptionNewParamsSelectedEventCategory = "account_transfer.created"
	EventSubscriptionNewParamsSelectedEventCategoryAccountTransferUpdated                               EventSubscriptionNewParamsSelectedEventCategory = "account_transfer.updated"
	EventSubscriptionNewParamsSelectedEventCategoryACHPrenotificationCreated                            EventSubscriptionNewParamsSelectedEventCategory = "ach_prenotification.created"
	EventSubscriptionNewParamsSelectedEventCategoryACHPrenotificationUpdated                            EventSubscriptionNewParamsSelectedEventCategory = "ach_prenotification.updated"
	EventSubscriptionNewParamsSelectedEventCategoryACHTransferCreated                                   EventSubscriptionNewParamsSelectedEventCategory = "ach_transfer.created"
	EventSubscriptionNewParamsSelectedEventCategoryACHTransferUpdated                                   EventSubscriptionNewParamsSelectedEventCategory = "ach_transfer.updated"
	EventSubscriptionNewParamsSelectedEventCategoryBookkeepingAccountCreated                            EventSubscriptionNewParamsSelectedEventCategory = "bookkeeping_account.created"
	EventSubscriptionNewParamsSelectedEventCategoryBookkeepingAccountUpdated                            EventSubscriptionNewParamsSelectedEventCategory = "bookkeeping_account.updated"
	EventSubscriptionNewParamsSelectedEventCategoryBookkeepingEntrySetUpdated                           EventSubscriptionNewParamsSelectedEventCategory = "bookkeeping_entry_set.updated"
	EventSubscriptionNewParamsSelectedEventCategoryCardCreated                                          EventSubscriptionNewParamsSelectedEventCategory = "card.created"
	EventSubscriptionNewParamsSelectedEventCategoryCardUpdated                                          EventSubscriptionNewParamsSelectedEventCategory = "card.updated"
	EventSubscriptionNewParamsSelectedEventCategoryCardPaymentCreated                                   EventSubscriptionNewParamsSelectedEventCategory = "card_payment.created"
	EventSubscriptionNewParamsSelectedEventCategoryCardPaymentUpdated                                   EventSubscriptionNewParamsSelectedEventCategory = "card_payment.updated"
	EventSubscriptionNewParamsSelectedEventCategoryCardProfileCreated                                   EventSubscriptionNewParamsSelectedEventCategory = "card_profile.created"
	EventSubscriptionNewParamsSelectedEventCategoryCardProfileUpdated                                   EventSubscriptionNewParamsSelectedEventCategory = "card_profile.updated"
	EventSubscriptionNewParamsSelectedEventCategoryCardDisputeCreated                                   EventSubscriptionNewParamsSelectedEventCategory = "card_dispute.created"
	EventSubscriptionNewParamsSelectedEventCategoryCardDisputeUpdated                                   EventSubscriptionNewParamsSelectedEventCategory = "card_dispute.updated"
	EventSubscriptionNewParamsSelectedEventCategoryCheckDepositCreated                                  EventSubscriptionNewParamsSelectedEventCategory = "check_deposit.created"
	EventSubscriptionNewParamsSelectedEventCategoryCheckDepositUpdated                                  EventSubscriptionNewParamsSelectedEventCategory = "check_deposit.updated"
	EventSubscriptionNewParamsSelectedEventCategoryCheckTransferCreated                                 EventSubscriptionNewParamsSelectedEventCategory = "check_transfer.created"
	EventSubscriptionNewParamsSelectedEventCategoryCheckTransferUpdated                                 EventSubscriptionNewParamsSelectedEventCategory = "check_transfer.updated"
	EventSubscriptionNewParamsSelectedEventCategoryDeclinedTransactionCreated                           EventSubscriptionNewParamsSelectedEventCategory = "declined_transaction.created"
	EventSubscriptionNewParamsSelectedEventCategoryDigitalCardProfileCreated                            EventSubscriptionNewParamsSelectedEventCategory = "digital_card_profile.created"
	EventSubscriptionNewParamsSelectedEventCategoryDigitalCardProfileUpdated                            EventSubscriptionNewParamsSelectedEventCategory = "digital_card_profile.updated"
	EventSubscriptionNewParamsSelectedEventCategoryDigitalWalletTokenCreated                            EventSubscriptionNewParamsSelectedEventCategory = "digital_wallet_token.created"
	EventSubscriptionNewParamsSelectedEventCategoryDigitalWalletTokenUpdated                            EventSubscriptionNewParamsSelectedEventCategory = "digital_wallet_token.updated"
	EventSubscriptionNewParamsSelectedEventCategoryDocumentCreated                                      EventSubscriptionNewParamsSelectedEventCategory = "document.created"
	EventSubscriptionNewParamsSelectedEventCategoryEntityCreated                                        EventSubscriptionNewParamsSelectedEventCategory = "entity.created"
	EventSubscriptionNewParamsSelectedEventCategoryEntityUpdated                                        EventSubscriptionNewParamsSelectedEventCategory = "entity.updated"
	EventSubscriptionNewParamsSelectedEventCategoryEventSubscriptionCreated                             EventSubscriptionNewParamsSelectedEventCategory = "event_subscription.created"
	EventSubscriptionNewParamsSelectedEventCategoryEventSubscriptionUpdated                             EventSubscriptionNewParamsSelectedEventCategory = "event_subscription.updated"
	EventSubscriptionNewParamsSelectedEventCategoryExportCreated                                        EventSubscriptionNewParamsSelectedEventCategory = "export.created"
	EventSubscriptionNewParamsSelectedEventCategoryExportUpdated                                        EventSubscriptionNewParamsSelectedEventCategory = "export.updated"
	EventSubscriptionNewParamsSelectedEventCategoryExternalAccountCreated                               EventSubscriptionNewParamsSelectedEventCategory = "external_account.created"
	EventSubscriptionNewParamsSelectedEventCategoryExternalAccountUpdated                               EventSubscriptionNewParamsSelectedEventCategory = "external_account.updated"
	EventSubscriptionNewParamsSelectedEventCategoryFednowTransferCreated                                EventSubscriptionNewParamsSelectedEventCategory = "fednow_transfer.created"
	EventSubscriptionNewParamsSelectedEventCategoryFednowTransferUpdated                                EventSubscriptionNewParamsSelectedEventCategory = "fednow_transfer.updated"
	EventSubscriptionNewParamsSelectedEventCategoryFileCreated                                          EventSubscriptionNewParamsSelectedEventCategory = "file.created"
	EventSubscriptionNewParamsSelectedEventCategoryGroupUpdated                                         EventSubscriptionNewParamsSelectedEventCategory = "group.updated"
	EventSubscriptionNewParamsSelectedEventCategoryGroupHeartbeat                                       EventSubscriptionNewParamsSelectedEventCategory = "group.heartbeat"
	EventSubscriptionNewParamsSelectedEventCategoryInboundACHTransferCreated                            EventSubscriptionNewParamsSelectedEventCategory = "inbound_ach_transfer.created"
	EventSubscriptionNewParamsSelectedEventCategoryInboundACHTransferUpdated                            EventSubscriptionNewParamsSelectedEventCategory = "inbound_ach_transfer.updated"
	EventSubscriptionNewParamsSelectedEventCategoryInboundACHTransferReturnCreated                      EventSubscriptionNewParamsSelectedEventCategory = "inbound_ach_transfer_return.created"
	EventSubscriptionNewParamsSelectedEventCategoryInboundACHTransferReturnUpdated                      EventSubscriptionNewParamsSelectedEventCategory = "inbound_ach_transfer_return.updated"
	EventSubscriptionNewParamsSelectedEventCategoryInboundCheckDepositCreated                           EventSubscriptionNewParamsSelectedEventCategory = "inbound_check_deposit.created"
	EventSubscriptionNewParamsSelectedEventCategoryInboundCheckDepositUpdated                           EventSubscriptionNewParamsSelectedEventCategory = "inbound_check_deposit.updated"
	EventSubscriptionNewParamsSelectedEventCategoryInboundFednowTransferCreated                         EventSubscriptionNewParamsSelectedEventCategory = "inbound_fednow_transfer.created"
	EventSubscriptionNewParamsSelectedEventCategoryInboundFednowTransferUpdated                         EventSubscriptionNewParamsSelectedEventCategory = "inbound_fednow_transfer.updated"
	EventSubscriptionNewParamsSelectedEventCategoryInboundMailItemCreated                               EventSubscriptionNewParamsSelectedEventCategory = "inbound_mail_item.created"
	EventSubscriptionNewParamsSelectedEventCategoryInboundMailItemUpdated                               EventSubscriptionNewParamsSelectedEventCategory = "inbound_mail_item.updated"
	EventSubscriptionNewParamsSelectedEventCategoryInboundRealTimePaymentsTransferCreated               EventSubscriptionNewParamsSelectedEventCategory = "inbound_real_time_payments_transfer.created"
	EventSubscriptionNewParamsSelectedEventCategoryInboundRealTimePaymentsTransferUpdated               EventSubscriptionNewParamsSelectedEventCategory = "inbound_real_time_payments_transfer.updated"
	EventSubscriptionNewParamsSelectedEventCategoryInboundWireDrawdownRequestCreated                    EventSubscriptionNewParamsSelectedEventCategory = "inbound_wire_drawdown_request.created"
	EventSubscriptionNewParamsSelectedEventCategoryInboundWireTransferCreated                           EventSubscriptionNewParamsSelectedEventCategory = "inbound_wire_transfer.created"
	EventSubscriptionNewParamsSelectedEventCategoryInboundWireTransferUpdated                           EventSubscriptionNewParamsSelectedEventCategory = "inbound_wire_transfer.updated"
	EventSubscriptionNewParamsSelectedEventCategoryIntrafiAccountEnrollmentCreated                      EventSubscriptionNewParamsSelectedEventCategory = "intrafi_account_enrollment.created"
	EventSubscriptionNewParamsSelectedEventCategoryIntrafiAccountEnrollmentUpdated                      EventSubscriptionNewParamsSelectedEventCategory = "intrafi_account_enrollment.updated"
	EventSubscriptionNewParamsSelectedEventCategoryIntrafiExclusionCreated                              EventSubscriptionNewParamsSelectedEventCategory = "intrafi_exclusion.created"
	EventSubscriptionNewParamsSelectedEventCategoryIntrafiExclusionUpdated                              EventSubscriptionNewParamsSelectedEventCategory = "intrafi_exclusion.updated"
	EventSubscriptionNewParamsSelectedEventCategoryLegacyCardDisputeCreated                             EventSubscriptionNewParamsSelectedEventCategory = "legacy_card_dispute.created"
	EventSubscriptionNewParamsSelectedEventCategoryLegacyCardDisputeUpdated                             EventSubscriptionNewParamsSelectedEventCategory = "legacy_card_dispute.updated"
	EventSubscriptionNewParamsSelectedEventCategoryLockboxCreated                                       EventSubscriptionNewParamsSelectedEventCategory = "lockbox.created"
	EventSubscriptionNewParamsSelectedEventCategoryLockboxUpdated                                       EventSubscriptionNewParamsSelectedEventCategory = "lockbox.updated"
	EventSubscriptionNewParamsSelectedEventCategoryOAuthConnectionCreated                               EventSubscriptionNewParamsSelectedEventCategory = "oauth_connection.created"
	EventSubscriptionNewParamsSelectedEventCategoryOAuthConnectionDeactivated                           EventSubscriptionNewParamsSelectedEventCategory = "oauth_connection.deactivated"
	EventSubscriptionNewParamsSelectedEventCategoryOutboundCardPushTransferCreated                      EventSubscriptionNewParamsSelectedEventCategory = "outbound_card_push_transfer.created"
	EventSubscriptionNewParamsSelectedEventCategoryOutboundCardPushTransferUpdated                      EventSubscriptionNewParamsSelectedEventCategory = "outbound_card_push_transfer.updated"
	EventSubscriptionNewParamsSelectedEventCategoryOutboundCardValidationCreated                        EventSubscriptionNewParamsSelectedEventCategory = "outbound_card_validation.created"
	EventSubscriptionNewParamsSelectedEventCategoryOutboundCardValidationUpdated                        EventSubscriptionNewParamsSelectedEventCategory = "outbound_card_validation.updated"
	EventSubscriptionNewParamsSelectedEventCategoryCardPushTransferCreated                              EventSubscriptionNewParamsSelectedEventCategory = "card_push_transfer.created"
	EventSubscriptionNewParamsSelectedEventCategoryCardPushTransferUpdated                              EventSubscriptionNewParamsSelectedEventCategory = "card_push_transfer.updated"
	EventSubscriptionNewParamsSelectedEventCategoryCardValidationCreated                                EventSubscriptionNewParamsSelectedEventCategory = "card_validation.created"
	EventSubscriptionNewParamsSelectedEventCategoryCardValidationUpdated                                EventSubscriptionNewParamsSelectedEventCategory = "card_validation.updated"
	EventSubscriptionNewParamsSelectedEventCategoryPendingTransactionCreated                            EventSubscriptionNewParamsSelectedEventCategory = "pending_transaction.created"
	EventSubscriptionNewParamsSelectedEventCategoryPendingTransactionUpdated                            EventSubscriptionNewParamsSelectedEventCategory = "pending_transaction.updated"
	EventSubscriptionNewParamsSelectedEventCategoryPhysicalCardCreated                                  EventSubscriptionNewParamsSelectedEventCategory = "physical_card.created"
	EventSubscriptionNewParamsSelectedEventCategoryPhysicalCardUpdated                                  EventSubscriptionNewParamsSelectedEventCategory = "physical_card.updated"
	EventSubscriptionNewParamsSelectedEventCategoryPhysicalCardProfileCreated                           EventSubscriptionNewParamsSelectedEventCategory = "physical_card_profile.created"
	EventSubscriptionNewParamsSelectedEventCategoryPhysicalCardProfileUpdated                           EventSubscriptionNewParamsSelectedEventCategory = "physical_card_profile.updated"
	EventSubscriptionNewParamsSelectedEventCategoryProgramCreated                                       EventSubscriptionNewParamsSelectedEventCategory = "program.created"
	EventSubscriptionNewParamsSelectedEventCategoryProgramUpdated                                       EventSubscriptionNewParamsSelectedEventCategory = "program.updated"
	EventSubscriptionNewParamsSelectedEventCategoryProofOfAuthorizationRequestCreated                   EventSubscriptionNewParamsSelectedEventCategory = "proof_of_authorization_request.created"
	EventSubscriptionNewParamsSelectedEventCategoryProofOfAuthorizationRequestUpdated                   EventSubscriptionNewParamsSelectedEventCategory = "proof_of_authorization_request.updated"
	EventSubscriptionNewParamsSelectedEventCategoryRealTimeDecisionCardAuthorizationRequested           EventSubscriptionNewParamsSelectedEventCategory = "real_time_decision.card_authorization_requested"
	EventSubscriptionNewParamsSelectedEventCategoryRealTimeDecisionDigitalWalletTokenRequested          EventSubscriptionNewParamsSelectedEventCategory = "real_time_decision.digital_wallet_token_requested"
	EventSubscriptionNewParamsSelectedEventCategoryRealTimeDecisionDigitalWalletAuthenticationRequested EventSubscriptionNewParamsSelectedEventCategory = "real_time_decision.digital_wallet_authentication_requested"
	EventSubscriptionNewParamsSelectedEventCategoryRealTimeDecisionCardAuthenticationRequested          EventSubscriptionNewParamsSelectedEventCategory = "real_time_decision.card_authentication_requested"
	EventSubscriptionNewParamsSelectedEventCategoryRealTimeDecisionCardAuthenticationChallengeRequested EventSubscriptionNewParamsSelectedEventCategory = "real_time_decision.card_authentication_challenge_requested"
	EventSubscriptionNewParamsSelectedEventCategoryRealTimePaymentsTransferCreated                      EventSubscriptionNewParamsSelectedEventCategory = "real_time_payments_transfer.created"
	EventSubscriptionNewParamsSelectedEventCategoryRealTimePaymentsTransferUpdated                      EventSubscriptionNewParamsSelectedEventCategory = "real_time_payments_transfer.updated"
	EventSubscriptionNewParamsSelectedEventCategoryRealTimePaymentsRequestForPaymentCreated             EventSubscriptionNewParamsSelectedEventCategory = "real_time_payments_request_for_payment.created"
	EventSubscriptionNewParamsSelectedEventCategoryRealTimePaymentsRequestForPaymentUpdated             EventSubscriptionNewParamsSelectedEventCategory = "real_time_payments_request_for_payment.updated"
	EventSubscriptionNewParamsSelectedEventCategorySwiftTransferCreated                                 EventSubscriptionNewParamsSelectedEventCategory = "swift_transfer.created"
	EventSubscriptionNewParamsSelectedEventCategorySwiftTransferUpdated                                 EventSubscriptionNewParamsSelectedEventCategory = "swift_transfer.updated"
	EventSubscriptionNewParamsSelectedEventCategoryTransactionCreated                                   EventSubscriptionNewParamsSelectedEventCategory = "transaction.created"
	EventSubscriptionNewParamsSelectedEventCategoryWireDrawdownRequestCreated                           EventSubscriptionNewParamsSelectedEventCategory = "wire_drawdown_request.created"
	EventSubscriptionNewParamsSelectedEventCategoryWireDrawdownRequestUpdated                           EventSubscriptionNewParamsSelectedEventCategory = "wire_drawdown_request.updated"
	EventSubscriptionNewParamsSelectedEventCategoryWireTransferCreated                                  EventSubscriptionNewParamsSelectedEventCategory = "wire_transfer.created"
	EventSubscriptionNewParamsSelectedEventCategoryWireTransferUpdated                                  EventSubscriptionNewParamsSelectedEventCategory = "wire_transfer.updated"
)

func (r EventSubscriptionNewParamsSelectedEventCategory) IsKnown() bool {
	switch r {
	case EventSubscriptionNewParamsSelectedEventCategoryAccountCreated, EventSubscriptionNewParamsSelectedEventCategoryAccountUpdated, EventSubscriptionNewParamsSelectedEventCategoryAccountNumberCreated, EventSubscriptionNewParamsSelectedEventCategoryAccountNumberUpdated, EventSubscriptionNewParamsSelectedEventCategoryAccountStatementCreated, EventSubscriptionNewParamsSelectedEventCategoryAccountTransferCreated, EventSubscriptionNewParamsSelectedEventCategoryAccountTransferUpdated, EventSubscriptionNewParamsSelectedEventCategoryACHPrenotificationCreated, EventSubscriptionNewParamsSelectedEventCategoryACHPrenotificationUpdated, EventSubscriptionNewParamsSelectedEventCategoryACHTransferCreated, EventSubscriptionNewParamsSelectedEventCategoryACHTransferUpdated, EventSubscriptionNewParamsSelectedEventCategoryBookkeepingAccountCreated, EventSubscriptionNewParamsSelectedEventCategoryBookkeepingAccountUpdated, EventSubscriptionNewParamsSelectedEventCategoryBookkeepingEntrySetUpdated, EventSubscriptionNewParamsSelectedEventCategoryCardCreated, EventSubscriptionNewParamsSelectedEventCategoryCardUpdated, EventSubscriptionNewParamsSelectedEventCategoryCardPaymentCreated, EventSubscriptionNewParamsSelectedEventCategoryCardPaymentUpdated, EventSubscriptionNewParamsSelectedEventCategoryCardProfileCreated, EventSubscriptionNewParamsSelectedEventCategoryCardProfileUpdated, EventSubscriptionNewParamsSelectedEventCategoryCardDisputeCreated, EventSubscriptionNewParamsSelectedEventCategoryCardDisputeUpdated, EventSubscriptionNewParamsSelectedEventCategoryCheckDepositCreated, EventSubscriptionNewParamsSelectedEventCategoryCheckDepositUpdated, EventSubscriptionNewParamsSelectedEventCategoryCheckTransferCreated, EventSubscriptionNewParamsSelectedEventCategoryCheckTransferUpdated, EventSubscriptionNewParamsSelectedEventCategoryDeclinedTransactionCreated, EventSubscriptionNewParamsSelectedEventCategoryDigitalCardProfileCreated, EventSubscriptionNewParamsSelectedEventCategoryDigitalCardProfileUpdated, EventSubscriptionNewParamsSelectedEventCategoryDigitalWalletTokenCreated, EventSubscriptionNewParamsSelectedEventCategoryDigitalWalletTokenUpdated, EventSubscriptionNewParamsSelectedEventCategoryDocumentCreated, EventSubscriptionNewParamsSelectedEventCategoryEntityCreated, EventSubscriptionNewParamsSelectedEventCategoryEntityUpdated, EventSubscriptionNewParamsSelectedEventCategoryEventSubscriptionCreated, EventSubscriptionNewParamsSelectedEventCategoryEventSubscriptionUpdated, EventSubscriptionNewParamsSelectedEventCategoryExportCreated, EventSubscriptionNewParamsSelectedEventCategoryExportUpdated, EventSubscriptionNewParamsSelectedEventCategoryExternalAccountCreated, EventSubscriptionNewParamsSelectedEventCategoryExternalAccountUpdated, EventSubscriptionNewParamsSelectedEventCategoryFednowTransferCreated, EventSubscriptionNewParamsSelectedEventCategoryFednowTransferUpdated, EventSubscriptionNewParamsSelectedEventCategoryFileCreated, EventSubscriptionNewParamsSelectedEventCategoryGroupUpdated, EventSubscriptionNewParamsSelectedEventCategoryGroupHeartbeat, EventSubscriptionNewParamsSelectedEventCategoryInboundACHTransferCreated, EventSubscriptionNewParamsSelectedEventCategoryInboundACHTransferUpdated, EventSubscriptionNewParamsSelectedEventCategoryInboundACHTransferReturnCreated, EventSubscriptionNewParamsSelectedEventCategoryInboundACHTransferReturnUpdated, EventSubscriptionNewParamsSelectedEventCategoryInboundCheckDepositCreated, EventSubscriptionNewParamsSelectedEventCategoryInboundCheckDepositUpdated, EventSubscriptionNewParamsSelectedEventCategoryInboundFednowTransferCreated, EventSubscriptionNewParamsSelectedEventCategoryInboundFednowTransferUpdated, EventSubscriptionNewParamsSelectedEventCategoryInboundMailItemCreated, EventSubscriptionNewParamsSelectedEventCategoryInboundMailItemUpdated, EventSubscriptionNewParamsSelectedEventCategoryInboundRealTimePaymentsTransferCreated, EventSubscriptionNewParamsSelectedEventCategoryInboundRealTimePaymentsTransferUpdated, EventSubscriptionNewParamsSelectedEventCategoryInboundWireDrawdownRequestCreated, EventSubscriptionNewParamsSelectedEventCategoryInboundWireTransferCreated, EventSubscriptionNewParamsSelectedEventCategoryInboundWireTransferUpdated, EventSubscriptionNewParamsSelectedEventCategoryIntrafiAccountEnrollmentCreated, EventSubscriptionNewParamsSelectedEventCategoryIntrafiAccountEnrollmentUpdated, EventSubscriptionNewParamsSelectedEventCategoryIntrafiExclusionCreated, EventSubscriptionNewParamsSelectedEventCategoryIntrafiExclusionUpdated, EventSubscriptionNewParamsSelectedEventCategoryLegacyCardDisputeCreated, EventSubscriptionNewParamsSelectedEventCategoryLegacyCardDisputeUpdated, EventSubscriptionNewParamsSelectedEventCategoryLockboxCreated, EventSubscriptionNewParamsSelectedEventCategoryLockboxUpdated, EventSubscriptionNewParamsSelectedEventCategoryOAuthConnectionCreated, EventSubscriptionNewParamsSelectedEventCategoryOAuthConnectionDeactivated, EventSubscriptionNewParamsSelectedEventCategoryOutboundCardPushTransferCreated, EventSubscriptionNewParamsSelectedEventCategoryOutboundCardPushTransferUpdated, EventSubscriptionNewParamsSelectedEventCategoryOutboundCardValidationCreated, EventSubscriptionNewParamsSelectedEventCategoryOutboundCardValidationUpdated, EventSubscriptionNewParamsSelectedEventCategoryCardPushTransferCreated, EventSubscriptionNewParamsSelectedEventCategoryCardPushTransferUpdated, EventSubscriptionNewParamsSelectedEventCategoryCardValidationCreated, EventSubscriptionNewParamsSelectedEventCategoryCardValidationUpdated, EventSubscriptionNewParamsSelectedEventCategoryPendingTransactionCreated, EventSubscriptionNewParamsSelectedEventCategoryPendingTransactionUpdated, EventSubscriptionNewParamsSelectedEventCategoryPhysicalCardCreated, EventSubscriptionNewParamsSelectedEventCategoryPhysicalCardUpdated, EventSubscriptionNewParamsSelectedEventCategoryPhysicalCardProfileCreated, EventSubscriptionNewParamsSelectedEventCategoryPhysicalCardProfileUpdated, EventSubscriptionNewParamsSelectedEventCategoryProgramCreated, EventSubscriptionNewParamsSelectedEventCategoryProgramUpdated, EventSubscriptionNewParamsSelectedEventCategoryProofOfAuthorizationRequestCreated, EventSubscriptionNewParamsSelectedEventCategoryProofOfAuthorizationRequestUpdated, EventSubscriptionNewParamsSelectedEventCategoryRealTimeDecisionCardAuthorizationRequested, EventSubscriptionNewParamsSelectedEventCategoryRealTimeDecisionDigitalWalletTokenRequested, EventSubscriptionNewParamsSelectedEventCategoryRealTimeDecisionDigitalWalletAuthenticationRequested, EventSubscriptionNewParamsSelectedEventCategoryRealTimeDecisionCardAuthenticationRequested, EventSubscriptionNewParamsSelectedEventCategoryRealTimeDecisionCardAuthenticationChallengeRequested, EventSubscriptionNewParamsSelectedEventCategoryRealTimePaymentsTransferCreated, EventSubscriptionNewParamsSelectedEventCategoryRealTimePaymentsTransferUpdated, EventSubscriptionNewParamsSelectedEventCategoryRealTimePaymentsRequestForPaymentCreated, EventSubscriptionNewParamsSelectedEventCategoryRealTimePaymentsRequestForPaymentUpdated, EventSubscriptionNewParamsSelectedEventCategorySwiftTransferCreated, EventSubscriptionNewParamsSelectedEventCategorySwiftTransferUpdated, EventSubscriptionNewParamsSelectedEventCategoryTransactionCreated, EventSubscriptionNewParamsSelectedEventCategoryWireDrawdownRequestCreated, EventSubscriptionNewParamsSelectedEventCategoryWireDrawdownRequestUpdated, EventSubscriptionNewParamsSelectedEventCategoryWireTransferCreated, EventSubscriptionNewParamsSelectedEventCategoryWireTransferUpdated:
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
	EventSubscriptionUpdateParamsStatusActive   EventSubscriptionUpdateParamsStatus = "active"
	EventSubscriptionUpdateParamsStatusDisabled EventSubscriptionUpdateParamsStatus = "disabled"
	EventSubscriptionUpdateParamsStatusDeleted  EventSubscriptionUpdateParamsStatus = "deleted"
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

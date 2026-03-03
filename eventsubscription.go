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
	ID string `json:"id" api:"required"`
	// The time the event subscription was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The idempotency key you chose for this object. This value is unique across
	// Increase and is used to ensure that a request is only processed once. Learn more
	// about [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey string `json:"idempotency_key" api:"required,nullable"`
	// If specified, this subscription will only receive webhooks for Events associated
	// with this OAuth Connection.
	OAuthConnectionID string `json:"oauth_connection_id" api:"required,nullable"`
	// If specified, this subscription will only receive webhooks for Events with the
	// specified `category`.
	SelectedEventCategories []EventSubscriptionSelectedEventCategory `json:"selected_event_categories" api:"required,nullable"`
	// This indicates if we'll send notifications to this subscription.
	Status EventSubscriptionStatus `json:"status" api:"required"`
	// A constant representing the object's type. For this resource it will always be
	// `event_subscription`.
	Type EventSubscriptionType `json:"type" api:"required"`
	// The webhook url where we'll send notifications.
	URL         string                 `json:"url" api:"required"`
	ExtraFields map[string]interface{} `json:"-" api:"extrafields"`
	JSON        eventSubscriptionJSON  `json:"-"`
}

// eventSubscriptionJSON contains the JSON metadata for the struct
// [EventSubscription]
type eventSubscriptionJSON struct {
	ID                      apijson.Field
	CreatedAt               apijson.Field
	IdempotencyKey          apijson.Field
	OAuthConnectionID       apijson.Field
	SelectedEventCategories apijson.Field
	Status                  apijson.Field
	Type                    apijson.Field
	URL                     apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *EventSubscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventSubscriptionJSON) RawJSON() string {
	return r.raw
}

type EventSubscriptionSelectedEventCategory struct {
	// The category of the Event.
	EventCategory EventSubscriptionSelectedEventCategoriesEventCategory `json:"event_category" api:"required,nullable"`
	JSON          eventSubscriptionSelectedEventCategoryJSON            `json:"-"`
}

// eventSubscriptionSelectedEventCategoryJSON contains the JSON metadata for the
// struct [EventSubscriptionSelectedEventCategory]
type eventSubscriptionSelectedEventCategoryJSON struct {
	EventCategory apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *EventSubscriptionSelectedEventCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventSubscriptionSelectedEventCategoryJSON) RawJSON() string {
	return r.raw
}

// The category of the Event.
type EventSubscriptionSelectedEventCategoriesEventCategory string

const (
	EventSubscriptionSelectedEventCategoriesEventCategoryAccountCreated                                       EventSubscriptionSelectedEventCategoriesEventCategory = "account.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryAccountUpdated                                       EventSubscriptionSelectedEventCategoriesEventCategory = "account.updated"
	EventSubscriptionSelectedEventCategoriesEventCategoryAccountNumberCreated                                 EventSubscriptionSelectedEventCategoriesEventCategory = "account_number.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryAccountNumberUpdated                                 EventSubscriptionSelectedEventCategoriesEventCategory = "account_number.updated"
	EventSubscriptionSelectedEventCategoriesEventCategoryAccountStatementCreated                              EventSubscriptionSelectedEventCategoriesEventCategory = "account_statement.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryAccountTransferCreated                               EventSubscriptionSelectedEventCategoriesEventCategory = "account_transfer.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryAccountTransferUpdated                               EventSubscriptionSelectedEventCategoriesEventCategory = "account_transfer.updated"
	EventSubscriptionSelectedEventCategoriesEventCategoryACHPrenotificationCreated                            EventSubscriptionSelectedEventCategoriesEventCategory = "ach_prenotification.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryACHPrenotificationUpdated                            EventSubscriptionSelectedEventCategoriesEventCategory = "ach_prenotification.updated"
	EventSubscriptionSelectedEventCategoriesEventCategoryACHTransferCreated                                   EventSubscriptionSelectedEventCategoriesEventCategory = "ach_transfer.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryACHTransferUpdated                                   EventSubscriptionSelectedEventCategoriesEventCategory = "ach_transfer.updated"
	EventSubscriptionSelectedEventCategoriesEventCategoryBlockchainAddressCreated                             EventSubscriptionSelectedEventCategoriesEventCategory = "blockchain_address.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryBlockchainAddressUpdated                             EventSubscriptionSelectedEventCategoriesEventCategory = "blockchain_address.updated"
	EventSubscriptionSelectedEventCategoriesEventCategoryBlockchainOfframpTransferCreated                     EventSubscriptionSelectedEventCategoriesEventCategory = "blockchain_offramp_transfer.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryBlockchainOfframpTransferUpdated                     EventSubscriptionSelectedEventCategoriesEventCategory = "blockchain_offramp_transfer.updated"
	EventSubscriptionSelectedEventCategoriesEventCategoryBlockchainOnrampTransferCreated                      EventSubscriptionSelectedEventCategoriesEventCategory = "blockchain_onramp_transfer.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryBlockchainOnrampTransferUpdated                      EventSubscriptionSelectedEventCategoriesEventCategory = "blockchain_onramp_transfer.updated"
	EventSubscriptionSelectedEventCategoriesEventCategoryBookkeepingAccountCreated                            EventSubscriptionSelectedEventCategoriesEventCategory = "bookkeeping_account.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryBookkeepingAccountUpdated                            EventSubscriptionSelectedEventCategoriesEventCategory = "bookkeeping_account.updated"
	EventSubscriptionSelectedEventCategoriesEventCategoryBookkeepingEntrySetUpdated                           EventSubscriptionSelectedEventCategoriesEventCategory = "bookkeeping_entry_set.updated"
	EventSubscriptionSelectedEventCategoriesEventCategoryCardCreated                                          EventSubscriptionSelectedEventCategoriesEventCategory = "card.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryCardUpdated                                          EventSubscriptionSelectedEventCategoriesEventCategory = "card.updated"
	EventSubscriptionSelectedEventCategoriesEventCategoryCardPaymentCreated                                   EventSubscriptionSelectedEventCategoriesEventCategory = "card_payment.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryCardPaymentUpdated                                   EventSubscriptionSelectedEventCategoriesEventCategory = "card_payment.updated"
	EventSubscriptionSelectedEventCategoriesEventCategoryCardProfileCreated                                   EventSubscriptionSelectedEventCategoriesEventCategory = "card_profile.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryCardProfileUpdated                                   EventSubscriptionSelectedEventCategoriesEventCategory = "card_profile.updated"
	EventSubscriptionSelectedEventCategoriesEventCategoryCardDisputeCreated                                   EventSubscriptionSelectedEventCategoriesEventCategory = "card_dispute.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryCardDisputeUpdated                                   EventSubscriptionSelectedEventCategoriesEventCategory = "card_dispute.updated"
	EventSubscriptionSelectedEventCategoriesEventCategoryCheckDepositCreated                                  EventSubscriptionSelectedEventCategoriesEventCategory = "check_deposit.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryCheckDepositUpdated                                  EventSubscriptionSelectedEventCategoriesEventCategory = "check_deposit.updated"
	EventSubscriptionSelectedEventCategoriesEventCategoryCheckTransferCreated                                 EventSubscriptionSelectedEventCategoriesEventCategory = "check_transfer.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryCheckTransferUpdated                                 EventSubscriptionSelectedEventCategoriesEventCategory = "check_transfer.updated"
	EventSubscriptionSelectedEventCategoriesEventCategoryDeclinedTransactionCreated                           EventSubscriptionSelectedEventCategoriesEventCategory = "declined_transaction.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryDigitalCardProfileCreated                            EventSubscriptionSelectedEventCategoriesEventCategory = "digital_card_profile.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryDigitalCardProfileUpdated                            EventSubscriptionSelectedEventCategoriesEventCategory = "digital_card_profile.updated"
	EventSubscriptionSelectedEventCategoriesEventCategoryDigitalWalletTokenCreated                            EventSubscriptionSelectedEventCategoriesEventCategory = "digital_wallet_token.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryDigitalWalletTokenUpdated                            EventSubscriptionSelectedEventCategoriesEventCategory = "digital_wallet_token.updated"
	EventSubscriptionSelectedEventCategoriesEventCategoryDocumentCreated                                      EventSubscriptionSelectedEventCategoriesEventCategory = "document.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryEntityCreated                                        EventSubscriptionSelectedEventCategoriesEventCategory = "entity.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryEntityUpdated                                        EventSubscriptionSelectedEventCategoriesEventCategory = "entity.updated"
	EventSubscriptionSelectedEventCategoriesEventCategoryEventSubscriptionCreated                             EventSubscriptionSelectedEventCategoriesEventCategory = "event_subscription.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryEventSubscriptionUpdated                             EventSubscriptionSelectedEventCategoriesEventCategory = "event_subscription.updated"
	EventSubscriptionSelectedEventCategoriesEventCategoryExportCreated                                        EventSubscriptionSelectedEventCategoriesEventCategory = "export.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryExportUpdated                                        EventSubscriptionSelectedEventCategoriesEventCategory = "export.updated"
	EventSubscriptionSelectedEventCategoriesEventCategoryExternalAccountCreated                               EventSubscriptionSelectedEventCategoriesEventCategory = "external_account.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryExternalAccountUpdated                               EventSubscriptionSelectedEventCategoriesEventCategory = "external_account.updated"
	EventSubscriptionSelectedEventCategoriesEventCategoryFednowTransferCreated                                EventSubscriptionSelectedEventCategoriesEventCategory = "fednow_transfer.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryFednowTransferUpdated                                EventSubscriptionSelectedEventCategoriesEventCategory = "fednow_transfer.updated"
	EventSubscriptionSelectedEventCategoriesEventCategoryFileCreated                                          EventSubscriptionSelectedEventCategoriesEventCategory = "file.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryGroupUpdated                                         EventSubscriptionSelectedEventCategoriesEventCategory = "group.updated"
	EventSubscriptionSelectedEventCategoriesEventCategoryGroupHeartbeat                                       EventSubscriptionSelectedEventCategoriesEventCategory = "group.heartbeat"
	EventSubscriptionSelectedEventCategoriesEventCategoryInboundACHTransferCreated                            EventSubscriptionSelectedEventCategoriesEventCategory = "inbound_ach_transfer.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryInboundACHTransferUpdated                            EventSubscriptionSelectedEventCategoriesEventCategory = "inbound_ach_transfer.updated"
	EventSubscriptionSelectedEventCategoriesEventCategoryInboundACHTransferReturnCreated                      EventSubscriptionSelectedEventCategoriesEventCategory = "inbound_ach_transfer_return.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryInboundACHTransferReturnUpdated                      EventSubscriptionSelectedEventCategoriesEventCategory = "inbound_ach_transfer_return.updated"
	EventSubscriptionSelectedEventCategoriesEventCategoryInboundCheckDepositCreated                           EventSubscriptionSelectedEventCategoriesEventCategory = "inbound_check_deposit.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryInboundCheckDepositUpdated                           EventSubscriptionSelectedEventCategoriesEventCategory = "inbound_check_deposit.updated"
	EventSubscriptionSelectedEventCategoriesEventCategoryInboundFednowTransferCreated                         EventSubscriptionSelectedEventCategoriesEventCategory = "inbound_fednow_transfer.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryInboundFednowTransferUpdated                         EventSubscriptionSelectedEventCategoriesEventCategory = "inbound_fednow_transfer.updated"
	EventSubscriptionSelectedEventCategoriesEventCategoryInboundMailItemCreated                               EventSubscriptionSelectedEventCategoriesEventCategory = "inbound_mail_item.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryInboundMailItemUpdated                               EventSubscriptionSelectedEventCategoriesEventCategory = "inbound_mail_item.updated"
	EventSubscriptionSelectedEventCategoriesEventCategoryInboundRealTimePaymentsTransferCreated               EventSubscriptionSelectedEventCategoriesEventCategory = "inbound_real_time_payments_transfer.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryInboundRealTimePaymentsTransferUpdated               EventSubscriptionSelectedEventCategoriesEventCategory = "inbound_real_time_payments_transfer.updated"
	EventSubscriptionSelectedEventCategoriesEventCategoryInboundWireDrawdownRequestCreated                    EventSubscriptionSelectedEventCategoriesEventCategory = "inbound_wire_drawdown_request.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryInboundWireTransferCreated                           EventSubscriptionSelectedEventCategoriesEventCategory = "inbound_wire_transfer.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryInboundWireTransferUpdated                           EventSubscriptionSelectedEventCategoriesEventCategory = "inbound_wire_transfer.updated"
	EventSubscriptionSelectedEventCategoriesEventCategoryIntrafiAccountEnrollmentCreated                      EventSubscriptionSelectedEventCategoriesEventCategory = "intrafi_account_enrollment.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryIntrafiAccountEnrollmentUpdated                      EventSubscriptionSelectedEventCategoriesEventCategory = "intrafi_account_enrollment.updated"
	EventSubscriptionSelectedEventCategoriesEventCategoryIntrafiExclusionCreated                              EventSubscriptionSelectedEventCategoriesEventCategory = "intrafi_exclusion.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryIntrafiExclusionUpdated                              EventSubscriptionSelectedEventCategoriesEventCategory = "intrafi_exclusion.updated"
	EventSubscriptionSelectedEventCategoriesEventCategoryLegacyCardDisputeCreated                             EventSubscriptionSelectedEventCategoriesEventCategory = "legacy_card_dispute.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryLegacyCardDisputeUpdated                             EventSubscriptionSelectedEventCategoriesEventCategory = "legacy_card_dispute.updated"
	EventSubscriptionSelectedEventCategoriesEventCategoryLockboxCreated                                       EventSubscriptionSelectedEventCategoriesEventCategory = "lockbox.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryLockboxUpdated                                       EventSubscriptionSelectedEventCategoriesEventCategory = "lockbox.updated"
	EventSubscriptionSelectedEventCategoriesEventCategoryOAuthConnectionCreated                               EventSubscriptionSelectedEventCategoriesEventCategory = "oauth_connection.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryOAuthConnectionDeactivated                           EventSubscriptionSelectedEventCategoriesEventCategory = "oauth_connection.deactivated"
	EventSubscriptionSelectedEventCategoriesEventCategoryCardPushTransferCreated                              EventSubscriptionSelectedEventCategoriesEventCategory = "card_push_transfer.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryCardPushTransferUpdated                              EventSubscriptionSelectedEventCategoriesEventCategory = "card_push_transfer.updated"
	EventSubscriptionSelectedEventCategoriesEventCategoryCardValidationCreated                                EventSubscriptionSelectedEventCategoriesEventCategory = "card_validation.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryCardValidationUpdated                                EventSubscriptionSelectedEventCategoriesEventCategory = "card_validation.updated"
	EventSubscriptionSelectedEventCategoriesEventCategoryPendingTransactionCreated                            EventSubscriptionSelectedEventCategoriesEventCategory = "pending_transaction.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryPendingTransactionUpdated                            EventSubscriptionSelectedEventCategoriesEventCategory = "pending_transaction.updated"
	EventSubscriptionSelectedEventCategoriesEventCategoryPhysicalCardCreated                                  EventSubscriptionSelectedEventCategoriesEventCategory = "physical_card.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryPhysicalCardUpdated                                  EventSubscriptionSelectedEventCategoriesEventCategory = "physical_card.updated"
	EventSubscriptionSelectedEventCategoriesEventCategoryPhysicalCardProfileCreated                           EventSubscriptionSelectedEventCategoriesEventCategory = "physical_card_profile.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryPhysicalCardProfileUpdated                           EventSubscriptionSelectedEventCategoriesEventCategory = "physical_card_profile.updated"
	EventSubscriptionSelectedEventCategoriesEventCategoryPhysicalCheckCreated                                 EventSubscriptionSelectedEventCategoriesEventCategory = "physical_check.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryPhysicalCheckUpdated                                 EventSubscriptionSelectedEventCategoriesEventCategory = "physical_check.updated"
	EventSubscriptionSelectedEventCategoriesEventCategoryProgramCreated                                       EventSubscriptionSelectedEventCategoriesEventCategory = "program.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryProgramUpdated                                       EventSubscriptionSelectedEventCategoriesEventCategory = "program.updated"
	EventSubscriptionSelectedEventCategoriesEventCategoryProofOfAuthorizationRequestCreated                   EventSubscriptionSelectedEventCategoriesEventCategory = "proof_of_authorization_request.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryProofOfAuthorizationRequestUpdated                   EventSubscriptionSelectedEventCategoriesEventCategory = "proof_of_authorization_request.updated"
	EventSubscriptionSelectedEventCategoriesEventCategoryRealTimeDecisionCardAuthorizationRequested           EventSubscriptionSelectedEventCategoriesEventCategory = "real_time_decision.card_authorization_requested"
	EventSubscriptionSelectedEventCategoriesEventCategoryRealTimeDecisionCardBalanceInquiryRequested          EventSubscriptionSelectedEventCategoriesEventCategory = "real_time_decision.card_balance_inquiry_requested"
	EventSubscriptionSelectedEventCategoriesEventCategoryRealTimeDecisionDigitalWalletTokenRequested          EventSubscriptionSelectedEventCategoriesEventCategory = "real_time_decision.digital_wallet_token_requested"
	EventSubscriptionSelectedEventCategoriesEventCategoryRealTimeDecisionDigitalWalletAuthenticationRequested EventSubscriptionSelectedEventCategoriesEventCategory = "real_time_decision.digital_wallet_authentication_requested"
	EventSubscriptionSelectedEventCategoriesEventCategoryRealTimeDecisionCardAuthenticationRequested          EventSubscriptionSelectedEventCategoriesEventCategory = "real_time_decision.card_authentication_requested"
	EventSubscriptionSelectedEventCategoriesEventCategoryRealTimeDecisionCardAuthenticationChallengeRequested EventSubscriptionSelectedEventCategoriesEventCategory = "real_time_decision.card_authentication_challenge_requested"
	EventSubscriptionSelectedEventCategoriesEventCategoryRealTimePaymentsTransferCreated                      EventSubscriptionSelectedEventCategoriesEventCategory = "real_time_payments_transfer.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryRealTimePaymentsTransferUpdated                      EventSubscriptionSelectedEventCategoriesEventCategory = "real_time_payments_transfer.updated"
	EventSubscriptionSelectedEventCategoriesEventCategoryRealTimePaymentsRequestForPaymentCreated             EventSubscriptionSelectedEventCategoriesEventCategory = "real_time_payments_request_for_payment.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryRealTimePaymentsRequestForPaymentUpdated             EventSubscriptionSelectedEventCategoriesEventCategory = "real_time_payments_request_for_payment.updated"
	EventSubscriptionSelectedEventCategoriesEventCategorySwiftTransferCreated                                 EventSubscriptionSelectedEventCategoriesEventCategory = "swift_transfer.created"
	EventSubscriptionSelectedEventCategoriesEventCategorySwiftTransferUpdated                                 EventSubscriptionSelectedEventCategoriesEventCategory = "swift_transfer.updated"
	EventSubscriptionSelectedEventCategoriesEventCategoryTransactionCreated                                   EventSubscriptionSelectedEventCategoriesEventCategory = "transaction.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryWireDrawdownRequestCreated                           EventSubscriptionSelectedEventCategoriesEventCategory = "wire_drawdown_request.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryWireDrawdownRequestUpdated                           EventSubscriptionSelectedEventCategoriesEventCategory = "wire_drawdown_request.updated"
	EventSubscriptionSelectedEventCategoriesEventCategoryWireTransferCreated                                  EventSubscriptionSelectedEventCategoriesEventCategory = "wire_transfer.created"
	EventSubscriptionSelectedEventCategoriesEventCategoryWireTransferUpdated                                  EventSubscriptionSelectedEventCategoriesEventCategory = "wire_transfer.updated"
)

func (r EventSubscriptionSelectedEventCategoriesEventCategory) IsKnown() bool {
	switch r {
	case EventSubscriptionSelectedEventCategoriesEventCategoryAccountCreated, EventSubscriptionSelectedEventCategoriesEventCategoryAccountUpdated, EventSubscriptionSelectedEventCategoriesEventCategoryAccountNumberCreated, EventSubscriptionSelectedEventCategoriesEventCategoryAccountNumberUpdated, EventSubscriptionSelectedEventCategoriesEventCategoryAccountStatementCreated, EventSubscriptionSelectedEventCategoriesEventCategoryAccountTransferCreated, EventSubscriptionSelectedEventCategoriesEventCategoryAccountTransferUpdated, EventSubscriptionSelectedEventCategoriesEventCategoryACHPrenotificationCreated, EventSubscriptionSelectedEventCategoriesEventCategoryACHPrenotificationUpdated, EventSubscriptionSelectedEventCategoriesEventCategoryACHTransferCreated, EventSubscriptionSelectedEventCategoriesEventCategoryACHTransferUpdated, EventSubscriptionSelectedEventCategoriesEventCategoryBlockchainAddressCreated, EventSubscriptionSelectedEventCategoriesEventCategoryBlockchainAddressUpdated, EventSubscriptionSelectedEventCategoriesEventCategoryBlockchainOfframpTransferCreated, EventSubscriptionSelectedEventCategoriesEventCategoryBlockchainOfframpTransferUpdated, EventSubscriptionSelectedEventCategoriesEventCategoryBlockchainOnrampTransferCreated, EventSubscriptionSelectedEventCategoriesEventCategoryBlockchainOnrampTransferUpdated, EventSubscriptionSelectedEventCategoriesEventCategoryBookkeepingAccountCreated, EventSubscriptionSelectedEventCategoriesEventCategoryBookkeepingAccountUpdated, EventSubscriptionSelectedEventCategoriesEventCategoryBookkeepingEntrySetUpdated, EventSubscriptionSelectedEventCategoriesEventCategoryCardCreated, EventSubscriptionSelectedEventCategoriesEventCategoryCardUpdated, EventSubscriptionSelectedEventCategoriesEventCategoryCardPaymentCreated, EventSubscriptionSelectedEventCategoriesEventCategoryCardPaymentUpdated, EventSubscriptionSelectedEventCategoriesEventCategoryCardProfileCreated, EventSubscriptionSelectedEventCategoriesEventCategoryCardProfileUpdated, EventSubscriptionSelectedEventCategoriesEventCategoryCardDisputeCreated, EventSubscriptionSelectedEventCategoriesEventCategoryCardDisputeUpdated, EventSubscriptionSelectedEventCategoriesEventCategoryCheckDepositCreated, EventSubscriptionSelectedEventCategoriesEventCategoryCheckDepositUpdated, EventSubscriptionSelectedEventCategoriesEventCategoryCheckTransferCreated, EventSubscriptionSelectedEventCategoriesEventCategoryCheckTransferUpdated, EventSubscriptionSelectedEventCategoriesEventCategoryDeclinedTransactionCreated, EventSubscriptionSelectedEventCategoriesEventCategoryDigitalCardProfileCreated, EventSubscriptionSelectedEventCategoriesEventCategoryDigitalCardProfileUpdated, EventSubscriptionSelectedEventCategoriesEventCategoryDigitalWalletTokenCreated, EventSubscriptionSelectedEventCategoriesEventCategoryDigitalWalletTokenUpdated, EventSubscriptionSelectedEventCategoriesEventCategoryDocumentCreated, EventSubscriptionSelectedEventCategoriesEventCategoryEntityCreated, EventSubscriptionSelectedEventCategoriesEventCategoryEntityUpdated, EventSubscriptionSelectedEventCategoriesEventCategoryEventSubscriptionCreated, EventSubscriptionSelectedEventCategoriesEventCategoryEventSubscriptionUpdated, EventSubscriptionSelectedEventCategoriesEventCategoryExportCreated, EventSubscriptionSelectedEventCategoriesEventCategoryExportUpdated, EventSubscriptionSelectedEventCategoriesEventCategoryExternalAccountCreated, EventSubscriptionSelectedEventCategoriesEventCategoryExternalAccountUpdated, EventSubscriptionSelectedEventCategoriesEventCategoryFednowTransferCreated, EventSubscriptionSelectedEventCategoriesEventCategoryFednowTransferUpdated, EventSubscriptionSelectedEventCategoriesEventCategoryFileCreated, EventSubscriptionSelectedEventCategoriesEventCategoryGroupUpdated, EventSubscriptionSelectedEventCategoriesEventCategoryGroupHeartbeat, EventSubscriptionSelectedEventCategoriesEventCategoryInboundACHTransferCreated, EventSubscriptionSelectedEventCategoriesEventCategoryInboundACHTransferUpdated, EventSubscriptionSelectedEventCategoriesEventCategoryInboundACHTransferReturnCreated, EventSubscriptionSelectedEventCategoriesEventCategoryInboundACHTransferReturnUpdated, EventSubscriptionSelectedEventCategoriesEventCategoryInboundCheckDepositCreated, EventSubscriptionSelectedEventCategoriesEventCategoryInboundCheckDepositUpdated, EventSubscriptionSelectedEventCategoriesEventCategoryInboundFednowTransferCreated, EventSubscriptionSelectedEventCategoriesEventCategoryInboundFednowTransferUpdated, EventSubscriptionSelectedEventCategoriesEventCategoryInboundMailItemCreated, EventSubscriptionSelectedEventCategoriesEventCategoryInboundMailItemUpdated, EventSubscriptionSelectedEventCategoriesEventCategoryInboundRealTimePaymentsTransferCreated, EventSubscriptionSelectedEventCategoriesEventCategoryInboundRealTimePaymentsTransferUpdated, EventSubscriptionSelectedEventCategoriesEventCategoryInboundWireDrawdownRequestCreated, EventSubscriptionSelectedEventCategoriesEventCategoryInboundWireTransferCreated, EventSubscriptionSelectedEventCategoriesEventCategoryInboundWireTransferUpdated, EventSubscriptionSelectedEventCategoriesEventCategoryIntrafiAccountEnrollmentCreated, EventSubscriptionSelectedEventCategoriesEventCategoryIntrafiAccountEnrollmentUpdated, EventSubscriptionSelectedEventCategoriesEventCategoryIntrafiExclusionCreated, EventSubscriptionSelectedEventCategoriesEventCategoryIntrafiExclusionUpdated, EventSubscriptionSelectedEventCategoriesEventCategoryLegacyCardDisputeCreated, EventSubscriptionSelectedEventCategoriesEventCategoryLegacyCardDisputeUpdated, EventSubscriptionSelectedEventCategoriesEventCategoryLockboxCreated, EventSubscriptionSelectedEventCategoriesEventCategoryLockboxUpdated, EventSubscriptionSelectedEventCategoriesEventCategoryOAuthConnectionCreated, EventSubscriptionSelectedEventCategoriesEventCategoryOAuthConnectionDeactivated, EventSubscriptionSelectedEventCategoriesEventCategoryCardPushTransferCreated, EventSubscriptionSelectedEventCategoriesEventCategoryCardPushTransferUpdated, EventSubscriptionSelectedEventCategoriesEventCategoryCardValidationCreated, EventSubscriptionSelectedEventCategoriesEventCategoryCardValidationUpdated, EventSubscriptionSelectedEventCategoriesEventCategoryPendingTransactionCreated, EventSubscriptionSelectedEventCategoriesEventCategoryPendingTransactionUpdated, EventSubscriptionSelectedEventCategoriesEventCategoryPhysicalCardCreated, EventSubscriptionSelectedEventCategoriesEventCategoryPhysicalCardUpdated, EventSubscriptionSelectedEventCategoriesEventCategoryPhysicalCardProfileCreated, EventSubscriptionSelectedEventCategoriesEventCategoryPhysicalCardProfileUpdated, EventSubscriptionSelectedEventCategoriesEventCategoryPhysicalCheckCreated, EventSubscriptionSelectedEventCategoriesEventCategoryPhysicalCheckUpdated, EventSubscriptionSelectedEventCategoriesEventCategoryProgramCreated, EventSubscriptionSelectedEventCategoriesEventCategoryProgramUpdated, EventSubscriptionSelectedEventCategoriesEventCategoryProofOfAuthorizationRequestCreated, EventSubscriptionSelectedEventCategoriesEventCategoryProofOfAuthorizationRequestUpdated, EventSubscriptionSelectedEventCategoriesEventCategoryRealTimeDecisionCardAuthorizationRequested, EventSubscriptionSelectedEventCategoriesEventCategoryRealTimeDecisionCardBalanceInquiryRequested, EventSubscriptionSelectedEventCategoriesEventCategoryRealTimeDecisionDigitalWalletTokenRequested, EventSubscriptionSelectedEventCategoriesEventCategoryRealTimeDecisionDigitalWalletAuthenticationRequested, EventSubscriptionSelectedEventCategoriesEventCategoryRealTimeDecisionCardAuthenticationRequested, EventSubscriptionSelectedEventCategoriesEventCategoryRealTimeDecisionCardAuthenticationChallengeRequested, EventSubscriptionSelectedEventCategoriesEventCategoryRealTimePaymentsTransferCreated, EventSubscriptionSelectedEventCategoriesEventCategoryRealTimePaymentsTransferUpdated, EventSubscriptionSelectedEventCategoriesEventCategoryRealTimePaymentsRequestForPaymentCreated, EventSubscriptionSelectedEventCategoriesEventCategoryRealTimePaymentsRequestForPaymentUpdated, EventSubscriptionSelectedEventCategoriesEventCategorySwiftTransferCreated, EventSubscriptionSelectedEventCategoriesEventCategorySwiftTransferUpdated, EventSubscriptionSelectedEventCategoriesEventCategoryTransactionCreated, EventSubscriptionSelectedEventCategoriesEventCategoryWireDrawdownRequestCreated, EventSubscriptionSelectedEventCategoriesEventCategoryWireDrawdownRequestUpdated, EventSubscriptionSelectedEventCategoriesEventCategoryWireTransferCreated, EventSubscriptionSelectedEventCategoriesEventCategoryWireTransferUpdated:
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
	URL param.Field[string] `json:"url" api:"required"`
	// If specified, this subscription will only receive webhooks for Events associated
	// with the specified OAuth Connection.
	OAuthConnectionID param.Field[string] `json:"oauth_connection_id"`
	// If specified, this subscription will only receive webhooks for Events with the
	// specified `category`. If specifying a Real-Time Decision event category, only
	// one Event Category can be specified for the Event Subscription.
	SelectedEventCategories param.Field[[]EventSubscriptionNewParamsSelectedEventCategory] `json:"selected_event_categories"`
	// The key that will be used to sign webhooks. If no value is passed, a random
	// string will be used as default.
	SharedSecret param.Field[string] `json:"shared_secret"`
	// The status of the event subscription. Defaults to `active` if not specified.
	Status param.Field[EventSubscriptionNewParamsStatus] `json:"status"`
}

func (r EventSubscriptionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EventSubscriptionNewParamsSelectedEventCategory struct {
	// The category of the Event to subscribe to.
	EventCategory param.Field[EventSubscriptionNewParamsSelectedEventCategoriesEventCategory] `json:"event_category" api:"required"`
}

func (r EventSubscriptionNewParamsSelectedEventCategory) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The category of the Event to subscribe to.
type EventSubscriptionNewParamsSelectedEventCategoriesEventCategory string

const (
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryAccountCreated                                       EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "account.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryAccountUpdated                                       EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "account.updated"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryAccountNumberCreated                                 EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "account_number.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryAccountNumberUpdated                                 EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "account_number.updated"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryAccountStatementCreated                              EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "account_statement.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryAccountTransferCreated                               EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "account_transfer.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryAccountTransferUpdated                               EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "account_transfer.updated"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryACHPrenotificationCreated                            EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "ach_prenotification.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryACHPrenotificationUpdated                            EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "ach_prenotification.updated"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryACHTransferCreated                                   EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "ach_transfer.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryACHTransferUpdated                                   EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "ach_transfer.updated"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryBlockchainAddressCreated                             EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "blockchain_address.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryBlockchainAddressUpdated                             EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "blockchain_address.updated"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryBlockchainOfframpTransferCreated                     EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "blockchain_offramp_transfer.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryBlockchainOfframpTransferUpdated                     EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "blockchain_offramp_transfer.updated"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryBlockchainOnrampTransferCreated                      EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "blockchain_onramp_transfer.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryBlockchainOnrampTransferUpdated                      EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "blockchain_onramp_transfer.updated"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryBookkeepingAccountCreated                            EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "bookkeeping_account.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryBookkeepingAccountUpdated                            EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "bookkeeping_account.updated"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryBookkeepingEntrySetUpdated                           EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "bookkeeping_entry_set.updated"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryCardCreated                                          EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "card.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryCardUpdated                                          EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "card.updated"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryCardPaymentCreated                                   EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "card_payment.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryCardPaymentUpdated                                   EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "card_payment.updated"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryCardProfileCreated                                   EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "card_profile.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryCardProfileUpdated                                   EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "card_profile.updated"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryCardDisputeCreated                                   EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "card_dispute.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryCardDisputeUpdated                                   EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "card_dispute.updated"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryCheckDepositCreated                                  EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "check_deposit.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryCheckDepositUpdated                                  EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "check_deposit.updated"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryCheckTransferCreated                                 EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "check_transfer.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryCheckTransferUpdated                                 EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "check_transfer.updated"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryDeclinedTransactionCreated                           EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "declined_transaction.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryDigitalCardProfileCreated                            EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "digital_card_profile.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryDigitalCardProfileUpdated                            EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "digital_card_profile.updated"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryDigitalWalletTokenCreated                            EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "digital_wallet_token.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryDigitalWalletTokenUpdated                            EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "digital_wallet_token.updated"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryDocumentCreated                                      EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "document.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryEntityCreated                                        EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "entity.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryEntityUpdated                                        EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "entity.updated"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryEventSubscriptionCreated                             EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "event_subscription.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryEventSubscriptionUpdated                             EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "event_subscription.updated"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryExportCreated                                        EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "export.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryExportUpdated                                        EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "export.updated"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryExternalAccountCreated                               EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "external_account.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryExternalAccountUpdated                               EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "external_account.updated"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryFednowTransferCreated                                EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "fednow_transfer.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryFednowTransferUpdated                                EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "fednow_transfer.updated"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryFileCreated                                          EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "file.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryGroupUpdated                                         EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "group.updated"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryGroupHeartbeat                                       EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "group.heartbeat"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryInboundACHTransferCreated                            EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "inbound_ach_transfer.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryInboundACHTransferUpdated                            EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "inbound_ach_transfer.updated"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryInboundACHTransferReturnCreated                      EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "inbound_ach_transfer_return.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryInboundACHTransferReturnUpdated                      EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "inbound_ach_transfer_return.updated"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryInboundCheckDepositCreated                           EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "inbound_check_deposit.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryInboundCheckDepositUpdated                           EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "inbound_check_deposit.updated"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryInboundFednowTransferCreated                         EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "inbound_fednow_transfer.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryInboundFednowTransferUpdated                         EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "inbound_fednow_transfer.updated"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryInboundMailItemCreated                               EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "inbound_mail_item.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryInboundMailItemUpdated                               EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "inbound_mail_item.updated"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryInboundRealTimePaymentsTransferCreated               EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "inbound_real_time_payments_transfer.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryInboundRealTimePaymentsTransferUpdated               EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "inbound_real_time_payments_transfer.updated"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryInboundWireDrawdownRequestCreated                    EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "inbound_wire_drawdown_request.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryInboundWireTransferCreated                           EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "inbound_wire_transfer.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryInboundWireTransferUpdated                           EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "inbound_wire_transfer.updated"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryIntrafiAccountEnrollmentCreated                      EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "intrafi_account_enrollment.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryIntrafiAccountEnrollmentUpdated                      EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "intrafi_account_enrollment.updated"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryIntrafiExclusionCreated                              EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "intrafi_exclusion.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryIntrafiExclusionUpdated                              EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "intrafi_exclusion.updated"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryLegacyCardDisputeCreated                             EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "legacy_card_dispute.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryLegacyCardDisputeUpdated                             EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "legacy_card_dispute.updated"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryLockboxCreated                                       EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "lockbox.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryLockboxUpdated                                       EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "lockbox.updated"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryOAuthConnectionCreated                               EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "oauth_connection.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryOAuthConnectionDeactivated                           EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "oauth_connection.deactivated"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryCardPushTransferCreated                              EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "card_push_transfer.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryCardPushTransferUpdated                              EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "card_push_transfer.updated"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryCardValidationCreated                                EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "card_validation.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryCardValidationUpdated                                EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "card_validation.updated"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryPendingTransactionCreated                            EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "pending_transaction.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryPendingTransactionUpdated                            EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "pending_transaction.updated"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryPhysicalCardCreated                                  EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "physical_card.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryPhysicalCardUpdated                                  EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "physical_card.updated"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryPhysicalCardProfileCreated                           EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "physical_card_profile.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryPhysicalCardProfileUpdated                           EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "physical_card_profile.updated"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryPhysicalCheckCreated                                 EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "physical_check.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryPhysicalCheckUpdated                                 EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "physical_check.updated"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryProgramCreated                                       EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "program.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryProgramUpdated                                       EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "program.updated"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryProofOfAuthorizationRequestCreated                   EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "proof_of_authorization_request.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryProofOfAuthorizationRequestUpdated                   EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "proof_of_authorization_request.updated"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryRealTimeDecisionCardAuthorizationRequested           EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "real_time_decision.card_authorization_requested"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryRealTimeDecisionCardBalanceInquiryRequested          EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "real_time_decision.card_balance_inquiry_requested"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryRealTimeDecisionDigitalWalletTokenRequested          EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "real_time_decision.digital_wallet_token_requested"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryRealTimeDecisionDigitalWalletAuthenticationRequested EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "real_time_decision.digital_wallet_authentication_requested"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryRealTimeDecisionCardAuthenticationRequested          EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "real_time_decision.card_authentication_requested"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryRealTimeDecisionCardAuthenticationChallengeRequested EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "real_time_decision.card_authentication_challenge_requested"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryRealTimePaymentsTransferCreated                      EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "real_time_payments_transfer.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryRealTimePaymentsTransferUpdated                      EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "real_time_payments_transfer.updated"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryRealTimePaymentsRequestForPaymentCreated             EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "real_time_payments_request_for_payment.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryRealTimePaymentsRequestForPaymentUpdated             EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "real_time_payments_request_for_payment.updated"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategorySwiftTransferCreated                                 EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "swift_transfer.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategorySwiftTransferUpdated                                 EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "swift_transfer.updated"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryTransactionCreated                                   EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "transaction.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryWireDrawdownRequestCreated                           EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "wire_drawdown_request.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryWireDrawdownRequestUpdated                           EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "wire_drawdown_request.updated"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryWireTransferCreated                                  EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "wire_transfer.created"
	EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryWireTransferUpdated                                  EventSubscriptionNewParamsSelectedEventCategoriesEventCategory = "wire_transfer.updated"
)

func (r EventSubscriptionNewParamsSelectedEventCategoriesEventCategory) IsKnown() bool {
	switch r {
	case EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryAccountCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryAccountUpdated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryAccountNumberCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryAccountNumberUpdated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryAccountStatementCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryAccountTransferCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryAccountTransferUpdated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryACHPrenotificationCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryACHPrenotificationUpdated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryACHTransferCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryACHTransferUpdated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryBlockchainAddressCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryBlockchainAddressUpdated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryBlockchainOfframpTransferCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryBlockchainOfframpTransferUpdated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryBlockchainOnrampTransferCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryBlockchainOnrampTransferUpdated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryBookkeepingAccountCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryBookkeepingAccountUpdated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryBookkeepingEntrySetUpdated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryCardCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryCardUpdated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryCardPaymentCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryCardPaymentUpdated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryCardProfileCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryCardProfileUpdated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryCardDisputeCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryCardDisputeUpdated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryCheckDepositCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryCheckDepositUpdated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryCheckTransferCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryCheckTransferUpdated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryDeclinedTransactionCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryDigitalCardProfileCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryDigitalCardProfileUpdated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryDigitalWalletTokenCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryDigitalWalletTokenUpdated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryDocumentCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryEntityCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryEntityUpdated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryEventSubscriptionCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryEventSubscriptionUpdated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryExportCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryExportUpdated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryExternalAccountCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryExternalAccountUpdated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryFednowTransferCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryFednowTransferUpdated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryFileCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryGroupUpdated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryGroupHeartbeat, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryInboundACHTransferCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryInboundACHTransferUpdated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryInboundACHTransferReturnCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryInboundACHTransferReturnUpdated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryInboundCheckDepositCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryInboundCheckDepositUpdated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryInboundFednowTransferCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryInboundFednowTransferUpdated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryInboundMailItemCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryInboundMailItemUpdated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryInboundRealTimePaymentsTransferCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryInboundRealTimePaymentsTransferUpdated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryInboundWireDrawdownRequestCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryInboundWireTransferCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryInboundWireTransferUpdated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryIntrafiAccountEnrollmentCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryIntrafiAccountEnrollmentUpdated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryIntrafiExclusionCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryIntrafiExclusionUpdated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryLegacyCardDisputeCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryLegacyCardDisputeUpdated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryLockboxCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryLockboxUpdated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryOAuthConnectionCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryOAuthConnectionDeactivated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryCardPushTransferCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryCardPushTransferUpdated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryCardValidationCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryCardValidationUpdated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryPendingTransactionCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryPendingTransactionUpdated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryPhysicalCardCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryPhysicalCardUpdated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryPhysicalCardProfileCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryPhysicalCardProfileUpdated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryPhysicalCheckCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryPhysicalCheckUpdated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryProgramCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryProgramUpdated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryProofOfAuthorizationRequestCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryProofOfAuthorizationRequestUpdated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryRealTimeDecisionCardAuthorizationRequested, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryRealTimeDecisionCardBalanceInquiryRequested, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryRealTimeDecisionDigitalWalletTokenRequested, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryRealTimeDecisionDigitalWalletAuthenticationRequested, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryRealTimeDecisionCardAuthenticationRequested, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryRealTimeDecisionCardAuthenticationChallengeRequested, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryRealTimePaymentsTransferCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryRealTimePaymentsTransferUpdated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryRealTimePaymentsRequestForPaymentCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryRealTimePaymentsRequestForPaymentUpdated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategorySwiftTransferCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategorySwiftTransferUpdated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryTransactionCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryWireDrawdownRequestCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryWireDrawdownRequestUpdated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryWireTransferCreated, EventSubscriptionNewParamsSelectedEventCategoriesEventCategoryWireTransferUpdated:
		return true
	}
	return false
}

// The status of the event subscription. Defaults to `active` if not specified.
type EventSubscriptionNewParamsStatus string

const (
	EventSubscriptionNewParamsStatusActive   EventSubscriptionNewParamsStatus = "active"
	EventSubscriptionNewParamsStatusDisabled EventSubscriptionNewParamsStatus = "disabled"
)

func (r EventSubscriptionNewParamsStatus) IsKnown() bool {
	switch r {
	case EventSubscriptionNewParamsStatusActive, EventSubscriptionNewParamsStatusDisabled:
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

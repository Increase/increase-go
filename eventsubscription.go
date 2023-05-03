package increase

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/apiquery"
	"github.com/increase/increase-go/internal/field"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/internal/shared"
	"github.com/increase/increase-go/option"
)

// EventSubscriptionService contains methods and other services that help with
// interacting with the increase API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewEventSubscriptionService] method
// instead.
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
func (r *EventSubscriptionService) Get(ctx context.Context, event_subscription_id string, opts ...option.RequestOption) (res *EventSubscription, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("event_subscriptions/%s", event_subscription_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an Event Subscription
func (r *EventSubscriptionService) Update(ctx context.Context, event_subscription_id string, body EventSubscriptionUpdateParams, opts ...option.RequestOption) (res *EventSubscription, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("event_subscriptions/%s", event_subscription_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List Event Subscriptions
func (r *EventSubscriptionService) List(ctx context.Context, query EventSubscriptionListParams, opts ...option.RequestOption) (res *shared.Page[EventSubscription], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
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
func (r *EventSubscriptionService) ListAutoPaging(ctx context.Context, query EventSubscriptionListParams, opts ...option.RequestOption) *shared.PageAutoPager[EventSubscription] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
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
	// This indicates if we'll send notifications to this subscription.
	Status EventSubscriptionStatus `json:"status,required"`
	// If specified, this subscription will only receive webhooks for Events with the
	// specified `category`.
	SelectedEventCategory EventSubscriptionSelectedEventCategory `json:"selected_event_category,required,nullable"`
	// The webhook url where we'll send notifications.
	URL string `json:"url,required"`
	// The key that will be used to sign webhooks.
	SharedSecret string `json:"shared_secret,required"`
	// A constant representing the object's type. For this resource it will always be
	// `event_subscription`.
	Type EventSubscriptionType `json:"type,required"`
	JSON eventSubscriptionJSON
}

// eventSubscriptionJSON contains the JSON metadata for the struct
// [EventSubscription]
type eventSubscriptionJSON struct {
	ID                    apijson.Field
	CreatedAt             apijson.Field
	Status                apijson.Field
	SelectedEventCategory apijson.Field
	URL                   apijson.Field
	SharedSecret          apijson.Field
	Type                  apijson.Field
	raw                   string
	Extras                map[string]apijson.Field
}

func (r *EventSubscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EventSubscriptionStatus string

const (
	EventSubscriptionStatusActive            EventSubscriptionStatus = "active"
	EventSubscriptionStatusDisabled          EventSubscriptionStatus = "disabled"
	EventSubscriptionStatusDeleted           EventSubscriptionStatus = "deleted"
	EventSubscriptionStatusRequiresAttention EventSubscriptionStatus = "requires_attention"
)

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
	EventSubscriptionSelectedEventCategoryCardCreated                                          EventSubscriptionSelectedEventCategory = "card.created"
	EventSubscriptionSelectedEventCategoryCardUpdated                                          EventSubscriptionSelectedEventCategory = "card.updated"
	EventSubscriptionSelectedEventCategoryCardPaymentCreated                                   EventSubscriptionSelectedEventCategory = "card_payment.created"
	EventSubscriptionSelectedEventCategoryCardPaymentUpdated                                   EventSubscriptionSelectedEventCategory = "card_payment.updated"
	EventSubscriptionSelectedEventCategoryCardDisputeCreated                                   EventSubscriptionSelectedEventCategory = "card_dispute.created"
	EventSubscriptionSelectedEventCategoryCardDisputeUpdated                                   EventSubscriptionSelectedEventCategory = "card_dispute.updated"
	EventSubscriptionSelectedEventCategoryCheckDepositCreated                                  EventSubscriptionSelectedEventCategory = "check_deposit.created"
	EventSubscriptionSelectedEventCategoryCheckDepositUpdated                                  EventSubscriptionSelectedEventCategory = "check_deposit.updated"
	EventSubscriptionSelectedEventCategoryCheckTransferCreated                                 EventSubscriptionSelectedEventCategory = "check_transfer.created"
	EventSubscriptionSelectedEventCategoryCheckTransferUpdated                                 EventSubscriptionSelectedEventCategory = "check_transfer.updated"
	EventSubscriptionSelectedEventCategoryDeclinedTransactionCreated                           EventSubscriptionSelectedEventCategory = "declined_transaction.created"
	EventSubscriptionSelectedEventCategoryDigitalWalletTokenCreated                            EventSubscriptionSelectedEventCategory = "digital_wallet_token.created"
	EventSubscriptionSelectedEventCategoryDigitalWalletTokenUpdated                            EventSubscriptionSelectedEventCategory = "digital_wallet_token.updated"
	EventSubscriptionSelectedEventCategoryDocumentCreated                                      EventSubscriptionSelectedEventCategory = "document.created"
	EventSubscriptionSelectedEventCategoryEntityCreated                                        EventSubscriptionSelectedEventCategory = "entity.created"
	EventSubscriptionSelectedEventCategoryEntityUpdated                                        EventSubscriptionSelectedEventCategory = "entity.updated"
	EventSubscriptionSelectedEventCategoryExternalAccountCreated                               EventSubscriptionSelectedEventCategory = "external_account.created"
	EventSubscriptionSelectedEventCategoryFileCreated                                          EventSubscriptionSelectedEventCategory = "file.created"
	EventSubscriptionSelectedEventCategoryGroupUpdated                                         EventSubscriptionSelectedEventCategory = "group.updated"
	EventSubscriptionSelectedEventCategoryGroupHeartbeat                                       EventSubscriptionSelectedEventCategory = "group.heartbeat"
	EventSubscriptionSelectedEventCategoryInboundACHTransferReturnCreated                      EventSubscriptionSelectedEventCategory = "inbound_ach_transfer_return.created"
	EventSubscriptionSelectedEventCategoryInboundACHTransferReturnUpdated                      EventSubscriptionSelectedEventCategory = "inbound_ach_transfer_return.updated"
	EventSubscriptionSelectedEventCategoryInboundWireDrawdownRequestCreated                    EventSubscriptionSelectedEventCategory = "inbound_wire_drawdown_request.created"
	EventSubscriptionSelectedEventCategoryOauthConnectionCreated                               EventSubscriptionSelectedEventCategory = "oauth_connection.created"
	EventSubscriptionSelectedEventCategoryOauthConnectionDeactivated                           EventSubscriptionSelectedEventCategory = "oauth_connection.deactivated"
	EventSubscriptionSelectedEventCategoryPendingTransactionCreated                            EventSubscriptionSelectedEventCategory = "pending_transaction.created"
	EventSubscriptionSelectedEventCategoryPendingTransactionUpdated                            EventSubscriptionSelectedEventCategory = "pending_transaction.updated"
	EventSubscriptionSelectedEventCategoryRealTimeDecisionCardAuthorizationRequested           EventSubscriptionSelectedEventCategory = "real_time_decision.card_authorization_requested"
	EventSubscriptionSelectedEventCategoryRealTimeDecisionDigitalWalletTokenRequested          EventSubscriptionSelectedEventCategory = "real_time_decision.digital_wallet_token_requested"
	EventSubscriptionSelectedEventCategoryRealTimeDecisionDigitalWalletAuthenticationRequested EventSubscriptionSelectedEventCategory = "real_time_decision.digital_wallet_authentication_requested"
	EventSubscriptionSelectedEventCategoryRealTimePaymentsTransferCreated                      EventSubscriptionSelectedEventCategory = "real_time_payments_transfer.created"
	EventSubscriptionSelectedEventCategoryRealTimePaymentsTransferUpdated                      EventSubscriptionSelectedEventCategory = "real_time_payments_transfer.updated"
	EventSubscriptionSelectedEventCategoryRealTimePaymentsRequestForPaymentCreated             EventSubscriptionSelectedEventCategory = "real_time_payments_request_for_payment.created"
	EventSubscriptionSelectedEventCategoryRealTimePaymentsRequestForPaymentUpdated             EventSubscriptionSelectedEventCategory = "real_time_payments_request_for_payment.updated"
	EventSubscriptionSelectedEventCategoryTransactionCreated                                   EventSubscriptionSelectedEventCategory = "transaction.created"
	EventSubscriptionSelectedEventCategoryWireDrawdownRequestCreated                           EventSubscriptionSelectedEventCategory = "wire_drawdown_request.created"
	EventSubscriptionSelectedEventCategoryWireDrawdownRequestUpdated                           EventSubscriptionSelectedEventCategory = "wire_drawdown_request.updated"
	EventSubscriptionSelectedEventCategoryWireTransferCreated                                  EventSubscriptionSelectedEventCategory = "wire_transfer.created"
	EventSubscriptionSelectedEventCategoryWireTransferUpdated                                  EventSubscriptionSelectedEventCategory = "wire_transfer.updated"
)

type EventSubscriptionType string

const (
	EventSubscriptionTypeEventSubscription EventSubscriptionType = "event_subscription"
)

type EventSubscriptionNewParams struct {
	// The URL you'd like us to send webhooks to.
	URL field.Field[string] `json:"url,required"`
	// The key that will be used to sign webhooks. If no value is passed, a random
	// string will be used as default.
	SharedSecret field.Field[string] `json:"shared_secret"`
	// If specified, this subscription will only receive webhooks for Events with the
	// specified `category`.
	SelectedEventCategory field.Field[EventSubscriptionNewParamsSelectedEventCategory] `json:"selected_event_category"`
}

func (r EventSubscriptionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

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
	EventSubscriptionNewParamsSelectedEventCategoryCardCreated                                          EventSubscriptionNewParamsSelectedEventCategory = "card.created"
	EventSubscriptionNewParamsSelectedEventCategoryCardUpdated                                          EventSubscriptionNewParamsSelectedEventCategory = "card.updated"
	EventSubscriptionNewParamsSelectedEventCategoryCardPaymentCreated                                   EventSubscriptionNewParamsSelectedEventCategory = "card_payment.created"
	EventSubscriptionNewParamsSelectedEventCategoryCardPaymentUpdated                                   EventSubscriptionNewParamsSelectedEventCategory = "card_payment.updated"
	EventSubscriptionNewParamsSelectedEventCategoryCardDisputeCreated                                   EventSubscriptionNewParamsSelectedEventCategory = "card_dispute.created"
	EventSubscriptionNewParamsSelectedEventCategoryCardDisputeUpdated                                   EventSubscriptionNewParamsSelectedEventCategory = "card_dispute.updated"
	EventSubscriptionNewParamsSelectedEventCategoryCheckDepositCreated                                  EventSubscriptionNewParamsSelectedEventCategory = "check_deposit.created"
	EventSubscriptionNewParamsSelectedEventCategoryCheckDepositUpdated                                  EventSubscriptionNewParamsSelectedEventCategory = "check_deposit.updated"
	EventSubscriptionNewParamsSelectedEventCategoryCheckTransferCreated                                 EventSubscriptionNewParamsSelectedEventCategory = "check_transfer.created"
	EventSubscriptionNewParamsSelectedEventCategoryCheckTransferUpdated                                 EventSubscriptionNewParamsSelectedEventCategory = "check_transfer.updated"
	EventSubscriptionNewParamsSelectedEventCategoryDeclinedTransactionCreated                           EventSubscriptionNewParamsSelectedEventCategory = "declined_transaction.created"
	EventSubscriptionNewParamsSelectedEventCategoryDigitalWalletTokenCreated                            EventSubscriptionNewParamsSelectedEventCategory = "digital_wallet_token.created"
	EventSubscriptionNewParamsSelectedEventCategoryDigitalWalletTokenUpdated                            EventSubscriptionNewParamsSelectedEventCategory = "digital_wallet_token.updated"
	EventSubscriptionNewParamsSelectedEventCategoryDocumentCreated                                      EventSubscriptionNewParamsSelectedEventCategory = "document.created"
	EventSubscriptionNewParamsSelectedEventCategoryEntityCreated                                        EventSubscriptionNewParamsSelectedEventCategory = "entity.created"
	EventSubscriptionNewParamsSelectedEventCategoryEntityUpdated                                        EventSubscriptionNewParamsSelectedEventCategory = "entity.updated"
	EventSubscriptionNewParamsSelectedEventCategoryExternalAccountCreated                               EventSubscriptionNewParamsSelectedEventCategory = "external_account.created"
	EventSubscriptionNewParamsSelectedEventCategoryFileCreated                                          EventSubscriptionNewParamsSelectedEventCategory = "file.created"
	EventSubscriptionNewParamsSelectedEventCategoryGroupUpdated                                         EventSubscriptionNewParamsSelectedEventCategory = "group.updated"
	EventSubscriptionNewParamsSelectedEventCategoryGroupHeartbeat                                       EventSubscriptionNewParamsSelectedEventCategory = "group.heartbeat"
	EventSubscriptionNewParamsSelectedEventCategoryInboundACHTransferReturnCreated                      EventSubscriptionNewParamsSelectedEventCategory = "inbound_ach_transfer_return.created"
	EventSubscriptionNewParamsSelectedEventCategoryInboundACHTransferReturnUpdated                      EventSubscriptionNewParamsSelectedEventCategory = "inbound_ach_transfer_return.updated"
	EventSubscriptionNewParamsSelectedEventCategoryInboundWireDrawdownRequestCreated                    EventSubscriptionNewParamsSelectedEventCategory = "inbound_wire_drawdown_request.created"
	EventSubscriptionNewParamsSelectedEventCategoryOauthConnectionCreated                               EventSubscriptionNewParamsSelectedEventCategory = "oauth_connection.created"
	EventSubscriptionNewParamsSelectedEventCategoryOauthConnectionDeactivated                           EventSubscriptionNewParamsSelectedEventCategory = "oauth_connection.deactivated"
	EventSubscriptionNewParamsSelectedEventCategoryPendingTransactionCreated                            EventSubscriptionNewParamsSelectedEventCategory = "pending_transaction.created"
	EventSubscriptionNewParamsSelectedEventCategoryPendingTransactionUpdated                            EventSubscriptionNewParamsSelectedEventCategory = "pending_transaction.updated"
	EventSubscriptionNewParamsSelectedEventCategoryRealTimeDecisionCardAuthorizationRequested           EventSubscriptionNewParamsSelectedEventCategory = "real_time_decision.card_authorization_requested"
	EventSubscriptionNewParamsSelectedEventCategoryRealTimeDecisionDigitalWalletTokenRequested          EventSubscriptionNewParamsSelectedEventCategory = "real_time_decision.digital_wallet_token_requested"
	EventSubscriptionNewParamsSelectedEventCategoryRealTimeDecisionDigitalWalletAuthenticationRequested EventSubscriptionNewParamsSelectedEventCategory = "real_time_decision.digital_wallet_authentication_requested"
	EventSubscriptionNewParamsSelectedEventCategoryRealTimePaymentsTransferCreated                      EventSubscriptionNewParamsSelectedEventCategory = "real_time_payments_transfer.created"
	EventSubscriptionNewParamsSelectedEventCategoryRealTimePaymentsTransferUpdated                      EventSubscriptionNewParamsSelectedEventCategory = "real_time_payments_transfer.updated"
	EventSubscriptionNewParamsSelectedEventCategoryRealTimePaymentsRequestForPaymentCreated             EventSubscriptionNewParamsSelectedEventCategory = "real_time_payments_request_for_payment.created"
	EventSubscriptionNewParamsSelectedEventCategoryRealTimePaymentsRequestForPaymentUpdated             EventSubscriptionNewParamsSelectedEventCategory = "real_time_payments_request_for_payment.updated"
	EventSubscriptionNewParamsSelectedEventCategoryTransactionCreated                                   EventSubscriptionNewParamsSelectedEventCategory = "transaction.created"
	EventSubscriptionNewParamsSelectedEventCategoryWireDrawdownRequestCreated                           EventSubscriptionNewParamsSelectedEventCategory = "wire_drawdown_request.created"
	EventSubscriptionNewParamsSelectedEventCategoryWireDrawdownRequestUpdated                           EventSubscriptionNewParamsSelectedEventCategory = "wire_drawdown_request.updated"
	EventSubscriptionNewParamsSelectedEventCategoryWireTransferCreated                                  EventSubscriptionNewParamsSelectedEventCategory = "wire_transfer.created"
	EventSubscriptionNewParamsSelectedEventCategoryWireTransferUpdated                                  EventSubscriptionNewParamsSelectedEventCategory = "wire_transfer.updated"
)

type EventSubscriptionUpdateParams struct {
	// The status to update the Event Subscription with.
	Status field.Field[EventSubscriptionUpdateParamsStatus] `json:"status"`
}

func (r EventSubscriptionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EventSubscriptionUpdateParamsStatus string

const (
	EventSubscriptionUpdateParamsStatusActive   EventSubscriptionUpdateParamsStatus = "active"
	EventSubscriptionUpdateParamsStatusDisabled EventSubscriptionUpdateParamsStatus = "disabled"
	EventSubscriptionUpdateParamsStatusDeleted  EventSubscriptionUpdateParamsStatus = "deleted"
)

type EventSubscriptionListParams struct {
	// Return the page of entries after this one.
	Cursor field.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit field.Field[int64] `query:"limit"`
}

// URLQuery serializes [EventSubscriptionListParams]'s query parameters as
// `url.Values`.
func (r EventSubscriptionListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

// A list of Event Subscription objects
type EventSubscriptionListResponse struct {
	// The contents of the list.
	Data []EventSubscription `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       eventSubscriptionListResponseJSON
}

// eventSubscriptionListResponseJSON contains the JSON metadata for the struct
// [EventSubscriptionListResponse]
type eventSubscriptionListResponseJSON struct {
	Data       apijson.Field
	NextCursor apijson.Field
	raw        string
	Extras     map[string]apijson.Field
}

func (r *EventSubscriptionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

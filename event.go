// File generated from our OpenAPI spec by Stainless.

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

// EventService contains methods and other services that help with interacting with
// the increase API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewEventService] method instead.
type EventService struct {
	Options []option.RequestOption
}

// NewEventService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEventService(opts ...option.RequestOption) (r *EventService) {
	r = &EventService{}
	r.Options = opts
	return
}

// Retrieve an Event
func (r *EventService) Get(ctx context.Context, eventID string, opts ...option.RequestOption) (res *Event, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("events/%s", eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Events
func (r *EventService) List(ctx context.Context, query EventListParams, opts ...option.RequestOption) (res *shared.Page[Event], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "events"
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

// List Events
func (r *EventService) ListAutoPaging(ctx context.Context, query EventListParams, opts ...option.RequestOption) *shared.PageAutoPager[Event] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Events are records of things that happened to objects in the API. They also
// result in webhooks being generated.
type Event struct {
	// The identifier of the object that generated this Event.
	AssociatedObjectID string `json:"associated_object_id,required"`
	// The type of the object that generated this Event.
	AssociatedObjectType string `json:"associated_object_type,required"`
	// The category of the Event. We may add additional possible values for this enum
	// over time; your application should be able to handle such additions gracefully.
	Category EventCategory `json:"category,required"`
	// The time the Event was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The Event identifier.
	ID string `json:"id,required"`
	// A constant representing the object's type. For this resource it will always be
	// `event`.
	Type EventType `json:"type,required"`
	JSON eventJSON
}

// eventJSON contains the JSON metadata for the struct [Event]
type eventJSON struct {
	AssociatedObjectID   apijson.Field
	AssociatedObjectType apijson.Field
	Category             apijson.Field
	CreatedAt            apijson.Field
	ID                   apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *Event) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The category of the Event. We may add additional possible values for this enum
// over time; your application should be able to handle such additions gracefully.
type EventCategory string

const (
	EventCategoryAccountCreated                                       EventCategory = "account.created"
	EventCategoryAccountUpdated                                       EventCategory = "account.updated"
	EventCategoryAccountNumberCreated                                 EventCategory = "account_number.created"
	EventCategoryAccountNumberUpdated                                 EventCategory = "account_number.updated"
	EventCategoryAccountStatementCreated                              EventCategory = "account_statement.created"
	EventCategoryAccountTransferCreated                               EventCategory = "account_transfer.created"
	EventCategoryAccountTransferUpdated                               EventCategory = "account_transfer.updated"
	EventCategoryACHPrenotificationCreated                            EventCategory = "ach_prenotification.created"
	EventCategoryACHPrenotificationUpdated                            EventCategory = "ach_prenotification.updated"
	EventCategoryACHTransferCreated                                   EventCategory = "ach_transfer.created"
	EventCategoryACHTransferUpdated                                   EventCategory = "ach_transfer.updated"
	EventCategoryCardCreated                                          EventCategory = "card.created"
	EventCategoryCardUpdated                                          EventCategory = "card.updated"
	EventCategoryCardPaymentCreated                                   EventCategory = "card_payment.created"
	EventCategoryCardPaymentUpdated                                   EventCategory = "card_payment.updated"
	EventCategoryCardDisputeCreated                                   EventCategory = "card_dispute.created"
	EventCategoryCardDisputeUpdated                                   EventCategory = "card_dispute.updated"
	EventCategoryCheckDepositCreated                                  EventCategory = "check_deposit.created"
	EventCategoryCheckDepositUpdated                                  EventCategory = "check_deposit.updated"
	EventCategoryCheckTransferCreated                                 EventCategory = "check_transfer.created"
	EventCategoryCheckTransferUpdated                                 EventCategory = "check_transfer.updated"
	EventCategoryDeclinedTransactionCreated                           EventCategory = "declined_transaction.created"
	EventCategoryDigitalWalletTokenCreated                            EventCategory = "digital_wallet_token.created"
	EventCategoryDigitalWalletTokenUpdated                            EventCategory = "digital_wallet_token.updated"
	EventCategoryDocumentCreated                                      EventCategory = "document.created"
	EventCategoryEntityCreated                                        EventCategory = "entity.created"
	EventCategoryEntityUpdated                                        EventCategory = "entity.updated"
	EventCategoryExternalAccountCreated                               EventCategory = "external_account.created"
	EventCategoryFileCreated                                          EventCategory = "file.created"
	EventCategoryGroupUpdated                                         EventCategory = "group.updated"
	EventCategoryGroupHeartbeat                                       EventCategory = "group.heartbeat"
	EventCategoryInboundACHTransferReturnCreated                      EventCategory = "inbound_ach_transfer_return.created"
	EventCategoryInboundACHTransferReturnUpdated                      EventCategory = "inbound_ach_transfer_return.updated"
	EventCategoryInboundWireDrawdownRequestCreated                    EventCategory = "inbound_wire_drawdown_request.created"
	EventCategoryOauthConnectionCreated                               EventCategory = "oauth_connection.created"
	EventCategoryOauthConnectionDeactivated                           EventCategory = "oauth_connection.deactivated"
	EventCategoryPendingTransactionCreated                            EventCategory = "pending_transaction.created"
	EventCategoryPendingTransactionUpdated                            EventCategory = "pending_transaction.updated"
	EventCategoryRealTimeDecisionCardAuthorizationRequested           EventCategory = "real_time_decision.card_authorization_requested"
	EventCategoryRealTimeDecisionDigitalWalletTokenRequested          EventCategory = "real_time_decision.digital_wallet_token_requested"
	EventCategoryRealTimeDecisionDigitalWalletAuthenticationRequested EventCategory = "real_time_decision.digital_wallet_authentication_requested"
	EventCategoryRealTimePaymentsTransferCreated                      EventCategory = "real_time_payments_transfer.created"
	EventCategoryRealTimePaymentsTransferUpdated                      EventCategory = "real_time_payments_transfer.updated"
	EventCategoryRealTimePaymentsRequestForPaymentCreated             EventCategory = "real_time_payments_request_for_payment.created"
	EventCategoryRealTimePaymentsRequestForPaymentUpdated             EventCategory = "real_time_payments_request_for_payment.updated"
	EventCategoryTransactionCreated                                   EventCategory = "transaction.created"
	EventCategoryWireDrawdownRequestCreated                           EventCategory = "wire_drawdown_request.created"
	EventCategoryWireDrawdownRequestUpdated                           EventCategory = "wire_drawdown_request.updated"
	EventCategoryWireTransferCreated                                  EventCategory = "wire_transfer.created"
	EventCategoryWireTransferUpdated                                  EventCategory = "wire_transfer.updated"
)

// A constant representing the object's type. For this resource it will always be
// `event`.
type EventType string

const (
	EventTypeEvent EventType = "event"
)

type EventListParams struct {
	// Filter Events to those belonging to the object with the provided identifier.
	AssociatedObjectID param.Field[string]                   `query:"associated_object_id"`
	Category           param.Field[EventListParamsCategory]  `query:"category"`
	CreatedAt          param.Field[EventListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [EventListParams]'s query parameters as `url.Values`.
func (r EventListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type EventListParamsCategory struct {
	// Filter Events for those with the specified category or categories. For GET
	// requests, this should be encoded as a comma-delimited string, such as
	// `?in=one,two,three`.
	In param.Field[[]EventListParamsCategoryIn] `query:"in"`
}

// URLQuery serializes [EventListParamsCategory]'s query parameters as
// `url.Values`.
func (r EventListParamsCategory) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type EventListParamsCategoryIn string

const (
	EventListParamsCategoryInAccountCreated                                       EventListParamsCategoryIn = "account.created"
	EventListParamsCategoryInAccountUpdated                                       EventListParamsCategoryIn = "account.updated"
	EventListParamsCategoryInAccountNumberCreated                                 EventListParamsCategoryIn = "account_number.created"
	EventListParamsCategoryInAccountNumberUpdated                                 EventListParamsCategoryIn = "account_number.updated"
	EventListParamsCategoryInAccountStatementCreated                              EventListParamsCategoryIn = "account_statement.created"
	EventListParamsCategoryInAccountTransferCreated                               EventListParamsCategoryIn = "account_transfer.created"
	EventListParamsCategoryInAccountTransferUpdated                               EventListParamsCategoryIn = "account_transfer.updated"
	EventListParamsCategoryInACHPrenotificationCreated                            EventListParamsCategoryIn = "ach_prenotification.created"
	EventListParamsCategoryInACHPrenotificationUpdated                            EventListParamsCategoryIn = "ach_prenotification.updated"
	EventListParamsCategoryInACHTransferCreated                                   EventListParamsCategoryIn = "ach_transfer.created"
	EventListParamsCategoryInACHTransferUpdated                                   EventListParamsCategoryIn = "ach_transfer.updated"
	EventListParamsCategoryInCardCreated                                          EventListParamsCategoryIn = "card.created"
	EventListParamsCategoryInCardUpdated                                          EventListParamsCategoryIn = "card.updated"
	EventListParamsCategoryInCardPaymentCreated                                   EventListParamsCategoryIn = "card_payment.created"
	EventListParamsCategoryInCardPaymentUpdated                                   EventListParamsCategoryIn = "card_payment.updated"
	EventListParamsCategoryInCardDisputeCreated                                   EventListParamsCategoryIn = "card_dispute.created"
	EventListParamsCategoryInCardDisputeUpdated                                   EventListParamsCategoryIn = "card_dispute.updated"
	EventListParamsCategoryInCheckDepositCreated                                  EventListParamsCategoryIn = "check_deposit.created"
	EventListParamsCategoryInCheckDepositUpdated                                  EventListParamsCategoryIn = "check_deposit.updated"
	EventListParamsCategoryInCheckTransferCreated                                 EventListParamsCategoryIn = "check_transfer.created"
	EventListParamsCategoryInCheckTransferUpdated                                 EventListParamsCategoryIn = "check_transfer.updated"
	EventListParamsCategoryInDeclinedTransactionCreated                           EventListParamsCategoryIn = "declined_transaction.created"
	EventListParamsCategoryInDigitalWalletTokenCreated                            EventListParamsCategoryIn = "digital_wallet_token.created"
	EventListParamsCategoryInDigitalWalletTokenUpdated                            EventListParamsCategoryIn = "digital_wallet_token.updated"
	EventListParamsCategoryInDocumentCreated                                      EventListParamsCategoryIn = "document.created"
	EventListParamsCategoryInEntityCreated                                        EventListParamsCategoryIn = "entity.created"
	EventListParamsCategoryInEntityUpdated                                        EventListParamsCategoryIn = "entity.updated"
	EventListParamsCategoryInExternalAccountCreated                               EventListParamsCategoryIn = "external_account.created"
	EventListParamsCategoryInFileCreated                                          EventListParamsCategoryIn = "file.created"
	EventListParamsCategoryInGroupUpdated                                         EventListParamsCategoryIn = "group.updated"
	EventListParamsCategoryInGroupHeartbeat                                       EventListParamsCategoryIn = "group.heartbeat"
	EventListParamsCategoryInInboundACHTransferReturnCreated                      EventListParamsCategoryIn = "inbound_ach_transfer_return.created"
	EventListParamsCategoryInInboundACHTransferReturnUpdated                      EventListParamsCategoryIn = "inbound_ach_transfer_return.updated"
	EventListParamsCategoryInInboundWireDrawdownRequestCreated                    EventListParamsCategoryIn = "inbound_wire_drawdown_request.created"
	EventListParamsCategoryInOauthConnectionCreated                               EventListParamsCategoryIn = "oauth_connection.created"
	EventListParamsCategoryInOauthConnectionDeactivated                           EventListParamsCategoryIn = "oauth_connection.deactivated"
	EventListParamsCategoryInPendingTransactionCreated                            EventListParamsCategoryIn = "pending_transaction.created"
	EventListParamsCategoryInPendingTransactionUpdated                            EventListParamsCategoryIn = "pending_transaction.updated"
	EventListParamsCategoryInRealTimeDecisionCardAuthorizationRequested           EventListParamsCategoryIn = "real_time_decision.card_authorization_requested"
	EventListParamsCategoryInRealTimeDecisionDigitalWalletTokenRequested          EventListParamsCategoryIn = "real_time_decision.digital_wallet_token_requested"
	EventListParamsCategoryInRealTimeDecisionDigitalWalletAuthenticationRequested EventListParamsCategoryIn = "real_time_decision.digital_wallet_authentication_requested"
	EventListParamsCategoryInRealTimePaymentsTransferCreated                      EventListParamsCategoryIn = "real_time_payments_transfer.created"
	EventListParamsCategoryInRealTimePaymentsTransferUpdated                      EventListParamsCategoryIn = "real_time_payments_transfer.updated"
	EventListParamsCategoryInRealTimePaymentsRequestForPaymentCreated             EventListParamsCategoryIn = "real_time_payments_request_for_payment.created"
	EventListParamsCategoryInRealTimePaymentsRequestForPaymentUpdated             EventListParamsCategoryIn = "real_time_payments_request_for_payment.updated"
	EventListParamsCategoryInTransactionCreated                                   EventListParamsCategoryIn = "transaction.created"
	EventListParamsCategoryInWireDrawdownRequestCreated                           EventListParamsCategoryIn = "wire_drawdown_request.created"
	EventListParamsCategoryInWireDrawdownRequestUpdated                           EventListParamsCategoryIn = "wire_drawdown_request.updated"
	EventListParamsCategoryInWireTransferCreated                                  EventListParamsCategoryIn = "wire_transfer.created"
	EventListParamsCategoryInWireTransferUpdated                                  EventListParamsCategoryIn = "wire_transfer.updated"
)

type EventListParamsCreatedAt struct {
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

// URLQuery serializes [EventListParamsCreatedAt]'s query parameters as
// `url.Values`.
func (r EventListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// A list of Event objects
type EventListResponse struct {
	// The contents of the list.
	Data []Event `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       eventListResponseJSON
}

// eventListResponseJSON contains the JSON metadata for the struct
// [EventListResponse]
type eventListResponseJSON struct {
	Data        apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

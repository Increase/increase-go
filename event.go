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
	"github.com/Increase/increase-go/internal/param"
	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
	"github.com/Increase/increase-go/packages/pagination"
)

// EventService contains methods and other services that help with interacting with
// the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEventService] method instead.
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
	if eventID == "" {
		err = errors.New("missing required event_id parameter")
		return
	}
	path := fmt.Sprintf("events/%s", eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Events
func (r *EventService) List(ctx context.Context, query EventListParams, opts ...option.RequestOption) (res *pagination.Page[Event], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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
func (r *EventService) ListAutoPaging(ctx context.Context, query EventListParams, opts ...option.RequestOption) *pagination.PageAutoPager[Event] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Events are records of things that happened to objects at Increase. Events are
// accessible via the List Events endpoint and can be delivered to your application
// via webhooks. For more information, see our
// [webhooks guide](https://increase.com/documentation/webhooks).
type Event struct {
	// The Event identifier.
	ID string `json:"id,required"`
	// The identifier of the object that generated this Event.
	AssociatedObjectID string `json:"associated_object_id,required"`
	// The type of the object that generated this Event.
	AssociatedObjectType string `json:"associated_object_type,required"`
	// The category of the Event. We may add additional possible values for this enum
	// over time; your application should be able to handle such additions gracefully.
	Category EventCategory `json:"category,required"`
	// The time the Event was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A constant representing the object's type. For this resource it will always be
	// `event`.
	Type EventType `json:"type,required"`
	JSON eventJSON `json:"-"`
}

// eventJSON contains the JSON metadata for the struct [Event]
type eventJSON struct {
	ID                   apijson.Field
	AssociatedObjectID   apijson.Field
	AssociatedObjectType apijson.Field
	Category             apijson.Field
	CreatedAt            apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *Event) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventJSON) RawJSON() string {
	return r.raw
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
	EventCategoryBookkeepingAccountCreated                            EventCategory = "bookkeeping_account.created"
	EventCategoryBookkeepingAccountUpdated                            EventCategory = "bookkeeping_account.updated"
	EventCategoryBookkeepingEntrySetUpdated                           EventCategory = "bookkeeping_entry_set.updated"
	EventCategoryCardCreated                                          EventCategory = "card.created"
	EventCategoryCardUpdated                                          EventCategory = "card.updated"
	EventCategoryCardPaymentCreated                                   EventCategory = "card_payment.created"
	EventCategoryCardPaymentUpdated                                   EventCategory = "card_payment.updated"
	EventCategoryCardProfileCreated                                   EventCategory = "card_profile.created"
	EventCategoryCardProfileUpdated                                   EventCategory = "card_profile.updated"
	EventCategoryCardDisputeCreated                                   EventCategory = "card_dispute.created"
	EventCategoryCardDisputeUpdated                                   EventCategory = "card_dispute.updated"
	EventCategoryCheckDepositCreated                                  EventCategory = "check_deposit.created"
	EventCategoryCheckDepositUpdated                                  EventCategory = "check_deposit.updated"
	EventCategoryCheckTransferCreated                                 EventCategory = "check_transfer.created"
	EventCategoryCheckTransferUpdated                                 EventCategory = "check_transfer.updated"
	EventCategoryDeclinedTransactionCreated                           EventCategory = "declined_transaction.created"
	EventCategoryDigitalCardProfileCreated                            EventCategory = "digital_card_profile.created"
	EventCategoryDigitalCardProfileUpdated                            EventCategory = "digital_card_profile.updated"
	EventCategoryDigitalWalletTokenCreated                            EventCategory = "digital_wallet_token.created"
	EventCategoryDigitalWalletTokenUpdated                            EventCategory = "digital_wallet_token.updated"
	EventCategoryDocumentCreated                                      EventCategory = "document.created"
	EventCategoryEntityCreated                                        EventCategory = "entity.created"
	EventCategoryEntityUpdated                                        EventCategory = "entity.updated"
	EventCategoryEventSubscriptionCreated                             EventCategory = "event_subscription.created"
	EventCategoryEventSubscriptionUpdated                             EventCategory = "event_subscription.updated"
	EventCategoryExportCreated                                        EventCategory = "export.created"
	EventCategoryExportUpdated                                        EventCategory = "export.updated"
	EventCategoryExternalAccountCreated                               EventCategory = "external_account.created"
	EventCategoryExternalAccountUpdated                               EventCategory = "external_account.updated"
	EventCategoryFileCreated                                          EventCategory = "file.created"
	EventCategoryGroupUpdated                                         EventCategory = "group.updated"
	EventCategoryGroupHeartbeat                                       EventCategory = "group.heartbeat"
	EventCategoryInboundACHTransferCreated                            EventCategory = "inbound_ach_transfer.created"
	EventCategoryInboundACHTransferUpdated                            EventCategory = "inbound_ach_transfer.updated"
	EventCategoryInboundACHTransferReturnCreated                      EventCategory = "inbound_ach_transfer_return.created"
	EventCategoryInboundACHTransferReturnUpdated                      EventCategory = "inbound_ach_transfer_return.updated"
	EventCategoryInboundCheckDepositCreated                           EventCategory = "inbound_check_deposit.created"
	EventCategoryInboundCheckDepositUpdated                           EventCategory = "inbound_check_deposit.updated"
	EventCategoryInboundMailItemCreated                               EventCategory = "inbound_mail_item.created"
	EventCategoryInboundMailItemUpdated                               EventCategory = "inbound_mail_item.updated"
	EventCategoryInboundRealTimePaymentsTransferCreated               EventCategory = "inbound_real_time_payments_transfer.created"
	EventCategoryInboundRealTimePaymentsTransferUpdated               EventCategory = "inbound_real_time_payments_transfer.updated"
	EventCategoryInboundWireDrawdownRequestCreated                    EventCategory = "inbound_wire_drawdown_request.created"
	EventCategoryInboundWireTransferCreated                           EventCategory = "inbound_wire_transfer.created"
	EventCategoryInboundWireTransferUpdated                           EventCategory = "inbound_wire_transfer.updated"
	EventCategoryIntrafiAccountEnrollmentCreated                      EventCategory = "intrafi_account_enrollment.created"
	EventCategoryIntrafiAccountEnrollmentUpdated                      EventCategory = "intrafi_account_enrollment.updated"
	EventCategoryIntrafiExclusionCreated                              EventCategory = "intrafi_exclusion.created"
	EventCategoryIntrafiExclusionUpdated                              EventCategory = "intrafi_exclusion.updated"
	EventCategoryLockboxCreated                                       EventCategory = "lockbox.created"
	EventCategoryLockboxUpdated                                       EventCategory = "lockbox.updated"
	EventCategoryOAuthConnectionCreated                               EventCategory = "oauth_connection.created"
	EventCategoryOAuthConnectionDeactivated                           EventCategory = "oauth_connection.deactivated"
	EventCategoryOutboundCardPushTransferCreated                      EventCategory = "outbound_card_push_transfer.created"
	EventCategoryOutboundCardPushTransferUpdated                      EventCategory = "outbound_card_push_transfer.updated"
	EventCategoryOutboundCardValidationCreated                        EventCategory = "outbound_card_validation.created"
	EventCategoryOutboundCardValidationUpdated                        EventCategory = "outbound_card_validation.updated"
	EventCategoryPendingTransactionCreated                            EventCategory = "pending_transaction.created"
	EventCategoryPendingTransactionUpdated                            EventCategory = "pending_transaction.updated"
	EventCategoryPhysicalCardCreated                                  EventCategory = "physical_card.created"
	EventCategoryPhysicalCardUpdated                                  EventCategory = "physical_card.updated"
	EventCategoryPhysicalCardProfileCreated                           EventCategory = "physical_card_profile.created"
	EventCategoryPhysicalCardProfileUpdated                           EventCategory = "physical_card_profile.updated"
	EventCategoryProofOfAuthorizationRequestCreated                   EventCategory = "proof_of_authorization_request.created"
	EventCategoryProofOfAuthorizationRequestUpdated                   EventCategory = "proof_of_authorization_request.updated"
	EventCategoryRealTimeDecisionCardAuthorizationRequested           EventCategory = "real_time_decision.card_authorization_requested"
	EventCategoryRealTimeDecisionDigitalWalletTokenRequested          EventCategory = "real_time_decision.digital_wallet_token_requested"
	EventCategoryRealTimeDecisionDigitalWalletAuthenticationRequested EventCategory = "real_time_decision.digital_wallet_authentication_requested"
	EventCategoryRealTimeDecisionCardAuthenticationRequested          EventCategory = "real_time_decision.card_authentication_requested"
	EventCategoryRealTimeDecisionCardAuthenticationChallengeRequested EventCategory = "real_time_decision.card_authentication_challenge_requested"
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

func (r EventCategory) IsKnown() bool {
	switch r {
	case EventCategoryAccountCreated, EventCategoryAccountUpdated, EventCategoryAccountNumberCreated, EventCategoryAccountNumberUpdated, EventCategoryAccountStatementCreated, EventCategoryAccountTransferCreated, EventCategoryAccountTransferUpdated, EventCategoryACHPrenotificationCreated, EventCategoryACHPrenotificationUpdated, EventCategoryACHTransferCreated, EventCategoryACHTransferUpdated, EventCategoryBookkeepingAccountCreated, EventCategoryBookkeepingAccountUpdated, EventCategoryBookkeepingEntrySetUpdated, EventCategoryCardCreated, EventCategoryCardUpdated, EventCategoryCardPaymentCreated, EventCategoryCardPaymentUpdated, EventCategoryCardProfileCreated, EventCategoryCardProfileUpdated, EventCategoryCardDisputeCreated, EventCategoryCardDisputeUpdated, EventCategoryCheckDepositCreated, EventCategoryCheckDepositUpdated, EventCategoryCheckTransferCreated, EventCategoryCheckTransferUpdated, EventCategoryDeclinedTransactionCreated, EventCategoryDigitalCardProfileCreated, EventCategoryDigitalCardProfileUpdated, EventCategoryDigitalWalletTokenCreated, EventCategoryDigitalWalletTokenUpdated, EventCategoryDocumentCreated, EventCategoryEntityCreated, EventCategoryEntityUpdated, EventCategoryEventSubscriptionCreated, EventCategoryEventSubscriptionUpdated, EventCategoryExportCreated, EventCategoryExportUpdated, EventCategoryExternalAccountCreated, EventCategoryExternalAccountUpdated, EventCategoryFileCreated, EventCategoryGroupUpdated, EventCategoryGroupHeartbeat, EventCategoryInboundACHTransferCreated, EventCategoryInboundACHTransferUpdated, EventCategoryInboundACHTransferReturnCreated, EventCategoryInboundACHTransferReturnUpdated, EventCategoryInboundCheckDepositCreated, EventCategoryInboundCheckDepositUpdated, EventCategoryInboundMailItemCreated, EventCategoryInboundMailItemUpdated, EventCategoryInboundRealTimePaymentsTransferCreated, EventCategoryInboundRealTimePaymentsTransferUpdated, EventCategoryInboundWireDrawdownRequestCreated, EventCategoryInboundWireTransferCreated, EventCategoryInboundWireTransferUpdated, EventCategoryIntrafiAccountEnrollmentCreated, EventCategoryIntrafiAccountEnrollmentUpdated, EventCategoryIntrafiExclusionCreated, EventCategoryIntrafiExclusionUpdated, EventCategoryLockboxCreated, EventCategoryLockboxUpdated, EventCategoryOAuthConnectionCreated, EventCategoryOAuthConnectionDeactivated, EventCategoryOutboundCardPushTransferCreated, EventCategoryOutboundCardPushTransferUpdated, EventCategoryOutboundCardValidationCreated, EventCategoryOutboundCardValidationUpdated, EventCategoryPendingTransactionCreated, EventCategoryPendingTransactionUpdated, EventCategoryPhysicalCardCreated, EventCategoryPhysicalCardUpdated, EventCategoryPhysicalCardProfileCreated, EventCategoryPhysicalCardProfileUpdated, EventCategoryProofOfAuthorizationRequestCreated, EventCategoryProofOfAuthorizationRequestUpdated, EventCategoryRealTimeDecisionCardAuthorizationRequested, EventCategoryRealTimeDecisionDigitalWalletTokenRequested, EventCategoryRealTimeDecisionDigitalWalletAuthenticationRequested, EventCategoryRealTimeDecisionCardAuthenticationRequested, EventCategoryRealTimeDecisionCardAuthenticationChallengeRequested, EventCategoryRealTimePaymentsTransferCreated, EventCategoryRealTimePaymentsTransferUpdated, EventCategoryRealTimePaymentsRequestForPaymentCreated, EventCategoryRealTimePaymentsRequestForPaymentUpdated, EventCategoryTransactionCreated, EventCategoryWireDrawdownRequestCreated, EventCategoryWireDrawdownRequestUpdated, EventCategoryWireTransferCreated, EventCategoryWireTransferUpdated:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `event`.
type EventType string

const (
	EventTypeEvent EventType = "event"
)

func (r EventType) IsKnown() bool {
	switch r {
	case EventTypeEvent:
		return true
	}
	return false
}

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
	EventListParamsCategoryInBookkeepingAccountCreated                            EventListParamsCategoryIn = "bookkeeping_account.created"
	EventListParamsCategoryInBookkeepingAccountUpdated                            EventListParamsCategoryIn = "bookkeeping_account.updated"
	EventListParamsCategoryInBookkeepingEntrySetUpdated                           EventListParamsCategoryIn = "bookkeeping_entry_set.updated"
	EventListParamsCategoryInCardCreated                                          EventListParamsCategoryIn = "card.created"
	EventListParamsCategoryInCardUpdated                                          EventListParamsCategoryIn = "card.updated"
	EventListParamsCategoryInCardPaymentCreated                                   EventListParamsCategoryIn = "card_payment.created"
	EventListParamsCategoryInCardPaymentUpdated                                   EventListParamsCategoryIn = "card_payment.updated"
	EventListParamsCategoryInCardProfileCreated                                   EventListParamsCategoryIn = "card_profile.created"
	EventListParamsCategoryInCardProfileUpdated                                   EventListParamsCategoryIn = "card_profile.updated"
	EventListParamsCategoryInCardDisputeCreated                                   EventListParamsCategoryIn = "card_dispute.created"
	EventListParamsCategoryInCardDisputeUpdated                                   EventListParamsCategoryIn = "card_dispute.updated"
	EventListParamsCategoryInCheckDepositCreated                                  EventListParamsCategoryIn = "check_deposit.created"
	EventListParamsCategoryInCheckDepositUpdated                                  EventListParamsCategoryIn = "check_deposit.updated"
	EventListParamsCategoryInCheckTransferCreated                                 EventListParamsCategoryIn = "check_transfer.created"
	EventListParamsCategoryInCheckTransferUpdated                                 EventListParamsCategoryIn = "check_transfer.updated"
	EventListParamsCategoryInDeclinedTransactionCreated                           EventListParamsCategoryIn = "declined_transaction.created"
	EventListParamsCategoryInDigitalCardProfileCreated                            EventListParamsCategoryIn = "digital_card_profile.created"
	EventListParamsCategoryInDigitalCardProfileUpdated                            EventListParamsCategoryIn = "digital_card_profile.updated"
	EventListParamsCategoryInDigitalWalletTokenCreated                            EventListParamsCategoryIn = "digital_wallet_token.created"
	EventListParamsCategoryInDigitalWalletTokenUpdated                            EventListParamsCategoryIn = "digital_wallet_token.updated"
	EventListParamsCategoryInDocumentCreated                                      EventListParamsCategoryIn = "document.created"
	EventListParamsCategoryInEntityCreated                                        EventListParamsCategoryIn = "entity.created"
	EventListParamsCategoryInEntityUpdated                                        EventListParamsCategoryIn = "entity.updated"
	EventListParamsCategoryInEventSubscriptionCreated                             EventListParamsCategoryIn = "event_subscription.created"
	EventListParamsCategoryInEventSubscriptionUpdated                             EventListParamsCategoryIn = "event_subscription.updated"
	EventListParamsCategoryInExportCreated                                        EventListParamsCategoryIn = "export.created"
	EventListParamsCategoryInExportUpdated                                        EventListParamsCategoryIn = "export.updated"
	EventListParamsCategoryInExternalAccountCreated                               EventListParamsCategoryIn = "external_account.created"
	EventListParamsCategoryInExternalAccountUpdated                               EventListParamsCategoryIn = "external_account.updated"
	EventListParamsCategoryInFileCreated                                          EventListParamsCategoryIn = "file.created"
	EventListParamsCategoryInGroupUpdated                                         EventListParamsCategoryIn = "group.updated"
	EventListParamsCategoryInGroupHeartbeat                                       EventListParamsCategoryIn = "group.heartbeat"
	EventListParamsCategoryInInboundACHTransferCreated                            EventListParamsCategoryIn = "inbound_ach_transfer.created"
	EventListParamsCategoryInInboundACHTransferUpdated                            EventListParamsCategoryIn = "inbound_ach_transfer.updated"
	EventListParamsCategoryInInboundACHTransferReturnCreated                      EventListParamsCategoryIn = "inbound_ach_transfer_return.created"
	EventListParamsCategoryInInboundACHTransferReturnUpdated                      EventListParamsCategoryIn = "inbound_ach_transfer_return.updated"
	EventListParamsCategoryInInboundCheckDepositCreated                           EventListParamsCategoryIn = "inbound_check_deposit.created"
	EventListParamsCategoryInInboundCheckDepositUpdated                           EventListParamsCategoryIn = "inbound_check_deposit.updated"
	EventListParamsCategoryInInboundMailItemCreated                               EventListParamsCategoryIn = "inbound_mail_item.created"
	EventListParamsCategoryInInboundMailItemUpdated                               EventListParamsCategoryIn = "inbound_mail_item.updated"
	EventListParamsCategoryInInboundRealTimePaymentsTransferCreated               EventListParamsCategoryIn = "inbound_real_time_payments_transfer.created"
	EventListParamsCategoryInInboundRealTimePaymentsTransferUpdated               EventListParamsCategoryIn = "inbound_real_time_payments_transfer.updated"
	EventListParamsCategoryInInboundWireDrawdownRequestCreated                    EventListParamsCategoryIn = "inbound_wire_drawdown_request.created"
	EventListParamsCategoryInInboundWireTransferCreated                           EventListParamsCategoryIn = "inbound_wire_transfer.created"
	EventListParamsCategoryInInboundWireTransferUpdated                           EventListParamsCategoryIn = "inbound_wire_transfer.updated"
	EventListParamsCategoryInIntrafiAccountEnrollmentCreated                      EventListParamsCategoryIn = "intrafi_account_enrollment.created"
	EventListParamsCategoryInIntrafiAccountEnrollmentUpdated                      EventListParamsCategoryIn = "intrafi_account_enrollment.updated"
	EventListParamsCategoryInIntrafiExclusionCreated                              EventListParamsCategoryIn = "intrafi_exclusion.created"
	EventListParamsCategoryInIntrafiExclusionUpdated                              EventListParamsCategoryIn = "intrafi_exclusion.updated"
	EventListParamsCategoryInLockboxCreated                                       EventListParamsCategoryIn = "lockbox.created"
	EventListParamsCategoryInLockboxUpdated                                       EventListParamsCategoryIn = "lockbox.updated"
	EventListParamsCategoryInOAuthConnectionCreated                               EventListParamsCategoryIn = "oauth_connection.created"
	EventListParamsCategoryInOAuthConnectionDeactivated                           EventListParamsCategoryIn = "oauth_connection.deactivated"
	EventListParamsCategoryInOutboundCardPushTransferCreated                      EventListParamsCategoryIn = "outbound_card_push_transfer.created"
	EventListParamsCategoryInOutboundCardPushTransferUpdated                      EventListParamsCategoryIn = "outbound_card_push_transfer.updated"
	EventListParamsCategoryInOutboundCardValidationCreated                        EventListParamsCategoryIn = "outbound_card_validation.created"
	EventListParamsCategoryInOutboundCardValidationUpdated                        EventListParamsCategoryIn = "outbound_card_validation.updated"
	EventListParamsCategoryInPendingTransactionCreated                            EventListParamsCategoryIn = "pending_transaction.created"
	EventListParamsCategoryInPendingTransactionUpdated                            EventListParamsCategoryIn = "pending_transaction.updated"
	EventListParamsCategoryInPhysicalCardCreated                                  EventListParamsCategoryIn = "physical_card.created"
	EventListParamsCategoryInPhysicalCardUpdated                                  EventListParamsCategoryIn = "physical_card.updated"
	EventListParamsCategoryInPhysicalCardProfileCreated                           EventListParamsCategoryIn = "physical_card_profile.created"
	EventListParamsCategoryInPhysicalCardProfileUpdated                           EventListParamsCategoryIn = "physical_card_profile.updated"
	EventListParamsCategoryInProofOfAuthorizationRequestCreated                   EventListParamsCategoryIn = "proof_of_authorization_request.created"
	EventListParamsCategoryInProofOfAuthorizationRequestUpdated                   EventListParamsCategoryIn = "proof_of_authorization_request.updated"
	EventListParamsCategoryInRealTimeDecisionCardAuthorizationRequested           EventListParamsCategoryIn = "real_time_decision.card_authorization_requested"
	EventListParamsCategoryInRealTimeDecisionDigitalWalletTokenRequested          EventListParamsCategoryIn = "real_time_decision.digital_wallet_token_requested"
	EventListParamsCategoryInRealTimeDecisionDigitalWalletAuthenticationRequested EventListParamsCategoryIn = "real_time_decision.digital_wallet_authentication_requested"
	EventListParamsCategoryInRealTimeDecisionCardAuthenticationRequested          EventListParamsCategoryIn = "real_time_decision.card_authentication_requested"
	EventListParamsCategoryInRealTimeDecisionCardAuthenticationChallengeRequested EventListParamsCategoryIn = "real_time_decision.card_authentication_challenge_requested"
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

func (r EventListParamsCategoryIn) IsKnown() bool {
	switch r {
	case EventListParamsCategoryInAccountCreated, EventListParamsCategoryInAccountUpdated, EventListParamsCategoryInAccountNumberCreated, EventListParamsCategoryInAccountNumberUpdated, EventListParamsCategoryInAccountStatementCreated, EventListParamsCategoryInAccountTransferCreated, EventListParamsCategoryInAccountTransferUpdated, EventListParamsCategoryInACHPrenotificationCreated, EventListParamsCategoryInACHPrenotificationUpdated, EventListParamsCategoryInACHTransferCreated, EventListParamsCategoryInACHTransferUpdated, EventListParamsCategoryInBookkeepingAccountCreated, EventListParamsCategoryInBookkeepingAccountUpdated, EventListParamsCategoryInBookkeepingEntrySetUpdated, EventListParamsCategoryInCardCreated, EventListParamsCategoryInCardUpdated, EventListParamsCategoryInCardPaymentCreated, EventListParamsCategoryInCardPaymentUpdated, EventListParamsCategoryInCardProfileCreated, EventListParamsCategoryInCardProfileUpdated, EventListParamsCategoryInCardDisputeCreated, EventListParamsCategoryInCardDisputeUpdated, EventListParamsCategoryInCheckDepositCreated, EventListParamsCategoryInCheckDepositUpdated, EventListParamsCategoryInCheckTransferCreated, EventListParamsCategoryInCheckTransferUpdated, EventListParamsCategoryInDeclinedTransactionCreated, EventListParamsCategoryInDigitalCardProfileCreated, EventListParamsCategoryInDigitalCardProfileUpdated, EventListParamsCategoryInDigitalWalletTokenCreated, EventListParamsCategoryInDigitalWalletTokenUpdated, EventListParamsCategoryInDocumentCreated, EventListParamsCategoryInEntityCreated, EventListParamsCategoryInEntityUpdated, EventListParamsCategoryInEventSubscriptionCreated, EventListParamsCategoryInEventSubscriptionUpdated, EventListParamsCategoryInExportCreated, EventListParamsCategoryInExportUpdated, EventListParamsCategoryInExternalAccountCreated, EventListParamsCategoryInExternalAccountUpdated, EventListParamsCategoryInFileCreated, EventListParamsCategoryInGroupUpdated, EventListParamsCategoryInGroupHeartbeat, EventListParamsCategoryInInboundACHTransferCreated, EventListParamsCategoryInInboundACHTransferUpdated, EventListParamsCategoryInInboundACHTransferReturnCreated, EventListParamsCategoryInInboundACHTransferReturnUpdated, EventListParamsCategoryInInboundCheckDepositCreated, EventListParamsCategoryInInboundCheckDepositUpdated, EventListParamsCategoryInInboundMailItemCreated, EventListParamsCategoryInInboundMailItemUpdated, EventListParamsCategoryInInboundRealTimePaymentsTransferCreated, EventListParamsCategoryInInboundRealTimePaymentsTransferUpdated, EventListParamsCategoryInInboundWireDrawdownRequestCreated, EventListParamsCategoryInInboundWireTransferCreated, EventListParamsCategoryInInboundWireTransferUpdated, EventListParamsCategoryInIntrafiAccountEnrollmentCreated, EventListParamsCategoryInIntrafiAccountEnrollmentUpdated, EventListParamsCategoryInIntrafiExclusionCreated, EventListParamsCategoryInIntrafiExclusionUpdated, EventListParamsCategoryInLockboxCreated, EventListParamsCategoryInLockboxUpdated, EventListParamsCategoryInOAuthConnectionCreated, EventListParamsCategoryInOAuthConnectionDeactivated, EventListParamsCategoryInOutboundCardPushTransferCreated, EventListParamsCategoryInOutboundCardPushTransferUpdated, EventListParamsCategoryInOutboundCardValidationCreated, EventListParamsCategoryInOutboundCardValidationUpdated, EventListParamsCategoryInPendingTransactionCreated, EventListParamsCategoryInPendingTransactionUpdated, EventListParamsCategoryInPhysicalCardCreated, EventListParamsCategoryInPhysicalCardUpdated, EventListParamsCategoryInPhysicalCardProfileCreated, EventListParamsCategoryInPhysicalCardProfileUpdated, EventListParamsCategoryInProofOfAuthorizationRequestCreated, EventListParamsCategoryInProofOfAuthorizationRequestUpdated, EventListParamsCategoryInRealTimeDecisionCardAuthorizationRequested, EventListParamsCategoryInRealTimeDecisionDigitalWalletTokenRequested, EventListParamsCategoryInRealTimeDecisionDigitalWalletAuthenticationRequested, EventListParamsCategoryInRealTimeDecisionCardAuthenticationRequested, EventListParamsCategoryInRealTimeDecisionCardAuthenticationChallengeRequested, EventListParamsCategoryInRealTimePaymentsTransferCreated, EventListParamsCategoryInRealTimePaymentsTransferUpdated, EventListParamsCategoryInRealTimePaymentsRequestForPaymentCreated, EventListParamsCategoryInRealTimePaymentsRequestForPaymentUpdated, EventListParamsCategoryInTransactionCreated, EventListParamsCategoryInWireDrawdownRequestCreated, EventListParamsCategoryInWireDrawdownRequestUpdated, EventListParamsCategoryInWireTransferCreated, EventListParamsCategoryInWireTransferUpdated:
		return true
	}
	return false
}

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

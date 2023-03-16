package types

import (
	"fmt"
	"net/url"
	"time"

	"github.com/increase/increase-go/core"
	"github.com/increase/increase-go/core/pjson"
	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/pagination"
)

type EventSubscription struct {
	// The event subscription identifier.
	ID *string `pjson:"id"`
	// The time the event subscription was created.
	CreatedAt *time.Time `pjson:"created_at" format:"2006-01-02T15:04:05Z07:00"`
	// This indicates if we'll send notifications to this subscription.
	Status *EventSubscriptionStatus `pjson:"status"`
	// If specified, this subscription will only receive webhooks for Events with the
	// specified `category`.
	SelectedEventCategory *EventSubscriptionSelectedEventCategory `pjson:"selected_event_category"`
	// The webhook url where we'll send notifications.
	URL *string `pjson:"url"`
	// The key that will be used to sign webhooks.
	SharedSecret *string `pjson:"shared_secret"`
	// A constant representing the object's type. For this resource it will always be
	// `event_subscription`.
	Type       *EventSubscriptionType `pjson:"type"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into EventSubscription using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *EventSubscription) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes EventSubscription into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *EventSubscription) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The event subscription identifier.
func (r EventSubscription) GetID() (ID string) {
	if r.ID != nil {
		ID = *r.ID
	}
	return
}

// The time the event subscription was created.
func (r EventSubscription) GetCreatedAt() (CreatedAt time.Time) {
	if r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// This indicates if we'll send notifications to this subscription.
func (r EventSubscription) GetStatus() (Status EventSubscriptionStatus) {
	if r.Status != nil {
		Status = *r.Status
	}
	return
}

// If specified, this subscription will only receive webhooks for Events with the
// specified `category`.
func (r EventSubscription) GetSelectedEventCategory() (SelectedEventCategory EventSubscriptionSelectedEventCategory) {
	if r.SelectedEventCategory != nil {
		SelectedEventCategory = *r.SelectedEventCategory
	}
	return
}

// The webhook url where we'll send notifications.
func (r EventSubscription) GetURL() (URL string) {
	if r.URL != nil {
		URL = *r.URL
	}
	return
}

// The key that will be used to sign webhooks.
func (r EventSubscription) GetSharedSecret() (SharedSecret string) {
	if r.SharedSecret != nil {
		SharedSecret = *r.SharedSecret
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `event_subscription`.
func (r EventSubscription) GetType() (Type EventSubscriptionType) {
	if r.Type != nil {
		Type = *r.Type
	}
	return
}

func (r EventSubscription) String() (result string) {
	return fmt.Sprintf("&EventSubscription{ID:%s CreatedAt:%s Status:%s SelectedEventCategory:%s URL:%s SharedSecret:%s Type:%s}", core.FmtP(r.ID), core.FmtP(r.CreatedAt), core.FmtP(r.Status), core.FmtP(r.SelectedEventCategory), core.FmtP(r.URL), core.FmtP(r.SharedSecret), core.FmtP(r.Type))
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

type CreateAnEventSubscriptionParameters struct {
	// The URL you'd like us to send webhooks to.
	URL *string `pjson:"url"`
	// The key that will be used to sign webhooks. If no value is passed, a random
	// string will be used as default.
	SharedSecret *string `pjson:"shared_secret"`
	// If specified, this subscription will only receive webhooks for Events with the
	// specified `category`.
	SelectedEventCategory *CreateAnEventSubscriptionParametersSelectedEventCategory `pjson:"selected_event_category"`
	jsonFields            map[string]interface{}                                    `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateAnEventSubscriptionParameters using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *CreateAnEventSubscriptionParameters) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CreateAnEventSubscriptionParameters into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *CreateAnEventSubscriptionParameters) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The URL you'd like us to send webhooks to.
func (r CreateAnEventSubscriptionParameters) GetURL() (URL string) {
	if r.URL != nil {
		URL = *r.URL
	}
	return
}

// The key that will be used to sign webhooks. If no value is passed, a random
// string will be used as default.
func (r CreateAnEventSubscriptionParameters) GetSharedSecret() (SharedSecret string) {
	if r.SharedSecret != nil {
		SharedSecret = *r.SharedSecret
	}
	return
}

// If specified, this subscription will only receive webhooks for Events with the
// specified `category`.
func (r CreateAnEventSubscriptionParameters) GetSelectedEventCategory() (SelectedEventCategory CreateAnEventSubscriptionParametersSelectedEventCategory) {
	if r.SelectedEventCategory != nil {
		SelectedEventCategory = *r.SelectedEventCategory
	}
	return
}

func (r CreateAnEventSubscriptionParameters) String() (result string) {
	return fmt.Sprintf("&CreateAnEventSubscriptionParameters{URL:%s SharedSecret:%s SelectedEventCategory:%s}", core.FmtP(r.URL), core.FmtP(r.SharedSecret), core.FmtP(r.SelectedEventCategory))
}

type CreateAnEventSubscriptionParametersSelectedEventCategory string

const (
	CreateAnEventSubscriptionParametersSelectedEventCategoryAccountCreated                                       CreateAnEventSubscriptionParametersSelectedEventCategory = "account.created"
	CreateAnEventSubscriptionParametersSelectedEventCategoryAccountUpdated                                       CreateAnEventSubscriptionParametersSelectedEventCategory = "account.updated"
	CreateAnEventSubscriptionParametersSelectedEventCategoryAccountNumberCreated                                 CreateAnEventSubscriptionParametersSelectedEventCategory = "account_number.created"
	CreateAnEventSubscriptionParametersSelectedEventCategoryAccountNumberUpdated                                 CreateAnEventSubscriptionParametersSelectedEventCategory = "account_number.updated"
	CreateAnEventSubscriptionParametersSelectedEventCategoryAccountStatementCreated                              CreateAnEventSubscriptionParametersSelectedEventCategory = "account_statement.created"
	CreateAnEventSubscriptionParametersSelectedEventCategoryAccountTransferCreated                               CreateAnEventSubscriptionParametersSelectedEventCategory = "account_transfer.created"
	CreateAnEventSubscriptionParametersSelectedEventCategoryAccountTransferUpdated                               CreateAnEventSubscriptionParametersSelectedEventCategory = "account_transfer.updated"
	CreateAnEventSubscriptionParametersSelectedEventCategoryACHPrenotificationCreated                            CreateAnEventSubscriptionParametersSelectedEventCategory = "ach_prenotification.created"
	CreateAnEventSubscriptionParametersSelectedEventCategoryACHPrenotificationUpdated                            CreateAnEventSubscriptionParametersSelectedEventCategory = "ach_prenotification.updated"
	CreateAnEventSubscriptionParametersSelectedEventCategoryACHTransferCreated                                   CreateAnEventSubscriptionParametersSelectedEventCategory = "ach_transfer.created"
	CreateAnEventSubscriptionParametersSelectedEventCategoryACHTransferUpdated                                   CreateAnEventSubscriptionParametersSelectedEventCategory = "ach_transfer.updated"
	CreateAnEventSubscriptionParametersSelectedEventCategoryCardCreated                                          CreateAnEventSubscriptionParametersSelectedEventCategory = "card.created"
	CreateAnEventSubscriptionParametersSelectedEventCategoryCardUpdated                                          CreateAnEventSubscriptionParametersSelectedEventCategory = "card.updated"
	CreateAnEventSubscriptionParametersSelectedEventCategoryCardDisputeCreated                                   CreateAnEventSubscriptionParametersSelectedEventCategory = "card_dispute.created"
	CreateAnEventSubscriptionParametersSelectedEventCategoryCardDisputeUpdated                                   CreateAnEventSubscriptionParametersSelectedEventCategory = "card_dispute.updated"
	CreateAnEventSubscriptionParametersSelectedEventCategoryCheckDepositCreated                                  CreateAnEventSubscriptionParametersSelectedEventCategory = "check_deposit.created"
	CreateAnEventSubscriptionParametersSelectedEventCategoryCheckDepositUpdated                                  CreateAnEventSubscriptionParametersSelectedEventCategory = "check_deposit.updated"
	CreateAnEventSubscriptionParametersSelectedEventCategoryCheckTransferCreated                                 CreateAnEventSubscriptionParametersSelectedEventCategory = "check_transfer.created"
	CreateAnEventSubscriptionParametersSelectedEventCategoryCheckTransferUpdated                                 CreateAnEventSubscriptionParametersSelectedEventCategory = "check_transfer.updated"
	CreateAnEventSubscriptionParametersSelectedEventCategoryDeclinedTransactionCreated                           CreateAnEventSubscriptionParametersSelectedEventCategory = "declined_transaction.created"
	CreateAnEventSubscriptionParametersSelectedEventCategoryDigitalWalletTokenCreated                            CreateAnEventSubscriptionParametersSelectedEventCategory = "digital_wallet_token.created"
	CreateAnEventSubscriptionParametersSelectedEventCategoryDigitalWalletTokenUpdated                            CreateAnEventSubscriptionParametersSelectedEventCategory = "digital_wallet_token.updated"
	CreateAnEventSubscriptionParametersSelectedEventCategoryDocumentCreated                                      CreateAnEventSubscriptionParametersSelectedEventCategory = "document.created"
	CreateAnEventSubscriptionParametersSelectedEventCategoryEntityCreated                                        CreateAnEventSubscriptionParametersSelectedEventCategory = "entity.created"
	CreateAnEventSubscriptionParametersSelectedEventCategoryEntityUpdated                                        CreateAnEventSubscriptionParametersSelectedEventCategory = "entity.updated"
	CreateAnEventSubscriptionParametersSelectedEventCategoryExternalAccountCreated                               CreateAnEventSubscriptionParametersSelectedEventCategory = "external_account.created"
	CreateAnEventSubscriptionParametersSelectedEventCategoryFileCreated                                          CreateAnEventSubscriptionParametersSelectedEventCategory = "file.created"
	CreateAnEventSubscriptionParametersSelectedEventCategoryGroupUpdated                                         CreateAnEventSubscriptionParametersSelectedEventCategory = "group.updated"
	CreateAnEventSubscriptionParametersSelectedEventCategoryGroupHeartbeat                                       CreateAnEventSubscriptionParametersSelectedEventCategory = "group.heartbeat"
	CreateAnEventSubscriptionParametersSelectedEventCategoryInboundACHTransferReturnCreated                      CreateAnEventSubscriptionParametersSelectedEventCategory = "inbound_ach_transfer_return.created"
	CreateAnEventSubscriptionParametersSelectedEventCategoryInboundACHTransferReturnUpdated                      CreateAnEventSubscriptionParametersSelectedEventCategory = "inbound_ach_transfer_return.updated"
	CreateAnEventSubscriptionParametersSelectedEventCategoryInboundWireDrawdownRequestCreated                    CreateAnEventSubscriptionParametersSelectedEventCategory = "inbound_wire_drawdown_request.created"
	CreateAnEventSubscriptionParametersSelectedEventCategoryOauthConnectionCreated                               CreateAnEventSubscriptionParametersSelectedEventCategory = "oauth_connection.created"
	CreateAnEventSubscriptionParametersSelectedEventCategoryOauthConnectionDeactivated                           CreateAnEventSubscriptionParametersSelectedEventCategory = "oauth_connection.deactivated"
	CreateAnEventSubscriptionParametersSelectedEventCategoryPendingTransactionCreated                            CreateAnEventSubscriptionParametersSelectedEventCategory = "pending_transaction.created"
	CreateAnEventSubscriptionParametersSelectedEventCategoryPendingTransactionUpdated                            CreateAnEventSubscriptionParametersSelectedEventCategory = "pending_transaction.updated"
	CreateAnEventSubscriptionParametersSelectedEventCategoryRealTimeDecisionCardAuthorizationRequested           CreateAnEventSubscriptionParametersSelectedEventCategory = "real_time_decision.card_authorization_requested"
	CreateAnEventSubscriptionParametersSelectedEventCategoryRealTimeDecisionDigitalWalletTokenRequested          CreateAnEventSubscriptionParametersSelectedEventCategory = "real_time_decision.digital_wallet_token_requested"
	CreateAnEventSubscriptionParametersSelectedEventCategoryRealTimeDecisionDigitalWalletAuthenticationRequested CreateAnEventSubscriptionParametersSelectedEventCategory = "real_time_decision.digital_wallet_authentication_requested"
	CreateAnEventSubscriptionParametersSelectedEventCategoryRealTimePaymentsTransferCreated                      CreateAnEventSubscriptionParametersSelectedEventCategory = "real_time_payments_transfer.created"
	CreateAnEventSubscriptionParametersSelectedEventCategoryRealTimePaymentsTransferUpdated                      CreateAnEventSubscriptionParametersSelectedEventCategory = "real_time_payments_transfer.updated"
	CreateAnEventSubscriptionParametersSelectedEventCategoryRealTimePaymentsRequestForPaymentCreated             CreateAnEventSubscriptionParametersSelectedEventCategory = "real_time_payments_request_for_payment.created"
	CreateAnEventSubscriptionParametersSelectedEventCategoryRealTimePaymentsRequestForPaymentUpdated             CreateAnEventSubscriptionParametersSelectedEventCategory = "real_time_payments_request_for_payment.updated"
	CreateAnEventSubscriptionParametersSelectedEventCategoryTransactionCreated                                   CreateAnEventSubscriptionParametersSelectedEventCategory = "transaction.created"
	CreateAnEventSubscriptionParametersSelectedEventCategoryWireDrawdownRequestCreated                           CreateAnEventSubscriptionParametersSelectedEventCategory = "wire_drawdown_request.created"
	CreateAnEventSubscriptionParametersSelectedEventCategoryWireDrawdownRequestUpdated                           CreateAnEventSubscriptionParametersSelectedEventCategory = "wire_drawdown_request.updated"
	CreateAnEventSubscriptionParametersSelectedEventCategoryWireTransferCreated                                  CreateAnEventSubscriptionParametersSelectedEventCategory = "wire_transfer.created"
	CreateAnEventSubscriptionParametersSelectedEventCategoryWireTransferUpdated                                  CreateAnEventSubscriptionParametersSelectedEventCategory = "wire_transfer.updated"
)

type UpdateAnEventSubscriptionParameters struct {
	// The status to update the Event Subscription with.
	Status     *UpdateAnEventSubscriptionParametersStatus `pjson:"status"`
	jsonFields map[string]interface{}                     `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// UpdateAnEventSubscriptionParameters using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *UpdateAnEventSubscriptionParameters) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes UpdateAnEventSubscriptionParameters into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *UpdateAnEventSubscriptionParameters) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The status to update the Event Subscription with.
func (r UpdateAnEventSubscriptionParameters) GetStatus() (Status UpdateAnEventSubscriptionParametersStatus) {
	if r.Status != nil {
		Status = *r.Status
	}
	return
}

func (r UpdateAnEventSubscriptionParameters) String() (result string) {
	return fmt.Sprintf("&UpdateAnEventSubscriptionParameters{Status:%s}", core.FmtP(r.Status))
}

type UpdateAnEventSubscriptionParametersStatus string

const (
	UpdateAnEventSubscriptionParametersStatusActive   UpdateAnEventSubscriptionParametersStatus = "active"
	UpdateAnEventSubscriptionParametersStatusDisabled UpdateAnEventSubscriptionParametersStatus = "disabled"
	UpdateAnEventSubscriptionParametersStatusDeleted  UpdateAnEventSubscriptionParametersStatus = "deleted"
)

type EventSubscriptionListParams struct {
	// Return the page of entries after this one.
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit      *int64                 `query:"limit"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into EventSubscriptionListParams
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *EventSubscriptionListParams) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes EventSubscriptionListParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *EventSubscriptionListParams) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes EventSubscriptionListParams into a url.Values of the query
// parameters associated with this value
func (r *EventSubscriptionListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// Return the page of entries after this one.
func (r EventSubscriptionListParams) GetCursor() (Cursor string) {
	if r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r EventSubscriptionListParams) GetLimit() (Limit int64) {
	if r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

func (r EventSubscriptionListParams) String() (result string) {
	return fmt.Sprintf("&EventSubscriptionListParams{Cursor:%s Limit:%s}", core.FmtP(r.Cursor), core.FmtP(r.Limit))
}

type EventSubscriptionList struct {
	// The contents of the list.
	Data *[]EventSubscription `pjson:"data"`
	// A pointer to a place in the list.
	NextCursor *string                `pjson:"next_cursor"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into EventSubscriptionList using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *EventSubscriptionList) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes EventSubscriptionList into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *EventSubscriptionList) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes EventSubscriptionList into a url.Values of the query
// parameters associated with this value
func (r *EventSubscriptionList) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// The contents of the list.
func (r EventSubscriptionList) GetData() (Data []EventSubscription) {
	if r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r EventSubscriptionList) GetNextCursor() (NextCursor string) {
	if r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

func (r EventSubscriptionList) String() (result string) {
	return fmt.Sprintf("&EventSubscriptionList{Data:%s NextCursor:%s}", core.Fmt(r.Data), core.FmtP(r.NextCursor))
}

type EventSubscriptionsPage struct {
	*pagination.Page[EventSubscription]
}

func (r *EventSubscriptionsPage) EventSubscription() *EventSubscription {
	return r.Current()
}

func (r *EventSubscriptionsPage) NextPage() (*EventSubscriptionsPage, error) {
	if page, err := r.Page.NextPage(); err != nil {
		return nil, err
	} else {
		return &EventSubscriptionsPage{page}, nil
	}
}

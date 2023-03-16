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

type Event struct {
	// The identifier of the object that generated this Event.
	AssociatedObjectID *string `pjson:"associated_object_id"`
	// The type of the object that generated this Event.
	AssociatedObjectType *string `pjson:"associated_object_type"`
	// The category of the Event. We may add additional possible values for this enum
	// over time; your application should be able to handle such additions gracefully.
	Category *EventCategory `pjson:"category"`
	// The time the Event was created.
	CreatedAt *time.Time `pjson:"created_at" format:"2006-01-02T15:04:05Z07:00"`
	// The Event identifier.
	ID *string `pjson:"id"`
	// A constant representing the object's type. For this resource it will always be
	// `event`.
	Type       *EventType             `pjson:"type"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into Event using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Event) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes Event into an array of bytes using the gjson library.
// Members of the `jsonFields` field are serialized into the top-level, and will
// overwrite known members of the same name.
func (r *Event) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The identifier of the object that generated this Event.
func (r Event) GetAssociatedObjectID() (AssociatedObjectID string) {
	if r.AssociatedObjectID != nil {
		AssociatedObjectID = *r.AssociatedObjectID
	}
	return
}

// The type of the object that generated this Event.
func (r Event) GetAssociatedObjectType() (AssociatedObjectType string) {
	if r.AssociatedObjectType != nil {
		AssociatedObjectType = *r.AssociatedObjectType
	}
	return
}

// The category of the Event. We may add additional possible values for this enum
// over time; your application should be able to handle such additions gracefully.
func (r Event) GetCategory() (Category EventCategory) {
	if r.Category != nil {
		Category = *r.Category
	}
	return
}

// The time the Event was created.
func (r Event) GetCreatedAt() (CreatedAt time.Time) {
	if r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The Event identifier.
func (r Event) GetID() (ID string) {
	if r.ID != nil {
		ID = *r.ID
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `event`.
func (r Event) GetType() (Type EventType) {
	if r.Type != nil {
		Type = *r.Type
	}
	return
}

func (r Event) String() (result string) {
	return fmt.Sprintf("&Event{AssociatedObjectID:%s AssociatedObjectType:%s Category:%s CreatedAt:%s ID:%s Type:%s}", core.FmtP(r.AssociatedObjectID), core.FmtP(r.AssociatedObjectType), core.FmtP(r.Category), core.FmtP(r.CreatedAt), core.FmtP(r.ID), core.FmtP(r.Type))
}

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

type EventType string

const (
	EventTypeEvent EventType = "event"
)

type EventListParams struct {
	// Return the page of entries after this one.
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int64 `query:"limit"`
	// Filter Events to those belonging to the object with the provided identifier.
	AssociatedObjectID *string                   `query:"associated_object_id"`
	CreatedAt          *EventListParamsCreatedAt `query:"created_at"`
	Category           *EventListParamsCategory  `query:"category"`
	jsonFields         map[string]interface{}    `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into EventListParams using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *EventListParams) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes EventListParams into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *EventListParams) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes EventListParams into a url.Values of the query parameters
// associated with this value
func (r *EventListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// Return the page of entries after this one.
func (r EventListParams) GetCursor() (Cursor string) {
	if r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r EventListParams) GetLimit() (Limit int64) {
	if r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// Filter Events to those belonging to the object with the provided identifier.
func (r EventListParams) GetAssociatedObjectID() (AssociatedObjectID string) {
	if r.AssociatedObjectID != nil {
		AssociatedObjectID = *r.AssociatedObjectID
	}
	return
}

func (r EventListParams) GetCreatedAt() (CreatedAt EventListParamsCreatedAt) {
	if r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

func (r EventListParams) GetCategory() (Category EventListParamsCategory) {
	if r.Category != nil {
		Category = *r.Category
	}
	return
}

func (r EventListParams) String() (result string) {
	return fmt.Sprintf("&EventListParams{Cursor:%s Limit:%s AssociatedObjectID:%s CreatedAt:%s Category:%s}", core.FmtP(r.Cursor), core.FmtP(r.Limit), core.FmtP(r.AssociatedObjectID), r.CreatedAt, r.Category)
}

type EventListParamsCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After *time.Time `pjson:"after" format:"2006-01-02T15:04:05Z07:00"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before *time.Time `pjson:"before" format:"2006-01-02T15:04:05Z07:00"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter *time.Time `pjson:"on_or_after" format:"2006-01-02T15:04:05Z07:00"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore *time.Time             `pjson:"on_or_before" format:"2006-01-02T15:04:05Z07:00"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into EventListParamsCreatedAt
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *EventListParamsCreatedAt) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes EventListParamsCreatedAt into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *EventListParamsCreatedAt) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes EventListParamsCreatedAt into a url.Values of the query
// parameters associated with this value
func (r *EventListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r EventListParamsCreatedAt) GetAfter() (After time.Time) {
	if r.After != nil {
		After = *r.After
	}
	return
}

// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r EventListParamsCreatedAt) GetBefore() (Before time.Time) {
	if r.Before != nil {
		Before = *r.Before
	}
	return
}

// Return results on or after this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r EventListParamsCreatedAt) GetOnOrAfter() (OnOrAfter time.Time) {
	if r.OnOrAfter != nil {
		OnOrAfter = *r.OnOrAfter
	}
	return
}

// Return results on or before this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r EventListParamsCreatedAt) GetOnOrBefore() (OnOrBefore time.Time) {
	if r.OnOrBefore != nil {
		OnOrBefore = *r.OnOrBefore
	}
	return
}

func (r EventListParamsCreatedAt) String() (result string) {
	return fmt.Sprintf("&EventListParamsCreatedAt{After:%s Before:%s OnOrAfter:%s OnOrBefore:%s}", core.FmtP(r.After), core.FmtP(r.Before), core.FmtP(r.OnOrAfter), core.FmtP(r.OnOrBefore))
}

type EventListParamsCategory struct {
	// Return results whose value is in the provided list. For GET requests, this
	// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
	In         *[]EventListParamsCategoryIn `pjson:"in"`
	jsonFields map[string]interface{}       `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into EventListParamsCategory using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *EventListParamsCategory) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes EventListParamsCategory into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *EventListParamsCategory) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes EventListParamsCategory into a url.Values of the query
// parameters associated with this value
func (r *EventListParamsCategory) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// Return results whose value is in the provided list. For GET requests, this
// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
func (r EventListParamsCategory) GetIn() (In []EventListParamsCategoryIn) {
	if r.In != nil {
		In = *r.In
	}
	return
}

func (r EventListParamsCategory) String() (result string) {
	return fmt.Sprintf("&EventListParamsCategory{In:%s}", core.Fmt(r.In))
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

type EventList struct {
	// The contents of the list.
	Data *[]Event `pjson:"data"`
	// A pointer to a place in the list.
	NextCursor *string                `pjson:"next_cursor"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into EventList using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *EventList) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes EventList into an array of bytes using the gjson library.
// Members of the `jsonFields` field are serialized into the top-level, and will
// overwrite known members of the same name.
func (r *EventList) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes EventList into a url.Values of the query parameters
// associated with this value
func (r *EventList) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// The contents of the list.
func (r EventList) GetData() (Data []Event) {
	if r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r EventList) GetNextCursor() (NextCursor string) {
	if r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

func (r EventList) String() (result string) {
	return fmt.Sprintf("&EventList{Data:%s NextCursor:%s}", core.Fmt(r.Data), core.FmtP(r.NextCursor))
}

type EventsPage struct {
	*pagination.Page[Event]
}

func (r *EventsPage) Event() *Event {
	return r.Current()
}

func (r *EventsPage) NextPage() (*EventsPage, error) {
	if page, err := r.Page.NextPage(); err != nil {
		return nil, err
	} else {
		return &EventsPage{page}, nil
	}
}

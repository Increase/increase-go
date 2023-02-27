package types

import (
	"fmt"
	"increase/core"
	"increase/core/pjson"
	"increase/core/query"
	"increase/pagination"
	"net/url"
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
	CreatedAt *string `pjson:"created_at"`
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
func (r *Event) GetAssociatedObjectID() (AssociatedObjectID string) {
	if r != nil && r.AssociatedObjectID != nil {
		AssociatedObjectID = *r.AssociatedObjectID
	}
	return
}

// The type of the object that generated this Event.
func (r *Event) GetAssociatedObjectType() (AssociatedObjectType string) {
	if r != nil && r.AssociatedObjectType != nil {
		AssociatedObjectType = *r.AssociatedObjectType
	}
	return
}

// The category of the Event. We may add additional possible values for this enum
// over time; your application should be able to handle such additions gracefully.
func (r *Event) GetCategory() (Category EventCategory) {
	if r != nil && r.Category != nil {
		Category = *r.Category
	}
	return
}

// The time the Event was created.
func (r *Event) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The Event identifier.
func (r *Event) GetID() (ID string) {
	if r != nil && r.ID != nil {
		ID = *r.ID
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `event`.
func (r *Event) GetType() (Type EventType) {
	if r != nil && r.Type != nil {
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
	AssociatedObjectID *string                    `query:"associated_object_id"`
	CreatedAt          *EventsListParamsCreatedAt `query:"created_at"`
	Category           *EventsListParamsCategory  `query:"category"`
	jsonFields         map[string]interface{}     `pjson:"-,extras"`
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
func (r *EventListParams) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *EventListParams) GetLimit() (Limit int64) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// Filter Events to those belonging to the object with the provided identifier.
func (r *EventListParams) GetAssociatedObjectID() (AssociatedObjectID string) {
	if r != nil && r.AssociatedObjectID != nil {
		AssociatedObjectID = *r.AssociatedObjectID
	}
	return
}

func (r *EventListParams) GetCreatedAt() (CreatedAt EventsListParamsCreatedAt) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

func (r *EventListParams) GetCategory() (Category EventsListParamsCategory) {
	if r != nil && r.Category != nil {
		Category = *r.Category
	}
	return
}

func (r EventListParams) String() (result string) {
	return fmt.Sprintf("&EventListParams{Cursor:%s Limit:%s AssociatedObjectID:%s CreatedAt:%s Category:%s}", core.FmtP(r.Cursor), core.FmtP(r.Limit), core.FmtP(r.AssociatedObjectID), r.CreatedAt, r.Category)
}

type EventsListParamsCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After *string `pjson:"after"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before *string `pjson:"before"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter *string `pjson:"on_or_after"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore *string                `pjson:"on_or_before"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into EventsListParamsCreatedAt
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *EventsListParamsCreatedAt) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes EventsListParamsCreatedAt into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *EventsListParamsCreatedAt) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes EventsListParamsCreatedAt into a url.Values of the query
// parameters associated with this value
func (r *EventsListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r *EventsListParamsCreatedAt) GetAfter() (After string) {
	if r != nil && r.After != nil {
		After = *r.After
	}
	return
}

// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r *EventsListParamsCreatedAt) GetBefore() (Before string) {
	if r != nil && r.Before != nil {
		Before = *r.Before
	}
	return
}

// Return results on or after this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r *EventsListParamsCreatedAt) GetOnOrAfter() (OnOrAfter string) {
	if r != nil && r.OnOrAfter != nil {
		OnOrAfter = *r.OnOrAfter
	}
	return
}

// Return results on or before this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r *EventsListParamsCreatedAt) GetOnOrBefore() (OnOrBefore string) {
	if r != nil && r.OnOrBefore != nil {
		OnOrBefore = *r.OnOrBefore
	}
	return
}

func (r EventsListParamsCreatedAt) String() (result string) {
	return fmt.Sprintf("&EventsListParamsCreatedAt{After:%s Before:%s OnOrAfter:%s OnOrBefore:%s}", core.FmtP(r.After), core.FmtP(r.Before), core.FmtP(r.OnOrAfter), core.FmtP(r.OnOrBefore))
}

type EventsListParamsCategory struct {
	// Return results whose value is in the provided list. For GET requests, this
	// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
	In         *[]EventsListParamsCategoryIn `pjson:"in"`
	jsonFields map[string]interface{}        `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into EventsListParamsCategory
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *EventsListParamsCategory) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes EventsListParamsCategory into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *EventsListParamsCategory) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes EventsListParamsCategory into a url.Values of the query
// parameters associated with this value
func (r *EventsListParamsCategory) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// Return results whose value is in the provided list. For GET requests, this
// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
func (r *EventsListParamsCategory) GetIn() (In []EventsListParamsCategoryIn) {
	if r != nil && r.In != nil {
		In = *r.In
	}
	return
}

func (r EventsListParamsCategory) String() (result string) {
	return fmt.Sprintf("&EventsListParamsCategory{In:%s}", core.Fmt(r.In))
}

type EventsListParamsCategoryIn string

const (
	EventsListParamsCategoryInAccountCreated                                       EventsListParamsCategoryIn = "account.created"
	EventsListParamsCategoryInAccountUpdated                                       EventsListParamsCategoryIn = "account.updated"
	EventsListParamsCategoryInAccountNumberCreated                                 EventsListParamsCategoryIn = "account_number.created"
	EventsListParamsCategoryInAccountNumberUpdated                                 EventsListParamsCategoryIn = "account_number.updated"
	EventsListParamsCategoryInAccountStatementCreated                              EventsListParamsCategoryIn = "account_statement.created"
	EventsListParamsCategoryInAccountTransferCreated                               EventsListParamsCategoryIn = "account_transfer.created"
	EventsListParamsCategoryInAccountTransferUpdated                               EventsListParamsCategoryIn = "account_transfer.updated"
	EventsListParamsCategoryInACHPrenotificationCreated                            EventsListParamsCategoryIn = "ach_prenotification.created"
	EventsListParamsCategoryInACHPrenotificationUpdated                            EventsListParamsCategoryIn = "ach_prenotification.updated"
	EventsListParamsCategoryInACHTransferCreated                                   EventsListParamsCategoryIn = "ach_transfer.created"
	EventsListParamsCategoryInACHTransferUpdated                                   EventsListParamsCategoryIn = "ach_transfer.updated"
	EventsListParamsCategoryInCardCreated                                          EventsListParamsCategoryIn = "card.created"
	EventsListParamsCategoryInCardUpdated                                          EventsListParamsCategoryIn = "card.updated"
	EventsListParamsCategoryInCardDisputeCreated                                   EventsListParamsCategoryIn = "card_dispute.created"
	EventsListParamsCategoryInCardDisputeUpdated                                   EventsListParamsCategoryIn = "card_dispute.updated"
	EventsListParamsCategoryInCheckDepositCreated                                  EventsListParamsCategoryIn = "check_deposit.created"
	EventsListParamsCategoryInCheckDepositUpdated                                  EventsListParamsCategoryIn = "check_deposit.updated"
	EventsListParamsCategoryInCheckTransferCreated                                 EventsListParamsCategoryIn = "check_transfer.created"
	EventsListParamsCategoryInCheckTransferUpdated                                 EventsListParamsCategoryIn = "check_transfer.updated"
	EventsListParamsCategoryInDeclinedTransactionCreated                           EventsListParamsCategoryIn = "declined_transaction.created"
	EventsListParamsCategoryInDigitalWalletTokenCreated                            EventsListParamsCategoryIn = "digital_wallet_token.created"
	EventsListParamsCategoryInDigitalWalletTokenUpdated                            EventsListParamsCategoryIn = "digital_wallet_token.updated"
	EventsListParamsCategoryInDocumentCreated                                      EventsListParamsCategoryIn = "document.created"
	EventsListParamsCategoryInEntityCreated                                        EventsListParamsCategoryIn = "entity.created"
	EventsListParamsCategoryInEntityUpdated                                        EventsListParamsCategoryIn = "entity.updated"
	EventsListParamsCategoryInExternalAccountCreated                               EventsListParamsCategoryIn = "external_account.created"
	EventsListParamsCategoryInFileCreated                                          EventsListParamsCategoryIn = "file.created"
	EventsListParamsCategoryInGroupUpdated                                         EventsListParamsCategoryIn = "group.updated"
	EventsListParamsCategoryInGroupHeartbeat                                       EventsListParamsCategoryIn = "group.heartbeat"
	EventsListParamsCategoryInInboundWireDrawdownRequestCreated                    EventsListParamsCategoryIn = "inbound_wire_drawdown_request.created"
	EventsListParamsCategoryInOauthConnectionCreated                               EventsListParamsCategoryIn = "oauth_connection.created"
	EventsListParamsCategoryInOauthConnectionDeactivated                           EventsListParamsCategoryIn = "oauth_connection.deactivated"
	EventsListParamsCategoryInPendingTransactionCreated                            EventsListParamsCategoryIn = "pending_transaction.created"
	EventsListParamsCategoryInPendingTransactionUpdated                            EventsListParamsCategoryIn = "pending_transaction.updated"
	EventsListParamsCategoryInRealTimeDecisionCardAuthorizationRequested           EventsListParamsCategoryIn = "real_time_decision.card_authorization_requested"
	EventsListParamsCategoryInRealTimeDecisionDigitalWalletTokenRequested          EventsListParamsCategoryIn = "real_time_decision.digital_wallet_token_requested"
	EventsListParamsCategoryInRealTimeDecisionDigitalWalletAuthenticationRequested EventsListParamsCategoryIn = "real_time_decision.digital_wallet_authentication_requested"
	EventsListParamsCategoryInRealTimePaymentsTransferCreated                      EventsListParamsCategoryIn = "real_time_payments_transfer.created"
	EventsListParamsCategoryInRealTimePaymentsTransferUpdated                      EventsListParamsCategoryIn = "real_time_payments_transfer.updated"
	EventsListParamsCategoryInRealTimePaymentsRequestForPaymentCreated             EventsListParamsCategoryIn = "real_time_payments_request_for_payment.created"
	EventsListParamsCategoryInRealTimePaymentsRequestForPaymentUpdated             EventsListParamsCategoryIn = "real_time_payments_request_for_payment.updated"
	EventsListParamsCategoryInTransactionCreated                                   EventsListParamsCategoryIn = "transaction.created"
	EventsListParamsCategoryInWireDrawdownRequestCreated                           EventsListParamsCategoryIn = "wire_drawdown_request.created"
	EventsListParamsCategoryInWireDrawdownRequestUpdated                           EventsListParamsCategoryIn = "wire_drawdown_request.updated"
	EventsListParamsCategoryInWireTransferCreated                                  EventsListParamsCategoryIn = "wire_transfer.created"
	EventsListParamsCategoryInWireTransferUpdated                                  EventsListParamsCategoryIn = "wire_transfer.updated"
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
func (r *EventList) GetData() (Data []Event) {
	if r != nil && r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r *EventList) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
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

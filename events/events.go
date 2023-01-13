package events

import "context"
import "increase/core"
import "fmt"
import "increase/pagination"

type EventService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewEventService(requester core.Requester) (r *EventService) {
	r = &EventService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

type PreloadedEventService struct {
	Events *EventService
}

func (r *PreloadedEventService) Init(service *EventService) {
	r.Events = service
}

func NewPreloadedEventService(service *EventService) (r *PreloadedEventService) {
	r = &PreloadedEventService{}
	r.Init(service)
	return
}

//
type Event struct {
	// The identifier of the object that generated this Event.
	AssociatedObjectId *string `json:"associated_object_id"`
	// The type of the object that generated this Event.
	AssociatedObjectType *string `json:"associated_object_type"`
	// The category of the Event. We may add additional possible values for this enum
	// over time; your application should be able to handle such additions gracefully.
	Category *EventCategory `json:"category"`
	// The time the Event was created.
	CreatedAt *string `json:"created_at"`
	// The Event identifier.
	Id *string `json:"id"`
	// A constant representing the object's type. For this resource it will always be
	// `event`.
	Type *EventType `json:"type"`
}

// The identifier of the object that generated this Event.
func (r *Event) GetAssociatedObjectId() (AssociatedObjectId string) {
	if r != nil && r.AssociatedObjectId != nil {
		AssociatedObjectId = *r.AssociatedObjectId
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
func (r *Event) GetId() (Id string) {
	if r != nil && r.Id != nil {
		Id = *r.Id
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
	EventCategoryOauthConnectionCreated                               EventCategory = "oauth_connection.created"
	EventCategoryOauthConnectionDeactivated                           EventCategory = "oauth_connection.deactivated"
	EventCategoryPendingTransactionCreated                            EventCategory = "pending_transaction.created"
	EventCategoryPendingTransactionUpdated                            EventCategory = "pending_transaction.updated"
	EventCategoryRealTimeDecisionCardAuthorizationRequested           EventCategory = "real_time_decision.card_authorization_requested"
	EventCategoryRealTimeDecisionDigitalWalletTokenRequested          EventCategory = "real_time_decision.digital_wallet_token_requested"
	EventCategoryRealTimeDecisionDigitalWalletAuthenticationRequested EventCategory = "real_time_decision.digital_wallet_authentication_requested"
	EventCategoryRealTimePaymentsTransferCreated                      EventCategory = "real_time_payments_transfer.created"
	EventCategoryRealTimePaymentsTransferUpdated                      EventCategory = "real_time_payments_transfer.updated"
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

type ListEventsQuery struct {
	// Return the page of entries after this one.
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int `query:"limit"`
	// Filter Events to those belonging to the object with the provided identifier.
	AssociatedObjectId *string                   `query:"associated_object_id"`
	CreatedAt          *ListEventsQueryCreatedAt `query:"created_at"`
	Category           *ListEventsQueryCategory  `query:"category"`
}

// Return the page of entries after this one.
func (r *ListEventsQuery) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *ListEventsQuery) GetLimit() (Limit int) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// Filter Events to those belonging to the object with the provided identifier.
func (r *ListEventsQuery) GetAssociatedObjectId() (AssociatedObjectId string) {
	if r != nil && r.AssociatedObjectId != nil {
		AssociatedObjectId = *r.AssociatedObjectId
	}
	return
}

func (r *ListEventsQuery) GetCreatedAt() (CreatedAt ListEventsQueryCreatedAt) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

func (r *ListEventsQuery) GetCategory() (Category ListEventsQueryCategory) {
	if r != nil && r.Category != nil {
		Category = *r.Category
	}
	return
}

type ListEventsQueryCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After *string `json:"after,omitempty"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before *string `json:"before,omitempty"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter *string `json:"on_or_after,omitempty"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore *string `json:"on_or_before,omitempty"`
}

// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r *ListEventsQueryCreatedAt) GetAfter() (After string) {
	if r != nil && r.After != nil {
		After = *r.After
	}
	return
}

// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r *ListEventsQueryCreatedAt) GetBefore() (Before string) {
	if r != nil && r.Before != nil {
		Before = *r.Before
	}
	return
}

// Return results on or after this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r *ListEventsQueryCreatedAt) GetOnOrAfter() (OnOrAfter string) {
	if r != nil && r.OnOrAfter != nil {
		OnOrAfter = *r.OnOrAfter
	}
	return
}

// Return results on or before this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r *ListEventsQueryCreatedAt) GetOnOrBefore() (OnOrBefore string) {
	if r != nil && r.OnOrBefore != nil {
		OnOrBefore = *r.OnOrBefore
	}
	return
}

type ListEventsQueryCategory struct {
	// Return results whose value is in the provided list. For GET requests, this
	// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
	In *[]ListEventsQueryCategoryIn `json:"in,omitempty"`
}

// Return results whose value is in the provided list. For GET requests, this
// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
func (r *ListEventsQueryCategory) GetIn() (In []ListEventsQueryCategoryIn) {
	if r != nil && r.In != nil {
		In = *r.In
	}
	return
}

type ListEventsQueryCategoryIn string

const (
	ListEventsQueryCategoryInAccountCreated                                       ListEventsQueryCategoryIn = "account.created"
	ListEventsQueryCategoryInAccountUpdated                                       ListEventsQueryCategoryIn = "account.updated"
	ListEventsQueryCategoryInAccountNumberCreated                                 ListEventsQueryCategoryIn = "account_number.created"
	ListEventsQueryCategoryInAccountNumberUpdated                                 ListEventsQueryCategoryIn = "account_number.updated"
	ListEventsQueryCategoryInAccountStatementCreated                              ListEventsQueryCategoryIn = "account_statement.created"
	ListEventsQueryCategoryInAccountTransferCreated                               ListEventsQueryCategoryIn = "account_transfer.created"
	ListEventsQueryCategoryInAccountTransferUpdated                               ListEventsQueryCategoryIn = "account_transfer.updated"
	ListEventsQueryCategoryInACHPrenotificationCreated                            ListEventsQueryCategoryIn = "ach_prenotification.created"
	ListEventsQueryCategoryInACHPrenotificationUpdated                            ListEventsQueryCategoryIn = "ach_prenotification.updated"
	ListEventsQueryCategoryInACHTransferCreated                                   ListEventsQueryCategoryIn = "ach_transfer.created"
	ListEventsQueryCategoryInACHTransferUpdated                                   ListEventsQueryCategoryIn = "ach_transfer.updated"
	ListEventsQueryCategoryInCardCreated                                          ListEventsQueryCategoryIn = "card.created"
	ListEventsQueryCategoryInCardUpdated                                          ListEventsQueryCategoryIn = "card.updated"
	ListEventsQueryCategoryInCardDisputeCreated                                   ListEventsQueryCategoryIn = "card_dispute.created"
	ListEventsQueryCategoryInCardDisputeUpdated                                   ListEventsQueryCategoryIn = "card_dispute.updated"
	ListEventsQueryCategoryInCheckDepositCreated                                  ListEventsQueryCategoryIn = "check_deposit.created"
	ListEventsQueryCategoryInCheckDepositUpdated                                  ListEventsQueryCategoryIn = "check_deposit.updated"
	ListEventsQueryCategoryInCheckTransferCreated                                 ListEventsQueryCategoryIn = "check_transfer.created"
	ListEventsQueryCategoryInCheckTransferUpdated                                 ListEventsQueryCategoryIn = "check_transfer.updated"
	ListEventsQueryCategoryInDeclinedTransactionCreated                           ListEventsQueryCategoryIn = "declined_transaction.created"
	ListEventsQueryCategoryInDigitalWalletTokenCreated                            ListEventsQueryCategoryIn = "digital_wallet_token.created"
	ListEventsQueryCategoryInDigitalWalletTokenUpdated                            ListEventsQueryCategoryIn = "digital_wallet_token.updated"
	ListEventsQueryCategoryInDocumentCreated                                      ListEventsQueryCategoryIn = "document.created"
	ListEventsQueryCategoryInEntityCreated                                        ListEventsQueryCategoryIn = "entity.created"
	ListEventsQueryCategoryInEntityUpdated                                        ListEventsQueryCategoryIn = "entity.updated"
	ListEventsQueryCategoryInExternalAccountCreated                               ListEventsQueryCategoryIn = "external_account.created"
	ListEventsQueryCategoryInFileCreated                                          ListEventsQueryCategoryIn = "file.created"
	ListEventsQueryCategoryInGroupUpdated                                         ListEventsQueryCategoryIn = "group.updated"
	ListEventsQueryCategoryInGroupHeartbeat                                       ListEventsQueryCategoryIn = "group.heartbeat"
	ListEventsQueryCategoryInOauthConnectionCreated                               ListEventsQueryCategoryIn = "oauth_connection.created"
	ListEventsQueryCategoryInOauthConnectionDeactivated                           ListEventsQueryCategoryIn = "oauth_connection.deactivated"
	ListEventsQueryCategoryInPendingTransactionCreated                            ListEventsQueryCategoryIn = "pending_transaction.created"
	ListEventsQueryCategoryInPendingTransactionUpdated                            ListEventsQueryCategoryIn = "pending_transaction.updated"
	ListEventsQueryCategoryInRealTimeDecisionCardAuthorizationRequested           ListEventsQueryCategoryIn = "real_time_decision.card_authorization_requested"
	ListEventsQueryCategoryInRealTimeDecisionDigitalWalletTokenRequested          ListEventsQueryCategoryIn = "real_time_decision.digital_wallet_token_requested"
	ListEventsQueryCategoryInRealTimeDecisionDigitalWalletAuthenticationRequested ListEventsQueryCategoryIn = "real_time_decision.digital_wallet_authentication_requested"
	ListEventsQueryCategoryInRealTimePaymentsTransferCreated                      ListEventsQueryCategoryIn = "real_time_payments_transfer.created"
	ListEventsQueryCategoryInRealTimePaymentsTransferUpdated                      ListEventsQueryCategoryIn = "real_time_payments_transfer.updated"
	ListEventsQueryCategoryInTransactionCreated                                   ListEventsQueryCategoryIn = "transaction.created"
	ListEventsQueryCategoryInWireDrawdownRequestCreated                           ListEventsQueryCategoryIn = "wire_drawdown_request.created"
	ListEventsQueryCategoryInWireDrawdownRequestUpdated                           ListEventsQueryCategoryIn = "wire_drawdown_request.updated"
	ListEventsQueryCategoryInWireTransferCreated                                  ListEventsQueryCategoryIn = "wire_transfer.created"
	ListEventsQueryCategoryInWireTransferUpdated                                  ListEventsQueryCategoryIn = "wire_transfer.updated"
)

//
type EventList struct {
	// The contents of the list.
	Data *[]Event `json:"data"`
	// A pointer to a place in the list.
	NextCursor *string `json:"next_cursor"`
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

func (r *EventService) Retrieve(ctx context.Context, event_id string, opts ...*core.RequestOpts) (res *Event, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/events/%s", event_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *PreloadedEventService) Retrieve(ctx context.Context, event_id string, opts ...*core.RequestOpts) (res *Event, err error) {
	err = r.Events.get(
		ctx,
		fmt.Sprintf("/events/%s", event_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

type EventsPage struct {
	*pagination.Page[Event]
}

func (r *EventsPage) Event() *Event {
	return r.Current()
}

func (r *EventsPage) GetNextPage() (*EventsPage, error) {
	if page, err := r.Page.GetNextPage(); err != nil {
		return nil, err
	} else {
		return &EventsPage{page}, nil
	}
}

func (r *EventService) List(ctx context.Context, query *ListEventsQuery, opts ...*core.RequestOpts) (res *EventsPage, err error) {
	page := &EventsPage{
		Page: &pagination.Page[Event]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/events",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

func (r *PreloadedEventService) List(ctx context.Context, query *ListEventsQuery, opts ...*core.RequestOpts) (res *EventsPage, err error) {
	page := &EventsPage{
		Page: &pagination.Page[Event]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/events",
			},
			Requester: r.Events.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

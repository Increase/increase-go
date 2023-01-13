package event_subscriptions

import "context"
import "increase/core"
import "fmt"
import "increase/pagination"

type EventSubscriptionService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewEventSubscriptionService(requester core.Requester) (r *EventSubscriptionService) {
	r = &EventSubscriptionService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

type PreloadedEventSubscriptionService struct {
	EventSubscriptions *EventSubscriptionService
}

func (r *PreloadedEventSubscriptionService) Init(service *EventSubscriptionService) {
	r.EventSubscriptions = service
}

func NewPreloadedEventSubscriptionService(service *EventSubscriptionService) (r *PreloadedEventSubscriptionService) {
	r = &PreloadedEventSubscriptionService{}
	r.Init(service)
	return
}

//
type EventSubscription struct {
	// The event subscription identifier.
	Id *string `json:"id"`
	// The time the event subscription was created.
	CreatedAt *string `json:"created_at"`
	// This indicates if we'll send notifications to this subscription.
	Status *EventSubscriptionStatus `json:"status"`
	// If specified, this subscription will only receive webhooks for Events with the
	// specified `category`.
	SelectedEventCategory *EventSubscriptionSelectedEventCategory `json:"selected_event_category"`
	// The webhook url where we'll send notifications.
	URL *string `json:"url"`
	// The key that will be used to sign webhooks.
	SharedSecret *string `json:"shared_secret"`
	// A constant representing the object's type. For this resource it will always be
	// `event_subscription`.
	Type *EventSubscriptionType `json:"type"`
}

// The event subscription identifier.
func (r *EventSubscription) GetId() (Id string) {
	if r != nil && r.Id != nil {
		Id = *r.Id
	}
	return
}

// The time the event subscription was created.
func (r *EventSubscription) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// This indicates if we'll send notifications to this subscription.
func (r *EventSubscription) GetStatus() (Status EventSubscriptionStatus) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

// If specified, this subscription will only receive webhooks for Events with the
// specified `category`.
func (r *EventSubscription) GetSelectedEventCategory() (SelectedEventCategory EventSubscriptionSelectedEventCategory) {
	if r != nil && r.SelectedEventCategory != nil {
		SelectedEventCategory = *r.SelectedEventCategory
	}
	return
}

// The webhook url where we'll send notifications.
func (r *EventSubscription) GetURL() (URL string) {
	if r != nil && r.URL != nil {
		URL = *r.URL
	}
	return
}

// The key that will be used to sign webhooks.
func (r *EventSubscription) GetSharedSecret() (SharedSecret string) {
	if r != nil && r.SharedSecret != nil {
		SharedSecret = *r.SharedSecret
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `event_subscription`.
func (r *EventSubscription) GetType() (Type EventSubscriptionType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
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
	EventSubscriptionSelectedEventCategoryOauthConnectionCreated                               EventSubscriptionSelectedEventCategory = "oauth_connection.created"
	EventSubscriptionSelectedEventCategoryOauthConnectionDeactivated                           EventSubscriptionSelectedEventCategory = "oauth_connection.deactivated"
	EventSubscriptionSelectedEventCategoryPendingTransactionCreated                            EventSubscriptionSelectedEventCategory = "pending_transaction.created"
	EventSubscriptionSelectedEventCategoryPendingTransactionUpdated                            EventSubscriptionSelectedEventCategory = "pending_transaction.updated"
	EventSubscriptionSelectedEventCategoryRealTimeDecisionCardAuthorizationRequested           EventSubscriptionSelectedEventCategory = "real_time_decision.card_authorization_requested"
	EventSubscriptionSelectedEventCategoryRealTimeDecisionDigitalWalletTokenRequested          EventSubscriptionSelectedEventCategory = "real_time_decision.digital_wallet_token_requested"
	EventSubscriptionSelectedEventCategoryRealTimeDecisionDigitalWalletAuthenticationRequested EventSubscriptionSelectedEventCategory = "real_time_decision.digital_wallet_authentication_requested"
	EventSubscriptionSelectedEventCategoryRealTimePaymentsTransferCreated                      EventSubscriptionSelectedEventCategory = "real_time_payments_transfer.created"
	EventSubscriptionSelectedEventCategoryRealTimePaymentsTransferUpdated                      EventSubscriptionSelectedEventCategory = "real_time_payments_transfer.updated"
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
	URL *string `json:"url"`
	// The key that will be used to sign webhooks. If no value is passed, a random
	// string will be used as default.
	SharedSecret *string `json:"shared_secret,omitempty"`
	// If specified, this subscription will only receive webhooks for Events with the
	// specified `category`.
	SelectedEventCategory *CreateAnEventSubscriptionParametersSelectedEventCategory `json:"selected_event_category,omitempty"`
}

// The URL you'd like us to send webhooks to.
func (r *CreateAnEventSubscriptionParameters) GetURL() (URL string) {
	if r != nil && r.URL != nil {
		URL = *r.URL
	}
	return
}

// The key that will be used to sign webhooks. If no value is passed, a random
// string will be used as default.
func (r *CreateAnEventSubscriptionParameters) GetSharedSecret() (SharedSecret string) {
	if r != nil && r.SharedSecret != nil {
		SharedSecret = *r.SharedSecret
	}
	return
}

// If specified, this subscription will only receive webhooks for Events with the
// specified `category`.
func (r *CreateAnEventSubscriptionParameters) GetSelectedEventCategory() (SelectedEventCategory CreateAnEventSubscriptionParametersSelectedEventCategory) {
	if r != nil && r.SelectedEventCategory != nil {
		SelectedEventCategory = *r.SelectedEventCategory
	}
	return
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
	CreateAnEventSubscriptionParametersSelectedEventCategoryOauthConnectionCreated                               CreateAnEventSubscriptionParametersSelectedEventCategory = "oauth_connection.created"
	CreateAnEventSubscriptionParametersSelectedEventCategoryOauthConnectionDeactivated                           CreateAnEventSubscriptionParametersSelectedEventCategory = "oauth_connection.deactivated"
	CreateAnEventSubscriptionParametersSelectedEventCategoryPendingTransactionCreated                            CreateAnEventSubscriptionParametersSelectedEventCategory = "pending_transaction.created"
	CreateAnEventSubscriptionParametersSelectedEventCategoryPendingTransactionUpdated                            CreateAnEventSubscriptionParametersSelectedEventCategory = "pending_transaction.updated"
	CreateAnEventSubscriptionParametersSelectedEventCategoryRealTimeDecisionCardAuthorizationRequested           CreateAnEventSubscriptionParametersSelectedEventCategory = "real_time_decision.card_authorization_requested"
	CreateAnEventSubscriptionParametersSelectedEventCategoryRealTimeDecisionDigitalWalletTokenRequested          CreateAnEventSubscriptionParametersSelectedEventCategory = "real_time_decision.digital_wallet_token_requested"
	CreateAnEventSubscriptionParametersSelectedEventCategoryRealTimeDecisionDigitalWalletAuthenticationRequested CreateAnEventSubscriptionParametersSelectedEventCategory = "real_time_decision.digital_wallet_authentication_requested"
	CreateAnEventSubscriptionParametersSelectedEventCategoryRealTimePaymentsTransferCreated                      CreateAnEventSubscriptionParametersSelectedEventCategory = "real_time_payments_transfer.created"
	CreateAnEventSubscriptionParametersSelectedEventCategoryRealTimePaymentsTransferUpdated                      CreateAnEventSubscriptionParametersSelectedEventCategory = "real_time_payments_transfer.updated"
	CreateAnEventSubscriptionParametersSelectedEventCategoryTransactionCreated                                   CreateAnEventSubscriptionParametersSelectedEventCategory = "transaction.created"
	CreateAnEventSubscriptionParametersSelectedEventCategoryWireDrawdownRequestCreated                           CreateAnEventSubscriptionParametersSelectedEventCategory = "wire_drawdown_request.created"
	CreateAnEventSubscriptionParametersSelectedEventCategoryWireDrawdownRequestUpdated                           CreateAnEventSubscriptionParametersSelectedEventCategory = "wire_drawdown_request.updated"
	CreateAnEventSubscriptionParametersSelectedEventCategoryWireTransferCreated                                  CreateAnEventSubscriptionParametersSelectedEventCategory = "wire_transfer.created"
	CreateAnEventSubscriptionParametersSelectedEventCategoryWireTransferUpdated                                  CreateAnEventSubscriptionParametersSelectedEventCategory = "wire_transfer.updated"
)

type UpdateAnEventSubscriptionParameters struct {
	// The status to update the Event Subscription with.
	Status *UpdateAnEventSubscriptionParametersStatus `json:"status,omitempty"`
}

// The status to update the Event Subscription with.
func (r *UpdateAnEventSubscriptionParameters) GetStatus() (Status UpdateAnEventSubscriptionParametersStatus) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

type UpdateAnEventSubscriptionParametersStatus string

const (
	UpdateAnEventSubscriptionParametersStatusActive   UpdateAnEventSubscriptionParametersStatus = "active"
	UpdateAnEventSubscriptionParametersStatusDisabled UpdateAnEventSubscriptionParametersStatus = "disabled"
	UpdateAnEventSubscriptionParametersStatusDeleted  UpdateAnEventSubscriptionParametersStatus = "deleted"
)

type ListEventSubscriptionsQuery struct {
	// Return the page of entries after this one.
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int `query:"limit"`
}

// Return the page of entries after this one.
func (r *ListEventSubscriptionsQuery) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *ListEventSubscriptionsQuery) GetLimit() (Limit int) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

//
type EventSubscriptionList struct {
	// The contents of the list.
	Data *[]EventSubscription `json:"data"`
	// A pointer to a place in the list.
	NextCursor *string `json:"next_cursor"`
}

// The contents of the list.
func (r *EventSubscriptionList) GetData() (Data []EventSubscription) {
	if r != nil && r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r *EventSubscriptionList) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

func (r *EventSubscriptionService) Create(ctx context.Context, body *CreateAnEventSubscriptionParameters, opts ...*core.RequestOpts) (res *EventSubscription, err error) {
	err = r.post(
		ctx,
		"/event_subscriptions",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *PreloadedEventSubscriptionService) Create(ctx context.Context, body *CreateAnEventSubscriptionParameters, opts ...*core.RequestOpts) (res *EventSubscription, err error) {
	err = r.EventSubscriptions.post(
		ctx,
		"/event_subscriptions",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *EventSubscriptionService) Retrieve(ctx context.Context, event_subscription_id string, opts ...*core.RequestOpts) (res *EventSubscription, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/event_subscriptions/%s", event_subscription_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *PreloadedEventSubscriptionService) Retrieve(ctx context.Context, event_subscription_id string, opts ...*core.RequestOpts) (res *EventSubscription, err error) {
	err = r.EventSubscriptions.get(
		ctx,
		fmt.Sprintf("/event_subscriptions/%s", event_subscription_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *EventSubscriptionService) Update(ctx context.Context, event_subscription_id string, body *UpdateAnEventSubscriptionParameters, opts ...*core.RequestOpts) (res *EventSubscription, err error) {
	err = r.patch(
		ctx,
		fmt.Sprintf("/event_subscriptions/%s", event_subscription_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *PreloadedEventSubscriptionService) Update(ctx context.Context, event_subscription_id string, body *UpdateAnEventSubscriptionParameters, opts ...*core.RequestOpts) (res *EventSubscription, err error) {
	err = r.EventSubscriptions.patch(
		ctx,
		fmt.Sprintf("/event_subscriptions/%s", event_subscription_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

type EventSubscriptionsPage struct {
	*pagination.Page[EventSubscription]
}

func (r *EventSubscriptionsPage) EventSubscription() *EventSubscription {
	return r.Current()
}

func (r *EventSubscriptionsPage) GetNextPage() (*EventSubscriptionsPage, error) {
	if page, err := r.Page.GetNextPage(); err != nil {
		return nil, err
	} else {
		return &EventSubscriptionsPage{page}, nil
	}
}

func (r *EventSubscriptionService) List(ctx context.Context, query *ListEventSubscriptionsQuery, opts ...*core.RequestOpts) (res *EventSubscriptionsPage, err error) {
	page := &EventSubscriptionsPage{
		Page: &pagination.Page[EventSubscription]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/event_subscriptions",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

func (r *PreloadedEventSubscriptionService) List(ctx context.Context, query *ListEventSubscriptionsQuery, opts ...*core.RequestOpts) (res *EventSubscriptionsPage, err error) {
	page := &EventSubscriptionsPage{
		Page: &pagination.Page[EventSubscription]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/event_subscriptions",
			},
			Requester: r.EventSubscriptions.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

package requests

import (
	"net/url"

	"github.com/increase/increase-go/core/field"
	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
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

// MarshalJSON serializes EventSubscriptionNewParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *EventSubscriptionNewParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
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

// MarshalJSON serializes EventSubscriptionUpdateParams into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *EventSubscriptionUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
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

// URLQuery serializes EventSubscriptionListParams into a url.Values of the query
// parameters associated with this value
func (r *EventSubscriptionListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

package requests

import (
	"fmt"
	"net/url"

	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/fields"
)

type CreateAnEventSubscriptionParameters struct {
	// The URL you'd like us to send webhooks to.
	URL fields.Field[string] `json:"url,required"`
	// The key that will be used to sign webhooks. If no value is passed, a random
	// string will be used as default.
	SharedSecret fields.Field[string] `json:"shared_secret"`
	// If specified, this subscription will only receive webhooks for Events with the
	// specified `category`.
	SelectedEventCategory fields.Field[CreateAnEventSubscriptionParametersSelectedEventCategory] `json:"selected_event_category"`
}

// MarshalJSON serializes CreateAnEventSubscriptionParameters into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *CreateAnEventSubscriptionParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CreateAnEventSubscriptionParameters) String() (result string) {
	return fmt.Sprintf("&CreateAnEventSubscriptionParameters{URL:%s SharedSecret:%s SelectedEventCategory:%s}", r.URL, r.SharedSecret, r.SelectedEventCategory)
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
	Status fields.Field[UpdateAnEventSubscriptionParametersStatus] `json:"status"`
}

// MarshalJSON serializes UpdateAnEventSubscriptionParameters into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *UpdateAnEventSubscriptionParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r UpdateAnEventSubscriptionParameters) String() (result string) {
	return fmt.Sprintf("&UpdateAnEventSubscriptionParameters{Status:%s}", r.Status)
}

type UpdateAnEventSubscriptionParametersStatus string

const (
	UpdateAnEventSubscriptionParametersStatusActive   UpdateAnEventSubscriptionParametersStatus = "active"
	UpdateAnEventSubscriptionParametersStatusDisabled UpdateAnEventSubscriptionParametersStatus = "disabled"
	UpdateAnEventSubscriptionParametersStatusDeleted  UpdateAnEventSubscriptionParametersStatus = "deleted"
)

type EventSubscriptionListParams struct {
	// Return the page of entries after this one.
	Cursor fields.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit fields.Field[int64] `query:"limit"`
}

// URLQuery serializes EventSubscriptionListParams into a url.Values of the query
// parameters associated with this value
func (r *EventSubscriptionListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r EventSubscriptionListParams) String() (result string) {
	return fmt.Sprintf("&EventSubscriptionListParams{Cursor:%s Limit:%s}", r.Cursor, r.Limit)
}

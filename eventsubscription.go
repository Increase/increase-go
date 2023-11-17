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
func (r *EventSubscriptionService) Get(ctx context.Context, eventSubscriptionID string, opts ...option.RequestOption) (res *EventSubscription, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("event_subscriptions/%s", eventSubscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an Event Subscription
func (r *EventSubscriptionService) Update(ctx context.Context, eventSubscriptionID string, body EventSubscriptionUpdateParams, opts ...option.RequestOption) (res *EventSubscription, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("event_subscriptions/%s", eventSubscriptionID)
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
	// If specified, this subscription will only receive webhooks for Events with the
	// specified `category`.
	SelectedEventCategory EventSubscriptionSelectedEventCategory `json:"selected_event_category,required,nullable"`
	// This indicates if we'll send notifications to this subscription.
	Status EventSubscriptionStatus `json:"status,required"`
	// A constant representing the object's type. For this resource it will always be
	// `event_subscription`.
	Type EventSubscriptionType `json:"type,required"`
	// The webhook url where we'll send notifications.
	URL  string                `json:"url,required"`
	JSON eventSubscriptionJSON `json:"-"`
}

// eventSubscriptionJSON contains the JSON metadata for the struct
// [EventSubscription]
type eventSubscriptionJSON struct {
	ID                    apijson.Field
	CreatedAt             apijson.Field
	SelectedEventCategory apijson.Field
	Status                apijson.Field
	Type                  apijson.Field
	URL                   apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *EventSubscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// If specified, this subscription will only receive webhooks for Events with the
// specified `category`.
type EventSubscriptionSelectedEventCategory string

const (
	// Occurs whenever an Account is created.
	EventSubscriptionSelectedEventCategoryAccountCreated EventSubscriptionSelectedEventCategory = "account.created"
	// Occurs whenever an Account is updated.
	EventSubscriptionSelectedEventCategoryAccountUpdated EventSubscriptionSelectedEventCategory = "account.updated"
	// Occurs whenever an Account Number is created.
	EventSubscriptionSelectedEventCategoryAccountNumberCreated EventSubscriptionSelectedEventCategory = "account_number.created"
	// Occurs whenever an Account Number is updated.
	EventSubscriptionSelectedEventCategoryAccountNumberUpdated EventSubscriptionSelectedEventCategory = "account_number.updated"
	// Occurs whenever an Account Statement is created.
	EventSubscriptionSelectedEventCategoryAccountStatementCreated EventSubscriptionSelectedEventCategory = "account_statement.created"
	// Occurs whenever an Account Transfer is created.
	EventSubscriptionSelectedEventCategoryAccountTransferCreated EventSubscriptionSelectedEventCategory = "account_transfer.created"
	// Occurs whenever an Account Transfer is updated.
	EventSubscriptionSelectedEventCategoryAccountTransferUpdated EventSubscriptionSelectedEventCategory = "account_transfer.updated"
	// Occurs whenever an ACH Prenotification is created.
	EventSubscriptionSelectedEventCategoryACHPrenotificationCreated EventSubscriptionSelectedEventCategory = "ach_prenotification.created"
	// Occurs whenever an ACH Prenotification is updated.
	EventSubscriptionSelectedEventCategoryACHPrenotificationUpdated EventSubscriptionSelectedEventCategory = "ach_prenotification.updated"
	// Occurs whenever an ACH Transfer is created.
	EventSubscriptionSelectedEventCategoryACHTransferCreated EventSubscriptionSelectedEventCategory = "ach_transfer.created"
	// Occurs whenever an ACH Transfer is updated.
	EventSubscriptionSelectedEventCategoryACHTransferUpdated EventSubscriptionSelectedEventCategory = "ach_transfer.updated"
	// Occurs whenever a Bookkeeping Account is created.
	EventSubscriptionSelectedEventCategoryBookkeepingAccountCreated EventSubscriptionSelectedEventCategory = "bookkeeping_account.created"
	// Occurs whenever a Bookkeeping Account is updated.
	EventSubscriptionSelectedEventCategoryBookkeepingAccountUpdated EventSubscriptionSelectedEventCategory = "bookkeeping_account.updated"
	// Occurs whenever a Bookkeeping Entry Set is created.
	EventSubscriptionSelectedEventCategoryBookkeepingEntrySetUpdated EventSubscriptionSelectedEventCategory = "bookkeeping_entry_set.updated"
	// Occurs whenever a Card is created.
	EventSubscriptionSelectedEventCategoryCardCreated EventSubscriptionSelectedEventCategory = "card.created"
	// Occurs whenever a Card is updated.
	EventSubscriptionSelectedEventCategoryCardUpdated EventSubscriptionSelectedEventCategory = "card.updated"
	// Occurs whenever a Card Payment is created.
	EventSubscriptionSelectedEventCategoryCardPaymentCreated EventSubscriptionSelectedEventCategory = "card_payment.created"
	// Occurs whenever a Card Payment is updated.
	EventSubscriptionSelectedEventCategoryCardPaymentUpdated EventSubscriptionSelectedEventCategory = "card_payment.updated"
	// Occurs whenever a Card Profile is created.
	EventSubscriptionSelectedEventCategoryCardProfileCreated EventSubscriptionSelectedEventCategory = "card_profile.created"
	// Occurs whenever a Card Profile is updated.
	EventSubscriptionSelectedEventCategoryCardProfileUpdated EventSubscriptionSelectedEventCategory = "card_profile.updated"
	// Occurs whenever a Card Dispute is created.
	EventSubscriptionSelectedEventCategoryCardDisputeCreated EventSubscriptionSelectedEventCategory = "card_dispute.created"
	// Occurs whenever a Card Dispute is updated.
	EventSubscriptionSelectedEventCategoryCardDisputeUpdated EventSubscriptionSelectedEventCategory = "card_dispute.updated"
	// Occurs whenever a Check Deposit is created.
	EventSubscriptionSelectedEventCategoryCheckDepositCreated EventSubscriptionSelectedEventCategory = "check_deposit.created"
	// Occurs whenever a Check Deposit is updated.
	EventSubscriptionSelectedEventCategoryCheckDepositUpdated EventSubscriptionSelectedEventCategory = "check_deposit.updated"
	// Occurs whenever a Check Transfer is created.
	EventSubscriptionSelectedEventCategoryCheckTransferCreated EventSubscriptionSelectedEventCategory = "check_transfer.created"
	// Occurs whenever a Check Transfer is updated.
	EventSubscriptionSelectedEventCategoryCheckTransferUpdated EventSubscriptionSelectedEventCategory = "check_transfer.updated"
	// Occurs whenever a Declined Transaction is created.
	EventSubscriptionSelectedEventCategoryDeclinedTransactionCreated EventSubscriptionSelectedEventCategory = "declined_transaction.created"
	// Occurs whenever a Digital Wallet Token is created.
	EventSubscriptionSelectedEventCategoryDigitalWalletTokenCreated EventSubscriptionSelectedEventCategory = "digital_wallet_token.created"
	// Occurs whenever a Digital Wallet Token is updated.
	EventSubscriptionSelectedEventCategoryDigitalWalletTokenUpdated EventSubscriptionSelectedEventCategory = "digital_wallet_token.updated"
	// Occurs whenever a Document is created.
	EventSubscriptionSelectedEventCategoryDocumentCreated EventSubscriptionSelectedEventCategory = "document.created"
	// Occurs whenever an Entity is created.
	EventSubscriptionSelectedEventCategoryEntityCreated EventSubscriptionSelectedEventCategory = "entity.created"
	// Occurs whenever an Entity is updated.
	EventSubscriptionSelectedEventCategoryEntityUpdated EventSubscriptionSelectedEventCategory = "entity.updated"
	// Occurs whenever an Event Subscription is created.
	EventSubscriptionSelectedEventCategoryEventSubscriptionCreated EventSubscriptionSelectedEventCategory = "event_subscription.created"
	// Occurs whenever an Event Subscription is updated.
	EventSubscriptionSelectedEventCategoryEventSubscriptionUpdated EventSubscriptionSelectedEventCategory = "event_subscription.updated"
	// Occurs whenever an Export is created.
	EventSubscriptionSelectedEventCategoryExportCreated EventSubscriptionSelectedEventCategory = "export.created"
	// Occurs whenever an Export is updated.
	EventSubscriptionSelectedEventCategoryExportUpdated EventSubscriptionSelectedEventCategory = "export.updated"
	// Occurs whenever an External Account is created.
	EventSubscriptionSelectedEventCategoryExternalAccountCreated EventSubscriptionSelectedEventCategory = "external_account.created"
	// Occurs whenever an External Account is updated.
	EventSubscriptionSelectedEventCategoryExternalAccountUpdated EventSubscriptionSelectedEventCategory = "external_account.updated"
	// Occurs whenever a File is created.
	EventSubscriptionSelectedEventCategoryFileCreated EventSubscriptionSelectedEventCategory = "file.created"
	// Occurs whenever a Group is updated.
	EventSubscriptionSelectedEventCategoryGroupUpdated EventSubscriptionSelectedEventCategory = "group.updated"
	// Increase may send webhooks with this category to see if a webhook endpoint is
	// working properly.
	EventSubscriptionSelectedEventCategoryGroupHeartbeat EventSubscriptionSelectedEventCategory = "group.heartbeat"
	// Occurs whenever an Inbound ACH Transfer is created.
	EventSubscriptionSelectedEventCategoryInboundACHTransferCreated EventSubscriptionSelectedEventCategory = "inbound_ach_transfer.created"
	// Occurs whenever an Inbound ACH Transfer is updated.
	EventSubscriptionSelectedEventCategoryInboundACHTransferUpdated EventSubscriptionSelectedEventCategory = "inbound_ach_transfer.updated"
	// Occurs whenever an Inbound ACH Transfer Return is created.
	EventSubscriptionSelectedEventCategoryInboundACHTransferReturnCreated EventSubscriptionSelectedEventCategory = "inbound_ach_transfer_return.created"
	// Occurs whenever an Inbound ACH Transfer Return is updated.
	EventSubscriptionSelectedEventCategoryInboundACHTransferReturnUpdated EventSubscriptionSelectedEventCategory = "inbound_ach_transfer_return.updated"
	// Occurs whenever an Inbound Wire Drawdown Request is created.
	EventSubscriptionSelectedEventCategoryInboundWireDrawdownRequestCreated EventSubscriptionSelectedEventCategory = "inbound_wire_drawdown_request.created"
	// Occurs whenever an IntraFi Account Enrollment is created.
	EventSubscriptionSelectedEventCategoryIntrafiAccountEnrollmentCreated EventSubscriptionSelectedEventCategory = "intrafi_account_enrollment.created"
	// Occurs whenever an IntraFi Account Enrollment is updated.
	EventSubscriptionSelectedEventCategoryIntrafiAccountEnrollmentUpdated EventSubscriptionSelectedEventCategory = "intrafi_account_enrollment.updated"
	// Occurs whenever an IntraFi Exclusion is created.
	EventSubscriptionSelectedEventCategoryIntrafiExclusionCreated EventSubscriptionSelectedEventCategory = "intrafi_exclusion.created"
	// Occurs whenever an IntraFi Exclusion is updated.
	EventSubscriptionSelectedEventCategoryIntrafiExclusionUpdated EventSubscriptionSelectedEventCategory = "intrafi_exclusion.updated"
	// Occurs whenever an OAuth Connection is created.
	EventSubscriptionSelectedEventCategoryOauthConnectionCreated EventSubscriptionSelectedEventCategory = "oauth_connection.created"
	// Occurs whenever an OAuth Connection is deactivated.
	EventSubscriptionSelectedEventCategoryOauthConnectionDeactivated EventSubscriptionSelectedEventCategory = "oauth_connection.deactivated"
	// Occurs whenever a Pending Transaction is created.
	EventSubscriptionSelectedEventCategoryPendingTransactionCreated EventSubscriptionSelectedEventCategory = "pending_transaction.created"
	// Occurs whenever a Pending Transaction is updated.
	EventSubscriptionSelectedEventCategoryPendingTransactionUpdated EventSubscriptionSelectedEventCategory = "pending_transaction.updated"
	// Occurs whenever a Physical Card is created.
	EventSubscriptionSelectedEventCategoryPhysicalCardCreated EventSubscriptionSelectedEventCategory = "physical_card.created"
	// Occurs whenever a Physical Card is updated.
	EventSubscriptionSelectedEventCategoryPhysicalCardUpdated EventSubscriptionSelectedEventCategory = "physical_card.updated"
	// Occurs whenever a Real-Time Decision is created in response to a card
	// authorization.
	EventSubscriptionSelectedEventCategoryRealTimeDecisionCardAuthorizationRequested EventSubscriptionSelectedEventCategory = "real_time_decision.card_authorization_requested"
	// Occurs whenever a Real-Time Decision is created in response to a digital wallet
	// provisioning attempt.
	EventSubscriptionSelectedEventCategoryRealTimeDecisionDigitalWalletTokenRequested EventSubscriptionSelectedEventCategory = "real_time_decision.digital_wallet_token_requested"
	// Occurs whenever a Real-Time Decision is created in response to a digital wallet
	// requiring two-factor authentication.
	EventSubscriptionSelectedEventCategoryRealTimeDecisionDigitalWalletAuthenticationRequested EventSubscriptionSelectedEventCategory = "real_time_decision.digital_wallet_authentication_requested"
	// Occurs whenever a Real-Time Payments Transfer is created.
	EventSubscriptionSelectedEventCategoryRealTimePaymentsTransferCreated EventSubscriptionSelectedEventCategory = "real_time_payments_transfer.created"
	// Occurs whenever a Real-Time Payments Transfer is updated.
	EventSubscriptionSelectedEventCategoryRealTimePaymentsTransferUpdated EventSubscriptionSelectedEventCategory = "real_time_payments_transfer.updated"
	// Occurs whenever a Real-Time Payments Request for Payment is created.
	EventSubscriptionSelectedEventCategoryRealTimePaymentsRequestForPaymentCreated EventSubscriptionSelectedEventCategory = "real_time_payments_request_for_payment.created"
	// Occurs whenever a Real-Time Payments Request for Payment is updated.
	EventSubscriptionSelectedEventCategoryRealTimePaymentsRequestForPaymentUpdated EventSubscriptionSelectedEventCategory = "real_time_payments_request_for_payment.updated"
	// Occurs whenever a Transaction is created.
	EventSubscriptionSelectedEventCategoryTransactionCreated EventSubscriptionSelectedEventCategory = "transaction.created"
	// Occurs whenever a Wire Drawdown Request is created.
	EventSubscriptionSelectedEventCategoryWireDrawdownRequestCreated EventSubscriptionSelectedEventCategory = "wire_drawdown_request.created"
	// Occurs whenever a Wire Drawdown Request is updated.
	EventSubscriptionSelectedEventCategoryWireDrawdownRequestUpdated EventSubscriptionSelectedEventCategory = "wire_drawdown_request.updated"
	// Occurs whenever a Wire Transfer is created.
	EventSubscriptionSelectedEventCategoryWireTransferCreated EventSubscriptionSelectedEventCategory = "wire_transfer.created"
	// Occurs whenever a Wire Transfer is updated.
	EventSubscriptionSelectedEventCategoryWireTransferUpdated EventSubscriptionSelectedEventCategory = "wire_transfer.updated"
)

// This indicates if we'll send notifications to this subscription.
type EventSubscriptionStatus string

const (
	// The subscription is active and Events will be delivered normally.
	EventSubscriptionStatusActive EventSubscriptionStatus = "active"
	// The subscription is temporarily disabled and Events will not be delivered.
	EventSubscriptionStatusDisabled EventSubscriptionStatus = "disabled"
	// The subscription is permanently disabled and Events will not be delivered.
	EventSubscriptionStatusDeleted EventSubscriptionStatus = "deleted"
	// The subscription is temporarily disabled due to delivery errors and Events will
	// not be delivered.
	EventSubscriptionStatusRequiresAttention EventSubscriptionStatus = "requires_attention"
)

// A constant representing the object's type. For this resource it will always be
// `event_subscription`.
type EventSubscriptionType string

const (
	EventSubscriptionTypeEventSubscription EventSubscriptionType = "event_subscription"
)

type EventSubscriptionNewParams struct {
	// The URL you'd like us to send webhooks to.
	URL param.Field[string] `json:"url,required"`
	// If specified, this subscription will only receive webhooks for Events with the
	// specified `category`.
	SelectedEventCategory param.Field[EventSubscriptionNewParamsSelectedEventCategory] `json:"selected_event_category"`
	// The key that will be used to sign webhooks. If no value is passed, a random
	// string will be used as default.
	SharedSecret param.Field[string] `json:"shared_secret"`
}

func (r EventSubscriptionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If specified, this subscription will only receive webhooks for Events with the
// specified `category`.
type EventSubscriptionNewParamsSelectedEventCategory string

const (
	// Occurs whenever an Account is created.
	EventSubscriptionNewParamsSelectedEventCategoryAccountCreated EventSubscriptionNewParamsSelectedEventCategory = "account.created"
	// Occurs whenever an Account is updated.
	EventSubscriptionNewParamsSelectedEventCategoryAccountUpdated EventSubscriptionNewParamsSelectedEventCategory = "account.updated"
	// Occurs whenever an Account Number is created.
	EventSubscriptionNewParamsSelectedEventCategoryAccountNumberCreated EventSubscriptionNewParamsSelectedEventCategory = "account_number.created"
	// Occurs whenever an Account Number is updated.
	EventSubscriptionNewParamsSelectedEventCategoryAccountNumberUpdated EventSubscriptionNewParamsSelectedEventCategory = "account_number.updated"
	// Occurs whenever an Account Statement is created.
	EventSubscriptionNewParamsSelectedEventCategoryAccountStatementCreated EventSubscriptionNewParamsSelectedEventCategory = "account_statement.created"
	// Occurs whenever an Account Transfer is created.
	EventSubscriptionNewParamsSelectedEventCategoryAccountTransferCreated EventSubscriptionNewParamsSelectedEventCategory = "account_transfer.created"
	// Occurs whenever an Account Transfer is updated.
	EventSubscriptionNewParamsSelectedEventCategoryAccountTransferUpdated EventSubscriptionNewParamsSelectedEventCategory = "account_transfer.updated"
	// Occurs whenever an ACH Prenotification is created.
	EventSubscriptionNewParamsSelectedEventCategoryACHPrenotificationCreated EventSubscriptionNewParamsSelectedEventCategory = "ach_prenotification.created"
	// Occurs whenever an ACH Prenotification is updated.
	EventSubscriptionNewParamsSelectedEventCategoryACHPrenotificationUpdated EventSubscriptionNewParamsSelectedEventCategory = "ach_prenotification.updated"
	// Occurs whenever an ACH Transfer is created.
	EventSubscriptionNewParamsSelectedEventCategoryACHTransferCreated EventSubscriptionNewParamsSelectedEventCategory = "ach_transfer.created"
	// Occurs whenever an ACH Transfer is updated.
	EventSubscriptionNewParamsSelectedEventCategoryACHTransferUpdated EventSubscriptionNewParamsSelectedEventCategory = "ach_transfer.updated"
	// Occurs whenever a Bookkeeping Account is created.
	EventSubscriptionNewParamsSelectedEventCategoryBookkeepingAccountCreated EventSubscriptionNewParamsSelectedEventCategory = "bookkeeping_account.created"
	// Occurs whenever a Bookkeeping Account is updated.
	EventSubscriptionNewParamsSelectedEventCategoryBookkeepingAccountUpdated EventSubscriptionNewParamsSelectedEventCategory = "bookkeeping_account.updated"
	// Occurs whenever a Bookkeeping Entry Set is created.
	EventSubscriptionNewParamsSelectedEventCategoryBookkeepingEntrySetUpdated EventSubscriptionNewParamsSelectedEventCategory = "bookkeeping_entry_set.updated"
	// Occurs whenever a Card is created.
	EventSubscriptionNewParamsSelectedEventCategoryCardCreated EventSubscriptionNewParamsSelectedEventCategory = "card.created"
	// Occurs whenever a Card is updated.
	EventSubscriptionNewParamsSelectedEventCategoryCardUpdated EventSubscriptionNewParamsSelectedEventCategory = "card.updated"
	// Occurs whenever a Card Payment is created.
	EventSubscriptionNewParamsSelectedEventCategoryCardPaymentCreated EventSubscriptionNewParamsSelectedEventCategory = "card_payment.created"
	// Occurs whenever a Card Payment is updated.
	EventSubscriptionNewParamsSelectedEventCategoryCardPaymentUpdated EventSubscriptionNewParamsSelectedEventCategory = "card_payment.updated"
	// Occurs whenever a Card Profile is created.
	EventSubscriptionNewParamsSelectedEventCategoryCardProfileCreated EventSubscriptionNewParamsSelectedEventCategory = "card_profile.created"
	// Occurs whenever a Card Profile is updated.
	EventSubscriptionNewParamsSelectedEventCategoryCardProfileUpdated EventSubscriptionNewParamsSelectedEventCategory = "card_profile.updated"
	// Occurs whenever a Card Dispute is created.
	EventSubscriptionNewParamsSelectedEventCategoryCardDisputeCreated EventSubscriptionNewParamsSelectedEventCategory = "card_dispute.created"
	// Occurs whenever a Card Dispute is updated.
	EventSubscriptionNewParamsSelectedEventCategoryCardDisputeUpdated EventSubscriptionNewParamsSelectedEventCategory = "card_dispute.updated"
	// Occurs whenever a Check Deposit is created.
	EventSubscriptionNewParamsSelectedEventCategoryCheckDepositCreated EventSubscriptionNewParamsSelectedEventCategory = "check_deposit.created"
	// Occurs whenever a Check Deposit is updated.
	EventSubscriptionNewParamsSelectedEventCategoryCheckDepositUpdated EventSubscriptionNewParamsSelectedEventCategory = "check_deposit.updated"
	// Occurs whenever a Check Transfer is created.
	EventSubscriptionNewParamsSelectedEventCategoryCheckTransferCreated EventSubscriptionNewParamsSelectedEventCategory = "check_transfer.created"
	// Occurs whenever a Check Transfer is updated.
	EventSubscriptionNewParamsSelectedEventCategoryCheckTransferUpdated EventSubscriptionNewParamsSelectedEventCategory = "check_transfer.updated"
	// Occurs whenever a Declined Transaction is created.
	EventSubscriptionNewParamsSelectedEventCategoryDeclinedTransactionCreated EventSubscriptionNewParamsSelectedEventCategory = "declined_transaction.created"
	// Occurs whenever a Digital Wallet Token is created.
	EventSubscriptionNewParamsSelectedEventCategoryDigitalWalletTokenCreated EventSubscriptionNewParamsSelectedEventCategory = "digital_wallet_token.created"
	// Occurs whenever a Digital Wallet Token is updated.
	EventSubscriptionNewParamsSelectedEventCategoryDigitalWalletTokenUpdated EventSubscriptionNewParamsSelectedEventCategory = "digital_wallet_token.updated"
	// Occurs whenever a Document is created.
	EventSubscriptionNewParamsSelectedEventCategoryDocumentCreated EventSubscriptionNewParamsSelectedEventCategory = "document.created"
	// Occurs whenever an Entity is created.
	EventSubscriptionNewParamsSelectedEventCategoryEntityCreated EventSubscriptionNewParamsSelectedEventCategory = "entity.created"
	// Occurs whenever an Entity is updated.
	EventSubscriptionNewParamsSelectedEventCategoryEntityUpdated EventSubscriptionNewParamsSelectedEventCategory = "entity.updated"
	// Occurs whenever an Event Subscription is created.
	EventSubscriptionNewParamsSelectedEventCategoryEventSubscriptionCreated EventSubscriptionNewParamsSelectedEventCategory = "event_subscription.created"
	// Occurs whenever an Event Subscription is updated.
	EventSubscriptionNewParamsSelectedEventCategoryEventSubscriptionUpdated EventSubscriptionNewParamsSelectedEventCategory = "event_subscription.updated"
	// Occurs whenever an Export is created.
	EventSubscriptionNewParamsSelectedEventCategoryExportCreated EventSubscriptionNewParamsSelectedEventCategory = "export.created"
	// Occurs whenever an Export is updated.
	EventSubscriptionNewParamsSelectedEventCategoryExportUpdated EventSubscriptionNewParamsSelectedEventCategory = "export.updated"
	// Occurs whenever an External Account is created.
	EventSubscriptionNewParamsSelectedEventCategoryExternalAccountCreated EventSubscriptionNewParamsSelectedEventCategory = "external_account.created"
	// Occurs whenever an External Account is updated.
	EventSubscriptionNewParamsSelectedEventCategoryExternalAccountUpdated EventSubscriptionNewParamsSelectedEventCategory = "external_account.updated"
	// Occurs whenever a File is created.
	EventSubscriptionNewParamsSelectedEventCategoryFileCreated EventSubscriptionNewParamsSelectedEventCategory = "file.created"
	// Occurs whenever a Group is updated.
	EventSubscriptionNewParamsSelectedEventCategoryGroupUpdated EventSubscriptionNewParamsSelectedEventCategory = "group.updated"
	// Increase may send webhooks with this category to see if a webhook endpoint is
	// working properly.
	EventSubscriptionNewParamsSelectedEventCategoryGroupHeartbeat EventSubscriptionNewParamsSelectedEventCategory = "group.heartbeat"
	// Occurs whenever an Inbound ACH Transfer is created.
	EventSubscriptionNewParamsSelectedEventCategoryInboundACHTransferCreated EventSubscriptionNewParamsSelectedEventCategory = "inbound_ach_transfer.created"
	// Occurs whenever an Inbound ACH Transfer is updated.
	EventSubscriptionNewParamsSelectedEventCategoryInboundACHTransferUpdated EventSubscriptionNewParamsSelectedEventCategory = "inbound_ach_transfer.updated"
	// Occurs whenever an Inbound ACH Transfer Return is created.
	EventSubscriptionNewParamsSelectedEventCategoryInboundACHTransferReturnCreated EventSubscriptionNewParamsSelectedEventCategory = "inbound_ach_transfer_return.created"
	// Occurs whenever an Inbound ACH Transfer Return is updated.
	EventSubscriptionNewParamsSelectedEventCategoryInboundACHTransferReturnUpdated EventSubscriptionNewParamsSelectedEventCategory = "inbound_ach_transfer_return.updated"
	// Occurs whenever an Inbound Wire Drawdown Request is created.
	EventSubscriptionNewParamsSelectedEventCategoryInboundWireDrawdownRequestCreated EventSubscriptionNewParamsSelectedEventCategory = "inbound_wire_drawdown_request.created"
	// Occurs whenever an IntraFi Account Enrollment is created.
	EventSubscriptionNewParamsSelectedEventCategoryIntrafiAccountEnrollmentCreated EventSubscriptionNewParamsSelectedEventCategory = "intrafi_account_enrollment.created"
	// Occurs whenever an IntraFi Account Enrollment is updated.
	EventSubscriptionNewParamsSelectedEventCategoryIntrafiAccountEnrollmentUpdated EventSubscriptionNewParamsSelectedEventCategory = "intrafi_account_enrollment.updated"
	// Occurs whenever an IntraFi Exclusion is created.
	EventSubscriptionNewParamsSelectedEventCategoryIntrafiExclusionCreated EventSubscriptionNewParamsSelectedEventCategory = "intrafi_exclusion.created"
	// Occurs whenever an IntraFi Exclusion is updated.
	EventSubscriptionNewParamsSelectedEventCategoryIntrafiExclusionUpdated EventSubscriptionNewParamsSelectedEventCategory = "intrafi_exclusion.updated"
	// Occurs whenever an OAuth Connection is created.
	EventSubscriptionNewParamsSelectedEventCategoryOauthConnectionCreated EventSubscriptionNewParamsSelectedEventCategory = "oauth_connection.created"
	// Occurs whenever an OAuth Connection is deactivated.
	EventSubscriptionNewParamsSelectedEventCategoryOauthConnectionDeactivated EventSubscriptionNewParamsSelectedEventCategory = "oauth_connection.deactivated"
	// Occurs whenever a Pending Transaction is created.
	EventSubscriptionNewParamsSelectedEventCategoryPendingTransactionCreated EventSubscriptionNewParamsSelectedEventCategory = "pending_transaction.created"
	// Occurs whenever a Pending Transaction is updated.
	EventSubscriptionNewParamsSelectedEventCategoryPendingTransactionUpdated EventSubscriptionNewParamsSelectedEventCategory = "pending_transaction.updated"
	// Occurs whenever a Physical Card is created.
	EventSubscriptionNewParamsSelectedEventCategoryPhysicalCardCreated EventSubscriptionNewParamsSelectedEventCategory = "physical_card.created"
	// Occurs whenever a Physical Card is updated.
	EventSubscriptionNewParamsSelectedEventCategoryPhysicalCardUpdated EventSubscriptionNewParamsSelectedEventCategory = "physical_card.updated"
	// Occurs whenever a Real-Time Decision is created in response to a card
	// authorization.
	EventSubscriptionNewParamsSelectedEventCategoryRealTimeDecisionCardAuthorizationRequested EventSubscriptionNewParamsSelectedEventCategory = "real_time_decision.card_authorization_requested"
	// Occurs whenever a Real-Time Decision is created in response to a digital wallet
	// provisioning attempt.
	EventSubscriptionNewParamsSelectedEventCategoryRealTimeDecisionDigitalWalletTokenRequested EventSubscriptionNewParamsSelectedEventCategory = "real_time_decision.digital_wallet_token_requested"
	// Occurs whenever a Real-Time Decision is created in response to a digital wallet
	// requiring two-factor authentication.
	EventSubscriptionNewParamsSelectedEventCategoryRealTimeDecisionDigitalWalletAuthenticationRequested EventSubscriptionNewParamsSelectedEventCategory = "real_time_decision.digital_wallet_authentication_requested"
	// Occurs whenever a Real-Time Payments Transfer is created.
	EventSubscriptionNewParamsSelectedEventCategoryRealTimePaymentsTransferCreated EventSubscriptionNewParamsSelectedEventCategory = "real_time_payments_transfer.created"
	// Occurs whenever a Real-Time Payments Transfer is updated.
	EventSubscriptionNewParamsSelectedEventCategoryRealTimePaymentsTransferUpdated EventSubscriptionNewParamsSelectedEventCategory = "real_time_payments_transfer.updated"
	// Occurs whenever a Real-Time Payments Request for Payment is created.
	EventSubscriptionNewParamsSelectedEventCategoryRealTimePaymentsRequestForPaymentCreated EventSubscriptionNewParamsSelectedEventCategory = "real_time_payments_request_for_payment.created"
	// Occurs whenever a Real-Time Payments Request for Payment is updated.
	EventSubscriptionNewParamsSelectedEventCategoryRealTimePaymentsRequestForPaymentUpdated EventSubscriptionNewParamsSelectedEventCategory = "real_time_payments_request_for_payment.updated"
	// Occurs whenever a Transaction is created.
	EventSubscriptionNewParamsSelectedEventCategoryTransactionCreated EventSubscriptionNewParamsSelectedEventCategory = "transaction.created"
	// Occurs whenever a Wire Drawdown Request is created.
	EventSubscriptionNewParamsSelectedEventCategoryWireDrawdownRequestCreated EventSubscriptionNewParamsSelectedEventCategory = "wire_drawdown_request.created"
	// Occurs whenever a Wire Drawdown Request is updated.
	EventSubscriptionNewParamsSelectedEventCategoryWireDrawdownRequestUpdated EventSubscriptionNewParamsSelectedEventCategory = "wire_drawdown_request.updated"
	// Occurs whenever a Wire Transfer is created.
	EventSubscriptionNewParamsSelectedEventCategoryWireTransferCreated EventSubscriptionNewParamsSelectedEventCategory = "wire_transfer.created"
	// Occurs whenever a Wire Transfer is updated.
	EventSubscriptionNewParamsSelectedEventCategoryWireTransferUpdated EventSubscriptionNewParamsSelectedEventCategory = "wire_transfer.updated"
)

type EventSubscriptionUpdateParams struct {
	// The status to update the Event Subscription with.
	Status param.Field[EventSubscriptionUpdateParamsStatus] `json:"status"`
}

func (r EventSubscriptionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The status to update the Event Subscription with.
type EventSubscriptionUpdateParamsStatus string

const (
	// The subscription is active and Events will be delivered normally.
	EventSubscriptionUpdateParamsStatusActive EventSubscriptionUpdateParamsStatus = "active"
	// The subscription is temporarily disabled and Events will not be delivered.
	EventSubscriptionUpdateParamsStatusDisabled EventSubscriptionUpdateParamsStatus = "disabled"
	// The subscription is permanently disabled and Events will not be delivered.
	EventSubscriptionUpdateParamsStatusDeleted EventSubscriptionUpdateParamsStatus = "deleted"
)

type EventSubscriptionListParams struct {
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [EventSubscriptionListParams]'s query parameters as
// `url.Values`.
func (r EventSubscriptionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

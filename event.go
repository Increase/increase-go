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

// The category of the Event. We may add additional possible values for this enum
// over time; your application should be able to handle such additions gracefully.
type EventCategory string

const (
	// Occurs whenever an Account is created.
	EventCategoryAccountCreated EventCategory = "account.created"
	// Occurs whenever an Account is updated.
	EventCategoryAccountUpdated EventCategory = "account.updated"
	// Occurs whenever an Account Number is created.
	EventCategoryAccountNumberCreated EventCategory = "account_number.created"
	// Occurs whenever an Account Number is updated.
	EventCategoryAccountNumberUpdated EventCategory = "account_number.updated"
	// Occurs whenever an Account Statement is created.
	EventCategoryAccountStatementCreated EventCategory = "account_statement.created"
	// Occurs whenever an Account Transfer is created.
	EventCategoryAccountTransferCreated EventCategory = "account_transfer.created"
	// Occurs whenever an Account Transfer is updated.
	EventCategoryAccountTransferUpdated EventCategory = "account_transfer.updated"
	// Occurs whenever an ACH Prenotification is created.
	EventCategoryACHPrenotificationCreated EventCategory = "ach_prenotification.created"
	// Occurs whenever an ACH Prenotification is updated.
	EventCategoryACHPrenotificationUpdated EventCategory = "ach_prenotification.updated"
	// Occurs whenever an ACH Transfer is created.
	EventCategoryACHTransferCreated EventCategory = "ach_transfer.created"
	// Occurs whenever an ACH Transfer is updated.
	EventCategoryACHTransferUpdated EventCategory = "ach_transfer.updated"
	// Occurs whenever a Bookkeeping Account is created.
	EventCategoryBookkeepingAccountCreated EventCategory = "bookkeeping_account.created"
	// Occurs whenever a Bookkeeping Account is updated.
	EventCategoryBookkeepingAccountUpdated EventCategory = "bookkeeping_account.updated"
	// Occurs whenever a Bookkeeping Entry Set is created.
	EventCategoryBookkeepingEntrySetUpdated EventCategory = "bookkeeping_entry_set.updated"
	// Occurs whenever a Card is created.
	EventCategoryCardCreated EventCategory = "card.created"
	// Occurs whenever a Card is updated.
	EventCategoryCardUpdated EventCategory = "card.updated"
	// Occurs whenever a Card Payment is created.
	EventCategoryCardPaymentCreated EventCategory = "card_payment.created"
	// Occurs whenever a Card Payment is updated.
	EventCategoryCardPaymentUpdated EventCategory = "card_payment.updated"
	// Occurs whenever a Card Profile is created.
	EventCategoryCardProfileCreated EventCategory = "card_profile.created"
	// Occurs whenever a Card Profile is updated.
	EventCategoryCardProfileUpdated EventCategory = "card_profile.updated"
	// Occurs whenever a Card Dispute is created.
	EventCategoryCardDisputeCreated EventCategory = "card_dispute.created"
	// Occurs whenever a Card Dispute is updated.
	EventCategoryCardDisputeUpdated EventCategory = "card_dispute.updated"
	// Occurs whenever a Check Deposit is created.
	EventCategoryCheckDepositCreated EventCategory = "check_deposit.created"
	// Occurs whenever a Check Deposit is updated.
	EventCategoryCheckDepositUpdated EventCategory = "check_deposit.updated"
	// Occurs whenever a Check Transfer is created.
	EventCategoryCheckTransferCreated EventCategory = "check_transfer.created"
	// Occurs whenever a Check Transfer is updated.
	EventCategoryCheckTransferUpdated EventCategory = "check_transfer.updated"
	// Occurs whenever a Declined Transaction is created.
	EventCategoryDeclinedTransactionCreated EventCategory = "declined_transaction.created"
	// Occurs whenever a Digital Wallet Token is created.
	EventCategoryDigitalWalletTokenCreated EventCategory = "digital_wallet_token.created"
	// Occurs whenever a Digital Wallet Token is updated.
	EventCategoryDigitalWalletTokenUpdated EventCategory = "digital_wallet_token.updated"
	// Occurs whenever a Document is created.
	EventCategoryDocumentCreated EventCategory = "document.created"
	// Occurs whenever an Entity is created.
	EventCategoryEntityCreated EventCategory = "entity.created"
	// Occurs whenever an Entity is updated.
	EventCategoryEntityUpdated EventCategory = "entity.updated"
	// Occurs whenever an Event Subscription is created.
	EventCategoryEventSubscriptionCreated EventCategory = "event_subscription.created"
	// Occurs whenever an Event Subscription is updated.
	EventCategoryEventSubscriptionUpdated EventCategory = "event_subscription.updated"
	// Occurs whenever an Export is created.
	EventCategoryExportCreated EventCategory = "export.created"
	// Occurs whenever an Export is updated.
	EventCategoryExportUpdated EventCategory = "export.updated"
	// Occurs whenever an External Account is created.
	EventCategoryExternalAccountCreated EventCategory = "external_account.created"
	// Occurs whenever an External Account is updated.
	EventCategoryExternalAccountUpdated EventCategory = "external_account.updated"
	// Occurs whenever a File is created.
	EventCategoryFileCreated EventCategory = "file.created"
	// Occurs whenever a Group is updated.
	EventCategoryGroupUpdated EventCategory = "group.updated"
	// Increase may send webhooks with this category to see if a webhook endpoint is
	// working properly.
	EventCategoryGroupHeartbeat EventCategory = "group.heartbeat"
	// Occurs whenever an Inbound ACH Transfer is created.
	EventCategoryInboundACHTransferCreated EventCategory = "inbound_ach_transfer.created"
	// Occurs whenever an Inbound ACH Transfer is updated.
	EventCategoryInboundACHTransferUpdated EventCategory = "inbound_ach_transfer.updated"
	// Occurs whenever an Inbound ACH Transfer Return is created.
	EventCategoryInboundACHTransferReturnCreated EventCategory = "inbound_ach_transfer_return.created"
	// Occurs whenever an Inbound ACH Transfer Return is updated.
	EventCategoryInboundACHTransferReturnUpdated EventCategory = "inbound_ach_transfer_return.updated"
	// Occurs whenever an Inbound Wire Drawdown Request is created.
	EventCategoryInboundWireDrawdownRequestCreated EventCategory = "inbound_wire_drawdown_request.created"
	// Occurs whenever an IntraFi Account Enrollment is created.
	EventCategoryIntrafiAccountEnrollmentCreated EventCategory = "intrafi_account_enrollment.created"
	// Occurs whenever an IntraFi Account Enrollment is updated.
	EventCategoryIntrafiAccountEnrollmentUpdated EventCategory = "intrafi_account_enrollment.updated"
	// Occurs whenever an IntraFi Exclusion is created.
	EventCategoryIntrafiExclusionCreated EventCategory = "intrafi_exclusion.created"
	// Occurs whenever an IntraFi Exclusion is updated.
	EventCategoryIntrafiExclusionUpdated EventCategory = "intrafi_exclusion.updated"
	// Occurs whenever an OAuth Connection is created.
	EventCategoryOauthConnectionCreated EventCategory = "oauth_connection.created"
	// Occurs whenever an OAuth Connection is deactivated.
	EventCategoryOauthConnectionDeactivated EventCategory = "oauth_connection.deactivated"
	// Occurs whenever a Pending Transaction is created.
	EventCategoryPendingTransactionCreated EventCategory = "pending_transaction.created"
	// Occurs whenever a Pending Transaction is updated.
	EventCategoryPendingTransactionUpdated EventCategory = "pending_transaction.updated"
	// Occurs whenever a Physical Card is created.
	EventCategoryPhysicalCardCreated EventCategory = "physical_card.created"
	// Occurs whenever a Physical Card is updated.
	EventCategoryPhysicalCardUpdated EventCategory = "physical_card.updated"
	// Occurs whenever a Real-Time Decision is created in response to a card
	// authorization.
	EventCategoryRealTimeDecisionCardAuthorizationRequested EventCategory = "real_time_decision.card_authorization_requested"
	// Occurs whenever a Real-Time Decision is created in response to a digital wallet
	// provisioning attempt.
	EventCategoryRealTimeDecisionDigitalWalletTokenRequested EventCategory = "real_time_decision.digital_wallet_token_requested"
	// Occurs whenever a Real-Time Decision is created in response to a digital wallet
	// requiring two-factor authentication.
	EventCategoryRealTimeDecisionDigitalWalletAuthenticationRequested EventCategory = "real_time_decision.digital_wallet_authentication_requested"
	// Occurs whenever a Real-Time Payments Transfer is created.
	EventCategoryRealTimePaymentsTransferCreated EventCategory = "real_time_payments_transfer.created"
	// Occurs whenever a Real-Time Payments Transfer is updated.
	EventCategoryRealTimePaymentsTransferUpdated EventCategory = "real_time_payments_transfer.updated"
	// Occurs whenever a Real-Time Payments Request for Payment is created.
	EventCategoryRealTimePaymentsRequestForPaymentCreated EventCategory = "real_time_payments_request_for_payment.created"
	// Occurs whenever a Real-Time Payments Request for Payment is updated.
	EventCategoryRealTimePaymentsRequestForPaymentUpdated EventCategory = "real_time_payments_request_for_payment.updated"
	// Occurs whenever a Transaction is created.
	EventCategoryTransactionCreated EventCategory = "transaction.created"
	// Occurs whenever a Wire Drawdown Request is created.
	EventCategoryWireDrawdownRequestCreated EventCategory = "wire_drawdown_request.created"
	// Occurs whenever a Wire Drawdown Request is updated.
	EventCategoryWireDrawdownRequestUpdated EventCategory = "wire_drawdown_request.updated"
	// Occurs whenever a Wire Transfer is created.
	EventCategoryWireTransferCreated EventCategory = "wire_transfer.created"
	// Occurs whenever a Wire Transfer is updated.
	EventCategoryWireTransferUpdated EventCategory = "wire_transfer.updated"
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
	// Occurs whenever an Account is created.
	EventListParamsCategoryInAccountCreated EventListParamsCategoryIn = "account.created"
	// Occurs whenever an Account is updated.
	EventListParamsCategoryInAccountUpdated EventListParamsCategoryIn = "account.updated"
	// Occurs whenever an Account Number is created.
	EventListParamsCategoryInAccountNumberCreated EventListParamsCategoryIn = "account_number.created"
	// Occurs whenever an Account Number is updated.
	EventListParamsCategoryInAccountNumberUpdated EventListParamsCategoryIn = "account_number.updated"
	// Occurs whenever an Account Statement is created.
	EventListParamsCategoryInAccountStatementCreated EventListParamsCategoryIn = "account_statement.created"
	// Occurs whenever an Account Transfer is created.
	EventListParamsCategoryInAccountTransferCreated EventListParamsCategoryIn = "account_transfer.created"
	// Occurs whenever an Account Transfer is updated.
	EventListParamsCategoryInAccountTransferUpdated EventListParamsCategoryIn = "account_transfer.updated"
	// Occurs whenever an ACH Prenotification is created.
	EventListParamsCategoryInACHPrenotificationCreated EventListParamsCategoryIn = "ach_prenotification.created"
	// Occurs whenever an ACH Prenotification is updated.
	EventListParamsCategoryInACHPrenotificationUpdated EventListParamsCategoryIn = "ach_prenotification.updated"
	// Occurs whenever an ACH Transfer is created.
	EventListParamsCategoryInACHTransferCreated EventListParamsCategoryIn = "ach_transfer.created"
	// Occurs whenever an ACH Transfer is updated.
	EventListParamsCategoryInACHTransferUpdated EventListParamsCategoryIn = "ach_transfer.updated"
	// Occurs whenever a Bookkeeping Account is created.
	EventListParamsCategoryInBookkeepingAccountCreated EventListParamsCategoryIn = "bookkeeping_account.created"
	// Occurs whenever a Bookkeeping Account is updated.
	EventListParamsCategoryInBookkeepingAccountUpdated EventListParamsCategoryIn = "bookkeeping_account.updated"
	// Occurs whenever a Bookkeeping Entry Set is created.
	EventListParamsCategoryInBookkeepingEntrySetUpdated EventListParamsCategoryIn = "bookkeeping_entry_set.updated"
	// Occurs whenever a Card is created.
	EventListParamsCategoryInCardCreated EventListParamsCategoryIn = "card.created"
	// Occurs whenever a Card is updated.
	EventListParamsCategoryInCardUpdated EventListParamsCategoryIn = "card.updated"
	// Occurs whenever a Card Payment is created.
	EventListParamsCategoryInCardPaymentCreated EventListParamsCategoryIn = "card_payment.created"
	// Occurs whenever a Card Payment is updated.
	EventListParamsCategoryInCardPaymentUpdated EventListParamsCategoryIn = "card_payment.updated"
	// Occurs whenever a Card Profile is created.
	EventListParamsCategoryInCardProfileCreated EventListParamsCategoryIn = "card_profile.created"
	// Occurs whenever a Card Profile is updated.
	EventListParamsCategoryInCardProfileUpdated EventListParamsCategoryIn = "card_profile.updated"
	// Occurs whenever a Card Dispute is created.
	EventListParamsCategoryInCardDisputeCreated EventListParamsCategoryIn = "card_dispute.created"
	// Occurs whenever a Card Dispute is updated.
	EventListParamsCategoryInCardDisputeUpdated EventListParamsCategoryIn = "card_dispute.updated"
	// Occurs whenever a Check Deposit is created.
	EventListParamsCategoryInCheckDepositCreated EventListParamsCategoryIn = "check_deposit.created"
	// Occurs whenever a Check Deposit is updated.
	EventListParamsCategoryInCheckDepositUpdated EventListParamsCategoryIn = "check_deposit.updated"
	// Occurs whenever a Check Transfer is created.
	EventListParamsCategoryInCheckTransferCreated EventListParamsCategoryIn = "check_transfer.created"
	// Occurs whenever a Check Transfer is updated.
	EventListParamsCategoryInCheckTransferUpdated EventListParamsCategoryIn = "check_transfer.updated"
	// Occurs whenever a Declined Transaction is created.
	EventListParamsCategoryInDeclinedTransactionCreated EventListParamsCategoryIn = "declined_transaction.created"
	// Occurs whenever a Digital Wallet Token is created.
	EventListParamsCategoryInDigitalWalletTokenCreated EventListParamsCategoryIn = "digital_wallet_token.created"
	// Occurs whenever a Digital Wallet Token is updated.
	EventListParamsCategoryInDigitalWalletTokenUpdated EventListParamsCategoryIn = "digital_wallet_token.updated"
	// Occurs whenever a Document is created.
	EventListParamsCategoryInDocumentCreated EventListParamsCategoryIn = "document.created"
	// Occurs whenever an Entity is created.
	EventListParamsCategoryInEntityCreated EventListParamsCategoryIn = "entity.created"
	// Occurs whenever an Entity is updated.
	EventListParamsCategoryInEntityUpdated EventListParamsCategoryIn = "entity.updated"
	// Occurs whenever an Event Subscription is created.
	EventListParamsCategoryInEventSubscriptionCreated EventListParamsCategoryIn = "event_subscription.created"
	// Occurs whenever an Event Subscription is updated.
	EventListParamsCategoryInEventSubscriptionUpdated EventListParamsCategoryIn = "event_subscription.updated"
	// Occurs whenever an Export is created.
	EventListParamsCategoryInExportCreated EventListParamsCategoryIn = "export.created"
	// Occurs whenever an Export is updated.
	EventListParamsCategoryInExportUpdated EventListParamsCategoryIn = "export.updated"
	// Occurs whenever an External Account is created.
	EventListParamsCategoryInExternalAccountCreated EventListParamsCategoryIn = "external_account.created"
	// Occurs whenever an External Account is updated.
	EventListParamsCategoryInExternalAccountUpdated EventListParamsCategoryIn = "external_account.updated"
	// Occurs whenever a File is created.
	EventListParamsCategoryInFileCreated EventListParamsCategoryIn = "file.created"
	// Occurs whenever a Group is updated.
	EventListParamsCategoryInGroupUpdated EventListParamsCategoryIn = "group.updated"
	// Increase may send webhooks with this category to see if a webhook endpoint is
	// working properly.
	EventListParamsCategoryInGroupHeartbeat EventListParamsCategoryIn = "group.heartbeat"
	// Occurs whenever an Inbound ACH Transfer is created.
	EventListParamsCategoryInInboundACHTransferCreated EventListParamsCategoryIn = "inbound_ach_transfer.created"
	// Occurs whenever an Inbound ACH Transfer is updated.
	EventListParamsCategoryInInboundACHTransferUpdated EventListParamsCategoryIn = "inbound_ach_transfer.updated"
	// Occurs whenever an Inbound ACH Transfer Return is created.
	EventListParamsCategoryInInboundACHTransferReturnCreated EventListParamsCategoryIn = "inbound_ach_transfer_return.created"
	// Occurs whenever an Inbound ACH Transfer Return is updated.
	EventListParamsCategoryInInboundACHTransferReturnUpdated EventListParamsCategoryIn = "inbound_ach_transfer_return.updated"
	// Occurs whenever an Inbound Wire Drawdown Request is created.
	EventListParamsCategoryInInboundWireDrawdownRequestCreated EventListParamsCategoryIn = "inbound_wire_drawdown_request.created"
	// Occurs whenever an IntraFi Account Enrollment is created.
	EventListParamsCategoryInIntrafiAccountEnrollmentCreated EventListParamsCategoryIn = "intrafi_account_enrollment.created"
	// Occurs whenever an IntraFi Account Enrollment is updated.
	EventListParamsCategoryInIntrafiAccountEnrollmentUpdated EventListParamsCategoryIn = "intrafi_account_enrollment.updated"
	// Occurs whenever an IntraFi Exclusion is created.
	EventListParamsCategoryInIntrafiExclusionCreated EventListParamsCategoryIn = "intrafi_exclusion.created"
	// Occurs whenever an IntraFi Exclusion is updated.
	EventListParamsCategoryInIntrafiExclusionUpdated EventListParamsCategoryIn = "intrafi_exclusion.updated"
	// Occurs whenever an OAuth Connection is created.
	EventListParamsCategoryInOauthConnectionCreated EventListParamsCategoryIn = "oauth_connection.created"
	// Occurs whenever an OAuth Connection is deactivated.
	EventListParamsCategoryInOauthConnectionDeactivated EventListParamsCategoryIn = "oauth_connection.deactivated"
	// Occurs whenever a Pending Transaction is created.
	EventListParamsCategoryInPendingTransactionCreated EventListParamsCategoryIn = "pending_transaction.created"
	// Occurs whenever a Pending Transaction is updated.
	EventListParamsCategoryInPendingTransactionUpdated EventListParamsCategoryIn = "pending_transaction.updated"
	// Occurs whenever a Physical Card is created.
	EventListParamsCategoryInPhysicalCardCreated EventListParamsCategoryIn = "physical_card.created"
	// Occurs whenever a Physical Card is updated.
	EventListParamsCategoryInPhysicalCardUpdated EventListParamsCategoryIn = "physical_card.updated"
	// Occurs whenever a Real-Time Decision is created in response to a card
	// authorization.
	EventListParamsCategoryInRealTimeDecisionCardAuthorizationRequested EventListParamsCategoryIn = "real_time_decision.card_authorization_requested"
	// Occurs whenever a Real-Time Decision is created in response to a digital wallet
	// provisioning attempt.
	EventListParamsCategoryInRealTimeDecisionDigitalWalletTokenRequested EventListParamsCategoryIn = "real_time_decision.digital_wallet_token_requested"
	// Occurs whenever a Real-Time Decision is created in response to a digital wallet
	// requiring two-factor authentication.
	EventListParamsCategoryInRealTimeDecisionDigitalWalletAuthenticationRequested EventListParamsCategoryIn = "real_time_decision.digital_wallet_authentication_requested"
	// Occurs whenever a Real-Time Payments Transfer is created.
	EventListParamsCategoryInRealTimePaymentsTransferCreated EventListParamsCategoryIn = "real_time_payments_transfer.created"
	// Occurs whenever a Real-Time Payments Transfer is updated.
	EventListParamsCategoryInRealTimePaymentsTransferUpdated EventListParamsCategoryIn = "real_time_payments_transfer.updated"
	// Occurs whenever a Real-Time Payments Request for Payment is created.
	EventListParamsCategoryInRealTimePaymentsRequestForPaymentCreated EventListParamsCategoryIn = "real_time_payments_request_for_payment.created"
	// Occurs whenever a Real-Time Payments Request for Payment is updated.
	EventListParamsCategoryInRealTimePaymentsRequestForPaymentUpdated EventListParamsCategoryIn = "real_time_payments_request_for_payment.updated"
	// Occurs whenever a Transaction is created.
	EventListParamsCategoryInTransactionCreated EventListParamsCategoryIn = "transaction.created"
	// Occurs whenever a Wire Drawdown Request is created.
	EventListParamsCategoryInWireDrawdownRequestCreated EventListParamsCategoryIn = "wire_drawdown_request.created"
	// Occurs whenever a Wire Drawdown Request is updated.
	EventListParamsCategoryInWireDrawdownRequestUpdated EventListParamsCategoryIn = "wire_drawdown_request.updated"
	// Occurs whenever a Wire Transfer is created.
	EventListParamsCategoryInWireTransferCreated EventListParamsCategoryIn = "wire_transfer.created"
	// Occurs whenever a Wire Transfer is updated.
	EventListParamsCategoryInWireTransferUpdated EventListParamsCategoryIn = "wire_transfer.updated"
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

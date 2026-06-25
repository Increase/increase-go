// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"net/http"
	"slices"

	"github.com/Increase/increase-go/internal/apijson"
	"github.com/Increase/increase-go/internal/param"
	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
)

// SimulationCardAuthorizationService contains methods and other services that help
// with interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimulationCardAuthorizationService] method instead.
type SimulationCardAuthorizationService struct {
	Options []option.RequestOption
}

// NewSimulationCardAuthorizationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSimulationCardAuthorizationService(opts ...option.RequestOption) (r *SimulationCardAuthorizationService) {
	r = &SimulationCardAuthorizationService{}
	r.Options = opts
	return
}

// Simulates a purchase authorization on a [Card](#cards). Depending on the balance
// available to the card and the `amount` submitted, the authorization activity
// will result in a [Pending Transaction](#pending-transactions) of type
// `card_authorization` or a [Declined Transaction](#declined-transactions) of type
// `card_decline`. You can pass either a Card id or a
// [Digital Wallet Token](#digital-wallet-tokens) id to simulate the two different
// ways purchases can be made. The response will contain either a
// `pending_transaction` or a `declined_transaction`; the other attribute will be
// null. If the authorization is declined, the reason is available on the Declined
// Transaction at `source.card_decline.reason`.
func (r *SimulationCardAuthorizationService) New(ctx context.Context, body SimulationCardAuthorizationNewParams, opts ...option.RequestOption) (res *SimulationCardAuthorizationNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "simulations/card_authorizations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// The results of a Card Authorization simulation.
type SimulationCardAuthorizationNewResponse struct {
	// If the authorization attempt fails, this will contain the resulting
	// [Declined Transaction](#declined-transactions) object. The Declined
	// Transaction's `source` will be of `category: card_decline`.
	DeclinedTransaction DeclinedTransaction `json:"declined_transaction" api:"required,nullable"`
	// If the authorization attempt succeeds, this will contain the resulting Pending
	// Transaction object. The Pending Transaction's `source` will be of
	// `category: card_authorization`.
	PendingTransaction PendingTransaction `json:"pending_transaction" api:"required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `inbound_card_authorization_simulation_result`.
	Type SimulationCardAuthorizationNewResponseType `json:"type" api:"required"`
	JSON simulationCardAuthorizationNewResponseJSON `json:"-"`
}

// simulationCardAuthorizationNewResponseJSON contains the JSON metadata for the
// struct [SimulationCardAuthorizationNewResponse]
type simulationCardAuthorizationNewResponseJSON struct {
	DeclinedTransaction apijson.Field
	PendingTransaction  apijson.Field
	Type                apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *SimulationCardAuthorizationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r simulationCardAuthorizationNewResponseJSON) RawJSON() string {
	return r.raw
}

// A constant representing the object's type. For this resource it will always be
// `inbound_card_authorization_simulation_result`.
type SimulationCardAuthorizationNewResponseType string

const (
	SimulationCardAuthorizationNewResponseTypeInboundCardAuthorizationSimulationResult SimulationCardAuthorizationNewResponseType = "inbound_card_authorization_simulation_result"
)

func (r SimulationCardAuthorizationNewResponseType) IsKnown() bool {
	switch r {
	case SimulationCardAuthorizationNewResponseTypeInboundCardAuthorizationSimulationResult:
		return true
	}
	return false
}

type SimulationCardAuthorizationNewParams struct {
	// The authorization amount in cents.
	Amount param.Field[int64] `json:"amount" api:"required"`
	// The identifier of a Card Payment with a `card_authentication` if you want to
	// simulate an authenticated authorization.
	AuthenticatedCardPaymentID param.Field[string] `json:"authenticated_card_payment_id"`
	// The identifier of the Card to be authorized.
	CardID param.Field[string] `json:"card_id"`
	// Forces a card decline with a specific reason. No real time decision will be
	// sent.
	DeclineReason param.Field[SimulationCardAuthorizationNewParamsDeclineReason] `json:"decline_reason"`
	// The identifier of the Digital Wallet Token to be authorized.
	DigitalWalletTokenID param.Field[string] `json:"digital_wallet_token_id"`
	// The identifier of the Event Subscription to use. If provided, will override the
	// default real time event subscription. Because you can only create one real time
	// decision event subscription, you can use this field to route events to any
	// specified event subscription for testing purposes.
	EventSubscriptionID param.Field[string] `json:"event_subscription_id"`
	// The merchant identifier (commonly abbreviated as MID) of the merchant the card
	// is transacting with.
	MerchantAcceptorID param.Field[string] `json:"merchant_acceptor_id"`
	// The Merchant Category Code (commonly abbreviated as MCC) of the merchant the
	// card is transacting with.
	MerchantCategoryCode param.Field[string] `json:"merchant_category_code"`
	// The city the merchant resides in.
	MerchantCity param.Field[string] `json:"merchant_city"`
	// The country the merchant resides in.
	MerchantCountry param.Field[string] `json:"merchant_country"`
	// The merchant descriptor of the merchant the card is transacting with.
	MerchantDescriptor param.Field[string] `json:"merchant_descriptor"`
	// The state the merchant resides in.
	MerchantState param.Field[string] `json:"merchant_state"`
	// Fields specific to a given card network.
	NetworkDetails param.Field[SimulationCardAuthorizationNewParamsNetworkDetails] `json:"network_details"`
	// The risk score generated by the card network. For Visa this is the Visa Advanced
	// Authorization risk score, from 0 to 99, where 99 is the riskiest.
	NetworkRiskScore param.Field[int64] `json:"network_risk_score"`
	// The identifier of the Physical Card to be authorized.
	PhysicalCardID param.Field[string] `json:"physical_card_id"`
	// Fields specific to a specific type of authorization, such as Automatic Fuel
	// Dispensers, Refund Authorizations, or Cash Disbursements.
	ProcessingCategory param.Field[SimulationCardAuthorizationNewParamsProcessingCategory] `json:"processing_category"`
	// The terminal identifier (commonly abbreviated as TID) of the terminal the card
	// is transacting with.
	TerminalID param.Field[string] `json:"terminal_id"`
}

func (r SimulationCardAuthorizationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Forces a card decline with a specific reason. No real time decision will be
// sent.
type SimulationCardAuthorizationNewParamsDeclineReason string

const (
	SimulationCardAuthorizationNewParamsDeclineReasonAccountClosed                SimulationCardAuthorizationNewParamsDeclineReason = "account_closed"
	SimulationCardAuthorizationNewParamsDeclineReasonCardNotActive                SimulationCardAuthorizationNewParamsDeclineReason = "card_not_active"
	SimulationCardAuthorizationNewParamsDeclineReasonCardCanceled                 SimulationCardAuthorizationNewParamsDeclineReason = "card_canceled"
	SimulationCardAuthorizationNewParamsDeclineReasonPhysicalCardNotActive        SimulationCardAuthorizationNewParamsDeclineReason = "physical_card_not_active"
	SimulationCardAuthorizationNewParamsDeclineReasonEntityNotActive              SimulationCardAuthorizationNewParamsDeclineReason = "entity_not_active"
	SimulationCardAuthorizationNewParamsDeclineReasonGroupLocked                  SimulationCardAuthorizationNewParamsDeclineReason = "group_locked"
	SimulationCardAuthorizationNewParamsDeclineReasonInsufficientFunds            SimulationCardAuthorizationNewParamsDeclineReason = "insufficient_funds"
	SimulationCardAuthorizationNewParamsDeclineReasonCvv2Mismatch                 SimulationCardAuthorizationNewParamsDeclineReason = "cvv2_mismatch"
	SimulationCardAuthorizationNewParamsDeclineReasonPinMismatch                  SimulationCardAuthorizationNewParamsDeclineReason = "pin_mismatch"
	SimulationCardAuthorizationNewParamsDeclineReasonCardExpirationMismatch       SimulationCardAuthorizationNewParamsDeclineReason = "card_expiration_mismatch"
	SimulationCardAuthorizationNewParamsDeclineReasonTransactionNotAllowed        SimulationCardAuthorizationNewParamsDeclineReason = "transaction_not_allowed"
	SimulationCardAuthorizationNewParamsDeclineReasonBreachesLimit                SimulationCardAuthorizationNewParamsDeclineReason = "breaches_limit"
	SimulationCardAuthorizationNewParamsDeclineReasonWebhookDeclined              SimulationCardAuthorizationNewParamsDeclineReason = "webhook_declined"
	SimulationCardAuthorizationNewParamsDeclineReasonWebhookTimedOut              SimulationCardAuthorizationNewParamsDeclineReason = "webhook_timed_out"
	SimulationCardAuthorizationNewParamsDeclineReasonDeclinedByStandInProcessing  SimulationCardAuthorizationNewParamsDeclineReason = "declined_by_stand_in_processing"
	SimulationCardAuthorizationNewParamsDeclineReasonInvalidPhysicalCard          SimulationCardAuthorizationNewParamsDeclineReason = "invalid_physical_card"
	SimulationCardAuthorizationNewParamsDeclineReasonMissingOriginalAuthorization SimulationCardAuthorizationNewParamsDeclineReason = "missing_original_authorization"
	SimulationCardAuthorizationNewParamsDeclineReasonInvalidCryptogram            SimulationCardAuthorizationNewParamsDeclineReason = "invalid_cryptogram"
	SimulationCardAuthorizationNewParamsDeclineReasonFailed3DSAuthentication      SimulationCardAuthorizationNewParamsDeclineReason = "failed_3ds_authentication"
	SimulationCardAuthorizationNewParamsDeclineReasonSuspectedCardTesting         SimulationCardAuthorizationNewParamsDeclineReason = "suspected_card_testing"
	SimulationCardAuthorizationNewParamsDeclineReasonSuspectedFraud               SimulationCardAuthorizationNewParamsDeclineReason = "suspected_fraud"
)

func (r SimulationCardAuthorizationNewParamsDeclineReason) IsKnown() bool {
	switch r {
	case SimulationCardAuthorizationNewParamsDeclineReasonAccountClosed, SimulationCardAuthorizationNewParamsDeclineReasonCardNotActive, SimulationCardAuthorizationNewParamsDeclineReasonCardCanceled, SimulationCardAuthorizationNewParamsDeclineReasonPhysicalCardNotActive, SimulationCardAuthorizationNewParamsDeclineReasonEntityNotActive, SimulationCardAuthorizationNewParamsDeclineReasonGroupLocked, SimulationCardAuthorizationNewParamsDeclineReasonInsufficientFunds, SimulationCardAuthorizationNewParamsDeclineReasonCvv2Mismatch, SimulationCardAuthorizationNewParamsDeclineReasonPinMismatch, SimulationCardAuthorizationNewParamsDeclineReasonCardExpirationMismatch, SimulationCardAuthorizationNewParamsDeclineReasonTransactionNotAllowed, SimulationCardAuthorizationNewParamsDeclineReasonBreachesLimit, SimulationCardAuthorizationNewParamsDeclineReasonWebhookDeclined, SimulationCardAuthorizationNewParamsDeclineReasonWebhookTimedOut, SimulationCardAuthorizationNewParamsDeclineReasonDeclinedByStandInProcessing, SimulationCardAuthorizationNewParamsDeclineReasonInvalidPhysicalCard, SimulationCardAuthorizationNewParamsDeclineReasonMissingOriginalAuthorization, SimulationCardAuthorizationNewParamsDeclineReasonInvalidCryptogram, SimulationCardAuthorizationNewParamsDeclineReasonFailed3DSAuthentication, SimulationCardAuthorizationNewParamsDeclineReasonSuspectedCardTesting, SimulationCardAuthorizationNewParamsDeclineReasonSuspectedFraud:
		return true
	}
	return false
}

// Fields specific to a given card network.
type SimulationCardAuthorizationNewParamsNetworkDetails struct {
	// Fields specific to the Visa network.
	Visa param.Field[SimulationCardAuthorizationNewParamsNetworkDetailsVisa] `json:"visa" api:"required"`
}

func (r SimulationCardAuthorizationNewParamsNetworkDetails) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Fields specific to the Visa network.
type SimulationCardAuthorizationNewParamsNetworkDetailsVisa struct {
	// For electronic commerce transactions, this identifies the level of security used
	// in obtaining the customer's payment credential. For mail or telephone order
	// transactions, identifies the type of mail or telephone order.
	ElectronicCommerceIndicator param.Field[SimulationCardAuthorizationNewParamsNetworkDetailsVisaElectronicCommerceIndicator] `json:"electronic_commerce_indicator"`
	// The method used to enter the cardholder's primary account number and card
	// expiration date.
	PointOfServiceEntryMode param.Field[SimulationCardAuthorizationNewParamsNetworkDetailsVisaPointOfServiceEntryMode] `json:"point_of_service_entry_mode"`
	// The reason code for the stand-in processing.
	StandInProcessingReason param.Field[SimulationCardAuthorizationNewParamsNetworkDetailsVisaStandInProcessingReason] `json:"stand_in_processing_reason"`
	// The capability of the terminal being used to read the card. Shows whether a
	// terminal can e.g., accept chip cards or if it only supports magnetic stripe
	// reads. This reflects the highest capability of the terminal — for example, a
	// terminal that supports both chip and magnetic stripe will be identified as
	// chip-capable.
	TerminalEntryCapability param.Field[SimulationCardAuthorizationNewParamsNetworkDetailsVisaTerminalEntryCapability] `json:"terminal_entry_capability"`
}

func (r SimulationCardAuthorizationNewParamsNetworkDetailsVisa) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For electronic commerce transactions, this identifies the level of security used
// in obtaining the customer's payment credential. For mail or telephone order
// transactions, identifies the type of mail or telephone order.
type SimulationCardAuthorizationNewParamsNetworkDetailsVisaElectronicCommerceIndicator string

const (
	SimulationCardAuthorizationNewParamsNetworkDetailsVisaElectronicCommerceIndicatorMailPhoneOrder                                          SimulationCardAuthorizationNewParamsNetworkDetailsVisaElectronicCommerceIndicator = "mail_phone_order"
	SimulationCardAuthorizationNewParamsNetworkDetailsVisaElectronicCommerceIndicatorRecurring                                               SimulationCardAuthorizationNewParamsNetworkDetailsVisaElectronicCommerceIndicator = "recurring"
	SimulationCardAuthorizationNewParamsNetworkDetailsVisaElectronicCommerceIndicatorInstallment                                             SimulationCardAuthorizationNewParamsNetworkDetailsVisaElectronicCommerceIndicator = "installment"
	SimulationCardAuthorizationNewParamsNetworkDetailsVisaElectronicCommerceIndicatorUnknownMailPhoneOrder                                   SimulationCardAuthorizationNewParamsNetworkDetailsVisaElectronicCommerceIndicator = "unknown_mail_phone_order"
	SimulationCardAuthorizationNewParamsNetworkDetailsVisaElectronicCommerceIndicatorSecureElectronicCommerce                                SimulationCardAuthorizationNewParamsNetworkDetailsVisaElectronicCommerceIndicator = "secure_electronic_commerce"
	SimulationCardAuthorizationNewParamsNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransactionAt3DSCapableMerchant SimulationCardAuthorizationNewParamsNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction_at_3ds_capable_merchant"
	SimulationCardAuthorizationNewParamsNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransaction                     SimulationCardAuthorizationNewParamsNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction"
	SimulationCardAuthorizationNewParamsNetworkDetailsVisaElectronicCommerceIndicatorNonSecureTransaction                                    SimulationCardAuthorizationNewParamsNetworkDetailsVisaElectronicCommerceIndicator = "non_secure_transaction"
)

func (r SimulationCardAuthorizationNewParamsNetworkDetailsVisaElectronicCommerceIndicator) IsKnown() bool {
	switch r {
	case SimulationCardAuthorizationNewParamsNetworkDetailsVisaElectronicCommerceIndicatorMailPhoneOrder, SimulationCardAuthorizationNewParamsNetworkDetailsVisaElectronicCommerceIndicatorRecurring, SimulationCardAuthorizationNewParamsNetworkDetailsVisaElectronicCommerceIndicatorInstallment, SimulationCardAuthorizationNewParamsNetworkDetailsVisaElectronicCommerceIndicatorUnknownMailPhoneOrder, SimulationCardAuthorizationNewParamsNetworkDetailsVisaElectronicCommerceIndicatorSecureElectronicCommerce, SimulationCardAuthorizationNewParamsNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransactionAt3DSCapableMerchant, SimulationCardAuthorizationNewParamsNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransaction, SimulationCardAuthorizationNewParamsNetworkDetailsVisaElectronicCommerceIndicatorNonSecureTransaction:
		return true
	}
	return false
}

// The method used to enter the cardholder's primary account number and card
// expiration date.
type SimulationCardAuthorizationNewParamsNetworkDetailsVisaPointOfServiceEntryMode string

const (
	SimulationCardAuthorizationNewParamsNetworkDetailsVisaPointOfServiceEntryModeUnknown                    SimulationCardAuthorizationNewParamsNetworkDetailsVisaPointOfServiceEntryMode = "unknown"
	SimulationCardAuthorizationNewParamsNetworkDetailsVisaPointOfServiceEntryModeManual                     SimulationCardAuthorizationNewParamsNetworkDetailsVisaPointOfServiceEntryMode = "manual"
	SimulationCardAuthorizationNewParamsNetworkDetailsVisaPointOfServiceEntryModeMagneticStripeNoCvv        SimulationCardAuthorizationNewParamsNetworkDetailsVisaPointOfServiceEntryMode = "magnetic_stripe_no_cvv"
	SimulationCardAuthorizationNewParamsNetworkDetailsVisaPointOfServiceEntryModeOpticalCode                SimulationCardAuthorizationNewParamsNetworkDetailsVisaPointOfServiceEntryMode = "optical_code"
	SimulationCardAuthorizationNewParamsNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCard      SimulationCardAuthorizationNewParamsNetworkDetailsVisaPointOfServiceEntryMode = "integrated_circuit_card"
	SimulationCardAuthorizationNewParamsNetworkDetailsVisaPointOfServiceEntryModeContactless                SimulationCardAuthorizationNewParamsNetworkDetailsVisaPointOfServiceEntryMode = "contactless"
	SimulationCardAuthorizationNewParamsNetworkDetailsVisaPointOfServiceEntryModeCredentialOnFile           SimulationCardAuthorizationNewParamsNetworkDetailsVisaPointOfServiceEntryMode = "credential_on_file"
	SimulationCardAuthorizationNewParamsNetworkDetailsVisaPointOfServiceEntryModeMagneticStripe             SimulationCardAuthorizationNewParamsNetworkDetailsVisaPointOfServiceEntryMode = "magnetic_stripe"
	SimulationCardAuthorizationNewParamsNetworkDetailsVisaPointOfServiceEntryModeContactlessMagneticStripe  SimulationCardAuthorizationNewParamsNetworkDetailsVisaPointOfServiceEntryMode = "contactless_magnetic_stripe"
	SimulationCardAuthorizationNewParamsNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCardNoCvv SimulationCardAuthorizationNewParamsNetworkDetailsVisaPointOfServiceEntryMode = "integrated_circuit_card_no_cvv"
)

func (r SimulationCardAuthorizationNewParamsNetworkDetailsVisaPointOfServiceEntryMode) IsKnown() bool {
	switch r {
	case SimulationCardAuthorizationNewParamsNetworkDetailsVisaPointOfServiceEntryModeUnknown, SimulationCardAuthorizationNewParamsNetworkDetailsVisaPointOfServiceEntryModeManual, SimulationCardAuthorizationNewParamsNetworkDetailsVisaPointOfServiceEntryModeMagneticStripeNoCvv, SimulationCardAuthorizationNewParamsNetworkDetailsVisaPointOfServiceEntryModeOpticalCode, SimulationCardAuthorizationNewParamsNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCard, SimulationCardAuthorizationNewParamsNetworkDetailsVisaPointOfServiceEntryModeContactless, SimulationCardAuthorizationNewParamsNetworkDetailsVisaPointOfServiceEntryModeCredentialOnFile, SimulationCardAuthorizationNewParamsNetworkDetailsVisaPointOfServiceEntryModeMagneticStripe, SimulationCardAuthorizationNewParamsNetworkDetailsVisaPointOfServiceEntryModeContactlessMagneticStripe, SimulationCardAuthorizationNewParamsNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCardNoCvv:
		return true
	}
	return false
}

// The reason code for the stand-in processing.
type SimulationCardAuthorizationNewParamsNetworkDetailsVisaStandInProcessingReason string

const (
	SimulationCardAuthorizationNewParamsNetworkDetailsVisaStandInProcessingReasonIssuerError                                              SimulationCardAuthorizationNewParamsNetworkDetailsVisaStandInProcessingReason = "issuer_error"
	SimulationCardAuthorizationNewParamsNetworkDetailsVisaStandInProcessingReasonInvalidPhysicalCard                                      SimulationCardAuthorizationNewParamsNetworkDetailsVisaStandInProcessingReason = "invalid_physical_card"
	SimulationCardAuthorizationNewParamsNetworkDetailsVisaStandInProcessingReasonInvalidCryptogram                                        SimulationCardAuthorizationNewParamsNetworkDetailsVisaStandInProcessingReason = "invalid_cryptogram"
	SimulationCardAuthorizationNewParamsNetworkDetailsVisaStandInProcessingReasonInvalidCardholderAuthenticationVerificationValue         SimulationCardAuthorizationNewParamsNetworkDetailsVisaStandInProcessingReason = "invalid_cardholder_authentication_verification_value"
	SimulationCardAuthorizationNewParamsNetworkDetailsVisaStandInProcessingReasonInternalVisaError                                        SimulationCardAuthorizationNewParamsNetworkDetailsVisaStandInProcessingReason = "internal_visa_error"
	SimulationCardAuthorizationNewParamsNetworkDetailsVisaStandInProcessingReasonMerchantTransactionAdvisoryServiceAuthenticationRequired SimulationCardAuthorizationNewParamsNetworkDetailsVisaStandInProcessingReason = "merchant_transaction_advisory_service_authentication_required"
	SimulationCardAuthorizationNewParamsNetworkDetailsVisaStandInProcessingReasonPaymentFraudDisruptionAcquirerBlock                      SimulationCardAuthorizationNewParamsNetworkDetailsVisaStandInProcessingReason = "payment_fraud_disruption_acquirer_block"
	SimulationCardAuthorizationNewParamsNetworkDetailsVisaStandInProcessingReasonOther                                                    SimulationCardAuthorizationNewParamsNetworkDetailsVisaStandInProcessingReason = "other"
)

func (r SimulationCardAuthorizationNewParamsNetworkDetailsVisaStandInProcessingReason) IsKnown() bool {
	switch r {
	case SimulationCardAuthorizationNewParamsNetworkDetailsVisaStandInProcessingReasonIssuerError, SimulationCardAuthorizationNewParamsNetworkDetailsVisaStandInProcessingReasonInvalidPhysicalCard, SimulationCardAuthorizationNewParamsNetworkDetailsVisaStandInProcessingReasonInvalidCryptogram, SimulationCardAuthorizationNewParamsNetworkDetailsVisaStandInProcessingReasonInvalidCardholderAuthenticationVerificationValue, SimulationCardAuthorizationNewParamsNetworkDetailsVisaStandInProcessingReasonInternalVisaError, SimulationCardAuthorizationNewParamsNetworkDetailsVisaStandInProcessingReasonMerchantTransactionAdvisoryServiceAuthenticationRequired, SimulationCardAuthorizationNewParamsNetworkDetailsVisaStandInProcessingReasonPaymentFraudDisruptionAcquirerBlock, SimulationCardAuthorizationNewParamsNetworkDetailsVisaStandInProcessingReasonOther:
		return true
	}
	return false
}

// The capability of the terminal being used to read the card. Shows whether a
// terminal can e.g., accept chip cards or if it only supports magnetic stripe
// reads. This reflects the highest capability of the terminal — for example, a
// terminal that supports both chip and magnetic stripe will be identified as
// chip-capable.
type SimulationCardAuthorizationNewParamsNetworkDetailsVisaTerminalEntryCapability string

const (
	SimulationCardAuthorizationNewParamsNetworkDetailsVisaTerminalEntryCapabilityUnknown                     SimulationCardAuthorizationNewParamsNetworkDetailsVisaTerminalEntryCapability = "unknown"
	SimulationCardAuthorizationNewParamsNetworkDetailsVisaTerminalEntryCapabilityTerminalNotUsed             SimulationCardAuthorizationNewParamsNetworkDetailsVisaTerminalEntryCapability = "terminal_not_used"
	SimulationCardAuthorizationNewParamsNetworkDetailsVisaTerminalEntryCapabilityMagneticStripe              SimulationCardAuthorizationNewParamsNetworkDetailsVisaTerminalEntryCapability = "magnetic_stripe"
	SimulationCardAuthorizationNewParamsNetworkDetailsVisaTerminalEntryCapabilityBarcode                     SimulationCardAuthorizationNewParamsNetworkDetailsVisaTerminalEntryCapability = "barcode"
	SimulationCardAuthorizationNewParamsNetworkDetailsVisaTerminalEntryCapabilityOpticalCharacterRecognition SimulationCardAuthorizationNewParamsNetworkDetailsVisaTerminalEntryCapability = "optical_character_recognition"
	SimulationCardAuthorizationNewParamsNetworkDetailsVisaTerminalEntryCapabilityChipOrContactless           SimulationCardAuthorizationNewParamsNetworkDetailsVisaTerminalEntryCapability = "chip_or_contactless"
	SimulationCardAuthorizationNewParamsNetworkDetailsVisaTerminalEntryCapabilityContactlessOnly             SimulationCardAuthorizationNewParamsNetworkDetailsVisaTerminalEntryCapability = "contactless_only"
	SimulationCardAuthorizationNewParamsNetworkDetailsVisaTerminalEntryCapabilityNoCapability                SimulationCardAuthorizationNewParamsNetworkDetailsVisaTerminalEntryCapability = "no_capability"
)

func (r SimulationCardAuthorizationNewParamsNetworkDetailsVisaTerminalEntryCapability) IsKnown() bool {
	switch r {
	case SimulationCardAuthorizationNewParamsNetworkDetailsVisaTerminalEntryCapabilityUnknown, SimulationCardAuthorizationNewParamsNetworkDetailsVisaTerminalEntryCapabilityTerminalNotUsed, SimulationCardAuthorizationNewParamsNetworkDetailsVisaTerminalEntryCapabilityMagneticStripe, SimulationCardAuthorizationNewParamsNetworkDetailsVisaTerminalEntryCapabilityBarcode, SimulationCardAuthorizationNewParamsNetworkDetailsVisaTerminalEntryCapabilityOpticalCharacterRecognition, SimulationCardAuthorizationNewParamsNetworkDetailsVisaTerminalEntryCapabilityChipOrContactless, SimulationCardAuthorizationNewParamsNetworkDetailsVisaTerminalEntryCapabilityContactlessOnly, SimulationCardAuthorizationNewParamsNetworkDetailsVisaTerminalEntryCapabilityNoCapability:
		return true
	}
	return false
}

// Fields specific to a specific type of authorization, such as Automatic Fuel
// Dispensers, Refund Authorizations, or Cash Disbursements.
type SimulationCardAuthorizationNewParamsProcessingCategory struct {
	// The processing category describes the intent behind the authorization, such as
	// whether it was used for bill payments or an automatic fuel dispenser.
	Category param.Field[SimulationCardAuthorizationNewParamsProcessingCategoryCategory] `json:"category" api:"required"`
	// Details related to refund authorizations.
	Refund param.Field[SimulationCardAuthorizationNewParamsProcessingCategoryRefund] `json:"refund"`
}

func (r SimulationCardAuthorizationNewParamsProcessingCategory) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The processing category describes the intent behind the authorization, such as
// whether it was used for bill payments or an automatic fuel dispenser.
type SimulationCardAuthorizationNewParamsProcessingCategoryCategory string

const (
	SimulationCardAuthorizationNewParamsProcessingCategoryCategoryAccountFunding         SimulationCardAuthorizationNewParamsProcessingCategoryCategory = "account_funding"
	SimulationCardAuthorizationNewParamsProcessingCategoryCategoryAutomaticFuelDispenser SimulationCardAuthorizationNewParamsProcessingCategoryCategory = "automatic_fuel_dispenser"
	SimulationCardAuthorizationNewParamsProcessingCategoryCategoryBillPayment            SimulationCardAuthorizationNewParamsProcessingCategoryCategory = "bill_payment"
	SimulationCardAuthorizationNewParamsProcessingCategoryCategoryOriginalCredit         SimulationCardAuthorizationNewParamsProcessingCategoryCategory = "original_credit"
	SimulationCardAuthorizationNewParamsProcessingCategoryCategoryPurchase               SimulationCardAuthorizationNewParamsProcessingCategoryCategory = "purchase"
	SimulationCardAuthorizationNewParamsProcessingCategoryCategoryQuasiCash              SimulationCardAuthorizationNewParamsProcessingCategoryCategory = "quasi_cash"
	SimulationCardAuthorizationNewParamsProcessingCategoryCategoryRefund                 SimulationCardAuthorizationNewParamsProcessingCategoryCategory = "refund"
	SimulationCardAuthorizationNewParamsProcessingCategoryCategoryCashDisbursement       SimulationCardAuthorizationNewParamsProcessingCategoryCategory = "cash_disbursement"
	SimulationCardAuthorizationNewParamsProcessingCategoryCategoryCashDeposit            SimulationCardAuthorizationNewParamsProcessingCategoryCategory = "cash_deposit"
	SimulationCardAuthorizationNewParamsProcessingCategoryCategoryBalanceInquiry         SimulationCardAuthorizationNewParamsProcessingCategoryCategory = "balance_inquiry"
)

func (r SimulationCardAuthorizationNewParamsProcessingCategoryCategory) IsKnown() bool {
	switch r {
	case SimulationCardAuthorizationNewParamsProcessingCategoryCategoryAccountFunding, SimulationCardAuthorizationNewParamsProcessingCategoryCategoryAutomaticFuelDispenser, SimulationCardAuthorizationNewParamsProcessingCategoryCategoryBillPayment, SimulationCardAuthorizationNewParamsProcessingCategoryCategoryOriginalCredit, SimulationCardAuthorizationNewParamsProcessingCategoryCategoryPurchase, SimulationCardAuthorizationNewParamsProcessingCategoryCategoryQuasiCash, SimulationCardAuthorizationNewParamsProcessingCategoryCategoryRefund, SimulationCardAuthorizationNewParamsProcessingCategoryCategoryCashDisbursement, SimulationCardAuthorizationNewParamsProcessingCategoryCategoryCashDeposit, SimulationCardAuthorizationNewParamsProcessingCategoryCategoryBalanceInquiry:
		return true
	}
	return false
}

// Details related to refund authorizations.
type SimulationCardAuthorizationNewParamsProcessingCategoryRefund struct {
	// The card payment to link this refund to.
	OriginalCardPaymentID param.Field[string] `json:"original_card_payment_id"`
}

func (r SimulationCardAuthorizationNewParamsProcessingCategoryRefund) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

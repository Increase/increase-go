// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/Increase/increase-go/internal/apijson"
	"github.com/Increase/increase-go/internal/param"
	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
)

// SimulationCardTokenService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimulationCardTokenService] method instead.
type SimulationCardTokenService struct {
	Options []option.RequestOption
}

// NewSimulationCardTokenService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSimulationCardTokenService(opts ...option.RequestOption) (r *SimulationCardTokenService) {
	r = &SimulationCardTokenService{}
	r.Options = opts
	return
}

// Simulates tokenizing a card in the sandbox environment.
func (r *SimulationCardTokenService) New(ctx context.Context, body SimulationCardTokenNewParams, opts ...option.RequestOption) (res *CardToken, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "simulations/card_tokens"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type SimulationCardTokenNewParams struct {
	// The capabilities of the outbound card token.
	Capabilities param.Field[[]SimulationCardTokenNewParamsCapability] `json:"capabilities"`
	// The expiration date of the card.
	Expiration param.Field[time.Time] `json:"expiration" format:"date"`
	// The last 4 digits of the card number.
	Last4 param.Field[string] `json:"last4"`
	// The outcome to simulate for card push transfers using this token.
	Outcome param.Field[SimulationCardTokenNewParamsOutcome] `json:"outcome"`
	// The prefix of the card number, usually the first 8 digits.
	Prefix param.Field[string] `json:"prefix"`
	// The total length of the card number, including prefix and last4.
	PrimaryAccountNumberLength param.Field[int64] `json:"primary_account_number_length"`
}

func (r SimulationCardTokenNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SimulationCardTokenNewParamsCapability struct {
	// The cross-border push transfers capability.
	CrossBorderPushTransfers param.Field[SimulationCardTokenNewParamsCapabilitiesCrossBorderPushTransfers] `json:"cross_border_push_transfers" api:"required"`
	// The domestic push transfers capability.
	DomesticPushTransfers param.Field[SimulationCardTokenNewParamsCapabilitiesDomesticPushTransfers] `json:"domestic_push_transfers" api:"required"`
	// The route of the capability.
	Route param.Field[SimulationCardTokenNewParamsCapabilitiesRoute] `json:"route" api:"required"`
}

func (r SimulationCardTokenNewParamsCapability) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cross-border push transfers capability.
type SimulationCardTokenNewParamsCapabilitiesCrossBorderPushTransfers string

const (
	SimulationCardTokenNewParamsCapabilitiesCrossBorderPushTransfersSupported    SimulationCardTokenNewParamsCapabilitiesCrossBorderPushTransfers = "supported"
	SimulationCardTokenNewParamsCapabilitiesCrossBorderPushTransfersNotSupported SimulationCardTokenNewParamsCapabilitiesCrossBorderPushTransfers = "not_supported"
)

func (r SimulationCardTokenNewParamsCapabilitiesCrossBorderPushTransfers) IsKnown() bool {
	switch r {
	case SimulationCardTokenNewParamsCapabilitiesCrossBorderPushTransfersSupported, SimulationCardTokenNewParamsCapabilitiesCrossBorderPushTransfersNotSupported:
		return true
	}
	return false
}

// The domestic push transfers capability.
type SimulationCardTokenNewParamsCapabilitiesDomesticPushTransfers string

const (
	SimulationCardTokenNewParamsCapabilitiesDomesticPushTransfersSupported    SimulationCardTokenNewParamsCapabilitiesDomesticPushTransfers = "supported"
	SimulationCardTokenNewParamsCapabilitiesDomesticPushTransfersNotSupported SimulationCardTokenNewParamsCapabilitiesDomesticPushTransfers = "not_supported"
)

func (r SimulationCardTokenNewParamsCapabilitiesDomesticPushTransfers) IsKnown() bool {
	switch r {
	case SimulationCardTokenNewParamsCapabilitiesDomesticPushTransfersSupported, SimulationCardTokenNewParamsCapabilitiesDomesticPushTransfersNotSupported:
		return true
	}
	return false
}

// The route of the capability.
type SimulationCardTokenNewParamsCapabilitiesRoute string

const (
	SimulationCardTokenNewParamsCapabilitiesRouteVisa       SimulationCardTokenNewParamsCapabilitiesRoute = "visa"
	SimulationCardTokenNewParamsCapabilitiesRouteMastercard SimulationCardTokenNewParamsCapabilitiesRoute = "mastercard"
)

func (r SimulationCardTokenNewParamsCapabilitiesRoute) IsKnown() bool {
	switch r {
	case SimulationCardTokenNewParamsCapabilitiesRouteVisa, SimulationCardTokenNewParamsCapabilitiesRouteMastercard:
		return true
	}
	return false
}

// The outcome to simulate for card push transfers using this token.
type SimulationCardTokenNewParamsOutcome struct {
	// Whether card push transfers or validations will be approved or declined.
	Result param.Field[SimulationCardTokenNewParamsOutcomeResult] `json:"result" api:"required"`
	// If the result is declined, the details of the decline.
	Decline param.Field[SimulationCardTokenNewParamsOutcomeDecline] `json:"decline"`
}

func (r SimulationCardTokenNewParamsOutcome) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether card push transfers or validations will be approved or declined.
type SimulationCardTokenNewParamsOutcomeResult string

const (
	SimulationCardTokenNewParamsOutcomeResultApprove SimulationCardTokenNewParamsOutcomeResult = "approve"
	SimulationCardTokenNewParamsOutcomeResultDecline SimulationCardTokenNewParamsOutcomeResult = "decline"
)

func (r SimulationCardTokenNewParamsOutcomeResult) IsKnown() bool {
	switch r {
	case SimulationCardTokenNewParamsOutcomeResultApprove, SimulationCardTokenNewParamsOutcomeResultDecline:
		return true
	}
	return false
}

// If the result is declined, the details of the decline.
type SimulationCardTokenNewParamsOutcomeDecline struct {
	// The reason for the decline.
	Reason param.Field[SimulationCardTokenNewParamsOutcomeDeclineReason] `json:"reason"`
}

func (r SimulationCardTokenNewParamsOutcomeDecline) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The reason for the decline.
type SimulationCardTokenNewParamsOutcomeDeclineReason string

const (
	SimulationCardTokenNewParamsOutcomeDeclineReasonDoNotHonor                                              SimulationCardTokenNewParamsOutcomeDeclineReason = "do_not_honor"
	SimulationCardTokenNewParamsOutcomeDeclineReasonActivityCountLimitExceeded                              SimulationCardTokenNewParamsOutcomeDeclineReason = "activity_count_limit_exceeded"
	SimulationCardTokenNewParamsOutcomeDeclineReasonReferToCardIssuer                                       SimulationCardTokenNewParamsOutcomeDeclineReason = "refer_to_card_issuer"
	SimulationCardTokenNewParamsOutcomeDeclineReasonReferToCardIssuerSpecialCondition                       SimulationCardTokenNewParamsOutcomeDeclineReason = "refer_to_card_issuer_special_condition"
	SimulationCardTokenNewParamsOutcomeDeclineReasonInvalidMerchant                                         SimulationCardTokenNewParamsOutcomeDeclineReason = "invalid_merchant"
	SimulationCardTokenNewParamsOutcomeDeclineReasonPickUpCard                                              SimulationCardTokenNewParamsOutcomeDeclineReason = "pick_up_card"
	SimulationCardTokenNewParamsOutcomeDeclineReasonError                                                   SimulationCardTokenNewParamsOutcomeDeclineReason = "error"
	SimulationCardTokenNewParamsOutcomeDeclineReasonPickUpCardSpecial                                       SimulationCardTokenNewParamsOutcomeDeclineReason = "pick_up_card_special"
	SimulationCardTokenNewParamsOutcomeDeclineReasonInvalidTransaction                                      SimulationCardTokenNewParamsOutcomeDeclineReason = "invalid_transaction"
	SimulationCardTokenNewParamsOutcomeDeclineReasonInvalidAmount                                           SimulationCardTokenNewParamsOutcomeDeclineReason = "invalid_amount"
	SimulationCardTokenNewParamsOutcomeDeclineReasonInvalidAccountNumber                                    SimulationCardTokenNewParamsOutcomeDeclineReason = "invalid_account_number"
	SimulationCardTokenNewParamsOutcomeDeclineReasonNoSuchIssuer                                            SimulationCardTokenNewParamsOutcomeDeclineReason = "no_such_issuer"
	SimulationCardTokenNewParamsOutcomeDeclineReasonReEnterTransaction                                      SimulationCardTokenNewParamsOutcomeDeclineReason = "re_enter_transaction"
	SimulationCardTokenNewParamsOutcomeDeclineReasonNoCreditAccount                                         SimulationCardTokenNewParamsOutcomeDeclineReason = "no_credit_account"
	SimulationCardTokenNewParamsOutcomeDeclineReasonPickUpCardLost                                          SimulationCardTokenNewParamsOutcomeDeclineReason = "pick_up_card_lost"
	SimulationCardTokenNewParamsOutcomeDeclineReasonPickUpCardStolen                                        SimulationCardTokenNewParamsOutcomeDeclineReason = "pick_up_card_stolen"
	SimulationCardTokenNewParamsOutcomeDeclineReasonClosedAccount                                           SimulationCardTokenNewParamsOutcomeDeclineReason = "closed_account"
	SimulationCardTokenNewParamsOutcomeDeclineReasonInsufficientFunds                                       SimulationCardTokenNewParamsOutcomeDeclineReason = "insufficient_funds"
	SimulationCardTokenNewParamsOutcomeDeclineReasonNoCheckingAccount                                       SimulationCardTokenNewParamsOutcomeDeclineReason = "no_checking_account"
	SimulationCardTokenNewParamsOutcomeDeclineReasonNoSavingsAccount                                        SimulationCardTokenNewParamsOutcomeDeclineReason = "no_savings_account"
	SimulationCardTokenNewParamsOutcomeDeclineReasonExpiredCard                                             SimulationCardTokenNewParamsOutcomeDeclineReason = "expired_card"
	SimulationCardTokenNewParamsOutcomeDeclineReasonTransactionNotPermittedToCardholder                     SimulationCardTokenNewParamsOutcomeDeclineReason = "transaction_not_permitted_to_cardholder"
	SimulationCardTokenNewParamsOutcomeDeclineReasonTransactionNotAllowedAtTerminal                         SimulationCardTokenNewParamsOutcomeDeclineReason = "transaction_not_allowed_at_terminal"
	SimulationCardTokenNewParamsOutcomeDeclineReasonSuspectedFraud                                          SimulationCardTokenNewParamsOutcomeDeclineReason = "suspected_fraud"
	SimulationCardTokenNewParamsOutcomeDeclineReasonActivityAmountLimitExceeded                             SimulationCardTokenNewParamsOutcomeDeclineReason = "activity_amount_limit_exceeded"
	SimulationCardTokenNewParamsOutcomeDeclineReasonRestrictedCard                                          SimulationCardTokenNewParamsOutcomeDeclineReason = "restricted_card"
	SimulationCardTokenNewParamsOutcomeDeclineReasonSecurityViolation                                       SimulationCardTokenNewParamsOutcomeDeclineReason = "security_violation"
	SimulationCardTokenNewParamsOutcomeDeclineReasonTransactionDoesNotFulfillAntiMoneyLaunderingRequirement SimulationCardTokenNewParamsOutcomeDeclineReason = "transaction_does_not_fulfill_anti_money_laundering_requirement"
	SimulationCardTokenNewParamsOutcomeDeclineReasonBlockedFirstUse                                         SimulationCardTokenNewParamsOutcomeDeclineReason = "blocked_first_use"
	SimulationCardTokenNewParamsOutcomeDeclineReasonCreditIssuerUnavailable                                 SimulationCardTokenNewParamsOutcomeDeclineReason = "credit_issuer_unavailable"
	SimulationCardTokenNewParamsOutcomeDeclineReasonNegativeCardVerificationValueResults                    SimulationCardTokenNewParamsOutcomeDeclineReason = "negative_card_verification_value_results"
	SimulationCardTokenNewParamsOutcomeDeclineReasonIssuerUnavailable                                       SimulationCardTokenNewParamsOutcomeDeclineReason = "issuer_unavailable"
	SimulationCardTokenNewParamsOutcomeDeclineReasonFinancialInstitutionCannotBeFound                       SimulationCardTokenNewParamsOutcomeDeclineReason = "financial_institution_cannot_be_found"
	SimulationCardTokenNewParamsOutcomeDeclineReasonTransactionCannotBeCompleted                            SimulationCardTokenNewParamsOutcomeDeclineReason = "transaction_cannot_be_completed"
	SimulationCardTokenNewParamsOutcomeDeclineReasonDuplicateTransaction                                    SimulationCardTokenNewParamsOutcomeDeclineReason = "duplicate_transaction"
	SimulationCardTokenNewParamsOutcomeDeclineReasonSystemMalfunction                                       SimulationCardTokenNewParamsOutcomeDeclineReason = "system_malfunction"
	SimulationCardTokenNewParamsOutcomeDeclineReasonAdditionalCustomerAuthenticationRequired                SimulationCardTokenNewParamsOutcomeDeclineReason = "additional_customer_authentication_required"
	SimulationCardTokenNewParamsOutcomeDeclineReasonSurchargeAmountNotPermitted                             SimulationCardTokenNewParamsOutcomeDeclineReason = "surcharge_amount_not_permitted"
	SimulationCardTokenNewParamsOutcomeDeclineReasonDeclineForCvv2Failure                                   SimulationCardTokenNewParamsOutcomeDeclineReason = "decline_for_cvv2_failure"
	SimulationCardTokenNewParamsOutcomeDeclineReasonStopPaymentOrder                                        SimulationCardTokenNewParamsOutcomeDeclineReason = "stop_payment_order"
	SimulationCardTokenNewParamsOutcomeDeclineReasonRevocationOfAuthorizationOrder                          SimulationCardTokenNewParamsOutcomeDeclineReason = "revocation_of_authorization_order"
	SimulationCardTokenNewParamsOutcomeDeclineReasonRevocationOfAllAuthorizationsOrder                      SimulationCardTokenNewParamsOutcomeDeclineReason = "revocation_of_all_authorizations_order"
)

func (r SimulationCardTokenNewParamsOutcomeDeclineReason) IsKnown() bool {
	switch r {
	case SimulationCardTokenNewParamsOutcomeDeclineReasonDoNotHonor, SimulationCardTokenNewParamsOutcomeDeclineReasonActivityCountLimitExceeded, SimulationCardTokenNewParamsOutcomeDeclineReasonReferToCardIssuer, SimulationCardTokenNewParamsOutcomeDeclineReasonReferToCardIssuerSpecialCondition, SimulationCardTokenNewParamsOutcomeDeclineReasonInvalidMerchant, SimulationCardTokenNewParamsOutcomeDeclineReasonPickUpCard, SimulationCardTokenNewParamsOutcomeDeclineReasonError, SimulationCardTokenNewParamsOutcomeDeclineReasonPickUpCardSpecial, SimulationCardTokenNewParamsOutcomeDeclineReasonInvalidTransaction, SimulationCardTokenNewParamsOutcomeDeclineReasonInvalidAmount, SimulationCardTokenNewParamsOutcomeDeclineReasonInvalidAccountNumber, SimulationCardTokenNewParamsOutcomeDeclineReasonNoSuchIssuer, SimulationCardTokenNewParamsOutcomeDeclineReasonReEnterTransaction, SimulationCardTokenNewParamsOutcomeDeclineReasonNoCreditAccount, SimulationCardTokenNewParamsOutcomeDeclineReasonPickUpCardLost, SimulationCardTokenNewParamsOutcomeDeclineReasonPickUpCardStolen, SimulationCardTokenNewParamsOutcomeDeclineReasonClosedAccount, SimulationCardTokenNewParamsOutcomeDeclineReasonInsufficientFunds, SimulationCardTokenNewParamsOutcomeDeclineReasonNoCheckingAccount, SimulationCardTokenNewParamsOutcomeDeclineReasonNoSavingsAccount, SimulationCardTokenNewParamsOutcomeDeclineReasonExpiredCard, SimulationCardTokenNewParamsOutcomeDeclineReasonTransactionNotPermittedToCardholder, SimulationCardTokenNewParamsOutcomeDeclineReasonTransactionNotAllowedAtTerminal, SimulationCardTokenNewParamsOutcomeDeclineReasonSuspectedFraud, SimulationCardTokenNewParamsOutcomeDeclineReasonActivityAmountLimitExceeded, SimulationCardTokenNewParamsOutcomeDeclineReasonRestrictedCard, SimulationCardTokenNewParamsOutcomeDeclineReasonSecurityViolation, SimulationCardTokenNewParamsOutcomeDeclineReasonTransactionDoesNotFulfillAntiMoneyLaunderingRequirement, SimulationCardTokenNewParamsOutcomeDeclineReasonBlockedFirstUse, SimulationCardTokenNewParamsOutcomeDeclineReasonCreditIssuerUnavailable, SimulationCardTokenNewParamsOutcomeDeclineReasonNegativeCardVerificationValueResults, SimulationCardTokenNewParamsOutcomeDeclineReasonIssuerUnavailable, SimulationCardTokenNewParamsOutcomeDeclineReasonFinancialInstitutionCannotBeFound, SimulationCardTokenNewParamsOutcomeDeclineReasonTransactionCannotBeCompleted, SimulationCardTokenNewParamsOutcomeDeclineReasonDuplicateTransaction, SimulationCardTokenNewParamsOutcomeDeclineReasonSystemMalfunction, SimulationCardTokenNewParamsOutcomeDeclineReasonAdditionalCustomerAuthenticationRequired, SimulationCardTokenNewParamsOutcomeDeclineReasonSurchargeAmountNotPermitted, SimulationCardTokenNewParamsOutcomeDeclineReasonDeclineForCvv2Failure, SimulationCardTokenNewParamsOutcomeDeclineReasonStopPaymentOrder, SimulationCardTokenNewParamsOutcomeDeclineReasonRevocationOfAuthorizationOrder, SimulationCardTokenNewParamsOutcomeDeclineReasonRevocationOfAllAuthorizationsOrder:
		return true
	}
	return false
}

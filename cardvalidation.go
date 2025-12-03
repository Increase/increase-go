// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Increase/increase-go/internal/apijson"
	"github.com/Increase/increase-go/internal/apiquery"
	"github.com/Increase/increase-go/internal/param"
	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
)

// CardValidationService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCardValidationService] method instead.
type CardValidationService struct {
	Options []option.RequestOption
}

// NewCardValidationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCardValidationService(opts ...option.RequestOption) (r *CardValidationService) {
	r = &CardValidationService{}
	r.Options = opts
	return
}

// Create a Card Validation
func (r *CardValidationService) New(ctx context.Context, body CardValidationNewParams, opts ...option.RequestOption) (res *CardValidation, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "card_validations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a Card Validation
func (r *CardValidationService) Get(ctx context.Context, cardValidationID string, opts ...option.RequestOption) (res *CardValidation, err error) {
	opts = slices.Concat(r.Options, opts)
	if cardValidationID == "" {
		err = errors.New("missing required card_validation_id parameter")
		return
	}
	path := fmt.Sprintf("card_validations/%s", cardValidationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Card Validations
func (r *CardValidationService) List(ctx context.Context, query CardValidationListParams, opts ...option.RequestOption) (res *CardValidationListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "card_validations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Card Validations are used to validate a card and its cardholder before sending
// funds to or pulling funds from a card.
type CardValidation struct {
	// The Card Validation's identifier.
	ID string `json:"id,required"`
	// If the validation is accepted by the recipient bank, this will contain
	// supplemental details.
	Acceptance CardValidationAcceptance `json:"acceptance,required,nullable"`
	// The identifier of the Account from which to send the validation.
	AccountID string `json:"account_id,required"`
	// The ID of the Card Token that was used to validate the card.
	CardTokenID string `json:"card_token_id,required"`
	// The cardholder's first name.
	CardholderFirstName string `json:"cardholder_first_name,required,nullable"`
	// The cardholder's last name.
	CardholderLastName string `json:"cardholder_last_name,required,nullable"`
	// The cardholder's middle name.
	CardholderMiddleName string `json:"cardholder_middle_name,required,nullable"`
	// The postal code of the cardholder's address.
	CardholderPostalCode string `json:"cardholder_postal_code,required,nullable"`
	// The cardholder's street address.
	CardholderStreetAddress string `json:"cardholder_street_address,required,nullable"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the validation was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// What object created the validation, either via the API or the dashboard.
	CreatedBy CardValidationCreatedBy `json:"created_by,required,nullable"`
	// If the validation is rejected by the card network or the destination financial
	// institution, this will contain supplemental details.
	Decline CardValidationDecline `json:"decline,required,nullable"`
	// The idempotency key you chose for this object. This value is unique across
	// Increase and is used to ensure that a request is only processed once. Learn more
	// about [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey string `json:"idempotency_key,required,nullable"`
	// A four-digit code (MCC) identifying the type of business or service provided by
	// the merchant.
	MerchantCategoryCode string `json:"merchant_category_code,required"`
	// The city where the merchant (typically your business) is located.
	MerchantCityName string `json:"merchant_city_name,required"`
	// The merchant name that will appear in the cardholder’s statement descriptor.
	// Typically your business name.
	MerchantName string `json:"merchant_name,required"`
	// The postal code for the merchant’s (typically your business’s) location.
	MerchantPostalCode string `json:"merchant_postal_code,required"`
	// The U.S. state where the merchant (typically your business) is located.
	MerchantState string `json:"merchant_state,required"`
	// The lifecycle status of the validation.
	Status CardValidationStatus `json:"status,required"`
	// After the validation is submitted to the card network, this will contain
	// supplemental details.
	Submission CardValidationSubmission `json:"submission,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `card_validation`.
	Type        CardValidationType     `json:"type,required"`
	ExtraFields map[string]interface{} `json:"-,extras"`
	JSON        cardValidationJSON     `json:"-"`
}

// cardValidationJSON contains the JSON metadata for the struct [CardValidation]
type cardValidationJSON struct {
	ID                      apijson.Field
	Acceptance              apijson.Field
	AccountID               apijson.Field
	CardTokenID             apijson.Field
	CardholderFirstName     apijson.Field
	CardholderLastName      apijson.Field
	CardholderMiddleName    apijson.Field
	CardholderPostalCode    apijson.Field
	CardholderStreetAddress apijson.Field
	CreatedAt               apijson.Field
	CreatedBy               apijson.Field
	Decline                 apijson.Field
	IdempotencyKey          apijson.Field
	MerchantCategoryCode    apijson.Field
	MerchantCityName        apijson.Field
	MerchantName            apijson.Field
	MerchantPostalCode      apijson.Field
	MerchantState           apijson.Field
	Status                  apijson.Field
	Submission              apijson.Field
	Type                    apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *CardValidation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardValidationJSON) RawJSON() string {
	return r.raw
}

// If the validation is accepted by the recipient bank, this will contain
// supplemental details.
type CardValidationAcceptance struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the validation was accepted by the issuing bank.
	AcceptedAt time.Time `json:"accepted_at,required" format:"date-time"`
	// The authorization identification response from the issuing bank.
	AuthorizationIdentificationResponse string `json:"authorization_identification_response,required"`
	// The result of the Card Verification Value 2 match.
	CardVerificationValue2Result CardValidationAcceptanceCardVerificationValue2Result `json:"card_verification_value2_result,required,nullable"`
	// The result of the cardholder first name match.
	CardholderFirstNameResult CardValidationAcceptanceCardholderFirstNameResult `json:"cardholder_first_name_result,required,nullable"`
	// The result of the cardholder full name match.
	CardholderFullNameResult CardValidationAcceptanceCardholderFullNameResult `json:"cardholder_full_name_result,required,nullable"`
	// The result of the cardholder last name match.
	CardholderLastNameResult CardValidationAcceptanceCardholderLastNameResult `json:"cardholder_last_name_result,required,nullable"`
	// The result of the cardholder middle name match.
	CardholderMiddleNameResult CardValidationAcceptanceCardholderMiddleNameResult `json:"cardholder_middle_name_result,required,nullable"`
	// The result of the cardholder postal code match.
	CardholderPostalCodeResult CardValidationAcceptanceCardholderPostalCodeResult `json:"cardholder_postal_code_result,required,nullable"`
	// The result of the cardholder street address match.
	CardholderStreetAddressResult CardValidationAcceptanceCardholderStreetAddressResult `json:"cardholder_street_address_result,required,nullable"`
	// A unique identifier for the transaction on the card network.
	NetworkTransactionIdentifier string                       `json:"network_transaction_identifier,required,nullable"`
	JSON                         cardValidationAcceptanceJSON `json:"-"`
}

// cardValidationAcceptanceJSON contains the JSON metadata for the struct
// [CardValidationAcceptance]
type cardValidationAcceptanceJSON struct {
	AcceptedAt                          apijson.Field
	AuthorizationIdentificationResponse apijson.Field
	CardVerificationValue2Result        apijson.Field
	CardholderFirstNameResult           apijson.Field
	CardholderFullNameResult            apijson.Field
	CardholderLastNameResult            apijson.Field
	CardholderMiddleNameResult          apijson.Field
	CardholderPostalCodeResult          apijson.Field
	CardholderStreetAddressResult       apijson.Field
	NetworkTransactionIdentifier        apijson.Field
	raw                                 string
	ExtraFields                         map[string]apijson.Field
}

func (r *CardValidationAcceptance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardValidationAcceptanceJSON) RawJSON() string {
	return r.raw
}

// The result of the Card Verification Value 2 match.
type CardValidationAcceptanceCardVerificationValue2Result string

const (
	CardValidationAcceptanceCardVerificationValue2ResultMatch   CardValidationAcceptanceCardVerificationValue2Result = "match"
	CardValidationAcceptanceCardVerificationValue2ResultNoMatch CardValidationAcceptanceCardVerificationValue2Result = "no_match"
)

func (r CardValidationAcceptanceCardVerificationValue2Result) IsKnown() bool {
	switch r {
	case CardValidationAcceptanceCardVerificationValue2ResultMatch, CardValidationAcceptanceCardVerificationValue2ResultNoMatch:
		return true
	}
	return false
}

// The result of the cardholder first name match.
type CardValidationAcceptanceCardholderFirstNameResult string

const (
	CardValidationAcceptanceCardholderFirstNameResultMatch        CardValidationAcceptanceCardholderFirstNameResult = "match"
	CardValidationAcceptanceCardholderFirstNameResultNoMatch      CardValidationAcceptanceCardholderFirstNameResult = "no_match"
	CardValidationAcceptanceCardholderFirstNameResultPartialMatch CardValidationAcceptanceCardholderFirstNameResult = "partial_match"
)

func (r CardValidationAcceptanceCardholderFirstNameResult) IsKnown() bool {
	switch r {
	case CardValidationAcceptanceCardholderFirstNameResultMatch, CardValidationAcceptanceCardholderFirstNameResultNoMatch, CardValidationAcceptanceCardholderFirstNameResultPartialMatch:
		return true
	}
	return false
}

// The result of the cardholder full name match.
type CardValidationAcceptanceCardholderFullNameResult string

const (
	CardValidationAcceptanceCardholderFullNameResultMatch        CardValidationAcceptanceCardholderFullNameResult = "match"
	CardValidationAcceptanceCardholderFullNameResultNoMatch      CardValidationAcceptanceCardholderFullNameResult = "no_match"
	CardValidationAcceptanceCardholderFullNameResultPartialMatch CardValidationAcceptanceCardholderFullNameResult = "partial_match"
)

func (r CardValidationAcceptanceCardholderFullNameResult) IsKnown() bool {
	switch r {
	case CardValidationAcceptanceCardholderFullNameResultMatch, CardValidationAcceptanceCardholderFullNameResultNoMatch, CardValidationAcceptanceCardholderFullNameResultPartialMatch:
		return true
	}
	return false
}

// The result of the cardholder last name match.
type CardValidationAcceptanceCardholderLastNameResult string

const (
	CardValidationAcceptanceCardholderLastNameResultMatch        CardValidationAcceptanceCardholderLastNameResult = "match"
	CardValidationAcceptanceCardholderLastNameResultNoMatch      CardValidationAcceptanceCardholderLastNameResult = "no_match"
	CardValidationAcceptanceCardholderLastNameResultPartialMatch CardValidationAcceptanceCardholderLastNameResult = "partial_match"
)

func (r CardValidationAcceptanceCardholderLastNameResult) IsKnown() bool {
	switch r {
	case CardValidationAcceptanceCardholderLastNameResultMatch, CardValidationAcceptanceCardholderLastNameResultNoMatch, CardValidationAcceptanceCardholderLastNameResultPartialMatch:
		return true
	}
	return false
}

// The result of the cardholder middle name match.
type CardValidationAcceptanceCardholderMiddleNameResult string

const (
	CardValidationAcceptanceCardholderMiddleNameResultMatch        CardValidationAcceptanceCardholderMiddleNameResult = "match"
	CardValidationAcceptanceCardholderMiddleNameResultNoMatch      CardValidationAcceptanceCardholderMiddleNameResult = "no_match"
	CardValidationAcceptanceCardholderMiddleNameResultPartialMatch CardValidationAcceptanceCardholderMiddleNameResult = "partial_match"
)

func (r CardValidationAcceptanceCardholderMiddleNameResult) IsKnown() bool {
	switch r {
	case CardValidationAcceptanceCardholderMiddleNameResultMatch, CardValidationAcceptanceCardholderMiddleNameResultNoMatch, CardValidationAcceptanceCardholderMiddleNameResultPartialMatch:
		return true
	}
	return false
}

// The result of the cardholder postal code match.
type CardValidationAcceptanceCardholderPostalCodeResult string

const (
	CardValidationAcceptanceCardholderPostalCodeResultMatch   CardValidationAcceptanceCardholderPostalCodeResult = "match"
	CardValidationAcceptanceCardholderPostalCodeResultNoMatch CardValidationAcceptanceCardholderPostalCodeResult = "no_match"
)

func (r CardValidationAcceptanceCardholderPostalCodeResult) IsKnown() bool {
	switch r {
	case CardValidationAcceptanceCardholderPostalCodeResultMatch, CardValidationAcceptanceCardholderPostalCodeResultNoMatch:
		return true
	}
	return false
}

// The result of the cardholder street address match.
type CardValidationAcceptanceCardholderStreetAddressResult string

const (
	CardValidationAcceptanceCardholderStreetAddressResultMatch   CardValidationAcceptanceCardholderStreetAddressResult = "match"
	CardValidationAcceptanceCardholderStreetAddressResultNoMatch CardValidationAcceptanceCardholderStreetAddressResult = "no_match"
)

func (r CardValidationAcceptanceCardholderStreetAddressResult) IsKnown() bool {
	switch r {
	case CardValidationAcceptanceCardholderStreetAddressResultMatch, CardValidationAcceptanceCardholderStreetAddressResultNoMatch:
		return true
	}
	return false
}

// What object created the validation, either via the API or the dashboard.
type CardValidationCreatedBy struct {
	// If present, details about the API key that created the transfer.
	APIKey CardValidationCreatedByAPIKey `json:"api_key,required,nullable"`
	// The type of object that created this transfer.
	Category CardValidationCreatedByCategory `json:"category,required"`
	// If present, details about the OAuth Application that created the transfer.
	OAuthApplication CardValidationCreatedByOAuthApplication `json:"oauth_application,required,nullable"`
	// If present, details about the User that created the transfer.
	User CardValidationCreatedByUser `json:"user,required,nullable"`
	JSON cardValidationCreatedByJSON `json:"-"`
}

// cardValidationCreatedByJSON contains the JSON metadata for the struct
// [CardValidationCreatedBy]
type cardValidationCreatedByJSON struct {
	APIKey           apijson.Field
	Category         apijson.Field
	OAuthApplication apijson.Field
	User             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CardValidationCreatedBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardValidationCreatedByJSON) RawJSON() string {
	return r.raw
}

// If present, details about the API key that created the transfer.
type CardValidationCreatedByAPIKey struct {
	// The description set for the API key when it was created.
	Description string                            `json:"description,required,nullable"`
	JSON        cardValidationCreatedByAPIKeyJSON `json:"-"`
}

// cardValidationCreatedByAPIKeyJSON contains the JSON metadata for the struct
// [CardValidationCreatedByAPIKey]
type cardValidationCreatedByAPIKeyJSON struct {
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardValidationCreatedByAPIKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardValidationCreatedByAPIKeyJSON) RawJSON() string {
	return r.raw
}

// The type of object that created this transfer.
type CardValidationCreatedByCategory string

const (
	CardValidationCreatedByCategoryAPIKey           CardValidationCreatedByCategory = "api_key"
	CardValidationCreatedByCategoryOAuthApplication CardValidationCreatedByCategory = "oauth_application"
	CardValidationCreatedByCategoryUser             CardValidationCreatedByCategory = "user"
)

func (r CardValidationCreatedByCategory) IsKnown() bool {
	switch r {
	case CardValidationCreatedByCategoryAPIKey, CardValidationCreatedByCategoryOAuthApplication, CardValidationCreatedByCategoryUser:
		return true
	}
	return false
}

// If present, details about the OAuth Application that created the transfer.
type CardValidationCreatedByOAuthApplication struct {
	// The name of the OAuth Application.
	Name string                                      `json:"name,required"`
	JSON cardValidationCreatedByOAuthApplicationJSON `json:"-"`
}

// cardValidationCreatedByOAuthApplicationJSON contains the JSON metadata for the
// struct [CardValidationCreatedByOAuthApplication]
type cardValidationCreatedByOAuthApplicationJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardValidationCreatedByOAuthApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardValidationCreatedByOAuthApplicationJSON) RawJSON() string {
	return r.raw
}

// If present, details about the User that created the transfer.
type CardValidationCreatedByUser struct {
	// The email address of the User.
	Email string                          `json:"email,required"`
	JSON  cardValidationCreatedByUserJSON `json:"-"`
}

// cardValidationCreatedByUserJSON contains the JSON metadata for the struct
// [CardValidationCreatedByUser]
type cardValidationCreatedByUserJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardValidationCreatedByUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardValidationCreatedByUserJSON) RawJSON() string {
	return r.raw
}

// If the validation is rejected by the card network or the destination financial
// institution, this will contain supplemental details.
type CardValidationDecline struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the validation was declined.
	DeclinedAt time.Time `json:"declined_at,required" format:"date-time"`
	// A unique identifier for the transaction on the card network.
	NetworkTransactionIdentifier string `json:"network_transaction_identifier,required,nullable"`
	// The reason why the validation was declined.
	Reason CardValidationDeclineReason `json:"reason,required"`
	JSON   cardValidationDeclineJSON   `json:"-"`
}

// cardValidationDeclineJSON contains the JSON metadata for the struct
// [CardValidationDecline]
type cardValidationDeclineJSON struct {
	DeclinedAt                   apijson.Field
	NetworkTransactionIdentifier apijson.Field
	Reason                       apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *CardValidationDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardValidationDeclineJSON) RawJSON() string {
	return r.raw
}

// The reason why the validation was declined.
type CardValidationDeclineReason string

const (
	CardValidationDeclineReasonDoNotHonor                                              CardValidationDeclineReason = "do_not_honor"
	CardValidationDeclineReasonActivityCountLimitExceeded                              CardValidationDeclineReason = "activity_count_limit_exceeded"
	CardValidationDeclineReasonReferToCardIssuer                                       CardValidationDeclineReason = "refer_to_card_issuer"
	CardValidationDeclineReasonReferToCardIssuerSpecialCondition                       CardValidationDeclineReason = "refer_to_card_issuer_special_condition"
	CardValidationDeclineReasonInvalidMerchant                                         CardValidationDeclineReason = "invalid_merchant"
	CardValidationDeclineReasonPickUpCard                                              CardValidationDeclineReason = "pick_up_card"
	CardValidationDeclineReasonError                                                   CardValidationDeclineReason = "error"
	CardValidationDeclineReasonPickUpCardSpecial                                       CardValidationDeclineReason = "pick_up_card_special"
	CardValidationDeclineReasonInvalidTransaction                                      CardValidationDeclineReason = "invalid_transaction"
	CardValidationDeclineReasonInvalidAmount                                           CardValidationDeclineReason = "invalid_amount"
	CardValidationDeclineReasonInvalidAccountNumber                                    CardValidationDeclineReason = "invalid_account_number"
	CardValidationDeclineReasonNoSuchIssuer                                            CardValidationDeclineReason = "no_such_issuer"
	CardValidationDeclineReasonReEnterTransaction                                      CardValidationDeclineReason = "re_enter_transaction"
	CardValidationDeclineReasonNoCreditAccount                                         CardValidationDeclineReason = "no_credit_account"
	CardValidationDeclineReasonPickUpCardLost                                          CardValidationDeclineReason = "pick_up_card_lost"
	CardValidationDeclineReasonPickUpCardStolen                                        CardValidationDeclineReason = "pick_up_card_stolen"
	CardValidationDeclineReasonClosedAccount                                           CardValidationDeclineReason = "closed_account"
	CardValidationDeclineReasonInsufficientFunds                                       CardValidationDeclineReason = "insufficient_funds"
	CardValidationDeclineReasonNoCheckingAccount                                       CardValidationDeclineReason = "no_checking_account"
	CardValidationDeclineReasonNoSavingsAccount                                        CardValidationDeclineReason = "no_savings_account"
	CardValidationDeclineReasonExpiredCard                                             CardValidationDeclineReason = "expired_card"
	CardValidationDeclineReasonTransactionNotPermittedToCardholder                     CardValidationDeclineReason = "transaction_not_permitted_to_cardholder"
	CardValidationDeclineReasonTransactionNotAllowedAtTerminal                         CardValidationDeclineReason = "transaction_not_allowed_at_terminal"
	CardValidationDeclineReasonSuspectedFraud                                          CardValidationDeclineReason = "suspected_fraud"
	CardValidationDeclineReasonActivityAmountLimitExceeded                             CardValidationDeclineReason = "activity_amount_limit_exceeded"
	CardValidationDeclineReasonRestrictedCard                                          CardValidationDeclineReason = "restricted_card"
	CardValidationDeclineReasonSecurityViolation                                       CardValidationDeclineReason = "security_violation"
	CardValidationDeclineReasonTransactionDoesNotFulfillAntiMoneyLaunderingRequirement CardValidationDeclineReason = "transaction_does_not_fulfill_anti_money_laundering_requirement"
	CardValidationDeclineReasonBlockedFirstUse                                         CardValidationDeclineReason = "blocked_first_use"
	CardValidationDeclineReasonCreditIssuerUnavailable                                 CardValidationDeclineReason = "credit_issuer_unavailable"
	CardValidationDeclineReasonNegativeCardVerificationValueResults                    CardValidationDeclineReason = "negative_card_verification_value_results"
	CardValidationDeclineReasonIssuerUnavailable                                       CardValidationDeclineReason = "issuer_unavailable"
	CardValidationDeclineReasonFinancialInstitutionCannotBeFound                       CardValidationDeclineReason = "financial_institution_cannot_be_found"
	CardValidationDeclineReasonTransactionCannotBeCompleted                            CardValidationDeclineReason = "transaction_cannot_be_completed"
	CardValidationDeclineReasonDuplicateTransaction                                    CardValidationDeclineReason = "duplicate_transaction"
	CardValidationDeclineReasonSystemMalfunction                                       CardValidationDeclineReason = "system_malfunction"
	CardValidationDeclineReasonAdditionalCustomerAuthenticationRequired                CardValidationDeclineReason = "additional_customer_authentication_required"
	CardValidationDeclineReasonSurchargeAmountNotPermitted                             CardValidationDeclineReason = "surcharge_amount_not_permitted"
	CardValidationDeclineReasonDeclineForCvv2Failure                                   CardValidationDeclineReason = "decline_for_cvv2_failure"
	CardValidationDeclineReasonStopPaymentOrder                                        CardValidationDeclineReason = "stop_payment_order"
	CardValidationDeclineReasonRevocationOfAuthorizationOrder                          CardValidationDeclineReason = "revocation_of_authorization_order"
	CardValidationDeclineReasonRevocationOfAllAuthorizationsOrder                      CardValidationDeclineReason = "revocation_of_all_authorizations_order"
)

func (r CardValidationDeclineReason) IsKnown() bool {
	switch r {
	case CardValidationDeclineReasonDoNotHonor, CardValidationDeclineReasonActivityCountLimitExceeded, CardValidationDeclineReasonReferToCardIssuer, CardValidationDeclineReasonReferToCardIssuerSpecialCondition, CardValidationDeclineReasonInvalidMerchant, CardValidationDeclineReasonPickUpCard, CardValidationDeclineReasonError, CardValidationDeclineReasonPickUpCardSpecial, CardValidationDeclineReasonInvalidTransaction, CardValidationDeclineReasonInvalidAmount, CardValidationDeclineReasonInvalidAccountNumber, CardValidationDeclineReasonNoSuchIssuer, CardValidationDeclineReasonReEnterTransaction, CardValidationDeclineReasonNoCreditAccount, CardValidationDeclineReasonPickUpCardLost, CardValidationDeclineReasonPickUpCardStolen, CardValidationDeclineReasonClosedAccount, CardValidationDeclineReasonInsufficientFunds, CardValidationDeclineReasonNoCheckingAccount, CardValidationDeclineReasonNoSavingsAccount, CardValidationDeclineReasonExpiredCard, CardValidationDeclineReasonTransactionNotPermittedToCardholder, CardValidationDeclineReasonTransactionNotAllowedAtTerminal, CardValidationDeclineReasonSuspectedFraud, CardValidationDeclineReasonActivityAmountLimitExceeded, CardValidationDeclineReasonRestrictedCard, CardValidationDeclineReasonSecurityViolation, CardValidationDeclineReasonTransactionDoesNotFulfillAntiMoneyLaunderingRequirement, CardValidationDeclineReasonBlockedFirstUse, CardValidationDeclineReasonCreditIssuerUnavailable, CardValidationDeclineReasonNegativeCardVerificationValueResults, CardValidationDeclineReasonIssuerUnavailable, CardValidationDeclineReasonFinancialInstitutionCannotBeFound, CardValidationDeclineReasonTransactionCannotBeCompleted, CardValidationDeclineReasonDuplicateTransaction, CardValidationDeclineReasonSystemMalfunction, CardValidationDeclineReasonAdditionalCustomerAuthenticationRequired, CardValidationDeclineReasonSurchargeAmountNotPermitted, CardValidationDeclineReasonDeclineForCvv2Failure, CardValidationDeclineReasonStopPaymentOrder, CardValidationDeclineReasonRevocationOfAuthorizationOrder, CardValidationDeclineReasonRevocationOfAllAuthorizationsOrder:
		return true
	}
	return false
}

// The lifecycle status of the validation.
type CardValidationStatus string

const (
	CardValidationStatusRequiresAttention CardValidationStatus = "requires_attention"
	CardValidationStatusPendingSubmission CardValidationStatus = "pending_submission"
	CardValidationStatusSubmitted         CardValidationStatus = "submitted"
	CardValidationStatusComplete          CardValidationStatus = "complete"
	CardValidationStatusDeclined          CardValidationStatus = "declined"
)

func (r CardValidationStatus) IsKnown() bool {
	switch r {
	case CardValidationStatusRequiresAttention, CardValidationStatusPendingSubmission, CardValidationStatusSubmitted, CardValidationStatusComplete, CardValidationStatusDeclined:
		return true
	}
	return false
}

// After the validation is submitted to the card network, this will contain
// supplemental details.
type CardValidationSubmission struct {
	// A 12-digit retrieval reference number that identifies the validation. Usually a
	// combination of a timestamp and the trace number.
	RetrievalReferenceNumber string `json:"retrieval_reference_number,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the validation was submitted to the card network.
	SubmittedAt time.Time `json:"submitted_at,required" format:"date-time"`
	// A 6-digit trace number that identifies the validation within a short time
	// window.
	TraceNumber string                       `json:"trace_number,required"`
	JSON        cardValidationSubmissionJSON `json:"-"`
}

// cardValidationSubmissionJSON contains the JSON metadata for the struct
// [CardValidationSubmission]
type cardValidationSubmissionJSON struct {
	RetrievalReferenceNumber apijson.Field
	SubmittedAt              apijson.Field
	TraceNumber              apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *CardValidationSubmission) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardValidationSubmissionJSON) RawJSON() string {
	return r.raw
}

// A constant representing the object's type. For this resource it will always be
// `card_validation`.
type CardValidationType string

const (
	CardValidationTypeCardValidation CardValidationType = "card_validation"
)

func (r CardValidationType) IsKnown() bool {
	switch r {
	case CardValidationTypeCardValidation:
		return true
	}
	return false
}

// A list of Card Validation objects.
type CardValidationListResponse struct {
	// The contents of the list.
	Data []CardValidation `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor  string                         `json:"next_cursor,required,nullable"`
	ExtraFields map[string]interface{}         `json:"-,extras"`
	JSON        cardValidationListResponseJSON `json:"-"`
}

// cardValidationListResponseJSON contains the JSON metadata for the struct
// [CardValidationListResponse]
type cardValidationListResponseJSON struct {
	Data        apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardValidationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardValidationListResponseJSON) RawJSON() string {
	return r.raw
}

type CardValidationNewParams struct {
	// The identifier of the Account from which to send the validation.
	AccountID param.Field[string] `json:"account_id,required"`
	// The Increase identifier for the Card Token that represents the card number
	// you're validating.
	CardTokenID param.Field[string] `json:"card_token_id,required"`
	// A four-digit code (MCC) identifying the type of business or service provided by
	// the merchant.
	MerchantCategoryCode param.Field[string] `json:"merchant_category_code,required"`
	// The city where the merchant (typically your business) is located.
	MerchantCityName param.Field[string] `json:"merchant_city_name,required"`
	// The merchant name that will appear in the cardholder’s statement descriptor.
	// Typically your business name.
	MerchantName param.Field[string] `json:"merchant_name,required"`
	// The postal code for the merchant’s (typically your business’s) location.
	MerchantPostalCode param.Field[string] `json:"merchant_postal_code,required"`
	// The U.S. state where the merchant (typically your business) is located.
	MerchantState param.Field[string] `json:"merchant_state,required"`
	// The cardholder's first name.
	CardholderFirstName param.Field[string] `json:"cardholder_first_name"`
	// The cardholder's last name.
	CardholderLastName param.Field[string] `json:"cardholder_last_name"`
	// The cardholder's middle name.
	CardholderMiddleName param.Field[string] `json:"cardholder_middle_name"`
	// The postal code of the cardholder's address.
	CardholderPostalCode param.Field[string] `json:"cardholder_postal_code"`
	// The cardholder's street address.
	CardholderStreetAddress param.Field[string] `json:"cardholder_street_address"`
}

func (r CardValidationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CardValidationListParams struct {
	// Filter Card Validations to ones belonging to the specified Account.
	AccountID param.Field[string]                            `query:"account_id"`
	CreatedAt param.Field[CardValidationListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Filter records to the one with the specified `idempotency_key` you chose for
	// that object. This value is unique across Increase and is used to ensure that a
	// request is only processed once. Learn more about
	// [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey param.Field[string] `query:"idempotency_key"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit  param.Field[int64]                          `query:"limit"`
	Status param.Field[CardValidationListParamsStatus] `query:"status"`
}

// URLQuery serializes [CardValidationListParams]'s query parameters as
// `url.Values`.
func (r CardValidationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type CardValidationListParamsCreatedAt struct {
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

// URLQuery serializes [CardValidationListParamsCreatedAt]'s query parameters as
// `url.Values`.
func (r CardValidationListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type CardValidationListParamsStatus struct {
	// Filter Card Validations by status. For GET requests, this should be encoded as a
	// comma-delimited string, such as `?in=one,two,three`.
	In param.Field[[]CardValidationListParamsStatusIn] `query:"in"`
}

// URLQuery serializes [CardValidationListParamsStatus]'s query parameters as
// `url.Values`.
func (r CardValidationListParamsStatus) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type CardValidationListParamsStatusIn string

const (
	CardValidationListParamsStatusInRequiresAttention CardValidationListParamsStatusIn = "requires_attention"
	CardValidationListParamsStatusInPendingSubmission CardValidationListParamsStatusIn = "pending_submission"
	CardValidationListParamsStatusInSubmitted         CardValidationListParamsStatusIn = "submitted"
	CardValidationListParamsStatusInComplete          CardValidationListParamsStatusIn = "complete"
	CardValidationListParamsStatusInDeclined          CardValidationListParamsStatusIn = "declined"
)

func (r CardValidationListParamsStatusIn) IsKnown() bool {
	switch r {
	case CardValidationListParamsStatusInRequiresAttention, CardValidationListParamsStatusInPendingSubmission, CardValidationListParamsStatusInSubmitted, CardValidationListParamsStatusInComplete, CardValidationListParamsStatusInDeclined:
		return true
	}
	return false
}

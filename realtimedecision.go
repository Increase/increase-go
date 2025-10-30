// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/Increase/increase-go/internal/apijson"
	"github.com/Increase/increase-go/internal/param"
	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
)

// RealTimeDecisionService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRealTimeDecisionService] method instead.
type RealTimeDecisionService struct {
	Options []option.RequestOption
}

// NewRealTimeDecisionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRealTimeDecisionService(opts ...option.RequestOption) (r *RealTimeDecisionService) {
	r = &RealTimeDecisionService{}
	r.Options = opts
	return
}

// Retrieve a Real-Time Decision
func (r *RealTimeDecisionService) Get(ctx context.Context, realTimeDecisionID string, opts ...option.RequestOption) (res *RealTimeDecision, err error) {
	opts = slices.Concat(r.Options, opts)
	if realTimeDecisionID == "" {
		err = errors.New("missing required real_time_decision_id parameter")
		return
	}
	path := fmt.Sprintf("real_time_decisions/%s", realTimeDecisionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Action a Real-Time Decision
func (r *RealTimeDecisionService) Action(ctx context.Context, realTimeDecisionID string, body RealTimeDecisionActionParams, opts ...option.RequestOption) (res *RealTimeDecision, err error) {
	opts = slices.Concat(r.Options, opts)
	if realTimeDecisionID == "" {
		err = errors.New("missing required real_time_decision_id parameter")
		return
	}
	path := fmt.Sprintf("real_time_decisions/%s/action", realTimeDecisionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Real Time Decisions are created when your application needs to take action in
// real-time to some event such as a card authorization. For more information, see
// our
// [Real-Time Decisions guide](https://increase.com/documentation/real-time-decisions).
type RealTimeDecision struct {
	// The Real-Time Decision identifier.
	ID string `json:"id,required"`
	// Fields related to a 3DS authentication attempt.
	CardAuthentication RealTimeDecisionCardAuthentication `json:"card_authentication,required,nullable"`
	// Fields related to a 3DS authentication attempt.
	CardAuthenticationChallenge RealTimeDecisionCardAuthenticationChallenge `json:"card_authentication_challenge,required,nullable"`
	// Fields related to a card authorization.
	CardAuthorization RealTimeDecisionCardAuthorization `json:"card_authorization,required,nullable"`
	// The category of the Real-Time Decision.
	Category RealTimeDecisionCategory `json:"category,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Real-Time Decision was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Fields related to a digital wallet authentication attempt.
	DigitalWalletAuthentication RealTimeDecisionDigitalWalletAuthentication `json:"digital_wallet_authentication,required,nullable"`
	// Fields related to a digital wallet token provisioning attempt.
	DigitalWalletToken RealTimeDecisionDigitalWalletToken `json:"digital_wallet_token,required,nullable"`
	// The status of the Real-Time Decision.
	Status RealTimeDecisionStatus `json:"status,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// your application can no longer respond to the Real-Time Decision.
	TimeoutAt time.Time `json:"timeout_at,required" format:"date-time"`
	// A constant representing the object's type. For this resource it will always be
	// `real_time_decision`.
	Type RealTimeDecisionType `json:"type,required"`
	JSON realTimeDecisionJSON `json:"-"`
}

// realTimeDecisionJSON contains the JSON metadata for the struct
// [RealTimeDecision]
type realTimeDecisionJSON struct {
	ID                          apijson.Field
	CardAuthentication          apijson.Field
	CardAuthenticationChallenge apijson.Field
	CardAuthorization           apijson.Field
	Category                    apijson.Field
	CreatedAt                   apijson.Field
	DigitalWalletAuthentication apijson.Field
	DigitalWalletToken          apijson.Field
	Status                      apijson.Field
	TimeoutAt                   apijson.Field
	Type                        apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *RealTimeDecision) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimeDecisionJSON) RawJSON() string {
	return r.raw
}

// Fields related to a 3DS authentication attempt.
type RealTimeDecisionCardAuthentication struct {
	// The identifier of the Account the card belongs to.
	AccountID string `json:"account_id,required"`
	// The identifier of the Card that is being tokenized.
	CardID string `json:"card_id,required"`
	// Whether or not the authentication attempt was approved.
	Decision RealTimeDecisionCardAuthenticationDecision `json:"decision,required,nullable"`
	// The identifier of the Card Payment this authentication attempt will belong to.
	// Available in the API once the card authentication has completed.
	UpcomingCardPaymentID string                                 `json:"upcoming_card_payment_id,required"`
	JSON                  realTimeDecisionCardAuthenticationJSON `json:"-"`
}

// realTimeDecisionCardAuthenticationJSON contains the JSON metadata for the struct
// [RealTimeDecisionCardAuthentication]
type realTimeDecisionCardAuthenticationJSON struct {
	AccountID             apijson.Field
	CardID                apijson.Field
	Decision              apijson.Field
	UpcomingCardPaymentID apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *RealTimeDecisionCardAuthentication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimeDecisionCardAuthenticationJSON) RawJSON() string {
	return r.raw
}

// Whether or not the authentication attempt was approved.
type RealTimeDecisionCardAuthenticationDecision string

const (
	RealTimeDecisionCardAuthenticationDecisionApprove   RealTimeDecisionCardAuthenticationDecision = "approve"
	RealTimeDecisionCardAuthenticationDecisionChallenge RealTimeDecisionCardAuthenticationDecision = "challenge"
	RealTimeDecisionCardAuthenticationDecisionDeny      RealTimeDecisionCardAuthenticationDecision = "deny"
)

func (r RealTimeDecisionCardAuthenticationDecision) IsKnown() bool {
	switch r {
	case RealTimeDecisionCardAuthenticationDecisionApprove, RealTimeDecisionCardAuthenticationDecisionChallenge, RealTimeDecisionCardAuthenticationDecisionDeny:
		return true
	}
	return false
}

// Fields related to a 3DS authentication attempt.
type RealTimeDecisionCardAuthenticationChallenge struct {
	// The identifier of the Account the card belongs to.
	AccountID string `json:"account_id,required"`
	// The identifier of the Card that is being tokenized.
	CardID string `json:"card_id,required"`
	// The identifier of the Card Payment this authentication challenge attempt belongs
	// to.
	CardPaymentID string `json:"card_payment_id,required"`
	// The one-time code delivered to the cardholder.
	OneTimeCode string `json:"one_time_code,required"`
	// Whether or not the challenge was delivered to the cardholder.
	Result RealTimeDecisionCardAuthenticationChallengeResult `json:"result,required,nullable"`
	JSON   realTimeDecisionCardAuthenticationChallengeJSON   `json:"-"`
}

// realTimeDecisionCardAuthenticationChallengeJSON contains the JSON metadata for
// the struct [RealTimeDecisionCardAuthenticationChallenge]
type realTimeDecisionCardAuthenticationChallengeJSON struct {
	AccountID     apijson.Field
	CardID        apijson.Field
	CardPaymentID apijson.Field
	OneTimeCode   apijson.Field
	Result        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RealTimeDecisionCardAuthenticationChallenge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimeDecisionCardAuthenticationChallengeJSON) RawJSON() string {
	return r.raw
}

// Whether or not the challenge was delivered to the cardholder.
type RealTimeDecisionCardAuthenticationChallengeResult string

const (
	RealTimeDecisionCardAuthenticationChallengeResultSuccess RealTimeDecisionCardAuthenticationChallengeResult = "success"
	RealTimeDecisionCardAuthenticationChallengeResultFailure RealTimeDecisionCardAuthenticationChallengeResult = "failure"
)

func (r RealTimeDecisionCardAuthenticationChallengeResult) IsKnown() bool {
	switch r {
	case RealTimeDecisionCardAuthenticationChallengeResultSuccess, RealTimeDecisionCardAuthenticationChallengeResultFailure:
		return true
	}
	return false
}

// Fields related to a card authorization.
type RealTimeDecisionCardAuthorization struct {
	// The identifier of the Account the authorization will debit.
	AccountID string `json:"account_id,required"`
	// Additional amounts associated with the card authorization, such as ATM
	// surcharges fees. These are usually a subset of the `amount` field and are used
	// to provide more detailed information about the transaction.
	AdditionalAmounts RealTimeDecisionCardAuthorizationAdditionalAmounts `json:"additional_amounts,required"`
	// The identifier of the Card that is being authorized.
	CardID string `json:"card_id,required"`
	// Whether or not the authorization was approved.
	Decision RealTimeDecisionCardAuthorizationDecision `json:"decision,required,nullable"`
	// Present if and only if `decision` is `decline`. Contains information related to
	// the reason the authorization was declined.
	Decline RealTimeDecisionCardAuthorizationDecline `json:"decline,required,nullable"`
	// If the authorization was made via a Digital Wallet Token (such as an Apple Pay
	// purchase), the identifier of the token that was used.
	DigitalWalletTokenID string `json:"digital_wallet_token_id,required,nullable"`
	// The direction describes the direction the funds will move, either from the
	// cardholder to the merchant or from the merchant to the cardholder.
	Direction RealTimeDecisionCardAuthorizationDirection `json:"direction,required"`
	// The merchant identifier (commonly abbreviated as MID) of the merchant the card
	// is transacting with.
	MerchantAcceptorID string `json:"merchant_acceptor_id,required"`
	// The Merchant Category Code (commonly abbreviated as MCC) of the merchant the
	// card is transacting with.
	MerchantCategoryCode string `json:"merchant_category_code,required"`
	// The city the merchant resides in.
	MerchantCity string `json:"merchant_city,required,nullable"`
	// The country the merchant resides in.
	MerchantCountry string `json:"merchant_country,required"`
	// The merchant descriptor of the merchant the card is transacting with.
	MerchantDescriptor string `json:"merchant_descriptor,required"`
	// The merchant's postal code. For US merchants this is either a 5-digit or 9-digit
	// ZIP code, where the first 5 and last 4 are separated by a dash.
	MerchantPostalCode string `json:"merchant_postal_code,required,nullable"`
	// The state the merchant resides in.
	MerchantState string `json:"merchant_state,required,nullable"`
	// Fields specific to the `network`.
	NetworkDetails RealTimeDecisionCardAuthorizationNetworkDetails `json:"network_details,required"`
	// Network-specific identifiers for a specific request or transaction.
	NetworkIdentifiers RealTimeDecisionCardAuthorizationNetworkIdentifiers `json:"network_identifiers,required"`
	// The risk score generated by the card network. For Visa this is the Visa Advanced
	// Authorization risk score, from 0 to 99, where 99 is the riskiest. For Pulse the
	// score is from 0 to 999, where 999 is the riskiest.
	NetworkRiskScore int64 `json:"network_risk_score,required,nullable"`
	// If the authorization was made in-person with a physical card, the Physical Card
	// that was used.
	PhysicalCardID string `json:"physical_card_id,required,nullable"`
	// The amount of the attempted authorization in the currency the card user sees at
	// the time of purchase, in the minor unit of that currency. For dollars, for
	// example, this is cents.
	PresentmentAmount int64 `json:"presentment_amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the currency the
	// user sees at the time of purchase.
	PresentmentCurrency string `json:"presentment_currency,required"`
	// The processing category describes the intent behind the authorization, such as
	// whether it was used for bill payments or an automatic fuel dispenser.
	ProcessingCategory RealTimeDecisionCardAuthorizationProcessingCategory `json:"processing_category,required"`
	// Fields specific to the type of request, such as an incremental authorization.
	RequestDetails RealTimeDecisionCardAuthorizationRequestDetails `json:"request_details,required"`
	// The amount of the attempted authorization in the currency it will be settled in.
	// This currency is the same as that of the Account the card belongs to.
	SettlementAmount int64 `json:"settlement_amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the currency the
	// transaction will be settled in.
	SettlementCurrency string `json:"settlement_currency,required"`
	// The terminal identifier (commonly abbreviated as TID) of the terminal the card
	// is transacting with.
	TerminalID string `json:"terminal_id,required,nullable"`
	// The identifier of the Card Payment this authorization will belong to. Available
	// in the API once the card authorization has completed.
	UpcomingCardPaymentID string `json:"upcoming_card_payment_id,required"`
	// Fields related to verification of cardholder-provided values.
	Verification RealTimeDecisionCardAuthorizationVerification `json:"verification,required"`
	ExtraFields  map[string]interface{}                        `json:"-,extras"`
	JSON         realTimeDecisionCardAuthorizationJSON         `json:"-"`
}

// realTimeDecisionCardAuthorizationJSON contains the JSON metadata for the struct
// [RealTimeDecisionCardAuthorization]
type realTimeDecisionCardAuthorizationJSON struct {
	AccountID             apijson.Field
	AdditionalAmounts     apijson.Field
	CardID                apijson.Field
	Decision              apijson.Field
	Decline               apijson.Field
	DigitalWalletTokenID  apijson.Field
	Direction             apijson.Field
	MerchantAcceptorID    apijson.Field
	MerchantCategoryCode  apijson.Field
	MerchantCity          apijson.Field
	MerchantCountry       apijson.Field
	MerchantDescriptor    apijson.Field
	MerchantPostalCode    apijson.Field
	MerchantState         apijson.Field
	NetworkDetails        apijson.Field
	NetworkIdentifiers    apijson.Field
	NetworkRiskScore      apijson.Field
	PhysicalCardID        apijson.Field
	PresentmentAmount     apijson.Field
	PresentmentCurrency   apijson.Field
	ProcessingCategory    apijson.Field
	RequestDetails        apijson.Field
	SettlementAmount      apijson.Field
	SettlementCurrency    apijson.Field
	TerminalID            apijson.Field
	UpcomingCardPaymentID apijson.Field
	Verification          apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *RealTimeDecisionCardAuthorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimeDecisionCardAuthorizationJSON) RawJSON() string {
	return r.raw
}

// Additional amounts associated with the card authorization, such as ATM
// surcharges fees. These are usually a subset of the `amount` field and are used
// to provide more detailed information about the transaction.
type RealTimeDecisionCardAuthorizationAdditionalAmounts struct {
	// The part of this transaction amount that was for clinic-related services.
	Clinic RealTimeDecisionCardAuthorizationAdditionalAmountsClinic `json:"clinic,required,nullable"`
	// The part of this transaction amount that was for dental-related services.
	Dental RealTimeDecisionCardAuthorizationAdditionalAmountsDental `json:"dental,required,nullable"`
	// The original pre-authorized amount.
	Original RealTimeDecisionCardAuthorizationAdditionalAmountsOriginal `json:"original,required,nullable"`
	// The part of this transaction amount that was for healthcare prescriptions.
	Prescription RealTimeDecisionCardAuthorizationAdditionalAmountsPrescription `json:"prescription,required,nullable"`
	// The surcharge amount charged for this transaction by the merchant.
	Surcharge RealTimeDecisionCardAuthorizationAdditionalAmountsSurcharge `json:"surcharge,required,nullable"`
	// The total amount of a series of incremental authorizations, optionally provided.
	TotalCumulative RealTimeDecisionCardAuthorizationAdditionalAmountsTotalCumulative `json:"total_cumulative,required,nullable"`
	// The total amount of healthcare-related additional amounts.
	TotalHealthcare RealTimeDecisionCardAuthorizationAdditionalAmountsTotalHealthcare `json:"total_healthcare,required,nullable"`
	// The part of this transaction amount that was for transit-related services.
	Transit RealTimeDecisionCardAuthorizationAdditionalAmountsTransit `json:"transit,required,nullable"`
	// An unknown additional amount.
	Unknown RealTimeDecisionCardAuthorizationAdditionalAmountsUnknown `json:"unknown,required,nullable"`
	// The part of this transaction amount that was for vision-related services.
	Vision RealTimeDecisionCardAuthorizationAdditionalAmountsVision `json:"vision,required,nullable"`
	JSON   realTimeDecisionCardAuthorizationAdditionalAmountsJSON   `json:"-"`
}

// realTimeDecisionCardAuthorizationAdditionalAmountsJSON contains the JSON
// metadata for the struct [RealTimeDecisionCardAuthorizationAdditionalAmounts]
type realTimeDecisionCardAuthorizationAdditionalAmountsJSON struct {
	Clinic          apijson.Field
	Dental          apijson.Field
	Original        apijson.Field
	Prescription    apijson.Field
	Surcharge       apijson.Field
	TotalCumulative apijson.Field
	TotalHealthcare apijson.Field
	Transit         apijson.Field
	Unknown         apijson.Field
	Vision          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RealTimeDecisionCardAuthorizationAdditionalAmounts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimeDecisionCardAuthorizationAdditionalAmountsJSON) RawJSON() string {
	return r.raw
}

// The part of this transaction amount that was for clinic-related services.
type RealTimeDecisionCardAuthorizationAdditionalAmountsClinic struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                       `json:"currency,required"`
	JSON     realTimeDecisionCardAuthorizationAdditionalAmountsClinicJSON `json:"-"`
}

// realTimeDecisionCardAuthorizationAdditionalAmountsClinicJSON contains the JSON
// metadata for the struct
// [RealTimeDecisionCardAuthorizationAdditionalAmountsClinic]
type realTimeDecisionCardAuthorizationAdditionalAmountsClinicJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RealTimeDecisionCardAuthorizationAdditionalAmountsClinic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimeDecisionCardAuthorizationAdditionalAmountsClinicJSON) RawJSON() string {
	return r.raw
}

// The part of this transaction amount that was for dental-related services.
type RealTimeDecisionCardAuthorizationAdditionalAmountsDental struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                       `json:"currency,required"`
	JSON     realTimeDecisionCardAuthorizationAdditionalAmountsDentalJSON `json:"-"`
}

// realTimeDecisionCardAuthorizationAdditionalAmountsDentalJSON contains the JSON
// metadata for the struct
// [RealTimeDecisionCardAuthorizationAdditionalAmountsDental]
type realTimeDecisionCardAuthorizationAdditionalAmountsDentalJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RealTimeDecisionCardAuthorizationAdditionalAmountsDental) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimeDecisionCardAuthorizationAdditionalAmountsDentalJSON) RawJSON() string {
	return r.raw
}

// The original pre-authorized amount.
type RealTimeDecisionCardAuthorizationAdditionalAmountsOriginal struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                         `json:"currency,required"`
	JSON     realTimeDecisionCardAuthorizationAdditionalAmountsOriginalJSON `json:"-"`
}

// realTimeDecisionCardAuthorizationAdditionalAmountsOriginalJSON contains the JSON
// metadata for the struct
// [RealTimeDecisionCardAuthorizationAdditionalAmountsOriginal]
type realTimeDecisionCardAuthorizationAdditionalAmountsOriginalJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RealTimeDecisionCardAuthorizationAdditionalAmountsOriginal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimeDecisionCardAuthorizationAdditionalAmountsOriginalJSON) RawJSON() string {
	return r.raw
}

// The part of this transaction amount that was for healthcare prescriptions.
type RealTimeDecisionCardAuthorizationAdditionalAmountsPrescription struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                             `json:"currency,required"`
	JSON     realTimeDecisionCardAuthorizationAdditionalAmountsPrescriptionJSON `json:"-"`
}

// realTimeDecisionCardAuthorizationAdditionalAmountsPrescriptionJSON contains the
// JSON metadata for the struct
// [RealTimeDecisionCardAuthorizationAdditionalAmountsPrescription]
type realTimeDecisionCardAuthorizationAdditionalAmountsPrescriptionJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RealTimeDecisionCardAuthorizationAdditionalAmountsPrescription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimeDecisionCardAuthorizationAdditionalAmountsPrescriptionJSON) RawJSON() string {
	return r.raw
}

// The surcharge amount charged for this transaction by the merchant.
type RealTimeDecisionCardAuthorizationAdditionalAmountsSurcharge struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                          `json:"currency,required"`
	JSON     realTimeDecisionCardAuthorizationAdditionalAmountsSurchargeJSON `json:"-"`
}

// realTimeDecisionCardAuthorizationAdditionalAmountsSurchargeJSON contains the
// JSON metadata for the struct
// [RealTimeDecisionCardAuthorizationAdditionalAmountsSurcharge]
type realTimeDecisionCardAuthorizationAdditionalAmountsSurchargeJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RealTimeDecisionCardAuthorizationAdditionalAmountsSurcharge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimeDecisionCardAuthorizationAdditionalAmountsSurchargeJSON) RawJSON() string {
	return r.raw
}

// The total amount of a series of incremental authorizations, optionally provided.
type RealTimeDecisionCardAuthorizationAdditionalAmountsTotalCumulative struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                                `json:"currency,required"`
	JSON     realTimeDecisionCardAuthorizationAdditionalAmountsTotalCumulativeJSON `json:"-"`
}

// realTimeDecisionCardAuthorizationAdditionalAmountsTotalCumulativeJSON contains
// the JSON metadata for the struct
// [RealTimeDecisionCardAuthorizationAdditionalAmountsTotalCumulative]
type realTimeDecisionCardAuthorizationAdditionalAmountsTotalCumulativeJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RealTimeDecisionCardAuthorizationAdditionalAmountsTotalCumulative) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimeDecisionCardAuthorizationAdditionalAmountsTotalCumulativeJSON) RawJSON() string {
	return r.raw
}

// The total amount of healthcare-related additional amounts.
type RealTimeDecisionCardAuthorizationAdditionalAmountsTotalHealthcare struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                                `json:"currency,required"`
	JSON     realTimeDecisionCardAuthorizationAdditionalAmountsTotalHealthcareJSON `json:"-"`
}

// realTimeDecisionCardAuthorizationAdditionalAmountsTotalHealthcareJSON contains
// the JSON metadata for the struct
// [RealTimeDecisionCardAuthorizationAdditionalAmountsTotalHealthcare]
type realTimeDecisionCardAuthorizationAdditionalAmountsTotalHealthcareJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RealTimeDecisionCardAuthorizationAdditionalAmountsTotalHealthcare) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimeDecisionCardAuthorizationAdditionalAmountsTotalHealthcareJSON) RawJSON() string {
	return r.raw
}

// The part of this transaction amount that was for transit-related services.
type RealTimeDecisionCardAuthorizationAdditionalAmountsTransit struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                        `json:"currency,required"`
	JSON     realTimeDecisionCardAuthorizationAdditionalAmountsTransitJSON `json:"-"`
}

// realTimeDecisionCardAuthorizationAdditionalAmountsTransitJSON contains the JSON
// metadata for the struct
// [RealTimeDecisionCardAuthorizationAdditionalAmountsTransit]
type realTimeDecisionCardAuthorizationAdditionalAmountsTransitJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RealTimeDecisionCardAuthorizationAdditionalAmountsTransit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimeDecisionCardAuthorizationAdditionalAmountsTransitJSON) RawJSON() string {
	return r.raw
}

// An unknown additional amount.
type RealTimeDecisionCardAuthorizationAdditionalAmountsUnknown struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                        `json:"currency,required"`
	JSON     realTimeDecisionCardAuthorizationAdditionalAmountsUnknownJSON `json:"-"`
}

// realTimeDecisionCardAuthorizationAdditionalAmountsUnknownJSON contains the JSON
// metadata for the struct
// [RealTimeDecisionCardAuthorizationAdditionalAmountsUnknown]
type realTimeDecisionCardAuthorizationAdditionalAmountsUnknownJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RealTimeDecisionCardAuthorizationAdditionalAmountsUnknown) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimeDecisionCardAuthorizationAdditionalAmountsUnknownJSON) RawJSON() string {
	return r.raw
}

// The part of this transaction amount that was for vision-related services.
type RealTimeDecisionCardAuthorizationAdditionalAmountsVision struct {
	// The amount in minor units of the `currency` field. The amount is positive if it
	// is added to the amount (such as an ATM surcharge fee) and negative if it is
	// subtracted from the amount (such as a discount).
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the additional
	// amount's currency.
	Currency string                                                       `json:"currency,required"`
	JSON     realTimeDecisionCardAuthorizationAdditionalAmountsVisionJSON `json:"-"`
}

// realTimeDecisionCardAuthorizationAdditionalAmountsVisionJSON contains the JSON
// metadata for the struct
// [RealTimeDecisionCardAuthorizationAdditionalAmountsVision]
type realTimeDecisionCardAuthorizationAdditionalAmountsVisionJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RealTimeDecisionCardAuthorizationAdditionalAmountsVision) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimeDecisionCardAuthorizationAdditionalAmountsVisionJSON) RawJSON() string {
	return r.raw
}

// Whether or not the authorization was approved.
type RealTimeDecisionCardAuthorizationDecision string

const (
	RealTimeDecisionCardAuthorizationDecisionApprove RealTimeDecisionCardAuthorizationDecision = "approve"
	RealTimeDecisionCardAuthorizationDecisionDecline RealTimeDecisionCardAuthorizationDecision = "decline"
)

func (r RealTimeDecisionCardAuthorizationDecision) IsKnown() bool {
	switch r {
	case RealTimeDecisionCardAuthorizationDecisionApprove, RealTimeDecisionCardAuthorizationDecisionDecline:
		return true
	}
	return false
}

// Present if and only if `decision` is `decline`. Contains information related to
// the reason the authorization was declined.
type RealTimeDecisionCardAuthorizationDecline struct {
	// The reason the authorization was declined.
	Reason string                                       `json:"reason,required"`
	JSON   realTimeDecisionCardAuthorizationDeclineJSON `json:"-"`
}

// realTimeDecisionCardAuthorizationDeclineJSON contains the JSON metadata for the
// struct [RealTimeDecisionCardAuthorizationDecline]
type realTimeDecisionCardAuthorizationDeclineJSON struct {
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RealTimeDecisionCardAuthorizationDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimeDecisionCardAuthorizationDeclineJSON) RawJSON() string {
	return r.raw
}

// The direction describes the direction the funds will move, either from the
// cardholder to the merchant or from the merchant to the cardholder.
type RealTimeDecisionCardAuthorizationDirection string

const (
	RealTimeDecisionCardAuthorizationDirectionSettlement RealTimeDecisionCardAuthorizationDirection = "settlement"
	RealTimeDecisionCardAuthorizationDirectionRefund     RealTimeDecisionCardAuthorizationDirection = "refund"
)

func (r RealTimeDecisionCardAuthorizationDirection) IsKnown() bool {
	switch r {
	case RealTimeDecisionCardAuthorizationDirectionSettlement, RealTimeDecisionCardAuthorizationDirectionRefund:
		return true
	}
	return false
}

// Fields specific to the `network`.
type RealTimeDecisionCardAuthorizationNetworkDetails struct {
	// The payment network used to process this card authorization.
	Category RealTimeDecisionCardAuthorizationNetworkDetailsCategory `json:"category,required"`
	// Fields specific to the `pulse` network.
	Pulse interface{} `json:"pulse,required,nullable"`
	// Fields specific to the `visa` network.
	Visa RealTimeDecisionCardAuthorizationNetworkDetailsVisa `json:"visa,required,nullable"`
	JSON realTimeDecisionCardAuthorizationNetworkDetailsJSON `json:"-"`
}

// realTimeDecisionCardAuthorizationNetworkDetailsJSON contains the JSON metadata
// for the struct [RealTimeDecisionCardAuthorizationNetworkDetails]
type realTimeDecisionCardAuthorizationNetworkDetailsJSON struct {
	Category    apijson.Field
	Pulse       apijson.Field
	Visa        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RealTimeDecisionCardAuthorizationNetworkDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimeDecisionCardAuthorizationNetworkDetailsJSON) RawJSON() string {
	return r.raw
}

// The payment network used to process this card authorization.
type RealTimeDecisionCardAuthorizationNetworkDetailsCategory string

const (
	RealTimeDecisionCardAuthorizationNetworkDetailsCategoryVisa  RealTimeDecisionCardAuthorizationNetworkDetailsCategory = "visa"
	RealTimeDecisionCardAuthorizationNetworkDetailsCategoryPulse RealTimeDecisionCardAuthorizationNetworkDetailsCategory = "pulse"
)

func (r RealTimeDecisionCardAuthorizationNetworkDetailsCategory) IsKnown() bool {
	switch r {
	case RealTimeDecisionCardAuthorizationNetworkDetailsCategoryVisa, RealTimeDecisionCardAuthorizationNetworkDetailsCategoryPulse:
		return true
	}
	return false
}

// Fields specific to the `visa` network.
type RealTimeDecisionCardAuthorizationNetworkDetailsVisa struct {
	// For electronic commerce transactions, this identifies the level of security used
	// in obtaining the customer's payment credential. For mail or telephone order
	// transactions, identifies the type of mail or telephone order.
	ElectronicCommerceIndicator RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator `json:"electronic_commerce_indicator,required,nullable"`
	// The method used to enter the cardholder's primary account number and card
	// expiration date.
	PointOfServiceEntryMode RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode `json:"point_of_service_entry_mode,required,nullable"`
	// Only present when `actioner: network`. Describes why a card authorization was
	// approved or declined by Visa through stand-in processing.
	StandInProcessingReason RealTimeDecisionCardAuthorizationNetworkDetailsVisaStandInProcessingReason `json:"stand_in_processing_reason,required,nullable"`
	JSON                    realTimeDecisionCardAuthorizationNetworkDetailsVisaJSON                    `json:"-"`
}

// realTimeDecisionCardAuthorizationNetworkDetailsVisaJSON contains the JSON
// metadata for the struct [RealTimeDecisionCardAuthorizationNetworkDetailsVisa]
type realTimeDecisionCardAuthorizationNetworkDetailsVisaJSON struct {
	ElectronicCommerceIndicator apijson.Field
	PointOfServiceEntryMode     apijson.Field
	StandInProcessingReason     apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *RealTimeDecisionCardAuthorizationNetworkDetailsVisa) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimeDecisionCardAuthorizationNetworkDetailsVisaJSON) RawJSON() string {
	return r.raw
}

// For electronic commerce transactions, this identifies the level of security used
// in obtaining the customer's payment credential. For mail or telephone order
// transactions, identifies the type of mail or telephone order.
type RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator string

const (
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorMailPhoneOrder                                          RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "mail_phone_order"
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorRecurring                                               RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "recurring"
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorInstallment                                             RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "installment"
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorUnknownMailPhoneOrder                                   RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "unknown_mail_phone_order"
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorSecureElectronicCommerce                                RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "secure_electronic_commerce"
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransactionAt3DSCapableMerchant RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction_at_3ds_capable_merchant"
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransaction                     RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction"
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonSecureTransaction                                    RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "non_secure_transaction"
)

func (r RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator) IsKnown() bool {
	switch r {
	case RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorMailPhoneOrder, RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorRecurring, RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorInstallment, RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorUnknownMailPhoneOrder, RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorSecureElectronicCommerce, RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransactionAt3DSCapableMerchant, RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransaction, RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonSecureTransaction:
		return true
	}
	return false
}

// The method used to enter the cardholder's primary account number and card
// expiration date.
type RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode string

const (
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeUnknown                    RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "unknown"
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeManual                     RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "manual"
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeMagneticStripeNoCvv        RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "magnetic_stripe_no_cvv"
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeOpticalCode                RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "optical_code"
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCard      RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "integrated_circuit_card"
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeContactless                RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "contactless"
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeCredentialOnFile           RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "credential_on_file"
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeMagneticStripe             RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "magnetic_stripe"
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeContactlessMagneticStripe  RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "contactless_magnetic_stripe"
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCardNoCvv RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode = "integrated_circuit_card_no_cvv"
)

func (r RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryMode) IsKnown() bool {
	switch r {
	case RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeUnknown, RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeManual, RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeMagneticStripeNoCvv, RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeOpticalCode, RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCard, RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeContactless, RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeCredentialOnFile, RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeMagneticStripe, RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeContactlessMagneticStripe, RealTimeDecisionCardAuthorizationNetworkDetailsVisaPointOfServiceEntryModeIntegratedCircuitCardNoCvv:
		return true
	}
	return false
}

// Only present when `actioner: network`. Describes why a card authorization was
// approved or declined by Visa through stand-in processing.
type RealTimeDecisionCardAuthorizationNetworkDetailsVisaStandInProcessingReason string

const (
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaStandInProcessingReasonIssuerError                                              RealTimeDecisionCardAuthorizationNetworkDetailsVisaStandInProcessingReason = "issuer_error"
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaStandInProcessingReasonInvalidPhysicalCard                                      RealTimeDecisionCardAuthorizationNetworkDetailsVisaStandInProcessingReason = "invalid_physical_card"
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaStandInProcessingReasonInvalidCardholderAuthenticationVerificationValue         RealTimeDecisionCardAuthorizationNetworkDetailsVisaStandInProcessingReason = "invalid_cardholder_authentication_verification_value"
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaStandInProcessingReasonInternalVisaError                                        RealTimeDecisionCardAuthorizationNetworkDetailsVisaStandInProcessingReason = "internal_visa_error"
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaStandInProcessingReasonMerchantTransactionAdvisoryServiceAuthenticationRequired RealTimeDecisionCardAuthorizationNetworkDetailsVisaStandInProcessingReason = "merchant_transaction_advisory_service_authentication_required"
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaStandInProcessingReasonPaymentFraudDisruptionAcquirerBlock                      RealTimeDecisionCardAuthorizationNetworkDetailsVisaStandInProcessingReason = "payment_fraud_disruption_acquirer_block"
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaStandInProcessingReasonOther                                                    RealTimeDecisionCardAuthorizationNetworkDetailsVisaStandInProcessingReason = "other"
)

func (r RealTimeDecisionCardAuthorizationNetworkDetailsVisaStandInProcessingReason) IsKnown() bool {
	switch r {
	case RealTimeDecisionCardAuthorizationNetworkDetailsVisaStandInProcessingReasonIssuerError, RealTimeDecisionCardAuthorizationNetworkDetailsVisaStandInProcessingReasonInvalidPhysicalCard, RealTimeDecisionCardAuthorizationNetworkDetailsVisaStandInProcessingReasonInvalidCardholderAuthenticationVerificationValue, RealTimeDecisionCardAuthorizationNetworkDetailsVisaStandInProcessingReasonInternalVisaError, RealTimeDecisionCardAuthorizationNetworkDetailsVisaStandInProcessingReasonMerchantTransactionAdvisoryServiceAuthenticationRequired, RealTimeDecisionCardAuthorizationNetworkDetailsVisaStandInProcessingReasonPaymentFraudDisruptionAcquirerBlock, RealTimeDecisionCardAuthorizationNetworkDetailsVisaStandInProcessingReasonOther:
		return true
	}
	return false
}

// Network-specific identifiers for a specific request or transaction.
type RealTimeDecisionCardAuthorizationNetworkIdentifiers struct {
	// The randomly generated 6-character Authorization Identification Response code
	// sent back to the acquirer in an approved response.
	AuthorizationIdentificationResponse string `json:"authorization_identification_response,required,nullable"`
	// A life-cycle identifier used across e.g., an authorization and a reversal.
	// Expected to be unique per acquirer within a window of time. For some card
	// networks the retrieval reference number includes the trace counter.
	RetrievalReferenceNumber string `json:"retrieval_reference_number,required,nullable"`
	// A counter used to verify an individual authorization. Expected to be unique per
	// acquirer within a window of time.
	TraceNumber string `json:"trace_number,required,nullable"`
	// A globally unique transaction identifier provided by the card network, used
	// across multiple life-cycle requests.
	TransactionID string                                                  `json:"transaction_id,required,nullable"`
	JSON          realTimeDecisionCardAuthorizationNetworkIdentifiersJSON `json:"-"`
}

// realTimeDecisionCardAuthorizationNetworkIdentifiersJSON contains the JSON
// metadata for the struct [RealTimeDecisionCardAuthorizationNetworkIdentifiers]
type realTimeDecisionCardAuthorizationNetworkIdentifiersJSON struct {
	AuthorizationIdentificationResponse apijson.Field
	RetrievalReferenceNumber            apijson.Field
	TraceNumber                         apijson.Field
	TransactionID                       apijson.Field
	raw                                 string
	ExtraFields                         map[string]apijson.Field
}

func (r *RealTimeDecisionCardAuthorizationNetworkIdentifiers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimeDecisionCardAuthorizationNetworkIdentifiersJSON) RawJSON() string {
	return r.raw
}

// The processing category describes the intent behind the authorization, such as
// whether it was used for bill payments or an automatic fuel dispenser.
type RealTimeDecisionCardAuthorizationProcessingCategory string

const (
	RealTimeDecisionCardAuthorizationProcessingCategoryAccountFunding         RealTimeDecisionCardAuthorizationProcessingCategory = "account_funding"
	RealTimeDecisionCardAuthorizationProcessingCategoryAutomaticFuelDispenser RealTimeDecisionCardAuthorizationProcessingCategory = "automatic_fuel_dispenser"
	RealTimeDecisionCardAuthorizationProcessingCategoryBillPayment            RealTimeDecisionCardAuthorizationProcessingCategory = "bill_payment"
	RealTimeDecisionCardAuthorizationProcessingCategoryOriginalCredit         RealTimeDecisionCardAuthorizationProcessingCategory = "original_credit"
	RealTimeDecisionCardAuthorizationProcessingCategoryPurchase               RealTimeDecisionCardAuthorizationProcessingCategory = "purchase"
	RealTimeDecisionCardAuthorizationProcessingCategoryQuasiCash              RealTimeDecisionCardAuthorizationProcessingCategory = "quasi_cash"
	RealTimeDecisionCardAuthorizationProcessingCategoryRefund                 RealTimeDecisionCardAuthorizationProcessingCategory = "refund"
	RealTimeDecisionCardAuthorizationProcessingCategoryCashDisbursement       RealTimeDecisionCardAuthorizationProcessingCategory = "cash_disbursement"
	RealTimeDecisionCardAuthorizationProcessingCategoryUnknown                RealTimeDecisionCardAuthorizationProcessingCategory = "unknown"
)

func (r RealTimeDecisionCardAuthorizationProcessingCategory) IsKnown() bool {
	switch r {
	case RealTimeDecisionCardAuthorizationProcessingCategoryAccountFunding, RealTimeDecisionCardAuthorizationProcessingCategoryAutomaticFuelDispenser, RealTimeDecisionCardAuthorizationProcessingCategoryBillPayment, RealTimeDecisionCardAuthorizationProcessingCategoryOriginalCredit, RealTimeDecisionCardAuthorizationProcessingCategoryPurchase, RealTimeDecisionCardAuthorizationProcessingCategoryQuasiCash, RealTimeDecisionCardAuthorizationProcessingCategoryRefund, RealTimeDecisionCardAuthorizationProcessingCategoryCashDisbursement, RealTimeDecisionCardAuthorizationProcessingCategoryUnknown:
		return true
	}
	return false
}

// Fields specific to the type of request, such as an incremental authorization.
type RealTimeDecisionCardAuthorizationRequestDetails struct {
	// The type of this request (e.g., an initial authorization or an incremental
	// authorization).
	Category RealTimeDecisionCardAuthorizationRequestDetailsCategory `json:"category,required"`
	// Fields specific to the category `incremental_authorization`.
	IncrementalAuthorization RealTimeDecisionCardAuthorizationRequestDetailsIncrementalAuthorization `json:"incremental_authorization,required,nullable"`
	// Fields specific to the category `initial_authorization`.
	InitialAuthorization interface{}                                         `json:"initial_authorization,required,nullable"`
	JSON                 realTimeDecisionCardAuthorizationRequestDetailsJSON `json:"-"`
}

// realTimeDecisionCardAuthorizationRequestDetailsJSON contains the JSON metadata
// for the struct [RealTimeDecisionCardAuthorizationRequestDetails]
type realTimeDecisionCardAuthorizationRequestDetailsJSON struct {
	Category                 apijson.Field
	IncrementalAuthorization apijson.Field
	InitialAuthorization     apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *RealTimeDecisionCardAuthorizationRequestDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimeDecisionCardAuthorizationRequestDetailsJSON) RawJSON() string {
	return r.raw
}

// The type of this request (e.g., an initial authorization or an incremental
// authorization).
type RealTimeDecisionCardAuthorizationRequestDetailsCategory string

const (
	RealTimeDecisionCardAuthorizationRequestDetailsCategoryInitialAuthorization     RealTimeDecisionCardAuthorizationRequestDetailsCategory = "initial_authorization"
	RealTimeDecisionCardAuthorizationRequestDetailsCategoryIncrementalAuthorization RealTimeDecisionCardAuthorizationRequestDetailsCategory = "incremental_authorization"
)

func (r RealTimeDecisionCardAuthorizationRequestDetailsCategory) IsKnown() bool {
	switch r {
	case RealTimeDecisionCardAuthorizationRequestDetailsCategoryInitialAuthorization, RealTimeDecisionCardAuthorizationRequestDetailsCategoryIncrementalAuthorization:
		return true
	}
	return false
}

// Fields specific to the category `incremental_authorization`.
type RealTimeDecisionCardAuthorizationRequestDetailsIncrementalAuthorization struct {
	// The card payment for this authorization and increment.
	CardPaymentID string `json:"card_payment_id,required"`
	// The identifier of the card authorization this request is attempting to
	// increment.
	OriginalCardAuthorizationID string                                                                      `json:"original_card_authorization_id,required"`
	JSON                        realTimeDecisionCardAuthorizationRequestDetailsIncrementalAuthorizationJSON `json:"-"`
}

// realTimeDecisionCardAuthorizationRequestDetailsIncrementalAuthorizationJSON
// contains the JSON metadata for the struct
// [RealTimeDecisionCardAuthorizationRequestDetailsIncrementalAuthorization]
type realTimeDecisionCardAuthorizationRequestDetailsIncrementalAuthorizationJSON struct {
	CardPaymentID               apijson.Field
	OriginalCardAuthorizationID apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *RealTimeDecisionCardAuthorizationRequestDetailsIncrementalAuthorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimeDecisionCardAuthorizationRequestDetailsIncrementalAuthorizationJSON) RawJSON() string {
	return r.raw
}

// Fields related to verification of cardholder-provided values.
type RealTimeDecisionCardAuthorizationVerification struct {
	// Fields related to verification of the Card Verification Code, a 3-digit code on
	// the back of the card.
	CardVerificationCode RealTimeDecisionCardAuthorizationVerificationCardVerificationCode `json:"card_verification_code,required"`
	// Cardholder address provided in the authorization request and the address on file
	// we verified it against.
	CardholderAddress RealTimeDecisionCardAuthorizationVerificationCardholderAddress `json:"cardholder_address,required"`
	JSON              realTimeDecisionCardAuthorizationVerificationJSON              `json:"-"`
}

// realTimeDecisionCardAuthorizationVerificationJSON contains the JSON metadata for
// the struct [RealTimeDecisionCardAuthorizationVerification]
type realTimeDecisionCardAuthorizationVerificationJSON struct {
	CardVerificationCode apijson.Field
	CardholderAddress    apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *RealTimeDecisionCardAuthorizationVerification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimeDecisionCardAuthorizationVerificationJSON) RawJSON() string {
	return r.raw
}

// Fields related to verification of the Card Verification Code, a 3-digit code on
// the back of the card.
type RealTimeDecisionCardAuthorizationVerificationCardVerificationCode struct {
	// The result of verifying the Card Verification Code.
	Result RealTimeDecisionCardAuthorizationVerificationCardVerificationCodeResult `json:"result,required"`
	JSON   realTimeDecisionCardAuthorizationVerificationCardVerificationCodeJSON   `json:"-"`
}

// realTimeDecisionCardAuthorizationVerificationCardVerificationCodeJSON contains
// the JSON metadata for the struct
// [RealTimeDecisionCardAuthorizationVerificationCardVerificationCode]
type realTimeDecisionCardAuthorizationVerificationCardVerificationCodeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RealTimeDecisionCardAuthorizationVerificationCardVerificationCode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimeDecisionCardAuthorizationVerificationCardVerificationCodeJSON) RawJSON() string {
	return r.raw
}

// The result of verifying the Card Verification Code.
type RealTimeDecisionCardAuthorizationVerificationCardVerificationCodeResult string

const (
	RealTimeDecisionCardAuthorizationVerificationCardVerificationCodeResultNotChecked RealTimeDecisionCardAuthorizationVerificationCardVerificationCodeResult = "not_checked"
	RealTimeDecisionCardAuthorizationVerificationCardVerificationCodeResultMatch      RealTimeDecisionCardAuthorizationVerificationCardVerificationCodeResult = "match"
	RealTimeDecisionCardAuthorizationVerificationCardVerificationCodeResultNoMatch    RealTimeDecisionCardAuthorizationVerificationCardVerificationCodeResult = "no_match"
)

func (r RealTimeDecisionCardAuthorizationVerificationCardVerificationCodeResult) IsKnown() bool {
	switch r {
	case RealTimeDecisionCardAuthorizationVerificationCardVerificationCodeResultNotChecked, RealTimeDecisionCardAuthorizationVerificationCardVerificationCodeResultMatch, RealTimeDecisionCardAuthorizationVerificationCardVerificationCodeResultNoMatch:
		return true
	}
	return false
}

// Cardholder address provided in the authorization request and the address on file
// we verified it against.
type RealTimeDecisionCardAuthorizationVerificationCardholderAddress struct {
	// Line 1 of the address on file for the cardholder.
	ActualLine1 string `json:"actual_line1,required,nullable"`
	// The postal code of the address on file for the cardholder.
	ActualPostalCode string `json:"actual_postal_code,required,nullable"`
	// The cardholder address line 1 provided for verification in the authorization
	// request.
	ProvidedLine1 string `json:"provided_line1,required,nullable"`
	// The postal code provided for verification in the authorization request.
	ProvidedPostalCode string `json:"provided_postal_code,required,nullable"`
	// The address verification result returned to the card network.
	Result RealTimeDecisionCardAuthorizationVerificationCardholderAddressResult `json:"result,required"`
	JSON   realTimeDecisionCardAuthorizationVerificationCardholderAddressJSON   `json:"-"`
}

// realTimeDecisionCardAuthorizationVerificationCardholderAddressJSON contains the
// JSON metadata for the struct
// [RealTimeDecisionCardAuthorizationVerificationCardholderAddress]
type realTimeDecisionCardAuthorizationVerificationCardholderAddressJSON struct {
	ActualLine1        apijson.Field
	ActualPostalCode   apijson.Field
	ProvidedLine1      apijson.Field
	ProvidedPostalCode apijson.Field
	Result             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RealTimeDecisionCardAuthorizationVerificationCardholderAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimeDecisionCardAuthorizationVerificationCardholderAddressJSON) RawJSON() string {
	return r.raw
}

// The address verification result returned to the card network.
type RealTimeDecisionCardAuthorizationVerificationCardholderAddressResult string

const (
	RealTimeDecisionCardAuthorizationVerificationCardholderAddressResultNotChecked                       RealTimeDecisionCardAuthorizationVerificationCardholderAddressResult = "not_checked"
	RealTimeDecisionCardAuthorizationVerificationCardholderAddressResultPostalCodeMatchAddressNoMatch    RealTimeDecisionCardAuthorizationVerificationCardholderAddressResult = "postal_code_match_address_no_match"
	RealTimeDecisionCardAuthorizationVerificationCardholderAddressResultPostalCodeNoMatchAddressMatch    RealTimeDecisionCardAuthorizationVerificationCardholderAddressResult = "postal_code_no_match_address_match"
	RealTimeDecisionCardAuthorizationVerificationCardholderAddressResultMatch                            RealTimeDecisionCardAuthorizationVerificationCardholderAddressResult = "match"
	RealTimeDecisionCardAuthorizationVerificationCardholderAddressResultNoMatch                          RealTimeDecisionCardAuthorizationVerificationCardholderAddressResult = "no_match"
	RealTimeDecisionCardAuthorizationVerificationCardholderAddressResultPostalCodeMatchAddressNotChecked RealTimeDecisionCardAuthorizationVerificationCardholderAddressResult = "postal_code_match_address_not_checked"
)

func (r RealTimeDecisionCardAuthorizationVerificationCardholderAddressResult) IsKnown() bool {
	switch r {
	case RealTimeDecisionCardAuthorizationVerificationCardholderAddressResultNotChecked, RealTimeDecisionCardAuthorizationVerificationCardholderAddressResultPostalCodeMatchAddressNoMatch, RealTimeDecisionCardAuthorizationVerificationCardholderAddressResultPostalCodeNoMatchAddressMatch, RealTimeDecisionCardAuthorizationVerificationCardholderAddressResultMatch, RealTimeDecisionCardAuthorizationVerificationCardholderAddressResultNoMatch, RealTimeDecisionCardAuthorizationVerificationCardholderAddressResultPostalCodeMatchAddressNotChecked:
		return true
	}
	return false
}

// The category of the Real-Time Decision.
type RealTimeDecisionCategory string

const (
	RealTimeDecisionCategoryCardAuthorizationRequested           RealTimeDecisionCategory = "card_authorization_requested"
	RealTimeDecisionCategoryCardAuthenticationRequested          RealTimeDecisionCategory = "card_authentication_requested"
	RealTimeDecisionCategoryCardAuthenticationChallengeRequested RealTimeDecisionCategory = "card_authentication_challenge_requested"
	RealTimeDecisionCategoryDigitalWalletTokenRequested          RealTimeDecisionCategory = "digital_wallet_token_requested"
	RealTimeDecisionCategoryDigitalWalletAuthenticationRequested RealTimeDecisionCategory = "digital_wallet_authentication_requested"
)

func (r RealTimeDecisionCategory) IsKnown() bool {
	switch r {
	case RealTimeDecisionCategoryCardAuthorizationRequested, RealTimeDecisionCategoryCardAuthenticationRequested, RealTimeDecisionCategoryCardAuthenticationChallengeRequested, RealTimeDecisionCategoryDigitalWalletTokenRequested, RealTimeDecisionCategoryDigitalWalletAuthenticationRequested:
		return true
	}
	return false
}

// Fields related to a digital wallet authentication attempt.
type RealTimeDecisionDigitalWalletAuthentication struct {
	// The identifier of the Card that is being tokenized.
	CardID string `json:"card_id,required"`
	// The channel to send the card user their one-time passcode.
	Channel RealTimeDecisionDigitalWalletAuthenticationChannel `json:"channel,required"`
	// The digital wallet app being used.
	DigitalWallet RealTimeDecisionDigitalWalletAuthenticationDigitalWallet `json:"digital_wallet,required"`
	// The email to send the one-time passcode to if `channel` is equal to `email`.
	Email string `json:"email,required,nullable"`
	// The one-time passcode to send the card user.
	OneTimePasscode string `json:"one_time_passcode,required"`
	// The phone number to send the one-time passcode to if `channel` is equal to
	// `sms`.
	Phone string `json:"phone,required,nullable"`
	// Whether your application successfully delivered the one-time passcode.
	Result RealTimeDecisionDigitalWalletAuthenticationResult `json:"result,required,nullable"`
	JSON   realTimeDecisionDigitalWalletAuthenticationJSON   `json:"-"`
}

// realTimeDecisionDigitalWalletAuthenticationJSON contains the JSON metadata for
// the struct [RealTimeDecisionDigitalWalletAuthentication]
type realTimeDecisionDigitalWalletAuthenticationJSON struct {
	CardID          apijson.Field
	Channel         apijson.Field
	DigitalWallet   apijson.Field
	Email           apijson.Field
	OneTimePasscode apijson.Field
	Phone           apijson.Field
	Result          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RealTimeDecisionDigitalWalletAuthentication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimeDecisionDigitalWalletAuthenticationJSON) RawJSON() string {
	return r.raw
}

// The channel to send the card user their one-time passcode.
type RealTimeDecisionDigitalWalletAuthenticationChannel string

const (
	RealTimeDecisionDigitalWalletAuthenticationChannelSMS   RealTimeDecisionDigitalWalletAuthenticationChannel = "sms"
	RealTimeDecisionDigitalWalletAuthenticationChannelEmail RealTimeDecisionDigitalWalletAuthenticationChannel = "email"
)

func (r RealTimeDecisionDigitalWalletAuthenticationChannel) IsKnown() bool {
	switch r {
	case RealTimeDecisionDigitalWalletAuthenticationChannelSMS, RealTimeDecisionDigitalWalletAuthenticationChannelEmail:
		return true
	}
	return false
}

// The digital wallet app being used.
type RealTimeDecisionDigitalWalletAuthenticationDigitalWallet string

const (
	RealTimeDecisionDigitalWalletAuthenticationDigitalWalletApplePay   RealTimeDecisionDigitalWalletAuthenticationDigitalWallet = "apple_pay"
	RealTimeDecisionDigitalWalletAuthenticationDigitalWalletGooglePay  RealTimeDecisionDigitalWalletAuthenticationDigitalWallet = "google_pay"
	RealTimeDecisionDigitalWalletAuthenticationDigitalWalletSamsungPay RealTimeDecisionDigitalWalletAuthenticationDigitalWallet = "samsung_pay"
	RealTimeDecisionDigitalWalletAuthenticationDigitalWalletUnknown    RealTimeDecisionDigitalWalletAuthenticationDigitalWallet = "unknown"
)

func (r RealTimeDecisionDigitalWalletAuthenticationDigitalWallet) IsKnown() bool {
	switch r {
	case RealTimeDecisionDigitalWalletAuthenticationDigitalWalletApplePay, RealTimeDecisionDigitalWalletAuthenticationDigitalWalletGooglePay, RealTimeDecisionDigitalWalletAuthenticationDigitalWalletSamsungPay, RealTimeDecisionDigitalWalletAuthenticationDigitalWalletUnknown:
		return true
	}
	return false
}

// Whether your application successfully delivered the one-time passcode.
type RealTimeDecisionDigitalWalletAuthenticationResult string

const (
	RealTimeDecisionDigitalWalletAuthenticationResultSuccess RealTimeDecisionDigitalWalletAuthenticationResult = "success"
	RealTimeDecisionDigitalWalletAuthenticationResultFailure RealTimeDecisionDigitalWalletAuthenticationResult = "failure"
)

func (r RealTimeDecisionDigitalWalletAuthenticationResult) IsKnown() bool {
	switch r {
	case RealTimeDecisionDigitalWalletAuthenticationResultSuccess, RealTimeDecisionDigitalWalletAuthenticationResultFailure:
		return true
	}
	return false
}

// Fields related to a digital wallet token provisioning attempt.
type RealTimeDecisionDigitalWalletToken struct {
	// The identifier of the Card that is being tokenized.
	CardID string `json:"card_id,required"`
	// Whether or not the provisioning request was approved. This will be null until
	// the real time decision is responded to.
	Decision RealTimeDecisionDigitalWalletTokenDecision `json:"decision,required,nullable"`
	// Device that is being used to provision the digital wallet token.
	Device RealTimeDecisionDigitalWalletTokenDevice `json:"device,required"`
	// The digital wallet app being used.
	DigitalWallet RealTimeDecisionDigitalWalletTokenDigitalWallet `json:"digital_wallet,required"`
	JSON          realTimeDecisionDigitalWalletTokenJSON          `json:"-"`
}

// realTimeDecisionDigitalWalletTokenJSON contains the JSON metadata for the struct
// [RealTimeDecisionDigitalWalletToken]
type realTimeDecisionDigitalWalletTokenJSON struct {
	CardID        apijson.Field
	Decision      apijson.Field
	Device        apijson.Field
	DigitalWallet apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RealTimeDecisionDigitalWalletToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimeDecisionDigitalWalletTokenJSON) RawJSON() string {
	return r.raw
}

// Whether or not the provisioning request was approved. This will be null until
// the real time decision is responded to.
type RealTimeDecisionDigitalWalletTokenDecision string

const (
	RealTimeDecisionDigitalWalletTokenDecisionApprove RealTimeDecisionDigitalWalletTokenDecision = "approve"
	RealTimeDecisionDigitalWalletTokenDecisionDecline RealTimeDecisionDigitalWalletTokenDecision = "decline"
)

func (r RealTimeDecisionDigitalWalletTokenDecision) IsKnown() bool {
	switch r {
	case RealTimeDecisionDigitalWalletTokenDecisionApprove, RealTimeDecisionDigitalWalletTokenDecisionDecline:
		return true
	}
	return false
}

// Device that is being used to provision the digital wallet token.
type RealTimeDecisionDigitalWalletTokenDevice struct {
	// ID assigned to the device by the digital wallet provider.
	Identifier string                                       `json:"identifier,required,nullable"`
	JSON       realTimeDecisionDigitalWalletTokenDeviceJSON `json:"-"`
}

// realTimeDecisionDigitalWalletTokenDeviceJSON contains the JSON metadata for the
// struct [RealTimeDecisionDigitalWalletTokenDevice]
type realTimeDecisionDigitalWalletTokenDeviceJSON struct {
	Identifier  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RealTimeDecisionDigitalWalletTokenDevice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realTimeDecisionDigitalWalletTokenDeviceJSON) RawJSON() string {
	return r.raw
}

// The digital wallet app being used.
type RealTimeDecisionDigitalWalletTokenDigitalWallet string

const (
	RealTimeDecisionDigitalWalletTokenDigitalWalletApplePay   RealTimeDecisionDigitalWalletTokenDigitalWallet = "apple_pay"
	RealTimeDecisionDigitalWalletTokenDigitalWalletGooglePay  RealTimeDecisionDigitalWalletTokenDigitalWallet = "google_pay"
	RealTimeDecisionDigitalWalletTokenDigitalWalletSamsungPay RealTimeDecisionDigitalWalletTokenDigitalWallet = "samsung_pay"
	RealTimeDecisionDigitalWalletTokenDigitalWalletUnknown    RealTimeDecisionDigitalWalletTokenDigitalWallet = "unknown"
)

func (r RealTimeDecisionDigitalWalletTokenDigitalWallet) IsKnown() bool {
	switch r {
	case RealTimeDecisionDigitalWalletTokenDigitalWalletApplePay, RealTimeDecisionDigitalWalletTokenDigitalWalletGooglePay, RealTimeDecisionDigitalWalletTokenDigitalWalletSamsungPay, RealTimeDecisionDigitalWalletTokenDigitalWalletUnknown:
		return true
	}
	return false
}

// The status of the Real-Time Decision.
type RealTimeDecisionStatus string

const (
	RealTimeDecisionStatusPending   RealTimeDecisionStatus = "pending"
	RealTimeDecisionStatusResponded RealTimeDecisionStatus = "responded"
	RealTimeDecisionStatusTimedOut  RealTimeDecisionStatus = "timed_out"
)

func (r RealTimeDecisionStatus) IsKnown() bool {
	switch r {
	case RealTimeDecisionStatusPending, RealTimeDecisionStatusResponded, RealTimeDecisionStatusTimedOut:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `real_time_decision`.
type RealTimeDecisionType string

const (
	RealTimeDecisionTypeRealTimeDecision RealTimeDecisionType = "real_time_decision"
)

func (r RealTimeDecisionType) IsKnown() bool {
	switch r {
	case RealTimeDecisionTypeRealTimeDecision:
		return true
	}
	return false
}

type RealTimeDecisionActionParams struct {
	// If the Real-Time Decision relates to a 3DS card authentication attempt, this
	// object contains your response to the authentication.
	CardAuthentication param.Field[RealTimeDecisionActionParamsCardAuthentication] `json:"card_authentication"`
	// If the Real-Time Decision relates to 3DS card authentication challenge delivery,
	// this object contains your response.
	CardAuthenticationChallenge param.Field[RealTimeDecisionActionParamsCardAuthenticationChallenge] `json:"card_authentication_challenge"`
	// If the Real-Time Decision relates to a card authorization attempt, this object
	// contains your response to the authorization.
	CardAuthorization param.Field[RealTimeDecisionActionParamsCardAuthorization] `json:"card_authorization"`
	// If the Real-Time Decision relates to a digital wallet authentication attempt,
	// this object contains your response to the authentication.
	DigitalWalletAuthentication param.Field[RealTimeDecisionActionParamsDigitalWalletAuthentication] `json:"digital_wallet_authentication"`
	// If the Real-Time Decision relates to a digital wallet token provisioning
	// attempt, this object contains your response to the attempt.
	DigitalWalletToken param.Field[RealTimeDecisionActionParamsDigitalWalletToken] `json:"digital_wallet_token"`
}

func (r RealTimeDecisionActionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If the Real-Time Decision relates to a 3DS card authentication attempt, this
// object contains your response to the authentication.
type RealTimeDecisionActionParamsCardAuthentication struct {
	// Whether the card authentication attempt should be approved or declined.
	Decision param.Field[RealTimeDecisionActionParamsCardAuthenticationDecision] `json:"decision,required"`
}

func (r RealTimeDecisionActionParamsCardAuthentication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether the card authentication attempt should be approved or declined.
type RealTimeDecisionActionParamsCardAuthenticationDecision string

const (
	RealTimeDecisionActionParamsCardAuthenticationDecisionApprove   RealTimeDecisionActionParamsCardAuthenticationDecision = "approve"
	RealTimeDecisionActionParamsCardAuthenticationDecisionChallenge RealTimeDecisionActionParamsCardAuthenticationDecision = "challenge"
	RealTimeDecisionActionParamsCardAuthenticationDecisionDeny      RealTimeDecisionActionParamsCardAuthenticationDecision = "deny"
)

func (r RealTimeDecisionActionParamsCardAuthenticationDecision) IsKnown() bool {
	switch r {
	case RealTimeDecisionActionParamsCardAuthenticationDecisionApprove, RealTimeDecisionActionParamsCardAuthenticationDecisionChallenge, RealTimeDecisionActionParamsCardAuthenticationDecisionDeny:
		return true
	}
	return false
}

// If the Real-Time Decision relates to 3DS card authentication challenge delivery,
// this object contains your response.
type RealTimeDecisionActionParamsCardAuthenticationChallenge struct {
	// Whether the card authentication challenge was successfully delivered to the
	// cardholder.
	Result param.Field[RealTimeDecisionActionParamsCardAuthenticationChallengeResult] `json:"result,required"`
}

func (r RealTimeDecisionActionParamsCardAuthenticationChallenge) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether the card authentication challenge was successfully delivered to the
// cardholder.
type RealTimeDecisionActionParamsCardAuthenticationChallengeResult string

const (
	RealTimeDecisionActionParamsCardAuthenticationChallengeResultSuccess RealTimeDecisionActionParamsCardAuthenticationChallengeResult = "success"
	RealTimeDecisionActionParamsCardAuthenticationChallengeResultFailure RealTimeDecisionActionParamsCardAuthenticationChallengeResult = "failure"
)

func (r RealTimeDecisionActionParamsCardAuthenticationChallengeResult) IsKnown() bool {
	switch r {
	case RealTimeDecisionActionParamsCardAuthenticationChallengeResultSuccess, RealTimeDecisionActionParamsCardAuthenticationChallengeResultFailure:
		return true
	}
	return false
}

// If the Real-Time Decision relates to a card authorization attempt, this object
// contains your response to the authorization.
type RealTimeDecisionActionParamsCardAuthorization struct {
	// Whether the card authorization should be approved or declined.
	Decision param.Field[RealTimeDecisionActionParamsCardAuthorizationDecision] `json:"decision,required"`
	// If your application approves the authorization, this contains metadata about
	// your decision to approve. Your response here is advisory to the acquiring bank.
	// The bank may choose to reverse the authorization if you approve the transaction
	// but indicate the address does not match.
	Approval param.Field[RealTimeDecisionActionParamsCardAuthorizationApproval] `json:"approval"`
	// If your application declines the authorization, this contains details about the
	// decline.
	Decline param.Field[RealTimeDecisionActionParamsCardAuthorizationDecline] `json:"decline"`
}

func (r RealTimeDecisionActionParamsCardAuthorization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether the card authorization should be approved or declined.
type RealTimeDecisionActionParamsCardAuthorizationDecision string

const (
	RealTimeDecisionActionParamsCardAuthorizationDecisionApprove RealTimeDecisionActionParamsCardAuthorizationDecision = "approve"
	RealTimeDecisionActionParamsCardAuthorizationDecisionDecline RealTimeDecisionActionParamsCardAuthorizationDecision = "decline"
)

func (r RealTimeDecisionActionParamsCardAuthorizationDecision) IsKnown() bool {
	switch r {
	case RealTimeDecisionActionParamsCardAuthorizationDecisionApprove, RealTimeDecisionActionParamsCardAuthorizationDecisionDecline:
		return true
	}
	return false
}

// If your application approves the authorization, this contains metadata about
// your decision to approve. Your response here is advisory to the acquiring bank.
// The bank may choose to reverse the authorization if you approve the transaction
// but indicate the address does not match.
type RealTimeDecisionActionParamsCardAuthorizationApproval struct {
	// Your decisions on whether or not each provided address component is a match.
	// Your response here is evaluated against the customer's provided `postal_code`
	// and `line1`, and an appropriate network response is generated. For more
	// information, see our
	// [Address Verification System Codes and Overrides](https://increase.com/documentation/address-verification-system-codes-and-overrides)
	// guide.
	CardholderAddressVerificationResult param.Field[RealTimeDecisionActionParamsCardAuthorizationApprovalCardholderAddressVerificationResult] `json:"cardholder_address_verification_result"`
}

func (r RealTimeDecisionActionParamsCardAuthorizationApproval) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Your decisions on whether or not each provided address component is a match.
// Your response here is evaluated against the customer's provided `postal_code`
// and `line1`, and an appropriate network response is generated. For more
// information, see our
// [Address Verification System Codes and Overrides](https://increase.com/documentation/address-verification-system-codes-and-overrides)
// guide.
type RealTimeDecisionActionParamsCardAuthorizationApprovalCardholderAddressVerificationResult struct {
	// Your decision on the address line of the provided address.
	Line1 param.Field[RealTimeDecisionActionParamsCardAuthorizationApprovalCardholderAddressVerificationResultLine1] `json:"line1,required"`
	// Your decision on the postal code of the provided address.
	PostalCode param.Field[RealTimeDecisionActionParamsCardAuthorizationApprovalCardholderAddressVerificationResultPostalCode] `json:"postal_code,required"`
}

func (r RealTimeDecisionActionParamsCardAuthorizationApprovalCardholderAddressVerificationResult) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Your decision on the address line of the provided address.
type RealTimeDecisionActionParamsCardAuthorizationApprovalCardholderAddressVerificationResultLine1 string

const (
	RealTimeDecisionActionParamsCardAuthorizationApprovalCardholderAddressVerificationResultLine1Match   RealTimeDecisionActionParamsCardAuthorizationApprovalCardholderAddressVerificationResultLine1 = "match"
	RealTimeDecisionActionParamsCardAuthorizationApprovalCardholderAddressVerificationResultLine1NoMatch RealTimeDecisionActionParamsCardAuthorizationApprovalCardholderAddressVerificationResultLine1 = "no_match"
)

func (r RealTimeDecisionActionParamsCardAuthorizationApprovalCardholderAddressVerificationResultLine1) IsKnown() bool {
	switch r {
	case RealTimeDecisionActionParamsCardAuthorizationApprovalCardholderAddressVerificationResultLine1Match, RealTimeDecisionActionParamsCardAuthorizationApprovalCardholderAddressVerificationResultLine1NoMatch:
		return true
	}
	return false
}

// Your decision on the postal code of the provided address.
type RealTimeDecisionActionParamsCardAuthorizationApprovalCardholderAddressVerificationResultPostalCode string

const (
	RealTimeDecisionActionParamsCardAuthorizationApprovalCardholderAddressVerificationResultPostalCodeMatch   RealTimeDecisionActionParamsCardAuthorizationApprovalCardholderAddressVerificationResultPostalCode = "match"
	RealTimeDecisionActionParamsCardAuthorizationApprovalCardholderAddressVerificationResultPostalCodeNoMatch RealTimeDecisionActionParamsCardAuthorizationApprovalCardholderAddressVerificationResultPostalCode = "no_match"
)

func (r RealTimeDecisionActionParamsCardAuthorizationApprovalCardholderAddressVerificationResultPostalCode) IsKnown() bool {
	switch r {
	case RealTimeDecisionActionParamsCardAuthorizationApprovalCardholderAddressVerificationResultPostalCodeMatch, RealTimeDecisionActionParamsCardAuthorizationApprovalCardholderAddressVerificationResultPostalCodeNoMatch:
		return true
	}
	return false
}

// If your application declines the authorization, this contains details about the
// decline.
type RealTimeDecisionActionParamsCardAuthorizationDecline struct {
	// The reason the card authorization was declined. This translates to a specific
	// decline code that is sent to the card network.
	Reason param.Field[RealTimeDecisionActionParamsCardAuthorizationDeclineReason] `json:"reason,required"`
}

func (r RealTimeDecisionActionParamsCardAuthorizationDecline) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The reason the card authorization was declined. This translates to a specific
// decline code that is sent to the card network.
type RealTimeDecisionActionParamsCardAuthorizationDeclineReason string

const (
	RealTimeDecisionActionParamsCardAuthorizationDeclineReasonInsufficientFunds       RealTimeDecisionActionParamsCardAuthorizationDeclineReason = "insufficient_funds"
	RealTimeDecisionActionParamsCardAuthorizationDeclineReasonTransactionNeverAllowed RealTimeDecisionActionParamsCardAuthorizationDeclineReason = "transaction_never_allowed"
	RealTimeDecisionActionParamsCardAuthorizationDeclineReasonExceedsApprovalLimit    RealTimeDecisionActionParamsCardAuthorizationDeclineReason = "exceeds_approval_limit"
	RealTimeDecisionActionParamsCardAuthorizationDeclineReasonCardTemporarilyDisabled RealTimeDecisionActionParamsCardAuthorizationDeclineReason = "card_temporarily_disabled"
	RealTimeDecisionActionParamsCardAuthorizationDeclineReasonSuspectedFraud          RealTimeDecisionActionParamsCardAuthorizationDeclineReason = "suspected_fraud"
	RealTimeDecisionActionParamsCardAuthorizationDeclineReasonOther                   RealTimeDecisionActionParamsCardAuthorizationDeclineReason = "other"
)

func (r RealTimeDecisionActionParamsCardAuthorizationDeclineReason) IsKnown() bool {
	switch r {
	case RealTimeDecisionActionParamsCardAuthorizationDeclineReasonInsufficientFunds, RealTimeDecisionActionParamsCardAuthorizationDeclineReasonTransactionNeverAllowed, RealTimeDecisionActionParamsCardAuthorizationDeclineReasonExceedsApprovalLimit, RealTimeDecisionActionParamsCardAuthorizationDeclineReasonCardTemporarilyDisabled, RealTimeDecisionActionParamsCardAuthorizationDeclineReasonSuspectedFraud, RealTimeDecisionActionParamsCardAuthorizationDeclineReasonOther:
		return true
	}
	return false
}

// If the Real-Time Decision relates to a digital wallet authentication attempt,
// this object contains your response to the authentication.
type RealTimeDecisionActionParamsDigitalWalletAuthentication struct {
	// Whether your application was able to deliver the one-time passcode.
	Result param.Field[RealTimeDecisionActionParamsDigitalWalletAuthenticationResult] `json:"result,required"`
	// If your application was able to deliver the one-time passcode, this contains
	// metadata about the delivery. Exactly one of `phone` or `email` must be provided.
	Success param.Field[RealTimeDecisionActionParamsDigitalWalletAuthenticationSuccess] `json:"success"`
}

func (r RealTimeDecisionActionParamsDigitalWalletAuthentication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether your application was able to deliver the one-time passcode.
type RealTimeDecisionActionParamsDigitalWalletAuthenticationResult string

const (
	RealTimeDecisionActionParamsDigitalWalletAuthenticationResultSuccess RealTimeDecisionActionParamsDigitalWalletAuthenticationResult = "success"
	RealTimeDecisionActionParamsDigitalWalletAuthenticationResultFailure RealTimeDecisionActionParamsDigitalWalletAuthenticationResult = "failure"
)

func (r RealTimeDecisionActionParamsDigitalWalletAuthenticationResult) IsKnown() bool {
	switch r {
	case RealTimeDecisionActionParamsDigitalWalletAuthenticationResultSuccess, RealTimeDecisionActionParamsDigitalWalletAuthenticationResultFailure:
		return true
	}
	return false
}

// If your application was able to deliver the one-time passcode, this contains
// metadata about the delivery. Exactly one of `phone` or `email` must be provided.
type RealTimeDecisionActionParamsDigitalWalletAuthenticationSuccess struct {
	// The email address that was used to verify the cardholder via one-time passcode.
	Email param.Field[string] `json:"email" format:"email"`
	// The phone number that was used to verify the cardholder via one-time passcode
	// over SMS.
	Phone param.Field[string] `json:"phone"`
}

func (r RealTimeDecisionActionParamsDigitalWalletAuthenticationSuccess) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If the Real-Time Decision relates to a digital wallet token provisioning
// attempt, this object contains your response to the attempt.
type RealTimeDecisionActionParamsDigitalWalletToken struct {
	// If your application approves the provisioning attempt, this contains metadata
	// about the digital wallet token that will be generated.
	Approval param.Field[RealTimeDecisionActionParamsDigitalWalletTokenApproval] `json:"approval"`
	// If your application declines the provisioning attempt, this contains details
	// about the decline.
	Decline param.Field[RealTimeDecisionActionParamsDigitalWalletTokenDecline] `json:"decline"`
}

func (r RealTimeDecisionActionParamsDigitalWalletToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If your application approves the provisioning attempt, this contains metadata
// about the digital wallet token that will be generated.
type RealTimeDecisionActionParamsDigitalWalletTokenApproval struct {
	// An email address that can be used to verify the cardholder via one-time
	// passcode.
	Email param.Field[string] `json:"email" format:"email"`
	// A phone number that can be used to verify the cardholder via one-time passcode
	// over SMS.
	Phone param.Field[string] `json:"phone"`
}

func (r RealTimeDecisionActionParamsDigitalWalletTokenApproval) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If your application declines the provisioning attempt, this contains details
// about the decline.
type RealTimeDecisionActionParamsDigitalWalletTokenDecline struct {
	// Why the tokenization attempt was declined. This is for logging purposes only and
	// is not displayed to the end-user.
	Reason param.Field[string] `json:"reason"`
}

func (r RealTimeDecisionActionParamsDigitalWalletTokenDecline) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

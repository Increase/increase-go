// File generated from our OpenAPI spec by Stainless.

package increase

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/internal/shared"
	"github.com/increase/increase-go/option"
)

// RealTimeDecisionService contains methods and other services that help with
// interacting with the increase API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRealTimeDecisionService] method
// instead.
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
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("real_time_decisions/%s", realTimeDecisionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Action a Real-Time Decision
func (r *RealTimeDecisionService) Action(ctx context.Context, realTimeDecisionID string, body RealTimeDecisionActionParams, opts ...option.RequestOption) (res *RealTimeDecision, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("real_time_decisions/%s/action", realTimeDecisionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Real Time Decisions are created when your application needs to take action in
// real-time to some event such as a card authorization. Real time decisions are
// currently in beta; please contact support@increase.com if you're interested in
// trying them out!
type RealTimeDecision struct {
	// The Real-Time Decision identifier.
	ID string `json:"id,required"`
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
	JSON realTimeDecisionJSON
}

// realTimeDecisionJSON contains the JSON metadata for the struct
// [RealTimeDecision]
type realTimeDecisionJSON struct {
	ID                          apijson.Field
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

// Fields related to a card authorization.
type RealTimeDecisionCardAuthorization struct {
	// The identifier of the Account the authorization will debit.
	AccountID string `json:"account_id,required"`
	// The identifier of the Card that is being authorized.
	CardID string `json:"card_id,required"`
	// Whether or not the authorization was approved.
	Decision RealTimeDecisionCardAuthorizationDecision `json:"decision,required,nullable"`
	// The merchant identifier (commonly abbreviated as MID) of the merchant the card
	// is transacting with.
	MerchantAcceptorID string `json:"merchant_acceptor_id,required"`
	// The Merchant Category Code (commonly abbreviated as MCC) of the merchant the
	// card is transacting with.
	MerchantCategoryCode string `json:"merchant_category_code,required,nullable"`
	// The city the merchant resides in.
	MerchantCity string `json:"merchant_city,required,nullable"`
	// The country the merchant resides in.
	MerchantCountry string `json:"merchant_country,required,nullable"`
	// The merchant descriptor of the merchant the card is transacting with.
	MerchantDescriptor string `json:"merchant_descriptor,required"`
	// The payment network used to process this card authorization
	Network RealTimeDecisionCardAuthorizationNetwork `json:"network,required"`
	// Fields specific to the `network`
	NetworkDetails RealTimeDecisionCardAuthorizationNetworkDetails `json:"network_details,required"`
	// The amount of the attempted authorization in the currency the card user sees at
	// the time of purchase, in the minor unit of that currency. For dollars, for
	// example, this is cents.
	PresentmentAmount int64 `json:"presentment_amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the currency the
	// user sees at the time of purchase.
	PresentmentCurrency string `json:"presentment_currency,required"`
	// The amount of the attempted authorization in the currency it will be settled in.
	// This currency is the same as that of the Account the card belongs to.
	SettlementAmount int64 `json:"settlement_amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the currency the
	// transaction will be settled in.
	SettlementCurrency string `json:"settlement_currency,required"`
	JSON               realTimeDecisionCardAuthorizationJSON
}

// realTimeDecisionCardAuthorizationJSON contains the JSON metadata for the struct
// [RealTimeDecisionCardAuthorization]
type realTimeDecisionCardAuthorizationJSON struct {
	AccountID            apijson.Field
	CardID               apijson.Field
	Decision             apijson.Field
	MerchantAcceptorID   apijson.Field
	MerchantCategoryCode apijson.Field
	MerchantCity         apijson.Field
	MerchantCountry      apijson.Field
	MerchantDescriptor   apijson.Field
	Network              apijson.Field
	NetworkDetails       apijson.Field
	PresentmentAmount    apijson.Field
	PresentmentCurrency  apijson.Field
	SettlementAmount     apijson.Field
	SettlementCurrency   apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *RealTimeDecisionCardAuthorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not the authorization was approved.
type RealTimeDecisionCardAuthorizationDecision string

const (
	// Approve the authorization.
	RealTimeDecisionCardAuthorizationDecisionApprove RealTimeDecisionCardAuthorizationDecision = "approve"
	// Decline the authorization.
	RealTimeDecisionCardAuthorizationDecisionDecline RealTimeDecisionCardAuthorizationDecision = "decline"
)

// The payment network used to process this card authorization
type RealTimeDecisionCardAuthorizationNetwork string

const (
	// Visa
	RealTimeDecisionCardAuthorizationNetworkVisa RealTimeDecisionCardAuthorizationNetwork = "visa"
)

// Fields specific to the `network`
type RealTimeDecisionCardAuthorizationNetworkDetails struct {
	// Fields specific to the `visa` network
	Visa RealTimeDecisionCardAuthorizationNetworkDetailsVisa `json:"visa,required"`
	JSON realTimeDecisionCardAuthorizationNetworkDetailsJSON
}

// realTimeDecisionCardAuthorizationNetworkDetailsJSON contains the JSON metadata
// for the struct [RealTimeDecisionCardAuthorizationNetworkDetails]
type realTimeDecisionCardAuthorizationNetworkDetailsJSON struct {
	Visa        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RealTimeDecisionCardAuthorizationNetworkDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Fields specific to the `visa` network
type RealTimeDecisionCardAuthorizationNetworkDetailsVisa struct {
	// For electronic commerce transactions, this identifies the level of security used
	// in obtaining the customer's payment credential. For mail or telephone order
	// transactions, identifies the type of mail or telephone order.
	ElectronicCommerceIndicator RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator `json:"electronic_commerce_indicator,required,nullable"`
	// The method used to enter the cardholder's primary account number and card
	// expiration date
	PointOfServiceEntryMode shared.PointOfServiceEntryMode `json:"point_of_service_entry_mode,required,nullable"`
	JSON                    realTimeDecisionCardAuthorizationNetworkDetailsVisaJSON
}

// realTimeDecisionCardAuthorizationNetworkDetailsVisaJSON contains the JSON
// metadata for the struct [RealTimeDecisionCardAuthorizationNetworkDetailsVisa]
type realTimeDecisionCardAuthorizationNetworkDetailsVisaJSON struct {
	ElectronicCommerceIndicator apijson.Field
	PointOfServiceEntryMode     apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *RealTimeDecisionCardAuthorizationNetworkDetailsVisa) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// For electronic commerce transactions, this identifies the level of security used
// in obtaining the customer's payment credential. For mail or telephone order
// transactions, identifies the type of mail or telephone order.
type RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator string

const (
	// Single transaction of a mail/phone order: Use to indicate that the transaction
	// is a mail/phone order purchase, not a recurring transaction or installment
	// payment. For domestic transactions in the US region, this value may also
	// indicate one bill payment transaction in the card-present or card-absent
	// environments.
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorMailPhoneOrder RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "mail_phone_order"
	// Recurring transaction: Payment indicator used to indicate a recurring
	// transaction that originates from an acquirer in the US region.
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorRecurring RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "recurring"
	// Installment payment: Payment indicator used to indicate one purchase of goods or
	// services that is billed to the account in multiple charges over a period of time
	// agreed upon by the cardholder and merchant from transactions that originate from
	// an acquirer in the US region.
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorInstallment RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "installment"
	// Unknown classification: other mail order: Use to indicate that the type of
	// mail/telephone order is unknown.
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorUnknownMailPhoneOrder RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "unknown_mail_phone_order"
	// Secure electronic commerce transaction: Use to indicate that the electronic
	// commerce transaction has been authenticated using e.g., 3-D Secure
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorSecureElectronicCommerce RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "secure_electronic_commerce"
	// Non-authenticated security transaction at a 3-D Secure-capable merchant, and
	// merchant attempted to authenticate the cardholder using 3-D Secure: Use to
	// identify an electronic commerce transaction where the merchant attempted to
	// authenticate the cardholder using 3-D Secure, but was unable to complete the
	// authentication because the issuer or cardholder does not participate in the 3-D
	// Secure program.
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransactionAt3DSCapableMerchant RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction_at_3ds_capable_merchant"
	// Non-authenticated security transaction: Use to identify an electronic commerce
	// transaction that uses data encryption for security however , cardholder
	// authentication is not performed using 3-D Secure.
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonAuthenticatedSecurityTransaction RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "non_authenticated_security_transaction"
	// Non-secure transaction: Use to identify an electronic commerce transaction that
	// has no data protection.
	RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicatorNonSecureTransaction RealTimeDecisionCardAuthorizationNetworkDetailsVisaElectronicCommerceIndicator = "non_secure_transaction"
)

// The category of the Real-Time Decision.
type RealTimeDecisionCategory string

const (
	// A card is being authorized.
	RealTimeDecisionCategoryCardAuthorizationRequested RealTimeDecisionCategory = "card_authorization_requested"
	// A card is being loaded into a digital wallet.
	RealTimeDecisionCategoryDigitalWalletTokenRequested RealTimeDecisionCategory = "digital_wallet_token_requested"
	// A card is being loaded into a digital wallet and requires cardholder
	// authentication.
	RealTimeDecisionCategoryDigitalWalletAuthenticationRequested RealTimeDecisionCategory = "digital_wallet_authentication_requested"
)

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
	JSON   realTimeDecisionDigitalWalletAuthenticationJSON
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

// The channel to send the card user their one-time passcode.
type RealTimeDecisionDigitalWalletAuthenticationChannel string

const (
	// Send one-time passcodes over SMS.
	RealTimeDecisionDigitalWalletAuthenticationChannelSMS RealTimeDecisionDigitalWalletAuthenticationChannel = "sms"
	// Send one-time passcodes over email.
	RealTimeDecisionDigitalWalletAuthenticationChannelEmail RealTimeDecisionDigitalWalletAuthenticationChannel = "email"
)

// The digital wallet app being used.
type RealTimeDecisionDigitalWalletAuthenticationDigitalWallet string

const (
	// Apple Pay
	RealTimeDecisionDigitalWalletAuthenticationDigitalWalletApplePay RealTimeDecisionDigitalWalletAuthenticationDigitalWallet = "apple_pay"
	// Google Pay
	RealTimeDecisionDigitalWalletAuthenticationDigitalWalletGooglePay RealTimeDecisionDigitalWalletAuthenticationDigitalWallet = "google_pay"
)

// Whether your application successfully delivered the one-time passcode.
type RealTimeDecisionDigitalWalletAuthenticationResult string

const (
	// Your application successfully delivered the one-time passcode to the cardholder.
	RealTimeDecisionDigitalWalletAuthenticationResultSuccess RealTimeDecisionDigitalWalletAuthenticationResult = "success"
	// Your application failed to deliver the one-time passcode to the cardholder.
	RealTimeDecisionDigitalWalletAuthenticationResultFailure RealTimeDecisionDigitalWalletAuthenticationResult = "failure"
)

// Fields related to a digital wallet token provisioning attempt.
type RealTimeDecisionDigitalWalletToken struct {
	// The identifier of the Card that is being tokenized.
	CardID string `json:"card_id,required"`
	// The identifier of the Card Profile that was set via the real time decision. This
	// will be null until the real time decision is responded to or if the real time
	// decision did not set a card profile.
	CardProfileID string `json:"card_profile_id,required,nullable"`
	// Whether or not the provisioning request was approved. This will be null until
	// the real time decision is responded to.
	Decision RealTimeDecisionDigitalWalletTokenDecision `json:"decision,required,nullable"`
	// The digital wallet app being used.
	DigitalWallet RealTimeDecisionDigitalWalletTokenDigitalWallet `json:"digital_wallet,required"`
	JSON          realTimeDecisionDigitalWalletTokenJSON
}

// realTimeDecisionDigitalWalletTokenJSON contains the JSON metadata for the struct
// [RealTimeDecisionDigitalWalletToken]
type realTimeDecisionDigitalWalletTokenJSON struct {
	CardID        apijson.Field
	CardProfileID apijson.Field
	Decision      apijson.Field
	DigitalWallet apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RealTimeDecisionDigitalWalletToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not the provisioning request was approved. This will be null until
// the real time decision is responded to.
type RealTimeDecisionDigitalWalletTokenDecision string

const (
	// Approve the provisioning request.
	RealTimeDecisionDigitalWalletTokenDecisionApprove RealTimeDecisionDigitalWalletTokenDecision = "approve"
	// Decline the provisioning request.
	RealTimeDecisionDigitalWalletTokenDecisionDecline RealTimeDecisionDigitalWalletTokenDecision = "decline"
)

// The digital wallet app being used.
type RealTimeDecisionDigitalWalletTokenDigitalWallet string

const (
	// Apple Pay
	RealTimeDecisionDigitalWalletTokenDigitalWalletApplePay RealTimeDecisionDigitalWalletTokenDigitalWallet = "apple_pay"
	// Google Pay
	RealTimeDecisionDigitalWalletTokenDigitalWalletGooglePay RealTimeDecisionDigitalWalletTokenDigitalWallet = "google_pay"
)

// The status of the Real-Time Decision.
type RealTimeDecisionStatus string

const (
	// The decision is pending action via real-time webhook.
	RealTimeDecisionStatusPending RealTimeDecisionStatus = "pending"
	// Your webhook actioned the real-time decision.
	RealTimeDecisionStatusResponded RealTimeDecisionStatus = "responded"
	// Your webhook failed to respond to the authorization in time.
	RealTimeDecisionStatusTimedOut RealTimeDecisionStatus = "timed_out"
)

// A constant representing the object's type. For this resource it will always be
// `real_time_decision`.
type RealTimeDecisionType string

const (
	RealTimeDecisionTypeRealTimeDecision RealTimeDecisionType = "real_time_decision"
)

type RealTimeDecisionActionParams struct {
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

// If the Real-Time Decision relates to a card authorization attempt, this object
// contains your response to the authorization.
type RealTimeDecisionActionParamsCardAuthorization struct {
	// Whether the card authorization should be approved or declined.
	Decision param.Field[RealTimeDecisionActionParamsCardAuthorizationDecision] `json:"decision,required"`
}

func (r RealTimeDecisionActionParamsCardAuthorization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether the card authorization should be approved or declined.
type RealTimeDecisionActionParamsCardAuthorizationDecision string

const (
	// Approve the authorization.
	RealTimeDecisionActionParamsCardAuthorizationDecisionApprove RealTimeDecisionActionParamsCardAuthorizationDecision = "approve"
	// Decline the authorization.
	RealTimeDecisionActionParamsCardAuthorizationDecisionDecline RealTimeDecisionActionParamsCardAuthorizationDecision = "decline"
)

// If the Real-Time Decision relates to a digital wallet authentication attempt,
// this object contains your response to the authentication.
type RealTimeDecisionActionParamsDigitalWalletAuthentication struct {
	// Whether your application was able to deliver the one-time passcode.
	Result param.Field[RealTimeDecisionActionParamsDigitalWalletAuthenticationResult] `json:"result,required"`
}

func (r RealTimeDecisionActionParamsDigitalWalletAuthentication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether your application was able to deliver the one-time passcode.
type RealTimeDecisionActionParamsDigitalWalletAuthenticationResult string

const (
	// Your application successfully delivered the one-time passcode to the cardholder.
	RealTimeDecisionActionParamsDigitalWalletAuthenticationResultSuccess RealTimeDecisionActionParamsDigitalWalletAuthenticationResult = "success"
	// Your application failed to deliver the one-time passcode to the cardholder.
	RealTimeDecisionActionParamsDigitalWalletAuthenticationResultFailure RealTimeDecisionActionParamsDigitalWalletAuthenticationResult = "failure"
)

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
	// The identifier of the Card Profile to assign to the Digital Wallet token.
	CardProfileID param.Field[string] `json:"card_profile_id,required"`
	// An email address that can be used to verify the cardholder via one-time
	// passcode.
	Email param.Field[string] `json:"email"`
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

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
func (r *RealTimeDecisionService) Get(ctx context.Context, real_time_decision_id string, opts ...option.RequestOption) (res *RealTimeDecision, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("real_time_decisions/%s", real_time_decision_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Action a Real-Time Decision
func (r *RealTimeDecisionService) Action(ctx context.Context, real_time_decision_id string, body RealTimeDecisionActionParams, opts ...option.RequestOption) (res *RealTimeDecision, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("real_time_decisions/%s/action", real_time_decision_id)
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
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Real-Time Decision was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// your application can no longer respond to the Real-Time Decision.
	TimeoutAt time.Time `json:"timeout_at,required" format:"date-time"`
	// The status of the Real-Time Decision.
	Status RealTimeDecisionStatus `json:"status,required"`
	// The category of the Real-Time Decision.
	Category RealTimeDecisionCategory `json:"category,required"`
	// Fields related to a card authorization.
	CardAuthorization RealTimeDecisionCardAuthorization `json:"card_authorization,required,nullable"`
	// Fields related to a digital wallet token provisioning attempt.
	DigitalWalletToken RealTimeDecisionDigitalWalletToken `json:"digital_wallet_token,required,nullable"`
	// Fields related to a digital wallet authentication attempt.
	DigitalWalletAuthentication RealTimeDecisionDigitalWalletAuthentication `json:"digital_wallet_authentication,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `real_time_decision`.
	Type RealTimeDecisionType `json:"type,required"`
	JSON realTimeDecisionJSON
}

// realTimeDecisionJSON contains the JSON metadata for the struct
// [RealTimeDecision]
type realTimeDecisionJSON struct {
	ID                          apijson.Field
	CreatedAt                   apijson.Field
	TimeoutAt                   apijson.Field
	Status                      apijson.Field
	Category                    apijson.Field
	CardAuthorization           apijson.Field
	DigitalWalletToken          apijson.Field
	DigitalWalletAuthentication apijson.Field
	Type                        apijson.Field
	raw                         string
	Extras                      map[string]apijson.Field
}

func (r *RealTimeDecision) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RealTimeDecisionStatus string

const (
	RealTimeDecisionStatusPending   RealTimeDecisionStatus = "pending"
	RealTimeDecisionStatusResponded RealTimeDecisionStatus = "responded"
	RealTimeDecisionStatusTimedOut  RealTimeDecisionStatus = "timed_out"
)

type RealTimeDecisionCategory string

const (
	RealTimeDecisionCategoryCardAuthorizationRequested           RealTimeDecisionCategory = "card_authorization_requested"
	RealTimeDecisionCategoryDigitalWalletTokenRequested          RealTimeDecisionCategory = "digital_wallet_token_requested"
	RealTimeDecisionCategoryDigitalWalletAuthenticationRequested RealTimeDecisionCategory = "digital_wallet_authentication_requested"
)

// Fields related to a card authorization.
type RealTimeDecisionCardAuthorization struct {
	// The merchant identifier (commonly abbreviated as MID) of the merchant the card
	// is transacting with.
	MerchantAcceptorID string `json:"merchant_acceptor_id,required"`
	// The merchant descriptor of the merchant the card is transacting with.
	MerchantDescriptor string `json:"merchant_descriptor,required"`
	// The Merchant Category Code (commonly abbreviated as MCC) of the merchant the
	// card is transacting with.
	MerchantCategoryCode string `json:"merchant_category_code,required,nullable"`
	// The city the merchant resides in.
	MerchantCity string `json:"merchant_city,required,nullable"`
	// The country the merchant resides in.
	MerchantCountry string `json:"merchant_country,required,nullable"`
	// The payment network used to process this card authorization
	Network RealTimeDecisionCardAuthorizationNetwork `json:"network,required"`
	// Fields specific to the `network`
	NetworkDetails RealTimeDecisionCardAuthorizationNetworkDetails `json:"network_details,required"`
	// Whether or not the authorization was approved.
	Decision RealTimeDecisionCardAuthorizationDecision `json:"decision,required,nullable"`
	// The identifier of the Card that is being authorized.
	CardID string `json:"card_id,required"`
	// The identifier of the Account the authorization will debit.
	AccountID string `json:"account_id,required"`
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
	MerchantAcceptorID   apijson.Field
	MerchantDescriptor   apijson.Field
	MerchantCategoryCode apijson.Field
	MerchantCity         apijson.Field
	MerchantCountry      apijson.Field
	Network              apijson.Field
	NetworkDetails       apijson.Field
	Decision             apijson.Field
	CardID               apijson.Field
	AccountID            apijson.Field
	PresentmentAmount    apijson.Field
	PresentmentCurrency  apijson.Field
	SettlementAmount     apijson.Field
	SettlementCurrency   apijson.Field
	raw                  string
	Extras               map[string]apijson.Field
}

func (r *RealTimeDecisionCardAuthorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RealTimeDecisionCardAuthorizationNetwork string

const (
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
	Visa   apijson.Field
	raw    string
	Extras map[string]apijson.Field
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
	Extras                      map[string]apijson.Field
}

func (r *RealTimeDecisionCardAuthorizationNetworkDetailsVisa) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

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

type RealTimeDecisionCardAuthorizationDecision string

const (
	RealTimeDecisionCardAuthorizationDecisionApprove RealTimeDecisionCardAuthorizationDecision = "approve"
	RealTimeDecisionCardAuthorizationDecisionDecline RealTimeDecisionCardAuthorizationDecision = "decline"
)

// Fields related to a digital wallet token provisioning attempt.
type RealTimeDecisionDigitalWalletToken struct {
	// Whether or not the provisioning request was approved. This will be null until
	// the real time decision is responded to.
	Decision RealTimeDecisionDigitalWalletTokenDecision `json:"decision,required,nullable"`
	// The identifier of the Card that is being tokenized.
	CardID string `json:"card_id,required"`
	// The digital wallet app being used.
	DigitalWallet RealTimeDecisionDigitalWalletTokenDigitalWallet `json:"digital_wallet,required"`
	// The identifier of the Card Profile that was set via the real time decision. This
	// will be null until the real time decision is responded to or if the real time
	// decision did not set a card profile.
	CardProfileID string `json:"card_profile_id,required,nullable"`
	JSON          realTimeDecisionDigitalWalletTokenJSON
}

// realTimeDecisionDigitalWalletTokenJSON contains the JSON metadata for the struct
// [RealTimeDecisionDigitalWalletToken]
type realTimeDecisionDigitalWalletTokenJSON struct {
	Decision      apijson.Field
	CardID        apijson.Field
	DigitalWallet apijson.Field
	CardProfileID apijson.Field
	raw           string
	Extras        map[string]apijson.Field
}

func (r *RealTimeDecisionDigitalWalletToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RealTimeDecisionDigitalWalletTokenDecision string

const (
	RealTimeDecisionDigitalWalletTokenDecisionApprove RealTimeDecisionDigitalWalletTokenDecision = "approve"
	RealTimeDecisionDigitalWalletTokenDecisionDecline RealTimeDecisionDigitalWalletTokenDecision = "decline"
)

type RealTimeDecisionDigitalWalletTokenDigitalWallet string

const (
	RealTimeDecisionDigitalWalletTokenDigitalWalletApplePay  RealTimeDecisionDigitalWalletTokenDigitalWallet = "apple_pay"
	RealTimeDecisionDigitalWalletTokenDigitalWalletGooglePay RealTimeDecisionDigitalWalletTokenDigitalWallet = "google_pay"
)

// Fields related to a digital wallet authentication attempt.
type RealTimeDecisionDigitalWalletAuthentication struct {
	// Whether your application successfully delivered the one-time passcode.
	Result RealTimeDecisionDigitalWalletAuthenticationResult `json:"result,required,nullable"`
	// The identifier of the Card that is being tokenized.
	CardID string `json:"card_id,required"`
	// The digital wallet app being used.
	DigitalWallet RealTimeDecisionDigitalWalletAuthenticationDigitalWallet `json:"digital_wallet,required"`
	// The channel to send the card user their one-time passcode.
	Channel RealTimeDecisionDigitalWalletAuthenticationChannel `json:"channel,required"`
	// The one-time passcode to send the card user.
	OneTimePasscode string `json:"one_time_passcode,required"`
	// The phone number to send the one-time passcode to if `channel` is equal to
	// `sms`.
	Phone string `json:"phone,required,nullable"`
	// The email to send the one-time passcode to if `channel` is equal to `email`.
	Email string `json:"email,required,nullable"`
	JSON  realTimeDecisionDigitalWalletAuthenticationJSON
}

// realTimeDecisionDigitalWalletAuthenticationJSON contains the JSON metadata for
// the struct [RealTimeDecisionDigitalWalletAuthentication]
type realTimeDecisionDigitalWalletAuthenticationJSON struct {
	Result          apijson.Field
	CardID          apijson.Field
	DigitalWallet   apijson.Field
	Channel         apijson.Field
	OneTimePasscode apijson.Field
	Phone           apijson.Field
	Email           apijson.Field
	raw             string
	Extras          map[string]apijson.Field
}

func (r *RealTimeDecisionDigitalWalletAuthentication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RealTimeDecisionDigitalWalletAuthenticationResult string

const (
	RealTimeDecisionDigitalWalletAuthenticationResultSuccess RealTimeDecisionDigitalWalletAuthenticationResult = "success"
	RealTimeDecisionDigitalWalletAuthenticationResultFailure RealTimeDecisionDigitalWalletAuthenticationResult = "failure"
)

type RealTimeDecisionDigitalWalletAuthenticationDigitalWallet string

const (
	RealTimeDecisionDigitalWalletAuthenticationDigitalWalletApplePay  RealTimeDecisionDigitalWalletAuthenticationDigitalWallet = "apple_pay"
	RealTimeDecisionDigitalWalletAuthenticationDigitalWalletGooglePay RealTimeDecisionDigitalWalletAuthenticationDigitalWallet = "google_pay"
)

type RealTimeDecisionDigitalWalletAuthenticationChannel string

const (
	RealTimeDecisionDigitalWalletAuthenticationChannelSMS   RealTimeDecisionDigitalWalletAuthenticationChannel = "sms"
	RealTimeDecisionDigitalWalletAuthenticationChannelEmail RealTimeDecisionDigitalWalletAuthenticationChannel = "email"
)

type RealTimeDecisionType string

const (
	RealTimeDecisionTypeRealTimeDecision RealTimeDecisionType = "real_time_decision"
)

type RealTimeDecisionActionParams struct {
	// If the Real-Time Decision relates to a card authorization attempt, this object
	// contains your response to the authorization.
	CardAuthorization param.Field[RealTimeDecisionActionParamsCardAuthorization] `json:"card_authorization"`
	// If the Real-Time Decision relates to a digital wallet token provisioning
	// attempt, this object contains your response to the attempt.
	DigitalWalletToken param.Field[RealTimeDecisionActionParamsDigitalWalletToken] `json:"digital_wallet_token"`
	// If the Real-Time Decision relates to a digital wallet authentication attempt,
	// this object contains your response to the authentication.
	DigitalWalletAuthentication param.Field[RealTimeDecisionActionParamsDigitalWalletAuthentication] `json:"digital_wallet_authentication"`
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

type RealTimeDecisionActionParamsCardAuthorizationDecision string

const (
	RealTimeDecisionActionParamsCardAuthorizationDecisionApprove RealTimeDecisionActionParamsCardAuthorizationDecision = "approve"
	RealTimeDecisionActionParamsCardAuthorizationDecisionDecline RealTimeDecisionActionParamsCardAuthorizationDecision = "decline"
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

// If your application approves the provisioning attempt, this contains metadata
// about the digital wallet token that will be generated.
type RealTimeDecisionActionParamsDigitalWalletTokenApproval struct {
	// The identifier of the Card Profile to assign to the Digital Wallet token.
	CardProfileID param.Field[string] `json:"card_profile_id,required"`
	// A phone number that can be used to verify the cardholder via one-time passcode
	// over SMS.
	Phone param.Field[string] `json:"phone"`
	// An email address that can be used to verify the cardholder via one-time
	// passcode.
	Email param.Field[string] `json:"email"`
}

// If your application declines the provisioning attempt, this contains details
// about the decline.
type RealTimeDecisionActionParamsDigitalWalletTokenDecline struct {
	// Why the tokenization attempt was declined. This is for logging purposes only and
	// is not displayed to the end-user.
	Reason param.Field[string] `json:"reason"`
}

// If the Real-Time Decision relates to a digital wallet authentication attempt,
// this object contains your response to the authentication.
type RealTimeDecisionActionParamsDigitalWalletAuthentication struct {
	// Whether your application was able to deliver the one-time passcode.
	Result param.Field[RealTimeDecisionActionParamsDigitalWalletAuthenticationResult] `json:"result,required"`
}

type RealTimeDecisionActionParamsDigitalWalletAuthenticationResult string

const (
	RealTimeDecisionActionParamsDigitalWalletAuthenticationResultSuccess RealTimeDecisionActionParamsDigitalWalletAuthenticationResult = "success"
	RealTimeDecisionActionParamsDigitalWalletAuthenticationResultFailure RealTimeDecisionActionParamsDigitalWalletAuthenticationResult = "failure"
)

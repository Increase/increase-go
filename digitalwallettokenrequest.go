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
	"github.com/Increase/increase-go/packages/pagination"
)

// DigitalWalletTokenRequestService contains methods and other services that help
// with interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDigitalWalletTokenRequestService] method instead.
type DigitalWalletTokenRequestService struct {
	Options []option.RequestOption
}

// NewDigitalWalletTokenRequestService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDigitalWalletTokenRequestService(opts ...option.RequestOption) (r *DigitalWalletTokenRequestService) {
	r = &DigitalWalletTokenRequestService{}
	r.Options = opts
	return
}

// Retrieve a Digital Wallet Token Request
func (r *DigitalWalletTokenRequestService) Get(ctx context.Context, digitalWalletTokenRequestID string, opts ...option.RequestOption) (res *DigitalWalletTokenRequest, err error) {
	opts = slices.Concat(r.Options, opts)
	if digitalWalletTokenRequestID == "" {
		err = errors.New("missing required digital_wallet_token_request_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("digital_wallet_token_requests/%s", digitalWalletTokenRequestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List Digital Wallet Token Requests
func (r *DigitalWalletTokenRequestService) List(ctx context.Context, query DigitalWalletTokenRequestListParams, opts ...option.RequestOption) (res *pagination.Page[DigitalWalletTokenRequest], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "digital_wallet_token_requests"
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

// List Digital Wallet Token Requests
func (r *DigitalWalletTokenRequestService) ListAutoPaging(ctx context.Context, query DigitalWalletTokenRequestListParams, opts ...option.RequestOption) *pagination.PageAutoPager[DigitalWalletTokenRequest] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// A Digital Wallet Token Request is created each time a digital wallet app, such
// as Apple Pay or Google Pay, requests to tokenize a Card.
type DigitalWalletTokenRequest struct {
	// The Digital Wallet Token Request identifier.
	ID string `json:"id" api:"required"`
	// The identifier of the Card the tokenization was requested for.
	CardID string `json:"card_id" api:"required"`
	// The cardholder adding the card to the digital wallet.
	Cardholder DigitalWalletTokenRequestCardholder `json:"cardholder" api:"required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Digital Wallet Token Request was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Details of the decline. Present if and only if `outcome` is `declined`.
	Declined DigitalWalletTokenRequestDeclined `json:"declined" api:"required,nullable"`
	// The device that requested the tokenization.
	Device DigitalWalletTokenRequestDevice `json:"device" api:"required"`
	// The card network the token is being provisioned on.
	Network DigitalWalletTokenRequestNetwork `json:"network" api:"required"`
	// The outcome of the tokenization request.
	Outcome DigitalWalletTokenRequestOutcome `json:"outcome" api:"required"`
	// How the card number was provided to the digital wallet.
	PrimaryAccountNumberSource DigitalWalletTokenRequestPrimaryAccountNumberSource `json:"primary_account_number_source" api:"required,nullable"`
	// Details of the provisioned Digital Wallet Token. Present if and only if
	// `outcome` is `provisioned`.
	Provisioned DigitalWalletTokenRequestProvisioned `json:"provisioned" api:"required,nullable"`
	// The reference identifier assigned by the card network to the token.
	TokenReferenceIdentifier string `json:"token_reference_identifier" api:"required"`
	// The digital wallet app being used.
	TokenRequestor DigitalWalletTokenRequestTokenRequestor `json:"token_requestor" api:"required"`
	// The digital wallet or merchant that requested the token.
	TokenRequestorDetails DigitalWalletTokenRequestTokenRequestorDetails `json:"token_requestor_details" api:"required"`
	// A constant representing the object's type. For this resource it will always be
	// `digital_wallet_token_request`.
	Type DigitalWalletTokenRequestType `json:"type" api:"required"`
	// Fields related to verification of cardholder-provided values.
	Verification DigitalWalletTokenRequestVerification `json:"verification" api:"required"`
	// Risk scores provided by Visa for the provisioning request. This field will be
	// present in the JSON response if and only if `network` is equal to `visa`.
	Visa DigitalWalletTokenRequestVisa `json:"visa" api:"required,nullable"`
	// Details about the digital wallet.
	Wallet DigitalWalletTokenRequestWallet `json:"wallet" api:"required"`
	JSON   digitalWalletTokenRequestJSON   `json:"-"`
}

// digitalWalletTokenRequestJSON contains the JSON metadata for the struct
// [DigitalWalletTokenRequest]
type digitalWalletTokenRequestJSON struct {
	ID                         apijson.Field
	CardID                     apijson.Field
	Cardholder                 apijson.Field
	CreatedAt                  apijson.Field
	Declined                   apijson.Field
	Device                     apijson.Field
	Network                    apijson.Field
	Outcome                    apijson.Field
	PrimaryAccountNumberSource apijson.Field
	Provisioned                apijson.Field
	TokenReferenceIdentifier   apijson.Field
	TokenRequestor             apijson.Field
	TokenRequestorDetails      apijson.Field
	Type                       apijson.Field
	Verification               apijson.Field
	Visa                       apijson.Field
	Wallet                     apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *DigitalWalletTokenRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalWalletTokenRequestJSON) RawJSON() string {
	return r.raw
}

// The cardholder adding the card to the digital wallet.
type DigitalWalletTokenRequestCardholder struct {
	// The cardholder's name, as entered in the digital wallet.
	Name string                                  `json:"name" api:"required,nullable"`
	JSON digitalWalletTokenRequestCardholderJSON `json:"-"`
}

// digitalWalletTokenRequestCardholderJSON contains the JSON metadata for the
// struct [DigitalWalletTokenRequestCardholder]
type digitalWalletTokenRequestCardholderJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalWalletTokenRequestCardholder) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalWalletTokenRequestCardholderJSON) RawJSON() string {
	return r.raw
}

// Details of the decline. Present if and only if `outcome` is `declined`.
type DigitalWalletTokenRequestDeclined struct {
	// The reason the tokenization was declined.
	Reason DigitalWalletTokenRequestDeclinedReason `json:"reason" api:"required"`
	JSON   digitalWalletTokenRequestDeclinedJSON   `json:"-"`
}

// digitalWalletTokenRequestDeclinedJSON contains the JSON metadata for the struct
// [DigitalWalletTokenRequestDeclined]
type digitalWalletTokenRequestDeclinedJSON struct {
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalWalletTokenRequestDeclined) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalWalletTokenRequestDeclinedJSON) RawJSON() string {
	return r.raw
}

// The reason the tokenization was declined.
type DigitalWalletTokenRequestDeclinedReason string

const (
	DigitalWalletTokenRequestDeclinedReasonCardNotActive                 DigitalWalletTokenRequestDeclinedReason = "card_not_active"
	DigitalWalletTokenRequestDeclinedReasonNoVerificationMethod          DigitalWalletTokenRequestDeclinedReason = "no_verification_method"
	DigitalWalletTokenRequestDeclinedReasonWebhookTimedOut               DigitalWalletTokenRequestDeclinedReason = "webhook_timed_out"
	DigitalWalletTokenRequestDeclinedReasonWebhookDeclined               DigitalWalletTokenRequestDeclinedReason = "webhook_declined"
	DigitalWalletTokenRequestDeclinedReasonIncorrectCardVerificationCode DigitalWalletTokenRequestDeclinedReason = "incorrect_card_verification_code"
	DigitalWalletTokenRequestDeclinedReasonDeclinedByTokenRequestor      DigitalWalletTokenRequestDeclinedReason = "declined_by_token_requestor"
	DigitalWalletTokenRequestDeclinedReasonGroupLocked                   DigitalWalletTokenRequestDeclinedReason = "group_locked"
	DigitalWalletTokenRequestDeclinedReasonAccountClosed                 DigitalWalletTokenRequestDeclinedReason = "account_closed"
	DigitalWalletTokenRequestDeclinedReasonEntityNotActive               DigitalWalletTokenRequestDeclinedReason = "entity_not_active"
	DigitalWalletTokenRequestDeclinedReasonCardExpirationMismatch        DigitalWalletTokenRequestDeclinedReason = "card_expiration_mismatch"
)

func (r DigitalWalletTokenRequestDeclinedReason) IsKnown() bool {
	switch r {
	case DigitalWalletTokenRequestDeclinedReasonCardNotActive, DigitalWalletTokenRequestDeclinedReasonNoVerificationMethod, DigitalWalletTokenRequestDeclinedReasonWebhookTimedOut, DigitalWalletTokenRequestDeclinedReasonWebhookDeclined, DigitalWalletTokenRequestDeclinedReasonIncorrectCardVerificationCode, DigitalWalletTokenRequestDeclinedReasonDeclinedByTokenRequestor, DigitalWalletTokenRequestDeclinedReasonGroupLocked, DigitalWalletTokenRequestDeclinedReasonAccountClosed, DigitalWalletTokenRequestDeclinedReasonEntityNotActive, DigitalWalletTokenRequestDeclinedReasonCardExpirationMismatch:
		return true
	}
	return false
}

// The device that requested the tokenization.
type DigitalWalletTokenRequestDevice struct {
	// Device type.
	DeviceType DigitalWalletTokenRequestDeviceDeviceType `json:"device_type" api:"required,nullable"`
	// ID assigned to the device by the digital wallet provider.
	Identifier string `json:"identifier" api:"required,nullable"`
	// IP address of the device.
	IPAddress string `json:"ip_address" api:"required,nullable"`
	// Name of the device, for example "My Work Phone".
	Name string                              `json:"name" api:"required,nullable"`
	JSON digitalWalletTokenRequestDeviceJSON `json:"-"`
}

// digitalWalletTokenRequestDeviceJSON contains the JSON metadata for the struct
// [DigitalWalletTokenRequestDevice]
type digitalWalletTokenRequestDeviceJSON struct {
	DeviceType  apijson.Field
	Identifier  apijson.Field
	IPAddress   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalWalletTokenRequestDevice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalWalletTokenRequestDeviceJSON) RawJSON() string {
	return r.raw
}

// Device type.
type DigitalWalletTokenRequestDeviceDeviceType string

const (
	DigitalWalletTokenRequestDeviceDeviceTypeUnknown             DigitalWalletTokenRequestDeviceDeviceType = "unknown"
	DigitalWalletTokenRequestDeviceDeviceTypeMobilePhone         DigitalWalletTokenRequestDeviceDeviceType = "mobile_phone"
	DigitalWalletTokenRequestDeviceDeviceTypeTablet              DigitalWalletTokenRequestDeviceDeviceType = "tablet"
	DigitalWalletTokenRequestDeviceDeviceTypeWatch               DigitalWalletTokenRequestDeviceDeviceType = "watch"
	DigitalWalletTokenRequestDeviceDeviceTypeMobilephoneOrTablet DigitalWalletTokenRequestDeviceDeviceType = "mobilephone_or_tablet"
	DigitalWalletTokenRequestDeviceDeviceTypePc                  DigitalWalletTokenRequestDeviceDeviceType = "pc"
	DigitalWalletTokenRequestDeviceDeviceTypeHouseholdDevice     DigitalWalletTokenRequestDeviceDeviceType = "household_device"
	DigitalWalletTokenRequestDeviceDeviceTypeWearableDevice      DigitalWalletTokenRequestDeviceDeviceType = "wearable_device"
	DigitalWalletTokenRequestDeviceDeviceTypeAutomobileDevice    DigitalWalletTokenRequestDeviceDeviceType = "automobile_device"
)

func (r DigitalWalletTokenRequestDeviceDeviceType) IsKnown() bool {
	switch r {
	case DigitalWalletTokenRequestDeviceDeviceTypeUnknown, DigitalWalletTokenRequestDeviceDeviceTypeMobilePhone, DigitalWalletTokenRequestDeviceDeviceTypeTablet, DigitalWalletTokenRequestDeviceDeviceTypeWatch, DigitalWalletTokenRequestDeviceDeviceTypeMobilephoneOrTablet, DigitalWalletTokenRequestDeviceDeviceTypePc, DigitalWalletTokenRequestDeviceDeviceTypeHouseholdDevice, DigitalWalletTokenRequestDeviceDeviceTypeWearableDevice, DigitalWalletTokenRequestDeviceDeviceTypeAutomobileDevice:
		return true
	}
	return false
}

// The card network the token is being provisioned on.
type DigitalWalletTokenRequestNetwork string

const (
	DigitalWalletTokenRequestNetworkVisa DigitalWalletTokenRequestNetwork = "visa"
)

func (r DigitalWalletTokenRequestNetwork) IsKnown() bool {
	switch r {
	case DigitalWalletTokenRequestNetworkVisa:
		return true
	}
	return false
}

// The outcome of the tokenization request.
type DigitalWalletTokenRequestOutcome string

const (
	DigitalWalletTokenRequestOutcomeProvisioned DigitalWalletTokenRequestOutcome = "provisioned"
	DigitalWalletTokenRequestOutcomeDeclined    DigitalWalletTokenRequestOutcome = "declined"
)

func (r DigitalWalletTokenRequestOutcome) IsKnown() bool {
	switch r {
	case DigitalWalletTokenRequestOutcomeProvisioned, DigitalWalletTokenRequestOutcomeDeclined:
		return true
	}
	return false
}

// How the card number was provided to the digital wallet.
type DigitalWalletTokenRequestPrimaryAccountNumberSource string

const (
	DigitalWalletTokenRequestPrimaryAccountNumberSourceKeyEntered     DigitalWalletTokenRequestPrimaryAccountNumberSource = "key_entered"
	DigitalWalletTokenRequestPrimaryAccountNumberSourceOnFile         DigitalWalletTokenRequestPrimaryAccountNumberSource = "on_file"
	DigitalWalletTokenRequestPrimaryAccountNumberSourceMobileApp      DigitalWalletTokenRequestPrimaryAccountNumberSource = "mobile_app"
	DigitalWalletTokenRequestPrimaryAccountNumberSourceToken          DigitalWalletTokenRequestPrimaryAccountNumberSource = "token"
	DigitalWalletTokenRequestPrimaryAccountNumberSourceChipDip        DigitalWalletTokenRequestPrimaryAccountNumberSource = "chip_dip"
	DigitalWalletTokenRequestPrimaryAccountNumberSourceContactlessTap DigitalWalletTokenRequestPrimaryAccountNumberSource = "contactless_tap"
)

func (r DigitalWalletTokenRequestPrimaryAccountNumberSource) IsKnown() bool {
	switch r {
	case DigitalWalletTokenRequestPrimaryAccountNumberSourceKeyEntered, DigitalWalletTokenRequestPrimaryAccountNumberSourceOnFile, DigitalWalletTokenRequestPrimaryAccountNumberSourceMobileApp, DigitalWalletTokenRequestPrimaryAccountNumberSourceToken, DigitalWalletTokenRequestPrimaryAccountNumberSourceChipDip, DigitalWalletTokenRequestPrimaryAccountNumberSourceContactlessTap:
		return true
	}
	return false
}

// Details of the provisioned Digital Wallet Token. Present if and only if
// `outcome` is `provisioned`.
type DigitalWalletTokenRequestProvisioned struct {
	// The identifier of the Digital Wallet Token that was provisioned.
	DigitalWalletTokenID string                                   `json:"digital_wallet_token_id" api:"required"`
	JSON                 digitalWalletTokenRequestProvisionedJSON `json:"-"`
}

// digitalWalletTokenRequestProvisionedJSON contains the JSON metadata for the
// struct [DigitalWalletTokenRequestProvisioned]
type digitalWalletTokenRequestProvisionedJSON struct {
	DigitalWalletTokenID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *DigitalWalletTokenRequestProvisioned) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalWalletTokenRequestProvisionedJSON) RawJSON() string {
	return r.raw
}

// The digital wallet app being used.
type DigitalWalletTokenRequestTokenRequestor string

const (
	DigitalWalletTokenRequestTokenRequestorApplePay   DigitalWalletTokenRequestTokenRequestor = "apple_pay"
	DigitalWalletTokenRequestTokenRequestorGooglePay  DigitalWalletTokenRequestTokenRequestor = "google_pay"
	DigitalWalletTokenRequestTokenRequestorSamsungPay DigitalWalletTokenRequestTokenRequestor = "samsung_pay"
	DigitalWalletTokenRequestTokenRequestorGarminPay  DigitalWalletTokenRequestTokenRequestor = "garmin_pay"
	DigitalWalletTokenRequestTokenRequestorUnknown    DigitalWalletTokenRequestTokenRequestor = "unknown"
)

func (r DigitalWalletTokenRequestTokenRequestor) IsKnown() bool {
	switch r {
	case DigitalWalletTokenRequestTokenRequestorApplePay, DigitalWalletTokenRequestTokenRequestorGooglePay, DigitalWalletTokenRequestTokenRequestorSamsungPay, DigitalWalletTokenRequestTokenRequestorGarminPay, DigitalWalletTokenRequestTokenRequestorUnknown:
		return true
	}
	return false
}

// The digital wallet or merchant that requested the token.
type DigitalWalletTokenRequestTokenRequestorDetails struct {
	// The identifier assigned to the token requestor by the card network.
	Identifier string `json:"identifier" api:"required"`
	// The name of the token requestor, for example "Apple Inc.".
	Name string                                             `json:"name" api:"required,nullable"`
	JSON digitalWalletTokenRequestTokenRequestorDetailsJSON `json:"-"`
}

// digitalWalletTokenRequestTokenRequestorDetailsJSON contains the JSON metadata
// for the struct [DigitalWalletTokenRequestTokenRequestorDetails]
type digitalWalletTokenRequestTokenRequestorDetailsJSON struct {
	Identifier  apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalWalletTokenRequestTokenRequestorDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalWalletTokenRequestTokenRequestorDetailsJSON) RawJSON() string {
	return r.raw
}

// A constant representing the object's type. For this resource it will always be
// `digital_wallet_token_request`.
type DigitalWalletTokenRequestType string

const (
	DigitalWalletTokenRequestTypeDigitalWalletTokenRequest DigitalWalletTokenRequestType = "digital_wallet_token_request"
)

func (r DigitalWalletTokenRequestType) IsKnown() bool {
	switch r {
	case DigitalWalletTokenRequestTypeDigitalWalletTokenRequest:
		return true
	}
	return false
}

// Fields related to verification of cardholder-provided values.
type DigitalWalletTokenRequestVerification struct {
	// Cardholder address provided in the provisioning request and the address on file
	// we verified it against.
	CardholderAddress DigitalWalletTokenRequestVerificationCardholderAddress `json:"cardholder_address" api:"required"`
	JSON              digitalWalletTokenRequestVerificationJSON              `json:"-"`
}

// digitalWalletTokenRequestVerificationJSON contains the JSON metadata for the
// struct [DigitalWalletTokenRequestVerification]
type digitalWalletTokenRequestVerificationJSON struct {
	CardholderAddress apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DigitalWalletTokenRequestVerification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalWalletTokenRequestVerificationJSON) RawJSON() string {
	return r.raw
}

// Cardholder address provided in the provisioning request and the address on file
// we verified it against.
type DigitalWalletTokenRequestVerificationCardholderAddress struct {
	// Line 1 of the address on file for the cardholder.
	ActualLine1 string `json:"actual_line1" api:"required,nullable"`
	// The postal code of the address on file for the cardholder.
	ActualPostalCode string `json:"actual_postal_code" api:"required,nullable"`
	// The cardholder address line 1 provided for verification in the provisioning
	// request.
	ProvidedLine1 string `json:"provided_line1" api:"required,nullable"`
	// The postal code provided for verification in the provisioning request.
	ProvidedPostalCode string `json:"provided_postal_code" api:"required,nullable"`
	// The address verification result.
	Result DigitalWalletTokenRequestVerificationCardholderAddressResult `json:"result" api:"required"`
	JSON   digitalWalletTokenRequestVerificationCardholderAddressJSON   `json:"-"`
}

// digitalWalletTokenRequestVerificationCardholderAddressJSON contains the JSON
// metadata for the struct [DigitalWalletTokenRequestVerificationCardholderAddress]
type digitalWalletTokenRequestVerificationCardholderAddressJSON struct {
	ActualLine1        apijson.Field
	ActualPostalCode   apijson.Field
	ProvidedLine1      apijson.Field
	ProvidedPostalCode apijson.Field
	Result             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DigitalWalletTokenRequestVerificationCardholderAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalWalletTokenRequestVerificationCardholderAddressJSON) RawJSON() string {
	return r.raw
}

// The address verification result.
type DigitalWalletTokenRequestVerificationCardholderAddressResult string

const (
	DigitalWalletTokenRequestVerificationCardholderAddressResultNotChecked                       DigitalWalletTokenRequestVerificationCardholderAddressResult = "not_checked"
	DigitalWalletTokenRequestVerificationCardholderAddressResultPostalCodeMatchAddressNoMatch    DigitalWalletTokenRequestVerificationCardholderAddressResult = "postal_code_match_address_no_match"
	DigitalWalletTokenRequestVerificationCardholderAddressResultPostalCodeNoMatchAddressMatch    DigitalWalletTokenRequestVerificationCardholderAddressResult = "postal_code_no_match_address_match"
	DigitalWalletTokenRequestVerificationCardholderAddressResultMatch                            DigitalWalletTokenRequestVerificationCardholderAddressResult = "match"
	DigitalWalletTokenRequestVerificationCardholderAddressResultNoMatch                          DigitalWalletTokenRequestVerificationCardholderAddressResult = "no_match"
	DigitalWalletTokenRequestVerificationCardholderAddressResultPostalCodeMatchAddressNotChecked DigitalWalletTokenRequestVerificationCardholderAddressResult = "postal_code_match_address_not_checked"
)

func (r DigitalWalletTokenRequestVerificationCardholderAddressResult) IsKnown() bool {
	switch r {
	case DigitalWalletTokenRequestVerificationCardholderAddressResultNotChecked, DigitalWalletTokenRequestVerificationCardholderAddressResultPostalCodeMatchAddressNoMatch, DigitalWalletTokenRequestVerificationCardholderAddressResultPostalCodeNoMatchAddressMatch, DigitalWalletTokenRequestVerificationCardholderAddressResultMatch, DigitalWalletTokenRequestVerificationCardholderAddressResultNoMatch, DigitalWalletTokenRequestVerificationCardholderAddressResultPostalCodeMatchAddressNotChecked:
		return true
	}
	return false
}

// Risk scores provided by Visa for the provisioning request. This field will be
// present in the JSON response if and only if `network` is equal to `visa`.
type DigitalWalletTokenRequestVisa struct {
	// Visa's risk score for the card, based on its transaction history, from 0 to 99.
	// A higher score indicates a higher risk.
	NetworkRiskScore int64 `json:"network_risk_score" api:"required,nullable"`
	// Visa's risk score for the provisioning request, from 1 to 99. A higher score
	// indicates a higher risk.
	ProvisioningRiskScore int64                             `json:"provisioning_risk_score" api:"required,nullable"`
	JSON                  digitalWalletTokenRequestVisaJSON `json:"-"`
}

// digitalWalletTokenRequestVisaJSON contains the JSON metadata for the struct
// [DigitalWalletTokenRequestVisa]
type digitalWalletTokenRequestVisaJSON struct {
	NetworkRiskScore      apijson.Field
	ProvisioningRiskScore apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *DigitalWalletTokenRequestVisa) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalWalletTokenRequestVisaJSON) RawJSON() string {
	return r.raw
}

// Details about the digital wallet.
type DigitalWalletTokenRequestWallet struct {
	// The number of days since the email address was added to the cardholder's digital
	// wallet account, up to 999.
	AccountEmailAgeDays int64 `json:"account_email_age_days" api:"required,nullable"`
	// An identifier for the cardholder's account with the digital wallet.
	AccountIdentifier string `json:"account_identifier" api:"required,nullable"`
	// How much the digital wallet trusts the cardholder's account, from 1 to 5. A
	// higher score indicates more trust.
	AccountScore int64 `json:"account_score" api:"required,nullable"`
	// How much the digital wallet trusts the device the card is being added to, from 1
	// to 5. A higher score indicates more trust.
	DeviceScore int64 `json:"device_score" api:"required,nullable"`
	// A hash of the email address on the cardholder's digital wallet account.
	EmailHash string                              `json:"email_hash" api:"required,nullable"`
	JSON      digitalWalletTokenRequestWalletJSON `json:"-"`
}

// digitalWalletTokenRequestWalletJSON contains the JSON metadata for the struct
// [DigitalWalletTokenRequestWallet]
type digitalWalletTokenRequestWalletJSON struct {
	AccountEmailAgeDays apijson.Field
	AccountIdentifier   apijson.Field
	AccountScore        apijson.Field
	DeviceScore         apijson.Field
	EmailHash           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *DigitalWalletTokenRequestWallet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalWalletTokenRequestWalletJSON) RawJSON() string {
	return r.raw
}

type DigitalWalletTokenRequestListParams struct {
	// Filter Digital Wallet Token Requests to ones for the specified Card.
	CardID    param.Field[string]                                       `query:"card_id"`
	CreatedAt param.Field[DigitalWalletTokenRequestListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	//
	// Defaults to `100`.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [DigitalWalletTokenRequestListParams]'s query parameters as
// `url.Values`.
func (r DigitalWalletTokenRequestListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DigitalWalletTokenRequestListParamsCreatedAt struct {
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

// URLQuery serializes [DigitalWalletTokenRequestListParamsCreatedAt]'s query
// parameters as `url.Values`.
func (r DigitalWalletTokenRequestListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

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
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Digital Wallet Token Request was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Details of the decline. Present if and only if `outcome` is `declined`.
	Declined DigitalWalletTokenRequestDeclined `json:"declined" api:"required,nullable"`
	// The device that requested the tokenization.
	Device DigitalWalletTokenRequestDevice `json:"device" api:"required"`
	// The outcome of the tokenization request.
	Outcome DigitalWalletTokenRequestOutcome `json:"outcome" api:"required"`
	// Details of the provisioned Digital Wallet Token. Present if and only if
	// `outcome` is `provisioned`.
	Provisioned DigitalWalletTokenRequestProvisioned `json:"provisioned" api:"required,nullable"`
	// The reference identifier assigned by the card network to the token.
	TokenReferenceIdentifier string `json:"token_reference_identifier" api:"required"`
	// The digital wallet app being used.
	TokenRequestor DigitalWalletTokenRequestTokenRequestor `json:"token_requestor" api:"required"`
	// A constant representing the object's type. For this resource it will always be
	// `digital_wallet_token_request`.
	Type DigitalWalletTokenRequestType `json:"type" api:"required"`
	JSON digitalWalletTokenRequestJSON `json:"-"`
}

// digitalWalletTokenRequestJSON contains the JSON metadata for the struct
// [DigitalWalletTokenRequest]
type digitalWalletTokenRequestJSON struct {
	ID                       apijson.Field
	CardID                   apijson.Field
	CreatedAt                apijson.Field
	Declined                 apijson.Field
	Device                   apijson.Field
	Outcome                  apijson.Field
	Provisioned              apijson.Field
	TokenReferenceIdentifier apijson.Field
	TokenRequestor           apijson.Field
	Type                     apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *DigitalWalletTokenRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalWalletTokenRequestJSON) RawJSON() string {
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
)

func (r DigitalWalletTokenRequestDeclinedReason) IsKnown() bool {
	switch r {
	case DigitalWalletTokenRequestDeclinedReasonCardNotActive, DigitalWalletTokenRequestDeclinedReasonNoVerificationMethod, DigitalWalletTokenRequestDeclinedReasonWebhookTimedOut, DigitalWalletTokenRequestDeclinedReasonWebhookDeclined, DigitalWalletTokenRequestDeclinedReasonIncorrectCardVerificationCode, DigitalWalletTokenRequestDeclinedReasonDeclinedByTokenRequestor, DigitalWalletTokenRequestDeclinedReasonGroupLocked, DigitalWalletTokenRequestDeclinedReasonAccountClosed, DigitalWalletTokenRequestDeclinedReasonEntityNotActive:
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

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

// DigitalWalletTokenService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDigitalWalletTokenService] method instead.
type DigitalWalletTokenService struct {
	Options []option.RequestOption
}

// NewDigitalWalletTokenService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDigitalWalletTokenService(opts ...option.RequestOption) (r *DigitalWalletTokenService) {
	r = &DigitalWalletTokenService{}
	r.Options = opts
	return
}

// Retrieve a Digital Wallet Token
func (r *DigitalWalletTokenService) Get(ctx context.Context, digitalWalletTokenID string, opts ...option.RequestOption) (res *DigitalWalletToken, err error) {
	opts = slices.Concat(r.Options, opts)
	if digitalWalletTokenID == "" {
		err = errors.New("missing required digital_wallet_token_id parameter")
		return
	}
	path := fmt.Sprintf("digital_wallet_tokens/%s", digitalWalletTokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Digital Wallet Tokens
func (r *DigitalWalletTokenService) List(ctx context.Context, query DigitalWalletTokenListParams, opts ...option.RequestOption) (res *pagination.Page[DigitalWalletToken], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "digital_wallet_tokens"
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

// List Digital Wallet Tokens
func (r *DigitalWalletTokenService) ListAutoPaging(ctx context.Context, query DigitalWalletTokenListParams, opts ...option.RequestOption) *pagination.PageAutoPager[DigitalWalletToken] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// A Digital Wallet Token is created when a user adds a Card to their Apple Pay or
// Google Pay app. The Digital Wallet Token can be used for purchases just like a
// Card.
type DigitalWalletToken struct {
	// The Digital Wallet Token identifier.
	ID string `json:"id,required"`
	// The identifier for the Card this Digital Wallet Token belongs to.
	CardID string `json:"card_id,required"`
	// The cardholder information given when the Digital Wallet Token was created.
	Cardholder DigitalWalletTokenCardholder `json:"cardholder,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Digital Wallet Token was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The device that was used to create the Digital Wallet Token.
	Device DigitalWalletTokenDevice `json:"device,required"`
	// The redacted Dynamic Primary Account Number.
	DynamicPrimaryAccountNumber DigitalWalletTokenDynamicPrimaryAccountNumber `json:"dynamic_primary_account_number,required,nullable"`
	// This indicates if payments can be made with the Digital Wallet Token.
	Status DigitalWalletTokenStatus `json:"status,required"`
	// The digital wallet app being used.
	TokenRequestor DigitalWalletTokenTokenRequestor `json:"token_requestor,required"`
	// A constant representing the object's type. For this resource it will always be
	// `digital_wallet_token`.
	Type DigitalWalletTokenType `json:"type,required"`
	// Updates to the Digital Wallet Token.
	Updates []DigitalWalletTokenUpdate `json:"updates,required"`
	JSON    digitalWalletTokenJSON     `json:"-"`
}

// digitalWalletTokenJSON contains the JSON metadata for the struct
// [DigitalWalletToken]
type digitalWalletTokenJSON struct {
	ID                          apijson.Field
	CardID                      apijson.Field
	Cardholder                  apijson.Field
	CreatedAt                   apijson.Field
	Device                      apijson.Field
	DynamicPrimaryAccountNumber apijson.Field
	Status                      apijson.Field
	TokenRequestor              apijson.Field
	Type                        apijson.Field
	Updates                     apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *DigitalWalletToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalWalletTokenJSON) RawJSON() string {
	return r.raw
}

// The cardholder information given when the Digital Wallet Token was created.
type DigitalWalletTokenCardholder struct {
	// Name of the cardholder, for example "John Smith".
	Name string                           `json:"name,required,nullable"`
	JSON digitalWalletTokenCardholderJSON `json:"-"`
}

// digitalWalletTokenCardholderJSON contains the JSON metadata for the struct
// [DigitalWalletTokenCardholder]
type digitalWalletTokenCardholderJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalWalletTokenCardholder) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalWalletTokenCardholderJSON) RawJSON() string {
	return r.raw
}

// The device that was used to create the Digital Wallet Token.
type DigitalWalletTokenDevice struct {
	// Device type.
	DeviceType DigitalWalletTokenDeviceDeviceType `json:"device_type,required,nullable"`
	// ID assigned to the device by the digital wallet provider.
	Identifier string `json:"identifier,required,nullable"`
	// IP address of the device.
	IPAddress string `json:"ip_address,required,nullable"`
	// Name of the device, for example "My Work Phone".
	Name string                       `json:"name,required,nullable"`
	JSON digitalWalletTokenDeviceJSON `json:"-"`
}

// digitalWalletTokenDeviceJSON contains the JSON metadata for the struct
// [DigitalWalletTokenDevice]
type digitalWalletTokenDeviceJSON struct {
	DeviceType  apijson.Field
	Identifier  apijson.Field
	IPAddress   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalWalletTokenDevice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalWalletTokenDeviceJSON) RawJSON() string {
	return r.raw
}

// Device type.
type DigitalWalletTokenDeviceDeviceType string

const (
	DigitalWalletTokenDeviceDeviceTypeUnknown             DigitalWalletTokenDeviceDeviceType = "unknown"
	DigitalWalletTokenDeviceDeviceTypeMobilePhone         DigitalWalletTokenDeviceDeviceType = "mobile_phone"
	DigitalWalletTokenDeviceDeviceTypeTablet              DigitalWalletTokenDeviceDeviceType = "tablet"
	DigitalWalletTokenDeviceDeviceTypeWatch               DigitalWalletTokenDeviceDeviceType = "watch"
	DigitalWalletTokenDeviceDeviceTypeMobilephoneOrTablet DigitalWalletTokenDeviceDeviceType = "mobilephone_or_tablet"
	DigitalWalletTokenDeviceDeviceTypePc                  DigitalWalletTokenDeviceDeviceType = "pc"
	DigitalWalletTokenDeviceDeviceTypeHouseholdDevice     DigitalWalletTokenDeviceDeviceType = "household_device"
	DigitalWalletTokenDeviceDeviceTypeWearableDevice      DigitalWalletTokenDeviceDeviceType = "wearable_device"
	DigitalWalletTokenDeviceDeviceTypeAutomobileDevice    DigitalWalletTokenDeviceDeviceType = "automobile_device"
)

func (r DigitalWalletTokenDeviceDeviceType) IsKnown() bool {
	switch r {
	case DigitalWalletTokenDeviceDeviceTypeUnknown, DigitalWalletTokenDeviceDeviceTypeMobilePhone, DigitalWalletTokenDeviceDeviceTypeTablet, DigitalWalletTokenDeviceDeviceTypeWatch, DigitalWalletTokenDeviceDeviceTypeMobilephoneOrTablet, DigitalWalletTokenDeviceDeviceTypePc, DigitalWalletTokenDeviceDeviceTypeHouseholdDevice, DigitalWalletTokenDeviceDeviceTypeWearableDevice, DigitalWalletTokenDeviceDeviceTypeAutomobileDevice:
		return true
	}
	return false
}

// The redacted Dynamic Primary Account Number.
type DigitalWalletTokenDynamicPrimaryAccountNumber struct {
	// The first 6 digits of the token's Dynamic Primary Account Number.
	First6 string `json:"first6,required"`
	// The last 4 digits of the token's Dynamic Primary Account Number.
	Last4 string                                            `json:"last4,required"`
	JSON  digitalWalletTokenDynamicPrimaryAccountNumberJSON `json:"-"`
}

// digitalWalletTokenDynamicPrimaryAccountNumberJSON contains the JSON metadata for
// the struct [DigitalWalletTokenDynamicPrimaryAccountNumber]
type digitalWalletTokenDynamicPrimaryAccountNumberJSON struct {
	First6      apijson.Field
	Last4       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalWalletTokenDynamicPrimaryAccountNumber) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalWalletTokenDynamicPrimaryAccountNumberJSON) RawJSON() string {
	return r.raw
}

// This indicates if payments can be made with the Digital Wallet Token.
type DigitalWalletTokenStatus string

const (
	DigitalWalletTokenStatusActive      DigitalWalletTokenStatus = "active"
	DigitalWalletTokenStatusInactive    DigitalWalletTokenStatus = "inactive"
	DigitalWalletTokenStatusSuspended   DigitalWalletTokenStatus = "suspended"
	DigitalWalletTokenStatusDeactivated DigitalWalletTokenStatus = "deactivated"
)

func (r DigitalWalletTokenStatus) IsKnown() bool {
	switch r {
	case DigitalWalletTokenStatusActive, DigitalWalletTokenStatusInactive, DigitalWalletTokenStatusSuspended, DigitalWalletTokenStatusDeactivated:
		return true
	}
	return false
}

// The digital wallet app being used.
type DigitalWalletTokenTokenRequestor string

const (
	DigitalWalletTokenTokenRequestorApplePay   DigitalWalletTokenTokenRequestor = "apple_pay"
	DigitalWalletTokenTokenRequestorGooglePay  DigitalWalletTokenTokenRequestor = "google_pay"
	DigitalWalletTokenTokenRequestorSamsungPay DigitalWalletTokenTokenRequestor = "samsung_pay"
	DigitalWalletTokenTokenRequestorUnknown    DigitalWalletTokenTokenRequestor = "unknown"
)

func (r DigitalWalletTokenTokenRequestor) IsKnown() bool {
	switch r {
	case DigitalWalletTokenTokenRequestorApplePay, DigitalWalletTokenTokenRequestorGooglePay, DigitalWalletTokenTokenRequestorSamsungPay, DigitalWalletTokenTokenRequestorUnknown:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `digital_wallet_token`.
type DigitalWalletTokenType string

const (
	DigitalWalletTokenTypeDigitalWalletToken DigitalWalletTokenType = "digital_wallet_token"
)

func (r DigitalWalletTokenType) IsKnown() bool {
	switch r {
	case DigitalWalletTokenTypeDigitalWalletToken:
		return true
	}
	return false
}

type DigitalWalletTokenUpdate struct {
	// The status the update changed this Digital Wallet Token to.
	Status DigitalWalletTokenUpdatesStatus `json:"status,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the update happened.
	Timestamp time.Time                    `json:"timestamp,required" format:"date-time"`
	JSON      digitalWalletTokenUpdateJSON `json:"-"`
}

// digitalWalletTokenUpdateJSON contains the JSON metadata for the struct
// [DigitalWalletTokenUpdate]
type digitalWalletTokenUpdateJSON struct {
	Status      apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalWalletTokenUpdate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalWalletTokenUpdateJSON) RawJSON() string {
	return r.raw
}

// The status the update changed this Digital Wallet Token to.
type DigitalWalletTokenUpdatesStatus string

const (
	DigitalWalletTokenUpdatesStatusActive      DigitalWalletTokenUpdatesStatus = "active"
	DigitalWalletTokenUpdatesStatusInactive    DigitalWalletTokenUpdatesStatus = "inactive"
	DigitalWalletTokenUpdatesStatusSuspended   DigitalWalletTokenUpdatesStatus = "suspended"
	DigitalWalletTokenUpdatesStatusDeactivated DigitalWalletTokenUpdatesStatus = "deactivated"
)

func (r DigitalWalletTokenUpdatesStatus) IsKnown() bool {
	switch r {
	case DigitalWalletTokenUpdatesStatusActive, DigitalWalletTokenUpdatesStatusInactive, DigitalWalletTokenUpdatesStatusSuspended, DigitalWalletTokenUpdatesStatusDeactivated:
		return true
	}
	return false
}

type DigitalWalletTokenListParams struct {
	// Filter Digital Wallet Tokens to ones belonging to the specified Card.
	CardID    param.Field[string]                                `query:"card_id"`
	CreatedAt param.Field[DigitalWalletTokenListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [DigitalWalletTokenListParams]'s query parameters as
// `url.Values`.
func (r DigitalWalletTokenListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DigitalWalletTokenListParamsCreatedAt struct {
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

// URLQuery serializes [DigitalWalletTokenListParamsCreatedAt]'s query parameters
// as `url.Values`.
func (r DigitalWalletTokenListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

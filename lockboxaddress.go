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

// LockboxAddressService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLockboxAddressService] method instead.
type LockboxAddressService struct {
	Options []option.RequestOption
}

// NewLockboxAddressService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLockboxAddressService(opts ...option.RequestOption) (r *LockboxAddressService) {
	r = &LockboxAddressService{}
	r.Options = opts
	return
}

// Create a Lockbox Address
func (r *LockboxAddressService) New(ctx context.Context, body LockboxAddressNewParams, opts ...option.RequestOption) (res *LockboxAddress, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "lockbox_addresses"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve a Lockbox Address
func (r *LockboxAddressService) Get(ctx context.Context, lockboxAddressID string, opts ...option.RequestOption) (res *LockboxAddress, err error) {
	opts = slices.Concat(r.Options, opts)
	if lockboxAddressID == "" {
		err = errors.New("missing required lockbox_address_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("lockbox_addresses/%s", lockboxAddressID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update a Lockbox Address
func (r *LockboxAddressService) Update(ctx context.Context, lockboxAddressID string, body LockboxAddressUpdateParams, opts ...option.RequestOption) (res *LockboxAddress, err error) {
	opts = slices.Concat(r.Options, opts)
	if lockboxAddressID == "" {
		err = errors.New("missing required lockbox_address_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("lockbox_addresses/%s", lockboxAddressID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List Lockbox Addresses
func (r *LockboxAddressService) List(ctx context.Context, query LockboxAddressListParams, opts ...option.RequestOption) (res *pagination.Page[LockboxAddress], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "lockbox_addresses"
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

// List Lockbox Addresses
func (r *LockboxAddressService) ListAutoPaging(ctx context.Context, query LockboxAddressListParams, opts ...option.RequestOption) *pagination.PageAutoPager[LockboxAddress] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Lockbox Addresses are physical locations that can receive mail containing paper
// checks.
type LockboxAddress struct {
	// The Lockbox Address identifier.
	ID string `json:"id" api:"required"`
	// The mailing address for the Lockbox Address. It will be present after Increase
	// generates it.
	Address LockboxAddressAddress `json:"address" api:"required,nullable"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Lockbox
	// Address was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The description you choose for the Lockbox Address.
	Description string `json:"description" api:"required,nullable"`
	// The idempotency key you chose for this object. This value is unique across
	// Increase and is used to ensure that a request is only processed once. Learn more
	// about [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey string `json:"idempotency_key" api:"required,nullable"`
	// The status of the Lockbox Address.
	Status LockboxAddressStatus `json:"status" api:"required"`
	// A constant representing the object's type. For this resource it will always be
	// `lockbox_address`.
	Type LockboxAddressType `json:"type" api:"required"`
	JSON lockboxAddressJSON `json:"-"`
}

// lockboxAddressJSON contains the JSON metadata for the struct [LockboxAddress]
type lockboxAddressJSON struct {
	ID             apijson.Field
	Address        apijson.Field
	CreatedAt      apijson.Field
	Description    apijson.Field
	IdempotencyKey apijson.Field
	Status         apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *LockboxAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockboxAddressJSON) RawJSON() string {
	return r.raw
}

// The mailing address for the Lockbox Address. It will be present after Increase
// generates it.
type LockboxAddressAddress struct {
	// The city of the address.
	City string `json:"city" api:"required"`
	// The first line of the address.
	Line1 string `json:"line1" api:"required"`
	// The second line of the address.
	Line2 string `json:"line2" api:"required"`
	// The postal code of the address.
	PostalCode string `json:"postal_code" api:"required"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State string                    `json:"state" api:"required"`
	JSON  lockboxAddressAddressJSON `json:"-"`
}

// lockboxAddressAddressJSON contains the JSON metadata for the struct
// [LockboxAddressAddress]
type lockboxAddressAddressJSON struct {
	City        apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LockboxAddressAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockboxAddressAddressJSON) RawJSON() string {
	return r.raw
}

// The status of the Lockbox Address.
type LockboxAddressStatus string

const (
	LockboxAddressStatusPending  LockboxAddressStatus = "pending"
	LockboxAddressStatusActive   LockboxAddressStatus = "active"
	LockboxAddressStatusDisabled LockboxAddressStatus = "disabled"
	LockboxAddressStatusCanceled LockboxAddressStatus = "canceled"
)

func (r LockboxAddressStatus) IsKnown() bool {
	switch r {
	case LockboxAddressStatusPending, LockboxAddressStatusActive, LockboxAddressStatusDisabled, LockboxAddressStatusCanceled:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `lockbox_address`.
type LockboxAddressType string

const (
	LockboxAddressTypeLockboxAddress LockboxAddressType = "lockbox_address"
)

func (r LockboxAddressType) IsKnown() bool {
	switch r {
	case LockboxAddressTypeLockboxAddress:
		return true
	}
	return false
}

type LockboxAddressNewParams struct {
	// The description you choose for the Lockbox Address.
	Description param.Field[string] `json:"description"`
}

func (r LockboxAddressNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LockboxAddressUpdateParams struct {
	// The description you choose for the Lockbox Address.
	Description param.Field[string] `json:"description"`
	// The status of the Lockbox Address.
	Status param.Field[LockboxAddressUpdateParamsStatus] `json:"status"`
}

func (r LockboxAddressUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The status of the Lockbox Address.
type LockboxAddressUpdateParamsStatus string

const (
	LockboxAddressUpdateParamsStatusActive   LockboxAddressUpdateParamsStatus = "active"
	LockboxAddressUpdateParamsStatusDisabled LockboxAddressUpdateParamsStatus = "disabled"
	LockboxAddressUpdateParamsStatusCanceled LockboxAddressUpdateParamsStatus = "canceled"
)

func (r LockboxAddressUpdateParamsStatus) IsKnown() bool {
	switch r {
	case LockboxAddressUpdateParamsStatusActive, LockboxAddressUpdateParamsStatusDisabled, LockboxAddressUpdateParamsStatusCanceled:
		return true
	}
	return false
}

type LockboxAddressListParams struct {
	CreatedAt param.Field[LockboxAddressListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Filter records to the one with the specified `idempotency_key` you chose for
	// that object. This value is unique across Increase and is used to ensure that a
	// request is only processed once. Learn more about
	// [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey param.Field[string] `query:"idempotency_key"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [LockboxAddressListParams]'s query parameters as
// `url.Values`.
func (r LockboxAddressListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LockboxAddressListParamsCreatedAt struct {
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

// URLQuery serializes [LockboxAddressListParamsCreatedAt]'s query parameters as
// `url.Values`.
func (r LockboxAddressListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

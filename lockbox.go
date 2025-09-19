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

// LockboxService contains methods and other services that help with interacting
// with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLockboxService] method instead.
type LockboxService struct {
	Options []option.RequestOption
}

// NewLockboxService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewLockboxService(opts ...option.RequestOption) (r *LockboxService) {
	r = &LockboxService{}
	r.Options = opts
	return
}

// Create a Lockbox
func (r *LockboxService) New(ctx context.Context, body LockboxNewParams, opts ...option.RequestOption) (res *Lockbox, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "lockboxes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a Lockbox
func (r *LockboxService) Get(ctx context.Context, lockboxID string, opts ...option.RequestOption) (res *Lockbox, err error) {
	opts = slices.Concat(r.Options, opts)
	if lockboxID == "" {
		err = errors.New("missing required lockbox_id parameter")
		return
	}
	path := fmt.Sprintf("lockboxes/%s", lockboxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a Lockbox
func (r *LockboxService) Update(ctx context.Context, lockboxID string, body LockboxUpdateParams, opts ...option.RequestOption) (res *Lockbox, err error) {
	opts = slices.Concat(r.Options, opts)
	if lockboxID == "" {
		err = errors.New("missing required lockbox_id parameter")
		return
	}
	path := fmt.Sprintf("lockboxes/%s", lockboxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List Lockboxes
func (r *LockboxService) List(ctx context.Context, query LockboxListParams, opts ...option.RequestOption) (res *pagination.Page[Lockbox], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "lockboxes"
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

// List Lockboxes
func (r *LockboxService) ListAutoPaging(ctx context.Context, query LockboxListParams, opts ...option.RequestOption) *pagination.PageAutoPager[Lockbox] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Lockboxes are physical locations that can receive mail containing paper checks.
// Increase will automatically create a Check Deposit for checks received this way.
type Lockbox struct {
	// The Lockbox identifier.
	ID string `json:"id,required"`
	// The identifier for the Account checks sent to this lockbox will be deposited
	// into.
	AccountID string `json:"account_id,required"`
	// The mailing address for the Lockbox.
	Address LockboxAddress `json:"address,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Lockbox
	// was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The description you choose for the Lockbox.
	Description string `json:"description,required,nullable"`
	// The idempotency key you chose for this object. This value is unique across
	// Increase and is used to ensure that a request is only processed once. Learn more
	// about [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey string `json:"idempotency_key,required,nullable"`
	// The recipient name you choose for the Lockbox.
	RecipientName string `json:"recipient_name,required,nullable"`
	// This indicates if mail can be sent to this address.
	Status LockboxStatus `json:"status,required"`
	// A constant representing the object's type. For this resource it will always be
	// `lockbox`.
	Type LockboxType `json:"type,required"`
	JSON lockboxJSON `json:"-"`
}

// lockboxJSON contains the JSON metadata for the struct [Lockbox]
type lockboxJSON struct {
	ID             apijson.Field
	AccountID      apijson.Field
	Address        apijson.Field
	CreatedAt      apijson.Field
	Description    apijson.Field
	IdempotencyKey apijson.Field
	RecipientName  apijson.Field
	Status         apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *Lockbox) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockboxJSON) RawJSON() string {
	return r.raw
}

// The mailing address for the Lockbox.
type LockboxAddress struct {
	// The city of the address.
	City string `json:"city,required"`
	// The first line of the address.
	Line1 string `json:"line1,required"`
	// The second line of the address.
	Line2 string `json:"line2,required"`
	// The postal code of the address.
	PostalCode string `json:"postal_code,required"`
	// The recipient line of the address. This will include the recipient name you
	// provide when creating the address, as well as an ATTN suffix to help route the
	// mail to your lockbox. Mail senders must include this ATTN line to receive mail
	// at this Lockbox.
	Recipient string `json:"recipient,required,nullable"`
	// The two-letter United States Postal Service (USPS) abbreviation for the state of
	// the address.
	State string             `json:"state,required"`
	JSON  lockboxAddressJSON `json:"-"`
}

// lockboxAddressJSON contains the JSON metadata for the struct [LockboxAddress]
type lockboxAddressJSON struct {
	City        apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	PostalCode  apijson.Field
	Recipient   apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LockboxAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockboxAddressJSON) RawJSON() string {
	return r.raw
}

// This indicates if mail can be sent to this address.
type LockboxStatus string

const (
	LockboxStatusActive   LockboxStatus = "active"
	LockboxStatusInactive LockboxStatus = "inactive"
)

func (r LockboxStatus) IsKnown() bool {
	switch r {
	case LockboxStatusActive, LockboxStatusInactive:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `lockbox`.
type LockboxType string

const (
	LockboxTypeLockbox LockboxType = "lockbox"
)

func (r LockboxType) IsKnown() bool {
	switch r {
	case LockboxTypeLockbox:
		return true
	}
	return false
}

type LockboxNewParams struct {
	// The Account checks sent to this Lockbox should be deposited into.
	AccountID param.Field[string] `json:"account_id,required"`
	// The description you choose for the Lockbox, for display purposes.
	Description param.Field[string] `json:"description"`
	// The name of the recipient that will receive mail at this location.
	RecipientName param.Field[string] `json:"recipient_name"`
}

func (r LockboxNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LockboxUpdateParams struct {
	// The description you choose for the Lockbox.
	Description param.Field[string] `json:"description"`
	// The recipient name you choose for the Lockbox.
	RecipientName param.Field[string] `json:"recipient_name"`
	// This indicates if checks can be sent to the Lockbox.
	Status param.Field[LockboxUpdateParamsStatus] `json:"status"`
}

func (r LockboxUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// This indicates if checks can be sent to the Lockbox.
type LockboxUpdateParamsStatus string

const (
	LockboxUpdateParamsStatusActive   LockboxUpdateParamsStatus = "active"
	LockboxUpdateParamsStatusInactive LockboxUpdateParamsStatus = "inactive"
)

func (r LockboxUpdateParamsStatus) IsKnown() bool {
	switch r {
	case LockboxUpdateParamsStatusActive, LockboxUpdateParamsStatusInactive:
		return true
	}
	return false
}

type LockboxListParams struct {
	// Filter Lockboxes to those associated with the provided Account.
	AccountID param.Field[string]                     `query:"account_id"`
	CreatedAt param.Field[LockboxListParamsCreatedAt] `query:"created_at"`
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

// URLQuery serializes [LockboxListParams]'s query parameters as `url.Values`.
func (r LockboxListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LockboxListParamsCreatedAt struct {
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

// URLQuery serializes [LockboxListParamsCreatedAt]'s query parameters as
// `url.Values`.
func (r LockboxListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

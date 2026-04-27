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

// LockboxRecipientService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLockboxRecipientService] method instead.
type LockboxRecipientService struct {
	Options []option.RequestOption
}

// NewLockboxRecipientService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLockboxRecipientService(opts ...option.RequestOption) (r *LockboxRecipientService) {
	r = &LockboxRecipientService{}
	r.Options = opts
	return
}

// Create a Lockbox Recipient
func (r *LockboxRecipientService) New(ctx context.Context, body LockboxRecipientNewParams, opts ...option.RequestOption) (res *LockboxRecipient, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "lockbox_recipients"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve a Lockbox Recipient
func (r *LockboxRecipientService) Get(ctx context.Context, lockboxRecipientID string, opts ...option.RequestOption) (res *LockboxRecipient, err error) {
	opts = slices.Concat(r.Options, opts)
	if lockboxRecipientID == "" {
		err = errors.New("missing required lockbox_recipient_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("lockbox_recipients/%s", lockboxRecipientID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update a Lockbox Recipient
func (r *LockboxRecipientService) Update(ctx context.Context, lockboxRecipientID string, body LockboxRecipientUpdateParams, opts ...option.RequestOption) (res *LockboxRecipient, err error) {
	opts = slices.Concat(r.Options, opts)
	if lockboxRecipientID == "" {
		err = errors.New("missing required lockbox_recipient_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("lockbox_recipients/%s", lockboxRecipientID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List Lockbox Recipients
func (r *LockboxRecipientService) List(ctx context.Context, query LockboxRecipientListParams, opts ...option.RequestOption) (res *pagination.Page[LockboxRecipient], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "lockbox_recipients"
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

// List Lockbox Recipients
func (r *LockboxRecipientService) ListAutoPaging(ctx context.Context, query LockboxRecipientListParams, opts ...option.RequestOption) *pagination.PageAutoPager[LockboxRecipient] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Lockbox Recipients represent an inbox at a Lockbox Address. Checks received for
// a Lockbox Recipient are deposited into its associated Account.
type LockboxRecipient struct {
	// The Lockbox Recipient identifier.
	ID string `json:"id" api:"required"`
	// The identifier for the Account that checks sent to this Lockbox Recipient will
	// be deposited into.
	AccountID string `json:"account_id" api:"required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Lockbox
	// Recipient was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The description of the Lockbox Recipient.
	Description string `json:"description" api:"required,nullable"`
	// The idempotency key you chose for this object. This value is unique across
	// Increase and is used to ensure that a request is only processed once. Learn more
	// about [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey string `json:"idempotency_key" api:"required,nullable"`
	// The identifier for the Lockbox Address where this Lockbox Recipient may receive
	// physical mail.
	LockboxAddressID string `json:"lockbox_address_id" api:"required"`
	// The mail stop code uniquely identifying this Lockbox Recipient at its Lockbox
	// Address. It should be included in the mailing address intended for this Lockbox
	// Recipient.
	MailStopCode string `json:"mail_stop_code" api:"required"`
	// The name of the Lockbox Recipient.
	RecipientName string `json:"recipient_name" api:"required,nullable"`
	// The status of the Lockbox Recipient.
	Status LockboxRecipientStatus `json:"status" api:"required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `lockbox_recipient`.
	Type LockboxRecipientType `json:"type" api:"required"`
	JSON lockboxRecipientJSON `json:"-"`
}

// lockboxRecipientJSON contains the JSON metadata for the struct
// [LockboxRecipient]
type lockboxRecipientJSON struct {
	ID               apijson.Field
	AccountID        apijson.Field
	CreatedAt        apijson.Field
	Description      apijson.Field
	IdempotencyKey   apijson.Field
	LockboxAddressID apijson.Field
	MailStopCode     apijson.Field
	RecipientName    apijson.Field
	Status           apijson.Field
	Type             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LockboxRecipient) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockboxRecipientJSON) RawJSON() string {
	return r.raw
}

// The status of the Lockbox Recipient.
type LockboxRecipientStatus string

const (
	LockboxRecipientStatusActive   LockboxRecipientStatus = "active"
	LockboxRecipientStatusDisabled LockboxRecipientStatus = "disabled"
	LockboxRecipientStatusCanceled LockboxRecipientStatus = "canceled"
)

func (r LockboxRecipientStatus) IsKnown() bool {
	switch r {
	case LockboxRecipientStatusActive, LockboxRecipientStatusDisabled, LockboxRecipientStatusCanceled:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `lockbox_recipient`.
type LockboxRecipientType string

const (
	LockboxRecipientTypeLockboxRecipient LockboxRecipientType = "lockbox_recipient"
)

func (r LockboxRecipientType) IsKnown() bool {
	switch r {
	case LockboxRecipientTypeLockboxRecipient:
		return true
	}
	return false
}

type LockboxRecipientNewParams struct {
	// The Account that checks sent to this Lockbox Recipient should be deposited into.
	AccountID param.Field[string] `json:"account_id" api:"required"`
	// The Lockbox Address where this Lockbox Recipient may receive mail.
	LockboxAddressID param.Field[string] `json:"lockbox_address_id" api:"required"`
	// The description you choose for the Lockbox Recipient.
	Description param.Field[string] `json:"description"`
	// The name of the Lockbox Recipient
	RecipientName param.Field[string] `json:"recipient_name"`
}

func (r LockboxRecipientNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LockboxRecipientUpdateParams struct {
	// The description you choose for the Lockbox Recipient.
	Description param.Field[string] `json:"description"`
	// The name of the Lockbox Recipient.
	RecipientName param.Field[string] `json:"recipient_name"`
	// The status of the Lockbox Recipient.
	Status param.Field[LockboxRecipientUpdateParamsStatus] `json:"status"`
}

func (r LockboxRecipientUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The status of the Lockbox Recipient.
type LockboxRecipientUpdateParamsStatus string

const (
	LockboxRecipientUpdateParamsStatusActive   LockboxRecipientUpdateParamsStatus = "active"
	LockboxRecipientUpdateParamsStatusDisabled LockboxRecipientUpdateParamsStatus = "disabled"
	LockboxRecipientUpdateParamsStatusCanceled LockboxRecipientUpdateParamsStatus = "canceled"
)

func (r LockboxRecipientUpdateParamsStatus) IsKnown() bool {
	switch r {
	case LockboxRecipientUpdateParamsStatusActive, LockboxRecipientUpdateParamsStatusDisabled, LockboxRecipientUpdateParamsStatusCanceled:
		return true
	}
	return false
}

type LockboxRecipientListParams struct {
	// Filter Lockbox Recipients to those associated with the provided Account.
	AccountID param.Field[string]                              `query:"account_id"`
	CreatedAt param.Field[LockboxRecipientListParamsCreatedAt] `query:"created_at"`
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
	// Filter Lockbox Recipients to those associated with the provided Lockbox Address.
	LockboxAddressID param.Field[string] `query:"lockbox_address_id"`
}

// URLQuery serializes [LockboxRecipientListParams]'s query parameters as
// `url.Values`.
func (r LockboxRecipientListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LockboxRecipientListParamsCreatedAt struct {
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

// URLQuery serializes [LockboxRecipientListParamsCreatedAt]'s query parameters as
// `url.Values`.
func (r LockboxRecipientListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

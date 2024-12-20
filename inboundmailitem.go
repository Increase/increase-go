// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Increase/increase-go/internal/apijson"
	"github.com/Increase/increase-go/internal/apiquery"
	"github.com/Increase/increase-go/internal/param"
	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
	"github.com/Increase/increase-go/packages/pagination"
)

// InboundMailItemService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInboundMailItemService] method instead.
type InboundMailItemService struct {
	Options []option.RequestOption
}

// NewInboundMailItemService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInboundMailItemService(opts ...option.RequestOption) (r *InboundMailItemService) {
	r = &InboundMailItemService{}
	r.Options = opts
	return
}

// Retrieve an Inbound Mail Item
func (r *InboundMailItemService) Get(ctx context.Context, inboundMailItemID string, opts ...option.RequestOption) (res *InboundMailItem, err error) {
	opts = append(r.Options[:], opts...)
	if inboundMailItemID == "" {
		err = errors.New("missing required inbound_mail_item_id parameter")
		return
	}
	path := fmt.Sprintf("inbound_mail_items/%s", inboundMailItemID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Inbound Mail Items
func (r *InboundMailItemService) List(ctx context.Context, query InboundMailItemListParams, opts ...option.RequestOption) (res *pagination.Page[InboundMailItem], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "inbound_mail_items"
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

// List Inbound Mail Items
func (r *InboundMailItemService) ListAutoPaging(ctx context.Context, query InboundMailItemListParams, opts ...option.RequestOption) *pagination.PageAutoPager[InboundMailItem] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Inbound Mail Items represent pieces of physical mail delivered to a Lockbox.
type InboundMailItem struct {
	// The Inbound Mail Item identifier.
	ID string `json:"id,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Inbound
	// Mail Item was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The identifier for the File containing the scanned contents of the mail item.
	FileID string `json:"file_id,required"`
	// The identifier for the Lockbox that received this mail item. For mail items that
	// could not be processed due to an invalid address, this will be null.
	LockboxID string `json:"lockbox_id,required,nullable"`
	// The recipient name as written on the mail item.
	RecipientName string `json:"recipient_name,required,nullable"`
	// If the mail item has been rejected, why it was rejected.
	RejectionReason InboundMailItemRejectionReason `json:"rejection_reason,required,nullable"`
	// If the mail item has been processed.
	Status InboundMailItemStatus `json:"status,required"`
	// A constant representing the object's type. For this resource it will always be
	// `inbound_mail_item`.
	Type InboundMailItemType `json:"type,required"`
	JSON inboundMailItemJSON `json:"-"`
}

// inboundMailItemJSON contains the JSON metadata for the struct [InboundMailItem]
type inboundMailItemJSON struct {
	ID              apijson.Field
	CreatedAt       apijson.Field
	FileID          apijson.Field
	LockboxID       apijson.Field
	RecipientName   apijson.Field
	RejectionReason apijson.Field
	Status          apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *InboundMailItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboundMailItemJSON) RawJSON() string {
	return r.raw
}

// If the mail item has been rejected, why it was rejected.
type InboundMailItemRejectionReason string

const (
	// The mail item does not match any lockbox.
	InboundMailItemRejectionReasonNoMatchingLockbox InboundMailItemRejectionReason = "no_matching_lockbox"
	// The mail item does not contain a check.
	InboundMailItemRejectionReasonNoCheck InboundMailItemRejectionReason = "no_check"
	// The Lockbox or its associated Account is not active.
	InboundMailItemRejectionReasonLockboxNotActive InboundMailItemRejectionReason = "lockbox_not_active"
)

func (r InboundMailItemRejectionReason) IsKnown() bool {
	switch r {
	case InboundMailItemRejectionReasonNoMatchingLockbox, InboundMailItemRejectionReasonNoCheck, InboundMailItemRejectionReasonLockboxNotActive:
		return true
	}
	return false
}

// If the mail item has been processed.
type InboundMailItemStatus string

const (
	// The mail item is pending processing.
	InboundMailItemStatusPending InboundMailItemStatus = "pending"
	// The mail item has been processed.
	InboundMailItemStatusProcessed InboundMailItemStatus = "processed"
	// The mail item has been rejected.
	InboundMailItemStatusRejected InboundMailItemStatus = "rejected"
)

func (r InboundMailItemStatus) IsKnown() bool {
	switch r {
	case InboundMailItemStatusPending, InboundMailItemStatusProcessed, InboundMailItemStatusRejected:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `inbound_mail_item`.
type InboundMailItemType string

const (
	InboundMailItemTypeInboundMailItem InboundMailItemType = "inbound_mail_item"
)

func (r InboundMailItemType) IsKnown() bool {
	switch r {
	case InboundMailItemTypeInboundMailItem:
		return true
	}
	return false
}

type InboundMailItemListParams struct {
	CreatedAt param.Field[InboundMailItemListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
	// Filter Inbound Mail Items to ones sent to the provided Lockbox.
	LockboxID param.Field[string] `query:"lockbox_id"`
}

// URLQuery serializes [InboundMailItemListParams]'s query parameters as
// `url.Values`.
func (r InboundMailItemListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type InboundMailItemListParamsCreatedAt struct {
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

// URLQuery serializes [InboundMailItemListParamsCreatedAt]'s query parameters as
// `url.Values`.
func (r InboundMailItemListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

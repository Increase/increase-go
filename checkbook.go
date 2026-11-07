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

// CheckbookService contains methods and other services that help with interacting
// with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCheckbookService] method instead.
type CheckbookService struct {
	Options []option.RequestOption
}

// NewCheckbookService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCheckbookService(opts ...option.RequestOption) (r *CheckbookService) {
	r = &CheckbookService{}
	r.Options = opts
	return
}

// Create a Checkbook
func (r *CheckbookService) New(ctx context.Context, body CheckbookNewParams, opts ...option.RequestOption) (res *Checkbook, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "checkbooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve a Checkbook
func (r *CheckbookService) Get(ctx context.Context, checkbookID string, opts ...option.RequestOption) (res *Checkbook, err error) {
	opts = slices.Concat(r.Options, opts)
	if checkbookID == "" {
		err = errors.New("missing required checkbook_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("checkbooks/%s", checkbookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List Checkbooks
func (r *CheckbookService) List(ctx context.Context, query CheckbookListParams, opts ...option.RequestOption) (res *pagination.Page[Checkbook], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "checkbooks"
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

// List Checkbooks
func (r *CheckbookService) ListAutoPaging(ctx context.Context, query CheckbookListParams, opts ...option.RequestOption) *pagination.PageAutoPager[Checkbook] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Stop a Checkbook. Stopping a Checkbook stops all of its checks that have not yet
// been deposited.
func (r *CheckbookService) Stop(ctx context.Context, checkbookID string, opts ...option.RequestOption) (res *Checkbook, err error) {
	opts = slices.Concat(r.Options, opts)
	if checkbookID == "" {
		err = errors.New("missing required checkbook_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("checkbooks/%s/stop", checkbookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Checkbooks are sets of blank checks that Increase prints and mails on your
// behalf.
type Checkbook struct {
	// The Checkbook's identifier.
	ID string `json:"id" api:"required"`
	// The identifier of the Account Number the Checkbook is attached to.
	AccountNumberID string `json:"account_number_id" api:"required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Checkbook was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The idempotency key you chose for this object. This value is unique across
	// Increase and is used to ensure that a request is only processed once. Learn more
	// about [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey string `json:"idempotency_key" api:"required,nullable"`
	// The destination address that the Checkbook will be mailed to.
	MailingAddress CheckbookMailingAddress `json:"mailing_address" api:"required"`
	// The payer information printed on the top-left corner of each check.
	Payer []CheckbookPayer `json:"payer" api:"required"`
	// The address the Checkbook will return to if there is a delivery issue.
	ReturnAddress CheckbookReturnAddress `json:"return_address" api:"required"`
	// The lifecycle status of the Checkbook.
	Status CheckbookStatus `json:"status" api:"required"`
	// The physical format of the Checkbook.
	Style CheckbookStyle `json:"style" api:"required"`
	// Tracking updates relating to the Checkbook's delivery.
	TrackingUpdates []CheckbookTrackingUpdate `json:"tracking_updates" api:"required"`
	// A constant representing the object's type. For this resource it will always be
	// `checkbook`.
	Type CheckbookType `json:"type" api:"required"`
	JSON checkbookJSON `json:"-"`
}

// checkbookJSON contains the JSON metadata for the struct [Checkbook]
type checkbookJSON struct {
	ID              apijson.Field
	AccountNumberID apijson.Field
	CreatedAt       apijson.Field
	IdempotencyKey  apijson.Field
	MailingAddress  apijson.Field
	Payer           apijson.Field
	ReturnAddress   apijson.Field
	Status          apijson.Field
	Style           apijson.Field
	TrackingUpdates apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *Checkbook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkbookJSON) RawJSON() string {
	return r.raw
}

// The destination address that the Checkbook will be mailed to.
type CheckbookMailingAddress struct {
	// The city of the destination address.
	City string `json:"city" api:"required"`
	// The first line of the destination address.
	Line1 string `json:"line1" api:"required"`
	// The second line of the destination address.
	Line2 string `json:"line2" api:"required,nullable"`
	// The recipient at the destination address.
	Name string `json:"name" api:"required"`
	// The postal code of the destination address.
	PostalCode string `json:"postal_code" api:"required"`
	// The state of the destination address.
	State string                      `json:"state" api:"required"`
	JSON  checkbookMailingAddressJSON `json:"-"`
}

// checkbookMailingAddressJSON contains the JSON metadata for the struct
// [CheckbookMailingAddress]
type checkbookMailingAddressJSON struct {
	City        apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	Name        apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CheckbookMailingAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkbookMailingAddressJSON) RawJSON() string {
	return r.raw
}

type CheckbookPayer struct {
	// The contents of the line.
	Contents string             `json:"contents" api:"required"`
	JSON     checkbookPayerJSON `json:"-"`
}

// checkbookPayerJSON contains the JSON metadata for the struct [CheckbookPayer]
type checkbookPayerJSON struct {
	Contents    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CheckbookPayer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkbookPayerJSON) RawJSON() string {
	return r.raw
}

// The address the Checkbook will return to if there is a delivery issue.
type CheckbookReturnAddress struct {
	// The city of the return address.
	City string `json:"city" api:"required"`
	// The first line of the return address.
	Line1 string `json:"line1" api:"required"`
	// The second line of the return address.
	Line2 string `json:"line2" api:"required,nullable"`
	// The recipient at the return address.
	Name string `json:"name" api:"required"`
	// The postal code of the return address.
	PostalCode string `json:"postal_code" api:"required"`
	// The state of the return address.
	State string                     `json:"state" api:"required"`
	JSON  checkbookReturnAddressJSON `json:"-"`
}

// checkbookReturnAddressJSON contains the JSON metadata for the struct
// [CheckbookReturnAddress]
type checkbookReturnAddressJSON struct {
	City        apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	Name        apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CheckbookReturnAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkbookReturnAddressJSON) RawJSON() string {
	return r.raw
}

// The lifecycle status of the Checkbook.
type CheckbookStatus string

const (
	CheckbookStatusPendingSubmission CheckbookStatus = "pending_submission"
	CheckbookStatusPendingMailing    CheckbookStatus = "pending_mailing"
	CheckbookStatusMailed            CheckbookStatus = "mailed"
	CheckbookStatusStopped           CheckbookStatus = "stopped"
)

func (r CheckbookStatus) IsKnown() bool {
	switch r {
	case CheckbookStatusPendingSubmission, CheckbookStatusPendingMailing, CheckbookStatusMailed, CheckbookStatusStopped:
		return true
	}
	return false
}

// The physical format of the Checkbook.
type CheckbookStyle string

const (
	CheckbookStyleThreeRingBinder CheckbookStyle = "three_ring_binder"
	CheckbookStyleFolder          CheckbookStyle = "folder"
)

func (r CheckbookStyle) IsKnown() bool {
	switch r {
	case CheckbookStyleThreeRingBinder, CheckbookStyleFolder:
		return true
	}
	return false
}

type CheckbookTrackingUpdate struct {
	// The type of tracking event.
	Category CheckbookTrackingUpdatesCategory `json:"category" api:"required"`
	// The ISO 3166-1 alpha-2 country code for the country where the event took place.
	Country string `json:"country" api:"required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the tracking event took place.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The postal code where the event took place.
	PostalCode string                      `json:"postal_code" api:"required"`
	JSON       checkbookTrackingUpdateJSON `json:"-"`
}

// checkbookTrackingUpdateJSON contains the JSON metadata for the struct
// [CheckbookTrackingUpdate]
type checkbookTrackingUpdateJSON struct {
	Category    apijson.Field
	Country     apijson.Field
	CreatedAt   apijson.Field
	PostalCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CheckbookTrackingUpdate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkbookTrackingUpdateJSON) RawJSON() string {
	return r.raw
}

// The type of tracking event.
type CheckbookTrackingUpdatesCategory string

const (
	CheckbookTrackingUpdatesCategoryInTransit            CheckbookTrackingUpdatesCategory = "in_transit"
	CheckbookTrackingUpdatesCategoryProcessedForDelivery CheckbookTrackingUpdatesCategory = "processed_for_delivery"
	CheckbookTrackingUpdatesCategoryDelivered            CheckbookTrackingUpdatesCategory = "delivered"
	CheckbookTrackingUpdatesCategoryDeliveryIssue        CheckbookTrackingUpdatesCategory = "delivery_issue"
	CheckbookTrackingUpdatesCategoryReturningToSender    CheckbookTrackingUpdatesCategory = "returning_to_sender"
	CheckbookTrackingUpdatesCategoryReturnedToSender     CheckbookTrackingUpdatesCategory = "returned_to_sender"
)

func (r CheckbookTrackingUpdatesCategory) IsKnown() bool {
	switch r {
	case CheckbookTrackingUpdatesCategoryInTransit, CheckbookTrackingUpdatesCategoryProcessedForDelivery, CheckbookTrackingUpdatesCategoryDelivered, CheckbookTrackingUpdatesCategoryDeliveryIssue, CheckbookTrackingUpdatesCategoryReturningToSender, CheckbookTrackingUpdatesCategoryReturnedToSender:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `checkbook`.
type CheckbookType string

const (
	CheckbookTypeCheckbook CheckbookType = "checkbook"
)

func (r CheckbookType) IsKnown() bool {
	switch r {
	case CheckbookTypeCheckbook:
		return true
	}
	return false
}

type CheckbookNewParams struct {
	// The identifier of the Account Number that the checks will draw from. The Account
	// Number must be active.
	AccountNumberID param.Field[string] `json:"account_number_id" api:"required"`
	// The destination address that the Checkbook will be mailed to.
	MailingAddress param.Field[CheckbookNewParamsMailingAddress] `json:"mailing_address" api:"required"`
	// The payer of the checks. This will be printed on the top-left corner of each
	// check. This should be an array of up to 4 elements, each of which represents a
	// line of the payer.
	Payer param.Field[[]CheckbookNewParamsPayer] `json:"payer" api:"required"`
}

func (r CheckbookNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The destination address that the Checkbook will be mailed to.
type CheckbookNewParamsMailingAddress struct {
	// The city of the destination address.
	City param.Field[string] `json:"city" api:"required"`
	// The first line of the destination address.
	Line1 param.Field[string] `json:"line1" api:"required"`
	// The recipient at the destination address.
	Name param.Field[string] `json:"name" api:"required"`
	// The postal code of the destination address.
	PostalCode param.Field[string] `json:"postal_code" api:"required"`
	// The US state of the destination address.
	State param.Field[string] `json:"state" api:"required"`
	// The second line of the destination address.
	Line2 param.Field[string] `json:"line2"`
}

func (r CheckbookNewParamsMailingAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CheckbookNewParamsPayer struct {
	// The contents of the line.
	Contents param.Field[string] `json:"contents" api:"required"`
}

func (r CheckbookNewParamsPayer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CheckbookListParams struct {
	// Filter Checkbooks to those attached to the specified Account Number.
	AccountNumberID param.Field[string] `query:"account_number_id"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Filter records to the one with the specified `idempotency_key` you chose for
	// that object. This value is unique across Increase and is used to ensure that a
	// request is only processed once. Learn more about
	// [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey param.Field[string] `query:"idempotency_key"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	//
	// Defaults to `100`.
	Limit  param.Field[int64]                     `query:"limit"`
	Status param.Field[CheckbookListParamsStatus] `query:"status"`
}

// URLQuery serializes [CheckbookListParams]'s query parameters as `url.Values`.
func (r CheckbookListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type CheckbookListParamsStatus struct {
	// Filter Checkbooks for those with the specified status or statuses. For GET
	// requests, this should be encoded as a comma-delimited string, such as
	// `?in=one,two,three`.
	In param.Field[[]CheckbookListParamsStatusIn] `query:"in"`
}

// URLQuery serializes [CheckbookListParamsStatus]'s query parameters as
// `url.Values`.
func (r CheckbookListParamsStatus) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type CheckbookListParamsStatusIn string

const (
	CheckbookListParamsStatusInPendingSubmission CheckbookListParamsStatusIn = "pending_submission"
	CheckbookListParamsStatusInPendingMailing    CheckbookListParamsStatusIn = "pending_mailing"
	CheckbookListParamsStatusInMailed            CheckbookListParamsStatusIn = "mailed"
	CheckbookListParamsStatusInStopped           CheckbookListParamsStatusIn = "stopped"
)

func (r CheckbookListParamsStatusIn) IsKnown() bool {
	switch r {
	case CheckbookListParamsStatusInPendingSubmission, CheckbookListParamsStatusInPendingMailing, CheckbookListParamsStatusInMailed, CheckbookListParamsStatusInStopped:
		return true
	}
	return false
}

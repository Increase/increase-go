// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/Increase/increase-go/internal/apijson"
	"github.com/Increase/increase-go/internal/param"
	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
)

// PhysicalCheckBatchService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPhysicalCheckBatchService] method instead.
type PhysicalCheckBatchService struct {
	Options []option.RequestOption
}

// NewPhysicalCheckBatchService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPhysicalCheckBatchService(opts ...option.RequestOption) (r *PhysicalCheckBatchService) {
	r = &PhysicalCheckBatchService{}
	r.Options = opts
	return
}

// Create a Physical Check Batch
func (r *PhysicalCheckBatchService) New(ctx context.Context, body PhysicalCheckBatchNewParams, opts ...option.RequestOption) (res *PhysicalCheckBatch, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "physical_check_batches"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Cancel a pending Physical Check Batch, which cancels all of its related checks.
func (r *PhysicalCheckBatchService) Cancel(ctx context.Context, physicalCheckBatchID string, opts ...option.RequestOption) (res *PhysicalCheckBatch, err error) {
	opts = slices.Concat(r.Options, opts)
	if physicalCheckBatchID == "" {
		err = errors.New("missing required physical_check_batch_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("physical_check_batches/%s/cancel", physicalCheckBatchID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Completing a Physical Check Batch closes it to new Physical Checks and begins
// the process of printing and mailing it.
func (r *PhysicalCheckBatchService) Complete(ctx context.Context, physicalCheckBatchID string, opts ...option.RequestOption) (res *PhysicalCheckBatch, err error) {
	opts = slices.Concat(r.Options, opts)
	if physicalCheckBatchID == "" {
		err = errors.New("missing required physical_check_batch_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("physical_check_batches/%s/complete", physicalCheckBatchID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Physical Check Batches are groups of checks that are mailed in the same parcel.
// Tracking updates are propagated to every related Check Transfer.
type PhysicalCheckBatch struct {
	// The Physical Check Batch's identifier.
	ID string `json:"id" api:"required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Physical Check Batch was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The idempotency key you chose for this object. This value is unique across
	// Increase and is used to ensure that a request is only processed once. Learn more
	// about [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey string `json:"idempotency_key" api:"required,nullable"`
	// The mailing address of the parcel.
	MailingAddress PhysicalCheckBatchMailingAddress `json:"mailing_address" api:"required"`
	// The return address of the parcel.
	ReturnAddress PhysicalCheckBatchReturnAddress `json:"return_address" api:"required"`
	// The shipping method for the parcel.
	ShippingMethod PhysicalCheckBatchShippingMethod `json:"shipping_method" api:"required"`
	// The lifecycle status of the Physical Check Batch.
	Status PhysicalCheckBatchStatus `json:"status" api:"required"`
	// A constant representing the object's type. For this resource it will always be
	// `physical_check_batch`.
	Type PhysicalCheckBatchType `json:"type" api:"required"`
	JSON physicalCheckBatchJSON `json:"-"`
}

// physicalCheckBatchJSON contains the JSON metadata for the struct
// [PhysicalCheckBatch]
type physicalCheckBatchJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	IdempotencyKey apijson.Field
	MailingAddress apijson.Field
	ReturnAddress  apijson.Field
	ShippingMethod apijson.Field
	Status         apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhysicalCheckBatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r physicalCheckBatchJSON) RawJSON() string {
	return r.raw
}

// The mailing address of the parcel.
type PhysicalCheckBatchMailingAddress struct {
	// The city of the address.
	City string `json:"city" api:"required"`
	// The first line of the address.
	Line1 string `json:"line1" api:"required"`
	// The second line of the address.
	Line2 string `json:"line2" api:"required,nullable"`
	// The name component of the address.
	Name string `json:"name" api:"required"`
	// The phone number that is used for delivery issues.
	Phone string `json:"phone" api:"required,nullable"`
	// The postal code of the address.
	PostalCode string `json:"postal_code" api:"required"`
	// The state of the address.
	State string                               `json:"state" api:"required"`
	JSON  physicalCheckBatchMailingAddressJSON `json:"-"`
}

// physicalCheckBatchMailingAddressJSON contains the JSON metadata for the struct
// [PhysicalCheckBatchMailingAddress]
type physicalCheckBatchMailingAddressJSON struct {
	City        apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	Name        apijson.Field
	Phone       apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhysicalCheckBatchMailingAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r physicalCheckBatchMailingAddressJSON) RawJSON() string {
	return r.raw
}

// The return address of the parcel.
type PhysicalCheckBatchReturnAddress struct {
	// The city of the return address.
	City string `json:"city" api:"required"`
	// The first line of the return address.
	Line1 string `json:"line1" api:"required"`
	// The second line of the return address.
	Line2 string `json:"line2" api:"required,nullable"`
	// The name component of the return address.
	Name string `json:"name" api:"required"`
	// The phone number that is used for delivery issues.
	Phone string `json:"phone" api:"required,nullable"`
	// The postal code of the return address.
	PostalCode string `json:"postal_code" api:"required"`
	// The state of the return address.
	State string                              `json:"state" api:"required"`
	JSON  physicalCheckBatchReturnAddressJSON `json:"-"`
}

// physicalCheckBatchReturnAddressJSON contains the JSON metadata for the struct
// [PhysicalCheckBatchReturnAddress]
type physicalCheckBatchReturnAddressJSON struct {
	City        apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	Name        apijson.Field
	Phone       apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhysicalCheckBatchReturnAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r physicalCheckBatchReturnAddressJSON) RawJSON() string {
	return r.raw
}

// The shipping method for the parcel.
type PhysicalCheckBatchShippingMethod string

const (
	PhysicalCheckBatchShippingMethodUspsFirstClass PhysicalCheckBatchShippingMethod = "usps_first_class"
	PhysicalCheckBatchShippingMethodFedexOvernight PhysicalCheckBatchShippingMethod = "fedex_overnight"
)

func (r PhysicalCheckBatchShippingMethod) IsKnown() bool {
	switch r {
	case PhysicalCheckBatchShippingMethodUspsFirstClass, PhysicalCheckBatchShippingMethodFedexOvernight:
		return true
	}
	return false
}

// The lifecycle status of the Physical Check Batch.
type PhysicalCheckBatchStatus string

const (
	PhysicalCheckBatchStatusPending           PhysicalCheckBatchStatus = "pending"
	PhysicalCheckBatchStatusCompleted         PhysicalCheckBatchStatus = "completed"
	PhysicalCheckBatchStatusCanceled          PhysicalCheckBatchStatus = "canceled"
	PhysicalCheckBatchStatusRequiresAttention PhysicalCheckBatchStatus = "requires_attention"
)

func (r PhysicalCheckBatchStatus) IsKnown() bool {
	switch r {
	case PhysicalCheckBatchStatusPending, PhysicalCheckBatchStatusCompleted, PhysicalCheckBatchStatusCanceled, PhysicalCheckBatchStatusRequiresAttention:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `physical_check_batch`.
type PhysicalCheckBatchType string

const (
	PhysicalCheckBatchTypePhysicalCheckBatch PhysicalCheckBatchType = "physical_check_batch"
)

func (r PhysicalCheckBatchType) IsKnown() bool {
	switch r {
	case PhysicalCheckBatchTypePhysicalCheckBatch:
		return true
	}
	return false
}

type PhysicalCheckBatchNewParams struct {
	// Details for where the parcel will be mailed.
	MailingAddress param.Field[PhysicalCheckBatchNewParamsMailingAddress] `json:"mailing_address" api:"required"`
	// Details for where the parcel should return if it is unable to be delivered.
	ReturnAddress param.Field[PhysicalCheckBatchNewParamsReturnAddress] `json:"return_address" api:"required"`
	// How to ship the batch.
	//
	// Defaults to `usps_first_class`.
	ShippingMethod param.Field[PhysicalCheckBatchNewParamsShippingMethod] `json:"shipping_method"`
}

func (r PhysicalCheckBatchNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Details for where the parcel will be mailed.
type PhysicalCheckBatchNewParamsMailingAddress struct {
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
	// The phone number used for delivery issues at the destination address. Only used
	// when `shipping_method` is `fedex_overnight`.
	Phone param.Field[string] `json:"phone"`
}

func (r PhysicalCheckBatchNewParamsMailingAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Details for where the parcel should return if it is unable to be delivered.
type PhysicalCheckBatchNewParamsReturnAddress struct {
	// The city of the return address.
	City param.Field[string] `json:"city" api:"required"`
	// The first line of the return address.
	Line1 param.Field[string] `json:"line1" api:"required"`
	// The recipient at the return address.
	Name param.Field[string] `json:"name" api:"required"`
	// The postal code of the return address.
	PostalCode param.Field[string] `json:"postal_code" api:"required"`
	// The US state of the return address.
	State param.Field[string] `json:"state" api:"required"`
	// The second line of the return address.
	Line2 param.Field[string] `json:"line2"`
	// The phone number used for delivery issues at the return address. Only used when
	// `shipping_method` is `fedex_overnight`.
	Phone param.Field[string] `json:"phone"`
}

func (r PhysicalCheckBatchNewParamsReturnAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// How to ship the batch.
type PhysicalCheckBatchNewParamsShippingMethod string

const (
	PhysicalCheckBatchNewParamsShippingMethodUspsFirstClass PhysicalCheckBatchNewParamsShippingMethod = "usps_first_class"
	PhysicalCheckBatchNewParamsShippingMethodFedexOvernight PhysicalCheckBatchNewParamsShippingMethod = "fedex_overnight"
)

func (r PhysicalCheckBatchNewParamsShippingMethod) IsKnown() bool {
	switch r {
	case PhysicalCheckBatchNewParamsShippingMethodUspsFirstClass, PhysicalCheckBatchNewParamsShippingMethodFedexOvernight:
		return true
	}
	return false
}

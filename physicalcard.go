// File generated from our OpenAPI spec by Stainless.

package increase

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/apiquery"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/internal/shared"
	"github.com/increase/increase-go/option"
)

// PhysicalCardService contains methods and other services that help with
// interacting with the increase API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewPhysicalCardService] method
// instead.
type PhysicalCardService struct {
	Options []option.RequestOption
}

// NewPhysicalCardService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPhysicalCardService(opts ...option.RequestOption) (r *PhysicalCardService) {
	r = &PhysicalCardService{}
	r.Options = opts
	return
}

// Create a Physical Card
func (r *PhysicalCardService) New(ctx context.Context, body PhysicalCardNewParams, opts ...option.RequestOption) (res *PhysicalCard, err error) {
	opts = append(r.Options[:], opts...)
	path := "physical_cards"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a Physical Card
func (r *PhysicalCardService) Get(ctx context.Context, physicalCardID string, opts ...option.RequestOption) (res *PhysicalCard, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("physical_cards/%s", physicalCardID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a Physical Card
func (r *PhysicalCardService) Update(ctx context.Context, physicalCardID string, body PhysicalCardUpdateParams, opts ...option.RequestOption) (res *PhysicalCard, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("physical_cards/%s", physicalCardID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List Physical Cards
func (r *PhysicalCardService) List(ctx context.Context, query PhysicalCardListParams, opts ...option.RequestOption) (res *shared.Page[PhysicalCard], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "physical_cards"
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

// List Physical Cards
func (r *PhysicalCardService) ListAutoPaging(ctx context.Context, query PhysicalCardListParams, opts ...option.RequestOption) *shared.PageAutoPager[PhysicalCard] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Custom physical Visa cards that are shipped to your customers. The artwork is
// configurable by a connected [Card Profile](/documentation/api#card-profiles).
// The same Card can be used for multiple Physical Cards. Printing cards incurs a
// fee. Please contact [support@increase.com](mailto:support@increase.com) for
// pricing!
type PhysicalCard struct {
	// The physical card identifier.
	ID string `json:"id,required"`
	// The identifier for the Card this Physical Card represents.
	CardID string `json:"card_id,required"`
	// The Card Profile used for this Physical Card.
	CardProfileID string `json:"card_profile_id,required,nullable"`
	// Details about the cardholder, as it appears on the printed card.
	Cardholder PhysicalCardCardholder `json:"cardholder,required,nullable"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Physical Card was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The idempotency key you chose for this object. This value is unique across
	// Increase and is used to ensure that a request is only processed once. Learn more
	// about [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey string `json:"idempotency_key,required,nullable"`
	// The Physical Card Profile used for this Physical Card.
	PhysicalCardProfileID string `json:"physical_card_profile_id,required,nullable"`
	// The details used to ship this physical card.
	Shipment PhysicalCardShipment `json:"shipment,required"`
	// The status of the Physical Card.
	Status PhysicalCardStatus `json:"status,required"`
	// A constant representing the object's type. For this resource it will always be
	// `physical_card`.
	Type PhysicalCardType `json:"type,required"`
	JSON physicalCardJSON `json:"-"`
}

// physicalCardJSON contains the JSON metadata for the struct [PhysicalCard]
type physicalCardJSON struct {
	ID                    apijson.Field
	CardID                apijson.Field
	CardProfileID         apijson.Field
	Cardholder            apijson.Field
	CreatedAt             apijson.Field
	IdempotencyKey        apijson.Field
	PhysicalCardProfileID apijson.Field
	Shipment              apijson.Field
	Status                apijson.Field
	Type                  apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *PhysicalCard) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details about the cardholder, as it appears on the printed card.
type PhysicalCardCardholder struct {
	// The cardholder's first name.
	FirstName string `json:"first_name,required"`
	// The cardholder's last name.
	LastName string                     `json:"last_name,required"`
	JSON     physicalCardCardholderJSON `json:"-"`
}

// physicalCardCardholderJSON contains the JSON metadata for the struct
// [PhysicalCardCardholder]
type physicalCardCardholderJSON struct {
	FirstName   apijson.Field
	LastName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhysicalCardCardholder) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The details used to ship this physical card.
type PhysicalCardShipment struct {
	// The location to where the card's packing label is addressed.
	Address PhysicalCardShipmentAddress `json:"address,required"`
	// The shipping method.
	Method PhysicalCardShipmentMethod `json:"method,required"`
	// The status of this shipment.
	Status PhysicalCardShipmentStatus `json:"status,required"`
	// Tracking details for the shipment.
	Tracking PhysicalCardShipmentTracking `json:"tracking,required,nullable"`
	JSON     physicalCardShipmentJSON     `json:"-"`
}

// physicalCardShipmentJSON contains the JSON metadata for the struct
// [PhysicalCardShipment]
type physicalCardShipmentJSON struct {
	Address     apijson.Field
	Method      apijson.Field
	Status      apijson.Field
	Tracking    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhysicalCardShipment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The location to where the card's packing label is addressed.
type PhysicalCardShipmentAddress struct {
	// The city of the shipping address.
	City string `json:"city,required"`
	// The first line of the shipping address.
	Line1 string `json:"line1,required"`
	// The second line of the shipping address.
	Line2 string `json:"line2,required,nullable"`
	// The third line of the shipping address.
	Line3 string `json:"line3,required,nullable"`
	// The name of the recipient.
	Name string `json:"name,required"`
	// The postal code of the shipping address.
	PostalCode string `json:"postal_code,required"`
	// The US state of the shipping address.
	State string                          `json:"state,required"`
	JSON  physicalCardShipmentAddressJSON `json:"-"`
}

// physicalCardShipmentAddressJSON contains the JSON metadata for the struct
// [PhysicalCardShipmentAddress]
type physicalCardShipmentAddressJSON struct {
	City        apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	Line3       apijson.Field
	Name        apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhysicalCardShipmentAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The shipping method.
type PhysicalCardShipmentMethod string

const (
	// USPS Post with tracking.
	PhysicalCardShipmentMethodUsps PhysicalCardShipmentMethod = "usps"
	// FedEx Priority Overnight, no signature.
	PhysicalCardShipmentMethodFedexPriorityOvernight PhysicalCardShipmentMethod = "fedex_priority_overnight"
	// FedEx 2-day.
	PhysicalCardShipmentMethodFedex2Day PhysicalCardShipmentMethod = "fedex_2_day"
)

// The status of this shipment.
type PhysicalCardShipmentStatus string

const (
	// The physical card has not yet been shipped.
	PhysicalCardShipmentStatusPending PhysicalCardShipmentStatus = "pending"
	// The physical card shipment was canceled prior to submission.
	PhysicalCardShipmentStatusCanceled PhysicalCardShipmentStatus = "canceled"
	// The physical card shipment has been submitted to the card fulfillment provider.
	PhysicalCardShipmentStatusSubmitted PhysicalCardShipmentStatus = "submitted"
	// The physical card shipment has been acknowledged by the card fulfillment
	// provider and will be processed in their next batch.
	PhysicalCardShipmentStatusAcknowledged PhysicalCardShipmentStatus = "acknowledged"
	// The physical card shipment was rejected by the card printer due to an error.
	PhysicalCardShipmentStatusRejected PhysicalCardShipmentStatus = "rejected"
	// The physical card has been shipped.
	PhysicalCardShipmentStatusShipped PhysicalCardShipmentStatus = "shipped"
	// The physical card shipment was returned to the sender and destroyed by the
	// production facility.
	PhysicalCardShipmentStatusReturned PhysicalCardShipmentStatus = "returned"
)

// Tracking details for the shipment.
type PhysicalCardShipmentTracking struct {
	// The tracking number.
	Number string `json:"number,required"`
	// For returned shipments, the tracking number of the return shipment.
	ReturnNumber string `json:"return_number,required,nullable"`
	// For returned shipments, this describes why the package was returned.
	ReturnReason string `json:"return_reason,required,nullable"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the fulfillment provider marked the card as ready for pick-up by the shipment
	// carrier.
	ShippedAt time.Time                        `json:"shipped_at,required" format:"date-time"`
	JSON      physicalCardShipmentTrackingJSON `json:"-"`
}

// physicalCardShipmentTrackingJSON contains the JSON metadata for the struct
// [PhysicalCardShipmentTracking]
type physicalCardShipmentTrackingJSON struct {
	Number       apijson.Field
	ReturnNumber apijson.Field
	ReturnReason apijson.Field
	ShippedAt    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PhysicalCardShipmentTracking) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the Physical Card.
type PhysicalCardStatus string

const (
	// The physical card is active.
	PhysicalCardStatusActive PhysicalCardStatus = "active"
	// The physical card is temporarily disabled.
	PhysicalCardStatusDisabled PhysicalCardStatus = "disabled"
	// The physical card is permanently canceled.
	PhysicalCardStatusCanceled PhysicalCardStatus = "canceled"
)

// A constant representing the object's type. For this resource it will always be
// `physical_card`.
type PhysicalCardType string

const (
	PhysicalCardTypePhysicalCard PhysicalCardType = "physical_card"
)

type PhysicalCardNewParams struct {
	// The underlying card representing this physical card.
	CardID param.Field[string] `json:"card_id,required"`
	// Details about the cardholder, as it will appear on the physical card.
	Cardholder param.Field[PhysicalCardNewParamsCardholder] `json:"cardholder,required"`
	// The details used to ship this physical card.
	Shipment param.Field[PhysicalCardNewParamsShipment] `json:"shipment,required"`
	// The physical card profile to use for this physical card. The latest default
	// physical card profile will be used if not provided.
	PhysicalCardProfileID param.Field[string] `json:"physical_card_profile_id"`
}

func (r PhysicalCardNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Details about the cardholder, as it will appear on the physical card.
type PhysicalCardNewParamsCardholder struct {
	// The cardholder's first name.
	FirstName param.Field[string] `json:"first_name,required"`
	// The cardholder's last name.
	LastName param.Field[string] `json:"last_name,required"`
}

func (r PhysicalCardNewParamsCardholder) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The details used to ship this physical card.
type PhysicalCardNewParamsShipment struct {
	// The address to where the card should be shipped.
	Address param.Field[PhysicalCardNewParamsShipmentAddress] `json:"address,required"`
	// The shipping method to use.
	Method param.Field[PhysicalCardNewParamsShipmentMethod] `json:"method,required"`
}

func (r PhysicalCardNewParamsShipment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The address to where the card should be shipped.
type PhysicalCardNewParamsShipmentAddress struct {
	// The city of the shipping address.
	City param.Field[string] `json:"city,required"`
	// The first line of the shipping address.
	Line1 param.Field[string] `json:"line1,required"`
	// The name of the recipient.
	Name param.Field[string] `json:"name,required"`
	// The postal code of the shipping address.
	PostalCode param.Field[string] `json:"postal_code,required"`
	// The US state of the shipping address.
	State param.Field[string] `json:"state,required"`
	// The second line of the shipping address.
	Line2 param.Field[string] `json:"line2"`
	// The third line of the shipping address.
	Line3 param.Field[string] `json:"line3"`
	// The phone number of the recipient.
	PhoneNumber param.Field[string] `json:"phone_number"`
}

func (r PhysicalCardNewParamsShipmentAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The shipping method to use.
type PhysicalCardNewParamsShipmentMethod string

const (
	// USPS Post with tracking.
	PhysicalCardNewParamsShipmentMethodUsps PhysicalCardNewParamsShipmentMethod = "usps"
	// FedEx Priority Overnight, no signature.
	PhysicalCardNewParamsShipmentMethodFedexPriorityOvernight PhysicalCardNewParamsShipmentMethod = "fedex_priority_overnight"
	// FedEx 2-day.
	PhysicalCardNewParamsShipmentMethodFedex2Day PhysicalCardNewParamsShipmentMethod = "fedex_2_day"
)

type PhysicalCardUpdateParams struct {
	// The status to update the Physical Card to.
	Status param.Field[PhysicalCardUpdateParamsStatus] `json:"status,required"`
}

func (r PhysicalCardUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The status to update the Physical Card to.
type PhysicalCardUpdateParamsStatus string

const (
	// The physical card is active.
	PhysicalCardUpdateParamsStatusActive PhysicalCardUpdateParamsStatus = "active"
	// The physical card is temporarily disabled.
	PhysicalCardUpdateParamsStatusDisabled PhysicalCardUpdateParamsStatus = "disabled"
	// The physical card is permanently canceled.
	PhysicalCardUpdateParamsStatusCanceled PhysicalCardUpdateParamsStatus = "canceled"
)

type PhysicalCardListParams struct {
	// Filter Physical Cards to ones belonging to the specified Card.
	CardID    param.Field[string]                          `query:"card_id"`
	CreatedAt param.Field[PhysicalCardListParamsCreatedAt] `query:"created_at"`
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

// URLQuery serializes [PhysicalCardListParams]'s query parameters as `url.Values`.
func (r PhysicalCardListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type PhysicalCardListParamsCreatedAt struct {
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

// URLQuery serializes [PhysicalCardListParamsCreatedAt]'s query parameters as
// `url.Values`.
func (r PhysicalCardListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

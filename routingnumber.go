// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/Increase/increase-go/internal/apijson"
	"github.com/Increase/increase-go/internal/apiquery"
	"github.com/Increase/increase-go/internal/param"
	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
)

// RoutingNumberService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRoutingNumberService] method instead.
type RoutingNumberService struct {
	Options []option.RequestOption
}

// NewRoutingNumberService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRoutingNumberService(opts ...option.RequestOption) (r *RoutingNumberService) {
	r = &RoutingNumberService{}
	r.Options = opts
	return
}

// You can use this API to confirm if a routing number is valid, such as when a
// user is providing you with bank account details. Since routing numbers uniquely
// identify a bank, this will always return 0 or 1 entry. In Sandbox, the only
// valid routing number for this method is 110000000.
func (r *RoutingNumberService) List(ctx context.Context, query RoutingNumberListParams, opts ...option.RequestOption) (res *RoutingNumberListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "routing_numbers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// A list of Routing Number objects.
type RoutingNumberListResponse struct {
	// The contents of the list.
	Data []RoutingNumberListResponseData `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor  string                        `json:"next_cursor,required,nullable"`
	ExtraFields map[string]interface{}        `json:"-,extras"`
	JSON        routingNumberListResponseJSON `json:"-"`
}

// routingNumberListResponseJSON contains the JSON metadata for the struct
// [RoutingNumberListResponse]
type routingNumberListResponseJSON struct {
	Data        apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingNumberListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingNumberListResponseJSON) RawJSON() string {
	return r.raw
}

// Routing numbers are used to identify your bank in a financial transaction.
type RoutingNumberListResponseData struct {
	// This routing number's support for ACH Transfers.
	ACHTransfers RoutingNumberListResponseDataACHTransfers `json:"ach_transfers,required"`
	// This routing number's support for FedNow Transfers.
	FednowTransfers RoutingNumberListResponseDataFednowTransfers `json:"fednow_transfers,required"`
	// The name of the financial institution belonging to a routing number.
	Name string `json:"name,required"`
	// This routing number's support for Real-Time Payments Transfers.
	RealTimePaymentsTransfers RoutingNumberListResponseDataRealTimePaymentsTransfers `json:"real_time_payments_transfers,required"`
	// The nine digit routing number identifier.
	RoutingNumber string `json:"routing_number,required"`
	// A constant representing the object's type. For this resource it will always be
	// `routing_number`.
	Type RoutingNumberListResponseDataType `json:"type,required"`
	// This routing number's support for Wire Transfers.
	WireTransfers RoutingNumberListResponseDataWireTransfers `json:"wire_transfers,required"`
	JSON          routingNumberListResponseDataJSON          `json:"-"`
}

// routingNumberListResponseDataJSON contains the JSON metadata for the struct
// [RoutingNumberListResponseData]
type routingNumberListResponseDataJSON struct {
	ACHTransfers              apijson.Field
	FednowTransfers           apijson.Field
	Name                      apijson.Field
	RealTimePaymentsTransfers apijson.Field
	RoutingNumber             apijson.Field
	Type                      apijson.Field
	WireTransfers             apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *RoutingNumberListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingNumberListResponseDataJSON) RawJSON() string {
	return r.raw
}

// This routing number's support for ACH Transfers.
type RoutingNumberListResponseDataACHTransfers string

const (
	RoutingNumberListResponseDataACHTransfersSupported    RoutingNumberListResponseDataACHTransfers = "supported"
	RoutingNumberListResponseDataACHTransfersNotSupported RoutingNumberListResponseDataACHTransfers = "not_supported"
)

func (r RoutingNumberListResponseDataACHTransfers) IsKnown() bool {
	switch r {
	case RoutingNumberListResponseDataACHTransfersSupported, RoutingNumberListResponseDataACHTransfersNotSupported:
		return true
	}
	return false
}

// This routing number's support for FedNow Transfers.
type RoutingNumberListResponseDataFednowTransfers string

const (
	RoutingNumberListResponseDataFednowTransfersSupported    RoutingNumberListResponseDataFednowTransfers = "supported"
	RoutingNumberListResponseDataFednowTransfersNotSupported RoutingNumberListResponseDataFednowTransfers = "not_supported"
)

func (r RoutingNumberListResponseDataFednowTransfers) IsKnown() bool {
	switch r {
	case RoutingNumberListResponseDataFednowTransfersSupported, RoutingNumberListResponseDataFednowTransfersNotSupported:
		return true
	}
	return false
}

// This routing number's support for Real-Time Payments Transfers.
type RoutingNumberListResponseDataRealTimePaymentsTransfers string

const (
	RoutingNumberListResponseDataRealTimePaymentsTransfersSupported    RoutingNumberListResponseDataRealTimePaymentsTransfers = "supported"
	RoutingNumberListResponseDataRealTimePaymentsTransfersNotSupported RoutingNumberListResponseDataRealTimePaymentsTransfers = "not_supported"
)

func (r RoutingNumberListResponseDataRealTimePaymentsTransfers) IsKnown() bool {
	switch r {
	case RoutingNumberListResponseDataRealTimePaymentsTransfersSupported, RoutingNumberListResponseDataRealTimePaymentsTransfersNotSupported:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `routing_number`.
type RoutingNumberListResponseDataType string

const (
	RoutingNumberListResponseDataTypeRoutingNumber RoutingNumberListResponseDataType = "routing_number"
)

func (r RoutingNumberListResponseDataType) IsKnown() bool {
	switch r {
	case RoutingNumberListResponseDataTypeRoutingNumber:
		return true
	}
	return false
}

// This routing number's support for Wire Transfers.
type RoutingNumberListResponseDataWireTransfers string

const (
	RoutingNumberListResponseDataWireTransfersSupported    RoutingNumberListResponseDataWireTransfers = "supported"
	RoutingNumberListResponseDataWireTransfersNotSupported RoutingNumberListResponseDataWireTransfers = "not_supported"
)

func (r RoutingNumberListResponseDataWireTransfers) IsKnown() bool {
	switch r {
	case RoutingNumberListResponseDataWireTransfersSupported, RoutingNumberListResponseDataWireTransfersNotSupported:
		return true
	}
	return false
}

type RoutingNumberListParams struct {
	// Filter financial institutions by routing number.
	RoutingNumber param.Field[string] `query:"routing_number,required"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [RoutingNumberListParams]'s query parameters as
// `url.Values`.
func (r RoutingNumberListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

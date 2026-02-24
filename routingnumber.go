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
	"github.com/Increase/increase-go/packages/pagination"
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
func (r *RoutingNumberService) List(ctx context.Context, query RoutingNumberListParams, opts ...option.RequestOption) (res *pagination.Page[RoutingNumberListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "routing_numbers"
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

// You can use this API to confirm if a routing number is valid, such as when a
// user is providing you with bank account details. Since routing numbers uniquely
// identify a bank, this will always return 0 or 1 entry. In Sandbox, the only
// valid routing number for this method is 110000000.
func (r *RoutingNumberService) ListAutoPaging(ctx context.Context, query RoutingNumberListParams, opts ...option.RequestOption) *pagination.PageAutoPager[RoutingNumberListResponse] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Routing numbers are used to identify your bank in a financial transaction.
type RoutingNumberListResponse struct {
	// This routing number's support for ACH Transfers.
	ACHTransfers RoutingNumberListResponseACHTransfers `json:"ach_transfers" api:"required"`
	// This routing number's support for FedNow Transfers.
	FednowTransfers RoutingNumberListResponseFednowTransfers `json:"fednow_transfers" api:"required"`
	// The name of the financial institution belonging to a routing number.
	Name string `json:"name" api:"required"`
	// This routing number's support for Real-Time Payments Transfers.
	RealTimePaymentsTransfers RoutingNumberListResponseRealTimePaymentsTransfers `json:"real_time_payments_transfers" api:"required"`
	// The nine digit routing number identifier.
	RoutingNumber string `json:"routing_number" api:"required"`
	// A constant representing the object's type. For this resource it will always be
	// `routing_number`.
	Type RoutingNumberListResponseType `json:"type" api:"required"`
	// This routing number's support for Wire Transfers.
	WireTransfers RoutingNumberListResponseWireTransfers `json:"wire_transfers" api:"required"`
	JSON          routingNumberListResponseJSON          `json:"-"`
}

// routingNumberListResponseJSON contains the JSON metadata for the struct
// [RoutingNumberListResponse]
type routingNumberListResponseJSON struct {
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

func (r *RoutingNumberListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingNumberListResponseJSON) RawJSON() string {
	return r.raw
}

// This routing number's support for ACH Transfers.
type RoutingNumberListResponseACHTransfers string

const (
	RoutingNumberListResponseACHTransfersSupported    RoutingNumberListResponseACHTransfers = "supported"
	RoutingNumberListResponseACHTransfersNotSupported RoutingNumberListResponseACHTransfers = "not_supported"
)

func (r RoutingNumberListResponseACHTransfers) IsKnown() bool {
	switch r {
	case RoutingNumberListResponseACHTransfersSupported, RoutingNumberListResponseACHTransfersNotSupported:
		return true
	}
	return false
}

// This routing number's support for FedNow Transfers.
type RoutingNumberListResponseFednowTransfers string

const (
	RoutingNumberListResponseFednowTransfersSupported    RoutingNumberListResponseFednowTransfers = "supported"
	RoutingNumberListResponseFednowTransfersNotSupported RoutingNumberListResponseFednowTransfers = "not_supported"
)

func (r RoutingNumberListResponseFednowTransfers) IsKnown() bool {
	switch r {
	case RoutingNumberListResponseFednowTransfersSupported, RoutingNumberListResponseFednowTransfersNotSupported:
		return true
	}
	return false
}

// This routing number's support for Real-Time Payments Transfers.
type RoutingNumberListResponseRealTimePaymentsTransfers string

const (
	RoutingNumberListResponseRealTimePaymentsTransfersSupported    RoutingNumberListResponseRealTimePaymentsTransfers = "supported"
	RoutingNumberListResponseRealTimePaymentsTransfersNotSupported RoutingNumberListResponseRealTimePaymentsTransfers = "not_supported"
)

func (r RoutingNumberListResponseRealTimePaymentsTransfers) IsKnown() bool {
	switch r {
	case RoutingNumberListResponseRealTimePaymentsTransfersSupported, RoutingNumberListResponseRealTimePaymentsTransfersNotSupported:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `routing_number`.
type RoutingNumberListResponseType string

const (
	RoutingNumberListResponseTypeRoutingNumber RoutingNumberListResponseType = "routing_number"
)

func (r RoutingNumberListResponseType) IsKnown() bool {
	switch r {
	case RoutingNumberListResponseTypeRoutingNumber:
		return true
	}
	return false
}

// This routing number's support for Wire Transfers.
type RoutingNumberListResponseWireTransfers string

const (
	RoutingNumberListResponseWireTransfersSupported    RoutingNumberListResponseWireTransfers = "supported"
	RoutingNumberListResponseWireTransfersNotSupported RoutingNumberListResponseWireTransfers = "not_supported"
)

func (r RoutingNumberListResponseWireTransfers) IsKnown() bool {
	switch r {
	case RoutingNumberListResponseWireTransfersSupported, RoutingNumberListResponseWireTransfersNotSupported:
		return true
	}
	return false
}

type RoutingNumberListParams struct {
	// Filter financial institutions by routing number.
	RoutingNumber param.Field[string] `query:"routing_number" api:"required"`
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

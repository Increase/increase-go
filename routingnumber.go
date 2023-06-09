// File generated from our OpenAPI spec by Stainless.

package increase

import (
	"context"
	"net/http"
	"net/url"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/apiquery"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/internal/shared"
	"github.com/increase/increase-go/option"
)

// RoutingNumberService contains methods and other services that help with
// interacting with the increase API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRoutingNumberService] method
// instead.
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
func (r *RoutingNumberService) List(ctx context.Context, query RoutingNumberListParams, opts ...option.RequestOption) (res *shared.Page[RoutingNumber], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
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
func (r *RoutingNumberService) ListAutoPaging(ctx context.Context, query RoutingNumberListParams, opts ...option.RequestOption) *shared.PageAutoPager[RoutingNumber] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Routing numbers are used to identify your bank in a financial transaction.
type RoutingNumber struct {
	// The name of the financial institution belonging to a routing number.
	Name string `json:"name,required"`
	// The nine digit routing number identifier.
	RoutingNumber string `json:"routing_number,required"`
	// A constant representing the object's type. For this resource it will always be
	// `routing_number`.
	Type RoutingNumberType `json:"type,required"`
	// This routing number's support for ACH Transfers.
	ACHTransfers RoutingNumberACHTransfers `json:"ach_transfers,required"`
	// This routing number's support for Real Time Payments Transfers.
	RealTimePaymentsTransfers RoutingNumberRealTimePaymentsTransfers `json:"real_time_payments_transfers,required"`
	// This routing number's support for Wire Transfers.
	WireTransfers RoutingNumberWireTransfers `json:"wire_transfers,required"`
	JSON          routingNumberJSON
}

// routingNumberJSON contains the JSON metadata for the struct [RoutingNumber]
type routingNumberJSON struct {
	Name                      apijson.Field
	RoutingNumber             apijson.Field
	Type                      apijson.Field
	ACHTransfers              apijson.Field
	RealTimePaymentsTransfers apijson.Field
	WireTransfers             apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *RoutingNumber) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A constant representing the object's type. For this resource it will always be
// `routing_number`.
type RoutingNumberType string

const (
	RoutingNumberTypeRoutingNumber RoutingNumberType = "routing_number"
)

// This routing number's support for ACH Transfers.
type RoutingNumberACHTransfers string

const (
	RoutingNumberACHTransfersSupported    RoutingNumberACHTransfers = "supported"
	RoutingNumberACHTransfersNotSupported RoutingNumberACHTransfers = "not_supported"
)

// This routing number's support for Real Time Payments Transfers.
type RoutingNumberRealTimePaymentsTransfers string

const (
	RoutingNumberRealTimePaymentsTransfersSupported    RoutingNumberRealTimePaymentsTransfers = "supported"
	RoutingNumberRealTimePaymentsTransfersNotSupported RoutingNumberRealTimePaymentsTransfers = "not_supported"
)

// This routing number's support for Wire Transfers.
type RoutingNumberWireTransfers string

const (
	RoutingNumberWireTransfersSupported    RoutingNumberWireTransfers = "supported"
	RoutingNumberWireTransfersNotSupported RoutingNumberWireTransfers = "not_supported"
)

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

// A list of Routing Number objects
type RoutingNumberListResponse struct {
	// The contents of the list.
	Data []RoutingNumber `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       routingNumberListResponseJSON
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

package responses

import (
	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/pagination"
)

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
	JSON          RoutingNumberJSON
}

type RoutingNumberJSON struct {
	Name                      pjson.Metadata
	RoutingNumber             pjson.Metadata
	Type                      pjson.Metadata
	ACHTransfers              pjson.Metadata
	RealTimePaymentsTransfers pjson.Metadata
	WireTransfers             pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into RoutingNumber using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *RoutingNumber) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type RoutingNumberType string

const (
	RoutingNumberTypeRoutingNumber RoutingNumberType = "routing_number"
)

type RoutingNumberACHTransfers string

const (
	RoutingNumberACHTransfersSupported    RoutingNumberACHTransfers = "supported"
	RoutingNumberACHTransfersNotSupported RoutingNumberACHTransfers = "not_supported"
)

type RoutingNumberRealTimePaymentsTransfers string

const (
	RoutingNumberRealTimePaymentsTransfersSupported    RoutingNumberRealTimePaymentsTransfers = "supported"
	RoutingNumberRealTimePaymentsTransfersNotSupported RoutingNumberRealTimePaymentsTransfers = "not_supported"
)

type RoutingNumberWireTransfers string

const (
	RoutingNumberWireTransfersSupported    RoutingNumberWireTransfers = "supported"
	RoutingNumberWireTransfersNotSupported RoutingNumberWireTransfers = "not_supported"
)

type RoutingNumberList struct {
	// The contents of the list.
	Data []RoutingNumber `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       RoutingNumberListJSON
}

type RoutingNumberListJSON struct {
	Data       pjson.Metadata
	NextCursor pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into RoutingNumberList using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *RoutingNumberList) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type RoutingNumbersPage struct {
	*pagination.Page[RoutingNumber]
}

func (r *RoutingNumbersPage) RoutingNumber() *RoutingNumber {
	return r.Current()
}

func (r *RoutingNumbersPage) NextPage() (*RoutingNumbersPage, error) {
	if page, err := r.Page.NextPage(); err != nil {
		return nil, err
	} else {
		return &RoutingNumbersPage{page}, nil
	}
}

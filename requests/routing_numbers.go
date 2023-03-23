package requests

import (
	"fmt"
	"net/url"

	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/fields"
)

type RoutingNumber struct {
	// The name of the financial institution belonging to a routing number.
	Name fields.Field[string] `json:"name,required"`
	// The nine digit routing number identifier.
	RoutingNumber fields.Field[string] `json:"routing_number,required"`
	// A constant representing the object's type. For this resource it will always be
	// `routing_number`.
	Type fields.Field[RoutingNumberType] `json:"type,required"`
	// This routing number's support for ACH Transfers.
	ACHTransfers fields.Field[RoutingNumberACHTransfers] `json:"ach_transfers,required"`
	// This routing number's support for Real Time Payments Transfers.
	RealTimePaymentsTransfers fields.Field[RoutingNumberRealTimePaymentsTransfers] `json:"real_time_payments_transfers,required"`
	// This routing number's support for Wire Transfers.
	WireTransfers fields.Field[RoutingNumberWireTransfers] `json:"wire_transfers,required"`
}

// MarshalJSON serializes RoutingNumber into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *RoutingNumber) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r RoutingNumber) String() (result string) {
	return fmt.Sprintf("&RoutingNumber{Name:%s RoutingNumber:%s Type:%s ACHTransfers:%s RealTimePaymentsTransfers:%s WireTransfers:%s}", r.Name, r.RoutingNumber, r.Type, r.ACHTransfers, r.RealTimePaymentsTransfers, r.WireTransfers)
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

type RoutingNumberListParams struct {
	// Return the page of entries after this one.
	Cursor fields.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit fields.Field[int64] `query:"limit"`
	// Filter financial institutions by routing number.
	RoutingNumber fields.Field[string] `query:"routing_number,required"`
}

// URLQuery serializes RoutingNumberListParams into a url.Values of the query
// parameters associated with this value
func (r *RoutingNumberListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r RoutingNumberListParams) String() (result string) {
	return fmt.Sprintf("&RoutingNumberListParams{Cursor:%s Limit:%s RoutingNumber:%s}", r.Cursor, r.Limit, r.RoutingNumber)
}

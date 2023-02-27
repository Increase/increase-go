package types

import (
	"fmt"
	"increase/core"
	"increase/core/pjson"
	"increase/core/query"
	"increase/pagination"
	"net/url"
)

type RoutingNumber struct {
	// The name of the financial institution belonging to a routing number.
	Name *string `pjson:"name"`
	// The nine digit routing number identifier.
	RoutingNumber *string `pjson:"routing_number"`
	// A constant representing the object's type. For this resource it will always be
	// `routing_number`.
	Type *RoutingNumberType `pjson:"type"`
	// This routing number's support for ACH Transfers.
	ACHTransfers *RoutingNumberACHTransfers `pjson:"ach_transfers"`
	// This routing number's support for Real Time Payments Transfers.
	RealTimePaymentsTransfers *RoutingNumberRealTimePaymentsTransfers `pjson:"real_time_payments_transfers"`
	// This routing number's support for Wire Transfers.
	WireTransfers *RoutingNumberWireTransfers `pjson:"wire_transfers"`
	jsonFields    map[string]interface{}      `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into RoutingNumber using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *RoutingNumber) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes RoutingNumber into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *RoutingNumber) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The name of the financial institution belonging to a routing number.
func (r *RoutingNumber) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}

// The nine digit routing number identifier.
func (r *RoutingNumber) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `routing_number`.
func (r *RoutingNumber) GetType() (Type RoutingNumberType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

// This routing number's support for ACH Transfers.
func (r *RoutingNumber) GetACHTransfers() (ACHTransfers RoutingNumberACHTransfers) {
	if r != nil && r.ACHTransfers != nil {
		ACHTransfers = *r.ACHTransfers
	}
	return
}

// This routing number's support for Real Time Payments Transfers.
func (r *RoutingNumber) GetRealTimePaymentsTransfers() (RealTimePaymentsTransfers RoutingNumberRealTimePaymentsTransfers) {
	if r != nil && r.RealTimePaymentsTransfers != nil {
		RealTimePaymentsTransfers = *r.RealTimePaymentsTransfers
	}
	return
}

// This routing number's support for Wire Transfers.
func (r *RoutingNumber) GetWireTransfers() (WireTransfers RoutingNumberWireTransfers) {
	if r != nil && r.WireTransfers != nil {
		WireTransfers = *r.WireTransfers
	}
	return
}

func (r RoutingNumber) String() (result string) {
	return fmt.Sprintf("&RoutingNumber{Name:%s RoutingNumber:%s Type:%s ACHTransfers:%s RealTimePaymentsTransfers:%s WireTransfers:%s}", core.FmtP(r.Name), core.FmtP(r.RoutingNumber), core.FmtP(r.Type), core.FmtP(r.ACHTransfers), core.FmtP(r.RealTimePaymentsTransfers), core.FmtP(r.WireTransfers))
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
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int64 `query:"limit"`
	// Filter financial institutions by routing number.
	RoutingNumber *string                `query:"routing_number,omitempty"`
	jsonFields    map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into RoutingNumberListParams using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *RoutingNumberListParams) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes RoutingNumberListParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *RoutingNumberListParams) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes RoutingNumberListParams into a url.Values of the query
// parameters associated with this value
func (r *RoutingNumberListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// Return the page of entries after this one.
func (r *RoutingNumberListParams) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *RoutingNumberListParams) GetLimit() (Limit int64) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// Filter financial institutions by routing number.
func (r *RoutingNumberListParams) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

func (r RoutingNumberListParams) String() (result string) {
	return fmt.Sprintf("&RoutingNumberListParams{Cursor:%s Limit:%s RoutingNumber:%s}", core.FmtP(r.Cursor), core.FmtP(r.Limit), core.FmtP(r.RoutingNumber))
}

type RoutingNumberList struct {
	// The contents of the list.
	Data *[]RoutingNumber `pjson:"data"`
	// A pointer to a place in the list.
	NextCursor *string                `pjson:"next_cursor"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into RoutingNumberList using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *RoutingNumberList) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes RoutingNumberList into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *RoutingNumberList) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes RoutingNumberList into a url.Values of the query parameters
// associated with this value
func (r *RoutingNumberList) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// The contents of the list.
func (r *RoutingNumberList) GetData() (Data []RoutingNumber) {
	if r != nil && r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r *RoutingNumberList) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

func (r RoutingNumberList) String() (result string) {
	return fmt.Sprintf("&RoutingNumberList{Data:%s NextCursor:%s}", core.Fmt(r.Data), core.FmtP(r.NextCursor))
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

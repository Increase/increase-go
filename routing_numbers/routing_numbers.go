package routing_numbers

import "context"
import "increase/core"
import "increase/pagination"

type RoutingNumberService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewRoutingNumberService(requester core.Requester) (r *RoutingNumberService) {
	r = &RoutingNumberService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

//
type RoutingNumber struct {
	// The name of the financial institution belonging to a routing number.
	Name string `json:"name"`
	// The nine digit routing number identifier.
	RoutingNumber string `json:"routing_number"`
	// A constant representing the object's type. For this resource it will always be
	// `routing_number`.
	Type RoutingNumberType `json:"type"`
	// This routing number's support for ACH Transfers.
	ACHTransfers RoutingNumberACHTransfers `json:"ach_transfers"`
	// This routing number's support for Real Time Payments Transfers.
	RealTimePaymentsTransfers RoutingNumberRealTimePaymentsTransfers `json:"real_time_payments_transfers"`
	// This routing number's support for Wire Transfers.
	WireTransfers RoutingNumberWireTransfers `json:"wire_transfers"`
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

type ListRoutingNumbersQuery struct {
	// Return the page of entries after this one.
	Cursor string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit int `query:"limit"`
	// Filter financial institutions by routing number.
	RoutingNumber string `query:"routing_number,omitempty"`
}

//
type RoutingNumberList struct {
	// The contents of the list.
	Data []RoutingNumber `json:"data"`
	// A pointer to a place in the list.
	NextCursor *string `json:"next_cursor"`
}

// A pointer to a place in the list.
func (r *RoutingNumberList) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

type RoutingNumbersPage struct {
	*pagination.Page[RoutingNumber]
}

func (r *RoutingNumbersPage) RoutingNumber() *RoutingNumber {
	return r.Current()
}

func (r *RoutingNumbersPage) GetNextPage() (*RoutingNumbersPage, error) {
	if page, err := r.Page.GetNextPage(); err != nil {
		return nil, err
	} else {
		return &RoutingNumbersPage{page}, nil
	}
}

// You can use this API to confirm if a routing number is valid, such as when a
// user is providing you with bank account details. Since routing numbers uniquely
// identify a bank, this will always return 0 or 1 entry. In Sandbox, the only
// valid routing number for this method is 110000000.
func (r *RoutingNumberService) List(ctx context.Context, query *ListRoutingNumbersQuery, opts ...*core.RequestOpts) (res *RoutingNumbersPage, err error) {
	page := &RoutingNumbersPage{
		Page: &pagination.Page[RoutingNumber]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/routing_numbers",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

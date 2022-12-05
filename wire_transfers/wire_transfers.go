package wire_transfers

import "increase/core"
import "fmt"

type WireTransfers struct {
	Requester core.Requester
	get       func(string, *core.CoreRequest, interface{}) error
	post      func(string, *core.CoreRequest, interface{}) error
	patch     func(string, *core.CoreRequest, interface{}) error
	put       func(string, *core.CoreRequest, interface{}) error
	delete    func(string, *core.CoreRequest, interface{}) error
}

func NewWireTransfers(requster core.Requester) (r *WireTransfers) {
	r = &WireTransfers{}
	r.Requester = requster
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

func (r *WireTransfers) Create(body WireTransfersCreateParams, opts ...*core.RequestOpts) (res WireTransfer, err error) {
	err = r.post(
		"/wire_transfers",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *WireTransfers) Retrieve(wire_transfer_id string, opts ...*core.RequestOpts) (res WireTransfer, err error) {
	err = r.get(
		fmt.Sprintf("/wire_transfers/%s", wire_transfer_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *WireTransfers) List(query *WireTransfersListQuery, opts ...*core.RequestOpts) (res WireTransfersListResponse, err error) {
	err = r.get(
		"/wire_transfers",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Query:  query,
		},
		&res,
	)
	return
}

// Simulates the reversal of an Wire Transfer by the Federal Reserve due to error
// conditions. This will also create a Transaction to account for the returned
// funds. This transfer must first have a `status` of `complete`.
func (r *WireTransfers) Reverse(wire_transfer_id string, opts ...*core.RequestOpts) (res WireTransfer, err error) {
	err = r.post(
		fmt.Sprintf("/simulations/wire_transfers/%s/reverse", wire_transfer_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

// Simulates the submission of a Wire Transfer to the Federal Reserve. This
// transfer must first have a `status` of `pending_approval` or `pending_creating`.
func (r *WireTransfers) Submit(wire_transfer_id string, opts ...*core.RequestOpts) (res WireTransfer, err error) {
	err = r.post(
		fmt.Sprintf("/simulations/wire_transfers/%s/submit", wire_transfer_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

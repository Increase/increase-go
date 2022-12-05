package simulations

import "increase/core"
import "increase/ach_transfers"
import "fmt"

type ACHTransfers struct {
	Requester core.Requester
	get       func(string, *core.CoreRequest, interface{}) error
	post      func(string, *core.CoreRequest, interface{}) error
	patch     func(string, *core.CoreRequest, interface{}) error
	put       func(string, *core.CoreRequest, interface{}) error
	delete    func(string, *core.CoreRequest, interface{}) error
}

func NewACHTransfers(requster core.Requester) (r *ACHTransfers) {
	r = &ACHTransfers{}
	r.Requester = requster
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

// Simulates an inbound ACH transfer to your account. The transfer may be either a
// credit or a debit depending on if the `amount` is positive or negative. This
// will result in either a Transaction or a Declined Transaction depending on if
// the transfer is allowed.
func (r *ACHTransfers) CreateInbound(body ACHTransfersCreateInboundParams, opts ...*core.RequestOpts) (res ACHTransferSimulation, err error) {
	err = r.post(
		"/simulations/inbound_ach_transfers",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

// Simulates the return of an ACH Transfer by the Federal Reserve due to error
// conditions. This will also create a Transaction to account for the returned
// funds. This transfer must first have a `status` of `submitted`.
func (r *ACHTransfers) Return(ach_transfer_id string, opts ...*core.RequestOpts) (res ach_transfers.ACHTransfer, err error) {
	err = r.post(
		fmt.Sprintf("/simulations/ach_transfers/%s/return", ach_transfer_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

// Simulates the submission of an ACH Transfer to the Federal Reserve. This
// transfer must first have a `status` of `pending_approval` or
// `pending_submission`.
func (r *ACHTransfers) Submit(ach_transfer_id string, opts ...*core.RequestOpts) (res ach_transfers.ACHTransfer, err error) {
	err = r.post(
		fmt.Sprintf("/simulations/ach_transfers/%s/submit", ach_transfer_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

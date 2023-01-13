package simulations

import "context"
import "increase/core"
import "increase/check_transfers"
import "fmt"

type CheckTransferService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewCheckTransferService(requester core.Requester) (r *CheckTransferService) {
	r = &CheckTransferService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

type PreloadedCheckTransferService struct {
	CheckTransfers *CheckTransferService
}

func (r *PreloadedCheckTransferService) Init(service *CheckTransferService) {
	r.CheckTransfers = service
}

func NewPreloadedCheckTransferService(service *CheckTransferService) (r *PreloadedCheckTransferService) {
	r = &PreloadedCheckTransferService{}
	r.Init(service)
	return
}

// Simulates the mailing of a Check Transfer. This transfer must first have a
// `status` of `pending_approval` or `pending_submission`.
func (r *CheckTransferService) Mail(ctx context.Context, check_transfer_id string, opts ...*core.RequestOpts) (res *check_transfers.CheckTransfer, err error) {
	err = r.post(
		ctx,
		fmt.Sprintf("/simulations/check_transfers/%s/mail", check_transfer_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

// Simulates the mailing of a Check Transfer. This transfer must first have a
// `status` of `pending_approval` or `pending_submission`.
func (r *PreloadedCheckTransferService) Mail(ctx context.Context, check_transfer_id string, opts ...*core.RequestOpts) (res *check_transfers.CheckTransfer, err error) {
	err = r.CheckTransfers.post(
		ctx,
		fmt.Sprintf("/simulations/check_transfers/%s/mail", check_transfer_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

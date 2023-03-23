package services

import (
	"context"
	"fmt"

	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/responses"
)

type SimulationsAccountTransferService struct {
	Options []options.RequestOption
}

func NewSimulationsAccountTransferService(opts ...options.RequestOption) (r *SimulationsAccountTransferService) {
	r = &SimulationsAccountTransferService{}
	r.Options = opts
	return
}

// If your account is configured to require approval for each transfer, this
// endpoint simulates the approval of an [Account Transfer](#account-transfers).
// You can also approve sandbox Account Transfers in the dashboard. This transfer
// must first have a `status` of `pending_approval`.
func (r *SimulationsAccountTransferService) Complete(ctx context.Context, account_transfer_id string, opts ...options.RequestOption) (res *responses.AccountTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("simulations/account_transfers/%s/complete", account_transfer_id)
	err = options.ExecuteNewRequest(ctx, "POST", path, nil, &res, opts...)
	return
}

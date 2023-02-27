package services

import (
	"context"
	"fmt"
	"increase/options"
	"increase/types"
	"net/url"
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
func (r *SimulationsAccountTransferService) Complete(ctx context.Context, account_transfer_id string, opts ...options.RequestOption) (res *types.AccountTransfer, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]options.RequestOption{options.WithHeader("Content-Type", "")}, opts...)
	u, err := url.Parse(fmt.Sprintf("simulations/account_transfers/%s/complete", account_transfer_id))
	if err != nil {
		return
	}
	cfg, err := options.NewRequestConfig(ctx, "POST", u, nil, opts...)
	if err != nil {
		return
	}
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return
	}

	return
}

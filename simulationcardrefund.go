package increase

import (
	"context"
	"net/http"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/field"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/option"
)

type SimulationCardRefundService struct {
	Options []option.RequestOption
}

func NewSimulationCardRefundService(opts ...option.RequestOption) (r *SimulationCardRefundService) {
	r = &SimulationCardRefundService{}
	r.Options = opts
	return
}

// Simulates refunding a card transaction. The full value of the original sandbox
// transaction is refunded.
func (r *SimulationCardRefundService) New(ctx context.Context, body SimulationCardRefundNewParams, opts ...option.RequestOption) (res *Transaction, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/card_refunds"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SimulationCardRefundNewParams struct {
	// The identifier for the Transaction to refund. The Transaction's source must have
	// a category of card_settlement.
	TransactionID field.Field[string] `json:"transaction_id,required"`
}

// MarshalJSON serializes SimulationCardRefundNewParams into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r SimulationCardRefundNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

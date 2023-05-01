package increase

import (
	"context"
	"net/http"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/field"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/option"
)

type SimulationDocumentService struct {
	Options []option.RequestOption
}

func NewSimulationDocumentService(opts ...option.RequestOption) (r *SimulationDocumentService) {
	r = &SimulationDocumentService{}
	r.Options = opts
	return
}

// Simulates an tax document being created for an account.
func (r *SimulationDocumentService) New(ctx context.Context, body SimulationDocumentNewParams, opts ...option.RequestOption) (res *Document, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/documents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SimulationDocumentNewParams struct {
	// The identifier of the Account the tax document is for.
	AccountID field.Field[string] `json:"account_id,required"`
}

// MarshalJSON serializes SimulationDocumentNewParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r SimulationDocumentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

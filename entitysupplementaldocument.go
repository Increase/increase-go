package increase

import (
	"context"
	"fmt"
	"net/http"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/option"
)

// EntitySupplementalDocumentService contains methods and other services that help
// with interacting with the increase API. Note, unlike clients, this service does
// not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewEntitySupplementalDocumentService] method instead.
type EntitySupplementalDocumentService struct {
	Options []option.RequestOption
}

// NewEntitySupplementalDocumentService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewEntitySupplementalDocumentService(opts ...option.RequestOption) (r *EntitySupplementalDocumentService) {
	r = &EntitySupplementalDocumentService{}
	r.Options = opts
	return
}

// Create a supplemental document for an Entity
func (r *EntitySupplementalDocumentService) New(ctx context.Context, entityID string, body EntitySupplementalDocumentNewParams, opts ...option.RequestOption) (res *Entity, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("entities/%s/supplemental_documents", entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type EntitySupplementalDocumentNewParams struct {
	// The identifier of the File containing the document.
	FileID param.Field[string] `json:"file_id,required"`
}

func (r EntitySupplementalDocumentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

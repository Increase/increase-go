package increase

import (
	"context"
	"fmt"
	"net/http"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/field"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/option"
)

type EntitySupplementalDocumentService struct {
	Options []option.RequestOption
}

func NewEntitySupplementalDocumentService(opts ...option.RequestOption) (r *EntitySupplementalDocumentService) {
	r = &EntitySupplementalDocumentService{}
	r.Options = opts
	return
}

// Create a supplemental document for an Entity
func (r *EntitySupplementalDocumentService) New(ctx context.Context, entity_id string, body EntitySupplementalDocumentNewParams, opts ...option.RequestOption) (res *Entity, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("entities/%s/supplemental_documents", entity_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type EntitySupplementalDocumentNewParams struct {
	// The identifier of the File containing the document.
	FileID field.Field[string] `json:"file_id,required"`
}

// MarshalJSON serializes EntitySupplementalDocumentNewParams into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r EntitySupplementalDocumentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

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

// EntityIndustryCodeService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEntityIndustryCodeService] method instead.
type EntityIndustryCodeService struct {
	Options []option.RequestOption
}

// NewEntityIndustryCodeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEntityIndustryCodeService(opts ...option.RequestOption) (r *EntityIndustryCodeService) {
	r = &EntityIndustryCodeService{}
	r.Options = opts
	return
}

// Update the industry code for a corporate Entity
func (r *EntityIndustryCodeService) New(ctx context.Context, entityID string, body EntityIndustryCodeNewParams, opts ...option.RequestOption) (res *Entity, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("entities/%s/industry_code", entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type EntityIndustryCodeNewParams struct {
	// The North American Industry Classification System (NAICS) code for the
	// corporation's primary line of business. This is a number, like `5132` for
	// `Software Publishers`. A full list of classification codes is available
	// [here](https://increase.com/documentation/data-dictionary#north-american-industry-classification-system-codes).
	IndustryCode param.Field[string] `json:"industry_code,required"`
}

func (r EntityIndustryCodeNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

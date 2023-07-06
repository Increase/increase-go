// File generated from our OpenAPI spec by Stainless.

package increase

import (
	"context"
	"fmt"
	"net/http"

	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/option"
)

// SimulationCardProfileService contains methods and other services that help with
// interacting with the increase API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSimulationCardProfileService]
// method instead.
type SimulationCardProfileService struct {
	Options []option.RequestOption
}

// NewSimulationCardProfileService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSimulationCardProfileService(opts ...option.RequestOption) (r *SimulationCardProfileService) {
	r = &SimulationCardProfileService{}
	r.Options = opts
	return
}

// After a [Card Profile](#card-profile) is created in production, the images will
// be uploaded to Visa within one day. Since Card Profiles are not uploaded to Visa
// in sandbox, this endpoint simulates that step. Invoking this simulation triggers
// the webhooks Increase sends when the Card Profile is approved and updates the
// status of the Card Profile.
func (r *SimulationCardProfileService) Approve(ctx context.Context, cardProfileID string, opts ...option.RequestOption) (res *CardProfile, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("simulations/card_profiles/%s/approve", cardProfileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

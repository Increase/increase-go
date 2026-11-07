// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
)

// SimulationCheckbookService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimulationCheckbookService] method instead.
type SimulationCheckbookService struct {
	Options []option.RequestOption
}

// NewSimulationCheckbookService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSimulationCheckbookService(opts ...option.RequestOption) (r *SimulationCheckbookService) {
	r = &SimulationCheckbookService{}
	r.Options = opts
	return
}

// Simulates the mailing of a [Checkbook](#checkbooks), which happens periodically
// throughout the day in production but can be sped up in sandbox. This Checkbook
// must have a `status` of `pending_submission` or `pending_mailing`.
func (r *SimulationCheckbookService) Mail(ctx context.Context, checkbookID string, opts ...option.RequestOption) (res *Checkbook, err error) {
	opts = slices.Concat(r.Options, opts)
	if checkbookID == "" {
		err = errors.New("missing required checkbook_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("simulations/checkbooks/%s/mail", checkbookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

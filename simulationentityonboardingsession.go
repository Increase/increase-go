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

// SimulationEntityOnboardingSessionService contains methods and other services
// that help with interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimulationEntityOnboardingSessionService] method instead.
type SimulationEntityOnboardingSessionService struct {
	Options []option.RequestOption
}

// NewSimulationEntityOnboardingSessionService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSimulationEntityOnboardingSessionService(opts ...option.RequestOption) (r *SimulationEntityOnboardingSessionService) {
	r = &SimulationEntityOnboardingSessionService{}
	r.Options = opts
	return
}

// Simulates the submission of an
// [Entity Onboarding Session](#entity-onboarding-sessions). This session must have
// a `status` of `active`. After submission, the session will transition to
// `expired` and a new Entity will be created.
func (r *SimulationEntityOnboardingSessionService) Submit(ctx context.Context, entityOnboardingSessionID string, opts ...option.RequestOption) (res *EntityOnboardingSession, err error) {
	opts = slices.Concat(r.Options, opts)
	if entityOnboardingSessionID == "" {
		err = errors.New("missing required entity_onboarding_session_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("simulations/entity_onboarding_sessions/%s/submit", entityOnboardingSessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/Increase/increase-go/internal/apijson"
	"github.com/Increase/increase-go/internal/param"
	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
)

// SimulationEntityService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimulationEntityService] method instead.
type SimulationEntityService struct {
	Options []option.RequestOption
}

// NewSimulationEntityService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSimulationEntityService(opts ...option.RequestOption) (r *SimulationEntityService) {
	r = &SimulationEntityService{}
	r.Options = opts
	return
}

// Simulate updates to an
// [Entity's validation](/documentation/api/entities#entity-object.validation). In
// production, Know Your Customer validations
// [run automatically](/documentation/entity-validation#entity-validation) for
// eligible programs. While developing, use this API to simulate issues with
// information submissions.
func (r *SimulationEntityService) UpdateValidation(ctx context.Context, entityID string, body SimulationEntityUpdateValidationParams, opts ...option.RequestOption) (res *Entity, err error) {
	opts = slices.Concat(r.Options, opts)
	if entityID == "" {
		err = errors.New("missing required entity_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("simulations/entities/%s/update_validation", entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type SimulationEntityUpdateValidationParams struct {
	// The validation issues to attach. If no issues are provided, the validation
	// status will be set to `valid`.
	Issues param.Field[[]SimulationEntityUpdateValidationParamsIssue] `json:"issues" api:"required"`
}

func (r SimulationEntityUpdateValidationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SimulationEntityUpdateValidationParamsIssue struct {
	// The type of issue.
	Category param.Field[SimulationEntityUpdateValidationParamsIssuesCategory] `json:"category" api:"required"`
}

func (r SimulationEntityUpdateValidationParamsIssue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of issue.
type SimulationEntityUpdateValidationParamsIssuesCategory string

const (
	SimulationEntityUpdateValidationParamsIssuesCategoryEntityTaxIdentifier     SimulationEntityUpdateValidationParamsIssuesCategory = "entity_tax_identifier"
	SimulationEntityUpdateValidationParamsIssuesCategoryEntityAddress           SimulationEntityUpdateValidationParamsIssuesCategory = "entity_address"
	SimulationEntityUpdateValidationParamsIssuesCategoryBeneficialOwnerIdentity SimulationEntityUpdateValidationParamsIssuesCategory = "beneficial_owner_identity"
	SimulationEntityUpdateValidationParamsIssuesCategoryBeneficialOwnerAddress  SimulationEntityUpdateValidationParamsIssuesCategory = "beneficial_owner_address"
)

func (r SimulationEntityUpdateValidationParamsIssuesCategory) IsKnown() bool {
	switch r {
	case SimulationEntityUpdateValidationParamsIssuesCategoryEntityTaxIdentifier, SimulationEntityUpdateValidationParamsIssuesCategoryEntityAddress, SimulationEntityUpdateValidationParamsIssuesCategoryBeneficialOwnerIdentity, SimulationEntityUpdateValidationParamsIssuesCategoryBeneficialOwnerAddress:
		return true
	}
	return false
}

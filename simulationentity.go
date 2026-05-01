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

// Set the status for an
// [Entity's validation](/documentation/api/entities#entity-object.validation). In
// production, Know Your Customer validations
// [run automatically](/documentation/entity-validation#entity-validation). While
// developing, it can be helpful to override the behavior in Sandbox.
func (r *SimulationEntityService) Validation(ctx context.Context, entityID string, body SimulationEntityValidationParams, opts ...option.RequestOption) (res *Entity, err error) {
	opts = slices.Concat(r.Options, opts)
	if entityID == "" {
		err = errors.New("missing required entity_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("simulations/entities/%s/validation", entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type SimulationEntityValidationParams struct {
	// The validation issues to attach. Only allowed when `status` is `invalid`.
	Issues param.Field[[]SimulationEntityValidationParamsIssue] `json:"issues" api:"required"`
	// The validation status to set on the Entity.
	Status param.Field[SimulationEntityValidationParamsStatus] `json:"status" api:"required"`
}

func (r SimulationEntityValidationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SimulationEntityValidationParamsIssue struct {
	// The type of issue.
	Category param.Field[SimulationEntityValidationParamsIssuesCategory] `json:"category" api:"required"`
}

func (r SimulationEntityValidationParamsIssue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of issue.
type SimulationEntityValidationParamsIssuesCategory string

const (
	SimulationEntityValidationParamsIssuesCategoryEntityTaxIdentifier     SimulationEntityValidationParamsIssuesCategory = "entity_tax_identifier"
	SimulationEntityValidationParamsIssuesCategoryEntityAddress           SimulationEntityValidationParamsIssuesCategory = "entity_address"
	SimulationEntityValidationParamsIssuesCategoryBeneficialOwnerIdentity SimulationEntityValidationParamsIssuesCategory = "beneficial_owner_identity"
	SimulationEntityValidationParamsIssuesCategoryBeneficialOwnerAddress  SimulationEntityValidationParamsIssuesCategory = "beneficial_owner_address"
)

func (r SimulationEntityValidationParamsIssuesCategory) IsKnown() bool {
	switch r {
	case SimulationEntityValidationParamsIssuesCategoryEntityTaxIdentifier, SimulationEntityValidationParamsIssuesCategoryEntityAddress, SimulationEntityValidationParamsIssuesCategoryBeneficialOwnerIdentity, SimulationEntityValidationParamsIssuesCategoryBeneficialOwnerAddress:
		return true
	}
	return false
}

// The validation status to set on the Entity.
type SimulationEntityValidationParamsStatus string

const (
	SimulationEntityValidationParamsStatusValid   SimulationEntityValidationParamsStatus = "valid"
	SimulationEntityValidationParamsStatusInvalid SimulationEntityValidationParamsStatus = "invalid"
	SimulationEntityValidationParamsStatusPending SimulationEntityValidationParamsStatus = "pending"
)

func (r SimulationEntityValidationParamsStatus) IsKnown() bool {
	switch r {
	case SimulationEntityValidationParamsStatusValid, SimulationEntityValidationParamsStatusInvalid, SimulationEntityValidationParamsStatusPending:
		return true
	}
	return false
}

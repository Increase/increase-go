// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Increase/increase-go/internal/apijson"
	"github.com/Increase/increase-go/internal/apiquery"
	"github.com/Increase/increase-go/internal/param"
	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
	"github.com/Increase/increase-go/packages/pagination"
)

// EntityOnboardingSessionService contains methods and other services that help
// with interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEntityOnboardingSessionService] method instead.
type EntityOnboardingSessionService struct {
	Options []option.RequestOption
}

// NewEntityOnboardingSessionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEntityOnboardingSessionService(opts ...option.RequestOption) (r *EntityOnboardingSessionService) {
	r = &EntityOnboardingSessionService{}
	r.Options = opts
	return
}

// Create an Entity Onboarding Session
func (r *EntityOnboardingSessionService) New(ctx context.Context, body EntityOnboardingSessionNewParams, opts ...option.RequestOption) (res *EntityOnboardingSession, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "entity_onboarding_sessions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve an Entity Onboarding Session
func (r *EntityOnboardingSessionService) Get(ctx context.Context, entityOnboardingSessionID string, opts ...option.RequestOption) (res *EntityOnboardingSession, err error) {
	opts = slices.Concat(r.Options, opts)
	if entityOnboardingSessionID == "" {
		err = errors.New("missing required entity_onboarding_session_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("entity_onboarding_sessions/%s", entityOnboardingSessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List Entity Onboarding Session
func (r *EntityOnboardingSessionService) List(ctx context.Context, query EntityOnboardingSessionListParams, opts ...option.RequestOption) (res *pagination.Page[EntityOnboardingSession], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "entity_onboarding_sessions"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List Entity Onboarding Session
func (r *EntityOnboardingSessionService) ListAutoPaging(ctx context.Context, query EntityOnboardingSessionListParams, opts ...option.RequestOption) *pagination.PageAutoPager[EntityOnboardingSession] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Expire an Entity Onboarding Session
func (r *EntityOnboardingSessionService) Expire(ctx context.Context, entityOnboardingSessionID string, opts ...option.RequestOption) (res *EntityOnboardingSession, err error) {
	opts = slices.Concat(r.Options, opts)
	if entityOnboardingSessionID == "" {
		err = errors.New("missing required entity_onboarding_session_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("entity_onboarding_sessions/%s/expire", entityOnboardingSessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Entity Onboarding Sessions let your customers onboard themselves by completing
// Increase-hosted forms. Create a session and redirect your customer to the
// returned URL. When they're done, they'll be redirected back to your site. This
// API is used for [hosted onboarding](/documentation/hosted-onboarding).
type EntityOnboardingSession struct {
	// The Entity Onboarding Session's identifier.
	ID string `json:"id" api:"required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Entity Onboarding Session was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The identifier of the Entity associated with this session, if one has been
	// created or was provided when creating the session.
	EntityID string `json:"entity_id" api:"required,nullable"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Entity Onboarding Session will expire.
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// The idempotency key you chose for this object. This value is unique across
	// Increase and is used to ensure that a request is only processed once. Learn more
	// about [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey string `json:"idempotency_key" api:"required,nullable"`
	// The identifier of the Program the Entity will be onboarded to.
	ProgramID string `json:"program_id" api:"required"`
	// The URL to redirect to after the onboarding session is complete. Increase will
	// include the query parameters `entity_onboarding_session_id` and `entity_id` when
	// redirecting.
	RedirectURL string `json:"redirect_url" api:"required"`
	// The URL containing the onboarding form. You should share this link with your
	// customer. Only present when the session is active.
	SessionURL string `json:"session_url" api:"required,nullable"`
	// The status of the onboarding session.
	Status EntityOnboardingSessionStatus `json:"status" api:"required"`
	// A constant representing the object's type. For this resource it will always be
	// `entity_onboarding_session`.
	Type EntityOnboardingSessionType `json:"type" api:"required"`
	JSON entityOnboardingSessionJSON `json:"-"`
}

// entityOnboardingSessionJSON contains the JSON metadata for the struct
// [EntityOnboardingSession]
type entityOnboardingSessionJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	EntityID       apijson.Field
	ExpiresAt      apijson.Field
	IdempotencyKey apijson.Field
	ProgramID      apijson.Field
	RedirectURL    apijson.Field
	SessionURL     apijson.Field
	Status         apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EntityOnboardingSession) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityOnboardingSessionJSON) RawJSON() string {
	return r.raw
}

// The status of the onboarding session.
type EntityOnboardingSessionStatus string

const (
	EntityOnboardingSessionStatusActive  EntityOnboardingSessionStatus = "active"
	EntityOnboardingSessionStatusExpired EntityOnboardingSessionStatus = "expired"
)

func (r EntityOnboardingSessionStatus) IsKnown() bool {
	switch r {
	case EntityOnboardingSessionStatusActive, EntityOnboardingSessionStatusExpired:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `entity_onboarding_session`.
type EntityOnboardingSessionType string

const (
	EntityOnboardingSessionTypeEntityOnboardingSession EntityOnboardingSessionType = "entity_onboarding_session"
)

func (r EntityOnboardingSessionType) IsKnown() bool {
	switch r {
	case EntityOnboardingSessionTypeEntityOnboardingSession:
		return true
	}
	return false
}

type EntityOnboardingSessionNewParams struct {
	// The identifier of the Program the Entity will be onboarded to.
	ProgramID param.Field[string] `json:"program_id" api:"required"`
	// The URL to redirect the customer to after they complete the onboarding form. The
	// redirect will include `entity_onboarding_session_id` and `entity_id` query
	// parameters.
	RedirectURL param.Field[string] `json:"redirect_url" api:"required"`
	// The identifier of an existing Entity to associate with the onboarding session.
	// If provided, the onboarding form will display any outstanding tasks required to
	// complete the Entity's onboarding.
	EntityID param.Field[string] `json:"entity_id"`
}

func (r EntityOnboardingSessionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EntityOnboardingSessionListParams struct {
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Filter records to the one with the specified `idempotency_key` you chose for
	// that object. This value is unique across Increase and is used to ensure that a
	// request is only processed once. Learn more about
	// [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey param.Field[string] `query:"idempotency_key"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit  param.Field[int64]                                   `query:"limit"`
	Status param.Field[EntityOnboardingSessionListParamsStatus] `query:"status"`
}

// URLQuery serializes [EntityOnboardingSessionListParams]'s query parameters as
// `url.Values`.
func (r EntityOnboardingSessionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type EntityOnboardingSessionListParamsStatus struct {
	// Filter Entity Onboarding Session for those with the specified status or
	// statuses. For GET requests, this should be encoded as a comma-delimited string,
	// such as `?in=one,two,three`.
	In param.Field[[]EntityOnboardingSessionListParamsStatusIn] `query:"in"`
}

// URLQuery serializes [EntityOnboardingSessionListParamsStatus]'s query parameters
// as `url.Values`.
func (r EntityOnboardingSessionListParamsStatus) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type EntityOnboardingSessionListParamsStatusIn string

const (
	EntityOnboardingSessionListParamsStatusInActive  EntityOnboardingSessionListParamsStatusIn = "active"
	EntityOnboardingSessionListParamsStatusInExpired EntityOnboardingSessionListParamsStatusIn = "expired"
)

func (r EntityOnboardingSessionListParamsStatusIn) IsKnown() bool {
	switch r {
	case EntityOnboardingSessionListParamsStatusInActive, EntityOnboardingSessionListParamsStatusInExpired:
		return true
	}
	return false
}

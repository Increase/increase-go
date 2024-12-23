// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Increase/increase-go/internal/apijson"
	"github.com/Increase/increase-go/internal/apiquery"
	"github.com/Increase/increase-go/internal/param"
	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
	"github.com/Increase/increase-go/packages/pagination"
)

// OAuthApplicationService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOAuthApplicationService] method instead.
type OAuthApplicationService struct {
	Options []option.RequestOption
}

// NewOAuthApplicationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOAuthApplicationService(opts ...option.RequestOption) (r *OAuthApplicationService) {
	r = &OAuthApplicationService{}
	r.Options = opts
	return
}

// Retrieve an OAuth Application
func (r *OAuthApplicationService) Get(ctx context.Context, oauthApplicationID string, opts ...option.RequestOption) (res *OAuthApplication, err error) {
	opts = append(r.Options[:], opts...)
	if oauthApplicationID == "" {
		err = errors.New("missing required oauth_application_id parameter")
		return
	}
	path := fmt.Sprintf("oauth_applications/%s", oauthApplicationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List OAuth Applications
func (r *OAuthApplicationService) List(ctx context.Context, query OAuthApplicationListParams, opts ...option.RequestOption) (res *pagination.Page[OAuthApplication], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "oauth_applications"
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

// List OAuth Applications
func (r *OAuthApplicationService) ListAutoPaging(ctx context.Context, query OAuthApplicationListParams, opts ...option.RequestOption) *pagination.PageAutoPager[OAuthApplication] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Create an OAuth Application via the Dashboard. Learn more about OAuth
// [here](https://increase.com/documentation/oauth).
type OAuthApplication struct {
	// The OAuth Application's identifier.
	ID string `json:"id,required"`
	// The OAuth Application's client_id. Use this to authenticate with the OAuth
	// Application.
	ClientID string `json:"client_id,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp when the OAuth
	// Application was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp when the OAuth
	// Application was deleted.
	DeletedAt time.Time `json:"deleted_at,required,nullable" format:"date-time"`
	// The name you chose for this OAuth Application.
	Name string `json:"name,required,nullable"`
	// Whether the application is active.
	Status OAuthApplicationStatus `json:"status,required"`
	// A constant representing the object's type. For this resource it will always be
	// `oauth_application`.
	Type OAuthApplicationType `json:"type,required"`
	JSON oauthApplicationJSON `json:"-"`
}

// oauthApplicationJSON contains the JSON metadata for the struct
// [OAuthApplication]
type oauthApplicationJSON struct {
	ID          apijson.Field
	ClientID    apijson.Field
	CreatedAt   apijson.Field
	DeletedAt   apijson.Field
	Name        apijson.Field
	Status      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OAuthApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthApplicationJSON) RawJSON() string {
	return r.raw
}

// Whether the application is active.
type OAuthApplicationStatus string

const (
	// The application is active and can be used by your users.
	OAuthApplicationStatusActive OAuthApplicationStatus = "active"
	// The application is deleted.
	OAuthApplicationStatusDeleted OAuthApplicationStatus = "deleted"
)

func (r OAuthApplicationStatus) IsKnown() bool {
	switch r {
	case OAuthApplicationStatusActive, OAuthApplicationStatusDeleted:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `oauth_application`.
type OAuthApplicationType string

const (
	OAuthApplicationTypeOAuthApplication OAuthApplicationType = "oauth_application"
)

func (r OAuthApplicationType) IsKnown() bool {
	switch r {
	case OAuthApplicationTypeOAuthApplication:
		return true
	}
	return false
}

type OAuthApplicationListParams struct {
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [OAuthApplicationListParams]'s query parameters as
// `url.Values`.
func (r OAuthApplicationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

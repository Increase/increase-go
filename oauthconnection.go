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

// OAuthConnectionService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOAuthConnectionService] method instead.
type OAuthConnectionService struct {
	Options []option.RequestOption
}

// NewOAuthConnectionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOAuthConnectionService(opts ...option.RequestOption) (r *OAuthConnectionService) {
	r = &OAuthConnectionService{}
	r.Options = opts
	return
}

// Retrieve an OAuth Connection
func (r *OAuthConnectionService) Get(ctx context.Context, oauthConnectionID string, opts ...option.RequestOption) (res *OAuthConnection, err error) {
	opts = append(r.Options[:], opts...)
	if oauthConnectionID == "" {
		err = errors.New("missing required oauth_connection_id parameter")
		return
	}
	path := fmt.Sprintf("oauth_connections/%s", oauthConnectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List OAuth Connections
func (r *OAuthConnectionService) List(ctx context.Context, query OAuthConnectionListParams, opts ...option.RequestOption) (res *pagination.Page[OAuthConnection], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "oauth_connections"
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

// List OAuth Connections
func (r *OAuthConnectionService) ListAutoPaging(ctx context.Context, query OAuthConnectionListParams, opts ...option.RequestOption) *pagination.PageAutoPager[OAuthConnection] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// When a user authorizes your OAuth application, an OAuth Connection object is
// created. Learn more about OAuth
// [here](https://increase.com/documentation/oauth).
type OAuthConnection struct {
	// The OAuth Connection's identifier.
	ID string `json:"id,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp when the OAuth
	// Connection was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp when the OAuth
	// Connection was deleted.
	DeletedAt time.Time `json:"deleted_at,required,nullable" format:"date-time"`
	// The identifier of the Group that has authorized your OAuth application.
	GroupID string `json:"group_id,required"`
	// The identifier of the OAuth application this connection is for.
	OAuthApplicationID string `json:"oauth_application_id,required"`
	// Whether the connection is active.
	Status OAuthConnectionStatus `json:"status,required"`
	// A constant representing the object's type. For this resource it will always be
	// `oauth_connection`.
	Type OAuthConnectionType `json:"type,required"`
	JSON oauthConnectionJSON `json:"-"`
}

// oauthConnectionJSON contains the JSON metadata for the struct [OAuthConnection]
type oauthConnectionJSON struct {
	ID                 apijson.Field
	CreatedAt          apijson.Field
	DeletedAt          apijson.Field
	GroupID            apijson.Field
	OAuthApplicationID apijson.Field
	Status             apijson.Field
	Type               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *OAuthConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthConnectionJSON) RawJSON() string {
	return r.raw
}

// Whether the connection is active.
type OAuthConnectionStatus string

const (
	OAuthConnectionStatusActive   OAuthConnectionStatus = "active"
	OAuthConnectionStatusInactive OAuthConnectionStatus = "inactive"
)

func (r OAuthConnectionStatus) IsKnown() bool {
	switch r {
	case OAuthConnectionStatusActive, OAuthConnectionStatusInactive:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `oauth_connection`.
type OAuthConnectionType string

const (
	OAuthConnectionTypeOAuthConnection OAuthConnectionType = "oauth_connection"
)

func (r OAuthConnectionType) IsKnown() bool {
	switch r {
	case OAuthConnectionTypeOAuthConnection:
		return true
	}
	return false
}

type OAuthConnectionListParams struct {
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
	// Filter results to only include OAuth Connections for a specific OAuth
	// Application.
	OAuthApplicationID param.Field[string]                          `query:"oauth_application_id"`
	Status             param.Field[OAuthConnectionListParamsStatus] `query:"status"`
}

// URLQuery serializes [OAuthConnectionListParams]'s query parameters as
// `url.Values`.
func (r OAuthConnectionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type OAuthConnectionListParamsStatus struct {
	// Filter to OAuth Connections by their status. By default, return only the
	// `active` ones. For GET requests, this should be encoded as a comma-delimited
	// string, such as `?in=one,two,three`.
	In param.Field[[]OAuthConnectionListParamsStatusIn] `query:"in"`
}

// URLQuery serializes [OAuthConnectionListParamsStatus]'s query parameters as
// `url.Values`.
func (r OAuthConnectionListParamsStatus) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type OAuthConnectionListParamsStatusIn string

const (
	OAuthConnectionListParamsStatusInActive   OAuthConnectionListParamsStatusIn = "active"
	OAuthConnectionListParamsStatusInInactive OAuthConnectionListParamsStatusIn = "inactive"
)

func (r OAuthConnectionListParamsStatusIn) IsKnown() bool {
	switch r {
	case OAuthConnectionListParamsStatusInActive, OAuthConnectionListParamsStatusInInactive:
		return true
	}
	return false
}

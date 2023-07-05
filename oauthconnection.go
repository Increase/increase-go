// File generated from our OpenAPI spec by Stainless.

package increase

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/apiquery"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/internal/shared"
	"github.com/increase/increase-go/option"
)

// OauthConnectionService contains methods and other services that help with
// interacting with the increase API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewOauthConnectionService] method
// instead.
type OauthConnectionService struct {
	Options []option.RequestOption
}

// NewOauthConnectionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOauthConnectionService(opts ...option.RequestOption) (r *OauthConnectionService) {
	r = &OauthConnectionService{}
	r.Options = opts
	return
}

// Retrieve an OAuth Connection
func (r *OauthConnectionService) Get(ctx context.Context, oauthConnectionID string, opts ...option.RequestOption) (res *OauthConnection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("oauth_connections/%s", oauthConnectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List OAuth Connections
func (r *OauthConnectionService) List(ctx context.Context, query OauthConnectionListParams, opts ...option.RequestOption) (res *shared.Page[OauthConnection], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
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
func (r *OauthConnectionService) ListAutoPaging(ctx context.Context, query OauthConnectionListParams, opts ...option.RequestOption) *shared.PageAutoPager[OauthConnection] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// When a user authorizes your OAuth application, an OAuth Connection object is
// created.
type OauthConnection struct {
	// The OAuth Connection's identifier.
	ID string `json:"id,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp when the OAuth
	// Connection was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The identifier of the Group that has authorized your OAuth application.
	GroupID string `json:"group_id,required"`
	// Whether the connection is active.
	Status OauthConnectionStatus `json:"status,required"`
	// A constant representing the object's type. For this resource it will always be
	// `oauth_connection`.
	Type OauthConnectionType `json:"type,required"`
	JSON oauthConnectionJSON
}

// oauthConnectionJSON contains the JSON metadata for the struct [OauthConnection]
type oauthConnectionJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	GroupID     apijson.Field
	Status      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OauthConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the connection is active.
type OauthConnectionStatus string

const (
	// The OAuth connection is active.
	OauthConnectionStatusActive OauthConnectionStatus = "active"
	// The OAuth connection is permanently deactivated.
	OauthConnectionStatusInactive OauthConnectionStatus = "inactive"
)

// A constant representing the object's type. For this resource it will always be
// `oauth_connection`.
type OauthConnectionType string

const (
	OauthConnectionTypeOauthConnection OauthConnectionType = "oauth_connection"
)

type OauthConnectionListParams struct {
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [OauthConnectionListParams]'s query parameters as
// `url.Values`.
func (r OauthConnectionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

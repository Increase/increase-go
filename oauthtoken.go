// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"net/http"

	"github.com/Increase/increase-go/internal/apijson"
	"github.com/Increase/increase-go/internal/param"
	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
)

// OAuthTokenService contains methods and other services that help with interacting
// with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOAuthTokenService] method instead.
type OAuthTokenService struct {
	Options []option.RequestOption
}

// NewOAuthTokenService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOAuthTokenService(opts ...option.RequestOption) (r *OAuthTokenService) {
	r = &OAuthTokenService{}
	r.Options = opts
	return
}

// Create an OAuth Token
func (r *OAuthTokenService) New(ctx context.Context, body OAuthTokenNewParams, opts ...option.RequestOption) (res *OAuthToken, err error) {
	opts = append(r.Options[:], opts...)
	path := "oauth/tokens"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// A token that is returned to your application when a user completes the OAuth
// flow and may be used to authenticate requests. Learn more about OAuth
// [here](/documentation/oauth).
type OAuthToken struct {
	// You may use this token in place of an API key to make OAuth requests on a user's
	// behalf.
	AccessToken string `json:"access_token,required"`
	// The type of OAuth token.
	TokenType OAuthTokenTokenType `json:"token_type,required"`
	// A constant representing the object's type. For this resource it will always be
	// `oauth_token`.
	Type OAuthTokenType `json:"type,required"`
	JSON oauthTokenJSON `json:"-"`
}

// oauthTokenJSON contains the JSON metadata for the struct [OAuthToken]
type oauthTokenJSON struct {
	AccessToken apijson.Field
	TokenType   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OAuthToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthTokenJSON) RawJSON() string {
	return r.raw
}

// The type of OAuth token.
type OAuthTokenTokenType string

const (
	OAuthTokenTokenTypeBearer OAuthTokenTokenType = "bearer"
)

func (r OAuthTokenTokenType) IsKnown() bool {
	switch r {
	case OAuthTokenTokenTypeBearer:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `oauth_token`.
type OAuthTokenType string

const (
	OAuthTokenTypeOAuthToken OAuthTokenType = "oauth_token"
)

func (r OAuthTokenType) IsKnown() bool {
	switch r {
	case OAuthTokenTypeOAuthToken:
		return true
	}
	return false
}

type OAuthTokenNewParams struct {
	// The credential you request in exchange for the code. In Production, this is
	// always `authorization_code`. In Sandbox, you can pass either enum value.
	GrantType param.Field[OAuthTokenNewParamsGrantType] `json:"grant_type,required"`
	// The public identifier for your application.
	ClientID param.Field[string] `json:"client_id"`
	// The secret that confirms you own the application. This is redundent given that
	// the request is made with your API key but it's a required component of OAuth
	// 2.0.
	ClientSecret param.Field[string] `json:"client_secret"`
	// The authorization code generated by the user and given to you as a query
	// parameter.
	Code param.Field[string] `json:"code"`
	// The production token you want to exchange for a sandbox token. This is only
	// available in Sandbox. Set `grant_type` to `production_token` to use this
	// parameter.
	ProductionToken param.Field[string] `json:"production_token"`
}

func (r OAuthTokenNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The credential you request in exchange for the code. In Production, this is
// always `authorization_code`. In Sandbox, you can pass either enum value.
type OAuthTokenNewParamsGrantType string

const (
	OAuthTokenNewParamsGrantTypeAuthorizationCode OAuthTokenNewParamsGrantType = "authorization_code"
	OAuthTokenNewParamsGrantTypeProductionToken   OAuthTokenNewParamsGrantType = "production_token"
)

func (r OAuthTokenNewParamsGrantType) IsKnown() bool {
	switch r {
	case OAuthTokenNewParamsGrantTypeAuthorizationCode, OAuthTokenNewParamsGrantTypeProductionToken:
		return true
	}
	return false
}
